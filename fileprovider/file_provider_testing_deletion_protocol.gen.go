// Code generated from Apple documentation for FileProvider. DO NOT EDIT.

package fileprovider

import (
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// An operation that syncs the deletion of the source item to the target location.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderTestingDeletion
type NSFileProviderTestingDeletion interface {
	objectivec.IObject
	NSFileProviderTestingOperation

	// The unique identifier for the source item.
	//
	// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderTestingDeletion/sourceItemIdentifier
	SourceItemIdentifier() NSFileProviderItemIdentifier

	// The target location for the delete operation.
	//
	// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderTestingDeletion/targetSide
	TargetSide() NSFileProviderTestingOperationSide

	// The unique identifier for the target item.
	//
	// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderTestingDeletion/targetItemIdentifier
	TargetItemIdentifier() NSFileProviderItemIdentifier

	// The version of the deleted item.
	//
	// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderTestingDeletion/targetItemBaseVersion
	TargetItemBaseVersion() INSFileProviderItemVersion

	// The domain’s version when the source location deleted the item.
	//
	// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderTestingDeletion/domainVersion
	DomainVersion() INSFileProviderDomainVersion
}

// NSFileProviderTestingDeletionObject wraps an existing Objective-C object that conforms to the NSFileProviderTestingDeletion protocol.
type NSFileProviderTestingDeletionObject struct {
	objectivec.Object
}

func (o NSFileProviderTestingDeletionObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSFileProviderTestingDeletionObjectFromID constructs a [NSFileProviderTestingDeletionObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSFileProviderTestingDeletionObjectFromID(id objc.ID) NSFileProviderTestingDeletionObject {
	return NSFileProviderTestingDeletionObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The operation’s type.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderTestingOperation/type
func (o NSFileProviderTestingDeletionObject) Type() NSFileProviderTestingOperationType {
	rv := objc.Send[NSFileProviderTestingOperationType](o.ID, objc.Sel("type"))
	return rv
}

// The unique identifier for the source item.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderTestingDeletion/sourceItemIdentifier
func (o NSFileProviderTestingDeletionObject) SourceItemIdentifier() NSFileProviderItemIdentifier {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("sourceItemIdentifier"))
	return NSFileProviderItemIdentifier(foundation.NSStringFromID(rv).String())
}

// The target location for the delete operation.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderTestingDeletion/targetSide
func (o NSFileProviderTestingDeletionObject) TargetSide() NSFileProviderTestingOperationSide {
	rv := objc.Send[NSFileProviderTestingOperationSide](o.ID, objc.Sel("targetSide"))
	return NSFileProviderTestingOperationSide(rv)
}

// The unique identifier for the target item.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderTestingDeletion/targetItemIdentifier
func (o NSFileProviderTestingDeletionObject) TargetItemIdentifier() NSFileProviderItemIdentifier {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("targetItemIdentifier"))
	return NSFileProviderItemIdentifier(foundation.NSStringFromID(rv).String())
}

// The version of the deleted item.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderTestingDeletion/targetItemBaseVersion
func (o NSFileProviderTestingDeletionObject) TargetItemBaseVersion() INSFileProviderItemVersion {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("targetItemBaseVersion"))
	return NSFileProviderItemVersionFromID(rv)
}

// The domain’s version when the source location deleted the item.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderTestingDeletion/domainVersion
func (o NSFileProviderTestingDeletionObject) DomainVersion() INSFileProviderDomainVersion {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("domainVersion"))
	return NSFileProviderDomainVersionFromID(rv)
}
