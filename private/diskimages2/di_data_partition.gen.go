// Code generated from Apple documentation for diskimages2. DO NOT EDIT.

package diskimages2

import (
	"errors"
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [DIDataPartition] class.
var (
	_DIDataPartitionClass     DIDataPartitionClass
	_DIDataPartitionClassOnce sync.Once
)

func getDIDataPartitionClass() DIDataPartitionClass {
	_DIDataPartitionClassOnce.Do(func() {
		_DIDataPartitionClass = DIDataPartitionClass{class: objc.GetClass("DIDataPartition")}
	})
	return _DIDataPartitionClass
}

// GetDIDataPartitionClass returns the class object for DIDataPartition.
func GetDIDataPartitionClass() DIDataPartitionClass {
	return getDIDataPartitionClass()
}

type DIDataPartitionClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (dc DIDataPartitionClass) Class() objc.Class {
	return dc.class
}

// Alloc allocates memory for a new instance of the class.
func (dc DIDataPartitionClass) Alloc() DIDataPartition {
	rv := objc.Send[DIDataPartition](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [DIDataPartition.GPTTypeID]
//   - [DIDataPartition.BlockSize]
//   - [DIDataPartition.SetBlockSize]
//   - [DIDataPartition.FindPartitionWithDeviceBSDNameError]
//   - [DIDataPartition.FindVolumeBSDNameWithError]
//   - [DIDataPartition.FormatWithVolumeNameError]
//   - [DIDataPartition.GetFileSystemMinimalBlocks]
//   - [DIDataPartition.IoMedia]
//   - [DIDataPartition.SetIoMedia]
//   - [DIDataPartition.IoMediaUUID]
//   - [DIDataPartition.SetIoMediaUUID]
//   - [DIDataPartition.MediaContentUUID]
//   - [DIDataPartition.NewMountVolumeWithDiskArbError]
//   - [DIDataPartition.NumBlocks]
//   - [DIDataPartition.SetNumBlocks]
//   - [DIDataPartition.ResizeFileSystemToMinimumWithError]
//   - [DIDataPartition.ResizeFileSystemWithNumBlocksError]
//   - [DIDataPartition.VolumeBSDName]
//   - [DIDataPartition.SetVolumeBSDName]
//
// See: https://developer.apple.com/documentation/DiskImages2/DIDataPartition
type DIDataPartition struct {
	objectivec.Object
}

// DIDataPartitionFromID constructs a [DIDataPartition] from an objc.ID.
func DIDataPartitionFromID(id objc.ID) DIDataPartition {
	return DIDataPartition{objectivec.Object{ID: id}}
}

// Ensure DIDataPartition implements IDIDataPartition.
var _ IDIDataPartition = DIDataPartition{}

// An interface definition for the [DIDataPartition] class.
//
// # Methods
//
//   - [IDIDataPartition.GPTTypeID]
//   - [IDIDataPartition.BlockSize]
//   - [IDIDataPartition.SetBlockSize]
//   - [IDIDataPartition.FindPartitionWithDeviceBSDNameError]
//   - [IDIDataPartition.FindVolumeBSDNameWithError]
//   - [IDIDataPartition.FormatWithVolumeNameError]
//   - [IDIDataPartition.GetFileSystemMinimalBlocks]
//   - [IDIDataPartition.IoMedia]
//   - [IDIDataPartition.SetIoMedia]
//   - [IDIDataPartition.IoMediaUUID]
//   - [IDIDataPartition.SetIoMediaUUID]
//   - [IDIDataPartition.MediaContentUUID]
//   - [IDIDataPartition.NewMountVolumeWithDiskArbError]
//   - [IDIDataPartition.NumBlocks]
//   - [IDIDataPartition.SetNumBlocks]
//   - [IDIDataPartition.ResizeFileSystemToMinimumWithError]
//   - [IDIDataPartition.ResizeFileSystemWithNumBlocksError]
//   - [IDIDataPartition.VolumeBSDName]
//   - [IDIDataPartition.SetVolumeBSDName]
//
// See: https://developer.apple.com/documentation/DiskImages2/DIDataPartition
type IDIDataPartition interface {
	objectivec.IObject

	// Topic: Methods

	GPTTypeID() uint64
	BlockSize() uint64
	SetBlockSize(value uint64)
	FindPartitionWithDeviceBSDNameError(bSDName objectivec.IObject) (bool, error)
	FindVolumeBSDNameWithError() (bool, error)
	FormatWithVolumeNameError(name objectivec.IObject) (bool, error)
	GetFileSystemMinimalBlocks(blocks []objectivec.IObject) uint64
	IoMedia() IDIIOMedia
	SetIoMedia(value IDIIOMedia)
	IoMediaUUID() string
	SetIoMediaUUID(value string)
	MediaContentUUID() string
	NewMountVolumeWithDiskArbError(arb objectivec.IObject) (objectivec.IObject, error)
	NumBlocks() uint64
	SetNumBlocks(value uint64)
	ResizeFileSystemToMinimumWithError() (bool, error)
	ResizeFileSystemWithNumBlocksError(blocks uint64) (bool, error)
	VolumeBSDName() string
	SetVolumeBSDName(value string)
}

// Init initializes the instance.
func (d DIDataPartition) Init() DIDataPartition {
	rv := objc.Send[DIDataPartition](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d DIDataPartition) Autorelease() DIDataPartition {
	rv := objc.Send[DIDataPartition](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewDIDataPartition creates a new DIDataPartition instance.
func NewDIDataPartition() DIDataPartition {
	class := getDIDataPartitionClass()
	rv := objc.Send[DIDataPartition](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/DiskImages2/DIDataPartition/findPartitionWithDeviceBSDName:error:
func (d DIDataPartition) FindPartitionWithDeviceBSDNameError(bSDName objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](d.ID, objc.Sel("findPartitionWithDeviceBSDName:error:"), bSDName, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("findPartitionWithDeviceBSDName:error: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/DiskImages2/DIDataPartition/findVolumeBSDNameWithError:
func (d DIDataPartition) FindVolumeBSDNameWithError() (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](d.ID, objc.Sel("findVolumeBSDNameWithError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("findVolumeBSDNameWithError: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/DiskImages2/DIDataPartition/formatWithVolumeName:error:
func (d DIDataPartition) FormatWithVolumeNameError(name objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](d.ID, objc.Sel("formatWithVolumeName:error:"), name, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("formatWithVolumeName:error: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/DiskImages2/DIDataPartition/getFileSystemMinimalBlocks:
func (d DIDataPartition) GetFileSystemMinimalBlocks(blocks []objectivec.IObject) uint64 {
	rv := objc.Send[uint64](d.ID, objc.Sel("getFileSystemMinimalBlocks:"), objectivec.IObjectSliceToNSArray(blocks))
	return rv
}

// See: https://developer.apple.com/documentation/DiskImages2/DIDataPartition/newMountVolumeWithDiskArb:error:
func (d DIDataPartition) NewMountVolumeWithDiskArbError(arb objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](d.ID, objc.Sel("newMountVolumeWithDiskArb:error:"), arb, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// See: https://developer.apple.com/documentation/DiskImages2/DIDataPartition/resizeFileSystemToMinimumWithError:
func (d DIDataPartition) ResizeFileSystemToMinimumWithError() (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](d.ID, objc.Sel("resizeFileSystemToMinimumWithError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("resizeFileSystemToMinimumWithError: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/DiskImages2/DIDataPartition/resizeFileSystemWithNumBlocks:error:
func (d DIDataPartition) ResizeFileSystemWithNumBlocksError(blocks uint64) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](d.ID, objc.Sel("resizeFileSystemWithNumBlocks:error:"), blocks, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("resizeFileSystemWithNumBlocks:error: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/DiskImages2/DIDataPartition/findPartitionSchemeWithIOMedia:error:
func (_DIDataPartitionClass DIDataPartitionClass) FindPartitionSchemeWithIOMediaError(iOMedia objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_DIDataPartitionClass.class), objc.Sel("findPartitionSchemeWithIOMedia:error:"), iOMedia, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// See: https://developer.apple.com/documentation/DiskImages2/DIDataPartition/newMountURLWithError:
func (_DIDataPartitionClass DIDataPartitionClass) NewMountURLWithError() (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_DIDataPartitionClass.class), objc.Sel("newMountURLWithError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// See: https://developer.apple.com/documentation/DiskImages2/DIDataPartition/GPTTypeID
func (d DIDataPartition) GPTTypeID() uint64 {
	rv := objc.Send[uint64](d.ID, objc.Sel("GPTTypeID"))
	return rv
}

// See: https://developer.apple.com/documentation/DiskImages2/DIDataPartition/blockSize
func (d DIDataPartition) BlockSize() uint64 {
	rv := objc.Send[uint64](d.ID, objc.Sel("blockSize"))
	return rv
}
func (d DIDataPartition) SetBlockSize(value uint64) {
	objc.Send[struct{}](d.ID, objc.Sel("setBlockSize:"), value)
}

// See: https://developer.apple.com/documentation/DiskImages2/DIDataPartition/ioMedia
func (d DIDataPartition) IoMedia() IDIIOMedia {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("ioMedia"))
	return DIIOMediaFromID(objc.ID(rv))
}
func (d DIDataPartition) SetIoMedia(value IDIIOMedia) {
	objc.Send[struct{}](d.ID, objc.Sel("setIoMedia:"), value)
}

// See: https://developer.apple.com/documentation/DiskImages2/DIDataPartition/ioMediaUUID
func (d DIDataPartition) IoMediaUUID() string {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("ioMediaUUID"))
	return foundation.NSStringFromID(rv).String()
}
func (d DIDataPartition) SetIoMediaUUID(value string) {
	objc.Send[struct{}](d.ID, objc.Sel("setIoMediaUUID:"), objc.String(value))
}

// See: https://developer.apple.com/documentation/DiskImages2/DIDataPartition/mediaContentUUID
func (d DIDataPartition) MediaContentUUID() string {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("mediaContentUUID"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/DiskImages2/DIDataPartition/numBlocks
func (d DIDataPartition) NumBlocks() uint64 {
	rv := objc.Send[uint64](d.ID, objc.Sel("numBlocks"))
	return rv
}
func (d DIDataPartition) SetNumBlocks(value uint64) {
	objc.Send[struct{}](d.ID, objc.Sel("setNumBlocks:"), value)
}

// See: https://developer.apple.com/documentation/DiskImages2/DIDataPartition/volumeBSDName
func (d DIDataPartition) VolumeBSDName() string {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("volumeBSDName"))
	return foundation.NSStringFromID(rv).String()
}
func (d DIDataPartition) SetVolumeBSDName(value string) {
	objc.Send[struct{}](d.ID, objc.Sel("setVolumeBSDName:"), objc.String(value))
}
