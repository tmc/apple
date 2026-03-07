// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSNNStateNode] class.
var (
	_MPSNNStateNodeClass     MPSNNStateNodeClass
	_MPSNNStateNodeClassOnce sync.Once
)

func getMPSNNStateNodeClass() MPSNNStateNodeClass {
	_MPSNNStateNodeClassOnce.Do(func() {
		_MPSNNStateNodeClass = MPSNNStateNodeClass{class: objc.GetClass("MPSNNStateNode")}
	})
	return _MPSNNStateNodeClass
}

// GetMPSNNStateNodeClass returns the class object for MPSNNStateNode.
func GetMPSNNStateNodeClass() MPSNNStateNodeClass {
	return getMPSNNStateNodeClass()
}

type MPSNNStateNodeClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNStateNodeClass) Alloc() MPSNNStateNode {
	rv := objc.Send[MPSNNStateNode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}







//
// # Methods
//
//   - [MPSNNStateNode.DebugQuickLookObject]
//   - [MPSNNStateNode.ExportFromGraph]
//   - [MPSNNStateNode.SetExportFromGraph]
//   - [MPSNNStateNode.Handle]
//   - [MPSNNStateNode.SetHandle]
//   - [MPSNNStateNode.SynchronizeResource]
//   - [MPSNNStateNode.SetSynchronizeResource]
//   - [MPSNNStateNode.InitWithParent]
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNStateNode
type MPSNNStateNode struct {
	objectivec.Object
}

// MPSNNStateNodeFromID constructs a [MPSNNStateNode] from an objc.ID.
func MPSNNStateNodeFromID(id objc.ID) MPSNNStateNode {
	return MPSNNStateNode{objectivec.Object{id}}
}
// Ensure MPSNNStateNode implements IMPSNNStateNode.
var _ IMPSNNStateNode = MPSNNStateNode{}





// An interface definition for the [MPSNNStateNode] class.
//
// # Methods
//
//   - [IMPSNNStateNode.DebugQuickLookObject]
//   - [IMPSNNStateNode.ExportFromGraph]
//   - [IMPSNNStateNode.SetExportFromGraph]
//   - [IMPSNNStateNode.Handle]
//   - [IMPSNNStateNode.SetHandle]
//   - [IMPSNNStateNode.SynchronizeResource]
//   - [IMPSNNStateNode.SetSynchronizeResource]
//   - [IMPSNNStateNode.InitWithParent]
//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNStateNode
type IMPSNNStateNode interface {
	objectivec.IObject

	// Topic: Methods

	DebugQuickLookObject() objectivec.IObject
	ExportFromGraph() bool
	SetExportFromGraph(value bool)
	Handle() unsafe.Pointer
	SetHandle(value unsafe.Pointer)
	SynchronizeResource() bool
	SetSynchronizeResource(value bool)
	InitWithParent(parent objectivec.IObject) MPSNNStateNode
}





// Init initializes the instance.
func (s MPSNNStateNode) Init() MPSNNStateNode {
	rv := objc.Send[MPSNNStateNode](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s MPSNNStateNode) Autorelease() MPSNNStateNode {
	rv := objc.Send[MPSNNStateNode](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNStateNode creates a new MPSNNStateNode instance.
func NewMPSNNStateNode() MPSNNStateNode {
	class := getMPSNNStateNodeClass()
	rv := objc.Send[MPSNNStateNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}






//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNStateNode/initWithParent:
func NewStateNodeWithParent(parent objectivec.IObject) MPSNNStateNode {
	instance := getMPSNNStateNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithParent:"), parent)
	return MPSNNStateNodeFromID(rv)
}







// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNStateNode/debugQuickLookObject
func (s MPSNNStateNode) DebugQuickLookObject() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("debugQuickLookObject"))
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNStateNode/initWithParent:
func (s MPSNNStateNode) InitWithParent(parent objectivec.IObject) MPSNNStateNode {
	rv := objc.Send[MPSNNStateNode](s.ID, objc.Sel("initWithParent:"), parent)
	return rv
}












// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNStateNode/exportFromGraph
func (s MPSNNStateNode) ExportFromGraph() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("exportFromGraph"))
	return rv
}
func (s MPSNNStateNode) SetExportFromGraph(value bool) {
	objc.Send[struct{}](s.ID, objc.Sel("setExportFromGraph:"), value)
}



// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNStateNode/handle
func (s MPSNNStateNode) Handle() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s.ID, objc.Sel("handle"))
	return rv
}
func (s MPSNNStateNode) SetHandle(value unsafe.Pointer) {
	objc.Send[struct{}](s.ID, objc.Sel("setHandle:"), value)
}



// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNStateNode/synchronizeResource
func (s MPSNNStateNode) SynchronizeResource() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("synchronizeResource"))
	return rv
}
func (s MPSNNStateNode) SetSynchronizeResource(value bool) {
	objc.Send[struct{}](s.ID, objc.Sel("setSynchronizeResource:"), value)
}

















