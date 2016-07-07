package board

import (
	"math/rand"
	"reflect"
	"testing"
)

func TestPoolCreatesFreshCopyWhenEmpty(t *testing.T) {
	pool := NewPool()
	b := newRandomBoard()
	bc := pool.Copy(b)
	assertEqualBoards(t, bc, b)
	assertDiffBoardAddr(t, bc, b)
}

func TestPoolUsesRecycledBoards(t *testing.T) {
	pool := NewPool()
	b := newRandomBoard()
	b1 := newRandomBoard()
	pool.Recycle(b1)
	bc := pool.Copy(b)
	assertEqualBoards(t, bc, b)
	assertSameBoardAddr(t, bc, b1)
}

func TestPoolRecyclesInLIFO(t *testing.T) {
	pool := NewPool()
	b := newRandomBoard()
	b1 := newRandomBoard()
	b2 := newRandomBoard()
	pool.Recycle(b1)
	pool.Recycle(b2)
	bc := pool.Copy(b)
	assertSameBoardAddr(t, bc, b2)
	bc = pool.Copy(b)
	assertSameBoardAddr(t, bc, b1)
}

func newRandomBoard() *Board {
	size := Size{3, 4}
	b := New(size)
	b.setSquare(newRandomPos(size), squareWithPiece)
	b.CurPos = newRandomPos(size)
	return b
}

func newRandomPos(s Size) Pos {
	r := func(n int8) int8 {
		return int8(rand.Intn(int(n)))
	}
	return Pos{r(s.Rows), r(s.Cols)}
}

func assertEqualBoards(t *testing.T, got, want *Board) {
	if !reflect.DeepEqual(want, got) {
		t.Errorf("Boards are not equal, want %v, got %v", want, got)
	}
}

func assertSameBoardAddr(t *testing.T, left, right *Board) {
	if left != right {
		t.Errorf("Board addresses are different")
	}
}

func assertDiffBoardAddr(t *testing.T, left, right *Board) {
	if left == right {
		t.Errorf("Board addresses are the same")
	}
}
