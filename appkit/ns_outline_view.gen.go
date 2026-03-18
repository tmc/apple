// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSOutlineView] class.
var (
	_NSOutlineViewClass     NSOutlineViewClass
	_NSOutlineViewClassOnce sync.Once
)

func getNSOutlineViewClass() NSOutlineViewClass {
	_NSOutlineViewClassOnce.Do(func() {
		_NSOutlineViewClass = NSOutlineViewClass{class: objc.GetClass("NSOutlineView")}
	})
	return _NSOutlineViewClass
}

// GetNSOutlineViewClass returns the class object for NSOutlineView.
func GetNSOutlineViewClass() NSOutlineViewClass {
	return getNSOutlineViewClass()
}

type NSOutlineViewClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSOutlineViewClass) Alloc() NSOutlineView {
	rv := objc.Send[NSOutlineView](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A view that uses a row-and-column format to display hierarchical data like
// directories and files that can be expanded and collapsed.
//
// # Overview
// 
// Like a table view, an outline view does not store its own data, instead it
// retrieves data values as needed from a data source to which it has a weak
// reference (see [Delegates and Data Sources]). See
// [NSOutlineViewDataSource], which declares the methods that an
// [NSOutlineView] object uses to access the contents of its data source
// object.
// 
// An outline view has the following features:
// 
// - A user can expand and collapse rows, edit values, and resize and
// rearrange columns. - Each item in the outline view must be unique. In order
// for the collapsed state to remain consistent between reloads the item’s
// pointer must remain the same and the item must maintain [isEqual(_:)]
// sameness. - The view gets data from a data source (see
// [NSOutlineViewDataSource]). - The view retrieves only the data that needs
// to be displayed.
// 
// For more information about using NSOutlineView in your app, see [Navigating
// Hierarchical Data Using Outline and Split Views].
// 
// # Subclassing
// 
// Subclassing [NSOutlineView] is not recommended. Customization can be
// accomplished in your data source class implementation (conforming to
// [NSOutlineViewDataSource]) or your delegate class implementation
// (conforming to [NSOutlineViewDelegate]).
//
// [Delegates and Data Sources]: https://developer.apple.com/library/archive/documentation/General/Conceptual/CocoaEncyclopedia/DelegatesandDataSources/DelegatesandDataSources.html#//apple_ref/doc/uid/TP40010810-CH11
// [Navigating Hierarchical Data Using Outline and Split Views]: https://developer.apple.com/documentation/AppKit/navigating-hierarchical-data-using-outline-and-split-views
// [isEqual(_:)]: https://developer.apple.com/documentation/ObjectiveC/NSObjectProtocol/isEqual(_:)
//
// # Accessing the Data Source
//
//   - [NSOutlineView.StronglyReferencesItems]: A Boolean value that indicates whether the outline view retains and releases the objects returned from its data source.
//   - [NSOutlineView.SetStronglyReferencesItems]
//
// # Working with Expandability
//
//   - [NSOutlineView.IsExpandable]: Returns a Boolean value that indicates whether a given item is expandable.
//   - [NSOutlineView.IsItemExpanded]: Returns a Boolean value that indicates whether a given item is expanded.
//
// # Expanding and Collapsing the Outline
//
//   - [NSOutlineView.ExpandItem]: Expands a given item.
//   - [NSOutlineView.ExpandItemExpandChildren]: Expands a specified item and, optionally, its children.
//   - [NSOutlineView.CollapseItem]: Collapses a given item.
//   - [NSOutlineView.CollapseItemCollapseChildren]: Collapses a given item and, optionally, its children.
//
// # Redisplaying Information
//
//   - [NSOutlineView.ReloadItem]: Reloads and redisplays the data for the given item.
//   - [NSOutlineView.ReloadItemReloadChildren]: Reloads a given item and, optionally, its children.
//
// # Converting Between Items and Rows
//
//   - [NSOutlineView.ItemAtRow]: Returns the item associated with a given row.
//   - [NSOutlineView.RowForItem]: Returns the row associated with a given item.
//
// # Working with the Outline Column
//
//   - [NSOutlineView.OutlineTableColumn]: The table column in which hierarchical data is displayed.
//   - [NSOutlineView.SetOutlineTableColumn]
//   - [NSOutlineView.AutoresizesOutlineColumn]: A Boolean value that indicates whether the outline view resizes its outline column when the user expands or collapses items.
//   - [NSOutlineView.SetAutoresizesOutlineColumn]
//
// # Working with Indentation
//
//   - [NSOutlineView.LevelForItem]: Returns the indentation level for a given item.
//   - [NSOutlineView.LevelForRow]: Returns the indentation level for a given row.
//   - [NSOutlineView.IndentationPerLevel]: The per-level indentation, measured in points.
//   - [NSOutlineView.SetIndentationPerLevel]
//   - [NSOutlineView.IndentationMarkerFollowsCell]: A Boolean value indicating whether the indentation marker symbol displayed in the outline column should be indented along with the cell contents.
//   - [NSOutlineView.SetIndentationMarkerFollowsCell]
//
// # Working with Persistence
//
//   - [NSOutlineView.AutosaveExpandedItems]: A Boolean value indicating whether the expanded items are automatically saved across launches of the app.
//   - [NSOutlineView.SetAutosaveExpandedItems]
//
// # Supporting Drag and Drop
//
//   - [NSOutlineView.SetDropItemDropChildIndex]: Used to “retarget” a proposed drop.
//   - [NSOutlineView.ShouldCollapseAutoExpandedItemsForDeposited]: Returns a Boolean value that indicates whether auto-expanded items should return to their original collapsed state.
//
// # Getting Related Items
//
//   - [NSOutlineView.ParentForItem]: Returns the parent for a given item.
//   - [NSOutlineView.ChildIndexForItem]: Returns the child index of the specified item within its parent.
//   - [NSOutlineView.ChildOfItem]: Returns the specified child of an item.
//   - [NSOutlineView.NumberOfChildrenOfItem]: Returns the number of children for the specified parent item.
//
// # Getting the Frame for a Cell
//
//   - [NSOutlineView.FrameOfOutlineCellAtRow]: Returns the frame of the outline cell for a given row.
//
// # Manipulating Items
//
//   - [NSOutlineView.InsertItemsAtIndexesInParentWithAnimation]: Inserts new items at the given indexes in the given parent with the specified optional animations.
//   - [NSOutlineView.MoveItemAtIndexInParentToIndexInParent]: Moves an item at a given index in the given parent to a new index in a new parent.
//   - [NSOutlineView.RemoveItemsAtIndexesInParentWithAnimation]: Removes items at the given indexes in the given parent with the specified optional animations.
//
// See: https://developer.apple.com/documentation/AppKit/NSOutlineView
type NSOutlineView struct {
	NSTableView
}

// NSOutlineViewFromID constructs a [NSOutlineView] from an objc.ID.
//
// A view that uses a row-and-column format to display hierarchical data like
// directories and files that can be expanded and collapsed.
func NSOutlineViewFromID(id objc.ID) NSOutlineView {
	return NSOutlineView{NSTableView: NSTableViewFromID(id)}
}
// NOTE: NSOutlineView adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSOutlineView] class.
//
// # Accessing the Data Source
//
//   - [INSOutlineView.StronglyReferencesItems]: A Boolean value that indicates whether the outline view retains and releases the objects returned from its data source.
//   - [INSOutlineView.SetStronglyReferencesItems]
//
// # Working with Expandability
//
//   - [INSOutlineView.IsExpandable]: Returns a Boolean value that indicates whether a given item is expandable.
//   - [INSOutlineView.IsItemExpanded]: Returns a Boolean value that indicates whether a given item is expanded.
//
// # Expanding and Collapsing the Outline
//
//   - [INSOutlineView.ExpandItem]: Expands a given item.
//   - [INSOutlineView.ExpandItemExpandChildren]: Expands a specified item and, optionally, its children.
//   - [INSOutlineView.CollapseItem]: Collapses a given item.
//   - [INSOutlineView.CollapseItemCollapseChildren]: Collapses a given item and, optionally, its children.
//
// # Redisplaying Information
//
//   - [INSOutlineView.ReloadItem]: Reloads and redisplays the data for the given item.
//   - [INSOutlineView.ReloadItemReloadChildren]: Reloads a given item and, optionally, its children.
//
// # Converting Between Items and Rows
//
//   - [INSOutlineView.ItemAtRow]: Returns the item associated with a given row.
//   - [INSOutlineView.RowForItem]: Returns the row associated with a given item.
//
// # Working with the Outline Column
//
//   - [INSOutlineView.OutlineTableColumn]: The table column in which hierarchical data is displayed.
//   - [INSOutlineView.SetOutlineTableColumn]
//   - [INSOutlineView.AutoresizesOutlineColumn]: A Boolean value that indicates whether the outline view resizes its outline column when the user expands or collapses items.
//   - [INSOutlineView.SetAutoresizesOutlineColumn]
//
// # Working with Indentation
//
//   - [INSOutlineView.LevelForItem]: Returns the indentation level for a given item.
//   - [INSOutlineView.LevelForRow]: Returns the indentation level for a given row.
//   - [INSOutlineView.IndentationPerLevel]: The per-level indentation, measured in points.
//   - [INSOutlineView.SetIndentationPerLevel]
//   - [INSOutlineView.IndentationMarkerFollowsCell]: A Boolean value indicating whether the indentation marker symbol displayed in the outline column should be indented along with the cell contents.
//   - [INSOutlineView.SetIndentationMarkerFollowsCell]
//
// # Working with Persistence
//
//   - [INSOutlineView.AutosaveExpandedItems]: A Boolean value indicating whether the expanded items are automatically saved across launches of the app.
//   - [INSOutlineView.SetAutosaveExpandedItems]
//
// # Supporting Drag and Drop
//
//   - [INSOutlineView.SetDropItemDropChildIndex]: Used to “retarget” a proposed drop.
//   - [INSOutlineView.ShouldCollapseAutoExpandedItemsForDeposited]: Returns a Boolean value that indicates whether auto-expanded items should return to their original collapsed state.
//
// # Getting Related Items
//
//   - [INSOutlineView.ParentForItem]: Returns the parent for a given item.
//   - [INSOutlineView.ChildIndexForItem]: Returns the child index of the specified item within its parent.
//   - [INSOutlineView.ChildOfItem]: Returns the specified child of an item.
//   - [INSOutlineView.NumberOfChildrenOfItem]: Returns the number of children for the specified parent item.
//
// # Getting the Frame for a Cell
//
//   - [INSOutlineView.FrameOfOutlineCellAtRow]: Returns the frame of the outline cell for a given row.
//
// # Manipulating Items
//
//   - [INSOutlineView.InsertItemsAtIndexesInParentWithAnimation]: Inserts new items at the given indexes in the given parent with the specified optional animations.
//   - [INSOutlineView.MoveItemAtIndexInParentToIndexInParent]: Moves an item at a given index in the given parent to a new index in a new parent.
//   - [INSOutlineView.RemoveItemsAtIndexesInParentWithAnimation]: Removes items at the given indexes in the given parent with the specified optional animations.
//
// See: https://developer.apple.com/documentation/AppKit/NSOutlineView
type INSOutlineView interface {
	INSTableView
	NSAccessibilityOutline

	// Topic: Accessing the Data Source

	// A Boolean value that indicates whether the outline view retains and releases the objects returned from its data source.
	StronglyReferencesItems() bool
	SetStronglyReferencesItems(value bool)

	// Topic: Working with Expandability

	// Returns a Boolean value that indicates whether a given item is expandable.
	IsExpandable(item objectivec.IObject) bool
	// Returns a Boolean value that indicates whether a given item is expanded.
	IsItemExpanded(item objectivec.IObject) bool

	// Topic: Expanding and Collapsing the Outline

	// Expands a given item.
	ExpandItem(item objectivec.IObject)
	// Expands a specified item and, optionally, its children.
	ExpandItemExpandChildren(item objectivec.IObject, expandChildren bool)
	// Collapses a given item.
	CollapseItem(item objectivec.IObject)
	// Collapses a given item and, optionally, its children.
	CollapseItemCollapseChildren(item objectivec.IObject, collapseChildren bool)

	// Topic: Redisplaying Information

	// Reloads and redisplays the data for the given item.
	ReloadItem(item objectivec.IObject)
	// Reloads a given item and, optionally, its children.
	ReloadItemReloadChildren(item objectivec.IObject, reloadChildren bool)

	// Topic: Converting Between Items and Rows

	// Returns the item associated with a given row.
	ItemAtRow(row int) objectivec.IObject
	// Returns the row associated with a given item.
	RowForItem(item objectivec.IObject) int

	// Topic: Working with the Outline Column

	// The table column in which hierarchical data is displayed.
	OutlineTableColumn() INSTableColumn
	SetOutlineTableColumn(value INSTableColumn)
	// A Boolean value that indicates whether the outline view resizes its outline column when the user expands or collapses items.
	AutoresizesOutlineColumn() bool
	SetAutoresizesOutlineColumn(value bool)

	// Topic: Working with Indentation

	// Returns the indentation level for a given item.
	LevelForItem(item objectivec.IObject) int
	// Returns the indentation level for a given row.
	LevelForRow(row int) int
	// The per-level indentation, measured in points.
	IndentationPerLevel() float64
	SetIndentationPerLevel(value float64)
	// A Boolean value indicating whether the indentation marker symbol displayed in the outline column should be indented along with the cell contents.
	IndentationMarkerFollowsCell() bool
	SetIndentationMarkerFollowsCell(value bool)

	// Topic: Working with Persistence

	// A Boolean value indicating whether the expanded items are automatically saved across launches of the app.
	AutosaveExpandedItems() bool
	SetAutosaveExpandedItems(value bool)

	// Topic: Supporting Drag and Drop

	// Used to “retarget” a proposed drop.
	SetDropItemDropChildIndex(item objectivec.IObject, index int)
	// Returns a Boolean value that indicates whether auto-expanded items should return to their original collapsed state.
	ShouldCollapseAutoExpandedItemsForDeposited(deposited bool) bool

	// Topic: Getting Related Items

	// Returns the parent for a given item.
	ParentForItem(item objectivec.IObject) objectivec.IObject
	// Returns the child index of the specified item within its parent.
	ChildIndexForItem(item objectivec.IObject) int
	// Returns the specified child of an item.
	ChildOfItem(index int, item objectivec.IObject) objectivec.IObject
	// Returns the number of children for the specified parent item.
	NumberOfChildrenOfItem(item objectivec.IObject) int

	// Topic: Getting the Frame for a Cell

	// Returns the frame of the outline cell for a given row.
	FrameOfOutlineCellAtRow(row int) corefoundation.CGRect

	// Topic: Manipulating Items

	// Inserts new items at the given indexes in the given parent with the specified optional animations.
	InsertItemsAtIndexesInParentWithAnimation(indexes foundation.NSIndexSet, parent objectivec.IObject, animationOptions NSTableViewAnimationOptions)
	// Moves an item at a given index in the given parent to a new index in a new parent.
	MoveItemAtIndexInParentToIndexInParent(fromIndex int, oldParent objectivec.IObject, toIndex int, newParent objectivec.IObject)
	// Removes items at the given indexes in the given parent with the specified optional animations.
	RemoveItemsAtIndexesInParentWithAnimation(indexes foundation.NSIndexSet, parent objectivec.IObject, animationOptions NSTableViewAnimationOptions)
}

// Init initializes the instance.
func (o NSOutlineView) Init() NSOutlineView {
	rv := objc.Send[NSOutlineView](o.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (o NSOutlineView) Autorelease() NSOutlineView {
	rv := objc.Send[NSOutlineView](o.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSOutlineView creates a new NSOutlineView instance.
func NewNSOutlineView() NSOutlineView {
	class := getNSOutlineViewClass()
	rv := objc.Send[NSOutlineView](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/AppKit/NSTableView/init(coder:)
func NewOutlineViewWithCoder(coder foundation.INSCoder) NSOutlineView {
	instance := getNSOutlineViewClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSOutlineViewFromID(rv)
}

//
// See: https://developer.apple.com/documentation/AppKit/NSTableView/init(frame:)
func NewOutlineViewWithFrame(frameRect corefoundation.CGRect) NSOutlineView {
	instance := getNSOutlineViewClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFrame:"), frameRect)
	return NSOutlineViewFromID(rv)
}

// Returns a Boolean value that indicates whether a given item is expandable.
//
// item: An item in the receiver.
//
// # Return Value
// 
// [true] if `item` is expandable—that is, `item` can contain other items,
// otherwise [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSOutlineView/isExpandable(_:)
func (o NSOutlineView) IsExpandable(item objectivec.IObject) bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isExpandable:"), item)
	return rv
}

// Returns a Boolean value that indicates whether a given item is expanded.
//
// item: An item in the receiver.
//
// # Return Value
// 
// [true] if `item` is expanded, otherwise [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSOutlineView/isItemExpanded(_:)
func (o NSOutlineView) IsItemExpanded(item objectivec.IObject) bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isItemExpanded:"), item)
	return rv
}

// Expands a given item.
//
// item: An item in the receiver.
//
// # Discussion
// 
// If `item` is not expandable or is already expanded, does nothing.
// 
// If expanding takes place, posts an item expanded notification.
//
// See: https://developer.apple.com/documentation/AppKit/NSOutlineView/expandItem(_:)
func (o NSOutlineView) ExpandItem(item objectivec.IObject) {
	objc.Send[objc.ID](o.ID, objc.Sel("expandItem:"), item)
}

// Expands a specified item and, optionally, its children.
//
// item: An item in the receiver.
// 
// Starting in OS X version 10.5, passing `'nil'` will expand each item under
// the root in the outline view.
//
// expandChildren: If [true], recursively expands `item` and its children. If [false], expands
// `item` only (identical to [ExpandItem]).
// //
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// For example, this method is invoked with the `expandChildren` parameter set
// to [true] when a user Option-clicks the disclosure triangle for an item in
// the outline view (to expand the item and all its contained items).
// 
// For each item expanded, posts an item expanded notification.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSOutlineView/expandItem(_:expandChildren:)
func (o NSOutlineView) ExpandItemExpandChildren(item objectivec.IObject, expandChildren bool) {
	objc.Send[objc.ID](o.ID, objc.Sel("expandItem:expandChildren:"), item, expandChildren)
}

// Collapses a given item.
//
// item: An item in the receiver.
//
// # Discussion
// 
// If `item` is not expanded or not expandable, does nothing
// 
// If collapsing takes place, posts item collapse notification.
//
// See: https://developer.apple.com/documentation/AppKit/NSOutlineView/collapseItem(_:)
func (o NSOutlineView) CollapseItem(item objectivec.IObject) {
	objc.Send[objc.ID](o.ID, objc.Sel("collapseItem:"), item)
}

// Collapses a given item and, optionally, its children.
//
// item: An item in the receiver.
// 
// Starting in OS X version 10.5, passing `'nil'` will collapse each item
// under the root in the outline view.
//
// collapseChildren: If [true], recursively collapses `item` and its children. If [false],
// collapses `item` only (identical to [CollapseItem]).
// //
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// For example, this method is invoked with the `collapseChildren` parameter
// set to [true] when a user Option-clicks the disclosure triangle for an item
// in the outline view (to collapse the item and all its contained items).
// 
// For each item collapsed, posts an item collapsed notification.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSOutlineView/collapseItem(_:collapseChildren:)
func (o NSOutlineView) CollapseItemCollapseChildren(item objectivec.IObject, collapseChildren bool) {
	objc.Send[objc.ID](o.ID, objc.Sel("collapseItem:collapseChildren:"), item, collapseChildren)
}

// Reloads and redisplays the data for the given item.
//
// item: The item to reload and display.
//
// # Discussion
// 
// Reloading the cell views associated with `item` occurs only in apps that
// link against macOS 10.12 and later.
// 
// This method may cause the outline view to change its selection without
// calling the [OutlineViewSelectionDidChange] delegate method.
//
// See: https://developer.apple.com/documentation/AppKit/NSOutlineView/reloadItem(_:)
func (o NSOutlineView) ReloadItem(item objectivec.IObject) {
	objc.Send[objc.ID](o.ID, objc.Sel("reloadItem:"), item)
}

// Reloads a given item and, optionally, its children.
//
// item: An item in the receiver.
// 
// Starting in OS X version 10.5, passing `'nil'` will reload everything under
// the root in the outline view.
//
// reloadChildren: If [true], recursively reloads `item` and its children. If [false], reloads
// `item` only (identical to [ReloadItem]).
// 
// It is not necessary, or efficient, to reload children if the item is not
// expanded.
// //
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSOutlineView/reloadItem(_:reloadChildren:)
func (o NSOutlineView) ReloadItemReloadChildren(item objectivec.IObject, reloadChildren bool) {
	objc.Send[objc.ID](o.ID, objc.Sel("reloadItem:reloadChildren:"), item, reloadChildren)
}

// Returns the item associated with a given row.
//
// row: The index of a row in the receiver.
//
// # Return Value
// 
// The item associated with `row`.
//
// See: https://developer.apple.com/documentation/AppKit/NSOutlineView/item(atRow:)
func (o NSOutlineView) ItemAtRow(row int) objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("itemAtRow:"), row)
	return objectivec.Object{ID: rv}
}

// Returns the row associated with a given item.
//
// item: An item in the receiver.
//
// # Return Value
// 
// The row associated with `item`, or `–1` if `item` is `nil` or cannot be
// found.
//
// See: https://developer.apple.com/documentation/AppKit/NSOutlineView/row(forItem:)
func (o NSOutlineView) RowForItem(item objectivec.IObject) int {
	rv := objc.Send[int](o.ID, objc.Sel("rowForItem:"), item)
	return rv
}

// Returns the indentation level for a given item.
//
// item: An item in the receiver.
//
// # Return Value
// 
// The indentation level for `item`. If `item` is `nil` (which is the root
// item), returns `–1`.
//
// # Discussion
// 
// The levels are zero-based—that is, the first level of displayed items is
// level `0`.
//
// See: https://developer.apple.com/documentation/AppKit/NSOutlineView/level(forItem:)
func (o NSOutlineView) LevelForItem(item objectivec.IObject) int {
	rv := objc.Send[int](o.ID, objc.Sel("levelForItem:"), item)
	return rv
}

// Returns the indentation level for a given row.
//
// row: The index of a row in the receiver.
//
// # Return Value
// 
// The indentation level for `row`. For an invalid row, returns `–1`.
//
// # Discussion
// 
// The levels are zero-based—that is, the first level of displayed items is
// level `0`.
//
// See: https://developer.apple.com/documentation/AppKit/NSOutlineView/level(forRow:)
func (o NSOutlineView) LevelForRow(row int) int {
	rv := objc.Send[int](o.ID, objc.Sel("levelForRow:"), row)
	return rv
}

// Used to “retarget” a proposed drop.
//
// item: The target item.
//
// index: The drop index.
//
// # Discussion
// 
// For example, to specify a drop on `someOutlineItem`, you specify `item` as
// `someOutlineItem` and `index` as [OutlineViewDropOnItemIndex]. To specify a
// drop between child `2` and `3` of `someOutlineItem`, you specify `item` as
// `someOutlineItem` and `index` as `3` (children are a zero-based index). To
// specify a drop on an un-expandable `someOutlineItem`, you specify `item` as
// `someOutlineItem` and `index` as [OutlineViewDropOnItemIndex].
//
// See: https://developer.apple.com/documentation/AppKit/NSOutlineView/setDropItem(_:dropChildIndex:)
func (o NSOutlineView) SetDropItemDropChildIndex(item objectivec.IObject, index int) {
	objc.Send[objc.ID](o.ID, objc.Sel("setDropItem:dropChildIndex:"), item, index)
}

// Returns a Boolean value that indicates whether auto-expanded items should
// return to their original collapsed state.
//
// deposited: If [true], the drop terminated successfully; if [false] the drop failed.
// //
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Return Value
// 
// [true] if auto-expanded items should return to their original collapsed
// state; otherwise [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// Override this method to provide custom behavior. If the target of a drop is
// not auto-expanded (by hovering long enough) the drop target still gets
// expanded after a successful drop unless this method returns [true]. The
// default implementation returns [false] after a successful drop.
// 
// This method is called in a variety of situations. For example, it is called
// shortly after the [OutlineViewAcceptDropItemChildIndex] method is called
// and also if the drag exits the outline view (exiting the view is treated
// the same as a failed drop). The return value of the
// [OutlineViewAcceptDropItemChildIndex] method determines the incoming value
// of the `deposited` parameter.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSOutlineView/shouldCollapseAutoExpandedItems(forDeposited:)
func (o NSOutlineView) ShouldCollapseAutoExpandedItemsForDeposited(deposited bool) bool {
	rv := objc.Send[bool](o.ID, objc.Sel("shouldCollapseAutoExpandedItemsForDeposited:"), deposited)
	return rv
}

// Returns the parent for a given item.
//
// item: The item for which to return the parent.
//
// # Return Value
// 
// The parent for `item`, or `nil` if the parent is the root.
//
// See: https://developer.apple.com/documentation/AppKit/NSOutlineView/parent(forItem:)
func (o NSOutlineView) ParentForItem(item objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("parentForItem:"), item)
	return objectivec.Object{ID: rv}
}

// Returns the child index of the specified item within its parent.
//
// # Discussion
// 
// The performance of this method is O(1) at best and O(n) at worst.
//
// See: https://developer.apple.com/documentation/AppKit/NSOutlineView/childIndex(forItem:)
func (o NSOutlineView) ChildIndexForItem(item objectivec.IObject) int {
	rv := objc.Send[int](o.ID, objc.Sel("childIndexForItem:"), item)
	return rv
}

// Returns the specified child of an item.
//
// index: The index of the child item in the parent.
//
// item: The parent item whose child item you want to retrieve.
//
// # Return Value
// 
// The child item or `nil` if the item could not be found.
//
// # Discussion
// 
// You can call this method on an outline view with either a static or dynamic
// data source. For an outline view whose contents are dynamic, this method
// may call out to the [OutlineViewChildOfItem] method of the associated data
// source object.
//
// See: https://developer.apple.com/documentation/AppKit/NSOutlineView/child(_:ofItem:)
func (o NSOutlineView) ChildOfItem(index int, item objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("child:ofItem:"), index, item)
	return objectivec.Object{ID: rv}
}

// Returns the number of children for the specified parent item.
//
// item: The parent item.
//
// # Return Value
// 
// The number of children associated with the parent.
//
// # Discussion
// 
// You can call this method on an outline view with either a static or dynamic
// data source. For an outline view whose contents are dynamic, this method
// may call out to the [OutlineViewNumberOfChildrenOfItem] method of the
// associated data source object.
//
// See: https://developer.apple.com/documentation/AppKit/NSOutlineView/numberOfChildren(ofItem:)
func (o NSOutlineView) NumberOfChildrenOfItem(item objectivec.IObject) int {
	rv := objc.Send[int](o.ID, objc.Sel("numberOfChildrenOfItem:"), item)
	return rv
}

// Returns the frame of the outline cell for a given row.
//
// row: The index of the row for which to return the frame.
//
// # Return Value
// 
// The frame of the outline cell for the row at index `row`, considering the
// current indentation and the value in the [IndentationMarkerFollowsCell]
// property. If the row at index `row` is not an expandable row, returns
// [NSZeroRect].
//
// # Discussion
// 
// You can override this method in a subclass to return a custom frame for the
// outline button cell. If your override returns an empty rect, no outline
// cell is drawn for that row. You might do that, for example, so that the
// disclosure triangle will not be shown for a row that should never be
// expanded.
//
// See: https://developer.apple.com/documentation/AppKit/NSOutlineView/frameOfOutlineCell(atRow:)
func (o NSOutlineView) FrameOfOutlineCellAtRow(row int) corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](o.ID, objc.Sel("frameOfOutlineCellAtRow:"), row)
	return corefoundation.CGRect(rv)
}

