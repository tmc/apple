// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// MTLBinding protocol.
//
// See: https://developer.apple.com/documentation/Metal/MTLBinding
type MTLBinding interface {
	objectivec.IObject

	// Access protocol.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLBinding/access
	Access() MTLBindingAccess

	// Index protocol.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLBinding/index
	Index() uint

	// Argument protocol.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLBinding/isArgument
	IsArgument() bool

	// Used protocol.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLBinding/isUsed
	IsUsed() bool

	// Name protocol.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLBinding/name
	Name() string

	// Type protocol.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLBinding/type
	Type() MTLBindingType
}

// MTLBindingObject wraps an existing Objective-C object that conforms to the MTLBinding protocol.
type MTLBindingObject struct {
	objectivec.Object
}
func (o MTLBindingObject) BaseObject() objectivec.Object {
	return o.Object
}

// MTLBindingObjectFromID constructs a [MTLBindingObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func MTLBindingObjectFromID(id objc.ID) MTLBindingObject {
	return MTLBindingObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// See: https://developer.apple.com/documentation/Metal/MTLBinding/access

func (o MTLBindingObject) Access() MTLBindingAccess {
	
	rv := objc.Send[MTLBindingAccess](o.ID, objc.Sel("access"))
	return rv
	}

// See: https://developer.apple.com/documentation/Metal/MTLBinding/index

func (o MTLBindingObject) Index() uint {
	
	rv := objc.Send[uint](o.ID, objc.Sel("index"))
	return rv
	}

// See: https://developer.apple.com/documentation/Metal/MTLBinding/isArgument

func (o MTLBindingObject) IsArgument() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("isArgument"))
	return rv
	}

// See: https://developer.apple.com/documentation/Metal/MTLBinding/isUsed

func (o MTLBindingObject) IsUsed() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("isUsed"))
	return rv
	}

// See: https://developer.apple.com/documentation/Metal/MTLBinding/name

func (o MTLBindingObject) Name() string {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("name"))
	return foundation.NSStringFromID(rv).String()
	}

// See: https://developer.apple.com/documentation/Metal/MTLBinding/type

func (o MTLBindingObject) Type() MTLBindingType {
	
	rv := objc.Send[MTLBindingType](o.ID, objc.Sel("type"))
	return rv
	}

