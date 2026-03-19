// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
)

// The class instance for the [NSScrollView] class.
var (
	_NSScrollViewClass     NSScrollViewClass
	_NSScrollViewClassOnce sync.Once
)

func getNSScrollViewClass() NSScrollViewClass {
	_NSScrollViewClassOnce.Do(func() {
		_NSScrollViewClass = NSScrollViewClass{class: objc.GetClass("NSScrollView")}
	})
	return _NSScrollViewClass
}

// GetNSScrollViewClass returns the class object for NSScrollView.
func GetNSScrollViewClass() NSScrollViewClass {
	return getNSScrollViewClass()
}

type NSScrollViewClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSScrollViewClass) Alloc() NSScrollView {
	rv := objc.Send[NSScrollView](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A view that displays a portion of a document view and provides scroll bars
// that allow the user to move the document view within the scroll view.
//
// # Overview
// 
// The [NSScrollView] class is the central coordinator for AppKit’s
// scrolling machinery, which is composed of this class, and the [NSClipView]
// and [NSScroller] classes.
// 
// When using an [NSClipView] object within a scroll view (the usual
// configuration), you should issue messages that control background drawing
// state to the scroll view directly, rather than messaging the clip view.
//
// # Determining Component Sizes
//
//   - [NSScrollView.ContentSize]: The size of the scroll view’s content view.
//   - [NSScrollView.DocumentVisibleRect]: The portion of the document view, in its own coordinate system, visible through the scroll view’s content view.
//
// # Managing Graphics Attributes
//
//   - [NSScrollView.BackgroundColor]: The color of the content view’s background.
//   - [NSScrollView.SetBackgroundColor]
//   - [NSScrollView.DrawsBackground]: A Boolean that indicates whether the scroll view draws its background.
//   - [NSScrollView.SetDrawsBackground]
//   - [NSScrollView.BorderType]: A value that specifies the appearance of the scroll view’s border.
//   - [NSScrollView.SetBorderType]
//   - [NSScrollView.DocumentCursor]: The content view’s document cursor.
//   - [NSScrollView.SetDocumentCursor]
//
// # Managing the Views
//
//   - [NSScrollView.DocumentView]: The view the scroll view scrolls within its content view.
//   - [NSScrollView.SetDocumentView]
//   - [NSScrollView.AddFloatingSubviewForAxis]: Adds a floating subview to the document view.
//
// # Managing Scrollers
//
//   - [NSScrollView.HorizontalScroller]: The scroll view’s horizontal scroller.
//   - [NSScrollView.SetHorizontalScroller]
//   - [NSScrollView.HasHorizontalScroller]: A Boolean that indicates whether the scroll view has a horizontal scroller.
//   - [NSScrollView.SetHasHorizontalScroller]
//   - [NSScrollView.VerticalScroller]: The scroll view’s vertical scroller.
//   - [NSScrollView.SetVerticalScroller]
//   - [NSScrollView.HasVerticalScroller]: A Boolean that indicates whether the scroll view has a vertical scroller.
//   - [NSScrollView.SetHasVerticalScroller]
//   - [NSScrollView.AutohidesScrollers]: A Boolean that indicates whether the scroll view automatically hides its scroll bars when they are not needed.
//   - [NSScrollView.SetAutohidesScrollers]
//
// # Managing Rulers
//
//   - [NSScrollView.HasHorizontalRuler]: A Boolean that indicates whether the scroll view keeps a horizontal ruler object.
//   - [NSScrollView.SetHasHorizontalRuler]
//   - [NSScrollView.HorizontalRulerView]: The scroll view’s horizontal ruler view.
//   - [NSScrollView.SetHorizontalRulerView]
//   - [NSScrollView.HasVerticalRuler]: A Boolean that indicates whether the scroll view keeps a vertical ruler object.
//   - [NSScrollView.SetHasVerticalRuler]
//   - [NSScrollView.VerticalRulerView]: The scroll view’s vertical ruler view.
//   - [NSScrollView.SetVerticalRulerView]
//   - [NSScrollView.RulersVisible]: A Boolean that indicates whether the scroll view displays its rulers.
//   - [NSScrollView.SetRulersVisible]
//
// # Managing Insets
//
//   - [NSScrollView.AutomaticallyAdjustsContentInsets]: A Boolean that indicates whether the scroll view automatically adjusts its content insets.
//   - [NSScrollView.SetAutomaticallyAdjustsContentInsets]
//   - [NSScrollView.ContentInsets]: The distance that the scroll view’s subviews are inset from the enclosing scroll view during tiling.
//   - [NSScrollView.SetContentInsets]
//   - [NSScrollView.ScrollerInsets]: The distance the scrollers are inset from the edge of the scroll view.
//   - [NSScrollView.SetScrollerInsets]
//
// # Scroller Style
//
//   - [NSScrollView.ScrollerKnobStyle]: The knob style of scroll views that use the overlay scroller style.
//   - [NSScrollView.SetScrollerKnobStyle]
//   - [NSScrollView.ScrollerStyle]: The scroller style used by the scroll view.
//   - [NSScrollView.SetScrollerStyle]
//
// # Setting Scrolling Behavior
//
//   - [NSScrollView.LineScroll]: The scroll view’s line by line scroll amount.
//   - [NSScrollView.SetLineScroll]
//   - [NSScrollView.HorizontalLineScroll]: The scroll view’s horizontal line by line scroll amount.
//   - [NSScrollView.SetHorizontalLineScroll]
//   - [NSScrollView.VerticalLineScroll]: The scroll view’s vertical line by line scroll amount.
//   - [NSScrollView.SetVerticalLineScroll]
//   - [NSScrollView.PageScroll]: The amount of the document view kept visible when scrolling page by page.
//   - [NSScrollView.SetPageScroll]
//   - [NSScrollView.HorizontalPageScroll]: The amount of the document view kept visible when scrolling horizontally page by page.
//   - [NSScrollView.SetHorizontalPageScroll]
//   - [NSScrollView.VerticalPageScroll]: The amount of the document view kept visible when scrolling vertically page by page.
//   - [NSScrollView.SetVerticalPageScroll]
//   - [NSScrollView.ScrollsDynamically]: A Boolean that indicates whether the scroll view redraws its document view while scrolling continuously.
//   - [NSScrollView.SetScrollsDynamically]
//
// # Arranging Components
//
//   - [NSScrollView.Tile]: Lays out the components of the receiver: the content view, the scrollers, and the ruler views.
//
// # Find Bar Positioning
//
//   - [NSScrollView.FindBarPosition]: The position of the find bar.
//   - [NSScrollView.SetFindBarPosition]
//
// # Specifying a Document’s Predominant Scrolling Behavior
//
//   - [NSScrollView.UsesPredominantAxisScrolling]: A Boolean that indicates whether the scroll view uses a predominant scrolling axis for content.
//   - [NSScrollView.SetUsesPredominantAxisScrolling]
//
// # Specifying the Scroll View Elasticity
//
//   - [NSScrollView.HorizontalScrollElasticity]: The scroll view’s horizontal scrolling elasticity mode.
//   - [NSScrollView.SetHorizontalScrollElasticity]
//   - [NSScrollView.VerticalScrollElasticity]: The scroll view’s vertical scrolling elasticity mode.
//   - [NSScrollView.SetVerticalScrollElasticity]
//
// # Flashing Overlay Scroll Bars
//
//   - [NSScrollView.FlashScrollers]: Flash the overlay scroll bars.
//
// # Zooming the Scroll View
//
//   - [NSScrollView.AllowsMagnification]: Allows the user to magnify the scroll view.
//   - [NSScrollView.SetAllowsMagnification]
//   - [NSScrollView.Magnification]: The amount by which the content is currently scaled.
//   - [NSScrollView.SetMagnification]
//   - [NSScrollView.MagnifyToFitRect]: Magnifies the content view proportionally such that the given rectangle fits centered in the scroll view.
//   - [NSScrollView.MaxMagnification]: The maximum value to which the content can be magnified.
//   - [NSScrollView.SetMaxMagnification]
//   - [NSScrollView.MinMagnification]: The minimum value to which the content can be magnified.
//   - [NSScrollView.SetMinMagnification]
//   - [NSScrollView.SetMagnificationCenteredAtPoint]: Magnify the content by the given amount and center the result on the given point.
//
// See: https://developer.apple.com/documentation/AppKit/NSScrollView
type NSScrollView struct {
	NSView
}

// NSScrollViewFromID constructs a [NSScrollView] from an objc.ID.
//
// A view that displays a portion of a document view and provides scroll bars
// that allow the user to move the document view within the scroll view.
func NSScrollViewFromID(id objc.ID) NSScrollView {
	return NSScrollView{NSView: NSViewFromID(id)}
}
// NOTE: NSScrollView adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSScrollView] class.
//
// # Determining Component Sizes
//
//   - [INSScrollView.ContentSize]: The size of the scroll view’s content view.
//   - [INSScrollView.DocumentVisibleRect]: The portion of the document view, in its own coordinate system, visible through the scroll view’s content view.
//
// # Managing Graphics Attributes
//
//   - [INSScrollView.BackgroundColor]: The color of the content view’s background.
//   - [INSScrollView.SetBackgroundColor]
//   - [INSScrollView.DrawsBackground]: A Boolean that indicates whether the scroll view draws its background.
//   - [INSScrollView.SetDrawsBackground]
//   - [INSScrollView.BorderType]: A value that specifies the appearance of the scroll view’s border.
//   - [INSScrollView.SetBorderType]
//   - [INSScrollView.DocumentCursor]: The content view’s document cursor.
//   - [INSScrollView.SetDocumentCursor]
//
// # Managing the Views
//
//   - [INSScrollView.DocumentView]: The view the scroll view scrolls within its content view.
//   - [INSScrollView.SetDocumentView]
//   - [INSScrollView.AddFloatingSubviewForAxis]: Adds a floating subview to the document view.
//
// # Managing Scrollers
//
//   - [INSScrollView.HorizontalScroller]: The scroll view’s horizontal scroller.
//   - [INSScrollView.SetHorizontalScroller]
//   - [INSScrollView.HasHorizontalScroller]: A Boolean that indicates whether the scroll view has a horizontal scroller.
//   - [INSScrollView.SetHasHorizontalScroller]
//   - [INSScrollView.VerticalScroller]: The scroll view’s vertical scroller.
//   - [INSScrollView.SetVerticalScroller]
//   - [INSScrollView.HasVerticalScroller]: A Boolean that indicates whether the scroll view has a vertical scroller.
//   - [INSScrollView.SetHasVerticalScroller]
//   - [INSScrollView.AutohidesScrollers]: A Boolean that indicates whether the scroll view automatically hides its scroll bars when they are not needed.
//   - [INSScrollView.SetAutohidesScrollers]
//
// # Managing Rulers
//
//   - [INSScrollView.HasHorizontalRuler]: A Boolean that indicates whether the scroll view keeps a horizontal ruler object.
//   - [INSScrollView.SetHasHorizontalRuler]
//   - [INSScrollView.HorizontalRulerView]: The scroll view’s horizontal ruler view.
//   - [INSScrollView.SetHorizontalRulerView]
//   - [INSScrollView.HasVerticalRuler]: A Boolean that indicates whether the scroll view keeps a vertical ruler object.
//   - [INSScrollView.SetHasVerticalRuler]
//   - [INSScrollView.VerticalRulerView]: The scroll view’s vertical ruler view.
//   - [INSScrollView.SetVerticalRulerView]
//   - [INSScrollView.RulersVisible]: A Boolean that indicates whether the scroll view displays its rulers.
//   - [INSScrollView.SetRulersVisible]
//
// # Managing Insets
//
//   - [INSScrollView.AutomaticallyAdjustsContentInsets]: A Boolean that indicates whether the scroll view automatically adjusts its content insets.
//   - [INSScrollView.SetAutomaticallyAdjustsContentInsets]
//   - [INSScrollView.ContentInsets]: The distance that the scroll view’s subviews are inset from the enclosing scroll view during tiling.
//   - [INSScrollView.SetContentInsets]
//   - [INSScrollView.ScrollerInsets]: The distance the scrollers are inset from the edge of the scroll view.
//   - [INSScrollView.SetScrollerInsets]
//
// # Scroller Style
//
//   - [INSScrollView.ScrollerKnobStyle]: The knob style of scroll views that use the overlay scroller style.
//   - [INSScrollView.SetScrollerKnobStyle]
//   - [INSScrollView.ScrollerStyle]: The scroller style used by the scroll view.
//   - [INSScrollView.SetScrollerStyle]
//
// # Setting Scrolling Behavior
//
//   - [INSScrollView.LineScroll]: The scroll view’s line by line scroll amount.
//   - [INSScrollView.SetLineScroll]
//   - [INSScrollView.HorizontalLineScroll]: The scroll view’s horizontal line by line scroll amount.
//   - [INSScrollView.SetHorizontalLineScroll]
//   - [INSScrollView.VerticalLineScroll]: The scroll view’s vertical line by line scroll amount.
//   - [INSScrollView.SetVerticalLineScroll]
//   - [INSScrollView.PageScroll]: The amount of the document view kept visible when scrolling page by page.
//   - [INSScrollView.SetPageScroll]
//   - [INSScrollView.HorizontalPageScroll]: The amount of the document view kept visible when scrolling horizontally page by page.
//   - [INSScrollView.SetHorizontalPageScroll]
//   - [INSScrollView.VerticalPageScroll]: The amount of the document view kept visible when scrolling vertically page by page.
//   - [INSScrollView.SetVerticalPageScroll]
//   - [INSScrollView.ScrollsDynamically]: A Boolean that indicates whether the scroll view redraws its document view while scrolling continuously.
//   - [INSScrollView.SetScrollsDynamically]
//
// # Arranging Components
//
//   - [INSScrollView.Tile]: Lays out the components of the receiver: the content view, the scrollers, and the ruler views.
//
// # Find Bar Positioning
//
//   - [INSScrollView.FindBarPosition]: The position of the find bar.
//   - [INSScrollView.SetFindBarPosition]
//
// # Specifying a Document’s Predominant Scrolling Behavior
//
//   - [INSScrollView.UsesPredominantAxisScrolling]: A Boolean that indicates whether the scroll view uses a predominant scrolling axis for content.
//   - [INSScrollView.SetUsesPredominantAxisScrolling]
//
// # Specifying the Scroll View Elasticity
//
//   - [INSScrollView.HorizontalScrollElasticity]: The scroll view’s horizontal scrolling elasticity mode.
//   - [INSScrollView.SetHorizontalScrollElasticity]
//   - [INSScrollView.VerticalScrollElasticity]: The scroll view’s vertical scrolling elasticity mode.
//   - [INSScrollView.SetVerticalScrollElasticity]
//
// # Flashing Overlay Scroll Bars
//
//   - [INSScrollView.FlashScrollers]: Flash the overlay scroll bars.
//
// # Zooming the Scroll View
//
//   - [INSScrollView.AllowsMagnification]: Allows the user to magnify the scroll view.
//   - [INSScrollView.SetAllowsMagnification]
//   - [INSScrollView.Magnification]: The amount by which the content is currently scaled.
//   - [INSScrollView.SetMagnification]
//   - [INSScrollView.MagnifyToFitRect]: Magnifies the content view proportionally such that the given rectangle fits centered in the scroll view.
//   - [INSScrollView.MaxMagnification]: The maximum value to which the content can be magnified.
//   - [INSScrollView.SetMaxMagnification]
//   - [INSScrollView.MinMagnification]: The minimum value to which the content can be magnified.
//   - [INSScrollView.SetMinMagnification]
//   - [INSScrollView.SetMagnificationCenteredAtPoint]: Magnify the content by the given amount and center the result on the given point.
//
// See: https://developer.apple.com/documentation/AppKit/NSScrollView
type INSScrollView interface {
	INSView

	// Topic: Determining Component Sizes

	// The size of the scroll view’s content view.
	ContentSize() corefoundation.CGSize
	// The portion of the document view, in its own coordinate system, visible through the scroll view’s content view.
	DocumentVisibleRect() corefoundation.CGRect

	// Topic: Managing Graphics Attributes

	// The color of the content view’s background.
	BackgroundColor() INSColor
	SetBackgroundColor(value INSColor)
	// A Boolean that indicates whether the scroll view draws its background.
	DrawsBackground() bool
	SetDrawsBackground(value bool)
	// A value that specifies the appearance of the scroll view’s border.
	BorderType() NSBorderType
	SetBorderType(value NSBorderType)
	// The content view’s document cursor.
	DocumentCursor() INSCursor
	SetDocumentCursor(value INSCursor)

	// Topic: Managing the Views

	// The view the scroll view scrolls within its content view.
	DocumentView() INSView
	SetDocumentView(value INSView)
	// Adds a floating subview to the document view.
	AddFloatingSubviewForAxis(view INSView, axis NSEventGestureAxis)

	// Topic: Managing Scrollers

	// The scroll view’s horizontal scroller.
	HorizontalScroller() INSScroller
	SetHorizontalScroller(value INSScroller)
	// A Boolean that indicates whether the scroll view has a horizontal scroller.
	HasHorizontalScroller() bool
	SetHasHorizontalScroller(value bool)
	// The scroll view’s vertical scroller.
	VerticalScroller() INSScroller
	SetVerticalScroller(value INSScroller)
	// A Boolean that indicates whether the scroll view has a vertical scroller.
	HasVerticalScroller() bool
	SetHasVerticalScroller(value bool)
	// A Boolean that indicates whether the scroll view automatically hides its scroll bars when they are not needed.
	AutohidesScrollers() bool
	SetAutohidesScrollers(value bool)

	// Topic: Managing Rulers

	// A Boolean that indicates whether the scroll view keeps a horizontal ruler object.
	HasHorizontalRuler() bool
	SetHasHorizontalRuler(value bool)
	// The scroll view’s horizontal ruler view.
	HorizontalRulerView() INSRulerView
	SetHorizontalRulerView(value INSRulerView)
	// A Boolean that indicates whether the scroll view keeps a vertical ruler object.
	HasVerticalRuler() bool
	SetHasVerticalRuler(value bool)
	// The scroll view’s vertical ruler view.
	VerticalRulerView() INSRulerView
	SetVerticalRulerView(value INSRulerView)
	// A Boolean that indicates whether the scroll view displays its rulers.
	RulersVisible() bool
	SetRulersVisible(value bool)

	// Topic: Managing Insets

	// A Boolean that indicates whether the scroll view automatically adjusts its content insets.
	AutomaticallyAdjustsContentInsets() bool
	SetAutomaticallyAdjustsContentInsets(value bool)
	// The distance that the scroll view’s subviews are inset from the enclosing scroll view during tiling.
	ContentInsets() foundation.NSEdgeInsets
	SetContentInsets(value foundation.NSEdgeInsets)
	// The distance the scrollers are inset from the edge of the scroll view.
	ScrollerInsets() foundation.NSEdgeInsets
	SetScrollerInsets(value foundation.NSEdgeInsets)

	// Topic: Scroller Style

	// The knob style of scroll views that use the overlay scroller style.
	ScrollerKnobStyle() NSScrollerKnobStyle
	SetScrollerKnobStyle(value NSScrollerKnobStyle)
	// The scroller style used by the scroll view.
	ScrollerStyle() NSScrollerStyle
	SetScrollerStyle(value NSScrollerStyle)

	// Topic: Setting Scrolling Behavior

	// The scroll view’s line by line scroll amount.
	LineScroll() float64
	SetLineScroll(value float64)
	// The scroll view’s horizontal line by line scroll amount.
	HorizontalLineScroll() float64
	SetHorizontalLineScroll(value float64)
	// The scroll view’s vertical line by line scroll amount.
	VerticalLineScroll() float64
	SetVerticalLineScroll(value float64)
	// The amount of the document view kept visible when scrolling page by page.
	PageScroll() float64
	SetPageScroll(value float64)
	// The amount of the document view kept visible when scrolling horizontally page by page.
	HorizontalPageScroll() float64
	SetHorizontalPageScroll(value float64)
	// The amount of the document view kept visible when scrolling vertically page by page.
	VerticalPageScroll() float64
	SetVerticalPageScroll(value float64)
	// A Boolean that indicates whether the scroll view redraws its document view while scrolling continuously.
	ScrollsDynamically() bool
	SetScrollsDynamically(value bool)

	// Topic: Arranging Components

	// Lays out the components of the receiver: the content view, the scrollers, and the ruler views.
	Tile()

	// Topic: Find Bar Positioning

	// The position of the find bar.
	FindBarPosition() NSScrollViewFindBarPosition
	SetFindBarPosition(value NSScrollViewFindBarPosition)

	// Topic: Specifying a Document’s Predominant Scrolling Behavior

	// A Boolean that indicates whether the scroll view uses a predominant scrolling axis for content.
	UsesPredominantAxisScrolling() bool
	SetUsesPredominantAxisScrolling(value bool)

	// Topic: Specifying the Scroll View Elasticity

	// The scroll view’s horizontal scrolling elasticity mode.
	HorizontalScrollElasticity() NSScrollElasticity
	SetHorizontalScrollElasticity(value NSScrollElasticity)
	// The scroll view’s vertical scrolling elasticity mode.
	VerticalScrollElasticity() NSScrollElasticity
	SetVerticalScrollElasticity(value NSScrollElasticity)

	// Topic: Flashing Overlay Scroll Bars

	// Flash the overlay scroll bars.
	FlashScrollers()

	// Topic: Zooming the Scroll View

	// Allows the user to magnify the scroll view.
	AllowsMagnification() bool
	SetAllowsMagnification(value bool)
	// The amount by which the content is currently scaled.
	Magnification() float64
	SetMagnification(value float64)
	// Magnifies the content view proportionally such that the given rectangle fits centered in the scroll view.
	MagnifyToFitRect(rect corefoundation.CGRect)
	// The maximum value to which the content can be magnified.
	MaxMagnification() float64
	SetMaxMagnification(value float64)
	// The minimum value to which the content can be magnified.
	MinMagnification() float64
	SetMinMagnification(value float64)
	// Magnify the content by the given amount and center the result on the given point.
	SetMagnificationCenteredAtPoint(magnification float64, point corefoundation.CGPoint)

	// Notifies the find bar container that the find bar has changed its height.
	FindBarViewDidChangeHeight()
	// Returns whether the container should display its find bar.
	IsFindBarVisible() bool
}

// Init initializes the instance.
func (s NSScrollView) Init() NSScrollView {
	rv := objc.Send[NSScrollView](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s NSScrollView) Autorelease() NSScrollView {
	rv := objc.Send[NSScrollView](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSScrollView creates a new NSScrollView instance.
func NewNSScrollView() NSScrollView {
	class := getNSScrollViewClass()
	rv := objc.Send[NSScrollView](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/AppKit/NSScrollView/init(coder:)
func NewScrollViewWithCoder(coder foundation.INSCoder) NSScrollView {
	instance := getNSScrollViewClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSScrollViewFromID(rv)
}

//
// See: https://developer.apple.com/documentation/AppKit/NSScrollView/init(frame:)
func NewScrollViewWithFrame(frameRect corefoundation.CGRect) NSScrollView {
	instance := getNSScrollViewClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFrame:"), frameRect)
	return NSScrollViewFromID(rv)
}

// Adds a floating subview to the document view.
//
// view: The view that can float.
//
// axis: The event gesture axis on which the view can float. A view can float on
// only one axis at a time.
//
// # Discussion
// 
// Floating subviews of the document view do not scroll like the rest of the
// document. Instead these views appear to float over the document. For
// example, see [NSTableView] floating group rows ([FloatsGroupRows]).
// 
// [NSScrollView] ensures that any scrolling on the non-floating axis is
// performed visually synchronously with the document content.
//
// See: https://developer.apple.com/documentation/AppKit/NSScrollView/addFloatingSubview(_:for:)
func (s NSScrollView) AddFloatingSubviewForAxis(view INSView, axis NSEventGestureAxis) {
	objc.Send[objc.ID](s.ID, objc.Sel("addFloatingSubview:forAxis:"), view, axis)
}
// Lays out the components of the receiver: the content view, the scrollers,
// and the ruler views.
//
// # Discussion
// 
// You rarely need to invoke this method, but subclasses may override it to
// manage additional components.
//
// See: https://developer.apple.com/documentation/AppKit/NSScrollView/tile()
func (s NSScrollView) Tile() {
	objc.Send[objc.ID](s.ID, objc.Sel("tile"))
}
// Flash the overlay scroll bars.
//
// # Discussion
// 
// This method only applies to scroll views that use overlay scrollers.
// 
// This method can be invoked to cause the overlay scroller knobs to be
// momentarily shown. This may be desirable when changing a document view’s
// size or swapping new content into the view, or to give the user a sense of
// the current position within the scrollable range at each step of an
// incremental search or similar operation.
//
// See: https://developer.apple.com/documentation/AppKit/NSScrollView/flashScrollers()
func (s NSScrollView) FlashScrollers() {
	objc.Send[objc.ID](s.ID, objc.Sel("flashScrollers"))
}
// Magnifies the content view proportionally such that the given rectangle
// fits centered in the scroll view.
//
// rect: The rectangle (in content view space) to which the content view is
// magnified.
//
// # Discussion
// 
// The resulting magnification value is clipped to the [MinMagnification] and
// [MaxMagnification] values. To animate the magnification, use the object’s
// animator.
//
// See: https://developer.apple.com/documentation/AppKit/NSScrollView/magnify(toFit:)
func (s NSScrollView) MagnifyToFitRect(rect corefoundation.CGRect) {
	objc.Send[objc.ID](s.ID, objc.Sel("magnifyToFitRect:"), rect)
}
// Magnify the content by the given amount and center the result on the given
// point.
//
// magnification: The amount by which to magnify the content.
//
// point: The point (in content view space) on which to center magnification.
//
// # Discussion
// 
// This method scales the content view such that the passed in point (in
// content view space) remains at the same screen location once the scaling is
// completed. The resulting magnification value is clipped to the
// [MinMagnification] and [MaxMagnification] values. To animate the
// magnification, use the object’s animator.
//
// See: https://developer.apple.com/documentation/AppKit/NSScrollView/setMagnification(_:centeredAt:)
func (s NSScrollView) SetMagnificationCenteredAtPoint(magnification float64, point corefoundation.CGPoint) {
	objc.Send[objc.ID](s.ID, objc.Sel("setMagnification:centeredAtPoint:"), magnification, point)
}
// Notifies the find bar container that the find bar has changed its height.
//
// # Discussion
// 
// Upon receiving this message the container may be required to re-tile its
// contents.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextFinderBarContainer/findBarViewDidChangeHeight()
func (s NSScrollView) FindBarViewDidChangeHeight() {
	objc.Send[objc.ID](s.ID, objc.Sel("findBarViewDidChangeHeight"))
}
// Returns whether the container should display its find bar.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextFinderBarContainer/isFindBarVisible
func (s NSScrollView) IsFindBarVisible() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("isFindBarVisible"))
	return rv
}

// Returns the frame size of a scroll view that contains a content view with
// the specified size.
//
// cSize: The content size.
//
// horizontalScrollerClass: The class used as the horizontal scroller. A value of `nil` specifies that
// no horizontal scroller is used.
//
// verticalScrollerClass: The class used as the vertical scroller. A value of `nil` specifies that no
// horizontal scroller is used.
//
// type: Specifies the appearance of the style of the scroll view’s border. See
// [NSBorderType] for a list of possible values.
// //
// [NSBorderType]: https://developer.apple.com/documentation/AppKit/NSBorderType
//
// controlSize: The control size. The possible values are specified in
// [NSControl.ControlSize]. [NSMiniControlSize] is not supported.
// //
// [NSControl.ControlSize]: https://developer.apple.com/documentation/AppKit/NSControl/ControlSize-swift.enum
// [NSMiniControlSize]: https://developer.apple.com/documentation/AppKit/NSMiniControlSize
//
// scrollerStyle: Specifies the scroll style. See [NSScroller.Style] for supported values.
// //
// [NSScroller.Style]: https://developer.apple.com/documentation/AppKit/NSScroller/Style
//
// # Return Value
// 
// The size of the frame for the specified `contentSize`.
//
// # Discussion
// 
// For an existing scroll view, you can simply use the `frame` method and
// extract its size.
//
// See: https://developer.apple.com/documentation/AppKit/NSScrollView/frameSize(forContentSize:horizontalScrollerClass:verticalScrollerClass:borderType:controlSize:scrollerStyle:)
func (_NSScrollViewClass NSScrollViewClass) FrameSizeForContentSizeHorizontalScrollerClassVerticalScrollerClassBorderTypeControlSizeScrollerStyle(cSize corefoundation.CGSize, horizontalScrollerClass objc.Class, verticalScrollerClass objc.Class, type_ NSBorderType, controlSize NSControlSize, scrollerStyle NSScrollerStyle) corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](objc.ID(_NSScrollViewClass.class), objc.Sel("frameSizeForContentSize:horizontalScrollerClass:verticalScrollerClass:borderType:controlSize:scrollerStyle:"), cSize, horizontalScrollerClass, verticalScrollerClass, type_, controlSize, scrollerStyle)
	return corefoundation.CGSize(rv)
}
// Returns the content size calculated from the frame size and the specified
// specifications.
//
// fSize: The frame size in screen coordinates.
//
// horizontalScrollerClass: The class used as the horizontal scroller. A value of `nil` specifies that
// no horizontal scroller is used.
//
// verticalScrollerClass: The class used as the vertical scroller. A value of `nil` specifies that no
// horizontal scroller is used.
//
// type: Specifies the appearance of the style of the scroll view’s border. See
// [NSBorderType] for a list of possible values.
// //
// [NSBorderType]: https://developer.apple.com/documentation/AppKit/NSBorderType
//
// controlSize: The control size. The possible values are specified in
// [NSControl.ControlSize]. [NSMiniControlSize] is not supported.
// //
// [NSControl.ControlSize]: https://developer.apple.com/documentation/AppKit/NSControl/ControlSize-swift.enum
// [NSMiniControlSize]: https://developer.apple.com/documentation/AppKit/NSMiniControlSize
//
// scrollerStyle: Specifies the scroll style. See [NSScroller.Style] for supported values.
// //
// [NSScroller.Style]: https://developer.apple.com/documentation/AppKit/NSScroller/Style
//
// # Return Value
// 
// The content view frame size.
//
// # Discussion
// 
// For an existing scroll view, you can simply use the [ContentSize] property.
//
// See: https://developer.apple.com/documentation/AppKit/NSScrollView/contentSize(forFrameSize:horizontalScrollerClass:verticalScrollerClass:borderType:controlSize:scrollerStyle:)
func (_NSScrollViewClass NSScrollViewClass) ContentSizeForFrameSizeHorizontalScrollerClassVerticalScrollerClassBorderTypeControlSizeScrollerStyle(fSize corefoundation.CGSize, horizontalScrollerClass objc.Class, verticalScrollerClass objc.Class, type_ NSBorderType, controlSize NSControlSize, scrollerStyle NSScrollerStyle) corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](objc.ID(_NSScrollViewClass.class), objc.Sel("contentSizeForFrameSize:horizontalScrollerClass:verticalScrollerClass:borderType:controlSize:scrollerStyle:"), fSize, horizontalScrollerClass, verticalScrollerClass, type_, controlSize, scrollerStyle)
	return corefoundation.CGSize(rv)
}

// The size of the scroll view’s content view.
//
// See: https://developer.apple.com/documentation/AppKit/NSScrollView/contentSize
func (s NSScrollView) ContentSize() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](s.ID, objc.Sel("contentSize"))
	return corefoundation.CGSize(rv)
}
// The portion of the document view, in its own coordinate system, visible
// through the scroll view’s content view.
//
// See: https://developer.apple.com/documentation/AppKit/NSScrollView/documentVisibleRect
func (s NSScrollView) DocumentVisibleRect() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](s.ID, objc.Sel("documentVisibleRect"))
	return corefoundation.CGRect(rv)
}
// The color of the content view’s background.
//
// # Discussion
// 
// This color is used to paint areas inside the content view that aren’t
// covered by the document view.
//
// See: https://developer.apple.com/documentation/AppKit/NSScrollView/backgroundColor
func (s NSScrollView) BackgroundColor() INSColor {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("backgroundColor"))
	return NSColorFromID(objc.ID(rv))
}
func (s NSScrollView) SetBackgroundColor(value INSColor) {
	objc.Send[struct{}](s.ID, objc.Sel("setBackgroundColor:"), value)
}
// A Boolean that indicates whether the scroll view draws its background.
//
// # Discussion
// 
// When the value of this property is [true], the scroll view cell fills the
// background with its background color.
// 
// If the scroll view encloses an [NSClipView], setting this property to
// [false] also sets the [NSClipView] property [copiesOnScroll] to [false].
// The side effect of setting `drawsBackground` directly on the [NSClipView]
// instead is the appearance of “trails” (vestiges of previous drawing) in
// the document view as it is scrolled.
//
// [copiesOnScroll]: https://developer.apple.com/documentation/AppKit/NSClipView/copiesOnScroll
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSScrollView/drawsBackground
func (s NSScrollView) DrawsBackground() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("drawsBackground"))
	return rv
}
func (s NSScrollView) SetDrawsBackground(value bool) {
	objc.Send[struct{}](s.ID, objc.Sel("setDrawsBackground:"), value)
}
// A value that specifies the appearance of the scroll view’s border.
//
// # Discussion
// 
// See [NSBorderType] for a list of possible values.
//
// [NSBorderType]: https://developer.apple.com/documentation/AppKit/NSBorderType
//
// See: https://developer.apple.com/documentation/AppKit/NSScrollView/borderType
func (s NSScrollView) BorderType() NSBorderType {
	rv := objc.Send[NSBorderType](s.ID, objc.Sel("borderType"))
	return NSBorderType(rv)
}
func (s NSScrollView) SetBorderType(value NSBorderType) {
	objc.Send[struct{}](s.ID, objc.Sel("setBorderType:"), value)
}
// The content view’s document cursor.
//
// See: https://developer.apple.com/documentation/AppKit/NSScrollView/documentCursor
func (s NSScrollView) DocumentCursor() INSCursor {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("documentCursor"))
	return NSCursorFromID(objc.ID(rv))
}
func (s NSScrollView) SetDocumentCursor(value INSCursor) {
	objc.Send[struct{}](s.ID, objc.Sel("setDocumentCursor:"), value)
}
// The scroll view’s content view, the view that clips the document view.
//
// # Discussion
// 
// Setting the value of this property to an [NSClipView] that has a document
// view also sets the scroll view’s document view to be the document view of
// that [NSClipView]. The original content view retains its document view.
//
// See: https://developer.apple.com/documentation/AppKit/NSScrollView/contentView
func (s NSScrollView) ContentView() INSClipView {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("contentView"))
	return NSClipViewFromID(objc.ID(rv))
}
func (s NSScrollView) SetContentView(value INSClipView) {
	objc.Send[struct{}](s.ID, objc.Sel("setContentView:"), value)
}
// The view the scroll view scrolls within its content view.
//
// See: https://developer.apple.com/documentation/AppKit/NSScrollView/documentView
func (s NSScrollView) DocumentView() INSView {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("documentView"))
	return NSViewFromID(objc.ID(rv))
}
func (s NSScrollView) SetDocumentView(value INSView) {
	objc.Send[struct{}](s.ID, objc.Sel("setDocumentView:"), value)
}
// The scroll view’s horizontal scroller.
//
// # Discussion
// 
// The value of this property is `nil` if the scroll view has no horizontal
// scroller.
// 
// You can access the horizontal scroller using this property even if the
// scroll view isn’t currently displaying it. To make sure the scroller is
// visible, set [HasHorizontalScroller] to [true].
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSScrollView/horizontalScroller
func (s NSScrollView) HorizontalScroller() INSScroller {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("horizontalScroller"))
	return NSScrollerFromID(objc.ID(rv))
}
func (s NSScrollView) SetHorizontalScroller(value INSScroller) {
	objc.Send[struct{}](s.ID, objc.Sel("setHorizontalScroller:"), value)
}
// A Boolean that indicates whether the scroll view has a horizontal scroller.
//
// # Discussion
// 
// When the value of this property is [true], the scroll view allocates and
// displays a horizontal scroller as needed. The default value of this
// property is [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSScrollView/hasHorizontalScroller
func (s NSScrollView) HasHorizontalScroller() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("hasHorizontalScroller"))
	return rv
}
func (s NSScrollView) SetHasHorizontalScroller(value bool) {
	objc.Send[struct{}](s.ID, objc.Sel("setHasHorizontalScroller:"), value)
}
// The scroll view’s vertical scroller.
//
// # Discussion
// 
// The value of this property is `nil` if the scroll view has no vertical
// scroller.
// 
// You can access the vertical scroller using this property even if the scroll
// view isn’t currently displaying it. To make sure the scroller is visible,
// set [HasVerticalScroller] to [true].
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSScrollView/verticalScroller
func (s NSScrollView) VerticalScroller() INSScroller {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("verticalScroller"))
	return NSScrollerFromID(objc.ID(rv))
}
func (s NSScrollView) SetVerticalScroller(value INSScroller) {
	objc.Send[struct{}](s.ID, objc.Sel("setVerticalScroller:"), value)
}
// A Boolean that indicates whether the scroll view has a vertical scroller.
//
// # Discussion
// 
// When the value of this property is [true], the scroll view allocates and
// displays a vertical scroller as needed. The default value of this property
// is [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSScrollView/hasVerticalScroller
func (s NSScrollView) HasVerticalScroller() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("hasVerticalScroller"))
	return rv
}
func (s NSScrollView) SetHasVerticalScroller(value bool) {
	objc.Send[struct{}](s.ID, objc.Sel("setHasVerticalScroller:"), value)
}
// A Boolean that indicates whether the scroll view automatically hides its
// scroll bars when they are not needed.
//
// # Discussion
// 
// The horizontal and vertical scroll bars are hidden independently of each
// other. When the value of this property is [true] and the content of the
// scroll view doesn’t extend beyond the size of the clip view on a given
// axis, the scroller on that axis is removed to leave more room for the
// content.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSScrollView/autohidesScrollers
func (s NSScrollView) AutohidesScrollers() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("autohidesScrollers"))
	return rv
}
func (s NSScrollView) SetAutohidesScrollers(value bool) {
	objc.Send[struct{}](s.ID, objc.Sel("setAutohidesScrollers:"), value)
}
// A Boolean that indicates whether the scroll view keeps a horizontal ruler
// object.
//
// # Discussion
// 
// When the value of this method is [true], the scroll view allocates a
// horizontal ruler the first time it’s needed.
// 
// Display of rulers is controlled using the [RulersVisible] property.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSScrollView/hasHorizontalRuler
func (s NSScrollView) HasHorizontalRuler() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("hasHorizontalRuler"))
	return rv
}
func (s NSScrollView) SetHasHorizontalRuler(value bool) {
	objc.Send[struct{}](s.ID, objc.Sel("setHasHorizontalRuler:"), value)
}
// The scroll view’s horizontal ruler view.
//
// # Discussion
// 
// The value of this property is `nil` when the scroll view has no horizontal
// ruler view.
// 
// If the scroll view is set to display a horizontal ruler view and doesn’t
// yet have one, this property creates an instance of the ruler view class set
// using the class method `setRulerViewClass(_:)`. You can use this property
// to override the default ruler class set using the class method
// `setRulerViewClass(_:)`.
// 
// Display of rulers is controlled using the [RulersVisible] property.
//
// See: https://developer.apple.com/documentation/AppKit/NSScrollView/horizontalRulerView
func (s NSScrollView) HorizontalRulerView() INSRulerView {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("horizontalRulerView"))
	return NSRulerViewFromID(objc.ID(rv))
}
func (s NSScrollView) SetHorizontalRulerView(value INSRulerView) {
	objc.Send[struct{}](s.ID, objc.Sel("setHorizontalRulerView:"), value)
}
// A Boolean that indicates whether the scroll view keeps a vertical ruler
// object.
//
// # Discussion
// 
// When the value of this method is [true], the scroll view allocates a
// vertical ruler the first time it’s needed.
// 
// Display of rulers is controlled using the [RulersVisible] property.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSScrollView/hasVerticalRuler
func (s NSScrollView) HasVerticalRuler() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("hasVerticalRuler"))
	return rv
}
func (s NSScrollView) SetHasVerticalRuler(value bool) {
	objc.Send[struct{}](s.ID, objc.Sel("setHasVerticalRuler:"), value)
}
// The scroll view’s vertical ruler view.
//
// # Discussion
// 
// The value of this property is `nil` when the scroll view has no vertical
// ruler view.
// 
// If the scroll view is set to display a vertical ruler view and doesn’t
// yet have one, this property creates an instance of the ruler view class set
// using the class method `setRulerViewClass(_:)`. You can use this property
// to override the default ruler class set using the class method
// `setRulerViewClass(_:)`.
// 
// Display of rulers is controlled using the [RulersVisible] property.
//
// See: https://developer.apple.com/documentation/AppKit/NSScrollView/verticalRulerView
func (s NSScrollView) VerticalRulerView() INSRulerView {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("verticalRulerView"))
	return NSRulerViewFromID(objc.ID(rv))
}
func (s NSScrollView) SetVerticalRulerView(value INSRulerView) {
	objc.Send[struct{}](s.ID, objc.Sel("setVerticalRulerView:"), value)
}
// A Boolean that indicates whether the scroll view displays its rulers.
//
// # Discussion
// 
// When the value of this property is [true], the scroll view displays its
// rulers (creating them if needed). When the value of this property is
// [false], the scroll view doesn’t display its rulers.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSScrollView/rulersVisible
func (s NSScrollView) RulersVisible() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("rulersVisible"))
	return rv
}
func (s NSScrollView) SetRulersVisible(value bool) {
	objc.Send[struct{}](s.ID, objc.Sel("setRulersVisible:"), value)
}
// A Boolean that indicates whether the scroll view automatically adjusts its
// content insets.
//
// # Discussion
// 
// When the value of this property is [true], the scroll view automatically
// sets its [ContentInsets] property to account for any overlapping title or
// tool bar. To overlap with the title or tool bar, the window style mask must
// include [NSFullSizeContentViewWindowMask] and the title bar must not be
// transparent.
// 
// The default value of this property is [true].
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSScrollView/automaticallyAdjustsContentInsets
func (s NSScrollView) AutomaticallyAdjustsContentInsets() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("automaticallyAdjustsContentInsets"))
	return rv
}
func (s NSScrollView) SetAutomaticallyAdjustsContentInsets(value bool) {
	objc.Send[struct{}](s.ID, objc.Sel("setAutomaticallyAdjustsContentInsets:"), value)
}
// The distance that the scroll view’s subviews are inset from the enclosing
// scroll view during tiling.
//
// # Discussion
// 
// When the value of this property is equal to [NSEdgeInsetsZero], traditional
// tiling is performed. Rulers, headers, and other subviews are tiled with the
// [ContentView] frame filling the remaining space. When the value of this
// property is not equal to [NSEdgeInsetsZero], the rulers, headers, and other
// subviews are inset as specified. The [ContentView] is placed underneath
// these sibling views and is only inset by the scroll view border and
// non-overlay scrollers.
// 
// See [NSEdgeInsets] for possible values.
// 
// When the value of the [AutomaticallyAdjustsContentInsets] property is
// [true], any value of this property is overridden during tiling.
//
// [NSEdgeInsets]: https://developer.apple.com/documentation/Foundation/NSEdgeInsets
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSScrollView/contentInsets
func (s NSScrollView) ContentInsets() foundation.NSEdgeInsets {
	rv := objc.Send[foundation.NSEdgeInsets](s.ID, objc.Sel("contentInsets"))
	return foundation.NSEdgeInsets(rv)
}
func (s NSScrollView) SetContentInsets(value foundation.NSEdgeInsets) {
	objc.Send[struct{}](s.ID, objc.Sel("setContentInsets:"), value)
}
// The distance the scrollers are inset from the edge of the scroll view.
//
// See: https://developer.apple.com/documentation/AppKit/NSScrollView/scrollerInsets
func (s NSScrollView) ScrollerInsets() foundation.NSEdgeInsets {
	rv := objc.Send[foundation.NSEdgeInsets](s.ID, objc.Sel("scrollerInsets"))
	return foundation.NSEdgeInsets(rv)
}
func (s NSScrollView) SetScrollerInsets(value foundation.NSEdgeInsets) {
	objc.Send[struct{}](s.ID, objc.Sel("setScrollerInsets:"), value)
}
// The knob style of scroll views that use the overlay scroller style.
//
// # Discussion
// 
// See [NSScroller.KnobStyle] for possible values.
// 
// Applicable only to scroll views that use overlay scrollers.
//
// [NSScroller.KnobStyle]: https://developer.apple.com/documentation/AppKit/NSScroller/KnobStyle-swift.enum
//
// See: https://developer.apple.com/documentation/AppKit/NSScrollView/scrollerKnobStyle
func (s NSScrollView) ScrollerKnobStyle() NSScrollerKnobStyle {
	rv := objc.Send[NSScrollerKnobStyle](s.ID, objc.Sel("scrollerKnobStyle"))
	return NSScrollerKnobStyle(rv)
}
func (s NSScrollView) SetScrollerKnobStyle(value NSScrollerKnobStyle) {
	objc.Send[struct{}](s.ID, objc.Sel("setScrollerKnobStyle:"), value)
}
// The scroller style used by the scroll view.
//
// # Discussion
// 
// See [NSScroller.Style] for possible values.
// 
// This setting is automatically set at runtime, based on the user’s
// preference setting and, if relevant, the set of connected pointing devices
// and their configured scroll capabilities, as determined by the [NSScroller]
// method [PreferredScrollerStyle].
// 
// Setting an scroll view’s scroller style sets the style of both the
// horizontal and vertical scrollers. If the scroll view subsequently creates
// or is assigned a new horizontal or vertical scroller, they are assigned the
// same scroller style assigned to the scroll view.
//
// [NSScroller.Style]: https://developer.apple.com/documentation/AppKit/NSScroller/Style
//
// See: https://developer.apple.com/documentation/AppKit/NSScrollView/scrollerStyle
func (s NSScrollView) ScrollerStyle() NSScrollerStyle {
	rv := objc.Send[NSScrollerStyle](s.ID, objc.Sel("scrollerStyle"))
	return NSScrollerStyle(rv)
}
func (s NSScrollView) SetScrollerStyle(value NSScrollerStyle) {
	objc.Send[struct{}](s.ID, objc.Sel("setScrollerStyle:"), value)
}
// The scroll view’s line by line scroll amount.
//
// # Discussion
// 
// The value of this property is the amount by which the scroll view scrolls
// itself when scrolling line by line, expressed in the content view’s
// coordinate system. This value is used when the user clicks the scroll
// arrows without holding down a modifier key. When displaying text in a
// scroll view, for example, you might set this value to the height of a
// single line of text in the default font. As part of its implementation,
// this property accesses [VerticalLineScroll].
// 
// Note that a scroll view can have two different line scroll amounts:
// [VerticalLineScroll] and [HorizontalLineScroll]. Set this property only if
// you can be sure they’re both the same; setting this property sets both
// [VerticalLineScroll] and [HorizontalLineScroll] to the same value.
//
// See: https://developer.apple.com/documentation/AppKit/NSScrollView/lineScroll
func (s NSScrollView) LineScroll() float64 {
	rv := objc.Send[float64](s.ID, objc.Sel("lineScroll"))
	return rv
}
func (s NSScrollView) SetLineScroll(value float64) {
	objc.Send[struct{}](s.ID, objc.Sel("setLineScroll:"), value)
}
// The scroll view’s horizontal line by line scroll amount.
//
// # Discussion
// 
// The value of this property is the amount by which the scroll view scrolls
// itself horizontally when scrolling line by line, expressed in the content
// view’s coordinate system. This value is used when the user clicks the
// scroll arrows on the horizontal scroll bar without holding down a modifier
// key.
//
// See: https://developer.apple.com/documentation/AppKit/NSScrollView/horizontalLineScroll
func (s NSScrollView) HorizontalLineScroll() float64 {
	rv := objc.Send[float64](s.ID, objc.Sel("horizontalLineScroll"))
	return rv
}
func (s NSScrollView) SetHorizontalLineScroll(value float64) {
	objc.Send[struct{}](s.ID, objc.Sel("setHorizontalLineScroll:"), value)
}
// The scroll view’s vertical line by line scroll amount.
//
// # Discussion
// 
// The value of this property is the amount by which the scroll view scrolls
// itself vertically when scrolling line by line, expressed in the content
// view’s coordinate system. This value is used when the user clicks the
// scroll arrows on the vertical scroll bar without holding down a modifier
// key.
//
// See: https://developer.apple.com/documentation/AppKit/NSScrollView/verticalLineScroll
func (s NSScrollView) VerticalLineScroll() float64 {
	rv := objc.Send[float64](s.ID, objc.Sel("verticalLineScroll"))
	return rv
}
func (s NSScrollView) SetVerticalLineScroll(value float64) {
	objc.Send[struct{}](s.ID, objc.Sel("setVerticalLineScroll:"), value)
}
// The amount of the document view kept visible when scrolling page by page.
//
// # Discussion
// 
// The value of this property is the amount of the document view kept visible
// when scrolling page by page, expressed in the content view’s coordinate
// system. This value is used when the user clicks the scroll arrows while
// holding down the Option key. As part of its implementation, this property
// accesses [VerticalPageScroll].
// 
// This amount expresses the context that remains when the scroll view scrolls
// by one page, allowing the user to orient to the new display. It differs
// from the line scroll amount, which indicates how far the document view
// moves. The page scroll amount is the amount common to the content view
// before and after the document view is scrolled by one page. Thus, setting
// the page scroll amount to `0.0` implies that the entire visible portion of
// the document view is replaced when a page scroll occurs.
// 
// Note that a scroll view can have two different page scroll amounts:
// [VerticalPageScroll] and [HorizontalPageScroll]. Set this property only if
// you can be sure they’re both the same; setting this property sets both
// [VerticalPageScroll] and [HorizontalPageScroll] to the same value.
//
// See: https://developer.apple.com/documentation/AppKit/NSScrollView/pageScroll
func (s NSScrollView) PageScroll() float64 {
	rv := objc.Send[float64](s.ID, objc.Sel("pageScroll"))
	return rv
}
func (s NSScrollView) SetPageScroll(value float64) {
	objc.Send[struct{}](s.ID, objc.Sel("setPageScroll:"), value)
}
// The amount of the document view kept visible when scrolling horizontally
// page by page.
//
// # Discussion
// 
// The value of this property is the amount of the document view kept visible
// when scrolling horizontally page by page, expressed in the content view’s
// coordinate system. This value is used when the user clicks the scroll
// arrows on the horizontal scroll bar while holding down the Option key.
// 
// This amount expresses the context that remains when the receiver scrolls by
// one page, allowing the user to orient to the new display. It differs from
// the line scroll amount, which indicates how far the document view moves.
// The page scroll amount is the amount common to the content view before and
// after the document view is scrolled by one page. Thus, setting the page
// scroll amount to `0.0` implies that the entire visible portion of the
// document view is replaced when a page scroll occurs.
//
// See: https://developer.apple.com/documentation/AppKit/NSScrollView/horizontalPageScroll
func (s NSScrollView) HorizontalPageScroll() float64 {
	rv := objc.Send[float64](s.ID, objc.Sel("horizontalPageScroll"))
	return rv
}
func (s NSScrollView) SetHorizontalPageScroll(value float64) {
	objc.Send[struct{}](s.ID, objc.Sel("setHorizontalPageScroll:"), value)
}
// The amount of the document view kept visible when scrolling vertically page
// by page.
//
// # Discussion
// 
// The value of this property is the amount of the document view kept visible
// when scrolling vertically page by page, expressed in the content view’s
// coordinate system. This value is used when the user clicks the scroll
// arrows on the vertical scroll bar while holding down the Option key.
// 
// This amount expresses the context that remains when the receiver scrolls by
// one page, allowing the user to orient to the new display. It differs from
// the line scroll amount, which indicates how far the document view moves.
// The page scroll amount is the amount common to the content view before and
// after the document view is scrolled by one page. Thus, setting the page
// scroll amount to `0.0` implies that the entire visible portion of the
// document view is replaced when a page scroll occurs.
//
// See: https://developer.apple.com/documentation/AppKit/NSScrollView/verticalPageScroll
func (s NSScrollView) VerticalPageScroll() float64 {
	rv := objc.Send[float64](s.ID, objc.Sel("verticalPageScroll"))
	return rv
}
func (s NSScrollView) SetVerticalPageScroll(value float64) {
	objc.Send[struct{}](s.ID, objc.Sel("setVerticalPageScroll:"), value)
}
// A Boolean that indicates whether the scroll view redraws its document view
// while scrolling continuously.
//
// # Discussion
// 
// When the value of this property is [true], the scroll view redraws its
// document view while scrolling. When the value of this property is[false],
// the scroll view redraws only when the scroller knob is released. The
// default value of this property is [true].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSScrollView/scrollsDynamically
func (s NSScrollView) ScrollsDynamically() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("scrollsDynamically"))
	return rv
}
func (s NSScrollView) SetScrollsDynamically(value bool) {
	objc.Send[struct{}](s.ID, objc.Sel("setScrollsDynamically:"), value)
}
// The position of the find bar.
//
// # Discussion
// 
// See [NSScrollView.FindBarPosition] for possible values.
//
// [NSScrollView.FindBarPosition]: https://developer.apple.com/documentation/AppKit/NSScrollView/FindBarPosition-swift.enum
//
// See: https://developer.apple.com/documentation/AppKit/NSScrollView/findBarPosition-swift.property
func (s NSScrollView) FindBarPosition() NSScrollViewFindBarPosition {
	rv := objc.Send[NSScrollViewFindBarPosition](s.ID, objc.Sel("findBarPosition"))
	return NSScrollViewFindBarPosition(rv)
}
func (s NSScrollView) SetFindBarPosition(value NSScrollViewFindBarPosition) {
	objc.Send[struct{}](s.ID, objc.Sel("setFindBarPosition:"), value)
}
// A Boolean that indicates whether the scroll view uses a predominant
// scrolling axis for content.
//
// # Discussion
// 
// Some content is scrollable in both the horizontal and vertical axes, but is
// predominantly scrolled one axis at a time. Other content (such as a drawing
// canvas) should scroll freely in both axes.
// 
// Traditionally this is not an issue with scroll wheels since they can only
// scroll in one direction at a time. With scroll balls and touch surfaces, it
// becomes more difficult to determine the user’s intention.
// 
// This property helps a scroll view determine the user’s intention by
// specifying if there is a predominant scrolling axis for content.
// 
// When the value of this property is [true], the scroll view uses a
// predominant scrolling access. The default value of this property is [true].
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSScrollView/usesPredominantAxisScrolling
func (s NSScrollView) UsesPredominantAxisScrolling() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("usesPredominantAxisScrolling"))
	return rv
}
func (s NSScrollView) SetUsesPredominantAxisScrolling(value bool) {
	objc.Send[struct{}](s.ID, objc.Sel("setUsesPredominantAxisScrolling:"), value)
}
// The scroll view’s horizontal scrolling elasticity mode.
//
// # Discussion
// 
// A scroll view can scroll its contents past its bounds to achieve an elastic
// effect.
// 
// When set to [NSScrollView.Elasticity.automatic], scrolling the horizontal
// axis beyond its document bounds only occurs if the document width is
// greater than the view width, or the vertical scroller is hidden and the
// horizontal scroller is visible. The default value is
// [NSScrollView.Elasticity.automatic].
// 
// See [NSScrollView.Elasticity] for possible values.
//
// [NSScrollView.Elasticity.automatic]: https://developer.apple.com/documentation/AppKit/NSScrollView/Elasticity/automatic
// [NSScrollView.Elasticity]: https://developer.apple.com/documentation/AppKit/NSScrollView/Elasticity
//
// See: https://developer.apple.com/documentation/AppKit/NSScrollView/horizontalScrollElasticity
func (s NSScrollView) HorizontalScrollElasticity() NSScrollElasticity {
	rv := objc.Send[NSScrollElasticity](s.ID, objc.Sel("horizontalScrollElasticity"))
	return NSScrollElasticity(rv)
}
func (s NSScrollView) SetHorizontalScrollElasticity(value NSScrollElasticity) {
	objc.Send[struct{}](s.ID, objc.Sel("setHorizontalScrollElasticity:"), value)
}
// The scroll view’s vertical scrolling elasticity mode.
//
// # Discussion
// 
// A scroll view can scroll its contents past its bounds to achieve an elastic
// effect.
// 
// When set to [NSScrollView.Elasticity.automatic], scrolling the vertical
// axis beyond its document bounds only occurs if any of the following are
// true: the vertical scroller is visible, the content height is greater than
// view height, or the horizontal scroller hidden. The default value is
// [NSScrollView.Elasticity.automatic].
// 
// See [NSScrollView.Elasticity] for possible values.
//
// [NSScrollView.Elasticity.automatic]: https://developer.apple.com/documentation/AppKit/NSScrollView/Elasticity/automatic
// [NSScrollView.Elasticity]: https://developer.apple.com/documentation/AppKit/NSScrollView/Elasticity
//
// See: https://developer.apple.com/documentation/AppKit/NSScrollView/verticalScrollElasticity
func (s NSScrollView) VerticalScrollElasticity() NSScrollElasticity {
	rv := objc.Send[NSScrollElasticity](s.ID, objc.Sel("verticalScrollElasticity"))
	return NSScrollElasticity(rv)
}
func (s NSScrollView) SetVerticalScrollElasticity(value NSScrollElasticity) {
	objc.Send[struct{}](s.ID, objc.Sel("setVerticalScrollElasticity:"), value)
}
// Allows the user to magnify the scroll view.
//
// # Discussion
// 
// This property does not prevent the developer from manually adjusting the
// magnification value. If magnification exceeds either the maximum or minimum
// limits for magnification, and [AllowsMagnification] is [true], the scroll
// view temporarily animates the content magnification just past those limits
// before returning to them. The default value is [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSScrollView/allowsMagnification
func (s NSScrollView) AllowsMagnification() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("allowsMagnification"))
	return rv
}
func (s NSScrollView) SetAllowsMagnification(value bool) {
	objc.Send[struct{}](s.ID, objc.Sel("setAllowsMagnification:"), value)
}
// The amount by which the content is currently scaled.
//
// # Discussion
// 
// To animate the magnification, use the object’s animator. The default
// value is `1.0`.
//
// See: https://developer.apple.com/documentation/AppKit/NSScrollView/magnification
func (s NSScrollView) Magnification() float64 {
	rv := objc.Send[float64](s.ID, objc.Sel("magnification"))
	return rv
}
func (s NSScrollView) SetMagnification(value float64) {
	objc.Send[struct{}](s.ID, objc.Sel("setMagnification:"), value)
}
// The maximum value to which the content can be magnified.
//
// # Discussion
// 
// This value must be greater than or equal to the minimum magnification. The
// default value is `4.0`.
//
// See: https://developer.apple.com/documentation/AppKit/NSScrollView/maxMagnification
func (s NSScrollView) MaxMagnification() float64 {
	rv := objc.Send[float64](s.ID, objc.Sel("maxMagnification"))
	return rv
}
func (s NSScrollView) SetMaxMagnification(value float64) {
	objc.Send[struct{}](s.ID, objc.Sel("setMaxMagnification:"), value)
}
// The minimum value to which the content can be magnified.
//
// # Discussion
// 
// The default value is `0.25`.
//
// See: https://developer.apple.com/documentation/AppKit/NSScrollView/minMagnification
func (s NSScrollView) MinMagnification() float64 {
	rv := objc.Send[float64](s.ID, objc.Sel("minMagnification"))
	return rv
}
func (s NSScrollView) SetMinMagnification(value float64) {
	objc.Send[struct{}](s.ID, objc.Sel("setMinMagnification:"), value)
}
// The view assigned by the text bar as the find bar view for the container.
//
// # Discussion
// 
// This property is managed by [NSTextFinder] and you must not set this
// property.
// 
// The container may freely modify the view’s width, but should not modify
// its height.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextFinderBarContainer/findBarView
func (s NSScrollView) FindBarView() INSView {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("findBarView"))
	return NSViewFromID(objc.ID(rv))
}
func (s NSScrollView) SetFindBarView(value INSView) {
	objc.Send[struct{}](s.ID, objc.Sel("setFindBarView:"), value)
}
// Returns whether the container should display its find bar.
//
// # Discussion
// 
// When this property is [true] and the [FindBarView] property is set, then
// the find bar is displayed by the container. Otherwise, the find bar is not
// displayed.
// 
// The default value should be [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSTextFinderBarContainer/isFindBarVisible
func (s NSScrollView) FindBarVisible() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("isFindBarVisible"))
	return rv
}
func (s NSScrollView) SetFindBarVisible(value bool) {
	objc.Send[struct{}](s.ID, objc.Sel("setFindBarVisible:"), value)
}

// Returns the default class to be used for ruler objects in NSScrollViews.
//
// # Discussion
// 
// This class is normally NSRulerView.
//
// See: https://developer.apple.com/documentation/AppKit/NSScrollView/rulerViewClass
func (_NSScrollViewClass NSScrollViewClass) RulerViewClass() objc.Class {
	rv := objc.Send[objc.Class](objc.ID(_NSScrollViewClass.class), objc.Sel("rulerViewClass"))
	return rv
}
func (_NSScrollViewClass NSScrollViewClass) SetRulerViewClass(value objc.Class) {
	objc.Send[struct{}](objc.ID(_NSScrollViewClass.class), objc.Sel("setRulerViewClass:"), value)
}

