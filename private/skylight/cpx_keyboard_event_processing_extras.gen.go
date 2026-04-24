// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CPXKeyboardEventProcessingExtras] class.
var (
	_CPXKeyboardEventProcessingExtrasClass     CPXKeyboardEventProcessingExtrasClass
	_CPXKeyboardEventProcessingExtrasClassOnce sync.Once
)

func getCPXKeyboardEventProcessingExtrasClass() CPXKeyboardEventProcessingExtrasClass {
	_CPXKeyboardEventProcessingExtrasClassOnce.Do(func() {
		_CPXKeyboardEventProcessingExtrasClass = CPXKeyboardEventProcessingExtrasClass{class: objc.GetClass("CPXKeyboardEventProcessingExtras")}
	})
	return _CPXKeyboardEventProcessingExtrasClass
}

// GetCPXKeyboardEventProcessingExtrasClass returns the class object for CPXKeyboardEventProcessingExtras.
func GetCPXKeyboardEventProcessingExtrasClass() CPXKeyboardEventProcessingExtrasClass {
	return getCPXKeyboardEventProcessingExtrasClass()
}

type CPXKeyboardEventProcessingExtrasClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CPXKeyboardEventProcessingExtrasClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CPXKeyboardEventProcessingExtrasClass) Alloc() CPXKeyboardEventProcessingExtras {
	rv := objc.Send[CPXKeyboardEventProcessingExtras](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [CPXKeyboardEventProcessingExtras.AppendDescriptionToStream]
//   - [CPXKeyboardEventProcessingExtras.ApplyToEvent]
//   - [CPXKeyboardEventProcessingExtras.MainDisplayHeight]
//   - [CPXKeyboardEventProcessingExtras.Window]
//   - [CPXKeyboardEventProcessingExtras.WindowHeight]
//   - [CPXKeyboardEventProcessingExtras.InitWithWindowWindowHeightMainDisplayHeight]
//   - [CPXKeyboardEventProcessingExtras.DebugDescription]
//   - [CPXKeyboardEventProcessingExtras.Description]
//   - [CPXKeyboardEventProcessingExtras.Hash]
//   - [CPXKeyboardEventProcessingExtras.Superclass]
//
// See: https://developer.apple.com/documentation/SkyLight/CPXKeyboardEventProcessingExtras
type CPXKeyboardEventProcessingExtras struct {
	objectivec.Object
}

// CPXKeyboardEventProcessingExtrasFromID constructs a [CPXKeyboardEventProcessingExtras] from an objc.ID.
func CPXKeyboardEventProcessingExtrasFromID(id objc.ID) CPXKeyboardEventProcessingExtras {
	return CPXKeyboardEventProcessingExtras{objectivec.Object{ID: id}}
}

// Ensure CPXKeyboardEventProcessingExtras implements ICPXKeyboardEventProcessingExtras.
var _ ICPXKeyboardEventProcessingExtras = CPXKeyboardEventProcessingExtras{}

// An interface definition for the [CPXKeyboardEventProcessingExtras] class.
//
// # Methods
//
//   - [ICPXKeyboardEventProcessingExtras.AppendDescriptionToStream]
//   - [ICPXKeyboardEventProcessingExtras.ApplyToEvent]
//   - [ICPXKeyboardEventProcessingExtras.MainDisplayHeight]
//   - [ICPXKeyboardEventProcessingExtras.Window]
//   - [ICPXKeyboardEventProcessingExtras.WindowHeight]
//   - [ICPXKeyboardEventProcessingExtras.InitWithWindowWindowHeightMainDisplayHeight]
//   - [ICPXKeyboardEventProcessingExtras.DebugDescription]
//   - [ICPXKeyboardEventProcessingExtras.Description]
//   - [ICPXKeyboardEventProcessingExtras.Hash]
//   - [ICPXKeyboardEventProcessingExtras.Superclass]
//
// See: https://developer.apple.com/documentation/SkyLight/CPXKeyboardEventProcessingExtras
type ICPXKeyboardEventProcessingExtras interface {
	objectivec.IObject

	// Topic: Methods

	AppendDescriptionToStream(stream objectivec.IObject)
	ApplyToEvent(event *SLSEventRecordRef)
	MainDisplayHeight() uint16
	Window() uint32
	WindowHeight() uint16
	InitWithWindowWindowHeightMainDisplayHeight(window uint32, height uint16, height2 uint16) CPXKeyboardEventProcessingExtras
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objc.Class
}

// Init initializes the instance.
func (c CPXKeyboardEventProcessingExtras) Init() CPXKeyboardEventProcessingExtras {
	rv := objc.Send[CPXKeyboardEventProcessingExtras](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CPXKeyboardEventProcessingExtras) Autorelease() CPXKeyboardEventProcessingExtras {
	rv := objc.Send[CPXKeyboardEventProcessingExtras](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCPXKeyboardEventProcessingExtras creates a new CPXKeyboardEventProcessingExtras instance.
func NewCPXKeyboardEventProcessingExtras() CPXKeyboardEventProcessingExtras {
	class := getCPXKeyboardEventProcessingExtrasClass()
	rv := objc.Send[CPXKeyboardEventProcessingExtras](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/CPXKeyboardEventProcessingExtras/initWithWindow:windowHeight:mainDisplayHeight:
func NewCPXKeyboardEventProcessingExtrasWithWindowWindowHeightMainDisplayHeight(window uint32, height uint16, height2 uint16) CPXKeyboardEventProcessingExtras {
	instance := getCPXKeyboardEventProcessingExtrasClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithWindow:windowHeight:mainDisplayHeight:"), window, height, height2)
	return CPXKeyboardEventProcessingExtrasFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/CPXKeyboardEventProcessingExtras/appendDescriptionToStream:
func (c CPXKeyboardEventProcessingExtras) AppendDescriptionToStream(stream objectivec.IObject) {
	objc.Send[objc.ID](c.ID, objc.Sel("appendDescriptionToStream:"), stream)
}

// See: https://developer.apple.com/documentation/SkyLight/CPXKeyboardEventProcessingExtras/applyToEvent:
func (c CPXKeyboardEventProcessingExtras) ApplyToEvent(event *SLSEventRecordRef) {
	objc.Send[objc.ID](c.ID, objc.Sel("applyToEvent:"), event)
}

// See: https://developer.apple.com/documentation/SkyLight/CPXKeyboardEventProcessingExtras/initWithWindow:windowHeight:mainDisplayHeight:
func (c CPXKeyboardEventProcessingExtras) InitWithWindowWindowHeightMainDisplayHeight(window uint32, height uint16, height2 uint16) CPXKeyboardEventProcessingExtras {
	rv := objc.Send[CPXKeyboardEventProcessingExtras](c.ID, objc.Sel("initWithWindow:windowHeight:mainDisplayHeight:"), window, height, height2)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/CPXKeyboardEventProcessingExtras/debugDescription
func (c CPXKeyboardEventProcessingExtras) DebugDescription() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/SkyLight/CPXKeyboardEventProcessingExtras/description
func (c CPXKeyboardEventProcessingExtras) Description() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/SkyLight/CPXKeyboardEventProcessingExtras/hash
func (c CPXKeyboardEventProcessingExtras) Hash() uint64 {
	rv := objc.Send[uint64](c.ID, objc.Sel("hash"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/CPXKeyboardEventProcessingExtras/mainDisplayHeight
func (c CPXKeyboardEventProcessingExtras) MainDisplayHeight() uint16 {
	rv := objc.Send[uint16](c.ID, objc.Sel("mainDisplayHeight"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/CPXKeyboardEventProcessingExtras/superclass
func (c CPXKeyboardEventProcessingExtras) Superclass() objc.Class {
	rv := objc.Send[objc.Class](c.ID, objc.Sel("superclass"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/CPXKeyboardEventProcessingExtras/window
func (c CPXKeyboardEventProcessingExtras) Window() uint32 {
	rv := objc.Send[uint32](c.ID, objc.Sel("window"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/CPXKeyboardEventProcessingExtras/windowHeight
func (c CPXKeyboardEventProcessingExtras) WindowHeight() uint16 {
	rv := objc.Send[uint16](c.ID, objc.Sel("windowHeight"))
	return rv
}
