// Code generated from Apple documentation for gtshaderprofiler. DO NOT EDIT.

package gtshaderprofiler

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [GTMioShaderExecutionHistoryLoopNode] class.
var (
	_GTMioShaderExecutionHistoryLoopNodeClass     GTMioShaderExecutionHistoryLoopNodeClass
	_GTMioShaderExecutionHistoryLoopNodeClassOnce sync.Once
)

func getGTMioShaderExecutionHistoryLoopNodeClass() GTMioShaderExecutionHistoryLoopNodeClass {
	_GTMioShaderExecutionHistoryLoopNodeClassOnce.Do(func() {
		_GTMioShaderExecutionHistoryLoopNodeClass = GTMioShaderExecutionHistoryLoopNodeClass{class: objc.GetClass("GTMioShaderExecutionHistoryLoopNode")}
	})
	return _GTMioShaderExecutionHistoryLoopNodeClass
}

// GetGTMioShaderExecutionHistoryLoopNodeClass returns the class object for GTMioShaderExecutionHistoryLoopNode.
func GetGTMioShaderExecutionHistoryLoopNodeClass() GTMioShaderExecutionHistoryLoopNodeClass {
	return getGTMioShaderExecutionHistoryLoopNodeClass()
}

type GTMioShaderExecutionHistoryLoopNodeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (gc GTMioShaderExecutionHistoryLoopNodeClass) Class() objc.Class {
	return gc.class
}

