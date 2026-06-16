package core

import (
	"log"
	"testing"
)

type TestMove struct {
	algebraicSquare string
	correctMoves    []Move
	piecePosition   PiecePosition
	state           State
}

func TestPawnMoves(t *testing.T) {
	set := make(map[Move]struct{})

	grid := [8][8]Piece{
		{nil, WhitePawn, nil, nil, nil, WhitePawn, BlackPawn, nil},
		{nil, WhitePawn, BlackPawn, nil, nil, nil, BlackPawn, nil},
		{nil, WhitePawn, nil, nil, nil, WhitePawn, BlackPawn, nil},
		{nil, WhitePawn, nil, nil, nil, nil, BlackPawn, nil},
		{WhiteKing, WhitePawn, nil, nil, nil, nil, BlackPawn, BlackKing},
		{nil, WhitePawn, BlackPawn, nil, nil, nil, BlackPawn, nil},
		{nil, WhitePawn, nil, nil, nil, WhitePawn, BlackPawn, nil},
		{nil, WhitePawn, BlackPawn, nil, nil, nil, BlackPawn, nil},
	}

	state := NewGameState(
		White,
		grid,
		0,
		PiecePosition{FileIndex: -1, RankIndex: -1},
	)

	testMoves := []TestMove{
		{
			algebraicSquare: "a2",
			correctMoves: []Move{
				{
					Target: PiecePosition{
						FileIndex: 0,
						RankIndex: 2,
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
						FileIndex: 0,
						RankIndex: 3,
					},
				},
			},
			piecePosition: PiecePosition{FileIndex: 0, RankIndex: 1},
			state:         state,
		},
		{
			algebraicSquare: "b2",
			correctMoves:    []Move{},
			piecePosition:   PiecePosition{FileIndex: 1, RankIndex: 1},
			state:           state,
		},
		{
			algebraicSquare: "f2",
			correctMoves:    []Move{},
			piecePosition:   PiecePosition{FileIndex: 5, RankIndex: 1},
			state:           state,
		},
		{
			algebraicSquare: "g2",
			correctMoves: []Move{
				{
					Target: PiecePosition{
						FileIndex: 7,
						RankIndex: 2,
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
						FileIndex: 6,
						RankIndex: 2,
					},
				},
			},
			piecePosition: PiecePosition{FileIndex: 6, RankIndex: 1},
			state:         state,
		},
		{
			algebraicSquare: "h2",
			correctMoves:    []Move{},
			piecePosition:   PiecePosition{FileIndex: 7, RankIndex: 1},
			state:           state,
		},
		{
			algebraicSquare: "a7",
			correctMoves:    []Move{},
			piecePosition:   PiecePosition{FileIndex: 0, RankIndex: 6},
			state:           NewGameState(Black, grid, 0, PiecePosition{-1, -1}),
		},
		{
			algebraicSquare: "b7",
			correctMoves: []Move{
				{
					Target: PiecePosition{
						FileIndex: 1,
						RankIndex: 5,
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
			piecePosition: PiecePosition{FileIndex: 1, RankIndex: 6},
			state:         NewGameState(Black, grid, 0, PiecePosition{-1, -1}),
		},
		{
			algebraicSquare: "f7",
			correctMoves: []Move{
				{
					Target: PiecePosition{
						FileIndex: 5,
						RankIndex: 5,
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
						RankIndex: 5,
					},
				},
			},
			piecePosition: PiecePosition{FileIndex: 5, RankIndex: 6},
			state:         NewGameState(Black, grid, 0, PiecePosition{-1, -1}),
		},
		{
			algebraicSquare: "g7",
			correctMoves:    []Move{},
			piecePosition:   PiecePosition{FileIndex: 6, RankIndex: 6},
			state:           NewGameState(Black, grid, 0, PiecePosition{-1, -1}),
		},
		{
			algebraicSquare: "h7",
			correctMoves: []Move{
				{
					Target: PiecePosition{
						FileIndex: 6,
						RankIndex: 5,
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
						FileIndex: 7,
						RankIndex: 5,
					},
				},
			},
			piecePosition: PiecePosition{FileIndex: 7, RankIndex: 6},
			state:         NewGameState(Black, grid, 0, PiecePosition{-1, -1}),
		},
		{
			algebraicSquare: "e4",
			correctMoves: []Move{
				{
					Target: PiecePosition{
						FileIndex: 4,
						RankIndex: 2,
					},
				},
				{
					Target: PiecePosition{
						FileIndex: 3,
						RankIndex: 2,
					},
				},
			},
			piecePosition: PiecePosition{FileIndex: 4, RankIndex: 3},
			state: NewGameState(Black, [8][8]Piece{
				{nil, WhitePawn, nil, nil, nil, nil, BlackPawn, nil},
				{nil, WhitePawn, nil, nil, nil, nil, BlackPawn, nil},
				{nil, WhitePawn, nil, nil, nil, nil, BlackPawn, nil},
				{nil, nil, nil, WhitePawn, nil, nil, BlackPawn, nil},
				{WhiteKing, WhitePawn, nil, BlackPawn, nil, nil, nil, BlackKing},
				{nil, WhitePawn, nil, nil, nil, nil, BlackPawn, nil},
				{nil, WhitePawn, nil, nil, nil, nil, BlackPawn, nil},
				{nil, WhitePawn, nil, nil, nil, nil, BlackPawn, nil},
			}, 0, PiecePosition{3, 2}),
		},
		{
			algebraicSquare: "c7",
			correctMoves: []Move{
				{
					Target: PiecePosition{
						FileIndex: 2,
						RankIndex: 7,
					},
					PromoteTo: WhiteQueen,
				},
				{
					Target: PiecePosition{
						FileIndex: 2,
						RankIndex: 7,
					},
					PromoteTo: WhiteRook,
				},
				{
					Target: PiecePosition{
						FileIndex: 2,
						RankIndex: 7,
					},
					PromoteTo: WhiteBishop,
				},
				{
					Target: PiecePosition{
						FileIndex: 2,
						RankIndex: 7,
					},
					PromoteTo: WhiteKnight,
				},
				{
					Target: PiecePosition{
						FileIndex: 3,
						RankIndex: 7,
					},
					PromoteTo: WhiteQueen,
				},
				{
					Target: PiecePosition{
						FileIndex: 3,
						RankIndex: 7,
					},
					PromoteTo: WhiteRook,
				},
				{
					Target: PiecePosition{
						FileIndex: 3,
						RankIndex: 7,
					},
					PromoteTo: WhiteBishop,
				},
				{
					Target: PiecePosition{
						FileIndex: 3,
						RankIndex: 7,
					},
					PromoteTo: WhiteKnight,
				},
				{
					Target: PiecePosition{
						FileIndex: 1,
						RankIndex: 7,
					},
					PromoteTo: WhiteQueen,
				},
				{
					Target: PiecePosition{
						FileIndex: 1,
						RankIndex: 7,
					},
					PromoteTo: WhiteRook,
				},
				{
					Target: PiecePosition{
						FileIndex: 1,
						RankIndex: 7,
					},
					PromoteTo: WhiteBishop,
				},
				{
					Target: PiecePosition{
						FileIndex: 1,
						RankIndex: 7,
					},
					PromoteTo: WhiteKnight,
				},
			},
			piecePosition: PiecePosition{FileIndex: 2, RankIndex: 6},
			state: NewGameState(White, [8][8]Piece{
				{nil, WhitePawn, nil, nil, nil, nil, BlackPawn, nil},
				{nil, WhitePawn, nil, nil, nil, nil, BlackPawn, BlackKnight},
				{nil, nil, nil, nil, nil, nil, WhitePawn, nil},
				{nil, WhitePawn, nil, nil, nil, nil, BlackPawn, BlackQueen},
				{WhiteKing, WhitePawn, nil, nil, nil, nil, BlackPawn, BlackKing},
				{nil, WhitePawn, nil, nil, nil, nil, BlackPawn, nil},
				{nil, WhitePawn, nil, nil, nil, nil, BlackPawn, nil},
				{nil, WhitePawn, nil, nil, nil, nil, BlackPawn, nil},
			}, 0, PiecePosition{-1, -1}),
		},
		{
			algebraicSquare: "f4 (Ghost Pawn En Passant Pin)",
			correctMoves: []Move{{
				Target: PiecePosition{
					FileIndex: 5,
					RankIndex: 2,
				},
			},
			},
			piecePosition: PiecePosition{FileIndex: 5, RankIndex: 3},
			state: NewGameState(Black, [8][8]Piece{
				{nil, nil, nil, nil, WhiteKing, nil, nil, nil},
				{nil, nil, nil, WhiteRook, WhitePawn, nil, nil, nil},
				{nil, nil, nil, nil, nil, nil, BlackPawn, nil},
				{nil, nil, nil, nil, nil, BlackPawn, nil, nil},
				{nil, nil, nil, WhitePawn, nil, nil, nil, nil},
				{nil, nil, nil, BlackPawn, nil, nil, nil, nil},
				{nil, WhitePawn, nil, nil, nil, nil, nil, nil},
				{nil, nil, nil, BlackKing, BlackRook, nil, nil, nil},
			}, 0, PiecePosition{FileIndex: 4, RankIndex: 2}),
		},
	}

	for _, move := range testMoves {
		chosenPiece := move.state.GetPieceAt(move.piecePosition.FileIndex, move.piecePosition.RankIndex)
		possibleMoves := chosenPiece.GetMoves(move.state, move.piecePosition)
		for _, item := range possibleMoves {
			set[item] = struct{}{}
		}
		if len(possibleMoves) != len(move.correctMoves) {
			t.Errorf("Invalid number of moves for %v pawn for %v. Expected %v, Got %v", move.algebraicSquare, move.state.turn, len(move.correctMoves), len(possibleMoves))
		}
		for _, correctMove := range move.correctMoves {
			if _, found := set[correctMove]; !found {
				t.Errorf("Valid move %v missing from the set of moves returned for %v pawn for %v", correctMove, move.algebraicSquare, move.state.turn)
			}
		}
	}
}

func TestGetEnPassantTargetSquare(t *testing.T) {

	type EnPassantTest struct {
		state                 State
		sourceRank            int
		targetSquare          PiecePosition
		enPassantTargetSquare PiecePosition
	}

	tests := []EnPassantTest{
		{
			state: NewGameState(
				White,
				[8][8]Piece{
					{nil, WhitePawn, nil, nil, nil, nil, BlackPawn, nil},
					{nil, WhitePawn, nil, nil, nil, nil, BlackPawn, nil},
					{nil, WhitePawn, nil, nil, nil, nil, BlackPawn, nil},
					{nil, WhitePawn, nil, nil, nil, nil, BlackPawn, nil},
					{WhiteKing, WhitePawn, nil, nil, nil, nil, BlackPawn, BlackKing},
					{nil, WhitePawn, nil, nil, nil, nil, BlackPawn, nil},
					{nil, WhitePawn, nil, nil, nil, nil, BlackPawn, nil},
					{nil, WhitePawn, nil, nil, nil, nil, BlackPawn, nil},
				},
				0,
				PiecePosition{-1, -1},
			),
			sourceRank:            1,
			targetSquare:          PiecePosition{FileIndex: 0, RankIndex: 3},
			enPassantTargetSquare: PiecePosition{-1, -1},
		},
		{
			state: NewGameState(
				White,
				[8][8]Piece{
					{nil, WhitePawn, nil, nil, nil, nil, BlackPawn, nil},
					{nil, WhitePawn, nil, nil, nil, nil, BlackPawn, nil},
					{nil, WhitePawn, nil, nil, nil, nil, BlackPawn, nil},
					{nil, WhitePawn, nil, nil, nil, nil, BlackPawn, nil},
					{WhiteKing, WhitePawn, nil, BlackPawn, nil, nil, nil, BlackKing},
					{nil, WhitePawn, nil, nil, nil, nil, BlackPawn, nil},
					{nil, WhitePawn, nil, nil, nil, nil, BlackPawn, nil},
					{nil, WhitePawn, nil, nil, nil, nil, BlackPawn, nil},
				},
				0,
				PiecePosition{-1, -1},
			),
			sourceRank:            1,
			targetSquare:          PiecePosition{FileIndex: 3, RankIndex: 3},
			enPassantTargetSquare: PiecePosition{FileIndex: 3, RankIndex: 2},
		},
		{
			state: NewGameState(
				White,
				[8][8]Piece{
					{nil, WhitePawn, nil, nil, nil, nil, BlackPawn, nil},
					{nil, WhitePawn, nil, nil, nil, nil, BlackPawn, nil},
					{nil, WhitePawn, nil, nil, nil, nil, BlackPawn, nil},
					{nil, WhitePawn, nil, nil, nil, nil, BlackPawn, nil},
					{WhiteKing, WhitePawn, nil, BlackPawn, nil, nil, nil, BlackKing},
					{nil, WhitePawn, nil, nil, nil, nil, BlackPawn, nil},
					{nil, WhitePawn, nil, nil, nil, nil, BlackPawn, nil},
					{nil, WhitePawn, nil, nil, nil, nil, BlackPawn, nil},
				},
				0,
				PiecePosition{-1, -1},
			),
			sourceRank:            1,
			targetSquare:          PiecePosition{FileIndex: 5, RankIndex: 3},
			enPassantTargetSquare: PiecePosition{FileIndex: 5, RankIndex: 2},
		},
		{
			state: NewGameState(
				White,
				[8][8]Piece{
					{nil, WhitePawn, nil, nil, nil, nil, BlackPawn, nil},
					{nil, WhitePawn, nil, nil, nil, nil, BlackPawn, nil},
					{nil, WhitePawn, nil, nil, nil, nil, BlackPawn, nil},
					{nil, WhitePawn, nil, nil, nil, nil, BlackPawn, nil},
					{WhiteKing, WhitePawn, nil, BlackPawn, nil, nil, nil, BlackKing},
					{nil, WhitePawn, nil, nil, nil, nil, BlackPawn, nil},
					{nil, WhitePawn, nil, nil, nil, nil, BlackPawn, nil},
					{nil, WhitePawn, nil, nil, nil, nil, BlackPawn, nil},
				},
				0,
				PiecePosition{-1, -1},
			),
			sourceRank:            1,
			targetSquare:          PiecePosition{FileIndex: 5, RankIndex: 2},
			enPassantTargetSquare: PiecePosition{FileIndex: -1, RankIndex: -1},
		},
		{
			state: NewGameState(
				White,
				[8][8]Piece{
					{nil, WhitePawn, nil, nil, nil, nil, BlackPawn, nil},
					{nil, WhitePawn, nil, nil, nil, nil, BlackPawn, nil},
					{nil, WhitePawn, nil, nil, nil, nil, BlackPawn, nil},
					{nil, WhitePawn, nil, nil, nil, nil, BlackPawn, nil},
					{WhiteKing, WhitePawn, nil, BlackPawn, nil, nil, nil, BlackKing},
					{nil, nil, WhitePawn, nil, nil, nil, BlackPawn, nil},
					{nil, WhitePawn, nil, nil, nil, nil, BlackPawn, nil},
					{nil, WhitePawn, nil, nil, nil, nil, BlackPawn, nil},
				},
				0,
				PiecePosition{-1, -1},
			),
			sourceRank:            2,
			targetSquare:          PiecePosition{FileIndex: 5, RankIndex: 3},
			enPassantTargetSquare: PiecePosition{FileIndex: -1, RankIndex: -1},
		},
		{
			state: NewGameState(
				Black,
				[8][8]Piece{
					{nil, WhitePawn, nil, nil, nil, nil, BlackPawn, nil},
					{nil, WhitePawn, nil, nil, nil, nil, BlackPawn, nil},
					{nil, WhitePawn, nil, nil, nil, nil, BlackPawn, nil},
					{nil, WhitePawn, nil, nil, nil, nil, BlackPawn, nil},
					{WhiteKing, nil, nil, nil, WhitePawn, nil, BlackPawn, BlackKing},
					{nil, WhitePawn, nil, nil, nil, nil, BlackPawn, nil},
					{nil, WhitePawn, nil, nil, nil, nil, BlackPawn, nil},
					{nil, WhitePawn, nil, nil, nil, nil, BlackPawn, nil},
				},
				0,
				PiecePosition{-1, -1},
			),
			sourceRank:            6,
			targetSquare:          PiecePosition{FileIndex: 3, RankIndex: 4},
			enPassantTargetSquare: PiecePosition{FileIndex: 3, RankIndex: 5},
		},
		{
			state: NewGameState(
				Black,
				[8][8]Piece{
					{nil, WhitePawn, nil, nil, nil, nil, BlackPawn, nil},
					{nil, WhitePawn, nil, nil, nil, nil, BlackPawn, nil},
					{nil, WhitePawn, nil, nil, nil, nil, BlackPawn, nil},
					{nil, WhitePawn, nil, nil, nil, nil, BlackPawn, nil},
					{WhiteKing, nil, nil, nil, WhitePawn, nil, BlackPawn, BlackKing},
					{nil, WhitePawn, nil, nil, nil, nil, BlackPawn, nil},
					{nil, WhitePawn, nil, nil, nil, nil, BlackPawn, nil},
					{nil, WhitePawn, nil, nil, nil, nil, BlackPawn, nil},
				},
				0,
				PiecePosition{-1, -1},
			),
			sourceRank:            6,
			targetSquare:          PiecePosition{FileIndex: 5, RankIndex: 4},
			enPassantTargetSquare: PiecePosition{FileIndex: 5, RankIndex: 5},
		},
		{
			state: NewGameState(
				Black,
				[8][8]Piece{
					{nil, WhitePawn, nil, nil, nil, nil, BlackPawn, nil},
					{nil, WhitePawn, nil, nil, nil, nil, BlackPawn, nil},
					{nil, WhitePawn, nil, nil, nil, nil, BlackPawn, nil},
					{nil, WhitePawn, nil, nil, nil, nil, BlackPawn, nil},
					{WhiteKing, nil, nil, nil, WhitePawn, nil, BlackPawn, BlackKing},
					{nil, WhitePawn, nil, nil, nil, BlackPawn, nil, nil},
					{nil, WhitePawn, nil, nil, nil, nil, BlackPawn, nil},
					{nil, WhitePawn, nil, nil, nil, nil, BlackPawn, nil},
				},
				0,
				PiecePosition{-1, -1},
			),
			sourceRank:            5,
			targetSquare:          PiecePosition{FileIndex: 5, RankIndex: 4},
			enPassantTargetSquare: PiecePosition{FileIndex: -1, RankIndex: -1},
		},
		{
			state: NewGameState(
				Black,
				[8][8]Piece{
					{nil, WhitePawn, nil, nil, nil, nil, BlackPawn, nil},
					{nil, WhitePawn, nil, nil, nil, nil, BlackPawn, nil},
					{nil, WhitePawn, nil, nil, nil, nil, BlackPawn, nil},
					{nil, WhitePawn, nil, nil, nil, nil, BlackPawn, nil},
					{WhiteKing, WhitePawn, nil, nil, nil, nil, BlackPawn, BlackKing},
					{nil, WhitePawn, nil, nil, nil, nil, BlackPawn, nil},
					{nil, WhitePawn, nil, nil, nil, nil, BlackPawn, nil},
					{nil, WhitePawn, nil, nil, nil, nil, BlackPawn, nil},
				},
				0,
				PiecePosition{-1, -1},
			),
			sourceRank:            6,
			targetSquare:          PiecePosition{FileIndex: 7, RankIndex: 4},
			enPassantTargetSquare: PiecePosition{FileIndex: -1, RankIndex: -1},
		},
	}

	for _, test := range tests {
		chosenPiece := test.state.GetPieceAt(test.targetSquare.FileIndex, test.sourceRank)
		chosenPawn, ok := chosenPiece.(pawn)
		if !ok {
			log.Fatal("Incorrect test case, chosen piece is not a pawn")
		}
		enPassantTargetSquare := chosenPawn.getEnPassantTargetSquare(test.state, test.sourceRank, test.targetSquare)
		if enPassantTargetSquare != test.enPassantTargetSquare {
			t.Errorf("Invalid en-passant target square. Expected %v, got %v", test.enPassantTargetSquare, enPassantTargetSquare)
		}

	}
}

func TestPawnStringRep(t *testing.T) {
	if WhitePawn.String() != "P" {
		t.Errorf("Incorrect string representation of WhitePawn, Expected P, got %s", WhitePawn.String())
	}

	if BlackPawn.String() != "p" {
		t.Errorf("Incorrect string representation of BlackPawn, Expected p, got %s", BlackPawn.String())
	}

	var neitherColorPawn pawn = 0
	if neitherColorPawn.String() != "" {
		t.Errorf("Incorrect string representation of non-colored pawn, Expected blank string, got %s", neitherColorPawn.String())
	}
}
