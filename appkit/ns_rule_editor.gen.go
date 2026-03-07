// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSRuleEditor] class.
var (
	_NSRuleEditorClass     NSRuleEditorClass
	_NSRuleEditorClassOnce sync.Once
)

func getNSRuleEditorClass() NSRuleEditorClass {
	_NSRuleEditorClassOnce.Do(func() {
		_NSRuleEditorClass = NSRuleEditorClass{class: objc.GetClass("NSRuleEditor")}
	})
	return _NSRuleEditorClass
}

// GetNSRuleEditorClass returns the class object for NSRuleEditor.
func GetNSRuleEditorClass() NSRuleEditorClass {
	return getNSRuleEditorClass()
}

type NSRuleEditorClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSRuleEditorClass) Alloc() NSRuleEditor {
	rv := objc.Send[NSRuleEditor](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// An interface for configuring a rule-based list of options.
//
// # Overview
// 
// A rule editor lets the user visually create and configure a list of options
// that are expressed as a predicate (as described in [Predicate Programming
// Guide]). Each row displayed by the rule editor represents a particular path
// down a tree of choices. The rule editor’s delegate provides the tree of
// choices to be displayed. The rule editor presents those choices to the user
// as a row of popup buttons, static text fields, and custom views.
// 
// [NSRuleEditor] exposes one binding, `rows`. You can bind `rows` to an
// ordered collection (such as an instance of [NSMutableArray]). Each object
// in the collection should have the following properties:
// 
// @“rowType”: An integer representing the type of the row
// ([NSRuleEditorRowType]). @“subrows”: An ordered to-many relation (such
// as an instance of [NSMutableArray]) containing the directly nested subrows
// for the given row. @“displayValues”: An ordered to-many relation
// containing the display values for the row. @“criteria”: An ordered
// to-many relation containing the criteria for the row.
//
// [Predicate Programming Guide]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/Predicates/AdditionalChapters/Introduction.html#//apple_ref/doc/uid/TP40001789
//
// # Configuring the Delegate
//
//   - [NSRuleEditor.Delegate]: The rule editor’s delegate.
//   - [NSRuleEditor.SetDelegate]
//
// # Configuring a Rule Editor
//
//   - [NSRuleEditor.Editable]: A Boolean value that determines whether the rule editor is editable.
//   - [NSRuleEditor.SetEditable]
//   - [NSRuleEditor.NestingMode]: The rule editor’s nesting mode.
//   - [NSRuleEditor.SetNestingMode]
//   - [NSRuleEditor.CanRemoveAllRows]: A Boolean value that indicates whether all the rows can be removed.
//   - [NSRuleEditor.SetCanRemoveAllRows]
//   - [NSRuleEditor.RowHeight]: The rule editor’s row height.
//   - [NSRuleEditor.SetRowHeight]
//
// # Working with Formatting
//
//   - [NSRuleEditor.FormattingDictionary]: The formatting dictionary for the rule editor.
//   - [NSRuleEditor.SetFormattingDictionary]
//   - [NSRuleEditor.FormattingStringsFilename]: The name of the rule editor’s strings file.
//   - [NSRuleEditor.SetFormattingStringsFilename]
//
// # Providing Data
//
//   - [NSRuleEditor.ReloadCriteria]: Instructs the receiver to refetch criteria from its delegate.
//   - [NSRuleEditor.SetCriteriaAndDisplayValuesForRowAtIndex]: Modifies the row at a given index to contain the given items and values.
//   - [NSRuleEditor.CriteriaForRow]: Returns the currently chosen items for a given row.
//   - [NSRuleEditor.DisplayValuesForRow]: Returns the chosen values for a given row.
//
// # Obtaining Row Information
//
//   - [NSRuleEditor.NumberOfRows]: The number of rows in the rule editor.
//   - [NSRuleEditor.ParentRowForRow]: Returns the index of the parent of a given row.
//   - [NSRuleEditor.RowForDisplayValue]: Returns the index of the row containing a given value.
//   - [NSRuleEditor.RowTypeForRow]: Returns the type of a given row.
//   - [NSRuleEditor.SubrowIndexesForRow]: Returns the immediate subrows of a given row.
//
// # Working with the Selection
//
//   - [NSRuleEditor.SelectedRowIndexes]: The indexes of the rule editor’s selected rows.
//   - [NSRuleEditor.SelectRowIndexesByExtendingSelection]: Sets in the receiver the indexes of rows that are selected.
//
// # Manipulating Rows
//
//   - [NSRuleEditor.AddRow]: Adds a row to the receiver.
//   - [NSRuleEditor.InsertRowAtIndexWithTypeAsSubrowOfRowAnimate]: Adds a new row of a given type at a given location.
//   - [NSRuleEditor.RemoveRowAtIndex]: Removes the row at a given index.
//   - [NSRuleEditor.RemoveRowsAtIndexesIncludeSubrows]: Removes the rows at given indexes.
//
// # Working with Predicates
//
//   - [NSRuleEditor.Predicate]: The rule editor’s predicate.
//   - [NSRuleEditor.ReloadPredicate]: Instructs the receiver to regenerate its predicate by invoking the corresponding delegate method.
//   - [NSRuleEditor.PredicateForRow]: Returns the predicate for a given row.
//
// # Supporting Bindings
//
//   - [NSRuleEditor.RowClass]: The class used to create a new row in the “rows” binding.
//   - [NSRuleEditor.SetRowClass]
//   - [NSRuleEditor.RowTypeKeyPath]: The key path for the row type.
//   - [NSRuleEditor.SetRowTypeKeyPath]
//   - [NSRuleEditor.SubrowsKeyPath]: The key path for the subrows.
//   - [NSRuleEditor.SetSubrowsKeyPath]
//   - [NSRuleEditor.CriteriaKeyPath]: The criteria key path.
//   - [NSRuleEditor.SetCriteriaKeyPath]
//   - [NSRuleEditor.DisplayValuesKeyPath]: The display values key path.
//   - [NSRuleEditor.SetDisplayValuesKeyPath]
//
// See: https://developer.apple.com/documentation/AppKit/NSRuleEditor
type NSRuleEditor struct {
	NSControl
}

// NSRuleEditorFromID constructs a [NSRuleEditor] from an objc.ID.
//
// An interface for configuring a rule-based list of options.
func NSRuleEditorFromID(id objc.ID) NSRuleEditor {
	return NSRuleEditor{NSControl: NSControlFromID(id)}
}
// NOTE: NSRuleEditor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [NSRuleEditor] class.
//
// # Configuring the Delegate
//
//   - [INSRuleEditor.Delegate]: The rule editor’s delegate.
//   - [INSRuleEditor.SetDelegate]
//
// # Configuring a Rule Editor
//
//   - [INSRuleEditor.Editable]: A Boolean value that determines whether the rule editor is editable.
//   - [INSRuleEditor.SetEditable]
//   - [INSRuleEditor.NestingMode]: The rule editor’s nesting mode.
//   - [INSRuleEditor.SetNestingMode]
//   - [INSRuleEditor.CanRemoveAllRows]: A Boolean value that indicates whether all the rows can be removed.
//   - [INSRuleEditor.SetCanRemoveAllRows]
//   - [INSRuleEditor.RowHeight]: The rule editor’s row height.
//   - [INSRuleEditor.SetRowHeight]
//
// # Working with Formatting
//
//   - [INSRuleEditor.FormattingDictionary]: The formatting dictionary for the rule editor.
//   - [INSRuleEditor.SetFormattingDictionary]
//   - [INSRuleEditor.FormattingStringsFilename]: The name of the rule editor’s strings file.
//   - [INSRuleEditor.SetFormattingStringsFilename]
//
// # Providing Data
//
//   - [INSRuleEditor.ReloadCriteria]: Instructs the receiver to refetch criteria from its delegate.
//   - [INSRuleEditor.SetCriteriaAndDisplayValuesForRowAtIndex]: Modifies the row at a given index to contain the given items and values.
//   - [INSRuleEditor.CriteriaForRow]: Returns the currently chosen items for a given row.
//   - [INSRuleEditor.DisplayValuesForRow]: Returns the chosen values for a given row.
//
// # Obtaining Row Information
//
//   - [INSRuleEditor.NumberOfRows]: The number of rows in the rule editor.
//   - [INSRuleEditor.ParentRowForRow]: Returns the index of the parent of a given row.
//   - [INSRuleEditor.RowForDisplayValue]: Returns the index of the row containing a given value.
//   - [INSRuleEditor.RowTypeForRow]: Returns the type of a given row.
//   - [INSRuleEditor.SubrowIndexesForRow]: Returns the immediate subrows of a given row.
//
// # Working with the Selection
//
//   - [INSRuleEditor.SelectedRowIndexes]: The indexes of the rule editor’s selected rows.
//   - [INSRuleEditor.SelectRowIndexesByExtendingSelection]: Sets in the receiver the indexes of rows that are selected.
//
// # Manipulating Rows
//
//   - [INSRuleEditor.AddRow]: Adds a row to the receiver.
//   - [INSRuleEditor.InsertRowAtIndexWithTypeAsSubrowOfRowAnimate]: Adds a new row of a given type at a given location.
//   - [INSRuleEditor.RemoveRowAtIndex]: Removes the row at a given index.
//   - [INSRuleEditor.RemoveRowsAtIndexesIncludeSubrows]: Removes the rows at given indexes.
//
// # Working with Predicates
//
//   - [INSRuleEditor.Predicate]: The rule editor’s predicate.
//   - [INSRuleEditor.ReloadPredicate]: Instructs the receiver to regenerate its predicate by invoking the corresponding delegate method.
//   - [INSRuleEditor.PredicateForRow]: Returns the predicate for a given row.
//
// # Supporting Bindings
//
//   - [INSRuleEditor.RowClass]: The class used to create a new row in the “rows” binding.
//   - [INSRuleEditor.SetRowClass]
//   - [INSRuleEditor.RowTypeKeyPath]: The key path for the row type.
//   - [INSRuleEditor.SetRowTypeKeyPath]
//   - [INSRuleEditor.SubrowsKeyPath]: The key path for the subrows.
//   - [INSRuleEditor.SetSubrowsKeyPath]
//   - [INSRuleEditor.CriteriaKeyPath]: The criteria key path.
//   - [INSRuleEditor.SetCriteriaKeyPath]
//   - [INSRuleEditor.DisplayValuesKeyPath]: The display values key path.
//   - [INSRuleEditor.SetDisplayValuesKeyPath]
//
// See: https://developer.apple.com/documentation/AppKit/NSRuleEditor
type INSRuleEditor interface {
	INSControl

	// Topic: Configuring the Delegate

	// The rule editor’s delegate.
	Delegate() NSRuleEditorDelegate
	SetDelegate(value NSRuleEditorDelegate)

	// Topic: Configuring a Rule Editor

	// A Boolean value that determines whether the rule editor is editable.
	Editable() bool
	SetEditable(value bool)
	// The rule editor’s nesting mode.
	NestingMode() NSRuleEditorNestingMode
	SetNestingMode(value NSRuleEditorNestingMode)
	// A Boolean value that indicates whether all the rows can be removed.
	CanRemoveAllRows() bool
	SetCanRemoveAllRows(value bool)
	// The rule editor’s row height.
	RowHeight() float64
	SetRowHeight(value float64)

	// Topic: Working with Formatting

	// The formatting dictionary for the rule editor.
	FormattingDictionary() foundation.INSDictionary
	SetFormattingDictionary(value foundation.INSDictionary)
	// The name of the rule editor’s strings file.
	FormattingStringsFilename() string
	SetFormattingStringsFilename(value string)

	// Topic: Providing Data

	// Instructs the receiver to refetch criteria from its delegate.
	ReloadCriteria()
	// Modifies the row at a given index to contain the given items and values.
	SetCriteriaAndDisplayValuesForRowAtIndex(criteria foundation.INSArray, values foundation.INSArray, rowIndex int)
	// Returns the currently chosen items for a given row.
	CriteriaForRow(row int) foundation.INSArray
	// Returns the chosen values for a given row.
	DisplayValuesForRow(row int) foundation.INSArray

	// Topic: Obtaining Row Information

	// The number of rows in the rule editor.
	NumberOfRows() int
	// Returns the index of the parent of a given row.
	ParentRowForRow(rowIndex int) int
	// Returns the index of the row containing a given value.
	RowForDisplayValue(displayValue objectivec.IObject) int
	// Returns the type of a given row.
	RowTypeForRow(rowIndex int) NSRuleEditorRowType
	// Returns the immediate subrows of a given row.
	SubrowIndexesForRow(rowIndex int) foundation.NSIndexSet

	// Topic: Working with the Selection

	// The indexes of the rule editor’s selected rows.
	SelectedRowIndexes() foundation.NSIndexSet
	// Sets in the receiver the indexes of rows that are selected.
	SelectRowIndexesByExtendingSelection(indexes foundation.NSIndexSet, extend bool)

	// Topic: Manipulating Rows

	// Adds a row to the receiver.
	AddRow(sender objectivec.IObject)
	// Adds a new row of a given type at a given location.
	InsertRowAtIndexWithTypeAsSubrowOfRowAnimate(rowIndex int, rowType NSRuleEditorRowType, parentRow int, shouldAnimate bool)
	// Removes the row at a given index.
	RemoveRowAtIndex(rowIndex int)
	// Removes the rows at given indexes.
	RemoveRowsAtIndexesIncludeSubrows(rowIndexes foundation.NSIndexSet, includeSubrows bool)

	// Topic: Working with Predicates

	// The rule editor’s predicate.
	Predicate() foundation.INSPredicate
	// Instructs the receiver to regenerate its predicate by invoking the corresponding delegate method.
	ReloadPredicate()
	// Returns the predicate for a given row.
	PredicateForRow(row int) foundation.INSPredicate

	// Topic: Supporting Bindings

	// The class used to create a new row in the “rows” binding.
	RowClass() objc.Class
	SetRowClass(value objc.Class)
	// The key path for the row type.
	RowTypeKeyPath() string
	SetRowTypeKeyPath(value string)
	// The key path for the subrows.
	SubrowsKeyPath() string
	SetSubrowsKeyPath(value string)
	// The criteria key path.
	CriteriaKeyPath() string
	SetCriteriaKeyPath(value string)
	// The display values key path.
	DisplayValuesKeyPath() string
	SetDisplayValuesKeyPath(value string)

	EncodeWithCoder(coder foundation.INSCoder)
}





// Init initializes the instance.
func (r NSRuleEditor) Init() NSRuleEditor {
	rv := objc.Send[NSRuleEditor](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r NSRuleEditor) Autorelease() NSRuleEditor {
	rv := objc.Send[NSRuleEditor](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSRuleEditor creates a new NSRuleEditor instance.
func NewNSRuleEditor() NSRuleEditor {
	class := getNSRuleEditorClass()
	rv := objc.Send[NSRuleEditor](objc.ID(class.class), objc.Sel("new"))
	return rv
}






// Initializes a control with data in an unarchiver.
//
// See: https://developer.apple.com/documentation/AppKit/NSControl/init(coder:)
func NewRuleEditorWithCoder(coder foundation.INSCoder) NSRuleEditor {
	instance := getNSRuleEditorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSRuleEditorFromID(rv)
}


// Initializes a control with the specified frame rectangle.
//
// frameRect: The rectangle of the control, specified in points in the coordinate space
// of the enclosing view.
//
// # Return Value
// 
// An initialized control object, or `nil` if the object couldn’t be
// initialized.
//
// # Discussion
// 
// If a cell has been specified for controls of this type, this method also
// creates an instance of the cell. Because [NSControl] is an abstract class,
// invocations of this method should appear only in the designated
// initializers of subclasses; that is, there should always be a more specific
// designated initializer for the subclass, as this method is the designated
// initializer for [NSControl].
//
// See: https://developer.apple.com/documentation/AppKit/NSControl/init(frame:)
func NewRuleEditorWithFrame(frameRect corefoundation.CGRect) NSRuleEditor {
	instance := getNSRuleEditorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFrame:"), frameRect)
	return NSRuleEditorFromID(rv)
}







// Instructs the receiver to refetch criteria from its delegate.
//
// # Discussion
// 
// You can use this method to indicate that the available criteria may have
// changed and should be refetched from the delegate and the popups
// recalculated. If any item in a given row is “orphaned” (that is, is no
// longer reported as a child of its previous parent), its criteria and
// display values are set to valid choices.
//
// See: https://developer.apple.com/documentation/AppKit/NSRuleEditor/reloadCriteria()
func (r NSRuleEditor) ReloadCriteria() {
	objc.Send[objc.ID](r.ID, objc.Sel("reloadCriteria"))
}

// Modifies the row at a given index to contain the given items and values.
//
// criteria: The array of criteria for the row at `rowIndex`. Pass an empty array to
// force the receiver to query its delegate. This value must not be `nil`.
//
// values: The array of values for the row at `rowIndex`. Pass an empty array to force
// the receiver to query its delegate. This value must not be `nil`.
//
// rowIndex: The index of a row in the receiver.
//
// # Discussion
// 
// It is your responsibility to ensure that each item in the array is a child
// of the previous item, and that the first item is a root item for the row
// type. If the last item has child items, then the items array will be
// extended by querying the delegate for child items until a childless item is
// reached. If `values` contains fewer objects than the (possibly extended)
// criteria array, then the delegate is queried to construct the remaining
// display values. If you want the delegate to be queried for all the criteria
// or all the display values, pass empty arrays; do not pass `nil`.
//
// See: https://developer.apple.com/documentation/AppKit/NSRuleEditor/setCriteria(_:andDisplayValues:forRowAt:)
func (r NSRuleEditor) SetCriteriaAndDisplayValuesForRowAtIndex(criteria foundation.INSArray, values foundation.INSArray, rowIndex int) {
	objc.Send[objc.ID](r.ID, objc.Sel("setCriteria:andDisplayValues:forRowAtIndex:"), criteria, values, rowIndex)
}

// Returns the currently chosen items for a given row.
//
// row: The index of a row in the receiver.
//
// # Return Value
// 
// The currently chosen items for row `row`.
//
// # Discussion
// 
// The items returned are the same as those returned by calling the
// delegate’s [RuleEditorChildForCriterionWithRowType] method once for each
// item in the row.
//
// See: https://developer.apple.com/documentation/AppKit/NSRuleEditor/criteria(forRow:)
func (r NSRuleEditor) CriteriaForRow(row int) foundation.INSArray {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("criteriaForRow:"), row)
	return foundation.NSArrayFromID(rv)
}

// Returns the chosen values for a given row.
//
// row: The index of a row in the receiver.
//
// # Return Value
// 
// The chosen values (strings, views, or menu items) for row `row`.
//
// # Discussion
// 
// The values returned are the same as those returned from the delegate’s
// [RuleEditorDisplayValueForCriterionInRow] method.
//
// See: https://developer.apple.com/documentation/AppKit/NSRuleEditor/displayValues(forRow:)
func (r NSRuleEditor) DisplayValuesForRow(row int) foundation.INSArray {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("displayValuesForRow:"), row)
	return foundation.NSArrayFromID(rv)
}

// Returns the index of the parent of a given row.
//
// rowIndex: The index of a row in the receiver.
//
// # Return Value
// 
// The index of the parent of the row at `rowIndex`. If the row at `rowIndex`
// is a root row, returns `-1`.
//
// See: https://developer.apple.com/documentation/AppKit/NSRuleEditor/parentRow(forRow:)
func (r NSRuleEditor) ParentRowForRow(rowIndex int) int {
	rv := objc.Send[int](r.ID, objc.Sel("parentRowForRow:"), rowIndex)
	return rv
}

// Returns the index of the row containing a given value.
//
// displayValue: The display value (string, view, or menu item) of an item in the receiver.
// This value must not be `nil`.
//
// # Return Value
// 
// The index of the row containing `displayValue`, or [NSNotFound].
//
// # Discussion
// 
// This method searches each row via pointer equality for the given display
// value, which may be present as an alternative in a popup menu for that row.
//
// See: https://developer.apple.com/documentation/AppKit/NSRuleEditor/row(forDisplayValue:)
func (r NSRuleEditor) RowForDisplayValue(displayValue objectivec.IObject) int {
	rv := objc.Send[int](r.ID, objc.Sel("rowForDisplayValue:"), displayValue)
	return rv
}

// Returns the type of a given row.
//
// rowIndex: The index of a row in the receiver.
//
// # Return Value
// 
// The type of the row at `rowIndex`.
//
// See: https://developer.apple.com/documentation/AppKit/NSRuleEditor/rowType(forRow:)
func (r NSRuleEditor) RowTypeForRow(rowIndex int) NSRuleEditorRowType {
	rv := objc.Send[NSRuleEditorRowType](r.ID, objc.Sel("rowTypeForRow:"), rowIndex)
	return NSRuleEditorRowType(rv)
}

// Returns the immediate subrows of a given row.
//
// rowIndex: The index of a row in the receiver, or `-1` to get the top-level rows.
//
// # Return Value
// 
// The immediate subrows of the row at `rowIndex`.
//
// # Discussion
// 
// Rows are numbered starting at `0`.
//
// See: https://developer.apple.com/documentation/AppKit/NSRuleEditor/subrowIndexes(forRow:)
func (r NSRuleEditor) SubrowIndexesForRow(rowIndex int) foundation.NSIndexSet {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("subrowIndexesForRow:"), rowIndex)
	return foundation.NSIndexSetFromID(rv)
}

// Sets in the receiver the indexes of rows that are selected.
//
// indexes: The indexes of rows in the receiver to select.
//
// extend: If [false], the selected rows are specified by `indexes`. If [true], the
// rows indicated by `indexes` are added to the collection of already selected
// rows, providing multiple selection.
// //
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSRuleEditor/selectRowIndexes(_:byExtendingSelection:)
func (r NSRuleEditor) SelectRowIndexesByExtendingSelection(indexes foundation.NSIndexSet, extend bool) {
	objc.Send[objc.ID](r.ID, objc.Sel("selectRowIndexes:byExtendingSelection:"), indexes, extend)
}

// Adds a row to the receiver.
//
// sender: Typically the object that sent the message.
//
// See: https://developer.apple.com/documentation/AppKit/NSRuleEditor/addRow(_:)
func (r NSRuleEditor) AddRow(sender objectivec.IObject) {
	objc.Send[objc.ID](r.ID, objc.Sel("addRow:"), sender)
}

// Adds a new row of a given type at a given location.
//
// rowIndex: The index at which the new row should be inserted. `rowIndex` must be
// greater than `parentRow`, and much specify a row that does not fall amongst
// the children of some other parent.
//
// rowType: The type of the new row.
//
// parentRow: The index of the row of which the new row is a child. Pass `-1` to indicate
// that the new row should be a root row.
//
// shouldAnimate: [true] if creation of the new row should be animated, otherwise [false].
// //
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
//
// See: https://developer.apple.com/documentation/AppKit/NSRuleEditor/insertRow(at:with:asSubrowOfRow:animate:)
func (r NSRuleEditor) InsertRowAtIndexWithTypeAsSubrowOfRowAnimate(rowIndex int, rowType NSRuleEditorRowType, parentRow int, shouldAnimate bool) {
	objc.Send[objc.ID](r.ID, objc.Sel("insertRowAtIndex:withType:asSubrowOfRow:animate:"), rowIndex, rowType, parentRow, shouldAnimate)
}

// Removes the row at a given index.
//
// rowIndex: The index of a row in the receiver.
//
// # Discussion
// 
// Any subrows of the deleted row are adopted by the parent of the deleted
// row, or are made root rows.
//
// See: https://developer.apple.com/documentation/AppKit/NSRuleEditor/removeRow(at:)
func (r NSRuleEditor) RemoveRowAtIndex(rowIndex int) {
	objc.Send[objc.ID](r.ID, objc.Sel("removeRowAtIndex:"), rowIndex)
}

// Removes the rows at given indexes.
//
// rowIndexes: Indexes of one or more rows in the receiver.
//
// includeSubrows: If [true], then sub-rows of deleted rows are also deleted; if [false], then
// each sub-row is adopted by its first non-deleted ancestor, or becomes a
// root row.
// //
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSRuleEditor/removeRows(at:includeSubrows:)
func (r NSRuleEditor) RemoveRowsAtIndexesIncludeSubrows(rowIndexes foundation.NSIndexSet, includeSubrows bool) {
	objc.Send[objc.ID](r.ID, objc.Sel("removeRowsAtIndexes:includeSubrows:"), rowIndexes, includeSubrows)
}

// Instructs the receiver to regenerate its predicate by invoking the
// corresponding delegate method.
//
// # Discussion
// 
// You typically invoke this method because something has changed (for
// example, a view’s value).
//
// See: https://developer.apple.com/documentation/AppKit/NSRuleEditor/reloadPredicate()
func (r NSRuleEditor) ReloadPredicate() {
	objc.Send[objc.ID](r.ID, objc.Sel("reloadPredicate"))
}

// Returns the predicate for a given row.
//
// row: The index of a row in the receiver.
//
// # Return Value
// 
// The predicate for the row at `row`.
//
// # Discussion
// 
// You should rarely have a need to call this directly, but you can override
// this method in a subclass to perform specialized predicate handling for
// certain criteria or display values.
//
// See: https://developer.apple.com/documentation/AppKit/NSRuleEditor/predicate(forRow:)
func (r NSRuleEditor) PredicateForRow(row int) foundation.INSPredicate {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("predicateForRow:"), row)
	return foundation.NSPredicateFromID(rv)
}
func (r NSRuleEditor) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](r.ID, objc.Sel("encodeWithCoder:"), coder)
}












