// Code generated from Apple documentation for appleneuralengine. DO NOT EDIT.

package appleneuralengine

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// _ANEStorageMaintainerProtocol protocol.
type ANEStorageMaintainerProtocol interface {
	objectivec.IObject

	// PurgeDanglingModelsAtWithReply protocol.
	PurgeDanglingModelsAtWithReply(at objectivec.IObject, reply VoidHandler)
}

// ANEStorageMaintainerProtocolObject wraps an existing Objective-C object that conforms to the ANEStorageMaintainerProtocol protocol.
type ANEStorageMaintainerProtocolObject struct {
	objectivec.Object
}

func (o ANEStorageMaintainerProtocolObject) BaseObject() objectivec.Object {
	return o.Object
}

// ANEStorageMaintainerProtocolObjectFromID constructs a [ANEStorageMaintainerProtocolObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func ANEStorageMaintainerProtocolObjectFromID(id objc.ID) ANEStorageMaintainerProtocolObject {
	return ANEStorageMaintainerProtocolObject{
		Object: objectivec.ObjectFromID(id),
	}
}

func (o ANEStorageMaintainerProtocolObject) PurgeDanglingModelsAtWithReply(at objectivec.IObject, reply VoidHandler) {
	_block1, _cleanup1 := NewVoidBlock(reply)
	defer _cleanup1()
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("purgeDanglingModelsAt:withReply:"), at, objc.ID(_block1))
}
