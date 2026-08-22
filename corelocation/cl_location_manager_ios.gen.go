// Code generated from Apple documentation for CoreLocation. DO NOT EDIT.
//go:build ios
// +build ios

package corelocation

import (
	"github.com/tmc/apple/objc"
)

// Stops the generation of heading updates.
//
// # Discussion
//
// Call this method whenever your code no longer needs to receive
// heading-related events. Disabling event delivery gives the receiver the
// option of disabling the appropriate hardware (and thereby saving power)
// when no clients need location data. You can always restart the generation
// of heading updates by calling the [CLLocationManager.StartUpdatingHeading]
// method again.
//
// If a compatible iPad or iPhone app calls this method when running in
// visionOS, the method does nothing.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLLocationManager/stopUpdatingHeading()
func (l CLLocationManager) StopUpdatingHeading() {
	objc.Send[objc.ID](l.ID, objc.Sel("stopUpdatingHeading"))
}

// Starts monitoring for the delivery of Apple Push Notification service
// (APNs) location pushes, and provides a device-specific token for sending
// pushes.
//
// completion: The completion handler to call after you start monitoring location pushes.
// The completion handler takes the following parameters:
//
// `token`: A globally unique token that identifies this device to APNs. Send
// this `token` to the server that you use to generate location pushes. Your
// server passes this `token` — unmodified — back to APNs when sending
// pushes. APNs device tokens are of variable length. Don’t hard-code their
// size. If an error occurs, `token` is `nil`. `error`: If your app is unable
// to register for location pushes, the system sets this parameter to an error
// object that contains information about why it failed; otherwise it’s
// `nil`. The error type is [CLLocationPushServiceError].
//
// # Discussion
//
// This function requests an Apple Push Notification service (APNs) token that
// the system uses to launch your Location Push Service Extension and deliver
// pushes. Devices need an Internet connection to receive the token. Your
// completion block receives the token if the call succeeds, otherwise it
// receives error information. If a compatible iPad or iPhone app calls this
// method when running in visionOS, the method does nothing.
//
// To use location push notifications, your app must have the
// `com.AppleXCUIElementTypeDeveloperXCUIElementTypeLocationXCUIElementTypePush()`
// entitlement. For more information about implementing location pushes in
// your app, see [Creating a location push service extension].
//
// See: https://developer.apple.com/documentation/CoreLocation/CLLocationManager/startMonitoringLocationPushes(completion:)
//
// [CLLocationPushServiceError]: https://developer.apple.com/documentation/CoreLocation/CLLocationPushServiceError-swift.struct
// [Creating a location push service extension]: https://developer.apple.com/documentation/CoreLocation/creating-a-location-push-service-extension
func (l CLLocationManager) StartMonitoringLocationPushesWithCompletion(completion DataErrorHandler) {
	_block0, _ := NewDataErrorBlock(completion)
	objc.Send[objc.ID](l.ID, objc.Sel("startMonitoringLocationPushesWithCompletion:"), _block0)
}

// Stops monitoring for Apple Push Notification service (APNs) location
// pushes.
//
// # Discussion
//
// Call this method to stop the device from monitoring for APNs location
// pushes. If a compatible iPad or iPhone app calls this method when running
// in visionOS, the method does nothing.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLLocationManager/stopMonitoringLocationPushes()
func (l CLLocationManager) StopMonitoringLocationPushes() {
	objc.Send[objc.ID](l.ID, objc.Sel("stopMonitoringLocationPushes"))
}

// # Discussion
//
// See: https://developer.apple.com/documentation/CoreLocation/CLLocationManager/requestHistoricalLocations(purposeKey:sampleCount:completionHandler:)
func (l CLLocationManager) RequestHistoricalLocationsWithPurposeKeySampleCountCompletionHandler(purposeKey string, sampleCount int, handler CLLocationArrayErrorHandler) {
	_block2, _ := NewCLLocationArrayErrorBlock(handler)
	objc.Send[objc.ID](l.ID, objc.Sel("requestHistoricalLocationsWithPurposeKey:sampleCount:completionHandler:"), objc.String(purposeKey), sampleCount, _block2)
}

// A Boolean value that indicates whether the status bar changes its
// appearance when an app uses location services in the background.
//
// # Discussion
//
// The default value of this property is `false`. The background location
// usage indicator is a blue bar or a blue pill in the status bar on iOS; on
// watchOS the indicator is a small icon. Users can tap the indicator to
// return to your app.
//
// This property affects only apps that received Always authorization. When
// such an app moves to the background, the system uses this property to
// determine whether to change the status bar appearance to indicate that
// location services are in use. Set this value to `true` to maintain
// transparency with the user.
//
// For apps with When In Use authorization, the system changes the appearance
// of the status bar when the app uses location services in the background.
//
// For more information, see [Handling location updates in the background].
//
// See: https://developer.apple.com/documentation/CoreLocation/CLLocationManager/showsBackgroundLocationIndicator
//
// [Handling location updates in the background]: https://developer.apple.com/documentation/CoreLocation/handling-location-updates-in-the-background
func (l CLLocationManager) ShowsBackgroundLocationIndicator() bool {
	rv := objc.Send[bool](l.ID, objc.Sel("showsBackgroundLocationIndicator"))
	return rv
}
func (l CLLocationManager) SetShowsBackgroundLocationIndicator(value bool) {
	objc.Send[struct{}](l.ID, objc.Sel("setShowsBackgroundLocationIndicator:"), value)
}
