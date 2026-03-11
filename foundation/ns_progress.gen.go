// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"context"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [Progress] class.
var (
	_ProgressClass     ProgressClass
	_ProgressClassOnce sync.Once
)

func getProgressClass() ProgressClass {
	_ProgressClassOnce.Do(func() {
		_ProgressClass = ProgressClass{class: objc.GetClass("NSProgress")}
	})
	return _ProgressClass
}

// GetProgressClass returns the class object for NSProgress.
func GetProgressClass() ProgressClass {
	return getProgressClass()
}

type ProgressClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (pc ProgressClass) Alloc() Progress {
	rv := objc.Send[Progress](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}







// An object that conveys ongoing progress to the user for a specified task.
//
// # Overview
// 
// The [NSProgress] class provides a self-contained mechanism for progress
// reporting. It makes it easy for code that performs work to report the
// progress of that work, and for user interface code to observe that progress
// for presentation to the user. Specifically, you can use a progress object
// to show the user a progress bar and explanatory text that update as you do
// work. It also allows the user to cancel or pause work.
// 
// # Reporting Progress
// 
// Using the methods of this class, your code can report the progress it’s
// currently making toward completing a task, including progress in related
// subtasks. You can create instances of this class using the
// [InitWithParentUserInfo] instance method or the
// [ProgressWithTotalUnitCount] class method.
// 
// Progress objects have many properties that you can use to observe and
// report current progress. For instance, the [TotalUnitCount] property
// represents the total number of units of work, and the [CompletedUnitCount]
// and [FractionCompleted] properties represent how much of that work is
// complete. The [FractionCompleted] property is useful for updating progress
// indicators or textual descriptors. To check whether progress is complete,
// test the [Finished] property.
// 
// The following code shows a sample method that reports the progress of
// performing an operation on a piece of data. When creating the progress
// object, you set the value of its [TotalUnitCount] property to a suitable
// batch size for the operation, and the [CompletedUnitCount] count is `0`.
// Each time the loop executes and processes a batch of data, you increment
// the progress object’s [CompletedUnitCount] property appropriately.
// 
// Each of the properties of a progress object, including [TotalUnitCount],
// [CompletedUnitCount], and [FractionCompleted], support key-value observing
// (KVO). This makes it extremely easy for a view or window controller object
// to observe the properties, and update UI elements, such as progress
// indicators, when the values change. It also means that there is a nonzero
// cost to updating the values of those properties, so avoid using a unit
// count that is too granular. If you’re iterating over a large dataset, for
// example, and each operation takes only a trivial amount of time, divide the
// work into batches so you can update the unit count once per batch rather
// than once per iteration.
// 
// # Reporting Progress for Multiple Operations
// 
// Sometimes, your code may need to report the progress of an operation that
// consists of several suboperations. To accomplish this, your code can report
// the progress of each suboperation by building up a tree of progress
// objects.
// 
// The [NSProgress] reporting mechanism supports a loosely coupled
// relationship between progress objects. Suboperations don’t need to know
// anything about the containing progress item — you can create new progress
// objects as suboperations of another progress instance. When you assign a
// progress instance, the system allocates a portion of the containing
// progress instance’s pending unit count. When the suboperation’s
// progress object completes, the containing progress object’s
// [CompletedUnitCount] property automatically increases by a predefined
// amount.
// 
// You add suboperation progress objects to your tree implicitly or
// explicitly.
// 
// # Adding a Progress Operation Implicitly
// 
// Add a suboperation implicitly by setting a pending unit count for the
// containing progress object and creating a new [NSProgress] instance. When
// you create the new progress instance, the system sets it as a suboperation
// of the containing progress object, and assigns the pending unit count.
// 
// As an example, consider that you’re tracking the progress of code
// downloading and copying files on disk. You can use a single progress object
// to track the entire task, but it’s easier to manage each subtask using a
// separate progress object. You start by creating an overall progress object
// with a suitable total unit count, call [BecomeCurrentWithPendingUnitCount],
// then create your suboperation progress objects before finally calling
// [ResignCurrent].
// 
// The system divides the pending unit count that you specify in the first
// method equally among the suboperation progress objects you create between
// these two method calls. Each suboperation progress object maintains its own
// internal unit count. When the suboperation object’s [CompletedUnitCount]
// equals or exceeds its [TotalUnitCount], the system increases the containing
// progress object’s [CompletedUnitCount] by the assigned portion of the
// original pending unit count.
// 
// In the following example, the overall progress object has 100 units. The
// two suboperation objects, therefore, get 50 pending units each, and keep
// track internally of 10 units of work each. When each suboperation completes
// its 10 units, the system increases the overall progress object’s
// completed unit count by 50.
// 
// If you don’t create any suboperation progress objects between the calls
// to [BecomeCurrentWithPendingUnitCount] and [ResignCurrent], the containing
// progress object automatically updates its [CompletedUnitCount] by adding
// the pending units.
// 
// # Adding a Progress Operation Explicitly
// 
// To add a progress operation explicitly, call [AddChildWithPendingUnitCount]
// on the containing progress object. The value for the pending unit count is
// the amount of the containing progress object’s [TotalUnitCount] that the
// suboperation consumes, which conforms to the [NSProgressReporting]
// protocol.
// 
// In the following example, the overall progress object has 10 units. The
// suboperation progress for the download gets eight units and tracks the
// download of a photo. The progress for the filter takes a lot less time and
// gets the remaining two units. When the download completes, the system
// updates the containing progress object’s completed unit count by eight.
// When the filter completes, the system updates it by the remaining two
// units.
//
// # Creating Progress Objects
//
//   - [Progress.InitWithParentUserInfo]: Creates a new progress instance.
//
// # Accessing the Current Progress Object
//
//   - [Progress.BecomeCurrentWithPendingUnitCount]: Sets the progress object as the current object of the current thread, and assigns the amount of work for the next suboperation progress object to perform.
//   - [Progress.AddChildWithPendingUnitCount]: Adds a process object as a suboperation of a progress tree.
//   - [Progress.ResignCurrent]: Restores the previous progress object to become the current progress object on the thread.
//
// # Reporting Progress
//
//   - [Progress.TotalUnitCount]: The total number of tracked units of work for the current progress.
//   - [Progress.SetTotalUnitCount]
//   - [Progress.CompletedUnitCount]: The number of completed units of work for the current job.
//   - [Progress.SetCompletedUnitCount]
//   - [Progress.LocalizedDescription]: A localized description of tracked progress for the receiver.
//   - [Progress.SetLocalizedDescription]
//   - [Progress.LocalizedAdditionalDescription]: A more specific localized description of tracked progress for the receiver.
//   - [Progress.SetLocalizedAdditionalDescription]
//   - [Progress.Cancellable]: A Boolean value that indicates whether the receiver is tracking work that you can cancel.
//   - [Progress.SetCancellable]
//   - [Progress.Cancelled]: A Boolean value that Indicates whether the receiver is tracking canceled work.
//   - [Progress.CancellationHandler]: The block to invoke when canceling progress.
//   - [Progress.SetCancellationHandler]
//   - [Progress.Pausable]: A Boolean value that indicates whether the receiver is tracking work that you can pause.
//   - [Progress.SetPausable]
//   - [Progress.Paused]: A Boolean value that indicates whether the receiver is tracking paused work.
//   - [Progress.PausingHandler]: The block to invoke when pausing progress.
//   - [Progress.SetPausingHandler]
//
// # Observing Progress
//
//   - [Progress.Indeterminate]: A Boolean value that indicates whether the tracked progress is indeterminate.
//   - [Progress.FractionCompleted]: The fraction of the overall work that the progress object completes, including work from its suboperations.
//   - [Progress.Finished]: A Boolean value that indicates the progress object is complete.
//
// # Controlling Progress
//
//   - [Progress.Cancel]: Cancels progress tracking.
//   - [Progress.Pause]: Pauses progress tracking.
//   - [Progress.Resume]: Resumes progress tracking.
//   - [Progress.ResumingHandler]: The block to invoke when progress resumes.
//   - [Progress.SetResumingHandler]
//
// # Inspecting Progress Information
//
//   - [Progress.Kind]: An object that represents the kind of progress for the progress object.
//   - [Progress.SetKind]
//   - [Progress.SetUserInfoObjectForKey]: Sets a value in the user info dictionary.
//   - [Progress.UserInfo]: A dictionary of arbitrary values for the receiver.
//
// # Inspecting File Operation Progress Information
//
//   - [Progress.FileOperationKind]: The kind of file operation for the progress object.
//   - [Progress.SetFileOperationKind]
//   - [Progress.FileURL]: A URL that represents the file for the current progress object.
//   - [Progress.SetFileURL]
//
// # Reporting Progress to Other Processes
//
//   - [Progress.Publish]: Publishes the progress object for other processes to observe it.
//   - [Progress.Unpublish]: Removes a progress object from publication, making it unobservable by other processes.
//
// # Observing and Controlling File Progress by Other Processes
//
//   - [Progress.Old]: A Boolean value that indicates when the observed progress object invokes the publish method before you subscribe to it.
//
// See: https://developer.apple.com/documentation/Foundation/Progress
type Progress struct {
	objectivec.Object
}

// ProgressFromID constructs a [Progress] from an objc.ID.
//
// An object that conveys ongoing progress to the user for a specified task.
func ProgressFromID(id objc.ID) Progress {
	return NSProgress{objectivec.Object{id}}
}

// NSProgressFromID is an alias for [ProgressFromID] for cross-framework compatibility.
func NSProgressFromID(id objc.ID) Progress { return ProgressFromID(id) }
// NOTE: Progress adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [Progress] class.
//
// # Creating Progress Objects
//
//   - [IProgress.InitWithParentUserInfo]: Creates a new progress instance.
//
// # Accessing the Current Progress Object
//
//   - [IProgress.BecomeCurrentWithPendingUnitCount]: Sets the progress object as the current object of the current thread, and assigns the amount of work for the next suboperation progress object to perform.
//   - [IProgress.AddChildWithPendingUnitCount]: Adds a process object as a suboperation of a progress tree.
//   - [IProgress.ResignCurrent]: Restores the previous progress object to become the current progress object on the thread.
//
// # Reporting Progress
//
//   - [IProgress.TotalUnitCount]: The total number of tracked units of work for the current progress.
//   - [IProgress.SetTotalUnitCount]
//   - [IProgress.CompletedUnitCount]: The number of completed units of work for the current job.
//   - [IProgress.SetCompletedUnitCount]
//   - [IProgress.LocalizedDescription]: A localized description of tracked progress for the receiver.
//   - [IProgress.SetLocalizedDescription]
//   - [IProgress.LocalizedAdditionalDescription]: A more specific localized description of tracked progress for the receiver.
//   - [IProgress.SetLocalizedAdditionalDescription]
//   - [IProgress.Cancellable]: A Boolean value that indicates whether the receiver is tracking work that you can cancel.
//   - [IProgress.SetCancellable]
//   - [IProgress.Cancelled]: A Boolean value that Indicates whether the receiver is tracking canceled work.
//   - [IProgress.CancellationHandler]: The block to invoke when canceling progress.
//   - [IProgress.SetCancellationHandler]
//   - [IProgress.Pausable]: A Boolean value that indicates whether the receiver is tracking work that you can pause.
//   - [IProgress.SetPausable]
//   - [IProgress.Paused]: A Boolean value that indicates whether the receiver is tracking paused work.
//   - [IProgress.PausingHandler]: The block to invoke when pausing progress.
//   - [IProgress.SetPausingHandler]
//
// # Observing Progress
//
//   - [IProgress.Indeterminate]: A Boolean value that indicates whether the tracked progress is indeterminate.
//   - [IProgress.FractionCompleted]: The fraction of the overall work that the progress object completes, including work from its suboperations.
//   - [IProgress.Finished]: A Boolean value that indicates the progress object is complete.
//
// # Controlling Progress
//
//   - [IProgress.Cancel]: Cancels progress tracking.
//   - [IProgress.Pause]: Pauses progress tracking.
//   - [IProgress.Resume]: Resumes progress tracking.
//   - [IProgress.ResumingHandler]: The block to invoke when progress resumes.
//   - [IProgress.SetResumingHandler]
//
// # Inspecting Progress Information
//
//   - [IProgress.Kind]: An object that represents the kind of progress for the progress object.
//   - [IProgress.SetKind]
//   - [IProgress.SetUserInfoObjectForKey]: Sets a value in the user info dictionary.
//   - [IProgress.UserInfo]: A dictionary of arbitrary values for the receiver.
//
// # Inspecting File Operation Progress Information
//
//   - [IProgress.FileOperationKind]: The kind of file operation for the progress object.
//   - [IProgress.SetFileOperationKind]
//   - [IProgress.FileURL]: A URL that represents the file for the current progress object.
//   - [IProgress.SetFileURL]
//
// # Reporting Progress to Other Processes
//
//   - [IProgress.Publish]: Publishes the progress object for other processes to observe it.
//   - [IProgress.Unpublish]: Removes a progress object from publication, making it unobservable by other processes.
//
// # Observing and Controlling File Progress by Other Processes
//
//   - [IProgress.Old]: A Boolean value that indicates when the observed progress object invokes the publish method before you subscribe to it.
//
// See: https://developer.apple.com/documentation/Foundation/Progress
type IProgress interface {
	objectivec.IObject

	// Topic: Creating Progress Objects

	// Creates a new progress instance.
	InitWithParentUserInfo(parentProgressOrNil INSProgress, userInfoOrNil INSDictionary) Progress

	// Topic: Accessing the Current Progress Object

	// Sets the progress object as the current object of the current thread, and assigns the amount of work for the next suboperation progress object to perform.
	BecomeCurrentWithPendingUnitCount(unitCount int64)
	// Adds a process object as a suboperation of a progress tree.
	AddChildWithPendingUnitCount(child INSProgress, inUnitCount int64)
	// Restores the previous progress object to become the current progress object on the thread.
	ResignCurrent()

	// Topic: Reporting Progress

	// The total number of tracked units of work for the current progress.
	TotalUnitCount() int64
	SetTotalUnitCount(value int64)
	// The number of completed units of work for the current job.
	CompletedUnitCount() int64
	SetCompletedUnitCount(value int64)
	// A localized description of tracked progress for the receiver.
	LocalizedDescription() string
	SetLocalizedDescription(value string)
	// A more specific localized description of tracked progress for the receiver.
	LocalizedAdditionalDescription() string
	SetLocalizedAdditionalDescription(value string)
	// A Boolean value that indicates whether the receiver is tracking work that you can cancel.
	Cancellable() bool
	SetCancellable(value bool)
	// A Boolean value that Indicates whether the receiver is tracking canceled work.
	Cancelled() bool
	// The block to invoke when canceling progress.
	CancellationHandler() VoidHandler
	SetCancellationHandler(value VoidHandler)
	// A Boolean value that indicates whether the receiver is tracking work that you can pause.
	Pausable() bool
	SetPausable(value bool)
	// A Boolean value that indicates whether the receiver is tracking paused work.
	Paused() bool
	// The block to invoke when pausing progress.
	PausingHandler() VoidHandler
	SetPausingHandler(value VoidHandler)

	// Topic: Observing Progress

	// A Boolean value that indicates whether the tracked progress is indeterminate.
	Indeterminate() bool
	// The fraction of the overall work that the progress object completes, including work from its suboperations.
	FractionCompleted() float64
	// A Boolean value that indicates the progress object is complete.
	Finished() bool

	// Topic: Controlling Progress

	// Cancels progress tracking.
	Cancel()
	// Pauses progress tracking.
	Pause()
	// Resumes progress tracking.
	Resume()
	// The block to invoke when progress resumes.
	ResumingHandler() VoidHandler
	SetResumingHandler(value VoidHandler)

	// Topic: Inspecting Progress Information

	// An object that represents the kind of progress for the progress object.
	Kind() NSProgressKind
	SetKind(value NSProgressKind)
	// Sets a value in the user info dictionary.
	SetUserInfoObjectForKey(objectOrNil objectivec.IObject, key NSProgressUserInfoKey)
	// A dictionary of arbitrary values for the receiver.
	UserInfo() INSDictionary

	// Topic: Inspecting File Operation Progress Information

	// The kind of file operation for the progress object.
	FileOperationKind() NSProgressFileOperationKind
	SetFileOperationKind(value NSProgressFileOperationKind)
	// A URL that represents the file for the current progress object.
	FileURL() INSURL
	SetFileURL(value INSURL)

	// Topic: Reporting Progress to Other Processes

	// Publishes the progress object for other processes to observe it.
	Publish()
	// Removes a progress object from publication, making it unobservable by other processes.
	Unpublish()

	// Topic: Observing and Controlling File Progress by Other Processes

	// A Boolean value that indicates when the observed progress object invokes the publish method before you subscribe to it.
	Old() bool

	// A value that indicates the estimated amount of time remaining to complete the progress.
	EstimatedTimeRemaining() INSNumber
	SetEstimatedTimeRemaining(value INSNumber)
	// The number of completed files for a file progress object.
	FileCompletedCount() INSNumber
	SetFileCompletedCount(value INSNumber)
	// The total number of files for a file progress object.
	FileTotalCount() INSNumber
	SetFileTotalCount(value INSNumber)
	// A value that represents the speed of data processing, in bytes per second.
	Throughput() INSNumber
	SetThroughput(value INSNumber)
	// Retrieves the current thread’s progress object, executes the specified block, and increments the progress object by the specified units of work.
	PerformAsCurrentWithPendingUnitCountUsingBlock(unitCount int64, work VoidHandler)
}





// Init initializes the instance.
func (p Progress) Init() Progress {
	rv := objc.Send[Progress](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p Progress) Autorelease() Progress {
	rv := objc.Send[Progress](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewProgress creates a new Progress instance.
func NewProgress() Progress {
	class := getProgressClass()
	rv := objc.Send[Progress](objc.ID(class.class), objc.Sel("new"))
	return rv
}






// Creates a new progress instance.
//
// parentProgressOrNil: The containing [NSProgress] object, if any, to notify when reporting
// progress, or to consult when checking for cancellation.
// 
// The only valid values are [CurrentProgress] or `nil`.
//
// userInfoOrNil: The optional user information dictionary for the progress object.
//
// # Discussion
// 
// This is the designated initializer for the [NSProgress] class.
//
// See: https://developer.apple.com/documentation/Foundation/Progress/init(parent:userInfo:)
func NewProgressWithParentUserInfo(parentProgressOrNil INSProgress, userInfoOrNil INSDictionary) Progress {
	instance := getProgressClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithParent:userInfo:"), parentProgressOrNil, userInfoOrNil)
	return ProgressFromID(rv)
}


// Creates and returns a progress instance.
//
// unitCount: The total number of units of work to assign to the progress instance.
//
// # Discussion
// 
// If a current progress object exists, the initializer uses it to set the
// value of the [TotalUnitCount] property.
// 
// In many cases, you can precede code that does a substantial amount of work
// with an invocation of this method, then repeatedly set the
// [CompletedUnitCount] or [Cancelled] property in the loop that does the
// work.
// 
// You can invoke this method on one thread and then message the returned
// [NSProgress] on another thread. For example, you can capture the created
// progress instance in a block that you pass to [dispatch_async]. In that
// block, you can invoke methods like [BecomeCurrentWithPendingUnitCount] or
// [ResignCurrent], and set the [CompletedUnitCount] or [Cancelled] properties
// as your app finishes its work.
//
// [dispatch_async]: https://developer.apple.com/documentation/Dispatch/dispatch_async
//
// See: https://developer.apple.com/documentation/Foundation/Progress/init(totalUnitCount:)
func NewProgressWithTotalUnitCount(unitCount int64) Progress {
	rv := objc.Send[objc.ID](objc.ID(getProgressClass().class), objc.Sel("progressWithTotalUnitCount:"), unitCount)
	return ProgressFromID(rv)
}


// Creates a progress instance for the specified progress object with a unit
// count that’s a portion of the containing object’s total unit count.
//
// unitCount: The total number of units of work to assign to the progress instance.
//
// parent: The containing progress object for the created [NSProgress] object.
//
// portionOfParentTotalUnitCount: The unit count for the progress object.
//
// # Discussion
// 
// Use this method to initialize a progress object with a specified containing
// progress object and unit count.
// 
// In many cases, you can precede code that does a substantial amount of work
// with an invocation of this method, then repeatedly set the
// [CompletedUnitCount] or [Cancelled] property in the loop that does the
// work.
// 
// You can invoke this method on one thread and then message the returned
// [NSProgress] on another thread. For example, you can capture the created
// progress instance in a block that you pass to [dispatch_async]. In that
// block, you can invoke methods like [BecomeCurrentWithPendingUnitCount] or
// [ResignCurrent], and set the [CompletedUnitCount] or [Cancelled] properties
// as your app finishes its work.
//
// [dispatch_async]: https://developer.apple.com/documentation/Dispatch/dispatch_async
//
// See: https://developer.apple.com/documentation/Foundation/Progress/init(totalUnitCount:parent:pendingUnitCount:)
func NewProgressWithTotalUnitCountParentPendingUnitCount(unitCount int64, parent INSProgress, portionOfParentTotalUnitCount int64) Progress {
	rv := objc.Send[objc.ID](objc.ID(getProgressClass().class), objc.Sel("progressWithTotalUnitCount:parent:pendingUnitCount:"), unitCount, parent, portionOfParentTotalUnitCount)
	return ProgressFromID(rv)
}







// Creates a new progress instance.
//
// parentProgressOrNil: The containing [NSProgress] object, if any, to notify when reporting
// progress, or to consult when checking for cancellation.
// 
// The only valid values are [CurrentProgress] or `nil`.
//
// userInfoOrNil: The optional user information dictionary for the progress object.
//
// # Discussion
// 
// This is the designated initializer for the [NSProgress] class.
//
// See: https://developer.apple.com/documentation/Foundation/Progress/init(parent:userInfo:)
func (p Progress) InitWithParentUserInfo(parentProgressOrNil INSProgress, userInfoOrNil INSDictionary) Progress {
	rv := objc.Send[Progress](p.ID, objc.Sel("initWithParent:userInfo:"), parentProgressOrNil, userInfoOrNil)
	return rv
}

// Sets the progress object as the current object of the current thread, and
// assigns the amount of work for the next suboperation progress object to
// perform.
//
// unitCount: The number of units of work for the next progress object that initializes
// when you invoke [InitWithParentUserInfo] in the current thread with this
// progress object as the containing progress object.
// 
// The number represents the portion of work to perform in relation to the
// total number of units of work, which is the value of the progress
// object’s [TotalUnitCount] property. The units of work for this parameter
// must be the same units of work in the progress object’s [TotalUnitCount]
// property.
//
// # Discussion
// 
// Use this method to build a tree of progress objects, as [NSProgress]
// describes.
//
// See: https://developer.apple.com/documentation/Foundation/Progress/becomeCurrent(withPendingUnitCount:)
func (p Progress) BecomeCurrentWithPendingUnitCount(unitCount int64) {
	objc.Send[objc.ID](p.ID, objc.Sel("becomeCurrentWithPendingUnitCount:"), unitCount)
}

// Adds a process object as a suboperation of a progress tree.
//
// child: The progress instance to add to the progress tree.
//
// inUnitCount: The number of units of work for the new suboperation to complete.
//
// # Discussion
// 
// You assign the suboperation a portion of the receiver’s total unit count
// according to `inUnitCount`. For more information, see [NSProgress].
//
// See: https://developer.apple.com/documentation/Foundation/Progress/addChild(_:withPendingUnitCount:)
func (p Progress) AddChildWithPendingUnitCount(child INSProgress, inUnitCount int64) {
	objc.Send[objc.ID](p.ID, objc.Sel("addChild:withPendingUnitCount:"), child, inUnitCount)
}

// Restores the previous progress object to become the current progress object
// on the thread.
//
// # Discussion
// 
// This method restores the current progress object to what it was before
// invoking [BecomeCurrentWithPendingUnitCount].
// 
// Use this method after building your tree of progress objects, as
// [NSProgress] describes.
//
// See: https://developer.apple.com/documentation/Foundation/Progress/resignCurrent()
func (p Progress) ResignCurrent() {
	objc.Send[objc.ID](p.ID, objc.Sel("resignCurrent"))
}

// Cancels progress tracking.
//
// # Discussion
// 
// This method invokes the block for [CancellationHandler], if there is one,
// and ensures that any subsequent reads of the [Cancelled] property return
// [true].
// 
// If the receiver has suboperations, the system cancels their progress as
// well.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Foundation/Progress/cancel()
func (p Progress) Cancel() {
	objc.Send[objc.ID](p.ID, objc.Sel("cancel"))
}

// Pauses progress tracking.
//
// # Discussion
// 
// This method invokes the block for [PausingHandler], if there is one, and
// ensures that any subsequent reads of the [Paused] property return [true].
// 
// If the receiver has suboperations, the system pauses their progress as
// well.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Foundation/Progress/pause()
func (p Progress) Pause() {
	objc.Send[objc.ID](p.ID, objc.Sel("pause"))
}

// Resumes progress tracking.
//
// # Discussion
// 
// This method invokes the block for [ResumingHandler], if there is one, and
// ensures that any subsequent reads of the [Paused] property return [false].
// 
// If the receiver has suboperations, the system resumes their progress as
// well.
//
// [false]: https://developer.apple.com/documentation/Swift/false
//
// See: https://developer.apple.com/documentation/Foundation/Progress/resume()
func (p Progress) Resume() {
	objc.Send[objc.ID](p.ID, objc.Sel("resume"))
}

// Sets a value in the user info dictionary.
//
// objectOrNil: The object to set for the specified key, or `nil` to remove an existing
// entry in the dictionary.
//
// key: The key for storing the specified object.
//
// # Discussion
// 
// Use this method to set a value in the [UserInfo] dictionary, with
// appropriate KVO notification for properties with values that can depend on
// values in the user info dictionary, like [LocalizedDescription].
// 
// Supply a value of `nil` to remove an existing dictionary entry for a
// specified key.
//
// See: https://developer.apple.com/documentation/Foundation/Progress/setUserInfoObject(_:forKey:)
func (p Progress) SetUserInfoObjectForKey(objectOrNil objectivec.IObject, key NSProgressUserInfoKey) {
	objc.Send[objc.ID](p.ID, objc.Sel("setUserInfoObject:forKey:"), objectOrNil, objc.String(string(key)))
}

// Publishes the progress object for other processes to observe it.
//
// # Discussion
// 
// Entries in the user info dictionary determine whether another process can
// discover the progress object to observe it, and how it does that. For
// example, a [fileURLKey] entry makes a progress object discoverable by
// corresponding invokers of [AddSubscriberForFileURLWithPublishingHandler].
// The system constrains access to the published progress URL with your app
// sandbox. If you can’t see the file due to the app’s sandbox
// restrictions, you can’t observe the progress on it.
// 
// When you make a progress object observable by other processes, you must
// ensure that at least [LocalizedDescription], [Indeterminate], and
// [FractionCompleted] always work when you send proxies of your progress
// object in other processes. You make [Indeterminate] and [FractionCompleted]
// work by accurately setting the total and completed unit counts of the
// progress. You make [LocalizedDescription] work by setting the value of the
// kind property to something valid, like [file], and then fulfilling the
// requirements for that kind of progress.
// 
// You can instead set the value of [LocalizedDescription] directly, but
// that’s not perfectly reliable because other processes might be using a
// different localization than yours.
// 
// You can publish an instance of [NSProgress] one time only.
//
// [fileURLKey]: https://developer.apple.com/documentation/Foundation/ProgressUserInfoKey/fileURLKey
// [file]: https://developer.apple.com/documentation/Foundation/ProgressKind/file
//
// See: https://developer.apple.com/documentation/Foundation/Progress/publish()
func (p Progress) Publish() {
	objc.Send[objc.ID](p.ID, objc.Sel("publish"))
}

// Removes a progress object from publication, making it unobservable by other
// processes.
//
// See: https://developer.apple.com/documentation/Foundation/Progress/unpublish()
func (p Progress) Unpublish() {
	objc.Send[objc.ID](p.ID, objc.Sel("unpublish"))
}

// Retrieves the current thread’s progress object, executes the specified
// block, and increments the progress object by the specified units of work.
//
// # Discussion
// 
// Use this function as a convenience method to wrap an existing method or
// block to increment the current progress object. This function retrieves the
// current progress object, does the work you specify in the block. When the
// block is complete, this function increments the current progress object.
// This function is the same as calling [BecomeCurrentWithPendingUnitCount],
// doing the work you specify in the block, and calling [ResignCurrent].
//
// See: https://developer.apple.com/documentation/Foundation/NSProgress/performAsCurrentWithPendingUnitCount:usingBlock:
func (p Progress) PerformAsCurrentWithPendingUnitCountUsingBlock(unitCount int64, work VoidHandler) {
		_block1, _cleanup1 := NewVoidBlock(work)
	defer _cleanup1()
		objc.Send[objc.ID](p.ID, objc.Sel("performAsCurrentWithPendingUnitCount:usingBlock:"), unitCount, _block1)
}





// Creates and returns a progress instance with the specified unit count that
// isn’t part of any existing progress tree.
//
// unitCount: The total number of units of work to assign to the progress instance.
//
// # Return Value
// 
// A new progress instance with its containing progress object set to `nil.`
//
// # Discussion
// 
// Use this method to create the top-level progress object that your custom
// classes return. The receiver of the returned progress object can add it to
// a progress tree using [AddChildWithPendingUnitCount].
// 
// You’re responsible for updating the progress count of the created
// progress object. You can invoke this method on one thread and then message
// the returned [NSProgress] on another thread. For example, you can capture
// the created progress instance in a block that you pass to [dispatch_async].
// In that block, you can invoke methods like
// [BecomeCurrentWithPendingUnitCount] or [ResignCurrent], and set the
// [CompletedUnitCount] or [Cancelled] properties as your app finishes its
// work.
//
// [dispatch_async]: https://developer.apple.com/documentation/Dispatch/dispatch_async
//
// See: https://developer.apple.com/documentation/Foundation/Progress/discreteProgress(totalUnitCount:)
func (_ProgressClass ProgressClass) DiscreteProgressWithTotalUnitCount(unitCount int64) Progress {
	rv := objc.Send[objc.ID](objc.ID(_ProgressClass.class), objc.Sel("discreteProgressWithTotalUnitCount:"), unitCount)
	return NSProgressFromID(rv)
}

// Returns the progress instance, if any.
//
// # Return Value
// 
// The progress instance for the current thread, if any.
//
// # Discussion
// 
// If you invoke [BecomeCurrentWithPendingUnitCount] on the current thread,
// this method returns the progress instance.
// 
// Use this per-thread [CurrentProgress] value to allow code that performs
// work to report useful progress even when it’s widely separated from the
// code that actually presents progress information to the user, and without
// requiring layers of intervening code to pass around an [NSProgress]
// instance.
// 
// To ensure that you report progress in known units of work, you typically
// work with a suboperation progress object that you create by calling
// [DiscreteProgressWithTotalUnitCount].
//
// See: https://developer.apple.com/documentation/Foundation/Progress/current()
func (_ProgressClass ProgressClass) CurrentProgress() Progress {
	rv := objc.Send[objc.ID](objc.ID(_ProgressClass.class), objc.Sel("currentProgress"))
	return NSProgressFromID(rv)
}

// Registers a file URL to hear about the progress of a file operation.
//
// url: The URL of the file to observe.
//
// publishingHandler: A closure that the system invokes when a progress object that represents a
// file operation matching the specified URL calls [Publish].
//
// publishingHandler is a [foundation.NSProgressPublishingHandler].
//
// # Return Value
// 
// A proxy of the progress object to observe.
//
// # Discussion
// 
// The system invokes the passed-in block when a progress object calls
// [Publish] with a [fileURLKey] user info dictionary entry that’s a URL
// that is the same as this method’s URL, or that is an item that the URL
// directly contains. The progress object that passes to your block is a proxy
// of the published progress object. The passed-in block may return another
// block. If it does, the system invokes the returned block when the observed
// progress object invokes [Unpublish], the publishing process terminates, or
// you invoke [RemoveSubscriber]. The system invokes the blocks you provide on
// the main thread.
//
// [fileURLKey]: https://developer.apple.com/documentation/Foundation/ProgressUserInfoKey/fileURLKey
//
// See: https://developer.apple.com/documentation/Foundation/Progress/addSubscriber(forFileURL:withPublishingHandler:)
func (_ProgressClass ProgressClass) AddSubscriberForFileURLWithPublishingHandler(url INSURL, publishingHandler ErrorHandler) objectivec.IObject {
		_block1, _cleanup1 := NewErrorBlock(publishingHandler)
	defer _cleanup1()
		rv := objc.Send[objc.ID](objc.ID(_ProgressClass.class), objc.Sel("addSubscriberForFileURL:withPublishingHandler:"), url, _block1)
	return objectivec.Object{ID: rv}
}

// Removes a proxy progress object that the add subscriber method returns.
//
// subscriber: The proxy of the progress object to observe.
//
// # Discussion
// 
// If the block for [AddSubscriberForFileURLWithPublishingHandler] returns a
// closure, the system invokes that closure on the main thread when you invoke
// [RemoveSubscriber].
//
// See: https://developer.apple.com/documentation/Foundation/Progress/removeSubscriber(_:)
func (_ProgressClass ProgressClass) RemoveSubscriber(subscriber objectivec.IObject) {
	objc.Send[objc.ID](objc.ID(_ProgressClass.class), objc.Sel("removeSubscriber:"), subscriber)
}








// The total number of tracked units of work for the current progress.
//
// # Discussion
// 
// For an [NSProgress] with a kind of [file], the unit of this property is
// bytes, and the [fileTotalCountKey] and [fileCompletedCountKey] keys in the
// `userInfo` dictionary report the overall count of files.
// 
// For any other kind of [NSProgress], the unit of measurement doesn’t
// matter as long as it’s consistent. You can report the values to the user
// in the [LocalizedDescription] and [LocalizedAdditionalDescription].
//
// [fileCompletedCountKey]: https://developer.apple.com/documentation/Foundation/ProgressUserInfoKey/fileCompletedCountKey
// [fileTotalCountKey]: https://developer.apple.com/documentation/Foundation/ProgressUserInfoKey/fileTotalCountKey
// [file]: https://developer.apple.com/documentation/Foundation/ProgressKind/file
//
// See: https://developer.apple.com/documentation/Foundation/Progress/totalUnitCount
func (p Progress) TotalUnitCount() int64 {
	rv := objc.Send[int64](p.ID, objc.Sel("totalUnitCount"))
	return rv
}
func (p Progress) SetTotalUnitCount(value int64) {
	objc.Send[struct{}](p.ID, objc.Sel("setTotalUnitCount:"), value)
}



// The number of completed units of work for the current job.
//
// # Discussion
// 
// For an [NSProgress] with a kind of [file], the unit of this property is
// bytes, and the [fileTotalCountKey] and [fileCompletedCountKey] keys in the
// `userInfo` dictionary report the overall count of files.
// 
// For any other kind of [NSProgress], the unit of measurement doesn’t
// matter as long as it’s consistent. You can report the values to the user
// in the [LocalizedDescription] and [LocalizedAdditionalDescription].
//
// [fileCompletedCountKey]: https://developer.apple.com/documentation/Foundation/ProgressUserInfoKey/fileCompletedCountKey
// [fileTotalCountKey]: https://developer.apple.com/documentation/Foundation/ProgressUserInfoKey/fileTotalCountKey
// [file]: https://developer.apple.com/documentation/Foundation/ProgressKind/file
//
// See: https://developer.apple.com/documentation/Foundation/Progress/completedUnitCount
func (p Progress) CompletedUnitCount() int64 {
	rv := objc.Send[int64](p.ID, objc.Sel("completedUnitCount"))
	return rv
}
func (p Progress) SetCompletedUnitCount(value int64) {
	objc.Send[struct{}](p.ID, objc.Sel("setCompletedUnitCount:"), value)
}



// A localized description of tracked progress for the receiver.
//
// # Discussion
// 
// If you don’t specify your own custom value for this property,
// [NSProgress] uses the value of the [Kind] property to determine how to use
// the values of other properties, as well as values in the user info
// dictionary, to return an automatically computed string. If it fails to do
// that, it returns an empty string.
// 
// The `localizedDescription` represents a general description of the work the
// receiver tracks. Depending on the kind of progress, the completed and total
// unit counts, and other parameters, localized descriptions resemble the
// following:
// 
// - Copying 10 files…
// - 30% completed
// - Copying “TextEdit”…
// 
// By default, [NSProgress] is KVO-compliant for this property. It sends
// notifications on the same thread that updates the property.
//
// See: https://developer.apple.com/documentation/Foundation/Progress/localizedDescription
func (p Progress) LocalizedDescription() string {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("localizedDescription"))
	return NSStringFromID(rv).String()
}
func (p Progress) SetLocalizedDescription(value string) {
	objc.Send[struct{}](p.ID, objc.Sel("setLocalizedDescription:"), objc.String(value))
}



// A more specific localized description of tracked progress for the receiver.
//
// # Discussion
// 
// If you don’t specify your own custom value for this property,
// [NSProgress] uses the value of the [Kind] property to determine how to use
// the values of other properties, as well as values in the user info
// dictionary, to return an automatically computed string. If it fails to do
// that, it returns an empty string.
// 
// The [LocalizedAdditionalDescription] is more specific than
// [LocalizedDescription] about the work the receiver is tracking at any
// particular moment. Depending on the kind of progress, the completed and
// total unit counts, and other parameters, localized additional descriptions
// resemble the following:
// 
// - 3 of 10 files - 123 KB of 789.1 MB - 3.3 MB of 103.92 GB – 2 hours
// remaining - 1.61 GB of 3.22 GB (2 KB/sec) – 2 minutes remaining - 1
// minute remaining (1 KB/sec)
// 
// By default, [NSProgress] is KVO-compliant for this property. It sends
// notifications on the same thread that updates the property.
//
// See: https://developer.apple.com/documentation/Foundation/Progress/localizedAdditionalDescription
func (p Progress) LocalizedAdditionalDescription() string {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("localizedAdditionalDescription"))
	return NSStringFromID(rv).String()
}
func (p Progress) SetLocalizedAdditionalDescription(value string) {
	objc.Send[struct{}](p.ID, objc.Sel("setLocalizedAdditionalDescription:"), objc.String(value))
}



// A Boolean value that indicates whether the receiver is tracking work that
// you can cancel.
//
// # Discussion
// 
// By default, [NSProgress] objects are cancelable.
// 
// You typically use this property to communicate whether controls for
// canceling appear in a progress-reporting user interface. [NSProgress]
// itself doesn’t do anything with this property other than help pass the
// value from progress reporters to progress observers.
// 
// If an [NSProgress] is cancelable, implement the ability to cancel progress
// either by setting a block for the [CancellationHandler] property, or by
// polling the [Cancelled] property periodically while performing the relevant
// work.
// 
// It’s valid for the value of this property to change during the lifetime
// of an [NSProgress] object. By default, [NSProgress] is KVO-compliant for
// this property. It sends notifications on the same thread that updates the
// property.
//
// See: https://developer.apple.com/documentation/Foundation/Progress/isCancellable
func (p Progress) Cancellable() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("isCancellable"))
	return rv
}
func (p Progress) SetCancellable(value bool) {
	objc.Send[struct{}](p.ID, objc.Sel("setCancellable:"), value)
}



