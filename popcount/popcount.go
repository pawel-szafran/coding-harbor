package popcount

type CountFunc func(uint32) uint32

var (
	CountTotallyNaive = func(v uint32) (c uint32) {
		for i := 0; i < 32; i++ {
			if v&0x1 > 0 {
				c++
			}
			v >>= 1
		}
		return
	}
	CountNaive = func(v uint32) (c uint32) {
		for i := 0; i < 32; i++ {
			c += v & 0x1
			v >>= 1
		}
		return
	}
)

func CountSlice(values []uint32, count CountFunc) (c uint32) {
	for _, v := range values {
		c += count(v)
	}
	return
}
