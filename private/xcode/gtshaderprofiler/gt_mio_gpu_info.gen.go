// Code generated from Apple documentation for gtshaderprofiler. DO NOT EDIT.

package gtshaderprofiler

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [GTMioGPUInfo] class.
var (
	_GTMioGPUInfoClass     GTMioGPUInfoClass
	_GTMioGPUInfoClassOnce sync.Once
)

func getGTMioGPUInfoClass() GTMioGPUInfoClass {
	_GTMioGPUInfoClassOnce.Do(func() {
		_GTMioGPUInfoClass = GTMioGPUInfoClass{class: objc.GetClass("GTMioGPUInfo")}
	})
	return _GTMioGPUInfoClass
}

// GetGTMioGPUInfoClass returns the class object for GTMioGPUInfo.
func GetGTMioGPUInfoClass() GTMioGPUInfoClass {
	return getGTMioGPUInfoClass()
}

type GTMioGPUInfoClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (gc GTMioGPUInfoClass) Class() objc.Class {
	return gc.class
}

// Alloc allocates memory for a new instance of the class.
func (gc GTMioGPUInfoClass) Alloc() GTMioGPUInfo {
	rv := objc.Send[GTMioGPUInfo](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [GTMioGPUInfo.FormattedName]
//   - [GTMioGPUInfo.GpuGeneration]
//   - [GTMioGPUInfo.GpuRevision]
//   - [GTMioGPUInfo.GpuType]
//   - [GTMioGPUInfo.GpuVariant]
//   - [GTMioGPUInfo.NumGPs]
//   - [GTMioGPUInfo.NumMGPUs]
//   - [GTMioGPUInfo.NumShaderCores]
//   - [GTMioGPUInfo.InitWithGPUInfo]
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioGPUInfo
type GTMioGPUInfo struct {
	objectivec.Object
}

// GTMioGPUInfoFromID constructs a [GTMioGPUInfo] from an objc.ID.
func GTMioGPUInfoFromID(id objc.ID) GTMioGPUInfo {
	return GTMioGPUInfo{objectivec.Object{ID: id}}
}

// Ensure GTMioGPUInfo implements IGTMioGPUInfo.
var _ IGTMioGPUInfo = GTMioGPUInfo{}

// An interface definition for the [GTMioGPUInfo] class.
//
// # Methods
//
//   - [IGTMioGPUInfo.FormattedName]
//   - [IGTMioGPUInfo.GpuGeneration]
//   - [IGTMioGPUInfo.GpuRevision]
//   - [IGTMioGPUInfo.GpuType]
//   - [IGTMioGPUInfo.GpuVariant]
//   - [IGTMioGPUInfo.NumGPs]
//   - [IGTMioGPUInfo.NumMGPUs]
//   - [IGTMioGPUInfo.NumShaderCores]
//   - [IGTMioGPUInfo.InitWithGPUInfo]
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioGPUInfo
type IGTMioGPUInfo interface {
	objectivec.IObject

	// Topic: Methods

	FormattedName(name bool) objectivec.IObject
	GpuGeneration() uint64
	GpuRevision() uint64
	GpuType() uint64
	GpuVariant() uint64
	NumGPs() uint64
	NumMGPUs() uint64
	NumShaderCores() uint64
	InitWithGPUInfo(gPUInfo unsafe.Pointer) GTMioGPUInfo
}

// Init initializes the instance.
func (g GTMioGPUInfo) Init() GTMioGPUInfo {
	rv := objc.Send[GTMioGPUInfo](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g GTMioGPUInfo) Autorelease() GTMioGPUInfo {
	rv := objc.Send[GTMioGPUInfo](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewGTMioGPUInfo creates a new GTMioGPUInfo instance.
func NewGTMioGPUInfo() GTMioGPUInfo {
	class := getGTMioGPUInfoClass()
	rv := objc.Send[GTMioGPUInfo](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioGPUInfo/initWithGPUInfo:
func NewGTMioGPUInfoWithGPUInfo(gPUInfo unsafe.Pointer) GTMioGPUInfo {
	instance := getGTMioGPUInfoClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithGPUInfo:"), gPUInfo)
	return GTMioGPUInfoFromID(rv)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioGPUInfo/formattedName:
func (g GTMioGPUInfo) FormattedName(name bool) objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("formattedName:"), name)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioGPUInfo/gpuType
func (g GTMioGPUInfo) GpuType() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("gpuType"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioGPUInfo/initWithGPUInfo:
func (g GTMioGPUInfo) InitWithGPUInfo(gPUInfo unsafe.Pointer) GTMioGPUInfo {
	rv := objc.Send[GTMioGPUInfo](g.ID, objc.Sel("initWithGPUInfo:"), gPUInfo)
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioGPUInfo/gpuGeneration
func (g GTMioGPUInfo) GpuGeneration() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("gpuGeneration"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioGPUInfo/gpuRevision
func (g GTMioGPUInfo) GpuRevision() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("gpuRevision"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioGPUInfo/gpuVariant
func (g GTMioGPUInfo) GpuVariant() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("gpuVariant"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioGPUInfo/numGPs
func (g GTMioGPUInfo) NumGPs() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("numGPs"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioGPUInfo/numMGPUs
func (g GTMioGPUInfo) NumMGPUs() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("numMGPUs"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioGPUInfo/numShaderCores
func (g GTMioGPUInfo) NumShaderCores() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("numShaderCores"))
	return rv
}
