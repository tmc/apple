// Code generated from Apple documentation for CoreLocation. DO NOT EDIT.

package corelocation

import (
	"context"
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CLMonitor] class.
var (
	_CLMonitorClass     CLMonitorClass
	_CLMonitorClassOnce sync.Once
)

func getCLMonitorClass() CLMonitorClass {
	_CLMonitorClassOnce.Do(func() {
		_CLMonitorClass = CLMonitorClass{class: objc.GetClass("CLMonitor")}
	})
	return _CLMonitorClass
}

// GetCLMonitorClass returns the class object for CLMonitor.
func GetCLMonitorClass() CLMonitorClass {
	return getCLMonitorClass()
}

type CLMonitorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CLMonitorClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CLMonitorClass) Alloc() CLMonitor {
	rv := objc.Send[CLMonitor](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// An object that monitors the conditions you add to it.
//
// # Overview
//
// Use [CLMonitor] to monitor for and observe events such as the entry to a
// specific geographic area or proximity to a beacon with characteristics that
// you specify.
//
// This service is unavailable in a compatible iPad or iPhone app running in
// visionOS.
//
// # Accessing the location monitor’s identifiers
//
//   - [CLMonitor.MonitoredIdentifiers]: An array that contains all the identifiers for each condition that the monitor is monitoring.
//   - [CLMonitor.Name]: The name associated with the location monitor instance.
//
// # Adding and removing conditions
//
//   - [CLMonitor.AddConditionForMonitoringIdentifier]: Adds a condition to monitor with the identifier you provide.
//   - [CLMonitor.AddConditionForMonitoringIdentifierAssumedState]: Adds a condition to monitor with the state and identifier you provide.
//   - [CLMonitor.MonitoringRecordForIdentifier]: Gets the monitoring record containing the condition and most recent monitoring event for the identifier you supply, if applicable.
//   - [CLMonitor.RemoveConditionFromMonitoringWithIdentifier]: Removes the monitoring record with the identifier from monitoring.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLMonitor-6ynwz
type CLMonitor struct {
	objectivec.Object
}

// CLMonitorFromID constructs a [CLMonitor] from an objc.ID.
//
// An object that monitors the conditions you add to it.
func CLMonitorFromID(id objc.ID) CLMonitor {
	return CLMonitor{objectivec.Object{ID: id}}
}

// Ensure CLMonitor implements ICLMonitor.
var _ ICLMonitor = CLMonitor{}

// An interface definition for the [CLMonitor] class.
//
// # Accessing the location monitor’s identifiers
//
//   - [ICLMonitor.MonitoredIdentifiers]: An array that contains all the identifiers for each condition that the monitor is monitoring.
//   - [ICLMonitor.Name]: The name associated with the location monitor instance.
//
// # Adding and removing conditions
//
//   - [ICLMonitor.AddConditionForMonitoringIdentifier]: Adds a condition to monitor with the identifier you provide.
//   - [ICLMonitor.AddConditionForMonitoringIdentifierAssumedState]: Adds a condition to monitor with the state and identifier you provide.
//   - [ICLMonitor.MonitoringRecordForIdentifier]: Gets the monitoring record containing the condition and most recent monitoring event for the identifier you supply, if applicable.
//   - [ICLMonitor.RemoveConditionFromMonitoringWithIdentifier]: Removes the monitoring record with the identifier from monitoring.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLMonitor-6ynwz
type ICLMonitor interface {
	objectivec.IObject

	// Topic: Accessing the location monitor’s identifiers

	// An array that contains all the identifiers for each condition that the monitor is monitoring.
	MonitoredIdentifiers() []string
	// The name associated with the location monitor instance.
	Name() string

	// Topic: Adding and removing conditions

	// Adds a condition to monitor with the identifier you provide.
	AddConditionForMonitoringIdentifier(condition ICLCondition, identifier string)
	// Adds a condition to monitor with the state and identifier you provide.
	AddConditionForMonitoringIdentifierAssumedState(condition ICLCondition, identifier string, state CLMonitoringState)
	// Gets the monitoring record containing the condition and most recent monitoring event for the identifier you supply, if applicable.
	MonitoringRecordForIdentifier(identifier string) ICLMonitoringRecord
	// Removes the monitoring record with the identifier from monitoring.
	RemoveConditionFromMonitoringWithIdentifier(identifier string)
}

// Init initializes the instance.
func (m CLMonitor) Init() CLMonitor {
	rv := objc.Send[CLMonitor](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m CLMonitor) Autorelease() CLMonitor {
	rv := objc.Send[CLMonitor](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewCLMonitor creates a new CLMonitor instance.
func NewCLMonitor() CLMonitor {
	class := getCLMonitorClass()
	rv := objc.Send[CLMonitor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Adds a condition to monitor with the identifier you provide.
//
// condition: A [CLCondition] to monitor for.
//
// identifier: A string you use to identify this condition.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLMonitor-6ynwz/addConditionForMonitoring:identifier:
//
// [CLCondition]: https://developer.apple.com/documentation/CoreLocation/CLCondition-swift.protocol
func (m CLMonitor) AddConditionForMonitoringIdentifier(condition ICLCondition, identifier string) {
	objc.Send[objc.ID](m.ID, objc.Sel("addConditionForMonitoring:identifier:"), condition, objc.String(identifier))
}

// Adds a condition to monitor with the state and identifier you provide.
//
// condition: A [CLCondition] to monitor for.
//
// identifier: A string you use to identify this condition.
//
// state: A [CLMonitoringState] that satisfies `condition`.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLMonitor-6ynwz/addConditionForMonitoring:identifier:assumedState:
//
// [CLCondition]: https://developer.apple.com/documentation/CoreLocation/CLCondition-swift.protocol
// [CLMonitoringState]: https://developer.apple.com/documentation/CoreLocation/CLMonitoringState
func (m CLMonitor) AddConditionForMonitoringIdentifierAssumedState(condition ICLCondition, identifier string, state CLMonitoringState) {
	objc.Send[objc.ID](m.ID, objc.Sel("addConditionForMonitoring:identifier:assumedState:"), condition, objc.String(identifier), state)
}

// Gets the monitoring record containing the condition and most recent
// monitoring event for the identifier you supply, if applicable.
//
// identifier: A string that identifies the monitoring record.
//
// # Return Value
//
// The monitoring record; otherwise, nil.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLMonitor-6ynwz/monitoringRecordForIdentifier:
func (m CLMonitor) MonitoringRecordForIdentifier(identifier string) ICLMonitoringRecord {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("monitoringRecordForIdentifier:"), objc.String(identifier))
	return CLMonitoringRecordFromID(rv)
}

// Removes the monitoring record with the identifier from monitoring.
//
// identifier: A string that identifies the monitoring record.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLMonitor-6ynwz/removeConditionFromMonitoringWithIdentifier:
func (m CLMonitor) RemoveConditionFromMonitoringWithIdentifier(identifier string) {
	objc.Send[objc.ID](m.ID, objc.Sel("removeConditionFromMonitoringWithIdentifier:"), objc.String(identifier))
}

// Creates a location monitor with the configuration and event handler you
// provide.
//
// config: The configuration that describes the conditions that satisfy this location
// monitor.
//
// completionHandler: The handler the framework calls with events that satisfy the monitor’s
// conditions.
//
// # Discussion
//
// Configure the location monitor by adding or removing conditions for
// monitoring by this instance of [CLMonitor]. When an event occurs, the
// framework calls the block you pass in on the specified queue and delivers
// an instance of [CLMonitoringEvent]. The event contains the identifier for
// the condition you’re monitoring, an optional instance of [CLCondition]
// containing specifics, the new state, and the event’s timestamp.
//
// All interaction directly with the returned [CLMonitor] needs to occur on
// the specified queue. Failing to do so results in undefined behavior.
//
// Conditions you add to an instance of [CLMonitor] persist until you remove
// them from monitoring. However, Core Location stops monitoring conditions if
// an event is pending for them, and there isn’t a configured [CLMonitor] to
// receive it.
//
// The framework stores conditions in an opaque file at
// `~/Library/CoreLocation/BundleId`. Alternatively, you can access the
// conditions in the file at `~/Library/CoreLocation/`process
// name`/name.monitor`. You can determine your app’s process name using the
// ActivityMonitor app or by using the UNIX `ps -al` command in Terminal.
//
// Note that for containerized apps, this is inside the data container. Apps
// need to observe when protected data becomes available using
// [protectedDataDidBecomeAvailableNotification] before creating a [CLMonitor]
// instance. Persistence of conditions enables an app to query efficiently for
// conditions it’s currently monitoring and the most recent event it
// delivers for each.
//
// The app can choose to initialize the monitoring state for a condition. By
// default, the monitoring state is [CLMonitoringStateUnknown].
//
// See: https://developer.apple.com/documentation/CoreLocation/CLMonitor-6ynwz/requestMonitorWithConfiguration:completion:
//
// [CLCondition]: https://developer.apple.com/documentation/CoreLocation/CLCondition-swift.protocol
// [CLMonitor]: https://developer.apple.com/documentation/CoreLocation/CLMonitor-2r51v
// [protectedDataDidBecomeAvailableNotification]: https://developer.apple.com/documentation/UIKit/UIApplication/protectedDataDidBecomeAvailableNotification
func (_CLMonitorClass CLMonitorClass) RequestMonitorWithConfigurationCompletion(config ICLMonitorConfiguration, completionHandler CLMonitorHandler) {
	_block1, _ := NewCLMonitorBlock(completionHandler)
	objc.Send[objc.ID](objc.ID(_CLMonitorClass.class), objc.Sel("requestMonitorWithConfiguration:completion:"), config, _block1)
}

// An array that contains all the identifiers for each condition that the
// monitor is monitoring.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLMonitor-6ynwz/monitoredIdentifiers
func (m CLMonitor) MonitoredIdentifiers() []string {
	rv := objc.Send[[]objc.ID](m.ID, objc.Sel("monitoredIdentifiers"))
	return objc.ConvertSliceToStrings(rv)
}

// The name associated with the location monitor instance.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLMonitor-6ynwz/name
func (m CLMonitor) Name() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("name"))
	return foundation.NSStringFromID(rv).String()
}

// RequestMonitorWithConfigurationCompletionSync is a synchronous wrapper around [CLMonitor.RequestMonitorWithConfigurationCompletion].
// It blocks until the completion handler fires or the context is cancelled.
func (mc CLMonitorClass) RequestMonitorWithConfigurationCompletionSync(ctx context.Context, config ICLMonitorConfiguration) (*CLMonitor, error) {
	done := make(chan *CLMonitor, 1)
	mc.RequestMonitorWithConfigurationCompletion(config, func(val *CLMonitor) {
		done <- val
	})
	select {
	case r := <-done:
		return r, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}
