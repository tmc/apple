// Code generated from Apple documentation for diskimages2. DO NOT EDIT.

package diskimages2

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"errors"
	"github.com/tmc/apple/foundation"
)

// The class instance for the [BlankDiskImageCreator] class.
var (
	_BlankDiskImageCreatorClass     BlankDiskImageCreatorClass
	_BlankDiskImageCreatorClassOnce sync.Once
)

func getBlankDiskImageCreatorClass() BlankDiskImageCreatorClass {
	_BlankDiskImageCreatorClassOnce.Do(func() {
		_BlankDiskImageCreatorClass = BlankDiskImageCreatorClass{class: objc.GetClass("BlankDiskImageCreator")}
	})
	return _BlankDiskImageCreatorClass
}

// GetBlankDiskImageCreatorClass returns the class object for BlankDiskImageCreator.
func GetBlankDiskImageCreatorClass() BlankDiskImageCreatorClass {
	return getBlankDiskImageCreatorClass()
}

type BlankDiskImageCreatorClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (bc BlankDiskImageCreatorClass) Alloc() BlankDiskImageCreator {
	rv := objc.Send[BlankDiskImageCreator](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [BlankDiskImageCreator.CreateImageWithNumBlocksError]
//   - [BlankDiskImageCreator.InitWithURLError]
//   - [BlankDiskImageCreator.FileSystem]
//   - [BlankDiskImageCreator.SetFileSystem]
// See: https://developer.apple.com/documentation/DiskImages2/BlankDiskImageCreator
type BlankDiskImageCreator struct {
	BaseDiskImageCreator
}

// BlankDiskImageCreatorFromID constructs a [BlankDiskImageCreator] from an objc.ID.
func BlankDiskImageCreatorFromID(id objc.ID) BlankDiskImageCreator {
	return BlankDiskImageCreator{BaseDiskImageCreator: BaseDiskImageCreatorFromID(id)}
}
// Ensure BlankDiskImageCreator implements IBlankDiskImageCreator.
var _ IBlankDiskImageCreator = BlankDiskImageCreator{}

// An interface definition for the [BlankDiskImageCreator] class.
//
// # Methods
//
//   - [IBlankDiskImageCreator.CreateImageWithNumBlocksError]
//   - [IBlankDiskImageCreator.InitWithURLError]
//   - [IBlankDiskImageCreator.FileSystem]
//   - [IBlankDiskImageCreator.SetFileSystem]
//
// See: https://developer.apple.com/documentation/DiskImages2/BlankDiskImageCreator
type IBlankDiskImageCreator interface {
	IBaseDiskImageCreator

	// Topic: Methods

	CreateImageWithNumBlocksError(numBlocks uint64) (bool, error)
	InitWithURLError(url foundation.INSURL) (BlankDiskImageCreator, error)
	FileSystem() uint64
	SetFileSystem(value uint64)
}

// Init initializes the instance.
func (b BlankDiskImageCreator) Init() BlankDiskImageCreator {
	rv := objc.Send[BlankDiskImageCreator](b.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b BlankDiskImageCreator) Autorelease() BlankDiskImageCreator {
	rv := objc.Send[BlankDiskImageCreator](b.ID, objc.Sel("autorelease"))
	return rv
}

// NewBlankDiskImageCreator creates a new BlankDiskImageCreator instance.
func NewBlankDiskImageCreator() BlankDiskImageCreator {
	class := getBlankDiskImageCreatorClass()
	rv := objc.Send[BlankDiskImageCreator](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/DiskImages2/BlankDiskImageCreator/initWithURL:error:
func NewBlankDiskImageCreatorWithURLError(url foundation.INSURL) (BlankDiskImageCreator, error) {
	var errorPtr objc.ID
	instance := getBlankDiskImageCreatorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithURL:error:"), url, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return BlankDiskImageCreatorFromID(rv), foundation.NSErrorFrom(errorPtr)
	}
	return BlankDiskImageCreatorFromID(rv), nil
}

//
// See: https://developer.apple.com/documentation/DiskImages2/BlankDiskImageCreator/createImageWithNumBlocks:error:
func (b BlankDiskImageCreator) CreateImageWithNumBlocksError(numBlocks uint64) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](b.ID, objc.Sel("createImageWithNumBlocks:error:"), numBlocks, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("createImageWithNumBlocks:error: returned NO with nil NSError")
	}
	return rv, nil

}
//
// See: https://developer.apple.com/documentation/DiskImages2/BlankDiskImageCreator/initWithURL:error:
func (b BlankDiskImageCreator) InitWithURLError(url foundation.INSURL) (BlankDiskImageCreator, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](b.ID, objc.Sel("initWithURL:error:"), url, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return BlankDiskImageCreator{}, foundation.NSErrorFrom(errorPtr)
	}
	return BlankDiskImageCreatorFromID(rv), nil

}

// See: https://developer.apple.com/documentation/DiskImages2/BlankDiskImageCreator/fileSystem
func (b BlankDiskImageCreator) FileSystem() uint64 {
	rv := objc.Send[uint64](b.ID, objc.Sel("fileSystem"))
	return rv
}
func (b BlankDiskImageCreator) SetFileSystem(value uint64) {
	objc.Send[struct{}](b.ID, objc.Sel("setFileSystem:"), value)
}

