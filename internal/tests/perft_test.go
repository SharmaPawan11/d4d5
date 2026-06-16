//go:build perft

package tests

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/SharmaPawan11/d4d5/internal/core"
)

type PerftCounters struct {
	Nodes        int
	Captures     int
	Castles      int
	Checks       int
	Promotions   int
	EnPassants   int
	DoubleChecks int
	CheckMates   int
}

func TestAutomatedPerftSuite(t *testing.T) {
	fmt.Println("Starting Automated Perft Build Pipeline...")

	// Start the global timer
	defer func(start time.Time) {
		fmt.Printf("\n========================================\n🏁 TOTAL SUITE TIME: %s\n========================================\n", time.Since(start))
	}(time.Now())

	file, err := os.Open("perftsuite.epd")
	if err != nil {
		t.Fatalf("Could not open perftsuite.epd. Ensure it is in the internal/tests/ directory: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	positionNumber := 1

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		parts := strings.Split(line, ";")
		fen := strings.TrimSpace(parts[0])

		state, err := core.NewGameStateFromFEN(fen)
		if err != nil {
			t.Fatalf("Position %d: Unable to create game state from FEN: %s", positionNumber, err)
		}

		// Beautiful header for the position
		fmt.Printf("\nTesting Position %d: %s\n", positionNumber, fen)

		for i := 1; i < len(parts); i++ {
			depthInfo := strings.Fields(strings.TrimSpace(parts[i]))
			if len(depthInfo) < 2 {
				continue
			}

			depthStr := strings.TrimPrefix(depthInfo[0], "D")
			depth, err := strconv.Atoi(depthStr)
			if err != nil {
				continue
			}
			expectedNodes, _ := strconv.Atoi(depthInfo[1])

			start := time.Now()
			actualNodes := RunPerftFastConcurrent(state, depth)
			elapsed := time.Since(start)

			if actualNodes == expectedNodes {
				// Use fmt.Printf to bypass the ugly t.Logf formatting
				fmt.Printf(" [PASS] Depth %d | Nodes: %d | Time: %s\n", depth, actualNodes, elapsed)
			} else {
				fmt.Printf(" [FAIL] Depth %d | Expected: %d | Got: %d\n", depth, expectedNodes, actualNodes)
				fmt.Println("        -> Node mismatch detected. Triggering Detailed Diagnostic Analysis...")

				RunDiagnosticFallback(t, fen, depth)

				// Officially fail the test so CI/CD turns red
				t.Errorf("Position %d Depth %d failed", positionNumber, depth)
			}
		}
		positionNumber++
	}
}

func RunPerftFastConcurrent(state core.State, targetDepth int) int {
	if targetDepth == 0 {
		return 1
	}

	moves := state.GetAllPossibleMoves()

	if targetDepth == 1 {
		total := 0
		for _, targets := range moves {
			total += len(targets)
		}
		return total
	}

	type moveData struct {
		Source [2]int
		Target core.Move
	}
	var rootMoves []moveData
	for source, targets := range moves {
		for _, target := range targets {
			rootMoves = append(rootMoves, moveData{Source: source, Target: target})
		}
	}

	var wg sync.WaitGroup
	results := make(chan int, len(rootMoves))

	for _, rm := range rootMoves {
		wg.Add(1)
		go func(s core.State, src [2]int, tgt core.Move) {
			defer wg.Done()
			newState, _ := s.MakeMove(
				core.PiecePosition{FileIndex: src[0], RankIndex: src[1]},
				core.PiecePosition{FileIndex: tgt.Target.FileIndex, RankIndex: tgt.Target.RankIndex},
				tgt.PromoteTo,
			)
			results <- recursivePerftFast(newState, targetDepth-1)
		}(state, rm.Source, rm.Target)
	}

	wg.Wait()
	close(results)

	totalNodes := 0
	for nodes := range results {
		totalNodes += nodes
	}
	return totalNodes
}

func recursivePerftFast(state core.State, depthLeft int) int {
	moves := state.GetAllPossibleMoves()

	if depthLeft == 1 {
		nodes := 0
		for _, targets := range moves {
			nodes += len(targets)
		}
		return nodes
	}

	nodes := 0
	for source, targets := range moves {
		for _, target := range targets {
			newState, _ := state.MakeMove(
				core.PiecePosition{FileIndex: source[0], RankIndex: source[1]},
				core.PiecePosition{FileIndex: target.Target.FileIndex, RankIndex: target.Target.RankIndex},
				target.PromoteTo,
			)
			nodes += recursivePerftFast(newState, depthLeft-1)
		}
	}
	return nodes
}

func RunDiagnosticFallback(t *testing.T, fen string, targetDepth int) {
	state, err := core.NewGameStateFromFEN(fen)
	if err != nil {
		t.Fatalf("Unable to create game state from FEN: %v", err)
	}

	actualStats := RunPerftDetailedConcurrent(state, targetDepth)
	expectedStats := ParseDetailedEPD("perftsuite_detailed.epd", fen, targetDepth)

	fmt.Println("\n========================================")
	fmt.Printf(" DIAGNOSTIC REPORT FOR DEPTH %d\n", targetDepth)
	fmt.Println("========================================")

	printDiff(t, "Nodes", expectedStats.Nodes, actualStats.Nodes)
	printDiff(t, "Captures", expectedStats.Captures, actualStats.Captures)
	printDiff(t, "En Passants", expectedStats.EnPassants, actualStats.EnPassants)
	printDiff(t, "Castles", expectedStats.Castles, actualStats.Castles)
	printDiff(t, "Promotions", expectedStats.Promotions, actualStats.Promotions)
	printDiff(t, "Checks", expectedStats.Checks, actualStats.Checks)
	printDiff(t, "Double Checks", expectedStats.DoubleChecks, actualStats.DoubleChecks)
	fmt.Println("========================================")
}

func printDiff(t *testing.T, name string, expected int, actual int) {
	if expected == actual {
		fmt.Printf("[OK]   %-15s: %d\n", name, actual)
	} else {
		fmt.Printf("[FAIL] %-15s: Expected %d, Got %d (Diff: %d)\n", name, expected, actual, actual-expected)
	}
}

func ParseDetailedEPD(filename string, targetFen string, targetDepth int) PerftCounters {
	var stats PerftCounters
	file, err := os.Open(filename)
	if err != nil {
		return stats // Will just return 0s if file isn't found
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, targetFen) {
			parts := strings.Split(line, ";")
			for i := 1; i < len(parts); i++ {
				depthInfo := strings.Fields(strings.TrimSpace(parts[i]))
				if len(depthInfo) == 0 {
					continue
				}

				depthStr := strings.TrimPrefix(depthInfo[0], "D")
				d, _ := strconv.Atoi(depthStr)

				if d == targetDepth && len(depthInfo) >= 9 {
					stats.Nodes, _ = strconv.Atoi(depthInfo[1])
					stats.Captures, _ = strconv.Atoi(depthInfo[2])
					stats.EnPassants, _ = strconv.Atoi(depthInfo[3])
					stats.Castles, _ = strconv.Atoi(depthInfo[4])
					stats.Promotions, _ = strconv.Atoi(depthInfo[5])
					stats.Checks, _ = strconv.Atoi(depthInfo[6])
					stats.DoubleChecks, _ = strconv.Atoi(depthInfo[8])
					return stats
				}
			}
		}
	}
	return stats
}

