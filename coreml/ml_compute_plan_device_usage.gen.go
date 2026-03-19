// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
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

// Alloc allocates memory for a new instance of the class.
func (mc MLComputePlanDeviceUsageClass) Alloc() MLComputePlanDeviceUsage {
	rv := objc.Send[MLComputePlanDeviceUsage](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// The anticipated compute devices to use for executing a layer or operation.
//
// # Getting the compute device
//
//   - [MLComputePlanDeviceUsage.PreferredComputeDevice]: The compute device that the framework prefers to execute the layer/operation.
//   - [MLComputePlanDeviceUsage.SupportedComputeDevices]: The compute devices that can execute the layer/operation.
//
// See: https://developer.apple.com/documentation/CoreML/MLComputePlanDeviceUsage
type MLComputePlanDeviceUsage struct {
	objectivec.Object
}

// MLComputePlanDeviceUsageFromID constructs a [MLComputePlanDeviceUsage] from an objc.ID.
//
// The anticipated compute devices to use for executing a layer or operation.
func MLComputePlanDeviceUsageFromID(id objc.ID) MLComputePlanDeviceUsage {
	return MLComputePlanDeviceUsage{objectivec.Object{ID: id}}
}
// Ensure MLComputePlanDeviceUsage implements IMLComputePlanDeviceUsage.
var _ IMLComputePlanDeviceUsage = MLComputePlanDeviceUsage{}

// An interface definition for the [MLComputePlanDeviceUsage] class.
//
// # Getting the compute device
//
//   - [IMLComputePlanDeviceUsage.PreferredComputeDevice]: The compute device that the framework prefers to execute the layer/operation.
//   - [IMLComputePlanDeviceUsage.SupportedComputeDevices]: The compute devices that can execute the layer/operation.
//
// See: https://developer.apple.com/documentation/CoreML/MLComputePlanDeviceUsage
type IMLComputePlanDeviceUsage interface {
	objectivec.IObject

	// Topic: Getting the compute device

	// The compute device that the framework prefers to execute the layer/operation.
	PreferredComputeDevice() MLComputeDeviceProtocol
	// The compute devices that can execute the layer/operation.
	SupportedComputeDevices() []objectivec.IObject
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

// The compute device that the framework prefers to execute the
// layer/operation.
//
// See: https://developer.apple.com/documentation/CoreML/MLComputePlanDeviceUsage/preferredComputeDevice
func (c MLComputePlanDeviceUsage) PreferredComputeDevice() MLComputeDeviceProtocol {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("preferredComputeDevice"))
	return MLComputeDeviceProtocolObjectFromID(rv)
}
// The compute devices that can execute the layer/operation.
//
// See: https://developer.apple.com/documentation/CoreML/MLComputePlanDeviceUsage/supportedComputeDevices
func (c MLComputePlanDeviceUsage) SupportedComputeDevices() []objectivec.IObject {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("supportedComputeDevices"))
	return objc.ConvertSlice(rv, func(id objc.ID) objectivec.IObject {
		return objectivec.Object{ID: id}
	})
}