// Inserts new items at the given indexes in the given parent with the
// specified optional animations.
//
// indexes: Indexes at which to insert items.
//
// parent: The parent for the items, or `nil` if the parent is the root.
//
// animationOptions: Animated slide effects used when inserting items.
//
// # Discussion
// 
// This method parallels the [InsertRowsAtIndexesWithAnimation] method of
// [NSTableView] and is used in a way similar to the [insert(_:at:)] method of
// [NSMutableArray]. The method does nothing if `parent` is not expanded. The
// actual item values are determined by the data source’s
// [OutlineViewChildOfItem] method (which is called only after [EndUpdates] to
// ensure data source integrity).
// 
// You can call this method multiple times within the same
// [BeginUpdates]/[EndUpdates] block; new insertions move previously inserted
// new items, just like modifying an array. Inserting an index beyond what is
// available throws an exception.
//
// [NSMutableArray]: https://developer.apple.com/documentation/Foundation/NSMutableArray
// [insert(_:at:)]: https://developer.apple.com/documentation/Foundation/NSMutableArray/insert(_:at:)-73pln
//
// See: https://developer.apple.com/documentation/AppKit/NSOutlineView/insertItems(at:inParent:withAnimation:)
func (o NSOutlineView) InsertItemsAtIndexesInParentWithAnimation(indexes foundation.NSIndexSet, parent objectivec.IObject, animationOptions NSTableViewAnimationOptions) {
	objc.Send[objc.ID](o.ID, objc.Sel("insertItemsAtIndexes:inParent:withAnimation:"), indexes, parent, animationOptions)
}

