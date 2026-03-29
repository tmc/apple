// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
)

// The class instance for the [NSClipView] class.
var (
	_NSClipViewClass     NSClipViewClass
	_NSClipViewClassOnce sync.Once
)

func getNSClipViewClass() NSClipViewClass {
	_NSClipViewClassOnce.Do(func() {
		_NSClipViewClass = NSClipViewClass{class: objc.GetClass("NSClipView")}
	})
	return _NSClipViewClass
}

// GetNSClipViewClass returns the class object for NSClipView.
func GetNSClipViewClass() NSClipViewClass {
	return getNSClipViewClass()
}

type NSClipViewClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSClipViewClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSClipViewClass) Alloc() NSClipView {
	rv := objc.Send[NSClipView](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// An object that clips a document view to a scroll view’s frame.
//
// # Overview
// 
// An [NSClipView] holds the document view of an [NSScrollView], clipping the
// document view to its frame, handling the details of scrolling in an
// efficient manner, and updating the [NSScrollView] when the document
// view’s size or position changes.
// 
// You don’t typically use the [NSClipView] class directly; it’s provided
// primarily as the scrolling machinery for the [NSScrollView] class. However,
// you might use the [NSClipView] class to implement a class similar to
// [NSScrollView].
// 
// # Interaction with NSScrollView
// 
// When using an [NSClipView] within an [NSScrollView] (the usual
// configuration), you should access the [NSScrollView] properties that
// control background drawing state, rather than accessing these properties of
// the [NSClipView]. This recommendation applies to the following properties:
// 
// - [NSClipView.BackgroundColor]
// - [NSClipView.DrawsBackground]
// 
// The [NSClipView] methods are intended for when the [NSClipView] is used
// independently of a containing [NSScrollView]. In the usual case,
// [NSScrollView] should be allowed to manage the background-drawing
// properties of its associated [NSClipView].
// 
// There is only one background-drawing state per [NSScrollView]/[NSClipView]
// pair. The two objects do not maintain independent and distinct
// [NSClipView.DrawsBackground] and [NSClipView.BackgroundColor] properties; rather, the accessors
// for these properties on [NSScrollView] largely defer to the associated
// [NSClipView] and allow the [NSClipView] to maintain the state. Note that
// this state is not cached by the [NSScrollView] object.
// 
// It is also important to note that setting [NSClipView.DrawsBackground] to [false] in
// an [NSScrollView] has the added effect of setting the [NSClipView] property
// [copiesOnScroll] to [false]. The side effect of setting the
// [NSClipView.DrawsBackground] property directly to the [NSClipView] is the appearance
// of “trails” (vestiges of previous drawing) in the document view as it
// is scrolled.
//
// [copiesOnScroll]: https://developer.apple.com/documentation/AppKit/NSClipView/copiesOnScroll
// [false]: https://developer.apple.com/documentation/Swift/false
//
// # Setting the Document View
//
//   - [NSClipView.DocumentView]: The clip view’s document view.
//   - [NSClipView.SetDocumentView]
//
// # Scrolling
//
//   - [NSClipView.ScrollToPoint]: Changes the origin of the clip view’s bounds rectangle to `newOrigin`.
//   - [NSClipView.ConstrainBoundsRect]: Constrains the bounds of the clip view while the user is magnifying and scrolling.
//
// # Accessing the Content Insets
//
//   - [NSClipView.ContentInsets]: The distance that the content view is inset from the enclosing scroll view.
//   - [NSClipView.SetContentInsets]
//   - [NSClipView.AutomaticallyAdjustsContentInsets]: A Boolean value that indicates if the clip view automatically accounts for other scroll view subviews.
//   - [NSClipView.SetAutomaticallyAdjustsContentInsets]
//
// # Accessing the Visible Portion
//
//   - [NSClipView.DocumentRect]: The rectangle defining the document view’s frame, adjusted to the size of the clip view if the document view is smaller.
//   - [NSClipView.DocumentVisibleRect]: The exposed rectangle of the clip view’s document view, in the document view’s own coordinate system.
//
// # Setting the Document Cursor
//
//   - [NSClipView.DocumentCursor]: The cursor object used when the pointer lies over the view.
//   - [NSClipView.SetDocumentCursor]
//
// # Working with Background Color
//
//   - [NSClipView.DrawsBackground]: A Boolean value that indicates if the clip view draws its background color.
//   - [NSClipView.SetDrawsBackground]
//   - [NSClipView.BackgroundColor]: The color of the clip view’s background.
//   - [NSClipView.SetBackgroundColor]
//
// # Overriding NSView Methods
//
//   - [NSClipView.ViewBoundsChanged]: Handles an [boundsDidChangeNotification](<doc://com.apple.appkit/documentation/AppKit/NSView/boundsDidChangeNotification>), passed in the `aNotification` argument, by updating a containing [NSScrollView](<doc://com.apple.appkit/documentation/AppKit/NSScrollView>) based on the new bounds.
//   - [NSClipView.ViewFrameChanged]: Handles an [frameDidChangeNotification](<doc://com.apple.appkit/documentation/AppKit/NSView/frameDidChangeNotification>), passed in the `aNotification` argument, by updating a containing [NSScrollView](<doc://com.apple.appkit/documentation/AppKit/NSScrollView>) based on the new frame.
//
// See: https://developer.apple.com/documentation/AppKit/NSClipView
type NSClipView struct {
	NSView
}

// NSClipViewFromID constructs a [NSClipView] from an objc.ID.
//
// An object that clips a document view to a scroll view’s frame.
func NSClipViewFromID(id objc.ID) NSClipView {
	return NSClipView{NSView: NSViewFromID(id)}
}
// NOTE: NSClipView adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSClipView] class.
//
// # Setting the Document View
//
//   - [INSClipView.DocumentView]: The clip view’s document view.
//   - [INSClipView.SetDocumentView]
//
// # Scrolling
//
//   - [INSClipView.ScrollToPoint]: Changes the origin of the clip view’s bounds rectangle to `newOrigin`.
//   - [INSClipView.ConstrainBoundsRect]: Constrains the bounds of the clip view while the user is magnifying and scrolling.
//
// # Accessing the Content Insets
//
//   - [INSClipView.ContentInsets]: The distance that the content view is inset from the enclosing scroll view.
//   - [INSClipView.SetContentInsets]
//   - [INSClipView.AutomaticallyAdjustsContentInsets]: A Boolean value that indicates if the clip view automatically accounts for other scroll view subviews.
//   - [INSClipView.SetAutomaticallyAdjustsContentInsets]
//
// # Accessing the Visible Portion
//
//   - [INSClipView.DocumentRect]: The rectangle defining the document view’s frame, adjusted to the size of the clip view if the document view is smaller.
//   - [INSClipView.DocumentVisibleRect]: The exposed rectangle of the clip view’s document view, in the document view’s own coordinate system.
//
// # Setting the Document Cursor
//
//   - [INSClipView.DocumentCursor]: The cursor object used when the pointer lies over the view.
//   - [INSClipView.SetDocumentCursor]
//
// # Working with Background Color
//
//   - [INSClipView.DrawsBackground]: A Boolean value that indicates if the clip view draws its background color.
//   - [INSClipView.SetDrawsBackground]
//   - [INSClipView.BackgroundColor]: The color of the clip view’s background.
//   - [INSClipView.SetBackgroundColor]
//
// # Overriding NSView Methods
//
//   - [INSClipView.ViewBoundsChanged]: Handles an [boundsDidChangeNotification](<doc://com.apple.appkit/documentation/AppKit/NSView/boundsDidChangeNotification>), passed in the `aNotification` argument, by updating a containing [NSScrollView](<doc://com.apple.appkit/documentation/AppKit/NSScrollView>) based on the new bounds.
//   - [INSClipView.ViewFrameChanged]: Handles an [frameDidChangeNotification](<doc://com.apple.appkit/documentation/AppKit/NSView/frameDidChangeNotification>), passed in the `aNotification` argument, by updating a containing [NSScrollView](<doc://com.apple.appkit/documentation/AppKit/NSScrollView>) based on the new frame.
//
// See: https://developer.apple.com/documentation/AppKit/NSClipView
type INSClipView interface {
	INSView

	// Topic: Setting the Document View

	// The clip view’s document view.
	DocumentView() INSView
	SetDocumentView(value INSView)

	// Topic: Scrolling

	// Changes the origin of the clip view’s bounds rectangle to `newOrigin`.
	ScrollToPoint(newOrigin corefoundation.CGPoint)
	// Constrains the bounds of the clip view while the user is magnifying and scrolling.
	ConstrainBoundsRect(proposedBounds corefoundation.CGRect) corefoundation.CGRect

	// Topic: Accessing the Content Insets

	// The distance that the content view is inset from the enclosing scroll view.
	ContentInsets() foundation.NSEdgeInsets
	SetContentInsets(value foundation.NSEdgeInsets)
	// A Boolean value that indicates if the clip view automatically accounts for other scroll view subviews.
	AutomaticallyAdjustsContentInsets() bool
	SetAutomaticallyAdjustsContentInsets(value bool)

	// Topic: Accessing the Visible Portion

	// The rectangle defining the document view’s frame, adjusted to the size of the clip view if the document view is smaller.
	DocumentRect() corefoundation.CGRect
	// The exposed rectangle of the clip view’s document view, in the document view’s own coordinate system.
	DocumentVisibleRect() corefoundation.CGRect

	// Topic: Setting the Document Cursor

	// The cursor object used when the pointer lies over the view.
	DocumentCursor() INSCursor
	SetDocumentCursor(value INSCursor)

	// Topic: Working with Background Color

	// A Boolean value that indicates if the clip view draws its background color.
	DrawsBackground() bool
	SetDrawsBackground(value bool)
	// The color of the clip view’s background.
	BackgroundColor() INSColor
	SetBackgroundColor(value INSColor)

	// Topic: Overriding NSView Methods

	// Handles an [boundsDidChangeNotification](<doc://com.apple.appkit/documentation/AppKit/NSView/boundsDidChangeNotification>), passed in the `aNotification` argument, by updating a containing [NSScrollView](<doc://com.apple.appkit/documentation/AppKit/NSScrollView>) based on the new bounds.
	ViewBoundsChanged(notification foundation.NSNotification)
	// Handles an [frameDidChangeNotification](<doc://com.apple.appkit/documentation/AppKit/NSView/frameDidChangeNotification>), passed in the `aNotification` argument, by updating a containing [NSScrollView](<doc://com.apple.appkit/documentation/AppKit/NSScrollView>) based on the new frame.
	ViewFrameChanged(notification foundation.NSNotification)
}

// Init initializes the instance.
func (c NSClipView) Init() NSClipView {
	rv := objc.Send[NSClipView](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c NSClipView) Autorelease() NSClipView {
	rv := objc.Send[NSClipView](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSClipView creates a new NSClipView instance.
func NewNSClipView() NSClipView {
	class := getNSClipViewClass()
	rv := objc.Send[NSClipView](objc.ID(class.class), objc.Sel("new"))
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
func NewClipViewWithCoder(coder foundation.INSCoder) NSClipView {
	instance := getNSClipViewClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSClipViewFromID(rv)
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
func NewClipViewWithFrame(frameRect corefoundation.CGRect) NSClipView {
	instance := getNSClipViewClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFrame:"), frameRect)
	return NSClipViewFromID(rv)
}

// Changes the origin of the clip view’s bounds rectangle to `newOrigin`.
//
// See: https://developer.apple.com/documentation/AppKit/NSClipView/scroll(to:)
func (c NSClipView) ScrollToPoint(newOrigin corefoundation.CGPoint) {
	objc.Send[objc.ID](c.ID, objc.Sel("scrollToPoint:"), newOrigin)
}
// Constrains the bounds of the clip view while the user is magnifying and
// scrolling.
//
// proposedBounds: The bounds to use to ensure that the view will still lie within its
// document view.
//
// # Return Value
// 
// A bounds rectangle.
//
// # Discussion
// 
// Note that you can move an implementation of the deprecated
// [constrainScroll(_:)] to this method by adjusting the origin of
// `proposedBounds` (instead of using the `newOrigin` parameter in `-`). To
// preserve compatibility, if a subclass overrides `-`, the default behavior
// of [ConstrainBoundsRect] will be to use that `-` to adjust the origin of
// `proposedBounds`, and to not change the size.
//
// [constrainScroll(_:)]: https://developer.apple.com/documentation/AppKit/NSClipView/constrainScroll(_:)
//
// See: https://developer.apple.com/documentation/AppKit/NSClipView/constrainBoundsRect(_:)
func (c NSClipView) ConstrainBoundsRect(proposedBounds corefoundation.CGRect) corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](c.ID, objc.Sel("constrainBoundsRect:"), proposedBounds)
	return corefoundation.CGRect(rv)
}
// Handles an [boundsDidChangeNotification], passed in the `aNotification`
// argument, by updating a containing [NSScrollView] based on the new bounds.
//
// [boundsDidChangeNotification]: https://developer.apple.com/documentation/AppKit/NSView/boundsDidChangeNotification
//
// See: https://developer.apple.com/documentation/AppKit/NSClipView/viewBoundsChanged(_:)
func (c NSClipView) ViewBoundsChanged(notification foundation.NSNotification) {
	objc.Send[objc.ID](c.ID, objc.Sel("viewBoundsChanged:"), notification)
}
// Handles an [frameDidChangeNotification], passed in the `aNotification`
// argument, by updating a containing [NSScrollView] based on the new frame.
//
// [frameDidChangeNotification]: https://developer.apple.com/documentation/AppKit/NSView/frameDidChangeNotification
//
// See: https://developer.apple.com/documentation/AppKit/NSClipView/viewFrameChanged(_:)
func (c NSClipView) ViewFrameChanged(notification foundation.NSNotification) {
	objc.Send[objc.ID](c.ID, objc.Sel("viewFrameChanged:"), notification)
}

// The clip view’s document view.
//
// # Discussion
// 
// If the clip view is contained in an [NSScrollView], you should send the
// [NSScrollView] a [DocumentView] message instead, so it can perform whatever
// updating it needs. Setting this property to a view removes any previous
// document view, and sets the origin of the clip view’s bounds rectangle to
// the origin of the new view’s frame rectangle. Doing so also registers the
// clip view for the notifications [frameDidChangeNotification] and
// [boundsDidChangeNotification], adjusts the key view loop to include the new
// document view, and updates a parent [NSScrollView] display if needed using
// [ReflectScrolledClipView].
//
// [boundsDidChangeNotification]: https://developer.apple.com/documentation/AppKit/NSView/boundsDidChangeNotification
// [frameDidChangeNotification]: https://developer.apple.com/documentation/AppKit/NSView/frameDidChangeNotification
//
// See: https://developer.apple.com/documentation/AppKit/NSClipView/documentView
func (c NSClipView) DocumentView() INSView {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("documentView"))
	return NSViewFromID(objc.ID(rv))
}
func (c NSClipView) SetDocumentView(value INSView) {
	objc.Send[struct{}](c.ID, objc.Sel("setDocumentView:"), value)
}
// The distance that the content view is inset from the enclosing scroll view.
//
// # Discussion
// 
// When the enclosing scroll view’s [ContentInsets] value is nonzero (that
// is, the value is not {0,0,0,0}), the scroll view sets the frame of its
// content view to the scroll view’s bounds minus the scroll view’s
// border, if it has one. (When the [ContentInsets] value is equal to zero,
// the scroll view adjusts its `contentView.Frame()` to fit inside all the
// other views the scroll view maintains.) When the value of
// `contentView.AutomaticallyAdjustsContentInsets()` is [true] (which is the
// default value), the header, rulers, and other views are overlaid on top of
// the content view and the scroll view sets the correct [ContentInsets] value
// on the [ContentView]. Note that you can animate the clip view when this
// property changes by calling `[self animator]`.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSClipView/contentInsets
func (c NSClipView) ContentInsets() foundation.NSEdgeInsets {
	rv := objc.Send[foundation.NSEdgeInsets](c.ID, objc.Sel("contentInsets"))
	return foundation.NSEdgeInsets(rv)
}
func (c NSClipView) SetContentInsets(value foundation.NSEdgeInsets) {
	objc.Send[struct{}](c.ID, objc.Sel("setContentInsets:"), value)
}
// A Boolean value that indicates if the clip view automatically accounts for
// other scroll view subviews.
//
// # Discussion
// 
// When the value of this property is [true], and the clip view is used as the
// [ContentView] of an [NSScrollView], the clip view automatically accounts
// for other scroll view subviews, such as rulers and headers. The default
// value is [true].
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSClipView/automaticallyAdjustsContentInsets
func (c NSClipView) AutomaticallyAdjustsContentInsets() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("automaticallyAdjustsContentInsets"))
	return rv
}
func (c NSClipView) SetAutomaticallyAdjustsContentInsets(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setAutomaticallyAdjustsContentInsets:"), value)
}
// The rectangle defining the document view’s frame, adjusted to the size of
// the clip view if the document view is smaller.
//
// # Discussion
// 
// The document rectangle is used in conjunction with an [NSClipView]
// object’s bounds rectangle to determine values for the indicators of
// relative position and size between the [NSClipView] and its document view.
// For example, [NSScrollView] uses these rectangles to set the size and
// position of the knobs in its scrollers. When the document view is much
// larger than the [NSClipView], the knob is small; when the document view is
// near the same size, the knob is large; and when the document view is the
// same size or smaller, there is no knob.
//
// See: https://developer.apple.com/documentation/AppKit/NSClipView/documentRect
func (c NSClipView) DocumentRect() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](c.ID, objc.Sel("documentRect"))
	return corefoundation.CGRect(rv)
}
// The exposed rectangle of the clip view’s document view, in the document
// view’s own coordinate system.
//
// # Discussion
// 
// Note that this rectangle doesn’t reflect the effects of any clipping that
// may occur above the [NSClipView] itself. To get the portion of the document
// view that’s guaranteed to be visible, send it a `visibleRect` message.
//
// See: https://developer.apple.com/documentation/AppKit/NSClipView/documentVisibleRect
func (c NSClipView) DocumentVisibleRect() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](c.ID, objc.Sel("documentVisibleRect"))
	return corefoundation.CGRect(rv)
}
// The cursor object used when the pointer lies over the view.
//
// # Discussion
// 
// The default value of this property is `nil`, unless you specify a value in
// the xib file associated with the clip view (or scroll view). Note that the
// clip view’s document view may specify a cursor for its enclosing scroll
// view by setting [enclosingScrollView].
//
// [enclosingScrollView]: https://developer.apple.com/documentation/AppKit/NSView/enclosingScrollView
//
// See: https://developer.apple.com/documentation/AppKit/NSClipView/documentCursor
func (c NSClipView) DocumentCursor() INSCursor {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("documentCursor"))
	return NSCursorFromID(objc.ID(rv))
}
func (c NSClipView) SetDocumentCursor(value INSCursor) {
	objc.Send[struct{}](c.ID, objc.Sel("setDocumentCursor:"), value)
}
// A Boolean value that indicates if the clip view draws its background color.
//
// # Discussion
// 
// If your [NSClipView] is enclosed in an [NSScrollView], you should set the
// [DrawsBackground] property on the [NSScrollView]. Setting this property to
// [false] on an [NSScrollView] has the added effect of setting the
// [NSClipView] property [copiesOnScroll] to [false]. The side effect of
// setting the [DrawsBackground] property on the [NSClipView] is the
// appearance of “trails” (vestiges of previous drawing) in the document
// view as it is scrolled.
//
// [copiesOnScroll]: https://developer.apple.com/documentation/AppKit/NSClipView/copiesOnScroll
// [false]: https://developer.apple.com/documentation/Swift/false
//
// See: https://developer.apple.com/documentation/AppKit/NSClipView/drawsBackground
func (c NSClipView) DrawsBackground() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("drawsBackground"))
	return rv
}
func (c NSClipView) SetDrawsBackground(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setDrawsBackground:"), value)
}
// The color of the clip view’s background.
//
// # Discussion
// 
// The default value of this property is supplied by the current
// [controlBackgroundColor].
//
// [controlBackgroundColor]: https://developer.apple.com/documentation/AppKit/NSColor/controlBackgroundColor
//
// See: https://developer.apple.com/documentation/AppKit/NSClipView/backgroundColor
func (c NSClipView) BackgroundColor() INSColor {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("backgroundColor"))
	return NSColorFromID(objc.ID(rv))
}
func (c NSClipView) SetBackgroundColor(value INSColor) {
	objc.Send[struct{}](c.ID, objc.Sel("setBackgroundColor:"), value)
}

// A notification that posts when the view’s bounds rectangle changes to a
// new value independently of the frame rectangle.
//
// See: https://developer.apple.com/documentation/appkit/nsview/boundsdidchangenotification
func (_NSClipViewClass NSClipViewClass) BoundsDidChangeNotification() foundation.NSString {
	rv := objc.Send[objc.ID](objc.ID(_NSClipViewClass.class), objc.Sel("NSViewBoundsDidChangeNotification"))
	return foundation.NSStringFromID(objc.ID(rv))
}
// A notification that posts when the view’s frame rectangle changes to a
// new value.
//
// See: https://developer.apple.com/documentation/appkit/nsview/framedidchangenotification
func (_NSClipViewClass NSClipViewClass) FrameDidChangeNotification() foundation.NSString {
	rv := objc.Send[objc.ID](objc.ID(_NSClipViewClass.class), objc.Sel("NSViewFrameDidChangeNotification"))
	return foundation.NSStringFromID(objc.ID(rv))
}

