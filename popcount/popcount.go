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

	CountKernighan = func(v uint32) (c uint32) {
		for v > 0 {
			c++
			v &= v - 1
		}
		return
	}

	CountMapLookup8 = func() CountFunc {

		count8 := make(map[uint32]uint32)
		count8[0] = 0
		for i := uint32(1); i < 1<<8; i++ {
			count8[i] = i&0x1 + count8[i>>1]
		}

		return func(v uint32) (c uint32) {
			const m8 = 0xff
			return count8[v&m8] +
				count8[(v>>8)&m8] +
				count8[(v>>16)&m8] +
				count8[(v>>24)&m8]
		}
	}()

	CountMapLookup16 = func() CountFunc {

		count16 := make(map[uint32]uint32)
		count16[0] = 0
		for i := uint32(1); i < 1<<16; i++ {
			count16[i] = i&0x1 + count16[i>>1]
		}

		return func(v uint32) (c uint32) {
			const m16 = 0xffff
			return count16[v&m16] + count16[(v>>16)&m16]
		}
	}()

	CountTableLookup8 = func() CountFunc {

		const size = 1 << 8
		var count8 [size]uint32
		count8[0] = 0
		for i := uint32(1); i < size; i++ {
			count8[i] = i&0x1 + count8[i>>1]
		}

		return func(v uint32) (c uint32) {
			const m8 = 0xff
			return count8[v&m8] +
				count8[(v>>8)&m8] +
				count8[(v>>16)&m8] +
				count8[(v>>24)&m8]
		}
	}()

	CountTableLookup16 = func() CountFunc {

		const size = 1 << 16
		var count16 [size]uint32
		count16[0] = 0
		for i := uint32(1); i < size; i++ {
			count16[i] = i&0x1 + count16[i>>1]
		}

		return func(v uint32) (c uint32) {
			const m16 = 0xffff
			return count16[v&m16] + count16[(v>>16)&m16]
		}
	}()

	CountParallelNaive = func(v uint32) uint32 {
		const (
			m1  = 0x55555555
			m2  = 0x33333333
			m4  = 0x0f0f0f0f
			m8  = 0x00ff00ff
			m16 = 0x0000ffff
		)
		v = (v & m1) + ((v >> 1) & m1)
		v = (v & m2) + ((v >> 2) & m2)
		v = (v & m4) + ((v >> 4) & m4)
		v = (v & m8) + ((v >> 8) & m8)
		v = (v & m16) + ((v >> 16) & m16)
		return v
	}

	CountParallelSmart = func(v uint32) uint32 {
		const (
			m1 = 0x55555555
			m2 = 0x33333333
			m4 = 0x0f0f0f0f
		)
		v -= ((v >> 1) & m1)
		v = (v & m2) + ((v >> 2) & m2)
		return ((v + (v >> 4)) & m4) * 0x01010101 >> 24
	}

	CountParallelSmartNoMul = func(v uint32) uint32 {
		const (
			m1 = 0x55555555
			m2 = 0x33333333
			m4 = 0x0f0f0f0f
		)
		v -= ((v >> 1) & m1)
		v = (v & m2) + ((v >> 2) & m2)
		v = (v + (v >> 4)) & m4
		v = v + (v >> 8)
		v = v + (v >> 16)
		return v & 0x3f
	}
)

func CountSlice(values []uint32, count CountFunc) (c uint32) {
	for _, v := range values {
		c += count(v)
	}
	return
}
