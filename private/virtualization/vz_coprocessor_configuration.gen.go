// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
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
	rv := objc.SendIfResponds[VZCoprocessorConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [VZCoprocessorConfiguration._init]
//   - [VZCoprocessorConfiguration.MakeCoprocessorForVirtualMachineCoprocessorIndex]
//   - [VZCoprocessorConfiguration.DebugDescription]
//   - [VZCoprocessorConfiguration.Description]
//   - [VZCoprocessorConfiguration.Hash]
//   - [VZCoprocessorConfiguration.Superclass]
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
//   - [IVZCoprocessorConfiguration._init]
//   - [IVZCoprocessorConfiguration.MakeCoprocessorForVirtualMachineCoprocessorIndex]
//   - [IVZCoprocessorConfiguration.DebugDescription]
//   - [IVZCoprocessorConfiguration.Description]
//   - [IVZCoprocessorConfiguration.Hash]
//   - [IVZCoprocessorConfiguration.Superclass]
type IVZCoprocessorConfiguration interface {
	objectivec.IObject

	// Topic: Methods

	_init() objectivec.IObject
	MakeCoprocessorForVirtualMachineCoprocessorIndex(machine objectivec.IObject, index uint64) objectivec.IObject
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objectivec.Class
}

// Init initializes the instance.
func (v VZCoprocessorConfiguration) Init() VZCoprocessorConfiguration {
	rv := objc.SendIfResponds[VZCoprocessorConfiguration](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZCoprocessorConfiguration) Autorelease() VZCoprocessorConfiguration {
	rv := objc.SendIfResponds[VZCoprocessorConfiguration](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZCoprocessorConfiguration creates a new VZCoprocessorConfiguration instance.
func NewVZCoprocessorConfiguration() VZCoprocessorConfiguration {
	class := getVZCoprocessorConfigurationClass()
	rv := objc.SendIfResponds[VZCoprocessorConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (v VZCoprocessorConfiguration) _init() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](v.ID, objc.Sel("_init"))
	return objectivec.Object{ID: rv}
}
func (v VZCoprocessorConfiguration) MakeCoprocessorForVirtualMachineCoprocessorIndex(machine objectivec.IObject, index uint64) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](v.ID, objc.Sel("makeCoprocessorForVirtualMachine:coprocessorIndex:"), machine, index)
	return objectivec.Object{ID: rv}
}

func (v VZCoprocessorConfiguration) DebugDescription() string {
	rv := objc.SendIfResponds[objc.ID](v.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
func (v VZCoprocessorConfiguration) Description() string {
	rv := objc.SendIfResponds[objc.ID](v.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
func (v VZCoprocessorConfiguration) Hash() uint64 {
	rv := objc.SendIfResponds[uint64](v.ID, objc.Sel("hash"))
	return rv
}
func (v VZCoprocessorConfiguration) Superclass() objectivec.Class {
	rv := objc.SendIfResponds[objectivec.Class](v.ID, objc.Sel("superclass"))
	return objectivec.Class(rv)
}
