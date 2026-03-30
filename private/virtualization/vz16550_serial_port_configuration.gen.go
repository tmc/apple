// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZ16550SerialPortConfiguration] class.
var (
	_VZ16550SerialPortConfigurationClass     VZ16550SerialPortConfigurationClass
	_VZ16550SerialPortConfigurationClassOnce sync.Once
)

func getVZ16550SerialPortConfigurationClass() VZ16550SerialPortConfigurationClass {
	_VZ16550SerialPortConfigurationClassOnce.Do(func() {
		_VZ16550SerialPortConfigurationClass = VZ16550SerialPortConfigurationClass{class: objc.GetClass("_VZ16550SerialPortConfiguration")}
	})
	return _VZ16550SerialPortConfigurationClass
}

// GetVZ16550SerialPortConfigurationClass returns the class object for _VZ16550SerialPortConfiguration.
func GetVZ16550SerialPortConfigurationClass() VZ16550SerialPortConfigurationClass {
	return getVZ16550SerialPortConfigurationClass()
}

type VZ16550SerialPortConfigurationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZ16550SerialPortConfigurationClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZ16550SerialPortConfigurationClass) Alloc() VZ16550SerialPortConfiguration {
	rv := objc.Send[VZ16550SerialPortConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [VZ16550SerialPortConfiguration.EncodeWithEncoder]
//
// See: https://developer.apple.com/documentation/Virtualization/_VZ16550SerialPortConfiguration
type VZ16550SerialPortConfiguration struct {
	VZSerialPortConfiguration
}

// VZ16550SerialPortConfigurationFromID constructs a [VZ16550SerialPortConfiguration] from an objc.ID.
func VZ16550SerialPortConfigurationFromID(id objc.ID) VZ16550SerialPortConfiguration {
	return VZ16550SerialPortConfiguration{VZSerialPortConfiguration: VZSerialPortConfigurationFromID(id)}
}

// Ensure VZ16550SerialPortConfiguration implements IVZ16550SerialPortConfiguration.
var _ IVZ16550SerialPortConfiguration = VZ16550SerialPortConfiguration{}

// An interface definition for the [VZ16550SerialPortConfiguration] class.
//
// # Methods
//
//   - [IVZ16550SerialPortConfiguration.EncodeWithEncoder]
//
// See: https://developer.apple.com/documentation/Virtualization/_VZ16550SerialPortConfiguration
type IVZ16550SerialPortConfiguration interface {
	IVZSerialPortConfiguration

	// Topic: Methods

	EncodeWithEncoder(encoder objectivec.IObject) objectivec.IObject
}

// Init initializes the instance.
func (v VZ16550SerialPortConfiguration) Init() VZ16550SerialPortConfiguration {
	rv := objc.Send[VZ16550SerialPortConfiguration](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZ16550SerialPortConfiguration) Autorelease() VZ16550SerialPortConfiguration {
	rv := objc.Send[VZ16550SerialPortConfiguration](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZ16550SerialPortConfiguration creates a new VZ16550SerialPortConfiguration instance.
func NewVZ16550SerialPortConfiguration() VZ16550SerialPortConfiguration {
	class := getVZ16550SerialPortConfigurationClass()
	rv := objc.Send[VZ16550SerialPortConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/_VZ16550SerialPortConfiguration/encodeWithEncoder:
func (v VZ16550SerialPortConfiguration) EncodeWithEncoder(encoder objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("encodeWithEncoder:"), encoder)
	return objectivec.Object{ID: rv}
}
