// Code generated from Apple documentation for iogpu. DO NOT EDIT.

package iogpu

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [IOGPUMetalDeviceAssertion] class.
var (
	_IOGPUMetalDeviceAssertionClass     IOGPUMetalDeviceAssertionClass
	_IOGPUMetalDeviceAssertionClassOnce sync.Once
)

func getIOGPUMetalDeviceAssertionClass() IOGPUMetalDeviceAssertionClass {
	_IOGPUMetalDeviceAssertionClassOnce.Do(func() {
		_IOGPUMetalDeviceAssertionClass = IOGPUMetalDeviceAssertionClass{class: objc.GetClass("IOGPUMetalDeviceAssertion")}
	})
	return _IOGPUMetalDeviceAssertionClass
}

// GetIOGPUMetalDeviceAssertionClass returns the class object for IOGPUMetalDeviceAssertion.
func GetIOGPUMetalDeviceAssertionClass() IOGPUMetalDeviceAssertionClass {
	return getIOGPUMetalDeviceAssertionClass()
}

type IOGPUMetalDeviceAssertionClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ic IOGPUMetalDeviceAssertionClass) Class() objc.Class {
	return ic.class
}

// Alloc allocates memory for a new instance of the class.
func (ic IOGPUMetalDeviceAssertionClass) Alloc() IOGPUMetalDeviceAssertion {
	rv := objc.SendIfResponds[IOGPUMetalDeviceAssertion](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [IOGPUMetalDeviceAssertion.InitWithDeviceAssertionTypeOptions]
type IOGPUMetalDeviceAssertion struct {
	objectivec.Object
}

// IOGPUMetalDeviceAssertionFromID constructs a [IOGPUMetalDeviceAssertion] from an objc.ID.
func IOGPUMetalDeviceAssertionFromID(id objc.ID) IOGPUMetalDeviceAssertion {
	return IOGPUMetalDeviceAssertion{objectivec.Object{ID: id}}
}

// Ensure IOGPUMetalDeviceAssertion implements IIOGPUMetalDeviceAssertion.
var _ IIOGPUMetalDeviceAssertion = IOGPUMetalDeviceAssertion{}

// An interface definition for the [IOGPUMetalDeviceAssertion] class.
//
// # Methods
//
//   - [IIOGPUMetalDeviceAssertion.InitWithDeviceAssertionTypeOptions]
type IIOGPUMetalDeviceAssertion interface {
	objectivec.IObject

	// Topic: Methods

	InitWithDeviceAssertionTypeOptions(device objectivec.IObject, type_ uint64, options uint64) IOGPUMetalDeviceAssertion
}

// Init initializes the instance.
func (i IOGPUMetalDeviceAssertion) Init() IOGPUMetalDeviceAssertion {
	rv := objc.SendIfResponds[IOGPUMetalDeviceAssertion](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i IOGPUMetalDeviceAssertion) Autorelease() IOGPUMetalDeviceAssertion {
	rv := objc.SendIfResponds[IOGPUMetalDeviceAssertion](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewIOGPUMetalDeviceAssertion creates a new IOGPUMetalDeviceAssertion instance.
func NewIOGPUMetalDeviceAssertion() IOGPUMetalDeviceAssertion {
	class := getIOGPUMetalDeviceAssertionClass()
	rv := objc.SendIfResponds[IOGPUMetalDeviceAssertion](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewGPUMetalDeviceAssertionWithDeviceAssertionTypeOptions(device objectivec.IObject, type_ uint64, options uint64) IOGPUMetalDeviceAssertion {
	instance := getIOGPUMetalDeviceAssertionClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithDevice:assertionType:options:"), device, type_, options)
	return IOGPUMetalDeviceAssertionFromID(rv)
}

func (i IOGPUMetalDeviceAssertion) InitWithDeviceAssertionTypeOptions(device objectivec.IObject, type_ uint64, options uint64) IOGPUMetalDeviceAssertion {
	rv := objc.SendIfResponds[IOGPUMetalDeviceAssertion](i.ID, objc.Sel("initWithDevice:assertionType:options:"), device, type_, options)
	return rv
}
