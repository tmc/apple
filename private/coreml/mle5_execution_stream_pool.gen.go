// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLE5ExecutionStreamPool] class.
var (
	_MLE5ExecutionStreamPoolClass     MLE5ExecutionStreamPoolClass
	_MLE5ExecutionStreamPoolClassOnce sync.Once
)

func getMLE5ExecutionStreamPoolClass() MLE5ExecutionStreamPoolClass {
	_MLE5ExecutionStreamPoolClassOnce.Do(func() {
		_MLE5ExecutionStreamPoolClass = MLE5ExecutionStreamPoolClass{class: objc.GetClass("MLE5ExecutionStreamPool")}
	})
	return _MLE5ExecutionStreamPoolClass
}

// GetMLE5ExecutionStreamPoolClass returns the class object for MLE5ExecutionStreamPool.
func GetMLE5ExecutionStreamPoolClass() MLE5ExecutionStreamPoolClass {
	return getMLE5ExecutionStreamPoolClass()
}

type MLE5ExecutionStreamPoolClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLE5ExecutionStreamPoolClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLE5ExecutionStreamPoolClass) Alloc() MLE5ExecutionStreamPool {
	rv := objc.Send[MLE5ExecutionStreamPool](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLE5ExecutionStreamPool._emitMappingTracepointForStream]
//   - [MLE5ExecutionStreamPool.AllStreams]
//   - [MLE5ExecutionStreamPool.EnableInstrumentsTracing]
//   - [MLE5ExecutionStreamPool.ModelConfiguration]
//   - [MLE5ExecutionStreamPool.ModelSignpostId]
//   - [MLE5ExecutionStreamPool.Pool]
//   - [MLE5ExecutionStreamPool.PutBack]
//   - [MLE5ExecutionStreamPool.SerialQueue]
//   - [MLE5ExecutionStreamPool.TakeOut]
//   - [MLE5ExecutionStreamPool.InitWithModelConfigurationModelSignpostId]
//
// See: https://developer.apple.com/documentation/CoreML/MLE5ExecutionStreamPool
type MLE5ExecutionStreamPool struct {
	objectivec.Object
}

// MLE5ExecutionStreamPoolFromID constructs a [MLE5ExecutionStreamPool] from an objc.ID.
func MLE5ExecutionStreamPoolFromID(id objc.ID) MLE5ExecutionStreamPool {
	return MLE5ExecutionStreamPool{objectivec.Object{ID: id}}
}

// Ensure MLE5ExecutionStreamPool implements IMLE5ExecutionStreamPool.
var _ IMLE5ExecutionStreamPool = MLE5ExecutionStreamPool{}

// An interface definition for the [MLE5ExecutionStreamPool] class.
//
// # Methods
//
//   - [IMLE5ExecutionStreamPool._emitMappingTracepointForStream]
//   - [IMLE5ExecutionStreamPool.AllStreams]
//   - [IMLE5ExecutionStreamPool.EnableInstrumentsTracing]
//   - [IMLE5ExecutionStreamPool.ModelConfiguration]
//   - [IMLE5ExecutionStreamPool.ModelSignpostId]
//   - [IMLE5ExecutionStreamPool.Pool]
//   - [IMLE5ExecutionStreamPool.PutBack]
//   - [IMLE5ExecutionStreamPool.SerialQueue]
//   - [IMLE5ExecutionStreamPool.TakeOut]
//   - [IMLE5ExecutionStreamPool.InitWithModelConfigurationModelSignpostId]
//
// See: https://developer.apple.com/documentation/CoreML/MLE5ExecutionStreamPool
type IMLE5ExecutionStreamPool interface {
	objectivec.IObject

	// Topic: Methods

	_emitMappingTracepointForStream(stream objectivec.IObject)
	AllStreams() foundation.INSSet
	EnableInstrumentsTracing()
	ModelConfiguration() IMLModelConfiguration
	ModelSignpostId() uint64
	Pool() foundation.INSSet
	PutBack(back objectivec.IObject)
	SerialQueue() objectivec.Object
	TakeOut() objectivec.IObject
	InitWithModelConfigurationModelSignpostId(configuration objectivec.IObject, id uint64) MLE5ExecutionStreamPool
}

// Init initializes the instance.
func (e MLE5ExecutionStreamPool) Init() MLE5ExecutionStreamPool {
	rv := objc.Send[MLE5ExecutionStreamPool](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e MLE5ExecutionStreamPool) Autorelease() MLE5ExecutionStreamPool {
	rv := objc.Send[MLE5ExecutionStreamPool](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLE5ExecutionStreamPool creates a new MLE5ExecutionStreamPool instance.
func NewMLE5ExecutionStreamPool() MLE5ExecutionStreamPool {
	class := getMLE5ExecutionStreamPoolClass()
	rv := objc.Send[MLE5ExecutionStreamPool](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLE5ExecutionStreamPool/initWithModelConfiguration:modelSignpostId:
func NewE5ExecutionStreamPoolWithModelConfigurationModelSignpostId(configuration objectivec.IObject, id uint64) MLE5ExecutionStreamPool {
	instance := getMLE5ExecutionStreamPoolClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithModelConfiguration:modelSignpostId:"), configuration, id)
	return MLE5ExecutionStreamPoolFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLE5ExecutionStreamPool/_emitMappingTracepointForStream:
func (e MLE5ExecutionStreamPool) _emitMappingTracepointForStream(stream objectivec.IObject) {
	objc.Send[objc.ID](e.ID, objc.Sel("_emitMappingTracepointForStream:"), stream)
}

// EmitMappingTracepointForStream is an exported wrapper for the private method _emitMappingTracepointForStream.
func (e MLE5ExecutionStreamPool) EmitMappingTracepointForStream(stream objectivec.IObject) {
	e._emitMappingTracepointForStream(stream)
}

// See: https://developer.apple.com/documentation/CoreML/MLE5ExecutionStreamPool/enableInstrumentsTracing
func (e MLE5ExecutionStreamPool) EnableInstrumentsTracing() {
	objc.Send[objc.ID](e.ID, objc.Sel("enableInstrumentsTracing"))
}

// See: https://developer.apple.com/documentation/CoreML/MLE5ExecutionStreamPool/putBack:
func (e MLE5ExecutionStreamPool) PutBack(back objectivec.IObject) {
	objc.Send[objc.ID](e.ID, objc.Sel("putBack:"), back)
}

// See: https://developer.apple.com/documentation/CoreML/MLE5ExecutionStreamPool/takeOut
func (e MLE5ExecutionStreamPool) TakeOut() objectivec.IObject {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("takeOut"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/CoreML/MLE5ExecutionStreamPool/initWithModelConfiguration:modelSignpostId:
func (e MLE5ExecutionStreamPool) InitWithModelConfigurationModelSignpostId(configuration objectivec.IObject, id uint64) MLE5ExecutionStreamPool {
	rv := objc.Send[MLE5ExecutionStreamPool](e.ID, objc.Sel("initWithModelConfiguration:modelSignpostId:"), configuration, id)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLE5ExecutionStreamPool/allStreams
func (e MLE5ExecutionStreamPool) AllStreams() foundation.INSSet {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("allStreams"))
	return foundation.NSSetFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/MLE5ExecutionStreamPool/modelConfiguration
func (e MLE5ExecutionStreamPool) ModelConfiguration() IMLModelConfiguration {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("modelConfiguration"))
	return MLModelConfigurationFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/MLE5ExecutionStreamPool/modelSignpostId
func (e MLE5ExecutionStreamPool) ModelSignpostId() uint64 {
	rv := objc.Send[uint64](e.ID, objc.Sel("modelSignpostId"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLE5ExecutionStreamPool/pool
func (e MLE5ExecutionStreamPool) Pool() foundation.INSSet {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("pool"))
	return foundation.NSSetFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/MLE5ExecutionStreamPool/serialQueue
func (e MLE5ExecutionStreamPool) SerialQueue() objectivec.Object {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("serialQueue"))
	return objectivec.ObjectFromID(objc.ID(rv))
}
