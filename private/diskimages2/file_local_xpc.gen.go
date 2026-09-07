// Code generated from Apple documentation for diskimages2. DO NOT EDIT.

package diskimages2

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
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
	rv := objc.SendIfResponds[FileLocalXPC](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [FileLocalXPC.FileUrl]
//   - [FileLocalXPC.SandboxExtensionTokenWithError]
//   - [FileLocalXPC.SandboxToken]
//   - [FileLocalXPC.InitWithBackend]
//   - [FileLocalXPC.InitWithFileDescriptorWritableLocked]
//   - [FileLocalXPC.InitWithURLFileOpenFlags]
//   - [FileLocalXPC.DebugDescription]
//   - [FileLocalXPC.Description]
//   - [FileLocalXPC.Hash]
//   - [FileLocalXPC.Superclass]
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
//   - [IFileLocalXPC.FileUrl]
//   - [IFileLocalXPC.SandboxExtensionTokenWithError]
//   - [IFileLocalXPC.SandboxToken]
//   - [IFileLocalXPC.InitWithBackend]
//   - [IFileLocalXPC.InitWithFileDescriptorWritableLocked]
//   - [IFileLocalXPC.InitWithURLFileOpenFlags]
//   - [IFileLocalXPC.DebugDescription]
//   - [IFileLocalXPC.Description]
//   - [IFileLocalXPC.Hash]
//   - [IFileLocalXPC.Superclass]
type IFileLocalXPC interface {
	IBackendXPC

	// Topic: Methods

	FileUrl() foundation.NSURL
	SandboxExtensionTokenWithError() (objectivec.IObject, error)
	SandboxToken() string
	InitWithBackend(backend unsafe.Pointer) FileLocalXPC
	InitWithFileDescriptorWritableLocked(descriptor int32, writable bool, locked bool) FileLocalXPC
	InitWithURLFileOpenFlags(url foundation.NSURL, flags int32) FileLocalXPC
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objectivec.Class
}

// Init initializes the instance.
func (f FileLocalXPC) Init() FileLocalXPC {
	rv := objc.SendIfResponds[FileLocalXPC](f.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f FileLocalXPC) Autorelease() FileLocalXPC {
	rv := objc.SendIfResponds[FileLocalXPC](f.ID, objc.Sel("autorelease"))
	return rv
}

// NewFileLocalXPC creates a new FileLocalXPC instance.
func NewFileLocalXPC() FileLocalXPC {
	class := getFileLocalXPCClass()
	rv := objc.SendIfResponds[FileLocalXPC](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewFileLocalXPCWithBackend(backend unsafe.Pointer) FileLocalXPC {
	instance := getFileLocalXPCClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithBackend:"), backend)
	return FileLocalXPCFromID(rv)
}

func NewFileLocalXPCWithCoder(coder objectivec.IObject) FileLocalXPC {
	instance := getFileLocalXPCClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return FileLocalXPCFromID(rv)
}

func NewFileLocalXPCWithFileDescriptorWritableLocked(descriptor int32, writable bool, locked bool) FileLocalXPC {
	instance := getFileLocalXPCClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithFileDescriptor:writable:locked:"), descriptor, writable, locked)
	return FileLocalXPCFromID(rv)
}

func NewFileLocalXPCWithURLFileOpenFlags(url foundation.NSURL, flags int32) FileLocalXPC {
	instance := getFileLocalXPCClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithURL:fileOpenFlags:"), url, flags)
	return FileLocalXPCFromID(rv)
}

func (f FileLocalXPC) SandboxExtensionTokenWithError() (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](f.ID, objc.Sel("sandboxExtensionTokenWithError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (f FileLocalXPC) InitWithBackend(backend unsafe.Pointer) FileLocalXPC {
	rv := objc.SendIfResponds[FileLocalXPC](f.ID, objc.Sel("initWithBackend:"), backend)
	return rv
}
func (f FileLocalXPC) InitWithFileDescriptorWritableLocked(descriptor int32, writable bool, locked bool) FileLocalXPC {
	rv := objc.SendIfResponds[FileLocalXPC](f.ID, objc.Sel("initWithFileDescriptor:writable:locked:"), descriptor, writable, locked)
	return rv
}
func (f FileLocalXPC) InitWithURLFileOpenFlags(url foundation.NSURL, flags int32) FileLocalXPC {
	rv := objc.SendIfResponds[FileLocalXPC](f.ID, objc.Sel("initWithURL:fileOpenFlags:"), url, flags)
	return rv
}

func (f FileLocalXPC) DebugDescription() string {
	rv := objc.SendIfResponds[objc.ID](f.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
func (f FileLocalXPC) Description() string {
	rv := objc.SendIfResponds[objc.ID](f.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
func (f FileLocalXPC) FileUrl() foundation.NSURL {
	rv := objc.SendIfResponds[foundation.NSURL](f.ID, objc.Sel("fileUrl"))
	return foundation.NSURL(rv)
}
func (f FileLocalXPC) Hash() uint64 {
	rv := objc.SendIfResponds[uint64](f.ID, objc.Sel("hash"))
	return rv
}
func (f FileLocalXPC) SandboxToken() string {
	rv := objc.SendIfResponds[objc.ID](f.ID, objc.Sel("sandboxToken"))
	return foundation.NSStringFromID(rv).String()
}
func (f FileLocalXPC) Superclass() objectivec.Class {
	rv := objc.SendIfResponds[objectivec.Class](f.ID, objc.Sel("superclass"))
	return objectivec.Class(rv)
}
