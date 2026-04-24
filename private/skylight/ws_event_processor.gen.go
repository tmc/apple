// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [WSEventProcessor] class.
var (
	_WSEventProcessorClass     WSEventProcessorClass
	_WSEventProcessorClassOnce sync.Once
)

func getWSEventProcessorClass() WSEventProcessorClass {
	_WSEventProcessorClassOnce.Do(func() {
		_WSEventProcessorClass = WSEventProcessorClass{class: objc.GetClass("WSEventProcessor")}
	})
	return _WSEventProcessorClass
}

// GetWSEventProcessorClass returns the class object for WSEventProcessor.
func GetWSEventProcessorClass() WSEventProcessorClass {
	return getWSEventProcessorClass()
}

type WSEventProcessorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (wc WSEventProcessorClass) Class() objc.Class {
	return wc.class
}

// Alloc allocates memory for a new instance of the class.
func (wc WSEventProcessorClass) Alloc() WSEventProcessor {
	rv := objc.Send[WSEventProcessor](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [WSEventProcessor.AnnotateAnnotationParams]
//   - [WSEventProcessor.Annotate_internalAnnotationParamsWindowConnEventRegionIDIsCapturedDefaultWindowOverrideCaptureCid]
//   - [WSEventProcessor.Can_handle]
//   - [WSEventProcessor.ClearEventState]
//   - [WSEventProcessor.CreateAnnotationParams]
//   - [WSEventProcessor.Event_dispatchAnnotationParamsDispatcher]
//   - [WSEventProcessor.Event_find_window]
//   - [WSEventProcessor.Post_event_annotateAnnotationParamsIsCapturedIsInkingEventAnnotateWindowAnnotateConnectionEventRegion]
//   - [WSEventProcessor.ProcessEventContextDispatcher]
//   - [WSEventProcessor.Route_annotate_eventAnnotationParamsDispatcher]
//   - [WSEventProcessor.InitWithSession]
//   - [WSEventProcessor.DebugDescription]
//   - [WSEventProcessor.Description]
//   - [WSEventProcessor.Hash]
//   - [WSEventProcessor.Superclass]
//
// See: https://developer.apple.com/documentation/SkyLight/WSEventProcessor
type WSEventProcessor struct {
	objectivec.Object
}

// WSEventProcessorFromID constructs a [WSEventProcessor] from an objc.ID.
func WSEventProcessorFromID(id objc.ID) WSEventProcessor {
	return WSEventProcessor{objectivec.Object{ID: id}}
}

// Ensure WSEventProcessor implements IWSEventProcessor.
var _ IWSEventProcessor = WSEventProcessor{}

// An interface definition for the [WSEventProcessor] class.
//
// # Methods
//
//   - [IWSEventProcessor.AnnotateAnnotationParams]
//   - [IWSEventProcessor.Annotate_internalAnnotationParamsWindowConnEventRegionIDIsCapturedDefaultWindowOverrideCaptureCid]
//   - [IWSEventProcessor.Can_handle]
//   - [IWSEventProcessor.ClearEventState]
//   - [IWSEventProcessor.CreateAnnotationParams]
//   - [IWSEventProcessor.Event_dispatchAnnotationParamsDispatcher]
//   - [IWSEventProcessor.Event_find_window]
//   - [IWSEventProcessor.Post_event_annotateAnnotationParamsIsCapturedIsInkingEventAnnotateWindowAnnotateConnectionEventRegion]
//   - [IWSEventProcessor.ProcessEventContextDispatcher]
//   - [IWSEventProcessor.Route_annotate_eventAnnotationParamsDispatcher]
//   - [IWSEventProcessor.InitWithSession]
//   - [IWSEventProcessor.DebugDescription]
//   - [IWSEventProcessor.Description]
//   - [IWSEventProcessor.Hash]
//   - [IWSEventProcessor.Superclass]
//
// See: https://developer.apple.com/documentation/SkyLight/WSEventProcessor
type IWSEventProcessor interface {
	objectivec.IObject

	// Topic: Methods

	AnnotateAnnotationParams(annotate *SLSEventRecordRef, params objectivec.IObject) int
	Annotate_internalAnnotationParamsWindowConnEventRegionIDIsCapturedDefaultWindowOverrideCaptureCid(annotate_internal *SLSEventRecordRef, params objectivec.IObject, conn uint32, id unsafe.Pointer, captured bool, window bool, cid unsafe.Pointer) int
	Can_handle(can_handle *SLSEventRecordRef) bool
	ClearEventState()
	CreateAnnotationParams(params *SLSEventRecordRef) objectivec.IObject
	Event_dispatchAnnotationParamsDispatcher(event_dispatch *SLSEventRecordRef, params objectivec.IObject, dispatcher objectivec.IObject)
	Event_find_window(event_find_window *SLSEventRecordRef)
	Post_event_annotateAnnotationParamsIsCapturedIsInkingEventAnnotateWindowAnnotateConnectionEventRegion(post_event_annotate *SLSEventRecordRef, params objectivec.IObject, captured bool, event bool, window unsafe.Pointer, connection unsafe.Pointer, region WSStructuralRegionRef) bool
	ProcessEventContextDispatcher(event *SLSEventRecordRef, context unsafe.Pointer, dispatcher objectivec.IObject) int64
	Route_annotate_eventAnnotationParamsDispatcher(route_annotate_event *SLSEventRecordRef, params objectivec.IObject, dispatcher objectivec.IObject)
	InitWithSession(session unsafe.Pointer) WSEventProcessor
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objc.Class
}

// Init initializes the instance.
func (w WSEventProcessor) Init() WSEventProcessor {
	rv := objc.Send[WSEventProcessor](w.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (w WSEventProcessor) Autorelease() WSEventProcessor {
	rv := objc.Send[WSEventProcessor](w.ID, objc.Sel("autorelease"))
	return rv
}

// NewWSEventProcessor creates a new WSEventProcessor instance.
func NewWSEventProcessor() WSEventProcessor {
	class := getWSEventProcessorClass()
	rv := objc.Send[WSEventProcessor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/WSEventProcessor/initWithSession:
func NewWSEventProcessorWithSession(session unsafe.Pointer) WSEventProcessor {
	instance := getWSEventProcessorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSession:"), session)
	return WSEventProcessorFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/WSEventProcessor/annotate:annotationParams:
func (w WSEventProcessor) AnnotateAnnotationParams(annotate *SLSEventRecordRef, params objectivec.IObject) int {
	rv := objc.Send[int](w.ID, objc.Sel("annotate:annotationParams:"), annotate, params)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/WSEventProcessor/annotate_internal:annotationParams:windowConn:eventRegionID:isCaptured:defaultWindow:overrideCaptureCid:
func (w WSEventProcessor) Annotate_internalAnnotationParamsWindowConnEventRegionIDIsCapturedDefaultWindowOverrideCaptureCid(annotate_internal *SLSEventRecordRef, params objectivec.IObject, conn uint32, id unsafe.Pointer, captured bool, window bool, cid unsafe.Pointer) int {
	rv := objc.Send[int](w.ID, objc.Sel("annotate_internal:annotationParams:windowConn:eventRegionID:isCaptured:defaultWindow:overrideCaptureCid:"), annotate_internal, params, conn, id, captured, window, cid)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/WSEventProcessor/can_handle:
func (w WSEventProcessor) Can_handle(can_handle *SLSEventRecordRef) bool {
	rv := objc.Send[bool](w.ID, objc.Sel("can_handle:"), can_handle)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/WSEventProcessor/clearEventState
func (w WSEventProcessor) ClearEventState() {
	objc.Send[objc.ID](w.ID, objc.Sel("clearEventState"))
}

// See: https://developer.apple.com/documentation/SkyLight/WSEventProcessor/createAnnotationParams:
func (w WSEventProcessor) CreateAnnotationParams(params *SLSEventRecordRef) objectivec.IObject {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("createAnnotationParams:"), params)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/WSEventProcessor/event_dispatch:annotationParams:dispatcher:
func (w WSEventProcessor) Event_dispatchAnnotationParamsDispatcher(event_dispatch *SLSEventRecordRef, params objectivec.IObject, dispatcher objectivec.IObject) {
	objc.Send[objc.ID](w.ID, objc.Sel("event_dispatch:annotationParams:dispatcher:"), event_dispatch, params, dispatcher)
}

// See: https://developer.apple.com/documentation/SkyLight/WSEventProcessor/event_find_window:
func (w WSEventProcessor) Event_find_window(event_find_window *SLSEventRecordRef) {
	objc.Send[objc.ID](w.ID, objc.Sel("event_find_window:"), event_find_window)
}

// See: https://developer.apple.com/documentation/SkyLight/WSEventProcessor/post_event_annotate:annotationParams:isCaptured:isInkingEvent:annotateWindow:annotateConnection:eventRegion:
func (w WSEventProcessor) Post_event_annotateAnnotationParamsIsCapturedIsInkingEventAnnotateWindowAnnotateConnectionEventRegion(post_event_annotate *SLSEventRecordRef, params objectivec.IObject, captured bool, event bool, window unsafe.Pointer, connection unsafe.Pointer, region WSStructuralRegionRef) bool {
	rv := objc.Send[bool](w.ID, objc.Sel("post_event_annotate:annotationParams:isCaptured:isInkingEvent:annotateWindow:annotateConnection:eventRegion:"), post_event_annotate, params, captured, event, window, connection, region)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/WSEventProcessor/processEvent:context:dispatcher:
func (w WSEventProcessor) ProcessEventContextDispatcher(event *SLSEventRecordRef, context unsafe.Pointer, dispatcher objectivec.IObject) int64 {
	rv := objc.Send[int64](w.ID, objc.Sel("processEvent:context:dispatcher:"), event, context, dispatcher)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/WSEventProcessor/route_annotate_event:annotationParams:dispatcher:
func (w WSEventProcessor) Route_annotate_eventAnnotationParamsDispatcher(route_annotate_event *SLSEventRecordRef, params objectivec.IObject, dispatcher objectivec.IObject) {
	objc.Send[objc.ID](w.ID, objc.Sel("route_annotate_event:annotationParams:dispatcher:"), route_annotate_event, params, dispatcher)
}

// See: https://developer.apple.com/documentation/SkyLight/WSEventProcessor/initWithSession:
func (w WSEventProcessor) InitWithSession(session unsafe.Pointer) WSEventProcessor {
	rv := objc.Send[WSEventProcessor](w.ID, objc.Sel("initWithSession:"), session)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/WSEventProcessor/debugDescription
func (w WSEventProcessor) DebugDescription() string {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/SkyLight/WSEventProcessor/description
func (w WSEventProcessor) Description() string {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/SkyLight/WSEventProcessor/hash
func (w WSEventProcessor) Hash() uint64 {
	rv := objc.Send[uint64](w.ID, objc.Sel("hash"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/WSEventProcessor/superclass
func (w WSEventProcessor) Superclass() objc.Class {
	rv := objc.Send[objc.Class](w.ID, objc.Sel("superclass"))
	return rv
}
