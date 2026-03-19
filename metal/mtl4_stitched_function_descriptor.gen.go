// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MTL4StitchedFunctionDescriptor] class.
var (
	_MTL4StitchedFunctionDescriptorClass     MTL4StitchedFunctionDescriptorClass
	_MTL4StitchedFunctionDescriptorClassOnce sync.Once
)

func getMTL4StitchedFunctionDescriptorClass() MTL4StitchedFunctionDescriptorClass {
	_MTL4StitchedFunctionDescriptorClassOnce.Do(func() {
		_MTL4StitchedFunctionDescriptorClass = MTL4StitchedFunctionDescriptorClass{class: objc.GetClass("MTL4StitchedFunctionDescriptor")}
	})
	return _MTL4StitchedFunctionDescriptorClass
}

// GetMTL4StitchedFunctionDescriptorClass returns the class object for MTL4StitchedFunctionDescriptor.
func GetMTL4StitchedFunctionDescriptorClass() MTL4StitchedFunctionDescriptorClass {
	return getMTL4StitchedFunctionDescriptorClass()
}

type MTL4StitchedFunctionDescriptorClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MTL4StitchedFunctionDescriptorClass) Alloc() MTL4StitchedFunctionDescriptor {
	rv := objc.Send[MTL4StitchedFunctionDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// Groups together properties that describe a shader function suitable for
// stitching.
//
// # Instance Properties
//
//   - [MTL4StitchedFunctionDescriptor.FunctionDescriptors]: Configures an array of function descriptors with references to functions that contribute to the stitching process.
//   - [MTL4StitchedFunctionDescriptor.SetFunctionDescriptors]
//   - [MTL4StitchedFunctionDescriptor.FunctionGraph]: Sets the graph representing how to stitch functions together.
//   - [MTL4StitchedFunctionDescriptor.SetFunctionGraph]
//
// See: https://developer.apple.com/documentation/Metal/MTL4StitchedFunctionDescriptor
type MTL4StitchedFunctionDescriptor struct {
	MTL4FunctionDescriptor
}

// MTL4StitchedFunctionDescriptorFromID constructs a [MTL4StitchedFunctionDescriptor] from an objc.ID.
//
// Groups together properties that describe a shader function suitable for
// stitching.
func MTL4StitchedFunctionDescriptorFromID(id objc.ID) MTL4StitchedFunctionDescriptor {
	return MTL4StitchedFunctionDescriptor{MTL4FunctionDescriptor: MTL4FunctionDescriptorFromID(id)}
}
// NOTE: MTL4StitchedFunctionDescriptor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MTL4StitchedFunctionDescriptor] class.
//
// # Instance Properties
//
//   - [IMTL4StitchedFunctionDescriptor.FunctionDescriptors]: Configures an array of function descriptors with references to functions that contribute to the stitching process.
//   - [IMTL4StitchedFunctionDescriptor.SetFunctionDescriptors]
//   - [IMTL4StitchedFunctionDescriptor.FunctionGraph]: Sets the graph representing how to stitch functions together.
//   - [IMTL4StitchedFunctionDescriptor.SetFunctionGraph]
//
// See: https://developer.apple.com/documentation/Metal/MTL4StitchedFunctionDescriptor
type IMTL4StitchedFunctionDescriptor interface {
	IMTL4FunctionDescriptor

	// Topic: Instance Properties

	// Configures an array of function descriptors with references to functions that contribute to the stitching process.
	FunctionDescriptors() []MTL4FunctionDescriptor
	SetFunctionDescriptors(value []MTL4FunctionDescriptor)
	// Sets the graph representing how to stitch functions together.
	FunctionGraph() IMTLFunctionStitchingGraph
	SetFunctionGraph(value IMTLFunctionStitchingGraph)
}

// Init initializes the instance.
func (m MTL4StitchedFunctionDescriptor) Init() MTL4StitchedFunctionDescriptor {
	rv := objc.Send[MTL4StitchedFunctionDescriptor](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MTL4StitchedFunctionDescriptor) Autorelease() MTL4StitchedFunctionDescriptor {
	rv := objc.Send[MTL4StitchedFunctionDescriptor](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTL4StitchedFunctionDescriptor creates a new MTL4StitchedFunctionDescriptor instance.
func NewMTL4StitchedFunctionDescriptor() MTL4StitchedFunctionDescriptor {
	class := getMTL4StitchedFunctionDescriptorClass()
	rv := objc.Send[MTL4StitchedFunctionDescriptor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Configures an array of function descriptors with references to functions
// that contribute to the stitching process.
//
// See: https://developer.apple.com/documentation/Metal/MTL4StitchedFunctionDescriptor/functionDescriptors
func (m MTL4StitchedFunctionDescriptor) FunctionDescriptors() []MTL4FunctionDescriptor {
	rv := objc.Send[[]objc.ID](m.ID, objc.Sel("functionDescriptors"))
	return objc.ConvertSlice(rv, func(id objc.ID) MTL4FunctionDescriptor {
		return MTL4FunctionDescriptorFromID(id)
	})
}
func (m MTL4StitchedFunctionDescriptor) SetFunctionDescriptors(value []MTL4FunctionDescriptor) {
	objc.Send[struct{}](m.ID, objc.Sel("setFunctionDescriptors:"), objectivec.IObjectSliceToNSArray(value))
}
// Sets the graph representing how to stitch functions together.
//
// See: https://developer.apple.com/documentation/Metal/MTL4StitchedFunctionDescriptor/functionGraph
func (m MTL4StitchedFunctionDescriptor) FunctionGraph() IMTLFunctionStitchingGraph {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("functionGraph"))
	return MTLFunctionStitchingGraphFromID(objc.ID(rv))
}
func (m MTL4StitchedFunctionDescriptor) SetFunctionGraph(value IMTLFunctionStitchingGraph) {
	objc.Send[struct{}](m.ID, objc.Sel("setFunctionGraph:"), value)
}

