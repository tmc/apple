// Code generated from Apple documentation for diskimages2. DO NOT EDIT.

package diskimages2

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [FileLocalXPC] class.
var (
	_FileLocalXPCClass     FileLocalXPCClass
	_FileLocalXPCClassOnce sync.Once
)

func getFileLocalXPCClass() FileLocalXPCClass {
	_FileLocalXPCClassOnce.Do(func() {
		_FileLocalXPCClass = FileLocalXPCClass{class: objc.GetClass("FileLocalXPC")}
	})
	return _FileLocalXPCClass
}

// GetFileLocalXPCClass returns the class object for FileLocalXPC.
func GetFileLocalXPCClass() FileLocalXPCClass {
	return getFileLocalXPCClass()
}

type FileLocalXPCClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (fc FileLocalXPCClass) Class() objc.Class {
	return fc.class
}

// Alloc allocates memory for a new instance of the class.
func (fc FileLocalXPCClass) Alloc() FileLocalXPC {
	rv := objc.Send[FileLocalXPC](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [FileLocalXPC.InitWithBackend]
//   - [FileLocalXPC.InitWithFileDescriptorWritableLocked]
//   - [FileLocalXPC.InitWithURLFileOpenFlags]
// See: https://developer.apple.com/documentation/DiskImages2/FileLocalXPC
type FileLocalXPC struct {
	BackendXPC
}

// FileLocalXPCFromID constructs a [FileLocalXPC] from an objc.ID.
func FileLocalXPCFromID(id objc.ID) FileLocalXPC {
	return FileLocalXPC{BackendXPC: BackendXPCFromID(id)}
}
// Ensure FileLocalXPC implements IFileLocalXPC.
var _ IFileLocalXPC = FileLocalXPC{}

// An interface definition for the [FileLocalXPC] class.
//
// # Methods
//
//   - [IFileLocalXPC.InitWithBackend]
//   - [IFileLocalXPC.InitWithFileDescriptorWritableLocked]
//   - [IFileLocalXPC.InitWithURLFileOpenFlags]
//
// See: https://developer.apple.com/documentation/DiskImages2/FileLocalXPC
type IFileLocalXPC interface {
	IBackendXPC

	// Topic: Methods

	InitWithBackend(backend unsafe.Pointer) FileLocalXPC
	InitWithFileDescriptorWritableLocked(descriptor int, writable bool, locked bool) FileLocalXPC
	InitWithURLFileOpenFlags(url foundation.INSURL, flags int) FileLocalXPC
}

// Init initializes the instance.
func (f FileLocalXPC) Init() FileLocalXPC {
	rv := objc.Send[FileLocalXPC](f.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f FileLocalXPC) Autorelease() FileLocalXPC {
	rv := objc.Send[FileLocalXPC](f.ID, objc.Sel("autorelease"))
	return rv
}

// NewFileLocalXPC creates a new FileLocalXPC instance.
func NewFileLocalXPC() FileLocalXPC {
	class := getFileLocalXPCClass()
	rv := objc.Send[FileLocalXPC](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/DiskImages2/FileLocalXPC/initWithBackend:
func NewFileLocalXPCWithBackend(backend unsafe.Pointer) FileLocalXPC {
	instance := getFileLocalXPCClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithBackend:"), backend)
	return FileLocalXPCFromID(rv)
}

//
// See: https://developer.apple.com/documentation/DiskImages2/FileLocalXPC/initWithCoder:
func NewFileLocalXPCWithCoder(coder objectivec.IObject) FileLocalXPC {
	instance := getFileLocalXPCClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return FileLocalXPCFromID(rv)
}

//
// See: https://developer.apple.com/documentation/DiskImages2/FileLocalXPC/initWithFileDescriptor:writable:locked:
func NewFileLocalXPCWithFileDescriptorWritableLocked(descriptor int, writable bool, locked bool) FileLocalXPC {
	instance := getFileLocalXPCClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFileDescriptor:writable:locked:"), descriptor, writable, locked)
	return FileLocalXPCFromID(rv)
}

//
// See: https://developer.apple.com/documentation/DiskImages2/FileLocalXPC/initWithURL:fileOpenFlags:
func NewFileLocalXPCWithURLFileOpenFlags(url foundation.INSURL, flags int) FileLocalXPC {
	instance := getFileLocalXPCClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithURL:fileOpenFlags:"), url, flags)
	return FileLocalXPCFromID(rv)
}

//
// See: https://developer.apple.com/documentation/DiskImages2/FileLocalXPC/initWithBackend:
func (f FileLocalXPC) InitWithBackend(backend unsafe.Pointer) FileLocalXPC {
	rv := objc.Send[FileLocalXPC](f.ID, objc.Sel("initWithBackend:"), backend)
	return rv
}
//
// See: https://developer.apple.com/documentation/DiskImages2/FileLocalXPC/initWithFileDescriptor:writable:locked:
func (f FileLocalXPC) InitWithFileDescriptorWritableLocked(descriptor int, writable bool, locked bool) FileLocalXPC {
	rv := objc.Send[FileLocalXPC](f.ID, objc.Sel("initWithFileDescriptor:writable:locked:"), descriptor, writable, locked)
	return rv
}
//
// See: https://developer.apple.com/documentation/DiskImages2/FileLocalXPC/initWithURL:fileOpenFlags:
func (f FileLocalXPC) InitWithURLFileOpenFlags(url foundation.INSURL, flags int) FileLocalXPC {
	rv := objc.Send[FileLocalXPC](f.ID, objc.Sel("initWithURL:fileOpenFlags:"), url, flags)
	return rv
}

