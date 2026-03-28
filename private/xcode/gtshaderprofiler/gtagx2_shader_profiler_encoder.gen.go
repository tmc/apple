// Code generated from Apple documentation for gtshaderprofiler. DO NOT EDIT.

package gtshaderprofiler

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [GTAGX2ShaderProfilerEncoder] class.
var (
	_GTAGX2ShaderProfilerEncoderClass     GTAGX2ShaderProfilerEncoderClass
	_GTAGX2ShaderProfilerEncoderClassOnce sync.Once
)

func getGTAGX2ShaderProfilerEncoderClass() GTAGX2ShaderProfilerEncoderClass {
	_GTAGX2ShaderProfilerEncoderClassOnce.Do(func() {
		_GTAGX2ShaderProfilerEncoderClass = GTAGX2ShaderProfilerEncoderClass{class: objc.GetClass("GTAGX2ShaderProfilerEncoder")}
	})
	return _GTAGX2ShaderProfilerEncoderClass
}

// GetGTAGX2ShaderProfilerEncoderClass returns the class object for GTAGX2ShaderProfilerEncoder.
func GetGTAGX2ShaderProfilerEncoderClass() GTAGX2ShaderProfilerEncoderClass {
	return getGTAGX2ShaderProfilerEncoderClass()
}

type GTAGX2ShaderProfilerEncoderClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (gc GTAGX2ShaderProfilerEncoderClass) Class() objc.Class {
	return gc.class
}

