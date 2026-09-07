package coremlcompiler

import "math"

// Float16bits returns the IEEE 754 binary16 representation of f, with the sign,
// exponent, and fraction in the same bit positions the format defines. It
// rounds to nearest, ties to even, the mode IEEE 754 specifies and the mode
// coremltools' numpy float16 conversion uses; values too large for binary16
// become an infinity of the same sign, and values too small become a subnormal
// or a signed zero.
//
// It is the inverse of [Float16frombits], and the analogue of
// [math.Float32bits] for the half-precision weights Core ML models are built
// from.
func Float16bits(f float32) uint16 {
	b := math.Float32bits(f)
	sign := uint16(b >> 16 & 0x8000)
	exp := int32(b >> 23 & 0xff)
	mant := b & 0x7fffff

	switch exp {
	case 0xff: // Inf or NaN
		if mant != 0 {
			// Keep it a NaN: a zeroed payload would land on infinity.
			return sign | 0x7e00
		}
		return sign | 0x7c00
	case 0: // signed zero (float32 subnormals are far below binary16 range)
		return sign
	}

	e := exp - 127 + 15
	switch {
	case e >= 0x1f:
		return sign | 0x7c00 // overflow to infinity
	case e <= 0:
		if e < -10 {
			// Below half the smallest subnormal: rounds to zero.
			return sign
		}
		// Subnormal: restore the implicit bit and shift into place.
		mant |= 0x800000
		shift := uint32(14 - e)
		return sign | uint16(roundToNearestEven(mant, shift))
	default:
		m := roundToNearestEven(mant, 13)
		if m > 0x3ff { // rounding carried into the exponent
			m = 0
			e++
			if e >= 0x1f {
				return sign | 0x7c00
			}
		}
		return sign | uint16(e)<<10 | uint16(m)
	}
}

// roundToNearestEven shifts v right by shift bits, rounding the discarded bits
// to nearest and breaking ties toward an even result.
func roundToNearestEven(v, shift uint32) uint32 {
	if shift == 0 || shift >= 32 {
		if shift >= 32 {
			return 0
		}
		return v
	}
	half := uint32(1) << (shift - 1)
	rem := v & (1<<shift - 1)
	q := v >> shift
	if rem > half || (rem == half && q&1 == 1) {
		q++
	}
	return q
}

// Float16frombits returns the float32 value of the IEEE 754 binary16 number
// h. Every binary16 value is exactly representable in float32, so the
// conversion is lossless. It is the inverse of [Float16bits].
func Float16frombits(h uint16) float32 {
	return float16ToFloat32(h)
}

// Float16Bytes encodes vals as little-endian binary16, the layout MIL blob
// storage holds an fp16 tensor in.
func Float16Bytes(vals []float32) []byte {
	out := make([]byte, len(vals)*2)
	for i, v := range vals {
		h := Float16bits(v)
		out[i*2] = byte(h)
		out[i*2+1] = byte(h >> 8)
	}
	return out
}
