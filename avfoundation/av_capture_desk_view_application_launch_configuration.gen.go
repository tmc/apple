// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVCaptureDeskViewApplicationLaunchConfiguration] class.
var (
	_AVCaptureDeskViewApplicationLaunchConfigurationClass     AVCaptureDeskViewApplicationLaunchConfigurationClass
	_AVCaptureDeskViewApplicationLaunchConfigurationClassOnce sync.Once
)

func getAVCaptureDeskViewApplicationLaunchConfigurationClass() AVCaptureDeskViewApplicationLaunchConfigurationClass {
	_AVCaptureDeskViewApplicationLaunchConfigurationClassOnce.Do(func() {
		_AVCaptureDeskViewApplicationLaunchConfigurationClass = AVCaptureDeskViewApplicationLaunchConfigurationClass{class: objc.GetClass("AVCaptureDeskViewApplicationLaunchConfiguration")}
	})
	return _AVCaptureDeskViewApplicationLaunchConfigurationClass
}

// GetAVCaptureDeskViewApplicationLaunchConfigurationClass returns the class object for AVCaptureDeskViewApplicationLaunchConfiguration.
func GetAVCaptureDeskViewApplicationLaunchConfigurationClass() AVCaptureDeskViewApplicationLaunchConfigurationClass {
	return getAVCaptureDeskViewApplicationLaunchConfigurationClass()
}

type AVCaptureDeskViewApplicationLaunchConfigurationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVCaptureDeskViewApplicationLaunchConfigurationClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVCaptureDeskViewApplicationLaunchConfigurationClass) Alloc() AVCaptureDeskViewApplicationLaunchConfiguration {
	rv := objc.Send[AVCaptureDeskViewApplicationLaunchConfiguration](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that configures how to present Desk View.
//
// # Overview
// 
// Use this object to specify the frame for Desk View when it launches, and
// when to execute the completion handler. You can specify whether to perform
// the completion handler as soon as Desk View is visible to the user, or only
// after they start Desk View.
//
// # Customizing the presentation
//
//   - [AVCaptureDeskViewApplicationLaunchConfiguration.MainWindowFrame]: The frame for Desk View after it launches.
//   - [AVCaptureDeskViewApplicationLaunchConfiguration.SetMainWindowFrame]
//   - [AVCaptureDeskViewApplicationLaunchConfiguration.RequiresSetUpModeCompletion]: A Boolean value that specifies whether the system requires the user to complete setup mode before it executes the completion handler.
//   - [AVCaptureDeskViewApplicationLaunchConfiguration.SetRequiresSetUpModeCompletion]
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDeskViewApplication/LaunchConfiguration
type AVCaptureDeskViewApplicationLaunchConfiguration struct {
	objectivec.Object
}

// AVCaptureDeskViewApplicationLaunchConfigurationFromID constructs a [AVCaptureDeskViewApplicationLaunchConfiguration] from an objc.ID.
//
// An object that configures how to present Desk View.
func AVCaptureDeskViewApplicationLaunchConfigurationFromID(id objc.ID) AVCaptureDeskViewApplicationLaunchConfiguration {
	return AVCaptureDeskViewApplicationLaunchConfiguration{objectivec.Object{ID: id}}
}
// NOTE: AVCaptureDeskViewApplicationLaunchConfiguration adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVCaptureDeskViewApplicationLaunchConfiguration] class.
//
// # Customizing the presentation
//
//   - [IAVCaptureDeskViewApplicationLaunchConfiguration.MainWindowFrame]: The frame for Desk View after it launches.
//   - [IAVCaptureDeskViewApplicationLaunchConfiguration.SetMainWindowFrame]
//   - [IAVCaptureDeskViewApplicationLaunchConfiguration.RequiresSetUpModeCompletion]: A Boolean value that specifies whether the system requires the user to complete setup mode before it executes the completion handler.
//   - [IAVCaptureDeskViewApplicationLaunchConfiguration.SetRequiresSetUpModeCompletion]
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDeskViewApplication/LaunchConfiguration
type IAVCaptureDeskViewApplicationLaunchConfiguration interface {
	objectivec.IObject

	// Topic: Customizing the presentation

	// The frame for Desk View after it launches.
	MainWindowFrame() corefoundation.CGRect
	SetMainWindowFrame(value corefoundation.CGRect)
	// A Boolean value that specifies whether the system requires the user to complete setup mode before it executes the completion handler.
	RequiresSetUpModeCompletion() bool
	SetRequiresSetUpModeCompletion(value bool)
}

// Init initializes the instance.
func (c AVCaptureDeskViewApplicationLaunchConfiguration) Init() AVCaptureDeskViewApplicationLaunchConfiguration {
	rv := objc.Send[AVCaptureDeskViewApplicationLaunchConfiguration](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c AVCaptureDeskViewApplicationLaunchConfiguration) Autorelease() AVCaptureDeskViewApplicationLaunchConfiguration {
	rv := objc.Send[AVCaptureDeskViewApplicationLaunchConfiguration](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVCaptureDeskViewApplicationLaunchConfiguration creates a new AVCaptureDeskViewApplicationLaunchConfiguration instance.
func NewAVCaptureDeskViewApplicationLaunchConfiguration() AVCaptureDeskViewApplicationLaunchConfiguration {
	class := getAVCaptureDeskViewApplicationLaunchConfigurationClass()
	rv := objc.Send[AVCaptureDeskViewApplicationLaunchConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The frame for Desk View after it launches.
//
// # Discussion
// 
// The default value is [zero], which tells the system to use the previously
// set frame. The system uses global screen coordinates to display the frame.
// When Desk View launches from a native macOS app, the window origin is
// bottom-left. When it launches from a [Mac Catalyst] app, the window origin
// is top-left.
//
// [Mac Catalyst]: https://developer.apple.com/documentation/UIKit/mac-catalyst
// [zero]: https://developer.apple.com/documentation/CoreFoundation/CGRect/zero
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDeskViewApplication/LaunchConfiguration/mainWindowFrame
func (c AVCaptureDeskViewApplicationLaunchConfiguration) MainWindowFrame() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](c.ID, objc.Sel("mainWindowFrame"))
	return corefoundation.CGRect(rv)
}
func (c AVCaptureDeskViewApplicationLaunchConfiguration) SetMainWindowFrame(value corefoundation.CGRect) {
	objc.Send[struct{}](c.ID, objc.Sel("setMainWindowFrame:"), value)
}
// A Boolean value that specifies whether the system requires the user to
// complete setup mode before it executes the completion handler.
//
// # Discussion
// 
// The default value is [false], which tells the system to execute the
// completion handler as soon as it displays Desk View. If [true], the system
// executes the completion handler after the user completes setup and starts
// Desk View.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDeskViewApplication/LaunchConfiguration/requiresSetUpModeCompletion
func (c AVCaptureDeskViewApplicationLaunchConfiguration) RequiresSetUpModeCompletion() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("requiresSetUpModeCompletion"))
	return rv
}
func (c AVCaptureDeskViewApplicationLaunchConfiguration) SetRequiresSetUpModeCompletion(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setRequiresSetUpModeCompletion:"), value)
}

