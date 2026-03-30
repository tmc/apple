// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSPressureConfiguration] class.
var (
	_NSPressureConfigurationClass     NSPressureConfigurationClass
	_NSPressureConfigurationClassOnce sync.Once
)

func getNSPressureConfigurationClass() NSPressureConfigurationClass {
	_NSPressureConfigurationClassOnce.Do(func() {
		_NSPressureConfigurationClass = NSPressureConfigurationClass{class: objc.GetClass("NSPressureConfiguration")}
	})
	return _NSPressureConfigurationClass
}

// GetNSPressureConfigurationClass returns the class object for NSPressureConfiguration.
func GetNSPressureConfigurationClass() NSPressureConfigurationClass {
	return getNSPressureConfigurationClass()
}

type NSPressureConfigurationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSPressureConfigurationClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSPressureConfigurationClass) Alloc() NSPressureConfiguration {
	rv := objc.Send[NSPressureConfiguration](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// An encapsulation of the behavior and progression of a Force Touch trackpad
// as it responds to specific events.
//
// # Overview
//
// Use an [NSPressureConfiguration] object to configure the behavior and
// progression of a Force Touch trackpad when it responds to a mouse drag or
// pressure event sequence. Pressure configurations are assigned to views
// ([NSView]) and gesture recognizers ([NSGestureRecognizer]).
//
// # Creating a Pressure Configuration Object
//
//   - [NSPressureConfiguration.InitWithPressureBehavior]: Initializes a pressure configuration object with a specified pressure behavior.
//   - [NSPressureConfiguration.Set]: Changes the pressure configuration of the trackpad to the initialized pressure configuration.
//
// # Accessing Pressure Configuration Object Properties
//
//   - [NSPressureConfiguration.PressureBehavior]: The pressure behavior of the pressure configuration object.
//
// See: https://developer.apple.com/documentation/AppKit/NSPressureConfiguration
type NSPressureConfiguration struct {
	objectivec.Object
}

// NSPressureConfigurationFromID constructs a [NSPressureConfiguration] from an objc.ID.
//
// An encapsulation of the behavior and progression of a Force Touch trackpad
// as it responds to specific events.
func NSPressureConfigurationFromID(id objc.ID) NSPressureConfiguration {
	return NSPressureConfiguration{objectivec.Object{ID: id}}
}

// NOTE: NSPressureConfiguration adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSPressureConfiguration] class.
//
// # Creating a Pressure Configuration Object
//
//   - [INSPressureConfiguration.InitWithPressureBehavior]: Initializes a pressure configuration object with a specified pressure behavior.
//   - [INSPressureConfiguration.Set]: Changes the pressure configuration of the trackpad to the initialized pressure configuration.
//
// # Accessing Pressure Configuration Object Properties
//
//   - [INSPressureConfiguration.PressureBehavior]: The pressure behavior of the pressure configuration object.
//
// See: https://developer.apple.com/documentation/AppKit/NSPressureConfiguration
type INSPressureConfiguration interface {
	objectivec.IObject

	// Topic: Creating a Pressure Configuration Object

	// Initializes a pressure configuration object with a specified pressure behavior.
	InitWithPressureBehavior(pressureBehavior NSPressureBehavior) NSPressureConfiguration
	// Changes the pressure configuration of the trackpad to the initialized pressure configuration.
	Set()

	// Topic: Accessing Pressure Configuration Object Properties

	// The pressure behavior of the pressure configuration object.
	PressureBehavior() NSPressureBehavior
}

// Init initializes the instance.
func (p NSPressureConfiguration) Init() NSPressureConfiguration {
	rv := objc.Send[NSPressureConfiguration](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p NSPressureConfiguration) Autorelease() NSPressureConfiguration {
	rv := objc.Send[NSPressureConfiguration](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSPressureConfiguration creates a new NSPressureConfiguration instance.
func NewNSPressureConfiguration() NSPressureConfiguration {
	class := getNSPressureConfigurationClass()
	rv := objc.Send[NSPressureConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Initializes a pressure configuration object with a specified pressure
// behavior.
//
// pressureBehavior: An [NSPressureBehavior] value that describes the behavior and progression
// for responding to pressure events.
//
// # Return Value
//
// A new pressure configuration object of type [NSPressureConfiguration] that
// describes how pressure events behave and progress.
//
// # Discussion
//
// The initialized pressure configuration object is used to change the
// behavior and progression of the trackpad when responding to a mouse drag or
// pressure event sequence.
//
// See: https://developer.apple.com/documentation/AppKit/NSPressureConfiguration/init(pressureBehavior:)
func NewPressureConfigurationWithPressureBehavior(pressureBehavior NSPressureBehavior) NSPressureConfiguration {
	instance := getNSPressureConfigurationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithPressureBehavior:"), pressureBehavior)
	return NSPressureConfigurationFromID(rv)
}

// Initializes a pressure configuration object with a specified pressure
// behavior.
//
// pressureBehavior: An [NSPressureBehavior] value that describes the behavior and progression
// for responding to pressure events.
//
// # Return Value
//
// A new pressure configuration object of type [NSPressureConfiguration] that
// describes how pressure events behave and progress.
//
// # Discussion
//
// The initialized pressure configuration object is used to change the
// behavior and progression of the trackpad when responding to a mouse drag or
// pressure event sequence.
//
// See: https://developer.apple.com/documentation/AppKit/NSPressureConfiguration/init(pressureBehavior:)
func (p NSPressureConfiguration) InitWithPressureBehavior(pressureBehavior NSPressureBehavior) NSPressureConfiguration {
	rv := objc.Send[NSPressureConfiguration](p.ID, objc.Sel("initWithPressureBehavior:"), pressureBehavior)
	return rv
}

// Changes the pressure configuration of the trackpad to the initialized
// pressure configuration.
//
// # Discussion
//
// During a mouse drag or pressure event sequence, this method may be called
// to change the pressure configuration of the trackpad to the initialized
// pressure configuration. The trackpad’s pressure configuration is
// automatically reset when the user releases the mouse or ends the gesture.
// If called outside of a mouse drag or pressure event sequence, this method
// has no effect on the trackpad.
//
// Ideally, pressure behavior should be configured by adjusting the
// `pressureConfiguration` property of a view prior to a mouse drag or gesture
// event occurring, such as before the view is displayed. Otherwise, the user
// may complete the mouse drag or gesture before the configuration has time to
// take effect.
//
// See: https://developer.apple.com/documentation/AppKit/NSPressureConfiguration/set()
func (p NSPressureConfiguration) Set() {
	objc.Send[objc.ID](p.ID, objc.Sel("set"))
}

// The pressure behavior of the pressure configuration object.
//
// # Discussion
//
// The value of this property is set when the pressure configuration object is
// created and cannot be changed later. It contains the pressure behavior type
// of the pressure configuration object.
//
// See: https://developer.apple.com/documentation/AppKit/NSPressureConfiguration/pressureBehavior
func (p NSPressureConfiguration) PressureBehavior() NSPressureBehavior {
	rv := objc.Send[NSPressureBehavior](p.ID, objc.Sel("pressureBehavior"))
	return NSPressureBehavior(rv)
}
