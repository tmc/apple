// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The interface a file coordinator uses to inform an object presenting a file about changes to that file made elsewhere in the system.
//
// See: https://developer.apple.com/documentation/Foundation/NSFilePresenter
type NSFilePresenter interface {
	objectivec.IObject

	// The URL of the presented file or directory.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSFilePresenter/presentedItemURL
	PresentedItemURL() INSURL

	// The operation queue in which to execute presenter-related messages.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSFilePresenter/presentedItemOperationQueue
	PresentedItemOperationQueue() INSOperationQueue

	// The URL of a secondary item’s primary presented file or directory.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSFilePresenter/primaryPresentedItemURL
	PrimaryPresentedItemURL() INSURL

	// A list of ubiquity attributes used to generate and send notifications whenever an attribute in the list changes.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSFilePresenter/observedPresentedItemUbiquityAttributes
	ObservedPresentedItemUbiquityAttributes() INSSet
}



// NSFilePresenterObject wraps an existing Objective-C object that conforms to the NSFilePresenter protocol.
type NSFilePresenterObject struct {
	objectivec.Object
}
func (o NSFilePresenterObject) BaseObject() objectivec.Object {
	return o.Object
}



// NSFilePresenterObjectFromID constructs a [NSFilePresenterObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSFilePresenterObjectFromID(id objc.ID) NSFilePresenterObject {
	return NSFilePresenterObject{
		Object: objectivec.ObjectFromID(id),
	}
}




// The URL of the presented file or directory.
//
// See: https://developer.apple.com/documentation/Foundation/NSFilePresenter/presentedItemURL

func (o NSFilePresenterObject) PresentedItemURL() INSURL {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("presentedItemURL"))
	return NSURLFromID(rv)
	}

// The operation queue in which to execute presenter-related messages.
//
// See: https://developer.apple.com/documentation/Foundation/NSFilePresenter/presentedItemOperationQueue

func (o NSFilePresenterObject) PresentedItemOperationQueue() INSOperationQueue {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("presentedItemOperationQueue"))
	return NSOperationQueueFromID(rv)
	}

// The URL of a secondary item’s primary presented file or directory.
//
// See: https://developer.apple.com/documentation/Foundation/NSFilePresenter/primaryPresentedItemURL

func (o NSFilePresenterObject) PrimaryPresentedItemURL() INSURL {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("primaryPresentedItemURL"))
	return NSURLFromID(rv)
	}

// A list of ubiquity attributes used to generate and send notifications
// whenever an attribute in the list changes.
//
// See: https://developer.apple.com/documentation/Foundation/NSFilePresenter/observedPresentedItemUbiquityAttributes

func (o NSFilePresenterObject) ObservedPresentedItemUbiquityAttributes() INSSet {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("observedPresentedItemUbiquityAttributes"))
	return NSSetFromID(rv)
	}

// Notifies your object that another object or process wants to read the
// presented file or directory.
//
// reader: A [Block object] that takes another block as a parameter and returns no
// value. The `reacquirer` block is one you pass to the `reader` block so that
// your object can be notified when the `reader` is done. If your object does
// not need to be notified, it can pass `nil` for the `reacquirer` block.
// //
// [Block object]: https://developer.apple.com/library/archive/documentation/General/Conceptual/DevPedia-CocoaCore/Block.html#//apple_ref/doc/uid/TP40008195-CH3
//
// # Discussion
// 
// You use this method to provide an appropriate response when another object
// wants to read from your presented URL. For example, when this method is
// called, you might temporarily stop making changes to the file or directory.
// After taking any appropriate steps, you must execute the block in the
// `reader` parameter to let the waiting object know that it may now proceed
// with its task. If you want to be notified when the reader has completed its
// task, pass your own block to the reader and use that block to reacquire the
// file or URL for your own uses.
// 
// The following listing shows a simple implementation of this method that
// sets a Boolean flag that the file being monitored is not writable at the
// moment. After setting the flag, it executes the reader block and passes in
// yet another block for the reader to execute when it is done.
// 
// Your implementation of this method is executed using the queue in the
// [PresentedItemOperationQueue] property. Your reacquirer block is executed
// on the queue associated with the reader.
//
// See: https://developer.apple.com/documentation/Foundation/NSFilePresenter/relinquishPresentedItem(toReader:)

