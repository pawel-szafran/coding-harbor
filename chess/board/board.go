package board

type (
	Size struct {
		Rows, Cols int8
	}
	Board struct {
		squares [][]square
		CurPos  Pos
	}
	Pos struct {
		Row, Col int8
	}
	Piece interface {
		CaptureSquares(b *Board, capture func(ps ...Pos) bool) bool
	}
	square int8
)

//go:generate stringer -type=square
const (
	safeSquare square = iota
	squareWithPiece
	capturedSquare
)

func New(size Size) *Board {
	squares := make([][]square, size.Rows)
	allSquares := make([]square, size.Rows*size.Cols)
	for row := range squares {
		squares[row], allSquares = allSquares[:size.Cols], allSquares[size.Cols:]
	}
	return &Board{squares, Pos{0, -1}}
}

func (b *Board) Copy() *Board {
	squares := make([][]square, b.Rows())
	allSquares := make([]square, b.Rows()*b.Cols())
	for row := range squares {
		squares[row], allSquares = allSquares[:b.Cols()], allSquares[b.Cols():]
		copy(squares[row], b.squares[row])
	}
	return &Board{squares, b.CurPos}
}

func (b *Board) Rows() int8 { return int8(len(b.squares)) }
func (b *Board) Cols() int8 { return int8(len(b.squares[0])) }

func (b *Board) checkSquare(p Pos) square {
	return b.squares[p.Row][p.Col]
}

func (b *Board) setSquare(p Pos, s square) {
	b.squares[p.Row][p.Col] = s
}

func (b *Board) MoveToNextSafeSquare() bool {
	for b.moveToNextSquare() {
		if b.checkSquare(b.CurPos) == safeSquare {
			return true
		}
	}
	return false
}

func (b *Board) moveToNextSquare() bool {
	switch {
	case b.CurPos.Col+1 < b.Cols():
		b.CurPos.Col++
	case b.CurPos.Row+1 < b.Rows():
		b.CurPos.Row++
		b.CurPos.Col = 0
	default:
		return false
	}
	return true
}

func (b *Board) PlacePiece(piece Piece) bool {
	b.setSquare(b.CurPos, squareWithPiece)
	captured := piece.CaptureSquares(b, func(ps ...Pos) bool {
		for _, pos := range ps {
			if b.contains(pos) {
				if b.checkSquare(pos) == squareWithPiece {
					return true
				}
				b.setSquare(pos, capturedSquare)
			}
		}
		return false
	})
	return !captured
}

func (b *Board) contains(p Pos) bool {
	return p.Row >= 0 && p.Row < b.Rows() &&
		p.Col >= 0 && p.Col < b.Cols()
}
