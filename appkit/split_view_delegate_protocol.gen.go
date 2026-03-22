// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"fmt"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)
var _ = fmt.Sprintf

// A set of optional methods that a delegate of a split view implements.
//
// See: https://developer.apple.com/documentation/AppKit/NSSplitViewDelegate
type NSSplitViewDelegate interface {
	objectivec.IObject
}

// NSSplitViewDelegateObject wraps an existing Objective-C object that conforms to the NSSplitViewDelegate protocol.
type NSSplitViewDelegateObject struct {
	objectivec.Object
}
func (o NSSplitViewDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSSplitViewDelegateObjectFromID constructs a [NSSplitViewDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSSplitViewDelegateObjectFromID(id objc.ID) NSSplitViewDelegateObject {
	return NSSplitViewDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Notifies the delegate when the split view is about to resize its subviews.
//
// notification: A notification named [willResizeSubviewsNotification], which posts before a
// change to the size of some or all subviews of a split view.
// //
// [willResizeSubviewsNotification]: https://developer.apple.com/documentation/AppKit/NSSplitView/willResizeSubviewsNotification
//
// # Discussion
// 
// If the delegate implements this method, the system automatically registers
// it to receive this notification.
// 
// The default notification center invokes this method before the split view
// resizes two of its subviews in response to the repositioning of a divider.
//
// See: https://developer.apple.com/documentation/AppKit/NSSplitViewDelegate/splitViewWillResizeSubviews(_:)
func (o NSSplitViewDelegateObject) SplitViewWillResizeSubviews(notification foundation.NSNotification) {
	objc.Send[struct{}](o.ID, objc.Sel("splitViewWillResizeSubviews:"), notification)
	}
// Notifies the delegate when the split view resizes its subviews.
//
// notification: A notification named [didResizeSubviewsNotification], which posts after a
// change to the size of some or all subviews of a split view.
// //
// [didResizeSubviewsNotification]: https://developer.apple.com/documentation/AppKit/NSSplitView/didResizeSubviewsNotification
//
// # Discussion
// 
// If the delegate implements this method, the system automatically registers
// it to receive this notification.
// 
// The default notification center invokes this method after the split view
// resizes two of its subviews in response to the repositioning of a divider.
//
// See: https://developer.apple.com/documentation/AppKit/NSSplitViewDelegate/splitViewDidResizeSubviews(_:)
func (o NSSplitViewDelegateObject) SplitViewDidResizeSubviews(notification foundation.NSNotification) {
	objc.Send[struct{}](o.ID, objc.Sel("splitViewDidResizeSubviews:"), notification)
	}
// Allows the delegate to determine whether the user can collapse and expand
// the specified subview.
//
// splitView: The split view that sends the message.
//
// subview: The subview to collapse.
//
// # Return Value
// 
// [true] if `subview` collapses when the user drags a divider beyond the
// halfway mark between its minimum size and its edge; otherwise, [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// The `subview` expands when the user drags the divider beyond the halfway
// mark between its minimum size and its edge.
// 
// To specify the minimum size, define the methods
// [SplitViewConstrainMaxCoordinateOfSubviewAt] and
// [SplitViewConstrainMinCoordinateOfSubviewAt]. A subview can collapse only
// if you also define [SplitViewConstrainMinCoordinateOfSubviewAt].
// 
// A collapsed subview isn’t visible, but the split view object retains it
// with the same size as before the collapse.
// 
// If the delegate doesn’t implement this method, the subviews can’t
// collapse.
//
// See: https://developer.apple.com/documentation/AppKit/NSSplitViewDelegate/splitView(_:canCollapseSubview:)
func (o NSSplitViewDelegateObject) SplitViewCanCollapseSubview(splitView INSSplitView, subview INSView) bool {
	rv := objc.Send[bool](o.ID, objc.Sel("splitView:canCollapseSubview:"), splitView, subview)
	return rv
	}
// Allows the delegate to modify the rectangle where mouse clicks initiate
// divider dragging.
//
// splitView: The split view that sends the message.
//
// proposedEffectiveRect: The proposed rectangle where mouse clicks initiate divider dragging. The
// rectangle uses the coordinate system that `splitView` defines.
//
// drawnRect: The frame of the divider in the coordinate system that `splitView` defines.
//
// dividerIndex: The index of the divider.
//
// # Return Value
// 
// A rectangle, in the coordinate system that `splitView` defines, where mouse
// clicks initiate divider dragging.
//
// # Discussion
// 
// A split view with thick dividers proposes the drawn frame as the effective
// frame. A split view with thin dividers proposes an effective frame that’s
// a little larger than the drawn frame to make it easier for the user to grab
// the divider.
// 
// If a split view has no delegate, or if its delegate doesn’t respond to
// this message, the split view behaves as if it has a delegate that returns
// `proposedEffectiveRect` when it receives this message.
//
// See: https://developer.apple.com/documentation/AppKit/NSSplitViewDelegate/splitView(_:effectiveRect:forDrawnRect:ofDividerAt:)
func (o NSSplitViewDelegateObject) SplitViewEffectiveRectForDrawnRectOfDividerAtIndex(splitView INSSplitView, proposedEffectiveRect corefoundation.CGRect, drawnRect corefoundation.CGRect, dividerIndex int) corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](o.ID, objc.Sel("splitView:effectiveRect:forDrawnRect:ofDividerAtIndex:"), splitView, proposedEffectiveRect, drawnRect, dividerIndex)
	return rv
	}
// Allows the delegate to determine whether the user can drag a divider or
// adjust it off the edge of the split view.
//
// splitView: The split view that sends the message.
//
// dividerIndex: The zero-based index of the divider.
//
// # Return Value
// 
// [true] if the user can drag the divider off the edge of the split view,
// resulting in it not being visible.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// If a split view has no delegate, or if its delegate doesn’t respond to
// this message, the split view behaves as if it has a delegate that returns
// [false] when it receives this message.
//
// [false]: https://developer.apple.com/documentation/Swift/false
//
// See: https://developer.apple.com/documentation/AppKit/NSSplitViewDelegate/splitView(_:shouldHideDividerAt:)
func (o NSSplitViewDelegateObject) SplitViewShouldHideDividerAtIndex(splitView INSSplitView, dividerIndex int) bool {
	rv := objc.Send[bool](o.ID, objc.Sel("splitView:shouldHideDividerAtIndex:"), splitView, dividerIndex)
	return rv
	}
// Allows the delegate to return an additional rectangle where mouse clicks
// can initiate divider dragging.
//
// splitView: The split view that sends the message.
//
// dividerIndex: The index of the divider.
//
// # Return Value
// 
// An additional rectangle, in the coordinate system that `splitView` defines,
// where mouse clicks can initiate divider dragging. Returning [NSZeroRect]
// indicates no additional dragging rectangle is necessary.
//
// [NSZeroRect]: https://developer.apple.com/documentation/Foundation/NSZeroRect
//
// # Discussion
// 
// If a split view has no delegate, or if its delegate doesn’t respond to
// this message, only mouse clicks within the effective frame of a divider
// initiate divider dragging.
//
// See: https://developer.apple.com/documentation/AppKit/NSSplitViewDelegate/splitView(_:additionalEffectiveRectOfDividerAt:)
func (o NSSplitViewDelegateObject) SplitViewAdditionalEffectiveRectOfDividerAtIndex(splitView INSSplitView, dividerIndex int) corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](o.ID, objc.Sel("splitView:additionalEffectiveRectOfDividerAtIndex:"), splitView, dividerIndex)
	return rv
	}
