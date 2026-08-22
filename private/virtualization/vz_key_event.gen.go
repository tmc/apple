// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZKeyEvent] class.
var (
	_VZKeyEventClass     VZKeyEventClass
	_VZKeyEventClassOnce sync.Once
)

func getVZKeyEventClass() VZKeyEventClass {
	_VZKeyEventClassOnce.Do(func() {
		_VZKeyEventClass = VZKeyEventClass{class: objc.GetClass("_VZKeyEvent")}
	})
	return _VZKeyEventClass
}

// GetVZKeyEventClass returns the class object for _VZKeyEvent.
func GetVZKeyEventClass() VZKeyEventClass {
	return getVZKeyEventClass()
}

type VZKeyEventClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZKeyEventClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZKeyEventClass) Alloc() VZKeyEvent {
	rv := objc.SendIfResponds[VZKeyEvent](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [VZKeyEvent.KeyCode]
//   - [VZKeyEvent.Type]
//   - [VZKeyEvent.InitWithEvent]
//   - [VZKeyEvent.InitWithTypeKeyCode]
type VZKeyEvent struct {
	objectivec.Object
}

// VZKeyEventFromID constructs a [VZKeyEvent] from an objc.ID.
func VZKeyEventFromID(id objc.ID) VZKeyEvent {
	return VZKeyEvent{objectivec.Object{ID: id}}
}

// Ensure VZKeyEvent implements IVZKeyEvent.
var _ IVZKeyEvent = VZKeyEvent{}

// An interface definition for the [VZKeyEvent] class.
//
// # Methods
//
//   - [IVZKeyEvent.KeyCode]
//   - [IVZKeyEvent.Type]
//   - [IVZKeyEvent.InitWithEvent]
//   - [IVZKeyEvent.InitWithTypeKeyCode]
type IVZKeyEvent interface {
	objectivec.IObject

	// Topic: Methods

	KeyCode() uint16
	Type() int64
	InitWithEvent(event objectivec.IObject) VZKeyEvent
	InitWithTypeKeyCode(type_ int64, code uint16) VZKeyEvent
}

// Init initializes the instance.
func (v VZKeyEvent) Init() VZKeyEvent {
	rv := objc.SendIfResponds[VZKeyEvent](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZKeyEvent) Autorelease() VZKeyEvent {
	rv := objc.SendIfResponds[VZKeyEvent](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZKeyEvent creates a new VZKeyEvent instance.
func NewVZKeyEvent() VZKeyEvent {
	class := getVZKeyEventClass()
	rv := objc.SendIfResponds[VZKeyEvent](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewVZKeyEventWithEvent(event objectivec.IObject) VZKeyEvent {
	instance := getVZKeyEventClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithEvent:"), event)
	return VZKeyEventFromID(rv)
}

func NewVZKeyEventWithTypeKeyCode(type_ int64, code uint16) VZKeyEvent {
	instance := getVZKeyEventClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithType:keyCode:"), type_, code)
	return VZKeyEventFromID(rv)
}

func (v VZKeyEvent) InitWithEvent(event objectivec.IObject) VZKeyEvent {
	rv := objc.SendIfResponds[VZKeyEvent](v.ID, objc.Sel("initWithEvent:"), event)
	return rv
}
func (v VZKeyEvent) InitWithTypeKeyCode(type_ int64, code uint16) VZKeyEvent {
	rv := objc.SendIfResponds[VZKeyEvent](v.ID, objc.Sel("initWithType:keyCode:"), type_, code)
	return rv
}

func (v VZKeyEvent) KeyCode() uint16 {
	rv := objc.SendIfResponds[uint16](v.ID, objc.Sel("keyCode"))
	return rv
}
func (v VZKeyEvent) Type() int64 {
	rv := objc.SendIfResponds[int64](v.ID, objc.Sel("type"))
	return rv
}