// A Boolean value that Indicates whether the receiver is tracking canceled
// work.
//
// # Discussion
// 
// By default, [NSProgress] is KVO-compliant for this property. It sends
// notifications on the same thread that updates the property.
// 
// If the receiver has a canceled containing progress object, the receiver
// reports a canceled status.
//
// See: https://developer.apple.com/documentation/Foundation/Progress/isCancelled
func (p Progress) Cancelled() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("isCancelled"))
	return rv
}



// The block to invoke when canceling progress.
//
// # Discussion
// 
// If the receiver is a suboperation of another progress object, the system
// invokes the [CancellationHandler] block when canceling the containing
// progress object.
// 
// # Special Considerations
// 
// You’re responsible for canceling any work for the progress object.
// 
// You can invoke the cancellation handler on any queue. If you must do work
// on a specific queue, dispatch to that queue from within the cancellation
// handler block.
//
// See: https://developer.apple.com/documentation/Foundation/Progress/cancellationHandler
func (p Progress) CancellationHandler() VoidHandler {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("cancellationHandler"))
	_ = rv
	return nil
}
func (p Progress) SetCancellationHandler(value VoidHandler) {
	block, cleanup := NewVoidBlock(value)
	defer cleanup()
	objc.Send[struct{}](p.ID, objc.Sel("setCancellationHandler:"), block)
}



