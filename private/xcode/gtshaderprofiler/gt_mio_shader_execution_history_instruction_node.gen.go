// Code generated from Apple documentation for gtshaderprofiler. DO NOT EDIT.

package gtshaderprofiler

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [GTMioShaderExecutionHistoryInstructionNode] class.
var (
	_GTMioShaderExecutionHistoryInstructionNodeClass     GTMioShaderExecutionHistoryInstructionNodeClass
	_GTMioShaderExecutionHistoryInstructionNodeClassOnce sync.Once
)

func getGTMioShaderExecutionHistoryInstructionNodeClass() GTMioShaderExecutionHistoryInstructionNodeClass {
	_GTMioShaderExecutionHistoryInstructionNodeClassOnce.Do(func() {
		_GTMioShaderExecutionHistoryInstructionNodeClass = GTMioShaderExecutionHistoryInstructionNodeClass{class: objc.GetClass("GTMioShaderExecutionHistoryInstructionNode")}
	})
	return _GTMioShaderExecutionHistoryInstructionNodeClass
}

// GetGTMioShaderExecutionHistoryInstructionNodeClass returns the class object for GTMioShaderExecutionHistoryInstructionNode.
func GetGTMioShaderExecutionHistoryInstructionNodeClass() GTMioShaderExecutionHistoryInstructionNodeClass {
	return getGTMioShaderExecutionHistoryInstructionNodeClass()
}

type GTMioShaderExecutionHistoryInstructionNodeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (gc GTMioShaderExecutionHistoryInstructionNodeClass) Class() objc.Class {
	return gc.class
}

