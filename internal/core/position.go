package core

import (
	"errors"
	"fmt"
	"slices"
	"strconv"
	"strings"
	"unicode"
)

type PlayerTurn bool

//goland:noinspection GoBoolExpressions
const (
	White PlayerTurn = iota%2 == 0
	Black
)

func (pt PlayerTurn) String() string {
	if pt {
		return "White"
	}
	return "Black"
}

type PiecePosition struct {
	FileIndex, RankIndex int
}

type Move struct {
	Target    PiecePosition
	PromoteTo Piece
}

type KingsPosition struct {
	WhiteKing PiecePosition
	BlackKing PiecePosition
}

type State struct {
	turn                  PlayerTurn
	grid                  [8][8]Piece
	kingsPosition         KingsPosition
	castlingRights        CastlingRights
	enPassantTargetSquare PiecePosition
	halfMoveClock         int
	fullMoveNumber        int
}

func (s State) String() string {
	var sb strings.Builder

	horizontal := "  +---+---+---+---+---+---+---+---+\n"

	sb.WriteString(horizontal)

	for rank := 7; rank >= 0; rank-- {
		sb.WriteString(fmt.Sprintf("%d |", rank+1))

		for file := 0; file < 8; file++ {
			piece := s.grid[file][rank]

			symbol := "."
			if piece != nil {
				switch piece.Value() {
				case 1:
					symbol = "P"
				case 2:
					symbol = "N"
				case 4:
					symbol = "B"
				case 8:
					symbol = "R"
				case 16:
					symbol = "Q"
				case 32:
					symbol = "K"
				case 64:
					symbol = "p"
				case 128:
					symbol = "n"
				case 256:
					symbol = "b"
				case 512:
					symbol = "r"
				case 1024:
					symbol = "q"
				case 2048:
					symbol = "k"
				}
			}

			sb.WriteString(fmt.Sprintf(" %s |", symbol))
		}

		sb.WriteString("\n")
		sb.WriteString(horizontal)
	}

	// file labels
	sb.WriteString("    a   b   c   d   e   f   g   h\n")

	return sb.String()
}

