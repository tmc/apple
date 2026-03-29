// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"context"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVCaptureDeskViewApplication] class.
var (
	_AVCaptureDeskViewApplicationClass     AVCaptureDeskViewApplicationClass
	_AVCaptureDeskViewApplicationClassOnce sync.Once
)

func getAVCaptureDeskViewApplicationClass() AVCaptureDeskViewApplicationClass {
	_AVCaptureDeskViewApplicationClassOnce.Do(func() {
		_AVCaptureDeskViewApplicationClass = AVCaptureDeskViewApplicationClass{class: objc.GetClass("AVCaptureDeskViewApplication")}
	})
	return _AVCaptureDeskViewApplicationClass
}

// GetAVCaptureDeskViewApplicationClass returns the class object for AVCaptureDeskViewApplication.
func GetAVCaptureDeskViewApplicationClass() AVCaptureDeskViewApplicationClass {
	return getAVCaptureDeskViewApplicationClass()
}

type AVCaptureDeskViewApplicationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVCaptureDeskViewApplicationClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVCaptureDeskViewApplicationClass) Alloc() AVCaptureDeskViewApplication {
	rv := objc.Send[AVCaptureDeskViewApplication](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that programmatically presents Desk View.
//
// # Overview
// 
// Use this class to programmatically launch Desk View from your app. You can
// optionally customize the presentation and specifiy an action to take
// afterward.
// 
// The following example shows how to configure and present Desk View with a
// completion handler:
//
// # Presenting the Desk View app
//
//   - [AVCaptureDeskViewApplication.PresentWithCompletionHandler]: Launches Desk View with no additional configuration and then performs a completion handler if you specify it.
//   - [AVCaptureDeskViewApplication.PresentWithLaunchConfigurationCompletionHandler]: Launches Desk View with the configuration and completion handler that you specify.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDeskViewApplication
type AVCaptureDeskViewApplication struct {
	objectivec.Object
}

// AVCaptureDeskViewApplicationFromID constructs a [AVCaptureDeskViewApplication] from an objc.ID.
//
// An object that programmatically presents Desk View.
func AVCaptureDeskViewApplicationFromID(id objc.ID) AVCaptureDeskViewApplication {
	return AVCaptureDeskViewApplication{objectivec.Object{ID: id}}
}
// NOTE: AVCaptureDeskViewApplication adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVCaptureDeskViewApplication] class.
//
// # Presenting the Desk View app
//
//   - [IAVCaptureDeskViewApplication.PresentWithCompletionHandler]: Launches Desk View with no additional configuration and then performs a completion handler if you specify it.
//   - [IAVCaptureDeskViewApplication.PresentWithLaunchConfigurationCompletionHandler]: Launches Desk View with the configuration and completion handler that you specify.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDeskViewApplication
type IAVCaptureDeskViewApplication interface {
	objectivec.IObject

	// Topic: Presenting the Desk View app

	// Launches Desk View with no additional configuration and then performs a completion handler if you specify it.
	PresentWithCompletionHandler(completionHandler ErrorHandler)
	// Launches Desk View with the configuration and completion handler that you specify.
	PresentWithLaunchConfigurationCompletionHandler(launchConfiguration IAVCaptureDeskViewApplicationLaunchConfiguration, completionHandler ErrorHandler)
}

// Init initializes the instance.
func (c AVCaptureDeskViewApplication) Init() AVCaptureDeskViewApplication {
	rv := objc.Send[AVCaptureDeskViewApplication](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c AVCaptureDeskViewApplication) Autorelease() AVCaptureDeskViewApplication {
	rv := objc.Send[AVCaptureDeskViewApplication](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVCaptureDeskViewApplication creates a new AVCaptureDeskViewApplication instance.
func NewAVCaptureDeskViewApplication() AVCaptureDeskViewApplication {
	class := getAVCaptureDeskViewApplicationClass()
	rv := objc.Send[AVCaptureDeskViewApplication](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Launches Desk View with no additional configuration and then performs a
// completion handler if you specify it.
//
// completionHandler: The code to perform after the system displays Desk View.
//
// # Discussion
// 
// If the Desk View app is already running, this method brings it to the
// front. If Desk View is in the Dock, this method opens it and brings it to
// the front.
// 
// Desk View launches in setup mode. This mode shows the full field of view of
// an ultrawide camera with a superimposed trapezoid that indicates the
// cropped desk region to display. The system displays this region after the
// user completes setup and starts Desk View.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDeskViewApplication/present(completionHandler:)
func (c AVCaptureDeskViewApplication) PresentWithCompletionHandler(completionHandler ErrorHandler) {
_block0, _ := NewErrorBlock(completionHandler)
	objc.Send[objc.ID](c.ID, objc.Sel("presentWithCompletionHandler:"), _block0)
}
// Launches Desk View with the configuration and completion handler that you
// specify.
//
// launchConfiguration: A configuration that specifies how to present Desk View.
//
// completionHandler: The code to perform after the system displays Desk View or the user
// transitions to Desk View after setup, depending on the configuration.
//
// # Discussion
// 
// If Desk View is already running, this method brings it to the front. If
// Desk View is in the Dock, this method opens it and brings it to the front.
// 
// Desk View launches in setup mode. This mode shows the full field of view of
// an ultrawide camera with a superimposed trapezoid that indicates the
// cropped desk region to display. The system displays this region after the
// user completes setup and starts Desk View.
// 
// Create an instance of [AVCaptureDeskViewApplicationLaunchConfiguration] and
// set it for `launchConfiguration` to specify the frame for Desk View and
// when to perform the `completionHandler`.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDeskViewApplication/present(launchConfiguration:completionHandler:)
func (c AVCaptureDeskViewApplication) PresentWithLaunchConfigurationCompletionHandler(launchConfiguration IAVCaptureDeskViewApplicationLaunchConfiguration, completionHandler ErrorHandler) {
_block1, _ := NewErrorBlock(completionHandler)
	objc.Send[objc.ID](c.ID, objc.Sel("presentWithLaunchConfiguration:completionHandler:"), launchConfiguration, _block1)
}

// Present is a synchronous wrapper around [AVCaptureDeskViewApplication.PresentWithCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (c AVCaptureDeskViewApplication) Present(ctx context.Context) error {
	done := make(chan error, 1)
	c.PresentWithCompletionHandler(func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

// PresentWithLaunchConfiguration is a synchronous wrapper around [AVCaptureDeskViewApplication.PresentWithLaunchConfigurationCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (c AVCaptureDeskViewApplication) PresentWithLaunchConfiguration(ctx context.Context, launchConfiguration IAVCaptureDeskViewApplicationLaunchConfiguration) error {
	done := make(chan error, 1)
	c.PresentWithLaunchConfigurationCompletionHandler(launchConfiguration, func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