// A Boolean value that indicates whether the receiver is tracking work that
// you can pause.
//
// # Discussion
// 
// By default, [NSProgress] objects aren’t pausable.
// 
// You typically use this property to communicate whether controls for pausing
// appear in a progress-reporting user interface. [NSProgress] itself
// doesn’t do anything with this property other than help pass the value
// from progress reporters to progress observers.
// 
// If an [NSProgress] is pausable, implement the ability to pause either by
// setting a block for the [PausingHandler] property, or by polling the
// [Paused] property periodically while performing the relevant work.
// 
// It’s valid for the value of this property to change during the lifetime
// of an [NSProgress] object. By default, [NSProgress] is KVO-compliant for
// this property. It sends notifications on the same thread that updates the
// property.
//
// See: https://developer.apple.com/documentation/Foundation/Progress/isPausable
func (p Progress) Pausable() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("isPausable"))
	return rv
}
func (p Progress) SetPausable(value bool) {
	objc.Send[struct{}](p.ID, objc.Sel("setPausable:"), value)
}



// A Boolean value that indicates whether the receiver is tracking paused
// work.
//
// # Discussion
// 
// By default, [NSProgress] is KVO-compliant for this property. It sends
// notifications on the same thread that updates the property.
// 
// If the receiver has a paused containing progress object, the receiver
// reports a paused status.
//
// See: https://developer.apple.com/documentation/Foundation/Progress/isPaused
func (p Progress) Paused() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("isPaused"))
	return rv
}



