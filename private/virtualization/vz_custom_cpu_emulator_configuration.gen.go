// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// The class instance for the [VZCustomCPUEmulatorConfiguration] class.
var (
	_VZCustomCPUEmulatorConfigurationClass     VZCustomCPUEmulatorConfigurationClass
	_VZCustomCPUEmulatorConfigurationClassOnce sync.Once
)

func getVZCustomCPUEmulatorConfigurationClass() VZCustomCPUEmulatorConfigurationClass {
	_VZCustomCPUEmulatorConfigurationClassOnce.Do(func() {
		_VZCustomCPUEmulatorConfigurationClass = VZCustomCPUEmulatorConfigurationClass{class: objc.GetClass("_VZCustomCPUEmulatorConfiguration")}
	})
	return _VZCustomCPUEmulatorConfigurationClass
}

// GetVZCustomCPUEmulatorConfigurationClass returns the class object for _VZCustomCPUEmulatorConfiguration.
func GetVZCustomCPUEmulatorConfigurationClass() VZCustomCPUEmulatorConfigurationClass {
	return getVZCustomCPUEmulatorConfigurationClass()
}

type VZCustomCPUEmulatorConfigurationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZCustomCPUEmulatorConfigurationClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZCustomCPUEmulatorConfigurationClass) Alloc() VZCustomCPUEmulatorConfiguration {
	rv := objc.Send[VZCustomCPUEmulatorConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [VZCustomCPUEmulatorConfiguration.EmulatorURL]
//   - [VZCustomCPUEmulatorConfiguration.SetEmulatorURL]
//   - [VZCustomCPUEmulatorConfiguration.MemorySize]
//   - [VZCustomCPUEmulatorConfiguration.SetMemorySize]
//   - [VZCustomCPUEmulatorConfiguration.Options]
//   - [VZCustomCPUEmulatorConfiguration.SetOptions]
//
// See: https://developer.apple.com/documentation/Virtualization/_VZCustomCPUEmulatorConfiguration
type VZCustomCPUEmulatorConfiguration struct {
	VZCPUEmulatorConfiguration
}

// VZCustomCPUEmulatorConfigurationFromID constructs a [VZCustomCPUEmulatorConfiguration] from an objc.ID.
func VZCustomCPUEmulatorConfigurationFromID(id objc.ID) VZCustomCPUEmulatorConfiguration {
	return VZCustomCPUEmulatorConfiguration{VZCPUEmulatorConfiguration: VZCPUEmulatorConfigurationFromID(id)}
}

// Ensure VZCustomCPUEmulatorConfiguration implements IVZCustomCPUEmulatorConfiguration.
var _ IVZCustomCPUEmulatorConfiguration = VZCustomCPUEmulatorConfiguration{}

// An interface definition for the [VZCustomCPUEmulatorConfiguration] class.
//
// # Methods
//
//   - [IVZCustomCPUEmulatorConfiguration.EmulatorURL]
//   - [IVZCustomCPUEmulatorConfiguration.SetEmulatorURL]
//   - [IVZCustomCPUEmulatorConfiguration.MemorySize]
//   - [IVZCustomCPUEmulatorConfiguration.SetMemorySize]
//   - [IVZCustomCPUEmulatorConfiguration.Options]
//   - [IVZCustomCPUEmulatorConfiguration.SetOptions]
//
// See: https://developer.apple.com/documentation/Virtualization/_VZCustomCPUEmulatorConfiguration
type IVZCustomCPUEmulatorConfiguration interface {
	IVZCPUEmulatorConfiguration

	// Topic: Methods

	EmulatorURL() foundation.INSURL
	SetEmulatorURL(value foundation.INSURL)
	MemorySize() foundation.NSNumber
	SetMemorySize(value foundation.NSNumber)
	Options() string
	SetOptions(value string)
}

// Init initializes the instance.
func (v VZCustomCPUEmulatorConfiguration) Init() VZCustomCPUEmulatorConfiguration {
	rv := objc.Send[VZCustomCPUEmulatorConfiguration](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZCustomCPUEmulatorConfiguration) Autorelease() VZCustomCPUEmulatorConfiguration {
	rv := objc.Send[VZCustomCPUEmulatorConfiguration](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZCustomCPUEmulatorConfiguration creates a new VZCustomCPUEmulatorConfiguration instance.
func NewVZCustomCPUEmulatorConfiguration() VZCustomCPUEmulatorConfiguration {
	class := getVZCustomCPUEmulatorConfigurationClass()
	rv := objc.Send[VZCustomCPUEmulatorConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/_VZCustomCPUEmulatorConfiguration/emulatorURL
func (v VZCustomCPUEmulatorConfiguration) EmulatorURL() foundation.INSURL {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("emulatorURL"))
	return foundation.NSURLFromID(objc.ID(rv))
}
func (v VZCustomCPUEmulatorConfiguration) SetEmulatorURL(value foundation.INSURL) {
	objc.Send[struct{}](v.ID, objc.Sel("setEmulatorURL:"), value)
}

// See: https://developer.apple.com/documentation/Virtualization/_VZCustomCPUEmulatorConfiguration/memorySize
func (v VZCustomCPUEmulatorConfiguration) MemorySize() foundation.NSNumber {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("memorySize"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
func (v VZCustomCPUEmulatorConfiguration) SetMemorySize(value foundation.NSNumber) {
	objc.Send[struct{}](v.ID, objc.Sel("setMemorySize:"), value)
}

// See: https://developer.apple.com/documentation/Virtualization/_VZCustomCPUEmulatorConfiguration/options
func (v VZCustomCPUEmulatorConfiguration) Options() string {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("options"))
	return foundation.NSStringFromID(rv).String()
}
func (v VZCustomCPUEmulatorConfiguration) SetOptions(value string) {
	objc.Send[struct{}](v.ID, objc.Sel("setOptions:"), objc.String(value))
}
