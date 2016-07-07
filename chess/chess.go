package chess

import "github.com/pawel-szafran/coding-harbor/chess/board"

func CountSafeBoards(size board.Size, pieces Pieces) int {
	return countSafeBoards(board.New(size), pieces.compact())
}

func countSafeBoards(board *board.Board, pieces pieces) (count int) {
	if pieces.areEmpty() {
		return 1
	}
	if !board.MoveToNextSafeSquare() {
		return 0
	}
	pieces.forEachType(func(piece Piece) {
		boardCopy := board.Copy()
		if boardCopy.PlacePiece(piece) {
			count += countSafeBoards(boardCopy, pieces.copyRemovingOne(piece))
		}
	})
	count += countSafeBoards(board, pieces)
	return
}
