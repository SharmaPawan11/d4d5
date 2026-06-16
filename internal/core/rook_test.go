package core

import (
	"testing"
)

func TestRookMoves(t *testing.T) {
	set := make(map[Move]struct{})

	grid := [8][8]Piece{
		{WhiteRook, WhitePawn, nil, nil, nil, nil, BlackPawn, nil},
		{nil, WhitePawn, nil, nil, BlackRook, nil, BlackPawn, nil},
		{nil, WhitePawn, nil, nil, nil, nil, BlackPawn, nil},
		{nil, WhitePawn, nil, nil, nil, nil, BlackPawn, nil},
		{nil, WhitePawn, nil, WhiteKing, nil, BlackKing, BlackPawn, nil},
		{nil, WhitePawn, nil, nil, WhiteRook, nil, BlackPawn, BlackRook},
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
			algebraicSquare: "a1",
			correctMoves: []Move{
				{
					Target: PiecePosition{
						FileIndex: 1,
						RankIndex: 0,
					},
				},
				{
					Target: PiecePosition{
						FileIndex: 2,
						RankIndex: 0,
					},
				},
				{
					Target: PiecePosition{
						FileIndex: 3,
						RankIndex: 0,
					},
				},
				{
					Target: PiecePosition{
						FileIndex: 4,
						RankIndex: 0,
					},
				},
				{
					Target: PiecePosition{
						FileIndex: 5,
						RankIndex: 0,
					},
				},
				{
					Target: PiecePosition{
						FileIndex: 6,
						RankIndex: 0,
					},
				},
				{
					Target: PiecePosition{
						FileIndex: 7,
						RankIndex: 0,
					},
				},
			},
			piecePosition: PiecePosition{FileIndex: 0, RankIndex: 0},
			state:         state,
		},
		{
			algebraicSquare: "f5",
			correctMoves: []Move{
				{
					Target: PiecePosition{
						FileIndex: 5,
						RankIndex: 3,
					},
				},
				{
					Target: PiecePosition{
						FileIndex: 5,
						RankIndex: 2,
					},
				},
				{
					Target: PiecePosition{
						FileIndex: 5,
						RankIndex: 5,
					},
				},
				{
					Target: PiecePosition{
						FileIndex: 5,
						RankIndex: 6,
					},
				},
				{
					Target: PiecePosition{
						FileIndex: 6,
						RankIndex: 4,
					},
				},
				{
					Target: PiecePosition{
						FileIndex: 7,
						RankIndex: 4,
					},
				},
				{
					Target: PiecePosition{
						FileIndex: 4,
						RankIndex: 4,
					},
				},
				{
					Target: PiecePosition{
						FileIndex: 3,
						RankIndex: 4,
					},
				},
				{
					Target: PiecePosition{
						FileIndex: 2,
						RankIndex: 4,
					},
				},
				{
					Target: PiecePosition{
						FileIndex: 1,
						RankIndex: 4,
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
						RankIndex: 4,
					},
				},
				{
					Target: PiecePosition{
						FileIndex: 2,
						RankIndex: 4,
					},
				},
				{
					Target: PiecePosition{
						FileIndex: 3,
						RankIndex: 4,
					},
				},
				{
					Target: PiecePosition{
						FileIndex: 4,
						RankIndex: 4,
					},
				},
				{
					Target: PiecePosition{
						FileIndex: 5,
						RankIndex: 4,
					},
				},
				{
					Target: PiecePosition{
						FileIndex: 1,
						RankIndex: 5,
					},
				},
				{
					Target: PiecePosition{
						FileIndex: 1,
						RankIndex: 3,
					},
				},
				{
					Target: PiecePosition{
						FileIndex: 1,
						RankIndex: 2,
					},
				},
				{
					Target: PiecePosition{
						FileIndex: 1,
						RankIndex: 1,
					},
				},
			},
			piecePosition: PiecePosition{FileIndex: 1, RankIndex: 4},
			state:         NewGameState(Black, grid, 0, PiecePosition{-1, -1}),
		},
		{
			algebraicSquare: "f8",
			correctMoves: []Move{
				{
					Target: PiecePosition{
						FileIndex: 4,
						RankIndex: 7,
					},
				},
				{
					Target: PiecePosition{
						FileIndex: 3,
						RankIndex: 7,
					},
				},
				{
					Target: PiecePosition{
						FileIndex: 2,
						RankIndex: 7,
					},
				},
				{
					Target: PiecePosition{
						FileIndex: 1,
						RankIndex: 7,
					},
				},
				{
					Target: PiecePosition{
						FileIndex: 0,
						RankIndex: 7,
					},
				},
				{
					Target: PiecePosition{
						FileIndex: 6,
						RankIndex: 7,
					},
				},
				{
					Target: PiecePosition{
						FileIndex: 7,
						RankIndex: 7,
					},
				},
			},
			piecePosition: PiecePosition{FileIndex: 5, RankIndex: 7},
			state:         NewGameState(Black, grid, 0, PiecePosition{-1, -1}),
		},
	}

	for _, move := range testMoves {
		chosenPiece := move.state.GetPieceAt(move.piecePosition.FileIndex, move.piecePosition.RankIndex)
		possibleMoves := chosenPiece.GetMoves(move.state, move.piecePosition)
		for _, item := range possibleMoves {
			set[item] = struct{}{}
		}

		if len(possibleMoves) != len(move.correctMoves) {
			t.Errorf("Invalid number of moves for %v rook for %v. Expected %v, Got %v", move.algebraicSquare, move.state.turn, len(move.correctMoves), len(possibleMoves))
		}

		for _, correctMove := range move.correctMoves {
			if _, found := set[correctMove]; !found {
				t.Errorf(`Valid move %v missing from the set of moves returned for %v rook for %v`, correctMove.Target, move.algebraicSquare, move.state.turn)

			}
		}
	}
}

func TestRookStringRep(t *testing.T) {
	if WhiteRook.String() != "R" {
		t.Errorf("Incorrect string representation of WhiteRook, Expected R, got %s", WhiteRook.String())
	}

	if BlackRook.String() != "r" {
		t.Errorf("Incorrect string representation of BlackRook, Expected r, got %s", BlackRook.String())
	}

	var neitherColorRook rook = 0
	if neitherColorRook.String() != "" {
		t.Errorf("Incorrect string representation of non-colored rook, Expected blank string, got %s", neitherColorRook.String())
	}
}