func (s State) MakeMove(sourceSquare PiecePosition, targetSquare PiecePosition, promoteTo Piece) (State, error) {
	chosenPiece := s.GetPieceAt(sourceSquare.FileIndex, sourceSquare.RankIndex)
	if chosenPiece == nil {
		return State{}, fmt.Errorf("no piece exist at the chosen position: %v", sourceSquare)
	}

	// Ensures that player selects only his piece
	if s.turn == White && chosenPiece.Value()&AnyBlackPiece > 0 ||
		s.turn == Black && chosenPiece.Value()&AnyWhitePiece > 0 {
		return State{}, fmt.Errorf("invalid piece selected, you can only select your pieces")
	}

	// Ensures that player is making a valid move
	moves := chosenPiece.GetMoves(s, sourceSquare)

	if !slices.Contains(moves, Move{
		Target:    targetSquare,
		PromoteTo: promoteTo,
	}) {
		return State{}, fmt.Errorf("invalid move for given piece")
	}

	/*
		We'll now perform the following steps in the new game state -
		1. Give turn to the other player
		2. Reset En-Passant Target Square to either nil or a new target square
		3. Vacant the square where the piece was originally placed and occupy the target square
		4. Reset the Half move clock if the move isn't a capture or a pawn move
	*/
	newState := s
	newState.turn = !s.turn

	targetPiece := newState.GetPieceAt(targetSquare.FileIndex, targetSquare.RankIndex)
	newState.grid[targetSquare.FileIndex][targetSquare.RankIndex] = chosenPiece
	newState.grid[sourceSquare.FileIndex][sourceSquare.RankIndex] = nil

	newState.enPassantTargetSquare = PiecePosition{-1, -1}
	// En passant is only available for one move, so we reset it once the move is made.
	if chosenPiece.Value()&AnyPawn > 0 {
		chosenPawn := chosenPiece.(pawn)
		newState.enPassantTargetSquare = chosenPawn.getEnPassantTargetSquare(s, sourceSquare.RankIndex, targetSquare)

		// If en passant then the opponent pawn to be captured will be on the same file but previous rank
		// of target rank from the direction of player e.g.
		// 1. (White to Move) If target square is e6, e5 will go vacant	due to en passant
		// 2. (Black to Move) If target square is e3, e4 will go vacant due to en passant
		if s.enPassantTargetSquare.FileIndex != -1 &&
			targetSquare.RankIndex == s.enPassantTargetSquare.RankIndex &&
			targetSquare.FileIndex == s.enPassantTargetSquare.FileIndex {
			direction := -1
			if s.turn == Black {
				direction = 1
			}
			newState.grid[targetSquare.FileIndex][targetSquare.RankIndex+direction] = nil
		}

		// Handle Promotion
		promotionRank := 7
		if s.turn == Black {
			promotionRank = 0
		}

		if targetSquare.RankIndex == promotionRank {
			newState.grid[targetSquare.FileIndex][targetSquare.RankIndex] = promoteTo
		}
	}

	// This resets half move counter in case of capture or pawn move
	if targetPiece != nil || chosenPiece == WhitePawn || chosenPiece == BlackPawn {
		newState.halfMoveClock = 0
	}

	kingStartRank := 0
	rookKsFile := 7
	rookQsFile := 0

	if !s.turn == White {
		kingStartRank = 7
	}

	if chosenPiece == WhiteKing || chosenPiece == BlackKing {
		if s.turn == White {
			newState.castlingRights &^= WhiteCanKsCastle | WhiteCanQsCastle
		} else {
			newState.castlingRights &^= BlackCanKsCastle | BlackCanQsCastle
		}

		fileDiff := sourceSquare.FileIndex - targetSquare.FileIndex
		if fileDiff == 2 || fileDiff == -2 {
			// king side castle
			if targetSquare.FileIndex == 6 && targetSquare.RankIndex == kingStartRank {
				newState.grid[5][kingStartRank] = newState.grid[7][kingStartRank]
				newState.grid[7][kingStartRank] = nil
			}

			// queen side castle
			if targetSquare.FileIndex == 2 && targetSquare.RankIndex == kingStartRank {
				newState.grid[3][kingStartRank] = newState.grid[0][kingStartRank]
				newState.grid[0][kingStartRank] = nil
			}
		}

	}

	if chosenPiece == WhiteRook || chosenPiece == BlackRook {
		if sourceSquare.FileIndex == rookKsFile && sourceSquare.RankIndex == kingStartRank {
			if s.turn == White {
				newState.castlingRights &^= WhiteCanKsCastle
			} else {
				newState.castlingRights &^= BlackCanKsCastle
			}
		}

		if sourceSquare.FileIndex == rookQsFile && sourceSquare.RankIndex == kingStartRank {
			if s.turn == White {
				newState.castlingRights &^= WhiteCanQsCastle
			} else {
				newState.castlingRights &^= BlackCanQsCastle
			}
		}
	}

	// Revoke castling rights if a rook is captured on its starting square
	if targetPiece == WhiteRook {
		if targetSquare.FileIndex == 7 && targetSquare.RankIndex == 0 {
			newState.castlingRights &^= WhiteCanKsCastle
		} else if targetSquare.FileIndex == 0 && targetSquare.RankIndex == 0 {
			newState.castlingRights &^= WhiteCanQsCastle
		}
	} else if targetPiece == BlackRook {
		if targetSquare.FileIndex == 7 && targetSquare.RankIndex == 7 {
			newState.castlingRights &^= BlackCanKsCastle
		} else if targetSquare.FileIndex == 0 && targetSquare.RankIndex == 7 {
			newState.castlingRights &^= BlackCanQsCastle
		}
	}

	return newState, nil
}

func (s State) GetPieceAt(fileIndex, rankIndex int) Piece {
	return s.grid[fileIndex][rankIndex]
}

