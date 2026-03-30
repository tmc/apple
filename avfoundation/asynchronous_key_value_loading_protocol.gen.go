// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// A protocol that defines the interface to load media data asynchronously.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAsynchronousKeyValueLoading
type AVAsynchronousKeyValueLoading interface {
	objectivec.IObject
}

// AVAsynchronousKeyValueLoadingObject wraps an existing Objective-C object that conforms to the AVAsynchronousKeyValueLoading protocol.
type AVAsynchronousKeyValueLoadingObject struct {
	objectivec.Object
}

func (o AVAsynchronousKeyValueLoadingObject) BaseObject() objectivec.Object {
	return o.Object
}

// AVAsynchronousKeyValueLoadingObjectFromID constructs a [AVAsynchronousKeyValueLoadingObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func AVAsynchronousKeyValueLoadingObjectFromID(id objc.ID) AVAsynchronousKeyValueLoadingObject {
	return AVAsynchronousKeyValueLoadingObject{
		Object: objectivec.ObjectFromID(id),
	}
}
