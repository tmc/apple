// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"context"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSFileCoordinator] class.
var (
	_NSFileCoordinatorClass     NSFileCoordinatorClass
	_NSFileCoordinatorClassOnce sync.Once
)

func getNSFileCoordinatorClass() NSFileCoordinatorClass {
	_NSFileCoordinatorClassOnce.Do(func() {
		_NSFileCoordinatorClass = NSFileCoordinatorClass{class: objc.GetClass("NSFileCoordinator")}
	})
	return _NSFileCoordinatorClass
}

// GetNSFileCoordinatorClass returns the class object for NSFileCoordinator.
func GetNSFileCoordinatorClass() NSFileCoordinatorClass {
	return getNSFileCoordinatorClass()
}

type NSFileCoordinatorClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSFileCoordinatorClass) Alloc() NSFileCoordinator {
	rv := objc.Send[NSFileCoordinator](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// An object that coordinates the reading and writing of files and directories
// among file presenters.
//
// # Overview
// 
// The [NSFileCoordinator] class coordinates the reading and writing of files
// and directories among multiple processes and objects in the same process.
// You use instances of this class as is to read from, write to, modify the
// attributes of, change the location of, or delete a file or directory, but
// before your code to perform those actions executes, the file coordinator
// lets registered file presenter objects perform any tasks that they might
// require to ensure their own integrity. For example, if you want to change
// the location of a file, other objects interested in that file need to know
// where you intend to move it so that they can update their references.
// 
// Objects that adopt the [NSFilePresenter] protocol must register themselves
// with the [NSFileCoordinator] class to be notified of any pending changes.
// They do this by calling the [NSFileCoordinator.AddFilePresenter] class method. A file
// presenter must balance calls to [NSFileCoordinator.AddFilePresenter] with a call to
// [NSFileCoordinator.RemoveFilePresenter] before being released, even in a garbage-collected
// application. The file presenter class maintains a list of active file
// presenter objects in the current application and uses that list, plus the
// file coordinator classes in other processes, to deliver notifications to
// all of the objects interested in a particular item.
// 
// Instances of [NSFileCoordinator] are meant to be used on a
// per-file-operation basis, where a file operation is something like opening
// and reading the contents of a file or moving a batch of files and
// directories to a new location. There is no benefit to keeping a file
// coordinator object past the length of the planned operation. In fact,
// because file coordinators retain file presenter objects, keeping one around
// could prevent the file presenter objects from being released in a timely
// manner.
// 
// For information about implementing a file presenter object to receive
// file-related notifications, see [NSFilePresenter].
// 
// # File Presenters and iOS
// 
// If your app or extension enters the background with an active file
// presenter, it may be terminated by the system in order to prevent deadlock
// on that file. To prevent this situation, call [NSFileCoordinator.RemoveFilePresenter] to
// remove the file presenter in the [applicationDidEnterBackground(_:)] method
// or in response to a [NSFileCoordinator.DidEnterBackgroundNotification] notification. Call
// [NSFileCoordinator.AddFilePresenter] to add the file presenter again in the
// [applicationWillEnterForeground(_:)] method or in response to a
// [NSFileCoordinator.WillEnterForegroundNotification] notification.
// 
// # File Coordinators and iOS
// 
// A coordinated read or write will automatically begin a background task when
// granted, similar to one created with the
// [beginBackgroundTask(expirationHandler:)] method. This helps ensure that
// your app or extension has sufficient time to finish the read or write
// operation if it’s suspended, without creating a deadlock on access to
// that file by other processes. If a process is suspended while waiting for a
// coordinated read or write to be granted, the request is canceled, and an
// [NSError] object with the code [NSFileCoordinator.NSUserCancelledError] is produced. If the
// background task expires, the process is terminated.
// 
// # Threading Considerations
// 
// Each file coordinator object you create should be used on a single thread
// only. If you need to coordinate file operations across multiple objects in
// different threads, each object should create its own file coordinator.
//
// [NSFileCoordinator.NSUserCancelledError]: https://developer.apple.com/documentation/Foundation/NSUserCancelledError-swift.var
// [applicationDidEnterBackground(_:)]: https://developer.apple.com/documentation/UIKit/UIApplicationDelegate/applicationDidEnterBackground(_:)
// [applicationWillEnterForeground(_:)]: https://developer.apple.com/documentation/UIKit/UIApplicationDelegate/applicationWillEnterForeground(_:)
// [beginBackgroundTask(expirationHandler:)]: https://developer.apple.com/documentation/UIKit/UIApplication/beginBackgroundTask(expirationHandler:)
//
// # Initializing a File Coordinator
//
//   - [NSFileCoordinator.InitWithFilePresenter]: Initializes and returns a file coordinator object using the specified file presenter.
//
// # Managing File Presenters
//
//   - [NSFileCoordinator.PurposeIdentifier]: A string that uniquely identifies the file access that was performed by this file coordinator.
//   - [NSFileCoordinator.SetPurposeIdentifier]
//
// # Coordinating File Operations Asynchronously
//
//   - [NSFileCoordinator.CoordinateAccessWithIntentsQueueByAccessor]: Performs a number of coordinated-read or -write operations asynchronously.
//
// # Coordinating File Operations Synchronously
//
//   - [NSFileCoordinator.CoordinateReadingItemAtURLOptionsErrorByAccessor]: Initiates a read operation on a single file or directory using the specified options.
//   - [NSFileCoordinator.CoordinateWritingItemAtURLOptionsErrorByAccessor]: Initiates a write operation on a single file or directory using the specified options.
//   - [NSFileCoordinator.CoordinateReadingItemAtURLOptionsWritingItemAtURLOptionsErrorByAccessor]: Initiates a read operation that contains a follow-up write operation.
//   - [NSFileCoordinator.CoordinateWritingItemAtURLOptionsWritingItemAtURLOptionsErrorByAccessor]: Initiates a write operation that involves a secondary write operation.
//   - [NSFileCoordinator.PrepareForReadingItemsAtURLsOptionsWritingItemsAtURLsOptionsErrorByAccessor]: Prepare to read or write from multiple files in a single batch operation.
//   - [NSFileCoordinator.ItemAtURLWillMoveToURL]: Announces that your app is moving a file to a new URL.
//   - [NSFileCoordinator.ItemAtURLDidMoveToURL]: Notifies relevant file presenters that the location of a file or directory changed.
//   - [NSFileCoordinator.Cancel]: Cancels any active file coordination calls.
//
// # Ubiquity Change Notifications
//
//   - [NSFileCoordinator.ItemAtURLDidChangeUbiquityAttributes]: Tells observing file providers that the item’s ubiquity attributes have changed.
//
// See: https://developer.apple.com/documentation/Foundation/NSFileCoordinator
type NSFileCoordinator struct {
	objectivec.Object
}

// NSFileCoordinatorFromID constructs a [NSFileCoordinator] from an objc.ID.
//
// An object that coordinates the reading and writing of files and directories
// among file presenters.
func NSFileCoordinatorFromID(id objc.ID) NSFileCoordinator {
	return NSFileCoordinator{objectivec.Object{ID: id}}
}
// NOTE: NSFileCoordinator adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSFileCoordinator] class.
//
// # Initializing a File Coordinator
//
//   - [INSFileCoordinator.InitWithFilePresenter]: Initializes and returns a file coordinator object using the specified file presenter.
//
// # Managing File Presenters
//
//   - [INSFileCoordinator.PurposeIdentifier]: A string that uniquely identifies the file access that was performed by this file coordinator.
//   - [INSFileCoordinator.SetPurposeIdentifier]
//
// # Coordinating File Operations Asynchronously
//
//   - [INSFileCoordinator.CoordinateAccessWithIntentsQueueByAccessor]: Performs a number of coordinated-read or -write operations asynchronously.
//
// # Coordinating File Operations Synchronously
//
//   - [INSFileCoordinator.CoordinateReadingItemAtURLOptionsErrorByAccessor]: Initiates a read operation on a single file or directory using the specified options.
//   - [INSFileCoordinator.CoordinateWritingItemAtURLOptionsErrorByAccessor]: Initiates a write operation on a single file or directory using the specified options.
//   - [INSFileCoordinator.CoordinateReadingItemAtURLOptionsWritingItemAtURLOptionsErrorByAccessor]: Initiates a read operation that contains a follow-up write operation.
//   - [INSFileCoordinator.CoordinateWritingItemAtURLOptionsWritingItemAtURLOptionsErrorByAccessor]: Initiates a write operation that involves a secondary write operation.
//   - [INSFileCoordinator.PrepareForReadingItemsAtURLsOptionsWritingItemsAtURLsOptionsErrorByAccessor]: Prepare to read or write from multiple files in a single batch operation.
//   - [INSFileCoordinator.ItemAtURLWillMoveToURL]: Announces that your app is moving a file to a new URL.
//   - [INSFileCoordinator.ItemAtURLDidMoveToURL]: Notifies relevant file presenters that the location of a file or directory changed.
//   - [INSFileCoordinator.Cancel]: Cancels any active file coordination calls.
//
// # Ubiquity Change Notifications
//
//   - [INSFileCoordinator.ItemAtURLDidChangeUbiquityAttributes]: Tells observing file providers that the item’s ubiquity attributes have changed.
//
// See: https://developer.apple.com/documentation/Foundation/NSFileCoordinator
type INSFileCoordinator interface {
	objectivec.IObject

	// Topic: Initializing a File Coordinator

	// Initializes and returns a file coordinator object using the specified file presenter.
	InitWithFilePresenter(filePresenterOrNil NSFilePresenter) NSFileCoordinator

	// Topic: Managing File Presenters

	// A string that uniquely identifies the file access that was performed by this file coordinator.
	PurposeIdentifier() string
	SetPurposeIdentifier(value string)

	// Topic: Coordinating File Operations Asynchronously

	// Performs a number of coordinated-read or -write operations asynchronously.
	CoordinateAccessWithIntentsQueueByAccessor(intents []NSFileAccessIntent, queue INSOperationQueue, accessor ErrorHandler)

	// Topic: Coordinating File Operations Synchronously

	// Initiates a read operation on a single file or directory using the specified options.
	CoordinateReadingItemAtURLOptionsErrorByAccessor(url INSURL, options NSFileCoordinatorReadingOptions, outError INSError, reader URLHandler)
	// Initiates a write operation on a single file or directory using the specified options.
	CoordinateWritingItemAtURLOptionsErrorByAccessor(url INSURL, options NSFileCoordinatorWritingOptions, outError INSError, writer URLHandler)
	// Initiates a read operation that contains a follow-up write operation.
	CoordinateReadingItemAtURLOptionsWritingItemAtURLOptionsErrorByAccessor(readingURL INSURL, readingOptions NSFileCoordinatorReadingOptions, writingURL INSURL, writingOptions NSFileCoordinatorWritingOptions, outError INSError, readerWriter URLURLHandler)
	// Initiates a write operation that involves a secondary write operation.
	CoordinateWritingItemAtURLOptionsWritingItemAtURLOptionsErrorByAccessor(url1 INSURL, options1 NSFileCoordinatorWritingOptions, url2 INSURL, options2 NSFileCoordinatorWritingOptions, outError INSError, writer URLURLHandler)
	// Prepare to read or write from multiple files in a single batch operation.
	PrepareForReadingItemsAtURLsOptionsWritingItemsAtURLsOptionsErrorByAccessor(readingURLs []NSURL, readingOptions NSFileCoordinatorReadingOptions, writingURLs []NSURL, writingOptions NSFileCoordinatorWritingOptions, outError INSError, batchAccessor VoidHandler)
	// Announces that your app is moving a file to a new URL.
	ItemAtURLWillMoveToURL(oldURL INSURL, newURL INSURL)
	// Notifies relevant file presenters that the location of a file or directory changed.
	ItemAtURLDidMoveToURL(oldURL INSURL, newURL INSURL)
	// Cancels any active file coordination calls.
	Cancel()

	// Topic: Ubiquity Change Notifications

	// Tells observing file providers that the item’s ubiquity attributes have changed.
	ItemAtURLDidChangeUbiquityAttributes(url INSURL, attributes INSSet)

	// The user canceled the operation (for example, by pressing Command-period).
	NSUserCancelledError() int
	SetNSUserCancelledError(value int)
}

// Init initializes the instance.
func (f NSFileCoordinator) Init() NSFileCoordinator {
	rv := objc.Send[NSFileCoordinator](f.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f NSFileCoordinator) Autorelease() NSFileCoordinator {
	rv := objc.Send[NSFileCoordinator](f.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSFileCoordinator creates a new NSFileCoordinator instance.
func NewNSFileCoordinator() NSFileCoordinator {
	class := getNSFileCoordinatorClass()
	rv := objc.Send[NSFileCoordinator](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Initializes and returns a file coordinator object using the specified file
// presenter.
//
// filePresenterOrNil: The file presenter object that is initiating some action on its file or
// directory. This object is assumed to be performing the relevant file or
// directory operations and therefore does not receive notifications about
// those operations from the returned file coordinator object. This parameter
// may be `nil`.
//
// # Return Value
// 
// A file coordinator object to use to coordinate file-related operations.
//
// # Discussion
// 
// Specifying a file presenter at initialization time is strongly recommended,
// especially if the file presenter is initiating the file operation.
// Otherwise, the file presenter itself would receive notifications when it
// made changes to the file and would have to compensate for that fact.
// Receiving such notifications could also deadlock if the file presenter’s
// code and its notifications run on the same thread.
// 
// Specifically, associating an NSFileCoordinator with an NSFilePresenter
// accomplishes the following:
// 
// - Prevents the file coordinator from sending messages to the file
// presenter. This means that the file presenter won’t receive notifications
// about its own file operations. There is one exception: Messages about
// versions of the presented item being added, remove, or resolved during
// coordinated writing are sent to every relevant file presenter. - Prevents
// deadlocks that could occur when the file presenter performs a coordinated
// write operation in response to a
// [SavePresentedItemChangesWithCompletionHandler] message. Usually,
// coordinated writes must wait for all the coordinated read operations on the
// same file or directory. However, when a coordinated read forces a file
// presenter to write its contents, the write operation must proceed before
// the read operation can finish. - Prevents race conditions that could occur
// when the file presenter is sent a [PresentedItemDidMoveToURL] message and
// at the same time—but before this message is dequeued—the file presenter
// enqueues an operation using the old URL on a different queue. For the file
// coordinators to work effectively, however, the coordinator must be
// initialized on the same operation queue that the file presenter uses to
// receive its messages. - Allows the file coordination mechanism to
// gracefully handle file presenters that initially contain `nil` in the
// [PresentedItemURL] property, but that can later contain a non-`nil` value
// after creating the item using a coordinated write operation. For example,
// AppKit uses this feature to instantiate new [NSDocument] objects
// immediately, instead of waiting until after the user saves the document.
//
// [NSDocument]: https://developer.apple.com/documentation/AppKit/NSDocument
//
// See: https://developer.apple.com/documentation/Foundation/NSFileCoordinator/init(filePresenter:)
func NewFileCoordinatorWithFilePresenter(filePresenterOrNil NSFilePresenter) NSFileCoordinator {
	instance := getNSFileCoordinatorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFilePresenter:"), filePresenterOrNil)
	return NSFileCoordinatorFromID(rv)
}

// Initializes and returns a file coordinator object using the specified file
// presenter.
//
// filePresenterOrNil: The file presenter object that is initiating some action on its file or
// directory. This object is assumed to be performing the relevant file or
// directory operations and therefore does not receive notifications about
// those operations from the returned file coordinator object. This parameter
// may be `nil`.
//
// # Return Value
// 
// A file coordinator object to use to coordinate file-related operations.
//
// # Discussion
// 
// Specifying a file presenter at initialization time is strongly recommended,
// especially if the file presenter is initiating the file operation.
// Otherwise, the file presenter itself would receive notifications when it
// made changes to the file and would have to compensate for that fact.
// Receiving such notifications could also deadlock if the file presenter’s
// code and its notifications run on the same thread.
// 
// Specifically, associating an NSFileCoordinator with an NSFilePresenter
// accomplishes the following:
// 
// - Prevents the file coordinator from sending messages to the file
// presenter. This means that the file presenter won’t receive notifications
// about its own file operations. There is one exception: Messages about
// versions of the presented item being added, remove, or resolved during
// coordinated writing are sent to every relevant file presenter. - Prevents
// deadlocks that could occur when the file presenter performs a coordinated
// write operation in response to a
// [SavePresentedItemChangesWithCompletionHandler] message. Usually,
// coordinated writes must wait for all the coordinated read operations on the
// same file or directory. However, when a coordinated read forces a file
// presenter to write its contents, the write operation must proceed before
// the read operation can finish. - Prevents race conditions that could occur
// when the file presenter is sent a [PresentedItemDidMoveToURL] message and
// at the same time—but before this message is dequeued—the file presenter
// enqueues an operation using the old URL on a different queue. For the file
// coordinators to work effectively, however, the coordinator must be
// initialized on the same operation queue that the file presenter uses to
// receive its messages. - Allows the file coordination mechanism to
// gracefully handle file presenters that initially contain `nil` in the
// [PresentedItemURL] property, but that can later contain a non-`nil` value
// after creating the item using a coordinated write operation. For example,
// AppKit uses this feature to instantiate new [NSDocument] objects
// immediately, instead of waiting until after the user saves the document.
//
// [NSDocument]: https://developer.apple.com/documentation/AppKit/NSDocument
//
// See: https://developer.apple.com/documentation/Foundation/NSFileCoordinator/init(filePresenter:)
func (f NSFileCoordinator) InitWithFilePresenter(filePresenterOrNil NSFilePresenter) NSFileCoordinator {
	rv := objc.Send[NSFileCoordinator](f.ID, objc.Sel("initWithFilePresenter:"), filePresenterOrNil)
	return rv
}

// Performs a number of coordinated-read or -write operations asynchronously.
//
// intents: An array of file access intent objects, representing the individual read
// and write operations.
//
// queue: The operation queue on which the accessor block is executed. The queue must
// not be `nil`.
//
// accessor: A [Block object] containing the file operations corresponding to the file
// access intent objects in the intents array.
// 
// The accessor block takes the following parameter:
// 
// error: If an error occurs while waiting for access, this parameter contains
// an [NSError] object that describes the problem. If access is successfully
// granted, it is set to `nil`, and you may perform the intended file access.
// 
// Do not attempt to access the files if the error parameter contains a
// non-`nil` value.
// //
// [Block object]: https://developer.apple.com/library/archive/documentation/General/Conceptual/DevPedia-CocoaCore/Block.html#//apple_ref/doc/uid/TP40008195-CH3
//
// # Discussion
// 
// This method lets you asynchronously perform coordinated read or writes. You
// can specify any combination of individual read or write operations. The
// file coordinator waits asynchronously to get access to the files and then
// invokes the accessor block on the specified queue.
// 
// If an error occurs while waiting for access, an error message is passed to
// the block. You must check the block’s error parameter before accessing
// any of the files. If the error parameter is set to `nil`, you can freely
// perform the read and write operations described by your intents. Otherwise,
// you may not access the files.
// 
// Additionally, always use the URL property of your intent objects when
// accessing the files inside the accessor block. The system updates this
// property to account for any changes to the underlying files. For example,
// if the files are moved or renamed, the system updates this URL property to
// match.
// 
// Your file coordinator has access to the files only until the accessor block
// returns. Do not dispatch tasks that continue to access these files onto
// other threads or queues. That can cause your app to access the files
// outside of file coordination, and could result in data corruption or data
// loss.
// 
// In general, coordinated-read operations wait for coordinated-write
// operations on the same URL. Coordinated-write operations wait for both
// coordinated-read and coordinated-write operations on the same URL. Multiple
// coordinated reads can occur simultaneously without blocking each other.
// 
// Performing a coordinated read or coordinated write on the contents of a
// file package is treated as a coordinated read or write to the package as a
// whole. In general, coordinated access to a directory that is not a file
// package is not affected by coordinated access to the directory’s
// contents. Similarly, accessing the contents does not affect the directory.
// However, if you make a coordinated write operation using the
// [FileCoordinatorWritingForDeleting], [FileCoordinatorWritingForMoving], or
// [FileCoordinatorWritingForReplacing] options, all coordinated access to the
// directory and its contents interact as if they were accessing the same URL.
// 
// Coordinated reads and writes from the same file coordinator instance never
// block each other. However, if you make multiple, concurrent calls to
// [CoordinateAccessWithIntentsQueueByAccessor], you risk deadlocking with
// another process that is similarly making multiple concurrent calls to its
// file coordinator. Wherever possible, invoke
// [CoordinateAccessWithIntentsQueueByAccessor] once, passing in multiple file
// access intent objects.
// 
// Coordinated-read and -write operations also wait on any file presenters
// methods that are triggered as part of the coordinated access. Coordinated
// access triggers method calls on all the file presenters for the same
// URL—even on file presenters in other processes. There is only one
// exception: file coordinator never sends messages to the file presenter that
// was passed to its [InitWithFilePresenter] method.
// 
// Coordinated reads trigger the following method calls:
// 
// - The [RelinquishPresentedItemToReader] method is called. - If the
// [FileCoordinatorReadingWithoutChanges] option is not used then
// [SavePresentedItemChangesWithCompletionHandler] is called. - If the file or
// directory is part of a file package, these methods are also called on the
// package’s file presenter. If there are nested packages, the methods are
// called on all the packages’ file presenters.
// 
// Coordinated writes trigger the following method calls:
// 
// - The [RelinquishPresentedItemToWriter] method is called. If the target is
// part of a file package, the method is called on the package’s presenter,
// just like in coordinated reads. - If the target is a directory and the
// [FileCoordinatorWritingForDeleting], [FileCoordinatorWritingForMoving], or
// [FileCoordinatorWritingForReplacing] options are used,
// [RelinquishPresentedItemToWriter] is also called on all the file presenters
// for the directory’s contents. - If the
// [FileCoordinatorWritingForDeleting] or [FileCoordinatorWritingForReplacing]
// options are used, the
// [AccommodatePresentedItemDeletionWithCompletionHandler] method is called.
// If the target is a directory, that method is called on the file presenters
// for the directory’s contents as well. If the target is inside a file
// package, the [AccommodatePresentedSubitemDeletionAtURLCompletionHandler]
// method is called on the package’s presenter. - If the
// [FileCoordinatorWritingForReplacing] option is used, the definition of the
// target file or directory can vary depending on how this write operation
// interacts with other coordinated write operations. For more details, see
// [FileCoordinatorWritingForReplacing]. - if the
// [FileCoordinatorWritingForMerging] option is used, the
// [SavePresentedItemChangesWithCompletionHandler] method is called. If the
// target is part of a file package, the method is called on the package’s
// presenter, just as with coordinated reads.
// 
// For both reading and writing, if there are multiple file presenters
// involved, the order in which the methods are called is undefined. If any of
// the file presenters signal an error, the coordinated access fails and the
// error is passed to the accessor block.
//
// See: https://developer.apple.com/documentation/Foundation/NSFileCoordinator/coordinate(with:queue:byAccessor:)
func (f NSFileCoordinator) CoordinateAccessWithIntentsQueueByAccessor(intents []NSFileAccessIntent, queue INSOperationQueue, accessor ErrorHandler) {
_block2, _cleanup2 := NewErrorBlock(accessor)
	defer _cleanup2()
	objc.Send[objc.ID](f.ID, objc.Sel("coordinateAccessWithIntents:queue:byAccessor:"), intents, queue, _block2)
}

// Initiates a read operation on a single file or directory using the
// specified options.
//
// url: A URL identifying the file or directory to read. If other objects or
// processes are acting on the item at the URL, the actual URL passed to the
// `reader` parameter may be different than the one in this parameter.
//
// options: One of the reading options described in [NSFileCoordinator.ReadingOptions].
// If you pass no options, the [SavePresentedItemChangesWithCompletionHandler]
// method of relevant file presenters is called before your block executes.
// //
// [NSFileCoordinator.ReadingOptions]: https://developer.apple.com/documentation/Foundation/NSFileCoordinator/ReadingOptions
//
// outError: On input, a pointer to a pointer for an error object. If a file presenter
// encounters an error while preparing for this read operation, that error is
// returned in this parameter and the block in the `reader` parameter is not
// executed. If you cancel this operation before the `reader` block is
// executed, this parameter contains an error object on output.
//
// reader: A [Block object] containing the file operations you want to perform in a
// coordinated manner. This block receives an [NSURL] object containing the
// URL of the item and returns no value. Always use the URL passed into the
// block instead of the value in the `url` parameter.
// //
// [Block object]: https://developer.apple.com/library/archive/documentation/General/Conceptual/DevPedia-CocoaCore/Block.html#//apple_ref/doc/uid/TP40008195-CH3
//
// # Discussion
// 
// You use this method to perform read-related operations on a file or
// directory in a coordinated manner. This method executes synchronously,
// blocking the current thread until the `reader` block finishes executing.
// Before executing that block, though, the file coordinator waits until other
// relevant file presenter objects finish in-progress actions. Similarly, your
// read operation may cause pending actions for other file presenters to wait
// until your operations are complete. Whether or not the file coordinator
// waits depends on whether the item being read is a file or a directory and
// also depends on other related operations.
// 
// - If the `url` parameter specifies a file: - This method waits for other
// writers of the exact same file to finish in-progress actions. - This method
// waits if the file is a file package or any item inside the file package and
// other writers are writing to other items in the package directory. - This
// method does not wait for other readers of the file. - This method does not
// wait for writers that are manipulating the parent directory of the file,
// unless one of those writers specified the
// [FileCoordinatorWritingForDeleting] or [FileCoordinatorWritingForMoving]
// option. - If the `url` parameter specifies a directory: - This method waits
// if other write operations are occurring on the exact same directory. - This
// method does not wait if write operations are occurring on items inside the
// directory (but not on the directory itself). - This method does not wait
// for other readers of the directory. - This method does not wait for writers
// that are manipulating the parent directory of the directory, unless one of
// those writers specified the [FileCoordinatorWritingForDeleting] or
// [FileCoordinatorWritingForMoving] option.
// 
// When invoking these methods, declare a `__block` variable before the
// accessor block and initialize it to a value that signals failure, and then
// inside the accessor block set it to a value that indicates success. If the
// coordinated operation fails, then the accessor block never runs. The
// sentinel variable still holds a value that indicates failure, and the
// [NSError] out parameter contains a reference that describes the error.
// 
// This method calls the [RelinquishPresentedItemToReader] method of any
// relevant file presenters. This method is called for file presenters in the
// current process and in other processes. Depending on the options you
// specify, other methods of the file presenters may also be called. When
// reading a file package directory, file presenter objects that are currently
// reading the contents of that file package also receive these notifications.
// All of the called methods must return successfully before the file
// coordinator executes your block. If multiple file presenters are operating
// on the item, the order in which those presenters are notified is undefined.
// 
// If the device has not yet downloaded the file at the given URL, this method
// blocks (potentially for a long time) while the file is downloaded. If the
// file cannot be downloaded, this method fails. Alternatively; use a metadata
// query to check for the [NSMetadataUbiquitousItemDownloadingStatusKey] key,
// and then call the [StartDownloadingUbiquitousItemAtURLError] method to
// download the file before trying to read it.
// 
// If you want to perform a write operation from inside a read block, use the
// [CoordinateWritingItemAtURLOptionsWritingItemAtURLOptionsErrorByAccessor]
// method.
// 
// If you want to perform a batch read operation on multiple files, use the
// [PrepareForReadingItemsAtURLsOptionsWritingItemsAtURLsOptionsErrorByAccessor]
// method instead.
//
// [NSMetadataUbiquitousItemDownloadingStatusKey]: https://developer.apple.com/documentation/Foundation/NSMetadataUbiquitousItemDownloadingStatusKey
//
// See: https://developer.apple.com/documentation/Foundation/NSFileCoordinator/coordinate(readingItemAt:options:error:byAccessor:)
func (f NSFileCoordinator) CoordinateReadingItemAtURLOptionsErrorByAccessor(url INSURL, options NSFileCoordinatorReadingOptions, outError INSError, reader URLHandler) {
_block3, _cleanup3 := NewURLBlock(reader)
	defer _cleanup3()
	objc.Send[objc.ID](f.ID, objc.Sel("coordinateReadingItemAtURL:options:error:byAccessor:"), url, options, outError, _block3)
}

// Initiates a write operation on a single file or directory using the
// specified options.
//
// url: A URL identifying the file or directory to write to. If other objects or
// processes are acting on the item at the URL, the actual URL passed to
// `writer` parameter may be different from the one in this parameter.
//
// options: One of the writing options described in [NSFileCoordinator.WritingOptions].
// The options you specify partially determine how file presenters are
// notified and how this file coordinator object waits to execute your block.
// //
// [NSFileCoordinator.WritingOptions]: https://developer.apple.com/documentation/Foundation/NSFileCoordinator/WritingOptions
//
// outError: On input, a pointer to a pointer for an error object. If a file presenter
// encounters an error while preparing for this write operation, that error is
// returned in this parameter and the block in the `writer` parameter is not
// executed. If you cancel this operation before the `writer` block is
// executed, this parameter contains an error object on output.
//
// writer: A [Block object] containing the file operations you want to perform in a
// coordinated manner. This block receives an [NSURL] object containing the
// URL of the item and returns no value. Always use the URL passed into the
// block instead of the value in the `url` parameter.
// //
// [Block object]: https://developer.apple.com/library/archive/documentation/General/Conceptual/DevPedia-CocoaCore/Block.html#//apple_ref/doc/uid/TP40008195-CH3
//
// # Discussion
// 
// When invoking these methods, declare a `__block` variable before the
// accessor block and initialize it to a value that signals failure, and then
// inside the accessor block set it to a value that indicates success. If the
// coordinated operation fails, then the accessor block never runs. The
// sentinel variable still holds a value that indicates failure, and the
// [NSError] out parameter contains a reference that describes the error.
// 
// You use this method to perform write-related operations on a file or
// directory in a coordinated manner. This method executes synchronously,
// blocking the current thread until the `writer` block finishes executing.
// Before executing the block, though, the file coordinator waits until other
// relevant file presenter objects finish in-progress actions. Similarly, your
// write operation may cause pending actions for other file presenters to wait
// until your operations are complete. Whether or not the file coordinator
// waits depends on whether the item being written is a file or directory and
// also depends on other related operations.
// 
// - If the `url` parameter specifies a file: - This method waits for other
// readers and writers of the exact same file to finish in-progress actions. -
// This method waits if the file is a file package or any item inside a file
// package and other writers are reading or writing to other items inside the
// package directory. - This method does not wait for readers or writers that
// are manipulating the parent directory of the file, unless one of those
// writers specified the [FileCoordinatorWritingForDeleting] or
// [FileCoordinatorWritingForMoving] option. - If the `url` parameter
// specifies a directory: - This method waits if other read or write
// operations are occurring on the exact same directory. - This method does
// not wait if read or write operations are occurring on items inside the
// directory (but not on the directory itself). - This method does not wait
// for readers or writers that are manipulating the parent directory of the
// directory, unless one of those writers specified the
// [FileCoordinatorWritingForDeleting] or [FileCoordinatorWritingForMoving]
// option.
// 
// This method calls the [RelinquishPresentedItemToWriter] method of any
// relevant file presenters. This method is called for file presenters in the
// current process and in other processes. Depending on the options you
// specify, other methods of the file presenters may also be called. When
// writing a file package directory, file presenter objects that are currently
// reading or writing the contents of that file package also receive these
// notifications. All of the called methods must return successfully before
// the file coordinator executes your block. If multiple file presenters are
// operating on the item, the order in which those presenters are notified is
// undefined.
// 
// With one exception, do not nest calls to file coordinator methods inside
// the block you pass to this method. You may call the
// [CoordinateReadingItemAtURLOptionsErrorByAccessor] method to read the file
// if you discover through modification-date checking that the contents of the
// file have changed. However, if you call this method from inside your block,
// the file coordinator object throws an exception.
//
// See: https://developer.apple.com/documentation/Foundation/NSFileCoordinator/coordinate(writingItemAt:options:error:byAccessor:)
func (f NSFileCoordinator) CoordinateWritingItemAtURLOptionsErrorByAccessor(url INSURL, options NSFileCoordinatorWritingOptions, outError INSError, writer URLHandler) {
_block3, _cleanup3 := NewURLBlock(writer)
	defer _cleanup3()
	objc.Send[objc.ID](f.ID, objc.Sel("coordinateWritingItemAtURL:options:error:byAccessor:"), url, options, outError, _block3)
}

// Initiates a read operation that contains a follow-up write operation.
//
// readingURL: A URL identifying the file or directory to read. If other objects or
// processes are acting on the item at the URL, the actual URL passed to the
// block in the `readerWriter` parameter may be different than the one in this
// parameter.
//
// readingOptions: One of the reading options described in [NSFileCoordinator.ReadingOptions].
// If you pass `0` for this parameter, the
// [SavePresentedItemChangesWithCompletionHandler] method of relevant file
// presenters is called before your block executes.
// //
// [NSFileCoordinator.ReadingOptions]: https://developer.apple.com/documentation/Foundation/NSFileCoordinator/ReadingOptions
//
// writingURL: A URL identifying the file or directory to write. If other objects or
// processes are acting on the item at the URL, the actual URL passed to the
// block in the `readerWriter` parameter may be different than the one in this
// parameter.
//
// writingOptions: One of the writing options described in [NSFileCoordinator.WritingOptions].
// The options you specify partially determine how file presenters are
// notified and how this file coordinator object waits to execute your block.
// //
// [NSFileCoordinator.WritingOptions]: https://developer.apple.com/documentation/Foundation/NSFileCoordinator/WritingOptions
//
// outError: On input, a pointer to a pointer for an error object. If a file presenter
// encounters an error while preparing for this operation, that error is
// returned in this parameter and the block in the `readerWriter` parameter is
// not executed. If you cancel this operation before the `readerWriter` block
// is executed, this parameter contains an error object on output.
//
// readerWriter: A [Block object] containing the read and write operations you want to
// perform in a coordinated manner. This block receives [NSURL] objects
// containing the URLs of the items to read and write and returns no value.
// Always use the URLs passed into the block instead of the values in the
// `readingURL` and `writingURL` parameters.
// //
// [Block object]: https://developer.apple.com/library/archive/documentation/General/Conceptual/DevPedia-CocoaCore/Block.html#//apple_ref/doc/uid/TP40008195-CH3
//
// # Discussion
// 
// When invoking these methods, declare a `__block` variable before the
// accessor block and initialize it to a value that signals failure, and then
// inside the accessor block set it to a value that indicates success. If the
// coordinated operation fails, then the accessor block never runs. The
// sentinel variable still holds a value that indicates failure, and the
// [NSError] out parameter contains a reference that describes the error.
// 
// You use this method to perform a read operation that might also contain a
// write operation that needs to be coordinated. This method executes
// synchronously, blocking the current thread until the `readerWriter` block
// finishes executing. When performing the write operation, you may call the
// [CoordinateWritingItemAtURLOptionsErrorByAccessor] method from your
// `readerWriter` block. This method does the canonical lock ordering that is
// required to prevent a potential deadlock of the file operations.
// 
// This method makes the same calls to file presenters, and has the same
// general wait behavior, as the
// [CoordinateReadingItemAtURLOptionsErrorByAccessor] method.
//
// See: https://developer.apple.com/documentation/Foundation/NSFileCoordinator/coordinate(readingItemAt:options:writingItemAt:options:error:byAccessor:)
func (f NSFileCoordinator) CoordinateReadingItemAtURLOptionsWritingItemAtURLOptionsErrorByAccessor(readingURL INSURL, readingOptions NSFileCoordinatorReadingOptions, writingURL INSURL, writingOptions NSFileCoordinatorWritingOptions, outError INSError, readerWriter URLURLHandler) {
_block5, _cleanup5 := NewURLURLBlock(readerWriter)
	defer _cleanup5()
	objc.Send[objc.ID](f.ID, objc.Sel("coordinateReadingItemAtURL:options:writingItemAtURL:options:error:byAccessor:"), readingURL, readingOptions, writingURL, writingOptions, outError, _block5)
}

// Initiates a write operation that involves a secondary write operation.
//
// url1: A URL identifying the first file or directory to write. If other objects or
// processes are acting on the item at the URL, the actual URL passed to the
// block in the `writer` parameter may be different from the one in this
// parameter.
//
// options1: One of the writing options described in [NSFileCoordinator.WritingOptions].
// //
// [NSFileCoordinator.WritingOptions]: https://developer.apple.com/documentation/Foundation/NSFileCoordinator/WritingOptions
//
// url2: A URL identifying the other file or directory to write. If other objects or
// processes are acting on the item at the URL, the actual URL passed to the
// block in the `writer` parameter may be different from the one in this
// parameter.
//
// options2: One of the writing options described in [NSFileCoordinator.WritingOptions].
// The options you specify partially determine how file presenters are
// notified and how this file coordinator object waits to execute your block.
// //
// [NSFileCoordinator.WritingOptions]: https://developer.apple.com/documentation/Foundation/NSFileCoordinator/WritingOptions
//
// outError: On input, a pointer to a pointer for an error object. If a file presenter
// encounters an error while preparing for this operation, that error is
// returned in this parameter and the block in the `writer` parameter is not
// executed. If you cancel this operation before the `writer` block is
// executed, this parameter contains an error object on output.
//
// writer: A [Block object] containing the write operations you want to perform in a
// coordinated manner. This block receives [NSURL] objects containing the URLs
// of the items to write and returns no value. Always use the URLs passed into
// the block instead of the values in the `url1` and `url2` parameters.
// //
// [Block object]: https://developer.apple.com/library/archive/documentation/General/Conceptual/DevPedia-CocoaCore/Block.html#//apple_ref/doc/uid/TP40008195-CH3
//
// # Discussion
// 
// When invoking these methods, declare a `__block` variable before the
// accessor block and initialize it to a value that signals failure, and then
// inside the accessor block set it to a value that indicates success. If the
// coordinated operation fails, then the accessor block never runs. The
// sentinel variable still holds a value that indicates failure, and the
// [NSError] out parameter contains a reference that describes the error.
// 
// You use this method to perform two write operations without the risk of
// those operations creating a deadlock. This method executes synchronously,
// blocking the current thread until the `writer` block finishes executing.
// You may call the [CoordinateWritingItemAtURLOptionsErrorByAccessor] method
// from your `writer` block. This method does the canonical lock ordering that
// is required to prevent a potential deadlock of the file operations.
// 
// This method makes the same calls to file presenters, and has the same
// general wait behavior, as the
// [CoordinateWritingItemAtURLOptionsErrorByAccessor] method.
//
// See: https://developer.apple.com/documentation/Foundation/NSFileCoordinator/coordinate(writingItemAt:options:writingItemAt:options:error:byAccessor:)
func (f NSFileCoordinator) CoordinateWritingItemAtURLOptionsWritingItemAtURLOptionsErrorByAccessor(url1 INSURL, options1 NSFileCoordinatorWritingOptions, url2 INSURL, options2 NSFileCoordinatorWritingOptions, outError INSError, writer URLURLHandler) {
_block5, _cleanup5 := NewURLURLBlock(writer)
	defer _cleanup5()
	objc.Send[objc.ID](f.ID, objc.Sel("coordinateWritingItemAtURL:options:writingItemAtURL:options:error:byAccessor:"), url1, options1, url2, options2, outError, _block5)
}

// Prepare to read or write from multiple files in a single batch operation.
//
// readingURLs: An array of [NSURL] objects identifying the items you want to read.
//
// readingOptions: One of the reading options described in [NSFileCoordinator.ReadingOptions].
// If you pass `0` for this parameter, the
// [SavePresentedItemChangesWithCompletionHandler] method of relevant file
// presenters is called before your block executes.
// //
// [NSFileCoordinator.ReadingOptions]: https://developer.apple.com/documentation/Foundation/NSFileCoordinator/ReadingOptions
//
// writingURLs: An array of [NSURL] objects identifying the items you want to write.
//
// writingOptions: One of the writing options described in [NSFileCoordinator.WritingOptions].
// The options you specify partially determine how file presenters are
// notified and how this file coordinator object waits to execute your block.
// //
// [NSFileCoordinator.WritingOptions]: https://developer.apple.com/documentation/Foundation/NSFileCoordinator/WritingOptions
//
// outError: On input, a pointer to a pointer for an error object. If a file presenter
// encounters an error while preparing for this operation, that error is
// returned in this parameter and the block in the `writer` parameter is not
// executed. If you cancel this operation before the `batchAccessor` block is
// executed, this parameter contains an error object on output.
//
// batchAccessor: A [Block object] containing additional calls to methods of this class.
// 
// The block takes the following parameter:
// 
// completionHandler: A completion handler block. The batch accessor must call
// the completion handler when it has finished its read and write calls.
// //
// [Block object]: https://developer.apple.com/library/archive/documentation/General/Conceptual/DevPedia-CocoaCore/Block.html#//apple_ref/doc/uid/TP40008195-CH3
//
// # Discussion
// 
// Use this method to prepare the file coordinator for multiple read and write
// operations. Because file coordination requires interprocess communication,
// it is much more efficient to batch changes to large numbers of files and
// directories than to change each item individually. The file coordinator
// uses the values in the `readingURLs` and `writingURLs` parameters, together
// with reading and writing options, to prepare any relevant file presenters
// for the upcoming operations. Specifically, it uses these parameters in the
// same way as the [CoordinateReadingItemAtURLOptionsErrorByAccessor] and
// [CoordinateWritingItemAtURLOptionsErrorByAccessor] methods to determine
// which file presenter methods to call.
// 
// This method executes synchronously, blocking the current thread until the
// `batchAccessor` block finishes executing. The block you provide for the
// `batchAccessor` parameter does not perform the actual operations itself.
// Instead, you must call the individual coordinated read and write methods
// from inside the `batchAccessor` block. You must then call the completion
// handler after all the coordinated reads and writes have completed. You can
// call the completion handler from any thread.
// 
// Don’t simply pass this method all the URLs that are passed into the
// nested coordinate methods. Instead pass only the top-level files and
// directories involved in the operation. This method triggers messages to the
// file presenters of those items and to the file presenters of any items
// contained by those items.
// 
// In most cases, passing the same reading and writing options to both this
// method and the contained coordination operations is redundant. For example,
// it is often appropriate to pass [FileCoordinatorReadingWithoutChanges] to
// nested read operations. This method has already triggered a call to
// [SavePresentedItemChangesWithCompletionHandler]. The individual read
// operations do not need to trigger additional calls.
//
// See: https://developer.apple.com/documentation/Foundation/NSFileCoordinator/prepare(forReadingItemsAt:options:writingItemsAt:options:error:byAccessor:)
func (f NSFileCoordinator) PrepareForReadingItemsAtURLsOptionsWritingItemsAtURLsOptionsErrorByAccessor(readingURLs []NSURL, readingOptions NSFileCoordinatorReadingOptions, writingURLs []NSURL, writingOptions NSFileCoordinatorWritingOptions, outError INSError, batchAccessor VoidHandler) {
_block5, _cleanup5 := NewVoidBlock(batchAccessor)
	defer _cleanup5()
	objc.Send[objc.ID](f.ID, objc.Sel("prepareForReadingItemsAtURLs:options:writingItemsAtURLs:options:error:byAccessor:"), readingURLs, readingOptions, writingURLs, writingOptions, outError, _block5)
}

// Announces that your app is moving a file to a new URL.
//
// oldURL: The old location of the file or directory.
//
// newURL: The new location of the file or directory.
//
// # Discussion
// 
// This method is intended for apps that adopt App Sandbox.
// 
// Some apps need to rename files while saving them. For example, when a user
// adds an attachment to a rich text document, TextEdit changes the
// document’s filename extension from `XCUIElementTypeRtf` to
// `XCUIElementTypeRtfd`. In such a case, in a sandboxed app, you must call
// this method to declare your intent to rename a file without user approval.
// 
// After the renaming process succeeds, call the [ItemAtURLDidMoveToURL]
// method, with the same arguments, to provide your app with continued access
// to the file under its new name, while also giving up access to any file
// that appears with the old name.
// 
// If your macOS app is not sandboxed, this method serves no purpose. This
// method is nonfunctional in iOS.
//
// See: https://developer.apple.com/documentation/Foundation/NSFileCoordinator/item(at:willMoveTo:)
func (f NSFileCoordinator) ItemAtURLWillMoveToURL(oldURL INSURL, newURL INSURL) {
	objc.Send[objc.ID](f.ID, objc.Sel("itemAtURL:willMoveToURL:"), oldURL, newURL)
}

// Notifies relevant file presenters that the location of a file or directory
// changed.
//
// oldURL: The old location of the file or directory.
//
// newURL: The new location of the file or directory.
//
// # Discussion
// 
// If you move or rename a file or directory as part of a write operation,
// call this method to notify relevant file presenters that the change
// occurred. This method calls the [PresentedItemDidMoveToURL] method for any
// of the item’s file presenters. If the item is a directory, this method
// calls [PresentedItemDidMoveToURL] on the file presenters for the item’s
// contents. Finally, it calls [PresentedSubitemAtURLDidMoveToURL] on the file
// presenter of any directory containing the item.
// 
// You must call this method from a coordinated write block. Calling this
// method with the same URL in the `oldURL` and `newURL` parameters is
// harmless. This call must balance a call to [ItemAtURLWillMoveToURL].
//
// See: https://developer.apple.com/documentation/Foundation/NSFileCoordinator/item(at:didMoveTo:)
func (f NSFileCoordinator) ItemAtURLDidMoveToURL(oldURL INSURL, newURL INSURL) {
	objc.Send[objc.ID](f.ID, objc.Sel("itemAtURL:didMoveToURL:"), oldURL, newURL)
}

// Cancels any active file coordination calls.
//
// # Discussion
// 
// Use this method to cancel any active calls to coordinate the reading or
// writing of a file. If the block passed to the file coordination call has
// not yet been executed—perhaps because the file coordinator is still
// waiting for a response from other file presenters—the file coordinator
// method stops waiting for a response, does not execute its block, and
// returns an error object with the error code [NSUserCancelledError].
// However, if the block is already being executed, the file coordinator
// method does not return until the block finishes executing.
// 
// You can call this method from any thread of your application and it returns
// immediately without waiting for the file coordinator object to respond.
// Thus, when this method returns, you cannot assume that the read or write
// operation occurred or did not occur. (In fact, it is possible for this
// method to return while the file coordinator is in the middle of executing a
// block.) If you want to know whether the operation actually occurred, you
// must track it yourself by setting a flag when the block starts executing or
// by using some other means.
//
// [NSUserCancelledError]: https://developer.apple.com/documentation/Foundation/NSUserCancelledError-swift.var
//
// See: https://developer.apple.com/documentation/Foundation/NSFileCoordinator/cancel()
func (f NSFileCoordinator) Cancel() {
	objc.Send[objc.ID](f.ID, objc.Sel("cancel"))
}

// Tells observing file providers that the item’s ubiquity attributes have
// changed.
//
// # Discussion
// 
// This method triggers the [NSFilePresenter] protocol’s
// [PresentedItemDidChangeUbiquityAttributes] method on any file presenters
// that are observing the item, even if they are running in different
// processes.
// 
// For information about the types of attributes that can trigger
// notifications, see the [NSFilePresenter] protocol’s
// [ObservedPresentedItemUbiquityAttributes] property.
//
// [NSFilePresenter]: https://developer.apple.com/library/archive/releasenotes/Foundation/RN-FoundationOlderNotes/index.html#//apple_ref/doc/uid/TP40008080-TRANSLATED_CHAPTER_965-TRANSLATED_DEST_6
//
// See: https://developer.apple.com/documentation/Foundation/NSFileCoordinator/item(at:didChangeUbiquityAttributes:)
func (f NSFileCoordinator) ItemAtURLDidChangeUbiquityAttributes(url INSURL, attributes INSSet) {
	objc.Send[objc.ID](f.ID, objc.Sel("itemAtURL:didChangeUbiquityAttributes:"), url, attributes)
}

// Registers the specified file presenter object so that it can receive
// notifications.
//
// filePresenter: The file presenter object to register.
//
// # Discussion
// 
// This method registers the file presenter object process wide. Thus, any
// file coordinator objects you create later automatically know about the file
// presenter object and know to message it when its file or directory is
// affected.
// 
// Be sure to balance calls to this method with a corresponding call to the
// [RemoveFilePresenter] method. You must remove file presenters from the
// process wide registry before the object is deallocated, even in a
// garbage-collected application.
// 
// If you call this method while coordinated file operations are already under
// way in another process, your file presenter may not receive notifications
// for that operation. To prevent missing such notifications, create a file
// coordinator, call its [CoordinateReadingItemAtURLOptionsErrorByAccessor]
// method, and register your file presenter object there. If you are going to
// read a file and then create a file presenter for that file, both actions
// should occur in the same coordinated read block. Synchronizing on the
// presented file or directory guarantees that when your block executes, all
// other objects have completed any tasks and you have sole access to the
// item.
//
// See: https://developer.apple.com/documentation/Foundation/NSFileCoordinator/addFilePresenter(_:)
func (_NSFileCoordinatorClass NSFileCoordinatorClass) AddFilePresenter(filePresenter NSFilePresenter) {
	objc.Send[objc.ID](objc.ID(_NSFileCoordinatorClass.class), objc.Sel("addFilePresenter:"), filePresenter)
}

// Unregisters the specified file presenter object.
//
// filePresenter: The file presenter object to unregister. If the object is not currently
// registered, this method does nothing.
//
// # Discussion
// 
// Call this method to unregister file presenters before those objects are
// deallocated, even in a garbage-collected application.
//
// See: https://developer.apple.com/documentation/Foundation/NSFileCoordinator/removeFilePresenter(_:)
func (_NSFileCoordinatorClass NSFileCoordinatorClass) RemoveFilePresenter(filePresenter NSFilePresenter) {
	objc.Send[objc.ID](objc.ID(_NSFileCoordinatorClass.class), objc.Sel("removeFilePresenter:"), filePresenter)
}

// A string that uniquely identifies the file access that was performed by
// this file coordinator.
//
// # Discussion
// 
// Coordinated reads and writes performed using the same purpose identifier
// never block each other, even if they occur in different processes. If you
// are coordinating file access on behalf of a file presenter, use
// [InitWithFilePresenter] and do not attempt to set a custom purpose
// identifier. Every file coordinator instance initialized with the same file
// presenter has the same purpose identifier.
// 
// You may need to set a custom purpose identifier for the following reasons:
// 
// - Your application has a File Provider extension. Any file coordination
// done on behalf of the File Provider needs to be done using the File
// Provider’s purpose identifier. - You have two separate subsystems that
// need to work together to perform a single high-level operation, and both
// subsystems perform their own coordinated reads or writes. Using the same
// purpose identifier in both subsystems prevents possible deadlocks between
// the two subsystems.
// 
// When creating custom purpose identifiers, you can use a reverse DNS style
// string, such as `com.Example().MyApplication.MyPurpose`, or a UUID string.
// You cannot use `nil` or zero-length strings.
//
// See: https://developer.apple.com/documentation/Foundation/NSFileCoordinator/purposeIdentifier
func (f NSFileCoordinator) PurposeIdentifier() string {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("purposeIdentifier"))
	return NSStringFromID(rv).String()
}
func (f NSFileCoordinator) SetPurposeIdentifier(value string) {
	objc.Send[struct{}](f.ID, objc.Sel("setPurposeIdentifier:"), objc.String(value))
}

// The user canceled the operation (for example, by pressing Command-period).
//
// See: https://developer.apple.com/documentation/foundation/nsusercancellederror-swift.var
func (f NSFileCoordinator) NSUserCancelledError() int {
	rv := objc.Send[int](f.ID, objc.Sel("NSUserCancelledError"))
	return rv
}
func (f NSFileCoordinator) SetNSUserCancelledError(value int) {
	objc.Send[struct{}](f.ID, objc.Sel("setNSUserCancelledError:"), value)
}

// Returns an array containing the currently registered file presenter
// objects.
//
// # Return Value
// 
// An array of objects that conform to the [NSFilePresenter] protocol.
//
// See: https://developer.apple.com/documentation/Foundation/NSFileCoordinator/filePresenters
func (_NSFileCoordinatorClass NSFileCoordinatorClass) FilePresenters() []objectivec.IObject {
	rv := objc.Send[[]objc.ID](objc.ID(_NSFileCoordinatorClass.class), objc.Sel("filePresenters"))
	return objc.ConvertSlice(rv, func(id objc.ID) objectivec.IObject {
		return objectivec.Object{ID: id}
	})
}

// A notification that posts when the app enters the background.
//
// See: https://developer.apple.com/documentation/UIKit/UIApplication/didEnterBackgroundNotification
func (_NSFileCoordinatorClass NSFileCoordinatorClass) DidEnterBackgroundNotification() NSNotificationName {
	rv := objc.Send[objc.ID](objc.ID(_NSFileCoordinatorClass.class), objc.Sel("didEnterBackgroundNotification"))
	return NSNotificationName(NSStringFromID(rv).String())
}

// A notification that posts shortly before an app leaves the background state
// on its way to becoming the active app.
//
// See: https://developer.apple.com/documentation/UIKit/UIApplication/willEnterForegroundNotification
func (_NSFileCoordinatorClass NSFileCoordinatorClass) WillEnterForegroundNotification() NSNotificationName {
	rv := objc.Send[objc.ID](objc.ID(_NSFileCoordinatorClass.class), objc.Sel("willEnterForegroundNotification"))
	return NSNotificationName(NSStringFromID(rv).String())
}

// CoordinateReadingItemAtURLOptionsErrorByAccessorSync is a synchronous wrapper around [NSFileCoordinator.CoordinateReadingItemAtURLOptionsErrorByAccessor].
// It blocks until the completion handler fires or the context is cancelled.
func (f NSFileCoordinator) CoordinateReadingItemAtURLOptionsErrorByAccessorSync(ctx context.Context, url INSURL, options NSFileCoordinatorReadingOptions, outError INSError) (*NSURL, error) {
	done := make(chan *NSURL, 1)
	f.CoordinateReadingItemAtURLOptionsErrorByAccessor(url, options, outError, func(val *NSURL) {
		done <- val
	})
	select {
	case r := <-done:
		return r, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

// CoordinateWritingItemAtURLOptionsErrorByAccessorSync is a synchronous wrapper around [NSFileCoordinator.CoordinateWritingItemAtURLOptionsErrorByAccessor].
// It blocks until the completion handler fires or the context is cancelled.
func (f NSFileCoordinator) CoordinateWritingItemAtURLOptionsErrorByAccessorSync(ctx context.Context, url INSURL, options NSFileCoordinatorWritingOptions, outError INSError) (*NSURL, error) {
	done := make(chan *NSURL, 1)
	f.CoordinateWritingItemAtURLOptionsErrorByAccessor(url, options, outError, func(val *NSURL) {
		done <- val
	})
	select {
	case r := <-done:
		return r, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

