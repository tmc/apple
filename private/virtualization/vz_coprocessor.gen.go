// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZCoprocessor] class.
var (
	_VZCoprocessorClass     VZCoprocessorClass
	_VZCoprocessorClassOnce sync.Once
)

func getVZCoprocessorClass() VZCoprocessorClass {
	_VZCoprocessorClassOnce.Do(func() {
		_VZCoprocessorClass = VZCoprocessorClass{class: objc.GetClass("_VZCoprocessor")}
	})
	return _VZCoprocessorClass
}

// GetVZCoprocessorClass returns the class object for _VZCoprocessor.
func GetVZCoprocessorClass() VZCoprocessorClass {
	return getVZCoprocessorClass()
}

type VZCoprocessorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZCoprocessorClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZCoprocessorClass) Alloc() VZCoprocessor {
	rv := objc.Send[VZCoprocessor](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [VZCoprocessor._init]
// See: https://developer.apple.com/documentation/Virtualization/_VZCoprocessor
type VZCoprocessor struct {
	objectivec.Object
}

// VZCoprocessorFromID constructs a [VZCoprocessor] from an objc.ID.
func VZCoprocessorFromID(id objc.ID) VZCoprocessor {
	return VZCoprocessor{objectivec.Object{ID: id}}
}
// Ensure VZCoprocessor implements IVZCoprocessor.
var _ IVZCoprocessor = VZCoprocessor{}

// An interface definition for the [VZCoprocessor] class.
//
// # Methods
//
//   - [IVZCoprocessor._init]
//
// See: https://developer.apple.com/documentation/Virtualization/_VZCoprocessor
type IVZCoprocessor interface {
	objectivec.IObject

	// Topic: Methods

	_init() objectivec.IObject
}

// Init initializes the instance.
func (v VZCoprocessor) Init() VZCoprocessor {
	rv := objc.Send[VZCoprocessor](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZCoprocessor) Autorelease() VZCoprocessor {
	rv := objc.Send[VZCoprocessor](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZCoprocessor creates a new VZCoprocessor instance.
func NewVZCoprocessor() VZCoprocessor {
	class := getVZCoprocessorClass()
	rv := objc.Send[VZCoprocessor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/_VZCoprocessor/_init
func (v VZCoprocessor) _init() objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("_init"))
	return objectivec.Object{ID: rv}
}

