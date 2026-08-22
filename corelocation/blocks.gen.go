// Code generated from Apple documentation. DO NOT EDIT.

package corelocation

import (
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// CLGeocodeCompletionHandler handles A block to be called when a geocoding request is complete.

// CLLocationArrayErrorHandler is the signature for a completion handler block.
//
// Used by:
//   - [CLLocationManager.RequestHistoricalLocationsWithPurposeKeySampleCountCompletionHandler]
type CLLocationArrayErrorHandler = func(*[]CLLocation, error)

// NewCLLocationArrayErrorBlock wraps a Go [CLLocationArrayErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [CLLocationManager.RequestHistoricalLocationsWithPurposeKeySampleCountCompletionHandler]
func NewCLLocationArrayErrorBlock(handler CLLocationArrayErrorHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, errID objc.ID) {
		var result *[]CLLocation
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			obj := foundation.NSArrayFromID(resultID)
			count := obj.Count()
			res := make([]CLLocation, count)
			for i := uint(0); i < count; i++ {
				item := obj.ObjectAtIndex(i)
				res[i] = CLLocationFromID(item.GetID())
			}
			result = &res
		}
		handler(result, foundation.SafeErrorFrom(errID))
	})
	return objc.ID(block), func() { block.Release() }
}

// CLMonitorCLMonitoringEventHandler is the signature for a completion handler block.
//
// Used by:
//   - [CLMonitorConfiguration.ConfigWithMonitorNameQueueEventHandler]
type CLMonitorCLMonitoringEventHandler = func(*CLMonitor, *CLMonitoringEvent)

// NewCLMonitorCLMonitoringEventBlock wraps a Go [CLMonitorCLMonitoringEventHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [CLMonitorConfiguration.ConfigWithMonitorNameQueueEventHandler]
func NewCLMonitorCLMonitoringEventBlock(handler CLMonitorCLMonitoringEventHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, extra0ID objc.ID) {
		var result *CLMonitor
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := CLMonitorFromID(resultID)
			result = &v
		}
		var extra0 *CLMonitoringEvent
		if extra0ID != 0 {
			objc.Send[objc.ID](extra0ID, objc.Sel("retain"))
			v := CLMonitoringEventFromID(extra0ID)
			extra0 = &v
		}
		handler(result, extra0)
	})
	return objc.ID(block), func() { block.Release() }
}

// CLMonitorHandler handles The handler the framework calls with events that satisfy the monitor’s conditions.
//
// Used by:
//   - [CLMonitor.RequestMonitorWithConfigurationCompletion]
type CLMonitorHandler = func(*CLMonitor)

// NewCLMonitorBlock wraps a Go [CLMonitorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [CLMonitor.RequestMonitorWithConfigurationCompletion]
func NewCLMonitorBlock(handler CLMonitorHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID) {
		var result *CLMonitor
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := CLMonitorFromID(resultID)
			result = &v
		}
		handler(result)
	})
	return objc.ID(block), func() { block.Release() }
}

// CLPlacemarkArrayErrorHandler handles The handler block to execute with the results.
// The error can be type-asserted to *foundation.NSError for Domain, Code, and UserInfo.
//
// Used by:
//   - [CLGeocoder.GeocodeAddressStringCompletionHandler]
//   - [CLGeocoder.GeocodeAddressStringInRegionCenteredAtInRegionRadiusPreferredLocaleCompletionHandler]
//   - [CLGeocoder.GeocodeAddressStringInRegionCompletionHandler]
//   - [CLGeocoder.GeocodeAddressStringInRegionPreferredLocaleCompletionHandler]
//   - [CLGeocoder.GeocodePostalAddressCompletionHandler]
//   - [CLGeocoder.GeocodePostalAddressPreferredLocaleCompletionHandler]
//   - [CLGeocoder.ReverseGeocodeLocationCompletionHandler]
//   - [CLGeocoder.ReverseGeocodeLocationPreferredLocaleCompletionHandler]
type CLPlacemarkArrayErrorHandler = func(*[]CLPlacemark, error)

