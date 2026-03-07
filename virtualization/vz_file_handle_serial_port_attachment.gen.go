// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
)

// The class instance for the [VZFileHandleSerialPortAttachment] class.
var (
	_VZFileHandleSerialPortAttachmentClass     VZFileHandleSerialPortAttachmentClass
	_VZFileHandleSerialPortAttachmentClassOnce sync.Once
)

func getVZFileHandleSerialPortAttachmentClass() VZFileHandleSerialPortAttachmentClass {
	_VZFileHandleSerialPortAttachmentClassOnce.Do(func() {
		_VZFileHandleSerialPortAttachmentClass = VZFileHandleSerialPortAttachmentClass{class: objc.GetClass("VZFileHandleSerialPortAttachment")}
	})
	return _VZFileHandleSerialPortAttachmentClass
}

// GetVZFileHandleSerialPortAttachmentClass returns the class object for VZFileHandleSerialPortAttachment.
func GetVZFileHandleSerialPortAttachmentClass() VZFileHandleSerialPortAttachmentClass {
	return getVZFileHandleSerialPortAttachmentClass()
}

type VZFileHandleSerialPortAttachmentClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZFileHandleSerialPortAttachmentClass) Alloc() VZFileHandleSerialPortAttachment {
	rv := objc.Send[VZFileHandleSerialPortAttachment](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}







// An attachment point that allows bidirectional communication using file
// handles.
//
// # Overview
// 
// Use a [VZFileHandleSerialPortAttachment] object to configure a serial port
// using separate file handles for reading and writing data. In your virtual
// machine, use the file handles in this object in the following way:
// 
// - To send data to the guest operating system, write data to the file handle
// in the [VZFileHandleSerialPortAttachment.FileHandleForReading] property. - To receive data from the guest
// operating system, read data from the file handle in the
// [VZFileHandleSerialPortAttachment.FileHandleForWriting] property.
//
// # Creating the attachment point
//
//   - [VZFileHandleSerialPortAttachment.InitWithFileHandleForReadingFileHandleForWriting]: Creates a serial port attachment object from the specified file handles.
//
// # Getting the file handles
//
//   - [VZFileHandleSerialPortAttachment.FileHandleForReading]: The file handle that the guest operating system uses to read data.
//   - [VZFileHandleSerialPortAttachment.FileHandleForWriting]: The file handle that the guest operating system uses to write data.
//
// See: https://developer.apple.com/documentation/Virtualization/VZFileHandleSerialPortAttachment
type VZFileHandleSerialPortAttachment struct {
	VZSerialPortAttachment
}

// VZFileHandleSerialPortAttachmentFromID constructs a [VZFileHandleSerialPortAttachment] from an objc.ID.
//
// An attachment point that allows bidirectional communication using file
// handles.
func VZFileHandleSerialPortAttachmentFromID(id objc.ID) VZFileHandleSerialPortAttachment {
	return VZFileHandleSerialPortAttachment{VZSerialPortAttachment: VZSerialPortAttachmentFromID(id)}
}
// NOTE: VZFileHandleSerialPortAttachment adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [VZFileHandleSerialPortAttachment] class.
//
// # Creating the attachment point
//
//   - [IVZFileHandleSerialPortAttachment.InitWithFileHandleForReadingFileHandleForWriting]: Creates a serial port attachment object from the specified file handles.
//
// # Getting the file handles
//
//   - [IVZFileHandleSerialPortAttachment.FileHandleForReading]: The file handle that the guest operating system uses to read data.
//   - [IVZFileHandleSerialPortAttachment.FileHandleForWriting]: The file handle that the guest operating system uses to write data.
//
// See: https://developer.apple.com/documentation/Virtualization/VZFileHandleSerialPortAttachment
type IVZFileHandleSerialPortAttachment interface {
	IVZSerialPortAttachment

	// Topic: Creating the attachment point

	// Creates a serial port attachment object from the specified file handles.
	InitWithFileHandleForReadingFileHandleForWriting(fileHandleForReading foundation.NSFileHandle, fileHandleForWriting foundation.NSFileHandle) VZFileHandleSerialPortAttachment

	// Topic: Getting the file handles

	// The file handle that the guest operating system uses to read data.
	FileHandleForReading() foundation.NSFileHandle
	// The file handle that the guest operating system uses to write data.
	FileHandleForWriting() foundation.NSFileHandle
}





// Init initializes the instance.
func (f VZFileHandleSerialPortAttachment) Init() VZFileHandleSerialPortAttachment {
	rv := objc.Send[VZFileHandleSerialPortAttachment](f.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f VZFileHandleSerialPortAttachment) Autorelease() VZFileHandleSerialPortAttachment {
	rv := objc.Send[VZFileHandleSerialPortAttachment](f.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZFileHandleSerialPortAttachment creates a new VZFileHandleSerialPortAttachment instance.
func NewVZFileHandleSerialPortAttachment() VZFileHandleSerialPortAttachment {
	class := getVZFileHandleSerialPortAttachmentClass()
	rv := objc.Send[VZFileHandleSerialPortAttachment](objc.ID(class.class), objc.Sel("new"))
	return rv
}






// Creates a serial port attachment object from the specified file handles.
//
// fileHandleForReading: The file handle that the guest operating system uses to read data. Your
// virtual machine writes data to this file handle. If the file descriptor for
// the file handle is invalid, this method raises an exception.
//
// fileHandleForWriting: The file handle to which the guest operating system writes data. Your
// virtual machine reads data from this file handle. If its file descriptor
// for the file handle is invalid, this method raises an exception.
//
// # Return Value
// 
// A serial port attachment object.
//
// See: https://developer.apple.com/documentation/Virtualization/VZFileHandleSerialPortAttachment/init(fileHandleForReading:fileHandleForWriting:)
func NewFileHandleSerialPortAttachmentWithFileHandleForReadingFileHandleForWriting(fileHandleForReading foundation.NSFileHandle, fileHandleForWriting foundation.NSFileHandle) VZFileHandleSerialPortAttachment {
	instance := getVZFileHandleSerialPortAttachmentClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFileHandleForReading:fileHandleForWriting:"), fileHandleForReading, fileHandleForWriting)
	return VZFileHandleSerialPortAttachmentFromID(rv)
}







// Creates a serial port attachment object from the specified file handles.
//
// fileHandleForReading: The file handle that the guest operating system uses to read data. Your
// virtual machine writes data to this file handle. If the file descriptor for
// the file handle is invalid, this method raises an exception.
//
// fileHandleForWriting: The file handle to which the guest operating system writes data. Your
// virtual machine reads data from this file handle. If its file descriptor
// for the file handle is invalid, this method raises an exception.
//
// # Return Value
// 
// A serial port attachment object.
//
// See: https://developer.apple.com/documentation/Virtualization/VZFileHandleSerialPortAttachment/init(fileHandleForReading:fileHandleForWriting:)
func (f VZFileHandleSerialPortAttachment) InitWithFileHandleForReadingFileHandleForWriting(fileHandleForReading foundation.NSFileHandle, fileHandleForWriting foundation.NSFileHandle) VZFileHandleSerialPortAttachment {
	rv := objc.Send[VZFileHandleSerialPortAttachment](f.ID, objc.Sel("initWithFileHandleForReading:fileHandleForWriting:"), fileHandleForReading, fileHandleForWriting)
	return rv
}












// The file handle that the guest operating system uses to read data.
//
// # Discussion
// 
// When you want to send data to the guest operating system, write data to the
// file handle in this property.
//
// See: https://developer.apple.com/documentation/Virtualization/VZFileHandleSerialPortAttachment/fileHandleForReading
func (f VZFileHandleSerialPortAttachment) FileHandleForReading() foundation.NSFileHandle {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("fileHandleForReading"))
	return foundation.NSFileHandleFromID(objc.ID(rv))
}



// The file handle that the guest operating system uses to write data.
//
// # Discussion
// 
// When you want to receive data from the guest operating system, read data
// from the file handle in this property.
//
// See: https://developer.apple.com/documentation/Virtualization/VZFileHandleSerialPortAttachment/fileHandleForWriting
func (f VZFileHandleSerialPortAttachment) FileHandleForWriting() foundation.NSFileHandle {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("fileHandleForWriting"))
	return foundation.NSFileHandleFromID(objc.ID(rv))
}























