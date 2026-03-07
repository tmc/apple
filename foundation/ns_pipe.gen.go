// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [Pipe] class.
var (
	_PipeClass     PipeClass
	_PipeClassOnce sync.Once
)

func getPipeClass() PipeClass {
	_PipeClassOnce.Do(func() {
		_PipeClass = PipeClass{class: objc.GetClass("NSPipe")}
	})
	return _PipeClass
}

// GetPipeClass returns the class object for NSPipe.
func GetPipeClass() PipeClass {
	return getPipeClass()
}

type PipeClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (pc PipeClass) Alloc() Pipe {
	rv := objc.Send[Pipe](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}







// A one-way communications channel between related processes.
//
// # Overview
// 
// [NSPipe] objects provide an object-oriented interface for accessing pipes.
// An [NSPipe] object represents both ends of a pipe and enables communication
// through the pipe. A pipe is a one-way communications channel between
// related processes; one process writes data, while the other process reads
// that data. The data that passes through the pipe is buffered; the size of
// the buffer is determined by the underlying operating system. [NSPipe] is an
// abstract class, the public interface of a class cluster.
//
// # Getting the File Handles for a Pipe
//
//   - [Pipe.FileHandleForReading]: The receiver’s read file handle.
//   - [Pipe.FileHandleForWriting]: The receiver’s write file handle.
//
// See: https://developer.apple.com/documentation/Foundation/Pipe
type Pipe struct {
	objectivec.Object
}

// PipeFromID constructs a [Pipe] from an objc.ID.
//
// A one-way communications channel between related processes.
func PipeFromID(id objc.ID) Pipe {
	return NSPipe{objectivec.Object{id}}
}

// NSPipeFromID is an alias for [PipeFromID] for cross-framework compatibility.
func NSPipeFromID(id objc.ID) Pipe { return PipeFromID(id) }
// NOTE: Pipe adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [Pipe] class.
//
// # Getting the File Handles for a Pipe
//
//   - [IPipe.FileHandleForReading]: The receiver’s read file handle.
//   - [IPipe.FileHandleForWriting]: The receiver’s write file handle.
//
// See: https://developer.apple.com/documentation/Foundation/Pipe
type IPipe interface {
	objectivec.IObject

	// Topic: Getting the File Handles for a Pipe

	// The receiver’s read file handle.
	FileHandleForReading() INSFileHandle
	// The receiver’s write file handle.
	FileHandleForWriting() INSFileHandle
}





// Init initializes the instance.
func (p Pipe) Init() Pipe {
	rv := objc.Send[Pipe](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p Pipe) Autorelease() Pipe {
	rv := objc.Send[Pipe](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewPipe creates a new Pipe instance.
func NewPipe() Pipe {
	class := getPipeClass()
	rv := objc.Send[Pipe](objc.ID(class.class), objc.Sel("new"))
	return rv
}














// Returns an [NSPipe] object.
//
// # Return Value
// 
// An initialized [NSPipe] object. Returns `nil` if the method encounters
// errors while attempting to create the pipe or the [NSFileHandle] objects
// that serve as endpoints of the pipe.
//
// See: https://developer.apple.com/documentation/Foundation/NSPipe/pipe
func (_PipeClass PipeClass) Pipe() Pipe {
	rv := objc.Send[objc.ID](objc.ID(_PipeClass.class), objc.Sel("pipe"))
	return NSPipeFromID(rv)
}








// The receiver’s read file handle.
//
// # Discussion
// 
// The descriptor represented by this object is deleted, and the object itself
// is automatically deallocated when the receiver is deallocated.
// 
// You use the returned file handle to read from the pipe using
// [NSFileHandle]’s read methods—[AvailableData], [ReadDataToEndOfFile],
// and [ReadDataOfLength].
// 
// You don’t need to send [CloseFile] to this object or explicitly release
// the object after you have finished using it.
//
// See: https://developer.apple.com/documentation/Foundation/Pipe/fileHandleForReading
func (p Pipe) FileHandleForReading() INSFileHandle {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("fileHandleForReading"))
	return NSFileHandleFromID(objc.ID(rv))
}



// The receiver’s write file handle.
//
// # Discussion
// 
// This object is automatically deallocated when the receiver is deallocated.
// 
// You use the returned file handle to write to the pipe using
// [NSFileHandle]’s [WriteData] method. When you are finished writing data
// to this object, send it a [CloseFile] message to delete the descriptor.
// Deleting the descriptor causes the reading process to receive an
// end-of-data signal (an empty [NSData] object).
//
// See: https://developer.apple.com/documentation/Foundation/Pipe/fileHandleForWriting
func (p Pipe) FileHandleForWriting() INSFileHandle {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("fileHandleForWriting"))
	return NSFileHandleFromID(objc.ID(rv))
}

























