// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZDebugStubConfiguration] class.
var (
	_VZDebugStubConfigurationClass     VZDebugStubConfigurationClass
	_VZDebugStubConfigurationClassOnce sync.Once
)

func getVZDebugStubConfigurationClass() VZDebugStubConfigurationClass {
	_VZDebugStubConfigurationClassOnce.Do(func() {
		_VZDebugStubConfigurationClass = VZDebugStubConfigurationClass{class: objc.GetClass("_VZDebugStubConfiguration")}
	})
	return _VZDebugStubConfigurationClass
}

// GetVZDebugStubConfigurationClass returns the class object for _VZDebugStubConfiguration.
func GetVZDebugStubConfigurationClass() VZDebugStubConfigurationClass {
	return getVZDebugStubConfigurationClass()
}

type VZDebugStubConfigurationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZDebugStubConfigurationClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZDebugStubConfigurationClass) Alloc() VZDebugStubConfiguration {
	rv := objc.Send[VZDebugStubConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [VZDebugStubConfiguration._init]
//   - [VZDebugStubConfiguration.EncodeWithEncoder]
//   - [VZDebugStubConfiguration.MakeDebugStubForCoprocessor]
//   - [VZDebugStubConfiguration.MakeDebugStubForVirtualMachine]
//   - [VZDebugStubConfiguration.DebugDescription]
//   - [VZDebugStubConfiguration.Description]
//   - [VZDebugStubConfiguration.Hash]
//   - [VZDebugStubConfiguration.Superclass]
// See: https://developer.apple.com/documentation/Virtualization/_VZDebugStubConfiguration
type VZDebugStubConfiguration struct {
	objectivec.Object
}

// VZDebugStubConfigurationFromID constructs a [VZDebugStubConfiguration] from an objc.ID.
func VZDebugStubConfigurationFromID(id objc.ID) VZDebugStubConfiguration {
	return VZDebugStubConfiguration{objectivec.Object{ID: id}}
}
// Ensure VZDebugStubConfiguration implements IVZDebugStubConfiguration.
var _ IVZDebugStubConfiguration = VZDebugStubConfiguration{}

// An interface definition for the [VZDebugStubConfiguration] class.
//
// # Methods
//
//   - [IVZDebugStubConfiguration._init]
//   - [IVZDebugStubConfiguration.EncodeWithEncoder]
//   - [IVZDebugStubConfiguration.MakeDebugStubForCoprocessor]
//   - [IVZDebugStubConfiguration.MakeDebugStubForVirtualMachine]
//   - [IVZDebugStubConfiguration.DebugDescription]
//   - [IVZDebugStubConfiguration.Description]
//   - [IVZDebugStubConfiguration.Hash]
//   - [IVZDebugStubConfiguration.Superclass]
//
// See: https://developer.apple.com/documentation/Virtualization/_VZDebugStubConfiguration
type IVZDebugStubConfiguration interface {
	objectivec.IObject

	// Topic: Methods

	_init() objectivec.IObject
	EncodeWithEncoder(encoder objectivec.IObject) objectivec.IObject
	MakeDebugStubForCoprocessor() objectivec.IObject
	MakeDebugStubForVirtualMachine(machine objectivec.IObject) objectivec.IObject
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objc.Class
}

// Init initializes the instance.
func (v VZDebugStubConfiguration) Init() VZDebugStubConfiguration {
	rv := objc.Send[VZDebugStubConfiguration](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZDebugStubConfiguration) Autorelease() VZDebugStubConfiguration {
	rv := objc.Send[VZDebugStubConfiguration](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZDebugStubConfiguration creates a new VZDebugStubConfiguration instance.
func NewVZDebugStubConfiguration() VZDebugStubConfiguration {
	class := getVZDebugStubConfigurationClass()
	rv := objc.Send[VZDebugStubConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/_VZDebugStubConfiguration/_init
func (v VZDebugStubConfiguration) _init() objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("_init"))
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/Virtualization/_VZDebugStubConfiguration/encodeWithEncoder:
func (v VZDebugStubConfiguration) EncodeWithEncoder(encoder objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("encodeWithEncoder:"), encoder)
	return objectivec.Object{ID: rv}
}
// See: https://developer.apple.com/documentation/Virtualization/_VZDebugStubConfiguration/makeDebugStubForCoprocessor
func (v VZDebugStubConfiguration) MakeDebugStubForCoprocessor() objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("makeDebugStubForCoprocessor"))
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/Virtualization/_VZDebugStubConfiguration/makeDebugStubForVirtualMachine:
func (v VZDebugStubConfiguration) MakeDebugStubForVirtualMachine(machine objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("makeDebugStubForVirtualMachine:"), machine)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/Virtualization/_VZDebugStubConfiguration/debugDescription
func (v VZDebugStubConfiguration) DebugDescription() string {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/Virtualization/_VZDebugStubConfiguration/description
func (v VZDebugStubConfiguration) Description() string {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/Virtualization/_VZDebugStubConfiguration/hash
func (v VZDebugStubConfiguration) Hash() uint64 {
	rv := objc.Send[uint64](v.ID, objc.Sel("hash"))
	return rv
}
// See: https://developer.apple.com/documentation/Virtualization/_VZDebugStubConfiguration/superclass
func (v VZDebugStubConfiguration) Superclass() objc.Class {
	rv := objc.Send[objc.Class](v.ID, objc.Sel("superclass"))
	return rv
}

