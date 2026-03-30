// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"errors"
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLModelDescriptionUtils] class.
var (
	_MLModelDescriptionUtilsClass     MLModelDescriptionUtilsClass
	_MLModelDescriptionUtilsClassOnce sync.Once
)

func getMLModelDescriptionUtilsClass() MLModelDescriptionUtilsClass {
	_MLModelDescriptionUtilsClassOnce.Do(func() {
		_MLModelDescriptionUtilsClass = MLModelDescriptionUtilsClass{class: objc.GetClass("MLModelDescriptionUtils")}
	})
	return _MLModelDescriptionUtilsClass
}

// GetMLModelDescriptionUtilsClass returns the class object for MLModelDescriptionUtils.
func GetMLModelDescriptionUtilsClass() MLModelDescriptionUtilsClass {
	return getMLModelDescriptionUtilsClass()
}

type MLModelDescriptionUtilsClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLModelDescriptionUtilsClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLModelDescriptionUtilsClass) Alloc() MLModelDescriptionUtils {
	rv := objc.Send[MLModelDescriptionUtils](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLModelDescriptionUtils
type MLModelDescriptionUtils struct {
	objectivec.Object
}

// MLModelDescriptionUtilsFromID constructs a [MLModelDescriptionUtils] from an objc.ID.
func MLModelDescriptionUtilsFromID(id objc.ID) MLModelDescriptionUtils {
	return MLModelDescriptionUtils{objectivec.Object{ID: id}}
}

// Ensure MLModelDescriptionUtils implements IMLModelDescriptionUtils.
var _ IMLModelDescriptionUtils = MLModelDescriptionUtils{}

// An interface definition for the [MLModelDescriptionUtils] class.
//
// See: https://developer.apple.com/documentation/CoreML/MLModelDescriptionUtils
type IMLModelDescriptionUtils interface {
	objectivec.IObject
}

// Init initializes the instance.
func (m MLModelDescriptionUtils) Init() MLModelDescriptionUtils {
	rv := objc.Send[MLModelDescriptionUtils](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLModelDescriptionUtils) Autorelease() MLModelDescriptionUtils {
	rv := objc.Send[MLModelDescriptionUtils](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLModelDescriptionUtils creates a new MLModelDescriptionUtils instance.
func NewMLModelDescriptionUtils() MLModelDescriptionUtils {
	class := getMLModelDescriptionUtilsClass()
	rv := objc.Send[MLModelDescriptionUtils](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLModelDescriptionUtils/copyFeatureDescriptionFrom:to:error:
func (_MLModelDescriptionUtilsClass MLModelDescriptionUtilsClass) CopyFeatureDescriptionFromToError(from objectivec.IObject, to unsafe.Pointer) error {
	var errorPtr objc.ID
	objc.Send[struct{}](objc.ID(_MLModelDescriptionUtilsClass.class), objc.Sel("copyFeatureDescriptionFrom:to:error:"), from, to, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return foundation.NSErrorFrom(errorPtr)
	}
	return nil

}

// See: https://developer.apple.com/documentation/CoreML/MLModelDescriptionUtils/createFeatureTypeFromFeatureDescription:error:
func (_MLModelDescriptionUtilsClass MLModelDescriptionUtilsClass) CreateFeatureTypeFromFeatureDescriptionError(description objectivec.IObject) (unsafe.Pointer, error) {
	var errorPtr objc.ID
	rv := objc.Send[unsafe.Pointer](objc.ID(_MLModelDescriptionUtilsClass.class), objc.Sel("createFeatureTypeFromFeatureDescription:error:"), description, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/CoreML/MLModelDescriptionUtils/createMetaData:
func (_MLModelDescriptionUtilsClass MLModelDescriptionUtilsClass) CreateMetaData(data objectivec.IObject) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(_MLModelDescriptionUtilsClass.class), objc.Sel("createMetaData:"), data)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLModelDescriptionUtils/createModelDescription:error:
func (_MLModelDescriptionUtilsClass MLModelDescriptionUtilsClass) CreateModelDescriptionError(description objectivec.IObject) (unsafe.Pointer, error) {
	var errorPtr objc.ID
	rv := objc.Send[unsafe.Pointer](objc.ID(_MLModelDescriptionUtilsClass.class), objc.Sel("createModelDescription:error:"), description, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/CoreML/MLModelDescriptionUtils/getArrayFeatureTypeFromConstraint:
func (_MLModelDescriptionUtilsClass MLModelDescriptionUtilsClass) GetArrayFeatureTypeFromConstraint(constraint objectivec.IObject) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(_MLModelDescriptionUtilsClass.class), objc.Sel("getArrayFeatureTypeFromConstraint:"), constraint)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLModelDescriptionUtils/getDictionaryFeatureTypeFromConstraint:error:
func (_MLModelDescriptionUtilsClass MLModelDescriptionUtilsClass) GetDictionaryFeatureTypeFromConstraintError(constraint objectivec.IObject) (unsafe.Pointer, error) {
	var errorPtr objc.ID
	rv := objc.Send[unsafe.Pointer](objc.ID(_MLModelDescriptionUtilsClass.class), objc.Sel("getDictionaryFeatureTypeFromConstraint:error:"), constraint, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/CoreML/MLModelDescriptionUtils/getImageFeatureTypeFromConstraint:
func (_MLModelDescriptionUtilsClass MLModelDescriptionUtilsClass) GetImageFeatureTypeFromConstraint(constraint objectivec.IObject) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(_MLModelDescriptionUtilsClass.class), objc.Sel("getImageFeatureTypeFromConstraint:"), constraint)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLModelDescriptionUtils/getSequenceFeatureTypeFromConstraint:error:
func (_MLModelDescriptionUtilsClass MLModelDescriptionUtilsClass) GetSequenceFeatureTypeFromConstraintError(constraint objectivec.IObject) (unsafe.Pointer, error) {
	var errorPtr objc.ID
	rv := objc.Send[unsafe.Pointer](objc.ID(_MLModelDescriptionUtilsClass.class), objc.Sel("getSequenceFeatureTypeFromConstraint:error:"), constraint, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/CoreML/MLModelDescriptionUtils/getStateFeatureTypeFromConstraint:
func (_MLModelDescriptionUtilsClass MLModelDescriptionUtilsClass) GetStateFeatureTypeFromConstraint(constraint objectivec.IObject) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(_MLModelDescriptionUtilsClass.class), objc.Sel("getStateFeatureTypeFromConstraint:"), constraint)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLModelDescriptionUtils/saveModelDescription:toSpecification:error:
func (_MLModelDescriptionUtilsClass MLModelDescriptionUtilsClass) SaveModelDescriptionToSpecificationError(description objectivec.IObject, specification unsafe.Pointer) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](objc.ID(_MLModelDescriptionUtilsClass.class), objc.Sel("saveModelDescription:toSpecification:error:"), description, specification, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("saveModelDescription:toSpecification:error: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/CoreML/MLModelDescriptionUtils/validateAllFeatureDescriptions:hasAnyFeatureTypeIn:minimalCount:maximumCount:debugLabel:error:
func (_MLModelDescriptionUtilsClass MLModelDescriptionUtilsClass) ValidateAllFeatureDescriptionsHasAnyFeatureTypeInMinimalCountMaximumCountDebugLabelError(descriptions objectivec.IObject, in objectivec.IObject, count uint64, count2 uint64, label objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](objc.ID(_MLModelDescriptionUtilsClass.class), objc.Sel("validateAllFeatureDescriptions:hasAnyFeatureTypeIn:minimalCount:maximumCount:debugLabel:error:"), descriptions, in, count, count2, label, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("validateAllFeatureDescriptions:hasAnyFeatureTypeIn:minimalCount:maximumCount:debugLabel:error: returned NO with nil NSError")
	}
	return rv, nil

}
