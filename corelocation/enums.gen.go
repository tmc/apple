// Code generated from Apple documentation for CoreLocation. DO NOT EDIT.

package corelocation

import (
	"fmt"
)

// See: https://developer.apple.com/documentation/CoreLocation/CLAccuracyAuthorization
type CLAccuracyAuthorization int

const (
	// CLAccuracyAuthorizationFullAccuracy: The user authorized the app to access location data with full accuracy.
	CLAccuracyAuthorizationFullAccuracy CLAccuracyAuthorization = 0
	// CLAccuracyAuthorizationReducedAccuracy: The user authorized the app to access location data with reduced accuracy.
	CLAccuracyAuthorizationReducedAccuracy CLAccuracyAuthorization = 1
)

func (e CLAccuracyAuthorization) String() string {
	switch e {
	case CLAccuracyAuthorizationFullAccuracy:
		return "CLAccuracyAuthorizationFullAccuracy"
	case CLAccuracyAuthorizationReducedAccuracy:
		return "CLAccuracyAuthorizationReducedAccuracy"
	default:
		return fmt.Sprintf("CLAccuracyAuthorization(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreLocation/CLActivityType
type CLActivityType int

const (
	// CLActivityTypeAirborne: The value that indicates activities in the air.
	CLActivityTypeAirborne CLActivityType = 5
	// CLActivityTypeAutomotiveNavigation: The value that indicates positioning in an automobile following a road network.
	CLActivityTypeAutomotiveNavigation CLActivityType = 2
	// CLActivityTypeFitness: The value that indicates positioning during dedicated fitness sessions, such as walking workouts, running workouts, cycling workouts, and so on.
	CLActivityTypeFitness CLActivityType = 3
	// CLActivityTypeOther: The value that indicates the app is using location manager for an unspecified activity.
	CLActivityTypeOther CLActivityType = 1
	// CLActivityTypeOtherNavigation: The value that indicates positioning for activities that don’t or may not adhere to roads such as cycling, scooters, trains, boats and off-road vehicles.
	CLActivityTypeOtherNavigation CLActivityType = 4
)

func (e CLActivityType) String() string {
	switch e {
	case CLActivityTypeAirborne:
		return "CLActivityTypeAirborne"
	case CLActivityTypeAutomotiveNavigation:
		return "CLActivityTypeAutomotiveNavigation"
	case CLActivityTypeFitness:
		return "CLActivityTypeFitness"
	case CLActivityTypeOther:
		return "CLActivityTypeOther"
	case CLActivityTypeOtherNavigation:
		return "CLActivityTypeOtherNavigation"
	default:
		return fmt.Sprintf("CLActivityType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreLocation/CLAuthorizationStatus
type CLAuthorizationStatus int32

const (
	// KCLAuthorizationStatusAuthorizedAlways: The user authorized the app to start location services at any time.
	KCLAuthorizationStatusAuthorizedAlways CLAuthorizationStatus = 3
	// KCLAuthorizationStatusAuthorizedWhenInUse: The user authorized the app to start location services while it is in use.
	KCLAuthorizationStatusAuthorizedWhenInUse CLAuthorizationStatus = 4
	// KCLAuthorizationStatusDenied: The user denied the use of location services for the app or they are disabled globally in Settings.
	KCLAuthorizationStatusDenied CLAuthorizationStatus = 2
	// KCLAuthorizationStatusNotDetermined: The user has not chosen whether the app can use location services.
	KCLAuthorizationStatusNotDetermined CLAuthorizationStatus = 0
	// KCLAuthorizationStatusRestricted: The app is not authorized to use location services.
	KCLAuthorizationStatusRestricted CLAuthorizationStatus = 1
	// Deprecated.
	KCLAuthorizationStatusAuthorized CLAuthorizationStatus = 3
)

func (e CLAuthorizationStatus) String() string {
	switch e {
	case KCLAuthorizationStatusAuthorizedAlways:
		return "KCLAuthorizationStatusAuthorizedAlways"
	case KCLAuthorizationStatusAuthorizedWhenInUse:
		return "KCLAuthorizationStatusAuthorizedWhenInUse"
	case KCLAuthorizationStatusDenied:
		return "KCLAuthorizationStatusDenied"
	case KCLAuthorizationStatusNotDetermined:
		return "KCLAuthorizationStatusNotDetermined"
	case KCLAuthorizationStatusRestricted:
		return "KCLAuthorizationStatusRestricted"
	default:
		return fmt.Sprintf("CLAuthorizationStatus(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreLocation/CLDeviceOrientation
type CLDeviceOrientation int32

const (
	// CLDeviceOrientationFaceDown: The device is held parallel to the ground with the screen facing downwards.
	CLDeviceOrientationFaceDown CLDeviceOrientation = 6
	// CLDeviceOrientationFaceUp: The device is held parallel to the ground with the screen facing upwards.
	CLDeviceOrientationFaceUp CLDeviceOrientation = 5
	// CLDeviceOrientationLandscapeLeft: The device is in landscape mode, with the device held upright and the home button on the right side.
	CLDeviceOrientationLandscapeLeft CLDeviceOrientation = 3
	// CLDeviceOrientationLandscapeRight: The device is in landscape mode, with the device held upright and the home button on the left side.
	CLDeviceOrientationLandscapeRight CLDeviceOrientation = 4
	// CLDeviceOrientationPortrait: The device is in portrait mode, with the device held upright and the home button at the bottom.
	CLDeviceOrientationPortrait CLDeviceOrientation = 1
	// CLDeviceOrientationPortraitUpsideDown: The device is in portrait mode but upside down, with the device held upright and the home button at the top.
	CLDeviceOrientationPortraitUpsideDown CLDeviceOrientation = 2
	// CLDeviceOrientationUnknown: The orientation is currently not known.
	CLDeviceOrientationUnknown CLDeviceOrientation = 0
)

func (e CLDeviceOrientation) String() string {
	switch e {
	case CLDeviceOrientationFaceDown:
		return "CLDeviceOrientationFaceDown"
	case CLDeviceOrientationFaceUp:
		return "CLDeviceOrientationFaceUp"
	case CLDeviceOrientationLandscapeLeft:
		return "CLDeviceOrientationLandscapeLeft"
	case CLDeviceOrientationLandscapeRight:
		return "CLDeviceOrientationLandscapeRight"
	case CLDeviceOrientationPortrait:
		return "CLDeviceOrientationPortrait"
	case CLDeviceOrientationPortraitUpsideDown:
		return "CLDeviceOrientationPortraitUpsideDown"
	case CLDeviceOrientationUnknown:
		return "CLDeviceOrientationUnknown"
	default:
		return fmt.Sprintf("CLDeviceOrientation(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreLocation/CLError-swift.struct/Code
type CLError int

const (
	// KCLErrorDeferredAccuracyTooLow: A constant that indicates deferred mode isn’t supported for the requested accuracy.
	KCLErrorDeferredAccuracyTooLow CLError = 13
	// KCLErrorDeferredCanceled: A constant that indicates your app or the location manager canceled the request for deferred updates.
	KCLErrorDeferredCanceled CLError = 15
	// KCLErrorDeferredDistanceFiltered: A constant that indicates deferred mode doesn’t support distance filters.
	KCLErrorDeferredDistanceFiltered CLError = 14
	// KCLErrorDeferredFailed: A constant that indicates the location manager didn’t enter deferred mode for an unknown reason.
	KCLErrorDeferredFailed CLError = 11
	// KCLErrorDeferredNotUpdatingLocation: A constant that indicates the location manager didn’t enter deferred mode because location updates were already disabled or paused.
	KCLErrorDeferredNotUpdatingLocation CLError = 12
	// KCLErrorDenied: A constant that indicates the user denied access to the location service.
	KCLErrorDenied CLError = 1
	// KCLErrorGeocodeCanceled: A constant that indicates the geocode request was canceled.
	KCLErrorGeocodeCanceled CLError = 10
	// KCLErrorGeocodeFoundNoResult: A constant that indicates the geocode request yielded no result.
	KCLErrorGeocodeFoundNoResult CLError = 8
	// KCLErrorGeocodeFoundPartialResult: A constant that indicates the geocode request yielded a partial result.
	KCLErrorGeocodeFoundPartialResult CLError = 9
	// KCLErrorHeadingFailure: A constant that indicates the location manager can’t determine the heading.
	KCLErrorHeadingFailure          CLError = 3
	KCLErrorHistoricalLocationError CLError = 19
	// KCLErrorLocationUnknown: A constant that indicates the location manager was unable to obtain a location value right now.
	KCLErrorLocationUnknown CLError = 0
	// KCLErrorNetwork: A constant that indicates the network was unavailable or a network error occurred.
	KCLErrorNetwork CLError = 2
	// KCLErrorPromptDeclined: A constant that indicates the user didn’t grant the requested temporary authorization.
	KCLErrorPromptDeclined CLError = 18
	// KCLErrorRangingFailure: A constant that indicates a general ranging error occurred.
	KCLErrorRangingFailure CLError = 17
	// KCLErrorRangingUnavailable: A constant that indicates ranging is disabled.
	KCLErrorRangingUnavailable CLError = 16
	// KCLErrorRegionMonitoringDenied: A constant that indicates the user denied access to the region monitoring service.
	KCLErrorRegionMonitoringDenied CLError = 4
	// KCLErrorRegionMonitoringFailure: A constant that indicates the location manager failed to monitor a registered region.
	KCLErrorRegionMonitoringFailure CLError = 5
	// KCLErrorRegionMonitoringResponseDelayed: A constant that indicates Core Location will deliver events but they may be delayed.
	KCLErrorRegionMonitoringResponseDelayed CLError = 7
	// KCLErrorRegionMonitoringSetupDelayed: A constant that indicates Core Location failed to initialize the region monitoring feature.
	KCLErrorRegionMonitoringSetupDelayed CLError = 6
)

func (e CLError) String() string {
	switch e {
	case KCLErrorDeferredAccuracyTooLow:
		return "KCLErrorDeferredAccuracyTooLow"
	case KCLErrorDeferredCanceled:
		return "KCLErrorDeferredCanceled"
	case KCLErrorDeferredDistanceFiltered:
		return "KCLErrorDeferredDistanceFiltered"
	case KCLErrorDeferredFailed:
		return "KCLErrorDeferredFailed"
	case KCLErrorDeferredNotUpdatingLocation:
		return "KCLErrorDeferredNotUpdatingLocation"
	case KCLErrorDenied:
		return "KCLErrorDenied"
	case KCLErrorGeocodeCanceled:
		return "KCLErrorGeocodeCanceled"
	case KCLErrorGeocodeFoundNoResult:
		return "KCLErrorGeocodeFoundNoResult"
	case KCLErrorGeocodeFoundPartialResult:
		return "KCLErrorGeocodeFoundPartialResult"
	case KCLErrorHeadingFailure:
		return "KCLErrorHeadingFailure"
	case KCLErrorHistoricalLocationError:
		return "KCLErrorHistoricalLocationError"
	case KCLErrorLocationUnknown:
		return "KCLErrorLocationUnknown"
	case KCLErrorNetwork:
		return "KCLErrorNetwork"
	case KCLErrorPromptDeclined:
		return "KCLErrorPromptDeclined"
	case KCLErrorRangingFailure:
		return "KCLErrorRangingFailure"
	case KCLErrorRangingUnavailable:
		return "KCLErrorRangingUnavailable"
	case KCLErrorRegionMonitoringDenied:
		return "KCLErrorRegionMonitoringDenied"
	case KCLErrorRegionMonitoringFailure:
		return "KCLErrorRegionMonitoringFailure"
	case KCLErrorRegionMonitoringResponseDelayed:
		return "KCLErrorRegionMonitoringResponseDelayed"
	case KCLErrorRegionMonitoringSetupDelayed:
		return "KCLErrorRegionMonitoringSetupDelayed"
	default:
		return fmt.Sprintf("CLError(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreLocation/CLLiveUpdateConfiguration
type CLLiveUpdateConfiguration int

const (
	// CLLiveUpdateConfigurationAirborne: A configuration for airborne use cases.
	CLLiveUpdateConfigurationAirborne CLLiveUpdateConfiguration = 4
	// CLLiveUpdateConfigurationAutomotiveNavigation: A configuration for automotive navigation use cases.
	CLLiveUpdateConfigurationAutomotiveNavigation CLLiveUpdateConfiguration = 1
	// CLLiveUpdateConfigurationDefault: The default configuration.
	CLLiveUpdateConfigurationDefault CLLiveUpdateConfiguration = 0
	// CLLiveUpdateConfigurationFitness: A configuration for fitness use cases.
	CLLiveUpdateConfigurationFitness CLLiveUpdateConfiguration = 3
	// CLLiveUpdateConfigurationOtherNavigation: A configuration for other navigation use cases.
	CLLiveUpdateConfigurationOtherNavigation CLLiveUpdateConfiguration = 2
)

func (e CLLiveUpdateConfiguration) String() string {
	switch e {
	case CLLiveUpdateConfigurationAirborne:
		return "CLLiveUpdateConfigurationAirborne"
	case CLLiveUpdateConfigurationAutomotiveNavigation:
		return "CLLiveUpdateConfigurationAutomotiveNavigation"
	case CLLiveUpdateConfigurationDefault:
		return "CLLiveUpdateConfigurationDefault"
	case CLLiveUpdateConfigurationFitness:
		return "CLLiveUpdateConfigurationFitness"
	case CLLiveUpdateConfigurationOtherNavigation:
		return "CLLiveUpdateConfigurationOtherNavigation"
	default:
		return fmt.Sprintf("CLLiveUpdateConfiguration(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreLocation/CLLocationPushServiceError-swift.struct/Code
type CLLocationPushServiceError int

const (
	// CLLocationPushServiceErrorMissingEntitlement: An error code that indicates the app is missing the entitlement it needs to use the location push service.
	CLLocationPushServiceErrorMissingEntitlement CLLocationPushServiceError = 3
	// CLLocationPushServiceErrorMissingPushExtension: An error code that indicates the app is missing a Location Push Service Extension.
	CLLocationPushServiceErrorMissingPushExtension CLLocationPushServiceError = 1
	// CLLocationPushServiceErrorMissingPushServerEnvironment: An error code that indicates the app is missing an Apple Push Notification service (APNs) environment entitlement.
	CLLocationPushServiceErrorMissingPushServerEnvironment CLLocationPushServiceError = 2
	// CLLocationPushServiceErrorUnknown: An error code that indicates the app was unable to start the location push service for an unknown reason.
	CLLocationPushServiceErrorUnknown CLLocationPushServiceError = 0
	// CLLocationPushServiceErrorUnsupportedPlatform: An error code that indicates the location push service isn’t available on this platform.
	CLLocationPushServiceErrorUnsupportedPlatform CLLocationPushServiceError = 4
)

func (e CLLocationPushServiceError) String() string {
	switch e {
	case CLLocationPushServiceErrorMissingEntitlement:
		return "CLLocationPushServiceErrorMissingEntitlement"
	case CLLocationPushServiceErrorMissingPushExtension:
		return "CLLocationPushServiceErrorMissingPushExtension"
	case CLLocationPushServiceErrorMissingPushServerEnvironment:
		return "CLLocationPushServiceErrorMissingPushServerEnvironment"
	case CLLocationPushServiceErrorUnknown:
		return "CLLocationPushServiceErrorUnknown"
	case CLLocationPushServiceErrorUnsupportedPlatform:
		return "CLLocationPushServiceErrorUnsupportedPlatform"
	default:
		return fmt.Sprintf("CLLocationPushServiceError(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreLocation/CLMonitoringState
type CLMonitoringState uint

const (
	// CLMonitoringStateSatisfied: The condition is in a satisfied state.
	CLMonitoringStateSatisfied CLMonitoringState = 1
	// CLMonitoringStateUnknown: The condition is in an unknown state.
	CLMonitoringStateUnknown CLMonitoringState = 0
	// CLMonitoringStateUnmonitored: The condition is in an unmonitored state.
	CLMonitoringStateUnmonitored CLMonitoringState = 3
	// CLMonitoringStateUnsatisfied: The condition is in an unsatisfied state.
	CLMonitoringStateUnsatisfied CLMonitoringState = 2
)

func (e CLMonitoringState) String() string {
	switch e {
	case CLMonitoringStateSatisfied:
		return "CLMonitoringStateSatisfied"
	case CLMonitoringStateUnknown:
		return "CLMonitoringStateUnknown"
	case CLMonitoringStateUnmonitored:
		return "CLMonitoringStateUnmonitored"
	case CLMonitoringStateUnsatisfied:
		return "CLMonitoringStateUnsatisfied"
	default:
		return fmt.Sprintf("CLMonitoringState(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreLocation/CLProximity
type CLProximity int

const (
	// CLProximityFar: The beacon is far away.
	CLProximityFar CLProximity = 3
	// CLProximityImmediate: The beacon is in the user’s immediate vicinity.
	CLProximityImmediate CLProximity = 1
	// CLProximityNear: The beacon is relatively close to the user.
	CLProximityNear CLProximity = 2
	// CLProximityUnknown: The proximity of the beacon could not be determined.
	CLProximityUnknown CLProximity = 0
)

func (e CLProximity) String() string {
	switch e {
	case CLProximityFar:
		return "CLProximityFar"
	case CLProximityImmediate:
		return "CLProximityImmediate"
	case CLProximityNear:
		return "CLProximityNear"
	case CLProximityUnknown:
		return "CLProximityUnknown"
	default:
		return fmt.Sprintf("CLProximity(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreLocation/CLRegionState
type CLRegionState int

const (
	// CLRegionStateInside: The location is inside of the given region.
	CLRegionStateInside CLRegionState = 1
	// CLRegionStateOutside: The location is outside of the given region.
	CLRegionStateOutside CLRegionState = 2
	// CLRegionStateUnknown: It is unknown whether the location is inside or outside of the region.
	CLRegionStateUnknown CLRegionState = 0
)

func (e CLRegionState) String() string {
	switch e {
	case CLRegionStateInside:
		return "CLRegionStateInside"
	case CLRegionStateOutside:
		return "CLRegionStateOutside"
	case CLRegionStateUnknown:
		return "CLRegionStateUnknown"
	default:
		return fmt.Sprintf("CLRegionState(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreLocation/CLServiceSessionAuthorizationRequirement
type CLServiceSessionAuthorizationRequirement int

const (
	CLServiceSessionAuthorizationRequirementAlways    CLServiceSessionAuthorizationRequirement = 2
	CLServiceSessionAuthorizationRequirementNone      CLServiceSessionAuthorizationRequirement = 0
	CLServiceSessionAuthorizationRequirementWhenInUse CLServiceSessionAuthorizationRequirement = 1
)

func (e CLServiceSessionAuthorizationRequirement) String() string {
	switch e {
	case CLServiceSessionAuthorizationRequirementAlways:
		return "CLServiceSessionAuthorizationRequirementAlways"
	case CLServiceSessionAuthorizationRequirementNone:
		return "CLServiceSessionAuthorizationRequirementNone"
	case CLServiceSessionAuthorizationRequirementWhenInUse:
		return "CLServiceSessionAuthorizationRequirementWhenInUse"
	default:
		return fmt.Sprintf("CLServiceSessionAuthorizationRequirement(%d)", e)
	}
}