func (s State) GetAttackers(fileIndex, rankIndex int) map[[2]int]Piece {
	var attackers = make(map[[2]int]Piece)
	var possibleAttackers, direction int

	if s.turn == White {
		direction = 1
	} else {
		direction = -1
	}

	jumps := [][2]int{
		{2, 1},
		{2, -1},
		{1, 2},
		{1, -2},
		{-2, 1},
		{-2, -1},
		{-1, 2},
		{-1, -2},
	}

	diagonals := [][2]int{
		{1, 1},
		{-1, 1},
		{-1, -1},
		{1, -1},
	}

	lines := [][2]int{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
	}

	pawnDeltas := [][2]int{
		{1, direction},
		{-1, direction},
	}

	for _, jump := range jumps {
		attackerFileIndex := fileIndex + jump[0]
		attackerRankIndex := rankIndex + jump[1]

		if attackerFileIndex > 7 || attackerRankIndex > 7 || attackerFileIndex < 0 || attackerRankIndex < 0 {
			continue
		}

		attackerPiece := s.GetPieceAt(attackerFileIndex, attackerRankIndex)
		if attackerPiece == nil {
			continue
		}
		if s.isFriendlyCapture(attackerFileIndex, attackerRankIndex) {
			continue
		}
		if attackerPiece == BlackKnight && s.turn == White || attackerPiece == WhiteKnight && s.turn == Black {
			attackers[[2]int{attackerFileIndex, attackerRankIndex}] = attackerPiece
		}
	}

	for _, line := range lines {
		for multiplier := 1; ; multiplier++ {
			if multiplier == 1 {
				if s.turn == White {
					possibleAttackers = BlackRookQueenOrKing
				} else {
					possibleAttackers = WhiteRookQueenOrKing
				}
			} else {
				if s.turn == White {
					possibleAttackers = BlackRookOrQueen
				} else {
					possibleAttackers = WhiteRookOrQueen
				}
			}
			attackerFileIndex := fileIndex + line[0]*multiplier
			attackerRankIndex := rankIndex + line[1]*multiplier

			// Out of bounds - so the line is done and next line should be picked
			if attackerFileIndex >= 8 || attackerRankIndex >= 8 || attackerFileIndex <= -1 || attackerRankIndex <= -1 {
				break
			}
			attackerPiece := s.GetPieceAt(attackerFileIndex, attackerRankIndex)
			// The square ain't vacant, so we can not proceed on this line any further
			// provided that we can still capture this square if an enemy piece exists on it
			if attackerPiece == nil {
				continue
			}
			if attackerPiece.Value()&possibleAttackers > 0 {
				attackers[[2]int{attackerFileIndex, attackerRankIndex}] = attackerPiece
			}
			break
		}
	}

	for _, diagonal := range diagonals {
		for multiplier := 1; ; multiplier++ {
			if multiplier == 1 {
				if s.turn == White {
					possibleAttackers = BlackBishopQueenOrKing
				} else {
					possibleAttackers = WhiteBishopQueenOrKing
				}
			} else {
				if s.turn == White {
					possibleAttackers = BlackBishopOrQueen
				} else {
					possibleAttackers = WhiteBishopOrQueen
				}
			}
			attackerFileIndex := fileIndex + diagonal[0]*multiplier
			attackerRankIndex := rankIndex + diagonal[1]*multiplier

			// Out of bounds - so the line is done and next diagonal should be picked
			if attackerFileIndex >= 8 || attackerRankIndex >= 8 || attackerFileIndex <= -1 || attackerRankIndex <= -1 {
				break
			}

			attackerPiece := s.GetPieceAt(attackerFileIndex, attackerRankIndex)

			// The square ain't vacant, so we can not proceed on this diagonal any further
			// provided that we can still capture this square if an enemy piece exists on it
			if attackerPiece == nil {
				continue
			}

			if attackerPiece.Value()&possibleAttackers > 0 {
				attackers[[2]int{attackerFileIndex, attackerRankIndex}] = attackerPiece
			}
			break
		}
	}

	for _, delta := range pawnDeltas {
		attackerFileIndex := fileIndex + delta[0]
		attackerRankIndex := rankIndex + delta[1]
		if attackerFileIndex < 0 || attackerRankIndex < 0 || attackerFileIndex >= 8 || attackerRankIndex >= 8 {
			continue
		}
		attackerPiece := s.GetPieceAt(attackerFileIndex, attackerRankIndex)
		if attackerPiece == nil {
			continue
		}
		if s.turn == White && attackerPiece == BlackPawn ||
			s.turn == Black && attackerPiece == WhitePawn {
			attackers[[2]int{attackerFileIndex, attackerRankIndex}] = attackerPiece
		}
	}

	return attackers
}

