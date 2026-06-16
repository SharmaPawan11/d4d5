package core

import (
	"testing"
)

func TestBishopMoves(t *testing.T) {
	set := make(map[Move]struct{})

	grid := [8][8]Piece{
		{nil, WhitePawn, nil, nil, nil, nil, BlackPawn, nil},
		{nil, WhitePawn, nil, nil, BlackBishop, nil, BlackPawn, nil},
		{WhiteBishop, WhitePawn, nil, nil, nil, nil, BlackPawn, nil},
		{nil, WhitePawn, nil, nil, nil, nil, BlackPawn, nil},
		{WhiteKing, WhitePawn, nil, nil, nil, nil, BlackPawn, BlackKing},
		{nil, WhitePawn, nil, nil, WhiteBishop, nil, BlackPawn, BlackBishop},
		{nil, WhitePawn, nil, nil, nil, nil, BlackPawn, nil},
		{nil, WhitePawn, nil, nil, nil, nil, BlackPawn, nil},
	}

	state := NewGameState(
		White,
		grid,
		0,
		PiecePosition{-1, -1},
	)

	testMoves := []TestMove{
		{
			algebraicSquare: "c1",
			correctMoves:    []Move{},
			piecePosition:   PiecePosition{FileIndex: 2, RankIndex: 0},
			state:           state,
		},
		{
			algebraicSquare: "f5",
			correctMoves: []Move{
				{
					Target: PiecePosition{
						FileIndex: 6,
						RankIndex: 3,
					},
				},
				{
					Target: PiecePosition{
						FileIndex: 7,
						RankIndex: 2,
					},
				},
				{
					Target: PiecePosition{
						FileIndex: 4,
						RankIndex: 3,
					},
				},
				{
					Target: PiecePosition{
						FileIndex: 3,
						RankIndex: 2,
					},
				},
				{
					Target: PiecePosition{
						FileIndex: 6,
						RankIndex: 5,
					},
				},
				{
					Target: PiecePosition{
						FileIndex: 7,
						RankIndex: 6,
					},
				},
				{
					Target: PiecePosition{
						FileIndex: 4,
						RankIndex: 5,
					},
				},
				{
					Target: PiecePosition{
						FileIndex: 3,
						RankIndex: 6,
					},
				},
			},
			piecePosition: PiecePosition{FileIndex: 5, RankIndex: 4},
			state:         state,
		},
		{
			algebraicSquare: "b5",
			correctMoves: []Move{
				{
					Target: PiecePosition{
						FileIndex: 0,
						RankIndex: 3,
					},
				},
				{
					Target: PiecePosition{
						FileIndex: 0,
						RankIndex: 5,
					},
				},
				{
					Target: PiecePosition{
						FileIndex: 2,
						RankIndex: 3,
					},
				},
				{
					Target: PiecePosition{
						FileIndex: 3,
						RankIndex: 2,
					},
				},
				{
					Target: PiecePosition{
						FileIndex: 4,
						RankIndex: 1,
					},
				},
				{
					Target: PiecePosition{
						FileIndex: 2,
						RankIndex: 5,
					},
				},
			},
			piecePosition: PiecePosition{FileIndex: 1, RankIndex: 4},
			state:         NewGameState(Black, grid, 0, PiecePosition{-1, -1}),
		},
		{
			algebraicSquare: "f8",
			correctMoves:    []Move{},
			piecePosition:   PiecePosition{FileIndex: 5, RankIndex: 7},
			state:           NewGameState(Black, grid, 0, PiecePosition{-1, -1}),
		},
	}

	for _, move := range testMoves {
		chosenPiece := move.state.GetPieceAt(move.piecePosition.FileIndex, move.piecePosition.RankIndex)
		possibleMoves := chosenPiece.GetMoves(move.state, move.piecePosition)
		for _, item := range possibleMoves {
			set[item] = struct{}{}
		}
		if len(possibleMoves) != len(move.correctMoves) {
			t.Errorf("Invalid number of moves for %v bishop for %v. Expected %v, Got %v", move.algebraicSquare, move.state.turn, len(move.correctMoves), len(possibleMoves))

		}
		for _, correctMove := range move.correctMoves {
			if _, found := set[correctMove]; !found {
				t.Errorf("Valid move %v missing from the set of moves returned for %v bishop for %v", correctMove.Target, move.algebraicSquare, move.state.turn)
			}
		}
	}
}

func TestBishopStringRep(t *testing.T) {
	if WhiteBishop.String() != "B" {
		t.Errorf("Incorrect string representation of WhiteBishop, Expected B, got %s", WhiteBishop.String())
	}

	if BlackBishop.String() != "b" {
		t.Errorf("Incorrect string representation of BlackBishop, Expected b, got %s", BlackBishop.String())
	}

	var neitherColorBishop bishop = 0
	if neitherColorBishop.String() != "" {
		t.Errorf("Incorrect string representation of non-colored bishop, Expected blank string, got %s", neitherColorBishop.String())
	}
}
