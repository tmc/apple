// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLComputePlanDeviceUsage] class.
var (
	_MLComputePlanDeviceUsageClass     MLComputePlanDeviceUsageClass
	_MLComputePlanDeviceUsageClassOnce sync.Once
)

func getMLComputePlanDeviceUsageClass() MLComputePlanDeviceUsageClass {
	_MLComputePlanDeviceUsageClassOnce.Do(func() {
		_MLComputePlanDeviceUsageClass = MLComputePlanDeviceUsageClass{class: objc.GetClass("MLComputePlanDeviceUsage")}
	})
	return _MLComputePlanDeviceUsageClass
}

// GetMLComputePlanDeviceUsageClass returns the class object for MLComputePlanDeviceUsage.
func GetMLComputePlanDeviceUsageClass() MLComputePlanDeviceUsageClass {
	return getMLComputePlanDeviceUsageClass()
}

type MLComputePlanDeviceUsageClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLComputePlanDeviceUsageClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLComputePlanDeviceUsageClass) Alloc() MLComputePlanDeviceUsage {
	rv := objc.Send[MLComputePlanDeviceUsage](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLComputePlanDeviceUsage.DeviceSupportInfoArray]
//   - [MLComputePlanDeviceUsage.SupportInfoForComputeDevice]
//   - [MLComputePlanDeviceUsage.InitWithSupportedComputeDevicesPreferredComputeDeviceDeviceSupportInfoArray]
//
// See: https://developer.apple.com/documentation/CoreML/MLComputePlanDeviceUsage
type MLComputePlanDeviceUsage struct {
	objectivec.Object
}

// MLComputePlanDeviceUsageFromID constructs a [MLComputePlanDeviceUsage] from an objc.ID.
func MLComputePlanDeviceUsageFromID(id objc.ID) MLComputePlanDeviceUsage {
	return MLComputePlanDeviceUsage{objectivec.Object{ID: id}}
}

// Ensure MLComputePlanDeviceUsage implements IMLComputePlanDeviceUsage.
var _ IMLComputePlanDeviceUsage = MLComputePlanDeviceUsage{}

// An interface definition for the [MLComputePlanDeviceUsage] class.
//
// # Methods
//
//   - [IMLComputePlanDeviceUsage.DeviceSupportInfoArray]
//   - [IMLComputePlanDeviceUsage.SupportInfoForComputeDevice]
//   - [IMLComputePlanDeviceUsage.InitWithSupportedComputeDevicesPreferredComputeDeviceDeviceSupportInfoArray]
//
// See: https://developer.apple.com/documentation/CoreML/MLComputePlanDeviceUsage
type IMLComputePlanDeviceUsage interface {
	objectivec.IObject

	// Topic: Methods

	DeviceSupportInfoArray() foundation.INSArray
	SupportInfoForComputeDevice(device objectivec.IObject) objectivec.IObject
	InitWithSupportedComputeDevicesPreferredComputeDeviceDeviceSupportInfoArray(devices objectivec.IObject, device objectivec.IObject, array objectivec.IObject) MLComputePlanDeviceUsage
}

// Init initializes the instance.
func (c MLComputePlanDeviceUsage) Init() MLComputePlanDeviceUsage {
	rv := objc.Send[MLComputePlanDeviceUsage](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MLComputePlanDeviceUsage) Autorelease() MLComputePlanDeviceUsage {
	rv := objc.Send[MLComputePlanDeviceUsage](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLComputePlanDeviceUsage creates a new MLComputePlanDeviceUsage instance.
func NewMLComputePlanDeviceUsage() MLComputePlanDeviceUsage {
	class := getMLComputePlanDeviceUsageClass()
	rv := objc.Send[MLComputePlanDeviceUsage](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLComputePlanDeviceUsage/initWithSupportedComputeDevices:preferredComputeDevice:deviceSupportInfoArray:
func NewComputePlanDeviceUsageWithSupportedComputeDevicesPreferredComputeDeviceDeviceSupportInfoArray(devices objectivec.IObject, device objectivec.IObject, array objectivec.IObject) MLComputePlanDeviceUsage {
	instance := getMLComputePlanDeviceUsageClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSupportedComputeDevices:preferredComputeDevice:deviceSupportInfoArray:"), devices, device, array)
	return MLComputePlanDeviceUsageFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLComputePlanDeviceUsage/supportInfoForComputeDevice:
func (c MLComputePlanDeviceUsage) SupportInfoForComputeDevice(device objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("supportInfoForComputeDevice:"), device)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/CoreML/MLComputePlanDeviceUsage/initWithSupportedComputeDevices:preferredComputeDevice:deviceSupportInfoArray:
func (c MLComputePlanDeviceUsage) InitWithSupportedComputeDevicesPreferredComputeDeviceDeviceSupportInfoArray(devices objectivec.IObject, device objectivec.IObject, array objectivec.IObject) MLComputePlanDeviceUsage {
	rv := objc.Send[MLComputePlanDeviceUsage](c.ID, objc.Sel("initWithSupportedComputeDevices:preferredComputeDevice:deviceSupportInfoArray:"), devices, device, array)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLComputePlanDeviceUsage/deviceSupportInfoArray
func (c MLComputePlanDeviceUsage) DeviceSupportInfoArray() foundation.INSArray {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("deviceSupportInfoArray"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
