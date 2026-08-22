// Code generated from Apple documentation for CoreLocation. DO NOT EDIT.

package corelocation

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CLMonitoringEvent] class.
var (
	_CLMonitoringEventClass     CLMonitoringEventClass
	_CLMonitoringEventClassOnce sync.Once
)

func getCLMonitoringEventClass() CLMonitoringEventClass {
	_CLMonitoringEventClassOnce.Do(func() {
		_CLMonitoringEventClass = CLMonitoringEventClass{class: objc.GetClass("CLMonitoringEvent")}
	})
	return _CLMonitoringEventClass
}

// GetCLMonitoringEventClass returns the class object for CLMonitoringEvent.
func GetCLMonitoringEventClass() CLMonitoringEventClass {
	return getCLMonitoringEventClass()
}

type CLMonitoringEventClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CLMonitoringEventClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CLMonitoringEventClass) Alloc() CLMonitoringEvent {
	rv := objc.Send[CLMonitoringEvent](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// The object that the framework passes to the monitor’s callback handler
// upon receiving an event.
//
// # Overview
//
// Instances of [CLMonitoringEvent] contain detailed information about an
// event in the monitoring of a [CLCondition] by a [CLMonitor].
//
// # Event states
//
//   - [CLMonitoringEvent.AccuracyLimited]: A Boolean value that indicates whether the app receives accuracy-limited location updates.
//   - [CLMonitoringEvent.AuthorizationDenied]: A Boolean value that indicates whether the app has local authorization.
//   - [CLMonitoringEvent.AuthorizationDeniedGlobally]: A Boolean value that indicates whether the app has system-wide authorization.
//   - [CLMonitoringEvent.AuthorizationRequestInProgress]
//   - [CLMonitoringEvent.AuthorizationRestricted]: A Boolean value that indicates whether the app can make authorization changes.
//   - [CLMonitoringEvent.ConditionLimitExceeded]: A Boolean value that indicates whether the app receives location updates based on other monitoring conditions.
//   - [CLMonitoringEvent.ConditionUnsupported]: A Boolean value that indicates whether the app receives location updates based on the supported condition.
//   - [CLMonitoringEvent.InsufficientlyInUse]: A Boolean value that indicates whether the app receives location updates because it’s insufficiently in use.
//   - [CLMonitoringEvent.PersistenceUnavailable]: A Boolean value that indicates whether it receives location updates based on successful persistence.
//   - [CLMonitoringEvent.ServiceSessionRequired]
//
// # Event properties
//
//   - [CLMonitoringEvent.Date]: The date the event occurs.
//   - [CLMonitoringEvent.Identifier]: A string that represents the identifier of a monitored condition.
//   - [CLMonitoringEvent.Refinement]: An optional instance of a condition that represents the most specific condition to that this event can apply to.
//   - [CLMonitoringEvent.State]: The state of the condition at the time of the event.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLMonitoringEvent
//
// [CLCondition]: https://developer.apple.com/documentation/CoreLocation/CLCondition-swift.protocol
// [CLMonitor]: https://developer.apple.com/documentation/CoreLocation/CLMonitor-2r51v
type CLMonitoringEvent struct {
	objectivec.Object
}

// CLMonitoringEventFromID constructs a [CLMonitoringEvent] from an objc.ID.
//
// The object that the framework passes to the monitor’s callback handler
// upon receiving an event.
func CLMonitoringEventFromID(id objc.ID) CLMonitoringEvent {
	return CLMonitoringEvent{objectivec.Object{ID: id}}
}

// NOTE: CLMonitoringEvent adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CLMonitoringEvent] class.
//
// # Event states
//
//   - [ICLMonitoringEvent.AccuracyLimited]: A Boolean value that indicates whether the app receives accuracy-limited location updates.
//   - [ICLMonitoringEvent.AuthorizationDenied]: A Boolean value that indicates whether the app has local authorization.
//   - [ICLMonitoringEvent.AuthorizationDeniedGlobally]: A Boolean value that indicates whether the app has system-wide authorization.
//   - [ICLMonitoringEvent.AuthorizationRequestInProgress]
//   - [ICLMonitoringEvent.AuthorizationRestricted]: A Boolean value that indicates whether the app can make authorization changes.
//   - [ICLMonitoringEvent.ConditionLimitExceeded]: A Boolean value that indicates whether the app receives location updates based on other monitoring conditions.
//   - [ICLMonitoringEvent.ConditionUnsupported]: A Boolean value that indicates whether the app receives location updates based on the supported condition.
//   - [ICLMonitoringEvent.InsufficientlyInUse]: A Boolean value that indicates whether the app receives location updates because it’s insufficiently in use.
//   - [ICLMonitoringEvent.PersistenceUnavailable]: A Boolean value that indicates whether it receives location updates based on successful persistence.
//   - [ICLMonitoringEvent.ServiceSessionRequired]
//
// # Event properties
//
//   - [ICLMonitoringEvent.Date]: The date the event occurs.
//   - [ICLMonitoringEvent.Identifier]: A string that represents the identifier of a monitored condition.
//   - [ICLMonitoringEvent.Refinement]: An optional instance of a condition that represents the most specific condition to that this event can apply to.
//   - [ICLMonitoringEvent.State]: The state of the condition at the time of the event.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLMonitoringEvent
type ICLMonitoringEvent interface {
	objectivec.IObject

	// Topic: Event states

	// A Boolean value that indicates whether the app receives accuracy-limited location updates.
	AccuracyLimited() bool
	// A Boolean value that indicates whether the app has local authorization.
	AuthorizationDenied() bool
	// A Boolean value that indicates whether the app has system-wide authorization.
	AuthorizationDeniedGlobally() bool
	AuthorizationRequestInProgress() bool
	// A Boolean value that indicates whether the app can make authorization changes.
	AuthorizationRestricted() bool
	// A Boolean value that indicates whether the app receives location updates based on other monitoring conditions.
	ConditionLimitExceeded() bool
	// A Boolean value that indicates whether the app receives location updates based on the supported condition.
	ConditionUnsupported() bool
	// A Boolean value that indicates whether the app receives location updates because it’s insufficiently in use.
	InsufficientlyInUse() bool
	// A Boolean value that indicates whether it receives location updates based on successful persistence.
	PersistenceUnavailable() bool
	ServiceSessionRequired() bool

	// Topic: Event properties

	// The date the event occurs.
	Date() foundation.NSDate
	// A string that represents the identifier of a monitored condition.
	Identifier() string
	// An optional instance of a condition that represents the most specific condition to that this event can apply to.
	Refinement() ICLCondition
	// The state of the condition at the time of the event.
	State() CLMonitoringState

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (m CLMonitoringEvent) Init() CLMonitoringEvent {
	rv := objc.Send[CLMonitoringEvent](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m CLMonitoringEvent) Autorelease() CLMonitoringEvent {
	rv := objc.Send[CLMonitoringEvent](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewCLMonitoringEvent creates a new CLMonitoringEvent instance.
func NewCLMonitoringEvent() CLMonitoringEvent {
	class := getCLMonitoringEventClass()
	rv := objc.Send[CLMonitoringEvent](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (m CLMonitoringEvent) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](m.ID, objc.Sel("encodeWithCoder:"), coder)
}

// A Boolean value that indicates whether the app receives accuracy-limited
// location updates.
//
// # Discussion
//
// If this property is true, then the app receives accuracy-limited location
// updates.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLMonitoringEvent/accuracyLimited
func (m CLMonitoringEvent) AccuracyLimited() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("accuracyLimited"))
	return rv
}

// A Boolean value that indicates whether the app has local authorization.
//
// # Discussion
//
// If this property is true, then the app isn’t receiving location updates
// because it’s denied local authorization.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLMonitoringEvent/authorizationDenied
func (m CLMonitoringEvent) AuthorizationDenied() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("authorizationDenied"))
	return rv
}

// A Boolean value that indicates whether the app has system-wide
// authorization.
//
// # Discussion
//
// If this property is true, then the app isn’t receiving location updates
// because it’s denied system-wide authorization.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLMonitoringEvent/authorizationDeniedGlobally
func (m CLMonitoringEvent) AuthorizationDeniedGlobally() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("authorizationDeniedGlobally"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreLocation/CLMonitoringEvent/authorizationRequestInProgress
func (m CLMonitoringEvent) AuthorizationRequestInProgress() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("authorizationRequestInProgress"))
	return rv
}

// A Boolean value that indicates whether the app can make authorization
// changes.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLMonitoringEvent/authorizationRestricted
func (m CLMonitoringEvent) AuthorizationRestricted() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("authorizationRestricted"))
	return rv
}

