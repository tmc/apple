// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"context"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSTableViewDiffableDataSource] class.
var (
	_NSTableViewDiffableDataSourceClass     NSTableViewDiffableDataSourceClass
	_NSTableViewDiffableDataSourceClassOnce sync.Once
)

func getNSTableViewDiffableDataSourceClass() NSTableViewDiffableDataSourceClass {
	_NSTableViewDiffableDataSourceClassOnce.Do(func() {
		_NSTableViewDiffableDataSourceClass = NSTableViewDiffableDataSourceClass{class: objc.GetClass("NSTableViewDiffableDataSource")}
	})
	return _NSTableViewDiffableDataSourceClass
}

// GetNSTableViewDiffableDataSourceClass returns the class object for NSTableViewDiffableDataSource.
func GetNSTableViewDiffableDataSourceClass() NSTableViewDiffableDataSourceClass {
	return getNSTableViewDiffableDataSourceClass()
}

type NSTableViewDiffableDataSourceClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSTableViewDiffableDataSourceClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSTableViewDiffableDataSourceClass) Alloc() NSTableViewDiffableDataSource {
	rv := objc.Send[NSTableViewDiffableDataSource](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// The object you use to manage data and provide items for a table view.
//
// # Overview
// 
// A object is a specialized type of data source that works together with your
// table view object. It provides the behavior you need to manage updates to
// your table view’s data and UI in a simple, efficient way. It also
// conforms to the [NSTableViewDataSource] protocol.
// 
// To fill a table view with data:
// 
// - Connect a diffable data source to your table view. - Implement a cell
// provider to configure your table view’s cells. - Generate the current
// state of the data. - Display the data in the UI.
// 
// To connect a diffable data source to a table view, you create the diffable
// data source using its [NSTableViewDiffableDataSource.InitWithTableViewCellProvider] initializer, passing
// in the table view you want to associate with that data source. You also
// pass in a cell provider, where you configure each of your cells to
// determine how to display your data in the UI.
// 
// Then, you generate the current state of the data and display the data in
// the UI by constructing and applying a snapshot. For more information, see
// [NSDiffableDataSourceSnapshot].
//
// # Creating a Diffable Data Source
//
//   - [NSTableViewDiffableDataSource.InitWithTableViewCellProvider]: Creates a diffable data source with the specified cell provider, and connects it to the specified table view.
//
// # Creating Row and Section Views
//
//   - [NSTableViewDiffableDataSource.RowViewProvider]: The closure that configures and returns the table view’s row views from the diffable data source.
//   - [NSTableViewDiffableDataSource.SetRowViewProvider]
//   - [NSTableViewDiffableDataSource.SectionHeaderViewProvider]: The closure that configures and returns the table view’s section header views from the diffable data source.
//   - [NSTableViewDiffableDataSource.SetSectionHeaderViewProvider]
//
// # Identifying Items and Sections
//
//   - [NSTableViewDiffableDataSource.ItemIdentifierForRow]: Returns an identifier for the item at the specified row in the table view.
//   - [NSTableViewDiffableDataSource.RowForItemIdentifier]: Returns a row for the item with the specified identifier in the table view.
//   - [NSTableViewDiffableDataSource.SectionIdentifierForRow]: Returns the identifier of the section containing the specified row in the snapshot.
//   - [NSTableViewDiffableDataSource.RowForSectionIdentifier]: Returns a row for the section with the specified identifier in the table view.
//
// # Updating Data
//
//   - [NSTableViewDiffableDataSource.Snapshot]: Returns a representation of the current state of the data in the table view.
//   - [NSTableViewDiffableDataSource.ApplySnapshotAnimatingDifferences]: Updates the UI to reflect the state of the data in the specified snapshot, optionally animating the UI changes.
//   - [NSTableViewDiffableDataSource.ApplySnapshotAnimatingDifferencesCompletion]: Updates the UI to reflect the state of the data in the specified snapshot, optionally animating the UI changes and executing a completion handler.
//   - [NSTableViewDiffableDataSource.DefaultRowAnimation]: The default animation the UI uses to show differences between rows.
//   - [NSTableViewDiffableDataSource.SetDefaultRowAnimation]
//
// See: https://developer.apple.com/documentation/AppKit/NSTableViewDiffableDataSourceReference
type NSTableViewDiffableDataSource struct {
	objectivec.Object
}

// NSTableViewDiffableDataSourceFromID constructs a [NSTableViewDiffableDataSource] from an objc.ID.
//
// The object you use to manage data and provide items for a table view.
func NSTableViewDiffableDataSourceFromID(id objc.ID) NSTableViewDiffableDataSource {
	return NSTableViewDiffableDataSource{objectivec.Object{ID: id}}
}
// NOTE: NSTableViewDiffableDataSource adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSTableViewDiffableDataSource] class.
//
// # Creating a Diffable Data Source
//
//   - [INSTableViewDiffableDataSource.InitWithTableViewCellProvider]: Creates a diffable data source with the specified cell provider, and connects it to the specified table view.
//
// # Creating Row and Section Views
//
//   - [INSTableViewDiffableDataSource.RowViewProvider]: The closure that configures and returns the table view’s row views from the diffable data source.
//   - [INSTableViewDiffableDataSource.SetRowViewProvider]
//   - [INSTableViewDiffableDataSource.SectionHeaderViewProvider]: The closure that configures and returns the table view’s section header views from the diffable data source.
//   - [INSTableViewDiffableDataSource.SetSectionHeaderViewProvider]
//
// # Identifying Items and Sections
//
//   - [INSTableViewDiffableDataSource.ItemIdentifierForRow]: Returns an identifier for the item at the specified row in the table view.
//   - [INSTableViewDiffableDataSource.RowForItemIdentifier]: Returns a row for the item with the specified identifier in the table view.
//   - [INSTableViewDiffableDataSource.SectionIdentifierForRow]: Returns the identifier of the section containing the specified row in the snapshot.
//   - [INSTableViewDiffableDataSource.RowForSectionIdentifier]: Returns a row for the section with the specified identifier in the table view.
//
// # Updating Data
//
//   - [INSTableViewDiffableDataSource.Snapshot]: Returns a representation of the current state of the data in the table view.
//   - [INSTableViewDiffableDataSource.ApplySnapshotAnimatingDifferences]: Updates the UI to reflect the state of the data in the specified snapshot, optionally animating the UI changes.
//   - [INSTableViewDiffableDataSource.ApplySnapshotAnimatingDifferencesCompletion]: Updates the UI to reflect the state of the data in the specified snapshot, optionally animating the UI changes and executing a completion handler.
//   - [INSTableViewDiffableDataSource.DefaultRowAnimation]: The default animation the UI uses to show differences between rows.
//   - [INSTableViewDiffableDataSource.SetDefaultRowAnimation]
//
// See: https://developer.apple.com/documentation/AppKit/NSTableViewDiffableDataSourceReference
type INSTableViewDiffableDataSource interface {
	objectivec.IObject
	NSTableViewDataSource

	// Topic: Creating a Diffable Data Source

	// Creates a diffable data source with the specified cell provider, and connects it to the specified table view.
	InitWithTableViewCellProvider(tableView INSTableView, cellProvider NSTableViewDiffableDataSourceCellProvider) NSTableViewDiffableDataSource

	// Topic: Creating Row and Section Views

	// The closure that configures and returns the table view’s row views from the diffable data source.
	RowViewProvider() NSTableViewDiffableDataSourceRowProvider
	SetRowViewProvider(value NSTableViewDiffableDataSourceRowProvider)
	// The closure that configures and returns the table view’s section header views from the diffable data source.
	SectionHeaderViewProvider() NSTableViewDiffableDataSourceSectionHeaderViewProvider
	SetSectionHeaderViewProvider(value NSTableViewDiffableDataSourceSectionHeaderViewProvider)

	// Topic: Identifying Items and Sections

	// Returns an identifier for the item at the specified row in the table view.
	ItemIdentifierForRow(row int) objectivec.IObject
	// Returns a row for the item with the specified identifier in the table view.
	RowForItemIdentifier(identifier objectivec.IObject) int
	// Returns the identifier of the section containing the specified row in the snapshot.
	SectionIdentifierForRow(row int) objectivec.IObject
	// Returns a row for the section with the specified identifier in the table view.
	RowForSectionIdentifier(identifier objectivec.IObject) int

	// Topic: Updating Data

	// Returns a representation of the current state of the data in the table view.
	Snapshot() INSDiffableDataSourceSnapshot
	// Updates the UI to reflect the state of the data in the specified snapshot, optionally animating the UI changes.
	ApplySnapshotAnimatingDifferences(snapshot INSDiffableDataSourceSnapshot, animatingDifferences bool)
	// Updates the UI to reflect the state of the data in the specified snapshot, optionally animating the UI changes and executing a completion handler.
	ApplySnapshotAnimatingDifferencesCompletion(snapshot INSDiffableDataSourceSnapshot, animatingDifferences bool, completion VoidHandler)
	// The default animation the UI uses to show differences between rows.
	DefaultRowAnimation() NSTableViewAnimationOptions
	SetDefaultRowAnimation(value NSTableViewAnimationOptions)
}

// Init initializes the instance.
func (t NSTableViewDiffableDataSource) Init() NSTableViewDiffableDataSource {
	rv := objc.Send[NSTableViewDiffableDataSource](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t NSTableViewDiffableDataSource) Autorelease() NSTableViewDiffableDataSource {
	rv := objc.Send[NSTableViewDiffableDataSource](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSTableViewDiffableDataSource creates a new NSTableViewDiffableDataSource instance.
func NewNSTableViewDiffableDataSource() NSTableViewDiffableDataSource {
	class := getNSTableViewDiffableDataSourceClass()
	rv := objc.Send[NSTableViewDiffableDataSource](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a diffable data source with the specified cell provider, and
// connects it to the specified table view.
//
// tableView: The initialized table view object to connect to the diffable data source.
//
// cellProvider: A closure that creates and returns each of the cells for the table view
// from the data the diffable data source provides. This replaces the
// [TableViewViewForTableColumnRow] delegate method.
//
// See: https://developer.apple.com/documentation/AppKit/NSTableViewDiffableDataSourceReference/init(tableView:cellProvider:)
func NewTableViewDiffableDataSourceWithTableViewCellProvider(tableView INSTableView, cellProvider NSTableViewDiffableDataSourceCellProvider) NSTableViewDiffableDataSource {
	instance := getNSTableViewDiffableDataSourceClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTableView:cellProvider:"), tableView, cellProvider)
	return NSTableViewDiffableDataSourceFromID(rv)
}

// Creates a diffable data source with the specified cell provider, and
// connects it to the specified table view.
//
// tableView: The initialized table view object to connect to the diffable data source.
//
// cellProvider: A closure that creates and returns each of the cells for the table view
// from the data the diffable data source provides. This replaces the
// [TableViewViewForTableColumnRow] delegate method.
//
// See: https://developer.apple.com/documentation/AppKit/NSTableViewDiffableDataSourceReference/init(tableView:cellProvider:)
func (t NSTableViewDiffableDataSource) InitWithTableViewCellProvider(tableView INSTableView, cellProvider NSTableViewDiffableDataSourceCellProvider) NSTableViewDiffableDataSource {
	rv := objc.Send[NSTableViewDiffableDataSource](t.ID, objc.Sel("initWithTableView:cellProvider:"), tableView, cellProvider)
	return rv
}
// Returns an identifier for the item at the specified row in the table view.
//
// row: The row of the item in the table view.
//
// # Return Value
// 
// The item’s identifier, or `nil` if the method doesn’t find an item at
// the provided row.
//
// See: https://developer.apple.com/documentation/AppKit/NSTableViewDiffableDataSourceReference/itemIdentifier(forRow:)
func (t NSTableViewDiffableDataSource) ItemIdentifierForRow(row int) objectivec.IObject {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("itemIdentifierForRow:"), row)
	return objectivec.Object{ID: rv}
}
// Returns a row for the item with the specified identifier in the table view.
//
// identifier: The identifier of the item in the table view.
//
// # Return Value
// 
// The item’s row, or `nil` if the method doesn’t find an item with the
// provided item identifier.
//
// See: https://developer.apple.com/documentation/AppKit/NSTableViewDiffableDataSourceReference/row(forItemIdentifier:)
func (t NSTableViewDiffableDataSource) RowForItemIdentifier(identifier objectivec.IObject) int {
	rv := objc.Send[int](t.ID, objc.Sel("rowForItemIdentifier:"), identifier)
	return rv
}
// Returns the identifier of the section containing the specified row in the
// snapshot.
//
// row: The row of the section in the table view.
//
// # Return Value
// 
// The section’s identifier, or `nil` if the method doesn’t find an item
// with the provided item identifier.
//
// See: https://developer.apple.com/documentation/AppKit/NSTableViewDiffableDataSourceReference/sectionIdentifier(forRow:)
func (t NSTableViewDiffableDataSource) SectionIdentifierForRow(row int) objectivec.IObject {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("sectionIdentifierForRow:"), row)
	return objectivec.Object{ID: rv}
}
// Returns a row for the section with the specified identifier in the table
// view.
//
// identifier: The identifier of the section in the table view.
//
// # Return Value
// 
// The item’s section, or `nil` if the method doesn’t find an item with
// the provided item identifier.
//
// See: https://developer.apple.com/documentation/AppKit/NSTableViewDiffableDataSourceReference/row(forSectionIdentifier:)
func (t NSTableViewDiffableDataSource) RowForSectionIdentifier(identifier objectivec.IObject) int {
	rv := objc.Send[int](t.ID, objc.Sel("rowForSectionIdentifier:"), identifier)
	return rv
}
// Returns a representation of the current state of the data in the table
// view.
//
// # Return Value
// 
// A snapshot that contains row and item identifiers in the order they appear
// in the UI.
//
// See: https://developer.apple.com/documentation/AppKit/NSTableViewDiffableDataSourceReference/snapshot()
func (t NSTableViewDiffableDataSource) Snapshot() INSDiffableDataSourceSnapshot {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("snapshot"))
	return NSDiffableDataSourceSnapshotFromID(rv)
}
// Updates the UI to reflect the state of the data in the specified snapshot,
// optionally animating the UI changes.
//
// snapshot: The snapshot that reflects the new state of the data in the table view.
//
// animatingDifferences: If `true`, the diffable data source computes the difference between the
// table view’s current state and the new state in the snapshot, which is an
// O(n) operation, where n is the number of items in the snapshot. The system
// animates the differences in the UI between the current state and new state.
// If `false`, the system sets the table view UI to the new state without any
// animations, with no additional overhead for computing a diff. Any ongoing
// row animations are interrupted and the table view’s content immediately
// reloads.
//
// See: https://developer.apple.com/documentation/AppKit/NSTableViewDiffableDataSourceReference/applySnapshot(_:animatingDifferences:)
func (t NSTableViewDiffableDataSource) ApplySnapshotAnimatingDifferences(snapshot INSDiffableDataSourceSnapshot, animatingDifferences bool) {
	objc.Send[objc.ID](t.ID, objc.Sel("applySnapshot:animatingDifferences:"), snapshot, animatingDifferences)
}
// Updates the UI to reflect the state of the data in the specified snapshot,
// optionally animating the UI changes and executing a completion handler.
//
// snapshot: The snapshot that reflects the new state of the data in the table view.
//
// animatingDifferences: If `true`, the diffable data source computes the difference between the
// table view’s current state and the new state in the snapshot, which is an
// O(n) operation, where n is the number of items in the snapshot. The system
// animates the differences in the UI between the current state and new state.
// If `false`, the system sets the table view UI to the new state without any
// animations, with no additional overhead for computing a diff. Any ongoing
// row animations are interrupted and the table view’s content immediately
// reloads.
//
// completion: The closure the system executes when the animations are complete. This
// closure has no return value and takes no parameters. The system calls this
// closure from the main queue.
//
// # Discussion
//
// See: https://developer.apple.com/documentation/AppKit/NSTableViewDiffableDataSourceReference/applySnapshot(_:animatingDifferences:completion:)
func (t NSTableViewDiffableDataSource) ApplySnapshotAnimatingDifferencesCompletion(snapshot INSDiffableDataSourceSnapshot, animatingDifferences bool, completion VoidHandler) {
_block2, _ := NewVoidBlock(completion)
	objc.Send[objc.ID](t.ID, objc.Sel("applySnapshot:animatingDifferences:completion:"), snapshot, animatingDifferences, _block2)
}
// Returns the number of records managed for `aTableView` by the data source
// object.
//
// tableView: The table view that sent the message.
//
// # Return Value
// 
// The number of rows in `aTableView`.
//
// # Discussion
// 
// An instance of [NSTableView] uses this method to determine how many rows it
// should create and display. Your [NumberOfRowsInTableView] implementation is
// called very frequently, so it must be efficient.
// 
// Both view-based table views and cell-based table views must implement this
// method.
//
// See: https://developer.apple.com/documentation/AppKit/NSTableViewDataSource/numberOfRows(in:)
func (t NSTableViewDiffableDataSource) NumberOfRowsInTableView(tableView INSTableView) int {
	rv := objc.Send[int](t.ID, objc.Sel("numberOfRowsInTableView:"), tableView)
	return rv
}
// Called by `aTableView` when the mouse button is released over a table view
// that previously decided to allow a drop.
//
// tableView: The table view that sent the message.
//
// info: An object that contains more information about this dragging operation.
//
// row: The index of the proposed target row.
//
// dropOperation: The type of dragging operation.
//
// # Return Value
// 
// [true] if the drop operation was successful, otherwise [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// The data source should incorporate the data from the dragging pasteboard in
// the implementation of this method. You can use the [DraggingPasteboard]
// method to get the data for the drop operation from `info`.
// 
// To accept a drop on the second row, `row` would be 2 and `operation` would
// be [NSTableViewDropOn]. To accept a drop below the last row, `row` would be
// `[aTableView numberOfRows]` and `operation` would be
// [NSTableViewDropAbove].
//
// See: https://developer.apple.com/documentation/AppKit/NSTableViewDataSource/tableView(_:acceptDrop:row:dropOperation:)
func (t NSTableViewDiffableDataSource) TableViewAcceptDropRowDropOperation(tableView INSTableView, info NSDraggingInfo, row int, dropOperation NSTableViewDropOperation) bool {
	rv := objc.Send[bool](t.ID, objc.Sel("tableView:acceptDrop:row:dropOperation:"), tableView, info, row, dropOperation)
	return rv
}
// Implement this method to determine when a dragging session has ended.
//
// tableView: The table view.
//
// session: The dragging session.
//
// screenPoint: The ending drag location in screen coordinates.
//
// operation: The drag operation. See [NSDragOperation] for supported values.
// //
// [NSDragOperation]: https://developer.apple.com/documentation/AppKit/NSDragOperation
//
// # Discussion
// 
// This delegate method can be used to determine when the dragging source
// operation ended at a specific location, such as the trash, by checking for
// an operation of [DragOperationDelete].
//
// See: https://developer.apple.com/documentation/AppKit/NSTableViewDataSource/tableView(_:draggingSession:endedAt:operation:)
func (t NSTableViewDiffableDataSource) TableViewDraggingSessionEndedAtPointOperation(tableView INSTableView, session INSDraggingSession, screenPoint corefoundation.CGPoint, operation NSDragOperation) {
	objc.Send[objc.ID](t.ID, objc.Sel("tableView:draggingSession:endedAtPoint:operation:"), tableView, session, screenPoint, operation)
}
// Implement this method to determine when a dragging session will begin.
//
// tableView: The table view.
//
// session: The dragging session.
//
// screenPoint: The initial drag location in screen coordinates.
//
// rowIndexes: The indexes of the rows to be dragged, excluding rows that were not dragged
// due to [TableViewPasteboardWriterForRow] returning `nil`.
//
// # Discussion
// 
// Implement this method to know when the dragging session is about to begin
// and to potentially modify the dragging session.
// 
// The dragged item order will directly match the pasteboard writer array used
// to begin the dragging session with the [NSView] method
// [BeginDraggingSessionWithItemsEventSource]. Hence, the order is
// deterministic, and can be used in [TableViewAcceptDropRowDropOperation]
// when enumerating the [NSDraggingInfo] pasteboard classes.
//
// See: https://developer.apple.com/documentation/AppKit/NSTableViewDataSource/tableView(_:draggingSession:willBeginAt:forRowIndexes:)
func (t NSTableViewDiffableDataSource) TableViewDraggingSessionWillBeginAtPointForRowIndexes(tableView INSTableView, session INSDraggingSession, screenPoint corefoundation.CGPoint, rowIndexes foundation.NSIndexSet) {
	objc.Send[objc.ID](t.ID, objc.Sel("tableView:draggingSession:willBeginAtPoint:forRowIndexes:"), tableView, session, screenPoint, rowIndexes)
}
// Called by the table view to return the data object associated with the
// specified row and column.
//
// tableView: The table view that sent the message.
//
// tableColumn: A column in `aTableView`.
//
// row: The row of the item in `aTableColumn`.
//
// # Return Value
// 
// An item in the data source in the specified table column of the view.
//
// # Discussion
// 
// [TableViewObjectValueForTableColumnRow] is called each time the table cell
// needs to be redisplayed, so it must be efficient.
//
// See: https://developer.apple.com/documentation/AppKit/NSTableViewDataSource/tableView(_:objectValueFor:row:)
func (t NSTableViewDiffableDataSource) TableViewObjectValueForTableColumnRow(tableView INSTableView, tableColumn INSTableColumn, row int) objectivec.IObject {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("tableView:objectValueForTableColumn:row:"), tableView, tableColumn, row)
	return objectivec.Object{ID: rv}
}
// Called to allow the table to support multiple item dragging.
//
// tableView: The table view.
//
// row: The row.
//
// # Return Value
// 
// Returns an instance of [NSPasteboardItem] or a custom object that
// implements the [NSPasteboardWriting] protocol. Returning `nil` excludes the
// row from being dragged.
//
// # Discussion
// 
// This method is required for multi-image dragging.
// 
// If this method is implemented, then [tableView(_:writeRowsWith:to:)] will
// not be called.
//
// [tableView(_:writeRowsWith:to:)]: https://developer.apple.com/documentation/AppKit/NSTableViewDataSource/tableView(_:writeRowsWith:to:)
//
// See: https://developer.apple.com/documentation/AppKit/NSTableViewDataSource/tableView(_:pasteboardWriterForRow:)
func (t NSTableViewDiffableDataSource) TableViewPasteboardWriterForRow(tableView INSTableView, row int) NSPasteboardWriting {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("tableView:pasteboardWriterForRow:"), tableView, row)
	return NSPasteboardWritingObjectFromID(rv)
}
// Sets the data object for an item in the specified row and column.
//
// tableView: The table view that sent the message.
//
// object: The new value for the item.
//
// tableColumn: A column in `aTableView`.
//
// row: The row of the item in `aTableColumn`.
//
// # Discussion
// 
// This method is intended for use with cell-based table views, it must not be
// used with view-based table views. In view-based tables, use target/action
// to set each item in the view cell.
//
// See: https://developer.apple.com/documentation/AppKit/NSTableViewDataSource/tableView(_:setObjectValue:for:row:)
func (t NSTableViewDiffableDataSource) TableViewSetObjectValueForTableColumnRow(tableView INSTableView, object objectivec.IObject, tableColumn INSTableColumn, row int) {
	objc.Send[objc.ID](t.ID, objc.Sel("tableView:setObjectValue:forTableColumn:row:"), tableView, object, tableColumn, row)
}
// Called by `aTableView` to indicate that sorting may need to be done.
//
// tableView: The table view that sent the message.
//
// oldDescriptors: An array that contains the previous descriptors.
//
// # Discussion
// 
// The data source typically sorts and reloads the data, and adjusts the
// selections accordingly. If you need to know the current sort descriptors
// and the data source doesn’t manage them itself, you can get the current
// sort descriptors by sending `aTableView` a [SortDescriptors] message.
//
// See: https://developer.apple.com/documentation/AppKit/NSTableViewDataSource/tableView(_:sortDescriptorsDidChange:)
func (t NSTableViewDiffableDataSource) TableViewSortDescriptorsDidChange(tableView INSTableView, oldDescriptors []foundation.NSSortDescriptor) {
	objc.Send[objc.ID](t.ID, objc.Sel("tableView:sortDescriptorsDidChange:"), tableView, objectivec.IObjectSliceToNSArray(oldDescriptors))
}
// Implement this method to allow the table to update dragging items as they
// are dragged over a view.
//
// tableView: The table view.
//
// draggingInfo: The dragging information.
//
// # Discussion
// 
// Required for multi-image dragging. Typically this will involve invoking
// [EnumerateDraggingItemsWithOptionsForViewClassesSearchOptionsUsingBlock] on
// the `draggingInfo` parameter value and setting the `draggingItem`
// object’s [ImageComponentsProvider] to a proper image based on the
// content.
// 
// For view-based table views, you can use the [NSTableCellView] method
// [DraggingImageComponents]. For cell-based tables, use the [NSCell] method
// [DraggingImageComponentsWithFrameInView].
//
// See: https://developer.apple.com/documentation/AppKit/NSTableViewDataSource/tableView(_:updateDraggingItemsForDrag:)
func (t NSTableViewDiffableDataSource) TableViewUpdateDraggingItemsForDrag(tableView INSTableView, draggingInfo NSDraggingInfo) {
	objc.Send[objc.ID](t.ID, objc.Sel("tableView:updateDraggingItemsForDrag:"), tableView, draggingInfo)
}
// Used by `aTableView` to determine a valid drop target.
//
// tableView: The table view that sent the message.
//
// info: An object that contains more information about this dragging operation.
//
// row: The index of the proposed target row.
//
// dropOperation: The type of dragging operation proposed.
//
// # Return Value
// 
// The dragging operation the data source will perform.
//
// # Discussion
// 
// The data source may retarget a drop by calling [SetDropRowDropOperation]
// and returning something other than [NSDragOperationNone]. A data source
// might retarget for various reasons, such as to provide better visual
// feedback when inserting into a sorted position.
// 
// To propose a drop on the second row, `row` would be 2 and `operation` would
// be [NSTableViewDropOn]. To propose a drop below the last row, `row` would
// be `[aTableView numberOfRows]` and `operation` would be
// [NSTableViewDropAbove].
//
// See: https://developer.apple.com/documentation/AppKit/NSTableViewDataSource/tableView(_:validateDrop:proposedRow:proposedDropOperation:)
func (t NSTableViewDiffableDataSource) TableViewValidateDropProposedRowProposedDropOperation(tableView INSTableView, info NSDraggingInfo, row int, dropOperation NSTableViewDropOperation) NSDragOperation {
	rv := objc.Send[NSDragOperation](t.ID, objc.Sel("tableView:validateDrop:proposedRow:proposedDropOperation:"), tableView, info, row, dropOperation)
	return NSDragOperation(rv)
}

// The closure that configures and returns the table view’s row views from
// the diffable data source.
//
// # Discussion
// 
// This property replaces the [TableViewRowViewForRow] delegate method.
//
// See: https://developer.apple.com/documentation/AppKit/NSTableViewDiffableDataSourceReference/rowViewProvider
func (t NSTableViewDiffableDataSource) RowViewProvider() NSTableViewDiffableDataSourceRowProvider {
	rv := objc.Send[NSTableViewDiffableDataSourceRowProvider](t.ID, objc.Sel("rowViewProvider"))
	return NSTableViewDiffableDataSourceRowProvider(rv)
}
func (t NSTableViewDiffableDataSource) SetRowViewProvider(value NSTableViewDiffableDataSourceRowProvider) {
	objc.Send[struct{}](t.ID, objc.Sel("setRowViewProvider:"), value)
}
// The closure that configures and returns the table view’s section header
// views from the diffable data source.
//
// # Discussion
// 
// This property replaces the [TableViewViewForTableColumnRow] method for
// group rows when the `tableColumn` parameter in
// [TableViewViewForTableColumnRow] would be `nil`. Setting this property
// means the table view never invokes its delegate’s [TableViewIsGroupRow]
// method. Instead, it uses the current snapshot’s sections.
//
// See: https://developer.apple.com/documentation/AppKit/NSTableViewDiffableDataSourceReference/sectionHeaderViewProvider
func (t NSTableViewDiffableDataSource) SectionHeaderViewProvider() NSTableViewDiffableDataSourceSectionHeaderViewProvider {
	rv := objc.Send[NSTableViewDiffableDataSourceSectionHeaderViewProvider](t.ID, objc.Sel("sectionHeaderViewProvider"))
	return NSTableViewDiffableDataSourceSectionHeaderViewProvider(rv)
}
func (t NSTableViewDiffableDataSource) SetSectionHeaderViewProvider(value NSTableViewDiffableDataSourceSectionHeaderViewProvider) {
	objc.Send[struct{}](t.ID, objc.Sel("setSectionHeaderViewProvider:"), value)
}
// The default animation the UI uses to show differences between rows.
//
// # Discussion
// 
// The default value of this property is [TableViewAnimationEffectFade].
// 
// If you set the value of this property, the new value becomes the default
// row animation for the next update that uses
// [ApplySnapshotAnimatingDifferences].
//
// See: https://developer.apple.com/documentation/AppKit/NSTableViewDiffableDataSourceReference/defaultRowAnimation
func (t NSTableViewDiffableDataSource) DefaultRowAnimation() NSTableViewAnimationOptions {
	rv := objc.Send[NSTableViewAnimationOptions](t.ID, objc.Sel("defaultRowAnimation"))
	return NSTableViewAnimationOptions(rv)
}
func (t NSTableViewDiffableDataSource) SetDefaultRowAnimation(value NSTableViewAnimationOptions) {
	objc.Send[struct{}](t.ID, objc.Sel("setDefaultRowAnimation:"), value)
}

			// Protocol methods for NSTableViewDataSource
			

// ApplySnapshotAnimatingDifferencesCompletionSync is a synchronous wrapper around [NSTableViewDiffableDataSource.ApplySnapshotAnimatingDifferencesCompletion].
// It blocks until the completion handler fires or the context is cancelled.
func (t NSTableViewDiffableDataSource) ApplySnapshotAnimatingDifferencesCompletionSync(ctx context.Context, snapshot INSDiffableDataSourceSnapshot, animatingDifferences bool) error {
	done := make(chan struct{}, 1)
	t.ApplySnapshotAnimatingDifferencesCompletion(snapshot, animatingDifferences, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

