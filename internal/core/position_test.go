package core

import (
	"strings"
	"testing"
)

func TestGetPieceAt(t *testing.T) {
	var grid = [8][8]Piece{
		{WhiteRook, WhitePawn, nil, nil, nil, nil, BlackPawn, BlackRook},
		{nil, WhitePawn, nil, nil, nil, nil, BlackPawn, nil},
		{WhiteBishop, WhitePawn, nil, nil, nil, nil, BlackPawn, BlackBishop},
		{nil, WhitePawn, nil, nil, nil, nil, BlackPawn, nil},
		{nil, WhitePawn, nil, nil, nil, nil, BlackPawn, nil},
		{WhiteBishop, WhitePawn, nil, nil, nil, nil, BlackPawn, BlackBishop},
		{nil, WhitePawn, nil, nil, nil, nil, BlackPawn, nil},
		{WhiteRook, WhitePawn, nil, nil, nil, nil, BlackPawn, BlackRook},
	}
	state := NewGameState(White, grid, 0, PiecePosition{-1, -1})
	fileIndex := 0
	rankIndex := 0
	piece := state.GetPieceAt(1, 6)
	if piece != BlackPawn {
		t.Errorf("Invalid piece returned at %v,%v. Expected :%v, Got: %v", fileIndex, rankIndex, BlackPawn, piece)
	}
}

func TestIsFriendlyCapture(t *testing.T) {
	var grid = [8][8]Piece{
		{WhiteRook, WhitePawn, nil, nil, nil, nil, BlackPawn, BlackRook},
		{nil, WhitePawn, nil, nil, nil, nil, BlackPawn, nil},
		{WhiteBishop, WhitePawn, nil, nil, nil, nil, BlackPawn, BlackBishop},
		{nil, WhitePawn, nil, nil, nil, nil, BlackPawn, nil},
		{nil, WhitePawn, nil, nil, nil, nil, BlackPawn, nil},
		{WhiteBishop, WhitePawn, nil, nil, nil, nil, BlackPawn, BlackBishop},
		{nil, WhitePawn, nil, nil, nil, nil, BlackPawn, nil},
		{WhiteRook, WhitePawn, nil, nil, nil, nil, BlackPawn, BlackRook},
	}
	state := NewGameState(White, grid, 0, PiecePosition{-1, -1})
	fileIndex := 0
	rankIndex := 1
	friendlyCapture := state.isFriendlyCapture(fileIndex, rankIndex)
	if friendlyCapture != true {
		t.Errorf("Friendly capture should be true for %v, %v for %v but got false", fileIndex, rankIndex, state.turn)
	}

	state = NewGameState(Black, grid, 0, PiecePosition{-1, -1})
	fileIndex = 0
	rankIndex = 6
	friendlyCapture = state.isFriendlyCapture(fileIndex, rankIndex)
	if friendlyCapture != true {
		t.Errorf("Friendly capture should be true for %v, %v for %v but got false", fileIndex, rankIndex, state.turn)
	}
}

