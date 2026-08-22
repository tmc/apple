//go:build darwin && arm64

package ane

import (
	"fmt"
	"math"
	"testing"
)

// asmF16ToF32 converts a single fp16 bit pattern using the NEON helper.
func asmF16ToF32(h uint16) float32 {
	src := []byte{byte(h), byte(h >> 8)}
	dst := make([]float32, 1)
	cvtF16ToF32(&dst[0], &src[0], 1)
	return dst[0]
}

// asmF32ToF16 converts a single float32 using the NEON helper.
func asmF32ToF16(f float32) uint16 {
	src := []float32{f}
	dst := make([]byte, 2)
	cvtF32ToF16(&dst[0], &src[0], 1)
	return uint16(dst[0]) | uint16(dst[1])<<8
}

func classifyF16(h uint16) string {
	exp := (h >> 10) & 0x1F
	frac := h & 0x3FF
	switch {
	case exp == 0 && frac == 0:
		return "zero"
	case exp == 0:
		return "subnormal"
	case exp == 0x1F && frac == 0:
		return "inf"
	case exp == 0x1F:
		return "nan"
	default:
		return "normal"
	}
}

// classifyF32 classifies a float32 by what fp16 category it should land in.
func classifyF32(b uint32) string {
	exp := int((b>>23)&0xFF) - 127
	frac := b & 0x7FFFFF
	switch {
	case exp == 128 && frac != 0:
		return "nan"
	case exp == 128:
		return "inf"
	case exp == -127 && frac == 0:
		return "zero"
	case exp > 15:
		return "overflow" // above fp16 max magnitude
	case exp >= -14:
		return "normal"
	case exp >= -25:
		return "subnormal" // rounds into fp16 subnormal range
	default:
		return "underflow" // rounds to zero
	}
}

type divergence struct {
	in   uint32
	got  uint32 // scalar result
	want uint32 // assembly result
}

type tally struct {
	counts  map[string]int
	total   map[string]int
	samples map[string][]divergence
}

func newTally() *tally {
	return &tally{
		counts:  map[string]int{},
		total:   map[string]int{},
		samples: map[string][]divergence{},
	}
}

func (t *tally) seen(cat string) { t.total[cat]++ }

func (t *tally) add(cat string, d divergence) {
	t.counts[cat]++
	if len(t.samples[cat]) < 5 {
		t.samples[cat] = append(t.samples[cat], d)
	}
}

func (t *tally) report(tb testing.TB, title string, inW, outW int) int {
	tb.Helper()
	cats := []string{"zero", "subnormal", "normal", "overflow", "underflow", "inf", "nan"}
	sum := 0
	tb.Logf("=== %s ===", title)
	tb.Logf("%-12s %12s %12s", "category", "population", "divergences")
	for _, c := range cats {
		if t.total[c] == 0 && t.counts[c] == 0 {
			continue
		}
		tb.Logf("%-12s %12d %12d", c, t.total[c], t.counts[c])
		sum += t.counts[c]
	}
	tb.Logf("%-12s %12d %12d", "TOTAL", sumMap(t.total), sum)
	for _, c := range cats {
		for _, d := range t.samples[c] {
			tb.Logf("  [%s] in=%#0*x scalar=%#0*x asm=%#0*x", c, inW, d.in, outW, d.got, outW, d.want)
		}
	}
	return sum
}

func sumMap(m map[string]int) int {
	s := 0
	for _, v := range m {
		s += v
	}
	return s
}

// TestFP16ToFloat32Exhaustive compares the scalar fp16ToFloat32 against the
// NEON FCVTL path over every one of the 65536 fp16 bit patterns.
func TestFP16ToFloat32Exhaustive(t *testing.T) {
	tl := newTally()
	for i := range 1 << 16 {
		h := uint16(i)
		cat := classifyF16(h)
		tl.seen(cat)
		got := math.Float32bits(fp16ToFloat32(h))
		want := math.Float32bits(asmF16ToF32(h))
		if got != want {
			tl.add(cat, divergence{in: uint32(h), got: got, want: want})
		}
	}
	if n := tl.report(t, "fp16 -> float32 (exhaustive, 65536 inputs)", 4, 8); n != 0 {
		t.Errorf("fp16ToFloat32 diverges from NEON FCVTL on %d of 65536 inputs", n)
	}
}

