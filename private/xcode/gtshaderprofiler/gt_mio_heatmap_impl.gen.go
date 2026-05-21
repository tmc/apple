// Code generated from Apple documentation for gtshaderprofiler. DO NOT EDIT.

package gtshaderprofiler

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/coregraphics"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [GTMioHeatmapImpl] class.
var (
	_GTMioHeatmapImplClass     GTMioHeatmapImplClass
	_GTMioHeatmapImplClassOnce sync.Once
)

func getGTMioHeatmapImplClass() GTMioHeatmapImplClass {
	_GTMioHeatmapImplClassOnce.Do(func() {
		_GTMioHeatmapImplClass = GTMioHeatmapImplClass{class: objc.GetClass("GTMioHeatmapImpl")}
	})
	return _GTMioHeatmapImplClass
}

// GetGTMioHeatmapImplClass returns the class object for GTMioHeatmapImpl.
func GetGTMioHeatmapImplClass() GTMioHeatmapImplClass {
	return getGTMioHeatmapImplClass()
}

type GTMioHeatmapImplClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (gc GTMioHeatmapImplClass) Class() objc.Class {
	return gc.class
}

// Alloc allocates memory for a new instance of the class.
func (gc GTMioHeatmapImplClass) Alloc() GTMioHeatmapImpl {
	rv := objc.Send[GTMioHeatmapImpl](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [GTMioHeatmapImpl.Depth]
//   - [GTMioHeatmapImpl.EncoderInfo]
//   - [GTMioHeatmapImpl.SetEncoderInfo]
//   - [GTMioHeatmapImpl.GenerateImage]
//   - [GTMioHeatmapImpl.GenerateTexture]
//   - [GTMioHeatmapImpl.GenerationOptions]
//   - [GTMioHeatmapImpl.HeatmapData]
//   - [GTMioHeatmapImpl.SetHeatmapData]
//   - [GTMioHeatmapImpl.HeatmapType]
//   - [GTMioHeatmapImpl.Height]
//   - [GTMioHeatmapImpl.MaxTimestamp]
//   - [GTMioHeatmapImpl.MaxValue]
//   - [GTMioHeatmapImpl.MinTimestamp]
//   - [GTMioHeatmapImpl.MinValue]
//   - [GTMioHeatmapImpl.NormalizedValueForPixelXPixelYSlice]
//   - [GTMioHeatmapImpl.NormalizedValueForValue]
//   - [GTMioHeatmapImpl.Options]
//   - [GTMioHeatmapImpl.ProgramType]
//   - [GTMioHeatmapImpl.QuadLocationForPixelXPixelYSlice]
//   - [GTMioHeatmapImpl.SetPixelToQuadLocMap]
//   - [GTMioHeatmapImpl.ThreadPositionForPixelXPixelYSliceXYZ]
//   - [GTMioHeatmapImpl.ThreadRangeForPixelXPixelYSliceMinXMaxXMinYMaxYMinZMaxZ]
//   - [GTMioHeatmapImpl.Type]
//   - [GTMioHeatmapImpl.ValueCount]
//   - [GTMioHeatmapImpl.ValueForPixelXPixelYSlice]
//   - [GTMioHeatmapImpl.Values]
//   - [GTMioHeatmapImpl.Width]
//   - [GTMioHeatmapImpl.InitWithEncoderInfoWitdhHeightDepthQuadDataType]
type GTMioHeatmapImpl struct {
	objectivec.Object
}

// GTMioHeatmapImplFromID constructs a [GTMioHeatmapImpl] from an objc.ID.
func GTMioHeatmapImplFromID(id objc.ID) GTMioHeatmapImpl {
	return GTMioHeatmapImpl{objectivec.Object{ID: id}}
}

// Ensure GTMioHeatmapImpl implements IGTMioHeatmapImpl.
var _ IGTMioHeatmapImpl = GTMioHeatmapImpl{}

// An interface definition for the [GTMioHeatmapImpl] class.
//
// # Methods
//
//   - [IGTMioHeatmapImpl.Depth]
//   - [IGTMioHeatmapImpl.EncoderInfo]
//   - [IGTMioHeatmapImpl.SetEncoderInfo]
//   - [IGTMioHeatmapImpl.GenerateImage]
//   - [IGTMioHeatmapImpl.GenerateTexture]
//   - [IGTMioHeatmapImpl.GenerationOptions]
//   - [IGTMioHeatmapImpl.HeatmapData]
//   - [IGTMioHeatmapImpl.SetHeatmapData]
//   - [IGTMioHeatmapImpl.HeatmapType]
//   - [IGTMioHeatmapImpl.Height]
//   - [IGTMioHeatmapImpl.MaxTimestamp]
//   - [IGTMioHeatmapImpl.MaxValue]
//   - [IGTMioHeatmapImpl.MinTimestamp]
//   - [IGTMioHeatmapImpl.MinValue]
//   - [IGTMioHeatmapImpl.NormalizedValueForPixelXPixelYSlice]
//   - [IGTMioHeatmapImpl.NormalizedValueForValue]
//   - [IGTMioHeatmapImpl.Options]
//   - [IGTMioHeatmapImpl.ProgramType]
//   - [IGTMioHeatmapImpl.QuadLocationForPixelXPixelYSlice]
//   - [IGTMioHeatmapImpl.SetPixelToQuadLocMap]
//   - [IGTMioHeatmapImpl.ThreadPositionForPixelXPixelYSliceXYZ]
//   - [IGTMioHeatmapImpl.ThreadRangeForPixelXPixelYSliceMinXMaxXMinYMaxYMinZMaxZ]
//   - [IGTMioHeatmapImpl.Type]
//   - [IGTMioHeatmapImpl.ValueCount]
//   - [IGTMioHeatmapImpl.ValueForPixelXPixelYSlice]
//   - [IGTMioHeatmapImpl.Values]
//   - [IGTMioHeatmapImpl.Width]
//   - [IGTMioHeatmapImpl.InitWithEncoderInfoWitdhHeightDepthQuadDataType]
type IGTMioHeatmapImpl interface {
	objectivec.IObject

	// Topic: Methods

	Depth() uint64
	EncoderInfo() unsafe.Pointer
	SetEncoderInfo(value *GTMioEncoderMetadata)
	GenerateImage(image uint64) coregraphics.CGImageRef
	GenerateTexture(texture uint64) objectivec.IObject
	GenerationOptions() unsafe.Pointer
	HeatmapData() foundation.NSData
	SetHeatmapData(value foundation.NSData)
	HeatmapType() uint64
	Height() uint64
	MaxTimestamp() uint64
	MaxValue() uint64
	MinTimestamp() uint64
	MinValue() uint64
	NormalizedValueForPixelXPixelYSlice(x uint64, y uint64, slice uint64) float64
	NormalizedValueForValue(value uint64) float64
	Options() uint64
	ProgramType() uint16
	QuadLocationForPixelXPixelYSlice(x uint64, y uint64, slice uint64) uint64
	SetPixelToQuadLocMap(map_ unsafe.Pointer)
	ThreadPositionForPixelXPixelYSliceXYZ(x uint64, y uint64, slice uint64, x2 *uint32, y2 *uint32, z *uint32) bool
	ThreadRangeForPixelXPixelYSliceMinXMaxXMinYMaxYMinZMaxZ(x uint64, y uint64, slice uint64, x2 *uint32, x3 *uint32, y2 *uint32, y3 *uint32, z *uint32, z2 *uint32) bool
	Type() uint64
	ValueCount() uint64
	ValueForPixelXPixelYSlice(x uint64, y uint64, slice uint64) uint64
	Values() unsafe.Pointer
	Width() uint64
	InitWithEncoderInfoWitdhHeightDepthQuadDataType(info GTMioEncoderMetadata, witdh uint64, height uint64, depth uint64, data objectivec.IObject, type_ uint64) GTMioHeatmapImpl
}

// Init initializes the instance.
func (g GTMioHeatmapImpl) Init() GTMioHeatmapImpl {
	rv := objc.Send[GTMioHeatmapImpl](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g GTMioHeatmapImpl) Autorelease() GTMioHeatmapImpl {
	rv := objc.Send[GTMioHeatmapImpl](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewGTMioHeatmapImpl creates a new GTMioHeatmapImpl instance.
func NewGTMioHeatmapImpl() GTMioHeatmapImpl {
	class := getGTMioHeatmapImplClass()
	rv := objc.Send[GTMioHeatmapImpl](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewGTMioHeatmapImplWithEncoderInfoWitdhHeightDepthQuadDataType(info GTMioEncoderMetadata, witdh uint64, height uint64, depth uint64, data objectivec.IObject, type_ uint64) GTMioHeatmapImpl {
	instance := getGTMioHeatmapImplClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithEncoderInfo:witdh:height:depth:quadData:type:"), info, witdh, height, depth, data, type_)
	return GTMioHeatmapImplFromID(rv)
}

func (g GTMioHeatmapImpl) GenerateImage(image uint64) coregraphics.CGImageRef {
	rv := objc.Send[coregraphics.CGImageRef](g.ID, objc.Sel("generateImage:"), image)
	return coregraphics.CGImageRef(rv)
}
func (g GTMioHeatmapImpl) GenerateTexture(texture uint64) objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("generateTexture:"), texture)
	return objectivec.Object{ID: rv}
}
func (g GTMioHeatmapImpl) NormalizedValueForPixelXPixelYSlice(x uint64, y uint64, slice uint64) float64 {
	rv := objc.Send[float64](g.ID, objc.Sel("normalizedValueForPixelX:pixelY:slice:"), x, y, slice)
	return rv
}
func (g GTMioHeatmapImpl) NormalizedValueForValue(value uint64) float64 {
	rv := objc.Send[float64](g.ID, objc.Sel("normalizedValueForValue:"), value)
	return rv
}
func (g GTMioHeatmapImpl) QuadLocationForPixelXPixelYSlice(x uint64, y uint64, slice uint64) uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("quadLocationForPixelX:pixelY:slice:"), x, y, slice)
	return rv
}
func (g GTMioHeatmapImpl) SetPixelToQuadLocMap(map_ unsafe.Pointer) {
	objc.Send[objc.ID](g.ID, objc.Sel("setPixelToQuadLocMap:"), map_)
}
func (g GTMioHeatmapImpl) ThreadPositionForPixelXPixelYSliceXYZ(x uint64, y uint64, slice uint64, x2 *uint32, y2 *uint32, z *uint32) bool {
	rv := objc.Send[bool](g.ID, objc.Sel("threadPositionForPixelX:pixelY:slice:x:y:z:"), x, y, slice, x2, y2, z)
	return rv
}
func (g GTMioHeatmapImpl) ThreadRangeForPixelXPixelYSliceMinXMaxXMinYMaxYMinZMaxZ(x uint64, y uint64, slice uint64, x2 *uint32, x3 *uint32, y2 *uint32, y3 *uint32, z *uint32, z2 *uint32) bool {
	rv := objc.Send[bool](g.ID, objc.Sel("threadRangeForPixelX:pixelY:slice:minX:maxX:minY:maxY:minZ:maxZ:"), x, y, slice, x2, x3, y2, y3, z, z2)
	return rv
}
func (g GTMioHeatmapImpl) ValueForPixelXPixelYSlice(x uint64, y uint64, slice uint64) uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("valueForPixelX:pixelY:slice:"), x, y, slice)
	return rv
}
func (g GTMioHeatmapImpl) InitWithEncoderInfoWitdhHeightDepthQuadDataType(info GTMioEncoderMetadata, witdh uint64, height uint64, depth uint64, data objectivec.IObject, type_ uint64) GTMioHeatmapImpl {
	rv := objc.Send[GTMioHeatmapImpl](g.ID, objc.Sel("initWithEncoderInfo:witdh:height:depth:quadData:type:"), info, witdh, height, depth, data, type_)
	return rv
}

func (g GTMioHeatmapImpl) Depth() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("depth"))
	return rv
}
func (g GTMioHeatmapImpl) EncoderInfo() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g.ID, objc.Sel("encoderInfo"))
	return rv
}
func (g GTMioHeatmapImpl) SetEncoderInfo(value *GTMioEncoderMetadata) {
	objc.Send[struct{}](g.ID, objc.Sel("setEncoderInfo:"), value)
}
func (g GTMioHeatmapImpl) GenerationOptions() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g.ID, objc.Sel("generationOptions"))
	return rv
}
func (g GTMioHeatmapImpl) HeatmapData() foundation.NSData {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("heatmapData"))
	return foundation.NSDataFromID(objc.ID(rv))
}
func (g GTMioHeatmapImpl) SetHeatmapData(value foundation.NSData) {
	objc.Send[struct{}](g.ID, objc.Sel("setHeatmapData:"), value)
}
func (g GTMioHeatmapImpl) HeatmapType() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("heatmapType"))
	return rv
}
func (g GTMioHeatmapImpl) Height() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("height"))
	return rv
}
func (g GTMioHeatmapImpl) MaxTimestamp() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("maxTimestamp"))
	return rv
}
func (g GTMioHeatmapImpl) MaxValue() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("maxValue"))
	return rv
}
func (g GTMioHeatmapImpl) MinTimestamp() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("minTimestamp"))
	return rv
}
func (g GTMioHeatmapImpl) MinValue() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("minValue"))
	return rv
}
func (g GTMioHeatmapImpl) Options() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("options"))
	return rv
}
func (g GTMioHeatmapImpl) ProgramType() uint16 {
	rv := objc.Send[uint16](g.ID, objc.Sel("programType"))
	return rv
}
func (g GTMioHeatmapImpl) Type() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("type"))
	return rv
}
func (g GTMioHeatmapImpl) ValueCount() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("valueCount"))
	return rv
}
func (g GTMioHeatmapImpl) Values() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g.ID, objc.Sel("values"))
	return rv
}
func (g GTMioHeatmapImpl) Width() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("width"))
	return rv
}
