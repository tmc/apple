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
//   - [MLNearestNeighborsSingleKdTreeIndex.ConstructTree]
//   - [MLNearestNeighborsSingleKdTreeIndex.ConstructTreeForPointsBoundedByStartingIndexCount]
//   - [MLNearestNeighborsSingleKdTreeIndex.DataPointCount]
//   - [MLNearestNeighborsSingleKdTreeIndex.EncodeWithCoder]
//   - [MLNearestNeighborsSingleKdTreeIndex.FindKNearestNeighborsToQueryPointInTree]
//   - [MLNearestNeighborsSingleKdTreeIndex.LeafSize]
//   - [MLNearestNeighborsSingleKdTreeIndex.SetLeafSize]
//   - [MLNearestNeighborsSingleKdTreeIndex.NumDimensions]
//   - [MLNearestNeighborsSingleKdTreeIndex.SetNumDimensions]
//   - [MLNearestNeighborsSingleKdTreeIndex.Root]
//   - [MLNearestNeighborsSingleKdTreeIndex.SetRoot]
//   - [MLNearestNeighborsSingleKdTreeIndex.UpdateWithDataError]
//   - [MLNearestNeighborsSingleKdTreeIndex.InitWithCoder]
//   - [MLNearestNeighborsSingleKdTreeIndex.InitWithDatasetNumberOfDimensionsLeafSizeError]
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
//   - [IMLNearestNeighborsSingleKdTreeIndex.ConstructTree]
//   - [IMLNearestNeighborsSingleKdTreeIndex.ConstructTreeForPointsBoundedByStartingIndexCount]
//   - [IMLNearestNeighborsSingleKdTreeIndex.DataPointCount]
//   - [IMLNearestNeighborsSingleKdTreeIndex.EncodeWithCoder]
//   - [IMLNearestNeighborsSingleKdTreeIndex.FindKNearestNeighborsToQueryPointInTree]
//   - [IMLNearestNeighborsSingleKdTreeIndex.LeafSize]
//   - [IMLNearestNeighborsSingleKdTreeIndex.SetLeafSize]
//   - [IMLNearestNeighborsSingleKdTreeIndex.NumDimensions]
//   - [IMLNearestNeighborsSingleKdTreeIndex.SetNumDimensions]
//   - [IMLNearestNeighborsSingleKdTreeIndex.Root]
//   - [IMLNearestNeighborsSingleKdTreeIndex.SetRoot]
//   - [IMLNearestNeighborsSingleKdTreeIndex.UpdateWithDataError]
//   - [IMLNearestNeighborsSingleKdTreeIndex.InitWithCoder]
//   - [IMLNearestNeighborsSingleKdTreeIndex.InitWithDatasetNumberOfDimensionsLeafSizeError]
type IMLNearestNeighborsSingleKdTreeIndex interface {
	objectivec.IObject

	// Topic: Methods

	ConstructTree() objectivec.IObject
	ConstructTreeForPointsBoundedByStartingIndexCount(by unsafe.Pointer, index uint64, count uint64) objectivec.IObject
	DataPointCount() uint64
	EncodeWithCoder(coder foundation.INSCoder)
	FindKNearestNeighborsToQueryPointInTree(k uint64, neighbors unsafe.Pointer, point unsafe.Pointer, tree objectivec.IObject)
	LeafSize() uint64
	SetLeafSize(value uint64)
	NumDimensions() uint64
	SetNumDimensions(value uint64)
	Root() IKDNode
	SetRoot(value IKDNode)
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

func NewNearestNeighborsSingleKdTreeIndexWithCoder(coder objectivec.IObject) MLNearestNeighborsSingleKdTreeIndex {
	instance := getMLNearestNeighborsSingleKdTreeIndexClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return MLNearestNeighborsSingleKdTreeIndexFromID(rv)
}

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

func (m MLNearestNeighborsSingleKdTreeIndex) ConstructTree() objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("constructTree"))
	return objectivec.Object{ID: rv}
}
func (m MLNearestNeighborsSingleKdTreeIndex) ConstructTreeForPointsBoundedByStartingIndexCount(by unsafe.Pointer, index uint64, count uint64) objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("constructTreeForPointsBoundedBy:startingIndex:count:"), objc.CArray(by), index, count)
	return objectivec.Object{ID: rv}
}
func (m MLNearestNeighborsSingleKdTreeIndex) DataPointCount() uint64 {
	rv := objc.Send[uint64](m.ID, objc.Sel("dataPointCount"))
	return rv
}
func (m MLNearestNeighborsSingleKdTreeIndex) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](m.ID, objc.Sel("encodeWithCoder:"), coder)
}
func (m MLNearestNeighborsSingleKdTreeIndex) FindKNearestNeighborsToQueryPointInTree(k uint64, neighbors unsafe.Pointer, point unsafe.Pointer, tree objectivec.IObject) {
	objc.Send[objc.ID](m.ID, objc.Sel("findK:nearestNeighbors:toQueryPoint:inTree:"), k, neighbors, point, tree)
}
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
func (m MLNearestNeighborsSingleKdTreeIndex) InitWithCoder(coder foundation.INSCoder) MLNearestNeighborsSingleKdTreeIndex {
	rv := objc.Send[MLNearestNeighborsSingleKdTreeIndex](m.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}
func (m MLNearestNeighborsSingleKdTreeIndex) InitWithDatasetNumberOfDimensionsLeafSizeError(dataset unsafe.Pointer, dimensions uint64, size uint64) (MLNearestNeighborsSingleKdTreeIndex, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("initWithDataset:numberOfDimensions:leafSize:error:"), dataset, dimensions, size, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLNearestNeighborsSingleKdTreeIndex{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLNearestNeighborsSingleKdTreeIndexFromID(rv), nil

}

func (_MLNearestNeighborsSingleKdTreeIndexClass MLNearestNeighborsSingleKdTreeIndexClass) SupportsSecureCoding() bool {
	rv := objc.Send[bool](objc.ID(_MLNearestNeighborsSingleKdTreeIndexClass.class), objc.Sel("supportsSecureCoding"))
	return rv
}

func (m MLNearestNeighborsSingleKdTreeIndex) LeafSize() uint64 {
	rv := objc.Send[uint64](m.ID, objc.Sel("leafSize"))
	return rv
}
func (m MLNearestNeighborsSingleKdTreeIndex) SetLeafSize(value uint64) {
	objc.Send[struct{}](m.ID, objc.Sel("setLeafSize:"), value)
}
func (m MLNearestNeighborsSingleKdTreeIndex) NumDimensions() uint64 {
	rv := objc.Send[uint64](m.ID, objc.Sel("numDimensions"))
	return rv
}
func (m MLNearestNeighborsSingleKdTreeIndex) SetNumDimensions(value uint64) {
	objc.Send[struct{}](m.ID, objc.Sel("setNumDimensions:"), value)
}
func (m MLNearestNeighborsSingleKdTreeIndex) Root() IKDNode {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("root"))
	return KDNodeFromID(objc.ID(rv))
}
func (m MLNearestNeighborsSingleKdTreeIndex) SetRoot(value IKDNode) {
	objc.Send[struct{}](m.ID, objc.Sel("setRoot:"), value)
}
