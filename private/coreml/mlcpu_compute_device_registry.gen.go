// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLCPUComputeDeviceRegistry] class.
var (
	_MLCPUComputeDeviceRegistryClass     MLCPUComputeDeviceRegistryClass
	_MLCPUComputeDeviceRegistryClassOnce sync.Once
)

func getMLCPUComputeDeviceRegistryClass() MLCPUComputeDeviceRegistryClass {
	_MLCPUComputeDeviceRegistryClassOnce.Do(func() {
		_MLCPUComputeDeviceRegistryClass = MLCPUComputeDeviceRegistryClass{class: objc.GetClass("MLCPUComputeDeviceRegistry")}
	})
	return _MLCPUComputeDeviceRegistryClass
}

// GetMLCPUComputeDeviceRegistryClass returns the class object for MLCPUComputeDeviceRegistry.
func GetMLCPUComputeDeviceRegistryClass() MLCPUComputeDeviceRegistryClass {
	return getMLCPUComputeDeviceRegistryClass()
}

type MLCPUComputeDeviceRegistryClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLCPUComputeDeviceRegistryClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLCPUComputeDeviceRegistryClass) Alloc() MLCPUComputeDeviceRegistry {
	rv := objc.Send[MLCPUComputeDeviceRegistry](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLCPUComputeDeviceRegistry.CpuDevice]
//   - [MLCPUComputeDeviceRegistry.RegisteredComputeDevices]
//   - [MLCPUComputeDeviceRegistry.InitWithCpuDevice]
//   - [MLCPUComputeDeviceRegistry.DebugDescription]
//   - [MLCPUComputeDeviceRegistry.Description]
//   - [MLCPUComputeDeviceRegistry.Hash]
//   - [MLCPUComputeDeviceRegistry.Superclass]
type MLCPUComputeDeviceRegistry struct {
	objectivec.Object
}

// MLCPUComputeDeviceRegistryFromID constructs a [MLCPUComputeDeviceRegistry] from an objc.ID.
func MLCPUComputeDeviceRegistryFromID(id objc.ID) MLCPUComputeDeviceRegistry {
	return MLCPUComputeDeviceRegistry{objectivec.Object{ID: id}}
}

// Ensure MLCPUComputeDeviceRegistry implements IMLCPUComputeDeviceRegistry.
var _ IMLCPUComputeDeviceRegistry = MLCPUComputeDeviceRegistry{}

// An interface definition for the [MLCPUComputeDeviceRegistry] class.
//
// # Methods
//
//   - [IMLCPUComputeDeviceRegistry.CpuDevice]
//   - [IMLCPUComputeDeviceRegistry.RegisteredComputeDevices]
//   - [IMLCPUComputeDeviceRegistry.InitWithCpuDevice]
//   - [IMLCPUComputeDeviceRegistry.DebugDescription]
//   - [IMLCPUComputeDeviceRegistry.Description]
//   - [IMLCPUComputeDeviceRegistry.Hash]
//   - [IMLCPUComputeDeviceRegistry.Superclass]
type IMLCPUComputeDeviceRegistry interface {
	objectivec.IObject

	// Topic: Methods

	CpuDevice() IMLCPUComputeDevice
	RegisteredComputeDevices() foundation.INSArray
	InitWithCpuDevice(device objectivec.IObject) MLCPUComputeDeviceRegistry
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objectivec.Class
}

// Init initializes the instance.
func (m MLCPUComputeDeviceRegistry) Init() MLCPUComputeDeviceRegistry {
	rv := objc.Send[MLCPUComputeDeviceRegistry](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLCPUComputeDeviceRegistry) Autorelease() MLCPUComputeDeviceRegistry {
	rv := objc.Send[MLCPUComputeDeviceRegistry](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLCPUComputeDeviceRegistry creates a new MLCPUComputeDeviceRegistry instance.
func NewMLCPUComputeDeviceRegistry() MLCPUComputeDeviceRegistry {
	class := getMLCPUComputeDeviceRegistryClass()
	rv := objc.Send[MLCPUComputeDeviceRegistry](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewCPUComputeDeviceRegistryWithCpuDevice(device objectivec.IObject) MLCPUComputeDeviceRegistry {
	instance := getMLCPUComputeDeviceRegistryClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCpuDevice:"), device)
	return MLCPUComputeDeviceRegistryFromID(rv)
}

func (m MLCPUComputeDeviceRegistry) InitWithCpuDevice(device objectivec.IObject) MLCPUComputeDeviceRegistry {
	rv := objc.Send[MLCPUComputeDeviceRegistry](m.ID, objc.Sel("initWithCpuDevice:"), device)
	return rv
}

func (_MLCPUComputeDeviceRegistryClass MLCPUComputeDeviceRegistryClass) SharedRegistry() MLCPUComputeDeviceRegistry {
	rv := objc.Send[objc.ID](objc.ID(_MLCPUComputeDeviceRegistryClass.class), objc.Sel("sharedRegistry"))
	return MLCPUComputeDeviceRegistryFromID(rv)
}

func (m MLCPUComputeDeviceRegistry) CpuDevice() IMLCPUComputeDevice {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("cpuDevice"))
	return MLCPUComputeDeviceFromID(objc.ID(rv))
}
func (m MLCPUComputeDeviceRegistry) DebugDescription() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
func (m MLCPUComputeDeviceRegistry) Description() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
func (m MLCPUComputeDeviceRegistry) Hash() uint64 {
	rv := objc.Send[uint64](m.ID, objc.Sel("hash"))
	return rv
}
func (m MLCPUComputeDeviceRegistry) RegisteredComputeDevices() foundation.INSArray {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("registeredComputeDevices"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (m MLCPUComputeDeviceRegistry) Superclass() objectivec.Class {
	rv := objc.Send[objectivec.Class](m.ID, objc.Sel("superclass"))
	return objectivec.Class(rv)
}
