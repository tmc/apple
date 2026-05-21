// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSTrackingArea] class.
var (
	_NSTrackingAreaClass     NSTrackingAreaClass
	_NSTrackingAreaClassOnce sync.Once
)

func getNSTrackingAreaClass() NSTrackingAreaClass {
	_NSTrackingAreaClassOnce.Do(func() {
		_NSTrackingAreaClass = NSTrackingAreaClass{class: objc.GetClass("NSTrackingArea")}
	})
	return _NSTrackingAreaClass
}

// GetNSTrackingAreaClass returns the class object for NSTrackingArea.
func GetNSTrackingAreaClass() NSTrackingAreaClass {
	return getNSTrackingAreaClass()
}

type NSTrackingAreaClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSTrackingAreaClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSTrackingAreaClass) Alloc() NSTrackingArea {
	rv := objc.Send[NSTrackingArea](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A region of a view that generates mouse-tracking and cursor-update events
// when the pointer is over that region.
//
// # Overview
//
// When creating a tracking-area object, you specify a rectangle (in the
// view’s coordinate system), an owning object, and one or more options,
// along with (optionally) a dictionary of data. After it’s created, you add
// the tracking-area object to a view using the [NSView.AddTrackingArea]
// method. Depending on the options specified, the owner of the tracking area
// receives [NSResponder.MouseEntered], [NSResponder.MouseExited],
// [NSResponder.MouseMoved], and [NSResponder.CursorUpdate] messages when the
// mouse cursor enters, moves within, and leaves the tracking area. Currently
// the tracking area is restricted to rectangles.
//
// An [NSTrackingArea] object belongs to its view rather than to its window.
// Consequently, you can add and remove tracking rectangles without needing to
// worry if the view has been added to a window. In addition, this design
// makes it possible for the AppKit to compute the geometry of tracking areas
// automatically when a view moves and, in some cases, when a view changes
// size.
//
// Using [NSTrackingArea], you can configure the scope of activity for mouse
// tracking. There are four options:
//
// - The tracking area is active only when the view is first responder. - The
// tracking area is active when the view is in the key window. - The tracking
// area is active when the application is active. - The tracking area is
// active always (even when the application is inactive).
//
// Other options for [NSTrackingArea] objects include specifying that the
// tracking area should be synchronized with the visible rectangle of the view
// ([NSView.VisibleRect]) and for generating “ and `mouseExited`: events when
// the mouse is dragged.
//
// Other [NSView] methods related to [NSTrackingArea] objects (in addition to
// [NSView.AddTrackingArea]) include [NSView.RemoveTrackingArea] and
// [NSView.UpdateTrackingAreas]. Views can override the latter method to
// recompute and replace their [NSTrackingArea] objects in certain situations,
// such as a change in the size of the `visibleRect`.
//
// # Initializing the Tracking-Area Object
//
//   - [NSTrackingArea.InitWithRectOptionsOwnerUserInfo]: Initializes and returns an object defining a region of a view to receive mouse-tracking events, mouse-moved events, cursor-update events, or possibly all these events.
//
// # Getting Object Attributes
//
//   - [NSTrackingArea.Options]: The options specified for the receiver.
//   - [NSTrackingArea.Owner]: The object owning the receiver, which is the recipient of mouse-tracking, mouse-movement, and cursor-update messages.
//   - [NSTrackingArea.Rect]: The rectangle defining the area encompassed by the receiver.
//   - [NSTrackingArea.UserInfo]: The dictionary containing the data associated with the receiver when it was created.
//
// See: https://developer.apple.com/documentation/AppKit/NSTrackingArea
type NSTrackingArea struct {
	objectivec.Object
}

// NSTrackingAreaFromID constructs a [NSTrackingArea] from an objc.ID.
//
// A region of a view that generates mouse-tracking and cursor-update events
// when the pointer is over that region.
func NSTrackingAreaFromID(id objc.ID) NSTrackingArea {
	return NSTrackingArea{objectivec.Object{ID: id}}
}

// NOTE: NSTrackingArea adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSTrackingArea] class.
//
// # Initializing the Tracking-Area Object
//
//   - [INSTrackingArea.InitWithRectOptionsOwnerUserInfo]: Initializes and returns an object defining a region of a view to receive mouse-tracking events, mouse-moved events, cursor-update events, or possibly all these events.
//
// # Getting Object Attributes
//
//   - [INSTrackingArea.Options]: The options specified for the receiver.
//   - [INSTrackingArea.Owner]: The object owning the receiver, which is the recipient of mouse-tracking, mouse-movement, and cursor-update messages.
//   - [INSTrackingArea.Rect]: The rectangle defining the area encompassed by the receiver.
//   - [INSTrackingArea.UserInfo]: The dictionary containing the data associated with the receiver when it was created.
//
// See: https://developer.apple.com/documentation/AppKit/NSTrackingArea
type INSTrackingArea interface {
	objectivec.IObject

	// Topic: Initializing the Tracking-Area Object

	// Initializes and returns an object defining a region of a view to receive mouse-tracking events, mouse-moved events, cursor-update events, or possibly all these events.
	InitWithRectOptionsOwnerUserInfo(rect corefoundation.CGRect, options NSTrackingAreaOptions, owner objectivec.IObject, userInfo foundation.INSDictionary) NSTrackingArea

	// Topic: Getting Object Attributes

	// The options specified for the receiver.
	Options() NSTrackingAreaOptions
	// The object owning the receiver, which is the recipient of mouse-tracking, mouse-movement, and cursor-update messages.
	Owner() objectivec.IObject
	// The rectangle defining the area encompassed by the receiver.
	Rect() corefoundation.CGRect
	// The dictionary containing the data associated with the receiver when it was created.
	UserInfo() foundation.INSDictionary

	// The owner receives messages regardless of first-responder status, window status, or application status. The
	ActiveAlways() NSTrackingAreaOptions
	SetNSTrackingActiveAlways(value NSTrackingAreaOptions)
	// The owner receives messages when the application is active. This value specifies when the tracking area defined by an
	ActiveInActiveApp() NSTrackingAreaOptions
	SetNSTrackingActiveInActiveApp(value NSTrackingAreaOptions)
	// The owner receives messages when the view is in the key window. This value specifies when the tracking area defined by an
	ActiveInKeyWindow() NSTrackingAreaOptions
	SetNSTrackingActiveInKeyWindow(value NSTrackingAreaOptions)
	// The owner receives messages when the view is the first responder. This value specifies when the tracking area defined by an
	ActiveWhenFirstResponder() NSTrackingAreaOptions
	SetNSTrackingActiveWhenFirstResponder(value NSTrackingAreaOptions)
	// The first event is generated when the cursor leaves the tracking area, regardless if the cursor is inside the area when the
	AssumeInside() NSTrackingAreaOptions
	SetNSTrackingAssumeInside(value NSTrackingAreaOptions)
	// A tracking option that receives events when the mouse cursor enters and exits the tracking area.
	CursorUpdate() NSTrackingAreaOptions
	SetNSTrackingCursorUpdate(value NSTrackingAreaOptions)
	// The owner receives
	EnabledDuringMouseDrag() NSTrackingAreaOptions
	SetNSTrackingEnabledDuringMouseDrag(value NSTrackingAreaOptions)
	// Mouse tracking occurs only in the visible rectangle of the view—in other words, that region of the tracking rectangle that is unobscured. Otherwise, the entire tracking area is active regardless of overlapping views. The
	InVisibleRect() NSTrackingAreaOptions
	SetNSTrackingInVisibleRect(value NSTrackingAreaOptions)
	// The owner of the tracking area receives
	MouseEnteredAndExited() NSTrackingAreaOptions
	SetNSTrackingMouseEnteredAndExited(value NSTrackingAreaOptions)
	// The owner of the tracking area receives
	MouseMoved() NSTrackingAreaOptions
	SetNSTrackingMouseMoved(value NSTrackingAreaOptions)
	InitWithCoder(coder foundation.INSCoder) NSTrackingArea
	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (t NSTrackingArea) Init() NSTrackingArea {
	rv := objc.Send[NSTrackingArea](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t NSTrackingArea) Autorelease() NSTrackingArea {
	rv := objc.Send[NSTrackingArea](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSTrackingArea creates a new NSTrackingArea instance.
func NewNSTrackingArea() NSTrackingArea {
	class := getNSTrackingAreaClass()
	rv := objc.Send[NSTrackingArea](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/AppKit/NSTrackingArea/init(coder:)
func NewTrackingAreaWithCoder(coder foundation.INSCoder) NSTrackingArea {
	instance := getNSTrackingAreaClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSTrackingAreaFromID(rv)
}

// Initializes and returns an object defining a region of a view to receive
// mouse-tracking events, mouse-moved events, cursor-update events, or
// possibly all these events.
//
// rect: A rectangle that defines a region of a target view, in the view’s
// coordinate system, for tracking events related to mouse tracking and cursor
// updating. The specified rectangle should not exceed the view’s bounds
// rectangle.
//
// options: One or more constants that specify the type of tracking area, the
// situations when the area is active, and special behaviors of the tracking
// area. See the description of [NSTrackingArea.Options] and related constants
// for details. You must specify one or more options for the initialized
// object for the type of tracking area and for when the tracking area is
// active; zero is not a valid value.
//
// owner: The object to receive the requested mouse-tracking, mouse-moved, or
// cursor-update messages. It does not necessarily have to be the view
// associated with the created [NSTrackingArea] object, but should be an
// object capable of responding to the [NSResponder] methods
// [NSResponder.MouseEntered], [NSResponder.MouseExited],
// [NSResponder.MouseMoved], and [NSResponder.CursorUpdate].
//
// userInfo: A dictionary containing arbitrary data for each mouse-entered,
// mouse-exited, and cursor-update event. When handling such an event you can
// obtain the dictionary by sending [NSEvent.UserData] to the [NSEvent]
// object. (The dictionary is not available for mouse-moved events.) This
// parameter may be `nil`.
//
// # Return Value
//
// The newly-initialized tracking area object.
//
// # Discussion
//
// After creating and initializing an [NSTrackingArea] object with this
// method, you must add it to a target view using the [NSView.AddTrackingArea]
// method. When changes in the view require changes in the geometry of its
// tracking areas, the Application Kit invokes [NSView.UpdateTrackingAreas].
// The view should implement this method to replace the current
// [NSTrackingArea] object with one with a recomputed area.
//
// # Special Considerations
//
// Beginning with OS X v10.5, the
// [NSTrackingArea.InitWithRectOptionsOwnerUserInfo], along with the
// [NSView.AddTrackingArea] method of [NSView], replace the [NSView] method
// [NSView.AddTrackingRectOwnerUserDataAssumeInside].
//
// See: https://developer.apple.com/documentation/AppKit/NSTrackingArea/init(rect:options:owner:userInfo:)
//
// [NSTrackingArea.Options]: https://developer.apple.com/documentation/AppKit/NSTrackingArea/Options-swift.struct
func NewTrackingAreaWithRectOptionsOwnerUserInfo(rect corefoundation.CGRect, options NSTrackingAreaOptions, owner objectivec.IObject, userInfo foundation.INSDictionary) NSTrackingArea {
	instance := getNSTrackingAreaClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithRect:options:owner:userInfo:"), rect, options, owner, userInfo)
	return NSTrackingAreaFromID(rv)
}

// Initializes and returns an object defining a region of a view to receive
// mouse-tracking events, mouse-moved events, cursor-update events, or
// possibly all these events.
//
// rect: A rectangle that defines a region of a target view, in the view’s
// coordinate system, for tracking events related to mouse tracking and cursor
// updating. The specified rectangle should not exceed the view’s bounds
// rectangle.
//
// options: One or more constants that specify the type of tracking area, the
// situations when the area is active, and special behaviors of the tracking
// area. See the description of [NSTrackingArea.Options] and related constants
// for details. You must specify one or more options for the initialized
// object for the type of tracking area and for when the tracking area is
// active; zero is not a valid value.
//
// owner: The object to receive the requested mouse-tracking, mouse-moved, or
// cursor-update messages. It does not necessarily have to be the view
// associated with the created [NSTrackingArea] object, but should be an
// object capable of responding to the [NSResponder] methods
// [NSResponder.MouseEntered], [NSResponder.MouseExited],
// [NSResponder.MouseMoved], and [NSResponder.CursorUpdate].
//
// userInfo: A dictionary containing arbitrary data for each mouse-entered,
// mouse-exited, and cursor-update event. When handling such an event you can
// obtain the dictionary by sending [NSEvent.UserData] to the [NSEvent]
// object. (The dictionary is not available for mouse-moved events.) This
// parameter may be `nil`.
//
// # Return Value
//
// The newly-initialized tracking area object.
//
// # Discussion
//
// After creating and initializing an [NSTrackingArea] object with this
// method, you must add it to a target view using the [NSView.AddTrackingArea]
// method. When changes in the view require changes in the geometry of its
// tracking areas, the Application Kit invokes [NSView.UpdateTrackingAreas].
// The view should implement this method to replace the current
// [NSTrackingArea] object with one with a recomputed area.
//
// # Special Considerations
//
// Beginning with OS X v10.5, the
// [NSTrackingArea.InitWithRectOptionsOwnerUserInfo], along with the
// [NSView.AddTrackingArea] method of [NSView], replace the [NSView] method
// [NSView.AddTrackingRectOwnerUserDataAssumeInside].
//
// See: https://developer.apple.com/documentation/AppKit/NSTrackingArea/init(rect:options:owner:userInfo:)
//
// [NSTrackingArea.Options]: https://developer.apple.com/documentation/AppKit/NSTrackingArea/Options-swift.struct
func (t NSTrackingArea) InitWithRectOptionsOwnerUserInfo(rect corefoundation.CGRect, options NSTrackingAreaOptions, owner objectivec.IObject, userInfo foundation.INSDictionary) NSTrackingArea {
	rv := objc.Send[NSTrackingArea](t.ID, objc.Sel("initWithRect:options:owner:userInfo:"), rect, options, owner, userInfo)
	return rv
}

// See: https://developer.apple.com/documentation/AppKit/NSTrackingArea/init(coder:)
func (t NSTrackingArea) InitWithCoder(coder foundation.INSCoder) NSTrackingArea {
	rv := objc.Send[NSTrackingArea](t.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}
func (t NSTrackingArea) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](t.ID, objc.Sel("encodeWithCoder:"), coder)
}

// The options specified for the receiver.
//
// # Discussion
//
// The options for an [NSTrackingArea] object are specified when the object is
// created.
//
// See: https://developer.apple.com/documentation/AppKit/NSTrackingArea/options-swift.property
func (t NSTrackingArea) Options() NSTrackingAreaOptions {
	rv := objc.Send[NSTrackingAreaOptions](t.ID, objc.Sel("options"))
	return NSTrackingAreaOptions(rv)
}

// The object owning the receiver, which is the recipient of mouse-tracking,
// mouse-movement, and cursor-update messages.
//
// See: https://developer.apple.com/documentation/AppKit/NSTrackingArea/owner
func (t NSTrackingArea) Owner() objectivec.IObject {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("owner"))
	return objectivec.Object{ID: rv}
}

// The rectangle defining the area encompassed by the receiver.
//
// # Discussion
//
// The rectangle is specified in the local coordinate system of the associated
// view. If the [NSTrackingInVisibleRect] option is specified, the receiver is
// automatically synchronized with changes in the view’s visible area
// ([NSView.VisibleRect]) and the value of this property is ignored.
//
// See: https://developer.apple.com/documentation/AppKit/NSTrackingArea/rect
func (t NSTrackingArea) Rect() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](t.ID, objc.Sel("rect"))
	return corefoundation.CGRect(rv)
}

// The dictionary containing the data associated with the receiver when it was
// created.
//
// # Discussion
//
// You can obtain this dictionary per event in each [NSResponder.MouseEntered]
// and [NSResponder.MouseExited] method by querying the passed-in [NSEvent]
// object with `[[event trackingArea] userData]`.
//
// See: https://developer.apple.com/documentation/AppKit/NSTrackingArea/userInfo
func (t NSTrackingArea) UserInfo() foundation.INSDictionary {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("userInfo"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}

// The owner receives messages regardless of first-responder status, window
// status, or application status. The
//
// See: https://developer.apple.com/documentation/appkit/nstrackingarea/options-swift.struct/activealways
func (t NSTrackingArea) ActiveAlways() NSTrackingAreaOptions {
	rv := objc.Send[NSTrackingAreaOptions](t.ID, objc.Sel("NSTrackingActiveAlways"))
	return NSTrackingAreaOptions(rv)
}
func (t NSTrackingArea) SetNSTrackingActiveAlways(value NSTrackingAreaOptions) {
	objc.Send[struct{}](t.ID, objc.Sel("setNSTrackingActiveAlways:"), value)
}

// The owner receives messages when the application is active. This value
// specifies when the tracking area defined by an
//
// See: https://developer.apple.com/documentation/appkit/nstrackingarea/options-swift.struct/activeinactiveapp
func (t NSTrackingArea) ActiveInActiveApp() NSTrackingAreaOptions {
	rv := objc.Send[NSTrackingAreaOptions](t.ID, objc.Sel("NSTrackingActiveInActiveApp"))
	return NSTrackingAreaOptions(rv)
}
func (t NSTrackingArea) SetNSTrackingActiveInActiveApp(value NSTrackingAreaOptions) {
	objc.Send[struct{}](t.ID, objc.Sel("setNSTrackingActiveInActiveApp:"), value)
}

// The owner receives messages when the view is in the key window. This value
// specifies when the tracking area defined by an
//
// See: https://developer.apple.com/documentation/appkit/nstrackingarea/options-swift.struct/activeinkeywindow
func (t NSTrackingArea) ActiveInKeyWindow() NSTrackingAreaOptions {
	rv := objc.Send[NSTrackingAreaOptions](t.ID, objc.Sel("NSTrackingActiveInKeyWindow"))
	return NSTrackingAreaOptions(rv)
}
func (t NSTrackingArea) SetNSTrackingActiveInKeyWindow(value NSTrackingAreaOptions) {
	objc.Send[struct{}](t.ID, objc.Sel("setNSTrackingActiveInKeyWindow:"), value)
}

// The owner receives messages when the view is the first responder. This
// value specifies when the tracking area defined by an
//
// See: https://developer.apple.com/documentation/appkit/nstrackingarea/options-swift.struct/activewhenfirstresponder
func (t NSTrackingArea) ActiveWhenFirstResponder() NSTrackingAreaOptions {
	rv := objc.Send[NSTrackingAreaOptions](t.ID, objc.Sel("NSTrackingActiveWhenFirstResponder"))
	return NSTrackingAreaOptions(rv)
}
func (t NSTrackingArea) SetNSTrackingActiveWhenFirstResponder(value NSTrackingAreaOptions) {
	objc.Send[struct{}](t.ID, objc.Sel("setNSTrackingActiveWhenFirstResponder:"), value)
}

// The first event is generated when the cursor leaves the tracking area,
// regardless if the cursor is inside the area when the
//
// See: https://developer.apple.com/documentation/appkit/nstrackingarea/options-swift.struct/assumeinside
func (t NSTrackingArea) AssumeInside() NSTrackingAreaOptions {
	rv := objc.Send[NSTrackingAreaOptions](t.ID, objc.Sel("NSTrackingAssumeInside"))
	return NSTrackingAreaOptions(rv)
}
func (t NSTrackingArea) SetNSTrackingAssumeInside(value NSTrackingAreaOptions) {
	objc.Send[struct{}](t.ID, objc.Sel("setNSTrackingAssumeInside:"), value)
}

// A tracking option that receives events when the mouse cursor enters and
// exits the tracking area.
//
// See: https://developer.apple.com/documentation/appkit/nstrackingarea/options-swift.struct/cursorupdate
func (t NSTrackingArea) CursorUpdate() NSTrackingAreaOptions {
	rv := objc.Send[NSTrackingAreaOptions](t.ID, objc.Sel("NSTrackingCursorUpdate"))
	return NSTrackingAreaOptions(rv)
}
func (t NSTrackingArea) SetNSTrackingCursorUpdate(value NSTrackingAreaOptions) {
	objc.Send[struct{}](t.ID, objc.Sel("setNSTrackingCursorUpdate:"), value)
}

// The owner receives
//
// See: https://developer.apple.com/documentation/appkit/nstrackingarea/options-swift.struct/enabledduringmousedrag
func (t NSTrackingArea) EnabledDuringMouseDrag() NSTrackingAreaOptions {
	rv := objc.Send[NSTrackingAreaOptions](t.ID, objc.Sel("NSTrackingEnabledDuringMouseDrag"))
	return NSTrackingAreaOptions(rv)
}
func (t NSTrackingArea) SetNSTrackingEnabledDuringMouseDrag(value NSTrackingAreaOptions) {
	objc.Send[struct{}](t.ID, objc.Sel("setNSTrackingEnabledDuringMouseDrag:"), value)
}

// Mouse tracking occurs only in the visible rectangle of the view—in other
// words, that region of the tracking rectangle that is unobscured. Otherwise,
// the entire tracking area is active regardless of overlapping views. The
//
// See: https://developer.apple.com/documentation/appkit/nstrackingarea/options-swift.struct/invisiblerect
func (t NSTrackingArea) InVisibleRect() NSTrackingAreaOptions {
	rv := objc.Send[NSTrackingAreaOptions](t.ID, objc.Sel("NSTrackingInVisibleRect"))
	return NSTrackingAreaOptions(rv)
}
func (t NSTrackingArea) SetNSTrackingInVisibleRect(value NSTrackingAreaOptions) {
	objc.Send[struct{}](t.ID, objc.Sel("setNSTrackingInVisibleRect:"), value)
}

// The owner of the tracking area receives
//
// See: https://developer.apple.com/documentation/appkit/nstrackingarea/options-swift.struct/mouseenteredandexited
func (t NSTrackingArea) MouseEnteredAndExited() NSTrackingAreaOptions {
	rv := objc.Send[NSTrackingAreaOptions](t.ID, objc.Sel("NSTrackingMouseEnteredAndExited"))
	return NSTrackingAreaOptions(rv)
}
func (t NSTrackingArea) SetNSTrackingMouseEnteredAndExited(value NSTrackingAreaOptions) {
	objc.Send[struct{}](t.ID, objc.Sel("setNSTrackingMouseEnteredAndExited:"), value)
}

// The owner of the tracking area receives
//
// See: https://developer.apple.com/documentation/appkit/nstrackingarea/options-swift.struct/mousemoved
func (t NSTrackingArea) MouseMoved() NSTrackingAreaOptions {
	rv := objc.Send[NSTrackingAreaOptions](t.ID, objc.Sel("NSTrackingMouseMoved"))
	return NSTrackingAreaOptions(rv)
}
func (t NSTrackingArea) SetNSTrackingMouseMoved(value NSTrackingAreaOptions) {
	objc.Send[struct{}](t.ID, objc.Sel("setNSTrackingMouseMoved:"), value)
}
