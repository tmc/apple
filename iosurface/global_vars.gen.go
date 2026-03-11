// Code generated from Apple documentation. DO NOT EDIT.

package iosurface

import (
	"unsafe"
	"github.com/ebitengine/purego"
	"github.com/tmc/apple/objc"
)

var KIOSurfaceAllocSize string

var KIOSurfaceBytesPerElement string

var KIOSurfaceBytesPerRow string

var KIOSurfaceCacheMode string

var KIOSurfaceColorSpace string

var KIOSurfaceContentHeadroom string

var KIOSurfaceElementHeight string

var KIOSurfaceElementWidth string

var KIOSurfaceHeight string

var KIOSurfaceICCProfile string

var KIOSurfaceName string

var KIOSurfaceOffset string

var KIOSurfacePixelFormat string

var KIOSurfacePixelSizeCastingAllowed string

var KIOSurfacePlaneBase string

var KIOSurfacePlaneBitsPerElement string

var KIOSurfacePlaneBytesPerElement string

var KIOSurfacePlaneBytesPerRow string

var KIOSurfacePlaneComponentBitDepths string

var KIOSurfacePlaneComponentBitOffsets string

var KIOSurfacePlaneComponentNames string

var KIOSurfacePlaneComponentRanges string

var KIOSurfacePlaneComponentTypes string

var KIOSurfacePlaneElementHeight string

var KIOSurfacePlaneElementWidth string

var KIOSurfacePlaneHeight string

var KIOSurfacePlaneInfo string

var KIOSurfacePlaneOffset string

var KIOSurfacePlaneSize string

var KIOSurfacePlaneWidth string

var KIOSurfaceSubsampling string

var KIOSurfaceWidth string

