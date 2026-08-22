//go:build darwin

package ane

import (
	"math"
	"testing"
)

// This file measures whether this package's private scalar fp16 helpers agree
// with the copies in the other packages (x/ane, x/ane/mil, x/ane/model,
// x/coremlcompiler). Each package hashes its helpers' outputs over an
// identical deterministic population; identical hashes mean bit-identical
// behaviour. The population code below is duplicated verbatim in each package.

type fp16prng uint64

func (s *fp16prng) next() uint64 {
	*s += 0x9E3779B97F4A7C15
	z := uint64(*s)
	z = (z ^ (z >> 30)) * 0xBF58476D1CE4E5B9
	z = (z ^ (z >> 27)) * 0x94D049BB133111EB
	return z ^ (z >> 31)
}

func fp16mix(h *uint64, v uint64) {
	for i := range 8 {
		*h ^= (v >> (8 * i)) & 0xFF
		*h *= 1099511628211
	}
}

const fp16HashInit = 14695981039346656037

// Reference hashes. Every package's scalar copy must produce these; a
// mismatch means the copies have drifted apart.
const (
	fp16WantReadHash  = 0x1b4d98100719b5e5
	fp16WantWriteHash = 0xf46383627634c577
)

// fp16ReadPopulation hashes f(h) over all 65536 fp16 bit patterns.
func fp16ReadPopulation(f func(uint16) float32) uint64 {
	h := uint64(fp16HashInit)
	for i := range 1 << 16 {
		fp16mix(&h, uint64(math.Float32bits(f(uint16(i)))))
	}
	return h
}

// fp16WritePopulation hashes f(v) over 1<<20 deterministic float32 samples
// biased into the fp16-representable exponent range, plus exact-tie cases.
func fp16WritePopulation(f func(float32) uint16) uint64 {
	h := uint64(fp16HashInit)
	for e := -25; e <= 16; e++ {
		base := uint32((e+127)<<23) | 0x1000
		for _, b := range []uint32{base, base | 0x2000, base | 1, base | 0x0FFF, base | 0x80000000} {
			fp16mix(&h, uint64(f(math.Float32frombits(b))))
		}
	}
	rng := fp16prng(0x5EED123456789ABC)
	for range 1 << 20 {
		r := rng.next()
		exp := uint32(112+(r>>52)%36) << 23
		fp16mix(&h, uint64(f(math.Float32frombits(uint32(r)&0x807FFFFF|exp))))
	}
	return h
}

// TestFP16CrossPackageHashesAne reports this package's scalar fp16 helper hashes.
func TestFP16CrossPackageHashesAne(t *testing.T) {
	if got := fp16ReadPopulation(fp16ToFloat32); got != fp16WantReadHash {
		t.Errorf("pkg ane fp16->float32 hash = %#016x, want %#016x (scalar copies have drifted)", got, uint64(fp16WantReadHash))
	} else {
		t.Logf("pkg ane fp16->float32 hash = %#016x (agrees with all other copies)", got)
	}
	if got := fp16WritePopulation(float32ToFP16); got != fp16WantWriteHash {
		t.Errorf("pkg ane float32->fp16 hash = %#016x, want %#016x (scalar copies have drifted)", got, uint64(fp16WantWriteHash))
	} else {
		t.Logf("pkg ane float32->fp16 hash = %#016x (agrees with all other copies)", got)
	}
}
