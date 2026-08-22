// Code generated from Apple documentation for ParavirtualizedGraphics. DO NOT EDIT.

package paravirtualizedgraphics

import (
	"sync"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/dispatch"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [PGDisplayDescriptor] class.
var (
	_PGDisplayDescriptorClass     PGDisplayDescriptorClass
	_PGDisplayDescriptorClassOnce sync.Once
)

func getPGDisplayDescriptorClass() PGDisplayDescriptorClass {
	_PGDisplayDescriptorClassOnce.Do(func() {
		_PGDisplayDescriptorClass = PGDisplayDescriptorClass{class: objc.GetClass("PGDisplayDescriptor")}
	})
	return _PGDisplayDescriptorClass
}

// GetPGDisplayDescriptorClass returns the class object for PGDisplayDescriptor.
func GetPGDisplayDescriptorClass() PGDisplayDescriptorClass {
	return getPGDisplayDescriptorClass()
}

type PGDisplayDescriptorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (pc PGDisplayDescriptorClass) Class() objc.Class {
	return pc.class
}

// Alloc allocates memory for a new instance of the class.
func (pc PGDisplayDescriptorClass) Alloc() PGDisplayDescriptor {
	rv := objc.Send[PGDisplayDescriptor](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// A descriptor for a virtual display.
//
// # Specifying the Display Properties
//
//   - [PGDisplayDescriptor.Name]: The display’s name as seen in the guest operating environment.
//   - [PGDisplayDescriptor.SetName]
//   - [PGDisplayDescriptor.SizeInMillimeters]: The size in millimeters of the virtual display.
//   - [PGDisplayDescriptor.SetSizeInMillimeters]
//
// # Setting the Dispatch Queue
//
//   - [PGDisplayDescriptor.Queue]: The queue that the framework uses when dispatching messages to any of the display’s registered handlers.
//   - [PGDisplayDescriptor.SetQueue]
//
// # Managing Cursor Events
//
//   - [PGDisplayDescriptor.CursorGlyphHandler]: A handler that the framework calls to change the cursor’s appearance.
//   - [PGDisplayDescriptor.SetCursorGlyphHandler]
//   - [PGDisplayDescriptor.CursorShowHandler]: A handler that the framework calls to change the cursor’s visibility.
//   - [PGDisplayDescriptor.SetCursorShowHandler]
//
// # Handling Mode Changes
//
//   - [PGDisplayDescriptor.ModeChangeHandler]: A handler that the framework calls to change the virtual display’s graphics mode.
//   - [PGDisplayDescriptor.SetModeChangeHandler]
//
// # Handling Frame Events
//
//   - [PGDisplayDescriptor.NewFrameEventHandler]: A handler that the framework calls when the guest environment has a new frame to display.
//   - [PGDisplayDescriptor.SetNewFrameEventHandler]
//
// # Instance Properties
//
//   - [PGDisplayDescriptor.CursorMoveHandler]
//   - [PGDisplayDescriptor.SetCursorMoveHandler]
//
// See: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDisplayDescriptor
type PGDisplayDescriptor struct {
	objectivec.Object
}

// PGDisplayDescriptorFromID constructs a [PGDisplayDescriptor] from an objc.ID.
//
// A descriptor for a virtual display.
func PGDisplayDescriptorFromID(id objc.ID) PGDisplayDescriptor {
	return PGDisplayDescriptor{objectivec.Object{ID: id}}
}

// NOTE: PGDisplayDescriptor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [PGDisplayDescriptor] class.
//
// # Specifying the Display Properties
//
//   - [IPGDisplayDescriptor.Name]: The display’s name as seen in the guest operating environment.
//   - [IPGDisplayDescriptor.SetName]
//   - [IPGDisplayDescriptor.SizeInMillimeters]: The size in millimeters of the virtual display.
//   - [IPGDisplayDescriptor.SetSizeInMillimeters]
//
// # Setting the Dispatch Queue
//
//   - [IPGDisplayDescriptor.Queue]: The queue that the framework uses when dispatching messages to any of the display’s registered handlers.
//   - [IPGDisplayDescriptor.SetQueue]
//
// # Managing Cursor Events
//
//   - [IPGDisplayDescriptor.CursorGlyphHandler]: A handler that the framework calls to change the cursor’s appearance.
//   - [IPGDisplayDescriptor.SetCursorGlyphHandler]
//   - [IPGDisplayDescriptor.CursorShowHandler]: A handler that the framework calls to change the cursor’s visibility.
//   - [IPGDisplayDescriptor.SetCursorShowHandler]
//
// # Handling Mode Changes
//
//   - [IPGDisplayDescriptor.ModeChangeHandler]: A handler that the framework calls to change the virtual display’s graphics mode.
//   - [IPGDisplayDescriptor.SetModeChangeHandler]
//
// # Handling Frame Events
//
//   - [IPGDisplayDescriptor.NewFrameEventHandler]: A handler that the framework calls when the guest environment has a new frame to display.
//   - [IPGDisplayDescriptor.SetNewFrameEventHandler]
//
// # Instance Properties
//
//   - [IPGDisplayDescriptor.CursorMoveHandler]
//   - [IPGDisplayDescriptor.SetCursorMoveHandler]
//
// See: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDisplayDescriptor
type IPGDisplayDescriptor interface {
	objectivec.IObject

	// Topic: Specifying the Display Properties

	// The display’s name as seen in the guest operating environment.
	Name() string
	SetName(value string)
	// The size in millimeters of the virtual display.
	SizeInMillimeters() corefoundation.CGSize
	SetSizeInMillimeters(value corefoundation.CGSize)

	// Topic: Setting the Dispatch Queue

	// The queue that the framework uses when dispatching messages to any of the display’s registered handlers.
	Queue() dispatch.Queue
	SetQueue(value dispatch.Queue)

	// Topic: Managing Cursor Events

	// A handler that the framework calls to change the cursor’s appearance.
	CursorGlyphHandler() BitmapImageRepPGDisplayCoord_tHandler
	SetCursorGlyphHandler(value BitmapImageRepPGDisplayCoord_tHandler)
	// A handler that the framework calls to change the cursor’s visibility.
	CursorShowHandler() PGDisplayCursorShowHandler
	SetCursorShowHandler(value PGDisplayCursorShowHandler)

	// Topic: Handling Mode Changes

	// A handler that the framework calls to change the virtual display’s graphics mode.
	ModeChangeHandler() PGDisplayCoord_tUint32Handler
	SetModeChangeHandler(value PGDisplayCoord_tUint32Handler)

	// Topic: Handling Frame Events

	// A handler that the framework calls when the guest environment has a new frame to display.
	NewFrameEventHandler() VoidHandler
	SetNewFrameEventHandler(value VoidHandler)

	// Topic: Instance Properties

	CursorMoveHandler() VoidHandler
	SetCursorMoveHandler(value VoidHandler)
}

// Init initializes the instance.
func (p PGDisplayDescriptor) Init() PGDisplayDescriptor {
	rv := objc.Send[PGDisplayDescriptor](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p PGDisplayDescriptor) Autorelease() PGDisplayDescriptor {
	rv := objc.Send[PGDisplayDescriptor](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewPGDisplayDescriptor creates a new PGDisplayDescriptor instance.
func NewPGDisplayDescriptor() PGDisplayDescriptor {
	class := getPGDisplayDescriptorClass()
	rv := objc.Send[PGDisplayDescriptor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The display’s name as seen in the guest operating environment.
//
// # Discussion
//
// The default value is `Apple Virtual`. The framework truncates the name to
// 13 characters.
//
// The device propagates the display name into the guest environment, so the
// name may be visible in the guest operating system’s user interface.
//
// See: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDisplayDescriptor/name
func (p PGDisplayDescriptor) Name() string {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("name"))
	return foundation.NSStringFromID(rv).String()
}
func (p PGDisplayDescriptor) SetName(value string) {
	objc.Send[struct{}](p.ID, objc.Sel("setName:"), objc.String(value))
}

// The size in millimeters of the virtual display.
//
// # Discussion
//
// The device propagates the display size to the guest operating system. The
// app can scale the resulting screen data to a different size in its user
// interface.
//
// See: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDisplayDescriptor/sizeInMillimeters
func (p PGDisplayDescriptor) SizeInMillimeters() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](p.ID, objc.Sel("sizeInMillimeters"))
	return corefoundation.CGSize(rv)
}
func (p PGDisplayDescriptor) SetSizeInMillimeters(value corefoundation.CGSize) {
	objc.Send[struct{}](p.ID, objc.Sel("setSizeInMillimeters:"), value)
}

// The queue that the framework uses when dispatching messages to any of the
// display’s registered handlers.
//
// # Discussion
//
// Most often, your app provides a serial queue. If you can benefit from
// dispatching events out of order, handle the messages on the serial queue
// and redispatch them to other queues as necessary.
//
// See: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDisplayDescriptor/queue
func (p PGDisplayDescriptor) Queue() dispatch.Queue {
	rv := objc.Send[uintptr](p.ID, objc.Sel("queue"))
	return dispatch.QueueFromHandle(rv)
}
func (p PGDisplayDescriptor) SetQueue(value dispatch.Queue) {
	objc.Send[struct{}](p.ID, objc.Sel("setQueue:"), uintptr(value.Handle()))
}

// A handler that the framework calls to change the cursor’s appearance.
//
// See: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDisplayDescriptor/cursorGlyphHandler
func (p PGDisplayDescriptor) CursorGlyphHandler() BitmapImageRepPGDisplayCoord_tHandler {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("cursorGlyphHandler"))
	_ = rv
	return nil
}
func (p PGDisplayDescriptor) SetCursorGlyphHandler(value BitmapImageRepPGDisplayCoord_tHandler) {
	block, cleanup := NewBitmapImageRepPGDisplayCoord_tBlock(value)
	defer cleanup()
	objc.Send[struct{}](p.ID, objc.Sel("setCursorGlyphHandler:"), block)
}

// A handler that the framework calls to change the cursor’s visibility.
//
// See: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDisplayDescriptor/cursorShowHandler
func (p PGDisplayDescriptor) CursorShowHandler() PGDisplayCursorShowHandler {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("cursorShowHandler"))
	_ = rv
	return nil
}
func (p PGDisplayDescriptor) SetCursorShowHandler(value PGDisplayCursorShowHandler) {
	block, cleanup := NewPGDisplayCursorShowBlock(value)
	defer cleanup()
	objc.Send[struct{}](p.ID, objc.Sel("setCursorShowHandler:"), block)
}

// A handler that the framework calls to change the virtual display’s
// graphics mode.
//
// See: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDisplayDescriptor/modeChangeHandler
func (p PGDisplayDescriptor) ModeChangeHandler() PGDisplayCoord_tUint32Handler {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("modeChangeHandler"))
	_ = rv
	return nil
}
func (p PGDisplayDescriptor) SetModeChangeHandler(value PGDisplayCoord_tUint32Handler) {
	block, cleanup := NewPGDisplayCoord_tUint32Block(value)
	defer cleanup()
	objc.Send[struct{}](p.ID, objc.Sel("setModeChangeHandler:"), block)
}

// A handler that the framework calls when the guest environment has a new
// frame to display.
//
// See: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDisplayDescriptor/newFrameEventHandler
func (p PGDisplayDescriptor) NewFrameEventHandler() VoidHandler {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("newFrameEventHandler"))
	_ = rv
	return nil
}
func (p PGDisplayDescriptor) SetNewFrameEventHandler(value VoidHandler) {
	block, cleanup := NewVoidBlock(value)
	defer cleanup()
	objc.Send[struct{}](p.ID, objc.Sel("setNewFrameEventHandler:"), block)
}

// See: https://developer.apple.com/documentation/ParavirtualizedGraphics/PGDisplayDescriptor/cursorMoveHandler
func (p PGDisplayDescriptor) CursorMoveHandler() VoidHandler {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("cursorMoveHandler"))
	_ = rv
	return nil
}
func (p PGDisplayDescriptor) SetCursorMoveHandler(value VoidHandler) {
	block, cleanup := NewVoidBlock(value)
	defer cleanup()
	objc.Send[struct{}](p.ID, objc.Sel("setCursorMoveHandler:"), block)
}
