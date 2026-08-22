//go:build darwin && arm64

package ane

import (
	"encoding/binary"
	"math"
	"testing"
	"unsafe"
)

// benchWeights is a realistic ANE weight-array size.
const benchWeights = 1 << 20 // 1,048,576 float32

func benchFloat32Data() []float32 {
	data := make([]float32, benchWeights)
	rng := splitmix64(0xB00F_1234_5678_9ABC)
	for i := range data {
		r := rng.next()
		exp := uint32(112+(r>>52)%30) << 23
		data[i] = math.Float32frombits(uint32(r)&0x807FFFFF | exp)
	}
	return data
}

// BenchmarkFP16ConvertWrite compares the per-element scalar helper used by the
// seven bulk weight-serialization call sites against one contiguous NEON call.
//
// NOTE: the two arms do NOT compute the same result — see
// TestFloat32ToFP16Differential. This measures cost, not equivalence.
func BenchmarkFP16ConvertWrite(b *testing.B) {
	data := benchFloat32Data()
	dst := make([]byte, 2*len(data))

	b.Run("scalar/PutUint16", func(b *testing.B) {
		b.SetBytes(int64(len(dst)))
		for range b.N {
			for i, w := range data {
				binary.LittleEndian.PutUint16(dst[i*2:], float32ToFP16(w))
			}
		}
	})

	b.Run("asm/contiguous", func(b *testing.B) {
		b.SetBytes(int64(len(dst)))
		for range b.N {
			cvtF32ToF16(unsafe.SliceData(dst), unsafe.SliceData(data), len(data))
		}
	})
}

// BenchmarkFP16ConvertRead is the read-side counterpart.
func BenchmarkFP16ConvertRead(b *testing.B) {
	data := benchFloat32Data()
	src := make([]byte, 2*len(data))
	cvtF32ToF16(unsafe.SliceData(src), unsafe.SliceData(data), len(data))
	out := make([]float32, len(data))

	b.Run("scalar/Uint16", func(b *testing.B) {
		b.SetBytes(int64(len(src)))
		for range b.N {
			for i := range out {
				out[i] = fp16ToFloat32(binary.LittleEndian.Uint16(src[i*2:]))
			}
		}
	})

	b.Run("asm/contiguous", func(b *testing.B) {
		b.SetBytes(int64(len(src)))
		for range b.N {
			cvtF16ToF32(unsafe.SliceData(out), unsafe.SliceData(src), len(out))
		}
	})
}
