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

// The class instance for the [XRGPUAPSDataContainer] class.
var (
	_XRGPUAPSDataContainerClass     XRGPUAPSDataContainerClass
	_XRGPUAPSDataContainerClassOnce sync.Once
)

func getXRGPUAPSDataContainerClass() XRGPUAPSDataContainerClass {
	_XRGPUAPSDataContainerClassOnce.Do(func() {
		_XRGPUAPSDataContainerClass = XRGPUAPSDataContainerClass{class: objc.GetClass("XRGPUAPSDataContainer")}
	})
	return _XRGPUAPSDataContainerClass
}

// GetXRGPUAPSDataContainerClass returns the class object for XRGPUAPSDataContainer.
func GetXRGPUAPSDataContainerClass() XRGPUAPSDataContainerClass {
	return getXRGPUAPSDataContainerClass()
}

type XRGPUAPSDataContainerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (xc XRGPUAPSDataContainerClass) Class() objc.Class {
	return xc.class
}

// Alloc allocates memory for a new instance of the class.
func (xc XRGPUAPSDataContainerClass) Alloc() XRGPUAPSDataContainer {
	rv := objc.Send[XRGPUAPSDataContainer](objc.ID(xc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [XRGPUAPSDataContainer.AddConfigEntryValue]
//   - [XRGPUAPSDataContainer.AddCustomDataData]
//   - [XRGPUAPSDataContainer.AddDataForRDESourceIndexBufferIndexData]
//   - [XRGPUAPSDataContainer.AddDataForUSCAtIndexData]
//   - [XRGPUAPSDataContainer.AddShaderMapDataVariant]
//   - [XRGPUAPSDataContainer.Config]
//   - [XRGPUAPSDataContainer.ConfigVariant]
//   - [XRGPUAPSDataContainer.Encode]
//   - [XRGPUAPSDataContainer.EncodeWithCoder]
//   - [XRGPUAPSDataContainer.EnumerateRDEData]
//   - [XRGPUAPSDataContainer.EnumerateUSCData]
//   - [XRGPUAPSDataContainer.GetData]
//   - [XRGPUAPSDataContainer.LoadATRCConfig]
//   - [XRGPUAPSDataContainer.LoadGTAConfig]
//   - [XRGPUAPSDataContainer.LoadInstrumentsConfig]
//   - [XRGPUAPSDataContainer.NumBuffersAtRDEIndex]
//   - [XRGPUAPSDataContainer.NumRDEs]
//   - [XRGPUAPSDataContainer.NumUSCs]
//   - [XRGPUAPSDataContainer.ParseGRCConfig]
//   - [XRGPUAPSDataContainer.ReleaseData]
//   - [XRGPUAPSDataContainer.SetRDECounterNames]
//   - [XRGPUAPSDataContainer.ShaderMapData]
//   - [XRGPUAPSDataContainer.ShaderMapDataVariant]
//   - [XRGPUAPSDataContainer.InitWithCoder]
//   - [XRGPUAPSDataContainer.InitWithConfigBaseFolder]
//   - [XRGPUAPSDataContainer.InitWithConfigBaseFolderVariant]
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/XRGPUAPSDataContainer
type XRGPUAPSDataContainer struct {
	objectivec.Object
}

// XRGPUAPSDataContainerFromID constructs a [XRGPUAPSDataContainer] from an objc.ID.
func XRGPUAPSDataContainerFromID(id objc.ID) XRGPUAPSDataContainer {
	return XRGPUAPSDataContainer{objectivec.Object{ID: id}}
}

// Ensure XRGPUAPSDataContainer implements IXRGPUAPSDataContainer.
var _ IXRGPUAPSDataContainer = XRGPUAPSDataContainer{}

// An interface definition for the [XRGPUAPSDataContainer] class.
//
// # Methods
//
//   - [IXRGPUAPSDataContainer.AddConfigEntryValue]
//   - [IXRGPUAPSDataContainer.AddCustomDataData]
//   - [IXRGPUAPSDataContainer.AddDataForRDESourceIndexBufferIndexData]
//   - [IXRGPUAPSDataContainer.AddDataForUSCAtIndexData]
//   - [IXRGPUAPSDataContainer.AddShaderMapDataVariant]
//   - [IXRGPUAPSDataContainer.Config]
//   - [IXRGPUAPSDataContainer.ConfigVariant]
//   - [IXRGPUAPSDataContainer.Encode]
//   - [IXRGPUAPSDataContainer.EncodeWithCoder]
//   - [IXRGPUAPSDataContainer.EnumerateRDEData]
//   - [IXRGPUAPSDataContainer.EnumerateUSCData]
//   - [IXRGPUAPSDataContainer.GetData]
//   - [IXRGPUAPSDataContainer.LoadATRCConfig]
//   - [IXRGPUAPSDataContainer.LoadGTAConfig]
//   - [IXRGPUAPSDataContainer.LoadInstrumentsConfig]
//   - [IXRGPUAPSDataContainer.NumBuffersAtRDEIndex]
//   - [IXRGPUAPSDataContainer.NumRDEs]
//   - [IXRGPUAPSDataContainer.NumUSCs]
//   - [IXRGPUAPSDataContainer.ParseGRCConfig]
//   - [IXRGPUAPSDataContainer.ReleaseData]
//   - [IXRGPUAPSDataContainer.SetRDECounterNames]
//   - [IXRGPUAPSDataContainer.ShaderMapData]
//   - [IXRGPUAPSDataContainer.ShaderMapDataVariant]
//   - [IXRGPUAPSDataContainer.InitWithCoder]
//   - [IXRGPUAPSDataContainer.InitWithConfigBaseFolder]
//   - [IXRGPUAPSDataContainer.InitWithConfigBaseFolderVariant]
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/XRGPUAPSDataContainer
type IXRGPUAPSDataContainer interface {
	objectivec.IObject

	// Topic: Methods

	AddConfigEntryValue(entry objectivec.IObject, value objectivec.IObject)
	AddCustomDataData(data objectivec.IObject, data2 objectivec.IObject)
	AddDataForRDESourceIndexBufferIndexData(index uint32, index2 uint32, data objectivec.IObject)
	AddDataForUSCAtIndexData(index uint32, data objectivec.IObject)
	AddShaderMapDataVariant(data objectivec.IObject, variant uint64)
	Config() foundation.INSDictionary
	ConfigVariant() uint64
	Encode() objectivec.IObject
	EncodeWithCoder(coder foundation.INSCoder)
	EnumerateRDEData(rDEData VoidHandler)
	EnumerateUSCData(uSCData VoidHandler)
	GetData(data objectivec.IObject) objectivec.IObject
	LoadATRCConfig() bool
	LoadGTAConfig() bool
	LoadInstrumentsConfig() bool
	NumBuffersAtRDEIndex(rDEIndex uint64) uint64
	NumRDEs() uint64
	NumUSCs() uint64
	ParseGRCConfig(gRCConfig objectivec.IObject) bool
	ReleaseData()
	SetRDECounterNames(names objectivec.IObject)
	ShaderMapData() objectivec.IObject
	ShaderMapDataVariant() uint64
	InitWithCoder(coder foundation.INSCoder) XRGPUAPSDataContainer
	InitWithConfigBaseFolder(config objectivec.IObject, folder objectivec.IObject) XRGPUAPSDataContainer
	InitWithConfigBaseFolderVariant(config objectivec.IObject, folder objectivec.IObject, variant uint64) XRGPUAPSDataContainer
}

// Init initializes the instance.
func (x XRGPUAPSDataContainer) Init() XRGPUAPSDataContainer {
	rv := objc.Send[XRGPUAPSDataContainer](x.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (x XRGPUAPSDataContainer) Autorelease() XRGPUAPSDataContainer {
	rv := objc.Send[XRGPUAPSDataContainer](x.ID, objc.Sel("autorelease"))
	return rv
}

// NewXRGPUAPSDataContainer creates a new XRGPUAPSDataContainer instance.
func NewXRGPUAPSDataContainer() XRGPUAPSDataContainer {
	class := getXRGPUAPSDataContainerClass()
	rv := objc.Send[XRGPUAPSDataContainer](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/XRGPUAPSDataContainer/initWithCoder:
func NewXRGPUAPSDataContainerWithCoder(coder objectivec.IObject) XRGPUAPSDataContainer {
	instance := getXRGPUAPSDataContainerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return XRGPUAPSDataContainerFromID(rv)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/XRGPUAPSDataContainer/initWithConfig:baseFolder:
func NewXRGPUAPSDataContainerWithConfigBaseFolder(config objectivec.IObject, folder objectivec.IObject) XRGPUAPSDataContainer {
	instance := getXRGPUAPSDataContainerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithConfig:baseFolder:"), config, folder)
	return XRGPUAPSDataContainerFromID(rv)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/XRGPUAPSDataContainer/initWithConfig:baseFolder:variant:
func NewXRGPUAPSDataContainerWithConfigBaseFolderVariant(config objectivec.IObject, folder objectivec.IObject, variant uint64) XRGPUAPSDataContainer {
	instance := getXRGPUAPSDataContainerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithConfig:baseFolder:variant:"), config, folder, variant)
	return XRGPUAPSDataContainerFromID(rv)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/XRGPUAPSDataContainer/addConfigEntry:value:
func (x XRGPUAPSDataContainer) AddConfigEntryValue(entry objectivec.IObject, value objectivec.IObject) {
	objc.Send[objc.ID](x.ID, objc.Sel("addConfigEntry:value:"), entry, value)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/XRGPUAPSDataContainer/addCustomData:data:
func (x XRGPUAPSDataContainer) AddCustomDataData(data objectivec.IObject, data2 objectivec.IObject) {
	objc.Send[objc.ID](x.ID, objc.Sel("addCustomData:data:"), data, data2)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/XRGPUAPSDataContainer/addDataForRDESourceIndex:bufferIndex:data:
func (x XRGPUAPSDataContainer) AddDataForRDESourceIndexBufferIndexData(index uint32, index2 uint32, data objectivec.IObject) {
	objc.Send[objc.ID](x.ID, objc.Sel("addDataForRDESourceIndex:bufferIndex:data:"), index, index2, data)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/XRGPUAPSDataContainer/addDataForUSCAtIndex:data:
func (x XRGPUAPSDataContainer) AddDataForUSCAtIndexData(index uint32, data objectivec.IObject) {
	objc.Send[objc.ID](x.ID, objc.Sel("addDataForUSCAtIndex:data:"), index, data)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/XRGPUAPSDataContainer/addShaderMapData:variant:
func (x XRGPUAPSDataContainer) AddShaderMapDataVariant(data objectivec.IObject, variant uint64) {
	objc.Send[objc.ID](x.ID, objc.Sel("addShaderMapData:variant:"), data, variant)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/XRGPUAPSDataContainer/encode
func (x XRGPUAPSDataContainer) Encode() objectivec.IObject {
	rv := objc.Send[objc.ID](x.ID, objc.Sel("encode"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/XRGPUAPSDataContainer/encodeWithCoder:
func (x XRGPUAPSDataContainer) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](x.ID, objc.Sel("encodeWithCoder:"), coder)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/XRGPUAPSDataContainer/enumerateRDEData:
func (x XRGPUAPSDataContainer) EnumerateRDEData(rDEData VoidHandler) {
	_block0, _ := NewVoidBlock(rDEData)
	objc.Send[objc.ID](x.ID, objc.Sel("enumerateRDEData:"), _block0)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/XRGPUAPSDataContainer/enumerateUSCData:
func (x XRGPUAPSDataContainer) EnumerateUSCData(uSCData VoidHandler) {
	_block0, _ := NewVoidBlock(uSCData)
	objc.Send[objc.ID](x.ID, objc.Sel("enumerateUSCData:"), _block0)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/XRGPUAPSDataContainer/getData:
func (x XRGPUAPSDataContainer) GetData(data objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](x.ID, objc.Sel("getData:"), data)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/XRGPUAPSDataContainer/loadATRCConfig
func (x XRGPUAPSDataContainer) LoadATRCConfig() bool {
	rv := objc.Send[bool](x.ID, objc.Sel("loadATRCConfig"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/XRGPUAPSDataContainer/loadGTAConfig
func (x XRGPUAPSDataContainer) LoadGTAConfig() bool {
	rv := objc.Send[bool](x.ID, objc.Sel("loadGTAConfig"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/XRGPUAPSDataContainer/loadInstrumentsConfig
func (x XRGPUAPSDataContainer) LoadInstrumentsConfig() bool {
	rv := objc.Send[bool](x.ID, objc.Sel("loadInstrumentsConfig"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/XRGPUAPSDataContainer/numBuffersAtRDEIndex:
func (x XRGPUAPSDataContainer) NumBuffersAtRDEIndex(rDEIndex uint64) uint64 {
	rv := objc.Send[uint64](x.ID, objc.Sel("numBuffersAtRDEIndex:"), rDEIndex)
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/XRGPUAPSDataContainer/parseGRCConfig:
func (x XRGPUAPSDataContainer) ParseGRCConfig(gRCConfig objectivec.IObject) bool {
	rv := objc.Send[bool](x.ID, objc.Sel("parseGRCConfig:"), gRCConfig)
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/XRGPUAPSDataContainer/releaseData
func (x XRGPUAPSDataContainer) ReleaseData() {
	objc.Send[objc.ID](x.ID, objc.Sel("releaseData"))
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/XRGPUAPSDataContainer/setRDECounterNames:
func (x XRGPUAPSDataContainer) SetRDECounterNames(names objectivec.IObject) {
	objc.Send[objc.ID](x.ID, objc.Sel("setRDECounterNames:"), names)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/XRGPUAPSDataContainer/initWithCoder:
func (x XRGPUAPSDataContainer) InitWithCoder(coder foundation.INSCoder) XRGPUAPSDataContainer {
	rv := objc.Send[XRGPUAPSDataContainer](x.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/XRGPUAPSDataContainer/initWithConfig:baseFolder:
func (x XRGPUAPSDataContainer) InitWithConfigBaseFolder(config objectivec.IObject, folder objectivec.IObject) XRGPUAPSDataContainer {
	rv := objc.Send[XRGPUAPSDataContainer](x.ID, objc.Sel("initWithConfig:baseFolder:"), config, folder)
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/XRGPUAPSDataContainer/initWithConfig:baseFolder:variant:
func (x XRGPUAPSDataContainer) InitWithConfigBaseFolderVariant(config objectivec.IObject, folder objectivec.IObject, variant uint64) XRGPUAPSDataContainer {
	rv := objc.Send[XRGPUAPSDataContainer](x.ID, objc.Sel("initWithConfig:baseFolder:variant:"), config, folder, variant)
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/XRGPUAPSDataContainer/configVariantFromConfig:
func (_XRGPUAPSDataContainerClass XRGPUAPSDataContainerClass) ConfigVariantFromConfig(config objectivec.IObject) uint64 {
	rv := objc.Send[uint64](objc.ID(_XRGPUAPSDataContainerClass.class), objc.Sel("configVariantFromConfig:"), config)
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/XRGPUAPSDataContainer/fromData:error:
func (_XRGPUAPSDataContainerClass XRGPUAPSDataContainerClass) FromDataError(data objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_XRGPUAPSDataContainerClass.class), objc.Sel("fromData:error:"), data, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// See: https://developer.apple.com/documentation/GTShaderProfiler/XRGPUAPSDataContainer/supportsSecureCoding
func (_XRGPUAPSDataContainerClass XRGPUAPSDataContainerClass) SupportsSecureCoding() bool {
	rv := objc.Send[bool](objc.ID(_XRGPUAPSDataContainerClass.class), objc.Sel("supportsSecureCoding"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/XRGPUAPSDataContainer/config
func (x XRGPUAPSDataContainer) Config() foundation.INSDictionary {
	rv := objc.Send[objc.ID](x.ID, objc.Sel("config"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/XRGPUAPSDataContainer/configVariant
func (x XRGPUAPSDataContainer) ConfigVariant() uint64 {
	rv := objc.Send[uint64](x.ID, objc.Sel("configVariant"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/XRGPUAPSDataContainer/numRDEs
func (x XRGPUAPSDataContainer) NumRDEs() uint64 {
	rv := objc.Send[uint64](x.ID, objc.Sel("numRDEs"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/XRGPUAPSDataContainer/numUSCs
func (x XRGPUAPSDataContainer) NumUSCs() uint64 {
	rv := objc.Send[uint64](x.ID, objc.Sel("numUSCs"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/XRGPUAPSDataContainer/shaderMapData
func (x XRGPUAPSDataContainer) ShaderMapData() objectivec.IObject {
	rv := objc.Send[objc.ID](x.ID, objc.Sel("shaderMapData"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/XRGPUAPSDataContainer/shaderMapDataVariant
func (x XRGPUAPSDataContainer) ShaderMapDataVariant() uint64 {
	rv := objc.Send[uint64](x.ID, objc.Sel("shaderMapDataVariant"))
	return rv
}

// EnumerateRDEDataSync is a synchronous wrapper around [XRGPUAPSDataContainer.EnumerateRDEData].
// It blocks until the completion handler fires or the context is cancelled.
func (x XRGPUAPSDataContainer) EnumerateRDEDataSync(ctx context.Context) error {
	done := make(chan struct{}, 1)
	x.EnumerateRDEData(func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// EnumerateUSCDataSync is a synchronous wrapper around [XRGPUAPSDataContainer.EnumerateUSCData].
// It blocks until the completion handler fires or the context is cancelled.
func (x XRGPUAPSDataContainer) EnumerateUSCDataSync(ctx context.Context) error {
	done := make(chan struct{}, 1)
	x.EnumerateUSCData(func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
