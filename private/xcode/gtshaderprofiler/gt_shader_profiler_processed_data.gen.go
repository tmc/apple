// Code generated from Apple documentation for gtshaderprofiler. DO NOT EDIT.

package gtshaderprofiler

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"errors"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [GTShaderProfilerProcessedData] class.
var (
	_GTShaderProfilerProcessedDataClass     GTShaderProfilerProcessedDataClass
	_GTShaderProfilerProcessedDataClassOnce sync.Once
)

func getGTShaderProfilerProcessedDataClass() GTShaderProfilerProcessedDataClass {
	_GTShaderProfilerProcessedDataClassOnce.Do(func() {
		_GTShaderProfilerProcessedDataClass = GTShaderProfilerProcessedDataClass{class: objc.GetClass("GTShaderProfilerProcessedData")}
	})
	return _GTShaderProfilerProcessedDataClass
}

// GetGTShaderProfilerProcessedDataClass returns the class object for GTShaderProfilerProcessedData.
func GetGTShaderProfilerProcessedDataClass() GTShaderProfilerProcessedDataClass {
	return getGTShaderProfilerProcessedDataClass()
}

type GTShaderProfilerProcessedDataClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (gc GTShaderProfilerProcessedDataClass) Class() objc.Class {
	return gc.class
}

