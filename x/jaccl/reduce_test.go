package jaccl

import (
	"encoding/binary"
	"math"
	"testing"
)

func TestReduce(t *testing.T) {
	tests := []struct {
		name  string
		dtype DType
		op    ReduceOp
		dst   []byte
		src   []byte
		want  []byte
	}{
		{"uint8 sum", Uint8, Sum, []byte{1, 250}, []byte{2, 10}, []byte{3, 4}},
		{"uint8 min", Uint8, Min, []byte{1, 250}, []byte{2, 10}, []byte{1, 10}},
		{"uint8 max", Uint8, Max, []byte{1, 250}, []byte{2, 10}, []byte{2, 250}},
		{"int32 sum", Int32, Sum, int32Bytes(-2, 10), int32Bytes(5, -3), int32Bytes(3, 7)},
		{"int32 min", Int32, Min, int32Bytes(-2, 10), int32Bytes(5, -3), int32Bytes(-2, -3)},
		{"float32 max", Float32, Max, float32Bytes(1.5, -2), float32Bytes(1.25, 3), float32Bytes(1.5, 3)},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if err := reduce(test.dst, test.src, test.dtype, test.op); err != nil {
				t.Fatal(err)
			}
			if string(test.dst) != string(test.want) {
				t.Fatalf("reduce = %x, want %x", test.dst, test.want)
			}
		})
	}
}

func TestReduceFloat32NaN(t *testing.T) {
	dst := float32Bytes(2)
	src := float32Bytes(float32(math.NaN()))
	if err := reduce(dst, src, Float32, Sum); err != nil {
		t.Fatal(err)
	}
	if !math.IsNaN(float64(math.Float32frombits(binary.LittleEndian.Uint32(dst)))) {
		t.Fatalf("sum = %x, want NaN", dst)
	}
}

func TestReduceFloatMinMaxUsesRightOperandOnUnorderedComparison(t *testing.T) {
	tests := []struct {
		name string
		op   ReduceOp
		dst  float32
		src  float32
		want uint32
	}{
		{"max nan", Max, 1, float32(math.NaN()), math.Float32bits(float32(math.NaN()))},
		{"min nan", Min, 1, float32(math.NaN()), math.Float32bits(float32(math.NaN()))},
		{"max signed zero", Max, math.Float32frombits(0x80000000), 0, 0},
		{"min signed zero", Min, math.Float32frombits(0x80000000), 0, 0},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			dst := float32Bytes(test.dst)
			if err := reduce(dst, float32Bytes(test.src), Float32, test.op); err != nil {
				t.Fatal(err)
			}
			if got := binary.LittleEndian.Uint32(dst); got != test.want {
				t.Fatalf("bits = %#x, want %#x", got, test.want)
			}
		})
	}
}

func TestReduceFloatSumNaNBits(t *testing.T) {
	tests := []struct {
		name  string
		dtype DType
		dst   []byte
		src   []byte
		want  []byte
	}{
		{"float32 positive then negative", Float32, uint32Bytes(0x7fc00000), uint32Bytes(0xffc00000), uint32Bytes(0x7fc00000)},
		{"float32 negative then positive", Float32, uint32Bytes(0xffc00000), uint32Bytes(0x7fc00000), uint32Bytes(0xffc00000)},
		{"float64 positive then negative", Float64, uint64Bytes(0x7ff8000000000000), uint64Bytes(0xfff8000000000000), uint64Bytes(0x7ff8000000000000)},
		{"float64 negative then positive", Float64, uint64Bytes(0xfff8000000000000), uint64Bytes(0x7ff8000000000000), uint64Bytes(0xfff8000000000000)},
		{"float32 signaling left", Float32, uint32Bytes(0x7f800001), uint32Bytes(0x3f800000), uint32Bytes(0x7fc00001)},
		{"float32 signaling right wins", Float32, uint32Bytes(0x7fc00000), uint32Bytes(0x7f800001), uint32Bytes(0x7fc00001)},
		{"float64 signaling left", Float64, uint64Bytes(0x7ff0000000000001), uint64Bytes(0x3ff0000000000000), uint64Bytes(0x7ff8000000000001)},
		{"float64 signaling right wins", Float64, uint64Bytes(0x7ff8000000000000), uint64Bytes(0x7ff0000000000001), uint64Bytes(0x7ff8000000000001)},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if err := reduce(test.dst, test.src, test.dtype, Sum); err != nil {
				t.Fatal(err)
			}
			if string(test.dst) != string(test.want) {
				t.Fatalf("sum = %x, want %x", test.dst, test.want)
			}
		})
	}
}

