// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// NSCollectionViewPrefetching protocol.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewPrefetching
type NSCollectionViewPrefetching interface {
	objectivec.IObject

	// CollectionViewPrefetchItemsAtIndexPaths protocol.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewPrefetching/collectionView(_:prefetchItemsAt:)
	CollectionViewPrefetchItemsAtIndexPaths(collectionView INSCollectionView, indexPaths []objc.ID)
}

// NSCollectionViewPrefetchingObject wraps an existing Objective-C object that conforms to the NSCollectionViewPrefetching protocol.
type NSCollectionViewPrefetchingObject struct {
	objectivec.Object
}
func (o NSCollectionViewPrefetchingObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSCollectionViewPrefetchingObjectFromID constructs a [NSCollectionViewPrefetchingObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSCollectionViewPrefetchingObjectFromID(id objc.ID) NSCollectionViewPrefetchingObject {
	return NSCollectionViewPrefetchingObject{
		Object: objectivec.ObjectFromID(id),
	}
}

//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewPrefetching/collectionView(_:prefetchItemsAt:)
func (o NSCollectionViewPrefetchingObject) CollectionViewPrefetchItemsAtIndexPaths(collectionView INSCollectionView, indexPaths []objc.ID) {
	
	objc.Send[struct{}](o.ID, objc.Sel("collectionView:prefetchItemsAtIndexPaths:"), collectionView, objectivec.IDSliceToNSArray(indexPaths))
	}
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewPrefetching/collectionView(_:cancelPrefetchingForItemsAt:)
func (o NSCollectionViewPrefetchingObject) CollectionViewCancelPrefetchingForItemsAtIndexPaths(collectionView INSCollectionView, indexPaths []objc.ID) {
	
	objc.Send[struct{}](o.ID, objc.Sel("collectionView:cancelPrefetchingForItemsAtIndexPaths:"), collectionView, objectivec.IDSliceToNSArray(indexPaths))
	}

