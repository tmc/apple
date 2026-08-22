// Code generated from Apple documentation. DO NOT EDIT.

package glkit

import (
	"github.com/ebitengine/purego"
	"github.com/tmc/apple/objc"
)

var (
	// GLKMatrix3Identity is a `3x3` identity matrix.
	//
	// See: https://developer.apple.com/documentation/GLKit/GLKMatrix3Identity
	GLKMatrix3Identity GLKMatrix3
)

var (
	// GLKMatrix4Identity is a `4x4` identity matrix.
	//
	// See: https://developer.apple.com/documentation/GLKit/GLKMatrix4Identity
	GLKMatrix4Identity GLKMatrix4
)

var (
	// GLKQuaternionIdentity is an identity quaternion.
	//
	// See: https://developer.apple.com/documentation/GLKit/GLKQuaternionIdentity
	GLKQuaternionIdentity GLKQuaternion
)

var (
	// GLKTextureLoaderApplyPremultiplication is whether image data should be premultiplied before being loaded into the sharegroup.
	//
	// Deprecated: Deprecated.
	//
	// See: https://developer.apple.com/documentation/GLKit/GLKTextureLoaderApplyPremultiplication
	GLKTextureLoaderApplyPremultiplication string
	// GLKTextureLoaderGenerateMipmaps is whether or not to create mipmaps for a texture.
	//
	// Deprecated: Deprecated.
	//
	// See: https://developer.apple.com/documentation/GLKit/GLKTextureLoaderGenerateMipmaps
	GLKTextureLoaderGenerateMipmaps string
	// GLKTextureLoaderGrayscaleAsAlpha is whether or not to treat greyscale image data as alpha information.
	//
	// See: https://developer.apple.com/documentation/GLKit/GLKTextureLoaderGrayscaleAsAlpha
	GLKTextureLoaderGrayscaleAsAlpha string
	// GLKTextureLoaderOriginBottomLeft is whether or not to vertically flip image data to match OpenGL’s coordinate system.
	//
	// Deprecated: Deprecated.
	//
	// See: https://developer.apple.com/documentation/GLKit/GLKTextureLoaderOriginBottomLeft
	GLKTextureLoaderOriginBottomLeft string
)

func init() {
	if frameworkHandle == 0 {
		return
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GLKMatrix3Identity"); err == nil && ptr != 0 {
		GLKMatrix3Identity = objc.ValueAt[GLKMatrix3](ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GLKMatrix4Identity"); err == nil && ptr != 0 {
		GLKMatrix4Identity = objc.ValueAt[GLKMatrix4](ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GLKQuaternionIdentity"); err == nil && ptr != 0 {
		GLKQuaternionIdentity = objc.ValueAt[GLKQuaternion](ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GLKTextureLoaderApplyPremultiplication"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GLKTextureLoaderApplyPremultiplication = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GLKTextureLoaderGenerateMipmaps"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GLKTextureLoaderGenerateMipmaps = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GLKTextureLoaderGrayscaleAsAlpha"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GLKTextureLoaderGrayscaleAsAlpha = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "GLKTextureLoaderOriginBottomLeft"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				GLKTextureLoaderOriginBottomLeft = objc.GoString(cstr)
			}
		}
	}

}
