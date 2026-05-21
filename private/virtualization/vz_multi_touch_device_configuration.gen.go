// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZMultiTouchDeviceConfiguration] class.
var (
	_VZMultiTouchDeviceConfigurationClass     VZMultiTouchDeviceConfigurationClass
	_VZMultiTouchDeviceConfigurationClassOnce sync.Once
)

func getVZMultiTouchDeviceConfigurationClass() VZMultiTouchDeviceConfigurationClass {
	_VZMultiTouchDeviceConfigurationClassOnce.Do(func() {
		_VZMultiTouchDeviceConfigurationClass = VZMultiTouchDeviceConfigurationClass{class: objc.GetClass("_VZMultiTouchDeviceConfiguration")}
	})
	return _VZMultiTouchDeviceConfigurationClass
}

// GetVZMultiTouchDeviceConfigurationClass returns the class object for _VZMultiTouchDeviceConfiguration.
func GetVZMultiTouchDeviceConfigurationClass() VZMultiTouchDeviceConfigurationClass {
	return getVZMultiTouchDeviceConfigurationClass()
}

type VZMultiTouchDeviceConfigurationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZMultiTouchDeviceConfigurationClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZMultiTouchDeviceConfigurationClass) Alloc() VZMultiTouchDeviceConfiguration {
	rv := objc.Send[VZMultiTouchDeviceConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [VZMultiTouchDeviceConfiguration._init]
//   - [VZMultiTouchDeviceConfiguration._multiTouchDevice]
//   - [VZMultiTouchDeviceConfiguration.AssociationIdentifier]
//   - [VZMultiTouchDeviceConfiguration.SetAssociationIdentifier]
//   - [VZMultiTouchDeviceConfiguration.MakeMultiTouchDeviceForVirtualMachineMultiTouchDeviceIndex]
//   - [VZMultiTouchDeviceConfiguration.DebugDescription]
//   - [VZMultiTouchDeviceConfiguration.Description]
//   - [VZMultiTouchDeviceConfiguration.Hash]
//   - [VZMultiTouchDeviceConfiguration.Superclass]
type VZMultiTouchDeviceConfiguration struct {
	objectivec.Object
}

// VZMultiTouchDeviceConfigurationFromID constructs a [VZMultiTouchDeviceConfiguration] from an objc.ID.
func VZMultiTouchDeviceConfigurationFromID(id objc.ID) VZMultiTouchDeviceConfiguration {
	return VZMultiTouchDeviceConfiguration{objectivec.Object{ID: id}}
}

// Ensure VZMultiTouchDeviceConfiguration implements IVZMultiTouchDeviceConfiguration.
var _ IVZMultiTouchDeviceConfiguration = VZMultiTouchDeviceConfiguration{}

// An interface definition for the [VZMultiTouchDeviceConfiguration] class.
//
// # Methods
//
//   - [IVZMultiTouchDeviceConfiguration._init]
//   - [IVZMultiTouchDeviceConfiguration._multiTouchDevice]
//   - [IVZMultiTouchDeviceConfiguration.AssociationIdentifier]
//   - [IVZMultiTouchDeviceConfiguration.SetAssociationIdentifier]
//   - [IVZMultiTouchDeviceConfiguration.MakeMultiTouchDeviceForVirtualMachineMultiTouchDeviceIndex]
//   - [IVZMultiTouchDeviceConfiguration.DebugDescription]
//   - [IVZMultiTouchDeviceConfiguration.Description]
//   - [IVZMultiTouchDeviceConfiguration.Hash]
//   - [IVZMultiTouchDeviceConfiguration.Superclass]
type IVZMultiTouchDeviceConfiguration interface {
	objectivec.IObject

	// Topic: Methods

	_init() objectivec.IObject
	_multiTouchDevice() unsafe.Pointer
	AssociationIdentifier() foundation.NSUUID
	SetAssociationIdentifier(value foundation.NSUUID)
	MakeMultiTouchDeviceForVirtualMachineMultiTouchDeviceIndex(machine objectivec.IObject, index uint64) objectivec.IObject
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objectivec.Class
}

// Init initializes the instance.
func (v VZMultiTouchDeviceConfiguration) Init() VZMultiTouchDeviceConfiguration {
	rv := objc.Send[VZMultiTouchDeviceConfiguration](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZMultiTouchDeviceConfiguration) Autorelease() VZMultiTouchDeviceConfiguration {
	rv := objc.Send[VZMultiTouchDeviceConfiguration](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZMultiTouchDeviceConfiguration creates a new VZMultiTouchDeviceConfiguration instance.
func NewVZMultiTouchDeviceConfiguration() VZMultiTouchDeviceConfiguration {
	class := getVZMultiTouchDeviceConfigurationClass()
	rv := objc.Send[VZMultiTouchDeviceConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (v VZMultiTouchDeviceConfiguration) _init() objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("_init"))
	return objectivec.Object{ID: rv}
}
func (v VZMultiTouchDeviceConfiguration) MakeMultiTouchDeviceForVirtualMachineMultiTouchDeviceIndex(machine objectivec.IObject, index uint64) objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("makeMultiTouchDeviceForVirtualMachine:multiTouchDeviceIndex:"), machine, index)
	return objectivec.Object{ID: rv}
}

func (v VZMultiTouchDeviceConfiguration) _multiTouchDevice() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v.ID, objc.Sel("_multiTouchDevice"))
	return rv
}

// CanMultiTouchDevice reports whether the receiver responds to the private selector _multiTouchDevice.
func (v VZMultiTouchDeviceConfiguration) CanMultiTouchDevice() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_multiTouchDevice"))
}

// MultiTouchDevice is an exported wrapper for the private property _multiTouchDevice.
func (v VZMultiTouchDeviceConfiguration) MultiTouchDevice() (unsafe.Pointer, error) {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_multiTouchDevice")) {
		return nil, &objc.UnrecognizedSelectorError{Selector: "_multiTouchDevice"}
	}
	return v._multiTouchDevice(), nil
}
func (v VZMultiTouchDeviceConfiguration) AssociationIdentifier() foundation.NSUUID {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("associationIdentifier"))
	return foundation.NSUUIDFromID(objc.ID(rv))
}
func (v VZMultiTouchDeviceConfiguration) SetAssociationIdentifier(value foundation.NSUUID) {
	objc.Send[struct{}](v.ID, objc.Sel("setAssociationIdentifier:"), value)
}
func (v VZMultiTouchDeviceConfiguration) DebugDescription() string {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
func (v VZMultiTouchDeviceConfiguration) Description() string {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
func (v VZMultiTouchDeviceConfiguration) Hash() uint64 {
	rv := objc.Send[uint64](v.ID, objc.Sel("hash"))
	return rv
}
func (v VZMultiTouchDeviceConfiguration) Superclass() objectivec.Class {
	rv := objc.Send[objectivec.Class](v.ID, objc.Sel("superclass"))
	return objectivec.Class(rv)
}
