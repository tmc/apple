package ext4

// bitmap is an ext4 allocation bitmap: one bit per object, bit n stored in the
// least-significant-bit-first order the kernel uses (byte n/8, bit n%8).
type bitmap struct {
	bits []byte
	n    int // number of meaningful bits
}

// newBitmap returns a bitmap able to hold n bits, all clear.
func newBitmap(n int) *bitmap {
	return &bitmap{bits: make([]byte, (n+7)/8), n: n}
}

// set marks bit i.
func (b *bitmap) set(i int) {
	b.bits[i/8] |= 1 << uint(i%8)
}

// get reports whether bit i is set.
func (b *bitmap) get(i int) bool {
	return b.bits[i/8]&(1<<uint(i%8)) != 0
}

// markPadding sets every bit at or above n through the end of the final byte,
// matching the kernel convention that trailing bits beyond the valid range are
// recorded as in-use.
func (b *bitmap) markPadding() {
	for i := b.n; i < len(b.bits)*8; i++ {
		b.set(i)
	}
}

// countClear returns the number of clear bits in [0, n).
func (b *bitmap) countClear() int {
	c := 0
	for i := 0; i < b.n; i++ {
		if !b.get(i) {
			c++
		}
	}
	return c
}
