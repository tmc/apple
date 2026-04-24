// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/objc"
)

// The class instance for the [WSTouchBarEventProcessor] class.
var (
	_WSTouchBarEventProcessorClass     WSTouchBarEventProcessorClass
	_WSTouchBarEventProcessorClassOnce sync.Once
)

func getWSTouchBarEventProcessorClass() WSTouchBarEventProcessorClass {
	_WSTouchBarEventProcessorClassOnce.Do(func() {
		_WSTouchBarEventProcessorClass = WSTouchBarEventProcessorClass{class: objc.GetClass("WSTouchBarEventProcessor")}
	})
	return _WSTouchBarEventProcessorClass
}

// GetWSTouchBarEventProcessorClass returns the class object for WSTouchBarEventProcessor.
func GetWSTouchBarEventProcessorClass() WSTouchBarEventProcessorClass {
	return getWSTouchBarEventProcessorClass()
}

type WSTouchBarEventProcessorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (wc WSTouchBarEventProcessorClass) Class() objc.Class {
	return wc.class
}

// Alloc allocates memory for a new instance of the class.
func (wc WSTouchBarEventProcessorClass) Alloc() WSTouchBarEventProcessor {
	rv := objc.Send[WSTouchBarEventProcessor](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/WSTouchBarEventProcessor
type WSTouchBarEventProcessor struct {
	WSEventProcessor
}

// WSTouchBarEventProcessorFromID constructs a [WSTouchBarEventProcessor] from an objc.ID.
func WSTouchBarEventProcessorFromID(id objc.ID) WSTouchBarEventProcessor {
	return WSTouchBarEventProcessor{WSEventProcessor: WSEventProcessorFromID(id)}
}

// Ensure WSTouchBarEventProcessor implements IWSTouchBarEventProcessor.
var _ IWSTouchBarEventProcessor = WSTouchBarEventProcessor{}

// An interface definition for the [WSTouchBarEventProcessor] class.
//
// See: https://developer.apple.com/documentation/SkyLight/WSTouchBarEventProcessor
type IWSTouchBarEventProcessor interface {
	IWSEventProcessor
}

// Init initializes the instance.
func (w WSTouchBarEventProcessor) Init() WSTouchBarEventProcessor {
	rv := objc.Send[WSTouchBarEventProcessor](w.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (w WSTouchBarEventProcessor) Autorelease() WSTouchBarEventProcessor {
	rv := objc.Send[WSTouchBarEventProcessor](w.ID, objc.Sel("autorelease"))
	return rv
}

// NewWSTouchBarEventProcessor creates a new WSTouchBarEventProcessor instance.
func NewWSTouchBarEventProcessor() WSTouchBarEventProcessor {
	class := getWSTouchBarEventProcessorClass()
	rv := objc.Send[WSTouchBarEventProcessor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/WSEventProcessor/initWithSession:
func NewWSTouchBarEventProcessorWithSession(session unsafe.Pointer) WSTouchBarEventProcessor {
	instance := getWSTouchBarEventProcessorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSession:"), session)
	return WSTouchBarEventProcessorFromID(rv)
}
