package popcount

func CountNaive(v uint32) (c uint32) {
	for i := 0; i < 32; i++ {
		if v&0x1 > 0 {
			c++
		}
		v >>= 1
	}
	return
}
