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

// The class instance for the [MLGLMClassification] class.
var (
	_MLGLMClassificationClass     MLGLMClassificationClass
	_MLGLMClassificationClassOnce sync.Once
)

func getMLGLMClassificationClass() MLGLMClassificationClass {
	_MLGLMClassificationClassOnce.Do(func() {
		_MLGLMClassificationClass = MLGLMClassificationClass{class: objc.GetClass("MLGLMClassification")}
	})
	return _MLGLMClassificationClass
}

// GetMLGLMClassificationClass returns the class object for MLGLMClassification.
func GetMLGLMClassificationClass() MLGLMClassificationClass {
	return getMLGLMClassificationClass()
}

type MLGLMClassificationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLGLMClassificationClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLGLMClassificationClass) Alloc() MLGLMClassification {
	rv := objc.Send[MLGLMClassification](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLGLMClassification.CalculateClassProbabilityInputError]
//   - [MLGLMClassification.ClassifyError]
//   - [MLGLMClassification.ClassifyOptionsError]
//   - [MLGLMClassification.ClassifyTopKError]
//   - [MLGLMClassification.InitWithSpecificationConfigurationError]
//   - [MLGLMClassification.DebugDescription]
//   - [MLGLMClassification.Description]
//   - [MLGLMClassification.Hash]
//   - [MLGLMClassification.Superclass]
//
// See: https://developer.apple.com/documentation/CoreML/MLGLMClassification
type MLGLMClassification struct {
	objectivec.Object
}

// MLGLMClassificationFromID constructs a [MLGLMClassification] from an objc.ID.
func MLGLMClassificationFromID(id objc.ID) MLGLMClassification {
	return MLGLMClassification{objectivec.Object{ID: id}}
}

// NOTE: MLGLMClassification struct embeds objectivec.Object (parent type unavailable) but
// IMLGLMClassification embeds the parent interface; skip compile-time assertion.

// An interface definition for the [MLGLMClassification] class.
//
// # Methods
//
//   - [IMLGLMClassification.CalculateClassProbabilityInputError]
//   - [IMLGLMClassification.ClassifyError]
//   - [IMLGLMClassification.ClassifyOptionsError]
//   - [IMLGLMClassification.ClassifyTopKError]
//   - [IMLGLMClassification.InitWithSpecificationConfigurationError]
//   - [IMLGLMClassification.DebugDescription]
//   - [IMLGLMClassification.Description]
//   - [IMLGLMClassification.Hash]
//   - [IMLGLMClassification.Superclass]
//
// See: https://developer.apple.com/documentation/CoreML/MLGLMClassification
type IMLGLMClassification interface {
	IMLClassifier

	// Topic: Methods

	CalculateClassProbabilityInputError(input objectivec.IObject) (float64, error)
	ClassifyError(classify objectivec.IObject) (objectivec.IObject, error)
	ClassifyOptionsError(classify objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error)
	ClassifyTopKError(classify objectivec.IObject, k uint64) (objectivec.IObject, error)
	InitWithSpecificationConfigurationError(specification unsafe.Pointer, configuration objectivec.IObject) (MLGLMClassification, error)
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objc.Class
}

// Init initializes the instance.
func (g MLGLMClassification) Init() MLGLMClassification {
	rv := objc.Send[MLGLMClassification](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g MLGLMClassification) Autorelease() MLGLMClassification {
	rv := objc.Send[MLGLMClassification](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLGLMClassification creates a new MLGLMClassification instance.
func NewMLGLMClassification() MLGLMClassification {
	class := getMLGLMClassificationClass()
	rv := objc.Send[MLGLMClassification](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLGLMClassification/initWithSpecification:configuration:error:
func NewGLMClassificationWithSpecificationConfigurationError(specification unsafe.Pointer, configuration objectivec.IObject) (MLGLMClassification, error) {
	var errorPtr objc.ID
	instance := getMLGLMClassificationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSpecification:configuration:error:"), specification, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLGLMClassification{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLGLMClassificationFromID(rv), nil
}

// See: https://developer.apple.com/documentation/CoreML/MLGLMClassification/calculateClassProbability:input:error:
func (g MLGLMClassification) CalculateClassProbabilityInputError(input objectivec.IObject) (float64, error) {
	var probability float64
	var errorPtr objc.ID
	rv := objc.Send[bool](g.ID, objc.Sel("calculateClassProbability:input:error:"), unsafe.Pointer(&probability), input, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return 0.0, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return 0.0, errors.New("calculateClassProbability:input:error: returned NO with nil NSError")
	}
	return probability, nil
}

// See: https://developer.apple.com/documentation/CoreML/MLGLMClassification/classify:error:
func (g MLGLMClassification) ClassifyError(classify objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](g.ID, objc.Sel("classify:error:"), classify, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// See: https://developer.apple.com/documentation/CoreML/MLGLMClassification/classify:options:error:
func (g MLGLMClassification) ClassifyOptionsError(classify objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](g.ID, objc.Sel("classify:options:error:"), classify, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// See: https://developer.apple.com/documentation/CoreML/MLGLMClassification/classify:topK:error:
func (g MLGLMClassification) ClassifyTopKError(classify objectivec.IObject, k uint64) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](g.ID, objc.Sel("classify:topK:error:"), classify, k, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// See: https://developer.apple.com/documentation/CoreML/MLGLMClassification/initWithSpecification:configuration:error:
func (g MLGLMClassification) InitWithSpecificationConfigurationError(specification unsafe.Pointer, configuration objectivec.IObject) (MLGLMClassification, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](g.ID, objc.Sel("initWithSpecification:configuration:error:"), specification, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLGLMClassification{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLGLMClassificationFromID(rv), nil

}

// See: https://developer.apple.com/documentation/CoreML/MLGLMClassification/loadModelFromSpecification:configuration:error:
func (_MLGLMClassificationClass MLGLMClassificationClass) LoadModelFromSpecificationConfigurationError(specification unsafe.Pointer, configuration objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLGLMClassificationClass.class), objc.Sel("loadModelFromSpecification:configuration:error:"), specification, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// See: https://developer.apple.com/documentation/CoreML/MLGLMClassification/debugDescription
func (g MLGLMClassification) DebugDescription() string {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/CoreML/MLGLMClassification/description
func (g MLGLMClassification) Description() string {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/CoreML/MLGLMClassification/hash
func (g MLGLMClassification) Hash() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("hash"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLGLMClassification/superclass
func (g MLGLMClassification) Superclass() objc.Class {
	rv := objc.Send[objc.Class](g.ID, objc.Sel("superclass"))
	return rv
}
