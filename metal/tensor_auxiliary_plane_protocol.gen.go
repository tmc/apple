// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// A type that represents the configuration and storage of an auxiliary plane in a multi-plane tensor.
//
// See: https://developer.apple.com/documentation/Metal/MTLTensorAuxiliaryPlane
type MTLTensorAuxiliaryPlane interface {
	objectivec.IObject

	// The number of data plane elements that correspond to one element in this auxiliary plane.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLTensorAuxiliaryPlane/blockFactors
	BlockFactors() IMTLTensorExtents

	// The buffer that provides the underlying storage for this plane, or `nil` if no buffer was provided at initialization.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLTensorAuxiliaryPlane/buffer
	Buffer() MTLBuffer

	// The byte offset into [buffer](<https://developer.apple.com/documentation/Metal/MTLTensorAuxiliaryPlane/buffer>) where this plane’s data begins, or `0` if no buffer was provided at initialization.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLTensorAuxiliaryPlane/bufferOffset
	BufferOffset() uint

	// The data format of all elements in the plane.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLTensorAuxiliaryPlane/dataType
	DataType() MTLTensorDataType

	// The type of information this plane stores.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLTensorAuxiliaryPlane/planeType
	PlaneType() MTLTensorPlaneType
}

// MTLTensorAuxiliaryPlaneObject wraps an existing Objective-C object that conforms to the MTLTensorAuxiliaryPlane protocol.
type MTLTensorAuxiliaryPlaneObject struct {
	objectivec.Object
}

func (o MTLTensorAuxiliaryPlaneObject) BaseObject() objectivec.Object {
	return o.Object
}

// MTLTensorAuxiliaryPlaneObjectFromID constructs a [MTLTensorAuxiliaryPlaneObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func MTLTensorAuxiliaryPlaneObjectFromID(id objc.ID) MTLTensorAuxiliaryPlaneObject {
	return MTLTensorAuxiliaryPlaneObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The number of data plane elements that correspond to one element in this
// auxiliary plane.
//
// See: https://developer.apple.com/documentation/Metal/MTLTensorAuxiliaryPlane/blockFactors
func (o MTLTensorAuxiliaryPlaneObject) BlockFactors() IMTLTensorExtents {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("blockFactors"))
	return MTLTensorExtentsFromID(rv)
}

// The buffer that provides the underlying storage for this plane, or `nil` if
// no buffer was provided at initialization.
//
// See: https://developer.apple.com/documentation/Metal/MTLTensorAuxiliaryPlane/buffer
func (o MTLTensorAuxiliaryPlaneObject) Buffer() MTLBuffer {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("buffer"))
	return MTLBufferObjectFromID(rv)
}

// The byte offset into [Buffer] where this plane’s data begins, or `0` if
// no buffer was provided at initialization.
//
// See: https://developer.apple.com/documentation/Metal/MTLTensorAuxiliaryPlane/bufferOffset
func (o MTLTensorAuxiliaryPlaneObject) BufferOffset() uint {
	rv := objc.Send[uint](o.ID, objc.Sel("bufferOffset"))
	return uint(rv)
}

// The data format of all elements in the plane.
//
// See: https://developer.apple.com/documentation/Metal/MTLTensorAuxiliaryPlane/dataType
func (o MTLTensorAuxiliaryPlaneObject) DataType() MTLTensorDataType {
	rv := objc.Send[MTLTensorDataType](o.ID, objc.Sel("dataType"))
	return MTLTensorDataType(rv)
}

// The type of information this plane stores.
//
// See: https://developer.apple.com/documentation/Metal/MTLTensorAuxiliaryPlane/planeType
func (o MTLTensorAuxiliaryPlaneObject) PlaneType() MTLTensorPlaneType {
	rv := objc.Send[MTLTensorPlaneType](o.ID, objc.Sel("planeType"))
	return MTLTensorPlaneType(rv)
}
