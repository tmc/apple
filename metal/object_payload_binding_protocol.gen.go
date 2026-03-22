// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// MTLObjectPayloadBinding protocol.
//
// See: https://developer.apple.com/documentation/Metal/MTLObjectPayloadBinding
type MTLObjectPayloadBinding interface {
	objectivec.IObject
	MTLBinding

	// ObjectPayloadAlignment protocol.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLObjectPayloadBinding/objectPayloadAlignment
	ObjectPayloadAlignment() uint

	// ObjectPayloadDataSize protocol.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLObjectPayloadBinding/objectPayloadDataSize
	ObjectPayloadDataSize() uint
}

// MTLObjectPayloadBindingObject wraps an existing Objective-C object that conforms to the MTLObjectPayloadBinding protocol.
type MTLObjectPayloadBindingObject struct {
	objectivec.Object
}
func (o MTLObjectPayloadBindingObject) BaseObject() objectivec.Object {
	return o.Object
}

// MTLObjectPayloadBindingObjectFromID constructs a [MTLObjectPayloadBindingObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func MTLObjectPayloadBindingObjectFromID(id objc.ID) MTLObjectPayloadBindingObject {
	return MTLObjectPayloadBindingObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// See: https://developer.apple.com/documentation/Metal/MTLObjectPayloadBinding/objectPayloadAlignment
func (o MTLObjectPayloadBindingObject) ObjectPayloadAlignment() uint {
	rv := objc.Send[uint](o.ID, objc.Sel("objectPayloadAlignment"))
	return rv
	}
// See: https://developer.apple.com/documentation/Metal/MTLObjectPayloadBinding/objectPayloadDataSize
func (o MTLObjectPayloadBindingObject) ObjectPayloadDataSize() uint {
	rv := objc.Send[uint](o.ID, objc.Sel("objectPayloadDataSize"))
	return rv
	}
// See: https://developer.apple.com/documentation/Metal/MTLBinding/access
func (o MTLObjectPayloadBindingObject) Access() MTLBindingAccess {
	rv := objc.Send[MTLBindingAccess](o.ID, objc.Sel("access"))
	return rv
	}
// See: https://developer.apple.com/documentation/Metal/MTLBinding/index
func (o MTLObjectPayloadBindingObject) Index() uint {
	rv := objc.Send[uint](o.ID, objc.Sel("index"))
	return rv
	}
// See: https://developer.apple.com/documentation/Metal/MTLBinding/isArgument
func (o MTLObjectPayloadBindingObject) IsArgument() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isArgument"))
	return rv
	}
// See: https://developer.apple.com/documentation/Metal/MTLBinding/isUsed
func (o MTLObjectPayloadBindingObject) IsUsed() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isUsed"))
	return rv
	}
// See: https://developer.apple.com/documentation/Metal/MTLBinding/name
func (o MTLObjectPayloadBindingObject) Name() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("name"))
	return foundation.NSStringFromID(rv).String()
	}
// See: https://developer.apple.com/documentation/Metal/MTLBinding/type
func (o MTLObjectPayloadBindingObject) Type() MTLBindingType {
	rv := objc.Send[MTLBindingType](o.ID, objc.Sel("type"))
	return rv
	}

