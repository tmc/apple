// Code generated from Apple documentation for FSKit. DO NOT EDIT.

package fskit

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [FSFileName] class.
var (
	_FSFileNameClass     FSFileNameClass
	_FSFileNameClassOnce sync.Once
)

func getFSFileNameClass() FSFileNameClass {
	_FSFileNameClassOnce.Do(func() {
		_FSFileNameClass = FSFileNameClass{class: objc.GetClass("FSFileName")}
	})
	return _FSFileNameClass
}

// GetFSFileNameClass returns the class object for FSFileName.
func GetFSFileNameClass() FSFileNameClass {
	return getFSFileNameClass()
}

type FSFileNameClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (fc FSFileNameClass) Class() objc.Class {
	return fc.class
}

// Alloc allocates memory for a new instance of the class.
func (fc FSFileNameClass) Alloc() FSFileName {
	rv := objc.Send[FSFileName](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// The name of a file, expressed as a data buffer.
//
// # Overview
//
// [FSFileName] is the class that carries filenames from the kernel to
// [FSModule] instances, and carries names back to the kernel as part of
// directory enumeration.
//
// A filename is usually a valid UTF-8 sequence, but can be an arbitrary byte
// sequence that doesn’t conform to that format. As a result, the [FSFileName.Data]
// property always contains a value, but the [FSFileName.String] property may be empty.
// An [FSModule] can receive an [FSFileName] that isn’t valid UTF-8 in two
// cases:
//
// - A program passes erroneous data to a system call. The [FSModule] treats
// this situation as an error. - An [FSModule] lacks the character encoding
// used for a file name. This situation occurs because some file system
// formats consider a filename to be an arbitrary “bag of bytes,” and
// leave character encoding up to the operating system. Without encoding
// information, the [FSModule] can only pass back the names it finds on disk.
// In this case, the behavior of upper layers such as [FileManager] is
// unspecified. However, the [FSModule] must support looking up such names and
// using them as the source name of rename operations. The [FSModule] must
// also be able to support filenames that are derivatives of filenames
// returned from directory enumeration. Derivative filenames include Apple
// Double filenames (`"._Name"`), and editor backup filenames.
//
// # Creating a filename
//
//   - [FSFileName.InitWithData]: Creates a filename by copying a character sequence data object.
//   - [FSFileName.InitWithString]: Creates a filename by copying a character sequence from a string instance.
//
// # Accessing filename properties
//
//   - [FSFileName.Data]: The byte sequence of the filename, as a data object.
//   - [FSFileName.String]: The filename, represented as a Unicode string.
//   - [FSFileName.DebugDescription]: The filename, represented as a potentially lossy conversion to a string.
//
// See: https://developer.apple.com/documentation/FSKit/FSFileName
//
// [FileManager]: https://developer.apple.com/documentation/Foundation/FileManager
type FSFileName struct {
	objectivec.Object
}

// FSFileNameFromID constructs a [FSFileName] from an objc.ID.
//
// The name of a file, expressed as a data buffer.
func FSFileNameFromID(id objc.ID) FSFileName {
	return FSFileName{objectivec.Object{ID: id}}
}

// NOTE: FSFileName adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [FSFileName] class.
//
// # Creating a filename
//
//   - [IFSFileName.InitWithData]: Creates a filename by copying a character sequence data object.
//   - [IFSFileName.InitWithString]: Creates a filename by copying a character sequence from a string instance.
//
// # Accessing filename properties
//
//   - [IFSFileName.Data]: The byte sequence of the filename, as a data object.
//   - [IFSFileName.String]: The filename, represented as a Unicode string.
//   - [IFSFileName.DebugDescription]: The filename, represented as a potentially lossy conversion to a string.
//
// See: https://developer.apple.com/documentation/FSKit/FSFileName
type IFSFileName interface {
	objectivec.IObject

	// Topic: Creating a filename

	// Creates a filename by copying a character sequence data object.
	InitWithData(name foundation.NSData) FSFileName
	// Creates a filename by copying a character sequence from a string instance.
	InitWithString(name string) FSFileName

	// Topic: Accessing filename properties

	// The byte sequence of the filename, as a data object.
	Data() foundation.NSData
	// The filename, represented as a Unicode string.
	String() string
	// The filename, represented as a potentially lossy conversion to a string.
	DebugDescription() string

	// Initializes a file name by copying a character sequence from a byte array.
	InitWithBytesLength(bytes string, length uint) FSFileName
	// Initializes a filename from a null-terminated character sequence.
	InitWithCString(name string) FSFileName
	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (f FSFileName) Init() FSFileName {
	rv := objc.Send[FSFileName](f.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f FSFileName) Autorelease() FSFileName {
	rv := objc.Send[FSFileName](f.ID, objc.Sel("autorelease"))
	return rv
}

// NewFSFileName creates a new FSFileName instance.
func NewFSFileName() FSFileName {
	class := getFSFileNameClass()
	rv := objc.Send[FSFileName](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Initializes a file name by copying a character sequence from a byte array.
//
// # Discussion
//
// - bytes: A pointer to the character data to copy, up to a maximum of
// `length`. The sequence terminates if a [NUL] character exists prior to
// `length`. - length: The size of the `bytes` array.
//
// See: https://developer.apple.com/documentation/FSKit/FSFileName/initWithBytes:length:
func NewFileNameWithBytesLength(bytes string, length uint) FSFileName {
	instance := getFSFileNameClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithBytes:length:"), unsafe.Pointer(unsafe.StringData(bytes+"\x00")), length)
	return FSFileNameFromID(rv)
}

// Initializes a filename from a null-terminated character sequence.
//
// name: A pointer to a C string.
//
// # Discussion
//
// See: https://developer.apple.com/documentation/FSKit/FSFileName/initWithCString:
func NewFileNameWithCString(name string) FSFileName {
	instance := getFSFileNameClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCString:"), unsafe.Pointer(unsafe.StringData(name+"\x00")))
	return FSFileNameFromID(rv)
}

// Creates a filename by copying a character sequence data object.
//
// name: The data object containing the character sequence to use for the filename.
// The sequence terminates if a [NUL] character exists prior to
// `name.Length()`.
//
// # Discussion
//
// This initializer copies up to `name.Length()` characters of the sequence
// pointed to by `bytes`.
//
// See: https://developer.apple.com/documentation/FSKit/FSFileName/init(data:)
func NewFileNameWithData(name foundation.NSData) FSFileName {
	instance := getFSFileNameClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithData:"), name)
	return FSFileNameFromID(rv)
}

// Creates a filename by copying a character sequence from a string instance.
//
// name: The string containing the character sequence to use for the filename.
//
// # Discussion
//
// This initializer copies the UTF-8 representation of the characters in
// `string`. If `string` contains a [NUL] character, the sequence terminates.
//
// See: https://developer.apple.com/documentation/FSKit/FSFileName/init(string:)
func NewFileNameWithString(name string) FSFileName {
	instance := getFSFileNameClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithString:"), objc.String(name))
	return FSFileNameFromID(rv)
}

// Creates a filename by copying a character sequence data object.
//
// name: The data object containing the character sequence to use for the filename.
// The sequence terminates if a [NUL] character exists prior to
// `name.Length()`.
//
// # Discussion
//
// This initializer copies up to `name.Length()` characters of the sequence
// pointed to by `bytes`.
//
// See: https://developer.apple.com/documentation/FSKit/FSFileName/init(data:)
func (f FSFileName) InitWithData(name foundation.NSData) FSFileName {
	rv := objc.Send[FSFileName](f.ID, objc.Sel("initWithData:"), name)
	return rv
}

// Creates a filename by copying a character sequence from a string instance.
//
// name: The string containing the character sequence to use for the filename.
//
// # Discussion
//
// This initializer copies the UTF-8 representation of the characters in
// `string`. If `string` contains a [NUL] character, the sequence terminates.
//
// See: https://developer.apple.com/documentation/FSKit/FSFileName/init(string:)
func (f FSFileName) InitWithString(name string) FSFileName {
	rv := objc.Send[FSFileName](f.ID, objc.Sel("initWithString:"), objc.String(name))
	return rv
}

// Initializes a file name by copying a character sequence from a byte array.
//
// # Discussion
//
// - bytes: A pointer to the character data to copy, up to a maximum of
// `length`. The sequence terminates if a [NUL] character exists prior to
// `length`. - length: The size of the `bytes` array.
//
// See: https://developer.apple.com/documentation/FSKit/FSFileName/initWithBytes:length:
func (f FSFileName) InitWithBytesLength(bytes string, length uint) FSFileName {
	rv := objc.Send[FSFileName](f.ID, objc.Sel("initWithBytes:length:"), unsafe.Pointer(unsafe.StringData(bytes+"\x00")), length)
	return rv
}

// Initializes a filename from a null-terminated character sequence.
//
// name: A pointer to a C string.
//
// # Discussion
//
// See: https://developer.apple.com/documentation/FSKit/FSFileName/initWithCString:
func (f FSFileName) InitWithCString(name string) FSFileName {
	rv := objc.Send[FSFileName](f.ID, objc.Sel("initWithCString:"), unsafe.Pointer(unsafe.StringData(name+"\x00")))
	return rv
}
func (f FSFileName) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](f.ID, objc.Sel("encodeWithCoder:"), coder)
}

// Creates a filename by copying a character sequence from a byte array.
//
// # Discussion
//
// - bytes: A pointer to the character data to copy, up to a maximum of
// `length`. The sequence terminates if a [NUL] character exists prior to
// `length`. - length: The size of the `bytes` array.
//
// See: https://developer.apple.com/documentation/FSKit/FSFileName/nameWithBytes:length:
func (_FSFileNameClass FSFileNameClass) NameWithBytesLength(bytes string, length uint) FSFileName {
	rv := objc.Send[objc.ID](objc.ID(_FSFileNameClass.class), objc.Sel("nameWithBytes:length:"), unsafe.Pointer(unsafe.StringData(bytes+"\x00")), length)
	return FSFileNameFromID(rv)
}

// Creates a filename from a null-terminated character sequence.
//
// name: A pointer to a C string.
//
// See: https://developer.apple.com/documentation/FSKit/FSFileName/nameWithCString:
func (_FSFileNameClass FSFileNameClass) NameWithCString(name string) FSFileName {
	rv := objc.Send[objc.ID](objc.ID(_FSFileNameClass.class), objc.Sel("nameWithCString:"), unsafe.Pointer(unsafe.StringData(name+"\x00")))
	return FSFileNameFromID(rv)
}

// Creates a filename by copying a character sequence data object.
//
// name: The data object containing the character sequence to use for the filename.
// The sequence terminates if a [NUL] character exists prior to
// `name.Length()`.
//
// # Discussion
//
// This initializer copies up to `name.Length()` characters of the sequence
// pointed to by `bytes`.
//
// See: https://developer.apple.com/documentation/FSKit/FSFileName/nameWithData:
func (_FSFileNameClass FSFileNameClass) NameWithData(name foundation.NSData) FSFileName {
	rv := objc.Send[objc.ID](objc.ID(_FSFileNameClass.class), objc.Sel("nameWithData:"), name)
	return FSFileNameFromID(rv)
}

// Creates a filename by copying a character sequence from a string instance.
//
// name: The string containing the character sequence to use for the filename.
//
// # Discussion
//
// This initializer copies the UTF-8 representation of the characters in
// `string`. If `string` contains a [NUL] character, the sequence terminates.
//
// See: https://developer.apple.com/documentation/FSKit/FSFileName/nameWithString:
func (_FSFileNameClass FSFileNameClass) NameWithString(name string) FSFileName {
	rv := objc.Send[objc.ID](objc.ID(_FSFileNameClass.class), objc.Sel("nameWithString:"), objc.String(name))
	return FSFileNameFromID(rv)
}

// The byte sequence of the filename, as a data object.
//
// # Discussion
//
// This property always provides a value.
//
// See: https://developer.apple.com/documentation/FSKit/FSFileName/data
func (f FSFileName) Data() foundation.NSData {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("data"))
	return foundation.NSDataFromID(objc.ID(rv))
}

// The filename, represented as a Unicode string.
//
// # Discussion
//
// If the value of the filename’s [Data] is not a valid UTF-8 byte sequence,
// this property is empty.
//
// See: https://developer.apple.com/documentation/FSKit/FSFileName/string
func (f FSFileName) String() string {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("string"))
	return foundation.NSStringFromID(rv).String()
}

// The filename, represented as a potentially lossy conversion to a string.
//
// # Discussion
//
// The exact details of the string conversion may change in the future.
//
// See: https://developer.apple.com/documentation/FSKit/FSFileName/debugDescription
func (f FSFileName) DebugDescription() string {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
