// Code generated from Apple documentation for gtshaderprofiler. DO NOT EDIT.

package gtshaderprofiler

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// GTShaderProfilerShaderDiassembly protocol.
type GTShaderProfilerShaderDiassembly interface {
	objectivec.IObject

	// Address protocol.
	Address() uint32

	// Cost protocol.
	Cost() float64

	// Opcode protocol.
	Opcode() uint32

	// OpcodeMask protocol.
	OpcodeMask() uint32

	// OpcodeType protocol.
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

func (o GTShaderProfilerShaderDiassemblyObject) Address() uint32 {
	rv := objc.Send[uint32](o.ID, objc.Sel("address"))
	return rv
}
func (o GTShaderProfilerShaderDiassemblyObject) Binary() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("binary"))
	return objectivec.Object{ID: rv}
}
func (o GTShaderProfilerShaderDiassemblyObject) Cost() float64 {
	rv := objc.Send[float64](o.ID, objc.Sel("cost"))
	return rv
}
func (o GTShaderProfilerShaderDiassemblyObject) Diassembly() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("diassembly"))
	return objectivec.Object{ID: rv}
}
func (o GTShaderProfilerShaderDiassemblyObject) Opcode() uint32 {
	rv := objc.Send[uint32](o.ID, objc.Sel("opcode"))
	return rv
}
func (o GTShaderProfilerShaderDiassemblyObject) OpcodeMask() uint32 {
	rv := objc.Send[uint32](o.ID, objc.Sel("opcodeMask"))
	return rv
}
func (o GTShaderProfilerShaderDiassemblyObject) OpcodeType() uint32 {
	rv := objc.Send[uint32](o.ID, objc.Sel("opcodeType"))
	return rv
}
