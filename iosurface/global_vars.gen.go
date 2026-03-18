// Code generated from Apple documentation. DO NOT EDIT.

package iosurface

import (
	"github.com/ebitengine/purego"
	"github.com/tmc/apple/objc"
)

var (
)

var (
	// KIOSurfaceAllocSize is cFNumber of the total allocation size of the buffer including all planes.
	//
	// See: https://developer.apple.com/documentation/IOSurface/kIOSurfaceAllocSize
	KIOSurfaceAllocSize string
	// KIOSurfaceBytesPerElement is the total number of bytes in an element.
	//
	// See: https://developer.apple.com/documentation/IOSurface/kIOSurfaceBytesPerElement
	KIOSurfaceBytesPerElement string
	// KIOSurfaceBytesPerRow is the bytes per row of the buffer.
	//
	// See: https://developer.apple.com/documentation/IOSurface/kIOSurfaceBytesPerRow
	KIOSurfaceBytesPerRow string
	// KIOSurfaceCacheMode is the CPU cache mode to be used for the allocation.
	//
	// See: https://developer.apple.com/documentation/IOSurface/kIOSurfaceCacheMode
	KIOSurfaceCacheMode string
	// See: https://developer.apple.com/documentation/IOSurface/kIOSurfaceColorSpace
	KIOSurfaceColorSpace string
	// See: https://developer.apple.com/documentation/IOSurface/kIOSurfaceContentHeadroom
	KIOSurfaceContentHeadroom string
	// KIOSurfaceElementHeight is cFNumber for how many pixels high each element is.
	//
	// See: https://developer.apple.com/documentation/IOSurface/kIOSurfaceElementHeight
	KIOSurfaceElementHeight string
	// KIOSurfaceElementWidth is cFNumber for how many pixels wide each element is.
	//
	// See: https://developer.apple.com/documentation/IOSurface/kIOSurfaceElementWidth
	KIOSurfaceElementWidth string
	// KIOSurfaceHeight is the height of the IOSurface buffer in pixels.
	//
	// See: https://developer.apple.com/documentation/IOSurface/kIOSurfaceHeight
	KIOSurfaceHeight string
	// See: https://developer.apple.com/documentation/IOSurface/kIOSurfaceICCProfile
	KIOSurfaceICCProfile string
	// See: https://developer.apple.com/documentation/IOSurface/kIOSurfaceName
	KIOSurfaceName string
	// KIOSurfaceOffset is the starting offset into the buffer.
	//
	// See: https://developer.apple.com/documentation/IOSurface/kIOSurfaceOffset
	KIOSurfaceOffset string
	// KIOSurfacePixelFormat is a 32-bit unsigned integer that stores the traditional macOS buffer format.
	//
	// See: https://developer.apple.com/documentation/IOSurface/kIOSurfacePixelFormat
	KIOSurfacePixelFormat string
	// See: https://developer.apple.com/documentation/IOSurface/kIOSurfacePixelSizeCastingAllowed
	KIOSurfacePixelSizeCastingAllowed string
	// KIOSurfacePlaneBase is the base offset into the buffer for this plane.
	//
	// See: https://developer.apple.com/documentation/IOSurface/kIOSurfacePlaneBase
	KIOSurfacePlaneBase string
	// See: https://developer.apple.com/documentation/IOSurface/kIOSurfacePlaneBitsPerElement
	KIOSurfacePlaneBitsPerElement string
	// KIOSurfacePlaneBytesPerElement is the bytes per element of this plane.
	//
	// See: https://developer.apple.com/documentation/IOSurface/kIOSurfacePlaneBytesPerElement
	KIOSurfacePlaneBytesPerElement string
	// KIOSurfacePlaneBytesPerRow is the bytes per row of this plane.
	//
	// See: https://developer.apple.com/documentation/IOSurface/kIOSurfacePlaneBytesPerRow
	KIOSurfacePlaneBytesPerRow string
	// See: https://developer.apple.com/documentation/IOSurface/kIOSurfacePlaneComponentBitDepths
	KIOSurfacePlaneComponentBitDepths string
	// See: https://developer.apple.com/documentation/IOSurface/kIOSurfacePlaneComponentBitOffsets
	KIOSurfacePlaneComponentBitOffsets string
	// See: https://developer.apple.com/documentation/IOSurface/kIOSurfacePlaneComponentNames
	KIOSurfacePlaneComponentNames string
	// See: https://developer.apple.com/documentation/IOSurface/kIOSurfacePlaneComponentRanges
	KIOSurfacePlaneComponentRanges string
	// See: https://developer.apple.com/documentation/IOSurface/kIOSurfacePlaneComponentTypes
	KIOSurfacePlaneComponentTypes string
	// KIOSurfacePlaneElementHeight is the element height of this plane.
	//
	// See: https://developer.apple.com/documentation/IOSurface/kIOSurfacePlaneElementHeight
	KIOSurfacePlaneElementHeight string
	// KIOSurfacePlaneElementWidth is the element width of this plane.
	//
	// See: https://developer.apple.com/documentation/IOSurface/kIOSurfacePlaneElementWidth
	KIOSurfacePlaneElementWidth string
	// KIOSurfacePlaneHeight is the height of this plane in pixels.
	//
	// See: https://developer.apple.com/documentation/IOSurface/kIOSurfacePlaneHeight
	KIOSurfacePlaneHeight string
	// KIOSurfacePlaneInfo is cFArray describing each image plane in the buffer as a CFDictionary.
	//
	// See: https://developer.apple.com/documentation/IOSurface/kIOSurfacePlaneInfo
	KIOSurfacePlaneInfo string
	// KIOSurfacePlaneOffset is the offset into the buffer for this plane.
	//
	// See: https://developer.apple.com/documentation/IOSurface/kIOSurfacePlaneOffset
	KIOSurfacePlaneOffset string
	// KIOSurfacePlaneSize is the total data size of this plane.
	//
	// See: https://developer.apple.com/documentation/IOSurface/kIOSurfacePlaneSize
	KIOSurfacePlaneSize string
	// KIOSurfacePlaneWidth is the width of this plane in pixels.
	//
	// See: https://developer.apple.com/documentation/IOSurface/kIOSurfacePlaneWidth
	KIOSurfacePlaneWidth string
	// See: https://developer.apple.com/documentation/IOSurface/kIOSurfaceSubsampling
	KIOSurfaceSubsampling string
	// KIOSurfaceWidth is the width of the IOSurface buffer in pixels.
	//
	// See: https://developer.apple.com/documentation/IOSurface/kIOSurfaceWidth
	KIOSurfaceWidth string
)
func init() {
	if frameworkHandle == 0 {
		return
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "IOSurfacePropertyKeyAllocSize"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				IOSurfacePropertyKeys.AllocSize = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "IOSurfacePropertyKeyBytesPerElement"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				IOSurfacePropertyKeys.BytesPerElement = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "IOSurfacePropertyKeyBytesPerRow"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				IOSurfacePropertyKeys.BytesPerRow = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "IOSurfacePropertyKeyCacheMode"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				IOSurfacePropertyKeys.CacheMode = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "IOSurfacePropertyKeyElementHeight"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				IOSurfacePropertyKeys.ElementHeight = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "IOSurfacePropertyKeyElementWidth"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				IOSurfacePropertyKeys.ElementWidth = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "IOSurfacePropertyKeyHeight"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				IOSurfacePropertyKeys.Height = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "IOSurfacePropertyKeyName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				IOSurfacePropertyKeys.Name = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "IOSurfacePropertyKeyOffset"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				IOSurfacePropertyKeys.Offset = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "IOSurfacePropertyKeyPixelFormat"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				IOSurfacePropertyKeys.PixelFormat = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "IOSurfacePropertyKeyPixelSizeCastingAllowed"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				IOSurfacePropertyKeys.PixelSizeCastingAllowed = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "IOSurfacePropertyKeyPlaneBase"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				IOSurfacePropertyKeys.PlaneBase = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "IOSurfacePropertyKeyPlaneBytesPerElement"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				IOSurfacePropertyKeys.PlaneBytesPerElement = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "IOSurfacePropertyKeyPlaneBytesPerRow"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				IOSurfacePropertyKeys.PlaneBytesPerRow = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "IOSurfacePropertyKeyPlaneElementHeight"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				IOSurfacePropertyKeys.PlaneElementHeight = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "IOSurfacePropertyKeyPlaneElementWidth"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				IOSurfacePropertyKeys.PlaneElementWidth = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "IOSurfacePropertyKeyPlaneHeight"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				IOSurfacePropertyKeys.PlaneHeight = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "IOSurfacePropertyKeyPlaneInfo"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				IOSurfacePropertyKeys.PlaneInfo = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "IOSurfacePropertyKeyPlaneOffset"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				IOSurfacePropertyKeys.PlaneOffset = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "IOSurfacePropertyKeyPlaneSize"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				IOSurfacePropertyKeys.PlaneSize = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "IOSurfacePropertyKeyPlaneWidth"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				IOSurfacePropertyKeys.PlaneWidth = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "IOSurfacePropertyKeyWidth"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				IOSurfacePropertyKeys.Width = string(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kIOSurfaceAllocSize"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KIOSurfaceAllocSize = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kIOSurfaceBytesPerElement"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KIOSurfaceBytesPerElement = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kIOSurfaceBytesPerRow"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KIOSurfaceBytesPerRow = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kIOSurfaceCacheMode"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KIOSurfaceCacheMode = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kIOSurfaceColorSpace"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KIOSurfaceColorSpace = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kIOSurfaceContentHeadroom"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KIOSurfaceContentHeadroom = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kIOSurfaceElementHeight"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KIOSurfaceElementHeight = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kIOSurfaceElementWidth"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KIOSurfaceElementWidth = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kIOSurfaceHeight"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KIOSurfaceHeight = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kIOSurfaceICCProfile"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KIOSurfaceICCProfile = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kIOSurfaceName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KIOSurfaceName = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kIOSurfaceOffset"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KIOSurfaceOffset = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kIOSurfacePixelFormat"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KIOSurfacePixelFormat = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kIOSurfacePixelSizeCastingAllowed"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KIOSurfacePixelSizeCastingAllowed = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kIOSurfacePlaneBase"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KIOSurfacePlaneBase = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kIOSurfacePlaneBitsPerElement"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KIOSurfacePlaneBitsPerElement = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kIOSurfacePlaneBytesPerElement"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KIOSurfacePlaneBytesPerElement = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kIOSurfacePlaneBytesPerRow"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KIOSurfacePlaneBytesPerRow = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kIOSurfacePlaneComponentBitDepths"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KIOSurfacePlaneComponentBitDepths = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kIOSurfacePlaneComponentBitOffsets"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KIOSurfacePlaneComponentBitOffsets = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kIOSurfacePlaneComponentNames"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KIOSurfacePlaneComponentNames = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kIOSurfacePlaneComponentRanges"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KIOSurfacePlaneComponentRanges = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kIOSurfacePlaneComponentTypes"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KIOSurfacePlaneComponentTypes = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kIOSurfacePlaneElementHeight"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KIOSurfacePlaneElementHeight = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kIOSurfacePlaneElementWidth"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KIOSurfacePlaneElementWidth = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kIOSurfacePlaneHeight"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KIOSurfacePlaneHeight = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kIOSurfacePlaneInfo"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KIOSurfacePlaneInfo = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kIOSurfacePlaneOffset"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KIOSurfacePlaneOffset = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kIOSurfacePlaneSize"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KIOSurfacePlaneSize = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kIOSurfacePlaneWidth"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KIOSurfacePlaneWidth = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kIOSurfaceSubsampling"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KIOSurfaceSubsampling = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kIOSurfaceWidth"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
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

