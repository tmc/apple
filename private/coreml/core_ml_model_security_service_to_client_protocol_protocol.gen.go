// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// CoreMLModelSecurityServiceToClientProtocol protocol.
//
// See: https://developer.apple.com/documentation/CoreML/CoreMLModelSecurityServiceToClientProtocol
type CoreMLModelSecurityServiceToClientProtocol interface {
	objectivec.IObject
}

// CoreMLModelSecurityServiceToClientProtocolObject wraps an existing Objective-C object that conforms to the CoreMLModelSecurityServiceToClientProtocol protocol.
type CoreMLModelSecurityServiceToClientProtocolObject struct {
	objectivec.Object
}

func (o CoreMLModelSecurityServiceToClientProtocolObject) BaseObject() objectivec.Object {
	return o.Object
}

// CoreMLModelSecurityServiceToClientProtocolObjectFromID constructs a [CoreMLModelSecurityServiceToClientProtocolObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CoreMLModelSecurityServiceToClientProtocolObjectFromID(id objc.ID) CoreMLModelSecurityServiceToClientProtocolObject {
	return CoreMLModelSecurityServiceToClientProtocolObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// See: https://developer.apple.com/documentation/CoreML/CoreMLModelSecurityServiceToClientProtocol/clientFeatureValueForName:uniqueKeyForProvider:withReply:
func (o CoreMLModelSecurityServiceToClientProtocolObject) ClientFeatureValueForNameUniqueKeyForProviderWithReply(name objectivec.IObject, provider objectivec.IObject, reply MLFeatureValueErrorHandler) {
	objc.Send[struct{}](o.ID, objc.Sel("clientFeatureValueForName:uniqueKeyForProvider:withReply:"), name, provider, reply)
}
