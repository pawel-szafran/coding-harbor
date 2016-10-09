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
	min, max int
}

func newIntRange(start int) *intRange {
	return &intRange{min: start, max: start}
}

func (r *intRange) update(n int) {
	switch {
	case n < r.min:
		r.min = n
	case n > r.max:
		r.max = n
	}
}

func (r *intRange) xor() (xor int) {
	for n := r.max; n >= r.min; n-- {
		xor ^= n
	}
	return
}
