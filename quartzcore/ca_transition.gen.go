// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/coregraphics"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CATransition] class.
var (
	_CATransitionClass     CATransitionClass
	_CATransitionClassOnce sync.Once
)

func getCATransitionClass() CATransitionClass {
	_CATransitionClassOnce.Do(func() {
		_CATransitionClass = CATransitionClass{class: objc.GetClass("CATransition")}
	})
	return _CATransitionClass
}

// GetCATransitionClass returns the class object for CATransition.
func GetCATransitionClass() CATransitionClass {
	return getCATransitionClass()
}

type CATransitionClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CATransitionClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CATransitionClass) Alloc() CATransition {
	rv := objc.Send[CATransition](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// An object that provides an animated transition between a layer’s states.
//
// # Overview
// 
// You can transition between a layer’s states by creating and adding a
// [CATransition] object to it. The default transition is a cross fade, but
// you can specify different effects from a set of predefined transitions.
// 
// The following code shows how you can transition between the two states of a
// [CATextLayer] named `transitioningLayer`. When the layer is first created,
// its [CATransition.BackgroundColor] is set to red and its [CATransition.String] property is set to
// [Red]. When the `runTransition()` function is called, a new [CATransition]
// object is created and added to `transitioningLayer`, and the state of the
// layer is changed so that its background color is blue and its rendered text
// reads [Blue].
// 
// The end result is that the push transition animates the red state from left
// to right with the blue state entering the scene from the left.
//
// # Transition start and end point
//
//   - [CATransition.StartProgress]: Indicates the start point of the receiver as a fraction of the entire transition.
//   - [CATransition.SetStartProgress]
//   - [CATransition.EndProgress]: Indicates the end point of the receiver as a fraction of the entire transition.
//   - [CATransition.SetEndProgress]
//
// # Transition Properties
//
//   - [CATransition.Type]: Specifies the predefined transition type.
//   - [CATransition.SetType]
//   - [CATransition.Subtype]: Specifies an optional subtype that indicates the direction for the predefined motion-based transitions.
//   - [CATransition.SetSubtype]
//
// # Custom transition filter
//
//   - [CATransition.Filter]: An optional Core Image filter object that provides the transition.
//   - [CATransition.SetFilter]
//
// See: https://developer.apple.com/documentation/QuartzCore/CATransition
type CATransition struct {
	CAAnimation
}

// CATransitionFromID constructs a [CATransition] from an objc.ID.
//
// An object that provides an animated transition between a layer’s states.
func CATransitionFromID(id objc.ID) CATransition {
	return CATransition{CAAnimation: CAAnimationFromID(id)}
}
// NOTE: CATransition adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CATransition] class.
//
// # Transition start and end point
//
//   - [ICATransition.StartProgress]: Indicates the start point of the receiver as a fraction of the entire transition.
//   - [ICATransition.SetStartProgress]
//   - [ICATransition.EndProgress]: Indicates the end point of the receiver as a fraction of the entire transition.
//   - [ICATransition.SetEndProgress]
//
// # Transition Properties
//
//   - [ICATransition.Type]: Specifies the predefined transition type.
//   - [ICATransition.SetType]
//   - [ICATransition.Subtype]: Specifies an optional subtype that indicates the direction for the predefined motion-based transitions.
//   - [ICATransition.SetSubtype]
//
// # Custom transition filter
//
//   - [ICATransition.Filter]: An optional Core Image filter object that provides the transition.
//   - [ICATransition.SetFilter]
//
// See: https://developer.apple.com/documentation/QuartzCore/CATransition
type ICATransition interface {
	ICAAnimation

	// Topic: Transition start and end point

	// Indicates the start point of the receiver as a fraction of the entire transition.
	StartProgress() float32
	SetStartProgress(value float32)
	// Indicates the end point of the receiver as a fraction of the entire transition.
	EndProgress() float32
	SetEndProgress(value float32)

	// Topic: Transition Properties

	// Specifies the predefined transition type.
	Type() CATransitionType
	SetType(value CATransitionType)
	// Specifies an optional subtype that indicates the direction for the predefined motion-based transitions.
	Subtype() CATransitionSubtype
	SetSubtype(value CATransitionSubtype)

	// Topic: Custom transition filter

	// An optional Core Image filter object that provides the transition.
	Filter() objectivec.IObject
	SetFilter(value objectivec.IObject)

	// The background color of the receiver. Animatable.
	BackgroundColor() coregraphics.CGColorRef
	SetBackgroundColor(value coregraphics.CGColorRef)
}

// Init initializes the instance.
func (t CATransition) Init() CATransition {
	rv := objc.Send[CATransition](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t CATransition) Autorelease() CATransition {
	rv := objc.Send[CATransition](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewCATransition creates a new CATransition instance.
func NewCATransition() CATransition {
	class := getCATransitionClass()
	rv := objc.Send[CATransition](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Indicates the start point of the receiver as a fraction of the entire
// transition.
//
// # Discussion
// 
// Legal values are numbers between 0.0 and 1.0. For example, to start the
// transition half way through its progress set `startProgress` to 0.5. The
// default value is 0.
//
// See: https://developer.apple.com/documentation/QuartzCore/CATransition/startProgress
func (t CATransition) StartProgress() float32 {
	rv := objc.Send[float32](t.ID, objc.Sel("startProgress"))
	return rv
}
func (t CATransition) SetStartProgress(value float32) {
	objc.Send[struct{}](t.ID, objc.Sel("setStartProgress:"), value)
}
// Indicates the end point of the receiver as a fraction of the entire
// transition.
//
// # Discussion
// 
// The value must be greater than or equal to [StartProgress], and not greater
// than 1.0. If `endProgress` is less than [StartProgress] the behavior is
// undefined. The default value is 1.0.
//
// See: https://developer.apple.com/documentation/QuartzCore/CATransition/endProgress
func (t CATransition) EndProgress() float32 {
	rv := objc.Send[float32](t.ID, objc.Sel("endProgress"))
	return rv
}
func (t CATransition) SetEndProgress(value float32) {
	objc.Send[struct{}](t.ID, objc.Sel("setEndProgress:"), value)
}
// Specifies the predefined transition type.
//
// # Discussion
// 
// The possible values are shown in [Common Transition Types]. This property
// is ignored if a custom transition is specified in the [Filter] property.
// The default is [fade].
//
// [Common Transition Types]: https://developer.apple.com/documentation/QuartzCore/common-transition-types
// [fade]: https://developer.apple.com/documentation/QuartzCore/CATransitionType/fade
//
// See: https://developer.apple.com/documentation/QuartzCore/CATransition/type
func (t CATransition) Type() CATransitionType {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("type"))
	return CATransitionType(foundation.NSStringFromID(rv).String())
}
func (t CATransition) SetType(value CATransitionType) {
	objc.Send[struct{}](t.ID, objc.Sel("setType:"), objc.String(string(value)))
}
// Specifies an optional subtype that indicates the direction for the
// predefined motion-based transitions.
//
// # Discussion
// 
// The possible values are shown in [Common Transition Subtypes]. The default
// is `nil`.
// 
// This property is ignored if a custom transition is specified in the
// [Filter] property.
//
// [Common Transition Subtypes]: https://developer.apple.com/documentation/QuartzCore/common-transition-subtypes
//
// See: https://developer.apple.com/documentation/QuartzCore/CATransition/subtype
func (t CATransition) Subtype() CATransitionSubtype {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("subtype"))
	return CATransitionSubtype(foundation.NSStringFromID(rv).String())
}
func (t CATransition) SetSubtype(value CATransitionSubtype) {
	objc.Send[struct{}](t.ID, objc.Sel("setSubtype:"), objc.String(string(value)))
}
// An optional Core Image filter object that provides the transition.
//
// # Discussion
// 
// If specified, the filter must support both [kCIInputImageKey] and
// [kCIInputTargetImageKey] input keys, and the [kCIOutputImageKey] output
// key. The filter may optionally support the [kCIInputExtentKey] input key,
// which is set to a rectangle describing the region in which the transition
// should run. If `filter` does not support the required input and output keys
// the behavior is undefined.
// 
// Defaults to `nil`. When a transition filter is specified the [Type] and
// [Subtype] properties are ignored.
// 
// The [NSView] that contains the transitioning layer must have its
// [layerUsesCoreImageFilters] set to [true].
// 
// The following code shows how you can transition between the two states of a
// [CATextLayer] named `transitioningLayer`. When the layer is first created,
// its [BackgroundColor] is set to red and its [String] property is set to
// [Red]. When the `runTransition()` function is called, a new [CATransition]
// object is created and added to `transitioningLayer`, and the state of the
// layer is changed so that its background color is blue and its rendered text
// reads [Blue].
// 
// The end result is that the transition animates from the red state to the
// blue state using a [CICopyMachineTransition] transition.
// 
// # Special Considerations
// 
// This property is not supported on layers in iOS.
//
// [CICopyMachineTransition]: https://developer.apple.com/library/archive/documentation/GraphicsImaging/Reference/CoreImageFilterReference/index.html#//apple_ref/doc/filter/ci/CICopyMachineTransition
// [NSView]: https://developer.apple.com/documentation/AppKit/NSView
// [kCIInputExtentKey]: https://developer.apple.com/documentation/CoreImage/kCIInputExtentKey
// [kCIInputImageKey]: https://developer.apple.com/documentation/CoreImage/kCIInputImageKey
// [kCIInputTargetImageKey]: https://developer.apple.com/documentation/CoreImage/kCIInputTargetImageKey
// [kCIOutputImageKey]: https://developer.apple.com/documentation/CoreImage/kCIOutputImageKey
// [layerUsesCoreImageFilters]: https://developer.apple.com/documentation/AppKit/NSView/layerUsesCoreImageFilters
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/QuartzCore/CATransition/filter
func (t CATransition) Filter() objectivec.IObject {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("filter"))
	return objectivec.Object{ID: rv}
}
func (t CATransition) SetFilter(value objectivec.IObject) {
	objc.Send[struct{}](t.ID, objc.Sel("setFilter:"), value)
}
// The background color of the receiver. Animatable.
//
// See: https://developer.apple.com/documentation/quartzcore/calayer/backgroundcolor
func (t CATransition) BackgroundColor() coregraphics.CGColorRef {
	rv := objc.Send[coregraphics.CGColorRef](t.ID, objc.Sel("backgroundColor"))
	return coregraphics.CGColorRef(rv)
}
func (t CATransition) SetBackgroundColor(value coregraphics.CGColorRef) {
	objc.Send[struct{}](t.ID, objc.Sel("setBackgroundColor:"), value)
}