func (s State) GetKingAttackers() map[[2]int]Piece {
	kingsPosition := s.GetKingsPosition()
	var kingFileIndex, kingRankIndex int
	if s.turn == White {
		kingFileIndex = kingsPosition.WhiteKing.FileIndex
		kingRankIndex = kingsPosition.WhiteKing.RankIndex
	} else {
		kingFileIndex = kingsPosition.BlackKing.FileIndex
		kingRankIndex = kingsPosition.BlackKing.RankIndex
	}
	attackers := s.GetAttackers(kingFileIndex, kingRankIndex)
	return attackers
}

func (s State) isKingInCheck(attackers map[[2]int]Piece) bool {
	return len(attackers) > 0
}

func (s State) SetTurn(turn PlayerTurn) State {
	s.turn = turn
	return s
}

func (s State) GetKingsPosition() KingsPosition {
	whiteKingFileIndex := -1
	whiteKingRankIndex := -1
	blackKingFileIndex := -1
	blackKingRankIndex := -1
	var piece Piece

	/** This loop is performed to find both kings position as early as possible by
	* 	searching opposite color king on the mirror square (mirror square of e1 is e8)
	*	i.e. If we are searching for white king on 1st rank, we simultaneously search for
	* 	black king on 8th rank as the probability of finding respective kings is highest
	* 	on these ranks. The likeliness of finding kings decreases as we move forward with
	* 	the ranks and such ranks should only be visited if the kings are not already found.
	 */
gridLoop:
	for rankIndex := 0; rankIndex < 8; rankIndex++ {
		for fileIndex := 0; fileIndex < 8; fileIndex++ {
			piece = s.GetPieceAt(fileIndex, rankIndex)
			if whiteKingFileIndex < 0 && piece == WhiteKing {
				whiteKingFileIndex = fileIndex
				whiteKingRankIndex = rankIndex
			}

			piece = s.GetPieceAt(fileIndex, 8-(rankIndex+1))
			if blackKingFileIndex < 0 && piece == BlackKing {
				blackKingFileIndex = fileIndex
				blackKingRankIndex = 8 - (rankIndex + 1)

			}

			if whiteKingRankIndex > -1 && blackKingRankIndex > -1 {
				break gridLoop
			}
		}
	}

	return KingsPosition{
		WhiteKing: PiecePosition{
			FileIndex: whiteKingFileIndex,
			RankIndex: whiteKingRankIndex,
		},
		BlackKing: PiecePosition{
			FileIndex: blackKingFileIndex,
			RankIndex: blackKingRankIndex,
		},
	}
}

func (s State) isKingInCheckAfter(sourceFileIndex, sourceRankIndex, targetFileIndex, targetRankIndex int) bool {
	piece := s.GetPieceAt(sourceFileIndex, sourceRankIndex)
	futureGrid := s.grid
	futureGrid[sourceFileIndex][sourceRankIndex] = nil
	futureGrid[targetFileIndex][targetRankIndex] = piece
	futureState := State{
		turn: s.turn,
		grid: futureGrid,
	}
	attackers := futureState.GetKingAttackers()
	return futureState.isKingInCheck(attackers)
}

func (s State) isKingInCheckAfterEnPassant(sourceFileIndex, sourceRankIndex, targetFileIndex, targetRankIndex int) bool {
	piece := s.GetPieceAt(sourceFileIndex, sourceRankIndex)
	futureGrid := s.grid

	futureGrid[sourceFileIndex][sourceRankIndex] = nil
	futureGrid[targetFileIndex][targetRankIndex] = piece
	futureGrid[targetFileIndex][sourceRankIndex] = nil

	futureState := State{
		turn: s.turn,
		grid: futureGrid,
	}
	attackers := futureState.GetKingAttackers()
	return futureState.isKingInCheck(attackers)
}

