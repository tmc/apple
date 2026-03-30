// Code generated from Apple documentation for appleneuralengine. DO NOT EDIT.

package appleneuralengine

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// _ANEStorageMaintainerProtocol protocol.
//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEStorageMaintainerProtocol
type ANEStorageMaintainerProtocol interface {
	objectivec.IObject
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

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEStorageMaintainerProtocol/purgeDanglingModelsAt:withReply:
func (o ANEStorageMaintainerProtocolObject) PurgeDanglingModelsAtWithReply(at objectivec.IObject, reply VoidHandler) {
	objc.Send[struct{}](o.ID, objc.Sel("purgeDanglingModelsAt:withReply:"), at, reply)
}
