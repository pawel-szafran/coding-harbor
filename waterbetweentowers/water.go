package waterbetweentowers

type Unit uint

func CalcWaterBetweenTowers(towers []Unit) (water Unit) {
	if len(towers) == 0 {
		return
	}
	left, right := leftScanner(towers), rightScanner(towers)
	for left.hasNotMet(right) {
		scanner := withLowerMaxHeight(left, right)
		water += scanner.calcWaterOnCurrTower()
		scanner.move()
	}
	return
}

type scanner struct {
	towers    []Unit
	idx       int
	maxHeight Unit
	moveIdx   func(int) int
}

func leftScanner(towers []Unit) *scanner {
	startIdx := 0
	return &scanner{
		towers:    towers,
		idx:       startIdx,
		maxHeight: towers[startIdx],
		moveIdx:   func(idx int) int { return idx + 1 },
	}
}

func rightScanner(towers []Unit) *scanner {
	startIdx := len(towers) - 1
	return &scanner{
		towers:    towers,
		idx:       startIdx,
		maxHeight: towers[startIdx],
		moveIdx:   func(idx int) int { return idx - 1 },
	}
}

func (s *scanner) hasNotMet(other *scanner) bool {
	return s.idx != other.idx
}

func withLowerMaxHeight(left, right *scanner) *scanner {
	if left.maxHeight <= right.maxHeight {
		return left
	} else {
		return right
	}
}

func (s *scanner) calcWaterOnCurrTower() Unit {
	return s.maxHeight - s.towers[s.idx]
}

func (s *scanner) move() {
	s.idx = s.moveIdx(s.idx)
	s.updateMaxHeight()
}

func (s *scanner) updateMaxHeight() {
	if s.towers[s.idx] > s.maxHeight {
		s.maxHeight = s.towers[s.idx]
	}
}
