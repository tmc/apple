// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// A protocol for requiring decryption keys for media data.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVContentKeyRecipient
type AVContentKeyRecipient interface {
	objectivec.IObject

	// A Boolean value that indicates whether the recipient requires decryption keys for media data to enable processing.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVContentKeyRecipient/mayRequireContentKeysForMediaDataProcessing
	MayRequireContentKeysForMediaDataProcessing() bool
}

// AVContentKeyRecipientObject wraps an existing Objective-C object that conforms to the AVContentKeyRecipient protocol.
type AVContentKeyRecipientObject struct {
	objectivec.Object
}
func (o AVContentKeyRecipientObject) BaseObject() objectivec.Object {
	return o.Object
}

// AVContentKeyRecipientObjectFromID constructs a [AVContentKeyRecipientObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func AVContentKeyRecipientObjectFromID(id objc.ID) AVContentKeyRecipientObject {
	return AVContentKeyRecipientObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// A Boolean value that indicates whether the recipient requires decryption
// keys for media data to enable processing.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVContentKeyRecipient/mayRequireContentKeysForMediaDataProcessing
func (o AVContentKeyRecipientObject) MayRequireContentKeysForMediaDataProcessing() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("mayRequireContentKeysForMediaDataProcessing"))
	return rv
	}
// Tells the recipient that a content key is available.
//
// contentKeySession: The current content key session.
//
// contentKey: A content key to use with objects that support manual attachment of keys,
// such as [CMSampleBuffer].
// //
// [CMSampleBuffer]: https://developer.apple.com/documentation/CoreMedia/CMSampleBuffer
//
// See: https://developer.apple.com/documentation/AVFoundation/AVContentKeyRecipient/contentKeySession(_:didProvide:)
func (o AVContentKeyRecipientObject) ContentKeySessionDidProvideContentKey(contentKeySession IAVContentKeySession, contentKey IAVContentKey) {
	objc.Send[struct{}](o.ID, objc.Sel("contentKeySession:didProvideContentKey:"), contentKeySession, contentKey)
	}

