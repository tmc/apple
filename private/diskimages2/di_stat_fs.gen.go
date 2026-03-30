// Code generated from Apple documentation for diskimages2. DO NOT EDIT.

package diskimages2

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [DIStatFS] class.
var (
	_DIStatFSClass     DIStatFSClass
	_DIStatFSClassOnce sync.Once
)

func getDIStatFSClass() DIStatFSClass {
	_DIStatFSClassOnce.Do(func() {
		_DIStatFSClass = DIStatFSClass{class: objc.GetClass("DIStatFS")}
	})
	return _DIStatFSClass
}

// GetDIStatFSClass returns the class object for DIStatFS.
func GetDIStatFSClass() DIStatFSClass {
	return getDIStatFSClass()
}

type DIStatFSClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (dc DIStatFSClass) Class() objc.Class {
	return dc.class
}

// Alloc allocates memory for a new instance of the class.
func (dc DIStatFSClass) Alloc() DIStatFS {
	rv := objc.Send[DIStatFS](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [DIStatFS.BlockSize]
//   - [DIStatFS.EncodeWithCoder]
//   - [DIStatFS.LogWithHeader]
//   - [DIStatFS.MountedFrom]
//   - [DIStatFS.MountedOnURL]
//   - [DIStatFS.SupportsBarrier]
//   - [DIStatFS.InitWithCoder]
//   - [DIStatFS.InitWithFileDescriptorError]
//
// See: https://developer.apple.com/documentation/DiskImages2/DIStatFS
type DIStatFS struct {
	objectivec.Object
}

// DIStatFSFromID constructs a [DIStatFS] from an objc.ID.
func DIStatFSFromID(id objc.ID) DIStatFS {
	return DIStatFS{objectivec.Object{ID: id}}
}

// Ensure DIStatFS implements IDIStatFS.
var _ IDIStatFS = DIStatFS{}

// An interface definition for the [DIStatFS] class.
//
// # Methods
//
//   - [IDIStatFS.BlockSize]
//   - [IDIStatFS.EncodeWithCoder]
//   - [IDIStatFS.LogWithHeader]
//   - [IDIStatFS.MountedFrom]
//   - [IDIStatFS.MountedOnURL]
//   - [IDIStatFS.SupportsBarrier]
//   - [IDIStatFS.InitWithCoder]
//   - [IDIStatFS.InitWithFileDescriptorError]
//
// See: https://developer.apple.com/documentation/DiskImages2/DIStatFS
type IDIStatFS interface {
	objectivec.IObject

	// Topic: Methods

	BlockSize() uint64
	EncodeWithCoder(coder foundation.INSCoder)
	LogWithHeader(header objectivec.IObject)
	MountedFrom() string
	MountedOnURL() foundation.INSURL
	SupportsBarrier() bool
	InitWithCoder(coder foundation.INSCoder) DIStatFS
	InitWithFileDescriptorError(descriptor int) (DIStatFS, error)
}

// Init initializes the instance.
func (d DIStatFS) Init() DIStatFS {
	rv := objc.Send[DIStatFS](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d DIStatFS) Autorelease() DIStatFS {
	rv := objc.Send[DIStatFS](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewDIStatFS creates a new DIStatFS instance.
func NewDIStatFS() DIStatFS {
	class := getDIStatFSClass()
	rv := objc.Send[DIStatFS](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/DiskImages2/DIStatFS/initWithCoder:
func NewDIStatFSWithCoder(coder objectivec.IObject) DIStatFS {
	instance := getDIStatFSClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return DIStatFSFromID(rv)
}

// See: https://developer.apple.com/documentation/DiskImages2/DIStatFS/initWithFileDescriptor:error:
func NewDIStatFSWithFileDescriptorError(descriptor int) (DIStatFS, error) {
	var errorPtr objc.ID
	instance := getDIStatFSClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFileDescriptor:error:"), descriptor, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return DIStatFS{}, foundation.NSErrorFrom(errorPtr)
	}
	return DIStatFSFromID(rv), nil
}

// See: https://developer.apple.com/documentation/DiskImages2/DIStatFS/encodeWithCoder:
func (d DIStatFS) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](d.ID, objc.Sel("encodeWithCoder:"), coder)
}

// See: https://developer.apple.com/documentation/DiskImages2/DIStatFS/logWithHeader:
func (d DIStatFS) LogWithHeader(header objectivec.IObject) {
	objc.Send[objc.ID](d.ID, objc.Sel("logWithHeader:"), header)
}

// See: https://developer.apple.com/documentation/DiskImages2/DIStatFS/initWithCoder:
func (d DIStatFS) InitWithCoder(coder foundation.INSCoder) DIStatFS {
	rv := objc.Send[DIStatFS](d.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}

// See: https://developer.apple.com/documentation/DiskImages2/DIStatFS/initWithFileDescriptor:error:
func (d DIStatFS) InitWithFileDescriptorError(descriptor int) (DIStatFS, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](d.ID, objc.Sel("initWithFileDescriptor:error:"), descriptor, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return DIStatFS{}, foundation.NSErrorFrom(errorPtr)
	}
	return DIStatFSFromID(rv), nil

}

// See: https://developer.apple.com/documentation/DiskImages2/DIStatFS/supportsSecureCoding
func (_DIStatFSClass DIStatFSClass) SupportsSecureCoding() bool {
	rv := objc.Send[bool](objc.ID(_DIStatFSClass.class), objc.Sel("supportsSecureCoding"))
	return rv
}

// See: https://developer.apple.com/documentation/DiskImages2/DIStatFS/blockSize
func (d DIStatFS) BlockSize() uint64 {
	rv := objc.Send[uint64](d.ID, objc.Sel("blockSize"))
	return rv
}

// See: https://developer.apple.com/documentation/DiskImages2/DIStatFS/mountedFrom
func (d DIStatFS) MountedFrom() string {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("mountedFrom"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/DiskImages2/DIStatFS/mountedOnURL
func (d DIStatFS) MountedOnURL() foundation.INSURL {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("mountedOnURL"))
	return foundation.NSURLFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/DiskImages2/DIStatFS/supportsBarrier
func (d DIStatFS) SupportsBarrier() bool {
	rv := objc.Send[bool](d.ID, objc.Sel("supportsBarrier"))
	return rv
}
