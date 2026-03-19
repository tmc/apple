// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSGlassEffectView] class.
var (
	_NSGlassEffectViewClass     NSGlassEffectViewClass
	_NSGlassEffectViewClassOnce sync.Once
)

func getNSGlassEffectViewClass() NSGlassEffectViewClass {
	_NSGlassEffectViewClassOnce.Do(func() {
		_NSGlassEffectViewClass = NSGlassEffectViewClass{class: objc.GetClass("NSGlassEffectView")}
	})
	return _NSGlassEffectViewClass
}

// GetNSGlassEffectViewClass returns the class object for NSGlassEffectView.
func GetNSGlassEffectViewClass() NSGlassEffectViewClass {
	return getNSGlassEffectViewClass()
}

type NSGlassEffectViewClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSGlassEffectViewClass) Alloc() NSGlassEffectView {
	rv := objc.Send[NSGlassEffectView](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A view that embeds its content view in a dynamic glass effect.
//
// # Instance Properties
//
//   - [NSGlassEffectView.ContentView]: The view to embed in glass.
//   - [NSGlassEffectView.SetContentView]
//   - [NSGlassEffectView.CornerRadius]: The amount of curvature for all corners of the glass.
//   - [NSGlassEffectView.SetCornerRadius]
//   - [NSGlassEffectView.Style]: The style of glass this view uses.
//   - [NSGlassEffectView.SetStyle]
//   - [NSGlassEffectView.TintColor]: The color the glass effect view uses to tint the background and glass effect toward.
//   - [NSGlassEffectView.SetTintColor]
//
// See: https://developer.apple.com/documentation/AppKit/NSGlassEffectView
type NSGlassEffectView struct {
	NSView
}

// NSGlassEffectViewFromID constructs a [NSGlassEffectView] from an objc.ID.
//
// A view that embeds its content view in a dynamic glass effect.
func NSGlassEffectViewFromID(id objc.ID) NSGlassEffectView {
	return NSGlassEffectView{NSView: NSViewFromID(id)}
}
// NOTE: NSGlassEffectView adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSGlassEffectView] class.
//
// # Instance Properties
//
//   - [INSGlassEffectView.ContentView]: The view to embed in glass.
//   - [INSGlassEffectView.SetContentView]
//   - [INSGlassEffectView.CornerRadius]: The amount of curvature for all corners of the glass.
//   - [INSGlassEffectView.SetCornerRadius]
//   - [INSGlassEffectView.Style]: The style of glass this view uses.
//   - [INSGlassEffectView.SetStyle]
//   - [INSGlassEffectView.TintColor]: The color the glass effect view uses to tint the background and glass effect toward.
//   - [INSGlassEffectView.SetTintColor]
//
// See: https://developer.apple.com/documentation/AppKit/NSGlassEffectView
type INSGlassEffectView interface {
	INSView

	// Topic: Instance Properties

	// The view to embed in glass.
	ContentView() INSView
	SetContentView(value INSView)
	// The amount of curvature for all corners of the glass.
	CornerRadius() float64
	SetCornerRadius(value float64)
	// The style of glass this view uses.
	Style() objectivec.IObject
	SetStyle(value objectivec.IObject)
	// The color the glass effect view uses to tint the background and glass effect toward.
	TintColor() INSColor
	SetTintColor(value INSColor)
}

// Init initializes the instance.
func (g NSGlassEffectView) Init() NSGlassEffectView {
	rv := objc.Send[NSGlassEffectView](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g NSGlassEffectView) Autorelease() NSGlassEffectView {
	rv := objc.Send[NSGlassEffectView](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSGlassEffectView creates a new NSGlassEffectView instance.
func NewNSGlassEffectView() NSGlassEffectView {
	class := getNSGlassEffectViewClass()
	rv := objc.Send[NSGlassEffectView](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Initializes a view using from data in the specified coder object.
//
// coder: The coder object that contains the view’s configuration details.
//
// # Return Value
// 
// An initialized view or `nil` if AppKit couldn’t create the object.
//
// See: https://developer.apple.com/documentation/AppKit/NSView/init(coder:)
func NewGlassEffectViewWithCoder(coder foundation.INSCoder) NSGlassEffectView {
	instance := getNSGlassEffectViewClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSGlassEffectViewFromID(rv)
}

// Initializes and returns a newly allocated [NSView] object with a specified
// frame rectangle.
//
// frameRect: The frame rectangle for the created view object.
//
// # Return Value
// 
// An initialized view or `nil` if AppKit couldn’t create the object.
//
// # Discussion
// 
// Insert the view into your window’s view hieararchy before you can do
// anything with it. This method is the designated initializer for the
// [NSView] class.
//
// See: https://developer.apple.com/documentation/AppKit/NSView/init(frame:)
func NewGlassEffectViewWithFrame(frameRect corefoundation.CGRect) NSGlassEffectView {
	instance := getNSGlassEffectViewClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFrame:"), frameRect)
	return NSGlassEffectViewFromID(rv)
}

// The view to embed in glass.
//
// # Discussion
//
// See: https://developer.apple.com/documentation/AppKit/NSGlassEffectView/contentView
func (g NSGlassEffectView) ContentView() INSView {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("contentView"))
	return NSViewFromID(objc.ID(rv))
}
func (g NSGlassEffectView) SetContentView(value INSView) {
	objc.Send[struct{}](g.ID, objc.Sel("setContentView:"), value)
}
// The amount of curvature for all corners of the glass.
//
// See: https://developer.apple.com/documentation/AppKit/NSGlassEffectView/cornerRadius
func (g NSGlassEffectView) CornerRadius() float64 {
	rv := objc.Send[float64](g.ID, objc.Sel("cornerRadius"))
	return rv
}
func (g NSGlassEffectView) SetCornerRadius(value float64) {
	objc.Send[struct{}](g.ID, objc.Sel("setCornerRadius:"), value)
}
// The style of glass this view uses.
//
// See: https://developer.apple.com/documentation/AppKit/NSGlassEffectView/style-swift.property
func (g NSGlassEffectView) Style() objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("style"))
	return objectivec.Object{ID: rv}
}
func (g NSGlassEffectView) SetStyle(value objectivec.IObject) {
	objc.Send[struct{}](g.ID, objc.Sel("setStyle:"), value)
}
// The color the glass effect view uses to tint the background and glass
// effect toward.
//
// See: https://developer.apple.com/documentation/AppKit/NSGlassEffectView/tintColor
func (g NSGlassEffectView) TintColor() INSColor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("tintColor"))
	return NSColorFromID(objc.ID(rv))
}
func (g NSGlassEffectView) SetTintColor(value INSColor) {
	objc.Send[struct{}](g.ID, objc.Sel("setTintColor:"), value)
}

