// Code generated from Apple documentation for diskimages2. DO NOT EDIT.

package diskimages2

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// DIController2ClientProtocol protocol.
type DIController2ClientProtocol interface {
	objectivec.IObject

	// AttachCompletedWithHandleReply protocol.
	AttachCompletedWithHandleReply(handle objectivec.IObject, reply ErrorHandler)
}

// DIController2ClientProtocolObject wraps an existing Objective-C object that conforms to the DIController2ClientProtocol protocol.
type DIController2ClientProtocolObject struct {
	objectivec.Object
}

func (o DIController2ClientProtocolObject) BaseObject() objectivec.Object {
	return o.Object
}

// DIController2ClientProtocolObjectFromID constructs a [DIController2ClientProtocolObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func DIController2ClientProtocolObjectFromID(id objc.ID) DIController2ClientProtocolObject {
	return DIController2ClientProtocolObject{
		Object: objectivec.ObjectFromID(id),
	}
}

func (o DIController2ClientProtocolObject) AttachCompletedWithHandleReply(handle objectivec.IObject, reply ErrorHandler) {
	_block1, _cleanup1 := NewErrorBlock(reply)
	defer _cleanup1()
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("attachCompletedWithHandle:reply:"), handle, objc.ID(_block1))
}
