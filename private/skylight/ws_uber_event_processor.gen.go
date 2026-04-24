// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [WSUberEventProcessor] class.
var (
	_WSUberEventProcessorClass     WSUberEventProcessorClass
	_WSUberEventProcessorClassOnce sync.Once
)

func getWSUberEventProcessorClass() WSUberEventProcessorClass {
	_WSUberEventProcessorClassOnce.Do(func() {
		_WSUberEventProcessorClass = WSUberEventProcessorClass{class: objc.GetClass("WSUberEventProcessor")}
	})
	return _WSUberEventProcessorClass
}

// GetWSUberEventProcessorClass returns the class object for WSUberEventProcessor.
func GetWSUberEventProcessorClass() WSUberEventProcessorClass {
	return getWSUberEventProcessorClass()
}

type WSUberEventProcessorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (wc WSUberEventProcessorClass) Class() objc.Class {
	return wc.class
}

// Alloc allocates memory for a new instance of the class.
func (wc WSUberEventProcessorClass) Alloc() WSUberEventProcessor {
	rv := objc.Send[WSUberEventProcessor](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [WSUberEventProcessor.ClearEventState]
//   - [WSUberEventProcessor.ProcessEventContextDispatcher]
//   - [WSUberEventProcessor.InitWithSession]
//   - [WSUberEventProcessor.DebugDescription]
//   - [WSUberEventProcessor.Description]
//   - [WSUberEventProcessor.Hash]
//   - [WSUberEventProcessor.Superclass]
//
// See: https://developer.apple.com/documentation/SkyLight/WSUberEventProcessor
type WSUberEventProcessor struct {
	objectivec.Object
}

// WSUberEventProcessorFromID constructs a [WSUberEventProcessor] from an objc.ID.
func WSUberEventProcessorFromID(id objc.ID) WSUberEventProcessor {
	return WSUberEventProcessor{objectivec.Object{ID: id}}
}

// Ensure WSUberEventProcessor implements IWSUberEventProcessor.
var _ IWSUberEventProcessor = WSUberEventProcessor{}

// An interface definition for the [WSUberEventProcessor] class.
//
// # Methods
//
//   - [IWSUberEventProcessor.ClearEventState]
//   - [IWSUberEventProcessor.ProcessEventContextDispatcher]
//   - [IWSUberEventProcessor.InitWithSession]
//   - [IWSUberEventProcessor.DebugDescription]
//   - [IWSUberEventProcessor.Description]
//   - [IWSUberEventProcessor.Hash]
//   - [IWSUberEventProcessor.Superclass]
//
// See: https://developer.apple.com/documentation/SkyLight/WSUberEventProcessor
type IWSUberEventProcessor interface {
	objectivec.IObject

	// Topic: Methods

	ClearEventState()
	ProcessEventContextDispatcher(event *SLSEventRecordRef, context unsafe.Pointer, dispatcher objectivec.IObject) int64
	InitWithSession(session unsafe.Pointer) WSUberEventProcessor
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objc.Class
}

// Init initializes the instance.
func (w WSUberEventProcessor) Init() WSUberEventProcessor {
	rv := objc.Send[WSUberEventProcessor](w.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (w WSUberEventProcessor) Autorelease() WSUberEventProcessor {
	rv := objc.Send[WSUberEventProcessor](w.ID, objc.Sel("autorelease"))
	return rv
}

// NewWSUberEventProcessor creates a new WSUberEventProcessor instance.
func NewWSUberEventProcessor() WSUberEventProcessor {
	class := getWSUberEventProcessorClass()
	rv := objc.Send[WSUberEventProcessor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/WSUberEventProcessor/initWithSession:
func NewWSUberEventProcessorWithSession(session unsafe.Pointer) WSUberEventProcessor {
	instance := getWSUberEventProcessorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSession:"), session)
	return WSUberEventProcessorFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/WSUberEventProcessor/clearEventState
func (w WSUberEventProcessor) ClearEventState() {
	objc.Send[objc.ID](w.ID, objc.Sel("clearEventState"))
}

// See: https://developer.apple.com/documentation/SkyLight/WSUberEventProcessor/processEvent:context:dispatcher:
func (w WSUberEventProcessor) ProcessEventContextDispatcher(event *SLSEventRecordRef, context unsafe.Pointer, dispatcher objectivec.IObject) int64 {
	rv := objc.Send[int64](w.ID, objc.Sel("processEvent:context:dispatcher:"), event, context, dispatcher)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/WSUberEventProcessor/initWithSession:
func (w WSUberEventProcessor) InitWithSession(session unsafe.Pointer) WSUberEventProcessor {
	rv := objc.Send[WSUberEventProcessor](w.ID, objc.Sel("initWithSession:"), session)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/WSUberEventProcessor/debugDescription
func (w WSUberEventProcessor) DebugDescription() string {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/SkyLight/WSUberEventProcessor/description
func (w WSUberEventProcessor) Description() string {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/SkyLight/WSUberEventProcessor/hash
func (w WSUberEventProcessor) Hash() uint64 {
	rv := objc.Send[uint64](w.ID, objc.Sel("hash"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/WSUberEventProcessor/superclass
func (w WSUberEventProcessor) Superclass() objc.Class {
	rv := objc.Send[objc.Class](w.ID, objc.Sel("superclass"))
	return rv
}