func (s State) isFriendlyCapture(fileIndex, rankIndex int) bool {
	if s.turn == White && (s.GetPieceAt(fileIndex, rankIndex).Value()&AnyWhitePiece != 0) {
		return true
	}
	if s.turn == Black && (s.GetPieceAt(fileIndex, rankIndex).Value()&AnyBlackPiece != 0) {
		return true
	}
	return false
}

func (s State) GetAllPossibleMoves() map[[2]int][]Move {
	piecesEncountered := 0
	possibleMoves := make(map[[2]int][]Move)
gridLoop:
	for fileIndex := 0; fileIndex < 8; fileIndex++ {
		for rankIndex := 0; rankIndex < 8; rankIndex++ {
			piece := s.GetPieceAt(fileIndex, rankIndex)
			if piece == nil ||
				(s.turn == White && (piece.Value()&AnyBlackPiece > 0)) ||
				(s.turn == Black && (piece.Value()&AnyWhitePiece > 0)) {
				continue
			}

			piecesEncountered++
			moves := piece.GetMoves(s, PiecePosition{FileIndex: fileIndex, RankIndex: rankIndex})
			if len(moves) > 0 {
				possibleMoves[[2]int{fileIndex, rankIndex}] = moves
			}

			if piecesEncountered >= 16 {
				break gridLoop
			}
		}
	}

	return possibleMoves
}

func (s State) ToFEN() string {
	var sb strings.Builder

	// 1. Board
	for rank := 7; rank >= 0; rank-- {
		empty := 0

		for file := 0; file < 8; file++ {
			piece := s.grid[file][rank]

			if piece == nil {
				empty++
				continue
			}

			if empty > 0 {
				sb.WriteString(fmt.Sprintf("%d", empty))
				empty = 0
			}

			switch piece.Value() {
			case 1:
				sb.WriteString("P")
			case 2:
				sb.WriteString("N")
			case 4:
				sb.WriteString("B")
			case 8:
				sb.WriteString("R")
			case 16:
				sb.WriteString("Q")
			case 32:
				sb.WriteString("K")
			case 64:
				sb.WriteString("p")
			case 128:
				sb.WriteString("n")
			case 256:
				sb.WriteString("b")
			case 512:
				sb.WriteString("r")
			case 1024:
				sb.WriteString("q")
			case 2048:
				sb.WriteString("k")
			}
		}

		if empty > 0 {
			sb.WriteString(fmt.Sprintf("%d", empty))
		}

		if rank > 0 {
			sb.WriteString("/")
		}
	}

	// 2. Turn
	if s.turn == White {
		sb.WriteString(" w ")
	} else {
		sb.WriteString(" b ")
	}

	// 3. Castling rights
	castling := ""
	if s.castlingRights&WhiteCanKsCastle > 0 {
		castling += "K"
	}
	if s.castlingRights&WhiteCanQsCastle > 0 {
		castling += "Q"
	}
	if s.castlingRights&BlackCanKsCastle > 0 {
		castling += "k"
	}
	if s.castlingRights&BlackCanQsCastle > 0 {
		castling += "q"
	}

	if castling == "" {
		castling = "-"
	}
	sb.WriteString(castling + " ")

	// 4. En passant
	if s.enPassantTargetSquare.FileIndex == -1 {
		sb.WriteString("- ")
	} else {
		file := 'a' + rune(s.enPassantTargetSquare.FileIndex)
		rank := s.enPassantTargetSquare.RankIndex + 1
		sb.WriteString(fmt.Sprintf("%c%d ", file, rank))
	}

	// 5. Halfmove clock
	sb.WriteString(fmt.Sprintf("%d ", s.halfMoveClock))

	// 6. Fullmove number
	fullMove := s.fullMoveNumber
	if fullMove == 0 {
		fullMove = 1 // fallback if not tracked yet
	}
	sb.WriteString(fmt.Sprintf("%d", fullMove))

	return sb.String()
}

