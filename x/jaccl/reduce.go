package jaccl

import (
	"encoding/binary"
	"fmt"
	"math"
)

// reduce combines src into dst using the standalone JACCL type table.
func reduce(dst, src []byte, dtype DType, op ReduceOp) error {
	width, err := dtype.Size()
	if err != nil {
		return err
	}
	if !op.valid() || len(dst) != len(src) || len(src)%width != 0 {
		return fmt.Errorf("reduce shape: %w", ErrProtocol)
	}
	for i := 0; i < len(dst); i += width {
		if err := reduceValue(dst[i:i+width], src[i:i+width], dtype, op); err != nil {
			return err
		}
	}
	return nil
}

// reduceRankMajor reduces equally sized rank-major blocks in ascending rank
// order. Standalone JACCL's mesh implementation promises this order, which is
// observable for non-associative floating-point sums.
func reduceRankMajor(dst, gathered []byte, size, blockSize int, dtype DType, op ReduceOp) error {
	if size < 2 || blockSize < 0 || len(dst) != blockSize || len(gathered) != size*blockSize {
		return fmt.Errorf("rank-major reduce shape: %w", ErrProtocol)
	}
	copy(dst, gathered[:blockSize])
	for rank := 1; rank < size; rank++ {
		if err := reduce(dst, gathered[rank*blockSize:(rank+1)*blockSize], dtype, op); err != nil {
			return fmt.Errorf("reduce rank %d: %w", rank, err)
		}
	}
	return nil
}

func reduceValue(dst, src []byte, dtype DType, op ReduceOp) error {
	switch dtype {
	case Bool:
		dst[0] = boolReduce(dst[0], src[0], op)
	case Int8:
		dst[0] = byte(reduceInt8(int8(dst[0]), int8(src[0]), op))
	case UInt8:
		dst[0] = reduceUint8(dst[0], src[0], op)
	case Int16:
		binary.LittleEndian.PutUint16(dst, uint16(reduceInt16(int16(binary.LittleEndian.Uint16(dst)), int16(binary.LittleEndian.Uint16(src)), op)))
	case UInt16:
		binary.LittleEndian.PutUint16(dst, reduceUint16(binary.LittleEndian.Uint16(dst), binary.LittleEndian.Uint16(src), op))
	case Int32:
		binary.LittleEndian.PutUint32(dst, uint32(reduceInt32(int32(binary.LittleEndian.Uint32(dst)), int32(binary.LittleEndian.Uint32(src)), op)))
	case UInt32:
		binary.LittleEndian.PutUint32(dst, reduceUint32(binary.LittleEndian.Uint32(dst), binary.LittleEndian.Uint32(src), op))
	case Int64:
		binary.LittleEndian.PutUint64(dst, uint64(reduceInt64(int64(binary.LittleEndian.Uint64(dst)), int64(binary.LittleEndian.Uint64(src)), op)))
	case UInt64:
		binary.LittleEndian.PutUint64(dst, reduceUint64(binary.LittleEndian.Uint64(dst), binary.LittleEndian.Uint64(src), op))
	case Float16:
		a, b := binary.LittleEndian.Uint16(dst), binary.LittleEndian.Uint16(src)
		if op == Sum {
			if result, ok := halfSumNaN(a, b); ok {
				binary.LittleEndian.PutUint16(dst, result)
				break
			}
			binary.LittleEndian.PutUint16(dst, float32ToHalf(halfToFloat32(a)+halfToFloat32(b)))
			break
		}
		if op == Min {
			if halfToFloat32(a) < halfToFloat32(b) {
				binary.LittleEndian.PutUint16(dst, a)
			} else {
				binary.LittleEndian.PutUint16(dst, b)
			}
			break
		}
		if halfToFloat32(a) > halfToFloat32(b) {
			binary.LittleEndian.PutUint16(dst, a)
		} else {
			binary.LittleEndian.PutUint16(dst, b)
		}
	case BFloat16:
		a, b := binary.LittleEndian.Uint16(dst), binary.LittleEndian.Uint16(src)
		if op == Sum {
			if result, ok := bfloat16SumNaN(a, b); ok {
				binary.LittleEndian.PutUint16(dst, result)
				break
			}
			binary.LittleEndian.PutUint16(dst, float32ToBFloat16(bfloat16ToFloat32(a)+bfloat16ToFloat32(b)))
			break
		}
		if op == Min {
			if bfloat16ToFloat32(a) < bfloat16ToFloat32(b) {
				binary.LittleEndian.PutUint16(dst, a)
			} else {
				binary.LittleEndian.PutUint16(dst, b)
			}
			break
		}
		if bfloat16ToFloat32(a) > bfloat16ToFloat32(b) {
			binary.LittleEndian.PutUint16(dst, a)
		} else {
			binary.LittleEndian.PutUint16(dst, b)
		}
	case Float32:
		a, b := binary.LittleEndian.Uint32(dst), binary.LittleEndian.Uint32(src)
		if op == Sum {
			if result, ok := float32SumNaN(a, b); ok {
				binary.LittleEndian.PutUint32(dst, result)
				break
			}
		}
		binary.LittleEndian.PutUint32(dst, math.Float32bits(reduceFloat32(math.Float32frombits(a), math.Float32frombits(b), op)))
	case Float64:
		a, b := binary.LittleEndian.Uint64(dst), binary.LittleEndian.Uint64(src)
		if op == Sum {
			if result, ok := float64SumNaN(a, b); ok {
				binary.LittleEndian.PutUint64(dst, result)
				break
			}
		}
		binary.LittleEndian.PutUint64(dst, math.Float64bits(reduceFloat64(math.Float64frombits(a), math.Float64frombits(b), op)))
	case Complex64:
		ar, ai := binary.LittleEndian.Uint32(dst), binary.LittleEndian.Uint32(dst[4:])
		br, bi := binary.LittleEndian.Uint32(src), binary.LittleEndian.Uint32(src[4:])
		if op == Sum {
			realBits, realNaN := float32SumNaN(ar, br)
			if !realNaN {
				realBits = math.Float32bits(math.Float32frombits(ar) + math.Float32frombits(br))
			}
			imagBits, imagNaN := float32SumNaN(ai, bi)
			if !imagNaN {
				imagBits = math.Float32bits(math.Float32frombits(ai) + math.Float32frombits(bi))
			}
			binary.LittleEndian.PutUint32(dst, realBits)
			binary.LittleEndian.PutUint32(dst[4:], imagBits)
			break
		}
		a := complex(math.Float32frombits(ar), math.Float32frombits(ai))
		b := complex(math.Float32frombits(br), math.Float32frombits(bi))
		v := reduceComplex64(a, b, op)
		binary.LittleEndian.PutUint32(dst, math.Float32bits(real(v)))
		binary.LittleEndian.PutUint32(dst[4:], math.Float32bits(imag(v)))
	default:
		return fmt.Errorf("reduce dtype %d: %w", dtype, ErrProtocol)
	}
	return nil
}

