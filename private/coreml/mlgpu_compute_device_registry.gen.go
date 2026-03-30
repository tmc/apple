// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLGPUComputeDeviceRegistry] class.
var (
	_MLGPUComputeDeviceRegistryClass     MLGPUComputeDeviceRegistryClass
	_MLGPUComputeDeviceRegistryClassOnce sync.Once
)

func getMLGPUComputeDeviceRegistryClass() MLGPUComputeDeviceRegistryClass {
	_MLGPUComputeDeviceRegistryClassOnce.Do(func() {
		_MLGPUComputeDeviceRegistryClass = MLGPUComputeDeviceRegistryClass{class: objc.GetClass("MLGPUComputeDeviceRegistry")}
	})
	return _MLGPUComputeDeviceRegistryClass
}

// GetMLGPUComputeDeviceRegistryClass returns the class object for MLGPUComputeDeviceRegistry.
func GetMLGPUComputeDeviceRegistryClass() MLGPUComputeDeviceRegistryClass {
	return getMLGPUComputeDeviceRegistryClass()
}

type MLGPUComputeDeviceRegistryClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLGPUComputeDeviceRegistryClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLGPUComputeDeviceRegistryClass) Alloc() MLGPUComputeDeviceRegistry {
	rv := objc.Send[MLGPUComputeDeviceRegistry](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLGPUComputeDeviceRegistry.AvailableGPUDevices]
//   - [MLGPUComputeDeviceRegistry.MetalDeviceObserver]
//   - [MLGPUComputeDeviceRegistry.ObservationToken]
//   - [MLGPUComputeDeviceRegistry.SetObservationToken]
//   - [MLGPUComputeDeviceRegistry.PendingChanges]
//   - [MLGPUComputeDeviceRegistry.RegisterGPUDevices]
//   - [MLGPUComputeDeviceRegistry.RegisteredComputeDevices]
//   - [MLGPUComputeDeviceRegistry.InitWithMetalDeviceObserver]
//   - [MLGPUComputeDeviceRegistry.DebugDescription]
//   - [MLGPUComputeDeviceRegistry.Description]
//   - [MLGPUComputeDeviceRegistry.Hash]
//   - [MLGPUComputeDeviceRegistry.Superclass]
//
// See: https://developer.apple.com/documentation/CoreML/MLGPUComputeDeviceRegistry
type MLGPUComputeDeviceRegistry struct {
	objectivec.Object
}

// MLGPUComputeDeviceRegistryFromID constructs a [MLGPUComputeDeviceRegistry] from an objc.ID.
func MLGPUComputeDeviceRegistryFromID(id objc.ID) MLGPUComputeDeviceRegistry {
	return MLGPUComputeDeviceRegistry{objectivec.Object{ID: id}}
}

// Ensure MLGPUComputeDeviceRegistry implements IMLGPUComputeDeviceRegistry.
var _ IMLGPUComputeDeviceRegistry = MLGPUComputeDeviceRegistry{}

// An interface definition for the [MLGPUComputeDeviceRegistry] class.
//
// # Methods
//
//   - [IMLGPUComputeDeviceRegistry.AvailableGPUDevices]
//   - [IMLGPUComputeDeviceRegistry.MetalDeviceObserver]
//   - [IMLGPUComputeDeviceRegistry.ObservationToken]
//   - [IMLGPUComputeDeviceRegistry.SetObservationToken]
//   - [IMLGPUComputeDeviceRegistry.PendingChanges]
//   - [IMLGPUComputeDeviceRegistry.RegisterGPUDevices]
//   - [IMLGPUComputeDeviceRegistry.RegisteredComputeDevices]
//   - [IMLGPUComputeDeviceRegistry.InitWithMetalDeviceObserver]
//   - [IMLGPUComputeDeviceRegistry.DebugDescription]
//   - [IMLGPUComputeDeviceRegistry.Description]
//   - [IMLGPUComputeDeviceRegistry.Hash]
//   - [IMLGPUComputeDeviceRegistry.Superclass]
//
// See: https://developer.apple.com/documentation/CoreML/MLGPUComputeDeviceRegistry
type IMLGPUComputeDeviceRegistry interface {
	objectivec.IObject

	// Topic: Methods

	AvailableGPUDevices() foundation.INSDictionary
	MetalDeviceObserver() IMLMetalDeviceObserver
	ObservationToken() objectivec.IObject
	SetObservationToken(value objectivec.IObject)
	PendingChanges() foundation.INSArray
	RegisterGPUDevices()
	RegisteredComputeDevices() foundation.INSArray
	InitWithMetalDeviceObserver(observer objectivec.IObject) MLGPUComputeDeviceRegistry
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objc.Class
}

// Init initializes the instance.
func (g MLGPUComputeDeviceRegistry) Init() MLGPUComputeDeviceRegistry {
	rv := objc.Send[MLGPUComputeDeviceRegistry](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g MLGPUComputeDeviceRegistry) Autorelease() MLGPUComputeDeviceRegistry {
	rv := objc.Send[MLGPUComputeDeviceRegistry](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLGPUComputeDeviceRegistry creates a new MLGPUComputeDeviceRegistry instance.
func NewMLGPUComputeDeviceRegistry() MLGPUComputeDeviceRegistry {
	class := getMLGPUComputeDeviceRegistryClass()
	rv := objc.Send[MLGPUComputeDeviceRegistry](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLGPUComputeDeviceRegistry/initWithMetalDeviceObserver:
func NewGPUComputeDeviceRegistryWithMetalDeviceObserver(observer objectivec.IObject) MLGPUComputeDeviceRegistry {
	instance := getMLGPUComputeDeviceRegistryClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithMetalDeviceObserver:"), observer)
	return MLGPUComputeDeviceRegistryFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLGPUComputeDeviceRegistry/registerGPUDevices
func (g MLGPUComputeDeviceRegistry) RegisterGPUDevices() {
	objc.Send[objc.ID](g.ID, objc.Sel("registerGPUDevices"))
}

// See: https://developer.apple.com/documentation/CoreML/MLGPUComputeDeviceRegistry/initWithMetalDeviceObserver:
func (g MLGPUComputeDeviceRegistry) InitWithMetalDeviceObserver(observer objectivec.IObject) MLGPUComputeDeviceRegistry {
	rv := objc.Send[MLGPUComputeDeviceRegistry](g.ID, objc.Sel("initWithMetalDeviceObserver:"), observer)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLGPUComputeDeviceRegistry/registryWithMetalDeviceObserver:
func (_MLGPUComputeDeviceRegistryClass MLGPUComputeDeviceRegistryClass) RegistryWithMetalDeviceObserver(observer objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLGPUComputeDeviceRegistryClass.class), objc.Sel("registryWithMetalDeviceObserver:"), observer)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/CoreML/MLGPUComputeDeviceRegistry/sharedRegistry
func (_MLGPUComputeDeviceRegistryClass MLGPUComputeDeviceRegistryClass) SharedRegistry() MLGPUComputeDeviceRegistry {
	rv := objc.Send[objc.ID](objc.ID(_MLGPUComputeDeviceRegistryClass.class), objc.Sel("sharedRegistry"))
	return MLGPUComputeDeviceRegistryFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLGPUComputeDeviceRegistry/availableGPUDevices
func (g MLGPUComputeDeviceRegistry) AvailableGPUDevices() foundation.INSDictionary {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("availableGPUDevices"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/MLGPUComputeDeviceRegistry/debugDescription
func (g MLGPUComputeDeviceRegistry) DebugDescription() string {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/CoreML/MLGPUComputeDeviceRegistry/description
func (g MLGPUComputeDeviceRegistry) Description() string {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/CoreML/MLGPUComputeDeviceRegistry/hash
func (g MLGPUComputeDeviceRegistry) Hash() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("hash"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLGPUComputeDeviceRegistry/metalDeviceObserver
func (g MLGPUComputeDeviceRegistry) MetalDeviceObserver() IMLMetalDeviceObserver {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("metalDeviceObserver"))
	return MLMetalDeviceObserverFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/MLGPUComputeDeviceRegistry/observationToken
func (g MLGPUComputeDeviceRegistry) ObservationToken() objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("observationToken"))
	return objectivec.Object{ID: rv}
}
func (g MLGPUComputeDeviceRegistry) SetObservationToken(value objectivec.IObject) {
	objc.Send[struct{}](g.ID, objc.Sel("setObservationToken:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLGPUComputeDeviceRegistry/pendingChanges
func (g MLGPUComputeDeviceRegistry) PendingChanges() foundation.INSArray {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("pendingChanges"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/MLGPUComputeDeviceRegistry/registeredComputeDevices
func (g MLGPUComputeDeviceRegistry) RegisteredComputeDevices() foundation.INSArray {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("registeredComputeDevices"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/MLGPUComputeDeviceRegistry/superclass
func (g MLGPUComputeDeviceRegistry) Superclass() objc.Class {
	rv := objc.Send[objc.Class](g.ID, objc.Sel("superclass"))
	return rv
}
