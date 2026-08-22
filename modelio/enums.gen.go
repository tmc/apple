// Code generated from Apple documentation for ModelIO. DO NOT EDIT.

package modelio

import (
	"fmt"
)

// See: https://developer.apple.com/documentation/ModelIO/MDLAnimatedValueInterpolation
type MDLAnimatedValueInterpolation uint

const (
	MDLAnimatedValueInterpolationConstant MDLAnimatedValueInterpolation = 0
	MDLAnimatedValueInterpolationLinear   MDLAnimatedValueInterpolation = 1
)

func (e MDLAnimatedValueInterpolation) String() string {
	switch e {
	case MDLAnimatedValueInterpolationConstant:
		return "MDLAnimatedValueInterpolationConstant"
	case MDLAnimatedValueInterpolationLinear:
		return "MDLAnimatedValueInterpolationLinear"
	default:
		return fmt.Sprintf("MDLAnimatedValueInterpolation(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/ModelIO/MDLCameraProjection
type MDLCameraProjection uint

const (
	// MDLCameraProjectionOrthographic: An orthographic projection.
	MDLCameraProjectionOrthographic MDLCameraProjection = 1
	// MDLCameraProjectionPerspective: A perspective projection.
	MDLCameraProjectionPerspective MDLCameraProjection = 0
)

func (e MDLCameraProjection) String() string {
	switch e {
	case MDLCameraProjectionOrthographic:
		return "MDLCameraProjectionOrthographic"
	case MDLCameraProjectionPerspective:
		return "MDLCameraProjectionPerspective"
	default:
		return fmt.Sprintf("MDLCameraProjection(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/ModelIO/MDLDataPrecision
type MDLDataPrecision uint

const (
	MDLDataPrecisionDouble    MDLDataPrecision = 2
	MDLDataPrecisionFloat     MDLDataPrecision = 1
	MDLDataPrecisionUndefined MDLDataPrecision = 0
)

func (e MDLDataPrecision) String() string {
	switch e {
	case MDLDataPrecisionDouble:
		return "MDLDataPrecisionDouble"
	case MDLDataPrecisionFloat:
		return "MDLDataPrecisionFloat"
	case MDLDataPrecisionUndefined:
		return "MDLDataPrecisionUndefined"
	default:
		return fmt.Sprintf("MDLDataPrecision(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/ModelIO/MDLGeometryType
type MDLGeometryType int

const (
	// MDLGeometryTypeLines: Each pair of consecutive indices in the submesh refers to two vertices to be rendered as a line segment.
	MDLGeometryTypeLines MDLGeometryType = 1
	// MDLGeometryTypePoints: Each index in the submesh refers to a vertex to be rendered as a single point.
	MDLGeometryTypePoints MDLGeometryType = 0
	// MDLGeometryTypeQuads: Each set of four consecutive indices in the submesh refers to four vertices to be rendered as a quadrilateral.
	MDLGeometryTypeQuads MDLGeometryType = 4
	// MDLGeometryTypeTriangleStrips: The first three consecutive indices in the submesh refer to three vertices to be rendered as a triangle.
	MDLGeometryTypeTriangleStrips MDLGeometryType = 3
	// MDLGeometryTypeTriangles: Each set of three consecutive indices in the submesh refers to three vertices to be rendered as a triangle.
	MDLGeometryTypeTriangles MDLGeometryType = 2
	// MDLGeometryTypeVariableTopology: The submesh’s index buffer does not contain a uniform set of primitives.
	MDLGeometryTypeVariableTopology MDLGeometryType = 5
)

func (e MDLGeometryType) String() string {
	switch e {
	case MDLGeometryTypeLines:
		return "MDLGeometryTypeLines"
	case MDLGeometryTypePoints:
		return "MDLGeometryTypePoints"
	case MDLGeometryTypeQuads:
		return "MDLGeometryTypeQuads"
	case MDLGeometryTypeTriangleStrips:
		return "MDLGeometryTypeTriangleStrips"
	case MDLGeometryTypeTriangles:
		return "MDLGeometryTypeTriangles"
	case MDLGeometryTypeVariableTopology:
		return "MDLGeometryTypeVariableTopology"
	default:
		return fmt.Sprintf("MDLGeometryType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/ModelIO/MDLIndexBitDepth
type MDLIndexBitDepth uint

const (
	// MDLIndexBitDepthInvalid: The submesh has not been initialized or its data type is unknown.
	MDLIndexBitDepthInvalid MDLIndexBitDepth = 0
	// MDLIndexBitDepthUInt16: Each index in the submesh’s index buffer is a 16-bit integer.
	MDLIndexBitDepthUInt16 MDLIndexBitDepth = 16
	// MDLIndexBitDepthUInt32: Each index in the submesh’s index buffer is a 32-bit integer.
	MDLIndexBitDepthUInt32 MDLIndexBitDepth = 32
	// MDLIndexBitDepthUInt8: Each index in the submesh’s index buffer is an 8-bit integer.
	MDLIndexBitDepthUInt8  MDLIndexBitDepth = 8
	MDLIndexBitDepthUint16 MDLIndexBitDepth = 16
	MDLIndexBitDepthUint32 MDLIndexBitDepth = 32
	MDLIndexBitDepthUint8  MDLIndexBitDepth = 8
)

func (e MDLIndexBitDepth) String() string {
	switch e {
	case MDLIndexBitDepthInvalid:
		return "MDLIndexBitDepthInvalid"
	case MDLIndexBitDepthUInt16:
		return "MDLIndexBitDepthUInt16"
	case MDLIndexBitDepthUInt32:
		return "MDLIndexBitDepthUInt32"
	case MDLIndexBitDepthUInt8:
		return "MDLIndexBitDepthUInt8"
	default:
		return fmt.Sprintf("MDLIndexBitDepth(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/ModelIO/MDLLightType
type MDLLightType uint

const (
	// MDLLightTypeAmbient: The light source should illuminate a scene evenly regardless of position or direction.
	MDLLightTypeAmbient MDLLightType = 1
	// MDLLightTypeDirectional: The light source illuminates a scene from a uniform direction regardless of its position.
	MDLLightTypeDirectional MDLLightType = 2
	// MDLLightTypeDiscArea: The light source illuminates a scene in all directions from an area in the shape of a disc.
	MDLLightTypeDiscArea MDLLightType = 6
	// MDLLightTypeEnvironment: The illumination from the light is determined by texture images representing a sample of the surrounding environment for a scene.
	MDLLightTypeEnvironment MDLLightType = 11
	// MDLLightTypeLinear: The light source illuminates a scene in all directions from an area in the shape of a line.
	MDLLightTypeLinear MDLLightType = 5
	// MDLLightTypePhotometric: The illumination from the light is determined by a photometric profile.
	MDLLightTypePhotometric MDLLightType = 9
	// MDLLightTypePoint: The light source illuminates a scene in all directions from a specific position.
	MDLLightTypePoint MDLLightType = 4
	// MDLLightTypeProbe: The illumination from the light is determined by texture images representing a sample of a scene at a specific point.
	MDLLightTypeProbe MDLLightType = 10
	// MDLLightTypeRectangularArea: The light source illuminates a scene in all directions from an area in the shape of a rectangle.
	MDLLightTypeRectangularArea MDLLightType = 7
	// MDLLightTypeSpot: The light source illuminates a scene from a specific position and direction.
	MDLLightTypeSpot MDLLightType = 3
	// MDLLightTypeSuperElliptical: The light source illuminates a scene in all directions from an area in the shape of a superellipse.
	MDLLightTypeSuperElliptical MDLLightType = 8
	// MDLLightTypeUnknown: The type of the light is unknown or has not been initialized.
	MDLLightTypeUnknown MDLLightType = 0
)

func (e MDLLightType) String() string {
	switch e {
	case MDLLightTypeAmbient:
		return "MDLLightTypeAmbient"
	case MDLLightTypeDirectional:
		return "MDLLightTypeDirectional"
	case MDLLightTypeDiscArea:
		return "MDLLightTypeDiscArea"
	case MDLLightTypeEnvironment:
		return "MDLLightTypeEnvironment"
	case MDLLightTypeLinear:
		return "MDLLightTypeLinear"
	case MDLLightTypePhotometric:
		return "MDLLightTypePhotometric"
	case MDLLightTypePoint:
		return "MDLLightTypePoint"
	case MDLLightTypeProbe:
		return "MDLLightTypeProbe"
	case MDLLightTypeRectangularArea:
		return "MDLLightTypeRectangularArea"
	case MDLLightTypeSpot:
		return "MDLLightTypeSpot"
	case MDLLightTypeSuperElliptical:
		return "MDLLightTypeSuperElliptical"
	case MDLLightTypeUnknown:
		return "MDLLightTypeUnknown"
	default:
		return fmt.Sprintf("MDLLightType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/ModelIO/MDLMaterialFace
type MDLMaterialFace uint

const (
	MDLMaterialFaceBack        MDLMaterialFace = 1
	MDLMaterialFaceDoubleSided MDLMaterialFace = 2
	MDLMaterialFaceFront       MDLMaterialFace = 0
)

func (e MDLMaterialFace) String() string {
	switch e {
	case MDLMaterialFaceBack:
		return "MDLMaterialFaceBack"
	case MDLMaterialFaceDoubleSided:
		return "MDLMaterialFaceDoubleSided"
	case MDLMaterialFaceFront:
		return "MDLMaterialFaceFront"
	default:
		return fmt.Sprintf("MDLMaterialFace(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/ModelIO/MDLMaterialMipMapFilterMode
type MDLMaterialMipMapFilterMode uint

const (
	// MDLMaterialMipMapFilterModeLinear: Sampling a texture at a size between mipmap levels should linearly interpolate between mipmap levels.
	MDLMaterialMipMapFilterModeLinear MDLMaterialMipMapFilterMode = 1
	// MDLMaterialMipMapFilterModeNearest: Sampling a texture at a size between mipmap levels should return a texel value from the nearest mipmap level.
	MDLMaterialMipMapFilterModeNearest MDLMaterialMipMapFilterMode = 0
)

func (e MDLMaterialMipMapFilterMode) String() string {
	switch e {
	case MDLMaterialMipMapFilterModeLinear:
		return "MDLMaterialMipMapFilterModeLinear"
	case MDLMaterialMipMapFilterModeNearest:
		return "MDLMaterialMipMapFilterModeNearest"
	default:
		return fmt.Sprintf("MDLMaterialMipMapFilterMode(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/ModelIO/MDLMaterialPropertyType
type MDLMaterialPropertyType uint

const (
	MDLMaterialPropertyTypeBuffer MDLMaterialPropertyType = 10
	// MDLMaterialPropertyTypeColor: The material property’s value is a uniform color.
	MDLMaterialPropertyTypeColor MDLMaterialPropertyType = 4
	// MDLMaterialPropertyTypeFloat: The material property’s value is a floating-point scalar.
	MDLMaterialPropertyTypeFloat MDLMaterialPropertyType = 5
	// MDLMaterialPropertyTypeFloat2: The material property’s value is a 2-component floating-point vector.
	MDLMaterialPropertyTypeFloat2 MDLMaterialPropertyType = 6
	// MDLMaterialPropertyTypeFloat3: The material property’s value is a 3-component floating-point vector.
	MDLMaterialPropertyTypeFloat3 MDLMaterialPropertyType = 7
	// MDLMaterialPropertyTypeFloat4: The material property’s value is a 4-component floating-point vector.
	MDLMaterialPropertyTypeFloat4 MDLMaterialPropertyType = 8
	// MDLMaterialPropertyTypeMatrix44: The material property’s value is a 4 x 4 floating-point matrix.
	MDLMaterialPropertyTypeMatrix44 MDLMaterialPropertyType = 9
	// MDLMaterialPropertyTypeNone: The material property’s value has not been initialized.
	MDLMaterialPropertyTypeNone MDLMaterialPropertyType = 0
	// MDLMaterialPropertyTypeString: The material’s value is a string.
	MDLMaterialPropertyTypeString MDLMaterialPropertyType = 1
	// MDLMaterialPropertyTypeTexture: The material property’s value is a  object that provides both a texture image and texture rendering parameters.
	MDLMaterialPropertyTypeTexture MDLMaterialPropertyType = 3
	// MDLMaterialPropertyTypeURL: The material property’s value is a URL—typically, a URL referencing a texture image.
	MDLMaterialPropertyTypeURL MDLMaterialPropertyType = 2
)

func (e MDLMaterialPropertyType) String() string {
	switch e {
	case MDLMaterialPropertyTypeBuffer:
		return "MDLMaterialPropertyTypeBuffer"
	case MDLMaterialPropertyTypeColor:
		return "MDLMaterialPropertyTypeColor"
	case MDLMaterialPropertyTypeFloat:
		return "MDLMaterialPropertyTypeFloat"
	case MDLMaterialPropertyTypeFloat2:
		return "MDLMaterialPropertyTypeFloat2"
	case MDLMaterialPropertyTypeFloat3:
		return "MDLMaterialPropertyTypeFloat3"
	case MDLMaterialPropertyTypeFloat4:
		return "MDLMaterialPropertyTypeFloat4"
	case MDLMaterialPropertyTypeMatrix44:
		return "MDLMaterialPropertyTypeMatrix44"
	case MDLMaterialPropertyTypeNone:
		return "MDLMaterialPropertyTypeNone"
	case MDLMaterialPropertyTypeString:
		return "MDLMaterialPropertyTypeString"
	case MDLMaterialPropertyTypeTexture:
		return "MDLMaterialPropertyTypeTexture"
	case MDLMaterialPropertyTypeURL:
		return "MDLMaterialPropertyTypeURL"
	default:
		return fmt.Sprintf("MDLMaterialPropertyType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/ModelIO/MDLMaterialSemantic
type MDLMaterialSemantic uint

const (
	// MDLMaterialSemanticAmbientOcclusion: The attenuation of ambient light due to local geometry variations on a surface.
	MDLMaterialSemanticAmbientOcclusion MDLMaterialSemantic = 22
	// MDLMaterialSemanticAmbientOcclusionScale: The scaling factor for ambient occlusion shading.
	MDLMaterialSemanticAmbientOcclusionScale MDLMaterialSemantic = 23
	// MDLMaterialSemanticAnisotropic: The degree to which specular highlights elongate in the direction of the local tangent basis.
	MDLMaterialSemanticAnisotropic MDLMaterialSemantic = 7
	// MDLMaterialSemanticAnisotropicRotation: The angle at which anisotropic effects are rotated relative to the local tangent basis.
	MDLMaterialSemanticAnisotropicRotation MDLMaterialSemantic = 8
	// MDLMaterialSemanticBaseColor: The inherent color of a surface, to be used as a modulator during shading.
	MDLMaterialSemanticBaseColor MDLMaterialSemantic = 0
	// MDLMaterialSemanticBump: The degree of perturbation in a material’s surface.
	MDLMaterialSemanticBump MDLMaterialSemantic = 14
	// MDLMaterialSemanticClearcoat: The intensity of a second specular highlight, similar to the gloss that results from a clear coat on an automotive finish.
	MDLMaterialSemanticClearcoat MDLMaterialSemantic = 11
	// MDLMaterialSemanticClearcoatGloss: The spread of a second specular highlight, similar to the gloss that results from a clear coat on an automotive finish.
	MDLMaterialSemanticClearcoatGloss MDLMaterialSemantic = 12
	// MDLMaterialSemanticDisplacement: The displacement of a material’s surface relative to the surface normal.
	MDLMaterialSemanticDisplacement MDLMaterialSemantic = 20
	// MDLMaterialSemanticDisplacementScale: The scaling factor for displacement of a material’s surface.
	MDLMaterialSemanticDisplacementScale MDLMaterialSemantic = 21
	// MDLMaterialSemanticEmission: The color emitted as radiance from a material’s surface.
	MDLMaterialSemanticEmission MDLMaterialSemantic = 13
	// MDLMaterialSemanticInterfaceIndexOfRefraction: The index of refraction for the medium surrounding a material.
	MDLMaterialSemanticInterfaceIndexOfRefraction MDLMaterialSemantic = 16
	// MDLMaterialSemanticMaterialIndexOfRefraction: The index of refraction for a material itself.
	MDLMaterialSemanticMaterialIndexOfRefraction MDLMaterialSemantic = 17
	// MDLMaterialSemanticMetallic: The degree to which a material appears as a dielectric surface (lower values) or as a metal (higher values).
	MDLMaterialSemanticMetallic MDLMaterialSemantic = 2
	// MDLMaterialSemanticNone: The material property’s  property has not been initialized.
	MDLMaterialSemanticNone MDLMaterialSemantic = 0x8000
	// MDLMaterialSemanticObjectSpaceNormal: The variation in the surface normal vectors in a material, relative to model coordinate space.
	MDLMaterialSemanticObjectSpaceNormal MDLMaterialSemantic = 18
	// MDLMaterialSemanticOpacity: The opacity of a material’s surface.
	MDLMaterialSemanticOpacity MDLMaterialSemantic = 15
	// MDLMaterialSemanticRoughness: The degree to which a material appears smooth, affecting both diffuse and specular response.
	MDLMaterialSemanticRoughness MDLMaterialSemantic = 6
	// MDLMaterialSemanticSheen: The intensity of highlights that appear only at glancing angles on a material’s surface.
	MDLMaterialSemanticSheen MDLMaterialSemantic = 9
	// MDLMaterialSemanticSheenTint: The balance of color for highlights that appear only at glancing angles, between the light color (lower values) and the material’s base color (at higher values).
	MDLMaterialSemanticSheenTint MDLMaterialSemantic = 10
	// MDLMaterialSemanticSpecular: The intensity of specular highlights that appear on the material’s surface.
	MDLMaterialSemanticSpecular MDLMaterialSemantic = 3
	// MDLMaterialSemanticSpecularExponent: The exponent to be used in Blinn-Phong approximation of the material’s specular response.
	MDLMaterialSemanticSpecularExponent MDLMaterialSemantic = 4
	// MDLMaterialSemanticSpecularTint: The balance of color for specular highlights, between the light color (lower values) and the material’s base color (at higher values).
	MDLMaterialSemanticSpecularTint MDLMaterialSemantic = 5
	// MDLMaterialSemanticSubsurface: The degree to which light scatters under the surface of a material.
	MDLMaterialSemanticSubsurface MDLMaterialSemantic = 1
	// MDLMaterialSemanticTangentSpaceNormal: The variation in the surface normal vectors in a material, relative to surface tangent coordinate space.
	MDLMaterialSemanticTangentSpaceNormal MDLMaterialSemantic = 19
	// MDLMaterialSemanticUserDefined: The meaning of the material property’s value is not one of the standard semantic uses recognized by Model I/O.
	MDLMaterialSemanticUserDefined MDLMaterialSemantic = 0x8001
)

func (e MDLMaterialSemantic) String() string {
	switch e {
	case MDLMaterialSemanticAmbientOcclusion:
		return "MDLMaterialSemanticAmbientOcclusion"
	case MDLMaterialSemanticAmbientOcclusionScale:
		return "MDLMaterialSemanticAmbientOcclusionScale"
	case MDLMaterialSemanticAnisotropic:
		return "MDLMaterialSemanticAnisotropic"
	case MDLMaterialSemanticAnisotropicRotation:
		return "MDLMaterialSemanticAnisotropicRotation"
	case MDLMaterialSemanticBaseColor:
		return "MDLMaterialSemanticBaseColor"
	case MDLMaterialSemanticBump:
		return "MDLMaterialSemanticBump"
	case MDLMaterialSemanticClearcoat:
		return "MDLMaterialSemanticClearcoat"
	case MDLMaterialSemanticClearcoatGloss:
		return "MDLMaterialSemanticClearcoatGloss"
	case MDLMaterialSemanticDisplacement:
		return "MDLMaterialSemanticDisplacement"
	case MDLMaterialSemanticDisplacementScale:
		return "MDLMaterialSemanticDisplacementScale"
	case MDLMaterialSemanticEmission:
		return "MDLMaterialSemanticEmission"
	case MDLMaterialSemanticInterfaceIndexOfRefraction:
		return "MDLMaterialSemanticInterfaceIndexOfRefraction"
	case MDLMaterialSemanticMaterialIndexOfRefraction:
		return "MDLMaterialSemanticMaterialIndexOfRefraction"
	case MDLMaterialSemanticMetallic:
		return "MDLMaterialSemanticMetallic"
	case MDLMaterialSemanticNone:
		return "MDLMaterialSemanticNone"
	case MDLMaterialSemanticObjectSpaceNormal:
		return "MDLMaterialSemanticObjectSpaceNormal"
	case MDLMaterialSemanticOpacity:
		return "MDLMaterialSemanticOpacity"
	case MDLMaterialSemanticRoughness:
		return "MDLMaterialSemanticRoughness"
	case MDLMaterialSemanticSheen:
		return "MDLMaterialSemanticSheen"
	case MDLMaterialSemanticSheenTint:
		return "MDLMaterialSemanticSheenTint"
	case MDLMaterialSemanticSpecular:
		return "MDLMaterialSemanticSpecular"
	case MDLMaterialSemanticSpecularExponent:
		return "MDLMaterialSemanticSpecularExponent"
	case MDLMaterialSemanticSpecularTint:
		return "MDLMaterialSemanticSpecularTint"
	case MDLMaterialSemanticSubsurface:
		return "MDLMaterialSemanticSubsurface"
	case MDLMaterialSemanticTangentSpaceNormal:
		return "MDLMaterialSemanticTangentSpaceNormal"
	case MDLMaterialSemanticUserDefined:
		return "MDLMaterialSemanticUserDefined"
	default:
		return fmt.Sprintf("MDLMaterialSemantic(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/ModelIO/MDLMaterialTextureFilterMode
type MDLMaterialTextureFilterMode uint

const (
	// MDLMaterialTextureFilterModeLinear: Sampling at texture coordinates between texels should linearly interpolate between texel values.
	MDLMaterialTextureFilterModeLinear MDLMaterialTextureFilterMode = 1
	// MDLMaterialTextureFilterModeNearest: Sampling at texture coordinates between texels should return the value of the nearest texel.
	MDLMaterialTextureFilterModeNearest MDLMaterialTextureFilterMode = 0
)

func (e MDLMaterialTextureFilterMode) String() string {
	switch e {
	case MDLMaterialTextureFilterModeLinear:
		return "MDLMaterialTextureFilterModeLinear"
	case MDLMaterialTextureFilterModeNearest:
		return "MDLMaterialTextureFilterModeNearest"
	default:
		return fmt.Sprintf("MDLMaterialTextureFilterMode(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/ModelIO/MDLMaterialTextureWrapMode
type MDLMaterialTextureWrapMode uint

const (
	// MDLMaterialTextureWrapModeClamp: Sampling at any texture coordinate outside the  to  range returns the texel color from the nearest edge.
	MDLMaterialTextureWrapModeClamp MDLMaterialTextureWrapMode = 0
	// MDLMaterialTextureWrapModeMirror: Sampling at texture coordinates outside the  to  range results in a mirrored tiling effect.
	MDLMaterialTextureWrapModeMirror MDLMaterialTextureWrapMode = 2
	// MDLMaterialTextureWrapModeRepeat: Sampling at texture coordinates outside the  to  range results in a repeated tiling effect.
	MDLMaterialTextureWrapModeRepeat MDLMaterialTextureWrapMode = 1
)

func (e MDLMaterialTextureWrapMode) String() string {
	switch e {
	case MDLMaterialTextureWrapModeClamp:
		return "MDLMaterialTextureWrapModeClamp"
	case MDLMaterialTextureWrapModeMirror:
		return "MDLMaterialTextureWrapModeMirror"
	case MDLMaterialTextureWrapModeRepeat:
		return "MDLMaterialTextureWrapModeRepeat"
	default:
		return fmt.Sprintf("MDLMaterialTextureWrapMode(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/ModelIO/MDLMeshBufferType
type MDLMeshBufferType uint

const (
	MDLMeshBufferTypeCustom MDLMeshBufferType = 3
	// MDLMeshBufferTypeIndex: The buffer contains index data for a  object.
	MDLMeshBufferTypeIndex MDLMeshBufferType = 2
	// MDLMeshBufferTypeVertex: The buffer contains per-vertex data for one or more vertex attributes of a  object.
	MDLMeshBufferTypeVertex MDLMeshBufferType = 1
)

func (e MDLMeshBufferType) String() string {
	switch e {
	case MDLMeshBufferTypeCustom:
		return "MDLMeshBufferTypeCustom"
	case MDLMeshBufferTypeIndex:
		return "MDLMeshBufferTypeIndex"
	case MDLMeshBufferTypeVertex:
		return "MDLMeshBufferTypeVertex"
	default:
		return fmt.Sprintf("MDLMeshBufferType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/ModelIO/MDLProbePlacement
type MDLProbePlacement int

const (
	// MDLProbePlacementIrradianceDistribution: An option to examine the lighting conditions at various positions in the scene being evaluated, then place light probes only at the locations where each contributes optimally to scene lighting.
	MDLProbePlacementIrradianceDistribution MDLProbePlacement = 1
	// MDLProbePlacementUniformGrid: An option to place light probes at each unit coordinate in a three-dimensional grid that evenly divides the region being evaluated.
	MDLProbePlacementUniformGrid MDLProbePlacement = 0
)

func (e MDLProbePlacement) String() string {
	switch e {
	case MDLProbePlacementIrradianceDistribution:
		return "MDLProbePlacementIrradianceDistribution"
	case MDLProbePlacementUniformGrid:
		return "MDLProbePlacementUniformGrid"
	default:
		return fmt.Sprintf("MDLProbePlacement(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/ModelIO/MDLTextureChannelEncoding
type MDLTextureChannelEncoding int

const (
	// MDLTextureChannelEncodingFloat16: Each channel value per texel is a 16-bit floating-point value.
	MDLTextureChannelEncodingFloat16   MDLTextureChannelEncoding = 0x102
	MDLTextureChannelEncodingFloat16SR MDLTextureChannelEncoding = 0x302
	// MDLTextureChannelEncodingFloat32: Each channel value per texel is a 32-bit floating-point value.
	MDLTextureChannelEncodingFloat32 MDLTextureChannelEncoding = 0x104
	// MDLTextureChannelEncodingUInt16: Each channel value per texel is a 16-bit unsigned integer.
	MDLTextureChannelEncodingUInt16 MDLTextureChannelEncoding = 2
	// MDLTextureChannelEncodingUInt24: Each channel value per texel is a 24-bit unsigned integer.
	MDLTextureChannelEncodingUInt24 MDLTextureChannelEncoding = 3
	// MDLTextureChannelEncodingUInt32: Each channel value per texel is a 32-bit unsigned integer.
	MDLTextureChannelEncodingUInt32 MDLTextureChannelEncoding = 4
	// MDLTextureChannelEncodingUInt8: Each channel value per texel is an 8-bit unsigned integer.
	MDLTextureChannelEncodingUInt8  MDLTextureChannelEncoding = 1
	MDLTextureChannelEncodingUint16 MDLTextureChannelEncoding = 2
	MDLTextureChannelEncodingUint24 MDLTextureChannelEncoding = 3
	MDLTextureChannelEncodingUint32 MDLTextureChannelEncoding = 4
	MDLTextureChannelEncodingUint8  MDLTextureChannelEncoding = 1
)

func (e MDLTextureChannelEncoding) String() string {
	switch e {
	case MDLTextureChannelEncodingFloat16:
		return "MDLTextureChannelEncodingFloat16"
	case MDLTextureChannelEncodingFloat16SR:
		return "MDLTextureChannelEncodingFloat16SR"
	case MDLTextureChannelEncodingFloat32:
		return "MDLTextureChannelEncodingFloat32"
	case MDLTextureChannelEncodingUInt16:
		return "MDLTextureChannelEncodingUInt16"
	case MDLTextureChannelEncodingUInt24:
		return "MDLTextureChannelEncodingUInt24"
	case MDLTextureChannelEncodingUInt32:
		return "MDLTextureChannelEncodingUInt32"
	case MDLTextureChannelEncodingUInt8:
		return "MDLTextureChannelEncodingUInt8"
	default:
		return fmt.Sprintf("MDLTextureChannelEncoding(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/ModelIO/MDLTransformOpRotationOrder
type MDLTransformOpRotationOrder uint

const (
	MDLTransformOpRotationOrderXYZ MDLTransformOpRotationOrder = 1
	MDLTransformOpRotationOrderXZY MDLTransformOpRotationOrder = 2
	MDLTransformOpRotationOrderYXZ MDLTransformOpRotationOrder = 3
	MDLTransformOpRotationOrderYZX MDLTransformOpRotationOrder = 4
	MDLTransformOpRotationOrderZXY MDLTransformOpRotationOrder = 5
	MDLTransformOpRotationOrderZYX MDLTransformOpRotationOrder = 6
)

func (e MDLTransformOpRotationOrder) String() string {
	switch e {
	case MDLTransformOpRotationOrderXYZ:
		return "MDLTransformOpRotationOrderXYZ"
	case MDLTransformOpRotationOrderXZY:
		return "MDLTransformOpRotationOrderXZY"
	case MDLTransformOpRotationOrderYXZ:
		return "MDLTransformOpRotationOrderYXZ"
	case MDLTransformOpRotationOrderYZX:
		return "MDLTransformOpRotationOrderYZX"
	case MDLTransformOpRotationOrderZXY:
		return "MDLTransformOpRotationOrderZXY"
	case MDLTransformOpRotationOrderZYX:
		return "MDLTransformOpRotationOrderZYX"
	default:
		return fmt.Sprintf("MDLTransformOpRotationOrder(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/ModelIO/MDLVertexFormat
type MDLVertexFormat uint

const (
	// MDLVertexFormatChar: The attribute value for each vertex is a scalar of signed 8-bit integer type.
	MDLVertexFormatChar MDLVertexFormat = 131073
	// MDLVertexFormatChar2: The attribute value for each vertex is a vector with 2 components, each of signed 8-bit integer type.
	MDLVertexFormatChar2 MDLVertexFormat = 131074
	// MDLVertexFormatChar2Normalized: The attribute value for each vertex is a vector with 2 components, each with a normalized value of signed 8-bit integer type.
	MDLVertexFormatChar2Normalized MDLVertexFormat = 262146
	// MDLVertexFormatChar3: The attribute value for each vertex is a vector with 3 components, each of signed 8-bit integer type.
	MDLVertexFormatChar3 MDLVertexFormat = 131075
	// MDLVertexFormatChar3Normalized: The attribute value for each vertex is a vector with 3 components, each with a normalized value of signed 8-bit integer type.
	MDLVertexFormatChar3Normalized MDLVertexFormat = 262147
	// MDLVertexFormatChar4: The attribute value for each vertex is a vector with 4 components, each of signed 8-bit integer type.
	MDLVertexFormatChar4 MDLVertexFormat = 131076
	// MDLVertexFormatChar4Normalized: The attribute value for each vertex is a vector with 4 components, each with a normalized value of signed 8-bit integer type.
	MDLVertexFormatChar4Normalized MDLVertexFormat = 262148
	// MDLVertexFormatCharBits: A bit mask for vertex attributes whose components are in 8-bit signed integer format.
	MDLVertexFormatCharBits MDLVertexFormat = 0x20000
	// MDLVertexFormatCharNormalized: The attribute value for each vertex is a normalized scalar of signed 8-bit integer type.
	MDLVertexFormatCharNormalized MDLVertexFormat = 262145
	// MDLVertexFormatCharNormalizedBits: A bit mask for vertex attributes whose components are in 8-bit signed normalized integer format.
	MDLVertexFormatCharNormalizedBits MDLVertexFormat = 0x40000
	// MDLVertexFormatFloat: The attribute value for each vertex is a scalar of 32-bit floating-point type.
	MDLVertexFormatFloat MDLVertexFormat = 786433
	// MDLVertexFormatFloat2: The attribute value for each vertex is a vector with 2 components, each of 32-bit floating-point type.
	MDLVertexFormatFloat2 MDLVertexFormat = 786434
	// MDLVertexFormatFloat3: The attribute value for each vertex is a vector with 3 components, each of 32-bit floating-point type.
	MDLVertexFormatFloat3 MDLVertexFormat = 786435
	// MDLVertexFormatFloat4: The attribute value for each vertex is a vector with 4 components, each of 32-bit floating-point type.
	MDLVertexFormatFloat4 MDLVertexFormat = 786436
	// MDLVertexFormatFloatBits: A bit mask for vertex attributes whose components are in 32-bit floating-point format.
	MDLVertexFormatFloatBits MDLVertexFormat = 0xc0000
	// MDLVertexFormatHalf: The attribute value for each vertex is a scalar of 16-bit floating-point type.
	MDLVertexFormatHalf MDLVertexFormat = 720897
	// MDLVertexFormatHalf2: The attribute value for each vertex is a vector with 2 components, each of 16-bit floating-point type.
	MDLVertexFormatHalf2 MDLVertexFormat = 720898
	// MDLVertexFormatHalf3: The attribute value for each vertex is a vector with 3 components, each of 16-bit floating-point type.
	MDLVertexFormatHalf3 MDLVertexFormat = 720899
	// MDLVertexFormatHalf4: The attribute value for each vertex is a vector with 4 components, each of 16-bit floating-point type.
	MDLVertexFormatHalf4 MDLVertexFormat = 720900
	// MDLVertexFormatHalfBits: A bit mask for vertex attributes whose components are in 16-bit floating-point format.
	MDLVertexFormatHalfBits MDLVertexFormat = 0xb0000
	// MDLVertexFormatInt: The attribute value for each vertex is a scalar of signed 32-bit integer type.
	MDLVertexFormatInt MDLVertexFormat = 655361
	// MDLVertexFormatInt1010102Normalized: The attribute value for each vertex is a packed vector with 4 components of signed integer type.
	MDLVertexFormatInt1010102Normalized MDLVertexFormat = 659460
	// MDLVertexFormatInt2: The attribute value for each vertex is a vector with 2 components, each of signed 32-bit integer type.
	MDLVertexFormatInt2 MDLVertexFormat = 655362
	// MDLVertexFormatInt3: The attribute value for each vertex is a vector with 3 components, each of signed 32-bit integer type.
	MDLVertexFormatInt3 MDLVertexFormat = 655363
	// MDLVertexFormatInt4: The attribute value for each vertex is a vector with 4 components, each of signed 32-bit integer type.
	MDLVertexFormatInt4 MDLVertexFormat = 655364
	// MDLVertexFormatIntBits: A bit mask for vertex attributes whose components are in 32-bit signed integer format.
	MDLVertexFormatIntBits MDLVertexFormat = 0xa0000
	// MDLVertexFormatInvalid: The vertex attribute has just been initialized or its format is unknown.
	MDLVertexFormatInvalid MDLVertexFormat = 0
	// MDLVertexFormatPackedBit: A bit mask for vertex attributes in packed vector formats.
	MDLVertexFormatPackedBit MDLVertexFormat = 0x1000
	// MDLVertexFormatShort: The attribute value for each vertex is a scalar of signed 16-bit integer type.
	MDLVertexFormatShort MDLVertexFormat = 393217
	// MDLVertexFormatShort2: The attribute value for each vertex is a vector with 2 components, each of signed 16-bit integer type.
	MDLVertexFormatShort2 MDLVertexFormat = 393218
	// MDLVertexFormatShort2Normalized: The attribute value for each vertex is a vector with 2 components, each with a normalized value of signed 16-bit integer type.
	MDLVertexFormatShort2Normalized MDLVertexFormat = 524290
	// MDLVertexFormatShort3: The attribute value for each vertex is a vector with 3 components, each of signed 16-bit integer type.
	MDLVertexFormatShort3 MDLVertexFormat = 393219
	// MDLVertexFormatShort3Normalized: The attribute value for each vertex is a vector with 3 components, each with a normalized value of signed 16-bit integer type.
	MDLVertexFormatShort3Normalized MDLVertexFormat = 524291
	// MDLVertexFormatShort4: The attribute value for each vertex is a vector with 4 components, each of signed 16-bit integer type.
	MDLVertexFormatShort4 MDLVertexFormat = 393220
	// MDLVertexFormatShort4Normalized: The attribute value for each vertex is a vector with 4 components, each with a normalized value of signed 16-bit integer type.
	MDLVertexFormatShort4Normalized MDLVertexFormat = 524292
	// MDLVertexFormatShortBits: A bit mask for vertex attributes whose components are in 16-bit signed integer format.
	MDLVertexFormatShortBits MDLVertexFormat = 0x60000
	// MDLVertexFormatShortNormalized: The attribute value for each vertex is a normalized scalar of signed 16-bit integer type.
	MDLVertexFormatShortNormalized MDLVertexFormat = 524289
	// MDLVertexFormatShortNormalizedBits: A bit mask for vertex attributes whose components are in 16-bit signed normalized integer format.
	MDLVertexFormatShortNormalizedBits MDLVertexFormat = 0x80000
	// MDLVertexFormatUChar: The attribute value for each vertex is a scalar of unsigned 8-bit integer type.
	MDLVertexFormatUChar MDLVertexFormat = 65537
	// MDLVertexFormatUChar2: The attribute value for each vertex is a vector with 2 components, each of unsigned 8-bit integer type.
	MDLVertexFormatUChar2 MDLVertexFormat = 65538
	// MDLVertexFormatUChar2Normalized: The attribute value for each vertex is a vector with 2 components, each with a normalized value of unsigned 8-bit integer type.
	MDLVertexFormatUChar2Normalized MDLVertexFormat = 196610
	// MDLVertexFormatUChar3: The attribute value for each vertex is a vector with 3 components, each of unsigned 8-bit integer type.
	MDLVertexFormatUChar3 MDLVertexFormat = 65539
	// MDLVertexFormatUChar3Normalized: The attribute value for each vertex is a vector with 3 components, each with a normalized value of unsigned 8-bit integer type.
	MDLVertexFormatUChar3Normalized MDLVertexFormat = 196611
	// MDLVertexFormatUChar4: The attribute value for each vertex is a vector with 4 components, each of unsigned 8-bit integer type.
	MDLVertexFormatUChar4 MDLVertexFormat = 65540
	// MDLVertexFormatUChar4Normalized: The attribute value for each vertex is a vector with 4 components, each with a normalized value of unsigned 8-bit integer type.
	MDLVertexFormatUChar4Normalized MDLVertexFormat = 196612
	// MDLVertexFormatUCharBits: A bit mask for vertex attributes whose components are in 8-bit unsigned integer format.
	MDLVertexFormatUCharBits MDLVertexFormat = 0x10000
	// MDLVertexFormatUCharNormalized: The attribute value for each vertex is a normalized scalar of unsigned 8-bit integer type.
	MDLVertexFormatUCharNormalized MDLVertexFormat = 196609
	// MDLVertexFormatUCharNormalizedBits: A bit mask for vertex attributes whose components are in 8-bit unsigned normalized integer format.
	MDLVertexFormatUCharNormalizedBits MDLVertexFormat = 0x30000
	// MDLVertexFormatUInt: The attribute value for each vertex is a scalar of unsigned 32-bit integer type.
	MDLVertexFormatUInt MDLVertexFormat = 589825
	// MDLVertexFormatUInt1010102Normalized: The attribute value for each vertex is a packed vector with 4 components of unsigned integer type.
	MDLVertexFormatUInt1010102Normalized MDLVertexFormat = 593924
	// MDLVertexFormatUInt2: The attribute value for each vertex is a vector with 2 components, each of unsigned 32-bit integer type.
	MDLVertexFormatUInt2 MDLVertexFormat = 589826
	// MDLVertexFormatUInt3: The attribute value for each vertex is a vector with 3 components, each of unsigned 32-bit integer type.
	MDLVertexFormatUInt3 MDLVertexFormat = 589827
	// MDLVertexFormatUInt4: The attribute value for each vertex is a vector with 4 components, each of unsigned 32-bit integer type.
	MDLVertexFormatUInt4 MDLVertexFormat = 589828
	// MDLVertexFormatUIntBits: A bit mask for vertex attributes whose components are in 32-bit unsigned integer format.
	MDLVertexFormatUIntBits MDLVertexFormat = 0x90000
	// MDLVertexFormatUShort: The attribute value for each vertex is a scalar of unsigned 16-bit integer type.
	MDLVertexFormatUShort MDLVertexFormat = 327681
	// MDLVertexFormatUShort2: The attribute value for each vertex is a vector with 2 components, each of unsigned 16-bit integer type.
	MDLVertexFormatUShort2 MDLVertexFormat = 327682
	// MDLVertexFormatUShort2Normalized: The attribute value for each vertex is a vector with 2 components, each with a normalized value of unsigned 16-bit integer type.
	MDLVertexFormatUShort2Normalized MDLVertexFormat = 458754
	// MDLVertexFormatUShort3: The attribute value for each vertex is a vector with 3 components, each of unsigned 16-bit integer type.
	MDLVertexFormatUShort3 MDLVertexFormat = 327683
	// MDLVertexFormatUShort3Normalized: The attribute value for each vertex is a vector with 3 components, each with a normalized value of unsigned 16-bit integer type.
	MDLVertexFormatUShort3Normalized MDLVertexFormat = 458755
	// MDLVertexFormatUShort4: The attribute value for each vertex is a vector with 4 components, each of unsigned 16-bit integer type.
	MDLVertexFormatUShort4 MDLVertexFormat = 327684
	// MDLVertexFormatUShort4Normalized: The attribute value for each vertex is a vector with 4 components, each with a normalized value of unsigned 16-bit integer type.
	MDLVertexFormatUShort4Normalized MDLVertexFormat = 458756
	// MDLVertexFormatUShortBits: A bit mask for vertex attributes whose components are in 16-bit unsigned integer format.
	MDLVertexFormatUShortBits MDLVertexFormat = 0x50000
	// MDLVertexFormatUShortNormalized: The attribute value for each vertex is a normalized scalar of unsigned 16-bit integer type.
	MDLVertexFormatUShortNormalized MDLVertexFormat = 458753
	// MDLVertexFormatUShortNormalizedBits: A bit mask for vertex attributes whose components are in 16-bit unsigned normalized integer format.
	MDLVertexFormatUShortNormalizedBits MDLVertexFormat = 0x70000
)

func (e MDLVertexFormat) String() string {
	switch e {
	case MDLVertexFormatChar:
		return "MDLVertexFormatChar"
	case MDLVertexFormatChar2:
		return "MDLVertexFormatChar2"
	case MDLVertexFormatChar2Normalized:
		return "MDLVertexFormatChar2Normalized"
	case MDLVertexFormatChar3:
		return "MDLVertexFormatChar3"
	case MDLVertexFormatChar3Normalized:
		return "MDLVertexFormatChar3Normalized"
	case MDLVertexFormatChar4:
		return "MDLVertexFormatChar4"
	case MDLVertexFormatChar4Normalized:
		return "MDLVertexFormatChar4Normalized"
	case MDLVertexFormatCharBits:
		return "MDLVertexFormatCharBits"
	case MDLVertexFormatCharNormalized:
		return "MDLVertexFormatCharNormalized"
	case MDLVertexFormatCharNormalizedBits:
		return "MDLVertexFormatCharNormalizedBits"
	case MDLVertexFormatFloat:
		return "MDLVertexFormatFloat"
	case MDLVertexFormatFloat2:
		return "MDLVertexFormatFloat2"
	case MDLVertexFormatFloat3:
		return "MDLVertexFormatFloat3"
	case MDLVertexFormatFloat4:
		return "MDLVertexFormatFloat4"
	case MDLVertexFormatFloatBits:
		return "MDLVertexFormatFloatBits"
	case MDLVertexFormatHalf:
		return "MDLVertexFormatHalf"
	case MDLVertexFormatHalf2:
		return "MDLVertexFormatHalf2"
	case MDLVertexFormatHalf3:
		return "MDLVertexFormatHalf3"
	case MDLVertexFormatHalf4:
		return "MDLVertexFormatHalf4"
	case MDLVertexFormatHalfBits:
		return "MDLVertexFormatHalfBits"
	case MDLVertexFormatInt:
		return "MDLVertexFormatInt"
	case MDLVertexFormatInt1010102Normalized:
		return "MDLVertexFormatInt1010102Normalized"
	case MDLVertexFormatInt2:
		return "MDLVertexFormatInt2"
	case MDLVertexFormatInt3:
		return "MDLVertexFormatInt3"
	case MDLVertexFormatInt4:
		return "MDLVertexFormatInt4"
	case MDLVertexFormatIntBits:
		return "MDLVertexFormatIntBits"
	case MDLVertexFormatInvalid:
		return "MDLVertexFormatInvalid"
	case MDLVertexFormatPackedBit:
		return "MDLVertexFormatPackedBit"
	case MDLVertexFormatShort:
		return "MDLVertexFormatShort"
	case MDLVertexFormatShort2:
		return "MDLVertexFormatShort2"
	case MDLVertexFormatShort2Normalized:
		return "MDLVertexFormatShort2Normalized"
	case MDLVertexFormatShort3:
		return "MDLVertexFormatShort3"
	case MDLVertexFormatShort3Normalized:
		return "MDLVertexFormatShort3Normalized"
	case MDLVertexFormatShort4:
		return "MDLVertexFormatShort4"
	case MDLVertexFormatShort4Normalized:
		return "MDLVertexFormatShort4Normalized"
	case MDLVertexFormatShortBits:
		return "MDLVertexFormatShortBits"
	case MDLVertexFormatShortNormalized:
		return "MDLVertexFormatShortNormalized"
	case MDLVertexFormatShortNormalizedBits:
		return "MDLVertexFormatShortNormalizedBits"
	case MDLVertexFormatUChar:
		return "MDLVertexFormatUChar"
	case MDLVertexFormatUChar2:
		return "MDLVertexFormatUChar2"
	case MDLVertexFormatUChar2Normalized:
		return "MDLVertexFormatUChar2Normalized"
	case MDLVertexFormatUChar3:
		return "MDLVertexFormatUChar3"
	case MDLVertexFormatUChar3Normalized:
		return "MDLVertexFormatUChar3Normalized"
	case MDLVertexFormatUChar4:
		return "MDLVertexFormatUChar4"
	case MDLVertexFormatUChar4Normalized:
		return "MDLVertexFormatUChar4Normalized"
	case MDLVertexFormatUCharBits:
		return "MDLVertexFormatUCharBits"
	case MDLVertexFormatUCharNormalized:
		return "MDLVertexFormatUCharNormalized"
	case MDLVertexFormatUCharNormalizedBits:
		return "MDLVertexFormatUCharNormalizedBits"
	case MDLVertexFormatUInt:
		return "MDLVertexFormatUInt"
	case MDLVertexFormatUInt1010102Normalized:
		return "MDLVertexFormatUInt1010102Normalized"
	case MDLVertexFormatUInt2:
		return "MDLVertexFormatUInt2"
	case MDLVertexFormatUInt3:
		return "MDLVertexFormatUInt3"
	case MDLVertexFormatUInt4:
		return "MDLVertexFormatUInt4"
	case MDLVertexFormatUIntBits:
		return "MDLVertexFormatUIntBits"
	case MDLVertexFormatUShort:
		return "MDLVertexFormatUShort"
	case MDLVertexFormatUShort2:
		return "MDLVertexFormatUShort2"
	case MDLVertexFormatUShort2Normalized:
		return "MDLVertexFormatUShort2Normalized"
	case MDLVertexFormatUShort3:
		return "MDLVertexFormatUShort3"
	case MDLVertexFormatUShort3Normalized:
		return "MDLVertexFormatUShort3Normalized"
	case MDLVertexFormatUShort4:
		return "MDLVertexFormatUShort4"
	case MDLVertexFormatUShort4Normalized:
		return "MDLVertexFormatUShort4Normalized"
	case MDLVertexFormatUShortBits:
		return "MDLVertexFormatUShortBits"
	case MDLVertexFormatUShortNormalized:
		return "MDLVertexFormatUShortNormalized"
	case MDLVertexFormatUShortNormalizedBits:
		return "MDLVertexFormatUShortNormalizedBits"
	default:
		return fmt.Sprintf("MDLVertexFormat(%d)", e)
	}
}
