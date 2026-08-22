// Code generated from Apple documentation. DO NOT EDIT.

package modelio

import (
	"github.com/ebitengine/purego"
	"github.com/tmc/apple/objc"
)

var (
	// MDLVertexAttributeAnisotropy is the attribute data describes the degree to which a surface’s appearance changes in appearance when rotated about its normal vector.
	//
	// See: https://developer.apple.com/documentation/ModelIO/MDLVertexAttributeAnisotropy
	MDLVertexAttributeAnisotropy string
	// MDLVertexAttributeBinormal is the attribute data describes surface binormal vectors.
	//
	// See: https://developer.apple.com/documentation/ModelIO/MDLVertexAttributeBinormal
	MDLVertexAttributeBinormal string
	// MDLVertexAttributeBitangent is the attribute data describes surface bitangent vectors.
	//
	// See: https://developer.apple.com/documentation/ModelIO/MDLVertexAttributeBitangent
	MDLVertexAttributeBitangent string
	// MDLVertexAttributeColor is the attribute data describes vertex colors.
	//
	// See: https://developer.apple.com/documentation/ModelIO/MDLVertexAttributeColor
	MDLVertexAttributeColor string
	// MDLVertexAttributeEdgeCrease is the attribute data describes edges that should be left unmodified by surface subdivision operations.
	//
	// See: https://developer.apple.com/documentation/ModelIO/MDLVertexAttributeEdgeCrease
	MDLVertexAttributeEdgeCrease string
	// MDLVertexAttributeJointIndices is the attribute data describes the indices of bones or joints in a skeletal animation rig.
	//
	// See: https://developer.apple.com/documentation/ModelIO/MDLVertexAttributeJointIndices
	MDLVertexAttributeJointIndices string
	// MDLVertexAttributeJointWeights is the attribute data describes the influence factors of bones or joints on a vertex’s position for use in skeletal animation.
	//
	// See: https://developer.apple.com/documentation/ModelIO/MDLVertexAttributeJointWeights
	MDLVertexAttributeJointWeights string
	// MDLVertexAttributeNormal is the attribute data describes surface normal vectors.
	//
	// See: https://developer.apple.com/documentation/ModelIO/MDLVertexAttributeNormal
	MDLVertexAttributeNormal string
	// MDLVertexAttributeOcclusionValue is the attribute data describes per-vertex ambient occlusion values.
	//
	// See: https://developer.apple.com/documentation/ModelIO/MDLVertexAttributeOcclusionValue
	MDLVertexAttributeOcclusionValue string
	// MDLVertexAttributePosition is the attribute data describes vertex positions.
	//
	// See: https://developer.apple.com/documentation/ModelIO/MDLVertexAttributePosition
	MDLVertexAttributePosition string
	// MDLVertexAttributeShadingBasisU is the attribute data describes the U component of a vector basis for use in shading.
	//
	// See: https://developer.apple.com/documentation/ModelIO/MDLVertexAttributeShadingBasisU
	MDLVertexAttributeShadingBasisU string
	// MDLVertexAttributeShadingBasisV is the attribute data describes the V component of a vector basis for use in shading.
	//
	// See: https://developer.apple.com/documentation/ModelIO/MDLVertexAttributeShadingBasisV
	MDLVertexAttributeShadingBasisV string
	// MDLVertexAttributeSubdivisionStencil is the attribute data describes which neighboring vertices influence the effect of surface subdivision on the area around a vertex.
	//
	// See: https://developer.apple.com/documentation/ModelIO/MDLVertexAttributeSubdivisionStencil
	MDLVertexAttributeSubdivisionStencil string
	// MDLVertexAttributeTangent is the attribute data describes surface tangent vectors.
	//
	// See: https://developer.apple.com/documentation/ModelIO/MDLVertexAttributeTangent
	MDLVertexAttributeTangent string
	// MDLVertexAttributeTextureCoordinate is the attribute data describes texture coordinates.
	//
	// See: https://developer.apple.com/documentation/ModelIO/MDLVertexAttributeTextureCoordinate
	MDLVertexAttributeTextureCoordinate string
	// KUTType3dObject is the Object file format (common extension `XCUIElementTypeObj`).
	//
	// See: https://developer.apple.com/documentation/ModelIO/kUTType3dObject
	KUTType3dObject string
	// KUTTypeAlembic is the Alembic file format (common extension `XCUIElementTypeAbc`).
	//
	// See: https://developer.apple.com/documentation/ModelIO/kUTTypeAlembic
	KUTTypeAlembic string
	// KUTTypePolygon is the Polygon file format (common extension `XCUIElementTypePly`).
	//
	// See: https://developer.apple.com/documentation/ModelIO/kUTTypePolygon
	KUTTypePolygon string
	// KUTTypeStereolithography is the Stereolithography file format (common extension `XCUIElementTypeStl`), also known as the Standard Tessellated Geometry format.
	//
	// See: https://developer.apple.com/documentation/ModelIO/kUTTypeStereolithography
	KUTTypeStereolithography string
	// KUTTypeUniversalSceneDescription is the Universal Scene Description file format (common extension `XCUIElementTypeUsd`).
	//
	// See: https://developer.apple.com/documentation/ModelIO/kUTTypeUniversalSceneDescription
	KUTTypeUniversalSceneDescription string
	// See: https://developer.apple.com/documentation/ModelIO/kUTTypeUniversalSceneDescriptionMobile
	KUTTypeUniversalSceneDescriptionMobile string
)

func init() {
	if frameworkHandle == 0 {
		return
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "MDLVertexAttributeAnisotropy"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MDLVertexAttributeAnisotropy = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "MDLVertexAttributeBinormal"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MDLVertexAttributeBinormal = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "MDLVertexAttributeBitangent"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MDLVertexAttributeBitangent = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "MDLVertexAttributeColor"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MDLVertexAttributeColor = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "MDLVertexAttributeEdgeCrease"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MDLVertexAttributeEdgeCrease = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "MDLVertexAttributeJointIndices"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MDLVertexAttributeJointIndices = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "MDLVertexAttributeJointWeights"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MDLVertexAttributeJointWeights = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "MDLVertexAttributeNormal"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MDLVertexAttributeNormal = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "MDLVertexAttributeOcclusionValue"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MDLVertexAttributeOcclusionValue = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "MDLVertexAttributePosition"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MDLVertexAttributePosition = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "MDLVertexAttributeShadingBasisU"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MDLVertexAttributeShadingBasisU = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "MDLVertexAttributeShadingBasisV"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MDLVertexAttributeShadingBasisV = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "MDLVertexAttributeSubdivisionStencil"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MDLVertexAttributeSubdivisionStencil = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "MDLVertexAttributeTangent"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MDLVertexAttributeTangent = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "MDLVertexAttributeTextureCoordinate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				MDLVertexAttributeTextureCoordinate = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kUTType3dObject"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KUTType3dObject = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kUTTypeAlembic"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KUTTypeAlembic = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kUTTypePolygon"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KUTTypePolygon = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kUTTypeStereolithography"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KUTTypeStereolithography = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kUTTypeUniversalSceneDescription"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KUTTypeUniversalSceneDescription = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kUTTypeUniversalSceneDescriptionMobile"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KUTTypeUniversalSceneDescriptionMobile = objc.GoString(cstr)
			}
		}
	}

}