// A Boolean value that indicates whether the app receives location updates
// based on other monitoring conditions.
//
// # Discussion
//
// If this property is true, then the app isn’t receiving location updates
// because it’s monitoring too many other conditions of this type.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLMonitoringEvent/conditionLimitExceeded
func (m CLMonitoringEvent) ConditionLimitExceeded() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("conditionLimitExceeded"))
	return rv
}

// A Boolean value that indicates whether the app receives location updates
// based on the supported condition.
//
// # Discussion
//
// If this property is true, then the app isn’t receiving location updates
// because it’s monitoring a type of condition that isn’t supported.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLMonitoringEvent/conditionUnsupported
func (m CLMonitoringEvent) ConditionUnsupported() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("conditionUnsupported"))
	return rv
}

// A Boolean value that indicates whether the app receives location updates
// because it’s insufficiently in use.
//
// # Discussion
//
// If this property is true, then the app isn’t receiving location updates
// because it’s insufficiently in use.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLMonitoringEvent/insufficientlyInUse
func (m CLMonitoringEvent) InsufficientlyInUse() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("insufficientlyInUse"))
	return rv
}

// A Boolean value that indicates whether it receives location updates based
// on successful persistence.
//
// # Discussion
//
// If this property is true, then location updates are suspended because the
// app has a persistence failure.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLMonitoringEvent/persistenceUnavailable
func (m CLMonitoringEvent) PersistenceUnavailable() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("persistenceUnavailable"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreLocation/CLMonitoringEvent/serviceSessionRequired
func (m CLMonitoringEvent) ServiceSessionRequired() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("serviceSessionRequired"))
	return rv
}

// The date the event occurs.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLMonitoringEvent/date
func (m CLMonitoringEvent) Date() foundation.NSDate {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("date"))
	return foundation.NSDateFromID(objc.ID(rv))
}

// A string that represents the identifier of a monitored condition.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLMonitoringEvent/identifier
func (m CLMonitoringEvent) Identifier() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("identifier"))
	return foundation.NSStringFromID(rv).String()
}

// An optional instance of a condition that represents the most specific
// condition to that this event can apply to.
//
// # Discussion
//
// The type of the refinement condition depends on the monitored condition
// itself.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLMonitoringEvent/refinement
func (m CLMonitoringEvent) Refinement() ICLCondition {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("refinement"))
	return CLConditionFromID(objc.ID(rv))
}

// The state of the condition at the time of the event.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLMonitoringEvent/state
func (m CLMonitoringEvent) State() CLMonitoringState {
	rv := objc.Send[CLMonitoringState](m.ID, objc.Sel("state"))
	return CLMonitoringState(rv)
}