// The block to invoke when pausing progress.
//
// # Discussion
// 
// If the receiver is a suboperation of another progress object, the system
// invokes the [PausingHandler] block when pausing the containing progress
// object.
// 
// # Special Considerations
// 
// You’re responsible for pausing any work for the progress object.
// 
// You can invoke the pausing handler on any queue. If you must do work on a
// specific queue, dispatch to that queue from within the pausing handler
// block.
//
// See: https://developer.apple.com/documentation/Foundation/Progress/pausingHandler
func (p Progress) PausingHandler() VoidHandler {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("pausingHandler"))
	_ = rv
	return nil
}
func (p Progress) SetPausingHandler(value VoidHandler) {
	block, cleanup := NewVoidBlock(value)
	defer cleanup()
	objc.Send[struct{}](p.ID, objc.Sel("setPausingHandler:"), block)
}



// A Boolean value that indicates whether the tracked progress is
// indeterminate.
//
// # Discussion
// 
// Use [Indeterminate] progress only when you’re unable to determine a
// reasonable value for either [CompletedUnitCount] or [TotalUnitCount].
// Progress is indeterminate when the value of the [TotalUnitCount] or
// [CompletedUnitCount] is less than zero or if both values are zero. When
// progress is indeterminate, [FractionCompleted] returns `0.0` and [Finished]
// returns `false`.
// 
// By default, [NSProgress] is KVO-compliant for this property. It sends
// notifications on the same thread that updates the property.
// 
// The following code snippet clarifies the behavior for setting both
// [TotalUnitCount] and [CompletedUnitCount] to `0`.
//
// See: https://developer.apple.com/documentation/Foundation/Progress/isIndeterminate
func (p Progress) Indeterminate() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("isIndeterminate"))
	return rv
}



