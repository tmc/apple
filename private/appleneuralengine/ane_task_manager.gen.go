// Code generated from Apple documentation for appleneuralengine. DO NOT EDIT.

package appleneuralengine

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [ANETaskManager] class.
var (
	_ANETaskManagerClass     ANETaskManagerClass
	_ANETaskManagerClassOnce sync.Once
)

func getANETaskManagerClass() ANETaskManagerClass {
	_ANETaskManagerClassOnce.Do(func() {
		_ANETaskManagerClass = ANETaskManagerClass{class: objc.GetClass("_ANETaskManager")}
	})
	return _ANETaskManagerClass
}

// GetANETaskManagerClass returns the class object for _ANETaskManager.
func GetANETaskManagerClass() ANETaskManagerClass {
	return getANETaskManagerClass()
}

type ANETaskManagerClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (ac ANETaskManagerClass) Alloc() ANETaskManager {
	rv := objc.Send[ANETaskManager](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANETaskManager
type ANETaskManager struct {
	objectivec.Object
}

// ANETaskManagerFromID constructs a [ANETaskManager] from an objc.ID.
func ANETaskManagerFromID(id objc.ID) ANETaskManager {
	return ANETaskManager{objectivec.Object{ID: id}}
}
// Ensure ANETaskManager implements IANETaskManager.
var _ IANETaskManager = ANETaskManager{}

// An interface definition for the [ANETaskManager] class.
//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANETaskManager
type IANETaskManager interface {
	objectivec.IObject
}

// Init initializes the instance.
func (a ANETaskManager) Init() ANETaskManager {
	rv := objc.Send[ANETaskManager](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a ANETaskManager) Autorelease() ANETaskManager {
	rv := objc.Send[ANETaskManager](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewANETaskManager creates a new ANETaskManager instance.
func NewANETaskManager() ANETaskManager {
	class := getANETaskManagerClass()
	rv := objc.Send[ANETaskManager](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANETaskManager/registerTask:
func (_ANETaskManagerClass ANETaskManagerClass) RegisterTask(task objectivec.IObject) {
	objc.Send[objc.ID](objc.ID(_ANETaskManagerClass.class), objc.Sel("registerTask:"), task)
}
//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANETaskManager/unregisterTask:
func (_ANETaskManagerClass ANETaskManagerClass) UnregisterTask(task objectivec.IObject) {
	objc.Send[objc.ID](objc.ID(_ANETaskManagerClass.class), objc.Sel("unregisterTask:"), task)
}

