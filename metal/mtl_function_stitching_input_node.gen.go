// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MTLFunctionStitchingInputNode] class.
var (
	_MTLFunctionStitchingInputNodeClass     MTLFunctionStitchingInputNodeClass
	_MTLFunctionStitchingInputNodeClassOnce sync.Once
)

func getMTLFunctionStitchingInputNodeClass() MTLFunctionStitchingInputNodeClass {
	_MTLFunctionStitchingInputNodeClassOnce.Do(func() {
		_MTLFunctionStitchingInputNodeClass = MTLFunctionStitchingInputNodeClass{class: objc.GetClass("MTLFunctionStitchingInputNode")}
	})
	return _MTLFunctionStitchingInputNodeClass
}

// GetMTLFunctionStitchingInputNodeClass returns the class object for MTLFunctionStitchingInputNode.
func GetMTLFunctionStitchingInputNodeClass() MTLFunctionStitchingInputNodeClass {
	return getMTLFunctionStitchingInputNodeClass()
}

type MTLFunctionStitchingInputNodeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MTLFunctionStitchingInputNodeClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MTLFunctionStitchingInputNodeClass) Alloc() MTLFunctionStitchingInputNode {
	rv := objc.Send[MTLFunctionStitchingInputNode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A call graph node that describes an input to the call graph.
//
// # Overview
//
// An input node contains data from one of the stitched function’s
// parameters. The output data type of an input node has the same type as the
// matching parameter.
//
// # Initializing an input node
//
//   - [MTLFunctionStitchingInputNode.InitWithArgumentIndex]: Creates a new input node.
//
// # Configuring an input node
//
//   - [MTLFunctionStitchingInputNode.ArgumentIndex]: The index in the command’s buffer argument table that declares which data to read for this input node.
//   - [MTLFunctionStitchingInputNode.SetArgumentIndex]
//
// See: https://developer.apple.com/documentation/Metal/MTLFunctionStitchingInputNode
type MTLFunctionStitchingInputNode struct {
	objectivec.Object
}

// MTLFunctionStitchingInputNodeFromID constructs a [MTLFunctionStitchingInputNode] from an objc.ID.
//
// A call graph node that describes an input to the call graph.
func MTLFunctionStitchingInputNodeFromID(id objc.ID) MTLFunctionStitchingInputNode {
	return MTLFunctionStitchingInputNode{objectivec.Object{ID: id}}
}

// NOTE: MTLFunctionStitchingInputNode adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MTLFunctionStitchingInputNode] class.
//
// # Initializing an input node
//
//   - [IMTLFunctionStitchingInputNode.InitWithArgumentIndex]: Creates a new input node.
//
// # Configuring an input node
//
//   - [IMTLFunctionStitchingInputNode.ArgumentIndex]: The index in the command’s buffer argument table that declares which data to read for this input node.
//   - [IMTLFunctionStitchingInputNode.SetArgumentIndex]
//
// See: https://developer.apple.com/documentation/Metal/MTLFunctionStitchingInputNode
type IMTLFunctionStitchingInputNode interface {
	objectivec.IObject
	MTLFunctionStitchingNode

	// Topic: Initializing an input node

	// Creates a new input node.
	InitWithArgumentIndex(argument uint) MTLFunctionStitchingInputNode

	// Topic: Configuring an input node

	// The index in the command’s buffer argument table that declares which data to read for this input node.
	ArgumentIndex() uint
	SetArgumentIndex(value uint)
}

// Init initializes the instance.
func (f MTLFunctionStitchingInputNode) Init() MTLFunctionStitchingInputNode {
	rv := objc.Send[MTLFunctionStitchingInputNode](f.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f MTLFunctionStitchingInputNode) Autorelease() MTLFunctionStitchingInputNode {
	rv := objc.Send[MTLFunctionStitchingInputNode](f.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTLFunctionStitchingInputNode creates a new MTLFunctionStitchingInputNode instance.
func NewMTLFunctionStitchingInputNode() MTLFunctionStitchingInputNode {
	class := getMTLFunctionStitchingInputNodeClass()
	rv := objc.Send[MTLFunctionStitchingInputNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a new input node.
//
// argument: The index of the parameter in the stitched function’s parameter list. The
// first parameter is `0`, the second is `1`, and so on.
//
// See: https://developer.apple.com/documentation/Metal/MTLFunctionStitchingInputNode/init(argumentIndex:)
func NewFunctionStitchingInputNodeWithArgumentIndex(argument uint) MTLFunctionStitchingInputNode {
	instance := getMTLFunctionStitchingInputNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithArgumentIndex:"), argument)
	return MTLFunctionStitchingInputNodeFromID(rv)
}

// Creates a new input node.
//
// argument: The index of the parameter in the stitched function’s parameter list. The
// first parameter is `0`, the second is `1`, and so on.
//
// See: https://developer.apple.com/documentation/Metal/MTLFunctionStitchingInputNode/init(argumentIndex:)
func (f MTLFunctionStitchingInputNode) InitWithArgumentIndex(argument uint) MTLFunctionStitchingInputNode {
	rv := objc.Send[MTLFunctionStitchingInputNode](f.ID, objc.Sel("initWithArgumentIndex:"), argument)
	return rv
}

// The index in the command’s buffer argument table that declares which data
// to read for this input node.
//
// See: https://developer.apple.com/documentation/Metal/MTLFunctionStitchingInputNode/argumentIndex
func (f MTLFunctionStitchingInputNode) ArgumentIndex() uint {
	rv := objc.Send[uint](f.ID, objc.Sel("argumentIndex"))
	return rv
}
func (f MTLFunctionStitchingInputNode) SetArgumentIndex(value uint) {
	objc.Send[struct{}](f.ID, objc.Sel("setArgumentIndex:"), value)
}

// Protocol methods for MTLFunctionStitchingNode