// Alloc allocates memory for a new instance of the class.
func (gc GTShaderProfilerProcessedDataClass) Alloc() GTShaderProfilerProcessedData {
	rv := objc.Send[GTShaderProfilerProcessedData](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [GTShaderProfilerProcessedData.ArchiveToURLError]
//   - [GTShaderProfilerProcessedData.EncodeWithCoder]
//   - [GTShaderProfilerProcessedData.GpuGeneration]
//   - [GTShaderProfilerProcessedData.SetGpuGeneration]
//   - [GTShaderProfilerProcessedData.MioData]
//   - [GTShaderProfilerProcessedData.ShaderProfilerResult]
//   - [GTShaderProfilerProcessedData.SetShaderProfilerResult]
//   - [GTShaderProfilerProcessedData.StreamData]
//   - [GTShaderProfilerProcessedData.SetStreamData]
//   - [GTShaderProfilerProcessedData.TimelineInfo]
//   - [GTShaderProfilerProcessedData.SetTimelineInfo]
//   - [GTShaderProfilerProcessedData.InitWithCoder]
//   - [GTShaderProfilerProcessedData.InitWithMioData]
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerProcessedData
type GTShaderProfilerProcessedData struct {
	objectivec.Object
}

// GTShaderProfilerProcessedDataFromID constructs a [GTShaderProfilerProcessedData] from an objc.ID.
func GTShaderProfilerProcessedDataFromID(id objc.ID) GTShaderProfilerProcessedData {
	return GTShaderProfilerProcessedData{objectivec.Object{ID: id}}
}
// Ensure GTShaderProfilerProcessedData implements IGTShaderProfilerProcessedData.
var _ IGTShaderProfilerProcessedData = GTShaderProfilerProcessedData{}

// An interface definition for the [GTShaderProfilerProcessedData] class.
//
// # Methods
//
//   - [IGTShaderProfilerProcessedData.ArchiveToURLError]
//   - [IGTShaderProfilerProcessedData.EncodeWithCoder]
//   - [IGTShaderProfilerProcessedData.GpuGeneration]
//   - [IGTShaderProfilerProcessedData.SetGpuGeneration]
//   - [IGTShaderProfilerProcessedData.MioData]
//   - [IGTShaderProfilerProcessedData.ShaderProfilerResult]
//   - [IGTShaderProfilerProcessedData.SetShaderProfilerResult]
//   - [IGTShaderProfilerProcessedData.StreamData]
//   - [IGTShaderProfilerProcessedData.SetStreamData]
//   - [IGTShaderProfilerProcessedData.TimelineInfo]
//   - [IGTShaderProfilerProcessedData.SetTimelineInfo]
//   - [IGTShaderProfilerProcessedData.InitWithCoder]
//   - [IGTShaderProfilerProcessedData.InitWithMioData]
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerProcessedData
type IGTShaderProfilerProcessedData interface {
	objectivec.IObject

	// Topic: Methods

	ArchiveToURLError(url foundation.INSURL) (bool, error)
	EncodeWithCoder(coder foundation.INSCoder)
	GpuGeneration() uint32
	SetGpuGeneration(value uint32)
	MioData() IGTMioTraceData
	ShaderProfilerResult() objectivec.IObject
	SetShaderProfilerResult(value objectivec.IObject)
	StreamData() IGTShaderProfilerStreamData
	SetStreamData(value IGTShaderProfilerStreamData)
	TimelineInfo() IDYWorkloadGPUTimelineInfo
	SetTimelineInfo(value IDYWorkloadGPUTimelineInfo)
	InitWithCoder(coder foundation.INSCoder) GTShaderProfilerProcessedData
	InitWithMioData(data objectivec.IObject) GTShaderProfilerProcessedData
}

// Init initializes the instance.
func (g GTShaderProfilerProcessedData) Init() GTShaderProfilerProcessedData {
	rv := objc.Send[GTShaderProfilerProcessedData](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g GTShaderProfilerProcessedData) Autorelease() GTShaderProfilerProcessedData {
	rv := objc.Send[GTShaderProfilerProcessedData](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewGTShaderProfilerProcessedData creates a new GTShaderProfilerProcessedData instance.
func NewGTShaderProfilerProcessedData() GTShaderProfilerProcessedData {
	class := getGTShaderProfilerProcessedDataClass()
	rv := objc.Send[GTShaderProfilerProcessedData](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerProcessedData/initWithCoder:
func NewGTShaderProfilerProcessedDataWithCoder(coder objectivec.IObject) GTShaderProfilerProcessedData {
	instance := getGTShaderProfilerProcessedDataClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return GTShaderProfilerProcessedDataFromID(rv)
}

//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerProcessedData/initWithMioData:
func NewGTShaderProfilerProcessedDataWithMioData(data objectivec.IObject) GTShaderProfilerProcessedData {
	instance := getGTShaderProfilerProcessedDataClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithMioData:"), data)
	return GTShaderProfilerProcessedDataFromID(rv)
}

//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerProcessedData/archiveToURL:error:
func (g GTShaderProfilerProcessedData) ArchiveToURLError(url foundation.INSURL) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](g.ID, objc.Sel("archiveToURL:error:"), url, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("archiveToURL:error: returned NO with nil NSError")
	}
	return rv, nil

}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerProcessedData/encodeWithCoder:
func (g GTShaderProfilerProcessedData) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](g.ID, objc.Sel("encodeWithCoder:"), coder)
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerProcessedData/initWithCoder:
func (g GTShaderProfilerProcessedData) InitWithCoder(coder foundation.INSCoder) GTShaderProfilerProcessedData {
	rv := objc.Send[GTShaderProfilerProcessedData](g.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerProcessedData/initWithMioData:
func (g GTShaderProfilerProcessedData) InitWithMioData(data objectivec.IObject) GTShaderProfilerProcessedData {
	rv := objc.Send[GTShaderProfilerProcessedData](g.ID, objc.Sel("initWithMioData:"), data)
	return rv
}

//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerProcessedData/dataFromData:error:
func (_GTShaderProfilerProcessedDataClass GTShaderProfilerProcessedDataClass) DataFromDataError(data objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_GTShaderProfilerProcessedDataClass.class), objc.Sel("dataFromData:error:"), data, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerProcessedData/dataFromURL:error:
func (_GTShaderProfilerProcessedDataClass GTShaderProfilerProcessedDataClass) DataFromURLError(url foundation.INSURL) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_GTShaderProfilerProcessedDataClass.class), objc.Sel("dataFromURL:error:"), url, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerProcessedData/supportsSecureCoding
func (_GTShaderProfilerProcessedDataClass GTShaderProfilerProcessedDataClass) SupportsSecureCoding() bool {
	rv := objc.Send[bool](objc.ID(_GTShaderProfilerProcessedDataClass.class), objc.Sel("supportsSecureCoding"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerProcessedData/gpuGeneration
func (g GTShaderProfilerProcessedData) GpuGeneration() uint32 {
	rv := objc.Send[uint32](g.ID, objc.Sel("gpuGeneration"))
	return rv
}
func (g GTShaderProfilerProcessedData) SetGpuGeneration(value uint32) {
	objc.Send[struct{}](g.ID, objc.Sel("setGpuGeneration:"), value)
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerProcessedData/mioData
func (g GTShaderProfilerProcessedData) MioData() IGTMioTraceData {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("mioData"))
	return GTMioTraceDataFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerProcessedData/shaderProfilerResult
func (g GTShaderProfilerProcessedData) ShaderProfilerResult() objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("shaderProfilerResult"))
	return objectivec.Object{ID: rv}
}
func (g GTShaderProfilerProcessedData) SetShaderProfilerResult(value objectivec.IObject) {
	objc.Send[struct{}](g.ID, objc.Sel("setShaderProfilerResult:"), value)
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerProcessedData/streamData
func (g GTShaderProfilerProcessedData) StreamData() IGTShaderProfilerStreamData {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("streamData"))
	return GTShaderProfilerStreamDataFromID(objc.ID(rv))
}
func (g GTShaderProfilerProcessedData) SetStreamData(value IGTShaderProfilerStreamData) {
	objc.Send[struct{}](g.ID, objc.Sel("setStreamData:"), value)
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerProcessedData/timelineInfo
func (g GTShaderProfilerProcessedData) TimelineInfo() IDYWorkloadGPUTimelineInfo {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("timelineInfo"))
	return DYWorkloadGPUTimelineInfoFromID(objc.ID(rv))
}
func (g GTShaderProfilerProcessedData) SetTimelineInfo(value IDYWorkloadGPUTimelineInfo) {
	objc.Send[struct{}](g.ID, objc.Sel("setTimelineInfo:"), value)
}

