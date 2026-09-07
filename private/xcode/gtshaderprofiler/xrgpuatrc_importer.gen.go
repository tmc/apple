// Code generated from Apple documentation for gtshaderprofiler. DO NOT EDIT.

package gtshaderprofiler

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [XRGPUATRCImporter] class.
var (
	_XRGPUATRCImporterClass     XRGPUATRCImporterClass
	_XRGPUATRCImporterClassOnce sync.Once
)

func getXRGPUATRCImporterClass() XRGPUATRCImporterClass {
	_XRGPUATRCImporterClassOnce.Do(func() {
		_XRGPUATRCImporterClass = XRGPUATRCImporterClass{class: objc.GetClass("XRGPUATRCImporter")}
	})
	return _XRGPUATRCImporterClass
}

// GetXRGPUATRCImporterClass returns the class object for XRGPUATRCImporter.
func GetXRGPUATRCImporterClass() XRGPUATRCImporterClass {
	return getXRGPUATRCImporterClass()
}

type XRGPUATRCImporterClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (xc XRGPUATRCImporterClass) Class() objc.Class {
	return xc.class
}

// Alloc allocates memory for a new instance of the class.
func (xc XRGPUATRCImporterClass) Alloc() XRGPUATRCImporter {
	rv := objc.SendIfResponds[XRGPUATRCImporter](objc.ID(xc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [XRGPUATRCImporter._loadKTrace]
//   - [XRGPUATRCImporter._parseAGXBlockLengthContainer]
//   - [XRGPUATRCImporter.AddBlockData]
//   - [XRGPUATRCImporter.AgxDriverConfig]
//   - [XRGPUATRCImporter.AgxTraceConfig]
//   - [XRGPUATRCImporter.BlocksWithTag]
//   - [XRGPUATRCImporter.ContainsBlockWithTag]
//   - [XRGPUATRCImporter.ParseRDEBufferSizeSampleCountCounterCountRdeSourceIndexRdeBufferIndexContainer]
//   - [XRGPUATRCImporter.ParseStringListWithBufferSizeCount]
//   - [XRGPUATRCImporter.RemoveBlocksWithTag]
//   - [XRGPUATRCImporter.InitWithkTraceFile]
type XRGPUATRCImporter struct {
	objectivec.Object
}

// XRGPUATRCImporterFromID constructs a [XRGPUATRCImporter] from an objc.ID.
func XRGPUATRCImporterFromID(id objc.ID) XRGPUATRCImporter {
	return XRGPUATRCImporter{objectivec.Object{ID: id}}
}

// Ensure XRGPUATRCImporter implements IXRGPUATRCImporter.
var _ IXRGPUATRCImporter = XRGPUATRCImporter{}

// An interface definition for the [XRGPUATRCImporter] class.
//
// # Methods
//
//   - [IXRGPUATRCImporter._loadKTrace]
//   - [IXRGPUATRCImporter._parseAGXBlockLengthContainer]
//   - [IXRGPUATRCImporter.AddBlockData]
//   - [IXRGPUATRCImporter.AgxDriverConfig]
//   - [IXRGPUATRCImporter.AgxTraceConfig]
//   - [IXRGPUATRCImporter.BlocksWithTag]
//   - [IXRGPUATRCImporter.ContainsBlockWithTag]
//   - [IXRGPUATRCImporter.ParseRDEBufferSizeSampleCountCounterCountRdeSourceIndexRdeBufferIndexContainer]
//   - [IXRGPUATRCImporter.ParseStringListWithBufferSizeCount]
//   - [IXRGPUATRCImporter.RemoveBlocksWithTag]
//   - [IXRGPUATRCImporter.InitWithkTraceFile]
type IXRGPUATRCImporter interface {
	objectivec.IObject

	// Topic: Methods

	_loadKTrace()
	_parseAGXBlockLengthContainer(aGXBlock string, length uint64, container objectivec.IObject) bool
	AddBlockData(block uint32, data objectivec.IObject)
	AgxDriverConfig() objectivec.IObject
	AgxTraceConfig() objectivec.IObject
	BlocksWithTag(tag uint64) objectivec.IObject
	ContainsBlockWithTag(tag uint64) bool
	ParseRDEBufferSizeSampleCountCounterCountRdeSourceIndexRdeBufferIndexContainer(rDEBuffer string, size uint32, count uint32, count2 uint32, index uint32, index2 uint32, container objectivec.IObject) bool
	ParseStringListWithBufferSizeCount(buffer string, size uint32, count uint32) objectivec.IObject
	RemoveBlocksWithTag(tag uint64)
	InitWithkTraceFile(file objectivec.IObject) XRGPUATRCImporter
}

// Init initializes the instance.
func (x XRGPUATRCImporter) Init() XRGPUATRCImporter {
	rv := objc.SendIfResponds[XRGPUATRCImporter](x.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (x XRGPUATRCImporter) Autorelease() XRGPUATRCImporter {
	rv := objc.SendIfResponds[XRGPUATRCImporter](x.ID, objc.Sel("autorelease"))
	return rv
}

// NewXRGPUATRCImporter creates a new XRGPUATRCImporter instance.
func NewXRGPUATRCImporter() XRGPUATRCImporter {
	class := getXRGPUATRCImporterClass()
	rv := objc.SendIfResponds[XRGPUATRCImporter](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewXRGPUATRCImporterWithkTraceFile(file objectivec.IObject) XRGPUATRCImporter {
	instance := getXRGPUATRCImporterClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithkTraceFile:"), file)
	return XRGPUATRCImporterFromID(rv)
}

func (x XRGPUATRCImporter) _loadKTrace() {
	objc.SendIfResponds[objc.ID](x.ID, objc.Sel("_loadKTrace"))
}

// LoadKTrace is an exported wrapper for the private method _loadKTrace.
func (x XRGPUATRCImporter) LoadKTrace() error {
	if !objc.RespondsToSelector(x.ID, objc.Sel("_loadKTrace")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_loadKTrace"}
		return err
	}
	x._loadKTrace()
	return nil
}

// CanLoadKTrace reports whether the receiver responds to the private selector _loadKTrace.
func (x XRGPUATRCImporter) CanLoadKTrace() bool {
	return objc.RespondsToSelector(x.ID, objc.Sel("_loadKTrace"))
}
func (x XRGPUATRCImporter) _parseAGXBlockLengthContainer(aGXBlock string, length uint64, container objectivec.IObject) bool {
	rv := objc.SendIfResponds[bool](x.ID, objc.Sel("_parseAGXBlock:length:container:"), unsafe.Pointer(unsafe.StringData(aGXBlock+"\x00")), length, container)
	return rv
}

// ParseAGXBlockLengthContainer is an exported wrapper for the private method _parseAGXBlockLengthContainer.
func (x XRGPUATRCImporter) ParseAGXBlockLengthContainer(aGXBlock string, length uint64, container objectivec.IObject) (bool, error) {
	if !objc.RespondsToSelector(x.ID, objc.Sel("_parseAGXBlock:length:container:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_parseAGXBlock:length:container:"}
		return false, err
	}
	return x._parseAGXBlockLengthContainer(aGXBlock, length, container), nil
}

// CanParseAGXBlockLengthContainer reports whether the receiver responds to the private selector _parseAGXBlock:length:container:.
func (x XRGPUATRCImporter) CanParseAGXBlockLengthContainer() bool {
	return objc.RespondsToSelector(x.ID, objc.Sel("_parseAGXBlock:length:container:"))
}
func (x XRGPUATRCImporter) AddBlockData(block uint32, data objectivec.IObject) {
	objc.SendIfResponds[objc.ID](x.ID, objc.Sel("addBlock:data:"), block, data)
}
func (x XRGPUATRCImporter) AgxDriverConfig() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](x.ID, objc.Sel("agxDriverConfig"))
	return objectivec.Object{ID: rv}
}
func (x XRGPUATRCImporter) AgxTraceConfig() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](x.ID, objc.Sel("agxTraceConfig"))
	return objectivec.Object{ID: rv}
}
func (x XRGPUATRCImporter) BlocksWithTag(tag uint64) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](x.ID, objc.Sel("blocksWithTag:"), tag)
	return objectivec.Object{ID: rv}
}
func (x XRGPUATRCImporter) ContainsBlockWithTag(tag uint64) bool {
	rv := objc.SendIfResponds[bool](x.ID, objc.Sel("containsBlockWithTag:"), tag)
	return rv
}
func (x XRGPUATRCImporter) ParseRDEBufferSizeSampleCountCounterCountRdeSourceIndexRdeBufferIndexContainer(rDEBuffer string, size uint32, count uint32, count2 uint32, index uint32, index2 uint32, container objectivec.IObject) bool {
	rv := objc.SendIfResponds[bool](x.ID, objc.Sel("parseRDEBuffer:size:sampleCount:counterCount:rdeSourceIndex:rdeBufferIndex:container:"), unsafe.Pointer(unsafe.StringData(rDEBuffer+"\x00")), size, count, count2, index, index2, container)
	return rv
}
func (x XRGPUATRCImporter) ParseStringListWithBufferSizeCount(buffer string, size uint32, count uint32) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](x.ID, objc.Sel("parseStringListWithBuffer:size:count:"), unsafe.Pointer(unsafe.StringData(buffer+"\x00")), size, count)
	return objectivec.Object{ID: rv}
}
func (x XRGPUATRCImporter) RemoveBlocksWithTag(tag uint64) {
	objc.SendIfResponds[objc.ID](x.ID, objc.Sel("removeBlocksWithTag:"), tag)
}
func (x XRGPUATRCImporter) InitWithkTraceFile(file objectivec.IObject) XRGPUATRCImporter {
	rv := objc.SendIfResponds[XRGPUATRCImporter](x.ID, objc.Sel("initWithkTraceFile:"), file)
	return rv
}

func (_XRGPUATRCImporterClass XRGPUATRCImporterClass) ContainerFromKtraceFile(file objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_XRGPUATRCImporterClass.class), objc.Sel("containerFromKtraceFile:"), file)
	return objectivec.Object{ID: rv}
}
