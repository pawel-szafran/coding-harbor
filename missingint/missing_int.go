package missingint

// FindMissingInt finds missing int in a slice of ints in linear time and constant space
func FindMissingInt(nums []int) int {
	xorAll, intRange := nums[0], newIntRange(nums[0])
	for _, n := range nums[1:] {
		xorAll ^= n
		intRange.update(n)
	}
	return xorAll ^ intRange.xor()
}

type intRange struct {
	Min, Max int
}

func newIntRange(start int) *intRange {
	return &intRange{Min: start, Max: start}
}

func (r *intRange) update(n int) {
	switch {
	case n < r.Min:
		r.Min = n
	case n > r.Max:
		r.Max = n
	}
}

func (r *intRange) xor() (xor int) {
	for i := r.Max; i >= r.Min; i-- {
		xor ^= i
	}
	return
}