func RunPerftDetailedConcurrent(state core.State, targetDepth int) PerftCounters {
	var totalCounters PerftCounters
	moves := state.GetAllPossibleMoves()

	type moveData struct {
		Source [2]int
		Target core.Move
	}
	var rootMoves []moveData
	for source, targets := range moves {
		for _, target := range targets {
			rootMoves = append(rootMoves, moveData{Source: source, Target: target})
		}
	}

	var wg sync.WaitGroup
	results := make(chan PerftCounters, len(rootMoves))

	for _, rm := range rootMoves {
		wg.Add(1)

		go func(s core.State, src [2]int, tgt core.Move) {
			defer wg.Done()
			var localCounters PerftCounters

			chosenPiece := s.GetPieceAt(src[0], src[1])
			targetPiece := s.GetPieceAt(tgt.Target.FileIndex, tgt.Target.RankIndex)

			isPromotion := tgt.PromoteTo != nil
			isEnPassant := (chosenPiece.Value()&core.AnyPawn > 0) && (src[0] != tgt.Target.FileIndex) && targetPiece == nil
			isCapture := targetPiece != nil || isEnPassant

			fileDiff := src[0] - tgt.Target.FileIndex
			if fileDiff < 0 {
				fileDiff = -fileDiff
			}
			isCastle := (chosenPiece.Value()&core.AnyKing > 0) && fileDiff == 2

			newState, _ := s.MakeMove(
				core.PiecePosition{FileIndex: src[0], RankIndex: src[1]},
				core.PiecePosition{FileIndex: tgt.Target.FileIndex, RankIndex: tgt.Target.RankIndex},
				tgt.PromoteTo,
			)

			if targetDepth == 1 {
				localCounters.Nodes++
				if isCapture {
					localCounters.Captures++
				}
				if isCastle {
					localCounters.Castles++
				}
				if isPromotion {
					localCounters.Promotions++
				}
				if isEnPassant {
					localCounters.EnPassants++
				}

				attackers := newState.GetKingAttackers()
				numAttackers := len(attackers)
				if numAttackers > 0 {
					localCounters.Checks++
					if numAttackers == 2 {
						localCounters.DoubleChecks++
					}
					respMoves := newState.GetAllPossibleMoves()
					if len(respMoves) == 0 {
						localCounters.CheckMates++
					}
				}
			} else {
				recursivePerftDetailed(newState, targetDepth-1, &localCounters)
			}

			results <- localCounters
		}(state, rm.Source, rm.Target)
	}

	wg.Wait()
	close(results)

	for c := range results {
		totalCounters.Nodes += c.Nodes
		totalCounters.Captures += c.Captures
		totalCounters.Castles += c.Castles
		totalCounters.Checks += c.Checks
		totalCounters.DoubleChecks += c.DoubleChecks
		totalCounters.Promotions += c.Promotions
		totalCounters.EnPassants += c.EnPassants
		totalCounters.CheckMates += c.CheckMates
	}

	return totalCounters
}

