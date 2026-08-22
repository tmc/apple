// Code generated from Apple documentation for CoreLocation. DO NOT EDIT.

// Package corelocation provides Go bindings for the CoreLocation framework.
//
// Obtain the geographic location and orientation of a device.
//
// Core Location provides services that determine a device’s geographic
// location, altitude, and orientation, or its position relative to a nearby
// iBeacon device. The framework gathers data using all available components
// on the device, including the Wi-Fi, GPS, Bluetooth, magnetometer,
// barometer, and cellular hardware.
//
// # Essentials
//
//   - [Configuring your app to use location services]: Prepare your app to start collecting location data.
//   - [Supporting live updates in SwiftUI and Mac Catalyst apps]: Enable background events by adding lifecycle event support.
//   - [CLLocationManager]: The object you use to start and stop the delivery of location-related events to your app. ([CLLocationManagerDelegate], [CLAuthorizationStatus], [CLLocationDistance], [CLLocationAccuracy], [CLActivityType])
//
// # Authorization
//
//   - [Requesting authorization to use location services]: Obtain authorization to use location services and manage changes to your app’s authorization status.
//   - [Suspending authorization requests]: Defer the system’s authorization request dialog until your app is ready.
//   - [CLAuthorizationStatus]: Constants that indicate the app’s authorization to use location services.
//   - [CLAccuracyAuthorization]: Constants that indicate the level of location accuracy the app has authorization to use.
//
// # Monitoring
//
//   - [CLMonitor]: An object that monitors the conditions you add to it. ([CLMonitorConfiguration], [CLMonitoringEvent], [CLMonitoringRecord], [CLCircularGeographicCondition], [CLBeaconIdentityCondition])
//   - [CLUpdate]: An object that represents a location update.
//
// # Location updates
//
//   - [Getting the current location of a device]: Start location services and provide information the system needs to optimize power usage for those services.
//   - [Handling location updates in the background]: Configure your app to receive location updates when it isn’t running in the foreground.
//   - [Creating a location push service extension]: Add and configure an extension to enable your location-sharing app to access a person’s location in response to a request from someone else.
//   - [CLLocation]: The latitude, longitude, and course information reported by the system. ([CLLocationDistance], [CLLocationAccuracy], [CLLocationSpeed], [CLLocationDirection], [CLLocationSpeedAccuracy])
//   - [CLLocationCoordinate2D]: The latitude and longitude associated with a location, specified using the WGS 84 reference frame.
//   - [CLFloor]: The floor of a building on which the user’s device is located.
//   - [CLVisit]: Information about the user’s location during a specific period of time.
//   - [CLLocationSourceInformation]: Information about the source that provides a location.
//   - [CLLocationUpdater]: An object that provides device location updates. ([CLLiveUpdateConfiguration])
//
// # Region monitoring
//
//   - [Monitoring the user’s proximity to geographic regions]: Use condition monitoring to determine when the user enters or leaves a geographic region.
//   - [CLRegion]: A base class representing an area that can be monitored.
//
// # iBeacon
//
//   - [Ranging for Beacons]: Configure a device to act as a beacon and to detect surrounding beacons.
//   - [Determining the proximity to an iBeacon device]: Detect beacons and determine the relative distance to them.
//   - [Turning an iOS device into an iBeacon device]: Broadcast iBeacon signals from an iOS device.
//   - [CLBeacon]: Information about an observed iBeacon device and its relative distance to a person’s device. ([CLProximity])
//   - [CLCondition]: The abstract base class that all other conditions derive from.
//   - [CLBeaconIdentityCondition]: A condition that describes the identity characteristics of a beacon.
//   - [CLCircularGeographicCondition]: A circular geographic condition that a center point and radius define.
//
// # Compass headings
//
//   - [Getting heading and course information]: Use a device’s orientation and course information for navigation.
//   - [CLHeading]: The orientation of the user’s device, relative to true or magnetic north. ([CLHeadingComponentValue])
//
// # Geocoding
//
//   - [Converting between coordinates and user-friendly place names]: Convert between a latitude and longitude pair and a more user-friendly description of that location.
//   - [Converting a user’s location to a descriptive placemark]: Transform the user’s location that displays on a map into an informative textual description by reverse geocoding.
//
// # Location push service extension
//
//   - [Location Push Service Extension]: An entitlement to enable a location-sharing app to query someone’s location in response to a push notification.
//   - [CLLocationPushServiceError]: Error codes the location manager returns if starting to monitor for location push notifications fails.
//
// # Errors
//
//   - [KCLErrorDomain]: The domain for Core Location errors.
//   - [KCLErrorUserInfoAlternateRegionKey]: A key in the user information dictionary of an error relating to a delayed region-monitoring response.
//
// # Protocols
//
//   - [CLBodyIdentifiable]
//
// # Enumerations
//
//   - [CLServiceSessionAuthorizationRequirement]//
//
// # Key Types
//
//   - [CLLocationManager] - The object you use to start and stop the delivery of location-related events to your app.
//   - [CLLocation] - The latitude, longitude, and course information reported by the system.
//   - [CLMonitoringEvent] - The object that the framework passes to the monitor’s callback handler upon receiving an event.
//   - [CLUpdate] - An object that represents a location update.
//   - [CLBeacon] - Information about an observed iBeacon device and its relative distance to a person’s device.
//   - [CLHeading] - The orientation of the user’s device, relative to true or magnetic north.
//   - [CLMonitor] - An object that monitors the conditions you add to it.
//   - [CLBeaconIdentityCondition] - A condition that describes the identity characteristics of a beacon.
//   - [CLLocationUpdater] - An object that provides device location updates.
//   - [CLVisit] - Information about the user’s location during a specific period of time.
//
// [Configuring your app to use location services]: https://developer.apple.com/documentation/corelocation/configuring-your-app-to-use-location-services
// [Converting a user’s location to a descriptive placemark]: https://developer.apple.com/documentation/corelocation/converting-a-user-s-location-to-a-descriptive-placemark
// [Converting between coordinates and user-friendly place names]: https://developer.apple.com/documentation/corelocation/converting-between-coordinates-and-user-friendly-place-names
// [Creating a location push service extension]: https://developer.apple.com/documentation/corelocation/creating-a-location-push-service-extension
// [Determining the proximity to an iBeacon device]: https://developer.apple.com/documentation/corelocation/determining-the-proximity-to-an-ibeacon-device
// [Getting heading and course information]: https://developer.apple.com/documentation/corelocation/getting-heading-and-course-information
// [Getting the current location of a device]: https://developer.apple.com/documentation/corelocation/getting-the-current-location-of-a-device
// [Handling location updates in the background]: https://developer.apple.com/documentation/corelocation/handling-location-updates-in-the-background
// [Location Push Service Extension]: https://developer.apple.com/documentation/BundleResources/Entitlements/com.apple.developer.location.push
// [Monitoring the user’s proximity to geographic regions]: https://developer.apple.com/documentation/corelocation/monitoring-the-user-s-proximity-to-geographic-regions
// [Ranging for Beacons]: https://developer.apple.com/documentation/corelocation/ranging-for-beacons
// [Requesting authorization to use location services]: https://developer.apple.com/documentation/corelocation/requesting-authorization-to-use-location-services
// [Supporting live updates in SwiftUI and Mac Catalyst apps]: https://developer.apple.com/documentation/corelocation/supporting-live-updates-in-swiftui-and-mac-catalyst-apps
// [Suspending authorization requests]: https://developer.apple.com/documentation/corelocation/suspending-authorization-requests
// [Turning an iOS device into an iBeacon device]: https://developer.apple.com/documentation/corelocation/turning-an-ios-device-into-an-ibeacon-device
package corelocation

import (
	"fmt"
	"os"

	"github.com/ebitengine/purego"
)

// frameworkPaths lists paths to try when loading the CoreLocation library,
// in order. Frameworks whose symbols live in a known dylib resolve to that
// dylib alone; the rest try the framework bundle first and then a /usr/lib
// dylib fallback, which covers C-API frameworks that are not in the dyld
// shared cache as bundles.
var frameworkPaths = []string{
	"/System/Library/Frameworks/CoreLocation.framework/CoreLocation",
	"/usr/lib/libCoreLocation.dylib",
}

// frameworkHandle is the handle to the loaded framework.
var frameworkHandle uintptr

func init() {
	for _, path := range frameworkPaths {
		h, err := purego.Dlopen(path, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
		if err == nil {
			frameworkHandle = h
			return
		}
	}
	// Loading is best-effort: the warning is silent by default because a missing
	// framework is harmless unless one of its symbols is actually called. Set
	// APPLE_FRAMEWORK_LOAD_DEBUG to surface load failures while diagnosing.
	if os.Getenv("APPLE_FRAMEWORK_LOAD_DEBUG") != "" {
		fmt.Fprintf(os.Stderr, "warning: CoreLocation: failed to load framework from any known path\n")
	}
}
