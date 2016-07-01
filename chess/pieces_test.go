package chess

import (
	"reflect"
	"testing"

	"github.com/pawel-szafran/coding-harbor/chess/board"
)

var captureSquaresTests = []struct {
	piece Piece
	board testBoard
}{
	{King, testBoard{
		{0, 0, 0, 0, 0, 0},
		{0, 0, C, C, C, 0},
		{0, 0, C, P, C, 0},
		{0, 0, C, C, C, 0},
		{0, 0, 0, 0, 0, 0},
	}},
	{Rook, testBoard{
		{0, 0, 0, C, 0, 0},
		{0, 0, 0, C, 0, 0},
		{C, C, C, P, C, C},
		{0, 0, 0, C, 0, 0},
		{0, 0, 0, C, 0, 0},
	}},
	{Knight, testBoard{
		{0, 0, C, 0, C, 0},
		{0, C, 0, 0, 0, C},
		{0, 0, 0, P, 0, 0},
		{0, C, 0, 0, 0, C},
		{0, 0, C, 0, C, 0},
	}},
	{Queen, testBoard{
		{0, C, 0, C, 0, C},
		{0, 0, C, C, C, 0},
		{C, C, C, P, C, C},
		{0, 0, C, C, C, 0},
		{0, C, 0, C, 0, C},
	}},
	{Bishop, testBoard{
		{0, C, 0, 0, 0, C},
		{0, 0, C, 0, C, 0},
		{0, 0, 0, P, 0, 0},
		{0, 0, C, 0, C, 0},
		{0, C, 0, 0, 0, C},
	}},
}

func TestCaptureSquares(t *testing.T) {
	for _, tt := range captureSquaresTests {
		b := board.New(tt.board.size())
		b.CurPos = tt.board.curPos()
		captured := make(map[board.Pos]struct{})
		tt.piece.CaptureSquares(b, func(ps ...board.Pos) bool {
			for _, p := range ps {
				captured[p] = struct{}{}
			}
			return false
		})
		if !reflect.DeepEqual(captured, tt.board.captured()) {
			t.Errorf("For %s want %v, got %v", tt.piece, tt.board.captured(), captured)
		}
	}
}

type (
	testBoard [][]square
	square    int8
)

const (
	P square = iota + 1
	C
)

func (b testBoard) size() board.Size {
	return board.Size{Rows: int8(len(b)), Cols: int8(len(b[0]))}
}

func (b testBoard) curPos() board.Pos {
	for r := 0; r < len(b); r++ {
		for c := 0; c < len(b[0]); c++ {
			if b[r][c] == P {
				return board.Pos{Row: int8(r), Col: int8(c)}
			}
		}
	}
	panic("Can't find current position")
}

func (b testBoard) captured() map[board.Pos]struct{} {
	captured := make(map[board.Pos]struct{})
	for r := 0; r < len(b); r++ {
		for c := 0; c < len(b[0]); c++ {
			if b[r][c] == C {
				captured[board.Pos{Row: int8(r), Col: int8(c)}] = struct{}{}
			}
		}
	}
	return captured
}
