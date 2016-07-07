package chess

import "github.com/pawel-szafran/coding-harbor/chess/board"

type (
	Piece  int8
	Pieces map[Piece]int8
	pieces [5]int8
)

//go:generate stringer -type=Piece
const (
	King Piece = iota
	Rook
	Knight
	Queen
	Bishop
)

func (ps Pieces) compact() (pieces pieces) {
	for p, n := range ps {
		pieces[p] = n
	}
	return
}

func (ps pieces) areEmpty() bool {
	for _, n := range ps {
		if n > 0 {
			return false
		}
	}
	return true
}

func (ps pieces) forEachType(op func(p Piece)) {
	for p, n := range ps {
		if n > 0 {
			op(Piece(p))
		}
	}
}

func (ps pieces) copyRemovingOne(p Piece) pieces {
	psCopy := ps
	psCopy[p]--
	return psCopy
}

func (p Piece) CaptureSquares(b *board.Board, capture func(ps ...board.Pos) bool) bool {

	row, col := b.CurPos.Row, b.CurPos.Col

	captureRowAndCol := func() bool {
		for r := int8(0); r < b.Rows(); r++ {
			if r != row {
				if capture(pos(r, col)) {
					return true
				}
			}
		}
		for c := int8(0); c < b.Cols(); c++ {
			if c != col {
				if capture(pos(row, c)) {
					return true
				}
			}
		}
		return false
	}
	captureDiagonals := func() bool {
		offset := col - row
		for r := int8(0); r < b.Rows(); r++ {
			if r != row {
				width := (row - r) * 2
				if capture(pos(r, r+offset), pos(r, r+offset+width)) {
					return true
				}
			}
		}
		return false
	}

	switch p {
	case Rook:
		return captureRowAndCol()
	case Bishop:
		return captureDiagonals()
	case Queen:
		return captureRowAndCol() || captureDiagonals()
	case King:
		return capture(
			pos(row-1, col),
			pos(row+1, col),
			pos(row, col-1),
			pos(row, col+1),
			pos(row-1, col-1),
			pos(row-1, col+1),
			pos(row+1, col+1),
			pos(row+1, col-1),
		)
	case Knight:
		return capture(
			pos(row-2, col-1),
			pos(row-2, col+1),
			pos(row+2, col-1),
			pos(row+2, col+1),
			pos(row-1, col+2),
			pos(row+1, col+2),
			pos(row-1, col-2),
			pos(row+1, col-2),
		)
	}
	return false
}

func pos(row, col int8) board.Pos {
	return board.Pos{Row: row, Col: col}
}
