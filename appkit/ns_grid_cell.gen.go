// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSGridCell] class.
var (
	_NSGridCellClass     NSGridCellClass
	_NSGridCellClassOnce sync.Once
)

func getNSGridCellClass() NSGridCellClass {
	_NSGridCellClassOnce.Do(func() {
		_NSGridCellClass = NSGridCellClass{class: objc.GetClass("NSGridCell")}
	})
	return _NSGridCellClass
}

// GetNSGridCellClass returns the class object for NSGridCell.
func GetNSGridCellClass() NSGridCellClass {
	return getNSGridCellClass()
}

type NSGridCellClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSGridCellClass) Alloc() NSGridCell {
	rv := objc.Send[NSGridCell](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// An individual content area within a grid view, typically at the
// intersection of a row and a column.
//
// # Overview
// 
// Use a grid cell to specify the content view to display and to position the
// content view within the cell’s area.
//
// # Getting the Cell Containers
//
//   - [NSGridCell.Column]
//   - [NSGridCell.Row]
//   - [NSGridCell.ContentView]
//   - [NSGridCell.SetContentView]
//
// # Formatting the Cell
//
//   - [NSGridCell.CustomPlacementConstraints]
//   - [NSGridCell.SetCustomPlacementConstraints]
//   - [NSGridCell.RowAlignment]
//   - [NSGridCell.SetRowAlignment]
//   - [NSGridCell.XPlacement]
//   - [NSGridCell.SetXPlacement]
//   - [NSGridCell.YPlacement]
//   - [NSGridCell.SetYPlacement]
//
// See: https://developer.apple.com/documentation/AppKit/NSGridCell
type NSGridCell struct {
	objectivec.Object
}

// NSGridCellFromID constructs a [NSGridCell] from an objc.ID.
//
// An individual content area within a grid view, typically at the
// intersection of a row and a column.
func NSGridCellFromID(id objc.ID) NSGridCell {
	return NSGridCell{objectivec.Object{ID: id}}
}
// NOTE: NSGridCell adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSGridCell] class.
//
// # Getting the Cell Containers
//
//   - [INSGridCell.Column]
//   - [INSGridCell.Row]
//   - [INSGridCell.ContentView]
//   - [INSGridCell.SetContentView]
//
// # Formatting the Cell
//
//   - [INSGridCell.CustomPlacementConstraints]
//   - [INSGridCell.SetCustomPlacementConstraints]
//   - [INSGridCell.RowAlignment]
//   - [INSGridCell.SetRowAlignment]
//   - [INSGridCell.XPlacement]
//   - [INSGridCell.SetXPlacement]
//   - [INSGridCell.YPlacement]
//   - [INSGridCell.SetYPlacement]
//
// See: https://developer.apple.com/documentation/AppKit/NSGridCell
type INSGridCell interface {
	objectivec.IObject

	// Topic: Getting the Cell Containers

	Column() INSGridColumn
	Row() INSGridRow
	ContentView() INSView
	SetContentView(value INSView)

	// Topic: Formatting the Cell

	CustomPlacementConstraints() []NSLayoutConstraint
	SetCustomPlacementConstraints(value []NSLayoutConstraint)
	RowAlignment() NSGridRowAlignment
	SetRowAlignment(value NSGridRowAlignment)
	XPlacement() NSGridCellPlacement
	SetXPlacement(value NSGridCellPlacement)
	YPlacement() NSGridCellPlacement
	SetYPlacement(value NSGridCellPlacement)

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (g NSGridCell) Init() NSGridCell {
	rv := objc.Send[NSGridCell](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g NSGridCell) Autorelease() NSGridCell {
	rv := objc.Send[NSGridCell](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSGridCell creates a new NSGridCell instance.
func NewNSGridCell() NSGridCell {
	class := getNSGridCellClass()
	rv := objc.Send[NSGridCell](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (g NSGridCell) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](g.ID, objc.Sel("encodeWithCoder:"), coder)
}

// See: https://developer.apple.com/documentation/AppKit/NSGridCell/column
func (g NSGridCell) Column() INSGridColumn {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("column"))
	return NSGridColumnFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/AppKit/NSGridCell/row
func (g NSGridCell) Row() INSGridRow {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("row"))
	return NSGridRowFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/AppKit/NSGridCell/contentView
func (g NSGridCell) ContentView() INSView {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("contentView"))
	return NSViewFromID(objc.ID(rv))
}
func (g NSGridCell) SetContentView(value INSView) {
	objc.Send[struct{}](g.ID, objc.Sel("setContentView:"), value)
}
// See: https://developer.apple.com/documentation/AppKit/NSGridCell/customPlacementConstraints
func (g NSGridCell) CustomPlacementConstraints() []NSLayoutConstraint {
	rv := objc.Send[[]objc.ID](g.ID, objc.Sel("customPlacementConstraints"))
	return objc.ConvertSlice(rv, func(id objc.ID) NSLayoutConstraint {
		return NSLayoutConstraintFromID(id)
	})
}
func (g NSGridCell) SetCustomPlacementConstraints(value []NSLayoutConstraint) {
	objc.Send[struct{}](g.ID, objc.Sel("setCustomPlacementConstraints:"), objectivec.IObjectSliceToNSArray(value))
}
// See: https://developer.apple.com/documentation/AppKit/NSGridCell/rowAlignment
func (g NSGridCell) RowAlignment() NSGridRowAlignment {
	rv := objc.Send[NSGridRowAlignment](g.ID, objc.Sel("rowAlignment"))
	return NSGridRowAlignment(rv)
}
func (g NSGridCell) SetRowAlignment(value NSGridRowAlignment) {
	objc.Send[struct{}](g.ID, objc.Sel("setRowAlignment:"), value)
}
// See: https://developer.apple.com/documentation/AppKit/NSGridCell/xPlacement
func (g NSGridCell) XPlacement() NSGridCellPlacement {
	rv := objc.Send[NSGridCellPlacement](g.ID, objc.Sel("xPlacement"))
	return NSGridCellPlacement(rv)
}
func (g NSGridCell) SetXPlacement(value NSGridCellPlacement) {
	objc.Send[struct{}](g.ID, objc.Sel("setXPlacement:"), value)
}
// See: https://developer.apple.com/documentation/AppKit/NSGridCell/yPlacement
func (g NSGridCell) YPlacement() NSGridCellPlacement {
	rv := objc.Send[NSGridCellPlacement](g.ID, objc.Sel("yPlacement"))
	return NSGridCellPlacement(rv)
}
func (g NSGridCell) SetYPlacement(value NSGridCellPlacement) {
	objc.Send[struct{}](g.ID, objc.Sel("setYPlacement:"), value)
}

// See: https://developer.apple.com/documentation/AppKit/NSGridCell/emptyContentView
func (_NSGridCellClass NSGridCellClass) EmptyContentView() NSView {
	rv := objc.Send[objc.ID](objc.ID(_NSGridCellClass.class), objc.Sel("emptyContentView"))
	return NSViewFromID(objc.ID(rv))
}

