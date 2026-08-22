// Code generated from Apple documentation for CoreLocation. DO NOT EDIT.

package corelocation

import (
	"fmt"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

var _ = fmt.Sprintf

// The methods you use to receive events from an associated location-manager object.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLLocationManagerDelegate
type CLLocationManagerDelegate interface {
	objectivec.IObject
}

// CLLocationManagerDelegateObject wraps an existing Objective-C object that conforms to the CLLocationManagerDelegate protocol.
type CLLocationManagerDelegateObject struct {
	objectivec.Object
}

func (o CLLocationManagerDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}

// CLLocationManagerDelegateObjectFromID constructs a [CLLocationManagerDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CLLocationManagerDelegateObjectFromID(id objc.ID) CLLocationManagerDelegateObject {
	return CLLocationManagerDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Tells the delegate when the app creates the location manager and when the
// authorization status changes.
//
// manager: The location manager object reporting the event.
//
// # Discussion
//
// The system calls this method when the app creates the related object’s
// [CLLocationManager] instance, and when the app’s authorization status
// changes. The status informs the app whether it can access the user’s
// location.
//
// Use this delegate method to manage your app’s state changes in response
// to its ability to use location information. For example, you may wish to
// enable or disable your app’s location-related features, as appropriate.
// To determine the app’s current authorization, read the new value of the
// [CLLocationManager.AuthorizationStatus] and
// [CLLocationManager.AccuracyAuthorization] properties of the location
// manager.
//
// If the user’s choice doesn’t change the authorization status after you
// call the [CLLocationManager.RequestWhenInUseAuthorization] or
// [CLLocationManager.RequestAlwaysAuthorization] method, the location manager
// doesn’t report the current authorization status to this method—the
// location manager only reports changes. For example, the location manager
// calls this method when the status changes from
// [KCLAuthorizationStatusNotDetermined] to
// [KCLAuthorizationStatusAuthorizedWhenInUse].
//
// # Events that Cause Authorization Status Changes
//
// An app’s authorization status changes in response to users’ actions.
// Users can change permission for apps to use location information at any
// time. The user can:
//
// - Change an app’s location authorization in Settings > Privacy > Location
// Services, or in Settings > (the app) > Location Services. - Turn location
// services on or off globally in Settings > Privacy > Location Services. -
// Choose Reset Location & Privacy in Settings > General > Reset.
//
// A user’s response to location manager prompts can also change
// authorization status. For instance, users may change the authorization
// status by responding to the prompts initiated by calls to
// [CLLocationManager.RequestWhenInUseAuthorization] or
// [CLLocationManager.RequestAlwaysAuthorization] methods. For apps with
// Always authorization, users may change the authorization status to When In
// Use when responding to the location usage reminder alert.
//
// When an app has temporary authorization, the authorization changes when the
// user ceases to use the app.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLLocationManagerDelegate/locationManagerDidChangeAuthorization(_:)
func (o CLLocationManagerDelegateObject) LocationManagerDidChangeAuthorization(manager ICLLocationManager) {
	objc.Send[struct{}](o.ID, objc.Sel("locationManagerDidChangeAuthorization:"), manager)
}

// Tells the delegate that the location manager was unable to retrieve a
// location value.
//
// manager: The location manager object that was unable to retrieve the location.
//
// error: The error object containing the reason the location or heading could not be
// retrieved.
//
// # Discussion
//
// If you do not implement this method, Core Location throws an exception when
// attempting to use location services.
//
// The location manager calls this method when it encounters an error trying
// to get the location or heading data. If the location service is unable to
// retrieve a location right away, it reports a [KCLErrorLocationUnknown]
// error and keeps trying. In such a situation, you can simply ignore the
// error and wait for a new event. If a heading could not be determined
// because of strong interference from nearby magnetic fields, this method
// returns [KCLErrorHeadingFailure].
//
// If the user denies your app’s use of the location service, this method
// reports a [KCLErrorDenied] error. Upon receiving such an error, you should
// stop the location service.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLLocationManagerDelegate/locationManager(_:didFailWithError:)
func (o CLLocationManagerDelegateObject) LocationManagerDidFailWithError(manager ICLLocationManager, error_ foundation.NSError) {
	objc.Send[struct{}](o.ID, objc.Sel("locationManager:didFailWithError:"), manager, error_)
}

// Tells the delegate that new location data is available.
//
// manager: The location manager object that generated the update event.
//
// locations: An array of [CLLocation] objects containing the location data. This array
// always contains at least one object representing the current location. If
// updates were deferred or if multiple locations arrived before they could be
// delivered, the array may contain additional entries. The objects in the
// array are organized in the order in which they occurred. Therefore, the
// most recent location update is at the end of the array.
//
// # Discussion
//
// Implementation of this method is optional but recommended.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLLocationManagerDelegate/locationManager(_:didUpdateLocations:)
func (o CLLocationManagerDelegateObject) LocationManagerDidUpdateLocations(manager ICLLocationManager, locations []CLLocation) {
	objc.Send[struct{}](o.ID, objc.Sel("locationManager:didUpdateLocations:"), manager, objectivec.IObjectSliceToNSArray(locations))
}

// Tells the delegate that a new location value is available.
//
// manager: The location manager object that generated the update event.
//
// newLocation: The new location data.
//
// oldLocation: The location data from the previous update. If this is the first update
// event delivered by this location manager, this parameter is `nil`.
//
// # Discussion
//
// By the time this message is delivered to your delegate, the new location
// data is also available directly from the [CLLocationManager] object. The
// `newLocation` parameter may contain the data that was cached from a
// previous usage of the location service. You can use the
// [CLLocation.Timestamp] property of the location object to determine how
// recent the location data is.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLLocationManagerDelegate/locationManager(_:didUpdateTo:from:)
func (o CLLocationManagerDelegateObject) LocationManagerDidUpdateToLocationFromLocation(manager ICLLocationManager, newLocation ICLLocation, oldLocation ICLLocation) {
	objc.Send[struct{}](o.ID, objc.Sel("locationManager:didUpdateToLocation:fromLocation:"), manager, newLocation, oldLocation)
}

// Tells the delegate that updates will no longer be deferred.
//
// manager: The location manager object that generated the update event.
//
// error: The error object containing the reason deferred location updates could not
// be delivered.
//
// # Discussion
//
// The location manager object calls this method to let you know that it has
// stopped deferring the delivery of location events. The manager may call
// this method for any number of reasons. For example, it calls it when you
// stop location updates altogether, when you ask the location manager to
// disallow deferred updates, or when a condition for deferring updates (such
// as exceeding a timeout or distance parameter) is met.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLLocationManagerDelegate/locationManager(_:didFinishDeferredUpdatesWithError:)
func (o CLLocationManagerDelegateObject) LocationManagerDidFinishDeferredUpdatesWithError(manager ICLLocationManager, error_ foundation.NSError) {
	objc.Send[struct{}](o.ID, objc.Sel("locationManager:didFinishDeferredUpdatesWithError:"), manager, error_)
}

// Tells the delegate that location updates were paused.
//
// manager: The location manager object that paused the delivery of events.
//
// # Discussion
//
// When the location manager detects that the device’s location is not
// changing, it can pause the delivery of updates in order to shut down the
// appropriate hardware and save power. When it does this, it calls this
// method to let your app know that this has happened.
//
// After a pause occurs, it is your responsibility to restart location
// services again at an appropriate time. You might use your implementation of
// this method to start region monitoring at the user’s current location or
// enable the visits location service to determine when the user starts moving
// again. Another alternative is to restart location services immediately with
// a reduced accuracy (which can save power) and then return to a greater
// accuracy only after the user starts moving again.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLLocationManagerDelegate/locationManagerDidPauseLocationUpdates(_:)
func (o CLLocationManagerDelegateObject) LocationManagerDidPauseLocationUpdates(manager ICLLocationManager) {
	objc.Send[struct{}](o.ID, objc.Sel("locationManagerDidPauseLocationUpdates:"), manager)
}

// Tells the delegate that the delivery of location updates has resumed.
//
// manager: The location manager that resumed the delivery of events.
//
// # Discussion
//
// When you restart location services after an automatic pause, Core Location
// calls this method to notify your app that services have resumed. You are
// responsible for restarting location services in your app. Core Location
// does not resume updates automatically after it pauses them. For tips on how
// to restart location services when a pause occurs, see the discussion of the
// [LocationManagerDidPauseLocationUpdates] method.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLLocationManagerDelegate/locationManagerDidResumeLocationUpdates(_:)
func (o CLLocationManagerDelegateObject) LocationManagerDidResumeLocationUpdates(manager ICLLocationManager) {
	objc.Send[struct{}](o.ID, objc.Sel("locationManagerDidResumeLocationUpdates:"), manager)
}

// Tells the delegate that a new visit-related event was received.
//
// manager: The location manager object reporting the event.
//
// visit: The visit object that contains the information about the event.
//
// # Discussion
//
// The location manager calls this method whenever it has new visit event to
// report to your app.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLLocationManagerDelegate/locationManager(_:didVisit:)
func (o CLLocationManagerDelegateObject) LocationManagerDidVisit(manager ICLLocationManager, visit ICLVisit) {
	objc.Send[struct{}](o.ID, objc.Sel("locationManager:didVisit:"), manager, visit)
}

// Tells the delegate that the location manager received updated heading
// information.
//
// manager: The location manager object that generated the update event.
//
// newHeading: The new heading data.
//
// # Discussion
//
// Implementation of this method is optional but expected if you start heading
// updates using the [CLLocationManager.StartUpdatingHeading] method.
//
// The location manager object calls this method after you initially start the
// heading service. Subsequent events are delivered when the previously
// reported value changes by more than the value specified in the
// [CLLocationManager.HeadingFilter] property of the location manager object.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLLocationManagerDelegate/locationManager(_:didUpdateHeading:)
func (o CLLocationManagerDelegateObject) LocationManagerDidUpdateHeading(manager ICLLocationManager, newHeading ICLHeading) {
	objc.Send[struct{}](o.ID, objc.Sel("locationManager:didUpdateHeading:"), manager, newHeading)
}

// Asks the delegate whether the heading calibration alert should be
// displayed.
//
// manager: The location manager object coordinating the display of the heading
// calibration alert.
//
// # Return Value
//
// true if you want to allow the heading calibration alert to be displayed;
// false if you do not.
//
// # Discussion
//
// Core Location may call this method in an effort to calibrate the onboard
// hardware used to determine heading values. Typically, Core Location calls
// this method at the following times:
//
// - The first time heading updates are ever requested - When Core Location
// observes a significant change in magnitude or inclination of the observed
// magnetic field
//
// If you return true from this method, Core Location displays the heading
// calibration alert on top of the current window immediately. The calibration
// alert prompts the user to move the device in a particular pattern so that
// Core Location can distinguish between the Earth’s magnetic field and any
// local magnetic fields. The alert remains visible until calibration is
// complete or until you explicitly dismiss it by calling the
// [CLLocationManager.DismissHeadingCalibrationDisplay] method. In the latter
// case, you can use this method to set up a timer and dismiss the interface
// after a specified amount of time has elapsed.
//
// If you return false from this method or do not provide an implementation
// for it in your delegate, Core Location does not display the heading
// calibration alert. Even if the alert is not displayed, calibration can
// still occur naturally when any interfering magnetic fields move away from
// the device. However, if the device is unable to calibrate itself for any
// reason, the value in the [CLHeading.HeadingAccuracy] property of any
// subsequent events will reflect the uncalibrated readings.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLLocationManagerDelegate/locationManagerShouldDisplayHeadingCalibration(_:)
func (o CLLocationManagerDelegateObject) LocationManagerShouldDisplayHeadingCalibration(manager ICLLocationManager) bool {
	rv := objc.Send[bool](o.ID, objc.Sel("locationManagerShouldDisplayHeadingCalibration:"), manager)
	return rv
}

// Tells the delegate that the user entered the specified region.
//
// manager: The location manager object reporting the event.
//
// region: An object containing information about the region that was entered.
//
// # Discussion
//
// Because regions are a shared application resource, every active location
// manager object delivers this message to its associated delegate. It
// doesn’t matter which location manager actually registered the specified
// region. If multiple location managers share a delegate object, that
// delegate receives the message multiple times.
//
// The region object provided may not be the same one that was registered. As
// a result, you should never perform pointer-level comparisons to determine
// equality. Instead, use the region’s identifier string to determine if
// your delegate should respond.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLLocationManagerDelegate/locationManager(_:didEnterRegion:)
func (o CLLocationManagerDelegateObject) LocationManagerDidEnterRegion(manager ICLLocationManager, region ICLRegion) {
	objc.Send[struct{}](o.ID, objc.Sel("locationManager:didEnterRegion:"), manager, region)
}

// Tells the delegate that the user left the specified region.
//
// manager: The location manager object reporting the event.
//
// region: An object containing information about the region that was exited.
//
// # Discussion
//
// Because regions are a shared application resource, every active location
// manager object delivers this message to its associated delegate. It
// doesn’t matter which location manager actually registered the specified
// region. If multiple location managers share a delegate object, that
// delegate receives the message multiple times.
//
// The region object provided may not be the same one that was registered. As
// a result, you should never perform pointer-level comparisons to determine
// equality. Instead, use the region’s identifier string to determine if
// your delegate should respond.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLLocationManagerDelegate/locationManager(_:didExitRegion:)
func (o CLLocationManagerDelegateObject) LocationManagerDidExitRegion(manager ICLLocationManager, region ICLRegion) {
	objc.Send[struct{}](o.ID, objc.Sel("locationManager:didExitRegion:"), manager, region)
}

// Tells the delegate about the state of the specified region.
//
// manager: The location manager object reporting the event.
//
// state: The state of the specified region. For a list of possible values, see the
// [CLRegionState] type.
//
// region: The region whose state was determined.
//
// # Discussion
//
// The location manager calls this method whenever there is a boundary
// transition for a region. It calls this method in addition to calling the
// [LocationManagerDidEnterRegion] and [LocationManagerDidExitRegion] methods.
// The location manager also calls this method in response to a call to its
// [CLLocationManager.RequestStateForRegion] method, which runs
// asynchronously.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLLocationManagerDelegate/locationManager(_:didDetermineState:for:)
//
// [CLRegionState]: https://developer.apple.com/documentation/CoreLocation/CLRegionState
func (o CLLocationManagerDelegateObject) LocationManagerDidDetermineStateForRegion(manager ICLLocationManager, state CLRegionState, region ICLRegion) {
	objc.Send[struct{}](o.ID, objc.Sel("locationManager:didDetermineState:forRegion:"), manager, state, region)
}

// Tells the delegate that a region monitoring error occurred.
//
// manager: The location manager object reporting the event.
//
// region: The region for which the error occurred.
//
// error: An error object containing the error code that indicates why region
// monitoring failed.
//
// # Discussion
//
// If an error occurs while trying to monitor a given region, the location
// manager sends this message to its delegate. Region monitoring might fail
// because the region itself cannot be monitored or because there was a more
// general failure in configuring the region monitoring service.
//
// Although implementation of this method is optional, it is recommended that
// you implement it if you use region monitoring in your application.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLLocationManagerDelegate/locationManager(_:monitoringDidFailFor:withError:)
func (o CLLocationManagerDelegateObject) LocationManagerMonitoringDidFailForRegionWithError(manager ICLLocationManager, region ICLRegion, error_ foundation.NSError) {
	objc.Send[struct{}](o.ID, objc.Sel("locationManager:monitoringDidFailForRegion:withError:"), manager, region, error_)
}

// Tells the delegate that a new region is being monitored.
//
// manager: The location manager object reporting the event.
//
// region: The region that is being monitored.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLLocationManagerDelegate/locationManager(_:didStartMonitoringFor:)
func (o CLLocationManagerDelegateObject) LocationManagerDidStartMonitoringForRegion(manager ICLLocationManager, region ICLRegion) {
	objc.Send[struct{}](o.ID, objc.Sel("locationManager:didStartMonitoringForRegion:"), manager, region)
}

// Tells the delegate that the location manager detected at least one beacon
// that satisfies the provided constraint.
//
// manager: The [CLLocationManager] that corresponds to this delegate.
//
// beacons: An array of [CLBeacon] objects.
//
// beaconConstraint: The [CLBeaconIdentityConstraint] that describes the characteristics of the
// beacons the location manager is looking for.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLLocationManagerDelegate/locationManager(_:didRange:satisfying:)
func (o CLLocationManagerDelegateObject) LocationManagerDidRangeBeaconsSatisfyingConstraint(manager ICLLocationManager, beacons []CLBeacon, beaconConstraint ICLBeaconIdentityConstraint) {
	objc.Send[struct{}](o.ID, objc.Sel("locationManager:didRangeBeacons:satisfyingConstraint:"), manager, objectivec.IObjectSliceToNSArray(beacons), beaconConstraint)
}

// Tells the delegate that the location manager couldn’t detect any beacons
// that satisfy the provided constraint.
//
// manager: The [CLLocationManager] that corresponds to this delegate.
//
// beaconConstraint: The [CLBeaconIdentityConstraint] that describes the characteristics of the
// beacons the location manager is looking for.
//
// error: An [NSError] object that describes the error.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLLocationManagerDelegate/locationManager(_:didFailRangingFor:error:)
//
// [NSError]: https://developer.apple.com/documentation/Foundation/NSError
func (o CLLocationManagerDelegateObject) LocationManagerDidFailRangingBeaconsForConstraintError(manager ICLLocationManager, beaconConstraint ICLBeaconIdentityConstraint, error_ foundation.NSError) {
	objc.Send[struct{}](o.ID, objc.Sel("locationManager:didFailRangingBeaconsForConstraint:error:"), manager, beaconConstraint, error_)
}

// CLLocationManagerDelegateConfig holds optional typed callbacks for [CLLocationManagerDelegate] methods.
// Set non-nil fields to register the corresponding Objective-C delegate method.
// Methods with nil callbacks are not registered, so [NSObject.RespondsToSelector]
// returns false for them — matching the Objective-C delegate pattern exactly.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/corelocation/cllocationmanagerdelegate
type CLLocationManagerDelegateConfig struct {

	// Responding to authorization changes
	// LocationManagerDidChangeAuthorization — Tells the delegate when the app creates the location manager and when the authorization status changes.
	LocationManagerDidChangeAuthorization func(manager CLLocationManager)

	// Handling errors
	// LocationManagerDidFailWithError — Tells the delegate that the location manager was unable to retrieve a location value.
	LocationManagerDidFailWithError func(manager CLLocationManager, error_ foundation.NSError)

	// Receiving location updates
	// LocationManagerDidFinishDeferredUpdatesWithError — Tells the delegate that updates will no longer be deferred.
	LocationManagerDidFinishDeferredUpdatesWithError func(manager CLLocationManager, error_ foundation.NSError)

	// Pausing location updates
	// LocationManagerDidPauseLocationUpdates — Tells the delegate that location updates were paused.
	LocationManagerDidPauseLocationUpdates func(manager CLLocationManager)
	// LocationManagerDidResumeLocationUpdates — Tells the delegate that the delivery of location updates has resumed.
	LocationManagerDidResumeLocationUpdates func(manager CLLocationManager)

	// Receiving visit updates
	// LocationManagerDidVisit — Tells the delegate that a new visit-related event was received.
	LocationManagerDidVisit func(manager CLLocationManager, visit CLVisit)

	// Receiving heading updates
	// LocationManagerDidUpdateHeading — Tells the delegate that the location manager received updated heading information.
	LocationManagerDidUpdateHeading func(manager CLLocationManager, newHeading CLHeading)
	// LocationManagerShouldDisplayHeadingCalibration — Asks the delegate whether the heading calibration alert should be displayed.
	LocationManagerShouldDisplayHeadingCalibration func(manager CLLocationManager) bool

	// Receiving region-related updates
	// LocationManagerDidEnterRegion — Tells the delegate that the user entered the specified region.
	LocationManagerDidEnterRegion func(manager CLLocationManager, region CLRegion)
	// LocationManagerDidExitRegion — Tells the delegate that the user left the specified region.
	LocationManagerDidExitRegion func(manager CLLocationManager, region CLRegion)

	// Other Methods
	// LocationManagerDidUpdateToLocationFromLocation — Tells the delegate that a new location value is available.
	LocationManagerDidUpdateToLocationFromLocation func(manager CLLocationManager, newLocation CLLocation, oldLocation CLLocation)
	// LocationManagerDidDetermineStateForRegion — Tells the delegate about the state of the specified region.
	LocationManagerDidDetermineStateForRegion func(manager CLLocationManager, state CLRegionState, region CLRegion)
	// LocationManagerMonitoringDidFailForRegionWithError — Tells the delegate that a region monitoring error occurred.
	LocationManagerMonitoringDidFailForRegionWithError func(manager CLLocationManager, region CLRegion, error_ foundation.NSError)
	// LocationManagerDidStartMonitoringForRegion — Tells the delegate that a new region is being monitored.
	LocationManagerDidStartMonitoringForRegion func(manager CLLocationManager, region CLRegion)
	// LocationManagerDidFailRangingBeaconsForConstraintError — Tells the delegate that the location manager couldn’t detect any beacons that satisfy the provided constraint.
	LocationManagerDidFailRangingBeaconsForConstraintError func(manager CLLocationManager, beaconConstraint CLBeaconIdentityConstraint, error_ foundation.NSError)
}

// NewCLLocationManagerDelegate creates an Objective-C object implementing the [CLLocationManagerDelegate] protocol.
//
// Each call registers a unique Objective-C class containing only the methods
// set in config. This means [NSObject.RespondsToSelector] works correctly
// for optional delegate methods — only non-nil callbacks are registered.
//
// The returned [CLLocationManagerDelegateObject] satisfies the [CLLocationManagerDelegate] interface
// and can be passed directly to SetDelegate and similar methods.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/corelocation/cllocationmanagerdelegate
func NewCLLocationManagerDelegate(config CLLocationManagerDelegateConfig) CLLocationManagerDelegateObject {
	n := delegateClassCounter.Add(1)
	className := fmt.Sprintf("GoCLLocationManagerDelegate_%d", n)

	var methods []objc.MethodDef

	if config.LocationManagerDidChangeAuthorization != nil {
		fn := config.LocationManagerDidChangeAuthorization
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("locationManagerDidChangeAuthorization:"),
			Fn: func(self objc.ID, _cmd objc.SEL, managerID objc.ID) {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("CLLocationManagerDelegate", "locationManagerDidChangeAuthorization:")
					}
				}()
				manager := CLLocationManagerFromID(managerID)
				fn(manager)
				_delegateDone = true
			},
		})
	}

	if config.LocationManagerDidFailWithError != nil {
		fn := config.LocationManagerDidFailWithError
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("locationManager:didFailWithError:"),
			Fn: func(self objc.ID, _cmd objc.SEL, managerID objc.ID, error_ID objc.ID) {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("CLLocationManagerDelegate", "locationManager:didFailWithError:")
					}
				}()
				manager := CLLocationManagerFromID(managerID)
				error_ := foundation.NSErrorFromID(error_ID)
				fn(manager, error_)
				_delegateDone = true
			},
		})
	}

	if config.LocationManagerDidUpdateToLocationFromLocation != nil {
		fn := config.LocationManagerDidUpdateToLocationFromLocation
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("locationManager:didUpdateToLocation:fromLocation:"),
			Fn: func(self objc.ID, _cmd objc.SEL, managerID objc.ID, newLocationID objc.ID, oldLocationID objc.ID) {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("CLLocationManagerDelegate", "locationManager:didUpdateToLocation:fromLocation:")
					}
				}()
				manager := CLLocationManagerFromID(managerID)
				newLocation := CLLocationFromID(newLocationID)
				oldLocation := CLLocationFromID(oldLocationID)
				fn(manager, newLocation, oldLocation)
				_delegateDone = true
			},
		})
	}

	if config.LocationManagerDidFinishDeferredUpdatesWithError != nil {
		fn := config.LocationManagerDidFinishDeferredUpdatesWithError
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("locationManager:didFinishDeferredUpdatesWithError:"),
			Fn: func(self objc.ID, _cmd objc.SEL, managerID objc.ID, error_ID objc.ID) {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("CLLocationManagerDelegate", "locationManager:didFinishDeferredUpdatesWithError:")
					}
				}()
				manager := CLLocationManagerFromID(managerID)
				error_ := foundation.NSErrorFromID(error_ID)
				fn(manager, error_)
				_delegateDone = true
			},
		})
	}

	if config.LocationManagerDidPauseLocationUpdates != nil {
		fn := config.LocationManagerDidPauseLocationUpdates
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("locationManagerDidPauseLocationUpdates:"),
			Fn: func(self objc.ID, _cmd objc.SEL, managerID objc.ID) {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("CLLocationManagerDelegate", "locationManagerDidPauseLocationUpdates:")
					}
				}()
				manager := CLLocationManagerFromID(managerID)
				fn(manager)
				_delegateDone = true
			},
		})
	}

	if config.LocationManagerDidResumeLocationUpdates != nil {
		fn := config.LocationManagerDidResumeLocationUpdates
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("locationManagerDidResumeLocationUpdates:"),
			Fn: func(self objc.ID, _cmd objc.SEL, managerID objc.ID) {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("CLLocationManagerDelegate", "locationManagerDidResumeLocationUpdates:")
					}
				}()
				manager := CLLocationManagerFromID(managerID)
				fn(manager)
				_delegateDone = true
			},
		})
	}

	if config.LocationManagerDidVisit != nil {
		fn := config.LocationManagerDidVisit
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("locationManager:didVisit:"),
			Fn: func(self objc.ID, _cmd objc.SEL, managerID objc.ID, visitID objc.ID) {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("CLLocationManagerDelegate", "locationManager:didVisit:")
					}
				}()
				manager := CLLocationManagerFromID(managerID)
				visit := CLVisitFromID(visitID)
				fn(manager, visit)
				_delegateDone = true
			},
		})
	}

	if config.LocationManagerDidUpdateHeading != nil {
		fn := config.LocationManagerDidUpdateHeading
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("locationManager:didUpdateHeading:"),
			Fn: func(self objc.ID, _cmd objc.SEL, managerID objc.ID, newHeadingID objc.ID) {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("CLLocationManagerDelegate", "locationManager:didUpdateHeading:")
					}
				}()
				manager := CLLocationManagerFromID(managerID)
				newHeading := CLHeadingFromID(newHeadingID)
				fn(manager, newHeading)
				_delegateDone = true
			},
		})
	}

	if config.LocationManagerShouldDisplayHeadingCalibration != nil {
		fn := config.LocationManagerShouldDisplayHeadingCalibration
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("locationManagerShouldDisplayHeadingCalibration:"),
			Fn: func(self objc.ID, _cmd objc.SEL, managerID objc.ID) bool {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("CLLocationManagerDelegate", "locationManagerShouldDisplayHeadingCalibration:")
					}
				}()
				manager := CLLocationManagerFromID(managerID)
				_delegateResult := fn(manager)
				_delegateDone = true
				return _delegateResult
			},
		})
	}

	if config.LocationManagerDidEnterRegion != nil {
		fn := config.LocationManagerDidEnterRegion
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("locationManager:didEnterRegion:"),
			Fn: func(self objc.ID, _cmd objc.SEL, managerID objc.ID, regionID objc.ID) {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("CLLocationManagerDelegate", "locationManager:didEnterRegion:")
					}
				}()
				manager := CLLocationManagerFromID(managerID)
				region := CLRegionFromID(regionID)
				fn(manager, region)
				_delegateDone = true
			},
		})
	}

	if config.LocationManagerDidExitRegion != nil {
		fn := config.LocationManagerDidExitRegion
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("locationManager:didExitRegion:"),
			Fn: func(self objc.ID, _cmd objc.SEL, managerID objc.ID, regionID objc.ID) {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("CLLocationManagerDelegate", "locationManager:didExitRegion:")
					}
				}()
				manager := CLLocationManagerFromID(managerID)
				region := CLRegionFromID(regionID)
				fn(manager, region)
				_delegateDone = true
			},
		})
	}

	if config.LocationManagerDidDetermineStateForRegion != nil {
		fn := config.LocationManagerDidDetermineStateForRegion
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("locationManager:didDetermineState:forRegion:"),
			Fn: func(self objc.ID, _cmd objc.SEL, managerID objc.ID, state CLRegionState, regionID objc.ID) {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("CLLocationManagerDelegate", "locationManager:didDetermineState:forRegion:")
					}
				}()
				manager := CLLocationManagerFromID(managerID)
				region := CLRegionFromID(regionID)
				fn(manager, state, region)
				_delegateDone = true
			},
		})
	}

	if config.LocationManagerMonitoringDidFailForRegionWithError != nil {
		fn := config.LocationManagerMonitoringDidFailForRegionWithError
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("locationManager:monitoringDidFailForRegion:withError:"),
			Fn: func(self objc.ID, _cmd objc.SEL, managerID objc.ID, regionID objc.ID, error_ID objc.ID) {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("CLLocationManagerDelegate", "locationManager:monitoringDidFailForRegion:withError:")
					}
				}()
				manager := CLLocationManagerFromID(managerID)
				region := CLRegionFromID(regionID)
				error_ := foundation.NSErrorFromID(error_ID)
				fn(manager, region, error_)
				_delegateDone = true
			},
		})
	}

	if config.LocationManagerDidStartMonitoringForRegion != nil {
		fn := config.LocationManagerDidStartMonitoringForRegion
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("locationManager:didStartMonitoringForRegion:"),
			Fn: func(self objc.ID, _cmd objc.SEL, managerID objc.ID, regionID objc.ID) {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("CLLocationManagerDelegate", "locationManager:didStartMonitoringForRegion:")
					}
				}()
				manager := CLLocationManagerFromID(managerID)
				region := CLRegionFromID(regionID)
				fn(manager, region)
				_delegateDone = true
			},
		})
	}

	if config.LocationManagerDidFailRangingBeaconsForConstraintError != nil {
		fn := config.LocationManagerDidFailRangingBeaconsForConstraintError
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("locationManager:didFailRangingBeaconsForConstraint:error:"),
			Fn: func(self objc.ID, _cmd objc.SEL, managerID objc.ID, beaconConstraintID objc.ID, error_ID objc.ID) {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("CLLocationManagerDelegate", "locationManager:didFailRangingBeaconsForConstraint:error:")
					}
				}()
				manager := CLLocationManagerFromID(managerID)
				beaconConstraint := CLBeaconIdentityConstraintFromID(beaconConstraintID)
				error_ := foundation.NSErrorFromID(error_ID)
				fn(manager, beaconConstraint, error_)
				_delegateDone = true
			},
		})
	}

	nsObjectClass := objc.GetClass("NSObject")
	proto := objc.GetProtocol("CLLocationManagerDelegate")

	var protocols []*objc.Protocol
	if proto != nil {
		protocols = append(protocols, proto)
	}

	cls, err := objc.RegisterClass(className, nsObjectClass, protocols, nil, methods)
	if err != nil {
		panic(fmt.Sprintf("NewCLLocationManagerDelegate: RegisterClass %s: %v", className, err))
	}

	instance := objc.ID(cls).Send(objc.RegisterName("alloc")).Send(objc.RegisterName("init"))
	return CLLocationManagerDelegateObjectFromID(instance)
}
