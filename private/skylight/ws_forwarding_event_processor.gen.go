// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [WSForwardingEventProcessor] class.
var (
	_WSForwardingEventProcessorClass     WSForwardingEventProcessorClass
	_WSForwardingEventProcessorClassOnce sync.Once
)

func getWSForwardingEventProcessorClass() WSForwardingEventProcessorClass {
	_WSForwardingEventProcessorClassOnce.Do(func() {
		_WSForwardingEventProcessorClass = WSForwardingEventProcessorClass{class: objc.GetClass("WSForwardingEventProcessor")}
	})
	return _WSForwardingEventProcessorClass
}

// GetWSForwardingEventProcessorClass returns the class object for WSForwardingEventProcessor.
func GetWSForwardingEventProcessorClass() WSForwardingEventProcessorClass {
	return getWSForwardingEventProcessorClass()
}

type WSForwardingEventProcessorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (wc WSForwardingEventProcessorClass) Class() objc.Class {
	return wc.class
}

// Alloc allocates memory for a new instance of the class.
func (wc WSForwardingEventProcessorClass) Alloc() WSForwardingEventProcessor {
	rv := objc.SendIfResponds[WSForwardingEventProcessor](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [WSForwardingEventProcessor.ClearEventState]
//   - [WSForwardingEventProcessor.ProcessEventDispatcher]
//   - [WSForwardingEventProcessor.InitWithHidSystem]
//   - [WSForwardingEventProcessor.DebugDescription]
//   - [WSForwardingEventProcessor.Description]
//   - [WSForwardingEventProcessor.Hash]
//   - [WSForwardingEventProcessor.Superclass]
type WSForwardingEventProcessor struct {
	objectivec.Object
}

// WSForwardingEventProcessorFromID constructs a [WSForwardingEventProcessor] from an objc.ID.
func WSForwardingEventProcessorFromID(id objc.ID) WSForwardingEventProcessor {
	return WSForwardingEventProcessor{objectivec.Object{ID: id}}
}

// Ensure WSForwardingEventProcessor implements IWSForwardingEventProcessor.
var _ IWSForwardingEventProcessor = WSForwardingEventProcessor{}

// An interface definition for the [WSForwardingEventProcessor] class.
//
// # Methods
//
//   - [IWSForwardingEventProcessor.ClearEventState]
//   - [IWSForwardingEventProcessor.ProcessEventDispatcher]
//   - [IWSForwardingEventProcessor.InitWithHidSystem]
//   - [IWSForwardingEventProcessor.DebugDescription]
//   - [IWSForwardingEventProcessor.Description]
//   - [IWSForwardingEventProcessor.Hash]
//   - [IWSForwardingEventProcessor.Superclass]
type IWSForwardingEventProcessor interface {
	objectivec.IObject

	// Topic: Methods

	ClearEventState()
	ProcessEventDispatcher(event *SLSEventRecord, dispatcher objectivec.IObject) int64
	InitWithHidSystem(system objectivec.IObject) WSForwardingEventProcessor
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objectivec.Class
}

// Init initializes the instance.
func (w WSForwardingEventProcessor) Init() WSForwardingEventProcessor {
	rv := objc.SendIfResponds[WSForwardingEventProcessor](w.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (w WSForwardingEventProcessor) Autorelease() WSForwardingEventProcessor {
	rv := objc.SendIfResponds[WSForwardingEventProcessor](w.ID, objc.Sel("autorelease"))
	return rv
}

// NewWSForwardingEventProcessor creates a new WSForwardingEventProcessor instance.
func NewWSForwardingEventProcessor() WSForwardingEventProcessor {
	class := getWSForwardingEventProcessorClass()
	rv := objc.SendIfResponds[WSForwardingEventProcessor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewWSForwardingEventProcessorWithHidSystem(system objectivec.IObject) WSForwardingEventProcessor {
	instance := getWSForwardingEventProcessorClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithHidSystem:"), system)
	return WSForwardingEventProcessorFromID(rv)
}

func (w WSForwardingEventProcessor) ClearEventState() {
	objc.SendIfResponds[objc.ID](w.ID, objc.Sel("clearEventState"))
}
func (w WSForwardingEventProcessor) ProcessEventDispatcher(event *SLSEventRecord, dispatcher objectivec.IObject) int64 {
	rv := objc.SendIfResponds[int64](w.ID, objc.Sel("processEvent:dispatcher:"), unsafe.Pointer(event), dispatcher)
	return rv
}
func (w WSForwardingEventProcessor) InitWithHidSystem(system objectivec.IObject) WSForwardingEventProcessor {
	rv := objc.SendIfResponds[WSForwardingEventProcessor](w.ID, objc.Sel("initWithHidSystem:"), system)
	return rv
}

func (w WSForwardingEventProcessor) DebugDescription() string {
	rv := objc.SendIfResponds[objc.ID](w.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
func (w WSForwardingEventProcessor) Description() string {
	rv := objc.SendIfResponds[objc.ID](w.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
func (w WSForwardingEventProcessor) Hash() uint64 {
	rv := objc.SendIfResponds[uint64](w.ID, objc.Sel("hash"))
	return rv
}
func (w WSForwardingEventProcessor) Superclass() objectivec.Class {
	rv := objc.SendIfResponds[objectivec.Class](w.ID, objc.Sel("superclass"))
	return objectivec.Class(rv)
}