// Allows the delegate to constrain the divider to certain positions.
//
// splitView: The split view that sends the message.
//
// proposedPosition: The cursor’s current position, and the proposed position of the divider.
//
// dividerIndex: The index of the divider the user is moving, with the first divider being
// `0` and increasing from top to bottom (or left to right).
//
// # Return Value
// 
// The position for constraining the divider.
//
// # Discussion
// 
// If the delegate implements this method, the split view calls it repeatedly
// as the user moves the divider.
// 
// If a subview’s height must be a multiple of a certain number, use this
// method to return the multiple nearest to `proposedPosition`.
//
// See: https://developer.apple.com/documentation/AppKit/NSSplitViewDelegate/splitView(_:constrainSplitPosition:ofSubviewAt:)
func (o NSSplitViewDelegateObject) SplitViewConstrainSplitPositionOfSubviewAt(splitView INSSplitView, proposedPosition float64, dividerIndex int) float64 {
	rv := objc.Send[float64](o.ID, objc.Sel("splitView:constrainSplitPosition:ofSubviewAt:"), splitView, proposedPosition, dividerIndex)
	return rv
	}
// Allows the delegate to constrain the minimum coordinate limit of a divider
// when the user drags it.
//
// splitView: The split view that sends the message.
//
// proposedMinimumPosition: The proposed minimum coordinate limit of the subview in the split view’s
// flipped coordinate system.
//
// dividerIndex: Specifies the divider the user is moving, with the first divider being 0
// and increasing from top to bottom (or left to right).
//
// # Return Value
// 
// The minimum coordinate limit of the divider.
//
// # Discussion
// 
// The delegate receives this message before the split view begins tracking
// the cursor to position a divider. You can further constrain the limits, but
// you can’t extend the divider limits.
// 
// If the split bars are horizontal (views are one on top of the other),
// `proposedMin` is the top limit. If the split bars are vertical (views are
// side by side), `proposedMin` is the left limit. The initial value of
// `proposedMin` is the top (or left side) of the subview before the divider.
//
// See: https://developer.apple.com/documentation/AppKit/NSSplitViewDelegate/splitView(_:constrainMinCoordinate:ofSubviewAt:)
func (o NSSplitViewDelegateObject) SplitViewConstrainMinCoordinateOfSubviewAt(splitView INSSplitView, proposedMinimumPosition float64, dividerIndex int) float64 {
	rv := objc.Send[float64](o.ID, objc.Sel("splitView:constrainMinCoordinate:ofSubviewAt:"), splitView, proposedMinimumPosition, dividerIndex)
	return rv
	}
