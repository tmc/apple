// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSCollectionView] class.
var (
	_NSCollectionViewClass     NSCollectionViewClass
	_NSCollectionViewClassOnce sync.Once
)

func getNSCollectionViewClass() NSCollectionViewClass {
	_NSCollectionViewClassOnce.Do(func() {
		_NSCollectionViewClass = NSCollectionViewClass{class: objc.GetClass("NSCollectionView")}
	})
	return _NSCollectionViewClass
}

// GetNSCollectionViewClass returns the class object for NSCollectionView.
func GetNSCollectionViewClass() NSCollectionViewClass {
	return getNSCollectionViewClass()
}

type NSCollectionViewClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSCollectionViewClass) Alloc() NSCollectionView {
	rv := objc.Send[NSCollectionView](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// An ordered collection of data items displayed in a customizable layout.
//
// # Overview
// 
// The simplest type of collection view displays its items in a grid, but you
// can define layouts to arrange items however you like. For example, you
// might create a layout where items are arranged in a circle. You can also
// change layouts dynamically at runtime whenever you need to present items
// differently.
// 
// You can add collection views to your interface using Interface Builder or
// create them programmatically in your view controller or window controller
// code. It is recommended that you configure your collection view with a data
// source object, which is an object that conforms to the
// [NSCollectionViewDataSource] protocol. Data sources support multiple
// sections and the modern layout architecture and are the preferred way for
// specifying your data.
// 
// In addition to displaying items, collection views support the display of
// supplementary and decoration views. Support for supplementary and
// decoration views is defined by the current layout object, but both types of
// views add to the visual presentation of your content. Supplementary views
// are associated with a specific section and can be used to create header and
// footer views for a related group of items. Decoration views are purely
// visual adornments and can be used to implement dynamic backgrounds or other
// types of configurable visual content.
// 
// The layout of a collection view can be changed dynamically by assigning a
// new layout object to the [CollectionViewLayout] property. Changing the
// layout object updates the appearance of the collection view without
// animating the changes.
// 
// # The Objects of a Collection View Interface
// 
// An [NSCollectionView] object itself is a facilitator, taking information
// from disparate sources and merging them together to create an overall
// interface:
// 
// - The data source object provides both the data and the views used to
// display that data. You define the data source object by implementing the
// methods of the [NSCollectionViewDataSource] protocol in one of your app’s
// objects. - The visual representation of items is provided by the
// [NSCollectionViewItem] class. Item objects are view controllers and you use
// their views to display your app’s data. The data source creates items on
// demand and returns them to the collection view for display. - The
// collection view delegate makes decisions about behaviors. The delegate also
// coordinates the dragging and dropping of items. You define the delegate by
// implementing the methods of the [NSCollectionViewDelegate] protocol in one
// of your app’s objects. - The layout object specifies the position and
// appearance of items onscreen. AppKit defines layout objects that you can
// use as-is, but you can also define custom layouts by subclassing
// [NSCollectionViewLayout].
// 
// [Figure 1] illustrates how the collection view works with its other objects
// to create its final appearance. The collection view obtains the views for
// items and supplementary views from its data source, which creates the views
// and fills them with data. The layout object provides the layout attributes
// needed to position those items and supplementary views onscreen. The
// collection view merges the two sets of information to create the final
// appearance that the user sees onscreen.
// 
// [media-1965644]
// 
// There are other helper classes and protocols that you can use to customize
// the layout behavior and other aspects of the collection view interface. For
// example, when using a flow layout object ([NSCollectionViewFlowLayout]),
// you can modify the flow layout’s behavior using the methods of the
// [NSCollectionViewDelegateFlowLayout] protocol. When implementing a custom
// layout, you might also work with [NSCollectionViewUpdateItem] and
// [NSCollectionViewLayoutInvalidationContext] objects, which help the layout
// object manage updates.
// 
// # Managing the Collection View’s Content
// 
// Data for the collection view is managed by the —that is an object that
// adopts the methods of the [NSCollectionViewDataSource] protocol. You are
// responsible for defining the data source used by your collection view. The
// data source provides information about the number of sections and items in
// the collection view and it provides the visual representation of that data.
// Every data source object is required to implement the following methods:
// 
// - [CollectionViewNumberOfItemsInSection] -
// [CollectionViewItemForRepresentedObjectAtIndexPath]
// 
// The [NSCollectionViewItem] class defines the visual appearance of items in
// the collection view. Your data source object vends items from its
// [CollectionViewItemForRepresentedObjectAtIndexPath] method, creating and
// configuring the item in one step. Each item is essentially a snapshot of
// the data it represents. Items are often short-lived because they can be
// recycled by the collection view and reused to display new data. As a
// result, never store references to items in your app.
// 
// Supplementary views are another way to display data in your interface. Each
// layout object defines the supplementary views it supports, and different
// layouts can define supplementary views for different purposes. For example,
// an [NSCollectionViewFlowLayout] object lets you add header and footer views
// to each section. Your data source must know enough about the layout to know
// which supplementary views are supported by the layout object and how those
// views are displayed. The data source can then provide supplementary views
// when asked for them.
// 
// When your content changes in a way that requires you to update what the
// collection view displays, call the [NSCollectionView.ReloadData], [NSCollectionView.ReloadSections], or
// [NSCollectionView.ReloadItemsAtIndexPaths] method to perform that update. These methods
// cause the collection view to discard the views currently being used to
// display your content and ask for new ones. Never try to modify the views
// associated with your items directly. The collection view does not maintain
// views for all items, only those that are currently being displayed.
// Reloading the items ensures that the views are updated correctly.
// 
// For more information on defining your data source object, see
// [NSCollectionViewDataSource].
// 
// # Inserting, Deleting, and Moving Content
// 
// The collection view includes methods for inserting, deleting, and moving
// items and sections. All of these methods affect only what the collection
// view displays onscreen; they do not change the data in the associated data
// source object. As a result, when updating your collection view’s content,
// always do the following:
// 
// - Update the internal structures of your data source object first. - Call
// the [NSCollectionView] methods to insert, delete, or move items and
// sections.
// 
// When you call methods like [NSCollectionView.InsertItemsAtIndexPaths] or [NSCollectionView.DeleteSections],
// the collection view fetches any new data from your data source object and
// then updates the layout. When inserting, moving, or deleting items, the
// collection view updates the layout for all affected items, which might
// include items not directly affected by the operation. For example,
// inserting one item might require adjusting the onscreen position of many
// other items. When the layout attributes for any visible items changes, the
// collection view animates those changes into place automatically.
// 
// The layout object determines how inserted and deleted items are animated
// into position. Because newly inserted items are not onscreen initially, the
// layout object provides the initial layout attributes for those items.
// Similarly, the layout object provides the final layout attributes for any
// items that are being deleted. For example, the layout object might specify
// final layout attributes that are offscreen so that a deleted item animates
// out of the visible rectangle.
// 
// Because individual methods for inserting, deleting, and moving content
// animate their changes right away, you must use the
// [NSCollectionView.PerformBatchUpdatesCompletionHandler] method when you want to animate
// multiple changes together. The [NSCollectionView.PerformBatchUpdatesCompletionHandler]
// method takes a block containing all of the insert, delete, move, and reload
// method calls you need to update the collection view. All of those
// operations are captured and performed as a single animated sequence.
// 
// # Interface Builder Configuration Options
// 
// Xcode lets you configure information about your collection view in your
// storyboard and nib files. The table below shows the basic collection view
// attributes. Additional attributes are available based on the selected value
// for the Layout attribute.
// 
// [Table data omitted]
// 
// The table below shows the attributes you can configure when you set the
// Layout attribute to Flow.
// 
// [Table data omitted]
// 
// The table below shows the attributes you can configure when you set the
// Layout attribute to Grid.
// 
// [Table data omitted]
// 
// The table below shows the attributes you can configure when you set the
// Layout attribute to Custom.
// 
// [Table data omitted]
// 
// The table below shows the attributes you can configure when you set the
// Layout attribute to Content Array (Legacy).
// 
// [Table data omitted]
// 
// # Legacy Collection View Support
// 
// Prior to OS X v10.11, the collection view always displayed its contents in
// a grid structure that could not be changed. The data for the collection
// view was stored in the [NSCollectionView.Content] property, which was often populated with
// data using bindings. You specified the visual appearance for the collection
// view’s data by creating an [NSCollectionViewItem] object and assigning it
// to the [itemPrototype] property. That item object acted as a template and
// was used to create all of the items in the collection view.
// 
// You are encouraged to use the modern collection view architecture when
// configuring collection views in macOS 10.11 and later. Use the legacy
// architecture only for apps that must run in earlier versions of macOS.
// 
// For more information about how to configure a collection view using the
// legacy architecture, see Collection View Programming Guide for macOS.
//
// [itemPrototype]: https://developer.apple.com/documentation/AppKit/NSCollectionView/itemPrototype
//
// # Providing the Collection View’s Data
//
//   - [NSCollectionView.DataSource]: An object that provides data for the collection view.
//   - [NSCollectionView.SetDataSource]
//
// # Configuring the Collection View
//
//   - [NSCollectionView.Delegate]: The collection view’s delegate object.
//   - [NSCollectionView.SetDelegate]
//   - [NSCollectionView.Content]: An array that provides data for the collection view.
//   - [NSCollectionView.SetContent]
//   - [NSCollectionView.BackgroundView]: The background view placed behind all items and supplementary views.
//   - [NSCollectionView.SetBackgroundView]
//   - [NSCollectionView.BackgroundColors]: An array containing the collection view’s background colors.
//   - [NSCollectionView.SetBackgroundColors]
//   - [NSCollectionView.BackgroundViewScrollsWithContent]: A Boolean value that indicates whether the collection view’s background view scrolls with the items and other content.
//   - [NSCollectionView.SetBackgroundViewScrollsWithContent]
//
// # Creating Collection View Items
//
//   - [NSCollectionView.MakeItemWithIdentifierForIndexPath]: Creates or returns a reusable item object of the specified type.
//   - [NSCollectionView.RegisterClassForItemWithIdentifier]: Registers a class to use when creating new items in the collection view.
//   - [NSCollectionView.RegisterNibForItemWithIdentifier]: Registers a nib file to use when creating items in the collection view.
//   - [NSCollectionView.MakeSupplementaryViewOfKindWithIdentifierForIndexPath]: Creates or returns a reusable supplementary view of the specified type.
//   - [NSCollectionView.RegisterClassForSupplementaryViewOfKindWithIdentifier]: Registers a class to use when creating new supplementary views in the collection view.
//   - [NSCollectionView.RegisterNibForSupplementaryViewOfKindWithIdentifier]: Registers a nib file to use when creating supplementary views in the collection view.
//
// # Changing the Layout
//
//   - [NSCollectionView.CollectionViewLayout]: The layout object used to organize the collection view’s content.
//   - [NSCollectionView.SetCollectionViewLayout]
//
// # Reloading Content
//
//   - [NSCollectionView.ReloadData]: Reloads all of the data for the collection view.
//   - [NSCollectionView.ReloadSections]: Reloads the data in the specified sections of the collection view.
//   - [NSCollectionView.ReloadItemsAtIndexPaths]: Reloads only the specified items.
//
// # Prefetching Collection View Cells and Data
//
//   - [NSCollectionView.PrefetchDataSource]
//   - [NSCollectionView.SetPrefetchDataSource]
//
// # Getting the State of the Collection View
//
//   - [NSCollectionView.NumberOfSections]: The number of sections in the collection view.
//   - [NSCollectionView.NumberOfItemsInSection]: Returns the number of items in the specified section.
//
// # Inserting, Moving, and Deleting Items
//
//   - [NSCollectionView.InsertItemsAtIndexPaths]: Inserts new items into the collection view at the specified locations.
//   - [NSCollectionView.MoveItemAtIndexPathToIndexPath]: Moves an item from one location to another in the collection view.
//   - [NSCollectionView.DeleteItemsAtIndexPaths]: Deletes the items at the specified index paths.
//
// # Inserting, Moving, Deleting, and Collapsing Sections
//
//   - [NSCollectionView.InsertSections]: Inserts new sections at the specified indexes.
//   - [NSCollectionView.MoveSectionToSection]: Moves a section from its current location to a new location.
//   - [NSCollectionView.DeleteSections]: Deletes the specified sections and their contained items.
//   - [NSCollectionView.ToggleSectionCollapse]: Collapses the section in which the sender resides into a single horizontally scrollable row.
//
// # Managing the Selection
//
//   - [NSCollectionView.Selectable]: A Boolean value that indicates whether the user may select items in the collection view.
//   - [NSCollectionView.SetSelectable]
//   - [NSCollectionView.AllowsMultipleSelection]: A Boolean value that indicates whether the user may select more than one item in the collection view.
//   - [NSCollectionView.SetAllowsMultipleSelection]
//   - [NSCollectionView.AllowsEmptySelection]: A Boolean value indicating whether the collection view may have no selected items.
//   - [NSCollectionView.SetAllowsEmptySelection]
//   - [NSCollectionView.SelectionIndexPaths]: The set of index paths representing the currently selected items.
//   - [NSCollectionView.SetSelectionIndexPaths]
//   - [NSCollectionView.DeselectAll]: Deselects all items in the collection view.
//   - [NSCollectionView.SelectItemsAtIndexPathsScrollPosition]: Adds the specified items to the current selection and optionally scrolls the items into position.
//   - [NSCollectionView.DeselectItemsAtIndexPaths]: Removes the specified items from the current selection.
//
// # Locating Items and Views
//
//   - [NSCollectionView.VisibleItems]: Returns an array of the actively managed items in the collection view.
//   - [NSCollectionView.IndexPathsForVisibleItems]: Returns the index paths of the currently active items.
//   - [NSCollectionView.VisibleSupplementaryViewsOfKind]: Returns an array of the actively managed supplementary views in the collection view.
//   - [NSCollectionView.IndexPathsForVisibleSupplementaryElementsOfKind]: Returns the index paths of the currently active supplementary views.
//   - [NSCollectionView.IndexPathForItem]: Returns the index path of the specified item.
//   - [NSCollectionView.IndexPathForItemAtPoint]: Returns the index path of the item at the specified point.
//   - [NSCollectionView.ItemAtIndexPath]: Returns the item associated with the specified index path.
//   - [NSCollectionView.SupplementaryViewForElementKindAtIndexPath]: Returns the supplementary view associated with the specified index path.
//   - [NSCollectionView.ScrollToItemsAtIndexPathsScrollPosition]: Scrolls the collection view contents until the specified items are visible.
//
// # Getting Layout Information
//
//   - [NSCollectionView.LayoutAttributesForItemAtIndexPath]: Returns the layout information for the item at the specified index path.
//   - [NSCollectionView.LayoutAttributesForSupplementaryElementOfKindAtIndexPath]: Returns the layout information for the supplementary view at the specified index path.
//
// # Animating Multiple Changes
//
//   - [NSCollectionView.PerformBatchUpdatesCompletionHandler]: Encapsulates multiple insert, delete, reload, and move operations into a single animated operation.
//
// # Working with the Responder Chain
//
//   - [NSCollectionView.FirstResponder]: A Boolean value indicating whether the collection view is the first responder.
//
// # Getting a Drag Image
//
//   - [NSCollectionView.DraggingImageForItemsAtIndexPathsWithEventOffset]: Returns an image to use for dragging the specified items.
//
// # Legacy Collection View Support
//
//   - [NSCollectionView.SelectionIndexes]: The indexes of the currently selected items.
//   - [NSCollectionView.SetSelectionIndexes]
//   - [NSCollectionView.ItemAtIndex]: Returns the collection view item for the represented object at the specified index.
//   - [NSCollectionView.FrameForItemAtIndex]: Returns the frame of the collection view item at the specified index.
//   - [NSCollectionView.FrameForItemAtIndexWithNumberOfItems]: Returns the frame of an item based on the number of items in the collection view.
//   - [NSCollectionView.DraggingImageForItemsAtIndexesWithEventOffset]: This method computes and returns an image to use for dragging.
//   - [NSCollectionView.SetDraggingSourceOperationMaskForLocal]: Configures the drag operation mask.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionView
type NSCollectionView struct {
	NSView
}

// NSCollectionViewFromID constructs a [NSCollectionView] from an objc.ID.
//
// An ordered collection of data items displayed in a customizable layout.
func NSCollectionViewFromID(id objc.ID) NSCollectionView {
	return NSCollectionView{NSView: NSViewFromID(id)}
}
// NOTE: NSCollectionView adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [NSCollectionView] class.
//
// # Providing the Collection View’s Data
//
//   - [INSCollectionView.DataSource]: An object that provides data for the collection view.
//   - [INSCollectionView.SetDataSource]
//
// # Configuring the Collection View
//
//   - [INSCollectionView.Delegate]: The collection view’s delegate object.
//   - [INSCollectionView.SetDelegate]
//   - [INSCollectionView.Content]: An array that provides data for the collection view.
//   - [INSCollectionView.SetContent]
//   - [INSCollectionView.BackgroundView]: The background view placed behind all items and supplementary views.
//   - [INSCollectionView.SetBackgroundView]
//   - [INSCollectionView.BackgroundColors]: An array containing the collection view’s background colors.
//   - [INSCollectionView.SetBackgroundColors]
//   - [INSCollectionView.BackgroundViewScrollsWithContent]: A Boolean value that indicates whether the collection view’s background view scrolls with the items and other content.
//   - [INSCollectionView.SetBackgroundViewScrollsWithContent]
//
// # Creating Collection View Items
//
//   - [INSCollectionView.MakeItemWithIdentifierForIndexPath]: Creates or returns a reusable item object of the specified type.
//   - [INSCollectionView.RegisterClassForItemWithIdentifier]: Registers a class to use when creating new items in the collection view.
//   - [INSCollectionView.RegisterNibForItemWithIdentifier]: Registers a nib file to use when creating items in the collection view.
//   - [INSCollectionView.MakeSupplementaryViewOfKindWithIdentifierForIndexPath]: Creates or returns a reusable supplementary view of the specified type.
//   - [INSCollectionView.RegisterClassForSupplementaryViewOfKindWithIdentifier]: Registers a class to use when creating new supplementary views in the collection view.
//   - [INSCollectionView.RegisterNibForSupplementaryViewOfKindWithIdentifier]: Registers a nib file to use when creating supplementary views in the collection view.
//
// # Changing the Layout
//
//   - [INSCollectionView.CollectionViewLayout]: The layout object used to organize the collection view’s content.
//   - [INSCollectionView.SetCollectionViewLayout]
//
// # Reloading Content
//
//   - [INSCollectionView.ReloadData]: Reloads all of the data for the collection view.
//   - [INSCollectionView.ReloadSections]: Reloads the data in the specified sections of the collection view.
//   - [INSCollectionView.ReloadItemsAtIndexPaths]: Reloads only the specified items.
//
// # Prefetching Collection View Cells and Data
//
//   - [INSCollectionView.PrefetchDataSource]
//   - [INSCollectionView.SetPrefetchDataSource]
//
// # Getting the State of the Collection View
//
//   - [INSCollectionView.NumberOfSections]: The number of sections in the collection view.
//   - [INSCollectionView.NumberOfItemsInSection]: Returns the number of items in the specified section.
//
// # Inserting, Moving, and Deleting Items
//
//   - [INSCollectionView.InsertItemsAtIndexPaths]: Inserts new items into the collection view at the specified locations.
//   - [INSCollectionView.MoveItemAtIndexPathToIndexPath]: Moves an item from one location to another in the collection view.
//   - [INSCollectionView.DeleteItemsAtIndexPaths]: Deletes the items at the specified index paths.
//
// # Inserting, Moving, Deleting, and Collapsing Sections
//
//   - [INSCollectionView.InsertSections]: Inserts new sections at the specified indexes.
//   - [INSCollectionView.MoveSectionToSection]: Moves a section from its current location to a new location.
//   - [INSCollectionView.DeleteSections]: Deletes the specified sections and their contained items.
//   - [INSCollectionView.ToggleSectionCollapse]: Collapses the section in which the sender resides into a single horizontally scrollable row.
//
// # Managing the Selection
//
//   - [INSCollectionView.Selectable]: A Boolean value that indicates whether the user may select items in the collection view.
//   - [INSCollectionView.SetSelectable]
//   - [INSCollectionView.AllowsMultipleSelection]: A Boolean value that indicates whether the user may select more than one item in the collection view.
//   - [INSCollectionView.SetAllowsMultipleSelection]
//   - [INSCollectionView.AllowsEmptySelection]: A Boolean value indicating whether the collection view may have no selected items.
//   - [INSCollectionView.SetAllowsEmptySelection]
//   - [INSCollectionView.SelectionIndexPaths]: The set of index paths representing the currently selected items.
//   - [INSCollectionView.SetSelectionIndexPaths]
//   - [INSCollectionView.DeselectAll]: Deselects all items in the collection view.
//   - [INSCollectionView.SelectItemsAtIndexPathsScrollPosition]: Adds the specified items to the current selection and optionally scrolls the items into position.
//   - [INSCollectionView.DeselectItemsAtIndexPaths]: Removes the specified items from the current selection.
//
// # Locating Items and Views
//
//   - [INSCollectionView.VisibleItems]: Returns an array of the actively managed items in the collection view.
//   - [INSCollectionView.IndexPathsForVisibleItems]: Returns the index paths of the currently active items.
//   - [INSCollectionView.VisibleSupplementaryViewsOfKind]: Returns an array of the actively managed supplementary views in the collection view.
//   - [INSCollectionView.IndexPathsForVisibleSupplementaryElementsOfKind]: Returns the index paths of the currently active supplementary views.
//   - [INSCollectionView.IndexPathForItem]: Returns the index path of the specified item.
//   - [INSCollectionView.IndexPathForItemAtPoint]: Returns the index path of the item at the specified point.
//   - [INSCollectionView.ItemAtIndexPath]: Returns the item associated with the specified index path.
//   - [INSCollectionView.SupplementaryViewForElementKindAtIndexPath]: Returns the supplementary view associated with the specified index path.
//   - [INSCollectionView.ScrollToItemsAtIndexPathsScrollPosition]: Scrolls the collection view contents until the specified items are visible.
//
// # Getting Layout Information
//
//   - [INSCollectionView.LayoutAttributesForItemAtIndexPath]: Returns the layout information for the item at the specified index path.
//   - [INSCollectionView.LayoutAttributesForSupplementaryElementOfKindAtIndexPath]: Returns the layout information for the supplementary view at the specified index path.
//
// # Animating Multiple Changes
//
//   - [INSCollectionView.PerformBatchUpdatesCompletionHandler]: Encapsulates multiple insert, delete, reload, and move operations into a single animated operation.
//
// # Working with the Responder Chain
//
//   - [INSCollectionView.FirstResponder]: A Boolean value indicating whether the collection view is the first responder.
//
// # Getting a Drag Image
//
//   - [INSCollectionView.DraggingImageForItemsAtIndexPathsWithEventOffset]: Returns an image to use for dragging the specified items.
//
// # Legacy Collection View Support
//
//   - [INSCollectionView.SelectionIndexes]: The indexes of the currently selected items.
//   - [INSCollectionView.SetSelectionIndexes]
//   - [INSCollectionView.ItemAtIndex]: Returns the collection view item for the represented object at the specified index.
//   - [INSCollectionView.FrameForItemAtIndex]: Returns the frame of the collection view item at the specified index.
//   - [INSCollectionView.FrameForItemAtIndexWithNumberOfItems]: Returns the frame of an item based on the number of items in the collection view.
//   - [INSCollectionView.DraggingImageForItemsAtIndexesWithEventOffset]: This method computes and returns an image to use for dragging.
//   - [INSCollectionView.SetDraggingSourceOperationMaskForLocal]: Configures the drag operation mask.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionView
type INSCollectionView interface {
	INSView
	NSDraggingSource

	// Topic: Providing the Collection View’s Data

	// An object that provides data for the collection view.
	DataSource() NSCollectionViewDataSource
	SetDataSource(value NSCollectionViewDataSource)

	// Topic: Configuring the Collection View

	// The collection view’s delegate object.
	Delegate() NSCollectionViewDelegate
	SetDelegate(value NSCollectionViewDelegate)
	// An array that provides data for the collection view.
	Content() []objectivec.IObject
	SetContent(value []objectivec.IObject)
	// The background view placed behind all items and supplementary views.
	BackgroundView() INSView
	SetBackgroundView(value INSView)
	// An array containing the collection view’s background colors.
	BackgroundColors() []NSColor
	SetBackgroundColors(value []NSColor)
	// A Boolean value that indicates whether the collection view’s background view scrolls with the items and other content.
	BackgroundViewScrollsWithContent() bool
	SetBackgroundViewScrollsWithContent(value bool)

	// Topic: Creating Collection View Items

	// Creates or returns a reusable item object of the specified type.
	MakeItemWithIdentifierForIndexPath(identifier NSUserInterfaceItemIdentifier, indexPath objectivec.IObject) INSCollectionViewItem
	// Registers a class to use when creating new items in the collection view.
	RegisterClassForItemWithIdentifier(itemClass objc.Class, identifier NSUserInterfaceItemIdentifier)
	// Registers a nib file to use when creating items in the collection view.
	RegisterNibForItemWithIdentifier(nib INSNib, identifier NSUserInterfaceItemIdentifier)
	// Creates or returns a reusable supplementary view of the specified type.
	MakeSupplementaryViewOfKindWithIdentifierForIndexPath(elementKind NSCollectionViewSupplementaryElementKind, identifier NSUserInterfaceItemIdentifier, indexPath objectivec.IObject) INSView
	// Registers a class to use when creating new supplementary views in the collection view.
	RegisterClassForSupplementaryViewOfKindWithIdentifier(viewClass objc.Class, kind NSCollectionViewSupplementaryElementKind, identifier NSUserInterfaceItemIdentifier)
	// Registers a nib file to use when creating supplementary views in the collection view.
	RegisterNibForSupplementaryViewOfKindWithIdentifier(nib INSNib, kind NSCollectionViewSupplementaryElementKind, identifier NSUserInterfaceItemIdentifier)

	// Topic: Changing the Layout

	// The layout object used to organize the collection view’s content.
	CollectionViewLayout() INSCollectionViewLayout
	SetCollectionViewLayout(value INSCollectionViewLayout)

	// Topic: Reloading Content

	// Reloads all of the data for the collection view.
	ReloadData()
	// Reloads the data in the specified sections of the collection view.
	ReloadSections(sections foundation.NSIndexSet)
	// Reloads only the specified items.
	ReloadItemsAtIndexPaths(indexPaths foundation.INSSet)

	// Topic: Prefetching Collection View Cells and Data

	PrefetchDataSource() NSCollectionViewPrefetching
	SetPrefetchDataSource(value NSCollectionViewPrefetching)

	// Topic: Getting the State of the Collection View

	// The number of sections in the collection view.
	NumberOfSections() int
	// Returns the number of items in the specified section.
	NumberOfItemsInSection(section int) int

	// Topic: Inserting, Moving, and Deleting Items

	// Inserts new items into the collection view at the specified locations.
	InsertItemsAtIndexPaths(indexPaths foundation.INSSet)
	// Moves an item from one location to another in the collection view.
	MoveItemAtIndexPathToIndexPath(indexPath objectivec.IObject, newIndexPath objectivec.IObject)
	// Deletes the items at the specified index paths.
	DeleteItemsAtIndexPaths(indexPaths foundation.INSSet)

	// Topic: Inserting, Moving, Deleting, and Collapsing Sections

	// Inserts new sections at the specified indexes.
	InsertSections(sections foundation.NSIndexSet)
	// Moves a section from its current location to a new location.
	MoveSectionToSection(section int, newSection int)
	// Deletes the specified sections and their contained items.
	DeleteSections(sections foundation.NSIndexSet)
	// Collapses the section in which the sender resides into a single horizontally scrollable row.
	ToggleSectionCollapse(sender objectivec.IObject)

	// Topic: Managing the Selection

	// A Boolean value that indicates whether the user may select items in the collection view.
	Selectable() bool
	SetSelectable(value bool)
	// A Boolean value that indicates whether the user may select more than one item in the collection view.
	AllowsMultipleSelection() bool
	SetAllowsMultipleSelection(value bool)
	// A Boolean value indicating whether the collection view may have no selected items.
	AllowsEmptySelection() bool
	SetAllowsEmptySelection(value bool)
	// The set of index paths representing the currently selected items.
	SelectionIndexPaths() foundation.INSSet
	SetSelectionIndexPaths(value foundation.INSSet)
	// Deselects all items in the collection view.
	DeselectAll(sender objectivec.IObject)
	// Adds the specified items to the current selection and optionally scrolls the items into position.
	SelectItemsAtIndexPathsScrollPosition(indexPaths foundation.INSSet, scrollPosition NSCollectionViewScrollPosition)
	// Removes the specified items from the current selection.
	DeselectItemsAtIndexPaths(indexPaths foundation.INSSet)

	// Topic: Locating Items and Views

	// Returns an array of the actively managed items in the collection view.
	VisibleItems() []NSCollectionViewItem
	// Returns the index paths of the currently active items.
	IndexPathsForVisibleItems() foundation.INSSet
	// Returns an array of the actively managed supplementary views in the collection view.
	VisibleSupplementaryViewsOfKind(elementKind NSCollectionViewSupplementaryElementKind) []NSView
	// Returns the index paths of the currently active supplementary views.
	IndexPathsForVisibleSupplementaryElementsOfKind(elementKind NSCollectionViewSupplementaryElementKind) foundation.INSSet
	// Returns the index path of the specified item.
	IndexPathForItem(item INSCollectionViewItem) objc.ID
	// Returns the index path of the item at the specified point.
	IndexPathForItemAtPoint(point corefoundation.CGPoint) objc.ID
	// Returns the item associated with the specified index path.
	ItemAtIndexPath(indexPath objectivec.IObject) INSCollectionViewItem
	// Returns the supplementary view associated with the specified index path.
	SupplementaryViewForElementKindAtIndexPath(elementKind NSCollectionViewSupplementaryElementKind, indexPath objectivec.IObject) INSView
	// Scrolls the collection view contents until the specified items are visible.
	ScrollToItemsAtIndexPathsScrollPosition(indexPaths foundation.INSSet, scrollPosition NSCollectionViewScrollPosition)

	// Topic: Getting Layout Information

	// Returns the layout information for the item at the specified index path.
	LayoutAttributesForItemAtIndexPath(indexPath objectivec.IObject) INSCollectionViewLayoutAttributes
	// Returns the layout information for the supplementary view at the specified index path.
	LayoutAttributesForSupplementaryElementOfKindAtIndexPath(kind NSCollectionViewSupplementaryElementKind, indexPath objectivec.IObject) INSCollectionViewLayoutAttributes

	// Topic: Animating Multiple Changes

	// Encapsulates multiple insert, delete, reload, and move operations into a single animated operation.
	PerformBatchUpdatesCompletionHandler(updates VoidHandler, completionHandler ErrorHandler)

	// Topic: Working with the Responder Chain

	// A Boolean value indicating whether the collection view is the first responder.
	FirstResponder() bool

	// Topic: Getting a Drag Image

	// Returns an image to use for dragging the specified items.
	DraggingImageForItemsAtIndexPathsWithEventOffset(indexPaths foundation.INSSet, event INSEvent, dragImageOffset foundation.NSPoint) INSImage

	// Topic: Legacy Collection View Support

	// The indexes of the currently selected items.
	SelectionIndexes() foundation.NSIndexSet
	SetSelectionIndexes(value foundation.NSIndexSet)
	// Returns the collection view item for the represented object at the specified index.
	ItemAtIndex(index uint) INSCollectionViewItem
	// Returns the frame of the collection view item at the specified index.
	FrameForItemAtIndex(index uint) corefoundation.CGRect
	// Returns the frame of an item based on the number of items in the collection view.
	FrameForItemAtIndexWithNumberOfItems(index uint, numberOfItems uint) corefoundation.CGRect
	// This method computes and returns an image to use for dragging.
	DraggingImageForItemsAtIndexesWithEventOffset(indexes foundation.NSIndexSet, event INSEvent, dragImageOffset foundation.NSPoint) INSImage
	// Configures the drag operation mask.
	SetDraggingSourceOperationMaskForLocal(dragOperationMask NSDragOperation, localDestination bool)

	EncodeWithCoder(coder foundation.INSCoder)
}





// Init initializes the instance.
func (c NSCollectionView) Init() NSCollectionView {
	rv := objc.Send[NSCollectionView](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c NSCollectionView) Autorelease() NSCollectionView {
	rv := objc.Send[NSCollectionView](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSCollectionView creates a new NSCollectionView instance.
func NewNSCollectionView() NSCollectionView {
	class := getNSCollectionViewClass()
	rv := objc.Send[NSCollectionView](objc.ID(class.class), objc.Sel("new"))
	return rv
}






// Initializes a view using from data in the specified coder object.
//
// coder: The coder object that contains the view’s configuration details.
//
// # Return Value
// 
// An initialized view or `nil` if AppKit couldn’t create the object.
//
// See: https://developer.apple.com/documentation/AppKit/NSView/init(coder:)
func NewCollectionViewWithCoder(coder foundation.INSCoder) NSCollectionView {
	instance := getNSCollectionViewClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSCollectionViewFromID(rv)
}


// Initializes and returns a newly allocated [NSView] object with a specified
// frame rectangle.
//
// frameRect: The frame rectangle for the created view object.
//
// # Return Value
// 
// An initialized view or `nil` if AppKit couldn’t create the object.
//
// # Discussion
// 
// Insert the view into your window’s view hieararchy before you can do
// anything with it. This method is the designated initializer for the
// [NSView] class.
//
// See: https://developer.apple.com/documentation/AppKit/NSView/init(frame:)
func NewCollectionViewWithFrame(frameRect corefoundation.CGRect) NSCollectionView {
	instance := getNSCollectionViewClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFrame:"), frameRect)
	return NSCollectionViewFromID(rv)
}







// Creates or returns a reusable item object of the specified type.
//
// identifier: The reuse identifier for the specified item. This is the identifier you
// specified when registering the item. This parameter must not be `nil`.
//
// indexPath: The index path specifying the location of the item. The data source object
// receives this information in its
// [CollectionViewItemForRepresentedObjectAtIndexPath] method and you should
// just pass it along.
//
// # Return Value
// 
// A valid [NSCollectionViewItem] object.
//
// # Discussion
// 
// This method looks for a recycled item object of the specified type and
// returns it if one exists. If one does not exist, it creates it using one of
// the following techniques:
// 
// - If you used the [RegisterClassForItemWithIdentifier] method to register a
// class for the identifier, this method instantiates your class and returns
// it. - If you used the [RegisterNibForItemWithIdentifier] method to register
// a nib file for the identifier, this method loads the item from the nib file
// and returns it.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionView/makeItem(withIdentifier:for:)
func (c NSCollectionView) MakeItemWithIdentifierForIndexPath(identifier NSUserInterfaceItemIdentifier, indexPath objectivec.IObject) INSCollectionViewItem {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("makeItemWithIdentifier:forIndexPath:"), objc.String(string(identifier)), indexPath)
	return NSCollectionViewItemFromID(rv)
}

// Registers a class to use when creating new items in the collection view.
//
// itemClass: A class to use for creating items. The class must be descended from
// [NSCollectionViewItem]. Specify `nil` to unregister a previously registered
// class or nib file.
//
// identifier: The string that identifies the type of item. You use this string later when
// requesting new items and it must be unique among the other registered item
// and view classes of this collection view. This parameter must not be an
// empty string or `nil`.
//
// # Discussion
// 
// Use this method to register the classes that represent items in your
// collection view. When you request an item using the
// [ItemWithIdentifierForIndexPath] method, the collection view recycles an
// existing item with the same `identifier` or creates a new one by
// instantiating your class and calling the [init()] method of the resulting
// object.
// 
// Because items are recycled to improve performance, it is recommended that
// your custom classes conform to the [NSCollectionViewElement] protocol. You
// can use the methods of that protocol to prepare your classes for reuse.
// 
// Typically, you register your items when initializing your collection view
// interface. Although you can register new items at any time, you must not
// call the [ItemWithIdentifierForIndexPath] method until after you register
// the corresponding item.
//
// [init()]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/init()
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionView/register(_:forItemWithIdentifier:)-6s4i
func (c NSCollectionView) RegisterClassForItemWithIdentifier(itemClass objc.Class, identifier NSUserInterfaceItemIdentifier) {
	objc.Send[objc.ID](c.ID, objc.Sel("registerClass:forItemWithIdentifier:"), itemClass, objc.String(string(identifier)))
}

// Registers a nib file to use when creating items in the collection view.
//
// nib: The nib object containing the item’s definition. The nib file must
// contain exactly one [NSCollectionViewItem] object at the top level. You may
// use a custom subclass when configuring the object in the nib file. Specify
// `nil` to unregister a previously registered class or nib file.
//
// identifier: The string that identifies the type of item. You use this string later when
// requesting new items and it must be unique among the other registered item
// and view classes of this collection view. This parameter must not be an
// empty string or `nil`.
//
// # Discussion
// 
// Use this method to register nib files containing prototype items to use in
// your collection view. When you request an item using the
// [ItemWithIdentifierForIndexPath] method, the collection view recycles an
// existing item with the same `identifier` or creates a new one by loading
// the contents of your nib file.
// 
// Because items are recycled to improve performance, it is recommended that
// your custom item classes conform to the [NSCollectionViewElement] protocol.
// You can use the methods of that protocol to prepare your classes for reuse.
// 
// Typically, you register your items when initializing your collection view
// interface. Although you can register new items at any time, you must not
// call the [ItemWithIdentifierForIndexPath] method until after you register
// the corresponding item.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionView/register(_:forItemWithIdentifier:)-90h1i
func (c NSCollectionView) RegisterNibForItemWithIdentifier(nib INSNib, identifier NSUserInterfaceItemIdentifier) {
	objc.Send[objc.ID](c.ID, objc.Sel("registerNib:forItemWithIdentifier:"), nib, objc.String(string(identifier)))
}

// Creates or returns a reusable supplementary view of the specified type.
//
// elementKind: The kind of supplementary view to create. This value is defined by the
// layout object. This parameter must not be an empty string or `nil`.
//
// identifier: The reuse identifier for the specified item. This is the identifier you
// specified when registering the supplementary view. This parameter must not
// be `nil`.
//
// indexPath: The index path specifying the location of the supplementary view. The data
// source object receives this information in its
// [CollectionViewViewForSupplementaryElementOfKindAtIndexPath] method and you
// should just pass it along.
//
// # Return Value
// 
// A view that adopts the [NSCollectionViewElement] protocol.
//
// # Discussion
// 
// This method looks for a recycled supplementary view of the specified type
// and returns it if one exists. If one does not exist, it creates it using
// one of the following techniques:
// 
// - If you used the [RegisterClassForSupplementaryViewOfKindWithIdentifier]
// method to register a class for the identifier, this method instantiates
// your view class and returns it. - If you used the
// [RegisterNibForSupplementaryViewOfKindWithIdentifier] method to register a
// nib file for the identifier, this method loads the view from the nib file
// and returns it.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionView/makeSupplementaryView(ofKind:withIdentifier:for:)
func (c NSCollectionView) MakeSupplementaryViewOfKindWithIdentifierForIndexPath(elementKind NSCollectionViewSupplementaryElementKind, identifier NSUserInterfaceItemIdentifier, indexPath objectivec.IObject) INSView {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("makeSupplementaryViewOfKind:withIdentifier:forIndexPath:"), objc.String(string(elementKind)), objc.String(string(identifier)), indexPath)
	return NSViewFromID(rv)
}

// Registers a class to use when creating new supplementary views in the
// collection view.
//
// viewClass: The view class to use for the supplementary view. This class must be
// descended from [NSView] and must conform to the [NSCollectionViewElement]
// protocol. Specify `nil` to unregister a previously registered class or nib
// file.
//
// kind: The kind of the supplementary view. Layout objects define the kinds of
// supplementary views they support and are responsible for providing
// appropriate strings that you can pass for this parameter. This parameter
// must not be an empty string or `nil`.
//
// identifier: The string that identifies the type of supplementary view. You use this
// string later when requesting new views and it must be unique among the
// other registered item and view classes of this collection view. This
// parameter must not be an empty string or `nil`.
//
// # Discussion
// 
// Use this method to register the classes that represent the supplementary
// views in your collection view. When you request a view using the
// [SupplementaryViewOfKindWithIdentifierForIndexPath] method, the collection
// view recycles an existing view with the same `identifier` and `kind` values
// or creates a new one by instantiating your class and calling the
// [InitWithFrame] method of the resulting object.
// 
// The layout object is responsible for defining the kind of supplementary
// views it supports and how those views are used. For example, the flow
// layout ([NSCollectionViewFlowLayout] class) lets you specify supplementary
// views to act as headers and footers for each section.
// 
// Typically, you register your supplementary views when initializing your
// collection view interface. Although you can register new views at any time,
// you must not call the [SupplementaryViewOfKindWithIdentifierForIndexPath]
// method until after you register the corresponding view.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionView/register(_:forSupplementaryViewOfKind:withIdentifier:)-3dqa
func (c NSCollectionView) RegisterClassForSupplementaryViewOfKindWithIdentifier(viewClass objc.Class, kind NSCollectionViewSupplementaryElementKind, identifier NSUserInterfaceItemIdentifier) {
	objc.Send[objc.ID](c.ID, objc.Sel("registerClass:forSupplementaryViewOfKind:withIdentifier:"), viewClass, objc.String(string(kind)), objc.String(string(identifier)))
}

// Registers a nib file to use when creating supplementary views in the
// collection view.
//
// nib: The nib object containing the supplementary view’s definition. The nib
// file must contain exactly one [NSView] object at the top level and that
// view must conform to the [NSCollectionViewElement] protocol. Specify `nil`
// to unregister a previously registered class or nib file.
//
// kind: The kind of the supplementary view. Layout objects define the kinds of
// supplementary views they support and are responsible for providing
// appropriate strings that you can pass for this parameter. This parameter
// must not be an empty string or `nil`.
//
// identifier: The string that identifies the type of supplementary view. You use this
// string later when requesting new views and it must be unique among the
// other registered item and view classes of this collection view. This
// parameter must not be an empty string or `nil`.
//
// # Discussion
// 
// Use this method to register nib files containing prototype supplementary
// views in your collection view. When you request a view using the
// [SupplementaryViewOfKindWithIdentifierForIndexPath] method, the collection
// view recycles an existing view with the same `identifier` and `kind` values
// or creates a new one by loading the contents of your nib file.
// 
// The layout object is responsible for defining the kind of supplementary
// views it supports and how those views are used. For example, the flow
// layout ([NSCollectionViewFlowLayout] class) lets you specify supplementary
// views to act as headers and footers for each section.
// 
// Typically, you register your supplementary views when initializing your
// collection view interface. Although you can register new views at any time,
// you must not call the [SupplementaryViewOfKindWithIdentifierForIndexPath]
// method until after you register the corresponding view.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionView/register(_:forSupplementaryViewOfKind:withIdentifier:)-7gvf2
func (c NSCollectionView) RegisterNibForSupplementaryViewOfKindWithIdentifier(nib INSNib, kind NSCollectionViewSupplementaryElementKind, identifier NSUserInterfaceItemIdentifier) {
	objc.Send[objc.ID](c.ID, objc.Sel("registerNib:forSupplementaryViewOfKind:withIdentifier:"), nib, objc.String(string(kind)), objc.String(string(identifier)))
}

// Reloads all of the data for the collection view.
//
// # Discussion
// 
// Call this method when the data in your data source object changes or when
// you want to force the collection view to update its contents. When you call
// this method, the collection view discards any currently visible items and
// views and redisplays them. For efficiency, the collection view displays
// only the items and supplementary views that are visible after reloading the
// data. If the collection view’s size changes as a result of reloading the
// data, the collection view adjusts its scrolling offsets accordingly.
// 
// Do not call this method in the middle of animation blocks where items are
// being inserted or deleted. The methods used to insert and delete items
// automatically update the collection view’s contents.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionView/reloadData()
func (c NSCollectionView) ReloadData() {
	objc.Send[objc.ID](c.ID, objc.Sel("reloadData"))
}

// Reloads the data in the specified sections of the collection view.
//
// sections: The indexes of the sections that you want to reload. Specifying `nil` for
// this parameter raises an exception.
//
// # Discussion
// 
// Call this method when the data for the specified sections changes or when
// you want to force the appearance of those sections to be updated. When you
// call this method, the collection view discards visible elements in the
// section along with any cached attributes for those elements. For
// efficiency, it then asks the layout object to provide fresh attributes for
// only the visible items and views and requests new views for those elements.
// 
// Do not call this method in the middle of animation blocks where items are
// being inserted or deleted. The methods used to insert and delete items
// automatically update the collection view’s contents.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionView/reloadSections(_:)
func (c NSCollectionView) ReloadSections(sections foundation.NSIndexSet) {
	objc.Send[objc.ID](c.ID, objc.Sel("reloadSections:"), sections)
}

// Reloads only the specified items.
//
// indexPaths: The index paths of the specific items that you want to reload. Specifying
// `nil` for this parameter raises an exception.
//
// # Discussion
// 
// Call this method to update specific items in your collection view. You call
// this method when the underlying data for those items changes and you want
// to update the visual appearance of those items. When you call this method,
// the collection view discards the specified items and asks your data source
// to provide new ones. For efficiency, the collection view requests only the
// items that are visible.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionView/reloadItems(at:)
func (c NSCollectionView) ReloadItemsAtIndexPaths(indexPaths foundation.INSSet) {
	objc.Send[objc.ID](c.ID, objc.Sel("reloadItemsAtIndexPaths:"), indexPaths)
}

// Returns the number of items in the specified section.
//
// section: The index of the section whose item count you want. This index is 0-based.
//
// # Return Value
// 
// The number of items in the section.
//
// # Discussion
// 
// Use this method to get the number of items currently displayed by the
// collection view for the specified section. Do not call the methods of the
// data source to get this information.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionView/numberOfItems(inSection:)
func (c NSCollectionView) NumberOfItemsInSection(section int) int {
	rv := objc.Send[int](c.ID, objc.Sel("numberOfItemsInSection:"), section)
	return rv
}

// Inserts new items into the collection view at the specified locations.
//
// indexPaths: A set of [NSIndexPath] objects, each of which includes a section and item
// index corresponding to the insertion point of a single item. Specifying
// `nil` for this parameter raises an exception.
// //
// [NSIndexPath]: https://developer.apple.com/documentation/Foundation/NSIndexPath
//
// # Discussion
// 
// After adding new items to your data source object, use this method to
// synchronize those changes with the collection view. Calling this method
// lets the collection view know that it must update its internal data
// structures and possibly update its visual appearance. In response, the
// collection view asks the layout object for information about the new
// objects. If the layout object indicates that the new items should appear
// onscreen, the collection view asks the data source to provide the
// appropriate content, animating that content into position as needed.
// 
// When inserting or deleting multiple sections and items, you can animate all
// of your changes at once using the [PerformBatchUpdatesCompletionHandler]
// method.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionView/insertItems(at:)
func (c NSCollectionView) InsertItemsAtIndexPaths(indexPaths foundation.INSSet) {
	objc.Send[objc.ID](c.ID, objc.Sel("insertItemsAtIndexPaths:"), indexPaths)
}

// Moves an item from one location to another in the collection view.
//
// indexPath: The index path of the item that you want to move. This parameter must not
// be `nil`.
//
// newIndexPath: The index path of the item’s new location. This parameter must not be
// `nil`.
//
// # Discussion
// 
// After rearranging items in your data source object, use this method to
// synchronize those changes with the collection view. Calling this method
// lets the collection view know that it must update its internal data
// structures and possibly update its visual appearance. You can move the item
// to a different section or to a new location in the same section. The
// collection view updates the layout as needed to account for the move,
// animating cells into position in response.
// 
// When inserting or deleting multiple sections and items, you can animate all
// of your changes at once using the [PerformBatchUpdatesCompletionHandler]
// method.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionView/moveItem(at:to:)
func (c NSCollectionView) MoveItemAtIndexPathToIndexPath(indexPath objectivec.IObject, newIndexPath objectivec.IObject) {
	objc.Send[objc.ID](c.ID, objc.Sel("moveItemAtIndexPath:toIndexPath:"), indexPath, newIndexPath)
}

// Deletes the items at the specified index paths.
//
// indexPaths: A set of [NSIndexPath] objects, each of which includes a section and item
// index corresponding to the insertion point of a single item. Specifying
// `nil` for this parameter raises an exception.
// //
// [NSIndexPath]: https://developer.apple.com/documentation/Foundation/NSIndexPath
//
// # Discussion
// 
// After removing items from your data source object, use this method to
// synchronize those changes with the collection view. Calling this method
// lets the collection view know that it must update its internal data
// structures and possibly update its visual appearance. In response, the
// collection view asks the layout object to update the positions of the
// remaining objects. If the layout object indicates that there are changes to
// the visible items, the collection view animates the affected items into
// place.
// 
// When inserting or deleting multiple sections and items, you can animate all
// of your changes at once using the [PerformBatchUpdatesCompletionHandler]
// method.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionView/deleteItems(at:)
func (c NSCollectionView) DeleteItemsAtIndexPaths(indexPaths foundation.INSSet) {
	objc.Send[objc.ID](c.ID, objc.Sel("deleteItemsAtIndexPaths:"), indexPaths)
}

// Inserts new sections at the specified indexes.
//
// sections: An index set containing the indexes at which you want to insert new
// sections. This parameter must not be `nil`.
//
// # Discussion
// 
// This method tells the collection view to insert the specified sections and
// update itself. Always update your data source object before calling this
// method. Calling this method kicks off an update (and possible animations)
// to add the new sections. Specifically, the collection view asks the layout
// object for any updated layout attributes related to the new sections or any
// existing sections. If the layout attributes of any visible items changed,
// those changes are animated into place.
// 
// When inserting or deleting multiple sections and items, you can animate all
// of your changes at once using the [PerformBatchUpdatesCompletionHandler]
// method.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionView/insertSections(_:)
func (c NSCollectionView) InsertSections(sections foundation.NSIndexSet) {
	objc.Send[objc.ID](c.ID, objc.Sel("insertSections:"), sections)
}

// Moves a section from its current location to a new location.
//
// section: The index of the section that you want to move.
//
// newSection: The new index at which to insert the section.
//
// # Discussion
// 
// Use this method to reorganize sections and their contained items. Always
// update your data source object before calling this method. Calling this
// method kicks off an update (and possible animations) to move the specified
// section to its new location. Specifically, the collection view asks the
// layout object for any updated layout attributes related to the new sections
// or any existing sections. If the layout attributes of any visible items
// changed, those changes are animated into place.
// 
// When inserting or deleting multiple sections and items, you can animate all
// of your changes at once using the [PerformBatchUpdatesCompletionHandler]
// method.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionView/moveSection(_:toSection:)
func (c NSCollectionView) MoveSectionToSection(section int, newSection int) {
	objc.Send[objc.ID](c.ID, objc.Sel("moveSection:toSection:"), section, newSection)
}

// Deletes the specified sections and their contained items.
//
// sections: An index set containing the indexes of the sections that you want to
// delete. This parameter must not be `nil`.
//
// # Discussion
// 
// Use this method to delete entire sections and their contained items. Always
// update your data source object before calling this method. Calling this
// method kicks off an update (and possible animations) to delete the
// specified sections. Specifically, the collection view asks the layout
// object for the final layout attributes for any deleted sections and may
// also ask for updated layout attributes for any remaining sections. If the
// layout attributes of any visible items changed, those changes are animated
// into place.
// 
// When inserting or deleting multiple sections and items, you can animate all
// of your changes at once using the [PerformBatchUpdatesCompletionHandler]
// method.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionView/deleteSections(_:)
func (c NSCollectionView) DeleteSections(sections foundation.NSIndexSet) {
	objc.Send[objc.ID](c.ID, objc.Sel("deleteSections:"), sections)
}

// Collapses the section in which the sender resides into a single
// horizontally scrollable row.
//
// sender: The object that requested the action.
//
// # Discussion
// 
// The icon view in Finder offers this type of collapsible section behavior
// when users choose the Show Less and Show More buttons. To enable this
// behavior, your header view must conform to the
// [NSCollectionViewSectionHeaderView] protocol, because the collection view
// uses the `sectionCollapseButton` property to identify the button that
// controls the collapse action.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionView/toggleSectionCollapse(_:)
func (c NSCollectionView) ToggleSectionCollapse(sender objectivec.IObject) {
	objc.Send[objc.ID](c.ID, objc.Sel("toggleSectionCollapse:"), sender)
}

// Deselects all items in the collection view.
//
// sender: The object that requested the action. You may specify `nil` for this
// property.
//
// # Discussion
// 
// This method works only when the [Selectable] and [AllowsEmptySelection]
// properties are both true [true]. If either property is set to [false], this
// method quietly does nothing and any connected menu item is disabled.
// 
// This method consults the delegate object regarding the selection.
// Specifically, it calls the delegate’s
// [CollectionViewShouldDeselectItemsAtIndexPaths] method to see if the items
// should be selected. For any items that are selected, it calls the
// [CollectionViewDidDeselectItemsAtIndexPaths] method.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionView/deselectAll(_:)
func (c NSCollectionView) DeselectAll(sender objectivec.IObject) {
	objc.Send[objc.ID](c.ID, objc.Sel("deselectAll:"), sender)
}

// Adds the specified items to the current selection and optionally scrolls
// the items into position.
//
// indexPaths: The index paths of the items you want to select.
//
// scrollPosition: The options for scrolling the newly selected items into view. You may
// combine one vertical and one horizontal scrolling option when calling this
// method. Specifying more than one option for either the vertical or
// horizontal directions raises an exception.
//
// # Discussion
// 
// Use this method to extend the current selection. If you want to animate the
// selection of the new items, call this method on the collection view’s
// [Animator] proxy object instead. This method does not call any methods of
// the delegate object when making the selection.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionView/selectItems(at:scrollPosition:)
func (c NSCollectionView) SelectItemsAtIndexPathsScrollPosition(indexPaths foundation.INSSet, scrollPosition NSCollectionViewScrollPosition) {
	objc.Send[objc.ID](c.ID, objc.Sel("selectItemsAtIndexPaths:scrollPosition:"), indexPaths, scrollPosition)
}

// Removes the specified items from the current selection.
//
// indexPaths: The index paths of the items you want to deselect.
//
// # Discussion
// 
// Use this method to reduce the current selection. If you want to animate the
// deselection of the new items, call this method on the collection view’s
// [Animator] proxy object instead. This method does not call any methods of
// the delegate object when making the selection.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionView/deselectItems(at:)
func (c NSCollectionView) DeselectItemsAtIndexPaths(indexPaths foundation.INSSet) {
	objc.Send[objc.ID](c.ID, objc.Sel("deselectItemsAtIndexPaths:"), indexPaths)
}

// Returns an array of the actively managed items in the collection view.
//
// # Return Value
// 
// An array of [NSCollectionViewItem] objects. The returned array may be
// empty.
//
// # Discussion
// 
// The items returned by this method represent the ones that are active and
// currently being managed by the collection view. This array may contain
// items that are outside of the collection view’s actual visible rectangle.
// For example, it may contain items that were recently visible but have since
// been scrolled out of view. To test whether an item is actually visible,
// check to see if its frame rectangle intersects the [visibleRect] of the
// collection view.
//
// [visibleRect]: https://developer.apple.com/documentation/AppKit/NSView/visibleRect
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionView/visibleItems()
func (c NSCollectionView) VisibleItems() []NSCollectionViewItem {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("visibleItems"))
	return objc.ConvertSlice(rv, func(id objc.ID) NSCollectionViewItem {
		return NSCollectionViewItemFromID(id)
	})
}

// Returns the index paths of the currently active items.
//
// # Return Value
// 
// The set of [NSIndexPath] objects corresponding to the currently visible
// items.
//
// [NSIndexPath]: https://developer.apple.com/documentation/Foundation/NSIndexPath
//
// # Discussion
// 
// The index paths returned by this method belong to items that are active and
// currently being managed by the collection view. As a result, the returned
// set may include index paths for items that are outside of the collection
// view’s actual visible rectangle. For example, it may contain index paths
// for items that were recently visible but have since been scrolled out of
// view. To test whether an item is visible, check to see if its frame
// rectangle intersects the [visibleRect] of the collection view.
//
// [visibleRect]: https://developer.apple.com/documentation/AppKit/NSView/visibleRect
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionView/indexPathsForVisibleItems()
func (c NSCollectionView) IndexPathsForVisibleItems() foundation.INSSet {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("indexPathsForVisibleItems"))
	return foundation.NSSetFromID(rv)
}

// Returns an array of the actively managed supplementary views in the
// collection view.
//
// elementKind: The kind of the supplementary views you want returned. The layout object
// defines the kinds of supplementary views it supports. This parameter must
// not be `nil`.
//
// # Return Value
// 
// An array of view objects. The returned array may be empty.
//
// # Discussion
// 
// The views returned by this method represent the ones that are active and
// are currently being managed by the collection view. The array may contain
// supplementary views that are outside of the collection view’s actual
// visible rectangle. For example, it might contain views that were recently
// visible but have since been scrolled out of the visible rectangle. To test
// whether a view is actually visible, check to see if its frame rectangle
// intersects the [visibleRect] of the collection view.
//
// [visibleRect]: https://developer.apple.com/documentation/AppKit/NSView/visibleRect
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionView/visibleSupplementaryViews(ofKind:)
func (c NSCollectionView) VisibleSupplementaryViewsOfKind(elementKind NSCollectionViewSupplementaryElementKind) []NSView {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("visibleSupplementaryViewsOfKind:"), objc.String(string(elementKind)))
	return objc.ConvertSlice(rv, func(id objc.ID) NSView {
		return NSViewFromID(id)
	})
}