// Moves an item at a given index in the given parent to a new index in a new
// parent.
//
// fromIndex: Index of the item to be moved.
//
// oldParent: The parent of the item to be moved.
//
// toIndex: Index in the new parent to which the item is moved.
//
// newParent: The parent of the item after it is moved.
//
// # Discussion
// 
// This method parallels the [MoveRowAtIndexToIndex] method of [NSTableView].
// The `newParent` can be the same as `oldParent` to reorder an item within
// the same parent.
// 
// You can call this method multiple times within the same
// [BeginUpdates]/[EndUpdates] block. Moving from an invalid index, or to an
// invalid index, throws an exception.
//
// See: https://developer.apple.com/documentation/AppKit/NSOutlineView/moveItem(at:inParent:to:inParent:)
func (o NSOutlineView) MoveItemAtIndexInParentToIndexInParent(fromIndex int, oldParent objectivec.IObject, toIndex int, newParent objectivec.IObject) {
	objc.Send[objc.ID](o.ID, objc.Sel("moveItemAtIndex:inParent:toIndex:inParent:"), fromIndex, oldParent, toIndex, newParent)
}

// Removes items at the given indexes in the given parent with the specified
// optional animations.
//
// indexes: Indexes of the items to be removed.
//
// parent: The parent of the items to be removed.
//
// animationOptions: Animated slide effects used when removing items.
//
// # Discussion
// 
// This method parallels the [RemoveRowsAtIndexesWithAnimation] method of
// [NSTableView] and is used in a way similar to the [removeObjects(at:)]
// method of [NSMutableArray]. The method does nothing if `parent` is not
// expanded. If any of the child items is expanded, then all of its child rows
// are also be removed.
// 
// You can call this method multiple times within the same
// [BeginUpdates]/[EndUpdates] block; changes work just like modifying an
// array. Removing an item at an index beyond what is available throws an
// exception.
//
// [NSMutableArray]: https://developer.apple.com/documentation/Foundation/NSMutableArray
// [removeObjects(at:)]: https://developer.apple.com/documentation/Foundation/NSMutableArray/removeObjects(at:)
//
// See: https://developer.apple.com/documentation/AppKit/NSOutlineView/removeItems(at:inParent:withAnimation:)
func (o NSOutlineView) RemoveItemsAtIndexesInParentWithAnimation(indexes foundation.NSIndexSet, parent objectivec.IObject, animationOptions NSTableViewAnimationOptions) {
	objc.Send[objc.ID](o.ID, objc.Sel("removeItemsAtIndexes:inParent:withAnimation:"), indexes, parent, animationOptions)
}

