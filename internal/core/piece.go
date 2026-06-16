package core

type Piece interface {
	GetMoves(state State, piecePos PiecePosition) []Move
	Value() int
}

const (
	AnyWhitePiece          = int(WhiteKing) | int(WhiteQueen) | int(WhiteRook) | int(WhiteBishop) | int(WhiteKnight) | int(WhitePawn)
	AnyBlackPiece          = int(BlackKing) | int(BlackQueen) | int(BlackRook) | int(BlackBishop) | int(BlackKnight) | int(BlackPawn)
	AnyPawn                = int(WhitePawn) | int(BlackPawn)
	AnyKing                = int(WhiteKing) | int(BlackKing)
	BlackRookOrQueen       = int(BlackRook) | int(BlackQueen)
	WhiteRookOrQueen       = int(WhiteRook) | int(WhiteQueen)
	BlackRookQueenOrKing   = BlackRookOrQueen | int(BlackKing)
	WhiteRookQueenOrKing   = WhiteRookOrQueen | int(WhiteKing)
	BlackBishopOrQueen     = int(BlackBishop) | int(BlackQueen)
	WhiteBishopOrQueen     = int(WhiteBishop) | int(WhiteQueen)
	BlackBishopQueenOrKing = BlackBishopOrQueen | int(BlackKing)
	WhiteBishopQueenOrKing = WhiteBishopOrQueen | int(WhiteKing)
)
