// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLAppleImageFeatureExtractorObjectPrintParameters] class.
var (
	_MLAppleImageFeatureExtractorObjectPrintParametersClass     MLAppleImageFeatureExtractorObjectPrintParametersClass
	_MLAppleImageFeatureExtractorObjectPrintParametersClassOnce sync.Once
)

func getMLAppleImageFeatureExtractorObjectPrintParametersClass() MLAppleImageFeatureExtractorObjectPrintParametersClass {
	_MLAppleImageFeatureExtractorObjectPrintParametersClassOnce.Do(func() {
		_MLAppleImageFeatureExtractorObjectPrintParametersClass = MLAppleImageFeatureExtractorObjectPrintParametersClass{class: objc.GetClass("MLAppleImageFeatureExtractorObjectPrintParameters")}
	})
	return _MLAppleImageFeatureExtractorObjectPrintParametersClass
}

// GetMLAppleImageFeatureExtractorObjectPrintParametersClass returns the class object for MLAppleImageFeatureExtractorObjectPrintParameters.
func GetMLAppleImageFeatureExtractorObjectPrintParametersClass() MLAppleImageFeatureExtractorObjectPrintParametersClass {
	return getMLAppleImageFeatureExtractorObjectPrintParametersClass()
}

type MLAppleImageFeatureExtractorObjectPrintParametersClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLAppleImageFeatureExtractorObjectPrintParametersClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLAppleImageFeatureExtractorObjectPrintParametersClass) Alloc() MLAppleImageFeatureExtractorObjectPrintParameters {
	rv := objc.Send[MLAppleImageFeatureExtractorObjectPrintParameters](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLAppleImageFeatureExtractorObjectPrintParameters.ExpectedKeys]
//   - [MLAppleImageFeatureExtractorObjectPrintParameters.ExpectedShapes]
//   - [MLAppleImageFeatureExtractorObjectPrintParameters.ObjectPrintVersion]
//   - [MLAppleImageFeatureExtractorObjectPrintParameters.InitObjectPrintParametersExpectedShapesExpectedKeysError]
//
// See: https://developer.apple.com/documentation/CoreML/MLAppleImageFeatureExtractorObjectPrintParameters
type MLAppleImageFeatureExtractorObjectPrintParameters struct {
	objectivec.Object
}

// MLAppleImageFeatureExtractorObjectPrintParametersFromID constructs a [MLAppleImageFeatureExtractorObjectPrintParameters] from an objc.ID.
func MLAppleImageFeatureExtractorObjectPrintParametersFromID(id objc.ID) MLAppleImageFeatureExtractorObjectPrintParameters {
	return MLAppleImageFeatureExtractorObjectPrintParameters{objectivec.Object{ID: id}}
}

// Ensure MLAppleImageFeatureExtractorObjectPrintParameters implements IMLAppleImageFeatureExtractorObjectPrintParameters.
var _ IMLAppleImageFeatureExtractorObjectPrintParameters = MLAppleImageFeatureExtractorObjectPrintParameters{}

// An interface definition for the [MLAppleImageFeatureExtractorObjectPrintParameters] class.
//
// # Methods
//
//   - [IMLAppleImageFeatureExtractorObjectPrintParameters.ExpectedKeys]
//   - [IMLAppleImageFeatureExtractorObjectPrintParameters.ExpectedShapes]
//   - [IMLAppleImageFeatureExtractorObjectPrintParameters.ObjectPrintVersion]
//   - [IMLAppleImageFeatureExtractorObjectPrintParameters.InitObjectPrintParametersExpectedShapesExpectedKeysError]
//
// See: https://developer.apple.com/documentation/CoreML/MLAppleImageFeatureExtractorObjectPrintParameters
type IMLAppleImageFeatureExtractorObjectPrintParameters interface {
	objectivec.IObject

	// Topic: Methods

	ExpectedKeys() foundation.INSArray
	ExpectedShapes() foundation.INSArray
	ObjectPrintVersion() uint64
	InitObjectPrintParametersExpectedShapesExpectedKeysError(parameters uint64, shapes objectivec.IObject, keys objectivec.IObject) (MLAppleImageFeatureExtractorObjectPrintParameters, error)
}

// Init initializes the instance.
func (a MLAppleImageFeatureExtractorObjectPrintParameters) Init() MLAppleImageFeatureExtractorObjectPrintParameters {
	rv := objc.Send[MLAppleImageFeatureExtractorObjectPrintParameters](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a MLAppleImageFeatureExtractorObjectPrintParameters) Autorelease() MLAppleImageFeatureExtractorObjectPrintParameters {
	rv := objc.Send[MLAppleImageFeatureExtractorObjectPrintParameters](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLAppleImageFeatureExtractorObjectPrintParameters creates a new MLAppleImageFeatureExtractorObjectPrintParameters instance.
func NewMLAppleImageFeatureExtractorObjectPrintParameters() MLAppleImageFeatureExtractorObjectPrintParameters {
	class := getMLAppleImageFeatureExtractorObjectPrintParametersClass()
	rv := objc.Send[MLAppleImageFeatureExtractorObjectPrintParameters](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLAppleImageFeatureExtractorObjectPrintParameters/initObjectPrintParameters:expectedShapes:expectedKeys:error:
func NewAppleImageFeatureExtractorObjectPrintParametersObjectPrintParametersExpectedShapesExpectedKeysError(parameters uint64, shapes objectivec.IObject, keys objectivec.IObject) (MLAppleImageFeatureExtractorObjectPrintParameters, error) {
	var errorPtr objc.ID
	instance := getMLAppleImageFeatureExtractorObjectPrintParametersClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initObjectPrintParameters:expectedShapes:expectedKeys:error:"), parameters, shapes, keys, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLAppleImageFeatureExtractorObjectPrintParameters{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLAppleImageFeatureExtractorObjectPrintParametersFromID(rv), nil
}

// See: https://developer.apple.com/documentation/CoreML/MLAppleImageFeatureExtractorObjectPrintParameters/initObjectPrintParameters:expectedShapes:expectedKeys:error:
func (a MLAppleImageFeatureExtractorObjectPrintParameters) InitObjectPrintParametersExpectedShapesExpectedKeysError(parameters uint64, shapes objectivec.IObject, keys objectivec.IObject) (MLAppleImageFeatureExtractorObjectPrintParameters, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](a.ID, objc.Sel("initObjectPrintParameters:expectedShapes:expectedKeys:error:"), parameters, shapes, keys, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLAppleImageFeatureExtractorObjectPrintParameters{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLAppleImageFeatureExtractorObjectPrintParametersFromID(rv), nil

}

// See: https://developer.apple.com/documentation/CoreML/MLAppleImageFeatureExtractorObjectPrintParameters/expectedKeys
func (a MLAppleImageFeatureExtractorObjectPrintParameters) ExpectedKeys() foundation.INSArray {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("expectedKeys"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/MLAppleImageFeatureExtractorObjectPrintParameters/expectedShapes
func (a MLAppleImageFeatureExtractorObjectPrintParameters) ExpectedShapes() foundation.INSArray {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("expectedShapes"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/MLAppleImageFeatureExtractorObjectPrintParameters/objectPrintVersion
func (a MLAppleImageFeatureExtractorObjectPrintParameters) ObjectPrintVersion() uint64 {
	rv := objc.Send[uint64](a.ID, objc.Sel("objectPrintVersion"))
	return rv
}
