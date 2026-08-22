// Code generated from Apple documentation for CoreLocation. DO NOT EDIT.

package corelocation

import (
	"context"
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CLLocationManager] class.
var (
	_CLLocationManagerClass     CLLocationManagerClass
	_CLLocationManagerClassOnce sync.Once
)

func getCLLocationManagerClass() CLLocationManagerClass {
	_CLLocationManagerClassOnce.Do(func() {
		_CLLocationManagerClass = CLLocationManagerClass{class: objc.GetClass("CLLocationManager")}
	})
	return _CLLocationManagerClass
}

// GetCLLocationManagerClass returns the class object for CLLocationManager.
func GetCLLocationManagerClass() CLLocationManagerClass {
	return getCLLocationManagerClass()
}

type CLLocationManagerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CLLocationManagerClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CLLocationManagerClass) Alloc() CLLocationManager {
	rv := objc.Send[CLLocationManager](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// The object you use to start and stop the delivery of location-related
// events to your app.
//
// # Overview
//
// A [CLLocationManager] object is the central place to manage your app’s
// location-related behaviors. Use a location-manager object to configure,
// start, and stop location services. You might use these services to:
//
// - Track large or small changes in the user’s current location with a
// configurable degree of accuracy. - Report heading changes from the onboard
// compass. - Monitor geographical regions of interest and generate events
// when someone enters or leaves those regions. - Report the range to nearby
// Bluetooth beacons.
//
// Create one or more location-manager objects in your app and use them where
// you need location data. After you create a location-manager object,
// configure it so that Core Location knows how often to report location
// changes. In particular, configure the [CLLocationManager.DistanceFilter]
// and [CLLocationManager.DesiredAccuracy] properties with values that reflect
// your app’s needs.
//
// A [CLLocationManager] object reports all location-related updates to its
// [CLLocationManager.Delegate] object, which is an object that conforms to
// the [CLLocationManagerDelegate] protocol. Assign the delegate immediately
// when you configure your location manager, because the system reports the
// app’s authorization status to the delegate’s
// [LocationManagerDidChangeAuthorization] method after the location manager
// finishes initializing itself. Core Location calls the methods of your
// delegate object using the [RunLoop] of the thread on which you initialized
// the [CLLocationManager] object. That thread must itself have an active
// [RunLoop], like the one found in your app’s main thread.
//
// For more information, see [Configuring your app to use location services].
//
// # Determining the availability of services
//
//   - [CLLocationManager.IsAuthorizedForWidgetUpdates]: A Boolean value that indicates whether a widget is eligible to receive location updates.
//   - [CLLocationManager.AccuracyAuthorization]: A value that indicates the level of location accuracy the app has permission to use.
//
// # Receiving data from location services
//
//   - [CLLocationManager.Delegate]: The delegate object to receive update events.
//   - [CLLocationManager.SetDelegate]
//
// # Requesting authorization for location services
//
//   - [CLLocationManager.RequestWhenInUseAuthorization]: Requests the user’s permission to use location services while the app is in use.
//   - [CLLocationManager.RequestAlwaysAuthorization]: Requests the user’s permission to use location services regardless of whether the app is in use.
//   - [CLLocationManager.RequestTemporaryFullAccuracyAuthorizationWithPurposeKeyCompletion]: Requests permission to temporarily use location services with full accuracy and reports the results to the provided completion handler.
//   - [CLLocationManager.RequestTemporaryFullAccuracyAuthorizationWithPurposeKey]: Requests permission to temporarily use location services with full accuracy.
//   - [CLLocationManager.AuthorizationStatus]: The current authorization status for the app.
//
// # Specifying distance and accuracy
//
//   - [CLLocationManager.DistanceFilter]: The minimum distance in meters the device must move horizontally before an update event is generated.
//   - [CLLocationManager.SetDistanceFilter]
//   - [CLLocationManager.DesiredAccuracy]: The accuracy of the location data that your app wants to receive.
//   - [CLLocationManager.SetDesiredAccuracy]
//
// # Running the standard location service
//
//   - [CLLocationManager.StartUpdatingLocation]: Starts the generation of updates that report the user’s current location.
//   - [CLLocationManager.StopUpdatingLocation]: Stops the generation of location updates.
//   - [CLLocationManager.RequestLocation]: Requests the one-time delivery of the user’s current location.
//   - [CLLocationManager.PausesLocationUpdatesAutomatically]: A Boolean value that indicates whether the location-manager object may pause location updates.
//   - [CLLocationManager.SetPausesLocationUpdatesAutomatically]
//   - [CLLocationManager.AllowsBackgroundLocationUpdates]: A Boolean value that indicates whether the app receives location updates when running in the background.
//   - [CLLocationManager.SetAllowsBackgroundLocationUpdates]
//   - [CLLocationManager.ActivityType]: The type of activity the app expects the user to typically perform while in the app’s location session.
//   - [CLLocationManager.SetActivityType]
//
// # Running the significant change location service
//
//   - [CLLocationManager.StartMonitoringSignificantLocationChanges]: Starts the generation of updates based on significant location changes.
//   - [CLLocationManager.StopMonitoringSignificantLocationChanges]: Stops the delivery of location events based on significant location changes.
//
// # Running the visits location service
//
//   - [CLLocationManager.StartMonitoringVisits]: Starts the delivery of visit-related events.
//   - [CLLocationManager.StopMonitoringVisits]: Stops the delivery of visit-related events.
//
// # Running the heading service
//
//   - [CLLocationManager.StartUpdatingHeading]: Starts the generation of updates that report the user’s current heading.
//   - [CLLocationManager.DismissHeadingCalibrationDisplay]: Dismisses the heading calibration view from the screen immediately.
//   - [CLLocationManager.HeadingFilter]: The minimum angular change in degrees required to generate new heading events.
//   - [CLLocationManager.SetHeadingFilter]
//   - [CLLocationManager.HeadingOrientation]: The device orientation to use when computing heading values.
//   - [CLLocationManager.SetHeadingOrientation]
//
// # Running the region-monitoring service
//
//   - [CLLocationManager.MonitoredRegions]: The set of shared regions monitored by all location-manager objects.
//   - [CLLocationManager.MaximumRegionMonitoringDistance]: The largest boundary distance that can be assigned to a region.
//
// # Performing beacon ranging
//
//   - [CLLocationManager.StartRangingBeaconsSatisfyingConstraint]: Starts the delivery of notifications for the specified beacon constraints.
//   - [CLLocationManager.StopRangingBeaconsSatisfyingConstraint]: Stops the delivery of notifications for the specified beacon constraints.
//   - [CLLocationManager.RangedBeaconConstraints]: The set of beacon constraints currently being tracked using ranging.
//
// # Getting recent location and heading data
//
//   - [CLLocationManager.Location]: The most recently retrieved user location.
//   - [CLLocationManager.Heading]: The most recently reported heading.
//
// # Instance Properties
//
//   - [CLLocationManager.HeadingBody]
//   - [CLLocationManager.SetHeadingBody]
//
// See: https://developer.apple.com/documentation/CoreLocation/CLLocationManager
//
// [Configuring your app to use location services]: https://developer.apple.com/documentation/CoreLocation/configuring-your-app-to-use-location-services
// [RunLoop]: https://developer.apple.com/documentation/Foundation/RunLoop
type CLLocationManager struct {
	objectivec.Object
}

// CLLocationManagerFromID constructs a [CLLocationManager] from an objc.ID.
//
// The object you use to start and stop the delivery of location-related
// events to your app.
func CLLocationManagerFromID(id objc.ID) CLLocationManager {
	return CLLocationManager{objectivec.Object{ID: id}}
}

// NOTE: CLLocationManager adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CLLocationManager] class.
//
// # Determining the availability of services
//
//   - [ICLLocationManager.IsAuthorizedForWidgetUpdates]: A Boolean value that indicates whether a widget is eligible to receive location updates.
//   - [ICLLocationManager.AccuracyAuthorization]: A value that indicates the level of location accuracy the app has permission to use.
//
// # Receiving data from location services
//
//   - [ICLLocationManager.Delegate]: The delegate object to receive update events.
//   - [ICLLocationManager.SetDelegate]
//
// # Requesting authorization for location services
//
//   - [ICLLocationManager.RequestWhenInUseAuthorization]: Requests the user’s permission to use location services while the app is in use.
//   - [ICLLocationManager.RequestAlwaysAuthorization]: Requests the user’s permission to use location services regardless of whether the app is in use.
//   - [ICLLocationManager.RequestTemporaryFullAccuracyAuthorizationWithPurposeKeyCompletion]: Requests permission to temporarily use location services with full accuracy and reports the results to the provided completion handler.
//   - [ICLLocationManager.RequestTemporaryFullAccuracyAuthorizationWithPurposeKey]: Requests permission to temporarily use location services with full accuracy.
//   - [ICLLocationManager.AuthorizationStatus]: The current authorization status for the app.
//
// # Specifying distance and accuracy
//
//   - [ICLLocationManager.DistanceFilter]: The minimum distance in meters the device must move horizontally before an update event is generated.
//   - [ICLLocationManager.SetDistanceFilter]
//   - [ICLLocationManager.DesiredAccuracy]: The accuracy of the location data that your app wants to receive.
//   - [ICLLocationManager.SetDesiredAccuracy]
//
// # Running the standard location service
//
//   - [ICLLocationManager.StartUpdatingLocation]: Starts the generation of updates that report the user’s current location.
//   - [ICLLocationManager.StopUpdatingLocation]: Stops the generation of location updates.
//   - [ICLLocationManager.RequestLocation]: Requests the one-time delivery of the user’s current location.
//   - [ICLLocationManager.PausesLocationUpdatesAutomatically]: A Boolean value that indicates whether the location-manager object may pause location updates.
//   - [ICLLocationManager.SetPausesLocationUpdatesAutomatically]
//   - [ICLLocationManager.AllowsBackgroundLocationUpdates]: A Boolean value that indicates whether the app receives location updates when running in the background.
//   - [ICLLocationManager.SetAllowsBackgroundLocationUpdates]
//   - [ICLLocationManager.ActivityType]: The type of activity the app expects the user to typically perform while in the app’s location session.
//   - [ICLLocationManager.SetActivityType]
//
// # Running the significant change location service
//
//   - [ICLLocationManager.StartMonitoringSignificantLocationChanges]: Starts the generation of updates based on significant location changes.
//   - [ICLLocationManager.StopMonitoringSignificantLocationChanges]: Stops the delivery of location events based on significant location changes.
//
// # Running the visits location service
//
//   - [ICLLocationManager.StartMonitoringVisits]: Starts the delivery of visit-related events.
//   - [ICLLocationManager.StopMonitoringVisits]: Stops the delivery of visit-related events.
//
// # Running the heading service
//
//   - [ICLLocationManager.StartUpdatingHeading]: Starts the generation of updates that report the user’s current heading.
//   - [ICLLocationManager.DismissHeadingCalibrationDisplay]: Dismisses the heading calibration view from the screen immediately.
//   - [ICLLocationManager.HeadingFilter]: The minimum angular change in degrees required to generate new heading events.
//   - [ICLLocationManager.SetHeadingFilter]
//   - [ICLLocationManager.HeadingOrientation]: The device orientation to use when computing heading values.
//   - [ICLLocationManager.SetHeadingOrientation]
//
// # Running the region-monitoring service
//
//   - [ICLLocationManager.MonitoredRegions]: The set of shared regions monitored by all location-manager objects.
//   - [ICLLocationManager.MaximumRegionMonitoringDistance]: The largest boundary distance that can be assigned to a region.
//
// # Performing beacon ranging
//
//   - [ICLLocationManager.StartRangingBeaconsSatisfyingConstraint]: Starts the delivery of notifications for the specified beacon constraints.
//   - [ICLLocationManager.StopRangingBeaconsSatisfyingConstraint]: Stops the delivery of notifications for the specified beacon constraints.
//   - [ICLLocationManager.RangedBeaconConstraints]: The set of beacon constraints currently being tracked using ranging.
//
// # Getting recent location and heading data
//
//   - [ICLLocationManager.Location]: The most recently retrieved user location.
//   - [ICLLocationManager.Heading]: The most recently reported heading.
//
// # Instance Properties
//
//   - [ICLLocationManager.HeadingBody]
//   - [ICLLocationManager.SetHeadingBody]
//
// See: https://developer.apple.com/documentation/CoreLocation/CLLocationManager
type ICLLocationManager interface {
	objectivec.IObject

	// Topic: Determining the availability of services

	// A Boolean value that indicates whether a widget is eligible to receive location updates.
	IsAuthorizedForWidgetUpdates() bool
	// A value that indicates the level of location accuracy the app has permission to use.
	AccuracyAuthorization() CLAccuracyAuthorization

	// Topic: Receiving data from location services

	// The delegate object to receive update events.
	Delegate() CLLocationManagerDelegate
	SetDelegate(value CLLocationManagerDelegate)

	// Topic: Requesting authorization for location services

	// Requests the user’s permission to use location services while the app is in use.
	RequestWhenInUseAuthorization()
	// Requests the user’s permission to use location services regardless of whether the app is in use.
	RequestAlwaysAuthorization()
	// Requests permission to temporarily use location services with full accuracy and reports the results to the provided completion handler.
	RequestTemporaryFullAccuracyAuthorizationWithPurposeKeyCompletion(purposeKey string, completion ErrorHandler)
	// Requests permission to temporarily use location services with full accuracy.
	RequestTemporaryFullAccuracyAuthorizationWithPurposeKey(purposeKey string)
	// The current authorization status for the app.
	AuthorizationStatus() CLAuthorizationStatus

	// Topic: Specifying distance and accuracy

	// The minimum distance in meters the device must move horizontally before an update event is generated.
	DistanceFilter() CLLocationDistance
	SetDistanceFilter(value CLLocationDistance)
	// The accuracy of the location data that your app wants to receive.
	DesiredAccuracy() CLLocationAccuracy
	SetDesiredAccuracy(value CLLocationAccuracy)

	// Topic: Running the standard location service

	// Starts the generation of updates that report the user’s current location.
	StartUpdatingLocation()
	// Stops the generation of location updates.
	StopUpdatingLocation()
	// Requests the one-time delivery of the user’s current location.
	RequestLocation()
	// A Boolean value that indicates whether the location-manager object may pause location updates.
	PausesLocationUpdatesAutomatically() bool
	SetPausesLocationUpdatesAutomatically(value bool)
	// A Boolean value that indicates whether the app receives location updates when running in the background.
	AllowsBackgroundLocationUpdates() bool
	SetAllowsBackgroundLocationUpdates(value bool)
	// The type of activity the app expects the user to typically perform while in the app’s location session.
	ActivityType() CLActivityType
	SetActivityType(value CLActivityType)

	// Topic: Running the significant change location service

	// Starts the generation of updates based on significant location changes.
	StartMonitoringSignificantLocationChanges()
	// Stops the delivery of location events based on significant location changes.
	StopMonitoringSignificantLocationChanges()

	// Topic: Running the visits location service

	// Starts the delivery of visit-related events.
	StartMonitoringVisits()
	// Stops the delivery of visit-related events.
	StopMonitoringVisits()

	// Topic: Running the heading service

	// Starts the generation of updates that report the user’s current heading.
	StartUpdatingHeading()
	// Dismisses the heading calibration view from the screen immediately.
	DismissHeadingCalibrationDisplay()
	// The minimum angular change in degrees required to generate new heading events.
	HeadingFilter() CLLocationDegrees
	SetHeadingFilter(value CLLocationDegrees)
	// The device orientation to use when computing heading values.
	HeadingOrientation() CLDeviceOrientation
	SetHeadingOrientation(value CLDeviceOrientation)

	// Topic: Running the region-monitoring service

	// The set of shared regions monitored by all location-manager objects.
	MonitoredRegions() foundation.INSSet
	// The largest boundary distance that can be assigned to a region.
	MaximumRegionMonitoringDistance() CLLocationDistance

	// Topic: Performing beacon ranging

	// Starts the delivery of notifications for the specified beacon constraints.
	StartRangingBeaconsSatisfyingConstraint(constraint ICLBeaconIdentityConstraint)
	// Stops the delivery of notifications for the specified beacon constraints.
	StopRangingBeaconsSatisfyingConstraint(constraint ICLBeaconIdentityConstraint)
	// The set of beacon constraints currently being tracked using ranging.
	RangedBeaconConstraints() foundation.INSSet

	// Topic: Getting recent location and heading data

	// The most recently retrieved user location.
	Location() ICLLocation
	// The most recently reported heading.
	Heading() ICLHeading

	// Topic: Instance Properties

	HeadingBody() CLBodyIdentifiable
	SetHeadingBody(value CLBodyIdentifiable)
}

// Init initializes the instance.
func (l CLLocationManager) Init() CLLocationManager {
	rv := objc.Send[CLLocationManager](l.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (l CLLocationManager) Autorelease() CLLocationManager {
	rv := objc.Send[CLLocationManager](l.ID, objc.Sel("autorelease"))
	return rv
}

// NewCLLocationManager creates a new CLLocationManager instance.
func NewCLLocationManager() CLLocationManager {
	class := getCLLocationManagerClass()
	rv := objc.Send[CLLocationManager](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Requests the user’s permission to use location services while the app is
// in use.
//
// # Discussion
//
// You must call this method or [CLLocationManager.RequestAlwaysAuthorization]
// before you can receive location-related information. You may call
// [CLLocationManager.RequestWhenInUseAuthorization] whenever the current
// authorization status is not determined
// ([KCLAuthorizationStatusNotDetermined]).
//
// This method runs asynchronously and prompts the user to grant permission to
// the app to use location services. The user prompt contains the text from
// the [NSLocationWhenInUseUsageDescription] key in your app `Info.Plist()`
// file, and the presence of that key is required when calling this method.
// The user prompt displays the following options, which determine the
// authorization your app can receive.
//
// [Table data omitted]
//
// After the user makes a selection and determines the status, the location
// manager delivers the results to the delegate’s
// [locationManager(_:didChangeAuthorization:)] method. If the initial
// authorization status is anything other than
// [KCLAuthorizationStatusNotDetermined], this method does nothing and
// doesn’t call the [locationManager(_:didChangeAuthorization:)] method.
//
// If the user’s choice grants When In Use authorization to your app, your
// app can start any location service and is eligible to receive the results
// while it’s in use. If the user’s choice grants temporary When In Use
// authorization, the authorization expires when the app is no longer in use,
// reverting to Not Determined status ([KCLAuthorizationStatusNotDetermined]).
// For information about when an app is considered to be in use, see [Choosing
// the Location Services Authorization to Request].
//
// When your app starts standard location services in the foreground, they
// continue to run in the background if your app has enabled background
// location updates in the Capabilities tab of your Xcode project. Attempts to
// start location updates while your app runs in the background will fail. The
// system displays a location services indicator in the status bar when your
// app moves to the background with active location services.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLLocationManager/requestWhenInUseAuthorization()
//
// [NSLocationWhenInUseUsageDescription]: https://developer.apple.com/documentation/BundleResources/Information-Property-List/NSLocationWhenInUseUsageDescription
// [locationManager(_:didChangeAuthorization:)]: https://developer.apple.com/documentation/CoreLocation/CLLocationManagerDelegate/locationManager(_:didChangeAuthorization:)
//
// [Choosing the  Location Services Authorization to Request]: https://developer.apple.com/documentation/BundleResources/choosing-the-location-services-authorization-to-request
func (l CLLocationManager) RequestWhenInUseAuthorization() {
	objc.Send[objc.ID](l.ID, objc.Sel("requestWhenInUseAuthorization"))
}

// Requests the user’s permission to use location services regardless of
// whether the app is in use.
//
// # Discussion
//
// You must call this or the [CLLocationManager.RequestWhenInUseAuthorization]
// method before your app can receive location information. To call this
// method, you must have both [NSLocationAlwaysUsageDescription] and
// [NSLocationWhenInUseUsageDescription] keys in your app’s `Info.Plist()`
// file. You may call [CLLocationManager.RequestAlwaysAuthorization] when the
// current authorization state is either:
//
// - Not Determined — [KCLAuthorizationStatusNotDetermined] - When In Use
// — [KCLAuthorizationStatusAuthorizedWhenInUse]
//
// Use the [LocationManagerDidUpdateLocations] method on the
// [CLLocationManager] delegate to receive updates when the user makes
// permission choices.
//
// Core Location limits calls to
// [CLLocationManager.RequestAlwaysAuthorization]. After your app calls this
// method, further calls have no effect. If a compatible iPad or iPhone app
// calls this method when running in visionOS, the method treats it as a
// request for When in Use authorization instead.
//
// # Request Always Authorization After Getting When In Use
//
// To obtain Always authorization, your app must first request When In Use
// permission followed by requesting Always authorization.
//
// If the user grants When In Use permission after your app calls
// [CLLocationManager.RequestWhenInUseAuthorization], then calling
// [CLLocationManager.RequestAlwaysAuthorization] immediately prompts the user
// to request Always permission. If the user responded to
// [CLLocationManager.RequestWhenInUseAuthorization] with Allow Once, then
// Core Location ignores further calls to
// [CLLocationManager.RequestAlwaysAuthorization] due to the temporary
// authorization.
//
// Core Location prompts the user to grant permission with the string from
// [NSLocationAlwaysUsageDescription]. The user prompt displays the following
// options, which determine the authorization your app can receive:
//
// [Table data omitted]
//
// # Request Always Authorization Directly
//
// If your app’s current state is [KCLAuthorizationStatusNotDetermined] and
// you call [CLLocationManager.RequestAlwaysAuthorization], Core Location uses
// two prompts before it fully enables Always authorization.
//
// The first prompt displays immediately with the string from
// [NSLocationWhenInUseUsageDescription]. The user prompt displays the
// following options, which determine the authorization your app receives:
//
// [Table data omitted]
//
// The second prompt displays when Core Location prepares to deliver an event
// to your app requiring [KCLAuthorizationStatusAuthorizedAlways]. If the app
// is in the Provisional Always state, the system displays the second prompt
// with the string from [NSLocationAlwaysUsageDescription]. Core Location will
// typically display the second prompt when your app isn’t running.
//
// Your app receives permanent Always authorization if the user chooses to
// grant permission when the second prompt appears while in the Provisional
// Always state. When the user responds, your app receives either the location
// event or a call to your delegate with the modified authorization.
//
// When displaying the second prompt, the user sees one of the following
// options:
//
// [Table data omitted]
//
// If the user responds to the prompt near the time it was delivered and
// chooses to allow the Always permission, the location event will be
// delivered to your app.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLLocationManager/requestAlwaysAuthorization()
//
// [NSLocationAlwaysUsageDescription]: https://developer.apple.com/documentation/BundleResources/Information-Property-List/NSLocationAlwaysUsageDescription
// [NSLocationWhenInUseUsageDescription]: https://developer.apple.com/documentation/BundleResources/Information-Property-List/NSLocationWhenInUseUsageDescription
func (l CLLocationManager) RequestAlwaysAuthorization() {
	objc.Send[objc.ID](l.ID, objc.Sel("requestAlwaysAuthorization"))
}

// Requests permission to temporarily use location services with full accuracy
// and reports the results to the provided completion handler.
//
// purposeKey: A key in the [NSLocationTemporaryUsageDescriptionDictionary] dictionary of
// the app’s `Info.Plist()` file. The value for this key is an app-provided
// string that describes the reason for accessing location data with full
// accuracy. To localize a usage description, add an entry to your
// `InfoPlist.Strings()` file with the same key you provide for this
// parameter.
//
// completion: A closure to execute after authorization status changes. This closure takes
// a single `error` parameter, which is `nil` if the prompt was displayed to
// the user, or an error object describing why the prompt couldn’t be
// displayed.
//
// # Discussion
//
// After the user gives permission for your app to use location data with full
// accuracy, your app can access that data in the foreground or in the
// background, until its permission automatically expires. Expiration is
// postponed while your app is actively in use. For example, expiration is
// postponed while your app in the foreground, and while a Continuous
// Background Location session is active with the background location
// indicator enabled. This approach to expiration allows apps to provide
// experiences that require full accuracy, such as fitness and navigation
// apps, even if the user doesn’t grant persistent access for full accuracy.
//
// The completion closure is guaranteed to be called after the request is
// completed, which includes the user granting access, the user declining, or
// an error that prevented displaying the prompt. The closure is always called
// in the same threading context as [CLLocationManagerDelegate] methods. If
// the prompt was successfully displayed to the user, the callback’s `error`
// parameter is `nil`.
//
// The request always fails with a [KCLErrorPromptDeclined] error in the
// following cases:
//
// - The `Info.Plist()` file doesn’t have an entry for the given
// `purposeKey` value. - The app is already authorized for full accuracy. -
// The app is in the background.
//
// If the closure is called with an error, log the error for debugging
// purposes, and retry the request again the next time the user performs the
// action that caused you to request precise location information.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLLocationManager/requestTemporaryFullAccuracyAuthorization(withPurposeKey:completion:)
//
// [NSLocationTemporaryUsageDescriptionDictionary]: https://developer.apple.com/documentation/BundleResources/Information-Property-List/NSLocationTemporaryUsageDescriptionDictionary
func (l CLLocationManager) RequestTemporaryFullAccuracyAuthorizationWithPurposeKeyCompletion(purposeKey string, completion ErrorHandler) {
	_block1, _ := NewErrorBlock(completion)
	objc.Send[objc.ID](l.ID, objc.Sel("requestTemporaryFullAccuracyAuthorizationWithPurposeKey:completion:"), objc.String(purposeKey), _block1)
}

// Requests permission to temporarily use location services with full
// accuracy.
//
// purposeKey: A key in the [NSLocationTemporaryUsageDescriptionDictionary] dictionary of
// the app’s `Info.Plist()` file. The value for this key is an app-provided
// string that describes the reason for accessing location data with full
// accuracy. To localize a usage description, add an entry to your
// `InfoPlist.Strings()` file with the same key you provide for this
// parameter.
//
// # Discussion
//
// This method behaves the same as calling the
// [CLLocationManager.RequestTemporaryFullAccuracyAuthorizationWithPurposeKeyCompletion]
// method, passing `nil` as the completion closure. Use this method if your
// app’s logic to respond to changes in location data accuracy is already
// handled by the [LocationManagerDidChangeAuthorization] delegate method, and
// your app doesn’t have any work to do in the closure.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLLocationManager/requestTemporaryFullAccuracyAuthorization(withPurposeKey:)
//
// [NSLocationTemporaryUsageDescriptionDictionary]: https://developer.apple.com/documentation/BundleResources/Information-Property-List/NSLocationTemporaryUsageDescriptionDictionary
func (l CLLocationManager) RequestTemporaryFullAccuracyAuthorizationWithPurposeKey(purposeKey string) {
	objc.Send[objc.ID](l.ID, objc.Sel("requestTemporaryFullAccuracyAuthorizationWithPurposeKey:"), objc.String(purposeKey))
}

// Starts the generation of updates that report the user’s current location.
//
// # Discussion
//
// This method returns immediately. Calling this method causes the location
// manager to obtain an initial location fix (which may take several seconds)
// and notify your delegate by calling its [LocationManagerDidUpdateLocations]
// method. After that, the receiver generates update events primarily when the
// value in the [CLLocationManager.DistanceFilter] property is exceeded.
// Updates may be delivered in other situations though. For example, the
// receiver may send another notification if the hardware gathers a more
// accurate location reading.
//
// Calling this method several times in succession does not automatically
// result in new events being generated. Calling
// [CLLocationManager.StopUpdatingLocation] in between, however, does cause a
// new initial event to be sent the next time you call this method.
//
// If you start this service and your app is suspended, the system stops the
// delivery of events until your app starts running again (either in the
// foreground or background). If your app is terminated, the delivery of new
// location events stops altogether. Therefore, if your app needs to receive
// location events while in the background, it must include the
// [UIBackgroundModes] key (with the `location` value) in its `Info.Plist()`
// file.
//
// In addition to your delegate object implementing the
// [LocationManagerDidUpdateLocations] method, it should also implement the
// [LocationManagerDidFailWithError] method to respond to potential errors.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLLocationManager/startUpdatingLocation()
func (l CLLocationManager) StartUpdatingLocation() {
	objc.Send[objc.ID](l.ID, objc.Sel("startUpdatingLocation"))
}

// Stops the generation of location updates.
//
// # Discussion
//
// Call this method whenever your code no longer needs to receive
// location-related events. Disabling event delivery gives the receiver the
// option of disabling the appropriate hardware (and thereby saving power)
// when no clients need location data. You can always restart the generation
// of location updates by calling the
// [CLLocationManager.StartUpdatingLocation] method again.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLLocationManager/stopUpdatingLocation()
func (l CLLocationManager) StopUpdatingLocation() {
	objc.Send[objc.ID](l.ID, objc.Sel("stopUpdatingLocation"))
}

// Requests the one-time delivery of the user’s current location.
//
// # Discussion
//
// This method returns immediately. Calling it causes the location manager to
// obtain a location fix (which may take several seconds) and call the
// delegate’s [LocationManagerDidUpdateLocations] method with the result.
// The location fix is obtained at the accuracy level indicated by the
// [CLLocationManager.DesiredAccuracy] property. Only one location fix is
// reported to the delegate, after which location services are stopped. If a
// location fix cannot be determined in a timely manner, the location manager
// calls the delegate’s [LocationManagerDidFailWithError] method instead and
// reports a [KCLErrorLocationUnknown] error.
//
// Use this method when you want the user’s current location but do not need
// to leave location services running. This method starts location services
// long enough to return a result or report an error and then stops them
// again. Calling the [CLLocationManager.StartUpdatingLocation] or
// [allowDeferredLocationUpdates(untilTraveled:timeout:)] method cancels any
// pending request made using this method. Calling this method while location
// services are already running does nothing. To cancel a pending request,
// call the [CLLocationManager.StopUpdatingLocation] method.
//
// If obtaining the desired accuracy would take too long, the location manager
// delivers a less accurate location value rather than reporting an error.
//
// When using this method, the associated delegate must implement the
// [LocationManagerDidUpdateLocations] and [LocationManagerDidFailWithError]
// methods. Failure to do so is a programmer error.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLLocationManager/requestLocation()
//
// [allowDeferredLocationUpdates(untilTraveled:timeout:)]: https://developer.apple.com/documentation/CoreLocation/CLLocationManager/allowDeferredLocationUpdates(untilTraveled:timeout:)
func (l CLLocationManager) RequestLocation() {
	objc.Send[objc.ID](l.ID, objc.Sel("requestLocation"))
}

// Starts the generation of updates based on significant location changes.
//
// # Discussion
//
// This method initiates the delivery of location events asynchronously,
// returning shortly after you call it. Location events are delivered to your
// delegate’s [LocationManagerDidUpdateLocations] method. The first event to
// be delivered is usually the most recently cached location event (if any)
// but may be a newer event in some circumstances. Obtaining a current
// location fix may take several additional seconds, so be sure to check the
// timestamps on the location events in your delegate method.
//
// After returning a current location fix, the receiver generates update
// events only when a significant change in the user’s location is detected.
// It does not rely on the value in the [CLLocationManager.DistanceFilter]
// property to generate events. Calling this method several times in
// succession does not automatically result in new events being generated.
// Calling [CLLocationManager.StopMonitoringSignificantLocationChanges] in
// between, however, does cause a new initial event to be sent the next time
// you call this method.
//
// If you start this service and your app is subsequently terminated, the
// system automatically relaunches the app into the background if a new event
// arrives. In such a case, the options dictionary passed to the
// [application(_:willFinishLaunchingWithOptions:)] and
// [application(_:didFinishLaunchingWithOptions:)] methods of your app
// delegate contains the key [location] to indicate that your app was launched
// because of a location event. Upon relaunch, you must still configure a
// location manager object and call this method to continue receiving location
// events. When you restart location services, the current event is delivered
// to your delegate immediately. In addition, the [CLLocationManager.Location]
// property of your location manager object is populated with the most recent
// location object even before you start location services.
//
// In addition to your delegate object implementing the
// [LocationManagerDidUpdateLocations] method, it should also implement the
// [LocationManagerDidFailWithError] method to respond to potential errors.
//
// If a compatible iPad or iPhone app calls this method when running in
// visionOS, the method does nothing.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLLocationManager/startMonitoringSignificantLocationChanges()
//
// [application(_:didFinishLaunchingWithOptions:)]: https://developer.apple.com/documentation/UIKit/UIApplicationDelegate/application(_:didFinishLaunchingWithOptions:)
// [application(_:willFinishLaunchingWithOptions:)]: https://developer.apple.com/documentation/UIKit/UIApplicationDelegate/application(_:willFinishLaunchingWithOptions:)
// [location]: https://developer.apple.com/documentation/UIKit/UIApplication/LaunchOptionsKey/location
func (l CLLocationManager) StartMonitoringSignificantLocationChanges() {
	objc.Send[objc.ID](l.ID, objc.Sel("startMonitoringSignificantLocationChanges"))
}

// Stops the delivery of location events based on significant location
// changes.
//
// # Discussion
//
// Use this method to stop the delivery of location events that was started
// using the [CLLocationManager.StartMonitoringSignificantLocationChanges]
// method. If a compatible iPad or iPhone app calls this method when running
// in visionOS, the method does nothing.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLLocationManager/stopMonitoringSignificantLocationChanges()
func (l CLLocationManager) StopMonitoringSignificantLocationChanges() {
	objc.Send[objc.ID](l.ID, objc.Sel("stopMonitoringSignificantLocationChanges"))
}

// Starts the delivery of visit-related events.
//
// # Discussion
//
// Calling this method begins the delivery of visit-related events to your
// app. Enabling visit events for one location manager enables visit events
// for all other location manager objects in your app. When a new visit event
// arrives, the location manager object delivers the event to the
// [LocationManagerDidVisit] method of its delegate.
//
// Your app can monitor for visit events without calling
// `requestTemporaryPreciseLocationAuthorization()`. In that case, the visit
// events use reduced accuracy, as reflected by the
// [CLVisit.HorizontalAccuracy] property of [CLVisit].
//
// If your app is terminated while this service is active, the system
// relaunches your app when new visit events are ready to be delivered. Upon
// relaunch, recreate your location manager object and assign a delegate to
// begin receiving visit events. You don’t need to call this method again to
// restart the delivery of visit events, but calling it does no harm.
//
// If a compatible iPad or iPhone app calls this method when running in
// visionOS, the method does nothing.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLLocationManager/startMonitoringVisits()
func (l CLLocationManager) StartMonitoringVisits() {
	objc.Send[objc.ID](l.ID, objc.Sel("startMonitoringVisits"))
}

// Stops the delivery of visit-related events.
//
// # Discussion
//
// Calling this method disables the delivery of visit-related events for your
// app. If a compatible iPad or iPhone app calls this method when running in
// visionOS, the method does nothing.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLLocationManager/stopMonitoringVisits()
func (l CLLocationManager) StopMonitoringVisits() {
	objc.Send[objc.ID](l.ID, objc.Sel("stopMonitoringVisits"))
}

// Starts the generation of updates that report the user’s current heading.
//
// # Discussion
//
// This method returns immediately. Calling this method when the receiver is
// stopped causes it to obtain an initial heading and notify your delegate.
// After that, the receiver generates update events when the value in the
// [CLLocationManager.HeadingFilter] property is exceeded.
//
// Before calling this method, you should always check the [headingAvailable]
// property to see whether heading information is supported on the current
// device. If heading information is not supported, calling this method has no
// effect and does not result in the delivery of events to your delegate.
//
// Calling this method several times in succession does not automatically
// result in new events being generated. Calling
// [CLLocationManager.StopUpdatingHeading] in between, however, does cause a
// new initial event to be sent the next time you call this method.
//
// If you start this service and your app is suspended, the system stops the
// delivery of events until your app starts running again (either in the
// foreground or background). If your app is terminated, the delivery of new
// heading events stops altogether and must be restarted by your code when the
// app is relaunched.
//
// Heading events are delivered to the [LocationManagerDidUpdateHeading]
// method of your delegate. If there is an error, the location manager calls
// the [LocationManagerDidFailWithError] method of your delegate instead.
//
// If a compatible iPad or iPhone app calls this method when running in
// visionOS, the method does nothing.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLLocationManager/startUpdatingHeading()
//
// [headingAvailable]: https://developer.apple.com/documentation/CoreLocation/CLLocationManager/headingAvailable-swift.property
func (l CLLocationManager) StartUpdatingHeading() {
	objc.Send[objc.ID](l.ID, objc.Sel("startUpdatingHeading"))
}

// Dismisses the heading calibration view from the screen immediately.
//
// # Discussion
//
// Core Location uses the heading calibration alert to calibrate the available
// heading hardware as needed. The display of this view is automatic, assuming
// your delegate supports displaying the view at all. If the view is
// displayed, you can use this method to dismiss it after an appropriate
// amount of time to ensure that your app’s user interface is not unduly
// disrupted.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLLocationManager/dismissHeadingCalibrationDisplay()
func (l CLLocationManager) DismissHeadingCalibrationDisplay() {
	objc.Send[objc.ID](l.ID, objc.Sel("dismissHeadingCalibrationDisplay"))
}

// Starts the delivery of notifications for the specified beacon constraints.
//
// constraint: A [CLBeaconIdentityConstraint] constraint.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLLocationManager/startRangingBeacons(satisfying:)
func (l CLLocationManager) StartRangingBeaconsSatisfyingConstraint(constraint ICLBeaconIdentityConstraint) {
	objc.Send[objc.ID](l.ID, objc.Sel("startRangingBeaconsSatisfyingConstraint:"), constraint)
}

// Stops the delivery of notifications for the specified beacon constraints.
//
// constraint: A [CLBeaconIdentityConstraint] constraint.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLLocationManager/stopRangingBeacons(satisfying:)
func (l CLLocationManager) StopRangingBeaconsSatisfyingConstraint(constraint ICLBeaconIdentityConstraint) {
	objc.Send[objc.ID](l.ID, objc.Sel("stopRangingBeaconsSatisfyingConstraint:"), constraint)
}

// Returns a Boolean value indicating whether the significant-change location
// service is available on the device.
//
// # Return Value
//
// true if location change monitoring is available; false if it is not.
//
// # Discussion
//
// This method indicates whether the device is able to report updates based on
// significant location changes only. This capability provides tremendous
// power savings for apps that want to track a user’s approximate location
// and don’t need highly accurate position information.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLLocationManager/significantLocationChangeMonitoringAvailable()
func (_CLLocationManagerClass CLLocationManagerClass) SignificantLocationChangeMonitoringAvailable() bool {
	rv := objc.Send[bool](objc.ID(_CLLocationManagerClass.class), objc.Sel("significantLocationChangeMonitoringAvailable"))
	return rv
}

// Returns a Boolean value indicating whether the location manager is able to
// generate heading-related events.
//
// # Return Value
//
// true if heading data is available; false if it is not.
//
// # Discussion
//
// Heading data may not be available on all iOS-based devices. You should
// check the value returned by this method before asking the location manager
// to deliver heading-related events.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLLocationManager/headingAvailable()
func (_CLLocationManagerClass CLLocationManagerClass) HeadingAvailable() bool {
	rv := objc.Send[bool](objc.ID(_CLLocationManagerClass.class), objc.Sel("headingAvailable"))
	return rv
}

// Returns a Boolean value indicating whether the device supports region
// monitoring using the specified class.
//
// regionClass: A region monitoring class from the MapKit framework. This class must
// descend from the [CLRegion] class.
//
// # Return Value
//
// true if the device is capable of monitoring regions using the specified
// class or false if it is not.
//
// # Discussion
//
// The availability of region monitoring support is dependent on the hardware
// present on the device. This method does not take into account the
// availability of location services or the fact that the user might have
// disabled them for the app or system; you must determine your app’s
// authorization status separately.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLLocationManager/isMonitoringAvailable(for:)
func (_CLLocationManagerClass CLLocationManagerClass) IsMonitoringAvailableForClass(regionClass objectivec.Class) bool {
	rv := objc.Send[bool](objc.ID(_CLLocationManagerClass.class), objc.Sel("isMonitoringAvailableForClass:"), regionClass)
	return rv
}

// Returns a Boolean value indicating whether the device supports ranging of
// beacons that use the iBeacon protocol.
//
// # Return Value
//
// true if the device supports ranging or false if it does not.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLLocationManager/isRangingAvailable()
func (_CLLocationManagerClass CLLocationManagerClass) IsRangingAvailable() bool {
	rv := objc.Send[bool](objc.ID(_CLLocationManagerClass.class), objc.Sel("isRangingAvailable"))
	return rv
}

// Returns a Boolean value indicating whether location services are enabled on
// the device.
//
// # Return Value
//
// true if location services are enabled on the device; false if they are not.
//
// # Discussion
//
// Users can enable or disable location services by toggling the Location
// Services switch in Settings > Privacy.
//
// - When users disable the switch, the system calls your delegate’s
// [locationManager(_:didChangeAuthorization:)] method with a denied
// authorization status ([KCLAuthorizationStatusDenied]). - When users enable
// the switch, the system returns your app’s authorization to its previous
// state and calls your delegate’s
// [locationManager(_:didChangeAuthorization:)] method.
//
// You are not required to call
// [CLLocationManagerClass.LocationServicesEnabled]. However, If you wish to
// display instructions about enabling location services, you may check the
// return value of this method to find out if the services are disabled for
// the entire device, or just for your app. If the result is `true`, provide
// instructions for enabling services for your app; otherwise, provide
// instructions for enabling the Location Services switch in Settings >
// Privacy.
//
// If users disable or deny location services and you attempt to start
// location updates anyway, the location manager reports an error to its
// delegate. See [LocationManagerDidFailWithError] and
// [LocationManagerMonitoringDidFailForRegionWithError] for more information.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLLocationManager/locationServicesEnabled()
//
// [locationManager(_:didChangeAuthorization:)]: https://developer.apple.com/documentation/CoreLocation/CLLocationManagerDelegate/locationManager(_:didChangeAuthorization:)
func (_CLLocationManagerClass CLLocationManagerClass) LocationServicesEnabled() bool {
	rv := objc.Send[bool](objc.ID(_CLLocationManagerClass.class), objc.Sel("locationServicesEnabled"))
	return rv
}

// A Boolean value that indicates whether a widget is eligible to receive
// location updates.
//
// # Discussion
//
// This property is `true` when either of the following is true:
//
// - The app’s authorization status is
// [KCLAuthorizationStatusAuthorizedAlways]. - The app’s authorization
// status is [KCLAuthorizationStatusAuthorizedWhenInUse] and the user agrees
// to extend the app’s authorization status to widgets.
//
// For details about using location information in widgets with
// [KCLAuthorizationStatusAuthorizedWhenInUse], see [Accessing location
// information in widgets].
//
// See: https://developer.apple.com/documentation/CoreLocation/CLLocationManager/isAuthorizedForWidgetUpdates
//
// [Accessing location information in widgets]: https://developer.apple.com/documentation/WidgetKit/Accessing-Location-Information-in-Widgets
func (l CLLocationManager) IsAuthorizedForWidgetUpdates() bool {
	rv := objc.Send[bool](l.ID, objc.Sel("isAuthorizedForWidgetUpdates"))
	return rv
}

// A value that indicates the level of location accuracy the app has
// permission to use.
//
// # Discussion
//
// If the value of this property is [CLAccuracyAuthorizationFullAccuracy], you
// can set the [CLLocationManager.DesiredAccuracy] property to any value. If
// the value is [CLAccuracyAuthorizationReducedAccuracy], setting
// [CLLocationManager.DesiredAccuracy] to a value other than
// [kCLLocationAccuracyReduced] has no effect on the location information, and
// your app can’t use region monitoring or beacon ranging.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLLocationManager/accuracyAuthorization
//
// [kCLLocationAccuracyReduced]: https://developer.apple.com/documentation/CoreLocation/kCLLocationAccuracyReduced
func (l CLLocationManager) AccuracyAuthorization() CLAccuracyAuthorization {
	rv := objc.Send[CLAccuracyAuthorization](l.ID, objc.Sel("accuracyAuthorization"))
	return CLAccuracyAuthorization(rv)
}

// The delegate object to receive update events.
//
// # Discussion
//
// In iOS, this property is declared as `nonatomic`. In macOS, it is declared
// as `atomic`.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLLocationManager/delegate
func (l CLLocationManager) Delegate() CLLocationManagerDelegate {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("delegate"))
	return CLLocationManagerDelegateObjectFromID(rv)
}
func (l CLLocationManager) SetDelegate(value CLLocationManagerDelegate) {
	objc.Send[struct{}](l.ID, objc.Sel("setDelegate:"), value)
}

// The current authorization status for the app.
//
// # Return Value
//
// A value indicating whether the app is authorized to use location services.
//
// # Discussion
//
// Check this value when the [LocationManagerDidChangeAuthorization] delegate
// callback indicates that the authorization status has changed.
//
// The system is guaranteed to call the delegate method with the app’s
// initial authorization state and all authorization status changes.
//
// The system manages the authorization status of a given app according to
// several factors. Users must authorize the app to use location services
// explicitly, and location services must be enabled in Settings > Privacy.
// See [Choosing the Location Services Authorization to Request] for more
// information.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLLocationManager/authorizationStatus-swift.property
//
// [Choosing the  Location Services Authorization to Request]: https://developer.apple.com/documentation/BundleResources/choosing-the-location-services-authorization-to-request
func (l CLLocationManager) AuthorizationStatus() CLAuthorizationStatus {
	rv := objc.Send[CLAuthorizationStatus](l.ID, objc.Sel("authorizationStatus"))
	return CLAuthorizationStatus(rv)
}

// The minimum distance in meters the device must move horizontally before an
// update event is generated.
//
// # Discussion
//
// This location manager measures this relative to the previously delivered
// location. Specify the value [kCLDistanceFilterNone] to receive
// notifications for all movements. The default value of this property is
// [kCLDistanceFilterNone].
//
// Use this property only in conjunction with the Standard location services
// and not with the Significant-change or Visits services.
//
// # Special Considerations
//
// In iOS, this property is declared as `nonatomic`. In macOS, it is declared
// as `atomic`.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLLocationManager/distanceFilter
//
// [kCLDistanceFilterNone]: https://developer.apple.com/documentation/CoreLocation/kCLDistanceFilterNone
func (l CLLocationManager) DistanceFilter() CLLocationDistance {
	rv := objc.Send[CLLocationDistance](l.ID, objc.Sel("distanceFilter"))
	return CLLocationDistance(rv)
}
func (l CLLocationManager) SetDistanceFilter(value CLLocationDistance) {
	objc.Send[struct{}](l.ID, objc.Sel("setDistanceFilter:"), value)
}

// The accuracy of the location data that your app wants to receive.
//
// # Discussion
//
// The location service does its best to achieve the requested accuracy;
// however, apps must be prepared to use less accurate data. If your app
// isn’t authorized to access precise location information
// (`isAuthorizedForPreciseLocation` is false), changes to this property’s
// value have no effect; the accuracy is always [kCLLocationAccuracyReduced].
//
// To reduce your app’s impact on battery life, assign a value to this
// property that’s appropriate for your usage. For example, if you need the
// current location only within a kilometer, specify
// [kCLLocationAccuracyKilometer]. More accurate location data also takes more
// time to become available.
//
// After you request high-accuracy location data, your app might still get
// data with a lower accuracy for a period of time. During the time it takes
// to determine the location within the requested accuracy, the location
// service keeps providing the data that’s available, even though that data
// isn’t as accurate as your app requested. Your app receives more accurate
// location data as that data becomes available.
//
// For iOS, the default value of this property is [kCLLocationAccuracyBest].
// For macOS, watchOS, and tvOS, the default value is
// [kCLLocationAccuracyHundredMeters].
//
// This property effects only the standard location services, not for
// monitoring significant location changes.
//
// # Special Considerations
//
// In iOS, this property is declared as `nonatomic`. In macOS, it is declared
// as `atomic`.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLLocationManager/desiredAccuracy
//
// [kCLLocationAccuracyBest]: https://developer.apple.com/documentation/CoreLocation/kCLLocationAccuracyBest
// [kCLLocationAccuracyHundredMeters]: https://developer.apple.com/documentation/CoreLocation/kCLLocationAccuracyHundredMeters
// [kCLLocationAccuracyKilometer]: https://developer.apple.com/documentation/CoreLocation/kCLLocationAccuracyKilometer
// [kCLLocationAccuracyReduced]: https://developer.apple.com/documentation/CoreLocation/kCLLocationAccuracyReduced
func (l CLLocationManager) DesiredAccuracy() CLLocationAccuracy {
	rv := objc.Send[CLLocationAccuracy](l.ID, objc.Sel("desiredAccuracy"))
	return CLLocationAccuracy(rv)
}
func (l CLLocationManager) SetDesiredAccuracy(value CLLocationAccuracy) {
	objc.Send[struct{}](l.ID, objc.Sel("setDesiredAccuracy:"), value)
}

// A Boolean value that indicates whether the location-manager object may
// pause location updates.
//
// # Discussion
//
// Allowing the location manager to pause updates can improve battery life on
// the target device without sacrificing location data. Setting this property
// to true causes the location manager to pause updates (and powers down the
// appropriate hardware) at times when the location data is unlikely to
// change. For example, if the user stops for food while using a navigation
// app, the location manager might pause updates for a period of time. You can
// help the determination of when to pause location updates by assigning a
// value to the [CLLocationManager.ActivityType] property.
//
// After a pause occurs, it’s your responsibility to restart location
// services again when you determine that they’re needed. Core Location
// calls the [LocationManagerDidPauseLocationUpdates] method of your location
// manager’s delegate to let you know that a pause has occurred. In that
// method configure a local notification that has a
// [UNLocationNotificationTrigger] to notify when the user exits the current
// region. The message for the local notification should prompt the user to
// launch your app again so that it can resume updates.
//
// On supported platforms the default value of this property is true;
// otherwise the default value is false and is immutable.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLLocationManager/pausesLocationUpdatesAutomatically
//
// [UNLocationNotificationTrigger]: https://developer.apple.com/documentation/UserNotifications/UNLocationNotificationTrigger
func (l CLLocationManager) PausesLocationUpdatesAutomatically() bool {
	rv := objc.Send[bool](l.ID, objc.Sel("pausesLocationUpdatesAutomatically"))
	return rv
}
func (l CLLocationManager) SetPausesLocationUpdatesAutomatically(value bool) {
	objc.Send[struct{}](l.ID, objc.Sel("setPausesLocationUpdatesAutomatically:"), value)
}

// A Boolean value that indicates whether the app receives location updates
// when running in the background.
//
// # Discussion
//
// Apps that receive location updates when running in the background must
// include the [UIBackgroundModes] key (with the `location` value) in their
// app’s `Info.Plist()` file. After including the [UIBackgroundModes] key,
// set the value of [CLLocationManager.AllowsBackgroundLocationUpdates] to
// true. Use this property to enable and disable background updates
// programmatically. For example, you might set this property to true only
// after the user enables features in your app that require background
// location updates.
//
// When the value of this property is true and you start location updates
// while the app is in the foreground, Core Location configures the system to
// keep the app running to receive continuous background location updates, and
// arranges to show the background location indicator (blue bar or pill) if
// needed. Updates continue even if the app subsequently enters the
// background.
//
// When the value of this property is false, location updates may or may not
// continue in the background depending on other factors, including other
// background modes. In this configuration Core Location doesn’t configure
// the system to keep the app running for delivery, or display the background
// location indicator to extend the effectiveness of the
// [KCLAuthorizationStatusAuthorizedWhenInUse] authorization while the app is
// running in the background.
//
// The default value of this property is false.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLLocationManager/allowsBackgroundLocationUpdates
//
// [UIBackgroundModes]: https://developer.apple.com/documentation/BundleResources/Information-Property-List/UIBackgroundModes
func (l CLLocationManager) AllowsBackgroundLocationUpdates() bool {
	rv := objc.Send[bool](l.ID, objc.Sel("allowsBackgroundLocationUpdates"))
	return rv
}
func (l CLLocationManager) SetAllowsBackgroundLocationUpdates(value bool) {
	objc.Send[struct{}](l.ID, objc.Sel("setAllowsBackgroundLocationUpdates:"), value)
}

// The type of activity the app expects the user to typically perform while in
// the app’s location session.
//
// # Discussion
//
// An app should use `activityType` to communicate to Core Location algorithms
// the type of activity the app’s users typically perform while using the
// app. The location manager uses the information in this property as a cue to
// determine when the system may pause location updates. Pausing updates gives
// the system the opportunity to save power in situations where the user’s
// location isn’t likely to be changing. For example, if the activity type
// is [CLActivityTypeAutomotiveNavigation] and no location changes have
// occurred recently, the system might power down radios until the system
// detects movement again.
//
// After a pause occurs, it’s your responsibility to restart location
// services again when you determine that they’re needed. For more
// information on ways to restart location services after a pause, see the
// discussion of the [CLLocationManager.PausesLocationUpdatesAutomatically]
// property.
//
// The default value of this property is [CLActivityTypeOther].
//
// If your app allows the user to change the type activity they’re
// performing, for example in a navigation app that allows the user to switch
// between driving or walking modes, you should also update the `activityType`
// as well.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLLocationManager/activityType
func (l CLLocationManager) ActivityType() CLActivityType {
	rv := objc.Send[CLActivityType](l.ID, objc.Sel("activityType"))
	return CLActivityType(rv)
}
func (l CLLocationManager) SetActivityType(value CLActivityType) {
	objc.Send[struct{}](l.ID, objc.Sel("setActivityType:"), value)
}

// The minimum angular change in degrees required to generate new heading
// events.
//
// # Discussion
//
// The angular distance is measured relative to the last delivered heading
// event. Use the value [kCLHeadingFilterNone] to be notified of all
// movements. The default value of this property is `1` degree.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLLocationManager/headingFilter
//
// [kCLHeadingFilterNone]: https://developer.apple.com/documentation/CoreLocation/kCLHeadingFilterNone
func (l CLLocationManager) HeadingFilter() CLLocationDegrees {
	rv := objc.Send[CLLocationDegrees](l.ID, objc.Sel("headingFilter"))
	return CLLocationDegrees(rv)
}
func (l CLLocationManager) SetHeadingFilter(value CLLocationDegrees) {
	objc.Send[struct{}](l.ID, objc.Sel("setHeadingFilter:"), value)
}

// The device orientation to use when computing heading values.
//
// # Discussion
//
// When computing heading values, the location manager assumes that the top of
// the device in portrait mode represents due north (0 degrees) by default.
// For apps that run in other orientations, this may not always be the most
// convenient orientation. This property allows you to specify which device
// orientation you want the location manager to use as the reference point for
// due north.
//
// Although you can set the value of this property to
// [CLDeviceOrientationUnknown], [CLDeviceOrientationFaceUp], or
// [CLDeviceOrientationFaceDown], doing so has no effect on the orientation
// reference point. The original reference point is retained instead.
//
// Changing the value in this property affects only those heading values
// reported after the change is made.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLLocationManager/headingOrientation
func (l CLLocationManager) HeadingOrientation() CLDeviceOrientation {
	rv := objc.Send[CLDeviceOrientation](l.ID, objc.Sel("headingOrientation"))
	return CLDeviceOrientation(rv)
}
func (l CLLocationManager) SetHeadingOrientation(value CLDeviceOrientation) {
	objc.Send[struct{}](l.ID, objc.Sel("setHeadingOrientation:"), value)
}

// The set of shared regions monitored by all location-manager objects.
//
// # Discussion
//
// You cannot add regions to this property directly. Instead, you must
// register regions by calling the
// [CLLocationManager.StartMonitoringForRegion] method. The regions in this
// property are shared by all instances of the [CLLocationManager] class in
// your app.
//
// The objects in this set may not necessarily be the same objects you
// specified at registration time. Only the region data itself is maintained
// by the system. Therefore, the only way to uniquely identify a registered
// region is using its [CLRegion.Identifier] property.
//
// The location manager persists region data between launches of your app. If
// your app is terminated and then relaunched, the contents of this property
// are repopulated with region objects that contain the previously registered
// data.
//
// In a compatible iPad or iPhone app running in visionOS, the property
// contains an empty set.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLLocationManager/monitoredRegions
func (l CLLocationManager) MonitoredRegions() foundation.INSSet {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("monitoredRegions"))
	return foundation.NSSetFromID(objc.ID(rv))
}

// The largest boundary distance that can be assigned to a region.
//
// # Discussion
//
// This property defines the largest boundary distance allowed from a
// region’s center point. Attempting to monitor a region with a distance
// larger than this value causes the location manager to send a
// [KCLErrorRegionMonitoringFailure] error to the delegate.
//
// If region monitoring is unavailable or not supported, the value in this
// property is `-1`.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLLocationManager/maximumRegionMonitoringDistance
func (l CLLocationManager) MaximumRegionMonitoringDistance() CLLocationDistance {
	rv := objc.Send[CLLocationDistance](l.ID, objc.Sel("maximumRegionMonitoringDistance"))
	return CLLocationDistance(rv)
}

// The set of beacon constraints currently being tracked using ranging.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLLocationManager/rangedBeaconConstraints
func (l CLLocationManager) RangedBeaconConstraints() foundation.INSSet {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("rangedBeaconConstraints"))
	return foundation.NSSetFromID(objc.ID(rv))
}

// The most recently retrieved user location.
//
// # Discussion
//
// The value of this property is `nil` if no location data has ever been
// retrieved.
//
// In iOS 4.0 and later, this property may contain a more recent location
// object at launch time. Specifically, if significant location updates are
// running and your app is terminated, this property is updated with the most
// recent location data when your app is relaunched (and you create a new
// location manager object). This location data may be more recent than the
// last location event processed by your app.
//
// It is always a good idea to check the timestamp of the location stored in
// this property. If the receiver is currently gathering location data, but
// the minimum distance filter is large, the returned location might be
// relatively old. If it is, you can stop the receiver and start it again to
// force an update.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLLocationManager/location
func (l CLLocationManager) Location() ICLLocation {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("location"))
	return CLLocationFromID(objc.ID(rv))
}

// The most recently reported heading.
//
// # Discussion
//
// The value of this property is `nil` if heading updates have never been
// initiated.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLLocationManager/heading
func (l CLLocationManager) Heading() ICLHeading {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("heading"))
	return CLHeadingFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreLocation/CLLocationManager/headingBody
func (l CLLocationManager) HeadingBody() CLBodyIdentifiable {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("headingBody"))
	return CLBodyIdentifiableObjectFromID(rv)
}
func (l CLLocationManager) SetHeadingBody(value CLBodyIdentifiable) {
	objc.Send[struct{}](l.ID, objc.Sel("setHeadingBody:"), value)
}

// RequestTemporaryFullAccuracyAuthorizationWithPurposeKeyCompletionSync is a synchronous wrapper around [CLLocationManager.RequestTemporaryFullAccuracyAuthorizationWithPurposeKeyCompletion].
// It blocks until the completion handler fires or the context is cancelled.
func (l CLLocationManager) RequestTemporaryFullAccuracyAuthorizationWithPurposeKeyCompletionSync(ctx context.Context, purposeKey string) error {
	done := make(chan error, 1)
	l.RequestTemporaryFullAccuracyAuthorizationWithPurposeKeyCompletion(purposeKey, func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}
