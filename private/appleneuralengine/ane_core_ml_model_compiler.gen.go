// Code generated from Apple documentation for appleneuralengine. DO NOT EDIT.

package appleneuralengine

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [ANECoreMLModelCompiler] class.
var (
	_ANECoreMLModelCompilerClass     ANECoreMLModelCompilerClass
	_ANECoreMLModelCompilerClassOnce sync.Once
)

func getANECoreMLModelCompilerClass() ANECoreMLModelCompilerClass {
	_ANECoreMLModelCompilerClassOnce.Do(func() {
		_ANECoreMLModelCompilerClass = ANECoreMLModelCompilerClass{class: objc.GetClass("_ANECoreMLModelCompiler")}
	})
	return _ANECoreMLModelCompilerClass
}

// GetANECoreMLModelCompilerClass returns the class object for _ANECoreMLModelCompiler.
func GetANECoreMLModelCompilerClass() ANECoreMLModelCompilerClass {
	return getANECoreMLModelCompilerClass()
}

type ANECoreMLModelCompilerClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (ac ANECoreMLModelCompilerClass) Alloc() ANECoreMLModelCompiler {
	rv := objc.Send[ANECoreMLModelCompiler](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANECoreMLModelCompiler
type ANECoreMLModelCompiler struct {
	objectivec.Object
}

// ANECoreMLModelCompilerFromID constructs a [ANECoreMLModelCompiler] from an objc.ID.
func ANECoreMLModelCompilerFromID(id objc.ID) ANECoreMLModelCompiler {
	return ANECoreMLModelCompiler{objectivec.Object{id}}
}
// Ensure ANECoreMLModelCompiler implements IANECoreMLModelCompiler.
var _ IANECoreMLModelCompiler = ANECoreMLModelCompiler{}





// An interface definition for the [ANECoreMLModelCompiler] class.
//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANECoreMLModelCompiler
type IANECoreMLModelCompiler interface {
	objectivec.IObject
}





// Init initializes the instance.
func (a ANECoreMLModelCompiler) Init() ANECoreMLModelCompiler {
	rv := objc.Send[ANECoreMLModelCompiler](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a ANECoreMLModelCompiler) Autorelease() ANECoreMLModelCompiler {
	rv := objc.Send[ANECoreMLModelCompiler](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewANECoreMLModelCompiler creates a new ANECoreMLModelCompiler instance.
func NewANECoreMLModelCompiler() ANECoreMLModelCompiler {
	class := getANECoreMLModelCompilerClass()
	rv := objc.Send[ANECoreMLModelCompiler](objc.ID(class.class), objc.Sel("new"))
	return rv
}














//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANECoreMLModelCompiler/compileModelAt:csIdentity:key:optionsFilename:tempDirectory:outputURL:saveSourceModelPath:aotModelBinaryPath:isEncryptedModel:options:ok:error:
func (_ANECoreMLModelCompilerClass ANECoreMLModelCompilerClass) CompileModelAtCsIdentityKeyOptionsFilenameTempDirectoryOutputURLSaveSourceModelPathAotModelBinaryPathIsEncryptedModelOptionsOkError(at objectivec.IObject, identity objectivec.IObject, key objectivec.IObject, filename objectivec.IObject, directory objectivec.IObject, url foundation.INSURL, path objectivec.IObject, path2 objectivec.IObject, model bool, options objectivec.IObject, ok unsafe.Pointer) (objectivec.IObject, error) {
			var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_ANECoreMLModelCompilerClass.class), objc.Sel("compileModelAt:csIdentity:key:optionsFilename:tempDirectory:outputURL:saveSourceModelPath:aotModelBinaryPath:isEncryptedModel:options:ok:error:"), at, identity, key, filename, directory, url, path, path2, model, options, ok, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANECoreMLModelCompiler/createErrorWithString:
func (_ANECoreMLModelCompilerClass ANECoreMLModelCompilerClass) CreateErrorWithString(string_ objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANECoreMLModelCompilerClass.class), objc.Sel("createErrorWithString:"), string_)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANECoreMLModelCompiler/pathsForModelURL:
func (_ANECoreMLModelCompilerClass ANECoreMLModelCompilerClass) PathsForModelURL(url foundation.INSURL) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANECoreMLModelCompilerClass.class), objc.Sel("pathsForModelURL:"), url)
	return objectivec.Object{ID: rv}
}






