// NewGameStateFromFEN parses a FEN safely and returns an error on malicious/invalid input.
func NewGameStateFromFEN(fen string) (State, error) {
	var grid [8][8]Piece

	parts := strings.Split(strings.TrimSpace(fen), " ")
	if len(parts) < 4 {
		return State{}, errors.New("invalid FEN: requires at least 4 segments (placement, turn, castling, en passant)")
	}

	rank := 7
	file := 0

	var whiteKings, blackKings int

	for _, r := range parts[0] {
		if r == '/' {
			rank--
			file = 0
			if rank < 0 {
				return State{}, errors.New("invalid FEN: board has more than 8 ranks")
			}
			continue
		}

		if unicode.IsDigit(r) {
			file += int(r - '0')
			if file > 8 {
				return State{}, errors.New("invalid FEN: rank contains more than 8 squares")
			}
			continue
		}

		if file >= 8 {
			return State{}, errors.New("invalid FEN: rank contains more than 8 squares")
		}

		var piece Piece
		switch r {
		case 'P':
			piece = WhitePawn
		case 'N':
			piece = WhiteKnight
		case 'B':
			piece = WhiteBishop
		case 'R':
			piece = WhiteRook
		case 'Q':
			piece = WhiteQueen
		case 'K':
			piece = WhiteKing
			whiteKings++
		case 'p':
			piece = BlackPawn
		case 'n':
			piece = BlackKnight
		case 'b':
			piece = BlackBishop
		case 'r':
			piece = BlackRook
		case 'q':
			piece = BlackQueen
		case 'k':
			piece = BlackKing
			blackKings++
		default:
			return State{}, fmt.Errorf("invalid FEN: unknown piece character '%c'", r)
		}

		grid[file][rank] = piece
		file++
	}

	// Basic Sanity Check to prevent mathematical anomalies
	if whiteKings != 1 || blackKings != 1 {
		return State{}, errors.New("invalid FEN: board must have exactly one white king and one black king")
	}

	turn := White
	if parts[1] == "b" {
		turn = Black
	}

	var castlingRights CastlingRights
	if parts[2] != "-" {
		if strings.ContainsRune(parts[2], 'K') {
			castlingRights |= WhiteCanKsCastle
		}
		if strings.ContainsRune(parts[2], 'Q') {
			castlingRights |= WhiteCanQsCastle
		}
		if strings.ContainsRune(parts[2], 'k') {
			castlingRights |= BlackCanKsCastle
		}
		if strings.ContainsRune(parts[2], 'q') {
			castlingRights |= BlackCanQsCastle
		}
	}

	epSquare := PiecePosition{FileIndex: -1, RankIndex: -1}
	if parts[3] != "-" && len(parts[3]) == 2 {
		epFile := int(parts[3][0] - 'a')
		epRank := int(parts[3][1] - '1')

		// Ensure En Passant target is on the board
		if epFile >= 0 && epFile <= 7 && epRank >= 0 && epRank <= 7 {
			epSquare = PiecePosition{FileIndex: epFile, RankIndex: epRank}
		}
	}

	state := NewGameState(turn, grid, castlingRights, epSquare)

	if len(parts) > 4 {
		if hm, err := strconv.Atoi(parts[4]); err == nil {
			state.halfMoveClock = hm
		}
	}
	if len(parts) > 5 {
		if fm, err := strconv.Atoi(parts[5]); err == nil {
			state.fullMoveNumber = fm
		}
	} else {
		state.fullMoveNumber = 1
	}

	return state, nil
}

func NewGameState(turn PlayerTurn, grid [8][8]Piece, castlingRights CastlingRights, enPassantTargetSquare PiecePosition) State {
	state := State{
		turn:                  turn,
		grid:                  grid,
		castlingRights:        castlingRights,
		enPassantTargetSquare: enPassantTargetSquare,
	}
	state.kingsPosition = state.GetKingsPosition()
	return state
}
