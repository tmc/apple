// Code generated from Apple documentation for mlruntime. DO NOT EDIT.

package mlruntime

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLRTrialDediscoRecipe] class.
var (
	_MLRTrialDediscoRecipeClass     MLRTrialDediscoRecipeClass
	_MLRTrialDediscoRecipeClassOnce sync.Once
)

func getMLRTrialDediscoRecipeClass() MLRTrialDediscoRecipeClass {
	_MLRTrialDediscoRecipeClassOnce.Do(func() {
		_MLRTrialDediscoRecipeClass = MLRTrialDediscoRecipeClass{class: objc.GetClass("MLRTrialDediscoRecipe")}
	})
	return _MLRTrialDediscoRecipeClass
}

// GetMLRTrialDediscoRecipeClass returns the class object for MLRTrialDediscoRecipe.
func GetMLRTrialDediscoRecipeClass() MLRTrialDediscoRecipeClass {
	return getMLRTrialDediscoRecipeClass()
}

type MLRTrialDediscoRecipeClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLRTrialDediscoRecipeClass) Alloc() MLRTrialDediscoRecipe {
	rv := objc.Send[MLRTrialDediscoRecipe](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [MLRTrialDediscoRecipe.BaseKeyFormat]
//   - [MLRTrialDediscoRecipe.DediscoTaskConfig]
//   - [MLRTrialDediscoRecipe.DpConfig]
//   - [MLRTrialDediscoRecipe.EncodingSchema]
//   - [MLRTrialDediscoRecipe.MlrDediscoMetadata]
//   - [MLRTrialDediscoRecipe.InitWithAssetURLConfigOverrideError]
// See: https://developer.apple.com/documentation/MLRuntime/MLRTrialDediscoRecipe
type MLRTrialDediscoRecipe struct {
	objectivec.Object
}

// MLRTrialDediscoRecipeFromID constructs a [MLRTrialDediscoRecipe] from an objc.ID.
func MLRTrialDediscoRecipeFromID(id objc.ID) MLRTrialDediscoRecipe {
	return MLRTrialDediscoRecipe{objectivec.Object{id}}
}
// Ensure MLRTrialDediscoRecipe implements IMLRTrialDediscoRecipe.
var _ IMLRTrialDediscoRecipe = MLRTrialDediscoRecipe{}

// An interface definition for the [MLRTrialDediscoRecipe] class.
//
// # Methods
//
//   - [IMLRTrialDediscoRecipe.BaseKeyFormat]
//   - [IMLRTrialDediscoRecipe.DediscoTaskConfig]
//   - [IMLRTrialDediscoRecipe.DpConfig]
//   - [IMLRTrialDediscoRecipe.EncodingSchema]
//   - [IMLRTrialDediscoRecipe.MlrDediscoMetadata]
//   - [IMLRTrialDediscoRecipe.InitWithAssetURLConfigOverrideError]
//
// See: https://developer.apple.com/documentation/MLRuntime/MLRTrialDediscoRecipe
type IMLRTrialDediscoRecipe interface {
	objectivec.IObject

	// Topic: Methods

	BaseKeyFormat() string
	DediscoTaskConfig() foundation.INSDictionary
	DpConfig() foundation.INSDictionary
	EncodingSchema() foundation.INSDictionary
	MlrDediscoMetadata() objectivec.IObject
	InitWithAssetURLConfigOverrideError(url foundation.INSURL, override objectivec.IObject) (MLRTrialDediscoRecipe, error)
}

// Init initializes the instance.
func (r MLRTrialDediscoRecipe) Init() MLRTrialDediscoRecipe {
	rv := objc.Send[MLRTrialDediscoRecipe](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r MLRTrialDediscoRecipe) Autorelease() MLRTrialDediscoRecipe {
	rv := objc.Send[MLRTrialDediscoRecipe](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLRTrialDediscoRecipe creates a new MLRTrialDediscoRecipe instance.
func NewMLRTrialDediscoRecipe() MLRTrialDediscoRecipe {
	class := getMLRTrialDediscoRecipeClass()
	rv := objc.Send[MLRTrialDediscoRecipe](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/MLRuntime/MLRTrialDediscoRecipe/initWithAssetURL:configOverride:error:
func NewRTrialDediscoRecipeWithAssetURLConfigOverrideError(url foundation.INSURL, override objectivec.IObject) (MLRTrialDediscoRecipe, error) {
	var errorPtr objc.ID
	instance := getMLRTrialDediscoRecipeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithAssetURL:configOverride:error:"), url, override, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLRTrialDediscoRecipeFromID(rv), foundation.NSErrorFrom(errorPtr)
	}
	return MLRTrialDediscoRecipeFromID(rv), nil
}

// See: https://developer.apple.com/documentation/MLRuntime/MLRTrialDediscoRecipe/mlrDediscoMetadata
func (r MLRTrialDediscoRecipe) MlrDediscoMetadata() objectivec.IObject {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("mlrDediscoMetadata"))
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/MLRuntime/MLRTrialDediscoRecipe/initWithAssetURL:configOverride:error:
func (r MLRTrialDediscoRecipe) InitWithAssetURLConfigOverrideError(url foundation.INSURL, override objectivec.IObject) (MLRTrialDediscoRecipe, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](r.ID, objc.Sel("initWithAssetURL:configOverride:error:"), url, override, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLRTrialDediscoRecipe{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLRTrialDediscoRecipeFromID(rv), nil

}

// See: https://developer.apple.com/documentation/MLRuntime/MLRTrialDediscoRecipe/baseKeyFormat
func (r MLRTrialDediscoRecipe) BaseKeyFormat() string {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("baseKeyFormat"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/MLRuntime/MLRTrialDediscoRecipe/dediscoTaskConfig
func (r MLRTrialDediscoRecipe) DediscoTaskConfig() foundation.INSDictionary {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("dediscoTaskConfig"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/MLRuntime/MLRTrialDediscoRecipe/dpConfig
func (r MLRTrialDediscoRecipe) DpConfig() foundation.INSDictionary {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("dpConfig"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/MLRuntime/MLRTrialDediscoRecipe/encodingSchema
func (r MLRTrialDediscoRecipe) EncodingSchema() foundation.INSDictionary {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("encodingSchema"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}

