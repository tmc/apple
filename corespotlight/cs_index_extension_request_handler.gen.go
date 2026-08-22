// Code generated from Apple documentation for CoreSpotlight. DO NOT EDIT.

package corespotlight

import (
	"context"
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CSIndexExtensionRequestHandler] class.
var (
	_CSIndexExtensionRequestHandlerClass     CSIndexExtensionRequestHandlerClass
	_CSIndexExtensionRequestHandlerClassOnce sync.Once
)

func getCSIndexExtensionRequestHandlerClass() CSIndexExtensionRequestHandlerClass {
	_CSIndexExtensionRequestHandlerClassOnce.Do(func() {
		_CSIndexExtensionRequestHandlerClass = CSIndexExtensionRequestHandlerClass{class: objc.GetClass("CSIndexExtensionRequestHandler")}
	})
	return _CSIndexExtensionRequestHandlerClass
}

// GetCSIndexExtensionRequestHandlerClass returns the class object for CSIndexExtensionRequestHandler.
func GetCSIndexExtensionRequestHandlerClass() CSIndexExtensionRequestHandlerClass {
	return getCSIndexExtensionRequestHandlerClass()
}

type CSIndexExtensionRequestHandlerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CSIndexExtensionRequestHandlerClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CSIndexExtensionRequestHandlerClass) Alloc() CSIndexExtensionRequestHandler {
	rv := objc.Send[CSIndexExtensionRequestHandler](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// An interface that implements an index-maintenance app extension.
//
// # Overview
//
// The [CSIndexExtensionRequestHandler] class provides the main entry point
// for an index-maintenance app extension. If any issues arise with your
// app’s indexes and your app isn’t running, the system loads your app
// extension and looks for an implementation of this class. It instantiates
// the class it finds and uses it to perform any index-related maintenance.
//
// Define a custom subclass of [CSIndexExtensionRequestHandler] in your app
// extension and implement methods of the [CSSearchableIndexDelegate] protocol
// in it. Use those methods to perform any required updates to your app’s
// index files. For example, use the
// [CSIndexExtensionRequestHandler.SearchableIndexReindexAllSearchableItemsWithAcknowledgementHandler] method
// to reindex all items in your app.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSIndexExtensionRequestHandler
type CSIndexExtensionRequestHandler struct {
	objectivec.Object
}

// CSIndexExtensionRequestHandlerFromID constructs a [CSIndexExtensionRequestHandler] from an objc.ID.
//
// An interface that implements an index-maintenance app extension.
func CSIndexExtensionRequestHandlerFromID(id objc.ID) CSIndexExtensionRequestHandler {
	return CSIndexExtensionRequestHandler{objectivec.Object{ID: id}}
}

// NOTE: CSIndexExtensionRequestHandler adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CSIndexExtensionRequestHandler] class.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSIndexExtensionRequestHandler
type ICSIndexExtensionRequestHandler interface {
	objectivec.IObject

	DataForSearchableIndexItemIdentifierTypeIdentifierError(searchableIndex ICSSearchableIndex, itemIdentifier string, typeIdentifier string) (foundation.NSData, error)
	FileURLForSearchableIndexItemIdentifierTypeIdentifierInPlaceError(searchableIndex ICSSearchableIndex, itemIdentifier string, typeIdentifier string, inPlace bool) (foundation.NSURL, error)
	// Tells the delegate that the index throttling has finished.
	SearchableIndexDidFinishThrottle(searchableIndex ICSSearchableIndex)
	// Tells the delegate that indexing is being throttled.
	SearchableIndexDidThrottle(searchableIndex ICSSearchableIndex)
	// Tells the delegate to reindex all searchable data and clear all local state information.
	SearchableIndexReindexAllSearchableItemsWithAcknowledgementHandler(searchableIndex ICSSearchableIndex, acknowledgementHandler VoidHandler)
	// Tells the delegate to reindex the searchable items associated with the specified identifiers.
	SearchableIndexReindexSearchableItemsWithIdentifiersAcknowledgementHandler(searchableIndex ICSSearchableIndex, identifiers []string, acknowledgementHandler VoidHandler)
	// Tells the delegate that the framework updated the list of searchable items.
	SearchableItemsDidUpdate(items []CSSearchableItem)
	SearchableItemsForIdentifiersProtectionClassSearchableItemsHandler(identifiers []string, protectionClass foundation.NSFileProtectionType, searchableItemsHandler CSSearchableItemArrayHandler)
	// Requests that the delegate provide searchable items for the provided identifiers.
	SearchableItemsForIdentifiersSearchableItemsHandler(identifiers []string, searchableItemsHandler CSSearchableItemArrayHandler)
}

// Init initializes the instance.
func (c CSIndexExtensionRequestHandler) Init() CSIndexExtensionRequestHandler {
	rv := objc.Send[CSIndexExtensionRequestHandler](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CSIndexExtensionRequestHandler) Autorelease() CSIndexExtensionRequestHandler {
	rv := objc.Send[CSIndexExtensionRequestHandler](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCSIndexExtensionRequestHandler creates a new CSIndexExtensionRequestHandler instance.
func NewCSIndexExtensionRequestHandler() CSIndexExtensionRequestHandler {
	class := getCSIndexExtensionRequestHandlerClass()
	rv := objc.Send[CSIndexExtensionRequestHandler](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableIndexDelegate/data(for:itemIdentifier:typeIdentifier:)
func (c CSIndexExtensionRequestHandler) DataForSearchableIndexItemIdentifierTypeIdentifierError(searchableIndex ICSSearchableIndex, itemIdentifier string, typeIdentifier string) (foundation.NSData, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](c.ID, objc.Sel("dataForSearchableIndex:itemIdentifier:typeIdentifier:error:"), searchableIndex, objc.String(itemIdentifier), objc.String(typeIdentifier), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return foundation.NSData{}, foundation.NSErrorFrom(errorPtr)
	}
	return foundation.NSDataFromID(rv), nil

}

// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableIndexDelegate/fileURL(for:itemIdentifier:typeIdentifier:inPlace:)
func (c CSIndexExtensionRequestHandler) FileURLForSearchableIndexItemIdentifierTypeIdentifierInPlaceError(searchableIndex ICSSearchableIndex, itemIdentifier string, typeIdentifier string, inPlace bool) (foundation.NSURL, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](c.ID, objc.Sel("fileURLForSearchableIndex:itemIdentifier:typeIdentifier:inPlace:error:"), searchableIndex, objc.String(itemIdentifier), objc.String(typeIdentifier), inPlace, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return foundation.NSURL{}, foundation.NSErrorFrom(errorPtr)
	}
	return foundation.NSURLFromID(rv), nil

}

// Tells the delegate that the index throttling has finished.
//
// searchableIndex: The index that was throttled.
//
// # Discussion
//
// In some situations, such as when the device is using battery only, the
// system may throttle indexing to save power. You can implement this method
// to be notified when throttling is finished so that your app can resume its
// standard indexing behavior.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableIndexDelegate/searchableIndexDidFinishThrottle(_:)
func (c CSIndexExtensionRequestHandler) SearchableIndexDidFinishThrottle(searchableIndex ICSSearchableIndex) {
	objc.Send[objc.ID](c.ID, objc.Sel("searchableIndexDidFinishThrottle:"), searchableIndex)
}

// Tells the delegate that indexing is being throttled.
//
// searchableIndex: The indexing that’s being throttled.
//
// # Discussion
//
// In some situations, such as when the device is using battery only, the
// system may throttle indexing to save power. You can implement this method
// to be notified of this situation so that you can respond by, for example,
// prioritizing the items to index.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableIndexDelegate/searchableIndexDidThrottle(_:)
func (c CSIndexExtensionRequestHandler) SearchableIndexDidThrottle(searchableIndex ICSSearchableIndex) {
	objc.Send[objc.ID](c.ID, objc.Sel("searchableIndexDidThrottle:"), searchableIndex)
}

// Tells the delegate to reindex all searchable data and clear all local state
// information.
//
// searchableIndex: The index in which to reindex the searchable data. The delegate or app
// extension should pass `searchableIndex` to
// [CSSearchableIndex.IndexSearchableItemsCompletionHandler].
//
// acknowledgementHandler: The handler to call after all client state has been saved. Note that if the
// app passes client state information in a batch (for example, by calling
// [CSSearchableIndex.EndIndexBatchWithClientStateCompletionHandler]), the
// acknowledgement handler can be called immediately.
//
// The delegate or app extension must call the acknowledgement handler after
// all client state information has been saved, so that the indexer can call
// this method again in case of a crash.
//
// # Discussion
//
// Typically, the index tells the delegate to reindex its searchable data and
// clear local state when the index has been lost. An app extension should not
// use the index passed in `searchableIndex` when a custom data protection
// class is needed.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableIndexDelegate/searchableIndex(_:reindexAllSearchableItemsWithAcknowledgementHandler:)
func (c CSIndexExtensionRequestHandler) SearchableIndexReindexAllSearchableItemsWithAcknowledgementHandler(searchableIndex ICSSearchableIndex, acknowledgementHandler VoidHandler) {
	_block1, _ := NewVoidBlock(acknowledgementHandler)
	objc.Send[objc.ID](c.ID, objc.Sel("searchableIndex:reindexAllSearchableItemsWithAcknowledgementHandler:"), searchableIndex, _block1)
}

// Tells the delegate to reindex the searchable items associated with the
// specified identifiers.
//
// searchableIndex: The index in which to reindex the specified searchable data. To update the
// state of the items, the delegate or app extension should call
// [CSSearchableIndex.IndexSearchableItemsCompletionHandler] passing in
// `searchableIndex`.
//
// identifiers: An array of identifiers that specify searchable items.
//
// acknowledgementHandler: The handler to call after all client state has been saved. Note that if the
// app passes client state information in a batch (for example, by calling
// [CSSearchableIndex.EndIndexBatchWithClientStateCompletionHandler]), the
// acknowledgement handler can be called immediately.
//
// The delegate or app extension must call the acknowledgement handler after
// all client state information has been saved, so that the indexer can call
// this method again in case of a crash.
//
// # Discussion
//
// An app extension should not use the index passed in `searchableIndex` when
// a custom data protection class is needed.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableIndexDelegate/searchableIndex(_:reindexSearchableItemsWithIdentifiers:acknowledgementHandler:)
func (c CSIndexExtensionRequestHandler) SearchableIndexReindexSearchableItemsWithIdentifiersAcknowledgementHandler(searchableIndex ICSSearchableIndex, identifiers []string, acknowledgementHandler VoidHandler) {
	_block2, _ := NewVoidBlock(acknowledgementHandler)
	objc.Send[objc.ID](c.ID, objc.Sel("searchableIndex:reindexSearchableItemsWithIdentifiers:acknowledgementHandler:"), searchableIndex, identifiers, _block2)
}

// Tells the delegate that the framework updated the list of searchable items.
//
// items: The items the framework updated.
//
// # Discussion
//
// The framework calls this method when it updates an item with specific
// attributes; see [CSSearchableItem.UpdateListenerOptions] for Apple
// Intelligence attributes.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableIndexDelegate/searchableItemsDidUpdate(_:)
//
// [CSSearchableItem.UpdateListenerOptions]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItem/UpdateListenerOptions-swift.struct
func (c CSIndexExtensionRequestHandler) SearchableItemsDidUpdate(items []CSSearchableItem) {
	objc.Send[objc.ID](c.ID, objc.Sel("searchableItemsDidUpdate:"), objectivec.IObjectSliceToNSArray(items))
}

// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableIndexDelegate/searchableItems(forIdentifiers:protectionClass:searchableItemsHandler:)
func (c CSIndexExtensionRequestHandler) SearchableItemsForIdentifiersProtectionClassSearchableItemsHandler(identifiers []string, protectionClass foundation.NSFileProtectionType, searchableItemsHandler CSSearchableItemArrayHandler) {
	_block2, _ := NewCSSearchableItemArrayBlock(searchableItemsHandler)
	objc.Send[objc.ID](c.ID, objc.Sel("searchableItemsForIdentifiers:protectionClass:searchableItemsHandler:"), identifiers, protectionClass, _block2)
}

// Requests that the delegate provide searchable items for the provided
// identifiers.
//
// identifiers: An array of strings that represent the identifiers.
//
// searchableItemsHandler: A method the framework calls that provides an array of [CSSearchableItem]
// objects.
//
// # Discussion
//
// Use this method to provide the framework with a list of identifiers to
// search for.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableIndexDelegate/searchableItems(forIdentifiers:searchableItemsHandler:)
func (c CSIndexExtensionRequestHandler) SearchableItemsForIdentifiersSearchableItemsHandler(identifiers []string, searchableItemsHandler CSSearchableItemArrayHandler) {
	_block1, _ := NewCSSearchableItemArrayBlock(searchableItemsHandler)
	objc.Send[objc.ID](c.ID, objc.Sel("searchableItemsForIdentifiers:searchableItemsHandler:"), identifiers, _block1)
}

// Protocol methods for CSSearchableIndexDelegate

// SearchableIndexReindexAllSearchableItemsWithAcknowledgementHandlerSync is a synchronous wrapper around [CSIndexExtensionRequestHandler.SearchableIndexReindexAllSearchableItemsWithAcknowledgementHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (c CSIndexExtensionRequestHandler) SearchableIndexReindexAllSearchableItemsWithAcknowledgementHandlerSync(ctx context.Context, searchableIndex ICSSearchableIndex) error {
	done := make(chan struct{}, 1)
	c.SearchableIndexReindexAllSearchableItemsWithAcknowledgementHandler(searchableIndex, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
