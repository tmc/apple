// Code generated from Apple documentation for gtshaderprofiler. DO NOT EDIT.

package gtshaderprofiler

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [XRGPUAPSDerivedCounter] class.
var (
	_XRGPUAPSDerivedCounterClass     XRGPUAPSDerivedCounterClass
	_XRGPUAPSDerivedCounterClassOnce sync.Once
)

func getXRGPUAPSDerivedCounterClass() XRGPUAPSDerivedCounterClass {
	_XRGPUAPSDerivedCounterClassOnce.Do(func() {
		_XRGPUAPSDerivedCounterClass = XRGPUAPSDerivedCounterClass{class: objc.GetClass("XRGPUAPSDerivedCounter")}
	})
	return _XRGPUAPSDerivedCounterClass
}

// GetXRGPUAPSDerivedCounterClass returns the class object for XRGPUAPSDerivedCounter.
func GetXRGPUAPSDerivedCounterClass() XRGPUAPSDerivedCounterClass {
	return getXRGPUAPSDerivedCounterClass()
}

type XRGPUAPSDerivedCounterClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (xc XRGPUAPSDerivedCounterClass) Class() objc.Class {
	return xc.class
}

// Alloc allocates memory for a new instance of the class.
func (xc XRGPUAPSDerivedCounterClass) Alloc() XRGPUAPSDerivedCounter {
	rv := objc.Send[XRGPUAPSDerivedCounter](objc.ID(xc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [XRGPUAPSDerivedCounter.CounterId]
//   - [XRGPUAPSDerivedCounter.CounterType]
//   - [XRGPUAPSDerivedCounter.DocString]
//   - [XRGPUAPSDerivedCounter.Name]
//   - [XRGPUAPSDerivedCounter.InitWithNameDocStringTypeCounterId]
// See: https://developer.apple.com/documentation/GTShaderProfiler/XRGPUAPSDerivedCounter
type XRGPUAPSDerivedCounter struct {
	objectivec.Object
}

// XRGPUAPSDerivedCounterFromID constructs a [XRGPUAPSDerivedCounter] from an objc.ID.
func XRGPUAPSDerivedCounterFromID(id objc.ID) XRGPUAPSDerivedCounter {
	return XRGPUAPSDerivedCounter{objectivec.Object{ID: id}}
}
// Ensure XRGPUAPSDerivedCounter implements IXRGPUAPSDerivedCounter.
var _ IXRGPUAPSDerivedCounter = XRGPUAPSDerivedCounter{}

// An interface definition for the [XRGPUAPSDerivedCounter] class.
//
// # Methods
//
//   - [IXRGPUAPSDerivedCounter.CounterId]
//   - [IXRGPUAPSDerivedCounter.CounterType]
//   - [IXRGPUAPSDerivedCounter.DocString]
//   - [IXRGPUAPSDerivedCounter.Name]
//   - [IXRGPUAPSDerivedCounter.InitWithNameDocStringTypeCounterId]
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/XRGPUAPSDerivedCounter
type IXRGPUAPSDerivedCounter interface {
	objectivec.IObject

	// Topic: Methods

	CounterId() uint64
	CounterType() uint32
	DocString() string
	Name() string
	InitWithNameDocStringTypeCounterId(name objectivec.IObject, string_ objectivec.IObject, type_ uint32, id uint64) XRGPUAPSDerivedCounter
}

// Init initializes the instance.
func (x XRGPUAPSDerivedCounter) Init() XRGPUAPSDerivedCounter {
	rv := objc.Send[XRGPUAPSDerivedCounter](x.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (x XRGPUAPSDerivedCounter) Autorelease() XRGPUAPSDerivedCounter {
	rv := objc.Send[XRGPUAPSDerivedCounter](x.ID, objc.Sel("autorelease"))
	return rv
}

// NewXRGPUAPSDerivedCounter creates a new XRGPUAPSDerivedCounter instance.
func NewXRGPUAPSDerivedCounter() XRGPUAPSDerivedCounter {
	class := getXRGPUAPSDerivedCounterClass()
	rv := objc.Send[XRGPUAPSDerivedCounter](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/GTShaderProfiler/XRGPUAPSDerivedCounter/initWithName:docString:type:counterId:
func NewXRGPUAPSDerivedCounterWithNameDocStringTypeCounterId(name objectivec.IObject, string_ objectivec.IObject, type_ uint32, id uint64) XRGPUAPSDerivedCounter {
	instance := getXRGPUAPSDerivedCounterClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithName:docString:type:counterId:"), name, string_, type_, id)
	return XRGPUAPSDerivedCounterFromID(rv)
}

//
// See: https://developer.apple.com/documentation/GTShaderProfiler/XRGPUAPSDerivedCounter/initWithName:docString:type:counterId:
func (x XRGPUAPSDerivedCounter) InitWithNameDocStringTypeCounterId(name objectivec.IObject, string_ objectivec.IObject, type_ uint32, id uint64) XRGPUAPSDerivedCounter {
	rv := objc.Send[XRGPUAPSDerivedCounter](x.ID, objc.Sel("initWithName:docString:type:counterId:"), name, string_, type_, id)
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/XRGPUAPSDerivedCounter/counterId
func (x XRGPUAPSDerivedCounter) CounterId() uint64 {
	rv := objc.Send[uint64](x.ID, objc.Sel("counterId"))
	return rv
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/XRGPUAPSDerivedCounter/counterType
func (x XRGPUAPSDerivedCounter) CounterType() uint32 {
	rv := objc.Send[uint32](x.ID, objc.Sel("counterType"))
	return rv
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/XRGPUAPSDerivedCounter/docString
func (x XRGPUAPSDerivedCounter) DocString() string {
	rv := objc.Send[objc.ID](x.ID, objc.Sel("docString"))
	return foundation.NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/XRGPUAPSDerivedCounter/name
func (x XRGPUAPSDerivedCounter) Name() string {
	rv := objc.Send[objc.ID](x.ID, objc.Sel("name"))
	return foundation.NSStringFromID(rv).String()
}

