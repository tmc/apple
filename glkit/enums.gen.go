// Code generated from Apple documentation for GLKit. DO NOT EDIT.

package glkit

import (
	"fmt"
)

// See: https://developer.apple.com/documentation/GLKit/GLKFogMode
type GLKFogMode int32

const (
	// GLKFogModeExp: The fog component is calculated as  and clamped to the range .
	GLKFogModeExp GLKFogMode = 0
	// GLKFogModeExp2: The fog component is calculated as  and clamped to the range .
	GLKFogModeExp2 GLKFogMode = 1
	// GLKFogModeLinear: The fog component is calculated as  and clamped to the range .
	GLKFogModeLinear GLKFogMode = 2
)

func (e GLKFogMode) String() string {
	switch e {
	case GLKFogModeExp:
		return "GLKFogModeExp"
	case GLKFogModeExp2:
		return "GLKFogModeExp2"
	case GLKFogModeLinear:
		return "GLKFogModeLinear"
	default:
		return fmt.Sprintf("GLKFogMode(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/GLKit/GLKLightingType
type GLKLightingType int32

const (
	// GLKLightingTypePerPixel: Indicates that the inputs to the lighting calculation are interpolated across a triangle and the lighting calculations are performed at each fragment.
	GLKLightingTypePerPixel GLKLightingType = 1
	// GLKLightingTypePerVertex: Indicates that the lighting calculations are performed at each vertex in a triangle and then interpolated across the triangle.
	GLKLightingTypePerVertex GLKLightingType = 0
)

func (e GLKLightingType) String() string {
	switch e {
	case GLKLightingTypePerPixel:
		return "GLKLightingTypePerPixel"
	case GLKLightingTypePerVertex:
		return "GLKLightingTypePerVertex"
	default:
		return fmt.Sprintf("GLKLightingType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/GLKit/GLKTextureEnvMode
type GLKTextureEnvMode int32

const (
	// GLKTextureEnvModeDecal: The output color is calculated by using the texture’s alpha component to blend the texture’s color with the input color.
	GLKTextureEnvModeDecal GLKTextureEnvMode = 2
	// GLKTextureEnvModeModulate: The output color is calculated by multiplying the texture’s color by the input color.
	GLKTextureEnvModeModulate GLKTextureEnvMode = 1
	// GLKTextureEnvModeReplace: The output color is set to the color fetched from the texture.
	GLKTextureEnvModeReplace GLKTextureEnvMode = 0
)

func (e GLKTextureEnvMode) String() string {
	switch e {
	case GLKTextureEnvModeDecal:
		return "GLKTextureEnvModeDecal"
	case GLKTextureEnvModeModulate:
		return "GLKTextureEnvModeModulate"
	case GLKTextureEnvModeReplace:
		return "GLKTextureEnvModeReplace"
	default:
		return fmt.Sprintf("GLKTextureEnvMode(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/GLKit/GLKTextureInfoAlphaState
type GLKTextureInfoAlphaState int32

const (
	// GLKTextureInfoAlphaStateNonPremultiplied: Indicates that the color values in the texture were not premultiplied by the alpha value.
	GLKTextureInfoAlphaStateNonPremultiplied GLKTextureInfoAlphaState = 1
	// GLKTextureInfoAlphaStateNone: Indicates that the texture has no alpha information.
	GLKTextureInfoAlphaStateNone GLKTextureInfoAlphaState = 0
	// GLKTextureInfoAlphaStatePremultiplied: Indicates that the color values in the texture have already been premultiplied by the alpha value.
	GLKTextureInfoAlphaStatePremultiplied GLKTextureInfoAlphaState = 2
)

func (e GLKTextureInfoAlphaState) String() string {
	switch e {
	case GLKTextureInfoAlphaStateNonPremultiplied:
		return "GLKTextureInfoAlphaStateNonPremultiplied"
	case GLKTextureInfoAlphaStateNone:
		return "GLKTextureInfoAlphaStateNone"
	case GLKTextureInfoAlphaStatePremultiplied:
		return "GLKTextureInfoAlphaStatePremultiplied"
	default:
		return fmt.Sprintf("GLKTextureInfoAlphaState(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/GLKit/GLKTextureInfoOrigin
type GLKTextureInfoOrigin int32

const (
	// GLKTextureInfoOriginBottomLeft: The origin of the texture is in the bottom-left corner.
	GLKTextureInfoOriginBottomLeft GLKTextureInfoOrigin = 2
	// GLKTextureInfoOriginTopLeft: The origin of the texture is in the top-left corner.
	GLKTextureInfoOriginTopLeft GLKTextureInfoOrigin = 1
	// GLKTextureInfoOriginUnknown: The origin of the texture is not supported.
	GLKTextureInfoOriginUnknown GLKTextureInfoOrigin = 0
)

func (e GLKTextureInfoOrigin) String() string {
	switch e {
	case GLKTextureInfoOriginBottomLeft:
		return "GLKTextureInfoOriginBottomLeft"
	case GLKTextureInfoOriginTopLeft:
		return "GLKTextureInfoOriginTopLeft"
	case GLKTextureInfoOriginUnknown:
		return "GLKTextureInfoOriginUnknown"
	default:
		return fmt.Sprintf("GLKTextureInfoOrigin(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/GLKit/GLKTextureLoaderError-swift.struct/Code
type GLKTextureLoaderError uint32

const (
	// GLKTextureLoaderErrorAlphaPremultiplicationFailure: The texture source data does not allow the alpha to be premultiplied.
	GLKTextureLoaderErrorAlphaPremultiplicationFailure GLKTextureLoaderError = 16
	// GLKTextureLoaderErrorCompressedTextureUpload: A compressed texture could not be uploaded.
	GLKTextureLoaderErrorCompressedTextureUpload GLKTextureLoaderError = 7
	// GLKTextureLoaderErrorCubeMapInvalidNumFiles: The incorrect number of files were specified for the cube map.
	GLKTextureLoaderErrorCubeMapInvalidNumFiles GLKTextureLoaderError = 6
	// GLKTextureLoaderErrorDataPreprocessingFailure: The data could not be preprocessed correctly.
	GLKTextureLoaderErrorDataPreprocessingFailure GLKTextureLoaderError = 12
	// GLKTextureLoaderErrorFileOrURLNotFound: A file could not be found at the path provided.
	GLKTextureLoaderErrorFileOrURLNotFound GLKTextureLoaderError = 0
	// GLKTextureLoaderErrorIncompatibleFormatSRGB: The decoded data was in an incompatible format for an sRGB texture.
	GLKTextureLoaderErrorIncompatibleFormatSRGB GLKTextureLoaderError = 18
	// GLKTextureLoaderErrorInvalidCGImage: The image provided was invalid.
	GLKTextureLoaderErrorInvalidCGImage GLKTextureLoaderError = 2
	// GLKTextureLoaderErrorInvalidEAGLContext: The EAGL context was not a valid context.
	GLKTextureLoaderErrorInvalidEAGLContext GLKTextureLoaderError = 17
	// GLKTextureLoaderErrorInvalidNSData: The data provided is not in a recognized image format.
	GLKTextureLoaderErrorInvalidNSData GLKTextureLoaderError = 1
	// GLKTextureLoaderErrorMipmapUnsupported: The texture source data does not allow mipmaps to be generated.
	GLKTextureLoaderErrorMipmapUnsupported GLKTextureLoaderError = 13
	// GLKTextureLoaderErrorPVRAtlasUnsupported: Cube maps may not be compressed in PVRTC format.
	GLKTextureLoaderErrorPVRAtlasUnsupported GLKTextureLoaderError = 5
	// GLKTextureLoaderErrorReorientationFailure: The texture source data does not allow the image to be reoriented.
	GLKTextureLoaderErrorReorientationFailure GLKTextureLoaderError = 15
	// GLKTextureLoaderErrorUncompressedTextureUpload: An uncompressed texture could not be uploaded.
	GLKTextureLoaderErrorUncompressedTextureUpload GLKTextureLoaderError = 8
	// GLKTextureLoaderErrorUnknownFileType: The file was in an unrecognized format.
	GLKTextureLoaderErrorUnknownFileType GLKTextureLoaderError = 4
	// GLKTextureLoaderErrorUnknownPathType: The path type was unrecognized.
	GLKTextureLoaderErrorUnknownPathType GLKTextureLoaderError = 3
	// GLKTextureLoaderErrorUnsupportedBitDepth: The data in the source image has an unsupported bit depth.
	GLKTextureLoaderErrorUnsupportedBitDepth GLKTextureLoaderError = 10
	// GLKTextureLoaderErrorUnsupportedCubeMapDimensions: The cube map’s dimensions are incorrect.
	GLKTextureLoaderErrorUnsupportedCubeMapDimensions GLKTextureLoaderError = 9
	// GLKTextureLoaderErrorUnsupportedOrientation: The texture source data is stored with an unsupported origin position.
	GLKTextureLoaderErrorUnsupportedOrientation GLKTextureLoaderError = 14
	// GLKTextureLoaderErrorUnsupportedPVRFormat: The data in the PVRTC compressed format is in an unsupported format.
	GLKTextureLoaderErrorUnsupportedPVRFormat     GLKTextureLoaderError = 11
	GLKTextureLoaderErrorUnsupportedTextureTarget GLKTextureLoaderError = 19
)

func (e GLKTextureLoaderError) String() string {
	switch e {
	case GLKTextureLoaderErrorAlphaPremultiplicationFailure:
		return "GLKTextureLoaderErrorAlphaPremultiplicationFailure"
	case GLKTextureLoaderErrorCompressedTextureUpload:
		return "GLKTextureLoaderErrorCompressedTextureUpload"
	case GLKTextureLoaderErrorCubeMapInvalidNumFiles:
		return "GLKTextureLoaderErrorCubeMapInvalidNumFiles"
	case GLKTextureLoaderErrorDataPreprocessingFailure:
		return "GLKTextureLoaderErrorDataPreprocessingFailure"
	case GLKTextureLoaderErrorFileOrURLNotFound:
		return "GLKTextureLoaderErrorFileOrURLNotFound"
	case GLKTextureLoaderErrorIncompatibleFormatSRGB:
		return "GLKTextureLoaderErrorIncompatibleFormatSRGB"
	case GLKTextureLoaderErrorInvalidCGImage:
		return "GLKTextureLoaderErrorInvalidCGImage"
	case GLKTextureLoaderErrorInvalidEAGLContext:
		return "GLKTextureLoaderErrorInvalidEAGLContext"
	case GLKTextureLoaderErrorInvalidNSData:
		return "GLKTextureLoaderErrorInvalidNSData"
	case GLKTextureLoaderErrorMipmapUnsupported:
		return "GLKTextureLoaderErrorMipmapUnsupported"
	case GLKTextureLoaderErrorPVRAtlasUnsupported:
		return "GLKTextureLoaderErrorPVRAtlasUnsupported"
	case GLKTextureLoaderErrorReorientationFailure:
		return "GLKTextureLoaderErrorReorientationFailure"
	case GLKTextureLoaderErrorUncompressedTextureUpload:
		return "GLKTextureLoaderErrorUncompressedTextureUpload"
	case GLKTextureLoaderErrorUnknownFileType:
		return "GLKTextureLoaderErrorUnknownFileType"
	case GLKTextureLoaderErrorUnknownPathType:
		return "GLKTextureLoaderErrorUnknownPathType"
	case GLKTextureLoaderErrorUnsupportedBitDepth:
		return "GLKTextureLoaderErrorUnsupportedBitDepth"
	case GLKTextureLoaderErrorUnsupportedCubeMapDimensions:
		return "GLKTextureLoaderErrorUnsupportedCubeMapDimensions"
	case GLKTextureLoaderErrorUnsupportedOrientation:
		return "GLKTextureLoaderErrorUnsupportedOrientation"
	case GLKTextureLoaderErrorUnsupportedPVRFormat:
		return "GLKTextureLoaderErrorUnsupportedPVRFormat"
	case GLKTextureLoaderErrorUnsupportedTextureTarget:
		return "GLKTextureLoaderErrorUnsupportedTextureTarget"
	default:
		return fmt.Sprintf("GLKTextureLoaderError(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/GLKit/GLKTextureTarget
type GLKTextureTarget uint32

const (
	// GLKTextureTarget2D: The texture is a 2D texture.
	GLKTextureTarget2D GLKTextureTarget = 0xde1
	// GLKTextureTargetCt: The number of items in the enumeration.
	GLKTextureTargetCt GLKTextureTarget = 2
	// GLKTextureTargetCubeMap: The texture is a set of six textures that make up a cube map.
	GLKTextureTargetCubeMap GLKTextureTarget = 0x8513
)

func (e GLKTextureTarget) String() string {
	switch e {
	case GLKTextureTarget2D:
		return "GLKTextureTarget2D"
	case GLKTextureTargetCt:
		return "GLKTextureTargetCt"
	case GLKTextureTargetCubeMap:
		return "GLKTextureTargetCubeMap"
	default:
		return fmt.Sprintf("GLKTextureTarget(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/GLKit/GLKVertexAttrib
type GLKVertexAttrib int32

const (
	// GLKVertexAttribColor: This index is used to provide the vertex color to a shader.
	GLKVertexAttribColor GLKVertexAttrib = 2
	// GLKVertexAttribNormal: This index is used to provide the vertex normal to a shader.
	GLKVertexAttribNormal GLKVertexAttrib = 1
	// GLKVertexAttribPosition: This index is used to provide the vertex position to a shader.
	GLKVertexAttribPosition GLKVertexAttrib = 0
	// GLKVertexAttribTexCoord0: This index is used to provide a set of texture coordinates to a shader.
	GLKVertexAttribTexCoord0 GLKVertexAttrib = 3
	// GLKVertexAttribTexCoord1: This index is used to provide the second set of texture coordinates to a shader.
	GLKVertexAttribTexCoord1 GLKVertexAttrib = 4
)

func (e GLKVertexAttrib) String() string {
	switch e {
	case GLKVertexAttribColor:
		return "GLKVertexAttribColor"
	case GLKVertexAttribNormal:
		return "GLKVertexAttribNormal"
	case GLKVertexAttribPosition:
		return "GLKVertexAttribPosition"
	case GLKVertexAttribTexCoord0:
		return "GLKVertexAttribTexCoord0"
	case GLKVertexAttribTexCoord1:
		return "GLKVertexAttribTexCoord1"
	default:
		return fmt.Sprintf("GLKVertexAttrib(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/GLKit/GLKViewDrawableColorFormat
type GLKViewDrawableColorFormat int

const (
	// GLKViewDrawableColorFormatRGB565: An RGB565 format.
	GLKViewDrawableColorFormatRGB565 GLKViewDrawableColorFormat = 0
	// GLKViewDrawableColorFormatRGBA8888: An RGBA8888 format.
	GLKViewDrawableColorFormatRGBA8888 GLKViewDrawableColorFormat = 0
	// GLKViewDrawableColorFormatSRGBA8888: An sRGBA8888 format.
	GLKViewDrawableColorFormatSRGBA8888 GLKViewDrawableColorFormat = 0
)

func (e GLKViewDrawableColorFormat) String() string {
	switch e {
	case GLKViewDrawableColorFormatRGB565:
		return "GLKViewDrawableColorFormatRGB565"
	default:
		return fmt.Sprintf("GLKViewDrawableColorFormat(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/GLKit/GLKViewDrawableDepthFormat
type GLKViewDrawableDepthFormat int

const (
	// GLKViewDrawableDepthFormat16: A 16-bit depth entry for each pixel.
	GLKViewDrawableDepthFormat16 GLKViewDrawableDepthFormat = 0
	// GLKViewDrawableDepthFormat24: A 24-bit depth entry for each pixel.
	GLKViewDrawableDepthFormat24 GLKViewDrawableDepthFormat = 0
	// GLKViewDrawableDepthFormatNone: The underlying framebuffer object has no depth buffer.
	GLKViewDrawableDepthFormatNone GLKViewDrawableDepthFormat = 0
)

func (e GLKViewDrawableDepthFormat) String() string {
	switch e {
	case GLKViewDrawableDepthFormat16:
		return "GLKViewDrawableDepthFormat16"
	default:
		return fmt.Sprintf("GLKViewDrawableDepthFormat(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/GLKit/GLKViewDrawableMultisample
type GLKViewDrawableMultisample int

const (
	// GLKViewDrawableMultisample4X: Multisampling is enabled.
	GLKViewDrawableMultisample4X GLKViewDrawableMultisample = 0
	// GLKViewDrawableMultisampleNone: Multisampling is not enabled.
	GLKViewDrawableMultisampleNone GLKViewDrawableMultisample = 0
)

func (e GLKViewDrawableMultisample) String() string {
	switch e {
	case GLKViewDrawableMultisample4X:
		return "GLKViewDrawableMultisample4X"
	default:
		return fmt.Sprintf("GLKViewDrawableMultisample(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/GLKit/GLKViewDrawableStencilFormat
type GLKViewDrawableStencilFormat int

const (
	// GLKViewDrawableStencilFormat8: An 8-bit stencil entry for each pixel.
	GLKViewDrawableStencilFormat8 GLKViewDrawableStencilFormat = 0
	// GLKViewDrawableStencilFormatNone: The underlying framebuffer object has no stencil buffer.
	GLKViewDrawableStencilFormatNone GLKViewDrawableStencilFormat = 0
)

func (e GLKViewDrawableStencilFormat) String() string {
	switch e {
	case GLKViewDrawableStencilFormat8:
		return "GLKViewDrawableStencilFormat8"
	default:
		return fmt.Sprintf("GLKViewDrawableStencilFormat(%d)", e)
	}
}
