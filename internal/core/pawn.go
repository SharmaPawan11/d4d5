package core

type pawn int

const WhitePawn pawn = 1 << 0
const BlackPawn pawn = 1 << 6

const WhiteSecondRank = 1
const BlackSecondRank = 6

func (p pawn) GetMoves(state State, piecePos PiecePosition) (possibleMoves []Move) {
	var direction int
	var targetFileIndex, targetRankIndex int

	if p == WhitePawn {
		direction = 1
	} else if p == BlackPawn {
		direction = -1
	}

	promotionRankIndex := 7
	possiblePromotionPieces := []Piece{WhiteKnight, WhiteQueen, WhiteBishop, WhiteRook}
	if state.turn == Black {
		promotionRankIndex = 0
		possiblePromotionPieces = []Piece{BlackKnight, BlackQueen, BlackBishop, BlackRook}
	}

	// Can move 1 step forward ?
	if state.GetPieceAt(
		piecePos.FileIndex,
		piecePos.RankIndex+direction,
	) == nil && !state.isKingInCheckAfter(
		piecePos.FileIndex,
		piecePos.RankIndex,
		piecePos.FileIndex,
		piecePos.RankIndex+direction,
	) {
		targetRankIndex = piecePos.RankIndex + direction
		possibleMoves = append(possibleMoves, Move{
			Target: PiecePosition{
				FileIndex: piecePos.FileIndex,
				RankIndex: targetRankIndex,
			},
		})
	}

	// Can capture sideways ?
	captureFiles := [2]int{-1, 1}
	for _, file := range captureFiles {
		targetFileIndex = piecePos.FileIndex + file
		targetRankIndex = piecePos.RankIndex + direction

		if targetFileIndex <= -1 || targetFileIndex >= 8 {
			continue
		}

		if state.GetPieceAt(
			targetFileIndex,
			targetRankIndex,
		) != nil && !state.isFriendlyCapture(
			targetFileIndex,
			targetRankIndex,
		) && !state.isKingInCheckAfter(
			piecePos.FileIndex,
			piecePos.RankIndex,
			targetFileIndex,
			targetRankIndex,
		) {
			possibleMoves = append(
				possibleMoves,
				Move{
					Target: PiecePosition{
						FileIndex: targetFileIndex,
						RankIndex: targetRankIndex,
					},
				},
			)
		}

		// Can capture En passant ?
		if state.enPassantTargetSquare.FileIndex != -1 &&
			!state.isKingInCheckAfterEnPassant(
				piecePos.FileIndex,
				piecePos.RankIndex,
				targetFileIndex,
				targetRankIndex,
			) && targetFileIndex == state.enPassantTargetSquare.FileIndex &&
			targetRankIndex == state.enPassantTargetSquare.RankIndex {
			possibleMoves = append(possibleMoves, Move{
				Target: PiecePosition{
					FileIndex: targetFileIndex,
					RankIndex: targetRankIndex,
				},
			})
		}
	}

	// If the targetRank is 8th rank or 1st rank, then we replace each move with 4 moves
	// i.e. Assuming that the pawn is on e7, e8 becomes e8=Q, e8=R, e8=B, e8=N. Similarly,
	// if there is a valid capture on say c8, it becomes cx8=Q, cx8=R, cx8=B, cx8=N.
	// We do not need to care about replacing an en-passant move as that move can't be in
	// our possibleMoves array if the pawn is on 7th rank or 2nd rank.
	if targetRankIndex == promotionRankIndex {
		var promotionMoves []Move
		for _, move := range possibleMoves {
			for _, piece := range possiblePromotionPieces {
				promotionMoves = append(promotionMoves, Move{
					Target: PiecePosition{
						RankIndex: move.Target.RankIndex,
						FileIndex: move.Target.FileIndex,
					},
					PromoteTo: piece,
				})
			}
		}
		clear(possibleMoves)
		possibleMoves = promotionMoves
	}

	// Can move 2 steps forward ?
	pawnStartingRank := WhiteSecondRank
	if state.turn == Black {
		pawnStartingRank = BlackSecondRank
	}
	if piecePos.RankIndex == pawnStartingRank &&
		state.GetPieceAt(piecePos.FileIndex, piecePos.RankIndex+direction) == nil &&
		state.GetPieceAt(piecePos.FileIndex, piecePos.RankIndex+direction*2) == nil &&
		!state.isKingInCheckAfter(
			piecePos.FileIndex,
			piecePos.RankIndex,
			piecePos.FileIndex,
			piecePos.RankIndex+direction*2,
		) {
		possibleMoves = append(
			possibleMoves,
			Move{
				Target: PiecePosition{
					FileIndex: piecePos.FileIndex,
					RankIndex: piecePos.RankIndex + direction*2,
				},
			},
		)
	}

	return
}

func (p pawn) getEnPassantTargetSquare(state State, sourceRank int, targetPos PiecePosition) PiecePosition {

	if (state.turn == White && p == WhitePawn) && (sourceRank == 1) && (targetPos.RankIndex == 3) {
		if (targetPos.FileIndex < 7) && (state.GetPieceAt(targetPos.FileIndex+1, targetPos.RankIndex) == BlackPawn) ||
			(targetPos.FileIndex > 0) && (state.GetPieceAt(targetPos.FileIndex-1, targetPos.RankIndex) == BlackPawn) {
			return PiecePosition{targetPos.FileIndex, 2}
		}
	}

	if (state.turn == Black && p == BlackPawn) && (sourceRank == 6) && (targetPos.RankIndex == 4) {
		if (targetPos.FileIndex < 7) && (state.GetPieceAt(targetPos.FileIndex+1, targetPos.RankIndex) == WhitePawn) ||
			(targetPos.FileIndex > 0) && (state.GetPieceAt(targetPos.FileIndex-1, targetPos.RankIndex) == WhitePawn) {
			return PiecePosition{targetPos.FileIndex, 5}
		}
	}
	return PiecePosition{-1, -1}
}

func (p pawn) Value() int {
	return int(p)
}

func (p pawn) String() string {
	if p == WhitePawn {
		return "P"
	} else if p == BlackPawn {
		return "p"
	}
	return ""
}
