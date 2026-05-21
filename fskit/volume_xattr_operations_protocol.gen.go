// Code generated from Apple documentation for FSKit. DO NOT EDIT.

package fskit

import (
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// Methods and properties implemented by volumes that natively or partially support extended attributes.
//
// See: https://developer.apple.com/documentation/FSKit/FSVolume/XattrOperations
type FSVolumeXattrOperations interface {
	objectivec.IObject

	// Gets the specified extended attribute of the given item.
	//
	// See: https://developer.apple.com/documentation/FSKit/FSVolume/XattrOperations/getXattr(named:of:replyHandler:)
	GetXattrNamedOfItemReplyHandler(name IFSFileName, item IFSItem, reply DataErrorHandler)

	// Gets the list of extended attributes currently set on the given item.
	//
	// See: https://developer.apple.com/documentation/FSKit/FSVolume/XattrOperations/listXattrs(of:replyHandler:)
	ListXattrsOfItemReplyHandler(item IFSItem, reply FSFileNameArrayErrorHandler)

	// Sets the specified extended attribute data on the given item.
	//
	// See: https://developer.apple.com/documentation/FSKit/FSVolume/XattrOperations/setXattr(named:to:on:policy:replyHandler:)
	SetXattrNamedToDataOnItemPolicyReplyHandler(name IFSFileName, value foundation.NSData, item IFSItem, policy FSSetXattrPolicy, reply ErrorHandler)

	// A Boolean value that instructs FSKit not to call this protocol’s methods, even if the volume conforms to it.
	//
	// See: https://developer.apple.com/documentation/FSKit/FSVolume/XattrOperations/xattrOperationsInhibited
	XattrOperationsInhibited() bool
	SetXattrOperationsInhibited(value bool)
}

// FSVolumeXattrOperationsObject wraps an existing Objective-C object that conforms to the FSVolumeXattrOperations protocol.
type FSVolumeXattrOperationsObject struct {
	objectivec.Object
}

func (o FSVolumeXattrOperationsObject) BaseObject() objectivec.Object {
	return o.Object
}

// FSVolumeXattrOperationsObjectFromID constructs a [FSVolumeXattrOperationsObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func FSVolumeXattrOperationsObjectFromID(id objc.ID) FSVolumeXattrOperationsObject {
	return FSVolumeXattrOperationsObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Gets the specified extended attribute of the given item.
//
// name: The extended attribute name.
//
// item: The item for which to get the extended attribute.
//
// reply: A block or closure to indicate success or failure. If getting the attribute
// succeeds, pass an data instance containing the extended attribute data and
// a `nil` error. If getting the attribute fails, pass the relevant error as
// the second parameter; FSKit ignores any data in this case. For an `async`
// Swift implementation, there’s no reply handler; simply return the data or
// throw an error.
//
// See: https://developer.apple.com/documentation/FSKit/FSVolume/XattrOperations/getXattr(named:of:replyHandler:)
func (o FSVolumeXattrOperationsObject) GetXattrNamedOfItemReplyHandler(name IFSFileName, item IFSItem, reply DataErrorHandler) {
	objc.Send[struct{}](o.ID, objc.Sel("getXattrNamed:ofItem:replyHandler:"), name, item, reply)
}

// Gets the list of extended attributes currently set on the given item.
//
// item: The item from which to get extended attributes.
//
// reply: A block or closure to indicate success or failure. If getting the list of
// extended attributes succeeds, pass the xattrs as an array of [FSFileName]
// instances and a `nil` error. If getting the attriubtes fails, pass `nil`
// along with the relevant error. For an `async` Swift implementation,
// there’s no reply handler; simply return the byte count or throw an error.
//
// See: https://developer.apple.com/documentation/FSKit/FSVolume/XattrOperations/listXattrs(of:replyHandler:)
func (o FSVolumeXattrOperationsObject) ListXattrsOfItemReplyHandler(item IFSItem, reply FSFileNameArrayErrorHandler) {
	objc.Send[struct{}](o.ID, objc.Sel("listXattrsOfItem:replyHandler:"), item, reply)
}

// Sets the specified extended attribute data on the given item.
//
// name: The extended attribute name.
//
// value: The extended attribute value to set. This can’t be `nil`, unless the
// policy is [FSSetXattrPolicyDelete].
//
// item: The item on which to set the extended attribute.
//
// policy: The policy to apply when setting the attribute. See
// [FSVolume.SetXattrPolicy] for possible values.
//
// reply: A block or closure to indicate success or failure. If setting the attribute
// fails, pass an error as the one parameter to the reply handler. If setting
// the attribute succeeds, pass `nil`. For an `async` Swift implementation,
// there’s no reply handler; simply throw an error or return normally.
//
// See: https://developer.apple.com/documentation/FSKit/FSVolume/XattrOperations/setXattr(named:to:on:policy:replyHandler:)
//
// [FSVolume.SetXattrPolicy]: https://developer.apple.com/documentation/FSKit/FSVolume/SetXattrPolicy
func (o FSVolumeXattrOperationsObject) SetXattrNamedToDataOnItemPolicyReplyHandler(name IFSFileName, value foundation.NSData, item IFSItem, policy FSSetXattrPolicy, reply ErrorHandler) {
	objc.Send[struct{}](o.ID, objc.Sel("setXattrNamed:toData:onItem:policy:replyHandler:"), name, value, item, policy, reply)
}

// Returns an array that specifies the extended attribute names the given item
// supports.
//
// item: The item for which to get information.
//
// # Discussion
//
// If `item` supports no extended attributes, this method returns `nil`.
//
// Only implement this method if your volume works with “limited” extended
// attributes. For purposes of this protocol, “limited” support means the
// volume doesn’t support extended attributes generally, but uses these APIs
// to expose specific file system data.
//
// See: https://developer.apple.com/documentation/FSKit/FSVolume/XattrOperations/supportedXattrNames(for:)
func (o FSVolumeXattrOperationsObject) SupportedXattrNamesForItem(item IFSItem) []FSFileName {
	rv := objc.Send[[]objc.ID](o.ID, objc.Sel("supportedXattrNamesForItem:"), item)
	return objc.ConvertSlice(rv, func(id objc.ID) FSFileName {
		return FSFileNameFromID(id)
	})
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
// See: https://developer.apple.com/documentation/FSKit/FSVolume/XattrOperations/xattrOperationsInhibited
func (o FSVolumeXattrOperationsObject) XattrOperationsInhibited() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("xattrOperationsInhibited"))
	return bool(rv)
}

func (o FSVolumeXattrOperationsObject) SetXattrOperationsInhibited(value bool) {
	objc.Send[struct{}](o.ID, objc.Sel("setXattrOperationsInhibited:"), value)
}
