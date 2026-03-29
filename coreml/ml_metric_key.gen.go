// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"github.com/tmc/apple/objc"
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

// A key for the metrics dictionary in an update context.
//
// See: https://developer.apple.com/documentation/CoreML/MLMetricKey
type MLMetricKey struct {
	MLKey
}

// MLMetricKeyFromID constructs a [MLMetricKey] from an objc.ID.
//
// A key for the metrics dictionary in an update context.
func MLMetricKeyFromID(id objc.ID) MLMetricKey {
	return MLMetricKey{MLKey: MLKeyFromID(id)}
}
// NOTE: MLMetricKey adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MLMetricKey] class.
//
// See: https://developer.apple.com/documentation/CoreML/MLMetricKey
type IMLMetricKey interface {
	IMLKey

	// The training metrics of the model for the update task, contained in a dictionary.
	Metrics() IMLMetricKey
	SetMetrics(value IMLMetricKey)
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

// The training metrics of the model for the update task, contained in a
// dictionary.
//
// See: https://developer.apple.com/documentation/coreml/mlupdatecontext/metrics
func (m MLMetricKey) Metrics() IMLMetricKey {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("metrics"))
	return MLMetricKeyFromID(objc.ID(rv))
}
func (m MLMetricKey) SetMetrics(value IMLMetricKey) {
	objc.Send[struct{}](m.ID, objc.Sel("setMetrics:"), value)
}

// The key you use to access the current loss (a `float` value).
//
// # Discussion
// 
// Use this key to fetch the loss value in the [Metrics] dictionary.
//
// See: https://developer.apple.com/documentation/CoreML/MLMetricKey/lossValue
func (_MLMetricKeyClass MLMetricKeyClass) LossValue() MLMetricKey {
	rv := objc.Send[objc.ID](objc.ID(_MLMetricKeyClass.class), objc.Sel("lossValue"))
	return MLMetricKeyFromID(objc.ID(rv))
}
// The key you use to access the epoch index (an [Int64] value).
//
// # Discussion
// 
// Use this key to fetch the epoch index value in the [Metrics] dictionary.
//
// See: https://developer.apple.com/documentation/CoreML/MLMetricKey/epochIndex
func (_MLMetricKeyClass MLMetricKeyClass) EpochIndex() MLMetricKey {
	rv := objc.Send[objc.ID](objc.ID(_MLMetricKeyClass.class), objc.Sel("epochIndex"))
	return MLMetricKeyFromID(objc.ID(rv))
}
// The key you use to access the mini-batch index (an [Int64] value) within an
// epoch.
//
// # Discussion
// 
// Use this key to fetch the mini-batch index value in the [Metrics]
// dictionary.
//
// See: https://developer.apple.com/documentation/CoreML/MLMetricKey/miniBatchIndex
func (_MLMetricKeyClass MLMetricKeyClass) MiniBatchIndex() MLMetricKey {
	rv := objc.Send[objc.ID](objc.ID(_MLMetricKeyClass.class), objc.Sel("miniBatchIndex"))
	return MLMetricKeyFromID(objc.ID(rv))
}

