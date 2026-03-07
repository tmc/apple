// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSRulerView] class.
var (
	_NSRulerViewClass     NSRulerViewClass
	_NSRulerViewClassOnce sync.Once
)

func getNSRulerViewClass() NSRulerViewClass {
	_NSRulerViewClassOnce.Do(func() {
		_NSRulerViewClass = NSRulerViewClass{class: objc.GetClass("NSRulerView")}
	})
	return _NSRulerViewClass
}

// GetNSRulerViewClass returns the class object for NSRulerView.
func GetNSRulerViewClass() NSRulerViewClass {
	return getNSRulerViewClass()
}

type NSRulerViewClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSRulerViewClass) Alloc() NSRulerView {
	rv := objc.Send[NSRulerView](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// A ruler and the markers above or to the side of a scroll view’s document
// view.
//
// # Overview
// 
// Views within the scroll view can become clients of the ruler view, having
// it display markers for their elements, and receiving messages from the
// ruler view when the user manipulates the markers.
// 
// # Principal Attributes
// 
// - Displays markers that represent elements of the client view. - Displays
// in arbitrary units. - Provides for an accessory view containing extra
// controls.
// 
// # Creation
// 
// - [NSRulerView.HasHorizontalRuler] ([NSScrollView]) - [NSRulerView.HasVerticalRuler]
// ([NSScrollView]) - [NSRulerView.InitWithScrollViewOrientation] Designated initializer.
// 
// # Commonly Used Methods
// 
// [NSRulerView.ClientView]: Changes the ruler’s client view. [NSRulerView.Markers]: Sets the
// markers displayed by the ruler view. [NSRulerView.AccessoryView]: Sets the accessory
// view. [NSRulerView.TrackMarkerWithMouseEvent]: Allows the user to add a new marker.
// 
// # Overview
// 
// See NSRulerMarkerClientViewDelegation for delegate methods that may be of
// interest.
//
// # Creating a Ruler View
//
//   - [NSRulerView.InitWithScrollViewOrientation]: Initializes a newly allocated NSRulerView to have `orientation` ([NSHorizontalRuler] or [NSVerticalRuler]) within `aScrollView`.
//
// # Altering measurement units
//
//   - [NSRulerView.MeasurementUnits]: The measurement units used by the ruler to `unitName`.
//   - [NSRulerView.SetMeasurementUnits]
//
// # Setting the client view
//
//   - [NSRulerView.ClientView]: The receiver’s client view, if it has one.
//   - [NSRulerView.SetClientView]
//
// # Setting an accessory view
//
//   - [NSRulerView.AccessoryView]: The receiver’s accessory view to `aView`.
//   - [NSRulerView.SetAccessoryView]
//
// # Setting the zero mark position
//
//   - [NSRulerView.OriginOffset]: The distance to the zero hash mark from the bounds origin of the NSScrollView’s document view (not of the receiver’s client view), in the document view’s coordinate system.
//   - [NSRulerView.SetOriginOffset]
//
// # Adding and removing markers
//
//   - [NSRulerView.Markers]: The receiver’s ruler markers to `markers`, removing any existing ruler markers and not consulting with the client view about the new markers.
//   - [NSRulerView.SetMarkers]
//   - [NSRulerView.AddMarker]: Adds `aMarker` to the receiver, without consulting the client view for approval.
//   - [NSRulerView.RemoveMarker]: Removes `aMarker` from the receiver, without consulting the client view for approval.
//   - [NSRulerView.TrackMarkerWithMouseEvent]: Tracks the mouse to add `aMarker` based on the initial mouse-down or mouse-dragged event `theEvent`.
//
// # Drawing temporary ruler lines
//
//   - [NSRulerView.MoveRulerlineFromLocationToLocation]: Draws temporary lines in the ruler area.
//
// # Drawing
//
//   - [NSRulerView.DrawHashMarksAndLabelsInRect]: Draws the receiver’s hash marks and labels in `aRect`, which is expressed in the receiver’s coordinate system.
//   - [NSRulerView.DrawMarkersInRect]: Draws the receiver’s markers in `aRect`, which is expressed in the receiver’s coordinate system.
//   - [NSRulerView.InvalidateHashMarks]: Forces recalculation of the hash mark spacing for the next time the receiver is displayed.
//
// # Ruler layout
//
//   - [NSRulerView.ScrollView]: The NSScrollView that owns the receiver to `scrollView`, without retaining it.
//   - [NSRulerView.SetScrollView]
//   - [NSRulerView.Orientation]: The orientation of the receiver to `orientation`.
//   - [NSRulerView.SetOrientation]
//   - [NSRulerView.ReservedThicknessForAccessoryView]: The room available for the receiver’s accessory view to `thickness`.
//   - [NSRulerView.SetReservedThicknessForAccessoryView]
//   - [NSRulerView.ReservedThicknessForMarkers]: The room available for ruler markers to `thickness`.
//   - [NSRulerView.SetReservedThicknessForMarkers]
//   - [NSRulerView.RuleThickness]: The thickness of the area where ruler hash marks and labels are drawn.
//   - [NSRulerView.SetRuleThickness]
//   - [NSRulerView.RequiredThickness]: The thickness needed for proper tiling of the receiver within an NSScrollView.
//   - [NSRulerView.BaselineLocation]: The location of the receiver’s baseline, in its own coordinate system.
//   - [NSRulerView.Flipped]: A Boolean that indicates if the ruler view’s coordinate system is flipped.
//
// See: https://developer.apple.com/documentation/AppKit/NSRulerView
type NSRulerView struct {
	NSView
}

// NSRulerViewFromID constructs a [NSRulerView] from an objc.ID.
//
// A ruler and the markers above or to the side of a scroll view’s document
// view.
func NSRulerViewFromID(id objc.ID) NSRulerView {
	return NSRulerView{NSView: NSViewFromID(id)}
}
// NOTE: NSRulerView adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [NSRulerView] class.
//
// # Creating a Ruler View
//
//   - [INSRulerView.InitWithScrollViewOrientation]: Initializes a newly allocated NSRulerView to have `orientation` ([NSHorizontalRuler] or [NSVerticalRuler]) within `aScrollView`.
//
// # Altering measurement units
//
//   - [INSRulerView.MeasurementUnits]: The measurement units used by the ruler to `unitName`.
//   - [INSRulerView.SetMeasurementUnits]
//
// # Setting the client view
//
//   - [INSRulerView.ClientView]: The receiver’s client view, if it has one.
//   - [INSRulerView.SetClientView]
//
// # Setting an accessory view
//
//   - [INSRulerView.AccessoryView]: The receiver’s accessory view to `aView`.
//   - [INSRulerView.SetAccessoryView]
//
// # Setting the zero mark position
//
//   - [INSRulerView.OriginOffset]: The distance to the zero hash mark from the bounds origin of the NSScrollView’s document view (not of the receiver’s client view), in the document view’s coordinate system.
//   - [INSRulerView.SetOriginOffset]
//
// # Adding and removing markers
//
//   - [INSRulerView.Markers]: The receiver’s ruler markers to `markers`, removing any existing ruler markers and not consulting with the client view about the new markers.
//   - [INSRulerView.SetMarkers]
//   - [INSRulerView.AddMarker]: Adds `aMarker` to the receiver, without consulting the client view for approval.
//   - [INSRulerView.RemoveMarker]: Removes `aMarker` from the receiver, without consulting the client view for approval.
//   - [INSRulerView.TrackMarkerWithMouseEvent]: Tracks the mouse to add `aMarker` based on the initial mouse-down or mouse-dragged event `theEvent`.
//
// # Drawing temporary ruler lines
//
//   - [INSRulerView.MoveRulerlineFromLocationToLocation]: Draws temporary lines in the ruler area.
//
// # Drawing
//
//   - [INSRulerView.DrawHashMarksAndLabelsInRect]: Draws the receiver’s hash marks and labels in `aRect`, which is expressed in the receiver’s coordinate system.
//   - [INSRulerView.DrawMarkersInRect]: Draws the receiver’s markers in `aRect`, which is expressed in the receiver’s coordinate system.
//   - [INSRulerView.InvalidateHashMarks]: Forces recalculation of the hash mark spacing for the next time the receiver is displayed.
//
// # Ruler layout
//
//   - [INSRulerView.ScrollView]: The NSScrollView that owns the receiver to `scrollView`, without retaining it.
//   - [INSRulerView.SetScrollView]
//   - [INSRulerView.Orientation]: The orientation of the receiver to `orientation`.
//   - [INSRulerView.SetOrientation]
//   - [INSRulerView.ReservedThicknessForAccessoryView]: The room available for the receiver’s accessory view to `thickness`.
//   - [INSRulerView.SetReservedThicknessForAccessoryView]
//   - [INSRulerView.ReservedThicknessForMarkers]: The room available for ruler markers to `thickness`.
//   - [INSRulerView.SetReservedThicknessForMarkers]
//   - [INSRulerView.RuleThickness]: The thickness of the area where ruler hash marks and labels are drawn.
//   - [INSRulerView.SetRuleThickness]
//   - [INSRulerView.RequiredThickness]: The thickness needed for proper tiling of the receiver within an NSScrollView.
//   - [INSRulerView.BaselineLocation]: The location of the receiver’s baseline, in its own coordinate system.
//   - [INSRulerView.Flipped]: A Boolean that indicates if the ruler view’s coordinate system is flipped.
//
// See: https://developer.apple.com/documentation/AppKit/NSRulerView
type INSRulerView interface {
	INSView

	// Topic: Creating a Ruler View

	// Initializes a newly allocated NSRulerView to have `orientation` ([NSHorizontalRuler] or [NSVerticalRuler]) within `aScrollView`.
	InitWithScrollViewOrientation(scrollView INSScrollView, orientation NSRulerOrientation) NSRulerView

	// Topic: Altering measurement units

	// The measurement units used by the ruler to `unitName`.
	MeasurementUnits() NSRulerViewUnitName
	SetMeasurementUnits(value NSRulerViewUnitName)

	// Topic: Setting the client view

	// The receiver’s client view, if it has one.
	ClientView() INSView
	SetClientView(value INSView)

	// Topic: Setting an accessory view

	// The receiver’s accessory view to `aView`.
	AccessoryView() INSView
	SetAccessoryView(value INSView)

	// Topic: Setting the zero mark position

	// The distance to the zero hash mark from the bounds origin of the NSScrollView’s document view (not of the receiver’s client view), in the document view’s coordinate system.
	OriginOffset() float64
	SetOriginOffset(value float64)

	// Topic: Adding and removing markers

	// The receiver’s ruler markers to `markers`, removing any existing ruler markers and not consulting with the client view about the new markers.
	Markers() []NSRulerMarker
	SetMarkers(value []NSRulerMarker)
	// Adds `aMarker` to the receiver, without consulting the client view for approval.
	AddMarker(marker INSRulerMarker)
	// Removes `aMarker` from the receiver, without consulting the client view for approval.
	RemoveMarker(marker INSRulerMarker)
	// Tracks the mouse to add `aMarker` based on the initial mouse-down or mouse-dragged event `theEvent`.
	TrackMarkerWithMouseEvent(marker INSRulerMarker, event INSEvent) bool

	// Topic: Drawing temporary ruler lines

	// Draws temporary lines in the ruler area.
	MoveRulerlineFromLocationToLocation(oldLocation float64, newLocation float64)

	// Topic: Drawing

	// Draws the receiver’s hash marks and labels in `aRect`, which is expressed in the receiver’s coordinate system.
	DrawHashMarksAndLabelsInRect(rect corefoundation.CGRect)
	// Draws the receiver’s markers in `aRect`, which is expressed in the receiver’s coordinate system.
	DrawMarkersInRect(rect corefoundation.CGRect)
	// Forces recalculation of the hash mark spacing for the next time the receiver is displayed.
	InvalidateHashMarks()

	// Topic: Ruler layout

	// The NSScrollView that owns the receiver to `scrollView`, without retaining it.
	ScrollView() INSScrollView
	SetScrollView(value INSScrollView)
	// The orientation of the receiver to `orientation`.
	Orientation() NSRulerOrientation
	SetOrientation(value NSRulerOrientation)
	// The room available for the receiver’s accessory view to `thickness`.
	ReservedThicknessForAccessoryView() float64
	SetReservedThicknessForAccessoryView(value float64)
	// The room available for ruler markers to `thickness`.
	ReservedThicknessForMarkers() float64
	SetReservedThicknessForMarkers(value float64)
	// The thickness of the area where ruler hash marks and labels are drawn.
	RuleThickness() float64
	SetRuleThickness(value float64)
	// The thickness needed for proper tiling of the receiver within an NSScrollView.
	RequiredThickness() float64
	// The location of the receiver’s baseline, in its own coordinate system.
	BaselineLocation() float64
	// A Boolean that indicates if the ruler view’s coordinate system is flipped.
	Flipped() bool

	// A Boolean that indicates whether the scroll view keeps a horizontal ruler object.
	HasHorizontalRuler() bool
	SetHasHorizontalRuler(value bool)
	// A Boolean that indicates whether the scroll view keeps a vertical ruler object.
	HasVerticalRuler() bool
	SetHasVerticalRuler(value bool)
	EncodeWithCoder(coder foundation.INSCoder)
}





// Init initializes the instance.
func (r NSRulerView) Init() NSRulerView {
	rv := objc.Send[NSRulerView](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r NSRulerView) Autorelease() NSRulerView {
	rv := objc.Send[NSRulerView](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSRulerView creates a new NSRulerView instance.
func NewNSRulerView() NSRulerView {
	class := getNSRulerViewClass()
	rv := objc.Send[NSRulerView](objc.ID(class.class), objc.Sel("new"))
	return rv
}






//
// See: https://developer.apple.com/documentation/AppKit/NSRulerView/init(coder:)
func NewRulerViewWithCoder(coder foundation.INSCoder) NSRulerView {
	instance := getNSRulerViewClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSRulerViewFromID(rv)
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
func NewRulerViewWithFrame(frameRect corefoundation.CGRect) NSRulerView {
	instance := getNSRulerViewClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFrame:"), frameRect)
	return NSRulerViewFromID(rv)
}


// Initializes a newly allocated NSRulerView to have `orientation`
// ([NSHorizontalRuler] or [NSVerticalRuler]) within `aScrollView`.
//
// # Discussion
// 
// The new ruler view displays the user’s preferred measurement units and
// has no client, markers, or accessory view. Unlike most subclasses of
// NSView, no initial frame rectangle is given for NSRulerView; its containing
// NSScrollView adjusts its frame rectangle as needed.
// 
// This method is the designated initializer for the NSRulerView class.
// Returns an initialized object.
//
// See: https://developer.apple.com/documentation/AppKit/NSRulerView/init(scrollView:orientation:)
func NewRulerViewWithScrollViewOrientation(scrollView INSScrollView, orientation NSRulerOrientation) NSRulerView {
	instance := getNSRulerViewClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithScrollView:orientation:"), scrollView, orientation)
	return NSRulerViewFromID(rv)
}







// Initializes a newly allocated NSRulerView to have `orientation`
// ([NSHorizontalRuler] or [NSVerticalRuler]) within `aScrollView`.
//
// # Discussion
// 
// The new ruler view displays the user’s preferred measurement units and
// has no client, markers, or accessory view. Unlike most subclasses of
// NSView, no initial frame rectangle is given for NSRulerView; its containing
// NSScrollView adjusts its frame rectangle as needed.
// 
// This method is the designated initializer for the NSRulerView class.
// Returns an initialized object.
//
// See: https://developer.apple.com/documentation/AppKit/NSRulerView/init(scrollView:orientation:)
func (r NSRulerView) InitWithScrollViewOrientation(scrollView INSScrollView, orientation NSRulerOrientation) NSRulerView {
	rv := objc.Send[NSRulerView](r.ID, objc.Sel("initWithScrollView:orientation:"), scrollView, orientation)
	return rv
}

// Adds `aMarker` to the receiver, without consulting the client view for
// approval.
//
// # Discussion
// 
// Raises an [NSInternalInconsistencyException] if the receiver has no client
// view.
//
// See: https://developer.apple.com/documentation/AppKit/NSRulerView/addMarker(_:)
func (r NSRulerView) AddMarker(marker INSRulerMarker) {
	objc.Send[objc.ID](r.ID, objc.Sel("addMarker:"), marker)
}

// Removes `aMarker` from the receiver, without consulting the client view for
// approval.
//
// See: https://developer.apple.com/documentation/AppKit/NSRulerView/removeMarker(_:)
func (r NSRulerView) RemoveMarker(marker INSRulerMarker) {
	objc.Send[objc.ID](r.ID, objc.Sel("removeMarker:"), marker)
}

// Tracks the mouse to add `aMarker` based on the initial mouse-down or
// mouse-dragged event `theEvent`.
//
// # Discussion
// 
// Returns [true] if the receiver adds `aMarker`, [false] if it doesn’t.
// This method works by sending [TrackMouseAdding] to `aMarker` with
// `theEvent` and [true] as arguments.
// 
// An application typically invokes this method in one of two cases. In the
// simpler case, the client view can implement [NSRulerView] to invoke this
// method when the user presses the mouse button while the cursor is in the
// NSRulerView’s ruler area. This technique is appropriate when it’s clear
// what kind of marker will be added by clicking the ruler area. The second,
// more general, case involves the application providing a palette of
// different kinds of markers that can be dragged onto the ruler, from the
// ruler’s accessory view or from some other place. With this technique the
// palette tracks the cursor until it enters the ruler view, at which time it
// hands over control to the ruler view by invoking
// [TrackMarkerWithMouseEvent].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSRulerView/trackMarker(_:withMouseEvent:)
func (r NSRulerView) TrackMarkerWithMouseEvent(marker INSRulerMarker, event INSEvent) bool {
	rv := objc.Send[bool](r.ID, objc.Sel("trackMarker:withMouseEvent:"), marker, event)
	return rv
}

// Draws temporary lines in the ruler area.
//
// # Discussion
// 
// If `oldLoc` is 0 or greater, erases the ruler line at that location; if
// `newLoc` is 0 or greater, draws a new rulerline at that location. `oldLoc`
// and `newLoc` are expressed in the coordinate system of the NSRulerView, not
// the client or document view, and are x coordinates for horizontal rulers
// and y coordinates for vertical rulers. Use NSView’s `convert...` methods
// to convert coordinates from the client or document view’s coordinate
// system to that of the NSRulerView.
// 
// This method is useful for drawing highlight lines in the ruler to show the
// position or extent of an object while it’s being dragged in the client
// view. The sender is responsible for keeping track of the number and
// positions of temporary lines—the NSRulerView only does the drawing.
//
// See: https://developer.apple.com/documentation/AppKit/NSRulerView/moveRulerline(fromLocation:toLocation:)
func (r NSRulerView) MoveRulerlineFromLocationToLocation(oldLocation float64, newLocation float64) {
	objc.Send[objc.ID](r.ID, objc.Sel("moveRulerlineFromLocation:toLocation:"), oldLocation, newLocation)
}

// Draws the receiver’s hash marks and labels in `aRect`, which is expressed
// in the receiver’s coordinate system.
//
// # Discussion
// 
// This method is invoked by [DrawRect]—you should never need to invoke it
// directly. You can define custom measurement units using the class method
// [RegisterUnitWithNameAbbreviationUnitToPointsConversionFactorStepUpCycleStepDownCycle].
// Override this method if you want to customize the appearance of the hash
// marks themselves.
//
// See: https://developer.apple.com/documentation/AppKit/NSRulerView/drawHashMarksAndLabels(in:)
func (r NSRulerView) DrawHashMarksAndLabelsInRect(rect corefoundation.CGRect) {
	objc.Send[objc.ID](r.ID, objc.Sel("drawHashMarksAndLabelsInRect:"), rect)
}

// Draws the receiver’s markers in `aRect`, which is expressed in the
// receiver’s coordinate system.
//
// # Discussion
// 
// This method is invoked by [DrawRect]; you should never need to invoke it
// directly, but you might want to override it if you want to do something
// different when drawing markers.
//
// See: https://developer.apple.com/documentation/AppKit/NSRulerView/drawMarkers(in:)
func (r NSRulerView) DrawMarkersInRect(rect corefoundation.CGRect) {
	objc.Send[objc.ID](r.ID, objc.Sel("drawMarkersInRect:"), rect)
}

// Forces recalculation of the hash mark spacing for the next time the
// receiver is displayed.
//
// # Discussion
// 
// You should never need to invoke this method directly, but might need to
// override it if you override [DrawHashMarksAndLabelsInRect].
//
// See: https://developer.apple.com/documentation/AppKit/NSRulerView/invalidateHashMarks()
func (r NSRulerView) InvalidateHashMarks() {
	objc.Send[objc.ID](r.ID, objc.Sel("invalidateHashMarks"))
}
func (r NSRulerView) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](r.ID, objc.Sel("encodeWithCoder:"), coder)
}





// Registers a new unit of measurement with the NSRulerView class, making it
// available to all instances of NSRulerView.
//
// # Discussion
// 
// `unitName` is the name of the unit in English, in plural form and
// capitalized by convention—“Inches”, for example. The unit name is
// used as a key to identify the measurement units and so shouldn’t be
// localized. `abbreviation` is a localized short form of the unit name, such
// as “in” for Inches. `conversionFactor` is the number of PostScript
// points in the specified unit; there are 72.0 points per inch, for example.
// `stepUpCycle` and `stepDownCycle` are arrays of NSNumbers that specify how
// hash marks are calculated, as explained in [Setting Up a Ruler View]. All
// numbers in `stepUpCycle` should be greater than 1.0, those in
// `stepDownCycle` should be less than 1.0.
// 
// NSRulerView supports these units by default:
// 
// [Table data omitted]
//
// [Setting Up a Ruler View]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/Rulers/Tasks/SettingUpRulerView.html#//apple_ref/doc/uid/20000874
//
// See: https://developer.apple.com/documentation/AppKit/NSRulerView/registerUnit(withName:abbreviation:unitToPointsConversionFactor:stepUpCycle:stepDownCycle:)
func (_NSRulerViewClass NSRulerViewClass) RegisterUnitWithNameAbbreviationUnitToPointsConversionFactorStepUpCycleStepDownCycle(unitName NSRulerViewUnitName, abbreviation string, conversionFactor float64, stepUpCycle []foundation.NSNumber, stepDownCycle []foundation.NSNumber) {
	objc.Send[objc.ID](objc.ID(_NSRulerViewClass.class), objc.Sel("registerUnitWithName:abbreviation:unitToPointsConversionFactor:stepUpCycle:stepDownCycle:"), objc.String(string(unitName)), objc.String(abbreviation), conversionFactor, objectivec.IObjectSliceToNSArray(stepUpCycle), objectivec.IObjectSliceToNSArray(stepDownCycle))
}








// The measurement units used by the ruler to `unitName`.
//
// # Discussion
// 
// `unitName` must have been registered with the NSRulerView class object
// prior to invoking this method. See the description of the class method
// [RegisterUnitWithNameAbbreviationUnitToPointsConversionFactorStepUpCycleStepDownCycle]
// for a list of predefined units.
//
// See: https://developer.apple.com/documentation/AppKit/NSRulerView/measurementUnits
func (r NSRulerView) MeasurementUnits() NSRulerViewUnitName {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("measurementUnits"))
	return NSRulerViewUnitName(foundation.NSStringFromID(rv).String())
}
func (r NSRulerView) SetMeasurementUnits(value NSRulerViewUnitName) {
	objc.Send[struct{}](r.ID, objc.Sel("setMeasurementUnits:"), objc.String(string(value)))
}



// The receiver’s client view, if it has one.
//
// # Discussion
// 
// `aView` is either the document view of the NSScrollView that contains the
// receiver or a subview of the document view.
//
// See: https://developer.apple.com/documentation/AppKit/NSRulerView/clientView
func (r NSRulerView) ClientView() INSView {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("clientView"))
	return NSViewFromID(objc.ID(rv))
}
func (r NSRulerView) SetClientView(value INSView) {
	objc.Send[struct{}](r.ID, objc.Sel("setClientView:"), value)
}



// The receiver’s accessory view to `aView`.
//
// # Discussion
// 
// Raises an [NSInternalInconsistencyException] if `aView` is not `nil` and
// the receiver has no client view.
//
// See: https://developer.apple.com/documentation/AppKit/NSRulerView/accessoryView
func (r NSRulerView) AccessoryView() INSView {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("accessoryView"))
	return NSViewFromID(objc.ID(rv))
}
func (r NSRulerView) SetAccessoryView(value INSView) {
	objc.Send[struct{}](r.ID, objc.Sel("setAccessoryView:"), value)
}



// The distance to the zero hash mark from the bounds origin of the
// NSScrollView’s document view (not of the receiver’s client view), in
// the document view’s coordinate system.
//
// # Discussion
// 
// The default offset is 0.0, meaning that the ruler origin coincides with the
// bounds origin of the document view.
//
// See: https://developer.apple.com/documentation/AppKit/NSRulerView/originOffset
func (r NSRulerView) OriginOffset() float64 {
	rv := objc.Send[float64](r.ID, objc.Sel("originOffset"))
	return rv
}
func (r NSRulerView) SetOriginOffset(value float64) {
	objc.Send[struct{}](r.ID, objc.Sel("setOriginOffset:"), value)
}



// The receiver’s ruler markers to `markers`, removing any existing ruler
// markers and not consulting with the client view about the new markers.
//
// # Discussion
// 
// `markers` can be `nil` or empty to remove all ruler markers. Raises an
// [NSInternalInconsistencyException] if `markers` is not `nil` and the
// receiver has no client view.
//
// See: https://developer.apple.com/documentation/AppKit/NSRulerView/markers
func (r NSRulerView) Markers() []NSRulerMarker {
	rv := objc.Send[[]objc.ID](r.ID, objc.Sel("markers"))
	return objc.ConvertSlice(rv, func(id objc.ID) NSRulerMarker {
		return NSRulerMarkerFromID(id)
	})
}
func (r NSRulerView) SetMarkers(value []NSRulerMarker) {
	objc.Send[struct{}](r.ID, objc.Sel("setMarkers:"), objectivec.IObjectSliceToNSArray(value))
}



// The NSScrollView that owns the receiver to `scrollView`, without retaining
// it.
//
// # Discussion
// 
// This method is generally invoked only by the ruler’s scroll view; you
// should rarely need to invoke it directly.
//
// See: https://developer.apple.com/documentation/AppKit/NSRulerView/scrollView
func (r NSRulerView) ScrollView() INSScrollView {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("scrollView"))
	return NSScrollViewFromID(objc.ID(rv))
}
func (r NSRulerView) SetScrollView(value INSScrollView) {
	objc.Send[struct{}](r.ID, objc.Sel("setScrollView:"), value)
}



