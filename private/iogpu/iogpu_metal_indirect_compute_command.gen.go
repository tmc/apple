// Code generated from Apple documentation for iogpu. DO NOT EDIT.

package iogpu

import (
	"sync"

	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [IOGPUMetalIndirectComputeCommand] class.
var (
	_IOGPUMetalIndirectComputeCommandClass     IOGPUMetalIndirectComputeCommandClass
	_IOGPUMetalIndirectComputeCommandClassOnce sync.Once
)

func getIOGPUMetalIndirectComputeCommandClass() IOGPUMetalIndirectComputeCommandClass {
	_IOGPUMetalIndirectComputeCommandClassOnce.Do(func() {
		_IOGPUMetalIndirectComputeCommandClass = IOGPUMetalIndirectComputeCommandClass{class: objc.GetClass("IOGPUMetalIndirectComputeCommand")}
	})
	return _IOGPUMetalIndirectComputeCommandClass
}

// GetIOGPUMetalIndirectComputeCommandClass returns the class object for IOGPUMetalIndirectComputeCommand.
func GetIOGPUMetalIndirectComputeCommandClass() IOGPUMetalIndirectComputeCommandClass {
	return getIOGPUMetalIndirectComputeCommandClass()
}

type IOGPUMetalIndirectComputeCommandClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ic IOGPUMetalIndirectComputeCommandClass) Class() objc.Class {
	return ic.class
}

// Alloc allocates memory for a new instance of the class.
func (ic IOGPUMetalIndirectComputeCommandClass) Alloc() IOGPUMetalIndirectComputeCommand {
	rv := objc.SendIfResponds[IOGPUMetalIndirectComputeCommand](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

type IOGPUMetalIndirectComputeCommand struct {
	metal.MTLIndirectComputeCommandObject
}

// IOGPUMetalIndirectComputeCommandFromID constructs a [IOGPUMetalIndirectComputeCommand] from an objc.ID.
func IOGPUMetalIndirectComputeCommandFromID(id objc.ID) IOGPUMetalIndirectComputeCommand {
	return IOGPUMetalIndirectComputeCommand{MTLIndirectComputeCommandObject: metal.MTLIndirectComputeCommandObjectFromID(id)}
}

// Ensure IOGPUMetalIndirectComputeCommand implements IIOGPUMetalIndirectComputeCommand.
var _ IIOGPUMetalIndirectComputeCommand = IOGPUMetalIndirectComputeCommand{}

// An interface definition for the [IOGPUMetalIndirectComputeCommand] class.
type IIOGPUMetalIndirectComputeCommand interface {
	metal.MTLIndirectComputeCommand
}

// Init initializes the instance.
func (i IOGPUMetalIndirectComputeCommand) Init() IOGPUMetalIndirectComputeCommand {
	rv := objc.SendIfResponds[IOGPUMetalIndirectComputeCommand](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i IOGPUMetalIndirectComputeCommand) Autorelease() IOGPUMetalIndirectComputeCommand {
	rv := objc.SendIfResponds[IOGPUMetalIndirectComputeCommand](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewIOGPUMetalIndirectComputeCommand creates a new IOGPUMetalIndirectComputeCommand instance.
func NewIOGPUMetalIndirectComputeCommand() IOGPUMetalIndirectComputeCommand {
	class := getIOGPUMetalIndirectComputeCommandClass()
	rv := objc.SendIfResponds[IOGPUMetalIndirectComputeCommand](objc.ID(class.class), objc.Sel("new"))
	return rv
}
