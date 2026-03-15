//go:build darwin && arm64

package ane

import "unsafe"

// cvtF32ToF16 converts n float32 values to fp16 using NEON FCVTN.
//
//go:noescape
func cvtF32ToF16(dst *byte, src *float32, n int)

// cvtF16ToF32 converts n fp16 values to float32 using NEON FCVTL.
//
//go:noescape
func cvtF16ToF32(dst *float32, src *byte, n int)

func writeStridedFP16Bytes(dst []byte, data []float32, layout TensorLayout) {
	if layout.Width <= 0 || layout.Height <= 0 || layout.Channels <= 0 {
		return
	}
	limit := logicalElementLimit(len(data), layout)
	if limit == 0 {
		return
	}
	hw := layout.Height * layout.Width
	for c := range layout.Channels {
		for h := range layout.Height {
			srcIdx := c*hw + h*layout.Width
			if srcIdx >= limit {
				return
			}
			n := layout.Width
			if remain := limit - srcIdx; remain < n {
				n = remain
			}
			off := c*layout.PlaneStride + h*layout.RowStride
			if off+n*2 > len(dst) {
				return
			}
			cvtF32ToF16(
				unsafe.SliceData(dst[off:]),
				unsafe.SliceData(data[srcIdx:]),
				n,
			)
		}
	}
}

func readStridedFP16Bytes(data []float32, src []byte, layout TensorLayout) {
	if layout.Width <= 0 || layout.Height <= 0 || layout.Channels <= 0 {
		return
	}
	limit := logicalElementLimit(len(data), layout)
	if limit == 0 {
		return
	}
	hw := layout.Height * layout.Width
	for c := range layout.Channels {
		for h := range layout.Height {
			dstIdx := c*hw + h*layout.Width
			if dstIdx >= limit {
				return
			}
			n := layout.Width
			if remain := limit - dstIdx; remain < n {
				n = remain
			}
			off := c*layout.PlaneStride + h*layout.RowStride
			if off+n*2 > len(src) {
				return
			}
			cvtF16ToF32(
				unsafe.SliceData(data[dstIdx:]),
				unsafe.SliceData(src[off:]),
				n,
			)
		}
	}
}
