// Code generated from Apple documentation for FSKit. DO NOT EDIT.

package fskit

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [FSDirectoryEntryPacker] class.
var (
	_FSDirectoryEntryPackerClass     FSDirectoryEntryPackerClass
	_FSDirectoryEntryPackerClassOnce sync.Once
)

func getFSDirectoryEntryPackerClass() FSDirectoryEntryPackerClass {
	_FSDirectoryEntryPackerClassOnce.Do(func() {
		_FSDirectoryEntryPackerClass = FSDirectoryEntryPackerClass{class: objc.GetClass("FSDirectoryEntryPacker")}
	})
	return _FSDirectoryEntryPackerClass
}

// GetFSDirectoryEntryPackerClass returns the class object for FSDirectoryEntryPacker.
func GetFSDirectoryEntryPackerClass() FSDirectoryEntryPackerClass {
	return getFSDirectoryEntryPackerClass()
}

type FSDirectoryEntryPackerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (fc FSDirectoryEntryPackerClass) Class() objc.Class {
	return fc.class
}

// Alloc allocates memory for a new instance of the class.
func (fc FSDirectoryEntryPackerClass) Alloc() FSDirectoryEntryPacker {
	rv := objc.Send[FSDirectoryEntryPacker](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// An object used to provide items during a directory enumeration.
//
// # Overview
//
// You use this type in your implementation of
// [EnumerateDirectoryStartingAtCookieVerifierProvidingAttributesUsingPackerReplyHandler].
//
// Packing allows your implementation to provide information FSKit needs,
// including each item’s name, type, and identifier (such as an inode
// number). Some directory enumerations require other attributes, as indicated
// by the [FSItemGetAttributesRequest] sent to the enumerate method.
//
// # Packing entries
//
//   - [FSDirectoryEntryPacker.PackEntryWithNameItemTypeItemIDNextCookieAttributes]: Provides a directory entry during enumeration.
//
// See: https://developer.apple.com/documentation/FSKit/FSDirectoryEntryPacker
type FSDirectoryEntryPacker struct {
	objectivec.Object
}

// FSDirectoryEntryPackerFromID constructs a [FSDirectoryEntryPacker] from an objc.ID.
//
// An object used to provide items during a directory enumeration.
func FSDirectoryEntryPackerFromID(id objc.ID) FSDirectoryEntryPacker {
	return FSDirectoryEntryPacker{objectivec.Object{ID: id}}
}

// NOTE: FSDirectoryEntryPacker adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [FSDirectoryEntryPacker] class.
//
// # Packing entries
//
//   - [IFSDirectoryEntryPacker.PackEntryWithNameItemTypeItemIDNextCookieAttributes]: Provides a directory entry during enumeration.
//
// See: https://developer.apple.com/documentation/FSKit/FSDirectoryEntryPacker
type IFSDirectoryEntryPacker interface {
	objectivec.IObject

	// Topic: Packing entries

	// Provides a directory entry during enumeration.
	PackEntryWithNameItemTypeItemIDNextCookieAttributes(name IFSFileName, itemType FSItemType, itemID FSItemID, nextCookie FSDirectoryCookie, attributes IFSItemAttributes) bool
}

// Init initializes the instance.
func (d FSDirectoryEntryPacker) Init() FSDirectoryEntryPacker {
	rv := objc.Send[FSDirectoryEntryPacker](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d FSDirectoryEntryPacker) Autorelease() FSDirectoryEntryPacker {
	rv := objc.Send[FSDirectoryEntryPacker](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewFSDirectoryEntryPacker creates a new FSDirectoryEntryPacker instance.
func NewFSDirectoryEntryPacker() FSDirectoryEntryPacker {
	class := getFSDirectoryEntryPackerClass()
	rv := objc.Send[FSDirectoryEntryPacker](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Provides a directory entry during enumeration.
//
// name: The item’s name.
//
// itemType: The type of the item.
//
// itemID: The item’s identifier. Typically this is an inode number, or one of the
// constants defined by [FSItem.Identifier] like
// [FSItem.Identifier.rootDirectory].
//
// nextCookie: A value to indicate the next entry in the directory to enumerate. FSKit
// passes this value as the `cookie` parameter on the next call to
// [EnumerateDirectoryStartingAtCookieVerifierProvidingAttributesUsingPackerReplyHandler].
// Use whatever value is appropriate for your implementation; the value is
// opaque to FSKit.
//
// attributes: The item’s attributes. Pass `nil` if the enumeration call didn’t
// request attributes.
//
// # Return Value
//
// `true` (Swift) or [YES] (Objective-C) if packing was successful and
// enumeration can continue with the next directory entry. If the value is
// `false` (Swift) or [NO] (Objective-C), stop enumerating. This result can
// happen when the entry is too big for the remaining space in the buffer.
//
// # Discussion
//
// You call this method in your implementation of
// [EnumerateDirectoryStartingAtCookieVerifierProvidingAttributesUsingPackerReplyHandler],
// for each directory entry you want to provide to the enumeration.
//
// See: https://developer.apple.com/documentation/FSKit/FSDirectoryEntryPacker/packEntry(name:itemType:itemID:nextCookie:attributes:)
//
// [FSItem.Identifier.rootDirectory]: https://developer.apple.com/documentation/FSKit/FSItem/Identifier/rootDirectory
// [FSItem.Identifier]: https://developer.apple.com/documentation/FSKit/FSItem/Identifier
func (d FSDirectoryEntryPacker) PackEntryWithNameItemTypeItemIDNextCookieAttributes(name IFSFileName, itemType FSItemType, itemID FSItemID, nextCookie FSDirectoryCookie, attributes IFSItemAttributes) bool {
	rv := objc.Send[bool](d.ID, objc.Sel("packEntryWithName:itemType:itemID:nextCookie:attributes:"), name, itemType, itemID, nextCookie, attributes)
	return rv
}
