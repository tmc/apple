// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// A protocol that a Cloud-sharing toolbar item uses to get validation of an item.
//
// See: https://developer.apple.com/documentation/AppKit/NSCloudSharingValidation
type NSCloudSharingValidation interface {
	objectivec.IObject

	// Returns the Cloud share object that corresponds to the specified item, if one exists.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSCloudSharingValidation/cloudShare(for:)
	CloudShareForUserInterfaceItem(item NSValidatedUserInterfaceItem) unsafe.Pointer
}

// NSCloudSharingValidationObject wraps an existing Objective-C object that conforms to the NSCloudSharingValidation protocol.
type NSCloudSharingValidationObject struct {
	objectivec.Object
}

func (o NSCloudSharingValidationObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSCloudSharingValidationObjectFromID constructs a [NSCloudSharingValidationObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSCloudSharingValidationObjectFromID(id objc.ID) NSCloudSharingValidationObject {
	return NSCloudSharingValidationObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Returns the Cloud share object that corresponds to the specified item, if
// one exists.
//
// See: https://developer.apple.com/documentation/AppKit/NSCloudSharingValidation/cloudShare(for:)
func (o NSCloudSharingValidationObject) CloudShareForUserInterfaceItem(item NSValidatedUserInterfaceItem) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o.ID, objc.Sel("cloudShareForUserInterfaceItem:"), item)
	return rv
}
