// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSAccelerationStructureGroup] class.
var (
	_MPSAccelerationStructureGroupClass     MPSAccelerationStructureGroupClass
	_MPSAccelerationStructureGroupClassOnce sync.Once
)

func getMPSAccelerationStructureGroupClass() MPSAccelerationStructureGroupClass {
	_MPSAccelerationStructureGroupClassOnce.Do(func() {
		_MPSAccelerationStructureGroupClass = MPSAccelerationStructureGroupClass{class: objc.GetClass("MPSAccelerationStructureGroup")}
	})
	return _MPSAccelerationStructureGroupClass
}

// GetMPSAccelerationStructureGroupClass returns the class object for MPSAccelerationStructureGroup.
func GetMPSAccelerationStructureGroupClass() MPSAccelerationStructureGroupClass {
	return getMPSAccelerationStructureGroupClass()
}

type MPSAccelerationStructureGroupClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSAccelerationStructureGroupClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSAccelerationStructureGroupClass) Alloc() MPSAccelerationStructureGroup {
	rv := objc.Send[MPSAccelerationStructureGroup](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A group of acceleration structures.
//
// # Instance Properties
//
//   - [MPSAccelerationStructureGroup.Device]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSAccelerationStructureGroup
type MPSAccelerationStructureGroup struct {
	objectivec.Object
}

// MPSAccelerationStructureGroupFromID constructs a [MPSAccelerationStructureGroup] from an objc.ID.
//
// A group of acceleration structures.
func MPSAccelerationStructureGroupFromID(id objc.ID) MPSAccelerationStructureGroup {
	return MPSAccelerationStructureGroup{objectivec.Object{ID: id}}
}

// NOTE: MPSAccelerationStructureGroup adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSAccelerationStructureGroup] class.
//
// # Instance Properties
//
//   - [IMPSAccelerationStructureGroup.Device]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSAccelerationStructureGroup
type IMPSAccelerationStructureGroup interface {
	objectivec.IObject

	// Topic: Instance Properties

	Device() metal.MTLDevice
}

// Init initializes the instance.
func (a MPSAccelerationStructureGroup) Init() MPSAccelerationStructureGroup {
	rv := objc.Send[MPSAccelerationStructureGroup](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a MPSAccelerationStructureGroup) Autorelease() MPSAccelerationStructureGroup {
	rv := objc.Send[MPSAccelerationStructureGroup](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSAccelerationStructureGroup creates a new MPSAccelerationStructureGroup instance.
func NewMPSAccelerationStructureGroup() MPSAccelerationStructureGroup {
	class := getMPSAccelerationStructureGroupClass()
	rv := objc.Send[MPSAccelerationStructureGroup](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSAccelerationStructureGroup/init(device:)
func NewAccelerationStructureGroupWithDevice(device metal.MTLDevice) MPSAccelerationStructureGroup {
	instance := getMPSAccelerationStructureGroupClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return MPSAccelerationStructureGroupFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSAccelerationStructureGroup/device
func (a MPSAccelerationStructureGroup) Device() metal.MTLDevice {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("device"))
	return metal.MTLDeviceObjectFromID(rv)
}
