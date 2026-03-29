// Code generated from Apple documentation for diskimages2. DO NOT EDIT.

package diskimages2

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"errors"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [DIAttachedDeviceInfo] class.
var (
	_DIAttachedDeviceInfoClass     DIAttachedDeviceInfoClass
	_DIAttachedDeviceInfoClassOnce sync.Once
)

func getDIAttachedDeviceInfoClass() DIAttachedDeviceInfoClass {
	_DIAttachedDeviceInfoClassOnce.Do(func() {
		_DIAttachedDeviceInfoClass = DIAttachedDeviceInfoClass{class: objc.GetClass("DIAttachedDeviceInfo")}
	})
	return _DIAttachedDeviceInfoClass
}

// GetDIAttachedDeviceInfoClass returns the class object for DIAttachedDeviceInfo.
func GetDIAttachedDeviceInfoClass() DIAttachedDeviceInfoClass {
	return getDIAttachedDeviceInfoClass()
}

type DIAttachedDeviceInfoClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (dc DIAttachedDeviceInfoClass) Class() objc.Class {
	return dc.class
}

// Alloc allocates memory for a new instance of the class.
func (dc DIAttachedDeviceInfoClass) Alloc() DIAttachedDeviceInfo {
	rv := objc.Send[DIAttachedDeviceInfo](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [DIAttachedDeviceInfo.BSDName]
//   - [DIAttachedDeviceInfo.BlockSize]
//   - [DIAttachedDeviceInfo.CacheURL]
//   - [DIAttachedDeviceInfo.CopyEntitiesList]
//   - [DIAttachedDeviceInfo.FillDI1InfoWithDeviceError]
//   - [DIAttachedDeviceInfo.FillDI2InfoWithDeviceError]
//   - [DIAttachedDeviceInfo.FrameworkNum]
//   - [DIAttachedDeviceInfo.ImageURL]
//   - [DIAttachedDeviceInfo.InstanceId]
//   - [DIAttachedDeviceInfo.IoMedia]
//   - [DIAttachedDeviceInfo.SetIoMedia]
//   - [DIAttachedDeviceInfo.MediaSize]
//   - [DIAttachedDeviceInfo.Pid]
//   - [DIAttachedDeviceInfo.SetDI2PIDWithDeviceError]
//   - [DIAttachedDeviceInfo.ShadowURL]
//   - [DIAttachedDeviceInfo.ToDictionary]
//   - [DIAttachedDeviceInfo.InitWithBSDNameError]
//   - [DIAttachedDeviceInfo.InitWithDeviceError]
// See: https://developer.apple.com/documentation/DiskImages2/DIAttachedDeviceInfo
type DIAttachedDeviceInfo struct {
	objectivec.Object
}

// DIAttachedDeviceInfoFromID constructs a [DIAttachedDeviceInfo] from an objc.ID.
func DIAttachedDeviceInfoFromID(id objc.ID) DIAttachedDeviceInfo {
	return DIAttachedDeviceInfo{objectivec.Object{ID: id}}
}
// Ensure DIAttachedDeviceInfo implements IDIAttachedDeviceInfo.
var _ IDIAttachedDeviceInfo = DIAttachedDeviceInfo{}

// An interface definition for the [DIAttachedDeviceInfo] class.
//
// # Methods
//
//   - [IDIAttachedDeviceInfo.BSDName]
//   - [IDIAttachedDeviceInfo.BlockSize]
//   - [IDIAttachedDeviceInfo.CacheURL]
//   - [IDIAttachedDeviceInfo.CopyEntitiesList]
//   - [IDIAttachedDeviceInfo.FillDI1InfoWithDeviceError]
//   - [IDIAttachedDeviceInfo.FillDI2InfoWithDeviceError]
//   - [IDIAttachedDeviceInfo.FrameworkNum]
//   - [IDIAttachedDeviceInfo.ImageURL]
//   - [IDIAttachedDeviceInfo.InstanceId]
//   - [IDIAttachedDeviceInfo.IoMedia]
//   - [IDIAttachedDeviceInfo.SetIoMedia]
//   - [IDIAttachedDeviceInfo.MediaSize]
//   - [IDIAttachedDeviceInfo.Pid]
//   - [IDIAttachedDeviceInfo.SetDI2PIDWithDeviceError]
//   - [IDIAttachedDeviceInfo.ShadowURL]
//   - [IDIAttachedDeviceInfo.ToDictionary]
//   - [IDIAttachedDeviceInfo.InitWithBSDNameError]
//   - [IDIAttachedDeviceInfo.InitWithDeviceError]
//
// See: https://developer.apple.com/documentation/DiskImages2/DIAttachedDeviceInfo
type IDIAttachedDeviceInfo interface {
	objectivec.IObject

	// Topic: Methods

	BSDName() string
	BlockSize() foundation.NSNumber
	CacheURL() foundation.INSURL
	CopyEntitiesList() objectivec.IObject
	FillDI1InfoWithDeviceError(device objectivec.IObject) (bool, error)
	FillDI2InfoWithDeviceError(device objectivec.IObject) (bool, error)
	FrameworkNum() int64
	ImageURL() foundation.INSURL
	InstanceId() string
	IoMedia() IDIIOMedia
	SetIoMedia(value IDIIOMedia)
	MediaSize() foundation.NSNumber
	Pid() foundation.NSNumber
	SetDI2PIDWithDeviceError(device objectivec.IObject) (bool, error)
	ShadowURL() foundation.INSURL
	ToDictionary() objectivec.IObject
	InitWithBSDNameError(bSDName objectivec.IObject) (DIAttachedDeviceInfo, error)
	InitWithDeviceError(device objectivec.IObject) (DIAttachedDeviceInfo, error)
}

// Init initializes the instance.
func (d DIAttachedDeviceInfo) Init() DIAttachedDeviceInfo {
	rv := objc.Send[DIAttachedDeviceInfo](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d DIAttachedDeviceInfo) Autorelease() DIAttachedDeviceInfo {
	rv := objc.Send[DIAttachedDeviceInfo](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewDIAttachedDeviceInfo creates a new DIAttachedDeviceInfo instance.
func NewDIAttachedDeviceInfo() DIAttachedDeviceInfo {
	class := getDIAttachedDeviceInfoClass()
	rv := objc.Send[DIAttachedDeviceInfo](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/DiskImages2/DIAttachedDeviceInfo/initWithBSDName:error:
func NewDIAttachedDeviceInfoWithBSDNameError(bSDName objectivec.IObject) (DIAttachedDeviceInfo, error) {
	var errorPtr objc.ID
	instance := getDIAttachedDeviceInfoClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithBSDName:error:"), bSDName, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return DIAttachedDeviceInfo{}, foundation.NSErrorFrom(errorPtr)
	}
	return DIAttachedDeviceInfoFromID(rv), nil
}

//
// See: https://developer.apple.com/documentation/DiskImages2/DIAttachedDeviceInfo/initWithDevice:error:
func NewDIAttachedDeviceInfoWithDeviceError(device objectivec.IObject) (DIAttachedDeviceInfo, error) {
	var errorPtr objc.ID
	instance := getDIAttachedDeviceInfoClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:error:"), device, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return DIAttachedDeviceInfo{}, foundation.NSErrorFrom(errorPtr)
	}
	return DIAttachedDeviceInfoFromID(rv), nil
}

// See: https://developer.apple.com/documentation/DiskImages2/DIAttachedDeviceInfo/copyEntitiesList
func (d DIAttachedDeviceInfo) CopyEntitiesList() objectivec.IObject {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("copyEntitiesList"))
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/DiskImages2/DIAttachedDeviceInfo/fillDI1InfoWithDevice:error:
func (d DIAttachedDeviceInfo) FillDI1InfoWithDeviceError(device objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](d.ID, objc.Sel("fillDI1InfoWithDevice:error:"), device, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("fillDI1InfoWithDevice:error: returned NO with nil NSError")
	}
	return rv, nil

}
//
// See: https://developer.apple.com/documentation/DiskImages2/DIAttachedDeviceInfo/fillDI2InfoWithDevice:error:
func (d DIAttachedDeviceInfo) FillDI2InfoWithDeviceError(device objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](d.ID, objc.Sel("fillDI2InfoWithDevice:error:"), device, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("fillDI2InfoWithDevice:error: returned NO with nil NSError")
	}
	return rv, nil

}
//
// See: https://developer.apple.com/documentation/DiskImages2/DIAttachedDeviceInfo/setDI2PIDWithDevice:error:
func (d DIAttachedDeviceInfo) SetDI2PIDWithDeviceError(device objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](d.ID, objc.Sel("setDI2PIDWithDevice:error:"), device, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("setDI2PIDWithDevice:error: returned NO with nil NSError")
	}
	return rv, nil

}
// See: https://developer.apple.com/documentation/DiskImages2/DIAttachedDeviceInfo/toDictionary
func (d DIAttachedDeviceInfo) ToDictionary() objectivec.IObject {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("toDictionary"))
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/DiskImages2/DIAttachedDeviceInfo/initWithBSDName:error:
func (d DIAttachedDeviceInfo) InitWithBSDNameError(bSDName objectivec.IObject) (DIAttachedDeviceInfo, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](d.ID, objc.Sel("initWithBSDName:error:"), bSDName, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return DIAttachedDeviceInfo{}, foundation.NSErrorFrom(errorPtr)
	}
	return DIAttachedDeviceInfoFromID(rv), nil

}
//
// See: https://developer.apple.com/documentation/DiskImages2/DIAttachedDeviceInfo/initWithDevice:error:
func (d DIAttachedDeviceInfo) InitWithDeviceError(device objectivec.IObject) (DIAttachedDeviceInfo, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](d.ID, objc.Sel("initWithDevice:error:"), device, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return DIAttachedDeviceInfo{}, foundation.NSErrorFrom(errorPtr)
	}
	return DIAttachedDeviceInfoFromID(rv), nil

}

