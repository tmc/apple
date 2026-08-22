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
	rv := objc.SendIfResponds[MLE5ExecutionStreamPool](objc.ID(mc.class), objc.Sel("alloc"))
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
func (m MLE5ExecutionStreamPool) Init() MLE5ExecutionStreamPool {
	rv := objc.SendIfResponds[MLE5ExecutionStreamPool](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLE5ExecutionStreamPool) Autorelease() MLE5ExecutionStreamPool {
	rv := objc.SendIfResponds[MLE5ExecutionStreamPool](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLE5ExecutionStreamPool creates a new MLE5ExecutionStreamPool instance.
func NewMLE5ExecutionStreamPool() MLE5ExecutionStreamPool {
	class := getMLE5ExecutionStreamPoolClass()
	rv := objc.SendIfResponds[MLE5ExecutionStreamPool](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewE5ExecutionStreamPoolWithModelConfigurationModelSignpostId(configuration objectivec.IObject, id uint64) MLE5ExecutionStreamPool {
	instance := getMLE5ExecutionStreamPoolClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithModelConfiguration:modelSignpostId:"), configuration, id)
	return MLE5ExecutionStreamPoolFromID(rv)
}

func (m MLE5ExecutionStreamPool) _emitMappingTracepointForStream(stream objectivec.IObject) {
	objc.SendIfResponds[objc.ID](m.ID, objc.Sel("_emitMappingTracepointForStream:"), stream)
}

// EmitMappingTracepointForStream is an exported wrapper for the private method _emitMappingTracepointForStream.
func (m MLE5ExecutionStreamPool) EmitMappingTracepointForStream(stream objectivec.IObject) error {
	if !objc.RespondsToSelector(m.ID, objc.Sel("_emitMappingTracepointForStream:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_emitMappingTracepointForStream:"}
		return err
	}
	m._emitMappingTracepointForStream(stream)
	return nil
}

// CanEmitMappingTracepointForStream reports whether the receiver responds to the private selector _emitMappingTracepointForStream:.
func (m MLE5ExecutionStreamPool) CanEmitMappingTracepointForStream() bool {
	return objc.RespondsToSelector(m.ID, objc.Sel("_emitMappingTracepointForStream:"))
}
func (m MLE5ExecutionStreamPool) EnableInstrumentsTracing() {
	objc.SendIfResponds[objc.ID](m.ID, objc.Sel("enableInstrumentsTracing"))
}
func (m MLE5ExecutionStreamPool) PutBack(back objectivec.IObject) {
	objc.SendIfResponds[objc.ID](m.ID, objc.Sel("putBack:"), back)
}
func (m MLE5ExecutionStreamPool) TakeOut() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("takeOut"))
	return objectivec.Object{ID: rv}
}
func (m MLE5ExecutionStreamPool) InitWithModelConfigurationModelSignpostId(configuration objectivec.IObject, id uint64) MLE5ExecutionStreamPool {
	rv := objc.SendIfResponds[MLE5ExecutionStreamPool](m.ID, objc.Sel("initWithModelConfiguration:modelSignpostId:"), configuration, id)
	return rv
}

func (m MLE5ExecutionStreamPool) AllStreams() foundation.INSSet {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("allStreams"))
	return foundation.NSSetFromID(objc.ID(rv))
}
func (m MLE5ExecutionStreamPool) ModelConfiguration() IMLModelConfiguration {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("modelConfiguration"))
	return MLModelConfigurationFromID(objc.ID(rv))
}
func (m MLE5ExecutionStreamPool) ModelSignpostId() uint64 {
	rv := objc.SendIfResponds[uint64](m.ID, objc.Sel("modelSignpostId"))
	return rv
}
func (m MLE5ExecutionStreamPool) Pool() foundation.INSSet {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("pool"))
	return foundation.NSSetFromID(objc.ID(rv))
}
func (m MLE5ExecutionStreamPool) SerialQueue() objectivec.Object {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("serialQueue"))
	return objectivec.ObjectFromID(objc.ID(rv))
}
