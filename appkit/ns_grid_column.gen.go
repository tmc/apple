// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSGridColumn] class.
var (
	_NSGridColumnClass     NSGridColumnClass
	_NSGridColumnClassOnce sync.Once
)

func getNSGridColumnClass() NSGridColumnClass {
	_NSGridColumnClassOnce.Do(func() {
		_NSGridColumnClass = NSGridColumnClass{class: objc.GetClass("NSGridColumn")}
	})
	return _NSGridColumnClass
}

// GetNSGridColumnClass returns the class object for NSGridColumn.
func GetNSGridColumnClass() NSGridColumnClass {
	return getNSGridColumnClass()
}

type NSGridColumnClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSGridColumnClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSGridColumnClass) Alloc() NSGridColumn {
	rv := objc.Send[NSGridColumn](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A column within a grid view.
//
// # Instance Properties
//
//   - [NSGridColumn.GridView]
//   - [NSGridColumn.Hidden]
//   - [NSGridColumn.SetHidden]
//   - [NSGridColumn.LeadingPadding]
//   - [NSGridColumn.SetLeadingPadding]
//   - [NSGridColumn.NumberOfCells]
//   - [NSGridColumn.TrailingPadding]
//   - [NSGridColumn.SetTrailingPadding]
//   - [NSGridColumn.Width]
//   - [NSGridColumn.SetWidth]
//   - [NSGridColumn.XPlacement]
//   - [NSGridColumn.SetXPlacement]
//
// # Instance Methods
//
//   - [NSGridColumn.CellAtIndex]
//   - [NSGridColumn.MergeCellsInRange]
//
// See: https://developer.apple.com/documentation/AppKit/NSGridColumn
type NSGridColumn struct {
	objectivec.Object
}

// NSGridColumnFromID constructs a [NSGridColumn] from an objc.ID.
//
// A column within a grid view.
func NSGridColumnFromID(id objc.ID) NSGridColumn {
	return NSGridColumn{objectivec.Object{ID: id}}
}

// NOTE: NSGridColumn adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSGridColumn] class.
//
// # Instance Properties
//
//   - [INSGridColumn.GridView]
//   - [INSGridColumn.Hidden]
//   - [INSGridColumn.SetHidden]
//   - [INSGridColumn.LeadingPadding]
//   - [INSGridColumn.SetLeadingPadding]
//   - [INSGridColumn.NumberOfCells]
//   - [INSGridColumn.TrailingPadding]
//   - [INSGridColumn.SetTrailingPadding]
//   - [INSGridColumn.Width]
//   - [INSGridColumn.SetWidth]
//   - [INSGridColumn.XPlacement]
//   - [INSGridColumn.SetXPlacement]
//
// # Instance Methods
//
//   - [INSGridColumn.CellAtIndex]
//   - [INSGridColumn.MergeCellsInRange]
//
// See: https://developer.apple.com/documentation/AppKit/NSGridColumn
type INSGridColumn interface {
	objectivec.IObject

	// Topic: Instance Properties

	GridView() INSGridView
	Hidden() bool
	SetHidden(value bool)
	LeadingPadding() float64
	SetLeadingPadding(value float64)
	NumberOfCells() int
	TrailingPadding() float64
	SetTrailingPadding(value float64)
	Width() float64
	SetWidth(value float64)
	XPlacement() NSGridCellPlacement
	SetXPlacement(value NSGridCellPlacement)

	// Topic: Instance Methods

	CellAtIndex(index int) INSGridCell
	MergeCellsInRange(range_ foundation.NSRange)

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (g NSGridColumn) Init() NSGridColumn {
	rv := objc.Send[NSGridColumn](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g NSGridColumn) Autorelease() NSGridColumn {
	rv := objc.Send[NSGridColumn](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSGridColumn creates a new NSGridColumn instance.
func NewNSGridColumn() NSGridColumn {
	class := getNSGridColumnClass()
	rv := objc.Send[NSGridColumn](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/AppKit/NSGridColumn/cell(at:)
func (g NSGridColumn) CellAtIndex(index int) INSGridCell {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("cellAtIndex:"), index)
	return NSGridCellFromID(rv)
}

// See: https://developer.apple.com/documentation/AppKit/NSGridColumn/mergeCells(in:)
func (g NSGridColumn) MergeCellsInRange(range_ foundation.NSRange) {
	objc.Send[objc.ID](g.ID, objc.Sel("mergeCellsInRange:"), range_)
}
func (g NSGridColumn) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](g.ID, objc.Sel("encodeWithCoder:"), coder)
}

// See: https://developer.apple.com/documentation/AppKit/NSGridColumn/gridView
func (g NSGridColumn) GridView() INSGridView {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("gridView"))
	return NSGridViewFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/AppKit/NSGridColumn/isHidden
func (g NSGridColumn) Hidden() bool {
	rv := objc.Send[bool](g.ID, objc.Sel("isHidden"))
	return rv
}
func (g NSGridColumn) SetHidden(value bool) {
	objc.Send[struct{}](g.ID, objc.Sel("setHidden:"), value)
}

// See: https://developer.apple.com/documentation/AppKit/NSGridColumn/leadingPadding
func (g NSGridColumn) LeadingPadding() float64 {
	rv := objc.Send[float64](g.ID, objc.Sel("leadingPadding"))
	return rv
}
func (g NSGridColumn) SetLeadingPadding(value float64) {
	objc.Send[struct{}](g.ID, objc.Sel("setLeadingPadding:"), value)
}

// See: https://developer.apple.com/documentation/AppKit/NSGridColumn/numberOfCells
func (g NSGridColumn) NumberOfCells() int {
	rv := objc.Send[int](g.ID, objc.Sel("numberOfCells"))
	return rv
}

// See: https://developer.apple.com/documentation/AppKit/NSGridColumn/trailingPadding
func (g NSGridColumn) TrailingPadding() float64 {
	rv := objc.Send[float64](g.ID, objc.Sel("trailingPadding"))
	return rv
}
func (g NSGridColumn) SetTrailingPadding(value float64) {
	objc.Send[struct{}](g.ID, objc.Sel("setTrailingPadding:"), value)
}

// See: https://developer.apple.com/documentation/AppKit/NSGridColumn/width
func (g NSGridColumn) Width() float64 {
	rv := objc.Send[float64](g.ID, objc.Sel("width"))
	return rv
}
func (g NSGridColumn) SetWidth(value float64) {
	objc.Send[struct{}](g.ID, objc.Sel("setWidth:"), value)
}

// See: https://developer.apple.com/documentation/AppKit/NSGridColumn/xPlacement
func (g NSGridColumn) XPlacement() NSGridCellPlacement {
	rv := objc.Send[NSGridCellPlacement](g.ID, objc.Sel("xPlacement"))
	return NSGridCellPlacement(rv)
}
func (g NSGridColumn) SetXPlacement(value NSGridCellPlacement) {
	objc.Send[struct{}](g.ID, objc.Sel("setXPlacement:"), value)
}
