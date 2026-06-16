package core

import (
	"testing"
)

func TestKingMoves(t *testing.T) {
	set := make(map[Move]struct{})

	testMoves := []TestMove{
		{
			algebraicSquare: "e6",
			correctMoves: []Move{
				{
					Target: PiecePosition{
						FileIndex: 3,
						RankIndex: 6,
					},
				},
				{
					Target: PiecePosition{
						FileIndex: 2,
						RankIndex: 6,
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
						FileIndex: 2,
						RankIndex: 4,
					},
				},
				{
					Target: PiecePosition{
						FileIndex: 4,
						RankIndex: 4,
					},
				},
			},
			piecePosition: PiecePosition{FileIndex: 3, RankIndex: 5},
			state: NewGameState(
				White,
				[8][8]Piece{
					{nil, WhitePawn, nil, nil, nil, nil, BlackPawn, nil},
					{nil, WhitePawn, nil, nil, nil, nil, BlackPawn, nil},
					{nil, WhitePawn, nil, nil, nil, nil, BlackPawn, nil},
					{nil, nil, BlackKing, BlackPawn, WhitePawn, WhiteKing, nil, nil},
					{nil, WhitePawn, nil, nil, nil, nil, BlackPawn, nil},
					{nil, WhitePawn, nil, nil, nil, nil, BlackPawn, nil},
					{nil, WhitePawn, nil, nil, nil, nil, BlackPawn, nil},
					{nil, WhitePawn, nil, nil, nil, nil, BlackPawn, nil},
				},
				0,
				PiecePosition{-1, -1},
			),
		},
		{
			algebraicSquare: "e3",
			correctMoves: []Move{
				{
					Target: PiecePosition{
						FileIndex: 3,
						RankIndex: 1,
					},
				},
				{
					Target: PiecePosition{
						FileIndex: 2,
						RankIndex: 1,
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
						RankIndex: 3,
					},
				},
				{
					Target: PiecePosition{
						FileIndex: 4,
						RankIndex: 3,
					},
				},
			},
			piecePosition: PiecePosition{FileIndex: 3, RankIndex: 2},
			state: NewGameState(
				Black,
				[8][8]Piece{
					{nil, WhitePawn, nil, nil, nil, nil, BlackPawn, nil},
					{nil, WhitePawn, nil, nil, nil, nil, BlackPawn, nil},
					{nil, WhitePawn, nil, nil, nil, nil, BlackPawn, nil},
					{nil, nil, BlackKing, BlackPawn, WhitePawn, WhiteKing, nil, nil},
					{nil, WhitePawn, nil, nil, nil, nil, BlackPawn, nil},
					{nil, WhitePawn, nil, nil, nil, nil, BlackPawn, nil},
					{nil, WhitePawn, nil, nil, nil, nil, BlackPawn, nil},
					{nil, WhitePawn, nil, nil, nil, nil, BlackPawn, nil},
				},
				0,
				PiecePosition{-1, -1},
			),
		},
		{
			algebraicSquare: "a6",
			correctMoves: []Move{
				{
					Target: PiecePosition{
						FileIndex: 0,
						RankIndex: 4,
					},
				},
				{
					Target: PiecePosition{
						FileIndex: 0,
						RankIndex: 6,
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
						FileIndex: 1,
						RankIndex: 6,
					},
				},
			},
			piecePosition: PiecePosition{FileIndex: 0, RankIndex: 5},
			state: NewGameState(
				White,
				[8][8]Piece{
					{nil, WhitePawn, BlackKing, nil, nil, WhiteKing, BlackPawn, nil},
					{nil, WhitePawn, nil, nil, nil, nil, BlackPawn, nil},
					{nil, WhitePawn, nil, nil, nil, nil, BlackPawn, nil},
					{nil, nil, nil, nil, nil, nil, nil, nil},
					{nil, WhitePawn, nil, nil, nil, nil, BlackPawn, nil},
					{nil, WhitePawn, nil, nil, nil, nil, BlackPawn, nil},
					{nil, WhitePawn, nil, nil, nil, nil, BlackPawn, nil},
					{nil, WhitePawn, nil, nil, nil, nil, BlackPawn, nil},
				},
				0,
				PiecePosition{-1, -1},
			),
		},
		{
			algebraicSquare: "e1",
			correctMoves: []Move{
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
						FileIndex: 3,
						RankIndex: 0,
					},
				},
				{
					Target: PiecePosition{
						FileIndex: 2,
						RankIndex: 0,
					},
				},
			},
			piecePosition: PiecePosition{FileIndex: 4, RankIndex: 0},
			state: NewGameState(
				White,
				[8][8]Piece{
					{nil, WhitePawn, nil, nil, nil, nil, BlackPawn, nil},
					{nil, WhitePawn, nil, nil, nil, nil, BlackPawn, nil},
					{nil, WhitePawn, nil, nil, nil, nil, BlackPawn, nil},
					{nil, WhitePawn, nil, nil, nil, nil, nil, nil},
					{WhiteKing, WhitePawn, nil, nil, nil, nil, BlackPawn, BlackKing},
					{nil, WhitePawn, nil, nil, nil, nil, BlackPawn, nil},
					{nil, WhitePawn, nil, nil, nil, nil, BlackPawn, nil},
					{nil, WhitePawn, nil, nil, nil, nil, BlackPawn, nil},
				},
				WhiteCanKsCastle|WhiteCanQsCastle,
				PiecePosition{-1, -1},
			),
		},
		{
			algebraicSquare: "e1",
			correctMoves: []Move{
				{
					Target: PiecePosition{
						FileIndex: 5,
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
						FileIndex: 2,
						RankIndex: 0,
					},
				},
			},
			piecePosition: PiecePosition{FileIndex: 4, RankIndex: 0},
			state: NewGameState(
				White,
				[8][8]Piece{
					{nil, WhitePawn, nil, nil, nil, nil, BlackPawn, nil},
					{nil, WhitePawn, nil, nil, nil, nil, BlackPawn, nil},
					{nil, WhitePawn, nil, nil, nil, nil, BlackPawn, nil},
					{nil, WhitePawn, nil, nil, nil, nil, nil, nil},
					{WhiteKing, WhitePawn, nil, nil, nil, nil, BlackPawn, BlackKing},
					{nil, WhitePawn, nil, nil, nil, nil, BlackPawn, nil},
					{nil, WhitePawn, nil, nil, nil, nil, BlackPawn, nil},
					{nil, WhitePawn, nil, nil, nil, nil, BlackPawn, nil},
				},
				WhiteCanQsCastle,
				PiecePosition{-1, -1},
			),
		},
		{
			algebraicSquare: "e1",
			correctMoves: []Move{
				{
					Target: PiecePosition{
						FileIndex: 5,
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
						FileIndex: 6,
						RankIndex: 0,
					},
				},
			},
			piecePosition: PiecePosition{FileIndex: 4, RankIndex: 0},
			state: NewGameState(
				White,
				[8][8]Piece{
					{nil, WhitePawn, nil, nil, nil, nil, BlackPawn, nil},
					{nil, WhitePawn, nil, nil, nil, nil, BlackPawn, nil},
					{nil, WhitePawn, nil, nil, nil, nil, BlackPawn, nil},
					{nil, WhitePawn, nil, nil, nil, nil, nil, nil},
					{WhiteKing, WhitePawn, nil, nil, nil, nil, BlackPawn, BlackKing},
					{nil, WhitePawn, nil, nil, nil, nil, BlackPawn, nil},
					{nil, WhitePawn, nil, nil, nil, nil, BlackPawn, nil},
					{nil, WhitePawn, nil, nil, nil, nil, BlackPawn, nil},
				},
				WhiteCanKsCastle,
				PiecePosition{-1, -1},
			),
		},
		{
			algebraicSquare: "e1",
			correctMoves:    []Move{},
			piecePosition:   PiecePosition{FileIndex: 4, RankIndex: 0},
			state: NewGameState(
				White,
				[8][8]Piece{
					{nil, WhitePawn, nil, nil, nil, nil, BlackPawn, nil},
					{nil, WhitePawn, nil, nil, nil, nil, BlackPawn, nil},
					{nil, WhitePawn, nil, BlackBishop, nil, nil, BlackPawn, nil},
					{nil, WhitePawn, nil, nil, nil, nil, nil, nil},
					{WhiteKing, nil, nil, nil, nil, nil, BlackPawn, BlackKing},
					{nil, WhitePawn, nil, nil, nil, nil, BlackPawn, nil},
					{nil, WhitePawn, nil, BlackQueen, nil, nil, BlackPawn, nil},
					{nil, WhitePawn, nil, nil, nil, nil, BlackPawn, nil},
				},
				WhiteCanKsCastle|WhiteCanQsCastle,
				PiecePosition{-1, -1},
			),
		},
		{
			algebraicSquare: "e1",
			correctMoves: []Move{
				{
					Target: PiecePosition{
						FileIndex: 3,
						RankIndex: 0,
					},
				},
				{
					Target: PiecePosition{
						FileIndex: 5,
						RankIndex: 0,
					},
				},
			},
			piecePosition: PiecePosition{FileIndex: 4, RankIndex: 0},
			state: NewGameState(
				White,
				[8][8]Piece{
					{WhiteRook, WhitePawn, nil, nil, nil, nil, BlackPawn, BlackRook},
					{WhiteKnight, WhitePawn, nil, nil, nil, nil, BlackPawn, nil},
					{nil, WhitePawn, nil, nil, nil, nil, BlackPawn, nil},
					{nil, WhitePawn, nil, nil, nil, nil, BlackPawn, nil},
					{WhiteKing, WhitePawn, nil, nil, nil, nil, BlackPawn, BlackKing},
					{nil, WhitePawn, nil, nil, nil, nil, BlackPawn, nil},
					{WhiteKnight, WhitePawn, nil, nil, nil, nil, BlackPawn, nil},
					{WhiteRook, WhitePawn, nil, nil, nil, nil, BlackPawn, BlackRook},
				},
				WhiteCanKsCastle|WhiteCanQsCastle,
				PiecePosition{-1, -1},
			),
		},
		{
			algebraicSquare: "e8",
			correctMoves: []Move{
				{
					Target: PiecePosition{
						FileIndex: 3,
						RankIndex: 7,
					},
				},
				{
					Target: PiecePosition{
						FileIndex: 5,
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
						FileIndex: 2,
						RankIndex: 7,
					},
				},
			},
			piecePosition: PiecePosition{FileIndex: 4, RankIndex: 7},
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
				BlackCanKsCastle|BlackCanQsCastle,
				PiecePosition{-1, -1},
			),
		},
		{
			algebraicSquare: "e8",
			correctMoves: []Move{
				{
					Target: PiecePosition{
						FileIndex: 3,
						RankIndex: 7,
					},
				},
				{
					Target: PiecePosition{
						FileIndex: 5,
						RankIndex: 7,
					},
				},
				{
					Target: PiecePosition{
						FileIndex: 6,
						RankIndex: 7,
					},
				},
			},
			piecePosition: PiecePosition{FileIndex: 4, RankIndex: 7},
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
				BlackCanKsCastle,
				PiecePosition{-1, -1},
			),
		},
		{
			algebraicSquare: "e8",
			correctMoves: []Move{
				{
					Target: PiecePosition{
						FileIndex: 3,
						RankIndex: 7,
					},
				},
				{
					Target: PiecePosition{
						FileIndex: 5,
						RankIndex: 7,
					},
				},
				{
					Target: PiecePosition{
						FileIndex: 2,
						RankIndex: 7,
					},
				},
			},
			piecePosition: PiecePosition{FileIndex: 4, RankIndex: 7},
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
				BlackCanQsCastle,
				PiecePosition{-1, -1},
			),
		},
		{
			algebraicSquare: "e8",
			correctMoves:    []Move{},
			piecePosition:   PiecePosition{FileIndex: 4, RankIndex: 7},
			state: NewGameState(
				Black,
				[8][8]Piece{
					{WhiteRook, WhitePawn, nil, nil, nil, nil, BlackPawn, nil},
					{nil, WhitePawn, nil, nil, nil, nil, BlackPawn, nil},
					{nil, WhitePawn, nil, nil, WhiteBishop, nil, BlackPawn, nil},
					{nil, WhitePawn, nil, nil, nil, nil, BlackPawn, nil},
					{WhiteKing, WhitePawn, nil, nil, nil, nil, nil, BlackKing},
					{nil, WhitePawn, nil, nil, nil, nil, BlackPawn, nil},
					{nil, WhitePawn, nil, nil, WhiteQueen, nil, BlackPawn, nil},
					{nil, WhitePawn, nil, nil, nil, nil, BlackPawn, nil},
				},
				BlackCanKsCastle|BlackCanQsCastle,
				PiecePosition{-1, -1},
			),
		},
		{
			algebraicSquare: "e1",
			correctMoves: []Move{
				{
					Target: PiecePosition{
						FileIndex: 3,
						RankIndex: 7,
					},
				},
				{
					Target: PiecePosition{
						FileIndex: 5,
						RankIndex: 7,
					},
				},
			},
			piecePosition: PiecePosition{FileIndex: 4, RankIndex: 7},
			state: NewGameState(
				Black,
				[8][8]Piece{
					{WhiteRook, WhitePawn, nil, nil, nil, nil, BlackPawn, BlackRook},
					{WhiteKnight, WhitePawn, nil, nil, nil, nil, BlackPawn, BlackKnight},
					{nil, WhitePawn, nil, nil, nil, nil, BlackPawn, nil},
					{nil, WhitePawn, nil, nil, nil, nil, BlackPawn, nil},
					{WhiteKing, WhitePawn, nil, nil, nil, nil, BlackPawn, BlackKing},
					{nil, WhitePawn, nil, nil, nil, nil, BlackPawn, nil},
					{WhiteKnight, WhitePawn, nil, nil, nil, nil, BlackPawn, BlackKnight},
					{WhiteRook, WhitePawn, nil, nil, nil, nil, BlackPawn, BlackRook},
				},
				BlackCanQsCastle|BlackCanKsCastle,
				PiecePosition{-1, -1},
			),
		},
		{
			algebraicSquare: "a3",
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
						RankIndex: 1,
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
						RankIndex: 1,
					},
				},
			},
			piecePosition: PiecePosition{FileIndex: 0, RankIndex: 2},
			state: NewGameState(
				Black,
				[8][8]Piece{
					{nil, WhitePawn, BlackKing, nil, nil, WhiteKing, BlackPawn, nil},
					{nil, WhitePawn, nil, nil, nil, nil, BlackPawn, nil},
					{nil, WhitePawn, nil, nil, nil, nil, BlackPawn, nil},
					{nil, nil, nil, nil, nil, nil, nil, nil},
					{nil, WhitePawn, nil, nil, nil, nil, BlackPawn, nil},
					{nil, WhitePawn, nil, nil, nil, nil, BlackPawn, nil},
					{nil, WhitePawn, nil, nil, nil, nil, BlackPawn, nil},
					{nil, WhitePawn, nil, nil, nil, nil, BlackPawn, nil},
				},
				0,
				PiecePosition{-1, -1},
			),
		},
	}

	for _, move := range testMoves {
		chosenPiece := move.state.GetPieceAt(move.piecePosition.FileIndex, move.piecePosition.RankIndex)
		possibleMoves := chosenPiece.GetMoves(move.state, move.piecePosition)
		for _, item := range possibleMoves {
			set[item] = struct{}{}
		}
		if len(possibleMoves) != len(move.correctMoves) {
			t.Errorf("Invalid number of moves for %v king for %v . Expected %v, Got %v", move.algebraicSquare, move.state.turn, len(move.correctMoves), len(possibleMoves))
		}
		for _, correctMove := range move.correctMoves {
			if _, found := set[correctMove]; !found {
				t.Errorf("Valid move %v missing from the set of moves returned for %v king for %v", correctMove.Target, move.algebraicSquare, move.state.turn)
			}
		}
		clear(set)
	}
}

func TestKingStringRep(t *testing.T) {
	if WhiteKing.String() != "K" {
		t.Errorf("Incorrect string representation of WhiteKing, Expected K, got %s", WhiteKing.String())
	}

	if BlackKing.String() != "k" {
		t.Errorf("Incorrect string representation of BlackKing, Expected k, got %s", BlackKing.String())
	}

	var neitherColorKing king = 0
	if neitherColorKing.String() != "" {
		t.Errorf("Incorrect string representation of non-colored king, Expected blank string, got %s", neitherColorKing.String())
	}
}
