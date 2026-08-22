// Code generated from Apple documentation. DO NOT EDIT.

package coredata

import (
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// AsynchronousFetchResultHandler is the signature for a completion handler block.
type AsynchronousFetchResultHandler = func(*NSAsynchronousFetchResult)

// NewAsynchronousFetchResultBlock wraps a Go [AsynchronousFetchResultHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewAsynchronousFetchResultBlock(handler AsynchronousFetchResultHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID) {
		var result *NSAsynchronousFetchResult
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := NSAsynchronousFetchResultFromID(resultID)
			result = &v
		}
		handler(result)
	})
	return objc.ID(block), func() { block.Release() }
}

// BoolManagedObjectHandler handles A closure that inserts data into the managed entity.
//
// Used by:
//   - [NSBatchInsertRequest.InitWithEntityManagedObjectHandler]
//   - [NSBatchInsertRequest.InitWithEntityNameManagedObjectHandler]
type BoolManagedObjectHandler = func(*NSManagedObject) bool

// NewBoolManagedObjectBlock wraps a Go [BoolManagedObjectHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [NSBatchInsertRequest.InitWithEntityManagedObjectHandler]
//   - [NSBatchInsertRequest.InitWithEntityNameManagedObjectHandler]
func NewBoolManagedObjectBlock(handler BoolManagedObjectHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID) bool {
		var result *NSManagedObject
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := NSManagedObjectFromID(resultID)
			result = &v
		}
		return handler(result)
	})
	return objc.ID(block), func() { block.Release() }
}

// BoolStringidMutableDictionaryHandler handles A closure that provides a dictionary that represents an object to insert.
//
// Used by:
//   - [NSBatchInsertRequest.InitWithEntityDictionaryHandler]
//   - [NSBatchInsertRequest.InitWithEntityNameDictionaryHandler]
type BoolStringidMutableDictionaryHandler = func(*foundation.INSDictionary) bool

// ErrorHandler is the signature for a completion handler block.
//
// Used by:
//   - [NSCoreDataCoreSpotlightDelegate.DeleteSpotlightIndexWithCompletionHandler]
type ErrorHandler = func(error)

// NewErrorBlock wraps a Go [ErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [NSCoreDataCoreSpotlightDelegate.DeleteSpotlightIndexWithCompletionHandler]
func NewErrorBlock(handler ErrorHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, errID objc.ID) {
		handler(foundation.SafeErrorFrom(errID))
	})
	objc.SetNSErrorBlockSignature(block)
	return objc.ID(block), func() { block.Release() }
}

// ManagedObjectContextHandler handles A closure that is executed by the persistent container against a newly created private context.
//
// Used by:
//   - [NSPersistentContainer.PerformBackgroundTask]
type ManagedObjectContextHandler = func(*NSManagedObjectContext)

// NewManagedObjectContextBlock wraps a Go [ManagedObjectContextHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [NSPersistentContainer.PerformBackgroundTask]
func NewManagedObjectContextBlock(handler ManagedObjectContextHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID) {
		var result *NSManagedObjectContext
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := NSManagedObjectContextFromID(resultID)
			result = &v
		}
		handler(result)
	})
	return objc.ID(block), func() { block.Release() }
}

// NSPersistentStoreAsynchronousFetchResultCompletionBlock handles A completion block that an asynchronous fetch request calls with a result.

// NewNSPersistentStoreAsynchronousFetchResultCompletionBlock wraps a Go [NSPersistentStoreAsynchronousFetchResultCompletionBlock] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewNSPersistentStoreAsynchronousFetchResultCompletionBlock(handler NSPersistentStoreAsynchronousFetchResultCompletionBlock) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitiveVal NSAsynchronousFetchResult) {
		handler(primitiveVal)
	})
	return objc.ID(block), func() { block.Release() }
}

// PersistentStoreDescriptionErrorHandler handles The completion handler block that’s invoked after the store is added.
// The error can be type-asserted to *foundation.NSError for Domain, Code, and UserInfo.
//
// Used by:
//   - [NSPersistentContainer.LoadPersistentStoresWithCompletionHandler]
//   - [NSPersistentStoreCoordinator.AddPersistentStoreWithDescriptionCompletionHandler]
type PersistentStoreDescriptionErrorHandler = func(*NSPersistentStoreDescription, error)

// NewPersistentStoreDescriptionErrorBlock wraps a Go [PersistentStoreDescriptionErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [NSPersistentContainer.LoadPersistentStoresWithCompletionHandler]
//   - [NSPersistentStoreCoordinator.AddPersistentStoreWithDescriptionCompletionHandler]
func NewPersistentStoreDescriptionErrorBlock(handler PersistentStoreDescriptionErrorHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, errID objc.ID) {
		var result *NSPersistentStoreDescription
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := NSPersistentStoreDescriptionFromID(resultID)
			result = &v
		}
		handler(result, foundation.SafeErrorFrom(errID))
	})
	return objc.ID(block), func() { block.Release() }
}

// VoidHandler handles The closure to perform.
//
// Used by:
//   - [NSCoreDataCoreSpotlightDelegate.SearchableIndexReindexAllSearchableItemsWithAcknowledgementHandler]
//   - [NSCoreDataCoreSpotlightDelegate.SearchableIndexReindexSearchableItemsWithIdentifiersAcknowledgementHandler]
//   - [NSManagedObjectContext.PerformBlockAndWait]
//   - [NSManagedObjectContext.PerformBlock]
//   - [NSPersistentStoreCoordinator.PerformBlockAndWait]
//   - [NSPersistentStoreCoordinator.PerformBlock]
type VoidHandler = func()

// NewVoidBlock wraps a Go [VoidHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [NSCoreDataCoreSpotlightDelegate.SearchableIndexReindexAllSearchableItemsWithAcknowledgementHandler]
//   - [NSCoreDataCoreSpotlightDelegate.SearchableIndexReindexSearchableItemsWithIdentifiersAcknowledgementHandler]
//   - [NSManagedObjectContext.PerformBlockAndWait]
//   - [NSManagedObjectContext.PerformBlock]
//   - [NSPersistentStoreCoordinator.PerformBlockAndWait]
//   - [NSPersistentStoreCoordinator.PerformBlock]
func NewVoidBlock(handler VoidHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block) {
		handler()
	})
	return objc.ID(block), func() { block.Release() }
}

// idNSFetchRequestResultAsynchronousFetchResultHandler is the signature for a completion handler block.
//
// Used by:
//   - [NSAsynchronousFetchRequest.InitWithFetchRequestCompletionBlock]
type idNSFetchRequestResultAsynchronousFetchResultHandler = func(*NSAsynchronousFetchResult)

// NewidNSFetchRequestResultAsynchronousFetchResultBlock wraps a Go [idNSFetchRequestResultAsynchronousFetchResultHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [NSAsynchronousFetchRequest.InitWithFetchRequestCompletionBlock]
func NewidNSFetchRequestResultAsynchronousFetchResultBlock(handler idNSFetchRequestResultAsynchronousFetchResultHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID) {
		var result *NSAsynchronousFetchResult
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := NSAsynchronousFetchResultFromID(resultID)
			result = &v
		}
		handler(result)
	})
	return objc.ID(block), func() { block.Release() }
}