// The rule editor’s delegate.
//
// See: https://developer.apple.com/documentation/AppKit/NSRuleEditor/delegate
func (r NSRuleEditor) Delegate() NSRuleEditorDelegate {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("delegate"))
	return NSRuleEditorDelegateObjectFromID(rv)
}
func (r NSRuleEditor) SetDelegate(value NSRuleEditorDelegate) {
	objc.Send[struct{}](r.ID, objc.Sel("setDelegate:"), value)
}



// A Boolean value that determines whether the rule editor is editable.
//
// # Discussion
// 
// The default is [true].
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSRuleEditor/isEditable
func (r NSRuleEditor) Editable() bool {
	rv := objc.Send[bool](r.ID, objc.Sel("isEditable"))
	return rv
}
func (r NSRuleEditor) SetEditable(value bool) {
	objc.Send[struct{}](r.ID, objc.Sel("setEditable:"), value)
}



// The rule editor’s nesting mode.
//
// # Discussion
// 
// You typically set the nesting mode at view creation time and do not
// subsequently modify it. The default is [NSRuleEditorNestingModeCompound].
// For a list of valid modes, see `Nesting Modes`.
//
// See: https://developer.apple.com/documentation/AppKit/NSRuleEditor/nestingMode-swift.property
func (r NSRuleEditor) NestingMode() NSRuleEditorNestingMode {
	rv := objc.Send[NSRuleEditorNestingMode](r.ID, objc.Sel("nestingMode"))
	return NSRuleEditorNestingMode(rv)
}
func (r NSRuleEditor) SetNestingMode(value NSRuleEditorNestingMode) {
	objc.Send[struct{}](r.ID, objc.Sel("setNestingMode:"), value)
}