func TestReduceComplexSumNaNBits(t *testing.T) {
	tests := []struct {
		name string
		dst  []byte
		src  []byte
		want []byte
	}{
		{"quiet", complex64BitBytes(0x7fc00000, 0xffc00000), complex64BitBytes(0xffc00000, 0x7fc00000), complex64BitBytes(0x7fc00000, 0xffc00000)},
		{"signaling", complex64BitBytes(0x7f800001, 0xffc00000), complex64BitBytes(0x7fc00000, 0xff800001), complex64BitBytes(0x7fc00001, 0xffc00001)},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if err := reduce(test.dst, test.src, Complex64, Sum); err != nil {
				t.Fatal(err)
			}
			if string(test.dst) != string(test.want) {
				t.Fatalf("sum = %x, want %x", test.dst, test.want)
			}
		})
	}
}

func TestReduceAllJACCLDTypes(t *testing.T) {
	tests := []struct {
		name  string
		dtype DType
		dst   []byte
		src   []byte
		want  []byte
	}{
		{"bool", Bool, []byte{0}, []byte{1}, []byte{1}},
		{"int8", Int8, []byte{0xfe}, []byte{5}, []byte{3}},
		{"int16", Int16, int16Bytes(-2), int16Bytes(5), int16Bytes(3)},
		{"int32", Int32, int32Bytes(-2), int32Bytes(5), int32Bytes(3)},
		{"int64", Int64, int64Bytes(-2), int64Bytes(5), int64Bytes(3)},
		{"uint8", UInt8, []byte{250}, []byte{10}, []byte{4}},
		{"uint16", UInt16, uint16Bytes(65530), uint16Bytes(10), uint16Bytes(4)},
		{"uint32", UInt32, uint32Bytes(0xfffffff0), uint32Bytes(0x20), uint32Bytes(0x10)},
		{"uint64", UInt64, uint64Bytes(^uint64(0) - 1), uint64Bytes(3), uint64Bytes(1)},
		{"float16", Float16, uint16Bytes(float32ToHalf(1.5)), uint16Bytes(float32ToHalf(2.25)), uint16Bytes(float32ToHalf(3.75))},
		{"bfloat16", BFloat16, uint16Bytes(float32ToBFloat16(1.5)), uint16Bytes(float32ToBFloat16(2.25)), uint16Bytes(float32ToBFloat16(3.75))},
		{"float32", Float32, float32Bytes(1.5), float32Bytes(2.25), float32Bytes(3.75)},
		{"float64", Float64, float64Bytes(1.5), float64Bytes(2.25), float64Bytes(3.75)},
		{"complex64", Complex64, complex64Bytes(complex(1, 2)), complex64Bytes(complex(3, 4)), complex64Bytes(complex(4, 6))},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if err := reduce(test.dst, test.src, test.dtype, Sum); err != nil {
				t.Fatal(err)
			}
			if string(test.dst) != string(test.want) {
				t.Fatalf("reduce = %x, want %x", test.dst, test.want)
			}
		})
	}
}

func TestFloat16Conversions(t *testing.T) {
	tests := []struct {
		bits uint16
		want float32
	}{
		{0x0000, 0},
		{0x0001, math.Float32frombits(0x33800000)},
		{0x3c00, 1},
		{0xc000, -2},
	}
	for _, test := range tests {
		if got := halfToFloat32(test.bits); got != test.want {
			t.Fatalf("halfToFloat32(%#x) = %g, want %g", test.bits, got, test.want)
		}
		if got := float32ToHalf(test.want); got != test.bits {
			t.Fatalf("float32ToHalf(%g) = %#x, want %#x", test.want, got, test.bits)
		}
	}
}

func TestFloat16JACCLBoundaryEncodings(t *testing.T) {
	tests := []struct {
		name  string
		value float32
		want  uint16
	}{
		{"largest finite", 65504, 0x7bff},
		{"round overflow", 65520, 0x7c00},
		{"positive infinity", float32(math.Inf(1)), 0x7c00},
		{"negative infinity", float32(math.Inf(-1)), 0xfc00},
		{"positive nan", math.Float32frombits(0x7fc12345), 0x7e00},
		{"negative nan", math.Float32frombits(0xffc12345), 0xfe00},
	}
	for _, test := range tests {
		if got := float32ToHalf(test.value); got != test.want {
			t.Fatalf("%s: float32ToHalf(%#x) = %#x, want %#x", test.name, math.Float32bits(test.value), got, test.want)
		}
	}
	if got := float32ToBFloat16(math.Float32frombits(0xffc12345)); got != 0x7fc0 {
		t.Fatalf("bfloat16 NaN = %#x, want 0x7fc0", got)
	}
}

