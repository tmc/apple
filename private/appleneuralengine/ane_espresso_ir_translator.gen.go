// Code generated from Apple documentation for appleneuralengine. DO NOT EDIT.

package appleneuralengine

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"errors"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [ANEEspressoIRTranslator] class.
var (
	_ANEEspressoIRTranslatorClass     ANEEspressoIRTranslatorClass
	_ANEEspressoIRTranslatorClassOnce sync.Once
)

func getANEEspressoIRTranslatorClass() ANEEspressoIRTranslatorClass {
	_ANEEspressoIRTranslatorClassOnce.Do(func() {
		_ANEEspressoIRTranslatorClass = ANEEspressoIRTranslatorClass{class: objc.GetClass("_ANEEspressoIRTranslator")}
	})
	return _ANEEspressoIRTranslatorClass
}

// GetANEEspressoIRTranslatorClass returns the class object for _ANEEspressoIRTranslator.
func GetANEEspressoIRTranslatorClass() ANEEspressoIRTranslatorClass {
	return getANEEspressoIRTranslatorClass()
}

type ANEEspressoIRTranslatorClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (ac ANEEspressoIRTranslatorClass) Alloc() ANEEspressoIRTranslator {
	rv := objc.Send[ANEEspressoIRTranslator](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEEspressoIRTranslator
type ANEEspressoIRTranslator struct {
	objectivec.Object
}

// ANEEspressoIRTranslatorFromID constructs a [ANEEspressoIRTranslator] from an objc.ID.
func ANEEspressoIRTranslatorFromID(id objc.ID) ANEEspressoIRTranslator {
	return ANEEspressoIRTranslator{objectivec.Object{id}}
}
// Ensure ANEEspressoIRTranslator implements IANEEspressoIRTranslator.
var _ IANEEspressoIRTranslator = ANEEspressoIRTranslator{}





// An interface definition for the [ANEEspressoIRTranslator] class.
//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEEspressoIRTranslator
type IANEEspressoIRTranslator interface {
	objectivec.IObject
}





// Init initializes the instance.
func (a ANEEspressoIRTranslator) Init() ANEEspressoIRTranslator {
	rv := objc.Send[ANEEspressoIRTranslator](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a ANEEspressoIRTranslator) Autorelease() ANEEspressoIRTranslator {
	rv := objc.Send[ANEEspressoIRTranslator](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewANEEspressoIRTranslator creates a new ANEEspressoIRTranslator instance.
func NewANEEspressoIRTranslator() ANEEspressoIRTranslator {
	class := getANEEspressoIRTranslatorClass()
	rv := objc.Send[ANEEspressoIRTranslator](objc.ID(class.class), objc.Sel("new"))
	return rv
}














//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEEspressoIRTranslator/createErrorForPlan:status:
func (_ANEEspressoIRTranslatorClass ANEEspressoIRTranslatorClass) CreateErrorForPlanStatus(plan unsafe.Pointer, status int) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEEspressoIRTranslatorClass.class), objc.Sel("createErrorForPlan:status:"), plan, status)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEEspressoIRTranslator/destroyEspresso:ctx:
func (_ANEEspressoIRTranslatorClass ANEEspressoIRTranslatorClass) DestroyEspressoCtx(espresso unsafe.Pointer, ctx unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(_ANEEspressoIRTranslatorClass.class), objc.Sel("destroyEspresso:ctx:"), espresso, ctx)
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEEspressoIRTranslator/translateModelAt:key:outputPath:isEncryptedModel:translationOptions:error:
func (_ANEEspressoIRTranslatorClass ANEEspressoIRTranslatorClass) TranslateModelAtKeyOutputPathIsEncryptedModelTranslationOptionsError(at objectivec.IObject, key objectivec.IObject, path objectivec.IObject, model bool, options objectivec.IObject) (bool, error) {
			var errorPtr objc.ID
	rv := objc.Send[bool](objc.ID(_ANEEspressoIRTranslatorClass.class), objc.Sel("translateModelAt:key:outputPath:isEncryptedModel:translationOptions:error:"), at, key, path, model, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("translateModelAt:key:outputPath:isEncryptedModel:translationOptions:error: returned NO with nil NSError")
	}
	return rv, nil

}






