func recursivePerftDetailed(state core.State, depthLeft int, counters *PerftCounters) {
	if depthLeft == 0 {
		return
	}

	moves := state.GetAllPossibleMoves()

	for source, targets := range moves {
		for _, target := range targets {
			chosenPiece := state.GetPieceAt(source[0], source[1])
			targetPiece := state.GetPieceAt(target.Target.FileIndex, target.Target.RankIndex)

			isPromotion := target.PromoteTo != nil
			isEnPassant := (chosenPiece.Value()&core.AnyPawn > 0) && (source[0] != target.Target.FileIndex) && targetPiece == nil
			isCapture := targetPiece != nil || isEnPassant

			fileDiff := source[0] - target.Target.FileIndex
			if fileDiff < 0 {
				fileDiff = -fileDiff
			}
			isCastle := (chosenPiece.Value()&core.AnyKing > 0) && fileDiff == 2

			newState, _ := state.MakeMove(
				core.PiecePosition{FileIndex: source[0], RankIndex: source[1]},
				core.PiecePosition{FileIndex: target.Target.FileIndex, RankIndex: target.Target.RankIndex},
				target.PromoteTo,
			)

			if depthLeft == 1 {
				counters.Nodes++
				if isCapture {
					counters.Captures++
				}
				if isCastle {
					counters.Castles++
				}
				if isPromotion {
					counters.Promotions++
				}
				if isEnPassant {
					counters.EnPassants++
				}

				attackers := newState.GetKingAttackers()
				numAttackers := len(attackers)
				if numAttackers > 0 {
					counters.Checks++
					if numAttackers == 2 {
						counters.DoubleChecks++
					}
					respMoves := newState.GetAllPossibleMoves()
					if len(respMoves) == 0 {
						counters.CheckMates++
					}
				}
			} else {
				recursivePerftDetailed(newState, depthLeft-1, counters)
			}
		}
	}
}
