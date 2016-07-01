package chess

import "github.com/pawel-szafran/coding-harbor/chess/board"

func CountSafeBoards(size board.Size, pieces Pieces) int {
	return countSafeBoards(board.New(size), pieces)
}

func countSafeBoards(board *board.Board, pieces Pieces) (count int) {
	if len(pieces) == 0 {
		return 1
	}
	if moved := board.MoveToNextSafeSquare(); !moved {
		if len(pieces) > 0 {
			return 0
		}
		return 1
	}
	for piece, _ := range pieces {
		boardCopy := board.Copy()
		if boardCopy.PlacePiece(piece) {
			count += countSafeBoards(boardCopy, pieces.copyRemovingOne(piece))
		}
	}
	count += countSafeBoards(board, pieces)
	return
}
