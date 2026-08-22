// Code generated from Apple documentation for CoreLocation. DO NOT EDIT.

package corelocation

import (
	"context"
	"sync"

	"github.com/tmc/apple/dispatch"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CLLocationUpdater] class.
var (
	_CLLocationUpdaterClass     CLLocationUpdaterClass
	_CLLocationUpdaterClassOnce sync.Once
)

func getCLLocationUpdaterClass() CLLocationUpdaterClass {
	_CLLocationUpdaterClassOnce.Do(func() {
		_CLLocationUpdaterClass = CLLocationUpdaterClass{class: objc.GetClass("CLLocationUpdater")}
	})
	return _CLLocationUpdaterClass
}

// GetCLLocationUpdaterClass returns the class object for CLLocationUpdater.
func GetCLLocationUpdaterClass() CLLocationUpdaterClass {
	return getCLLocationUpdaterClass()
}

type CLLocationUpdaterClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CLLocationUpdaterClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CLLocationUpdaterClass) Alloc() CLLocationUpdater {
	rv := objc.Send[CLLocationUpdater](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// An object that provides device location updates.
//
// # Controlling the updater
//
//   - [CLLocationUpdater.Invalidate]: Invalidates the updater.
//   - [CLLocationUpdater.Pause]: Pauses the updater.
//   - [CLLocationUpdater.Resume]: Resumes the updater.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLLocationUpdater
type CLLocationUpdater struct {
	objectivec.Object
}

// CLLocationUpdaterFromID constructs a [CLLocationUpdater] from an objc.ID.
//
// An object that provides device location updates.
func CLLocationUpdaterFromID(id objc.ID) CLLocationUpdater {
	return CLLocationUpdater{objectivec.Object{ID: id}}
}

// Ensure CLLocationUpdater implements ICLLocationUpdater.
var _ ICLLocationUpdater = CLLocationUpdater{}

// An interface definition for the [CLLocationUpdater] class.
//
// # Controlling the updater
//
//   - [ICLLocationUpdater.Invalidate]: Invalidates the updater.
//   - [ICLLocationUpdater.Pause]: Pauses the updater.
//   - [ICLLocationUpdater.Resume]: Resumes the updater.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLLocationUpdater
type ICLLocationUpdater interface {
	objectivec.IObject

	// Topic: Controlling the updater

	// Invalidates the updater.
	Invalidate()
	// Pauses the updater.
	Pause()
	// Resumes the updater.
	Resume()
}

// Init initializes the instance.
func (l CLLocationUpdater) Init() CLLocationUpdater {
	rv := objc.Send[CLLocationUpdater](l.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (l CLLocationUpdater) Autorelease() CLLocationUpdater {
	rv := objc.Send[CLLocationUpdater](l.ID, objc.Sel("autorelease"))
	return rv
}

// NewCLLocationUpdater creates a new CLLocationUpdater instance.
func NewCLLocationUpdater() CLLocationUpdater {
	class := getCLLocationUpdaterClass()
	rv := objc.Send[CLLocationUpdater](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Invalidates the updater.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLLocationUpdater/invalidate
func (l CLLocationUpdater) Invalidate() {
	objc.Send[objc.ID](l.ID, objc.Sel("invalidate"))
}

// Pauses the updater.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLLocationUpdater/pause
func (l CLLocationUpdater) Pause() {
	objc.Send[objc.ID](l.ID, objc.Sel("pause"))
}

// Resumes the updater.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLLocationUpdater/resume
func (l CLLocationUpdater) Resume() {
	objc.Send[objc.ID](l.ID, objc.Sel("resume"))
}

// Creates a location updater with the configuration and queue that you
// specify.
//
// configuration: Specifies the live update configuration that the framework uses.
//
// queue: Specifies the queue to which the framework submits the handler with each
// available update.
//
// handler: The block that the framework invokes with each update.
//
// # Return Value
//
// Returns a location updater instance with the specified configuration,
// queue, and update handler.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLLocationUpdater/liveUpdaterWithConfiguration:queue:handler:
func (_CLLocationUpdaterClass CLLocationUpdaterClass) LiveUpdaterWithConfigurationQueueHandler(configuration CLLiveUpdateConfiguration, queue dispatch.Queue, handler CLUpdateHandler) CLLocationUpdater {
	_block2, _ := NewCLUpdateBlock(handler)
	rv := objc.Send[objc.ID](objc.ID(_CLLocationUpdaterClass.class), objc.Sel("liveUpdaterWithConfiguration:queue:handler:"), configuration, uintptr(queue.Handle()), _block2)
	return CLLocationUpdaterFromID(rv)
}

// Creates a location updater on the queue you specify.
//
// queue: Specifies the queue to which the framework submits the handler with each
// available update
//
// handler: The block that the framework invokes with each update.
//
// # Return Value
//
// Returns a location updater instance with the specified queue and update
// handler.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLLocationUpdater/liveUpdaterWithQueue:handler:
func (_CLLocationUpdaterClass CLLocationUpdaterClass) LiveUpdaterWithQueueHandler(queue dispatch.Queue, handler CLUpdateHandler) CLLocationUpdater {
	_block1, _ := NewCLUpdateBlock(handler)
	rv := objc.Send[objc.ID](objc.ID(_CLLocationUpdaterClass.class), objc.Sel("liveUpdaterWithQueue:handler:"), uintptr(queue.Handle()), _block1)
	return CLLocationUpdaterFromID(rv)
}

// LiveUpdaterWithConfigurationQueueHandlerSync is a synchronous wrapper around [CLLocationUpdater.LiveUpdaterWithConfigurationQueueHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (lc CLLocationUpdaterClass) LiveUpdaterWithConfigurationQueueHandlerSync(ctx context.Context, configuration CLLiveUpdateConfiguration, queue dispatch.Queue) (*CLUpdate, error) {
	done := make(chan *CLUpdate, 1)
	lc.LiveUpdaterWithConfigurationQueueHandler(configuration, queue, func(val *CLUpdate) {
		done <- val
	})
	select {
	case r := <-done:
		return r, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

// LiveUpdaterWithQueueHandlerSync is a synchronous wrapper around [CLLocationUpdater.LiveUpdaterWithQueueHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (lc CLLocationUpdaterClass) LiveUpdaterWithQueueHandlerSync(ctx context.Context, queue dispatch.Queue) (*CLUpdate, error) {
	done := make(chan *CLUpdate, 1)
	lc.LiveUpdaterWithQueueHandler(queue, func(val *CLUpdate) {
		done <- val
	})
	select {
	case r := <-done:
		return r, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}