// Returns the index paths of the currently active supplementary views.
//
// elementKind: The kind of the supplementary views you want returned. The layout object
// defines the kinds of supplementary views it supports. This parameter must
// not be `nil`.
//
// # Return Value
// 
// The set of [NSIndexPath] objects. The returned array may be empty.
//
// [NSIndexPath]: https://developer.apple.com/documentation/Foundation/NSIndexPath
//
// # Discussion
// 
// The index paths returned by this method correspond to supplementary views
// that are active and currently being managed by the collection view. The set
// may include index paths for views that are outside of the collection
// view’s actual visible rectangle. For example, it might contain index
// paths for views that were recently visible but have since been scrolled out
// of the visible rectangle. To test whether a view is actually visible, check
// to see if its frame rectangle intersects the [visibleRect] of the
// collection view.
//
// [visibleRect]: https://developer.apple.com/documentation/AppKit/NSView/visibleRect
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionView/indexPathsForVisibleSupplementaryElements(ofKind:)
func (c NSCollectionView) IndexPathsForVisibleSupplementaryElementsOfKind(elementKind NSCollectionViewSupplementaryElementKind) foundation.INSSet {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("indexPathsForVisibleSupplementaryElementsOfKind:"), objc.String(string(elementKind)))
	return foundation.NSSetFromID(rv)
}

