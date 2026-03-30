// Code generated from Apple documentation for gtshaderprofiler. DO NOT EDIT.

package gtshaderprofiler

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// GTShaderProfilerShaderBinary protocol.
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerShaderBinary
type GTShaderProfilerShaderBinary interface {
	objectivec.IObject

	// AddrEnd protocol.
	//
	// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerShaderBinary/addrEnd
	AddrEnd() uint32

	// AddrStart protocol.
	//
	// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerShaderBinary/addrStart
	AddrStart() uint32

	// CostForAddress protocol.
	//
	// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerShaderBinary/costForAddress:
	CostForAddress(address uint32) float64

	// CostForDrawAtIndex protocol.
	//
	// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerShaderBinary/costForDrawAtIndex:
	CostForDrawAtIndex(index uint32) float64

	// CostPercentageForDrawAtIndex protocol.
	//
	// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerShaderBinary/costPercentageForDrawAtIndex:
	CostPercentageForDrawAtIndex(index uint32) float64

	// IsDylib protocol.
	//
	// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerShaderBinary/isDylib
	IsDylib() bool

	// NumSamples protocol.
	//
	// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerShaderBinary/numSamples
	NumSamples() uint64

	// TotalCost protocol.
	//
	// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerShaderBinary/totalCost
	TotalCost() float64

	// Type protocol.
	//
	// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerShaderBinary/type
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

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerShaderBinary/addrEnd
func (o GTShaderProfilerShaderBinaryObject) AddrEnd() uint32 {
	rv := objc.Send[uint32](o.ID, objc.Sel("addrEnd"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerShaderBinary/addrStart
func (o GTShaderProfilerShaderBinaryObject) AddrStart() uint32 {
	rv := objc.Send[uint32](o.ID, objc.Sel("addrStart"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerShaderBinary/binaryRanges
func (o GTShaderProfilerShaderBinaryObject) BinaryRanges() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("binaryRanges"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerShaderBinary/costForAddress:
func (o GTShaderProfilerShaderBinaryObject) CostForAddress(address uint32) float64 {
	rv := objc.Send[float64](o.ID, objc.Sel("costForAddress:"), address)
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerShaderBinary/costForDrawAtIndex:
func (o GTShaderProfilerShaderBinaryObject) CostForDrawAtIndex(index uint32) float64 {
	rv := objc.Send[float64](o.ID, objc.Sel("costForDrawAtIndex:"), index)
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerShaderBinary/costPercentageForDrawAtIndex:
func (o GTShaderProfilerShaderBinaryObject) CostPercentageForDrawAtIndex(index uint32) float64 {
	rv := objc.Send[float64](o.ID, objc.Sel("costPercentageForDrawAtIndex:"), index)
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerShaderBinary/diassemblies
func (o GTShaderProfilerShaderBinaryObject) Diassemblies() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("diassemblies"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerShaderBinary/diassemblyAtAddress:
func (o GTShaderProfilerShaderBinaryObject) DiassemblyAtAddress(address uint32) objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("diassemblyAtAddress:"), address)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerShaderBinary/fullPath
func (o GTShaderProfilerShaderBinaryObject) FullPath() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("fullPath"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerShaderBinary/isDylib
func (o GTShaderProfilerShaderBinaryObject) IsDylib() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isDylib"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerShaderBinary/key
func (o GTShaderProfilerShaderBinaryObject) Key() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("key"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerShaderBinary/numSamples
func (o GTShaderProfilerShaderBinaryObject) NumSamples() uint64 {
	rv := objc.Send[uint64](o.ID, objc.Sel("numSamples"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerShaderBinary/stringFromIndex:
func (o GTShaderProfilerShaderBinaryObject) StringFromIndex(index uint64) objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("stringFromIndex:"), index)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerShaderBinary/totalCost
func (o GTShaderProfilerShaderBinaryObject) TotalCost() float64 {
	rv := objc.Send[float64](o.ID, objc.Sel("totalCost"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerShaderBinary/type
func (o GTShaderProfilerShaderBinaryObject) Type() uint32 {
	rv := objc.Send[uint32](o.ID, objc.Sel("type"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerShaderBinary/typeName
func (o GTShaderProfilerShaderBinaryObject) TypeName() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("typeName"))
	return objectivec.Object{ID: rv}
}
