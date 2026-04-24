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

// The class instance for the [MKDIDevice] class.
var (
	_MKDIDeviceClass     MKDIDeviceClass
	_MKDIDeviceClassOnce sync.Once
)

func getMKDIDeviceClass() MKDIDeviceClass {
	_MKDIDeviceClassOnce.Do(func() {
		_MKDIDeviceClass = MKDIDeviceClass{class: objc.GetClass("MKDIDevice")}
	})
	return _MKDIDeviceClass
}

// GetMKDIDeviceClass returns the class object for MKDIDevice.
func GetMKDIDeviceClass() MKDIDeviceClass {
	return getMKDIDeviceClass()
}

type MKDIDeviceClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MKDIDeviceClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MKDIDeviceClass) Alloc() MKDIDevice {
	rv := objc.Send[MKDIDevice](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MKDIDevice.BlockSize]
//   - [MKDIDevice.MediaRef]
//   - [MKDIDevice.PartitionDiskWithGPTTypeIDError]
//   - [MKDIDevice.ResizeDataPartitionWithPartitionUUIDPartitionNumBlocksError]
//   - [MKDIDevice.UpdatePartitionMapWithError]
//   - [MKDIDevice.InitWithBSDNameNumBlocksBlockSizeError]
//
// See: https://developer.apple.com/documentation/DiskImages2/MKDIDevice
type MKDIDevice struct {
	objectivec.Object
}

// MKDIDeviceFromID constructs a [MKDIDevice] from an objc.ID.
func MKDIDeviceFromID(id objc.ID) MKDIDevice {
	return MKDIDevice{objectivec.Object{ID: id}}
}

// Ensure MKDIDevice implements IMKDIDevice.
var _ IMKDIDevice = MKDIDevice{}

// An interface definition for the [MKDIDevice] class.
//
// # Methods
//
//   - [IMKDIDevice.BlockSize]
//   - [IMKDIDevice.MediaRef]
//   - [IMKDIDevice.PartitionDiskWithGPTTypeIDError]
//   - [IMKDIDevice.ResizeDataPartitionWithPartitionUUIDPartitionNumBlocksError]
//   - [IMKDIDevice.UpdatePartitionMapWithError]
//   - [IMKDIDevice.InitWithBSDNameNumBlocksBlockSizeError]
//
// See: https://developer.apple.com/documentation/DiskImages2/MKDIDevice
type IMKDIDevice interface {
	objectivec.IObject

	// Topic: Methods

	BlockSize() int
	MediaRef() MKMediaRef
	PartitionDiskWithGPTTypeIDError(id uint64) (bool, error)
	ResizeDataPartitionWithPartitionUUIDPartitionNumBlocksError(uuid objectivec.IObject, blocks uint64) (bool, error)
	UpdatePartitionMapWithError() (bool, error)
	InitWithBSDNameNumBlocksBlockSizeError(bSDName objectivec.IObject, blocks uint64, size int) (MKDIDevice, error)
}

// Init initializes the instance.
func (m MKDIDevice) Init() MKDIDevice {
	rv := objc.Send[MKDIDevice](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MKDIDevice) Autorelease() MKDIDevice {
	rv := objc.Send[MKDIDevice](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMKDIDevice creates a new MKDIDevice instance.
func NewMKDIDevice() MKDIDevice {
	class := getMKDIDeviceClass()
	rv := objc.Send[MKDIDevice](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/DiskImages2/MKDIDevice/initWithBSDName:numBlocks:blockSize:error:
func NewMKDIDeviceWithBSDNameNumBlocksBlockSizeError(bSDName objectivec.IObject, blocks uint64, size int) (MKDIDevice, error) {
	var errorPtr objc.ID
	instance := getMKDIDeviceClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithBSDName:numBlocks:blockSize:error:"), bSDName, blocks, size, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MKDIDevice{}, foundation.NSErrorFrom(errorPtr)
	}
	return MKDIDeviceFromID(rv), nil
}

// See: https://developer.apple.com/documentation/DiskImages2/MKDIDevice/partitionDiskWithGPTTypeID:error:
func (m MKDIDevice) PartitionDiskWithGPTTypeIDError(id uint64) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](m.ID, objc.Sel("partitionDiskWithGPTTypeID:error:"), id, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("partitionDiskWithGPTTypeID:error: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/DiskImages2/MKDIDevice/resizeDataPartitionWithPartitionUUID:partitionNumBlocks:error:
func (m MKDIDevice) ResizeDataPartitionWithPartitionUUIDPartitionNumBlocksError(uuid objectivec.IObject, blocks uint64) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](m.ID, objc.Sel("resizeDataPartitionWithPartitionUUID:partitionNumBlocks:error:"), uuid, blocks, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("resizeDataPartitionWithPartitionUUID:partitionNumBlocks:error: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/DiskImages2/MKDIDevice/updatePartitionMapWithError:
func (m MKDIDevice) UpdatePartitionMapWithError() (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](m.ID, objc.Sel("updatePartitionMapWithError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("updatePartitionMapWithError: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/DiskImages2/MKDIDevice/initWithBSDName:numBlocks:blockSize:error:
func (m MKDIDevice) InitWithBSDNameNumBlocksBlockSizeError(bSDName objectivec.IObject, blocks uint64, size int) (MKDIDevice, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("initWithBSDName:numBlocks:blockSize:error:"), bSDName, blocks, size, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MKDIDevice{}, foundation.NSErrorFrom(errorPtr)
	}
	return MKDIDeviceFromID(rv), nil

}

// See: https://developer.apple.com/documentation/DiskImages2/MKDIDevice/blockSize
func (m MKDIDevice) BlockSize() int {
	rv := objc.Send[int](m.ID, objc.Sel("blockSize"))
	return rv
}

// See: https://developer.apple.com/documentation/DiskImages2/MKDIDevice/mediaRef
func (m MKDIDevice) MediaRef() MKMediaRef {
	rv := objc.Send[MKMediaRef](m.ID, objc.Sel("mediaRef"))
	return MKMediaRef(rv)
}
