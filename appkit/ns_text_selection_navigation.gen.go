// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSTextSelectionNavigation] class.
var (
	_NSTextSelectionNavigationClass     NSTextSelectionNavigationClass
	_NSTextSelectionNavigationClassOnce sync.Once
)

func getNSTextSelectionNavigationClass() NSTextSelectionNavigationClass {
	_NSTextSelectionNavigationClassOnce.Do(func() {
		_NSTextSelectionNavigationClass = NSTextSelectionNavigationClass{class: objc.GetClass("NSTextSelectionNavigation")}
	})
	return _NSTextSelectionNavigationClass
}

// GetNSTextSelectionNavigationClass returns the class object for NSTextSelectionNavigation.
func GetNSTextSelectionNavigationClass() NSTextSelectionNavigationClass {
	return getNSTextSelectionNavigationClass()
}

type NSTextSelectionNavigationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSTextSelectionNavigationClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSTextSelectionNavigationClass) Alloc() NSTextSelectionNavigation {
	rv := objc.Send[NSTextSelectionNavigation](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// An interface you use to expose methods for obtaining results from actions
// performed on text selections.
//
// # Creating a selection navigation
//
//   - [NSTextSelectionNavigation.InitWithDataSource]: Creates a new object using the text selection data source you provide.
//
// # Selection characteristics
//
//   - [NSTextSelectionNavigation.AllowsNonContiguousRanges]: Determines if the instance could produce selections with multiple noncontiguous selections.
//   - [NSTextSelectionNavigation.SetAllowsNonContiguousRanges]
//   - [NSTextSelectionNavigation.RotatesCoordinateSystemForLayoutOrientation]: Determines if the framework rotates the coordinate system to match the layout orientation.
//   - [NSTextSelectionNavigation.SetRotatesCoordinateSystemForLayoutOrientation]
//   - [NSTextSelectionNavigation.TextSelectionForSelectionGranularityEnclosingPointInContainerAtLocation]: Returns a text selection that expands to the nearest boundaries for selection granularity and an enclosing point you specify.
//
// # Accessing the data source
//
//   - [NSTextSelectionNavigation.TextSelectionDataSource]: The data source associated with this selection navigation.
//
// # Working with text selections
//
//   - [NSTextSelectionNavigation.TextSelectionForSelectionGranularityEnclosingTextSelection]: Returns a text selection expanded to the nearest boundaries for the selection granularity and enclosing text selection text ranges you specify.
//   - [NSTextSelectionNavigation.TextSelectionsInteractingAtPointInContainerAtLocationAnchorsModifiersSelectingBounds]: Returns an array of text selections produced by a tap or click at the point you specify.
//   - [NSTextSelectionNavigation.DestinationSelectionForTextSelectionDirectionDestinationExtendingConfined]: Returns a new selection that results from applying the navigation operations you specify to the text selection you provide.
//
// # Controlling cache behavior
//
//   - [NSTextSelectionNavigation.FlushLayoutCache]: Flushes cached layout information.
//
// # Finding the insertion point
//
//   - [NSTextSelectionNavigation.ResolvedInsertionLocationForTextSelectionWritingDirection]: Returns the location for inserting the next input depending on the state of the current and secondary selections.
//
// # Specifying deletion ranges
//
//   - [NSTextSelectionNavigation.DeletionRangesForTextSelectionDirectionDestinationAllowsDecomposition]: Returns the ranges for deleting the text based on the current selection and movement arguments.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextSelectionNavigation
type NSTextSelectionNavigation struct {
	objectivec.Object
}

// NSTextSelectionNavigationFromID constructs a [NSTextSelectionNavigation] from an objc.ID.
//
// An interface you use to expose methods for obtaining results from actions
// performed on text selections.
func NSTextSelectionNavigationFromID(id objc.ID) NSTextSelectionNavigation {
	return NSTextSelectionNavigation{objectivec.Object{ID: id}}
}
// NOTE: NSTextSelectionNavigation adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSTextSelectionNavigation] class.
//
// # Creating a selection navigation
//
//   - [INSTextSelectionNavigation.InitWithDataSource]: Creates a new object using the text selection data source you provide.
//
// # Selection characteristics
//
//   - [INSTextSelectionNavigation.AllowsNonContiguousRanges]: Determines if the instance could produce selections with multiple noncontiguous selections.
//   - [INSTextSelectionNavigation.SetAllowsNonContiguousRanges]
//   - [INSTextSelectionNavigation.RotatesCoordinateSystemForLayoutOrientation]: Determines if the framework rotates the coordinate system to match the layout orientation.
//   - [INSTextSelectionNavigation.SetRotatesCoordinateSystemForLayoutOrientation]
//   - [INSTextSelectionNavigation.TextSelectionForSelectionGranularityEnclosingPointInContainerAtLocation]: Returns a text selection that expands to the nearest boundaries for selection granularity and an enclosing point you specify.
//
// # Accessing the data source
//
//   - [INSTextSelectionNavigation.TextSelectionDataSource]: The data source associated with this selection navigation.
//
// # Working with text selections
//
//   - [INSTextSelectionNavigation.TextSelectionForSelectionGranularityEnclosingTextSelection]: Returns a text selection expanded to the nearest boundaries for the selection granularity and enclosing text selection text ranges you specify.
//   - [INSTextSelectionNavigation.TextSelectionsInteractingAtPointInContainerAtLocationAnchorsModifiersSelectingBounds]: Returns an array of text selections produced by a tap or click at the point you specify.
//   - [INSTextSelectionNavigation.DestinationSelectionForTextSelectionDirectionDestinationExtendingConfined]: Returns a new selection that results from applying the navigation operations you specify to the text selection you provide.
//
// # Controlling cache behavior
//
//   - [INSTextSelectionNavigation.FlushLayoutCache]: Flushes cached layout information.
//
// # Finding the insertion point
//
//   - [INSTextSelectionNavigation.ResolvedInsertionLocationForTextSelectionWritingDirection]: Returns the location for inserting the next input depending on the state of the current and secondary selections.
//
// # Specifying deletion ranges
//
//   - [INSTextSelectionNavigation.DeletionRangesForTextSelectionDirectionDestinationAllowsDecomposition]: Returns the ranges for deleting the text based on the current selection and movement arguments.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextSelectionNavigation
type INSTextSelectionNavigation interface {
	objectivec.IObject

	// Topic: Creating a selection navigation

	// Creates a new object using the text selection data source you provide.
	InitWithDataSource(dataSource NSTextSelectionDataSource) NSTextSelectionNavigation

	// Topic: Selection characteristics

	// Determines if the instance could produce selections with multiple noncontiguous selections.
	AllowsNonContiguousRanges() bool
	SetAllowsNonContiguousRanges(value bool)
	// Determines if the framework rotates the coordinate system to match the layout orientation.
	RotatesCoordinateSystemForLayoutOrientation() bool
	SetRotatesCoordinateSystemForLayoutOrientation(value bool)
	// Returns a text selection that expands to the nearest boundaries for selection granularity and an enclosing point you specify.
	TextSelectionForSelectionGranularityEnclosingPointInContainerAtLocation(selectionGranularity NSTextSelectionGranularity, point corefoundation.CGPoint, location NSTextLocation) INSTextSelection

	// Topic: Accessing the data source

	// The data source associated with this selection navigation.
	TextSelectionDataSource() NSTextSelectionDataSource

	// Topic: Working with text selections

	// Returns a text selection expanded to the nearest boundaries for the selection granularity and enclosing text selection text ranges you specify.
	TextSelectionForSelectionGranularityEnclosingTextSelection(selectionGranularity NSTextSelectionGranularity, textSelection INSTextSelection) INSTextSelection
	// Returns an array of text selections produced by a tap or click at the point you specify.
	TextSelectionsInteractingAtPointInContainerAtLocationAnchorsModifiersSelectingBounds(point corefoundation.CGPoint, containerLocation NSTextLocation, anchors []NSTextSelection, modifiers NSTextSelectionNavigationModifier, selecting bool, bounds corefoundation.CGRect) []NSTextSelection
	// Returns a new selection that results from applying the navigation operations you specify to the text selection you provide.
	DestinationSelectionForTextSelectionDirectionDestinationExtendingConfined(textSelection INSTextSelection, direction NSTextSelectionNavigationDirection, destination NSTextSelectionNavigationDestination, extending bool, confined bool) INSTextSelection

	// Topic: Controlling cache behavior

	// Flushes cached layout information.
	FlushLayoutCache()

	// Topic: Finding the insertion point

	// Returns the location for inserting the next input depending on the state of the current and secondary selections.
	ResolvedInsertionLocationForTextSelectionWritingDirection(textSelection INSTextSelection, writingDirection NSTextSelectionNavigationWritingDirection) NSTextLocation

	// Topic: Specifying deletion ranges

	// Returns the ranges for deleting the text based on the current selection and movement arguments.
	DeletionRangesForTextSelectionDirectionDestinationAllowsDecomposition(textSelection INSTextSelection, direction NSTextSelectionNavigationDirection, destination NSTextSelectionNavigationDestination, allowsDecomposition bool) []NSTextRange
}

// Init initializes the instance.
func (t NSTextSelectionNavigation) Init() NSTextSelectionNavigation {
	rv := objc.Send[NSTextSelectionNavigation](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t NSTextSelectionNavigation) Autorelease() NSTextSelectionNavigation {
	rv := objc.Send[NSTextSelectionNavigation](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSTextSelectionNavigation creates a new NSTextSelectionNavigation instance.
func NewNSTextSelectionNavigation() NSTextSelectionNavigation {
	class := getNSTextSelectionNavigationClass()
	rv := objc.Send[NSTextSelectionNavigation](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a new object using the text selection data source you provide.
//
// dataSource: An [NSTextSelectionDataSource].
//
// See: https://developer.apple.com/documentation/AppKit/NSTextSelectionNavigation/init(dataSource:)
func NewTextSelectionNavigationWithDataSource(dataSource NSTextSelectionDataSource) NSTextSelectionNavigation {
	instance := getNSTextSelectionNavigationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDataSource:"), dataSource)
	return NSTextSelectionNavigationFromID(rv)
}

// Creates a new object using the text selection data source you provide.
//
// dataSource: An [NSTextSelectionDataSource].
//
// See: https://developer.apple.com/documentation/AppKit/NSTextSelectionNavigation/init(dataSource:)
func (t NSTextSelectionNavigation) InitWithDataSource(dataSource NSTextSelectionDataSource) NSTextSelectionNavigation {
	rv := objc.Send[NSTextSelectionNavigation](t.ID, objc.Sel("initWithDataSource:"), dataSource)
	return rv
}
// Returns a text selection that expands to the nearest boundaries for
// selection granularity and an enclosing point you specify.
//
// selectionGranularity: One of the available [NSTextSelection.Granularity] options.
// //
// [NSTextSelection.Granularity]: https://developer.apple.com/documentation/AppKit/NSTextSelection/Granularity-swift.enum
//
// point: The point that encloses the text.
//
// location: An [NSTextLocation] that describes the container.
//
// # Return Value
// 
// A new [NSTextSelection], or `nil` if the text selection is not found.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextSelectionNavigation/textSelection(for:enclosing:inContainerAt:)
func (t NSTextSelectionNavigation) TextSelectionForSelectionGranularityEnclosingPointInContainerAtLocation(selectionGranularity NSTextSelectionGranularity, point corefoundation.CGPoint, location NSTextLocation) INSTextSelection {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("textSelectionForSelectionGranularity:enclosingPoint:inContainerAtLocation:"), selectionGranularity, point, location)
	return NSTextSelectionFromID(rv)
}
// Returns a text selection expanded to the nearest boundaries for the
// selection granularity and enclosing text selection text ranges you specify.
//
// selectionGranularity: One of the available [NSTextSelection.Granularity] options.
// //
// [NSTextSelection.Granularity]: https://developer.apple.com/documentation/AppKit/NSTextSelection/Granularity-swift.enum
//
// textSelection: The text selection that describes the text range of interest.
//
// # Return Value
// 
// A new text selection.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextSelectionNavigation/textSelection(for:enclosing:)
func (t NSTextSelectionNavigation) TextSelectionForSelectionGranularityEnclosingTextSelection(selectionGranularity NSTextSelectionGranularity, textSelection INSTextSelection) INSTextSelection {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("textSelectionForSelectionGranularity:enclosingTextSelection:"), selectionGranularity, textSelection)
	return NSTextSelectionFromID(rv)
}
// Returns an array of text selections produced by a tap or click at the point
// you specify.
//
// point: A [CGPoint] that represents the location of the tap or click.
//
// containerLocation: A `NSTextLocation that describes the contasiner location`.
//
// anchors: An array of [NSTextSelection] objects.
//
// modifiers: One or more [NSTextSelectionNavigation.Modifier] options.
// //
// [NSTextSelectionNavigation.Modifier]: https://developer.apple.com/documentation/AppKit/NSTextSelectionNavigation/Modifier
//
// selecting: A Boolean value that indicates if the selection is in drag session.
//
// bounds: A [CGRect] that defines the view area in the container’s coordinate
// system that can interact with events.
//
// # Return Value
// 
// An array of text selections.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextSelectionNavigation/textSelections(interactingAt:inContainerAt:anchors:modifiers:selecting:bounds:)
func (t NSTextSelectionNavigation) TextSelectionsInteractingAtPointInContainerAtLocationAnchorsModifiersSelectingBounds(point corefoundation.CGPoint, containerLocation NSTextLocation, anchors []NSTextSelection, modifiers NSTextSelectionNavigationModifier, selecting bool, bounds corefoundation.CGRect) []NSTextSelection {
	rv := objc.Send[[]objc.ID](t.ID, objc.Sel("textSelectionsInteractingAtPoint:inContainerAtLocation:anchors:modifiers:selecting:bounds:"), point, containerLocation, objectivec.IObjectSliceToNSArray(anchors), modifiers, selecting, bounds)
	return objc.ConvertSlice(rv, func(id objc.ID) NSTextSelection {
		return NSTextSelectionFromID(id)
	})
}
// Returns a new selection that results from applying the navigation
// operations you specify to the text selection you provide.
//
// textSelection: The source selection.
//
// direction: One of the available [NSTextSelectionNavigation.Direction] directions.
// //
// [NSTextSelectionNavigation.Direction]: https://developer.apple.com/documentation/AppKit/NSTextSelectionNavigation/Direction
//
// destination: One of the available [NSTextSelectionNavigation.Destination] destinations.
// //
// [NSTextSelectionNavigation.Destination]: https://developer.apple.com/documentation/AppKit/NSTextSelectionNavigation/Destination
//
// extending: Whether this selection extends an existing selection.
//
// confined: Whether to confine movement to the existing selection.
//
// # Return Value
// 
// A new [NSTextSelection], or `nil` if the operation doesn’t produce a
// logically valid result.
//
// # Discussion
// 
// If `confined` is `true`, confine any movement to the text element that the
// selection already lies within.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextSelectionNavigation/destinationSelection(for:direction:destination:extending:confined:)
func (t NSTextSelectionNavigation) DestinationSelectionForTextSelectionDirectionDestinationExtendingConfined(textSelection INSTextSelection, direction NSTextSelectionNavigationDirection, destination NSTextSelectionNavigationDestination, extending bool, confined bool) INSTextSelection {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("destinationSelectionForTextSelection:direction:destination:extending:confined:"), textSelection, direction, destination, extending, confined)
	return NSTextSelectionFromID(rv)
}
// Flushes cached layout information.
//
// # Discussion
// 
// You call this method whenever the contents of the document change.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextSelectionNavigation/flushLayoutCache()
func (t NSTextSelectionNavigation) FlushLayoutCache() {
	objc.Send[objc.ID](t.ID, objc.Sel("flushLayoutCache"))
}
// Returns the location for inserting the next input depending on the state of
// the current and secondary selections.
//
// textSelection: The text selection.
//
// writingDirection: The [NSTextSelectionNavigation.WritingDirection] direction.
// //
// [NSTextSelectionNavigation.WritingDirection]: https://developer.apple.com/documentation/AppKit/NSTextSelectionNavigation/WritingDirection
//
// # Return Value
// 
// Returns an [NSTextLocation] when the `textSelection.IsLogical() = false
// AND` `secondarySelectionLocation != nil`. Otherwise, returns nil.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextSelectionNavigation/resolvedInsertionLocation(for:writingDirection:)
func (t NSTextSelectionNavigation) ResolvedInsertionLocationForTextSelectionWritingDirection(textSelection INSTextSelection, writingDirection NSTextSelectionNavigationWritingDirection) NSTextLocation {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("resolvedInsertionLocationForTextSelection:writingDirection:"), textSelection, writingDirection)
	return NSTextLocationObjectFromID(rv)
}
// Returns the ranges for deleting the text based on the current selection and
// movement arguments.
//
// textSelection: The text selection.
//
// direction: The [NSTextSelectionNavigation.Direction] to consider when calculating the
// deletion ranges.
// //
// [NSTextSelectionNavigation.Direction]: https://developer.apple.com/documentation/AppKit/NSTextSelectionNavigation/Direction
//
// destination: The [NSTextSelectionNavigation.Destination] that describes the scope of the
// text selection to consider when calculating the deletion ranges.
// //
// [NSTextSelectionNavigation.Destination]: https://developer.apple.com/documentation/AppKit/NSTextSelectionNavigation/Destination
//
// allowsDecomposition: A Boolean value that determines if this method operates on composite
// characters which may be present depending on the characteristics of the
// script used by `textSelection`.
//
// # Return Value
// 
// An array of text ranges to delete.
//
// # Discussion
// 
// The selection after deletion contains a zero-length range starting at the
// location of the first range returned. The framework ignores the destination
// when `textSelection` has a non-empty selection. The `allowsDecomposition`
// parameter only applies to the
// [NSTextSelectionNavigation.Direction.backward] direction and
// [NSTextSelectionNavigation.Destination.character] with a zero-length
// selection.
//
// [NSTextSelectionNavigation.Destination.character]: https://developer.apple.com/documentation/AppKit/NSTextSelectionNavigation/Destination/character
// [NSTextSelectionNavigation.Direction.backward]: https://developer.apple.com/documentation/AppKit/NSTextSelectionNavigation/Direction/backward
//
// See: https://developer.apple.com/documentation/AppKit/NSTextSelectionNavigation/deletionRanges(for:direction:destination:allowsDecomposition:)
func (t NSTextSelectionNavigation) DeletionRangesForTextSelectionDirectionDestinationAllowsDecomposition(textSelection INSTextSelection, direction NSTextSelectionNavigationDirection, destination NSTextSelectionNavigationDestination, allowsDecomposition bool) []NSTextRange {
	rv := objc.Send[[]objc.ID](t.ID, objc.Sel("deletionRangesForTextSelection:direction:destination:allowsDecomposition:"), textSelection, direction, destination, allowsDecomposition)
	return objc.ConvertSlice(rv, func(id objc.ID) NSTextRange {
		return NSTextRangeFromID(id)
	})
}

// Determines if the instance could produce selections with multiple
// noncontiguous selections.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextSelectionNavigation/allowsNonContiguousRanges
func (t NSTextSelectionNavigation) AllowsNonContiguousRanges() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("allowsNonContiguousRanges"))
	return rv
}
func (t NSTextSelectionNavigation) SetAllowsNonContiguousRanges(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setAllowsNonContiguousRanges:"), value)
}
// Determines if the framework rotates the coordinate system to match the
// layout orientation.
//
// # Discussion
// 
// If set to `true`, the framework rotates the coordinate system for arguments
// passed to the navigation methods such as
// [TextSelectionsInteractingAtPointInContainerAtLocationAnchorsModifiersSelectingBounds]:
// based on the text container layout orientation. Defaults to `false`.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextSelectionNavigation/rotatesCoordinateSystemForLayoutOrientation
func (t NSTextSelectionNavigation) RotatesCoordinateSystemForLayoutOrientation() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("rotatesCoordinateSystemForLayoutOrientation"))
	return rv
}
func (t NSTextSelectionNavigation) SetRotatesCoordinateSystemForLayoutOrientation(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setRotatesCoordinateSystemForLayoutOrientation:"), value)
}
// The data source associated with this selection navigation.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextSelectionNavigation/textSelectionDataSource
func (t NSTextSelectionNavigation) TextSelectionDataSource() NSTextSelectionDataSource {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("textSelectionDataSource"))
	return NSTextSelectionDataSourceObjectFromID(rv)
}

