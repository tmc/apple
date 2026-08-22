// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSPersistentHistoryTransaction] class.
var (
	_NSPersistentHistoryTransactionClass     NSPersistentHistoryTransactionClass
	_NSPersistentHistoryTransactionClassOnce sync.Once
)

func getNSPersistentHistoryTransactionClass() NSPersistentHistoryTransactionClass {
	_NSPersistentHistoryTransactionClassOnce.Do(func() {
		_NSPersistentHistoryTransactionClass = NSPersistentHistoryTransactionClass{class: objc.GetClass("NSPersistentHistoryTransaction")}
	})
	return _NSPersistentHistoryTransactionClass
}

// GetNSPersistentHistoryTransactionClass returns the class object for NSPersistentHistoryTransaction.
func GetNSPersistentHistoryTransactionClass() NSPersistentHistoryTransactionClass {
	return getNSPersistentHistoryTransactionClass()
}

type NSPersistentHistoryTransactionClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSPersistentHistoryTransactionClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSPersistentHistoryTransactionClass) Alloc() NSPersistentHistoryTransaction {
	rv := objc.Send[NSPersistentHistoryTransaction](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A set of changes in the persistent history based on a context save or batch
// operation.
//
// # Requesting Notifications
//
//   - [NSPersistentHistoryTransaction.ObjectIDNotification]: Obtains a notification for use in merging the transaction’s changes into a managed object context.
//
// # Inspecting Transaction Details
//
//   - [NSPersistentHistoryTransaction.Author]: A granular description of the context that made the persistent history change, if available.
//   - [NSPersistentHistoryTransaction.BundleID]: The originating bundle’s identifier.
//   - [NSPersistentHistoryTransaction.Changes]: The array of persistent history changes.
//   - [NSPersistentHistoryTransaction.ContextName]: The originating context’s name.
//   - [NSPersistentHistoryTransaction.ProcessID]: The originating process’s identifier.
//   - [NSPersistentHistoryTransaction.StoreID]: The originating store’s identifier.
//   - [NSPersistentHistoryTransaction.Timestamp]: The date of the persistent history change.
//   - [NSPersistentHistoryTransaction.Token]: The token that represents this transaction in the persistent history.
//   - [NSPersistentHistoryTransaction.TransactionNumber]: The transaction’s numeric identifier.
//
// See: https://developer.apple.com/documentation/CoreData/NSPersistentHistoryTransaction
type NSPersistentHistoryTransaction struct {
	objectivec.Object
}

// NSPersistentHistoryTransactionFromID constructs a [NSPersistentHistoryTransaction] from an objc.ID.
//
// A set of changes in the persistent history based on a context save or batch
// operation.
func NSPersistentHistoryTransactionFromID(id objc.ID) NSPersistentHistoryTransaction {
	return NSPersistentHistoryTransaction{objectivec.Object{ID: id}}
}

// NOTE: NSPersistentHistoryTransaction adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSPersistentHistoryTransaction] class.
//
// # Requesting Notifications
//
//   - [INSPersistentHistoryTransaction.ObjectIDNotification]: Obtains a notification for use in merging the transaction’s changes into a managed object context.
//
// # Inspecting Transaction Details
//
//   - [INSPersistentHistoryTransaction.Author]: A granular description of the context that made the persistent history change, if available.
//   - [INSPersistentHistoryTransaction.BundleID]: The originating bundle’s identifier.
//   - [INSPersistentHistoryTransaction.Changes]: The array of persistent history changes.
//   - [INSPersistentHistoryTransaction.ContextName]: The originating context’s name.
//   - [INSPersistentHistoryTransaction.ProcessID]: The originating process’s identifier.
//   - [INSPersistentHistoryTransaction.StoreID]: The originating store’s identifier.
//   - [INSPersistentHistoryTransaction.Timestamp]: The date of the persistent history change.
//   - [INSPersistentHistoryTransaction.Token]: The token that represents this transaction in the persistent history.
//   - [INSPersistentHistoryTransaction.TransactionNumber]: The transaction’s numeric identifier.
//
// See: https://developer.apple.com/documentation/CoreData/NSPersistentHistoryTransaction
type INSPersistentHistoryTransaction interface {
	objectivec.IObject

	// Topic: Requesting Notifications

	// Obtains a notification for use in merging the transaction’s changes into a managed object context.
	ObjectIDNotification() foundation.NSNotification

	// Topic: Inspecting Transaction Details

	// A granular description of the context that made the persistent history change, if available.
	Author() string
	// The originating bundle’s identifier.
	BundleID() string
	// The array of persistent history changes.
	Changes() []NSPersistentHistoryChange
	// The originating context’s name.
	ContextName() string
	// The originating process’s identifier.
	ProcessID() string
	// The originating store’s identifier.
	StoreID() string
	// The date of the persistent history change.
	Timestamp() foundation.NSDate
	// The token that represents this transaction in the persistent history.
	Token() INSPersistentHistoryToken
	// The transaction’s numeric identifier.
	TransactionNumber() int64
}

// Init initializes the instance.
func (p NSPersistentHistoryTransaction) Init() NSPersistentHistoryTransaction {
	rv := objc.Send[NSPersistentHistoryTransaction](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p NSPersistentHistoryTransaction) Autorelease() NSPersistentHistoryTransaction {
	rv := objc.Send[NSPersistentHistoryTransaction](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSPersistentHistoryTransaction creates a new NSPersistentHistoryTransaction instance.
func NewNSPersistentHistoryTransaction() NSPersistentHistoryTransaction {
	class := getNSPersistentHistoryTransactionClass()
	rv := objc.Send[NSPersistentHistoryTransaction](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Obtains a notification for use in merging the transaction’s changes into
// a managed object context.
//
// # Return Value
//
// An [NSManagedObjectContextDidSaveObjectIDsNotification] notification.
//
// # Discussion
//
// To merge the relevant changes into your view context, first obtain a
// notification by calling `objectIDNotification()` on the transaction. Then,
// pass the notification to
// [NSManagedObjectContext.MergeChangesFromContextDidSaveNotification].
//
// See: https://developer.apple.com/documentation/CoreData/NSPersistentHistoryTransaction/objectIDNotification()
func (p NSPersistentHistoryTransaction) ObjectIDNotification() foundation.NSNotification {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("objectIDNotification"))
	return foundation.NSNotificationFromID(rv)
}

// Requests an entity description using the provided context for the managed
// object type affected by the transaction.
//
// context: The managed object context for this request.
//
// # Return Value
//
// The entity description ([NSEntityDescription]) of the persistent history
// transaction entity.
//
// See: https://developer.apple.com/documentation/CoreData/NSPersistentHistoryTransaction/entityDescription(with:)
func (_NSPersistentHistoryTransactionClass NSPersistentHistoryTransactionClass) EntityDescriptionWithContext(context INSManagedObjectContext) NSEntityDescription {
	rv := objc.Send[objc.ID](objc.ID(_NSPersistentHistoryTransactionClass.class), objc.Sel("entityDescriptionWithContext:"), context)
	return NSEntityDescriptionFromID(rv)
}

// A granular description of the context that made the persistent history
// change, if available.
//
// # Discussion
//
// This property has a value if the managed object context set a
// [NSManagedObjectContext.TransactionAuthor] before the save.
//
// See: https://developer.apple.com/documentation/CoreData/NSPersistentHistoryTransaction/author
func (p NSPersistentHistoryTransaction) Author() string {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("author"))
	return foundation.NSStringFromID(rv).String()
}

// The originating bundle’s identifier.
//
// See: https://developer.apple.com/documentation/CoreData/NSPersistentHistoryTransaction/bundleID
func (p NSPersistentHistoryTransaction) BundleID() string {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("bundleID"))
	return foundation.NSStringFromID(rv).String()
}

// The array of persistent history changes.
//
// See: https://developer.apple.com/documentation/CoreData/NSPersistentHistoryTransaction/changes
func (p NSPersistentHistoryTransaction) Changes() []NSPersistentHistoryChange {
	rv := objc.Send[[]objc.ID](p.ID, objc.Sel("changes"))
	return objc.ConvertSlice(rv, func(id objc.ID) NSPersistentHistoryChange {
		return NSPersistentHistoryChangeFromID(id)
	})
}

// The originating context’s name.
//
// See: https://developer.apple.com/documentation/CoreData/NSPersistentHistoryTransaction/contextName
func (p NSPersistentHistoryTransaction) ContextName() string {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("contextName"))
	return foundation.NSStringFromID(rv).String()
}

// The originating process’s identifier.
//
// See: https://developer.apple.com/documentation/CoreData/NSPersistentHistoryTransaction/processID
func (p NSPersistentHistoryTransaction) ProcessID() string {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("processID"))
	return foundation.NSStringFromID(rv).String()
}

// The originating store’s identifier.
//
// See: https://developer.apple.com/documentation/CoreData/NSPersistentHistoryTransaction/storeID
func (p NSPersistentHistoryTransaction) StoreID() string {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("storeID"))
	return foundation.NSStringFromID(rv).String()
}

// The date of the persistent history change.
//
// See: https://developer.apple.com/documentation/CoreData/NSPersistentHistoryTransaction/timestamp
func (p NSPersistentHistoryTransaction) Timestamp() foundation.NSDate {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("timestamp"))
	return foundation.NSDateFromID(objc.ID(rv))
}

