// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSLayoutGuide] class.
var (
	_NSLayoutGuideClass     NSLayoutGuideClass
	_NSLayoutGuideClassOnce sync.Once
)

func getNSLayoutGuideClass() NSLayoutGuideClass {
	_NSLayoutGuideClassOnce.Do(func() {
		_NSLayoutGuideClass = NSLayoutGuideClass{class: objc.GetClass("NSLayoutGuide")}
	})
	return _NSLayoutGuideClass
}

// GetNSLayoutGuideClass returns the class object for NSLayoutGuide.
func GetNSLayoutGuideClass() NSLayoutGuideClass {
	return getNSLayoutGuideClass()
}

type NSLayoutGuideClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSLayoutGuideClass) Alloc() NSLayoutGuide {
	rv := objc.Send[NSLayoutGuide](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A rectangular area that can interact with Auto Layout.
//
// # Overview
// 
// Use layout guides to replace the placeholder views you may have created to
// represent inter-view spaces or encapsulation in your user interface.
// Traditionally, there were a number of Auto Layout techniques that required
// placeholder views. A placeholder view is an empty view that does not have
// any visual elements of its own and serves only to define a rectangular
// region in the view hierarchy. For example, if you wanted to use constraints
// to define the size or location of an empty space between views, you needed
// to use a placeholder view to represent that space. If you wanted to center
// a group of objects, you needed a placeholder view to contain those objects.
// Similarly, placeholder views could be used to contain and encapsulate part
// of your user interface. Placeholder views let you break up a large, complex
// user interface into self-contained, modular chunks. When used properly,
// they could greatly simplify your Auto Layout constraint logic.
// 
// There are a number of costs associated with adding placeholder views to
// your view hierarchy. First, there is the cost of creating and maintaining
// the view itself. Second, the placeholder view is a full member of the view
// hierarchy, which means that it adds overhead to every task the hierarchy
// performs. Worst of all, the invisible placeholder view can intercept
// messages that are intended for other views, causing problems that are very
// difficult to find.
// 
// The [NSLayoutGuide] class is designed to perform all the tasks previously
// performed by placeholder views, but to do it in a safer, more efficient
// manner. Layout guides are not views. They do not use as much memory, and
// they do not participate in the view hierarchy. Instead, they simply define
// a rectangular region in their owning view’s coordinate system that can
// interact with Auto Layout.
// 
// # Creating Layout Guides
// 
// To create a layout guide, perform the following steps:
// 
// - Instantiate a new layout guide. - Add the layout guide to a view by
// calling the view’s [AddLayoutGuide] method . - Define the position and
// size of the layout guide using Auto Layout.
// 
// You can use these guides to define the space between elements in your
// layout. The following example shows how to use layout guides to define an
// equal spacing between a series of views.
// 
// A layout guide can also act as an opaque box that contains other views and
// controls, letting you encapsulate parts of your view and break up your
// layout into modular chunks.
//
// # Working With Layout Guides
//
//   - [NSLayoutGuide.Frame]: The layout guide’s frame in its owning view’s coordinate system.
//   - [NSLayoutGuide.OwningView]: The view that owns this layout guide.
//   - [NSLayoutGuide.SetOwningView]
//
// # Creating Constraints Using Layout Anchors
//
//   - [NSLayoutGuide.BottomAnchor]: A layout anchor representing the bottom edge of the layout guide’s frame.
//   - [NSLayoutGuide.CenterXAnchor]: A layout anchor representing the horizontal center of the layout guide’s frame.
//   - [NSLayoutGuide.CenterYAnchor]: A layout anchor representing the vertical center of the layout guide’s frame.
//   - [NSLayoutGuide.HeightAnchor]: A layout anchor representing the height of the layout guide’s frame.
//   - [NSLayoutGuide.LeadingAnchor]: A layout anchor representing the leading edge of the layout guide’s frame.
//   - [NSLayoutGuide.LeftAnchor]: A layout anchor representing the left edge of the layout guide’s frame.
//   - [NSLayoutGuide.RightAnchor]: A layout anchor representing the right edge of the layout guide’s frame.
//   - [NSLayoutGuide.TopAnchor]: A layout anchor representing the top edge of the layout guide’s frame.
//   - [NSLayoutGuide.TrailingAnchor]: A layout anchor representing the trailing edge of the layout guide’s frame.
//   - [NSLayoutGuide.WidthAnchor]: A layout anchor representing the width of the layout guide’s frame.
//
// # Instance Properties
//
//   - [NSLayoutGuide.HasAmbiguousLayout]
//
// # Instance Methods
//
//   - [NSLayoutGuide.ConstraintsAffectingLayoutForOrientation]
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutGuide
type NSLayoutGuide struct {
	objectivec.Object
}

// NSLayoutGuideFromID constructs a [NSLayoutGuide] from an objc.ID.
//
// A rectangular area that can interact with Auto Layout.
func NSLayoutGuideFromID(id objc.ID) NSLayoutGuide {
	return NSLayoutGuide{objectivec.Object{ID: id}}
}
// NOTE: NSLayoutGuide adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSLayoutGuide] class.
//
// # Working With Layout Guides
//
//   - [INSLayoutGuide.Frame]: The layout guide’s frame in its owning view’s coordinate system.
//   - [INSLayoutGuide.OwningView]: The view that owns this layout guide.
//   - [INSLayoutGuide.SetOwningView]
//
// # Creating Constraints Using Layout Anchors
//
//   - [INSLayoutGuide.BottomAnchor]: A layout anchor representing the bottom edge of the layout guide’s frame.
//   - [INSLayoutGuide.CenterXAnchor]: A layout anchor representing the horizontal center of the layout guide’s frame.
//   - [INSLayoutGuide.CenterYAnchor]: A layout anchor representing the vertical center of the layout guide’s frame.
//   - [INSLayoutGuide.HeightAnchor]: A layout anchor representing the height of the layout guide’s frame.
//   - [INSLayoutGuide.LeadingAnchor]: A layout anchor representing the leading edge of the layout guide’s frame.
//   - [INSLayoutGuide.LeftAnchor]: A layout anchor representing the left edge of the layout guide’s frame.
//   - [INSLayoutGuide.RightAnchor]: A layout anchor representing the right edge of the layout guide’s frame.
//   - [INSLayoutGuide.TopAnchor]: A layout anchor representing the top edge of the layout guide’s frame.
//   - [INSLayoutGuide.TrailingAnchor]: A layout anchor representing the trailing edge of the layout guide’s frame.
//   - [INSLayoutGuide.WidthAnchor]: A layout anchor representing the width of the layout guide’s frame.
//
// # Instance Properties
//
//   - [INSLayoutGuide.HasAmbiguousLayout]
//
// # Instance Methods
//
//   - [INSLayoutGuide.ConstraintsAffectingLayoutForOrientation]
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutGuide
type INSLayoutGuide interface {
	objectivec.IObject
	NSUserInterfaceItemIdentification

	// Topic: Working With Layout Guides

	// The layout guide’s frame in its owning view’s coordinate system.
	Frame() corefoundation.CGRect
	// The view that owns this layout guide.
	OwningView() INSView
	SetOwningView(value INSView)

	// Topic: Creating Constraints Using Layout Anchors

	// A layout anchor representing the bottom edge of the layout guide’s frame.
	BottomAnchor() INSLayoutYAxisAnchor
	// A layout anchor representing the horizontal center of the layout guide’s frame.
	CenterXAnchor() INSLayoutXAxisAnchor
	// A layout anchor representing the vertical center of the layout guide’s frame.
	CenterYAnchor() INSLayoutYAxisAnchor
	// A layout anchor representing the height of the layout guide’s frame.
	HeightAnchor() INSLayoutDimension
	// A layout anchor representing the leading edge of the layout guide’s frame.
	LeadingAnchor() INSLayoutXAxisAnchor
	// A layout anchor representing the left edge of the layout guide’s frame.
	LeftAnchor() INSLayoutXAxisAnchor
	// A layout anchor representing the right edge of the layout guide’s frame.
	RightAnchor() INSLayoutXAxisAnchor
	// A layout anchor representing the top edge of the layout guide’s frame.
	TopAnchor() INSLayoutYAxisAnchor
	// A layout anchor representing the trailing edge of the layout guide’s frame.
	TrailingAnchor() INSLayoutXAxisAnchor
	// A layout anchor representing the width of the layout guide’s frame.
	WidthAnchor() INSLayoutDimension

	// Topic: Instance Properties

	HasAmbiguousLayout() bool

	// Topic: Instance Methods

	ConstraintsAffectingLayoutForOrientation(orientation NSLayoutConstraintOrientation) []NSLayoutConstraint

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (l NSLayoutGuide) Init() NSLayoutGuide {
	rv := objc.Send[NSLayoutGuide](l.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (l NSLayoutGuide) Autorelease() NSLayoutGuide {
	rv := objc.Send[NSLayoutGuide](l.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSLayoutGuide creates a new NSLayoutGuide instance.
func NewNSLayoutGuide() NSLayoutGuide {
	class := getNSLayoutGuideClass()
	rv := objc.Send[NSLayoutGuide](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutGuide/constraintsAffectingLayout(for:)
func (l NSLayoutGuide) ConstraintsAffectingLayoutForOrientation(orientation NSLayoutConstraintOrientation) []NSLayoutConstraint {
	rv := objc.Send[[]objc.ID](l.ID, objc.Sel("constraintsAffectingLayoutForOrientation:"), orientation)
	return objc.ConvertSlice(rv, func(id objc.ID) NSLayoutConstraint {
		return NSLayoutConstraintFromID(id)
	})
}
func (l NSLayoutGuide) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](l.ID, objc.Sel("encodeWithCoder:"), coder)
}

// A string used to identify the layout guide.
//
// # Discussion
// 
// By default, the `identifier` property is `nil`. You can assign a string to
// help identify this guide. This string appears as part of the guide’s
// description when the guide is printed to the console. You can also use the
// identifier to find a particular layout guide from among the guides owned by
// a view.
// 
// Identifiers starting with “NS” or “UI” are reserved by the system.
// The system may assign these identifiers to the guides it creates.
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutGuide/identifier
func (l NSLayoutGuide) Identifier() NSUserInterfaceItemIdentifier {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("identifier"))
	return NSUserInterfaceItemIdentifier(foundation.NSStringFromID(rv).String())
}
func (l NSLayoutGuide) SetIdentifier(value NSUserInterfaceItemIdentifier) {
	objc.Send[struct{}](l.ID, objc.Sel("setIdentifier:"), objc.String(string(value)))
}
// The layout guide’s frame in its owning view’s coordinate system.
//
// # Discussion
// 
// The layout guide defines a rectangular space in its owning view’s
// coordinate system. This property contains a valid [CGRect] value by the
// time its owning view’s [Layout] method is called.
//
// [CGRect]: https://developer.apple.com/documentation/CoreFoundation/CGRect
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutGuide/frame
func (l NSLayoutGuide) Frame() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](l.ID, objc.Sel("frame"))
	return corefoundation.CGRect(rv)
}
// The view that owns this layout guide.
//
// # Discussion
// 
// By default, this property is `nil`. To participate in Auto Layout, the
// layout guide must be added to a view by calling its [AddLayoutGuide]
// method. Do not modify this property directly. Instead, use the view’s
// [AddLayoutGuide] and [RemoveLayoutGuide] methods, which update this
// property as necessary.
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutGuide/owningView
func (l NSLayoutGuide) OwningView() INSView {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("owningView"))
	return NSViewFromID(objc.ID(rv))
}
func (l NSLayoutGuide) SetOwningView(value INSView) {
	objc.Send[struct{}](l.ID, objc.Sel("setOwningView:"), value)
}
// A layout anchor representing the bottom edge of the layout guide’s frame.
//
// # Discussion
// 
// Use this anchor to create constraints with the layout guide’s bottom
// edge. You can combine this anchor only with other [NSLayoutYAxisAnchor]
// anchors. For more information, see [NSLayoutAnchor].
//
// [NSLayoutAnchor]: https://developer.apple.com/documentation/UIKit/NSLayoutAnchor
// [NSLayoutYAxisAnchor]: https://developer.apple.com/documentation/UIKit/NSLayoutYAxisAnchor
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutGuide/bottomAnchor
func (l NSLayoutGuide) BottomAnchor() INSLayoutYAxisAnchor {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("bottomAnchor"))
	return NSLayoutYAxisAnchorFromID(objc.ID(rv))
}
// A layout anchor representing the horizontal center of the layout guide’s
// frame.
//
// # Discussion
// 
// Use this anchor to create constraints with the layout guide’s horizontal
// center. You can combine this anchor only with other [NSLayoutXAxisAnchor]
// anchors. For more information, see [NSLayoutAnchor].
//
// [NSLayoutAnchor]: https://developer.apple.com/documentation/UIKit/NSLayoutAnchor
// [NSLayoutXAxisAnchor]: https://developer.apple.com/documentation/UIKit/NSLayoutXAxisAnchor
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutGuide/centerXAnchor
func (l NSLayoutGuide) CenterXAnchor() INSLayoutXAxisAnchor {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("centerXAnchor"))
	return NSLayoutXAxisAnchorFromID(objc.ID(rv))
}
// A layout anchor representing the vertical center of the layout guide’s
// frame.
//
// # Discussion
// 
// Use this anchor to create constraints with the layout guide’s vertical
// center. You can combine this anchor only with other [NSLayoutYAxisAnchor]
// anchors. For more information, see [NSLayoutAnchor].
//
// [NSLayoutAnchor]: https://developer.apple.com/documentation/UIKit/NSLayoutAnchor
// [NSLayoutYAxisAnchor]: https://developer.apple.com/documentation/UIKit/NSLayoutYAxisAnchor
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutGuide/centerYAnchor
func (l NSLayoutGuide) CenterYAnchor() INSLayoutYAxisAnchor {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("centerYAnchor"))
	return NSLayoutYAxisAnchorFromID(objc.ID(rv))
}
// A layout anchor representing the height of the layout guide’s frame.
//
// # Discussion
// 
// Use this anchor to create constraints with the layout guide’s height. You
// can combine this anchor only with other [NSLayoutDimension] anchors. For
// more information, see [NSLayoutAnchor].
//
// [NSLayoutAnchor]: https://developer.apple.com/documentation/UIKit/NSLayoutAnchor
// [NSLayoutDimension]: https://developer.apple.com/documentation/UIKit/NSLayoutDimension
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutGuide/heightAnchor
func (l NSLayoutGuide) HeightAnchor() INSLayoutDimension {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("heightAnchor"))
	return NSLayoutDimensionFromID(objc.ID(rv))
}
// A layout anchor representing the leading edge of the layout guide’s
// frame.
//
// # Discussion
// 
// Use this anchor to create constraints with the layout guide’s leading
// edge. You can combine this anchor only with a subset of the
// [NSLayoutXAxisAnchor] anchors. You can combine a [LeadingAnchor] with
// another `leadingAnchor`, a `trailingAnchor`, or a `centerXAnchor`. For more
// information, see [NSLayoutAnchor].
//
// [NSLayoutAnchor]: https://developer.apple.com/documentation/UIKit/NSLayoutAnchor
// [NSLayoutXAxisAnchor]: https://developer.apple.com/documentation/UIKit/NSLayoutXAxisAnchor
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutGuide/leadingAnchor
func (l NSLayoutGuide) LeadingAnchor() INSLayoutXAxisAnchor {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("leadingAnchor"))
	return NSLayoutXAxisAnchorFromID(objc.ID(rv))
}
// A layout anchor representing the left edge of the layout guide’s frame.
//
// # Discussion
// 
// Use this anchor to create constraints with the layout guide’s left edge.
// You can combine this anchor only with a subset of the [NSLayoutXAxisAnchor]
// anchors. You can combine a [LeftAnchor] with another `leftAnchor`, a
// `rightAnchor`, or a `centerXAnchor`. For more information, see
// [NSLayoutAnchor].
//
// [NSLayoutAnchor]: https://developer.apple.com/documentation/UIKit/NSLayoutAnchor
// [NSLayoutXAxisAnchor]: https://developer.apple.com/documentation/UIKit/NSLayoutXAxisAnchor
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutGuide/leftAnchor
func (l NSLayoutGuide) LeftAnchor() INSLayoutXAxisAnchor {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("leftAnchor"))
	return NSLayoutXAxisAnchorFromID(objc.ID(rv))
}
// A layout anchor representing the right edge of the layout guide’s frame.
//
// # Discussion
// 
// Use this anchor to create constraints with the layout guide’s right edge.
// You can combine this anchor only with a subset of the [NSLayoutXAxisAnchor]
// anchors. You can combine a [RightAnchor] with another `rightAnchor`, a
// `leftAnchor`, or a `centerXAnchor`. For more information, see
// [NSLayoutAnchor].
//
// [NSLayoutAnchor]: https://developer.apple.com/documentation/UIKit/NSLayoutAnchor
// [NSLayoutXAxisAnchor]: https://developer.apple.com/documentation/UIKit/NSLayoutXAxisAnchor
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutGuide/rightAnchor
func (l NSLayoutGuide) RightAnchor() INSLayoutXAxisAnchor {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("rightAnchor"))
	return NSLayoutXAxisAnchorFromID(objc.ID(rv))
}
// A layout anchor representing the top edge of the layout guide’s frame.
//
// # Discussion
// 
// Use this anchor to create constraints with the layout guide’s top edge.
// You can combine this anchor only with other [NSLayoutYAxisAnchor] anchors.
// For more information, see [NSLayoutAnchor].
//
// [NSLayoutAnchor]: https://developer.apple.com/documentation/UIKit/NSLayoutAnchor
// [NSLayoutYAxisAnchor]: https://developer.apple.com/documentation/UIKit/NSLayoutYAxisAnchor
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutGuide/topAnchor
func (l NSLayoutGuide) TopAnchor() INSLayoutYAxisAnchor {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("topAnchor"))
	return NSLayoutYAxisAnchorFromID(objc.ID(rv))
}
// A layout anchor representing the trailing edge of the layout guide’s
// frame.
//
// # Discussion
// 
// Use this anchor to create constraints with the layout guide’s trailing
// edge. You can combine this anchor only with a subset of the
// [NSLayoutXAxisAnchor] anchors. You can combine a [TrailingAnchor] with
// another `trailingAnchor`, a `leadingAnchor`, or a `centerXAnchor`. For more
// information, see [NSLayoutAnchor].
//
// [NSLayoutAnchor]: https://developer.apple.com/documentation/UIKit/NSLayoutAnchor
// [NSLayoutXAxisAnchor]: https://developer.apple.com/documentation/UIKit/NSLayoutXAxisAnchor
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutGuide/trailingAnchor
func (l NSLayoutGuide) TrailingAnchor() INSLayoutXAxisAnchor {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("trailingAnchor"))
	return NSLayoutXAxisAnchorFromID(objc.ID(rv))
}
// A layout anchor representing the width of the layout guide’s frame.
//
// # Discussion
// 
// Use this anchor to create constraints with the layout guide’s width. You
// can combine this anchor only with other [NSLayoutDimension] anchors. For
// more information, see [NSLayoutAnchor].
//
// [NSLayoutAnchor]: https://developer.apple.com/documentation/UIKit/NSLayoutAnchor
// [NSLayoutDimension]: https://developer.apple.com/documentation/UIKit/NSLayoutDimension
//
// See: https://developer.apple.com/documentation/AppKit/NSLayoutGuide/widthAnchor
func (l NSLayoutGuide) WidthAnchor() INSLayoutDimension {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("widthAnchor"))
	return NSLayoutDimensionFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/AppKit/NSLayoutGuide/hasAmbiguousLayout
func (l NSLayoutGuide) HasAmbiguousLayout() bool {
	rv := objc.Send[bool](l.ID, objc.Sel("hasAmbiguousLayout"))
	return rv
}

			// Protocol methods for NSUserInterfaceItemIdentification
			

