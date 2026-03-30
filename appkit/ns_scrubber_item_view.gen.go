// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// The class instance for the [NSScrubberItemView] class.
var (
	_NSScrubberItemViewClass     NSScrubberItemViewClass
	_NSScrubberItemViewClassOnce sync.Once
)

func getNSScrubberItemViewClass() NSScrubberItemViewClass {
	_NSScrubberItemViewClassOnce.Do(func() {
		_NSScrubberItemViewClass = NSScrubberItemViewClass{class: objc.GetClass("NSScrubberItemView")}
	})
	return _NSScrubberItemViewClass
}

// GetNSScrubberItemViewClass returns the class object for NSScrubberItemView.
func GetNSScrubberItemViewClass() NSScrubberItemViewClass {
	return getNSScrubberItemViewClass()
}

type NSScrubberItemViewClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSScrubberItemViewClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSScrubberItemViewClass) Alloc() NSScrubberItemView {
	rv := objc.Send[NSScrubberItemView](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// An item at a specific index position in the scrubber.
//
// See: https://developer.apple.com/documentation/AppKit/NSScrubberItemView
type NSScrubberItemView struct {
	NSScrubberArrangedView
}

// NSScrubberItemViewFromID constructs a [NSScrubberItemView] from an objc.ID.
//
// An item at a specific index position in the scrubber.
func NSScrubberItemViewFromID(id objc.ID) NSScrubberItemView {
	return NSScrubberItemView{NSScrubberArrangedView: NSScrubberArrangedViewFromID(id)}
}

// NOTE: NSScrubberItemView adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSScrubberItemView] class.
//
// See: https://developer.apple.com/documentation/AppKit/NSScrubberItemView
type INSScrubberItemView interface {
	INSScrubberArrangedView
}

// Init initializes the instance.
func (s NSScrubberItemView) Init() NSScrubberItemView {
	rv := objc.Send[NSScrubberItemView](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s NSScrubberItemView) Autorelease() NSScrubberItemView {
	rv := objc.Send[NSScrubberItemView](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSScrubberItemView creates a new NSScrubberItemView instance.
func NewNSScrubberItemView() NSScrubberItemView {
	class := getNSScrubberItemViewClass()
	rv := objc.Send[NSScrubberItemView](objc.ID(class.class), objc.Sel("new"))
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
func NewScrubberItemViewWithCoder(coder foundation.INSCoder) NSScrubberItemView {
	instance := getNSScrubberItemViewClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSScrubberItemViewFromID(rv)
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
func NewScrubberItemViewWithFrame(frameRect corefoundation.CGRect) NSScrubberItemView {
	instance := getNSScrubberItemViewClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFrame:"), frameRect)
	return NSScrubberItemViewFromID(rv)
}
