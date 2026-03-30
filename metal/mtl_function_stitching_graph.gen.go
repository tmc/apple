// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MTLFunctionStitchingGraph] class.
var (
	_MTLFunctionStitchingGraphClass     MTLFunctionStitchingGraphClass
	_MTLFunctionStitchingGraphClassOnce sync.Once
)

func getMTLFunctionStitchingGraphClass() MTLFunctionStitchingGraphClass {
	_MTLFunctionStitchingGraphClassOnce.Do(func() {
		_MTLFunctionStitchingGraphClass = MTLFunctionStitchingGraphClass{class: objc.GetClass("MTLFunctionStitchingGraph")}
	})
	return _MTLFunctionStitchingGraphClass
}

// GetMTLFunctionStitchingGraphClass returns the class object for MTLFunctionStitchingGraph.
func GetMTLFunctionStitchingGraphClass() MTLFunctionStitchingGraphClass {
	return getMTLFunctionStitchingGraphClass()
}

type MTLFunctionStitchingGraphClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MTLFunctionStitchingGraphClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MTLFunctionStitchingGraphClass) Alloc() MTLFunctionStitchingGraph {
	rv := objc.Send[MTLFunctionStitchingGraph](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A description of a new stitched function.
//
// # Overview
//
// An [MTLFunctionStitchingGraph] instance describes the function graph for a
// stitched function. A is a visible function you create by composing other
// Metal shader functions together in a function graph. A function stitching
// graph contains nodes for the function’s arguments and any functions it
// calls in the implementation. Data flows from the arguments to the end of
// the graph until the stitched function evaluates all of the graph’s nodes.
//
// The graph in the figure below constructs a new function that adds numbers
// from two source arrays, storing the result in a third array. The
// function’s parameters are pointers to the source and destination arrays,
// and an index for performing the array lookup. The graph uses three separate
// MSL functions to construct the stitched function: a function to look up a
// value from an array, a function that adds two numbers together, and a
// function that stores a value to an array.
//
// [media-3842304]
//
// Create an [MTLFunctionStitchingGraph] instance for each stitched function
// you want to create. Configure its properties to describe the new function
// and the nodes that define its behavior, as described below. To create a new
// library with these stitched functions, see [MTLStitchedLibraryDescriptor].
//
// # Configuring a function stitching graph
//
// To create a valid stitched function, the function stitching graph and
// shader code need to meet some requirements:
//
// - Implement the MSL functions that you use to create the new function,
// adding the `stitchable` attribute to each. Stitchable functions are visible
// functions that you can also use in a function graph. Stitchable functions
// may require the compiler to do additional work or emit larger instance
// code, so mark functions as stitchable only when necessary. - Declare the
// stitched function’s name and signature in a header file to include in any
// shader code that calls the new function. Alternatively, you can add the
// function to a function table with a matching type and pass the function
// table as an argument. - Create an [MTLFunctionStitchingInputNode] node for
// each of the function’s arguments, specifying which parameter each node
// references. The output type of each input node is the type of that
// parameter in your function signature. - Create an
// [MTLFunctionStitchingFunctionNode] for each function the implementation
// calls. A function node’s output type is the return type of the MSL
// function. - Make sure the output types of each node match the types of the
// node they pass to. For example, if a function takes a `float` parameter,
// the node that provides that data need to output a `float` value. If you
// don’t match the types correctly, Metal doesn’t define the behavior of
// the resulting function. - Create an array from the node instances and
// assign it to the [MTLFunctionStitchingGraph.Nodes] property. - If the function produces an output,
// create another node and assign it to the [MTLFunctionStitchingGraph.OutputNode] property. The output
// type of this node needs to match the function’s return type. Don’t
// include this node in the array of nodes you assign to the [MTLFunctionStitchingGraph.Nodes] property.
//
// The MSL code below implements the functions in the example graph above, as
// well as the function’s signature:
//
// The following code creates the graph above:
//
// # Initializing a function graph
//
//   - [MTLFunctionStitchingGraph.InitWithFunctionNameNodesOutputNodeAttributes]: Creates a description of a new function call graph.
//
// # Configuring a function graph
//
//   - [MTLFunctionStitchingGraph.FunctionName]: The name of the new stitched function.
//   - [MTLFunctionStitchingGraph.SetFunctionName]
//   - [MTLFunctionStitchingGraph.Nodes]: The nodes in the function’s call graph.
//   - [MTLFunctionStitchingGraph.SetNodes]
//   - [MTLFunctionStitchingGraph.OutputNode]: The node with the output that’s the output of the new stitched function.
//   - [MTLFunctionStitchingGraph.SetOutputNode]
//   - [MTLFunctionStitchingGraph.Attributes]: A list of attributes to configure how the Metal device object generates the new stitched function.
//   - [MTLFunctionStitchingGraph.SetAttributes]
//
// See: https://developer.apple.com/documentation/Metal/MTLFunctionStitchingGraph
type MTLFunctionStitchingGraph struct {
	objectivec.Object
}

// MTLFunctionStitchingGraphFromID constructs a [MTLFunctionStitchingGraph] from an objc.ID.
//
// A description of a new stitched function.
func MTLFunctionStitchingGraphFromID(id objc.ID) MTLFunctionStitchingGraph {
	return MTLFunctionStitchingGraph{objectivec.Object{ID: id}}
}

// NOTE: MTLFunctionStitchingGraph adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MTLFunctionStitchingGraph] class.
//
// # Initializing a function graph
//
//   - [IMTLFunctionStitchingGraph.InitWithFunctionNameNodesOutputNodeAttributes]: Creates a description of a new function call graph.
//
// # Configuring a function graph
//
//   - [IMTLFunctionStitchingGraph.FunctionName]: The name of the new stitched function.
//   - [IMTLFunctionStitchingGraph.SetFunctionName]
//   - [IMTLFunctionStitchingGraph.Nodes]: The nodes in the function’s call graph.
//   - [IMTLFunctionStitchingGraph.SetNodes]
//   - [IMTLFunctionStitchingGraph.OutputNode]: The node with the output that’s the output of the new stitched function.
//   - [IMTLFunctionStitchingGraph.SetOutputNode]
//   - [IMTLFunctionStitchingGraph.Attributes]: A list of attributes to configure how the Metal device object generates the new stitched function.
//   - [IMTLFunctionStitchingGraph.SetAttributes]
//
// See: https://developer.apple.com/documentation/Metal/MTLFunctionStitchingGraph
type IMTLFunctionStitchingGraph interface {
	objectivec.IObject

	// Topic: Initializing a function graph

	// Creates a description of a new function call graph.
	InitWithFunctionNameNodesOutputNodeAttributes(functionName string, nodes []MTLFunctionStitchingFunctionNode, outputNode IMTLFunctionStitchingFunctionNode, attributes []objectivec.IObject) MTLFunctionStitchingGraph

	// Topic: Configuring a function graph

	// The name of the new stitched function.
	FunctionName() string
	SetFunctionName(value string)
	// The nodes in the function’s call graph.
	Nodes() []MTLFunctionStitchingFunctionNode
	SetNodes(value []MTLFunctionStitchingFunctionNode)
	// The node with the output that’s the output of the new stitched function.
	OutputNode() IMTLFunctionStitchingFunctionNode
	SetOutputNode(value IMTLFunctionStitchingFunctionNode)
	// A list of attributes to configure how the Metal device object generates the new stitched function.
	Attributes() []objectivec.IObject
	SetAttributes(value []objectivec.IObject)
}

// Init initializes the instance.
func (f MTLFunctionStitchingGraph) Init() MTLFunctionStitchingGraph {
	rv := objc.Send[MTLFunctionStitchingGraph](f.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f MTLFunctionStitchingGraph) Autorelease() MTLFunctionStitchingGraph {
	rv := objc.Send[MTLFunctionStitchingGraph](f.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTLFunctionStitchingGraph creates a new MTLFunctionStitchingGraph instance.
func NewMTLFunctionStitchingGraph() MTLFunctionStitchingGraph {
	class := getMTLFunctionStitchingGraphClass()
	rv := objc.Send[MTLFunctionStitchingGraph](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a description of a new function call graph.
//
// functionName: The name of the new function.
//
// nodes: The nodes in the function’s call graph.
//
// outputNode: The node whose output is the output of the new stitched function.
//
// attributes: A list of attributes used to generate the new stitched function.
//
// See: https://developer.apple.com/documentation/Metal/MTLFunctionStitchingGraph/init(functionName:nodes:outputNode:attributes:)
func NewFunctionStitchingGraphWithFunctionNameNodesOutputNodeAttributes(functionName string, nodes []MTLFunctionStitchingFunctionNode, outputNode IMTLFunctionStitchingFunctionNode, attributes []objectivec.IObject) MTLFunctionStitchingGraph {
	instance := getMTLFunctionStitchingGraphClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFunctionName:nodes:outputNode:attributes:"), objc.String(functionName), objectivec.IObjectSliceToNSArray(nodes), outputNode, objectivec.IObjectSliceToNSArray(attributes))
	return MTLFunctionStitchingGraphFromID(rv)
}

// Creates a description of a new function call graph.
//
// functionName: The name of the new function.
//
// nodes: The nodes in the function’s call graph.
//
// outputNode: The node whose output is the output of the new stitched function.
//
// attributes: A list of attributes used to generate the new stitched function.
//
// See: https://developer.apple.com/documentation/Metal/MTLFunctionStitchingGraph/init(functionName:nodes:outputNode:attributes:)
func (f MTLFunctionStitchingGraph) InitWithFunctionNameNodesOutputNodeAttributes(functionName string, nodes []MTLFunctionStitchingFunctionNode, outputNode IMTLFunctionStitchingFunctionNode, attributes []objectivec.IObject) MTLFunctionStitchingGraph {
	rv := objc.Send[MTLFunctionStitchingGraph](f.ID, objc.Sel("initWithFunctionName:nodes:outputNode:attributes:"), objc.String(functionName), objectivec.IObjectSliceToNSArray(nodes), outputNode, objectivec.IObjectSliceToNSArray(attributes))
	return rv
}

// The name of the new stitched function.
//
// See: https://developer.apple.com/documentation/Metal/MTLFunctionStitchingGraph/functionName
func (f MTLFunctionStitchingGraph) FunctionName() string {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("functionName"))
	return foundation.NSStringFromID(rv).String()
}
func (f MTLFunctionStitchingGraph) SetFunctionName(value string) {
	objc.Send[struct{}](f.ID, objc.Sel("setFunctionName:"), objc.String(value))
}

// The nodes in the function’s call graph.
//
// See: https://developer.apple.com/documentation/Metal/MTLFunctionStitchingGraph/nodes
func (f MTLFunctionStitchingGraph) Nodes() []MTLFunctionStitchingFunctionNode {
	rv := objc.Send[[]objc.ID](f.ID, objc.Sel("nodes"))
	return objc.ConvertSlice(rv, func(id objc.ID) MTLFunctionStitchingFunctionNode {
		return MTLFunctionStitchingFunctionNodeFromID(id)
	})
}
func (f MTLFunctionStitchingGraph) SetNodes(value []MTLFunctionStitchingFunctionNode) {
	objc.Send[struct{}](f.ID, objc.Sel("setNodes:"), objectivec.IObjectSliceToNSArray(value))
}

// The node with the output that’s the output of the new stitched function.
//
// # Discussion
//
// The output type of the node needs to match the result type in the stitched
// function’s declaration.
//
// See: https://developer.apple.com/documentation/Metal/MTLFunctionStitchingGraph/outputNode
func (f MTLFunctionStitchingGraph) OutputNode() IMTLFunctionStitchingFunctionNode {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("outputNode"))
	return MTLFunctionStitchingFunctionNodeFromID(objc.ID(rv))
}
func (f MTLFunctionStitchingGraph) SetOutputNode(value IMTLFunctionStitchingFunctionNode) {
	objc.Send[struct{}](f.ID, objc.Sel("setOutputNode:"), value)
}

// A list of attributes to configure how the Metal device object generates the
// new stitched function.
//
// See: https://developer.apple.com/documentation/Metal/MTLFunctionStitchingGraph/attributes
func (f MTLFunctionStitchingGraph) Attributes() []objectivec.IObject {
	rv := objc.Send[[]objc.ID](f.ID, objc.Sel("attributes"))
	return objc.ConvertSlice(rv, func(id objc.ID) objectivec.IObject {
		return objectivec.Object{ID: id}
	})
}
func (f MTLFunctionStitchingGraph) SetAttributes(value []objectivec.IObject) {
	objc.Send[struct{}](f.ID, objc.Sel("setAttributes:"), objectivec.IObjectSliceToNSArray(value))
}
