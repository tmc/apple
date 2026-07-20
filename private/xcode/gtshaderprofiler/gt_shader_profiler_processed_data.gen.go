// Code generated from Apple documentation for gtshaderprofiler. DO NOT EDIT.

package gtshaderprofiler

import (
	"errors"
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
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
type IGTShaderProfilerProcessedData interface {
	objectivec.IObject

	// Topic: Methods

	ArchiveToURLError(url foundation.NSURL) (bool, error)
	EncodeWithCoder(coder foundation.INSCoder)
	GpuGeneration() uint32
	SetGpuGeneration(value uint32)
	MioData() IGTMioTraceData
	ShaderProfilerResult() unsafe.Pointer
	SetShaderProfilerResult(value unsafe.Pointer)
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

func NewGTShaderProfilerProcessedDataWithCoder(coder objectivec.IObject) GTShaderProfilerProcessedData {
	instance := getGTShaderProfilerProcessedDataClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return GTShaderProfilerProcessedDataFromID(rv)
}

func NewGTShaderProfilerProcessedDataWithMioData(data objectivec.IObject) GTShaderProfilerProcessedData {
	instance := getGTShaderProfilerProcessedDataClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithMioData:"), data)
	return GTShaderProfilerProcessedDataFromID(rv)
}

func (g GTShaderProfilerProcessedData) ArchiveToURLError(url foundation.NSURL) (bool, error) {
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
func (g GTShaderProfilerProcessedData) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](g.ID, objc.Sel("encodeWithCoder:"), coder)
}
func (g GTShaderProfilerProcessedData) InitWithCoder(coder foundation.INSCoder) GTShaderProfilerProcessedData {
	rv := objc.Send[GTShaderProfilerProcessedData](g.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}
func (g GTShaderProfilerProcessedData) InitWithMioData(data objectivec.IObject) GTShaderProfilerProcessedData {
	rv := objc.Send[GTShaderProfilerProcessedData](g.ID, objc.Sel("initWithMioData:"), data)
	return rv
}

func (_GTShaderProfilerProcessedDataClass GTShaderProfilerProcessedDataClass) DataFromDataError(data objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_GTShaderProfilerProcessedDataClass.class), objc.Sel("dataFromData:error:"), data, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (_GTShaderProfilerProcessedDataClass GTShaderProfilerProcessedDataClass) DataFromURLError(url foundation.NSURL) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_GTShaderProfilerProcessedDataClass.class), objc.Sel("dataFromURL:error:"), url, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (_GTShaderProfilerProcessedDataClass GTShaderProfilerProcessedDataClass) SupportsSecureCoding() bool {
	rv := objc.Send[bool](objc.ID(_GTShaderProfilerProcessedDataClass.class), objc.Sel("supportsSecureCoding"))
	return rv
}

func (g GTShaderProfilerProcessedData) GpuGeneration() uint32 {
	rv := objc.Send[uint32](g.ID, objc.Sel("gpuGeneration"))
	return rv
}
func (g GTShaderProfilerProcessedData) SetGpuGeneration(value uint32) {
	objc.Send[struct{}](g.ID, objc.Sel("setGpuGeneration:"), value)
}
func (g GTShaderProfilerProcessedData) MioData() IGTMioTraceData {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("mioData"))
	return GTMioTraceDataFromID(objc.ID(rv))
}
func (g GTShaderProfilerProcessedData) ShaderProfilerResult() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g.ID, objc.Sel("shaderProfilerResult"))
	return rv
}
func (g GTShaderProfilerProcessedData) SetShaderProfilerResult(value unsafe.Pointer) {
	objc.Send[struct{}](g.ID, objc.Sel("setShaderProfilerResult:"), value)
}
func (g GTShaderProfilerProcessedData) StreamData() IGTShaderProfilerStreamData {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("streamData"))
	return GTShaderProfilerStreamDataFromID(objc.ID(rv))
}
func (g GTShaderProfilerProcessedData) SetStreamData(value IGTShaderProfilerStreamData) {
	objc.Send[struct{}](g.ID, objc.Sel("setStreamData:"), value)
}
func (g GTShaderProfilerProcessedData) TimelineInfo() IDYWorkloadGPUTimelineInfo {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("timelineInfo"))
	return DYWorkloadGPUTimelineInfoFromID(objc.ID(rv))
}
func (g GTShaderProfilerProcessedData) SetTimelineInfo(value IDYWorkloadGPUTimelineInfo) {
	objc.Send[struct{}](g.ID, objc.Sel("setTimelineInfo:"), value)
}
