// Code generated from Apple documentation for FileProvider. DO NOT EDIT.

package fileprovider

import (
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// A protocol for enumerating pending items.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderPendingSetEnumerator
type NSFileProviderPendingSetEnumerator interface {
	objectivec.IObject
	NSFileProviderEnumerator

	// The domain version when the system last refreshed the pending set.
	//
	// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderPendingSetEnumerator/domainVersion
	DomainVersion() INSFileProviderDomainVersion

	// The amount of time, in seconds, between updates to the pending set.
	//
	// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderPendingSetEnumerator/refreshInterval
	RefreshInterval() foundation.NSTimeInterval

	// maximumSizeReached protocol.
	//
	// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderPendingSetEnumerator/isMaximumSizeReached
	IsMaximumSizeReached() bool
}

// NSFileProviderPendingSetEnumeratorObject wraps an existing Objective-C object that conforms to the NSFileProviderPendingSetEnumerator protocol.
type NSFileProviderPendingSetEnumeratorObject struct {
	objectivec.Object
}

func (o NSFileProviderPendingSetEnumeratorObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSFileProviderPendingSetEnumeratorObjectFromID constructs a [NSFileProviderPendingSetEnumeratorObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSFileProviderPendingSetEnumeratorObjectFromID(id objc.ID) NSFileProviderPendingSetEnumeratorObject {
	return NSFileProviderPendingSetEnumeratorObject{
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
func (o NSFileProviderPendingSetEnumeratorObject) EnumerateItemsForObserverStartingAtPage(observer NSFileProviderEnumerationObserver, page NSFileProviderPage) {
	objc.Send[struct{}](o.ID, objc.Sel("enumerateItemsForObserver:startingAtPage:"), observer, page)
}

// Stops the enumeration of items and changes.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderEnumerator/invalidate()
func (o NSFileProviderPendingSetEnumeratorObject) Invalidate() {
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
func (o NSFileProviderPendingSetEnumeratorObject) EnumerateChangesForObserverFromSyncAnchor(observer NSFileProviderChangeObserver, syncAnchor NSFileProviderSyncAnchor) {
	objc.Send[struct{}](o.ID, objc.Sel("enumerateChangesForObserver:fromSyncAnchor:"), observer, syncAnchor)
}

// Returns the current sync anchor.
//
// # Discussion
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderEnumerator/currentSyncAnchor(completionHandler:)
func (o NSFileProviderPendingSetEnumeratorObject) CurrentSyncAnchorWithCompletionHandler(completionHandler DataHandler) {
	_block0, _ := NewDataBlock(completionHandler)
	objc.Send[struct{}](o.ID, objc.Sel("currentSyncAnchorWithCompletionHandler:"), _block0)
}

// The domain version when the system last refreshed the pending set.
//
// # Discussion
//
// The system sets this property when you call the enumerator’s methods. The
// value is initially `nil`.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderPendingSetEnumerator/domainVersion
func (o NSFileProviderPendingSetEnumeratorObject) DomainVersion() INSFileProviderDomainVersion {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("domainVersion"))
	return NSFileProviderDomainVersionFromID(rv)
}

// The amount of time, in seconds, between updates to the pending set.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderPendingSetEnumerator/refreshInterval
func (o NSFileProviderPendingSetEnumeratorObject) RefreshInterval() foundation.NSTimeInterval {
	rv := objc.Send[foundation.NSTimeInterval](o.ID, objc.Sel("refreshInterval"))
	return foundation.NSTimeInterval(rv)
}

// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderPendingSetEnumerator/isMaximumSizeReached
func (o NSFileProviderPendingSetEnumeratorObject) IsMaximumSizeReached() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isMaximumSizeReached"))
	return bool(rv)
}
