// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLNeuralEngineComputeDeviceRegistry] class.
var (
	_MLNeuralEngineComputeDeviceRegistryClass     MLNeuralEngineComputeDeviceRegistryClass
	_MLNeuralEngineComputeDeviceRegistryClassOnce sync.Once
)

func getMLNeuralEngineComputeDeviceRegistryClass() MLNeuralEngineComputeDeviceRegistryClass {
	_MLNeuralEngineComputeDeviceRegistryClassOnce.Do(func() {
		_MLNeuralEngineComputeDeviceRegistryClass = MLNeuralEngineComputeDeviceRegistryClass{class: objc.GetClass("MLNeuralEngineComputeDeviceRegistry")}
	})
	return _MLNeuralEngineComputeDeviceRegistryClass
}

// GetMLNeuralEngineComputeDeviceRegistryClass returns the class object for MLNeuralEngineComputeDeviceRegistry.
func GetMLNeuralEngineComputeDeviceRegistryClass() MLNeuralEngineComputeDeviceRegistryClass {
	return getMLNeuralEngineComputeDeviceRegistryClass()
}

type MLNeuralEngineComputeDeviceRegistryClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLNeuralEngineComputeDeviceRegistryClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLNeuralEngineComputeDeviceRegistryClass) Alloc() MLNeuralEngineComputeDeviceRegistry {
	rv := objc.SendIfResponds[MLNeuralEngineComputeDeviceRegistry](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLNeuralEngineComputeDeviceRegistry.NeuralEngineDevice]
//   - [MLNeuralEngineComputeDeviceRegistry.RegisteredComputeDevices]
//   - [MLNeuralEngineComputeDeviceRegistry.InitWithNeuralEngineDevice]
//   - [MLNeuralEngineComputeDeviceRegistry.DebugDescription]
//   - [MLNeuralEngineComputeDeviceRegistry.Description]
//   - [MLNeuralEngineComputeDeviceRegistry.Hash]
//   - [MLNeuralEngineComputeDeviceRegistry.Superclass]
type MLNeuralEngineComputeDeviceRegistry struct {
	objectivec.Object
}

// MLNeuralEngineComputeDeviceRegistryFromID constructs a [MLNeuralEngineComputeDeviceRegistry] from an objc.ID.
func MLNeuralEngineComputeDeviceRegistryFromID(id objc.ID) MLNeuralEngineComputeDeviceRegistry {
	return MLNeuralEngineComputeDeviceRegistry{objectivec.Object{ID: id}}
}

// Ensure MLNeuralEngineComputeDeviceRegistry implements IMLNeuralEngineComputeDeviceRegistry.
var _ IMLNeuralEngineComputeDeviceRegistry = MLNeuralEngineComputeDeviceRegistry{}

// An interface definition for the [MLNeuralEngineComputeDeviceRegistry] class.
//
// # Methods
//
//   - [IMLNeuralEngineComputeDeviceRegistry.NeuralEngineDevice]
//   - [IMLNeuralEngineComputeDeviceRegistry.RegisteredComputeDevices]
//   - [IMLNeuralEngineComputeDeviceRegistry.InitWithNeuralEngineDevice]
//   - [IMLNeuralEngineComputeDeviceRegistry.DebugDescription]
//   - [IMLNeuralEngineComputeDeviceRegistry.Description]
//   - [IMLNeuralEngineComputeDeviceRegistry.Hash]
//   - [IMLNeuralEngineComputeDeviceRegistry.Superclass]
type IMLNeuralEngineComputeDeviceRegistry interface {
	objectivec.IObject

	// Topic: Methods

	NeuralEngineDevice() IMLNeuralEngineComputeDevice
	RegisteredComputeDevices() foundation.INSArray
	InitWithNeuralEngineDevice(device objectivec.IObject) MLNeuralEngineComputeDeviceRegistry
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objectivec.Class
}

// Init initializes the instance.
func (m MLNeuralEngineComputeDeviceRegistry) Init() MLNeuralEngineComputeDeviceRegistry {
	rv := objc.SendIfResponds[MLNeuralEngineComputeDeviceRegistry](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLNeuralEngineComputeDeviceRegistry) Autorelease() MLNeuralEngineComputeDeviceRegistry {
	rv := objc.SendIfResponds[MLNeuralEngineComputeDeviceRegistry](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLNeuralEngineComputeDeviceRegistry creates a new MLNeuralEngineComputeDeviceRegistry instance.
func NewMLNeuralEngineComputeDeviceRegistry() MLNeuralEngineComputeDeviceRegistry {
	class := getMLNeuralEngineComputeDeviceRegistryClass()
	rv := objc.SendIfResponds[MLNeuralEngineComputeDeviceRegistry](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewNeuralEngineComputeDeviceRegistryWithNeuralEngineDevice(device objectivec.IObject) MLNeuralEngineComputeDeviceRegistry {
	instance := getMLNeuralEngineComputeDeviceRegistryClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithNeuralEngineDevice:"), device)
	return MLNeuralEngineComputeDeviceRegistryFromID(rv)
}

func (m MLNeuralEngineComputeDeviceRegistry) InitWithNeuralEngineDevice(device objectivec.IObject) MLNeuralEngineComputeDeviceRegistry {
	rv := objc.SendIfResponds[MLNeuralEngineComputeDeviceRegistry](m.ID, objc.Sel("initWithNeuralEngineDevice:"), device)
	return rv
}

func (_MLNeuralEngineComputeDeviceRegistryClass MLNeuralEngineComputeDeviceRegistryClass) SharedRegistry() MLNeuralEngineComputeDeviceRegistry {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_MLNeuralEngineComputeDeviceRegistryClass.class), objc.Sel("sharedRegistry"))
	return MLNeuralEngineComputeDeviceRegistryFromID(rv)
}

func (m MLNeuralEngineComputeDeviceRegistry) DebugDescription() string {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
func (m MLNeuralEngineComputeDeviceRegistry) Description() string {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
func (m MLNeuralEngineComputeDeviceRegistry) Hash() uint64 {
	rv := objc.SendIfResponds[uint64](m.ID, objc.Sel("hash"))
	return rv
}
func (m MLNeuralEngineComputeDeviceRegistry) NeuralEngineDevice() IMLNeuralEngineComputeDevice {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("neuralEngineDevice"))
	return MLNeuralEngineComputeDeviceFromID(objc.ID(rv))
}
func (m MLNeuralEngineComputeDeviceRegistry) RegisteredComputeDevices() foundation.INSArray {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("registeredComputeDevices"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (m MLNeuralEngineComputeDeviceRegistry) Superclass() objectivec.Class {
	rv := objc.SendIfResponds[objectivec.Class](m.ID, objc.Sel("superclass"))
	return objectivec.Class(rv)
}
