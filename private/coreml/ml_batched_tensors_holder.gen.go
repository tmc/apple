// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLBatchedTensorsHolder] class.
var (
	_MLBatchedTensorsHolderClass     MLBatchedTensorsHolderClass
	_MLBatchedTensorsHolderClassOnce sync.Once
)

func getMLBatchedTensorsHolderClass() MLBatchedTensorsHolderClass {
	_MLBatchedTensorsHolderClassOnce.Do(func() {
		_MLBatchedTensorsHolderClass = MLBatchedTensorsHolderClass{class: objc.GetClass("MLBatchedTensorsHolder")}
	})
	return _MLBatchedTensorsHolderClass
}

// GetMLBatchedTensorsHolderClass returns the class object for MLBatchedTensorsHolder.
func GetMLBatchedTensorsHolderClass() MLBatchedTensorsHolderClass {
	return getMLBatchedTensorsHolderClass()
}

type MLBatchedTensorsHolderClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLBatchedTensorsHolderClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLBatchedTensorsHolderClass) Alloc() MLBatchedTensorsHolder {
	rv := objc.Send[MLBatchedTensorsHolder](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [MLBatchedTensorsHolder.BatchedTensors]
//   - [MLBatchedTensorsHolder.NumberOfTensors]
//   - [MLBatchedTensorsHolder.InitWithBatchedTensorsNumberOfTensors]
// See: https://developer.apple.com/documentation/CoreML/MLBatchedTensorsHolder
type MLBatchedTensorsHolder struct {
	objectivec.Object
}

// MLBatchedTensorsHolderFromID constructs a [MLBatchedTensorsHolder] from an objc.ID.
func MLBatchedTensorsHolderFromID(id objc.ID) MLBatchedTensorsHolder {
	return MLBatchedTensorsHolder{objectivec.Object{ID: id}}
}
// Ensure MLBatchedTensorsHolder implements IMLBatchedTensorsHolder.
var _ IMLBatchedTensorsHolder = MLBatchedTensorsHolder{}

// An interface definition for the [MLBatchedTensorsHolder] class.
//
// # Methods
//
//   - [IMLBatchedTensorsHolder.BatchedTensors]
//   - [IMLBatchedTensorsHolder.NumberOfTensors]
//   - [IMLBatchedTensorsHolder.InitWithBatchedTensorsNumberOfTensors]
//
// See: https://developer.apple.com/documentation/CoreML/MLBatchedTensorsHolder
type IMLBatchedTensorsHolder interface {
	objectivec.IObject

	// Topic: Methods

	BatchedTensors() foundation.INSDictionary
	NumberOfTensors() uint64
	InitWithBatchedTensorsNumberOfTensors(tensors objectivec.IObject, tensors2 uint64) MLBatchedTensorsHolder
}

// Init initializes the instance.
func (b MLBatchedTensorsHolder) Init() MLBatchedTensorsHolder {
	rv := objc.Send[MLBatchedTensorsHolder](b.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b MLBatchedTensorsHolder) Autorelease() MLBatchedTensorsHolder {
	rv := objc.Send[MLBatchedTensorsHolder](b.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLBatchedTensorsHolder creates a new MLBatchedTensorsHolder instance.
func NewMLBatchedTensorsHolder() MLBatchedTensorsHolder {
	class := getMLBatchedTensorsHolderClass()
	rv := objc.Send[MLBatchedTensorsHolder](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/CoreML/MLBatchedTensorsHolder/initWithBatchedTensors:numberOfTensors:
func NewBatchedTensorsHolderWithBatchedTensorsNumberOfTensors(tensors objectivec.IObject, tensors2 uint64) MLBatchedTensorsHolder {
	instance := getMLBatchedTensorsHolderClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithBatchedTensors:numberOfTensors:"), tensors, tensors2)
	return MLBatchedTensorsHolderFromID(rv)
}

//
// See: https://developer.apple.com/documentation/CoreML/MLBatchedTensorsHolder/initWithBatchedTensors:numberOfTensors:
func (b MLBatchedTensorsHolder) InitWithBatchedTensorsNumberOfTensors(tensors objectivec.IObject, tensors2 uint64) MLBatchedTensorsHolder {
	rv := objc.Send[MLBatchedTensorsHolder](b.ID, objc.Sel("initWithBatchedTensors:numberOfTensors:"), tensors, tensors2)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLBatchedTensorsHolder/batchedTensors
func (b MLBatchedTensorsHolder) BatchedTensors() foundation.INSDictionary {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("batchedTensors"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/CoreML/MLBatchedTensorsHolder/numberOfTensors
func (b MLBatchedTensorsHolder) NumberOfTensors() uint64 {
	rv := objc.Send[uint64](b.ID, objc.Sel("numberOfTensors"))
	return rv
}

