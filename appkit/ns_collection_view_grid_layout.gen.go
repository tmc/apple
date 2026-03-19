// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSCollectionViewGridLayout] class.
var (
	_NSCollectionViewGridLayoutClass     NSCollectionViewGridLayoutClass
	_NSCollectionViewGridLayoutClassOnce sync.Once
)

func getNSCollectionViewGridLayoutClass() NSCollectionViewGridLayoutClass {
	_NSCollectionViewGridLayoutClassOnce.Do(func() {
		_NSCollectionViewGridLayoutClass = NSCollectionViewGridLayoutClass{class: objc.GetClass("NSCollectionViewGridLayout")}
	})
	return _NSCollectionViewGridLayoutClass
}

// GetNSCollectionViewGridLayoutClass returns the class object for NSCollectionViewGridLayout.
func GetNSCollectionViewGridLayoutClass() NSCollectionViewGridLayoutClass {
	return getNSCollectionViewGridLayoutClass()
}

type NSCollectionViewGridLayoutClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSCollectionViewGridLayoutClass) Alloc() NSCollectionViewGridLayout {
	rv := objc.Send[NSCollectionViewGridLayout](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A layout that displays a single section of items in a row and column grid.
//
// # Overview
// 
// The [NSCollectionViewGridLayout] object provides the same layout behavior
// offered by the [NSCollectionView] class prior to macOS 10.11, and you can
// use it in cases where you want to maintain the old appearance while still
// taking advantage of newer collection view features.
// 
// # Configuring a Collection View to Use a Grid Layout
// 
// You can configure a collection view to use a grid layout object
// programmatically or at design time:
// 
// - At design time, set the Layout attribute of your collection view to Grid.
// - Create an [NSCollectionViewGridLayout] object programmatically and assign
// it to the collection view’s [CollectionViewLayout] property.
// 
// A grid layout displays only items and does not display supplementary views
// or decoration views. Use the properties of this class to configure the
// number of rows and columns in the grid. You can also use these properties
// to configure the spacing between items and the minimum sizes.
//
// # Specifying the Grid Parameters
//
//   - [NSCollectionViewGridLayout.MaximumNumberOfRows]: The maximum number of rows to display in the collection view’s visible area.
//   - [NSCollectionViewGridLayout.SetMaximumNumberOfRows]
//   - [NSCollectionViewGridLayout.MaximumNumberOfColumns]: The maximum number of columns to display in the collection view’s visible area.
//   - [NSCollectionViewGridLayout.SetMaximumNumberOfColumns]
//   - [NSCollectionViewGridLayout.MinimumItemSize]: The smallest allowable size for an item’s view.
//   - [NSCollectionViewGridLayout.SetMinimumItemSize]
//   - [NSCollectionViewGridLayout.MaximumItemSize]: The largest allowable size for an item’s view.
//   - [NSCollectionViewGridLayout.SetMaximumItemSize]
//
// # Specifying the Grid Layout Attributes
//
//   - [NSCollectionViewGridLayout.MinimumInteritemSpacing]: The minimum spacing (in points) to use between items in the same row or column.
//   - [NSCollectionViewGridLayout.SetMinimumInteritemSpacing]
//   - [NSCollectionViewGridLayout.MinimumLineSpacing]: The minimum spacing (in points) to use between rows or columns.
//   - [NSCollectionViewGridLayout.SetMinimumLineSpacing]
//   - [NSCollectionViewGridLayout.Margins]: The amount of empty space (in points) around the grid’s content.
//   - [NSCollectionViewGridLayout.SetMargins]
//
// # Specifying the Grid Background Color
//
//   - [NSCollectionViewGridLayout.BackgroundColors]: The array of background colors to use when drawing the grid.
//   - [NSCollectionViewGridLayout.SetBackgroundColors]
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewGridLayout
type NSCollectionViewGridLayout struct {
	NSCollectionViewLayout
}

// NSCollectionViewGridLayoutFromID constructs a [NSCollectionViewGridLayout] from an objc.ID.
//
// A layout that displays a single section of items in a row and column grid.
func NSCollectionViewGridLayoutFromID(id objc.ID) NSCollectionViewGridLayout {
	return NSCollectionViewGridLayout{NSCollectionViewLayout: NSCollectionViewLayoutFromID(id)}
}
// NOTE: NSCollectionViewGridLayout adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSCollectionViewGridLayout] class.
//
// # Specifying the Grid Parameters
//
//   - [INSCollectionViewGridLayout.MaximumNumberOfRows]: The maximum number of rows to display in the collection view’s visible area.
//   - [INSCollectionViewGridLayout.SetMaximumNumberOfRows]
//   - [INSCollectionViewGridLayout.MaximumNumberOfColumns]: The maximum number of columns to display in the collection view’s visible area.
//   - [INSCollectionViewGridLayout.SetMaximumNumberOfColumns]
//   - [INSCollectionViewGridLayout.MinimumItemSize]: The smallest allowable size for an item’s view.
//   - [INSCollectionViewGridLayout.SetMinimumItemSize]
//   - [INSCollectionViewGridLayout.MaximumItemSize]: The largest allowable size for an item’s view.
//   - [INSCollectionViewGridLayout.SetMaximumItemSize]
//
// # Specifying the Grid Layout Attributes
//
//   - [INSCollectionViewGridLayout.MinimumInteritemSpacing]: The minimum spacing (in points) to use between items in the same row or column.
//   - [INSCollectionViewGridLayout.SetMinimumInteritemSpacing]
//   - [INSCollectionViewGridLayout.MinimumLineSpacing]: The minimum spacing (in points) to use between rows or columns.
//   - [INSCollectionViewGridLayout.SetMinimumLineSpacing]
//   - [INSCollectionViewGridLayout.Margins]: The amount of empty space (in points) around the grid’s content.
//   - [INSCollectionViewGridLayout.SetMargins]
//
// # Specifying the Grid Background Color
//
//   - [INSCollectionViewGridLayout.BackgroundColors]: The array of background colors to use when drawing the grid.
//   - [INSCollectionViewGridLayout.SetBackgroundColors]
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewGridLayout
type INSCollectionViewGridLayout interface {
	INSCollectionViewLayout

	// Topic: Specifying the Grid Parameters

	// The maximum number of rows to display in the collection view’s visible area.
	MaximumNumberOfRows() uint
	SetMaximumNumberOfRows(value uint)
	// The maximum number of columns to display in the collection view’s visible area.
	MaximumNumberOfColumns() uint
	SetMaximumNumberOfColumns(value uint)
	// The smallest allowable size for an item’s view.
	MinimumItemSize() corefoundation.CGSize
	SetMinimumItemSize(value corefoundation.CGSize)
	// The largest allowable size for an item’s view.
	MaximumItemSize() corefoundation.CGSize
	SetMaximumItemSize(value corefoundation.CGSize)

	// Topic: Specifying the Grid Layout Attributes

	// The minimum spacing (in points) to use between items in the same row or column.
	MinimumInteritemSpacing() float64
	SetMinimumInteritemSpacing(value float64)
	// The minimum spacing (in points) to use between rows or columns.
	MinimumLineSpacing() float64
	SetMinimumLineSpacing(value float64)
	// The amount of empty space (in points) around the grid’s content.
	Margins() foundation.NSEdgeInsets
	SetMargins(value foundation.NSEdgeInsets)

	// Topic: Specifying the Grid Background Color

	// The array of background colors to use when drawing the grid.
	BackgroundColors() []NSColor
	SetBackgroundColors(value []NSColor)
}

// Init initializes the instance.
func (c NSCollectionViewGridLayout) Init() NSCollectionViewGridLayout {
	rv := objc.Send[NSCollectionViewGridLayout](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c NSCollectionViewGridLayout) Autorelease() NSCollectionViewGridLayout {
	rv := objc.Send[NSCollectionViewGridLayout](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSCollectionViewGridLayout creates a new NSCollectionViewGridLayout instance.
func NewNSCollectionViewGridLayout() NSCollectionViewGridLayout {
	class := getNSCollectionViewGridLayoutClass()
	rv := objc.Send[NSCollectionViewGridLayout](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The maximum number of rows to display in the collection view’s visible
// area.
//
// # Discussion
// 
// Use this value to specify the maximum number of rows to display in the
// collection view at any given time. The grid layout object uses this value
// during layout to configure the position and spacing of items. The default
// value of this property is `0`, which means that there is no maximum number
// of rows.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewGridLayout/maximumNumberOfRows
func (c NSCollectionViewGridLayout) MaximumNumberOfRows() uint {
	rv := objc.Send[uint](c.ID, objc.Sel("maximumNumberOfRows"))
	return rv
}
func (c NSCollectionViewGridLayout) SetMaximumNumberOfRows(value uint) {
	objc.Send[struct{}](c.ID, objc.Sel("setMaximumNumberOfRows:"), value)
}
// The maximum number of columns to display in the collection view’s visible
// area.
//
// # Discussion
// 
// Use this value to specify the maximum number of columns that should be
// visible in the collection view at any given time. The grid layout object
// uses this value during layout to configure the position and spacing of
// items. The default value of this property is `0`, which means that there is
// no maximum number of columns.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewGridLayout/maximumNumberOfColumns
func (c NSCollectionViewGridLayout) MaximumNumberOfColumns() uint {
	rv := objc.Send[uint](c.ID, objc.Sel("maximumNumberOfColumns"))
	return rv
}
func (c NSCollectionViewGridLayout) SetMaximumNumberOfColumns(value uint) {
	objc.Send[struct{}](c.ID, objc.Sel("setMaximumNumberOfColumns:"), value)
}
// The smallest allowable size for an item’s view.
//
// # Discussion
// 
// Use this property to ensure that items have a minimum size when displayed
// in the grid. The default value of this property is (`0.0`, `0.0`), which
// imposes no minimum size for items.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewGridLayout/minimumItemSize
func (c NSCollectionViewGridLayout) MinimumItemSize() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](c.ID, objc.Sel("minimumItemSize"))
	return corefoundation.CGSize(rv)
}
func (c NSCollectionViewGridLayout) SetMinimumItemSize(value corefoundation.CGSize) {
	objc.Send[struct{}](c.ID, objc.Sel("setMinimumItemSize:"), value)
}
// The largest allowable size for an item’s view.
//
// # Discussion
// 
// Use this property to limit the maximum size of items displayed in the grid.
// The default value of this property is (`0.0`, `0.0`), which imposes no
// maximum size for items.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewGridLayout/maximumItemSize
func (c NSCollectionViewGridLayout) MaximumItemSize() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](c.ID, objc.Sel("maximumItemSize"))
	return corefoundation.CGSize(rv)
}
func (c NSCollectionViewGridLayout) SetMaximumItemSize(value corefoundation.CGSize) {
	objc.Send[struct{}](c.ID, objc.Sel("setMaximumItemSize:"), value)
}
// The minimum spacing (in points) to use between items in the same row or
// column.
//
// # Discussion
// 
// For a vertically scrolling layout, the value represents the minimum spacing
// between items in the same row. For a horizontally scrolling layout, the
// value represents the minimum spacing between items in the same column. The
// layout object uses this spacing only to compute how many items can fit in a
// single row or column. The actual spacing may be increased after the number
// of items has been determined.
// 
// The default value of this property is `0.0`.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewGridLayout/minimumInteritemSpacing
func (c NSCollectionViewGridLayout) MinimumInteritemSpacing() float64 {
	rv := objc.Send[float64](c.ID, objc.Sel("minimumInteritemSpacing"))
	return rv
}
func (c NSCollectionViewGridLayout) SetMinimumInteritemSpacing(value float64) {
	objc.Send[struct{}](c.ID, objc.Sel("setMinimumInteritemSpacing:"), value)
}
// The minimum spacing (in points) to use between rows or columns.
//
// # Discussion
// 
// For a vertically scrolling layout, the value represents the minimum spacing
// between successive rows. For a horizontally scrolling layout, the value
// represents the minimum spacing between successive columns. This spacing is
// not applied to the space between the header view and the first line or
// between the last line and the footer view.
// 
// The default value of this property is `0.0`.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewGridLayout/minimumLineSpacing
func (c NSCollectionViewGridLayout) MinimumLineSpacing() float64 {
	rv := objc.Send[float64](c.ID, objc.Sel("minimumLineSpacing"))
	return rv
}
func (c NSCollectionViewGridLayout) SetMinimumLineSpacing(value float64) {
	objc.Send[struct{}](c.ID, objc.Sel("setMinimumLineSpacing:"), value)
}
// The amount of empty space (in points) around the grid’s content.
//
// # Discussion
// 
// The default value of this property is [NSEdgeInsetsZero]. Changing this
// property to a new value invalidates the layout.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewGridLayout/margins
func (c NSCollectionViewGridLayout) Margins() foundation.NSEdgeInsets {
	rv := objc.Send[foundation.NSEdgeInsets](c.ID, objc.Sel("margins"))
	return foundation.NSEdgeInsets(rv)
}
func (c NSCollectionViewGridLayout) SetMargins(value foundation.NSEdgeInsets) {
	objc.Send[struct{}](c.ID, objc.Sel("setMargins:"), value)
}
// The array of background colors to use when drawing the grid.
//
// # Discussion
// 
// The [NSColor] objects in this property are used to draw the grid’s
// background. The appearance of the background depends on the value you
// specify:
// 
// - Specifying `nil` fills the background with the collection view’s
// default background color. - Specifying an empty array causes the collection
// view to draw no background color. - Specifying an array with one color
// object fills the background with the specified color. - Specifying an array
// with more than one color object causes the collection view to use the
// specified colors to create a checkerboard pattern. Each successive grid
// item is displayed with the next color in the array, cycling back to the
// beginning of the array when the last color is reached.
// 
// The default value of this property is `nil`.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewGridLayout/backgroundColors
func (c NSCollectionViewGridLayout) BackgroundColors() []NSColor {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("backgroundColors"))
	return objc.ConvertSlice(rv, func(id objc.ID) NSColor {
		return NSColorFromID(id)
	})
}
func (c NSCollectionViewGridLayout) SetBackgroundColors(value []NSColor) {
	objc.Send[struct{}](c.ID, objc.Sel("setBackgroundColors:"), objectivec.IObjectSliceToNSArray(value))
}

