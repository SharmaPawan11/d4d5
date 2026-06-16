package core

import "testing"

func TestKnightMoves(t *testing.T) {
	set := make(map[Move]struct{})

	grid := [8][8]Piece{
		{nil, WhitePawn, nil, nil, nil, nil, BlackPawn, nil},
		{WhiteKnight, WhitePawn, nil, nil, nil, nil, BlackPawn, BlackKnight},
		{nil, WhitePawn, nil, nil, nil, nil, BlackPawn, nil},
		{nil, WhitePawn, nil, nil, nil, nil, BlackPawn, nil},
		{WhiteKing, WhitePawn, nil, nil, nil, nil, BlackPawn, BlackKing},
		{nil, WhitePawn, nil, BlackKnight, WhiteKnight, nil, BlackPawn, nil},
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
			algebraicSquare: "b1",
			correctMoves: []Move{
				{
					Target: PiecePosition{
						FileIndex: 0,
						RankIndex: 2,
					},
				},
				{
					Target: PiecePosition{
						FileIndex: 2,
						RankIndex: 2,
					},
				},
			},
			piecePosition: PiecePosition{FileIndex: 1, RankIndex: 0},
			state:         state,
		},
		{
			algebraicSquare: "f5",
			correctMoves: []Move{
				{
					Target: PiecePosition{
						FileIndex: 7,
						RankIndex: 3,
					},
				},
				{
					Target: PiecePosition{
						FileIndex: 7,
						RankIndex: 5,
					},
				},
				{
					Target: PiecePosition{
						FileIndex: 3,
						RankIndex: 3,
					},
				},
				{
					Target: PiecePosition{
						FileIndex: 3,
						RankIndex: 5,
					},
				},
				{
					Target: PiecePosition{
						FileIndex: 6,
						RankIndex: 2,
					},
				},
				{
					Target: PiecePosition{
						FileIndex: 4,
						RankIndex: 2,
					},
				},
				{
					Target: PiecePosition{
						FileIndex: 4,
						RankIndex: 6,
					},
				},
				{
					Target: PiecePosition{
						FileIndex: 6,
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
						FileIndex: 7,
						RankIndex: 4,
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
						FileIndex: 3,
						RankIndex: 4,
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
						RankIndex: 5,
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
						FileIndex: 4,
						RankIndex: 1,
					},
				},
				{
					Target: PiecePosition{
						FileIndex: 6,
						RankIndex: 1,
					},
				},
			},
			piecePosition: PiecePosition{FileIndex: 5, RankIndex: 3},
			state:         NewGameState(Black, grid, 0, PiecePosition{-1, -1}),
		},
		{
			algebraicSquare: "b8",
			correctMoves: []Move{
				{
					Target: PiecePosition{
						FileIndex: 0,
						RankIndex: 5,
					},
				},
				{
					Target: PiecePosition{
						FileIndex: 2,
						RankIndex: 5,
					},
				},
			},
			piecePosition: PiecePosition{FileIndex: 1, RankIndex: 7},
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
			t.Errorf("Invalid number of moves for %v knight for %v. Expected %v, Got %v", move.algebraicSquare, move.state.turn, len(move.correctMoves), len(possibleMoves))
		}
		for _, correctMove := range move.correctMoves {
			if _, found := set[correctMove]; !found {
				t.Errorf("Valid move %v missing from the set of moves returned for %v knight for %v", correctMove.Target, move.algebraicSquare, move.state.turn)
			}
		}
	}
}

func TestKnightStringRep(t *testing.T) {
	if WhiteKnight.String() != "N" {
		t.Errorf("Incorrect string representation of WhiteKnight, Expected N, got %s", WhiteKnight.String())
	}

	if BlackKnight.String() != "n" {
		t.Errorf("Incorrect string representation of BlackKnight, Expected n, got %s", BlackKnight.String())
	}

	var neitherColorKnight knight = 0
	if neitherColorKnight.String() != "" {
		t.Errorf("Incorrect string representation of non-colored knight, Expected blank string, got %s", neitherColorKnight.String())
	}
}
