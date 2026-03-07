// Code generated from Apple documentation. DO NOT EDIT.

package metalkit

import (
	"unsafe"
	"github.com/ebitengine/purego"
	"github.com/tmc/apple/objc"
)

var mTKModelErrorDomain MTKModelError

var mTKModelErrorKey MTKModelError

var MTKTextureLoaderCubeLayoutVertical MTKTextureLoaderCubeLayout

var mTKTextureLoaderErrorDomain MTKTextureLoaderError

var mTKTextureLoaderErrorKey MTKTextureLoaderError

var mTKTextureLoaderOptionAllocateMipmaps MTKTextureLoaderOption

var mTKTextureLoaderOptionCubeLayout MTKTextureLoaderOption

var mTKTextureLoaderOptionGenerateMipmaps MTKTextureLoaderOption

var mTKTextureLoaderOptionLoadAsArray MTKTextureLoaderOption

var mTKTextureLoaderOptionOrigin MTKTextureLoaderOption

var mTKTextureLoaderOptionSRGB MTKTextureLoaderOption

var mTKTextureLoaderOptionTextureCPUCacheMode MTKTextureLoaderOption

var mTKTextureLoaderOptionTextureStorageMode MTKTextureLoaderOption

var mTKTextureLoaderOptionTextureUsage MTKTextureLoaderOption

var mTKTextureLoaderOriginBottomLeft MTKTextureLoaderOrigin

var mTKTextureLoaderOriginFlippedVertically MTKTextureLoaderOrigin

var mTKTextureLoaderOriginTopLeft MTKTextureLoaderOrigin

