// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [NSDirectoryEnumerator] class.
var (
	_NSDirectoryEnumeratorClass     NSDirectoryEnumeratorClass
	_NSDirectoryEnumeratorClassOnce sync.Once
)

func getNSDirectoryEnumeratorClass() NSDirectoryEnumeratorClass {
	_NSDirectoryEnumeratorClassOnce.Do(func() {
		_NSDirectoryEnumeratorClass = NSDirectoryEnumeratorClass{class: objc.GetClass("NSDirectoryEnumerator")}
	})
	return _NSDirectoryEnumeratorClass
}

// GetNSDirectoryEnumeratorClass returns the class object for NSDirectoryEnumerator.
func GetNSDirectoryEnumeratorClass() NSDirectoryEnumeratorClass {
	return getNSDirectoryEnumeratorClass()
}

type NSDirectoryEnumeratorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSDirectoryEnumeratorClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSDirectoryEnumeratorClass) Alloc() NSDirectoryEnumerator {
	rv := objc.Send[NSDirectoryEnumerator](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// An object that enumerates the contents of a directory.
//
// # Overview
// 
// You obtain a directory enumerator using [NSFileManager]’s
// [EnumeratorAtPath] method. The enumeration provides the pathnames of all
// files and directories contained within that directory. These pathnames are
// relative to the directory.
// 
// An enumeration is recursive, including the files of all subdirectories, and
// crosses device boundaries. An enumeration does not resolve symbolic links,
// or attempt to traverse symbolic links that point to directories.
//
// # Getting File and Directory Attributes
//
//   - [NSDirectoryEnumerator.DirectoryAttributes]: A dictionary with the attributes of the directory at which enumeration started.
//   - [NSDirectoryEnumerator.FileAttributes]: A dictionary with the attributes of the most recently returned file or subdirectory (as referenced by the pathname).
//   - [NSDirectoryEnumerator.Level]: The number of levels deep the current object is in the directory hierarchy being enumerated.
//
// # Skipping Subdirectories
//
//   - [NSDirectoryEnumerator.SkipDescendents]: Causes the receiver to skip recursion into the most recently obtained subdirectory.
//   - [NSDirectoryEnumerator.SkipDescendants]: Causes the receiver to skip recursion into the most recently obtained subdirectory.
//
// # Instance Properties
//
//   - [NSDirectoryEnumerator.IsEnumeratingDirectoryPostOrder]
//
// See: https://developer.apple.com/documentation/Foundation/FileManager/DirectoryEnumerator
type NSDirectoryEnumerator struct {
	NSEnumerator
}

// NSDirectoryEnumeratorFromID constructs a [NSDirectoryEnumerator] from an objc.ID.
//
// An object that enumerates the contents of a directory.
func NSDirectoryEnumeratorFromID(id objc.ID) NSDirectoryEnumerator {
	return NSDirectoryEnumerator{NSEnumerator: NSEnumeratorFromID(id)}
}
// NOTE: NSDirectoryEnumerator adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSDirectoryEnumerator] class.
//
// # Getting File and Directory Attributes
//
//   - [INSDirectoryEnumerator.DirectoryAttributes]: A dictionary with the attributes of the directory at which enumeration started.
//   - [INSDirectoryEnumerator.FileAttributes]: A dictionary with the attributes of the most recently returned file or subdirectory (as referenced by the pathname).
//   - [INSDirectoryEnumerator.Level]: The number of levels deep the current object is in the directory hierarchy being enumerated.
//
// # Skipping Subdirectories
//
//   - [INSDirectoryEnumerator.SkipDescendents]: Causes the receiver to skip recursion into the most recently obtained subdirectory.
//   - [INSDirectoryEnumerator.SkipDescendants]: Causes the receiver to skip recursion into the most recently obtained subdirectory.
//
// # Instance Properties
//
//   - [INSDirectoryEnumerator.IsEnumeratingDirectoryPostOrder]
//
// See: https://developer.apple.com/documentation/Foundation/FileManager/DirectoryEnumerator
type INSDirectoryEnumerator interface {
	INSEnumerator

	// Topic: Getting File and Directory Attributes

	// A dictionary with the attributes of the directory at which enumeration started.
	DirectoryAttributes() INSDictionary
	// A dictionary with the attributes of the most recently returned file or subdirectory (as referenced by the pathname).
	FileAttributes() INSDictionary
	// The number of levels deep the current object is in the directory hierarchy being enumerated.
	Level() uint

	// Topic: Skipping Subdirectories

	// Causes the receiver to skip recursion into the most recently obtained subdirectory.
	SkipDescendents()
	// Causes the receiver to skip recursion into the most recently obtained subdirectory.
	SkipDescendants()

	// Topic: Instance Properties

	IsEnumeratingDirectoryPostOrder() bool
}

// Init initializes the instance.
func (d NSDirectoryEnumerator) Init() NSDirectoryEnumerator {
	rv := objc.Send[NSDirectoryEnumerator](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d NSDirectoryEnumerator) Autorelease() NSDirectoryEnumerator {
	rv := objc.Send[NSDirectoryEnumerator](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSDirectoryEnumerator creates a new NSDirectoryEnumerator instance.
func NewNSDirectoryEnumerator() NSDirectoryEnumerator {
	class := getNSDirectoryEnumeratorClass()
	rv := objc.Send[NSDirectoryEnumerator](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Causes the receiver to skip recursion into the most recently obtained
// subdirectory.
//
// See: https://developer.apple.com/documentation/Foundation/FileManager/DirectoryEnumerator/skipDescendents()
func (d NSDirectoryEnumerator) SkipDescendents() {
	objc.Send[objc.ID](d.ID, objc.Sel("skipDescendents"))
}
// Causes the receiver to skip recursion into the most recently obtained
// subdirectory.
//
// # Discussion
// 
// This method is identical to [SkipDescendents] except for the spelling.
//
// See: https://developer.apple.com/documentation/Foundation/FileManager/DirectoryEnumerator/skipDescendants()
func (d NSDirectoryEnumerator) SkipDescendants() {
	objc.Send[objc.ID](d.ID, objc.Sel("skipDescendants"))
}

// A dictionary with the attributes of the directory at which enumeration
// started.
//
// # Discussion
// 
// See the description of the [FileAttributesAtPathTraverseLink] method of
// [NSFileManager] for details on obtaining the attributes from the
// dictionary.
//
// See: https://developer.apple.com/documentation/Foundation/FileManager/DirectoryEnumerator/directoryAttributes
func (d NSDirectoryEnumerator) DirectoryAttributes() INSDictionary {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("directoryAttributes"))
	return NSDictionaryFromID(objc.ID(rv))
}
// A dictionary with the attributes of the most recently returned file or
// subdirectory (as referenced by the pathname).
//
// # Discussion
// 
// See the description of the [FileAttributesAtPathTraverseLink] method of
// [NSFileManager] for details on obtaining the attributes from the
// dictionary.
//
// See: https://developer.apple.com/documentation/Foundation/FileManager/DirectoryEnumerator/fileAttributes
func (d NSDirectoryEnumerator) FileAttributes() INSDictionary {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("fileAttributes"))
	return NSDictionaryFromID(objc.ID(rv))
}
// The number of levels deep the current object is in the directory hierarchy
// being enumerated.
//
// # Discussion
// 
// The number of levels, with the directory passed to
// [EnumeratorAtURLIncludingPropertiesForKeysOptionsErrorHandler]
// ([NSFileManager]) considered to be level `0`.
//
// See: https://developer.apple.com/documentation/Foundation/FileManager/DirectoryEnumerator/level
func (d NSDirectoryEnumerator) Level() uint {
	rv := objc.Send[uint](d.ID, objc.Sel("level"))
	return rv
}
// See: https://developer.apple.com/documentation/Foundation/FileManager/DirectoryEnumerator/isEnumeratingDirectoryPostOrder
func (d NSDirectoryEnumerator) IsEnumeratingDirectoryPostOrder() bool {
	rv := objc.Send[bool](d.ID, objc.Sel("isEnumeratingDirectoryPostOrder"))
	return rv
}

