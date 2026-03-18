// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
)

// The class instance for the [NSGlassEffectContainerView] class.
var (
	_NSGlassEffectContainerViewClass     NSGlassEffectContainerViewClass
	_NSGlassEffectContainerViewClassOnce sync.Once
)

func getNSGlassEffectContainerViewClass() NSGlassEffectContainerViewClass {
	_NSGlassEffectContainerViewClassOnce.Do(func() {
		_NSGlassEffectContainerViewClass = NSGlassEffectContainerViewClass{class: objc.GetClass("NSGlassEffectContainerView")}
	})
	return _NSGlassEffectContainerViewClass
}

// GetNSGlassEffectContainerViewClass returns the class object for NSGlassEffectContainerView.
func GetNSGlassEffectContainerViewClass() NSGlassEffectContainerViewClass {
	return getNSGlassEffectContainerViewClass()
}

type NSGlassEffectContainerViewClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSGlassEffectContainerViewClass) Alloc() NSGlassEffectContainerView {
	rv := objc.Send[NSGlassEffectContainerView](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A view that efficiently merges descendant glass effect views together when
// they are within a specified proximity to each other.
//
// # Overview
//
// # Instance Properties
//
//   - [NSGlassEffectContainerView.ContentView]: The view that contains descendant views to merge together when in proximity to each other.
//   - [NSGlassEffectContainerView.SetContentView]
//   - [NSGlassEffectContainerView.Spacing]: The proximity at which the glass effect container view begins merging eligible descendent glass effect views.
//   - [NSGlassEffectContainerView.SetSpacing]
//
// See: https://developer.apple.com/documentation/AppKit/NSGlassEffectContainerView
type NSGlassEffectContainerView struct {
	NSView
}

// NSGlassEffectContainerViewFromID constructs a [NSGlassEffectContainerView] from an objc.ID.
//
// A view that efficiently merges descendant glass effect views together when
// they are within a specified proximity to each other.
func NSGlassEffectContainerViewFromID(id objc.ID) NSGlassEffectContainerView {
	return NSGlassEffectContainerView{NSView: NSViewFromID(id)}
}
// NOTE: NSGlassEffectContainerView adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSGlassEffectContainerView] class.
//
// # Instance Properties
//
//   - [INSGlassEffectContainerView.ContentView]: The view that contains descendant views to merge together when in proximity to each other.
//   - [INSGlassEffectContainerView.SetContentView]
//   - [INSGlassEffectContainerView.Spacing]: The proximity at which the glass effect container view begins merging eligible descendent glass effect views.
//   - [INSGlassEffectContainerView.SetSpacing]
//
// See: https://developer.apple.com/documentation/AppKit/NSGlassEffectContainerView
type INSGlassEffectContainerView interface {
	INSView

	// Topic: Instance Properties

	// The view that contains descendant views to merge together when in proximity to each other.
	ContentView() INSView
	SetContentView(value INSView)
	// The proximity at which the glass effect container view begins merging eligible descendent glass effect views.
	Spacing() float64
	SetSpacing(value float64)
}

// Init initializes the instance.
func (g NSGlassEffectContainerView) Init() NSGlassEffectContainerView {
	rv := objc.Send[NSGlassEffectContainerView](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g NSGlassEffectContainerView) Autorelease() NSGlassEffectContainerView {
	rv := objc.Send[NSGlassEffectContainerView](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSGlassEffectContainerView creates a new NSGlassEffectContainerView instance.
func NewNSGlassEffectContainerView() NSGlassEffectContainerView {
	class := getNSGlassEffectContainerViewClass()
	rv := objc.Send[NSGlassEffectContainerView](objc.ID(class.class), objc.Sel("new"))
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
func NewGlassEffectContainerViewWithCoder(coder foundation.INSCoder) NSGlassEffectContainerView {
	instance := getNSGlassEffectContainerViewClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSGlassEffectContainerViewFromID(rv)
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
func NewGlassEffectContainerViewWithFrame(frameRect corefoundation.CGRect) NSGlassEffectContainerView {
	instance := getNSGlassEffectContainerViewClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFrame:"), frameRect)
	return NSGlassEffectContainerViewFromID(rv)
}

// The view that contains descendant views to merge together when in proximity
// to each other.
//
// # Discussion
// 
// The glass effect container view does the following:
// 
// - Elevates the z-order of descendants of `contentView` to position them
// above the `contentView`. - Merges descendants together if the views are
// sufficiently similar and within the proximity specified in [Spacing]. -
// Processes similar glass effect views as a batch to improve performance.
//
// See: https://developer.apple.com/documentation/AppKit/NSGlassEffectContainerView/contentView
func (g NSGlassEffectContainerView) ContentView() INSView {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("contentView"))
	return NSViewFromID(objc.ID(rv))
}
func (g NSGlassEffectContainerView) SetContentView(value INSView) {
	objc.Send[struct{}](g.ID, objc.Sel("setContentView:"), value)
}

// The proximity at which the glass effect container view begins merging
// eligible descendent glass effect views.
//
// # Discussion
// 
// The default value, zero, is sufficient for batch processing eligible glass
// effect views, while avoiding distortion and merging effects for other views
// in close proximity.
//
// See: https://developer.apple.com/documentation/AppKit/NSGlassEffectContainerView/spacing
func (g NSGlassEffectContainerView) Spacing() float64 {
	rv := objc.Send[float64](g.ID, objc.Sel("spacing"))
	return rv
}
func (g NSGlassEffectContainerView) SetSpacing(value float64) {
	objc.Send[struct{}](g.ID, objc.Sel("setSpacing:"), value)
}

