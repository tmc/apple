// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSScrubberLayout] class.
var (
	_NSScrubberLayoutClass     NSScrubberLayoutClass
	_NSScrubberLayoutClassOnce sync.Once
)

func getNSScrubberLayoutClass() NSScrubberLayoutClass {
	_NSScrubberLayoutClassOnce.Do(func() {
		_NSScrubberLayoutClass = NSScrubberLayoutClass{class: objc.GetClass("NSScrubberLayout")}
	})
	return _NSScrubberLayoutClass
}

// GetNSScrubberLayoutClass returns the class object for NSScrubberLayout.
func GetNSScrubberLayoutClass() NSScrubberLayoutClass {
	return getNSScrubberLayoutClass()
}

type NSScrubberLayoutClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSScrubberLayoutClass) Alloc() NSScrubberLayout {
	rv := objc.Send[NSScrubberLayout](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// An abstract class that describes the layout of items within a scrubber
// control.
//
// # Overview
// 
// To determine the layout of items in a scrubber, use one of the built-in
// subclasses ([NSScrubberProportionalLayout] or [NSScrubberFlowLayout]), or
// create a custom subclass to implement your own layout.
//
// # Creating a scrubber layout
//
//   - [NSScrubberLayout.InitWithCoder]: Initializes and returns a newly allocated scrubber layout object from a storyboard or nib file.
//
// # Configuring a scrubber layout
//
//   - [NSScrubberLayout.Scrubber]: The scrubber control that this layout is assigned to.
//   - [NSScrubberLayout.VisibleRect]: The currently visible rectangle, in the coordinate space of the scrubber content.
//   - [NSScrubberLayout.InvalidateLayout]: Signals that the layout has been invalidated, and that the scrubber control should perform a new layout pass.
//
// # Subclassing a scrubber layout
//
//   - [NSScrubberLayout.PrepareLayout]: Gives you an opportunity to perform layout calculations when the scrubber’s layout is invalidated.
//   - [NSScrubberLayout.ScrubberContentSize]: The size required to contain all elements within the scrubber.
//   - [NSScrubberLayout.LayoutAttributesForItemAtIndex]: The layout attributes for the item with the specified index.
//   - [NSScrubberLayout.LayoutAttributesForItemsInRect]: The set of layout attributes for all items within the provided rectangle.
//   - [NSScrubberLayout.ShouldInvalidateLayoutForSelectionChange]: Determines whether the scrubber should refresh its layout when the selection changes.
//   - [NSScrubberLayout.ShouldInvalidateLayoutForHighlightChange]: Determines whether the scrubber should refresh its layout when an item is highlighted.
//   - [NSScrubberLayout.ShouldInvalidateLayoutForChangeFromVisibleRectToVisibleRect]: Determines whether the scrubber should refresh its layout in response to a change of its visible region.
//   - [NSScrubberLayout.AutomaticallyMirrorsInRightToLeftLayout]: Determines whether the scrubber mirrors its layout for right-to-left layouts.
//
// See: https://developer.apple.com/documentation/AppKit/NSScrubberLayout
type NSScrubberLayout struct {
	objectivec.Object
}

// NSScrubberLayoutFromID constructs a [NSScrubberLayout] from an objc.ID.
//
// An abstract class that describes the layout of items within a scrubber
// control.
func NSScrubberLayoutFromID(id objc.ID) NSScrubberLayout {
	return NSScrubberLayout{objectivec.Object{ID: id}}
}
// NOTE: NSScrubberLayout adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSScrubberLayout] class.
//
// # Creating a scrubber layout
//
//   - [INSScrubberLayout.InitWithCoder]: Initializes and returns a newly allocated scrubber layout object from a storyboard or nib file.
//
// # Configuring a scrubber layout
//
//   - [INSScrubberLayout.Scrubber]: The scrubber control that this layout is assigned to.
//   - [INSScrubberLayout.VisibleRect]: The currently visible rectangle, in the coordinate space of the scrubber content.
//   - [INSScrubberLayout.InvalidateLayout]: Signals that the layout has been invalidated, and that the scrubber control should perform a new layout pass.
//
// # Subclassing a scrubber layout
//
//   - [INSScrubberLayout.PrepareLayout]: Gives you an opportunity to perform layout calculations when the scrubber’s layout is invalidated.
//   - [INSScrubberLayout.ScrubberContentSize]: The size required to contain all elements within the scrubber.
//   - [INSScrubberLayout.LayoutAttributesForItemAtIndex]: The layout attributes for the item with the specified index.
//   - [INSScrubberLayout.LayoutAttributesForItemsInRect]: The set of layout attributes for all items within the provided rectangle.
//   - [INSScrubberLayout.ShouldInvalidateLayoutForSelectionChange]: Determines whether the scrubber should refresh its layout when the selection changes.
//   - [INSScrubberLayout.ShouldInvalidateLayoutForHighlightChange]: Determines whether the scrubber should refresh its layout when an item is highlighted.
//   - [INSScrubberLayout.ShouldInvalidateLayoutForChangeFromVisibleRectToVisibleRect]: Determines whether the scrubber should refresh its layout in response to a change of its visible region.
//   - [INSScrubberLayout.AutomaticallyMirrorsInRightToLeftLayout]: Determines whether the scrubber mirrors its layout for right-to-left layouts.
//
// See: https://developer.apple.com/documentation/AppKit/NSScrubberLayout
type INSScrubberLayout interface {
	objectivec.IObject

	// Topic: Creating a scrubber layout

	// Initializes and returns a newly allocated scrubber layout object from a storyboard or nib file.
	InitWithCoder(coder foundation.INSCoder) NSScrubberLayout

	// Topic: Configuring a scrubber layout

	// The scrubber control that this layout is assigned to.
	Scrubber() INSScrubber
	// The currently visible rectangle, in the coordinate space of the scrubber content.
	VisibleRect() corefoundation.CGRect
	// Signals that the layout has been invalidated, and that the scrubber control should perform a new layout pass.
	InvalidateLayout()

	// Topic: Subclassing a scrubber layout

	// Gives you an opportunity to perform layout calculations when the scrubber’s layout is invalidated.
	PrepareLayout()
	// The size required to contain all elements within the scrubber.
	ScrubberContentSize() corefoundation.CGSize
	// The layout attributes for the item with the specified index.
	LayoutAttributesForItemAtIndex(index int) INSScrubberLayoutAttributes
	// The set of layout attributes for all items within the provided rectangle.
	LayoutAttributesForItemsInRect(rect corefoundation.CGRect) foundation.INSSet
	// Determines whether the scrubber should refresh its layout when the selection changes.
	ShouldInvalidateLayoutForSelectionChange() bool
	// Determines whether the scrubber should refresh its layout when an item is highlighted.
	ShouldInvalidateLayoutForHighlightChange() bool
	// Determines whether the scrubber should refresh its layout in response to a change of its visible region.
	ShouldInvalidateLayoutForChangeFromVisibleRectToVisibleRect(fromVisibleRect corefoundation.CGRect, toVisibleRect corefoundation.CGRect) bool
	// Determines whether the scrubber mirrors its layout for right-to-left layouts.
	AutomaticallyMirrorsInRightToLeftLayout() bool

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (s NSScrubberLayout) Init() NSScrubberLayout {
	rv := objc.Send[NSScrubberLayout](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s NSScrubberLayout) Autorelease() NSScrubberLayout {
	rv := objc.Send[NSScrubberLayout](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSScrubberLayout creates a new NSScrubberLayout instance.
func NewNSScrubberLayout() NSScrubberLayout {
	class := getNSScrubberLayoutClass()
	rv := objc.Send[NSScrubberLayout](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Initializes and returns a newly allocated scrubber layout object from a
// storyboard or nib file.
//
// # Return Value
// 
// A properly initialized scrubber layout object.
//
// See: https://developer.apple.com/documentation/AppKit/NSScrubberLayout/init(coder:)
func NewScrubberLayoutWithCoder(coder foundation.INSCoder) NSScrubberLayout {
	instance := getNSScrubberLayoutClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSScrubberLayoutFromID(rv)
}

// Initializes and returns a newly allocated scrubber layout object from a
// storyboard or nib file.
//
// # Return Value
// 
// A properly initialized scrubber layout object.
//
// See: https://developer.apple.com/documentation/AppKit/NSScrubberLayout/init(coder:)
func (s NSScrubberLayout) InitWithCoder(coder foundation.INSCoder) NSScrubberLayout {
	rv := objc.Send[NSScrubberLayout](s.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}
// Signals that the layout has been invalidated, and that the scrubber control
// should perform a new layout pass.
//
// See: https://developer.apple.com/documentation/AppKit/NSScrubberLayout/invalidateLayout()
func (s NSScrubberLayout) InvalidateLayout() {
	objc.Send[objc.ID](s.ID, objc.Sel("invalidateLayout"))
}
// Gives you an opportunity to perform layout calculations when the
// scrubber’s layout is invalidated.
//
// # Discussion
// 
// Use this method in subclasses to perform layout calculations and caching in
// advance of rendering the scrubber.
// 
// The system calls this method when the scrubber’s layout is invalidated.
// 
// The base implementation of this method does nothing.
//
// See: https://developer.apple.com/documentation/AppKit/NSScrubberLayout/prepare()
func (s NSScrubberLayout) PrepareLayout() {
	objc.Send[objc.ID](s.ID, objc.Sel("prepareLayout"))
}
// The layout attributes for the item with the specified index.
//
// # Discussion
// 
// The base implementation returns `nil`.
//
// See: https://developer.apple.com/documentation/AppKit/NSScrubberLayout/layoutAttributesForItem(at:)
func (s NSScrubberLayout) LayoutAttributesForItemAtIndex(index int) INSScrubberLayoutAttributes {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("layoutAttributesForItemAtIndex:"), index)
	return NSScrubberLayoutAttributesFromID(rv)
}
// The set of layout attributes for all items within the provided rectangle.
//
// # Discussion
// 
// The base implementation returns an empty [NSSet].
//
// [NSSet]: https://developer.apple.com/documentation/Foundation/NSSet
//
// See: https://developer.apple.com/documentation/AppKit/NSScrubberLayout/layoutAttributesForItems(in:)
func (s NSScrubberLayout) LayoutAttributesForItemsInRect(rect corefoundation.CGRect) foundation.INSSet {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("layoutAttributesForItemsInRect:"), rect)
	return foundation.NSSetFromID(rv)
}
// Determines whether the scrubber should refresh its layout in response to a
// change of its visible region.
//
// # Discussion
// 
// If [true], the scrubber invalidates its layout in response to a change of
// the visible region, in response to a user scroll action, or to the scrubber
// being resized. Subclasses that rely on the size or origin of the visible
// region should return [true].
// 
// The base implementation of this method returns [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSScrubberLayout/shouldInvalidateLayoutForChange(fromVisibleRect:toVisibleRect:)
func (s NSScrubberLayout) ShouldInvalidateLayoutForChangeFromVisibleRectToVisibleRect(fromVisibleRect corefoundation.CGRect, toVisibleRect corefoundation.CGRect) bool {
	rv := objc.Send[bool](s.ID, objc.Sel("shouldInvalidateLayoutForChangeFromVisibleRect:toVisibleRect:"), fromVisibleRect, toVisibleRect)
	return rv
}
func (s NSScrubberLayout) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](s.ID, objc.Sel("encodeWithCoder:"), coder)
}

// The scrubber control that this layout is assigned to.
//
// # Discussion
// 
// If this layout is not currently assigned to a scrubber control, the value
// of this property is `nil`.
//
// See: https://developer.apple.com/documentation/AppKit/NSScrubberLayout/scrubber
func (s NSScrubberLayout) Scrubber() INSScrubber {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("scrubber"))
	return NSScrubberFromID(objc.ID(rv))
}
// The currently visible rectangle, in the coordinate space of the scrubber
// content.
//
// # Discussion
// 
// The value of this property is [NSZeroRect] if the layout is not currently
// assigned to a scrubber control.
//
// [NSZeroRect]: https://developer.apple.com/documentation/Foundation/NSZeroRect
//
// See: https://developer.apple.com/documentation/AppKit/NSScrubberLayout/visibleRect
func (s NSScrubberLayout) VisibleRect() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](s.ID, objc.Sel("visibleRect"))
	return corefoundation.CGRect(rv)
}
// The size required to contain all elements within the scrubber.
//
// # Discussion
// 
// The base implementation returns [NSZeroSize].
//
// See: https://developer.apple.com/documentation/AppKit/NSScrubberLayout/scrubberContentSize
func (s NSScrubberLayout) ScrubberContentSize() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](s.ID, objc.Sel("scrubberContentSize"))
	return corefoundation.CGSize(rv)
}
// Determines whether the scrubber should refresh its layout when the
// selection changes.
//
// # Discussion
// 
// If [true], the scrubber invalidates its layout when the selection changes.
// Subclasses should return [true] if the selection index affects the item
// layout.
// 
// The base implementation returns [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSScrubberLayout/shouldInvalidateLayoutForSelectionChange
func (s NSScrubberLayout) ShouldInvalidateLayoutForSelectionChange() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("shouldInvalidateLayoutForSelectionChange"))
	return rv
}
// Determines whether the scrubber should refresh its layout when an item is
// highlighted.
//
// # Discussion
// 
// If [true], the scrubber invalidates its layout when an item is highlighted.
// Subclasses should return [true] if the highlight state affects the item
// layout.
// 
// The base implementation of this method returns [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSScrubberLayout/shouldInvalidateLayoutForHighlightChange
func (s NSScrubberLayout) ShouldInvalidateLayoutForHighlightChange() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("shouldInvalidateLayoutForHighlightChange"))
	return rv
}
// Determines whether the scrubber mirrors its layout for right-to-left
// layouts.
//
// # Discussion
// 
// If [true], the layout of a scrubber in a right-to-left interface is the
// mirror of the left-to-right version. If you wish to customize the behavior
// of the scrubber in right-to-left interfaces in a custom subclass, override
// this method to return [false].
// 
// The base implementation of this method returns [true].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSScrubberLayout/automaticallyMirrorsInRightToLeftLayout
func (s NSScrubberLayout) AutomaticallyMirrorsInRightToLeftLayout() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("automaticallyMirrorsInRightToLeftLayout"))
	return rv
}

// A property containing a class that describes layout attributes.
//
// # Return Value
// 
// The default return value is [NSScrubberLayoutAttributes].
// 
// # Discussion
// 
// Create a custom subclass of [NSScrubberLayout] and override this method if
// you wish to use a custom layout attributes subclass.
//
// See: https://developer.apple.com/documentation/AppKit/NSScrubberLayout/layoutAttributesClass
func (_NSScrubberLayoutClass NSScrubberLayoutClass) LayoutAttributesClass() objc.Class {
	rv := objc.Send[objc.Class](objc.ID(_NSScrubberLayoutClass.class), objc.Sel("layoutAttributesClass"))
	return rv
}

