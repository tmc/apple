// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSTextContainer] class.
var (
	_NSTextContainerClass     NSTextContainerClass
	_NSTextContainerClassOnce sync.Once
)

func getNSTextContainerClass() NSTextContainerClass {
	_NSTextContainerClassOnce.Do(func() {
		_NSTextContainerClass = NSTextContainerClass{class: objc.GetClass("NSTextContainer")}
	})
	return _NSTextContainerClass
}

// GetNSTextContainerClass returns the class object for NSTextContainer.
func GetNSTextContainerClass() NSTextContainerClass {
	return getNSTextContainerClass()
}

type NSTextContainerClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSTextContainerClass) Alloc() NSTextContainer {
	rv := objc.Send[NSTextContainer](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A region where text layout occurs.
//
// # Overview
// 
// An [NSLayoutManager] uses [NSTextContainer] to determine where to break
// lines, lay out portions of text, and so on. An [NSTextContainer] object
// typically defines rectangular regions, but you can define exclusion paths
// inside the text container to create regions where text doesn’t flow. You
// can also subclass to create text containers with nonrectangular regions,
// such as circular regions, regions with holes in them, or regions that flow
// alongside graphics.
// 
// You can access instances of the [NSTextContainer], [NSLayoutManager], and
// [NSTextStorage] classes from threads other than the main thread as long as
// the app guarantees access from only one thread at a time.
//
// # Creating a text container
//
//   - [NSTextContainer.InitWithSize]: Initializes a text container with a specified bounding rectangle.
//   - [NSTextContainer.InitWithCoder]: Creates a text container from data in an unarchiver.
//
// # Managing text components
//
//   - [NSTextContainer.LayoutManager]: The text container’s layout manager.
//   - [NSTextContainer.SetLayoutManager]
//   - [NSTextContainer.TextLayoutManager]
//   - [NSTextContainer.ReplaceLayoutManager]: Replaces the layout manager for the group of text system objects that contains the text container.
//   - [NSTextContainer.TextView]: The text container’s text view.
//   - [NSTextContainer.SetTextView]
//
// # Defining the container shape
//
//   - [NSTextContainer.Size]: The size of the text container’s bounding rectangle.
//   - [NSTextContainer.SetSize]
//   - [NSTextContainer.ExclusionPaths]: An array of path objects that represents the regions where text doesn’t display in the text container.
//   - [NSTextContainer.SetExclusionPaths]
//   - [NSTextContainer.LineBreakMode]: The behavior of the last line inside the text container.
//   - [NSTextContainer.SetLineBreakMode]
//   - [NSTextContainer.WidthTracksTextView]: A Boolean that controls whether the text container adjusts the width of its bounding rectangle when its text view resizes.
//   - [NSTextContainer.SetWidthTracksTextView]
//   - [NSTextContainer.HeightTracksTextView]: A Boolean that controls whether the text container adjusts the height of its bounding rectangle when its text view resizes.
//   - [NSTextContainer.SetHeightTracksTextView]
//
// # Constraining text layout
//
//   - [NSTextContainer.MaximumNumberOfLines]: The maximum number of lines that the text container can store.
//   - [NSTextContainer.SetMaximumNumberOfLines]
//   - [NSTextContainer.LineFragmentPadding]: The value for the text inset within line fragment rectangles.
//   - [NSTextContainer.SetLineFragmentPadding]
//   - [NSTextContainer.LineFragmentRectForProposedRectAtIndexWritingDirectionRemainingRect]: Returns the bounds of a line fragment rectangle inside the text container for the proposed rectangle.
//   - [NSTextContainer.SimpleRectangularTextContainer]: A Boolean that indicates whether the text container’s region is a rectangle with no holes or gaps, and whose edges are parallel to the text view’s coordinate system axes.
//
// # Deprecated
//
//   - [NSTextContainer.InitWithContainerSize]: Initializes a text container with a specified bounding rectangle.
//   - [NSTextContainer.LineFragmentRectForProposedRectSweepDirectionMovementDirectionRemainingRect]: Calculates and returns the longest rectangle available in the proposed rectangle for displaying text.
//   - [NSTextContainer.ContainerSize]: The size of the text container’s bounding rectangle.
//   - [NSTextContainer.SetContainerSize]
//
// See: https://developer.apple.com/documentation/AppKit/NSTextContainer
type NSTextContainer struct {
	objectivec.Object
}

// NSTextContainerFromID constructs a [NSTextContainer] from an objc.ID.
//
// A region where text layout occurs.
func NSTextContainerFromID(id objc.ID) NSTextContainer {
	return NSTextContainer{objectivec.Object{ID: id}}
}
// NOTE: NSTextContainer adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSTextContainer] class.
//
// # Creating a text container
//
//   - [INSTextContainer.InitWithSize]: Initializes a text container with a specified bounding rectangle.
//   - [INSTextContainer.InitWithCoder]: Creates a text container from data in an unarchiver.
//
// # Managing text components
//
//   - [INSTextContainer.LayoutManager]: The text container’s layout manager.
//   - [INSTextContainer.SetLayoutManager]
//   - [INSTextContainer.TextLayoutManager]
//   - [INSTextContainer.ReplaceLayoutManager]: Replaces the layout manager for the group of text system objects that contains the text container.
//   - [INSTextContainer.TextView]: The text container’s text view.
//   - [INSTextContainer.SetTextView]
//
// # Defining the container shape
//
//   - [INSTextContainer.Size]: The size of the text container’s bounding rectangle.
//   - [INSTextContainer.SetSize]
//   - [INSTextContainer.ExclusionPaths]: An array of path objects that represents the regions where text doesn’t display in the text container.
//   - [INSTextContainer.SetExclusionPaths]
//   - [INSTextContainer.LineBreakMode]: The behavior of the last line inside the text container.
//   - [INSTextContainer.SetLineBreakMode]
//   - [INSTextContainer.WidthTracksTextView]: A Boolean that controls whether the text container adjusts the width of its bounding rectangle when its text view resizes.
//   - [INSTextContainer.SetWidthTracksTextView]
//   - [INSTextContainer.HeightTracksTextView]: A Boolean that controls whether the text container adjusts the height of its bounding rectangle when its text view resizes.
//   - [INSTextContainer.SetHeightTracksTextView]
//
// # Constraining text layout
//
//   - [INSTextContainer.MaximumNumberOfLines]: The maximum number of lines that the text container can store.
//   - [INSTextContainer.SetMaximumNumberOfLines]
//   - [INSTextContainer.LineFragmentPadding]: The value for the text inset within line fragment rectangles.
//   - [INSTextContainer.SetLineFragmentPadding]
//   - [INSTextContainer.LineFragmentRectForProposedRectAtIndexWritingDirectionRemainingRect]: Returns the bounds of a line fragment rectangle inside the text container for the proposed rectangle.
//   - [INSTextContainer.SimpleRectangularTextContainer]: A Boolean that indicates whether the text container’s region is a rectangle with no holes or gaps, and whose edges are parallel to the text view’s coordinate system axes.
//
// # Deprecated
//
//   - [INSTextContainer.InitWithContainerSize]: Initializes a text container with a specified bounding rectangle.
//   - [INSTextContainer.LineFragmentRectForProposedRectSweepDirectionMovementDirectionRemainingRect]: Calculates and returns the longest rectangle available in the proposed rectangle for displaying text.
//   - [INSTextContainer.ContainerSize]: The size of the text container’s bounding rectangle.
//   - [INSTextContainer.SetContainerSize]
//
// See: https://developer.apple.com/documentation/AppKit/NSTextContainer
type INSTextContainer interface {
	objectivec.IObject
	NSTextLayoutOrientationProvider

	// Topic: Creating a text container

	// Initializes a text container with a specified bounding rectangle.
	InitWithSize(size corefoundation.CGSize) NSTextContainer
	// Creates a text container from data in an unarchiver.
	InitWithCoder(coder foundation.INSCoder) NSTextContainer

	// Topic: Managing text components

	// The text container’s layout manager.
	LayoutManager() INSLayoutManager
	SetLayoutManager(value INSLayoutManager)
	TextLayoutManager() INSTextLayoutManager
	// Replaces the layout manager for the group of text system objects that contains the text container.
	ReplaceLayoutManager(newLayoutManager INSLayoutManager)
	// The text container’s text view.
	TextView() INSTextView
	SetTextView(value INSTextView)

	// Topic: Defining the container shape

	// The size of the text container’s bounding rectangle.
	Size() corefoundation.CGSize
	SetSize(value corefoundation.CGSize)
	// An array of path objects that represents the regions where text doesn’t display in the text container.
	ExclusionPaths() []NSBezierPath
	SetExclusionPaths(value []NSBezierPath)
	// The behavior of the last line inside the text container.
	LineBreakMode() NSLineBreakMode
	SetLineBreakMode(value NSLineBreakMode)
	// A Boolean that controls whether the text container adjusts the width of its bounding rectangle when its text view resizes.
	WidthTracksTextView() bool
	SetWidthTracksTextView(value bool)
	// A Boolean that controls whether the text container adjusts the height of its bounding rectangle when its text view resizes.
	HeightTracksTextView() bool
	SetHeightTracksTextView(value bool)

	// Topic: Constraining text layout

	// The maximum number of lines that the text container can store.
	MaximumNumberOfLines() uint
	SetMaximumNumberOfLines(value uint)
	// The value for the text inset within line fragment rectangles.
	LineFragmentPadding() float64
	SetLineFragmentPadding(value float64)
	// Returns the bounds of a line fragment rectangle inside the text container for the proposed rectangle.
	LineFragmentRectForProposedRectAtIndexWritingDirectionRemainingRect(proposedRect corefoundation.CGRect, characterIndex uint, baseWritingDirection NSWritingDirection, remainingRect *corefoundation.CGRect) corefoundation.CGRect
	// A Boolean that indicates whether the text container’s region is a rectangle with no holes or gaps, and whose edges are parallel to the text view’s coordinate system axes.
	SimpleRectangularTextContainer() bool

	// Topic: Deprecated

	// Initializes a text container with a specified bounding rectangle.
	InitWithContainerSize(aContainerSize corefoundation.CGSize) NSTextContainer
	// Calculates and returns the longest rectangle available in the proposed rectangle for displaying text.
	LineFragmentRectForProposedRectSweepDirectionMovementDirectionRemainingRect(proposedRect corefoundation.CGRect, sweepDirection NSLineSweepDirection, movementDirection NSLineMovementDirection, remainingRect foundation.NSRect) corefoundation.CGRect
	// The size of the text container’s bounding rectangle.
	ContainerSize() corefoundation.CGSize
	SetContainerSize(value corefoundation.CGSize)

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (t NSTextContainer) Init() NSTextContainer {
	rv := objc.Send[NSTextContainer](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t NSTextContainer) Autorelease() NSTextContainer {
	rv := objc.Send[NSTextContainer](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSTextContainer creates a new NSTextContainer instance.
func NewNSTextContainer() NSTextContainer {
	class := getNSTextContainerClass()
	rv := objc.Send[NSTextContainer](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a text container from data in an unarchiver.
//
// coder: A coder that implements [NSCoder].
// //
// [NSCoder]: https://developer.apple.com/documentation/Foundation/NSCoder
//
// See: https://developer.apple.com/documentation/AppKit/NSTextContainer/init(coder:)
func NewTextContainerWithCoder(coder foundation.INSCoder) NSTextContainer {
	instance := getNSTextContainerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSTextContainerFromID(rv)
}

// Initializes a text container with a specified bounding rectangle.
//
// aContainerSize: The size of the text container’s bounding rectangle.
//
// # Return Value
// 
// The newly initialized text container.
//
// # Discussion
// 
// The new text container must be added to an [NSLayoutManager] object before
// it can be used. The text container must also have an [NSTextView] object
// set for text to be displayed. This method is the designated initializer for
// the [NSTextContainer] class.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextContainer/init(containerSize:)
func NewTextContainerWithContainerSize(aContainerSize corefoundation.CGSize) NSTextContainer {
	instance := getNSTextContainerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithContainerSize:"), aContainerSize)
	return NSTextContainerFromID(rv)
}

// Initializes a text container with a specified bounding rectangle.
//
// size: The size of the text container’s bounding rectangle.
//
// # Discussion
// 
// The new text container must be added to an [NSLayoutManager] object before
// it can be used. The text container must also have an associated
// [NSTextView] object for text to be displayed. This method is the designated
// initializer for the [NSTextContainer] class.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextContainer/init(size:)
func NewTextContainerWithSize(size corefoundation.CGSize) NSTextContainer {
	instance := getNSTextContainerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSize:"), size)
	return NSTextContainerFromID(rv)
}

// Initializes a text container with a specified bounding rectangle.
//
// size: The size of the text container’s bounding rectangle.
//
// # Discussion
// 
// The new text container must be added to an [NSLayoutManager] object before
// it can be used. The text container must also have an associated
// [NSTextView] object for text to be displayed. This method is the designated
// initializer for the [NSTextContainer] class.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextContainer/init(size:)
func (t NSTextContainer) InitWithSize(size corefoundation.CGSize) NSTextContainer {
	rv := objc.Send[NSTextContainer](t.ID, objc.Sel("initWithSize:"), size)
	return rv
}
// Creates a text container from data in an unarchiver.
//
// coder: A coder that implements [NSCoder].
// //
// [NSCoder]: https://developer.apple.com/documentation/Foundation/NSCoder
//
// See: https://developer.apple.com/documentation/AppKit/NSTextContainer/init(coder:)
func (t NSTextContainer) InitWithCoder(coder foundation.INSCoder) NSTextContainer {
	rv := objc.Send[NSTextContainer](t.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}
// Replaces the layout manager for the group of text system objects that
// contains the text container.
//
// newLayoutManager: The new layout manager.
//
// # Discussion
// 
// The framework reassigns all text containers and text views attached to the
// old layout manager to the new layout manager. Unlike setting the
// [LayoutManager] property directly, this method makes all the adjustments
// necessary to keep the text object relationships intact.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextContainer/replaceLayoutManager(_:)
func (t NSTextContainer) ReplaceLayoutManager(newLayoutManager INSLayoutManager) {
	objc.Send[objc.ID](t.ID, objc.Sel("replaceLayoutManager:"), newLayoutManager)
}
// Returns the bounds of a line fragment rectangle inside the text container
// for the proposed rectangle.
//
// proposedRect: A rectangle in which to lay out text proposed by the layout manager.
//
// characterIndex: The character location inside the text storage for the line fragment being
// processed.
//
// baseWritingDirection: The direction of advancement for line fragments inside a visual horizontal
// line. The values passed into the method are either
// [WritingDirectionLeftToRight] or [WritingDirectionRightToLeft].
//
// remainingRect: The remainder of the proposed rectangle that was excluded from returned
// rectangle. It can be passed in as the proposed rectangle for the next
// iteration.
//
// # Discussion
// 
// The bounds of the line fragment rectangle are determined by the
// intersection of `proposedRect` and the text container’s bounding
// rectangle defined by its [NSTextContainer] property. The regions defined by
// the [NSTextContainer] property are excluded from the return value. It is
// possible that `proposedRect` can be divided into multiple line fragments
// due to exclusion paths. In that case, `remainingRect` returns the remainder
// that can be passed in as the proposed rectangle for the next iteration.
// 
// This method can be overridden by subclasses for further text container
// region customization.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextContainer/lineFragmentRect(forProposedRect:at:writingDirection:remaining:)
func (t NSTextContainer) LineFragmentRectForProposedRectAtIndexWritingDirectionRemainingRect(proposedRect corefoundation.CGRect, characterIndex uint, baseWritingDirection NSWritingDirection, remainingRect *corefoundation.CGRect) corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](t.ID, objc.Sel("lineFragmentRectForProposedRect:atIndex:writingDirection:remainingRect:"), proposedRect, characterIndex, baseWritingDirection, remainingRect)
	return corefoundation.CGRect(rv)
}
// Initializes a text container with a specified bounding rectangle.
//
// aContainerSize: The size of the text container’s bounding rectangle.
//
// # Return Value
// 
// The newly initialized text container.
//
// # Discussion
// 
// The new text container must be added to an [NSLayoutManager] object before
// it can be used. The text container must also have an [NSTextView] object
// set for text to be displayed. This method is the designated initializer for
// the [NSTextContainer] class.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextContainer/init(containerSize:)
func (t NSTextContainer) InitWithContainerSize(aContainerSize corefoundation.CGSize) NSTextContainer {
	rv := objc.Send[NSTextContainer](t.ID, objc.Sel("initWithContainerSize:"), aContainerSize)
	return rv
}
// Calculates and returns the longest rectangle available in the proposed
// rectangle for displaying text.
//
// proposedRect: The proposed rectangle in which to layout text.
//
// sweepDirection: The line sweep direction.
//
// movementDirection: The line movement direction.
//
// remainingRect: Upon return, the unused, possibly shifted, portion of `proposedRect`
// that’s available for further text, or [NSZeroRect] if there is no
// remainder.
//
// # Return Value
// 
// The longest rectangle available in the proposed rectangle for displaying
// text, or [NSZeroRect] if there is none according to the receiver’s region
// definition.
//
// # Discussion
// 
// There is no guarantee as to the width of the proposed rectangle or to its
// location. For example, the proposed rectangle is likely to be much wider
// than the width of the receiver. The receiver should examine `proposedRect`
// to see that it intersects its bounding rectangle and should return a
// modified rectangle based on `sweepDirection` and `movementDirection`, whose
// possible values are listed in the class description. If `sweepDirection` is
// [NSLineSweepRight], for example, the receiver uses this information to trim
// the right end of `proposedRect` as needed rather than the left end.
// 
// If `proposedRect` doesn’t completely overlap the region along the axis of
// `movementDirection` and `movementDirection` isn’t [NSLineDoesntMove],
// this method can either shift the rectangle in that direction as much as
// needed so that it does completely overlap, or return [NSZeroRect] to
// indicate that the proposed rectangle simply doesn’t fit.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextContainer/lineFragmentRect(forProposedRect:sweepDirection:movementDirection:remaining:)
func (t NSTextContainer) LineFragmentRectForProposedRectSweepDirectionMovementDirectionRemainingRect(proposedRect corefoundation.CGRect, sweepDirection NSLineSweepDirection, movementDirection NSLineMovementDirection, remainingRect foundation.NSRect) corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](t.ID, objc.Sel("lineFragmentRectForProposedRect:sweepDirection:movementDirection:remainingRect:"), proposedRect, sweepDirection, movementDirection, remainingRect)
	return corefoundation.CGRect(rv)
}
func (t NSTextContainer) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](t.ID, objc.Sel("encodeWithCoder:"), coder)
}

// The text container’s layout manager.
//
// # Discussion
// 
// Avoid assigning a layout manager directly through this property. Instead,
// use the [ReplaceLayoutManager] method when you want to replace the layout
// manager. The framework sets the value of this property automatically when
// you add a text container to your layout manager using the
// [AddTextContainer] method.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextContainer/layoutManager
func (t NSTextContainer) LayoutManager() INSLayoutManager {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("layoutManager"))
	return NSLayoutManagerFromID(objc.ID(rv))
}
func (t NSTextContainer) SetLayoutManager(value INSLayoutManager) {
	objc.Send[struct{}](t.ID, objc.Sel("setLayoutManager:"), value)
}
// See: https://developer.apple.com/documentation/AppKit/NSTextContainer/textLayoutManager
func (t NSTextContainer) TextLayoutManager() INSTextLayoutManager {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("textLayoutManager"))
	return NSTextLayoutManagerFromID(objc.ID(rv))
}
// The text container’s text view.
//
// # Discussion
// 
// A text container doesn’t need a text view to calculate line fragment
// rectangles, but must have one to display text.
// 
// You can use this property to disconnect a text view from a group of text
// system objects by sending this message to its text container and passing
// `nil` as `aTextView`.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextContainer/textView
func (t NSTextContainer) TextView() INSTextView {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("textView"))
	return NSTextViewFromID(objc.ID(rv))
}
func (t NSTextContainer) SetTextView(value INSTextView) {
	objc.Send[struct{}](t.ID, objc.Sel("setTextView:"), value)
}
// The size of the text container’s bounding rectangle.
//
// # Discussion
// 
// This property defines the maximum size for the layout area returned from
// [LineFragmentRectForProposedRectAtIndexWritingDirectionRemainingRect]. A
// value of `0.0` or less means no limitation.
// 
// If you don’t specify an explicit size when you initialize a text
// container, the system uses a default large size of (`10000000.0`,
// `10000000.0`).
//
// See: https://developer.apple.com/documentation/AppKit/NSTextContainer/size
func (t NSTextContainer) Size() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](t.ID, objc.Sel("size"))
	return corefoundation.CGSize(rv)
}
func (t NSTextContainer) SetSize(value corefoundation.CGSize) {
	objc.Send[struct{}](t.ID, objc.Sel("setSize:"), value)
}
// An array of path objects that represents the regions where text doesn’t
// display in the text container.
//
// # Discussion
// 
// The default value of this property is an empty array. Depending on the
// platform, you can assign an array of [NSBezierPath] or [UIBezierPath]
// objects to exclude text from one or more regions in the text container’s
// bounds. When the layout manager proposes a line fragment rectangle
// intersecting one of the regions defined by the exclusion paths, the text
// container returns an adjusted line fragment rectangle excluding that
// region.
//
// [UIBezierPath]: https://developer.apple.com/documentation/UIKit/UIBezierPath
//
// See: https://developer.apple.com/documentation/AppKit/NSTextContainer/exclusionPaths
func (t NSTextContainer) ExclusionPaths() []NSBezierPath {
	rv := objc.Send[[]objc.ID](t.ID, objc.Sel("exclusionPaths"))
	return objc.ConvertSlice(rv, func(id objc.ID) NSBezierPath {
		return NSBezierPathFromID(id)
	})
}
func (t NSTextContainer) SetExclusionPaths(value []NSBezierPath) {
	objc.Send[struct{}](t.ID, objc.Sel("setExclusionPaths:"), objectivec.IObjectSliceToNSArray(value))
}
// The behavior of the last line inside the text container.
//
// # Discussion
// 
// The [NSLineBreakMode] constants specify what happens when a line is too
// long for its container. For example, wrapping can occur on word boundaries
// (the default) or character boundaries, or the line can be clipped or
// truncated. The default value of this property is
// [NSLineBreakMode.byWordWrapping].
//
// [NSLineBreakMode.byWordWrapping]: https://developer.apple.com/documentation/AppKit/NSLineBreakMode/byWordWrapping
// [NSLineBreakMode]: https://developer.apple.com/documentation/AppKit/NSLineBreakMode
//
// See: https://developer.apple.com/documentation/AppKit/NSTextContainer/lineBreakMode
func (t NSTextContainer) LineBreakMode() NSLineBreakMode {
	rv := objc.Send[NSLineBreakMode](t.ID, objc.Sel("lineBreakMode"))
	return NSLineBreakMode(rv)
}
func (t NSTextContainer) SetLineBreakMode(value NSLineBreakMode) {
	objc.Send[struct{}](t.ID, objc.Sel("setLineBreakMode:"), value)
}
// A Boolean that controls whether the text container adjusts the width of its
// bounding rectangle when its text view resizes.
//
// # Discussion
// 
// When the value of this property is [true], the text container adjusts its
// width when the width of its text view changes. The default value of this
// property is [false].
// 
// For more information about size tracking, see [Text System Storage Layer
// Overview].
//
// [Text System Storage Layer Overview]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/TextStorageLayer/TextStorageLayer.html#//apple_ref/doc/uid/10000087i
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSTextContainer/widthTracksTextView
func (t NSTextContainer) WidthTracksTextView() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("widthTracksTextView"))
	return rv
}
func (t NSTextContainer) SetWidthTracksTextView(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setWidthTracksTextView:"), value)
}
// A Boolean that controls whether the text container adjusts the height of
// its bounding rectangle when its text view resizes.
//
// # Discussion
// 
// When the value of this property is [true], the text container adjusts its
// height when the height of its text view changes. The default value of this
// property is [false].
// 
// For more information about size tracking, see [Text System Storage Layer
// Overview].
//
// [Text System Storage Layer Overview]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/TextStorageLayer/TextStorageLayer.html#//apple_ref/doc/uid/10000087i
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSTextContainer/heightTracksTextView
func (t NSTextContainer) HeightTracksTextView() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("heightTracksTextView"))
	return rv
}
func (t NSTextContainer) SetHeightTracksTextView(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setHeightTracksTextView:"), value)
}
// The maximum number of lines that the text container can store.
//
// # Discussion
// 
// The layout manager uses the value of this property to determine the maximum
// number of lines associated with the text container. The default value of
// this property is `0`, which indicates that there is no limit.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextContainer/maximumNumberOfLines
func (t NSTextContainer) MaximumNumberOfLines() uint {
	rv := objc.Send[uint](t.ID, objc.Sel("maximumNumberOfLines"))
	return rv
}
func (t NSTextContainer) SetMaximumNumberOfLines(value uint) {
	objc.Send[struct{}](t.ID, objc.Sel("setMaximumNumberOfLines:"), value)
}
// The value for the text inset within line fragment rectangles.
//
// # Discussion
// 
// The padding appears at the beginning and end of the line fragment
// rectangles. The layout manager uses this value to determine the layout
// width. The default value of this property is `5.0`.
// 
// Line fragment padding is not designed to express text margins. Instead, you
// should use insets on your text view, adjust the paragraph margin
// attributes, or change the position of the text view within its superview.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextContainer/lineFragmentPadding
func (t NSTextContainer) LineFragmentPadding() float64 {
	rv := objc.Send[float64](t.ID, objc.Sel("lineFragmentPadding"))
	return rv
}
func (t NSTextContainer) SetLineFragmentPadding(value float64) {
	objc.Send[struct{}](t.ID, objc.Sel("setLineFragmentPadding:"), value)
}
// A Boolean that indicates whether the text container’s region is a
// rectangle with no holes or gaps, and whose edges are parallel to the text
// view’s coordinate system axes.
//
// # Discussion
// 
// The value of this property is [true] when the text container’s region is
// a rectangle with no holes or gaps and the edges are parallel to the text
// view’s coordinate system axes. The default value of this property is
// [false] when the [ExclusionPaths] property contains one or more items, when
// the [MaximumNumberOfLines] property is not zero, or when you override the
// [LineFragmentRectForProposedRectAtIndexWritingDirectionRemainingRect]
// method. Otherwise, the default value is [true].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSTextContainer/isSimpleRectangularTextContainer
func (t NSTextContainer) SimpleRectangularTextContainer() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("isSimpleRectangularTextContainer"))
	return rv
}
// The size of the text container’s bounding rectangle.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextContainer/containerSize
func (t NSTextContainer) ContainerSize() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](t.ID, objc.Sel("containerSize"))
	return corefoundation.CGSize(rv)
}
func (t NSTextContainer) SetContainerSize(value corefoundation.CGSize) {
	objc.Send[struct{}](t.ID, objc.Sel("setContainerSize:"), value)
}
// The default layout orientation.
//
// # Discussion
// 
// This property contains the default layout orientation for text in the
// object that adopts the protocol. If the text contains an explicit
// [verticalGlyphForm] attribute, that attribute overrides the value in this
// property. When rendering, TextKit assumes the coordinate system is
// appropriately rotated.
//
// [verticalGlyphForm]: https://developer.apple.com/documentation/Foundation/NSAttributedString/Key/verticalGlyphForm
//
// See: https://developer.apple.com/documentation/AppKit/NSTextLayoutOrientationProvider/layoutOrientation
func (t NSTextContainer) LayoutOrientation() NSTextLayoutOrientation {
	rv := objc.Send[NSTextLayoutOrientation](t.ID, objc.Sel("layoutOrientation"))
	return NSTextLayoutOrientation(rv)
}

			// Protocol methods for NSTextLayoutOrientationProvider
			