// Alloc allocates memory for a new instance of the class.
func (gc GTMioShaderExecutionHistoryLoopNodeClass) Alloc() GTMioShaderExecutionHistoryLoopNode {
	rv := objc.SendIfResponds[GTMioShaderExecutionHistoryLoopNode](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [GTMioShaderExecutionHistoryLoopNode.Binary]
//   - [GTMioShaderExecutionHistoryLoopNode.BinaryRange]
//   - [GTMioShaderExecutionHistoryLoopNode.CurrentLoopIndex]
//   - [GTMioShaderExecutionHistoryLoopNode.DebugFilePath]
//   - [GTMioShaderExecutionHistoryLoopNode.DebugFunctionName]
//   - [GTMioShaderExecutionHistoryLoopNode.InstructionBegin]
//   - [GTMioShaderExecutionHistoryLoopNode.InstructionEnd]
//   - [GTMioShaderExecutionHistoryLoopNode.IsLoopRoot]
//   - [GTMioShaderExecutionHistoryLoopNode.Location]
//   - [GTMioShaderExecutionHistoryLoopNode.LoopCount]
//   - [GTMioShaderExecutionHistoryLoopNode.InitWithLoopInstructionBeginEndLoopCountCurrentLoopIndexBinaryParent]
type GTMioShaderExecutionHistoryLoopNode struct {
	GTMioShaderExecutionHistoryNode
}

// GTMioShaderExecutionHistoryLoopNodeFromID constructs a [GTMioShaderExecutionHistoryLoopNode] from an objc.ID.
func GTMioShaderExecutionHistoryLoopNodeFromID(id objc.ID) GTMioShaderExecutionHistoryLoopNode {
	return GTMioShaderExecutionHistoryLoopNode{GTMioShaderExecutionHistoryNode: GTMioShaderExecutionHistoryNodeFromID(id)}
}

// Ensure GTMioShaderExecutionHistoryLoopNode implements IGTMioShaderExecutionHistoryLoopNode.
var _ IGTMioShaderExecutionHistoryLoopNode = GTMioShaderExecutionHistoryLoopNode{}

// An interface definition for the [GTMioShaderExecutionHistoryLoopNode] class.
//
// # Methods
//
//   - [IGTMioShaderExecutionHistoryLoopNode.Binary]
//   - [IGTMioShaderExecutionHistoryLoopNode.BinaryRange]
//   - [IGTMioShaderExecutionHistoryLoopNode.CurrentLoopIndex]
//   - [IGTMioShaderExecutionHistoryLoopNode.DebugFilePath]
//   - [IGTMioShaderExecutionHistoryLoopNode.DebugFunctionName]
//   - [IGTMioShaderExecutionHistoryLoopNode.InstructionBegin]
//   - [IGTMioShaderExecutionHistoryLoopNode.InstructionEnd]
//   - [IGTMioShaderExecutionHistoryLoopNode.IsLoopRoot]
//   - [IGTMioShaderExecutionHistoryLoopNode.Location]
//   - [IGTMioShaderExecutionHistoryLoopNode.LoopCount]
//   - [IGTMioShaderExecutionHistoryLoopNode.InitWithLoopInstructionBeginEndLoopCountCurrentLoopIndexBinaryParent]
type IGTMioShaderExecutionHistoryLoopNode interface {
	IGTMioShaderExecutionHistoryNode

	// Topic: Methods

	Binary() IGTMioShaderBinaryData
	BinaryRange() *GTMioShaderBinaryDebugBinaryRange
	CurrentLoopIndex() uint32
	DebugFilePath() string
	DebugFunctionName() string
	InstructionBegin() uint32
	InstructionEnd() uint32
	IsLoopRoot() bool
	Location() *GTMioShaderBinaryDebugLocation
	LoopCount() uint32
	InitWithLoopInstructionBeginEndLoopCountCurrentLoopIndexBinaryParent(begin uint32, end uint32, count uint32, index uint32, binary objectivec.IObject, parent objectivec.IObject) GTMioShaderExecutionHistoryLoopNode
}

// Init initializes the instance.
func (g GTMioShaderExecutionHistoryLoopNode) Init() GTMioShaderExecutionHistoryLoopNode {
	rv := objc.SendIfResponds[GTMioShaderExecutionHistoryLoopNode](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g GTMioShaderExecutionHistoryLoopNode) Autorelease() GTMioShaderExecutionHistoryLoopNode {
	rv := objc.SendIfResponds[GTMioShaderExecutionHistoryLoopNode](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewGTMioShaderExecutionHistoryLoopNode creates a new GTMioShaderExecutionHistoryLoopNode instance.
func NewGTMioShaderExecutionHistoryLoopNode() GTMioShaderExecutionHistoryLoopNode {
	class := getGTMioShaderExecutionHistoryLoopNodeClass()
	rv := objc.SendIfResponds[GTMioShaderExecutionHistoryLoopNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewGTMioShaderExecutionHistoryLoopNodeWithLoopInstructionBeginEndLoopCountCurrentLoopIndexBinaryParent(begin uint32, end uint32, count uint32, index uint32, binary objectivec.IObject, parent objectivec.IObject) GTMioShaderExecutionHistoryLoopNode {
	instance := getGTMioShaderExecutionHistoryLoopNodeClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithLoopInstructionBegin:end:loopCount:currentLoopIndex:binary:parent:"), begin, end, count, index, binary, parent)
	return GTMioShaderExecutionHistoryLoopNodeFromID(rv)
}

func NewGTMioShaderExecutionHistoryLoopNodeWithTypeParent(type_ uint32, parent objectivec.IObject) GTMioShaderExecutionHistoryLoopNode {
	instance := getGTMioShaderExecutionHistoryLoopNodeClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithType:parent:"), type_, parent)
	return GTMioShaderExecutionHistoryLoopNodeFromID(rv)
}

func (g GTMioShaderExecutionHistoryLoopNode) InitWithLoopInstructionBeginEndLoopCountCurrentLoopIndexBinaryParent(begin uint32, end uint32, count uint32, index uint32, binary objectivec.IObject, parent objectivec.IObject) GTMioShaderExecutionHistoryLoopNode {
	rv := objc.SendIfResponds[GTMioShaderExecutionHistoryLoopNode](g.ID, objc.Sel("initWithLoopInstructionBegin:end:loopCount:currentLoopIndex:binary:parent:"), begin, end, count, index, binary, parent)
	return rv
}

func (g GTMioShaderExecutionHistoryLoopNode) Binary() IGTMioShaderBinaryData {
	rv := objc.SendIfResponds[objc.ID](g.ID, objc.Sel("binary"))
	return GTMioShaderBinaryDataFromID(objc.ID(rv))
}
func (g GTMioShaderExecutionHistoryLoopNode) BinaryRange() *GTMioShaderBinaryDebugBinaryRange {
	rv := objc.SendIfResponds[unsafe.Pointer](g.ID, objc.Sel("binaryRange"))
	return (*GTMioShaderBinaryDebugBinaryRange)(rv)
}
func (g GTMioShaderExecutionHistoryLoopNode) CurrentLoopIndex() uint32 {
	rv := objc.SendIfResponds[uint32](g.ID, objc.Sel("currentLoopIndex"))
	return rv
}
func (g GTMioShaderExecutionHistoryLoopNode) DebugFilePath() string {
	rv := objc.SendIfResponds[objc.ID](g.ID, objc.Sel("debugFilePath"))
	return foundation.NSStringFromID(rv).String()
}
func (g GTMioShaderExecutionHistoryLoopNode) DebugFunctionName() string {
	rv := objc.SendIfResponds[objc.ID](g.ID, objc.Sel("debugFunctionName"))
	return foundation.NSStringFromID(rv).String()
}
func (g GTMioShaderExecutionHistoryLoopNode) InstructionBegin() uint32 {
	rv := objc.SendIfResponds[uint32](g.ID, objc.Sel("instructionBegin"))
	return rv
}
func (g GTMioShaderExecutionHistoryLoopNode) InstructionEnd() uint32 {
	rv := objc.SendIfResponds[uint32](g.ID, objc.Sel("instructionEnd"))
	return rv
}
func (g GTMioShaderExecutionHistoryLoopNode) IsLoopRoot() bool {
	rv := objc.SendIfResponds[bool](g.ID, objc.Sel("isLoopRoot"))
	return rv
}
func (g GTMioShaderExecutionHistoryLoopNode) Location() *GTMioShaderBinaryDebugLocation {
	rv := objc.SendIfResponds[unsafe.Pointer](g.ID, objc.Sel("location"))
	return (*GTMioShaderBinaryDebugLocation)(rv)
}
func (g GTMioShaderExecutionHistoryLoopNode) LoopCount() uint32 {
	rv := objc.SendIfResponds[uint32](g.ID, objc.Sel("loopCount"))
	return rv
}
