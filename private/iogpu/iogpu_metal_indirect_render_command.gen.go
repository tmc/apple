// Code generated from Apple documentation for iogpu. DO NOT EDIT.

package iogpu

import (
	"sync"

	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
)

// The class instance for the [IOGPUMetalIndirectRenderCommand] class.
var (
	_IOGPUMetalIndirectRenderCommandClass     IOGPUMetalIndirectRenderCommandClass
	_IOGPUMetalIndirectRenderCommandClassOnce sync.Once
)

func getIOGPUMetalIndirectRenderCommandClass() IOGPUMetalIndirectRenderCommandClass {
	_IOGPUMetalIndirectRenderCommandClassOnce.Do(func() {
		_IOGPUMetalIndirectRenderCommandClass = IOGPUMetalIndirectRenderCommandClass{class: objc.GetClass("IOGPUMetalIndirectRenderCommand")}
	})
	return _IOGPUMetalIndirectRenderCommandClass
}

// GetIOGPUMetalIndirectRenderCommandClass returns the class object for IOGPUMetalIndirectRenderCommand.
func GetIOGPUMetalIndirectRenderCommandClass() IOGPUMetalIndirectRenderCommandClass {
	return getIOGPUMetalIndirectRenderCommandClass()
}

type IOGPUMetalIndirectRenderCommandClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ic IOGPUMetalIndirectRenderCommandClass) Class() objc.Class {
	return ic.class
}

// Alloc allocates memory for a new instance of the class.
func (ic IOGPUMetalIndirectRenderCommandClass) Alloc() IOGPUMetalIndirectRenderCommand {
	rv := objc.SendIfResponds[IOGPUMetalIndirectRenderCommand](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

type IOGPUMetalIndirectRenderCommand struct {
	metal.MTLIndirectRenderCommandObject
}

// IOGPUMetalIndirectRenderCommandFromID constructs a [IOGPUMetalIndirectRenderCommand] from an objc.ID.
func IOGPUMetalIndirectRenderCommandFromID(id objc.ID) IOGPUMetalIndirectRenderCommand {
	return IOGPUMetalIndirectRenderCommand{MTLIndirectRenderCommandObject: metal.MTLIndirectRenderCommandObjectFromID(id)}
}

// Ensure IOGPUMetalIndirectRenderCommand implements IIOGPUMetalIndirectRenderCommand.
var _ IIOGPUMetalIndirectRenderCommand = IOGPUMetalIndirectRenderCommand{}

// An interface definition for the [IOGPUMetalIndirectRenderCommand] class.
type IIOGPUMetalIndirectRenderCommand interface {
	metal.MTLIndirectRenderCommand
}

// Init initializes the instance.
func (i IOGPUMetalIndirectRenderCommand) Init() IOGPUMetalIndirectRenderCommand {
	rv := objc.SendIfResponds[IOGPUMetalIndirectRenderCommand](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i IOGPUMetalIndirectRenderCommand) Autorelease() IOGPUMetalIndirectRenderCommand {
	rv := objc.SendIfResponds[IOGPUMetalIndirectRenderCommand](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewIOGPUMetalIndirectRenderCommand creates a new IOGPUMetalIndirectRenderCommand instance.
func NewIOGPUMetalIndirectRenderCommand() IOGPUMetalIndirectRenderCommand {
	class := getIOGPUMetalIndirectRenderCommandClass()
	rv := objc.SendIfResponds[IOGPUMetalIndirectRenderCommand](objc.ID(class.class), objc.Sel("new"))
	return rv
}
