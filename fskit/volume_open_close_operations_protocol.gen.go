// Code generated from Apple documentation for FSKit. DO NOT EDIT.

package fskit

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// Methods and properties implemented by volumes that want to receive open and close calls for each item.
//
// See: https://developer.apple.com/documentation/FSKit/FSVolume/OpenCloseOperations
type FSVolumeOpenCloseOperations interface {
	objectivec.IObject

	// Opens a file for access.
	//
	// See: https://developer.apple.com/documentation/FSKit/FSVolume/OpenCloseOperations/openItem(_:modes:replyHandler:)
	OpenItemWithModesReplyHandler(item IFSItem, modes FSVolumeOpenModes, reply ErrorHandler)

	// Closes a file from further access.
	//
	// See: https://developer.apple.com/documentation/FSKit/FSVolume/OpenCloseOperations/closeItem(_:modes:replyHandler:)
	CloseItemKeepingModesReplyHandler(item IFSItem, modes FSVolumeOpenModes, reply ErrorHandler)

	// A Boolean value that instructs FSKit not to call this protocol’s methods, even if the volume conforms to it.
	//
	// See: https://developer.apple.com/documentation/FSKit/FSVolume/OpenCloseOperations/isOpenCloseInhibited
	IsOpenCloseInhibited() bool

	// A Boolean value that instructs FSKit not to call this protocol’s methods, even if the volume conforms to it.
	//
	// See: https://developer.apple.com/documentation/FSKit/FSVolume/OpenCloseOperations/isOpenCloseInhibited
	SetOpenCloseInhibited(value bool)
}

// FSVolumeOpenCloseOperationsObject wraps an existing Objective-C object that conforms to the FSVolumeOpenCloseOperations protocol.
type FSVolumeOpenCloseOperationsObject struct {
	objectivec.Object
}

func (o FSVolumeOpenCloseOperationsObject) BaseObject() objectivec.Object {
	return o.Object
}

// FSVolumeOpenCloseOperationsObjectFromID constructs a [FSVolumeOpenCloseOperationsObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func FSVolumeOpenCloseOperationsObjectFromID(id objc.ID) FSVolumeOpenCloseOperationsObject {
	return FSVolumeOpenCloseOperationsObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Opens a file for access.
//
// item: The item to open.
//
// modes: The set of mode flags to open the item with.
//
// reply: A block or closure to indicate success or failure. If opening fails, pass
// an error as the one parameter to the reply handler. If opening succeeds,
// pass `nil`. For an `async` Swift implementation, there’s no reply
// handler; simply throw an error or return normally.
//
// See: https://developer.apple.com/documentation/FSKit/FSVolume/OpenCloseOperations/openItem(_:modes:replyHandler:)
func (o FSVolumeOpenCloseOperationsObject) OpenItemWithModesReplyHandler(item IFSItem, modes FSVolumeOpenModes, reply ErrorHandler) {
	objc.Send[struct{}](o.ID, objc.Sel("openItem:withModes:replyHandler:"), item, modes, reply)
}

// Closes a file from further access.
//
// item: The item to close.
//
// modes: The set of mode flags to keep after this close.
//
// reply: A block or closure to indicate success or failure. If closing fails, pass
// an error as the one parameter to the reply handler. If closing succeeds,
// pass `nil`. For an `async` Swift implementation, there’s no reply
// handler; simply throw an error or return normally.
//
// See: https://developer.apple.com/documentation/FSKit/FSVolume/OpenCloseOperations/closeItem(_:modes:replyHandler:)
func (o FSVolumeOpenCloseOperationsObject) CloseItemKeepingModesReplyHandler(item IFSItem, modes FSVolumeOpenModes, reply ErrorHandler) {
	objc.Send[struct{}](o.ID, objc.Sel("closeItem:keepingModes:replyHandler:"), item, modes, reply)
}

// A Boolean value that instructs FSKit not to call this protocol’s methods,
// even if the volume conforms to it.
//
// See: https://developer.apple.com/documentation/FSKit/FSVolume/OpenCloseOperations/isOpenCloseInhibited
func (o FSVolumeOpenCloseOperationsObject) IsOpenCloseInhibited() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isOpenCloseInhibited"))
	return rv
}

// A Boolean value that instructs FSKit not to call this protocol’s methods,
// even if the volume conforms to it.
//
// # Discussion
//
// FSKit reads this value after the file system replies to the `loadResource`
// message. Changing the returned value during the runtime of the volume has
// no effect.
//
// See: https://developer.apple.com/documentation/FSKit/FSVolume/OpenCloseOperations/isOpenCloseInhibited
func (o FSVolumeOpenCloseOperationsObject) SetOpenCloseInhibited(value bool) {
	objc.Send[struct{}](o.ID, objc.Sel("setOpenCloseInhibited:"), value)
}
