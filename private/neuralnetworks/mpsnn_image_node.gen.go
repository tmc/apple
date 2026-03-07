// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSNNImageNode] class.
var (
	_MPSNNImageNodeClass     MPSNNImageNodeClass
	_MPSNNImageNodeClassOnce sync.Once
)

func getMPSNNImageNodeClass() MPSNNImageNodeClass {
	_MPSNNImageNodeClassOnce.Do(func() {
		_MPSNNImageNodeClass = MPSNNImageNodeClass{class: objc.GetClass("MPSNNImageNode")}
	})
	return _MPSNNImageNodeClass
}

// GetMPSNNImageNodeClass returns the class object for MPSNNImageNode.
func GetMPSNNImageNodeClass() MPSNNImageNodeClass {
	return getMPSNNImageNodeClass()
}

type MPSNNImageNodeClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNImageNodeClass) Alloc() MPSNNImageNode {
	rv := objc.Send[MPSNNImageNode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}







//
// # Methods
//
//   - [MPSNNImageNode.DebugQuickLookObject]
//   - [MPSNNImageNode.ExportFromGraph]
//   - [MPSNNImageNode.SetExportFromGraph]
//   - [MPSNNImageNode.Format]
//   - [MPSNNImageNode.SetFormat]
//   - [MPSNNImageNode.Handle]
//   - [MPSNNImageNode.SetHandle]
//   - [MPSNNImageNode.ImageAllocator]
//   - [MPSNNImageNode.SetImageAllocator]
//   - [MPSNNImageNode.StopGradient]
//   - [MPSNNImageNode.SetStopGradient]
//   - [MPSNNImageNode.SynchronizeResource]
//   - [MPSNNImageNode.SetSynchronizeResource]
//   - [MPSNNImageNode.InitWithHandle]
//   - [MPSNNImageNode.InitWithParent]
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNImageNode
type MPSNNImageNode struct {
	objectivec.Object
}

// MPSNNImageNodeFromID constructs a [MPSNNImageNode] from an objc.ID.
func MPSNNImageNodeFromID(id objc.ID) MPSNNImageNode {
	return MPSNNImageNode{objectivec.Object{id}}
}
// Ensure MPSNNImageNode implements IMPSNNImageNode.
var _ IMPSNNImageNode = MPSNNImageNode{}





// An interface definition for the [MPSNNImageNode] class.
//
// # Methods
//
//   - [IMPSNNImageNode.DebugQuickLookObject]
//   - [IMPSNNImageNode.ExportFromGraph]
//   - [IMPSNNImageNode.SetExportFromGraph]
//   - [IMPSNNImageNode.Format]
//   - [IMPSNNImageNode.SetFormat]
//   - [IMPSNNImageNode.Handle]
//   - [IMPSNNImageNode.SetHandle]
//   - [IMPSNNImageNode.ImageAllocator]
//   - [IMPSNNImageNode.SetImageAllocator]
//   - [IMPSNNImageNode.StopGradient]
//   - [IMPSNNImageNode.SetStopGradient]
//   - [IMPSNNImageNode.SynchronizeResource]
//   - [IMPSNNImageNode.SetSynchronizeResource]
//   - [IMPSNNImageNode.InitWithHandle]
//   - [IMPSNNImageNode.InitWithParent]
//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNImageNode
type IMPSNNImageNode interface {
	objectivec.IObject

	// Topic: Methods

	DebugQuickLookObject() objectivec.IObject
	ExportFromGraph() bool
	SetExportFromGraph(value bool)
	Format() uint64
	SetFormat(value uint64)
	Handle() unsafe.Pointer
	SetHandle(value unsafe.Pointer)
	ImageAllocator() unsafe.Pointer
	SetImageAllocator(value unsafe.Pointer)
	StopGradient() bool
	SetStopGradient(value bool)
	SynchronizeResource() bool
	SetSynchronizeResource(value bool)
	InitWithHandle(handle objectivec.IObject) MPSNNImageNode
	InitWithParent(parent objectivec.IObject) MPSNNImageNode
}





// Init initializes the instance.
func (i MPSNNImageNode) Init() MPSNNImageNode {
	rv := objc.Send[MPSNNImageNode](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i MPSNNImageNode) Autorelease() MPSNNImageNode {
	rv := objc.Send[MPSNNImageNode](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSNNImageNode creates a new MPSNNImageNode instance.
func NewMPSNNImageNode() MPSNNImageNode {
	class := getMPSNNImageNodeClass()
	rv := objc.Send[MPSNNImageNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}






//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNImageNode/initWithHandle:
func NewImageNodeWithHandle(handle objectivec.IObject) MPSNNImageNode {
	instance := getMPSNNImageNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithHandle:"), handle)
	return MPSNNImageNodeFromID(rv)
}


//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNImageNode/initWithParent:
func NewImageNodeWithParent(parent objectivec.IObject) MPSNNImageNode {
	instance := getMPSNNImageNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithParent:"), parent)
	return MPSNNImageNodeFromID(rv)
}







// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNImageNode/debugQuickLookObject
func (i MPSNNImageNode) DebugQuickLookObject() objectivec.IObject {
	rv := objc.Send[objc.ID](i.ID, objc.Sel("debugQuickLookObject"))
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNImageNode/initWithHandle:
func (i MPSNNImageNode) InitWithHandle(handle objectivec.IObject) MPSNNImageNode {
	rv := objc.Send[MPSNNImageNode](i.ID, objc.Sel("initWithHandle:"), handle)
	return rv
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNImageNode/initWithParent:
func (i MPSNNImageNode) InitWithParent(parent objectivec.IObject) MPSNNImageNode {
	rv := objc.Send[MPSNNImageNode](i.ID, objc.Sel("initWithParent:"), parent)
	return rv
}





//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNImageNode/exportedNodeWithHandle:
func (_MPSNNImageNodeClass MPSNNImageNodeClass) ExportedNodeWithHandle(handle objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MPSNNImageNodeClass.class), objc.Sel("exportedNodeWithHandle:"), handle)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNImageNode/nodeWithHandle:
func (_MPSNNImageNodeClass MPSNNImageNodeClass) NodeWithHandle(handle objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MPSNNImageNodeClass.class), objc.Sel("nodeWithHandle:"), handle)
	return objectivec.Object{ID: rv}
}








// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNImageNode/exportFromGraph
func (i MPSNNImageNode) ExportFromGraph() bool {
	rv := objc.Send[bool](i.ID, objc.Sel("exportFromGraph"))
	return rv
}
func (i MPSNNImageNode) SetExportFromGraph(value bool) {
	objc.Send[struct{}](i.ID, objc.Sel("setExportFromGraph:"), value)
}



// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNImageNode/format
func (i MPSNNImageNode) Format() uint64 {
	rv := objc.Send[uint64](i.ID, objc.Sel("format"))
	return rv
}
func (i MPSNNImageNode) SetFormat(value uint64) {
	objc.Send[struct{}](i.ID, objc.Sel("setFormat:"), value)
}



// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNImageNode/handle
func (i MPSNNImageNode) Handle() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i.ID, objc.Sel("handle"))
	return rv
}
func (i MPSNNImageNode) SetHandle(value unsafe.Pointer) {
	objc.Send[struct{}](i.ID, objc.Sel("setHandle:"), value)
}



// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNImageNode/imageAllocator
func (i MPSNNImageNode) ImageAllocator() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i.ID, objc.Sel("imageAllocator"))
	return rv
}
func (i MPSNNImageNode) SetImageAllocator(value unsafe.Pointer) {
	objc.Send[struct{}](i.ID, objc.Sel("setImageAllocator:"), value)
}



// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNImageNode/stopGradient
func (i MPSNNImageNode) StopGradient() bool {
	rv := objc.Send[bool](i.ID, objc.Sel("stopGradient"))
	return rv
}
func (i MPSNNImageNode) SetStopGradient(value bool) {
	objc.Send[struct{}](i.ID, objc.Sel("setStopGradient:"), value)
}



// See: https://developer.apple.com/documentation/NeuralNetworks/MPSNNImageNode/synchronizeResource
func (i MPSNNImageNode) SynchronizeResource() bool {
	rv := objc.Send[bool](i.ID, objc.Sel("synchronizeResource"))
	return rv
}
func (i MPSNNImageNode) SetSynchronizeResource(value bool) {
	objc.Send[struct{}](i.ID, objc.Sel("setSynchronizeResource:"), value)
}

