// The orientation of the receiver to `orientation`.
//
// # Discussion
// 
// Possible values for `orientation` are described in [Constants].
//
// See: https://developer.apple.com/documentation/AppKit/NSRulerView/orientation-swift.property
func (r NSRulerView) Orientation() NSRulerOrientation {
	rv := objc.Send[NSRulerOrientation](r.ID, objc.Sel("orientation"))
	return NSRulerOrientation(rv)
}
func (r NSRulerView) SetOrientation(value NSRulerOrientation) {
	objc.Send[struct{}](r.ID, objc.Sel("setOrientation:"), value)
}



// The room available for the receiver’s accessory view to `thickness`.
//
// # Discussion
// 
// If the ruler is horizontal, `thickness` is the height of the accessory
// view; otherwise, it’s the width. NSRulerViews by default reserve no space
// for an accessory view.
// 
// An NSRulerView automatically increases the reserved thickness as necessary
// to that of the accessory view. When the accessory view is thinner than the
// reserved space, it’s centered in that space. If you plan to use several
// accessory views of different sizes, you should set the reserved thickness
// beforehand to that of the thickest accessory view, in order to avoid
// retiling of the NSScrollView.
//
// See: https://developer.apple.com/documentation/AppKit/NSRulerView/reservedThicknessForAccessoryView
func (r NSRulerView) ReservedThicknessForAccessoryView() float64 {
	rv := objc.Send[float64](r.ID, objc.Sel("reservedThicknessForAccessoryView"))
	return rv
}
func (r NSRulerView) SetReservedThicknessForAccessoryView(value float64) {
	objc.Send[struct{}](r.ID, objc.Sel("setReservedThicknessForAccessoryView:"), value)
}



