// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"errors"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLNearestNeighborsLinearIndex] class.
var (
	_MLNearestNeighborsLinearIndexClass     MLNearestNeighborsLinearIndexClass
	_MLNearestNeighborsLinearIndexClassOnce sync.Once
)

func getMLNearestNeighborsLinearIndexClass() MLNearestNeighborsLinearIndexClass {
	_MLNearestNeighborsLinearIndexClassOnce.Do(func() {
		_MLNearestNeighborsLinearIndexClass = MLNearestNeighborsLinearIndexClass{class: objc.GetClass("MLNearestNeighborsLinearIndex")}
	})
	return _MLNearestNeighborsLinearIndexClass
}

// GetMLNearestNeighborsLinearIndexClass returns the class object for MLNearestNeighborsLinearIndex.
func GetMLNearestNeighborsLinearIndexClass() MLNearestNeighborsLinearIndexClass {
	return getMLNearestNeighborsLinearIndexClass()
}

type MLNearestNeighborsLinearIndexClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLNearestNeighborsLinearIndexClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLNearestNeighborsLinearIndexClass) Alloc() MLNearestNeighborsLinearIndex {
	rv := objc.Send[MLNearestNeighborsLinearIndex](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [MLNearestNeighborsLinearIndex.DataPointCount]
//   - [MLNearestNeighborsLinearIndex.EncodeWithCoder]
//   - [MLNearestNeighborsLinearIndex.FindNearestNeighborsToIndex]
//   - [MLNearestNeighborsLinearIndex.FindNearestNeighborsToQueryPoint]
//   - [MLNearestNeighborsLinearIndex.NumDataPoints]
//   - [MLNearestNeighborsLinearIndex.SetNumDataPoints]
//   - [MLNearestNeighborsLinearIndex.NumDimensions]
//   - [MLNearestNeighborsLinearIndex.SetNumDimensions]
//   - [MLNearestNeighborsLinearIndex.UpdateWithDataError]
//   - [MLNearestNeighborsLinearIndex.InitWithCoder]
//   - [MLNearestNeighborsLinearIndex.InitWithDatasetNumberOfDimensions]
// See: https://developer.apple.com/documentation/CoreML/MLNearestNeighborsLinearIndex
type MLNearestNeighborsLinearIndex struct {
	objectivec.Object
}

// MLNearestNeighborsLinearIndexFromID constructs a [MLNearestNeighborsLinearIndex] from an objc.ID.
func MLNearestNeighborsLinearIndexFromID(id objc.ID) MLNearestNeighborsLinearIndex {
	return MLNearestNeighborsLinearIndex{objectivec.Object{ID: id}}
}
// Ensure MLNearestNeighborsLinearIndex implements IMLNearestNeighborsLinearIndex.
var _ IMLNearestNeighborsLinearIndex = MLNearestNeighborsLinearIndex{}

// An interface definition for the [MLNearestNeighborsLinearIndex] class.
//
// # Methods
//
//   - [IMLNearestNeighborsLinearIndex.DataPointCount]
//   - [IMLNearestNeighborsLinearIndex.EncodeWithCoder]
//   - [IMLNearestNeighborsLinearIndex.FindNearestNeighborsToIndex]
//   - [IMLNearestNeighborsLinearIndex.FindNearestNeighborsToQueryPoint]
//   - [IMLNearestNeighborsLinearIndex.NumDataPoints]
//   - [IMLNearestNeighborsLinearIndex.SetNumDataPoints]
//   - [IMLNearestNeighborsLinearIndex.NumDimensions]
//   - [IMLNearestNeighborsLinearIndex.SetNumDimensions]
//   - [IMLNearestNeighborsLinearIndex.UpdateWithDataError]
//   - [IMLNearestNeighborsLinearIndex.InitWithCoder]
//   - [IMLNearestNeighborsLinearIndex.InitWithDatasetNumberOfDimensions]
//
// See: https://developer.apple.com/documentation/CoreML/MLNearestNeighborsLinearIndex
type IMLNearestNeighborsLinearIndex interface {
	objectivec.IObject

	// Topic: Methods

	DataPointCount() uint64
	EncodeWithCoder(coder foundation.INSCoder)
	FindNearestNeighborsToIndex(neighbors uint64, index uint64) objectivec.IObject
	FindNearestNeighborsToQueryPoint(neighbors uint64, point unsafe.Pointer) objectivec.IObject
	NumDataPoints() uint64
	SetNumDataPoints(value uint64)
	NumDimensions() uint64
	SetNumDimensions(value uint64)
	UpdateWithDataError(data unsafe.Pointer) (bool, error)
	InitWithCoder(coder foundation.INSCoder) MLNearestNeighborsLinearIndex
	InitWithDatasetNumberOfDimensions(dataset objectivec.IObject, dimensions uint64) MLNearestNeighborsLinearIndex
}

// Init initializes the instance.
func (n MLNearestNeighborsLinearIndex) Init() MLNearestNeighborsLinearIndex {
	rv := objc.Send[MLNearestNeighborsLinearIndex](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n MLNearestNeighborsLinearIndex) Autorelease() MLNearestNeighborsLinearIndex {
	rv := objc.Send[MLNearestNeighborsLinearIndex](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLNearestNeighborsLinearIndex creates a new MLNearestNeighborsLinearIndex instance.
func NewMLNearestNeighborsLinearIndex() MLNearestNeighborsLinearIndex {
	class := getMLNearestNeighborsLinearIndexClass()
	rv := objc.Send[MLNearestNeighborsLinearIndex](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/CoreML/MLNearestNeighborsLinearIndex/initWithCoder:
func NewNearestNeighborsLinearIndexWithCoder(coder objectivec.IObject) MLNearestNeighborsLinearIndex {
	instance := getMLNearestNeighborsLinearIndexClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return MLNearestNeighborsLinearIndexFromID(rv)
}

//
// See: https://developer.apple.com/documentation/CoreML/MLNearestNeighborsLinearIndex/initWithDataset:numberOfDimensions:
func NewNearestNeighborsLinearIndexWithDatasetNumberOfDimensions(dataset objectivec.IObject, dimensions uint64) MLNearestNeighborsLinearIndex {
	instance := getMLNearestNeighborsLinearIndexClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDataset:numberOfDimensions:"), dataset, dimensions)
	return MLNearestNeighborsLinearIndexFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLNearestNeighborsLinearIndex/dataPointCount
func (n MLNearestNeighborsLinearIndex) DataPointCount() uint64 {
	rv := objc.Send[uint64](n.ID, objc.Sel("dataPointCount"))
	return rv
}
//
// See: https://developer.apple.com/documentation/CoreML/MLNearestNeighborsLinearIndex/encodeWithCoder:
func (n MLNearestNeighborsLinearIndex) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](n.ID, objc.Sel("encodeWithCoder:"), coder)
}
//
// See: https://developer.apple.com/documentation/CoreML/MLNearestNeighborsLinearIndex/findNearestNeighbors:toIndex:
func (n MLNearestNeighborsLinearIndex) FindNearestNeighborsToIndex(neighbors uint64, index uint64) objectivec.IObject {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("findNearestNeighbors:toIndex:"), neighbors, index)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/CoreML/MLNearestNeighborsLinearIndex/findNearestNeighbors:toQueryPoint:
func (n MLNearestNeighborsLinearIndex) FindNearestNeighborsToQueryPoint(neighbors uint64, point unsafe.Pointer) objectivec.IObject {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("findNearestNeighbors:toQueryPoint:"), neighbors, point)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/CoreML/MLNearestNeighborsLinearIndex/updateWithData:error:
func (n MLNearestNeighborsLinearIndex) UpdateWithDataError(data unsafe.Pointer) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](n.ID, objc.Sel("updateWithData:error:"), data, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("updateWithData:error: returned NO with nil NSError")
	}
	return rv, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLNearestNeighborsLinearIndex/initWithCoder:
func (n MLNearestNeighborsLinearIndex) InitWithCoder(coder foundation.INSCoder) MLNearestNeighborsLinearIndex {
	rv := objc.Send[MLNearestNeighborsLinearIndex](n.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}
//
// See: https://developer.apple.com/documentation/CoreML/MLNearestNeighborsLinearIndex/initWithDataset:numberOfDimensions:
func (n MLNearestNeighborsLinearIndex) InitWithDatasetNumberOfDimensions(dataset objectivec.IObject, dimensions uint64) MLNearestNeighborsLinearIndex {
	rv := objc.Send[MLNearestNeighborsLinearIndex](n.ID, objc.Sel("initWithDataset:numberOfDimensions:"), dataset, dimensions)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLNearestNeighborsLinearIndex/supportsSecureCoding
func (_MLNearestNeighborsLinearIndexClass MLNearestNeighborsLinearIndexClass) SupportsSecureCoding() bool {
	rv := objc.Send[bool](objc.ID(_MLNearestNeighborsLinearIndexClass.class), objc.Sel("supportsSecureCoding"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLNearestNeighborsLinearIndex/numDataPoints
func (n MLNearestNeighborsLinearIndex) NumDataPoints() uint64 {
	rv := objc.Send[uint64](n.ID, objc.Sel("numDataPoints"))
	return rv
}
func (n MLNearestNeighborsLinearIndex) SetNumDataPoints(value uint64) {
	objc.Send[struct{}](n.ID, objc.Sel("setNumDataPoints:"), value)
}
// See: https://developer.apple.com/documentation/CoreML/MLNearestNeighborsLinearIndex/numDimensions
func (n MLNearestNeighborsLinearIndex) NumDimensions() uint64 {
	rv := objc.Send[uint64](n.ID, objc.Sel("numDimensions"))
	return rv
}
func (n MLNearestNeighborsLinearIndex) SetNumDimensions(value uint64) {
	objc.Send[struct{}](n.ID, objc.Sel("setNumDimensions:"), value)
}

