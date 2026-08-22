// Code generated from Apple documentation for CoreSpotlight. DO NOT EDIT.

package corespotlight

import (
	"context"
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
	"github.com/tmc/apple/uniformtypeidentifiers"
)

// The class instance for the [CSSearchableIndex] class.
var (
	_CSSearchableIndexClass     CSSearchableIndexClass
	_CSSearchableIndexClassOnce sync.Once
)

func getCSSearchableIndexClass() CSSearchableIndexClass {
	_CSSearchableIndexClassOnce.Do(func() {
		_CSSearchableIndexClass = CSSearchableIndexClass{class: objc.GetClass("CSSearchableIndex")}
	})
	return _CSSearchableIndexClass
}

// GetCSSearchableIndexClass returns the class object for CSSearchableIndex.
func GetCSSearchableIndexClass() CSSearchableIndexClass {
	return getCSSearchableIndexClass()
}

type CSSearchableIndexClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CSSearchableIndexClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CSSearchableIndexClass) Alloc() CSSearchableIndex {
	rv := objc.Send[CSSearchableIndex](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// An on-device index for your app’s searchable content.
//
// # Overview
//
// A [CSSearchableIndex] object manages an on-device index for your app’s
// searchable content. To make your app’s content searchable, create one or
// more [CSSearchableItem] objects for your content and add those items to the
// index. If your app defines [AppEntity] types, you can also index those
// types directly or associate them with your [CSSearchableItem] objects. When
// you execute a query, Core Spotlight searches your indexes for the requested
// information and returns the results to your code.
//
// Create custom [CSSearchableIndex] objects in your production code to store
// your app’s content, instead of using the default index. Custom indexes
// support data protection, which allows you to encrypt your data and protect
// it from unauthorized access. Custom indexes also support batch operations,
// which allow you to index large amounts of data more efficiently and with
// less risk. For example, you can add custom state information to each batch
// operation to make it easier to restart the indexing process if your app or
// app extension crashes. Use the default index only during testing or
// prototyping of your features.
//
// Modify a [CSSearchableIndex] object from only one thread or task at a time,
// and modify it only from your signed app or app extension. It’s a
// programming error to access a custom index from multiple threads
// simultaneously or from an unsigned bundle. When performing batch updates on
// an index, start each new batch operation only after calling the
// [CSSearchableIndex.EndIndexBatchWithClientStateCompletionHandler] or
// [CSSearchableIndex.EndIndexBatchWithExpectedClientStateNewClientStateCompletionHandler]
// method of the previous batch operation.
//
// # Creating an index
//
//   - [CSSearchableIndex.InitWithName]: Returns an on-device index with the specified name.
//   - [CSSearchableIndex.InitWithNameProtectionClass]: Returns an on-device index with the specified name and data protection class.
//
// # Responding to index-related changes
//
//   - [CSSearchableIndex.IndexDelegate]: The delegate object that can handle index-management tasks.
//   - [CSSearchableIndex.SetIndexDelegate]
//
// # Managing items in an index
//
//   - [CSSearchableIndex.IndexSearchableItemsCompletionHandler]: Adds or updates items in the index.
//   - [CSSearchableIndex.DeleteAllSearchableItemsWithCompletionHandler]: Deletes all searchable items from the index.
//   - [CSSearchableIndex.DeleteSearchableItemsWithDomainIdentifiersCompletionHandler]: Removes from the index all searchable items associated with the specified domain.
//   - [CSSearchableIndex.DeleteSearchableItemsWithIdentifiersCompletionHandler]: Removes from the index all items with the specified identifiers.
//
// # Batching index updates
//
//   - [CSSearchableIndex.BeginIndexBatch]: Begins a batch of updates to an index.
//   - [CSSearchableIndex.EndIndexBatchWithClientStateCompletionHandler]: Ends a batch of index updates and stores the specified state information.
//   - [CSSearchableIndex.EndIndexBatchWithExpectedClientStateNewClientStateCompletionHandler]: Ends a batch of index updates and stores the specified state information.
//   - [CSSearchableIndex.FetchLastClientStateWithCompletionHandler]: Fetches the app’s most recent client state information asynchronously.
//
// # Handling drag and drop content
//
//   - [CSSearchableIndex.FetchDataForBundleIdentifierItemIdentifierContentTypeCompletionHandler]: Fetches data from an external provider.
//
// # Getting the protection class
//
//   - [CSSearchableIndex.ProtectionClass]
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableIndex
//
// [AppEntity]: https://developer.apple.com/documentation/AppIntents/AppEntity
type CSSearchableIndex struct {
	objectivec.Object
}

// CSSearchableIndexFromID constructs a [CSSearchableIndex] from an objc.ID.
//
// An on-device index for your app’s searchable content.
func CSSearchableIndexFromID(id objc.ID) CSSearchableIndex {
	return CSSearchableIndex{objectivec.Object{ID: id}}
}

// NOTE: CSSearchableIndex adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CSSearchableIndex] class.
//
// # Creating an index
//
//   - [ICSSearchableIndex.InitWithName]: Returns an on-device index with the specified name.
//   - [ICSSearchableIndex.InitWithNameProtectionClass]: Returns an on-device index with the specified name and data protection class.
//
// # Responding to index-related changes
//
//   - [ICSSearchableIndex.IndexDelegate]: The delegate object that can handle index-management tasks.
//   - [ICSSearchableIndex.SetIndexDelegate]
//
// # Managing items in an index
//
//   - [ICSSearchableIndex.IndexSearchableItemsCompletionHandler]: Adds or updates items in the index.
//   - [ICSSearchableIndex.DeleteAllSearchableItemsWithCompletionHandler]: Deletes all searchable items from the index.
//   - [ICSSearchableIndex.DeleteSearchableItemsWithDomainIdentifiersCompletionHandler]: Removes from the index all searchable items associated with the specified domain.
//   - [ICSSearchableIndex.DeleteSearchableItemsWithIdentifiersCompletionHandler]: Removes from the index all items with the specified identifiers.
//
// # Batching index updates
//
//   - [ICSSearchableIndex.BeginIndexBatch]: Begins a batch of updates to an index.
//   - [ICSSearchableIndex.EndIndexBatchWithClientStateCompletionHandler]: Ends a batch of index updates and stores the specified state information.
//   - [ICSSearchableIndex.EndIndexBatchWithExpectedClientStateNewClientStateCompletionHandler]: Ends a batch of index updates and stores the specified state information.
//   - [ICSSearchableIndex.FetchLastClientStateWithCompletionHandler]: Fetches the app’s most recent client state information asynchronously.
//
// # Handling drag and drop content
//
//   - [ICSSearchableIndex.FetchDataForBundleIdentifierItemIdentifierContentTypeCompletionHandler]: Fetches data from an external provider.
//
// # Getting the protection class
//
//   - [ICSSearchableIndex.ProtectionClass]
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableIndex
type ICSSearchableIndex interface {
	objectivec.IObject

	// Topic: Creating an index

	// Returns an on-device index with the specified name.
	InitWithName(name string) CSSearchableIndex
	// Returns an on-device index with the specified name and data protection class.
	InitWithNameProtectionClass(name string, protectionClass foundation.NSFileProtectionType) CSSearchableIndex

	// Topic: Responding to index-related changes

	// The delegate object that can handle index-management tasks.
	IndexDelegate() CSSearchableIndexDelegate
	SetIndexDelegate(value CSSearchableIndexDelegate)

	// Topic: Managing items in an index

	// Adds or updates items in the index.
	IndexSearchableItemsCompletionHandler(items []CSSearchableItem, completionHandler ErrorHandler)
	// Deletes all searchable items from the index.
	DeleteAllSearchableItemsWithCompletionHandler(completionHandler ErrorHandler)
	// Removes from the index all searchable items associated with the specified domain.
	DeleteSearchableItemsWithDomainIdentifiersCompletionHandler(domainIdentifiers []string, completionHandler ErrorHandler)
	// Removes from the index all items with the specified identifiers.
	DeleteSearchableItemsWithIdentifiersCompletionHandler(identifiers []string, completionHandler ErrorHandler)

	// Topic: Batching index updates

	// Begins a batch of updates to an index.
	BeginIndexBatch()
	// Ends a batch of index updates and stores the specified state information.
	EndIndexBatchWithClientStateCompletionHandler(clientState foundation.NSData, completionHandler ErrorHandler)
	// Ends a batch of index updates and stores the specified state information.
	EndIndexBatchWithExpectedClientStateNewClientStateCompletionHandler(expectedClientState foundation.NSData, newClientState foundation.NSData, completionHandler ErrorHandler)
	// Fetches the app’s most recent client state information asynchronously.
	FetchLastClientStateWithCompletionHandler(completionHandler DataErrorHandler)

	// Topic: Handling drag and drop content

	// Fetches data from an external provider.
	FetchDataForBundleIdentifierItemIdentifierContentTypeCompletionHandler(bundleIdentifier string, itemIdentifier string, contentType uniformtypeidentifiers.UTType, completionHandler DataErrorHandler)

	// Topic: Getting the protection class

	ProtectionClass() foundation.NSFileProtectionType
}

// Init initializes the instance.
func (c CSSearchableIndex) Init() CSSearchableIndex {
	rv := objc.Send[CSSearchableIndex](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CSSearchableIndex) Autorelease() CSSearchableIndex {
	rv := objc.Send[CSSearchableIndex](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCSSearchableIndex creates a new CSSearchableIndex instance.
func NewCSSearchableIndex() CSSearchableIndex {
	class := getCSSearchableIndexClass()
	rv := objc.Send[CSSearchableIndex](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Returns an on-device index with the specified name.
//
// name: A name that pertains to your custom organization.
//
// # Return Value
//
// An on-device index.
//
// # Discussion
//
// If you want to use batching or you want to index items in a specific
// protection class, you need to use your own index (you can’t perform batch
// updates on the default index).
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableIndex/init(name:)
func NewCSSearchableIndexWithName(name string) CSSearchableIndex {
	instance := getCSSearchableIndexClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithName:"), objc.String(name))
	return CSSearchableIndexFromID(rv)
}

// Returns an on-device index with the specified name and data protection
// class.
//
// name: A name that pertains to your custom organization.
//
// protectionClass: The file protection class. Acceptable values are [none], [complete],
// [completeUnlessOpen], or [completeUntilFirstUserAuthentication].
//
// # Return Value
//
// An index that can handle items within the specified protection class.
//
// # Discussion
//
// Use this method to specify a protection class for an index. You can specify
// a default protection class for index items in the entitlements for your
// app.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableIndex/init(name:protectionClass:)
//
// [completeUnlessOpen]: https://developer.apple.com/documentation/Foundation/FileProtectionType/completeUnlessOpen
// [completeUntilFirstUserAuthentication]: https://developer.apple.com/documentation/Foundation/FileProtectionType/completeUntilFirstUserAuthentication
// [complete]: https://developer.apple.com/documentation/Foundation/FileProtectionType/complete
// [none]: https://developer.apple.com/documentation/Foundation/FileProtectionType/none
func NewCSSearchableIndexWithNameProtectionClass(name string, protectionClass foundation.NSFileProtectionType) CSSearchableIndex {
	instance := getCSSearchableIndexClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithName:protectionClass:"), objc.String(name), protectionClass)
	return CSSearchableIndexFromID(rv)
}

// Returns an on-device index with the specified name.
//
// name: A name that pertains to your custom organization.
//
// # Return Value
//
// An on-device index.
//
// # Discussion
//
// If you want to use batching or you want to index items in a specific
// protection class, you need to use your own index (you can’t perform batch
// updates on the default index).
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableIndex/init(name:)
func (c CSSearchableIndex) InitWithName(name string) CSSearchableIndex {
	rv := objc.Send[CSSearchableIndex](c.ID, objc.Sel("initWithName:"), objc.String(name))
	return rv
}

// Returns an on-device index with the specified name and data protection
// class.
//
// name: A name that pertains to your custom organization.
//
// protectionClass: The file protection class. Acceptable values are [none], [complete],
// [completeUnlessOpen], or [completeUntilFirstUserAuthentication].
//
// # Return Value
//
// An index that can handle items within the specified protection class.
//
// # Discussion
//
// Use this method to specify a protection class for an index. You can specify
// a default protection class for index items in the entitlements for your
// app.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableIndex/init(name:protectionClass:)
//
// [completeUnlessOpen]: https://developer.apple.com/documentation/Foundation/FileProtectionType/completeUnlessOpen
// [completeUntilFirstUserAuthentication]: https://developer.apple.com/documentation/Foundation/FileProtectionType/completeUntilFirstUserAuthentication
// [complete]: https://developer.apple.com/documentation/Foundation/FileProtectionType/complete
// [none]: https://developer.apple.com/documentation/Foundation/FileProtectionType/none
func (c CSSearchableIndex) InitWithNameProtectionClass(name string, protectionClass foundation.NSFileProtectionType) CSSearchableIndex {
	rv := objc.Send[CSSearchableIndex](c.ID, objc.Sel("initWithName:protectionClass:"), objc.String(name), protectionClass)
	return rv
}

// Adds or updates items in the index.
//
// items: An array of searchable items to add or update.
//
// completionHandler: The block that’s called when the data has been journaled by the index,
// which means that the index makes a note that it has to perform this
// operation. If the completion handler returns an error, it means that the
// data wasn’t journaled correctly and the client should retry the request.
//
// The block receives the following parameter:
//
// error: If an error occurred, this parameter holds an error object that
// explains the error. Otherwise, the value of this parameter is `nil`.
//
// # Discussion
//
// The
// [SearchableIndexReindexSearchableItemsWithIdentifiersAcknowledgementHandler]
// protocol method is called in the case that the journaling completed
// successfully but the data was not able to be indexed for some reason.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableIndex/indexSearchableItems(_:completionHandler:)
func (c CSSearchableIndex) IndexSearchableItemsCompletionHandler(items []CSSearchableItem, completionHandler ErrorHandler) {
	_block1, _ := NewErrorBlock(completionHandler)
	objc.Send[objc.ID](c.ID, objc.Sel("indexSearchableItems:completionHandler:"), items, _block1)
}

// Deletes all searchable items from the index.
//
// completionHandler: The block that’s called when the request has been journaled by the index
// (“journaled” means that the index makes a note that it has to perform
// this operation). Note that the request may not have completed.
//
// The block receives the following parameter:
//
// error: If an error occurred, this parameter holds an error object that
// explains the error. Otherwise, the value of this parameter is `nil`.
//
// # Discussion
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableIndex/deleteAllSearchableItems(completionHandler:)
func (c CSSearchableIndex) DeleteAllSearchableItemsWithCompletionHandler(completionHandler ErrorHandler) {
	_block0, _ := NewErrorBlock(completionHandler)
	objc.Send[objc.ID](c.ID, objc.Sel("deleteAllSearchableItemsWithCompletionHandler:"), _block0)
}

// Removes from the index all searchable items associated with the specified
// domain.
//
// domainIdentifiers: The domain identifier that describes the group of items to delete. To learn
// more about domain identifiers, see [CSSearchableItem.DomainIdentifier].
//
// completionHandler: The block that’s called when the request has been journaled by the index
// (“journaled” means that the index makes a note that it has to perform
// this operation). Note that the request may not have completed.
//
// The block receives the following parameter:
//
// error: If an error occurred, this parameter holds an error object that
// explains the error. Otherwise, the value of this parameter is `nil`.
//
// # Discussion
//
// Use this method to delete groups of items. Note that the delete operation
// is recursive. For example, if domain identifiers are of the form `.`,
// calling this method and specifying “ deletes items with the specified
// account and any mailbox.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableIndex/deleteSearchableItems(withDomainIdentifiers:completionHandler:)
func (c CSSearchableIndex) DeleteSearchableItemsWithDomainIdentifiersCompletionHandler(domainIdentifiers []string, completionHandler ErrorHandler) {
	_block1, _ := NewErrorBlock(completionHandler)
	objc.Send[objc.ID](c.ID, objc.Sel("deleteSearchableItemsWithDomainIdentifiers:completionHandler:"), domainIdentifiers, _block1)
}

// Removes from the index all items with the specified identifiers.
//
// identifiers: An array of identifiers that specify the items to delete.
//
// completionHandler: The block that’s called when the data has been journaled by the index,
// which means that the index makes a note that it has to perform this
// operation. If the completion handler returns an error, it means that the
// data wasn’t journaled correctly and the client should retry the request.
//
// The block receives the following parameter:
//
// error: If an error occurred, this parameter holds an error object that
// explains the error. Otherwise, the value of this parameter is `nil`.
//
// # Discussion
//
// The
// [SearchableIndexReindexSearchableItemsWithIdentifiersAcknowledgementHandler]
// protocol method is called in the case that the journaling completed
// successfully, but the data was not able to be indexed for some reason.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableIndex/deleteSearchableItems(withIdentifiers:completionHandler:)
func (c CSSearchableIndex) DeleteSearchableItemsWithIdentifiersCompletionHandler(identifiers []string, completionHandler ErrorHandler) {
	_block1, _ := NewErrorBlock(completionHandler)
	objc.Send[objc.ID](c.ID, objc.Sel("deleteSearchableItemsWithIdentifiers:completionHandler:"), identifiers, _block1)
}

// Begins a batch of updates to an index.
//
// # Discussion
//
// Don’t call this method again before
// [CSSearchableIndex.EndIndexBatchWithClientStateCompletionHandler] has
// returned. (You can call it again before the completion handler passed to
// [CSSearchableIndex.EndIndexBatchWithClientStateCompletionHandler] has been
// called.)
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableIndex/beginBatch()
func (c CSSearchableIndex) BeginIndexBatch() {
	objc.Send[objc.ID](c.ID, objc.Sel("beginIndexBatch"))
}

// Ends a batch of index updates and stores the specified state information.
//
// clientState: Up to 250 bytes of information that can help you recover from a crash and
// resume indexing.
//
// completionHandler: The block that’s called after the client state has been stored.
//
// The block receives the following parameter:
//
// error: If an error occurred, this parameter holds an error object that
// explains the error. Otherwise, the value of this parameter is `nil`.
//
// # Discussion
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableIndex/endBatch(withClientState:completionHandler:)
func (c CSSearchableIndex) EndIndexBatchWithClientStateCompletionHandler(clientState foundation.NSData, completionHandler ErrorHandler) {
	_block1, _ := NewErrorBlock(completionHandler)
	objc.Send[objc.ID](c.ID, objc.Sel("endIndexBatchWithClientState:completionHandler:"), clientState, _block1)
}

// Ends a batch of index updates and stores the specified state information.
//
// expectedClientState: The client state data from the previous batch.
//
// newClientState: Up to 250 bytes of app-specific data that can help you recover from a crash
// and resume indexing.
//
// completionHandler: The block to call with the results. The block receives the following
// parameter:
//
// error: If an error occurred, this parameter holds an error object that
// explains the error. Otherwise, the value of this parameter is `nil`.
//
// # Discussion
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableIndex/endIndexBatch(expectedClientState:newClientState:completionHandler:)
func (c CSSearchableIndex) EndIndexBatchWithExpectedClientStateNewClientStateCompletionHandler(expectedClientState foundation.NSData, newClientState foundation.NSData, completionHandler ErrorHandler) {
	_block2, _ := NewErrorBlock(completionHandler)
	objc.Send[objc.ID](c.ID, objc.Sel("endIndexBatchWithExpectedClientState:newClientState:completionHandler:"), expectedClientState, newClientState, _block2)
}

// Fetches the app’s most recent client state information asynchronously.
//
// completionHandler: The block to call when the request has been journaled by the index, which
// means that the index makes a note that it has to perform this operation.
// Note that the request may not have completed.
//
// The block receives the following parameter:
//
// error: If an error occurred, this parameter holds an error object that
// explains the error. Otherwise, the value of this parameter is `nil`.
//
// # Discussion
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableIndex/fetchLastClientState(completionHandler:)
func (c CSSearchableIndex) FetchLastClientStateWithCompletionHandler(completionHandler DataErrorHandler) {
	_block0, _ := NewDataErrorBlock(completionHandler)
	objc.Send[objc.ID](c.ID, objc.Sel("fetchLastClientStateWithCompletionHandler:"), _block0)
}

// Fetches data from an external provider.
//
// bundleIdentifier: The bundle identifier of the app to search.
//
// itemIdentifier: The app-specific identifier of the item you want.
//
// contentType: The type of data to fetch.
//
// completionHandler: The block to execute with the results. The block has no return value and
// takes the following parameters:
//
// data: The data for the specified item, if successful. error: An error
// object, or `nil` if the method retrieved the data successfully.
//
// # Discussion
//
// Clients with the appropriate entitlements can use this method to fetch data
// from an external app such as Mail.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableIndex/fetchData(forBundleIdentifier:itemIdentifier:contentType:completionHandler:)
func (c CSSearchableIndex) FetchDataForBundleIdentifierItemIdentifierContentTypeCompletionHandler(bundleIdentifier string, itemIdentifier string, contentType uniformtypeidentifiers.UTType, completionHandler DataErrorHandler) {
	_block3, _ := NewDataErrorBlock(completionHandler)
	objc.Send[objc.ID](c.ID, objc.Sel("fetchDataForBundleIdentifier:itemIdentifier:contentType:completionHandler:"), objc.String(bundleIdentifier), objc.String(itemIdentifier), contentType, _block3)
}

// Returns the default on-device index.
//
// # Return Value
//
// The default on-device index.
//
// # Discussion
//
// If you want to use batching or you want to index items in a specific
// protection class, you need to use your own index (you can’t perform batch
// updates on the default index).
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableIndex/default()
func (_CSSearchableIndexClass CSSearchableIndexClass) DefaultSearchableIndex() CSSearchableIndex {
	rv := objc.Send[objc.ID](objc.ID(_CSSearchableIndexClass.class), objc.Sel("defaultSearchableIndex"))
	return CSSearchableIndexFromID(rv)
}

// Returns a Boolean value that indicates whether indexing is available on the
// current device.
//
// # Return Value
//
// `true` if indexing is available on the current device, or `false` if it
// isn’t.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableIndex/isIndexingAvailable()
func (_CSSearchableIndexClass CSSearchableIndexClass) IsIndexingAvailable() bool {
	rv := objc.Send[bool](objc.ID(_CSSearchableIndexClass.class), objc.Sel("isIndexingAvailable"))
	return rv
}

// The delegate object that can handle index-management tasks.
//
// # Discussion
//
// The delegate should conform to the [CSSearchableIndexDelegate] protocol.
// Set this property to handle communication with the index and perform
// index-management tasks for your app. In particular, long-running apps
// should set a delegate so that the index can be updated while the app is in
// the background. Alternatively, you can create an extension with a request
// handler that conforms to the [CSSearchableIndexDelegate] protocol and let
// the extension perform index updates when your app isn’t running.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableIndex/indexDelegate
func (c CSSearchableIndex) IndexDelegate() CSSearchableIndexDelegate {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("indexDelegate"))
	return CSSearchableIndexDelegateObjectFromID(rv)
}
func (c CSSearchableIndex) SetIndexDelegate(value CSSearchableIndexDelegate) {
	objc.Send[struct{}](c.ID, objc.Sel("setIndexDelegate:"), value)
}

// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableIndex/protectionClass
func (c CSSearchableIndex) ProtectionClass() foundation.NSFileProtectionType {
	rv := objc.Send[foundation.NSFileProtectionType](c.ID, objc.Sel("protectionClass"))
	return foundation.NSFileProtectionType(rv)
}

// DeleteAllSearchableItems is a synchronous wrapper around [CSSearchableIndex.DeleteAllSearchableItemsWithCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (c CSSearchableIndex) DeleteAllSearchableItems(ctx context.Context) error {
	done := make(chan error, 1)
	c.DeleteAllSearchableItemsWithCompletionHandler(func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

// EndIndexBatchWithClientState is a synchronous wrapper around [CSSearchableIndex.EndIndexBatchWithClientStateCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (c CSSearchableIndex) EndIndexBatchWithClientState(ctx context.Context, clientState foundation.NSData) error {
	done := make(chan error, 1)
	c.EndIndexBatchWithClientStateCompletionHandler(clientState, func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

// EndIndexBatchWithExpectedClientStateNewClientState is a synchronous wrapper around [CSSearchableIndex.EndIndexBatchWithExpectedClientStateNewClientStateCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (c CSSearchableIndex) EndIndexBatchWithExpectedClientStateNewClientState(ctx context.Context, expectedClientState foundation.NSData, newClientState foundation.NSData) error {
	done := make(chan error, 1)
	c.EndIndexBatchWithExpectedClientStateNewClientStateCompletionHandler(expectedClientState, newClientState, func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

// FetchLastClientState is a synchronous wrapper around [CSSearchableIndex.FetchLastClientStateWithCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (c CSSearchableIndex) FetchLastClientState(ctx context.Context) (*foundation.NSData, error) {
	type result struct {
		val *foundation.NSData
		err error
	}
	done := make(chan result, 1)
	c.FetchLastClientStateWithCompletionHandler(func(val *foundation.NSData, err error) {
		done <- result{val, err}
	})
	select {
	case r := <-done:
		return r.val, r.err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

// FetchDataForBundleIdentifierItemIdentifierContentType is a synchronous wrapper around [CSSearchableIndex.FetchDataForBundleIdentifierItemIdentifierContentTypeCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (c CSSearchableIndex) FetchDataForBundleIdentifierItemIdentifierContentType(ctx context.Context, bundleIdentifier string, itemIdentifier string, contentType uniformtypeidentifiers.UTType) (*foundation.NSData, error) {
	type result struct {
		val *foundation.NSData
		err error
	}
	done := make(chan result, 1)
	c.FetchDataForBundleIdentifierItemIdentifierContentTypeCompletionHandler(bundleIdentifier, itemIdentifier, contentType, func(val *foundation.NSData, err error) {
		done <- result{val, err}
	})
	select {
	case r := <-done:
		return r.val, r.err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}
