// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLPredictionOptions] class.
var (
	_MLPredictionOptionsClass     MLPredictionOptionsClass
	_MLPredictionOptionsClassOnce sync.Once
)

func getMLPredictionOptionsClass() MLPredictionOptionsClass {
	_MLPredictionOptionsClassOnce.Do(func() {
		_MLPredictionOptionsClass = MLPredictionOptionsClass{class: objc.GetClass("MLPredictionOptions")}
	})
	return _MLPredictionOptionsClass
}

// GetMLPredictionOptionsClass returns the class object for MLPredictionOptions.
func GetMLPredictionOptionsClass() MLPredictionOptionsClass {
	return getMLPredictionOptionsClass()
}

type MLPredictionOptionsClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLPredictionOptionsClass) Alloc() MLPredictionOptions {
	rv := objc.Send[MLPredictionOptions](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}







// The options available when making a prediction.
//
// # Getting features
//
//   - [MLPredictionOptions.OutputBackings]: A dictionary of feature names and client-allocated buffers.
//   - [MLPredictionOptions.SetOutputBackings]
//
// See: https://developer.apple.com/documentation/CoreML/MLPredictionOptions
type MLPredictionOptions struct {
	objectivec.Object
}

// MLPredictionOptionsFromID constructs a [MLPredictionOptions] from an objc.ID.
//
// The options available when making a prediction.
func MLPredictionOptionsFromID(id objc.ID) MLPredictionOptions {
	return MLPredictionOptions{objectivec.Object{id}}
}
// NOTE: MLPredictionOptions adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [MLPredictionOptions] class.
//
// # Getting features
//
//   - [IMLPredictionOptions.OutputBackings]: A dictionary of feature names and client-allocated buffers.
//   - [IMLPredictionOptions.SetOutputBackings]
//
// See: https://developer.apple.com/documentation/CoreML/MLPredictionOptions
type IMLPredictionOptions interface {
	objectivec.IObject

	// Topic: Getting features

	// A dictionary of feature names and client-allocated buffers.
	OutputBackings() foundation.INSDictionary
	SetOutputBackings(value foundation.INSDictionary)
}





// Init initializes the instance.
func (p MLPredictionOptions) Init() MLPredictionOptions {
	rv := objc.Send[MLPredictionOptions](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p MLPredictionOptions) Autorelease() MLPredictionOptions {
	rv := objc.Send[MLPredictionOptions](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLPredictionOptions creates a new MLPredictionOptions instance.
func NewMLPredictionOptions() MLPredictionOptions {
	class := getMLPredictionOptionsClass()
	rv := objc.Send[MLPredictionOptions](objc.ID(class.class), objc.Sel("new"))
	return rv
}





















// A dictionary of feature names and client-allocated buffers.
//
// See: https://developer.apple.com/documentation/CoreML/MLPredictionOptions/outputBackings
func (p MLPredictionOptions) OutputBackings() foundation.INSDictionary {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("outputBackings"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (p MLPredictionOptions) SetOutputBackings(value foundation.INSDictionary) {
	objc.Send[struct{}](p.ID, objc.Sel("setOutputBackings:"), value)
}