// The fraction of the overall work that the progress object completes,
// including work from its suboperations.
//
// # Discussion
// 
// If the receiver object doesn’t have any suboperations,
// [FractionCompleted] is generally the result of dividing
// [CompletedUnitCount] by [TotalUnitCount]. Setting both [TotalUnitCount] and
// [CompletedUnitCount] properties to zero indicates that there is no progress
// to track. In this case, the [Indeterminate] property returns [false] and
// the [FractionCompleted] property returns `0.0`.
// 
// If the receiver does have suboperations, [FractionCompleted] reflects
// progress from those progress objects in addition to its own
// [CompletedUnitCount]. When the suboperations finish, the
// [CompletedUnitCount] of the containing progress object updates.
//
// [false]: https://developer.apple.com/documentation/Swift/false
//
// See: https://developer.apple.com/documentation/Foundation/Progress/fractionCompleted
func (p Progress) FractionCompleted() float64 {
	rv := objc.Send[float64](p.ID, objc.Sel("fractionCompleted"))
	return rv
}



// A Boolean value that indicates the progress object is complete.
//
// # Discussion
// 
// A progress object finishes when the [CompletedUnitCount] equals or exceeds
// the [TotalUnitCount].
// 
// By default, [NSProgress] is KVO-compliant for this property. It sends
// notifications on the same thread that updates the property.
//
// See: https://developer.apple.com/documentation/Foundation/Progress/isFinished
func (p Progress) Finished() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("isFinished"))
	return rv
}



