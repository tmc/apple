// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"context"
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CKSyncEngine] class.
var (
	_CKSyncEngineClass     CKSyncEngineClass
	_CKSyncEngineClassOnce sync.Once
)

func getCKSyncEngineClass() CKSyncEngineClass {
	_CKSyncEngineClassOnce.Do(func() {
		_CKSyncEngineClass = CKSyncEngineClass{class: objc.GetClass("CKSyncEngine")}
	})
	return _CKSyncEngineClass
}

// GetCKSyncEngineClass returns the class object for CKSyncEngine.
func GetCKSyncEngineClass() CKSyncEngineClass {
	return getCKSyncEngineClass()
}

type CKSyncEngineClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CKSyncEngineClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CKSyncEngineClass) Alloc() CKSyncEngine {
	rv := objc.Send[CKSyncEngine](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// An object that manages the synchronization of local and remote record data.
//
// # Overview
//
// Use [CKSyncEngine] to handle your app’s CloudKit sync operations and
// benefit from the performance and reliability it provides. To use the class,
// create an instance early in your app’s launch process and specify a
// database to sync. Thereafter, and depending on good system conditions, the
// sync engine periodically pushes and pulls database and record zone changes
// on the app’s behalf. To participate in those sync operations and to
// provide the engine with the changes to send, create an object that conforms
// to [CKSyncEngineDelegate] and assign an instance of it to the engine’s
// configuration. You can have multiple instances of [CKSyncEngine] in a
// single process, each targeting a different database. For example, you may
// have one syncing a person’s private database and another syncing their
// shared database.
//
// Because periodic sync relies on good system conditions — adequate battery
// charge, an active network connection, a signed-in iCloud account, and so on
// — the engine’s sync schedule is indeterminate. If you need to sync
// immediately, like when your app requires it has the most recent changes
// before continuing, use the
// [CKSyncEngine.FetchChangesWithOptionsCompletionHandler] and
// [CKSyncEngine.SendChangesWithOptionsCompletionHandler] methods.
//
// The sync engine uses an opaque type to track its internal state, and it’s
// your responsibility to persist that state to disk and make it available
// across app launches so the engine can function properly. For more
// information, see [SyncEngineHandleEvent] and
// [CKSyncEngineStateUpdateEvent].
//
// [CKSyncEngine] requires the CloudKit and Remote notifications entitlements.
// For more information, see [Configuring iCloud services] and [Configuring
// background execution modes].
//
// # Send changes to iCloud
//
// A sync engine requires you to tell it about any changes to send, which you
// do by invoking the [CKSyncEngineState.AddPendingDatabaseChanges] and
// [CKSyncEngineState.AddPendingRecordZoneChanges] methods on the engine’s
// [CKSyncEngine.State] property. If there are no scheduled sync operations
// when you invoke these methods, the engine automatically schedules one.
// Database changes don’t require any additional input, but the sync engine
// does expect you to provide the individual record zone changes — in
// batches — and return them from your delegate’s implementation of
// [SyncEngineNextRecordZoneChangeBatchForContext]. After the engine sends the
// changes, it notifies your delegate about their success (or failure) by
// dispatching events of type [CKSyncEngineSentDatabaseChangesEvent] and
// [CKSyncEngineSentRecordZoneChangesEvent].
//
// # Fetch changes from iCloud
//
// By default, a sync engine attempts to discover an existing
// [CKDatabaseSubscription] for the associated database and uses that to
// receive silent notifications about remote record changes. If the engine
// doesn’t find a subscription, it automatically creates one to use. On
// receipt of a notification, the engine schedules a sync operation to fetch
// the related changes. When that operation runs, the engine dispatches an
// instance of [CKSyncEngineWillFetchChangesEvent] to your delegate. As it
// receives fetched changes, the engine dispatches
// [CKSyncEngineFetchedDatabaseChangesEvent] and
// [CKSyncEngineFetchedRecordZoneChangesEvent], accordingly. After the
// operation finishes, the sync engine notifies your delegate by dispatching
// an instance of [CKSyncEngineDidFetchChangesEvent]. You handle all
// dispatched events in your delegate’s implementation of
// [SyncEngineHandleEvent].
//
// # Sync Scheduling
//
// # Automatic sync
//
// By default, the sync engine automatically schedules sync tasks on your
// behalf. If the user is signed in, the device has a network connection, and
// the system is generally in a good state, these scheduled syncs happen
// relatively quickly. However, if the device has no network, is low on power,
// or is otherwise under a heavy load, these automatic syncs might be delayed.
// Similarly, if the user isn’t signed in to an account, the sync engine
// won’t perform any sync tasks at all.
//
// # Manual sync
//
// There may be some cases where you want to manually trigger a sync. For
// example, if you have a pull-to-refresh UI, you can call
// [CKSyncEngine.FetchChangesWithOptionsCompletionHandler] to tell the sync
// engine to fetch immediately. Or, if you have a “backup now” UI, you can
// call [CKSyncEngine.SendChangesWithOptionsCompletionHandler] to send to the
// server immediately.
//
// # Error Handling
//
// There are some transient errors that the sync engine handles automatically
// behind the scenes. The sync engine retries the operations for these
// transient errors automatically when it makes sense to do so. Specifically,
// the sync engine will handle the following errors on your behalf:
//
// - [CKError.Code.notAuthenticated] -
// [CKError.Code.accountTemporarilyUnavailable] -
// [CKError.Code.networkFailure] - [CKError.Code.networkUnavailable] -
// [CKError.Code.requestRateLimited] - [CKError.Code.serviceUnavailable] -
// [CKError.Code.zoneBusy]
//
// When the sync engine encounters one of these errors, it waits for the
// system to be in a good state, and tries again. For example, if the server
// sends back a [CKError.Code.requestRateLimited] error, the sync engine
// respects this throttle and tries again after the error’s retry-after
// time.
//
// [CKSyncEngine] does handle errors that require application-specific logic.
// For example, if you try to save a record and get a
// [CKError.Code.serverRecordChanged], you need to handle that error yourself.
// There are plenty of errors that the sync engine cannot handle on your
// behalf, see [CKError] for a list of all the possible errors.
//
// # Accounts
//
// [CKSyncEngine] monitors for account status, and it only syncs if there’s
// an account signed in. Because of this, you can initialize your
// [CKSyncEngine] at any time, regardless of account status. If there is no
// account, or if the user disabled sync in settings, the sync engine stays
// dormant in the background. Once an account is available, the sync engine
// starts syncing automatically.
//
// The sync engine listens for when the user signs in or out of their account.
// When it notices an account change, it sends an
// [CKSyncEngineAccountChangeEvent] to your delegate. It’s your
// responsibility to react appropriately to this change and update your local
// persistence.
//
// # Creating a sync engine
//
//   - [CKSyncEngine.InitWithConfiguration]: Creates a sync engine with the specified configuration.
//
// # Accessing the engine’s attributes
//
//   - [CKSyncEngine.Database]: The associated database.
//   - [CKSyncEngine.State]: The sync engine’s state.
//
// # Invoking manual sync operations
//
//   - [CKSyncEngine.FetchChangesWithCompletionHandler]: Fetches pending remote changes from the server.
//   - [CKSyncEngine.FetchChangesWithOptionsCompletionHandler]: Fetches pending remote changes from the server using the specified options.
//   - [CKSyncEngine.SendChangesWithCompletionHandler]: Sends pending local changes to the server.
//   - [CKSyncEngine.SendChangesWithOptionsCompletionHandler]: Sends pending local changes to the server using the specified options.
//
// # Canceling operations
//
//   - [CKSyncEngine.CancelOperationsWithCompletionHandler]: Cancels any in-progress or pending sync operations.
//
// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngine-4b4w9
//
// [CKError.Code.accountTemporarilyUnavailable]: https://developer.apple.com/documentation/CloudKit/CKError/Code/accountTemporarilyUnavailable
// [CKError.Code.networkFailure]: https://developer.apple.com/documentation/CloudKit/CKError/Code/networkFailure
// [CKError.Code.networkUnavailable]: https://developer.apple.com/documentation/CloudKit/CKError/Code/networkUnavailable
// [CKError.Code.notAuthenticated]: https://developer.apple.com/documentation/CloudKit/CKError/Code/notAuthenticated
// [CKError.Code.requestRateLimited]: https://developer.apple.com/documentation/CloudKit/CKError/Code/requestRateLimited
// [CKError.Code.serverRecordChanged]: https://developer.apple.com/documentation/CloudKit/CKError/Code/serverRecordChanged
// [CKError.Code.serviceUnavailable]: https://developer.apple.com/documentation/CloudKit/CKError/Code/serviceUnavailable
// [CKError.Code.zoneBusy]: https://developer.apple.com/documentation/CloudKit/CKError/Code/zoneBusy
// [CKError]: https://developer.apple.com/documentation/CloudKit/CKError
// [Configuring background execution modes]: https://developer.apple.com/documentation/Xcode/configuring-background-execution-modes
// [Configuring iCloud services]: https://developer.apple.com/documentation/Xcode/configuring-icloud-services
type CKSyncEngine struct {
	objectivec.Object
}

// CKSyncEngineFromID constructs a [CKSyncEngine] from an objc.ID.
//
// An object that manages the synchronization of local and remote record data.
func CKSyncEngineFromID(id objc.ID) CKSyncEngine {
	return CKSyncEngine{objectivec.Object{ID: id}}
}

// Ensure CKSyncEngine implements ICKSyncEngine.
var _ ICKSyncEngine = CKSyncEngine{}

// An interface definition for the [CKSyncEngine] class.
//
// # Creating a sync engine
//
//   - [ICKSyncEngine.InitWithConfiguration]: Creates a sync engine with the specified configuration.
//
// # Accessing the engine’s attributes
//
//   - [ICKSyncEngine.Database]: The associated database.
//   - [ICKSyncEngine.State]: The sync engine’s state.
//
// # Invoking manual sync operations
//
//   - [ICKSyncEngine.FetchChangesWithCompletionHandler]: Fetches pending remote changes from the server.
//   - [ICKSyncEngine.FetchChangesWithOptionsCompletionHandler]: Fetches pending remote changes from the server using the specified options.
//   - [ICKSyncEngine.SendChangesWithCompletionHandler]: Sends pending local changes to the server.
//   - [ICKSyncEngine.SendChangesWithOptionsCompletionHandler]: Sends pending local changes to the server using the specified options.
//
// # Canceling operations
//
//   - [ICKSyncEngine.CancelOperationsWithCompletionHandler]: Cancels any in-progress or pending sync operations.
//
// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngine-4b4w9
type ICKSyncEngine interface {
	objectivec.IObject

	// Topic: Creating a sync engine

	// Creates a sync engine with the specified configuration.
	InitWithConfiguration(configuration ICKSyncEngineConfiguration) CKSyncEngine

	// Topic: Accessing the engine’s attributes

	// The associated database.
	Database() ICKDatabase
	// The sync engine’s state.
	State() ICKSyncEngineState

	// Topic: Invoking manual sync operations

	// Fetches pending remote changes from the server.
	FetchChangesWithCompletionHandler(completionHandler ErrorHandler)
	// Fetches pending remote changes from the server using the specified options.
	FetchChangesWithOptionsCompletionHandler(options ICKSyncEngineFetchChangesOptions, completionHandler ErrorHandler)
	// Sends pending local changes to the server.
	SendChangesWithCompletionHandler(completionHandler ErrorHandler)
	// Sends pending local changes to the server using the specified options.
	SendChangesWithOptionsCompletionHandler(options ICKSyncEngineSendChangesOptions, completionHandler ErrorHandler)

	// Topic: Canceling operations

	// Cancels any in-progress or pending sync operations.
	CancelOperationsWithCompletionHandler(completionHandler VoidHandler)
}

// Init initializes the instance.
func (c CKSyncEngine) Init() CKSyncEngine {
	rv := objc.Send[CKSyncEngine](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CKSyncEngine) Autorelease() CKSyncEngine {
	rv := objc.Send[CKSyncEngine](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKSyncEngine creates a new CKSyncEngine instance.
func NewCKSyncEngine() CKSyncEngine {
	class := getCKSyncEngineClass()
	rv := objc.Send[CKSyncEngine](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a sync engine with the specified configuration.
//
// configuration: The attributes of the new sync engine, such as the associated database and
// the object to use as the engine’s delegate. For more information, see
// [CKSyncEngineConfiguration].
//
// # Return Value
//
// A configured sync engine, or `nil` if CloudKit can’t create one.
//
// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngine-4b4w9/initWithConfiguration:
func NewCKSyncEngineWithConfiguration(configuration ICKSyncEngineConfiguration) CKSyncEngine {
	instance := getCKSyncEngineClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithConfiguration:"), configuration)
	return CKSyncEngineFromID(rv)
}

// Creates a sync engine with the specified configuration.
//
// configuration: The attributes of the new sync engine, such as the associated database and
// the object to use as the engine’s delegate. For more information, see
// [CKSyncEngineConfiguration].
//
// # Return Value
//
// A configured sync engine, or `nil` if CloudKit can’t create one.
//
// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngine-4b4w9/initWithConfiguration:
func (c CKSyncEngine) InitWithConfiguration(configuration ICKSyncEngineConfiguration) CKSyncEngine {
	rv := objc.Send[CKSyncEngine](c.ID, objc.Sel("initWithConfiguration:"), configuration)
	return rv
}

// Fetches pending remote changes from the server.
//
// completionHandler: The block to execute when the fetch completes.
//
// # Discussion
//
// If the fetch fails, the completion handler’s `error` parameter is an
// object that describes that failure; otherwise, it’s `nil`.
//
// Use this method to ensure the sync engine immediatley fetches all pending
// remote changes before your app continues. This isn’t necessary in normal
// use, as the engine automatically syncs your app’s records. It is useful,
// however, in scenarios where you require more control over sync, such as
// pull-to-refresh or unit tests.
//
// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngine-4b4w9/fetchChangesWithCompletionHandler:
func (c CKSyncEngine) FetchChangesWithCompletionHandler(completionHandler ErrorHandler) {
	_block0, _ := NewErrorBlock(completionHandler)
	objc.Send[objc.ID](c.ID, objc.Sel("fetchChangesWithCompletionHandler:"), _block0)
}

// Fetches pending remote changes from the server using the specified options.
//
// options: The options to use when fetching changes. For more information, see
// [CKSyncEngineFetchChangesOptions].
//
// completionHandler: The block to execute when the fetch completes.
//
// # Discussion
//
// If the fetch fails, the completion handler’s `error` parameter is an
// object that describes that failure; otherwise, it’s `nil`.
//
// Use this method to ensure the sync engine immediatley fetches all pending
// remote changes before your app continues. This isn’t necessary in normal
// use, as the engine automatically syncs your app’s records. It is useful,
// however, in scenarios where you require more control over sync, such as
// pull-to-refresh or unit tests.
//
// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngine-4b4w9/fetchChangesWithOptions:completionHandler:
func (c CKSyncEngine) FetchChangesWithOptionsCompletionHandler(options ICKSyncEngineFetchChangesOptions, completionHandler ErrorHandler) {
	_block1, _ := NewErrorBlock(completionHandler)
	objc.Send[objc.ID](c.ID, objc.Sel("fetchChangesWithOptions:completionHandler:"), options, _block1)
}

// Sends pending local changes to the server.
//
// completionHandler: The block to execute when the send completes.
//
// # Discussion
//
// If the send fails, the completion handler’s `error` parameter is an
// object that describes that failure; otherwise, it’s `nil`.
//
// Use this method to ensure the sync engine sends all pending local changes
// to the server before your app continues. This isn’t necessary in normal
// use, as the engine automatically syncs your app’s records. It is useful,
// however, in scenarios where you require greater control over sync, such as
// a “Backup now” button or unit tests.
//
// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngine-4b4w9/sendChangesWithCompletionHandler:
func (c CKSyncEngine) SendChangesWithCompletionHandler(completionHandler ErrorHandler) {
	_block0, _ := NewErrorBlock(completionHandler)
	objc.Send[objc.ID](c.ID, objc.Sel("sendChangesWithCompletionHandler:"), _block0)
}

// Sends pending local changes to the server using the specified options.
//
// options: The options to use when sending changes. For more information, see
// [CKSyncEngineSendChangesOptions].
//
// completionHandler: The block to execute when the send completes.
//
// # Discussion
//
// If the send fails, the completion handler’s `error` parameter is an
// object that describes that failure; otherwise, it’s `nil`.
//
// Use this method to ensure the sync engine sends all pending local changes
// to the server before your app continues. This isn’t necessary in normal
// use, as the engine automatically syncs your app’s records. It is useful,
// however, in scenarios where you require greater control over sync, such as
// a “Backup now” button or unit tests.
//
// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngine-4b4w9/sendChangesWithOptions:completionHandler:
func (c CKSyncEngine) SendChangesWithOptionsCompletionHandler(options ICKSyncEngineSendChangesOptions, completionHandler ErrorHandler) {
	_block1, _ := NewErrorBlock(completionHandler)
	objc.Send[objc.ID](c.ID, objc.Sel("sendChangesWithOptions:completionHandler:"), options, _block1)
}

// Cancels any in-progress or pending sync operations.
//
// # Discussion
//
// The sync engine processes cancelation requests asynchronously, meaning
// it’s possible for in-progress operations to complete even after this
// method returns.
//
// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngine-4b4w9/cancelOperationsWithCompletionHandler:
func (c CKSyncEngine) CancelOperationsWithCompletionHandler(completionHandler VoidHandler) {
	_block0, _ := NewVoidBlock(completionHandler)
	objc.Send[objc.ID](c.ID, objc.Sel("cancelOperationsWithCompletionHandler:"), _block0)
}

// The associated database.
//
// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngine-4b4w9/database
func (c CKSyncEngine) Database() ICKDatabase {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("database"))
	return CKDatabaseFromID(objc.ID(rv))
}

// The sync engine’s state.
//
// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngine-4b4w9/state
func (c CKSyncEngine) State() ICKSyncEngineState {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("state"))
	return CKSyncEngineStateFromID(objc.ID(rv))
}

// FetchChanges is a synchronous wrapper around [CKSyncEngine.FetchChangesWithCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (c CKSyncEngine) FetchChanges(ctx context.Context) error {
	done := make(chan error, 1)
	c.FetchChangesWithCompletionHandler(func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

// FetchChangesWithOptions is a synchronous wrapper around [CKSyncEngine.FetchChangesWithOptionsCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (c CKSyncEngine) FetchChangesWithOptions(ctx context.Context, options ICKSyncEngineFetchChangesOptions) error {
	done := make(chan error, 1)
	c.FetchChangesWithOptionsCompletionHandler(options, func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

// SendChanges is a synchronous wrapper around [CKSyncEngine.SendChangesWithCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (c CKSyncEngine) SendChanges(ctx context.Context) error {
	done := make(chan error, 1)
	c.SendChangesWithCompletionHandler(func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

// SendChangesWithOptions is a synchronous wrapper around [CKSyncEngine.SendChangesWithOptionsCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (c CKSyncEngine) SendChangesWithOptions(ctx context.Context, options ICKSyncEngineSendChangesOptions) error {
	done := make(chan error, 1)
	c.SendChangesWithOptionsCompletionHandler(options, func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

// CancelOperations is a synchronous wrapper around [CKSyncEngine.CancelOperationsWithCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (c CKSyncEngine) CancelOperations(ctx context.Context) error {
	done := make(chan struct{}, 1)
	c.CancelOperationsWithCompletionHandler(func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