// Returns the index path of the specified item.
//
// item: The item whose index path you want to know.
//
// # Return Value
// 
// The item’s index path or `nil` if the item is not in the collection view.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionView/indexPath(for:)
func (c NSCollectionView) IndexPathForItem(item INSCollectionViewItem) objc.ID {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("indexPathForItem:"), item)
	return rv
}

// Returns the index path of the item at the specified point.
//
// point: The point in the collection view’s bounds that you want to test.
//
// # Return Value
// 
// The item at the specified point or `nil` if no item was found at that
// point.
//
// # Discussion
// 
// This method uses the available layout attributes to determine which item is
// at the specified point. If more than one item is at the point, this method
// returns only the top-most item. This method ignores the opacity of the
// item, so items that are fully transparent are still returned by this
// method. Hidden items are never returned.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionView/indexPathForItem(at:)
func (c NSCollectionView) IndexPathForItemAtPoint(point corefoundation.CGPoint) objc.ID {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("indexPathForItemAtPoint:"), point)
	return rv
}

// Returns the item associated with the specified index path.
//
// indexPath: The index path whose item you want.
//
// # Return Value
// 
// The item for the specified index path or `nil` if no item is available.
//
// # Discussion
// 
// For efficiency, the collection view does not create items until they are
// needed, and usually it creates them only when they need to be displayed
// onscreen. If the collection view does not currently have an item for the
// specified index path, because that item would be positioned offscreen, this
// method returns `nil`.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionView/item(at:)-2vx2h
func (c NSCollectionView) ItemAtIndexPath(indexPath objectivec.IObject) INSCollectionViewItem {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("itemAtIndexPath:"), indexPath)
	return NSCollectionViewItemFromID(rv)
}

