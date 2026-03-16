// Code generated from Apple documentation for appleneuralengine. DO NOT EDIT.

package appleneuralengine

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [ANECompiler] class.
var (
	_ANECompilerClass     ANECompilerClass
	_ANECompilerClassOnce sync.Once
)

func getANECompilerClass() ANECompilerClass {
	_ANECompilerClassOnce.Do(func() {
		_ANECompilerClass = ANECompilerClass{class: objc.GetClass("_ANECompiler")}
	})
	return _ANECompilerClass
}

// GetANECompilerClass returns the class object for _ANECompiler.
func GetANECompilerClass() ANECompilerClass {
	return getANECompilerClass()
}

type ANECompilerClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (ac ANECompilerClass) Alloc() ANECompiler {
	rv := objc.Send[ANECompiler](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANECompiler
type ANECompiler struct {
	objectivec.Object
}

// ANECompilerFromID constructs a [ANECompiler] from an objc.ID.
func ANECompilerFromID(id objc.ID) ANECompiler {
	return ANECompiler{objectivec.Object{id}}
}
// Ensure ANECompiler implements IANECompiler.
var _ IANECompiler = ANECompiler{}

// An interface definition for the [ANECompiler] class.
//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANECompiler
type IANECompiler interface {
	objectivec.IObject
}

// Init initializes the instance.
func (a ANECompiler) Init() ANECompiler {
	rv := objc.Send[ANECompiler](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a ANECompiler) Autorelease() ANECompiler {
	rv := objc.Send[ANECompiler](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewANECompiler creates a new ANECompiler instance.
func NewANECompiler() ANECompiler {
	class := getANECompilerClass()
	rv := objc.Send[ANECompiler](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANECompiler/compileModel:options:ok:error:
func (_ANECompilerClass ANECompilerClass) CompileModelOptionsOkError(model objectivec.IObject, options objectivec.IObject, ok unsafe.Pointer) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_ANECompilerClass.class), objc.Sel("compileModel:options:ok:error:"), model, options, ok, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANECompiler/compileModelJIT:ok:error:
func (_ANECompilerClass ANECompilerClass) CompileModelJITOkError(jit objectivec.IObject, ok unsafe.Pointer) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_ANECompilerClass.class), objc.Sel("compileModelJIT:ok:error:"), jit, ok, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANECompiler/createErrorWithUnderlyingError:
func (_ANECompilerClass ANECompilerClass) CreateErrorWithUnderlyingError(error_ objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANECompilerClass.class), objc.Sel("createErrorWithUnderlyingError:"), error_)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANECompiler/createInMemoryConstants:
func (_ANECompilerClass ANECompilerClass) CreateInMemoryConstants(constants objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANECompilerClass.class), objc.Sel("createInMemoryConstants:"), constants)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANECompiler/createJITNetworkFromModelAtPath:modelFilename:aotModelAtPath:aotModelFilename:
func (_ANECompilerClass ANECompilerClass) CreateJITNetworkFromModelAtPathModelFilenameAotModelAtPathAotModelFilename(path objectivec.IObject, filename objectivec.IObject, path2 objectivec.IObject, filename2 objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANECompilerClass.class), objc.Sel("createJITNetworkFromModelAtPath:modelFilename:aotModelAtPath:aotModelFilename:"), path, filename, path2, filename2)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANECompiler/createNetworkFromModelAtPath:modelFilename:
func (_ANECompilerClass ANECompilerClass) CreateNetworkFromModelAtPathModelFilename(path objectivec.IObject, filename objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANECompilerClass.class), objc.Sel("createNetworkFromModelAtPath:modelFilename:"), path, filename)
	return objectivec.Object{ID: rv}
}

