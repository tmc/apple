// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLClassifier] class.
var (
	_MLClassifierClass     MLClassifierClass
	_MLClassifierClassOnce sync.Once
)

func getMLClassifierClass() MLClassifierClass {
	_MLClassifierClassOnce.Do(func() {
		_MLClassifierClass = MLClassifierClass{class: objc.GetClass("MLClassifier")}
	})
	return _MLClassifierClass
}

// GetMLClassifierClass returns the class object for MLClassifier.
func GetMLClassifierClass() MLClassifierClass {
	return getMLClassifierClass()
}

type MLClassifierClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLClassifierClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLClassifierClass) Alloc() MLClassifier {
	rv := objc.Send[MLClassifier](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A parent class referenced by other CoreML classes. [Full Topic]
type MLClassifier struct {
	objectivec.Object
}

// MLClassifierFromID constructs a [MLClassifier] from an objc.ID.
//
// A parent class referenced by other CoreML classes.
func MLClassifierFromID(id objc.ID) MLClassifier {
	return MLClassifier{objectivec.Object{ID: id}}
}

// Ensure MLClassifier implements IMLClassifier.
var _ IMLClassifier = MLClassifier{}

// An interface definition for the [MLClassifier] class.
type IMLClassifier interface {
	objectivec.IObject
}

// Init initializes the instance.
func (c MLClassifier) Init() MLClassifier {
	rv := objc.Send[MLClassifier](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MLClassifier) Autorelease() MLClassifier {
	rv := objc.Send[MLClassifier](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLClassifier creates a new MLClassifier instance.
func NewMLClassifier() MLClassifier {
	class := getMLClassifierClass()
	rv := objc.Send[MLClassifier](objc.ID(class.class), objc.Sel("new"))
	return rv
}