// Returns the supplementary view associated with the specified index path.
//
// elementKind: The kind of the supplementary views you want returned. The layout object
// defines the kinds of supplementary views it supports. This parameter must
// not be `nil`.
//
// indexPath: The index path whose supplementary view you want.
//
// # Return Value
// 
// The view for the specified index path or `nil` if no view is available.
//
// # Discussion
// 
// For efficiency, the collection view does not create supplementary views
// until they are needed. Typically, views are created only when they need to
// be displayed onscreen. If the collection view does not currently have a
// supplementary view for the specified index path, because that view would be
// positioned offscreen, this method returns `nil`.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionView/supplementaryView(forElementKind:at:)
func (c NSCollectionView) SupplementaryViewForElementKindAtIndexPath(elementKind NSCollectionViewSupplementaryElementKind, indexPath objectivec.IObject) INSView {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("supplementaryViewForElementKind:atIndexPath:"), objc.String(string(elementKind)), indexPath)
	return NSViewFromID(rv)
}

// Scrolls the collection view contents until the specified items are visible.
//
// indexPaths: The index paths of the items. The layout attributes of these items define
// the bounding box that needs to be scrolled onscreen.
//
// scrollPosition: The options for scrolling the bounding box of the specified items into
// view. You may combine one vertical and one horizontal scrolling option when
// calling this method. Specifying more than one option for either the
// vertical or horizontal directions raises an exception.
//
// # Discussion
// 
// To animate the scrolling operation, call this method on the collection
// view’s [Animator] proxy object instead.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionView/scrollToItems(at:scrollPosition:)
func (c NSCollectionView) ScrollToItemsAtIndexPathsScrollPosition(indexPaths foundation.INSSet, scrollPosition NSCollectionViewScrollPosition) {
	objc.Send[objc.ID](c.ID, objc.Sel("scrollToItemsAtIndexPaths:scrollPosition:"), indexPaths, scrollPosition)
}