func (o NSFilePresenterObject) RelinquishPresentedItemToReader(reader VoidHandler) {
	
	objc.Send[struct{}](o.ID, objc.Sel("relinquishPresentedItemToReader:"), reader)
	}

// Notifies your object that another object or process wants to write to the
// presented file or directory.
//
// writer: A [Block object] that takes another block as a parameter and returns no
// value. The `reacquirer` block is one you pass to the `writer` block so that
// your object can be notified when the `writer` is done. If your object does
// not need to be notified, it can pass `nil` for the `reacquirer` block.
// //
// [Block object]: https://developer.apple.com/library/archive/documentation/General/Conceptual/DevPedia-CocoaCore/Block.html#//apple_ref/doc/uid/TP40008195-CH3
//
// # Discussion
// 
// You use this method to provide an appropriate response when another object
// wants to write to your presented URL. For example, when this method is
// called, you would likely stop making changes to the file or directory.
// After taking any appropriate steps, you must execute the block in the
// `writer` parameter to let the waiting object know that it may now proceed
// with its task. If you want to be notified when the writer has completed its
// task, pass your own block to the writer and use that block to reacquire the
// file or URL for your own uses.
// 
// If the writer changes the file or directory, you do not need to incorporate
// those changes in your reacquirer block. Instead, implement the
// [PresentedItemDidChange] method and use it to detect when a writer actually
// wrote its changes to disk.
// 
// The following listing shows a simple implementation of this method that
// sets a Boolean flag that the file being monitored is not writable at the
// moment. After setting the flag, it executes the writer block and passes in
// yet another block for the writer to execute when it is done.
// 
// Your implementation of this method is executed using the queue in the
// [PresentedItemOperationQueue] property. Your reacquirer block is executed
// on the queue associated with the writer.
//
// See: https://developer.apple.com/documentation/Foundation/NSFilePresenter/relinquishPresentedItem(toWriter:)

func (o NSFilePresenterObject) RelinquishPresentedItemToWriter(writer VoidHandler) {
	
	objc.Send[struct{}](o.ID, objc.Sel("relinquishPresentedItemToWriter:"), writer)
	}

// Tells your object to save any unsaved changes for the presented item.
//
// completionHandler: The [Block object] to call after you save your changes. If you saved your
// changes successfully, pass `nil` for the block’s `errorOrNil` parameter;
// otherwise, pass an error object indicating why the changes could not be
// saved.
// //
// [Block object]: https://developer.apple.com/library/archive/documentation/General/Conceptual/DevPedia-CocoaCore/Block.html#//apple_ref/doc/uid/TP40008195-CH3
//
// # Discussion
// 
// The file coordinator calls this method to ensure that all objects trying to
// access the file or directory see the same contents. Implement this method
// if your object can change the presented item in a way that requires you to
// write those changes back to disk. If your presenter object does not make
// changes that need to be saved, you do not need to implement this method.
//
// See: https://developer.apple.com/documentation/Foundation/NSFilePresenter/savePresentedItemChanges(completionHandler:)

func (o NSFilePresenterObject) SavePresentedItemChangesWithCompletionHandler(completionHandler ErrorHandler) {
	
	objc.Send[struct{}](o.ID, objc.Sel("savePresentedItemChangesWithCompletionHandler:"), completionHandler)
	}

// Tells your object that its presented item is about to be deleted.
//
// completionHandler: The [Block object] to call after updating your data structures. Pass `nil`
// to the block’s `errorOrNil` parameter if you were able to successfully
// prepare for the deletion of the item. Pass an error object if your object
// could not prepare itself properly.
// //
// [Block object]: https://developer.apple.com/library/archive/documentation/General/Conceptual/DevPedia-CocoaCore/Block.html#//apple_ref/doc/uid/TP40008195-CH3
//
// # Discussion
// 
// A file coordinator calls this method when your object’s presented item is
// about to be deleted. You can use this method to perform any actions that
// are needed to prepare for the deletion. For example, document objects
// typically use this method to close the document.
//
// See: https://developer.apple.com/documentation/Foundation/NSFilePresenter/accommodatePresentedItemDeletion(completionHandler:)

