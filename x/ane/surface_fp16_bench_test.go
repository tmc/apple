//go:build darwin

package ane

import "testing"

func BenchmarkWriteStridedFP16(b *testing.B) {
	for _, bc := range []struct {
		name    string
		layout  TensorLayout
		dataLen int
	}{
		{
			name:    "small/2x5x2",
			layout:  TensorLayout{Channels: 2, Width: 5, Height: 2, ElemSize: 2, RowStride: 64, PlaneStride: 128},
			dataLen: 20,
		},
		{
			name:    "medium/64x1x64",
			layout:  TensorLayout{Channels: 64, Width: 64, Height: 1, ElemSize: 2, RowStride: 128, PlaneStride: 128},
			dataLen: 64 * 64,
		},
		{
			name:    "large/256x1x256",
			layout:  TensorLayout{Channels: 256, Width: 256, Height: 1, ElemSize: 2, RowStride: 512, PlaneStride: 512},
			dataLen: 256 * 256,
		},
	} {
		data := make([]float32, bc.dataLen)
		for i := range data {
			data[i] = float32(i) * 0.1
		}
		dst := make([]byte, bc.layout.AllocSize())

		b.Run(bc.name+"/write", func(b *testing.B) {
			for range b.N {
				writeStridedFP16Bytes(dst, data, bc.layout)
			}
		})

		// Pre-fill dst for read benchmark.
		writeStridedFP16Bytes(dst, data, bc.layout)
		out := make([]float32, bc.dataLen)

		b.Run(bc.name+"/read", func(b *testing.B) {
			for range b.N {
				readStridedFP16Bytes(out, dst, bc.layout)
			}
		})
	}
}
