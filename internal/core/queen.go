package core

type queen int

const WhiteQueen queen = 1 << 4
const BlackQueen queen = 1 << 10

func (q queen) GetMoves(state State, piecePos PiecePosition) (possibleMoves []Move) {
	var rookMoves, bishopMoves []Move
	if q == WhiteQueen {
		rookMoves = WhiteRook.GetMoves(state, piecePos)
		bishopMoves = WhiteBishop.GetMoves(state, piecePos)
	} else if q == BlackQueen {
		rookMoves = BlackRook.GetMoves(state, piecePos)
		bishopMoves = BlackBishop.GetMoves(state, piecePos)
	}
	possibleMoves = append(rookMoves, bishopMoves...)
	return
}

func (q queen) Value() int {
	return int(q)
}

func (q queen) String() string {
	if q == WhiteQueen {
		return "Q"
	} else if q == BlackQueen {
		return "q"
	}
	return ""
}
