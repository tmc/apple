// Code generated from Apple documentation for gtshaderprofiler. DO NOT EDIT.

package gtshaderprofiler

import (
	"context"
	"sync"
	"unsafe"

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
	rv := objc.Send[GTShaderProfilerStreamData](objc.ID(gc.class), objc.Sel("alloc"))
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
//   - [GTShaderProfilerStreamData.GPUCommandInfoFromFunctionIndexSubCommandIndex]
//   - [GTShaderProfilerStreamData._setupDataPath]
//   - [GTShaderProfilerStreamData._writeLocalDataDataPathTo]
//   - [GTShaderProfilerStreamData.BatchIdFilterableCounters]
//   - [GTShaderProfilerStreamData.BlitCallCount]
//   - [GTShaderProfilerStreamData.CommandBufferInfoCount]
//   - [GTShaderProfilerStreamData.CommandBufferInfoData]
//   - [GTShaderProfilerStreamData.CommandBuffers]
//   - [GTShaderProfilerStreamData.DataFileURL]
//   - [GTShaderProfilerStreamData.DataFromUnarchvedMetadata]
//   - [GTShaderProfilerStreamData.DataSourceCaptureRange]
//   - [GTShaderProfilerStreamData.DataSourceHasUnusedResources]
//   - [GTShaderProfilerStreamData.DebugDump]
//   - [GTShaderProfilerStreamData.EncodeError]
//   - [GTShaderProfilerStreamData.EncodeAPSArrayForOldHostArray]
//   - [GTShaderProfilerStreamData.EncodeWithCoder]
//   - [GTShaderProfilerStreamData.EncoderInfoCount]
//   - [GTShaderProfilerStreamData.EncoderInfoData]
//   - [GTShaderProfilerStreamData.EncoderInfoFromFunctionIndex]
//   - [GTShaderProfilerStreamData.Encoders]
//   - [GTShaderProfilerStreamData.EnumerateUnarchivedBatchIdFilteredCounterData]
//   - [GTShaderProfilerStreamData.EnumerateUnarchivedGPUTimelineData]
//   - [GTShaderProfilerStreamData.EnumerateUnarchivedShaderProfilerData]
//   - [GTShaderProfilerStreamData.FunctionInfo]
//   - [GTShaderProfilerStreamData.FunctionInfoCount]
//   - [GTShaderProfilerStreamData.FunctionInfoData]
//   - [GTShaderProfilerStreamData.GpuCommandInfoCount]
//   - [GTShaderProfilerStreamData.GpuCommandInfoData]
//   - [GTShaderProfilerStreamData.GpuCommands]
//   - [GTShaderProfilerStreamData.IsPreSiData]
//   - [GTShaderProfilerStreamData.PatchObjectIds]
//   - [GTShaderProfilerStreamData.PipelinePerformanceStatistics]
//   - [GTShaderProfilerStreamData.PipelineStateInfoCount]
//   - [GTShaderProfilerStreamData.PipelineStateInfoData]
//   - [GTShaderProfilerStreamData.PipelineStates]
//   - [GTShaderProfilerStreamData.PreSiBundleURL]
//   - [GTShaderProfilerStreamData.ProfiledExecutionMode]
//   - [GTShaderProfilerStreamData.SetProfiledExecutionMode]
//   - [GTShaderProfilerStreamData.ProfiledPerformanceState]
//   - [GTShaderProfilerStreamData.SetProfiledPerformanceState]
//   - [GTShaderProfilerStreamData.ProfiledProfilerMode]
//   - [GTShaderProfilerStreamData.SetProfiledProfilerMode]
//   - [GTShaderProfilerStreamData.ShortDescription]
//   - [GTShaderProfilerStreamData.Strings]
//   - [GTShaderProfilerStreamData.SupportsFileFormatV2]
//   - [GTShaderProfilerStreamData.SetSupportsFileFormatV2]
//   - [GTShaderProfilerStreamData.TraceName]
//   - [GTShaderProfilerStreamData.SetTraceName]
//   - [GTShaderProfilerStreamData.UnarchivedAPSCounterData]
//   - [GTShaderProfilerStreamData.UnarchivedAPSData]
//   - [GTShaderProfilerStreamData.UnarchivedAPSTimelineData]
//   - [GTShaderProfilerStreamData.UnarchivedBatchIdFilteredCounterData]
//   - [GTShaderProfilerStreamData.UnarchivedGPUTimelineData]
//   - [GTShaderProfilerStreamData.UnarchivedShaderProfilerData]
//   - [GTShaderProfilerStreamData.UnixTimestamp]
//   - [GTShaderProfilerStreamData.InitWithCoder]
//   - [GTShaderProfilerStreamData.InitWithNewFileFormatV2Support]
//   - [GTShaderProfilerStreamData.InitWithPreSiBundle]
//   - [GTShaderProfilerStreamData.Version]
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
//   - [IGTShaderProfilerStreamData.GPUCommandInfoFromFunctionIndexSubCommandIndex]
//   - [IGTShaderProfilerStreamData._setupDataPath]
//   - [IGTShaderProfilerStreamData._writeLocalDataDataPathTo]
//   - [IGTShaderProfilerStreamData.BatchIdFilterableCounters]
//   - [IGTShaderProfilerStreamData.BlitCallCount]
//   - [IGTShaderProfilerStreamData.CommandBufferInfoCount]
//   - [IGTShaderProfilerStreamData.CommandBufferInfoData]
//   - [IGTShaderProfilerStreamData.CommandBuffers]
//   - [IGTShaderProfilerStreamData.DataFileURL]
//   - [IGTShaderProfilerStreamData.DataFromUnarchvedMetadata]
//   - [IGTShaderProfilerStreamData.DataSourceCaptureRange]
//   - [IGTShaderProfilerStreamData.DataSourceHasUnusedResources]
//   - [IGTShaderProfilerStreamData.DebugDump]
//   - [IGTShaderProfilerStreamData.EncodeError]
//   - [IGTShaderProfilerStreamData.EncodeAPSArrayForOldHostArray]
//   - [IGTShaderProfilerStreamData.EncodeWithCoder]
//   - [IGTShaderProfilerStreamData.EncoderInfoCount]
//   - [IGTShaderProfilerStreamData.EncoderInfoData]
//   - [IGTShaderProfilerStreamData.EncoderInfoFromFunctionIndex]
//   - [IGTShaderProfilerStreamData.Encoders]
//   - [IGTShaderProfilerStreamData.EnumerateUnarchivedBatchIdFilteredCounterData]
//   - [IGTShaderProfilerStreamData.EnumerateUnarchivedGPUTimelineData]
//   - [IGTShaderProfilerStreamData.EnumerateUnarchivedShaderProfilerData]
//   - [IGTShaderProfilerStreamData.FunctionInfo]
//   - [IGTShaderProfilerStreamData.FunctionInfoCount]
//   - [IGTShaderProfilerStreamData.FunctionInfoData]
//   - [IGTShaderProfilerStreamData.GpuCommandInfoCount]
//   - [IGTShaderProfilerStreamData.GpuCommandInfoData]
//   - [IGTShaderProfilerStreamData.GpuCommands]
//   - [IGTShaderProfilerStreamData.IsPreSiData]
//   - [IGTShaderProfilerStreamData.PatchObjectIds]
//   - [IGTShaderProfilerStreamData.PipelinePerformanceStatistics]
//   - [IGTShaderProfilerStreamData.PipelineStateInfoCount]
//   - [IGTShaderProfilerStreamData.PipelineStateInfoData]
//   - [IGTShaderProfilerStreamData.PipelineStates]
//   - [IGTShaderProfilerStreamData.PreSiBundleURL]
//   - [IGTShaderProfilerStreamData.ProfiledExecutionMode]
//   - [IGTShaderProfilerStreamData.SetProfiledExecutionMode]
//   - [IGTShaderProfilerStreamData.ProfiledPerformanceState]
//   - [IGTShaderProfilerStreamData.SetProfiledPerformanceState]
//   - [IGTShaderProfilerStreamData.ProfiledProfilerMode]
//   - [IGTShaderProfilerStreamData.SetProfiledProfilerMode]
//   - [IGTShaderProfilerStreamData.ShortDescription]
//   - [IGTShaderProfilerStreamData.Strings]
//   - [IGTShaderProfilerStreamData.SupportsFileFormatV2]
//   - [IGTShaderProfilerStreamData.SetSupportsFileFormatV2]
//   - [IGTShaderProfilerStreamData.TraceName]
//   - [IGTShaderProfilerStreamData.SetTraceName]
//   - [IGTShaderProfilerStreamData.UnarchivedAPSCounterData]
//   - [IGTShaderProfilerStreamData.UnarchivedAPSData]
//   - [IGTShaderProfilerStreamData.UnarchivedAPSTimelineData]
//   - [IGTShaderProfilerStreamData.UnarchivedBatchIdFilteredCounterData]
//   - [IGTShaderProfilerStreamData.UnarchivedGPUTimelineData]
//   - [IGTShaderProfilerStreamData.UnarchivedShaderProfilerData]
//   - [IGTShaderProfilerStreamData.UnixTimestamp]
//   - [IGTShaderProfilerStreamData.InitWithCoder]
//   - [IGTShaderProfilerStreamData.InitWithNewFileFormatV2Support]
//   - [IGTShaderProfilerStreamData.InitWithPreSiBundle]
//   - [IGTShaderProfilerStreamData.Version]
type IGTShaderProfilerStreamData interface {
	objectivec.IObject

	// Topic: Methods

	ArchivedAPSCounterData() foundation.INSArray
	ArchivedAPSData() foundation.INSArray
	ArchivedAPSTimelineData() foundation.INSArray
	ArchivedBatchIdFilteredCounterData() foundation.INSArray
	ArchivedGPUTimelineData() foundation.INSArray
	ArchivedShaderProfilerData() foundation.INSArray
	CleanupLocalFiles()
	DeviceInfo() objectivec.IObject
	SetDeviceInfo(value objectivec.IObject)
	GpuGeneration() uint32
	MetalDeviceName() string
	MetalPluginName() string
	GPUCommandInfoFromFunctionIndexSubCommandIndex(index uint32, index2 int) unsafe.Pointer
	_setupDataPath() objectivec.IObject
	_writeLocalDataDataPathTo(data objectivec.IObject, path objectivec.IObject, to objectivec.IObject)
	BatchIdFilterableCounters() foundation.INSArray
	BlitCallCount() uint64
	CommandBufferInfoCount() uint64
	CommandBufferInfoData() foundation.NSData
	CommandBuffers() objectivec.IObject
	DataFileURL() foundation.NSURL
	DataFromUnarchvedMetadata(metadata objectivec.IObject) objectivec.IObject
	DataSourceCaptureRange() foundation.NSRange
	DataSourceHasUnusedResources() bool
	DebugDump(dump objectivec.IObject)
	EncodeError(encode objectivec.IObject) (objectivec.IObject, error)
	EncodeAPSArrayForOldHostArray(host objectivec.IObject, array objectivec.IObject)
	EncodeWithCoder(coder foundation.INSCoder)
	EncoderInfoCount() uint64
	EncoderInfoData() foundation.NSData
	EncoderInfoFromFunctionIndex(index uint32) unsafe.Pointer
	Encoders() objectivec.IObject
	EnumerateUnarchivedBatchIdFilteredCounterData(data VoidHandler)
	EnumerateUnarchivedGPUTimelineData(data VoidHandler)
	EnumerateUnarchivedShaderProfilerData(data VoidHandler)
	FunctionInfo() objectivec.IObject
	FunctionInfoCount() uint64
	FunctionInfoData() foundation.NSData
	GpuCommandInfoCount() uint64
	GpuCommandInfoData() foundation.NSData
	GpuCommands() objectivec.IObject
	IsPreSiData() bool
	PatchObjectIds(ids objectivec.IObject)
	PipelinePerformanceStatistics() foundation.INSDictionary
	PipelineStateInfoCount() uint64
	PipelineStateInfoData() foundation.NSData
	PipelineStates() objectivec.IObject
	PreSiBundleURL() foundation.NSURL
	ProfiledExecutionMode() uint32
	SetProfiledExecutionMode(value uint32)
	ProfiledPerformanceState() uint32
	SetProfiledPerformanceState(value uint32)
	ProfiledProfilerMode() uint32
	SetProfiledProfilerMode(value uint32)
	ShortDescription() string
	Strings() foundation.INSArray
	SupportsFileFormatV2() bool
	SetSupportsFileFormatV2(value bool)
	TraceName() string
	SetTraceName(value string)
	UnarchivedAPSCounterData() foundation.INSArray
	UnarchivedAPSData() foundation.INSArray
	UnarchivedAPSTimelineData() foundation.INSArray
	UnarchivedBatchIdFilteredCounterData() foundation.INSArray
	UnarchivedGPUTimelineData() foundation.INSArray
	UnarchivedShaderProfilerData() foundation.INSArray
	UnixTimestamp() int64
	InitWithCoder(coder foundation.INSCoder) GTShaderProfilerStreamData
	InitWithNewFileFormatV2Support(v2Support bool) GTShaderProfilerStreamData
	InitWithPreSiBundle(bundle objectivec.IObject) GTShaderProfilerStreamData
	Version() uint64
}

// Init initializes the instance.
func (g GTShaderProfilerStreamData) Init() GTShaderProfilerStreamData {
	rv := objc.Send[GTShaderProfilerStreamData](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g GTShaderProfilerStreamData) Autorelease() GTShaderProfilerStreamData {
	rv := objc.Send[GTShaderProfilerStreamData](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewGTShaderProfilerStreamData creates a new GTShaderProfilerStreamData instance.
func NewGTShaderProfilerStreamData() GTShaderProfilerStreamData {
	class := getGTShaderProfilerStreamDataClass()
	rv := objc.Send[GTShaderProfilerStreamData](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewGTShaderProfilerStreamDataWithCoder(coder objectivec.IObject) GTShaderProfilerStreamData {
	instance := getGTShaderProfilerStreamDataClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return GTShaderProfilerStreamDataFromID(rv)
}

func NewGTShaderProfilerStreamDataWithNewFileFormatV2Support(v2Support bool) GTShaderProfilerStreamData {
	instance := getGTShaderProfilerStreamDataClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithNewFileFormatV2Support:"), v2Support)
	return GTShaderProfilerStreamDataFromID(rv)
}

func NewGTShaderProfilerStreamDataWithPreSiBundle(bundle objectivec.IObject) GTShaderProfilerStreamData {
	instance := getGTShaderProfilerStreamDataClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithPreSiBundle:"), bundle)
	return GTShaderProfilerStreamDataFromID(rv)
}

func (g GTShaderProfilerStreamData) CleanupLocalFiles() {
	objc.Send[objc.ID](g.ID, objc.Sel("cleanupLocalFiles"))
}
func (g GTShaderProfilerStreamData) GPUCommandInfoFromFunctionIndexSubCommandIndex(index uint32, index2 int) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g.ID, objc.Sel("GPUCommandInfoFromFunctionIndex:subCommandIndex:"), index, index2)
	return rv
}
func (g GTShaderProfilerStreamData) _setupDataPath() objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("_setupDataPath"))
	return objectivec.Object{ID: rv}
}

// SetupDataPath is an exported wrapper for the private method _setupDataPath.
func (g GTShaderProfilerStreamData) SetupDataPath() (objectivec.IObject, error) {
	if !objc.RespondsToSelector(g.ID, objc.Sel("_setupDataPath")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_setupDataPath"}
		return nil, err
	}
	return g._setupDataPath(), nil
}

// CanSetupDataPath reports whether the receiver responds to the private selector _setupDataPath.
func (g GTShaderProfilerStreamData) CanSetupDataPath() bool {
	return objc.RespondsToSelector(g.ID, objc.Sel("_setupDataPath"))
}
func (g GTShaderProfilerStreamData) _writeLocalDataDataPathTo(data objectivec.IObject, path objectivec.IObject, to objectivec.IObject) {
	objc.Send[objc.ID](g.ID, objc.Sel("_writeLocalData:dataPath:to:"), data, path, to)
}

// WriteLocalDataDataPathTo is an exported wrapper for the private method _writeLocalDataDataPathTo.
func (g GTShaderProfilerStreamData) WriteLocalDataDataPathTo(data objectivec.IObject, path objectivec.IObject, to objectivec.IObject) error {
	if !objc.RespondsToSelector(g.ID, objc.Sel("_writeLocalData:dataPath:to:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_writeLocalData:dataPath:to:"}
		return err
	}
	g._writeLocalDataDataPathTo(data, path, to)
	return nil
}

// CanWriteLocalDataDataPathTo reports whether the receiver responds to the private selector _writeLocalData:dataPath:to:.
func (g GTShaderProfilerStreamData) CanWriteLocalDataDataPathTo() bool {
	return objc.RespondsToSelector(g.ID, objc.Sel("_writeLocalData:dataPath:to:"))
}
func (g GTShaderProfilerStreamData) DataFromUnarchvedMetadata(metadata objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("dataFromUnarchvedMetadata:"), metadata)
	return objectivec.Object{ID: rv}
}
func (g GTShaderProfilerStreamData) DebugDump(dump objectivec.IObject) {
	objc.Send[objc.ID](g.ID, objc.Sel("debugDump:"), dump)
}
func (g GTShaderProfilerStreamData) EncodeError(encode objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](g.ID, objc.Sel("encode:error:"), encode, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (g GTShaderProfilerStreamData) EncodeAPSArrayForOldHostArray(host objectivec.IObject, array objectivec.IObject) {
	objc.Send[objc.ID](g.ID, objc.Sel("encodeAPSArrayForOldHost:array:"), host, array)
}
func (g GTShaderProfilerStreamData) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](g.ID, objc.Sel("encodeWithCoder:"), coder)
}
func (g GTShaderProfilerStreamData) EncoderInfoFromFunctionIndex(index uint32) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g.ID, objc.Sel("encoderInfoFromFunctionIndex:"), index)
	return rv
}
func (g GTShaderProfilerStreamData) EnumerateUnarchivedBatchIdFilteredCounterData(data VoidHandler) {
	_block0, _ := NewVoidBlock(data)
	objc.Send[objc.ID](g.ID, objc.Sel("enumerateUnarchivedBatchIdFilteredCounterData:"), _block0)
}
func (g GTShaderProfilerStreamData) EnumerateUnarchivedGPUTimelineData(data VoidHandler) {
	_block0, _ := NewVoidBlock(data)
	objc.Send[objc.ID](g.ID, objc.Sel("enumerateUnarchivedGPUTimelineData:"), _block0)
}
func (g GTShaderProfilerStreamData) EnumerateUnarchivedShaderProfilerData(data VoidHandler) {
	_block0, _ := NewVoidBlock(data)
	objc.Send[objc.ID](g.ID, objc.Sel("enumerateUnarchivedShaderProfilerData:"), _block0)
}
func (g GTShaderProfilerStreamData) PatchObjectIds(ids objectivec.IObject) {
	objc.Send[objc.ID](g.ID, objc.Sel("patchObjectIds:"), ids)
}
func (g GTShaderProfilerStreamData) InitWithCoder(coder foundation.INSCoder) GTShaderProfilerStreamData {
	rv := objc.Send[GTShaderProfilerStreamData](g.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}
func (g GTShaderProfilerStreamData) InitWithNewFileFormatV2Support(v2Support bool) GTShaderProfilerStreamData {
	rv := objc.Send[GTShaderProfilerStreamData](g.ID, objc.Sel("initWithNewFileFormatV2Support:"), v2Support)
	return rv
}
func (g GTShaderProfilerStreamData) InitWithPreSiBundle(bundle objectivec.IObject) GTShaderProfilerStreamData {
	rv := objc.Send[GTShaderProfilerStreamData](g.ID, objc.Sel("initWithPreSiBundle:"), bundle)
	return rv
}

func (_GTShaderProfilerStreamDataClass GTShaderProfilerStreamDataClass) DataForMetadataFromArchivedDataURL(url foundation.NSURL) GTShaderProfilerStreamData {
	rv := objc.Send[objc.ID](objc.ID(_GTShaderProfilerStreamDataClass.class), objc.Sel("dataForMetadataFromArchivedDataURL:"), url)
	return GTShaderProfilerStreamDataFromID(rv)
}
func (_GTShaderProfilerStreamDataClass GTShaderProfilerStreamDataClass) DataFromArchivedDataURL(url foundation.NSURL) GTShaderProfilerStreamData {
	rv := objc.Send[objc.ID](objc.ID(_GTShaderProfilerStreamDataClass.class), objc.Sel("dataFromArchivedDataURL:"), url)
	return GTShaderProfilerStreamDataFromID(rv)
}
func (_GTShaderProfilerStreamDataClass GTShaderProfilerStreamDataClass) StreamDataClasses() foundation.INSArray {
	rv := objc.Send[objc.ID](objc.ID(_GTShaderProfilerStreamDataClass.class), objc.Sel("streamDataClasses"))
	return foundation.NSArrayFromID(rv)
}
func (_GTShaderProfilerStreamDataClass GTShaderProfilerStreamDataClass) SupportsSecureCoding() bool {
	rv := objc.Send[bool](objc.ID(_GTShaderProfilerStreamDataClass.class), objc.Sel("supportsSecureCoding"))
	return rv
}
func (_GTShaderProfilerStreamDataClass GTShaderProfilerStreamDataClass) SavedStreamDataFromCaptureArchive(archive objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_GTShaderProfilerStreamDataClass.class), objc.Sel("savedStreamDataFromCaptureArchive:"), archive)
	return objectivec.Object{ID: rv}
}
func (_GTShaderProfilerStreamDataClass GTShaderProfilerStreamDataClass) SteamDataFromData(data objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_GTShaderProfilerStreamDataClass.class), objc.Sel("steamDataFromData:"), data)
	return objectivec.Object{ID: rv}
}

