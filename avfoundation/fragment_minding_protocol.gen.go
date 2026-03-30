// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// A protocol that defines whether an asset supports fragment minding.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVFragmentMinding
type AVFragmentMinding interface {
	objectivec.IObject

	// A Boolean value that indicates whether an asset that supports fragment minding is currently associated with a fragment minder.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVFragmentMinding/isAssociatedWithFragmentMinder
	IsAssociatedWithFragmentMinder() bool
}

// AVFragmentMindingObject wraps an existing Objective-C object that conforms to the AVFragmentMinding protocol.
type AVFragmentMindingObject struct {
	objectivec.Object
}

func (o AVFragmentMindingObject) BaseObject() objectivec.Object {
	return o.Object
}

// AVFragmentMindingObjectFromID constructs a [AVFragmentMindingObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func AVFragmentMindingObjectFromID(id objc.ID) AVFragmentMindingObject {
	return AVFragmentMindingObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// A Boolean value that indicates whether an asset that supports fragment
// minding is currently associated with a fragment minder.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVFragmentMinding/isAssociatedWithFragmentMinder
func (o AVFragmentMindingObject) IsAssociatedWithFragmentMinder() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAssociatedWithFragmentMinder"))
	return rv
}
