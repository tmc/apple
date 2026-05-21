// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

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
type IMLComputePlanDeviceUsageSupportInfo interface {
	objectivec.IObject

	// Topic: Methods

	ComputeDevice() unsafe.Pointer
	State() int64
	InitWithComputeDeviceSupportState(device objectivec.IObject, state int64) MLComputePlanDeviceUsageSupportInfo
}

// Init initializes the instance.
func (m MLComputePlanDeviceUsageSupportInfo) Init() MLComputePlanDeviceUsageSupportInfo {
	rv := objc.Send[MLComputePlanDeviceUsageSupportInfo](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLComputePlanDeviceUsageSupportInfo) Autorelease() MLComputePlanDeviceUsageSupportInfo {
	rv := objc.Send[MLComputePlanDeviceUsageSupportInfo](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLComputePlanDeviceUsageSupportInfo creates a new MLComputePlanDeviceUsageSupportInfo instance.
func NewMLComputePlanDeviceUsageSupportInfo() MLComputePlanDeviceUsageSupportInfo {
	class := getMLComputePlanDeviceUsageSupportInfoClass()
	rv := objc.Send[MLComputePlanDeviceUsageSupportInfo](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewComputePlanDeviceUsageSupportInfoWithComputeDeviceSupportState(device objectivec.IObject, state int64) MLComputePlanDeviceUsageSupportInfo {
	instance := getMLComputePlanDeviceUsageSupportInfoClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithComputeDevice:supportState:"), device, state)
	return MLComputePlanDeviceUsageSupportInfoFromID(rv)
}

func (m MLComputePlanDeviceUsageSupportInfo) InitWithComputeDeviceSupportState(device objectivec.IObject, state int64) MLComputePlanDeviceUsageSupportInfo {
	rv := objc.Send[MLComputePlanDeviceUsageSupportInfo](m.ID, objc.Sel("initWithComputeDevice:supportState:"), device, state)
	return rv
}

func (m MLComputePlanDeviceUsageSupportInfo) ComputeDevice() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m.ID, objc.Sel("computeDevice"))
	return rv
}
func (m MLComputePlanDeviceUsageSupportInfo) State() int64 {
	rv := objc.Send[int64](m.ID, objc.Sel("state"))
	return rv
}
