package coremlcompiler

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
const fp16WantReadHash = 0x1b4d98100719b5e5

// fp16ReadPopulation hashes f(h) over all 65536 fp16 bit patterns.
func fp16ReadPopulation(f func(uint16) float32) uint64 {
	h := uint64(fp16HashInit)
	for i := range 1 << 16 {
		fp16mix(&h, uint64(math.Float32bits(f(uint16(i)))))
	}
	return h
}

// TestFP16CrossPackageHashesCoremlcompiler reports this package's scalar fp16 helper hashes.
func TestFP16CrossPackageHashesCoremlcompiler(t *testing.T) {
	if got := fp16ReadPopulation(float16ToFloat32); got != fp16WantReadHash {
		t.Errorf("pkg coremlcompiler fp16->float32 hash = %#016x, want %#016x (scalar copies have drifted)", got, uint64(fp16WantReadHash))
	} else {
		t.Logf("pkg coremlcompiler fp16->float32 hash = %#016x (agrees with all other copies)", got)
	}
	t.Logf("pkg coremlcompiler float32->fp16: UNMEASURED (no float32ToFP16 helper in this package)")
}
