// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
)

// The class instance for the [NSScrubberFlowLayout] class.
var (
	_NSScrubberFlowLayoutClass     NSScrubberFlowLayoutClass
	_NSScrubberFlowLayoutClassOnce sync.Once
)

func getNSScrubberFlowLayoutClass() NSScrubberFlowLayoutClass {
	_NSScrubberFlowLayoutClassOnce.Do(func() {
		_NSScrubberFlowLayoutClass = NSScrubberFlowLayoutClass{class: objc.GetClass("NSScrubberFlowLayout")}
	})
	return _NSScrubberFlowLayoutClass
}

// GetNSScrubberFlowLayoutClass returns the class object for NSScrubberFlowLayout.
func GetNSScrubberFlowLayoutClass() NSScrubberFlowLayoutClass {
	return getNSScrubberFlowLayoutClass()
}

type NSScrubberFlowLayoutClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSScrubberFlowLayoutClass) Alloc() NSScrubberFlowLayout {
	rv := objc.Send[NSScrubberFlowLayout](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A concrete layout object that arranges items end-to-end in a linear strip.
//
// # Overview
// 
// To set the size of items on a per-item basis, ensure that your scrubber
// delegate conforms to the [NSScrubberFlowLayoutDelegate] protocol, and
// provides an implementation of the [ScrubberLayoutSizeForItemAtIndex]
// method.
//
// # Configuring the layout
//
//   - [NSScrubberFlowLayout.ItemSpacing]: The horizontal spacing between items, specified in points.
//   - [NSScrubberFlowLayout.SetItemSpacing]
//   - [NSScrubberFlowLayout.ItemSize]: The frame size for each item in the scrubber.
//   - [NSScrubberFlowLayout.SetItemSize]
//
// # Invalidating the layout
//
//   - [NSScrubberFlowLayout.InvalidateLayoutForItemsAtIndexes]: Informs the scrubber that it should perform a new layout pass for the items at the specified indexes.
//
// See: https://developer.apple.com/documentation/AppKit/NSScrubberFlowLayout
type NSScrubberFlowLayout struct {
	NSScrubberLayout
}

// NSScrubberFlowLayoutFromID constructs a [NSScrubberFlowLayout] from an objc.ID.
//
// A concrete layout object that arranges items end-to-end in a linear strip.
func NSScrubberFlowLayoutFromID(id objc.ID) NSScrubberFlowLayout {
	return NSScrubberFlowLayout{NSScrubberLayout: NSScrubberLayoutFromID(id)}
}
// NOTE: NSScrubberFlowLayout adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSScrubberFlowLayout] class.
//
// # Configuring the layout
//
//   - [INSScrubberFlowLayout.ItemSpacing]: The horizontal spacing between items, specified in points.
//   - [INSScrubberFlowLayout.SetItemSpacing]
//   - [INSScrubberFlowLayout.ItemSize]: The frame size for each item in the scrubber.
//   - [INSScrubberFlowLayout.SetItemSize]
//
// # Invalidating the layout
//
//   - [INSScrubberFlowLayout.InvalidateLayoutForItemsAtIndexes]: Informs the scrubber that it should perform a new layout pass for the items at the specified indexes.
//
// See: https://developer.apple.com/documentation/AppKit/NSScrubberFlowLayout
type INSScrubberFlowLayout interface {
	INSScrubberLayout

	// Topic: Configuring the layout

	// The horizontal spacing between items, specified in points.
	ItemSpacing() float64
	SetItemSpacing(value float64)
	// The frame size for each item in the scrubber.
	ItemSize() corefoundation.CGSize
	SetItemSize(value corefoundation.CGSize)

	// Topic: Invalidating the layout

	// Informs the scrubber that it should perform a new layout pass for the items at the specified indexes.
	InvalidateLayoutForItemsAtIndexes(invalidItemIndexes foundation.NSIndexSet)
}

// Init initializes the instance.
func (s NSScrubberFlowLayout) Init() NSScrubberFlowLayout {
	rv := objc.Send[NSScrubberFlowLayout](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s NSScrubberFlowLayout) Autorelease() NSScrubberFlowLayout {
	rv := objc.Send[NSScrubberFlowLayout](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSScrubberFlowLayout creates a new NSScrubberFlowLayout instance.
func NewNSScrubberFlowLayout() NSScrubberFlowLayout {
	class := getNSScrubberFlowLayoutClass()
	rv := objc.Send[NSScrubberFlowLayout](objc.ID(class.class), objc.Sel("new"))
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
func NewScrubberFlowLayoutWithCoder(coder foundation.INSCoder) NSScrubberFlowLayout {
	instance := getNSScrubberFlowLayoutClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSScrubberFlowLayoutFromID(rv)
}

// Informs the scrubber that it should perform a new layout pass for the items
// at the specified indexes.
//
// invalidItemIndexes: An index set containing the indexes of the items whose layout should be
// invalidated.
//
// See: https://developer.apple.com/documentation/AppKit/NSScrubberFlowLayout/invalidateLayoutForItems(at:)
func (s NSScrubberFlowLayout) InvalidateLayoutForItemsAtIndexes(invalidItemIndexes foundation.NSIndexSet) {
	objc.Send[objc.ID](s.ID, objc.Sel("invalidateLayoutForItemsAtIndexes:"), invalidItemIndexes)
}

// The horizontal spacing between items, specified in points.
//
// # Discussion
// 
// The default value of this property is `0.0`.
//
// See: https://developer.apple.com/documentation/AppKit/NSScrubberFlowLayout/itemSpacing
func (s NSScrubberFlowLayout) ItemSpacing() float64 {
	rv := objc.Send[float64](s.ID, objc.Sel("itemSpacing"))
	return rv
}
func (s NSScrubberFlowLayout) SetItemSpacing(value float64) {
	objc.Send[struct{}](s.ID, objc.Sel("setItemSpacing:"), value)
}

// The frame size for each item in the scrubber.
//
// # Discussion
// 
// You can override the value of this property on a per-item basis, by
// providing a delegate object to the scrubber that conforms to the
// [NSScrubberFlowLayoutDelegate] protocol. This delegate object must
// implement the [ScrubberLayoutSizeForItemAtIndex] method.
// 
// The default value for this property has a width of `50.0` points and a
// height of `30.0` points.
//
// See: https://developer.apple.com/documentation/AppKit/NSScrubberFlowLayout/itemSize
func (s NSScrubberFlowLayout) ItemSize() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](s.ID, objc.Sel("itemSize"))
	return corefoundation.CGSize(rv)
}
func (s NSScrubberFlowLayout) SetItemSize(value corefoundation.CGSize) {
	objc.Send[struct{}](s.ID, objc.Sel("setItemSize:"), value)
}