// NewCLPlacemarkArrayErrorBlock wraps a Go [CLPlacemarkArrayErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [CLGeocoder.GeocodeAddressStringCompletionHandler]
//   - [CLGeocoder.GeocodeAddressStringInRegionCenteredAtInRegionRadiusPreferredLocaleCompletionHandler]
//   - [CLGeocoder.GeocodeAddressStringInRegionCompletionHandler]
//   - [CLGeocoder.GeocodeAddressStringInRegionPreferredLocaleCompletionHandler]
//   - [CLGeocoder.GeocodePostalAddressCompletionHandler]
//   - [CLGeocoder.GeocodePostalAddressPreferredLocaleCompletionHandler]
//   - [CLGeocoder.ReverseGeocodeLocationCompletionHandler]
//   - [CLGeocoder.ReverseGeocodeLocationPreferredLocaleCompletionHandler]
func NewCLPlacemarkArrayErrorBlock(handler CLPlacemarkArrayErrorHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, errID objc.ID) {
		var result *[]CLPlacemark
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			obj := foundation.NSArrayFromID(resultID)
			count := obj.Count()
			res := make([]CLPlacemark, count)
			for i := uint(0); i < count; i++ {
				item := obj.ObjectAtIndex(i)
				res[i] = CLPlacemarkFromID(item.GetID())
			}
			result = &res
		}
		handler(result, foundation.SafeErrorFrom(errID))
	})
	return objc.ID(block), func() { block.Release() }
}

// CLUpdateHandler handles The block that the framework invokes with each update.
//
// Used by:
//   - [CLLocationUpdater.LiveUpdaterWithConfigurationQueueHandler]
//   - [CLLocationUpdater.LiveUpdaterWithQueueHandler]
type CLUpdateHandler = func(*CLUpdate)

// NewCLUpdateBlock wraps a Go [CLUpdateHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [CLLocationUpdater.LiveUpdaterWithConfigurationQueueHandler]
//   - [CLLocationUpdater.LiveUpdaterWithQueueHandler]
func NewCLUpdateBlock(handler CLUpdateHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID) {
		var result *CLUpdate
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := CLUpdateFromID(resultID)
			result = &v
		}
		handler(result)
	})
	return objc.ID(block), func() { block.Release() }
}

// DataErrorHandler handles The completion handler to call after you start monitoring location pushes.
// The error can be type-asserted to *foundation.NSError for Domain, Code, and UserInfo.
//
// Used by:
//   - [CLLocationManager.StartMonitoringLocationPushesWithCompletion]
type DataErrorHandler = func(*foundation.NSData, error)

// NewDataErrorBlock wraps a Go [DataErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [CLLocationManager.StartMonitoringLocationPushesWithCompletion]
func NewDataErrorBlock(handler DataErrorHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, errID objc.ID) {
		var result *foundation.NSData
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := foundation.NSDataFromID(resultID)
			result = &v
		}
		handler(result, foundation.SafeErrorFrom(errID))
	})
	return objc.ID(block), func() { block.Release() }
}

// ErrorHandler handles A closure to execute after authorization status changes.
// The error can be type-asserted to *foundation.NSError for Domain, Code, and UserInfo.
//
// Used by:
//   - [CLLocationManager.RequestTemporaryFullAccuracyAuthorizationWithPurposeKeyCompletion]
type ErrorHandler = func(error)

// NewErrorBlock wraps a Go [ErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [CLLocationManager.RequestTemporaryFullAccuracyAuthorizationWithPurposeKeyCompletion]
func NewErrorBlock(handler ErrorHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, errID objc.ID) {
		handler(foundation.SafeErrorFrom(errID))
	})
	objc.SetNSErrorBlockSignature(block)
	return objc.ID(block), func() { block.Release() }
}
