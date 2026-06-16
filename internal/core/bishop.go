package core

type bishop int

const WhiteBishop bishop = 1 << 2
const BlackBishop bishop = 1 << 8

func (b bishop) GetMoves(state State, piecePos PiecePosition) (possibleMoves []Move) {
	diagonals := [][2]int{
		{1, 1},
		{-1, 1},
		{-1, -1},
		{1, -1},
	}

	//Outer:
	for _, diagonal := range diagonals {
		for multiplier := 1; ; multiplier++ {
			fileIncrement := diagonal[0] * multiplier
			rankIncrement := diagonal[1] * multiplier

			fileIndex := piecePos.FileIndex + fileIncrement
			rankIndex := piecePos.RankIndex + rankIncrement

			// Out of bounds - so the diagonals is done and next diagonal should be picked
			if fileIndex >= 8 || rankIndex >= 8 || fileIndex <= -1 || rankIndex <= -1 {
				break
			}

			// The square ain't vacant, so we can not proceed on this diagonal any further
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

func (b bishop) Value() int {
	return int(b)
}

func (b bishop) String() string {
	if b == WhiteBishop {
		return "B"
	} else if b == BlackBishop {
		return "b"
	}
	return ""
}