func init() {
	if frameworkHandle == 0 {
		return
	}


	if ptr, err := purego.Dlsym(frameworkHandle, "IOSurfacePropertyKeyAllocSize"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				IOSurfacePropertyKeys.AllocSize = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "IOSurfacePropertyKeyBytesPerElement"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				IOSurfacePropertyKeys.BytesPerElement = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "IOSurfacePropertyKeyBytesPerRow"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				IOSurfacePropertyKeys.BytesPerRow = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "IOSurfacePropertyKeyCacheMode"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				IOSurfacePropertyKeys.CacheMode = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "IOSurfacePropertyKeyElementHeight"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				IOSurfacePropertyKeys.ElementHeight = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "IOSurfacePropertyKeyElementWidth"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				IOSurfacePropertyKeys.ElementWidth = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "IOSurfacePropertyKeyHeight"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				IOSurfacePropertyKeys.Height = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "IOSurfacePropertyKeyName"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				IOSurfacePropertyKeys.Name = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "IOSurfacePropertyKeyOffset"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				IOSurfacePropertyKeys.Offset = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "IOSurfacePropertyKeyPixelFormat"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				IOSurfacePropertyKeys.PixelFormat = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "IOSurfacePropertyKeyPixelSizeCastingAllowed"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				IOSurfacePropertyKeys.PixelSizeCastingAllowed = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "IOSurfacePropertyKeyPlaneBase"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				IOSurfacePropertyKeys.PlaneBase = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "IOSurfacePropertyKeyPlaneBytesPerElement"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				IOSurfacePropertyKeys.PlaneBytesPerElement = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "IOSurfacePropertyKeyPlaneBytesPerRow"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				IOSurfacePropertyKeys.PlaneBytesPerRow = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "IOSurfacePropertyKeyPlaneElementHeight"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				IOSurfacePropertyKeys.PlaneElementHeight = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "IOSurfacePropertyKeyPlaneElementWidth"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				IOSurfacePropertyKeys.PlaneElementWidth = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "IOSurfacePropertyKeyPlaneHeight"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				IOSurfacePropertyKeys.PlaneHeight = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "IOSurfacePropertyKeyPlaneInfo"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				IOSurfacePropertyKeys.PlaneInfo = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "IOSurfacePropertyKeyPlaneOffset"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				IOSurfacePropertyKeys.PlaneOffset = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "IOSurfacePropertyKeyPlaneSize"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				IOSurfacePropertyKeys.PlaneSize = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "IOSurfacePropertyKeyPlaneWidth"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				IOSurfacePropertyKeys.PlaneWidth = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "IOSurfacePropertyKeyWidth"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				IOSurfacePropertyKeys.Width = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kIOSurfaceAllocSize"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KIOSurfaceAllocSize = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kIOSurfaceBytesPerElement"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KIOSurfaceBytesPerElement = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kIOSurfaceBytesPerRow"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KIOSurfaceBytesPerRow = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kIOSurfaceCacheMode"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KIOSurfaceCacheMode = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kIOSurfaceColorSpace"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KIOSurfaceColorSpace = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kIOSurfaceContentHeadroom"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KIOSurfaceContentHeadroom = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kIOSurfaceElementHeight"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KIOSurfaceElementHeight = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kIOSurfaceElementWidth"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KIOSurfaceElementWidth = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kIOSurfaceHeight"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KIOSurfaceHeight = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kIOSurfaceICCProfile"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KIOSurfaceICCProfile = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kIOSurfaceName"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KIOSurfaceName = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kIOSurfaceOffset"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KIOSurfaceOffset = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kIOSurfacePixelFormat"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KIOSurfacePixelFormat = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kIOSurfacePixelSizeCastingAllowed"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KIOSurfacePixelSizeCastingAllowed = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kIOSurfacePlaneBase"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KIOSurfacePlaneBase = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kIOSurfacePlaneBitsPerElement"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KIOSurfacePlaneBitsPerElement = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kIOSurfacePlaneBytesPerElement"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KIOSurfacePlaneBytesPerElement = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kIOSurfacePlaneBytesPerRow"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KIOSurfacePlaneBytesPerRow = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kIOSurfacePlaneComponentBitDepths"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KIOSurfacePlaneComponentBitDepths = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kIOSurfacePlaneComponentBitOffsets"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KIOSurfacePlaneComponentBitOffsets = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kIOSurfacePlaneComponentNames"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KIOSurfacePlaneComponentNames = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kIOSurfacePlaneComponentRanges"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KIOSurfacePlaneComponentRanges = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kIOSurfacePlaneComponentTypes"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KIOSurfacePlaneComponentTypes = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kIOSurfacePlaneElementHeight"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KIOSurfacePlaneElementHeight = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kIOSurfacePlaneElementWidth"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KIOSurfacePlaneElementWidth = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kIOSurfacePlaneHeight"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KIOSurfacePlaneHeight = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kIOSurfacePlaneInfo"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KIOSurfacePlaneInfo = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kIOSurfacePlaneOffset"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KIOSurfacePlaneOffset = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kIOSurfacePlaneSize"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KIOSurfacePlaneSize = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kIOSurfacePlaneWidth"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KIOSurfacePlaneWidth = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kIOSurfaceSubsampling"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KIOSurfaceSubsampling = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kIOSurfaceWidth"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KIOSurfaceWidth = objc.GoString(cstr)
			}
		}
	}

}

// IOSurfacePropertyKeys provides typed accessors for [IOSurfacePropertyKey] constants.
var IOSurfacePropertyKeys struct {
	AllocSize IOSurfacePropertyKey
	BytesPerElement IOSurfacePropertyKey
	BytesPerRow IOSurfacePropertyKey
	CacheMode IOSurfacePropertyKey
	ElementHeight IOSurfacePropertyKey
	ElementWidth IOSurfacePropertyKey
	Height IOSurfacePropertyKey
	Name IOSurfacePropertyKey
	Offset IOSurfacePropertyKey
	PixelFormat IOSurfacePropertyKey
	PixelSizeCastingAllowed IOSurfacePropertyKey
	PlaneBase IOSurfacePropertyKey
	PlaneBytesPerElement IOSurfacePropertyKey
	PlaneBytesPerRow IOSurfacePropertyKey
	PlaneElementHeight IOSurfacePropertyKey
	PlaneElementWidth IOSurfacePropertyKey
	PlaneHeight IOSurfacePropertyKey
	PlaneInfo IOSurfacePropertyKey
	PlaneOffset IOSurfacePropertyKey
	PlaneSize IOSurfacePropertyKey
	PlaneWidth IOSurfacePropertyKey
	Width IOSurfacePropertyKey
}

