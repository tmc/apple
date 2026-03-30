// Code generated from Apple documentation for diskimages2. DO NOT EDIT.

package diskimages2

import (
	"errors"
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
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

// Class returns the underlying Objective-C class pointer.
func (bc BlankDiskImageCreatorClass) Class() objc.Class {
	return bc.class
}

// Alloc allocates memory for a new instance of the class.
func (bc BlankDiskImageCreatorClass) Alloc() BlankDiskImageCreator {
	rv := objc.Send[BlankDiskImageCreator](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [BlankDiskImageCreator.CreateImageWithNumBlocksError]
//   - [BlankDiskImageCreator.InitWithURLError]
//
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
//
// See: https://developer.apple.com/documentation/DiskImages2/BlankDiskImageCreator
type IBlankDiskImageCreator interface {
	IBaseDiskImageCreator

	// Topic: Methods

	CreateImageWithNumBlocksError(numBlocks uint64) (bool, error)
	InitWithURLError(url foundation.INSURL) (BlankDiskImageCreator, error)
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

// See: https://developer.apple.com/documentation/DiskImages2/BaseDiskImageCreator/initWithURL:defaultFormat:error:
func NewBlankDiskImageCreatorWithURLDefaultFormatError(url foundation.INSURL, format int64) (BlankDiskImageCreator, error) {
	var errorPtr objc.ID
	instance := getBlankDiskImageCreatorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithURL:defaultFormat:error:"), url, format, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return BlankDiskImageCreator{}, foundation.NSErrorFrom(errorPtr)
	}
	return BlankDiskImageCreatorFromID(rv), nil
}

// See: https://developer.apple.com/documentation/DiskImages2/BlankDiskImageCreator/initWithURL:error:
func NewBlankDiskImageCreatorWithURLError(url foundation.INSURL) (BlankDiskImageCreator, error) {
	var errorPtr objc.ID
	instance := getBlankDiskImageCreatorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithURL:error:"), url, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return BlankDiskImageCreator{}, foundation.NSErrorFrom(errorPtr)
	}
	return BlankDiskImageCreatorFromID(rv), nil
}

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
