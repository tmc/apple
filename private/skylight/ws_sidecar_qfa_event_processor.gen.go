// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [WSSidecarQFAEventProcessor] class.
var (
	_WSSidecarQFAEventProcessorClass     WSSidecarQFAEventProcessorClass
	_WSSidecarQFAEventProcessorClassOnce sync.Once
)

func getWSSidecarQFAEventProcessorClass() WSSidecarQFAEventProcessorClass {
	_WSSidecarQFAEventProcessorClassOnce.Do(func() {
		_WSSidecarQFAEventProcessorClass = WSSidecarQFAEventProcessorClass{class: objc.GetClass("WSSidecarQFAEventProcessor")}
	})
	return _WSSidecarQFAEventProcessorClass
}

// GetWSSidecarQFAEventProcessorClass returns the class object for WSSidecarQFAEventProcessor.
func GetWSSidecarQFAEventProcessorClass() WSSidecarQFAEventProcessorClass {
	return getWSSidecarQFAEventProcessorClass()
}

type WSSidecarQFAEventProcessorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (wc WSSidecarQFAEventProcessorClass) Class() objc.Class {
	return wc.class
}

// Alloc allocates memory for a new instance of the class.
func (wc WSSidecarQFAEventProcessorClass) Alloc() WSSidecarQFAEventProcessor {
	rv := objc.SendIfResponds[WSSidecarQFAEventProcessor](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

type WSSidecarQFAEventProcessor struct {
	WSGestureEventProcessor
}

// WSSidecarQFAEventProcessorFromID constructs a [WSSidecarQFAEventProcessor] from an objc.ID.
func WSSidecarQFAEventProcessorFromID(id objc.ID) WSSidecarQFAEventProcessor {
	return WSSidecarQFAEventProcessor{WSGestureEventProcessor: WSGestureEventProcessorFromID(id)}
}

// Ensure WSSidecarQFAEventProcessor implements IWSSidecarQFAEventProcessor.
var _ IWSSidecarQFAEventProcessor = WSSidecarQFAEventProcessor{}

// An interface definition for the [WSSidecarQFAEventProcessor] class.
type IWSSidecarQFAEventProcessor interface {
	IWSGestureEventProcessor
}

// Init initializes the instance.
func (w WSSidecarQFAEventProcessor) Init() WSSidecarQFAEventProcessor {
	rv := objc.SendIfResponds[WSSidecarQFAEventProcessor](w.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (w WSSidecarQFAEventProcessor) Autorelease() WSSidecarQFAEventProcessor {
	rv := objc.SendIfResponds[WSSidecarQFAEventProcessor](w.ID, objc.Sel("autorelease"))
	return rv
}

// NewWSSidecarQFAEventProcessor creates a new WSSidecarQFAEventProcessor instance.
func NewWSSidecarQFAEventProcessor() WSSidecarQFAEventProcessor {
	class := getWSSidecarQFAEventProcessorClass()
	rv := objc.SendIfResponds[WSSidecarQFAEventProcessor](objc.ID(class.class), objc.Sel("new"))
	return rv
}