// A Boolean value that indicates whether all the rows can be removed.
//
// See: https://developer.apple.com/documentation/AppKit/NSRuleEditor/canRemoveAllRows
func (r NSRuleEditor) CanRemoveAllRows() bool {
	rv := objc.Send[bool](r.ID, objc.Sel("canRemoveAllRows"))
	return rv
}
func (r NSRuleEditor) SetCanRemoveAllRows(value bool) {
	objc.Send[struct{}](r.ID, objc.Sel("setCanRemoveAllRows:"), value)
}



// The rule editor’s row height.
//
// See: https://developer.apple.com/documentation/AppKit/NSRuleEditor/rowHeight
func (r NSRuleEditor) RowHeight() float64 {
	rv := objc.Send[float64](r.ID, objc.Sel("rowHeight"))
	return rv
}
func (r NSRuleEditor) SetRowHeight(value float64) {
	objc.Send[struct{}](r.ID, objc.Sel("setRowHeight:"), value)
}



// The formatting dictionary for the rule editor.
//
// # Discussion
// 
// If you assign a new the formatting dictionary to this property, it sets the
// current to formatting strings file name to `nil`.
//
// See: https://developer.apple.com/documentation/AppKit/NSRuleEditor/formattingDictionary
func (r NSRuleEditor) FormattingDictionary() foundation.INSDictionary {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("formattingDictionary"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (r NSRuleEditor) SetFormattingDictionary(value foundation.INSDictionary) {
	objc.Send[struct{}](r.ID, objc.Sel("setFormattingDictionary:"), value)
}



// The name of the rule editor’s strings file.
//
// # Discussion
// 
// The [NSRuleEditor] class looks for a strings file with the given name in
// the main bundle and (if appropriate) the bundle containing the nib file
// from which it was loaded. If it finds a strings file resource with the
// given name, [NSRuleEditor] loads it and sets it as the formatting
// dictionary for the receiver. You can obtain the resulting dictionary using
// the [FormattingDictionary] property].
// 
// If you assign a new dictionary to the [FormattingDictionary] property, it
// sets the current to formatting strings file name to `nil`.
//
// See: https://developer.apple.com/documentation/AppKit/NSRuleEditor/formattingStringsFilename
func (r NSRuleEditor) FormattingStringsFilename() string {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("formattingStringsFilename"))
	return foundation.NSStringFromID(rv).String()
}
func (r NSRuleEditor) SetFormattingStringsFilename(value string) {
	objc.Send[struct{}](r.ID, objc.Sel("setFormattingStringsFilename:"), objc.String(value))
}



// The number of rows in the rule editor.
//
// See: https://developer.apple.com/documentation/AppKit/NSRuleEditor/numberOfRows
func (r NSRuleEditor) NumberOfRows() int {
	rv := objc.Send[int](r.ID, objc.Sel("numberOfRows"))
	return rv
}



// The indexes of the rule editor’s selected rows.
//
// See: https://developer.apple.com/documentation/AppKit/NSRuleEditor/selectedRowIndexes
func (r NSRuleEditor) SelectedRowIndexes() foundation.NSIndexSet {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("selectedRowIndexes"))
	return foundation.NSIndexSetFromID(objc.ID(rv))
}



