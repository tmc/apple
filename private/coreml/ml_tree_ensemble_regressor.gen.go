// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLTreeEnsembleRegressor] class.
var (
	_MLTreeEnsembleRegressorClass     MLTreeEnsembleRegressorClass
	_MLTreeEnsembleRegressorClassOnce sync.Once
)

func getMLTreeEnsembleRegressorClass() MLTreeEnsembleRegressorClass {
	_MLTreeEnsembleRegressorClassOnce.Do(func() {
		_MLTreeEnsembleRegressorClass = MLTreeEnsembleRegressorClass{class: objc.GetClass("MLTreeEnsembleRegressor")}
	})
	return _MLTreeEnsembleRegressorClass
}

// GetMLTreeEnsembleRegressorClass returns the class object for MLTreeEnsembleRegressor.
func GetMLTreeEnsembleRegressorClass() MLTreeEnsembleRegressorClass {
	return getMLTreeEnsembleRegressorClass()
}

type MLTreeEnsembleRegressorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLTreeEnsembleRegressorClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLTreeEnsembleRegressorClass) Alloc() MLTreeEnsembleRegressor {
	rv := objc.Send[MLTreeEnsembleRegressor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLTreeEnsembleRegressor.ModelData]
//   - [MLTreeEnsembleRegressor.RegressOptionsError]
//   - [MLTreeEnsembleRegressor.ScalarRegress]
//   - [MLTreeEnsembleRegressor.ScalarRegressError]
//   - [MLTreeEnsembleRegressor.VectorRegressDest]
//   - [MLTreeEnsembleRegressor.DebugDescription]
//   - [MLTreeEnsembleRegressor.Description]
//   - [MLTreeEnsembleRegressor.Hash]
//   - [MLTreeEnsembleRegressor.Superclass]
//
// See: https://developer.apple.com/documentation/CoreML/MLTreeEnsembleRegressor
type MLTreeEnsembleRegressor struct {
	objectivec.Object
}

// MLTreeEnsembleRegressorFromID constructs a [MLTreeEnsembleRegressor] from an objc.ID.
func MLTreeEnsembleRegressorFromID(id objc.ID) MLTreeEnsembleRegressor {
	return MLTreeEnsembleRegressor{objectivec.Object{ID: id}}
}

// NOTE: MLTreeEnsembleRegressor struct embeds objectivec.Object (parent type unavailable) but
// IMLTreeEnsembleRegressor embeds the parent interface; skip compile-time assertion.

// An interface definition for the [MLTreeEnsembleRegressor] class.
//
// # Methods
//
//   - [IMLTreeEnsembleRegressor.ModelData]
//   - [IMLTreeEnsembleRegressor.RegressOptionsError]
//   - [IMLTreeEnsembleRegressor.ScalarRegress]
//   - [IMLTreeEnsembleRegressor.ScalarRegressError]
//   - [IMLTreeEnsembleRegressor.VectorRegressDest]
//   - [IMLTreeEnsembleRegressor.DebugDescription]
//   - [IMLTreeEnsembleRegressor.Description]
//   - [IMLTreeEnsembleRegressor.Hash]
//   - [IMLTreeEnsembleRegressor.Superclass]
//
// See: https://developer.apple.com/documentation/CoreML/MLTreeEnsembleRegressor
type IMLTreeEnsembleRegressor interface {
	IMLRegressor

	// Topic: Methods

	ModelData() string
	RegressOptionsError(regress objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error)
	ScalarRegress(regress []float64) float64
	ScalarRegressError(regress objectivec.IObject) (float64, error)
	VectorRegressDest(regress []float64, dest []float64)
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objc.Class
}

// Init initializes the instance.
func (t MLTreeEnsembleRegressor) Init() MLTreeEnsembleRegressor {
	rv := objc.Send[MLTreeEnsembleRegressor](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t MLTreeEnsembleRegressor) Autorelease() MLTreeEnsembleRegressor {
	rv := objc.Send[MLTreeEnsembleRegressor](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLTreeEnsembleRegressor creates a new MLTreeEnsembleRegressor instance.
func NewMLTreeEnsembleRegressor() MLTreeEnsembleRegressor {
	class := getMLTreeEnsembleRegressorClass()
	rv := objc.Send[MLTreeEnsembleRegressor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLTreeEnsembleRegressor/modelData
func (t MLTreeEnsembleRegressor) ModelData() string {
	rv := objc.Send[*byte](t.ID, objc.Sel("modelData"))
	return objc.GoString(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLTreeEnsembleRegressor/regress:options:error:
func (t MLTreeEnsembleRegressor) RegressOptionsError(regress objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](t.ID, objc.Sel("regress:options:error:"), regress, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// See: https://developer.apple.com/documentation/CoreML/MLTreeEnsembleRegressor/scalarRegress:
func (t MLTreeEnsembleRegressor) ScalarRegress(regress []float64) float64 {
	rv := objc.Send[float64](t.ID, objc.Sel("scalarRegress:"), regress)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLTreeEnsembleRegressor/scalarRegress:error:
func (t MLTreeEnsembleRegressor) ScalarRegressError(regress objectivec.IObject) (float64, error) {
	var errorPtr objc.ID
	rv := objc.Send[float64](t.ID, objc.Sel("scalarRegress:error:"), regress, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return 0.0, foundation.NSErrorFrom(errorPtr)
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/CoreML/MLTreeEnsembleRegressor/vectorRegress:dest:
func (t MLTreeEnsembleRegressor) VectorRegressDest(regress []float64, dest []float64) {
	objc.Send[objc.ID](t.ID, objc.Sel("vectorRegress:dest:"), regress, dest)
}

// See: https://developer.apple.com/documentation/CoreML/MLTreeEnsembleRegressor/compileSpecification:toArchive:options:error:
func (_MLTreeEnsembleRegressorClass MLTreeEnsembleRegressorClass) CompileSpecificationToArchiveOptionsError(specification unsafe.Pointer, archive unsafe.Pointer, options objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLTreeEnsembleRegressorClass.class), objc.Sel("compileSpecification:toArchive:options:error:"), specification, archive, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// See: https://developer.apple.com/documentation/CoreML/MLTreeEnsembleRegressor/compiledVersionForSpecification:options:error:
func (_MLTreeEnsembleRegressorClass MLTreeEnsembleRegressorClass) CompiledVersionForSpecificationOptionsError(specification unsafe.Pointer, options objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLTreeEnsembleRegressorClass.class), objc.Sel("compiledVersionForSpecification:options:error:"), specification, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// See: https://developer.apple.com/documentation/CoreML/MLTreeEnsembleRegressor/loadModelFromCompiledArchive:modelVersionInfo:compilerVersionInfo:configuration:error:
func (_MLTreeEnsembleRegressorClass MLTreeEnsembleRegressorClass) LoadModelFromCompiledArchiveModelVersionInfoCompilerVersionInfoConfigurationError(archive unsafe.Pointer, info objectivec.IObject, info2 objectivec.IObject, configuration objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLTreeEnsembleRegressorClass.class), objc.Sel("loadModelFromCompiledArchive:modelVersionInfo:compilerVersionInfo:configuration:error:"), archive, info, info2, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// See: https://developer.apple.com/documentation/CoreML/MLTreeEnsembleRegressor/loadModelFromSpecification:configuration:error:
func (_MLTreeEnsembleRegressorClass MLTreeEnsembleRegressorClass) LoadModelFromSpecificationConfigurationError(specification unsafe.Pointer, configuration objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLTreeEnsembleRegressorClass.class), objc.Sel("loadModelFromSpecification:configuration:error:"), specification, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// See: https://developer.apple.com/documentation/CoreML/MLTreeEnsembleRegressor/loadModelFromSpecificationWithCompilationOptions:options:error:
func (_MLTreeEnsembleRegressorClass MLTreeEnsembleRegressorClass) LoadModelFromSpecificationWithCompilationOptionsOptionsError(options unsafe.Pointer, options2 objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLTreeEnsembleRegressorClass.class), objc.Sel("loadModelFromSpecificationWithCompilationOptions:options:error:"), options, options2, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// See: https://developer.apple.com/documentation/CoreML/MLTreeEnsembleRegressor/debugDescription
func (t MLTreeEnsembleRegressor) DebugDescription() string {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/CoreML/MLTreeEnsembleRegressor/description
func (t MLTreeEnsembleRegressor) Description() string {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/CoreML/MLTreeEnsembleRegressor/hash
func (t MLTreeEnsembleRegressor) Hash() uint64 {
	rv := objc.Send[uint64](t.ID, objc.Sel("hash"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLTreeEnsembleRegressor/superclass
func (t MLTreeEnsembleRegressor) Superclass() objc.Class {
	rv := objc.Send[objc.Class](t.ID, objc.Sel("superclass"))
	return rv
}
