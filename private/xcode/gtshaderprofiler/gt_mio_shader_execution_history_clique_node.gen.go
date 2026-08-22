// Code generated from Apple documentation for gtshaderprofiler. DO NOT EDIT.

package gtshaderprofiler

import (
	"sync"
	"unsafe"

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
	rv := objc.SendIfResponds[GTMioShaderExecutionHistoryCliqueNode](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [GTMioShaderExecutionHistoryCliqueNode.Clique]
//   - [GTMioShaderExecutionHistoryCliqueNode.Usc]
//   - [GTMioShaderExecutionHistoryCliqueNode.InitWithCliqueUscParent]
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
type IGTMioShaderExecutionHistoryCliqueNode interface {
	IGTMioShaderExecutionHistoryNode

	// Topic: Methods

	Clique() *GTMioUSCCliqueMetadata
	Usc() IGTMioUSCTraceData
	InitWithCliqueUscParent(clique *GTMioUSCCliqueMetadata, usc objectivec.IObject, parent objectivec.IObject) GTMioShaderExecutionHistoryCliqueNode
}

// Init initializes the instance.
func (g GTMioShaderExecutionHistoryCliqueNode) Init() GTMioShaderExecutionHistoryCliqueNode {
	rv := objc.SendIfResponds[GTMioShaderExecutionHistoryCliqueNode](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g GTMioShaderExecutionHistoryCliqueNode) Autorelease() GTMioShaderExecutionHistoryCliqueNode {
	rv := objc.SendIfResponds[GTMioShaderExecutionHistoryCliqueNode](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewGTMioShaderExecutionHistoryCliqueNode creates a new GTMioShaderExecutionHistoryCliqueNode instance.
func NewGTMioShaderExecutionHistoryCliqueNode() GTMioShaderExecutionHistoryCliqueNode {
	class := getGTMioShaderExecutionHistoryCliqueNodeClass()
	rv := objc.SendIfResponds[GTMioShaderExecutionHistoryCliqueNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewGTMioShaderExecutionHistoryCliqueNodeWithCliqueUscParent(clique *GTMioUSCCliqueMetadata, usc objectivec.IObject, parent objectivec.IObject) GTMioShaderExecutionHistoryCliqueNode {
	instance := getGTMioShaderExecutionHistoryCliqueNodeClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithClique:usc:parent:"), unsafe.Pointer(clique), usc, parent)
	return GTMioShaderExecutionHistoryCliqueNodeFromID(rv)
}

func NewGTMioShaderExecutionHistoryCliqueNodeWithTypeParent(type_ uint32, parent objectivec.IObject) GTMioShaderExecutionHistoryCliqueNode {
	instance := getGTMioShaderExecutionHistoryCliqueNodeClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithType:parent:"), type_, parent)
	return GTMioShaderExecutionHistoryCliqueNodeFromID(rv)
}

func (g GTMioShaderExecutionHistoryCliqueNode) InitWithCliqueUscParent(clique *GTMioUSCCliqueMetadata, usc objectivec.IObject, parent objectivec.IObject) GTMioShaderExecutionHistoryCliqueNode {
	rv := objc.SendIfResponds[GTMioShaderExecutionHistoryCliqueNode](g.ID, objc.Sel("initWithClique:usc:parent:"), unsafe.Pointer(clique), usc, parent)
	return rv
}

func (g GTMioShaderExecutionHistoryCliqueNode) Clique() *GTMioUSCCliqueMetadata {
	rv := objc.SendIfResponds[unsafe.Pointer](g.ID, objc.Sel("clique"))
	return (*GTMioUSCCliqueMetadata)(rv)
}
func (g GTMioShaderExecutionHistoryCliqueNode) Usc() IGTMioUSCTraceData {
	rv := objc.SendIfResponds[objc.ID](g.ID, objc.Sel("usc"))
	return GTMioUSCTraceDataFromID(objc.ID(rv))
}