// float32ToFP16Cases returns the mandatory boundary population.
func float32ToFP16Cases() []uint32 {
	var cases []uint32
	add := func(f float32) { cases = append(cases, math.Float32bits(f)) }
	addBits := func(b uint32) { cases = append(cases, b) }

	// Zeros.
	add(0)
	add(float32(math.Copysign(0, -1)))

	// Every fp16 value round-tripped through float32: covers all fp16
	// subnormal-producing exponents, all normals, inf and NaN payloads.
	for i := range 1 << 16 {
		v := asmF16ToF32(uint16(i))
		add(v)
	}

	// fp16 boundaries.
	add(6.103515625e-05)  // fp16 min normal 2^-14
	add(6.0975552e-05)    // just below min normal
	add(5.9604645e-08)    // fp16 min subnormal 2^-24
	add(2.9802322e-08)    // half of min subnormal: exact tie to zero/min-sub
	add(2.9802322e-08*3)  // 1.5x min subnormal: tie
	add(65504)            // fp16 max
	add(65519.99)         // just below the overflow threshold 65520
	add(65520)            // exactly halfway: rounds to inf under RNE
	add(65535)            // above -> inf
	add(131008)
	add(math.Float32frombits(0x477fefff)) // 65503.99...

	// Exact ties across the normal range: mantissa low 13 bits == 0x1000.
	for e := -14; e <= 15; e++ {
		base := uint32((e+127)<<23) | 0x1000
		addBits(base)          // tie, even -> down
		addBits(base | 0x2000) // tie, odd  -> up
		addBits(base | 0x0001) // just above tie -> up
		addBits(base | 0x0fff) // just below next tie
		addBits(base | 0x8000_0000)
		addBits(base | 0x2000 | 0x8000_0000)
	}
	// Ties in the subnormal-producing range.
	for e := -25; e < -14; e++ {
		base := uint32((e+127)<<23)
		addBits(base)
		addBits(base | 0x0040_0000)
		addBits(base | 0x7fffff)
		addBits(base | 0x8000_0000)
	}

	// Infinities.
	addBits(0x7F800000)
	addBits(0xFF800000)
	// NaNs with payloads: quiet, signaling, low-bit-only payloads.
	addBits(0x7FC00000) // qNaN
	addBits(0xFFC00000)
	addBits(0x7F800001) // sNaN, payload in bits the fp16 mantissa drops
	addBits(0x7FBFFFFF) // sNaN max payload
	addBits(0x7FC00001)
	addBits(0x7FFFFFFF)
	addBits(0xFF800001)
	addBits(0x7F802000) // sNaN with payload surviving the >>13 shift
	addBits(0x7FA00000)
	return cases
}

// splitmix64 is a deterministic, seedless-at-runtime PRNG so the sampled
// population is identical on every run.
type splitmix64 uint64

func (s *splitmix64) next() uint64 {
	*s += 0x9E3779B97F4A7C15
	z := uint64(*s)
	z = (z ^ (z >> 30)) * 0xBF58476D1CE4E5B9
	z = (z ^ (z >> 27)) * 0x94D049BB133111EB
	return z ^ (z >> 31)
}

const sampleCount = 2_000_000

// TestFloat32ToFP16Differential compares the scalar float32ToFP16 against the
// NEON FCVTN path over a boundary population plus a fixed 2M-sample sweep.
func TestFloat32ToFP16Differential(t *testing.T) {
	tl := newTally()
	check := func(b uint32) {
		f := math.Float32frombits(b)
		cat := classifyF32(b)
		tl.seen(cat)
		got := float32ToFP16(f)
		want := asmF32ToF16(f)
		if got != want {
			tl.add(cat, divergence{in: b, got: uint32(got), want: uint32(want)})
		}
	}

	for _, b := range float32ToFP16Cases() {
		check(b)
	}
	rng := splitmix64(0x5EED_1234_5678_9ABC)
	for range sampleCount {
		check(uint32(rng.next()))
	}
	// A second sample biased into the fp16-representable exponent range, so
	// the sweep is not dominated by overflow/underflow.
	for range sampleCount {
		r := rng.next()
		exp := uint32(112+(r>>52)%36) << 23 // exponents 2^-15 .. 2^20
		b := uint32(r)&0x807FFFFF | exp
		check(b)
	}

	if n := tl.report(t, fmt.Sprintf("float32 -> fp16 (boundaries + %d samples)", 2*sampleCount), 8, 4); n != 0 {
		t.Errorf("float32ToFP16 diverges from NEON FCVTN on %d inputs", n)
	}
}

