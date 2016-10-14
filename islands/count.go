package islands

type Map [][]bool

func CountIslands(m Map) (islands uint) {
	landToVisit := filterLand(m)
	for pos := range landToVisit {
		islands++
		visitWholeIsland(landToVisit, pos)
	}
	return
}

func filterLand(m Map) map[pos]bool {
	land := make(map[pos]bool)
	for row := 0; row < len(m); row++ {
		for col := 0; col < len(m[0]); col++ {
			if m[row][col] {
				land[pos{row, col}] = true
			}
		}
	}
	return land
}

func visitWholeIsland(landToVisit map[pos]bool, pos pos) {
	delete(landToVisit, pos)
	for _, nextPos := range pos.posibleMoves() {
		if landToVisit[nextPos] {
			visitWholeIsland(landToVisit, nextPos)
		}
	}
}

type pos struct {
	row, col int
}

func (p pos) posibleMoves() []pos {
	return []pos{
		pos{p.row + 1, p.col},
		pos{p.row - 1, p.col},
		pos{p.row, p.col + 1},
		pos{p.row, p.col - 1},
	}
}
