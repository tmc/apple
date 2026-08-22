// Code generated from Apple documentation for AudioToolbox. DO NOT EDIT.

package audiotoolbox

import (
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// A specification for a bidirectional communication message channel.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUMessageChannel
type AUMessageChannel interface {
	objectivec.IObject

	// A callback for the audio unit to send a message to the host.
	//
	// See: https://developer.apple.com/documentation/AudioToolbox/AUMessageChannel/callHostBlock
	CallHostBlock() CallHostBlock
	SetCallHostBlock(value objc.ID)
}

// AUMessageChannelObject wraps an existing Objective-C object that conforms to the AUMessageChannel protocol.
type AUMessageChannelObject struct {
	objectivec.Object
}

func (o AUMessageChannelObject) BaseObject() objectivec.Object {
	return o.Object
}

// AUMessageChannelObjectFromID constructs a [AUMessageChannelObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func AUMessageChannelObjectFromID(id objc.ID) AUMessageChannelObject {
	return AUMessageChannelObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Sends an audio unit a custom data message.
//
// message: The data to send the audio unit.
//
// # Return Value
//
// A dictionary with custom data.
//
// # Discussion
//
// The valid values for key and value types are [NSArray], [NSDictionary],
// [NSOrderedSet], [NSSet], [NSString], [NSData], [NSNull], [NSNumber], and
// [NSDate].
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUMessageChannel/callAudioUnit(_:)
func (o AUMessageChannelObject) CallAudioUnit(message foundation.INSDictionary) foundation.INSDictionary {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("callAudioUnit:"), message)
	return foundation.NSDictionaryFromID(rv)
}

// A callback for the audio unit to send a message to the host.
//
// # Discussion
//
// The host must set a block on this property to use it.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUMessageChannel/callHostBlock
func (o AUMessageChannelObject) CallHostBlock() CallHostBlock {
	rv := objc.Send[CallHostBlock](o.ID, objc.Sel("callHostBlock"))
	return CallHostBlock(rv)
}

func (o AUMessageChannelObject) SetCallHostBlock(value objc.ID) {
	objc.Send[struct{}](o.ID, objc.Sel("setCallHostBlock:"), value)
}
