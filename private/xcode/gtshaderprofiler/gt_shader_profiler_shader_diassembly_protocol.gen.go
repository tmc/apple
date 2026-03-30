// Code generated from Apple documentation for gtshaderprofiler. DO NOT EDIT.

package gtshaderprofiler

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// GTShaderProfilerShaderDiassembly protocol.
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerShaderDiassembly
type GTShaderProfilerShaderDiassembly interface {
	objectivec.IObject

	// Address protocol.
	//
	// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerShaderDiassembly/address
	Address() uint32

	// Cost protocol.
	//
	// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerShaderDiassembly/cost
	Cost() float64

	// Opcode protocol.
	//
	// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerShaderDiassembly/opcode
	Opcode() uint32

	// OpcodeMask protocol.
	//
	// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerShaderDiassembly/opcodeMask
	OpcodeMask() uint32

	// OpcodeType protocol.
	//
	// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerShaderDiassembly/opcodeType
	OpcodeType() uint32
}

// GTShaderProfilerShaderDiassemblyObject wraps an existing Objective-C object that conforms to the GTShaderProfilerShaderDiassembly protocol.
type GTShaderProfilerShaderDiassemblyObject struct {
	objectivec.Object
}

func (o GTShaderProfilerShaderDiassemblyObject) BaseObject() objectivec.Object {
	return o.Object
}

// GTShaderProfilerShaderDiassemblyObjectFromID constructs a [GTShaderProfilerShaderDiassemblyObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func GTShaderProfilerShaderDiassemblyObjectFromID(id objc.ID) GTShaderProfilerShaderDiassemblyObject {
	return GTShaderProfilerShaderDiassemblyObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerShaderDiassembly/address
func (o GTShaderProfilerShaderDiassemblyObject) Address() uint32 {
	rv := objc.Send[uint32](o.ID, objc.Sel("address"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerShaderDiassembly/binary
func (o GTShaderProfilerShaderDiassemblyObject) Binary() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("binary"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerShaderDiassembly/cost
func (o GTShaderProfilerShaderDiassemblyObject) Cost() float64 {
	rv := objc.Send[float64](o.ID, objc.Sel("cost"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerShaderDiassembly/diassembly
func (o GTShaderProfilerShaderDiassemblyObject) Diassembly() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("diassembly"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerShaderDiassembly/opcode
func (o GTShaderProfilerShaderDiassemblyObject) Opcode() uint32 {
	rv := objc.Send[uint32](o.ID, objc.Sel("opcode"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerShaderDiassembly/opcodeMask
func (o GTShaderProfilerShaderDiassemblyObject) OpcodeMask() uint32 {
	rv := objc.Send[uint32](o.ID, objc.Sel("opcodeMask"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerShaderDiassembly/opcodeType
func (o GTShaderProfilerShaderDiassemblyObject) OpcodeType() uint32 {
	rv := objc.Send[uint32](o.ID, objc.Sel("opcodeType"))
	return rv
}
