// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import "unsafe"

// toSameSize converts between int and uint (same underlying size).
func toSameSize[From, To int | uint](v From) To {
	return *(*To)(unsafe.Pointer(&v))
}

// Inline constructor helpers for Metal struct types.

func mtl4BufferRangeMake(bufferAddress MTLGPUAddress, length uint64) MTL4BufferRange {
	return MTL4BufferRange{BufferAddress: bufferAddress, Length: length}
}

func mtlClearColorMake(red, green, blue, alpha float64) MTLClearColor {
	return MTLClearColor{Red: red, Green: green, Blue: blue, Alpha: alpha}
}

func mtlCoordinate2DMake(x, y float32) MTLCoordinate2D {
	return MTLCoordinate2D{X: x, Y: y}
}

func mtlIndirectCommandBufferExecutionRangeMake(location, length uint32) MTLIndirectCommandBufferExecutionRange {
	return MTLIndirectCommandBufferExecutionRange{Location: location, Length: length}
}

func mtlOriginMake(x, y, z uint) MTLOrigin {
	type origin struct{ x, y, z uint }
	o := origin{x, y, z}
	return *(*MTLOrigin)(unsafe.Pointer(&o))
}

func mtlPackedFloat3Make(x, y, z float32) MTLPackedFloat3 {
	return MTLPackedFloat3{X: x, Y: y, Z: z}
}

func mtlPackedFloatQuaternionMake(x, y, z, w float32) MTLPackedFloatQuaternion {
	return MTLPackedFloatQuaternion{X: x, Y: y, Z: z, W: w}
}

func mtlRegionMake1D(x, width uint) MTLRegion {
	return MTLRegion{
		Origin: mtlOriginMake(x, 0, 0),
		Size:   mtlSizeMake(width, 1, 1),
	}
}

func mtlRegionMake2D(x, y, width, height uint) MTLRegion {
	return MTLRegion{
		Origin: mtlOriginMake(x, y, 0),
		Size:   mtlSizeMake(width, height, 1),
	}
}

func mtlRegionMake3D(x, y, z, width, height, depth uint) MTLRegion {
	return MTLRegion{
		Origin: mtlOriginMake(x, y, z),
		Size:   mtlSizeMake(width, height, depth),
	}
}

func mtlSamplePositionMake(x, y float32) MTLSamplePosition {
	return MTLSamplePosition{X: x, Y: y}
}

func mtlSizeMake(width, height, depth uint) MTLSize {
	type size struct{ w, h, d uint }
	s := size{width, height, depth}
	return *(*MTLSize)(unsafe.Pointer(&s))
}

func mtlTextureSwizzleChannelsMake(r, g, b, a MTLTextureSwizzle) MTLTextureSwizzleChannels {
	return MTLTextureSwizzleChannels{Red: r, Green: g, Blue: b, Alpha: a}
}

