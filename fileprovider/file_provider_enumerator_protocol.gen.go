// Code generated from Apple documentation for FileProvider. DO NOT EDIT.

package fileprovider

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// A protocol for enumerating items and changes.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderEnumerator
type NSFileProviderEnumerator interface {
	objectivec.IObject

	// Requests the next batch of items, starting at the specified page.
	//
	// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderEnumerator/enumerateItems(for:startingAt:)
	EnumerateItemsForObserverStartingAtPage(observer NSFileProviderEnumerationObserver, page NSFileProviderPage)

	// Stops the enumeration of items and changes.
	//
	// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderEnumerator/invalidate()
	Invalidate()
}

// NSFileProviderEnumeratorObject wraps an existing Objective-C object that conforms to the NSFileProviderEnumerator protocol.
type NSFileProviderEnumeratorObject struct {
	objectivec.Object
}

func (o NSFileProviderEnumeratorObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSFileProviderEnumeratorObjectFromID constructs a [NSFileProviderEnumeratorObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSFileProviderEnumeratorObjectFromID(id objc.ID) NSFileProviderEnumeratorObject {
	return NSFileProviderEnumeratorObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Requests the next batch of items, starting at the specified page.
//
// # Discussion
//
// The system requests an enumerator in the following situations:
//
// - For directories. The system requests an enumerator when a document
// browser displays the contents of a directory. For performance reasons, the
// system may retain the enumerator even after the browser has moved to a
// different directory. - For files. The system requests an enumerator when a
// file presenter begins managing an item. The enumerator is invalidated after
// the file presenter is removed. - For the working set. The system requests
// an enumerator when it begins indexing the working set. It invalidates the
// enumerator after the indexing operation has completed.
//
// Once requested, an enumerator is used to provide both the content and any
// changes for an item. When the system is finished with the item, it calls
// the enumerator’s [Invalidate] method. For example, if you return an
// enumerator that provides the content of a directory, as long as the
// enumerator is active, the system also uses it to enumerate changes to the
// directory.
//
// The system may have multiple, active enumerators for a number of different
// items. Some of these may represent items that are currently displayed on
// screen. Others may be items that are no longer displayed, but the
// enumerator is retained for performance reasons. You need to inform the
// system of any changes to the content managed by any of the active
// enumerators, as well as any changes to the working set (whether or not it
// has an active enumerator).
//
// For more information on enumerating items, see [Defining Your File
// Provider’s Content]. For more information on tracking changes, see
// [Tracking Your File Provider’s Changes].
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderEnumerator/enumerateItems(for:startingAt:)
//
// [Defining Your File Provider’s Content]: https://developer.apple.com/documentation/FileProvider/defining-your-file-provider-s-content
// [Tracking Your File Provider’s Changes]: https://developer.apple.com/documentation/FileProvider/tracking-your-file-provider-s-changes
func (o NSFileProviderEnumeratorObject) EnumerateItemsForObserverStartingAtPage(observer NSFileProviderEnumerationObserver, page NSFileProviderPage) {
	objc.Send[struct{}](o.ID, objc.Sel("enumerateItemsForObserver:startingAtPage:"), observer, page)
}

// Stops the enumeration of items and changes.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderEnumerator/invalidate()
func (o NSFileProviderEnumeratorObject) Invalidate() {
	objc.Send[struct{}](o.ID, objc.Sel("invalidate"))
}

// Requests the next batch of changes after the specified sync anchor.
//
// # Discussion
//
// The system requests an enumerator in the following situations:
//
// - For directories. The system requests an enumerator when a document
// browser displays the contents of a directory. For performance reasons, the
// system may retain the enumerator even after the browser has moved to a
// different directory. - For files. The system requests an enumerator when a
// file presenter begins managing an item. The enumerator is invalidated after
// the file presenter is removed. - For the working set. The system requests
// an enumerator when it begins indexing the working set. It invalidates the
// enumerator after the indexing operation has completed.
//
// Once requested, an enumerator is used to provide both the content and any
// changes for an item. When the system is finished with the item, it calls
// the enumerator’s [Invalidate] method. For example, if you return an
// enumerator that provides the content of a directory, as long as the
// enumerator is active, the system also uses it to enumerate changes to the
// directory.
//
// The system may have multiple, active enumerators for a number of different
// items. Some of these may represent items that are currently displayed on
// screen. Others may be items that are no longer displayed, but the
// enumerator is retained for performance reasons. You need to inform the
// system of any changes to the content managed by any of the active
// enumerators, as well as any changes to the working set (whether or not it
// has an active enumerator).
//
// For more information on enumerating items, see [Defining Your File
// Provider’s Content]. For more information on tracking changes, see
// [Tracking Your File Provider’s Changes].
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderEnumerator/enumerateChanges(for:from:)
//
// [Defining Your File Provider’s Content]: https://developer.apple.com/documentation/FileProvider/defining-your-file-provider-s-content
// [Tracking Your File Provider’s Changes]: https://developer.apple.com/documentation/FileProvider/tracking-your-file-provider-s-changes
func (o NSFileProviderEnumeratorObject) EnumerateChangesForObserverFromSyncAnchor(observer NSFileProviderChangeObserver, syncAnchor NSFileProviderSyncAnchor) {
	objc.Send[struct{}](o.ID, objc.Sel("enumerateChangesForObserver:fromSyncAnchor:"), observer, syncAnchor)
}

// Returns the current sync anchor.
//
// # Discussion
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderEnumerator/currentSyncAnchor(completionHandler:)
func (o NSFileProviderEnumeratorObject) CurrentSyncAnchorWithCompletionHandler(completionHandler DataHandler) {
	_block0, _ := NewDataBlock(completionHandler)
	objc.Send[struct{}](o.ID, objc.Sel("currentSyncAnchorWithCompletionHandler:"), _block0)
}
