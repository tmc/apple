// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"errors"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLModelIOUtils] class.
var (
	_MLModelIOUtilsClass     MLModelIOUtilsClass
	_MLModelIOUtilsClassOnce sync.Once
)

func getMLModelIOUtilsClass() MLModelIOUtilsClass {
	_MLModelIOUtilsClassOnce.Do(func() {
		_MLModelIOUtilsClass = MLModelIOUtilsClass{class: objc.GetClass("MLModelIOUtils")}
	})
	return _MLModelIOUtilsClass
}

// GetMLModelIOUtilsClass returns the class object for MLModelIOUtils.
func GetMLModelIOUtilsClass() MLModelIOUtilsClass {
	return getMLModelIOUtilsClass()
}

type MLModelIOUtilsClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLModelIOUtilsClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLModelIOUtilsClass) Alloc() MLModelIOUtils {
	rv := objc.Send[MLModelIOUtils](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLModelIOUtils
type MLModelIOUtils struct {
	objectivec.Object
}

// MLModelIOUtilsFromID constructs a [MLModelIOUtils] from an objc.ID.
func MLModelIOUtilsFromID(id objc.ID) MLModelIOUtils {
	return MLModelIOUtils{objectivec.Object{ID: id}}
}
// Ensure MLModelIOUtils implements IMLModelIOUtils.
var _ IMLModelIOUtils = MLModelIOUtils{}

// An interface definition for the [MLModelIOUtils] class.
//
// See: https://developer.apple.com/documentation/CoreML/MLModelIOUtils
type IMLModelIOUtils interface {
	objectivec.IObject
}

// Init initializes the instance.
func (m MLModelIOUtils) Init() MLModelIOUtils {
	rv := objc.Send[MLModelIOUtils](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLModelIOUtils) Autorelease() MLModelIOUtils {
	rv := objc.Send[MLModelIOUtils](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLModelIOUtils creates a new MLModelIOUtils instance.
func NewMLModelIOUtils() MLModelIOUtils {
	class := getMLModelIOUtilsClass()
	rv := objc.Send[MLModelIOUtils](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/CoreML/MLModelIOUtils/defaultFunctionNameFromDescriptionSpecification:
func (_MLModelIOUtilsClass MLModelIOUtilsClass) DefaultFunctionNameFromDescriptionSpecification(specification unsafe.Pointer) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLModelIOUtilsClass.class), objc.Sel("defaultFunctionNameFromDescriptionSpecification:"), specification)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/CoreML/MLModelIOUtils/descriptionFromProto:
func (_MLModelIOUtilsClass MLModelIOUtilsClass) DescriptionFromProto(proto unsafe.Pointer) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLModelIOUtilsClass.class), objc.Sel("descriptionFromProto:"), proto)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/CoreML/MLModelIOUtils/deserializeInterfaceFormat:archive:error:
func (_MLModelIOUtilsClass MLModelIOUtilsClass) DeserializeInterfaceFormatArchiveError(format unsafe.Pointer, archive unsafe.Pointer) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](objc.ID(_MLModelIOUtilsClass.class), objc.Sel("deserializeInterfaceFormat:archive:error:"), format, archive, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("deserializeInterfaceFormat:archive:error: returned NO with nil NSError")
	}
	return rv, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLModelIOUtils/deserializeMetadataAndInterfaceFromArchive:error:
func (_MLModelIOUtilsClass MLModelIOUtilsClass) DeserializeMetadataAndInterfaceFromArchiveError(archive unsafe.Pointer) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLModelIOUtilsClass.class), objc.Sel("deserializeMetadataAndInterfaceFromArchive:error:"), archive, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLModelIOUtils/deserializeVersionInfoFromArchive:error:
func (_MLModelIOUtilsClass MLModelIOUtilsClass) DeserializeVersionInfoFromArchiveError(archive unsafe.Pointer) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLModelIOUtilsClass.class), objc.Sel("deserializeVersionInfoFromArchive:error:"), archive, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLModelIOUtils/featureDescriptionsFromProto:
func (_MLModelIOUtilsClass MLModelIOUtilsClass) FeatureDescriptionsFromProto(proto unsafe.Pointer) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLModelIOUtilsClass.class), objc.Sel("featureDescriptionsFromProto:"), proto)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/CoreML/MLModelIOUtils/functionDescriptionsFromDescriptionSpecification:
func (_MLModelIOUtilsClass MLModelIOUtilsClass) FunctionDescriptionsFromDescriptionSpecification(specification unsafe.Pointer) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLModelIOUtilsClass.class), objc.Sel("functionDescriptionsFromDescriptionSpecification:"), specification)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/CoreML/MLModelIOUtils/inputDescriptionFromInterface:
func (_MLModelIOUtilsClass MLModelIOUtilsClass) InputDescriptionFromInterface(interface_ unsafe.Pointer) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLModelIOUtilsClass.class), objc.Sel("inputDescriptionFromInterface:"), interface_)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/CoreML/MLModelIOUtils/loadFromModelSpecificationInArchive:withClass:versionInfo:configuration:error:
func (_MLModelIOUtilsClass MLModelIOUtilsClass) LoadFromModelSpecificationInArchiveWithClassVersionInfoConfigurationError(archive unsafe.Pointer, class objc.Class, info objectivec.IObject, configuration objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLModelIOUtilsClass.class), objc.Sel("loadFromModelSpecificationInArchive:withClass:versionInfo:configuration:error:"), archive, class, info, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLModelIOUtils/orderedFeatureNamesFromInterface:forInput:
func (_MLModelIOUtilsClass MLModelIOUtilsClass) OrderedFeatureNamesFromInterfaceForInput(interface_ unsafe.Pointer, input bool) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLModelIOUtilsClass.class), objc.Sel("orderedFeatureNamesFromInterface:forInput:"), interface_, input)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/CoreML/MLModelIOUtils/orderedInputFeatureNamesFromInterface:
func (_MLModelIOUtilsClass MLModelIOUtilsClass) OrderedInputFeatureNamesFromInterface(interface_ unsafe.Pointer) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLModelIOUtilsClass.class), objc.Sel("orderedInputFeatureNamesFromInterface:"), interface_)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/CoreML/MLModelIOUtils/orderedNamesFromProto:
func (_MLModelIOUtilsClass MLModelIOUtilsClass) OrderedNamesFromProto(proto unsafe.Pointer) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLModelIOUtilsClass.class), objc.Sel("orderedNamesFromProto:"), proto)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/CoreML/MLModelIOUtils/orderedOutputFeatureNamesFromInterface:
func (_MLModelIOUtilsClass MLModelIOUtilsClass) OrderedOutputFeatureNamesFromInterface(interface_ unsafe.Pointer) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLModelIOUtilsClass.class), objc.Sel("orderedOutputFeatureNamesFromInterface:"), interface_)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/CoreML/MLModelIOUtils/orderedStateFeatureNamesFromInterface:
func (_MLModelIOUtilsClass MLModelIOUtilsClass) OrderedStateFeatureNamesFromInterface(interface_ unsafe.Pointer) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLModelIOUtilsClass.class), objc.Sel("orderedStateFeatureNamesFromInterface:"), interface_)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/CoreML/MLModelIOUtils/outputDescriptionFromInterface:
func (_MLModelIOUtilsClass MLModelIOUtilsClass) OutputDescriptionFromInterface(interface_ unsafe.Pointer) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLModelIOUtilsClass.class), objc.Sel("outputDescriptionFromInterface:"), interface_)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/CoreML/MLModelIOUtils/populateConstraintsForArrayFeatureType:dataType:constraintClass:defaultOptionalValue:
func (_MLModelIOUtilsClass MLModelIOUtilsClass) PopulateConstraintsForArrayFeatureTypeDataTypeConstraintClassDefaultOptionalValue(type_ unsafe.Pointer, type_2 int64, class objc.Class, value objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLModelIOUtilsClass.class), objc.Sel("populateConstraintsForArrayFeatureType:dataType:constraintClass:defaultOptionalValue:"), type_, type_2, class, value)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/CoreML/MLModelIOUtils/populateConstraintsForFeatureDescription:
func (_MLModelIOUtilsClass MLModelIOUtilsClass) PopulateConstraintsForFeatureDescription(description unsafe.Pointer) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLModelIOUtilsClass.class), objc.Sel("populateConstraintsForFeatureDescription:"), description)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/CoreML/MLModelIOUtils/populateConstraintsForImageFeatureDescription:
func (_MLModelIOUtilsClass MLModelIOUtilsClass) PopulateConstraintsForImageFeatureDescription(description unsafe.Pointer) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLModelIOUtilsClass.class), objc.Sel("populateConstraintsForImageFeatureDescription:"), description)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/CoreML/MLModelIOUtils/populateConstraintsForImageFeatureDescriptionElement:
func (_MLModelIOUtilsClass MLModelIOUtilsClass) PopulateConstraintsForImageFeatureDescriptionElement(element unsafe.Pointer) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLModelIOUtilsClass.class), objc.Sel("populateConstraintsForImageFeatureDescriptionElement:"), element)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/CoreML/MLModelIOUtils/rangeFromAllowedSizeRangeProtoMessage:
func (_MLModelIOUtilsClass MLModelIOUtilsClass) RangeFromAllowedSizeRangeProtoMessage(message unsafe.Pointer) foundation.NSRange {
	rv := objc.Send[foundation.NSRange](objc.ID(_MLModelIOUtilsClass.class), objc.Sel("rangeFromAllowedSizeRangeProtoMessage:"), message)
	return foundation.NSRange(rv)
}
//
// See: https://developer.apple.com/documentation/CoreML/MLModelIOUtils/serializeInterfaceFormat:archive:error:
func (_MLModelIOUtilsClass MLModelIOUtilsClass) SerializeInterfaceFormatArchiveError(format unsafe.Pointer, archive unsafe.Pointer) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](objc.ID(_MLModelIOUtilsClass.class), objc.Sel("serializeInterfaceFormat:archive:error:"), format, archive, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("serializeInterfaceFormat:archive:error: returned NO with nil NSError")
	}
	return rv, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLModelIOUtils/serializeMetadataAndInterfaceFromSpecification:archive:error:
