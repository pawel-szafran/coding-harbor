package chess

import "github.com/pawel-szafran/coding-harbor/chess/board"

func CountSafeBoards(size board.Size, pieces Pieces) int {
	return countSafeBoards(board.NewPool(), board.New(size), pieces.compact())
}

func countSafeBoards(pool *board.Pool, board *board.Board, pieces pieces) (count int) {
	if pieces.areEmpty() {
		return 1
	}
	if !board.MoveToNextSafeSquare() {
		return 0
	}
	pieces.forEachType(func(piece Piece) {
		boardCopy := pool.Copy(board)
		if boardCopy.PlacePiece(piece) {
			count += countSafeBoards(pool, boardCopy, pieces.copyRemovingOne(piece))
		}
		pool.Recycle(boardCopy)
	})
	count += countSafeBoards(pool, board, pieces)
	return
}
