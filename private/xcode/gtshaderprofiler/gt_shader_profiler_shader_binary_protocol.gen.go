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

	// BinaryRanges protocol.
	BinaryRanges() objectivec.IObject

	// CostForAddress protocol.
	CostForAddress(address uint32) float64

	// CostForDrawAtIndex protocol.
	CostForDrawAtIndex(index uint32) float64

	// CostPercentageForDrawAtIndex protocol.
	CostPercentageForDrawAtIndex(index uint32) float64

	// Diassemblies protocol.
	Diassemblies() objectivec.IObject

	// DiassemblyAtAddress protocol.
	DiassemblyAtAddress(address uint32) objectivec.IObject

	// FullPath protocol.
	FullPath() objectivec.IObject

	// IsDylib protocol.
	IsDylib() bool

	// Key protocol.
	Key() objectivec.IObject

	// NumSamples protocol.
	NumSamples() uint64

	// StringFromIndex protocol.
	StringFromIndex(index uint64) objectivec.IObject

	// TotalCost protocol.
	TotalCost() float64

	// Type protocol.
	Type() uint32

	// TypeName protocol.
	TypeName() objectivec.IObject
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
	rv := objc.SendIfResponds[uint32](o.ID, objc.Sel("addrEnd"))
	return rv
}
func (o GTShaderProfilerShaderBinaryObject) AddrStart() uint32 {
	rv := objc.SendIfResponds[uint32](o.ID, objc.Sel("addrStart"))
	return rv
}
func (o GTShaderProfilerShaderBinaryObject) BinaryRanges() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](o.ID, objc.Sel("binaryRanges"))
	return objectivec.Object{ID: rv}
}
func (o GTShaderProfilerShaderBinaryObject) CostForAddress(address uint32) float64 {
	rv := objc.SendIfResponds[float64](o.ID, objc.Sel("costForAddress:"), address)
	return rv
}
func (o GTShaderProfilerShaderBinaryObject) CostForDrawAtIndex(index uint32) float64 {
	rv := objc.SendIfResponds[float64](o.ID, objc.Sel("costForDrawAtIndex:"), index)
	return rv
}
func (o GTShaderProfilerShaderBinaryObject) CostPercentageForDrawAtIndex(index uint32) float64 {
	rv := objc.SendIfResponds[float64](o.ID, objc.Sel("costPercentageForDrawAtIndex:"), index)
	return rv
}
func (o GTShaderProfilerShaderBinaryObject) Diassemblies() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](o.ID, objc.Sel("diassemblies"))
	return objectivec.Object{ID: rv}
}
func (o GTShaderProfilerShaderBinaryObject) DiassemblyAtAddress(address uint32) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](o.ID, objc.Sel("diassemblyAtAddress:"), address)
	return objectivec.Object{ID: rv}
}
func (o GTShaderProfilerShaderBinaryObject) FullPath() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](o.ID, objc.Sel("fullPath"))
	return objectivec.Object{ID: rv}
}
func (o GTShaderProfilerShaderBinaryObject) IsDylib() bool {
	rv := objc.SendIfResponds[bool](o.ID, objc.Sel("isDylib"))
	return rv
}
func (o GTShaderProfilerShaderBinaryObject) Key() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](o.ID, objc.Sel("key"))
	return objectivec.Object{ID: rv}
}
func (o GTShaderProfilerShaderBinaryObject) NumSamples() uint64 {
	rv := objc.SendIfResponds[uint64](o.ID, objc.Sel("numSamples"))
	return rv
}
func (o GTShaderProfilerShaderBinaryObject) StringFromIndex(index uint64) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](o.ID, objc.Sel("stringFromIndex:"), index)
	return objectivec.Object{ID: rv}
}
func (o GTShaderProfilerShaderBinaryObject) TotalCost() float64 {
	rv := objc.SendIfResponds[float64](o.ID, objc.Sel("totalCost"))
	return rv
}
func (o GTShaderProfilerShaderBinaryObject) Type() uint32 {
	rv := objc.SendIfResponds[uint32](o.ID, objc.Sel("type"))
	return rv
}
func (o GTShaderProfilerShaderBinaryObject) TypeName() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](o.ID, objc.Sel("typeName"))
	return objectivec.Object{ID: rv}
}
