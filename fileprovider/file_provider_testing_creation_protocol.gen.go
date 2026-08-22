// Code generated from Apple documentation for FileProvider. DO NOT EDIT.

package fileprovider

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// An operation that syncs the creation of the source item to the target location.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderTestingCreation
type NSFileProviderTestingCreation interface {
	objectivec.IObject
	NSFileProviderTestingOperation

	// A description of the item stored at the source.
	//
	// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderTestingCreation/sourceItem
	SourceItem() NSFileProviderItem

	// The target location for the new item.
	//
	// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderTestingCreation/targetSide
	TargetSide() NSFileProviderTestingOperationSide

	// The domain’s version when the system discovered the item at the source location.
	//
	// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderTestingCreation/domainVersion
	DomainVersion() INSFileProviderDomainVersion
}

// NSFileProviderTestingCreationObject wraps an existing Objective-C object that conforms to the NSFileProviderTestingCreation protocol.
type NSFileProviderTestingCreationObject struct {
	objectivec.Object
}

func (o NSFileProviderTestingCreationObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSFileProviderTestingCreationObjectFromID constructs a [NSFileProviderTestingCreationObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSFileProviderTestingCreationObjectFromID(id objc.ID) NSFileProviderTestingCreationObject {
	return NSFileProviderTestingCreationObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The operation’s type.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderTestingOperation/type
func (o NSFileProviderTestingCreationObject) Type() NSFileProviderTestingOperationType {
	rv := objc.Send[NSFileProviderTestingOperationType](o.ID, objc.Sel("type"))
	return rv
}

// A description of the item stored at the source.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderTestingCreation/sourceItem
func (o NSFileProviderTestingCreationObject) SourceItem() NSFileProviderItem {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("sourceItem"))
	return NSFileProviderItemObjectFromID(rv)
}

// The target location for the new item.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderTestingCreation/targetSide
func (o NSFileProviderTestingCreationObject) TargetSide() NSFileProviderTestingOperationSide {
	rv := objc.Send[NSFileProviderTestingOperationSide](o.ID, objc.Sel("targetSide"))
	return NSFileProviderTestingOperationSide(rv)
}

// The domain’s version when the system discovered the item at the source
// location.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderTestingCreation/domainVersion
func (o NSFileProviderTestingCreationObject) DomainVersion() INSFileProviderDomainVersion {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("domainVersion"))
	return NSFileProviderDomainVersionFromID(rv)
}
