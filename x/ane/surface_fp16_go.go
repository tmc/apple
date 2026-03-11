//go:build darwin && (!arm64 || !cgo)

package ane

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
			row := dst[off : off+n*2]
			for i, v := range data[srcIdx : srcIdx+n] {
				bits := float32ToFP16(v)
				row[i*2] = byte(bits)
				row[i*2+1] = byte(bits >> 8)
			}
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
			row := src[off : off+n*2]
			for i := range n {
				bits := uint16(row[i*2]) | uint16(row[i*2+1])<<8
				data[dstIdx+i] = fp16ToFloat32(bits)
			}
		}
	}
}
