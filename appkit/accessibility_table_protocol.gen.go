// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// A role-based protocol that declares the minimum interface necessary for an accessibility element to act as a table view.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityTable
type NSAccessibilityTable interface {
	objectivec.IObject
	NSAccessibilityElementProtocol
	NSAccessibilityGroup

	// Returns a short description of the table.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityTable/accessibilityLabel()
	AccessibilityLabel() string

	// Returns the row accessibility elements for the table.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityTable/accessibilityRows()
	AccessibilityRows() foundation.INSArray
}

// NSAccessibilityTableObject wraps an existing Objective-C object that conforms to the NSAccessibilityTable protocol.
type NSAccessibilityTableObject struct {
	objectivec.Object
}

func (o NSAccessibilityTableObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSAccessibilityTableObjectFromID constructs a [NSAccessibilityTableObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSAccessibilityTableObjectFromID(id objc.ID) NSAccessibilityTableObject {
	return NSAccessibilityTableObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Returns a short description of the table.
//
// # Return Value
//
// The description of the table.
//
// # Discussion
//
// This method is the getter for the [NSAccessibilityProtocol] protocol’s
// [NSWindow.AccessibilityLabel] property.
//
// Do not include the control’s type in the label (for example, use
// [Employees], not `Employees Table`). If possible use a single word. To help
// ensure that accessibility clients such as VoiceOver read the label with the
// correct intonation, start this label with a capital letter. Do not put a
// period at the end. Always localize the label.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityTable/accessibilityLabel()
func (o NSAccessibilityTableObject) AccessibilityLabel() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityLabel"))
	return foundation.NSStringFromID(rv).String()
}

// Returns the row accessibility elements for the table.
//
// # Return Value
//
// An array containing the table’s row elements.
//
// # Discussion
//
// This method is the getter for the [NSAccessibilityProtocol] protocol’s
// [NSWindow.AccessibilityRows] property.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityTable/accessibilityRows()
func (o NSAccessibilityTableObject) AccessibilityRows() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityRows"))
	return foundation.NSArrayFromID(rv)
}

// Returns the column header accessibility elements for the table.
//
// # Return Value
//
// The column header element.
//
// # Discussion
//
// This method is the getter for the [NSAccessibilityProtocol] protocol’s
// [NSWindow.AccessibilityColumnHeaderUIElements] property.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityTable/accessibilityColumnHeaderUIElements()
func (o NSAccessibilityTableObject) AccessibilityColumnHeaderUIElements() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityColumnHeaderUIElements"))
	return foundation.NSArrayFromID(rv)
}

// Returns the column accessibility elements for the table.
//
// # Return Value
//
// An array containing the table’s column elements.
//
// # Discussion
//
// This method is the getter for the [NSAccessibilityProtocol] protocol’s
// [NSWindow.AccessibilityColumns] property.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityTable/accessibilityColumns()
func (o NSAccessibilityTableObject) AccessibilityColumns() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityColumns"))
	return foundation.NSArrayFromID(rv)
}

// Returns the row header accessibility elements for the table.
//
// # Return Value
//
// The row header elements for the table.
//
// # Discussion
//
// This method is the getter for the [NSAccessibilityProtocol] protocol’s
// [NSWindow.AccessibilityRowHeaderUIElements] property.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityTable/accessibilityRowHeaderUIElements()
func (o NSAccessibilityTableObject) AccessibilityRowHeaderUIElements() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityRowHeaderUIElements"))
	return foundation.NSArrayFromID(rv)
}

// The currently selected cells for the table.
//
// # Return Value
//
// An array containing the currently selected cells for the table.
//
// # Discussion
//
// This method is the getter for the [NSAccessibilityProtocol] protocol’s
// [NSWindow.AccessibilitySelectedCells] property. Additionally, your class
// needs to send a [selectedCellsChanged] notification whenever the table’s
// selected cells change.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityTable/accessibilitySelectedCells()
//
// [selectedCellsChanged]: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Notification/selectedCellsChanged
func (o NSAccessibilityTableObject) AccessibilitySelectedCells() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilitySelectedCells"))
	return foundation.NSArrayFromID(rv)
}

// Returns the currently selected columns for the table.
//
// # Return Value
//
// An array containing the currently selected columns for the table.
//
// # Discussion
//
// This method is the getter for the [NSAccessibilityProtocol] protocol’s
// [NSWindow.AccessibilitySelectedColumns] property. Additionally, your class
// needs to send a [selectedColumnsChanged] notification whenever the
// table’s selected columns change.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityTable/accessibilitySelectedColumns()
//
// [selectedColumnsChanged]: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Notification/selectedColumnsChanged
func (o NSAccessibilityTableObject) AccessibilitySelectedColumns() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilitySelectedColumns"))
	return foundation.NSArrayFromID(rv)
}

