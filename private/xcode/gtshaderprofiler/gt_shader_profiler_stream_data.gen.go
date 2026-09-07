// Code generated from Apple documentation for gtshaderprofiler. DO NOT EDIT.

package gtshaderprofiler

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [GTShaderProfilerStreamData] class.
var (
	_GTShaderProfilerStreamDataClass     GTShaderProfilerStreamDataClass
	_GTShaderProfilerStreamDataClassOnce sync.Once
)

func getGTShaderProfilerStreamDataClass() GTShaderProfilerStreamDataClass {
	_GTShaderProfilerStreamDataClassOnce.Do(func() {
		_GTShaderProfilerStreamDataClass = GTShaderProfilerStreamDataClass{class: objc.GetClass("GTShaderProfilerStreamData")}
	})
	return _GTShaderProfilerStreamDataClass
}

// GetGTShaderProfilerStreamDataClass returns the class object for GTShaderProfilerStreamData.
func GetGTShaderProfilerStreamDataClass() GTShaderProfilerStreamDataClass {
	return getGTShaderProfilerStreamDataClass()
}

type GTShaderProfilerStreamDataClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (gc GTShaderProfilerStreamDataClass) Class() objc.Class {
	return gc.class
}

// Alloc allocates memory for a new instance of the class.
func (gc GTShaderProfilerStreamDataClass) Alloc() GTShaderProfilerStreamData {
	rv := objc.SendIfResponds[GTShaderProfilerStreamData](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [GTShaderProfilerStreamData.ArchivedAPSCounterData]
//   - [GTShaderProfilerStreamData.ArchivedAPSData]
//   - [GTShaderProfilerStreamData.ArchivedAPSTimelineData]
//   - [GTShaderProfilerStreamData.ArchivedBatchIdFilteredCounterData]
//   - [GTShaderProfilerStreamData.ArchivedGPUTimelineData]
//   - [GTShaderProfilerStreamData.ArchivedShaderProfilerData]
//   - [GTShaderProfilerStreamData.CleanupLocalFiles]
//   - [GTShaderProfilerStreamData.DeviceInfo]
//   - [GTShaderProfilerStreamData.SetDeviceInfo]
//   - [GTShaderProfilerStreamData.GpuGeneration]
//   - [GTShaderProfilerStreamData.MetalDeviceName]
//   - [GTShaderProfilerStreamData.MetalPluginName]
type GTShaderProfilerStreamData struct {
	objectivec.Object
}

// GTShaderProfilerStreamDataFromID constructs a [GTShaderProfilerStreamData] from an objc.ID.
func GTShaderProfilerStreamDataFromID(id objc.ID) GTShaderProfilerStreamData {
	return GTShaderProfilerStreamData{objectivec.Object{ID: id}}
}

// Ensure GTShaderProfilerStreamData implements IGTShaderProfilerStreamData.
var _ IGTShaderProfilerStreamData = GTShaderProfilerStreamData{}

// An interface definition for the [GTShaderProfilerStreamData] class.
//
// # Methods
//
//   - [IGTShaderProfilerStreamData.ArchivedAPSCounterData]
//   - [IGTShaderProfilerStreamData.ArchivedAPSData]
//   - [IGTShaderProfilerStreamData.ArchivedAPSTimelineData]
//   - [IGTShaderProfilerStreamData.ArchivedBatchIdFilteredCounterData]
//   - [IGTShaderProfilerStreamData.ArchivedGPUTimelineData]
//   - [IGTShaderProfilerStreamData.ArchivedShaderProfilerData]
//   - [IGTShaderProfilerStreamData.CleanupLocalFiles]
//   - [IGTShaderProfilerStreamData.DeviceInfo]
//   - [IGTShaderProfilerStreamData.SetDeviceInfo]
//   - [IGTShaderProfilerStreamData.GpuGeneration]
//   - [IGTShaderProfilerStreamData.MetalDeviceName]
//   - [IGTShaderProfilerStreamData.MetalPluginName]
type IGTShaderProfilerStreamData interface {
	objectivec.IObject

	// Topic: Methods

	ArchivedAPSCounterData() foundation.NSData
	ArchivedAPSData() foundation.NSData
	ArchivedAPSTimelineData() foundation.NSData
	ArchivedBatchIdFilteredCounterData() foundation.NSData
	ArchivedGPUTimelineData() foundation.NSData
	ArchivedShaderProfilerData() foundation.NSData
	CleanupLocalFiles()
	DeviceInfo() objectivec.IObject
	SetDeviceInfo(value objectivec.IObject)
	GpuGeneration() uint32
	MetalDeviceName() string
	MetalPluginName() string
}

// Init initializes the instance.
func (g GTShaderProfilerStreamData) Init() GTShaderProfilerStreamData {
	rv := objc.SendIfResponds[GTShaderProfilerStreamData](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g GTShaderProfilerStreamData) Autorelease() GTShaderProfilerStreamData {
	rv := objc.SendIfResponds[GTShaderProfilerStreamData](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewGTShaderProfilerStreamData creates a new GTShaderProfilerStreamData instance.
func NewGTShaderProfilerStreamData() GTShaderProfilerStreamData {
	class := getGTShaderProfilerStreamDataClass()
	rv := objc.SendIfResponds[GTShaderProfilerStreamData](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (g GTShaderProfilerStreamData) ArchivedAPSCounterData() foundation.NSData {
	rv := objc.SendIfResponds[foundation.NSData](g.ID, objc.Sel("archivedAPSCounterData"))
	return foundation.NSData(rv)
}
func (g GTShaderProfilerStreamData) ArchivedAPSData() foundation.NSData {
	rv := objc.SendIfResponds[foundation.NSData](g.ID, objc.Sel("archivedAPSData"))
	return foundation.NSData(rv)
}
func (g GTShaderProfilerStreamData) ArchivedAPSTimelineData() foundation.NSData {
	rv := objc.SendIfResponds[foundation.NSData](g.ID, objc.Sel("archivedAPSTimelineData"))
	return foundation.NSData(rv)
}
func (g GTShaderProfilerStreamData) ArchivedBatchIdFilteredCounterData() foundation.NSData {
	rv := objc.SendIfResponds[foundation.NSData](g.ID, objc.Sel("archivedBatchIdFilteredCounterData"))
	return foundation.NSData(rv)
}
func (g GTShaderProfilerStreamData) ArchivedGPUTimelineData() foundation.NSData {
	rv := objc.SendIfResponds[foundation.NSData](g.ID, objc.Sel("archivedGPUTimelineData"))
	return foundation.NSData(rv)
}
func (g GTShaderProfilerStreamData) ArchivedShaderProfilerData() foundation.NSData {
	rv := objc.SendIfResponds[foundation.NSData](g.ID, objc.Sel("archivedShaderProfilerData"))
	return foundation.NSData(rv)
}
func (g GTShaderProfilerStreamData) CleanupLocalFiles() {
	objc.SendIfResponds[objc.ID](g.ID, objc.Sel("cleanupLocalFiles"))
}

func (_GTShaderProfilerStreamDataClass GTShaderProfilerStreamDataClass) DataForMetadataFromArchivedDataURL(url foundation.NSURL) GTShaderProfilerStreamData {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_GTShaderProfilerStreamDataClass.class), objc.Sel("dataForMetadataFromArchivedDataURL:"), url)
	return GTShaderProfilerStreamDataFromID(rv)
}
func (_GTShaderProfilerStreamDataClass GTShaderProfilerStreamDataClass) DataFromArchivedDataURL(url foundation.NSURL) GTShaderProfilerStreamData {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_GTShaderProfilerStreamDataClass.class), objc.Sel("dataFromArchivedDataURL:"), url)
	return GTShaderProfilerStreamDataFromID(rv)
}
func (_GTShaderProfilerStreamDataClass GTShaderProfilerStreamDataClass) StreamDataClasses() foundation.INSArray {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_GTShaderProfilerStreamDataClass.class), objc.Sel("streamDataClasses"))
	return foundation.NSArrayFromID(rv)
}
func (_GTShaderProfilerStreamDataClass GTShaderProfilerStreamDataClass) SupportsSecureCoding() bool {
	rv := objc.SendIfResponds[bool](objc.ID(_GTShaderProfilerStreamDataClass.class), objc.Sel("supportsSecureCoding"))
	return rv
}

func (g GTShaderProfilerStreamData) DeviceInfo() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](g.ID, objc.Sel("deviceInfo"))
	return objectivec.Object{ID: rv}
}
func (g GTShaderProfilerStreamData) SetDeviceInfo(value objectivec.IObject) {
	objc.SendIfResponds[struct{}](g.ID, objc.Sel("setDeviceInfo:"), value)
}
func (g GTShaderProfilerStreamData) GpuGeneration() uint32 {
	rv := objc.SendIfResponds[uint32](g.ID, objc.Sel("gpuGeneration"))
	return rv
}
func (g GTShaderProfilerStreamData) MetalDeviceName() string {
	rv := objc.SendIfResponds[objc.ID](g.ID, objc.Sel("metalDeviceName"))
	return foundation.NSStringFromID(rv).String()
}
func (g GTShaderProfilerStreamData) MetalPluginName() string {
	rv := objc.SendIfResponds[objc.ID](g.ID, objc.Sel("metalPluginName"))
	return foundation.NSStringFromID(rv).String()
}
