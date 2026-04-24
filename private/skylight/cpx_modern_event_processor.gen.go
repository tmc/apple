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
	rv := objc.Send[CPXModernEventProcessor](objc.ID(cc.class), objc.Sel("alloc"))
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
//
// See: https://developer.apple.com/documentation/SkyLight/CPXModernEventProcessor
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
//
// See: https://developer.apple.com/documentation/SkyLight/CPXModernEventProcessor
type ICPXModernEventProcessor interface {
	objectivec.IObject

	// Topic: Methods

	ClearEventState()
	ProcessEventContextDispatcher(event SLSEventRecordRef, context unsafe.Pointer, dispatcher objectivec.IObject) int64
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objc.Class
}

// Init initializes the instance.
func (c CPXModernEventProcessor) Init() CPXModernEventProcessor {
	rv := objc.Send[CPXModernEventProcessor](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CPXModernEventProcessor) Autorelease() CPXModernEventProcessor {
	rv := objc.Send[CPXModernEventProcessor](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCPXModernEventProcessor creates a new CPXModernEventProcessor instance.
func NewCPXModernEventProcessor() CPXModernEventProcessor {
	class := getCPXModernEventProcessorClass()
	rv := objc.Send[CPXModernEventProcessor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/CPXModernEventProcessor/clearEventState
func (c CPXModernEventProcessor) ClearEventState() {
	objc.Send[objc.ID](c.ID, objc.Sel("clearEventState"))
}

// See: https://developer.apple.com/documentation/SkyLight/CPXModernEventProcessor/processEvent:context:dispatcher:
func (c CPXModernEventProcessor) ProcessEventContextDispatcher(event SLSEventRecordRef, context unsafe.Pointer, dispatcher objectivec.IObject) int64 {
	rv := objc.Send[int64](c.ID, objc.Sel("processEvent:context:dispatcher:"), event, context, dispatcher)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/CPXModernEventProcessor/debugDescription
func (c CPXModernEventProcessor) DebugDescription() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/SkyLight/CPXModernEventProcessor/description
func (c CPXModernEventProcessor) Description() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/SkyLight/CPXModernEventProcessor/hash
func (c CPXModernEventProcessor) Hash() uint64 {
	rv := objc.Send[uint64](c.ID, objc.Sel("hash"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/CPXModernEventProcessor/superclass
func (c CPXModernEventProcessor) Superclass() objc.Class {
	rv := objc.Send[objc.Class](c.ID, objc.Sel("superclass"))
	return rv
}
