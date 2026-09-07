// Code generated from Apple documentation for gtshaderprofiler. DO NOT EDIT.

package gtshaderprofiler

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
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
	rv := objc.SendIfResponds[GTAGX2ShaderProfilerEncoder](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [GTAGX2ShaderProfilerEncoder.EncodeWithCoder]
//   - [GTAGX2ShaderProfilerEncoder.FunctionIndex]
//   - [GTAGX2ShaderProfilerEncoder.SetFunctionIndex]
//   - [GTAGX2ShaderProfilerEncoder.GpuCommandStartIndex]
//   - [GTAGX2ShaderProfilerEncoder.SetGpuCommandStartIndex]
//   - [GTAGX2ShaderProfilerEncoder.Index]
//   - [GTAGX2ShaderProfilerEncoder.SetIndex]
//   - [GTAGX2ShaderProfilerEncoder.LoadTime]
//   - [GTAGX2ShaderProfilerEncoder.SetLoadTime]
//   - [GTAGX2ShaderProfilerEncoder.NumGPUCommands]
//   - [GTAGX2ShaderProfilerEncoder.SetNumGPUCommands]
//   - [GTAGX2ShaderProfilerEncoder.ObjectId]
//   - [GTAGX2ShaderProfilerEncoder.SetObjectId]
//   - [GTAGX2ShaderProfilerEncoder.PointerId]
//   - [GTAGX2ShaderProfilerEncoder.SetPointerId]
//   - [GTAGX2ShaderProfilerEncoder.StoreTime]
//   - [GTAGX2ShaderProfilerEncoder.SetStoreTime]
//   - [GTAGX2ShaderProfilerEncoder.TimingInfo]
//   - [GTAGX2ShaderProfilerEncoder.SetTimingInfo]
//   - [GTAGX2ShaderProfilerEncoder.InitWithCoder]
//   - [GTAGX2ShaderProfilerEncoder.DebugDescription]
//   - [GTAGX2ShaderProfilerEncoder.Description]
//   - [GTAGX2ShaderProfilerEncoder.Hash]
//   - [GTAGX2ShaderProfilerEncoder.Superclass]
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
//   - [IGTAGX2ShaderProfilerEncoder.Index]
//   - [IGTAGX2ShaderProfilerEncoder.SetIndex]
//   - [IGTAGX2ShaderProfilerEncoder.LoadTime]
//   - [IGTAGX2ShaderProfilerEncoder.SetLoadTime]
//   - [IGTAGX2ShaderProfilerEncoder.NumGPUCommands]
//   - [IGTAGX2ShaderProfilerEncoder.SetNumGPUCommands]
//   - [IGTAGX2ShaderProfilerEncoder.ObjectId]
//   - [IGTAGX2ShaderProfilerEncoder.SetObjectId]
//   - [IGTAGX2ShaderProfilerEncoder.PointerId]
//   - [IGTAGX2ShaderProfilerEncoder.SetPointerId]
//   - [IGTAGX2ShaderProfilerEncoder.StoreTime]
//   - [IGTAGX2ShaderProfilerEncoder.SetStoreTime]
//   - [IGTAGX2ShaderProfilerEncoder.TimingInfo]
//   - [IGTAGX2ShaderProfilerEncoder.SetTimingInfo]
//   - [IGTAGX2ShaderProfilerEncoder.InitWithCoder]
//   - [IGTAGX2ShaderProfilerEncoder.DebugDescription]
//   - [IGTAGX2ShaderProfilerEncoder.Description]
//   - [IGTAGX2ShaderProfilerEncoder.Hash]
//   - [IGTAGX2ShaderProfilerEncoder.Superclass]
type IGTAGX2ShaderProfilerEncoder interface {
	objectivec.IObject

	// Topic: Methods

	EncodeWithCoder(coder foundation.INSCoder)
	FunctionIndex() uint64
	SetFunctionIndex(value uint64)
	GpuCommandStartIndex() uint32
	SetGpuCommandStartIndex(value uint32)
	Index() uint32
	SetIndex(value uint32)
	LoadTime() uint64
	SetLoadTime(value uint64)
	NumGPUCommands() uint32
	SetNumGPUCommands(value uint32)
	ObjectId() uint64
	SetObjectId(value uint64)
	PointerId() uint64
	SetPointerId(value uint64)
	StoreTime() uint64
	SetStoreTime(value uint64)
	TimingInfo() IGTShaderProfilerTimingInfo
	SetTimingInfo(value IGTShaderProfilerTimingInfo)
	InitWithCoder(coder foundation.INSCoder) GTAGX2ShaderProfilerEncoder
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objectivec.Class
}

// Init initializes the instance.
func (g GTAGX2ShaderProfilerEncoder) Init() GTAGX2ShaderProfilerEncoder {
	rv := objc.SendIfResponds[GTAGX2ShaderProfilerEncoder](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g GTAGX2ShaderProfilerEncoder) Autorelease() GTAGX2ShaderProfilerEncoder {
	rv := objc.SendIfResponds[GTAGX2ShaderProfilerEncoder](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewGTAGX2ShaderProfilerEncoder creates a new GTAGX2ShaderProfilerEncoder instance.
func NewGTAGX2ShaderProfilerEncoder() GTAGX2ShaderProfilerEncoder {
	class := getGTAGX2ShaderProfilerEncoderClass()
	rv := objc.SendIfResponds[GTAGX2ShaderProfilerEncoder](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewGTAGX2ShaderProfilerEncoderWithCoder(coder objectivec.IObject) GTAGX2ShaderProfilerEncoder {
	instance := getGTAGX2ShaderProfilerEncoderClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return GTAGX2ShaderProfilerEncoderFromID(rv)
}

func (g GTAGX2ShaderProfilerEncoder) EncodeWithCoder(coder foundation.INSCoder) {
	objc.SendIfResponds[objc.ID](g.ID, objc.Sel("encodeWithCoder:"), coder)
}
func (g GTAGX2ShaderProfilerEncoder) InitWithCoder(coder foundation.INSCoder) GTAGX2ShaderProfilerEncoder {
	rv := objc.SendIfResponds[GTAGX2ShaderProfilerEncoder](g.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}

func (_GTAGX2ShaderProfilerEncoderClass GTAGX2ShaderProfilerEncoderClass) SupportsSecureCoding() bool {
	rv := objc.SendIfResponds[bool](objc.ID(_GTAGX2ShaderProfilerEncoderClass.class), objc.Sel("supportsSecureCoding"))
	return rv
}

func (g GTAGX2ShaderProfilerEncoder) DebugDescription() string {
	rv := objc.SendIfResponds[objc.ID](g.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
func (g GTAGX2ShaderProfilerEncoder) Description() string {
	rv := objc.SendIfResponds[objc.ID](g.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
func (g GTAGX2ShaderProfilerEncoder) FunctionIndex() uint64 {
	rv := objc.SendIfResponds[uint64](g.ID, objc.Sel("functionIndex"))
	return rv
}
func (g GTAGX2ShaderProfilerEncoder) SetFunctionIndex(value uint64) {
	objc.SendIfResponds[struct{}](g.ID, objc.Sel("setFunctionIndex:"), value)
}
func (g GTAGX2ShaderProfilerEncoder) GpuCommandStartIndex() uint32 {
	rv := objc.SendIfResponds[uint32](g.ID, objc.Sel("gpuCommandStartIndex"))
	return rv
}
func (g GTAGX2ShaderProfilerEncoder) SetGpuCommandStartIndex(value uint32) {
	objc.SendIfResponds[struct{}](g.ID, objc.Sel("setGpuCommandStartIndex:"), value)
}
func (g GTAGX2ShaderProfilerEncoder) Hash() uint64 {
	rv := objc.SendIfResponds[uint64](g.ID, objc.Sel("hash"))
	return rv
}
func (g GTAGX2ShaderProfilerEncoder) Index() uint32 {
	rv := objc.SendIfResponds[uint32](g.ID, objc.Sel("index"))
	return rv
}
func (g GTAGX2ShaderProfilerEncoder) SetIndex(value uint32) {
	objc.SendIfResponds[struct{}](g.ID, objc.Sel("setIndex:"), value)
}
func (g GTAGX2ShaderProfilerEncoder) LoadTime() uint64 {
	rv := objc.SendIfResponds[uint64](g.ID, objc.Sel("loadTime"))
	return rv
}
func (g GTAGX2ShaderProfilerEncoder) SetLoadTime(value uint64) {
	objc.SendIfResponds[struct{}](g.ID, objc.Sel("setLoadTime:"), value)
}
func (g GTAGX2ShaderProfilerEncoder) NumGPUCommands() uint32 {
	rv := objc.SendIfResponds[uint32](g.ID, objc.Sel("numGPUCommands"))
	return rv
}
func (g GTAGX2ShaderProfilerEncoder) SetNumGPUCommands(value uint32) {
	objc.SendIfResponds[struct{}](g.ID, objc.Sel("setNumGPUCommands:"), value)
}
func (g GTAGX2ShaderProfilerEncoder) ObjectId() uint64 {
	rv := objc.SendIfResponds[uint64](g.ID, objc.Sel("objectId"))
	return rv
}
func (g GTAGX2ShaderProfilerEncoder) SetObjectId(value uint64) {
	objc.SendIfResponds[struct{}](g.ID, objc.Sel("setObjectId:"), value)
}
func (g GTAGX2ShaderProfilerEncoder) PointerId() uint64 {
	rv := objc.SendIfResponds[uint64](g.ID, objc.Sel("pointerId"))
	return rv
}
func (g GTAGX2ShaderProfilerEncoder) SetPointerId(value uint64) {
	objc.SendIfResponds[struct{}](g.ID, objc.Sel("setPointerId:"), value)
}
func (g GTAGX2ShaderProfilerEncoder) StoreTime() uint64 {
	rv := objc.SendIfResponds[uint64](g.ID, objc.Sel("storeTime"))
	return rv
}
func (g GTAGX2ShaderProfilerEncoder) SetStoreTime(value uint64) {
	objc.SendIfResponds[struct{}](g.ID, objc.Sel("setStoreTime:"), value)
}
func (g GTAGX2ShaderProfilerEncoder) Superclass() objectivec.Class {
	rv := objc.SendIfResponds[objectivec.Class](g.ID, objc.Sel("superclass"))
	return objectivec.Class(rv)
}
func (g GTAGX2ShaderProfilerEncoder) TimingInfo() IGTShaderProfilerTimingInfo {
	rv := objc.SendIfResponds[objc.ID](g.ID, objc.Sel("timingInfo"))
	return GTShaderProfilerTimingInfoFromID(objc.ID(rv))
}
func (g GTAGX2ShaderProfilerEncoder) SetTimingInfo(value IGTShaderProfilerTimingInfo) {
	objc.SendIfResponds[struct{}](g.ID, objc.Sel("setTimingInfo:"), value)
}
