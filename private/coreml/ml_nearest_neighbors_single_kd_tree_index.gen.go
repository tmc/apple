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

// The class instance for the [MLNearestNeighborsSingleKdTreeIndex] class.
var (
	_MLNearestNeighborsSingleKdTreeIndexClass     MLNearestNeighborsSingleKdTreeIndexClass
	_MLNearestNeighborsSingleKdTreeIndexClassOnce sync.Once
)

func getMLNearestNeighborsSingleKdTreeIndexClass() MLNearestNeighborsSingleKdTreeIndexClass {
	_MLNearestNeighborsSingleKdTreeIndexClassOnce.Do(func() {
		_MLNearestNeighborsSingleKdTreeIndexClass = MLNearestNeighborsSingleKdTreeIndexClass{class: objc.GetClass("MLNearestNeighborsSingleKdTreeIndex")}
	})
	return _MLNearestNeighborsSingleKdTreeIndexClass
}

// GetMLNearestNeighborsSingleKdTreeIndexClass returns the class object for MLNearestNeighborsSingleKdTreeIndex.
func GetMLNearestNeighborsSingleKdTreeIndexClass() MLNearestNeighborsSingleKdTreeIndexClass {
	return getMLNearestNeighborsSingleKdTreeIndexClass()
}

type MLNearestNeighborsSingleKdTreeIndexClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLNearestNeighborsSingleKdTreeIndexClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLNearestNeighborsSingleKdTreeIndexClass) Alloc() MLNearestNeighborsSingleKdTreeIndex {
	rv := objc.Send[MLNearestNeighborsSingleKdTreeIndex](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLNearestNeighborsSingleKdTreeIndex.CalculateDistancesForNodesBetweenLeftAndRightToQueryPoint]
//   - [MLNearestNeighborsSingleKdTreeIndex.ConstructTree]
//   - [MLNearestNeighborsSingleKdTreeIndex.ConstructTreeForPointsBoundedByStartingIndexCount]
//   - [MLNearestNeighborsSingleKdTreeIndex.DataPointCount]
//   - [MLNearestNeighborsSingleKdTreeIndex.EncodeWithCoder]
//   - [MLNearestNeighborsSingleKdTreeIndex.FindKNearestNeighborsToQueryPointInTree]
//   - [MLNearestNeighborsSingleKdTreeIndex.FindNearestNeighborsToIndex]
//   - [MLNearestNeighborsSingleKdTreeIndex.FindNearestNeighborsToQueryPoint]
//   - [MLNearestNeighborsSingleKdTreeIndex.LeafSize]
//   - [MLNearestNeighborsSingleKdTreeIndex.SetLeafSize]
//   - [MLNearestNeighborsSingleKdTreeIndex.NumDimensions]
//   - [MLNearestNeighborsSingleKdTreeIndex.SetNumDimensions]
//   - [MLNearestNeighborsSingleKdTreeIndex.Root]
//   - [MLNearestNeighborsSingleKdTreeIndex.SetRoot]
//   - [MLNearestNeighborsSingleKdTreeIndex.UpdateWithDataError]
//   - [MLNearestNeighborsSingleKdTreeIndex.InitWithCoder]
//   - [MLNearestNeighborsSingleKdTreeIndex.InitWithDatasetNumberOfDimensionsLeafSizeError]
//
// See: https://developer.apple.com/documentation/CoreML/MLNearestNeighborsSingleKdTreeIndex
type MLNearestNeighborsSingleKdTreeIndex struct {
	objectivec.Object
}

// MLNearestNeighborsSingleKdTreeIndexFromID constructs a [MLNearestNeighborsSingleKdTreeIndex] from an objc.ID.
func MLNearestNeighborsSingleKdTreeIndexFromID(id objc.ID) MLNearestNeighborsSingleKdTreeIndex {
	return MLNearestNeighborsSingleKdTreeIndex{objectivec.Object{ID: id}}
}

// Ensure MLNearestNeighborsSingleKdTreeIndex implements IMLNearestNeighborsSingleKdTreeIndex.
var _ IMLNearestNeighborsSingleKdTreeIndex = MLNearestNeighborsSingleKdTreeIndex{}

// An interface definition for the [MLNearestNeighborsSingleKdTreeIndex] class.
//
// # Methods
//
//   - [IMLNearestNeighborsSingleKdTreeIndex.CalculateDistancesForNodesBetweenLeftAndRightToQueryPoint]
//   - [IMLNearestNeighborsSingleKdTreeIndex.ConstructTree]
//   - [IMLNearestNeighborsSingleKdTreeIndex.ConstructTreeForPointsBoundedByStartingIndexCount]
//   - [IMLNearestNeighborsSingleKdTreeIndex.DataPointCount]
//   - [IMLNearestNeighborsSingleKdTreeIndex.EncodeWithCoder]
//   - [IMLNearestNeighborsSingleKdTreeIndex.FindKNearestNeighborsToQueryPointInTree]
//   - [IMLNearestNeighborsSingleKdTreeIndex.FindNearestNeighborsToIndex]
//   - [IMLNearestNeighborsSingleKdTreeIndex.FindNearestNeighborsToQueryPoint]
//   - [IMLNearestNeighborsSingleKdTreeIndex.LeafSize]
//   - [IMLNearestNeighborsSingleKdTreeIndex.SetLeafSize]
//   - [IMLNearestNeighborsSingleKdTreeIndex.NumDimensions]
//   - [IMLNearestNeighborsSingleKdTreeIndex.SetNumDimensions]
//   - [IMLNearestNeighborsSingleKdTreeIndex.Root]
//   - [IMLNearestNeighborsSingleKdTreeIndex.SetRoot]
//   - [IMLNearestNeighborsSingleKdTreeIndex.UpdateWithDataError]
//   - [IMLNearestNeighborsSingleKdTreeIndex.InitWithCoder]
//   - [IMLNearestNeighborsSingleKdTreeIndex.InitWithDatasetNumberOfDimensionsLeafSizeError]
//
// See: https://developer.apple.com/documentation/CoreML/MLNearestNeighborsSingleKdTreeIndex
type IMLNearestNeighborsSingleKdTreeIndex interface {
	objectivec.IObject

	// Topic: Methods

	CalculateDistancesForNodesBetweenLeftAndRightToQueryPoint(left uint64, right uint64, point unsafe.Pointer) objectivec.IObject
	ConstructTree() objectivec.IObject
	ConstructTreeForPointsBoundedByStartingIndexCount(by unsafe.Pointer, index uint64, count uint64) objectivec.IObject
	DataPointCount() uint64
	EncodeWithCoder(coder foundation.INSCoder)
	FindKNearestNeighborsToQueryPointInTree(k uint64, neighbors unsafe.Pointer, point unsafe.Pointer, tree objectivec.IObject)
	FindNearestNeighborsToIndex(neighbors uint64, index uint64) objectivec.IObject
	FindNearestNeighborsToQueryPoint(neighbors uint64, point unsafe.Pointer) objectivec.IObject
	LeafSize() uint64
	SetLeafSize(value uint64)
	NumDimensions() uint64
	SetNumDimensions(value uint64)
	Root() *KDNode
	SetRoot(value *KDNode)
	UpdateWithDataError(data unsafe.Pointer) (bool, error)
	InitWithCoder(coder foundation.INSCoder) MLNearestNeighborsSingleKdTreeIndex
	InitWithDatasetNumberOfDimensionsLeafSizeError(dataset unsafe.Pointer, dimensions uint64, size uint64) (MLNearestNeighborsSingleKdTreeIndex, error)
}

// Init initializes the instance.
func (m MLNearestNeighborsSingleKdTreeIndex) Init() MLNearestNeighborsSingleKdTreeIndex {
	rv := objc.Send[MLNearestNeighborsSingleKdTreeIndex](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLNearestNeighborsSingleKdTreeIndex) Autorelease() MLNearestNeighborsSingleKdTreeIndex {
	rv := objc.Send[MLNearestNeighborsSingleKdTreeIndex](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLNearestNeighborsSingleKdTreeIndex creates a new MLNearestNeighborsSingleKdTreeIndex instance.
func NewMLNearestNeighborsSingleKdTreeIndex() MLNearestNeighborsSingleKdTreeIndex {
	class := getMLNearestNeighborsSingleKdTreeIndexClass()
	rv := objc.Send[MLNearestNeighborsSingleKdTreeIndex](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLNearestNeighborsSingleKdTreeIndex/initWithCoder:
func NewNearestNeighborsSingleKdTreeIndexWithCoder(coder objectivec.IObject) MLNearestNeighborsSingleKdTreeIndex {
	instance := getMLNearestNeighborsSingleKdTreeIndexClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return MLNearestNeighborsSingleKdTreeIndexFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLNearestNeighborsSingleKdTreeIndex/initWithDataset:numberOfDimensions:leafSize:error:
func NewNearestNeighborsSingleKdTreeIndexWithDatasetNumberOfDimensionsLeafSizeError(dataset unsafe.Pointer, dimensions uint64, size uint64) (MLNearestNeighborsSingleKdTreeIndex, error) {
	var errorPtr objc.ID
	instance := getMLNearestNeighborsSingleKdTreeIndexClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDataset:numberOfDimensions:leafSize:error:"), dataset, dimensions, size, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLNearestNeighborsSingleKdTreeIndex{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLNearestNeighborsSingleKdTreeIndexFromID(rv), nil
}

// See: https://developer.apple.com/documentation/CoreML/MLNearestNeighborsSingleKdTreeIndex/calculateDistancesForNodesBetweenLeft:andRight:toQueryPoint:
func (m MLNearestNeighborsSingleKdTreeIndex) CalculateDistancesForNodesBetweenLeftAndRightToQueryPoint(left uint64, right uint64, point unsafe.Pointer) objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("calculateDistancesForNodesBetweenLeft:andRight:toQueryPoint:"), left, right, point)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/CoreML/MLNearestNeighborsSingleKdTreeIndex/constructTree
func (m MLNearestNeighborsSingleKdTreeIndex) ConstructTree() objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("constructTree"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/CoreML/MLNearestNeighborsSingleKdTreeIndex/constructTreeForPointsBoundedBy:startingIndex:count:
func (m MLNearestNeighborsSingleKdTreeIndex) ConstructTreeForPointsBoundedByStartingIndexCount(by unsafe.Pointer, index uint64, count uint64) objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("constructTreeForPointsBoundedBy:startingIndex:count:"), objc.CArray(by), index, count)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/CoreML/MLNearestNeighborsSingleKdTreeIndex/dataPointCount
func (m MLNearestNeighborsSingleKdTreeIndex) DataPointCount() uint64 {
	rv := objc.Send[uint64](m.ID, objc.Sel("dataPointCount"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLNearestNeighborsSingleKdTreeIndex/encodeWithCoder:
func (m MLNearestNeighborsSingleKdTreeIndex) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](m.ID, objc.Sel("encodeWithCoder:"), coder)
}

// See: https://developer.apple.com/documentation/CoreML/MLNearestNeighborsSingleKdTreeIndex/findK:nearestNeighbors:toQueryPoint:inTree:
func (m MLNearestNeighborsSingleKdTreeIndex) FindKNearestNeighborsToQueryPointInTree(k uint64, neighbors unsafe.Pointer, point unsafe.Pointer, tree objectivec.IObject) {
	objc.Send[objc.ID](m.ID, objc.Sel("findK:nearestNeighbors:toQueryPoint:inTree:"), k, neighbors, point, tree)
}

// See: https://developer.apple.com/documentation/CoreML/MLNearestNeighborsSingleKdTreeIndex/findNearestNeighbors:toIndex:
func (m MLNearestNeighborsSingleKdTreeIndex) FindNearestNeighborsToIndex(neighbors uint64, index uint64) objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("findNearestNeighbors:toIndex:"), neighbors, index)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/CoreML/MLNearestNeighborsSingleKdTreeIndex/findNearestNeighbors:toQueryPoint:
func (m MLNearestNeighborsSingleKdTreeIndex) FindNearestNeighborsToQueryPoint(neighbors uint64, point unsafe.Pointer) objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("findNearestNeighbors:toQueryPoint:"), neighbors, point)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/CoreML/MLNearestNeighborsSingleKdTreeIndex/updateWithData:error:
func (m MLNearestNeighborsSingleKdTreeIndex) UpdateWithDataError(data unsafe.Pointer) (bool, error) {
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

// See: https://developer.apple.com/documentation/CoreML/MLNearestNeighborsSingleKdTreeIndex/initWithCoder:
func (m MLNearestNeighborsSingleKdTreeIndex) InitWithCoder(coder foundation.INSCoder) MLNearestNeighborsSingleKdTreeIndex {
	rv := objc.Send[MLNearestNeighborsSingleKdTreeIndex](m.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLNearestNeighborsSingleKdTreeIndex/initWithDataset:numberOfDimensions:leafSize:error:
func (m MLNearestNeighborsSingleKdTreeIndex) InitWithDatasetNumberOfDimensionsLeafSizeError(dataset unsafe.Pointer, dimensions uint64, size uint64) (MLNearestNeighborsSingleKdTreeIndex, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("initWithDataset:numberOfDimensions:leafSize:error:"), dataset, dimensions, size, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLNearestNeighborsSingleKdTreeIndex{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLNearestNeighborsSingleKdTreeIndexFromID(rv), nil

}

// See: https://developer.apple.com/documentation/CoreML/MLNearestNeighborsSingleKdTreeIndex/supportsSecureCoding
func (_MLNearestNeighborsSingleKdTreeIndexClass MLNearestNeighborsSingleKdTreeIndexClass) SupportsSecureCoding() bool {
	rv := objc.Send[bool](objc.ID(_MLNearestNeighborsSingleKdTreeIndexClass.class), objc.Sel("supportsSecureCoding"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLNearestNeighborsSingleKdTreeIndex/leafSize
func (m MLNearestNeighborsSingleKdTreeIndex) LeafSize() uint64 {
	rv := objc.Send[uint64](m.ID, objc.Sel("leafSize"))
	return rv
}
func (m MLNearestNeighborsSingleKdTreeIndex) SetLeafSize(value uint64) {
	objc.Send[struct{}](m.ID, objc.Sel("setLeafSize:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLNearestNeighborsSingleKdTreeIndex/numDimensions
func (m MLNearestNeighborsSingleKdTreeIndex) NumDimensions() uint64 {
	rv := objc.Send[uint64](m.ID, objc.Sel("numDimensions"))
	return rv
}
func (m MLNearestNeighborsSingleKdTreeIndex) SetNumDimensions(value uint64) {
	objc.Send[struct{}](m.ID, objc.Sel("setNumDimensions:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLNearestNeighborsSingleKdTreeIndex/root
func (m MLNearestNeighborsSingleKdTreeIndex) Root() *KDNode {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("root"))
	if rv == 0 {
		return nil
	}
	val := KDNodeFromID(objc.ID(rv))
	return &val
}
func (m MLNearestNeighborsSingleKdTreeIndex) SetRoot(value *KDNode) {
	if value == nil {
		objc.Send[struct{}](m.ID, objc.Sel("setRoot:"), objc.ID(0))
		return
	}
	objc.Send[struct{}](m.ID, objc.Sel("setRoot:"), value)
}
