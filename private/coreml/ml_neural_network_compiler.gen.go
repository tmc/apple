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
	rv := objc.SendIfResponds[MLNeuralNetworkCompiler](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

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
type IMLNeuralNetworkCompiler interface {
	IMLModel
}

// Init initializes the instance.
func (m MLNeuralNetworkCompiler) Init() MLNeuralNetworkCompiler {
	rv := objc.SendIfResponds[MLNeuralNetworkCompiler](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLNeuralNetworkCompiler) Autorelease() MLNeuralNetworkCompiler {
	rv := objc.SendIfResponds[MLNeuralNetworkCompiler](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLNeuralNetworkCompiler creates a new MLNeuralNetworkCompiler instance.
func NewMLNeuralNetworkCompiler() MLNeuralNetworkCompiler {
	class := getMLNeuralNetworkCompilerClass()
	rv := objc.SendIfResponds[MLNeuralNetworkCompiler](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewNeuralNetworkCompilerDescriptionOnlyWithSpecificationConfigurationError(specification unsafe.Pointer, configuration objectivec.IObject) (MLNeuralNetworkCompiler, error) {
	var errorPtr objc.ID
	instance := getMLNeuralNetworkCompilerClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initDescriptionOnlyWithSpecification:configuration:error:"), specification, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLNeuralNetworkCompiler{}, foundation.NSErrorFrom(errorPtr)
	}
	if rv == 0 {
		return MLNeuralNetworkCompiler{}, objc.ErrInitFailed
	}
	return MLNeuralNetworkCompilerFromID(rv), nil
}

func NewNeuralNetworkCompilerInterfaceAndMetadataWithCompiledArchiveError(archive unsafe.Pointer) (MLNeuralNetworkCompiler, error) {
	var errorPtr objc.ID
	instance := getMLNeuralNetworkCompilerClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initInterfaceAndMetadataWithCompiledArchive:error:"), archive, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLNeuralNetworkCompiler{}, foundation.NSErrorFrom(errorPtr)
	}
	if rv == 0 {
		return MLNeuralNetworkCompiler{}, objc.ErrInitFailed
	}
	return MLNeuralNetworkCompilerFromID(rv), nil
}

func NewNeuralNetworkCompilerWithConfiguration(configuration objectivec.IObject) MLNeuralNetworkCompiler {
	instance := getMLNeuralNetworkCompilerClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithConfiguration:"), configuration)
	return MLNeuralNetworkCompilerFromID(rv)
}

func NewNeuralNetworkCompilerWithDescription(description objectivec.IObject) MLNeuralNetworkCompiler {
	instance := getMLNeuralNetworkCompilerClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithDescription:"), description)
	return MLNeuralNetworkCompilerFromID(rv)
}

func NewNeuralNetworkCompilerWithDescriptionConfiguration(description objectivec.IObject, configuration objectivec.IObject) MLNeuralNetworkCompiler {
	instance := getMLNeuralNetworkCompilerClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithDescription:configuration:"), description, configuration)
	return MLNeuralNetworkCompilerFromID(rv)
}

func NewNeuralNetworkCompilerWithNameInputDescriptionOutputDescriptionOrderedInputFeatureNamesOrderedOutputFeatureNamesConfiguration(name objectivec.IObject, description objectivec.IObject, description2 objectivec.IObject, names objectivec.IObject, names2 objectivec.IObject, configuration objectivec.IObject) MLNeuralNetworkCompiler {
	instance := getMLNeuralNetworkCompilerClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithName:inputDescription:outputDescription:orderedInputFeatureNames:orderedOutputFeatureNames:configuration:"), name, description, description2, names, names2, configuration)
	return MLNeuralNetworkCompilerFromID(rv)
}

func (_MLNeuralNetworkCompilerClass MLNeuralNetworkCompilerClass) _compileSpecificationBlobMappingToArchiveOptionsError(specification unsafe.Pointer, mapping objectivec.IObject, archive unsafe.Pointer, options objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLNeuralNetworkCompilerClass.class), objc.Sel("_compileSpecification:blobMapping:toArchive:options:error:"), specification, mapping, archive, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (_MLNeuralNetworkCompilerClass MLNeuralNetworkCompilerClass) CollectEspressoModelDetailsModelPath(details unsafe.Pointer, path unsafe.Pointer) {
	objc.SendIfResponds[objc.ID](objc.ID(_MLNeuralNetworkCompilerClass.class), objc.Sel("collectEspressoModelDetails:modelPath:"), details, path)
}
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
func (_MLNeuralNetworkCompilerClass MLNeuralNetworkCompilerClass) CompileSpecificationBlobMappingToArchiveOptionsError(specification unsafe.Pointer, mapping objectivec.IObject, archive unsafe.Pointer, options objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLNeuralNetworkCompilerClass.class), objc.Sel("compileSpecification:blobMapping:toArchive:options:error:"), specification, mapping, archive, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (_MLNeuralNetworkCompilerClass MLNeuralNetworkCompilerClass) CompileSpecificationToArchiveOptionsError(specification unsafe.Pointer, archive unsafe.Pointer, options objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLNeuralNetworkCompilerClass.class), objc.Sel("compileSpecification:toArchive:options:error:"), specification, archive, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (_MLNeuralNetworkCompilerClass MLNeuralNetworkCompilerClass) CompiledVersionForSpecificationOptionsError(specification unsafe.Pointer, options objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLNeuralNetworkCompilerClass.class), objc.Sel("compiledVersionForSpecification:options:error:"), specification, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (_MLNeuralNetworkCompilerClass MLNeuralNetworkCompilerClass) IOS17CompilerVersionInfo() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_MLNeuralNetworkCompilerClass.class), objc.Sel("iOS17CompilerVersionInfo"))
	return objectivec.Object{ID: rv}
}
func (_MLNeuralNetworkCompilerClass MLNeuralNetworkCompilerClass) IOS18CompilerVersionInfo() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_MLNeuralNetworkCompilerClass.class), objc.Sel("iOS18CompilerVersionInfo"))
	return objectivec.Object{ID: rv}
}
