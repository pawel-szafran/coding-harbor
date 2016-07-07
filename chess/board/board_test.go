package board

import "testing"

const (
	S = safeSquare
	P = squareWithPiece
	C = capturedSquare
)

func TestNewBoard(t *testing.T) {
	b := New(Size{2, 3})
	assertBoard(t, b, wantBoard{
		squares: [][]square{
			{S, S, S},
			{S, S, S},
		},
		curPos: Pos{0, -1},
	})
}

func TestCopyBoard(t *testing.T) {
	b := New(Size{3, 4})
	b.setSquare(Pos{1, 0}, squareWithPiece)
	b.CurPos = Pos{0, 1}
	bc := b.Copy()
	assertBoardWithDesc(t, "Board Copy", bc, wantBoard{
		squares: [][]square{
			{S, S, S, S},
			{P, S, S, S},
			{S, S, S, S},
		},
		curPos: Pos{0, 1},
	})
	b.setSquare(Pos{0, 2}, capturedSquare)
	b.CurPos = Pos{1, 0}
	bc.setSquare(Pos{2, 1}, capturedSquare)
	bc.CurPos = Pos{2, 3}
	assertBoardWithDesc(t, "Modified Board", b, wantBoard{
		squares: [][]square{
			{S, S, C, S},
			{P, S, S, S},
			{S, S, S, S},
		},
		curPos: Pos{1, 0},
	})
	assertBoardWithDesc(t, "Modified Board Copy", bc, wantBoard{
		squares: [][]square{
			{S, S, S, S},
			{P, S, S, S},
			{S, C, S, S},
		},
		curPos: Pos{2, 3},
	})
}

func TestMoveToNextSafeSquare(t *testing.T) {
	b := New(Size{3, 2})
	assertMoveTo(t, b, Pos{0, 0})
	assertMoveTo(t, b, Pos{0, 1})
	assertMoveTo(t, b, Pos{1, 0})
	b.setSquare(Pos{1, 1}, squareWithPiece)
	b.setSquare(Pos{2, 0}, capturedSquare)
	assertMoveTo(t, b, Pos{2, 1})
	assertNoMoreSafeSquares(t, b, Pos{2, 1})
	assertNoMoreSafeSquares(t, b, Pos{2, 1})
}

func assertMoveTo(t *testing.T, b *Board, pos Pos) {
	if !b.MoveToNextSafeSquare() {
		t.Fatalf("Doesn't move when current position is %v", b.CurPos)
	}
	assertCurPos(t, b, pos)
}

func assertNoMoreSafeSquares(t *testing.T, b *Board, pos Pos) {
	if b.MoveToNextSafeSquare() {
		t.Fatalf("Still has safe square at %v", b.CurPos)
	}
	assertCurPos(t, b, pos)
}

func assertCurPos(t *testing.T, b *Board, pos Pos) {
	if b.CurPos != pos {
		t.Fatalf("Want current position %v, got %v", pos, b.CurPos)
	}
}

func TestPlacePieceOk(t *testing.T) {
	b := New(Size{2, 3})
	b.CurPos = Pos{0, 1}
	b.setSquare(Pos{0, 2}, capturedSquare)
	p := testPiece{captureSquares: []Pos{{0, 2}, {1, 1}}}
	if !b.PlacePiece(p) {
		t.Errorf("Doesn't place")
	}
	assertBoard(t, b, wantBoard{
		squares: [][]square{
			{S, P, C},
			{S, C, S},
		},
		curPos: Pos{0, 1},
	})
}

func TestPlacePieceFails(t *testing.T) {
	b := New(Size{2, 3})
	b.CurPos = Pos{0, 1}
	b.setSquare(Pos{0, 2}, squareWithPiece)
	p := testPiece{captureSquares: []Pos{{0, 2}, {1, 1}}}
	if b.PlacePiece(p) {
		t.Errorf("Places")
	}
}

type testPiece struct {
	captureSquares []Pos
}

func (p testPiece) CaptureSquares(b *Board, capture func(Pos) bool) bool {
	captured := false
	for _, pos := range p.captureSquares {
		captured = captured || capture(pos)
	}
	return captured
}

type wantBoard struct {
	squares [][]square
	curPos  Pos
}

func assertBoard(t *testing.T, b *Board, want wantBoard) {
	assertBoardWithDesc(t, "", b, want)
}

func assertBoardWithDesc(t *testing.T, comment string, b *Board, want wantBoard) {
	if comment != "" {
		comment += ": "
	}
	if wantRows := int8(len(want.squares)); b.Rows() != wantRows {
		t.Errorf("%sWant %d rows, got %d", comment, b.Rows(), wantRows)
	}
	if wantCols := int8(len(want.squares[0])); b.Cols() != wantCols {
		t.Errorf("%sWant %d rows, got %d", comment, b.Cols(), wantCols)
	}
	for row := int8(0); row < b.Rows(); row++ {
		for col := int8(0); col < b.Cols(); col++ {
			wantSquare := want.squares[row][col]
			if !b.checkSquare(Pos{row, col}, wantSquare) {
				t.Errorf("%sWant square at [%d,%d] to be %s", comment, row, col, wantSquare)
			}
		}
	}
	if b.CurPos != want.curPos {
		t.Errorf("Want current position %v, got %v", want.curPos, b.CurPos)
	}
}