func (o NSFilePresenterObject) AccommodatePresentedItemDeletionWithCompletionHandler(completionHandler ErrorHandler) {
	
	objc.Send[struct{}](o.ID, objc.Sel("accommodatePresentedItemDeletionWithCompletionHandler:"), completionHandler)
	}

// Tells your object that the presented item moved or was renamed.
//
// newURL: The URL containing the new path to the presented item.
//
// # Discussion
// 
// Use this method to update the value returned by the [PresentedItemURL]
// property of your object.
//
// See: https://developer.apple.com/documentation/Foundation/NSFilePresenter/presentedItemDidMove(to:)

func (o NSFilePresenterObject) PresentedItemDidMoveToURL(newURL INSURL) {
	
	objc.Send[struct{}](o.ID, objc.Sel("presentedItemDidMoveToURL:"), newURL)
	}

// Tells your object that the presented item’s contents or attributes
// changed.
//
// # Discussion
// 
// You can use this method to update your internal data structures to reflect
// the changes to the presented item. This method reports both changes to the
// file’s contents, such as the data in a file or the files in a directory,
// or the attributes of the item, such as whether the Hide extension checkbox
// of a file was toggled.
// 
// Because this method notifies you of both attribute and content changes, you
// might want to check the modification date before needlessly rereading the
// contents of a file. To do that, you must store the date when your object
// last made changes to the file and compare that date with the item’s
// current modification date. Use the
// [CoordinateReadingItemAtURLOptionsErrorByAccessor] method of a file
// coordinator to ensure exclusive access to the file when reading the current
// modification date.
//
// See: https://developer.apple.com/documentation/Foundation/NSFilePresenter/presentedItemDidChange()

func (o NSFilePresenterObject) PresentedItemDidChange() {
	
	objc.Send[struct{}](o.ID, objc.Sel("presentedItemDidChange"))
	}

// Tells the delegate that a new version of the file or file package was
// added.
//
// version: The file version object containing information about the new file version.
//
// # Discussion
// 
// Your delegate can use this method to determine how to incorporate data from
// the new version of the file or file package. If the file has not been
// modified by your code, you might simply update to the new version quietly.
// However, if your application has its own changes, you might need to ask the
// user how to proceed.
//
// See: https://developer.apple.com/documentation/Foundation/NSFilePresenter/presentedItemDidGain(_:)

func (o NSFilePresenterObject) PresentedItemDidGainVersion(version INSFileVersion) {
	
	objc.Send[struct{}](o.ID, objc.Sel("presentedItemDidGainVersion:"), version)
	}

// Tells the delegate that a version of the file or file package was removed.
//
// version: The file version object containing information about the version that was
// removed.
//
// # Discussion
// 
// Your delegate can use this method to determine how to handle the loss of
// the specified file version. You can try to revert the presented document to
// a previous version or you might want to prompt the user about how to
// proceed.
//
// See: https://developer.apple.com/documentation/Foundation/NSFilePresenter/presentedItemDidLose(_:)

func (o NSFilePresenterObject) PresentedItemDidLoseVersion(version INSFileVersion) {
	
	objc.Send[struct{}](o.ID, objc.Sel("presentedItemDidLoseVersion:"), version)
	}

// Tells the delegate that some other entity resolved a version conflict for
// the presenter’s file or file package.
//
// version: The version object containing the conflicting change.
//
// # Discussion
// 
// Your delegate can use this method to respond to the resolution of a version
// conflict by a different file presenter. This might occur if a version of
// your application running on another device resolves the conflict first. You
// might then use this method to update your user interface to indicate that
// there is no longer a conflict.
//
// See: https://developer.apple.com/documentation/Foundation/NSFilePresenter/presentedItemDidResolveConflict(_:)

func (o NSFilePresenterObject) PresentedItemDidResolveConflictVersion(version INSFileVersion) {
	
	objc.Send[struct{}](o.ID, objc.Sel("presentedItemDidResolveConflictVersion:"), version)
	}

