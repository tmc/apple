// Code generated from Apple documentation for gtshaderprofiler. DO NOT EDIT.

package gtshaderprofiler

import (
	"unsafe"

	"github.com/tmc/apple/coregraphics"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// GTMioHeatmap protocol.
type GTMioHeatmap interface {
	objectivec.IObject

	// Depth protocol.
	Depth() uint64

	// EncoderInfo protocol.
	EncoderInfo() unsafe.Pointer

	// GenerateImage protocol.
	GenerateImage(image uint64) coregraphics.CGImageRef

	// GenerationOptions protocol.
	GenerationOptions() unsafe.Pointer

	// HeatmapType protocol.
	HeatmapType() uint64

	// Height protocol.
	Height() uint64

	// MaxTimestamp protocol.
	MaxTimestamp() uint64

	// MaxValue protocol.
	MaxValue() uint64

	// MinTimestamp protocol.
	MinTimestamp() uint64

	// MinValue protocol.
	MinValue() uint64

	// NormalizedValueForPixelXPixelYSlice protocol.
	NormalizedValueForPixelXPixelYSlice(x uint64, y uint64, slice uint64) float64

	// NormalizedValueForValue protocol.
	NormalizedValueForValue(value uint64) float64

	// Options protocol.
	Options() uint64

	// ProgramType protocol.
	ProgramType() uint16

	// QuadLocationForPixelXPixelYSlice protocol.
	QuadLocationForPixelXPixelYSlice(x uint64, y uint64, slice uint64) uint64

	// ThreadPositionForPixelXPixelYSliceXYZ protocol.
	ThreadPositionForPixelXPixelYSliceXYZ(x uint64, y uint64, slice uint64, x2 *uint32, y2 *uint32, z *uint32) bool

	// ThreadRangeForPixelXPixelYSliceMinXMaxXMinYMaxYMinZMaxZ protocol.
	ThreadRangeForPixelXPixelYSliceMinXMaxXMinYMaxYMinZMaxZ(x uint64, y uint64, slice uint64, x2 *uint32, x3 *uint32, y2 *uint32, y3 *uint32, z *uint32, z2 *uint32) bool

	// Type protocol.
	Type() uint64

	// ValueCount protocol.
	ValueCount() uint64

	// ValueForPixelXPixelYSlice protocol.
	ValueForPixelXPixelYSlice(x uint64, y uint64, slice uint64) uint64

	// Values protocol.
	Values() unsafe.Pointer

	// Width protocol.
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

func (o GTMioHeatmapObject) Depth() uint64 {
	rv := objc.Send[uint64](o.ID, objc.Sel("depth"))
	return rv
}
func (o GTMioHeatmapObject) EncoderInfo() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o.ID, objc.Sel("encoderInfo"))
	return rv
}
func (o GTMioHeatmapObject) GenerateImage(image uint64) coregraphics.CGImageRef {
	rv := objc.Send[coregraphics.CGImageRef](o.ID, objc.Sel("generateImage:"), image)
	return rv
}
func (o GTMioHeatmapObject) GenerateTexture(texture uint64) objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("generateTexture:"), texture)
	return objectivec.Object{ID: rv}
}
func (o GTMioHeatmapObject) GenerationOptions() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o.ID, objc.Sel("generationOptions"))
	return rv
}
func (o GTMioHeatmapObject) HeatmapData() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("heatmapData"))
	return objectivec.Object{ID: rv}
}
func (o GTMioHeatmapObject) HeatmapType() uint64 {
	rv := objc.Send[uint64](o.ID, objc.Sel("heatmapType"))
	return rv
}
func (o GTMioHeatmapObject) Height() uint64 {
	rv := objc.Send[uint64](o.ID, objc.Sel("height"))
	return rv
}
func (o GTMioHeatmapObject) MaxTimestamp() uint64 {
	rv := objc.Send[uint64](o.ID, objc.Sel("maxTimestamp"))
	return rv
}
func (o GTMioHeatmapObject) MaxValue() uint64 {
	rv := objc.Send[uint64](o.ID, objc.Sel("maxValue"))
	return rv
}
func (o GTMioHeatmapObject) MinTimestamp() uint64 {
	rv := objc.Send[uint64](o.ID, objc.Sel("minTimestamp"))
	return rv
}
func (o GTMioHeatmapObject) MinValue() uint64 {
	rv := objc.Send[uint64](o.ID, objc.Sel("minValue"))
	return rv
}
func (o GTMioHeatmapObject) NormalizedValueForPixelXPixelYSlice(x uint64, y uint64, slice uint64) float64 {
	rv := objc.Send[float64](o.ID, objc.Sel("normalizedValueForPixelX:pixelY:slice:"), x, y, slice)
	return rv
}
func (o GTMioHeatmapObject) NormalizedValueForValue(value uint64) float64 {
	rv := objc.Send[float64](o.ID, objc.Sel("normalizedValueForValue:"), value)
	return rv
}
func (o GTMioHeatmapObject) Options() uint64 {
	rv := objc.Send[uint64](o.ID, objc.Sel("options"))
	return rv
}
func (o GTMioHeatmapObject) ProgramType() uint16 {
	rv := objc.Send[uint16](o.ID, objc.Sel("programType"))
	return rv
}
func (o GTMioHeatmapObject) QuadLocationForPixelXPixelYSlice(x uint64, y uint64, slice uint64) uint64 {
	rv := objc.Send[uint64](o.ID, objc.Sel("quadLocationForPixelX:pixelY:slice:"), x, y, slice)
	return rv
}
func (o GTMioHeatmapObject) ThreadPositionForPixelXPixelYSliceXYZ(x uint64, y uint64, slice uint64, x2 *uint32, y2 *uint32, z *uint32) bool {
	rv := objc.Send[bool](o.ID, objc.Sel("threadPositionForPixelX:pixelY:slice:x:y:z:"), x, y, slice, x2, y2, z)
	return rv
}
func (o GTMioHeatmapObject) ThreadRangeForPixelXPixelYSliceMinXMaxXMinYMaxYMinZMaxZ(x uint64, y uint64, slice uint64, x2 *uint32, x3 *uint32, y2 *uint32, y3 *uint32, z *uint32, z2 *uint32) bool {
	rv := objc.Send[bool](o.ID, objc.Sel("threadRangeForPixelX:pixelY:slice:minX:maxX:minY:maxY:minZ:maxZ:"), x, y, slice, x2, x3, y2, y3, z, z2)
	return rv
}
func (o GTMioHeatmapObject) Type() uint64 {
	rv := objc.Send[uint64](o.ID, objc.Sel("type"))
	return rv
}
func (o GTMioHeatmapObject) ValueCount() uint64 {
	rv := objc.Send[uint64](o.ID, objc.Sel("valueCount"))
	return rv
}
func (o GTMioHeatmapObject) ValueForPixelXPixelYSlice(x uint64, y uint64, slice uint64) uint64 {
	rv := objc.Send[uint64](o.ID, objc.Sel("valueForPixelX:pixelY:slice:"), x, y, slice)
	return rv
}
func (o GTMioHeatmapObject) Values() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o.ID, objc.Sel("values"))
	return rv
}
func (o GTMioHeatmapObject) Width() uint64 {
	rv := objc.Send[uint64](o.ID, objc.Sel("width"))
	return rv
}
