package board

type (
	Size struct {
		Rows, Cols int8
	}
	Board struct {
		size    Size
		squares []byte
		CurPos  Pos
	}
	Pos struct {
		Row, Col int8
	}
	Piece interface {
		CaptureSquares(b *Board, capture func(ps ...Pos) bool) bool
	}
	square byte
)

//go:generate stringer -type=square
const (
	safeSquare square = iota
	squareWithPiece
	capturedSquare
)

func New(size Size) *Board {
	return &Board{
		size:    size,
		squares: newSquares(size),
		CurPos:  Pos{0, -1},
	}
}

func (b *Board) Copy() *Board {
	return &Board{b.size, copySquares(b.squares), b.CurPos}
}

func (b *Board) Rows() int8 { return b.size.Rows }
func (b *Board) Cols() int8 { return b.size.Cols }

func (b *Board) MoveToNextSafeSquare() bool {
	for b.moveToNextSquare() {
		if b.checkSquare(b.CurPos, safeSquare) {
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
				if b.checkSquare(pos, squareWithPiece) {
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

const (
	bitsPerSquare  = 2
	squaresPerByte = 8 / bitsPerSquare
	squareMask     = 0x3
)

func newSquares(size Size) []byte {
	squaresCount := size.Rows * size.Cols
	bytesCount := squaresCount / squaresPerByte
	if squaresCount%squaresPerByte != 0 {
		bytesCount++
	}
	return make([]byte, bytesCount)
}

func copySquares(sqs []byte) []byte {
	sqsCopy := make([]byte, len(sqs))
	copy(sqsCopy, sqs)
	return sqsCopy
}

func (b *Board) checkSquare(p Pos, s square) bool {
	idx, shift := b.findSquare(p)
	return (b.squares[idx]>>shift)&squareMask == byte(s)
}

func (b *Board) setSquare(p Pos, s square) {
	idx, shift := b.findSquare(p)
	b.squares[idx] |= byte(s) << shift
}

func (b *Board) findSquare(p Pos) (idx int8, shift uint8) {
	squareIdx := p.Row*b.size.Cols + p.Col
	idx = squareIdx / squaresPerByte
	shift = uint8((squareIdx % squaresPerByte) * bitsPerSquare)
	return
}
