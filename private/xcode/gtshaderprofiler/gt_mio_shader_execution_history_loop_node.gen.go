// Code generated from Apple documentation for gtshaderprofiler. DO NOT EDIT.

package gtshaderprofiler

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
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
	rv := objc.Send[GTMioShaderExecutionHistoryLoopNode](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

//
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
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderExecutionHistoryLoopNode
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
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderExecutionHistoryLoopNode
type IGTMioShaderExecutionHistoryLoopNode interface {
	IGTMioShaderExecutionHistoryNode

	// Topic: Methods

	Binary() IGTMioShaderBinaryData
	BinaryRange() unsafe.Pointer
	CurrentLoopIndex() uint32
	DebugFilePath() string
	DebugFunctionName() string
	InstructionBegin() uint32
	InstructionEnd() uint32
	IsLoopRoot() bool
	Location() unsafe.Pointer
	LoopCount() uint32
	InitWithLoopInstructionBeginEndLoopCountCurrentLoopIndexBinaryParent(begin uint32, end uint32, count uint32, index uint32, binary objectivec.IObject, parent objectivec.IObject) GTMioShaderExecutionHistoryLoopNode
}

// Init initializes the instance.
func (g GTMioShaderExecutionHistoryLoopNode) Init() GTMioShaderExecutionHistoryLoopNode {
	rv := objc.Send[GTMioShaderExecutionHistoryLoopNode](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g GTMioShaderExecutionHistoryLoopNode) Autorelease() GTMioShaderExecutionHistoryLoopNode {
	rv := objc.Send[GTMioShaderExecutionHistoryLoopNode](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewGTMioShaderExecutionHistoryLoopNode creates a new GTMioShaderExecutionHistoryLoopNode instance.
func NewGTMioShaderExecutionHistoryLoopNode() GTMioShaderExecutionHistoryLoopNode {
	class := getGTMioShaderExecutionHistoryLoopNodeClass()
	rv := objc.Send[GTMioShaderExecutionHistoryLoopNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderExecutionHistoryLoopNode/initWithLoopInstructionBegin:end:loopCount:currentLoopIndex:binary:parent:
func NewGTMioShaderExecutionHistoryLoopNodeWithLoopInstructionBeginEndLoopCountCurrentLoopIndexBinaryParent(begin uint32, end uint32, count uint32, index uint32, binary objectivec.IObject, parent objectivec.IObject) GTMioShaderExecutionHistoryLoopNode {
	instance := getGTMioShaderExecutionHistoryLoopNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithLoopInstructionBegin:end:loopCount:currentLoopIndex:binary:parent:"), begin, end, count, index, binary, parent)
	return GTMioShaderExecutionHistoryLoopNodeFromID(rv)
}

//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderExecutionHistoryNode/initWithType:parent:
func NewGTMioShaderExecutionHistoryLoopNodeWithTypeParent(type_ uint32, parent objectivec.IObject) GTMioShaderExecutionHistoryLoopNode {
	instance := getGTMioShaderExecutionHistoryLoopNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithType:parent:"), type_, parent)
	return GTMioShaderExecutionHistoryLoopNodeFromID(rv)
}

//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderExecutionHistoryLoopNode/initWithLoopInstructionBegin:end:loopCount:currentLoopIndex:binary:parent:
func (g GTMioShaderExecutionHistoryLoopNode) InitWithLoopInstructionBeginEndLoopCountCurrentLoopIndexBinaryParent(begin uint32, end uint32, count uint32, index uint32, binary objectivec.IObject, parent objectivec.IObject) GTMioShaderExecutionHistoryLoopNode {
	rv := objc.Send[GTMioShaderExecutionHistoryLoopNode](g.ID, objc.Sel("initWithLoopInstructionBegin:end:loopCount:currentLoopIndex:binary:parent:"), begin, end, count, index, binary, parent)
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderExecutionHistoryLoopNode/binary
func (g GTMioShaderExecutionHistoryLoopNode) Binary() IGTMioShaderBinaryData {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("binary"))
	return GTMioShaderBinaryDataFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderExecutionHistoryLoopNode/binaryRange
func (g GTMioShaderExecutionHistoryLoopNode) BinaryRange() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g.ID, objc.Sel("binaryRange"))
	return rv
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderExecutionHistoryLoopNode/currentLoopIndex
func (g GTMioShaderExecutionHistoryLoopNode) CurrentLoopIndex() uint32 {
	rv := objc.Send[uint32](g.ID, objc.Sel("currentLoopIndex"))
	return rv
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderExecutionHistoryLoopNode/debugFilePath
func (g GTMioShaderExecutionHistoryLoopNode) DebugFilePath() string {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("debugFilePath"))
	return foundation.NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderExecutionHistoryLoopNode/debugFunctionName
func (g GTMioShaderExecutionHistoryLoopNode) DebugFunctionName() string {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("debugFunctionName"))
	return foundation.NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderExecutionHistoryLoopNode/instructionBegin
func (g GTMioShaderExecutionHistoryLoopNode) InstructionBegin() uint32 {
	rv := objc.Send[uint32](g.ID, objc.Sel("instructionBegin"))
	return rv
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderExecutionHistoryLoopNode/instructionEnd
func (g GTMioShaderExecutionHistoryLoopNode) InstructionEnd() uint32 {
	rv := objc.Send[uint32](g.ID, objc.Sel("instructionEnd"))
	return rv
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderExecutionHistoryLoopNode/isLoopRoot
func (g GTMioShaderExecutionHistoryLoopNode) IsLoopRoot() bool {
	rv := objc.Send[bool](g.ID, objc.Sel("isLoopRoot"))
	return rv
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderExecutionHistoryLoopNode/location
func (g GTMioShaderExecutionHistoryLoopNode) Location() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g.ID, objc.Sel("location"))
	return rv
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderExecutionHistoryLoopNode/loopCount
func (g GTMioShaderExecutionHistoryLoopNode) LoopCount() uint32 {
	rv := objc.Send[uint32](g.ID, objc.Sel("loopCount"))
	return rv
}

