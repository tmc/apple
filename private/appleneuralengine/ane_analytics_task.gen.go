// Code generated from Apple documentation for appleneuralengine. DO NOT EDIT.

package appleneuralengine

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [ANEAnalyticsTask] class.
var (
	_ANEAnalyticsTaskClass     ANEAnalyticsTaskClass
	_ANEAnalyticsTaskClassOnce sync.Once
)

func getANEAnalyticsTaskClass() ANEAnalyticsTaskClass {
	_ANEAnalyticsTaskClassOnce.Do(func() {
		_ANEAnalyticsTaskClass = ANEAnalyticsTaskClass{class: objc.GetClass("_ANEAnalyticsTask")}
	})
	return _ANEAnalyticsTaskClass
}

// GetANEAnalyticsTaskClass returns the class object for _ANEAnalyticsTask.
func GetANEAnalyticsTaskClass() ANEAnalyticsTaskClass {
	return getANEAnalyticsTaskClass()
}

type ANEAnalyticsTaskClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (ac ANEAnalyticsTaskClass) Alloc() ANEAnalyticsTask {
	rv := objc.Send[ANEAnalyticsTask](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}







//
// # Methods
//
//   - [ANEAnalyticsTask.Metrics]
//   - [ANEAnalyticsTask.Serialize]
//   - [ANEAnalyticsTask.InitWithMetrics]
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEAnalyticsTask
type ANEAnalyticsTask struct {
	objectivec.Object
}

// ANEAnalyticsTaskFromID constructs a [ANEAnalyticsTask] from an objc.ID.
func ANEAnalyticsTaskFromID(id objc.ID) ANEAnalyticsTask {
	return ANEAnalyticsTask{objectivec.Object{id}}
}
// Ensure ANEAnalyticsTask implements IANEAnalyticsTask.
var _ IANEAnalyticsTask = ANEAnalyticsTask{}





// An interface definition for the [ANEAnalyticsTask] class.
//
// # Methods
//
//   - [IANEAnalyticsTask.Metrics]
//   - [IANEAnalyticsTask.Serialize]
//   - [IANEAnalyticsTask.InitWithMetrics]
//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEAnalyticsTask
type IANEAnalyticsTask interface {
	objectivec.IObject

	// Topic: Methods

	Metrics() foundation.INSDictionary
	Serialize() objectivec.IObject
	InitWithMetrics(metrics objectivec.IObject) ANEAnalyticsTask
}





// Init initializes the instance.
func (a ANEAnalyticsTask) Init() ANEAnalyticsTask {
	rv := objc.Send[ANEAnalyticsTask](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a ANEAnalyticsTask) Autorelease() ANEAnalyticsTask {
	rv := objc.Send[ANEAnalyticsTask](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewANEAnalyticsTask creates a new ANEAnalyticsTask instance.
func NewANEAnalyticsTask() ANEAnalyticsTask {
	class := getANEAnalyticsTaskClass()
	rv := objc.Send[ANEAnalyticsTask](objc.ID(class.class), objc.Sel("new"))
	return rv
}






//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEAnalyticsTask/initWithMetrics:
func NewANEAnalyticsTaskWithMetrics(metrics objectivec.IObject) ANEAnalyticsTask {
	instance := getANEAnalyticsTaskClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithMetrics:"), metrics)
	return ANEAnalyticsTaskFromID(rv)
}







// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEAnalyticsTask/serialize
func (a ANEAnalyticsTask) Serialize() objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("serialize"))
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEAnalyticsTask/initWithMetrics:
func (a ANEAnalyticsTask) InitWithMetrics(metrics objectivec.IObject) ANEAnalyticsTask {
	rv := objc.Send[ANEAnalyticsTask](a.ID, objc.Sel("initWithMetrics:"), metrics)
	return rv
}





//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEAnalyticsTask/objectWithMetrics:
func (_ANEAnalyticsTaskClass ANEAnalyticsTaskClass) ObjectWithMetrics(metrics objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEAnalyticsTaskClass.class), objc.Sel("objectWithMetrics:"), metrics)
	return objectivec.Object{ID: rv}
}








// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEAnalyticsTask/metrics
func (a ANEAnalyticsTask) Metrics() foundation.INSDictionary {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("metrics"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}

















