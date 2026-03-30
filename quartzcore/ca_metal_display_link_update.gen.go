// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CAMetalDisplayLinkUpdate] class.
var (
	_CAMetalDisplayLinkUpdateClass     CAMetalDisplayLinkUpdateClass
	_CAMetalDisplayLinkUpdateClassOnce sync.Once
)

func getCAMetalDisplayLinkUpdateClass() CAMetalDisplayLinkUpdateClass {
	_CAMetalDisplayLinkUpdateClassOnce.Do(func() {
		_CAMetalDisplayLinkUpdateClass = CAMetalDisplayLinkUpdateClass{class: objc.GetClass("CAMetalDisplayLinkUpdate")}
	})
	return _CAMetalDisplayLinkUpdateClass
}

// GetCAMetalDisplayLinkUpdateClass returns the class object for CAMetalDisplayLinkUpdate.
func GetCAMetalDisplayLinkUpdateClass() CAMetalDisplayLinkUpdateClass {
	return getCAMetalDisplayLinkUpdateClass()
}

type CAMetalDisplayLinkUpdateClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CAMetalDisplayLinkUpdateClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CAMetalDisplayLinkUpdateClass) Alloc() CAMetalDisplayLinkUpdate {
	rv := objc.Send[CAMetalDisplayLinkUpdate](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// Stores information about a single update from a Metal display link
// instance.
//
// # Timing the Next Animation Frame
//
//   - [CAMetalDisplayLinkUpdate.TargetPresentationTimestamp]: The time the system estimates until the display of the next frame.
//
// # Drawing the Next Frame
//
//   - [CAMetalDisplayLinkUpdate.TargetTimestamp]: A deadline that indicates when your app needs to finish rendering to the drawable.
//   - [CAMetalDisplayLinkUpdate.Drawable]: The Metal drawable your app uses to render the next frame.
//
// See: https://developer.apple.com/documentation/QuartzCore/CAMetalDisplayLink/Update
type CAMetalDisplayLinkUpdate struct {
	objectivec.Object
}

// CAMetalDisplayLinkUpdateFromID constructs a [CAMetalDisplayLinkUpdate] from an objc.ID.
//
// Stores information about a single update from a Metal display link
// instance.
func CAMetalDisplayLinkUpdateFromID(id objc.ID) CAMetalDisplayLinkUpdate {
	return CAMetalDisplayLinkUpdate{objectivec.Object{ID: id}}
}

// NOTE: CAMetalDisplayLinkUpdate adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CAMetalDisplayLinkUpdate] class.
//
// # Timing the Next Animation Frame
//
//   - [ICAMetalDisplayLinkUpdate.TargetPresentationTimestamp]: The time the system estimates until the display of the next frame.
//
// # Drawing the Next Frame
//
//   - [ICAMetalDisplayLinkUpdate.TargetTimestamp]: A deadline that indicates when your app needs to finish rendering to the drawable.
//   - [ICAMetalDisplayLinkUpdate.Drawable]: The Metal drawable your app uses to render the next frame.
//
// See: https://developer.apple.com/documentation/QuartzCore/CAMetalDisplayLink/Update
type ICAMetalDisplayLinkUpdate interface {
	objectivec.IObject

	// Topic: Timing the Next Animation Frame

	// The time the system estimates until the display of the next frame.
	TargetPresentationTimestamp() float64

	// Topic: Drawing the Next Frame

	// A deadline that indicates when your app needs to finish rendering to the drawable.
	TargetTimestamp() float64
	// The Metal drawable your app uses to render the next frame.
	Drawable() CAMetalDrawable
}

// Init initializes the instance.
func (m CAMetalDisplayLinkUpdate) Init() CAMetalDisplayLinkUpdate {
	rv := objc.Send[CAMetalDisplayLinkUpdate](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m CAMetalDisplayLinkUpdate) Autorelease() CAMetalDisplayLinkUpdate {
	rv := objc.Send[CAMetalDisplayLinkUpdate](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewCAMetalDisplayLinkUpdate creates a new CAMetalDisplayLinkUpdate instance.
func NewCAMetalDisplayLinkUpdate() CAMetalDisplayLinkUpdate {
	class := getCAMetalDisplayLinkUpdateClass()
	rv := objc.Send[CAMetalDisplayLinkUpdate](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The time the system estimates until the display of the next frame.
//
// # Discussion
//
// Update your animations based on the time difference between this timestamp
// and the previous timestamp.
//
// See: https://developer.apple.com/documentation/QuartzCore/CAMetalDisplayLink/Update/targetPresentationTimestamp
func (m CAMetalDisplayLinkUpdate) TargetPresentationTimestamp() float64 {
	rv := objc.Send[float64](m.ID, objc.Sel("targetPresentationTimestamp"))
	return rv
}

// A deadline that indicates when your app needs to finish rendering to the
// drawable.
//
// # Discussion
//
// Your app needs to call the [Drawable] instance’s [present()] method
// before the deadline. GPU rendering can continue after this time, based on
// [PreferredFrameLatency]. For more information on timing your app’s
// rendering, see [MetalDisplayLinkNeedsUpdate].
//
// See: https://developer.apple.com/documentation/QuartzCore/CAMetalDisplayLink/Update/targetTimestamp
//
// [present()]: https://developer.apple.com/documentation/Metal/MTLDrawable/present()
func (m CAMetalDisplayLinkUpdate) TargetTimestamp() float64 {
	rv := objc.Send[float64](m.ID, objc.Sel("targetTimestamp"))
	return rv
}

// The Metal drawable your app uses to render the next frame.
//
// See: https://developer.apple.com/documentation/QuartzCore/CAMetalDisplayLink/Update/drawable
func (m CAMetalDisplayLinkUpdate) Drawable() CAMetalDrawable {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("drawable"))
	return CAMetalDrawableObjectFromID(rv)
}