// Tells the delegate that the item inside the presented directory gained a
// new version.
//
// url: The URL of the item inside the presented directory that gained a new
// version. The item need not be at the top level of the presented directory
// but may itself be inside a nested subdirectory.
//
// version: The file version object containing information about the new file version.
//
// # Discussion
// 
// Your delegate can use this method to determine how to incorporate data from
// the new version of the item. This might involve incorporating the version
// silently or asking the user about how to proceed.
//
// See: https://developer.apple.com/documentation/Foundation/NSFilePresenter/presentedSubitem(at:didGain:)

func (o NSFilePresenterObject) PresentedSubitemAtURLDidGainVersion(url INSURL, version INSFileVersion) {
	
	objc.Send[struct{}](o.ID, objc.Sel("presentedSubitemAtURL:didGainVersion:"), url, version)
	}

// Tells the delegate that the item inside the presented directory lost an
// existing version.
//
// url: The URL of the item inside the presented directory that lost a version. The
// item need not be at the top level of the presented directory but may itself
// be inside a nested subdirectory.
//
// version: The file version object containing information about the version that was
// removed.
//
// # Discussion
// 
// Your delegate can use this method to determine how to handle the loss of
// the specified file version. For an old version, you might not have to do
// anything. However, if your application is currently using the lost version,
// you would need to update your application’s user interface or prompt the
// user about how to proceed.
//
// See: https://developer.apple.com/documentation/Foundation/NSFilePresenter/presentedSubitem(at:didLose:)

func (o NSFilePresenterObject) PresentedSubitemAtURLDidLoseVersion(url INSURL, version INSFileVersion) {
	
	objc.Send[struct{}](o.ID, objc.Sel("presentedSubitemAtURL:didLoseVersion:"), url, version)
	}

// Tells the delegate that the item inside the presented directory had a
// version conflict resolved by an outside entity.
//
// url: The URL of the item inside the presented directory that was in conflict.
// The item need not be at the top level of the presented directory but may
// itself be inside a nested subdirectory.
//
// version: The version object containing the conflicting change.
//
// # Discussion
// 
// Your delegate can use this method to respond to the resolution of a version
// conflict by a different file presenter. This might occur if a version of
// your application running on another device resolves the conflict first. You
// might then use this method to update your user interface to indicate that
// there is no longer a conflict.
//
// See: https://developer.apple.com/documentation/Foundation/NSFilePresenter/presentedSubitem(at:didResolve:)

func (o NSFilePresenterObject) PresentedSubitemAtURLDidResolveConflictVersion(url INSURL, version INSFileVersion) {
	
	objc.Send[struct{}](o.ID, objc.Sel("presentedSubitemAtURL:didResolveConflictVersion:"), url, version)
	}

// Tells the delegate that some entity wants to delete an item that is inside
// of a presented directory.
//
// url: The URL of the item being deleted from the presented directory. The item
// need not be at the top level of the presented directory but may itself be
// inside a nested subdirectory.
//
// completionHandler: The [Block object] to call after updating your data structures. Pass `nil`
// to the block’s `errorOrNil` parameter if you were able to successfully
// prepare for the deletion of the item. Pass an error object if your object
// could not prepare itself properly.
// //
// [Block object]: https://developer.apple.com/library/archive/documentation/General/Conceptual/DevPedia-CocoaCore/Block.html#//apple_ref/doc/uid/TP40008195-CH3
//
// # Discussion
// 
// This method is relevant for applications that present directories. This
// might occur if the delegate manages the contents of a directory or manages
// a file that is implemented as a file package. When called, your
// implementation of this method should take whatever actions needed to update
// your application to handle the deletion of the specified file.
//
// See: https://developer.apple.com/documentation/Foundation/NSFilePresenter/accommodatePresentedSubitemDeletion(at:completionHandler:)

func (o NSFilePresenterObject) AccommodatePresentedSubitemDeletionAtURLCompletionHandler(url INSURL, completionHandler ErrorHandler) {
	
	objc.Send[struct{}](o.ID, objc.Sel("accommodatePresentedSubitemDeletionAtURL:completionHandler:"), url, completionHandler)
	}

