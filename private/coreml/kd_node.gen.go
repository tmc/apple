// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [KDNode] class.
var (
	_KDNodeClass     KDNodeClass
	_KDNodeClassOnce sync.Once
)

func getKDNodeClass() KDNodeClass {
	_KDNodeClassOnce.Do(func() {
		_KDNodeClass = KDNodeClass{class: objc.GetClass("_KDNode")}
	})
	return _KDNodeClass
}

// GetKDNodeClass returns the class object for _KDNode.
func GetKDNodeClass() KDNodeClass {
	return getKDNodeClass()
}

type KDNodeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (kc KDNodeClass) Class() objc.Class {
	return kc.class
}

// Alloc allocates memory for a new instance of the class.
func (kc KDNodeClass) Alloc() KDNode {
	rv := objc.Send[KDNode](objc.ID(kc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [KDNode.AssignSplitsForDataIndicesNumDimensions]
//   - [KDNode.BoundingBox]
//   - [KDNode.SetBoundingBox]
//   - [KDNode.Count]
//   - [KDNode.SetCount]
//   - [KDNode.EncodeWithCoder]
//   - [KDNode.FindMinAndMaxAlongDimensionDataIndicesNumDimensions]
//   - [KDNode.IsLeaf]
//   - [KDNode.SetIsLeaf]
//   - [KDNode.LeftChild]
//   - [KDNode.SetLeftChild]
//   - [KDNode.PartitionDataPointsIndicesNumDimensions]
//   - [KDNode.Print]
//   - [KDNode.RightChild]
//   - [KDNode.SetRightChild]
//   - [KDNode.SplitDimension]
//   - [KDNode.SetSplitDimension]
//   - [KDNode.SplitIndex]
//   - [KDNode.SetSplitIndex]
//   - [KDNode.SplitValue]
//   - [KDNode.SetSplitValue]
//   - [KDNode.StartingIndex]
//   - [KDNode.SetStartingIndex]
//   - [KDNode.InitWithCoder]
//
// See: https://developer.apple.com/documentation/CoreML/_KDNode
type KDNode struct {
	objectivec.Object
}

// KDNodeFromID constructs a [KDNode] from an objc.ID.
func KDNodeFromID(id objc.ID) KDNode {
	return KDNode{objectivec.Object{ID: id}}
}

// Ensure KDNode implements IKDNode.
var _ IKDNode = KDNode{}

// An interface definition for the [KDNode] class.
//
// # Methods
//
//   - [IKDNode.AssignSplitsForDataIndicesNumDimensions]
//   - [IKDNode.BoundingBox]
//   - [IKDNode.SetBoundingBox]
//   - [IKDNode.Count]
//   - [IKDNode.SetCount]
//   - [IKDNode.EncodeWithCoder]
//   - [IKDNode.FindMinAndMaxAlongDimensionDataIndicesNumDimensions]
//   - [IKDNode.IsLeaf]
//   - [IKDNode.SetIsLeaf]
//   - [IKDNode.LeftChild]
//   - [IKDNode.SetLeftChild]
//   - [IKDNode.PartitionDataPointsIndicesNumDimensions]
//   - [IKDNode.Print]
//   - [IKDNode.RightChild]
//   - [IKDNode.SetRightChild]
//   - [IKDNode.SplitDimension]
//   - [IKDNode.SetSplitDimension]
//   - [IKDNode.SplitIndex]
//   - [IKDNode.SetSplitIndex]
//   - [IKDNode.SplitValue]
//   - [IKDNode.SetSplitValue]
//   - [IKDNode.StartingIndex]
//   - [IKDNode.SetStartingIndex]
//   - [IKDNode.InitWithCoder]
//
// See: https://developer.apple.com/documentation/CoreML/_KDNode
type IKDNode interface {
	objectivec.IObject

	// Topic: Methods

	AssignSplitsForDataIndicesNumDimensions(data unsafe.Pointer, indices unsafe.Pointer, dimensions uint64)
	BoundingBox() objectivec.IObject
	SetBoundingBox(value objectivec.IObject)
	Count() uint64
	SetCount(value uint64)
	EncodeWithCoder(coder foundation.INSCoder)
	FindMinAndMaxAlongDimensionDataIndicesNumDimensions(min unsafe.Pointer, max unsafe.Pointer, dimension uint64, data unsafe.Pointer, indices unsafe.Pointer, dimensions uint64)
	IsLeaf() bool
	SetIsLeaf(value bool)
	LeftChild() *KDNode
	SetLeftChild(value *KDNode)
	PartitionDataPointsIndicesNumDimensions(points unsafe.Pointer, indices unsafe.Pointer, dimensions uint64)
	Print()
	RightChild() *KDNode
	SetRightChild(value *KDNode)
	SplitDimension() uint64
	SetSplitDimension(value uint64)
	SplitIndex() uint64
	SetSplitIndex(value uint64)
	SplitValue() float32
	SetSplitValue(value float32)
	StartingIndex() uint64
	SetStartingIndex(value uint64)
	InitWithCoder(coder foundation.INSCoder) KDNode
}

// Init initializes the instance.
func (k KDNode) Init() KDNode {
	rv := objc.Send[KDNode](k.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (k KDNode) Autorelease() KDNode {
	rv := objc.Send[KDNode](k.ID, objc.Sel("autorelease"))
	return rv
}

// NewKDNode creates a new KDNode instance.
func NewKDNode() KDNode {
	class := getKDNodeClass()
	rv := objc.Send[KDNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/_KDNode/initWithCoder:
func NewKDNodeWithCoder(coder objectivec.IObject) KDNode {
	instance := getKDNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return KDNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/_KDNode/assignSplitsForData:indices:numDimensions:
func (k KDNode) AssignSplitsForDataIndicesNumDimensions(data unsafe.Pointer, indices unsafe.Pointer, dimensions uint64) {
	objc.Send[objc.ID](k.ID, objc.Sel("assignSplitsForData:indices:numDimensions:"), data, indices, dimensions)
}

// See: https://developer.apple.com/documentation/CoreML/_KDNode/encodeWithCoder:
func (k KDNode) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](k.ID, objc.Sel("encodeWithCoder:"), coder)
}

// See: https://developer.apple.com/documentation/CoreML/_KDNode/findMin:andMax:alongDimension:data:indices:numDimensions:
func (k KDNode) FindMinAndMaxAlongDimensionDataIndicesNumDimensions(min unsafe.Pointer, max unsafe.Pointer, dimension uint64, data unsafe.Pointer, indices unsafe.Pointer, dimensions uint64) {
	objc.Send[objc.ID](k.ID, objc.Sel("findMin:andMax:alongDimension:data:indices:numDimensions:"), min, max, dimension, data, indices, dimensions)
}

// See: https://developer.apple.com/documentation/CoreML/_KDNode/partitionDataPoints:indices:numDimensions:
func (k KDNode) PartitionDataPointsIndicesNumDimensions(points unsafe.Pointer, indices unsafe.Pointer, dimensions uint64) {
	objc.Send[objc.ID](k.ID, objc.Sel("partitionDataPoints:indices:numDimensions:"), points, indices, dimensions)
}

// See: https://developer.apple.com/documentation/CoreML/_KDNode/print
func (k KDNode) Print() {
	objc.Send[objc.ID](k.ID, objc.Sel("print"))
}

// See: https://developer.apple.com/documentation/CoreML/_KDNode/initWithCoder:
func (k KDNode) InitWithCoder(coder foundation.INSCoder) KDNode {
	rv := objc.Send[KDNode](k.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/_KDNode/supportsSecureCoding
func (_KDNodeClass KDNodeClass) SupportsSecureCoding() bool {
	rv := objc.Send[bool](objc.ID(_KDNodeClass.class), objc.Sel("supportsSecureCoding"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/_KDNode/boundingBox
func (k KDNode) BoundingBox() objectivec.IObject {
	rv := objc.Send[objc.ID](k.ID, objc.Sel("boundingBox"))
	return objectivec.Object{ID: rv}
}
func (k KDNode) SetBoundingBox(value objectivec.IObject) {
	objc.Send[struct{}](k.ID, objc.Sel("setBoundingBox:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/_KDNode/count
func (k KDNode) Count() uint64 {
	rv := objc.Send[uint64](k.ID, objc.Sel("count"))
	return rv
}
func (k KDNode) SetCount(value uint64) {
	objc.Send[struct{}](k.ID, objc.Sel("setCount:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/_KDNode/isLeaf
func (k KDNode) IsLeaf() bool {
	rv := objc.Send[bool](k.ID, objc.Sel("isLeaf"))
	return rv
}
func (k KDNode) SetIsLeaf(value bool) {
	objc.Send[struct{}](k.ID, objc.Sel("setIsLeaf:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/_KDNode/leftChild
func (k KDNode) LeftChild() *KDNode {
	rv := objc.Send[objc.ID](k.ID, objc.Sel("leftChild"))
	if rv == 0 {
		return nil
	}
	val := KDNodeFromID(objc.ID(rv))
	return &val
}
func (k KDNode) SetLeftChild(value *KDNode) {
	if value == nil {
		objc.Send[struct{}](k.ID, objc.Sel("setLeftChild:"), objc.ID(0))
		return
	}
	objc.Send[struct{}](k.ID, objc.Sel("setLeftChild:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/_KDNode/rightChild
func (k KDNode) RightChild() *KDNode {
	rv := objc.Send[objc.ID](k.ID, objc.Sel("rightChild"))
	if rv == 0 {
		return nil
	}
	val := KDNodeFromID(objc.ID(rv))
	return &val
}
func (k KDNode) SetRightChild(value *KDNode) {
	if value == nil {
		objc.Send[struct{}](k.ID, objc.Sel("setRightChild:"), objc.ID(0))
		return
	}
	objc.Send[struct{}](k.ID, objc.Sel("setRightChild:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/_KDNode/splitDimension
func (k KDNode) SplitDimension() uint64 {
	rv := objc.Send[uint64](k.ID, objc.Sel("splitDimension"))
	return rv
}
func (k KDNode) SetSplitDimension(value uint64) {
	objc.Send[struct{}](k.ID, objc.Sel("setSplitDimension:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/_KDNode/splitIndex
func (k KDNode) SplitIndex() uint64 {
	rv := objc.Send[uint64](k.ID, objc.Sel("splitIndex"))
	return rv
}
func (k KDNode) SetSplitIndex(value uint64) {
	objc.Send[struct{}](k.ID, objc.Sel("setSplitIndex:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/_KDNode/splitValue
func (k KDNode) SplitValue() float32 {
	rv := objc.Send[float32](k.ID, objc.Sel("splitValue"))
	return rv
}
func (k KDNode) SetSplitValue(value float32) {
	objc.Send[struct{}](k.ID, objc.Sel("setSplitValue:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/_KDNode/startingIndex
func (k KDNode) StartingIndex() uint64 {
	rv := objc.Send[uint64](k.ID, objc.Sel("startingIndex"))
	return rv
}
func (k KDNode) SetStartingIndex(value uint64) {
	objc.Send[struct{}](k.ID, objc.Sel("setStartingIndex:"), value)
}
