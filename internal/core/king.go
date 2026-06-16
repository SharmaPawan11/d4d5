package core

type king int

const WhiteKing king = 1 << 5
const BlackKing king = 1 << 11

func (k king) GetMoves(state State, piecePos PiecePosition) (possibleMoves []Move) {
	deltas := [][2]int{
		{1, 1},
		{1, 0},
		{1, -1},
		{0, -1},
		{0, 1},
		{-1, -1},
		{-1, 0},
		{-1, 1},
	}

	for _, delta := range deltas {
		fileIndex := piecePos.FileIndex + delta[0]
		rankIndex := piecePos.RankIndex + delta[1]

		if fileIndex > 7 || rankIndex > 7 || fileIndex < 0 || rankIndex < 0 {
			continue
		}

		if state.GetPieceAt(fileIndex, rankIndex) != nil {
			if !state.isFriendlyCapture(fileIndex, rankIndex) && !state.isKingInCheckAfter(piecePos.FileIndex, piecePos.RankIndex, fileIndex, rankIndex) {
				possibleMoves = append(
					possibleMoves,
					Move{
						Target: PiecePosition{
							FileIndex: fileIndex,
							RankIndex: rankIndex,
						},
					},
				)
			}
		} else if !state.isKingInCheckAfter(piecePos.FileIndex, piecePos.RankIndex, fileIndex, rankIndex) {
			possibleMoves = append(
				possibleMoves,
				Move{
					Target: PiecePosition{
						FileIndex: fileIndex,
						RankIndex: rankIndex,
					},
				},
			)
		}
	}

	isKingCurrentlyInCheck := len(state.GetKingAttackers()) > 0
	castlingRank := 0
	if state.turn == Black {
		castlingRank = 7
	}

	if (state.turn == White && state.castlingRights&WhiteCanKsCastle > 0) ||
		(state.turn == Black && state.castlingRights&BlackCanKsCastle > 0) {
		if !isKingCurrentlyInCheck &&
			(state.GetPieceAt(5, castlingRank) == nil && len(state.GetAttackers(5, castlingRank)) == 0) &&
			(state.GetPieceAt(6, castlingRank) == nil && len(state.GetAttackers(6, castlingRank)) == 0) {
			possibleMoves = append(
				possibleMoves,
				Move{
					Target: PiecePosition{
						FileIndex: 6,
						RankIndex: castlingRank,
					},
				},
			)
		}
	}
	if (state.turn == White && state.castlingRights&WhiteCanQsCastle > 0) ||
		(state.turn == Black && state.castlingRights&BlackCanQsCastle > 0) {

		if !isKingCurrentlyInCheck &&
			(state.GetPieceAt(3, castlingRank) == nil && len(state.GetAttackers(3, castlingRank)) == 0) &&
			(state.GetPieceAt(2, castlingRank) == nil && len(state.GetAttackers(2, castlingRank)) == 0) &&
			(state.GetPieceAt(1, castlingRank) == nil) {
			possibleMoves = append(
				possibleMoves,
				Move{
					Target: PiecePosition{
						FileIndex: 2,
						RankIndex: castlingRank,
					},
				},
			)
		}
	}

	return
}

func (k king) Value() int {
	return int(k)
}

func (k king) String() string {
	if k == WhiteKing {
		return "K"
	} else if k == BlackKing {
		return "k"
	}
	return ""
}
