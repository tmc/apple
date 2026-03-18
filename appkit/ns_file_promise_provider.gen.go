// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSFilePromiseProvider] class.
var (
	_NSFilePromiseProviderClass     NSFilePromiseProviderClass
	_NSFilePromiseProviderClassOnce sync.Once
)

func getNSFilePromiseProviderClass() NSFilePromiseProviderClass {
	_NSFilePromiseProviderClassOnce.Do(func() {
		_NSFilePromiseProviderClass = NSFilePromiseProviderClass{class: objc.GetClass("NSFilePromiseProvider")}
	})
	return _NSFilePromiseProviderClass
}

// GetNSFilePromiseProviderClass returns the class object for NSFilePromiseProvider.
func GetNSFilePromiseProviderClass() NSFilePromiseProviderClass {
	return getNSFilePromiseProviderClass()
}

type NSFilePromiseProviderClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSFilePromiseProviderClass) Alloc() NSFilePromiseProvider {
	rv := objc.Send[NSFilePromiseProvider](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// An object that provides a promise for the pasteboard.
//
// # Overview
// 
// A file promise is a possible future file of a specified type. When you’re
// working with drag and drop, use promises to indicate intent for future
// action. Avoid loading or performing any actions on the file until the
// promise completes.
// 
// Use the [NSFilePromiseProvider] class when creating file promises.
// Instantiate one [NSFilePromiseProvider] for each file promised. Set the
// [NSFilePromiseProvider.FileType] and [NSFilePromiseProvider.Delegate] properties before writing any
// [NSFilePromiseProvider] to the pasteboard. The file type must be a Uniform
// Type Identifier (UTI) that ultimately conforms to `kUTTypeData` or
// `kUTTypeDirectory`. The [NSFilePromiseProviderDelegate] will write the
// promised file to the destination directory.
// 
// Optionally, you may attach a `userInfo` object of your choosing to the
// [NSFilePromiseProvider] to determine which promise is being referenced when
// promising multiple files under the same [NSFilePromiseProviderDelegate]
// instance.
//
// # Initializers
//
//   - [NSFilePromiseProvider.InitWithFileTypeDelegate]: Initializes a file promise provider for a certain file type.
//
// # Instance Properties
//
//   - [NSFilePromiseProvider.Delegate]
//   - [NSFilePromiseProvider.SetDelegate]
//   - [NSFilePromiseProvider.FileType]: The file type of the file promise provider.
//   - [NSFilePromiseProvider.SetFileType]
//   - [NSFilePromiseProvider.UserInfo]: Optional user information to pass to the file promise provider.
//   - [NSFilePromiseProvider.SetUserInfo]
//
// See: https://developer.apple.com/documentation/AppKit/NSFilePromiseProvider
type NSFilePromiseProvider struct {
	objectivec.Object
}

// NSFilePromiseProviderFromID constructs a [NSFilePromiseProvider] from an objc.ID.
//
// An object that provides a promise for the pasteboard.
func NSFilePromiseProviderFromID(id objc.ID) NSFilePromiseProvider {
	return NSFilePromiseProvider{objectivec.Object{ID: id}}
}
// NOTE: NSFilePromiseProvider adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSFilePromiseProvider] class.
//
// # Initializers
//
//   - [INSFilePromiseProvider.InitWithFileTypeDelegate]: Initializes a file promise provider for a certain file type.
//
// # Instance Properties
//
//   - [INSFilePromiseProvider.Delegate]
//   - [INSFilePromiseProvider.SetDelegate]
//   - [INSFilePromiseProvider.FileType]: The file type of the file promise provider.
//   - [INSFilePromiseProvider.SetFileType]
//   - [INSFilePromiseProvider.UserInfo]: Optional user information to pass to the file promise provider.
//   - [INSFilePromiseProvider.SetUserInfo]
//
// See: https://developer.apple.com/documentation/AppKit/NSFilePromiseProvider
type INSFilePromiseProvider interface {
	objectivec.IObject
	NSPasteboardWriting

	// Topic: Initializers

	// Initializes a file promise provider for a certain file type.
	InitWithFileTypeDelegate(fileType string, delegate NSFilePromiseProviderDelegate) NSFilePromiseProvider

	// Topic: Instance Properties

	Delegate() NSFilePromiseProviderDelegate
	SetDelegate(value NSFilePromiseProviderDelegate)
	// The file type of the file promise provider.
	FileType() string
	SetFileType(value string)
	// Optional user information to pass to the file promise provider.
	UserInfo() objectivec.IObject
	SetUserInfo(value objectivec.IObject)
}

// Init initializes the instance.
func (f NSFilePromiseProvider) Init() NSFilePromiseProvider {
	rv := objc.Send[NSFilePromiseProvider](f.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f NSFilePromiseProvider) Autorelease() NSFilePromiseProvider {
	rv := objc.Send[NSFilePromiseProvider](f.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSFilePromiseProvider creates a new NSFilePromiseProvider instance.
func NewNSFilePromiseProvider() NSFilePromiseProvider {
	class := getNSFilePromiseProviderClass()
	rv := objc.Send[NSFilePromiseProvider](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Initializes a file promise provider for a certain file type.
//
// fileType: A string describing the file type.
//
// delegate: An object that conforms to the NSFilePromiseProviderDelegate protocol, for
// providing promised file data.
//
// See: https://developer.apple.com/documentation/AppKit/NSFilePromiseProvider/init(fileType:delegate:)
func NewFilePromiseProviderWithFileTypeDelegate(fileType string, delegate NSFilePromiseProviderDelegate) NSFilePromiseProvider {
	instance := getNSFilePromiseProviderClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFileType:delegate:"), objc.String(fileType), delegate)
	return NSFilePromiseProviderFromID(rv)
}

// Initializes a file promise provider for a certain file type.
//
// fileType: A string describing the file type.
//
// delegate: An object that conforms to the NSFilePromiseProviderDelegate protocol, for
// providing promised file data.
//
// See: https://developer.apple.com/documentation/AppKit/NSFilePromiseProvider/init(fileType:delegate:)
func (f NSFilePromiseProvider) InitWithFileTypeDelegate(fileType string, delegate NSFilePromiseProviderDelegate) NSFilePromiseProvider {
	rv := objc.Send[NSFilePromiseProvider](f.ID, objc.Sel("initWithFileType:delegate:"), objc.String(fileType), delegate)
	return rv
}

// Returns a property list object to represent the receiver on a pasteboard as
// an object of a specified type.
//
// type: One of the types the receiver supports for writing (one of the UTIs
// returned by its implementation of [WritableTypesForPasteboard]).
//
// # Return Value
// 
// A property list object to represent the receiver on a pasteboard as an
// object of type `type`.
//
// # Discussion
// 
// The returned value will commonly be the [NSData] object for the specified
// data type. However, if this method returns either a string, or any other
// property-list type, the pasteboard will automatically convert these items
// to the correct data format required for the pasteboard.
//
// See: https://developer.apple.com/documentation/AppKit/NSPasteboardWriting/pasteboardPropertyList(forType:)
func (f NSFilePromiseProvider) PasteboardPropertyListForType(type_ NSPasteboardType) objectivec.IObject {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("pasteboardPropertyListForType:"), objc.String(string(type_)))
	return objectivec.Object{ID: rv}
}

// Returns an array of UTI strings of data types the receiver can write to a
// given pasteboard.
//
// pasteboard: A pasteboard.
// 
// You can use this argument to provide different options based on the
// pasteboard name, if you need to.
//
// # Return Value
// 
// An array of UTI strings of data types the receiver can write to
// `pasteboard`.
//
// # Discussion
// 
// By default, data for the first returned type is put onto the pasteboard
// immediately, with the remaining types being promised.
// 
// To change the default behavior, implement
// -writingOptionsForType:pasteboard: and return [PasteboardWritingPromised]
// to lazily provide data for types, return no option to provide the data for
// that type immediately. Use the pasteboard argument to provide different
// types based on the pasteboard name, if desired. Do not perform other
// pasteboard operations in the method implementation.
//
// See: https://developer.apple.com/documentation/AppKit/NSPasteboardWriting/writableTypes(for:)
func (f NSFilePromiseProvider) WritableTypesForPasteboard(pasteboard INSPasteboard) []string {
	rv := objc.Send[[]objc.ID](f.ID, objc.Sel("writableTypesForPasteboard:"), pasteboard)
	return objc.ConvertSliceToStrings(rv)
}

// Returns options for writing data of a specified type to a given pasteboard.
//
// type: One of the types the receiver supports for writing (one of the UTIs
// returned by its implementation of [WritableTypesForPasteboard]).
//
// pasteboard: A pasteboard.
// 
// You can use this argument to provide different options based on the
// pasteboard name, if you need to.
//
// # Return Value
// 
// Options for writing data of type type to `pasteboard`. Return `0` for no
// options, or a value given in [Pasteboard Writing Options].
//
// [Pasteboard Writing Options]: https://developer.apple.com/documentation/AppKit/pasteboard-writing-options
//
// # Discussion
// 
// Do not perform other pasteboard operations in the method implementation.
//
// See: https://developer.apple.com/documentation/AppKit/NSPasteboardWriting/writingOptions(forType:pasteboard:)
func (f NSFilePromiseProvider) WritingOptionsForTypePasteboard(type_ NSPasteboardType, pasteboard INSPasteboard) NSPasteboardWritingOptions {
	rv := objc.Send[NSPasteboardWritingOptions](f.ID, objc.Sel("writingOptionsForType:pasteboard:"), objc.String(string(type_)), pasteboard)
	return NSPasteboardWritingOptions(rv)
}

// See: https://developer.apple.com/documentation/AppKit/NSFilePromiseProvider/delegate
func (f NSFilePromiseProvider) Delegate() NSFilePromiseProviderDelegate {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("delegate"))
	return NSFilePromiseProviderDelegateObjectFromID(rv)
}
func (f NSFilePromiseProvider) SetDelegate(value NSFilePromiseProviderDelegate) {
	objc.Send[struct{}](f.ID, objc.Sel("setDelegate:"), value)
}

// The file type of the file promise provider.
//
// # Discussion
// 
// The type is Universal Type Identifier (UTI).
//
// See: https://developer.apple.com/documentation/AppKit/NSFilePromiseProvider/fileType
func (f NSFilePromiseProvider) FileType() string {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("fileType"))
	return foundation.NSStringFromID(rv).String()
}
func (f NSFilePromiseProvider) SetFileType(value string) {
	objc.Send[struct{}](f.ID, objc.Sel("setFileType:"), objc.String(value))
}

// Optional user information to pass to the file promise provider.
//
// See: https://developer.apple.com/documentation/AppKit/NSFilePromiseProvider/userInfo
func (f NSFilePromiseProvider) UserInfo() objectivec.IObject {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("userInfo"))
	return objectivec.Object{ID: rv}
}
func (f NSFilePromiseProvider) SetUserInfo(value objectivec.IObject) {
	objc.Send[struct{}](f.ID, objc.Sel("setUserInfo:"), value)
}

			// Protocol methods for NSPasteboardWriting
			