// Returns the currently selected rows for the table.
//
// # Return Value
//
// An array containing the currently selected rows for the table.
//
// # Discussion
//
// This method is the getter for the [NSAccessibilityProtocol] protocol’s
// [NSWindow.AccessibilitySelectedRows] property. Additionally, your class
// needs to send a [selectedRowsChanged] notification whenever the table’s
// selected rows change.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityTable/accessibilitySelectedRows()
//
// [selectedRowsChanged]: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Notification/selectedRowsChanged
func (o NSAccessibilityTableObject) AccessibilitySelectedRows() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilitySelectedRows"))
	return foundation.NSArrayFromID(rv)
}

// Returns the visible cells for the table.
//
// # Return Value
//
// An array containing the currently visible cells.
//
// # Discussion
//
// This method is the getter for the [NSAccessibilityProtocol] protocol’s
// [NSWindow.AccessibilityVisibleCells] property.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityTable/accessibilityVisibleCells()
func (o NSAccessibilityTableObject) AccessibilityVisibleCells() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityVisibleCells"))
	return foundation.NSArrayFromID(rv)
}

// Returns the visible columns for the table.
//
// # Return Value
//
// An array containing the currently visible columns.
//
// # Discussion
//
// This method is the getter for the [NSAccessibilityProtocol] protocol’s
// [NSWindow.AccessibilityVisibleColumns] property.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityTable/accessibilityVisibleColumns()
func (o NSAccessibilityTableObject) AccessibilityVisibleColumns() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityVisibleColumns"))
	return foundation.NSArrayFromID(rv)
}

// Returns the visible rows for the table.
//
// # Return Value
//
// An array containing the currently visible rows.
//
// # Discussion
//
// This method is the getter for the [NSAccessibilityProtocol] protocol’s
// [NSWindow.AccessibilityVisibleRows] property.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityTable/accessibilityVisibleRows()
func (o NSAccessibilityTableObject) AccessibilityVisibleRows() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityVisibleRows"))
	return foundation.NSArrayFromID(rv)
}

// Sets the table’s currently selected rows.
//
// selectedRows: An array containing the row elements to be selected.
//
// # Discussion
//
// This method is the setter for the [NSAccessibilityProtocol] protocol’s
// [NSWindow.AccessibilitySelectedRows] property. Implementing this method
// allows the user to change the selected row using an accessibility client.
// Additionally, your class needs to send a [selectedRowsChanged] notification
// whenever the table’s selected rows change.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityTable/setAccessibilitySelectedRows(_:)
//
// [selectedRowsChanged]: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Notification/selectedRowsChanged
func (o NSAccessibilityTableObject) SetAccessibilitySelectedRows(selectedRows foundation.INSArray) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySelectedRows:"), selectedRows)
}

// Returns the accessibility element’s frame in screen coordinates.
//
// # Return Value
//
// The element’s frame in screen coordinates.
//
// # Discussion
//
// This method is the getter for the [NSAccessibilityProtocol] protocol’s
// [NSWindow.AccessibilityFrame] property. This method is called whenever
// accessibility clients request the [size] or [position] attributes.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityElementProtocol/accessibilityFrame()
//
// [position]: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Attribute/position
// [size]: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Attribute/size
func (o NSAccessibilityTableObject) AccessibilityFrame() corefoundation.CGRect {
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
// [NSWindow.AccessibilityParent] property.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityElementProtocol/accessibilityParent()
func (o NSAccessibilityTableObject) AccessibilityParent() objectivec.IObject {
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
// [NSWindow.AccessibilityIdentifier] property.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityElementProtocol/accessibilityIdentifier()
func (o NSAccessibilityTableObject) AccessibilityIdentifier() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityIdentifier"))
	return foundation.NSStringFromID(rv).String()
}

// Returns a Boolean value that indicates whether the accessibility element
// has the keyboard focus.
//
// # Return Value
//
// true if this element has the keyboard focus; otherwise, false.
//
// # Discussion
//
// This method is the getter for the [NSAccessibilityProtocol] protocol’s
// [NSWindow.AccessibilityFocused] property.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityElementProtocol/isAccessibilityFocused()
func (o NSAccessibilityTableObject) IsAccessibilityFocused() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityFocused"))
	return rv
}
