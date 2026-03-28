// Code generated from Apple documentation for gtshaderprofiler. DO NOT EDIT.

package gtshaderprofiler

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [GTMioShaderExecutionHistoryFunctionNode] class.
var (
	_GTMioShaderExecutionHistoryFunctionNodeClass     GTMioShaderExecutionHistoryFunctionNodeClass
	_GTMioShaderExecutionHistoryFunctionNodeClassOnce sync.Once
)

func getGTMioShaderExecutionHistoryFunctionNodeClass() GTMioShaderExecutionHistoryFunctionNodeClass {
	_GTMioShaderExecutionHistoryFunctionNodeClassOnce.Do(func() {
		_GTMioShaderExecutionHistoryFunctionNodeClass = GTMioShaderExecutionHistoryFunctionNodeClass{class: objc.GetClass("GTMioShaderExecutionHistoryFunctionNode")}
	})
	return _GTMioShaderExecutionHistoryFunctionNodeClass
}

// GetGTMioShaderExecutionHistoryFunctionNodeClass returns the class object for GTMioShaderExecutionHistoryFunctionNode.
func GetGTMioShaderExecutionHistoryFunctionNodeClass() GTMioShaderExecutionHistoryFunctionNodeClass {
	return getGTMioShaderExecutionHistoryFunctionNodeClass()
}

type GTMioShaderExecutionHistoryFunctionNodeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (gc GTMioShaderExecutionHistoryFunctionNodeClass) Class() objc.Class {
	return gc.class
}

