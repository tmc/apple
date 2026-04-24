// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// SLSWindowManagementBridgeDelegate protocol.
//
// See: https://developer.apple.com/documentation/SkyLight/SLSWindowManagementBridgeDelegate
type SLSWindowManagementBridgeDelegate interface {
	objectivec.IObject

	// PerformWindowManagementBridgeTransactionUsingBlock protocol.
	//
	// See: https://developer.apple.com/documentation/SkyLight/SLSWindowManagementBridgeDelegate/performWindowManagementBridgeTransactionUsingBlock:
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

// See: https://developer.apple.com/documentation/SkyLight/SLSWindowManagementBridgeDelegate/performAsynchronousBridgedWindowManagementOperation:
func (o SLSWindowManagementBridgeDelegateObject) PerformAsynchronousBridgedWindowManagementOperation(operation objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("performAsynchronousBridgedWindowManagementOperation:"), operation)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSWindowManagementBridgeDelegate/performSynchronousBridgedWindowManagementOperation:
func (o SLSWindowManagementBridgeDelegateObject) PerformSynchronousBridgedWindowManagementOperation(operation objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("performSynchronousBridgedWindowManagementOperation:"), operation)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/SLSWindowManagementBridgeDelegate/performWindowManagementBridgeTransactionUsingBlock:
func (o SLSWindowManagementBridgeDelegateObject) PerformWindowManagementBridgeTransactionUsingBlock(block VoidHandler) {
	objc.Send[struct{}](o.ID, objc.Sel("performWindowManagementBridgeTransactionUsingBlock:"), block)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSWindowManagementBridgeDelegate/setWindowTags:onWindow:clear:
func (o SLSWindowManagementBridgeDelegateObject) SetWindowTagsOnWindowClear(tags objectivec.IObject, window uint32, clear bool) {
	objc.Send[struct{}](o.ID, objc.Sel("setWindowTags:onWindow:clear:"), tags, window, clear)
}
