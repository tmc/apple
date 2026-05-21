// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [ANEAnalyticsGroup] class.
var (
	_ANEAnalyticsGroupClass     ANEAnalyticsGroupClass
	_ANEAnalyticsGroupClassOnce sync.Once
)

func getANEAnalyticsGroupClass() ANEAnalyticsGroupClass {
	_ANEAnalyticsGroupClassOnce.Do(func() {
		_ANEAnalyticsGroupClass = ANEAnalyticsGroupClass{class: objc.GetClass("_ANEAnalyticsGroup")}
	})
	return _ANEAnalyticsGroupClass
}

// GetANEAnalyticsGroupClass returns the class object for _ANEAnalyticsGroup.
func GetANEAnalyticsGroupClass() ANEAnalyticsGroupClass {
	return getANEAnalyticsGroupClass()
}

type ANEAnalyticsGroupClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac ANEAnalyticsGroupClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac ANEAnalyticsGroupClass) Alloc() ANEAnalyticsGroup {
	rv := objc.Send[ANEAnalyticsGroup](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [ANEAnalyticsGroup.GroupID]
//   - [ANEAnalyticsGroup.LayerInfo]
//   - [ANEAnalyticsGroup.Serialize]
//   - [ANEAnalyticsGroup.TaskInfo]
//   - [ANEAnalyticsGroup.InitWithIDLayersTasks]
type ANEAnalyticsGroup struct {
	objectivec.Object
}

// ANEAnalyticsGroupFromID constructs a [ANEAnalyticsGroup] from an objc.ID.
func ANEAnalyticsGroupFromID(id objc.ID) ANEAnalyticsGroup {
	return ANEAnalyticsGroup{objectivec.Object{ID: id}}
}

// Ensure ANEAnalyticsGroup implements IANEAnalyticsGroup.
var _ IANEAnalyticsGroup = ANEAnalyticsGroup{}

// An interface definition for the [ANEAnalyticsGroup] class.
//
// # Methods
//
//   - [IANEAnalyticsGroup.GroupID]
//   - [IANEAnalyticsGroup.LayerInfo]
//   - [IANEAnalyticsGroup.Serialize]
//   - [IANEAnalyticsGroup.TaskInfo]
//   - [IANEAnalyticsGroup.InitWithIDLayersTasks]
type IANEAnalyticsGroup interface {
	objectivec.IObject

	// Topic: Methods

	GroupID() foundation.NSNumber
	LayerInfo() foundation.INSArray
	Serialize() objectivec.IObject
	TaskInfo() foundation.INSArray
	InitWithIDLayersTasks(id objectivec.IObject, layers objectivec.IObject, tasks objectivec.IObject) ANEAnalyticsGroup
}

// Init initializes the instance.
func (a ANEAnalyticsGroup) Init() ANEAnalyticsGroup {
	rv := objc.Send[ANEAnalyticsGroup](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a ANEAnalyticsGroup) Autorelease() ANEAnalyticsGroup {
	rv := objc.Send[ANEAnalyticsGroup](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewANEAnalyticsGroup creates a new ANEAnalyticsGroup instance.
func NewANEAnalyticsGroup() ANEAnalyticsGroup {
	class := getANEAnalyticsGroupClass()
	rv := objc.Send[ANEAnalyticsGroup](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewANEAnalyticsGroupWithIDLayersTasks(id objectivec.IObject, layers objectivec.IObject, tasks objectivec.IObject) ANEAnalyticsGroup {
	instance := getANEAnalyticsGroupClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithID:layers:tasks:"), id, layers, tasks)
	return ANEAnalyticsGroupFromID(rv)
}

func (a ANEAnalyticsGroup) Serialize() objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("serialize"))
	return objectivec.Object{ID: rv}
}
func (a ANEAnalyticsGroup) InitWithIDLayersTasks(id objectivec.IObject, layers objectivec.IObject, tasks objectivec.IObject) ANEAnalyticsGroup {
	rv := objc.Send[ANEAnalyticsGroup](a.ID, objc.Sel("initWithID:layers:tasks:"), id, layers, tasks)
	return rv
}

func (_ANEAnalyticsGroupClass ANEAnalyticsGroupClass) ObjectWithIDLayersTasks(id objectivec.IObject, layers objectivec.IObject, tasks objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEAnalyticsGroupClass.class), objc.Sel("objectWithID:layers:tasks:"), id, layers, tasks)
	return objectivec.Object{ID: rv}
}

func (a ANEAnalyticsGroup) GroupID() foundation.NSNumber {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("groupID"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
func (a ANEAnalyticsGroup) LayerInfo() foundation.INSArray {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("layerInfo"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (a ANEAnalyticsGroup) TaskInfo() foundation.INSArray {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("taskInfo"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
