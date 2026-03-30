// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLRegressor] class.
var (
	_MLRegressorClass     MLRegressorClass
	_MLRegressorClassOnce sync.Once
)

func getMLRegressorClass() MLRegressorClass {
	_MLRegressorClassOnce.Do(func() {
		_MLRegressorClass = MLRegressorClass{class: objc.GetClass("MLRegressor")}
	})
	return _MLRegressorClass
}

// GetMLRegressorClass returns the class object for MLRegressor.
func GetMLRegressorClass() MLRegressorClass {
	return getMLRegressorClass()
}

type MLRegressorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLRegressorClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLRegressorClass) Alloc() MLRegressor {
	rv := objc.Send[MLRegressor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A parent class referenced by other CoreML classes. [Full Topic]
type MLRegressor struct {
	objectivec.Object
}

// MLRegressorFromID constructs a [MLRegressor] from an objc.ID.
//
// A parent class referenced by other CoreML classes.
func MLRegressorFromID(id objc.ID) MLRegressor {
	return MLRegressor{objectivec.Object{ID: id}}
}

// Ensure MLRegressor implements IMLRegressor.
var _ IMLRegressor = MLRegressor{}

// An interface definition for the [MLRegressor] class.
type IMLRegressor interface {
	objectivec.IObject
}

// Init initializes the instance.
func (r MLRegressor) Init() MLRegressor {
	rv := objc.Send[MLRegressor](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r MLRegressor) Autorelease() MLRegressor {
	rv := objc.Send[MLRegressor](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLRegressor creates a new MLRegressor instance.
func NewMLRegressor() MLRegressor {
	class := getMLRegressorClass()
	rv := objc.Send[MLRegressor](objc.ID(class.class), objc.Sel("new"))
	return rv
}
