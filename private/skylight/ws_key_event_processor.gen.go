// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/objc"
)

// The class instance for the [WSKeyEventProcessor] class.
var (
	_WSKeyEventProcessorClass     WSKeyEventProcessorClass
	_WSKeyEventProcessorClassOnce sync.Once
)

func getWSKeyEventProcessorClass() WSKeyEventProcessorClass {
	_WSKeyEventProcessorClassOnce.Do(func() {
		_WSKeyEventProcessorClass = WSKeyEventProcessorClass{class: objc.GetClass("WSKeyEventProcessor")}
	})
	return _WSKeyEventProcessorClass
}

// GetWSKeyEventProcessorClass returns the class object for WSKeyEventProcessor.
func GetWSKeyEventProcessorClass() WSKeyEventProcessorClass {
	return getWSKeyEventProcessorClass()
}

type WSKeyEventProcessorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (wc WSKeyEventProcessorClass) Class() objc.Class {
	return wc.class
}

// Alloc allocates memory for a new instance of the class.
func (wc WSKeyEventProcessorClass) Alloc() WSKeyEventProcessor {
	rv := objc.Send[WSKeyEventProcessor](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/WSKeyEventProcessor
type WSKeyEventProcessor struct {
	WSEventProcessor
}

// WSKeyEventProcessorFromID constructs a [WSKeyEventProcessor] from an objc.ID.
func WSKeyEventProcessorFromID(id objc.ID) WSKeyEventProcessor {
	return WSKeyEventProcessor{WSEventProcessor: WSEventProcessorFromID(id)}
}

// Ensure WSKeyEventProcessor implements IWSKeyEventProcessor.
var _ IWSKeyEventProcessor = WSKeyEventProcessor{}

// An interface definition for the [WSKeyEventProcessor] class.
//
// See: https://developer.apple.com/documentation/SkyLight/WSKeyEventProcessor
type IWSKeyEventProcessor interface {
	IWSEventProcessor
}

// Init initializes the instance.
func (w WSKeyEventProcessor) Init() WSKeyEventProcessor {
	rv := objc.Send[WSKeyEventProcessor](w.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (w WSKeyEventProcessor) Autorelease() WSKeyEventProcessor {
	rv := objc.Send[WSKeyEventProcessor](w.ID, objc.Sel("autorelease"))
	return rv
}

// NewWSKeyEventProcessor creates a new WSKeyEventProcessor instance.
func NewWSKeyEventProcessor() WSKeyEventProcessor {
	class := getWSKeyEventProcessorClass()
	rv := objc.Send[WSKeyEventProcessor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/WSEventProcessor/initWithSession:
func NewWSKeyEventProcessorWithSession(session unsafe.Pointer) WSKeyEventProcessor {
	instance := getWSKeyEventProcessorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSession:"), session)
	return WSKeyEventProcessorFromID(rv)
}