//
// See: https://developer.apple.com/documentation/DiskImages2/DIAttachedDeviceInfo/DI1URLWithData:error:
func (_DIAttachedDeviceInfoClass DIAttachedDeviceInfoClass) DI1URLWithDataError(data objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_DIAttachedDeviceInfoClass.class), objc.Sel("DI1URLWithData:error:"), data, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
// See: https://developer.apple.com/documentation/DiskImages2/DIAttachedDeviceInfo/copyAllMountPoints
func (_DIAttachedDeviceInfoClass DIAttachedDeviceInfoClass) CopyAllMountPoints() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_DIAttachedDeviceInfoClass.class), objc.Sel("copyAllMountPoints"))
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/DiskImages2/DIAttachedDeviceInfo/newDI1DevicesArrayWithError:
func (_DIAttachedDeviceInfoClass DIAttachedDeviceInfoClass) NewDI1DevicesArrayWithError() (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_DIAttachedDeviceInfoClass.class), objc.Sel("newDI1DevicesArrayWithError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/DiskImages2/DIAttachedDeviceInfo/newDI2DevicesArrayWithError:
func (_DIAttachedDeviceInfoClass DIAttachedDeviceInfoClass) NewDI2DevicesArrayWithError() (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_DIAttachedDeviceInfoClass.class), objc.Sel("newDI2DevicesArrayWithError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/DiskImages2/DIAttachedDeviceInfo/newDevicesArrayWithError:
func (_DIAttachedDeviceInfoClass DIAttachedDeviceInfoClass) NewDevicesArrayWithError() (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_DIAttachedDeviceInfoClass.class), objc.Sel("newDevicesArrayWithError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/DiskImages2/DIAttachedDeviceInfo/newEntityDictWithIOMedia:mountPoints:
func (_DIAttachedDeviceInfoClass DIAttachedDeviceInfoClass) NewEntityDictWithIOMediaMountPoints(iOMedia objectivec.IObject, points objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_DIAttachedDeviceInfoClass.class), objc.Sel("newEntityDictWithIOMedia:mountPoints:"), iOMedia, points)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/DiskImages2/DIAttachedDeviceInfo/BSDName
func (d DIAttachedDeviceInfo) BSDName() string {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("BSDName"))
	return foundation.NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/DiskImages2/DIAttachedDeviceInfo/blockSize
func (d DIAttachedDeviceInfo) BlockSize() foundation.NSNumber {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("blockSize"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/DiskImages2/DIAttachedDeviceInfo/cacheURL
func (d DIAttachedDeviceInfo) CacheURL() foundation.INSURL {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("cacheURL"))
	return foundation.NSURLFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/DiskImages2/DIAttachedDeviceInfo/frameworkNum
func (d DIAttachedDeviceInfo) FrameworkNum() int64 {
	rv := objc.Send[int64](d.ID, objc.Sel("frameworkNum"))
	return rv
}
// See: https://developer.apple.com/documentation/DiskImages2/DIAttachedDeviceInfo/imageURL
func (d DIAttachedDeviceInfo) ImageURL() foundation.INSURL {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("imageURL"))
	return foundation.NSURLFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/DiskImages2/DIAttachedDeviceInfo/instanceId
func (d DIAttachedDeviceInfo) InstanceId() string {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("instanceId"))
	return foundation.NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/DiskImages2/DIAttachedDeviceInfo/ioMedia
func (d DIAttachedDeviceInfo) IoMedia() IDIIOMedia {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("ioMedia"))
	return DIIOMediaFromID(objc.ID(rv))
}
func (d DIAttachedDeviceInfo) SetIoMedia(value IDIIOMedia) {
	objc.Send[struct{}](d.ID, objc.Sel("setIoMedia:"), value)
}
// See: https://developer.apple.com/documentation/DiskImages2/DIAttachedDeviceInfo/mediaSize
func (d DIAttachedDeviceInfo) MediaSize() foundation.NSNumber {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("mediaSize"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/DiskImages2/DIAttachedDeviceInfo/pid
func (d DIAttachedDeviceInfo) Pid() foundation.NSNumber {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("pid"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/DiskImages2/DIAttachedDeviceInfo/shadowURL
func (d DIAttachedDeviceInfo) ShadowURL() foundation.INSURL {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("shadowURL"))
	return foundation.NSURLFromID(objc.ID(rv))
}

