// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

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
type IMLGPUComputeDeviceRegistry interface {
	objectivec.IObject

	// Topic: Methods

	AvailableGPUDevices() foundation.INSDictionary
	MetalDeviceObserver() IMLMetalDeviceObserver
	ObservationToken() unsafe.Pointer
	SetObservationToken(value unsafe.Pointer)
	PendingChanges() foundation.INSArray
	RegisterGPUDevices()
	RegisteredComputeDevices() foundation.INSArray
	InitWithMetalDeviceObserver(observer objectivec.IObject) MLGPUComputeDeviceRegistry
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objectivec.Class
}

// Init initializes the instance.
func (m MLGPUComputeDeviceRegistry) Init() MLGPUComputeDeviceRegistry {
	rv := objc.Send[MLGPUComputeDeviceRegistry](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLGPUComputeDeviceRegistry) Autorelease() MLGPUComputeDeviceRegistry {
	rv := objc.Send[MLGPUComputeDeviceRegistry](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLGPUComputeDeviceRegistry creates a new MLGPUComputeDeviceRegistry instance.
func NewMLGPUComputeDeviceRegistry() MLGPUComputeDeviceRegistry {
	class := getMLGPUComputeDeviceRegistryClass()
	rv := objc.Send[MLGPUComputeDeviceRegistry](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewGPUComputeDeviceRegistryWithMetalDeviceObserver(observer objectivec.IObject) MLGPUComputeDeviceRegistry {
	instance := getMLGPUComputeDeviceRegistryClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithMetalDeviceObserver:"), observer)
	return MLGPUComputeDeviceRegistryFromID(rv)
}

func (m MLGPUComputeDeviceRegistry) RegisterGPUDevices() {
	objc.Send[objc.ID](m.ID, objc.Sel("registerGPUDevices"))
}
func (m MLGPUComputeDeviceRegistry) InitWithMetalDeviceObserver(observer objectivec.IObject) MLGPUComputeDeviceRegistry {
	rv := objc.Send[MLGPUComputeDeviceRegistry](m.ID, objc.Sel("initWithMetalDeviceObserver:"), observer)
	return rv
}

func (_MLGPUComputeDeviceRegistryClass MLGPUComputeDeviceRegistryClass) RegistryWithMetalDeviceObserver(observer objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLGPUComputeDeviceRegistryClass.class), objc.Sel("registryWithMetalDeviceObserver:"), observer)
	return objectivec.Object{ID: rv}
}
func (_MLGPUComputeDeviceRegistryClass MLGPUComputeDeviceRegistryClass) SharedRegistry() MLGPUComputeDeviceRegistry {
	rv := objc.Send[objc.ID](objc.ID(_MLGPUComputeDeviceRegistryClass.class), objc.Sel("sharedRegistry"))
	return MLGPUComputeDeviceRegistryFromID(rv)
}

func (m MLGPUComputeDeviceRegistry) AvailableGPUDevices() foundation.INSDictionary {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("availableGPUDevices"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (m MLGPUComputeDeviceRegistry) DebugDescription() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
func (m MLGPUComputeDeviceRegistry) Description() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
func (m MLGPUComputeDeviceRegistry) Hash() uint64 {
	rv := objc.Send[uint64](m.ID, objc.Sel("hash"))
	return rv
}
func (m MLGPUComputeDeviceRegistry) MetalDeviceObserver() IMLMetalDeviceObserver {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("metalDeviceObserver"))
	return MLMetalDeviceObserverFromID(objc.ID(rv))
}
func (m MLGPUComputeDeviceRegistry) ObservationToken() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m.ID, objc.Sel("observationToken"))
	return rv
}
func (m MLGPUComputeDeviceRegistry) SetObservationToken(value unsafe.Pointer) {
	objc.Send[struct{}](m.ID, objc.Sel("setObservationToken:"), value)
}
func (m MLGPUComputeDeviceRegistry) PendingChanges() foundation.INSArray {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("pendingChanges"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (m MLGPUComputeDeviceRegistry) RegisteredComputeDevices() foundation.INSArray {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("registeredComputeDevices"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (m MLGPUComputeDeviceRegistry) Superclass() objectivec.Class {
	rv := objc.Send[objectivec.Class](m.ID, objc.Sel("superclass"))
	return objectivec.Class(rv)
}
