// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [WSLegacyEventProcessor] class.
var (
	_WSLegacyEventProcessorClass     WSLegacyEventProcessorClass
	_WSLegacyEventProcessorClassOnce sync.Once
)

func getWSLegacyEventProcessorClass() WSLegacyEventProcessorClass {
	_WSLegacyEventProcessorClassOnce.Do(func() {
		_WSLegacyEventProcessorClass = WSLegacyEventProcessorClass{class: objc.GetClass("WSLegacyEventProcessor")}
	})
	return _WSLegacyEventProcessorClass
}

// GetWSLegacyEventProcessorClass returns the class object for WSLegacyEventProcessor.
func GetWSLegacyEventProcessorClass() WSLegacyEventProcessorClass {
	return getWSLegacyEventProcessorClass()
}

type WSLegacyEventProcessorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (wc WSLegacyEventProcessorClass) Class() objc.Class {
	return wc.class
}

// Alloc allocates memory for a new instance of the class.
func (wc WSLegacyEventProcessorClass) Alloc() WSLegacyEventProcessor {
	rv := objc.SendIfResponds[WSLegacyEventProcessor](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [WSLegacyEventProcessor.AnnotateAnnotationParams]
//   - [WSLegacyEventProcessor.Annotate_internalAnnotationParamsWindowConnEventRegionIDIsCapturedDefaultWindowOverrideCaptureCid]
//   - [WSLegacyEventProcessor.Can_handle]
//   - [WSLegacyEventProcessor.ClearEventState]
//   - [WSLegacyEventProcessor.CreateAnnotationParams]
//   - [WSLegacyEventProcessor.Event_dispatchAnnotationParamsDispatcher]
//   - [WSLegacyEventProcessor.Event_find_windowAnnotationParams]
//   - [WSLegacyEventProcessor.Post_event_annotateAnnotationParamsIsCapturedAnnotateWindowAnnotateConnectionEventRegion]
//   - [WSLegacyEventProcessor.ProcessEventDispatcher]
//   - [WSLegacyEventProcessor.Route_annotate_eventAnnotationParamsDispatcher]
//   - [WSLegacyEventProcessor.DebugDescription]
//   - [WSLegacyEventProcessor.Description]
//   - [WSLegacyEventProcessor.Hash]
//   - [WSLegacyEventProcessor.Superclass]
type WSLegacyEventProcessor struct {
	objectivec.Object
}

// WSLegacyEventProcessorFromID constructs a [WSLegacyEventProcessor] from an objc.ID.
func WSLegacyEventProcessorFromID(id objc.ID) WSLegacyEventProcessor {
	return WSLegacyEventProcessor{objectivec.Object{ID: id}}
}

// Ensure WSLegacyEventProcessor implements IWSLegacyEventProcessor.
var _ IWSLegacyEventProcessor = WSLegacyEventProcessor{}

// An interface definition for the [WSLegacyEventProcessor] class.
//
// # Methods
//
//   - [IWSLegacyEventProcessor.AnnotateAnnotationParams]
//   - [IWSLegacyEventProcessor.Annotate_internalAnnotationParamsWindowConnEventRegionIDIsCapturedDefaultWindowOverrideCaptureCid]
//   - [IWSLegacyEventProcessor.Can_handle]
//   - [IWSLegacyEventProcessor.ClearEventState]
//   - [IWSLegacyEventProcessor.CreateAnnotationParams]
//   - [IWSLegacyEventProcessor.Event_dispatchAnnotationParamsDispatcher]
//   - [IWSLegacyEventProcessor.Event_find_windowAnnotationParams]
//   - [IWSLegacyEventProcessor.Post_event_annotateAnnotationParamsIsCapturedAnnotateWindowAnnotateConnectionEventRegion]
//   - [IWSLegacyEventProcessor.ProcessEventDispatcher]
//   - [IWSLegacyEventProcessor.Route_annotate_eventAnnotationParamsDispatcher]
//   - [IWSLegacyEventProcessor.DebugDescription]
//   - [IWSLegacyEventProcessor.Description]
//   - [IWSLegacyEventProcessor.Hash]
//   - [IWSLegacyEventProcessor.Superclass]
type IWSLegacyEventProcessor interface {
	objectivec.IObject

	// Topic: Methods

	AnnotateAnnotationParams(annotate *SLSEventRecord, params objectivec.IObject) int32
	Annotate_internalAnnotationParamsWindowConnEventRegionIDIsCapturedDefaultWindowOverrideCaptureCid(annotate_internal *SLSEventRecord, params objectivec.IObject, conn uint32, id *uint64, captured bool, window bool, cid *bool) int32
	Can_handle(can_handle *SLSEventRecord) bool
	ClearEventState()
	CreateAnnotationParams(params *SLSEventRecord) objectivec.IObject
	Event_dispatchAnnotationParamsDispatcher(event_dispatch *SLSEventRecord, params objectivec.IObject, dispatcher objectivec.IObject)
	Event_find_windowAnnotationParams(event_find_window *SLSEventRecord, params objectivec.IObject)
	Post_event_annotateAnnotationParamsIsCapturedAnnotateWindowAnnotateConnectionEventRegion(post_event_annotate *SLSEventRecord, params objectivec.IObject, captured bool, window unsafe.Pointer, connection *CGXConnection, region WSStructuralRegionRef) bool
	ProcessEventDispatcher(event *SLSEventRecord, dispatcher objectivec.IObject) int64
	Route_annotate_eventAnnotationParamsDispatcher(route_annotate_event *SLSEventRecord, params objectivec.IObject, dispatcher objectivec.IObject)
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objectivec.Class
}

// Init initializes the instance.
func (w WSLegacyEventProcessor) Init() WSLegacyEventProcessor {
	rv := objc.SendIfResponds[WSLegacyEventProcessor](w.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (w WSLegacyEventProcessor) Autorelease() WSLegacyEventProcessor {
	rv := objc.SendIfResponds[WSLegacyEventProcessor](w.ID, objc.Sel("autorelease"))
	return rv
}

// NewWSLegacyEventProcessor creates a new WSLegacyEventProcessor instance.
func NewWSLegacyEventProcessor() WSLegacyEventProcessor {
	class := getWSLegacyEventProcessorClass()
	rv := objc.SendIfResponds[WSLegacyEventProcessor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (w WSLegacyEventProcessor) AnnotateAnnotationParams(annotate *SLSEventRecord, params objectivec.IObject) int32 {
	rv := objc.SendIfResponds[int32](w.ID, objc.Sel("annotate:annotationParams:"), unsafe.Pointer(annotate), params)
	return rv
}
func (w WSLegacyEventProcessor) Annotate_internalAnnotationParamsWindowConnEventRegionIDIsCapturedDefaultWindowOverrideCaptureCid(annotate_internal *SLSEventRecord, params objectivec.IObject, conn uint32, id *uint64, captured bool, window bool, cid *bool) int32 {
	rv := objc.SendIfResponds[int32](w.ID, objc.Sel("annotate_internal:annotationParams:windowConn:eventRegionID:isCaptured:defaultWindow:overrideCaptureCid:"), unsafe.Pointer(annotate_internal), params, conn, id, captured, window, cid)
	return rv
}
func (w WSLegacyEventProcessor) Can_handle(can_handle *SLSEventRecord) bool {
	rv := objc.SendIfResponds[bool](w.ID, objc.Sel("can_handle:"), unsafe.Pointer(can_handle))
	return rv
}
func (w WSLegacyEventProcessor) ClearEventState() {
	objc.SendIfResponds[objc.ID](w.ID, objc.Sel("clearEventState"))
}
func (w WSLegacyEventProcessor) CreateAnnotationParams(params *SLSEventRecord) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](w.ID, objc.Sel("createAnnotationParams:"), unsafe.Pointer(params))
	return objectivec.Object{ID: rv}
}
func (w WSLegacyEventProcessor) Event_dispatchAnnotationParamsDispatcher(event_dispatch *SLSEventRecord, params objectivec.IObject, dispatcher objectivec.IObject) {
	objc.SendIfResponds[objc.ID](w.ID, objc.Sel("event_dispatch:annotationParams:dispatcher:"), unsafe.Pointer(event_dispatch), params, dispatcher)
}
func (w WSLegacyEventProcessor) Event_find_windowAnnotationParams(event_find_window *SLSEventRecord, params objectivec.IObject) {
	objc.SendIfResponds[objc.ID](w.ID, objc.Sel("event_find_window:annotationParams:"), unsafe.Pointer(event_find_window), params)
}
func (w WSLegacyEventProcessor) Post_event_annotateAnnotationParamsIsCapturedAnnotateWindowAnnotateConnectionEventRegion(post_event_annotate *SLSEventRecord, params objectivec.IObject, captured bool, window unsafe.Pointer, connection *CGXConnection, region WSStructuralRegionRef) bool {
	rv := objc.SendIfResponds[bool](w.ID, objc.Sel("post_event_annotate:annotationParams:isCaptured:annotateWindow:annotateConnection:eventRegion:"), unsafe.Pointer(post_event_annotate), params, captured, window, unsafe.Pointer(connection), region)
	return rv
}
func (w WSLegacyEventProcessor) ProcessEventDispatcher(event *SLSEventRecord, dispatcher objectivec.IObject) int64 {
	rv := objc.SendIfResponds[int64](w.ID, objc.Sel("processEvent:dispatcher:"), unsafe.Pointer(event), dispatcher)
	return rv
}
func (w WSLegacyEventProcessor) Route_annotate_eventAnnotationParamsDispatcher(route_annotate_event *SLSEventRecord, params objectivec.IObject, dispatcher objectivec.IObject) {
	objc.SendIfResponds[objc.ID](w.ID, objc.Sel("route_annotate_event:annotationParams:dispatcher:"), unsafe.Pointer(route_annotate_event), params, dispatcher)
}

func (w WSLegacyEventProcessor) DebugDescription() string {
	rv := objc.SendIfResponds[objc.ID](w.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
func (w WSLegacyEventProcessor) Description() string {
	rv := objc.SendIfResponds[objc.ID](w.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
func (w WSLegacyEventProcessor) Hash() uint64 {
	rv := objc.SendIfResponds[uint64](w.ID, objc.Sel("hash"))
	return rv
}
func (w WSLegacyEventProcessor) Superclass() objectivec.Class {
	rv := objc.SendIfResponds[objectivec.Class](w.ID, objc.Sel("superclass"))
	return objectivec.Class(rv)
}