// Alloc allocates memory for a new instance of the class.
func (gc GTMioShaderExecutionHistoryInstructionNodeClass) Alloc() GTMioShaderExecutionHistoryInstructionNode {
	rv := objc.Send[GTMioShaderExecutionHistoryInstructionNode](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [GTMioShaderExecutionHistoryInstructionNode.Binary]
//   - [GTMioShaderExecutionHistoryInstructionNode.DebugFilePath]
//   - [GTMioShaderExecutionHistoryInstructionNode.DebugFunctionName]
//   - [GTMioShaderExecutionHistoryInstructionNode.InstructionIndex]
//   - [GTMioShaderExecutionHistoryInstructionNode.InstructionInfo]
//   - [GTMioShaderExecutionHistoryInstructionNode.Isa]
//   - [GTMioShaderExecutionHistoryInstructionNode.Location]
//   - [GTMioShaderExecutionHistoryInstructionNode.InitWithInstructionIndexLocationBinaryParent]
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderExecutionHistoryInstructionNode
type GTMioShaderExecutionHistoryInstructionNode struct {
	GTMioShaderExecutionHistoryNode
}

// GTMioShaderExecutionHistoryInstructionNodeFromID constructs a [GTMioShaderExecutionHistoryInstructionNode] from an objc.ID.
func GTMioShaderExecutionHistoryInstructionNodeFromID(id objc.ID) GTMioShaderExecutionHistoryInstructionNode {
	return GTMioShaderExecutionHistoryInstructionNode{GTMioShaderExecutionHistoryNode: GTMioShaderExecutionHistoryNodeFromID(id)}
}
// Ensure GTMioShaderExecutionHistoryInstructionNode implements IGTMioShaderExecutionHistoryInstructionNode.
var _ IGTMioShaderExecutionHistoryInstructionNode = GTMioShaderExecutionHistoryInstructionNode{}

// An interface definition for the [GTMioShaderExecutionHistoryInstructionNode] class.
//
// # Methods
//
//   - [IGTMioShaderExecutionHistoryInstructionNode.Binary]
//   - [IGTMioShaderExecutionHistoryInstructionNode.DebugFilePath]
//   - [IGTMioShaderExecutionHistoryInstructionNode.DebugFunctionName]
//   - [IGTMioShaderExecutionHistoryInstructionNode.InstructionIndex]
//   - [IGTMioShaderExecutionHistoryInstructionNode.InstructionInfo]
//   - [IGTMioShaderExecutionHistoryInstructionNode.Isa]
//   - [IGTMioShaderExecutionHistoryInstructionNode.Location]
//   - [IGTMioShaderExecutionHistoryInstructionNode.InitWithInstructionIndexLocationBinaryParent]
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderExecutionHistoryInstructionNode
type IGTMioShaderExecutionHistoryInstructionNode interface {
	IGTMioShaderExecutionHistoryNode

	// Topic: Methods

	Binary() IGTMioShaderBinaryData
	DebugFilePath() string
	DebugFunctionName() string
	InstructionIndex() uint32
	InstructionInfo() unsafe.Pointer
	Isa() string
	Location() unsafe.Pointer
	InitWithInstructionIndexLocationBinaryParent(index uint32, location unsafe.Pointer, binary objectivec.IObject, parent objectivec.IObject) GTMioShaderExecutionHistoryInstructionNode
}

// Init initializes the instance.
func (g GTMioShaderExecutionHistoryInstructionNode) Init() GTMioShaderExecutionHistoryInstructionNode {
	rv := objc.Send[GTMioShaderExecutionHistoryInstructionNode](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g GTMioShaderExecutionHistoryInstructionNode) Autorelease() GTMioShaderExecutionHistoryInstructionNode {
	rv := objc.Send[GTMioShaderExecutionHistoryInstructionNode](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewGTMioShaderExecutionHistoryInstructionNode creates a new GTMioShaderExecutionHistoryInstructionNode instance.
func NewGTMioShaderExecutionHistoryInstructionNode() GTMioShaderExecutionHistoryInstructionNode {
	class := getGTMioShaderExecutionHistoryInstructionNodeClass()
	rv := objc.Send[GTMioShaderExecutionHistoryInstructionNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderExecutionHistoryInstructionNode/initWithInstructionIndex:location:binary:parent:
func NewGTMioShaderExecutionHistoryInstructionNodeWithInstructionIndexLocationBinaryParent(index uint32, location unsafe.Pointer, binary objectivec.IObject, parent objectivec.IObject) GTMioShaderExecutionHistoryInstructionNode {
	instance := getGTMioShaderExecutionHistoryInstructionNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithInstructionIndex:location:binary:parent:"), index, location, binary, parent)
	return GTMioShaderExecutionHistoryInstructionNodeFromID(rv)
}

//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderExecutionHistoryNode/initWithType:parent:
func NewGTMioShaderExecutionHistoryInstructionNodeWithTypeParent(type_ uint32, parent objectivec.IObject) GTMioShaderExecutionHistoryInstructionNode {
	instance := getGTMioShaderExecutionHistoryInstructionNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithType:parent:"), type_, parent)
	return GTMioShaderExecutionHistoryInstructionNodeFromID(rv)
}

//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderExecutionHistoryInstructionNode/initWithInstructionIndex:location:binary:parent:
func (g GTMioShaderExecutionHistoryInstructionNode) InitWithInstructionIndexLocationBinaryParent(index uint32, location unsafe.Pointer, binary objectivec.IObject, parent objectivec.IObject) GTMioShaderExecutionHistoryInstructionNode {
	rv := objc.Send[GTMioShaderExecutionHistoryInstructionNode](g.ID, objc.Sel("initWithInstructionIndex:location:binary:parent:"), index, location, binary, parent)
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderExecutionHistoryInstructionNode/binary
func (g GTMioShaderExecutionHistoryInstructionNode) Binary() IGTMioShaderBinaryData {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("binary"))
	return GTMioShaderBinaryDataFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderExecutionHistoryInstructionNode/debugFilePath
func (g GTMioShaderExecutionHistoryInstructionNode) DebugFilePath() string {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("debugFilePath"))
	return foundation.NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderExecutionHistoryInstructionNode/debugFunctionName
func (g GTMioShaderExecutionHistoryInstructionNode) DebugFunctionName() string {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("debugFunctionName"))
	return foundation.NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderExecutionHistoryInstructionNode/instructionIndex
func (g GTMioShaderExecutionHistoryInstructionNode) InstructionIndex() uint32 {
	rv := objc.Send[uint32](g.ID, objc.Sel("instructionIndex"))
	return rv
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderExecutionHistoryInstructionNode/instructionInfo
func (g GTMioShaderExecutionHistoryInstructionNode) InstructionInfo() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g.ID, objc.Sel("instructionInfo"))
	return rv
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderExecutionHistoryInstructionNode/isa
func (g GTMioShaderExecutionHistoryInstructionNode) Isa() string {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("isa"))
	return foundation.NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderExecutionHistoryInstructionNode/location
func (g GTMioShaderExecutionHistoryInstructionNode) Location() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g.ID, objc.Sel("location"))
	return rv
}

