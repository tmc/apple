// Code generated from Apple documentation for mlruntime. DO NOT EDIT.

package mlruntime

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLRTaskParameters] class.
var (
	_MLRTaskParametersClass     MLRTaskParametersClass
	_MLRTaskParametersClassOnce sync.Once
)

func getMLRTaskParametersClass() MLRTaskParametersClass {
	_MLRTaskParametersClassOnce.Do(func() {
		_MLRTaskParametersClass = MLRTaskParametersClass{class: objc.GetClass("MLRTaskParameters")}
	})
	return _MLRTaskParametersClass
}

// GetMLRTaskParametersClass returns the class object for MLRTaskParameters.
func GetMLRTaskParametersClass() MLRTaskParametersClass {
	return getMLRTaskParametersClass()
}

type MLRTaskParametersClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLRTaskParametersClass) Alloc() MLRTaskParameters {
	rv := objc.Send[MLRTaskParameters](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [MLRTaskParameters.BoolValueForKeyDefaultValue]
//   - [MLRTaskParameters.Count]
//   - [MLRTaskParameters.DictionaryRepresentation]
//   - [MLRTaskParameters.DoubleValueForKeyDefaultValue]
//   - [MLRTaskParameters.EncodeWithCoder]
//   - [MLRTaskParameters.FloatValueForKeyDefaultValue]
//   - [MLRTaskParameters.IntegerValueForKeyDefaultValue]
//   - [MLRTaskParameters.ObjectForKeyedSubscript]
//   - [MLRTaskParameters.RecipeUserInfo]
//   - [MLRTaskParameters.StringValueForKeyDefaultValue]
//   - [MLRTaskParameters.UnsignedIntegerValueForKeyDefaultValue]
//   - [MLRTaskParameters.InitWithCoder]
//   - [MLRTaskParameters.InitWithDESRecipe]
//   - [MLRTaskParameters.InitWithParametersDict]
//   - [MLRTaskParameters.InitWithURLError]
// See: https://developer.apple.com/documentation/MLRuntime/MLRTaskParameters
type MLRTaskParameters struct {
	objectivec.Object
}

// MLRTaskParametersFromID constructs a [MLRTaskParameters] from an objc.ID.
func MLRTaskParametersFromID(id objc.ID) MLRTaskParameters {
	return MLRTaskParameters{objectivec.Object{ID: id}}
}
// Ensure MLRTaskParameters implements IMLRTaskParameters.
var _ IMLRTaskParameters = MLRTaskParameters{}

// An interface definition for the [MLRTaskParameters] class.
//
// # Methods
//
//   - [IMLRTaskParameters.BoolValueForKeyDefaultValue]
//   - [IMLRTaskParameters.Count]
//   - [IMLRTaskParameters.DictionaryRepresentation]
//   - [IMLRTaskParameters.DoubleValueForKeyDefaultValue]
//   - [IMLRTaskParameters.EncodeWithCoder]
//   - [IMLRTaskParameters.FloatValueForKeyDefaultValue]
//   - [IMLRTaskParameters.IntegerValueForKeyDefaultValue]
//   - [IMLRTaskParameters.ObjectForKeyedSubscript]
//   - [IMLRTaskParameters.RecipeUserInfo]
//   - [IMLRTaskParameters.StringValueForKeyDefaultValue]
//   - [IMLRTaskParameters.UnsignedIntegerValueForKeyDefaultValue]
//   - [IMLRTaskParameters.InitWithCoder]
//   - [IMLRTaskParameters.InitWithDESRecipe]
//   - [IMLRTaskParameters.InitWithParametersDict]
//   - [IMLRTaskParameters.InitWithURLError]
//
// See: https://developer.apple.com/documentation/MLRuntime/MLRTaskParameters
type IMLRTaskParameters interface {
	objectivec.IObject

	// Topic: Methods

	BoolValueForKeyDefaultValue(key objectivec.IObject, value bool) bool
	Count() uint64
	DictionaryRepresentation() foundation.INSDictionary
	DoubleValueForKeyDefaultValue(key objectivec.IObject, value float64) float64
	EncodeWithCoder(coder foundation.INSCoder)
	FloatValueForKeyDefaultValue(key objectivec.IObject, value float32) float32
	IntegerValueForKeyDefaultValue(key objectivec.IObject, value int64) int64
	ObjectForKeyedSubscript(subscript objectivec.IObject) objectivec.IObject
	RecipeUserInfo() foundation.INSDictionary
	StringValueForKeyDefaultValue(key objectivec.IObject, value objectivec.IObject) objectivec.IObject
	UnsignedIntegerValueForKeyDefaultValue(key objectivec.IObject, value uint64) uint64
	InitWithCoder(coder foundation.INSCoder) MLRTaskParameters
	InitWithDESRecipe(dESRecipe objectivec.IObject) MLRTaskParameters
	InitWithParametersDict(dict objectivec.IObject) MLRTaskParameters
	InitWithURLError(url foundation.INSURL) (MLRTaskParameters, error)
}

// Init initializes the instance.
func (r MLRTaskParameters) Init() MLRTaskParameters {
	rv := objc.Send[MLRTaskParameters](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r MLRTaskParameters) Autorelease() MLRTaskParameters {
	rv := objc.Send[MLRTaskParameters](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLRTaskParameters creates a new MLRTaskParameters instance.
func NewMLRTaskParameters() MLRTaskParameters {
	class := getMLRTaskParametersClass()
	rv := objc.Send[MLRTaskParameters](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/MLRuntime/MLRTaskParameters/initWithCoder:
func NewRTaskParametersWithCoder(coder objectivec.IObject) MLRTaskParameters {
	instance := getMLRTaskParametersClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return MLRTaskParametersFromID(rv)
}

//
// See: https://developer.apple.com/documentation/MLRuntime/MLRTaskParameters/initWithDESRecipe:
func NewRTaskParametersWithDESRecipe(dESRecipe objectivec.IObject) MLRTaskParameters {
	instance := getMLRTaskParametersClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDESRecipe:"), dESRecipe)
	return MLRTaskParametersFromID(rv)
}

//
// See: https://developer.apple.com/documentation/MLRuntime/MLRTaskParameters/initWithParametersDict:
func NewRTaskParametersWithParametersDict(dict objectivec.IObject) MLRTaskParameters {
	instance := getMLRTaskParametersClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithParametersDict:"), dict)
	return MLRTaskParametersFromID(rv)
}

//
// See: https://developer.apple.com/documentation/MLRuntime/MLRTaskParameters/initWithURL:error:
func NewRTaskParametersWithURLError(url foundation.INSURL) (MLRTaskParameters, error) {
	var errorPtr objc.ID
	instance := getMLRTaskParametersClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithURL:error:"), url, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLRTaskParametersFromID(rv), foundation.NSErrorFrom(errorPtr)
	}
	return MLRTaskParametersFromID(rv), nil
}

//
// See: https://developer.apple.com/documentation/MLRuntime/MLRTaskParameters/boolValueForKey:defaultValue:
func (r MLRTaskParameters) BoolValueForKeyDefaultValue(key objectivec.IObject, value bool) bool {
	rv := objc.Send[bool](r.ID, objc.Sel("boolValueForKey:defaultValue:"), key, value)
	return rv
}
//
// See: https://developer.apple.com/documentation/MLRuntime/MLRTaskParameters/doubleValueForKey:defaultValue:
func (r MLRTaskParameters) DoubleValueForKeyDefaultValue(key objectivec.IObject, value float64) float64 {
	rv := objc.Send[float64](r.ID, objc.Sel("doubleValueForKey:defaultValue:"), key, value)
	return rv
}
//
// See: https://developer.apple.com/documentation/MLRuntime/MLRTaskParameters/encodeWithCoder:
func (r MLRTaskParameters) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](r.ID, objc.Sel("encodeWithCoder:"), coder)
}
//
// See: https://developer.apple.com/documentation/MLRuntime/MLRTaskParameters/floatValueForKey:defaultValue:
func (r MLRTaskParameters) FloatValueForKeyDefaultValue(key objectivec.IObject, value float32) float32 {
	rv := objc.Send[float32](r.ID, objc.Sel("floatValueForKey:defaultValue:"), key, value)
	return rv
}
//
// See: https://developer.apple.com/documentation/MLRuntime/MLRTaskParameters/integerValueForKey:defaultValue:
func (r MLRTaskParameters) IntegerValueForKeyDefaultValue(key objectivec.IObject, value int64) int64 {
	rv := objc.Send[int64](r.ID, objc.Sel("integerValueForKey:defaultValue:"), key, value)
	return rv
}
//
// See: https://developer.apple.com/documentation/MLRuntime/MLRTaskParameters/objectForKeyedSubscript:
func (r MLRTaskParameters) ObjectForKeyedSubscript(subscript objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("objectForKeyedSubscript:"), subscript)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/MLRuntime/MLRTaskParameters/stringValueForKey:defaultValue:
func (r MLRTaskParameters) StringValueForKeyDefaultValue(key objectivec.IObject, value objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("stringValueForKey:defaultValue:"), key, value)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/MLRuntime/MLRTaskParameters/unsignedIntegerValueForKey:defaultValue:
func (r MLRTaskParameters) UnsignedIntegerValueForKeyDefaultValue(key objectivec.IObject, value uint64) uint64 {
	rv := objc.Send[uint64](r.ID, objc.Sel("unsignedIntegerValueForKey:defaultValue:"), key, value)
	return rv
}
//
// See: https://developer.apple.com/documentation/MLRuntime/MLRTaskParameters/initWithCoder:
func (r MLRTaskParameters) InitWithCoder(coder foundation.INSCoder) MLRTaskParameters {
	rv := objc.Send[MLRTaskParameters](r.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}
//
// See: https://developer.apple.com/documentation/MLRuntime/MLRTaskParameters/initWithDESRecipe:
func (r MLRTaskParameters) InitWithDESRecipe(dESRecipe objectivec.IObject) MLRTaskParameters {
	rv := objc.Send[MLRTaskParameters](r.ID, objc.Sel("initWithDESRecipe:"), dESRecipe)
	return rv
}
//
// See: https://developer.apple.com/documentation/MLRuntime/MLRTaskParameters/initWithParametersDict:
func (r MLRTaskParameters) InitWithParametersDict(dict objectivec.IObject) MLRTaskParameters {
	rv := objc.Send[MLRTaskParameters](r.ID, objc.Sel("initWithParametersDict:"), dict)
	return rv
}
//
// See: https://developer.apple.com/documentation/MLRuntime/MLRTaskParameters/initWithURL:error:
func (r MLRTaskParameters) InitWithURLError(url foundation.INSURL) (MLRTaskParameters, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](r.ID, objc.Sel("initWithURL:error:"), url, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLRTaskParameters{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLRTaskParametersFromID(rv), nil

}

// See: https://developer.apple.com/documentation/MLRuntime/MLRTaskParameters/supportsSecureCoding
func (_MLRTaskParametersClass MLRTaskParametersClass) SupportsSecureCoding() bool {
	rv := objc.Send[bool](objc.ID(_MLRTaskParametersClass.class), objc.Sel("supportsSecureCoding"))
	return rv
}

// See: https://developer.apple.com/documentation/MLRuntime/MLRTaskParameters/count
func (r MLRTaskParameters) Count() uint64 {
	rv := objc.Send[uint64](r.ID, objc.Sel("count"))
	return rv
}
// See: https://developer.apple.com/documentation/MLRuntime/MLRTaskParameters/dictionaryRepresentation
func (r MLRTaskParameters) DictionaryRepresentation() foundation.INSDictionary {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("dictionaryRepresentation"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/MLRuntime/MLRTaskParameters/recipeUserInfo
func (r MLRTaskParameters) RecipeUserInfo() foundation.INSDictionary {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("recipeUserInfo"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}

