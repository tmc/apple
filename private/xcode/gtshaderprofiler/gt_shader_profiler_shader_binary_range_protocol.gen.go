// Code generated from Apple documentation for gtshaderprofiler. DO NOT EDIT.

package gtshaderprofiler

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// GTShaderProfilerShaderBinaryRange protocol.
type GTShaderProfilerShaderBinaryRange interface {
	objectivec.IObject

	// AddrEnd protocol.
	AddrEnd() uint32

	// AddrStart protocol.
	AddrStart() uint32

	// CostForAddress protocol.
	CostForAddress(address uint32) float64

	// NumSamples protocol.
	NumSamples() uint64

	// TotalCost protocol.
	TotalCost() float64
}

// GTShaderProfilerShaderBinaryRangeObject wraps an existing Objective-C object that conforms to the GTShaderProfilerShaderBinaryRange protocol.
type GTShaderProfilerShaderBinaryRangeObject struct {
	objectivec.Object
}

func (o GTShaderProfilerShaderBinaryRangeObject) BaseObject() objectivec.Object {
	return o.Object
}

// GTShaderProfilerShaderBinaryRangeObjectFromID constructs a [GTShaderProfilerShaderBinaryRangeObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func GTShaderProfilerShaderBinaryRangeObjectFromID(id objc.ID) GTShaderProfilerShaderBinaryRangeObject {
	return GTShaderProfilerShaderBinaryRangeObject{
		Object: objectivec.ObjectFromID(id),
	}
}

func (o GTShaderProfilerShaderBinaryRangeObject) AddrEnd() uint32 {
	rv := objc.Send[uint32](o.ID, objc.Sel("addrEnd"))
	return rv
}
func (o GTShaderProfilerShaderBinaryRangeObject) AddrStart() uint32 {
	rv := objc.Send[uint32](o.ID, objc.Sel("addrStart"))
	return rv
}
func (o GTShaderProfilerShaderBinaryRangeObject) Binary() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("binary"))
	return objectivec.Object{ID: rv}
}
func (o GTShaderProfilerShaderBinaryRangeObject) CallStack() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("callStack"))
	return objectivec.Object{ID: rv}
}
func (o GTShaderProfilerShaderBinaryRangeObject) CostForAddress(address uint32) float64 {
	rv := objc.Send[float64](o.ID, objc.Sel("costForAddress:"), address)
	return rv
}
func (o GTShaderProfilerShaderBinaryRangeObject) Location() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("location"))
	return objectivec.Object{ID: rv}
}
func (o GTShaderProfilerShaderBinaryRangeObject) NumSamples() uint64 {
	rv := objc.Send[uint64](o.ID, objc.Sel("numSamples"))
	return rv
}
func (o GTShaderProfilerShaderBinaryRangeObject) TotalCost() float64 {
	rv := objc.Send[float64](o.ID, objc.Sel("totalCost"))
	return rv
}