// The block to invoke when progress resumes.
//
// # Discussion
// 
// If the receiver is a suboperation of another progress object, the system
// invokes the [ResumingHandler] block when pausing the containing progress
// object.
// 
// # Special Considerations
// 
// You’re responsible for resuming any work for the progress object.
// 
// You can invoke the resuming handler on any queue. If you must do work on a
// specific queue, dispatch to that queue from within the resuming handler
// block.
//
// See: https://developer.apple.com/documentation/Foundation/Progress/resumingHandler
func (p Progress) ResumingHandler() VoidHandler {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("resumingHandler"))
	_ = rv
	return nil
}
func (p Progress) SetResumingHandler(value VoidHandler) {
	block, cleanup := NewVoidBlock(value)
	defer cleanup()
	objc.Send[struct{}](p.ID, objc.Sel("setResumingHandler:"), block)
}



// An object that represents the kind of progress for the progress object.
//
// # Discussion
// 
// This property identifies the kind of progress for the progress object, such
// as [file]. It can be `nil`.
// 
// If you set a non-`nil` value to [Kind], the default [LocalizedDescription]
// getter uses the kind of progress. The [LocalizedDescription] getter
// determines how to use the values of other properties, along with values in
// the user info dictionary, to create a string representation.
//
// [file]: https://developer.apple.com/documentation/Foundation/ProgressKind/file
//
// See: https://developer.apple.com/documentation/Foundation/Progress/kind
func (p Progress) Kind() NSProgressKind {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("kind"))
	return ProgressKind(NSStringFromID(rv).String())
}
func (p Progress) SetKind(value NSProgressKind) {
	objc.Send[struct{}](p.ID, objc.Sel("setKind:"), objc.String(string(value)))
}



