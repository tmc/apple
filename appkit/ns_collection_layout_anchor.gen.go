// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSCollectionLayoutAnchor] class.
var (
	_NSCollectionLayoutAnchorClass     NSCollectionLayoutAnchorClass
	_NSCollectionLayoutAnchorClassOnce sync.Once
)

func getNSCollectionLayoutAnchorClass() NSCollectionLayoutAnchorClass {
	_NSCollectionLayoutAnchorClassOnce.Do(func() {
		_NSCollectionLayoutAnchorClass = NSCollectionLayoutAnchorClass{class: objc.GetClass("NSCollectionLayoutAnchor")}
	})
	return _NSCollectionLayoutAnchorClass
}

// GetNSCollectionLayoutAnchorClass returns the class object for NSCollectionLayoutAnchor.
func GetNSCollectionLayoutAnchorClass() NSCollectionLayoutAnchorClass {
	return getNSCollectionLayoutAnchorClass()
}

type NSCollectionLayoutAnchorClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSCollectionLayoutAnchorClass) Alloc() NSCollectionLayoutAnchor {
	rv := objc.Send[NSCollectionLayoutAnchor](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// An object that defines how to attach a supplementary item to an item in a
// collection view.
//
// # Overview
// 
// You use an anchor to attach a supplementary item to a specific item. An
// anchor contains information about where on the item your supplementary item
// is attached, including:
// 
// - An edge or set of edges. You can attach a supplementary item to a single
// edge, or to a corner by specifying two adjacent edges. - An offset from the
// item. By default, the supplementary item is anchored within the specified
// edges of the item it’s attached to. You can change this location by
// providing a custom offset when you create an anchor.
// 
// # Edges
// 
// The leading and trailing edges for anchors differ in left-to-right versus
// right-to-left environments. In a left-to-right environment, the leading
// edge is on the left, and the trailing edge is on the right. In a
// right-to-left environment, the leading edge is on the right, and the
// trailing edge is on the left. This difference ensures that your collection
// view layout is built with support for right-to-left languages.
// 
// The following diagram shows anchor placement for the specified edges in a
// left-to-right environment.
// 
// [media-3570665]
// 
// # Offset
// 
// You can express anchor offset in these ways:
// 
// - Absolute value. The offset is calculated as a point value. For example,
// an absolute x offset of `30.0` means that the origin of the supplementary
// item is offset by 30 points in the positive x direction. - Fractional
// value. The offset is calculated as a fraction of the supplementary item’s
// dimensions. For example, a fractional x offset of `0.3` means that the
// origin of the supplementary item is offset by 30% of the supplementary
// item’s width in the positive x direction.
// 
// The following code creates a basic badge and attaches it to an item’s top
// trailing corner.
//
// # Getting the edges
//
//   - [NSCollectionLayoutAnchor.Edges]: The edges of the item an anchor is attached to.
//
// # Getting the offset
//
//   - [NSCollectionLayoutAnchor.Offset]: The floating-point value of the anchor’s offset from the item it’s attached to.
//   - [NSCollectionLayoutAnchor.IsAbsoluteOffset]: A Boolean value that indicates whether the anchor’s offset is expressed as an absolute value.
//   - [NSCollectionLayoutAnchor.IsFractionalOffset]: A Boolean value that indicates whether the anchor’s offset is expressed as a fraction of its supplementary item’s dimension.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutAnchor
type NSCollectionLayoutAnchor struct {
	objectivec.Object
}

// NSCollectionLayoutAnchorFromID constructs a [NSCollectionLayoutAnchor] from an objc.ID.
//
// An object that defines how to attach a supplementary item to an item in a
// collection view.
func NSCollectionLayoutAnchorFromID(id objc.ID) NSCollectionLayoutAnchor {
	return NSCollectionLayoutAnchor{objectivec.Object{id}}
}
// NOTE: NSCollectionLayoutAnchor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [NSCollectionLayoutAnchor] class.
//
// # Getting the edges
//
//   - [INSCollectionLayoutAnchor.Edges]: The edges of the item an anchor is attached to.
//
// # Getting the offset
//
//   - [INSCollectionLayoutAnchor.Offset]: The floating-point value of the anchor’s offset from the item it’s attached to.
//   - [INSCollectionLayoutAnchor.IsAbsoluteOffset]: A Boolean value that indicates whether the anchor’s offset is expressed as an absolute value.
//   - [INSCollectionLayoutAnchor.IsFractionalOffset]: A Boolean value that indicates whether the anchor’s offset is expressed as a fraction of its supplementary item’s dimension.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutAnchor
type INSCollectionLayoutAnchor interface {
	objectivec.IObject

	// Topic: Getting the edges

	// The edges of the item an anchor is attached to.
	Edges() NSDirectionalRectEdge

	// Topic: Getting the offset

	// The floating-point value of the anchor’s offset from the item it’s attached to.
	Offset() corefoundation.CGPoint
	// A Boolean value that indicates whether the anchor’s offset is expressed as an absolute value.
	IsAbsoluteOffset() bool
	// A Boolean value that indicates whether the anchor’s offset is expressed as a fraction of its supplementary item’s dimension.
	IsFractionalOffset() bool
}





// Init initializes the instance.
func (c NSCollectionLayoutAnchor) Init() NSCollectionLayoutAnchor {
	rv := objc.Send[NSCollectionLayoutAnchor](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c NSCollectionLayoutAnchor) Autorelease() NSCollectionLayoutAnchor {
	rv := objc.Send[NSCollectionLayoutAnchor](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSCollectionLayoutAnchor creates a new NSCollectionLayoutAnchor instance.
func NewNSCollectionLayoutAnchor() NSCollectionLayoutAnchor {
	class := getNSCollectionLayoutAnchorClass()
	rv := objc.Send[NSCollectionLayoutAnchor](objc.ID(class.class), objc.Sel("new"))
	return rv
}






// Creates an anchor with the specified edges to attach to.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutAnchor/init(edges:)
func NewCollectionLayoutAnchorWithEdges(edges NSDirectionalRectEdge) NSCollectionLayoutAnchor {
	rv := objc.Send[objc.ID](objc.ID(getNSCollectionLayoutAnchorClass().class), objc.Sel("layoutAnchorWithEdges:"), edges)
	return NSCollectionLayoutAnchorFromID(rv)
}


// Creates an anchor with the specified edges to attach to, offset by the
// provided absolute value.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutAnchor/init(edges:absoluteOffset:)
func NewCollectionLayoutAnchorWithEdgesAbsoluteOffset(edges NSDirectionalRectEdge, absoluteOffset corefoundation.CGPoint) NSCollectionLayoutAnchor {
	rv := objc.Send[objc.ID](objc.ID(getNSCollectionLayoutAnchorClass().class), objc.Sel("layoutAnchorWithEdges:absoluteOffset:"), edges, absoluteOffset)
	return NSCollectionLayoutAnchorFromID(rv)
}


// Creates an anchor with the specified edges to attach to, offset by the
// provided fractional value.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutAnchor/init(edges:fractionalOffset:)
func NewCollectionLayoutAnchorWithEdgesFractionalOffset(edges NSDirectionalRectEdge, fractionalOffset corefoundation.CGPoint) NSCollectionLayoutAnchor {
	rv := objc.Send[objc.ID](objc.ID(getNSCollectionLayoutAnchorClass().class), objc.Sel("layoutAnchorWithEdges:fractionalOffset:"), edges, fractionalOffset)
	return NSCollectionLayoutAnchorFromID(rv)
}


















// The edges of the item an anchor is attached to.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutAnchor/edges
func (c NSCollectionLayoutAnchor) Edges() NSDirectionalRectEdge {
	rv := objc.Send[NSDirectionalRectEdge](c.ID, objc.Sel("edges"))
	return NSDirectionalRectEdge(rv)
}



// The floating-point value of the anchor’s offset from the item it’s
// attached to.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutAnchor/offset
func (c NSCollectionLayoutAnchor) Offset() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](c.ID, objc.Sel("offset"))
	return corefoundation.CGPoint(rv)
}



// A Boolean value that indicates whether the anchor’s offset is expressed
// as an absolute value.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutAnchor/isAbsoluteOffset
func (c NSCollectionLayoutAnchor) IsAbsoluteOffset() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isAbsoluteOffset"))
	return rv
}



// A Boolean value that indicates whether the anchor’s offset is expressed
// as a fraction of its supplementary item’s dimension.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutAnchor/isFractionalOffset
func (c NSCollectionLayoutAnchor) IsFractionalOffset() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isFractionalOffset"))
	return rv
}

























