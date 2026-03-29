// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZCoprocessorConfiguration] class.
var (
	_VZCoprocessorConfigurationClass     VZCoprocessorConfigurationClass
	_VZCoprocessorConfigurationClassOnce sync.Once
)

func getVZCoprocessorConfigurationClass() VZCoprocessorConfigurationClass {
	_VZCoprocessorConfigurationClassOnce.Do(func() {
		_VZCoprocessorConfigurationClass = VZCoprocessorConfigurationClass{class: objc.GetClass("_VZCoprocessorConfiguration")}
	})
	return _VZCoprocessorConfigurationClass
}

// GetVZCoprocessorConfigurationClass returns the class object for _VZCoprocessorConfiguration.
func GetVZCoprocessorConfigurationClass() VZCoprocessorConfigurationClass {
	return getVZCoprocessorConfigurationClass()
}

type VZCoprocessorConfigurationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZCoprocessorConfigurationClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZCoprocessorConfigurationClass) Alloc() VZCoprocessorConfiguration {
	rv := objc.Send[VZCoprocessorConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [VZCoprocessorConfiguration._coprocessor]
//   - [VZCoprocessorConfiguration._init]
//   - [VZCoprocessorConfiguration.EncodeWithEncoder]
//   - [VZCoprocessorConfiguration.MakeCoprocessorForVirtualMachineCoprocessorIndex]
//   - [VZCoprocessorConfiguration.DebugDescription]
//   - [VZCoprocessorConfiguration.Description]
//   - [VZCoprocessorConfiguration.Hash]
//   - [VZCoprocessorConfiguration.Superclass]
// See: https://developer.apple.com/documentation/Virtualization/_VZCoprocessorConfiguration
type VZCoprocessorConfiguration struct {
	objectivec.Object
}

// VZCoprocessorConfigurationFromID constructs a [VZCoprocessorConfiguration] from an objc.ID.
func VZCoprocessorConfigurationFromID(id objc.ID) VZCoprocessorConfiguration {
	return VZCoprocessorConfiguration{objectivec.Object{ID: id}}
}
// Ensure VZCoprocessorConfiguration implements IVZCoprocessorConfiguration.
var _ IVZCoprocessorConfiguration = VZCoprocessorConfiguration{}

// An interface definition for the [VZCoprocessorConfiguration] class.
//
// # Methods
//
//   - [IVZCoprocessorConfiguration._coprocessor]
//   - [IVZCoprocessorConfiguration._init]
//   - [IVZCoprocessorConfiguration.EncodeWithEncoder]
//   - [IVZCoprocessorConfiguration.MakeCoprocessorForVirtualMachineCoprocessorIndex]
//   - [IVZCoprocessorConfiguration.DebugDescription]
//   - [IVZCoprocessorConfiguration.Description]
//   - [IVZCoprocessorConfiguration.Hash]
//   - [IVZCoprocessorConfiguration.Superclass]
//
// See: https://developer.apple.com/documentation/Virtualization/_VZCoprocessorConfiguration
type IVZCoprocessorConfiguration interface {
	objectivec.IObject

	// Topic: Methods

	_coprocessor() objectivec.IObject
	_init() objectivec.IObject
	EncodeWithEncoder(encoder objectivec.IObject) objectivec.IObject
	MakeCoprocessorForVirtualMachineCoprocessorIndex(machine objectivec.IObject, index uint64) objectivec.IObject
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objc.Class
}

// Init initializes the instance.
func (v VZCoprocessorConfiguration) Init() VZCoprocessorConfiguration {
	rv := objc.Send[VZCoprocessorConfiguration](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZCoprocessorConfiguration) Autorelease() VZCoprocessorConfiguration {
	rv := objc.Send[VZCoprocessorConfiguration](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZCoprocessorConfiguration creates a new VZCoprocessorConfiguration instance.
func NewVZCoprocessorConfiguration() VZCoprocessorConfiguration {
	class := getVZCoprocessorConfigurationClass()
	rv := objc.Send[VZCoprocessorConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/_VZCoprocessorConfiguration/_coprocessor
func (v VZCoprocessorConfiguration) _coprocessor() objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("_coprocessor"))
	return objectivec.Object{ID: rv}
}

// Coprocessor is an exported wrapper for the private method _coprocessor.
func (v VZCoprocessorConfiguration) Coprocessor() objectivec.IObject {
	return v._coprocessor()
}
// See: https://developer.apple.com/documentation/Virtualization/_VZCoprocessorConfiguration/_init
func (v VZCoprocessorConfiguration) _init() objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("_init"))
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/Virtualization/_VZCoprocessorConfiguration/encodeWithEncoder:
func (v VZCoprocessorConfiguration) EncodeWithEncoder(encoder objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("encodeWithEncoder:"), encoder)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/Virtualization/_VZCoprocessorConfiguration/makeCoprocessorForVirtualMachine:coprocessorIndex:
func (v VZCoprocessorConfiguration) MakeCoprocessorForVirtualMachineCoprocessorIndex(machine objectivec.IObject, index uint64) objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("makeCoprocessorForVirtualMachine:coprocessorIndex:"), machine, index)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/Virtualization/_VZCoprocessorConfiguration/debugDescription
func (v VZCoprocessorConfiguration) DebugDescription() string {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/Virtualization/_VZCoprocessorConfiguration/description
func (v VZCoprocessorConfiguration) Description() string {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/Virtualization/_VZCoprocessorConfiguration/hash
func (v VZCoprocessorConfiguration) Hash() uint64 {
	rv := objc.Send[uint64](v.ID, objc.Sel("hash"))
	return rv
}
// See: https://developer.apple.com/documentation/Virtualization/_VZCoprocessorConfiguration/superclass
func (v VZCoprocessorConfiguration) Superclass() objc.Class {
	rv := objc.Send[objc.Class](v.ID, objc.Sel("superclass"))
	return rv
}

