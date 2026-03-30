// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLAllComputeDeviceRegistry] class.
var (
	_MLAllComputeDeviceRegistryClass     MLAllComputeDeviceRegistryClass
	_MLAllComputeDeviceRegistryClassOnce sync.Once
)

func getMLAllComputeDeviceRegistryClass() MLAllComputeDeviceRegistryClass {
	_MLAllComputeDeviceRegistryClassOnce.Do(func() {
		_MLAllComputeDeviceRegistryClass = MLAllComputeDeviceRegistryClass{class: objc.GetClass("MLAllComputeDeviceRegistry")}
	})
	return _MLAllComputeDeviceRegistryClass
}

// GetMLAllComputeDeviceRegistryClass returns the class object for MLAllComputeDeviceRegistry.
func GetMLAllComputeDeviceRegistryClass() MLAllComputeDeviceRegistryClass {
	return getMLAllComputeDeviceRegistryClass()
}

type MLAllComputeDeviceRegistryClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLAllComputeDeviceRegistryClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLAllComputeDeviceRegistryClass) Alloc() MLAllComputeDeviceRegistry {
	rv := objc.Send[MLAllComputeDeviceRegistry](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLAllComputeDeviceRegistry.DeviceRegistries]
//   - [MLAllComputeDeviceRegistry.RegisteredComputeDevices]
//   - [MLAllComputeDeviceRegistry.InitWithDeviceRegistries]
//   - [MLAllComputeDeviceRegistry.DebugDescription]
//   - [MLAllComputeDeviceRegistry.Description]
//   - [MLAllComputeDeviceRegistry.Hash]
//   - [MLAllComputeDeviceRegistry.Superclass]
//
// See: https://developer.apple.com/documentation/CoreML/MLAllComputeDeviceRegistry
type MLAllComputeDeviceRegistry struct {
	objectivec.Object
}

// MLAllComputeDeviceRegistryFromID constructs a [MLAllComputeDeviceRegistry] from an objc.ID.
func MLAllComputeDeviceRegistryFromID(id objc.ID) MLAllComputeDeviceRegistry {
	return MLAllComputeDeviceRegistry{objectivec.Object{ID: id}}
}

// Ensure MLAllComputeDeviceRegistry implements IMLAllComputeDeviceRegistry.
var _ IMLAllComputeDeviceRegistry = MLAllComputeDeviceRegistry{}

// An interface definition for the [MLAllComputeDeviceRegistry] class.
//
// # Methods
//
//   - [IMLAllComputeDeviceRegistry.DeviceRegistries]
//   - [IMLAllComputeDeviceRegistry.RegisteredComputeDevices]
//   - [IMLAllComputeDeviceRegistry.InitWithDeviceRegistries]
//   - [IMLAllComputeDeviceRegistry.DebugDescription]
//   - [IMLAllComputeDeviceRegistry.Description]
//   - [IMLAllComputeDeviceRegistry.Hash]
//   - [IMLAllComputeDeviceRegistry.Superclass]
//
// See: https://developer.apple.com/documentation/CoreML/MLAllComputeDeviceRegistry
type IMLAllComputeDeviceRegistry interface {
	objectivec.IObject

	// Topic: Methods

	DeviceRegistries() foundation.INSArray
	RegisteredComputeDevices() foundation.INSArray
	InitWithDeviceRegistries(registries objectivec.IObject) MLAllComputeDeviceRegistry
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objc.Class
}

// Init initializes the instance.
func (a MLAllComputeDeviceRegistry) Init() MLAllComputeDeviceRegistry {
	rv := objc.Send[MLAllComputeDeviceRegistry](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a MLAllComputeDeviceRegistry) Autorelease() MLAllComputeDeviceRegistry {
	rv := objc.Send[MLAllComputeDeviceRegistry](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLAllComputeDeviceRegistry creates a new MLAllComputeDeviceRegistry instance.
func NewMLAllComputeDeviceRegistry() MLAllComputeDeviceRegistry {
	class := getMLAllComputeDeviceRegistryClass()
	rv := objc.Send[MLAllComputeDeviceRegistry](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLAllComputeDeviceRegistry/initWithDeviceRegistries:
func NewAllComputeDeviceRegistryWithDeviceRegistries(registries objectivec.IObject) MLAllComputeDeviceRegistry {
	instance := getMLAllComputeDeviceRegistryClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDeviceRegistries:"), registries)
	return MLAllComputeDeviceRegistryFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLAllComputeDeviceRegistry/initWithDeviceRegistries:
func (a MLAllComputeDeviceRegistry) InitWithDeviceRegistries(registries objectivec.IObject) MLAllComputeDeviceRegistry {
	rv := objc.Send[MLAllComputeDeviceRegistry](a.ID, objc.Sel("initWithDeviceRegistries:"), registries)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLAllComputeDeviceRegistry/registryWithDeviceRegistries:
func (_MLAllComputeDeviceRegistryClass MLAllComputeDeviceRegistryClass) RegistryWithDeviceRegistries(registries objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLAllComputeDeviceRegistryClass.class), objc.Sel("registryWithDeviceRegistries:"), registries)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/CoreML/MLAllComputeDeviceRegistry/sharedRegistry
func (_MLAllComputeDeviceRegistryClass MLAllComputeDeviceRegistryClass) SharedRegistry() MLAllComputeDeviceRegistry {
	rv := objc.Send[objc.ID](objc.ID(_MLAllComputeDeviceRegistryClass.class), objc.Sel("sharedRegistry"))
	return MLAllComputeDeviceRegistryFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLAllComputeDeviceRegistry/debugDescription
func (a MLAllComputeDeviceRegistry) DebugDescription() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/CoreML/MLAllComputeDeviceRegistry/description
func (a MLAllComputeDeviceRegistry) Description() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/CoreML/MLAllComputeDeviceRegistry/deviceRegistries
func (a MLAllComputeDeviceRegistry) DeviceRegistries() foundation.INSArray {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("deviceRegistries"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/MLAllComputeDeviceRegistry/hash
func (a MLAllComputeDeviceRegistry) Hash() uint64 {
	rv := objc.Send[uint64](a.ID, objc.Sel("hash"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLAllComputeDeviceRegistry/registeredComputeDevices
func (a MLAllComputeDeviceRegistry) RegisteredComputeDevices() foundation.INSArray {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("registeredComputeDevices"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/MLAllComputeDeviceRegistry/superclass
func (a MLAllComputeDeviceRegistry) Superclass() objc.Class {
	rv := objc.Send[objc.Class](a.ID, objc.Sel("superclass"))
	return rv
}
