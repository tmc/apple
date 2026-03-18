// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSCollectionViewDiffableDataSource] class.
var (
	_NSCollectionViewDiffableDataSourceClass     NSCollectionViewDiffableDataSourceClass
	_NSCollectionViewDiffableDataSourceClassOnce sync.Once
)

func getNSCollectionViewDiffableDataSourceClass() NSCollectionViewDiffableDataSourceClass {
	_NSCollectionViewDiffableDataSourceClassOnce.Do(func() {
		_NSCollectionViewDiffableDataSourceClass = NSCollectionViewDiffableDataSourceClass{class: objc.GetClass("NSCollectionViewDiffableDataSource")}
	})
	return _NSCollectionViewDiffableDataSourceClass
}

// GetNSCollectionViewDiffableDataSourceClass returns the class object for NSCollectionViewDiffableDataSource.
func GetNSCollectionViewDiffableDataSourceClass() NSCollectionViewDiffableDataSourceClass {
	return getNSCollectionViewDiffableDataSourceClass()
}

type NSCollectionViewDiffableDataSourceClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSCollectionViewDiffableDataSourceClass) Alloc() NSCollectionViewDiffableDataSource {
	rv := objc.Send[NSCollectionViewDiffableDataSource](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// The object you use to manage data and provide items for a collection view.
//
// # Overview
// 
// A object is a specialized type of data source that works together with your
// collection view object. It provides the behavior you need to manage updates
// to your collection view’s data and UI in a simple, efficient way. It also
// conforms to the [NSCollectionViewDataSource] protocol and provides
// implementations for all of the protocol’s methods.
// 
// To fill a collection view with data:
// 
// - Connect a diffable data source to your collection view. - Implement an
// item provider to configure your collection view’s items. - Generate the
// current state of the data. - Display the data in the UI.
// 
// To connect a diffable data source to a collection view, you create the
// diffable data source using its [NSCollectionViewDiffableDataSource.InitWithCollectionViewItemProvider]
// initializer, passing in the collection view you want to associate with that
// data source. You also pass in an item provider, where you configure each of
// your items to determine how to display your data in the UI.
// 
// Then, you generate the current state of the data and display the data in
// the UI by constructing and applying a snapshot. For more information, see
// [NSDiffableDataSourceSnapshot].
//
// # Creating a Diffable Data Source
//
//   - [NSCollectionViewDiffableDataSource.InitWithCollectionViewItemProvider]: Creates a diffable data source with the specified item provider, and connects it to the specified collection view.
//
// # Creating Supplementary Views
//
//   - [NSCollectionViewDiffableDataSource.SupplementaryViewProvider]: The closure that configures and returns the collection view’s supplementary views, such as headers and footers, from the diffable data source.
//   - [NSCollectionViewDiffableDataSource.SetSupplementaryViewProvider]
//
// # Identifying Items
//
//   - [NSCollectionViewDiffableDataSource.ItemIdentifierForIndexPath]: Returns an identifier for the item at the specified index path in the collection view.
//   - [NSCollectionViewDiffableDataSource.IndexPathForItemIdentifier]: Returns an index path for the item with the specified identifier in the collection view.
//
// # Updating Data
//
//   - [NSCollectionViewDiffableDataSource.Snapshot]: Returns a representation of the current state of the data in the collection view.
//   - [NSCollectionViewDiffableDataSource.ApplySnapshotAnimatingDifferences]: Updates the UI to reflect the state of the data in the specified snapshot, optionally animating the UI changes.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewDiffableDataSourceReference
type NSCollectionViewDiffableDataSource struct {
	objectivec.Object
}

// NSCollectionViewDiffableDataSourceFromID constructs a [NSCollectionViewDiffableDataSource] from an objc.ID.
//
// The object you use to manage data and provide items for a collection view.
func NSCollectionViewDiffableDataSourceFromID(id objc.ID) NSCollectionViewDiffableDataSource {
	return NSCollectionViewDiffableDataSource{objectivec.Object{ID: id}}
}
// NOTE: NSCollectionViewDiffableDataSource adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSCollectionViewDiffableDataSource] class.
//
// # Creating a Diffable Data Source
//
//   - [INSCollectionViewDiffableDataSource.InitWithCollectionViewItemProvider]: Creates a diffable data source with the specified item provider, and connects it to the specified collection view.
//
// # Creating Supplementary Views
//
//   - [INSCollectionViewDiffableDataSource.SupplementaryViewProvider]: The closure that configures and returns the collection view’s supplementary views, such as headers and footers, from the diffable data source.
//   - [INSCollectionViewDiffableDataSource.SetSupplementaryViewProvider]
//
// # Identifying Items
//
//   - [INSCollectionViewDiffableDataSource.ItemIdentifierForIndexPath]: Returns an identifier for the item at the specified index path in the collection view.
//   - [INSCollectionViewDiffableDataSource.IndexPathForItemIdentifier]: Returns an index path for the item with the specified identifier in the collection view.
//
// # Updating Data
//
//   - [INSCollectionViewDiffableDataSource.Snapshot]: Returns a representation of the current state of the data in the collection view.
//   - [INSCollectionViewDiffableDataSource.ApplySnapshotAnimatingDifferences]: Updates the UI to reflect the state of the data in the specified snapshot, optionally animating the UI changes.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewDiffableDataSourceReference
type INSCollectionViewDiffableDataSource interface {
	objectivec.IObject
	NSCollectionViewDataSource

	// Topic: Creating a Diffable Data Source

	// Creates a diffable data source with the specified item provider, and connects it to the specified collection view.
	InitWithCollectionViewItemProvider(collectionView INSCollectionView, itemProvider NSCollectionViewDiffableDataSourceItemProvider) NSCollectionViewDiffableDataSource

	// Topic: Creating Supplementary Views

	// The closure that configures and returns the collection view’s supplementary views, such as headers and footers, from the diffable data source.
	SupplementaryViewProvider() NSCollectionViewDiffableDataSourceSupplementaryViewProvider
	SetSupplementaryViewProvider(value NSCollectionViewDiffableDataSourceSupplementaryViewProvider)

	// Topic: Identifying Items

	// Returns an identifier for the item at the specified index path in the collection view.
	ItemIdentifierForIndexPath(indexPath objectivec.IObject) objectivec.IObject
	// Returns an index path for the item with the specified identifier in the collection view.
	IndexPathForItemIdentifier(identifier objectivec.IObject) objc.ID

	// Topic: Updating Data

	// Returns a representation of the current state of the data in the collection view.
	Snapshot() INSDiffableDataSourceSnapshot
	// Updates the UI to reflect the state of the data in the specified snapshot, optionally animating the UI changes.
	ApplySnapshotAnimatingDifferences(snapshot INSDiffableDataSourceSnapshot, animatingDifferences bool)
}

// Init initializes the instance.
func (c NSCollectionViewDiffableDataSource) Init() NSCollectionViewDiffableDataSource {
	rv := objc.Send[NSCollectionViewDiffableDataSource](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c NSCollectionViewDiffableDataSource) Autorelease() NSCollectionViewDiffableDataSource {
	rv := objc.Send[NSCollectionViewDiffableDataSource](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSCollectionViewDiffableDataSource creates a new NSCollectionViewDiffableDataSource instance.
func NewNSCollectionViewDiffableDataSource() NSCollectionViewDiffableDataSource {
	class := getNSCollectionViewDiffableDataSourceClass()
	rv := objc.Send[NSCollectionViewDiffableDataSource](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a diffable data source with the specified item provider, and
// connects it to the specified collection view.
//
// collectionView: The initialized collection view object to connect to the diffable data
// source.
//
// itemProvider: A closure that creates and returns each of the items for the collection
// view from the data the diffable data source provides.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewDiffableDataSourceReference/init(collectionView:itemProvider:)
func NewCollectionViewDiffableDataSourceWithCollectionViewItemProvider(collectionView INSCollectionView, itemProvider NSCollectionViewDiffableDataSourceItemProvider) NSCollectionViewDiffableDataSource {
	instance := getNSCollectionViewDiffableDataSourceClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCollectionView:itemProvider:"), collectionView, itemProvider)
	return NSCollectionViewDiffableDataSourceFromID(rv)
}

// Creates a diffable data source with the specified item provider, and
// connects it to the specified collection view.
//
// collectionView: The initialized collection view object to connect to the diffable data
// source.
//
// itemProvider: A closure that creates and returns each of the items for the collection
// view from the data the diffable data source provides.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewDiffableDataSourceReference/init(collectionView:itemProvider:)
func (c NSCollectionViewDiffableDataSource) InitWithCollectionViewItemProvider(collectionView INSCollectionView, itemProvider NSCollectionViewDiffableDataSourceItemProvider) NSCollectionViewDiffableDataSource {
	rv := objc.Send[NSCollectionViewDiffableDataSource](c.ID, objc.Sel("initWithCollectionView:itemProvider:"), collectionView, itemProvider)
	return rv
}

// Returns an identifier for the item at the specified index path in the
// collection view.
//
// indexPath: The index path of the item in the collection view.
//
// # Return Value
// 
// The item’s identifier, or `nil` if no item is found at the provided index
// path.
//
// # Discussion
// 
// This method is a constant time operation, O(1), which means you can look up
// an item identifier from its corresponding index path with no significant
// overhead.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewDiffableDataSourceReference/itemIdentifier(for:)
func (c NSCollectionViewDiffableDataSource) ItemIdentifierForIndexPath(indexPath objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("itemIdentifierForIndexPath:"), indexPath)
	return objectivec.Object{ID: rv}
}

// Returns an index path for the item with the specified identifier in the
// collection view.
//
// identifier: The identifier of the item in the collection view.
//
// # Return Value
// 
// The item’s index path, or `nil` if no item is found with the provided
// item identifier.
//
// # Discussion
// 
// This method is a constant time operation, O(1), which means you can look up
// an index path from its corresponding item identifier with no significant
// overhead.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewDiffableDataSourceReference/indexPath(forItemIdentifier:)
func (c NSCollectionViewDiffableDataSource) IndexPathForItemIdentifier(identifier objectivec.IObject) objc.ID {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("indexPathForItemIdentifier:"), identifier)
	return rv
}

// Returns a representation of the current state of the data in the collection
// view.
//
// # Return Value
// 
// A snapshot containing section and item identifiers in the order that they
// appear in the UI.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewDiffableDataSourceReference/snapshot()
func (c NSCollectionViewDiffableDataSource) Snapshot() INSDiffableDataSourceSnapshot {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("snapshot"))
	return NSDiffableDataSourceSnapshotFromID(rv)
}

// Updates the UI to reflect the state of the data in the specified snapshot,
// optionally animating the UI changes.
//
// snapshot: The snapshot reflecting the new state of the data in the collection view.
//
// animatingDifferences: If [true], the diffable data source computes the difference between the
// collection view’s current state and the new state in the snapshot, which
// is an O() operation, where is the number of items in the snapshot. The
// differences in the UI between the current state and new state are animated.
// If [false], the collection view UI is set to the new state without any
// animations, with no additional overhead for computing a diff. Any ongoing
// item animations are interrupted and the collection view’s content is
// reloaded immediately.
// //
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// It’s safe to call this method from a background queue, but you must do so
// consistently in your app. Always call this method exclusively from the main
// queue or from a background queue.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewDiffableDataSourceReference/applySnapshot(_:animatingDifferences:)
func (c NSCollectionViewDiffableDataSource) ApplySnapshotAnimatingDifferences(snapshot INSDiffableDataSourceSnapshot, animatingDifferences bool) {
	objc.Send[objc.ID](c.ID, objc.Sel("applySnapshot:animatingDifferences:"), snapshot, animatingDifferences)
}

// Asks your data source object to provide the item at the specified location
// in the collection view.
//
// collectionView: The collection view requesting the information.
//
// indexPath: The index path that specifies the location of the item. This index path
// contains both the section index and the item index within that section.
//
// # Return Value
// 
// A configured item object. You must not return `nil` from this method.
//
// # Discussion
// 
// All data source objects must implement this method. Your implementation is
// responsible for creating, configuring, and returning the appropriate item
// object based on the specified index path. You do this by calling the
// [ItemWithIdentifierForIndexPath] method of the collection view to retrieve
// an empty item object of the appropriate type. After receiving the item
// object, update its properties with the data from your app’s data
// structures and return it.
// 
// You do not need to set the frame of an item’s view from this method. The
// collection view gets the item’s location and other layout-related
// attributes from the layout object during a separate step.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewDataSource/collectionView(_:itemForRepresentedObjectAt:)
func (c NSCollectionViewDiffableDataSource) CollectionViewItemForRepresentedObjectAtIndexPath(collectionView INSCollectionView, indexPath objectivec.IObject) INSCollectionViewItem {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("collectionView:itemForRepresentedObjectAtIndexPath:"), collectionView, indexPath)
	return NSCollectionViewItemFromID(rv)
}

// Asks your data source object to provide the number of items in the
// specified section.
//
// collectionView: The collection view requesting the information.
//
// section: The index number of the section. Section indexes are zero based.
//
// # Return Value
// 
// The number of items in the specified section.
//
// # Discussion
// 
// All data source objects must implement this method. Your implementation
// should quickly return the number of items in the specified section.
// 
// Make sure the number of items you return is accurate. The
// [CollectionViewItemForRepresentedObjectAtIndexPath] method of your data
// source object must be able to provide a visual representation for each item
// in the section.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewDataSource/collectionView(_:numberOfItemsInSection:)
func (c NSCollectionViewDiffableDataSource) CollectionViewNumberOfItemsInSection(collectionView INSCollectionView, section int) int {
	rv := objc.Send[int](c.ID, objc.Sel("collectionView:numberOfItemsInSection:"), collectionView, section)
	return rv
}

// Asks your data source object to provide the supplementary view at the
// specified location in a section of the collection view.
//
// collectionView: The collection view requesting the information.
//
// kind: The kind of supplementary view to provide. The value of this string is
// defined by the current layout object associated with the collection view.
// Layouts may define additional views to add visual content that is unrelated
// to specific items.
//
// indexPath: The index path that identifies the section in which to place the
// supplementary view.
//
// # Return Value
// 
// A configured view object. You must not return `nil` from this method.
//
// # Discussion
// 
// Implement this method if the collection view’s layout object supports
// supplementary views. Your implementation is responsible for creating,
// configuring, and returning an appropriate view. You do this by calling the
// [SupplementaryViewOfKindWithIdentifierForIndexPath] method of the
// collection view to retrieve an unconfigured view of the appropriate type.
// After receiving the view, update its properties and content using your
// app’s data structures and return it.
// 
// You do not need to set the location of supplementary views inside the
// collection view’s bounds. The collection view gets the view’s location
// and other layout-related attributes from the layout object during a
// separate step.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewDataSource/collectionView(_:viewForSupplementaryElementOfKind:at:)
func (c NSCollectionViewDiffableDataSource) CollectionViewViewForSupplementaryElementOfKindAtIndexPath(collectionView INSCollectionView, kind NSCollectionViewSupplementaryElementKind, indexPath objectivec.IObject) INSView {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("collectionView:viewForSupplementaryElementOfKind:atIndexPath:"), collectionView, objc.String(string(kind)), indexPath)
	return NSViewFromID(rv)
}

// Asks your data source object to provide the total number of sections.
//
// collectionView: The collection view requesting the information.
//
// # Return Value
// 
// The number of sections in the specified collection view.
//
// # Discussion
// 
// Implement this method when the organization of your data requires more than
// one section. If you do not implement this method, the collection view
// creates only one section.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewDataSource/numberOfSections(in:)
func (c NSCollectionViewDiffableDataSource) NumberOfSectionsInCollectionView(collectionView INSCollectionView) int {
	rv := objc.Send[int](c.ID, objc.Sel("numberOfSectionsInCollectionView:"), collectionView)
	return rv
}

// The closure that configures and returns the collection view’s
// supplementary views, such as headers and footers, from the diffable data
// source.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewDiffableDataSourceReference/supplementaryViewProvider
func (c NSCollectionViewDiffableDataSource) SupplementaryViewProvider() NSCollectionViewDiffableDataSourceSupplementaryViewProvider {
	rv := objc.Send[NSCollectionViewDiffableDataSourceSupplementaryViewProvider](c.ID, objc.Sel("supplementaryViewProvider"))
	return NSCollectionViewDiffableDataSourceSupplementaryViewProvider(rv)
}
func (c NSCollectionViewDiffableDataSource) SetSupplementaryViewProvider(value NSCollectionViewDiffableDataSourceSupplementaryViewProvider) {
	objc.Send[struct{}](c.ID, objc.Sel("setSupplementaryViewProvider:"), value)
}

			// Protocol methods for NSCollectionViewDataSource
			

