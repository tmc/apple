// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"errors"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [FileHandle] class.
var (
	_FileHandleClass     FileHandleClass
	_FileHandleClassOnce sync.Once
)

func getFileHandleClass() FileHandleClass {
	_FileHandleClassOnce.Do(func() {
		_FileHandleClass = FileHandleClass{class: objc.GetClass("NSFileHandle")}
	})
	return _FileHandleClass
}

// GetFileHandleClass returns the class object for NSFileHandle.
func GetFileHandleClass() FileHandleClass {
	return getFileHandleClass()
}

type FileHandleClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (fc FileHandleClass) Class() objc.Class {
	return fc.class
}

// Alloc allocates memory for a new instance of the class.
func (fc FileHandleClass) Alloc() FileHandle {
	rv := objc.Send[FileHandle](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// An object-oriented wrapper for a file descriptor.
//
// # Overview
// 
// You use file handle objects to access data associated with files, sockets,
// pipes, and devices. For files, you can read, write, and seek within the
// file. For sockets, pipes, and devices, you can use a file handle object to
// monitor the device and process data asynchronously.
// 
// Most creation methods for [NSFileHandle] cause the file handle object to
// take ownership of the associated file descriptor. This means that the file
// handle object both creates the file descriptor and is responsible for
// closing it later, usually when the system deallocates the file handle
// object. If you want to use a file handle object with a file descriptor that
// you created, use the [InitWithFileDescriptor] method or use the
// [InitWithFileDescriptorCloseOnDealloc] method and pass [false] for the
// `flag` parameter.
// 
// # Run Loop Considerations
// 
// When using a file handle object to communicate asynchronously with a
// socket, you must initiate the corresponding operations from a thread with
// an active run loop. Although the read, accept, and wait operations
// themselves are performed asynchronously on background threads, the file
// handle uses a run loop source to monitor the operations and notify your
// code appropriately. Therefore, you must call those methods from your
// application’s main thread or from any thread where you’ve configured a
// run loop and are using it to process events.
//
// [false]: https://developer.apple.com/documentation/Swift/false
//
// # Creating a File Handle
//
//   - [FileHandle.InitWithFileDescriptor]: Creates and returns a file handle object associated with the specified file descriptor.
//   - [FileHandle.InitWithFileDescriptorCloseOnDealloc]: Creates and returns a file handle object associated with the specified file descriptor and deallocation policy.
//
// # Getting a File Descriptor
//
//   - [FileHandle.FileDescriptor]: The POSIX file descriptor associated with the receiver.
//
// # Reading from a File Handle Asynchronously
//
//   - [FileHandle.Bytes]: The file’s contents, as an asynchronous sequence of bytes.
//   - [FileHandle.SetBytes]
//
// # Reading from a File Handle Synchronously
//
//   - [FileHandle.AvailableData]: The data currently available in the receiver.
//
// # Reading Asynchronously with Notifications
//
//   - [FileHandle.AcceptConnectionInBackgroundAndNotify]: Accepts a socket connection (for stream-type sockets only) in the background and creates a file handle for the “near” (client) end of the communications channel.
//   - [FileHandle.AcceptConnectionInBackgroundAndNotifyForModes]: Accepts a socket connection (for stream-type sockets only) in the background and creates a file handle for the “near” (client) end of the communications channel.
//   - [FileHandle.ReadInBackgroundAndNotify]: Reads from the file or communications channel in the background and posts a notification when finished.
//   - [FileHandle.ReadInBackgroundAndNotifyForModes]: Reads from the file or communications channel in the background and posts a notification when finished.
//   - [FileHandle.ReadToEndOfFileInBackgroundAndNotify]: Reads to the end of file from the file or communications channel in the background and posts a notification when finished.
//   - [FileHandle.ReadToEndOfFileInBackgroundAndNotifyForModes]: Reads to the end of file from the file or communications channel in the background and posts a notification when finished.
//   - [FileHandle.WaitForDataInBackgroundAndNotify]: Asynchronously checks to see if data is available.
//   - [FileHandle.WaitForDataInBackgroundAndNotifyForModes]: Asynchronously checks to see if data is available.
//
// # Seeking Within a File
//
//   - [FileHandle.SeekToOffsetError]: Moves the file pointer to the specified offset within the file.
//
// # Operating on a File
//
//   - [FileHandle.CloseAndReturnError]: Disallows further access to the represented file or communications channel and signals end of file on communications channels that permit writing.
//   - [FileHandle.SynchronizeAndReturnError]: Causes all in-memory data and attributes of the file represented by the file handle to write to permanent storage.
//   - [FileHandle.TruncateAtOffsetError]: Truncates or extends the file represented by the file handle to a specified offset within the file and puts the file pointer at that position.
//
// # Monitoring for Readability and Writability
//
//   - [FileHandle.ReadabilityHandler]: The block to use for reading the contents of the file handle asynchronously.
//   - [FileHandle.SetReadabilityHandler]
//   - [FileHandle.WriteabilityHandler]: The block to use for writing the contents of the file handle asynchronously.
//   - [FileHandle.SetWriteabilityHandler]
//
// # Notifications
//
//   - [FileHandle.NSFileHandleConnectionAccepted]: Posted when a file handle object establishes a socket connection between two processes, creates a file handle object for one end of the connection, and makes this object available to observers.
//   - [FileHandle.NSFileHandleDataAvailable]: Posted when the file handle determines that data is currently available for reading in a file or at a communications channel.
//   - [FileHandle.NSFileHandleReadToEndOfFileCompletion]: Posted when the file handle reads all data in the file or, in a communications channel, until the other process signals the end of data.
//
// # Deprecated
//
//   - [FileHandle.OffsetInFile]: The position of the file pointer within the file represented by the file handle.
//   - [FileHandle.NSFileHandleNotificationMonitorModes]: Currently unused.
//
// See: https://developer.apple.com/documentation/Foundation/FileHandle
type FileHandle struct {
	objectivec.Object
}

// FileHandleFromID constructs a [FileHandle] from an objc.ID.
//
// An object-oriented wrapper for a file descriptor.
func FileHandleFromID(id objc.ID) FileHandle {
	return NSFileHandle{objectivec.Object{ID: id}}
}

// NSFileHandleFromID is an alias for [FileHandleFromID] for cross-framework compatibility.
func NSFileHandleFromID(id objc.ID) FileHandle { return FileHandleFromID(id) }
// NOTE: FileHandle adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [FileHandle] class.
//
// # Creating a File Handle
//
//   - [IFileHandle.InitWithFileDescriptor]: Creates and returns a file handle object associated with the specified file descriptor.
//   - [IFileHandle.InitWithFileDescriptorCloseOnDealloc]: Creates and returns a file handle object associated with the specified file descriptor and deallocation policy.
//
// # Getting a File Descriptor
//
//   - [IFileHandle.FileDescriptor]: The POSIX file descriptor associated with the receiver.
//
// # Reading from a File Handle Asynchronously
//
//   - [IFileHandle.Bytes]: The file’s contents, as an asynchronous sequence of bytes.
//   - [IFileHandle.SetBytes]
//
// # Reading from a File Handle Synchronously
//
//   - [IFileHandle.AvailableData]: The data currently available in the receiver.
//
// # Reading Asynchronously with Notifications
//
//   - [IFileHandle.AcceptConnectionInBackgroundAndNotify]: Accepts a socket connection (for stream-type sockets only) in the background and creates a file handle for the “near” (client) end of the communications channel.
//   - [IFileHandle.AcceptConnectionInBackgroundAndNotifyForModes]: Accepts a socket connection (for stream-type sockets only) in the background and creates a file handle for the “near” (client) end of the communications channel.
//   - [IFileHandle.ReadInBackgroundAndNotify]: Reads from the file or communications channel in the background and posts a notification when finished.
//   - [IFileHandle.ReadInBackgroundAndNotifyForModes]: Reads from the file or communications channel in the background and posts a notification when finished.
//   - [IFileHandle.ReadToEndOfFileInBackgroundAndNotify]: Reads to the end of file from the file or communications channel in the background and posts a notification when finished.
//   - [IFileHandle.ReadToEndOfFileInBackgroundAndNotifyForModes]: Reads to the end of file from the file or communications channel in the background and posts a notification when finished.
//   - [IFileHandle.WaitForDataInBackgroundAndNotify]: Asynchronously checks to see if data is available.
//   - [IFileHandle.WaitForDataInBackgroundAndNotifyForModes]: Asynchronously checks to see if data is available.
//
// # Seeking Within a File
//
//   - [IFileHandle.SeekToOffsetError]: Moves the file pointer to the specified offset within the file.
//
// # Operating on a File
//
//   - [IFileHandle.CloseAndReturnError]: Disallows further access to the represented file or communications channel and signals end of file on communications channels that permit writing.
//   - [IFileHandle.SynchronizeAndReturnError]: Causes all in-memory data and attributes of the file represented by the file handle to write to permanent storage.
//   - [IFileHandle.TruncateAtOffsetError]: Truncates or extends the file represented by the file handle to a specified offset within the file and puts the file pointer at that position.
//
// # Monitoring for Readability and Writability
//
//   - [IFileHandle.ReadabilityHandler]: The block to use for reading the contents of the file handle asynchronously.
//   - [IFileHandle.SetReadabilityHandler]
//   - [IFileHandle.WriteabilityHandler]: The block to use for writing the contents of the file handle asynchronously.
//   - [IFileHandle.SetWriteabilityHandler]
//
// # Notifications
//
//   - [IFileHandle.NSFileHandleConnectionAccepted]: Posted when a file handle object establishes a socket connection between two processes, creates a file handle object for one end of the connection, and makes this object available to observers.
//   - [IFileHandle.NSFileHandleDataAvailable]: Posted when the file handle determines that data is currently available for reading in a file or at a communications channel.
//   - [IFileHandle.NSFileHandleReadToEndOfFileCompletion]: Posted when the file handle reads all data in the file or, in a communications channel, until the other process signals the end of data.
//
// # Deprecated
//
//   - [IFileHandle.OffsetInFile]: The position of the file pointer within the file represented by the file handle.
//   - [IFileHandle.NSFileHandleNotificationMonitorModes]: Currently unused.
//
// See: https://developer.apple.com/documentation/Foundation/FileHandle
type IFileHandle interface {
	objectivec.IObject
	NSCoding
	NSSecureCoding

	// Topic: Creating a File Handle

	// Creates and returns a file handle object associated with the specified file descriptor.
	InitWithFileDescriptor(fd int) FileHandle
	// Creates and returns a file handle object associated with the specified file descriptor and deallocation policy.
	InitWithFileDescriptorCloseOnDealloc(fd int, closeopt bool) FileHandle

	// Topic: Getting a File Descriptor

	// The POSIX file descriptor associated with the receiver.
	FileDescriptor() int

	// Topic: Reading from a File Handle Asynchronously

	// The file’s contents, as an asynchronous sequence of bytes.
	Bytes() objectivec.IObject
	SetBytes(value objectivec.IObject)

	// Topic: Reading from a File Handle Synchronously

	// The data currently available in the receiver.
	AvailableData() INSData

	// Topic: Reading Asynchronously with Notifications

	// Accepts a socket connection (for stream-type sockets only) in the background and creates a file handle for the “near” (client) end of the communications channel.
	AcceptConnectionInBackgroundAndNotify()
	// Accepts a socket connection (for stream-type sockets only) in the background and creates a file handle for the “near” (client) end of the communications channel.
	AcceptConnectionInBackgroundAndNotifyForModes(modes []string)
	// Reads from the file or communications channel in the background and posts a notification when finished.
	ReadInBackgroundAndNotify()
	// Reads from the file or communications channel in the background and posts a notification when finished.
	ReadInBackgroundAndNotifyForModes(modes []string)
	// Reads to the end of file from the file or communications channel in the background and posts a notification when finished.
	ReadToEndOfFileInBackgroundAndNotify()
	// Reads to the end of file from the file or communications channel in the background and posts a notification when finished.
	ReadToEndOfFileInBackgroundAndNotifyForModes(modes []string)
	// Asynchronously checks to see if data is available.
	WaitForDataInBackgroundAndNotify()
	// Asynchronously checks to see if data is available.
	WaitForDataInBackgroundAndNotifyForModes(modes []string)

	// Topic: Seeking Within a File

	// Moves the file pointer to the specified offset within the file.
	SeekToOffsetError(offset uint64) (bool, error)

	// Topic: Operating on a File

	// Disallows further access to the represented file or communications channel and signals end of file on communications channels that permit writing.
	CloseAndReturnError() (bool, error)
	// Causes all in-memory data and attributes of the file represented by the file handle to write to permanent storage.
	SynchronizeAndReturnError() (bool, error)
	// Truncates or extends the file represented by the file handle to a specified offset within the file and puts the file pointer at that position.
	TruncateAtOffsetError(offset uint64) (bool, error)

	// Topic: Monitoring for Readability and Writability

	// The block to use for reading the contents of the file handle asynchronously.
	ReadabilityHandler() FileHandleHandler
	SetReadabilityHandler(value FileHandleHandler)
	// The block to use for writing the contents of the file handle asynchronously.
	WriteabilityHandler() FileHandleHandler
	SetWriteabilityHandler(value FileHandleHandler)

	// Topic: Notifications

	// Posted when a file handle object establishes a socket connection between two processes, creates a file handle object for one end of the connection, and makes this object available to observers.
	NSFileHandleConnectionAccepted() NSNotificationName
	// Posted when the file handle determines that data is currently available for reading in a file or at a communications channel.
	NSFileHandleDataAvailable() NSNotificationName
	// Posted when the file handle reads all data in the file or, in a communications channel, until the other process signals the end of data.
	NSFileHandleReadToEndOfFileCompletion() NSNotificationName

	// Topic: Deprecated

	// The position of the file pointer within the file represented by the file handle.
	OffsetInFile() uint64
	// Currently unused.
	NSFileHandleNotificationMonitorModes() string

	// Get the current position of the file pointer within the file.
	GetOffsetError() (uint64, error)
	// Reads the available data synchronously up to the end of file or maximum number of bytes.
	ReadDataToEndOfFileAndReturnError() (INSData, error)
	// Reads data synchronously up to the specified number of bytes.
	ReadDataUpToLengthError(length uint) (INSData, error)
	// Places the file pointer at the end of the file referenced by the file handle and returns the new file offset.
	SeekToEndReturningOffsetError() (uint64, error)
	// Writes the specified data synchronously to the file handle.
	WriteDataError(data INSData) (bool, error)
}

// Init initializes the instance.
func (f FileHandle) Init() FileHandle {
	rv := objc.Send[FileHandle](f.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f FileHandle) Autorelease() FileHandle {
	rv := objc.Send[FileHandle](f.ID, objc.Sel("autorelease"))
	return rv
}

// NewFileHandle creates a new FileHandle instance.
func NewFileHandle() FileHandle {
	class := getFileHandleClass()
	rv := objc.Send[FileHandle](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Returns a file handle initialized for reading the file, device, or named
// socket at the specified path.
//
// path: The path to the file, device, or named socket to access.
//
// # Return Value
// 
// The initialized file handle object or `nil` if no file exists at `path`.
//
// # Discussion
// 
// The system sets the file pointer to the beginning of the file. You can’t
// write data to the returned file handle object. Use the
// [ReadDataToEndOfFile] or [ReadDataOfLength] methods to read data from it.
// 
// When using this method to create a file handle object, the file handle owns
// its associated file descriptor and is responsible for closing it.
//
// See: https://developer.apple.com/documentation/Foundation/FileHandle/init(forReadingAtPath:)
func NewFileHandleForReadingAtPath(path string) FileHandle {
	rv := objc.Send[objc.ID](objc.ID(getFileHandleClass().class), objc.Sel("fileHandleForReadingAtPath:"), objc.String(path))
	return FileHandleFromID(rv)
}

//
// See: https://developer.apple.com/documentation/Foundation/FileHandle/init(forReadingFrom:)
func NewFileHandleForReadingFromURLError(url INSURL) (FileHandle, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(getFileHandleClass().class), objc.Sel("fileHandleForReadingFromURL:error:"), url, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return FileHandle{}, NSErrorFrom(errorPtr)
	}
	return FileHandleFromID(rv), nil
}

// Returns a file handle initialized for reading and writing to the file,
// device, or named socket at the specified path.
//
// path: The path to the file, device, or named socket to access.
//
// # Return Value
// 
// The initialized file handle object or `nil` if no file exists at `path`.
//
// # Discussion
// 
// The file pointer is set to the beginning of the file. The returned object
// responds to both `read...` messages and [WriteData].
// 
// When using this method to create a file handle object, the file handle owns
// its associated file descriptor and is responsible for closing it.
//
// See: https://developer.apple.com/documentation/Foundation/FileHandle/init(forUpdatingAtPath:)
func NewFileHandleForUpdatingAtPath(path string) FileHandle {
	rv := objc.Send[objc.ID](objc.ID(getFileHandleClass().class), objc.Sel("fileHandleForUpdatingAtPath:"), objc.String(path))
	return FileHandleFromID(rv)
}

//
// See: https://developer.apple.com/documentation/Foundation/FileHandle/init(forUpdating:)
func NewFileHandleForUpdatingURLError(url INSURL) (FileHandle, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(getFileHandleClass().class), objc.Sel("fileHandleForUpdatingURL:error:"), url, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return FileHandle{}, NSErrorFrom(errorPtr)
	}
	return FileHandleFromID(rv), nil
}

// Returns a file handle initialized for writing to the file, device, or named
// socket at the specified path.
//
// path: The path to the file, device, or named socket to access.
//
// # Return Value
// 
// The initialized file handle object or `nil` if no file exists at `path`.
//
// # Discussion
// 
// The file pointer is set to the beginning of the file. You cannot read data
// from the returned file handle object. Use the [WriteData] method to write
// data to the file handle.
// 
// When using this method to create a file handle object, the file handle owns
// its associated file descriptor and is responsible for closing it.
//
// See: https://developer.apple.com/documentation/Foundation/FileHandle/init(forWritingAtPath:)
func NewFileHandleForWritingAtPath(path string) FileHandle {
	rv := objc.Send[objc.ID](objc.ID(getFileHandleClass().class), objc.Sel("fileHandleForWritingAtPath:"), objc.String(path))
	return FileHandleFromID(rv)
}

//
// See: https://developer.apple.com/documentation/Foundation/FileHandle/init(forWritingTo:)
func NewFileHandleForWritingToURLError(url INSURL) (FileHandle, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(getFileHandleClass().class), objc.Sel("fileHandleForWritingToURL:error:"), url, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return FileHandle{}, NSErrorFrom(errorPtr)
	}
	return FileHandleFromID(rv), nil
}

// Returns a file handle initialized from data in an unarchiver.
//
// See: https://developer.apple.com/documentation/Foundation/FileHandle/init(coder:)
func NewFileHandleWithCoder(coder INSCoder) FileHandle {
	instance := getFileHandleClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return FileHandleFromID(rv)
}

// Creates and returns a file handle object associated with the specified file
// descriptor.
//
// fd: The POSIX file descriptor with which to initialize the file handle. This
// descriptor represents an open file or socket that you created previously.
// For example, when creating a file handle for a socket, you’d pass the
// value returned by the socket function.
//
// # Return Value
// 
// A file handle initialized with `fileDescriptor`.
//
// # Discussion
// 
// The file descriptor you pass in to this method isn’t owned by the file
// handle object. Therefore, you’re responsible for closing the file
// descriptor at some point after disposing of the file handle object.
// 
// You can create a file handle for a socket by using the result of a `socket`
// call as `fileDescriptor`.
//
// See: https://developer.apple.com/documentation/Foundation/FileHandle/init(fileDescriptor:)
func NewFileHandleWithFileDescriptor(fd int) FileHandle {
	instance := getFileHandleClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFileDescriptor:"), fd)
	return FileHandleFromID(rv)
}

// Creates and returns a file handle object associated with the specified file
// descriptor and deallocation policy.
//
// fd: The POSIX file descriptor with which to initialize the file handle.
//
// closeopt: [true] if the returned file handle object should take ownership of the file
// descriptor and close it for you or [false] if you want to maintain
// ownership of the file descriptor.
// //
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Return Value
// 
// An initialized file handle object.
//
// # Discussion
// 
// If `flag` is [false], the file descriptor you pass in to this method
// isn’t owned by the file handle object. In such a case, you’re
// responsible for closing the file descriptor at some point after disposing
// of the file handle object. If you want the file handle object to close the
// descriptor for you automatically, pass [true] for the `flag` parameter.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Foundation/FileHandle/init(fileDescriptor:closeOnDealloc:)
func NewFileHandleWithFileDescriptorCloseOnDealloc(fd int, closeopt bool) FileHandle {
	instance := getFileHandleClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFileDescriptor:closeOnDealloc:"), fd, closeopt)
	return FileHandleFromID(rv)
}

// Creates and returns a file handle object associated with the specified file
// descriptor.
//
// fd: The POSIX file descriptor with which to initialize the file handle. This
// descriptor represents an open file or socket that you created previously.
// For example, when creating a file handle for a socket, you’d pass the
// value returned by the socket function.
//
// # Return Value
// 
// A file handle initialized with `fileDescriptor`.
//
// # Discussion
// 
// The file descriptor you pass in to this method isn’t owned by the file
// handle object. Therefore, you’re responsible for closing the file
// descriptor at some point after disposing of the file handle object.
// 
// You can create a file handle for a socket by using the result of a `socket`
// call as `fileDescriptor`.
//
// See: https://developer.apple.com/documentation/Foundation/FileHandle/init(fileDescriptor:)
func (f FileHandle) InitWithFileDescriptor(fd int) FileHandle {
	rv := objc.Send[FileHandle](f.ID, objc.Sel("initWithFileDescriptor:"), fd)
	return rv
}
// Creates and returns a file handle object associated with the specified file
// descriptor and deallocation policy.
//
// fd: The POSIX file descriptor with which to initialize the file handle.
//
// closeopt: [true] if the returned file handle object should take ownership of the file
// descriptor and close it for you or [false] if you want to maintain
// ownership of the file descriptor.
// //
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Return Value
// 
// An initialized file handle object.
//
// # Discussion
// 
// If `flag` is [false], the file descriptor you pass in to this method
// isn’t owned by the file handle object. In such a case, you’re
// responsible for closing the file descriptor at some point after disposing
// of the file handle object. If you want the file handle object to close the
// descriptor for you automatically, pass [true] for the `flag` parameter.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Foundation/FileHandle/init(fileDescriptor:closeOnDealloc:)
func (f FileHandle) InitWithFileDescriptorCloseOnDealloc(fd int, closeopt bool) FileHandle {
	rv := objc.Send[FileHandle](f.ID, objc.Sel("initWithFileDescriptor:closeOnDealloc:"), fd, closeopt)
	return rv
}
// Returns a file handle initialized from data in an unarchiver.
//
// See: https://developer.apple.com/documentation/Foundation/FileHandle/init(coder:)
func (f FileHandle) InitWithCoder(coder INSCoder) FileHandle {
	rv := objc.Send[FileHandle](f.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}
// Accepts a socket connection (for stream-type sockets only) in the
// background and creates a file handle for the “near” (client) end of the
// communications channel.
//
// # Discussion
// 
// This method asynchronously creates a file handle for the other end of the
// socket connection and returns that object by posting a
// [NSFileHandleConnectionAccepted] notification in the current thread. The
// notification includes a `userInfo` dictionary with the created
// [NSFileHandle] object, which is accessible using the
// [NSFileHandleNotificationFileHandleItem] key.
// 
// You must call this method from a thread that has an active run loop.
// 
// # Special Considerations
// 
// The receiver must be created by an [InitWithFileDescriptor] message that
// takes as an argument a stream-type socket created by the appropriate system
// routine, . In other words, you must `bind()` the socket, and ensure that
// the socket has a connection backlog defined by `listen()`.
// 
// The object that will write data to the returned file handle must add itself
// as an observer of [NSFileHandleConnectionAccepted].
// 
// Note that this method does not continue to listen for connection requests
// after it posts [NSFileHandleConnectionAccepted]. If you want to keep
// getting notified, you need to call [AcceptConnectionInBackgroundAndNotify]
// again in your observer method.
//
// [NSFileHandleConnectionAccepted]: https://developer.apple.com/documentation/Foundation/NSNotification/Name-swift.struct/NSFileHandleConnectionAccepted
//
// See: https://developer.apple.com/documentation/Foundation/FileHandle/acceptConnectionInBackgroundAndNotify()
func (f FileHandle) AcceptConnectionInBackgroundAndNotify() {
	objc.Send[objc.ID](f.ID, objc.Sel("acceptConnectionInBackgroundAndNotify"))
}
// Accepts a socket connection (for stream-type sockets only) in the
// background and creates a file handle for the “near” (client) end of the
// communications channel.
//
// modes: The runloop modes in which the connection accepted notification can be
// posted.
//
// # Discussion
// 
// See [AcceptConnectionInBackgroundAndNotify] for details of how this method
// operates. This method differs from [AcceptConnectionInBackgroundAndNotify]
// in that `modes` specifies the run-loop mode (or modes) in which
// [NSFileHandleConnectionAccepted] can be posted.
// 
// You must call this method from a thread that has an active run loop.
//
// [NSFileHandleConnectionAccepted]: https://developer.apple.com/documentation/Foundation/NSNotification/Name-swift.struct/NSFileHandleConnectionAccepted
//
// See: https://developer.apple.com/documentation/Foundation/FileHandle/acceptConnectionInBackgroundAndNotify(forModes:)
func (f FileHandle) AcceptConnectionInBackgroundAndNotifyForModes(modes []string) {
	objc.Send[objc.ID](f.ID, objc.Sel("acceptConnectionInBackgroundAndNotifyForModes:"), objectivec.StringSliceToNSArray(modes))
}
// Reads from the file or communications channel in the background and posts a
// notification when finished.
//
// # Discussion
// 
// This method performs an asynchronous [AvailableData] operation on a file or
// communications channel and posts an [readCompletionNotification]
// notification on the current thread when that operation is complete. You
// must call this method from a thread that has an active run loop.
// 
// The length of the data is limited to the buffer size of the underlying
// operating system. The notification includes a `userInfo` dictionary that
// contains the data read; access this object using the
// [NSFileHandleNotificationDataItem] key.
// 
// Any object interested in receiving this data asynchronously must add itself
// as an observer of [readCompletionNotification]. In communication via
// stream-type sockets, the receiver is often the object returned in the
// `userInfo` dictionary of [NSFileHandleConnectionAccepted].
// 
// Note that this method does not cause a continuous stream of notifications
// to be sent. If you wish to keep getting notified, you’ll also need to
// call [ReadInBackgroundAndNotify] in your observer method.
//
// [NSFileHandleConnectionAccepted]: https://developer.apple.com/documentation/Foundation/NSNotification/Name-swift.struct/NSFileHandleConnectionAccepted
// [readCompletionNotification]: https://developer.apple.com/documentation/Foundation/FileHandle/readCompletionNotification
//
// See: https://developer.apple.com/documentation/Foundation/FileHandle/readInBackgroundAndNotify()
func (f FileHandle) ReadInBackgroundAndNotify() {
	objc.Send[objc.ID](f.ID, objc.Sel("readInBackgroundAndNotify"))
}
// Reads from the file or communications channel in the background and posts a
// notification when finished.
//
// modes: The runloop modes in which the read completion notification can be posted.
//
// # Discussion
// 
// See [ReadInBackgroundAndNotify] for details of how this method operates.
// This method differs from [ReadInBackgroundAndNotify] in that `modes`
// specifies the run-loop mode (or modes) in which
// [readCompletionNotification] can be posted.
// 
// You must call this method from a thread that has an active run loop.
//
// [readCompletionNotification]: https://developer.apple.com/documentation/Foundation/FileHandle/readCompletionNotification
//
// See: https://developer.apple.com/documentation/Foundation/FileHandle/readInBackgroundAndNotify(forModes:)
func (f FileHandle) ReadInBackgroundAndNotifyForModes(modes []string) {
	objc.Send[objc.ID](f.ID, objc.Sel("readInBackgroundAndNotifyForModes:"), objectivec.StringSliceToNSArray(modes))
}
// Reads to the end of file from the file or communications channel in the
// background and posts a notification when finished.
//
// # Discussion
// 
// This method performs an asynchronous `readToEndOfFile` operation on a file
// or communications channel and posts an
// [NSFileHandleReadToEndOfFileCompletion]. You must call this method from a
// thread that has an active run loop.
// 
// The notification includes a `userInfo` dictionary that contains the data
// read; access this object using the [NSFileHandleNotificationDataItem] key.
// 
// Any object interested in receiving this data asynchronously must add itself
// as an observer of [NSFileHandleReadToEndOfFileCompletion]. In communication
// via stream-type sockets, the receiver is often the object returned in the
// `userInfo` dictionary of [NSFileHandleConnectionAccepted].
//
// [NSFileHandleConnectionAccepted]: https://developer.apple.com/documentation/Foundation/NSNotification/Name-swift.struct/NSFileHandleConnectionAccepted
// [NSFileHandleReadToEndOfFileCompletion]: https://developer.apple.com/documentation/Foundation/NSNotification/Name-swift.struct/NSFileHandleReadToEndOfFileCompletion
//
// See: https://developer.apple.com/documentation/Foundation/FileHandle/readToEndOfFileInBackgroundAndNotify()
func (f FileHandle) ReadToEndOfFileInBackgroundAndNotify() {
	objc.Send[objc.ID](f.ID, objc.Sel("readToEndOfFileInBackgroundAndNotify"))
}
// Reads to the end of file from the file or communications channel in the
// background and posts a notification when finished.
//
// modes: The runloop modes in which the read completion notification can be posted.
//
// # Discussion
// 
// See [ReadToEndOfFileInBackgroundAndNotify] for details of this method’s
// operation. The method differs from [ReadToEndOfFileInBackgroundAndNotify]
// in that `modes` specifies the run-loop mode (or modes) in which
// [NSFileHandleReadToEndOfFileCompletion] can be posted.
// 
// You must call this method from a thread that has an active run loop.
//
// [NSFileHandleReadToEndOfFileCompletion]: https://developer.apple.com/documentation/Foundation/NSNotification/Name-swift.struct/NSFileHandleReadToEndOfFileCompletion
//
// See: https://developer.apple.com/documentation/Foundation/FileHandle/readToEndOfFileInBackgroundAndNotify(forModes:)
func (f FileHandle) ReadToEndOfFileInBackgroundAndNotifyForModes(modes []string) {
	objc.Send[objc.ID](f.ID, objc.Sel("readToEndOfFileInBackgroundAndNotifyForModes:"), objectivec.StringSliceToNSArray(modes))
}
// Asynchronously checks to see if data is available.
//
// # Discussion
// 
// When the data becomes available, this method posts a
// [NSFileHandleDataAvailable] notification on the current thread.
// 
// You must call this method from a thread that has an active run loop.
//
// [NSFileHandleDataAvailable]: https://developer.apple.com/documentation/Foundation/NSNotification/Name-swift.struct/NSFileHandleDataAvailable
//
// See: https://developer.apple.com/documentation/Foundation/FileHandle/waitForDataInBackgroundAndNotify()
func (f FileHandle) WaitForDataInBackgroundAndNotify() {
	objc.Send[objc.ID](f.ID, objc.Sel("waitForDataInBackgroundAndNotify"))
}
// Asynchronously checks to see if data is available.
//
// modes: The runloop modes in which the data available notification can be posted.
//
// # Discussion
// 
// When the data becomes available, this method posts a
// [NSFileHandleDataAvailable] notification on the current thread. This method
// differs from [WaitForDataInBackgroundAndNotify] in that `modes` specifies
// the run-loop mode (or modes) in which [NSFileHandleDataAvailable] can be
// posted.
// 
// You must call this method from a thread that has an active run loop.
//
// [NSFileHandleDataAvailable]: https://developer.apple.com/documentation/Foundation/NSNotification/Name-swift.struct/NSFileHandleDataAvailable
//
// See: https://developer.apple.com/documentation/Foundation/FileHandle/waitForDataInBackgroundAndNotify(forModes:)
func (f FileHandle) WaitForDataInBackgroundAndNotifyForModes(modes []string) {
	objc.Send[objc.ID](f.ID, objc.Sel("waitForDataInBackgroundAndNotifyForModes:"), objectivec.StringSliceToNSArray(modes))
}
// Moves the file pointer to the specified offset within the file.
//
// offset: The offset to seek to.
//
// # Discussion
//
// See: https://developer.apple.com/documentation/Foundation/FileHandle/seek(toOffset:)
func (f FileHandle) SeekToOffsetError(offset uint64) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](f.ID, objc.Sel("seekToOffset:error:"), offset, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("seekToOffset:error: returned NO with nil NSError")
	}
	return rv, nil

}
// Disallows further access to the represented file or communications channel
// and signals end of file on communications channels that permit writing.
//
// # Discussion
// 
// If the file handle object owns its file descriptor, it automatically closes
// that descriptor when deallocated. If you initialized the file handle object
// using the [InitWithFileDescriptor] method, or you initialized it using the
// [InitWithFileDescriptorCloseOnDealloc] and passed [false] for the `flag`
// parameter, you can use this method to close the file descriptor; otherwise,
// you must close the file descriptor yourself.
// 
// After calling this method, you may still use the file handle object, but
// you must not attempt to read or write data or use the object to operate on
// the file descriptor. Attempts to read or write a closed file descriptor
// raise an exception.
//
// [false]: https://developer.apple.com/documentation/Swift/false
//
// See: https://developer.apple.com/documentation/Foundation/FileHandle/close()
func (f FileHandle) CloseAndReturnError() (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](f.ID, objc.Sel("closeAndReturnError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("closeAndReturnError: returned NO with nil NSError")
	}
	return rv, nil

}
// Causes all in-memory data and attributes of the file represented by the
// file handle to write to permanent storage.
//
// # Discussion
// 
// Programs that require the file to always be in a known state should call
// this method. An invocation of this method doesn’t return until memory is
// flushed.
//
// See: https://developer.apple.com/documentation/Foundation/FileHandle/synchronize()
func (f FileHandle) SynchronizeAndReturnError() (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](f.ID, objc.Sel("synchronizeAndReturnError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("synchronizeAndReturnError: returned NO with nil NSError")
	}
	return rv, nil

}
// Truncates or extends the file represented by the file handle to a specified
// offset within the file and puts the file pointer at that position.
//
// offset: The offset within the file that marks the new end of the file.
//
// # Discussion
// 
// If the file is extended (if `offset` is beyond the current end of file),
// the added characters are null bytes.
//
// See: https://developer.apple.com/documentation/Foundation/FileHandle/truncate(atOffset:)
func (f FileHandle) TruncateAtOffsetError(offset uint64) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](f.ID, objc.Sel("truncateAtOffset:error:"), offset, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("truncateAtOffset:error: returned NO with nil NSError")
	}
	return rv, nil

}
// Encodes the receiver using a given archiver.
//
// coder: An archiver object.
//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/encode(with:)
func (f FileHandle) EncodeWithCoder(coder INSCoder) {
	objc.Send[objc.ID](f.ID, objc.Sel("encodeWithCoder:"), coder)
}
// Get the current position of the file pointer within the file.
//
// offsetInFile: When the return value is [true], this provides the current position of the
// file pointer within the file.
// //
// [true]: https://developer.apple.com/documentation/Swift/true
//
// error: When the return value is [false], this provides an [NSError] indicating why
// the operation failed.
// //
// [false]: https://developer.apple.com/documentation/Swift/false
//
// # Return Value
// 
// Returns [false] when there was an error. Otherwise, returns [true] and sets
// the `offsetInFile` parameter’s pointee to the current position of the
// file pointer within the file.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// Returns [false] if called on a file handle representing a pipe or socket,
// or if the file descriptor is closed.
//
// [false]: https://developer.apple.com/documentation/Swift/false
//
// See: https://developer.apple.com/documentation/Foundation/NSFileHandle/getOffset:error:
func (f FileHandle) GetOffsetError() (uint64, error) {
	var offsetInFile uint64
	var errorPtr objc.ID
	rv := objc.Send[bool](f.ID, objc.Sel("getOffset:error:"), unsafe.Pointer(&offsetInFile), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return 0, NSErrorFrom(errorPtr)
	}
	if !rv {
		return 0, errors.New("getOffset:error: returned NO with nil NSError")
	}
	return offsetInFile, nil
}
// Reads the available data synchronously up to the end of file or maximum
// number of bytes.
//
// error: When the return value is `nil`, this provides an [NSError] indicating why
// the read operation failed.
//
// # Return Value
// 
// The data available through the file handle up to the maximum size that can
// be represented by an [NSData] object or, if a communications channel, until
// an end-of-file indicator is returned.
//
// # Discussion
// 
// This method invokes [ReadDataOfLength] as part of its implementation.
//
// See: https://developer.apple.com/documentation/Foundation/NSFileHandle/readDataToEndOfFileAndReturnError:
func (f FileHandle) ReadDataToEndOfFileAndReturnError() (INSData, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](f.ID, objc.Sel("readDataToEndOfFileAndReturnError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return NSData{}, NSErrorFrom(errorPtr)
	}
	return NSDataFromID(rv), nil

}
// Reads data synchronously up to the specified number of bytes.
//
// length: The number of bytes to read from the file handle.
//
// error: When the return value is `nil`, this provides an [NSError] indicating why
// the read operation failed.
//
// # Return Value
// 
// The data available through the receiver up to a maximum of `length` bytes,
// or the maximum size that can be represented by an [NSData] object,
// whichever is the smaller.
//
// # Discussion
// 
// If the handle represents a file, this method returns the data obtained by
// reading `length` bytes starting at the current file pointer. If `length`
// bytes aren’t available, this method returns the data from the current
// file pointer to the end of the file. If the handle is a communications
// channel, the method reads up to `length` bytes from the channel. Returns an
// empty [NSData] object if the handle is at the file’s end or if the
// communications channel returns an end-of-file indicator.
// 
// This method provides an error if attempts to determine the file-handle type
// fail or if attempts to read from the file or channel fail.
//
// See: https://developer.apple.com/documentation/Foundation/NSFileHandle/readDataUpToLength:error:
func (f FileHandle) ReadDataUpToLengthError(length uint) (INSData, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](f.ID, objc.Sel("readDataUpToLength:error:"), length, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return NSData{}, NSErrorFrom(errorPtr)
	}
	return NSDataFromID(rv), nil

}
// Places the file pointer at the end of the file referenced by the file
// handle and returns the new file offset.
//
// offsetInFile: When the return value is [true], this provides the file pointer’s offset
// at the end of the file. This should therefore equal to the size of the
// file.
// //
// [true]: https://developer.apple.com/documentation/Swift/true
//
// error: When the return value is [false], this provides an [NSError] indicating why
// the operation failed.
// //
// [false]: https://developer.apple.com/documentation/Swift/false
//
// # Return Value
// 
// Returns [false] when there was an error. Otherwise, returns [true] and sets
// the `offsetInFile` parameter’s pointee to the current position of the
// file pointer within the file.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// Returns [false] if called on a file handle representing a pipe or socket,
// or if the file descriptor is closed.
//
// [false]: https://developer.apple.com/documentation/Swift/false
//
// See: https://developer.apple.com/documentation/Foundation/NSFileHandle/seekToEndReturningOffset:error:
func (f FileHandle) SeekToEndReturningOffsetError() (uint64, error) {
	var offsetInFile uint64
	var errorPtr objc.ID
	rv := objc.Send[bool](f.ID, objc.Sel("seekToEndReturningOffset:error:"), unsafe.Pointer(&offsetInFile), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return 0, NSErrorFrom(errorPtr)
	}
	if !rv {
		return 0, errors.New("seekToEndReturningOffset:error: returned NO with nil NSError")
	}
	return offsetInFile, nil
}
// Writes the specified data synchronously to the file handle.
//
// data: The data to write to the file handle.
//
// error: When the return value is [false], this provides an [NSError] indicating why
// the write operation failed.
// //
// [false]: https://developer.apple.com/documentation/Swift/false
//
// # Return Value
// 
// Returns [true] when the data is successfullly written to the file handle.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// If the handle represents a file, writing takes place at the file
// pointer’s current position. After it writes the data, the method advances
// the file pointer by the number of bytes written. This method provides an
// error if the file descriptor is closed or isn’t valid, if the handle
// represents an unconnected pipe or socket endpoint, if there isn’t any
// free space on the file system, or if any other writing error occurs.
//
// See: https://developer.apple.com/documentation/Foundation/NSFileHandle/writeData:error:
func (f FileHandle) WriteDataError(data INSData) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](f.ID, objc.Sel("writeData:error:"), data, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("writeData:error: returned NO with nil NSError")
	}
	return rv, nil

}

