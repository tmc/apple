// Code generated from Apple documentation for diskimages2. DO NOT EDIT.

package diskimages2

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// DIController2ClientProtocol protocol.
//
// See: https://developer.apple.com/documentation/DiskImages2/DIController2ClientProtocol
type DIController2ClientProtocol interface {
	objectivec.IObject
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

//
// See: https://developer.apple.com/documentation/DiskImages2/DIController2ClientProtocol/attachCompletedWithHandle:reply:
func (o DIController2ClientProtocolObject) AttachCompletedWithHandleReply(handle objectivec.IObject, reply ErrorHandler) {
	objc.Send[struct{}](o.ID, objc.Sel("attachCompletedWithHandle:reply:"), handle, reply)
	}