// Allows the delegate to constrain the maximum coordinate limit of a divider
// when the user drags it.
//
// splitView: The split view that sends the message.
//
// proposedMaximumPosition: The proposed maximum coordinate limit of the subview in the split view’s
// flipped coordinate system.
//
// dividerIndex: Specifies the divider the user is moving, with the first divider being 0
// and increasing from top to bottom (or left to right).
//
// # Return Value
// 
// The maximum coordinate limit of the divider.
//
// # Discussion
// 
// The delegate receives this message before the split view begins tracking
// the mouse to position a divider. You can further constrain the limits, but
// you can’t extend the divider limits.
// 
// If the split bars are horizontal (views are one on top of the other),
// `proposedMax` is the bottom limit. If the split bars are vertical (views
// are side by side), `proposedMax` is the right limit. The initial value of
// `proposedMax` is the bottom (or right side) of the subview after the
// divider.
//
// See: https://developer.apple.com/documentation/AppKit/NSSplitViewDelegate/splitView(_:constrainMaxCoordinate:ofSubviewAt:)
func (o NSSplitViewDelegateObject) SplitViewConstrainMaxCoordinateOfSubviewAt(splitView INSSplitView, proposedMaximumPosition float64, dividerIndex int) float64 {
	rv := objc.Send[float64](o.ID, objc.Sel("splitView:constrainMaxCoordinate:ofSubviewAt:"), splitView, proposedMaximumPosition, dividerIndex)
	return rv
	}
// Allows the delegate to specify custom sizing behavior for the subviews of
// the split view.
//
// splitView: The split view that sends the message.
//
// oldSize: The size of the split view before the user resizes it.
//
// # Discussion
// 
// If the delegate implements this method, it receives this message after the
// split view resizes.
// 
// Resize the subviews so that the sum of the sizes of the subviews plus the
// sum of the thickness of the dividers equals the size of the new frame of
// the [NSSplitView]. You can get the thickness of a divider through the
// [DividerThickness] method.
// 
// If you implement this delegate method to resize subviews on your own, the
// [NSSplitView] doesn’t perform any error checking for you. However, you
// can invoke [AdjustSubviews] to perform the default sizing behavior.
//
// See: https://developer.apple.com/documentation/AppKit/NSSplitViewDelegate/splitView(_:resizeSubviewsWithOldSize:)
func (o NSSplitViewDelegateObject) SplitViewResizeSubviewsWithOldSize(splitView INSSplitView, oldSize corefoundation.CGSize) {
	objc.Send[struct{}](o.ID, objc.Sel("splitView:resizeSubviewsWithOldSize:"), splitView, oldSize)
	}
