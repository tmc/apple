// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSWorkspaceAuthorization] class.
var (
	_NSWorkspaceAuthorizationClass     NSWorkspaceAuthorizationClass
	_NSWorkspaceAuthorizationClassOnce sync.Once
)

func getNSWorkspaceAuthorizationClass() NSWorkspaceAuthorizationClass {
	_NSWorkspaceAuthorizationClassOnce.Do(func() {
		_NSWorkspaceAuthorizationClass = NSWorkspaceAuthorizationClass{class: objc.GetClass("NSWorkspaceAuthorization")}
	})
	return _NSWorkspaceAuthorizationClass
}

// GetNSWorkspaceAuthorizationClass returns the class object for NSWorkspaceAuthorization.
func GetNSWorkspaceAuthorizationClass() NSWorkspaceAuthorizationClass {
	return getNSWorkspaceAuthorizationClass()
}

type NSWorkspaceAuthorizationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSWorkspaceAuthorizationClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSWorkspaceAuthorizationClass) Alloc() NSWorkspaceAuthorization {
	rv := objc.Send[NSWorkspaceAuthorization](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// The authorization granted to the app by the user.
//
// # Overview
//
// To enable your app to prompt the user for these file permissions, you must
// have a Privileged File Operation entitlement. If you have an app on the Mac
// App Store or plan to submit your app for review, you can [request this
// entitlement].
//
// See: https://developer.apple.com/documentation/AppKit/NSWorkspace/Authorization
//
// [request this entitlement]: https://developer.apple.com/go/?id=workspace-authorization
type NSWorkspaceAuthorization struct {
	objectivec.Object
}

// NSWorkspaceAuthorizationFromID constructs a [NSWorkspaceAuthorization] from an objc.ID.
//
// The authorization granted to the app by the user.
func NSWorkspaceAuthorizationFromID(id objc.ID) NSWorkspaceAuthorization {
	return NSWorkspaceAuthorization{objectivec.Object{ID: id}}
}

// NOTE: NSWorkspaceAuthorization adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSWorkspaceAuthorization] class.
//
// See: https://developer.apple.com/documentation/AppKit/NSWorkspace/Authorization
type INSWorkspaceAuthorization interface {
	objectivec.IObject
}

// Init initializes the instance.
func (w NSWorkspaceAuthorization) Init() NSWorkspaceAuthorization {
	rv := objc.Send[NSWorkspaceAuthorization](w.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (w NSWorkspaceAuthorization) Autorelease() NSWorkspaceAuthorization {
	rv := objc.Send[NSWorkspaceAuthorization](w.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSWorkspaceAuthorization creates a new NSWorkspaceAuthorization instance.
func NewNSWorkspaceAuthorization() NSWorkspaceAuthorization {
	class := getNSWorkspaceAuthorizationClass()
	rv := objc.Send[NSWorkspaceAuthorization](objc.ID(class.class), objc.Sel("new"))
	return rv
}
