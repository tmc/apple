// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSTextViewportLayoutController] class.
var (
	_NSTextViewportLayoutControllerClass     NSTextViewportLayoutControllerClass
	_NSTextViewportLayoutControllerClassOnce sync.Once
)

func getNSTextViewportLayoutControllerClass() NSTextViewportLayoutControllerClass {
	_NSTextViewportLayoutControllerClassOnce.Do(func() {
		_NSTextViewportLayoutControllerClass = NSTextViewportLayoutControllerClass{class: objc.GetClass("NSTextViewportLayoutController")}
	})
	return _NSTextViewportLayoutControllerClass
}

// GetNSTextViewportLayoutControllerClass returns the class object for NSTextViewportLayoutController.
func GetNSTextViewportLayoutControllerClass() NSTextViewportLayoutControllerClass {
	return getNSTextViewportLayoutControllerClass()
}

type NSTextViewportLayoutControllerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSTextViewportLayoutControllerClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSTextViewportLayoutControllerClass) Alloc() NSTextViewportLayoutController {
	rv := objc.Send[NSTextViewportLayoutController](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// Manages the layout process inside the viewport interacting with its
// delegate.
//
// # Overview
// 
// A viewport is a rectangular area within a flipped coordinate system
// expanding along the y-axis. With text contents, lines advance expanding the
// view in the current writing direction. The viewport defines the active area
// where the framework lays out text fragments. In most cases, the area
// corresponds to the user visible area with an additional over-scroll region.
//
// # Creating a viewport layout controller
//
//   - [NSTextViewportLayoutController.InitWithTextLayoutManager]: Creates a new instance with the text layout manager you provide.
//
// # Accessing the layout manager
//
//   - [NSTextViewportLayoutController.TextLayoutManager]: Returns the text layout manager for this viewport layout controller.
//
// # Responding to changes in viewport layout
//
//   - [NSTextViewportLayoutController.Delegate]: The delegate for the text layout manager object.
//   - [NSTextViewportLayoutController.SetDelegate]
//
// # Accessing the viewport characteristics
//
//   - [NSTextViewportLayoutController.ViewportBounds]: Returns the visible bounds of the view, plus the overdraw area.
//   - [NSTextViewportLayoutController.ViewportRange]: Returns the text range of the current viewport layout.
//   - [NSTextViewportLayoutController.AdjustViewportByVerticalOffset]: Adjusts the viewport rect by the specified offset if needed.
//   - [NSTextViewportLayoutController.LayoutViewport]: Performs layout in the viewport.
//   - [NSTextViewportLayoutController.RelocateViewportToTextLocation]: Relocates the viewport to the location you specify.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextViewportLayoutController
type NSTextViewportLayoutController struct {
	objectivec.Object
}

// NSTextViewportLayoutControllerFromID constructs a [NSTextViewportLayoutController] from an objc.ID.
//
// Manages the layout process inside the viewport interacting with its
// delegate.
func NSTextViewportLayoutControllerFromID(id objc.ID) NSTextViewportLayoutController {
	return NSTextViewportLayoutController{objectivec.Object{ID: id}}
}
// NOTE: NSTextViewportLayoutController adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSTextViewportLayoutController] class.
//
// # Creating a viewport layout controller
//
//   - [INSTextViewportLayoutController.InitWithTextLayoutManager]: Creates a new instance with the text layout manager you provide.
//
// # Accessing the layout manager
//
//   - [INSTextViewportLayoutController.TextLayoutManager]: Returns the text layout manager for this viewport layout controller.
//
// # Responding to changes in viewport layout
//
//   - [INSTextViewportLayoutController.Delegate]: The delegate for the text layout manager object.
//   - [INSTextViewportLayoutController.SetDelegate]
//
// # Accessing the viewport characteristics
//
//   - [INSTextViewportLayoutController.ViewportBounds]: Returns the visible bounds of the view, plus the overdraw area.
//   - [INSTextViewportLayoutController.ViewportRange]: Returns the text range of the current viewport layout.
//   - [INSTextViewportLayoutController.AdjustViewportByVerticalOffset]: Adjusts the viewport rect by the specified offset if needed.
//   - [INSTextViewportLayoutController.LayoutViewport]: Performs layout in the viewport.
//   - [INSTextViewportLayoutController.RelocateViewportToTextLocation]: Relocates the viewport to the location you specify.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextViewportLayoutController
type INSTextViewportLayoutController interface {
	objectivec.IObject

	// Topic: Creating a viewport layout controller

	// Creates a new instance with the text layout manager you provide.
	InitWithTextLayoutManager(textLayoutManager INSTextLayoutManager) NSTextViewportLayoutController

	// Topic: Accessing the layout manager

	// Returns the text layout manager for this viewport layout controller.
	TextLayoutManager() INSTextLayoutManager

	// Topic: Responding to changes in viewport layout

	// The delegate for the text layout manager object.
	Delegate() NSTextViewportLayoutControllerDelegate
	SetDelegate(value NSTextViewportLayoutControllerDelegate)

	// Topic: Accessing the viewport characteristics

	// Returns the visible bounds of the view, plus the overdraw area.
	ViewportBounds() corefoundation.CGRect
	// Returns the text range of the current viewport layout.
	ViewportRange() INSTextRange
	// Adjusts the viewport rect by the specified offset if needed.
	AdjustViewportByVerticalOffset(verticalOffset float64)
	// Performs layout in the viewport.
	LayoutViewport()
	// Relocates the viewport to the location you specify.
	RelocateViewportToTextLocation(textLocation NSTextLocation) float64
}

// Init initializes the instance.
func (t NSTextViewportLayoutController) Init() NSTextViewportLayoutController {
	rv := objc.Send[NSTextViewportLayoutController](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t NSTextViewportLayoutController) Autorelease() NSTextViewportLayoutController {
	rv := objc.Send[NSTextViewportLayoutController](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSTextViewportLayoutController creates a new NSTextViewportLayoutController instance.
func NewNSTextViewportLayoutController() NSTextViewportLayoutController {
	class := getNSTextViewportLayoutControllerClass()
	rv := objc.Send[NSTextViewportLayoutController](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a new instance with the text layout manager you provide.
//
// textLayoutManager: The [NSTextLayoutManager] to associate with this viewport layout
// controller.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextViewportLayoutController/init(textLayoutManager:)
func NewTextViewportLayoutControllerWithTextLayoutManager(textLayoutManager INSTextLayoutManager) NSTextViewportLayoutController {
	instance := getNSTextViewportLayoutControllerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTextLayoutManager:"), textLayoutManager)
	return NSTextViewportLayoutControllerFromID(rv)
}

// Creates a new instance with the text layout manager you provide.
//
// textLayoutManager: The [NSTextLayoutManager] to associate with this viewport layout
// controller.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextViewportLayoutController/init(textLayoutManager:)
func (t NSTextViewportLayoutController) InitWithTextLayoutManager(textLayoutManager INSTextLayoutManager) NSTextViewportLayoutController {
	rv := objc.Send[NSTextViewportLayoutController](t.ID, objc.Sel("initWithTextLayoutManager:"), textLayoutManager)
	return rv
}
// Adjusts the viewport rect by the specified offset if needed.
//
// verticalOffset: A [CGFloat] that represents the offset amount to apply to the viewport.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextViewportLayoutController/adjustViewport(byVerticalOffset:)
func (t NSTextViewportLayoutController) AdjustViewportByVerticalOffset(verticalOffset float64) {
	objc.Send[objc.ID](t.ID, objc.Sel("adjustViewportByVerticalOffset:"), verticalOffset)
}
// Performs layout in the viewport.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextViewportLayoutController/layoutViewport()
func (t NSTextViewportLayoutController) LayoutViewport() {
	objc.Send[objc.ID](t.ID, objc.Sel("layoutViewport"))
}
// Relocates the viewport to the location you specify.
//
// textLocation: An [NSTextLocation].
//
// See: https://developer.apple.com/documentation/AppKit/NSTextViewportLayoutController/relocateViewport(to:)
func (t NSTextViewportLayoutController) RelocateViewportToTextLocation(textLocation NSTextLocation) float64 {
	rv := objc.Send[float64](t.ID, objc.Sel("relocateViewportToTextLocation:"), textLocation)
	return rv
}

// Returns the text layout manager for this viewport layout controller.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextViewportLayoutController/textLayoutManager
func (t NSTextViewportLayoutController) TextLayoutManager() INSTextLayoutManager {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("textLayoutManager"))
	return NSTextLayoutManagerFromID(objc.ID(rv))
}
// The delegate for the text layout manager object.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextViewportLayoutController/delegate
func (t NSTextViewportLayoutController) Delegate() NSTextViewportLayoutControllerDelegate {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("delegate"))
	return NSTextViewportLayoutControllerDelegateObjectFromID(rv)
}
func (t NSTextViewportLayoutController) SetDelegate(value NSTextViewportLayoutControllerDelegate) {
	objc.Send[struct{}](t.ID, objc.Sel("setDelegate:"), value)
}
// Returns the visible bounds of the view, plus the overdraw area.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextViewportLayoutController/viewportBounds
func (t NSTextViewportLayoutController) ViewportBounds() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](t.ID, objc.Sel("viewportBounds"))
	return corefoundation.CGRect(rv)
}
// Returns the text range of the current viewport layout.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextViewportLayoutController/viewportRange
func (t NSTextViewportLayoutController) ViewportRange() INSTextRange {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("viewportRange"))
	return NSTextRangeFromID(objc.ID(rv))
}

