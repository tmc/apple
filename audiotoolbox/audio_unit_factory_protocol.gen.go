// Code generated from Apple documentation for AudioToolbox. DO NOT EDIT.

package audiotoolbox

import (
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// An object that creates a version 3 audio unit.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnitFactory
type AUAudioUnitFactory interface {
	objectivec.IObject
	foundation.NSExtensionRequestHandling

	// Creates an instance of an extension’s audio unit.
	//
	// See: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnitFactory/createAudioUnit(with:)
	CreateAudioUnitWithComponentDescriptionError(desc AudioComponentDescription) (IAUAudioUnit, error)
}

// AUAudioUnitFactoryObject wraps an existing Objective-C object that conforms to the AUAudioUnitFactory protocol.
type AUAudioUnitFactoryObject struct {
	foundation.NSExtensionRequestHandlingObject
}

func (o AUAudioUnitFactoryObject) BaseObject() objectivec.Object {
	return o.NSExtensionRequestHandlingObject.BaseObject()
}

// AUAudioUnitFactoryObjectFromID constructs a [AUAudioUnitFactoryObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func AUAudioUnitFactoryObjectFromID(id objc.ID) AUAudioUnitFactoryObject {
	return AUAudioUnitFactoryObject{
		NSExtensionRequestHandlingObject: foundation.NSExtensionRequestHandlingObjectFromID(id),
	}
}

// Creates an instance of an extension’s audio unit.
//
// desc: The description of the audio component.
//
// # Return Value
//
// An instance of the extension’s audio unit.
//
// # Discussion
//
// This method is called only once per factory instance.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnitFactory/createAudioUnit(with:)
func (o AUAudioUnitFactoryObject) CreateAudioUnitWithComponentDescriptionError(desc AudioComponentDescription) (IAUAudioUnit, error) {
	rv, err := objc.SendWithError[objc.ID](o.ID, objc.Sel("createAudioUnitWithComponentDescription:error:"), desc)
	if err != nil {
		return nil, err
	}
	return AUAudioUnitFromID(rv), nil
}
