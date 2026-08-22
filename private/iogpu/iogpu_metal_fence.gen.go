// Code generated from Apple documentation for iogpu. DO NOT EDIT.

package iogpu

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [IOGPUMetalFence] class.
var (
	_IOGPUMetalFenceClass     IOGPUMetalFenceClass
	_IOGPUMetalFenceClassOnce sync.Once
)

func getIOGPUMetalFenceClass() IOGPUMetalFenceClass {
	_IOGPUMetalFenceClassOnce.Do(func() {
		_IOGPUMetalFenceClass = IOGPUMetalFenceClass{class: objc.GetClass("IOGPUMetalFence")}
	})
	return _IOGPUMetalFenceClass
}

// GetIOGPUMetalFenceClass returns the class object for IOGPUMetalFence.
func GetIOGPUMetalFenceClass() IOGPUMetalFenceClass {
	return getIOGPUMetalFenceClass()
}

type IOGPUMetalFenceClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ic IOGPUMetalFenceClass) Class() objc.Class {
	return ic.class
}

// Alloc allocates memory for a new instance of the class.
func (ic IOGPUMetalFenceClass) Alloc() IOGPUMetalFence {
	rv := objc.SendIfResponds[IOGPUMetalFence](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [IOGPUMetalFence.InitWithDevice]
//   - [IOGPUMetalFence.DebugDescription]
//   - [IOGPUMetalFence.Description]
//   - [IOGPUMetalFence.Hash]
//   - [IOGPUMetalFence.Superclass]
type IOGPUMetalFence struct {
	metal.MTLFenceObject
}

// IOGPUMetalFenceFromID constructs a [IOGPUMetalFence] from an objc.ID.
func IOGPUMetalFenceFromID(id objc.ID) IOGPUMetalFence {
	return IOGPUMetalFence{MTLFenceObject: metal.MTLFenceObjectFromID(id)}
}

// Ensure IOGPUMetalFence implements IIOGPUMetalFence.
var _ IIOGPUMetalFence = IOGPUMetalFence{}

// An interface definition for the [IOGPUMetalFence] class.
//
// # Methods
//
//   - [IIOGPUMetalFence.InitWithDevice]
//   - [IIOGPUMetalFence.DebugDescription]
//   - [IIOGPUMetalFence.Description]
//   - [IIOGPUMetalFence.Hash]
//   - [IIOGPUMetalFence.Superclass]
type IIOGPUMetalFence interface {
	metal.MTLFence

	// Topic: Methods

	InitWithDevice(device objectivec.IObject) IOGPUMetalFence
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objectivec.Class
}

// Init initializes the instance.
func (i IOGPUMetalFence) Init() IOGPUMetalFence {
	rv := objc.SendIfResponds[IOGPUMetalFence](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i IOGPUMetalFence) Autorelease() IOGPUMetalFence {
	rv := objc.SendIfResponds[IOGPUMetalFence](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewIOGPUMetalFence creates a new IOGPUMetalFence instance.
func NewIOGPUMetalFence() IOGPUMetalFence {
	class := getIOGPUMetalFenceClass()
	rv := objc.SendIfResponds[IOGPUMetalFence](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewGPUMetalFenceWithDevice(device objectivec.IObject) IOGPUMetalFence {
	instance := getIOGPUMetalFenceClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return IOGPUMetalFenceFromID(rv)
}

func (i IOGPUMetalFence) InitWithDevice(device objectivec.IObject) IOGPUMetalFence {
	rv := objc.SendIfResponds[IOGPUMetalFence](i.ID, objc.Sel("initWithDevice:"), device)
	return rv
}

func (i IOGPUMetalFence) DebugDescription() string {
	rv := objc.SendIfResponds[objc.ID](i.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
func (i IOGPUMetalFence) Description() string {
	rv := objc.SendIfResponds[objc.ID](i.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
func (i IOGPUMetalFence) Hash() uint64 {
	rv := objc.SendIfResponds[uint64](i.ID, objc.Sel("hash"))
	return rv
}
func (i IOGPUMetalFence) Superclass() objectivec.Class {
	rv := objc.SendIfResponds[objectivec.Class](i.ID, objc.Sel("superclass"))
	return objectivec.Class(rv)
}
