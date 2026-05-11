// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"errors"
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSColorList] class.
var (
	_NSColorListClass     NSColorListClass
	_NSColorListClassOnce sync.Once
)

func getNSColorListClass() NSColorListClass {
	_NSColorListClassOnce.Do(func() {
		_NSColorListClass = NSColorListClass{class: objc.GetClass("NSColorList")}
	})
	return _NSColorListClass
}

// GetNSColorListClass returns the class object for NSColorList.
func GetNSColorListClass() NSColorListClass {
	return getNSColorListClass()
}

type NSColorListClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSColorListClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSColorListClass) Alloc() NSColorList {
	rv := objc.Send[NSColorList](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// An ordered list of color objects, identified by keys.
//
// # Overview
//
// A color list manages a list of [NSColor] objects, each of which has an
// associated name. The [NSColorPanel] list mode color picker uses instances
// of [NSColorList] to represent any lists of colors that come with the
// system, as well as any lists the user creates. An app can use a color list
// to manage document-specific color lists.
//
// # Creating Lists of Colors
//
//   - [NSColorList.InitWithName]: Initializes and returns a color list, registering it under the specified name if it isn’t in use already.
//   - [NSColorList.InitWithNameFromFile]: Initializes and returns a color list from the specified file, registering it under the specified name if it isn’t in use already.
//
// # Getting Information About Lists of Colors
//
//   - [NSColorList.Name]: The name of the color list.
//   - [NSColorList.Editable]: A Boolean value that indicates whether the color list can be modified.
//
// # Managing Colors By Key
//
//   - [NSColorList.AllKeys]: An array of the keys by which the color objects are stored in the color list.
//   - [NSColorList.ColorWithKey]: Returns the color object associated with the specified key.
//   - [NSColorList.InsertColorKeyAtIndex]: Inserts the specified color at the specified location in the color list.
//   - [NSColorList.RemoveColorWithKey]: Removes the color associated with the specified key from the color list.
//   - [NSColorList.SetColorForKey]: Associates the specified color object with the specified key.
//
// # Writing and Removing Color List Files
//
//   - [NSColorList.WriteToURLError]: Saves the color list to the file at the specified URL.
//   - [NSColorList.RemoveFile]: Removes the file from which the list was created, if the file is in a standard search path and owned by the user.
//
// See: https://developer.apple.com/documentation/AppKit/NSColorList
type NSColorList struct {
	objectivec.Object
}

// NSColorListFromID constructs a [NSColorList] from an objc.ID.
//
// An ordered list of color objects, identified by keys.
func NSColorListFromID(id objc.ID) NSColorList {
	return NSColorList{objectivec.Object{ID: id}}
}

// NOTE: NSColorList adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSColorList] class.
//
// # Creating Lists of Colors
//
//   - [INSColorList.InitWithName]: Initializes and returns a color list, registering it under the specified name if it isn’t in use already.
//   - [INSColorList.InitWithNameFromFile]: Initializes and returns a color list from the specified file, registering it under the specified name if it isn’t in use already.
//
// # Getting Information About Lists of Colors
//
//   - [INSColorList.Name]: The name of the color list.
//   - [INSColorList.Editable]: A Boolean value that indicates whether the color list can be modified.
//
// # Managing Colors By Key
//
//   - [INSColorList.AllKeys]: An array of the keys by which the color objects are stored in the color list.
//   - [INSColorList.ColorWithKey]: Returns the color object associated with the specified key.
//   - [INSColorList.InsertColorKeyAtIndex]: Inserts the specified color at the specified location in the color list.
//   - [INSColorList.RemoveColorWithKey]: Removes the color associated with the specified key from the color list.
//   - [INSColorList.SetColorForKey]: Associates the specified color object with the specified key.
//
// # Writing and Removing Color List Files
//
//   - [INSColorList.WriteToURLError]: Saves the color list to the file at the specified URL.
//   - [INSColorList.RemoveFile]: Removes the file from which the list was created, if the file is in a standard search path and owned by the user.
//
// See: https://developer.apple.com/documentation/AppKit/NSColorList
type INSColorList interface {
	objectivec.IObject

	// Topic: Creating Lists of Colors

	// Initializes and returns a color list, registering it under the specified name if it isn’t in use already.
	InitWithName(name string) NSColorList
	// Initializes and returns a color list from the specified file, registering it under the specified name if it isn’t in use already.
	InitWithNameFromFile(name string, path string) NSColorList

	// Topic: Getting Information About Lists of Colors

	// The name of the color list.
	Name() NSColorListName
	// A Boolean value that indicates whether the color list can be modified.
	Editable() bool

	// Topic: Managing Colors By Key

	// An array of the keys by which the color objects are stored in the color list.
	AllKeys() []string
	// Returns the color object associated with the specified key.
	ColorWithKey(key string) INSColor
	// Inserts the specified color at the specified location in the color list.
	InsertColorKeyAtIndex(color INSColor, key string, loc uint)
	// Removes the color associated with the specified key from the color list.
	RemoveColorWithKey(key string)
	// Associates the specified color object with the specified key.
	SetColorForKey(color INSColor, key string)

	// Topic: Writing and Removing Color List Files

	// Saves the color list to the file at the specified URL.
	WriteToURLError(url foundation.INSURL) (bool, error)
	// Removes the file from which the list was created, if the file is in a standard search path and owned by the user.
	RemoveFile()

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (c NSColorList) Init() NSColorList {
	rv := objc.Send[NSColorList](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c NSColorList) Autorelease() NSColorList {
	rv := objc.Send[NSColorList](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSColorList creates a new NSColorList instance.
func NewNSColorList() NSColorList {
	class := getNSColorListClass()
	rv := objc.Send[NSColorList](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Searches the available color lists array and returns the color list with
// the specified name.
//
// name: The name of the color list to retrieve. This name must not include the
// “`XCUIElementTypeClr`” suffix.
//
// # Return Value
//
// The color list with the specified name or `nil` if no such color list
// exists.
//
// See: https://developer.apple.com/documentation/AppKit/NSColorList/init(named:)
func NewColorListNamed(name string) NSColorList {
	rv := objc.Send[objc.ID](objc.ID(getNSColorListClass().class), objc.Sel("colorListNamed:"), objc.String(name))
	return NSColorListFromID(rv)
}

// Initializes and returns a color list, registering it under the specified
// name if it isn’t in use already.
//
// name: The name under which to register the color list. Specify `@””` if you
// don’t want a name.
//
// # Return Value
//
// The initialized color list.
//
// # Discussion
//
// This method invokes [InitWithNameFromFile] with a “ argument of `nil`,
// indicating that the color list doesn’t need to be initialized from a
// file. Note that this method does not add the color list to
// [AvailableColorLists] until the color list is saved into the user’s path
// with [write(toFile:)] with a value of `nil`.
//
// See: https://developer.apple.com/documentation/AppKit/NSColorList/init(name:)
//
// [write(toFile:)]: https://developer.apple.com/documentation/AppKit/NSColorList/write(toFile:)
func NewColorListWithName(name string) NSColorList {
	instance := getNSColorListClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithName:"), objc.String(name))
	return NSColorListFromID(rv)
}

// Initializes and returns a color list from the specified file, registering
// it under the specified name if it isn’t in use already.
//
// name: The name of the file for the color list (minus the
// `“XCUIElementTypeClr”` extension). Specify `@””` if you don’t
// want a name.
//
// path: The full path to the file for the color list. A `nil` path indicates the
// color list should be initialized with no colors.
//
// # Discussion
//
// Note that this method does not add the color list to [AvailableColorLists]
// until the color list is saved into the user’s path with [write(toFile:)]
// with a value of `nil`.
//
// See: https://developer.apple.com/documentation/AppKit/NSColorList/init(name:fromFile:)
//
// [write(toFile:)]: https://developer.apple.com/documentation/AppKit/NSColorList/write(toFile:)
func NewColorListWithNameFromFile(name string, path string) NSColorList {
	instance := getNSColorListClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithName:fromFile:"), objc.String(name), objc.String(path))
	return NSColorListFromID(rv)
}

// Initializes and returns a color list, registering it under the specified
// name if it isn’t in use already.
//
// name: The name under which to register the color list. Specify `@””` if you
// don’t want a name.
//
// # Return Value
//
// The initialized color list.
//
// # Discussion
//
// This method invokes [InitWithNameFromFile] with a “ argument of `nil`,
// indicating that the color list doesn’t need to be initialized from a
// file. Note that this method does not add the color list to
// [AvailableColorLists] until the color list is saved into the user’s path
// with [write(toFile:)] with a value of `nil`.
//
// See: https://developer.apple.com/documentation/AppKit/NSColorList/init(name:)
//
// [write(toFile:)]: https://developer.apple.com/documentation/AppKit/NSColorList/write(toFile:)
func (c NSColorList) InitWithName(name string) NSColorList {
	rv := objc.Send[NSColorList](c.ID, objc.Sel("initWithName:"), objc.String(name))
	return rv
}

// Initializes and returns a color list from the specified file, registering
// it under the specified name if it isn’t in use already.
//
// name: The name of the file for the color list (minus the
// `“XCUIElementTypeClr”` extension). Specify `@””` if you don’t
// want a name.
//
// path: The full path to the file for the color list. A `nil` path indicates the
// color list should be initialized with no colors.
//
// # Discussion
//
// Note that this method does not add the color list to [AvailableColorLists]
// until the color list is saved into the user’s path with [write(toFile:)]
// with a value of `nil`.
//
// See: https://developer.apple.com/documentation/AppKit/NSColorList/init(name:fromFile:)
//
// [write(toFile:)]: https://developer.apple.com/documentation/AppKit/NSColorList/write(toFile:)
func (c NSColorList) InitWithNameFromFile(name string, path string) NSColorList {
	rv := objc.Send[NSColorList](c.ID, objc.Sel("initWithName:fromFile:"), objc.String(name), objc.String(path))
	return rv
}

// Returns the color object associated with the specified key.
//
// key: The key for which to retrieve the color.
//
// # Return Value
//
// The color associated with the given key or `nil` if there is none.
//
// See: https://developer.apple.com/documentation/AppKit/NSColorList/color(withKey:)
func (c NSColorList) ColorWithKey(key string) INSColor {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("colorWithKey:"), objc.String(key))
	return NSColorFromID(rv)
}

// Inserts the specified color at the specified location in the color list.
//
// color: The color to add to the color list.
//
// key: The key with which to associate the color.
//
// loc: The location in the color list at which to place the specified color.
// Locations are numbered starting with 0.
//
// # Discussion
//
// If the list already contains a color with the same key at a different
// location, it’s removed from the old location. This method posts
// [didChangeNotification] to the default notification center. It raises
// [NSColorListNotEditableException] if the color list isn’t editable.
//
// See: https://developer.apple.com/documentation/AppKit/NSColorList/insertColor(_:key:at:)
//
// [didChangeNotification]: https://developer.apple.com/documentation/AppKit/NSColorList/didChangeNotification
func (c NSColorList) InsertColorKeyAtIndex(color INSColor, key string, loc uint) {
	objc.Send[objc.ID](c.ID, objc.Sel("insertColor:key:atIndex:"), color, objc.String(key), loc)
}

// Removes the color associated with the specified key from the color list.
//
// key: The key for which to remove the color.
//
// # Discussion
//
// This method does nothing if the receiver doesn’t contain the key. This
// method posts [didChangeNotification] to the default notification center. It
// raises [NSColorListNotEditableException] if the receiver is not editable.
//
// See: https://developer.apple.com/documentation/AppKit/NSColorList/removeColor(withKey:)
//
// [didChangeNotification]: https://developer.apple.com/documentation/AppKit/NSColorList/didChangeNotification
func (c NSColorList) RemoveColorWithKey(key string) {
	objc.Send[objc.ID](c.ID, objc.Sel("removeColorWithKey:"), objc.String(key))
}

// Associates the specified color object with the specified key.
//
// color: The color to associate with the given key.
//
// key: The key.
//
// # Discussion
//
// If the list already contains `key`, this method sets the corresponding
// color to `color`; otherwise, it inserts `color` at the end of the list by
// invoking [InsertColorKeyAtIndex].
//
// See: https://developer.apple.com/documentation/AppKit/NSColorList/setColor(_:forKey:)
func (c NSColorList) SetColorForKey(color INSColor, key string) {
	objc.Send[objc.ID](c.ID, objc.Sel("setColor:forKey:"), color, objc.String(key))
}

// Saves the color list to the file at the specified URL.
//
// url: The URL at which to store the color list. The URL must specify either a
// directory or file in the file system. Specify `nil` to save the color list
// to the user’s `~/Library/Colors` directory.
//
// # Discussion
//
// If `url` represents a directory, this method saves the color list in that
// directory in a file with the name `XCUIElementTypeClr`, where is the value
// of the [Name] property. If `url` represents a file, this method saves the
// color list using the name you provided.
//
// See: https://developer.apple.com/documentation/AppKit/NSColorList/write(to:)
func (c NSColorList) WriteToURLError(url foundation.INSURL) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](c.ID, objc.Sel("writeToURL:error:"), url, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("writeToURL:error: returned NO with nil NSError")
	}
	return rv, nil

}

// Removes the file from which the list was created, if the file is in a
// standard search path and owned by the user.
//
// # Discussion
//
// In addition to removing the file, this method removes the color list from
// the contents of the [AvailableColorLists] property. If there are no
// outstanding references to the color list, this method might also deallocate
// the object.
//
// See: https://developer.apple.com/documentation/AppKit/NSColorList/removeFile()
func (c NSColorList) RemoveFile() {
	objc.Send[objc.ID](c.ID, objc.Sel("removeFile"))
}
func (c NSColorList) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](c.ID, objc.Sel("encodeWithCoder:"), coder)
}

// The name of the color list.
//
// See: https://developer.apple.com/documentation/AppKit/NSColorList/name-swift.property
func (c NSColorList) Name() NSColorListName {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("name"))
	return NSColorListName(foundation.NSStringFromID(rv).String())
}

// A Boolean value that indicates whether the color list can be modified.
//
// See: https://developer.apple.com/documentation/AppKit/NSColorList/isEditable
func (c NSColorList) Editable() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isEditable"))
	return rv
}

// An array of the keys by which the color objects are stored in the color
// list.
//
// See: https://developer.apple.com/documentation/AppKit/NSColorList/allKeys
func (c NSColorList) AllKeys() []string {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("allKeys"))
	return objc.ConvertSliceToStrings(rv)
}

// Returns an array of all color lists found in the standard color list
// directories.
//
// # Return Value
//
// An array of [NSColorList] objects representing all of the color lists found
// in the standard color list directories, including color catalogs (lists of
// colors identified only by name). Color lists created at runtime aren’t
// included in this list unless they’re saved into one of the standard color
// list directories.
//
// See: https://developer.apple.com/documentation/AppKit/NSColorList/availableColorLists
func (_NSColorListClass NSColorListClass) AvailableColorLists() []NSColorList {
	rv := objc.Send[[]objc.ID](objc.ID(_NSColorListClass.class), objc.Sel("availableColorLists"))
	return objc.ConvertSlice(rv, func(id objc.ID) NSColorList {
		return NSColorListFromID(id)
	})
}