func TestGetKingAttackers(t *testing.T) {
	type attackerTest struct {
		state     State
		attackers map[[2]int]Piece
	}

	var testData = []attackerTest{
		{
			attackers: map[[2]int]Piece{
				[2]int{0, 3}: BlackQueen,
				[2]int{2, 1}: BlackKnight,
				[2]int{5, 2}: BlackKnight,
				[2]int{7, 3}: BlackRook,
				[2]int{6, 6}: BlackBishop,
				[2]int{3, 4}: BlackKing,
				[2]int{2, 4}: BlackPawn,
			},
			state: NewGameState(White, [8][8]Piece{
				{nil, nil, nil, BlackQueen, nil, nil, nil, nil},
				{nil, nil, nil, nil, nil, nil, nil, nil},
				{nil, BlackKnight, nil, nil, BlackPawn, nil, nil, nil},
				{nil, nil, nil, WhiteKing, BlackKing, nil, nil, nil},
				{nil, nil, nil, nil, nil, nil, nil, nil},
				{nil, nil, BlackKnight, nil, nil, nil, nil, nil},
				{nil, nil, nil, nil, nil, nil, BlackBishop, nil},
				{nil, nil, nil, BlackRook, nil, nil, nil, nil},
			}, 0, PiecePosition{-1, -1}),
		},
		{
			attackers: map[[2]int]Piece{
				[2]int{0, 3}: WhiteQueen,
				[2]int{2, 1}: WhiteKnight,
				[2]int{5, 2}: WhiteKnight,
				[2]int{7, 3}: WhiteRook,
				[2]int{6, 6}: WhiteBishop,
				[2]int{3, 4}: WhiteKing,
				[2]int{2, 2}: WhitePawn,
			},
			state: NewGameState(Black, [8][8]Piece{
				{nil, nil, nil, WhiteQueen, nil, nil, nil, nil},
				{nil, nil, nil, nil, nil, nil, nil, nil},
				{nil, WhiteKnight, WhitePawn, nil, nil, nil, nil, nil},
				{nil, nil, nil, BlackKing, WhiteKing, nil, nil, nil},
				{nil, nil, nil, nil, nil, nil, nil, nil},
				{nil, nil, WhiteKnight, nil, nil, nil, nil, nil},
				{nil, nil, nil, nil, nil, nil, WhiteBishop, nil},
				{nil, nil, nil, WhiteRook, nil, nil, nil, nil},
			}, 0, PiecePosition{-1, -1}),
		},
	}

	for _, test := range testData {
		attackers := test.state.GetKingAttackers()
		if len(attackers) != len(test.attackers) {
			t.Errorf("Invalid number of attackers of %v king on %+v. Expected %v; Got %v", test.state.turn, test.state.kingsPosition, len(test.attackers), len(attackers))
		}
		for position, attacker := range test.attackers {
			if piece, ok := attackers[position]; ok {
				if attacker != piece {
					t.Errorf("Invalid attacker on position %v. Expected %v; Got %v", position, attacker, piece)
				}
			} else {
				t.Errorf("Attacker missing for position %v. Expected %v", position, attacker)
			}
		}
	}
}

func TestIsKingInCheck(t *testing.T) {
	grid := [8][8]Piece{
		{nil, nil, nil, nil, nil, nil, nil, nil},
		{nil, nil, nil, nil, nil, nil, nil, nil},
		{nil, nil, nil, nil, nil, nil, nil, nil},
		{nil, nil, nil, nil, nil, nil, nil, nil},
		{WhiteKing, nil, nil, nil, nil, nil, nil, BlackKing},
		{nil, nil, BlackKnight, nil, nil, nil, nil, nil},
		{nil, nil, nil, nil, nil, WhiteBishop, nil, nil},
		{nil, nil, nil, nil, nil, nil, nil, nil},
	}
	var states = []State{NewGameState(White, grid, 0, PiecePosition{-1, -1}), NewGameState(Black, grid, 0, PiecePosition{-1, -1})}

	for _, state := range states {
		attackers := state.GetKingAttackers()
		isInCheck := state.isKingInCheck(attackers)
		if !isInCheck {
			t.Errorf("%v king not in check. Expected to be in check", state.turn)
		}
	}
}

