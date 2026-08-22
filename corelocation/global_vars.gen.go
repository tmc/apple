// Code generated from Apple documentation. DO NOT EDIT.

package corelocation

import (
	"github.com/ebitengine/purego"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

var (
	// CLLocationDistanceMax is a constant indicating the maximum distance.
	//
	// See: https://developer.apple.com/documentation/CoreLocation/CLLocationDistanceMax
	CLLocationDistanceMax CLLocationDistance
	// KCLDistanceFilterNone is a constant indicating that all movement should be reported.
	//
	// See: https://developer.apple.com/documentation/CoreLocation/kCLDistanceFilterNone
	KCLDistanceFilterNone CLLocationDistance
	// KCLHeadingFilterNone is a constant indicating that all header values should be reported.
	//
	// See: https://developer.apple.com/documentation/CoreLocation/kCLHeadingFilterNone
	KCLHeadingFilterNone CLLocationDegrees
	// KCLLocationAccuracyBest is the best level of accuracy available.
	//
	// See: https://developer.apple.com/documentation/CoreLocation/kCLLocationAccuracyBest
	KCLLocationAccuracyBest CLLocationAccuracy
	// KCLLocationAccuracyBestForNavigation is the highest possible accuracy that uses additional sensor data to facilitate navigation apps.
	//
	// See: https://developer.apple.com/documentation/CoreLocation/kCLLocationAccuracyBestForNavigation
	KCLLocationAccuracyBestForNavigation CLLocationAccuracy
	// KCLLocationAccuracyHundredMeters is accurate to within one hundred meters.
	//
	// See: https://developer.apple.com/documentation/CoreLocation/kCLLocationAccuracyHundredMeters
	KCLLocationAccuracyHundredMeters CLLocationAccuracy
	// KCLLocationAccuracyKilometer is accurate to the nearest kilometer.
	//
	// See: https://developer.apple.com/documentation/CoreLocation/kCLLocationAccuracyKilometer
	KCLLocationAccuracyKilometer CLLocationAccuracy
	// KCLLocationAccuracyNearestTenMeters is accurate to within ten meters of the desired target.
	//
	// See: https://developer.apple.com/documentation/CoreLocation/kCLLocationAccuracyNearestTenMeters
	KCLLocationAccuracyNearestTenMeters CLLocationAccuracy
	// KCLLocationAccuracyReduced is the level of accuracy used when an app isn’t authorized for full accuracy location data.
	//
	// See: https://developer.apple.com/documentation/CoreLocation/kCLLocationAccuracyReduced
	KCLLocationAccuracyReduced CLLocationAccuracy
	// KCLLocationAccuracyThreeKilometers is accurate to the nearest three kilometers.
	//
	// See: https://developer.apple.com/documentation/CoreLocation/kCLLocationAccuracyThreeKilometers
	KCLLocationAccuracyThreeKilometers CLLocationAccuracy
)

var (
	// CLTimeIntervalMax is a value representing an unlimited amount of time.
	//
	// See: https://developer.apple.com/documentation/CoreLocation/CLTimeIntervalMax
	CLTimeIntervalMax foundation.NSTimeInterval
)

var (
	// KCLErrorDomain is the domain for Core Location errors.
	//
	// See: https://developer.apple.com/documentation/CoreLocation/kCLErrorDomain
	KCLErrorDomain string
	// KCLErrorUserInfoAlternateRegionKey is a key in the user information dictionary of an error relating to a delayed region-monitoring response.
	//
	// See: https://developer.apple.com/documentation/CoreLocation/kCLErrorUserInfoAlternateRegionKey
	KCLErrorUserInfoAlternateRegionKey string
)

var (
	// KCLLocationCoordinate2DInvalid is an invalid coordinate value.
	//
	// See: https://developer.apple.com/documentation/CoreLocation/kCLLocationCoordinate2DInvalid
	KCLLocationCoordinate2DInvalid CLLocationCoordinate2D
)

func init() {
	if frameworkHandle == 0 {
		return
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "CLLocationDistanceMax"); err == nil && ptr != 0 {
		CLLocationDistanceMax = objc.ValueAt[CLLocationDistance](ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "CLTimeIntervalMax"); err == nil && ptr != 0 {
		CLTimeIntervalMax = objc.ValueAt[foundation.NSTimeInterval](ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCLDistanceFilterNone"); err == nil && ptr != 0 {
		KCLDistanceFilterNone = objc.ValueAt[CLLocationDistance](ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCLErrorDomain"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCLErrorDomain = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCLErrorUserInfoAlternateRegionKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				KCLErrorUserInfoAlternateRegionKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCLHeadingFilterNone"); err == nil && ptr != 0 {
		KCLHeadingFilterNone = objc.ValueAt[CLLocationDegrees](ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCLLocationAccuracyBest"); err == nil && ptr != 0 {
		KCLLocationAccuracyBest = objc.ValueAt[CLLocationAccuracy](ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCLLocationAccuracyBestForNavigation"); err == nil && ptr != 0 {
		KCLLocationAccuracyBestForNavigation = objc.ValueAt[CLLocationAccuracy](ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCLLocationAccuracyHundredMeters"); err == nil && ptr != 0 {
		KCLLocationAccuracyHundredMeters = objc.ValueAt[CLLocationAccuracy](ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCLLocationAccuracyKilometer"); err == nil && ptr != 0 {
		KCLLocationAccuracyKilometer = objc.ValueAt[CLLocationAccuracy](ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCLLocationAccuracyNearestTenMeters"); err == nil && ptr != 0 {
		KCLLocationAccuracyNearestTenMeters = objc.ValueAt[CLLocationAccuracy](ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCLLocationAccuracyReduced"); err == nil && ptr != 0 {
		KCLLocationAccuracyReduced = objc.ValueAt[CLLocationAccuracy](ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCLLocationAccuracyThreeKilometers"); err == nil && ptr != 0 {
		KCLLocationAccuracyThreeKilometers = objc.ValueAt[CLLocationAccuracy](ptr)
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kCLLocationCoordinate2DInvalid"); err == nil && ptr != 0 {
		KCLLocationCoordinate2DInvalid = objc.ValueAt[CLLocationCoordinate2D](ptr)
	}

}
