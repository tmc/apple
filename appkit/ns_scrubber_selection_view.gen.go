// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
)

// The class instance for the [NSScrubberSelectionView] class.
var (
	_NSScrubberSelectionViewClass     NSScrubberSelectionViewClass
	_NSScrubberSelectionViewClassOnce sync.Once
)

func getNSScrubberSelectionViewClass() NSScrubberSelectionViewClass {
	_NSScrubberSelectionViewClassOnce.Do(func() {
		_NSScrubberSelectionViewClass = NSScrubberSelectionViewClass{class: objc.GetClass("NSScrubberSelectionView")}
	})
	return _NSScrubberSelectionViewClass
}

// GetNSScrubberSelectionViewClass returns the class object for NSScrubberSelectionView.
func GetNSScrubberSelectionViewClass() NSScrubberSelectionViewClass {
	return getNSScrubberSelectionViewClass()
}

type NSScrubberSelectionViewClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSScrubberSelectionViewClass) Alloc() NSScrubberSelectionView {
	rv := objc.Send[NSScrubberSelectionView](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// An abstract base class for specifying the appearance of a highlighted or
// selected item in a scrubber.
//
// # Overview
// 
// Create a subclass to customize the selection or highlight appearance of an
// item in your scrubber control. You need to return an instance of your
// subclass from the [SelectionView] method on [NSScrubberSelectionStyle].
//
// See: https://developer.apple.com/documentation/AppKit/NSScrubberSelectionView
type NSScrubberSelectionView struct {
	NSScrubberArrangedView
}

// NSScrubberSelectionViewFromID constructs a [NSScrubberSelectionView] from an objc.ID.
//
// An abstract base class for specifying the appearance of a highlighted or
// selected item in a scrubber.
func NSScrubberSelectionViewFromID(id objc.ID) NSScrubberSelectionView {
	return NSScrubberSelectionView{NSScrubberArrangedView: NSScrubberArrangedViewFromID(id)}
}
// NOTE: NSScrubberSelectionView adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [NSScrubberSelectionView] class.
//
// See: https://developer.apple.com/documentation/AppKit/NSScrubberSelectionView
type INSScrubberSelectionView interface {
	INSScrubberArrangedView

	EncodeWithCoder(coder foundation.INSCoder)
}





// Init initializes the instance.
func (s NSScrubberSelectionView) Init() NSScrubberSelectionView {
	rv := objc.Send[NSScrubberSelectionView](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s NSScrubberSelectionView) Autorelease() NSScrubberSelectionView {
	rv := objc.Send[NSScrubberSelectionView](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSScrubberSelectionView creates a new NSScrubberSelectionView instance.
func NewNSScrubberSelectionView() NSScrubberSelectionView {
	class := getNSScrubberSelectionViewClass()
	rv := objc.Send[NSScrubberSelectionView](objc.ID(class.class), objc.Sel("new"))
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
func NewScrubberSelectionViewWithCoder(coder foundation.INSCoder) NSScrubberSelectionView {
	instance := getNSScrubberSelectionViewClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSScrubberSelectionViewFromID(rv)
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
func NewScrubberSelectionViewWithFrame(frameRect corefoundation.CGRect) NSScrubberSelectionView {
	instance := getNSScrubberSelectionViewClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFrame:"), frameRect)
	return NSScrubberSelectionViewFromID(rv)
}






func (s NSScrubberSelectionView) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](s.ID, objc.Sel("encodeWithCoder:"), coder)
}












