// A Boolean value that indicates whether the outline view retains and
// releases the objects returned from its data source.
//
// # Discussion
// 
// When the value of this property is [true], the outline view retains and
// releases the objects returned to it from [DataSource]. When the value is
// [false], the outline view treats the objects as opaque items and assumes
// that the client has a retain on them. The default value is [true] for
// applications linked on macOS 10.12 and later, and [false] for applications
// linked on earlier versions of macOS. If you require the legacy behavior and
// your app links in macOS 10.12 or later, the value of this property must be
// explicitly set to [false] in code, because it is not encoded in the nib. In
// general, this is required if the items themselves create a retain cycle.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSOutlineView/stronglyReferencesItems
func (o NSOutlineView) StronglyReferencesItems() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("stronglyReferencesItems"))
	return rv
}
func (o NSOutlineView) SetStronglyReferencesItems(value bool) {
	objc.Send[struct{}](o.ID, objc.Sel("setStronglyReferencesItems:"), value)
}

// The table column in which hierarchical data is displayed.
//
// # Discussion
// 
// Each level of hierarchical data is indented by the amount specified by the
// [IndentationPerLevel] property (the default is `16.0`), and decorated with
// the indentation marker (disclosure triangle) on rows that are expandable.
// Outline table column data is archived with the rest of the outline view’s
// state information.
// 
// Attempts to set the value of this property to `nil` are silently ignored.
//
// See: https://developer.apple.com/documentation/AppKit/NSOutlineView/outlineTableColumn
func (o NSOutlineView) OutlineTableColumn() INSTableColumn {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outlineTableColumn"))
	return NSTableColumnFromID(objc.ID(rv))
}
func (o NSOutlineView) SetOutlineTableColumn(value INSTableColumn) {
	objc.Send[struct{}](o.ID, objc.Sel("setOutlineTableColumn:"), value)
}

