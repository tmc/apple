// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [EspressoContext] class.
var (
	_EspressoContextClass     EspressoContextClass
	_EspressoContextClassOnce sync.Once
)

func getEspressoContextClass() EspressoContextClass {
	_EspressoContextClassOnce.Do(func() {
		_EspressoContextClass = EspressoContextClass{class: objc.GetClass("EspressoContext")}
	})
	return _EspressoContextClass
}

// GetEspressoContextClass returns the class object for EspressoContext.
func GetEspressoContextClass() EspressoContextClass {
	return getEspressoContextClass()
}

type EspressoContextClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ec EspressoContextClass) Class() objc.Class {
	return ec.class
}

// Alloc allocates memory for a new instance of the class.
func (ec EspressoContextClass) Alloc() EspressoContext {
	rv := objc.Send[EspressoContext](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [EspressoContext.Ctx]
//   - [EspressoContext.Platform]
//   - [EspressoContext.Set_priorityLow_priority_max_ms_per_command_bufferGpu_priority]
//   - [EspressoContext.InitWithContext]
//   - [EspressoContext.InitWithDeviceAndWisdomParams]
//   - [EspressoContext.InitWithNetworkContext]
//   - [EspressoContext.InitWithPlatform]
// See: https://developer.apple.com/documentation/Espresso/EspressoContext
type EspressoContext struct {
	objectivec.Object
}

// EspressoContextFromID constructs a [EspressoContext] from an objc.ID.
func EspressoContextFromID(id objc.ID) EspressoContext {
	return EspressoContext{objectivec.Object{ID: id}}
}
// Ensure EspressoContext implements IEspressoContext.
var _ IEspressoContext = EspressoContext{}

// An interface definition for the [EspressoContext] class.
//
// # Methods
//
//   - [IEspressoContext.Ctx]
//   - [IEspressoContext.Platform]
//   - [IEspressoContext.Set_priorityLow_priority_max_ms_per_command_bufferGpu_priority]
//   - [IEspressoContext.InitWithContext]
//   - [IEspressoContext.InitWithDeviceAndWisdomParams]
//   - [IEspressoContext.InitWithNetworkContext]
//   - [IEspressoContext.InitWithPlatform]
//
// See: https://developer.apple.com/documentation/Espresso/EspressoContext
type IEspressoContext interface {
	objectivec.IObject

	// Topic: Methods

	Ctx() objectivec.IObject
	Platform() int
	Set_priorityLow_priority_max_ms_per_command_bufferGpu_priority(set_priority bool, low_priority_max_ms_per_command_buffer float32, gpu_priority uint32)
	InitWithContext(context objectivec.IObject) EspressoContext
	InitWithDeviceAndWisdomParams(device objectivec.IObject, params objectivec.IObject) EspressoContext
	InitWithNetworkContext(context objectivec.IObject) EspressoContext
	InitWithPlatform(platform int) EspressoContext
}

// Init initializes the instance.
func (e EspressoContext) Init() EspressoContext {
	rv := objc.Send[EspressoContext](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e EspressoContext) Autorelease() EspressoContext {
	rv := objc.Send[EspressoContext](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewEspressoContext creates a new EspressoContext instance.
func NewEspressoContext() EspressoContext {
	class := getEspressoContextClass()
	rv := objc.Send[EspressoContext](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/Espresso/EspressoContext/initWithContext:
func NewEspressoContextWithContext(context objectivec.IObject) EspressoContext {
	instance := getEspressoContextClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithContext:"), context)
	return EspressoContextFromID(rv)
}

//
// See: https://developer.apple.com/documentation/Espresso/EspressoContext/initWithDevice:andWisdomParams:
func NewEspressoContextWithDeviceAndWisdomParams(device objectivec.IObject, params objectivec.IObject) EspressoContext {
	instance := getEspressoContextClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:andWisdomParams:"), device, params)
	return EspressoContextFromID(rv)
}

//
// See: https://developer.apple.com/documentation/Espresso/EspressoContext/initWithNetworkContext:
func NewEspressoContextWithNetworkContext(context objectivec.IObject) EspressoContext {
	instance := getEspressoContextClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithNetworkContext:"), context)
	return EspressoContextFromID(rv)
}

//
// See: https://developer.apple.com/documentation/Espresso/EspressoContext/initWithPlatform:
func NewEspressoContextWithPlatform(platform int) EspressoContext {
	instance := getEspressoContextClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithPlatform:"), platform)
	return EspressoContextFromID(rv)
}

//
// See: https://developer.apple.com/documentation/Espresso/EspressoContext/set_priority:low_priority_max_ms_per_command_buffer:gpu_priority:
func (e EspressoContext) Set_priorityLow_priority_max_ms_per_command_bufferGpu_priority(set_priority bool, low_priority_max_ms_per_command_buffer float32, gpu_priority uint32) {
	objc.Send[objc.ID](e.ID, objc.Sel("set_priority:low_priority_max_ms_per_command_buffer:gpu_priority:"), set_priority, low_priority_max_ms_per_command_buffer, gpu_priority)
}
//
// See: https://developer.apple.com/documentation/Espresso/EspressoContext/initWithContext:
func (e EspressoContext) InitWithContext(context objectivec.IObject) EspressoContext {
	rv := objc.Send[EspressoContext](e.ID, objc.Sel("initWithContext:"), context)
	return rv
}
//
// See: https://developer.apple.com/documentation/Espresso/EspressoContext/initWithDevice:andWisdomParams:
func (e EspressoContext) InitWithDeviceAndWisdomParams(device objectivec.IObject, params objectivec.IObject) EspressoContext {
	rv := objc.Send[EspressoContext](e.ID, objc.Sel("initWithDevice:andWisdomParams:"), device, params)
	return rv
}
//
// See: https://developer.apple.com/documentation/Espresso/EspressoContext/initWithNetworkContext:
func (e EspressoContext) InitWithNetworkContext(context objectivec.IObject) EspressoContext {
	rv := objc.Send[EspressoContext](e.ID, objc.Sel("initWithNetworkContext:"), context)
	return rv
}
//
// See: https://developer.apple.com/documentation/Espresso/EspressoContext/initWithPlatform:
func (e EspressoContext) InitWithPlatform(platform int) EspressoContext {
	rv := objc.Send[EspressoContext](e.ID, objc.Sel("initWithPlatform:"), platform)
	return rv
}

// See: https://developer.apple.com/documentation/Espresso/EspressoContext/ctx
func (e EspressoContext) Ctx() objectivec.IObject {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("ctx"))
	return objectivec.Object{ID: rv}
}
// See: https://developer.apple.com/documentation/Espresso/EspressoContext/platform
func (e EspressoContext) Platform() int {
	rv := objc.Send[int](e.ID, objc.Sel("platform"))
	return rv
}