// TestFloat32ToFP16ErrorMagnitude characterizes the divergence found by
// TestFloat32ToFP16Differential: how far the scalar result is from the
// hardware result, in fp16 ULPs, over the finite non-overflow population.
func TestFloat32ToFP16ErrorMagnitude(t *testing.T) {
	buckets := map[string]int{}
	rng := splitmix64(0x5EED_1234_5678_9ABC)
	n := 0
	for range sampleCount {
		r := rng.next()
		exp := uint32(112+(r>>52)%33) << 23 // 2^-15 .. 2^17
		b := uint32(r)&0x807FFFFF | exp
		f := math.Float32frombits(b)
		if cat := classifyF32(b); cat == "nan" || cat == "inf" || cat == "overflow" {
			continue
		}
		n++
		got, want := float32ToFP16(f), asmF32ToF16(f)
		switch {
		case got == want:
			buckets["exact"]++
		case want-got == 1 || got-want == 1:
			buckets["1 ULP"]++
		case want == 0 && got == 0x8000 || want == 0x8000 && got == 0:
			buckets["sign of zero"]++
		default:
			buckets["> 1 ULP"]++
		}
	}
	t.Logf("=== float32 -> fp16 error magnitude (%d finite in-range samples) ===", n)
	for _, k := range []string{"exact", "1 ULP", "> 1 ULP", "sign of zero"} {
		if v := buckets[k]; v > 0 || k == "exact" {
			t.Logf("%-14s %10d  (%.2f%%)", k, v, 100*float64(v)/float64(n))
		}
	}
}

// TestFP16BatchPaths exercises the assembly's 8-at-a-time loop and its scalar
// tail together, at every length from 0 to 40, against the scalar helpers.
func TestFP16BatchPaths(t *testing.T) {
	rng := splitmix64(0xABCD_EF01_2345_6789)
	const maxN = 40

	// float32 -> fp16 batches.
	f32 := make([]float32, maxN)
	for i := range f32 {
		r := rng.next()
		exp := uint32(112+(r>>52)%30) << 23
		f32[i] = math.Float32frombits(uint32(r)&0x807FFFFF | exp)
	}
	bad := 0
	for n := range maxN + 1 {
		dst := make([]byte, 2*maxN)
		if n > 0 {
			cvtF32ToF16(&dst[0], &f32[0], n)
		}
		for i := range n {
			asm := uint16(dst[2*i]) | uint16(dst[2*i+1])<<8
			sc := float32ToFP16(f32[i])
			if asm != sc {
				if bad < 5 {
					t.Logf("  batch w n=%d i=%d in=%#08x scalar=%#04x asm=%#04x",
						n, i, math.Float32bits(f32[i]), sc, asm)
				}
				bad++
			}
		}
		// Bytes past n must be untouched.
		for i := 2 * n; i < 2*maxN; i++ {
			if dst[i] != 0 {
				t.Errorf("cvtF32ToF16 wrote past n=%d at byte %d", n, i)
				break
			}
		}
	}
	t.Logf("batch float32->fp16: %d divergences over lengths 0..%d", bad, maxN)

	// fp16 -> float32 batches.
	f16 := make([]byte, 2*maxN)
	for i := range maxN {
		h := uint16(rng.next())
		f16[2*i] = byte(h)
		f16[2*i+1] = byte(h >> 8)
	}
	bad2 := 0
	for n := range maxN + 1 {
		dst := make([]float32, maxN)
		if n > 0 {
			cvtF16ToF32(&dst[0], &f16[0], n)
		}
		for i := range n {
			h := uint16(f16[2*i]) | uint16(f16[2*i+1])<<8
			asm := math.Float32bits(dst[i])
			sc := math.Float32bits(fp16ToFloat32(h))
			if asm != sc {
				if bad2 < 5 {
					t.Logf("  batch r n=%d i=%d in=%#04x scalar=%#08x asm=%#08x", n, i, h, sc, asm)
				}
				bad2++
			}
		}
		for i := n; i < maxN; i++ {
			if dst[i] != 0 {
				t.Errorf("cvtF16ToF32 wrote past n=%d at element %d", n, i)
				break
			}
		}
	}
	t.Logf("batch fp16->float32: %d divergences over lengths 0..%d", bad2, maxN)

	if bad != 0 || bad2 != 0 {
		t.Errorf("batch paths diverge: %d write-side, %d read-side", bad, bad2)
	}
}

// TestFP16LittleEndianByteOrder confirms the assembly stores fp16 in
// little-endian byte order, matching binary.LittleEndian.PutUint16 at the
// bulk call sites.
func TestFP16LittleEndianByteOrder(t *testing.T) {
	// 1.0 in fp16 is 0x3C00; little-endian bytes are 00 3C.
	src := []float32{1.0}
	dst := make([]byte, 2)
	cvtF32ToF16(&dst[0], &src[0], 1)
	if dst[0] != 0x00 || dst[1] != 0x3C {
		t.Fatalf("cvtF32ToF16(1.0) bytes = %#02x %#02x, want 0x00 0x3c", dst[0], dst[1])
	}
	t.Logf("byte order confirmed little-endian: 1.0 -> % x", dst)
}