// The POSIX file descriptor associated with the receiver.
//
// # Discussion
// 
// You can use this method to retrieve the file descriptor while it is open.
// If the file handle object owns the file descriptor, you must not close it
// yourself. However, you can use the [CloseFile] method to close the file
// descriptor programmatically. If you do call the [CloseFile] method,
// subsequent calls to this method raise an exception.
//
// See: https://developer.apple.com/documentation/Foundation/FileHandle/fileDescriptor
func (f FileHandle) FileDescriptor() int {
	rv := objc.Send[int](f.ID, objc.Sel("fileDescriptor"))
	return rv
}
// The file’s contents, as an asynchronous sequence of bytes.
//
// See: https://developer.apple.com/documentation/foundation/filehandle/bytes
func (f FileHandle) Bytes() objectivec.IObject {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("bytes"))
	return objectivec.Object{ID: rv}
}
func (f FileHandle) SetBytes(value objectivec.IObject) {
	objc.Send[struct{}](f.ID, objc.Sel("setBytes:"), value)
}
// The data currently available in the receiver.
//
// # Discussion
// 
// The data currently available through the receiver, up to the maximum size
// that can be represented by an [NSData] object.
// 
// If the receiver is a file, this method returns the data obtained by reading
// the file from the current file pointer to the end of the file. If the
// receiver is a communications channel, this method reads up to a buffer of
// data and returns it; if no data is available, the method blocks. Returns an
// empty data object if the end of file is reached. This method raises
// [NSFileHandleOperationException] if attempts to determine the file-handle
// type fail or if attempts to read from the file or channel fail.
//
// See: https://developer.apple.com/documentation/Foundation/FileHandle/availableData
func (f FileHandle) AvailableData() INSData {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("availableData"))
	return NSDataFromID(objc.ID(rv))
}
// The block to use for reading the contents of the file handle
// asynchronously.
//
// # Discussion
// 
// The default value of this property is `nil`. Assigning a valid [Block
// object] to this property creates a dispatch source for reading the contents
// of the file or socket. Your block is submitted to the file handle’s
// dispatch queue when there is data to read. When reading a file, your
// handler block is typically executed repeatedly until the entire contents of
// the file have been read. When reading data from a socket, your handler
// block is executed whenever there is data on the socket waiting to be read.
// 
// The block you provide must accept a single parameter that is the current
// file handle. The return type of your block should be `void`.
// 
// To stop reading the file or socket, set the value of this property to
// `nil`. Doing so cancels the dispatch source and cleans up the file
// handle’s structures appropriately.
//
// [Block object]: https://developer.apple.com/library/archive/documentation/General/Conceptual/DevPedia-CocoaCore/Block.html#//apple_ref/doc/uid/TP40008195-CH3
//
// See: https://developer.apple.com/documentation/Foundation/FileHandle/readabilityHandler
func (f FileHandle) ReadabilityHandler() FileHandleHandler {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("readabilityHandler"))
	_ = rv
	return nil
}
func (f FileHandle) SetReadabilityHandler(value FileHandleHandler) {
	block, cleanup := NewFileHandleBlock(value)
	defer cleanup()
	objc.Send[struct{}](f.ID, objc.Sel("setReadabilityHandler:"), block)
}
// The block to use for writing the contents of the file handle
// asynchronously.
//
// # Discussion
// 
// The default value of this property is `nil`. Assigning a valid [Block
// object] to this property creates a dispatch source for writing the contents
// of the file or socket. Your block is submitted to the file handle’s
// dispatch queue when there is room available to write more data. When
// writing a file, your handler block is typically executed repeatedly until
// the entire contents of the file have been written. When writing data to a
// socket, your handler block is executed whenever the socket is ready to
// accept more data.
// 
// The block you provide must accept a single parameter that is the current
// file handle. The return type of your block should be `void`.
// 
// To stop writing data to the file or socket, set the value of this property
// to `nil`. Doing so cancels the dispatch source and cleans up the file
// handle’s structures appropriately.
//
// [Block object]: https://developer.apple.com/library/archive/documentation/General/Conceptual/DevPedia-CocoaCore/Block.html#//apple_ref/doc/uid/TP40008195-CH3
//
// See: https://developer.apple.com/documentation/Foundation/FileHandle/writeabilityHandler
func (f FileHandle) WriteabilityHandler() FileHandleHandler {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("writeabilityHandler"))
	_ = rv
	return nil
}
func (f FileHandle) SetWriteabilityHandler(value FileHandleHandler) {
	block, cleanup := NewFileHandleBlock(value)
	defer cleanup()
	objc.Send[struct{}](f.ID, objc.Sel("setWriteabilityHandler:"), block)
}
// Posted when a file handle object establishes a socket connection between
// two processes, creates a file handle object for one end of the connection,
// and makes this object available to observers.
//
// See: https://developer.apple.com/documentation/foundation/nsnotification/name-swift.struct/nsfilehandleconnectionaccepted
func (f FileHandle) NSFileHandleConnectionAccepted() NSNotificationName {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("NSFileHandleConnectionAcceptedNotification"))
	return NSNotificationName(NSStringFromID(rv).String())
}
// Posted when the file handle determines that data is currently available for
// reading in a file or at a communications channel.
//
// See: https://developer.apple.com/documentation/foundation/nsnotification/name-swift.struct/nsfilehandledataavailable
func (f FileHandle) NSFileHandleDataAvailable() NSNotificationName {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("NSFileHandleDataAvailableNotification"))
	return NSNotificationName(NSStringFromID(rv).String())
}
// Posted when the file handle reads all data in the file or, in a
// communications channel, until the other process signals the end of data.
//
// See: https://developer.apple.com/documentation/foundation/nsnotification/name-swift.struct/nsfilehandlereadtoendoffilecompletion
func (f FileHandle) NSFileHandleReadToEndOfFileCompletion() NSNotificationName {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("NSFileHandleReadToEndOfFileCompletionNotification"))
	return NSNotificationName(NSStringFromID(rv).String())
}
// The position of the file pointer within the file represented by the file
// handle.
//
// # Discussion
// 
// Raises [fileHandleOperationException] if called on a file handle
// representing a pipe or socket, or if the file descriptor is closed.
//
// [fileHandleOperationException]: https://developer.apple.com/documentation/Foundation/NSExceptionName/fileHandleOperationException
//
// See: https://developer.apple.com/documentation/Foundation/FileHandle/offsetInFile
func (f FileHandle) OffsetInFile() uint64 {
	rv := objc.Send[uint64](f.ID, objc.Sel("offsetInFile"))
	return rv
}
// Currently unused.
//
// See: https://developer.apple.com/documentation/foundation/nsfilehandlenotificationmonitormodes
func (f FileHandle) NSFileHandleNotificationMonitorModes() string {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("NSFileHandleNotificationMonitorModes"))
	return NSStringFromID(rv).String()
}