// A dictionary of arbitrary values for the receiver.
//
// # Discussion
// 
// A KVO-compliant dictionary that changes in response to
// [SetUserInfoObjectForKey]. The dictionary sends all of its KVO
// notifications on the thread that updates the property.
// 
// Some entries have meanings that the [NSProgress] class recognizes. For more
// information, see Recognizing Kinds of Progress, Using General Keys, Using
// File Operation Keys, and Recognizing Kinds of File Operations.
//
// See: https://developer.apple.com/documentation/Foundation/Progress/userInfo
func (p Progress) UserInfo() INSDictionary {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("userInfo"))
	return NSDictionaryFromID(objc.ID(rv))
}



// The kind of file operation for the progress object.
//
// # Discussion
// 
// Set this value when the [Kind] property is [file] to describe the kind of
// file operation.
// 
// If present, [NSProgress] presents additional information in its localized
// description by setting a value in the `userInfo` dictionary.
//
// [file]: https://developer.apple.com/documentation/Foundation/ProgressKind/file
//
// See: https://developer.apple.com/documentation/Foundation/Progress/fileOperationKind-swift.property
func (p Progress) FileOperationKind() NSProgressFileOperationKind {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("fileOperationKind"))
	return NSProgressFileOperationKind(NSStringFromID(rv).String())
}
func (p Progress) SetFileOperationKind(value NSProgressFileOperationKind) {
	objc.Send[struct{}](p.ID, objc.Sel("setFileOperationKind:"), objc.String(string(value)))
}