func halfSumNaN(a, b uint16) (uint16, bool) {
	for _, value := range [2]uint16{a, b} {
		if value&0x7c00 == 0x7c00 && value&0x03ff != 0 {
			return value | 0x0200, true
		}
	}
	return 0, false
}

func bfloat16SumNaN(a, b uint16) (uint16, bool) {
	for _, value := range [2]uint16{a, b} {
		if value&0x7f80 == 0x7f80 && value&0x003f != 0 {
			return value | 0x0040, true
		}
	}
	for _, value := range [2]uint16{a, b} {
		if value&0x7f80 == 0x7f80 && value&0x007f != 0 {
			return value | 0x0040, true
		}
	}
	return 0, false
}

func float32SumNaN(a, b uint32) (uint32, bool) {
	for _, value := range [2]uint32{a, b} {
		if value&0x7f800000 == 0x7f800000 && value&0x003fffff != 0 {
			return value | 0x00400000, true
		}
	}
	if a&0x7f800000 == 0x7f800000 && a&0x007fffff != 0 {
		return a, true
	}
	if b&0x7f800000 == 0x7f800000 && b&0x007fffff != 0 {
		return b, true
	}
	return 0, false
}

func float64SumNaN(a, b uint64) (uint64, bool) {
	for _, value := range [2]uint64{a, b} {
		if value&0x7ff0000000000000 == 0x7ff0000000000000 && value&0x0007ffffffffffff != 0 {
			return value | 0x0008000000000000, true
		}
	}
	if a&0x7ff0000000000000 == 0x7ff0000000000000 && a&0x000fffffffffffff != 0 {
		return a, true
	}
	if b&0x7ff0000000000000 == 0x7ff0000000000000 && b&0x000fffffffffffff != 0 {
		return b, true
	}
	return 0, false
}

