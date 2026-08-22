// Code generated from Apple documentation for FileProvider. DO NOT EDIT.

package fileprovider

import (
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// An observer that receives changes and deletions during enumeration.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderChangeObserver
type NSFileProviderChangeObserver interface {
	objectivec.IObject

	// Tells the observer that the specified items have been deleted.
	//
	// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderChangeObserver/didDeleteItems(withIdentifiers:)
	DidDeleteItemsWithIdentifiers(deletedItemIdentifiers []string)

	// Tells the observer that the specified items have been updated.
	//
	// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderChangeObserver/didUpdate(_:)
	DidUpdateItems(updatedItems []objectivec.IObject)

	// Tells the observer that all of the changes have been enumerated up to the specified sync anchor.
	//
	// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderChangeObserver/finishEnumeratingChanges(upTo:moreComing:)
	FinishEnumeratingChangesUpToSyncAnchorMoreComing(anchor NSFileProviderSyncAnchor, moreComing bool)

	// Tells the observer that an error occurred during change notification.
	//
	// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderChangeObserver/finishEnumeratingWithError(_:)
	FinishEnumeratingWithError(error_ foundation.NSError)

	// The batch size that the system recommends.
	//
	// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderChangeObserver/suggestedBatchSize
	SuggestedBatchSize() int
}

// NSFileProviderChangeObserverObject wraps an existing Objective-C object that conforms to the NSFileProviderChangeObserver protocol.
type NSFileProviderChangeObserverObject struct {
	objectivec.Object
}

func (o NSFileProviderChangeObserverObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSFileProviderChangeObserverObjectFromID constructs a [NSFileProviderChangeObserverObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSFileProviderChangeObserverObjectFromID(id objc.ID) NSFileProviderChangeObserverObject {
	return NSFileProviderChangeObserverObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Tells the observer that the specified items have been deleted.
//
// deletedItemIdentifiers: An array of identifiers for the deleted items.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderChangeObserver/didDeleteItems(withIdentifiers:)
func (o NSFileProviderChangeObserverObject) DidDeleteItemsWithIdentifiers(deletedItemIdentifiers []string) {
	objc.Send[struct{}](o.ID, objc.Sel("didDeleteItemsWithIdentifiers:"), objectivec.StringSliceToNSArray(deletedItemIdentifiers))
}

// Tells the observer that the specified items have been updated.
//
// updatedItems: An array of updated items.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderChangeObserver/didUpdate(_:)
func (o NSFileProviderChangeObserverObject) DidUpdateItems(updatedItems []objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("didUpdateItems:"), objectivec.IObjectSliceToNSArray(updatedItems))
}

// Tells the observer that all of the changes have been enumerated up to the
// specified sync anchor.
//
// anchor: An object used to identify the end of the current batch of changes.
//
// moreComing: A Boolean value that indicates the file provider still has one or more
// batches of pending changes.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderChangeObserver/finishEnumeratingChanges(upTo:moreComing:)
func (o NSFileProviderChangeObserverObject) FinishEnumeratingChangesUpToSyncAnchorMoreComing(anchor NSFileProviderSyncAnchor, moreComing bool) {
	objc.Send[struct{}](o.ID, objc.Sel("finishEnumeratingChangesUpToSyncAnchor:moreComing:"), anchor, moreComing)
}

// Tells the observer that an error occurred during change notification.
//
// error: An object that contains information about the error.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderChangeObserver/finishEnumeratingWithError(_:)
func (o NSFileProviderChangeObserverObject) FinishEnumeratingWithError(error_ foundation.NSError) {
	objc.Send[struct{}](o.ID, objc.Sel("finishEnumeratingWithError:"), error_)
}

// The batch size that the system recommends.
//
// # Discussion
//
// The system suggests the batch size to optimize performance based on the
// context of the pending changes. The system can request the enumeration of a
// container for various reasons, such as if the user has the directory open
// in Finder, or the file open in an application. Each case has its own
// performance profile.
//
// If the enumerator has more pending changes than the suggested batch size,
// it should split the changes into batches that are equal to or smaller than
// the batch size. If the enumerator has fewer changes than the suggested
// batch size, return all the changes immediately and finish the enumeration.
// You don’t need to wait for more incoming changes.
//
// While using the suggested batch size helps ensure the best user experience,
// the system enforces a maximum of 100 times the suggested size.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderChangeObserver/suggestedBatchSize
func (o NSFileProviderChangeObserverObject) SuggestedBatchSize() int {
	rv := objc.Send[int](o.ID, objc.Sel("suggestedBatchSize"))
	return int(rv)
}
