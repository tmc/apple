// Code generated from Apple documentation for gtshaderprofiler. DO NOT EDIT.

package gtshaderprofiler

import (
	"unsafe"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/coregraphics"
	"github.com/tmc/apple/objectivec"
)

// GTMioHeatmap protocol.
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioHeatmap
type GTMioHeatmap interface {
	objectivec.IObject

	// Depth protocol.
	//
	// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioHeatmap/depth
	Depth() uint64

	// EncoderInfo protocol.
	//
	// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioHeatmap/encoderInfo
	EncoderInfo() unsafe.Pointer

	// GenerateImage protocol.
	//
	// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioHeatmap/generateImage:
	GenerateImage(image uint64) *coregraphics.CGImageRef

	// GenerationOptions protocol.
	//
	// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioHeatmap/generationOptions
	GenerationOptions() unsafe.Pointer

	// HeatmapType protocol.
	//
	// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioHeatmap/heatmapType
	HeatmapType() uint64

	// Height protocol.
	//
	// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioHeatmap/height
	Height() uint64

	// MaxTimestamp protocol.
	//
	// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioHeatmap/maxTimestamp
	MaxTimestamp() uint64

	// MaxValue protocol.
	//
	// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioHeatmap/maxValue
	MaxValue() uint64

	// MinTimestamp protocol.
	//
	// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioHeatmap/minTimestamp
	MinTimestamp() uint64

	// MinValue protocol.
	//
	// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioHeatmap/minValue
	MinValue() uint64

	// NormalizedValueForPixelXPixelYSlice protocol.
	//
	// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioHeatmap/normalizedValueForPixelX:pixelY:slice:
	NormalizedValueForPixelXPixelYSlice(x uint64, y uint64, slice uint64) float64

	// NormalizedValueForValue protocol.
	//
	// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioHeatmap/normalizedValueForValue:
	NormalizedValueForValue(value uint64) float64

	// Options protocol.
	//
	// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioHeatmap/options
	Options() uint64

	// ProgramType protocol.
	//
	// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioHeatmap/programType
	ProgramType() uint16

	// QuadLocationForPixelXPixelYSlice protocol.
	//
	// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioHeatmap/quadLocationForPixelX:pixelY:slice:
	QuadLocationForPixelXPixelYSlice(x uint64, y uint64, slice uint64) uint64

	// ThreadPositionForPixelXPixelYSliceXYZ protocol.
	//
	// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioHeatmap/threadPositionForPixelX:pixelY:slice:x:y:z:
	ThreadPositionForPixelXPixelYSliceXYZ(x uint64, y uint64, slice uint64, x2 unsafe.Pointer, y2 unsafe.Pointer, z unsafe.Pointer) bool

	// ThreadRangeForPixelXPixelYSliceMinXMaxXMinYMaxYMinZMaxZ protocol.
	//
	// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioHeatmap/threadRangeForPixelX:pixelY:slice:minX:maxX:minY:maxY:minZ:maxZ:
	ThreadRangeForPixelXPixelYSliceMinXMaxXMinYMaxYMinZMaxZ(x uint64, y uint64, slice uint64, x2 unsafe.Pointer, x3 unsafe.Pointer, y2 unsafe.Pointer, y3 unsafe.Pointer, z unsafe.Pointer, z2 unsafe.Pointer) bool

	// Type protocol.
	//
	// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioHeatmap/type
	Type() uint64

	// ValueCount protocol.
	//
	// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioHeatmap/valueCount
	ValueCount() uint64

	// ValueForPixelXPixelYSlice protocol.
	//
	// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioHeatmap/valueForPixelX:pixelY:slice:
	ValueForPixelXPixelYSlice(x uint64, y uint64, slice uint64) uint64

	// Values protocol.
	//
	// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioHeatmap/values
	Values() unsafe.Pointer

	// Width protocol.
	//
	// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioHeatmap/width
	Width() uint64
}

// GTMioHeatmapObject wraps an existing Objective-C object that conforms to the GTMioHeatmap protocol.
type GTMioHeatmapObject struct {
	objectivec.Object
}
func (o GTMioHeatmapObject) BaseObject() objectivec.Object {
	return o.Object
}

// GTMioHeatmapObjectFromID constructs a [GTMioHeatmapObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func GTMioHeatmapObjectFromID(id objc.ID) GTMioHeatmapObject {
	return GTMioHeatmapObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioHeatmap/depth
func (o GTMioHeatmapObject) Depth() uint64 {
	rv := objc.Send[uint64](o.ID, objc.Sel("depth"))
	return rv
	}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioHeatmap/encoderInfo
func (o GTMioHeatmapObject) EncoderInfo() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o.ID, objc.Sel("encoderInfo"))
	return rv
	}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioHeatmap/generateImage:
func (o GTMioHeatmapObject) GenerateImage(image uint64) *coregraphics.CGImageRef {
	rv := objc.Send[unsafe.Pointer](o.ID, objc.Sel("generateImage:"), image)
		return (*coregraphics.CGImageRef)(rv)
	}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioHeatmap/generateTexture:
func (o GTMioHeatmapObject) GenerateTexture(texture uint64) objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("generateTexture:"), texture)
	return objectivec.Object{ID: rv}
	}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioHeatmap/generationOptions
