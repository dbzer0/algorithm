package popcount

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))]) +
		int(pc[byte(x>>(1*8))]) +
		int(pc[byte(x>>(2*8))]) +
		int(pc[byte(x>>(3*8))]) +
		int(pc[byte(x>>(4*8))]) +
		int(pc[byte(x>>(5*8))]) +
		int(pc[byte(x>>(6*8))]) +
		int(pc[byte(x>>(7*8))])
}

func PopCount2(x uint64) int {
	var r int
	for i := 0; i < 8; i++ {
		r += int(pc[byte(x>>(uint(i)*8))])
	}
	return r
}

// PopCount считает кол-во единиц по алгоритму SWAR.
// https://graphics.stanford.edu/~seander/bithacks.html#CountBitsSetParallel
func PopCountSWAR(x uint64) uint64 {
	x -= (x >> 1) & 0x55555555
	x = ((x >> 2) & 0x33333333) + (x & 0x33333333)
	x = ((x >> 4) + x) & 0x0f0f0f0f
	x += x >> 8
	x += x >> 16
	return x & 0x0000003f
}

// PopCountOptSWAR - это оптимизированный SWAR.
func PopCountOptSWAR(i uint64) uint64 {
	i = i - ((i >> 1) & 0x55555555)
	i = (i & 0x33333333) + ((i >> 2) & 0x33333333)
	return (((i + (i >> 4)) & 0x0F0F0F0F) * 0x01010101) >> 24
}
