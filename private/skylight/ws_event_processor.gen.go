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
	rv := objc.SendIfResponds[WSEventProcessor](objc.ID(wc.class), objc.Sel("alloc"))
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
type IWSEventProcessor interface {
	objectivec.IObject

	// Topic: Methods

	AnnotateAnnotationParams(annotate *SLSEventRecord, params objectivec.IObject) int
	Annotate_internalAnnotationParamsWindowConnEventRegionIDIsCapturedDefaultWindowOverrideCaptureCid(annotate_internal *SLSEventRecord, params objectivec.IObject, conn uint32, id *uint64, captured bool, window bool, cid *bool) int
	Can_handle(can_handle *SLSEventRecord) bool
	ClearEventState()
	CreateAnnotationParams(params *SLSEventRecord) objectivec.IObject
	Event_dispatchAnnotationParamsDispatcher(event_dispatch *SLSEventRecord, params objectivec.IObject, dispatcher objectivec.IObject)
	Event_find_window(event_find_window *SLSEventRecord)
	Post_event_annotateAnnotationParamsIsCapturedIsInkingEventAnnotateWindowAnnotateConnectionEventRegion(post_event_annotate *SLSEventRecord, params objectivec.IObject, captured bool, event bool, window unsafe.Pointer, connection *CGXConnection, region WSStructuralRegionRef) bool
	ProcessEventContextDispatcher(event *SLSEventRecord, context *CPXEventProcessorContext, dispatcher objectivec.IObject) int64
	Route_annotate_eventAnnotationParamsDispatcher(route_annotate_event *SLSEventRecord, params objectivec.IObject, dispatcher objectivec.IObject)
	InitWithSession(session *CGXSession) WSEventProcessor
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objectivec.Class
}

// Init initializes the instance.
func (w WSEventProcessor) Init() WSEventProcessor {
	rv := objc.SendIfResponds[WSEventProcessor](w.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (w WSEventProcessor) Autorelease() WSEventProcessor {
	rv := objc.SendIfResponds[WSEventProcessor](w.ID, objc.Sel("autorelease"))
	return rv
}

// NewWSEventProcessor creates a new WSEventProcessor instance.
func NewWSEventProcessor() WSEventProcessor {
	class := getWSEventProcessorClass()
	rv := objc.SendIfResponds[WSEventProcessor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewWSEventProcessorWithSession(session *CGXSession) WSEventProcessor {
	instance := getWSEventProcessorClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithSession:"), unsafe.Pointer(session))
	return WSEventProcessorFromID(rv)
}

func (w WSEventProcessor) AnnotateAnnotationParams(annotate *SLSEventRecord, params objectivec.IObject) int {
	rv := objc.SendIfResponds[int](w.ID, objc.Sel("annotate:annotationParams:"), unsafe.Pointer(annotate), params)
	return rv
}
func (w WSEventProcessor) Annotate_internalAnnotationParamsWindowConnEventRegionIDIsCapturedDefaultWindowOverrideCaptureCid(annotate_internal *SLSEventRecord, params objectivec.IObject, conn uint32, id *uint64, captured bool, window bool, cid *bool) int {
	rv := objc.SendIfResponds[int](w.ID, objc.Sel("annotate_internal:annotationParams:windowConn:eventRegionID:isCaptured:defaultWindow:overrideCaptureCid:"), unsafe.Pointer(annotate_internal), params, conn, unsafe.Pointer(id), captured, window, cid)
	return rv
}
func (w WSEventProcessor) Can_handle(can_handle *SLSEventRecord) bool {
	rv := objc.SendIfResponds[bool](w.ID, objc.Sel("can_handle:"), unsafe.Pointer(can_handle))
	return rv
}
func (w WSEventProcessor) ClearEventState() {
	objc.SendIfResponds[objc.ID](w.ID, objc.Sel("clearEventState"))
}
func (w WSEventProcessor) CreateAnnotationParams(params *SLSEventRecord) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](w.ID, objc.Sel("createAnnotationParams:"), unsafe.Pointer(params))
	return objectivec.Object{ID: rv}
}
func (w WSEventProcessor) Event_dispatchAnnotationParamsDispatcher(event_dispatch *SLSEventRecord, params objectivec.IObject, dispatcher objectivec.IObject) {
	objc.SendIfResponds[objc.ID](w.ID, objc.Sel("event_dispatch:annotationParams:dispatcher:"), unsafe.Pointer(event_dispatch), params, dispatcher)
}
func (w WSEventProcessor) Event_find_window(event_find_window *SLSEventRecord) {
	objc.SendIfResponds[objc.ID](w.ID, objc.Sel("event_find_window:"), unsafe.Pointer(event_find_window))
}
func (w WSEventProcessor) Post_event_annotateAnnotationParamsIsCapturedIsInkingEventAnnotateWindowAnnotateConnectionEventRegion(post_event_annotate *SLSEventRecord, params objectivec.IObject, captured bool, event bool, window unsafe.Pointer, connection *CGXConnection, region WSStructuralRegionRef) bool {
	rv := objc.SendIfResponds[bool](w.ID, objc.Sel("post_event_annotate:annotationParams:isCaptured:isInkingEvent:annotateWindow:annotateConnection:eventRegion:"), unsafe.Pointer(post_event_annotate), params, captured, event, window, unsafe.Pointer(connection), region)
	return rv
}
func (w WSEventProcessor) ProcessEventContextDispatcher(event *SLSEventRecord, context *CPXEventProcessorContext, dispatcher objectivec.IObject) int64 {
	rv := objc.SendIfResponds[int64](w.ID, objc.Sel("processEvent:context:dispatcher:"), unsafe.Pointer(event), unsafe.Pointer(context), dispatcher)
	return rv
}
func (w WSEventProcessor) Route_annotate_eventAnnotationParamsDispatcher(route_annotate_event *SLSEventRecord, params objectivec.IObject, dispatcher objectivec.IObject) {
	objc.SendIfResponds[objc.ID](w.ID, objc.Sel("route_annotate_event:annotationParams:dispatcher:"), unsafe.Pointer(route_annotate_event), params, dispatcher)
}
func (w WSEventProcessor) InitWithSession(session *CGXSession) WSEventProcessor {
	rv := objc.SendIfResponds[WSEventProcessor](w.ID, objc.Sel("initWithSession:"), unsafe.Pointer(session))
	return rv
}

func (w WSEventProcessor) DebugDescription() string {
	rv := objc.SendIfResponds[objc.ID](w.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
func (w WSEventProcessor) Description() string {
	rv := objc.SendIfResponds[objc.ID](w.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
func (w WSEventProcessor) Hash() uint64 {
	rv := objc.SendIfResponds[uint64](w.ID, objc.Sel("hash"))
	return rv
}
func (w WSEventProcessor) Superclass() objectivec.Class {
	rv := objc.SendIfResponds[objectivec.Class](w.ID, objc.Sel("superclass"))
	return objectivec.Class(rv)
}
