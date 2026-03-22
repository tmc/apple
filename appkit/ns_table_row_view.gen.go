// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSTableRowView] class.
var (
	_NSTableRowViewClass     NSTableRowViewClass
	_NSTableRowViewClassOnce sync.Once
)

func getNSTableRowViewClass() NSTableRowViewClass {
	_NSTableRowViewClassOnce.Do(func() {
		_NSTableRowViewClass = NSTableRowViewClass{class: objc.GetClass("NSTableRowView")}
	})
	return _NSTableRowViewClass
}

// GetNSTableRowViewClass returns the class object for NSTableRowView.
func GetNSTableRowViewClass() NSTableRowViewClass {
	return getNSTableRowViewClass()
}

type NSTableRowViewClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSTableRowViewClass) Alloc() NSTableRowView {
	rv := objc.Send[NSTableRowView](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// The view shown for a row in a table view.
//
// # Overview
// 
// [NSTableRowView] is responsible for displaying attributes associated with
// the row, including the selection highlight, and group row look.
//
// # Display Style
//
//   - [NSTableRowView.Emphasized]: Determines whether the row will draw with the alternate or secondary color (unless overridden).
//   - [NSTableRowView.SetEmphasized]
//   - [NSTableRowView.InteriorBackgroundStyle]: Specifies how the subviews should draw.
//   - [NSTableRowView.Floating]: Specifies whether the row is drawn using the floating style.
//   - [NSTableRowView.SetFloating]
//
// # Row Selection
//
//   - [NSTableRowView.Selected]: Determines whether the row is selected.
//   - [NSTableRowView.SetSelected]
//   - [NSTableRowView.SelectionHighlightStyle]: Specifies the selection highlight style.
//   - [NSTableRowView.SetSelectionHighlightStyle]
//
// # Drag and Drop
//
//   - [NSTableRowView.DraggingDestinationFeedbackStyle]: Specifies the dragging destination feedback style.
//   - [NSTableRowView.SetDraggingDestinationFeedbackStyle]
//   - [NSTableRowView.IndentationForDropOperation]: Defines the amount the drag target for a row should be indented.
//   - [NSTableRowView.SetIndentationForDropOperation]
//   - [NSTableRowView.TargetForDropOperation]: Specifies whether this row will draw a drop indicator based on the current dragging feedback style.
//   - [NSTableRowView.SetTargetForDropOperation]
//
// # Row Grouping
//
//   - [NSTableRowView.GroupRowStyle]: Specifies whether this row view is a group row.
//   - [NSTableRowView.SetGroupRowStyle]
//   - [NSTableRowView.NumberOfColumns]: Returns the number of columns represented by views in the table row view.
//
// # Overriding Row View Display Characteristics
//
//   - [NSTableRowView.BackgroundColor]: The background color of the row.
//   - [NSTableRowView.SetBackgroundColor]
//   - [NSTableRowView.DrawBackgroundInRect]: Draws the background of the row in the rectangle.
//   - [NSTableRowView.DrawDraggingDestinationFeedbackInRect]: Draws the row’s dragging destination feedback when the entire row is a drop target.
//   - [NSTableRowView.DrawSelectionInRect]: Draws the selected row.
//   - [NSTableRowView.DrawSeparatorInRect]: Draws the horizontal separator between table rows.
//
// # Accessing A Row Column View
//
//   - [NSTableRowView.ViewAtColumn]: Provides access to the given view at a particular column.
//
// # Instance Properties
//
//   - [NSTableRowView.NextRowSelected]
//   - [NSTableRowView.SetNextRowSelected]
//   - [NSTableRowView.PreviousRowSelected]
//   - [NSTableRowView.SetPreviousRowSelected]
//
// See: https://developer.apple.com/documentation/AppKit/NSTableRowView
type NSTableRowView struct {
	NSView
}

// NSTableRowViewFromID constructs a [NSTableRowView] from an objc.ID.
//
// The view shown for a row in a table view.
func NSTableRowViewFromID(id objc.ID) NSTableRowView {
	return NSTableRowView{NSView: NSViewFromID(id)}
}
// NOTE: NSTableRowView adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSTableRowView] class.
//
// # Display Style
//
//   - [INSTableRowView.Emphasized]: Determines whether the row will draw with the alternate or secondary color (unless overridden).
//   - [INSTableRowView.SetEmphasized]
//   - [INSTableRowView.InteriorBackgroundStyle]: Specifies how the subviews should draw.
//   - [INSTableRowView.Floating]: Specifies whether the row is drawn using the floating style.
//   - [INSTableRowView.SetFloating]
//
// # Row Selection
//
//   - [INSTableRowView.Selected]: Determines whether the row is selected.
//   - [INSTableRowView.SetSelected]
//   - [INSTableRowView.SelectionHighlightStyle]: Specifies the selection highlight style.
//   - [INSTableRowView.SetSelectionHighlightStyle]
//
// # Drag and Drop
//
//   - [INSTableRowView.DraggingDestinationFeedbackStyle]: Specifies the dragging destination feedback style.
//   - [INSTableRowView.SetDraggingDestinationFeedbackStyle]
//   - [INSTableRowView.IndentationForDropOperation]: Defines the amount the drag target for a row should be indented.
//   - [INSTableRowView.SetIndentationForDropOperation]
//   - [INSTableRowView.TargetForDropOperation]: Specifies whether this row will draw a drop indicator based on the current dragging feedback style.
//   - [INSTableRowView.SetTargetForDropOperation]
//
// # Row Grouping
//
//   - [INSTableRowView.GroupRowStyle]: Specifies whether this row view is a group row.
//   - [INSTableRowView.SetGroupRowStyle]
//   - [INSTableRowView.NumberOfColumns]: Returns the number of columns represented by views in the table row view.
//
// # Overriding Row View Display Characteristics
//
//   - [INSTableRowView.BackgroundColor]: The background color of the row.
//   - [INSTableRowView.SetBackgroundColor]
//   - [INSTableRowView.DrawBackgroundInRect]: Draws the background of the row in the rectangle.
//   - [INSTableRowView.DrawDraggingDestinationFeedbackInRect]: Draws the row’s dragging destination feedback when the entire row is a drop target.
//   - [INSTableRowView.DrawSelectionInRect]: Draws the selected row.
//   - [INSTableRowView.DrawSeparatorInRect]: Draws the horizontal separator between table rows.
//
// # Accessing A Row Column View
//
//   - [INSTableRowView.ViewAtColumn]: Provides access to the given view at a particular column.
//
// # Instance Properties
//
//   - [INSTableRowView.NextRowSelected]
//   - [INSTableRowView.SetNextRowSelected]
//   - [INSTableRowView.PreviousRowSelected]
//   - [INSTableRowView.SetPreviousRowSelected]
//
// See: https://developer.apple.com/documentation/AppKit/NSTableRowView
type INSTableRowView interface {
	INSView
	NSAccessibilityGroup
	NSAccessibilityRow

	// Topic: Display Style

	// Determines whether the row will draw with the alternate or secondary color (unless overridden).
	Emphasized() bool
	SetEmphasized(value bool)
	// Specifies how the subviews should draw.
	InteriorBackgroundStyle() NSBackgroundStyle
	// Specifies whether the row is drawn using the floating style.
	Floating() bool
	SetFloating(value bool)

	// Topic: Row Selection

	// Determines whether the row is selected.
	Selected() bool
	SetSelected(value bool)
	// Specifies the selection highlight style.
	SelectionHighlightStyle() NSTableViewSelectionHighlightStyle
	SetSelectionHighlightStyle(value NSTableViewSelectionHighlightStyle)

	// Topic: Drag and Drop

	// Specifies the dragging destination feedback style.
	DraggingDestinationFeedbackStyle() NSTableViewDraggingDestinationFeedbackStyle
	SetDraggingDestinationFeedbackStyle(value NSTableViewDraggingDestinationFeedbackStyle)
	// Defines the amount the drag target for a row should be indented.
	IndentationForDropOperation() float64
	SetIndentationForDropOperation(value float64)
	// Specifies whether this row will draw a drop indicator based on the current dragging feedback style.
	TargetForDropOperation() bool
	SetTargetForDropOperation(value bool)

	// Topic: Row Grouping

	// Specifies whether this row view is a group row.
	GroupRowStyle() bool
	SetGroupRowStyle(value bool)
	// Returns the number of columns represented by views in the table row view.
	NumberOfColumns() int

	// Topic: Overriding Row View Display Characteristics

	// The background color of the row.
	BackgroundColor() INSColor
	SetBackgroundColor(value INSColor)
	// Draws the background of the row in the rectangle.
	DrawBackgroundInRect(dirtyRect corefoundation.CGRect)
	// Draws the row’s dragging destination feedback when the entire row is a drop target.
	DrawDraggingDestinationFeedbackInRect(dirtyRect corefoundation.CGRect)
	// Draws the selected row.
	DrawSelectionInRect(dirtyRect corefoundation.CGRect)
	// Draws the horizontal separator between table rows.
	DrawSeparatorInRect(dirtyRect corefoundation.CGRect)

	// Topic: Accessing A Row Column View

	// Provides access to the given view at a particular column.
	ViewAtColumn(column int) objectivec.IObject

	// Topic: Instance Properties

	NextRowSelected() bool
	SetNextRowSelected(value bool)
	PreviousRowSelected() bool
	SetPreviousRowSelected(value bool)
}

// Init initializes the instance.
func (t NSTableRowView) Init() NSTableRowView {
	rv := objc.Send[NSTableRowView](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t NSTableRowView) Autorelease() NSTableRowView {
	rv := objc.Send[NSTableRowView](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSTableRowView creates a new NSTableRowView instance.
func NewNSTableRowView() NSTableRowView {
	class := getNSTableRowViewClass()
	rv := objc.Send[NSTableRowView](objc.ID(class.class), objc.Sel("new"))
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
func NewTableRowViewWithCoder(coder foundation.INSCoder) NSTableRowView {
	instance := getNSTableRowViewClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSTableRowViewFromID(rv)
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
func NewTableRowViewWithFrame(frameRect corefoundation.CGRect) NSTableRowView {
	instance := getNSTableRowViewClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFrame:"), frameRect)
	return NSTableRowViewFromID(rv)
}

// Draws the background of the row in the rectangle.
//
// dirtyRect: The rectangle that requires drawing.
//
// # Discussion
// 
// Overriding this method allows an application to draw a custom background
// for a table row view.
// 
// By default, this method draws the background color or group row style as
// appropriate for the row. This method also draws the “below look” for a
// drop target if [TargetForDropOperation] is [true].
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSTableRowView/drawBackground(in:)
func (t NSTableRowView) DrawBackgroundInRect(dirtyRect corefoundation.CGRect) {
	objc.Send[objc.ID](t.ID, objc.Sel("drawBackgroundInRect:"), dirtyRect)
}
// Draws the row’s dragging destination feedback when the entire row is a
// drop target.
//
// dirtyRect: The rectangle that requires drawing.
//
// # Discussion
// 
// Overriding this method allows an application to draw custom dragging
// destination feedback when the entire row is a drop target.
// 
// This method only is called if [TargetForDropOperation] is [true], and is
// only drawn based on the properties set, such as the group row style.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSTableRowView/drawDraggingDestinationFeedback(in:)
func (t NSTableRowView) DrawDraggingDestinationFeedbackInRect(dirtyRect corefoundation.CGRect) {
	objc.Send[objc.ID](t.ID, objc.Sel("drawDraggingDestinationFeedbackInRect:"), dirtyRect)
}
// Draws the selected row.
//
// dirtyRect: The rectangle that requires drawing.
//
// # Discussion
// 
// This method is only called if the selection should be drawn.
// 
// The selection will automatically be alpha-blended if the selection is
// animating in or out.
// 
// The default selection drawn is dependent on the [SelectionHighlightStyle].
//
// See: https://developer.apple.com/documentation/AppKit/NSTableRowView/drawSelection(in:)
func (t NSTableRowView) DrawSelectionInRect(dirtyRect corefoundation.CGRect) {
	objc.Send[objc.ID](t.ID, objc.Sel("drawSelectionInRect:"), dirtyRect)
}
// Draws the horizontal separator between table rows.
//
// dirtyRect: The rectangle that requires drawing.
//
// # Discussion
// 
// By default, the separator is only drawn if the enclosing table’s
// [GridStyleMask] is set to include a horizontal separator.
// 
// The separator should be drawn at the bottom of the row view, indicating a
// separation from this row and the next.
//
// See: https://developer.apple.com/documentation/AppKit/NSTableRowView/drawSeparator(in:)
func (t NSTableRowView) DrawSeparatorInRect(dirtyRect corefoundation.CGRect) {
	objc.Send[objc.ID](t.ID, objc.Sel("drawSeparatorInRect:"), dirtyRect)
}
// Provides access to the given view at a particular column.
//
// column: The index of the column.
//
// # Return Value
// 
// The view for the specified column.
//
// # Discussion
// 
// This is the only way to access cell views after the row view has been
// removed from the table.
//
// See: https://developer.apple.com/documentation/AppKit/NSTableRowView/view(atColumn:)
func (t NSTableRowView) ViewAtColumn(column int) objectivec.IObject {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("viewAtColumn:"), column)
	return objectivec.Object{ID: rv}
}
// Returns the indention level for the row.
//
// # Return Value
// 
// The disclosure level for the row.
//
// # Discussion
// 
// This method is the getter for the [NSAccessibilityProtocol] protocol’s
// [accessibilityDisclosureLevel] property.
//
// [accessibilityDisclosureLevel]: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityDisclosureLevel
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityRow/accessibilityDisclosureLevel()
func (t NSTableRowView) AccessibilityDisclosureLevel() int {
	rv := objc.Send[int](t.ID, objc.Sel("accessibilityDisclosureLevel"))
	return rv
}
// Returns the index for the row.
//
// # Return Value
// 
// The index for the row.
//
// # Discussion
// 
// This method is the getter for the [NSAccessibilityProtocol] protocol’s
// [accessibilityIndex] property.
//
// [accessibilityIndex]: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityIndex
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityRow/accessibilityIndex()
func (t NSTableRowView) AccessibilityIndex() int {
	rv := objc.Send[int](t.ID, objc.Sel("accessibilityIndex"))
	return rv
}

// Determines whether the row will draw with the alternate or secondary color
// (unless overridden).
//
// # Discussion
// 
// When emphasized is [true], the view will draw with the
// [alternateSelectedControlColor] defined by [NSColor]. When [false] it will
// use the [secondarySelectedControlColor] defined by [NSColor].
//
// [alternateSelectedControlColor]: https://developer.apple.com/documentation/AppKit/NSColor/alternateSelectedControlColor
// [false]: https://developer.apple.com/documentation/Swift/false
// [secondarySelectedControlColor]: https://developer.apple.com/documentation/AppKit/NSColor/secondarySelectedControlColor
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSTableRowView/isEmphasized
func (t NSTableRowView) Emphasized() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("isEmphasized"))
	return rv
}
func (t NSTableRowView) SetEmphasized(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setEmphasized:"), value)
}
// Specifies how the subviews should draw.
//
// # Discussion
// 
// This value is dynamically computed based on the set of properties set for
// the [NSTableRowView].
// 
// Subclassers can override this value when they draw differently based on the
// currently displayed properties.
// 
// This property can also be set to determine the color a subview should use.
// See [NSView.BackgroundStyle] for supported values.
//
// [NSView.BackgroundStyle]: https://developer.apple.com/documentation/AppKit/NSView/BackgroundStyle
//
// See: https://developer.apple.com/documentation/AppKit/NSTableRowView/interiorBackgroundStyle
func (t NSTableRowView) InteriorBackgroundStyle() NSBackgroundStyle {
	rv := objc.Send[NSBackgroundStyle](t.ID, objc.Sel("interiorBackgroundStyle"))
	return NSBackgroundStyle(rv)
}
// Specifies whether the row is drawn using the floating style.
//
// # Discussion
// 
// Floating is a temporary attribute that is set when a particular group row
// is actually floating above other rows. The state may change dynamically
// based on the position of the group row. Drawing may be different for rows
// that are currently ‘floating’.
//
// See: https://developer.apple.com/documentation/AppKit/NSTableRowView/isFloating
func (t NSTableRowView) Floating() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("isFloating"))
	return rv
}
func (t NSTableRowView) SetFloating(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setFloating:"), value)
}
// Determines whether the row is selected.
//
// # Discussion
// 
// [true] if selected, otherwise [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSTableRowView/isSelected
func (t NSTableRowView) Selected() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("isSelected"))
	return rv
}
func (t NSTableRowView) SetSelected(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setSelected:"), value)
}
// Specifies the selection highlight style.
//
// # Discussion
// 
// The possible values are specified in [NSTableView.SelectionHighlightStyle]
// in [NSTableView].
//
// [NSTableView.SelectionHighlightStyle]: https://developer.apple.com/documentation/AppKit/NSTableView/SelectionHighlightStyle-swift.enum
//
// See: https://developer.apple.com/documentation/AppKit/NSTableRowView/selectionHighlightStyle
func (t NSTableRowView) SelectionHighlightStyle() NSTableViewSelectionHighlightStyle {
	rv := objc.Send[NSTableViewSelectionHighlightStyle](t.ID, objc.Sel("selectionHighlightStyle"))
	return NSTableViewSelectionHighlightStyle(rv)
}
func (t NSTableRowView) SetSelectionHighlightStyle(value NSTableViewSelectionHighlightStyle) {
	objc.Send[struct{}](t.ID, objc.Sel("setSelectionHighlightStyle:"), value)
}
// Specifies the dragging destination feedback style.
//
// # Discussion
// 
// Possible values are defined in
// [NSTableView.DraggingDestinationFeedbackStyle].
//
// [NSTableView.DraggingDestinationFeedbackStyle]: https://developer.apple.com/documentation/AppKit/NSTableView/DraggingDestinationFeedbackStyle-swift.enum
//
// See: https://developer.apple.com/documentation/AppKit/NSTableRowView/draggingDestinationFeedbackStyle
func (t NSTableRowView) DraggingDestinationFeedbackStyle() NSTableViewDraggingDestinationFeedbackStyle {
	rv := objc.Send[NSTableViewDraggingDestinationFeedbackStyle](t.ID, objc.Sel("draggingDestinationFeedbackStyle"))
	return NSTableViewDraggingDestinationFeedbackStyle(rv)
}
func (t NSTableRowView) SetDraggingDestinationFeedbackStyle(value NSTableViewDraggingDestinationFeedbackStyle) {
	objc.Send[struct{}](t.ID, objc.Sel("setDraggingDestinationFeedbackStyle:"), value)
}
// Defines the amount the drag target for a row should be indented.
//
// # Discussion
// 
// The default is `0`.
//
// See: https://developer.apple.com/documentation/AppKit/NSTableRowView/indentationForDropOperation
func (t NSTableRowView) IndentationForDropOperation() float64 {
	rv := objc.Send[float64](t.ID, objc.Sel("indentationForDropOperation"))
	return rv
}
func (t NSTableRowView) SetIndentationForDropOperation(value float64) {
	objc.Send[struct{}](t.ID, objc.Sel("setIndentationForDropOperation:"), value)
}
// Specifies whether this row will draw a drop indicator based on the current
// dragging feedback style.
//
// # Discussion
// 
// When [true], the row view will draw a drop on indicator based on the
// current [DraggingDestinationFeedbackStyle].
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSTableRowView/isTargetForDropOperation
func (t NSTableRowView) TargetForDropOperation() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("isTargetForDropOperation"))
	return rv
}
func (t NSTableRowView) SetTargetForDropOperation(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setTargetForDropOperation:"), value)
}
// Specifies whether this row view is a group row.
//
// # Discussion
// 
// When [true] this row is a group row and will draw appropriately.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSTableRowView/isGroupRowStyle
func (t NSTableRowView) GroupRowStyle() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("isGroupRowStyle"))
	return rv
}
func (t NSTableRowView) SetGroupRowStyle(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setGroupRowStyle:"), value)
}
// Returns the number of columns represented by views in the table row view.
//
// # Discussion
// 
// The number of columns may not be equal to the number of columns in the
// enclosing [NSTableView], if this row view is a group style and has a single
// view that spans the entire width of the row.
//
// See: https://developer.apple.com/documentation/AppKit/NSTableRowView/numberOfColumns
func (t NSTableRowView) NumberOfColumns() int {
	rv := objc.Send[int](t.ID, objc.Sel("numberOfColumns"))
	return rv
}
// The background color of the row.
//
// # Discussion
// 
// The property defaults to the table view’s [BackgroundColor], unless
// [UsesAlternatingRowBackgroundColors] is set to [true]. In that case, the
// colors alternate, and are automatically updated as required by insertions
// and deletions.
// 
// The value of the background color can be customized in the
// [NSTableViewDelegate] method ``. The property is animatable.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSTableRowView/backgroundColor
func (t NSTableRowView) BackgroundColor() INSColor {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("backgroundColor"))
	return NSColorFromID(objc.ID(rv))
}
func (t NSTableRowView) SetBackgroundColor(value INSColor) {
	objc.Send[struct{}](t.ID, objc.Sel("setBackgroundColor:"), value)
}
// See: https://developer.apple.com/documentation/AppKit/NSTableRowView/isNextRowSelected
func (t NSTableRowView) NextRowSelected() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("isNextRowSelected"))
	return rv
}
func (t NSTableRowView) SetNextRowSelected(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setNextRowSelected:"), value)
}
// See: https://developer.apple.com/documentation/AppKit/NSTableRowView/isPreviousRowSelected
func (t NSTableRowView) PreviousRowSelected() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("isPreviousRowSelected"))
	return rv
}
func (t NSTableRowView) SetPreviousRowSelected(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setPreviousRowSelected:"), value)
}

			// Protocol methods for NSAccessibilityRow
			
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
func (o NSTableRowView) AccessibilityFrame() corefoundation.CGRect {
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
func (o NSTableRowView) AccessibilityParent() objectivec.IObject {
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
func (o NSTableRowView) AccessibilityIdentifier() string {
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
func (o NSTableRowView) IsAccessibilityFocused() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityFocused"))
	return rv
	}

