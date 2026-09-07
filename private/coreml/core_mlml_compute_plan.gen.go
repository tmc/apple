// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CoreMLMLComputePlan] class.
var (
	_CoreMLMLComputePlanClass     CoreMLMLComputePlanClass
	_CoreMLMLComputePlanClassOnce sync.Once
)

func getCoreMLMLComputePlanClass() CoreMLMLComputePlanClass {
	_CoreMLMLComputePlanClassOnce.Do(func() {
		_CoreMLMLComputePlanClass = CoreMLMLComputePlanClass{class: objc.GetClass("CoreML.MLComputePlan")}
	})
	return _CoreMLMLComputePlanClass
}

// GetCoreMLMLComputePlanClass returns the class object for CoreML.MLComputePlan.
func GetCoreMLMLComputePlanClass() CoreMLMLComputePlanClass {
	return getCoreMLMLComputePlanClass()
}

type CoreMLMLComputePlanClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CoreMLMLComputePlanClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CoreMLMLComputePlanClass) Alloc() CoreMLMLComputePlan {
	rv := objc.SendIfResponds[CoreMLMLComputePlan](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

type CoreMLMLComputePlan struct {
	objectivec.Object
}

// CoreMLMLComputePlanFromID constructs a [CoreMLMLComputePlan] from an objc.ID.
func CoreMLMLComputePlanFromID(id objc.ID) CoreMLMLComputePlan {
	return CoreMLMLComputePlan{objectivec.Object{ID: id}}
}

// Ensure CoreMLMLComputePlan implements ICoreMLMLComputePlan.
var _ ICoreMLMLComputePlan = CoreMLMLComputePlan{}

// An interface definition for the [CoreMLMLComputePlan] class.
type ICoreMLMLComputePlan interface {
	objectivec.IObject
}

// Init initializes the instance.
func (c CoreMLMLComputePlan) Init() CoreMLMLComputePlan {
	rv := objc.SendIfResponds[CoreMLMLComputePlan](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CoreMLMLComputePlan) Autorelease() CoreMLMLComputePlan {
	rv := objc.SendIfResponds[CoreMLMLComputePlan](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCoreMLMLComputePlan creates a new CoreMLMLComputePlan instance.
func NewCoreMLMLComputePlan() CoreMLMLComputePlan {
	class := getCoreMLMLComputePlanClass()
	rv := objc.SendIfResponds[CoreMLMLComputePlan](objc.ID(class.class), objc.Sel("new"))
	return rv
}