// The room available for ruler markers to `thickness`.
//
// # Discussion
// 
// The default thickness reserved for markers is 15.0 PostScript units for a
// horizontal ruler and 0.0 PostScript units for a vertical ruler (under the
// assumption that vertical rulers rarely contain markers). If you don’t
// expect to have any markers on the ruler, you can set the reserved thickness
// to 0.0.
// 
// An NSRulerView automatically increases the reserved thickness as necessary
// to that of its thickest marker. If you plan to use markers of varying
// sizes, you should set the reserved thickness beforehand to that of the
// thickest one in order to avoid retiling of the NSScrollView.
//
// See: https://developer.apple.com/documentation/AppKit/NSRulerView/reservedThicknessForMarkers
func (r NSRulerView) ReservedThicknessForMarkers() float64 {
	rv := objc.Send[float64](r.ID, objc.Sel("reservedThicknessForMarkers"))
	return rv
}
func (r NSRulerView) SetReservedThicknessForMarkers(value float64) {
	objc.Send[struct{}](r.ID, objc.Sel("setReservedThicknessForMarkers:"), value)
}



// The thickness of the area where ruler hash marks and labels are drawn.
//
// # Discussion
// 
// This value is the height of the ruler area for a horizontal ruler or the
// width of the ruler area for a vertical ruler. Rulers are by default 16.0
// PostScript units thick. You should rarely need to change this layout
// attribute, but subclasses might do so to accommodate custom drawing.
//
// See: https://developer.apple.com/documentation/AppKit/NSRulerView/ruleThickness
func (r NSRulerView) RuleThickness() float64 {
	rv := objc.Send[float64](r.ID, objc.Sel("ruleThickness"))
	return rv
}
func (r NSRulerView) SetRuleThickness(value float64) {
	objc.Send[struct{}](r.ID, objc.Sel("setRuleThickness:"), value)
}



