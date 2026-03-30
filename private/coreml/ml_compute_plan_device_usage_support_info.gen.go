// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLComputePlanDeviceUsageSupportInfo] class.
var (
	_MLComputePlanDeviceUsageSupportInfoClass     MLComputePlanDeviceUsageSupportInfoClass
	_MLComputePlanDeviceUsageSupportInfoClassOnce sync.Once
)

func getMLComputePlanDeviceUsageSupportInfoClass() MLComputePlanDeviceUsageSupportInfoClass {
	_MLComputePlanDeviceUsageSupportInfoClassOnce.Do(func() {
		_MLComputePlanDeviceUsageSupportInfoClass = MLComputePlanDeviceUsageSupportInfoClass{class: objc.GetClass("MLComputePlanDeviceUsageSupportInfo")}
	})
	return _MLComputePlanDeviceUsageSupportInfoClass
}

// GetMLComputePlanDeviceUsageSupportInfoClass returns the class object for MLComputePlanDeviceUsageSupportInfo.
func GetMLComputePlanDeviceUsageSupportInfoClass() MLComputePlanDeviceUsageSupportInfoClass {
	return getMLComputePlanDeviceUsageSupportInfoClass()
}

type MLComputePlanDeviceUsageSupportInfoClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLComputePlanDeviceUsageSupportInfoClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLComputePlanDeviceUsageSupportInfoClass) Alloc() MLComputePlanDeviceUsageSupportInfo {
	rv := objc.Send[MLComputePlanDeviceUsageSupportInfo](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLComputePlanDeviceUsageSupportInfo.ComputeDevice]
//   - [MLComputePlanDeviceUsageSupportInfo.State]
//   - [MLComputePlanDeviceUsageSupportInfo.InitWithComputeDeviceSupportState]
//
// See: https://developer.apple.com/documentation/CoreML/MLComputePlanDeviceUsageSupportInfo
type MLComputePlanDeviceUsageSupportInfo struct {
	objectivec.Object
}

// MLComputePlanDeviceUsageSupportInfoFromID constructs a [MLComputePlanDeviceUsageSupportInfo] from an objc.ID.
func MLComputePlanDeviceUsageSupportInfoFromID(id objc.ID) MLComputePlanDeviceUsageSupportInfo {
	return MLComputePlanDeviceUsageSupportInfo{objectivec.Object{ID: id}}
}

// Ensure MLComputePlanDeviceUsageSupportInfo implements IMLComputePlanDeviceUsageSupportInfo.
var _ IMLComputePlanDeviceUsageSupportInfo = MLComputePlanDeviceUsageSupportInfo{}

// An interface definition for the [MLComputePlanDeviceUsageSupportInfo] class.
//
// # Methods
//
//   - [IMLComputePlanDeviceUsageSupportInfo.ComputeDevice]
//   - [IMLComputePlanDeviceUsageSupportInfo.State]
//   - [IMLComputePlanDeviceUsageSupportInfo.InitWithComputeDeviceSupportState]
//
// See: https://developer.apple.com/documentation/CoreML/MLComputePlanDeviceUsageSupportInfo
type IMLComputePlanDeviceUsageSupportInfo interface {
	objectivec.IObject

	// Topic: Methods

	ComputeDevice() objectivec.IObject
	State() int64
	InitWithComputeDeviceSupportState(device objectivec.IObject, state int64) MLComputePlanDeviceUsageSupportInfo
}

// Init initializes the instance.
func (c MLComputePlanDeviceUsageSupportInfo) Init() MLComputePlanDeviceUsageSupportInfo {
	rv := objc.Send[MLComputePlanDeviceUsageSupportInfo](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MLComputePlanDeviceUsageSupportInfo) Autorelease() MLComputePlanDeviceUsageSupportInfo {
	rv := objc.Send[MLComputePlanDeviceUsageSupportInfo](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLComputePlanDeviceUsageSupportInfo creates a new MLComputePlanDeviceUsageSupportInfo instance.
func NewMLComputePlanDeviceUsageSupportInfo() MLComputePlanDeviceUsageSupportInfo {
	class := getMLComputePlanDeviceUsageSupportInfoClass()
	rv := objc.Send[MLComputePlanDeviceUsageSupportInfo](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLComputePlanDeviceUsageSupportInfo/initWithComputeDevice:supportState:
func NewComputePlanDeviceUsageSupportInfoWithComputeDeviceSupportState(device objectivec.IObject, state int64) MLComputePlanDeviceUsageSupportInfo {
	instance := getMLComputePlanDeviceUsageSupportInfoClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithComputeDevice:supportState:"), device, state)
	return MLComputePlanDeviceUsageSupportInfoFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLComputePlanDeviceUsageSupportInfo/initWithComputeDevice:supportState:
func (c MLComputePlanDeviceUsageSupportInfo) InitWithComputeDeviceSupportState(device objectivec.IObject, state int64) MLComputePlanDeviceUsageSupportInfo {
	rv := objc.Send[MLComputePlanDeviceUsageSupportInfo](c.ID, objc.Sel("initWithComputeDevice:supportState:"), device, state)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLComputePlanDeviceUsageSupportInfo/computeDevice
func (c MLComputePlanDeviceUsageSupportInfo) ComputeDevice() objectivec.IObject {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("computeDevice"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/CoreML/MLComputePlanDeviceUsageSupportInfo/state
func (c MLComputePlanDeviceUsageSupportInfo) State() int64 {
	rv := objc.Send[int64](c.ID, objc.Sel("state"))
	return rv
}