func (o GTMioHeatmapObject) GenerationOptions() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o.ID, objc.Sel("generationOptions"))
	return rv
	}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioHeatmap/heatmapData
func (o GTMioHeatmapObject) HeatmapData() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("heatmapData"))
	return objectivec.Object{ID: rv}
	}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioHeatmap/heatmapType
func (o GTMioHeatmapObject) HeatmapType() uint64 {
	rv := objc.Send[uint64](o.ID, objc.Sel("heatmapType"))
	return rv
	}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioHeatmap/height
func (o GTMioHeatmapObject) Height() uint64 {
	rv := objc.Send[uint64](o.ID, objc.Sel("height"))
	return rv
	}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioHeatmap/maxTimestamp
func (o GTMioHeatmapObject) MaxTimestamp() uint64 {
	rv := objc.Send[uint64](o.ID, objc.Sel("maxTimestamp"))
	return rv
	}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioHeatmap/maxValue
func (o GTMioHeatmapObject) MaxValue() uint64 {
	rv := objc.Send[uint64](o.ID, objc.Sel("maxValue"))
	return rv
	}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioHeatmap/minTimestamp
func (o GTMioHeatmapObject) MinTimestamp() uint64 {
	rv := objc.Send[uint64](o.ID, objc.Sel("minTimestamp"))
	return rv
	}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioHeatmap/minValue
func (o GTMioHeatmapObject) MinValue() uint64 {
	rv := objc.Send[uint64](o.ID, objc.Sel("minValue"))
	return rv
	}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioHeatmap/normalizedValueForPixelX:pixelY:slice:
func (o GTMioHeatmapObject) NormalizedValueForPixelXPixelYSlice(x uint64, y uint64, slice uint64) float64 {
	rv := objc.Send[float64](o.ID, objc.Sel("normalizedValueForPixelX:pixelY:slice:"), x, y, slice)
	return rv
	}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioHeatmap/normalizedValueForValue:
func (o GTMioHeatmapObject) NormalizedValueForValue(value uint64) float64 {
	rv := objc.Send[float64](o.ID, objc.Sel("normalizedValueForValue:"), value)
	return rv
	}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioHeatmap/options
func (o GTMioHeatmapObject) Options() uint64 {
	rv := objc.Send[uint64](o.ID, objc.Sel("options"))
	return rv
	}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioHeatmap/programType
func (o GTMioHeatmapObject) ProgramType() uint16 {
	rv := objc.Send[uint16](o.ID, objc.Sel("programType"))
	return rv
	}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioHeatmap/quadLocationForPixelX:pixelY:slice:
func (o GTMioHeatmapObject) QuadLocationForPixelXPixelYSlice(x uint64, y uint64, slice uint64) uint64 {
	rv := objc.Send[uint64](o.ID, objc.Sel("quadLocationForPixelX:pixelY:slice:"), x, y, slice)
	return rv
	}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioHeatmap/threadPositionForPixelX:pixelY:slice:x:y:z:
func (o GTMioHeatmapObject) ThreadPositionForPixelXPixelYSliceXYZ(x uint64, y uint64, slice uint64, x2 unsafe.Pointer, y2 unsafe.Pointer, z unsafe.Pointer) bool {
	rv := objc.Send[bool](o.ID, objc.Sel("threadPositionForPixelX:pixelY:slice:x:y:z:"), x, y, slice, x2, y2, z)
	return rv
	}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioHeatmap/threadRangeForPixelX:pixelY:slice:minX:maxX:minY:maxY:minZ:maxZ:
func (o GTMioHeatmapObject) ThreadRangeForPixelXPixelYSliceMinXMaxXMinYMaxYMinZMaxZ(x uint64, y uint64, slice uint64, x2 unsafe.Pointer, x3 unsafe.Pointer, y2 unsafe.Pointer, y3 unsafe.Pointer, z unsafe.Pointer, z2 unsafe.Pointer) bool {
	rv := objc.Send[bool](o.ID, objc.Sel("threadRangeForPixelX:pixelY:slice:minX:maxX:minY:maxY:minZ:maxZ:"), x, y, slice, x2, x3, y2, y3, z, z2)
	return rv
	}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioHeatmap/type
func (o GTMioHeatmapObject) Type() uint64 {
	rv := objc.Send[uint64](o.ID, objc.Sel("type"))
	return rv
	}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioHeatmap/valueCount
func (o GTMioHeatmapObject) ValueCount() uint64 {
	rv := objc.Send[uint64](o.ID, objc.Sel("valueCount"))
	return rv
	}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioHeatmap/valueForPixelX:pixelY:slice:
func (o GTMioHeatmapObject) ValueForPixelXPixelYSlice(x uint64, y uint64, slice uint64) uint64 {
	rv := objc.Send[uint64](o.ID, objc.Sel("valueForPixelX:pixelY:slice:"), x, y, slice)
	return rv
	}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioHeatmap/values
func (o GTMioHeatmapObject) Values() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o.ID, objc.Sel("values"))
	return rv
	}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioHeatmap/width
func (o GTMioHeatmapObject) Width() uint64 {
	rv := objc.Send[uint64](o.ID, objc.Sel("width"))
	return rv
	}

