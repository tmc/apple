// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MTLStitchedLibraryDescriptor] class.
var (
	_MTLStitchedLibraryDescriptorClass     MTLStitchedLibraryDescriptorClass
	_MTLStitchedLibraryDescriptorClassOnce sync.Once
)

func getMTLStitchedLibraryDescriptorClass() MTLStitchedLibraryDescriptorClass {
	_MTLStitchedLibraryDescriptorClassOnce.Do(func() {
		_MTLStitchedLibraryDescriptorClass = MTLStitchedLibraryDescriptorClass{class: objc.GetClass("MTLStitchedLibraryDescriptor")}
	})
	return _MTLStitchedLibraryDescriptorClass
}

// GetMTLStitchedLibraryDescriptorClass returns the class object for MTLStitchedLibraryDescriptor.
func GetMTLStitchedLibraryDescriptorClass() MTLStitchedLibraryDescriptorClass {
	return getMTLStitchedLibraryDescriptorClass()
}

type MTLStitchedLibraryDescriptorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MTLStitchedLibraryDescriptorClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MTLStitchedLibraryDescriptorClass) Alloc() MTLStitchedLibraryDescriptor {
	rv := objc.Send[MTLStitchedLibraryDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A description of a new library of procedurally generated functions.
//
// # Overview
// 
// An [MTLStitchedLibraryDescriptor] describes a library of new stitched
// functions. A is a visible function you create by composing other Metal
// shader functions together in a function graph.
// 
// Configure a stitched library descriptor by assigning an array of one or
// more [MTLFunctionStitchingGraph] instances, each describing a stitched
// function, to the [MTLStitchedLibraryDescriptor.FunctionGraphs] property. Then assign an [MTLFunction]
// array that includes all the functions the graphs depend on to the
// [MTLStitchedLibraryDescriptor.Functions] property.
// 
// Create a stitched library from the descriptor by passing it to the
// [NewLibraryWithStitchedDescriptorError] method of an [MTLDevice]. You can
// change the descriptor to create other libraries without affecting any
// existing ones.
//
// # Configuring a stitched library
//
//   - [MTLStitchedLibraryDescriptor.Functions]: The list of functions for creating the stitched library.
//   - [MTLStitchedLibraryDescriptor.SetFunctions]
//   - [MTLStitchedLibraryDescriptor.FunctionGraphs]: The function graphs that define the new stitched library’s functions.
//   - [MTLStitchedLibraryDescriptor.SetFunctionGraphs]
//
// # Instance Properties
//
//   - [MTLStitchedLibraryDescriptor.BinaryArchives]
//   - [MTLStitchedLibraryDescriptor.SetBinaryArchives]
//   - [MTLStitchedLibraryDescriptor.Options]
//   - [MTLStitchedLibraryDescriptor.SetOptions]
//
// See: https://developer.apple.com/documentation/Metal/MTLStitchedLibraryDescriptor
type MTLStitchedLibraryDescriptor struct {
	objectivec.Object
}

// MTLStitchedLibraryDescriptorFromID constructs a [MTLStitchedLibraryDescriptor] from an objc.ID.
//
// A description of a new library of procedurally generated functions.
func MTLStitchedLibraryDescriptorFromID(id objc.ID) MTLStitchedLibraryDescriptor {
	return MTLStitchedLibraryDescriptor{objectivec.Object{ID: id}}
}
// NOTE: MTLStitchedLibraryDescriptor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MTLStitchedLibraryDescriptor] class.
//
// # Configuring a stitched library
//
//   - [IMTLStitchedLibraryDescriptor.Functions]: The list of functions for creating the stitched library.
//   - [IMTLStitchedLibraryDescriptor.SetFunctions]
//   - [IMTLStitchedLibraryDescriptor.FunctionGraphs]: The function graphs that define the new stitched library’s functions.
//   - [IMTLStitchedLibraryDescriptor.SetFunctionGraphs]
//
// # Instance Properties
//
//   - [IMTLStitchedLibraryDescriptor.BinaryArchives]
//   - [IMTLStitchedLibraryDescriptor.SetBinaryArchives]
//   - [IMTLStitchedLibraryDescriptor.Options]
//   - [IMTLStitchedLibraryDescriptor.SetOptions]
//
// See: https://developer.apple.com/documentation/Metal/MTLStitchedLibraryDescriptor
type IMTLStitchedLibraryDescriptor interface {
	objectivec.IObject

	// Topic: Configuring a stitched library

	// The list of functions for creating the stitched library.
	Functions() []objectivec.IObject
	SetFunctions(value []objectivec.IObject)
	// The function graphs that define the new stitched library’s functions.
	FunctionGraphs() []MTLFunctionStitchingGraph
	SetFunctionGraphs(value []MTLFunctionStitchingGraph)

	// Topic: Instance Properties

	BinaryArchives() []objectivec.IObject
	SetBinaryArchives(value []objectivec.IObject)
	Options() MTLStitchedLibraryOptions
	SetOptions(value MTLStitchedLibraryOptions)
}

// Init initializes the instance.
func (s MTLStitchedLibraryDescriptor) Init() MTLStitchedLibraryDescriptor {
	rv := objc.Send[MTLStitchedLibraryDescriptor](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s MTLStitchedLibraryDescriptor) Autorelease() MTLStitchedLibraryDescriptor {
	rv := objc.Send[MTLStitchedLibraryDescriptor](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTLStitchedLibraryDescriptor creates a new MTLStitchedLibraryDescriptor instance.
func NewMTLStitchedLibraryDescriptor() MTLStitchedLibraryDescriptor {
	class := getMTLStitchedLibraryDescriptorClass()
	rv := objc.Send[MTLStitchedLibraryDescriptor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The list of functions for creating the stitched library.
//
// # Discussion
// 
// The function objects need to all be created by the same Metal device object
// that you’ll use to create the library. The MSL functions referenced by
// these function objects need to be declared with the `stitchable` attribute,
// as in the example below:
//
// See: https://developer.apple.com/documentation/Metal/MTLStitchedLibraryDescriptor/functions
func (s MTLStitchedLibraryDescriptor) Functions() []objectivec.IObject {
	rv := objc.Send[[]objc.ID](s.ID, objc.Sel("functions"))
	return objc.ConvertSlice(rv, func(id objc.ID) objectivec.IObject {
		return objectivec.Object{ID: id}
	})
}
func (s MTLStitchedLibraryDescriptor) SetFunctions(value []objectivec.IObject) {
	objc.Send[struct{}](s.ID, objc.Sel("setFunctions:"), objectivec.IObjectSliceToNSArray(value))
}
// The function graphs that define the new stitched library’s functions.
//
// See: https://developer.apple.com/documentation/Metal/MTLStitchedLibraryDescriptor/functionGraphs
func (s MTLStitchedLibraryDescriptor) FunctionGraphs() []MTLFunctionStitchingGraph {
	rv := objc.Send[[]objc.ID](s.ID, objc.Sel("functionGraphs"))
	return objc.ConvertSlice(rv, func(id objc.ID) MTLFunctionStitchingGraph {
		return MTLFunctionStitchingGraphFromID(id)
	})
}
func (s MTLStitchedLibraryDescriptor) SetFunctionGraphs(value []MTLFunctionStitchingGraph) {
	objc.Send[struct{}](s.ID, objc.Sel("setFunctionGraphs:"), objectivec.IObjectSliceToNSArray(value))
}
// See: https://developer.apple.com/documentation/Metal/MTLStitchedLibraryDescriptor/binaryArchives
func (s MTLStitchedLibraryDescriptor) BinaryArchives() []objectivec.IObject {
	rv := objc.Send[[]objc.ID](s.ID, objc.Sel("binaryArchives"))
	return objc.ConvertSlice(rv, func(id objc.ID) objectivec.IObject {
		return objectivec.Object{ID: id}
	})
}
func (s MTLStitchedLibraryDescriptor) SetBinaryArchives(value []objectivec.IObject) {
	objc.Send[struct{}](s.ID, objc.Sel("setBinaryArchives:"), objectivec.IObjectSliceToNSArray(value))
}
// See: https://developer.apple.com/documentation/Metal/MTLStitchedLibraryDescriptor/options
func (s MTLStitchedLibraryDescriptor) Options() MTLStitchedLibraryOptions {
	rv := objc.Send[MTLStitchedLibraryOptions](s.ID, objc.Sel("options"))
	return MTLStitchedLibraryOptions(rv)
}
func (s MTLStitchedLibraryDescriptor) SetOptions(value MTLStitchedLibraryOptions) {
	objc.Send[struct{}](s.ID, objc.Sel("setOptions:"), value)
}

