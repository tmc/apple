// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/objc"
)

// The class instance for the [WSSystemDefinedEventProcessor] class.
var (
	_WSSystemDefinedEventProcessorClass     WSSystemDefinedEventProcessorClass
	_WSSystemDefinedEventProcessorClassOnce sync.Once
)

func getWSSystemDefinedEventProcessorClass() WSSystemDefinedEventProcessorClass {
	_WSSystemDefinedEventProcessorClassOnce.Do(func() {
		_WSSystemDefinedEventProcessorClass = WSSystemDefinedEventProcessorClass{class: objc.GetClass("WSSystemDefinedEventProcessor")}
	})
	return _WSSystemDefinedEventProcessorClass
}

// GetWSSystemDefinedEventProcessorClass returns the class object for WSSystemDefinedEventProcessor.
func GetWSSystemDefinedEventProcessorClass() WSSystemDefinedEventProcessorClass {
	return getWSSystemDefinedEventProcessorClass()
}

type WSSystemDefinedEventProcessorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (wc WSSystemDefinedEventProcessorClass) Class() objc.Class {
	return wc.class
}

// Alloc allocates memory for a new instance of the class.
func (wc WSSystemDefinedEventProcessorClass) Alloc() WSSystemDefinedEventProcessor {
	rv := objc.Send[WSSystemDefinedEventProcessor](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/WSSystemDefinedEventProcessor
type WSSystemDefinedEventProcessor struct {
	WSEventProcessor
}

// WSSystemDefinedEventProcessorFromID constructs a [WSSystemDefinedEventProcessor] from an objc.ID.
func WSSystemDefinedEventProcessorFromID(id objc.ID) WSSystemDefinedEventProcessor {
	return WSSystemDefinedEventProcessor{WSEventProcessor: WSEventProcessorFromID(id)}
}

// Ensure WSSystemDefinedEventProcessor implements IWSSystemDefinedEventProcessor.
var _ IWSSystemDefinedEventProcessor = WSSystemDefinedEventProcessor{}

// An interface definition for the [WSSystemDefinedEventProcessor] class.
//
// See: https://developer.apple.com/documentation/SkyLight/WSSystemDefinedEventProcessor
type IWSSystemDefinedEventProcessor interface {
	IWSEventProcessor
}

// Init initializes the instance.
func (w WSSystemDefinedEventProcessor) Init() WSSystemDefinedEventProcessor {
	rv := objc.Send[WSSystemDefinedEventProcessor](w.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (w WSSystemDefinedEventProcessor) Autorelease() WSSystemDefinedEventProcessor {
	rv := objc.Send[WSSystemDefinedEventProcessor](w.ID, objc.Sel("autorelease"))
	return rv
}

// NewWSSystemDefinedEventProcessor creates a new WSSystemDefinedEventProcessor instance.
func NewWSSystemDefinedEventProcessor() WSSystemDefinedEventProcessor {
	class := getWSSystemDefinedEventProcessorClass()
	rv := objc.Send[WSSystemDefinedEventProcessor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/WSEventProcessor/initWithSession:
func NewWSSystemDefinedEventProcessorWithSession(session unsafe.Pointer) WSSystemDefinedEventProcessor {
	instance := getWSSystemDefinedEventProcessorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSession:"), session)
	return WSSystemDefinedEventProcessorFromID(rv)
}
