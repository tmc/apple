// Code generated from Apple documentation for gtshaderprofiler. DO NOT EDIT.

package gtshaderprofiler

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [GTMioShaderProfilerEncoder] class.
var (
	_GTMioShaderProfilerEncoderClass     GTMioShaderProfilerEncoderClass
	_GTMioShaderProfilerEncoderClassOnce sync.Once
)

func getGTMioShaderProfilerEncoderClass() GTMioShaderProfilerEncoderClass {
	_GTMioShaderProfilerEncoderClassOnce.Do(func() {
		_GTMioShaderProfilerEncoderClass = GTMioShaderProfilerEncoderClass{class: objc.GetClass("GTMioShaderProfilerEncoder")}
	})
	return _GTMioShaderProfilerEncoderClass
}

// GetGTMioShaderProfilerEncoderClass returns the class object for GTMioShaderProfilerEncoder.
func GetGTMioShaderProfilerEncoderClass() GTMioShaderProfilerEncoderClass {
	return getGTMioShaderProfilerEncoderClass()
}

type GTMioShaderProfilerEncoderClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (gc GTMioShaderProfilerEncoderClass) Class() objc.Class {
	return gc.class
}

// Alloc allocates memory for a new instance of the class.
func (gc GTMioShaderProfilerEncoderClass) Alloc() GTMioShaderProfilerEncoder {
	rv := objc.Send[GTMioShaderProfilerEncoder](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [GTMioShaderProfilerEncoder.FunctionIndex]
//   - [GTMioShaderProfilerEncoder.GpuCommandStartIndex]
//   - [GTMioShaderProfilerEncoder.LoadTime]
//   - [GTMioShaderProfilerEncoder.NumGPUCommands]
//   - [GTMioShaderProfilerEncoder.PointerId]
//   - [GTMioShaderProfilerEncoder.StoreTime]
//   - [GTMioShaderProfilerEncoder.TimingInfo]
//   - [GTMioShaderProfilerEncoder.InitWithInfoMetadataTraceData]
//   - [GTMioShaderProfilerEncoder.DebugDescription]
//   - [GTMioShaderProfilerEncoder.Description]
//   - [GTMioShaderProfilerEncoder.Hash]
//   - [GTMioShaderProfilerEncoder.Superclass]
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderProfilerEncoder
type GTMioShaderProfilerEncoder struct {
	objectivec.Object
}

// GTMioShaderProfilerEncoderFromID constructs a [GTMioShaderProfilerEncoder] from an objc.ID.
func GTMioShaderProfilerEncoderFromID(id objc.ID) GTMioShaderProfilerEncoder {
	return GTMioShaderProfilerEncoder{objectivec.Object{ID: id}}
}
// Ensure GTMioShaderProfilerEncoder implements IGTMioShaderProfilerEncoder.
var _ IGTMioShaderProfilerEncoder = GTMioShaderProfilerEncoder{}

// An interface definition for the [GTMioShaderProfilerEncoder] class.
//
// # Methods
//
//   - [IGTMioShaderProfilerEncoder.FunctionIndex]
//   - [IGTMioShaderProfilerEncoder.GpuCommandStartIndex]
//   - [IGTMioShaderProfilerEncoder.LoadTime]
//   - [IGTMioShaderProfilerEncoder.NumGPUCommands]
//   - [IGTMioShaderProfilerEncoder.PointerId]
//   - [IGTMioShaderProfilerEncoder.StoreTime]
//   - [IGTMioShaderProfilerEncoder.TimingInfo]
//   - [IGTMioShaderProfilerEncoder.InitWithInfoMetadataTraceData]
//   - [IGTMioShaderProfilerEncoder.DebugDescription]
//   - [IGTMioShaderProfilerEncoder.Description]
//   - [IGTMioShaderProfilerEncoder.Hash]
//   - [IGTMioShaderProfilerEncoder.Superclass]
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderProfilerEncoder
type IGTMioShaderProfilerEncoder interface {
	objectivec.IObject

	// Topic: Methods

	FunctionIndex() uint64
	GpuCommandStartIndex() uint32
	LoadTime() uint64
	NumGPUCommands() uint32
	PointerId() uint64
	StoreTime() uint64
	TimingInfo() IGTShaderProfilerTimingInfo
	InitWithInfoMetadataTraceData(info objectivec.IObject, metadata unsafe.Pointer, data objectivec.IObject) GTMioShaderProfilerEncoder
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objc.Class
}

// Init initializes the instance.
func (g GTMioShaderProfilerEncoder) Init() GTMioShaderProfilerEncoder {
	rv := objc.Send[GTMioShaderProfilerEncoder](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g GTMioShaderProfilerEncoder) Autorelease() GTMioShaderProfilerEncoder {
	rv := objc.Send[GTMioShaderProfilerEncoder](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewGTMioShaderProfilerEncoder creates a new GTMioShaderProfilerEncoder instance.
func NewGTMioShaderProfilerEncoder() GTMioShaderProfilerEncoder {
	class := getGTMioShaderProfilerEncoderClass()
	rv := objc.Send[GTMioShaderProfilerEncoder](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderProfilerEncoder/initWithInfo:metadata:traceData:
func NewGTMioShaderProfilerEncoderWithInfoMetadataTraceData(info objectivec.IObject, metadata unsafe.Pointer, data objectivec.IObject) GTMioShaderProfilerEncoder {
	instance := getGTMioShaderProfilerEncoderClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithInfo:metadata:traceData:"), info, metadata, data)
	return GTMioShaderProfilerEncoderFromID(rv)
}

//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderProfilerEncoder/initWithInfo:metadata:traceData:
func (g GTMioShaderProfilerEncoder) InitWithInfoMetadataTraceData(info objectivec.IObject, metadata unsafe.Pointer, data objectivec.IObject) GTMioShaderProfilerEncoder {
	rv := objc.Send[GTMioShaderProfilerEncoder](g.ID, objc.Sel("initWithInfo:metadata:traceData:"), info, metadata, data)
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderProfilerEncoder/debugDescription
func (g GTMioShaderProfilerEncoder) DebugDescription() string {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderProfilerEncoder/description
func (g GTMioShaderProfilerEncoder) Description() string {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderProfilerEncoder/functionIndex
func (g GTMioShaderProfilerEncoder) FunctionIndex() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("functionIndex"))
	return rv
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderProfilerEncoder/gpuCommandStartIndex
func (g GTMioShaderProfilerEncoder) GpuCommandStartIndex() uint32 {
	rv := objc.Send[uint32](g.ID, objc.Sel("gpuCommandStartIndex"))
	return rv
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderProfilerEncoder/hash
func (g GTMioShaderProfilerEncoder) Hash() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("hash"))
	return rv
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderProfilerEncoder/loadTime
func (g GTMioShaderProfilerEncoder) LoadTime() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("loadTime"))
	return rv
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderProfilerEncoder/numGPUCommands
func (g GTMioShaderProfilerEncoder) NumGPUCommands() uint32 {
	rv := objc.Send[uint32](g.ID, objc.Sel("numGPUCommands"))
	return rv
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderProfilerEncoder/pointerId
func (g GTMioShaderProfilerEncoder) PointerId() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("pointerId"))
	return rv
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderProfilerEncoder/storeTime
func (g GTMioShaderProfilerEncoder) StoreTime() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("storeTime"))
	return rv
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderProfilerEncoder/superclass
func (g GTMioShaderProfilerEncoder) Superclass() objc.Class {
	rv := objc.Send[objc.Class](g.ID, objc.Sel("superclass"))
	return rv
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderProfilerEncoder/timingInfo
func (g GTMioShaderProfilerEncoder) TimingInfo() IGTShaderProfilerTimingInfo {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("timingInfo"))
	return GTShaderProfilerTimingInfoFromID(objc.ID(rv))
}