// Returns the layout information for the item at the specified index path.
//
// indexPath: The index path of the item.
//
// # Return Value
// 
// The layout attributes of the item or `nil` if no item exists at the
// specified path.
//
// # Discussion
// 
// This method updates the layout information as needed before returning the
// specified attributes. Always use this method to retrieve the layout
// attributes for items in the collection view. Do not query the layout object
// directly.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionView/layoutAttributesForItem(at:)
func (c NSCollectionView) LayoutAttributesForItemAtIndexPath(indexPath objectivec.IObject) INSCollectionViewLayoutAttributes {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("layoutAttributesForItemAtIndexPath:"), indexPath)
	return NSCollectionViewLayoutAttributesFromID(rv)
}

// Returns the layout information for the supplementary view at the specified
// index path.
//
// kind: The kind of the supplementary view whose attributes you want. The layout
// object defines the kinds of supplementary views it supports. This parameter
// must not be `nil`.
//
// indexPath: The index path of the supplementary view. Normally, this path
//
// # Return Value
// 
// The layout attributes of the supplementary view or `nil` if no item exists
// at the specified path.
//
// # Discussion
// 
// This method updates the layout information as needed before returning the
// specified attributes. Always use this method to retrieve the layout
// attributes for supplementary views in the collection view. Do not query the
// layout object directly.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionView/layoutAttributesForSupplementaryElement(ofKind:at:)
func (c NSCollectionView) LayoutAttributesForSupplementaryElementOfKindAtIndexPath(kind NSCollectionViewSupplementaryElementKind, indexPath objectivec.IObject) INSCollectionViewLayoutAttributes {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("layoutAttributesForSupplementaryElementOfKind:atIndexPath:"), objc.String(string(kind)), indexPath)
	return NSCollectionViewLayoutAttributesFromID(rv)
}

