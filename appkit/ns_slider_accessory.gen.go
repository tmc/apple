// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSSliderAccessory] class.
var (
	_NSSliderAccessoryClass     NSSliderAccessoryClass
	_NSSliderAccessoryClassOnce sync.Once
)

func getNSSliderAccessoryClass() NSSliderAccessoryClass {
	_NSSliderAccessoryClassOnce.Do(func() {
		_NSSliderAccessoryClass = NSSliderAccessoryClass{class: objc.GetClass("NSSliderAccessory")}
	})
	return _NSSliderAccessoryClass
}

// GetNSSliderAccessoryClass returns the class object for NSSliderAccessory.
func GetNSSliderAccessoryClass() NSSliderAccessoryClass {
	return getNSSliderAccessoryClass()
}

type NSSliderAccessoryClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSSliderAccessoryClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSSliderAccessoryClass) Alloc() NSSliderAccessory {
	rv := objc.Send[NSSliderAccessory](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// # Instance Properties
//
//   - [NSSliderAccessory.Behavior]: The effect on interaction with the accessory.
//   - [NSSliderAccessory.SetBehavior]
//   - [NSSliderAccessory.IsEnabled]
//   - [NSSliderAccessory.SetEnabled]
//
// See: https://developer.apple.com/documentation/AppKit/NSSliderAccessory
type NSSliderAccessory struct {
	objectivec.Object
}

// NSSliderAccessoryFromID constructs a [NSSliderAccessory] from an objc.ID.
func NSSliderAccessoryFromID(id objc.ID) NSSliderAccessory {
	return NSSliderAccessory{objectivec.Object{ID: id}}
}

// NOTE: NSSliderAccessory adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSSliderAccessory] class.
//
// # Instance Properties
//
//   - [INSSliderAccessory.Behavior]: The effect on interaction with the accessory.
//   - [INSSliderAccessory.SetBehavior]
//   - [INSSliderAccessory.IsEnabled]
//   - [INSSliderAccessory.SetEnabled]
//
// See: https://developer.apple.com/documentation/AppKit/NSSliderAccessory
type INSSliderAccessory interface {
	objectivec.IObject

	// Topic: Instance Properties

	// The effect on interaction with the accessory.
	Behavior() INSSliderAccessoryBehavior
	SetBehavior(value INSSliderAccessoryBehavior)
	IsEnabled() bool
	SetEnabled(value bool)

	InitWithCoder(coder foundation.INSCoder) NSSliderAccessory
	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (s NSSliderAccessory) Init() NSSliderAccessory {
	rv := objc.Send[NSSliderAccessory](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s NSSliderAccessory) Autorelease() NSSliderAccessory {
	rv := objc.Send[NSSliderAccessory](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSSliderAccessory creates a new NSSliderAccessory instance.
func NewNSSliderAccessory() NSSliderAccessory {
	class := getNSSliderAccessoryClass()
	rv := objc.Send[NSSliderAccessory](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/AppKit/NSSliderAccessory/init(coder:)
func NewSliderAccessoryWithCoder(coder foundation.INSCoder) NSSliderAccessory {
	instance := getNSSliderAccessoryClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSSliderAccessoryFromID(rv)
}

// See: https://developer.apple.com/documentation/AppKit/NSSliderAccessory/init(image:)
func NewSliderAccessoryWithImage(image INSImage) NSSliderAccessory {
	rv := objc.Send[objc.ID](objc.ID(getNSSliderAccessoryClass().class), objc.Sel("accessoryWithImage:"), image)
	return NSSliderAccessoryFromID(rv)
}

// See: https://developer.apple.com/documentation/AppKit/NSSliderAccessory/init(coder:)
func (s NSSliderAccessory) InitWithCoder(coder foundation.INSCoder) NSSliderAccessory {
	rv := objc.Send[NSSliderAccessory](s.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}
func (s NSSliderAccessory) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](s.ID, objc.Sel("encodeWithCoder:"), coder)
}

// The effect on interaction with the accessory.
//
// # Discussion
//
// The default value is `automaticBehavior`.
//
// See: https://developer.apple.com/documentation/AppKit/NSSliderAccessory/behavior
func (s NSSliderAccessory) Behavior() INSSliderAccessoryBehavior {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("behavior"))
	return NSSliderAccessoryBehaviorFromID(objc.ID(rv))
}
func (s NSSliderAccessory) SetBehavior(value INSSliderAccessoryBehavior) {
	objc.Send[struct{}](s.ID, objc.Sel("setBehavior:"), value)
}

// See: https://developer.apple.com/documentation/AppKit/NSSliderAccessory/isEnabled
func (s NSSliderAccessory) IsEnabled() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("isEnabled"))
	return rv
}
func (s NSSliderAccessory) SetEnabled(value bool) {
	objc.Send[struct{}](s.ID, objc.Sel("setEnabled:"), value)
}
