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

// The class instance for the [MLTreeEnsembleClassifier] class.
var (
	_MLTreeEnsembleClassifierClass     MLTreeEnsembleClassifierClass
	_MLTreeEnsembleClassifierClassOnce sync.Once
)

func getMLTreeEnsembleClassifierClass() MLTreeEnsembleClassifierClass {
	_MLTreeEnsembleClassifierClassOnce.Do(func() {
		_MLTreeEnsembleClassifierClass = MLTreeEnsembleClassifierClass{class: objc.GetClass("MLTreeEnsembleClassifier")}
	})
	return _MLTreeEnsembleClassifierClass
}

// GetMLTreeEnsembleClassifierClass returns the class object for MLTreeEnsembleClassifier.
func GetMLTreeEnsembleClassifierClass() MLTreeEnsembleClassifierClass {
	return getMLTreeEnsembleClassifierClass()
}

type MLTreeEnsembleClassifierClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLTreeEnsembleClassifierClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLTreeEnsembleClassifierClass) Alloc() MLTreeEnsembleClassifier {
	rv := objc.Send[MLTreeEnsembleClassifier](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [MLTreeEnsembleClassifier._buildClassificationClassesTopkError]
//   - [MLTreeEnsembleClassifier._setSingleArrayLookupField]
//   - [MLTreeEnsembleClassifier.ClassifyOptionsError]
//   - [MLTreeEnsembleClassifier.ModelData]
//   - [MLTreeEnsembleClassifier.PrepareInputError]
//   - [MLTreeEnsembleClassifier.DebugDescription]
//   - [MLTreeEnsembleClassifier.Description]
//   - [MLTreeEnsembleClassifier.Hash]
//   - [MLTreeEnsembleClassifier.Superclass]
// See: https://developer.apple.com/documentation/CoreML/MLTreeEnsembleClassifier
type MLTreeEnsembleClassifier struct {
	objectivec.Object
}

// MLTreeEnsembleClassifierFromID constructs a [MLTreeEnsembleClassifier] from an objc.ID.
func MLTreeEnsembleClassifierFromID(id objc.ID) MLTreeEnsembleClassifier {
	return MLTreeEnsembleClassifier{objectivec.Object{ID: id}}
}
// NOTE: MLTreeEnsembleClassifier struct embeds objectivec.Object (parent type unavailable) but
// IMLTreeEnsembleClassifier embeds the parent interface; skip compile-time assertion.

// An interface definition for the [MLTreeEnsembleClassifier] class.
//
// # Methods
//
//   - [IMLTreeEnsembleClassifier._buildClassificationClassesTopkError]
//   - [IMLTreeEnsembleClassifier._setSingleArrayLookupField]
//   - [IMLTreeEnsembleClassifier.ClassifyOptionsError]
//   - [IMLTreeEnsembleClassifier.ModelData]
//   - [IMLTreeEnsembleClassifier.PrepareInputError]
//   - [IMLTreeEnsembleClassifier.DebugDescription]
//   - [IMLTreeEnsembleClassifier.Description]
//   - [IMLTreeEnsembleClassifier.Hash]
//   - [IMLTreeEnsembleClassifier.Superclass]
//
// See: https://developer.apple.com/documentation/CoreML/MLTreeEnsembleClassifier
type IMLTreeEnsembleClassifier interface {
	IMLClassifier

	// Topic: Methods

	_buildClassificationClassesTopkError(classes []float64, topk uint64) (objectivec.IObject, error)
	_setSingleArrayLookupField()
	ClassifyOptionsError(classify objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error)
	ModelData() string
	PrepareInputError(input objectivec.IObject) (objectivec.IObject, error)
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objc.Class
}

// Init initializes the instance.
func (t MLTreeEnsembleClassifier) Init() MLTreeEnsembleClassifier {
	rv := objc.Send[MLTreeEnsembleClassifier](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t MLTreeEnsembleClassifier) Autorelease() MLTreeEnsembleClassifier {
	rv := objc.Send[MLTreeEnsembleClassifier](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLTreeEnsembleClassifier creates a new MLTreeEnsembleClassifier instance.
func NewMLTreeEnsembleClassifier() MLTreeEnsembleClassifier {
	class := getMLTreeEnsembleClassifierClass()
	rv := objc.Send[MLTreeEnsembleClassifier](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/CoreML/MLTreeEnsembleClassifier/_buildClassificationClasses:topk:error:
func (t MLTreeEnsembleClassifier) _buildClassificationClassesTopkError(classes []float64, topk uint64) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](t.ID, objc.Sel("_buildClassificationClasses:topk:error:"), classes, topk, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// BuildClassificationClassesTopkError is an exported wrapper for the private method _buildClassificationClassesTopkError.
func (t MLTreeEnsembleClassifier) BuildClassificationClassesTopkError(classes []float64, topk uint64) (objectivec.IObject, error) {
	return t._buildClassificationClassesTopkError(classes, topk)
}
// See: https://developer.apple.com/documentation/CoreML/MLTreeEnsembleClassifier/_setSingleArrayLookupField
func (t MLTreeEnsembleClassifier) _setSingleArrayLookupField() {
	objc.Send[objc.ID](t.ID, objc.Sel("_setSingleArrayLookupField"))
}

// SetSingleArrayLookupField is an exported wrapper for the private method _setSingleArrayLookupField.
func (t MLTreeEnsembleClassifier) SetSingleArrayLookupField() {
	t._setSingleArrayLookupField()
}
//
// See: https://developer.apple.com/documentation/CoreML/MLTreeEnsembleClassifier/classify:options:error:
func (t MLTreeEnsembleClassifier) ClassifyOptionsError(classify objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](t.ID, objc.Sel("classify:options:error:"), classify, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
// See: https://developer.apple.com/documentation/CoreML/MLTreeEnsembleClassifier/modelData
func (t MLTreeEnsembleClassifier) ModelData() string {
	rv := objc.Send[*byte](t.ID, objc.Sel("modelData"))
	return objc.GoString(rv)
}
//
// See: https://developer.apple.com/documentation/CoreML/MLTreeEnsembleClassifier/prepareInput:error:
func (t MLTreeEnsembleClassifier) PrepareInputError(input objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](t.ID, objc.Sel("prepareInput:error:"), input, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

//
// See: https://developer.apple.com/documentation/CoreML/MLTreeEnsembleClassifier/_convertStringClassVector:int64ClassVector:dimensions:toClassLabel:classType:andReturnError:
func (_MLTreeEnsembleClassifierClass MLTreeEnsembleClassifierClass) _convertStringClassVectorInt64ClassVectorDimensionsToClassLabelClassTypeAndReturnError(vector unsafe.Pointer, vector2 unsafe.Pointer, dimensions uint64, label []objectivec.IObject) (int64, error) {
	var type_ int64
	var errorPtr objc.ID
	rv := objc.Send[bool](objc.ID(_MLTreeEnsembleClassifierClass.class), objc.Sel("_convertStringClassVector:int64ClassVector:dimensions:toClassLabel:classType:andReturnError:"), vector, vector2, dimensions, label, unsafe.Pointer(&type_), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return 0, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return 0, errors.New("_convertStringClassVector:int64ClassVector:dimensions:toClassLabel:classType:andReturnError: returned NO with nil NSError")
	}
	return type_, nil
}

// ConvertStringClassVectorInt64ClassVectorDimensionsToClassLabelClassTypeAndReturnError is an exported wrapper for the private method _convertStringClassVectorInt64ClassVectorDimensionsToClassLabelClassTypeAndReturnError.
func (_MLTreeEnsembleClassifierClass MLTreeEnsembleClassifierClass) ConvertStringClassVectorInt64ClassVectorDimensionsToClassLabelClassTypeAndReturnError(vector unsafe.Pointer, vector2 unsafe.Pointer, dimensions uint64, label []objectivec.IObject) (int64, error) {
	return _MLTreeEnsembleClassifierClass._convertStringClassVectorInt64ClassVectorDimensionsToClassLabelClassTypeAndReturnError(vector, vector2, dimensions, label)
}
//
// See: https://developer.apple.com/documentation/CoreML/MLTreeEnsembleClassifier/compileSpecification:toArchive:options:error:
func (_MLTreeEnsembleClassifierClass MLTreeEnsembleClassifierClass) CompileSpecificationToArchiveOptionsError(specification unsafe.Pointer, archive unsafe.Pointer, options objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLTreeEnsembleClassifierClass.class), objc.Sel("compileSpecification:toArchive:options:error:"), specification, archive, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLTreeEnsembleClassifier/compiledVersionForSpecification:options:error:
func (_MLTreeEnsembleClassifierClass MLTreeEnsembleClassifierClass) CompiledVersionForSpecificationOptionsError(specification unsafe.Pointer, options objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLTreeEnsembleClassifierClass.class), objc.Sel("compiledVersionForSpecification:options:error:"), specification, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLTreeEnsembleClassifier/loadModelFromCompiledArchive:modelVersionInfo:compilerVersionInfo:configuration:error:
func (_MLTreeEnsembleClassifierClass MLTreeEnsembleClassifierClass) LoadModelFromCompiledArchiveModelVersionInfoCompilerVersionInfoConfigurationError(archive unsafe.Pointer, info objectivec.IObject, info2 objectivec.IObject, configuration objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLTreeEnsembleClassifierClass.class), objc.Sel("loadModelFromCompiledArchive:modelVersionInfo:compilerVersionInfo:configuration:error:"), archive, info, info2, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLTreeEnsembleClassifier/loadModelFromSpecification:configuration:error:
func (_MLTreeEnsembleClassifierClass MLTreeEnsembleClassifierClass) LoadModelFromSpecificationConfigurationError(specification unsafe.Pointer, configuration objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLTreeEnsembleClassifierClass.class), objc.Sel("loadModelFromSpecification:configuration:error:"), specification, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLTreeEnsembleClassifier/loadModelFromSpecificationWithCompilationOptions:options:error:
func (_MLTreeEnsembleClassifierClass MLTreeEnsembleClassifierClass) LoadModelFromSpecificationWithCompilationOptionsOptionsError(options unsafe.Pointer, options2 objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLTreeEnsembleClassifierClass.class), objc.Sel("loadModelFromSpecificationWithCompilationOptions:options:error:"), options, options2, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// See: https://developer.apple.com/documentation/CoreML/MLTreeEnsembleClassifier/debugDescription
func (t MLTreeEnsembleClassifier) DebugDescription() string {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/CoreML/MLTreeEnsembleClassifier/description
func (t MLTreeEnsembleClassifier) Description() string {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/CoreML/MLTreeEnsembleClassifier/hash
func (t MLTreeEnsembleClassifier) Hash() uint64 {
	rv := objc.Send[uint64](t.ID, objc.Sel("hash"))
	return rv
}
// See: https://developer.apple.com/documentation/CoreML/MLTreeEnsembleClassifier/superclass
func (t MLTreeEnsembleClassifier) Superclass() objc.Class {
	rv := objc.Send[objc.Class](t.ID, objc.Sel("superclass"))
	return rv
}

