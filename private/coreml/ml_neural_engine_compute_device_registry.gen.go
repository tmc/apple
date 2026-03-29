// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
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
	rv := objc.Send[MLNeuralEngineComputeDeviceRegistry](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [MLNeuralEngineComputeDeviceRegistry.NeuralEngineDevice]
//   - [MLNeuralEngineComputeDeviceRegistry.RegisteredComputeDevices]
//   - [MLNeuralEngineComputeDeviceRegistry.InitWithNeuralEngineDevice]
//   - [MLNeuralEngineComputeDeviceRegistry.DebugDescription]
//   - [MLNeuralEngineComputeDeviceRegistry.Description]
//   - [MLNeuralEngineComputeDeviceRegistry.Hash]
//   - [MLNeuralEngineComputeDeviceRegistry.Superclass]
// See: https://developer.apple.com/documentation/CoreML/MLNeuralEngineComputeDeviceRegistry
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
//
// See: https://developer.apple.com/documentation/CoreML/MLNeuralEngineComputeDeviceRegistry
type IMLNeuralEngineComputeDeviceRegistry interface {
	objectivec.IObject

	// Topic: Methods

	NeuralEngineDevice() IMLNeuralEngineComputeDevice
	RegisteredComputeDevices() foundation.INSArray
	InitWithNeuralEngineDevice(device objectivec.IObject) MLNeuralEngineComputeDeviceRegistry
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objc.Class
}

// Init initializes the instance.
func (n MLNeuralEngineComputeDeviceRegistry) Init() MLNeuralEngineComputeDeviceRegistry {
	rv := objc.Send[MLNeuralEngineComputeDeviceRegistry](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n MLNeuralEngineComputeDeviceRegistry) Autorelease() MLNeuralEngineComputeDeviceRegistry {
	rv := objc.Send[MLNeuralEngineComputeDeviceRegistry](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLNeuralEngineComputeDeviceRegistry creates a new MLNeuralEngineComputeDeviceRegistry instance.
func NewMLNeuralEngineComputeDeviceRegistry() MLNeuralEngineComputeDeviceRegistry {
	class := getMLNeuralEngineComputeDeviceRegistryClass()
	rv := objc.Send[MLNeuralEngineComputeDeviceRegistry](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/CoreML/MLNeuralEngineComputeDeviceRegistry/initWithNeuralEngineDevice:
func NewNeuralEngineComputeDeviceRegistryWithNeuralEngineDevice(device objectivec.IObject) MLNeuralEngineComputeDeviceRegistry {
	instance := getMLNeuralEngineComputeDeviceRegistryClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithNeuralEngineDevice:"), device)
	return MLNeuralEngineComputeDeviceRegistryFromID(rv)
}

//
// See: https://developer.apple.com/documentation/CoreML/MLNeuralEngineComputeDeviceRegistry/initWithNeuralEngineDevice:
func (n MLNeuralEngineComputeDeviceRegistry) InitWithNeuralEngineDevice(device objectivec.IObject) MLNeuralEngineComputeDeviceRegistry {
	rv := objc.Send[MLNeuralEngineComputeDeviceRegistry](n.ID, objc.Sel("initWithNeuralEngineDevice:"), device)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLNeuralEngineComputeDeviceRegistry/sharedRegistry
func (_MLNeuralEngineComputeDeviceRegistryClass MLNeuralEngineComputeDeviceRegistryClass) SharedRegistry() MLNeuralEngineComputeDeviceRegistry {
	rv := objc.Send[objc.ID](objc.ID(_MLNeuralEngineComputeDeviceRegistryClass.class), objc.Sel("sharedRegistry"))
	return MLNeuralEngineComputeDeviceRegistryFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLNeuralEngineComputeDeviceRegistry/debugDescription
func (n MLNeuralEngineComputeDeviceRegistry) DebugDescription() string {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/CoreML/MLNeuralEngineComputeDeviceRegistry/description
func (n MLNeuralEngineComputeDeviceRegistry) Description() string {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/CoreML/MLNeuralEngineComputeDeviceRegistry/hash
func (n MLNeuralEngineComputeDeviceRegistry) Hash() uint64 {
	rv := objc.Send[uint64](n.ID, objc.Sel("hash"))
	return rv
}
// See: https://developer.apple.com/documentation/CoreML/MLNeuralEngineComputeDeviceRegistry/neuralEngineDevice
func (n MLNeuralEngineComputeDeviceRegistry) NeuralEngineDevice() IMLNeuralEngineComputeDevice {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("neuralEngineDevice"))
	return MLNeuralEngineComputeDeviceFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/CoreML/MLNeuralEngineComputeDeviceRegistry/registeredComputeDevices
func (n MLNeuralEngineComputeDeviceRegistry) RegisteredComputeDevices() foundation.INSArray {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("registeredComputeDevices"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/CoreML/MLNeuralEngineComputeDeviceRegistry/superclass
func (n MLNeuralEngineComputeDeviceRegistry) Superclass() objc.Class {
	rv := objc.Send[objc.Class](n.ID, objc.Sel("superclass"))
	return rv
}

