// Code generated from Apple documentation for iogpu. DO NOT EDIT.

package iogpu

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [IOGPUMetalTextureLayout] class.
var (
	_IOGPUMetalTextureLayoutClass     IOGPUMetalTextureLayoutClass
	_IOGPUMetalTextureLayoutClassOnce sync.Once
)

func getIOGPUMetalTextureLayoutClass() IOGPUMetalTextureLayoutClass {
	_IOGPUMetalTextureLayoutClassOnce.Do(func() {
		_IOGPUMetalTextureLayoutClass = IOGPUMetalTextureLayoutClass{class: objc.GetClass("IOGPUMetalTextureLayout")}
	})
	return _IOGPUMetalTextureLayoutClass
}

// GetIOGPUMetalTextureLayoutClass returns the class object for IOGPUMetalTextureLayout.
func GetIOGPUMetalTextureLayoutClass() IOGPUMetalTextureLayoutClass {
	return getIOGPUMetalTextureLayoutClass()
}

type IOGPUMetalTextureLayoutClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ic IOGPUMetalTextureLayoutClass) Class() objc.Class {
	return ic.class
}

// Alloc allocates memory for a new instance of the class.
func (ic IOGPUMetalTextureLayoutClass) Alloc() IOGPUMetalTextureLayout {
	rv := objc.SendIfResponds[IOGPUMetalTextureLayout](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [IOGPUMetalTextureLayout.InitWithDeviceDescriptor]
type IOGPUMetalTextureLayout struct {
	objectivec.Object
}

// IOGPUMetalTextureLayoutFromID constructs a [IOGPUMetalTextureLayout] from an objc.ID.
func IOGPUMetalTextureLayoutFromID(id objc.ID) IOGPUMetalTextureLayout {
	return IOGPUMetalTextureLayout{objectivec.Object{ID: id}}
}

// Ensure IOGPUMetalTextureLayout implements IIOGPUMetalTextureLayout.
var _ IIOGPUMetalTextureLayout = IOGPUMetalTextureLayout{}

// An interface definition for the [IOGPUMetalTextureLayout] class.
//
// # Methods
//
//   - [IIOGPUMetalTextureLayout.InitWithDeviceDescriptor]
type IIOGPUMetalTextureLayout interface {
	objectivec.IObject

	// Topic: Methods

	InitWithDeviceDescriptor(device objectivec.IObject, descriptor objectivec.IObject) IOGPUMetalTextureLayout
}

// Init initializes the instance.
func (i IOGPUMetalTextureLayout) Init() IOGPUMetalTextureLayout {
	rv := objc.SendIfResponds[IOGPUMetalTextureLayout](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i IOGPUMetalTextureLayout) Autorelease() IOGPUMetalTextureLayout {
	rv := objc.SendIfResponds[IOGPUMetalTextureLayout](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewIOGPUMetalTextureLayout creates a new IOGPUMetalTextureLayout instance.
func NewIOGPUMetalTextureLayout() IOGPUMetalTextureLayout {
	class := getIOGPUMetalTextureLayoutClass()
	rv := objc.SendIfResponds[IOGPUMetalTextureLayout](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewGPUMetalTextureLayoutWithDeviceDescriptor(device objectivec.IObject, descriptor objectivec.IObject) IOGPUMetalTextureLayout {
	instance := getIOGPUMetalTextureLayoutClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithDevice:descriptor:"), device, descriptor)
	return IOGPUMetalTextureLayoutFromID(rv)
}

func (i IOGPUMetalTextureLayout) InitWithDeviceDescriptor(device objectivec.IObject, descriptor objectivec.IObject) IOGPUMetalTextureLayout {
	rv := objc.SendIfResponds[IOGPUMetalTextureLayout](i.ID, objc.Sel("initWithDevice:descriptor:"), device, descriptor)
	return rv
}