// A Boolean value that indicates whether the outline view resizes its outline
// column when the user expands or collapses items.
//
// # Discussion
// 
// The outline column contains the cells with the expansion symbols and is
// generally the first column. The default value of this property is [true],
// which causes the outline column to be resized.
// 
// The outline column is resized based on how many indentation levels are
// exposed or hidden. For example, if expanding a row exposes a single
// indentation level, the outline column width is increased by one
// [IndentationPerLevel].
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSOutlineView/autoresizesOutlineColumn
func (o NSOutlineView) AutoresizesOutlineColumn() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("autoresizesOutlineColumn"))
	return rv
}
func (o NSOutlineView) SetAutoresizesOutlineColumn(value bool) {
	objc.Send[struct{}](o.ID, objc.Sel("setAutoresizesOutlineColumn:"), value)
}

// The per-level indentation, measured in points.
//
// See: https://developer.apple.com/documentation/AppKit/NSOutlineView/indentationPerLevel
func (o NSOutlineView) IndentationPerLevel() float64 {
	rv := objc.Send[float64](o.ID, objc.Sel("indentationPerLevel"))
	return rv
}
func (o NSOutlineView) SetIndentationPerLevel(value float64) {
	objc.Send[struct{}](o.ID, objc.Sel("setIndentationPerLevel:"), value)
}