func init() {
	if frameworkHandle == 0 {
		return
	}


	if ptr, err := purego.Dlsym(frameworkHandle, "MTKModelErrorDomain"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				mTKModelErrorDomain = MTKModelError(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "MTKModelErrorKey"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				mTKModelErrorKey = MTKModelError(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "MTKTextureLoaderCubeLayoutVertical"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MTKTextureLoaderCubeLayoutVertical = MTKTextureLoaderCubeLayout(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "MTKTextureLoaderErrorDomain"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				mTKTextureLoaderErrorDomain = MTKTextureLoaderError(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "MTKTextureLoaderErrorKey"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				mTKTextureLoaderErrorKey = MTKTextureLoaderError(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "MTKTextureLoaderOptionAllocateMipmaps"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				mTKTextureLoaderOptionAllocateMipmaps = MTKTextureLoaderOption(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "MTKTextureLoaderOptionCubeLayout"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				mTKTextureLoaderOptionCubeLayout = MTKTextureLoaderOption(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "MTKTextureLoaderOptionGenerateMipmaps"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				mTKTextureLoaderOptionGenerateMipmaps = MTKTextureLoaderOption(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "MTKTextureLoaderOptionLoadAsArray"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				mTKTextureLoaderOptionLoadAsArray = MTKTextureLoaderOption(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "MTKTextureLoaderOptionOrigin"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				mTKTextureLoaderOptionOrigin = MTKTextureLoaderOption(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "MTKTextureLoaderOptionSRGB"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				mTKTextureLoaderOptionSRGB = MTKTextureLoaderOption(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "MTKTextureLoaderOptionTextureCPUCacheMode"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				mTKTextureLoaderOptionTextureCPUCacheMode = MTKTextureLoaderOption(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "MTKTextureLoaderOptionTextureStorageMode"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				mTKTextureLoaderOptionTextureStorageMode = MTKTextureLoaderOption(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "MTKTextureLoaderOptionTextureUsage"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				mTKTextureLoaderOptionTextureUsage = MTKTextureLoaderOption(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "MTKTextureLoaderOriginBottomLeft"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				mTKTextureLoaderOriginBottomLeft = MTKTextureLoaderOrigin(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "MTKTextureLoaderOriginFlippedVertically"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				mTKTextureLoaderOriginFlippedVertically = MTKTextureLoaderOrigin(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "MTKTextureLoaderOriginTopLeft"); err == nil && ptr != 0 {
		nsStringID := *(*objc.ID)(unsafe.Pointer(ptr))
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				mTKTextureLoaderOriginTopLeft = MTKTextureLoaderOrigin(objc.GoString(cstr))
			}
		}
	}

}

type MTKModelErrorValues struct{}

// MTKModelErrors provides typed accessors for [MTKModelError] constants.
var MTKModelErrors MTKModelErrorValues

// Domain returns The error domain used by MetalKit when returning mesh initialization errors.
func (MTKModelErrorValues) Domain() MTKModelError { return mTKModelErrorDomain }

// Key returns The key used to retrieve an error string from an error object’s [userInfo](<doc://com.apple.documentation/documentation/Foundation/NSError/userInfo>) dictionary.
func (MTKModelErrorValues) Key() MTKModelError { return mTKModelErrorKey }


type MTKTextureLoaderErrorValues struct{}

// MTKTextureLoaderErrors provides typed accessors for [MTKTextureLoaderError] constants.
var MTKTextureLoaderErrors MTKTextureLoaderErrorValues

// Domain returns The error domain used by [MetalKit] when returning texture loading errors.
func (MTKTextureLoaderErrorValues) Domain() MTKTextureLoaderError { return mTKTextureLoaderErrorDomain }

// Key returns The key used to retrieve an error string from an error object’s [userInfo](<doc://com.apple.documentation/documentation/Foundation/NSError/userInfo>) dictionary.
func (MTKTextureLoaderErrorValues) Key() MTKTextureLoaderError { return mTKTextureLoaderErrorKey }


type MTKTextureLoaderOptionValues struct{}

// MTKTextureLoaderOptions provides typed accessors for [MTKTextureLoaderOption] constants.
var MTKTextureLoaderOptions MTKTextureLoaderOptionValues

// AllocateMipmaps returns A key used to specify whether the texture loader should allocate memory for mipmaps in the texture.
func (MTKTextureLoaderOptionValues) AllocateMipmaps() MTKTextureLoaderOption { return mTKTextureLoaderOptionAllocateMipmaps }

// CubeLayout returns A key used to specify how cube texture data is arranged in the source image.
func (MTKTextureLoaderOptionValues) CubeLayout() MTKTextureLoaderOption { return mTKTextureLoaderOptionCubeLayout }

// GenerateMipmaps returns A key used to specify whether the texture loader should generate mipmaps for the texture.
func (MTKTextureLoaderOptionValues) GenerateMipmaps() MTKTextureLoaderOption { return mTKTextureLoaderOptionGenerateMipmaps }

func (MTKTextureLoaderOptionValues) LoadAsArray() MTKTextureLoaderOption { return mTKTextureLoaderOptionLoadAsArray }

// Origin returns A key used to specify when to flip the pixel coordinates of the texture.
func (MTKTextureLoaderOptionValues) Origin() MTKTextureLoaderOption { return mTKTextureLoaderOptionOrigin }

// SRGB returns A key used to specify whether the texture data is stored as sRGB image data.
func (MTKTextureLoaderOptionValues) SRGB() MTKTextureLoaderOption { return mTKTextureLoaderOptionSRGB }

// TextureCPUCacheMode returns A key used to specify the CPU cache mode for the texture.
func (MTKTextureLoaderOptionValues) TextureCPUCacheMode() MTKTextureLoaderOption { return mTKTextureLoaderOptionTextureCPUCacheMode }

// TextureStorageMode returns A key used to specify the storage mode for the texture.
func (MTKTextureLoaderOptionValues) TextureStorageMode() MTKTextureLoaderOption { return mTKTextureLoaderOptionTextureStorageMode }

// TextureUsage returns A key used to specify the intended usage of the texture.
func (MTKTextureLoaderOptionValues) TextureUsage() MTKTextureLoaderOption { return mTKTextureLoaderOptionTextureUsage }


type MTKTextureLoaderOriginValues struct{}

// MTKTextureLoaderOrigins provides typed accessors for [MTKTextureLoaderOrigin] constants.
var MTKTextureLoaderOrigins MTKTextureLoaderOriginValues

// BottomLeft returns An option for specifying images that should be flipped only to put their origin in the bottom-left corner.
func (MTKTextureLoaderOriginValues) BottomLeft() MTKTextureLoaderOrigin { return mTKTextureLoaderOriginBottomLeft }

// FlippedVertically returns An option that specifies that images should always be flipped.
func (MTKTextureLoaderOriginValues) FlippedVertically() MTKTextureLoaderOrigin { return mTKTextureLoaderOriginFlippedVertically }

// TopLeft returns An option for specifying images that should be flipped only to put their origin in the top-left corner.
func (MTKTextureLoaderOriginValues) TopLeft() MTKTextureLoaderOrigin { return mTKTextureLoaderOriginTopLeft }


