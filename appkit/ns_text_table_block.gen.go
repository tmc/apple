// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [NSTextTableBlock] class.
var (
	_NSTextTableBlockClass     NSTextTableBlockClass
	_NSTextTableBlockClassOnce sync.Once
)

func getNSTextTableBlockClass() NSTextTableBlockClass {
	_NSTextTableBlockClassOnce.Do(func() {
		_NSTextTableBlockClass = NSTextTableBlockClass{class: objc.GetClass("NSTextTableBlock")}
	})
	return _NSTextTableBlockClass
}

// GetNSTextTableBlockClass returns the class object for NSTextTableBlock.
func GetNSTextTableBlockClass() NSTextTableBlockClass {
	return getNSTextTableBlockClass()
}

type NSTextTableBlockClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSTextTableBlockClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSTextTableBlockClass) Alloc() NSTextTableBlock {
	rv := objc.Send[NSTextTableBlock](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A text block that appears as a cell in a text table.
//
// # Creation
//
//   - [NSTextTableBlock.InitWithTableStartingRowRowSpanStartingColumnColumnSpan]: Returns an initialized text table block.
//
// # Getting the block’s enclosing table
//
//   - [NSTextTableBlock.Table]: Returns the table containing this text table block.
//
// # Getting information about the block’s position in its enclosing table
//
//   - [NSTextTableBlock.StartingRow]: Returns the table row at which this text table block starts.
//   - [NSTextTableBlock.RowSpan]: Returns the number of table rows spanned by this text table block.
//   - [NSTextTableBlock.StartingColumn]: Returns the table column at which this text table block starts.
//   - [NSTextTableBlock.ColumnSpan]: Returns the number of table columns spanned by this text table block.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextTableBlock
type NSTextTableBlock struct {
	NSTextBlock
}

// NSTextTableBlockFromID constructs a [NSTextTableBlock] from an objc.ID.
//
// A text block that appears as a cell in a text table.
func NSTextTableBlockFromID(id objc.ID) NSTextTableBlock {
	return NSTextTableBlock{NSTextBlock: NSTextBlockFromID(id)}
}

// NOTE: NSTextTableBlock adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSTextTableBlock] class.
//
// # Creation
//
//   - [INSTextTableBlock.InitWithTableStartingRowRowSpanStartingColumnColumnSpan]: Returns an initialized text table block.
//
// # Getting the block’s enclosing table
//
//   - [INSTextTableBlock.Table]: Returns the table containing this text table block.
//
// # Getting information about the block’s position in its enclosing table
//
//   - [INSTextTableBlock.StartingRow]: Returns the table row at which this text table block starts.
//   - [INSTextTableBlock.RowSpan]: Returns the number of table rows spanned by this text table block.
//   - [INSTextTableBlock.StartingColumn]: Returns the table column at which this text table block starts.
//   - [INSTextTableBlock.ColumnSpan]: Returns the number of table columns spanned by this text table block.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextTableBlock
type INSTextTableBlock interface {
	INSTextBlock

	// Topic: Creation

	// Returns an initialized text table block.
	InitWithTableStartingRowRowSpanStartingColumnColumnSpan(table INSTextTable, row int, rowSpan int, col int, colSpan int) NSTextTableBlock

	// Topic: Getting the block’s enclosing table

	// Returns the table containing this text table block.
	Table() INSTextTable

	// Topic: Getting information about the block’s position in its enclosing table

	// Returns the table row at which this text table block starts.
	StartingRow() int
	// Returns the number of table rows spanned by this text table block.
	RowSpan() int
	// Returns the table column at which this text table block starts.
	StartingColumn() int
	// Returns the number of table columns spanned by this text table block.
	ColumnSpan() int
}

// Init initializes the instance.
func (t NSTextTableBlock) Init() NSTextTableBlock {
	rv := objc.Send[NSTextTableBlock](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t NSTextTableBlock) Autorelease() NSTextTableBlock {
	rv := objc.Send[NSTextTableBlock](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSTextTableBlock creates a new NSTextTableBlock instance.
func NewNSTextTableBlock() NSTextTableBlock {
	class := getNSTextTableBlockClass()
	rv := objc.Send[NSTextTableBlock](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Returns an initialized text table block.
//
// table: The text table containing this text table block.
//
// row: The table row at which the text table block starts.
//
// rowSpan: How many rows the text table block covers.
//
// col: The table column at which the text table block starts.
//
// colSpan: How many columns the text table block covers.
//
// # Discussion
//
// This is the designated initializer.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextTableBlock/init(table:startingRow:rowSpan:startingColumn:columnSpan:)
func NewTextTableBlockWithTableStartingRowRowSpanStartingColumnColumnSpan(table INSTextTable, row int, rowSpan int, col int, colSpan int) NSTextTableBlock {
	instance := getNSTextTableBlockClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTable:startingRow:rowSpan:startingColumn:columnSpan:"), table, row, rowSpan, col, colSpan)
	return NSTextTableBlockFromID(rv)
}

// Returns an initialized text table block.
//
// table: The text table containing this text table block.
//
// row: The table row at which the text table block starts.
//
// rowSpan: How many rows the text table block covers.
//
// col: The table column at which the text table block starts.
//
// colSpan: How many columns the text table block covers.
//
// # Discussion
//
// This is the designated initializer.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextTableBlock/init(table:startingRow:rowSpan:startingColumn:columnSpan:)
func (t NSTextTableBlock) InitWithTableStartingRowRowSpanStartingColumnColumnSpan(table INSTextTable, row int, rowSpan int, col int, colSpan int) NSTextTableBlock {
	rv := objc.Send[NSTextTableBlock](t.ID, objc.Sel("initWithTable:startingRow:rowSpan:startingColumn:columnSpan:"), table, row, rowSpan, col, colSpan)
	return rv
}

// Returns the table containing this text table block.
//
// # Return Value
//
// The table containing this text table block.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextTableBlock/table
func (t NSTextTableBlock) Table() INSTextTable {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("table"))
	return NSTextTableFromID(objc.ID(rv))
}

// Returns the table row at which this text table block starts.
//
// # Return Value
//
// The table row at which this text table block starts.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextTableBlock/startingRow
func (t NSTextTableBlock) StartingRow() int {
	rv := objc.Send[int](t.ID, objc.Sel("startingRow"))
	return rv
}

// Returns the number of table rows spanned by this text table block.
//
// # Return Value
//
// The number of table rows spanned by this text table block.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextTableBlock/rowSpan
func (t NSTextTableBlock) RowSpan() int {
	rv := objc.Send[int](t.ID, objc.Sel("rowSpan"))
	return rv
}

// Returns the table column at which this text table block starts.
//
// # Return Value
//
// The table column at which this text table block starts.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextTableBlock/startingColumn
func (t NSTextTableBlock) StartingColumn() int {
	rv := objc.Send[int](t.ID, objc.Sel("startingColumn"))
	return rv
}

// Returns the number of table columns spanned by this text table block.
//
// # Return Value
//
// The number of table columns spanned by this text table block.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextTableBlock/columnSpan
func (t NSTextTableBlock) ColumnSpan() int {
	rv := objc.Send[int](t.ID, objc.Sel("columnSpan"))
	return rv
}
