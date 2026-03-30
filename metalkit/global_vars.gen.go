// Code generated from Apple documentation. DO NOT EDIT.

package metalkit

import (
	"github.com/ebitengine/purego"
	"github.com/tmc/apple/objc"
)

var ()

var (
	// MTKTextureLoaderCubeLayoutVertical is specifies that the source 2D image is a vertical arrangement of six cube faces.
	//
	// See: https://developer.apple.com/documentation/MetalKit/MTKTextureLoader/CubeLayout/vertical
	MTKTextureLoaderCubeLayoutVertical MTKTextureLoaderCubeLayout
)

var ()

var ()

var ()

func init() {
	if frameworkHandle == 0 {
		return
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "MTKModelErrorDomain"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MTKModelErrors.Domain = MTKModelError(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "MTKModelErrorKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MTKModelErrors.Key = MTKModelError(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "MTKTextureLoaderCubeLayoutVertical"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MTKTextureLoaderCubeLayoutVertical = MTKTextureLoaderCubeLayout(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "MTKTextureLoaderErrorDomain"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MTKTextureLoaderErrors.Domain = MTKTextureLoaderError(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "MTKTextureLoaderErrorKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MTKTextureLoaderErrors.Key = MTKTextureLoaderError(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "MTKTextureLoaderOptionAllocateMipmaps"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MTKTextureLoaderOptions.AllocateMipmaps = MTKTextureLoaderOption(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "MTKTextureLoaderOptionCubeLayout"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MTKTextureLoaderOptions.CubeLayout = MTKTextureLoaderOption(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "MTKTextureLoaderOptionGenerateMipmaps"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MTKTextureLoaderOptions.GenerateMipmaps = MTKTextureLoaderOption(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "MTKTextureLoaderOptionLoadAsArray"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MTKTextureLoaderOptions.LoadAsArray = MTKTextureLoaderOption(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "MTKTextureLoaderOptionOrigin"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MTKTextureLoaderOptions.Origin = MTKTextureLoaderOption(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "MTKTextureLoaderOptionSRGB"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MTKTextureLoaderOptions.SRGB = MTKTextureLoaderOption(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "MTKTextureLoaderOptionTextureCPUCacheMode"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MTKTextureLoaderOptions.TextureCPUCacheMode = MTKTextureLoaderOption(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "MTKTextureLoaderOptionTextureStorageMode"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MTKTextureLoaderOptions.TextureStorageMode = MTKTextureLoaderOption(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "MTKTextureLoaderOptionTextureUsage"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MTKTextureLoaderOptions.TextureUsage = MTKTextureLoaderOption(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "MTKTextureLoaderOriginBottomLeft"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MTKTextureLoaderOrigins.BottomLeft = MTKTextureLoaderOrigin(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "MTKTextureLoaderOriginFlippedVertically"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MTKTextureLoaderOrigins.FlippedVertically = MTKTextureLoaderOrigin(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "MTKTextureLoaderOriginTopLeft"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MTKTextureLoaderOrigins.TopLeft = MTKTextureLoaderOrigin(objc.GoString(cstr))
			}
		}
	}

}

// MTKModelErrors provides typed accessors for [MTKModelError] constants.
var MTKModelErrors struct {
	// Domain: The error domain used by MetalKit when returning mesh initialization errors.
	Domain MTKModelError
	// Key: The key used to retrieve an error string from an error object’s [userInfo](<doc://com.apple.documentation/documentation/Foundation/NSError/userInfo>) dictionary.
	Key MTKModelError
}

// MTKTextureLoaderErrors provides typed accessors for [MTKTextureLoaderError] constants.
var MTKTextureLoaderErrors struct {
	// Domain: The error domain used by [MetalKit] when returning texture loading errors.
	Domain MTKTextureLoaderError
	// Key: The key used to retrieve an error string from an error object’s [userInfo](<doc://com.apple.documentation/documentation/Foundation/NSError/userInfo>) dictionary.
	Key MTKTextureLoaderError
}

// MTKTextureLoaderOptions provides typed accessors for [MTKTextureLoaderOption] constants.
var MTKTextureLoaderOptions struct {
	// AllocateMipmaps: A key used to specify whether the texture loader should allocate memory for mipmaps in the texture.
	AllocateMipmaps MTKTextureLoaderOption
	// CubeLayout: A key used to specify how cube texture data is arranged in the source image.
	CubeLayout MTKTextureLoaderOption
	// GenerateMipmaps: A key used to specify whether the texture loader should generate mipmaps for the texture.
	GenerateMipmaps MTKTextureLoaderOption
	LoadAsArray     MTKTextureLoaderOption
	// Origin: A key used to specify when to flip the pixel coordinates of the texture.
	Origin MTKTextureLoaderOption
	// SRGB: A key used to specify whether the texture data is stored as sRGB image data.
	SRGB MTKTextureLoaderOption
	// TextureCPUCacheMode: A key used to specify the CPU cache mode for the texture.
	TextureCPUCacheMode MTKTextureLoaderOption
	// TextureStorageMode: A key used to specify the storage mode for the texture.
	TextureStorageMode MTKTextureLoaderOption
	// TextureUsage: A key used to specify the intended usage of the texture.
	TextureUsage MTKTextureLoaderOption
}

// MTKTextureLoaderOrigins provides typed accessors for [MTKTextureLoaderOrigin] constants.
var MTKTextureLoaderOrigins struct {
	// BottomLeft: An option for specifying images that should be flipped only to put their origin in the bottom-left corner.
	BottomLeft MTKTextureLoaderOrigin
	// FlippedVertically: An option that specifies that images should always be flipped.
	FlippedVertically MTKTextureLoaderOrigin
	// TopLeft: An option for specifying images that should be flipped only to put their origin in the top-left corner.
	TopLeft MTKTextureLoaderOrigin
}
