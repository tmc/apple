// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CPXModernEventProcessor] class.
var (
	_CPXModernEventProcessorClass     CPXModernEventProcessorClass
	_CPXModernEventProcessorClassOnce sync.Once
)

func getCPXModernEventProcessorClass() CPXModernEventProcessorClass {
	_CPXModernEventProcessorClassOnce.Do(func() {
		_CPXModernEventProcessorClass = CPXModernEventProcessorClass{class: objc.GetClass("CPXModernEventProcessor")}
	})
	return _CPXModernEventProcessorClass
}

// GetCPXModernEventProcessorClass returns the class object for CPXModernEventProcessor.
func GetCPXModernEventProcessorClass() CPXModernEventProcessorClass {
	return getCPXModernEventProcessorClass()
}

type CPXModernEventProcessorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CPXModernEventProcessorClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CPXModernEventProcessorClass) Alloc() CPXModernEventProcessor {
	rv := objc.SendIfResponds[CPXModernEventProcessor](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [CPXModernEventProcessor.ClearEventState]
//   - [CPXModernEventProcessor.ProcessEventContextDispatcher]
//   - [CPXModernEventProcessor.DebugDescription]
//   - [CPXModernEventProcessor.Description]
//   - [CPXModernEventProcessor.Hash]
//   - [CPXModernEventProcessor.Superclass]
type CPXModernEventProcessor struct {
	objectivec.Object
}

// CPXModernEventProcessorFromID constructs a [CPXModernEventProcessor] from an objc.ID.
func CPXModernEventProcessorFromID(id objc.ID) CPXModernEventProcessor {
	return CPXModernEventProcessor{objectivec.Object{ID: id}}
}

// Ensure CPXModernEventProcessor implements ICPXModernEventProcessor.
var _ ICPXModernEventProcessor = CPXModernEventProcessor{}

// An interface definition for the [CPXModernEventProcessor] class.
//
// # Methods
//
//   - [ICPXModernEventProcessor.ClearEventState]
//   - [ICPXModernEventProcessor.ProcessEventContextDispatcher]
//   - [ICPXModernEventProcessor.DebugDescription]
//   - [ICPXModernEventProcessor.Description]
//   - [ICPXModernEventProcessor.Hash]
//   - [ICPXModernEventProcessor.Superclass]
type ICPXModernEventProcessor interface {
	objectivec.IObject

	// Topic: Methods

	ClearEventState()
	ProcessEventContextDispatcher(event SLSEventRecordRef, context *CPXEventProcessorContext, dispatcher objectivec.IObject) int64
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objectivec.Class
}

// Init initializes the instance.
func (c CPXModernEventProcessor) Init() CPXModernEventProcessor {
	rv := objc.SendIfResponds[CPXModernEventProcessor](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CPXModernEventProcessor) Autorelease() CPXModernEventProcessor {
	rv := objc.SendIfResponds[CPXModernEventProcessor](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCPXModernEventProcessor creates a new CPXModernEventProcessor instance.
func NewCPXModernEventProcessor() CPXModernEventProcessor {
	class := getCPXModernEventProcessorClass()
	rv := objc.SendIfResponds[CPXModernEventProcessor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (c CPXModernEventProcessor) ClearEventState() {
	objc.SendIfResponds[objc.ID](c.ID, objc.Sel("clearEventState"))
}
func (c CPXModernEventProcessor) ProcessEventContextDispatcher(event SLSEventRecordRef, context *CPXEventProcessorContext, dispatcher objectivec.IObject) int64 {
	rv := objc.SendIfResponds[int64](c.ID, objc.Sel("processEvent:context:dispatcher:"), event, unsafe.Pointer(context), dispatcher)
	return rv
}

func (c CPXModernEventProcessor) DebugDescription() string {
	rv := objc.SendIfResponds[objc.ID](c.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
func (c CPXModernEventProcessor) Description() string {
	rv := objc.SendIfResponds[objc.ID](c.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
func (c CPXModernEventProcessor) Hash() uint64 {
	rv := objc.SendIfResponds[uint64](c.ID, objc.Sel("hash"))
	return rv
}
func (c CPXModernEventProcessor) Superclass() objectivec.Class {
	rv := objc.SendIfResponds[objectivec.Class](c.ID, objc.Sel("superclass"))
	return objectivec.Class(rv)
}
