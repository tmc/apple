// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZKeyboardConfiguration] class.
var (
	_VZKeyboardConfigurationClass     VZKeyboardConfigurationClass
	_VZKeyboardConfigurationClassOnce sync.Once
)

func getVZKeyboardConfigurationClass() VZKeyboardConfigurationClass {
	_VZKeyboardConfigurationClassOnce.Do(func() {
		_VZKeyboardConfigurationClass = VZKeyboardConfigurationClass{class: objc.GetClass("VZKeyboardConfiguration")}
	})
	return _VZKeyboardConfigurationClass
}

// GetVZKeyboardConfigurationClass returns the class object for VZKeyboardConfiguration.
func GetVZKeyboardConfigurationClass() VZKeyboardConfigurationClass {
	return getVZKeyboardConfigurationClass()
}

type VZKeyboardConfigurationClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZKeyboardConfigurationClass) Alloc() VZKeyboardConfiguration {
	rv := objc.Send[VZKeyboardConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// The base class for a configuring a keyboard.
//
// # Overview
// 
// [VZKeyboardConfiguration] defines the abstract interface that defines a
// virtual keyboard that you connect to a guest operating system. Don’t
// instantiate [VZKeyboardConfiguration] directly, use one of its subclasses
// such as [VZUSBKeyboardConfiguration] instead.
//
// See: https://developer.apple.com/documentation/Virtualization/VZKeyboardConfiguration
type VZKeyboardConfiguration struct {
	objectivec.Object
}

// VZKeyboardConfigurationFromID constructs a [VZKeyboardConfiguration] from an objc.ID.
//
// The base class for a configuring a keyboard.
func VZKeyboardConfigurationFromID(id objc.ID) VZKeyboardConfiguration {
	return VZKeyboardConfiguration{objectivec.Object{ID: id}}
}
// NOTE: VZKeyboardConfiguration adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [VZKeyboardConfiguration] class.
//
// See: https://developer.apple.com/documentation/Virtualization/VZKeyboardConfiguration
type IVZKeyboardConfiguration interface {
	objectivec.IObject
}

// Init initializes the instance.
func (k VZKeyboardConfiguration) Init() VZKeyboardConfiguration {
	rv := objc.Send[VZKeyboardConfiguration](k.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (k VZKeyboardConfiguration) Autorelease() VZKeyboardConfiguration {
	rv := objc.Send[VZKeyboardConfiguration](k.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZKeyboardConfiguration creates a new VZKeyboardConfiguration instance.
func NewVZKeyboardConfiguration() VZKeyboardConfiguration {
	class := getVZKeyboardConfigurationClass()
	rv := objc.Send[VZKeyboardConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

