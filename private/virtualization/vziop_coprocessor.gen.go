// Code generated from Apple documentation for virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [VZIOPCoprocessor] class.
var (
	_VZIOPCoprocessorClass     VZIOPCoprocessorClass
	_VZIOPCoprocessorClassOnce sync.Once
)

func getVZIOPCoprocessorClass() VZIOPCoprocessorClass {
	_VZIOPCoprocessorClassOnce.Do(func() {
		_VZIOPCoprocessorClass = VZIOPCoprocessorClass{class: objc.GetClass("_VZIOPCoprocessor")}
	})
	return _VZIOPCoprocessorClass
}

// GetVZIOPCoprocessorClass returns the class object for _VZIOPCoprocessor.
func GetVZIOPCoprocessorClass() VZIOPCoprocessorClass {
	return getVZIOPCoprocessorClass()
}

type VZIOPCoprocessorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZIOPCoprocessorClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZIOPCoprocessorClass) Alloc() VZIOPCoprocessor {
	rv := objc.SendIfResponds[VZIOPCoprocessor](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

type VZIOPCoprocessor struct {
	VZCoprocessor
}

// VZIOPCoprocessorFromID constructs a [VZIOPCoprocessor] from an objc.ID.
func VZIOPCoprocessorFromID(id objc.ID) VZIOPCoprocessor {
	return VZIOPCoprocessor{VZCoprocessor: VZCoprocessorFromID(id)}
}

// Ensure VZIOPCoprocessor implements IVZIOPCoprocessor.
var _ IVZIOPCoprocessor = VZIOPCoprocessor{}

// An interface definition for the [VZIOPCoprocessor] class.
type IVZIOPCoprocessor interface {
	IVZCoprocessor
}

// Init initializes the instance.
func (v VZIOPCoprocessor) Init() VZIOPCoprocessor {
	rv := objc.SendIfResponds[VZIOPCoprocessor](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZIOPCoprocessor) Autorelease() VZIOPCoprocessor {
	rv := objc.SendIfResponds[VZIOPCoprocessor](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZIOPCoprocessor creates a new VZIOPCoprocessor instance.
func NewVZIOPCoprocessor() VZIOPCoprocessor {
	class := getVZIOPCoprocessorClass()
	rv := objc.SendIfResponds[VZIOPCoprocessor](objc.ID(class.class), objc.Sel("new"))
	return rv
}
