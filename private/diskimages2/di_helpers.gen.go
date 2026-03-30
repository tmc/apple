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

// The class instance for the [DIHelpers] class.
var (
	_DIHelpersClass     DIHelpersClass
	_DIHelpersClassOnce sync.Once
)

func getDIHelpersClass() DIHelpersClass {
	_DIHelpersClassOnce.Do(func() {
		_DIHelpersClass = DIHelpersClass{class: objc.GetClass("DIHelpers")}
	})
	return _DIHelpersClass
}

// GetDIHelpersClass returns the class object for DIHelpers.
func GetDIHelpersClass() DIHelpersClass {
	return getDIHelpersClass()
}

type DIHelpersClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (dc DIHelpersClass) Class() objc.Class {
	return dc.class
}

// Alloc allocates memory for a new instance of the class.
func (dc DIHelpersClass) Alloc() DIHelpers {
	rv := objc.Send[DIHelpers](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/DiskImages2/DIHelpers
type DIHelpers struct {
	objectivec.Object
}

// DIHelpersFromID constructs a [DIHelpers] from an objc.ID.
func DIHelpersFromID(id objc.ID) DIHelpers {
	return DIHelpers{objectivec.Object{ID: id}}
}

// Ensure DIHelpers implements IDIHelpers.
var _ IDIHelpers = DIHelpers{}

// An interface definition for the [DIHelpers] class.
//
// See: https://developer.apple.com/documentation/DiskImages2/DIHelpers
type IDIHelpers interface {
	objectivec.IObject
}

// Init initializes the instance.
func (d DIHelpers) Init() DIHelpers {
	rv := objc.Send[DIHelpers](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d DIHelpers) Autorelease() DIHelpers {
	rv := objc.Send[DIHelpers](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewDIHelpers creates a new DIHelpers instance.
func NewDIHelpers() DIHelpers {
	class := getDIHelpersClass()
	rv := objc.Send[DIHelpers](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/DiskImages2/DIHelpers/copyDevicePathWithStatfs:
func (_DIHelpersClass DIHelpersClass) CopyDevicePathWithStatfs(statfs objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_DIHelpersClass.class), objc.Sel("copyDevicePathWithStatfs:"), statfs)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/DiskImages2/DIHelpers/executeWithPath:arguments:error:
func (_DIHelpersClass DIHelpersClass) ExecuteWithPathArgumentsError(path objectivec.IObject, arguments objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](objc.ID(_DIHelpersClass.class), objc.Sel("executeWithPath:arguments:error:"), path, arguments, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("executeWithPath:arguments:error: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/DiskImages2/DIHelpers/getBlockSizeFromStr:error:
func (_DIHelpersClass DIHelpersClass) GetBlockSizeFromStrError(str objectivec.IObject) (uint32, error) {
	var errorPtr objc.ID
	rv := objc.Send[uint32](objc.ID(_DIHelpersClass.class), objc.Sel("getBlockSizeFromStr:error:"), str, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return 0, foundation.NSErrorFrom(errorPtr)
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/DiskImages2/DIHelpers/numBlocksWithSizeStr:blockSize:
func (_DIHelpersClass DIHelpersClass) NumBlocksWithSizeStrBlockSize(str objectivec.IObject, size uint32) uint64 {
	rv := objc.Send[uint64](objc.ID(_DIHelpersClass.class), objc.Sel("numBlocksWithSizeStr:blockSize:"), str, size)
	return rv
}

// See: https://developer.apple.com/documentation/DiskImages2/DIHelpers/stringWithImageFormat:
func (_DIHelpersClass DIHelpersClass) StringWithImageFormat(format int64) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_DIHelpersClass.class), objc.Sel("stringWithImageFormat:"), format)
	return objectivec.Object{ID: rv}
}
