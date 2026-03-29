// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLSNSoundPrint] class.
var (
	_MLSNSoundPrintClass     MLSNSoundPrintClass
	_MLSNSoundPrintClassOnce sync.Once
)

func getMLSNSoundPrintClass() MLSNSoundPrintClass {
	_MLSNSoundPrintClassOnce.Do(func() {
		_MLSNSoundPrintClass = MLSNSoundPrintClass{class: objc.GetClass("_MLSNSoundPrint")}
	})
	return _MLSNSoundPrintClass
}

// GetMLSNSoundPrintClass returns the class object for _MLSNSoundPrint.
func GetMLSNSoundPrintClass() MLSNSoundPrintClass {
	return getMLSNSoundPrintClass()
}

type MLSNSoundPrintClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLSNSoundPrintClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLSNSoundPrintClass) Alloc() MLSNSoundPrint {
	rv := objc.Send[MLSNSoundPrint](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [MLSNSoundPrint.ModelDescription]
//   - [MLSNSoundPrint.PredictionFromFeaturesOptionsError]
//   - [MLSNSoundPrint.InitWithModelDescriptionParameterDictionaryError]
// See: https://developer.apple.com/documentation/CoreML/_MLSNSoundPrint
type MLSNSoundPrint struct {
	objectivec.Object
}

// MLSNSoundPrintFromID constructs a [MLSNSoundPrint] from an objc.ID.
func MLSNSoundPrintFromID(id objc.ID) MLSNSoundPrint {
	return MLSNSoundPrint{objectivec.Object{ID: id}}
}
// Ensure MLSNSoundPrint implements IMLSNSoundPrint.
var _ IMLSNSoundPrint = MLSNSoundPrint{}

// An interface definition for the [MLSNSoundPrint] class.
//
// # Methods
//
//   - [IMLSNSoundPrint.ModelDescription]
//   - [IMLSNSoundPrint.PredictionFromFeaturesOptionsError]
//   - [IMLSNSoundPrint.InitWithModelDescriptionParameterDictionaryError]
//
// See: https://developer.apple.com/documentation/CoreML/_MLSNSoundPrint
type IMLSNSoundPrint interface {
	objectivec.IObject

	// Topic: Methods

	ModelDescription() IMLModelDescription
	PredictionFromFeaturesOptionsError(features objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error)
	InitWithModelDescriptionParameterDictionaryError(description objectivec.IObject, dictionary objectivec.IObject) (MLSNSoundPrint, error)
}

// Init initializes the instance.
func (m MLSNSoundPrint) Init() MLSNSoundPrint {
	rv := objc.Send[MLSNSoundPrint](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLSNSoundPrint) Autorelease() MLSNSoundPrint {
	rv := objc.Send[MLSNSoundPrint](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLSNSoundPrint creates a new MLSNSoundPrint instance.
func NewMLSNSoundPrint() MLSNSoundPrint {
	class := getMLSNSoundPrintClass()
	rv := objc.Send[MLSNSoundPrint](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/CoreML/_MLSNSoundPrint/initWithModelDescription:parameterDictionary:error:
func NewMLSNSoundPrintWithModelDescriptionParameterDictionaryError(description objectivec.IObject, dictionary objectivec.IObject) (MLSNSoundPrint, error) {
	var errorPtr objc.ID
	instance := getMLSNSoundPrintClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithModelDescription:parameterDictionary:error:"), description, dictionary, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLSNSoundPrint{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLSNSoundPrintFromID(rv), nil
}

//
// See: https://developer.apple.com/documentation/CoreML/_MLSNSoundPrint/predictionFromFeatures:options:error:
func (m MLSNSoundPrint) PredictionFromFeaturesOptionsError(features objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("predictionFromFeatures:options:error:"), features, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/_MLSNSoundPrint/initWithModelDescription:parameterDictionary:error:
func (m MLSNSoundPrint) InitWithModelDescriptionParameterDictionaryError(description objectivec.IObject, dictionary objectivec.IObject) (MLSNSoundPrint, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("initWithModelDescription:parameterDictionary:error:"), description, dictionary, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLSNSoundPrint{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLSNSoundPrintFromID(rv), nil

}

// See: https://developer.apple.com/documentation/CoreML/_MLSNSoundPrint/modelDescription
func (m MLSNSoundPrint) ModelDescription() IMLModelDescription {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("modelDescription"))
	return MLModelDescriptionFromID(objc.ID(rv))
}

