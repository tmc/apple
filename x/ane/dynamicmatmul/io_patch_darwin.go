//go:build darwin

package dynamicmatmul

import (
	"fmt"
	"unsafe"

	appleiosurface "github.com/tmc/apple/iosurface"
)

func writeFullTileInput(tile *tile) error {
	return tile.k.WriteInputF32(0, tile.inputPacked)
}

func writeTileRows(tile *tile, rows []int) error {
	if len(rows) == 0 {
		return nil
	}
	layout := tile.k.InputLayout(0)
	if layout.Height != 1 || layout.ElemSize != 4 {
		return writeFullTileInput(tile)
	}
	surf := appleiosurface.IOSurfaceRef(tile.k.InputSurface(0))
	appleiosurface.IOSurfaceLock(surf, 0, nil)
	defer appleiosurface.IOSurfaceUnlock(surf, 0, nil)

	base := appleiosurface.IOSurfaceGetBaseAddress(surf)
	if base == nil {
		return fmt.Errorf("dynamic matmul: nil IOSurface base address")
	}
	rowBytes := layout.Width * layout.ElemSize
	if rowBytes <= 0 {
		return fmt.Errorf("dynamic matmul: invalid row bytes %d", rowBytes)
	}
	dst := unsafe.Slice((*byte)(base), layout.AllocSize())
	for _, row := range rows {
		if row < 0 || row >= layout.Channels {
			return fmt.Errorf("dynamic matmul: row %d out of range [0,%d)", row, layout.Channels)
		}
		off := row * layout.PlaneStride
		if off < 0 || off+rowBytes > len(dst) {
			return fmt.Errorf("dynamic matmul: row %d offset out of range", row)
		}
		srcRow := tile.inputPacked[row*layout.Width : (row+1)*layout.Width]
		src := unsafe.Slice((*byte)(unsafe.Pointer(&srcRow[0])), rowBytes)
		copy(dst[off:off+rowBytes], src)
	}
	return nil
}
