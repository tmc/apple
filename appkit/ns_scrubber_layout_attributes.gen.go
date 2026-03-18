// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSScrubberLayoutAttributes] class.
var (
	_NSScrubberLayoutAttributesClass     NSScrubberLayoutAttributesClass
	_NSScrubberLayoutAttributesClassOnce sync.Once
)

func getNSScrubberLayoutAttributesClass() NSScrubberLayoutAttributesClass {
	_NSScrubberLayoutAttributesClassOnce.Do(func() {
		_NSScrubberLayoutAttributesClass = NSScrubberLayoutAttributesClass{class: objc.GetClass("NSScrubberLayoutAttributes")}
	})
	return _NSScrubberLayoutAttributesClass
}

// GetNSScrubberLayoutAttributesClass returns the class object for NSScrubberLayoutAttributes.
func GetNSScrubberLayoutAttributesClass() NSScrubberLayoutAttributesClass {
	return getNSScrubberLayoutAttributesClass()
}

type NSScrubberLayoutAttributesClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSScrubberLayoutAttributesClass) Alloc() NSScrubberLayoutAttributes {
	rv := objc.Send[NSScrubberLayoutAttributes](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// The layout of a scrubber item.
//
// # Overview
// 
// A layout attributes object is the model for the layout of a single item in
// a scrubber control.
// 
// If you require model attributes in addition to those provided by this
// class, create a subclass and add appropriate attributes. Subclasses must
// implement [isEqual(_:)], [NSScrubberLayoutAttributes.Hash] and the [NSCopying] protocol.
//
// [NSCopying]: https://developer.apple.com/documentation/Foundation/NSCopying
// [isEqual(_:)]: https://developer.apple.com/documentation/ObjectiveC/NSObjectProtocol/isEqual(_:)
//
// # Controlling the layout
//
//   - [NSScrubberLayoutAttributes.Alpha]: The item’s alpha value.
//   - [NSScrubberLayoutAttributes.SetAlpha]
//   - [NSScrubberLayoutAttributes.Frame]: The frame of the scrubber item.
//   - [NSScrubberLayoutAttributes.SetFrame]
//   - [NSScrubberLayoutAttributes.ItemIndex]: The index of the scrubber item that is represented by the item’s layout attributes.
//   - [NSScrubberLayoutAttributes.SetItemIndex]
//
// See: https://developer.apple.com/documentation/AppKit/NSScrubberLayoutAttributes
type NSScrubberLayoutAttributes struct {
	objectivec.Object
}

// NSScrubberLayoutAttributesFromID constructs a [NSScrubberLayoutAttributes] from an objc.ID.
//
// The layout of a scrubber item.
func NSScrubberLayoutAttributesFromID(id objc.ID) NSScrubberLayoutAttributes {
	return NSScrubberLayoutAttributes{objectivec.Object{ID: id}}
}
// NOTE: NSScrubberLayoutAttributes adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSScrubberLayoutAttributes] class.
//
// # Controlling the layout
//
//   - [INSScrubberLayoutAttributes.Alpha]: The item’s alpha value.
//   - [INSScrubberLayoutAttributes.SetAlpha]
//   - [INSScrubberLayoutAttributes.Frame]: The frame of the scrubber item.
//   - [INSScrubberLayoutAttributes.SetFrame]
//   - [INSScrubberLayoutAttributes.ItemIndex]: The index of the scrubber item that is represented by the item’s layout attributes.
//   - [INSScrubberLayoutAttributes.SetItemIndex]
//
// See: https://developer.apple.com/documentation/AppKit/NSScrubberLayoutAttributes
type INSScrubberLayoutAttributes interface {
	objectivec.IObject

	// Topic: Controlling the layout

	// The item’s alpha value.
	Alpha() float64
	SetAlpha(value float64)
	// The frame of the scrubber item.
	Frame() corefoundation.CGRect
	SetFrame(value corefoundation.CGRect)
	// The index of the scrubber item that is represented by the item’s layout attributes.
	ItemIndex() int
	SetItemIndex(value int)

	// Returns an integer that can be used as a table address in a hash table structure.
	Hash() int
	SetHash(value int)
}

// Init initializes the instance.
func (s NSScrubberLayoutAttributes) Init() NSScrubberLayoutAttributes {
	rv := objc.Send[NSScrubberLayoutAttributes](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s NSScrubberLayoutAttributes) Autorelease() NSScrubberLayoutAttributes {
	rv := objc.Send[NSScrubberLayoutAttributes](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSScrubberLayoutAttributes creates a new NSScrubberLayoutAttributes instance.
func NewNSScrubberLayoutAttributes() NSScrubberLayoutAttributes {
	class := getNSScrubberLayoutAttributesClass()
	rv := objc.Send[NSScrubberLayoutAttributes](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a new layout attributes object for the specified scrubber item
// index.
//
// index: The index of the scrubber item that this layout attributes object
// represents.
//
// # Return Value
// 
// A layout attributes object configured for the specified index.
//
// See: https://developer.apple.com/documentation/AppKit/NSScrubberLayoutAttributes/init(forItemAt:)
func NewScrubberLayoutAttributesForItemAtIndex(index int) NSScrubberLayoutAttributes {
	rv := objc.Send[objc.ID](objc.ID(getNSScrubberLayoutAttributesClass().class), objc.Sel("layoutAttributesForItemAtIndex:"), index)
	return NSScrubberLayoutAttributesFromID(rv)
}

// The item’s alpha value.
//
// See: https://developer.apple.com/documentation/AppKit/NSScrubberLayoutAttributes/alpha
func (s NSScrubberLayoutAttributes) Alpha() float64 {
	rv := objc.Send[float64](s.ID, objc.Sel("alpha"))
	return rv
}
func (s NSScrubberLayoutAttributes) SetAlpha(value float64) {
	objc.Send[struct{}](s.ID, objc.Sel("setAlpha:"), value)
}

// The frame of the scrubber item.
//
// See: https://developer.apple.com/documentation/AppKit/NSScrubberLayoutAttributes/frame
func (s NSScrubberLayoutAttributes) Frame() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](s.ID, objc.Sel("frame"))
	return corefoundation.CGRect(rv)
}
func (s NSScrubberLayoutAttributes) SetFrame(value corefoundation.CGRect) {
	objc.Send[struct{}](s.ID, objc.Sel("setFrame:"), value)
}

// The index of the scrubber item that is represented by the item’s layout
// attributes.
//
// See: https://developer.apple.com/documentation/AppKit/NSScrubberLayoutAttributes/itemIndex
func (s NSScrubberLayoutAttributes) ItemIndex() int {
	rv := objc.Send[int](s.ID, objc.Sel("itemIndex"))
	return rv
}
func (s NSScrubberLayoutAttributes) SetItemIndex(value int) {
	objc.Send[struct{}](s.ID, objc.Sel("setItemIndex:"), value)
}

// Returns an integer that can be used as a table address in a hash table
// structure.
//
// See: https://developer.apple.com/documentation/ObjectiveC/NSObjectProtocol/hash
func (s NSScrubberLayoutAttributes) Hash() int {
	rv := objc.Send[int](s.ID, objc.Sel("hash"))
	return rv
}
func (s NSScrubberLayoutAttributes) SetHash(value int) {
	objc.Send[struct{}](s.ID, objc.Sel("setHash:"), value)
}

