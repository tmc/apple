// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLMetricKey] class.
var (
	_MLMetricKeyClass     MLMetricKeyClass
	_MLMetricKeyClassOnce sync.Once
)

func getMLMetricKeyClass() MLMetricKeyClass {
	_MLMetricKeyClassOnce.Do(func() {
		_MLMetricKeyClass = MLMetricKeyClass{class: objc.GetClass("MLMetricKey")}
	})
	return _MLMetricKeyClass
}

// GetMLMetricKeyClass returns the class object for MLMetricKey.
func GetMLMetricKeyClass() MLMetricKeyClass {
	return getMLMetricKeyClass()
}

type MLMetricKeyClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLMetricKeyClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLMetricKeyClass) Alloc() MLMetricKey {
	rv := objc.Send[MLMetricKey](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLMetricKey
type MLMetricKey struct {
	MLKey
}

// MLMetricKeyFromID constructs a [MLMetricKey] from an objc.ID.
func MLMetricKeyFromID(id objc.ID) MLMetricKey {
	return MLMetricKey{MLKey: MLKeyFromID(id)}
}
// Ensure MLMetricKey implements IMLMetricKey.
var _ IMLMetricKey = MLMetricKey{}

// An interface definition for the [MLMetricKey] class.
//
// See: https://developer.apple.com/documentation/CoreML/MLMetricKey
type IMLMetricKey interface {
	IMLKey
}

// Init initializes the instance.
func (m MLMetricKey) Init() MLMetricKey {
	rv := objc.Send[MLMetricKey](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLMetricKey) Autorelease() MLMetricKey {
	rv := objc.Send[MLMetricKey](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLMetricKey creates a new MLMetricKey instance.
func NewMLMetricKey() MLMetricKey {
	class := getMLMetricKeyClass()
	rv := objc.Send[MLMetricKey](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/CoreML/MLKey/initWithCoder:
func NewMetricKeyWithCoder(coder objectivec.IObject) MLMetricKey {
	instance := getMLMetricKeyClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return MLMetricKeyFromID(rv)
}

//
// See: https://developer.apple.com/documentation/CoreML/MLMetricKey/initWithKeyName:
func NewMetricKeyWithKeyName(name objectivec.IObject) MLMetricKey {
	instance := getMLMetricKeyClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithKeyName:"), name)
	return MLMetricKeyFromID(rv)
}

//
// See: https://developer.apple.com/documentation/CoreML/MLKey/initWithKeyName:scope:
func NewMetricKeyWithKeyNameScope(name objectivec.IObject, scope objectivec.IObject) MLMetricKey {
	instance := getMLMetricKeyClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithKeyName:scope:"), name, scope)
	return MLMetricKeyFromID(rv)
}

