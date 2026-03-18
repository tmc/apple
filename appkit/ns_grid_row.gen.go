// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSGridRow] class.
var (
	_NSGridRowClass     NSGridRowClass
	_NSGridRowClassOnce sync.Once
)

func getNSGridRowClass() NSGridRowClass {
	_NSGridRowClassOnce.Do(func() {
		_NSGridRowClass = NSGridRowClass{class: objc.GetClass("NSGridRow")}
	})
	return _NSGridRowClass
}

// GetNSGridRowClass returns the class object for NSGridRow.
func GetNSGridRowClass() NSGridRowClass {
	return getNSGridRowClass()
}

type NSGridRowClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSGridRowClass) Alloc() NSGridRow {
	rv := objc.Send[NSGridRow](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A row within a grid view.
//
// # Getting the Row Details
//
//   - [NSGridRow.NumberOfCells]
//   - [NSGridRow.Hidden]
//   - [NSGridRow.SetHidden]
//
// # Formatting the Row
//
//   - [NSGridRow.TopPadding]
//   - [NSGridRow.SetTopPadding]
//   - [NSGridRow.BottomPadding]
//   - [NSGridRow.SetBottomPadding]
//   - [NSGridRow.Height]
//   - [NSGridRow.SetHeight]
//   - [NSGridRow.RowAlignment]
//   - [NSGridRow.SetRowAlignment]
//   - [NSGridRow.YPlacement]
//   - [NSGridRow.SetYPlacement]
//
// # Getting the Grid View
//
//   - [NSGridRow.GridView]
//
// # Getting Cells
//
//   - [NSGridRow.CellAtIndex]
//
// # Merging Cells in the Row
//
//   - [NSGridRow.MergeCellsInRange]
//
// See: https://developer.apple.com/documentation/AppKit/NSGridRow
type NSGridRow struct {
	objectivec.Object
}

// NSGridRowFromID constructs a [NSGridRow] from an objc.ID.
//
// A row within a grid view.
func NSGridRowFromID(id objc.ID) NSGridRow {
	return NSGridRow{objectivec.Object{ID: id}}
}
// NOTE: NSGridRow adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSGridRow] class.
//
// # Getting the Row Details
//
//   - [INSGridRow.NumberOfCells]
//   - [INSGridRow.Hidden]
//   - [INSGridRow.SetHidden]
//
// # Formatting the Row
//
//   - [INSGridRow.TopPadding]
//   - [INSGridRow.SetTopPadding]
//   - [INSGridRow.BottomPadding]
//   - [INSGridRow.SetBottomPadding]
//   - [INSGridRow.Height]
//   - [INSGridRow.SetHeight]
//   - [INSGridRow.RowAlignment]
//   - [INSGridRow.SetRowAlignment]
//   - [INSGridRow.YPlacement]
//   - [INSGridRow.SetYPlacement]
//
// # Getting the Grid View
//
//   - [INSGridRow.GridView]
//
// # Getting Cells
//
//   - [INSGridRow.CellAtIndex]
//
// # Merging Cells in the Row
//
//   - [INSGridRow.MergeCellsInRange]
//
// See: https://developer.apple.com/documentation/AppKit/NSGridRow
type INSGridRow interface {
	objectivec.IObject

	// Topic: Getting the Row Details

	NumberOfCells() int
	Hidden() bool
	SetHidden(value bool)

	// Topic: Formatting the Row

	TopPadding() float64
	SetTopPadding(value float64)
	BottomPadding() float64
	SetBottomPadding(value float64)
	Height() float64
	SetHeight(value float64)
	RowAlignment() NSGridRowAlignment
	SetRowAlignment(value NSGridRowAlignment)
	YPlacement() NSGridCellPlacement
	SetYPlacement(value NSGridCellPlacement)

	// Topic: Getting the Grid View

	GridView() INSGridView

	// Topic: Getting Cells

	CellAtIndex(index int) INSGridCell

	// Topic: Merging Cells in the Row

	MergeCellsInRange(range_ foundation.NSRange)

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (g NSGridRow) Init() NSGridRow {
	rv := objc.Send[NSGridRow](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g NSGridRow) Autorelease() NSGridRow {
	rv := objc.Send[NSGridRow](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSGridRow creates a new NSGridRow instance.
func NewNSGridRow() NSGridRow {
	class := getNSGridRowClass()
	rv := objc.Send[NSGridRow](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/AppKit/NSGridRow/cell(at:)
func (g NSGridRow) CellAtIndex(index int) INSGridCell {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("cellAtIndex:"), index)
	return NSGridCellFromID(rv)
}

//
// See: https://developer.apple.com/documentation/AppKit/NSGridRow/mergeCells(in:)
func (g NSGridRow) MergeCellsInRange(range_ foundation.NSRange) {
	objc.Send[objc.ID](g.ID, objc.Sel("mergeCellsInRange:"), range_)
}
func (g NSGridRow) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](g.ID, objc.Sel("encodeWithCoder:"), coder)
}

// See: https://developer.apple.com/documentation/AppKit/NSGridRow/numberOfCells
func (g NSGridRow) NumberOfCells() int {
	rv := objc.Send[int](g.ID, objc.Sel("numberOfCells"))
	return rv
}

// See: https://developer.apple.com/documentation/AppKit/NSGridRow/isHidden
func (g NSGridRow) Hidden() bool {
	rv := objc.Send[bool](g.ID, objc.Sel("isHidden"))
	return rv
}
func (g NSGridRow) SetHidden(value bool) {
	objc.Send[struct{}](g.ID, objc.Sel("setHidden:"), value)
}

// See: https://developer.apple.com/documentation/AppKit/NSGridRow/topPadding
func (g NSGridRow) TopPadding() float64 {
	rv := objc.Send[float64](g.ID, objc.Sel("topPadding"))
	return rv
}
func (g NSGridRow) SetTopPadding(value float64) {
	objc.Send[struct{}](g.ID, objc.Sel("setTopPadding:"), value)
}

// See: https://developer.apple.com/documentation/AppKit/NSGridRow/bottomPadding
func (g NSGridRow) BottomPadding() float64 {
	rv := objc.Send[float64](g.ID, objc.Sel("bottomPadding"))
	return rv
}
func (g NSGridRow) SetBottomPadding(value float64) {
	objc.Send[struct{}](g.ID, objc.Sel("setBottomPadding:"), value)
}

// See: https://developer.apple.com/documentation/AppKit/NSGridRow/height
func (g NSGridRow) Height() float64 {
	rv := objc.Send[float64](g.ID, objc.Sel("height"))
	return rv
}
func (g NSGridRow) SetHeight(value float64) {
	objc.Send[struct{}](g.ID, objc.Sel("setHeight:"), value)
}

// See: https://developer.apple.com/documentation/AppKit/NSGridRow/rowAlignment
func (g NSGridRow) RowAlignment() NSGridRowAlignment {
	rv := objc.Send[NSGridRowAlignment](g.ID, objc.Sel("rowAlignment"))
	return NSGridRowAlignment(rv)
}
func (g NSGridRow) SetRowAlignment(value NSGridRowAlignment) {
	objc.Send[struct{}](g.ID, objc.Sel("setRowAlignment:"), value)
}

// See: https://developer.apple.com/documentation/AppKit/NSGridRow/yPlacement
func (g NSGridRow) YPlacement() NSGridCellPlacement {
	rv := objc.Send[NSGridCellPlacement](g.ID, objc.Sel("yPlacement"))
	return NSGridCellPlacement(rv)
}
func (g NSGridRow) SetYPlacement(value NSGridCellPlacement) {
	objc.Send[struct{}](g.ID, objc.Sel("setYPlacement:"), value)
}

// See: https://developer.apple.com/documentation/AppKit/NSGridRow/gridView
func (g NSGridRow) GridView() INSGridView {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("gridView"))
	return NSGridViewFromID(objc.ID(rv))
}

