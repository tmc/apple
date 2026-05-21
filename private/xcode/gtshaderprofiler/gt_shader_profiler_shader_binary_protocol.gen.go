// Code generated from Apple documentation for gtshaderprofiler. DO NOT EDIT.

package gtshaderprofiler

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// GTShaderProfilerShaderBinary protocol.
type GTShaderProfilerShaderBinary interface {
	objectivec.IObject

	// AddrEnd protocol.
	AddrEnd() uint32

	// AddrStart protocol.
	AddrStart() uint32

	// CostForAddress protocol.
	CostForAddress(address uint32) float64

	// CostForDrawAtIndex protocol.
	CostForDrawAtIndex(index uint32) float64

	// CostPercentageForDrawAtIndex protocol.
	CostPercentageForDrawAtIndex(index uint32) float64

	// IsDylib protocol.
	IsDylib() bool

	// NumSamples protocol.
	NumSamples() uint64

	// TotalCost protocol.
	TotalCost() float64

	// Type protocol.
	Type() uint32
}

// GTShaderProfilerShaderBinaryObject wraps an existing Objective-C object that conforms to the GTShaderProfilerShaderBinary protocol.
type GTShaderProfilerShaderBinaryObject struct {
	objectivec.Object
}

func (o GTShaderProfilerShaderBinaryObject) BaseObject() objectivec.Object {
	return o.Object
}

// GTShaderProfilerShaderBinaryObjectFromID constructs a [GTShaderProfilerShaderBinaryObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func GTShaderProfilerShaderBinaryObjectFromID(id objc.ID) GTShaderProfilerShaderBinaryObject {
	return GTShaderProfilerShaderBinaryObject{
		Object: objectivec.ObjectFromID(id),
	}
}

func (o GTShaderProfilerShaderBinaryObject) AddrEnd() uint32 {
	rv := objc.Send[uint32](o.ID, objc.Sel("addrEnd"))
	return rv
}
func (o GTShaderProfilerShaderBinaryObject) AddrStart() uint32 {
	rv := objc.Send[uint32](o.ID, objc.Sel("addrStart"))
	return rv
}
func (o GTShaderProfilerShaderBinaryObject) BinaryRanges() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("binaryRanges"))
	return objectivec.Object{ID: rv}
}
func (o GTShaderProfilerShaderBinaryObject) CostForAddress(address uint32) float64 {
	rv := objc.Send[float64](o.ID, objc.Sel("costForAddress:"), address)
	return rv
}
func (o GTShaderProfilerShaderBinaryObject) CostForDrawAtIndex(index uint32) float64 {
	rv := objc.Send[float64](o.ID, objc.Sel("costForDrawAtIndex:"), index)
	return rv
}
func (o GTShaderProfilerShaderBinaryObject) CostPercentageForDrawAtIndex(index uint32) float64 {
	rv := objc.Send[float64](o.ID, objc.Sel("costPercentageForDrawAtIndex:"), index)
	return rv
}
func (o GTShaderProfilerShaderBinaryObject) Diassemblies() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("diassemblies"))
	return objectivec.Object{ID: rv}
}
func (o GTShaderProfilerShaderBinaryObject) DiassemblyAtAddress(address uint32) objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("diassemblyAtAddress:"), address)
	return objectivec.Object{ID: rv}
}
func (o GTShaderProfilerShaderBinaryObject) FullPath() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("fullPath"))
	return objectivec.Object{ID: rv}
}
func (o GTShaderProfilerShaderBinaryObject) IsDylib() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isDylib"))
	return rv
}
func (o GTShaderProfilerShaderBinaryObject) Key() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("key"))
	return objectivec.Object{ID: rv}
}
func (o GTShaderProfilerShaderBinaryObject) NumSamples() uint64 {
	rv := objc.Send[uint64](o.ID, objc.Sel("numSamples"))
	return rv
}
func (o GTShaderProfilerShaderBinaryObject) StringFromIndex(index uint64) objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("stringFromIndex:"), index)
	return objectivec.Object{ID: rv}
}
func (o GTShaderProfilerShaderBinaryObject) TotalCost() float64 {
	rv := objc.Send[float64](o.ID, objc.Sel("totalCost"))
	return rv
}
func (o GTShaderProfilerShaderBinaryObject) Type() uint32 {
	rv := objc.Send[uint32](o.ID, objc.Sel("type"))
	return rv
}
func (o GTShaderProfilerShaderBinaryObject) TypeName() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("typeName"))
	return objectivec.Object{ID: rv}
}
