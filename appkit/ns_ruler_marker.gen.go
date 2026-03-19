// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSRulerMarker] class.
var (
	_NSRulerMarkerClass     NSRulerMarkerClass
	_NSRulerMarkerClassOnce sync.Once
)

func getNSRulerMarkerClass() NSRulerMarkerClass {
	_NSRulerMarkerClassOnce.Do(func() {
		_NSRulerMarkerClass = NSRulerMarkerClass{class: objc.GetClass("NSRulerMarker")}
	})
	return _NSRulerMarkerClass
}

// GetNSRulerMarkerClass returns the class object for NSRulerMarker.
func GetNSRulerMarkerClass() NSRulerMarkerClass {
	return getNSRulerMarkerClass()
}

type NSRulerMarkerClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSRulerMarkerClass) Alloc() NSRulerMarker {
	rv := objc.Send[NSRulerMarker](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A symbol on a ruler view, indicating a location for the graphics element it
// represents in the client of the ruler view.
//
// # Overview
// 
// An example of a marker is the representation of a margin or tab setting, or
// the edges of a graphic on the page.
//
// # Creating instances
//
//   - [NSRulerMarker.InitWithRulerViewMarkerLocationImageImageOrigin]: Initializes a newly allocated ruler marker, associating it with (but not adding it to) a specified ruler view and assigning the attributes provided.
//
// # Getting the ruler view
//
//   - [NSRulerMarker.Ruler]: The receiver’s ruler view.
//
// # Setting the image
//
//   - [NSRulerMarker.Image]: The receiver’s image.
//   - [NSRulerMarker.SetImage]
//   - [NSRulerMarker.ImageOrigin]: The point in the receiver’s image that is positioned at the receiver’s location on the ruler view.
//   - [NSRulerMarker.SetImageOrigin]
//   - [NSRulerMarker.ImageRectInRuler]: The rectangle occupied by the receiver’s image.
//   - [NSRulerMarker.ThicknessRequiredInRuler]: The amount of the receiver’s image that’s displayed above or to the left of the ruler view’s baseline.
//
// # Setting movability
//
//   - [NSRulerMarker.Movable]: A Boolean that indicates whether the user can move the receiver in its ruler view.
//   - [NSRulerMarker.SetMovable]
//   - [NSRulerMarker.Removable]: A Boolean that indicates whether the user can remove the receiver from its ruler view.
//   - [NSRulerMarker.SetRemovable]
//
// # Setting the location
//
//   - [NSRulerMarker.MarkerLocation]: The location of the receiver in the coordinate system of the ruler view’s client view.
//   - [NSRulerMarker.SetMarkerLocation]
//
// # Setting the represented object
//
//   - [NSRulerMarker.RepresentedObject]: The object the receiver represents.
//   - [NSRulerMarker.SetRepresentedObject]
//
// # Drawing and event handling
//
//   - [NSRulerMarker.DrawRect]: Draws the receiver’s image that appears in the supplied rectangle.
//   - [NSRulerMarker.Dragging]: A Boolean that indicates whether the receiver is being dragged.
//   - [NSRulerMarker.TrackMouseAdding]: Handles user manipulation of the receiver in its ruler view.
//
// # Initializers
//
//   - [NSRulerMarker.InitWithCoder]
//
// See: https://developer.apple.com/documentation/AppKit/NSRulerMarker
type NSRulerMarker struct {
	objectivec.Object
}

// NSRulerMarkerFromID constructs a [NSRulerMarker] from an objc.ID.
//
// A symbol on a ruler view, indicating a location for the graphics element it
// represents in the client of the ruler view.
func NSRulerMarkerFromID(id objc.ID) NSRulerMarker {
	return NSRulerMarker{objectivec.Object{ID: id}}
}
// NOTE: NSRulerMarker adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSRulerMarker] class.
//
// # Creating instances
//
//   - [INSRulerMarker.InitWithRulerViewMarkerLocationImageImageOrigin]: Initializes a newly allocated ruler marker, associating it with (but not adding it to) a specified ruler view and assigning the attributes provided.
//
// # Getting the ruler view
//
//   - [INSRulerMarker.Ruler]: The receiver’s ruler view.
//
// # Setting the image
//
//   - [INSRulerMarker.Image]: The receiver’s image.
//   - [INSRulerMarker.SetImage]
//   - [INSRulerMarker.ImageOrigin]: The point in the receiver’s image that is positioned at the receiver’s location on the ruler view.
//   - [INSRulerMarker.SetImageOrigin]
//   - [INSRulerMarker.ImageRectInRuler]: The rectangle occupied by the receiver’s image.
//   - [INSRulerMarker.ThicknessRequiredInRuler]: The amount of the receiver’s image that’s displayed above or to the left of the ruler view’s baseline.
//
// # Setting movability
//
//   - [INSRulerMarker.Movable]: A Boolean that indicates whether the user can move the receiver in its ruler view.
//   - [INSRulerMarker.SetMovable]
//   - [INSRulerMarker.Removable]: A Boolean that indicates whether the user can remove the receiver from its ruler view.
//   - [INSRulerMarker.SetRemovable]
//
// # Setting the location
//
//   - [INSRulerMarker.MarkerLocation]: The location of the receiver in the coordinate system of the ruler view’s client view.
//   - [INSRulerMarker.SetMarkerLocation]
//
// # Setting the represented object
//
//   - [INSRulerMarker.RepresentedObject]: The object the receiver represents.
//   - [INSRulerMarker.SetRepresentedObject]
//
// # Drawing and event handling
//
//   - [INSRulerMarker.DrawRect]: Draws the receiver’s image that appears in the supplied rectangle.
//   - [INSRulerMarker.Dragging]: A Boolean that indicates whether the receiver is being dragged.
//   - [INSRulerMarker.TrackMouseAdding]: Handles user manipulation of the receiver in its ruler view.
//
// # Initializers
//
//   - [INSRulerMarker.InitWithCoder]
//
// See: https://developer.apple.com/documentation/AppKit/NSRulerMarker
type INSRulerMarker interface {
	objectivec.IObject

	// Topic: Creating instances

	// Initializes a newly allocated ruler marker, associating it with (but not adding it to) a specified ruler view and assigning the attributes provided.
	InitWithRulerViewMarkerLocationImageImageOrigin(ruler INSRulerView, location float64, image INSImage, imageOrigin corefoundation.CGPoint) NSRulerMarker

	// Topic: Getting the ruler view

	// The receiver’s ruler view.
	Ruler() INSRulerView

	// Topic: Setting the image

	// The receiver’s image.
	Image() INSImage
	SetImage(value INSImage)
	// The point in the receiver’s image that is positioned at the receiver’s location on the ruler view.
	ImageOrigin() corefoundation.CGPoint
	SetImageOrigin(value corefoundation.CGPoint)
	// The rectangle occupied by the receiver’s image.
	ImageRectInRuler() corefoundation.CGRect
	// The amount of the receiver’s image that’s displayed above or to the left of the ruler view’s baseline.
	ThicknessRequiredInRuler() float64

	// Topic: Setting movability

	// A Boolean that indicates whether the user can move the receiver in its ruler view.
	Movable() bool
	SetMovable(value bool)
	// A Boolean that indicates whether the user can remove the receiver from its ruler view.
	Removable() bool
	SetRemovable(value bool)

	// Topic: Setting the location

	// The location of the receiver in the coordinate system of the ruler view’s client view.
	MarkerLocation() float64
	SetMarkerLocation(value float64)

	// Topic: Setting the represented object

	// The object the receiver represents.
	RepresentedObject() foundation.NSCopying
	SetRepresentedObject(value foundation.NSCopying)

	// Topic: Drawing and event handling

	// Draws the receiver’s image that appears in the supplied rectangle.
	DrawRect(rect corefoundation.CGRect)
	// A Boolean that indicates whether the receiver is being dragged.
	Dragging() bool
	// Handles user manipulation of the receiver in its ruler view.
	TrackMouseAdding(mouseDownEvent INSEvent, isAdding bool) bool

	// Topic: Initializers

	InitWithCoder(coder foundation.INSCoder) NSRulerMarker

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (r NSRulerMarker) Init() NSRulerMarker {
	rv := objc.Send[NSRulerMarker](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r NSRulerMarker) Autorelease() NSRulerMarker {
	rv := objc.Send[NSRulerMarker](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSRulerMarker creates a new NSRulerMarker instance.
func NewNSRulerMarker() NSRulerMarker {
	class := getNSRulerMarkerClass()
	rv := objc.Send[NSRulerMarker](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/AppKit/NSRulerMarker/init(coder:)
func NewRulerMarkerWithCoder(coder foundation.INSCoder) NSRulerMarker {
	instance := getNSRulerMarkerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSRulerMarkerFromID(rv)
}

// Initializes a newly allocated ruler marker, associating it with (but not
// adding it to) a specified ruler view and assigning the attributes provided.
//
// ruler: The ruler view with which to associate the ruler marker. This method raises
// an [NSInvalidArgumentException] if `aRulerView` is `nil`.
//
// location: The x or y position of the marker in the client view’s coordinate system,
// depending on whether the ruler view is horizontal or vertical.
//
// image: The image displayed at the marker location. This method raises an
// [NSInvalidArgumentException] if `anImage` is `nil`.
//
// imageOrigin: The point within the image positioned at the marker location, expressed in
// pixels relative to the lower-left corner of the image.
//
// # Return Value
// 
// An initialized ruler marker object.
//
// # Discussion
// 
// The image used to draw the marker must be appropriate for the orientation
// of the ruler. Markers may need to look different on a horizontal ruler than
// on a vertical ruler, and the ruler view neither scales nor rotates the
// images.
// 
// To add the new ruler marker to `aRulerView`, use either of NSRulerView’s
// [AddMarker] or [TrackMarkerWithMouseEvent] methods. [AddMarker] immediately
// puts the marker on the ruler, while [TrackMarkerWithMouseEvent] allows the
// client view to intercede in the addition and placement of the marker.
// 
// A new ruler marker can be moved on its ruler view, but not removed. Use
// [Movable] and [Removable] to change these attributes. The new ruler marker
// also has no represented object; use [RepresentedObject] to set one.
// 
// This method is the designated initializer for the NSRulerMarker class.
//
// See: https://developer.apple.com/documentation/AppKit/NSRulerMarker/init(rulerView:markerLocation:image:imageOrigin:)
func NewRulerMarkerWithRulerViewMarkerLocationImageImageOrigin(ruler INSRulerView, location float64, image INSImage, imageOrigin corefoundation.CGPoint) NSRulerMarker {
	instance := getNSRulerMarkerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithRulerView:markerLocation:image:imageOrigin:"), ruler, location, image, imageOrigin)
	return NSRulerMarkerFromID(rv)
}

// Initializes a newly allocated ruler marker, associating it with (but not
// adding it to) a specified ruler view and assigning the attributes provided.
//
// ruler: The ruler view with which to associate the ruler marker. This method raises
// an [NSInvalidArgumentException] if `aRulerView` is `nil`.
//
// location: The x or y position of the marker in the client view’s coordinate system,
// depending on whether the ruler view is horizontal or vertical.
//
// image: The image displayed at the marker location. This method raises an
// [NSInvalidArgumentException] if `anImage` is `nil`.
//
// imageOrigin: The point within the image positioned at the marker location, expressed in
// pixels relative to the lower-left corner of the image.
//
// # Return Value
// 
// An initialized ruler marker object.
//
// # Discussion
// 
// The image used to draw the marker must be appropriate for the orientation
// of the ruler. Markers may need to look different on a horizontal ruler than
// on a vertical ruler, and the ruler view neither scales nor rotates the
// images.
// 
// To add the new ruler marker to `aRulerView`, use either of NSRulerView’s
// [AddMarker] or [TrackMarkerWithMouseEvent] methods. [AddMarker] immediately
// puts the marker on the ruler, while [TrackMarkerWithMouseEvent] allows the
// client view to intercede in the addition and placement of the marker.
// 
// A new ruler marker can be moved on its ruler view, but not removed. Use
// [Movable] and [Removable] to change these attributes. The new ruler marker
// also has no represented object; use [RepresentedObject] to set one.
// 
// This method is the designated initializer for the NSRulerMarker class.
//
// See: https://developer.apple.com/documentation/AppKit/NSRulerMarker/init(rulerView:markerLocation:image:imageOrigin:)
func (r NSRulerMarker) InitWithRulerViewMarkerLocationImageImageOrigin(ruler INSRulerView, location float64, image INSImage, imageOrigin corefoundation.CGPoint) NSRulerMarker {
	rv := objc.Send[NSRulerMarker](r.ID, objc.Sel("initWithRulerView:markerLocation:image:imageOrigin:"), ruler, location, image, imageOrigin)
	return rv
}
// Draws the receiver’s image that appears in the supplied rectangle.
//
// rect: The rectangle to be drawn, in the ruler view’s coordinate system.
//
// See: https://developer.apple.com/documentation/AppKit/NSRulerMarker/draw(_:)
func (r NSRulerMarker) DrawRect(rect corefoundation.CGRect) {
	objc.Send[objc.ID](r.ID, objc.Sel("drawRect:"), rect)
}
// Handles user manipulation of the receiver in its ruler view.
//
// mouseDownEvent: The event that represents the user manipulation being attempted on the
// ruler marker.
//
// isAdding: [true] to indicate that the receiver is a new marker being added to its
// ruler view, [false] otherwise.
// //
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Return Value
// 
// [true] if the user manipulation was allowed, [false] if it was not allowed.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// [NSRulerView] objects invoke this method automatically to add a new marker
// or to move or remove an existing marker. You should never need to invoke it
// directly.
// 
// If the receiver is a new marker being added to its ruler view (`flag` is
// [true]), the receiver queries the ruler view’s client before adding
// itself to the ruler view. If the client responds to
// [RulerViewShouldAddMarker] and that method returns [false], this method
// immediately returns [false], and the new marker isn’t added.
// 
// If the receiver is not a new marker being added to its ruler view (`flag`
// is [false]), this method attempts to move or remove an existing marker,
// once again based on responses from the ruler view’s client view. If the
// receiver is neither movable nor removable, this method immediately returns
// [false]. Further, if the ruler view’s client responds to
// [RulerViewShouldMoveMarker] and returns [false], this method returns
// [false], indicating the receiver can’t be moved.
// 
// If the receiver is being added or moved, this method queries the client
// view using [RulerViewWillAddMarkerAtLocation] or
// [RulerViewWillMoveMarkerToLocation], respectively. If the client responds
// to the method, the return value is used as the receiver’s location. These
// methods are invoked repeatedly as the receiver is dragged within the ruler
// view.
// 
// If the receiver is an existing marker being removed (dragged off the
// ruler), this method queries the client view using
// [RulerViewShouldRemoveMarker]. If the client responds to this method and
// returns [false], the marker is pinned to the ruler view’s baseline
// (following the cursor on the baseline if it’s movable).
// 
// When the user releases the mouse button, this method informs the client
// view of the marker’s new status using [RulerViewDidAddMarker],
// [RulerViewDidMoveMarker], or [RulerViewDidRemoveMarker] as appropriate. The
// client view can use this notification to set the marker’s represented
// object, modify its state and redisplay (for example, adjusting text layout
// around a new tab stop), or take whatever other action it might need.
// 
// If `flag` is [true] and the user dragged the new marker away from the
// ruler, the marker isn’t added, no message is sent, and this method
// returns [false].
// 
// See [Ruler and Paragraph Style Programming Topics] for more information on
// these client methods.
//
// [Ruler and Paragraph Style Programming Topics]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/Rulers/Rulers.html#//apple_ref/doc/uid/10000089i
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSRulerMarker/trackMouse(with:adding:)
func (r NSRulerMarker) TrackMouseAdding(mouseDownEvent INSEvent, isAdding bool) bool {
	rv := objc.Send[bool](r.ID, objc.Sel("trackMouse:adding:"), mouseDownEvent, isAdding)
	return rv
}
//
// See: https://developer.apple.com/documentation/AppKit/NSRulerMarker/init(coder:)
func (r NSRulerMarker) InitWithCoder(coder foundation.INSCoder) NSRulerMarker {
	rv := objc.Send[NSRulerMarker](r.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}
func (r NSRulerMarker) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](r.ID, objc.Sel("encodeWithCoder:"), coder)
}

// The receiver’s ruler view.
//
// See: https://developer.apple.com/documentation/AppKit/NSRulerMarker/ruler
func (r NSRulerMarker) Ruler() INSRulerView {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("ruler"))
	return NSRulerViewFromID(objc.ID(rv))
}
// The receiver’s image.
//
// # Discussion
// 
// The image used to draw the marker must be appropriate for the orientation
// of the ruler. Markers may need to look different on a horizontal ruler than
// on a vertical ruler, and the ruler view neither scales nor rotates the
// images.
//
// See: https://developer.apple.com/documentation/AppKit/NSRulerMarker/image
func (r NSRulerMarker) Image() INSImage {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("image"))
	return NSImageFromID(objc.ID(rv))
}
func (r NSRulerMarker) SetImage(value INSImage) {
	objc.Send[struct{}](r.ID, objc.Sel("setImage:"), value)
}
// The point in the receiver’s image that is positioned at the receiver’s
// location on the ruler view.
//
// # Discussion
// 
// For a horizontal ruler, the x coordinate of the image origin is aligned
// with the location of the marker, and the y coordinate lies on the baseline
// of the ruler. For vertical rulers, the y coordinate of the image origin is
// the location, and the x coordinate lies on the baseline.
//
// See: https://developer.apple.com/documentation/AppKit/NSRulerMarker/imageOrigin
func (r NSRulerMarker) ImageOrigin() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](r.ID, objc.Sel("imageOrigin"))
	return corefoundation.CGPoint(rv)
}
func (r NSRulerMarker) SetImageOrigin(value corefoundation.CGPoint) {
	objc.Send[struct{}](r.ID, objc.Sel("setImageOrigin:"), value)
}
// The rectangle occupied by the receiver’s image.
//
// # Discussion
// 
// The rectangle occupied by the receiver’s image, in the ruler view’s
// coordinate system, accounting for whether the ruler view’s coordinate
// system is flipped.
//
// See: https://developer.apple.com/documentation/AppKit/NSRulerMarker/imageRectInRuler
func (r NSRulerMarker) ImageRectInRuler() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](r.ID, objc.Sel("imageRectInRuler"))
	return corefoundation.CGRect(rv)
}
// The amount of the receiver’s image that’s displayed above or to the
// left of the ruler view’s baseline.
//
// # Discussion
// 
// The amount of the receiver’s image that’s displayed above or to the
// left of the ruler view’s baseline, the height for a horizontal ruler or
// width for a vertical ruler.
//
// See: https://developer.apple.com/documentation/AppKit/NSRulerMarker/thicknessRequiredInRuler
func (r NSRulerMarker) ThicknessRequiredInRuler() float64 {
	rv := objc.Send[float64](r.ID, objc.Sel("thicknessRequiredInRuler"))
	return rv
}
// A Boolean that indicates whether the user can move the receiver in its
// ruler view.
//
// # Discussion
// 
// [true] to allow the user to drag the marker image in the ruler, [false] to
// make it immobile.
// 
// By default, ruler markers are movable.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSRulerMarker/isMovable
func (r NSRulerMarker) Movable() bool {
	rv := objc.Send[bool](r.ID, objc.Sel("isMovable"))
	return rv
}
func (r NSRulerMarker) SetMovable(value bool) {
	objc.Send[struct{}](r.ID, objc.Sel("setMovable:"), value)
}
// A Boolean that indicates whether the user can remove the receiver from its
// ruler view.
//
// # Discussion
// 
// [true] to allow the user to drag the marker image off of the ruler and
// remove the marker, [false] to prevent the user from removing the marker.
// 
// By default ruler markers are not removable.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSRulerMarker/isRemovable
func (r NSRulerMarker) Removable() bool {
	rv := objc.Send[bool](r.ID, objc.Sel("isRemovable"))
	return rv
}
func (r NSRulerMarker) SetRemovable(value bool) {
	objc.Send[struct{}](r.ID, objc.Sel("setRemovable:"), value)
}
// The location of the receiver in the coordinate system of the ruler view’s
// client view.
//
// # Discussion
// 
// This is an x position for a horizontal ruler, a y position for a vertical
// ruler.
//
// See: https://developer.apple.com/documentation/AppKit/NSRulerMarker/markerLocation
func (r NSRulerMarker) MarkerLocation() float64 {
	rv := objc.Send[float64](r.ID, objc.Sel("markerLocation"))
	return rv
}
func (r NSRulerMarker) SetMarkerLocation(value float64) {
	objc.Send[struct{}](r.ID, objc.Sel("setMarkerLocation:"), value)
}
// The object the receiver represents.
//
// # Discussion
// 
// See [About Ruler Markers] for more information on the represented object.
//
// [About Ruler Markers]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/Rulers/Concepts/AboutRulerMarkers.html#//apple_ref/doc/uid/20000873
//
// See: https://developer.apple.com/documentation/AppKit/NSRulerMarker/representedObject
func (r NSRulerMarker) RepresentedObject() foundation.NSCopying {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("representedObject"))
	return foundation.NSCopyingObjectFromID(rv)
}
func (r NSRulerMarker) SetRepresentedObject(value foundation.NSCopying) {
	objc.Send[struct{}](r.ID, objc.Sel("setRepresentedObject:"), value)
}
// A Boolean that indicates whether the receiver is being dragged.
//
// # Discussion
// 
// [true] if the receiver is being dragged, [false] otherwise.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSRulerMarker/isDragging
func (r NSRulerMarker) Dragging() bool {
	rv := objc.Send[bool](r.ID, objc.Sel("isDragging"))
	return rv
}

