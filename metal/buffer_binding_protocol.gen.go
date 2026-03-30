// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// MTLBufferBinding protocol.
//
// See: https://developer.apple.com/documentation/Metal/MTLBufferBinding
type MTLBufferBinding interface {
	objectivec.IObject
	MTLBinding

	// BufferAlignment protocol.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLBufferBinding/bufferAlignment
	BufferAlignment() uint

	// BufferDataSize protocol.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLBufferBinding/bufferDataSize
	BufferDataSize() uint

	// BufferDataType protocol.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLBufferBinding/bufferDataType
	BufferDataType() MTLDataType

	// BufferPointerType protocol.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLBufferBinding/bufferPointerType
	BufferPointerType() IMTLPointerType

	// BufferStructType protocol.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLBufferBinding/bufferStructType
	BufferStructType() IMTLStructType
}

// MTLBufferBindingObject wraps an existing Objective-C object that conforms to the MTLBufferBinding protocol.
type MTLBufferBindingObject struct {
	objectivec.Object
}

func (o MTLBufferBindingObject) BaseObject() objectivec.Object {
	return o.Object
}

// MTLBufferBindingObjectFromID constructs a [MTLBufferBindingObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func MTLBufferBindingObjectFromID(id objc.ID) MTLBufferBindingObject {
	return MTLBufferBindingObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// See: https://developer.apple.com/documentation/Metal/MTLBufferBinding/bufferAlignment
func (o MTLBufferBindingObject) BufferAlignment() uint {
	rv := objc.Send[uint](o.ID, objc.Sel("bufferAlignment"))
	return rv
}

// See: https://developer.apple.com/documentation/Metal/MTLBufferBinding/bufferDataSize
func (o MTLBufferBindingObject) BufferDataSize() uint {
	rv := objc.Send[uint](o.ID, objc.Sel("bufferDataSize"))
	return rv
}

// See: https://developer.apple.com/documentation/Metal/MTLBufferBinding/bufferDataType
func (o MTLBufferBindingObject) BufferDataType() MTLDataType {
	rv := objc.Send[MTLDataType](o.ID, objc.Sel("bufferDataType"))
	return rv
}

// See: https://developer.apple.com/documentation/Metal/MTLBufferBinding/bufferPointerType
func (o MTLBufferBindingObject) BufferPointerType() IMTLPointerType {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("bufferPointerType"))
	return MTLPointerTypeFromID(rv)
}

// See: https://developer.apple.com/documentation/Metal/MTLBufferBinding/bufferStructType
func (o MTLBufferBindingObject) BufferStructType() IMTLStructType {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("bufferStructType"))
	return MTLStructTypeFromID(rv)
}

// See: https://developer.apple.com/documentation/Metal/MTLBinding/access
func (o MTLBufferBindingObject) Access() MTLBindingAccess {
	rv := objc.Send[MTLBindingAccess](o.ID, objc.Sel("access"))
	return rv
}

// See: https://developer.apple.com/documentation/Metal/MTLBinding/index
func (o MTLBufferBindingObject) Index() uint {
	rv := objc.Send[uint](o.ID, objc.Sel("index"))
	return rv
}

// See: https://developer.apple.com/documentation/Metal/MTLBinding/isArgument
func (o MTLBufferBindingObject) IsArgument() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isArgument"))
	return rv
}

// See: https://developer.apple.com/documentation/Metal/MTLBinding/isUsed
func (o MTLBufferBindingObject) IsUsed() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isUsed"))
	return rv
}

// See: https://developer.apple.com/documentation/Metal/MTLBinding/name
func (o MTLBufferBindingObject) Name() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("name"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/Metal/MTLBinding/type
func (o MTLBufferBindingObject) Type() MTLBindingType {
	rv := objc.Send[MTLBindingType](o.ID, objc.Sel("type"))
	return rv
}
