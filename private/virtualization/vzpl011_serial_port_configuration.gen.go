// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZPL011SerialPortConfiguration] class.
var (
	_VZPL011SerialPortConfigurationClass     VZPL011SerialPortConfigurationClass
	_VZPL011SerialPortConfigurationClassOnce sync.Once
)

func getVZPL011SerialPortConfigurationClass() VZPL011SerialPortConfigurationClass {
	_VZPL011SerialPortConfigurationClassOnce.Do(func() {
		_VZPL011SerialPortConfigurationClass = VZPL011SerialPortConfigurationClass{class: objc.GetClass("_VZPL011SerialPortConfiguration")}
	})
	return _VZPL011SerialPortConfigurationClass
}

// GetVZPL011SerialPortConfigurationClass returns the class object for _VZPL011SerialPortConfiguration.
func GetVZPL011SerialPortConfigurationClass() VZPL011SerialPortConfigurationClass {
	return getVZPL011SerialPortConfigurationClass()
}

type VZPL011SerialPortConfigurationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZPL011SerialPortConfigurationClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZPL011SerialPortConfigurationClass) Alloc() VZPL011SerialPortConfiguration {
	rv := objc.Send[VZPL011SerialPortConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [VZPL011SerialPortConfiguration.EncodeWithEncoder]
//
// See: https://developer.apple.com/documentation/Virtualization/_VZPL011SerialPortConfiguration
type VZPL011SerialPortConfiguration struct {
	VZSerialPortConfiguration
}

// VZPL011SerialPortConfigurationFromID constructs a [VZPL011SerialPortConfiguration] from an objc.ID.
func VZPL011SerialPortConfigurationFromID(id objc.ID) VZPL011SerialPortConfiguration {
	return VZPL011SerialPortConfiguration{VZSerialPortConfiguration: VZSerialPortConfigurationFromID(id)}
}

// Ensure VZPL011SerialPortConfiguration implements IVZPL011SerialPortConfiguration.
var _ IVZPL011SerialPortConfiguration = VZPL011SerialPortConfiguration{}

// An interface definition for the [VZPL011SerialPortConfiguration] class.
//
// # Methods
//
//   - [IVZPL011SerialPortConfiguration.EncodeWithEncoder]
//
// See: https://developer.apple.com/documentation/Virtualization/_VZPL011SerialPortConfiguration
type IVZPL011SerialPortConfiguration interface {
	IVZSerialPortConfiguration

	// Topic: Methods

	EncodeWithEncoder(encoder objectivec.IObject) objectivec.IObject
}

// Init initializes the instance.
func (v VZPL011SerialPortConfiguration) Init() VZPL011SerialPortConfiguration {
	rv := objc.Send[VZPL011SerialPortConfiguration](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZPL011SerialPortConfiguration) Autorelease() VZPL011SerialPortConfiguration {
	rv := objc.Send[VZPL011SerialPortConfiguration](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZPL011SerialPortConfiguration creates a new VZPL011SerialPortConfiguration instance.
func NewVZPL011SerialPortConfiguration() VZPL011SerialPortConfiguration {
	class := getVZPL011SerialPortConfigurationClass()
	rv := objc.Send[VZPL011SerialPortConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/_VZPL011SerialPortConfiguration/encodeWithEncoder:
func (v VZPL011SerialPortConfiguration) EncodeWithEncoder(encoder objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("encodeWithEncoder:"), encoder)
	return objectivec.Object{ID: rv}
}
