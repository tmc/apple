// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// MTLThreadgroupBinding protocol.
//
// See: https://developer.apple.com/documentation/Metal/MTLThreadgroupBinding
type MTLThreadgroupBinding interface {
	objectivec.IObject
	MTLBinding

	// ThreadgroupMemoryAlignment protocol.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLThreadgroupBinding/threadgroupMemoryAlignment
	ThreadgroupMemoryAlignment() uint

	// ThreadgroupMemoryDataSize protocol.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLThreadgroupBinding/threadgroupMemoryDataSize
	ThreadgroupMemoryDataSize() uint
}

// MTLThreadgroupBindingObject wraps an existing Objective-C object that conforms to the MTLThreadgroupBinding protocol.
type MTLThreadgroupBindingObject struct {
	objectivec.Object
}
func (o MTLThreadgroupBindingObject) BaseObject() objectivec.Object {
	return o.Object
}

// MTLThreadgroupBindingObjectFromID constructs a [MTLThreadgroupBindingObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func MTLThreadgroupBindingObjectFromID(id objc.ID) MTLThreadgroupBindingObject {
	return MTLThreadgroupBindingObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// See: https://developer.apple.com/documentation/Metal/MTLThreadgroupBinding/threadgroupMemoryAlignment
func (o MTLThreadgroupBindingObject) ThreadgroupMemoryAlignment() uint {
	rv := objc.Send[uint](o.ID, objc.Sel("threadgroupMemoryAlignment"))
	return rv
	}
// See: https://developer.apple.com/documentation/Metal/MTLThreadgroupBinding/threadgroupMemoryDataSize
func (o MTLThreadgroupBindingObject) ThreadgroupMemoryDataSize() uint {
	rv := objc.Send[uint](o.ID, objc.Sel("threadgroupMemoryDataSize"))
	return rv
	}
// See: https://developer.apple.com/documentation/Metal/MTLBinding/access
func (o MTLThreadgroupBindingObject) Access() MTLBindingAccess {
	rv := objc.Send[MTLBindingAccess](o.ID, objc.Sel("access"))
	return rv
	}
// See: https://developer.apple.com/documentation/Metal/MTLBinding/index
func (o MTLThreadgroupBindingObject) Index() uint {
	rv := objc.Send[uint](o.ID, objc.Sel("index"))
	return rv
	}
// See: https://developer.apple.com/documentation/Metal/MTLBinding/isArgument
func (o MTLThreadgroupBindingObject) IsArgument() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isArgument"))
	return rv
	}
// See: https://developer.apple.com/documentation/Metal/MTLBinding/isUsed
func (o MTLThreadgroupBindingObject) IsUsed() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isUsed"))
	return rv
	}
// See: https://developer.apple.com/documentation/Metal/MTLBinding/name
func (o MTLThreadgroupBindingObject) Name() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("name"))
	return foundation.NSStringFromID(rv).String()
	}
// See: https://developer.apple.com/documentation/Metal/MTLBinding/type
func (o MTLThreadgroupBindingObject) Type() MTLBindingType {
	rv := objc.Send[MTLBindingType](o.ID, objc.Sel("type"))
	return rv
	}

