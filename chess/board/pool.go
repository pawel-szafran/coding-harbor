package board

type Pool struct {
	boards []*Board
}

func NewPool() *Pool {
	return &Pool{boards: make([]*Board, 0, 1000)}
}

func (p *Pool) Copy(b *Board) *Board {
	if len(p.boards) > 0 {
		bc := p.boards[len(p.boards)-1]
		p.boards = p.boards[:len(p.boards)-1]
		bc.CurPos = b.CurPos
		copy(bc.squares, b.squares)
		return bc
	}
	return b.Copy()
}

func (p *Pool) Recycle(b *Board) {
	p.boards = append(p.boards, b)
}