// Alloc allocates memory for a new instance of the class.
func (gc GTAGX2ShaderProfilerEncoderClass) Alloc() GTAGX2ShaderProfilerEncoder {
	rv := objc.Send[GTAGX2ShaderProfilerEncoder](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [GTAGX2ShaderProfilerEncoder.EncodeWithCoder]
//   - [GTAGX2ShaderProfilerEncoder.FunctionIndex]
//   - [GTAGX2ShaderProfilerEncoder.SetFunctionIndex]
//   - [GTAGX2ShaderProfilerEncoder.GpuCommandStartIndex]
//   - [GTAGX2ShaderProfilerEncoder.SetGpuCommandStartIndex]
//   - [GTAGX2ShaderProfilerEncoder.LoadTime]
//   - [GTAGX2ShaderProfilerEncoder.SetLoadTime]
//   - [GTAGX2ShaderProfilerEncoder.NumGPUCommands]
//   - [GTAGX2ShaderProfilerEncoder.SetNumGPUCommands]
//   - [GTAGX2ShaderProfilerEncoder.ObjectId]
//   - [GTAGX2ShaderProfilerEncoder.SetObjectId]
//   - [GTAGX2ShaderProfilerEncoder.PointerId]
//   - [GTAGX2ShaderProfilerEncoder.SetPointerId]
//   - [GTAGX2ShaderProfilerEncoder.SetIndex]
//   - [GTAGX2ShaderProfilerEncoder.StoreTime]
//   - [GTAGX2ShaderProfilerEncoder.SetStoreTime]
//   - [GTAGX2ShaderProfilerEncoder.TimingInfo]
//   - [GTAGX2ShaderProfilerEncoder.SetTimingInfo]
//   - [GTAGX2ShaderProfilerEncoder.InitWithCoder]
//   - [GTAGX2ShaderProfilerEncoder.DebugDescription]
//   - [GTAGX2ShaderProfilerEncoder.Description]
//   - [GTAGX2ShaderProfilerEncoder.Hash]
//   - [GTAGX2ShaderProfilerEncoder.Superclass]
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderProfilerEncoder
type GTAGX2ShaderProfilerEncoder struct {
	objectivec.Object
}

// GTAGX2ShaderProfilerEncoderFromID constructs a [GTAGX2ShaderProfilerEncoder] from an objc.ID.
func GTAGX2ShaderProfilerEncoderFromID(id objc.ID) GTAGX2ShaderProfilerEncoder {
	return GTAGX2ShaderProfilerEncoder{objectivec.Object{ID: id}}
}
// Ensure GTAGX2ShaderProfilerEncoder implements IGTAGX2ShaderProfilerEncoder.
var _ IGTAGX2ShaderProfilerEncoder = GTAGX2ShaderProfilerEncoder{}

// An interface definition for the [GTAGX2ShaderProfilerEncoder] class.
//
// # Methods
//
//   - [IGTAGX2ShaderProfilerEncoder.EncodeWithCoder]
//   - [IGTAGX2ShaderProfilerEncoder.FunctionIndex]
//   - [IGTAGX2ShaderProfilerEncoder.SetFunctionIndex]
//   - [IGTAGX2ShaderProfilerEncoder.GpuCommandStartIndex]
//   - [IGTAGX2ShaderProfilerEncoder.SetGpuCommandStartIndex]
//   - [IGTAGX2ShaderProfilerEncoder.LoadTime]
//   - [IGTAGX2ShaderProfilerEncoder.SetLoadTime]
//   - [IGTAGX2ShaderProfilerEncoder.NumGPUCommands]
//   - [IGTAGX2ShaderProfilerEncoder.SetNumGPUCommands]
//   - [IGTAGX2ShaderProfilerEncoder.ObjectId]
//   - [IGTAGX2ShaderProfilerEncoder.SetObjectId]
//   - [IGTAGX2ShaderProfilerEncoder.PointerId]
//   - [IGTAGX2ShaderProfilerEncoder.SetPointerId]
//   - [IGTAGX2ShaderProfilerEncoder.SetIndex]
//   - [IGTAGX2ShaderProfilerEncoder.StoreTime]
//   - [IGTAGX2ShaderProfilerEncoder.SetStoreTime]
//   - [IGTAGX2ShaderProfilerEncoder.TimingInfo]
//   - [IGTAGX2ShaderProfilerEncoder.SetTimingInfo]
//   - [IGTAGX2ShaderProfilerEncoder.InitWithCoder]
//   - [IGTAGX2ShaderProfilerEncoder.DebugDescription]
//   - [IGTAGX2ShaderProfilerEncoder.Description]
//   - [IGTAGX2ShaderProfilerEncoder.Hash]
//   - [IGTAGX2ShaderProfilerEncoder.Superclass]
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderProfilerEncoder
type IGTAGX2ShaderProfilerEncoder interface {
	objectivec.IObject

	// Topic: Methods

	EncodeWithCoder(coder foundation.INSCoder)
	FunctionIndex() uint64
	SetFunctionIndex(value uint64)
	GpuCommandStartIndex() uint32
	SetGpuCommandStartIndex(value uint32)
	LoadTime() uint64
	SetLoadTime(value uint64)
	NumGPUCommands() uint32
	SetNumGPUCommands(value uint32)
	ObjectId() uint64
	SetObjectId(value uint64)
	PointerId() uint64
	SetPointerId(value uint64)
	SetIndex(index uint32)
	StoreTime() uint64
	SetStoreTime(value uint64)
	TimingInfo() IGTShaderProfilerTimingInfo
	SetTimingInfo(value IGTShaderProfilerTimingInfo)
	InitWithCoder(coder foundation.INSCoder) GTAGX2ShaderProfilerEncoder
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objc.Class
}

// Init initializes the instance.
func (g GTAGX2ShaderProfilerEncoder) Init() GTAGX2ShaderProfilerEncoder {
	rv := objc.Send[GTAGX2ShaderProfilerEncoder](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g GTAGX2ShaderProfilerEncoder) Autorelease() GTAGX2ShaderProfilerEncoder {
	rv := objc.Send[GTAGX2ShaderProfilerEncoder](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewGTAGX2ShaderProfilerEncoder creates a new GTAGX2ShaderProfilerEncoder instance.
func NewGTAGX2ShaderProfilerEncoder() GTAGX2ShaderProfilerEncoder {
	class := getGTAGX2ShaderProfilerEncoderClass()
	rv := objc.Send[GTAGX2ShaderProfilerEncoder](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderProfilerEncoder/initWithCoder:
func NewGTAGX2ShaderProfilerEncoderWithCoder(coder objectivec.IObject) GTAGX2ShaderProfilerEncoder {
	instance := getGTAGX2ShaderProfilerEncoderClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return GTAGX2ShaderProfilerEncoderFromID(rv)
}

//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderProfilerEncoder/encodeWithCoder:
func (g GTAGX2ShaderProfilerEncoder) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](g.ID, objc.Sel("encodeWithCoder:"), coder)
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderProfilerEncoder/setIndex:
func (g GTAGX2ShaderProfilerEncoder) SetIndex(index uint32) {
	objc.Send[objc.ID](g.ID, objc.Sel("setIndex:"), index)
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderProfilerEncoder/initWithCoder:
func (g GTAGX2ShaderProfilerEncoder) InitWithCoder(coder foundation.INSCoder) GTAGX2ShaderProfilerEncoder {
	rv := objc.Send[GTAGX2ShaderProfilerEncoder](g.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderProfilerEncoder/supportsSecureCoding
func (_GTAGX2ShaderProfilerEncoderClass GTAGX2ShaderProfilerEncoderClass) SupportsSecureCoding() bool {
	rv := objc.Send[bool](objc.ID(_GTAGX2ShaderProfilerEncoderClass.class), objc.Sel("supportsSecureCoding"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderProfilerEncoder/debugDescription
func (g GTAGX2ShaderProfilerEncoder) DebugDescription() string {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderProfilerEncoder/description
func (g GTAGX2ShaderProfilerEncoder) Description() string {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderProfilerEncoder/functionIndex
func (g GTAGX2ShaderProfilerEncoder) FunctionIndex() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("functionIndex"))
	return rv
}
func (g GTAGX2ShaderProfilerEncoder) SetFunctionIndex(value uint64) {
	objc.Send[struct{}](g.ID, objc.Sel("setFunctionIndex:"), value)
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderProfilerEncoder/gpuCommandStartIndex
func (g GTAGX2ShaderProfilerEncoder) GpuCommandStartIndex() uint32 {
	rv := objc.Send[uint32](g.ID, objc.Sel("gpuCommandStartIndex"))
	return rv
}
func (g GTAGX2ShaderProfilerEncoder) SetGpuCommandStartIndex(value uint32) {
	objc.Send[struct{}](g.ID, objc.Sel("setGpuCommandStartIndex:"), value)
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderProfilerEncoder/hash
func (g GTAGX2ShaderProfilerEncoder) Hash() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("hash"))
	return rv
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderProfilerEncoder/loadTime
func (g GTAGX2ShaderProfilerEncoder) LoadTime() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("loadTime"))
	return rv
}
func (g GTAGX2ShaderProfilerEncoder) SetLoadTime(value uint64) {
	objc.Send[struct{}](g.ID, objc.Sel("setLoadTime:"), value)
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderProfilerEncoder/numGPUCommands
func (g GTAGX2ShaderProfilerEncoder) NumGPUCommands() uint32 {
	rv := objc.Send[uint32](g.ID, objc.Sel("numGPUCommands"))
	return rv
}
func (g GTAGX2ShaderProfilerEncoder) SetNumGPUCommands(value uint32) {
	objc.Send[struct{}](g.ID, objc.Sel("setNumGPUCommands:"), value)
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderProfilerEncoder/objectId
func (g GTAGX2ShaderProfilerEncoder) ObjectId() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("objectId"))
	return rv
}
func (g GTAGX2ShaderProfilerEncoder) SetObjectId(value uint64) {
	objc.Send[struct{}](g.ID, objc.Sel("setObjectId:"), value)
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderProfilerEncoder/pointerId
func (g GTAGX2ShaderProfilerEncoder) PointerId() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("pointerId"))
	return rv
}
func (g GTAGX2ShaderProfilerEncoder) SetPointerId(value uint64) {
	objc.Send[struct{}](g.ID, objc.Sel("setPointerId:"), value)
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderProfilerEncoder/storeTime
func (g GTAGX2ShaderProfilerEncoder) StoreTime() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("storeTime"))
	return rv
}
func (g GTAGX2ShaderProfilerEncoder) SetStoreTime(value uint64) {
	objc.Send[struct{}](g.ID, objc.Sel("setStoreTime:"), value)
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderProfilerEncoder/superclass
func (g GTAGX2ShaderProfilerEncoder) Superclass() objc.Class {
	rv := objc.Send[objc.Class](g.ID, objc.Sel("superclass"))
	return rv
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTAGX2ShaderProfilerEncoder/timingInfo
func (g GTAGX2ShaderProfilerEncoder) TimingInfo() IGTShaderProfilerTimingInfo {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("timingInfo"))
	return GTShaderProfilerTimingInfoFromID(objc.ID(rv))
}
func (g GTAGX2ShaderProfilerEncoder) SetTimingInfo(value IGTShaderProfilerTimingInfo) {
	objc.Send[struct{}](g.ID, objc.Sel("setTimingInfo:"), value)
}

