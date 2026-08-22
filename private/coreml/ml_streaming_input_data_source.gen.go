// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLStreamingInputDataSource] class.
var (
	_MLStreamingInputDataSourceClass     MLStreamingInputDataSourceClass
	_MLStreamingInputDataSourceClassOnce sync.Once
)

func getMLStreamingInputDataSourceClass() MLStreamingInputDataSourceClass {
	_MLStreamingInputDataSourceClassOnce.Do(func() {
		_MLStreamingInputDataSourceClass = MLStreamingInputDataSourceClass{class: objc.GetClass("MLStreamingInputDataSource")}
	})
	return _MLStreamingInputDataSourceClass
}

// GetMLStreamingInputDataSourceClass returns the class object for MLStreamingInputDataSource.
func GetMLStreamingInputDataSourceClass() MLStreamingInputDataSourceClass {
	return getMLStreamingInputDataSourceClass()
}

type MLStreamingInputDataSourceClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLStreamingInputDataSourceClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLStreamingInputDataSourceClass) Alloc() MLStreamingInputDataSource {
	rv := objc.SendIfResponds[MLStreamingInputDataSource](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLStreamingInputDataSource.AppendBatchedTensorsNumberOfTensors]
//   - [MLStreamingInputDataSource.BatchAtIndexError]
//   - [MLStreamingInputDataSource.BatchSize]
//   - [MLStreamingInputDataSource.SetBatchSize]
//   - [MLStreamingInputDataSource.DataSources]
//   - [MLStreamingInputDataSource.SetDataSources]
//   - [MLStreamingInputDataSource.NumberOfBatches]
//   - [MLStreamingInputDataSource.SizeOfBatchAtIndex]
//   - [MLStreamingInputDataSource.InitWithBatchSize]
type MLStreamingInputDataSource struct {
	objectivec.Object
}

// MLStreamingInputDataSourceFromID constructs a [MLStreamingInputDataSource] from an objc.ID.
func MLStreamingInputDataSourceFromID(id objc.ID) MLStreamingInputDataSource {
	return MLStreamingInputDataSource{objectivec.Object{ID: id}}
}

// Ensure MLStreamingInputDataSource implements IMLStreamingInputDataSource.
var _ IMLStreamingInputDataSource = MLStreamingInputDataSource{}

// An interface definition for the [MLStreamingInputDataSource] class.
//
// # Methods
//
//   - [IMLStreamingInputDataSource.AppendBatchedTensorsNumberOfTensors]
//   - [IMLStreamingInputDataSource.BatchAtIndexError]
//   - [IMLStreamingInputDataSource.BatchSize]
//   - [IMLStreamingInputDataSource.SetBatchSize]
//   - [IMLStreamingInputDataSource.DataSources]
//   - [IMLStreamingInputDataSource.SetDataSources]
//   - [IMLStreamingInputDataSource.NumberOfBatches]
//   - [IMLStreamingInputDataSource.SizeOfBatchAtIndex]
//   - [IMLStreamingInputDataSource.InitWithBatchSize]
type IMLStreamingInputDataSource interface {
	objectivec.IObject

	// Topic: Methods

	AppendBatchedTensorsNumberOfTensors(tensors objectivec.IObject, tensors2 uint64)
	BatchAtIndexError(index uint64) (objectivec.IObject, error)
	BatchSize() uint64
	SetBatchSize(value uint64)
	DataSources() foundation.INSArray
	SetDataSources(value foundation.INSArray)
	NumberOfBatches() uint64
	SizeOfBatchAtIndex(index uint64) uint64
	InitWithBatchSize(size uint64) MLStreamingInputDataSource
}

// Init initializes the instance.
func (m MLStreamingInputDataSource) Init() MLStreamingInputDataSource {
	rv := objc.SendIfResponds[MLStreamingInputDataSource](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLStreamingInputDataSource) Autorelease() MLStreamingInputDataSource {
	rv := objc.SendIfResponds[MLStreamingInputDataSource](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLStreamingInputDataSource creates a new MLStreamingInputDataSource instance.
func NewMLStreamingInputDataSource() MLStreamingInputDataSource {
	class := getMLStreamingInputDataSourceClass()
	rv := objc.SendIfResponds[MLStreamingInputDataSource](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewStreamingInputDataSourceWithBatchSize(size uint64) MLStreamingInputDataSource {
	instance := getMLStreamingInputDataSourceClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithBatchSize:"), size)
	return MLStreamingInputDataSourceFromID(rv)
}

func (m MLStreamingInputDataSource) AppendBatchedTensorsNumberOfTensors(tensors objectivec.IObject, tensors2 uint64) {
	objc.SendIfResponds[objc.ID](m.ID, objc.Sel("appendBatchedTensors:numberOfTensors:"), tensors, tensors2)
}
func (m MLStreamingInputDataSource) BatchAtIndexError(index uint64) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("batchAtIndex:error:"), index, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (m MLStreamingInputDataSource) NumberOfBatches() uint64 {
	rv := objc.SendIfResponds[uint64](m.ID, objc.Sel("numberOfBatches"))
	return rv
}
func (m MLStreamingInputDataSource) SizeOfBatchAtIndex(index uint64) uint64 {
	rv := objc.SendIfResponds[uint64](m.ID, objc.Sel("sizeOfBatchAtIndex:"), index)
	return rv
}
func (m MLStreamingInputDataSource) InitWithBatchSize(size uint64) MLStreamingInputDataSource {
	rv := objc.SendIfResponds[MLStreamingInputDataSource](m.ID, objc.Sel("initWithBatchSize:"), size)
	return rv
}

func (m MLStreamingInputDataSource) BatchSize() uint64 {
	rv := objc.SendIfResponds[uint64](m.ID, objc.Sel("batchSize"))
	return rv
}
func (m MLStreamingInputDataSource) SetBatchSize(value uint64) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setBatchSize:"), value)
}
func (m MLStreamingInputDataSource) DataSources() foundation.INSArray {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("dataSources"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (m MLStreamingInputDataSource) SetDataSources(value foundation.INSArray) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setDataSources:"), value)
}
