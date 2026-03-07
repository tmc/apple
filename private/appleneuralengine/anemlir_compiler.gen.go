// Code generated from Apple documentation for appleneuralengine. DO NOT EDIT.

package appleneuralengine

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [ANEMLIRCompiler] class.
var (
	_ANEMLIRCompilerClass     ANEMLIRCompilerClass
	_ANEMLIRCompilerClassOnce sync.Once
)

func getANEMLIRCompilerClass() ANEMLIRCompilerClass {
	_ANEMLIRCompilerClassOnce.Do(func() {
		_ANEMLIRCompilerClass = ANEMLIRCompilerClass{class: objc.GetClass("_ANEMLIRCompiler")}
	})
	return _ANEMLIRCompilerClass
}

// GetANEMLIRCompilerClass returns the class object for _ANEMLIRCompiler.
func GetANEMLIRCompilerClass() ANEMLIRCompilerClass {
	return getANEMLIRCompilerClass()
}

type ANEMLIRCompilerClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (ac ANEMLIRCompilerClass) Alloc() ANEMLIRCompiler {
	rv := objc.Send[ANEMLIRCompiler](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEMLIRCompiler
type ANEMLIRCompiler struct {
	objectivec.Object
}

// ANEMLIRCompilerFromID constructs a [ANEMLIRCompiler] from an objc.ID.
func ANEMLIRCompilerFromID(id objc.ID) ANEMLIRCompiler {
	return ANEMLIRCompiler{objectivec.Object{id}}
}
// Ensure ANEMLIRCompiler implements IANEMLIRCompiler.
var _ IANEMLIRCompiler = ANEMLIRCompiler{}





// An interface definition for the [ANEMLIRCompiler] class.
//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEMLIRCompiler
type IANEMLIRCompiler interface {
	objectivec.IObject
}





// Init initializes the instance.
func (a ANEMLIRCompiler) Init() ANEMLIRCompiler {
	rv := objc.Send[ANEMLIRCompiler](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a ANEMLIRCompiler) Autorelease() ANEMLIRCompiler {
	rv := objc.Send[ANEMLIRCompiler](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewANEMLIRCompiler creates a new ANEMLIRCompiler instance.
func NewANEMLIRCompiler() ANEMLIRCompiler {
	class := getANEMLIRCompilerClass()
	rv := objc.Send[ANEMLIRCompiler](objc.ID(class.class), objc.Sel("new"))
	return rv
}














//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEMLIRCompiler/compileModelAt:modelName:csIdentity:optionsFilename:outputURL:saveSourceURL:aotModelBinaryPath:isEncryptedModel:options:mpsConstants:ok:error:
func (_ANEMLIRCompilerClass ANEMLIRCompilerClass) CompileModelAtModelNameCsIdentityOptionsFilenameOutputURLSaveSourceURLAotModelBinaryPathIsEncryptedModelOptionsMpsConstantsOkError(at objectivec.IObject, name objectivec.IObject, identity objectivec.IObject, filename objectivec.IObject, url foundation.INSURL, url2 foundation.INSURL, path objectivec.IObject, model bool, options objectivec.IObject, constants objectivec.IObject, ok unsafe.Pointer) (objectivec.IObject, error) {
			var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_ANEMLIRCompilerClass.class), objc.Sel("compileModelAt:modelName:csIdentity:optionsFilename:outputURL:saveSourceURL:aotModelBinaryPath:isEncryptedModel:options:mpsConstants:ok:error:"), at, name, identity, filename, url, url2, path, model, options, constants, ok, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}






















