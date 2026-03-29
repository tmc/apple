// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
)

// The class instance for the [NSScrubberTextItemView] class.
var (
	_NSScrubberTextItemViewClass     NSScrubberTextItemViewClass
	_NSScrubberTextItemViewClassOnce sync.Once
)

func getNSScrubberTextItemViewClass() NSScrubberTextItemViewClass {
	_NSScrubberTextItemViewClassOnce.Do(func() {
		_NSScrubberTextItemViewClass = NSScrubberTextItemViewClass{class: objc.GetClass("NSScrubberTextItemView")}
	})
	return _NSScrubberTextItemViewClass
}

// GetNSScrubberTextItemViewClass returns the class object for NSScrubberTextItemView.
func GetNSScrubberTextItemViewClass() NSScrubberTextItemViewClass {
	return getNSScrubberTextItemViewClass()
}

type NSScrubberTextItemViewClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSScrubberTextItemViewClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSScrubberTextItemViewClass) Alloc() NSScrubberTextItemView {
	rv := objc.Send[NSScrubberTextItemView](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A concrete view subclass for displaying text for an item in a scrubber.
//
// # Overview
// 
// Provide the text you want to display in the scrubber item to the [NSScrubberTextItemView.Title]
// property. If you want finer control over the appearance of the text, you
// can access the underlying text field using the [TextField] property.
//
// # Providing text content
//
//   - [NSScrubberTextItemView.Title]: The text displayed for the scrubber item.
//   - [NSScrubberTextItemView.SetTitle]
//   - [NSScrubberTextItemView.TextField]: The text field that the scrubber item uses to display its text.
//
// See: https://developer.apple.com/documentation/AppKit/NSScrubberTextItemView
type NSScrubberTextItemView struct {
	NSScrubberItemView
}

// NSScrubberTextItemViewFromID constructs a [NSScrubberTextItemView] from an objc.ID.
//
// A concrete view subclass for displaying text for an item in a scrubber.
func NSScrubberTextItemViewFromID(id objc.ID) NSScrubberTextItemView {
	return NSScrubberTextItemView{NSScrubberItemView: NSScrubberItemViewFromID(id)}
}
// NOTE: NSScrubberTextItemView adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSScrubberTextItemView] class.
//
// # Providing text content
//
//   - [INSScrubberTextItemView.Title]: The text displayed for the scrubber item.
//   - [INSScrubberTextItemView.SetTitle]
//   - [INSScrubberTextItemView.TextField]: The text field that the scrubber item uses to display its text.
//
// See: https://developer.apple.com/documentation/AppKit/NSScrubberTextItemView
type INSScrubberTextItemView interface {
	INSScrubberItemView

	// Topic: Providing text content

	// The text displayed for the scrubber item.
	Title() string
	SetTitle(value string)
	// The text field that the scrubber item uses to display its text.
	TextField() INSTextField
}

// Init initializes the instance.
func (s NSScrubberTextItemView) Init() NSScrubberTextItemView {
	rv := objc.Send[NSScrubberTextItemView](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s NSScrubberTextItemView) Autorelease() NSScrubberTextItemView {
	rv := objc.Send[NSScrubberTextItemView](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSScrubberTextItemView creates a new NSScrubberTextItemView instance.
func NewNSScrubberTextItemView() NSScrubberTextItemView {
	class := getNSScrubberTextItemViewClass()
	rv := objc.Send[NSScrubberTextItemView](objc.ID(class.class), objc.Sel("new"))
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
func NewScrubberTextItemViewWithCoder(coder foundation.INSCoder) NSScrubberTextItemView {
	instance := getNSScrubberTextItemViewClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSScrubberTextItemViewFromID(rv)
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
func NewScrubberTextItemViewWithFrame(frameRect corefoundation.CGRect) NSScrubberTextItemView {
	instance := getNSScrubberTextItemViewClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFrame:"), frameRect)
	return NSScrubberTextItemViewFromID(rv)
}

// The text displayed for the scrubber item.
//
// See: https://developer.apple.com/documentation/AppKit/NSScrubberTextItemView/title
func (s NSScrubberTextItemView) Title() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("title"))
	return foundation.NSStringFromID(rv).String()
}
func (s NSScrubberTextItemView) SetTitle(value string) {
	objc.Send[struct{}](s.ID, objc.Sel("setTitle:"), objc.String(value))
}
// The text field that the scrubber item uses to display its text.
//
// # Discussion
// 
// Access and configure the underlying text field used to display the text in
// the [Title] property.
//
// See: https://developer.apple.com/documentation/AppKit/NSScrubberTextItemView/textField
func (s NSScrubberTextItemView) TextField() INSTextField {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("textField"))
	return NSTextFieldFromID(objc.ID(rv))
}

