// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"context"
	"sync"

	"github.com/tmc/apple/corespotlight"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSCoreDataCoreSpotlightDelegate] class.
var (
	_NSCoreDataCoreSpotlightDelegateClass     NSCoreDataCoreSpotlightDelegateClass
	_NSCoreDataCoreSpotlightDelegateClassOnce sync.Once
)

func getNSCoreDataCoreSpotlightDelegateClass() NSCoreDataCoreSpotlightDelegateClass {
	_NSCoreDataCoreSpotlightDelegateClassOnce.Do(func() {
		_NSCoreDataCoreSpotlightDelegateClass = NSCoreDataCoreSpotlightDelegateClass{class: objc.GetClass("NSCoreDataCoreSpotlightDelegate")}
	})
	return _NSCoreDataCoreSpotlightDelegateClass
}

// GetNSCoreDataCoreSpotlightDelegateClass returns the class object for NSCoreDataCoreSpotlightDelegate.
func GetNSCoreDataCoreSpotlightDelegateClass() NSCoreDataCoreSpotlightDelegateClass {
	return getNSCoreDataCoreSpotlightDelegateClass()
}

type NSCoreDataCoreSpotlightDelegateClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSCoreDataCoreSpotlightDelegateClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSCoreDataCoreSpotlightDelegateClass) Alloc() NSCoreDataCoreSpotlightDelegate {
	rv := objc.Send[NSCoreDataCoreSpotlightDelegate](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A set of methods that enable integration with Core Spotlight.
//
// # Overview
//
// # Creating a Core Spotlight Delegate
//
//   - [NSCoreDataCoreSpotlightDelegate.InitForStoreWithDescriptionCoordinator]: Creates a Core Spotlight delegate with the specified store description and coordinator.
//
// # Configuring the Index
//
//   - [NSCoreDataCoreSpotlightDelegate.IsIndexingEnabled]: A Boolean value that indicates whether Core Data is currently updating the Core Spotlight index with the persistent store’s entities.
//   - [NSCoreDataCoreSpotlightDelegate.DomainIdentifier]: Returns the domain identifier.
//   - [NSCoreDataCoreSpotlightDelegate.IndexName]: Returns the index’s name.
//
// # Managing the Index
//
//   - [NSCoreDataCoreSpotlightDelegate.AttributeSetForObject]: Returns the searchable attributes for the specified managed object.
//   - [NSCoreDataCoreSpotlightDelegate.DeleteSpotlightIndexWithCompletionHandler]: Deletes all searchable items from the configured index.
//   - [NSCoreDataCoreSpotlightDelegate.StartSpotlightIndexing]: Starts the indexing of the store’s entities.
//   - [NSCoreDataCoreSpotlightDelegate.StopSpotlightIndexing]: Stops the indexing of the store’s entities.
//
// # Updating the Index
//
//   - [NSCoreDataCoreSpotlightDelegate.SearchableIndexReindexAllSearchableItemsWithAcknowledgementHandler]: Reindexes all searchable items and clears any local state.
//   - [NSCoreDataCoreSpotlightDelegate.SearchableIndexReindexSearchableItemsWithIdentifiersAcknowledgementHandler]: Reindexes the searchable items for the specified identifiers.
//
// See: https://developer.apple.com/documentation/CoreData/NSCoreDataCoreSpotlightDelegate
type NSCoreDataCoreSpotlightDelegate struct {
	objectivec.Object
}

// NSCoreDataCoreSpotlightDelegateFromID constructs a [NSCoreDataCoreSpotlightDelegate] from an objc.ID.
//
// A set of methods that enable integration with Core Spotlight.
func NSCoreDataCoreSpotlightDelegateFromID(id objc.ID) NSCoreDataCoreSpotlightDelegate {
	return NSCoreDataCoreSpotlightDelegate{objectivec.Object{ID: id}}
}

// NOTE: NSCoreDataCoreSpotlightDelegate adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSCoreDataCoreSpotlightDelegate] class.
//
// # Creating a Core Spotlight Delegate
//
//   - [INSCoreDataCoreSpotlightDelegate.InitForStoreWithDescriptionCoordinator]: Creates a Core Spotlight delegate with the specified store description and coordinator.
//
// # Configuring the Index
//
//   - [INSCoreDataCoreSpotlightDelegate.IsIndexingEnabled]: A Boolean value that indicates whether Core Data is currently updating the Core Spotlight index with the persistent store’s entities.
//   - [INSCoreDataCoreSpotlightDelegate.DomainIdentifier]: Returns the domain identifier.
//   - [INSCoreDataCoreSpotlightDelegate.IndexName]: Returns the index’s name.
//
// # Managing the Index
//
//   - [INSCoreDataCoreSpotlightDelegate.AttributeSetForObject]: Returns the searchable attributes for the specified managed object.
//   - [INSCoreDataCoreSpotlightDelegate.DeleteSpotlightIndexWithCompletionHandler]: Deletes all searchable items from the configured index.
//   - [INSCoreDataCoreSpotlightDelegate.StartSpotlightIndexing]: Starts the indexing of the store’s entities.
//   - [INSCoreDataCoreSpotlightDelegate.StopSpotlightIndexing]: Stops the indexing of the store’s entities.
//
// # Updating the Index
//
//   - [INSCoreDataCoreSpotlightDelegate.SearchableIndexReindexAllSearchableItemsWithAcknowledgementHandler]: Reindexes all searchable items and clears any local state.
//   - [INSCoreDataCoreSpotlightDelegate.SearchableIndexReindexSearchableItemsWithIdentifiersAcknowledgementHandler]: Reindexes the searchable items for the specified identifiers.
//
// See: https://developer.apple.com/documentation/CoreData/NSCoreDataCoreSpotlightDelegate
type INSCoreDataCoreSpotlightDelegate interface {
	objectivec.IObject

	// Topic: Creating a Core Spotlight Delegate

	// Creates a Core Spotlight delegate with the specified store description and coordinator.
	InitForStoreWithDescriptionCoordinator(description INSPersistentStoreDescription, psc INSPersistentStoreCoordinator) NSCoreDataCoreSpotlightDelegate

	// Topic: Configuring the Index

	// A Boolean value that indicates whether Core Data is currently updating the Core Spotlight index with the persistent store’s entities.
	IsIndexingEnabled() bool
	// Returns the domain identifier.
	DomainIdentifier() string
	// Returns the index’s name.
	IndexName() string

	// Topic: Managing the Index

	// Returns the searchable attributes for the specified managed object.
	AttributeSetForObject(object INSManagedObject) corespotlight.CSSearchableItemAttributeSet
	// Deletes all searchable items from the configured index.
	DeleteSpotlightIndexWithCompletionHandler(completionHandler ErrorHandler)
	// Starts the indexing of the store’s entities.
	StartSpotlightIndexing()
	// Stops the indexing of the store’s entities.
	StopSpotlightIndexing()

	// Topic: Updating the Index

	// Reindexes all searchable items and clears any local state.
	SearchableIndexReindexAllSearchableItemsWithAcknowledgementHandler(searchableIndex *corespotlight.CSSearchableIndex, acknowledgementHandler VoidHandler)
	// Reindexes the searchable items for the specified identifiers.
	SearchableIndexReindexSearchableItemsWithIdentifiersAcknowledgementHandler(searchableIndex *corespotlight.CSSearchableIndex, identifiers []string, acknowledgementHandler VoidHandler)
}

// Init initializes the instance.
func (c NSCoreDataCoreSpotlightDelegate) Init() NSCoreDataCoreSpotlightDelegate {
	rv := objc.Send[NSCoreDataCoreSpotlightDelegate](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c NSCoreDataCoreSpotlightDelegate) Autorelease() NSCoreDataCoreSpotlightDelegate {
	rv := objc.Send[NSCoreDataCoreSpotlightDelegate](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSCoreDataCoreSpotlightDelegate creates a new NSCoreDataCoreSpotlightDelegate instance.
func NewNSCoreDataCoreSpotlightDelegate() NSCoreDataCoreSpotlightDelegate {
	class := getNSCoreDataCoreSpotlightDelegateClass()
	rv := objc.Send[NSCoreDataCoreSpotlightDelegate](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a Core Spotlight delegate with the specified store description and
// coordinator.
//
// description: An object that describes the persistent store that contains the entities to
// index.
//
// psc: The persistent store coordinator, which you initialize with the managed
// object model that contains the definitions of the entities to index.
//
// # Discussion
//
// After you initialize a Core Spotlight delegate, call the
// [NSCoreDataCoreSpotlightDelegate.StartSpotlightIndexing] to begin indexing
// your store’s contents.
//
// See: https://developer.apple.com/documentation/CoreData/NSCoreDataCoreSpotlightDelegate/init(forStoreWith:coordinator:)
func NewCoreDataCoreSpotlightDelegateForStoreWithDescriptionCoordinator(description INSPersistentStoreDescription, psc INSPersistentStoreCoordinator) NSCoreDataCoreSpotlightDelegate {
	instance := getNSCoreDataCoreSpotlightDelegateClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initForStoreWithDescription:coordinator:"), description, psc)
	return NSCoreDataCoreSpotlightDelegateFromID(rv)
}

// Creates a Core Spotlight delegate with the specified store description and
// coordinator.
//
// description: An object that describes the persistent store that contains the entities to
// index.
//
// psc: The persistent store coordinator, which you initialize with the managed
// object model that contains the definitions of the entities to index.
//
// # Discussion
//
// After you initialize a Core Spotlight delegate, call the
// [NSCoreDataCoreSpotlightDelegate.StartSpotlightIndexing] to begin indexing
// your store’s contents.
//
// See: https://developer.apple.com/documentation/CoreData/NSCoreDataCoreSpotlightDelegate/init(forStoreWith:coordinator:)
func (c NSCoreDataCoreSpotlightDelegate) InitForStoreWithDescriptionCoordinator(description INSPersistentStoreDescription, psc INSPersistentStoreCoordinator) NSCoreDataCoreSpotlightDelegate {
	rv := objc.Send[NSCoreDataCoreSpotlightDelegate](c.ID, objc.Sel("initForStoreWithDescription:coordinator:"), description, psc)
	return rv
}

// Returns the domain identifier.
//
// # Discussion
//
// The default value is the persistent store’s identifier.
//
// See: https://developer.apple.com/documentation/CoreData/NSCoreDataCoreSpotlightDelegate/domainIdentifier()
func (c NSCoreDataCoreSpotlightDelegate) DomainIdentifier() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("domainIdentifier"))
	return foundation.NSStringFromID(rv).String()
}

// Returns the index’s name.
//
// # Discussion
//
// The default value is `nil`.
//
// See: https://developer.apple.com/documentation/CoreData/NSCoreDataCoreSpotlightDelegate/indexName()
func (c NSCoreDataCoreSpotlightDelegate) IndexName() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("indexName"))
	return foundation.NSStringFromID(rv).String()
}

// Returns the searchable attributes for the specified managed object.
//
// object: The managed object to index.
//
// # Return Value
//
// An instance of [CSSearchableItemAttributeSet] that provides the searchable
// item’s attributes.
//
// # Discussion
//
// To prevent Core Spotlight from indexing a specific managed object, override
// this method and return `nil` for that object.
//
// See: https://developer.apple.com/documentation/CoreData/NSCoreDataCoreSpotlightDelegate/attributeSet(for:)
//
// [CSSearchableItemAttributeSet]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet
func (c NSCoreDataCoreSpotlightDelegate) AttributeSetForObject(object INSManagedObject) corespotlight.CSSearchableItemAttributeSet {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("attributeSetForObject:"), object)
	return corespotlight.CSSearchableItemAttributeSetFromID(rv)
}

// Deletes all searchable items from the configured index.
//
// # Discussion
//
// The closure returns no value and takes only a single parameter, which is an
// error object that contains information about issues preventing the deletion
// of searchable items, or `nil` if Core Spotlight successfully deletes all
// searchable items.
//
// Depending on the cause of the issue, an error can originate from Core Data
// or from Core Spotlight. Make sure your app can handle both scenarios.
//
// See: https://developer.apple.com/documentation/CoreData/NSCoreDataCoreSpotlightDelegate/deleteSpotlightIndex(completionHandler:)
func (c NSCoreDataCoreSpotlightDelegate) DeleteSpotlightIndexWithCompletionHandler(completionHandler ErrorHandler) {
	_block0, _ := NewErrorBlock(completionHandler)
	objc.Send[objc.ID](c.ID, objc.Sel("deleteSpotlightIndexWithCompletionHandler:"), _block0)
}

// Starts the indexing of the store’s entities.
//
// # Discussion
//
// After you call this method, the delegate posts a notification whenever the
// index changes. The type of notification is [indexDidUpdateNotification],
// and its `userInfo` dictionary contains the keys [NSStoreUUIDKey] and
// [NSPersistentHistoryTokenKey].
//
// See: https://developer.apple.com/documentation/CoreData/NSCoreDataCoreSpotlightDelegate/startSpotlightIndexing()
//
// [NSPersistentHistoryTokenKey]: https://developer.apple.com/documentation/CoreData/NSPersistentHistoryTokenKey
// [NSStoreUUIDKey]: https://developer.apple.com/documentation/CoreData/NSStoreUUIDKey
// [indexDidUpdateNotification]: https://developer.apple.com/documentation/CoreData/NSCoreDataCoreSpotlightDelegate/indexDidUpdateNotification
func (c NSCoreDataCoreSpotlightDelegate) StartSpotlightIndexing() {
	objc.Send[objc.ID](c.ID, objc.Sel("startSpotlightIndexing"))
}

// Stops the indexing of the store’s entities.
//
// # Discussion
//
// After you call this method, the delegate no longer posts notifications
// about index changes.
//
// See: https://developer.apple.com/documentation/CoreData/NSCoreDataCoreSpotlightDelegate/stopSpotlightIndexing()
func (c NSCoreDataCoreSpotlightDelegate) StopSpotlightIndexing() {
	objc.Send[objc.ID](c.ID, objc.Sel("stopSpotlightIndexing"))
}

// Reindexes all searchable items and clears any local state.
//
// searchableIndex: The index that requires reindexing.
//
// acknowledgementHandler: The handler to call when you finish saving client state information.
//
// # Discussion
//
// For more information, see
// [searchableIndex(_:reindexAllSearchableItemsWithAcknowledgementHandler:)].
//
// See: https://developer.apple.com/documentation/CoreData/NSCoreDataCoreSpotlightDelegate/searchableIndex(_:reindexAllSearchableItemsWithAcknowledgementHandler:)
//
// [searchableIndex(_:reindexAllSearchableItemsWithAcknowledgementHandler:)]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableIndexDelegate/searchableIndex(_:reindexAllSearchableItemsWithAcknowledgementHandler:)
func (c NSCoreDataCoreSpotlightDelegate) SearchableIndexReindexAllSearchableItemsWithAcknowledgementHandler(searchableIndex *corespotlight.CSSearchableIndex, acknowledgementHandler VoidHandler) {
	_block1, _ := NewVoidBlock(acknowledgementHandler)
	objc.Send[objc.ID](c.ID, objc.Sel("searchableIndex:reindexAllSearchableItemsWithAcknowledgementHandler:"), searchableIndex, _block1)
}

// Reindexes the searchable items for the specified identifiers.
//
// searchableIndex: The index that contains the items that require reindexing.
//
// identifiers: An array of strings that identify the searchable items.
//
// acknowledgementHandler: The handler to call when you finish saving client state information.
//
// # Discussion
//
// For more information, see
// [NSCoreDataCoreSpotlightDelegate.SearchableIndexReindexSearchableItemsWithIdentifiersAcknowledgementHandler].
//
// See: https://developer.apple.com/documentation/CoreData/NSCoreDataCoreSpotlightDelegate/searchableIndex(_:reindexSearchableItemsWithIdentifiers:acknowledgementHandler:)
func (c NSCoreDataCoreSpotlightDelegate) SearchableIndexReindexSearchableItemsWithIdentifiersAcknowledgementHandler(searchableIndex *corespotlight.CSSearchableIndex, identifiers []string, acknowledgementHandler VoidHandler) {
	_block2, _ := NewVoidBlock(acknowledgementHandler)
	objc.Send[objc.ID](c.ID, objc.Sel("searchableIndex:reindexSearchableItemsWithIdentifiers:acknowledgementHandler:"), searchableIndex, identifiers, _block2)
}

// A Boolean value that indicates whether Core Data is currently updating the
// Core Spotlight index with the persistent store’s entities.
//
// See: https://developer.apple.com/documentation/CoreData/NSCoreDataCoreSpotlightDelegate/isIndexingEnabled
func (c NSCoreDataCoreSpotlightDelegate) IsIndexingEnabled() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isIndexingEnabled"))
	return rv
}

// DeleteSpotlightIndex is a synchronous wrapper around [NSCoreDataCoreSpotlightDelegate.DeleteSpotlightIndexWithCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (c NSCoreDataCoreSpotlightDelegate) DeleteSpotlightIndex(ctx context.Context) error {
	done := make(chan error, 1)
	c.DeleteSpotlightIndexWithCompletionHandler(func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

// SearchableIndexReindexAllSearchableItemsWithAcknowledgementHandlerSync is a synchronous wrapper around [NSCoreDataCoreSpotlightDelegate.SearchableIndexReindexAllSearchableItemsWithAcknowledgementHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (c NSCoreDataCoreSpotlightDelegate) SearchableIndexReindexAllSearchableItemsWithAcknowledgementHandlerSync(ctx context.Context, searchableIndex *corespotlight.CSSearchableIndex) error {
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