// Tells the delegate that an item was added to the presented directory.
//
// url: The URL of the item being added to the presented directory. The item need
// not be at the top level of the presented directory but may itself be inside
// a nested subdirectory.
//
// # Discussion
// 
// This method is relevant for applications that present directories. This
// might occur if the delegate manages the contents of a directory or manages
// a file that is implemented as a file package. Your implementation of this
// method should take whatever actions necessary to incorporate the new file
// or directory into the presented content. For example, you might add the new
// item to your application’s data structures and refresh your user
// interface.
// 
// If the presented directory is a file package, the system calls the
// [PresentedItemDidChange] method if your delegate does not implement this
// method.
//
// See: https://developer.apple.com/documentation/Foundation/NSFilePresenter/presentedSubitemDidAppear(at:)

func (o NSFilePresenterObject) PresentedSubitemDidAppearAtURL(url INSURL) {
	
	objc.Send[struct{}](o.ID, objc.Sel("presentedSubitemDidAppearAtURL:"), url)
	}

// Tells the delegate that an item in the presented directory moved to a new
// location.
//
// oldURL: The original URL of the item inside the presented directory. The item need
// not be at the top level of the presented directory but may itself be inside
// a nested subdirectory.
//
// newURL: The new URL for the item. This URL may or may not be located inside the
// presented directory.
//
// # Discussion
// 
// This method is relevant for applications that present directories. This
// might occur if the delegate manages the contents of a directory or manages
// a file that is implemented as a file package. Your implementation of this
// method should take whatever actions necessary to handle the change in
// location of the specified item. For example, you might update references to
// the item in your application’s data structures and refresh your user
// interface.
// 
// If the presented directory is a file package, the system calls the
// [PresentedItemDidChange] method if your delegate does not implement this
// method.
//
// See: https://developer.apple.com/documentation/Foundation/NSFilePresenter/presentedSubitem(at:didMoveTo:)

func (o NSFilePresenterObject) PresentedSubitemAtURLDidMoveToURL(oldURL INSURL, newURL INSURL) {
	
	objc.Send[struct{}](o.ID, objc.Sel("presentedSubitemAtURL:didMoveToURL:"), oldURL, newURL)
	}

// Tells the delegate that the contents or attributes of the specified item
// changed.
//
// url: The URL of the item in the presented directory that changed. The item need
// not be at the top level of the presented directory but may itself be inside
// a nested subdirectory.
//
// # Discussion
// 
// This method is relevant for applications that present directories. This
// might occur if the delegate manages the contents of a directory or manages
// a file that is implemented as a file package. Your implementation of this
// method should take whatever actions necessary to handle the change in
// content or attributes of the specified item.
// 
// If the presented directory is a file package, the system calls the
// [PresentedItemDidChange] method if your delegate does not implement this
// method.
//
// See: https://developer.apple.com/documentation/Foundation/NSFilePresenter/presentedSubitemDidChange(at:)

func (o NSFilePresenterObject) PresentedSubitemDidChangeAtURL(url INSURL) {
	
	objc.Send[struct{}](o.ID, objc.Sel("presentedSubitemDidChangeAtURL:"), url)
	}

// Tells your object that the file or file package’s ubiquity attributes
// have changed.
//
// attributes: The set of ubiquity attributes that have changed. For information about
// valid ubiquity attributes, see the
// [ObservedPresentedItemUbiquityAttributes] property.
//
// # Discussion
// 
// To specify the ubiquity attributes that trigger notifications, implement
// your file provider’s [ObservedPresentedItemUbiquityAttributes] property.
// If you do not implement this property, the system sends notifications when
// any ubiquity attribute changes.
//
// See: https://developer.apple.com/documentation/Foundation/NSFilePresenter/presentedItemDidChangeUbiquityAttributes(_:)

func (o NSFilePresenterObject) PresentedItemDidChangeUbiquityAttributes(attributes INSSet) {
	
	objc.Send[struct{}](o.ID, objc.Sel("presentedItemDidChangeUbiquityAttributes:"), attributes)
	}

//
// See: https://developer.apple.com/documentation/Foundation/NSFilePresenter/accommodatePresentedItemEviction(completionHandler:)

func (o NSFilePresenterObject) AccommodatePresentedItemEvictionWithCompletionHandler(completionHandler ErrorHandler) {
	
	objc.Send[struct{}](o.ID, objc.Sel("accommodatePresentedItemEvictionWithCompletionHandler:"), completionHandler)
	}