// The rule editor’s predicate.
//
// # Discussion
// 
// If the delegate implements [NSRuleEditor], this property contains the rule
// editor’s predicate. If the delegate does not implement [NSRuleEditor], or
// if the delegate does not return enough parts to construct a full predicate,
// this property contains `nil`.
//
// See: https://developer.apple.com/documentation/AppKit/NSRuleEditor/predicate
func (r NSRuleEditor) Predicate() foundation.INSPredicate {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("predicate"))
	return foundation.NSPredicateFromID(objc.ID(rv))
}



// The class used to create a new row in the “rows” binding.
//
// # Discussion
// 
// The default is [NSMutableDictionary].
//
// See: https://developer.apple.com/documentation/AppKit/NSRuleEditor/rowClass
func (r NSRuleEditor) RowClass() objc.Class {
	rv := objc.Send[objc.Class](r.ID, objc.Sel("rowClass"))
	return rv
}
func (r NSRuleEditor) SetRowClass(value objc.Class) {
	objc.Send[struct{}](r.ID, objc.Sel("setRowClass:"), value)
}



// The key path for the row type.
//
// # Discussion
// 
// The key path is used to get the row type in the “rows” binding. The
// corresponding property should be a number that specifies an
// [NSRuleEditorRowType] value (see `Row Types`).
// 
// The default value is `@"rowType"`.
//
// See: https://developer.apple.com/documentation/AppKit/NSRuleEditor/rowTypeKeyPath
func (r NSRuleEditor) RowTypeKeyPath() string {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("rowTypeKeyPath"))
	return foundation.NSStringFromID(rv).String()
}
func (r NSRuleEditor) SetRowTypeKeyPath(value string) {
	objc.Send[struct{}](r.ID, objc.Sel("setRowTypeKeyPath:"), objc.String(value))
}



