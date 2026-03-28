// Code generated from Apple documentation for gtshaderprofiler. DO NOT EDIT.

package gtshaderprofiler

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [GTMioShaderExecutionHistoryCliqueNode] class.
var (
	_GTMioShaderExecutionHistoryCliqueNodeClass     GTMioShaderExecutionHistoryCliqueNodeClass
	_GTMioShaderExecutionHistoryCliqueNodeClassOnce sync.Once
)

func getGTMioShaderExecutionHistoryCliqueNodeClass() GTMioShaderExecutionHistoryCliqueNodeClass {
	_GTMioShaderExecutionHistoryCliqueNodeClassOnce.Do(func() {
		_GTMioShaderExecutionHistoryCliqueNodeClass = GTMioShaderExecutionHistoryCliqueNodeClass{class: objc.GetClass("GTMioShaderExecutionHistoryCliqueNode")}
	})
	return _GTMioShaderExecutionHistoryCliqueNodeClass
}

// GetGTMioShaderExecutionHistoryCliqueNodeClass returns the class object for GTMioShaderExecutionHistoryCliqueNode.
func GetGTMioShaderExecutionHistoryCliqueNodeClass() GTMioShaderExecutionHistoryCliqueNodeClass {
	return getGTMioShaderExecutionHistoryCliqueNodeClass()
}

type GTMioShaderExecutionHistoryCliqueNodeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (gc GTMioShaderExecutionHistoryCliqueNodeClass) Class() objc.Class {
	return gc.class
}

// Alloc allocates memory for a new instance of the class.
func (gc GTMioShaderExecutionHistoryCliqueNodeClass) Alloc() GTMioShaderExecutionHistoryCliqueNode {
	rv := objc.Send[GTMioShaderExecutionHistoryCliqueNode](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [GTMioShaderExecutionHistoryCliqueNode.Clique]
//   - [GTMioShaderExecutionHistoryCliqueNode.Usc]
//   - [GTMioShaderExecutionHistoryCliqueNode.InitWithCliqueUscParent]
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderExecutionHistoryCliqueNode
type GTMioShaderExecutionHistoryCliqueNode struct {
	GTMioShaderExecutionHistoryNode
}

// GTMioShaderExecutionHistoryCliqueNodeFromID constructs a [GTMioShaderExecutionHistoryCliqueNode] from an objc.ID.
func GTMioShaderExecutionHistoryCliqueNodeFromID(id objc.ID) GTMioShaderExecutionHistoryCliqueNode {
	return GTMioShaderExecutionHistoryCliqueNode{GTMioShaderExecutionHistoryNode: GTMioShaderExecutionHistoryNodeFromID(id)}
}
// Ensure GTMioShaderExecutionHistoryCliqueNode implements IGTMioShaderExecutionHistoryCliqueNode.
var _ IGTMioShaderExecutionHistoryCliqueNode = GTMioShaderExecutionHistoryCliqueNode{}

// An interface definition for the [GTMioShaderExecutionHistoryCliqueNode] class.
//
// # Methods
//
//   - [IGTMioShaderExecutionHistoryCliqueNode.Clique]
//   - [IGTMioShaderExecutionHistoryCliqueNode.Usc]
//   - [IGTMioShaderExecutionHistoryCliqueNode.InitWithCliqueUscParent]
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderExecutionHistoryCliqueNode
type IGTMioShaderExecutionHistoryCliqueNode interface {
	IGTMioShaderExecutionHistoryNode

	// Topic: Methods

	Clique() unsafe.Pointer
	Usc() IGTMioUSCTraceData
	InitWithCliqueUscParent(clique unsafe.Pointer, usc objectivec.IObject, parent objectivec.IObject) GTMioShaderExecutionHistoryCliqueNode
}

// Init initializes the instance.
func (g GTMioShaderExecutionHistoryCliqueNode) Init() GTMioShaderExecutionHistoryCliqueNode {
	rv := objc.Send[GTMioShaderExecutionHistoryCliqueNode](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g GTMioShaderExecutionHistoryCliqueNode) Autorelease() GTMioShaderExecutionHistoryCliqueNode {
	rv := objc.Send[GTMioShaderExecutionHistoryCliqueNode](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewGTMioShaderExecutionHistoryCliqueNode creates a new GTMioShaderExecutionHistoryCliqueNode instance.
func NewGTMioShaderExecutionHistoryCliqueNode() GTMioShaderExecutionHistoryCliqueNode {
	class := getGTMioShaderExecutionHistoryCliqueNodeClass()
	rv := objc.Send[GTMioShaderExecutionHistoryCliqueNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderExecutionHistoryCliqueNode/initWithClique:usc:parent:
func NewGTMioShaderExecutionHistoryCliqueNodeWithCliqueUscParent(clique unsafe.Pointer, usc objectivec.IObject, parent objectivec.IObject) GTMioShaderExecutionHistoryCliqueNode {
	instance := getGTMioShaderExecutionHistoryCliqueNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithClique:usc:parent:"), clique, usc, parent)
	return GTMioShaderExecutionHistoryCliqueNodeFromID(rv)
}

//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderExecutionHistoryNode/initWithType:parent:
func NewGTMioShaderExecutionHistoryCliqueNodeWithTypeParent(type_ uint32, parent objectivec.IObject) GTMioShaderExecutionHistoryCliqueNode {
	instance := getGTMioShaderExecutionHistoryCliqueNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithType:parent:"), type_, parent)
	return GTMioShaderExecutionHistoryCliqueNodeFromID(rv)
}

//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderExecutionHistoryCliqueNode/initWithClique:usc:parent:
func (g GTMioShaderExecutionHistoryCliqueNode) InitWithCliqueUscParent(clique unsafe.Pointer, usc objectivec.IObject, parent objectivec.IObject) GTMioShaderExecutionHistoryCliqueNode {
	rv := objc.Send[GTMioShaderExecutionHistoryCliqueNode](g.ID, objc.Sel("initWithClique:usc:parent:"), clique, usc, parent)
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderExecutionHistoryCliqueNode/clique
func (g GTMioShaderExecutionHistoryCliqueNode) Clique() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g.ID, objc.Sel("clique"))
	return rv
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioShaderExecutionHistoryCliqueNode/usc
func (g GTMioShaderExecutionHistoryCliqueNode) Usc() IGTMioUSCTraceData {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("usc"))
	return GTMioUSCTraceDataFromID(objc.ID(rv))
}