// Encapsulates multiple insert, delete, reload, and move operations into a
// single animated operation.
//
// updates: The block that performs the needed inset, delete, reload, or move
// operations. This parameter may be `nil`.
//
// completionHandler: A completion handler block to execute when the changes made in the
// `updates` block have finished animating. This parameter may be `nil`. This
// block takes the following parameters:
// 
// finished: A Boolean value indicating whether the animations completed
// successfully. The value of this parameter is [true] if the animations ran
// to completion or [false] if they were interrupted.
// //
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// Use this method to make multiple changes to the collection view in one
// single animated operation. Normally, when you insert, delete, reload, or
// move items, the collection view animates each change separately. Making
// those same changes inside the `updates` block causes them to be animated at
// the same time. This method updates the current layout information as needed
// before and after performing the operations in the `updates` block.
// 
// The order in which you make calls to insert, delete, or otherwise modify
// the collection view is ignored. When executing your `updates` block, this
// method gathers information about the operations you requested without
// performing those operations. After it gathers that information, the method
// reorders the operations and performs all deletion operations first,
// followed by all insertion operations and then all move operations.
// (Reloading an item is treated as a delete operation followed by an insert
// operation at the same location.) As a result, the indexes you specify for
// each set of operations must reflect the changes made by any preceding
// operations.
// 
// You may call this method from inside your `updates` or `completionHandler`
// blocks.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionView/performBatchUpdates(_:completionHandler:)
func (c NSCollectionView) PerformBatchUpdatesCompletionHandler(updates VoidHandler, completionHandler ErrorHandler) {
		_block0, _cleanup0 := NewVoidBlock(updates)
	defer _cleanup0()
	_block1, _cleanup1 := NewErrorBlock(completionHandler)
	defer _cleanup1()
		objc.Send[objc.ID](c.ID, objc.Sel("performBatchUpdates:completionHandler:"), _block0, _block1)
}

