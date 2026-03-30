// Code generated from Apple documentation for ServiceManagement. DO NOT EDIT.

package servicemanagement

import (
	"context"
	"errors"
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SMAppService] class.
var (
	_SMAppServiceClass     SMAppServiceClass
	_SMAppServiceClassOnce sync.Once
)

func getSMAppServiceClass() SMAppServiceClass {
	_SMAppServiceClassOnce.Do(func() {
		_SMAppServiceClass = SMAppServiceClass{class: objc.GetClass("SMAppService")}
	})
	return _SMAppServiceClass
}

// GetSMAppServiceClass returns the class object for SMAppService.
func GetSMAppServiceClass() SMAppServiceClass {
	return getSMAppServiceClass()
}

type SMAppServiceClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SMAppServiceClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SMAppServiceClass) Alloc() SMAppService {
	rv := objc.Send[SMAppService](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// An object the framework uses to control helper executables that live inside
// an app’s main bundle.
//
// # Overview
//
// In macOS 13 and later, use [SMAppService] to register and control
// [LoginItems], [LaunchAgents], and [LaunchDaemons] as helper executables for
// your app. When converting code from earlier versions of macOS, use an
// [SMAppService] object and select one of the following methods depending on
// the type of service your helper executable provides:
//
// - For [SMAppServices] initialized as [LoginItems], the
// [SMAppService.RegisterAndReturnError] and [SMAppService.UnregisterAndReturnError] APIs provide a
// replacement for [SMLoginItemSetEnabled(_:_:)]. - For [SMAppServices]
// initialized as [LaunchAgents], the [SMAppService.RegisterAndReturnError] and
// [SMAppService.UnregisterAndReturnError] methods provide a replacement for installing
// property lists in `~/Library/LaunchAgents` or `/Library/LaunchAgents`. -
// For [SMAppServices] initialized as [LaunchDaemons], the
// [SMAppService.RegisterAndReturnError] and [SMAppService.UnregisterAndReturnError] methods provide a
// replacement for installing property lists in `/Library/LaunchDaemons`.
//
// # Registering services
//
//   - [SMAppService.RegisterAndReturnError]: Registers the service so it can begin launching subject to user approval.
//   - [SMAppService.UnregisterAndReturnError]: Unregisters the service so the system no longer launches it.
//   - [SMAppService.UnregisterWithCompletionHandler]: Unregisters the service so the system no longer launches it and calls a completion handler you provide with the resulting error value.
//
// # Getting the state of the service
//
//   - [SMAppService.Status]: A property that describes registration or authorization state of the service.
//
// See: https://developer.apple.com/documentation/ServiceManagement/SMAppService
//
// [SMLoginItemSetEnabled(_:_:)]: https://developer.apple.com/documentation/ServiceManagement/SMLoginItemSetEnabled(_:_:)
type SMAppService struct {
	objectivec.Object
}

// SMAppServiceFromID constructs a [SMAppService] from an objc.ID.
//
// An object the framework uses to control helper executables that live inside
// an app’s main bundle.
func SMAppServiceFromID(id objc.ID) SMAppService {
	return SMAppService{objectivec.Object{ID: id}}
}

// NOTE: SMAppService adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [SMAppService] class.
//
// # Registering services
//
//   - [ISMAppService.RegisterAndReturnError]: Registers the service so it can begin launching subject to user approval.
//   - [ISMAppService.UnregisterAndReturnError]: Unregisters the service so the system no longer launches it.
//   - [ISMAppService.UnregisterWithCompletionHandler]: Unregisters the service so the system no longer launches it and calls a completion handler you provide with the resulting error value.
//
// # Getting the state of the service
//
//   - [ISMAppService.Status]: A property that describes registration or authorization state of the service.
//
// See: https://developer.apple.com/documentation/ServiceManagement/SMAppService
type ISMAppService interface {
	objectivec.IObject

	// Topic: Registering services

	// Registers the service so it can begin launching subject to user approval.
	RegisterAndReturnError() (bool, error)
	// Unregisters the service so the system no longer launches it.
	UnregisterAndReturnError() (bool, error)
	// Unregisters the service so the system no longer launches it and calls a completion handler you provide with the resulting error value.
	UnregisterWithCompletionHandler(handler ErrorHandler)

	// Topic: Getting the state of the service

	// A property that describes registration or authorization state of the service.
	Status() SMAppServiceStatus
}

// Init initializes the instance.
func (a SMAppService) Init() SMAppService {
	rv := objc.Send[SMAppService](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a SMAppService) Autorelease() SMAppService {
	rv := objc.Send[SMAppService](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewSMAppService creates a new SMAppService instance.
func NewSMAppService() SMAppService {
	class := getSMAppServiceClass()
	rv := objc.Send[SMAppService](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Registers the service so it can begin launching subject to user approval.
//
// # Discussion
//
// The registration process applies to the following rules, depending upon the
// type of service:
//
// - If the service corresponds to a LoginItem bundle, the helper starts
// immediately and on subsequent logins. If the helper crashes or exits with a
// non-zero status, the system relaunches it. - If the service corresponds to
// the main application, the application launches on subsequent logins. - If
// the service corresponds to a LaunchAgent, the LaunchAgent is immediately
// bootstrapped and may begin running. In addition LaunchAgents registered
// with this method bootstrap on each subsequent login. - If an app needs to
// register a LaunchAgent for multiple users, you must call the API once per
// user while that user is running the app. - If the service corresponds to a
// LaunchDaemon, the system won’t bootstrap the LaunchDaemon until an admin
// approves the LaunchDaemon in System Preferences. The system bootstraps
// LaunchDaemons registered with this method and approved by an admin on each
// subsequent boot.
//
// If the service is already registered, this method returns
// [KSMErrorAlreadyRegistered].
//
// If the service isn’t approved by the user, this method returns
// [KSMErrorLaunchDeniedByUser].
//
// See: https://developer.apple.com/documentation/ServiceManagement/SMAppService/register()
func (a SMAppService) RegisterAndReturnError() (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](a.ID, objc.Sel("registerAndReturnError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("registerAndReturnError: returned NO with nil NSError")
	}
	return rv, nil

}

// Unregisters the service so the system no longer launches it.
//
// # Discussion
//
// This is the opposite operation of [RegisterAndReturnError].
//
// If the service corresponds to a LoginItem, LaunchAgent, or LaunchDaemon and
// the service is currently running it, the system terminates it. If the
// service corresponds to the main application, it continues running, but
// becomes unregistered to prevent future launches at login.
//
// If the service is already unregistered, this method returns
// [KSMErrorJobNotFound].
//
// See: https://developer.apple.com/documentation/ServiceManagement/SMAppService/unregister()
func (a SMAppService) UnregisterAndReturnError() (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](a.ID, objc.Sel("unregisterAndReturnError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("unregisterAndReturnError: returned NO with nil NSError")
	}
	return rv, nil

}

// Unregisters the service so the system no longer launches it and calls a
// completion handler you provide with the resulting error value.
//
// handler: A completion handler to call with the result of the unregistration
// operation. Upon an unsuccessful return, the handler contains a new
// [NSError] object describing the error. Upon successful return, this
// argument is [NULL].
//
// # Discussion
//
// See: https://developer.apple.com/documentation/ServiceManagement/SMAppService/unregister(completionHandler:)
//
// [NSError]: https://developer.apple.com/documentation/Foundation/NSError
func (a SMAppService) UnregisterWithCompletionHandler(handler ErrorHandler) {
	_block0, _ := NewErrorBlock(handler)
	objc.Send[objc.ID](a.ID, objc.Sel("unregisterWithCompletionHandler:"), _block0)
}

// Initializes an app service object with a launch agent with the property
// list name you provide.
//
// plistName: The name of the property list corresponding to the [SMAppService].
//
// # Discussion
//
// The property list name must correspond to a property list in the calling
// app’s `Contents/Library/LaunchAgents` directory.
//
// See: https://developer.apple.com/documentation/ServiceManagement/SMAppService/agent(plistName:)
func (_SMAppServiceClass SMAppServiceClass) AgentServiceWithPlistName(plistName string) SMAppService {
	rv := objc.Send[objc.ID](objc.ID(_SMAppServiceClass.class), objc.Sel("agentServiceWithPlistName:"), objc.String(plistName))
	return SMAppServiceFromID(rv)
}

// Initializes an app service object with a launch daemon with the property
// list name you provide.
//
// plistName: The name of the property list corresponding to the [SMAppService].
//
// # Return Value
//
// An [SMService] object
//
// # Discussion
//
// The property list name must correspond to a property list in the calling
// app’s `Contents/Library/LaunchDaemons` directory
//
// See: https://developer.apple.com/documentation/ServiceManagement/SMAppService/daemon(plistName:)
func (_SMAppServiceClass SMAppServiceClass) DaemonServiceWithPlistName(plistName string) SMAppService {
	rv := objc.Send[objc.ID](objc.ID(_SMAppServiceClass.class), objc.Sel("daemonServiceWithPlistName:"), objc.String(plistName))
	return SMAppServiceFromID(rv)
}

// Initializes an app service object for a login item corresponding to the
// bundle with the identifier you provide.
//
// identifier: The bundle identifier of the helper application.
//
// # Return Value
//
// An [SMService] object.
//
// # Discussion
//
// The property list name must correspond to a property list in the calling
// app’s `Contents/Library/LoginItems` directory
//
// See: https://developer.apple.com/documentation/ServiceManagement/SMAppService/loginItem(identifier:)
func (_SMAppServiceClass SMAppServiceClass) LoginItemServiceWithIdentifier(identifier string) SMAppService {
	rv := objc.Send[objc.ID](objc.ID(_SMAppServiceClass.class), objc.Sel("loginItemServiceWithIdentifier:"), objc.String(identifier))
	return SMAppServiceFromID(rv)
}

// Opens System Settings to the Login Items control panel.
//
// See: https://developer.apple.com/documentation/ServiceManagement/SMAppService/openSystemSettingsLoginItems()
func (_SMAppServiceClass SMAppServiceClass) OpenSystemSettingsLoginItems() {
	objc.Send[objc.ID](objc.ID(_SMAppServiceClass.class), objc.Sel("openSystemSettingsLoginItems"))
}

// Check the authorization status of an earlier OS version login item.
//
// url: The URL of the helper executable’s property list.
//
// # Return Value
//
// One of the [SMAppService.Status] constants that indicate the current
// authorization status.
//
// See: https://developer.apple.com/documentation/ServiceManagement/SMAppService/statusForLegacyPlist(at:)
//
// [SMAppService.Status]: https://developer.apple.com/documentation/ServiceManagement/SMAppService/Status-swift.enum
func (_SMAppServiceClass SMAppServiceClass) StatusForLegacyURL(url foundation.INSURL) SMAppServiceStatus {
	rv := objc.Send[SMAppServiceStatus](objc.ID(_SMAppServiceClass.class), objc.Sel("statusForLegacyURL:"), url)
	return SMAppServiceStatus(rv)
}

// A property that describes registration or authorization state of the
// service.
//
// See: https://developer.apple.com/documentation/ServiceManagement/SMAppService/status-swift.property
func (a SMAppService) Status() SMAppServiceStatus {
	rv := objc.Send[SMAppServiceStatus](a.ID, objc.Sel("status"))
	return SMAppServiceStatus(rv)
}

// An app service object that corresponds to the main application as a login
// item.
//
// # Discussion
//
// Use this [SMAppService] to configure the main app to launch at login.
//
// See: https://developer.apple.com/documentation/ServiceManagement/SMAppService/mainApp
func (_SMAppServiceClass SMAppServiceClass) MainAppService() SMAppService {
	rv := objc.Send[objc.ID](objc.ID(_SMAppServiceClass.class), objc.Sel("mainAppService"))
	return SMAppServiceFromID(objc.ID(rv))
}

// Unregister is a synchronous wrapper around [SMAppService.UnregisterWithCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (a SMAppService) Unregister(ctx context.Context) error {
	done := make(chan error, 1)
	a.UnregisterWithCompletionHandler(func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}
