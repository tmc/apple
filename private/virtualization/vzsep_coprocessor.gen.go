// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [VZSEPCoprocessor] class.
var (
	_VZSEPCoprocessorClass     VZSEPCoprocessorClass
	_VZSEPCoprocessorClassOnce sync.Once
)

func getVZSEPCoprocessorClass() VZSEPCoprocessorClass {
	_VZSEPCoprocessorClassOnce.Do(func() {
		_VZSEPCoprocessorClass = VZSEPCoprocessorClass{class: objc.GetClass("_VZSEPCoprocessor")}
	})
	return _VZSEPCoprocessorClass
}

// GetVZSEPCoprocessorClass returns the class object for _VZSEPCoprocessor.
func GetVZSEPCoprocessorClass() VZSEPCoprocessorClass {
	return getVZSEPCoprocessorClass()
}

type VZSEPCoprocessorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZSEPCoprocessorClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZSEPCoprocessorClass) Alloc() VZSEPCoprocessor {
	rv := objc.Send[VZSEPCoprocessor](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [VZSEPCoprocessor.DebugStub]
//
// See: https://developer.apple.com/documentation/Virtualization/_VZSEPCoprocessor
type VZSEPCoprocessor struct {
	VZCoprocessor
}

// VZSEPCoprocessorFromID constructs a [VZSEPCoprocessor] from an objc.ID.
func VZSEPCoprocessorFromID(id objc.ID) VZSEPCoprocessor {
	return VZSEPCoprocessor{VZCoprocessor: VZCoprocessorFromID(id)}
}

// Ensure VZSEPCoprocessor implements IVZSEPCoprocessor.
var _ IVZSEPCoprocessor = VZSEPCoprocessor{}

// An interface definition for the [VZSEPCoprocessor] class.
//
// # Methods
//
//   - [IVZSEPCoprocessor.DebugStub]
//
// See: https://developer.apple.com/documentation/Virtualization/_VZSEPCoprocessor
type IVZSEPCoprocessor interface {
	IVZCoprocessor

	// Topic: Methods

	DebugStub() *VZDebugStub
}

// Init initializes the instance.
func (v VZSEPCoprocessor) Init() VZSEPCoprocessor {
	rv := objc.Send[VZSEPCoprocessor](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZSEPCoprocessor) Autorelease() VZSEPCoprocessor {
	rv := objc.Send[VZSEPCoprocessor](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZSEPCoprocessor creates a new VZSEPCoprocessor instance.
func NewVZSEPCoprocessor() VZSEPCoprocessor {
	class := getVZSEPCoprocessorClass()
	rv := objc.Send[VZSEPCoprocessor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/_VZSEPCoprocessor/debugStub
func (v VZSEPCoprocessor) DebugStub() *VZDebugStub {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("debugStub"))
	if rv == 0 {
		return nil
	}
	val := VZDebugStubFromID(objc.ID(rv))
	return &val
}