func TestGetKingsPosition(t *testing.T) {
	state := NewGameState(Black, [8][8]Piece{
		{nil, nil, nil, nil, nil, nil, nil, nil},
		{nil, nil, nil, nil, nil, nil, nil, nil},
		{nil, nil, nil, nil, nil, nil, nil, nil},
		{nil, nil, nil, nil, nil, nil, nil, nil},
		{WhiteKing, nil, nil, nil, nil, nil, nil, BlackKing},
		{nil, nil, nil, nil, nil, nil, nil, nil},
		{nil, nil, nil, nil, nil, nil, nil, nil},
		{nil, nil, nil, nil, nil, nil, nil, nil},
	}, 0, PiecePosition{-1, -1})
	kingsPosition := state.GetKingsPosition()
	if kingsPosition.WhiteKing.FileIndex != 4 &&
		kingsPosition.WhiteKing.RankIndex != 0 {
		t.Errorf("Invalid King position. Expected %v %v for %v king, Got %v, %v ", 4, 0, state.turn, kingsPosition.WhiteKing.FileIndex, kingsPosition.WhiteKing.RankIndex)
	}

	if kingsPosition.BlackKing.FileIndex != 4 &&
		kingsPosition.BlackKing.RankIndex != 0 {
		t.Errorf("Invalid King position. Expected %v %v for %v king, Got %v, %v ", 4, 7, state.turn, kingsPosition.BlackKing.FileIndex, kingsPosition.BlackKing.RankIndex)
	}
}

func TestIsKingInCheckAfter(t *testing.T) {
	type checkTest struct {
		state           State
		sourceFileIndex int
		targetFileIndex int
		sourceRankIndex int
		targetRankIndex int
		shouldBeInCheck bool
	}

	grid := [8][8]Piece{
		{nil, BlackQueen, WhitePawn, nil, nil, BlackPawn, WhiteRook, WhiteQueen},
		{nil, nil, WhitePawn, nil, BlackBishop, BlackPawn, nil, nil},
		{nil, nil, WhitePawn, nil, nil, BlackPawn, nil, nil},
		{nil, WhiteKnight, WhitePawn, nil, nil, BlackPawn, BlackKnight, nil},
		{nil, WhiteKing, WhitePawn, nil, WhiteKnight, BlackPawn, BlackKing, nil},
		{nil, nil, WhitePawn, nil, nil, BlackPawn, nil, nil},
		{nil, nil, WhitePawn, BlackKnight, WhiteBishop, BlackPawn, nil, nil},
		{BlackRook, nil, WhitePawn, nil, nil, BlackPawn, nil, nil},
	}
	state := NewGameState(White, grid, 0, PiecePosition{-1, -1})

	tests := []checkTest{
		{
			state:           state,
			sourceFileIndex: 3,
			sourceRankIndex: 1,
			targetFileIndex: 2,
			targetRankIndex: 3,
			shouldBeInCheck: true,
		},
		{
			state:           state,
			sourceFileIndex: 4,
			sourceRankIndex: 1,
			targetFileIndex: 4,
			targetRankIndex: 0,
			shouldBeInCheck: true,
		},
		{
			state:           state,
			sourceFileIndex: 3,
			sourceRankIndex: 1,
			targetFileIndex: 2,
			targetRankIndex: 3,
			shouldBeInCheck: true,
		},
		{
			state:           state,
			sourceFileIndex: 3,
			sourceRankIndex: 2,
			targetFileIndex: 3,
			targetRankIndex: 3,
			shouldBeInCheck: true,
		},
		{
			state:           state,
			sourceFileIndex: 0,
			sourceRankIndex: 2,
			targetFileIndex: 0,
			targetRankIndex: 3,
			shouldBeInCheck: false,
		},
		{
			state:           state,
			sourceFileIndex: 4,
			sourceRankIndex: 1,
			targetFileIndex: 5,
			targetRankIndex: 1,
			shouldBeInCheck: true,
		},
		{
			state:           NewGameState(Black, grid, 0, PiecePosition{-1, -1}),
			sourceFileIndex: 4,
			sourceRankIndex: 6,
			targetFileIndex: 4,
			targetRankIndex: 7,
			shouldBeInCheck: true,
		},
		{
			state:           NewGameState(Black, grid, 0, PiecePosition{-1, -1}),
			sourceFileIndex: 4,
			sourceRankIndex: 6,
			targetFileIndex: 5,
			targetRankIndex: 6,
			shouldBeInCheck: true,
		},
		{
			state:           NewGameState(Black, grid, 0, PiecePosition{-1, -1}),
			sourceFileIndex: 3,
			sourceRankIndex: 6,
			targetFileIndex: 5,
			targetRankIndex: 7,
			shouldBeInCheck: true,
		},
		{
			state:           NewGameState(Black, grid, 0, PiecePosition{-1, -1}),
			sourceFileIndex: 5,
			sourceRankIndex: 5,
			targetFileIndex: 5,
			targetRankIndex: 4,
			shouldBeInCheck: true,
		},
		{
			state:           NewGameState(Black, grid, 0, PiecePosition{-1, -1}),
			sourceFileIndex: 0,
			sourceRankIndex: 6,
			targetFileIndex: 0,
			targetRankIndex: 5,
			shouldBeInCheck: false,
		},
	}

	for _, test := range tests {
		isInCheck := test.state.isKingInCheckAfter(test.sourceFileIndex, test.sourceRankIndex, test.targetFileIndex, test.targetRankIndex) //WhiteKnight
		if isInCheck != test.shouldBeInCheck {
			t.Errorf("Check status mismatch after move. Expected %v; Got %v for source [%v,%v] & target [%v,%v]", test.shouldBeInCheck, isInCheck, test.sourceFileIndex, test.sourceRankIndex, test.targetFileIndex, test.targetRankIndex)
		}
	}
}

