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

// The class instance for the [MLNeuralNetworkCompiler] class.
var (
	_MLNeuralNetworkCompilerClass     MLNeuralNetworkCompilerClass
	_MLNeuralNetworkCompilerClassOnce sync.Once
)

func getMLNeuralNetworkCompilerClass() MLNeuralNetworkCompilerClass {
	_MLNeuralNetworkCompilerClassOnce.Do(func() {
		_MLNeuralNetworkCompilerClass = MLNeuralNetworkCompilerClass{class: objc.GetClass("MLNeuralNetworkCompiler")}
	})
	return _MLNeuralNetworkCompilerClass
}

// GetMLNeuralNetworkCompilerClass returns the class object for MLNeuralNetworkCompiler.
func GetMLNeuralNetworkCompilerClass() MLNeuralNetworkCompilerClass {
	return getMLNeuralNetworkCompilerClass()
}

type MLNeuralNetworkCompilerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLNeuralNetworkCompilerClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLNeuralNetworkCompilerClass) Alloc() MLNeuralNetworkCompiler {
	rv := objc.Send[MLNeuralNetworkCompiler](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkCompiler
type MLNeuralNetworkCompiler struct {
	MLModel
}

// MLNeuralNetworkCompilerFromID constructs a [MLNeuralNetworkCompiler] from an objc.ID.
func MLNeuralNetworkCompilerFromID(id objc.ID) MLNeuralNetworkCompiler {
	return MLNeuralNetworkCompiler{MLModel: MLModelFromID(id)}
}
// Ensure MLNeuralNetworkCompiler implements IMLNeuralNetworkCompiler.
var _ IMLNeuralNetworkCompiler = MLNeuralNetworkCompiler{}

// An interface definition for the [MLNeuralNetworkCompiler] class.
//
// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkCompiler
type IMLNeuralNetworkCompiler interface {
	IMLModel
}

// Init initializes the instance.
func (n MLNeuralNetworkCompiler) Init() MLNeuralNetworkCompiler {
	rv := objc.Send[MLNeuralNetworkCompiler](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n MLNeuralNetworkCompiler) Autorelease() MLNeuralNetworkCompiler {
	rv := objc.Send[MLNeuralNetworkCompiler](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLNeuralNetworkCompiler creates a new MLNeuralNetworkCompiler instance.
func NewMLNeuralNetworkCompiler() MLNeuralNetworkCompiler {
	class := getMLNeuralNetworkCompilerClass()
	rv := objc.Send[MLNeuralNetworkCompiler](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/CoreML/MLModel/initDescriptionOnlyWithSpecification:configuration:error:
func NewNeuralNetworkCompilerDescriptionOnlyWithSpecificationConfigurationError(specification unsafe.Pointer, configuration objectivec.IObject) (MLNeuralNetworkCompiler, error) {
	var errorPtr objc.ID
	instance := getMLNeuralNetworkCompilerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initDescriptionOnlyWithSpecification:configuration:error:"), specification, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLNeuralNetworkCompiler{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLNeuralNetworkCompilerFromID(rv), nil
}

//
// See: https://developer.apple.com/documentation/CoreML/MLModel/initInterfaceAndMetadataWithCompiledArchive:error:
func NewNeuralNetworkCompilerInterfaceAndMetadataWithCompiledArchiveError(archive unsafe.Pointer) (MLNeuralNetworkCompiler, error) {
	var errorPtr objc.ID
	instance := getMLNeuralNetworkCompilerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initInterfaceAndMetadataWithCompiledArchive:error:"), archive, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLNeuralNetworkCompiler{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLNeuralNetworkCompilerFromID(rv), nil
}

//
// See: https://developer.apple.com/documentation/CoreML/MLModel/initWithConfiguration:
func NewNeuralNetworkCompilerWithConfiguration(configuration objectivec.IObject) MLNeuralNetworkCompiler {
	instance := getMLNeuralNetworkCompilerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithConfiguration:"), configuration)
	return MLNeuralNetworkCompilerFromID(rv)
}

//
// See: https://developer.apple.com/documentation/CoreML/MLModel/initWithDescription:
func NewNeuralNetworkCompilerWithDescription(description objectivec.IObject) MLNeuralNetworkCompiler {
	instance := getMLNeuralNetworkCompilerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDescription:"), description)
	return MLNeuralNetworkCompilerFromID(rv)
}

//
// See: https://developer.apple.com/documentation/CoreML/MLModel/initWithDescription:configuration:
func NewNeuralNetworkCompilerWithDescriptionConfiguration(description objectivec.IObject, configuration objectivec.IObject) MLNeuralNetworkCompiler {
	instance := getMLNeuralNetworkCompilerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDescription:configuration:"), description, configuration)
	return MLNeuralNetworkCompilerFromID(rv)
}

//
// See: https://developer.apple.com/documentation/CoreML/MLModel/initWithName:inputDescription:outputDescription:orderedInputFeatureNames:orderedOutputFeatureNames:configuration:
func NewNeuralNetworkCompilerWithNameInputDescriptionOutputDescriptionOrderedInputFeatureNamesOrderedOutputFeatureNamesConfiguration(name objectivec.IObject, description objectivec.IObject, description2 objectivec.IObject, names objectivec.IObject, names2 objectivec.IObject, configuration objectivec.IObject) MLNeuralNetworkCompiler {
	instance := getMLNeuralNetworkCompilerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithName:inputDescription:outputDescription:orderedInputFeatureNames:orderedOutputFeatureNames:configuration:"), name, description, description2, names, names2, configuration)
	return MLNeuralNetworkCompilerFromID(rv)
}

//
// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkCompiler/_compileSpecification:blobMapping:toArchive:options:error:
func (_MLNeuralNetworkCompilerClass MLNeuralNetworkCompilerClass) _compileSpecificationBlobMappingToArchiveOptionsError(specification unsafe.Pointer, mapping objectivec.IObject, archive unsafe.Pointer, options objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLNeuralNetworkCompilerClass.class), objc.Sel("_compileSpecification:blobMapping:toArchive:options:error:"), specification, mapping, archive, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkCompiler/collectEspressoModelDetails:modelPath:
func (_MLNeuralNetworkCompilerClass MLNeuralNetworkCompilerClass) CollectEspressoModelDetailsModelPath(details unsafe.Pointer, path unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(_MLNeuralNetworkCompilerClass.class), objc.Sel("collectEspressoModelDetails:modelPath:"), details, path)
}
//
// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkCompiler/collectNNModelDetailsFromArchive:spec:error:
func (_MLNeuralNetworkCompilerClass MLNeuralNetworkCompilerClass) CollectNNModelDetailsFromArchiveSpecError(archive unsafe.Pointer, spec unsafe.Pointer) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](objc.ID(_MLNeuralNetworkCompilerClass.class), objc.Sel("collectNNModelDetailsFromArchive:spec:error:"), archive, spec, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("collectNNModelDetailsFromArchive:spec:error: returned NO with nil NSError")
	}
	return rv, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkCompiler/compileSpecification:blobMapping:toArchive:options:error:
func (_MLNeuralNetworkCompilerClass MLNeuralNetworkCompilerClass) CompileSpecificationBlobMappingToArchiveOptionsError(specification unsafe.Pointer, mapping objectivec.IObject, archive unsafe.Pointer, options objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLNeuralNetworkCompilerClass.class), objc.Sel("compileSpecification:blobMapping:toArchive:options:error:"), specification, mapping, archive, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkCompiler/compileSpecification:toArchive:options:error:
func (_MLNeuralNetworkCompilerClass MLNeuralNetworkCompilerClass) CompileSpecificationToArchiveOptionsError(specification unsafe.Pointer, archive unsafe.Pointer, options objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLNeuralNetworkCompilerClass.class), objc.Sel("compileSpecification:toArchive:options:error:"), specification, archive, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkCompiler/compiledVersionForSpecification:options:error:
func (_MLNeuralNetworkCompilerClass MLNeuralNetworkCompilerClass) CompiledVersionForSpecificationOptionsError(specification unsafe.Pointer, options objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLNeuralNetworkCompilerClass.class), objc.Sel("compiledVersionForSpecification:options:error:"), specification, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkCompiler/iOS17CompilerVersionInfo
func (_MLNeuralNetworkCompilerClass MLNeuralNetworkCompilerClass) IOS17CompilerVersionInfo() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLNeuralNetworkCompilerClass.class), objc.Sel("iOS17CompilerVersionInfo"))
	return objectivec.Object{ID: rv}
}
// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkCompiler/iOS18CompilerVersionInfo
func (_MLNeuralNetworkCompilerClass MLNeuralNetworkCompilerClass) IOS18CompilerVersionInfo() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLNeuralNetworkCompilerClass.class), objc.Sel("iOS18CompilerVersionInfo"))
	return objectivec.Object{ID: rv}
}

