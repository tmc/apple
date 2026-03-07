// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSCollectionViewUpdateItem] class.
var (
	_NSCollectionViewUpdateItemClass     NSCollectionViewUpdateItemClass
	_NSCollectionViewUpdateItemClassOnce sync.Once
)

func getNSCollectionViewUpdateItemClass() NSCollectionViewUpdateItemClass {
	_NSCollectionViewUpdateItemClassOnce.Do(func() {
		_NSCollectionViewUpdateItemClass = NSCollectionViewUpdateItemClass{class: objc.GetClass("NSCollectionViewUpdateItem")}
	})
	return _NSCollectionViewUpdateItemClass
}

// GetNSCollectionViewUpdateItemClass returns the class object for NSCollectionViewUpdateItem.
func GetNSCollectionViewUpdateItemClass() NSCollectionViewUpdateItemClass {
	return getNSCollectionViewUpdateItemClass()
}

type NSCollectionViewUpdateItemClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSCollectionViewUpdateItemClass) Alloc() NSCollectionViewUpdateItem {
	rv := objc.Send[NSCollectionViewUpdateItem](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// A description of a single change to make to an item in a collection view.
//
// # Overview
// 
// You do not create instances of this class directly. When updating its
// content, the collection view object creates them and passes them to the
// layout object’s [PrepareForCollectionViewUpdates] method, which can use
// them to prepare for the upcoming changes.
//
// # Accessing the Item Changes
//
//   - [NSCollectionViewUpdateItem.IndexPathBeforeUpdate]: The index path of the item before the update.
//   - [NSCollectionViewUpdateItem.IndexPathAfterUpdate]: The index path of the item after the update.
//   - [NSCollectionViewUpdateItem.UpdateAction]: The action being performed on the item.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewUpdateItem
type NSCollectionViewUpdateItem struct {
	objectivec.Object
}

// NSCollectionViewUpdateItemFromID constructs a [NSCollectionViewUpdateItem] from an objc.ID.
//
// A description of a single change to make to an item in a collection view.
func NSCollectionViewUpdateItemFromID(id objc.ID) NSCollectionViewUpdateItem {
	return NSCollectionViewUpdateItem{objectivec.Object{id}}
}
// NOTE: NSCollectionViewUpdateItem adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [NSCollectionViewUpdateItem] class.
//
// # Accessing the Item Changes
//
//   - [INSCollectionViewUpdateItem.IndexPathBeforeUpdate]: The index path of the item before the update.
//   - [INSCollectionViewUpdateItem.IndexPathAfterUpdate]: The index path of the item after the update.
//   - [INSCollectionViewUpdateItem.UpdateAction]: The action being performed on the item.
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewUpdateItem
type INSCollectionViewUpdateItem interface {
	objectivec.IObject

	// Topic: Accessing the Item Changes

	// The index path of the item before the update.
	IndexPathBeforeUpdate() objc.ID
	// The index path of the item after the update.
	IndexPathAfterUpdate() objc.ID
	// The action being performed on the item.
	UpdateAction() NSCollectionUpdateAction
}





// Init initializes the instance.
func (c NSCollectionViewUpdateItem) Init() NSCollectionViewUpdateItem {
	rv := objc.Send[NSCollectionViewUpdateItem](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c NSCollectionViewUpdateItem) Autorelease() NSCollectionViewUpdateItem {
	rv := objc.Send[NSCollectionViewUpdateItem](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSCollectionViewUpdateItem creates a new NSCollectionViewUpdateItem instance.
func NewNSCollectionViewUpdateItem() NSCollectionViewUpdateItem {
	class := getNSCollectionViewUpdateItemClass()
	rv := objc.Send[NSCollectionViewUpdateItem](objc.ID(class.class), objc.Sel("new"))
	return rv
}





















// The index path of the item before the update.
//
// # Discussion
// 
// The value of this property is `nil` for an action of type
// [NSCollectionView.UpdateAction.insert].
//
// [NSCollectionView.UpdateAction.insert]: https://developer.apple.com/documentation/AppKit/NSCollectionView/UpdateAction/insert
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewUpdateItem/indexPathBeforeUpdate
func (c NSCollectionViewUpdateItem) IndexPathBeforeUpdate() objc.ID {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("indexPathBeforeUpdate"))
	return rv
}



// The index path of the item after the update.
//
// # Discussion
// 
// The value of this property is `nil` for an action of type
// [NSCollectionView.UpdateAction.delete].
//
// [NSCollectionView.UpdateAction.delete]: https://developer.apple.com/documentation/AppKit/NSCollectionView/UpdateAction/delete
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewUpdateItem/indexPathAfterUpdate
func (c NSCollectionViewUpdateItem) IndexPathAfterUpdate() objc.ID {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("indexPathAfterUpdate"))
	return rv
}



// The action being performed on the item.
//
// # Discussion
// 
// For a list of relevant actions, see [NSCollectionView.UpdateAction].
//
// [NSCollectionView.UpdateAction]: https://developer.apple.com/documentation/AppKit/NSCollectionView/UpdateAction
//
// See: https://developer.apple.com/documentation/AppKit/NSCollectionViewUpdateItem/updateAction
func (c NSCollectionViewUpdateItem) UpdateAction() NSCollectionUpdateAction {
	rv := objc.Send[NSCollectionUpdateAction](c.ID, objc.Sel("updateAction"))
	return NSCollectionUpdateAction(rv)
}
























