// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
)

// The class instance for the [NSScrubberArrangedView] class.
var (
	_NSScrubberArrangedViewClass     NSScrubberArrangedViewClass
	_NSScrubberArrangedViewClassOnce sync.Once
)

func getNSScrubberArrangedViewClass() NSScrubberArrangedViewClass {
	_NSScrubberArrangedViewClassOnce.Do(func() {
		_NSScrubberArrangedViewClass = NSScrubberArrangedViewClass{class: objc.GetClass("NSScrubberArrangedView")}
	})
	return _NSScrubberArrangedViewClass
}

// GetNSScrubberArrangedViewClass returns the class object for NSScrubberArrangedView.
func GetNSScrubberArrangedViewClass() NSScrubberArrangedViewClass {
	return getNSScrubberArrangedViewClass()
}

type NSScrubberArrangedViewClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSScrubberArrangedViewClass) Alloc() NSScrubberArrangedView {
	rv := objc.Send[NSScrubberArrangedView](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// An abstract base class for the views whose layout is managed by a scrubber.
//
// # Managing selection and highlighting
//
//   - [NSScrubberArrangedView.Highlighted]: A Boolean value that specifies whether the view is currently highlighted.
//   - [NSScrubberArrangedView.SetHighlighted]
//   - [NSScrubberArrangedView.Selected]: A Boolean value that specifies whether the current view is selected.
//   - [NSScrubberArrangedView.SetSelected]
//
// # Controlling the layout
//
//   - [NSScrubberArrangedView.ApplyLayoutAttributes]: Updates the layout of the arranged view to respect the provided layout attributes.
//
// See: https://developer.apple.com/documentation/AppKit/NSScrubberArrangedView
type NSScrubberArrangedView struct {
	NSView
}

// NSScrubberArrangedViewFromID constructs a [NSScrubberArrangedView] from an objc.ID.
//
// An abstract base class for the views whose layout is managed by a scrubber.
func NSScrubberArrangedViewFromID(id objc.ID) NSScrubberArrangedView {
	return NSScrubberArrangedView{NSView: NSViewFromID(id)}
}
// NOTE: NSScrubberArrangedView adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [NSScrubberArrangedView] class.
//
// # Managing selection and highlighting
//
//   - [INSScrubberArrangedView.Highlighted]: A Boolean value that specifies whether the view is currently highlighted.
//   - [INSScrubberArrangedView.SetHighlighted]
//   - [INSScrubberArrangedView.Selected]: A Boolean value that specifies whether the current view is selected.
//   - [INSScrubberArrangedView.SetSelected]
//
// # Controlling the layout
//
//   - [INSScrubberArrangedView.ApplyLayoutAttributes]: Updates the layout of the arranged view to respect the provided layout attributes.
//
// See: https://developer.apple.com/documentation/AppKit/NSScrubberArrangedView
type INSScrubberArrangedView interface {
	INSView

	// Topic: Managing selection and highlighting

	// A Boolean value that specifies whether the view is currently highlighted.
	Highlighted() bool
	SetHighlighted(value bool)
	// A Boolean value that specifies whether the current view is selected.
	Selected() bool
	SetSelected(value bool)

	// Topic: Controlling the layout

	// Updates the layout of the arranged view to respect the provided layout attributes.
	ApplyLayoutAttributes(layoutAttributes INSScrubberLayoutAttributes)

	EncodeWithCoder(coder foundation.INSCoder)
}





// Init initializes the instance.
func (s NSScrubberArrangedView) Init() NSScrubberArrangedView {
	rv := objc.Send[NSScrubberArrangedView](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s NSScrubberArrangedView) Autorelease() NSScrubberArrangedView {
	rv := objc.Send[NSScrubberArrangedView](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSScrubberArrangedView creates a new NSScrubberArrangedView instance.
func NewNSScrubberArrangedView() NSScrubberArrangedView {
	class := getNSScrubberArrangedViewClass()
	rv := objc.Send[NSScrubberArrangedView](objc.ID(class.class), objc.Sel("new"))
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
func NewScrubberArrangedViewWithCoder(coder foundation.INSCoder) NSScrubberArrangedView {
	instance := getNSScrubberArrangedViewClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSScrubberArrangedViewFromID(rv)
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
func NewScrubberArrangedViewWithFrame(frameRect corefoundation.CGRect) NSScrubberArrangedView {
	instance := getNSScrubberArrangedViewClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFrame:"), frameRect)
	return NSScrubberArrangedViewFromID(rv)
}







// Updates the layout of the arranged view to respect the provided layout
// attributes.
//
// layoutAttributes: An object that contains the parameters that this arranged view uses to
// determine its layout.
//
// See: https://developer.apple.com/documentation/AppKit/NSScrubberArrangedView/apply(_:)
func (s NSScrubberArrangedView) ApplyLayoutAttributes(layoutAttributes INSScrubberLayoutAttributes) {
	objc.Send[objc.ID](s.ID, objc.Sel("applyLayoutAttributes:"), layoutAttributes)
}
func (s NSScrubberArrangedView) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](s.ID, objc.Sel("encodeWithCoder:"), coder)
}












// A Boolean value that specifies whether the view is currently highlighted.
//
// See: https://developer.apple.com/documentation/AppKit/NSScrubberArrangedView/isHighlighted
func (s NSScrubberArrangedView) Highlighted() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("isHighlighted"))
	return rv
}
func (s NSScrubberArrangedView) SetHighlighted(value bool) {
	objc.Send[struct{}](s.ID, objc.Sel("setHighlighted:"), value)
}



// A Boolean value that specifies whether the current view is selected.
//
// See: https://developer.apple.com/documentation/AppKit/NSScrubberArrangedView/isSelected
func (s NSScrubberArrangedView) Selected() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("isSelected"))
	return rv
}
func (s NSScrubberArrangedView) SetSelected(value bool) {
	objc.Send[struct{}](s.ID, objc.Sel("setSelected:"), value)
}



































