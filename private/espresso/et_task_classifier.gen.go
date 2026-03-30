// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [ETTaskClassifier] class.
var (
	_ETTaskClassifierClass     ETTaskClassifierClass
	_ETTaskClassifierClassOnce sync.Once
)

func getETTaskClassifierClass() ETTaskClassifierClass {
	_ETTaskClassifierClassOnce.Do(func() {
		_ETTaskClassifierClass = ETTaskClassifierClass{class: objc.GetClass("ETTaskClassifier")}
	})
	return _ETTaskClassifierClass
}

// GetETTaskClassifierClass returns the class object for ETTaskClassifier.
func GetETTaskClassifierClass() ETTaskClassifierClass {
	return getETTaskClassifierClass()
}

type ETTaskClassifierClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ec ETTaskClassifierClass) Class() objc.Class {
	return ec.class
}

// Alloc allocates memory for a new instance of the class.
func (ec ETTaskClassifierClass) Alloc() ETTaskClassifier {
	rv := objc.Send[ETTaskClassifier](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [ETTaskClassifier.Class_names]
//   - [ETTaskClassifier.SetClass_names]
//
// See: https://developer.apple.com/documentation/Espresso/ETTaskClassifier
type ETTaskClassifier struct {
	ETTask
}

// ETTaskClassifierFromID constructs a [ETTaskClassifier] from an objc.ID.
func ETTaskClassifierFromID(id objc.ID) ETTaskClassifier {
	return ETTaskClassifier{ETTask: ETTaskFromID(id)}
}

// Ensure ETTaskClassifier implements IETTaskClassifier.
var _ IETTaskClassifier = ETTaskClassifier{}

// An interface definition for the [ETTaskClassifier] class.
//
// # Methods
//
//   - [IETTaskClassifier.Class_names]
//   - [IETTaskClassifier.SetClass_names]
//
// See: https://developer.apple.com/documentation/Espresso/ETTaskClassifier
type IETTaskClassifier interface {
	IETTask

	// Topic: Methods

	Class_names() foundation.INSArray
	SetClass_names(value foundation.INSArray)
}

// Init initializes the instance.
func (e ETTaskClassifier) Init() ETTaskClassifier {
	rv := objc.Send[ETTaskClassifier](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e ETTaskClassifier) Autorelease() ETTaskClassifier {
	rv := objc.Send[ETTaskClassifier](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewETTaskClassifier creates a new ETTaskClassifier instance.
func NewETTaskClassifier() ETTaskClassifier {
	class := getETTaskClassifierClass()
	rv := objc.Send[ETTaskClassifier](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Espresso/ETTaskClassifier/initWithModelDef:optimizerDef:extractor:
func NewETTaskClassifierWithModelDefOptimizerDefExtractor(def objectivec.IObject, def2 objectivec.IObject, extractor objectivec.IObject) ETTaskClassifier {
	instance := getETTaskClassifierClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithModelDef:optimizerDef:extractor:"), def, def2, extractor)
	return ETTaskClassifierFromID(rv)
}

// See: https://developer.apple.com/documentation/Espresso/ETTaskClassifier/initWithModelDef:optimizerDef:extractor:needWeightsInitialization:
func NewETTaskClassifierWithModelDefOptimizerDefExtractorNeedWeightsInitialization(def objectivec.IObject, def2 objectivec.IObject, extractor objectivec.IObject, initialization bool) ETTaskClassifier {
	instance := getETTaskClassifierClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithModelDef:optimizerDef:extractor:needWeightsInitialization:"), def, def2, extractor, initialization)
	return ETTaskClassifierFromID(rv)
}

// See: https://developer.apple.com/documentation/Espresso/ETTask/initWithModelDef:optimizerDef:lossConfig:
func NewETTaskClassifierWithModelDefOptimizerDefLossConfig(def objectivec.IObject, def2 objectivec.IObject, config objectivec.IObject) ETTaskClassifier {
	instance := getETTaskClassifierClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithModelDef:optimizerDef:lossConfig:"), def, def2, config)
	return ETTaskClassifierFromID(rv)
}

// See: https://developer.apple.com/documentation/Espresso/ETTask/initWithModelDef:optimizerDef:lossConfig:extractor:
func NewETTaskClassifierWithModelDefOptimizerDefLossConfigExtractor(def objectivec.IObject, def2 objectivec.IObject, config objectivec.IObject, extractor objectivec.IObject) ETTaskClassifier {
	instance := getETTaskClassifierClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithModelDef:optimizerDef:lossConfig:extractor:"), def, def2, config, extractor)
	return ETTaskClassifierFromID(rv)
}

// See: https://developer.apple.com/documentation/Espresso/ETTaskClassifier/class_names
func (e ETTaskClassifier) Class_names() foundation.INSArray {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("class_names"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (e ETTaskClassifier) SetClass_names(value foundation.INSArray) {
	objc.Send[struct{}](e.ID, objc.Sel("setClass_names:"), value)
}