func TestStateStringRepresentation(t *testing.T) {
	state := NewGameState(Black, [8][8]Piece{
		{WhiteRook, WhitePawn, nil, nil, nil, nil, BlackPawn, BlackRook},
		{WhiteKnight, WhitePawn, nil, nil, nil, nil, BlackPawn, BlackKnight},
		{WhiteBishop, WhitePawn, nil, nil, nil, nil, BlackPawn, BlackBishop},
		{WhiteQueen, WhitePawn, nil, nil, nil, nil, BlackPawn, BlackQueen},
		{WhiteKing, WhitePawn, nil, nil, nil, nil, BlackPawn, BlackKing},
		{WhiteBishop, WhitePawn, nil, nil, nil, nil, BlackPawn, BlackBishop},
		{WhiteKnight, WhitePawn, nil, nil, nil, nil, BlackPawn, BlackKnight},
		{WhiteRook, WhitePawn, nil, nil, nil, nil, BlackPawn, BlackRook},
	}, 0, PiecePosition{-1, -1})
	stringRep := state.String()
	shouldBe := `
+---+---+---+---+---+---+---+---+
8 | r | n | b | q | k | b | n | r |
  +---+---+---+---+---+---+---+---+
7 | p | p | p | p | p | p | p | p |
  +---+---+---+---+---+---+---+---+
6 | . | . | . | . | . | . | . | . |
  +---+---+---+---+---+---+---+---+
5 | . | . | . | . | . | . | . | . |
  +---+---+---+---+---+---+---+---+
4 | . | . | . | . | . | . | . | . |
  +---+---+---+---+---+---+---+---+
3 | . | . | . | . | . | . | . | . |
  +---+---+---+---+---+---+---+---+
2 | P | P | P | P | P | P | P | P |
  +---+---+---+---+---+---+---+---+
1 | R | N | B | Q | K | B | N | R |
  +---+---+---+---+---+---+---+---+
    a   b   c   d   e   f   g   h`
	if strings.TrimSpace(stringRep) != strings.TrimSpace(shouldBe) {
		t.Errorf("Expected string representation to be '%v', got '%v'", stringRep, state)
	}
}

func TestPlayerTurnStringRepresentation(t *testing.T) {
	if White.String() != "White" {
		t.Errorf("Expected string representation of White to be 'White', got '%v'", White.String())
	}

	if Black.String() != "Black" {
		t.Errorf("Expected string representation of Black to be 'Black', got '%v'", Black.String())
	}
}