func TestReduceFloat16NaNUsesJACCLCanonicalPayload(t *testing.T) {
	tests := []struct {
		name string
		dst  uint16
		src  uint16
		want uint16
	}{
		{"finite then nan", 0x3c00, 0x7e00, 0x7e00},
		{"positive then negative nan", 0x7e00, 0xfe00, 0x7e00},
		{"negative then positive nan", 0xfe00, 0x7e00, 0xfe00},
		{"signaling nan quieted", 0x7d00, 0x3c00, 0x7f00},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			dst := uint16Bytes(test.dst)
			if err := reduce(dst, uint16Bytes(test.src), Float16, Sum); err != nil {
				t.Fatal(err)
			}
			if got := binary.LittleEndian.Uint16(dst); got != test.want {
				t.Fatalf("float16 NaN sum = %#x, want %#x", got, test.want)
			}
		})
	}
}

func TestReduceBFloat16NaNUsesJACCLPayload(t *testing.T) {
	tests := []struct {
		name string
		dst  uint16
		src  uint16
		want uint16
	}{
		{"signaling nan quieted", 0x7f81, 0x3f80, 0x7fc1},
		{"signaling right wins", 0x7fc0, 0x7f81, 0x7fc1},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			dst := uint16Bytes(test.dst)
			if err := reduce(dst, uint16Bytes(test.src), BFloat16, Sum); err != nil {
				t.Fatal(err)
			}
			if got := binary.LittleEndian.Uint16(dst); got != test.want {
				t.Fatalf("bfloat16 NaN sum = %#x, want %#x", got, test.want)
			}
		})
	}
}

func TestReduceRankMajorUsesUpstreamOrder(t *testing.T) {
	gathered := float32Bytes(1e20, -1e20, 1)
	dst := make([]byte, 4)
	if err := reduceRankMajor(dst, gathered, 3, 4, Float32, Sum); err != nil {
		t.Fatal(err)
	}
	if got := math.Float32frombits(binary.LittleEndian.Uint32(dst)); got != 1 {
		t.Fatalf("rank-major sum = %g, want 1", got)
	}

	localFirst := float32Bytes(1)
	if err := reduce(localFirst, float32Bytes(1e20), Float32, Sum); err != nil {
		t.Fatal(err)
	}
	if err := reduce(localFirst, float32Bytes(-1e20), Float32, Sum); err != nil {
		t.Fatal(err)
	}
	if got := math.Float32frombits(binary.LittleEndian.Uint32(localFirst)); got != 0 {
		t.Fatalf("local-first control = %g, want 0", got)
	}
}

func int32Bytes(values ...int32) []byte {
	data := make([]byte, 4*len(values))
	for i, value := range values {
		binary.LittleEndian.PutUint32(data[4*i:], uint32(value))
	}
	return data
}

func int16Bytes(values ...int16) []byte {
	data := make([]byte, 2*len(values))
	for i, value := range values {
		binary.LittleEndian.PutUint16(data[2*i:], uint16(value))
	}
	return data
}

func uint16Bytes(values ...uint16) []byte {
	data := make([]byte, 2*len(values))
	for i, value := range values {
		binary.LittleEndian.PutUint16(data[2*i:], value)
	}
	return data
}

func int64Bytes(values ...int64) []byte {
	data := make([]byte, 8*len(values))
	for i, value := range values {
		binary.LittleEndian.PutUint64(data[8*i:], uint64(value))
	}
	return data
}

func uint32Bytes(values ...uint32) []byte {
	data := make([]byte, 4*len(values))
	for i, value := range values {
		binary.LittleEndian.PutUint32(data[4*i:], value)
	}
	return data
}

func uint64Bytes(values ...uint64) []byte {
	data := make([]byte, 8*len(values))
	for i, value := range values {
		binary.LittleEndian.PutUint64(data[8*i:], value)
	}
	return data
}

func float32Bytes(values ...float32) []byte {
	data := make([]byte, 4*len(values))
	for i, value := range values {
		binary.LittleEndian.PutUint32(data[4*i:], math.Float32bits(value))
	}
	return data
}

func float64Bytes(values ...float64) []byte {
	data := make([]byte, 8*len(values))
	for i, value := range values {
		binary.LittleEndian.PutUint64(data[8*i:], math.Float64bits(value))
	}
	return data
}

func complex64Bytes(values ...complex64) []byte {
	data := make([]byte, 8*len(values))
	for i, value := range values {
		binary.LittleEndian.PutUint32(data[8*i:], math.Float32bits(real(value)))
		binary.LittleEndian.PutUint32(data[8*i+4:], math.Float32bits(imag(value)))
	}
	return data
}

func complex64BitBytes(realBits, imagBits uint32) []byte {
	data := make([]byte, 8)
	binary.LittleEndian.PutUint32(data, realBits)
	binary.LittleEndian.PutUint32(data[4:], imagBits)
	return data
}