// A Boolean value indicating whether the indentation marker symbol displayed
// in the outline column should be indented along with the cell contents.
//
// # Discussion
// 
// When the value of this property is [true], the indentation marker is
// indented along with the cell contents. When the value is [false], the
// marker is always displayed left-justified in the column. The default value
// of this property is [true].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSOutlineView/indentationMarkerFollowsCell
func (o NSOutlineView) IndentationMarkerFollowsCell() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("indentationMarkerFollowsCell"))
	return rv
}
func (o NSOutlineView) SetIndentationMarkerFollowsCell(value bool) {
	objc.Send[struct{}](o.ID, objc.Sel("setIndentationMarkerFollowsCell:"), value)
}

// A Boolean value indicating whether the expanded items are automatically
// saved across launches of the app.
//
// # Discussion
// 
// When the value of this property is [true], the outline view saves the state
// of its expanded items and restores that state the next time the user
// launches the app. (If the outline view’s [AutosaveName] property is
// `nil`, or if you have not implemented the
// [OutlineViewItemForPersistentObject] and
// [OutlineViewPersistentObjectForItem] delegate methods, this setting is
// ignored and outline information is not saved.) The configuration data is
// saved separately for each user and for each app. The default value of this
// property is [false].
// 
// You can have separate settings for the [AutosaveExpandedItems] and
// [AutosaveTableColumns] properties, so you could, for example, save expanded
// item information, but not table column positions.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSOutlineView/autosaveExpandedItems
func (o NSOutlineView) AutosaveExpandedItems() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("autosaveExpandedItems"))
	return rv
}
func (o NSOutlineView) SetAutosaveExpandedItems(value bool) {
	objc.Send[struct{}](o.ID, objc.Sel("setAutosaveExpandedItems:"), value)
}

			// Protocol methods for NSAccessibilityOutline
			

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

func (o NSOutlineView) AccessibilityFrame() corefoundation.CGRect {
	
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

func (o NSOutlineView) AccessibilityParent() objectivec.IObject {
	
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

func (o NSOutlineView) AccessibilityIdentifier() string {
	
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

func (o NSOutlineView) IsAccessibilityFocused() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityFocused"))
	return rv
	}

