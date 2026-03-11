// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"context"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSFilePromiseReceiver] class.
var (
	_NSFilePromiseReceiverClass     NSFilePromiseReceiverClass
	_NSFilePromiseReceiverClassOnce sync.Once
)

func getNSFilePromiseReceiverClass() NSFilePromiseReceiverClass {
	_NSFilePromiseReceiverClassOnce.Do(func() {
		_NSFilePromiseReceiverClass = NSFilePromiseReceiverClass{class: objc.GetClass("NSFilePromiseReceiver")}
	})
	return _NSFilePromiseReceiverClass
}

// GetNSFilePromiseReceiverClass returns the class object for NSFilePromiseReceiver.
func GetNSFilePromiseReceiverClass() NSFilePromiseReceiverClass {
	return getNSFilePromiseReceiverClass()
}

type NSFilePromiseReceiverClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSFilePromiseReceiverClass) Alloc() NSFilePromiseReceiver {
	rv := objc.Send[NSFilePromiseReceiver](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// An object that receives a file promise from the pasteboard.
//
// # Overview
// 
// Because [NSFilePromiseReceiver] implements the [NSPasteboardReading]
// protocol, you receive all file promises on the drag pasteboard as follows:
// 
// Likewise, you can enumerate dragged items by calling the following:
//
// # Instance Properties
//
//   - [NSFilePromiseReceiver.FileNames]: An array containing names of the promised files being written to the destination location.
//   - [NSFilePromiseReceiver.FileTypes]: An array containing types of the promised files being written to the destination location.
//
// # Instance Methods
//
//   - [NSFilePromiseReceiver.ReceivePromisedFilesAtDestinationOptionsOperationQueueReader]: Fulfills the promises at the specified destination.
//
// See: https://developer.apple.com/documentation/AppKit/NSFilePromiseReceiver
type NSFilePromiseReceiver struct {
	objectivec.Object
}

// NSFilePromiseReceiverFromID constructs a [NSFilePromiseReceiver] from an objc.ID.
//
// An object that receives a file promise from the pasteboard.
func NSFilePromiseReceiverFromID(id objc.ID) NSFilePromiseReceiver {
	return NSFilePromiseReceiver{objectivec.Object{id}}
}
// NOTE: NSFilePromiseReceiver adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [NSFilePromiseReceiver] class.
//
// # Instance Properties
//
//   - [INSFilePromiseReceiver.FileNames]: An array containing names of the promised files being written to the destination location.
//   - [INSFilePromiseReceiver.FileTypes]: An array containing types of the promised files being written to the destination location.
//
// # Instance Methods
//
//   - [INSFilePromiseReceiver.ReceivePromisedFilesAtDestinationOptionsOperationQueueReader]: Fulfills the promises at the specified destination.
//
// See: https://developer.apple.com/documentation/AppKit/NSFilePromiseReceiver
type INSFilePromiseReceiver interface {
	objectivec.IObject

	// Topic: Instance Properties

	// An array containing names of the promised files being written to the destination location.
	FileNames() []string
	// An array containing types of the promised files being written to the destination location.
	FileTypes() []string

	// Topic: Instance Methods

	// Fulfills the promises at the specified destination.
	ReceivePromisedFilesAtDestinationOptionsOperationQueueReader(destinationDir foundation.INSURL, options foundation.INSDictionary, operationQueue foundation.NSOperationQueue, reader URLErrorHandler)

	// Initializes an instance with a property list object and a type string.
	InitWithPasteboardPropertyListOfType(propertyList objectivec.IObject, type_ NSPasteboardType) NSFilePromiseReceiver
}





// Init initializes the instance.
func (f NSFilePromiseReceiver) Init() NSFilePromiseReceiver {
	rv := objc.Send[NSFilePromiseReceiver](f.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f NSFilePromiseReceiver) Autorelease() NSFilePromiseReceiver {
	rv := objc.Send[NSFilePromiseReceiver](f.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSFilePromiseReceiver creates a new NSFilePromiseReceiver instance.
func NewNSFilePromiseReceiver() NSFilePromiseReceiver {
	class := getNSFilePromiseReceiverClass()
	rv := objc.Send[NSFilePromiseReceiver](objc.ID(class.class), objc.Sel("new"))
	return rv
}






// Initializes an instance with a property list object and a type string.
//
// propertyList: A property list containing data to initialize the receiver.
// 
// By default, the property list object is an instance of [NSData]. If you
// implement [ReadingOptionsForTypePasteboard] and specify an option other
// than [PasteboardReadingAsData], the `propertyList` may be any other
// property list object.
//
// type: A UTI supported by the receiver for reading (one of the types returned by
// [ReadableTypesForPasteboard]).
//
// # Return Value
// 
// An object initialized using the data in `propertyList`.
//
// # Discussion
// 
// This method is considered optional because, if [ReadableTypesForPasteboard]
// returns just a single type, and that type uses the
// [PasteboardReadingAsKeyedArchive] reading option, then instances are
// initialized using [init(coder:)] instead of this method.
//
// [init(coder:)]: https://developer.apple.com/documentation/Foundation/NSCoding/init(coder:)
//
// See: https://developer.apple.com/documentation/AppKit/NSPasteboardReading/init(pasteboardPropertyList:ofType:)
func NewFilePromiseReceiverWithPasteboardPropertyListOfType(propertyList objectivec.IObject, type_ NSPasteboardType) NSFilePromiseReceiver {
	instance := getNSFilePromiseReceiverClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithPasteboardPropertyList:ofType:"), propertyList, objc.String(string(type_)))
	return NSFilePromiseReceiverFromID(rv)
}







// Fulfills the promises at the specified destination.
//
// destinationDir: The destination location URL of the file promise.
//
// options: An options dictionary to pass additional data.
//
// operationQueue: The operation queue on which to call the reader block when the promised
// file is ready.
//
// reader: A block to be called on the supplied operationQueue when the promised file
// is ready to be read.
//
// # Discussion
// 
// Call this method only when you’re accepting the file promise. All file
// promise receivers in a drag must specify the same destination location. The
// `options` dictionary is ignored for now. The `reader` block is called on
// the supplied `operationQueue` when the promised file is ready to be read.
// 
// Avoid blocking the main thread while waiting for the file promise to be
// written (which can be a long process) by specifying an operation queue
// other than the main queue. When the source is an [NSFilePromiseProvider],
// the reader block call is wrapped in a file coordination read.
//
// See: https://developer.apple.com/documentation/AppKit/NSFilePromiseReceiver/receivePromisedFiles(atDestination:options:operationQueue:reader:)
func (f NSFilePromiseReceiver) ReceivePromisedFilesAtDestinationOptionsOperationQueueReader(destinationDir foundation.INSURL, options foundation.INSDictionary, operationQueue foundation.NSOperationQueue, reader URLErrorHandler) {
		_block3, _cleanup3 := NewURLErrorBlock(reader)
	defer _cleanup3()
		objc.Send[objc.ID](f.ID, objc.Sel("receivePromisedFilesAtDestination:options:operationQueue:reader:"), destinationDir, options, operationQueue, _block3)
}

// Initializes an instance with a property list object and a type string.
//
// propertyList: A property list containing data to initialize the receiver.
// 
// By default, the property list object is an instance of [NSData]. If you
// implement [ReadingOptionsForTypePasteboard] and specify an option other
// than [PasteboardReadingAsData], the `propertyList` may be any other
// property list object.
//
// type: A UTI supported by the receiver for reading (one of the types returned by
// [ReadableTypesForPasteboard]).
//
// # Return Value
// 
// An object initialized using the data in `propertyList`.
//
// # Discussion
// 
// This method is considered optional because, if [ReadableTypesForPasteboard]
// returns just a single type, and that type uses the
// [PasteboardReadingAsKeyedArchive] reading option, then instances are
// initialized using [init(coder:)] instead of this method.
//
// [init(coder:)]: https://developer.apple.com/documentation/Foundation/NSCoding/init(coder:)
//
// See: https://developer.apple.com/documentation/AppKit/NSPasteboardReading/init(pasteboardPropertyList:ofType:)
func (f NSFilePromiseReceiver) InitWithPasteboardPropertyListOfType(propertyList objectivec.IObject, type_ NSPasteboardType) NSFilePromiseReceiver {
	rv := objc.Send[NSFilePromiseReceiver](f.ID, objc.Sel("initWithPasteboardPropertyList:ofType:"), propertyList, objc.String(string(type_)))
	return rv
}





// Returns an array of uniform type identifier strings of data types the
// receiver can read from the pasteboard and initialize from.
//
// pasteboard: A pasteboard. You can use the pasteboard argument to provide different
// types based on the pasteboard name, if you need to.
//
// # Return Value
// 
// An array of uniform type identifier strings of data types instances that
// the receiver can read from the pasteboard and initialize from.
//
// # Discussion
// 
// By default, the system provides the data for a type to
// [InitWithPasteboardPropertyListOfType] as an instance of [NSData]. If you
// implement [ReadingOptionsForTypePasteboard] and specify a different option,
// the system converts the [NSData] object for a type to an [NSString] object
// or any other property list object.
// 
// # Special Considerations
// 
// Don’t perform other pasteboard operations in the method implementation.
//
// See: https://developer.apple.com/documentation/AppKit/NSPasteboardReading/readableTypes(for:)
func (_NSFilePromiseReceiverClass NSFilePromiseReceiverClass) ReadableTypesForPasteboard(pasteboard INSPasteboard) []string {
	rv := objc.Send[[]objc.ID](objc.ID(_NSFilePromiseReceiverClass.class), objc.Sel("readableTypesForPasteboard:"), pasteboard)
	return objc.ConvertSliceToStrings(rv)
}

// Returns options for reading data of a specified type from a given
// pasteboard.
//
// type: A UTI supported by instances of the receiver for reading (one of the types
// returned by [ReadableTypesForPasteboard]).
//
// pasteboard: A pasteboard.
// 
// You can use the pasteboard argument to provide return different based on
// the pasteboard name, should you need to do so.
//
// # Return Value
// 
// Options for reading data of `type` from `pasteboard`. For a list of valid
// values, see [NSPasteboard.ReadingOptions].
//
// [NSPasteboard.ReadingOptions]: https://developer.apple.com/documentation/AppKit/NSPasteboard/ReadingOptions
//
// # Discussion
// 
// Do not perform other pasteboard operations in this method implementation.
//
// See: https://developer.apple.com/documentation/AppKit/NSPasteboardReading/readingOptions(forType:pasteboard:)
func (_NSFilePromiseReceiverClass NSFilePromiseReceiverClass) ReadingOptionsForTypePasteboard(type_ NSPasteboardType, pasteboard INSPasteboard) NSPasteboardReadingOptions {
	rv := objc.Send[NSPasteboardReadingOptions](objc.ID(_NSFilePromiseReceiverClass.class), objc.Sel("readingOptionsForType:pasteboard:"), objc.String(string(type_)), pasteboard)
	return NSPasteboardReadingOptions(rv)
}








// An array containing names of the promised files being written to the
// destination location.
//
// # Discussion
// 
// This property returns an empty array until the file promise is called using
// [ReceivePromisedFilesAtDestinationOptionsOperationQueueReader].
//
// See: https://developer.apple.com/documentation/AppKit/NSFilePromiseReceiver/fileNames
func (f NSFilePromiseReceiver) FileNames() []string {
	rv := objc.Send[[]objc.ID](f.ID, objc.Sel("fileNames"))
	return objc.ConvertSliceToStrings(rv)
}



// An array containing types of the promised files being written to the
// destination location.
//
// # Discussion
// 
// [NSFilePromiseProvider] promises one file type per item. The `count` of
// `fileTypes` should tell you the number of promised files in this item, but
// that’s not always guaranteed. Some legacy file promisers list each unique
// `fileType` only once.
//
// See: https://developer.apple.com/documentation/AppKit/NSFilePromiseReceiver/fileTypes
func (f NSFilePromiseReceiver) FileTypes() []string {
	rv := objc.Send[[]objc.ID](f.ID, objc.Sel("fileTypes"))
	return objc.ConvertSliceToStrings(rv)
}







// An array containing dragged file types that are readable.
//
// # Discussion
// 
// A view must register what types it accepts via [RegisterForDraggedTypes].
// Use that class method to get the file promise drag types that
// [NSFilePromiseReceiver] can accept, in order to register a view to accept
// promised files. [NSFilePromiseReceiver] can accept file promises from both
// the item-based [NSFilePromiseProvider] and the non-item based API. If you
// don’t register all these drag types, you might not be notified about some
// file promise drags. Register using the following code:
//
// See: https://developer.apple.com/documentation/AppKit/NSFilePromiseReceiver/readableDraggedTypes
func (_NSFilePromiseReceiverClass NSFilePromiseReceiverClass) ReadableDraggedTypes() []string {
	rv := objc.Send[[]objc.ID](objc.ID(_NSFilePromiseReceiverClass.class), objc.Sel("readableDraggedTypes"))
	return objc.ConvertSliceToStrings(rv)
}















// ReceivePromisedFilesAtDestinationOptionsOperationQueueReaderSync is a synchronous wrapper around [NSFilePromiseReceiver.ReceivePromisedFilesAtDestinationOptionsOperationQueueReader].
// It blocks until the completion handler fires or the context is cancelled.
func (f NSFilePromiseReceiver) ReceivePromisedFilesAtDestinationOptionsOperationQueueReaderSync(ctx context.Context, destinationDir foundation.INSURL, options foundation.INSDictionary, operationQueue foundation.NSOperationQueue) (*foundation.NSURL, error) {
	type result struct {
		val *foundation.NSURL
		err error
	}
	done := make(chan result, 1)
	f.ReceivePromisedFilesAtDestinationOptionsOperationQueueReader(destinationDir, options, operationQueue, func(val *foundation.NSURL, err error) {
		done <- result{val, err}
	})
	select {
	case r := <-done:
		return r.val, r.err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}






