// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"sync"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CAScrollLayer] class.
var (
	_CAScrollLayerClass     CAScrollLayerClass
	_CAScrollLayerClassOnce sync.Once
)

func getCAScrollLayerClass() CAScrollLayerClass {
	_CAScrollLayerClassOnce.Do(func() {
		_CAScrollLayerClass = CAScrollLayerClass{class: objc.GetClass("CAScrollLayer")}
	})
	return _CAScrollLayerClass
}

// GetCAScrollLayerClass returns the class object for CAScrollLayer.
func GetCAScrollLayerClass() CAScrollLayerClass {
	return getCAScrollLayerClass()
}

type CAScrollLayerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CAScrollLayerClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CAScrollLayerClass) Alloc() CAScrollLayer {
	rv := objc.Send[CAScrollLayer](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// A layer that displays scrollable content larger than its own bounds.
//
// # Overview
//
// The [CAScrollLayer] class is a subclass of [CALayer] that simplifies
// displaying a portion of a layer. The extent of the scrollable area of the
// [CAScrollLayer] is defined by the layout of its sublayers. The visible
// portion of the layer content is set by specifying the origin as a point or
// a rectangular area of the contents to be displayed. [CAScrollLayer] does
// not provide keyboard or mouse event-handling, nor does it provide visible
// scrollers.
//
// # Scrolling constraints
//
//   - [CAScrollLayer.ScrollMode]: Defines the axes in which the layer may be scrolled.
//   - [CAScrollLayer.SetScrollMode]
//
// # Scrolling the layer
//
//   - [CAScrollLayer.ScrollToPoint]: Changes the origin of the receiver to the specified point.
//   - [CAScrollLayer.ScrollToRect]: Scroll the contents of the receiver to ensure that the rectangle is visible.
//
// See: https://developer.apple.com/documentation/QuartzCore/CAScrollLayer
type CAScrollLayer struct {
	CALayer
}

// CAScrollLayerFromID constructs a [CAScrollLayer] from an objc.ID.
//
// A layer that displays scrollable content larger than its own bounds.
func CAScrollLayerFromID(id objc.ID) CAScrollLayer {
	return CAScrollLayer{CALayer: CALayerFromID(id)}
}

// NOTE: CAScrollLayer adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CAScrollLayer] class.
//
// # Scrolling constraints
//
//   - [ICAScrollLayer.ScrollMode]: Defines the axes in which the layer may be scrolled.
//   - [ICAScrollLayer.SetScrollMode]
//
// # Scrolling the layer
//
//   - [ICAScrollLayer.ScrollToPoint]: Changes the origin of the receiver to the specified point.
//   - [ICAScrollLayer.ScrollToRect]: Scroll the contents of the receiver to ensure that the rectangle is visible.
//
// See: https://developer.apple.com/documentation/QuartzCore/CAScrollLayer
type ICAScrollLayer interface {
	ICALayer

	// Topic: Scrolling constraints

	// Defines the axes in which the layer may be scrolled.
	ScrollMode() CAScrollLayerScrollMode
	SetScrollMode(value CAScrollLayerScrollMode)

	// Topic: Scrolling the layer

	// Changes the origin of the receiver to the specified point.
	ScrollToPoint(p corefoundation.CGPoint)
	// Scroll the contents of the receiver to ensure that the rectangle is visible.
	ScrollToRect(r corefoundation.CGRect)
}

// Init initializes the instance.
func (s CAScrollLayer) Init() CAScrollLayer {
	rv := objc.Send[CAScrollLayer](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s CAScrollLayer) Autorelease() CAScrollLayer {
	rv := objc.Send[CAScrollLayer](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewCAScrollLayer creates a new CAScrollLayer instance.
func NewCAScrollLayer() CAScrollLayer {
	class := getCAScrollLayerClass()
	rv := objc.Send[CAScrollLayer](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Override to copy or initialize custom fields of the specified layer.
//
// layer: The layer from which custom fields should be copied.
//
// # Return Value
//
// A layer instance with any custom instance variables copied from `layer`.
//
// # Discussion
//
// This initializer is used to create shadow copies of layers, for example,
// for the [PresentationLayer] method. Using this method in any other
// situation will produce undefined behavior. For example, do not use this
// method to initialize a new layer with an existing layer’s content.
//
// If you are implementing a custom layer subclass, you can override this
// method and use it to copy the values of instance variables into the new
// object. Subclasses should always invoke the superclass implementation.
//
// This method is the designated initializer for layer objects in the
// presentation layer.
//
// See: https://developer.apple.com/documentation/QuartzCore/CALayer/init(layer:)
func NewScrollLayerWithLayer(layer objectivec.IObject) CAScrollLayer {
	instance := getCAScrollLayerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithLayer:"), layer)
	return CAScrollLayerFromID(rv)
}

// Changes the origin of the receiver to the specified point.
//
// p: The new origin.
//
// See: https://developer.apple.com/documentation/QuartzCore/CAScrollLayer/scroll(to:)-37q0p
func (s CAScrollLayer) ScrollToPoint(p corefoundation.CGPoint) {
	objc.Send[objc.ID](s.ID, objc.Sel("scrollToPoint:"), p)
}

// Scroll the contents of the receiver to ensure that the rectangle is
// visible.
//
// r: The rectangle that should be visible.
//
// See: https://developer.apple.com/documentation/QuartzCore/CAScrollLayer/scroll(to:)-782vd
func (s CAScrollLayer) ScrollToRect(r corefoundation.CGRect) {
	objc.Send[objc.ID](s.ID, objc.Sel("scrollToRect:"), r)
}

// Defines the axes in which the layer may be scrolled.
//
// # Discussion
//
// The possible values are described in [Scroll Modes]. The default is [both].
//
// See: https://developer.apple.com/documentation/QuartzCore/CAScrollLayer/scrollMode
//
// [Scroll Modes]: https://developer.apple.com/documentation/QuartzCore/scroll-modes
// [both]: https://developer.apple.com/documentation/QuartzCore/CAScrollLayerScrollMode/both
func (s CAScrollLayer) ScrollMode() CAScrollLayerScrollMode {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("scrollMode"))
	return CAScrollLayerScrollMode(foundation.NSStringFromID(rv).String())
}
func (s CAScrollLayer) SetScrollMode(value CAScrollLayerScrollMode) {
	objc.Send[struct{}](s.ID, objc.Sel("setScrollMode:"), objc.String(string(value)))
}