// A URL that represents the file for the current progress object.
//
// # Discussion
// 
// Set this value for a progress that you [Publish] to subscribers that
// register for updates using [AddSubscriberForFileURLWithPublishingHandler].
// 
// If present, [NSProgress] presents additional information in its localized
// description by setting a value in the `userInfo` dictionary.
//
// See: https://developer.apple.com/documentation/Foundation/Progress/fileURL
func (p Progress) FileURL() INSURL {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("fileURL"))
	return NSURLFromID(objc.ID(rv))
}
func (p Progress) SetFileURL(value INSURL) {
	objc.Send[struct{}](p.ID, objc.Sel("setFileURL:"), value)
}



// A Boolean value that indicates when the observed progress object invokes
// the publish method before you subscribe to it.
//
// # Discussion
// 
// The publish and subscribe mechanism is generally , in that when you invoke
// [AddSubscriberForFileURLWithPublishingHandler], the system invokes your
// block for every relevant published and unpublished progress object.
// Sometimes you need to implement behavior, in which you do something either
// exactly when new progress begins or not at all.
// 
// In the example above, the Dock doesn’t animate file icons when this
// method returns [true].
// 
// There’s no reliable definition of in this case, which involves multiple
// processes in a preemptively scheduled system. Don’t use this method for
// anything more important than best efforts at animating. It can be
// inaccurate due to processes coming and going from unpredictable user
// actions.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Foundation/Progress/isOld
func (p Progress) Old() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("isOld"))
	return rv
}



// A value that indicates the estimated amount of time remaining to complete
// the progress.
//
// # Discussion
// 
// If present, [NSProgress] presents additional information in its localized
// description by setting a value in the `userInfo` dictionary.
//
// See: https://developer.apple.com/documentation/Foundation/NSProgress/estimatedTimeRemaining
func (p Progress) EstimatedTimeRemaining() INSNumber {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("estimatedTimeRemaining"))
	return NSNumberFromID(objc.ID(rv))
}
func (p Progress) SetEstimatedTimeRemaining(value INSNumber) {
	objc.Send[struct{}](p.ID, objc.Sel("setEstimatedTimeRemaining:"), value)
}



// The number of completed files for a file progress object.
//
// # Discussion
// 
// If the current progress is operating on a set of files, set this property
// to the number of completed files in the operation.
// 
// If present, [NSProgress] presents additional information in its localized
// description by setting a value in the `userInfo` dictionary.
//
// See: https://developer.apple.com/documentation/Foundation/NSProgress/fileCompletedCount
func (p Progress) FileCompletedCount() INSNumber {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("fileCompletedCount"))
	return NSNumberFromID(objc.ID(rv))
}
func (p Progress) SetFileCompletedCount(value INSNumber) {
	objc.Send[struct{}](p.ID, objc.Sel("setFileCompletedCount:"), value)
}



// The total number of files for a file progress object.
//
// # Discussion
// 
// If the current progress is operating on a set of files, set this property
// to the total number of files in the operation.
// 
// If present, [NSProgress] presents additional information in its localized
// description by setting a value in the `userInfo` dictionary.
//
// See: https://developer.apple.com/documentation/Foundation/NSProgress/fileTotalCount
func (p Progress) FileTotalCount() INSNumber {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("fileTotalCount"))
	return NSNumberFromID(objc.ID(rv))
}
func (p Progress) SetFileTotalCount(value INSNumber) {
	objc.Send[struct{}](p.ID, objc.Sel("setFileTotalCount:"), value)
}



// A value that represents the speed of data processing, in bytes per second.
//
// # Discussion
// 
// If present, [NSProgress] presents additional information in its localized
// description by setting a value in the `userInfo` dictionary.
//
// See: https://developer.apple.com/documentation/Foundation/NSProgress/throughput
func (p Progress) Throughput() INSNumber {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("throughput"))
	return NSNumberFromID(objc.ID(rv))
}
func (p Progress) SetThroughput(value INSNumber) {
	objc.Send[struct{}](p.ID, objc.Sel("setThroughput:"), value)
}




















// PerformAsCurrentWithPendingUnitCountUsingBlockSync is a synchronous wrapper around [Progress.PerformAsCurrentWithPendingUnitCountUsingBlock].
// It blocks until the completion handler fires or the context is cancelled.
func (p Progress) PerformAsCurrentWithPendingUnitCountUsingBlockSync(ctx context.Context, unitCount int64) error {
	done := make(chan struct{}, 1)
	p.PerformAsCurrentWithPendingUnitCountUsingBlock(unitCount, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}






