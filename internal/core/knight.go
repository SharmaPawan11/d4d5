package core

type knight int

const WhiteKnight knight = 1 << 1
const BlackKnight knight = 1 << 7

func (n knight) GetMoves(state State, piecePos PiecePosition) (possibleMoves []Move) {
	jumps := [][2]int{
		{2, 1},
		{2, -1},
		{1, 2},
		{1, -2},
		{-2, 1},
		{-2, -1},
		{-1, 2},
		{-1, -2},
	}

	for _, jump := range jumps {
		fileIndex := piecePos.FileIndex + jump[0]
		rankIndex := piecePos.RankIndex + jump[1]

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

	return
}

func (n knight) Value() int {
	return int(n)
}

func (n knight) String() string {
	if n == WhiteKnight {
		return "N"
	} else if n == BlackKnight {
		return "n"
	}
	return ""
}
