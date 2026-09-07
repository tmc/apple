// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [WSIndirectSystemGestureEventProcessor] class.
var (
	_WSIndirectSystemGestureEventProcessorClass     WSIndirectSystemGestureEventProcessorClass
	_WSIndirectSystemGestureEventProcessorClassOnce sync.Once
)

func getWSIndirectSystemGestureEventProcessorClass() WSIndirectSystemGestureEventProcessorClass {
	_WSIndirectSystemGestureEventProcessorClassOnce.Do(func() {
		_WSIndirectSystemGestureEventProcessorClass = WSIndirectSystemGestureEventProcessorClass{class: objc.GetClass("WSIndirectSystemGestureEventProcessor")}
	})
	return _WSIndirectSystemGestureEventProcessorClass
}

// GetWSIndirectSystemGestureEventProcessorClass returns the class object for WSIndirectSystemGestureEventProcessor.
func GetWSIndirectSystemGestureEventProcessorClass() WSIndirectSystemGestureEventProcessorClass {
	return getWSIndirectSystemGestureEventProcessorClass()
}

type WSIndirectSystemGestureEventProcessorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (wc WSIndirectSystemGestureEventProcessorClass) Class() objc.Class {
	return wc.class
}

// Alloc allocates memory for a new instance of the class.
func (wc WSIndirectSystemGestureEventProcessorClass) Alloc() WSIndirectSystemGestureEventProcessor {
	rv := objc.SendIfResponds[WSIndirectSystemGestureEventProcessor](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [WSIndirectSystemGestureEventProcessor.CancelAllActiveGestures]
//   - [WSIndirectSystemGestureEventProcessor.ProcessEventPassthroughsTarget]
//   - [WSIndirectSystemGestureEventProcessor.InitWithSystemGestureEventProcessor]
type WSIndirectSystemGestureEventProcessor struct {
	objectivec.Object
}

// WSIndirectSystemGestureEventProcessorFromID constructs a [WSIndirectSystemGestureEventProcessor] from an objc.ID.
func WSIndirectSystemGestureEventProcessorFromID(id objc.ID) WSIndirectSystemGestureEventProcessor {
	return WSIndirectSystemGestureEventProcessor{objectivec.Object{ID: id}}
}

// Ensure WSIndirectSystemGestureEventProcessor implements IWSIndirectSystemGestureEventProcessor.
var _ IWSIndirectSystemGestureEventProcessor = WSIndirectSystemGestureEventProcessor{}

// An interface definition for the [WSIndirectSystemGestureEventProcessor] class.
//
// # Methods
//
//   - [IWSIndirectSystemGestureEventProcessor.CancelAllActiveGestures]
//   - [IWSIndirectSystemGestureEventProcessor.ProcessEventPassthroughsTarget]
//   - [IWSIndirectSystemGestureEventProcessor.InitWithSystemGestureEventProcessor]
type IWSIndirectSystemGestureEventProcessor interface {
	objectivec.IObject

	// Topic: Methods

	CancelAllActiveGestures()
	ProcessEventPassthroughsTarget(event *SLSEventRecord, passthroughs objectivec.IObject, target unsafe.Pointer)
	InitWithSystemGestureEventProcessor(processor objectivec.IObject) WSIndirectSystemGestureEventProcessor
}

// Init initializes the instance.
func (w WSIndirectSystemGestureEventProcessor) Init() WSIndirectSystemGestureEventProcessor {
	rv := objc.SendIfResponds[WSIndirectSystemGestureEventProcessor](w.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (w WSIndirectSystemGestureEventProcessor) Autorelease() WSIndirectSystemGestureEventProcessor {
	rv := objc.SendIfResponds[WSIndirectSystemGestureEventProcessor](w.ID, objc.Sel("autorelease"))
	return rv
}

// NewWSIndirectSystemGestureEventProcessor creates a new WSIndirectSystemGestureEventProcessor instance.
func NewWSIndirectSystemGestureEventProcessor() WSIndirectSystemGestureEventProcessor {
	class := getWSIndirectSystemGestureEventProcessorClass()
	rv := objc.SendIfResponds[WSIndirectSystemGestureEventProcessor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewWSIndirectSystemGestureEventProcessorWithSystemGestureEventProcessor(processor objectivec.IObject) WSIndirectSystemGestureEventProcessor {
	instance := getWSIndirectSystemGestureEventProcessorClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithSystemGestureEventProcessor:"), processor)
	return WSIndirectSystemGestureEventProcessorFromID(rv)
}

func (w WSIndirectSystemGestureEventProcessor) CancelAllActiveGestures() {
	objc.SendIfResponds[objc.ID](w.ID, objc.Sel("cancelAllActiveGestures"))
}
func (w WSIndirectSystemGestureEventProcessor) ProcessEventPassthroughsTarget(event *SLSEventRecord, passthroughs objectivec.IObject, target unsafe.Pointer) {
	objc.SendIfResponds[objc.ID](w.ID, objc.Sel("processEvent:passthroughs:target:"), unsafe.Pointer(event), passthroughs, target)
}
func (w WSIndirectSystemGestureEventProcessor) InitWithSystemGestureEventProcessor(processor objectivec.IObject) WSIndirectSystemGestureEventProcessor {
	rv := objc.SendIfResponds[WSIndirectSystemGestureEventProcessor](w.ID, objc.Sel("initWithSystemGestureEventProcessor:"), processor)
	return rv
}
