// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"context"
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSView] class.
var (
	_NSViewClass     NSViewClass
	_NSViewClassOnce sync.Once
)

func getNSViewClass() NSViewClass {
	_NSViewClassOnce.Do(func() {
		_NSViewClass = NSViewClass{class: objc.GetClass("NSView")}
	})
	return _NSViewClass
}

// GetNSViewClass returns the class object for NSView.
func GetNSViewClass() NSViewClass {
	return getNSViewClass()
}

type NSViewClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSViewClass) Alloc() NSView {
	rv := objc.Send[NSView](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// The infrastructure for drawing, printing, and handling events in an app.
//
// # Overview
// 
// You typically don’t use [NSView] objects directly. Instead, you use
// objects that descend from [NSView] or you subclass [NSView] yourself and
// override its methods to implement the behavior you need. An instance of the
// [NSView] class (or one of its subclasses) is commonly known as a view
// object, or simply as a view.
// 
// Views handle the presentation and interaction with your app’s visible
// content. You arrange one or more views inside an [NSWindow] object, which
// acts as a wrapper for your content. A view object defines a rectangular
// region for drawing and receiving mouse events. Views handle other chores as
// well, including the dragging of icons and working with the [NSScrollView]
// class to support efficient scrolling.
// 
// AppKit handles most of your app’s [NSView] management. Unless you’re
// implementing a concrete subclass of [NSView] or working intimately with the
// content of the view hierarchy at runtime, you don’t need to know much
// about this class’s interface. For any view, there are many methods that
// you can use as-is. The following methods are commonly used.
// 
// - [NSView.Frame] returns the location and size of the [NSView] object. - [bounds]
// returns the internal origin and size of the [NSView] object. -
// [needsDisplay] determines whether the [NSView] object needs to be redrawn.
// - [window] returns the [NSWindow] object that contains the [NSView] object.
// - [NSView.DrawRect] draws the [NSView] object. (All subclasses must implement this
// method, but it’s rarely invoked explicitly.) An alternative to drawing is
// to update the layer directly using the [NSView.UpdateLayer] method.
// 
// For more information on how [NSView] instances handle event and action
// messages, see [Cocoa Event Handling Guide]. For more information on
// displaying tooltips and contextual menus, see [NSMenu] and [NSWindow].
// 
// # Subclassing notes
// 
// [NSView] is perhaps the most important class in AppKit when it comes to
// subclassing and inheritance. Most user-interface objects you see in a Cocoa
// application are objects that inherit from [NSView]. If you want to create
// an object that draws itself in a special way, or that responds to mouse
// clicks in a special way, you would create a custom subclass of [NSView] (or
// of a class that inherits from [NSView]). Subclassing [NSView] is such a
// common and important procedure that several technical documents describe
// how to both draw in custom subclasses and respond to events in custom
// subclasses. See [Cocoa Drawing Guide] and [Cocoa Event Handling Guide]
// (especially “[Handling Mouse Events]” and “[Mouse Events]”).
// 
// # Handling events in your subclass
// 
// If you subclass [NSView] directly and handle specific types of events,
// don’t call `super` in the implementations of your event-related methods.
// Views inherit their event-handling capabilities from their [NSResponder]
// parent class. The default behavior for responders is to pass events up the
// responder chain, which isn’t the behavior you typically want for a custom
// view. Therefore, don’t call `super` if your view implements any of the
// following methods and handles the event:
// 
// - [MouseDown] - [MouseDragged] - [MouseUp] - [MouseMoved] - [MouseEntered]
// - [MouseExited] - [RightMouseDragged] - [RightMouseUp] - [OtherMouseDown] -
// [OtherMouseDragged] - [OtherMouseUp] - [ScrollWheel] - [KeyDown] - [KeyUp]
// - [FlagsChanged] - [TabletPoint] - [TabletProximity]
// 
// If your view descends from a class other than [NSView], call `super` to let
// the parent view handle any events that you don’t.
//
// [Cocoa Drawing Guide]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/CocoaDrawingGuide/Introduction/Introduction.html#//apple_ref/doc/uid/TP40003290
// [Cocoa Event Handling Guide]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/EventOverview/Introduction/Introduction.html#//apple_ref/doc/uid/10000060i
// [Handling Mouse Events]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/EventOverview/HandlingMouseEvents/HandlingMouseEvents.html#//apple_ref/doc/uid/10000060i-CH6
// [Mouse Events]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/EventOverview/EventObjectsTypes/EventObjectsTypes.html#//apple_ref/doc/uid/10000060i-CH4-SW10
// [bounds]: https://developer.apple.com/documentation/AppKit/NSView/bounds
// [needsDisplay]: https://developer.apple.com/documentation/AppKit/NSView/needsDisplay
// [window]: https://developer.apple.com/documentation/AppKit/NSView/window
//
// # Creating a view object
//
//   - [NSView.InitWithFrame]: Initializes and returns a newly allocated [NSView] object with a specified frame rectangle.
//   - [NSView.PrepareForReuse]: Restores the view to an initial state so that it can be reused.
//
// # Instance Properties
//
//   - [NSView.PrefersCompactControlSizeMetrics]: When this property is true, any NSControls in the view or its descendants will be sized with compact metrics compatible with macOS 15 and earlier. Defaults to false
//   - [NSView.SetPrefersCompactControlSizeMetrics]
//   - [NSView.WritingToolsCoordinator]
//   - [NSView.SetWritingToolsCoordinator]
//
// See: https://developer.apple.com/documentation/AppKit/NSView
type NSView struct {
	NSResponder
}

// NSViewFromID constructs a [NSView] from an objc.ID.
//
// The infrastructure for drawing, printing, and handling events in an app.
func NSViewFromID(id objc.ID) NSView {
	return NSView{NSResponder: NSResponderFromID(id)}
}
// NOTE: NSView adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [NSView] class.
//
// # Creating a view object
//
//   - [INSView.InitWithFrame]: Initializes and returns a newly allocated [NSView] object with a specified frame rectangle.
//   - [INSView.PrepareForReuse]: Restores the view to an initial state so that it can be reused.
//
// # Instance Properties
//
//   - [INSView.PrefersCompactControlSizeMetrics]: When this property is true, any NSControls in the view or its descendants will be sized with compact metrics compatible with macOS 15 and earlier. Defaults to false
//   - [INSView.SetPrefersCompactControlSizeMetrics]
//   - [INSView.WritingToolsCoordinator]
//   - [INSView.SetWritingToolsCoordinator]
//
// See: https://developer.apple.com/documentation/AppKit/NSView
type INSView interface {
	INSResponder
	
	
	NSAppearanceCustomization
	NSDraggingDestination
	NSUserInterfaceItemIdentification

	// Topic: Creating a view object

	// Initializes and returns a newly allocated [NSView] object with a specified frame rectangle.
	InitWithFrame(frameRect corefoundation.CGRect) NSView
	// Restores the view to an initial state so that it can be reused.
	PrepareForReuse()

	// Topic: Instance Properties

	// When this property is true, any NSControls in the view or its descendants will be sized with compact metrics compatible with macOS 15 and earlier. Defaults to false
	PrefersCompactControlSizeMetrics() bool
	SetPrefersCompactControlSizeMetrics(value bool)
	WritingToolsCoordinator() INSWritingToolsCoordinator
	SetWritingToolsCoordinator(value INSWritingToolsCoordinator)

	// The view’s frame rectangle, which defines its position and size in its superview’s coordinate system.
	Frame() corefoundation.CGRect
	SetFrame(value corefoundation.CGRect)
	// The view’s bounds rectangle, which expresses its location and size in its own coordinate system.
	Bounds() corefoundation.CGRect
	SetBounds(value corefoundation.CGRect)
	// A Boolean value that determines whether the view needs to be redrawn before being displayed.
	NeedsDisplay() bool
	SetNeedsDisplay(value bool)
	// The view’s window object, if it is installed in a window.
	Window() INSWindow
	SetWindow(value INSWindow)
	// The menu item containing the view or any of its superviews in the view hierarchy.
	EnclosingMenuItem() INSMenuItem
	SetEnclosingMenuItem(value INSMenuItem)
	// The nearest ancestor scroll view that contains the current view.
	EnclosingScrollView() INSScrollView
	SetEnclosingScrollView(value INSScrollView)
	// The view’s closest opaque ancestor, which might be the view itself.
	OpaqueAncestor() INSView
	SetOpaqueAncestor(value INSView)
	// The array of views embedded in the current view.
	Subviews() INSView
	SetSubviews(value INSView)
	// The view that is the parent of the current view.
	Superview() INSView
	SetSuperview(value INSView)
	// The view’s tag, which is an integer that you use to identify the view within your app.
	Tag() int
	SetTag(value int)
	// The angle of rotation, measured in degrees, applied to the view’s bounds rectangle relative to its frame rectangle.
	BoundsRotation() float64
	SetBoundsRotation(value float64)
	// The angle of rotation, measured in degrees, applied to the view’s frame rectangle relative to its superview’s coordinate system.
	FrameRotation() float64
	SetFrameRotation(value float64)
	// A Boolean value indicating whether the view uses a flipped coordinate system.
	IsFlipped() bool
	SetIsFlipped(value bool)
	// A Boolean value indicating whether the view or any of its ancestors has ever had a rotation factor applied to its frame or bounds.
	IsRotatedFromBase() bool
	SetIsRotatedFromBase(value bool)
	// A Boolean value indicating whether the view or any of its ancestors has ever had a rotation factor applied to its frame or bounds, or has been scaled from the window’s base coordinate system.
	IsRotatedOrScaledFromBase() bool
	SetIsRotatedOrScaledFromBase(value bool)
	// A Boolean value indicating whether the view posts notifications when its bounds rectangle changes.
	PostsBoundsChangedNotifications() bool
	SetPostsBoundsChangedNotifications(value bool)
	// A Boolean value indicating whether the view posts notifications when its frame rectangle changes.
	PostsFrameChangedNotifications() bool
	SetPostsFrameChangedNotifications(value bool)
	// A Boolean value indicating whether the view ensures it is vibrant on top of other content.
	AllowsVibrancy() bool
	SetAllowsVibrancy(value bool)
	// The focus ring mask bounds, specified in the view’s coordinate space.
	FocusRingMaskBounds() corefoundation.CGRect
	SetFocusRingMaskBounds(value corefoundation.CGRect)
	// The type of focus ring drawn around the view.
	FocusRingType() NSFocusRingType
	SetFocusRingType(value NSFocusRingType)
	// A Boolean value indicating whether the view or one of its ancestors is being drawn for a find indicator.
	IsDrawingFindIndicator() bool
	SetIsDrawingFindIndicator(value bool)
	// A Boolean value indicating whether the view is hidden.
	IsHidden() bool
	SetIsHidden(value bool)
	// A Boolean value indicating whether the view is hidden from sight because it, or one of its ancestors, is marked as hidden.
	IsHiddenOrHasHiddenAncestor() bool
	SetIsHiddenOrHasHiddenAncestor(value bool)
	// The opacity of the view.
	AlphaValue() float64
	SetAlphaValue(value float64)
	// An array of Core Image filters to apply to the view’s background.
	BackgroundFilters() objectivec.IObject
	SetBackgroundFilters(value objectivec.IObject)
	// A Boolean value indicating whether the view incorporates content from its subviews into its own layer.
	CanDrawSubviewsIntoLayer() bool
	SetCanDrawSubviewsIntoLayer(value bool)
	// The Core Image filter used to composite the view’s contents with its background.
	CompositingFilter() objectivec.IObject
	SetCompositingFilter(value objectivec.IObject)
	// An array of Core Image filters to apply to the contents of the view and its sublayers.
	ContentFilters() objectivec.IObject
	SetContentFilters(value objectivec.IObject)
	// The rotation angle of the view around the center of its layer.
	FrameCenterRotation() float64
	SetFrameCenterRotation(value float64)
	// The Core Animation layer that the view uses as its backing store.
	Layer() objectivec.IObject
	SetLayer(value objectivec.IObject)
	// The current layer contents placement policy.
	LayerContentsPlacement() objectivec.IObject
	SetLayerContentsPlacement(value objectivec.IObject)
	// The contents redraw policy for the view’s layer.
	LayerContentsRedrawPolicy() objectivec.IObject
	SetLayerContentsRedrawPolicy(value objectivec.IObject)
	// A Boolean value indicating whether the view’s layer uses Core Image filters and needs in-process rendering.
	LayerUsesCoreImageFilters() bool
	SetLayerUsesCoreImageFilters(value bool)
	// The shadow displayed underneath the view.
	Shadow() INSShadow
	SetShadow(value INSShadow)
	// A Boolean value indicating whether the view uses a layer as its backing store.
	WantsLayer() bool
	SetWantsLayer(value bool)
	// A Boolean value indicating which drawing path the view takes when updating its contents.
	WantsUpdateLayer() bool
	SetWantsUpdateLayer(value bool)
	// The text for the view’s tooltip.
	ToolTip() string
	SetToolTip(value string)
	// Custom insets that you specify to modify your view’s safe area
	AdditionalSafeAreaInsets() foundation.NSEdgeInsets
	SetAdditionalSafeAreaInsets(value foundation.NSEdgeInsets)
	// The insets (in points) from the view’s frame that define its content rectangle.
	AlignmentRectInsets() foundation.NSEdgeInsets
	SetAlignmentRectInsets(value foundation.NSEdgeInsets)
	// A Boolean value indicating whether the view applies the autoresizing behavior to its subviews when its frame size changes.
	AutoresizesSubviews() bool
	SetAutoresizesSubviews(value bool)
	// The options that determine how the view is resized relative to its superview.
	AutoresizingMask() uint
	SetAutoresizingMask(value uint)
	// The distance (in points) between the bottom of the view’s alignment rectangle and its baseline.
	BaselineOffsetFromBottom() float64
	SetBaselineOffsetFromBottom(value float64)
	// A layout anchor representing the bottom edge of the view’s frame.
	BottomAnchor() INSLayoutYAxisAnchor
	SetBottomAnchor(value INSLayoutYAxisAnchor)
	// A layout anchor representing the horizontal center of the view’s frame.
	CenterXAnchor() INSLayoutXAxisAnchor
	SetCenterXAnchor(value INSLayoutXAxisAnchor)
	// A layout anchor representing the vertical center of the view’s frame.
	CenterYAnchor() INSLayoutYAxisAnchor
	SetCenterYAnchor(value INSLayoutYAxisAnchor)
	// Returns the constraints held by the view.
	Constraints() INSLayoutConstraint
	SetConstraints(value INSLayoutConstraint)
	// A layout anchor representing the baseline for the topmost line of text in the view.
	FirstBaselineAnchor() INSLayoutYAxisAnchor
	SetFirstBaselineAnchor(value INSLayoutYAxisAnchor)
	// The distance (in points) between the top of the view’s alignment rectangle and its topmost baseline.
	FirstBaselineOffsetFromTop() float64
	SetFirstBaselineOffsetFromTop(value float64)
	// The minimum size of the view that satisfies the constraints it holds.
	FittingSize() corefoundation.CGSize
	SetFittingSize(value corefoundation.CGSize)
	// A Boolean value indicating whether the constraints impacting the layout of the view incompletely specify the location of the view.
	HasAmbiguousLayout() bool
	SetHasAmbiguousLayout(value bool)
	// A layout anchor representing the height of the view’s frame.
	HeightAnchor() INSLayoutDimension
	SetHeightAnchor(value INSLayoutDimension)
	// The natural size for the receiving view, considering only properties of the view itself.
	IntrinsicContentSize() corefoundation.CGSize
	SetIntrinsicContentSize(value corefoundation.CGSize)
	// A Boolean value that indicates whether the view’s horizontal size constraints are active.
	IsHorizontalContentSizeConstraintActive() bool
	SetIsHorizontalContentSizeConstraintActive(value bool)
	// A Boolean value that indicates whether the view’s vertical size constraints are active.
	IsVerticalContentSizeConstraintActive() bool
	SetIsVerticalContentSizeConstraintActive(value bool)
	// A layout anchor representing the baseline for the bottommost line of text in the view.
	LastBaselineAnchor() INSLayoutYAxisAnchor
	SetLastBaselineAnchor(value INSLayoutYAxisAnchor)
	// The distance (in points) between the bottom of the view’s alignment rectangle and its bottommost baseline.
	LastBaselineOffsetFromBottom() float64
	SetLastBaselineOffsetFromBottom(value float64)
	// The array of layout guide objects owned by this view.
	LayoutGuides() INSLayoutGuide
	SetLayoutGuides(value INSLayoutGuide)
	// A layout guide that provides the recommended amount of padding for content inside of a view.
	LayoutMarginsGuide() INSLayoutGuide
	SetLayoutMarginsGuide(value INSLayoutGuide)
	// A layout anchor representing the leading edge of the view’s frame.
	LeadingAnchor() INSLayoutXAxisAnchor
	SetLeadingAnchor(value INSLayoutXAxisAnchor)
	// A layout anchor representing the left edge of the view’s frame.
	LeftAnchor() INSLayoutXAxisAnchor
	SetLeftAnchor(value INSLayoutXAxisAnchor)
	// A Boolean value indicating whether the view needs a layout pass before it can be drawn.
	NeedsLayout() bool
	SetNeedsLayout(value bool)
	// A Boolean value indicating whether the view’s constraints need to be updated.
	NeedsUpdateConstraints() bool
	SetNeedsUpdateConstraints(value bool)
	// A layout anchor representing the right edge of the view’s frame.
	RightAnchor() INSLayoutXAxisAnchor
	SetRightAnchor(value INSLayoutXAxisAnchor)
	// The distances from the edges of your view that define the safe area.
	SafeAreaInsets() foundation.NSEdgeInsets
	SetSafeAreaInsets(value foundation.NSEdgeInsets)
	// The layout guide you use to position content inside your view’s safe area.
	SafeAreaLayoutGuide() INSLayoutGuide
	SetSafeAreaLayoutGuide(value INSLayoutGuide)
	// A rectangle in the view’s coordinate system that contains the unobscured portion of the view.
	SafeAreaRect() corefoundation.CGRect
	SetSafeAreaRect(value corefoundation.CGRect)
	// A layout anchor representing the top edge of the view’s frame.
	TopAnchor() INSLayoutYAxisAnchor
	SetTopAnchor(value INSLayoutYAxisAnchor)
	// A layout anchor representing the trailing edge of the view’s frame.
	TrailingAnchor() INSLayoutXAxisAnchor
	SetTrailingAnchor(value INSLayoutXAxisAnchor)
	// A Boolean value indicating whether the view’s autoresizing mask is translated into constraints for the constraint-based layout system.
	TranslatesAutoresizingMaskIntoConstraints() bool
	SetTranslatesAutoresizingMaskIntoConstraints(value bool)
	// The layout direction for content in the view.
	UserInterfaceLayoutDirection() NSUserInterfaceLayoutDirection
	SetUserInterfaceLayoutDirection(value NSUserInterfaceLayoutDirection)
	// A layout anchor representing the width of the view’s frame.
	WidthAnchor() INSLayoutDimension
	SetWidthAnchor(value INSLayoutDimension)
	// A Boolean value indicating whether the view can draw its contents on a background thread.
	CanDrawConcurrently() bool
	SetCanDrawConcurrently(value bool)
	// A Boolean value that indicates whether the view, and its subviews, confine their drawing areas to the bounds of the view.
	ClipsToBounds() bool
	SetClipsToBounds(value bool)
	// A Boolean value indicating whether the view is being rendered as part of a live resizing operation.
	InLiveResize() bool
	SetInLiveResize(value bool)
	// A Boolean value indicating whether the view is in full screen mode.
	IsInFullScreenMode() bool
	SetIsInFullScreenMode(value bool)
	// A Boolean value indicating whether the view fills its frame rectangle with opaque content.
	IsOpaque() bool
	SetIsOpaque(value bool)
	// A Boolean value indicating whether the view optimizes live-resize operations by preserving content that has not moved.
	PreservesContentDuringLiveResize() bool
	SetPreservesContentDuringLiveResize(value bool)
	// The rectangle identifying the portion of your view that did not change during a live resize operation.
	RectPreservedDuringLiveResize() corefoundation.CGRect
	SetRectPreservedDuringLiveResize(value corefoundation.CGRect)
	// The portion of the view that isn’t clipped by its superviews.
	VisibleRect() corefoundation.CGRect
	SetVisibleRect(value corefoundation.CGRect)
	// A Boolean value indicating whether AppKit’s default clipping behavior is in effect.
	WantsDefaultClipping() bool
	SetWantsDefaultClipping(value bool)
	// The fraction of the page that can be pushed onto the next page during automatic pagination to prevent items such as lines of text from being divided across pages.
	HeightAdjustLimit() float64
	SetHeightAdjustLimit(value float64)
	// A default footer string that includes the current page number and page count.
	PageFooter() foundation.NSAttributedString
	SetPageFooter(value foundation.NSAttributedString)
	// A default header string that includes the print job title and date.
	PageHeader() foundation.NSAttributedString
	SetPageHeader(value foundation.NSAttributedString)
	// The view’s print job title.
	PrintJobTitle() string
	SetPrintJobTitle(value string)
	// The fraction of the page that can be pushed onto the next page during automatic pagination to prevent items such as small images or text columns from being divided across pages.
	WidthAdjustLimit() float64
	SetWidthAdjustLimit(value float64)
	// The types of touch interactions the view allows.
	AllowedTouchTypes() NSTouchTypeMask
	SetAllowedTouchTypes(value NSTouchTypeMask)
	// A Boolean value indicating whether the view can become key view.
	CanBecomeKeyView() bool
	SetCanBecomeKeyView(value bool)
	CandidateListTouchBarItem() INSCandidateListTouchBarItem
	SetCandidateListTouchBarItem(value INSCandidateListTouchBarItem)
	// The gesture recognize objects currently attached to the view.
	GestureRecognizers() INSGestureRecognizer
	SetGestureRecognizers(value INSGestureRecognizer)
	// The text input context object for the view.
	InputContext() INSTextInputContext
	SetInputContext(value INSTextInputContext)
	// A Boolean value indicating whether the view can pass mouse down events through to its superviews.
	MouseDownCanMoveWindow() bool
	SetMouseDownCanMoveWindow(value bool)
	// A Boolean value indicating whether the view needs its panel to become the key window before it can handle keyboard input and navigation.
	NeedsPanelToBecomeKey() bool
	SetNeedsPanelToBecomeKey(value bool)
	// The view object that follows the current view in the key view loop.
	NextKeyView() INSView
	SetNextKeyView(value INSView)
	// The closest view object in the key view loop that follows the current view in the key view loop and accepts first responder status.
	NextValidKeyView() INSView
	SetNextValidKeyView(value INSView)
	// The portion of the view that has been rendered and is available for responsive scrolling.
	PreparedContentRect() corefoundation.CGRect
	SetPreparedContentRect(value corefoundation.CGRect)
	// Configures the behavior and progression of the Force Touch trackpad when responding to touch input produced by the user when the cursor is over the view.
	PressureConfiguration() INSPressureConfiguration
	SetPressureConfiguration(value INSPressureConfiguration)
	// The view object preceding the current view in the key view loop.
	PreviousKeyView() INSView
	SetPreviousKeyView(value INSView)
	// The closest view object in the key view loop that precedes the current view and accepts first responder status.
	PreviousValidKeyView() INSView
	SetPreviousValidKeyView(value INSView)
	// The array of pasteboard drag types that the view can accept.
	RegisteredDraggedTypes() NSPasteboardType
	SetRegisteredDraggedTypes(value NSPasteboardType)
	// An array of the view’s tracking areas.
	TrackingAreas() INSTrackingArea
	SetTrackingAreas(value INSTrackingArea)
	// A Boolean value indicating whether the view wants resting touches.
	WantsRestingTouches() bool
	SetWantsRestingTouches(value bool)
	// A Boolean value indicating whether the view accepts touch events.
	AcceptsTouchEvents() bool
	SetAcceptsTouchEvents(value bool)
	// A Boolean value indicating whether drawing commands will produce any results.
	CanDraw() bool
	SetCanDraw(value bool)
	// A Boolean value indicating whether the view wants an OpenGL backing surface with a resolution greater than 1 pixel per point.
	WantsBestResolutionOpenGLSurface() bool
	SetWantsBestResolutionOpenGLSurface(value bool)
	WantsExtendedDynamicRangeOpenGLSurface() bool
	SetWantsExtendedDynamicRangeOpenGLSurface(value bool)
	// Adds a constraint on the layout of the receiving view or its subviews.
	AddConstraint(constraint INSLayoutConstraint)
	// Attaches a gesture recognizer to the view.
	AddGestureRecognizer(gestureRecognizer INSGestureRecognizer)
	// Adds multiple constraints on the layout of the receiving view or its subviews.
	AddConstraints(constraints []NSLayoutConstraint)
	// Adds the provided layout guide to the view.
	AddLayoutGuide(guide INSLayoutGuide)
	// Adds a view to the view’s subviews so it’s displayed above its siblings.
	AddSubview(view INSView)
	// Inserts a view among the view’s subviews so it’s displayed immediately above or below another view.
	AddSubviewPositionedRelativeTo(view INSView, place NSWindowOrderingMode, otherView INSView)
	// Creates a tooltip for a defined area in the view and returns a tag that identifies the tooltip rectangle.
	AddToolTipRectOwnerUserData(rect corefoundation.CGRect, owner objectivec.IObject, data unsafe.Pointer) objectivec.IObject
	// Establishes  an area for tracking mouse-entered and mouse-exited events within the view and returns a tag that identifies the tracking rectangle.
	AddTrackingRectOwnerUserDataAssumeInside(rect corefoundation.CGRect, owner objectivec.IObject, data unsafe.Pointer, flag bool) objectivec.IObject
	// Establishes  the cursor to be used when the mouse pointer lies within a specified region.
	AddCursorRectCursor(rect corefoundation.CGRect, object INSCursor)
	// Overridden by subclasses to modify a given rectangle, returning the altered rectangle.
	AdjustScroll(newVisible corefoundation.CGRect) corefoundation.CGRect
	// Overridden by subclasses to adjust page width during automatic pagination.
	AdjustPageWidthNewLeftRightLimit(newRight float64, oldLeft float64, oldRight float64, rightLimit float64)
	// Returns the view’s alignment rectangle for a given frame.
	AlignmentRectForFrame(frame corefoundation.CGRect) corefoundation.CGRect
	// Adds a given tracking area to the view.
	AddTrackingArea(trackingArea INSTrackingArea)
	// Returns the closest ancestor shared by the view and another specified view.
	AncestorSharedWithView(view INSView) INSView
	// Invoked at the beginning of the printing session, this method sets up the current graphics context.
	BeginDocument()
	// Called at the beginning of each page, this method sets up the coordinate system so that a region inside the view’s bounds is translated to a specified location.
	BeginPageInRectAtPlacement(rect corefoundation.CGRect, location corefoundation.CGPoint)
	// Returns a bitmap-representation object suitable for caching the specified portion of the view.
	BitmapImageRepForCachingDisplayInRect(rect corefoundation.CGRect) INSBitmapImageRep
	// Returns a backing store pixel-aligned rectangle in local view coordinates.
	BackingAlignedRectOptions(rect corefoundation.CGRect, options foundation.AlignmentOptions) corefoundation.CGRect
	// Draws the specified area of the view, and its descendants, into a provided bitmap-representation object.
	CacheDisplayInRectToBitmapImageRep(rect corefoundation.CGRect, bitmapImageRep INSBitmapImageRep)
	// Returns the constraints impacting the layout of the view for a given orientation.
	ConstraintsAffectingLayoutForOrientation(orientation NSLayoutConstraintOrientation) []NSLayoutConstraint
	// Scrolls the view’s closest ancestor [NSClipView](<doc://com.apple.appkit/documentation/AppKit/NSClipView>) object proportionally to the distance of an event that occurs outside of it.
	Autoscroll(event INSEvent) bool
	// Initiates a dragging session with a group of dragging items.
	BeginDraggingSessionWithItemsEventSource(items []NSDraggingItem, event INSEvent, source NSDraggingSource) INSDraggingSession
	// Converts the corners of a specified rectangle to lie on the center of device pixels, which is useful in compensating for rendering overscanning when the coordinate system has been scaled.
	CenterScanRect(rect corefoundation.CGRect) corefoundation.CGRect
	// Returns the priority with which a view resists being made smaller than its intrinsic size.
	ContentCompressionResistancePriorityForOrientation(orientation NSLayoutConstraintOrientation) NSLayoutPriority
	// Converts a point from the coordinate system of a given view to that of the view.
	ConvertPointFromView(point corefoundation.CGPoint, view INSView) corefoundation.CGPoint
	// Returns the priority with which a view resists being made larger than its intrinsic size.
	ContentHuggingPriorityForOrientation(orientation NSLayoutConstraintOrientation) NSLayoutPriority
	// Converts a size from another view’s coordinate system to that of the view.
	ConvertSizeFromView(size corefoundation.CGSize, view INSView) corefoundation.CGSize
	// Converts a rectangle from the coordinate system of another view to that of the view.
	ConvertRectFromView(rect corefoundation.CGRect, view INSView) corefoundation.CGRect
	// Converts a rectangle from the view’s coordinate system to that of another view.
	ConvertRectToView(rect corefoundation.CGRect, view INSView) corefoundation.CGRect
	// Converts a point from its pixel aligned backing store coordinate system to the view’s interior coordinate system.
	ConvertPointFromBacking(point corefoundation.CGPoint) corefoundation.CGPoint
	// Converts a point from the view’s coordinate system to that of a given view.
	ConvertPointToView(point corefoundation.CGPoint, view INSView) corefoundation.CGPoint
	// Overridden by subclasses to adjust page height during automatic pagination.
	AdjustPageHeightNewTopBottomLimit(newBottom float64, oldTop float64, oldBottom float64, bottomLimit float64)
	// Convert the point from the layer’s interior coordinate system to the view’s interior coordinate system.
	ConvertPointFromLayer(point corefoundation.CGPoint) corefoundation.CGPoint
	// Converts a size from its pixel aligned backing store coordinate system to the view’s interior coordinate system.
	ConvertSizeFromBacking(size corefoundation.CGSize) corefoundation.CGSize
	// Converts a point from the view’s interior coordinate system to its pixel aligned backing store coordinate system.
	ConvertPointToBacking(point corefoundation.CGPoint) corefoundation.CGPoint
	// Convert the rectangle from the layer’s interior coordinate system to the view’s interior coordinate system.
	ConvertRectFromLayer(rect corefoundation.CGRect) corefoundation.CGRect
	// Returns a Boolean value that indicates whether the view accepts the initial mouse-down event.
	AcceptsFirstMouse(event INSEvent) bool
	// Converts a rectangle from the view’s interior coordinate system to its pixel aligned backing store coordinate system.
	ConvertRectToBacking(rect corefoundation.CGRect) corefoundation.CGRect
	// Converts a size from the view’s interior coordinate system to its pixel aligned backing store coordinate system.
	ConvertSizeToBacking(size corefoundation.CGSize) corefoundation.CGSize
	// Returns EPS data that draws the region of the view within a specified rectangle.
	DataWithEPSInsideRect(rect corefoundation.CGRect) foundation.INSData
	// Convert the size from the view’s interior coordinate system to the layer’s interior coordinate system.
	ConvertRectToLayer(rect corefoundation.CGRect) corefoundation.CGRect
	// Called after a contextual menu that was displayed from the receiving view has been closed.
	DidCloseMenuWithEvent(menu INSMenu, event INSEvent)
	// Convert the size from the view’s interior coordinate system to the layer’s interior coordinate system.
	ConvertSizeToLayer(size corefoundation.CGSize) corefoundation.CGSize
	// Invalidates all cursor rectangles set up using [addCursorRect(_:cursor:)](<doc://com.apple.appkit/documentation/AppKit/NSView/addCursorRect(_:cursor:)>).
	DiscardCursorRects()
	// Convert the size from the view’s interior coordinate system to the layer’s interior coordinate system.
	ConvertPointToLayer(point corefoundation.CGPoint) corefoundation.CGPoint
	// Displays the view and all its subviews if any part of the view has been marked as needing display.
	DisplayIfNeeded()
	// Overridden by subclasses to perform additional actions when subviews are added to the view.
	DidAddSubview(subview INSView)
	// Displays the view and all its subviews if possible, invoking each of the [NSView] methods [lockFocus()](<doc://com.apple.appkit/documentation/AppKit/NSView/lockFocus()>), [draw(_:)](<doc://com.apple.appkit/documentation/AppKit/NSView/draw(_:)>), and [unlockFocus()](<doc://com.apple.appkit/documentation/AppKit/NSView/unlockFocus()>) as necessary.
	Display()
	// Acts as [display()](<doc://com.apple.appkit/documentation/AppKit/NSView/display()>), but confining drawing to a rectangular region of the view.
	DisplayRect(rect corefoundation.CGRect)
	// Acts as [displayIfNeeded()](<doc://com.apple.appkit/documentation/AppKit/NSView/displayIfNeeded()>), except that this method doesn’t back up to the first opaque ancestor—it simply causes the view and its descendants to execute their drawing code.
	DisplayIfNeededIgnoringOpacity()
	// Acts as [displayIfNeeded()](<doc://com.apple.appkit/documentation/AppKit/NSView/displayIfNeeded()>), confining drawing to a specified region of the view.
	DisplayIfNeededInRect(rect corefoundation.CGRect)
	DisplayLinkWithTargetSelector(target objectivec.IObject, selector objc.SEL) objectivec.IObject
	// Acts as [displayIfNeeded()](<doc://com.apple.appkit/documentation/AppKit/NSView/displayIfNeeded()>), but confining drawing to `aRect` and not backing up to the first opaque ancestor—it simply causes the view and its descendants to execute their drawing code.
	DisplayIfNeededInRectIgnoringOpacity(rect corefoundation.CGRect)
	// Draws the focus ring mask for the view.
	DrawFocusRingMask()
	// Displays the view but confines drawing to a specified region and does not back up to the first opaque ancestor—it simply causes the view and its descendants to execute their drawing code.
	DisplayRectIgnoringOpacity(rect corefoundation.CGRect)
	// Returns PDF data that draws the region of the view within a specified rectangle.
	DataWithPDFInsideRect(rect corefoundation.CGRect) foundation.INSData
	// Converts a rectangle from its pixel aligned backing store coordinate system to the view’s interior coordinate system.
	ConvertRectFromBacking(rect corefoundation.CGRect) corefoundation.CGRect
	// Allows applications that use the AppKit pagination facility to draw additional marks on each logical page.
	DrawPageBorderWithSize(borderSize corefoundation.CGSize)
	// This method is invoked at the end of the printing session.
	EndDocument()
	// Writes the end of a conforming page.
	EndPage()
	// Randomly changes the frame of a view with an ambiguous layout between the different valid values.
	ExerciseAmbiguityInLayout()
	// Instructs the view to exit full screen mode.
	ExitFullScreenModeWithOptions(options foundation.INSDictionary)
	// Overridden by subclasses to draw the view’s image within the specified rectangle.
	DrawRect(dirtyRect corefoundation.CGRect)
	// Sets the view to full screen mode.
	EnterFullScreenModeWithOptions(screen INSScreen, options foundation.INSDictionary) bool
	// Returns the view’s frame for a given alignment rectangle.
	FrameForAlignmentRect(alignmentRect corefoundation.CGRect) corefoundation.CGRect
	// Causes the view and its descendants to be redrawn to the specified graphics context.
	DisplayRectIgnoringOpacityInContext(rect corefoundation.CGRect, context INSGraphicsContext)
	// Returns a list of rectangles indicating the newly exposed areas of the view.
	GetRectsExposedDuringLiveResizeCount(exposedRects []corefoundation.CGRect, count int)
	// Returns by indirection a list of nonoverlapping rectangles that define the area the view is being asked to draw in [draw(_:)](<doc://com.apple.appkit/documentation/AppKit/NSView/draw(_:)>).
	GetRectsBeingDrawnCount(rects []corefoundation.CGRect, count int)
	// Returns the farthest descendant of the view in the view hierarchy (including itself) that contains a specified point, or `nil` if that point lies completely outside the view.
	HitTest(point corefoundation.CGPoint) INSView
	// Returns a Boolean value that indicates whether the view handles page boundaries.
	KnowsPageRange(range_ foundation.NSRange) bool
	// Returns a Boolean value that indicates whether the view is a subview of the specified view.
	IsDescendantOf(view INSView) bool
	// Invalidates the view’s intrinsic content size.
	InvalidateIntrinsicContentSize()
	// Returns whether a region of the view contains a specified point, accounting for whether the view is flipped or not.
	MouseInRect(point corefoundation.CGPoint, rect corefoundation.CGRect) bool
	// Perform layout in concert with the constraint-based layout system.
	Layout()
	// Updates the layout of the receiving view and its subviews based on the current views and constraints.
	LayoutSubtreeIfNeeded()
	// Creates the view’s backing layer.
	MakeBackingLayer() objectivec.IObject
	// Returns a Boolean value indicating whether the specified rectangle intersects any part of the area that the view is being asked to draw.
	NeedsToDrawRect(rect corefoundation.CGRect) bool
	// Invoked to notify the view that the focus ring mask requires updating.
	NoteFocusRingMaskChanged()
	// Prepares the overdraw region for drawing.
	PrepareContentInRect(rect corefoundation.CGRect)
	// This action method opens the Print panel, and if the user chooses an option other than canceling, prints the view and all its subviews to the device specified in the Print panel.
	Print(sender objectivec.IObject)
	// Returns the appropriate rectangle to use when magnifying around the specified point.
	RectForSmartMagnificationAtPointInRect(location corefoundation.CGPoint, visibleRect corefoundation.CGRect) corefoundation.CGRect
	// Overridden by subclasses to return a context-sensitive pop-up menu for a given mouse-down event.
	MenuForEvent(event INSEvent) INSMenu
	// Implemented by subclasses to determine the portion of the view to be printed for the specified page number.
	RectForPage(page int) corefoundation.CGRect
	// Registers the pasteboard types that the view will accept as the destination of an image-dragging session.
	RegisterForDraggedTypes(newTypes []string)
	// Removes all tooltips assigned to the view.
	RemoveAllToolTips()
	// Notifies a clip view’s superview that either the clip view’s bounds rectangle or the document view’s frame rectangle has changed, and that any indicators of the scroll position need to be adjusted.
	ReflectScrolledClipView(clipView INSClipView)
	// Removes the specified constraint from the view.
	RemoveConstraint(constraint INSLayoutConstraint)
	// Unlinks the view from its superview and its window, removes it from the responder chain, and invalidates its cursor rectangles.
	RemoveFromSuperview()
	// Detaches a gesture recognizer from the view.
	RemoveGestureRecognizer(gestureRecognizer INSGestureRecognizer)
	// Unlinks the view from its superview and its window and removes it from the responder chain, but does not invalidate its cursor rectangles to cause redrawing.
	RemoveFromSuperviewWithoutNeedingDisplay()
	// Completely removes a cursor rectangle from the view.
	RemoveCursorRectCursor(rect corefoundation.CGRect, object INSCursor)
	// Removes the tooltip identified by specified tag.
	RemoveToolTip(tag objectivec.IObject)
	// Removes the provided layout guide from the view.
	RemoveLayoutGuide(guide INSLayoutGuide)
	// Removes the tracking rectangle identified by a tag.
	RemoveTrackingRect(tag objectivec.IObject)
	// Removes a given tracking area from the view.
	RemoveTrackingArea(trackingArea INSTrackingArea)
	// Informs the view that the bounds size of its superview has changed.
	ResizeWithOldSuperviewSize(oldSize corefoundation.CGSize)
	// Replaces one of the view’s subviews with another view.
	ReplaceSubviewWith(oldView INSView, newView INSView)
	// Informs the view’s subviews that the view’s bounds rectangle size has changed.
	ResizeSubviewsWithOldSize(oldSize corefoundation.CGSize)
	// Rotates the view’s bounds rectangle by a specified degree value around the origin of the coordinate system, (0.0, 0.0).
	RotateByAngle(angle float64)
	// Informs the client that `aRulerView` allowed the user to move `aMarker`.
	RulerViewDidMoveMarker(ruler INSRulerView, marker INSRulerMarker)
	// Overridden by subclasses to define their default cursor rectangles.
	ResetCursorRects()
	// Informs the client that `aRulerView` allowed the user to remove `aMarker`.
	RulerViewDidRemoveMarker(ruler INSRulerView, marker INSRulerMarker)
	// Converts a size from the view’s coordinate system to that of another view.
	ConvertSizeToView(size corefoundation.CGSize, view INSView) corefoundation.CGSize
	// Informs the client that the user has pressed the mouse button while the cursor is in the ruler area of `aRulerView`.
	RulerViewHandleMouseDown(ruler INSRulerView, event INSEvent)
	RulerViewPointForLocation(ruler INSRulerView, point float64) corefoundation.CGPoint
	// Requests permission for `aRulerView` to add `aMarker`, an NSRulerMarker being dragged onto the ruler by the user.
	RulerViewShouldAddMarker(ruler INSRulerView, marker INSRulerMarker) bool
	// Requests permission for `aRulerView` to move `aMarker`.
	RulerViewShouldMoveMarker(ruler INSRulerView, marker INSRulerMarker) bool
	// Requests permission for `aRulerView` to remove `aMarker`.
	RulerViewShouldRemoveMarker(ruler INSRulerView, marker INSRulerMarker) bool
	// Informs the client that `aRulerView` will add the new NSRulerMarker, `aMarker`.
	RulerViewWillAddMarkerAtLocation(ruler INSRulerView, marker INSRulerMarker, location float64) float64
	// Removes the specified constraints from the view.
	RemoveConstraints(constraints []NSLayoutConstraint)
	// Informs the client that `aRulerView` will move `aMarker`, an NSRulerMarker already on the ruler view.
	RulerViewWillMoveMarkerToLocation(ruler INSRulerView, marker INSRulerMarker, location float64) float64
	RulerViewLocationForPoint(ruler INSRulerView, point corefoundation.CGPoint) float64
	// Informs the client view that `aRulerView` is about to be appropriated by `newClient`.
	RulerViewWillSetClientView(ruler INSRulerView, newClient INSView)
	// Scales the view’s coordinate system so that the unit square scales to the specified dimensions.
	ScaleUnitSquareToSize(newUnitSize corefoundation.CGSize)
	// Scrolls the view’s closest ancestor [NSClipView](<doc://com.apple.appkit/documentation/AppKit/NSClipView>) object the minimum distance needed so a specified region of the view becomes visible in the clip view.
	ScrollRectToVisible(rect corefoundation.CGRect) bool
	// Sets the origin of the view’s bounds rectangle to a specified point.
	SetBoundsOrigin(newOrigin corefoundation.CGPoint)
	// Sets the size of the view’s bounds rectangle to specified dimensions, inversely scaling its coordinate system relative to its frame rectangle.
	SetBoundsSize(newSize corefoundation.CGSize)
	// Sets the priority with which a view resists being made smaller than its intrinsic size.
	SetContentCompressionResistancePriorityForOrientation(priority NSLayoutPriority, orientation NSLayoutConstraintOrientation)
	// Sets the origin of the view’s frame rectangle to the specified point, effectively repositioning it within its superview.
	SetFrameOrigin(newOrigin corefoundation.CGPoint)
	// Informs the client that `aRulerView` allowed the user to add `aMarker`.
	RulerViewDidAddMarker(ruler INSRulerView, marker INSRulerMarker)
	// Invalidates the area around the focus ring.
	SetKeyboardFocusRingNeedsDisplayInRect(rect corefoundation.CGRect)
	// Notifies the superview of a clip view that the clip view needs to reset the origin of its bounds rectangle.
	ScrollClipViewToPoint(clipView INSClipView, point corefoundation.CGPoint)
	// Sets the priority with which a view resists being made larger than its intrinsic size.
	SetContentHuggingPriorityForOrientation(priority NSLayoutPriority, orientation NSLayoutConstraintOrientation)
	// Shows a window displaying the definition of the attributed string at the specified point.
	ShowDefinitionForAttributedStringAtPoint(attrString foundation.NSAttributedString, textBaselineOrigin corefoundation.CGPoint)
	// Convert the size from the layer’s interior coordinate system to the view’s interior coordinate system.
	ConvertSizeFromLayer(size corefoundation.CGSize) corefoundation.CGSize
	// Scrolls the view’s closest ancestor [NSClipView](<doc://com.apple.appkit/documentation/AppKit/NSClipView>) object so a point in the view lies at the origin of the clip view’s bounds rectangle.
	ScrollPoint(point corefoundation.CGPoint)
	// Marks the region of the view within the specified rectangle as needing display, increasing the view’s existing invalid region to include it.
	SetNeedsDisplayInRect(invalidRect corefoundation.CGRect)
	// Translates the view’s coordinate system so that its origin moves to a new location.
	TranslateOriginToPoint(translation corefoundation.CGPoint)
	// Orders the view’s immediate subviews using the specified comparator function.
	SortSubviewsUsingFunctionContext(compare objectivec.IObject, context unsafe.Pointer)
	// Unregisters the view as a possible destination in a dragging session.
	UnregisterDraggedTypes()
	// Updates the constraints for the receiving view and its subviews.
	UpdateConstraintsForSubtreeIfNeeded()
	// Update constraints for the view.
	UpdateConstraints()
	// Translates the display rectangles by the specified delta.
	TranslateRectsNeedingDisplayInRectBy(clipRect corefoundation.CGRect, delta corefoundation.CGSize)
	// Invoked automatically when the view’s geometry changes such that its tracking areas need to be recalculated.
	UpdateTrackingAreas()
	// Responds when the view’s backing store properties change.
	ViewDidChangeBackingProperties()
	// Informs the view that its effective appearance changed.
	ViewDidChangeEffectiveAppearance()
	// Shows a window displaying the definition of the specified range of the attributed string.
	ShowDefinitionForAttributedStringRangeOptionsBaselineOriginProvider(attrString foundation.NSAttributedString, targetRange foundation.NSRange, options foundation.INSDictionary, originProvider RangeHandler)
	// Updates the view’s content by modifying its underlying layer.
	UpdateLayer()
	// Invoked when the view is hidden, either directly, or in response to an ancestor being hidden.
	ViewDidHide()
	// Informs the view that its superview has changed (possibly to `nil`).
	ViewDidMoveToSuperview()
	// Invoked when the view is unhidden, either directly, or in response to an ancestor being unhidden
	ViewDidUnhide()
	// Informs the view that it has been added to a new view hierarchy.
	ViewDidMoveToWindow()
	// Informs the view of the end of a live resize—the user has finished resizing the view.
	ViewDidEndLiveResize()
	// Returns the view’s nearest descendant (including itself) with a specific tag, or `nil` if no subview has that tag.
	ViewWithTag(tag int) INSView
	// Informs the view that it’s being added to the view hierarchy of the specified window object (which may be `nil`).
	ViewWillMoveToWindow(newWindow INSWindow)
	// Informs the view of the start of a live resize—the user has started resizing the view.
	ViewWillStartLiveResize()
	// Called just before a contextual menu for a view is opened on screen.
	WillOpenMenuWithEvent(menu INSMenu, event INSEvent)
	// Informs the view that it’s required to draw content.
	ViewWillDraw()
	// Writes EPS data that draws the region of the view within a specified rectangle onto a pasteboard.
	WriteEPSInsideRectToPasteboard(rect corefoundation.CGRect, pasteboard INSPasteboard)
	// Writes PDF data that draws the region of the view within a specified rectangle onto a pasteboard.
	WritePDFInsideRectToPasteboard(rect corefoundation.CGRect, pasteboard INSPasteboard)
	// Overridden by subclasses to perform additional actions before subviews are removed from the view.
	WillRemoveSubview(subview INSView)
	// Sets the size of the view’s frame rectangle to the specified dimensions, resizing it within its superview without affecting its coordinate system.
	SetFrameSize(newSize corefoundation.CGSize)
	// Invoked by [printView(_:)](<doc://com.apple.appkit/documentation/AppKit/NSView/printView(_:)>) to determine the location of the region of the view being printed on the physical page.
	LocationOfPrintRect(rect corefoundation.CGRect) corefoundation.CGPoint
	// Informs the view that its superview is about to change to the specified superview (which may be `nil`).
	ViewWillMoveToSuperview(newSuperview INSView)
	// Allows the user to drag objects from the view without activating the app or moving the window of the view forward, possibly obscuring the destination.
	ShouldDelayWindowOrderingForEvent(event INSEvent) bool
	// Returns a proxy object for the receiver that can be used to initiate implied animation for property changes.
	Animator() INSView
	// Returns the animation that should be performed for the specified key.
	AnimationForKey(key NSAnimatablePropertyKey) objectivec.IObject
	EncodeWithCoder(coder foundation.INSCoder)
}





// Init initializes the instance.
func (v NSView) Init() NSView {
	rv := objc.Send[NSView](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v NSView) Autorelease() NSView {
	rv := objc.Send[NSView](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSView creates a new NSView instance.
func NewNSView() NSView {
	class := getNSViewClass()
	rv := objc.Send[NSView](objc.ID(class.class), objc.Sel("new"))
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
func NewViewWithCoder(coder foundation.INSCoder) NSView {
	instance := getNSViewClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSViewFromID(rv)
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
func NewViewWithFrame(frameRect corefoundation.CGRect) NSView {
	instance := getNSViewClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFrame:"), frameRect)
	return NSViewFromID(rv)
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
func (v NSView) InitWithFrame(frameRect corefoundation.CGRect) NSView {
	rv := objc.Send[NSView](v.ID, objc.Sel("initWithFrame:"), frameRect)
	return rv
}

// Restores the view to an initial state so that it can be reused.
//
// # Discussion
// 
// The default implementation of this method sets the view’s alpha to `1.0`
// and its hidden state to [false]. Subclasses can override this method and
// use it to return the view to its initial state. Subclasses should call
// `super` at some point in their implementation.
// 
// This method offers a way to reset a view to some initial state so that it
// can be reused. For example, the [NSTableView] class uses it to prepare
// views for reuse and thereby avoid the expense of creating new views as they
// scroll into view. If you implement a view-reuse system in your own code,
// you can call this method from your own code prior to reusing them.
//
// [false]: https://developer.apple.com/documentation/Swift/false
//
// See: https://developer.apple.com/documentation/AppKit/NSView/prepareForReuse()
func (v NSView) PrepareForReuse() {
	objc.Send[objc.ID](v.ID, objc.Sel("prepareForReuse"))
}

// Adds a constraint on the layout of the receiving view or its subviews.
//
// constraint: The constraint to be added to the view. The constraint may only reference
// the view itself or its subviews.
//
// # Discussion
// 
// The constraint must involve only views that are within scope of the
// receiving view. Specifically, any views involved must be either the
// receiving view itself, or a subview of the receiving view. Constraints that
// are added to a view are said to be held by that view. The coordinate system
// used when evaluating the constraint is the coordinate system of the view
// that holds the constraint.
//
// See: https://developer.apple.com/documentation/AppKit/NSView/addConstraint(_:)
func (v NSView) AddConstraint(constraint INSLayoutConstraint) {
	objc.Send[objc.ID](v.ID, objc.Sel("addConstraint:"), constraint)
}

// Attaches a gesture recognizer to the view.
//
// gestureRecognizer: The gesture recognizer to attach to the view. This parameter must not be
// `nil`.
//
// # Discussion
// 
// Attaching a gesture recognizer to a view defines the scope of the
// represented gesture, causing it to receive touches occurring only in the
// view or one of its subviews. The view establishes a strong reference to the
// specified gesture recognizer.
//
// See: https://developer.apple.com/documentation/AppKit/NSView/addGestureRecognizer(_:)
func (v NSView) AddGestureRecognizer(gestureRecognizer INSGestureRecognizer) {
	objc.Send[objc.ID](v.ID, objc.Sel("addGestureRecognizer:"), gestureRecognizer)
}

// Adds multiple constraints on the layout of the receiving view or its
// subviews.
//
// constraints: An array of constraints to be added to the view. All constraints may only
// reference the view itself or its subviews.
//
// # Discussion
// 
// All constraints must involve only views that are within scope of the
// receiving view. Specifically, any views involved must be either the
// receiving view itself, or a subview of the receiving view. Constraints that
// are added to a view are said to be held by that view. The coordinate system
// used when evaluating each constraint is the coordinate system of the view
// that holds the constraint.
//
// See: https://developer.apple.com/documentation/AppKit/NSView/addConstraints(_:)
func (v NSView) AddConstraints(constraints []NSLayoutConstraint) {
	objc.Send[objc.ID](v.ID, objc.Sel("addConstraints:"), objectivec.IObjectSliceToNSArray(constraints))
}

// Adds the provided layout guide to the view.
//
// guide: The layout guide to be added.
//
// # Discussion
// 
// This method adds the provided layout guide to the end of the view’s
// [layoutGuides] array. It also assigns the view to the guide’s
// [OwningView] property. Each guide can have only one owning view.
// 
// After the guide has been added to a view, it can participate in Auto Layout
// constraints with that view’s hierarchy.
//
// [layoutGuides]: https://developer.apple.com/documentation/AppKit/NSView/layoutGuides
//
// See: https://developer.apple.com/documentation/AppKit/NSView/addLayoutGuide(_:)
func (v NSView) AddLayoutGuide(guide INSLayoutGuide) {
	objc.Send[objc.ID](v.ID, objc.Sel("addLayoutGuide:"), guide)
}

// Adds a view to the view’s subviews so it’s displayed above its
// siblings.
//
// view: The view to add to the view as a subview.
//
// # Discussion
// 
// This method also sets the view as the next responder of `aView`.
// 
// The view retains `aView`. If you use [RemoveFromSuperview] to remove
// `aView` from the view hierarchy, `aView` is released. If you want to keep
// using `aView` after removing it from the view hierarchy (if, for example,
// you are swapping through a number of views), you must retain it before
// invoking [RemoveFromSuperview].
//
// See: https://developer.apple.com/documentation/AppKit/NSView/addSubview(_:)
func (v NSView) AddSubview(view INSView) {
	objc.Send[objc.ID](v.ID, objc.Sel("addSubview:"), view)
}

// Inserts a view among the view’s subviews so it’s displayed immediately
// above or below another view.
//
// view: The view object to add to the view as a subview.
//
// place: An `enum` constant specifying the position of the `aView` relative to
// `otherView`. Valid values are [NSWindowAbove] or [NSWindowBelow].
//
// otherView: The other view `aView` is to be positioned relative to. If `otherView` is
// `nil` (or isn’t a subview of the view), `aView` is added above or below
// all of its new siblings.
//
// # Discussion
// 
// This method also sets the view as the next responder of `aView`.
// 
// The view retains `aView`. If you use [RemoveFromSuperview] to remove
// `aView` from the view hierarchy, `aView` is released. If you want to keep
// using `aView` after removing it from the view hierarchy (if, for example,
// you are swapping through a number of views), you must retain it before
// invoking [RemoveFromSuperview].
//
// See: https://developer.apple.com/documentation/AppKit/NSView/addSubview(_:positioned:relativeTo:)
func (v NSView) AddSubviewPositionedRelativeTo(view INSView, place NSWindowOrderingMode, otherView INSView) {
	objc.Send[objc.ID](v.ID, objc.Sel("addSubview:positioned:relativeTo:"), view, place, otherView)
}

// Creates a tooltip for a defined area in the view and returns a tag that
// identifies the tooltip rectangle.
//
// rect: A rectangle defining the region of the view to associate the tooltip with.
//
// owner: An object from which to obtain the tooltip string. The object should either
// implement [view:stringForToolTip:point:userData:], or return a suitable
// string from its [description] method. It can therefore simply be an
// [NSString] object.
// //
// [NSString]: https://developer.apple.com/documentation/Foundation/NSString
// [description]: https://developer.apple.com/documentation/ObjectiveC/NSObjectProtocol/description
// [view:stringForToolTip:point:userData:]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/view:stringForToolTip:point:userData:
//
// data: Any additional information you want to pass to
// [view:stringForToolTip:point:userData:]; it isn’t used if `owner`
// doesn’t implement this method.
// //
// [view:stringForToolTip:point:userData:]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/view:stringForToolTip:point:userData:
//
// # Return Value
// 
// An integer tag identifying the tooltip; you can use this tag to remove the
// tooltip.
//
// # Discussion
// 
// The tooltip string is obtained dynamically from `owner` by invoking either
// the [NSToolTipOwner] informal protocol method
// [view:stringForToolTip:point:userData:], if implemented, or the
// [NSObjectProtocol] protocol method [description].
//
// [NSObjectProtocol]: https://developer.apple.com/documentation/ObjectiveC/NSObjectProtocol
// [description]: https://developer.apple.com/documentation/ObjectiveC/NSObjectProtocol/description
// [view:stringForToolTip:point:userData:]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/view:stringForToolTip:point:userData:
//
// See: https://developer.apple.com/documentation/AppKit/NSView/addToolTip(_:owner:userData:)
func (v NSView) AddToolTipRectOwnerUserData(rect corefoundation.CGRect, owner objectivec.IObject, data unsafe.Pointer) objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("addToolTipRect:owner:userData:"), rect, owner, data)
	return objectivec.Object{ID: rv}
}

// Establishes an area for tracking mouse-entered and mouse-exited events
// within the view and returns a tag that identifies the tracking rectangle.
//
// rect: A rectangle that defines a region of the view for tracking mouse-entered
// and mouse-exited events.
//
// owner: The object that gets sent the event messages. It can be the view itself or
// some other object (such as an NSCursor or a custom drawing tool object), as
// long as it responds to both [MouseEntered] and [MouseExited].
//
// data: Data stored in the [NSEvent] object for each tracking event.
//
// flag: If [true], the first event will be generated when the cursor leaves
// `aRect`, regardless if the cursor is inside `aRect` when the tracking
// rectangle is added. If [false] the first event will be generated when the
// cursor leaves `aRect` if the cursor is initially inside `aRect`, or when
// the cursor enters `aRect` if the cursor is initially outside `aRect`. You
// usually want to set this flag to [false].
// //
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Return Value
// 
// A tag that identifies the tracking rectangle. It is stored in the
// associated [NSEvent] objects and can be used to remove the tracking
// rectangle.
//
// # Discussion
// 
// Tracking rectangles provide a general mechanism that can be used to trigger
// actions based on the cursor location (for example, a status bar or hint
// field that provides information on the item the cursor lies over). To
// simply change the cursor over a particular area, use [AddCursorRectCursor].
// If you must use tracking rectangles to change the cursor, the [NSCursor]
// class specification describes the additional methods that must be invoked
// to change cursors by using tracking rectangles.
// 
// In macOS 10.5 and later, tracking areas provide a greater range of
// functionality (see [AddTrackingArea]).
//
// See: https://developer.apple.com/documentation/AppKit/NSView/addTrackingRect(_:owner:userData:assumeInside:)
func (v NSView) AddTrackingRectOwnerUserDataAssumeInside(rect corefoundation.CGRect, owner objectivec.IObject, data unsafe.Pointer, flag bool) objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("addTrackingRect:owner:userData:assumeInside:"), rect, owner, data, flag)
	return objectivec.Object{ID: rv}
}

// Establishes the cursor to be used when the mouse pointer lies within a
// specified region.
//
// rect: A rectangle defining a region of the view.
//
// object: An object representing a cursor.
//
// # Discussion
// 
// Cursor rectangles aren’t subject to clipping by superviews, nor are they
// intended for use with rotated views. You should explicitly confine a cursor
// rectangle to the view’s visible rectangle to prevent improper behavior.
// 
// This method is intended to be invoked only by the [ResetCursorRects]
// method. If invoked in any other way, the resulting cursor rectangle will be
// discarded the next time the view’s cursor rectangles are rebuilt.
//
// See: https://developer.apple.com/documentation/AppKit/NSView/addCursorRect(_:cursor:)
func (v NSView) AddCursorRectCursor(rect corefoundation.CGRect, object INSCursor) {
	objc.Send[objc.ID](v.ID, objc.Sel("addCursorRect:cursor:"), rect, object)
}

// Overridden by subclasses to modify a given rectangle, returning the altered
// rectangle.
//
// newVisible: A rectangle that defines a region of the view.
//
// # Discussion
// 
// [NSClipView] invokes this method to allow its document view to adjust its
// position during scrolling. For example, a custom view object that displays
// a table of data can adjust the origin of `newVisible` so rows or columns
// aren’t cut off by the edge of the enclosing [NSClipView]. The [NSView]
// implementation simply returns `newVisible`.
// 
// [NSClipView] only invokes this method during automatic or user controlled
// scrolling. Its [ScrollToPoint] method doesn’t invoke this method, so you
// can still force a scroll to an arbitrary point.
//
// See: https://developer.apple.com/documentation/AppKit/NSView/adjustScroll(_:)
func (v NSView) AdjustScroll(newVisible corefoundation.CGRect) corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](v.ID, objc.Sel("adjustScroll:"), newVisible)
	return corefoundation.CGRect(rv)
}

// Overridden by subclasses to adjust page width during automatic pagination.
//
// newRight: Returns by indirection a new [CGFloat] value for the right edge of the
// pending page rectangle in the view’s coordinate system.
// //
// [CGFloat]: https://developer.apple.com/documentation/CoreFoundation/CGFloat-swift.struct
//
// oldLeft: A [CGFloat] value that sets the left edge of the pending page rectangle in
// the view’s coordinate system.
// //
// [CGFloat]: https://developer.apple.com/documentation/CoreFoundation/CGFloat-swift.struct
//
// oldRight: A [CGFloat] value that sets the right edge of the pending page rectangle in
// the view’s coordinate system.
// //
// [CGFloat]: https://developer.apple.com/documentation/CoreFoundation/CGFloat-swift.struct
//
// rightLimit: The leftmost [CGFloat] value `newRight` can be set to, as calculated using
// the value of the [widthAdjustLimit] property.
// //
// [CGFloat]: https://developer.apple.com/documentation/CoreFoundation/CGFloat-swift.struct
// [widthAdjustLimit]: https://developer.apple.com/documentation/AppKit/NSView/widthAdjustLimit
//
// # Discussion
// 
// This method is invoked by [Print]. The view can pull in the right edge and
// return the new value in `newRight`, allowing it to prevent items such as
// small images or text columns from being divided across pages. If
// `rightLimit` is exceeded, the pagination mechanism simply uses `rightLimit`
// for the right edge.
// 
// The default implementation of this method propagates the message to its
// subviews, allowing nested views to adjust page width for their drawing as
// well. An [NSButton] object or other small view, for example, will nudge the
// right edge out if necessary to prevent itself from being cut in two
// (thereby pushing it onto an adjacent page). Subclasses should invoke
// `super`’s implementation, if desired, after first making their own
// adjustments.
//
// See: https://developer.apple.com/documentation/AppKit/NSView/adjustPageWidthNew(_:left:right:limit:)
func (v NSView) AdjustPageWidthNewLeftRightLimit(newRight float64, oldLeft float64, oldRight float64, rightLimit float64) {
	objc.Send[objc.ID](v.ID, objc.Sel("adjustPageWidthNew:left:right:limit:"), newRight, oldLeft, oldRight, rightLimit)
}

// Returns the view’s alignment rectangle for a given frame.
//
// frame: The frame whose corresponding alignment rectangle is desired.
//
// # Return Value
// 
// The alignment rectangle for the specified frame.
//
// # Discussion
// 
// The constraint-based layout system uses alignment rectangles to align
// views, rather than their frame. This allows custom views to be aligned
// based on the location of their content while still having a frame that
// encompasses any ornamentation they need to draw around their content, such
// as shadows or reflections.
// 
// The default implementation returns the view’s frame modified by the
// insets specified by the view’s [alignmentRectInsets] method. Most custom
// views can override [alignmentRectInsets] to specify the location of their
// content within their frame. Custom views that require arbitrary
// transformations can override [AlignmentRectForFrame] and
// [FrameForAlignmentRect] to describe the location of their content. These
// two methods must always be inverses of each other.
//
// [alignmentRectInsets]: https://developer.apple.com/documentation/AppKit/NSView/alignmentRectInsets
//
// See: https://developer.apple.com/documentation/AppKit/NSView/alignmentRect(forFrame:)
func (v NSView) AlignmentRectForFrame(frame corefoundation.CGRect) corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](v.ID, objc.Sel("alignmentRectForFrame:"), frame)
	return corefoundation.CGRect(rv)
}

// Adds a given tracking area to the view.
//
// trackingArea: The tracking area to add to the view.
//
// See: https://developer.apple.com/documentation/AppKit/NSView/addTrackingArea(_:)
func (v NSView) AddTrackingArea(trackingArea INSTrackingArea) {
	objc.Send[objc.ID](v.ID, objc.Sel("addTrackingArea:"), trackingArea)
}

// Returns the closest ancestor shared by the view and another specified view.
//
// view: Another view to test for closest shared ancestor with the view.
//
// # Return Value
// 
// The closest ancestor or `nil` if there’s no such object. Returns `self`
// if `aView` is identical to the view.
//
// See: https://developer.apple.com/documentation/AppKit/NSView/ancestorShared(with:)
func (v NSView) AncestorSharedWithView(view INSView) INSView {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("ancestorSharedWithView:"), view)
	return NSViewFromID(rv)
}

// Invoked at the beginning of the printing session, this method sets up the
// current graphics context.
//
// # Discussion
// 
// Note that this method may be invoked in a subthread.
// 
// Override it to configure printing related settings. You should store your
// settings in the object returned by [NSPrintInfo]’s [SharedPrintInfo]
// class method, which is guaranteed to return an instance specific to the
// thread in which you invoke this method. If you override this method, call
// the superclass implementation.
//
// See: https://developer.apple.com/documentation/AppKit/NSView/beginDocument()
func (v NSView) BeginDocument() {
	objc.Send[objc.ID](v.ID, objc.Sel("beginDocument"))
}

// Called at the beginning of each page, this method sets up the coordinate
// system so that a region inside the view’s bounds is translated to a
// specified location.
//
// rect: A rectangle defining the region to be translated.
//
// location: A point that is the end-point of translation.
//
// # Discussion
// 
// If you override this method, be sure to call the superclass implementation.
//
// See: https://developer.apple.com/documentation/AppKit/NSView/beginPage(in:atPlacement:)
func (v NSView) BeginPageInRectAtPlacement(rect corefoundation.CGRect, location corefoundation.CGPoint) {
	objc.Send[objc.ID](v.ID, objc.Sel("beginPageInRect:atPlacement:"), rect, location)
}

// Returns a bitmap-representation object suitable for caching the specified
// portion of the view.
//
// rect: A rectangle defining the area of the view to be cached.
//
// # Return Value
// 
// An autoreleased [NSBitmapImageRep] object or `nil` if the object could not
// be created.
//
// # Discussion
// 
// Passing the visible rectangle of the view (`[self visibleRect]`) returns a
// bitmap suitable for caching the current contents of the view, including all
// of its descendants.
//
// See: https://developer.apple.com/documentation/AppKit/NSView/bitmapImageRepForCachingDisplay(in:)
func (v NSView) BitmapImageRepForCachingDisplayInRect(rect corefoundation.CGRect) INSBitmapImageRep {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("bitmapImageRepForCachingDisplayInRect:"), rect)
	return NSBitmapImageRepFromID(rv)
}

// Returns a backing store pixel-aligned rectangle in local view coordinates.
//
// rect: The rectangle in the view’s interior coordinate system.
//
// options: The alignment options. See [AlignmentOptions] for possible values. (Note
// that although the alignment options specify integral values, the rectangle
// returned by this method is pixel-aligned.)
// //
// [AlignmentOptions]: https://developer.apple.com/documentation/Foundation/AlignmentOptions
//
// # Return Value
// 
// A rectangle in the view’s interior coordinate system that is aligned to
// the backing store pixels using the specified options.
//
// # Discussion
// 
// Uses the [NSIntegralRectWithOptions(_:_:)] function and the given input
// rectangle and options to produce a backing store pixel-aligned rectangle in
// the view’s interior coordinates.
//
// [NSIntegralRectWithOptions(_:_:)]: https://developer.apple.com/documentation/Foundation/NSIntegralRectWithOptions(_:_:)
//
// See: https://developer.apple.com/documentation/AppKit/NSView/backingAlignedRect(_:options:)
func (v NSView) BackingAlignedRectOptions(rect corefoundation.CGRect, options foundation.AlignmentOptions) corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](v.ID, objc.Sel("backingAlignedRect:options:"), rect, options)
	return corefoundation.CGRect(rv)
}

// Draws the specified area of the view, and its descendants, into a provided
// bitmap-representation object.
//
// rect: A rectangle defining the region to be drawn into `bitmapImageRep`.
//
// bitmapImageRep: An [NSBitmapImageRep] object. For pixel-format compatibility,
// `bitmapImageRep` should have been obtained from
// [BitmapImageRepForCachingDisplayInRect].
//
// # Discussion
// 
// You are responsible for initializing the bitmap to the desired
// configuration before calling this method. However, once initialized, you
// can reuse the same bitmap multiple times to refresh the cached copy of your
// view’s contents.
// 
// The bitmap produced by this method is transparent (that is, has an alpha
// value of 0) wherever the view and its descendants do not draw any content.
//
// See: https://developer.apple.com/documentation/AppKit/NSView/cacheDisplay(in:to:)
func (v NSView) CacheDisplayInRectToBitmapImageRep(rect corefoundation.CGRect, bitmapImageRep INSBitmapImageRep) {
	objc.Send[objc.ID](v.ID, objc.Sel("cacheDisplayInRect:toBitmapImageRep:"), rect, bitmapImageRep)
}

// Returns the constraints impacting the layout of the view for a given
// orientation.
//
// orientation: The direction of the dimension for which the constraints should be found.
//
// # Return Value
// 
// The constraints impacting the layout of the view for the specified
// orientation.
//
// # Discussion
// 
// The returned set of constraints may not all include the view explicitly.
// Constraints that impact the location of the view implicitly may also be
// included. While this provides a good starting point for debugging, there is
// no guarantee that the returned set of constraints will include all of the
// constraints that have an impact on the view’s layout in the given
// orientation.
// 
// This method should only be used for debugging constraint-based layout. No
// application should ship with calls to this method as part of its operation.
//
// See: https://developer.apple.com/documentation/AppKit/NSView/constraintsAffectingLayout(for:)
func (v NSView) ConstraintsAffectingLayoutForOrientation(orientation NSLayoutConstraintOrientation) []NSLayoutConstraint {
	rv := objc.Send[[]objc.ID](v.ID, objc.Sel("constraintsAffectingLayoutForOrientation:"), orientation)
	return objc.ConvertSlice(rv, func(id objc.ID) NSLayoutConstraint {
		return NSLayoutConstraintFromID(id)
	})
}

// Scrolls the view’s closest ancestor [NSClipView] object proportionally to
// the distance of an event that occurs outside of it.
//
// event: An event object whose location should be expressed in the window’s base
// coordinate system (which it normally is), not the receiving view’s.
//
// # Return Value
// 
// Returns [true] if any scrolling is performed; otherwise returns [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// View objects that track mouse-dragged events can use this method to scroll
// automatically when the cursor is dragged outside of the [NSClipView]
// object. Repeated invocations of this method (with an appropriate delay)
// result in continual scrolling, even when the mouse doesn’t move.
//
// See: https://developer.apple.com/documentation/AppKit/NSView/autoscroll(with:)
func (v NSView) Autoscroll(event INSEvent) bool {
	rv := objc.Send[bool](v.ID, objc.Sel("autoscroll:"), event)
	return rv
}

// Initiates a dragging session with a group of dragging items.
//
// items: The dragging items. The frame property of each [NSDraggingItem] must be in
// the view’s coordinate system.
//
// event: The mouse-down event object from which to initiate the drag operation. In
// particular, its mouse location is used for the offset of the icon being
// dragged.
//
// source: An object that serves as the controller of the dragging operation. It must
// conform to the [NSDraggingSource] protocol and is typically the view itself
// or its [NSWindow] object.
//
// # Return Value
// 
// The dragging session for the drag.
//
// # Discussion
// 
// A basic drag starts by calling ``.
// 
// The caller can take the returned [NSDraggingSession] and continue to modify
// its properties. When the drag actually starts, the source is sent a
// [DraggingSessionWillBeginAtPoint] message followed by multiple
// [DraggingSessionMovedToPoint] messages as the user drags.
// 
// Once the drag is ended or cancelled, the source receives a
// [DraggingSessionEndedAtPointOperation] method and the drag is complete.
//
// See: https://developer.apple.com/documentation/AppKit/NSView/beginDraggingSession(with:event:source:)
func (v NSView) BeginDraggingSessionWithItemsEventSource(items []NSDraggingItem, event INSEvent, source NSDraggingSource) INSDraggingSession {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("beginDraggingSessionWithItems:event:source:"), objectivec.IObjectSliceToNSArray(items), event, source)
	return NSDraggingSessionFromID(rv)
}

// Converts the corners of a specified rectangle to lie on the center of
// device pixels, which is useful in compensating for rendering overscanning
// when the coordinate system has been scaled.
//
// rect: The rectangle whose corners are to be converted.
//
// # Return Value
// 
// The adjusted rectangle.
//
// # Discussion
// 
// This method converts the given rectangle to device coordinates, adjusts the
// rectangle to lie in the center of the pixels, and converts the resulting
// rectangle back to the view’s coordinate system. Note that this method
// does not take into account any transformations performed using the
// [NSAffineTransform] class or Quartz 2D routines.
//
// [NSAffineTransform]: https://developer.apple.com/documentation/Foundation/NSAffineTransform
//
// See: https://developer.apple.com/documentation/AppKit/NSView/centerScanRect(_:)
func (v NSView) CenterScanRect(rect corefoundation.CGRect) corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](v.ID, objc.Sel("centerScanRect:"), rect)
	return corefoundation.CGRect(rv)
}

// Returns the priority with which a view resists being made smaller than its
// intrinsic size.
//
// orientation: The orientation of the dimension of the view that might be reduced.
//
// # Return Value
// 
// The priority with which the view should resist being compressed from its
// intrinsic size in the specified orientation.
//
// # Discussion
// 
// The constraint-based layout system uses these priorities when determining
// the best layout for views that are encountering constraints that would
// require them to be smaller than their intrinsic size.
// 
// Subclasses should not override this method. Instead, custom views should
// set default values for their content on creation, typically to [defaultLow]
// or [defaultHigh].
//
// [defaultHigh]: https://developer.apple.com/documentation/AppKit/NSLayoutConstraint/Priority-swift.struct/defaultHigh
// [defaultLow]: https://developer.apple.com/documentation/AppKit/NSLayoutConstraint/Priority-swift.struct/defaultLow
//
// See: https://developer.apple.com/documentation/AppKit/NSView/contentCompressionResistancePriority(for:)
func (v NSView) ContentCompressionResistancePriorityForOrientation(orientation NSLayoutConstraintOrientation) NSLayoutPriority {
	rv := objc.Send[NSLayoutPriority](v.ID, objc.Sel("contentCompressionResistancePriorityForOrientation:"), orientation)
	return NSLayoutPriority(rv)
}

// Converts a point from the coordinate system of a given view to that of the
// view.
//
// point: A point specifying a location in the coordinate system of `view`.
//
// view: The view with `point` in its coordinate system. Both `view` and the view
// must belong to the same [NSWindow] object, and that window must not be
// `nil`. If `view` is `nil`, this method converts from window coordinates
// instead.
//
// # Return Value
// 
// The point converted to the coordinate system of the view.
//
// See: https://developer.apple.com/documentation/AppKit/NSView/convert(_:from:)-1dq9l
func (v NSView) ConvertPointFromView(point corefoundation.CGPoint, view INSView) corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](v.ID, objc.Sel("convertPoint:fromView:"), point, view)
	return corefoundation.CGPoint(rv)
}

// Returns the priority with which a view resists being made larger than its
// intrinsic size.
//
// orientation: The orientation of the dimension of the view that might be enlarged.
//
// # Return Value
// 
// The priority with which the view should resist being enlarged from its
// intrinsic size in the specified orientation.
//
// # Discussion
// 
// The constraint-based layout system uses these priorities when determining
// the best layout for views that are encountering constraints that would
// require them to be larger than their intrinsic size.
// 
// Subclasses should not override this method. Instead, custom views should
// set default values for their content on creation, typically to [defaultLow]
// or [defaultHigh].
//
// [defaultHigh]: https://developer.apple.com/documentation/AppKit/NSLayoutConstraint/Priority-swift.struct/defaultHigh
// [defaultLow]: https://developer.apple.com/documentation/AppKit/NSLayoutConstraint/Priority-swift.struct/defaultLow
//
// See: https://developer.apple.com/documentation/AppKit/NSView/contentHuggingPriority(for:)
func (v NSView) ContentHuggingPriorityForOrientation(orientation NSLayoutConstraintOrientation) NSLayoutPriority {
	rv := objc.Send[NSLayoutPriority](v.ID, objc.Sel("contentHuggingPriorityForOrientation:"), orientation)
	return NSLayoutPriority(rv)
}

// Converts a size from another view’s coordinate system to that of the
// view.
//
// size: The size (width and height) in view’s coordinate system.
//
// view: The view with `size` in its coordinate system. Both `view` and the view
// must belong to the same [NSWindow] object, and that window must not be
// `nil`. If `view` is `nil`, this method converts from window coordinates
// instead.
//
// # Return Value
// 
// The converted size, as an [NSSize] structure.
//
// [NSSize]: https://developer.apple.com/documentation/Foundation/NSSize
//
// # Discussion
// 
// The returned [NSSize] values are always forced to have positive a width and
// height.
// 
// You can also use this method to get a view’s current magnification or
// zoom level, if it’s been changed from the default scale. Specifically, if
// you convert a known size from the window’s base coordinate space to that
// of `view`, the result is the current zoom level.
//
// See: https://developer.apple.com/documentation/AppKit/NSView/convert(_:from:)-40x0w
func (v NSView) ConvertSizeFromView(size corefoundation.CGSize, view INSView) corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](v.ID, objc.Sel("convertSize:fromView:"), size, view)
	return corefoundation.CGSize(rv)
}

// Converts a rectangle from the coordinate system of another view to that of
// the view.
//
// rect: The rectangle in the `view` coordinate system.
//
// view: The view with `rect` in its coordinate system. Both `view` and the view
// must belong to the same [NSWindow] object, and that window must not be
// `nil`. If `view` is `nil`, this method converts from window coordinates
// instead.
//
// # Return Value
// 
// The converted rectangle.
//
// See: https://developer.apple.com/documentation/AppKit/NSView/convert(_:from:)-7fbb6
func (v NSView) ConvertRectFromView(rect corefoundation.CGRect, view INSView) corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](v.ID, objc.Sel("convertRect:fromView:"), rect, view)
	return corefoundation.CGRect(rv)
}

// Converts a rectangle from the view’s coordinate system to that of another
// view.
//
// rect: A rectangle in the view’s coordinate system.
//
// view: The view that’s the target of the conversion operation. Both `view` and
// the view must belong to the same [NSWindow] object, and that window must
// not be `nil`. If `view` is `nil`, this method converts the rectangle to
// window coordinates instead.
//
// # Return Value
// 
// The converted rectangle.
//
// See: https://developer.apple.com/documentation/AppKit/NSView/convert(_:to:)-3cqqt
func (v NSView) ConvertRectToView(rect corefoundation.CGRect, view INSView) corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](v.ID, objc.Sel("convertRect:toView:"), rect, view)
	return corefoundation.CGRect(rv)
}

// Converts a point from its pixel aligned backing store coordinate system to
// the view’s interior coordinate system.
//
// point: The point in the pixel backing store aligned coordinate system.
//
// # Return Value
// 
// A point in the view’s interior coordinate system.
//
// See: https://developer.apple.com/documentation/AppKit/NSView/convertFromBacking(_:)-229ps
func (v NSView) ConvertPointFromBacking(point corefoundation.CGPoint) corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](v.ID, objc.Sel("convertPointFromBacking:"), point)
	return corefoundation.CGPoint(rv)
}

// Converts a point from the view’s coordinate system to that of a given
// view.
//
// point: A point specifying a location in the coordinate system of the view.
//
// view: The view into whose coordinate system `point` is to be converted. Both
// `view` and the view must belong to the same [NSWindow] object, and that
// window must not be `nil`. If `view` is `nil`, this method converts to
// window coordinates instead.
//
// # Return Value
// 
// The point converted to the coordinate system of `aView`.
//
// See: https://developer.apple.com/documentation/AppKit/NSView/convert(_:to:)-6u9ir
func (v NSView) ConvertPointToView(point corefoundation.CGPoint, view INSView) corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](v.ID, objc.Sel("convertPoint:toView:"), point, view)
	return corefoundation.CGPoint(rv)
}

// Overridden by subclasses to adjust page height during automatic pagination.
//
// newBottom: Returns by indirection a new [CGFloat] value for the bottom edge of the
// pending page rectangle in the view’s coordinate system.
// //
// [CGFloat]: https://developer.apple.com/documentation/CoreFoundation/CGFloat-swift.struct
//
// oldTop: A [CGFloat] value that sets the top edge of the pending page rectangle in
// the view’s coordinate system.
// //
// [CGFloat]: https://developer.apple.com/documentation/CoreFoundation/CGFloat-swift.struct
//
// oldBottom: A [CGFloat] value that sets the bottom edge of the pending page rectangle
// in the view’s coordinate system.
// //
// [CGFloat]: https://developer.apple.com/documentation/CoreFoundation/CGFloat-swift.struct
//
// bottomLimit: The topmost [CGFloat] value `newBottom` can be set to, as calculated using
// the value of the [heightAdjustLimit] property.
// //
// [CGFloat]: https://developer.apple.com/documentation/CoreFoundation/CGFloat-swift.struct
// [heightAdjustLimit]: https://developer.apple.com/documentation/AppKit/NSView/heightAdjustLimit
//
// # Discussion
// 
// This method is invoked by [Print]. The view can raise the bottom edge and
// return the new value in `newBottom`, allowing it to prevent items such as
// lines of text from being divided across pages. If `bottomLimit` is
// exceeded, the pagination mechanism simply uses `bottomLimit` for the bottom
// edge.
// 
// The default implementation of this method propagates the message to its
// subviews, allowing nested views to adjust page height for their drawing as
// well. An [NSButton] object or other small view, for example, will nudge the
// bottom edge up if necessary to prevent itself from being cut in two
// (thereby pushing it onto an adjacent page). Subclasses should invoke
// `super`’s implementation, if desired, after first making their own
// adjustments.
//
// See: https://developer.apple.com/documentation/AppKit/NSView/adjustPageHeightNew(_:top:bottom:limit:)
func (v NSView) AdjustPageHeightNewTopBottomLimit(newBottom float64, oldTop float64, oldBottom float64, bottomLimit float64) {
	objc.Send[objc.ID](v.ID, objc.Sel("adjustPageHeightNew:top:bottom:limit:"), newBottom, oldTop, oldBottom, bottomLimit)
}

// Convert the point from the layer’s interior coordinate system to the
// view’s interior coordinate system.
//
// point: The point in the layer’s interior coordinate system.
//
// # Return Value
// 
// The point in the view’s interior coordinate system.
//
// # Discussion
// 
// The layer’s space is virtual, and doesn’t take into account the
// layer’s [contentsScale] setting.
//
// [contentsScale]: https://developer.apple.com/documentation/QuartzCore/CALayer/contentsScale
//
// See: https://developer.apple.com/documentation/AppKit/NSView/convertFromLayer(_:)-3nsbu
func (v NSView) ConvertPointFromLayer(point corefoundation.CGPoint) corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](v.ID, objc.Sel("convertPointFromLayer:"), point)
	return corefoundation.CGPoint(rv)
}

// Converts a size from its pixel aligned backing store coordinate system to
// the view’s interior coordinate system.
//
// size: The size in the pixel aligned coordinate system.
//
// # Return Value
// 
// The size in the view’s interior coordinate system.
//
// See: https://developer.apple.com/documentation/AppKit/NSView/convertFromBacking(_:)-4agf9
func (v NSView) ConvertSizeFromBacking(size corefoundation.CGSize) corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](v.ID, objc.Sel("convertSizeFromBacking:"), size)
	return corefoundation.CGSize(rv)
}

// Converts a point from the view’s interior coordinate system to its pixel
// aligned backing store coordinate system.
//
// point: The point in the view’s interior coordinate system.
//
// # Return Value
// 
// A point in its pixel aligned backing store coordinate system.
//
// See: https://developer.apple.com/documentation/AppKit/NSView/convertToBacking(_:)-2xx45
func (v NSView) ConvertPointToBacking(point corefoundation.CGPoint) corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](v.ID, objc.Sel("convertPointToBacking:"), point)
	return corefoundation.CGPoint(rv)
}

// Convert the rectangle from the layer’s interior coordinate system to the
// view’s interior coordinate system.
//
// rect: A rectangle in the layer’s interior coordinate system.
//
// # Return Value
// 
// A rectangle in the view’s interior coordinate system.
//
// # Discussion
// 
// The layer’s space is virtual, and doesn’t take into account the
// layer’s [contentsScale] setting.
//
// [contentsScale]: https://developer.apple.com/documentation/QuartzCore/CALayer/contentsScale
//
// See: https://developer.apple.com/documentation/AppKit/NSView/convertFromLayer(_:)-8s5bi
func (v NSView) ConvertRectFromLayer(rect corefoundation.CGRect) corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](v.ID, objc.Sel("convertRectFromLayer:"), rect)
	return corefoundation.CGRect(rv)
}

// Returns a Boolean value that indicates whether the view accepts the initial
// mouse-down event.
//
// event: The initial mouse-down event, which must be over the view in its window.
//
// # Discussion
// 
// Subclasses can override this method to return [true] if the view should be
// sent a [MouseDown] message for an initial mouse-down event, [false] if not.
// 
// The view can either return a value unconditionally or use the location of
// `event` to determine whether or not it wants the event. The default
// implementation ignores `event` and returns [false].
// 
// Override this method in a subclass to allow instances to respond to
// click-through. This allows the user to click on a view in an inactive
// window, activating the view with one click, instead of clicking first to
// make the window active and then clicking the view. Most view objects refuse
// a click-through attempt, so the event simply activates the window. Many
// control objects, however, such as instances of [NSButton] and [NSSlider],
// do accept them, so the user can immediately manipulate the control without
// having to release the mouse button.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSView/acceptsFirstMouse(for:)
func (v NSView) AcceptsFirstMouse(event INSEvent) bool {
	rv := objc.Send[bool](v.ID, objc.Sel("acceptsFirstMouse:"), event)
	return rv
}

// Converts a rectangle from the view’s interior coordinate system to its
// pixel aligned backing store coordinate system.
//
// rect: A rectangle in the view’s interior coordinate system.
//
// # Return Value
// 
// A rectangle in its pixel aligned backing store coordinate system.
//
// See: https://developer.apple.com/documentation/AppKit/NSView/convertToBacking(_:)-3zors
func (v NSView) ConvertRectToBacking(rect corefoundation.CGRect) corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](v.ID, objc.Sel("convertRectToBacking:"), rect)
	return corefoundation.CGRect(rv)
}

// Converts a size from the view’s interior coordinate system to its pixel
// aligned backing store coordinate system.
//
// size: The size in the view’s interior coordinate system.
//
// # Return Value
// 
// The size in the pixel aligned coordinate system.
//
// See: https://developer.apple.com/documentation/AppKit/NSView/convertToBacking(_:)-4ra9y
func (v NSView) ConvertSizeToBacking(size corefoundation.CGSize) corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](v.ID, objc.Sel("convertSizeToBacking:"), size)
	return corefoundation.CGSize(rv)
}

// Returns EPS data that draws the region of the view within a specified
// rectangle.
//
// rect: A rectangle defining the region.
//
// # Discussion
// 
// This data can be placed on an [NSPasteboard] object, written to a file, or
// used to create an [NSImage] object.
//
// See: https://developer.apple.com/documentation/AppKit/NSView/dataWithEPS(inside:)
func (v NSView) DataWithEPSInsideRect(rect corefoundation.CGRect) foundation.INSData {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("dataWithEPSInsideRect:"), rect)
	return foundation.NSDataFromID(rv)
}

// Convert the size from the view’s interior coordinate system to the
// layer’s interior coordinate system.
//
// rect: A rectangle in the view’s interior coordinate system.
//
// # Return Value
// 
// A rectangle in the layer’s interior coordinate system.
//
// # Discussion
// 
// The layer’s space is virtual, and doesn’t take into account the
// layer’s [contentsScale] setting.
//
// [contentsScale]: https://developer.apple.com/documentation/QuartzCore/CALayer/contentsScale
//
// See: https://developer.apple.com/documentation/AppKit/NSView/convertToLayer(_:)-160pw
func (v NSView) ConvertRectToLayer(rect corefoundation.CGRect) corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](v.ID, objc.Sel("convertRectToLayer:"), rect)
	return corefoundation.CGRect(rv)
}

// Called after a contextual menu that was displayed from the receiving view
// has been closed.
//
// menu: The menu that was closed.
//
// event: The event that caused the menu to close, if there was one. If an event did
// not cause the menu to close, this value is `nil`.
//
// # Discussion
// 
// This method is called only if the contextual menu had been opened and the
// view has previously received the [WillOpenMenuWithEvent] method. When the
// view receives [DidCloseMenuWithEvent], it should reset its visual state, if
// necessary. For example, if a table view selected a row in response to a
// contextual menu being displayed, this method could deselect the row.
//
// See: https://developer.apple.com/documentation/AppKit/NSView/didCloseMenu(_:with:)
func (v NSView) DidCloseMenuWithEvent(menu INSMenu, event INSEvent) {
	objc.Send[objc.ID](v.ID, objc.Sel("didCloseMenu:withEvent:"), menu, event)
}

// Convert the size from the view’s interior coordinate system to the
// layer’s interior coordinate system.
//
// size: A size in the view’s interior coordinate system.
//
// # Return Value
// 
// A size in the layer’s interior coordinate system.
//
// # Discussion
// 
// The layer’s space is virtual, and doesn’t take into account the
// layer’s [contentsScale] setting.
// 
// The returned [NSSize] values are always forced to have positive a width and
// height.
//
// [contentsScale]: https://developer.apple.com/documentation/QuartzCore/CALayer/contentsScale
//
// See: https://developer.apple.com/documentation/AppKit/NSView/convertToLayer(_:)-2vozx
func (v NSView) ConvertSizeToLayer(size corefoundation.CGSize) corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](v.ID, objc.Sel("convertSizeToLayer:"), size)
	return corefoundation.CGSize(rv)
}

// Invalidates all cursor rectangles set up using [AddCursorRectCursor].
//
// # Discussion
// 
// You need never invoke this method directly; neither is it typically invoked
// during the invalidation of cursor rectangles. [NSWindow] automatically
// invalidates cursor rectangles in response to [InvalidateCursorRectsForView]
// and before the view’s cursor rectangles are reestablished using
// [ResetCursorRects]. This method is invoked just before the view is removed
// from a window and when the view is deallocated.
//
// See: https://developer.apple.com/documentation/AppKit/NSView/discardCursorRects()
func (v NSView) DiscardCursorRects() {
	objc.Send[objc.ID](v.ID, objc.Sel("discardCursorRects"))
}

// Convert the size from the view’s interior coordinate system to the
// layer’s interior coordinate system.
//
// point: A point in the view’s interior coordinate system.
//
// # Return Value
// 
// A point in the view’s layer interior coordinate system.
//
// # Discussion
// 
// The layer’s space is virtual, and doesn’t take into account the
// layer’s [contentsScale] setting.
//
// [contentsScale]: https://developer.apple.com/documentation/QuartzCore/CALayer/contentsScale
//
// See: https://developer.apple.com/documentation/AppKit/NSView/convertToLayer(_:)-44u7d
func (v NSView) ConvertPointToLayer(point corefoundation.CGPoint) corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](v.ID, objc.Sel("convertPointToLayer:"), point)
	return corefoundation.CGPoint(rv)
}

// Displays the view and all its subviews if any part of the view has been
// marked as needing display.
//
// # Discussion
// 
// This method invokes the [NSView] methods [lockFocus()], [DrawRect], and
// [unlockFocus()] as necessary. If the view isn’t opaque, this method backs
// up the view hierarchy to the first opaque ancestor, calculates the portion
// of the opaque ancestor covered by the view, and begins displaying from
// there.
//
// [lockFocus()]: https://developer.apple.com/documentation/AppKit/NSView/lockFocus()
// [unlockFocus()]: https://developer.apple.com/documentation/AppKit/NSView/unlockFocus()
//
// See: https://developer.apple.com/documentation/AppKit/NSView/displayIfNeeded()
func (v NSView) DisplayIfNeeded() {
	objc.Send[objc.ID](v.ID, objc.Sel("displayIfNeeded"))
}

// Overridden by subclasses to perform additional actions when subviews are
// added to the view.
//
// subview: The view that was added as a subview.
//
// # Discussion
// 
// This method is invoked by [AddSubview].
//
// See: https://developer.apple.com/documentation/AppKit/NSView/didAddSubview(_:)
func (v NSView) DidAddSubview(subview INSView) {
	objc.Send[objc.ID](v.ID, objc.Sel("didAddSubview:"), subview)
}

// Displays the view and all its subviews if possible, invoking each of the
// [NSView] methods [lockFocus()], [DrawRect], and [unlockFocus()] as
// necessary.
//
// [lockFocus()]: https://developer.apple.com/documentation/AppKit/NSView/lockFocus()
// [unlockFocus()]: https://developer.apple.com/documentation/AppKit/NSView/unlockFocus()
//
// # Discussion
// 
// If the view isn’t opaque, this method backs up the view hierarchy to the
// first opaque ancestor, calculates the portion of the opaque ancestor
// covered by the view, and begins displaying from there.
//
// See: https://developer.apple.com/documentation/AppKit/NSView/display()
func (v NSView) Display() {
	objc.Send[objc.ID](v.ID, objc.Sel("display"))
}

// Acts as [Display], but confining drawing to a rectangular region of the
// view.
//
// rect: A rectangle defining the region of the view to be redrawn; it should be
// specified in the coordinate system of the view.
//
// See: https://developer.apple.com/documentation/AppKit/NSView/display(_:)
func (v NSView) DisplayRect(rect corefoundation.CGRect) {
	objc.Send[objc.ID](v.ID, objc.Sel("displayRect:"), rect)
}

// Acts as [DisplayIfNeeded], except that this method doesn’t back up to the
// first opaque ancestor—it simply causes the view and its descendants to
// execute their drawing code.
//
// See: https://developer.apple.com/documentation/AppKit/NSView/displayIfNeededIgnoringOpacity()
func (v NSView) DisplayIfNeededIgnoringOpacity() {
	objc.Send[objc.ID](v.ID, objc.Sel("displayIfNeededIgnoringOpacity"))
}

// Acts as [DisplayIfNeeded], confining drawing to a specified region of the
// view.
//
// rect: A rectangle defining the region to be redrawn. It should be specified in
// the coordinate system of the view.
//
// See: https://developer.apple.com/documentation/AppKit/NSView/displayIfNeeded(_:)
func (v NSView) DisplayIfNeededInRect(rect corefoundation.CGRect) {
	objc.Send[objc.ID](v.ID, objc.Sel("displayIfNeededInRect:"), rect)
}

//
// See: https://developer.apple.com/documentation/AppKit/NSView/displayLink(target:selector:)
func (v NSView) DisplayLinkWithTargetSelector(target objectivec.IObject, selector objc.SEL) objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("displayLinkWithTarget:selector:"), target, selector)
	return objectivec.Object{ID: rv}
}

// Acts as [DisplayIfNeeded], but confining drawing to `aRect` and not backing
// up to the first opaque ancestor—it simply causes the view and its
// descendants to execute their drawing code.
//
// rect: A rectangle defining the region to be redrawn. It should be specified in
// the coordinate system of the view.
//
// See: https://developer.apple.com/documentation/AppKit/NSView/displayIfNeededIgnoringOpacity(_:)
func (v NSView) DisplayIfNeededInRectIgnoringOpacity(rect corefoundation.CGRect) {
	objc.Send[objc.ID](v.ID, objc.Sel("displayIfNeededInRectIgnoringOpacity:"), rect)
}

// Draws the focus ring mask for the view.
//
// # Discussion
// 
// This method provides the shape of the focus ring mask by drawing the focus
// ring mask. An implementation of this method should draw in the view’s
// interior (bounds) coordinate space, that the focus ring style has been set
// (it will be set it to [NSFocusRingPlacement.only] to capture the focus ring
// itself), and that the fill and stroke colors have been set to an arbitrary
// fully opaque color.
// 
// Subclasses that find the default behavior insufficient should only draw the
// focus ring shape.
// 
// The [NSView] implementation of this method simply fills `[self bounds]`.
//
// [NSFocusRingPlacement.only]: https://developer.apple.com/documentation/AppKit/NSFocusRingPlacement/only
//
// See: https://developer.apple.com/documentation/AppKit/NSView/drawFocusRingMask()
func (v NSView) DrawFocusRingMask() {
	objc.Send[objc.ID](v.ID, objc.Sel("drawFocusRingMask"))
}

// Displays the view but confines drawing to a specified region and does not
// back up to the first opaque ancestor—it simply causes the view and its
// descendants to execute their drawing code.
//
// rect: A rectangle defining the region of the view to be redrawn; it should be
// specified in the coordinate system of the view.
//
// See: https://developer.apple.com/documentation/AppKit/NSView/displayIgnoringOpacity(_:)
func (v NSView) DisplayRectIgnoringOpacity(rect corefoundation.CGRect) {
	objc.Send[objc.ID](v.ID, objc.Sel("displayRectIgnoringOpacity:"), rect)
}

// Returns PDF data that draws the region of the view within a specified
// rectangle.
//
// rect: A rectangle defining the region.
//
// # Discussion
// 
// This data can be placed on an [NSPasteboard] object, written to a file, or
// used to create an [NSImage] object.
//
// See: https://developer.apple.com/documentation/AppKit/NSView/dataWithPDF(inside:)
func (v NSView) DataWithPDFInsideRect(rect corefoundation.CGRect) foundation.INSData {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("dataWithPDFInsideRect:"), rect)
	return foundation.NSDataFromID(rv)
}

// Converts a rectangle from its pixel aligned backing store coordinate system
// to the view’s interior coordinate system.
//
// rect: The rectangle in the pixel backing store coordinate system.
//
// # Return Value
// 
// A rectangle in the view’s interior coordinate system.
//
// See: https://developer.apple.com/documentation/AppKit/NSView/convertFromBacking(_:)-2njpa
func (v NSView) ConvertRectFromBacking(rect corefoundation.CGRect) corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](v.ID, objc.Sel("convertRectFromBacking:"), rect)
	return corefoundation.CGRect(rv)
}

// Allows applications that use the AppKit pagination facility to draw
// additional marks on each logical page.
//
// borderSize: An [NSSize] structure that defines a logical page.
//
// # Discussion
// 
// The marks can be such things as alignment marks or a virtual sheet border
// of size `borderSize`. The default implementation doesn’t draw anything.
//
// See: https://developer.apple.com/documentation/AppKit/NSView/drawPageBorder(with:)
func (v NSView) DrawPageBorderWithSize(borderSize corefoundation.CGSize) {
	objc.Send[objc.ID](v.ID, objc.Sel("drawPageBorderWithSize:"), borderSize)
}

// This method is invoked at the end of the printing session.
//
// # Discussion
// 
// If you override this method, call the superclass implementation.
//
// See: https://developer.apple.com/documentation/AppKit/NSView/endDocument()
func (v NSView) EndDocument() {
	objc.Send[objc.ID](v.ID, objc.Sel("endDocument"))
}

// Writes the end of a conforming page.
//
// # Discussion
// 
// This method is invoked after each page is printed. It invokes
// [unlockFocus()]. This method also generates comments for the bounding box
// and page fonts, if they were specified as being at the end of the page.
//
// [unlockFocus()]: https://developer.apple.com/documentation/AppKit/NSView/unlockFocus()
//
// See: https://developer.apple.com/documentation/AppKit/NSView/endPage()
func (v NSView) EndPage() {
	objc.Send[objc.ID](v.ID, objc.Sel("endPage"))
}

// Randomly changes the frame of a view with an ambiguous layout between the
// different valid values.
//
// # Discussion
// 
// This method randomly changes the frame of a view with an ambiguous layout
// between its different valid values, causing the view to move in the
// interface. This makes it easy to visually identify what the valid frames
// are and may enable the developer to discern what constraints need to be
// added to the layout to fully specify a location for the view.
// 
// This method should only be used for debugging constraint-based layout. No
// application should ship with calls to this method as part of its operation.
//
// See: https://developer.apple.com/documentation/AppKit/NSView/exerciseAmbiguityInLayout()
func (v NSView) ExerciseAmbiguityInLayout() {
	objc.Send[objc.ID](v.ID, objc.Sel("exerciseAmbiguityInLayout"))
}

// Instructs the view to exit full screen mode.
//
// options: A dictionary of options for the mode. For possible keys, see `Full Screen
// Mode Options`.
//
// # Discussion
// 
// When the [fullScreenModeApplicationPresentationOptions] options is
// specified when [EnterFullScreenModeWithOptions] is invoked, exiting full
// screen mode will restore the previously active [PresentationOptions].
//
// [fullScreenModeApplicationPresentationOptions]: https://developer.apple.com/documentation/AppKit/NSView/FullScreenModeOptionKey/fullScreenModeApplicationPresentationOptions
//
// See: https://developer.apple.com/documentation/AppKit/NSView/exitFullScreenMode(options:)
func (v NSView) ExitFullScreenModeWithOptions(options foundation.INSDictionary) {
	objc.Send[objc.ID](v.ID, objc.Sel("exitFullScreenModeWithOptions:"), options)
}

// Overridden by subclasses to draw the view’s image within the specified
// rectangle.
//
// dirtyRect: A rectangle defining the portion of the view that requires redrawing. This
// rectangle usually represents the portion of the view that requires
// updating. When responsive scrolling is enabled, this rectangle can also
// represent a nonvisible portion of the view that AppKit wants to cache.
//
// # Discussion
// 
// Use this method to draw the specified portion of your view’s content.
// Your implementation of this method should be as fast as possible and do as
// little work as possible. The `dirtyRect` parameter helps you achieve better
// performance by specifying the portion of the view that needs to be drawn.
// You should always limit drawing to the content inside this rectangle. For
// even better performance, you can call the [GetRectsBeingDrawnCount] method
// and use the list of rectangles returned by that method to limit drawing
// even further. You can also use the [NeedsToDrawRect] method test whether
// objects in a particular rectangle need to be drawn.
// 
// The default implementation does nothing. Subclasses should override this
// method if they do custom drawing. Prior to calling this method, AppKit
// creates an appropriate drawing context and configures it for drawing to the
// view; you do not need to configure the drawing context yourself. If your
// app manages content using its layer object instead, use the [UpdateLayer]
// method to update your layer instead of overriding this method.
// 
// If your custom view is a direct [NSView] subclass, you do not need to call
// `super`. For all other views, call `super` at some point in your
// implementation so that the parent class can perform any additional drawing.
// 
// For information about how to draw in your app, see [Cocoa Drawing Guide].
//
// [Cocoa Drawing Guide]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/CocoaDrawingGuide/Introduction/Introduction.html#//apple_ref/doc/uid/TP40003290
//
// See: https://developer.apple.com/documentation/AppKit/NSView/draw(_:)
func (v NSView) DrawRect(dirtyRect corefoundation.CGRect) {
	objc.Send[objc.ID](v.ID, objc.Sel("drawRect:"), dirtyRect)
}

// Sets the view to full screen mode.
//
// screen: The screen the view should cover.
//
// options: A dictionary of options for the mode. For possible keys, see `Full Screen
// Mode Options`.
//
// # Return Value
// 
// [true] if the view was able to enter full screen mode, otherwise [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// When the [fullScreenModeApplicationPresentationOptions] is contained in the
// options dictionary, the presentation options that were in effect when this
// method is invoked are not altered, and no displays are captured.
// 
// If you do not wish to capture the screen when going to full screen mode,
// you can add [fullScreenModeApplicationPresentationOptions] to the options
// dictionary with the value returned by the [PresentationOptions].
// 
// When the [fullScreenModeApplicationPresentationOptions] options is
// specified, exiting full screen mode using [ExitFullScreenModeWithOptions]
// will restore the previously active [PresentationOptions].
// 
// # Special Considerations
// 
// In OS X v 10.5, invoking this method when the view was not in a window
// would cause an exception. In macOS 10.6 and later, you can now send this
// message to a view not in a window. For applications that must also run in
// OS X v 10.5, a simple workaround is to place the view in an offscreen
// window.
//
// [fullScreenModeApplicationPresentationOptions]: https://developer.apple.com/documentation/AppKit/NSView/FullScreenModeOptionKey/fullScreenModeApplicationPresentationOptions
//
// See: https://developer.apple.com/documentation/AppKit/NSView/enterFullScreenMode(_:withOptions:)
func (v NSView) EnterFullScreenModeWithOptions(screen INSScreen, options foundation.INSDictionary) bool {
	rv := objc.Send[bool](v.ID, objc.Sel("enterFullScreenMode:withOptions:"), screen, options)
	return rv
}

// Returns the view’s frame for a given alignment rectangle.
//
// alignmentRect: The alignment rectangle whose corresponding frame is desired.
//
// # Return Value
// 
// The frame for the specified alignment rectangle
//
// # Discussion
// 
// The constraint-based layout system uses alignment rectangles to align
// views, rather than their frame. This allows custom views to be aligned
// based on the location of their content while still having a frame that
// encompasses any ornamentation they need to draw around their content, such
// as shadows or reflections.
// 
// The default implementation returns `alignmentRect` modified by the insets
// specified by the view’s [alignmentRectInsets] method. Most custom views
// can override [alignmentRectInsets] to specify the location of their content
// within their frame. Custom views that require arbitrary transformations can
// override [AlignmentRectForFrame] and [FrameForAlignmentRect] to describe
// the location of their content. These two methods must always be inverses of
// each other.
//
// [alignmentRectInsets]: https://developer.apple.com/documentation/AppKit/NSView/alignmentRectInsets
//
// See: https://developer.apple.com/documentation/AppKit/NSView/frame(forAlignmentRect:)
func (v NSView) FrameForAlignmentRect(alignmentRect corefoundation.CGRect) corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](v.ID, objc.Sel("frameForAlignmentRect:"), alignmentRect)
	return corefoundation.CGRect(rv)
}

// Causes the view and its descendants to be redrawn to the specified graphics
// context.
//
// rect: A rectangle defining the region of the view to be redrawn. It should be
// specified in the coordinate system of the view.
//
// context: The graphics context in which drawing will occur. See the discussion below
// for more about this parameter.
//
// # Discussion
// 
// Acts as [Display], but confines drawing to `aRect`. This method initiates
// drawing with the view, even if the view is not opaque. Appropriate scaling
// factors for the view are obtained from `context`.
// 
// If the `context` parameter represents the context for the window containing
// the view, then all of the necessary transformations are applied. This
// includes the application of the view’s bounds and frame transforms along
// with any transforms it inherited from its ancestors. In this situation, the
// view is also marked as no longer needing an update for the specified
// rectangle.
// 
// If `context` specifies any other graphics context, then only the view’s
// bounds transform is applied. This means that drawing is not constrained to
// the view’s visible rectangle. It also means that any dirty rectangles are
// not cleared, since they are not being redrawn to the window.
//
// See: https://developer.apple.com/documentation/AppKit/NSView/displayIgnoringOpacity(_:in:)
func (v NSView) DisplayRectIgnoringOpacityInContext(rect corefoundation.CGRect, context INSGraphicsContext) {
	objc.Send[objc.ID](v.ID, objc.Sel("displayRectIgnoringOpacity:inContext:"), rect, context)
}

// Returns a list of rectangles indicating the newly exposed areas of the
// view.
//
// exposedRects: On return, contains the list of rectangles. The returned rectangles are in
// the coordinate space of the view.
//
// count: Contains the number of rectangles in `exposedRects`; this value may be 0
// and is guaranteed to be no more than 4.
//
// # Discussion
// 
// If your view does not support content preservation during live resizing,
// the entire area of your view is returned in the `exposedRects` parameter.
// To support content preservation, override the
// [preservesContentDuringLiveResize] property in your view and have your
// implementation return [true].
// 
// If the view decreased in both height and width, the list of returned
// rectangles will be empty. If the view increased in both height and width
// and its upper-left corner stayed anchored in the same position, the list of
// returned rectangles will contain a vertical and horizontal component
// indicating the exposed area.
//
// [preservesContentDuringLiveResize]: https://developer.apple.com/documentation/AppKit/NSView/preservesContentDuringLiveResize
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSView/getRectsExposedDuringLiveResize(_:count:)
func (v NSView) GetRectsExposedDuringLiveResizeCount(exposedRects []corefoundation.CGRect, count int) {
	objc.Send[objc.ID](v.ID, objc.Sel("getRectsExposedDuringLiveResize:count:"), objc.CArray(exposedRects), count)
}

// Returns by indirection a list of nonoverlapping rectangles that define the
// area the view is being asked to draw in [DrawRect].
//
// rects: On return, contains a list of nonoverlapping rectangles defining areas to
// be drawn in. The rectangles returned in `rects` are in the coordinate space
// of the view.
//
// count: On return, the number of rectangles in the `rects` list.
//
// # Discussion
// 
// An implementation of [DrawRect] can use this information to test whether
// objects or regions within the view intersect with the rectangles in the
// list, and thereby avoid unnecessary drawing that would be completely
// clipped away.
// 
// The [NeedsToDrawRect] method gives you a convenient way to test individual
// objects for intersection with the area being drawn in [DrawRect]. However,
// you may want to retrieve and directly inspect the rectangle list if this is
// a more efficient way to perform intersection testing.
// 
// You should send this message only from within a [DrawRect] implementation.
// The `aRect` parameter of [DrawRect] is the rectangle enclosing the returned
// list of rectangles; you can use it in an initial pass to reject objects
// that are clearly outside the area to be drawn.
//
// See: https://developer.apple.com/documentation/AppKit/NSView/getRectsBeingDrawn(_:count:)
func (v NSView) GetRectsBeingDrawnCount(rects []corefoundation.CGRect, count int) {
	objc.Send[objc.ID](v.ID, objc.Sel("getRectsBeingDrawn:count:"), objc.CArray(rects), count)
}

// Returns the farthest descendant of the view in the view hierarchy
// (including itself) that contains a specified point, or `nil` if that point
// lies completely outside the view.
//
// point: A point that is in the coordinate system of the view’s superview, not of
// the view itself.
//
// # Return Value
// 
// A view object that is the farthest descendent of `aPoint`.
//
// # Discussion
// 
// This method is used primarily by an [NSWindow] object to determine which
// view should receive a mouse-down event. You’d rarely need to invoke this
// method, but you might want to override it to have a view object hide
// mouse-down events from its subviews. This method ignores hidden views.
//
// See: https://developer.apple.com/documentation/AppKit/NSView/hitTest(_:)
func (v NSView) HitTest(point corefoundation.CGPoint) INSView {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("hitTest:"), point)
	return NSViewFromID(rv)
}

// Returns a Boolean value that indicates whether the view handles page
// boundaries.
//
// range: On return, holds the page range if [true] is returned directly. Page
// numbers are one-based—that is pages run from one to .
// //
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Return Value
// 
// [true] if the view handles page boundaries; otherwise, [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// Returns [false] if the view uses the default auto-pagination mechanism. The
// default implementation returns [false]. Override this method if your class
// handles page boundaries.
//
// [false]: https://developer.apple.com/documentation/Swift/false
//
// See: https://developer.apple.com/documentation/AppKit/NSView/knowsPageRange(_:)
func (v NSView) KnowsPageRange(range_ foundation.NSRange) bool {
	rv := objc.Send[bool](v.ID, objc.Sel("knowsPageRange:"), range_)
	return rv
}

// Returns a Boolean value that indicates whether the view is a subview of the
// specified view.
//
// view: The view to test for subview relationship within the view hierarchy.
//
// # Return Value
// 
// [true] if the view is a subview, or distant subview, of the specified view.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// The method returns [true] if the view is either an immediate or distant
// subview of `aView`.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSView/isDescendant(of:)
func (v NSView) IsDescendantOf(view INSView) bool {
	rv := objc.Send[bool](v.ID, objc.Sel("isDescendantOf:"), view)
	return rv
}

// Invalidates the view’s intrinsic content size.
//
// # Discussion
// 
// Call this when something changes in your custom view that invalidates its
// intrinsic content size. This allows the constraint-based layout system to
// take the new intrinsic content size into account in its next layout pass.
//
// See: https://developer.apple.com/documentation/AppKit/NSView/invalidateIntrinsicContentSize()
func (v NSView) InvalidateIntrinsicContentSize() {
	objc.Send[objc.ID](v.ID, objc.Sel("invalidateIntrinsicContentSize"))
}

// Returns whether a region of the view contains a specified point, accounting
// for whether the view is flipped or not.
//
// point: A point that is expressed in the view’s coordinate system. This point
// generally represents the hot spot of the mouse cursor.
//
// rect: A rectangle that is expressed in the view’s coordinate system.
//
// # Return Value
// 
// [true] if `aRect` contains `aPoint`, [false] otherwise.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// Point-in-rectangle functions generally assume that the bottom edge of a
// rectangle is outside of the rectangle boundaries, while the upper edge is
// inside the boundaries. This method views `aRect` from the point of view of
// the user—that is, this method always treats the bottom edge of the
// rectangle as the one closest to the bottom edge of the user’s screen. By
// making this adjustment, this function ensures consistent mouse-detection
// behavior from the user’s perspective.
// 
// Never use the Foundation’s [NSPointInRect(_:_:)] function as a substitute
// for this method. It doesn’t account for flipped coordinate systems.
//
// [NSPointInRect(_:_:)]: https://developer.apple.com/documentation/Foundation/NSPointInRect(_:_:)
//
// See: https://developer.apple.com/documentation/AppKit/NSView/isMousePoint(_:in:)
func (v NSView) MouseInRect(point corefoundation.CGPoint, rect corefoundation.CGRect) bool {
	rv := objc.Send[bool](v.ID, objc.Sel("mouse:inRect:"), point, rect)
	return rv
}

// Perform layout in concert with the constraint-based layout system.
//
// # Discussion
// 
// Override this method if your custom view needs to perform custom layout not
// expressible using the constraint-based layout system. In this case you are
// responsible for setting [needsLayout] to [true] when something that impacts
// your custom layout changes.
// 
// You may not invalidate any constraints as part of your layout phase, nor
// invalidate the layout of your superview or views outside of your view
// hierarchy. You also may not invoke a drawing pass as part of layout.
// 
// You must call `[super layout]` as part of your implementation.
//
// [needsLayout]: https://developer.apple.com/documentation/AppKit/NSView/needsLayout
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSView/layout()
func (v NSView) Layout() {
	objc.Send[objc.ID](v.ID, objc.Sel("layout"))
}

// Updates the layout of the receiving view and its subviews based on the
// current views and constraints.
//
// # Discussion
// 
// Before displaying a view that uses constraints-based layout the system
// invokes this method to ensure that the layout of the view and its subviews
// is up to date. This method updates the layout if needed, first invoking
// [UpdateConstraintsForSubtreeIfNeeded] to ensure that all constraints are up
// to date. This method is called automatically by the system, but may be
// invoked manually if you need to examine the most up to date layout.
// 
// Subclasses should not override this method.
//
// See: https://developer.apple.com/documentation/AppKit/NSView/layoutSubtreeIfNeeded()
func (v NSView) LayoutSubtreeIfNeeded() {
	objc.Send[objc.ID](v.ID, objc.Sel("layoutSubtreeIfNeeded"))
}

// Creates the view’s backing layer.
//
// # Return Value
// 
// The layer to use as the view’s backing layer.
//
// See: https://developer.apple.com/documentation/AppKit/NSView/makeBackingLayer()
func (v NSView) MakeBackingLayer() objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("makeBackingLayer"))
	return objectivec.Object{ID: rv}
}

// Returns a Boolean value indicating whether the specified rectangle
// intersects any part of the area that the view is being asked to draw.
//
// rect: A rectangle defining a region of the view.
//
// # Discussion
// 
// You typically send this message from within a [DrawRect] implementation. It
// gives you a convenient way to determine whether any part of a given
// graphical entity might need to be drawn. It is optimized to efficiently
// reject any rectangle that lies outside the bounding box of the area that
// the view is being asked to draw in [DrawRect].
//
// See: https://developer.apple.com/documentation/AppKit/NSView/needsToDraw(_:)
func (v NSView) NeedsToDrawRect(rect corefoundation.CGRect) bool {
	rv := objc.Send[bool](v.ID, objc.Sel("needsToDrawRect:"), rect)
	return rv
}

// Invoked to notify the view that the focus ring mask requires updating.
//
// # Discussion
// 
// It is important to note that it is only necessary for developers to invoke
// this method when some internal state change of their application, that
// AppKit can’t determine, affects the shape of the focus ring mask.
// 
// It is assumed that if the view is marked as needing display, or is resized,
// its focus ring shape is likely to have changed, and there is no need for
// clients to explicitly send this message in such cases, they are handled
// automatically.
// 
// If, however, a view is showing a focus ring around some part of its content
// (an [NSImage], perhaps), and that content changes, the client must provide
// notification by invoking this method so that [focusRingMaskBounds] and
// [DrawFocusRingMask] will be invoked to redraw the focus ring.
//
// [focusRingMaskBounds]: https://developer.apple.com/documentation/AppKit/NSView/focusRingMaskBounds
//
// See: https://developer.apple.com/documentation/AppKit/NSView/noteFocusRingMaskChanged()
func (v NSView) NoteFocusRingMaskChanged() {
	objc.Send[objc.ID](v.ID, objc.Sel("noteFocusRingMaskChanged"))
}

// Prepares the overdraw region for drawing.
//
// rect: The current overdraw region, specified in the view’s coordinate system.
// This rectangle includes the view’s visible rectangle plus any space
// surrounding the visible rectangle that represents the overdraw region.
//
// # Discussion
// 
// During responsive scrolling, AppKit calls this method before asking your
// view to draw any content in the overdraw region. You can override this
// method in your own views and use it to prepare the content that is about to
// be drawn. For example, if your app defers the creation of subviews until
// they are scrolled into view, you would use this method to create them and
// add them to your view hierarchy.
// 
// Your implementation of this method must call `super` at some point. When
// calling `super`, you can extend the overdraw rectangle by passing a
// different rectangle for the `rect` parameter. For example, if you add a
// subview whose frame falls outside the current rectangle, you can grow the
// rectangle to include the entire frame of the subview.
// 
// AppKit may call this method multiple times to build up the current overdraw
// region slowly. Each time it calls the method, it extends the overdraw
// rectangle passed in the `rect` parameter. If you pass the same rectangle to
// `super` twice in succession, AppKit stops generating additional overdraw
// content. You can use this behavior to avoid generating more overdraw
// content than makes sense for your app. If the user scrolls the content,
// AppKit resets the current overdraw region and starts asking your app for
// content again. You can also reset the current overdraw region by assigning
// a value to the [preparedContentRect] property.
//
// [preparedContentRect]: https://developer.apple.com/documentation/AppKit/NSView/preparedContentRect
//
// See: https://developer.apple.com/documentation/AppKit/NSView/prepareContent(in:)
func (v NSView) PrepareContentInRect(rect corefoundation.CGRect) {
	objc.Send[objc.ID](v.ID, objc.Sel("prepareContentInRect:"), rect)
}

// This action method opens the Print panel, and if the user chooses an option
// other than canceling, prints the view and all its subviews to the device
// specified in the Print panel.
//
// sender: The object that sent the message.
//
// See: https://developer.apple.com/documentation/AppKit/NSView/printView(_:)
func (v NSView) Print(sender objectivec.IObject) {
	objc.Send[objc.ID](v.ID, objc.Sel("print:"), sender)
}

// Returns the appropriate rectangle to use when magnifying around the
// specified point.
//
// location: The location in your view’s coordinate system around which magnification
// is centered.
//
// visibleRect: The visible portion of the view. Use this value to help determine the
// specific content group you want to target for magnification.
//
// # Return Value
// 
// The rectangle to use for magnification, specified in the view’s
// coordinate system. To get the default magnification behavior, return
// [NSZeroRect].
//
// [NSZeroRect]: https://developer.apple.com/documentation/Foundation/NSZeroRect
//
// # Discussion
// 
// AppKit calls this method when magnifying content in a scroll view. If you
// do not override this method, or if you return [NSZeroRect], the scroll view
// magnifies the view’s content around the specified point. If you override
// this method and return a custom rectangle, the scroll view adjusts the
// magnification behavior to accommodate the rectangle you provide.
// 
// Use this method to provide AppKit with rectangles for your view’s custom
// content. If your view’s content can be divided into logical groups of
// content, use the provided `location` and `visibleRect` parameters to
// determine which group is being targeted and then return the rectangle that
// fully encloses that group. For example, a view with multiple columns of
// content could return the rectangle for the targeted column. The returned
// rectangle should always fully enclose the content, regardless of whether
// that rectangle is larger than the visible rectangle.
//
// [NSZeroRect]: https://developer.apple.com/documentation/Foundation/NSZeroRect
//
// See: https://developer.apple.com/documentation/AppKit/NSView/rectForSmartMagnification(at:in:)
func (v NSView) RectForSmartMagnificationAtPointInRect(location corefoundation.CGPoint, visibleRect corefoundation.CGRect) corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](v.ID, objc.Sel("rectForSmartMagnificationAtPoint:inRect:"), location, visibleRect)
	return corefoundation.CGRect(rv)
}

// Overridden by subclasses to return a context-sensitive pop-up menu for a
// given mouse-down event.
//
// event: An object representing a mouse-down event.
//
// # Discussion
// 
// The view can use information in the mouse event, such as its location over
// a particular element of the view, to determine what kind of menu to return.
// For example, a text object might display a text-editing menu when the
// cursor lies over text and a menu for changing graphics attributes when the
// cursor lies over an embedded image.
// 
// The default implementation returns the view’s normal menu.
//
// See: https://developer.apple.com/documentation/AppKit/NSView/menu(for:)
func (v NSView) MenuForEvent(event INSEvent) INSMenu {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("menuForEvent:"), event)
	return NSMenuFromID(rv)
}

// Implemented by subclasses to determine the portion of the view to be
// printed for the specified page number.
//
// page: An integer indicating a page number. Page numbers are one-based—that is
// pages run from one to .
//
// # Return Value
// 
// A rectangle defining the region of the view to be printed for `pageNumber`.
// This method returns [NSZeroRect] if `pageNumber` is outside the view’s
// bounds.
//
// # Discussion
// 
// If the view responded [true] to an earlier [KnowsPageRange] message, this
// method is invoked for each page it specified in the out parameters of that
// message. The view is later made to display this rectangle in order to
// generate the image for this page.
// 
// If an [NSView] object responds [false] to [KnowsPageRange], this method
// isn’t invoked by the printing mechanism.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSView/rectForPage(_:)
func (v NSView) RectForPage(page int) corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](v.ID, objc.Sel("rectForPage:"), page)
	return corefoundation.CGRect(rv)
}

// Registers the pasteboard types that the view will accept as the destination
// of an image-dragging session.
//
// newTypes: An array of [Uniform Type Identifier]. See [System-Declared Uniform Type
// Identifiers] for descriptions of the pasteboard type identifiers.
// //
// [System-Declared Uniform Type Identifiers]: https://developer.apple.com/library/archive/documentation/Miscellaneous/Reference/UTIRef/Articles/System-DeclaredUniformTypeIdentifiers.html#//apple_ref/doc/uid/TP40009259
// [Uniform Type Identifier]: https://developer.apple.com/library/archive/documentation/General/Conceptual/DevPedia-CocoaCore/UniformTypeIdentifier.html#//apple_ref/doc/uid/TP40008195-CH60
//
// # Discussion
// 
// Registering an [NSView] object for dragged types automatically makes it a
// candidate destination object for a dragging session. As such, it must
// properly implement some or all of the [NSDraggingDestination] protocol
// methods. As a convenience, [NSView] provides default implementations of
// these methods. See the [NSDraggingDestination] protocol specification for
// details.
//
// See: https://developer.apple.com/documentation/AppKit/NSView/registerForDraggedTypes(_:)
func (v NSView) RegisterForDraggedTypes(newTypes []string) {
	objc.Send[objc.ID](v.ID, objc.Sel("registerForDraggedTypes:"), objectivec.StringSliceToNSArray(newTypes))
}

// Removes all tooltips assigned to the view.
//
// # Discussion
// 
// This method operates on tooltips created using either
// [AddToolTipRectOwnerUserData] or set using the [toolTip] property.
//
// [toolTip]: https://developer.apple.com/documentation/AppKit/NSView/toolTip
//
// See: https://developer.apple.com/documentation/AppKit/NSView/removeAllToolTips()
func (v NSView) RemoveAllToolTips() {
	objc.Send[objc.ID](v.ID, objc.Sel("removeAllToolTips"))
}

// Notifies a clip view’s superview that either the clip view’s bounds
// rectangle or the document view’s frame rectangle has changed, and that
// any indicators of the scroll position need to be adjusted.
//
// clipView: The [NSClipView] object whose superview is to be notified.
//
// # Discussion
// 
// [NSScrollView] implements this method to update its [NSScroller] objects.
//
// See: https://developer.apple.com/documentation/AppKit/NSView/reflectScrolledClipView(_:)
func (v NSView) ReflectScrolledClipView(clipView INSClipView) {
	objc.Send[objc.ID](v.ID, objc.Sel("reflectScrolledClipView:"), clipView)
}

// Removes the specified constraint from the view.
//
// constraint: The constraint to remove. Removing a constraint not held by the view has no
// effect.
//
// See: https://developer.apple.com/documentation/AppKit/NSView/removeConstraint(_:)
func (v NSView) RemoveConstraint(constraint INSLayoutConstraint) {
	objc.Send[objc.ID](v.ID, objc.Sel("removeConstraint:"), constraint)
}

// Unlinks the view from its superview and its window, removes it from the
// responder chain, and invalidates its cursor rectangles.
//
// # Discussion
// 
// The view is also released; if you plan to reuse it, be sure to retain it
// before sending this message and to release it as appropriate when adding it
// as a subview of another [NSView].
// 
// Calling this method removes any constraints that refer to the view you are
// removing, or that refer to any view in the subtree of the view you are
// removing.
// 
// Never invoke this method during display.
//
// See: https://developer.apple.com/documentation/AppKit/NSView/removeFromSuperview()
func (v NSView) RemoveFromSuperview() {
	objc.Send[objc.ID](v.ID, objc.Sel("removeFromSuperview"))
}

// Detaches a gesture recognizer from the view.
//
// gestureRecognizer: The gesture recognizer to remove. This parameter must not be `nil`.
//
// # Discussion
// 
// Removing a gesture recognizer also removes the strong reference to it held
// by the view.
//
// See: https://developer.apple.com/documentation/AppKit/NSView/removeGestureRecognizer(_:)
func (v NSView) RemoveGestureRecognizer(gestureRecognizer INSGestureRecognizer) {
	objc.Send[objc.ID](v.ID, objc.Sel("removeGestureRecognizer:"), gestureRecognizer)
}

// Unlinks the view from its superview and its window and removes it from the
// responder chain, but does not invalidate its cursor rectangles to cause
// redrawing.
//
// # Discussion
// 
// The view is also released; if you plan to reuse it, be sure to retain it
// before sending this message and to release it as appropriate when adding it
// as a subview of another view.
// 
// Unlike its counterpart, [RemoveFromSuperview], this method can be safely
// invoked during display.
//
// See: https://developer.apple.com/documentation/AppKit/NSView/removeFromSuperviewWithoutNeedingDisplay()
func (v NSView) RemoveFromSuperviewWithoutNeedingDisplay() {
	objc.Send[objc.ID](v.ID, objc.Sel("removeFromSuperviewWithoutNeedingDisplay"))
}

// Completely removes a cursor rectangle from the view.
//
// rect: A rectangle defining a region of the view. Must match a value previously
// specified using [AddCursorRectCursor].
//
// object: An object representing a cursor. Must match a value previously specified
// using [AddCursorRectCursor].
//
// # Discussion
// 
// You should rarely need to use this method. The [ResetCursorRects] method,
// which is called when the cursor rectangles need to be rebuilt, should
// establish only the cursor rectangles needed. If you implement
// [ResetCursorRects] in this way, you can then simply modify the state that
// [ResetCursorRects] uses to build its cursor rectangles and then invoke the
// [NSWindow] method [InvalidateCursorRectsForView].
//
// See: https://developer.apple.com/documentation/AppKit/NSView/removeCursorRect(_:cursor:)
func (v NSView) RemoveCursorRectCursor(rect corefoundation.CGRect, object INSCursor) {
	objc.Send[objc.ID](v.ID, objc.Sel("removeCursorRect:cursor:"), rect, object)
}

// Removes the tooltip identified by specified tag.
//
// tag: An integer tag that is the value returned by a previous
// [AddToolTipRectOwnerUserData] message.
//
// See: https://developer.apple.com/documentation/AppKit/NSView/removeToolTip(_:)
func (v NSView) RemoveToolTip(tag objectivec.IObject) {
	objc.Send[objc.ID](v.ID, objc.Sel("removeToolTip:"), tag)
}

// Removes the provided layout guide from the view.
//
// guide: The layout guide to be removed.
//
// # Discussion
// 
// This method removes the provided layout guide from the view’s
// [layoutGuides] array. It also sets the guide’s [OwningView] property to
// `nil`. Finally, it removes any constraints to the layout guide.
// 
// Layout guides cannot participate in Auto Layout constraints unless they are
// added by a view in the view hierarchy.
//
// [layoutGuides]: https://developer.apple.com/documentation/AppKit/NSView/layoutGuides
//
// See: https://developer.apple.com/documentation/AppKit/NSView/removeLayoutGuide(_:)
func (v NSView) RemoveLayoutGuide(guide INSLayoutGuide) {
	objc.Send[objc.ID](v.ID, objc.Sel("removeLayoutGuide:"), guide)
}

// Removes the tracking rectangle identified by a tag.
//
// tag: An integer value identifying a tracking rectangle. It was returned by a
// previously sent [AddTrackingRectOwnerUserDataAssumeInside] message.
//
// See: https://developer.apple.com/documentation/AppKit/NSView/removeTrackingRect(_:)
func (v NSView) RemoveTrackingRect(tag objectivec.IObject) {
	objc.Send[objc.ID](v.ID, objc.Sel("removeTrackingRect:"), tag)
}

// Removes a given tracking area from the view.
//
// trackingArea: The tracking area to remove from the view.
//
// See: https://developer.apple.com/documentation/AppKit/NSView/removeTrackingArea(_:)
func (v NSView) RemoveTrackingArea(trackingArea INSTrackingArea) {
	objc.Send[objc.ID](v.ID, objc.Sel("removeTrackingArea:"), trackingArea)
}

// Informs the view that the bounds size of its superview has changed.
//
// oldSize: The previous size of the superview’s bounds rectangle.
//
// # Discussion
// 
// This method is normally invoked automatically from
// [ResizeSubviewsWithOldSize].
// 
// The default implementation resizes the view according to the autoresizing
// options specified by the [autoresizingMask] property. You shouldn’t
// invoke this method directly, but you can override it to define a specific
// resizing behavior.
// 
// If you override this method and call `super` as part of your
// implementation, you should be sure to call `super` before making changes to
// the receiving view’s frame yourself.
//
// [autoresizingMask]: https://developer.apple.com/documentation/AppKit/NSView/autoresizingMask-swift.property
//
// See: https://developer.apple.com/documentation/AppKit/NSView/resize(withOldSuperviewSize:)
func (v NSView) ResizeWithOldSuperviewSize(oldSize corefoundation.CGSize) {
	objc.Send[objc.ID](v.ID, objc.Sel("resizeWithOldSuperviewSize:"), oldSize)
}

// Replaces one of the view’s subviews with another view.
//
// oldView: The view to be replaced by `newView`. May not be `nil`.
//
// newView: The view to replace `oldView`. May not be `nil`.
//
// # Discussion
// 
// This method does nothing if `oldView` is not a subview of the view.
// 
// Neither `oldView` nor `newView` may be `nil`, and the behavior is undefined
// if either of these parameters is `nil`.
// 
// This method causes `oldView` to be released; if you plan to reuse it, be
// sure to retain it before sending this message and to release it as
// appropriate when adding it as a subview of another NSView.
// 
// Calling this method also removes any constraints associated with `oldView`
// and its subtree.
//
// See: https://developer.apple.com/documentation/AppKit/NSView/replaceSubview(_:with:)
func (v NSView) ReplaceSubviewWith(oldView INSView, newView INSView) {
	objc.Send[objc.ID](v.ID, objc.Sel("replaceSubview:with:"), oldView, newView)
}

// Informs the view’s subviews that the view’s bounds rectangle size has
// changed.
//
// oldSize: The previous size of the view’s bounds rectangle.
//
// # Discussion
// 
// If the view is configured to autoresize its subviews, this method is
// automatically invoked by any method that changes the view’s frame size.
// 
// The default implementation sends [ResizeWithOldSuperviewSize] to the
// view’s subviews with `oldBoundsSize` as the argument. You shouldn’t
// invoke this method directly, but you can override it to define a specific
// resizing behavior.
//
// See: https://developer.apple.com/documentation/AppKit/NSView/resizeSubviews(withOldSize:)
func (v NSView) ResizeSubviewsWithOldSize(oldSize corefoundation.CGSize) {
	objc.Send[objc.ID](v.ID, objc.Sel("resizeSubviewsWithOldSize:"), oldSize)
}

// Rotates the view’s bounds rectangle by a specified degree value around
// the origin of the coordinate system, (0.0, 0.0).
//
// angle: A `float` value specifying the angle of rotation, in degrees.
//
// # Discussion
// 
// See the [boundsRotation] method description for more information. This
// method neither redisplays the view nor marks it as needing display. You
// must do this yourself by calling the [Display] method or setting the
// [needsDisplay] property.
// 
// This method posts an [boundsDidChangeNotification] to the default
// notification center if the view is configured to do so.
//
// [boundsDidChangeNotification]: https://developer.apple.com/documentation/AppKit/NSView/boundsDidChangeNotification
// [boundsRotation]: https://developer.apple.com/documentation/AppKit/NSView/boundsRotation
// [needsDisplay]: https://developer.apple.com/documentation/AppKit/NSView/needsDisplay
//
// See: https://developer.apple.com/documentation/AppKit/NSView/rotate(byDegrees:)
func (v NSView) RotateByAngle(angle float64) {
	objc.Send[objc.ID](v.ID, objc.Sel("rotateByAngle:"), angle)
}

// Informs the client that `aRulerView` allowed the user to move `aMarker`.
//
// # Discussion
// 
// The client can take whatever action it needs based on this message, such as
// updating the location of a tab stop in the selected paragraph, moving a
// layout guideline, or resizing a graphics element.
//
// See: https://developer.apple.com/documentation/AppKit/NSView/rulerView(_:didMove:)
func (v NSView) RulerViewDidMoveMarker(ruler INSRulerView, marker INSRulerMarker) {
	objc.Send[objc.ID](v.ID, objc.Sel("rulerView:didMoveMarker:"), ruler, marker)
}

// Overridden by subclasses to define their default cursor rectangles.
//
// # Discussion
// 
// A subclass’s implementation must invoke [AddCursorRectCursor] for each
// cursor rectangle it wants to establish. The default implementation does
// nothing.
// 
// Application code should never invoke this method directly; it’s invoked
// automatically as described in “[Responding to User Events and
// Actions].” Use the [InvalidateCursorRectsForView] method instead to
// explicitly rebuild cursor rectangles.
//
// [Responding to User Events and Actions]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/CocoaViewsGuide/SubclassingNSView/SubclassingNSView.html#//apple_ref/doc/uid/TP40002978-CH7-SW26
//
// See: https://developer.apple.com/documentation/AppKit/NSView/resetCursorRects()
func (v NSView) ResetCursorRects() {
	objc.Send[objc.ID](v.ID, objc.Sel("resetCursorRects"))
}

// Informs the client that `aRulerView` allowed the user to remove `aMarker`.
//
// # Discussion
// 
// The client can take whatever action it needs based on this message, such as
// deleting a tab stop from the paragraph style or removing a layout
// guideline.
//
// See: https://developer.apple.com/documentation/AppKit/NSView/rulerView(_:didRemove:)
func (v NSView) RulerViewDidRemoveMarker(ruler INSRulerView, marker INSRulerMarker) {
	objc.Send[objc.ID](v.ID, objc.Sel("rulerView:didRemoveMarker:"), ruler, marker)
}

// Converts a size from the view’s coordinate system to that of another
// view.
//
// size: The size (width and height) in the view’s coordinate system.
//
// view: The view that’s the target of the conversion operation. Both `view` and
// the view must belong to the same [NSWindow] object, and that window must
// not be `nil`. If `view` is `nil`, this method converts to window
// coordinates instead.
//
// # Return Value
// 
// The converted size, as an [NSSize] structure.
//
// [NSSize]: https://developer.apple.com/documentation/Foundation/NSSize
//
// # Discussion
// 
// The returned [NSSize] values are always forced to have positive a width and
// height.
//
// See: https://developer.apple.com/documentation/AppKit/NSView/convert(_:to:)-5nptx
func (v NSView) ConvertSizeToView(size corefoundation.CGSize, view INSView) corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](v.ID, objc.Sel("convertSize:toView:"), size, view)
	return corefoundation.CGSize(rv)
}

// Informs the client that the user has pressed the mouse button while the
// cursor is in the ruler area of `aRulerView`.
//
// # Discussion
// 
// `theEvent` is the mouse-down event that triggered the message. The client
// view can implement this method to perform an action such as adding a new
// marker using [NSRulerMarkerClientViewDelegation] or adding layout
// guidelines.
//
// See: https://developer.apple.com/documentation/AppKit/NSView/rulerView(_:handleMouseDownWith:)
func (v NSView) RulerViewHandleMouseDown(ruler INSRulerView, event INSEvent) {
	objc.Send[objc.ID](v.ID, objc.Sel("rulerView:handleMouseDown:"), ruler, event)
}

//
// See: https://developer.apple.com/documentation/AppKit/NSView/rulerView(_:pointForLocation:)
func (v NSView) RulerViewPointForLocation(ruler INSRulerView, point float64) corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](v.ID, objc.Sel("rulerView:pointForLocation:"), ruler, point)
	return corefoundation.CGPoint(rv)
}

// Requests permission for `aRulerView` to add `aMarker`, an NSRulerMarker
// being dragged onto the ruler by the user.
//
// # Discussion
// 
// If the client returns [true] the ruler view accepts the new marker and
// begins tracking its movement; if the client returns [false] the ruler view
// refuses the new marker.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSView/rulerView(_:shouldAdd:)
func (v NSView) RulerViewShouldAddMarker(ruler INSRulerView, marker INSRulerMarker) bool {
	rv := objc.Send[bool](v.ID, objc.Sel("rulerView:shouldAddMarker:"), ruler, marker)
	return rv
}

// Requests permission for `aRulerView` to move `aMarker`.
//
// # Discussion
// 
// If the client returns [true] the ruler view allows the user to move the
// marker; if the client returns [false] the marker doesn’t move.
// 
// The user’s ability to move a marker is typically set on the marker
// itself, using NSRulerMarker’s [Movable] method. You should use this
// client view method only when the marker’s movability can vary depending
// on a variable condition (for example, if graphic items can be locked down
// to prevent them from being inadvertently moved).
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSView/rulerView(_:shouldMove:)
func (v NSView) RulerViewShouldMoveMarker(ruler INSRulerView, marker INSRulerMarker) bool {
	rv := objc.Send[bool](v.ID, objc.Sel("rulerView:shouldMoveMarker:"), ruler, marker)
	return rv
}

// Requests permission for `aRulerView` to remove `aMarker`.
//
// # Discussion
// 
// If the client returns [true] the ruler view allows the user to remove the
// marker; if the client returns [false] the marker is kept pinned to the
// ruler’s baseline. This message is sent as many times as needed while the
// user drags the marker.
// 
// The user’s ability to remove a marker is typically set on the marker
// itself, using NSRulerMarker’s [Removable] method. You should use this
// client view method only when the marker’s removability can vary while the
// user drags it (for example, if the user must press the Shift key to remove
// a marker).
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSView/rulerView(_:shouldRemove:)
func (v NSView) RulerViewShouldRemoveMarker(ruler INSRulerView, marker INSRulerMarker) bool {
	rv := objc.Send[bool](v.ID, objc.Sel("rulerView:shouldRemoveMarker:"), ruler, marker)
	return rv
}

// Informs the client that `aRulerView` will add the new NSRulerMarker,
// `aMarker`.
//
// # Discussion
// 
// `location` is the marker’s tentative new location, expressed in the
// client view’s coordinate system. The value returned by the client view is
// actually used; the client can simply return `location` unchanged or adjust
// it as needed. For example, it may snap the location to a grid. This message
// is sent repeatedly to the client as the user drags the marker.
//
// See: https://developer.apple.com/documentation/AppKit/NSView/rulerView(_:willAdd:atLocation:)
func (v NSView) RulerViewWillAddMarkerAtLocation(ruler INSRulerView, marker INSRulerMarker, location float64) float64 {
	rv := objc.Send[float64](v.ID, objc.Sel("rulerView:willAddMarker:atLocation:"), ruler, marker, location)
	return rv
}

// Removes the specified constraints from the view.
//
// constraints: The constraints to remove.
//
// See: https://developer.apple.com/documentation/AppKit/NSView/removeConstraints(_:)
func (v NSView) RemoveConstraints(constraints []NSLayoutConstraint) {
	objc.Send[objc.ID](v.ID, objc.Sel("removeConstraints:"), objectivec.IObjectSliceToNSArray(constraints))
}

// Informs the client that `aRulerView` will move `aMarker`, an NSRulerMarker
// already on the ruler view.
//
// # Discussion
// 
// `location` is the marker’s tentative new location, expressed in the
// client view’s coordinate system. The value returned by the client view is
// actually used; the client can simply return `location` unchanged or adjust
// it as needed. For example, it may snap the location to a grid. This message
// is sent repeatedly to the client as the user drags the marker.
//
// See: https://developer.apple.com/documentation/AppKit/NSView/rulerView(_:willMove:toLocation:)
func (v NSView) RulerViewWillMoveMarkerToLocation(ruler INSRulerView, marker INSRulerMarker, location float64) float64 {
	rv := objc.Send[float64](v.ID, objc.Sel("rulerView:willMoveMarker:toLocation:"), ruler, marker, location)
	return rv
}

//
// See: https://developer.apple.com/documentation/AppKit/NSView/rulerView(_:locationFor:)
func (v NSView) RulerViewLocationForPoint(ruler INSRulerView, point corefoundation.CGPoint) float64 {
	rv := objc.Send[float64](v.ID, objc.Sel("rulerView:locationForPoint:"), ruler, point)
	return rv
}

// Informs the client view that `aRulerView` is about to be appropriated by
// `newClient`.
//
// # Discussion
// 
// The client view can use this opportunity to clear any cached information
// related to the ruler.
//
// See: https://developer.apple.com/documentation/AppKit/NSView/rulerView(_:willSetClientView:)
func (v NSView) RulerViewWillSetClientView(ruler INSRulerView, newClient INSView) {
	objc.Send[objc.ID](v.ID, objc.Sel("rulerView:willSetClientView:"), ruler, newClient)
}

// Scales the view’s coordinate system so that the unit square scales to the
// specified dimensions.
//
// newUnitSize: An [NSSize] structure specifying the new unit size.
//
// # Discussion
// 
// For example, a `newUnitSize` of (0.5, 1.0) causes the view’s horizontal
// coordinates to be halved, in turn doubling the width of its bounds
// rectangle. Note that scaling is performed from the origin of the coordinate
// system, (0.0, 0.0), not the origin of the bounds rectangle; as a result,
// both the origin and size of the bounds rectangle are changed. The frame
// rectangle remains unchanged.
// 
// This method does not redisplay the view or mark it as needing display. You
// must do this yourself by calling the [Display] method or setting the
// [needsDisplay] property.
// 
// This method posts an [boundsDidChangeNotification] to the default
// notification center if the view is configured to do so.
//
// [boundsDidChangeNotification]: https://developer.apple.com/documentation/AppKit/NSView/boundsDidChangeNotification
// [needsDisplay]: https://developer.apple.com/documentation/AppKit/NSView/needsDisplay
//
// See: https://developer.apple.com/documentation/AppKit/NSView/scaleUnitSquare(to:)
func (v NSView) ScaleUnitSquareToSize(newUnitSize corefoundation.CGSize) {
	objc.Send[objc.ID](v.ID, objc.Sel("scaleUnitSquareToSize:"), newUnitSize)
}

// Scrolls the view’s closest ancestor [NSClipView] object the minimum
// distance needed so a specified region of the view becomes visible in the
// clip view.
//
// rect: The rectangle to be made visible in the clip view.
//
// # Discussion
// 
// [true] if any scrolling is performed; otherwise returns [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSView/scrollToVisible(_:)
func (v NSView) ScrollRectToVisible(rect corefoundation.CGRect) bool {
	rv := objc.Send[bool](v.ID, objc.Sel("scrollRectToVisible:"), rect)
	return rv
}

// Sets the origin of the view’s bounds rectangle to a specified point.
//
// newOrigin: A point specifying the new bounds origin of the view.
//
// # Discussion
// 
// In setting the new bounds origin, this method effectively shifts the
// view’s coordinate system so `newOrigin` lies at the origin of the
// view’s frame rectangle. It neither redisplays the view nor marks it as
// needing display. Set the [needsDisplay] property to [true] when you want
// the view to be redisplayed.
// 
// This method posts an [boundsDidChangeNotification] to the default
// notification center if the view is configured to do so.
// 
// After calling this method, [NSView] creates an internal transform (or
// appends these changes to an existing internal transform) to convert from
// frame coordinates to bounds coordinates in your view. As long as the
// width-to-height ratio of the two coordinate systems remains the same, your
// content appears normal. If the ratios differ, your content may appear
// skewed.
//
// [boundsDidChangeNotification]: https://developer.apple.com/documentation/AppKit/NSView/boundsDidChangeNotification
// [needsDisplay]: https://developer.apple.com/documentation/AppKit/NSView/needsDisplay
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSView/setBoundsOrigin(_:)
func (v NSView) SetBoundsOrigin(newOrigin corefoundation.CGPoint) {
	objc.Send[objc.ID](v.ID, objc.Sel("setBoundsOrigin:"), newOrigin)
}

// Sets the size of the view’s bounds rectangle to specified dimensions,
// inversely scaling its coordinate system relative to its frame rectangle.
//
// newSize: An [NSSize] structure specifying the new width and height of the view’s
// bounds rectangle.
//
// # Discussion
// 
// For example, a view object with a frame size of (100.0, 100.0) and a bounds
// size of (200.0, 100.0) draws half as wide along the x axis. This method
// neither redisplays the view nor marks it as needing display. Set the
// [needsDisplay] property to [true] when you want the view to be redisplayed.
// 
// This method posts an [boundsDidChangeNotification] to the default
// notification center if the view is configured to do so.
// 
// After calling this method, [NSView] creates an internal transform (or
// appends these changes to an existing internal transform) to convert from
// frame coordinates to bounds coordinates in your view. As long as the
// width-to-height ratio of the two coordinate systems remains the same, your
// content appears normal. If the ratios differ, your content may appear
// skewed.
//
// [boundsDidChangeNotification]: https://developer.apple.com/documentation/AppKit/NSView/boundsDidChangeNotification
// [needsDisplay]: https://developer.apple.com/documentation/AppKit/NSView/needsDisplay
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSView/setBoundsSize(_:)
func (v NSView) SetBoundsSize(newSize corefoundation.CGSize) {
	objc.Send[objc.ID](v.ID, objc.Sel("setBoundsSize:"), newSize)
}

// Sets the priority with which a view resists being made smaller than its
// intrinsic size.
//
// priority: The new priority.
//
// orientation: The orientation for which the compression resistance priority should be
// set.
//
// # Discussion
// 
// Custom views should set default values for both orientations on creation,
// based on their content, typically to [defaultLow] or [defaultHigh]. When
// creating user interfaces, the layout designer can modify these priorities
// for specific views when the overall layout design requires different
// tradeoffs than the natural priorities of the views being used in the
// interface.
// 
// Subclasses should not override this method.
//
// [defaultHigh]: https://developer.apple.com/documentation/AppKit/NSLayoutConstraint/Priority-swift.struct/defaultHigh
// [defaultLow]: https://developer.apple.com/documentation/AppKit/NSLayoutConstraint/Priority-swift.struct/defaultLow
//
// See: https://developer.apple.com/documentation/AppKit/NSView/setContentCompressionResistancePriority(_:for:)
func (v NSView) SetContentCompressionResistancePriorityForOrientation(priority NSLayoutPriority, orientation NSLayoutConstraintOrientation) {
	objc.Send[objc.ID](v.ID, objc.Sel("setContentCompressionResistancePriority:forOrientation:"), priority, orientation)
}

// Sets the origin of the view’s frame rectangle to the specified point,
// effectively repositioning it within its superview.
//
// newOrigin: The point that is the new origin of the view’s frame.
//
// # Discussion
// 
// Changing the frame does not mark the view as needing to be displayed. Set
// the [needsDisplay] property to [true] when you want the view to be
// redisplayed.
// 
// Changing the frame origin results in the posting of an
// [frameDidChangeNotification] to the default notification center if the view
// is configured to do so.
//
// [frameDidChangeNotification]: https://developer.apple.com/documentation/AppKit/NSView/frameDidChangeNotification
// [needsDisplay]: https://developer.apple.com/documentation/AppKit/NSView/needsDisplay
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSView/setFrameOrigin(_:)
func (v NSView) SetFrameOrigin(newOrigin corefoundation.CGPoint) {
	objc.Send[objc.ID](v.ID, objc.Sel("setFrameOrigin:"), newOrigin)
}

// Informs the client that `aRulerView` allowed the user to add `aMarker`.
//
// # Discussion
// 
// The client can take whatever action it needs based on this message, such as
// adding a new tab stop to the selected paragraph or creating a layout
// guideline.
//
// See: https://developer.apple.com/documentation/AppKit/NSView/rulerView(_:didAdd:)
func (v NSView) RulerViewDidAddMarker(ruler INSRulerView, marker INSRulerMarker) {
	objc.Send[objc.ID](v.ID, objc.Sel("rulerView:didAddMarker:"), ruler, marker)
}

// Invalidates the area around the focus ring.
//
// rect: The rectangle of the control or cell defining the area around the focus
// ring. `rect` will be expanded to include the focus ring for invalidation.
//
// See: https://developer.apple.com/documentation/AppKit/NSView/setKeyboardFocusRingNeedsDisplay(_:)
func (v NSView) SetKeyboardFocusRingNeedsDisplayInRect(rect corefoundation.CGRect) {
	objc.Send[objc.ID](v.ID, objc.Sel("setKeyboardFocusRingNeedsDisplayInRect:"), rect)
}

// Notifies the superview of a clip view that the clip view needs to reset the
// origin of its bounds rectangle.
//
// clipView: The [NSClipView] object whose superview is to be notified.
//
// point: A point that specifies the new origin of the clip view’s bounds
// rectangle.
//
// # Discussion
// 
// The superview of `clipView` should then send a [ScrollToPoint] message to
// `clipView` with `point` as the argument. This mechanism is provided so the
// [NSClipView] object’s superview can coordinate scrolling of multiple
// tiled clip views.
//
// See: https://developer.apple.com/documentation/AppKit/NSView/scroll(_:to:)
func (v NSView) ScrollClipViewToPoint(clipView INSClipView, point corefoundation.CGPoint) {
	objc.Send[objc.ID](v.ID, objc.Sel("scrollClipView:toPoint:"), clipView, point)
}

// Sets the priority with which a view resists being made larger than its
// intrinsic size.
//
// priority: The new priority.
//
// orientation: The orientation for which the content hugging priority should be set.
//
// # Discussion
// 
// Custom views should set default values for both orientations on creation,
// based on their content, typically to [defaultLow] or [defaultHigh]. When
// creating user interfaces, the layout designer can modify these priorities
// for specific views when the overall layout design requires different
// tradeoffs than the natural priorities of the views being used in the
// interface.
// 
// Subclasses should not override this method.
//
// [defaultHigh]: https://developer.apple.com/documentation/AppKit/NSLayoutConstraint/Priority-swift.struct/defaultHigh
// [defaultLow]: https://developer.apple.com/documentation/AppKit/NSLayoutConstraint/Priority-swift.struct/defaultLow
//
// See: https://developer.apple.com/documentation/AppKit/NSView/setContentHuggingPriority(_:for:)
func (v NSView) SetContentHuggingPriorityForOrientation(priority NSLayoutPriority, orientation NSLayoutConstraintOrientation) {
	objc.Send[objc.ID](v.ID, objc.Sel("setContentHuggingPriority:forOrientation:"), priority, orientation)
}

// Shows a window displaying the definition of the attributed string at the
// specified point.
//
// attrString: The attributed string for which to show the definition. If the view is an
// instance of NSTextView, the `attrString` can be `nil`, in which case the
// text view will automatically supply values suitable for displaying
// definitions for the specified range within its text content.
//
// textBaselineOrigin: Specifies the baseline origin of `attrString` in the view’s coordinate
// system.
//
// # Discussion
// 
// Shows a window that displays the definition (or other subject depending on
// available dictionaries) of the specified attributed string.
// 
// This method can be used for implementing the same functionality as the
// [NSTextView] “Look Up in Dictionary” contextual menu on a custom view.
//
// See: https://developer.apple.com/documentation/AppKit/NSView/showDefinition(for:at:)
func (v NSView) ShowDefinitionForAttributedStringAtPoint(attrString foundation.NSAttributedString, textBaselineOrigin corefoundation.CGPoint) {
	objc.Send[objc.ID](v.ID, objc.Sel("showDefinitionForAttributedString:atPoint:"), attrString, textBaselineOrigin)
}

// Convert the size from the layer’s interior coordinate system to the
// view’s interior coordinate system.
//
// size: A size in the layer’s interior coordinate system.
//
// # Return Value
// 
// A size in the view’s interior coordinate system.
//
// # Discussion
// 
// The layer’s space is virtual, and doesn’t take into account the
// layer’s [contentsScale] setting.
// 
// The returned [NSSize] values are always forced to have positive a width and
// height.
//
// [contentsScale]: https://developer.apple.com/documentation/QuartzCore/CALayer/contentsScale
//
// See: https://developer.apple.com/documentation/AppKit/NSView/convertFromLayer(_:)-3usqp
func (v NSView) ConvertSizeFromLayer(size corefoundation.CGSize) corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](v.ID, objc.Sel("convertSizeFromLayer:"), size)
	return corefoundation.CGSize(rv)
}

// Scrolls the view’s closest ancestor [NSClipView] object so a point in the
// view lies at the origin of the clip view’s bounds rectangle.
//
// point: The point in the view to scroll to.
//
// See: https://developer.apple.com/documentation/AppKit/NSView/scroll(_:)
func (v NSView) ScrollPoint(point corefoundation.CGPoint) {
	objc.Send[objc.ID](v.ID, objc.Sel("scrollPoint:"), point)
}

// Marks the region of the view within the specified rectangle as needing
// display, increasing the view’s existing invalid region to include it.
//
// invalidRect: The rectangular region of the view to mark as invalid; it should be
// specified in the coordinate system of the view.
//
// # Discussion
// 
// A later `displayIfNeeded` method will then perform drawing only within the
// invalid region. View objects marked as needing display are automatically
// redisplayed on each pass through the application’s event loop. (View
// objects that need to redisplay before the event loop comes around can of
// course immediately be sent the appropriate `display` method.)
//
// See: https://developer.apple.com/documentation/AppKit/NSView/setNeedsDisplay(_:)
func (v NSView) SetNeedsDisplayInRect(invalidRect corefoundation.CGRect) {
	objc.Send[objc.ID](v.ID, objc.Sel("setNeedsDisplayInRect:"), invalidRect)
}

// Translates the view’s coordinate system so that its origin moves to a new
// location.
//
// translation: A point that specifies the new origin.
//
// # Discussion
// 
// In the process, the origin of the view’s bounds rectangle is shifted by
// (`–translation.X()`, `–translation.Y()`). This method neither
// redisplays the view nor marks it as needing display. You must do this
// yourself by calling the [Display] method or setting the [needsDisplay]
// property.
// 
// Note the difference between this method and setting the bounds origin.
// Translation effectively moves the image inside the bounds rectangle, while
// setting the bounds origin effectively moves the rectangle over the image.
// The two are in a sense inverse, although translation is cumulative, and
// setting the bounds origin is absolute.
// 
// This method posts an [boundsDidChangeNotification] to the default
// notification center if the view is configured to do so.
//
// [boundsDidChangeNotification]: https://developer.apple.com/documentation/AppKit/NSView/boundsDidChangeNotification
// [needsDisplay]: https://developer.apple.com/documentation/AppKit/NSView/needsDisplay
//
// See: https://developer.apple.com/documentation/AppKit/NSView/translateOrigin(to:)
func (v NSView) TranslateOriginToPoint(translation corefoundation.CGPoint) {
	objc.Send[objc.ID](v.ID, objc.Sel("translateOriginToPoint:"), translation)
}

// Orders the view’s immediate subviews using the specified comparator
// function.
//
// compare: A pointer to the comparator function. This function must take as arguments
// two subviews to be ordered and contextual data (supplied in `context` which
// may be arbitrary data used to help in the comparison. The comparator
// function should return [NSOrderedAscending] if the first subview should be
// ordered lower, [NSOrderedDescending] if the second subview should be
// ordered lower, and [NSOrderedSame] if their ordering isn’t important.
//
// context: Arbitrary data that might help the comparator function `compare` in its
// decisions.
//
// See: https://developer.apple.com/documentation/AppKit/NSView/sortSubviews(_:context:)
func (v NSView) SortSubviewsUsingFunctionContext(compare objectivec.IObject, context unsafe.Pointer) {
	objc.Send[objc.ID](v.ID, objc.Sel("sortSubviewsUsingFunction:context:"), compare, context)
}

// Unregisters the view as a possible destination in a dragging session.
//
// See: https://developer.apple.com/documentation/AppKit/NSView/unregisterDraggedTypes()
func (v NSView) UnregisterDraggedTypes() {
	objc.Send[objc.ID](v.ID, objc.Sel("unregisterDraggedTypes"))
}

// Updates the constraints for the receiving view and its subviews.
//
// # Discussion
// 
// Whenever a new layout pass is triggered for a view, the system invokes this
// method to ensure that any constraints for the view and its subviews are
// updated with information from the current view hierarchy and its
// constraints. This method is called automatically by the system, but may be
// invoked manually if you need to examine the most up to date constraints.
// 
// Subclasses should not override this method.
//
// See: https://developer.apple.com/documentation/AppKit/NSView/updateConstraintsForSubtreeIfNeeded()
func (v NSView) UpdateConstraintsForSubtreeIfNeeded() {
	objc.Send[objc.ID](v.ID, objc.Sel("updateConstraintsForSubtreeIfNeeded"))
}

// Update constraints for the view.
//
// # Discussion
// 
// Override this method to optimize changes to your constraints.
// 
// To schedule a change, set the view’s [needsUpdateConstraints] property to
// [true]. The system then calls your implementation of [UpdateConstraints]
// before the layout occurs. This lets you verify that all necessary
// constraints for your content are in place at a time when your custom
// view’s properties are not changing.
// 
// Your implementation must be as efficient as possible. Do not deactivate all
// your constraints, then reactivate the ones you need. Instead, your app must
// have some way of tracking your constraints, and validating them during each
// update pass. Only change items that need to be changed. During each update
// pass, you must ensure that you have the appropriate constraints for the
// app’s current state.
// 
// Do not set the [needsUpdateConstraints] property inside your
// implementation. Setting [needsUpdateConstraints] to [true] schedules
// another update pass, creating a feedback loop.
//
// [needsUpdateConstraints]: https://developer.apple.com/documentation/AppKit/NSView/needsUpdateConstraints
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSView/updateConstraints()
func (v NSView) UpdateConstraints() {
	objc.Send[objc.ID](v.ID, objc.Sel("updateConstraints"))
}

// Translates the display rectangles by the specified delta.
//
// clipRect: A rectangle defining the region of the view, typically the view’s bounds.
//
// delta: A NSSize structure that specifies an offset from aRect’s origin.
//
// # Discussion
// 
// This method performs the shifting of dirty rectangles that an equivalent
// [scroll(_:by:)] operation would cause, without performing the actual scroll
// operation. It is only useful in very rare cases where a view implements its
// own low-level scrolling mechanics.
// 
// This method:
// 
// - Collects the receiving view’s dirty rectangles. - Clears all dirty
// rectangles in the intersection of `clipRect` and the view’s bounds. -
// Shifts the retrieved rectangles by the `delta` offset. - Clips the result
// to the intersection of `clipRect` and the view’s bounds - Marks the
// resultant rectangles as needing display.
// 
// The developer must ensure that `clipRect` and `delta` are pixel-aligned in
// order to guarantee correct drawing. See [Transforming View Coordinates To
// and From Base Space] for a description of how to pixel-align view
// coordinates.
//
// [Transforming View Coordinates To and From Base Space]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/CocoaViewsGuide/WorkingWithAViewHierarchy/WorkingWithAViewHierarchy.html#//apple_ref/doc/uid/TP40002978-CH4-SW25
// [scroll(_:by:)]: https://developer.apple.com/documentation/AppKit/NSView/scroll(_:by:)
//
// See: https://developer.apple.com/documentation/AppKit/NSView/translateRectsNeedingDisplay(in:by:)
func (v NSView) TranslateRectsNeedingDisplayInRectBy(clipRect corefoundation.CGRect, delta corefoundation.CGSize) {
	objc.Send[objc.ID](v.ID, objc.Sel("translateRectsNeedingDisplayInRect:by:"), clipRect, delta)
}

// Invoked automatically when the view’s geometry changes such that its
// tracking areas need to be recalculated.
//
// # Discussion
// 
// You should override this method to remove out of date tracking areas and
// add recomputed tracking areas; your implementation should call `super`.
//
// See: https://developer.apple.com/documentation/AppKit/NSView/updateTrackingAreas()
func (v NSView) UpdateTrackingAreas() {
	objc.Send[objc.ID](v.ID, objc.Sel("updateTrackingAreas"))
}

// Responds when the view’s backing store properties change.
//
// # Discussion
// 
// The view gets this message when the backing store scale or color space
// changes. Provide an implementation if you need to swap assets or make other
// adjustments when a view’s backing store properties change.
//
// See: https://developer.apple.com/documentation/AppKit/NSView/viewDidChangeBackingProperties()
func (v NSView) ViewDidChangeBackingProperties() {
	objc.Send[objc.ID](v.ID, objc.Sel("viewDidChangeBackingProperties"))
}

// Informs the view that its effective appearance changed.
//
// See: https://developer.apple.com/documentation/AppKit/NSView/viewDidChangeEffectiveAppearance()
func (v NSView) ViewDidChangeEffectiveAppearance() {
	objc.Send[objc.ID](v.ID, objc.Sel("viewDidChangeEffectiveAppearance"))
}

// Shows a window displaying the definition of the specified range of the
// attributed string.
//
// attrString: The attributed string for which to show the definition. If the view is an
// instance of NSTextView, the `attrString` value can be `nil`, in which case
// the text view will automatically supply values suitable for displaying
// definitions for the specified range within its text content.
//
// targetRange: The range of the attributed string to define. You can pass a zero-length
// range and the appropriate range will be auto-detected around the range’s
// offset. That’s the recommended approach when there is no selection.
//
// options: An optional dictionary that specifies how the definition is displayed. See
// `NSDefinition Presentation Constants` for the key and it’s possible
// values.
//
// originProvider: The originProvider block object should return the baseline origin for the
// first character at the adjusted range.
// 
// If the view is an instance of NSTextView, the originProvider can be [NULL],
// in which case the text view will automatically supply values suitable for
// displaying definitions for the specified range within its text content.
// 
// The block object takes a single argument:
// 
// adjustedRange: The adjusted range.
// 
// The block object returns an [NSPoint] to be used as the baseline origin of
// the first character, in the view’s view coordinate system.
//
// # Discussion
// 
// This method does not cause scrolling, so clients should perform any
// necessary scrolling before calling this method.
//
// See: https://developer.apple.com/documentation/AppKit/NSView/showDefinition(for:range:options:baselineOriginProvider:)
func (v NSView) ShowDefinitionForAttributedStringRangeOptionsBaselineOriginProvider(attrString foundation.NSAttributedString, targetRange foundation.NSRange, options foundation.INSDictionary, originProvider RangeHandler) {
		_block3, _cleanup3 := NewRangeBlock(originProvider)
	defer _cleanup3()
		objc.Send[objc.ID](v.ID, objc.Sel("showDefinitionForAttributedString:range:options:baselineOriginProvider:"), attrString, targetRange, options, _block3)
}

// Updates the view’s content by modifying its underlying layer.
//
// # Discussion
// 
// You use this method to optimize the rendering of your view in situations
// where you can represent your views contents entirely using a layer object.
// If your view’s [wantsUpdateLayer] property is [true], the view calls this
// method instead of [DrawRect] during the view update cycle. Custom views can
// override this method and use it to modify the properties of the underlying
// layer object. Modifying layer properties is a much more efficient way to
// update your view than is redrawing its content each time something changes.
// 
// When you want to update the contents of your layer, mark the view as dirty
// by setting its [needsDisplay] property to [true]. Doing so adds the view to
// the list of views that need to be refreshed during the next update cycle.
// During that update cycle, this method is called if the [wantsUpdateLayer]
// property is still [true].
// 
// Your implementation of this method should not call `super`.
//
// [needsDisplay]: https://developer.apple.com/documentation/AppKit/NSView/needsDisplay
// [true]: https://developer.apple.com/documentation/Swift/true
// [wantsUpdateLayer]: https://developer.apple.com/documentation/AppKit/NSView/wantsUpdateLayer
//
// See: https://developer.apple.com/documentation/AppKit/NSView/updateLayer()
func (v NSView) UpdateLayer() {
	objc.Send[objc.ID](v.ID, objc.Sel("updateLayer"))
}

// Invoked when the view is hidden, either directly, or in response to an
// ancestor being hidden.
//
// # Discussion
// 
// The view receives this message when its [isHiddenOrHasHiddenAncestor]
// property changes from [false] to [true]. This happens when the view or an
// ancestor is marked as hidden, or when the view or an ancestor is inserted
// into a new view hierarchy.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [isHiddenOrHasHiddenAncestor]: https://developer.apple.com/documentation/AppKit/NSView/isHiddenOrHasHiddenAncestor
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSView/viewDidHide()
func (v NSView) ViewDidHide() {
	objc.Send[objc.ID](v.ID, objc.Sel("viewDidHide"))
}

// Informs the view that its superview has changed (possibly to `nil`).
//
// # Discussion
// 
// The default implementation does nothing; subclasses can override this
// method to perform whatever actions are necessary.
//
// See: https://developer.apple.com/documentation/AppKit/NSView/viewDidMoveToSuperview()
func (v NSView) ViewDidMoveToSuperview() {
	objc.Send[objc.ID](v.ID, objc.Sel("viewDidMoveToSuperview"))
}

// Invoked when the view is unhidden, either directly, or in response to an
// ancestor being unhidden
//
// # Discussion
// 
// The view receives this message when its `isHiddenOrHasHiddenAncestor` state
// goes from [true] to [false]. This can happen when the view or an ancestor
// is marked as not hidden, or when the view or an ancestor is removed from
// its containing view hierarchy.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSView/viewDidUnhide()
func (v NSView) ViewDidUnhide() {
	objc.Send[objc.ID](v.ID, objc.Sel("viewDidUnhide"))
}

// Informs the view that it has been added to a new view hierarchy.
//
// # Discussion
// 
// The default implementation does nothing; subclasses can override this
// method to perform whatever actions are necessary.
// 
// If the view’s [window] property is `nil`, that result signifies that the
// view was removed from its window and does not currently reside in any
// window.
//
// [window]: https://developer.apple.com/documentation/AppKit/NSView/window
//
// See: https://developer.apple.com/documentation/AppKit/NSView/viewDidMoveToWindow()
func (v NSView) ViewDidMoveToWindow() {
	objc.Send[objc.ID](v.ID, objc.Sel("viewDidMoveToWindow"))
}

// Informs the view of the end of a live resize—the user has finished
// resizing the view.
//
// # Discussion
// 
// In the simple case, a view is sent [ViewWillStartLiveResize] before the
// first resize operation on the containing window and [ViewDidEndLiveResize]
// after the last resize operation. A view that is repeatedly added and
// removed from a window during live resize will receive only one
// [ViewWillStartLiveResize] (on the first time it is added to the window) and
// one [ViewDidEndLiveResize] (when the window has completed the live resize
// operation). This allows a superview such as [NSBrowser] object to add and
// remove its [NSMatrix] subviews during live resize without the [NSMatrix]
// receiving multiple calls to these methods.
// 
// A view might allocate data structures to cache-drawing information in
// [ViewWillStartLiveResize] and should clean up these data structures in
// [ViewDidEndLiveResize]. In addition, a view that does optimized drawing
// during live resize might need to do full drawing after
// [ViewDidEndLiveResize]. However, you should not assume that a view has a
// drawing context in [ViewDidEndLiveResize] (since it may have been removed
// from the window during live resize). A view that needs to redraw itself
// after live resize should set its [needsDisplay] property to [true] in
// [ViewDidEndLiveResize].
// 
// A view subclass should call `super` from these methods.
//
// [needsDisplay]: https://developer.apple.com/documentation/AppKit/NSView/needsDisplay
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSView/viewDidEndLiveResize()
func (v NSView) ViewDidEndLiveResize() {
	objc.Send[objc.ID](v.ID, objc.Sel("viewDidEndLiveResize"))
}

// Returns the view’s nearest descendant (including itself) with a specific
// tag, or `nil` if no subview has that tag.
//
// tag: An integer identifier associated with a view object.
//
// See: https://developer.apple.com/documentation/AppKit/NSView/viewWithTag(_:)
func (v NSView) ViewWithTag(tag int) INSView {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("viewWithTag:"), tag)
	return NSViewFromID(rv)
}

// Informs the view that it’s being added to the view hierarchy of the
// specified window object (which may be `nil`).
//
// newWindow: The window object that will be at the root of the view’s new view
// hierarchy. If the view is being removed from a window and there is no new
// window, this parameter is `nil`.
//
// # Discussion
// 
// AppKit calls this method when the window of a view changes. It also calls
// it in cases where a view stays in the same window but its position in its
// view hierarchy changes. The view that moved also calls this method on all
// of its subviews, giving each of them a chance to respond to the change.
// 
// Subclasses can override this method to perform whatever actions are
// necessary. For example, when a window is deallocated, you can use this
// method to remove notification observers and bindings associated with the
// view.
// 
// When a window is deallocated, AppKit calls this method for each view in the
// window, passing `nil` for the `newWindow` parameter. AppKit does not
// necessarily call this method when closing a window, though. Closing a
// window usually just hides the window. Closed windows are deallocated only
// if their [ReleasedWhenClosed] method returns [true].
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSView/viewWillMove(toWindow:)
func (v NSView) ViewWillMoveToWindow(newWindow INSWindow) {
	objc.Send[objc.ID](v.ID, objc.Sel("viewWillMoveToWindow:"), newWindow)
}

// Informs the view of the start of a live resize—the user has started
// resizing the view.
//
// # Discussion
// 
// In the simple case, a view is sent [ViewWillStartLiveResize] before the
// first resize operation on the containing window and [ViewDidEndLiveResize]
// after the last resize operation. A view that is repeatedly added and
// removed from a window during live resize will receive only one
// [ViewWillStartLiveResize] (on the first time it is added to the window) and
// one [ViewDidEndLiveResize] (when the window has completed the live resize
// operation). This allows a superview such as [NSBrowser] object to add and
// remove its [NSMatrix] subviews during live resize without the [NSMatrix]
// object receiving multiple calls to these methods.
// 
// A view might allocate data structures to cache-drawing information in
// [ViewWillStartLiveResize] and should clean up these data structures in
// [ViewDidEndLiveResize]. In addition, a view that does optimized drawing
// during live resize might need to do full drawing after
// [ViewDidEndLiveResize]. However, you should not assume that a view has a
// drawing context in [ViewDidEndLiveResize] (since it may have been removed
// from the window during live resize). A view that needs to redraw itself
// after live resize should set its [needsDisplay] property to [true] in
// [ViewDidEndLiveResize].
// 
// A view subclass should call `super` from these methods.
//
// [needsDisplay]: https://developer.apple.com/documentation/AppKit/NSView/needsDisplay
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSView/viewWillStartLiveResize()
func (v NSView) ViewWillStartLiveResize() {
	objc.Send[objc.ID](v.ID, objc.Sel("viewWillStartLiveResize"))
}

// Called just before a contextual menu for a view is opened on screen.
//
// menu: The menu that will be opened.
//
// event: The event that caused the menu to open.
//
// # Discussion
// 
// This method is called just before a contextual menu for a view is opened on
// screen. It provides an opportunity to make any desired changes to the
// visual state of the view. For example, a table view might select a row in
// response to the contextual menu being displayed.
//
// See: https://developer.apple.com/documentation/AppKit/NSView/willOpenMenu(_:with:)
func (v NSView) WillOpenMenuWithEvent(menu INSMenu, event INSEvent) {
	objc.Send[objc.ID](v.ID, objc.Sel("willOpenMenu:withEvent:"), menu, event)
}

// Informs the view that it’s required to draw content.
//
// # Discussion
// 
// In response to receiving one of the `display` methods, the view recurses
// down the view hierarchy, sending this message to each of the views that may
// be involved in the display operation.
// 
// Subclasses can override this method to move or resize views, mark
// additional areas as requiring display, or take other actions that can best
// be deferred until they are required for drawing. During the recursion,
// setting the [needsDisplay] property or sending the [SetNeedsDisplayInRect]
// message to views in the hierarchy that are about to be drawn is valid and
// supported, and affects the assessment of the total area to be rendered in
// that drawing pass.
// 
// The following is an example of a generic subclass implementation:
//
// [needsDisplay]: https://developer.apple.com/documentation/AppKit/NSView/needsDisplay
//
// See: https://developer.apple.com/documentation/AppKit/NSView/viewWillDraw()
func (v NSView) ViewWillDraw() {
	objc.Send[objc.ID](v.ID, objc.Sel("viewWillDraw"))
}

// Writes EPS data that draws the region of the view within a specified
// rectangle onto a pasteboard.
//
// rect: A rectangle defining the region.
//
// pasteboard: An object representing a pasteboard.
//
// See: https://developer.apple.com/documentation/AppKit/NSView/writeEPS(inside:to:)
func (v NSView) WriteEPSInsideRectToPasteboard(rect corefoundation.CGRect, pasteboard INSPasteboard) {
	objc.Send[objc.ID](v.ID, objc.Sel("writeEPSInsideRect:toPasteboard:"), rect, pasteboard)
}

// Writes PDF data that draws the region of the view within a specified
// rectangle onto a pasteboard.
//
// rect: A rectangle defining the region.
//
// pasteboard: An object representing a pasteboard.
//
// See: https://developer.apple.com/documentation/AppKit/NSView/writePDF(inside:to:)
func (v NSView) WritePDFInsideRectToPasteboard(rect corefoundation.CGRect, pasteboard INSPasteboard) {
	objc.Send[objc.ID](v.ID, objc.Sel("writePDFInsideRect:toPasteboard:"), rect, pasteboard)
}

// Overridden by subclasses to perform additional actions before subviews are
// removed from the view.
//
// subview: The subview that will be removed.
//
// # Discussion
// 
// This method is invoked when `subview` receives a [RemoveFromSuperview]
// message or `subview` is removed from the view due to it being added to
// another view with [AddSubview].
//
// See: https://developer.apple.com/documentation/AppKit/NSView/willRemoveSubview(_:)
func (v NSView) WillRemoveSubview(subview INSView) {
	objc.Send[objc.ID](v.ID, objc.Sel("willRemoveSubview:"), subview)
}

// Sets the size of the view’s frame rectangle to the specified dimensions,
// resizing it within its superview without affecting its coordinate system.
//
// newSize: An [NSSize] structure specifying the new height and width of the frame
// rectangle.
//
// # Discussion
// 
// Changing the frame does not mark the view as needing to be displayed. Set
// the [needsDisplay] property to [true] when you want the view to be
// redisplayed.
// 
// Changing the frame size results in the posting of an
// [frameDidChangeNotification] to the default notification center if the view
// is configured to do so.
// 
// In macOS 10.4 and later, you can override this method to support content
// preservation during live resizing. In your overridden implementation,
// include some conditional code to be executed only during a live resize
// operation. Your code must invalidate any portions of your view that need to
// be redrawn.
//
// [frameDidChangeNotification]: https://developer.apple.com/documentation/AppKit/NSView/frameDidChangeNotification
// [needsDisplay]: https://developer.apple.com/documentation/AppKit/NSView/needsDisplay
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSView/setFrameSize(_:)
func (v NSView) SetFrameSize(newSize corefoundation.CGSize) {
	objc.Send[objc.ID](v.ID, objc.Sel("setFrameSize:"), newSize)
}

// Invoked by [Print] to determine the location of the region of the view
// being printed on the physical page.
//
// rect: A rectangle defining a region of the view; it is expressed in the default
// coordinate system of the page.
//
// # Return Value
// 
// A point to be used for setting the origin for `aRect`, whose size the view
// can examine in order to properly place it. It is expressed in the default
// coordinate system of the page.
//
// # Discussion
// 
// The default implementation places `aRect` according to the status of the
// [NSPrintInfo] object for the print job. By default it places the image in
// the upper-left corner of the page, but if the [NSPrintInfo] methods
// [HorizontallyCentered] or [VerticallyCentered] return [true], it centers a
// single-page image along the appropriate axis. A multiple-page document,
// however, is always placed so the divided pieces can be assembled at their
// edges.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSView/locationOfPrintRect(_:)
func (v NSView) LocationOfPrintRect(rect corefoundation.CGRect) corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](v.ID, objc.Sel("locationOfPrintRect:"), rect)
	return corefoundation.CGPoint(rv)
}

// Informs the view that its superview is about to change to the specified
// superview (which may be `nil`).
//
// newSuperview: A view object that will be the new superview of the view.
//
// # Discussion
// 
// Subclasses can override this method to perform whatever actions are
// necessary.
//
// See: https://developer.apple.com/documentation/AppKit/NSView/viewWillMove(toSuperview:)
func (v NSView) ViewWillMoveToSuperview(newSuperview INSView) {
	objc.Send[objc.ID](v.ID, objc.Sel("viewWillMoveToSuperview:"), newSuperview)
}

// Allows the user to drag objects from the view without activating the app or
// moving the window of the view forward, possibly obscuring the destination.
//
// event: An object representing an initial mouse-down event.
//
// # Return Value
// 
// If this method returns [true], the normal window-ordering and activation
// mechanism is delayed (not necessarily prevented) until the next mouse-up
// event. If it returns [false], then normal ordering and activation occur.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// Never invoke this method directly; it’s invoked automatically for each
// mouse-down event directed at the NSView.
// 
// An [NSView] subclass that allows dragging should implement this method to
// return [true] if `theEvent` is potentially the beginning of a dragging
// session or of some other context where window ordering isn’t appropriate.
// This method is invoked before a [MouseDown] message for `theEvent` is sent.
// The default implementation returns [false].
// 
// If, after delaying window ordering, the view actually initiates a dragging
// session or similar operation, it should also send a [PreventWindowOrdering]
// message to [NSApp], which completely prevents the window from ordering
// forward and the activation from becoming active. [PreventWindowOrdering] is
// sent automatically by the `` and `` methods of [NSView].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSView/shouldDelayWindowOrdering(for:)
func (v NSView) ShouldDelayWindowOrderingForEvent(event INSEvent) bool {
	rv := objc.Send[bool](v.ID, objc.Sel("shouldDelayWindowOrderingForEvent:"), event)
	return rv
}

// Returns a proxy object for the receiver that can be used to initiate
// implied animation for property changes.
//
// # Return Value
// 
// Returns a proxy object for the receiver that can initiate implied
// animations in response to property changes.
//
// # Discussion
// 
// The animator proxy object should be treated as if it was the receiver
// itself, and may be passed to any code that accepts the receiver as a
// parameter.
// 
// Sending key-value coding compliant “set” messages to the proxy will
// trigger animation for automatically animated properties of its target
// object, if the active [NSAnimationContext] in the current thread has a
// duration value greater than zero, and an animation for the property key is
// found by the [NSAnimatablePropertyContainer] search mechanism.
//
// See: https://developer.apple.com/documentation/AppKit/NSAnimatablePropertyContainer/animator()
func (v NSView) Animator() INSView {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("animator"))
	return NSViewFromID(rv)
}

// Returns the animation that should be performed for the specified key.
//
// key: The action name or property specified as a string.
//
// # Return Value
// 
// The animation to perform. A subclass of [CAAnimation].
//
// [CAAnimation]: https://developer.apple.com/documentation/QuartzCore/CAAnimation
//
// # Discussion
// 
// When the action specified by `key` is triggered for an object, this method
// is consulted to find the animation, if any, that should be performed in
// response.
// 
// Like its Core Animation [CALayer] counterpart, [action(forKey:)], this
// method is a funnel point that defines the order in which the search for an
// animation proceeds.It first checks the receiver’s Getting the Animator
// Proxy dictionary for a value matching `key`, then falls back to [Animator]
// for the receiver’s class.
// 
// Subclasses should not typically need to override this method.
//
// [CALayer]: https://developer.apple.com/documentation/QuartzCore/CALayer
// [action(forKey:)]: https://developer.apple.com/documentation/QuartzCore/CALayer/action(forKey:)
//
// See: https://developer.apple.com/documentation/AppKit/NSAnimatablePropertyContainer/animation(forKey:)
func (v NSView) AnimationForKey(key NSAnimatablePropertyKey) objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("animationForKey:"), objc.String(string(key)))
	return objectivec.Object{ID: rv}
}

// Invoked when the dragging operation is complete, signaling the receiver to
// perform any necessary clean-up.
//
// sender: The object sending the message; use it to get details about the dragging
// operation.
//
// # Discussion
// 
// For this method to be invoked, the previous [PerformDragOperation] must
// have returned [true].
// 
// The destination implements this method to perform any tidying up that it
// needs to do, such as updating its visual representation now that it has
// incorporated the dragged data. This message is the last message sent from
// `sender` to the destination during a dragging session.
// 
// If the `sender` object’s [AnimatesToDestination] property was set to
// [true] in [PrepareForDragOperation], then the drag image is still visible.
// At this point you should draw the final visual representation in the view.
// When this method returns, the drag image is removed form the screen. If
// your final visual representation matches the visual representation in the
// drag, this is a seamless transition.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSDraggingDestination/concludeDragOperation(_:)
func (v NSView) ConcludeDragOperation(sender NSDraggingInfo) {
	objc.Send[objc.ID](v.ID, objc.Sel("concludeDragOperation:"), sender)
}

// Called when a drag operation ends.
//
// sender: The object sending the message; use it to get details about the dragging
// operation.
//
// # Discussion
// 
// Implement this method if you want to be notified when a drag operation
// ends, which can be useful if the drag ends in some other destination. For
// example, this method might be used by a destination doing auto-expansion in
// order to collapse any auto-expands.
//
// See: https://developer.apple.com/documentation/AppKit/NSDraggingDestination/draggingEnded(_:)
func (v NSView) DraggingEnded(sender NSDraggingInfo) {
	objc.Send[objc.ID](v.ID, objc.Sel("draggingEnded:"), sender)
}

// Invoked when the dragged image exits the destination’s bounds rectangle
// (in the case of a view object) or its frame rectangle (in the case of a
// window object).
//
// sender: The object sending the message; use it to get details about the dragging
// operation.
//
// See: https://developer.apple.com/documentation/AppKit/NSDraggingDestination/draggingExited(_:)
func (v NSView) DraggingExited(sender NSDraggingInfo) {
	objc.Send[objc.ID](v.ID, objc.Sel("draggingExited:"), sender)
}

// Invoked after the released image has been removed from the screen,
// signaling the receiver to import the pasteboard data.
//
// sender: The object sending the message; use it to get details about the dragging
// operation.
//
// # Return Value
// 
// If the destination accepts the data, it returns [true]; otherwise it
// returns [false]. The default is to return [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// For this method to be invoked, the previous [PrepareForDragOperation]
// message must have returned [true]. The destination should implement this
// method to do the real work of importing the pasteboard data represented by
// the image.
// 
// If the sender object’s [AnimatesToDestination] was set to [true] in
// [PrepareForDragOperation], then setup any animation to arrange space for
// the drag items to animate to. Also at this time, enumerate through the
// dragging items to set their destination frames and destination images.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSDraggingDestination/performDragOperation(_:)
func (v NSView) PerformDragOperation(sender NSDraggingInfo) bool {
	rv := objc.Send[bool](v.ID, objc.Sel("performDragOperation:"), sender)
	return rv
}

// Invoked periodically as the image is held within the destination area,
// allowing modification of the dragging operation or mouse-pointer position.
//
// sender: The object sending the message; use it to get details about the dragging
// operation.
//
// # Return Value
// 
// One (and only one) of the dragging operation constants described in
// [NSDragOperation] in the [NSDraggingInfo] reference. The default return
// value (if this method is not implemented by the destination) is the value
// returned by the previous [DraggingEntered] message.
//
// [NSDragOperation]: https://developer.apple.com/documentation/AppKit/NSDragOperation
//
// # Discussion
// 
// For this to be invoked, the destination must have registered for the
// pasteboard data type involved in the drag operation. The messages continue
// until the image is either released or dragged out of the window or view.
// 
// This method provides the destination with an opportunity to modify the
// dragging operation depending on the position of the mouse pointer inside of
// the destination view or window object. For example, you may have several
// graphics or areas of text contained within the same view and wish to tailor
// the dragging operation, or to ignore the drag event completely, depending
// upon which object is underneath the mouse pointer at the time when the user
// releases the dragged image and the [PerformDragOperation] method is
// invoked.
// 
// You typically examine the contents of the pasteboard in the
// [DraggingEntered] method, where this examination is performed only once,
// rather than in the [DraggingUpdated] method, which is invoked multiple
// times.
// 
// Only one destination at a time receives a sequence of [DraggingUpdated]
// messages. If the mouse pointer is within the bounds of two overlapping
// views that are both valid destinations, the uppermost view receives these
// messages until the image is either released or dragged out.
//
// See: https://developer.apple.com/documentation/AppKit/NSDraggingDestination/draggingUpdated(_:)
func (v NSView) DraggingUpdated(sender NSDraggingInfo) NSDragOperation {
	rv := objc.Send[NSDragOperation](v.ID, objc.Sel("draggingUpdated:"), sender)
	return NSDragOperation(rv)
}

// Invoked when the dragged image enters destination bounds or frame; delegate
// returns dragging operation to perform.
//
// sender: The object sending the message; use it to get details about the dragging
// operation.
//
// # Return Value
// 
// One (and only one) of the dragging operation constants described in
// [NSDragOperation] in the [NSDraggingInfo] reference. The default return
// value (if this method is not implemented by the destination) is the value
// returned by the previous [DraggingEntered] message.
//
// [NSDragOperation]: https://developer.apple.com/documentation/AppKit/NSDragOperation
//
// # Discussion
// 
// Invoked when a dragged image enters the destination but only if the
// destination has registered for the pasteboard data type involved in the
// drag operation. Specifically, this method is invoked when the mouse pointer
// enters the destination’s bounds rectangle (if it is a view object) or its
// frame rectangle (if it is a window object).
// 
// This method must return a value that indicates which dragging operation the
// destination will perform when the image is released. In deciding which
// dragging operation to return, the method should evaluate the overlap
// between both the dragging operations allowed by the source (obtained from
// `sender` with the [DraggingSourceOperationMask] method) and the dragging
// operations and pasteboard data types the destination itself supports.
// 
// If none of the operations is appropriate, this method should return
// [NSDragOperationNone] (this is the default response if the method is not
// implemented by the destination). A destination will still receive
// [DraggingUpdated] and [DraggingExited] even if [NSDragOperationNone] is
// returned by this method.
//
// See: https://developer.apple.com/documentation/AppKit/NSDraggingDestination/draggingEntered(_:)
func (v NSView) DraggingEntered(sender NSDraggingInfo) NSDragOperation {
	rv := objc.Send[NSDragOperation](v.ID, objc.Sel("draggingEntered:"), sender)
	return NSDragOperation(rv)
}

// Invoked when the image is released, allowing the receiver to agree to or
// refuse drag operation.
//
// sender: The object sending the message; use it to get details about the dragging
// operation.
//
// # Return Value
// 
// [true] if the receiver agrees to perform the drag operation and [false] if
// not.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// This method is invoked only if the most recent [DraggingEntered] or
// [DraggingUpdated] message returned an acceptable drag-operation value.
// 
// If you want the drag items to animate from their current location on screen
// to their final location in your view, set the sender object’s
// [AnimatesToDestination] property to [true] in your implementation of this
// method.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSDraggingDestination/prepareForDragOperation(_:)
func (v NSView) PrepareForDragOperation(sender NSDraggingInfo) bool {
	rv := objc.Send[bool](v.ID, objc.Sel("prepareForDragOperation:"), sender)
	return rv
}

// Invoked when the dragging images should be changed.
//
// sender: The object sending the message; use this object to get details about the
// dragging operation.
//
// # Discussion
// 
// While a destination may change the dragging images at any time, it is
// recommended to wait until this method is called before updating the
// dragging images.
// 
// This allows the system to delay changing the dragging images until it is
// likely that the user will drop on this destination. Otherwise, the dragging
// images will change too often during the drag which would be distracting to
// the user.
// 
// During `` you may set non-acceptable drag items images to `nil` to hide
// them or use the enumeration option of
// [DraggingItemEnumerationClearNonenumeratedImages] If there are items that
// you hide, then after enumeration, you need to set the
// [NumberOfValidItemsForDrop] to the number of non-hidden drag items.
// However, if the valid item count is `0`, then it is better to return
// [NSDragOperationNone] from your implementation of [DraggingEntered] and, or
// [DraggingUpdated] instead of hiding all drag items during enumeration.
//
// [NSDragOperationNone]: https://developer.apple.com/documentation/AppKit/NSDragOperation/NSDragOperationNone
//
// See: https://developer.apple.com/documentation/AppKit/NSDraggingDestination/updateDraggingItemsForDrag(_:)
func (v NSView) UpdateDraggingItemsForDrag(sender NSDraggingInfo) {
	objc.Send[objc.ID](v.ID, objc.Sel("updateDraggingItemsForDrag:"), sender)
}

// Asks the destination object whether it wants to receive periodic
// [DraggingUpdated] messages.
//
// # Return Value
// 
// [true] if the destination wants to receive periodic [DraggingUpdated]
// messages, [false] otherwise.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// If the destination returns [false], these messages are sent only when the
// mouse moves or a modifier flag changes. Otherwise the destination gets the
// default behavior, where it receives periodic dragging-updated messages even
// if nothing changes.
//
// [false]: https://developer.apple.com/documentation/Swift/false
//
// See: https://developer.apple.com/documentation/AppKit/NSDraggingDestination/wantsPeriodicDraggingUpdates()
func (v NSView) WantsPeriodicDraggingUpdates() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("wantsPeriodicDraggingUpdates"))
	return rv
}
func (v NSView) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](v.ID, objc.Sel("encodeWithCoder:"), coder)
}





// Returns the default animation that should be performed for the specified
// key.
//
// key: The action name or property specified as a string.
//
// # Return Value
// 
// The animation to perform. A subclass of [CAAnimation].
//
// [CAAnimation]: https://developer.apple.com/documentation/QuartzCore/CAAnimation
//
// # Discussion
// 
// The [NSAnimatablePropertyContainer] method consults this class method when
// its search of the receivers Getting the Animator Proxy dictionary fails to
// return an animation for `key`.
// 
// An animatable property container should implement this method to return a
// default animation to be performed for each key that it wants to make
// auto-animatable, where `key` usually references a property of the receiver,
// but can also specify a special animation trigger
// ([NSAnimationTriggerOrderIn] or [NSAnimationTriggerOrderOut]).
// 
// A developer implementing a custom view subclass, can enable automatic
// animation for properties by overriding this method, and having it return
// the desired default [CAAnimation] subclass to use for each of the property
// keys of interest. The override should defer to super for any keys it
// doesn’t specifically handle, facilitating inheritance of default
// animation specifications. The following is an example of such an
// implementation.
//
// [CAAnimation]: https://developer.apple.com/documentation/QuartzCore/CAAnimation
// [NSAnimationTriggerOrderIn]: https://developer.apple.com/documentation/AppKit/NSAnimationTriggerOrderIn
// [NSAnimationTriggerOrderOut]: https://developer.apple.com/documentation/AppKit/NSAnimationTriggerOrderOut
//
// See: https://developer.apple.com/documentation/AppKit/NSAnimatablePropertyContainer/defaultAnimation(forKey:)
func (_NSViewClass NSViewClass) DefaultAnimationForKey(key NSAnimatablePropertyKey) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_NSViewClass.class), objc.Sel("defaultAnimationForKey:"), objc.String(string(key)))
	return objectivec.Object{ID: rv}
}








// When this property is true, any NSControls in the view or its descendants
// will be sized with compact metrics compatible with macOS 15 and earlier.
// Defaults to false
//
// See: https://developer.apple.com/documentation/AppKit/NSView/prefersCompactControlSizeMetrics
func (v NSView) PrefersCompactControlSizeMetrics() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("prefersCompactControlSizeMetrics"))
	return rv
}
func (v NSView) SetPrefersCompactControlSizeMetrics(value bool) {
	objc.Send[struct{}](v.ID, objc.Sel("setPrefersCompactControlSizeMetrics:"), value)
}



// See: https://developer.apple.com/documentation/AppKit/NSView/writingToolsCoordinator
func (v NSView) WritingToolsCoordinator() INSWritingToolsCoordinator {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("writingToolsCoordinator"))
	return NSWritingToolsCoordinatorFromID(objc.ID(rv))
}
func (v NSView) SetWritingToolsCoordinator(value INSWritingToolsCoordinator) {
	objc.Send[struct{}](v.ID, objc.Sel("setWritingToolsCoordinator:"), value)
}



// The view’s frame rectangle, which defines its position and size in its
// superview’s coordinate system.
//
// # Discussion
// 
// Changing the value of this property repositions and resizes the view within
// the coordinate system of its superview. Changing the frame does not mark
// the view as needing to be displayed. Set the [needsDisplay] property to
// [true] when you want the view to be redisplayed.
// 
// If your view does not use a custom bounds rectangle, this method also sets
// the view’s bounds to match the size of the new frame. You can specify a
// custom bounds rectangle by changing the [bounds] property or by calling the
// [SetBoundsOrigin] or [SetBoundsSize] method explicitly. Once set, the view
// creates an internal transform to convert from frame coordinates to bounds
// coordinates. As long as the width-to-height ratio of the two coordinate
// systems remains the same, your content appears normal. If the ratios
// differ, your content may appear skewed.
// 
// The frame rectangle may be rotated relative to its superview’s coordinate
// system. For more information, see the [frameRotation] property.
// 
// Changing the value of this property results in the posting of an
// [frameDidChangeNotification] to the default notification center if the view
// is configured to do so.
//
// [bounds]: https://developer.apple.com/documentation/AppKit/NSView/bounds
// [frameDidChangeNotification]: https://developer.apple.com/documentation/AppKit/NSView/frameDidChangeNotification
// [frameRotation]: https://developer.apple.com/documentation/AppKit/NSView/frameRotation
// [needsDisplay]: https://developer.apple.com/documentation/AppKit/NSView/needsDisplay
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSView/frame
func (v NSView) Frame() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](v.ID, objc.Sel("frame"))
	return corefoundation.CGRect(rv)
}
func (v NSView) SetFrame(value corefoundation.CGRect) {
	objc.Send[struct{}](v.ID, objc.Sel("setFrame:"), value)
}



// The view’s bounds rectangle, which expresses its location and size in its
// own coordinate system.
//
// See: https://developer.apple.com/documentation/appkit/nsview/bounds
func (v NSView) Bounds() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](v.ID, objc.Sel("bounds"))
	return corefoundation.CGRect(rv)
}
func (v NSView) SetBounds(value corefoundation.CGRect) {
	objc.Send[struct{}](v.ID, objc.Sel("setBounds:"), value)
}



// A Boolean value that determines whether the view needs to be redrawn before
// being displayed.
//
// See: https://developer.apple.com/documentation/appkit/nsview/needsdisplay
func (v NSView) NeedsDisplay() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("needsDisplay"))
	return rv
}
func (v NSView) SetNeedsDisplay(value bool) {
	objc.Send[struct{}](v.ID, objc.Sel("setNeedsDisplay:"), value)
}



// The view’s window object, if it is installed in a window.
//
// See: https://developer.apple.com/documentation/appkit/nsview/window
func (v NSView) Window() INSWindow {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("window"))
	return NSWindowFromID(objc.ID(rv))
}
func (v NSView) SetWindow(value INSWindow) {
	objc.Send[struct{}](v.ID, objc.Sel("setWindow:"), value)
}



// The menu item containing the view or any of its superviews in the view
// hierarchy.
//
// See: https://developer.apple.com/documentation/appkit/nsview/enclosingmenuitem
func (v NSView) EnclosingMenuItem() INSMenuItem {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("enclosingMenuItem"))
	return NSMenuItemFromID(objc.ID(rv))
}
func (v NSView) SetEnclosingMenuItem(value INSMenuItem) {
	objc.Send[struct{}](v.ID, objc.Sel("setEnclosingMenuItem:"), value)
}



// The nearest ancestor scroll view that contains the current view.
//
// See: https://developer.apple.com/documentation/appkit/nsview/enclosingscrollview
func (v NSView) EnclosingScrollView() INSScrollView {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("enclosingScrollView"))
	return NSScrollViewFromID(objc.ID(rv))
}
func (v NSView) SetEnclosingScrollView(value INSScrollView) {
	objc.Send[struct{}](v.ID, objc.Sel("setEnclosingScrollView:"), value)
}



// The view’s closest opaque ancestor, which might be the view itself.
//
// See: https://developer.apple.com/documentation/appkit/nsview/opaqueancestor
func (v NSView) OpaqueAncestor() INSView {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("opaqueAncestor"))
	return NSViewFromID(objc.ID(rv))
}
func (v NSView) SetOpaqueAncestor(value INSView) {
	objc.Send[struct{}](v.ID, objc.Sel("setOpaqueAncestor:"), value)
}



// The array of views embedded in the current view.
//
// See: https://developer.apple.com/documentation/appkit/nsview/subviews
func (v NSView) Subviews() INSView {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("subviews"))
	return NSViewFromID(objc.ID(rv))
}
func (v NSView) SetSubviews(value INSView) {
	objc.Send[struct{}](v.ID, objc.Sel("setSubviews:"), value)
}



// The view that is the parent of the current view.
//
// See: https://developer.apple.com/documentation/appkit/nsview/superview
func (v NSView) Superview() INSView {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("superview"))
	return NSViewFromID(objc.ID(rv))
}
func (v NSView) SetSuperview(value INSView) {
	objc.Send[struct{}](v.ID, objc.Sel("setSuperview:"), value)
}



// The view’s tag, which is an integer that you use to identify the view
// within your app.
//
// See: https://developer.apple.com/documentation/appkit/nsview/tag
func (v NSView) Tag() int {
	rv := objc.Send[int](v.ID, objc.Sel("tag"))
	return rv
}
func (v NSView) SetTag(value int) {
	objc.Send[struct{}](v.ID, objc.Sel("setTag:"), value)
}



// The angle of rotation, measured in degrees, applied to the view’s bounds
// rectangle relative to its frame rectangle.
//
// See: https://developer.apple.com/documentation/appkit/nsview/boundsrotation
func (v NSView) BoundsRotation() float64 {
	rv := objc.Send[float64](v.ID, objc.Sel("boundsRotation"))
	return rv
}
func (v NSView) SetBoundsRotation(value float64) {
	objc.Send[struct{}](v.ID, objc.Sel("setBoundsRotation:"), value)
}



// The angle of rotation, measured in degrees, applied to the view’s frame
// rectangle relative to its superview’s coordinate system.
//
// See: https://developer.apple.com/documentation/appkit/nsview/framerotation
func (v NSView) FrameRotation() float64 {
	rv := objc.Send[float64](v.ID, objc.Sel("frameRotation"))
	return rv
}
func (v NSView) SetFrameRotation(value float64) {
	objc.Send[struct{}](v.ID, objc.Sel("setFrameRotation:"), value)
}



// A Boolean value indicating whether the view uses a flipped coordinate
// system.
//
// See: https://developer.apple.com/documentation/appkit/nsview/isflipped
func (v NSView) IsFlipped() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("flipped"))
	return rv
}
func (v NSView) SetIsFlipped(value bool) {
	objc.Send[struct{}](v.ID, objc.Sel("setFlipped:"), value)
}



// A Boolean value indicating whether the view or any of its ancestors has
// ever had a rotation factor applied to its frame or bounds.
//
// See: https://developer.apple.com/documentation/appkit/nsview/isrotatedfrombase
func (v NSView) IsRotatedFromBase() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("rotatedFromBase"))
	return rv
}
func (v NSView) SetIsRotatedFromBase(value bool) {
	objc.Send[struct{}](v.ID, objc.Sel("setRotatedFromBase:"), value)
}



// A Boolean value indicating whether the view or any of its ancestors has
// ever had a rotation factor applied to its frame or bounds, or has been
// scaled from the window’s base coordinate system.
//
// See: https://developer.apple.com/documentation/appkit/nsview/isrotatedorscaledfrombase
func (v NSView) IsRotatedOrScaledFromBase() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("rotatedOrScaledFromBase"))
	return rv
}
func (v NSView) SetIsRotatedOrScaledFromBase(value bool) {
	objc.Send[struct{}](v.ID, objc.Sel("setRotatedOrScaledFromBase:"), value)
}



// A Boolean value indicating whether the view posts notifications when its
// bounds rectangle changes.
//
// See: https://developer.apple.com/documentation/appkit/nsview/postsboundschangednotifications
func (v NSView) PostsBoundsChangedNotifications() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("postsBoundsChangedNotifications"))
	return rv
}
func (v NSView) SetPostsBoundsChangedNotifications(value bool) {
	objc.Send[struct{}](v.ID, objc.Sel("setPostsBoundsChangedNotifications:"), value)
}



// A Boolean value indicating whether the view posts notifications when its
// frame rectangle changes.
//
// See: https://developer.apple.com/documentation/appkit/nsview/postsframechangednotifications
func (v NSView) PostsFrameChangedNotifications() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("postsFrameChangedNotifications"))
	return rv
}
func (v NSView) SetPostsFrameChangedNotifications(value bool) {
	objc.Send[struct{}](v.ID, objc.Sel("setPostsFrameChangedNotifications:"), value)
}



// A Boolean value indicating whether the view ensures it is vibrant on top of
// other content.
//
// See: https://developer.apple.com/documentation/appkit/nsview/allowsvibrancy
func (v NSView) AllowsVibrancy() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("allowsVibrancy"))
	return rv
}
func (v NSView) SetAllowsVibrancy(value bool) {
	objc.Send[struct{}](v.ID, objc.Sel("setAllowsVibrancy:"), value)
}



// The focus ring mask bounds, specified in the view’s coordinate space.
//
// See: https://developer.apple.com/documentation/appkit/nsview/focusringmaskbounds
func (v NSView) FocusRingMaskBounds() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](v.ID, objc.Sel("focusRingMaskBounds"))
	return corefoundation.CGRect(rv)
}
func (v NSView) SetFocusRingMaskBounds(value corefoundation.CGRect) {
	objc.Send[struct{}](v.ID, objc.Sel("setFocusRingMaskBounds:"), value)
}



// The type of focus ring drawn around the view.
//
// See: https://developer.apple.com/documentation/appkit/nsview/focusringtype
func (v NSView) FocusRingType() NSFocusRingType {
	rv := objc.Send[NSFocusRingType](v.ID, objc.Sel("focusRingType"))
	return NSFocusRingType(rv)
}
func (v NSView) SetFocusRingType(value NSFocusRingType) {
	objc.Send[struct{}](v.ID, objc.Sel("setFocusRingType:"), value)
}



// A Boolean value indicating whether the view or one of its ancestors is
// being drawn for a find indicator.
//
// See: https://developer.apple.com/documentation/appkit/nsview/isdrawingfindindicator
func (v NSView) IsDrawingFindIndicator() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("drawingFindIndicator"))
	return rv
}
func (v NSView) SetIsDrawingFindIndicator(value bool) {
	objc.Send[struct{}](v.ID, objc.Sel("setDrawingFindIndicator:"), value)
}



// A Boolean value indicating whether the view is hidden.
//
// See: https://developer.apple.com/documentation/appkit/nsview/ishidden
func (v NSView) IsHidden() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("hidden"))
	return rv
}
func (v NSView) SetIsHidden(value bool) {
	objc.Send[struct{}](v.ID, objc.Sel("setHidden:"), value)
}



// A Boolean value indicating whether the view is hidden from sight because
// it, or one of its ancestors, is marked as hidden.
//
// See: https://developer.apple.com/documentation/appkit/nsview/ishiddenorhashiddenancestor
func (v NSView) IsHiddenOrHasHiddenAncestor() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("hiddenOrHasHiddenAncestor"))
	return rv
}
func (v NSView) SetIsHiddenOrHasHiddenAncestor(value bool) {
	objc.Send[struct{}](v.ID, objc.Sel("setHiddenOrHasHiddenAncestor:"), value)
}



// The opacity of the view.
//
// See: https://developer.apple.com/documentation/appkit/nsview/alphavalue
func (v NSView) AlphaValue() float64 {
	rv := objc.Send[float64](v.ID, objc.Sel("alphaValue"))
	return rv
}
func (v NSView) SetAlphaValue(value float64) {
	objc.Send[struct{}](v.ID, objc.Sel("setAlphaValue:"), value)
}



// An array of Core Image filters to apply to the view’s background.
//
// See: https://developer.apple.com/documentation/appkit/nsview/backgroundfilters
func (v NSView) BackgroundFilters() objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("backgroundFilters"))
	return objectivec.Object{ID: rv}
}
func (v NSView) SetBackgroundFilters(value objectivec.IObject) {
	objc.Send[struct{}](v.ID, objc.Sel("setBackgroundFilters:"), value)
}



// A Boolean value indicating whether the view incorporates content from its
// subviews into its own layer.
//
// See: https://developer.apple.com/documentation/appkit/nsview/candrawsubviewsintolayer
func (v NSView) CanDrawSubviewsIntoLayer() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("canDrawSubviewsIntoLayer"))
	return rv
}
func (v NSView) SetCanDrawSubviewsIntoLayer(value bool) {
	objc.Send[struct{}](v.ID, objc.Sel("setCanDrawSubviewsIntoLayer:"), value)
}



// The Core Image filter used to composite the view’s contents with its
// background.
//
// See: https://developer.apple.com/documentation/appkit/nsview/compositingfilter
func (v NSView) CompositingFilter() objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("compositingFilter"))
	return objectivec.Object{ID: rv}
}
func (v NSView) SetCompositingFilter(value objectivec.IObject) {
	objc.Send[struct{}](v.ID, objc.Sel("setCompositingFilter:"), value)
}



// An array of Core Image filters to apply to the contents of the view and its
// sublayers.
//
// See: https://developer.apple.com/documentation/appkit/nsview/contentfilters
func (v NSView) ContentFilters() objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("contentFilters"))
	return objectivec.Object{ID: rv}
}
func (v NSView) SetContentFilters(value objectivec.IObject) {
	objc.Send[struct{}](v.ID, objc.Sel("setContentFilters:"), value)
}



// The rotation angle of the view around the center of its layer.
//
// See: https://developer.apple.com/documentation/appkit/nsview/framecenterrotation
func (v NSView) FrameCenterRotation() float64 {
	rv := objc.Send[float64](v.ID, objc.Sel("frameCenterRotation"))
	return rv
}
func (v NSView) SetFrameCenterRotation(value float64) {
	objc.Send[struct{}](v.ID, objc.Sel("setFrameCenterRotation:"), value)
}



// The Core Animation layer that the view uses as its backing store.
//
// See: https://developer.apple.com/documentation/appkit/nsview/layer
func (v NSView) Layer() objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("layer"))
	return objectivec.Object{ID: rv}
}
func (v NSView) SetLayer(value objectivec.IObject) {
	objc.Send[struct{}](v.ID, objc.Sel("setLayer:"), value)
}



// The current layer contents placement policy.
//
// See: https://developer.apple.com/documentation/appkit/nsview/layercontentsplacement-swift.property
func (v NSView) LayerContentsPlacement() objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("layerContentsPlacement"))
	return objectivec.Object{ID: rv}
}
func (v NSView) SetLayerContentsPlacement(value objectivec.IObject) {
	objc.Send[struct{}](v.ID, objc.Sel("setLayerContentsPlacement:"), value)
}



// The contents redraw policy for the view’s layer.
//
// See: https://developer.apple.com/documentation/appkit/nsview/layercontentsredrawpolicy-swift.property
func (v NSView) LayerContentsRedrawPolicy() objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("layerContentsRedrawPolicy"))
	return objectivec.Object{ID: rv}
}
func (v NSView) SetLayerContentsRedrawPolicy(value objectivec.IObject) {
	objc.Send[struct{}](v.ID, objc.Sel("setLayerContentsRedrawPolicy:"), value)
}



// A Boolean value indicating whether the view’s layer uses Core Image
// filters and needs in-process rendering.
//
// See: https://developer.apple.com/documentation/appkit/nsview/layerusescoreimagefilters
func (v NSView) LayerUsesCoreImageFilters() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("layerUsesCoreImageFilters"))
	return rv
}
func (v NSView) SetLayerUsesCoreImageFilters(value bool) {
	objc.Send[struct{}](v.ID, objc.Sel("setLayerUsesCoreImageFilters:"), value)
}



// The shadow displayed underneath the view.
//
// See: https://developer.apple.com/documentation/appkit/nsview/shadow
func (v NSView) Shadow() INSShadow {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("shadow"))
	return NSShadowFromID(objc.ID(rv))
}
func (v NSView) SetShadow(value INSShadow) {
	objc.Send[struct{}](v.ID, objc.Sel("setShadow:"), value)
}



// A Boolean value indicating whether the view uses a layer as its backing
// store.
//
// See: https://developer.apple.com/documentation/appkit/nsview/wantslayer
func (v NSView) WantsLayer() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("wantsLayer"))
	return rv
}
func (v NSView) SetWantsLayer(value bool) {
	objc.Send[struct{}](v.ID, objc.Sel("setWantsLayer:"), value)
}



// A Boolean value indicating which drawing path the view takes when updating
// its contents.
//
// See: https://developer.apple.com/documentation/appkit/nsview/wantsupdatelayer
func (v NSView) WantsUpdateLayer() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("wantsUpdateLayer"))
	return rv
}
func (v NSView) SetWantsUpdateLayer(value bool) {
	objc.Send[struct{}](v.ID, objc.Sel("setWantsUpdateLayer:"), value)
}



// The text for the view’s tooltip.
//
// See: https://developer.apple.com/documentation/appkit/nsview/tooltip
func (v NSView) ToolTip() string {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("toolTip"))
	return foundation.NSStringFromID(rv).String()
}
func (v NSView) SetToolTip(value string) {
	objc.Send[struct{}](v.ID, objc.Sel("setToolTip:"), objc.String(value))
}



// Custom insets that you specify to modify your view’s safe area
//
// See: https://developer.apple.com/documentation/appkit/nsview/additionalsafeareainsets
func (v NSView) AdditionalSafeAreaInsets() foundation.NSEdgeInsets {
	rv := objc.Send[foundation.NSEdgeInsets](v.ID, objc.Sel("additionalSafeAreaInsets"))
	return foundation.NSEdgeInsets(rv)
}
func (v NSView) SetAdditionalSafeAreaInsets(value foundation.NSEdgeInsets) {
	objc.Send[struct{}](v.ID, objc.Sel("setAdditionalSafeAreaInsets:"), value)
}



// The insets (in points) from the view’s frame that define its content
// rectangle.
//
// See: https://developer.apple.com/documentation/appkit/nsview/alignmentrectinsets
func (v NSView) AlignmentRectInsets() foundation.NSEdgeInsets {
	rv := objc.Send[foundation.NSEdgeInsets](v.ID, objc.Sel("alignmentRectInsets"))
	return foundation.NSEdgeInsets(rv)
}
func (v NSView) SetAlignmentRectInsets(value foundation.NSEdgeInsets) {
	objc.Send[struct{}](v.ID, objc.Sel("setAlignmentRectInsets:"), value)
}



// A Boolean value indicating whether the view applies the autoresizing
// behavior to its subviews when its frame size changes.
//
// See: https://developer.apple.com/documentation/appkit/nsview/autoresizessubviews
func (v NSView) AutoresizesSubviews() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("autoresizesSubviews"))
	return rv
}
func (v NSView) SetAutoresizesSubviews(value bool) {
	objc.Send[struct{}](v.ID, objc.Sel("setAutoresizesSubviews:"), value)
}



// The options that determine how the view is resized relative to its
// superview.
//
// See: https://developer.apple.com/documentation/appkit/nsview/autoresizingmask-swift.property
func (v NSView) AutoresizingMask() uint {
	rv := objc.Send[uint](v.ID, objc.Sel("autoresizingMask"))
	return rv
}
func (v NSView) SetAutoresizingMask(value uint) {
	objc.Send[struct{}](v.ID, objc.Sel("setAutoresizingMask:"), value)
}



// The distance (in points) between the bottom of the view’s alignment
// rectangle and its baseline.
//
// See: https://developer.apple.com/documentation/appkit/nsview/baselineoffsetfrombottom
func (v NSView) BaselineOffsetFromBottom() float64 {
	rv := objc.Send[float64](v.ID, objc.Sel("baselineOffsetFromBottom"))
	return rv
}
func (v NSView) SetBaselineOffsetFromBottom(value float64) {
	objc.Send[struct{}](v.ID, objc.Sel("setBaselineOffsetFromBottom:"), value)
}



// A layout anchor representing the bottom edge of the view’s frame.
//
// See: https://developer.apple.com/documentation/appkit/nsview/bottomanchor
func (v NSView) BottomAnchor() INSLayoutYAxisAnchor {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("bottomAnchor"))
	return NSLayoutYAxisAnchorFromID(objc.ID(rv))
}
func (v NSView) SetBottomAnchor(value INSLayoutYAxisAnchor) {
	objc.Send[struct{}](v.ID, objc.Sel("setBottomAnchor:"), value)
}



// A layout anchor representing the horizontal center of the view’s frame.
//
// See: https://developer.apple.com/documentation/appkit/nsview/centerxanchor
func (v NSView) CenterXAnchor() INSLayoutXAxisAnchor {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("centerXAnchor"))
	return NSLayoutXAxisAnchorFromID(objc.ID(rv))
}
func (v NSView) SetCenterXAnchor(value INSLayoutXAxisAnchor) {
	objc.Send[struct{}](v.ID, objc.Sel("setCenterXAnchor:"), value)
}



// A layout anchor representing the vertical center of the view’s frame.
//
// See: https://developer.apple.com/documentation/appkit/nsview/centeryanchor
func (v NSView) CenterYAnchor() INSLayoutYAxisAnchor {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("centerYAnchor"))
	return NSLayoutYAxisAnchorFromID(objc.ID(rv))
}
func (v NSView) SetCenterYAnchor(value INSLayoutYAxisAnchor) {
	objc.Send[struct{}](v.ID, objc.Sel("setCenterYAnchor:"), value)
}



// Returns the constraints held by the view.
//
// See: https://developer.apple.com/documentation/appkit/nsview/constraints
func (v NSView) Constraints() INSLayoutConstraint {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("constraints"))
	return NSLayoutConstraintFromID(objc.ID(rv))
}
func (v NSView) SetConstraints(value INSLayoutConstraint) {
	objc.Send[struct{}](v.ID, objc.Sel("setConstraints:"), value)
}



// A layout anchor representing the baseline for the topmost line of text in
// the view.
//
// See: https://developer.apple.com/documentation/appkit/nsview/firstbaselineanchor
func (v NSView) FirstBaselineAnchor() INSLayoutYAxisAnchor {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("firstBaselineAnchor"))
	return NSLayoutYAxisAnchorFromID(objc.ID(rv))
}
func (v NSView) SetFirstBaselineAnchor(value INSLayoutYAxisAnchor) {
	objc.Send[struct{}](v.ID, objc.Sel("setFirstBaselineAnchor:"), value)
}



// The distance (in points) between the top of the view’s alignment
// rectangle and its topmost baseline.
//
// See: https://developer.apple.com/documentation/appkit/nsview/firstbaselineoffsetfromtop
func (v NSView) FirstBaselineOffsetFromTop() float64 {
	rv := objc.Send[float64](v.ID, objc.Sel("firstBaselineOffsetFromTop"))
	return rv
}
func (v NSView) SetFirstBaselineOffsetFromTop(value float64) {
	objc.Send[struct{}](v.ID, objc.Sel("setFirstBaselineOffsetFromTop:"), value)
}



// The minimum size of the view that satisfies the constraints it holds.
//
// See: https://developer.apple.com/documentation/appkit/nsview/fittingsize
func (v NSView) FittingSize() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](v.ID, objc.Sel("fittingSize"))
	return corefoundation.CGSize(rv)
}
func (v NSView) SetFittingSize(value corefoundation.CGSize) {
	objc.Send[struct{}](v.ID, objc.Sel("setFittingSize:"), value)
}



// A Boolean value indicating whether the constraints impacting the layout of
// the view incompletely specify the location of the view.
//
// See: https://developer.apple.com/documentation/appkit/nsview/hasambiguouslayout
func (v NSView) HasAmbiguousLayout() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("hasAmbiguousLayout"))
	return rv
}
func (v NSView) SetHasAmbiguousLayout(value bool) {
	objc.Send[struct{}](v.ID, objc.Sel("setHasAmbiguousLayout:"), value)
}



// A layout anchor representing the height of the view’s frame.
//
// See: https://developer.apple.com/documentation/appkit/nsview/heightanchor
func (v NSView) HeightAnchor() INSLayoutDimension {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("heightAnchor"))
	return NSLayoutDimensionFromID(objc.ID(rv))
}
func (v NSView) SetHeightAnchor(value INSLayoutDimension) {
	objc.Send[struct{}](v.ID, objc.Sel("setHeightAnchor:"), value)
}



// The natural size for the receiving view, considering only properties of the
// view itself.
//
// See: https://developer.apple.com/documentation/appkit/nsview/intrinsiccontentsize
func (v NSView) IntrinsicContentSize() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](v.ID, objc.Sel("intrinsicContentSize"))
	return corefoundation.CGSize(rv)
}
func (v NSView) SetIntrinsicContentSize(value corefoundation.CGSize) {
	objc.Send[struct{}](v.ID, objc.Sel("setIntrinsicContentSize:"), value)
}



// A Boolean value that indicates whether the view’s horizontal size
// constraints are active.
//
// See: https://developer.apple.com/documentation/appkit/nsview/ishorizontalcontentsizeconstraintactive
func (v NSView) IsHorizontalContentSizeConstraintActive() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("horizontalContentSizeConstraintActive"))
	return rv
}
func (v NSView) SetIsHorizontalContentSizeConstraintActive(value bool) {
	objc.Send[struct{}](v.ID, objc.Sel("setHorizontalContentSizeConstraintActive:"), value)
}



// A Boolean value that indicates whether the view’s vertical size
// constraints are active.
//
// See: https://developer.apple.com/documentation/appkit/nsview/isverticalcontentsizeconstraintactive
func (v NSView) IsVerticalContentSizeConstraintActive() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("verticalContentSizeConstraintActive"))
	return rv
}
func (v NSView) SetIsVerticalContentSizeConstraintActive(value bool) {
	objc.Send[struct{}](v.ID, objc.Sel("setVerticalContentSizeConstraintActive:"), value)
}



// A layout anchor representing the baseline for the bottommost line of text
// in the view.
//
// See: https://developer.apple.com/documentation/appkit/nsview/lastbaselineanchor
func (v NSView) LastBaselineAnchor() INSLayoutYAxisAnchor {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("lastBaselineAnchor"))
	return NSLayoutYAxisAnchorFromID(objc.ID(rv))
}
func (v NSView) SetLastBaselineAnchor(value INSLayoutYAxisAnchor) {
	objc.Send[struct{}](v.ID, objc.Sel("setLastBaselineAnchor:"), value)
}



// The distance (in points) between the bottom of the view’s alignment
// rectangle and its bottommost baseline.
//
// See: https://developer.apple.com/documentation/appkit/nsview/lastbaselineoffsetfrombottom
func (v NSView) LastBaselineOffsetFromBottom() float64 {
	rv := objc.Send[float64](v.ID, objc.Sel("lastBaselineOffsetFromBottom"))
	return rv
}
func (v NSView) SetLastBaselineOffsetFromBottom(value float64) {
	objc.Send[struct{}](v.ID, objc.Sel("setLastBaselineOffsetFromBottom:"), value)
}



// The array of layout guide objects owned by this view.
//
// See: https://developer.apple.com/documentation/appkit/nsview/layoutguides
func (v NSView) LayoutGuides() INSLayoutGuide {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("layoutGuides"))
	return NSLayoutGuideFromID(objc.ID(rv))
}
func (v NSView) SetLayoutGuides(value INSLayoutGuide) {
	objc.Send[struct{}](v.ID, objc.Sel("setLayoutGuides:"), value)
}



// A layout guide that provides the recommended amount of padding for content
// inside of a view.
//
// See: https://developer.apple.com/documentation/appkit/nsview/layoutmarginsguide
func (v NSView) LayoutMarginsGuide() INSLayoutGuide {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("layoutMarginsGuide"))
	return NSLayoutGuideFromID(objc.ID(rv))
}
func (v NSView) SetLayoutMarginsGuide(value INSLayoutGuide) {
	objc.Send[struct{}](v.ID, objc.Sel("setLayoutMarginsGuide:"), value)
}



// A layout anchor representing the leading edge of the view’s frame.
//
// See: https://developer.apple.com/documentation/appkit/nsview/leadinganchor
func (v NSView) LeadingAnchor() INSLayoutXAxisAnchor {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("leadingAnchor"))
	return NSLayoutXAxisAnchorFromID(objc.ID(rv))
}
func (v NSView) SetLeadingAnchor(value INSLayoutXAxisAnchor) {
	objc.Send[struct{}](v.ID, objc.Sel("setLeadingAnchor:"), value)
}



// A layout anchor representing the left edge of the view’s frame.
//
// See: https://developer.apple.com/documentation/appkit/nsview/leftanchor
func (v NSView) LeftAnchor() INSLayoutXAxisAnchor {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("leftAnchor"))
	return NSLayoutXAxisAnchorFromID(objc.ID(rv))
}
func (v NSView) SetLeftAnchor(value INSLayoutXAxisAnchor) {
	objc.Send[struct{}](v.ID, objc.Sel("setLeftAnchor:"), value)
}



// A Boolean value indicating whether the view needs a layout pass before it
// can be drawn.
//
// See: https://developer.apple.com/documentation/appkit/nsview/needslayout
func (v NSView) NeedsLayout() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("needsLayout"))
	return rv
}
func (v NSView) SetNeedsLayout(value bool) {
	objc.Send[struct{}](v.ID, objc.Sel("setNeedsLayout:"), value)
}



// A Boolean value indicating whether the view’s constraints need to be
// updated.
//
// See: https://developer.apple.com/documentation/appkit/nsview/needsupdateconstraints
func (v NSView) NeedsUpdateConstraints() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("needsUpdateConstraints"))
	return rv
}
func (v NSView) SetNeedsUpdateConstraints(value bool) {
	objc.Send[struct{}](v.ID, objc.Sel("setNeedsUpdateConstraints:"), value)
}



// A layout anchor representing the right edge of the view’s frame.
//
// See: https://developer.apple.com/documentation/appkit/nsview/rightanchor
func (v NSView) RightAnchor() INSLayoutXAxisAnchor {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("rightAnchor"))
	return NSLayoutXAxisAnchorFromID(objc.ID(rv))
}
func (v NSView) SetRightAnchor(value INSLayoutXAxisAnchor) {
	objc.Send[struct{}](v.ID, objc.Sel("setRightAnchor:"), value)
}



// The distances from the edges of your view that define the safe area.
//
// See: https://developer.apple.com/documentation/appkit/nsview/safeareainsets
func (v NSView) SafeAreaInsets() foundation.NSEdgeInsets {
	rv := objc.Send[foundation.NSEdgeInsets](v.ID, objc.Sel("safeAreaInsets"))
	return foundation.NSEdgeInsets(rv)
}
func (v NSView) SetSafeAreaInsets(value foundation.NSEdgeInsets) {
	objc.Send[struct{}](v.ID, objc.Sel("setSafeAreaInsets:"), value)
}



// The layout guide you use to position content inside your view’s safe
// area.
//
// See: https://developer.apple.com/documentation/appkit/nsview/safearealayoutguide
func (v NSView) SafeAreaLayoutGuide() INSLayoutGuide {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("safeAreaLayoutGuide"))
	return NSLayoutGuideFromID(objc.ID(rv))
}
func (v NSView) SetSafeAreaLayoutGuide(value INSLayoutGuide) {
	objc.Send[struct{}](v.ID, objc.Sel("setSafeAreaLayoutGuide:"), value)
}



// A rectangle in the view’s coordinate system that contains the unobscured
// portion of the view.
//
// See: https://developer.apple.com/documentation/appkit/nsview/safearearect
func (v NSView) SafeAreaRect() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](v.ID, objc.Sel("safeAreaRect"))
	return corefoundation.CGRect(rv)
}
func (v NSView) SetSafeAreaRect(value corefoundation.CGRect) {
	objc.Send[struct{}](v.ID, objc.Sel("setSafeAreaRect:"), value)
}



// A layout anchor representing the top edge of the view’s frame.
//
// See: https://developer.apple.com/documentation/appkit/nsview/topanchor
func (v NSView) TopAnchor() INSLayoutYAxisAnchor {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("topAnchor"))
	return NSLayoutYAxisAnchorFromID(objc.ID(rv))
}
func (v NSView) SetTopAnchor(value INSLayoutYAxisAnchor) {
	objc.Send[struct{}](v.ID, objc.Sel("setTopAnchor:"), value)
}



// A layout anchor representing the trailing edge of the view’s frame.
//
// See: https://developer.apple.com/documentation/appkit/nsview/trailinganchor
func (v NSView) TrailingAnchor() INSLayoutXAxisAnchor {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("trailingAnchor"))
	return NSLayoutXAxisAnchorFromID(objc.ID(rv))
}
func (v NSView) SetTrailingAnchor(value INSLayoutXAxisAnchor) {
	objc.Send[struct{}](v.ID, objc.Sel("setTrailingAnchor:"), value)
}



// A Boolean value indicating whether the view’s autoresizing mask is
// translated into constraints for the constraint-based layout system.
//
// See: https://developer.apple.com/documentation/appkit/nsview/translatesautoresizingmaskintoconstraints
func (v NSView) TranslatesAutoresizingMaskIntoConstraints() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("translatesAutoresizingMaskIntoConstraints"))
	return rv
}
func (v NSView) SetTranslatesAutoresizingMaskIntoConstraints(value bool) {
	objc.Send[struct{}](v.ID, objc.Sel("setTranslatesAutoresizingMaskIntoConstraints:"), value)
}



// The layout direction for content in the view.
//
// See: https://developer.apple.com/documentation/appkit/nsview/userinterfacelayoutdirection
func (v NSView) UserInterfaceLayoutDirection() NSUserInterfaceLayoutDirection {
	rv := objc.Send[NSUserInterfaceLayoutDirection](v.ID, objc.Sel("userInterfaceLayoutDirection"))
	return NSUserInterfaceLayoutDirection(rv)
}
func (v NSView) SetUserInterfaceLayoutDirection(value NSUserInterfaceLayoutDirection) {
	objc.Send[struct{}](v.ID, objc.Sel("setUserInterfaceLayoutDirection:"), value)
}



// A layout anchor representing the width of the view’s frame.
//
// See: https://developer.apple.com/documentation/appkit/nsview/widthanchor
func (v NSView) WidthAnchor() INSLayoutDimension {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("widthAnchor"))
	return NSLayoutDimensionFromID(objc.ID(rv))
}
func (v NSView) SetWidthAnchor(value INSLayoutDimension) {
	objc.Send[struct{}](v.ID, objc.Sel("setWidthAnchor:"), value)
}



// A Boolean value indicating whether the view can draw its contents on a
// background thread.
//
// See: https://developer.apple.com/documentation/appkit/nsview/candrawconcurrently
func (v NSView) CanDrawConcurrently() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("canDrawConcurrently"))
	return rv
}
func (v NSView) SetCanDrawConcurrently(value bool) {
	objc.Send[struct{}](v.ID, objc.Sel("setCanDrawConcurrently:"), value)
}



// A Boolean value that indicates whether the view, and its subviews, confine
// their drawing areas to the bounds of the view.
//
// See: https://developer.apple.com/documentation/appkit/nsview/clipstobounds
func (v NSView) ClipsToBounds() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("clipsToBounds"))
	return rv
}
func (v NSView) SetClipsToBounds(value bool) {
	objc.Send[struct{}](v.ID, objc.Sel("setClipsToBounds:"), value)
}



// A Boolean value indicating whether the view is being rendered as part of a
// live resizing operation.
//
// See: https://developer.apple.com/documentation/appkit/nsview/inliveresize
func (v NSView) InLiveResize() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("inLiveResize"))
	return rv
}
func (v NSView) SetInLiveResize(value bool) {
	objc.Send[struct{}](v.ID, objc.Sel("setInLiveResize:"), value)
}



// A Boolean value indicating whether the view is in full screen mode.
//
// See: https://developer.apple.com/documentation/appkit/nsview/isinfullscreenmode
func (v NSView) IsInFullScreenMode() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("inFullScreenMode"))
	return rv
}
func (v NSView) SetIsInFullScreenMode(value bool) {
	objc.Send[struct{}](v.ID, objc.Sel("setInFullScreenMode:"), value)
}



// A Boolean value indicating whether the view fills its frame rectangle with
// opaque content.
//
// See: https://developer.apple.com/documentation/appkit/nsview/isopaque
func (v NSView) IsOpaque() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("opaque"))
	return rv
}
func (v NSView) SetIsOpaque(value bool) {
	objc.Send[struct{}](v.ID, objc.Sel("setOpaque:"), value)
}



// A Boolean value indicating whether the view optimizes live-resize
// operations by preserving content that has not moved.
//
// See: https://developer.apple.com/documentation/appkit/nsview/preservescontentduringliveresize
func (v NSView) PreservesContentDuringLiveResize() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("preservesContentDuringLiveResize"))
	return rv
}
func (v NSView) SetPreservesContentDuringLiveResize(value bool) {
	objc.Send[struct{}](v.ID, objc.Sel("setPreservesContentDuringLiveResize:"), value)
}



// The rectangle identifying the portion of your view that did not change
// during a live resize operation.
//
// See: https://developer.apple.com/documentation/appkit/nsview/rectpreservedduringliveresize
func (v NSView) RectPreservedDuringLiveResize() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](v.ID, objc.Sel("rectPreservedDuringLiveResize"))
	return corefoundation.CGRect(rv)
}
func (v NSView) SetRectPreservedDuringLiveResize(value corefoundation.CGRect) {
	objc.Send[struct{}](v.ID, objc.Sel("setRectPreservedDuringLiveResize:"), value)
}



// The portion of the view that isn’t clipped by its superviews.
//
// See: https://developer.apple.com/documentation/appkit/nsview/visiblerect
func (v NSView) VisibleRect() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](v.ID, objc.Sel("visibleRect"))
	return corefoundation.CGRect(rv)
}
func (v NSView) SetVisibleRect(value corefoundation.CGRect) {
	objc.Send[struct{}](v.ID, objc.Sel("setVisibleRect:"), value)
}



// A Boolean value indicating whether AppKit’s default clipping behavior is
// in effect.
//
// See: https://developer.apple.com/documentation/appkit/nsview/wantsdefaultclipping
func (v NSView) WantsDefaultClipping() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("wantsDefaultClipping"))
	return rv
}
func (v NSView) SetWantsDefaultClipping(value bool) {
	objc.Send[struct{}](v.ID, objc.Sel("setWantsDefaultClipping:"), value)
}



// The fraction of the page that can be pushed onto the next page during
// automatic pagination to prevent items such as lines of text from being
// divided across pages.
//
// See: https://developer.apple.com/documentation/appkit/nsview/heightadjustlimit
func (v NSView) HeightAdjustLimit() float64 {
	rv := objc.Send[float64](v.ID, objc.Sel("heightAdjustLimit"))
	return rv
}
func (v NSView) SetHeightAdjustLimit(value float64) {
	objc.Send[struct{}](v.ID, objc.Sel("setHeightAdjustLimit:"), value)
}



// A default footer string that includes the current page number and page
// count.
//
// See: https://developer.apple.com/documentation/appkit/nsview/pagefooter
func (v NSView) PageFooter() foundation.NSAttributedString {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("pageFooter"))
	return foundation.NSAttributedStringFromID(objc.ID(rv))
}
func (v NSView) SetPageFooter(value foundation.NSAttributedString) {
	objc.Send[struct{}](v.ID, objc.Sel("setPageFooter:"), value)
}



// A default header string that includes the print job title and date.
//
// See: https://developer.apple.com/documentation/appkit/nsview/pageheader
func (v NSView) PageHeader() foundation.NSAttributedString {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("pageHeader"))
	return foundation.NSAttributedStringFromID(objc.ID(rv))
}
func (v NSView) SetPageHeader(value foundation.NSAttributedString) {
	objc.Send[struct{}](v.ID, objc.Sel("setPageHeader:"), value)
}



// The view’s print job title.
//
// See: https://developer.apple.com/documentation/appkit/nsview/printjobtitle
func (v NSView) PrintJobTitle() string {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("printJobTitle"))
	return foundation.NSStringFromID(rv).String()
}
func (v NSView) SetPrintJobTitle(value string) {
	objc.Send[struct{}](v.ID, objc.Sel("setPrintJobTitle:"), objc.String(value))
}



// The fraction of the page that can be pushed onto the next page during
// automatic pagination to prevent items such as small images or text columns
// from being divided across pages.
//
// See: https://developer.apple.com/documentation/appkit/nsview/widthadjustlimit
func (v NSView) WidthAdjustLimit() float64 {
	rv := objc.Send[float64](v.ID, objc.Sel("widthAdjustLimit"))
	return rv
}
func (v NSView) SetWidthAdjustLimit(value float64) {
	objc.Send[struct{}](v.ID, objc.Sel("setWidthAdjustLimit:"), value)
}



// The types of touch interactions the view allows.
//
// See: https://developer.apple.com/documentation/appkit/nsview/allowedtouchtypes
func (v NSView) AllowedTouchTypes() NSTouchTypeMask {
	rv := objc.Send[NSTouchTypeMask](v.ID, objc.Sel("allowedTouchTypes"))
	return NSTouchTypeMask(rv)
}
func (v NSView) SetAllowedTouchTypes(value NSTouchTypeMask) {
	objc.Send[struct{}](v.ID, objc.Sel("setAllowedTouchTypes:"), value)
}



// A Boolean value indicating whether the view can become key view.
//
// See: https://developer.apple.com/documentation/appkit/nsview/canbecomekeyview
func (v NSView) CanBecomeKeyView() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("canBecomeKeyView"))
	return rv
}
func (v NSView) SetCanBecomeKeyView(value bool) {
	objc.Send[struct{}](v.ID, objc.Sel("setCanBecomeKeyView:"), value)
}



// See: https://developer.apple.com/documentation/appkit/nsview/candidatelisttouchbaritem
func (v NSView) CandidateListTouchBarItem() INSCandidateListTouchBarItem {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("candidateListTouchBarItem"))
	return NSCandidateListTouchBarItemFromID(objc.ID(rv))
}
func (v NSView) SetCandidateListTouchBarItem(value INSCandidateListTouchBarItem) {
	objc.Send[struct{}](v.ID, objc.Sel("setCandidateListTouchBarItem:"), value)
}



// The gesture recognize objects currently attached to the view.
//
// See: https://developer.apple.com/documentation/appkit/nsview/gesturerecognizers
func (v NSView) GestureRecognizers() INSGestureRecognizer {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("gestureRecognizers"))
	return NSGestureRecognizerFromID(objc.ID(rv))
}
func (v NSView) SetGestureRecognizers(value INSGestureRecognizer) {
	objc.Send[struct{}](v.ID, objc.Sel("setGestureRecognizers:"), value)
}



// The text input context object for the view.
//
// See: https://developer.apple.com/documentation/appkit/nsview/inputcontext
func (v NSView) InputContext() INSTextInputContext {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("inputContext"))
	return NSTextInputContextFromID(objc.ID(rv))
}
func (v NSView) SetInputContext(value INSTextInputContext) {
	objc.Send[struct{}](v.ID, objc.Sel("setInputContext:"), value)
}



// A Boolean value indicating whether the view can pass mouse down events
// through to its superviews.
//
// See: https://developer.apple.com/documentation/appkit/nsview/mousedowncanmovewindow
func (v NSView) MouseDownCanMoveWindow() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("mouseDownCanMoveWindow"))
	return rv
}
func (v NSView) SetMouseDownCanMoveWindow(value bool) {
	objc.Send[struct{}](v.ID, objc.Sel("setMouseDownCanMoveWindow:"), value)
}



// A Boolean value indicating whether the view needs its panel to become the
// key window before it can handle keyboard input and navigation.
//
// See: https://developer.apple.com/documentation/appkit/nsview/needspaneltobecomekey
func (v NSView) NeedsPanelToBecomeKey() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("needsPanelToBecomeKey"))
	return rv
}
func (v NSView) SetNeedsPanelToBecomeKey(value bool) {
	objc.Send[struct{}](v.ID, objc.Sel("setNeedsPanelToBecomeKey:"), value)
}



// The view object that follows the current view in the key view loop.
//
// See: https://developer.apple.com/documentation/appkit/nsview/nextkeyview
func (v NSView) NextKeyView() INSView {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("nextKeyView"))
	return NSViewFromID(objc.ID(rv))
}
func (v NSView) SetNextKeyView(value INSView) {
	objc.Send[struct{}](v.ID, objc.Sel("setNextKeyView:"), value)
}



// The closest view object in the key view loop that follows the current view
// in the key view loop and accepts first responder status.
//
// See: https://developer.apple.com/documentation/appkit/nsview/nextvalidkeyview
func (v NSView) NextValidKeyView() INSView {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("nextValidKeyView"))
	return NSViewFromID(objc.ID(rv))
}
func (v NSView) SetNextValidKeyView(value INSView) {
	objc.Send[struct{}](v.ID, objc.Sel("setNextValidKeyView:"), value)
}



// The portion of the view that has been rendered and is available for
// responsive scrolling.
//
// See: https://developer.apple.com/documentation/appkit/nsview/preparedcontentrect
func (v NSView) PreparedContentRect() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](v.ID, objc.Sel("preparedContentRect"))
	return corefoundation.CGRect(rv)
}
func (v NSView) SetPreparedContentRect(value corefoundation.CGRect) {
	objc.Send[struct{}](v.ID, objc.Sel("setPreparedContentRect:"), value)
}



// Configures the behavior and progression of the Force Touch trackpad when
// responding to touch input produced by the user when the cursor is over the
// view.
//
// See: https://developer.apple.com/documentation/appkit/nsview/pressureconfiguration
func (v NSView) PressureConfiguration() INSPressureConfiguration {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("pressureConfiguration"))
	return NSPressureConfigurationFromID(objc.ID(rv))
}
func (v NSView) SetPressureConfiguration(value INSPressureConfiguration) {
	objc.Send[struct{}](v.ID, objc.Sel("setPressureConfiguration:"), value)
}



// The view object preceding the current view in the key view loop.
//
// See: https://developer.apple.com/documentation/appkit/nsview/previouskeyview
func (v NSView) PreviousKeyView() INSView {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("previousKeyView"))
	return NSViewFromID(objc.ID(rv))
}
func (v NSView) SetPreviousKeyView(value INSView) {
	objc.Send[struct{}](v.ID, objc.Sel("setPreviousKeyView:"), value)
}



// The closest view object in the key view loop that precedes the current view
// and accepts first responder status.
//
// See: https://developer.apple.com/documentation/appkit/nsview/previousvalidkeyview
func (v NSView) PreviousValidKeyView() INSView {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("previousValidKeyView"))
	return NSViewFromID(objc.ID(rv))
}
func (v NSView) SetPreviousValidKeyView(value INSView) {
	objc.Send[struct{}](v.ID, objc.Sel("setPreviousValidKeyView:"), value)
}



// The array of pasteboard drag types that the view can accept.
//
// See: https://developer.apple.com/documentation/appkit/nsview/registereddraggedtypes
func (v NSView) RegisteredDraggedTypes() NSPasteboardType {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("registeredDraggedTypes"))
	return NSPasteboardType(foundation.NSStringFromID(rv).String())
}
func (v NSView) SetRegisteredDraggedTypes(value NSPasteboardType) {
	objc.Send[struct{}](v.ID, objc.Sel("setRegisteredDraggedTypes:"), objc.String(string(value)))
}



// An array of the view’s tracking areas.
//
// See: https://developer.apple.com/documentation/appkit/nsview/trackingareas
func (v NSView) TrackingAreas() INSTrackingArea {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("trackingAreas"))
	return NSTrackingAreaFromID(objc.ID(rv))
}
func (v NSView) SetTrackingAreas(value INSTrackingArea) {
	objc.Send[struct{}](v.ID, objc.Sel("setTrackingAreas:"), value)
}



// A Boolean value indicating whether the view wants resting touches.
//
// See: https://developer.apple.com/documentation/appkit/nsview/wantsrestingtouches
func (v NSView) WantsRestingTouches() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("wantsRestingTouches"))
	return rv
}
func (v NSView) SetWantsRestingTouches(value bool) {
	objc.Send[struct{}](v.ID, objc.Sel("setWantsRestingTouches:"), value)
}



// A Boolean value indicating whether the view accepts touch events.
//
// See: https://developer.apple.com/documentation/appkit/nsview/acceptstouchevents
func (v NSView) AcceptsTouchEvents() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("acceptsTouchEvents"))
	return rv
}
func (v NSView) SetAcceptsTouchEvents(value bool) {
	objc.Send[struct{}](v.ID, objc.Sel("setAcceptsTouchEvents:"), value)
}



// A Boolean value indicating whether drawing commands will produce any
// results.
//
// See: https://developer.apple.com/documentation/appkit/nsview/candraw
func (v NSView) CanDraw() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("canDraw"))
	return rv
}
func (v NSView) SetCanDraw(value bool) {
	objc.Send[struct{}](v.ID, objc.Sel("setCanDraw:"), value)
}



// A Boolean value indicating whether the view wants an OpenGL backing surface
// with a resolution greater than 1 pixel per point.
//
// See: https://developer.apple.com/documentation/appkit/nsview/wantsbestresolutionopenglsurface
func (v NSView) WantsBestResolutionOpenGLSurface() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("wantsBestResolutionOpenGLSurface"))
	return rv
}
func (v NSView) SetWantsBestResolutionOpenGLSurface(value bool) {
	objc.Send[struct{}](v.ID, objc.Sel("setWantsBestResolutionOpenGLSurface:"), value)
}



// See: https://developer.apple.com/documentation/appkit/nsview/wantsextendeddynamicrangeopenglsurface
func (v NSView) WantsExtendedDynamicRangeOpenGLSurface() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("wantsExtendedDynamicRangeOpenGLSurface"))
	return rv
}
func (v NSView) SetWantsExtendedDynamicRangeOpenGLSurface(value bool) {
	objc.Send[struct{}](v.ID, objc.Sel("setWantsExtendedDynamicRangeOpenGLSurface:"), value)
}



// Sets the option dictionary that maps event trigger keys to animation
// objects.
//
// See: https://developer.apple.com/documentation/AppKit/NSAnimatablePropertyContainer/animations
func (v NSView) Animations() foundation.INSDictionary {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("animations"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (v NSView) SetAnimations(value foundation.INSDictionary) {
	objc.Send[struct{}](v.ID, objc.Sel("setAnimations:"), value)
}



// The appearance of the receiver, in an [NSAppearance] object.
//
// # Discussion
// 
// The default value for this property is `nil`, which means that the receiver
// uses the appearance it inherits from the nearest ancestor that has set an
// appearance. When you set `appearance` to a non-`nil` value, the receiver
// and the views it contains use the specified appearance.
//
// See: https://developer.apple.com/documentation/AppKit/NSAppearanceCustomization/appearance
func (v NSView) Appearance() INSAppearance {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("appearance"))
	return NSAppearanceFromID(objc.ID(rv))
}
func (v NSView) SetAppearance(value INSAppearance) {
	objc.Send[struct{}](v.ID, objc.Sel("setAppearance:"), value)
}



// The appearance that will be used when the receiver is drawn onscreen, in an
// [NSAppearance] object. (read-only)
//
// # Discussion
// 
// The default value for this property is provided by the nearest ancestor of
// the receiver that has set an appearance.
// 
// You can use this property to ensure that an offscreen view sets the
// appropriate current appearance when it draws onscreen.
//
// See: https://developer.apple.com/documentation/AppKit/NSAppearanceCustomization/effectiveAppearance
func (v NSView) EffectiveAppearance() INSAppearance {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("effectiveAppearance"))
	return NSAppearanceFromID(objc.ID(rv))
}



// A string that identifies the user interface item.
//
// # Discussion
// 
// Identifiers are used during window restoration operations to uniquely
// identify the windows of the application. You can set the value of this
// string programmatically or in Interface Builder. If you create an item in
// Interface Builder and do not set a value for this string, a unique value is
// created for the item when the nib file is loaded. For programmatically
// created views, you typically set this value after creating the item but
// before adding it to a window.
// 
// You should not change the value of a window’s identifier after adding any
// views to the window. For views and controls in a window, the value you
// specify for this string must be unique on a per-window basis.
// 
// The slash (`/`), backslash (`\`), or colon (`:`) characters are reserved
// and must not be used in your custom identifiers. Similarly, Apple reserves
// all identifiers beginning with an underscore (`_`) character. Applications
// and frameworks should use a consistent prefix for their identifiers to
// avoid collisions with other frameworks. For a list of prefixes used by the
// system frameworks, see [OS X Frameworks] in [Mac Technology Overview].
// 
// If you are subclassing a class from one of the system frameworks, do not
// override the accessor methods of this protocol.
//
// [Mac Technology Overview]: https://developer.apple.com/library/archive/documentation/MacOSX/Conceptual/OSX_Technology_Overview/About/About.html#//apple_ref/doc/uid/TP40001067
// [OS X Frameworks]: https://developer.apple.com/library/archive/documentation/MacOSX/Conceptual/OSX_Technology_Overview/SystemFrameworks/SystemFrameworks.html#//apple_ref/doc/uid/TP40001067-CH210
//
// See: https://developer.apple.com/documentation/AppKit/NSUserInterfaceItemIdentification/identifier
func (v NSView) Identifier() NSUserInterfaceItemIdentifier {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("identifier"))
	return NSUserInterfaceItemIdentifier(foundation.NSStringFromID(rv).String())
}
func (v NSView) SetIdentifier(value NSUserInterfaceItemIdentifier) {
	objc.Send[struct{}](v.ID, objc.Sel("setIdentifier:"), objc.String(string(value)))
}







// Returns the default focus ring type.
//
// See: https://developer.apple.com/documentation/appkit/nsview/defaultfocusringtype
func (_NSViewClass NSViewClass) DefaultFocusRingType() NSFocusRingType {
	rv := objc.Send[NSFocusRingType](objc.ID(_NSViewClass.class), objc.Sel("defaultFocusRingType"))
	return NSFocusRingType(rv)
}
func (_NSViewClass NSViewClass) SetDefaultFocusRingType(value NSFocusRingType) {
	objc.Send[struct{}](objc.ID(_NSViewClass.class), objc.Sel("setDefaultFocusRingType:"), value)
}



// Overridden by subclasses to return the default pop-up menu for instances of
// the receiving class.
//
// See: https://developer.apple.com/documentation/appkit/nsview/defaultmenu
func (_NSViewClass NSViewClass) DefaultMenu() NSMenu {
	rv := objc.Send[objc.ID](objc.ID(_NSViewClass.class), objc.Sel("defaultMenu"))
	return NSMenuFromID(objc.ID(rv))
}
func (_NSViewClass NSViewClass) SetDefaultMenu(value NSMenu) {
	objc.Send[struct{}](objc.ID(_NSViewClass.class), objc.Sel("setDefaultMenu:"), value)
}



// The currently focused view object.
//
// See: https://developer.apple.com/documentation/appkit/nsview/focusview
func (_NSViewClass NSViewClass) FocusView() NSView {
	rv := objc.Send[objc.ID](objc.ID(_NSViewClass.class), objc.Sel("focusView"))
	return NSViewFromID(objc.ID(rv))
}
func (_NSViewClass NSViewClass) SetFocusView(value NSView) {
	objc.Send[struct{}](objc.ID(_NSViewClass.class), objc.Sel("setFocusView:"), value)
}



// Returns a Boolean value indicating whether the view depends on the
// constraint-based layout system.
//
// See: https://developer.apple.com/documentation/appkit/nsview/requiresconstraintbasedlayout
func (_NSViewClass NSViewClass) RequiresConstraintBasedLayout() bool {
	rv := objc.Send[bool](objc.ID(_NSViewClass.class), objc.Sel("requiresConstraintBasedLayout"))
	return rv
}
func (_NSViewClass NSViewClass) SetRequiresConstraintBasedLayout(value bool) {
	objc.Send[struct{}](objc.ID(_NSViewClass.class), objc.Sel("setRequiresConstraintBasedLayout:"), value)
}



// A Boolean value that indicates whether views support responsive scrolling.
//
// See: https://developer.apple.com/documentation/appkit/nsview/iscompatiblewithresponsivescrolling
func (_NSViewClass NSViewClass) IsCompatibleWithResponsiveScrolling() bool {
	rv := objc.Send[bool](objc.ID(_NSViewClass.class), objc.Sel("compatibleWithResponsiveScrolling"))
	return rv
}
func (_NSViewClass NSViewClass) SetIsCompatibleWithResponsiveScrolling(value bool) {
	objc.Send[struct{}](objc.ID(_NSViewClass.class), objc.Sel("setCompatibleWithResponsiveScrolling:"), value)
}










			// Protocol methods for NSAccessibilityElementProtocol
			

// Returns the accessibility element’s frame in screen coordinates.
//
// # Return Value
// 
// The element’s frame in screen coordinates.
//
// # Discussion
// 
// This method is the getter for the [NSAccessibilityProtocol] protocol’s
// [accessibilityFrame] property. This method is called whenever accessibility
// clients request the [size] or [position] attributes.
//
// [accessibilityFrame]: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityFrame
// [position]: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Attribute/position
// [size]: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Attribute/size
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityElementProtocol/accessibilityFrame()

func (o NSView) AccessibilityFrame() corefoundation.CGRect {
	
	rv := objc.Send[corefoundation.CGRect](o.ID, objc.Sel("accessibilityFrame"))
	return rv
	}

// Returns the accessibility element’s parent in the accessibility
// hierarchy.
//
// # Return Value
// 
// The element’s parent in the accessibility hierarchy.
//
// # Discussion
// 
// This method is the getter for the [NSAccessibilityProtocol] protocol’s
// [accessibilityParent] property.
//
// [accessibilityParent]: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityParent
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityElementProtocol/accessibilityParent()

func (o NSView) AccessibilityParent() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityParent"))
	return objectivec.Object{ID: rv}
	}

// Returns the accessibility element’s identity.
//
// # Return Value
// 
// Returns the unique ID for the accessibility element. It is often used in
// automated testing.
//
// # Discussion
// 
// This method is the getter for the [NSAccessibilityProtocol] protocol’s
// [accessibilityIdentifier] property.
//
// [accessibilityIdentifier]: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityIdentifier
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityElementProtocol/accessibilityIdentifier()

func (o NSView) AccessibilityIdentifier() string {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityIdentifier"))
	return foundation.NSStringFromID(rv).String()
	}

// Returns a Boolean value that indicates whether the accessibility element
// has the keyboard focus.
//
// # Return Value
// 
// [true] if this element has the keyboard focus; otherwise, [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// This method is the getter for the [NSAccessibilityProtocol] protocol’s
// [accessibilityFocused] property.
//
// [accessibilityFocused]: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityFocused
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityElementProtocol/isAccessibilityFocused()

func (o NSView) IsAccessibilityFocused() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityFocused"))
	return rv
	}




			// Protocol methods for NSAccessibilityProtocol
			

// Returns a Boolean value that determines whether the accessibility element
// participates in the accessibility hierarchy.
//
// See: /documentation/appkit/nsaccessibilityprotocol/isaccessibilityelement()

func (o NSView) IsAccessibilityElement() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityElement"))
	return rv
	}

// Sets a Boolean value that determines whether the accessibility element
// participates in the accessibility hierarchy.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityelement(_:)

func (o NSView) SetAccessibilityElement(accessibilityElement bool) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityElement:"), accessibilityElement)
	}

// Returns a Boolean value that determines whether the accessibility element
// responds to user events.
//
// See: /documentation/appkit/nsaccessibilityprotocol/isaccessibilityenabled()

func (o NSView) IsAccessibilityEnabled() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityEnabled"))
	return rv
	}

// Sets a Boolean value that determines whether the accessibility element
// responds to user events.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityenabled(_:)

func (o NSView) SetAccessibilityEnabled(accessibilityEnabled bool) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityEnabled:"), accessibilityEnabled)
	}

// Sets the accessibility element’s frame in screen coordinates.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityframe(_:)

func (o NSView) SetAccessibilityFrame(accessibilityFrame corefoundation.CGRect) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityFrame:"), accessibilityFrame)
	}

// Returns the help text for the accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityhelp()

func (o NSView) AccessibilityHelp() string {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityHelp"))
	return foundation.NSStringFromID(rv).String()
	}

// Sets the help text for the accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityhelp(_:)

func (o NSView) SetAccessibilityHelp(accessibilityHelp string) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityHelp:"), objc.String(accessibilityHelp))
	}

// Returns a short description of the accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitylabel()

func (o NSView) AccessibilityLabel() string {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityLabel"))
	return foundation.NSStringFromID(rv).String()
	}

// Sets a short description of the accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitylabel(_:)

func (o NSView) SetAccessibilityLabel(accessibilityLabel string) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityLabel:"), objc.String(accessibilityLabel))
	}

// Returns the title of the accessibility element—for example, a button’s
// visible text.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitytitle()

func (o NSView) AccessibilityTitle() string {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityTitle"))
	return foundation.NSStringFromID(rv).String()
	}

// Sets the title of the accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitytitle(_:)

func (o NSView) SetAccessibilityTitle(accessibilityTitle string) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityTitle:"), objc.String(accessibilityTitle))
	}

// Returns the accessibility element’s value.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityvalue()

func (o NSView) AccessibilityValue() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityValue"))
	return objectivec.Object{ID: rv}
	}

// Sets the accessibility element’s value.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityvalue(_:)

func (o NSView) SetAccessibilityValue(accessibilityValue objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityValue:"), accessibilityValue)
	}

// Returns a Boolean value that indicates whether assistive apps can invoke
// the specified selector on the accessibility element.
//
// selector: The selector to check.
//
// # Return Value
// 
// [true], if accessibility clients can call the selector; otherwise, [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilitySelectorAllowed(_:)

func (o NSView) IsAccessibilitySelectorAllowed(selector objc.SEL) bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilitySelectorAllowed:"), selector)
	return rv
	}

// Returns the contents of the current accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitycontents()

func (o NSView) AccessibilityContents() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityContents"))
	return objectivec.Object{ID: rv}
	}

// Sets the contents of the current accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitycontents(_:)

func (o NSView) SetAccessibilityContents(accessibilityContents foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityContents:"), accessibilityContents)
	}

// Returns the critical value for the level indicator.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitycriticalvalue()

func (o NSView) AccessibilityCriticalValue() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityCriticalValue"))
	return objectivec.Object{ID: rv}
	}

// Sets the critical value for the level indicator.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitycriticalvalue(_:)

func (o NSView) SetAccessibilityCriticalValue(accessibilityCriticalValue objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityCriticalValue:"), accessibilityCriticalValue)
	}

// Sets the accessibility element’s identity.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityidentifier(_:)

func (o NSView) SetAccessibilityIdentifier(accessibilityIdentifier string) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityIdentifier:"), objc.String(accessibilityIdentifier))
	}

// Returns the maximum value for the accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitymaxvalue()

func (o NSView) AccessibilityMaxValue() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityMaxValue"))
	return objectivec.Object{ID: rv}
	}

// Sets the maximum value for the accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitymaxvalue(_:)

func (o NSView) SetAccessibilityMaxValue(accessibilityMaxValue objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityMaxValue:"), accessibilityMaxValue)
	}

// Returns the minimum value for the accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityminvalue()

func (o NSView) AccessibilityMinValue() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityMinValue"))
	return objectivec.Object{ID: rv}
	}

// Sets the minimum value for the accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityminvalue(_:)

func (o NSView) SetAccessibilityMinValue(accessibilityMinValue objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityMinValue:"), accessibilityMinValue)
	}

// Returns the orientation of the accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityorientation()

func (o NSView) AccessibilityOrientation() NSAccessibilityOrientation {
	
	rv := objc.Send[NSAccessibilityOrientation](o.ID, objc.Sel("accessibilityOrientation"))
	return rv
	}

// Sets the orientation of the accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityorientation(_:)

func (o NSView) SetAccessibilityOrientation(accessibilityOrientation NSAccessibilityOrientation) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityOrientation:"), accessibilityOrientation)
	}

// Returns a Boolean value that determines whether the accessibility element
// contains protected content.
//
// See: /documentation/appkit/nsaccessibilityprotocol/isaccessibilityprotectedcontent()

func (o NSView) IsAccessibilityProtectedContent() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityProtectedContent"))
	return rv
	}

// Sets a Boolean value that determines whether the accessibility element
// contains protected content.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityprotectedcontent(_:)

func (o NSView) SetAccessibilityProtectedContent(accessibilityProtectedContent bool) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityProtectedContent:"), accessibilityProtectedContent)
	}

// Returns a Boolean value that determines whether the accessibility element
// is currently in a selected state.
//
// See: /documentation/appkit/nsaccessibilityprotocol/isaccessibilityselected()

func (o NSView) IsAccessibilitySelected() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilitySelected"))
	return rv
	}

// Sets a Boolean value that determines whether the accessibility element is
// currently in a selected state.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityselected(_:)

func (o NSView) SetAccessibilitySelected(accessibilitySelected bool) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySelected:"), accessibilitySelected)
	}

// Returns the URL for the accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityurl()

func (o NSView) AccessibilityURL() foundation.INSURL {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityURL"))
	return foundation.NSURLFromID(rv)
	}

// Sets the URL for the accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityurl(_:)

func (o NSView) SetAccessibilityURL(accessibilityURL foundation.INSURL) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityURL:"), accessibilityURL)
	}

// Returns the human-readable description of the accessibility element’s
// value.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityvaluedescription()

func (o NSView) AccessibilityValueDescription() string {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityValueDescription"))
	return foundation.NSStringFromID(rv).String()
	}

// Sets the human-readable description of the accessibility element’s value.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityvaluedescription(_:)

func (o NSView) SetAccessibilityValueDescription(accessibilityValueDescription string) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityValueDescription:"), objc.String(accessibilityValueDescription))
	}

// Returns the warning value for the level indicator.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitywarningvalue()

func (o NSView) AccessibilityWarningValue() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityWarningValue"))
	return objectivec.Object{ID: rv}
	}

// Sets the warning value for the level indicator.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitywarningvalue(_:)

func (o NSView) SetAccessibilityWarningValue(accessibilityWarningValue objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityWarningValue:"), accessibilityWarningValue)
	}

// Returns the child accessibility elements in the accessibility hierarchy.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitychildren()

func (o NSView) AccessibilityChildren() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityChildren"))
	return objectivec.Object{ID: rv}
	}

// Sets the child accessibility elements in the accessibility hierarchy.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitychildren(_:)

func (o NSView) SetAccessibilityChildren(accessibilityChildren foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityChildren:"), accessibilityChildren)
	}

// Returns the array of child accessibility elements in order for linear
// navigation.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitychildreninnavigationorder()

func (o NSView) AccessibilityChildrenInNavigationOrder() unsafe.Pointer {
	
	rv := objc.Send[unsafe.Pointer](o.ID, objc.Sel("accessibilityChildrenInNavigationOrder"))
	return rv
	}

// Sets the array of child accessibility elements in order for linear
// navigation.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitychildreninnavigationorder(_:)

func (o NSView) SetAccessibilityChildrenInNavigationOrder(accessibilityChildrenInNavigationOrder foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityChildrenInNavigationOrder:"), accessibilityChildrenInNavigationOrder)
	}

// Sets the accessibility element’s parent in the accessibility hierarchy.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityparent(_:)

func (o NSView) SetAccessibilityParent(accessibilityParent objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityParent:"), accessibilityParent)
	}

// Returns the accessibility element’s currently selected children.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityselectedchildren()

func (o NSView) AccessibilitySelectedChildren() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilitySelectedChildren"))
	return objectivec.Object{ID: rv}
	}

// Sets the accessibility element’s currently selected children.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityselectedchildren(_:)

func (o NSView) SetAccessibilitySelectedChildren(accessibilitySelectedChildren foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySelectedChildren:"), accessibilitySelectedChildren)
	}

// Returns the top-level element that contains the accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitytopleveluielement()

func (o NSView) AccessibilityTopLevelUIElement() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityTopLevelUIElement"))
	return objectivec.Object{ID: rv}
	}

// Sets the top-level element that contains the accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitytopleveluielement(_:)

func (o NSView) SetAccessibilityTopLevelUIElement(accessibilityTopLevelUIElement objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityTopLevelUIElement:"), accessibilityTopLevelUIElement)
	}

// Returns the accessibility element’s visible child accessibility elements.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityvisiblechildren()

func (o NSView) AccessibilityVisibleChildren() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityVisibleChildren"))
	return objectivec.Object{ID: rv}
	}

// Sets the accessibility element’s visible child accessibility elements.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityvisiblechildren(_:)

func (o NSView) SetAccessibilityVisibleChildren(accessibilityVisibleChildren foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityVisibleChildren:"), accessibilityVisibleChildren)
	}

// Returns the child accessibility element with the current focus.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityapplicationfocuseduielement()

func (o NSView) AccessibilityApplicationFocusedUIElement() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityApplicationFocusedUIElement"))
	return objectivec.Object{ID: rv}
	}

// Sets the child accessibility element with the current focus.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityapplicationfocuseduielement(_:)

func (o NSView) SetAccessibilityApplicationFocusedUIElement(accessibilityApplicationFocusedUIElement objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityApplicationFocusedUIElement:"), accessibilityApplicationFocusedUIElement)
	}

// Sets a Boolean value that determines whether the accessibility element has
// the keyboard focus.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityfocused(_:)

func (o NSView) SetAccessibilityFocused(accessibilityFocused bool) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityFocused:"), accessibilityFocused)
	}

// Returns the child window with the current focus.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityfocusedwindow()

func (o NSView) AccessibilityFocusedWindow() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityFocusedWindow"))
	return objectivec.Object{ID: rv}
	}

// Sets the child window with the current focus.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityfocusedwindow(_:)

func (o NSView) SetAccessibilityFocusedWindow(accessibilityFocusedWindow objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityFocusedWindow:"), accessibilityFocusedWindow)
	}

// Returns the array of elements that shares the keyboard focus with the
// accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitysharedfocuselements()

func (o NSView) AccessibilitySharedFocusElements() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilitySharedFocusElements"))
	return objectivec.Object{ID: rv}
	}

// Sets the array of elements that shares the keyboard focus with the
// accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitysharedfocuselements(_:)

func (o NSView) SetAccessibilitySharedFocusElements(accessibilitySharedFocusElements foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySharedFocusElements:"), accessibilitySharedFocusElements)
	}

// Returns a Boolean value that determines whether the accessibility element
// must have content for successful submission of a form.
//
// See: /documentation/appkit/nsaccessibilityprotocol/isaccessibilityrequired()

func (o NSView) IsAccessibilityRequired() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityRequired"))
	return rv
	}

// Sets a Boolean value that determines whether the accessibility element must
// have content for successful submission of a form.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityrequired(_:)

func (o NSView) SetAccessibilityRequired(accessibilityRequired bool) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityRequired:"), accessibilityRequired)
	}

// Returns the type of interface element that the accessibility element
// represents.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityrole()

func (o NSView) AccessibilityRole() NSAccessibilityRole {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityRole"))
	return NSAccessibilityRole(foundation.NSStringFromID(rv).String())
	}

// Sets the type of interface element that the accessibility element
// represents.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityrole(_:)

func (o NSView) SetAccessibilityRole(accessibilityRole NSAccessibilityRole) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityRole:"), objc.String(string(accessibilityRole)))
	}

// Returns a localized, human-intelligible description of the accessibility
// element’s role, such as
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityroledescription()

func (o NSView) AccessibilityRoleDescription() string {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityRoleDescription"))
	return foundation.NSStringFromID(rv).String()
	}

// Sets the localized, human-intelligible description of the accessibility
// element’s role, such as
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityroledescription(_:)

func (o NSView) SetAccessibilityRoleDescription(accessibilityRoleDescription string) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityRoleDescription:"), objc.String(accessibilityRoleDescription))
	}

// Returns the specialized interface element type that the accessibility
// element represents.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitysubrole()

func (o NSView) AccessibilitySubrole() NSAccessibilitySubrole {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilitySubrole"))
	return NSAccessibilitySubrole(foundation.NSStringFromID(rv).String())
	}

// Sets the specialized interface element type that the accessibility element
// represents.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitysubrole(_:)

func (o NSView) SetAccessibilitySubrole(accessibilitySubrole NSAccessibilitySubrole) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySubrole:"), objc.String(string(accessibilitySubrole)))
	}

// Returns the custom actions of the current accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitycustomactions()

func (o NSView) AccessibilityCustomActions() INSAccessibilityCustomAction {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityCustomActions"))
	return NSAccessibilityCustomActionFromID(rv)
	}

// Sets the custom actions of the current accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitycustomactions(_:)

func (o NSView) SetAccessibilityCustomActions(accessibilityCustomActions foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityCustomActions:"), accessibilityCustomActions)
	}

// Returns the custom rotors of the current accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitycustomrotors()

func (o NSView) AccessibilityCustomRotors() INSAccessibilityCustomRotor {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityCustomRotors"))
	return NSAccessibilityCustomRotorFromID(rv)
	}

// Sets the custom rotors of the current accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitycustomrotors(_:)

func (o NSView) SetAccessibilityCustomRotors(accessibilityCustomRotors foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityCustomRotors:"), accessibilityCustomRotors)
	}

// Returns the line number that contains the insertion point.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityinsertionpointlinenumber()

func (o NSView) AccessibilityInsertionPointLineNumber() int {
	
	rv := objc.Send[int](o.ID, objc.Sel("accessibilityInsertionPointLineNumber"))
	return rv
	}

// Sets the line number that contains the insertion point.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityinsertionpointlinenumber(_:)

func (o NSView) SetAccessibilityInsertionPointLineNumber(accessibilityInsertionPointLineNumber int) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityInsertionPointLineNumber:"), accessibilityInsertionPointLineNumber)
	}

// Returns the number of characters in the text.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitynumberofcharacters()

func (o NSView) AccessibilityNumberOfCharacters() int {
	
	rv := objc.Send[int](o.ID, objc.Sel("accessibilityNumberOfCharacters"))
	return rv
	}

// Sets the number of characters in the text.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitynumberofcharacters(_:)

func (o NSView) SetAccessibilityNumberOfCharacters(accessibilityNumberOfCharacters int) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityNumberOfCharacters:"), accessibilityNumberOfCharacters)
	}

// Returns the placeholder value for the accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityplaceholdervalue()

func (o NSView) AccessibilityPlaceholderValue() string {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityPlaceholderValue"))
	return foundation.NSStringFromID(rv).String()
	}

// Sets the placeholder value for the accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityplaceholdervalue(_:)

func (o NSView) SetAccessibilityPlaceholderValue(accessibilityPlaceholderValue string) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityPlaceholderValue:"), objc.String(accessibilityPlaceholderValue))
	}

// Returns the currently selected text.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityselectedtext()

func (o NSView) AccessibilitySelectedText() string {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilitySelectedText"))
	return foundation.NSStringFromID(rv).String()
	}

// Sets the currently selected text.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityselectedtext(_:)

func (o NSView) SetAccessibilitySelectedText(accessibilitySelectedText string) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySelectedText:"), objc.String(accessibilitySelectedText))
	}

// Returns the range of the currently selected text.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityselectedtextrange()

func (o NSView) AccessibilitySelectedTextRange() foundation.NSRange {
	
	rv := objc.Send[foundation.NSRange](o.ID, objc.Sel("accessibilitySelectedTextRange"))
	return rv
	}

// Sets the range of the currently selected text.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityselectedtextrange(_:)

func (o NSView) SetAccessibilitySelectedTextRange(accessibilitySelectedTextRange foundation.NSRange) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySelectedTextRange:"), accessibilitySelectedTextRange)
	}

// Returns an array of ranges for the currently selected text.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityselectedtextranges()

func (o NSView) AccessibilitySelectedTextRanges() foundation.NSValue {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilitySelectedTextRanges"))
	return foundation.NSValueFromID(rv)
	}

// Sets an array of ranges for the currently selected text.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityselectedtextranges(_:)

func (o NSView) SetAccessibilitySelectedTextRanges(accessibilitySelectedTextRanges foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySelectedTextRanges:"), accessibilitySelectedTextRanges)
	}

// Returns the range of characters that the accessibility element displays.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitysharedcharacterrange()

func (o NSView) AccessibilitySharedCharacterRange() foundation.NSRange {
	
	rv := objc.Send[foundation.NSRange](o.ID, objc.Sel("accessibilitySharedCharacterRange"))
	return rv
	}

// Sets the range of characters that the accessibility element displays.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitysharedcharacterrange(_:)

func (o NSView) SetAccessibilitySharedCharacterRange(accessibilitySharedCharacterRange foundation.NSRange) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySharedCharacterRange:"), accessibilitySharedCharacterRange)
	}

// Returns the other elements that share text with the accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitysharedtextuielements()

func (o NSView) AccessibilitySharedTextUIElements() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilitySharedTextUIElements"))
	return objectivec.Object{ID: rv}
	}

// Sets the other elements that share text with the accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitysharedtextuielements(_:)

func (o NSView) SetAccessibilitySharedTextUIElements(accessibilitySharedTextUIElements foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySharedTextUIElements:"), accessibilitySharedTextUIElements)
	}

// Returns the range of visible characters in the document.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityvisiblecharacterrange()

func (o NSView) AccessibilityVisibleCharacterRange() foundation.NSRange {
	
	rv := objc.Send[foundation.NSRange](o.ID, objc.Sel("accessibilityVisibleCharacterRange"))
	return rv
	}

// Sets the range of visible characters in the document.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityvisiblecharacterrange(_:)

func (o NSView) SetAccessibilityVisibleCharacterRange(accessibilityVisibleCharacterRange foundation.NSRange) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityVisibleCharacterRange:"), accessibilityVisibleCharacterRange)
	}

// Returns the substring for the specified range.
//
// range: A range of characters contained by the element.
//
// # Return Value
// 
// The substring specified by the given range.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityString(for:)

func (o NSView) AccessibilityStringForRange(range_ foundation.NSRange) string {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityStringForRange:"), range_)
	return foundation.NSStringFromID(rv).String()
	}

// Returns the attributed substring for the specified range of characters.
//
// range: The range of characters.
//
// # Return Value
// 
// An attributed string representing the specified characters.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityAttributedString(for:)

func (o NSView) AccessibilityAttributedStringForRange(range_ foundation.NSRange) foundation.NSAttributedString {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityAttributedStringForRange:"), range_)
	return foundation.NSAttributedStringFromID(rv)
	}

// Returns the rich text format (RTF) data that describes the specified range
// of characters.
//
// range: The range of characters.
//
// # Return Value
// 
// A data object containing an RTF representation of the specified characters.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityRTF(for:)

func (o NSView) AccessibilityRTFForRange(range_ foundation.NSRange) foundation.INSData {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityRTFForRange:"), range_)
	return foundation.NSDataFromID(rv)
	}

// Returns the rectangle that encloses the specified range of characters.
//
// range: The range of characters.
//
// # Return Value
// 
// The rectangle that encloses the specified characters.
//
// # Discussion
// 
// If the range crosses a line boundary, the returned rectangle fully encloses
// all the lines of characters.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityFrame(for:)

func (o NSView) AccessibilityFrameForRange(range_ foundation.NSRange) corefoundation.CGRect {
	
	rv := objc.Send[corefoundation.CGRect](o.ID, objc.Sel("accessibilityFrameForRange:"), range_)
	return rv
	}

// Returns the line number for the line that contains the specified character
// index.
//
// index: The index for a character.
//
// # Return Value
// 
// The line number for the line holding the specified character index.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityLine(for:)

func (o NSView) AccessibilityLineForIndex(index int) int {
	
	rv := objc.Send[int](o.ID, objc.Sel("accessibilityLineForIndex:"), index)
	return rv
	}

// Returns the range of characters for the glyph that includes the specified
// character.
//
// index: The specified character.
//
// # Return Value
// 
// The range of characters for the glyph.
//
// # Discussion
// 
// This value always includes the specified character but may include
// additional characters if that character is part of a multicharacter glyph.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityRange(for:)-6kv3

func (o NSView) AccessibilityRangeForIndex(index int) foundation.NSRange {
	
	rv := objc.Send[foundation.NSRange](o.ID, objc.Sel("accessibilityRangeForIndex:"), index)
	return rv
	}

// Returns a range of characters that all have the same style as the specified
// character.
//
// index: The index of the specified character.
//
// # Return Value
// 
// A range of characters with the same style as the specified character.
//
// # Discussion
// 
// This method returns a range of characters that meet two conditions: The
// range must include the specified character, and all the other characters in
// the range must match the specified character’s style. If none of the
// adjacent characters match the specified character’s style, the method
// returns only the specified character.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityStyleRange(for:)

func (o NSView) AccessibilityStyleRangeForIndex(index int) foundation.NSRange {
	
	rv := objc.Send[foundation.NSRange](o.ID, objc.Sel("accessibilityStyleRangeForIndex:"), index)
	return rv
	}

// Returns the range of characters in the specified line.
//
// line: The line number to be examined.
//
// # Return Value
// 
// The range of characters for the specified line number. If the line ends
// with a newline character, including the newline is preferred.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityRange(forLine:)

func (o NSView) AccessibilityRangeForLine(line int) foundation.NSRange {
	
	rv := objc.Send[foundation.NSRange](o.ID, objc.Sel("accessibilityRangeForLine:"), line)
	return rv
	}

// Returns the range of characters for the glyph at the specified point.
//
// point: A point in screen coordinates.
//
// # Return Value
// 
// The range of characters that make up the glyph at the given point.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityRange(for:)-1iudm

func (o NSView) AccessibilityRangeForPosition(point corefoundation.CGPoint) foundation.NSRange {
	
	rv := objc.Send[foundation.NSRange](o.ID, objc.Sel("accessibilityRangeForPosition:"), point)
	return rv
	}

// Returns the activation point for the user interface element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityactivationpoint()

func (o NSView) AccessibilityActivationPoint() corefoundation.CGPoint {
	
	rv := objc.Send[corefoundation.CGPoint](o.ID, objc.Sel("accessibilityActivationPoint"))
	return rv
	}

// Sets the activation point for the user interface element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityactivationpoint(_:)

func (o NSView) SetAccessibilityActivationPoint(accessibilityActivationPoint corefoundation.CGPoint) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityActivationPoint:"), accessibilityActivationPoint)
	}

// Returns the Boolean value that determines whether the accessibility
// element’s alternative UI is currently visible.
//
// See: /documentation/appkit/nsaccessibilityprotocol/isaccessibilityalternateuivisible()

func (o NSView) IsAccessibilityAlternateUIVisible() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityAlternateUIVisible"))
	return rv
	}

// Sets the Boolean value that determines whether the accessibility
// element’s alternative UI is currently visible.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityalternateuivisible(_:)

func (o NSView) SetAccessibilityAlternateUIVisible(accessibilityAlternateUIVisible bool) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityAlternateUIVisible:"), accessibilityAlternateUIVisible)
	}

// Returns the child accessibility element that represents the window’s
// cancel button.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitycancelbutton()

func (o NSView) AccessibilityCancelButton() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityCancelButton"))
	return objectivec.Object{ID: rv}
	}

// Sets the child accessibility element that represents the window’s cancel
// button.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitycancelbutton(_:)

func (o NSView) SetAccessibilityCancelButton(accessibilityCancelButton objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityCancelButton:"), accessibilityCancelButton)
	}

// Returns the child accessibility element that represents the window’s
// close button.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityclosebutton()

func (o NSView) AccessibilityCloseButton() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityCloseButton"))
	return objectivec.Object{ID: rv}
	}

// Sets the child accessibility element that represents the window’s close
// button.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityclosebutton(_:)

func (o NSView) SetAccessibilityCloseButton(accessibilityCloseButton objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityCloseButton:"), accessibilityCloseButton)
	}

// Returns the child accessibility element that represents the window’s
// default button.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitydefaultbutton()

func (o NSView) AccessibilityDefaultButton() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityDefaultButton"))
	return objectivec.Object{ID: rv}
	}

// Sets the child accessibility element that represents the window’s default
// button.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitydefaultbutton(_:)

func (o NSView) SetAccessibilityDefaultButton(accessibilityDefaultButton objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityDefaultButton:"), accessibilityDefaultButton)
	}

// Returns the child accessibility element that represents the window’s
// full-screen button.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityfullscreenbutton()

func (o NSView) AccessibilityFullScreenButton() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityFullScreenButton"))
	return objectivec.Object{ID: rv}
	}

// Sets the child accessibility element that represents the window’s
// full-screen button.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityfullscreenbutton(_:)

func (o NSView) SetAccessibilityFullScreenButton(accessibilityFullScreenButton objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityFullScreenButton:"), accessibilityFullScreenButton)
	}

// Returns the child accessibility element that represents the window’s grow
// area.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitygrowarea()

func (o NSView) AccessibilityGrowArea() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityGrowArea"))
	return objectivec.Object{ID: rv}
	}

// Sets the child accessibility element that represents the window’s grow
// area.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitygrowarea(_:)

func (o NSView) SetAccessibilityGrowArea(accessibilityGrowArea objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityGrowArea:"), accessibilityGrowArea)
	}

// Returns a Boolean value that determines whether the window is the app’s
// main window.
//
// See: /documentation/appkit/nsaccessibilityprotocol/isaccessibilitymain()

func (o NSView) IsAccessibilityMain() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityMain"))
	return rv
	}

// Sets a Boolean value that determines whether the window is the app’s main
// window.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitymain(_:)

func (o NSView) SetAccessibilityMain(accessibilityMain bool) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityMain:"), accessibilityMain)
	}

// Returns the child accessibility element that represents the window’s
// minimize button.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityminimizebutton()

func (o NSView) AccessibilityMinimizeButton() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityMinimizeButton"))
	return objectivec.Object{ID: rv}
	}

// Sets the child accessibility element that represents the window’s
// minimize button.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityminimizebutton(_:)

func (o NSView) SetAccessibilityMinimizeButton(accessibilityMinimizeButton objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityMinimizeButton:"), accessibilityMinimizeButton)
	}

// Returns the Boolean value that determines whether the window is in a
// minimized state.
//
// See: /documentation/appkit/nsaccessibilityprotocol/isaccessibilityminimized()

func (o NSView) IsAccessibilityMinimized() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityMinimized"))
	return rv
	}

// Sets the Boolean value that determines whether the window is in a minimized
// state.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityminimized(_:)

func (o NSView) SetAccessibilityMinimized(accessibilityMinimized bool) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityMinimized:"), accessibilityMinimized)
	}

// Returns a Boolean value that determines whether the window is modal.
//
// See: /documentation/appkit/nsaccessibilityprotocol/isaccessibilitymodal()

func (o NSView) IsAccessibilityModal() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityModal"))
	return rv
	}

// Sets a Boolean value that determines whether the window is modal.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitymodal(_:)

func (o NSView) SetAccessibilityModal(accessibilityModal bool) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityModal:"), accessibilityModal)
	}

// Returns the child accessibility element that represents the window’s
// proxy icon.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityproxy()

func (o NSView) AccessibilityProxy() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityProxy"))
	return objectivec.Object{ID: rv}
	}

// Sets the child accessibility element that represents the window’s proxy
// icon.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityproxy(_:)

func (o NSView) SetAccessibilityProxy(accessibilityProxy objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityProxy:"), accessibilityProxy)
	}

// Returns the menu currently displaying for the accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityshownmenu()

func (o NSView) AccessibilityShownMenu() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityShownMenu"))
	return objectivec.Object{ID: rv}
	}

// Sets the menu currently displaying for the accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityshownmenu(_:)

func (o NSView) SetAccessibilityShownMenu(accessibilityShownMenu objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityShownMenu:"), accessibilityShownMenu)
	}

// Returns the child accessibility element that represents the window’s
// toolbar button.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitytoolbarbutton()

func (o NSView) AccessibilityToolbarButton() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityToolbarButton"))
	return objectivec.Object{ID: rv}
	}

// Sets the child accessibility element that represents the window’s toolbar
// button.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitytoolbarbutton(_:)

func (o NSView) SetAccessibilityToolbarButton(accessibilityToolbarButton objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityToolbarButton:"), accessibilityToolbarButton)
	}

// Returns the window that contains the accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitywindow()

func (o NSView) AccessibilityWindow() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityWindow"))
	return objectivec.Object{ID: rv}
	}

// Sets the window that contains the accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitywindow(_:)

func (o NSView) SetAccessibilityWindow(accessibilityWindow objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityWindow:"), accessibilityWindow)
	}

// Returns the child accessibility element that represents the window’s zoom
// button.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityzoombutton()

func (o NSView) AccessibilityZoomButton() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityZoomButton"))
	return objectivec.Object{ID: rv}
	}

// Sets the child accessibility element that represents the window’s zoom
// button.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityzoombutton(_:)

func (o NSView) SetAccessibilityZoomButton(accessibilityZoomButton objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityZoomButton:"), accessibilityZoomButton)
	}

// Returns the icon for the app’s menu bar extra.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityextrasmenubar()

func (o NSView) AccessibilityExtrasMenuBar() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityExtrasMenuBar"))
	return objectivec.Object{ID: rv}
	}

// Sets the icon for the app’s menu bar extra.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityextrasmenubar(_:)

func (o NSView) SetAccessibilityExtrasMenuBar(accessibilityExtrasMenuBar objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityExtrasMenuBar:"), accessibilityExtrasMenuBar)
	}

// Returns a Boolean value that determines whether the app is the frontmost
// app.
//
// See: /documentation/appkit/nsaccessibilityprotocol/isaccessibilityfrontmost()

func (o NSView) IsAccessibilityFrontmost() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityFrontmost"))
	return rv
	}

// Sets a Boolean value that determines whether the app is the frontmost app.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityfrontmost(_:)

func (o NSView) SetAccessibilityFrontmost(accessibilityFrontmost bool) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityFrontmost:"), accessibilityFrontmost)
	}

// Returns a Boolean value that determines whether the app is in a hidden
// state.
//
// See: /documentation/appkit/nsaccessibilityprotocol/isaccessibilityhidden()

func (o NSView) IsAccessibilityHidden() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityHidden"))
	return rv
	}

// Sets a Boolean value that determines whether the app is in a hidden state.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityhidden(_:)

func (o NSView) SetAccessibilityHidden(accessibilityHidden bool) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityHidden:"), accessibilityHidden)
	}

// Returns the app’s main window.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitymainwindow()

func (o NSView) AccessibilityMainWindow() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityMainWindow"))
	return objectivec.Object{ID: rv}
	}

// Sets the app’s main window.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitymainwindow(_:)

func (o NSView) SetAccessibilityMainWindow(accessibilityMainWindow objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityMainWindow:"), accessibilityMainWindow)
	}

// Returns the app’s menu bar.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitymenubar()

func (o NSView) AccessibilityMenuBar() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityMenuBar"))
	return objectivec.Object{ID: rv}
	}

// Sets the app’s menu bar.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitymenubar(_:)

func (o NSView) SetAccessibilityMenuBar(accessibilityMenuBar objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityMenuBar:"), accessibilityMenuBar)
	}

// Returns an array that contains all the app’s windows.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitywindows()

func (o NSView) AccessibilityWindows() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityWindows"))
	return objectivec.Object{ID: rv}
	}

// Sets the array that contains all the app’s windows.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitywindows(_:)

func (o NSView) SetAccessibilityWindows(accessibilityWindows foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityWindows:"), accessibilityWindows)
	}

// Returns the number of columns in the accessibility element’s grid.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitycolumncount()

func (o NSView) AccessibilityColumnCount() int {
	
	rv := objc.Send[int](o.ID, objc.Sel("accessibilityColumnCount"))
	return rv
	}

// Sets the number of columns in the accessibility element’s grid.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitycolumncount(_:)

func (o NSView) SetAccessibilityColumnCount(accessibilityColumnCount int) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityColumnCount:"), accessibilityColumnCount)
	}

// Returns a Boolean value that determines whether the accessibility
// element’s grid is in row major order or in column major order.
//
// See: /documentation/appkit/nsaccessibilityprotocol/isaccessibilityorderedbyrow()

func (o NSView) IsAccessibilityOrderedByRow() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityOrderedByRow"))
	return rv
	}

// Sets a Boolean value that determines whether the element’s grid is in row
// major order or in column major order.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityorderedbyrow(_:)

func (o NSView) SetAccessibilityOrderedByRow(accessibilityOrderedByRow bool) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityOrderedByRow:"), accessibilityOrderedByRow)
	}

// Returns the number of rows in the accessibility element’s grid.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityrowcount()

func (o NSView) AccessibilityRowCount() int {
	
	rv := objc.Send[int](o.ID, objc.Sel("accessibilityRowCount"))
	return rv
	}

// Sets the number of rows in the accessibility element’s grid.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityrowcount(_:)

func (o NSView) SetAccessibilityRowCount(accessibilityRowCount int) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityRowCount:"), accessibilityRowCount)
	}

// Returns the horizontal scroll bar for the scroll view.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityhorizontalscrollbar()

func (o NSView) AccessibilityHorizontalScrollBar() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityHorizontalScrollBar"))
	return objectivec.Object{ID: rv}
	}

// Sets the horizontal scroll bar for the scroll view.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityhorizontalscrollbar(_:)

func (o NSView) SetAccessibilityHorizontalScrollBar(accessibilityHorizontalScrollBar objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityHorizontalScrollBar:"), accessibilityHorizontalScrollBar)
	}

// Returns the vertical scroll bar for the scroll view.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityverticalscrollbar()

func (o NSView) AccessibilityVerticalScrollBar() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityVerticalScrollBar"))
	return objectivec.Object{ID: rv}
	}

// Sets the vertical scroll bar for the scroll view.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityverticalscrollbar(_:)

func (o NSView) SetAccessibilityVerticalScrollBar(accessibilityVerticalScrollBar objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityVerticalScrollBar:"), accessibilityVerticalScrollBar)
	}

// Returns the column header accessibility elements for the table or outline.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitycolumnheaderuielements()

func (o NSView) AccessibilityColumnHeaderUIElements() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityColumnHeaderUIElements"))
	return objectivec.Object{ID: rv}
	}

// Sets the column header accessibility elements for the table or outline.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitycolumnheaderuielements(_:)

func (o NSView) SetAccessibilityColumnHeaderUIElements(accessibilityColumnHeaderUIElements foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityColumnHeaderUIElements:"), accessibilityColumnHeaderUIElements)
	}

// Returns the column accessibility elements for the table or outline.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitycolumns()

func (o NSView) AccessibilityColumns() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityColumns"))
	return objectivec.Object{ID: rv}
	}

// Sets the column accessibility elements for the table or outline.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitycolumns(_:)

func (o NSView) SetAccessibilityColumns(accessibilityColumns foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityColumns:"), accessibilityColumns)
	}

// Returns the column titles for the accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitycolumntitles()

func (o NSView) AccessibilityColumnTitles() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityColumnTitles"))
	return objectivec.Object{ID: rv}
	}

// Sets the column titles for the accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitycolumntitles(_:)

func (o NSView) SetAccessibilityColumnTitles(accessibilityColumnTitles foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityColumnTitles:"), accessibilityColumnTitles)
	}

// Returns a Boolean value that determines whether the accessibility element
// is in an expanded state.
//
// See: /documentation/appkit/nsaccessibilityprotocol/isaccessibilityexpanded()

func (o NSView) IsAccessibilityExpanded() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityExpanded"))
	return rv
	}

// Sets a Boolean value that determines whether accessibility element is in an
// expanded state.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityexpanded(_:)

func (o NSView) SetAccessibilityExpanded(accessibilityExpanded bool) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityExpanded:"), accessibilityExpanded)
	}

// Returns the header for the table view.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityheader()

func (o NSView) AccessibilityHeader() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityHeader"))
	return objectivec.Object{ID: rv}
	}

// Sets the header for the table view.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityheader(_:)

func (o NSView) SetAccessibilityHeader(accessibilityHeader objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityHeader:"), accessibilityHeader)
	}

// Returns the index of the row or column that the accessibility element
// represents.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityindex()

func (o NSView) AccessibilityIndex() int {
	
	rv := objc.Send[int](o.ID, objc.Sel("accessibilityIndex"))
	return rv
	}

// Sets the index of the row or column that the accessibility element
// represents.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityindex(_:)

func (o NSView) SetAccessibilityIndex(accessibilityIndex int) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityIndex:"), accessibilityIndex)
	}

// Returns the row header accessibility elements for the table or outline.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityrowheaderuielements()

func (o NSView) AccessibilityRowHeaderUIElements() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityRowHeaderUIElements"))
	return objectivec.Object{ID: rv}
	}

// Sets the row header accessibility elements for the table or outline.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityrowheaderuielements(_:)

func (o NSView) SetAccessibilityRowHeaderUIElements(accessibilityRowHeaderUIElements foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityRowHeaderUIElements:"), accessibilityRowHeaderUIElements)
	}

// Returns the row accessibility elements for the table or outline.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityrows()

func (o NSView) AccessibilityRows() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityRows"))
	return objectivec.Object{ID: rv}
	}

// Sets the row accessibility elements for the table or outline.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityrows(_:)

func (o NSView) SetAccessibilityRows(accessibilityRows foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityRows:"), accessibilityRows)
	}

// Returns the currently selected columns for the table or outline.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityselectedcolumns()

func (o NSView) AccessibilitySelectedColumns() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilitySelectedColumns"))
	return objectivec.Object{ID: rv}
	}

// Sets the currently selected columns for the table or outline.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityselectedcolumns(_:)

func (o NSView) SetAccessibilitySelectedColumns(accessibilitySelectedColumns foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySelectedColumns:"), accessibilitySelectedColumns)
	}

// Returns the currently selected rows for the table or outline.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityselectedrows()

func (o NSView) AccessibilitySelectedRows() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilitySelectedRows"))
	return objectivec.Object{ID: rv}
	}

// Sets the currently selected rows for the table or outline.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityselectedrows(_:)

func (o NSView) SetAccessibilitySelectedRows(accessibilitySelectedRows foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySelectedRows:"), accessibilitySelectedRows)
	}

// Returns the accessibility element’s sort direction.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitysortdirection()

func (o NSView) AccessibilitySortDirection() NSAccessibilitySortDirection {
	
	rv := objc.Send[NSAccessibilitySortDirection](o.ID, objc.Sel("accessibilitySortDirection"))
	return rv
	}

// Sets the accessibility element’s sort direction.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitysortdirection(_:)

func (o NSView) SetAccessibilitySortDirection(accessibilitySortDirection NSAccessibilitySortDirection) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySortDirection:"), accessibilitySortDirection)
	}

// Returns the visible columns for the table or outline.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityvisiblecolumns()

func (o NSView) AccessibilityVisibleColumns() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityVisibleColumns"))
	return objectivec.Object{ID: rv}
	}

// Sets the visible columns for the table or outline.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityvisiblecolumns(_:)

func (o NSView) SetAccessibilityVisibleColumns(accessibilityVisibleColumns foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityVisibleColumns:"), accessibilityVisibleColumns)
	}

// Returns the visible rows for the table or outline.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityvisiblerows()

func (o NSView) AccessibilityVisibleRows() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityVisibleRows"))
	return objectivec.Object{ID: rv}
	}

// Sets the visible rows for the table or outline.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityvisiblerows(_:)

func (o NSView) SetAccessibilityVisibleRows(accessibilityVisibleRows foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityVisibleRows:"), accessibilityVisibleRows)
	}

// Returns a Boolean value that determines whether the row is disclosing other
// rows.
//
// See: /documentation/appkit/nsaccessibilityprotocol/isaccessibilitydisclosed()

func (o NSView) IsAccessibilityDisclosed() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityDisclosed"))
	return rv
	}

// Sets a Boolean value that determines whether the row is disclosing other
// rows.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitydisclosed(_:)

func (o NSView) SetAccessibilityDisclosed(accessibilityDisclosed bool) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityDisclosed:"), accessibilityDisclosed)
	}

// Returns the row disclosing the current row.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitydisclosedbyrow()

func (o NSView) AccessibilityDisclosedByRow() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityDisclosedByRow"))
	return objectivec.Object{ID: rv}
	}

// Sets the row disclosing the current row.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitydisclosedbyrow(_:)

func (o NSView) SetAccessibilityDisclosedByRow(accessibilityDisclosedByRow objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityDisclosedByRow:"), accessibilityDisclosedByRow)
	}

// Returns the rows that the current row discloses.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitydisclosedrows()

func (o NSView) AccessibilityDisclosedRows() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityDisclosedRows"))
	return objectivec.Object{ID: rv}
	}

// Sets the rows that the current row discloses.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitydisclosedrows(_:)

func (o NSView) SetAccessibilityDisclosedRows(accessibilityDisclosedRows objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityDisclosedRows:"), accessibilityDisclosedRows)
	}

// Returns the indention level for the row.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitydisclosurelevel()

func (o NSView) AccessibilityDisclosureLevel() int {
	
	rv := objc.Send[int](o.ID, objc.Sel("accessibilityDisclosureLevel"))
	return rv
	}

// Sets the indention level for the row.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitydisclosurelevel(_:)

func (o NSView) SetAccessibilityDisclosureLevel(accessibilityDisclosureLevel int) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityDisclosureLevel:"), accessibilityDisclosureLevel)
	}

// Returns the column index range of the cell.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitycolumnindexrange()

func (o NSView) AccessibilityColumnIndexRange() foundation.NSRange {
	
	rv := objc.Send[foundation.NSRange](o.ID, objc.Sel("accessibilityColumnIndexRange"))
	return rv
	}

// Sets the column index range of the cell.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitycolumnindexrange(_:)

func (o NSView) SetAccessibilityColumnIndexRange(accessibilityColumnIndexRange foundation.NSRange) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityColumnIndexRange:"), accessibilityColumnIndexRange)
	}

// Returns the row index range of the cell.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityrowindexrange()

func (o NSView) AccessibilityRowIndexRange() foundation.NSRange {
	
	rv := objc.Send[foundation.NSRange](o.ID, objc.Sel("accessibilityRowIndexRange"))
	return rv
	}

// Sets the row index range of the cell.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityrowindexrange(_:)

func (o NSView) SetAccessibilityRowIndexRange(accessibilityRowIndexRange foundation.NSRange) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityRowIndexRange:"), accessibilityRowIndexRange)
	}

// Returns the currently selected cells for the table.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityselectedcells()

func (o NSView) AccessibilitySelectedCells() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilitySelectedCells"))
	return objectivec.Object{ID: rv}
	}

// Sets the currently selected cells for the table.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityselectedcells(_:)

func (o NSView) SetAccessibilitySelectedCells(accessibilitySelectedCells foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySelectedCells:"), accessibilitySelectedCells)
	}

// Returns the visible cells for the table.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityvisiblecells()

func (o NSView) AccessibilityVisibleCells() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityVisibleCells"))
	return objectivec.Object{ID: rv}
	}

// Sets the visible cells for the table.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityvisiblecells(_:)

func (o NSView) SetAccessibilityVisibleCells(accessibilityVisibleCells foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityVisibleCells:"), accessibilityVisibleCells)
	}

// Returns the cell at the specified column and row.
//
// column: The column index.
//
// row: The row index.
//
// # Return Value
// 
// The cell specified by the column and row indexes.
//
// # Discussion
// 
// This property is required for all elements that function as cell-based
// tables.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityCell(forColumn:row:)

func (o NSView) AccessibilityCellForColumnRow(column int, row int) objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityCellForColumn:row:"), column, row)
	return objectivec.Object{ID: rv}
	}

// Returns the drag handle elements for the layout item element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityhandles()

func (o NSView) AccessibilityHandles() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityHandles"))
	return objectivec.Object{ID: rv}
	}

// Sets the drag handle accessibility elements for the layout item element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityhandles(_:)

func (o NSView) SetAccessibilityHandles(accessibilityHandles foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityHandles:"), accessibilityHandles)
	}

// Returns the units that the layout area uses for horizontal values.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityhorizontalunits()

func (o NSView) AccessibilityHorizontalUnits() NSAccessibilityUnits {
	
	rv := objc.Send[NSAccessibilityUnits](o.ID, objc.Sel("accessibilityHorizontalUnits"))
	return rv
	}

// Sets the units that the layout area uses for horizontal values.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityhorizontalunits(_:)

func (o NSView) SetAccessibilityHorizontalUnits(accessibilityHorizontalUnits NSAccessibilityUnits) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityHorizontalUnits:"), accessibilityHorizontalUnits)
	}

// Returns the description of the layout area’s horizontal units.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityhorizontalunitdescription()

func (o NSView) AccessibilityHorizontalUnitDescription() string {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityHorizontalUnitDescription"))
	return foundation.NSStringFromID(rv).String()
	}

// Sets the description of the layout area’s horizontal units.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityhorizontalunitdescription(_:)

func (o NSView) SetAccessibilityHorizontalUnitDescription(accessibilityHorizontalUnitDescription string) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityHorizontalUnitDescription:"), objc.String(accessibilityHorizontalUnitDescription))
	}

// Returns the units that the layout area uses for vertical values.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityverticalunits()

func (o NSView) AccessibilityVerticalUnits() NSAccessibilityUnits {
	
	rv := objc.Send[NSAccessibilityUnits](o.ID, objc.Sel("accessibilityVerticalUnits"))
	return rv
	}

// Sets the units that the layout area uses for vertical values.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityverticalunits(_:)

func (o NSView) SetAccessibilityVerticalUnits(accessibilityVerticalUnits NSAccessibilityUnits) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityVerticalUnits:"), accessibilityVerticalUnits)
	}

// Returns the description of the layout area’s vertical units.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityverticalunitdescription()

func (o NSView) AccessibilityVerticalUnitDescription() string {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityVerticalUnitDescription"))
	return foundation.NSStringFromID(rv).String()
	}

// Sets the description of the layout area’s vertical units.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityverticalunitdescription(_:)

func (o NSView) SetAccessibilityVerticalUnitDescription(accessibilityVerticalUnitDescription string) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityVerticalUnitDescription:"), objc.String(accessibilityVerticalUnitDescription))
	}

// Converts the provided point in screen coordinates to a point in the layout
// area’s coordinate system.
//
// point: A point in the screen’s coordinate system.
//
// # Return Value
// 
// A point in the layout area’s coordinate system.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityLayoutPoint(forScreenPoint:)

func (o NSView) AccessibilityLayoutPointForScreenPoint(point corefoundation.CGPoint) corefoundation.CGPoint {
	
	rv := objc.Send[corefoundation.CGPoint](o.ID, objc.Sel("accessibilityLayoutPointForScreenPoint:"), point)
	return rv
	}

// Converts the provided size in screen coordinates to a size in the layout
// area’s coordinate system.
//
// size: A size in the screen’s coordinate system.
//
// # Return Value
// 
// A size in the layout area’s coordinate system.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityLayoutSize(forScreenSize:)

func (o NSView) AccessibilityLayoutSizeForScreenSize(size corefoundation.CGSize) corefoundation.CGSize {
	
	rv := objc.Send[corefoundation.CGSize](o.ID, objc.Sel("accessibilityLayoutSizeForScreenSize:"), size)
	return rv
	}

// Converts the provided point in the layout area’s coordinates to a point
// in the screen’s coordinate system.
//
// point: A point in the layout area’s coordinate system.
//
// # Return Value
// 
// A point in the screen’s coordinate system.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityScreenPoint(forLayoutPoint:)

func (o NSView) AccessibilityScreenPointForLayoutPoint(point corefoundation.CGPoint) corefoundation.CGPoint {
	
	rv := objc.Send[corefoundation.CGPoint](o.ID, objc.Sel("accessibilityScreenPointForLayoutPoint:"), point)
	return rv
	}

// Converts the provided size in the layout area’s coordinates to a size in
// the screen’s coordinate system.
//
// size: A size in the layout area’s coordinate system.
//
// # Return Value
// 
// A size in the screen’s coordinate system.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityScreenSize(forLayoutSize:)

func (o NSView) AccessibilityScreenSizeForLayoutSize(size corefoundation.CGSize) corefoundation.CGSize {
	
	rv := objc.Send[corefoundation.CGSize](o.ID, objc.Sel("accessibilityScreenSizeForLayoutSize:"), size)
	return rv
	}

// Returns the allowed values for the slider accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityallowedvalues()

func (o NSView) AccessibilityAllowedValues() foundation.NSNumber {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityAllowedValues"))
	return foundation.NSNumberFromID(rv)
	}

// Sets the allowed values for the slider accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityallowedvalues(_:)

func (o NSView) SetAccessibilityAllowedValues(accessibilityAllowedValues foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityAllowedValues:"), accessibilityAllowedValues)
	}

// Returns the child label elements for the slider accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitylabeluielements()

func (o NSView) AccessibilityLabelUIElements() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityLabelUIElements"))
	return objectivec.Object{ID: rv}
	}

// Sets the child label elements for the slider accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitylabeluielements(_:)

func (o NSView) SetAccessibilityLabelUIElements(accessibilityLabelUIElements foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityLabelUIElements:"), accessibilityLabelUIElements)
	}

// Returns the value of the label accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitylabelvalue()

func (o NSView) AccessibilityLabelValue() float64 {
	
	rv := objc.Send[float64](o.ID, objc.Sel("accessibilityLabelValue"))
	return rv
	}

// Sets the value of the label accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitylabelvalue(_:)

func (o NSView) SetAccessibilityLabelValue(accessibilityLabelValue float64) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityLabelValue:"), accessibilityLabelValue)
	}

// Returns the contents that follow the divider accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitynextcontents()

func (o NSView) AccessibilityNextContents() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityNextContents"))
	return objectivec.Object{ID: rv}
	}

// Sets the contents that follow the divider accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitynextcontents(_:)

func (o NSView) SetAccessibilityNextContents(accessibilityNextContents foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityNextContents:"), accessibilityNextContents)
	}

// Returns the contents that precede the divider accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitypreviouscontents()

func (o NSView) AccessibilityPreviousContents() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityPreviousContents"))
	return objectivec.Object{ID: rv}
	}

// Sets the contents that precede the divider accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitypreviouscontents(_:)

func (o NSView) SetAccessibilityPreviousContents(accessibilityPreviousContents foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityPreviousContents:"), accessibilityPreviousContents)
	}

// Returns an array that contains the views and splitter bar from the split
// view.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitysplitters()

func (o NSView) AccessibilitySplitters() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilitySplitters"))
	return objectivec.Object{ID: rv}
	}

// Sets the array that contains the views and splitter bar from the split
// view.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitysplitters(_:)

func (o NSView) SetAccessibilitySplitters(accessibilitySplitters foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySplitters:"), accessibilitySplitters)
	}

// Returns the overflow button for the toolbar.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityoverflowbutton()

func (o NSView) AccessibilityOverflowButton() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityOverflowButton"))
	return objectivec.Object{ID: rv}
	}

// Sets the overflow button for the toolbar.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityoverflowbutton(_:)

func (o NSView) SetAccessibilityOverflowButton(accessibilityOverflowButton objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityOverflowButton:"), accessibilityOverflowButton)
	}

// Returns the tab accessibility elements for the tab view.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitytabs()

func (o NSView) AccessibilityTabs() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityTabs"))
	return objectivec.Object{ID: rv}
	}

// Sets the tab accessibility elements for the tab view.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitytabs(_:)

func (o NSView) SetAccessibilityTabs(accessibilityTabs foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityTabs:"), accessibilityTabs)
	}

// Returns the user interface element that functions as a marker group for the
// ruler accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitymarkergroupuielement()

func (o NSView) AccessibilityMarkerGroupUIElement() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityMarkerGroupUIElement"))
	return objectivec.Object{ID: rv}
	}

// Sets the user interface element that functions as a marker group for the
// ruler accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitymarkergroupuielement(_:)

func (o NSView) SetAccessibilityMarkerGroupUIElement(accessibilityMarkerGroupUIElement objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityMarkerGroupUIElement:"), accessibilityMarkerGroupUIElement)
	}

// Returns the human-readable description of the marker type.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitymarkertypedescription()

func (o NSView) AccessibilityMarkerTypeDescription() string {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityMarkerTypeDescription"))
	return foundation.NSStringFromID(rv).String()
	}

// Sets the human-readable description of the marker type.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitymarkertypedescription(_:)

func (o NSView) SetAccessibilityMarkerTypeDescription(accessibilityMarkerTypeDescription string) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityMarkerTypeDescription:"), objc.String(accessibilityMarkerTypeDescription))
	}

// Returns the array of marker accessibility elements for the ruler.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitymarkeruielements()

func (o NSView) AccessibilityMarkerUIElements() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityMarkerUIElements"))
	return objectivec.Object{ID: rv}
	}

// Sets the array of marker accessibility elements for the ruler.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitymarkeruielements(_:)

func (o NSView) SetAccessibilityMarkerUIElements(accessibilityMarkerUIElements foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityMarkerUIElements:"), accessibilityMarkerUIElements)
	}

// Returns the marker values for the ruler.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitymarkervalues()

func (o NSView) AccessibilityMarkerValues() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityMarkerValues"))
	return objectivec.Object{ID: rv}
	}

// Sets the marker values for the ruler.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitymarkervalues(_:)

func (o NSView) SetAccessibilityMarkerValues(accessibilityMarkerValues objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityMarkerValues:"), accessibilityMarkerValues)
	}

// Returns the type of markers for the ruler.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityrulermarkertype()

func (o NSView) AccessibilityRulerMarkerType() NSAccessibilityRulerMarkerType {
	
	rv := objc.Send[NSAccessibilityRulerMarkerType](o.ID, objc.Sel("accessibilityRulerMarkerType"))
	return rv
	}

// Sets the type of markers for the ruler.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityrulermarkertype(_:)

func (o NSView) SetAccessibilityRulerMarkerType(accessibilityRulerMarkerType NSAccessibilityRulerMarkerType) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityRulerMarkerType:"), accessibilityRulerMarkerType)
	}

// Returns the units for the ruler.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityunits()

func (o NSView) AccessibilityUnits() NSAccessibilityUnits {
	
	rv := objc.Send[NSAccessibilityUnits](o.ID, objc.Sel("accessibilityUnits"))
	return rv
	}

// Sets the units used for the ruler.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityunits(_:)

func (o NSView) SetAccessibilityUnits(accessibilityUnits NSAccessibilityUnits) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityUnits:"), accessibilityUnits)
	}

// Returns the human-readable description of the ruler’s units.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityunitdescription()

func (o NSView) AccessibilityUnitDescription() string {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityUnitDescription"))
	return foundation.NSStringFromID(rv).String()
	}

// Sets the human-readable description of the ruler’s units.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityunitdescription(_:)

func (o NSView) SetAccessibilityUnitDescription(accessibilityUnitDescription string) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityUnitDescription:"), objc.String(accessibilityUnitDescription))
	}

// Returns the URL for the file that the accessibility element represents.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitydocument()

func (o NSView) AccessibilityDocument() string {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityDocument"))
	return foundation.NSStringFromID(rv).String()
	}

// Sets the URL for the file that the accessibility element represents.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitydocument(_:)

func (o NSView) SetAccessibilityDocument(accessibilityDocument string) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityDocument:"), objc.String(accessibilityDocument))
	}

// Returns a Boolean value that indicates whether the accessibility element is
// in an edited state.
//
// See: /documentation/appkit/nsaccessibilityprotocol/isaccessibilityedited()

func (o NSView) IsAccessibilityEdited() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityEdited"))
	return rv
	}

// Sets a Boolean value that indicates whether the accessibility element is in
// an edited state.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityedited(_:)

func (o NSView) SetAccessibilityEdited(accessibilityEdited bool) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityEdited:"), accessibilityEdited)
	}

// Returns the filename for the file that the accessibility element
// represents.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityfilename()

func (o NSView) AccessibilityFilename() string {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityFilename"))
	return foundation.NSStringFromID(rv).String()
	}

// Sets the filename for the file that the accessibility element represents.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityfilename(_:)

func (o NSView) SetAccessibilityFilename(accessibilityFilename string) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityFilename:"), objc.String(accessibilityFilename))
	}

// Returns the elements that have links with the accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitylinkeduielements()

func (o NSView) AccessibilityLinkedUIElements() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityLinkedUIElements"))
	return objectivec.Object{ID: rv}
	}

// Sets the elements that have links with the accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitylinkeduielements(_:)

func (o NSView) SetAccessibilityLinkedUIElements(accessibilityLinkedUIElements foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityLinkedUIElements:"), accessibilityLinkedUIElements)
	}

// Returns the list of elements that the accessibility element is a title for.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityservesastitleforuielements()

func (o NSView) AccessibilityServesAsTitleForUIElements() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityServesAsTitleForUIElements"))
	return objectivec.Object{ID: rv}
	}

// Sets the list of elements that the accessibility element is a title for.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityservesastitleforuielements(_:)

func (o NSView) SetAccessibilityServesAsTitleForUIElements(accessibilityServesAsTitleForUIElements foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityServesAsTitleForUIElements:"), accessibilityServesAsTitleForUIElements)
	}

// Returns the static text element that represents the accessibility
// element’s title.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitytitleuielement()

func (o NSView) AccessibilityTitleUIElement() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityTitleUIElement"))
	return objectivec.Object{ID: rv}
	}

// Sets the static text element that represents the accessibility element’s
// title.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitytitleuielement(_:)

func (o NSView) SetAccessibilityTitleUIElement(accessibilityTitleUIElement objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityTitleUIElement:"), accessibilityTitleUIElement)
	}

// Returns the clear button for the search field.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityclearbutton()

func (o NSView) AccessibilityClearButton() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityClearButton"))
	return objectivec.Object{ID: rv}
	}

// Sets the clear button for the search field.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityclearbutton(_:)

func (o NSView) SetAccessibilityClearButton(accessibilityClearButton objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityClearButton:"), accessibilityClearButton)
	}

// Returns the search button for the search field.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitysearchbutton()

func (o NSView) AccessibilitySearchButton() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilitySearchButton"))
	return objectivec.Object{ID: rv}
	}

// Sets the search button for the search field.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitysearchbutton(_:)

func (o NSView) SetAccessibilitySearchButton(accessibilitySearchButton objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySearchButton:"), accessibilitySearchButton)
	}

// Returns the search menu for the search field.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitysearchmenu()

func (o NSView) AccessibilitySearchMenu() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilitySearchMenu"))
	return objectivec.Object{ID: rv}
	}

// Sets the search menu for the search field.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitysearchmenu(_:)

func (o NSView) SetAccessibilitySearchMenu(accessibilitySearchMenu objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySearchMenu:"), accessibilitySearchMenu)
	}

// Cancels the current operation.
//
// # Return Value
// 
// [true] if the action was successfully triggered; otherwise, [false]. This
// method does not indicate the success or failure of the action, just the
// fact that the action was successfully triggered.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityPerformCancel()

func (o NSView) AccessibilityPerformCancel() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("accessibilityPerformCancel"))
	return rv
	}

// Simulates pressing Return in the accessibility element.
//
// # Return Value
// 
// [true] if the action was successfully triggered; otherwise, [false]. This
// method does not indicate the success or failure of the action, just the
// fact that the action was successfully triggered.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// Use this method on elements that take keyboard input, such as a text field.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityPerformConfirm()

func (o NSView) AccessibilityPerformConfirm() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("accessibilityPerformConfirm"))
	return rv
	}

// Selects the accessibility element.
//
// # Return Value
// 
// [true] if the action was successfully triggered; otherwise, [false]. This
// method does not indicate the success or failure of the action, just the
// fact that the action was successfully triggered.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// Use this method on selectable elements, such as a menu item.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityPerformPick()

func (o NSView) AccessibilityPerformPick() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("accessibilityPerformPick"))
	return rv
	}

// Simulates clicking the accessibility element.
//
// # Return Value
// 
// [true] if the action was successfully triggered; otherwise, [false]. This
// method does not indicate the success or failure of the action, just the
// fact that the action was successfully triggered.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// Use this method on elements that behave like buttons.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityPerformPress()

func (o NSView) AccessibilityPerformPress() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("accessibilityPerformPress"))
	return rv
	}

// Displays the accessibility element’s alternative UI.
//
// # Return Value
// 
// [true] if the action was successfully triggered; otherwise, [false]. This
// method does not indicate the success or failure of the action, just the
// fact that the action was successfully triggered.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// Use this method to trigger changes to the UI due to a mouse-hover or
// similar event.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityPerformShowAlternateUI()

func (o NSView) AccessibilityPerformShowAlternateUI() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("accessibilityPerformShowAlternateUI"))
	return rv
	}

// Returns to the accessibility element’s original UI.
//
// # Return Value
// 
// [true] if the action was successfully triggered; otherwise, [false]. This
// method does not indicate the success or failure of the action, just the
// fact that the action was successfully triggered.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// Call this method after successfully calling
// [AccessibilityPerformShowAlternateUI] to return to the original UI.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityPerformShowDefaultUI()

func (o NSView) AccessibilityPerformShowDefaultUI() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("accessibilityPerformShowDefaultUI"))
	return rv
	}

// Displays the menu accessibility element.
//
// # Return Value
// 
// [true] if the action was successfully triggered; otherwise, [false]. This
// method does not indicate the success or failure of the action, just the
// fact that the action was successfully triggered.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// Use this method to display the contextual menu for the element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityPerformShowMenu()

func (o NSView) AccessibilityPerformShowMenu() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("accessibilityPerformShowMenu"))
	return rv
	}

// Brings the window to the front.
//
// # Return Value
// 
// [true] if the action was successfully triggered; otherwise, [false]. This
// method does not indicate the success or failure of the action, just the
// fact that the action was successfully triggered.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// The window behaves as if you had clicked on the window’s title bar.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityPerformRaise()

func (o NSView) AccessibilityPerformRaise() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("accessibilityPerformRaise"))
	return rv
	}

// Returns the increment button for the stepper accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityincrementbutton()

func (o NSView) AccessibilityIncrementButton() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityIncrementButton"))
	return objectivec.Object{ID: rv}
	}

// Sets the increment button for the stepper accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityincrementbutton(_:)

func (o NSView) SetAccessibilityIncrementButton(accessibilityIncrementButton objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityIncrementButton:"), accessibilityIncrementButton)
	}

// Returns the decrement button for the stepper accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitydecrementbutton()

func (o NSView) AccessibilityDecrementButton() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityDecrementButton"))
	return objectivec.Object{ID: rv}
	}

// Sets the decrement button for the stepper accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitydecrementbutton(_:)

func (o NSView) SetAccessibilityDecrementButton(accessibilityDecrementButton objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityDecrementButton:"), accessibilityDecrementButton)
	}

// Increments the accessibility element’s value.
//
// # Return Value
// 
// [true] if the action was successfully triggered; otherwise, [false]. This
// method does not indicate the success or failure of the action, just the
// fact that the action was successfully triggered.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// Use this method on elements that have an adjustable [accessibilityValue]
// property.
//
// [accessibilityValue]: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityValue
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityPerformIncrement()

func (o NSView) AccessibilityPerformIncrement() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("accessibilityPerformIncrement"))
	return rv
	}

// Decrements the accessibility element’s value.
//
// # Return Value
// 
// [true] if the action was successfully triggered; otherwise, [false]. This
// method does not indicate the success or failure of the action, just the
// fact that the action was successfully triggered.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// Use this method on elements that have an adjustable [accessibilityValue]
// property.
//
// [accessibilityValue]: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityValue
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityPerformDecrement()

func (o NSView) AccessibilityPerformDecrement() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("accessibilityPerformDecrement"))
	return rv
	}

// Deletes the accessibility element’s value.
//
// # Return Value
// 
// [true] if the action was successfully triggered; otherwise, [false]. This
// method does not indicate the success or failure of the action, just the
// fact that the action was successfully triggered.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// Use this method on elements with values.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityPerformDelete()

func (o NSView) AccessibilityPerformDelete() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("accessibilityPerformDelete"))
	return rv
	}

// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityattributeduserinputlabels()

func (o NSView) AccessibilityAttributedUserInputLabels() foundation.NSAttributedString {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityAttributedUserInputLabels"))
	return foundation.NSAttributedStringFromID(rv)
	}

// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityuserinputlabels()

func (o NSView) AccessibilityUserInputLabels() string {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityUserInputLabels"))
	return foundation.NSStringFromID(rv).String()
	}

//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityattributeduserinputlabels(_:)

func (o NSView) SetAccessibilityAttributedUserInputLabels(accessibilityAttributedUserInputLabels foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityAttributedUserInputLabels:"), accessibilityAttributedUserInputLabels)
	}

//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityuserinputlabels(_:)

func (o NSView) SetAccessibilityUserInputLabels(accessibilityUserInputLabels foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityUserInputLabels:"), accessibilityUserInputLabels)
	}





			// Protocol methods for NSAppearanceCustomization
			





			// Protocol methods for NSDraggingDestination
			








			// Protocol methods for NSUserInterfaceItemIdentification
			









// ShowDefinitionForAttributedStringRangeOptionsBaselineOriginProviderSync is a synchronous wrapper around [NSView.ShowDefinitionForAttributedStringRangeOptionsBaselineOriginProvider].
// It blocks until the completion handler fires or the context is cancelled.
func (v NSView) ShowDefinitionForAttributedStringRangeOptionsBaselineOriginProviderSync(ctx context.Context, attrString foundation.NSAttributedString, targetRange foundation.NSRange, options foundation.INSDictionary) (foundation.NSRange, error) {
	done := make(chan foundation.NSRange, 1)
	v.ShowDefinitionForAttributedStringRangeOptionsBaselineOriginProvider(attrString, targetRange, options, func(val foundation.NSRange) {
		done <- val
	})
	select {
	case r := <-done:
		return r, nil
	case <-ctx.Done():
		return foundation.NSRange{}, ctx.Err()
	}
}






