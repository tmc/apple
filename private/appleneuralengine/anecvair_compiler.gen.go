// Code generated from Apple documentation for appleneuralengine. DO NOT EDIT.

package appleneuralengine

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [ANECVAIRCompiler] class.
var (
	_ANECVAIRCompilerClass     ANECVAIRCompilerClass
	_ANECVAIRCompilerClassOnce sync.Once
)

func getANECVAIRCompilerClass() ANECVAIRCompilerClass {
	_ANECVAIRCompilerClassOnce.Do(func() {
		_ANECVAIRCompilerClass = ANECVAIRCompilerClass{class: objc.GetClass("_ANECVAIRCompiler")}
	})
	return _ANECVAIRCompilerClass
}

// GetANECVAIRCompilerClass returns the class object for _ANECVAIRCompiler.
func GetANECVAIRCompilerClass() ANECVAIRCompilerClass {
	return getANECVAIRCompilerClass()
}

type ANECVAIRCompilerClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (ac ANECVAIRCompilerClass) Alloc() ANECVAIRCompiler {
	rv := objc.Send[ANECVAIRCompiler](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANECVAIRCompiler
type ANECVAIRCompiler struct {
	objectivec.Object
}

// ANECVAIRCompilerFromID constructs a [ANECVAIRCompiler] from an objc.ID.
func ANECVAIRCompilerFromID(id objc.ID) ANECVAIRCompiler {
	return ANECVAIRCompiler{objectivec.Object{id}}
}
// Ensure ANECVAIRCompiler implements IANECVAIRCompiler.
var _ IANECVAIRCompiler = ANECVAIRCompiler{}





// An interface definition for the [ANECVAIRCompiler] class.
//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANECVAIRCompiler
type IANECVAIRCompiler interface {
	objectivec.IObject
}





// Init initializes the instance.
func (a ANECVAIRCompiler) Init() ANECVAIRCompiler {
	rv := objc.Send[ANECVAIRCompiler](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a ANECVAIRCompiler) Autorelease() ANECVAIRCompiler {
	rv := objc.Send[ANECVAIRCompiler](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewANECVAIRCompiler creates a new ANECVAIRCompiler instance.
func NewANECVAIRCompiler() ANECVAIRCompiler {
	class := getANECVAIRCompilerClass()
	rv := objc.Send[ANECVAIRCompiler](objc.ID(class.class), objc.Sel("new"))
	return rv
}














//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANECVAIRCompiler/compileModelAt:csIdentity:plistFilename:optionsFilename:outputURL:saveSourceURL:aotModelBinaryPath:isEncryptedModel:options:ok:error:
func (_ANECVAIRCompilerClass ANECVAIRCompilerClass) CompileModelAtCsIdentityPlistFilenameOptionsFilenameOutputURLSaveSourceURLAotModelBinaryPathIsEncryptedModelOptionsOkError(at objectivec.IObject, identity objectivec.IObject, filename objectivec.IObject, filename2 objectivec.IObject, url foundation.INSURL, url2 foundation.INSURL, path objectivec.IObject, model bool, options objectivec.IObject, ok unsafe.Pointer) (objectivec.IObject, error) {
			var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_ANECVAIRCompilerClass.class), objc.Sel("compileModelAt:csIdentity:plistFilename:optionsFilename:outputURL:saveSourceURL:aotModelBinaryPath:isEncryptedModel:options:ok:error:"), at, identity, filename, filename2, url, url2, path, model, options, ok, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}






