// Returns an image to use for dragging the specified items.
//
// indexPaths: The set of [NSIndexPath] objects corresponding to the items being dragged.
// //
// [NSIndexPath]: https://developer.apple.com/documentation/Foundation/NSIndexPath
//
// event: The mouse-down event that began the drag operation.
//
// dragImageOffset: The offset value to use when positioning the image. On input, the point is
// [NSZeroPoint], which centers the returned image under the mouse. Custom
// implementations can return a different point that repositions the drag
// image by the specified offset values.
// //
// [NSZeroPoint]: https://developer.apple.com/documentation/Foundation/NSZeroPoint
//
// # Return Value
// 
// The image to use for the dragged items.
//
// # Discussion
// 
// The default implementation of this method creates an image using the
// visible portions of the dragged items. The resulting image is a snapshot of
// the dragged items as they currently appear in the collection view but
// rendered with additional transparency to indicate they are part of a drag
// operation. This method also updates the `dragImageOffset` parameter to an
// appropriate point for dragging the resulting image.
// 
// If the delegate implements the
// [CollectionViewDraggingImageForItemsAtIndexPathsWithEventOffset] method,
// the collection view method obtains the drag image from that method instead.
// If the delegate does not implement that method, the collection view uses
// the image returned by this method.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionView/draggingImageForItems(at:with:offset:)-7rc4k
func (c NSCollectionView) DraggingImageForItemsAtIndexPathsWithEventOffset(indexPaths foundation.INSSet, event INSEvent, dragImageOffset foundation.NSPoint) INSImage {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("draggingImageForItemsAtIndexPaths:withEvent:offset:"), indexPaths, event, dragImageOffset)
	return NSImageFromID(rv)
}

// Returns the collection view item for the represented object at the
// specified index.
//
// index: The index of the collection view item.
//
// # Return Value
// 
// An instance of [NSCollectionViewItem].
//
// # Discussion
// 
// Rather than using the [NSCollectionViewItem] instance returned by this
// method to determine the frame of the collection item’s view you should
// use [Content], it is significantly more efficient.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionView/item(at:)-80xze
func (c NSCollectionView) ItemAtIndex(index uint) INSCollectionViewItem {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("itemAtIndex:"), index)
	return NSCollectionViewItemFromID(rv)
}

// Returns the frame of the collection view item at the specified index.
//
// index: The index of the collection view item.
//
// # Return Value
// 
// The frame calculated by the receiver where it intends to place the subview
// for the [NSCollectionViewItem] at the given index. The rectangle is
// returned in the collection view’s coordinate system.
//
// # Discussion
// 
// You can use this method in the
// [CollectionViewDraggingImageForItemsAtIndexesWithEventOffset] method to
// determine which views are in the visible portion of the enclosing scroll
// view.
// 
// Overriding this method will have no effect on the collection view’s
// subview layout.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionView/frameForItem(at:)
func (c NSCollectionView) FrameForItemAtIndex(index uint) corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](c.ID, objc.Sel("frameForItemAtIndex:"), index)
	return corefoundation.CGRect(rv)
}

// Returns the frame of an item based on the number of items in the collection
// view.
//
// index: The index of the item in the collection view.
//
// numberOfItems: The targeted number of items in the collection view. Use this parameter to
// specify the number of items you intend to have in the collection view, if
// that number is different than the actual number of items.
//
// # Return Value
// 
// The frame rectangle that reflects where the collection view would place the
// item.
//
// # Discussion
// 
// Using the value in the `numberOfItems` parameter, this method calculates
// the frame rectangle of the item at the specified `index` in the collection
// view.
// 
// When the collection view is a drag destination, use this method (instead of
// the [Content] method) to get the frame of items. Drag operations can change
// the number of items, which affects the layout of the item views.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionView/frameForItem(at:withNumberOfItems:)
func (c NSCollectionView) FrameForItemAtIndexWithNumberOfItems(index uint, numberOfItems uint) corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](c.ID, objc.Sel("frameForItemAtIndex:withNumberOfItems:"), index, numberOfItems)
	return corefoundation.CGRect(rv)
}

// This method computes and returns an image to use for dragging.
//
// indexes: The index set of the items to be dragged.
//
// event: Mouse drag event.
//
// dragImageOffset: An in/out parameter that will initially be set to [NSZeroPoint]. it can be
// modified to reposition the returned image. A `dragImageOffset` of
// [NSZeroPoint] will cause the image to be centered under the mouse.
// //
// [NSZeroPoint]: https://developer.apple.com/documentation/Foundation/NSZeroPoint
//
// # Return Value
// 
// An image containing a rendering of the visible portions of the views for
// each item.
//
// # Discussion
// 
// You can override the default image by subclassing NSCollectionView and
// overriding this method, or by implementing the
// [CollectionViewDraggingImageForItemsAtIndexesWithEventOffset] delegate
// method, it will be preferred over this method.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionView/draggingImageForItems(at:with:offset:)-951w7
func (c NSCollectionView) DraggingImageForItemsAtIndexesWithEventOffset(indexes foundation.NSIndexSet, event INSEvent, dragImageOffset foundation.NSPoint) INSImage {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("draggingImageForItemsAtIndexes:withEvent:offset:"), indexes, event, dragImageOffset)
	return NSImageFromID(rv)
}

// Configures the drag operation mask.
//
// dragOperationMask: The types of drag operations allowed.
//
// localDestination: If [true], mask applies when the drag destination object is in the same
// application as the receiver; if [false], mask applies when the destination
// object is outside the receiver’s application.
// //
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// This method configures the default value returned from
// [draggingSourceOperationMaskForLocal:]. By default, this method returns
// [DragOperationEvery] when `localDestination` is [true] and
// [NSDragOperationNone] when `localDestination` is [false].
// [NSCollectionView] will save the values you set for each `localDestination`
// value.
// 
// You typically will invoke this method, and not override it.
//
// [NSDragOperationNone]: https://developer.apple.com/documentation/AppKit/NSDragOperation/NSDragOperationNone
// [draggingSourceOperationMaskForLocal:]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/draggingSourceOperationMaskForLocal:
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionView/setDraggingSourceOperationMask(_:forLocal:)
func (c NSCollectionView) SetDraggingSourceOperationMaskForLocal(dragOperationMask NSDragOperation, localDestination bool) {
	objc.Send[objc.ID](c.ID, objc.Sel("setDraggingSourceOperationMask:forLocal:"), dragOperationMask, localDestination)
}

// Declares the types of operations the source allows to be performed.
//
// session: The dragging session.
//
// context: The dragging context. See [NSDraggingContext] for the supported values.
// //
// [NSDraggingContext]: https://developer.apple.com/documentation/AppKit/NSDraggingContext
//
// # Return Value
// 
// The appropriate dragging operation as defined in
//
// # Discussion
// 
// In the future Apple may provide more specific “within” values in the
// future. To account for this, for unrecognized localities, return the
// operation mask for the most specific context that you are concerned with.
// The following code is an example of how to implement this functionality:
//
// See: https://developer.apple.com/documentation/AppKit/NSDraggingSource/draggingSession(_:sourceOperationMaskFor:)
func (c NSCollectionView) DraggingSessionSourceOperationMaskForDraggingContext(session INSDraggingSession, context NSDraggingContext) NSDragOperation {
	rv := objc.Send[NSDragOperation](c.ID, objc.Sel("draggingSession:sourceOperationMaskForDraggingContext:"), session, context)
	return NSDragOperation(rv)
}

// Invoked when the dragging session has completed.
//
// session: The dragging session.
//
// screenPoint: The point where the drag ended, in screen coordinates.
//
// operation: The drag operation. See [NSDragOperation] for drag operation types.
// //
// [NSDragOperation]: https://developer.apple.com/documentation/AppKit/NSDragOperation
//
// See: https://developer.apple.com/documentation/AppKit/NSDraggingSource/draggingSession(_:endedAt:operation:)
func (c NSCollectionView) DraggingSessionEndedAtPointOperation(session INSDraggingSession, screenPoint corefoundation.CGPoint, operation NSDragOperation) {
	objc.Send[objc.ID](c.ID, objc.Sel("draggingSession:endedAtPoint:operation:"), session, screenPoint, operation)
}

// Invoked when the drag moves on the screen.
//
// session: The dragging session.
//
// screenPoint: The point where the drag moved to, in screen coordinates.
//
// See: https://developer.apple.com/documentation/AppKit/NSDraggingSource/draggingSession(_:movedTo:)
func (c NSCollectionView) DraggingSessionMovedToPoint(session INSDraggingSession, screenPoint corefoundation.CGPoint) {
	objc.Send[objc.ID](c.ID, objc.Sel("draggingSession:movedToPoint:"), session, screenPoint)
}

// Invoked when the drag will begin.
//
// session: The dragging session.
//
// screenPoint: The point where the drag will begin, in screen coordinates.
//
// See: https://developer.apple.com/documentation/AppKit/NSDraggingSource/draggingSession(_:willBeginAt:)
func (c NSCollectionView) DraggingSessionWillBeginAtPoint(session INSDraggingSession, screenPoint corefoundation.CGPoint) {
	objc.Send[objc.ID](c.ID, objc.Sel("draggingSession:willBeginAtPoint:"), session, screenPoint)
}

// Returns whether the modifier keys will be ignored for this dragging
// session.
//
// session: The dragging session.
//
// # Return Value
// 
// [true] if the modifier keys will be ignored, [false] otherwise.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSDraggingSource/ignoreModifierKeys(for:)
func (c NSCollectionView) IgnoreModifierKeysForDraggingSession(session INSDraggingSession) bool {
	rv := objc.Send[bool](c.ID, objc.Sel("ignoreModifierKeysForDraggingSession:"), session)
	return rv
}
func (c NSCollectionView) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](c.ID, objc.Sel("encodeWithCoder:"), coder)
}












// An object that provides data for the collection view.
//
// # Discussion
// 
// The data source object manages the data in the collection view. Use this
// object to specify how many items are in the collection view and to create
// the visual representation of those items. The object you specify must adopt
// the [NSCollectionViewDataSource] protocol.
// 
// To specify the data for your collection view, assign a value to this
// property or to the [Content] property, but not both. If you specify a value
// for this property, the collection view ignores the [Content] property and
// the `content` binding.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionView/dataSource
func (c NSCollectionView) DataSource() NSCollectionViewDataSource {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("dataSource"))
	return NSCollectionViewDataSourceObjectFromID(rv)
}
func (c NSCollectionView) SetDataSource(value NSCollectionViewDataSource) {
	objc.Send[struct{}](c.ID, objc.Sel("setDataSource:"), value)
}



// The collection view’s delegate object.
//
// # Discussion
// 
// Use the delegate object to manage the selection and highlighting of items,
// track the addition and removal of items, and manage drag and drop
// operations. The object you assign to this property must conform to the
// [NSCollectionViewDelegate] protocol. The default value of this property is
// `nil`.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionView/delegate
func (c NSCollectionView) Delegate() NSCollectionViewDelegate {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("delegate"))
	return NSCollectionViewDelegateObjectFromID(rv)
}
func (c NSCollectionView) SetDelegate(value NSCollectionViewDelegate) {
	objc.Send[struct{}](c.ID, objc.Sel("setDelegate:"), value)
}