// Alloc allocates memory for a new instance of the class.
func (gc GTMioShaderExecutionHistoryFunctionNodeClass) Alloc() GTMioShaderExecutionHistoryFunctionNode {
	rv := objc.Send[GTMioShaderExecutionHistoryFunctionNode](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [GTMioShaderExecutionHistoryFunctionNode.Binary]
//   - [GTMioShaderExecutionHistoryFunctionNode.BinaryRange]
//   - [GTMioShaderExecutionHistoryFunctionNode.BinaryRangeIndex]
//   - [GTMioShaderExecutionHistoryFunctionNode.CallCount]
//   - [GTMioShaderExecutionHistoryFunctionNode.CallerBinary]
//   - [GTMioShaderExecutionHistoryFunctionNode.CallerBinaryRange]
//   - [GTMioShaderExecutionHistoryFunctionNode.CallerBinaryRangeIndex]
//   - [GTMioShaderExecutionHistoryFunctionNode.CallerDebugFilePath]
//   - [GTMioShaderExecutionHistoryFunctionNode.CallerDebugFunctionName]
//   - [GTMioShaderExecutionHistoryFunctionNode.CallerLocation]
//   - [GTMioShaderExecutionHistoryFunctionNode.DebugFilePath]
//   - [GTMioShaderExecutionHistoryFunctionNode.DebugFunctionName]
//   - [GTMioShaderExecutionHistoryFunctionNode.Identifier]
//   - [GTMioShaderExecutionHistoryFunctionNode.IncrementCallCount]
//   - [GTMioShaderExecutionHistoryFunctionNode.Inlined]
//   - [GTMioShaderExecutionHistoryFunctionNode.Location]
//   - [GTMioShaderExecutionHistoryFunctionNode.Synthesized]
//   - [GTMioShaderExecutionHistoryFunctionNode.InitWithBinaryCliqueParentIdentifier]
//   - [GTMioShaderExecutionHistoryFunctionNode.InitWithLocationInlinedBinaryRangeIndexBinaryCallerLocationCallerBinaryRangeIndexCallerBinaryIdentifierParent]
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderExecutionHistoryFunctionNode
type GTMioShaderExecutionHistoryFunctionNode struct {
	GTMioShaderExecutionHistoryNode
}

// GTMioShaderExecutionHistoryFunctionNodeFromID constructs a [GTMioShaderExecutionHistoryFunctionNode] from an objc.ID.
func GTMioShaderExecutionHistoryFunctionNodeFromID(id objc.ID) GTMioShaderExecutionHistoryFunctionNode {
	return GTMioShaderExecutionHistoryFunctionNode{GTMioShaderExecutionHistoryNode: GTMioShaderExecutionHistoryNodeFromID(id)}
}
// Ensure GTMioShaderExecutionHistoryFunctionNode implements IGTMioShaderExecutionHistoryFunctionNode.
var _ IGTMioShaderExecutionHistoryFunctionNode = GTMioShaderExecutionHistoryFunctionNode{}

// An interface definition for the [GTMioShaderExecutionHistoryFunctionNode] class.
//
// # Methods
//
//   - [IGTMioShaderExecutionHistoryFunctionNode.Binary]
//   - [IGTMioShaderExecutionHistoryFunctionNode.BinaryRange]
//   - [IGTMioShaderExecutionHistoryFunctionNode.BinaryRangeIndex]
//   - [IGTMioShaderExecutionHistoryFunctionNode.CallCount]
//   - [IGTMioShaderExecutionHistoryFunctionNode.CallerBinary]
//   - [IGTMioShaderExecutionHistoryFunctionNode.CallerBinaryRange]
//   - [IGTMioShaderExecutionHistoryFunctionNode.CallerBinaryRangeIndex]
//   - [IGTMioShaderExecutionHistoryFunctionNode.CallerDebugFilePath]
//   - [IGTMioShaderExecutionHistoryFunctionNode.CallerDebugFunctionName]
//   - [IGTMioShaderExecutionHistoryFunctionNode.CallerLocation]
//   - [IGTMioShaderExecutionHistoryFunctionNode.DebugFilePath]
//   - [IGTMioShaderExecutionHistoryFunctionNode.DebugFunctionName]
//   - [IGTMioShaderExecutionHistoryFunctionNode.Identifier]
//   - [IGTMioShaderExecutionHistoryFunctionNode.IncrementCallCount]
//   - [IGTMioShaderExecutionHistoryFunctionNode.Inlined]
//   - [IGTMioShaderExecutionHistoryFunctionNode.Location]
//   - [IGTMioShaderExecutionHistoryFunctionNode.Synthesized]
//   - [IGTMioShaderExecutionHistoryFunctionNode.InitWithBinaryCliqueParentIdentifier]
//   - [IGTMioShaderExecutionHistoryFunctionNode.InitWithLocationInlinedBinaryRangeIndexBinaryCallerLocationCallerBinaryRangeIndexCallerBinaryIdentifierParent]
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderExecutionHistoryFunctionNode
type IGTMioShaderExecutionHistoryFunctionNode interface {
	IGTMioShaderExecutionHistoryNode

	// Topic: Methods

	Binary() IGTMioShaderBinaryData
	BinaryRange() unsafe.Pointer
	BinaryRangeIndex() uint32
	CallCount() uint64
	CallerBinary() IGTMioShaderBinaryData
	CallerBinaryRange() unsafe.Pointer
	CallerBinaryRangeIndex() uint32
	CallerDebugFilePath() string
	CallerDebugFunctionName() string
	CallerLocation() unsafe.Pointer
	DebugFilePath() string
	DebugFunctionName() string
	Identifier() uint64
	IncrementCallCount()
	Inlined() bool
	Location() unsafe.Pointer
	Synthesized() bool
	InitWithBinaryCliqueParentIdentifier(binary objectivec.IObject, clique unsafe.Pointer, parent objectivec.IObject, identifier uint64) GTMioShaderExecutionHistoryFunctionNode
	InitWithLocationInlinedBinaryRangeIndexBinaryCallerLocationCallerBinaryRangeIndexCallerBinaryIdentifierParent(location unsafe.Pointer, inlined bool, index uint32, binary objectivec.IObject, location2 unsafe.Pointer, index2 uint32, binary2 objectivec.IObject, identifier uint64, parent objectivec.IObject) GTMioShaderExecutionHistoryFunctionNode
}

// Init initializes the instance.
func (g GTMioShaderExecutionHistoryFunctionNode) Init() GTMioShaderExecutionHistoryFunctionNode {
	rv := objc.Send[GTMioShaderExecutionHistoryFunctionNode](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g GTMioShaderExecutionHistoryFunctionNode) Autorelease() GTMioShaderExecutionHistoryFunctionNode {
	rv := objc.Send[GTMioShaderExecutionHistoryFunctionNode](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewGTMioShaderExecutionHistoryFunctionNode creates a new GTMioShaderExecutionHistoryFunctionNode instance.
func NewGTMioShaderExecutionHistoryFunctionNode() GTMioShaderExecutionHistoryFunctionNode {
	class := getGTMioShaderExecutionHistoryFunctionNodeClass()
	rv := objc.Send[GTMioShaderExecutionHistoryFunctionNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderExecutionHistoryFunctionNode/initWithBinary:clique:parent:identifier:
func NewGTMioShaderExecutionHistoryFunctionNodeWithBinaryCliqueParentIdentifier(binary objectivec.IObject, clique unsafe.Pointer, parent objectivec.IObject, identifier uint64) GTMioShaderExecutionHistoryFunctionNode {
	instance := getGTMioShaderExecutionHistoryFunctionNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithBinary:clique:parent:identifier:"), binary, clique, parent, identifier)
	return GTMioShaderExecutionHistoryFunctionNodeFromID(rv)
}

//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderExecutionHistoryFunctionNode/initWithLocation:inlined:binaryRangeIndex:binary:callerLocation:callerBinaryRangeIndex:callerBinary:identifier:parent:
func NewGTMioShaderExecutionHistoryFunctionNodeWithLocationInlinedBinaryRangeIndexBinaryCallerLocationCallerBinaryRangeIndexCallerBinaryIdentifierParent(location unsafe.Pointer, inlined bool, index uint32, binary objectivec.IObject, location2 unsafe.Pointer, index2 uint32, binary2 objectivec.IObject, identifier uint64, parent objectivec.IObject) GTMioShaderExecutionHistoryFunctionNode {
	instance := getGTMioShaderExecutionHistoryFunctionNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithLocation:inlined:binaryRangeIndex:binary:callerLocation:callerBinaryRangeIndex:callerBinary:identifier:parent:"), location, inlined, index, binary, location2, index2, binary2, identifier, parent)
	return GTMioShaderExecutionHistoryFunctionNodeFromID(rv)
}

//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderExecutionHistoryNode/initWithType:parent:
func NewGTMioShaderExecutionHistoryFunctionNodeWithTypeParent(type_ uint32, parent objectivec.IObject) GTMioShaderExecutionHistoryFunctionNode {
	instance := getGTMioShaderExecutionHistoryFunctionNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithType:parent:"), type_, parent)
	return GTMioShaderExecutionHistoryFunctionNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderExecutionHistoryFunctionNode/incrementCallCount
func (g GTMioShaderExecutionHistoryFunctionNode) IncrementCallCount() {
	objc.Send[objc.ID](g.ID, objc.Sel("incrementCallCount"))
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderExecutionHistoryFunctionNode/initWithBinary:clique:parent:identifier:
func (g GTMioShaderExecutionHistoryFunctionNode) InitWithBinaryCliqueParentIdentifier(binary objectivec.IObject, clique unsafe.Pointer, parent objectivec.IObject, identifier uint64) GTMioShaderExecutionHistoryFunctionNode {
	rv := objc.Send[GTMioShaderExecutionHistoryFunctionNode](g.ID, objc.Sel("initWithBinary:clique:parent:identifier:"), binary, clique, parent, identifier)
	return rv
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderExecutionHistoryFunctionNode/initWithLocation:inlined:binaryRangeIndex:binary:callerLocation:callerBinaryRangeIndex:callerBinary:identifier:parent:
func (g GTMioShaderExecutionHistoryFunctionNode) InitWithLocationInlinedBinaryRangeIndexBinaryCallerLocationCallerBinaryRangeIndexCallerBinaryIdentifierParent(location unsafe.Pointer, inlined bool, index uint32, binary objectivec.IObject, location2 unsafe.Pointer, index2 uint32, binary2 objectivec.IObject, identifier uint64, parent objectivec.IObject) GTMioShaderExecutionHistoryFunctionNode {
	rv := objc.Send[GTMioShaderExecutionHistoryFunctionNode](g.ID, objc.Sel("initWithLocation:inlined:binaryRangeIndex:binary:callerLocation:callerBinaryRangeIndex:callerBinary:identifier:parent:"), location, inlined, index, binary, location2, index2, binary2, identifier, parent)
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderExecutionHistoryFunctionNode/binary
func (g GTMioShaderExecutionHistoryFunctionNode) Binary() IGTMioShaderBinaryData {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("binary"))
	return GTMioShaderBinaryDataFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderExecutionHistoryFunctionNode/binaryRange
func (g GTMioShaderExecutionHistoryFunctionNode) BinaryRange() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g.ID, objc.Sel("binaryRange"))
	return rv
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderExecutionHistoryFunctionNode/binaryRangeIndex
func (g GTMioShaderExecutionHistoryFunctionNode) BinaryRangeIndex() uint32 {
	rv := objc.Send[uint32](g.ID, objc.Sel("binaryRangeIndex"))
	return rv
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderExecutionHistoryFunctionNode/callCount
func (g GTMioShaderExecutionHistoryFunctionNode) CallCount() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("callCount"))
	return rv
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderExecutionHistoryFunctionNode/callerBinary
func (g GTMioShaderExecutionHistoryFunctionNode) CallerBinary() IGTMioShaderBinaryData {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("callerBinary"))
	return GTMioShaderBinaryDataFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderExecutionHistoryFunctionNode/callerBinaryRange
func (g GTMioShaderExecutionHistoryFunctionNode) CallerBinaryRange() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g.ID, objc.Sel("callerBinaryRange"))
	return rv
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderExecutionHistoryFunctionNode/callerBinaryRangeIndex
func (g GTMioShaderExecutionHistoryFunctionNode) CallerBinaryRangeIndex() uint32 {
	rv := objc.Send[uint32](g.ID, objc.Sel("callerBinaryRangeIndex"))
	return rv
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderExecutionHistoryFunctionNode/callerDebugFilePath
func (g GTMioShaderExecutionHistoryFunctionNode) CallerDebugFilePath() string {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("callerDebugFilePath"))
	return foundation.NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderExecutionHistoryFunctionNode/callerDebugFunctionName
func (g GTMioShaderExecutionHistoryFunctionNode) CallerDebugFunctionName() string {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("callerDebugFunctionName"))
	return foundation.NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderExecutionHistoryFunctionNode/callerLocation
func (g GTMioShaderExecutionHistoryFunctionNode) CallerLocation() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g.ID, objc.Sel("callerLocation"))
	return rv
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderExecutionHistoryFunctionNode/debugFilePath
func (g GTMioShaderExecutionHistoryFunctionNode) DebugFilePath() string {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("debugFilePath"))
	return foundation.NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderExecutionHistoryFunctionNode/debugFunctionName
func (g GTMioShaderExecutionHistoryFunctionNode) DebugFunctionName() string {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("debugFunctionName"))
	return foundation.NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderExecutionHistoryFunctionNode/identifier
func (g GTMioShaderExecutionHistoryFunctionNode) Identifier() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("identifier"))
	return rv
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderExecutionHistoryFunctionNode/inlined
func (g GTMioShaderExecutionHistoryFunctionNode) Inlined() bool {
	rv := objc.Send[bool](g.ID, objc.Sel("inlined"))
	return rv
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderExecutionHistoryFunctionNode/location
func (g GTMioShaderExecutionHistoryFunctionNode) Location() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g.ID, objc.Sel("location"))
	return rv
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderExecutionHistoryFunctionNode/synthesized
func (g GTMioShaderExecutionHistoryFunctionNode) Synthesized() bool {
	rv := objc.Send[bool](g.ID, objc.Sel("synthesized"))
	return rv
}