func (_MLModelIOUtilsClass MLModelIOUtilsClass) SerializeMetadataAndInterfaceFromSpecificationArchiveError(specification unsafe.Pointer, archive unsafe.Pointer) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](objc.ID(_MLModelIOUtilsClass.class), objc.Sel("serializeMetadataAndInterfaceFromSpecification:archive:error:"), specification, archive, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("serializeMetadataAndInterfaceFromSpecification:archive:error: returned NO with nil NSError")
	}
	return rv, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLModelIOUtils/serializeSpecification:toArchive:error:
func (_MLModelIOUtilsClass MLModelIOUtilsClass) SerializeSpecificationToArchiveError(specification unsafe.Pointer, archive unsafe.Pointer) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLModelIOUtilsClass.class), objc.Sel("serializeSpecification:toArchive:error:"), specification, archive, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLModelIOUtils/serializeVersionInfo:archive:error:
func (_MLModelIOUtilsClass MLModelIOUtilsClass) SerializeVersionInfoArchiveError(info objectivec.IObject, archive unsafe.Pointer) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](objc.ID(_MLModelIOUtilsClass.class), objc.Sel("serializeVersionInfo:archive:error:"), info, archive, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("serializeVersionInfo:archive:error: returned NO with nil NSError")
	}
	return rv, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLModelIOUtils/specificationURLFromModelAtURL:error:
func (_MLModelIOUtilsClass MLModelIOUtilsClass) SpecificationURLFromModelAtURLError(url foundation.INSURL) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLModelIOUtilsClass.class), objc.Sel("specificationURLFromModelAtURL:error:"), url, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLModelIOUtils/stateDescriptionFromInterface:
func (_MLModelIOUtilsClass MLModelIOUtilsClass) StateDescriptionFromInterface(interface_ unsafe.Pointer) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLModelIOUtilsClass.class), objc.Sel("stateDescriptionFromInterface:"), interface_)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/CoreML/MLModelIOUtils/trainingInputDescriptionFromInterface:
func (_MLModelIOUtilsClass MLModelIOUtilsClass) TrainingInputDescriptionFromInterface(interface_ unsafe.Pointer) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLModelIOUtilsClass.class), objc.Sel("trainingInputDescriptionFromInterface:"), interface_)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/CoreML/MLModelIOUtils/versionForSerializedSpecification:options:error:
func (_MLModelIOUtilsClass MLModelIOUtilsClass) VersionForSerializedSpecificationOptionsError(specification unsafe.Pointer, options objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLModelIOUtilsClass.class), objc.Sel("versionForSerializedSpecification:options:error:"), specification, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