// An array that provides data for the collection view.
//
// # Discussion
// 
// The content object manages the data in the collection view. Use this object
// to specify an array of items directly. This property is observable using
// key-value observing. The collection view also exposes a `content` binding
// value so that you can specify the array of items using an ArrayController
// object in Interface Builder.
// 
// To specify the data for your collection view, assign a value to this
// property or to the [DataSource] property, but not both. If you specify a
// value for the [DataSource] property, the collection view ignores the value
// in this property.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionView/content
func (c NSCollectionView) Content() []objectivec.IObject {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("content"))
	return objc.ConvertSlice(rv, func(id objc.ID) objectivec.IObject {
		return objectivec.Object{ID: id}
	})
}
func (c NSCollectionView) SetContent(value []objectivec.IObject) {
	objc.Send[struct{}](c.ID, objc.Sel("setContent:"), objectivec.IObjectSliceToNSArray(value))
}



// The background view placed behind all items and supplementary views.
//
// # Discussion
// 
// The view you assign to this property is positioned underneath all other
// content and sized automatically to match the enclosing clip view’s frame.
// The view itself does not scroll with the rest of the collection view
// content. The view’s layer redraw policy is also changed to
// [NSView.LayerContentsRedrawPolicy.never].
// 
// In macOS 10.12 and later, a collection view that sets both [BackgroundView]
// and [BackgroundColors] shows `backgroundColors[0]` through all areas that
// are not opaquely covered by the [BackgroundView].
//
// [NSView.LayerContentsRedrawPolicy.never]: https://developer.apple.com/documentation/AppKit/NSView/LayerContentsRedrawPolicy-swift.enum/never
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionView/backgroundView
func (c NSCollectionView) BackgroundView() INSView {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("backgroundView"))
	return NSViewFromID(objc.ID(rv))
}
func (c NSCollectionView) SetBackgroundView(value INSView) {
	objc.Send[struct{}](c.ID, objc.Sel("setBackgroundView:"), value)
}



// An array containing the collection view’s background colors.
//
// # Discussion
// 
// This property contains an array of [NSColor] objects, representing the
// colors to use when drawing the background grid. Specifying an empty array
// or `nil` causes the collection view to use the default colors returned by
// the [controlAlternatingRowBackgroundColors] method.
// 
// When a background view is specified for the collection view, the colors in
// this property are ignored.
//
// [controlAlternatingRowBackgroundColors]: https://developer.apple.com/documentation/AppKit/NSColor/controlAlternatingRowBackgroundColors
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionView/backgroundColors
func (c NSCollectionView) BackgroundColors() []NSColor {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("backgroundColors"))
	return objc.ConvertSlice(rv, func(id objc.ID) NSColor {
		return NSColorFromID(id)
	})
}
func (c NSCollectionView) SetBackgroundColors(value []NSColor) {
	objc.Send[struct{}](c.ID, objc.Sel("setBackgroundColors:"), objectivec.IObjectSliceToNSArray(value))
}



// A Boolean value that indicates whether the collection view’s background
// view scrolls with the items and other content.
//
// # Discussion
// 
// The default value of this property is [false], which means that
// [BackgroundView] (if it exists) fills the collection view’s visible area
// and remains stationary when the collection view’s content is scrolled.
// When the value of this property is [true], [BackgroundView] matches the
// collection view’s frame and scrolls with the collection view’s items
// and other content.
// 
// Changing the value of this property also changes the background view’s
// parent. When [BackgroundView] floats behind the scrolling content, it is a
// sibling of the collection view’s clip view. When it scrolls with the
// collection view’s content, it is a subview of the collection view.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionView/backgroundViewScrollsWithContent
func (c NSCollectionView) BackgroundViewScrollsWithContent() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("backgroundViewScrollsWithContent"))
	return rv
}
func (c NSCollectionView) SetBackgroundViewScrollsWithContent(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setBackgroundViewScrollsWithContent:"), value)
}



// The layout object used to organize the collection view’s content.
//
// # Discussion
// 
// Typically, you specify the layout object at design time in Interface
// Builder. The layout object works with your data source object to generate
// the needed items and views to display. In macOS 10.11 and later, using a
// data source object is recommended, but you may specify `nil` for this
// property if your collection view does not use a data source object. The
// collection view uses the [NSCollectionViewGridLayout] object by default.
// 
// Assigning a new value to this property changes the layout object and causes
// the collection view to update its contents immediately and without
// animations. If you want to animate a layout change, use an animator object
// to set the layout object as follows:
// 
// You can use the completion handler of the associated [NSAnimationContext]
// object to perform additional tasks when the animations finish.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionView/collectionViewLayout
func (c NSCollectionView) CollectionViewLayout() INSCollectionViewLayout {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("collectionViewLayout"))
	return NSCollectionViewLayoutFromID(objc.ID(rv))
}
func (c NSCollectionView) SetCollectionViewLayout(value INSCollectionViewLayout) {
	objc.Send[struct{}](c.ID, objc.Sel("setCollectionViewLayout:"), value)
}



// See: https://developer.apple.com/documentation/AppKit/NSCollectionView/prefetchDataSource
func (c NSCollectionView) PrefetchDataSource() NSCollectionViewPrefetching {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("prefetchDataSource"))
	return NSCollectionViewPrefetchingObjectFromID(rv)
}
func (c NSCollectionView) SetPrefetchDataSource(value NSCollectionViewPrefetching) {
	objc.Send[struct{}](c.ID, objc.Sel("setPrefetchDataSource:"), value)
}



// The number of sections in the collection view.
//
// # Discussion
// 
// This property contains the number of sections reported by the data source
// object. If the collection view does not use a data source object, the value
// in this property is `1`.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionView/numberOfSections
func (c NSCollectionView) NumberOfSections() int {
	rv := objc.Send[int](c.ID, objc.Sel("numberOfSections"))
	return rv
}



// A Boolean value that indicates whether the user may select items in the
// collection view.
//
// # Discussion
// 
// The value of this property is [true] when the collection view allows the
// user to select items, or [false] when it does not. You can set selections
// programmatically regardless of this setting.
// 
// The default value of this property is [false]. Changing the value from
// [true] to [false] removes the current selection if there is one.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionView/isSelectable
func (c NSCollectionView) Selectable() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isSelectable"))
	return rv
}
func (c NSCollectionView) SetSelectable(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setSelectable:"), value)
}



// A Boolean value that indicates whether the user may select more than one
// item in the collection view.
//
// # Discussion
// 
// The value of this property is [true] if the collection view supports the
// selection of more than one item at a time. The default value of this
// property is [false]. Changing the value from [true] to [false] reduces the
// current selection to the first item in the selected group.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionView/allowsMultipleSelection
func (c NSCollectionView) AllowsMultipleSelection() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("allowsMultipleSelection"))
	return rv
}
func (c NSCollectionView) SetAllowsMultipleSelection(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setAllowsMultipleSelection:"), value)
}



// A Boolean value indicating whether the collection view may have no selected
// items.
//
// # Discussion
// 
// The default value of this property is [true], which allows the collection
// view to have no selected items. Setting this property to [false] causes the
// collection view to always leave at least one item selected.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionView/allowsEmptySelection
func (c NSCollectionView) AllowsEmptySelection() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("allowsEmptySelection"))
	return rv
}
func (c NSCollectionView) SetAllowsEmptySelection(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setAllowsEmptySelection:"), value)
}



// The set of index paths representing the currently selected items.
//
// # Discussion
// 
// This property reflects the index paths of the currently selected items,
// where each index path contains a [section] number and an index number for
// the [item] in that section. This property is updated automatically when the
// user selects items interactively. You can also change the selection
// programmatically by assigning a new value to this property. To animate
// changes to the selection, call this method on the collection view’s
// [Animator] proxy object instead.
// 
// It is a programmer error to specify an index path that does not refer to a
// valid item in the data source. If you specify an invalid index path, this
// method raises an exception.
// 
// This property is key-value observable. Other methods that modify the
// selection automatically update this property.
//
// [item]: https://developer.apple.com/documentation/Foundation/NSIndexPath/item
// [section]: https://developer.apple.com/documentation/Foundation/NSIndexPath/section
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionView/selectionIndexPaths
func (c NSCollectionView) SelectionIndexPaths() foundation.INSSet {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("selectionIndexPaths"))
	return foundation.NSSetFromID(objc.ID(rv))
}
func (c NSCollectionView) SetSelectionIndexPaths(value foundation.INSSet) {
	objc.Send[struct{}](c.ID, objc.Sel("setSelectionIndexPaths:"), value)
}



// A Boolean value indicating whether the collection view is the first
// responder.
//
// # Discussion
// 
// The value of this property is [true] when the collection view is the first
// responder. This property is fully key-value observing compliant.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionView/isFirstResponder
func (c NSCollectionView) FirstResponder() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isFirstResponder"))
	return rv
}



// The indexes of the currently selected items.
//
// # Discussion
// 
// Whenever possible, use the [SelectionIndexPaths] property instead of this
// one to set selections.
// 
// This property is observable using key-value observing.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionView/selectionIndexes
func (c NSCollectionView) SelectionIndexes() foundation.NSIndexSet {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("selectionIndexes"))
	return foundation.NSIndexSetFromID(objc.ID(rv))
}
func (c NSCollectionView) SetSelectionIndexes(value foundation.NSIndexSet) {
	objc.Send[struct{}](c.ID, objc.Sel("setSelectionIndexes:"), value)
}







// The element kind string assigned to the attributes object when it
// represents an inter-item gap.
//
// See: https://developer.apple.com/documentation/appkit/nscollectionview/elementkindinteritemgapindicator
func (_NSCollectionViewClass NSCollectionViewClass) ElementKindInterItemGapIndicator() string {
	rv := objc.Send[objc.ID](objc.ID(_NSCollectionViewClass.class), objc.Sel("NSCollectionElementKindInterItemGapIndicator"))
	return foundation.NSStringFromID(rv).String()
}



// A supplementary view that acts as a footer for a given section.
//
// See: https://developer.apple.com/documentation/appkit/nscollectionview/elementkindsectionfooter
func (_NSCollectionViewClass NSCollectionViewClass) ElementKindSectionFooter() string {
	rv := objc.Send[objc.ID](objc.ID(_NSCollectionViewClass.class), objc.Sel("NSCollectionElementKindSectionFooter"))
	return foundation.NSStringFromID(rv).String()
}



// A supplementary view that acts as a header for a given section.
//
// See: https://developer.apple.com/documentation/appkit/nscollectionview/elementkindsectionheader
func (_NSCollectionViewClass NSCollectionViewClass) ElementKindSectionHeader() string {
	rv := objc.Send[objc.ID](objc.ID(_NSCollectionViewClass.class), objc.Sel("NSCollectionElementKindSectionHeader"))
	return foundation.NSStringFromID(rv).String()
}















			// Protocol methods for NSDraggingSource
			




