// The token that represents this transaction in the persistent history.
//
// See: https://developer.apple.com/documentation/CoreData/NSPersistentHistoryTransaction/token
func (p NSPersistentHistoryTransaction) Token() INSPersistentHistoryToken {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("token"))
	return NSPersistentHistoryTokenFromID(objc.ID(rv))
}

// The transaction’s numeric identifier.
//
// See: https://developer.apple.com/documentation/CoreData/NSPersistentHistoryTransaction/transactionNumber
func (p NSPersistentHistoryTransaction) TransactionNumber() int64 {
	rv := objc.Send[int64](p.ID, objc.Sel("transactionNumber"))
	return rv
}

// A fetch request that has the persistent history transaction as the entity.
//
// See: https://developer.apple.com/documentation/CoreData/NSPersistentHistoryTransaction/fetchRequest
func (_NSPersistentHistoryTransactionClass NSPersistentHistoryTransactionClass) FetchRequest() NSFetchRequest {
	rv := objc.Send[objc.ID](objc.ID(_NSPersistentHistoryTransactionClass.class), objc.Sel("fetchRequest"))
	return NSFetchRequestFromID(objc.ID(rv))
}

// The entity description of the persistent history transaction entity.
//
// # Discussion
//
// The entity description of [NSPersistentHistoryTransaction] lists the
// properties of the persistent history change. This can be useful for
// filtering your request.
//
// See: https://developer.apple.com/documentation/CoreData/NSPersistentHistoryTransaction/entityDescription
func (_NSPersistentHistoryTransactionClass NSPersistentHistoryTransactionClass) EntityDescription() NSEntityDescription {
	rv := objc.Send[objc.ID](objc.ID(_NSPersistentHistoryTransactionClass.class), objc.Sel("entityDescription"))
	return NSEntityDescriptionFromID(objc.ID(rv))
}
