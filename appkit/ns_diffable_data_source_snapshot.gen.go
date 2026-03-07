// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSDiffableDataSourceSnapshot] class.
var (
	_NSDiffableDataSourceSnapshotClass     NSDiffableDataSourceSnapshotClass
	_NSDiffableDataSourceSnapshotClassOnce sync.Once
)

func getNSDiffableDataSourceSnapshotClass() NSDiffableDataSourceSnapshotClass {
	_NSDiffableDataSourceSnapshotClassOnce.Do(func() {
		_NSDiffableDataSourceSnapshotClass = NSDiffableDataSourceSnapshotClass{class: objc.GetClass("NSDiffableDataSourceSnapshot")}
	})
	return _NSDiffableDataSourceSnapshotClass
}

// GetNSDiffableDataSourceSnapshotClass returns the class object for NSDiffableDataSourceSnapshot.
func GetNSDiffableDataSourceSnapshotClass() NSDiffableDataSourceSnapshotClass {
	return getNSDiffableDataSourceSnapshotClass()
}

type NSDiffableDataSourceSnapshotClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSDiffableDataSourceSnapshotClass) Alloc() NSDiffableDataSourceSnapshot {
	rv := objc.Send[NSDiffableDataSourceSnapshot](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// A representation of the state of the data in a view at a specific point in
// time.
//
// # Overview
// 
// Diffable data sources use to provide data for collection views and table
// views. Through a snapshot, you set up the initial state of the data that
// displays in a view, and later update that data.
// 
// The data in a snapshot is made up of the sections and items you want to
// display, in the specific order you want to display them. You configure what
// to display by adding, deleting, or moving the sections and items.
// 
// To display data in a view using a snapshot:
// 
// - Create a snapshot and populate it with the state of the data you want to
// display. - Apply the snapshot to reflect the changes in the UI.
// 
// You can create and configure a snapshot in one of these ways:
// 
// - Create an empty snapshot, then append sections and items to it. - Get the
// current snapshot by calling the diffable data source’s [Snapshot] method,
// then modify that snapshot to reflect the new state of the data that you
// want to display.
// 
// For example, the following code creates an empty snapshot, and populates it
// with a single section with three items. Then, it applies the snapshot,
// animating the UI updates between the previous state and the new state
// represented in the snapshot.
// 
// For more information, see the diffable data source types:
// 
// - [UICollectionViewDiffableDataSource] - [UITableViewDiffableDataSource] -
// [NSCollectionViewDiffableDataSource]
// 
// # Bridging
// 
// Avoid using this type in Swift code. Only use this type to bridge from
// Objective-C code to Swift code by typecasting from a snapshot reference to
// a snapshot:
//
// [UICollectionViewDiffableDataSource]: https://developer.apple.com/documentation/UIKit/UICollectionViewDiffableDataSource-9tqpa
// [UITableViewDiffableDataSource]: https://developer.apple.com/documentation/UIKit/UITableViewDiffableDataSource-2euir
//
// # Creating a snapshot
//
//   - [NSDiffableDataSourceSnapshot.AppendSectionsWithIdentifiers]: Adds the sections with the specified identifiers to the snapshot.
//   - [NSDiffableDataSourceSnapshot.AppendItemsWithIdentifiersIntoSectionWithIdentifier]: Adds the items with the specified identifiers to the specified section of the snapshot.
//   - [NSDiffableDataSourceSnapshot.AppendItemsWithIdentifiers]: Adds the items with the specified identifiers to the last section of the snapshot.
//
// # Getting item and section metrics
//
//   - [NSDiffableDataSourceSnapshot.NumberOfItems]: The number of items in the snapshot.
//   - [NSDiffableDataSourceSnapshot.NumberOfSections]: The number of sections in the snapshot.
//   - [NSDiffableDataSourceSnapshot.NumberOfItemsInSection]: Returns the number of items in the specified section of the snapshot.
//
// # Identifying items and sections
//
//   - [NSDiffableDataSourceSnapshot.ItemIdentifiers]: The identifiers of all of the items in the snapshot.
//   - [NSDiffableDataSourceSnapshot.SectionIdentifiers]: The identifiers of all of the sections in the snapshot.
//   - [NSDiffableDataSourceSnapshot.IndexOfItemIdentifier]: Returns the index of the item in the snapshot with the specified identifier.
//   - [NSDiffableDataSourceSnapshot.IndexOfSectionIdentifier]: Returns the index of the section of the snapshot with the specified identifier.
//   - [NSDiffableDataSourceSnapshot.ItemIdentifiersInSectionWithIdentifier]: Returns the identifiers of all of the items in the specified section of the snapshot.
//   - [NSDiffableDataSourceSnapshot.SectionIdentifierForSectionContainingItemIdentifier]: Returns the identifier of the section containing the specified item in the snapshot.
//
// # Inserting items and sections
//
//   - [NSDiffableDataSourceSnapshot.InsertItemsWithIdentifiersAfterItemWithIdentifier]: Inserts the provided items immediately after the item with the specified identifier in the snapshot.
//   - [NSDiffableDataSourceSnapshot.InsertItemsWithIdentifiersBeforeItemWithIdentifier]: Inserts the provided items immediately before the item with the specified identifier in the snapshot.
//   - [NSDiffableDataSourceSnapshot.InsertSectionsWithIdentifiersAfterSectionWithIdentifier]: Inserts the provided sections immediately after the section with the specified identifier in the snapshot.
//   - [NSDiffableDataSourceSnapshot.InsertSectionsWithIdentifiersBeforeSectionWithIdentifier]: Inserts the provided sections immediately before the section with the specified identifier in the snapshot.
//
// # Removing items and sections
//
//   - [NSDiffableDataSourceSnapshot.DeleteAllItems]: Deletes all of the items from the snapshot.
//   - [NSDiffableDataSourceSnapshot.DeleteItemsWithIdentifiers]: Deletes the items with the specified identifiers from the snapshot.
//   - [NSDiffableDataSourceSnapshot.DeleteSectionsWithIdentifiers]: Deletes the sections with the specified identifiers from the snapshot.
//
// # Reordering items and sections
//
//   - [NSDiffableDataSourceSnapshot.MoveItemWithIdentifierAfterItemWithIdentifier]: Moves the item from its current position in the snapshot to the position immediately after the specified item.
//   - [NSDiffableDataSourceSnapshot.MoveItemWithIdentifierBeforeItemWithIdentifier]: Moves the item from its current position in the snapshot to the position immediately before the specified item.
//   - [NSDiffableDataSourceSnapshot.MoveSectionWithIdentifierAfterSectionWithIdentifier]: Moves the section from its current position in the snapshot to the position immediately after the specified section.
//   - [NSDiffableDataSourceSnapshot.MoveSectionWithIdentifierBeforeSectionWithIdentifier]: Moves the section from its current position in the snapshot to the position immediately before the specified section.
//
// # Reloading data
//
//   - [NSDiffableDataSourceSnapshot.ReloadItemsWithIdentifiers]: Reloads the data within the specified items in the snapshot.
//   - [NSDiffableDataSourceSnapshot.ReloadSectionsWithIdentifiers]: Reloads the data within the specified sections of the snapshot.
//
// See: https://developer.apple.com/documentation/AppKit/NSDiffableDataSourceSnapshotReference
type NSDiffableDataSourceSnapshot struct {
	objectivec.Object
}

// NSDiffableDataSourceSnapshotFromID constructs a [NSDiffableDataSourceSnapshot] from an objc.ID.
//
// A representation of the state of the data in a view at a specific point in
// time.
func NSDiffableDataSourceSnapshotFromID(id objc.ID) NSDiffableDataSourceSnapshot {
	return NSDiffableDataSourceSnapshot{objectivec.Object{id}}
}
// NOTE: NSDiffableDataSourceSnapshot adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [NSDiffableDataSourceSnapshot] class.
//
// # Creating a snapshot
//
//   - [INSDiffableDataSourceSnapshot.AppendSectionsWithIdentifiers]: Adds the sections with the specified identifiers to the snapshot.
//   - [INSDiffableDataSourceSnapshot.AppendItemsWithIdentifiersIntoSectionWithIdentifier]: Adds the items with the specified identifiers to the specified section of the snapshot.
//   - [INSDiffableDataSourceSnapshot.AppendItemsWithIdentifiers]: Adds the items with the specified identifiers to the last section of the snapshot.
//
// # Getting item and section metrics
//
//   - [INSDiffableDataSourceSnapshot.NumberOfItems]: The number of items in the snapshot.
//   - [INSDiffableDataSourceSnapshot.NumberOfSections]: The number of sections in the snapshot.
//   - [INSDiffableDataSourceSnapshot.NumberOfItemsInSection]: Returns the number of items in the specified section of the snapshot.
//
// # Identifying items and sections
//
//   - [INSDiffableDataSourceSnapshot.ItemIdentifiers]: The identifiers of all of the items in the snapshot.
//   - [INSDiffableDataSourceSnapshot.SectionIdentifiers]: The identifiers of all of the sections in the snapshot.
//   - [INSDiffableDataSourceSnapshot.IndexOfItemIdentifier]: Returns the index of the item in the snapshot with the specified identifier.
//   - [INSDiffableDataSourceSnapshot.IndexOfSectionIdentifier]: Returns the index of the section of the snapshot with the specified identifier.
//   - [INSDiffableDataSourceSnapshot.ItemIdentifiersInSectionWithIdentifier]: Returns the identifiers of all of the items in the specified section of the snapshot.
//   - [INSDiffableDataSourceSnapshot.SectionIdentifierForSectionContainingItemIdentifier]: Returns the identifier of the section containing the specified item in the snapshot.
//
// # Inserting items and sections
//
//   - [INSDiffableDataSourceSnapshot.InsertItemsWithIdentifiersAfterItemWithIdentifier]: Inserts the provided items immediately after the item with the specified identifier in the snapshot.
//   - [INSDiffableDataSourceSnapshot.InsertItemsWithIdentifiersBeforeItemWithIdentifier]: Inserts the provided items immediately before the item with the specified identifier in the snapshot.
//   - [INSDiffableDataSourceSnapshot.InsertSectionsWithIdentifiersAfterSectionWithIdentifier]: Inserts the provided sections immediately after the section with the specified identifier in the snapshot.
//   - [INSDiffableDataSourceSnapshot.InsertSectionsWithIdentifiersBeforeSectionWithIdentifier]: Inserts the provided sections immediately before the section with the specified identifier in the snapshot.
//
// # Removing items and sections
//
//   - [INSDiffableDataSourceSnapshot.DeleteAllItems]: Deletes all of the items from the snapshot.
//   - [INSDiffableDataSourceSnapshot.DeleteItemsWithIdentifiers]: Deletes the items with the specified identifiers from the snapshot.
//   - [INSDiffableDataSourceSnapshot.DeleteSectionsWithIdentifiers]: Deletes the sections with the specified identifiers from the snapshot.
//
// # Reordering items and sections
//
//   - [INSDiffableDataSourceSnapshot.MoveItemWithIdentifierAfterItemWithIdentifier]: Moves the item from its current position in the snapshot to the position immediately after the specified item.
//   - [INSDiffableDataSourceSnapshot.MoveItemWithIdentifierBeforeItemWithIdentifier]: Moves the item from its current position in the snapshot to the position immediately before the specified item.
//   - [INSDiffableDataSourceSnapshot.MoveSectionWithIdentifierAfterSectionWithIdentifier]: Moves the section from its current position in the snapshot to the position immediately after the specified section.
//   - [INSDiffableDataSourceSnapshot.MoveSectionWithIdentifierBeforeSectionWithIdentifier]: Moves the section from its current position in the snapshot to the position immediately before the specified section.
//
// # Reloading data
//
//   - [INSDiffableDataSourceSnapshot.ReloadItemsWithIdentifiers]: Reloads the data within the specified items in the snapshot.
//   - [INSDiffableDataSourceSnapshot.ReloadSectionsWithIdentifiers]: Reloads the data within the specified sections of the snapshot.
//
// See: https://developer.apple.com/documentation/AppKit/NSDiffableDataSourceSnapshotReference
type INSDiffableDataSourceSnapshot interface {
	objectivec.IObject

	// Topic: Creating a snapshot

	// Adds the sections with the specified identifiers to the snapshot.
	AppendSectionsWithIdentifiers(sectionIdentifiers []objectivec.IObject)
	// Adds the items with the specified identifiers to the specified section of the snapshot.
	AppendItemsWithIdentifiersIntoSectionWithIdentifier(identifiers []objectivec.IObject, sectionIdentifier objectivec.IObject)
	// Adds the items with the specified identifiers to the last section of the snapshot.
	AppendItemsWithIdentifiers(identifiers []objectivec.IObject)

	// Topic: Getting item and section metrics

	// The number of items in the snapshot.
	NumberOfItems() int
	// The number of sections in the snapshot.
	NumberOfSections() int
	// Returns the number of items in the specified section of the snapshot.
	NumberOfItemsInSection(sectionIdentifier objectivec.IObject) int

	// Topic: Identifying items and sections

	// The identifiers of all of the items in the snapshot.
	ItemIdentifiers() []objectivec.IObject
	// The identifiers of all of the sections in the snapshot.
	SectionIdentifiers() []objectivec.IObject
	// Returns the index of the item in the snapshot with the specified identifier.
	IndexOfItemIdentifier(itemIdentifier objectivec.IObject) int
	// Returns the index of the section of the snapshot with the specified identifier.
	IndexOfSectionIdentifier(sectionIdentifier objectivec.IObject) int
	// Returns the identifiers of all of the items in the specified section of the snapshot.
	ItemIdentifiersInSectionWithIdentifier(sectionIdentifier objectivec.IObject) []objectivec.IObject
	// Returns the identifier of the section containing the specified item in the snapshot.
	SectionIdentifierForSectionContainingItemIdentifier(itemIdentifier objectivec.IObject) objectivec.IObject

	// Topic: Inserting items and sections

	// Inserts the provided items immediately after the item with the specified identifier in the snapshot.
	InsertItemsWithIdentifiersAfterItemWithIdentifier(identifiers []objectivec.IObject, itemIdentifier objectivec.IObject)
	// Inserts the provided items immediately before the item with the specified identifier in the snapshot.
	InsertItemsWithIdentifiersBeforeItemWithIdentifier(identifiers []objectivec.IObject, itemIdentifier objectivec.IObject)
	// Inserts the provided sections immediately after the section with the specified identifier in the snapshot.
	InsertSectionsWithIdentifiersAfterSectionWithIdentifier(sectionIdentifiers []objectivec.IObject, toSectionIdentifier objectivec.IObject)
	// Inserts the provided sections immediately before the section with the specified identifier in the snapshot.
	InsertSectionsWithIdentifiersBeforeSectionWithIdentifier(sectionIdentifiers []objectivec.IObject, toSectionIdentifier objectivec.IObject)

	// Topic: Removing items and sections

	// Deletes all of the items from the snapshot.
	DeleteAllItems()
	// Deletes the items with the specified identifiers from the snapshot.
	DeleteItemsWithIdentifiers(identifiers []objectivec.IObject)
	// Deletes the sections with the specified identifiers from the snapshot.
	DeleteSectionsWithIdentifiers(sectionIdentifiers []objectivec.IObject)

	// Topic: Reordering items and sections

	// Moves the item from its current position in the snapshot to the position immediately after the specified item.
	MoveItemWithIdentifierAfterItemWithIdentifier(fromIdentifier objectivec.IObject, toIdentifier objectivec.IObject)
	// Moves the item from its current position in the snapshot to the position immediately before the specified item.
	MoveItemWithIdentifierBeforeItemWithIdentifier(fromIdentifier objectivec.IObject, toIdentifier objectivec.IObject)
	// Moves the section from its current position in the snapshot to the position immediately after the specified section.
	MoveSectionWithIdentifierAfterSectionWithIdentifier(fromSectionIdentifier objectivec.IObject, toSectionIdentifier objectivec.IObject)
	// Moves the section from its current position in the snapshot to the position immediately before the specified section.
	MoveSectionWithIdentifierBeforeSectionWithIdentifier(fromSectionIdentifier objectivec.IObject, toSectionIdentifier objectivec.IObject)

	// Topic: Reloading data

	// Reloads the data within the specified items in the snapshot.
	ReloadItemsWithIdentifiers(identifiers []objectivec.IObject)
	// Reloads the data within the specified sections of the snapshot.
	ReloadSectionsWithIdentifiers(sectionIdentifiers []objectivec.IObject)
}





// Init initializes the instance.
func (d NSDiffableDataSourceSnapshot) Init() NSDiffableDataSourceSnapshot {
	rv := objc.Send[NSDiffableDataSourceSnapshot](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d NSDiffableDataSourceSnapshot) Autorelease() NSDiffableDataSourceSnapshot {
	rv := objc.Send[NSDiffableDataSourceSnapshot](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSDiffableDataSourceSnapshot creates a new NSDiffableDataSourceSnapshot instance.
func NewNSDiffableDataSourceSnapshot() NSDiffableDataSourceSnapshot {
	class := getNSDiffableDataSourceSnapshotClass()
	rv := objc.Send[NSDiffableDataSourceSnapshot](objc.ID(class.class), objc.Sel("new"))
	return rv
}










// Adds the sections with the specified identifiers to the snapshot.
//
// sectionIdentifiers: An array of identifiers specifying the sections to add to the snapshot.
//
// See: https://developer.apple.com/documentation/AppKit/NSDiffableDataSourceSnapshotReference/appendSections(withIdentifiers:)
func (d NSDiffableDataSourceSnapshot) AppendSectionsWithIdentifiers(sectionIdentifiers []objectivec.IObject) {
	objc.Send[objc.ID](d.ID, objc.Sel("appendSectionsWithIdentifiers:"), objectivec.IObjectSliceToNSArray(sectionIdentifiers))
}

// Adds the items with the specified identifiers to the specified section of
// the snapshot.
//
// identifiers: An array of identifiers specifying the items to add to the snapshot.
//
// sectionIdentifier: The section to which to add the items.
//
// See: https://developer.apple.com/documentation/AppKit/NSDiffableDataSourceSnapshotReference/appendItems(withIdentifiers:intoSectionWithIdentifier:)
func (d NSDiffableDataSourceSnapshot) AppendItemsWithIdentifiersIntoSectionWithIdentifier(identifiers []objectivec.IObject, sectionIdentifier objectivec.IObject) {
	objc.Send[objc.ID](d.ID, objc.Sel("appendItemsWithIdentifiers:intoSectionWithIdentifier:"), objectivec.IObjectSliceToNSArray(identifiers), sectionIdentifier)
}

// Adds the items with the specified identifiers to the last section of the
// snapshot.
//
// identifiers: An array of identifiers specifying the items to add to the snapshot.
//
// See: https://developer.apple.com/documentation/AppKit/NSDiffableDataSourceSnapshotReference/appendItems(withIdentifiers:)
func (d NSDiffableDataSourceSnapshot) AppendItemsWithIdentifiers(identifiers []objectivec.IObject) {
	objc.Send[objc.ID](d.ID, objc.Sel("appendItemsWithIdentifiers:"), objectivec.IObjectSliceToNSArray(identifiers))
}

// Returns the number of items in the specified section of the snapshot.
//
// sectionIdentifier: The identifier of the section of the snapshot.
//
// # Return Value
// 
// The number of items in the specified section. This method returns `0` if
// the section is empty.
//
// # Discussion
// 
// If you call this method with the identifier of a section that doesn’t
// exist in the snapshot, the app throws an error.
//
// See: https://developer.apple.com/documentation/AppKit/NSDiffableDataSourceSnapshotReference/numberOfItems(inSection:)
func (d NSDiffableDataSourceSnapshot) NumberOfItemsInSection(sectionIdentifier objectivec.IObject) int {
	rv := objc.Send[int](d.ID, objc.Sel("numberOfItemsInSection:"), sectionIdentifier)
	return rv
}

// Returns the index of the item in the snapshot with the specified
// identifier.
//
// itemIdentifier: The identifier of the item in the snapshot.
//
// # Return Value
// 
// The index of the item in the snapshot, or [NSNotFound] if the item with the
// specified identifier doesn’t exist in the snapshot. This index value is
// 0-based.
//
// See: https://developer.apple.com/documentation/AppKit/NSDiffableDataSourceSnapshotReference/index(ofItemIdentifier:)
func (d NSDiffableDataSourceSnapshot) IndexOfItemIdentifier(itemIdentifier objectivec.IObject) int {
	rv := objc.Send[int](d.ID, objc.Sel("indexOfItemIdentifier:"), itemIdentifier)
	return rv
}

// Returns the index of the section of the snapshot with the specified
// identifier.
//
// sectionIdentifier: The identifier of the section of the snapshot.
//
// # Return Value
// 
// The index of the section of the snapshot, or [NSNotFound] if the section
// with the specified identifier doesn’t exist in the snapshot. This index
// value is 0-based.
//
// See: https://developer.apple.com/documentation/AppKit/NSDiffableDataSourceSnapshotReference/index(ofSectionIdentifier:)
func (d NSDiffableDataSourceSnapshot) IndexOfSectionIdentifier(sectionIdentifier objectivec.IObject) int {
	rv := objc.Send[int](d.ID, objc.Sel("indexOfSectionIdentifier:"), sectionIdentifier)
	return rv
}

// Returns the identifiers of all of the items in the specified section of the
// snapshot.
//
// sectionIdentifier: The identifier of the section of the snapshot.
//
// # Return Value
// 
// An array of identifiers of the items contained in the section.
//
// See: https://developer.apple.com/documentation/AppKit/NSDiffableDataSourceSnapshotReference/itemIdentifiersInSection(withIdentifier:)
func (d NSDiffableDataSourceSnapshot) ItemIdentifiersInSectionWithIdentifier(sectionIdentifier objectivec.IObject) []objectivec.IObject {
	rv := objc.Send[[]objc.ID](d.ID, objc.Sel("itemIdentifiersInSectionWithIdentifier:"), sectionIdentifier)
	return objc.ConvertSlice(rv, func(id objc.ID) objectivec.IObject {
		return objectivec.Object{ID: id}
	})
}

// Returns the identifier of the section containing the specified item in the
// snapshot.
//
// itemIdentifier: The identifier of the item contained in the section of the snapshot.
//
// # Return Value
// 
// The identifier of the section containing the specified item, or `nil` if
// the specified item doesn’t exist in any section of the snapshot.
//
// See: https://developer.apple.com/documentation/AppKit/NSDiffableDataSourceSnapshotReference/sectionIdentifier(forSectionContainingItemIdentifier:)
func (d NSDiffableDataSourceSnapshot) SectionIdentifierForSectionContainingItemIdentifier(itemIdentifier objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("sectionIdentifierForSectionContainingItemIdentifier:"), itemIdentifier)
	return objectivec.Object{ID: rv}
}

// Inserts the provided items immediately after the item with the specified
// identifier in the snapshot.
//
// identifiers: The array of identifiers corresponding to the items to add to the snapshot.
//
// itemIdentifier: The identifier of the item after which to insert the new items.
//
// See: https://developer.apple.com/documentation/AppKit/NSDiffableDataSourceSnapshotReference/insertItems(withIdentifiers:afterItemWithIdentifier:)
func (d NSDiffableDataSourceSnapshot) InsertItemsWithIdentifiersAfterItemWithIdentifier(identifiers []objectivec.IObject, itemIdentifier objectivec.IObject) {
	objc.Send[objc.ID](d.ID, objc.Sel("insertItemsWithIdentifiers:afterItemWithIdentifier:"), objectivec.IObjectSliceToNSArray(identifiers), itemIdentifier)
}

// Inserts the provided items immediately before the item with the specified
// identifier in the snapshot.
//
// identifiers: The array of identifiers corresponding to the items to add to the snapshot.
//
// itemIdentifier: The identifier of the item before which to insert the new items.
//
// See: https://developer.apple.com/documentation/AppKit/NSDiffableDataSourceSnapshotReference/insertItems(withIdentifiers:beforeItemWithIdentifier:)
func (d NSDiffableDataSourceSnapshot) InsertItemsWithIdentifiersBeforeItemWithIdentifier(identifiers []objectivec.IObject, itemIdentifier objectivec.IObject) {
	objc.Send[objc.ID](d.ID, objc.Sel("insertItemsWithIdentifiers:beforeItemWithIdentifier:"), objectivec.IObjectSliceToNSArray(identifiers), itemIdentifier)
}

// Inserts the provided sections immediately after the section with the
// specified identifier in the snapshot.
//
// sectionIdentifiers: The array of identifiers corresponding to the sections to add to the
// snapshot.
//
// toSectionIdentifier: The identifier of the section after which to insert the new sections.
//
// See: https://developer.apple.com/documentation/AppKit/NSDiffableDataSourceSnapshotReference/insertSections(withIdentifiers:afterSectionWithIdentifier:)
func (d NSDiffableDataSourceSnapshot) InsertSectionsWithIdentifiersAfterSectionWithIdentifier(sectionIdentifiers []objectivec.IObject, toSectionIdentifier objectivec.IObject) {
	objc.Send[objc.ID](d.ID, objc.Sel("insertSectionsWithIdentifiers:afterSectionWithIdentifier:"), objectivec.IObjectSliceToNSArray(sectionIdentifiers), toSectionIdentifier)
}

// Inserts the provided sections immediately before the section with the
// specified identifier in the snapshot.
//
// sectionIdentifiers: The array of identifiers corresponding to the sections to add to the
// snapshot.
//
// toSectionIdentifier: The identifier of the section before which to insert the new sections.
//
// See: https://developer.apple.com/documentation/AppKit/NSDiffableDataSourceSnapshotReference/insertSections(withIdentifiers:beforeSectionWithIdentifier:)
func (d NSDiffableDataSourceSnapshot) InsertSectionsWithIdentifiersBeforeSectionWithIdentifier(sectionIdentifiers []objectivec.IObject, toSectionIdentifier objectivec.IObject) {
	objc.Send[objc.ID](d.ID, objc.Sel("insertSectionsWithIdentifiers:beforeSectionWithIdentifier:"), objectivec.IObjectSliceToNSArray(sectionIdentifiers), toSectionIdentifier)
}

// Deletes all of the items from the snapshot.
//
// See: https://developer.apple.com/documentation/AppKit/NSDiffableDataSourceSnapshotReference/deleteAllItems()
func (d NSDiffableDataSourceSnapshot) DeleteAllItems() {
	objc.Send[objc.ID](d.ID, objc.Sel("deleteAllItems"))
}

// Deletes the items with the specified identifiers from the snapshot.
//
// identifiers: The array of identifiers corresponding to the items to delete from the
// snapshot.
//
// See: https://developer.apple.com/documentation/AppKit/NSDiffableDataSourceSnapshotReference/deleteItems(withIdentifiers:)
func (d NSDiffableDataSourceSnapshot) DeleteItemsWithIdentifiers(identifiers []objectivec.IObject) {
	objc.Send[objc.ID](d.ID, objc.Sel("deleteItemsWithIdentifiers:"), objectivec.IObjectSliceToNSArray(identifiers))
}

// Deletes the sections with the specified identifiers from the snapshot.
//
// sectionIdentifiers: The array of identifiers corresponding to the sections to delete from the
// snapshot.
//
// See: https://developer.apple.com/documentation/AppKit/NSDiffableDataSourceSnapshotReference/deleteSections(withIdentifiers:)
func (d NSDiffableDataSourceSnapshot) DeleteSectionsWithIdentifiers(sectionIdentifiers []objectivec.IObject) {
	objc.Send[objc.ID](d.ID, objc.Sel("deleteSectionsWithIdentifiers:"), objectivec.IObjectSliceToNSArray(sectionIdentifiers))
}

// Moves the item from its current position in the snapshot to the position
// immediately after the specified item.
//
// fromIdentifier: The identifier of the item to move in the snapshot.
//
// toIdentifier: The identifier of the item after which to move the specified item.
//
// See: https://developer.apple.com/documentation/AppKit/NSDiffableDataSourceSnapshotReference/moveItem(withIdentifier:afterItemWithIdentifier:)
func (d NSDiffableDataSourceSnapshot) MoveItemWithIdentifierAfterItemWithIdentifier(fromIdentifier objectivec.IObject, toIdentifier objectivec.IObject) {
	objc.Send[objc.ID](d.ID, objc.Sel("moveItemWithIdentifier:afterItemWithIdentifier:"), fromIdentifier, toIdentifier)
}

// Moves the item from its current position in the snapshot to the position
// immediately before the specified item.
//
// fromIdentifier: The identifier of the item to move in the snapshot.
//
// toIdentifier: The identifier of the item before which to move the specified item.
//
// See: https://developer.apple.com/documentation/AppKit/NSDiffableDataSourceSnapshotReference/moveItem(withIdentifier:beforeItemWithIdentifier:)
func (d NSDiffableDataSourceSnapshot) MoveItemWithIdentifierBeforeItemWithIdentifier(fromIdentifier objectivec.IObject, toIdentifier objectivec.IObject) {
	objc.Send[objc.ID](d.ID, objc.Sel("moveItemWithIdentifier:beforeItemWithIdentifier:"), fromIdentifier, toIdentifier)
}

// Moves the section from its current position in the snapshot to the position
// immediately after the specified section.
//
// fromSectionIdentifier: The identifier of the section to move in the snapshot.
//
// toSectionIdentifier: The identifier of the section after which to move the specified section.
//
// See: https://developer.apple.com/documentation/AppKit/NSDiffableDataSourceSnapshotReference/moveSection(withIdentifier:afterSectionWithIdentifier:)
func (d NSDiffableDataSourceSnapshot) MoveSectionWithIdentifierAfterSectionWithIdentifier(fromSectionIdentifier objectivec.IObject, toSectionIdentifier objectivec.IObject) {
	objc.Send[objc.ID](d.ID, objc.Sel("moveSectionWithIdentifier:afterSectionWithIdentifier:"), fromSectionIdentifier, toSectionIdentifier)
}

// Moves the section from its current position in the snapshot to the position
// immediately before the specified section.
//
// fromSectionIdentifier: The identifier of the section to move in the snapshot.
//
// toSectionIdentifier: The identifier of the section before which to move the specified section.
//
// See: https://developer.apple.com/documentation/AppKit/NSDiffableDataSourceSnapshotReference/moveSection(withIdentifier:beforeSectionWithIdentifier:)
func (d NSDiffableDataSourceSnapshot) MoveSectionWithIdentifierBeforeSectionWithIdentifier(fromSectionIdentifier objectivec.IObject, toSectionIdentifier objectivec.IObject) {
	objc.Send[objc.ID](d.ID, objc.Sel("moveSectionWithIdentifier:beforeSectionWithIdentifier:"), fromSectionIdentifier, toSectionIdentifier)
}

// Reloads the data within the specified items in the snapshot.
//
// identifiers: The array of identifiers corresponding to the items to reload in the
// snapshot.
//
// See: https://developer.apple.com/documentation/AppKit/NSDiffableDataSourceSnapshotReference/reloadItems(withIdentifiers:)
func (d NSDiffableDataSourceSnapshot) ReloadItemsWithIdentifiers(identifiers []objectivec.IObject) {
	objc.Send[objc.ID](d.ID, objc.Sel("reloadItemsWithIdentifiers:"), objectivec.IObjectSliceToNSArray(identifiers))
}

// Reloads the data within the specified sections of the snapshot.
//
// sectionIdentifiers: The array of identifiers corresponding to the sections to reload in the
// snapshot.
//
// See: https://developer.apple.com/documentation/AppKit/NSDiffableDataSourceSnapshotReference/reloadSections(withIdentifiers:)
func (d NSDiffableDataSourceSnapshot) ReloadSectionsWithIdentifiers(sectionIdentifiers []objectivec.IObject) {
	objc.Send[objc.ID](d.ID, objc.Sel("reloadSectionsWithIdentifiers:"), objectivec.IObjectSliceToNSArray(sectionIdentifiers))
}












// The number of items in the snapshot.
//
// See: https://developer.apple.com/documentation/AppKit/NSDiffableDataSourceSnapshotReference/numberOfItems
func (d NSDiffableDataSourceSnapshot) NumberOfItems() int {
	rv := objc.Send[int](d.ID, objc.Sel("numberOfItems"))
	return rv
}



// The number of sections in the snapshot.
//
// See: https://developer.apple.com/documentation/AppKit/NSDiffableDataSourceSnapshotReference/numberOfSections
func (d NSDiffableDataSourceSnapshot) NumberOfSections() int {
	rv := objc.Send[int](d.ID, objc.Sel("numberOfSections"))
	return rv
}



// The identifiers of all of the items in the snapshot.
//
// See: https://developer.apple.com/documentation/AppKit/NSDiffableDataSourceSnapshotReference/itemIdentifiers
func (d NSDiffableDataSourceSnapshot) ItemIdentifiers() []objectivec.IObject {
	rv := objc.Send[[]objc.ID](d.ID, objc.Sel("itemIdentifiers"))
	return objc.ConvertSlice(rv, func(id objc.ID) objectivec.IObject {
		return objectivec.Object{ID: id}
	})
}



// The identifiers of all of the sections in the snapshot.
//
// See: https://developer.apple.com/documentation/AppKit/NSDiffableDataSourceSnapshotReference/sectionIdentifiers
func (d NSDiffableDataSourceSnapshot) SectionIdentifiers() []objectivec.IObject {
	rv := objc.Send[[]objc.ID](d.ID, objc.Sel("sectionIdentifiers"))
	return objc.ConvertSlice(rv, func(id objc.ID) objectivec.IObject {
		return objectivec.Object{ID: id}
	})
}
























