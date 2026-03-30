// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLGLMRegression] class.
var (
	_MLGLMRegressionClass     MLGLMRegressionClass
	_MLGLMRegressionClassOnce sync.Once
)

func getMLGLMRegressionClass() MLGLMRegressionClass {
	_MLGLMRegressionClassOnce.Do(func() {
		_MLGLMRegressionClass = MLGLMRegressionClass{class: objc.GetClass("MLGLMRegression")}
	})
	return _MLGLMRegressionClass
}

// GetMLGLMRegressionClass returns the class object for MLGLMRegression.
func GetMLGLMRegressionClass() MLGLMRegressionClass {
	return getMLGLMRegressionClass()
}

type MLGLMRegressionClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLGLMRegressionClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLGLMRegressionClass) Alloc() MLGLMRegression {
	rv := objc.Send[MLGLMRegression](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLGLMRegression.RegressOptionsError]
//   - [MLGLMRegression.InitWithLRSpecConfigurationError]
//   - [MLGLMRegression.InitWithSpecificationConfigurationError]
//   - [MLGLMRegression.DebugDescription]
//   - [MLGLMRegression.Description]
//   - [MLGLMRegression.Hash]
//   - [MLGLMRegression.Superclass]
//
// See: https://developer.apple.com/documentation/CoreML/MLGLMRegression
type MLGLMRegression struct {
	objectivec.Object
}

// MLGLMRegressionFromID constructs a [MLGLMRegression] from an objc.ID.
func MLGLMRegressionFromID(id objc.ID) MLGLMRegression {
	return MLGLMRegression{objectivec.Object{ID: id}}
}

// NOTE: MLGLMRegression struct embeds objectivec.Object (parent type unavailable) but
// IMLGLMRegression embeds the parent interface; skip compile-time assertion.

// An interface definition for the [MLGLMRegression] class.
//
// # Methods
//
//   - [IMLGLMRegression.RegressOptionsError]
//   - [IMLGLMRegression.InitWithLRSpecConfigurationError]
//   - [IMLGLMRegression.InitWithSpecificationConfigurationError]
//   - [IMLGLMRegression.DebugDescription]
//   - [IMLGLMRegression.Description]
//   - [IMLGLMRegression.Hash]
//   - [IMLGLMRegression.Superclass]
//
// See: https://developer.apple.com/documentation/CoreML/MLGLMRegression
type IMLGLMRegression interface {
	IMLRegressor

	// Topic: Methods

	RegressOptionsError(regress objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error)
	InitWithLRSpecConfigurationError(lRSpec unsafe.Pointer, configuration objectivec.IObject) (MLGLMRegression, error)
	InitWithSpecificationConfigurationError(specification unsafe.Pointer, configuration objectivec.IObject) (MLGLMRegression, error)
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objc.Class
}

// Init initializes the instance.
func (g MLGLMRegression) Init() MLGLMRegression {
	rv := objc.Send[MLGLMRegression](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g MLGLMRegression) Autorelease() MLGLMRegression {
	rv := objc.Send[MLGLMRegression](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLGLMRegression creates a new MLGLMRegression instance.
func NewMLGLMRegression() MLGLMRegression {
	class := getMLGLMRegressionClass()
	rv := objc.Send[MLGLMRegression](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLGLMRegression/initWithLRSpec:configuration:error:
func NewGLMRegressionWithLRSpecConfigurationError(lRSpec unsafe.Pointer, configuration objectivec.IObject) (MLGLMRegression, error) {
	var errorPtr objc.ID
	instance := getMLGLMRegressionClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithLRSpec:configuration:error:"), lRSpec, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLGLMRegression{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLGLMRegressionFromID(rv), nil
}

// See: https://developer.apple.com/documentation/CoreML/MLGLMRegression/initWithSpecification:configuration:error:
func NewGLMRegressionWithSpecificationConfigurationError(specification unsafe.Pointer, configuration objectivec.IObject) (MLGLMRegression, error) {
	var errorPtr objc.ID
	instance := getMLGLMRegressionClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSpecification:configuration:error:"), specification, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLGLMRegression{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLGLMRegressionFromID(rv), nil
}

// See: https://developer.apple.com/documentation/CoreML/MLGLMRegression/regress:options:error:
func (g MLGLMRegression) RegressOptionsError(regress objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](g.ID, objc.Sel("regress:options:error:"), regress, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// See: https://developer.apple.com/documentation/CoreML/MLGLMRegression/initWithLRSpec:configuration:error:
func (g MLGLMRegression) InitWithLRSpecConfigurationError(lRSpec unsafe.Pointer, configuration objectivec.IObject) (MLGLMRegression, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](g.ID, objc.Sel("initWithLRSpec:configuration:error:"), lRSpec, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLGLMRegression{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLGLMRegressionFromID(rv), nil

}

// See: https://developer.apple.com/documentation/CoreML/MLGLMRegression/initWithSpecification:configuration:error:
func (g MLGLMRegression) InitWithSpecificationConfigurationError(specification unsafe.Pointer, configuration objectivec.IObject) (MLGLMRegression, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](g.ID, objc.Sel("initWithSpecification:configuration:error:"), specification, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLGLMRegression{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLGLMRegressionFromID(rv), nil

}

// See: https://developer.apple.com/documentation/CoreML/MLGLMRegression/loadModelFromSpecification:configuration:error:
func (_MLGLMRegressionClass MLGLMRegressionClass) LoadModelFromSpecificationConfigurationError(specification unsafe.Pointer, configuration objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLGLMRegressionClass.class), objc.Sel("loadModelFromSpecification:configuration:error:"), specification, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// See: https://developer.apple.com/documentation/CoreML/MLGLMRegression/debugDescription
func (g MLGLMRegression) DebugDescription() string {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/CoreML/MLGLMRegression/description
func (g MLGLMRegression) Description() string {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/CoreML/MLGLMRegression/hash
func (g MLGLMRegression) Hash() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("hash"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLGLMRegression/superclass
func (g MLGLMRegression) Superclass() objc.Class {
	rv := objc.Send[objc.Class](g.ID, objc.Sel("superclass"))
	return rv
}