// The file handle associated with the standard error file.
//
// # Return Value
// 
// The shared file handle associated with the standard error file.
// 
// # Discussion
// 
// Conventionally this is a terminal device where the system sends error
// messages. There’s one standard error file handle per process; it’s a
// shared instance.
// 
// When using this method to create a file handle object, the file handle owns
// its associated file descriptor and is responsible for closing it.
//
// See: https://developer.apple.com/documentation/Foundation/FileHandle/standardError
func (_FileHandleClass FileHandleClass) FileHandleWithStandardError() FileHandle {
	rv := objc.Send[objc.ID](objc.ID(_FileHandleClass.class), objc.Sel("fileHandleWithStandardError"))
	return NSFileHandleFromID(objc.ID(rv))
}
// The file handle associated with the standard input file.
//
// # Return Value
// 
// The shared file handle associated with the standard input file.
// 
// # Discussion
// 
// Conventionally this is a terminal device on which the user enters a stream
// of data. There’s one standard input file handle per process; it’s a
// shared instance.
// 
// When using this method to create a file handle object, the file handle owns
// its associated file descriptor and is responsible for closing it.
//
// See: https://developer.apple.com/documentation/Foundation/FileHandle/standardInput
func (_FileHandleClass FileHandleClass) FileHandleWithStandardInput() FileHandle {
	rv := objc.Send[objc.ID](objc.ID(_FileHandleClass.class), objc.Sel("fileHandleWithStandardInput"))
	return NSFileHandleFromID(objc.ID(rv))
}
// The file handle associated with the standard output file.
//
// # Return Value
// 
// The shared file handle associated with the standard output file.
// 
// # Discussion
// 
// Conventionally this is a terminal device that receives a stream of data
// from a program. There’s one standard output file handle per process;
// it’s a shared instance.
// 
// When using this method to create a file handle object, the file handle owns
// its associated file descriptor and is responsible for closing it.
//
// See: https://developer.apple.com/documentation/Foundation/FileHandle/standardOutput
func (_FileHandleClass FileHandleClass) FileHandleWithStandardOutput() FileHandle {
	rv := objc.Send[objc.ID](objc.ID(_FileHandleClass.class), objc.Sel("fileHandleWithStandardOutput"))
	return NSFileHandleFromID(objc.ID(rv))
}
// The file handle associated with a null device.
//
// # Return Value
// 
// A file handle associated with a null device.
// 
// # Discussion
// 
// You can use null-device file handles as “placeholders” for
// standard-device file handles or in collection objects to avoid exceptions
// and other errors resulting from messages being sent to invalid file
// handles. Read messages sent to a null-device file handle return an
// end-of-file indicator (an empty [NSData] object) rather than raise an
// exception. Write messages are no-ops, whereas [FileDescriptor] returns an
// illegal value. Other methods are no-ops or return “sensible” values.
// 
// When using this method to create a file handle object, the file handle owns
// its associated file descriptor and is responsible for closing it.
//
// See: https://developer.apple.com/documentation/Foundation/FileHandle/nullDevice
func (_FileHandleClass FileHandleClass) FileHandleWithNullDevice() FileHandle {
	rv := objc.Send[objc.ID](objc.ID(_FileHandleClass.class), objc.Sel("fileHandleWithNullDevice"))
	return NSFileHandleFromID(objc.ID(rv))
}

			// Protocol methods for NSSecureCoding
			

