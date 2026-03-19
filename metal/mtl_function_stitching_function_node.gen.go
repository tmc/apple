// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MTLFunctionStitchingFunctionNode] class.
var (
	_MTLFunctionStitchingFunctionNodeClass     MTLFunctionStitchingFunctionNodeClass
	_MTLFunctionStitchingFunctionNodeClassOnce sync.Once
)

func getMTLFunctionStitchingFunctionNodeClass() MTLFunctionStitchingFunctionNodeClass {
	_MTLFunctionStitchingFunctionNodeClassOnce.Do(func() {
		_MTLFunctionStitchingFunctionNodeClass = MTLFunctionStitchingFunctionNodeClass{class: objc.GetClass("MTLFunctionStitchingFunctionNode")}
	})
	return _MTLFunctionStitchingFunctionNodeClass
}

// GetMTLFunctionStitchingFunctionNodeClass returns the class object for MTLFunctionStitchingFunctionNode.
func GetMTLFunctionStitchingFunctionNodeClass() MTLFunctionStitchingFunctionNodeClass {
	return getMTLFunctionStitchingFunctionNodeClass()
}

type MTLFunctionStitchingFunctionNodeClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MTLFunctionStitchingFunctionNodeClass) Alloc() MTLFunctionStitchingFunctionNode {
	rv := objc.Send[MTLFunctionStitchingFunctionNode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A call graph node that describes a function call and its inputs.
//
// # Overview
// 
// When the Metal device object evaluates the function graph to compile the
// stitched function, it evaluates the nodes stored in the [MTLFunctionStitchingFunctionNode.Arguments]
// property that it hasn’t already evaluated, and then calls the function
// specified by [MTLFunctionStitchingFunctionNode.Name] to generate the node’s output.
// 
// If the function has side effects on the input data, use the
// [MTLFunctionStitchingFunctionNode.ControlDependencies] property on other nodes to specify whether the Metal
// device object needs to evaluate this node first.
//
// # Initializing a function node
//
//   - [MTLFunctionStitchingFunctionNode.InitWithNameArgumentsControlDependencies]: Creates a new function node.
//
// # Configuring a function node
//
//   - [MTLFunctionStitchingFunctionNode.Name]: The name of the function to call.
//   - [MTLFunctionStitchingFunctionNode.SetName]
//   - [MTLFunctionStitchingFunctionNode.Arguments]: An ordered list of the nodes that provide the function’s arguments.
//   - [MTLFunctionStitchingFunctionNode.SetArguments]
//   - [MTLFunctionStitchingFunctionNode.ControlDependencies]: The list of nodes that need to execute before executing the node.
//   - [MTLFunctionStitchingFunctionNode.SetControlDependencies]
//
// See: https://developer.apple.com/documentation/Metal/MTLFunctionStitchingFunctionNode
type MTLFunctionStitchingFunctionNode struct {
	objectivec.Object
}

// MTLFunctionStitchingFunctionNodeFromID constructs a [MTLFunctionStitchingFunctionNode] from an objc.ID.
//
// A call graph node that describes a function call and its inputs.
func MTLFunctionStitchingFunctionNodeFromID(id objc.ID) MTLFunctionStitchingFunctionNode {
	return MTLFunctionStitchingFunctionNode{objectivec.Object{ID: id}}
}
// NOTE: MTLFunctionStitchingFunctionNode adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MTLFunctionStitchingFunctionNode] class.
//
// # Initializing a function node
//
//   - [IMTLFunctionStitchingFunctionNode.InitWithNameArgumentsControlDependencies]: Creates a new function node.
//
// # Configuring a function node
//
//   - [IMTLFunctionStitchingFunctionNode.Name]: The name of the function to call.
//   - [IMTLFunctionStitchingFunctionNode.SetName]
//   - [IMTLFunctionStitchingFunctionNode.Arguments]: An ordered list of the nodes that provide the function’s arguments.
//   - [IMTLFunctionStitchingFunctionNode.SetArguments]
//   - [IMTLFunctionStitchingFunctionNode.ControlDependencies]: The list of nodes that need to execute before executing the node.
//   - [IMTLFunctionStitchingFunctionNode.SetControlDependencies]
//
// See: https://developer.apple.com/documentation/Metal/MTLFunctionStitchingFunctionNode
type IMTLFunctionStitchingFunctionNode interface {
	objectivec.IObject
	MTLFunctionStitchingNode

	// Topic: Initializing a function node

	// Creates a new function node.
	InitWithNameArgumentsControlDependencies(name string, arguments []objectivec.IObject, controlDependencies []MTLFunctionStitchingFunctionNode) MTLFunctionStitchingFunctionNode

	// Topic: Configuring a function node

	// The name of the function to call.
	Name() string
	SetName(value string)
	// An ordered list of the nodes that provide the function’s arguments.
	Arguments() []objectivec.IObject
	SetArguments(value []objectivec.IObject)
	// The list of nodes that need to execute before executing the node.
	ControlDependencies() []MTLFunctionStitchingFunctionNode
	SetControlDependencies(value []MTLFunctionStitchingFunctionNode)
}

// Init initializes the instance.
func (f MTLFunctionStitchingFunctionNode) Init() MTLFunctionStitchingFunctionNode {
	rv := objc.Send[MTLFunctionStitchingFunctionNode](f.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f MTLFunctionStitchingFunctionNode) Autorelease() MTLFunctionStitchingFunctionNode {
	rv := objc.Send[MTLFunctionStitchingFunctionNode](f.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTLFunctionStitchingFunctionNode creates a new MTLFunctionStitchingFunctionNode instance.
func NewMTLFunctionStitchingFunctionNode() MTLFunctionStitchingFunctionNode {
	class := getMTLFunctionStitchingFunctionNodeClass()
	rv := objc.Send[MTLFunctionStitchingFunctionNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a new function node.
//
// name: The name of the function to call.
//
// arguments: An ordered list of the nodes that provide the function’s arguments.
//
// controlDependencies: The list of nodes that need to run before executing this node.
//
// See: https://developer.apple.com/documentation/Metal/MTLFunctionStitchingFunctionNode/init(name:arguments:controlDependencies:)
func NewFunctionStitchingFunctionNodeWithNameArgumentsControlDependencies(name string, arguments []objectivec.IObject, controlDependencies []MTLFunctionStitchingFunctionNode) MTLFunctionStitchingFunctionNode {
	instance := getMTLFunctionStitchingFunctionNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithName:arguments:controlDependencies:"), objc.String(name), objectivec.IObjectSliceToNSArray(arguments), objectivec.IObjectSliceToNSArray(controlDependencies))
	return MTLFunctionStitchingFunctionNodeFromID(rv)
}

// Creates a new function node.
//
// name: The name of the function to call.
//
// arguments: An ordered list of the nodes that provide the function’s arguments.
//
// controlDependencies: The list of nodes that need to run before executing this node.
//
// See: https://developer.apple.com/documentation/Metal/MTLFunctionStitchingFunctionNode/init(name:arguments:controlDependencies:)
func (f MTLFunctionStitchingFunctionNode) InitWithNameArgumentsControlDependencies(name string, arguments []objectivec.IObject, controlDependencies []MTLFunctionStitchingFunctionNode) MTLFunctionStitchingFunctionNode {
	rv := objc.Send[MTLFunctionStitchingFunctionNode](f.ID, objc.Sel("initWithName:arguments:controlDependencies:"), objc.String(name), objectivec.IObjectSliceToNSArray(arguments), objectivec.IObjectSliceToNSArray(controlDependencies))
	return rv
}

// The name of the function to call.
//
// # Discussion
// 
// The name needs to match one of the functions in the stitched library
// descriptor’s [Functions] property.
//
// See: https://developer.apple.com/documentation/Metal/MTLFunctionStitchingFunctionNode/name
func (f MTLFunctionStitchingFunctionNode) Name() string {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("name"))
	return foundation.NSStringFromID(rv).String()
}
func (f MTLFunctionStitchingFunctionNode) SetName(value string) {
	objc.Send[struct{}](f.ID, objc.Sel("setName:"), objc.String(value))
}
// An ordered list of the nodes that provide the function’s arguments.
//
// # Discussion
// 
// Each node’s output data types needs to match the input data type of the
// matching argument.
//
// See: https://developer.apple.com/documentation/Metal/MTLFunctionStitchingFunctionNode/arguments
func (f MTLFunctionStitchingFunctionNode) Arguments() []objectivec.IObject {
	rv := objc.Send[[]objc.ID](f.ID, objc.Sel("arguments"))
	return objc.ConvertSlice(rv, func(id objc.ID) objectivec.IObject {
		return objectivec.Object{ID: id}
	})
}
func (f MTLFunctionStitchingFunctionNode) SetArguments(value []objectivec.IObject) {
	objc.Send[struct{}](f.ID, objc.Sel("setArguments:"), objectivec.IObjectSliceToNSArray(value))
}
// The list of nodes that need to execute before executing the node.
//
// # Discussion
// 
// When a stitched function calls functions that have side effects on their
// input data, you often need the GPU to execute functions in a specific
// order. In such cases, use the [ControlDependencies] property to specify
// which nodes need to run before executing this node.
//
// See: https://developer.apple.com/documentation/Metal/MTLFunctionStitchingFunctionNode/controlDependencies
func (f MTLFunctionStitchingFunctionNode) ControlDependencies() []MTLFunctionStitchingFunctionNode {
	rv := objc.Send[[]objc.ID](f.ID, objc.Sel("controlDependencies"))
	return objc.ConvertSlice(rv, func(id objc.ID) MTLFunctionStitchingFunctionNode {
		return MTLFunctionStitchingFunctionNodeFromID(id)
	})
}
func (f MTLFunctionStitchingFunctionNode) SetControlDependencies(value []MTLFunctionStitchingFunctionNode) {
	objc.Send[struct{}](f.ID, objc.Sel("setControlDependencies:"), objectivec.IObjectSliceToNSArray(value))
}

			// Protocol methods for MTLFunctionStitchingNode
			

