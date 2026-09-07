// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// AVAudioRemoteInputPluginDelegate protocol.
type AVAudioRemoteInputPluginDelegate interface {
	objectivec.IObject

	// InputPluginDidPublishDevice protocol.
	InputPluginDidPublishDevice(plugin objectivec.IObject, device objectivec.IObject)

	// InputPluginDidUnpublishDevice protocol.
	InputPluginDidUnpublishDevice(plugin objectivec.IObject, device objectivec.IObject)
}

// AVAudioRemoteInputPluginDelegateObject wraps an existing Objective-C object that conforms to the AVAudioRemoteInputPluginDelegate protocol.
type AVAudioRemoteInputPluginDelegateObject struct {
	objectivec.Object
}

func (o AVAudioRemoteInputPluginDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}

// AVAudioRemoteInputPluginDelegateObjectFromID constructs a [AVAudioRemoteInputPluginDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func AVAudioRemoteInputPluginDelegateObjectFromID(id objc.ID) AVAudioRemoteInputPluginDelegateObject {
	return AVAudioRemoteInputPluginDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
}

func (o AVAudioRemoteInputPluginDelegateObject) InputPluginDidPublishDevice(plugin objectivec.IObject, device objectivec.IObject) {
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("inputPlugin:didPublishDevice:"), plugin, device)
}
func (o AVAudioRemoteInputPluginDelegateObject) InputPluginDidUnpublishDevice(plugin objectivec.IObject, device objectivec.IObject) {
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("inputPlugin:didUnpublishDevice:"), plugin, device)
}
