//go:build darwin

package ane

import "unsafe"

func logicalElementLimit(dataLen int, layout TensorLayout) int {
	n := layout.LogicalElements()
	if dataLen < n {
		return dataLen
	}
	return n
}

func writeStridedF32Bytes(dst []byte, data []float32, layout TensorLayout) {
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
			if off+n*4 > len(dst) {
				return
			}
			row := unsafe.Slice((*float32)(unsafe.Pointer(&dst[off])), n)
			copy(row, data[srcIdx:srcIdx+n])
		}
	}
}

func readStridedF32Bytes(data []float32, src []byte, layout TensorLayout) {
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
			if off+n*4 > len(src) {
				return
			}
			row := unsafe.Slice((*float32)(unsafe.Pointer(&src[off])), n)
			copy(data[dstIdx:dstIdx+n], row)
		}
	}
}
