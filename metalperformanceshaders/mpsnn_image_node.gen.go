// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
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

// Class returns the underlying Objective-C class pointer.
func (mc MPSNNImageNodeClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNImageNodeClass) Alloc() MPSNNImageNode {
	rv := objc.Send[MPSNNImageNode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A placeholder node denoting the position of a neural network image in a
// graph.
//
// # Initializers
//
//   - [MPSNNImageNode.InitWithHandle]
//
// # Instance Properties
//
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
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNImageNode
type MPSNNImageNode struct {
	objectivec.Object
}

// MPSNNImageNodeFromID constructs a [MPSNNImageNode] from an objc.ID.
//
// A placeholder node denoting the position of a neural network image in a
// graph.
func MPSNNImageNodeFromID(id objc.ID) MPSNNImageNode {
	return MPSNNImageNode{objectivec.Object{ID: id}}
}

// NOTE: MPSNNImageNode adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSNNImageNode] class.
//
// # Initializers
//
//   - [IMPSNNImageNode.InitWithHandle]
//
// # Instance Properties
//
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
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNImageNode
type IMPSNNImageNode interface {
	objectivec.IObject

	// Topic: Initializers

	InitWithHandle(handle objectivec.NSObject) MPSNNImageNode

	// Topic: Instance Properties

	ExportFromGraph() bool
	SetExportFromGraph(value bool)
	Format() MPSImageFeatureChannelFormat
	SetFormat(value MPSImageFeatureChannelFormat)
	Handle() MPSHandle
	SetHandle(value MPSHandle)
	ImageAllocator() MPSImageAllocator
	SetImageAllocator(value MPSImageAllocator)
	StopGradient() bool
	SetStopGradient(value bool)
	SynchronizeResource() bool
	SetSynchronizeResource(value bool)
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

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNImageNode/init(handle:)
func NewImageNodeWithHandle(handle objectivec.NSObject) MPSNNImageNode {
	instance := getMPSNNImageNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithHandle:"), handle)
	return MPSNNImageNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNImageNode/init(handle:)
func (i MPSNNImageNode) InitWithHandle(handle objectivec.NSObject) MPSNNImageNode {
	rv := objc.Send[MPSNNImageNode](i.ID, objc.Sel("initWithHandle:"), handle)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNImageNode/exportedNode(with:)
func (_MPSNNImageNodeClass MPSNNImageNodeClass) ExportedNodeWithHandle(handle objectivec.NSObject) MPSNNImageNode {
	rv := objc.Send[objc.ID](objc.ID(_MPSNNImageNodeClass.class), objc.Sel("exportedNodeWithHandle:"), handle)
	return MPSNNImageNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNImageNode/nodeWithHandle:
func (_MPSNNImageNodeClass MPSNNImageNodeClass) NodeWithHandle(handle objectivec.NSObject) MPSNNImageNode {
	rv := objc.Send[objc.ID](objc.ID(_MPSNNImageNodeClass.class), objc.Sel("nodeWithHandle:"), handle)
	return MPSNNImageNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNImageNode/exportFromGraph
func (i MPSNNImageNode) ExportFromGraph() bool {
	rv := objc.Send[bool](i.ID, objc.Sel("exportFromGraph"))
	return rv
}
func (i MPSNNImageNode) SetExportFromGraph(value bool) {
	objc.Send[struct{}](i.ID, objc.Sel("setExportFromGraph:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNImageNode/format
func (i MPSNNImageNode) Format() MPSImageFeatureChannelFormat {
	rv := objc.Send[MPSImageFeatureChannelFormat](i.ID, objc.Sel("format"))
	return MPSImageFeatureChannelFormat(rv)
}
func (i MPSNNImageNode) SetFormat(value MPSImageFeatureChannelFormat) {
	objc.Send[struct{}](i.ID, objc.Sel("setFormat:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNImageNode/handle
func (i MPSNNImageNode) Handle() MPSHandle {
	rv := objc.Send[objc.ID](i.ID, objc.Sel("handle"))
	return MPSHandleObjectFromID(rv)
}
func (i MPSNNImageNode) SetHandle(value MPSHandle) {
	objc.Send[struct{}](i.ID, objc.Sel("setHandle:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNImageNode/imageAllocator
func (i MPSNNImageNode) ImageAllocator() MPSImageAllocator {
	rv := objc.Send[objc.ID](i.ID, objc.Sel("imageAllocator"))
	return MPSImageAllocatorObjectFromID(rv)
}
func (i MPSNNImageNode) SetImageAllocator(value MPSImageAllocator) {
	objc.Send[struct{}](i.ID, objc.Sel("setImageAllocator:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNImageNode/stopGradient
func (i MPSNNImageNode) StopGradient() bool {
	rv := objc.Send[bool](i.ID, objc.Sel("stopGradient"))
	return rv
}
func (i MPSNNImageNode) SetStopGradient(value bool) {
	objc.Send[struct{}](i.ID, objc.Sel("setStopGradient:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNImageNode/synchronizeResource
func (i MPSNNImageNode) SynchronizeResource() bool {
	rv := objc.Send[bool](i.ID, objc.Sel("synchronizeResource"))
	return rv
}
func (i MPSNNImageNode) SetSynchronizeResource(value bool) {
	objc.Send[struct{}](i.ID, objc.Sel("setSynchronizeResource:"), value)
}
