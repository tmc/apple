// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// AVAudioRemoteInputPlugin_PrivateTest protocol.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioRemoteInputPlugin_PrivateTest
type AVAudioRemoteInputPlugin_PrivateTest interface {
	objectivec.IObject
}

// AVAudioRemoteInputPlugin_PrivateTestObject wraps an existing Objective-C object that conforms to the AVAudioRemoteInputPlugin_PrivateTest protocol.
type AVAudioRemoteInputPlugin_PrivateTestObject struct {
	objectivec.Object
}

func (o AVAudioRemoteInputPlugin_PrivateTestObject) BaseObject() objectivec.Object {
	return o.Object
}

// AVAudioRemoteInputPlugin_PrivateTestObjectFromID constructs a [AVAudioRemoteInputPlugin_PrivateTestObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func AVAudioRemoteInputPlugin_PrivateTestObjectFromID(id objc.ID) AVAudioRemoteInputPlugin_PrivateTestObject {
	return AVAudioRemoteInputPlugin_PrivateTestObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioRemoteInputPlugin_PrivateTest/mockPluginEndpoint
func (o AVAudioRemoteInputPlugin_PrivateTestObject) MockPluginEndpoint() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("mockPluginEndpoint"))
	return objectivec.Object{ID: rv}
}