func (g GTShaderProfilerStreamData) DeviceInfo() objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("deviceInfo"))
	return objectivec.Object{ID: rv}
}
func (g GTShaderProfilerStreamData) SetDeviceInfo(value objectivec.IObject) {
	objc.Send[struct{}](g.ID, objc.Sel("setDeviceInfo:"), value)
}
func (g GTShaderProfilerStreamData) GpuGeneration() uint32 {
	rv := objc.Send[uint32](g.ID, objc.Sel("gpuGeneration"))
	return rv
}
func (g GTShaderProfilerStreamData) MetalDeviceName() string {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("metalDeviceName"))
	return foundation.NSStringFromID(rv).String()
}
func (g GTShaderProfilerStreamData) MetalPluginName() string {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("metalPluginName"))
	return foundation.NSStringFromID(rv).String()
}
func (g GTShaderProfilerStreamData) ArchivedAPSCounterData() foundation.INSArray {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("archivedAPSCounterData"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (g GTShaderProfilerStreamData) ArchivedAPSData() foundation.INSArray {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("archivedAPSData"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (g GTShaderProfilerStreamData) ArchivedAPSTimelineData() foundation.INSArray {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("archivedAPSTimelineData"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (g GTShaderProfilerStreamData) ArchivedBatchIdFilteredCounterData() foundation.INSArray {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("archivedBatchIdFilteredCounterData"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (g GTShaderProfilerStreamData) ArchivedGPUTimelineData() foundation.INSArray {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("archivedGPUTimelineData"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (g GTShaderProfilerStreamData) ArchivedShaderProfilerData() foundation.INSArray {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("archivedShaderProfilerData"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (g GTShaderProfilerStreamData) BatchIdFilterableCounters() foundation.INSArray {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("batchIdFilterableCounters"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (g GTShaderProfilerStreamData) BlitCallCount() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("blitCallCount"))
	return rv
}
func (g GTShaderProfilerStreamData) CommandBufferInfoCount() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("commandBufferInfoCount"))
	return rv
}
func (g GTShaderProfilerStreamData) CommandBufferInfoData() foundation.NSData {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("commandBufferInfoData"))
	return foundation.NSDataFromID(objc.ID(rv))
}
func (g GTShaderProfilerStreamData) CommandBuffers() objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("commandBuffers"))
	return objectivec.Object{ID: rv}
}
func (g GTShaderProfilerStreamData) DataFileURL() foundation.NSURL {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("dataFileURL"))
	return foundation.NSURLFromID(objc.ID(rv))
}
func (g GTShaderProfilerStreamData) DataSourceCaptureRange() foundation.NSRange {
	rv := objc.Send[foundation.NSRange](g.ID, objc.Sel("dataSourceCaptureRange"))
	return foundation.NSRange(rv)
}
func (g GTShaderProfilerStreamData) DataSourceHasUnusedResources() bool {
	rv := objc.Send[bool](g.ID, objc.Sel("dataSourceHasUnusedResources"))
	return rv
}
func (g GTShaderProfilerStreamData) EncoderInfoCount() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("encoderInfoCount"))
	return rv
}
func (g GTShaderProfilerStreamData) EncoderInfoData() foundation.NSData {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("encoderInfoData"))
	return foundation.NSDataFromID(objc.ID(rv))
}
func (g GTShaderProfilerStreamData) Encoders() objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("encoders"))
	return objectivec.Object{ID: rv}
}
func (g GTShaderProfilerStreamData) FunctionInfo() objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("functionInfo"))
	return objectivec.Object{ID: rv}
}
func (g GTShaderProfilerStreamData) FunctionInfoCount() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("functionInfoCount"))
	return rv
}
func (g GTShaderProfilerStreamData) FunctionInfoData() foundation.NSData {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("functionInfoData"))
	return foundation.NSDataFromID(objc.ID(rv))
}
func (g GTShaderProfilerStreamData) GpuCommandInfoCount() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("gpuCommandInfoCount"))
	return rv
}
func (g GTShaderProfilerStreamData) GpuCommandInfoData() foundation.NSData {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("gpuCommandInfoData"))
	return foundation.NSDataFromID(objc.ID(rv))
}
func (g GTShaderProfilerStreamData) GpuCommands() objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("gpuCommands"))
	return objectivec.Object{ID: rv}
}
func (g GTShaderProfilerStreamData) IsPreSiData() bool {
	rv := objc.Send[bool](g.ID, objc.Sel("isPreSiData"))
	return rv
}
func (g GTShaderProfilerStreamData) PipelinePerformanceStatistics() foundation.INSDictionary {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("pipelinePerformanceStatistics"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (g GTShaderProfilerStreamData) PipelineStateInfoCount() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("pipelineStateInfoCount"))
	return rv
}
func (g GTShaderProfilerStreamData) PipelineStateInfoData() foundation.NSData {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("pipelineStateInfoData"))
	return foundation.NSDataFromID(objc.ID(rv))
}
func (g GTShaderProfilerStreamData) PipelineStates() objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("pipelineStates"))
	return objectivec.Object{ID: rv}
}
func (g GTShaderProfilerStreamData) PreSiBundleURL() foundation.NSURL {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("preSiBundleURL"))
	return foundation.NSURLFromID(objc.ID(rv))
}
func (g GTShaderProfilerStreamData) ProfiledExecutionMode() uint32 {
	rv := objc.Send[uint32](g.ID, objc.Sel("profiledExecutionMode"))
	return rv
}
func (g GTShaderProfilerStreamData) SetProfiledExecutionMode(value uint32) {
	objc.Send[struct{}](g.ID, objc.Sel("setProfiledExecutionMode:"), value)
}
func (g GTShaderProfilerStreamData) ProfiledPerformanceState() uint32 {
	rv := objc.Send[uint32](g.ID, objc.Sel("profiledPerformanceState"))
	return rv
}
func (g GTShaderProfilerStreamData) SetProfiledPerformanceState(value uint32) {
	objc.Send[struct{}](g.ID, objc.Sel("setProfiledPerformanceState:"), value)
}
func (g GTShaderProfilerStreamData) ProfiledProfilerMode() uint32 {
	rv := objc.Send[uint32](g.ID, objc.Sel("profiledProfilerMode"))
	return rv
}
func (g GTShaderProfilerStreamData) SetProfiledProfilerMode(value uint32) {
	objc.Send[struct{}](g.ID, objc.Sel("setProfiledProfilerMode:"), value)
}
func (g GTShaderProfilerStreamData) ShortDescription() string {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("shortDescription"))
	return foundation.NSStringFromID(rv).String()
}
func (g GTShaderProfilerStreamData) Strings() foundation.INSArray {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("strings"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (g GTShaderProfilerStreamData) SupportsFileFormatV2() bool {
	rv := objc.Send[bool](g.ID, objc.Sel("supportsFileFormatV2"))
	return rv
}
func (g GTShaderProfilerStreamData) SetSupportsFileFormatV2(value bool) {
	objc.Send[struct{}](g.ID, objc.Sel("setSupportsFileFormatV2:"), value)
}
func (g GTShaderProfilerStreamData) TraceName() string {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("traceName"))
	return foundation.NSStringFromID(rv).String()
}
func (g GTShaderProfilerStreamData) SetTraceName(value string) {
	objc.Send[struct{}](g.ID, objc.Sel("setTraceName:"), objc.String(value))
}
func (g GTShaderProfilerStreamData) UnarchivedAPSCounterData() foundation.INSArray {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("unarchivedAPSCounterData"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (g GTShaderProfilerStreamData) UnarchivedAPSData() foundation.INSArray {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("unarchivedAPSData"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (g GTShaderProfilerStreamData) UnarchivedAPSTimelineData() foundation.INSArray {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("unarchivedAPSTimelineData"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (g GTShaderProfilerStreamData) UnarchivedBatchIdFilteredCounterData() foundation.INSArray {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("unarchivedBatchIdFilteredCounterData"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (g GTShaderProfilerStreamData) UnarchivedGPUTimelineData() foundation.INSArray {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("unarchivedGPUTimelineData"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (g GTShaderProfilerStreamData) UnarchivedShaderProfilerData() foundation.INSArray {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("unarchivedShaderProfilerData"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (g GTShaderProfilerStreamData) UnixTimestamp() int64 {
	rv := objc.Send[int64](g.ID, objc.Sel("unixTimestamp"))
	return rv
}
func (g GTShaderProfilerStreamData) Version() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("version"))
	return rv
}

// EnumerateUnarchivedBatchIdFilteredCounterDataSync is a synchronous wrapper around [GTShaderProfilerStreamData.EnumerateUnarchivedBatchIdFilteredCounterData].
// It blocks until the completion handler fires or the context is cancelled.
func (g GTShaderProfilerStreamData) EnumerateUnarchivedBatchIdFilteredCounterDataSync(ctx context.Context) error {
	done := make(chan struct{}, 1)
	g.EnumerateUnarchivedBatchIdFilteredCounterData(func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// EnumerateUnarchivedGPUTimelineDataSync is a synchronous wrapper around [GTShaderProfilerStreamData.EnumerateUnarchivedGPUTimelineData].
// It blocks until the completion handler fires or the context is cancelled.
func (g GTShaderProfilerStreamData) EnumerateUnarchivedGPUTimelineDataSync(ctx context.Context) error {
	done := make(chan struct{}, 1)
	g.EnumerateUnarchivedGPUTimelineData(func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// EnumerateUnarchivedShaderProfilerDataSync is a synchronous wrapper around [GTShaderProfilerStreamData.EnumerateUnarchivedShaderProfilerData].
// It blocks until the completion handler fires or the context is cancelled.
func (g GTShaderProfilerStreamData) EnumerateUnarchivedShaderProfilerDataSync(ctx context.Context) error {
	done := make(chan struct{}, 1)
	g.EnumerateUnarchivedShaderProfilerData(func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