// Allows the delegate to specify whether to resize the subview.
//
// splitView: The split view that sends the message.
//
// view: The subview to resize.
//
// # Return Value
// 
// If [AdjustSubviews] can change the size of the subview, [true]; otherwise,
// [false]. By returning [false], you lock the size of the split view
// `subview` while the split view resizes.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// Regardless of the value that this method returns, [AdjustSubviews] may
// change the origin of the subview. Nonresizable subviews may resize to
// prevent an invalid subview layout.
// 
// If a split view has no delegate, or if its delegate doesn’t respond to
// this message, the split view behaves as if this method returns [true].
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSSplitViewDelegate/splitView(_:shouldAdjustSizeOfSubview:)
func (o NSSplitViewDelegateObject) SplitViewShouldAdjustSizeOfSubview(splitView INSSplitView, view INSView) bool {
	rv := objc.Send[bool](o.ID, objc.Sel("splitView:shouldAdjustSizeOfSubview:"), splitView, view)
	return rv
	}

// NSSplitViewDelegateConfig holds optional typed callbacks for [NSSplitViewDelegate] methods.
// Set non-nil fields to register the corresponding Objective-C delegate method.
// Methods with nil callbacks are not registered, so [NSObject.RespondsToSelector]
// returns false for them — matching the Objective-C delegate pattern exactly.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/appkit/nssplitviewdelegate
type NSSplitViewDelegateConfig struct {

	// Managing Subviews
	// WillResizeSubviews — Notifies the delegate when the split view is about to resize its subviews.
	WillResizeSubviews func(notification foundation.NSNotification)
	// DidResizeSubviews — Notifies the delegate when the split view resizes its subviews.
	DidResizeSubviews func(notification foundation.NSNotification)
	// CanCollapseSubview — Allows the delegate to determine whether the user can collapse and expand the specified subview.
	CanCollapseSubview func(splitView NSSplitView, subview NSView) bool

	// Constraining Split Position
	// ConstrainSplitPositionOfSubviewAt — Allows the delegate to constrain the divider to certain positions.
	ConstrainSplitPositionOfSubviewAt func(splitView NSSplitView, proposedPosition float64, dividerIndex int) float64

	// Adjusting Subviews Manually
	// ConstrainMinCoordinateOfSubviewAt — Allows the delegate to constrain the minimum coordinate limit of a divider when the user drags it.
	ConstrainMinCoordinateOfSubviewAt func(splitView NSSplitView, proposedMinimumPosition float64, dividerIndex int) float64
	// ConstrainMaxCoordinateOfSubviewAt — Allows the delegate to constrain the maximum coordinate limit of a divider when the user drags it.
	ConstrainMaxCoordinateOfSubviewAt func(splitView NSSplitView, proposedMaximumPosition float64, dividerIndex int) float64
	// ShouldAdjustSizeOfSubview — Allows the delegate to specify whether to resize the subview.
	ShouldAdjustSizeOfSubview func(splitView NSSplitView, view NSView) bool

	// Other Methods
	// ShouldHideDividerAtIndex — Allows the delegate to determine whether the user can drag a divider or adjust it off the edge of the split view.
	ShouldHideDividerAtIndex func(splitView NSSplitView, dividerIndex int) bool
}

