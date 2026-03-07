// Code generated from Apple documentation for appleneuralengine. DO NOT EDIT.

package appleneuralengine

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [ANEMILCompiler] class.
var (
	_ANEMILCompilerClass     ANEMILCompilerClass
	_ANEMILCompilerClassOnce sync.Once
)

func getANEMILCompilerClass() ANEMILCompilerClass {
	_ANEMILCompilerClassOnce.Do(func() {
		_ANEMILCompilerClass = ANEMILCompilerClass{class: objc.GetClass("_ANEMILCompiler")}
	})
	return _ANEMILCompilerClass
}

// GetANEMILCompilerClass returns the class object for _ANEMILCompiler.
func GetANEMILCompilerClass() ANEMILCompilerClass {
	return getANEMILCompilerClass()
}

type ANEMILCompilerClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (ac ANEMILCompilerClass) Alloc() ANEMILCompiler {
	rv := objc.Send[ANEMILCompiler](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEMILCompiler
type ANEMILCompiler struct {
	objectivec.Object
}

// ANEMILCompilerFromID constructs a [ANEMILCompiler] from an objc.ID.
func ANEMILCompilerFromID(id objc.ID) ANEMILCompiler {
	return ANEMILCompiler{objectivec.Object{id}}
}
// Ensure ANEMILCompiler implements IANEMILCompiler.
var _ IANEMILCompiler = ANEMILCompiler{}





// An interface definition for the [ANEMILCompiler] class.
//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEMILCompiler
type IANEMILCompiler interface {
	objectivec.IObject
}





// Init initializes the instance.
func (a ANEMILCompiler) Init() ANEMILCompiler {
	rv := objc.Send[ANEMILCompiler](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a ANEMILCompiler) Autorelease() ANEMILCompiler {
	rv := objc.Send[ANEMILCompiler](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewANEMILCompiler creates a new ANEMILCompiler instance.
func NewANEMILCompiler() ANEMILCompiler {
	class := getANEMILCompilerClass()
	rv := objc.Send[ANEMILCompiler](objc.ID(class.class), objc.Sel("new"))
	return rv
}














//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEMILCompiler/compileModelAt:modelName:csIdentity:optionsFilename:outputURL:saveSourceURL:aotModelBinaryPath:isEncryptedModel:options:ok:error:
func (_ANEMILCompilerClass ANEMILCompilerClass) CompileModelAtModelNameCsIdentityOptionsFilenameOutputURLSaveSourceURLAotModelBinaryPathIsEncryptedModelOptionsOkError(at objectivec.IObject, name objectivec.IObject, identity objectivec.IObject, filename objectivec.IObject, url foundation.INSURL, url2 foundation.INSURL, path objectivec.IObject, model bool, options objectivec.IObject, ok unsafe.Pointer) (objectivec.IObject, error) {
			var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_ANEMILCompilerClass.class), objc.Sel("compileModelAt:modelName:csIdentity:optionsFilename:outputURL:saveSourceURL:aotModelBinaryPath:isEncryptedModel:options:ok:error:"), at, name, identity, filename, url, url2, path, model, options, ok, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}






















