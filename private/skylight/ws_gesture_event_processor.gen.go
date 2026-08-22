// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [WSGestureEventProcessor] class.
var (
	_WSGestureEventProcessorClass     WSGestureEventProcessorClass
	_WSGestureEventProcessorClassOnce sync.Once
)

func getWSGestureEventProcessorClass() WSGestureEventProcessorClass {
	_WSGestureEventProcessorClassOnce.Do(func() {
		_WSGestureEventProcessorClass = WSGestureEventProcessorClass{class: objc.GetClass("WSGestureEventProcessor")}
	})
	return _WSGestureEventProcessorClass
}

// GetWSGestureEventProcessorClass returns the class object for WSGestureEventProcessor.
func GetWSGestureEventProcessorClass() WSGestureEventProcessorClass {
	return getWSGestureEventProcessorClass()
}

type WSGestureEventProcessorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (wc WSGestureEventProcessorClass) Class() objc.Class {
	return wc.class
}

// Alloc allocates memory for a new instance of the class.
func (wc WSGestureEventProcessorClass) Alloc() WSGestureEventProcessor {
	rv := objc.SendIfResponds[WSGestureEventProcessor](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

type WSGestureEventProcessor struct {
	WSEventProcessor
}

// WSGestureEventProcessorFromID constructs a [WSGestureEventProcessor] from an objc.ID.
func WSGestureEventProcessorFromID(id objc.ID) WSGestureEventProcessor {
	return WSGestureEventProcessor{WSEventProcessor: WSEventProcessorFromID(id)}
}

// Ensure WSGestureEventProcessor implements IWSGestureEventProcessor.
var _ IWSGestureEventProcessor = WSGestureEventProcessor{}

// An interface definition for the [WSGestureEventProcessor] class.
type IWSGestureEventProcessor interface {
	IWSEventProcessor
}

// Init initializes the instance.
func (w WSGestureEventProcessor) Init() WSGestureEventProcessor {
	rv := objc.SendIfResponds[WSGestureEventProcessor](w.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (w WSGestureEventProcessor) Autorelease() WSGestureEventProcessor {
	rv := objc.SendIfResponds[WSGestureEventProcessor](w.ID, objc.Sel("autorelease"))
	return rv
}

// NewWSGestureEventProcessor creates a new WSGestureEventProcessor instance.
func NewWSGestureEventProcessor() WSGestureEventProcessor {
	class := getWSGestureEventProcessorClass()
	rv := objc.SendIfResponds[WSGestureEventProcessor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewWSGestureEventProcessorWithSession(session *CGXSession) WSGestureEventProcessor {
	instance := getWSGestureEventProcessorClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithSession:"), unsafe.Pointer(session))
	return WSGestureEventProcessorFromID(rv)
}

func (_WSGestureEventProcessorClass WSGestureEventProcessorClass) Annotate_scroll_zoom_eventWindowConnEventRegionIDIsCapturedAnnotationParams(annotate_scroll_zoom_event *SLSEventRecord, conn uint32, id *uint64, captured bool, params objectivec.IObject) int {
	rv := objc.SendIfResponds[int](objc.ID(_WSGestureEventProcessorClass.class), objc.Sel("annotate_scroll_zoom_event:windowConn:eventRegionID:isCaptured:annotationParams:"), unsafe.Pointer(annotate_scroll_zoom_event), conn, unsafe.Pointer(id), captured, params)
	return rv
}
