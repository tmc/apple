// Code generated from Apple documentation for iogpu. DO NOT EDIT.

package iogpu

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [IOGPUMTLFence] class.
var (
	_IOGPUMTLFenceClass     IOGPUMTLFenceClass
	_IOGPUMTLFenceClassOnce sync.Once
)

func getIOGPUMTLFenceClass() IOGPUMTLFenceClass {
	_IOGPUMTLFenceClassOnce.Do(func() {
		_IOGPUMTLFenceClass = IOGPUMTLFenceClass{class: objc.GetClass("IOGPUMTLFence")}
	})
	return _IOGPUMTLFenceClass
}

// GetIOGPUMTLFenceClass returns the class object for IOGPUMTLFence.
func GetIOGPUMTLFenceClass() IOGPUMTLFenceClass {
	return getIOGPUMTLFenceClass()
}

type IOGPUMTLFenceClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ic IOGPUMTLFenceClass) Class() objc.Class {
	return ic.class
}

// Alloc allocates memory for a new instance of the class.
func (ic IOGPUMTLFenceClass) Alloc() IOGPUMTLFence {
	rv := objc.SendIfResponds[IOGPUMTLFence](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [IOGPUMTLFence.InitWithDevice]
type IOGPUMTLFence struct {
	objectivec.Object
}

// IOGPUMTLFenceFromID constructs a [IOGPUMTLFence] from an objc.ID.
func IOGPUMTLFenceFromID(id objc.ID) IOGPUMTLFence {
	return IOGPUMTLFence{objectivec.Object{ID: id}}
}

// Ensure IOGPUMTLFence implements IIOGPUMTLFence.
var _ IIOGPUMTLFence = IOGPUMTLFence{}

// An interface definition for the [IOGPUMTLFence] class.
//
// # Methods
//
//   - [IIOGPUMTLFence.InitWithDevice]
type IIOGPUMTLFence interface {
	objectivec.IObject

	// Topic: Methods

	InitWithDevice(device *uintptr) IOGPUMTLFence
}

// Init initializes the instance.
func (i IOGPUMTLFence) Init() IOGPUMTLFence {
	rv := objc.SendIfResponds[IOGPUMTLFence](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i IOGPUMTLFence) Autorelease() IOGPUMTLFence {
	rv := objc.SendIfResponds[IOGPUMTLFence](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewIOGPUMTLFence creates a new IOGPUMTLFence instance.
func NewIOGPUMTLFence() IOGPUMTLFence {
	class := getIOGPUMTLFenceClass()
	rv := objc.SendIfResponds[IOGPUMTLFence](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewGPUMTLFenceWithDevice(device *uintptr) IOGPUMTLFence {
	instance := getIOGPUMTLFenceClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return IOGPUMTLFenceFromID(rv)
}

func (i IOGPUMTLFence) InitWithDevice(device *uintptr) IOGPUMTLFence {
	rv := objc.SendIfResponds[IOGPUMTLFence](i.ID, objc.Sel("initWithDevice:"), device)
	return rv
}
