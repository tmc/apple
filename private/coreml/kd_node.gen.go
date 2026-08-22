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
	rv := objc.SendIfResponds[KDNode](objc.ID(kc.class), objc.Sel("alloc"))
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
type IKDNode interface {
	objectivec.IObject

	// Topic: Methods

	AssignSplitsForDataIndicesNumDimensions(data *float32, indices *uint64, dimensions uint64)
	BoundingBox() KDBoundingBox
	SetBoundingBox(value KDBoundingBox)
	Count() uint64
	SetCount(value uint64)
	EncodeWithCoder(coder foundation.INSCoder)
	FindMinAndMaxAlongDimensionDataIndicesNumDimensions(min *float32, max *float32, dimension uint64, data *float32, indices *uint64, dimensions uint64)
	IsLeaf() bool
	SetIsLeaf(value bool)
	LeftChild() IKDNode
	SetLeftChild(value IKDNode)
	PartitionDataPointsIndicesNumDimensions(points *float32, indices *uint64, dimensions uint64)
	Print()
	RightChild() IKDNode
	SetRightChild(value IKDNode)
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
	rv := objc.SendIfResponds[KDNode](k.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (k KDNode) Autorelease() KDNode {
	rv := objc.SendIfResponds[KDNode](k.ID, objc.Sel("autorelease"))
	return rv
}

// NewKDNode creates a new KDNode instance.
func NewKDNode() KDNode {
	class := getKDNodeClass()
	rv := objc.SendIfResponds[KDNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewKDNodeWithCoder(coder objectivec.IObject) KDNode {
	instance := getKDNodeClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return KDNodeFromID(rv)
}

func (k KDNode) AssignSplitsForDataIndicesNumDimensions(data *float32, indices *uint64, dimensions uint64) {
	objc.SendIfResponds[objc.ID](k.ID, objc.Sel("assignSplitsForData:indices:numDimensions:"), data, unsafe.Pointer(indices), dimensions)
}
func (k KDNode) EncodeWithCoder(coder foundation.INSCoder) {
	objc.SendIfResponds[objc.ID](k.ID, objc.Sel("encodeWithCoder:"), coder)
}
func (k KDNode) FindMinAndMaxAlongDimensionDataIndicesNumDimensions(min *float32, max *float32, dimension uint64, data *float32, indices *uint64, dimensions uint64) {
	objc.SendIfResponds[objc.ID](k.ID, objc.Sel("findMin:andMax:alongDimension:data:indices:numDimensions:"), min, max, dimension, data, unsafe.Pointer(indices), dimensions)
}
func (k KDNode) PartitionDataPointsIndicesNumDimensions(points *float32, indices *uint64, dimensions uint64) {
	objc.SendIfResponds[objc.ID](k.ID, objc.Sel("partitionDataPoints:indices:numDimensions:"), points, unsafe.Pointer(indices), dimensions)
}
func (k KDNode) Print() {
	objc.SendIfResponds[objc.ID](k.ID, objc.Sel("print"))
}
func (k KDNode) InitWithCoder(coder foundation.INSCoder) KDNode {
	rv := objc.SendIfResponds[KDNode](k.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}

func (_KDNodeClass KDNodeClass) SupportsSecureCoding() bool {
	rv := objc.SendIfResponds[bool](objc.ID(_KDNodeClass.class), objc.Sel("supportsSecureCoding"))
	return rv
}

func (k KDNode) BoundingBox() KDBoundingBox {
	rv := objc.SendIfResponds[KDBoundingBox](k.ID, objc.Sel("boundingBox"))
	return KDBoundingBox(rv)
}
func (k KDNode) SetBoundingBox(value KDBoundingBox) {
	objc.SendIfResponds[struct{}](k.ID, objc.Sel("setBoundingBox:"), value)
}
func (k KDNode) Count() uint64 {
	rv := objc.SendIfResponds[uint64](k.ID, objc.Sel("count"))
	return rv
}
func (k KDNode) SetCount(value uint64) {
	objc.SendIfResponds[struct{}](k.ID, objc.Sel("setCount:"), value)
}
func (k KDNode) IsLeaf() bool {
	rv := objc.SendIfResponds[bool](k.ID, objc.Sel("isLeaf"))
	return rv
}
func (k KDNode) SetIsLeaf(value bool) {
	objc.SendIfResponds[struct{}](k.ID, objc.Sel("setIsLeaf:"), value)
}
func (k KDNode) LeftChild() IKDNode {
	rv := objc.SendIfResponds[objc.ID](k.ID, objc.Sel("leftChild"))
	return KDNodeFromID(objc.ID(rv))
}
func (k KDNode) SetLeftChild(value IKDNode) {
	objc.SendIfResponds[struct{}](k.ID, objc.Sel("setLeftChild:"), value)
}
func (k KDNode) RightChild() IKDNode {
	rv := objc.SendIfResponds[objc.ID](k.ID, objc.Sel("rightChild"))
	return KDNodeFromID(objc.ID(rv))
}
func (k KDNode) SetRightChild(value IKDNode) {
	objc.SendIfResponds[struct{}](k.ID, objc.Sel("setRightChild:"), value)
}
func (k KDNode) SplitDimension() uint64 {
	rv := objc.SendIfResponds[uint64](k.ID, objc.Sel("splitDimension"))
	return rv
}
func (k KDNode) SetSplitDimension(value uint64) {
	objc.SendIfResponds[struct{}](k.ID, objc.Sel("setSplitDimension:"), value)
}
func (k KDNode) SplitIndex() uint64 {
	rv := objc.SendIfResponds[uint64](k.ID, objc.Sel("splitIndex"))
	return rv
}
func (k KDNode) SetSplitIndex(value uint64) {
	objc.SendIfResponds[struct{}](k.ID, objc.Sel("setSplitIndex:"), value)
}
func (k KDNode) SplitValue() float32 {
	rv := objc.SendIfResponds[float32](k.ID, objc.Sel("splitValue"))
	return rv
}
func (k KDNode) SetSplitValue(value float32) {
	objc.SendIfResponds[struct{}](k.ID, objc.Sel("setSplitValue:"), value)
}
func (k KDNode) StartingIndex() uint64 {
	rv := objc.SendIfResponds[uint64](k.ID, objc.Sel("startingIndex"))
	return rv
}
func (k KDNode) SetStartingIndex(value uint64) {
	objc.SendIfResponds[struct{}](k.ID, objc.Sel("setStartingIndex:"), value)
}
