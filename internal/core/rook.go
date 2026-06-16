package core

type rook int

const WhiteRook rook = 1 << 3
const BlackRook rook = 1 << 9

func (r rook) GetMoves(state State, piecePos PiecePosition) (possibleMoves []Move) {
	lines := [][2]int{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
	}

	for _, line := range lines {
		for multiplier := 1; ; multiplier++ {
			fileIncrement := line[0] * multiplier
			rankIncrement := line[1] * multiplier

			fileIndex := piecePos.FileIndex + fileIncrement
			rankIndex := piecePos.RankIndex + rankIncrement

			// Out of bounds - so the line is done and next line should be picked
			if fileIndex >= 8 || rankIndex >= 8 || fileIndex <= -1 || rankIndex <= -1 {
				break
			}

			// The square ain't vacant, so we can not proceed on this line any further
			// provided that we can still capture this square if an enemy piece exists on it
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
				break
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
	}

	return
}

func (r rook) Value() int {
	return int(r)
}

func (r rook) String() string {
	if r == WhiteRook {
		return "R"
	} else if r == BlackRook {
		return "r"
	}
	return ""
}