// The thickness needed for proper tiling of the receiver within an
// NSScrollView.
//
// # Discussion
// 
// This thickness is the height of a horizontal ruler and the width of a
// vertical ruler. The required thickness is the sum of the thicknesses of the
// ruler area, the marker area, and the accessory view.
//
// See: https://developer.apple.com/documentation/AppKit/NSRulerView/requiredThickness
func (r NSRulerView) RequiredThickness() float64 {
	rv := objc.Send[float64](r.ID, objc.Sel("requiredThickness"))
	return rv
}



// The location of the receiver’s baseline, in its own coordinate system.
//
// # Discussion
// 
// This is a y position for horizontal rulers and an x position for vertical
// ones.
//
// See: https://developer.apple.com/documentation/AppKit/NSRulerView/baselineLocation
func (r NSRulerView) BaselineLocation() float64 {
	rv := objc.Send[float64](r.ID, objc.Sel("baselineLocation"))
	return rv
}



// A Boolean that indicates if the ruler view’s coordinate system is
// flipped.
//
// # Discussion
// 
// [true] if the receiver’s coordinate system is flipped, [false] otherwise.
// 
// A vertical ruler takes into account whether the coordinate system of the
// [NSScrollView]‘s document view—not the receiver’s client view—is
// flipped. A horizontal ruler is always flipped.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSRulerView/isFlipped
func (r NSRulerView) Flipped() bool {
	rv := objc.Send[bool](r.ID, objc.Sel("isFlipped"))
	return rv
}



// A Boolean that indicates whether the scroll view keeps a horizontal ruler
// object.
//
// See: https://developer.apple.com/documentation/appkit/nsscrollview/hashorizontalruler
func (r NSRulerView) HasHorizontalRuler() bool {
	rv := objc.Send[bool](r.ID, objc.Sel("hasHorizontalRuler"))
	return rv
}
func (r NSRulerView) SetHasHorizontalRuler(value bool) {
	objc.Send[struct{}](r.ID, objc.Sel("setHasHorizontalRuler:"), value)
}



// A Boolean that indicates whether the scroll view keeps a vertical ruler
// object.
//
// See: https://developer.apple.com/documentation/appkit/nsscrollview/hasverticalruler
func (r NSRulerView) HasVerticalRuler() bool {
	rv := objc.Send[bool](r.ID, objc.Sel("hasVerticalRuler"))
	return rv
}
func (r NSRulerView) SetHasVerticalRuler(value bool) {
	objc.Send[struct{}](r.ID, objc.Sel("setHasVerticalRuler:"), value)
}



































