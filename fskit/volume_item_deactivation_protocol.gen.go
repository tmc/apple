// Code generated from Apple documentation for FSKit. DO NOT EDIT.

package fskit

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// Methods and properties implemented by volumes that support deactivating items.
//
// See: https://developer.apple.com/documentation/FSKit/FSVolume/ItemDeactivation
type FSVolumeItemDeactivation interface {
	objectivec.IObject

	// Notifies the file system that the kernel is no longer making immediate use of the given item.
	//
	// See: https://developer.apple.com/documentation/FSKit/FSVolume/ItemDeactivation/deactivateItem(_:replyHandler:)
	DeactivateItemReplyHandler(item IFSItem, reply ErrorHandler)

	// A property that tells FSKit to which types of items the deactivation applies, if any.
	//
	// See: https://developer.apple.com/documentation/FSKit/FSVolume/ItemDeactivation/itemDeactivationPolicy
	ItemDeactivationPolicy() FSItemDeactivationOptions
}

// FSVolumeItemDeactivationObject wraps an existing Objective-C object that conforms to the FSVolumeItemDeactivation protocol.
type FSVolumeItemDeactivationObject struct {
	objectivec.Object
}

func (o FSVolumeItemDeactivationObject) BaseObject() objectivec.Object {
	return o.Object
}

// FSVolumeItemDeactivationObjectFromID constructs a [FSVolumeItemDeactivationObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func FSVolumeItemDeactivationObjectFromID(id objc.ID) FSVolumeItemDeactivationObject {
	return FSVolumeItemDeactivationObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Notifies the file system that the kernel is no longer making immediate use
// of the given item.
//
// item: The item to deactivate.
//
// reply: A block or closure to indicate success or failure. If deactivation fails,
// pass an error as the one parameter to the reply handler. If deactivation
// succeeds, pass `nil`. For an `async` Swift implementation, there’s no
// reply handler; simply throw an error or return normally.
//
// # Discussion
//
// This method gives a file system a chance to release resources associated
// wtih an item. However, this method prescribes no specific action; it’s
// acceptable to defer all reclamation until [ReclaimItemReplyHandler]. This
// method is the equivalent of VFS’s `VNOP_INACTIVE`.
//
// FSKit restricts calls to this method based on the current value of
// [ItemDeactivationPolicy].
//
// See: https://developer.apple.com/documentation/FSKit/FSVolume/ItemDeactivation/deactivateItem(_:replyHandler:)
func (o FSVolumeItemDeactivationObject) DeactivateItemReplyHandler(item IFSItem, reply ErrorHandler) {
	objc.Send[struct{}](o.ID, objc.Sel("deactivateItem:replyHandler:"), item, reply)
}

// A property that tells FSKit to which types of items the deactivation
// applies, if any.
//
// See: https://developer.apple.com/documentation/FSKit/FSVolume/ItemDeactivation/itemDeactivationPolicy
func (o FSVolumeItemDeactivationObject) ItemDeactivationPolicy() FSItemDeactivationOptions {
	rv := objc.Send[FSItemDeactivationOptions](o.ID, objc.Sel("itemDeactivationPolicy"))
	return rv
}
