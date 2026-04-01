// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSCursor] class.
var (
	_NSCursorClass     NSCursorClass
	_NSCursorClassOnce sync.Once
)

func getNSCursorClass() NSCursorClass {
	_NSCursorClassOnce.Do(func() {
		_NSCursorClass = NSCursorClass{class: objc.GetClass("NSCursor")}
	})
	return _NSCursorClass
}

// GetNSCursorClass returns the class object for NSCursor.
func GetNSCursorClass() NSCursorClass {
	return getNSCursorClass()
}

type NSCursorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSCursorClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSCursorClass) Alloc() NSCursor {
	rv := objc.Send[NSCursor](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A pointer (also called a cursor).
//
// # Overview
//
// The following table shows and describes the system cursors, and indicates
// the class method for obtaining them:
//
// [Table data omitted]
//
// In macOS 10.3 and later, cursor size is no longer limited to 16 by 16
// pixels.
//
// # Cursor rectangles
//
// In Cocoa, you can change the currently displayed cursor based on the
// position of the mouse over one of your views. You might use this technique
// to provide visual feedback about what actions the user can take with the
// mouse. For example, you might display one of the resize cursors whenever
// the mouse moves over a portion of your view that acts as a custom resizing
// handle. To set this up, you associate a cursor object with one or more
// cursor rectangles in the view.
//
// Cursor rectangles are a specialized type of tracking rectangles, which are
// used to monitor the mouse location in a view. Views implement cursor
// rectangles using tracking rectangles but provide methods for setting and
// refreshing cursor rectangles that are distinct from the generic tracking
// rectangle interface. For information on how to set up cursor rectangles,
// see [Mouse-Tracking and Cursor-Update Events].
//
// # Balancing cursor hiding and unhiding
//
// Each call to [NSCursor.Hide] cursor must have a corresponding [NSCursor.Unhide] call. For
// example,
//
// Will result in the cursor still being hidden because the `hide` and
// `unhide` method invocations are not balanced. Instead you must balance the
// method calls, such as in the following example:
//
// There are corresponding cursor `hide` and `unhide` calls, thus the cursor
// will become visible.
//
// # Initializing a new cursor
//
//   - [NSCursor.InitWithImageHotSpot]: Initializes a cursor with the given image and hot spot.
//   - [NSCursor.InitWithCoder]
//
// # Setting cursor attributes
//
//   - [NSCursor.Image]: The cursor’s image.
//   - [NSCursor.HotSpot]: The position of the click location within the cursor.
//
// # Controlling which cursor is current
//
//   - [NSCursor.Push]: Puts the receiver on top of the cursor stack and makes it the current cursor.
//   - [NSCursor.Set]: Makes the receiver the current cursor.
//
// See: https://developer.apple.com/documentation/AppKit/NSCursor
//
// [Mouse-Tracking and Cursor-Update Events]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/EventOverview/MouseTrackingEvents/MouseTrackingEvents.html#//apple_ref/doc/uid/10000060i-CH11
type NSCursor struct {
	objectivec.Object
}

// NSCursorFromID constructs a [NSCursor] from an objc.ID.
//
// A pointer (also called a cursor).
func NSCursorFromID(id objc.ID) NSCursor {
	return NSCursor{objectivec.Object{ID: id}}
}

// NOTE: NSCursor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSCursor] class.
//
// # Initializing a new cursor
//
//   - [INSCursor.InitWithImageHotSpot]: Initializes a cursor with the given image and hot spot.
//   - [INSCursor.InitWithCoder]
//
// # Setting cursor attributes
//
//   - [INSCursor.Image]: The cursor’s image.
//   - [INSCursor.HotSpot]: The position of the click location within the cursor.
//
// # Controlling which cursor is current
//
//   - [INSCursor.Push]: Puts the receiver on top of the cursor stack and makes it the current cursor.
//   - [INSCursor.Set]: Makes the receiver the current cursor.
//
// See: https://developer.apple.com/documentation/AppKit/NSCursor
type INSCursor interface {
	objectivec.IObject

	// Topic: Initializing a new cursor

	// Initializes a cursor with the given image and hot spot.
	InitWithImageHotSpot(newImage INSImage, point corefoundation.CGPoint) NSCursor
	InitWithCoder(coder foundation.INSCoder) NSCursor

	// Topic: Setting cursor attributes

	// The cursor’s image.
	Image() INSImage
	// The position of the click location within the cursor.
	HotSpot() corefoundation.CGPoint

	// Topic: Controlling which cursor is current

	// Puts the receiver on top of the cursor stack and makes it the current cursor.
	Push()
	// Makes the receiver the current cursor.
	Set()

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (c NSCursor) Init() NSCursor {
	rv := objc.Send[NSCursor](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c NSCursor) Autorelease() NSCursor {
	rv := objc.Send[NSCursor](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSCursor creates a new NSCursor instance.
func NewNSCursor() NSCursor {
	class := getNSCursorClass()
	rv := objc.Send[NSCursor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/AppKit/NSCursor/init(coder:)
func NewCursorWithCoder(coder foundation.INSCoder) NSCursor {
	instance := getNSCursorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSCursorFromID(rv)
}

// Initializes a cursor with the given image and hot spot.
//
// newImage: The image to assign to the cursor.
//
// point: The point to set as the cursor’s hot spot.
//
// # Return Value
//
// An initialized cursor object.
//
// # Discussion
//
// This method is the designated initializer for the class.
//
// See: https://developer.apple.com/documentation/AppKit/NSCursor/init(image:hotSpot:)
func NewCursorWithImageHotSpot(newImage INSImage, point corefoundation.CGPoint) NSCursor {
	instance := getNSCursorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithImage:hotSpot:"), newImage, point)
	return NSCursorFromID(rv)
}

// Initializes a cursor with the given image and hot spot.
//
// newImage: The image to assign to the cursor.
//
// point: The point to set as the cursor’s hot spot.
//
// # Return Value
//
// An initialized cursor object.
//
// # Discussion
//
// This method is the designated initializer for the class.
//
// See: https://developer.apple.com/documentation/AppKit/NSCursor/init(image:hotSpot:)
func (c NSCursor) InitWithImageHotSpot(newImage INSImage, point corefoundation.CGPoint) NSCursor {
	rv := objc.Send[NSCursor](c.ID, objc.Sel("initWithImage:hotSpot:"), newImage, point)
	return rv
}

// See: https://developer.apple.com/documentation/AppKit/NSCursor/init(coder:)
func (c NSCursor) InitWithCoder(coder foundation.INSCoder) NSCursor {
	rv := objc.Send[NSCursor](c.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}

// Puts the receiver on top of the cursor stack and makes it the current
// cursor.
//
// See: https://developer.apple.com/documentation/AppKit/NSCursor/push()
func (c NSCursor) Push() {
	objc.Send[objc.ID](c.ID, objc.Sel("push"))
}

// Makes the receiver the current cursor.
//
// See: https://developer.apple.com/documentation/AppKit/NSCursor/set()
func (c NSCursor) Set() {
	objc.Send[objc.ID](c.ID, objc.Sel("set"))
}
func (c NSCursor) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](c.ID, objc.Sel("encodeWithCoder:"), coder)
}

// Makes the current cursor invisible.
//
// # Discussion
//
// If another cursor becomes current, that cursor will be invisible, too. It
// will remain invisible until you invoke the [Unhide] method.
//
// Each invocation of `hide` must be balanced by an invocation of [Unhide] in
// order for the cursor to be displayed.
//
// The [Hide] method overrides [SetHiddenUntilMouseMoves].
//
// See: https://developer.apple.com/documentation/AppKit/NSCursor/hide()
func (_NSCursorClass NSCursorClass) Hide() {
	objc.Send[objc.ID](objc.ID(_NSCursorClass.class), objc.Sel("hide"))
}

// Negates an earlier call to [Hide] by showing the current cursor.
//
// # Discussion
//
// Each invocation of `unhide` must be balanced by an invocation of [Hide] in
// order for the cursor display to be correct.
//
// See: https://developer.apple.com/documentation/AppKit/NSCursor/unhide()
func (_NSCursorClass NSCursorClass) Unhide() {
	objc.Send[objc.ID](objc.ID(_NSCursorClass.class), objc.Sel("unhide"))
}

// Sets whether the cursor is hidden until the mouse moves.
//
// flag: true to hide the cursor until one of the following occurs:
//
// - The mouse moves.
// - You invoke the method again, with `flag` set to false.
//
// # Discussion
//
// Do not try to counter this method by invoking [Unhide]. The results are
// undefined.
//
// See: https://developer.apple.com/documentation/AppKit/NSCursor/setHiddenUntilMouseMoves(_:)
func (_NSCursorClass NSCursorClass) SetHiddenUntilMouseMoves(flag bool) {
	objc.Send[objc.ID](objc.ID(_NSCursorClass.class), objc.Sel("setHiddenUntilMouseMoves:"), flag)
}

// Pops the current cursor off the top of the stack.
//
// # Discussion
//
// The new object on the top of the stack becomes the current cursor. If the
// current cursor is the only cursor on the stack, this method does nothing.
//
// See: https://developer.apple.com/documentation/AppKit/NSCursor/pop()-swift.type.method
func (_NSCursorClass NSCursorClass) Pop() {
	objc.Send[objc.ID](objc.ID(_NSCursorClass.class), objc.Sel("pop"))
}

// Returns the cursor for resizing a column (vertical divider) in the
// specified directions.
//
// directions: The direction in which a column can be resized.
//
// See: https://developer.apple.com/documentation/AppKit/NSCursor/columnResizeCursorInDirections:
func (_NSCursorClass NSCursorClass) ColumnResizeCursorInDirections(directions NSHorizontalDirections) NSCursor {
	rv := objc.Send[objc.ID](objc.ID(_NSCursorClass.class), objc.Sel("columnResizeCursorInDirections:"), directions)
	return NSCursorFromID(rv)
}

// Returns the cursor for resizing a rectangular frame from the specified edge
// or corner.
//
// position: The position along the perimeter of a rectangular frame (its edges and
// corners) from which it’s resized.
//
// directions: The directions in which a rectangular frame can be resized.
//
// See: https://developer.apple.com/documentation/AppKit/NSCursor/frameResizeCursorFromPosition:inDirections:
func (_NSCursorClass NSCursorClass) FrameResizeCursorFromPositionInDirections(position NSCursorFrameResizePosition, directions NSCursorFrameResizeDirections) NSCursor {
	rv := objc.Send[objc.ID](objc.ID(_NSCursorClass.class), objc.Sel("frameResizeCursorFromPosition:inDirections:"), position, directions)
	return NSCursorFromID(rv)
}

// Returns the cursor for resizing a row (horizontal divider) in the specified
// directions.
//
// directions: The direction in which a row can be resized.
//
// See: https://developer.apple.com/documentation/AppKit/NSCursor/rowResizeCursorInDirections:
func (_NSCursorClass NSCursorClass) RowResizeCursorInDirections(directions NSVerticalDirections) NSCursor {
	rv := objc.Send[objc.ID](objc.ID(_NSCursorClass.class), objc.Sel("rowResizeCursorInDirections:"), directions)
	return NSCursorFromID(rv)
}

// The cursor’s image.
//
// # Discussion
//
// The cursor image or `nil` if none exists. Note that an [NSCursor] object is
// immutable: you cannot change its image after it’s created. Instead, use
// [InitWithImageHotSpot] to create a new cursor with the new settings.
//
// See: https://developer.apple.com/documentation/AppKit/NSCursor/image
func (c NSCursor) Image() INSImage {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("image"))
	return NSImageFromID(objc.ID(rv))
}

// The position of the click location within the cursor.
//
// # Discussion
//
// The hot spot precisely determines the click location within the cursor’s
// image. Using its flipped coordinate system, you calculate the hot spot in
// points with the top-left corner acting as the origin. For example, the
// arrow cursor’s hot spot is at the intersection of its left and right
// edges, which is inset 4pts from the image’s corner to account for the
// arrow’s stroke and shadow.
//
// [media-4311497]
//
// Note that an [NSCursor] object is immutable: you can’t change its hot
// spot after it’s created. Instead, use [InitWithImageHotSpot] to create a
// new cursor with the new settings.
//
// See: https://developer.apple.com/documentation/AppKit/NSCursor/hotSpot
func (c NSCursor) HotSpot() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](c.ID, objc.Sel("hotSpot"))
	return corefoundation.CGPoint(rv)
}

// Returns the application’s current cursor.
//
// # Return Value
//
// The top cursor on the application’s cursor stack. This cursor may not be
// the visible cursor on the screen if a different application is currently
// active.
//
// # Discussion
//
// The method only returns the cursor set by your application using [NSCursor]
// methods. It does not return cursors set by other applications or cursors
// set by your application using Carbon APIs.
//
// See: https://developer.apple.com/documentation/AppKit/NSCursor/current
func (_NSCursorClass NSCursorClass) CurrentCursor() NSCursor {
	rv := objc.Send[objc.ID](objc.ID(_NSCursorClass.class), objc.Sel("currentCursor"))
	return NSCursorFromID(objc.ID(rv))
}

// Returns the current system cursor.
//
// # Return Value
//
// A cursor whose image and hot spot match those of the currently-displayed
// cursor on the system
//
// # Discussion
//
// This method returns the current system cursor regardless of which
// application set the cursor, and whether Cocoa or Carbon APIs were used to
// set it.
//
// This method replaces the now deprecated QDGetCursorData function.
//
// See: https://developer.apple.com/documentation/AppKit/NSCursor/currentSystem
func (_NSCursorClass NSCursorClass) CurrentSystemCursor() NSCursor {
	rv := objc.Send[objc.ID](objc.ID(_NSCursorClass.class), objc.Sel("currentSystemCursor"))
	return NSCursorFromID(objc.ID(rv))
}

// Returns the default cursor, the arrow cursor.
//
// # Return Value
//
// The default cursor, a slanted arrow with its hot spot at the tip. The arrow
// cursor is the one you’re used to seeing over buttons, scrollers, and many
// other objects in the window system.
//
// See: https://developer.apple.com/documentation/AppKit/NSCursor/arrow
func (_NSCursorClass NSCursorClass) ArrowCursor() NSCursor {
	rv := objc.Send[objc.ID](objc.ID(_NSCursorClass.class), objc.Sel("arrowCursor"))
	return NSCursorFromID(objc.ID(rv))
}

// Returns the contextual menu system cursor.
//
// # Return Value
//
// # The contextual menu cursor
//
// See: https://developer.apple.com/documentation/AppKit/NSCursor/contextualMenu
func (_NSCursorClass NSCursorClass) ContextualMenuCursor() NSCursor {
	rv := objc.Send[objc.ID](objc.ID(_NSCursorClass.class), objc.Sel("contextualMenuCursor"))
	return NSCursorFromID(objc.ID(rv))
}

// Returns the closed-hand system cursor.
//
// # Return Value
//
// The closed-hand cursor.
//
// See: https://developer.apple.com/documentation/AppKit/NSCursor/closedHand
func (_NSCursorClass NSCursorClass) ClosedHandCursor() NSCursor {
	rv := objc.Send[objc.ID](objc.ID(_NSCursorClass.class), objc.Sel("closedHandCursor"))
	return NSCursorFromID(objc.ID(rv))
}

// Returns the cross-hair system cursor.
//
// # Return Value
//
// The cross-hair cursor. This cursor is used for situations when precise
// location is required (where the lines cross is the hot spot).
//
// See: https://developer.apple.com/documentation/AppKit/NSCursor/crosshair
func (_NSCursorClass NSCursorClass) CrosshairCursor() NSCursor {
	rv := objc.Send[objc.ID](objc.ID(_NSCursorClass.class), objc.Sel("crosshairCursor"))
	return NSCursorFromID(objc.ID(rv))
}

// Returns a cursor indicating that the current operation will result in a
// disappearing item.
//
// # Return Value
//
// The system cursor that indicates that the current operation will result in
// a disappearing item (for example, when dragging an item from the dock or a
// toolbar).
//
// See: https://developer.apple.com/documentation/AppKit/NSCursor/disappearingItem
func (_NSCursorClass NSCursorClass) DisappearingItemCursor() NSCursor {
	rv := objc.Send[objc.ID](objc.ID(_NSCursorClass.class), objc.Sel("disappearingItemCursor"))
	return NSCursorFromID(objc.ID(rv))
}

// Returns a cursor indicating that the current operation will result in a
// copy action.
//
// # Return Value
//
// The drag copy cursor.
//
// See: https://developer.apple.com/documentation/AppKit/NSCursor/dragCopy
func (_NSCursorClass NSCursorClass) DragCopyCursor() NSCursor {
	rv := objc.Send[objc.ID](objc.ID(_NSCursorClass.class), objc.Sel("dragCopyCursor"))
	return NSCursorFromID(objc.ID(rv))
}

// Returns a cursor indicating that the current operation will result in a
// link action.
//
// # Return Value
//
// The drag link cursor.
//
// See: https://developer.apple.com/documentation/AppKit/NSCursor/dragLink
func (_NSCursorClass NSCursorClass) DragLinkCursor() NSCursor {
	rv := objc.Send[objc.ID](objc.ID(_NSCursorClass.class), objc.Sel("dragLinkCursor"))
	return NSCursorFromID(objc.ID(rv))
}

// Returns a cursor that looks like a capital I with a tiny crossbeam at its
// middle.
//
// # Return Value
//
// The I-beam cursor. This is the cursor that you’re used to seeing over
// editable or selectable text. The I-beam cursor’s default hot spot is
// where the crossbeam intersects the I.
//
// See: https://developer.apple.com/documentation/AppKit/NSCursor/iBeam
func (_NSCursorClass NSCursorClass) IBeamCursor() NSCursor {
	rv := objc.Send[objc.ID](objc.ID(_NSCursorClass.class), objc.Sel("IBeamCursor"))
	return NSCursorFromID(objc.ID(rv))
}

// Returns the cursor for editing vertical layout text.
//
// # Return Value
//
// The vertical layout text cursor. This cursor is used when editing vertical
// layout text.
//
// See: https://developer.apple.com/documentation/AppKit/NSCursor/iBeamCursorForVerticalLayout
func (_NSCursorClass NSCursorClass) IBeamCursorForVerticalLayout() NSCursor {
	rv := objc.Send[objc.ID](objc.ID(_NSCursorClass.class), objc.Sel("IBeamCursorForVerticalLayout"))
	return NSCursorFromID(objc.ID(rv))
}

// Returns the open-hand system cursor.
//
// # Return Value
//
// The open-hand cursor.
//
// See: https://developer.apple.com/documentation/AppKit/NSCursor/openHand
func (_NSCursorClass NSCursorClass) OpenHandCursor() NSCursor {
	rv := objc.Send[objc.ID](objc.ID(_NSCursorClass.class), objc.Sel("openHandCursor"))
	return NSCursorFromID(objc.ID(rv))
}

// Returns the operation not allowed cursor.
//
// # Return Value
//
// The operation not allowed cursor.
//
// # Discussion
//
// This cursor indicates that the operation that is being attempted, perhaps
// dragging to an item that can’t accept the drag type, is being denied.
//
// See: https://developer.apple.com/documentation/AppKit/NSCursor/operationNotAllowed
func (_NSCursorClass NSCursorClass) OperationNotAllowedCursor() NSCursor {
	rv := objc.Send[objc.ID](objc.ID(_NSCursorClass.class), objc.Sel("operationNotAllowedCursor"))
	return NSCursorFromID(objc.ID(rv))
}

// Returns the pointing-hand system cursor.
//
// # Return Value
//
// The pointing-hand cursor. The tip of the pointing finger is the hot spot.
//
// See: https://developer.apple.com/documentation/AppKit/NSCursor/pointingHand
func (_NSCursorClass NSCursorClass) PointingHandCursor() NSCursor {
	rv := objc.Send[objc.ID](objc.ID(_NSCursorClass.class), objc.Sel("pointingHandCursor"))
	return NSCursorFromID(objc.ID(rv))
}

// Returns the zoom-in cursor.
//
// # Discussion
//
// See: https://developer.apple.com/documentation/AppKit/NSCursor/zoomIn
func (_NSCursorClass NSCursorClass) ZoomInCursor() NSCursor {
	rv := objc.Send[objc.ID](objc.ID(_NSCursorClass.class), objc.Sel("zoomInCursor"))
	return NSCursorFromID(objc.ID(rv))
}

// Returns the zoom-out cursor.
//
// # Discussion
//
// See: https://developer.apple.com/documentation/AppKit/NSCursor/zoomOut
func (_NSCursorClass NSCursorClass) ZoomOutCursor() NSCursor {
	rv := objc.Send[objc.ID](objc.ID(_NSCursorClass.class), objc.Sel("zoomOutCursor"))
	return NSCursorFromID(objc.ID(rv))
}

// Returns the resize-down system cursor.
//
// # Return Value
//
// The resize-down cursor. This cursor is used when moving or resizing an
// object to indicate that the user can move only in the indicated direction.
//
// See: https://developer.apple.com/documentation/AppKit/NSCursor/resizeDown
func (_NSCursorClass NSCursorClass) ResizeDownCursor() NSCursor {
	rv := objc.Send[objc.ID](objc.ID(_NSCursorClass.class), objc.Sel("resizeDownCursor"))
	return NSCursorFromID(objc.ID(rv))
}

// Returns the resize-left system cursor.
//
// # Return Value
//
// The resize-left cursor. This cursor is used when moving or resizing an
// object to indicate that the user can move only in the indicated direction.
//
// See: https://developer.apple.com/documentation/AppKit/NSCursor/resizeLeft
func (_NSCursorClass NSCursorClass) ResizeLeftCursor() NSCursor {
	rv := objc.Send[objc.ID](objc.ID(_NSCursorClass.class), objc.Sel("resizeLeftCursor"))
	return NSCursorFromID(objc.ID(rv))
}

// Returns the resize-left-and-right system cursor.
//
// # Return Value
//
// The resize-left-and-right cursor. This cursor is used when moving or
// resizing an object and the object can be moved left or right.
//
// See: https://developer.apple.com/documentation/AppKit/NSCursor/resizeLeftRight
func (_NSCursorClass NSCursorClass) ResizeLeftRightCursor() NSCursor {
	rv := objc.Send[objc.ID](objc.ID(_NSCursorClass.class), objc.Sel("resizeLeftRightCursor"))
	return NSCursorFromID(objc.ID(rv))
}

// Returns the resize-right system cursor.
//
// # Return Value
//
// The resize-right cursor. This cursor is used when moving or resizing an
// object to indicate that the user can move only in the indicated direction.
//
// See: https://developer.apple.com/documentation/AppKit/NSCursor/resizeRight
func (_NSCursorClass NSCursorClass) ResizeRightCursor() NSCursor {
	rv := objc.Send[objc.ID](objc.ID(_NSCursorClass.class), objc.Sel("resizeRightCursor"))
	return NSCursorFromID(objc.ID(rv))
}

// Returns the resize-up system cursor.
//
// # Return Value
//
// The resize-up cursor. This cursor is used when moving or resizing an object
// to indicate that the user can move only in the indicated direction.
//
// See: https://developer.apple.com/documentation/AppKit/NSCursor/resizeUp
func (_NSCursorClass NSCursorClass) ResizeUpCursor() NSCursor {
	rv := objc.Send[objc.ID](objc.ID(_NSCursorClass.class), objc.Sel("resizeUpCursor"))
	return NSCursorFromID(objc.ID(rv))
}

// Returns the resize-up-and-down system cursor.
//
// # Return Value
//
// The resize-up-and-down cursor. This cursor is used when moving or resizing
// an object and the object can be moved up or down.
//
// See: https://developer.apple.com/documentation/AppKit/NSCursor/resizeUpDown
func (_NSCursorClass NSCursorClass) ResizeUpDownCursor() NSCursor {
	rv := objc.Send[objc.ID](objc.ID(_NSCursorClass.class), objc.Sel("resizeUpDownCursor"))
	return NSCursorFromID(objc.ID(rv))
}

// Returns the cursor for resizing a column (vertical divider) in either
// direction.
//
// See: https://developer.apple.com/documentation/AppKit/NSCursor/columnResize
func (_NSCursorClass NSCursorClass) ColumnResizeCursor() NSCursor {
	rv := objc.Send[objc.ID](objc.ID(_NSCursorClass.class), objc.Sel("columnResizeCursor"))
	return NSCursorFromID(objc.ID(rv))
}

// Returns the cursor for resizing a row (horizontal divider) in either
// direction.
//
// See: https://developer.apple.com/documentation/AppKit/NSCursor/rowResize
func (_NSCursorClass NSCursorClass) RowResizeCursor() NSCursor {
	rv := objc.Send[objc.ID](objc.ID(_NSCursorClass.class), objc.Sel("rowResizeCursor"))
	return NSCursorFromID(objc.ID(rv))
}
