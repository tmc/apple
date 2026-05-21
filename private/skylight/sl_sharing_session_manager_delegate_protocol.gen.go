// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// SLSharingSessionManagerDelegate protocol.
type SLSharingSessionManagerDelegate interface {
	objectivec.IObject
}

// SLSharingSessionManagerDelegateObject wraps an existing Objective-C object that conforms to the SLSharingSessionManagerDelegate protocol.
type SLSharingSessionManagerDelegateObject struct {
	objectivec.Object
}

func (o SLSharingSessionManagerDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}

// SLSharingSessionManagerDelegateObjectFromID constructs a [SLSharingSessionManagerDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func SLSharingSessionManagerDelegateObjectFromID(id objc.ID) SLSharingSessionManagerDelegateObject {
	return SLSharingSessionManagerDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
}

func (o SLSharingSessionManagerDelegateObject) SharingSessionManagerPickerCanceledForSession(manager objectivec.IObject, session objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("sharingSessionManager:pickerCanceledForSession:"), manager, session)
}
func (o SLSharingSessionManagerDelegateObject) SharingSessionManagerSessionDidBegin(manager objectivec.IObject, begin objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("sharingSessionManager:sessionDidBegin:"), manager, begin)
}
func (o SLSharingSessionManagerDelegateObject) SharingSessionManagerSessionDidChangeContent(manager objectivec.IObject, content objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("sharingSessionManager:sessionDidChangeContent:"), manager, content)
}
func (o SLSharingSessionManagerDelegateObject) SharingSessionManagerSessionDidChangeMetaData(manager objectivec.IObject, data objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("sharingSessionManager:sessionDidChangeMetaData:"), manager, data)
}
func (o SLSharingSessionManagerDelegateObject) SharingSessionManagerSessionDidEnd(manager objectivec.IObject, end objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("sharingSessionManager:sessionDidEnd:"), manager, end)
}
