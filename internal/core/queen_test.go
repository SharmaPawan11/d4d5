package core

import "testing"

func TestQueenMoves(t *testing.T) {
	set := make(map[Move]struct{})

	grid := [8][8]Piece{
		{nil, WhitePawn, nil, nil, nil, nil, BlackPawn, nil},
		{nil, WhitePawn, nil, nil, nil, nil, BlackPawn, nil},
		{nil, WhitePawn, nil, nil, nil, nil, BlackPawn, nil},
		{nil, WhitePawn, nil, WhiteQueen, nil, nil, BlackPawn, nil},
		{WhiteKing, WhitePawn, nil, nil, BlackQueen, nil, BlackPawn, BlackKing},
		{nil, WhitePawn, nil, nil, nil, nil, BlackPawn, nil},
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
			algebraicSquare: "e4",
			correctMoves: []Move{
				{
					Target: PiecePosition{
						FileIndex: 3,
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
						RankIndex: 5,
					},
				},
				{
					Target: PiecePosition{
						FileIndex: 3,
						RankIndex: 6,
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
						FileIndex: 1,
						RankIndex: 3,
					},
				},
				{
					Target: PiecePosition{
						FileIndex: 0,
						RankIndex: 3,
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
						FileIndex: 5,
						RankIndex: 3,
					},
				},
				{
					Target: PiecePosition{
						FileIndex: 6,
						RankIndex: 3,
					},
				},
				{
					Target: PiecePosition{
						FileIndex: 7,
						RankIndex: 3,
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
						FileIndex: 2,
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
						FileIndex: 2,
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
						FileIndex: 0,
						RankIndex: 6,
					},
				},
			},
			piecePosition: PiecePosition{FileIndex: 3, RankIndex: 3},
			state:         state,
		},
		{
			algebraicSquare: "d5",
			correctMoves: []Move{
				{
					Target: PiecePosition{
						FileIndex: 3,
						RankIndex: 3,
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
						FileIndex: 3,
						RankIndex: 5,
					},
				},
				{
					Target: PiecePosition{
						FileIndex: 5,
						RankIndex: 3,
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
						FileIndex: 7,
						RankIndex: 1,
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
						FileIndex: 4,
						RankIndex: 3,
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
						RankIndex: 1,
					},
				},
				{
					Target: PiecePosition{
						FileIndex: 0,
						RankIndex: 4,
					},
				},
				{
					Target: PiecePosition{
						FileIndex: 1,
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
						FileIndex: 5,
						RankIndex: 4,
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
			},
			piecePosition: PiecePosition{FileIndex: 4, RankIndex: 4},
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
			t.Errorf("Invalid number of moves for %v queen for %v. Expected %v, Got %v", move.algebraicSquare, move.state.turn, len(move.correctMoves), len(possibleMoves))

		}
		for _, correctMove := range move.correctMoves {
			if _, found := set[correctMove]; !found {
				t.Errorf("Valid move %v missing from the set of moves returned for %v queen for %v", correctMove.Target, move.algebraicSquare, move.state.turn)
			}
		}
	}
}

func TestQueenStringRep(t *testing.T) {
	if WhiteQueen.String() != "Q" {
		t.Errorf("Incorrect string representation of WhiteQueen, Expected Q, got %s", WhiteQueen.String())
	}

	if BlackQueen.String() != "q" {
		t.Errorf("Incorrect string representation of BlackQueen, Expected q, got %s", BlackQueen.String())
	}

	var neitherColorQueen queen = 0
	if neitherColorQueen.String() != "" {
		t.Errorf("Incorrect string representation of non-colored queen, Expected blank string, got %s", neitherColorQueen.String())
	}
}