func boolReduce(a, b byte, op ReduceOp) byte {
	if op == Min {
		if a != 0 && b != 0 {
			return 1
		}
		return 0
	}
	if a != 0 || b != 0 {
		return 1
	}
	return 0
}
func reduceInt8(a, b int8, op ReduceOp) int8 {
	if op == Sum {
		return a + b
	}
	if op == Min && b < a {
		return b
	}
	if op == Max && b > a {
		return b
	}
	return a
}
func reduceUint8(a, b byte, op ReduceOp) byte {
	if op == Sum {
		return a + b
	}
	if op == Min && b < a {
		return b
	}
	if op == Max && b > a {
		return b
	}
	return a
}
func reduceInt16(a, b int16, op ReduceOp) int16 {
	if op == Sum {
		return a + b
	}
	if op == Min && b < a {
		return b
	}
	if op == Max && b > a {
		return b
	}
	return a
}
func reduceUint16(a, b uint16, op ReduceOp) uint16 {
	if op == Sum {
		return a + b
	}
	if op == Min && b < a {
		return b
	}
	if op == Max && b > a {
		return b
	}
	return a
}
func reduceInt32(a, b int32, op ReduceOp) int32 {
	if op == Sum {
		return a + b
	}
	if op == Min && b < a {
		return b
	}
	if op == Max && b > a {
		return b
	}
	return a
}
func reduceUint32(a, b uint32, op ReduceOp) uint32 {
	if op == Sum {
		return a + b
	}
	if op == Min && b < a {
		return b
	}
	if op == Max && b > a {
		return b
	}
	return a
}
func reduceInt64(a, b int64, op ReduceOp) int64 {
	if op == Sum {
		return a + b
	}
	if op == Min && b < a {
		return b
	}
	if op == Max && b > a {
		return b
	}
	return a
}
func reduceUint64(a, b uint64, op ReduceOp) uint64 {
	if op == Sum {
		return a + b
	}
	if op == Min && b < a {
		return b
	}
	if op == Max && b > a {
		return b
	}
	return a
}
func reduceFloat32(a, b float32, op ReduceOp) float32 {
	if op == Sum {
		return a + b
	}
	if op == Min {
		if a < b {
			return a
		}
		return b
	}
	if a > b {
		return a
	}
	return b
}
func reduceFloat64(a, b float64, op ReduceOp) float64 {
	if op == Sum {
		return a + b
	}
	if op == Min {
		if a < b {
			return a
		}
		return b
	}
	if a > b {
		return a
	}
	return b
}
func reduceComplex64(a, b complex64, op ReduceOp) complex64 {
	if op == Sum {
		return a + b
	}
	if op == Min {
		if complexLess(a, b) {
			return a
		}
		return b
	}
	if complexLess(b, a) {
		return a
	}
	return b
}
func complexLess(a, b complex64) bool {
	return real(a) < real(b) || real(a) == real(b) && imag(a) < imag(b)
}

func halfToFloat32(v uint16) float32 {
	sign, exp, frac := uint32(v&0x8000)<<16, uint32(v>>10)&31, uint32(v&0x3ff)
	if exp == 0 {
		if frac == 0 {
			return math.Float32frombits(sign)
		}
		e := int32(-14)
		for frac&0x400 == 0 {
			frac <<= 1
			e--
		}
		frac &= 0x3ff
		return math.Float32frombits(sign | uint32(e+127)<<23 | frac<<13)
	}
	if exp == 31 {
		return math.Float32frombits(sign | 0x7f800000 | frac<<13)
	}
	return math.Float32frombits(sign | (exp+112)<<23 | frac<<13)
}
func float32ToHalf(v float32) uint16 {
	b := math.Float32bits(v)
	sign := uint16(b>>16) & 0x8000
	exponent := int(b>>23) & 0xff
	frac := b & 0x7fffff
	if exponent == 0xff {
		if frac != 0 {
			// The live arm64 JACCL build uses native Float16 arithmetic and
			// emits this canonical quiet NaN payload.
			return sign | 0x7e00
		}
		return sign | 0x7c00
	}
	exponent -= 127
	if exponent > 15 {
		return sign | 0x7c00
	}
	if exponent < -14 {
		if exponent < -24 {
			return sign
		}
		mantissa := frac | 0x800000
		shift := uint(-exponent - 1)
		return sign | uint16((mantissa+(1<<(shift-1))-1+((mantissa>>shift)&1))>>shift)
	}
	frac += 0xfff + ((frac >> 13) & 1)
	if frac&0x800000 != 0 {
		exponent++
		frac = 0
	}
	if exponent > 15 {
		return sign | 0x7c00
	}
	return sign | uint16(exponent+15)<<10 | uint16(frac>>13)
}
func bfloat16ToFloat32(v uint16) float32 { return math.Float32frombits(uint32(v) << 16) }
func float32ToBFloat16(v float32) uint16 {
	if math.IsNaN(float64(v)) {
		return 0x7fc0
	}
	b := math.Float32bits(v)
	return uint16((b + ((b >> 16) & 1) + 0x7fff) >> 16)
}
