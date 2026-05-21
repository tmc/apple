// Code generated from Apple documentation for FSKit. DO NOT EDIT.

package fskit

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [FSContainerStatus] class.
var (
	_FSContainerStatusClass     FSContainerStatusClass
	_FSContainerStatusClassOnce sync.Once
)

func getFSContainerStatusClass() FSContainerStatusClass {
	_FSContainerStatusClassOnce.Do(func() {
		_FSContainerStatusClass = FSContainerStatusClass{class: objc.GetClass("FSContainerStatus")}
	})
	return _FSContainerStatusClass
}

// GetFSContainerStatusClass returns the class object for FSContainerStatus.
func GetFSContainerStatusClass() FSContainerStatusClass {
	return getFSContainerStatusClass()
}

type FSContainerStatusClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (fc FSContainerStatusClass) Class() objc.Class {
	return fc.class
}

// Alloc allocates memory for a new instance of the class.
func (fc FSContainerStatusClass) Alloc() FSContainerStatus {
	rv := objc.Send[FSContainerStatus](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// A type that represents a container’s status.
//
// # Overview
//
// This type contains two properties:
//
// - The [FSContainerStatus.State] value that indicates the state of the
// container, such as [FSContainerStateReady] or [FSContainerStateBlocked]. -
// The [FSContainerStatus.Status] is an error (optional in Swift, nullable in
// Objective-C) that provides further information about the state, such as why
// the container is blocked.
//
// Examples of statuses that require intervention include errors that indicate
// the container isn’t ready (POSIX [EAGAIN] or [ENOTCONN]), the container
// needs authentication ([ENEEDAUTH]), or that authentication failed
// ([EAUTH]). The status can also be an informative error, such as the FSKit
// error [FSErrorStatusOperationInProgress].
//
// # Inspecting status properties
//
//   - [FSContainerStatus.State]: A value that represents the container state, such as ready, active, or blocked.
//   - [FSContainerStatus.Status]: An optional error that provides further information about the state.
//
// See: https://developer.apple.com/documentation/FSKit/FSContainerStatus
type FSContainerStatus struct {
	objectivec.Object
}

// FSContainerStatusFromID constructs a [FSContainerStatus] from an objc.ID.
//
// A type that represents a container’s status.
func FSContainerStatusFromID(id objc.ID) FSContainerStatus {
	return FSContainerStatus{objectivec.Object{ID: id}}
}

// NOTE: FSContainerStatus adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [FSContainerStatus] class.
//
// # Inspecting status properties
//
//   - [IFSContainerStatus.State]: A value that represents the container state, such as ready, active, or blocked.
//   - [IFSContainerStatus.Status]: An optional error that provides further information about the state.
//
// See: https://developer.apple.com/documentation/FSKit/FSContainerStatus
type IFSContainerStatus interface {
	objectivec.IObject

	// Topic: Inspecting status properties

	// A value that represents the container state, such as ready, active, or blocked.
	State() FSContainerState
	// An optional error that provides further information about the state.
	Status() foundation.NSError
}

// Init initializes the instance.
func (c FSContainerStatus) Init() FSContainerStatus {
	rv := objc.Send[FSContainerStatus](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c FSContainerStatus) Autorelease() FSContainerStatus {
	rv := objc.Send[FSContainerStatus](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewFSContainerStatus creates a new FSContainerStatus instance.
func NewFSContainerStatus() FSContainerStatus {
	class := getFSContainerStatusClass()
	rv := objc.Send[FSContainerStatus](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Returns a active container status instance with the provided error status.
//
// errorStatus: The error status, if any, for the new instance.
//
// See: https://developer.apple.com/documentation/FSKit/FSContainerStatus/active(status:)
func (_FSContainerStatusClass FSContainerStatusClass) ActiveWithStatus(errorStatus foundation.NSError) FSContainerStatus {
	rv := objc.Send[objc.ID](objc.ID(_FSContainerStatusClass.class), objc.Sel("activeWithStatus:"), errorStatus)
	return FSContainerStatusFromID(rv)
}

// Returns a blocked container status instance with the provided error status.
//
// errorStatus: The error status, if any, for the new instance.
//
// See: https://developer.apple.com/documentation/FSKit/FSContainerStatus/blocked(status:)
func (_FSContainerStatusClass FSContainerStatusClass) BlockedWithStatus(errorStatus foundation.NSError) FSContainerStatus {
	rv := objc.Send[objc.ID](objc.ID(_FSContainerStatusClass.class), objc.Sel("blockedWithStatus:"), errorStatus)
	return FSContainerStatusFromID(rv)
}

// Returns a not-ready container status instance with the provided error
// status.
//
// errorStatus: The error status, if any, for the new instance.
//
// See: https://developer.apple.com/documentation/FSKit/FSContainerStatus/notReady(status:)
func (_FSContainerStatusClass FSContainerStatusClass) NotReadyWithStatus(errorStatus foundation.NSError) FSContainerStatus {
	rv := objc.Send[objc.ID](objc.ID(_FSContainerStatusClass.class), objc.Sel("notReadyWithStatus:"), errorStatus)
	return FSContainerStatusFromID(rv)
}

// Returns a ready container status instance with the provided error status.
//
// errorStatus: The error status, if any, for the new instance.
//
// See: https://developer.apple.com/documentation/FSKit/FSContainerStatus/ready(status:)
func (_FSContainerStatusClass FSContainerStatusClass) ReadyWithStatus(errorStatus foundation.NSError) FSContainerStatus {
	rv := objc.Send[objc.ID](objc.ID(_FSContainerStatusClass.class), objc.Sel("readyWithStatus:"), errorStatus)
	return FSContainerStatusFromID(rv)
}

// A value that represents the container state, such as ready, active, or
// blocked.
//
// See: https://developer.apple.com/documentation/FSKit/FSContainerStatus/state
func (c FSContainerStatus) State() FSContainerState {
	rv := objc.Send[FSContainerState](c.ID, objc.Sel("state"))
	return FSContainerState(rv)
}

// An optional error that provides further information about the state.
//
// See: https://developer.apple.com/documentation/FSKit/FSContainerStatus/status
func (c FSContainerStatus) Status() foundation.NSError {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("status"))
	return foundation.NSErrorFromID(objc.ID(rv))
}

// A status that represents an active container with no error.
//
// # Discussion
//
// This value is a [FSContainerStatus] with a [FSContainerStatus.State] that
// is [FSContainerStatusClass.Active], and has a [FSContainerStatus.Status]
// that is `nil`.
//
// See: https://developer.apple.com/documentation/FSKit/FSContainerStatus/active
func (_FSContainerStatusClass FSContainerStatusClass) Active() FSContainerStatus {
	rv := objc.Send[objc.ID](objc.ID(_FSContainerStatusClass.class), objc.Sel("active"))
	return FSContainerStatusFromID(objc.ID(rv))
}

// A status that represents a ready container with no error.
//
// # Discussion
//
// This value is a [FSContainerStatus] with a [FSContainerStatus.State] that
// is [FSContainerStatusClass.Ready], and a [FSContainerStatus.Status] that is
// `nil`.
//
// See: https://developer.apple.com/documentation/FSKit/FSContainerStatus/ready
func (_FSContainerStatusClass FSContainerStatusClass) Ready() FSContainerStatus {
	rv := objc.Send[objc.ID](objc.ID(_FSContainerStatusClass.class), objc.Sel("ready"))
	return FSContainerStatusFromID(objc.ID(rv))
}
