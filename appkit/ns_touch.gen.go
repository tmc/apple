// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSTouch] class.
var (
	_NSTouchClass     NSTouchClass
	_NSTouchClassOnce sync.Once
)

func getNSTouchClass() NSTouchClass {
	_NSTouchClassOnce.Do(func() {
		_NSTouchClass = NSTouchClass{class: objc.GetClass("NSTouch")}
	})
	return _NSTouchClass
}

// GetNSTouchClass returns the class object for NSTouch.
func GetNSTouchClass() NSTouchClass {
	return getNSTouchClass()
}

type NSTouchClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSTouchClass) Alloc() NSTouch {
	rv := objc.Send[NSTouch](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A snapshot of a particular touch at an instant in time.
//
// # Overview
// 
// A touch event is not persistent throughout the touch. A touch creates new
// instances as it progresses. Use the identity property to follow a specific
// touch across its lifetime.
// 
// Touches do not have a corresponding screen location. The first touch of a
// touch collection latches to the view underlying the cursor using the same
// hit detection as mouse events. Additional touches on the same device latch
// to the same view. Latches remain on views until the user ends a touch or an
// event cancels it.
//
// # Getting the Touch Type
//
//   - [NSTouch.Type]: A type of touch from a Touch Bar interaction.
//
// # Using Touch Properties
//
//   - [NSTouch.Identity]: The changes to a particular touch during its lifetime.
//   - [NSTouch.Phase]: The current phase of the touch.
//   - [NSTouch.NormalizedPosition]: The normalized position of the touch.
//   - [NSTouch.Resting]: The indicator for a resting touch.
//
// # Using Touch Device Properties
//
//   - [NSTouch.Device]: The digitizer that generates the touch. Useful to distinguish touches emanating from multiple-device scenarios.
//   - [NSTouch.DeviceSize]: The range of the touch device in points, such as 72 ppi.
//
// # Getting the Touch Location
//
//   - [NSTouch.LocationInView]: Indicates the location of the touch in the view’s coordinates.
//   - [NSTouch.PreviousLocationInView]: Indicates the previous location of the touch in the view’s coordinates.
//
// See: https://developer.apple.com/documentation/AppKit/NSTouch
type NSTouch struct {
	objectivec.Object
}

// NSTouchFromID constructs a [NSTouch] from an objc.ID.
//
// A snapshot of a particular touch at an instant in time.
func NSTouchFromID(id objc.ID) NSTouch {
	return NSTouch{objectivec.Object{ID: id}}
}
// NOTE: NSTouch adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSTouch] class.
//
// # Getting the Touch Type
//
//   - [INSTouch.Type]: A type of touch from a Touch Bar interaction.
//
// # Using Touch Properties
//
//   - [INSTouch.Identity]: The changes to a particular touch during its lifetime.
//   - [INSTouch.Phase]: The current phase of the touch.
//   - [INSTouch.NormalizedPosition]: The normalized position of the touch.
//   - [INSTouch.Resting]: The indicator for a resting touch.
//
// # Using Touch Device Properties
//
//   - [INSTouch.Device]: The digitizer that generates the touch. Useful to distinguish touches emanating from multiple-device scenarios.
//   - [INSTouch.DeviceSize]: The range of the touch device in points, such as 72 ppi.
//
// # Getting the Touch Location
//
//   - [INSTouch.LocationInView]: Indicates the location of the touch in the view’s coordinates.
//   - [INSTouch.PreviousLocationInView]: Indicates the previous location of the touch in the view’s coordinates.
//
// See: https://developer.apple.com/documentation/AppKit/NSTouch
type INSTouch interface {
	objectivec.IObject

	// Topic: Getting the Touch Type

	// A type of touch from a Touch Bar interaction.
	Type() NSTouchType

	// Topic: Using Touch Properties

	// The changes to a particular touch during its lifetime.
	Identity() objectivec.IObject
	// The current phase of the touch.
	Phase() NSTouchPhase
	// The normalized position of the touch.
	NormalizedPosition() corefoundation.CGPoint
	// The indicator for a resting touch.
	Resting() bool

	// Topic: Using Touch Device Properties

	// The digitizer that generates the touch. Useful to distinguish touches emanating from multiple-device scenarios.
	Device() objectivec.IObject
	// The range of the touch device in points, such as 72 ppi.
	DeviceSize() corefoundation.CGSize

	// Topic: Getting the Touch Location

	// Indicates the location of the touch in the view’s coordinates.
	LocationInView(view INSView) corefoundation.CGPoint
	// Indicates the previous location of the touch in the view’s coordinates.
	PreviousLocationInView(view INSView) corefoundation.CGPoint
}

// Init initializes the instance.
func (t NSTouch) Init() NSTouch {
	rv := objc.Send[NSTouch](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t NSTouch) Autorelease() NSTouch {
	rv := objc.Send[NSTouch](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSTouch creates a new NSTouch instance.
func NewNSTouch() NSTouch {
	class := getNSTouchClass()
	rv := objc.Send[NSTouch](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Indicates the location of the touch in the view’s coordinates.
//
// See: https://developer.apple.com/documentation/AppKit/NSTouch/location(in:)
func (t NSTouch) LocationInView(view INSView) corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](t.ID, objc.Sel("locationInView:"), view)
	return corefoundation.CGPoint(rv)
}
// Indicates the previous location of the touch in the view’s coordinates.
//
// See: https://developer.apple.com/documentation/AppKit/NSTouch/previousLocation(in:)
func (t NSTouch) PreviousLocationInView(view INSView) corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](t.ID, objc.Sel("previousLocationInView:"), view)
	return corefoundation.CGPoint(rv)
}

// A type of touch from a Touch Bar interaction.
//
// See: https://developer.apple.com/documentation/AppKit/NSTouch/type
func (t NSTouch) Type() NSTouchType {
	rv := objc.Send[NSTouchType](t.ID, objc.Sel("type"))
	return NSTouchType(rv)
}
// The changes to a particular touch during its lifetime.
//
// # Discussion
// 
// While touch identities may be re-used, they are unique during the life of
// the touch, even when multiple devices are present.
// 
// Identity objects implement the [NSCopying] protocol so that they may be
// used as keys in an [NSDictionary]. Use isEqual: to compare two touch
// identities.
//
// [NSCopying]: https://developer.apple.com/documentation/Foundation/NSCopying
// [NSDictionary]: https://developer.apple.com/documentation/Foundation/NSDictionary
//
// See: https://developer.apple.com/documentation/AppKit/NSTouch/identity
func (t NSTouch) Identity() objectivec.IObject {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("identity"))
	return objectivec.Object{ID: rv}
}
// The current phase of the touch.
//
// # Discussion
// 
// See [NSTouch.Phase] for possible values.
//
// [NSTouch.Phase]: https://developer.apple.com/documentation/AppKit/NSTouch/Phase-swift.struct
//
// See: https://developer.apple.com/documentation/AppKit/NSTouch/phase-swift.property
func (t NSTouch) Phase() NSTouchPhase {
	rv := objc.Send[NSTouchPhase](t.ID, objc.Sel("phase"))
	return NSTouchPhase(rv)
}
// The normalized position of the touch.
//
// # Discussion
// 
// The normalized position is a scaled value between (0.0) and (1.0,1.0),
// where (0.0,0.0) is the lower-left position on the touch device.
//
// See: https://developer.apple.com/documentation/AppKit/NSTouch/normalizedPosition
func (t NSTouch) NormalizedPosition() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](t.ID, objc.Sel("normalizedPosition"))
	return corefoundation.CGPoint(rv)
}
// The indicator for a resting touch.
//
// # Discussion
// 
// Resting touches occur when a user simply rests their thumb on the trackpad
// device.
//
// See: https://developer.apple.com/documentation/AppKit/NSTouch/isResting
func (t NSTouch) Resting() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("isResting"))
	return rv
}
// The digitizer that generates the touch. Useful to distinguish touches
// emanating from multiple-device scenarios.
//
// See: https://developer.apple.com/documentation/AppKit/NSTouch/device
func (t NSTouch) Device() objectivec.IObject {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("device"))
	return objectivec.Object{ID: rv}
}
// The range of the touch device in points, such as 72 ppi.
//
// # Discussion
// 
// The lower-left corner of the surface is considered (0,0).
//
// See: https://developer.apple.com/documentation/AppKit/NSTouch/deviceSize
func (t NSTouch) DeviceSize() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](t.ID, objc.Sel("deviceSize"))
	return corefoundation.CGSize(rv)
}

