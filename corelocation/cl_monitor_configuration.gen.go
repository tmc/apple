// Code generated from Apple documentation for CoreLocation. DO NOT EDIT.

package corelocation

import (
	"sync"

	"github.com/tmc/apple/dispatch"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CLMonitorConfiguration] class.
var (
	_CLMonitorConfigurationClass     CLMonitorConfigurationClass
	_CLMonitorConfigurationClassOnce sync.Once
)

func getCLMonitorConfigurationClass() CLMonitorConfigurationClass {
	_CLMonitorConfigurationClassOnce.Do(func() {
		_CLMonitorConfigurationClass = CLMonitorConfigurationClass{class: objc.GetClass("CLMonitorConfiguration")}
	})
	return _CLMonitorConfigurationClass
}

// GetCLMonitorConfigurationClass returns the class object for CLMonitorConfiguration.
func GetCLMonitorConfigurationClass() CLMonitorConfigurationClass {
	return getCLMonitorConfigurationClass()
}

type CLMonitorConfigurationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CLMonitorConfigurationClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CLMonitorConfigurationClass) Alloc() CLMonitorConfiguration {
	rv := objc.Send[CLMonitorConfiguration](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// An object for configuring a location monitor instance.
//
// # Instance properties
//
//   - [CLMonitorConfiguration.EventHandler]: The block the framework calls as the event handler for the location monitor instance.
//   - [CLMonitorConfiguration.Name]: The name of the monitor instance.
//   - [CLMonitorConfiguration.Queue]: The dispatch queue to bind the instance of a location monitor to.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLMonitorConfiguration
type CLMonitorConfiguration struct {
	objectivec.Object
}

// CLMonitorConfigurationFromID constructs a [CLMonitorConfiguration] from an objc.ID.
//
// An object for configuring a location monitor instance.
func CLMonitorConfigurationFromID(id objc.ID) CLMonitorConfiguration {
	return CLMonitorConfiguration{objectivec.Object{ID: id}}
}

// Ensure CLMonitorConfiguration implements ICLMonitorConfiguration.
var _ ICLMonitorConfiguration = CLMonitorConfiguration{}

// An interface definition for the [CLMonitorConfiguration] class.
//
// # Instance properties
//
//   - [ICLMonitorConfiguration.EventHandler]: The block the framework calls as the event handler for the location monitor instance.
//   - [ICLMonitorConfiguration.Name]: The name of the monitor instance.
//   - [ICLMonitorConfiguration.Queue]: The dispatch queue to bind the instance of a location monitor to.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLMonitorConfiguration
type ICLMonitorConfiguration interface {
	objectivec.IObject

	// Topic: Instance properties

	// The block the framework calls as the event handler for the location monitor instance.
	EventHandler() CLMonitorCLMonitoringEventHandler
	// The name of the monitor instance.
	Name() string
	// The dispatch queue to bind the instance of a location monitor to.
	Queue() dispatch.Queue
}

// Init initializes the instance.
func (m CLMonitorConfiguration) Init() CLMonitorConfiguration {
	rv := objc.Send[CLMonitorConfiguration](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m CLMonitorConfiguration) Autorelease() CLMonitorConfiguration {
	rv := objc.Send[CLMonitorConfiguration](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewCLMonitorConfiguration creates a new CLMonitorConfiguration instance.
func NewCLMonitorConfiguration() CLMonitorConfiguration {
	class := getCLMonitorConfigurationClass()
	rv := objc.Send[CLMonitorConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a location monitor instance with the name, dispatch queue, and
// event handler you specify.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLMonitorConfiguration/configWithMonitorName:queue:eventHandler:
func (_CLMonitorConfigurationClass CLMonitorConfigurationClass) ConfigWithMonitorNameQueueEventHandler(name string, queue dispatch.Queue, eventHandler CLMonitorCLMonitoringEventHandler) CLMonitorConfiguration {
	_block2, _ := NewCLMonitorCLMonitoringEventBlock(eventHandler)
	rv := objc.Send[objc.ID](objc.ID(_CLMonitorConfigurationClass.class), objc.Sel("configWithMonitorName:queue:eventHandler:"), objc.String(name), uintptr(queue.Handle()), _block2)
	return CLMonitorConfigurationFromID(rv)
}

// The block the framework calls as the event handler for the location monitor
// instance.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLMonitorConfiguration/eventHandler
func (m CLMonitorConfiguration) EventHandler() CLMonitorCLMonitoringEventHandler {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("eventHandler"))
	_ = rv
	return nil
}

// The name of the monitor instance.
//
// # Discussion
//
// `name` can contain only alphanumeric characters and can’t start with an
// underscore (_).
//
// See: https://developer.apple.com/documentation/CoreLocation/CLMonitorConfiguration/name
func (m CLMonitorConfiguration) Name() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("name"))
	return foundation.NSStringFromID(rv).String()
}

// The dispatch queue to bind the instance of a location monitor to.
//
// # Discussion
//
// You need to perform all interactions related to the [CLMonitor] instance on
// this queue, and the framework delivers events that the [CLMonitor] instance
// generates to the handler on this queue.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLMonitorConfiguration/queue
//
// [CLMonitor]: https://developer.apple.com/documentation/CoreLocation/CLMonitor-2r51v
func (m CLMonitorConfiguration) Queue() dispatch.Queue {
	rv := objc.Send[uintptr](m.ID, objc.Sel("queue"))
	return dispatch.QueueFromHandle(rv)
}
