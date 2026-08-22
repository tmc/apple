// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
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

// Class returns the underlying Objective-C class pointer.
func (mc MPSNNStateNodeClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSNNStateNodeClass) Alloc() MPSNNStateNode {
	rv := objc.Send[MPSNNStateNode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A placeholder node denoting the position in the graph of a state object.
//
// # Instance Properties
//
//   - [MPSNNStateNode.Handle]
//   - [MPSNNStateNode.SetHandle]
//   - [MPSNNStateNode.ExportFromGraph]
//   - [MPSNNStateNode.SetExportFromGraph]
//   - [MPSNNStateNode.SynchronizeResource]
//   - [MPSNNStateNode.SetSynchronizeResource]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNStateNode
type MPSNNStateNode struct {
	objectivec.Object
}

// MPSNNStateNodeFromID constructs a [MPSNNStateNode] from an objc.ID.
//
// A placeholder node denoting the position in the graph of a state object.
func MPSNNStateNodeFromID(id objc.ID) MPSNNStateNode {
	return MPSNNStateNode{objectivec.Object{ID: id}}
}

// NOTE: MPSNNStateNode adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSNNStateNode] class.
//
// # Instance Properties
//
//   - [IMPSNNStateNode.Handle]
//   - [IMPSNNStateNode.SetHandle]
//   - [IMPSNNStateNode.ExportFromGraph]
//   - [IMPSNNStateNode.SetExportFromGraph]
//   - [IMPSNNStateNode.SynchronizeResource]
//   - [IMPSNNStateNode.SetSynchronizeResource]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNStateNode
type IMPSNNStateNode interface {
	objectivec.IObject

	// Topic: Instance Properties

	Handle() MPSHandle
	SetHandle(value MPSHandle)
	ExportFromGraph() bool
	SetExportFromGraph(value bool)
	SynchronizeResource() bool
	SetSynchronizeResource(value bool)
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

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNStateNode/handle
func (s MPSNNStateNode) Handle() MPSHandle {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("handle"))
	return MPSHandleObjectFromID(rv)
}
func (s MPSNNStateNode) SetHandle(value MPSHandle) {
	objc.Send[struct{}](s.ID, objc.Sel("setHandle:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNStateNode/exportFromGraph
func (s MPSNNStateNode) ExportFromGraph() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("exportFromGraph"))
	return rv
}
func (s MPSNNStateNode) SetExportFromGraph(value bool) {
	objc.Send[struct{}](s.ID, objc.Sel("setExportFromGraph:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNStateNode/synchronizeResource
func (s MPSNNStateNode) SynchronizeResource() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("synchronizeResource"))
	return rv
}
func (s MPSNNStateNode) SetSynchronizeResource(value bool) {
	objc.Send[struct{}](s.ID, objc.Sel("setSynchronizeResource:"), value)
}
