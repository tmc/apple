// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/objc"
)

// The class instance for the [WSMouseEventProcessor] class.
var (
	_WSMouseEventProcessorClass     WSMouseEventProcessorClass
	_WSMouseEventProcessorClassOnce sync.Once
)

func getWSMouseEventProcessorClass() WSMouseEventProcessorClass {
	_WSMouseEventProcessorClassOnce.Do(func() {
		_WSMouseEventProcessorClass = WSMouseEventProcessorClass{class: objc.GetClass("WSMouseEventProcessor")}
	})
	return _WSMouseEventProcessorClass
}

// GetWSMouseEventProcessorClass returns the class object for WSMouseEventProcessor.
func GetWSMouseEventProcessorClass() WSMouseEventProcessorClass {
	return getWSMouseEventProcessorClass()
}

type WSMouseEventProcessorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (wc WSMouseEventProcessorClass) Class() objc.Class {
	return wc.class
}

// Alloc allocates memory for a new instance of the class.
func (wc WSMouseEventProcessorClass) Alloc() WSMouseEventProcessor {
	rv := objc.Send[WSMouseEventProcessor](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/WSMouseEventProcessor
type WSMouseEventProcessor struct {
	WSEventProcessor
}

// WSMouseEventProcessorFromID constructs a [WSMouseEventProcessor] from an objc.ID.
func WSMouseEventProcessorFromID(id objc.ID) WSMouseEventProcessor {
	return WSMouseEventProcessor{WSEventProcessor: WSEventProcessorFromID(id)}
}

// Ensure WSMouseEventProcessor implements IWSMouseEventProcessor.
var _ IWSMouseEventProcessor = WSMouseEventProcessor{}

// An interface definition for the [WSMouseEventProcessor] class.
//
// See: https://developer.apple.com/documentation/SkyLight/WSMouseEventProcessor
type IWSMouseEventProcessor interface {
	IWSEventProcessor
}

// Init initializes the instance.
func (w WSMouseEventProcessor) Init() WSMouseEventProcessor {
	rv := objc.Send[WSMouseEventProcessor](w.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (w WSMouseEventProcessor) Autorelease() WSMouseEventProcessor {
	rv := objc.Send[WSMouseEventProcessor](w.ID, objc.Sel("autorelease"))
	return rv
}

// NewWSMouseEventProcessor creates a new WSMouseEventProcessor instance.
func NewWSMouseEventProcessor() WSMouseEventProcessor {
	class := getWSMouseEventProcessorClass()
	rv := objc.Send[WSMouseEventProcessor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/WSEventProcessor/initWithSession:
func NewWSMouseEventProcessorWithSession(session unsafe.Pointer) WSMouseEventProcessor {
	instance := getWSMouseEventProcessorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSession:"), session)
	return WSMouseEventProcessorFromID(rv)
}
