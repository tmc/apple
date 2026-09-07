// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [PKGSession] class.
var (
	_PKGSessionClass     PKGSessionClass
	_PKGSessionClassOnce sync.Once
)

func getPKGSessionClass() PKGSessionClass {
	_PKGSessionClassOnce.Do(func() {
		_PKGSessionClass = PKGSessionClass{class: objc.GetClass("PKGSession")}
	})
	return _PKGSessionClass
}

// GetPKGSessionClass returns the class object for PKGSession.
func GetPKGSessionClass() PKGSessionClass {
	return getPKGSessionClass()
}

type PKGSessionClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (pc PKGSessionClass) Class() objc.Class {
	return pc.class
}

// Alloc allocates memory for a new instance of the class.
func (pc PKGSessionClass) Alloc() PKGSession {
	rv := objc.SendIfResponds[PKGSession](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [PKGSession.AcknowledgeProposedWindowMovementLocationTimestamp]
//   - [PKGSession.AdjustDraggedWindowOffsetAbsolute]
//   - [PKGSession.BeginWindowDragRelativeToMouseWindowOptions]
//   - [PKGSession.BlockWindowOrdering]
//   - [PKGSession.CancelWindowDrag]
//   - [PKGSession.DisableTransformedDragMode]
//   - [PKGSession.DraggedWindowID]
//   - [PKGSession.DraggedWindowOffset]
//   - [PKGSession.EnableTransformedDragMode]
//   - [PKGSession.FinishSetup]
//   - [PKGSession.HasDeferredOperations]
//   - [PKGSession.IngestEventPsnFilterOnly]
//   - [PKGSession.Invalidate]
//   - [PKGSession.PerformDeferredOperations]
//   - [PKGSession.RebuildGesturesForWindow]
//   - [PKGSession.RequestDeferredKitActivationCountOrderingEventTimeWindowID]
//   - [PKGSession.SetAccessibilityDragTouchUp]
type PKGSession struct {
	objectivec.Object
}

// PKGSessionFromID constructs a [PKGSession] from an objc.ID.
func PKGSessionFromID(id objc.ID) PKGSession {
	return PKGSession{objectivec.Object{ID: id}}
}

// Ensure PKGSession implements IPKGSession.
var _ IPKGSession = PKGSession{}

// An interface definition for the [PKGSession] class.
//
// # Methods
//
//   - [IPKGSession.AcknowledgeProposedWindowMovementLocationTimestamp]
//   - [IPKGSession.AdjustDraggedWindowOffsetAbsolute]
//   - [IPKGSession.BeginWindowDragRelativeToMouseWindowOptions]
//   - [IPKGSession.BlockWindowOrdering]
//   - [IPKGSession.CancelWindowDrag]
//   - [IPKGSession.DisableTransformedDragMode]
//   - [IPKGSession.DraggedWindowID]
//   - [IPKGSession.DraggedWindowOffset]
//   - [IPKGSession.EnableTransformedDragMode]
//   - [IPKGSession.FinishSetup]
//   - [IPKGSession.HasDeferredOperations]
//   - [IPKGSession.IngestEventPsnFilterOnly]
//   - [IPKGSession.Invalidate]
//   - [IPKGSession.PerformDeferredOperations]
//   - [IPKGSession.RebuildGesturesForWindow]
//   - [IPKGSession.RequestDeferredKitActivationCountOrderingEventTimeWindowID]
//   - [IPKGSession.SetAccessibilityDragTouchUp]
type IPKGSession interface {
	objectivec.IObject

	// Topic: Methods

	AcknowledgeProposedWindowMovementLocationTimestamp(movement corefoundation.CGPoint, location corefoundation.CGPoint, timestamp uint64)
	AdjustDraggedWindowOffsetAbsolute(offset corefoundation.CGPoint, absolute bool)
	BeginWindowDragRelativeToMouseWindowOptions(mouse corefoundation.CGPoint, window uint32, options uint64) bool
	BlockWindowOrdering()
	CancelWindowDrag()
	DisableTransformedDragMode(mode bool)
	DraggedWindowID() uint32
	DraggedWindowOffset() corefoundation.CGPoint
	EnableTransformedDragMode()
	FinishSetup()
	HasDeferredOperations() bool
	IngestEventPsnFilterOnly(event *SLSEventRecord, psn CPSProcessSerNum, only bool)
	Invalidate()
	PerformDeferredOperations()
	RebuildGesturesForWindow(window uint32)
	RequestDeferredKitActivationCountOrderingEventTimeWindowID(activation bool, count uint64, ordering bool, time uint64, id uint32)
	SetAccessibilityDragTouchUp(up bool)
}

// Init initializes the instance.
func (p PKGSession) Init() PKGSession {
	rv := objc.SendIfResponds[PKGSession](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p PKGSession) Autorelease() PKGSession {
	rv := objc.SendIfResponds[PKGSession](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewPKGSession creates a new PKGSession instance.
func NewPKGSession() PKGSession {
	class := getPKGSessionClass()
	rv := objc.SendIfResponds[PKGSession](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (p PKGSession) AcknowledgeProposedWindowMovementLocationTimestamp(movement corefoundation.CGPoint, location corefoundation.CGPoint, timestamp uint64) {
	objc.SendIfResponds[objc.ID](p.ID, objc.Sel("acknowledgeProposedWindowMovement:location:timestamp:"), movement, location, timestamp)
}
func (p PKGSession) AdjustDraggedWindowOffsetAbsolute(offset corefoundation.CGPoint, absolute bool) {
	objc.SendIfResponds[objc.ID](p.ID, objc.Sel("adjustDraggedWindowOffset:absolute:"), offset, absolute)
}
func (p PKGSession) BeginWindowDragRelativeToMouseWindowOptions(mouse corefoundation.CGPoint, window uint32, options uint64) bool {
	rv := objc.SendIfResponds[bool](p.ID, objc.Sel("beginWindowDragRelativeToMouse:window:options:"), mouse, window, options)
	return rv
}
func (p PKGSession) BlockWindowOrdering() {
	objc.SendIfResponds[objc.ID](p.ID, objc.Sel("blockWindowOrdering"))
}
func (p PKGSession) CancelWindowDrag() {
	objc.SendIfResponds[objc.ID](p.ID, objc.Sel("cancelWindowDrag"))
}
func (p PKGSession) DisableTransformedDragMode(mode bool) {
	objc.SendIfResponds[objc.ID](p.ID, objc.Sel("disableTransformedDragMode:"), mode)
}
func (p PKGSession) EnableTransformedDragMode() {
	objc.SendIfResponds[objc.ID](p.ID, objc.Sel("enableTransformedDragMode"))
}
func (p PKGSession) FinishSetup() {
	objc.SendIfResponds[objc.ID](p.ID, objc.Sel("finishSetup"))
}
func (p PKGSession) IngestEventPsnFilterOnly(event *SLSEventRecord, psn CPSProcessSerNum, only bool) {
	objc.SendIfResponds[objc.ID](p.ID, objc.Sel("ingestEvent:psn:filterOnly:"), unsafe.Pointer(event), psn, only)
}
func (p PKGSession) Invalidate() {
	objc.SendIfResponds[objc.ID](p.ID, objc.Sel("invalidate"))
}
func (p PKGSession) PerformDeferredOperations() {
	objc.SendIfResponds[objc.ID](p.ID, objc.Sel("performDeferredOperations"))
}
func (p PKGSession) RebuildGesturesForWindow(window uint32) {
	objc.SendIfResponds[objc.ID](p.ID, objc.Sel("rebuildGesturesForWindow:"), window)
}
func (p PKGSession) RequestDeferredKitActivationCountOrderingEventTimeWindowID(activation bool, count uint64, ordering bool, time uint64, id uint32) {
	objc.SendIfResponds[objc.ID](p.ID, objc.Sel("requestDeferredKitActivation:count:ordering:eventTime:windowID:"), activation, count, ordering, time, id)
}
func (p PKGSession) SetAccessibilityDragTouchUp(up bool) {
	objc.SendIfResponds[objc.ID](p.ID, objc.Sel("setAccessibilityDragTouchUp:"), up)
}

func (p PKGSession) DraggedWindowID() uint32 {
	rv := objc.SendIfResponds[uint32](p.ID, objc.Sel("draggedWindowID"))
	return rv
}
func (p PKGSession) DraggedWindowOffset() corefoundation.CGPoint {
	rv := objc.SendIfResponds[corefoundation.CGPoint](p.ID, objc.Sel("draggedWindowOffset"))
	return corefoundation.CGPoint(rv)
}
func (p PKGSession) HasDeferredOperations() bool {
	rv := objc.SendIfResponds[bool](p.ID, objc.Sel("hasDeferredOperations"))
	return rv
}
