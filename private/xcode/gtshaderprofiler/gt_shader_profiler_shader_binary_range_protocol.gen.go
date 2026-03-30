// Code generated from Apple documentation for gtshaderprofiler. DO NOT EDIT.

package gtshaderprofiler

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// GTShaderProfilerShaderBinaryRange protocol.
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerShaderBinaryRange
type GTShaderProfilerShaderBinaryRange interface {
	objectivec.IObject

	// AddrEnd protocol.
	//
	// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerShaderBinaryRange/addrEnd
	AddrEnd() uint32

	// AddrStart protocol.
	//
	// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerShaderBinaryRange/addrStart
	AddrStart() uint32

	// CostForAddress protocol.
	//
	// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerShaderBinaryRange/costForAddress:
	CostForAddress(address uint32) float64

	// NumSamples protocol.
	//
	// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerShaderBinaryRange/numSamples
	NumSamples() uint64

	// TotalCost protocol.
	//
	// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerShaderBinaryRange/totalCost
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

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerShaderBinaryRange/addrEnd
func (o GTShaderProfilerShaderBinaryRangeObject) AddrEnd() uint32 {
	rv := objc.Send[uint32](o.ID, objc.Sel("addrEnd"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerShaderBinaryRange/addrStart
func (o GTShaderProfilerShaderBinaryRangeObject) AddrStart() uint32 {
	rv := objc.Send[uint32](o.ID, objc.Sel("addrStart"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerShaderBinaryRange/binary
func (o GTShaderProfilerShaderBinaryRangeObject) Binary() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("binary"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerShaderBinaryRange/callStack
func (o GTShaderProfilerShaderBinaryRangeObject) CallStack() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("callStack"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerShaderBinaryRange/costForAddress:
func (o GTShaderProfilerShaderBinaryRangeObject) CostForAddress(address uint32) float64 {
	rv := objc.Send[float64](o.ID, objc.Sel("costForAddress:"), address)
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerShaderBinaryRange/location
func (o GTShaderProfilerShaderBinaryRangeObject) Location() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("location"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerShaderBinaryRange/numSamples
func (o GTShaderProfilerShaderBinaryRangeObject) NumSamples() uint64 {
	rv := objc.Send[uint64](o.ID, objc.Sel("numSamples"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerShaderBinaryRange/totalCost
func (o GTShaderProfilerShaderBinaryRangeObject) TotalCost() float64 {
	rv := objc.Send[float64](o.ID, objc.Sel("totalCost"))
	return rv
}
