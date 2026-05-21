// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// SLSWindowManagementBridgeDelegate protocol.
type SLSWindowManagementBridgeDelegate interface {
	objectivec.IObject

	// PerformWindowManagementBridgeTransactionUsingBlock protocol.
	PerformWindowManagementBridgeTransactionUsingBlock(block VoidHandler)
}

// SLSWindowManagementBridgeDelegateObject wraps an existing Objective-C object that conforms to the SLSWindowManagementBridgeDelegate protocol.
type SLSWindowManagementBridgeDelegateObject struct {
	objectivec.Object
}

func (o SLSWindowManagementBridgeDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}

// SLSWindowManagementBridgeDelegateObjectFromID constructs a [SLSWindowManagementBridgeDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func SLSWindowManagementBridgeDelegateObjectFromID(id objc.ID) SLSWindowManagementBridgeDelegateObject {
	return SLSWindowManagementBridgeDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
}

func (o SLSWindowManagementBridgeDelegateObject) PerformAsynchronousBridgedWindowManagementOperation(operation objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("performAsynchronousBridgedWindowManagementOperation:"), operation)
}
func (o SLSWindowManagementBridgeDelegateObject) PerformSynchronousBridgedWindowManagementOperation(operation objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("performSynchronousBridgedWindowManagementOperation:"), operation)
	return objectivec.Object{ID: rv}
}
func (o SLSWindowManagementBridgeDelegateObject) PerformWindowManagementBridgeTransactionUsingBlock(block VoidHandler) {
	objc.Send[struct{}](o.ID, objc.Sel("performWindowManagementBridgeTransactionUsingBlock:"), block)
}
func (o SLSWindowManagementBridgeDelegateObject) SetWindowTagsOnWindowClear(tags objectivec.IObject, window uint32, clear bool) {
	objc.Send[struct{}](o.ID, objc.Sel("setWindowTags:onWindow:clear:"), tags, window, clear)
}
