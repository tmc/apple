// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// A set of methods that a data source object implements to provide the information and view objects that a collection view requires to present content.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewDataSource
type NSCollectionViewDataSource interface {
	objectivec.IObject

	// Asks your data source object to provide the number of items in the specified section.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewDataSource/collectionView(_:numberOfItemsInSection:)
	CollectionViewNumberOfItemsInSection(collectionView INSCollectionView, section int) int

	// Asks your data source object to provide the item at the specified location in the collection view.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewDataSource/collectionView(_:itemForRepresentedObjectAt:)
	CollectionViewItemForRepresentedObjectAtIndexPath(collectionView INSCollectionView, indexPath objectivec.IObject) INSCollectionViewItem
}

// NSCollectionViewDataSourceObject wraps an existing Objective-C object that conforms to the NSCollectionViewDataSource protocol.
type NSCollectionViewDataSourceObject struct {
	objectivec.Object
}
func (o NSCollectionViewDataSourceObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSCollectionViewDataSourceObjectFromID constructs a [NSCollectionViewDataSourceObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSCollectionViewDataSourceObjectFromID(id objc.ID) NSCollectionViewDataSourceObject {
	return NSCollectionViewDataSourceObject{
		Object: objectivec.ObjectFromID(id),
	}
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
func (o NSCollectionViewDataSourceObject) CollectionViewNumberOfItemsInSection(collectionView INSCollectionView, section int) int {
	
	rv := objc.Send[int](o.ID, objc.Sel("collectionView:numberOfItemsInSection:"), collectionView, section)
	return rv
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
func (o NSCollectionViewDataSourceObject) CollectionViewItemForRepresentedObjectAtIndexPath(collectionView INSCollectionView, indexPath objectivec.IObject) INSCollectionViewItem {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("collectionView:itemForRepresentedObjectAtIndexPath:"), collectionView, indexPath)
	return NSCollectionViewItemFromID(rv)
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
func (o NSCollectionViewDataSourceObject) NumberOfSectionsInCollectionView(collectionView INSCollectionView) int {
	
	rv := objc.Send[int](o.ID, objc.Sel("numberOfSectionsInCollectionView:"), collectionView)
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
func (o NSCollectionViewDataSourceObject) CollectionViewViewForSupplementaryElementOfKindAtIndexPath(collectionView INSCollectionView, kind NSCollectionViewSupplementaryElementKind, indexPath objectivec.IObject) INSView {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("collectionView:viewForSupplementaryElementOfKind:atIndexPath:"), collectionView, objc.String(string(kind)), indexPath)
	return NSViewFromID(rv)
	}

