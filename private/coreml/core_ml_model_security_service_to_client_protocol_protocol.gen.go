// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// CoreMLModelSecurityServiceToClientProtocol protocol.
type CoreMLModelSecurityServiceToClientProtocol interface {
	objectivec.IObject

	// ClientFeatureValueForNameUniqueKeyForProviderWithReply protocol.
	ClientFeatureValueForNameUniqueKeyForProviderWithReply(name objectivec.IObject, provider objectivec.IObject, reply MLFeatureValueErrorHandler)
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

func (o CoreMLModelSecurityServiceToClientProtocolObject) ClientFeatureValueForNameUniqueKeyForProviderWithReply(name objectivec.IObject, provider objectivec.IObject, reply MLFeatureValueErrorHandler) {
	_block2, _cleanup2 := NewMLFeatureValueErrorBlock(reply)
	defer _cleanup2()
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("clientFeatureValueForName:uniqueKeyForProvider:withReply:"), name, provider, objc.ID(_block2))
}