// NewNSSplitViewDelegate creates an Objective-C object implementing the [NSSplitViewDelegate] protocol.
//
// Each call registers a unique Objective-C class containing only the methods
// set in config. This means [NSObject.RespondsToSelector] works correctly
// for optional delegate methods — only non-nil callbacks are registered.
//
// The returned [NSSplitViewDelegateObject] satisfies the [NSSplitViewDelegate] interface
// and can be passed directly to SetDelegate and similar methods.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/appkit/nssplitviewdelegate
func NewNSSplitViewDelegate(config NSSplitViewDelegateConfig) NSSplitViewDelegateObject {
	n := delegateClassCounter.Add(1)
	className := fmt.Sprintf("GoNSSplitViewDelegate_%d", n)

	var methods []objc.MethodDef

	if config.WillResizeSubviews != nil {
		fn := config.WillResizeSubviews
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("splitViewWillResizeSubviews:"),
			Fn: func(self objc.ID, _cmd objc.SEL, notificationID objc.ID) {
				notification := foundation.NSNotificationFromID(notificationID)
				fn(notification)
			},
		})
	}

	if config.DidResizeSubviews != nil {
		fn := config.DidResizeSubviews
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("splitViewDidResizeSubviews:"),
			Fn: func(self objc.ID, _cmd objc.SEL, notificationID objc.ID) {
				notification := foundation.NSNotificationFromID(notificationID)
				fn(notification)
			},
		})
	}

	if config.CanCollapseSubview != nil {
		fn := config.CanCollapseSubview
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("splitView:canCollapseSubview:"),
			Fn: func(self objc.ID, _cmd objc.SEL, splitViewID objc.ID, subviewID objc.ID) bool {
				splitView := NSSplitViewFromID(splitViewID)
				subview := NSViewFromID(subviewID)
				return fn(splitView, subview)
			},
		})
	}

	if config.ShouldHideDividerAtIndex != nil {
		fn := config.ShouldHideDividerAtIndex
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("splitView:shouldHideDividerAtIndex:"),
			Fn: func(self objc.ID, _cmd objc.SEL, splitViewID objc.ID, dividerIndex int) bool {
				splitView := NSSplitViewFromID(splitViewID)
				return fn(splitView, dividerIndex)
			},
		})
	}

	if config.ConstrainSplitPositionOfSubviewAt != nil {
		fn := config.ConstrainSplitPositionOfSubviewAt
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("splitView:constrainSplitPosition:ofSubviewAt:"),
			Fn: func(self objc.ID, _cmd objc.SEL, splitViewID objc.ID, proposedPosition float64, dividerIndex int) float64 {
				splitView := NSSplitViewFromID(splitViewID)
				return fn(splitView, proposedPosition, dividerIndex)
			},
		})
	}

	if config.ConstrainMinCoordinateOfSubviewAt != nil {
		fn := config.ConstrainMinCoordinateOfSubviewAt
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("splitView:constrainMinCoordinate:ofSubviewAt:"),
			Fn: func(self objc.ID, _cmd objc.SEL, splitViewID objc.ID, proposedMinimumPosition float64, dividerIndex int) float64 {
				splitView := NSSplitViewFromID(splitViewID)
				return fn(splitView, proposedMinimumPosition, dividerIndex)
			},
		})
	}

	if config.ConstrainMaxCoordinateOfSubviewAt != nil {
		fn := config.ConstrainMaxCoordinateOfSubviewAt
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("splitView:constrainMaxCoordinate:ofSubviewAt:"),
			Fn: func(self objc.ID, _cmd objc.SEL, splitViewID objc.ID, proposedMaximumPosition float64, dividerIndex int) float64 {
				splitView := NSSplitViewFromID(splitViewID)
				return fn(splitView, proposedMaximumPosition, dividerIndex)
			},
		})
	}

	if config.ShouldAdjustSizeOfSubview != nil {
		fn := config.ShouldAdjustSizeOfSubview
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("splitView:shouldAdjustSizeOfSubview:"),
			Fn: func(self objc.ID, _cmd objc.SEL, splitViewID objc.ID, viewID objc.ID) bool {
				splitView := NSSplitViewFromID(splitViewID)
				view := NSViewFromID(viewID)
				return fn(splitView, view)
			},
		})
	}

	nsObjectClass := objc.GetClass("NSObject")
	proto := objc.GetProtocol("NSSplitViewDelegate")

	var protocols []*objc.Protocol
	if proto != nil {
		protocols = append(protocols, proto)
	}

	cls, err := objc.RegisterClass(className, nsObjectClass, protocols, nil, methods)
	if err != nil {
		panic(fmt.Sprintf("NewNSSplitViewDelegate: RegisterClass %s: %v", className, err))
	}

	instance := objc.ID(cls).Send(objc.RegisterName("alloc")).Send(objc.RegisterName("init"))
	return NSSplitViewDelegateObjectFromID(instance)
}

