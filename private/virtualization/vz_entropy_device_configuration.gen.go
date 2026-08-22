// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZEntropyDeviceConfiguration] class.
var (
	_VZEntropyDeviceConfigurationClass     VZEntropyDeviceConfigurationClass
	_VZEntropyDeviceConfigurationClassOnce sync.Once
)

func getVZEntropyDeviceConfigurationClass() VZEntropyDeviceConfigurationClass {
	_VZEntropyDeviceConfigurationClassOnce.Do(func() {
		_VZEntropyDeviceConfigurationClass = VZEntropyDeviceConfigurationClass{class: objc.GetClass("VZEntropyDeviceConfiguration")}
	})
	return _VZEntropyDeviceConfigurationClass
}

// GetVZEntropyDeviceConfigurationClass returns the class object for VZEntropyDeviceConfiguration.
func GetVZEntropyDeviceConfigurationClass() VZEntropyDeviceConfigurationClass {
	return getVZEntropyDeviceConfigurationClass()
}

type VZEntropyDeviceConfigurationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZEntropyDeviceConfigurationClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZEntropyDeviceConfigurationClass) Alloc() VZEntropyDeviceConfiguration {
	rv := objc.SendIfResponds[VZEntropyDeviceConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [VZEntropyDeviceConfiguration._entropyDevice]
//   - [VZEntropyDeviceConfiguration._init]
//   - [VZEntropyDeviceConfiguration.DebugDescription]
//   - [VZEntropyDeviceConfiguration.Description]
//   - [VZEntropyDeviceConfiguration.Hash]
//   - [VZEntropyDeviceConfiguration.Superclass]
type VZEntropyDeviceConfiguration struct {
	objectivec.Object
}

// VZEntropyDeviceConfigurationFromID constructs a [VZEntropyDeviceConfiguration] from an objc.ID.
func VZEntropyDeviceConfigurationFromID(id objc.ID) VZEntropyDeviceConfiguration {
	return VZEntropyDeviceConfiguration{objectivec.Object{ID: id}}
}

// Ensure VZEntropyDeviceConfiguration implements IVZEntropyDeviceConfiguration.
var _ IVZEntropyDeviceConfiguration = VZEntropyDeviceConfiguration{}

// An interface definition for the [VZEntropyDeviceConfiguration] class.
//
// # Methods
//
//   - [IVZEntropyDeviceConfiguration._entropyDevice]
//   - [IVZEntropyDeviceConfiguration._init]
//   - [IVZEntropyDeviceConfiguration.DebugDescription]
//   - [IVZEntropyDeviceConfiguration.Description]
//   - [IVZEntropyDeviceConfiguration.Hash]
//   - [IVZEntropyDeviceConfiguration.Superclass]
type IVZEntropyDeviceConfiguration interface {
	objectivec.IObject

	// Topic: Methods

	_entropyDevice() int
	_init() objectivec.IObject
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objectivec.Class
}

// Init initializes the instance.
func (v VZEntropyDeviceConfiguration) Init() VZEntropyDeviceConfiguration {
	rv := objc.SendIfResponds[VZEntropyDeviceConfiguration](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZEntropyDeviceConfiguration) Autorelease() VZEntropyDeviceConfiguration {
	rv := objc.SendIfResponds[VZEntropyDeviceConfiguration](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZEntropyDeviceConfiguration creates a new VZEntropyDeviceConfiguration instance.
func NewVZEntropyDeviceConfiguration() VZEntropyDeviceConfiguration {
	class := getVZEntropyDeviceConfigurationClass()
	rv := objc.SendIfResponds[VZEntropyDeviceConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (v VZEntropyDeviceConfiguration) _init() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](v.ID, objc.Sel("_init"))
	return objectivec.Object{ID: rv}
}

func (v VZEntropyDeviceConfiguration) _entropyDevice() int {
	rv := objc.SendIfResponds[int](v.ID, objc.Sel("_entropyDevice"))
	return rv
}

// CanEntropyDevice reports whether the receiver responds to the private selector _entropyDevice.
func (v VZEntropyDeviceConfiguration) CanEntropyDevice() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_entropyDevice"))
}

// EntropyDevice is an exported wrapper for the private property _entropyDevice.
func (v VZEntropyDeviceConfiguration) EntropyDevice() (int, error) {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_entropyDevice")) {
		return 0, &objc.UnrecognizedSelectorError{Selector: "_entropyDevice"}
	}
	return v._entropyDevice(), nil
}
func (v VZEntropyDeviceConfiguration) DebugDescription() string {
	rv := objc.SendIfResponds[objc.ID](v.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
func (v VZEntropyDeviceConfiguration) Description() string {
	rv := objc.SendIfResponds[objc.ID](v.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
func (v VZEntropyDeviceConfiguration) Hash() uint64 {
	rv := objc.SendIfResponds[uint64](v.ID, objc.Sel("hash"))
	return rv
}
func (v VZEntropyDeviceConfiguration) Superclass() objectivec.Class {
	rv := objc.SendIfResponds[objectivec.Class](v.ID, objc.Sel("superclass"))
	return objectivec.Class(rv)
}
