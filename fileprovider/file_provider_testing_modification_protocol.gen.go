// Code generated from Apple documentation for FileProvider. DO NOT EDIT.

package fileprovider

import (
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// An operation that syncs the modification of the source item to the target location.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderTestingModification
type NSFileProviderTestingModification interface {
	objectivec.IObject
	NSFileProviderTestingOperation

	// A description of the source item.
	//
	// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderTestingModification/sourceItem
	SourceItem() NSFileProviderItem

	// A list of the fields that changed.
	//
	// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderTestingModification/changedFields
	ChangedFields() NSFileProviderItemFields

	// The target location for the modification operation.
	//
	// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderTestingModification/targetSide
	TargetSide() NSFileProviderTestingOperationSide

	// The unique identifier for the target item.
	//
	// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderTestingModification/targetItemIdentifier
	TargetItemIdentifier() NSFileProviderItemIdentifier

	// The version of the changed item.
	//
	// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderTestingModification/targetItemBaseVersion
	TargetItemBaseVersion() INSFileProviderItemVersion

	// The domain’s version when the change occurred.
	//
	// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderTestingModification/domainVersion
	DomainVersion() INSFileProviderDomainVersion
}

// NSFileProviderTestingModificationObject wraps an existing Objective-C object that conforms to the NSFileProviderTestingModification protocol.
type NSFileProviderTestingModificationObject struct {
	objectivec.Object
}

func (o NSFileProviderTestingModificationObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSFileProviderTestingModificationObjectFromID constructs a [NSFileProviderTestingModificationObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSFileProviderTestingModificationObjectFromID(id objc.ID) NSFileProviderTestingModificationObject {
	return NSFileProviderTestingModificationObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The operation’s type.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderTestingOperation/type
func (o NSFileProviderTestingModificationObject) Type() NSFileProviderTestingOperationType {
	rv := objc.Send[NSFileProviderTestingOperationType](o.ID, objc.Sel("type"))
	return rv
}

// A description of the source item.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderTestingModification/sourceItem
func (o NSFileProviderTestingModificationObject) SourceItem() NSFileProviderItem {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("sourceItem"))
	return NSFileProviderItemObjectFromID(rv)
}

// A list of the fields that changed.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderTestingModification/changedFields
func (o NSFileProviderTestingModificationObject) ChangedFields() NSFileProviderItemFields {
	rv := objc.Send[NSFileProviderItemFields](o.ID, objc.Sel("changedFields"))
	return NSFileProviderItemFields(rv)
}

// The target location for the modification operation.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderTestingModification/targetSide
func (o NSFileProviderTestingModificationObject) TargetSide() NSFileProviderTestingOperationSide {
	rv := objc.Send[NSFileProviderTestingOperationSide](o.ID, objc.Sel("targetSide"))
	return NSFileProviderTestingOperationSide(rv)
}

// The unique identifier for the target item.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderTestingModification/targetItemIdentifier
func (o NSFileProviderTestingModificationObject) TargetItemIdentifier() NSFileProviderItemIdentifier {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("targetItemIdentifier"))
	return NSFileProviderItemIdentifier(foundation.NSStringFromID(rv).String())
}

// The version of the changed item.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderTestingModification/targetItemBaseVersion
func (o NSFileProviderTestingModificationObject) TargetItemBaseVersion() INSFileProviderItemVersion {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("targetItemBaseVersion"))
	return NSFileProviderItemVersionFromID(rv)
}

// The domain’s version when the change occurred.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderTestingModification/domainVersion
func (o NSFileProviderTestingModificationObject) DomainVersion() INSFileProviderDomainVersion {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("domainVersion"))
	return NSFileProviderDomainVersionFromID(rv)
}
