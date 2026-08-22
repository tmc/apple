// Code generated from Apple documentation for FileProvider. DO NOT EDIT.

package fileprovider

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// An operation that resolves a collision by renaming the new item.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderTestingCollisionResolution
type NSFileProviderTestingCollisionResolution interface {
	objectivec.IObject
	NSFileProviderTestingOperation

	// A description of the renamed item.
	//
	// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderTestingCollisionResolution/renamedItem
	RenamedItem() NSFileProviderItem

	// The item’s location.
	//
	// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderTestingCollisionResolution/side
	Side() NSFileProviderTestingOperationSide
}

// NSFileProviderTestingCollisionResolutionObject wraps an existing Objective-C object that conforms to the NSFileProviderTestingCollisionResolution protocol.
type NSFileProviderTestingCollisionResolutionObject struct {
	objectivec.Object
}

func (o NSFileProviderTestingCollisionResolutionObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSFileProviderTestingCollisionResolutionObjectFromID constructs a [NSFileProviderTestingCollisionResolutionObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSFileProviderTestingCollisionResolutionObjectFromID(id objc.ID) NSFileProviderTestingCollisionResolutionObject {
	return NSFileProviderTestingCollisionResolutionObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The operation’s type.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderTestingOperation/type
func (o NSFileProviderTestingCollisionResolutionObject) Type() NSFileProviderTestingOperationType {
	rv := objc.Send[NSFileProviderTestingOperationType](o.ID, objc.Sel("type"))
	return rv
}

// A description of the renamed item.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderTestingCollisionResolution/renamedItem
func (o NSFileProviderTestingCollisionResolutionObject) RenamedItem() NSFileProviderItem {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("renamedItem"))
	return NSFileProviderItemObjectFromID(rv)
}

// The item’s location.
//
// # Discussion
//
// Most operations are symmetrical. They can affect either items stored
// locally, or items in the File Provider extension’s remote storage.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderTestingCollisionResolution/side
func (o NSFileProviderTestingCollisionResolutionObject) Side() NSFileProviderTestingOperationSide {
	rv := objc.Send[NSFileProviderTestingOperationSide](o.ID, objc.Sel("side"))
	return NSFileProviderTestingOperationSide(rv)
}
