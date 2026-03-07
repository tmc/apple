// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// A role-based protocol that declares the minimum interface necessary for an accessibility element to act as an outline view.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityOutline
type NSAccessibilityOutline interface {
	objectivec.IObject
	NSAccessibilityGroup
	NSAccessibilityTable
}



// NSAccessibilityOutlineObject wraps an existing Objective-C object that conforms to the NSAccessibilityOutline protocol.
type NSAccessibilityOutlineObject struct {
	objectivec.Object
}
func (o NSAccessibilityOutlineObject) BaseObject() objectivec.Object {
	return o.Object
}



// NSAccessibilityOutlineObjectFromID constructs a [NSAccessibilityOutlineObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSAccessibilityOutlineObjectFromID(id objc.ID) NSAccessibilityOutlineObject {
	return NSAccessibilityOutlineObject{
		Object: objectivec.ObjectFromID(id),
	}
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
// [accessibilityFrame] property. This method is called whenever accessibility
// clients request the [size] or [position] attributes.
//
// [accessibilityFrame]: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityFrame
// [position]: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Attribute/position
// [size]: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Attribute/size
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityElementProtocol/accessibilityFrame()

func (o NSAccessibilityOutlineObject) AccessibilityFrame() corefoundation.CGRect {
	
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

func (o NSAccessibilityOutlineObject) AccessibilityParent() objectivec.IObject {
	
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

func (o NSAccessibilityOutlineObject) AccessibilityIdentifier() string {
	
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

func (o NSAccessibilityOutlineObject) IsAccessibilityFocused() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityFocused"))
	return rv
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
// [accessibilityLabel] property.
// 
// Do not include the control’s type in the label (for example, use
// [Employees], not `Employees Table`). If possible use a single word. To help
// ensure that accessibility clients such as VoiceOver read the label with the
// correct intonation, start this label with a capital letter. Do not put a
// period at the end. Always localize the label.
//
// [accessibilityLabel]: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityLabel
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityTable/accessibilityLabel()

func (o NSAccessibilityOutlineObject) AccessibilityLabel() string {
	
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
// [accessibilityRows] property.
//
// [accessibilityRows]: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityRows
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityTable/accessibilityRows()

func (o NSAccessibilityOutlineObject) AccessibilityRows() []objectivec.IObject {
	
	rv := objc.Send[[]objc.ID](o.ID, objc.Sel("accessibilityRows"))
	return objc.ConvertSlice(rv, func(id objc.ID) objectivec.IObject {
		return objectivec.Object{ID: id}
	})
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
// [accessibilityColumnHeaderUIElements] property.
//
// [accessibilityColumnHeaderUIElements]: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityColumnHeaderUIElements
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityTable/accessibilityColumnHeaderUIElements()

func (o NSAccessibilityOutlineObject) AccessibilityColumnHeaderUIElements() foundation.INSArray {
	
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
// [accessibilityColumns] property.
//
// [accessibilityColumns]: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityColumns
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityTable/accessibilityColumns()

func (o NSAccessibilityOutlineObject) AccessibilityColumns() foundation.INSArray {
	
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
// [accessibilityRowHeaderUIElements] property.
//
// [accessibilityRowHeaderUIElements]: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityRowHeaderUIElements
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityTable/accessibilityRowHeaderUIElements()

func (o NSAccessibilityOutlineObject) AccessibilityRowHeaderUIElements() foundation.INSArray {
	
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
// [accessibilitySelectedCells] property. Additionally, your class needs to
// send a [selectedCellsChanged] notification whenever the table’s selected
// cells change.
//
// [accessibilitySelectedCells]: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilitySelectedCells
// [selectedCellsChanged]: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Notification/selectedCellsChanged
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityTable/accessibilitySelectedCells()

func (o NSAccessibilityOutlineObject) AccessibilitySelectedCells() foundation.INSArray {
	
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
// [accessibilitySelectedColumns] property. Additionally, your class needs to
// send a [selectedColumnsChanged] notification whenever the table’s
// selected columns change.
//
// [accessibilitySelectedColumns]: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilitySelectedColumns
// [selectedColumnsChanged]: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Notification/selectedColumnsChanged
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityTable/accessibilitySelectedColumns()

func (o NSAccessibilityOutlineObject) AccessibilitySelectedColumns() foundation.INSArray {
	
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
// [accessibilitySelectedRows] property. Additionally, your class needs to
// send a [selectedRowsChanged] notification whenever the table’s selected
// rows change.
//
// [accessibilitySelectedRows]: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilitySelectedRows
// [selectedRowsChanged]: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Notification/selectedRowsChanged
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityTable/accessibilitySelectedRows()

func (o NSAccessibilityOutlineObject) AccessibilitySelectedRows() []objectivec.IObject {
	
	rv := objc.Send[[]objc.ID](o.ID, objc.Sel("accessibilitySelectedRows"))
	return objc.ConvertSlice(rv, func(id objc.ID) objectivec.IObject {
		return objectivec.Object{ID: id}
	})
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
// [accessibilityVisibleCells] property.
//
// [accessibilityVisibleCells]: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityVisibleCells
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityTable/accessibilityVisibleCells()

func (o NSAccessibilityOutlineObject) AccessibilityVisibleCells() foundation.INSArray {
	
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
// [accessibilityVisibleColumns] property.
//
// [accessibilityVisibleColumns]: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityVisibleColumns
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityTable/accessibilityVisibleColumns()

func (o NSAccessibilityOutlineObject) AccessibilityVisibleColumns() foundation.INSArray {
	
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
// [accessibilityVisibleRows] property.
//
// [accessibilityVisibleRows]: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityVisibleRows
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityTable/accessibilityVisibleRows()

func (o NSAccessibilityOutlineObject) AccessibilityVisibleRows() []objectivec.IObject {
	
	rv := objc.Send[[]objc.ID](o.ID, objc.Sel("accessibilityVisibleRows"))
	return objc.ConvertSlice(rv, func(id objc.ID) objectivec.IObject {
		return objectivec.Object{ID: id}
	})
	}

// Sets the table’s currently selected rows.
//
// selectedRows: An array containing the row elements to be selected.
//
// # Discussion
// 
// This method is the setter for the [NSAccessibilityProtocol] protocol’s
// [accessibilitySelectedRows] property. Implementing this method allows the
// user to change the selected row using an accessibility client.
// Additionally, your class needs to send a [selectedRowsChanged] notification
// whenever the table’s selected rows change.
//
// [accessibilitySelectedRows]: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilitySelectedRows
// [selectedRowsChanged]: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Notification/selectedRowsChanged
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityTable/setAccessibilitySelectedRows(_:)

func (o NSAccessibilityOutlineObject) SetAccessibilitySelectedRows(selectedRows []objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySelectedRows:"), objectivec.IObjectSliceToNSArray(selectedRows))
	}







