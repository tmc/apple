// Code generated from Apple documentation for gtshaderprofiler. DO NOT EDIT.

package gtshaderprofiler

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [XRGPUAPSDataProcessor] class.
var (
	_XRGPUAPSDataProcessorClass     XRGPUAPSDataProcessorClass
	_XRGPUAPSDataProcessorClassOnce sync.Once
)

func getXRGPUAPSDataProcessorClass() XRGPUAPSDataProcessorClass {
	_XRGPUAPSDataProcessorClassOnce.Do(func() {
		_XRGPUAPSDataProcessorClass = XRGPUAPSDataProcessorClass{class: objc.GetClass("XRGPUAPSDataProcessor")}
	})
	return _XRGPUAPSDataProcessorClass
}

// GetXRGPUAPSDataProcessorClass returns the class object for XRGPUAPSDataProcessor.
func GetXRGPUAPSDataProcessorClass() XRGPUAPSDataProcessorClass {
	return getXRGPUAPSDataProcessorClass()
}

type XRGPUAPSDataProcessorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (xc XRGPUAPSDataProcessorClass) Class() objc.Class {
	return xc.class
}

// Alloc allocates memory for a new instance of the class.
func (xc XRGPUAPSDataProcessorClass) Alloc() XRGPUAPSDataProcessor {
	rv := objc.SendIfResponds[XRGPUAPSDataProcessor](objc.ID(xc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [XRGPUAPSDataProcessor.DeriveAPSCountersNumCoresCounterSet]
//   - [XRGPUAPSDataProcessor.CounterConfigForGRCCounterSet]
//   - [XRGPUAPSDataProcessor.LoadCounterGraphConfig]
//   - [XRGPUAPSDataProcessor.InitWithGPUGenerationVariantRevConfigOptions]
//   - [XRGPUAPSDataProcessor.AcceleratorID]
//   - [XRGPUAPSDataProcessor.Config]
//   - [XRGPUAPSDataProcessor.SetConfig]
type XRGPUAPSDataProcessor struct {
	objectivec.Object
}

// XRGPUAPSDataProcessorFromID constructs a [XRGPUAPSDataProcessor] from an objc.ID.
func XRGPUAPSDataProcessorFromID(id objc.ID) XRGPUAPSDataProcessor {
	return XRGPUAPSDataProcessor{objectivec.Object{ID: id}}
}

// Ensure XRGPUAPSDataProcessor implements IXRGPUAPSDataProcessor.
var _ IXRGPUAPSDataProcessor = XRGPUAPSDataProcessor{}

// An interface definition for the [XRGPUAPSDataProcessor] class.
//
// # Methods
//
//   - [IXRGPUAPSDataProcessor.DeriveAPSCountersNumCoresCounterSet]
//   - [IXRGPUAPSDataProcessor.CounterConfigForGRCCounterSet]
//   - [IXRGPUAPSDataProcessor.LoadCounterGraphConfig]
//   - [IXRGPUAPSDataProcessor.InitWithGPUGenerationVariantRevConfigOptions]
//   - [IXRGPUAPSDataProcessor.AcceleratorID]
//   - [IXRGPUAPSDataProcessor.Config]
//   - [IXRGPUAPSDataProcessor.SetConfig]
type IXRGPUAPSDataProcessor interface {
	objectivec.IObject

	// Topic: Methods

	DeriveAPSCountersNumCoresCounterSet(apsCounters unsafe.Pointer, numCores uint32, counterSet uint64) bool
	CounterConfigForGRCCounterSet(grc bool, counterSet uint64) objectivec.IObject
	LoadCounterGraphConfig() objectivec.IObject
	InitWithGPUGenerationVariantRevConfigOptions(gPUGeneration uint32, variant uint32, rev uint32, config foundation.INSDictionary, options uint32) XRGPUAPSDataProcessor
	AcceleratorID() uint64
	Config() foundation.INSDictionary
	SetConfig(value foundation.INSDictionary)
}

// Init initializes the instance.
func (x XRGPUAPSDataProcessor) Init() XRGPUAPSDataProcessor {
	rv := objc.SendIfResponds[XRGPUAPSDataProcessor](x.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (x XRGPUAPSDataProcessor) Autorelease() XRGPUAPSDataProcessor {
	rv := objc.SendIfResponds[XRGPUAPSDataProcessor](x.ID, objc.Sel("autorelease"))
	return rv
}

// NewXRGPUAPSDataProcessor creates a new XRGPUAPSDataProcessor instance.
func NewXRGPUAPSDataProcessor() XRGPUAPSDataProcessor {
	class := getXRGPUAPSDataProcessorClass()
	rv := objc.SendIfResponds[XRGPUAPSDataProcessor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewXRGPUAPSDataProcessorWithGPUGenerationVariantRevConfigOptions(gPUGeneration uint32, variant uint32, rev uint32, config foundation.INSDictionary, options uint32) XRGPUAPSDataProcessor {
	instance := getXRGPUAPSDataProcessorClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithGPUGeneration:variant:rev:config:options:"), gPUGeneration, variant, rev, config, options)
	return XRGPUAPSDataProcessorFromID(rv)
}

func (x XRGPUAPSDataProcessor) DeriveAPSCountersNumCoresCounterSet(apsCounters unsafe.Pointer, numCores uint32, counterSet uint64) bool {
	rv := objc.SendIfResponds[bool](x.ID, objc.Sel("deriveAPSCounters:numCores:counterSet:"), apsCounters, numCores, counterSet)
	return rv
}
func (x XRGPUAPSDataProcessor) CounterConfigForGRCCounterSet(grc bool, counterSet uint64) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](x.ID, objc.Sel("counterConfigForGRC:counterSet:"), grc, counterSet)
	return objectivec.Object{ID: rv}
}
func (x XRGPUAPSDataProcessor) LoadCounterGraphConfig() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](x.ID, objc.Sel("loadCounterGraphConfig"))
	return objectivec.Object{ID: rv}
}
func (x XRGPUAPSDataProcessor) InitWithGPUGenerationVariantRevConfigOptions(gPUGeneration uint32, variant uint32, rev uint32, config foundation.INSDictionary, options uint32) XRGPUAPSDataProcessor {
	rv := objc.SendIfResponds[XRGPUAPSDataProcessor](x.ID, objc.Sel("initWithGPUGeneration:variant:rev:config:options:"), gPUGeneration, variant, rev, config, options)
	return rv
}

func (_XRGPUAPSDataProcessorClass XRGPUAPSDataProcessorClass) ProcessorFromConfigOptions(config foundation.INSDictionary, options uint32) XRGPUAPSDataProcessor {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_XRGPUAPSDataProcessorClass.class), objc.Sel("processorFromConfig:options:"), config, options)
	return XRGPUAPSDataProcessorFromID(rv)
}

func (x XRGPUAPSDataProcessor) AcceleratorID() uint64 {
	rv := objc.SendIfResponds[uint64](x.ID, objc.Sel("acceleratorID"))
	return rv
}
func (x XRGPUAPSDataProcessor) Config() foundation.INSDictionary {
	rv := objc.SendIfResponds[objc.ID](x.ID, objc.Sel("config"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (x XRGPUAPSDataProcessor) SetConfig(value foundation.INSDictionary) {
	objc.SendIfResponds[struct{}](x.ID, objc.Sel("setConfig:"), value)
}