// The key path for the subrows.
//
// # Discussion
// 
// The key path is used to get the nested rows in the “rows” binding. The
// corresponding property should be an ordered to-many relationship containing
// additional bound row objects.
// 
// The default value is `@"subrows"`.
//
// See: https://developer.apple.com/documentation/AppKit/NSRuleEditor/subrowsKeyPath
func (r NSRuleEditor) SubrowsKeyPath() string {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("subrowsKeyPath"))
	return foundation.NSStringFromID(rv).String()
}
func (r NSRuleEditor) SetSubrowsKeyPath(value string) {
	objc.Send[struct{}](r.ID, objc.Sel("setSubrowsKeyPath:"), objc.String(value))
}



// The criteria key path.
//
// # Discussion
// 
// The key path is used to get the criteria for a row in the “rows”
// binding. The criteria objects are the objects returned by calling the
// delegate’s [RuleEditorChildForCriterionWithRowType] method once for every
// child in the specified row. The corresponding property is an ordered
// to-many relationship.
// 
// The default value is `@"criteria"`.
//
// See: https://developer.apple.com/documentation/AppKit/NSRuleEditor/criteriaKeyPath
func (r NSRuleEditor) CriteriaKeyPath() string {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("criteriaKeyPath"))
	return foundation.NSStringFromID(rv).String()
}
func (r NSRuleEditor) SetCriteriaKeyPath(value string) {
	objc.Send[struct{}](r.ID, objc.Sel("setCriteriaKeyPath:"), objc.String(value))
}



// The display values key path.
//
// # Discussion
// 
// The key path is used to get the display values for a row in the “rows”
// binding. The display values are the objects returned by calling the
// delegate’s [RuleEditorDisplayValueForCriterionInRow] method for the
// specified row. The corresponding property is an ordered to-many
// relationship.
// 
// The default is `@"displayValues"`.
//
// See: https://developer.apple.com/documentation/AppKit/NSRuleEditor/displayValuesKeyPath
func (r NSRuleEditor) DisplayValuesKeyPath() string {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("displayValuesKeyPath"))
	return foundation.NSStringFromID(rv).String()
}
func (r NSRuleEditor) SetDisplayValuesKeyPath(value string) {
	objc.Send[struct{}](r.ID, objc.Sel("setDisplayValuesKeyPath:"), objc.String(value))
}



































