// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"errors"
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
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

// # Methods
//
//   - [MLNearestNeighborsLinearIndex.DataPointCount]
//   - [MLNearestNeighborsLinearIndex.EncodeWithCoder]
//   - [MLNearestNeighborsLinearIndex.NumDataPoints]
//   - [MLNearestNeighborsLinearIndex.SetNumDataPoints]
//   - [MLNearestNeighborsLinearIndex.NumDimensions]
//   - [MLNearestNeighborsLinearIndex.SetNumDimensions]
//   - [MLNearestNeighborsLinearIndex.UpdateWithDataError]
//   - [MLNearestNeighborsLinearIndex.InitWithCoder]
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
//   - [IMLNearestNeighborsLinearIndex.NumDataPoints]
//   - [IMLNearestNeighborsLinearIndex.SetNumDataPoints]
//   - [IMLNearestNeighborsLinearIndex.NumDimensions]
//   - [IMLNearestNeighborsLinearIndex.SetNumDimensions]
//   - [IMLNearestNeighborsLinearIndex.UpdateWithDataError]
//   - [IMLNearestNeighborsLinearIndex.InitWithCoder]
type IMLNearestNeighborsLinearIndex interface {
	objectivec.IObject

	// Topic: Methods

	DataPointCount() uint64
	EncodeWithCoder(coder foundation.INSCoder)
	NumDataPoints() uint64
	SetNumDataPoints(value uint64)
	NumDimensions() uint64
	SetNumDimensions(value uint64)
	UpdateWithDataError(data unsafe.Pointer) (bool, error)
	InitWithCoder(coder foundation.INSCoder) MLNearestNeighborsLinearIndex
}

// Init initializes the instance.
func (m MLNearestNeighborsLinearIndex) Init() MLNearestNeighborsLinearIndex {
	rv := objc.Send[MLNearestNeighborsLinearIndex](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLNearestNeighborsLinearIndex) Autorelease() MLNearestNeighborsLinearIndex {
	rv := objc.Send[MLNearestNeighborsLinearIndex](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLNearestNeighborsLinearIndex creates a new MLNearestNeighborsLinearIndex instance.
func NewMLNearestNeighborsLinearIndex() MLNearestNeighborsLinearIndex {
	class := getMLNearestNeighborsLinearIndexClass()
	rv := objc.Send[MLNearestNeighborsLinearIndex](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewNearestNeighborsLinearIndexWithCoder(coder objectivec.IObject) MLNearestNeighborsLinearIndex {
	instance := getMLNearestNeighborsLinearIndexClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return MLNearestNeighborsLinearIndexFromID(rv)
}

func NewNearestNeighborsLinearIndexWithDatasetNumberOfDimensions(dataset unsafe.Pointer, dimensions uint64) MLNearestNeighborsLinearIndex {
	instance := getMLNearestNeighborsLinearIndexClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDataset:numberOfDimensions:"), dataset, dimensions)
	return MLNearestNeighborsLinearIndexFromID(rv)
}

func (m MLNearestNeighborsLinearIndex) DataPointCount() uint64 {
	rv := objc.Send[uint64](m.ID, objc.Sel("dataPointCount"))
	return rv
}
func (m MLNearestNeighborsLinearIndex) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](m.ID, objc.Sel("encodeWithCoder:"), coder)
}
func (m MLNearestNeighborsLinearIndex) UpdateWithDataError(data unsafe.Pointer) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](m.ID, objc.Sel("updateWithData:error:"), data, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("updateWithData:error: returned NO with nil NSError")
	}
	return rv, nil

}
func (m MLNearestNeighborsLinearIndex) InitWithCoder(coder foundation.INSCoder) MLNearestNeighborsLinearIndex {
	rv := objc.Send[MLNearestNeighborsLinearIndex](m.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}

func (_MLNearestNeighborsLinearIndexClass MLNearestNeighborsLinearIndexClass) SupportsSecureCoding() bool {
	rv := objc.Send[bool](objc.ID(_MLNearestNeighborsLinearIndexClass.class), objc.Sel("supportsSecureCoding"))
	return rv
}

func (m MLNearestNeighborsLinearIndex) NumDataPoints() uint64 {
	rv := objc.Send[uint64](m.ID, objc.Sel("numDataPoints"))
	return rv
}
func (m MLNearestNeighborsLinearIndex) SetNumDataPoints(value uint64) {
	objc.Send[struct{}](m.ID, objc.Sel("setNumDataPoints:"), value)
}
func (m MLNearestNeighborsLinearIndex) NumDimensions() uint64 {
	rv := objc.Send[uint64](m.ID, objc.Sel("numDimensions"))
	return rv
}
func (m MLNearestNeighborsLinearIndex) SetNumDimensions(value uint64) {
	objc.Send[struct{}](m.ID, objc.Sel("setNumDimensions:"), value)
}
