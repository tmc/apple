// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/objc"
)

// The class instance for the [WSSidecar2EventProcessor] class.
var (
	_WSSidecar2EventProcessorClass     WSSidecar2EventProcessorClass
	_WSSidecar2EventProcessorClassOnce sync.Once
)

func getWSSidecar2EventProcessorClass() WSSidecar2EventProcessorClass {
	_WSSidecar2EventProcessorClassOnce.Do(func() {
		_WSSidecar2EventProcessorClass = WSSidecar2EventProcessorClass{class: objc.GetClass("WSSidecar2EventProcessor")}
	})
	return _WSSidecar2EventProcessorClass
}

// GetWSSidecar2EventProcessorClass returns the class object for WSSidecar2EventProcessor.
func GetWSSidecar2EventProcessorClass() WSSidecar2EventProcessorClass {
	return getWSSidecar2EventProcessorClass()
}

type WSSidecar2EventProcessorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (wc WSSidecar2EventProcessorClass) Class() objc.Class {
	return wc.class
}

// Alloc allocates memory for a new instance of the class.
func (wc WSSidecar2EventProcessorClass) Alloc() WSSidecar2EventProcessor {
	rv := objc.SendIfResponds[WSSidecar2EventProcessor](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

type WSSidecar2EventProcessor struct {
	WSGestureEventProcessor
}

// WSSidecar2EventProcessorFromID constructs a [WSSidecar2EventProcessor] from an objc.ID.
func WSSidecar2EventProcessorFromID(id objc.ID) WSSidecar2EventProcessor {
	return WSSidecar2EventProcessor{WSGestureEventProcessor: WSGestureEventProcessorFromID(id)}
}

// Ensure WSSidecar2EventProcessor implements IWSSidecar2EventProcessor.
var _ IWSSidecar2EventProcessor = WSSidecar2EventProcessor{}

// An interface definition for the [WSSidecar2EventProcessor] class.
type IWSSidecar2EventProcessor interface {
	IWSGestureEventProcessor
}

// Init initializes the instance.
func (w WSSidecar2EventProcessor) Init() WSSidecar2EventProcessor {
	rv := objc.SendIfResponds[WSSidecar2EventProcessor](w.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (w WSSidecar2EventProcessor) Autorelease() WSSidecar2EventProcessor {
	rv := objc.SendIfResponds[WSSidecar2EventProcessor](w.ID, objc.Sel("autorelease"))
	return rv
}

// NewWSSidecar2EventProcessor creates a new WSSidecar2EventProcessor instance.
func NewWSSidecar2EventProcessor() WSSidecar2EventProcessor {
	class := getWSSidecar2EventProcessorClass()
	rv := objc.SendIfResponds[WSSidecar2EventProcessor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewWSSidecar2EventProcessorWithSession(session *CGXSession) WSSidecar2EventProcessor {
	instance := getWSSidecar2EventProcessorClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithSession:"), unsafe.Pointer(session))
	return WSSidecar2EventProcessorFromID(rv)
}
