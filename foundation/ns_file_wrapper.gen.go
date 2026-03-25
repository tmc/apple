// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"errors"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [FileWrapper] class.
var (
	_FileWrapperClass     FileWrapperClass
	_FileWrapperClassOnce sync.Once
)

func getFileWrapperClass() FileWrapperClass {
	_FileWrapperClassOnce.Do(func() {
		_FileWrapperClass = FileWrapperClass{class: objc.GetClass("NSFileWrapper")}
	})
	return _FileWrapperClass
}

// GetFileWrapperClass returns the class object for NSFileWrapper.
func GetFileWrapperClass() FileWrapperClass {
	return getFileWrapperClass()
}

type FileWrapperClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (fc FileWrapperClass) Alloc() FileWrapper {
	rv := objc.Send[FileWrapper](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// A representation of a node (a file, directory, or symbolic link) in the
// file system.
//
// # Overview
// 
// The [NSFileWrapper] class provides access to the attributes and contents of
// file system nodes. A file system node is a file, directory, or symbolic
// link. Instances of this class are known as file wrappers.
// 
// File wrappers represent a file system node as an object that can be
// displayed as an image (and possibly edited in place), saved to the file
// system, or transmitted to another application.
// 
// There are three types of file wrappers:
// 
// - Regular-file file wrapper: Represents a regular file. - Directory file
// wrapper: Represents a directory. - Symbolic-link file wrapper: Represents a
// symbolic link.
// 
// A file wrapper has these attributes:
// 
// - Filename. Name of the file system node the file wrapper represents. -
// file-system attributes. See [NSFileManager] for information on the contents
// of the `attributes` dictionary. - Regular-file contents. Applicable only to
// regular-file file wrappers. - File wrappers. Applicable only to directory
// file wrappers. - Destination node. Applicable only to symbolic-link file
// wrappers.
//
// # Creating File Wrappers
//
//   - [FileWrapper.InitWithURLOptionsError]: Initializes a file wrapper instance whose kind is determined by the type of file-system node located by the URL.
//   - [FileWrapper.InitDirectoryWithFileWrappers]: Initializes the receiver as a directory file wrapper, with a given file-wrapper list.
//   - [FileWrapper.InitRegularFileWithContents]: Initializes the receiver as a regular-file file wrapper.
//   - [FileWrapper.InitSymbolicLinkWithDestinationURL]: Initializes the receiver as a symbolic-link file wrapper that links to a specified file.
//   - [FileWrapper.InitWithSerializedRepresentation]: Initializes the receiver as a regular-file file wrapper from given serialized data.
//
// # Querying File Wrappers
//
//   - [FileWrapper.RegularFile]: This property contains a boolean value that indicates whether the file wrapper object is a regular-file.
//   - [FileWrapper.Directory]: This property contains a boolean value indicating whether the file wrapper is a directory file wrapper.
//   - [FileWrapper.SymbolicLink]: A boolean that indicates whether the file wrapper object is a symbolic-link file wrapper.
//
// # Accessing File-Wrapper Information
//
//   - [FileWrapper.FileWrappers]: The file wrappers contained by a directory file wrapper.
//   - [FileWrapper.AddFileWrapper]: Adds a child file wrapper to the receiver, which must be a directory file wrapper.
//   - [FileWrapper.RemoveFileWrapper]: Removes a child file wrapper from the receiver, which must be a directory file wrapper.
//   - [FileWrapper.AddRegularFileWithContentsPreferredFilename]: Creates a regular-file file wrapper with the given contents and adds it to the receiver, which must be a directory file wrapper.
//   - [FileWrapper.KeyForFileWrapper]: Returns the dictionary key used by a directory to identify a given file wrapper.
//   - [FileWrapper.SymbolicLinkDestinationURL]: The URL referenced by the file wrapper object, which must be a symbolic-link file wrapper.
//
// # Updating File Wrappers
//
//   - [FileWrapper.MatchesContentsOfURL]: Indicates whether the contents of a file wrapper matches a directory, regular file, or symbolic link on disk.
//   - [FileWrapper.ReadFromURLOptionsError]: Recursively rereads the entire contents of a file wrapper from the specified location on disk.
//
// # Serializing
//
//   - [FileWrapper.SerializedRepresentation]: The contents of the file wrapper as an opaque data object.
//
// # Accessing Files
//
//   - [FileWrapper.Filename]: The filename of the file wrapper object
//   - [FileWrapper.SetFilename]
//   - [FileWrapper.PreferredFilename]: The preferred filename for the file wrapper object.
//   - [FileWrapper.SetPreferredFilename]
//   - [FileWrapper.FileAttributes]: A dictionary of file attributes.
//   - [FileWrapper.SetFileAttributes]
//   - [FileWrapper.RegularFileContents]: The contents of the file-system node associated with a regular-file file wrapper.
//
// # Writing Files
//
//   - [FileWrapper.WriteToURLOptionsOriginalContentsURLError]: Recursively writes the entire contents of a file wrapper to a given file-system URL.
//
// # Working with Icons
//
//   - [FileWrapper.Icon]: The icon that represents the file wrapper.
//   - [FileWrapper.SetIcon]
//
// See: https://developer.apple.com/documentation/Foundation/FileWrapper
type FileWrapper struct {
	objectivec.Object
}

// FileWrapperFromID constructs a [FileWrapper] from an objc.ID.
//
// A representation of a node (a file, directory, or symbolic link) in the
// file system.
func FileWrapperFromID(id objc.ID) FileWrapper {
	return NSFileWrapper{objectivec.Object{ID: id}}
}

// NSFileWrapperFromID is an alias for [FileWrapperFromID] for cross-framework compatibility.
func NSFileWrapperFromID(id objc.ID) FileWrapper { return FileWrapperFromID(id) }
// NOTE: FileWrapper adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [FileWrapper] class.
//
// # Creating File Wrappers
//
//   - [IFileWrapper.InitWithURLOptionsError]: Initializes a file wrapper instance whose kind is determined by the type of file-system node located by the URL.
//   - [IFileWrapper.InitDirectoryWithFileWrappers]: Initializes the receiver as a directory file wrapper, with a given file-wrapper list.
//   - [IFileWrapper.InitRegularFileWithContents]: Initializes the receiver as a regular-file file wrapper.
//   - [IFileWrapper.InitSymbolicLinkWithDestinationURL]: Initializes the receiver as a symbolic-link file wrapper that links to a specified file.
//   - [IFileWrapper.InitWithSerializedRepresentation]: Initializes the receiver as a regular-file file wrapper from given serialized data.
//
// # Querying File Wrappers
//
//   - [IFileWrapper.RegularFile]: This property contains a boolean value that indicates whether the file wrapper object is a regular-file.
//   - [IFileWrapper.Directory]: This property contains a boolean value indicating whether the file wrapper is a directory file wrapper.
//   - [IFileWrapper.SymbolicLink]: A boolean that indicates whether the file wrapper object is a symbolic-link file wrapper.
//
// # Accessing File-Wrapper Information
//
//   - [IFileWrapper.FileWrappers]: The file wrappers contained by a directory file wrapper.
//   - [IFileWrapper.AddFileWrapper]: Adds a child file wrapper to the receiver, which must be a directory file wrapper.
//   - [IFileWrapper.RemoveFileWrapper]: Removes a child file wrapper from the receiver, which must be a directory file wrapper.
//   - [IFileWrapper.AddRegularFileWithContentsPreferredFilename]: Creates a regular-file file wrapper with the given contents and adds it to the receiver, which must be a directory file wrapper.
//   - [IFileWrapper.KeyForFileWrapper]: Returns the dictionary key used by a directory to identify a given file wrapper.
//   - [IFileWrapper.SymbolicLinkDestinationURL]: The URL referenced by the file wrapper object, which must be a symbolic-link file wrapper.
//
// # Updating File Wrappers
//
//   - [IFileWrapper.MatchesContentsOfURL]: Indicates whether the contents of a file wrapper matches a directory, regular file, or symbolic link on disk.
//   - [IFileWrapper.ReadFromURLOptionsError]: Recursively rereads the entire contents of a file wrapper from the specified location on disk.
//
// # Serializing
//
//   - [IFileWrapper.SerializedRepresentation]: The contents of the file wrapper as an opaque data object.
//
// # Accessing Files
//
//   - [IFileWrapper.Filename]: The filename of the file wrapper object
//   - [IFileWrapper.SetFilename]
//   - [IFileWrapper.PreferredFilename]: The preferred filename for the file wrapper object.
//   - [IFileWrapper.SetPreferredFilename]
//   - [IFileWrapper.FileAttributes]: A dictionary of file attributes.
//   - [IFileWrapper.SetFileAttributes]
//   - [IFileWrapper.RegularFileContents]: The contents of the file-system node associated with a regular-file file wrapper.
//
// # Writing Files
//
//   - [IFileWrapper.WriteToURLOptionsOriginalContentsURLError]: Recursively writes the entire contents of a file wrapper to a given file-system URL.
//
// # Working with Icons
//
//   - [IFileWrapper.Icon]: The icon that represents the file wrapper.
//   - [IFileWrapper.SetIcon]
//
// See: https://developer.apple.com/documentation/Foundation/FileWrapper
type IFileWrapper interface {
	objectivec.IObject
	NSCoding
	NSSecureCoding

	// Topic: Creating File Wrappers

	// Initializes a file wrapper instance whose kind is determined by the type of file-system node located by the URL.
	InitWithURLOptionsError(url INSURL, options NSFileWrapperReadingOptions) (FileWrapper, error)
	// Initializes the receiver as a directory file wrapper, with a given file-wrapper list.
	InitDirectoryWithFileWrappers(childrenByPreferredName INSDictionary) FileWrapper
	// Initializes the receiver as a regular-file file wrapper.
	InitRegularFileWithContents(contents INSData) FileWrapper
	// Initializes the receiver as a symbolic-link file wrapper that links to a specified file.
	InitSymbolicLinkWithDestinationURL(url INSURL) FileWrapper
	// Initializes the receiver as a regular-file file wrapper from given serialized data.
	InitWithSerializedRepresentation(serializeRepresentation INSData) FileWrapper

	// Topic: Querying File Wrappers

	// This property contains a boolean value that indicates whether the file wrapper object is a regular-file.
	RegularFile() bool
	// This property contains a boolean value indicating whether the file wrapper is a directory file wrapper.
	Directory() bool
	// A boolean that indicates whether the file wrapper object is a symbolic-link file wrapper.
	SymbolicLink() bool

	// Topic: Accessing File-Wrapper Information

	// The file wrappers contained by a directory file wrapper.
	FileWrappers() INSDictionary
	// Adds a child file wrapper to the receiver, which must be a directory file wrapper.
	AddFileWrapper(child INSFileWrapper) string
	// Removes a child file wrapper from the receiver, which must be a directory file wrapper.
	RemoveFileWrapper(child INSFileWrapper)
	// Creates a regular-file file wrapper with the given contents and adds it to the receiver, which must be a directory file wrapper.
	AddRegularFileWithContentsPreferredFilename(data INSData, fileName string) string
	// Returns the dictionary key used by a directory to identify a given file wrapper.
	KeyForFileWrapper(child INSFileWrapper) string
	// The URL referenced by the file wrapper object, which must be a symbolic-link file wrapper.
	SymbolicLinkDestinationURL() INSURL

	// Topic: Updating File Wrappers

	// Indicates whether the contents of a file wrapper matches a directory, regular file, or symbolic link on disk.
	MatchesContentsOfURL(url INSURL) bool
	// Recursively rereads the entire contents of a file wrapper from the specified location on disk.
	ReadFromURLOptionsError(url INSURL, options NSFileWrapperReadingOptions) (bool, error)

	// Topic: Serializing

	// The contents of the file wrapper as an opaque data object.
	SerializedRepresentation() INSData

	// Topic: Accessing Files

	// The filename of the file wrapper object
	Filename() string
	SetFilename(value string)
	// The preferred filename for the file wrapper object.
	PreferredFilename() string
	SetPreferredFilename(value string)
	// A dictionary of file attributes.
	FileAttributes() INSDictionary
	SetFileAttributes(value INSDictionary)
	// The contents of the file-system node associated with a regular-file file wrapper.
	RegularFileContents() INSData

	// Topic: Writing Files

	// Recursively writes the entire contents of a file wrapper to a given file-system URL.
	WriteToURLOptionsOriginalContentsURLError(url INSURL, options NSFileWrapperWritingOptions, originalContentsURL INSURL) (bool, error)

	// Topic: Working with Icons

	// The icon that represents the file wrapper.
	Icon() objc.ID
	SetIcon(value objc.ID)
}

// Init initializes the instance.
func (f FileWrapper) Init() FileWrapper {
	rv := objc.Send[FileWrapper](f.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f FileWrapper) Autorelease() FileWrapper {
	rv := objc.Send[FileWrapper](f.ID, objc.Sel("autorelease"))
	return rv
}

// NewFileWrapper creates a new FileWrapper instance.
func NewFileWrapper() FileWrapper {
	class := getFileWrapperClass()
	rv := objc.Send[FileWrapper](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Initializes the receiver as a directory file wrapper, with a given
// file-wrapper list.
//
// childrenByPreferredName: Key-value dictionary of file wrappers with which to initialize the
// receiver. The dictionary must contain entries whose values are the file
// wrappers that are to become children and whose keys are filenames. See
// [Accessing File Wrapper Identities] in [File System Programming Guide] for
// more information about the file-wrapper list structure.
// //
// [Accessing File Wrapper Identities]: https://developer.apple.com/library/archive/documentation/FileManagement/Conceptual/FileSystemProgrammingGuide/FileWrappers/FileWrappers.html#//apple_ref/doc/uid/TP40010672-CH13-SW1
// [File System Programming Guide]: https://developer.apple.com/library/archive/documentation/FileManagement/Conceptual/FileSystemProgrammingGuide/Introduction/Introduction.html#//apple_ref/doc/uid/TP40010672
//
// # Return Value
// 
// Initialized file wrapper for `fileWrappers`.
//
// # Discussion
// 
// After initialization, the file wrapper is not associated with a file-system
// node until you save it using [WriteToURLOptionsOriginalContentsURLError].
// 
// The receiver is initialized with open permissions: anyone can read, write,
// or modify the directory on disk.
// 
// If any file wrapper in the directory doesn’t have a preferred filename,
// its preferred name is automatically set to its corresponding key in the
// `childrenByPreferredName` dictionary.
//
// See: https://developer.apple.com/documentation/Foundation/FileWrapper/init(directoryWithFileWrappers:)
func NewFileWrapperDirectoryWithFileWrappers(childrenByPreferredName INSDictionary) FileWrapper {
	instance := getFileWrapperClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initDirectoryWithFileWrappers:"), childrenByPreferredName)
	return FileWrapperFromID(rv)
}

// Initializes the receiver as a regular-file file wrapper.
//
// contents: Contents of the file.
//
// # Return Value
// 
// Initialized regular-file file wrapper containing `contents`.
//
// # Discussion
// 
// After initialization, the file wrapper is not associated with a file-system
// node until you save it using [WriteToURLOptionsOriginalContentsURLError].
// 
// The file wrapper is initialized with open permissions: anyone can write to
// or read the file wrapper.
//
// See: https://developer.apple.com/documentation/Foundation/FileWrapper/init(regularFileWithContents:)
func NewFileWrapperRegularFileWithContents(contents INSData) FileWrapper {
	instance := getFileWrapperClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initRegularFileWithContents:"), contents)
	return FileWrapperFromID(rv)
}

// Initializes the receiver as a symbolic-link file wrapper that links to a
// specified file.
//
// url: URL of the file the file wrapper is to reference.
//
// # Return Value
// 
// Initialized symbolic-link file wrapper referencing `url`.
//
// # Discussion
// 
// The file wrapper is not associated with a file-system node until you save
// it using [WriteToURLOptionsOriginalContentsURLError].
// 
// The file wrapper is initialized with open permissions: anyone can modify or
// read the file reference. .
//
// See: https://developer.apple.com/documentation/Foundation/FileWrapper/init(symbolicLinkWithDestinationURL:)
func NewFileWrapperSymbolicLinkWithDestinationURL(url INSURL) FileWrapper {
	instance := getFileWrapperClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initSymbolicLinkWithDestinationURL:"), url)
	return FileWrapperFromID(rv)
}

//
// See: https://developer.apple.com/documentation/Foundation/FileWrapper/init(coder:)
func NewFileWrapperWithCoder(inCoder INSCoder) FileWrapper {
	instance := getFileWrapperClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), inCoder)
	return FileWrapperFromID(rv)
}

// Initializes the receiver as a regular-file file wrapper from given
// serialized data.
//
// serializeRepresentation: Serialized representation of a file wrapper in the format used for the
// [NSFileContentsPboardType] pasteboard type. Data of this format is returned
// by such methods as [SerializedRepresentation] and
// [RTFDFromRangeDocumentAttributes] ([NSAttributedString]).
//
// # Return Value
// 
// Regular-file file wrapper initialized from `serializedRepresentation`.
//
// # Discussion
// 
// The file wrapper is not associated with a file-system node until you save
// it using [WriteToURLOptionsOriginalContentsURLError].
//
// See: https://developer.apple.com/documentation/Foundation/FileWrapper/init(serializedRepresentation:)
func NewFileWrapperWithSerializedRepresentation(serializeRepresentation INSData) FileWrapper {
	instance := getFileWrapperClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSerializedRepresentation:"), serializeRepresentation)
	return FileWrapperFromID(rv)
}

// Initializes a file wrapper instance whose kind is determined by the type of
// file-system node located by the URL.
//
// url: URL of the file-system node the file wrapper is to represent.
//
// options: Option flags for reading the node located at `url`. See
// [FileWrapper.ReadingOptions] for possible values.
// //
// [FileWrapper.ReadingOptions]: https://developer.apple.com/documentation/Foundation/FileWrapper/ReadingOptions
//
// # Return Value
// 
// File wrapper for the file-system node at `url`. May be a directory, file,
// or symbolic link, depending on what is located at the URL. Returns [false]
// (0) if reading is not successful.
//
// [false]: https://developer.apple.com/documentation/Swift/false
//
// # Discussion
// 
// If `url` is a directory, this method recursively creates file wrappers for
// each node within that directory. Use the [FileWrappers] property to get the
// file wrappers of the nodes contained by the directory.
//
// See: https://developer.apple.com/documentation/Foundation/FileWrapper/init(url:options:)
func NewFileWrapperWithURLOptionsError(url INSURL, options NSFileWrapperReadingOptions) (FileWrapper, error) {
	var errorPtr objc.ID
	instance := getFileWrapperClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithURL:options:error:"), url, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return FileWrapper{}, NSErrorFrom(errorPtr)
	}
	return FileWrapperFromID(rv), nil
}

// Initializes a file wrapper instance whose kind is determined by the type of
// file-system node located by the URL.
//
// url: URL of the file-system node the file wrapper is to represent.
//
// options: Option flags for reading the node located at `url`. See
// [FileWrapper.ReadingOptions] for possible values.
// //
// [FileWrapper.ReadingOptions]: https://developer.apple.com/documentation/Foundation/FileWrapper/ReadingOptions
//
// # Return Value
// 
// File wrapper for the file-system node at `url`. May be a directory, file,
// or symbolic link, depending on what is located at the URL. Returns [false]
// (0) if reading is not successful.
//
// [false]: https://developer.apple.com/documentation/Swift/false
//
// # Discussion
// 
// If `url` is a directory, this method recursively creates file wrappers for
// each node within that directory. Use the [FileWrappers] property to get the
// file wrappers of the nodes contained by the directory.
//
// See: https://developer.apple.com/documentation/Foundation/FileWrapper/init(url:options:)
func (f FileWrapper) InitWithURLOptionsError(url INSURL, options NSFileWrapperReadingOptions) (FileWrapper, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](f.ID, objc.Sel("initWithURL:options:error:"), url, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return FileWrapper{}, NSErrorFrom(errorPtr)
	}
	return NSFileWrapperFromID(rv), nil

}
// Initializes the receiver as a directory file wrapper, with a given
// file-wrapper list.
//
// childrenByPreferredName: Key-value dictionary of file wrappers with which to initialize the
// receiver. The dictionary must contain entries whose values are the file
// wrappers that are to become children and whose keys are filenames. See
// [Accessing File Wrapper Identities] in [File System Programming Guide] for
// more information about the file-wrapper list structure.
// //
// [Accessing File Wrapper Identities]: https://developer.apple.com/library/archive/documentation/FileManagement/Conceptual/FileSystemProgrammingGuide/FileWrappers/FileWrappers.html#//apple_ref/doc/uid/TP40010672-CH13-SW1
// [File System Programming Guide]: https://developer.apple.com/library/archive/documentation/FileManagement/Conceptual/FileSystemProgrammingGuide/Introduction/Introduction.html#//apple_ref/doc/uid/TP40010672
//
// # Return Value
// 
// Initialized file wrapper for `fileWrappers`.
//
// # Discussion
// 
// After initialization, the file wrapper is not associated with a file-system
// node until you save it using [WriteToURLOptionsOriginalContentsURLError].
// 
// The receiver is initialized with open permissions: anyone can read, write,
// or modify the directory on disk.
// 
// If any file wrapper in the directory doesn’t have a preferred filename,
// its preferred name is automatically set to its corresponding key in the
// `childrenByPreferredName` dictionary.
//
// See: https://developer.apple.com/documentation/Foundation/FileWrapper/init(directoryWithFileWrappers:)
func (f FileWrapper) InitDirectoryWithFileWrappers(childrenByPreferredName INSDictionary) FileWrapper {
	rv := objc.Send[FileWrapper](f.ID, objc.Sel("initDirectoryWithFileWrappers:"), childrenByPreferredName)
	return rv
}
// Initializes the receiver as a regular-file file wrapper.
//
// contents: Contents of the file.
//
// # Return Value
// 
// Initialized regular-file file wrapper containing `contents`.
//
// # Discussion
// 
// After initialization, the file wrapper is not associated with a file-system
// node until you save it using [WriteToURLOptionsOriginalContentsURLError].
// 
// The file wrapper is initialized with open permissions: anyone can write to
// or read the file wrapper.
//
// See: https://developer.apple.com/documentation/Foundation/FileWrapper/init(regularFileWithContents:)
func (f FileWrapper) InitRegularFileWithContents(contents INSData) FileWrapper {
	rv := objc.Send[FileWrapper](f.ID, objc.Sel("initRegularFileWithContents:"), contents)
	return rv
}
// Initializes the receiver as a symbolic-link file wrapper that links to a
// specified file.
//
// url: URL of the file the file wrapper is to reference.
//
// # Return Value
// 
// Initialized symbolic-link file wrapper referencing `url`.
//
// # Discussion
// 
// The file wrapper is not associated with a file-system node until you save
// it using [WriteToURLOptionsOriginalContentsURLError].
// 
// The file wrapper is initialized with open permissions: anyone can modify or
// read the file reference. .
//
// See: https://developer.apple.com/documentation/Foundation/FileWrapper/init(symbolicLinkWithDestinationURL:)
func (f FileWrapper) InitSymbolicLinkWithDestinationURL(url INSURL) FileWrapper {
	rv := objc.Send[FileWrapper](f.ID, objc.Sel("initSymbolicLinkWithDestinationURL:"), url)
	return rv
}
// Initializes the receiver as a regular-file file wrapper from given
// serialized data.
//
// serializeRepresentation: Serialized representation of a file wrapper in the format used for the
// [NSFileContentsPboardType] pasteboard type. Data of this format is returned
// by such methods as [SerializedRepresentation] and
// [RTFDFromRangeDocumentAttributes] ([NSAttributedString]).
//
// # Return Value
// 
// Regular-file file wrapper initialized from `serializedRepresentation`.
//
// # Discussion
// 
// The file wrapper is not associated with a file-system node until you save
// it using [WriteToURLOptionsOriginalContentsURLError].
//
// See: https://developer.apple.com/documentation/Foundation/FileWrapper/init(serializedRepresentation:)
func (f FileWrapper) InitWithSerializedRepresentation(serializeRepresentation INSData) FileWrapper {
	rv := objc.Send[FileWrapper](f.ID, objc.Sel("initWithSerializedRepresentation:"), serializeRepresentation)
	return rv
}
// Adds a child file wrapper to the receiver, which must be a directory file
// wrapper.
//
// child: File wrapper to add to the directory.
//
// # Return Value
// 
// Dictionary key used to store `fileWrapper` in the directory’s list of
// file wrappers. The dictionary key is a unique filename, which is the same
// as the passed-in file wrapper’s preferred filename unless that name is
// already in use as a key in the directory’s dictionary of children. See
// [Accessing File Wrapper Identities] in [File System Programming Guide] for
// more information about the file-wrapper list structure.
//
// [Accessing File Wrapper Identities]: https://developer.apple.com/library/archive/documentation/FileManagement/Conceptual/FileSystemProgrammingGuide/FileWrappers/FileWrappers.html#//apple_ref/doc/uid/TP40010672-CH13-SW1
// [File System Programming Guide]: https://developer.apple.com/library/archive/documentation/FileManagement/Conceptual/FileSystemProgrammingGuide/Introduction/Introduction.html#//apple_ref/doc/uid/TP40010672
//
// # Discussion
// 
// Use this method to add an existing file wrapper as a child of a directory
// file wrapper. If the file wrapper does not have a preferred filename, set
// the [PreferredFilename] property to give it one before calling
// [AddFileWrapper]. To create a new file wrapper and add it to a directory,
// use the [AddRegularFileWithContentsPreferredFilename] method.
// 
// # Special Considerations
// 
// This method raises [NSInternalInconsistencyException] if the receiver is
// not a directory file wrapper.
// 
// This method raises [NSInvalidArgumentException] if the child file wrapper
// doesn’t have a preferred name.
//
// See: https://developer.apple.com/documentation/Foundation/FileWrapper/addFileWrapper(_:)
func (f FileWrapper) AddFileWrapper(child INSFileWrapper) string {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("addFileWrapper:"), child)
	return NSStringFromID(rv).String()
}
// Removes a child file wrapper from the receiver, which must be a directory
// file wrapper.
//
// child: File wrapper to remove from the directory.
//
// # Discussion
// 
// This method raises [NSInternalInconsistencyException] if the receiver is
// not a directory file wrapper.
//
// See: https://developer.apple.com/documentation/Foundation/FileWrapper/removeFileWrapper(_:)
func (f FileWrapper) RemoveFileWrapper(child INSFileWrapper) {
	objc.Send[objc.ID](f.ID, objc.Sel("removeFileWrapper:"), child)
}
// Creates a regular-file file wrapper with the given contents and adds it to
// the receiver, which must be a directory file wrapper.
//
// data: Contents for the new regular-file file wrapper.
//
// fileName: Preferred filename for the new regular-file file wrapper.
//
// # Return Value
// 
// Dictionary key used to store the new file wrapper in the directory’s list
// of file wrappers. The dictionary key is a unique filename, which is the
// same as the passed-in file wrapper’s preferred filename unless that name
// is already in use as a key in the directory’s dictionary of children. See
// [Accessing File Wrapper Identities] in [File System Programming Guide] for
// more information about the file-wrapper list structure.
//
// [Accessing File Wrapper Identities]: https://developer.apple.com/library/archive/documentation/FileManagement/Conceptual/FileSystemProgrammingGuide/FileWrappers/FileWrappers.html#//apple_ref/doc/uid/TP40010672-CH13-SW1
// [File System Programming Guide]: https://developer.apple.com/library/archive/documentation/FileManagement/Conceptual/FileSystemProgrammingGuide/Introduction/Introduction.html#//apple_ref/doc/uid/TP40010672
//
// # Discussion
// 
// This is a convenience method. The default implementation allocates a new
// file wrapper, initializes it with [InitRegularFileWithContents], set its
// [PreferredFilename] property, adds it to the directory with
// [AddFileWrapper], and returns what [AddFileWrapper] returned.
// 
// # Special Considerations
// 
// This method raises [NSInternalInconsistencyException] if the receiver is
// not a directory file wrapper.
// 
// This method raises [NSInvalidArgumentException] if you pass `nil` or an
// empty value for `filename`.
//
// See: https://developer.apple.com/documentation/Foundation/FileWrapper/addRegularFile(withContents:preferredFilename:)
func (f FileWrapper) AddRegularFileWithContentsPreferredFilename(data INSData, fileName string) string {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("addRegularFileWithContents:preferredFilename:"), data, objc.String(fileName))
	return NSStringFromID(rv).String()
}
// Returns the dictionary key used by a directory to identify a given file
// wrapper.
//
// child: The child file wrapper for which you want the key.
//
// # Return Value
// 
// Dictionary key used to store the file wrapper in the directory’s list of
// file wrappers. The dictionary key is a unique filename, which may not be
// the same as the passed-in file wrapper’s preferred filename if more than
// one file wrapper in the directory’s dictionary of children has the same
// preferred filename. See [Accessing File Wrapper Identities] in [File System
// Programming Guide] for more information about the file-wrapper list
// structure. Returns `nil` if the file wrapper specified in `child` is not a
// child of the directory.
//
// [Accessing File Wrapper Identities]: https://developer.apple.com/library/archive/documentation/FileManagement/Conceptual/FileSystemProgrammingGuide/FileWrappers/FileWrappers.html#//apple_ref/doc/uid/TP40010672-CH13-SW1
// [File System Programming Guide]: https://developer.apple.com/library/archive/documentation/FileManagement/Conceptual/FileSystemProgrammingGuide/Introduction/Introduction.html#//apple_ref/doc/uid/TP40010672
//
// # Discussion
// 
// This method raises [NSInternalInconsistencyException] if the receiver is
// not a directory file wrapper.
//
// See: https://developer.apple.com/documentation/Foundation/FileWrapper/keyForChildFileWrapper(_:)
func (f FileWrapper) KeyForFileWrapper(child INSFileWrapper) string {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("keyForFileWrapper:"), child)
	return NSStringFromID(rv).String()
}
// Indicates whether the contents of a file wrapper matches a directory,
// regular file, or symbolic link on disk.
//
// url: URL of the file-system node with which to compare the file wrapper.
//
// # Return Value
// 
// [true] when the contents of the file wrapper match the contents of `url`,
// [false] otherwise.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// The contents of files are not compared; matching of regular files is based
// on file modification dates. For a directory, children are compared against
// the files in the directory, recursively.
// 
// Because children of directory file wrappers are not read immediately by the
// [InitWithURLOptionsError] method unless the [NSFileWrapperReadingImmediate]
// reading option is used, even a newly-created directory file wrapper might
// not have the same contents as the directory on disk. You can use this
// method to determine whether the file wrapper’s contents in memory need to
// be updated.
// 
// If the file wrapper needs updating, use the [ReadFromURLOptionsError]
// method with the [NSFileWrapperReadingImmediate] reading option.
// 
// This table describes which attributes of the file wrapper and file-system
// node are compared to determine whether the file wrapper matches the node on
// disk:
// 
// [Table data omitted]
//
// See: https://developer.apple.com/documentation/Foundation/FileWrapper/matchesContents(of:)
func (f FileWrapper) MatchesContentsOfURL(url INSURL) bool {
	rv := objc.Send[bool](f.ID, objc.Sel("matchesContentsOfURL:"), url)
	return rv
}
// Recursively rereads the entire contents of a file wrapper from the
// specified location on disk.
//
// url: URL of the file-system node corresponding to the file wrapper.
//
// options: Option flags for reading the node located at `url`. See
// [FileWrapper.ReadingOptions] for possible values.
// //
// [FileWrapper.ReadingOptions]: https://developer.apple.com/documentation/Foundation/FileWrapper/ReadingOptions
//
// # Discussion
// 
// When reading a directory, children are added and removed as necessary to
// match the file system.
//
// See: https://developer.apple.com/documentation/Foundation/FileWrapper/read(from:options:)
func (f FileWrapper) ReadFromURLOptionsError(url INSURL, options NSFileWrapperReadingOptions) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](f.ID, objc.Sel("readFromURL:options:error:"), url, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("readFromURL:options:error: returned NO with nil NSError")
	}
	return rv, nil

}
// Recursively writes the entire contents of a file wrapper to a given
// file-system URL.
//
// url: URL of the file-system node to which the file wrapper’s contents are
// written.
//
// options: Option flags for writing to the node located at `url`. See
// [FileWrapper.WritingOptions] for possible values.
// //
// [FileWrapper.WritingOptions]: https://developer.apple.com/documentation/Foundation/FileWrapper/WritingOptions
//
// originalContentsURL: The location of a previous revision of the contents being written. The
// default implementation of this method attempts to avoid unnecessary I/O by
// writing hard links to regular files instead of actually writing out their
// contents when the contents have not changed. The child file wrappers must
// return accurate values when its [Filename] property is accessed for this to
// work. Use the [NSFileWrapperWritingWithNameUpdating] writing option to
// increase the likelihood of that.
// 
// Specify `nil` for this parameter if there is no earlier version of the
// contents or if you want to ensure that all the contents are written to
// files.
//
// # Discussion
//
// See: https://developer.apple.com/documentation/Foundation/FileWrapper/write(to:options:originalContentsURL:)
func (f FileWrapper) WriteToURLOptionsOriginalContentsURLError(url INSURL, options NSFileWrapperWritingOptions, originalContentsURL INSURL) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](f.ID, objc.Sel("writeToURL:options:originalContentsURL:error:"), url, options, originalContentsURL, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("writeToURL:options:originalContentsURL:error: returned NO with nil NSError")
	}
	return rv, nil

}
//
// See: https://developer.apple.com/documentation/Foundation/FileWrapper/init(coder:)
func (f FileWrapper) InitWithCoder(inCoder INSCoder) FileWrapper {
	rv := objc.Send[FileWrapper](f.ID, objc.Sel("initWithCoder:"), inCoder)
	return rv
}
// Encodes the receiver using a given archiver.
//
// coder: An archiver object.
//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/encode(with:)
func (f FileWrapper) EncodeWithCoder(coder INSCoder) {
	objc.Send[objc.ID](f.ID, objc.Sel("encodeWithCoder:"), coder)
}

// This property contains a boolean value that indicates whether the file
// wrapper object is a regular-file.
//
// # Discussion
// 
// This property contains [true] when the file wrapper object is a
// regular-file wrapper, otherwise it contains [false]. Invocations of
// [ReadFromURLOptionsError] may change the value of this property if the type
// of the file on disk has changed.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Foundation/FileWrapper/isRegularFile
func (f FileWrapper) RegularFile() bool {
	rv := objc.Send[bool](f.ID, objc.Sel("isRegularFile"))
	return rv
}
// This property contains a boolean value indicating whether the file wrapper
// is a directory file wrapper.
//
// # Discussion
// 
// This property will contain YES when the file wrapper is a directory file
// wrapper, otherwise it contains NO.
//
// See: https://developer.apple.com/documentation/Foundation/FileWrapper/isDirectory
func (f FileWrapper) Directory() bool {
	rv := objc.Send[bool](f.ID, objc.Sel("isDirectory"))
	return rv
}
// A boolean that indicates whether the file wrapper object is a symbolic-link
// file wrapper.
//
// # Discussion
// 
// This property contains [true] when the file wrapper object is a
// symbolic-link file wrapper, [false] otherwise.
// 
// Invocations of [ReadFromURLOptionsError] may change the value contained by
// this property, if the type of the file on disk has changed.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Foundation/FileWrapper/isSymbolicLink
func (f FileWrapper) SymbolicLink() bool {
	rv := objc.Send[bool](f.ID, objc.Sel("isSymbolicLink"))
	return rv
}
// The file wrappers contained by a directory file wrapper.
//
// # Discussion
// 
// The dictionary contains entries whose values are the file wrappers and
// whose keys are the unique filenames that have been assigned to each one.
// See [Accessing File Wrapper Identities] in [File System Programming Guide]
// for more information about the file-wrapper list structure.
// 
// This property may contain `nil` if the user modifies the directory after
// you call [ReadFromURLOptionsError] or [InitWithURLOptionsError] but before
// [NSFileWrapper] has read the contents of the directory. Use the
// [FileWrapperReadingImmediate] reading option to reduce the likelihood of
// that problem.
// 
// # Special Considerations
// 
// This property raises [NSInternalInconsistencyException] if the file wrapper
// object is not a directory file wrapper.
//
// [Accessing File Wrapper Identities]: https://developer.apple.com/library/archive/documentation/FileManagement/Conceptual/FileSystemProgrammingGuide/FileWrappers/FileWrappers.html#//apple_ref/doc/uid/TP40010672-CH13-SW1
// [File System Programming Guide]: https://developer.apple.com/library/archive/documentation/FileManagement/Conceptual/FileSystemProgrammingGuide/Introduction/Introduction.html#//apple_ref/doc/uid/TP40010672
//
// See: https://developer.apple.com/documentation/Foundation/FileWrapper/fileWrappers
func (f FileWrapper) FileWrappers() INSDictionary {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("fileWrappers"))
	return NSDictionaryFromID(objc.ID(rv))
}
// The URL referenced by the file wrapper object, which must be a
// symbolic-link file wrapper.
//
// # Discussion
// 
// This property may contain `nil` if the user modifies the symbolic link
// after you call [ReadFromURLOptionsError] or [InitWithURLOptionsError] but
// before [NSFileWrapper] has read the contents of the link. Use the
// [FileWrapperReadingImmediate] reading option to reduce the likelihood of
// that problem.
// 
// # Special Considerations
// 
// This property raises [NSInternalInconsistencyException] if the file wrapper
// object is not a symbolic-link file wrapper.
//
// See: https://developer.apple.com/documentation/Foundation/FileWrapper/symbolicLinkDestinationURL
func (f FileWrapper) SymbolicLinkDestinationURL() INSURL {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("symbolicLinkDestinationURL"))
	return NSURLFromID(objc.ID(rv))
}
// The contents of the file wrapper as an opaque data object.
//
// # Discussion
// 
// This property contains a data object in the format used by the
// [fileContents] pasteboard type. This data object is also suitable for
// passing to [InitWithSerializedRepresentation].
// 
// This property may be `nil` if the user modifies the contents of the file
// system node after you call [ReadFromURLOptionsError] or
// [InitWithURLOptionsError], but before [NSFileWrapper] has read the contents
// of the file. You can use the [FileWrapperReadingImmediate] reading option
// to reduce the likelihood of this problem.
//
// [fileContents]: https://developer.apple.com/documentation/AppKit/NSPasteboard/PasteboardType/fileContents
//
// See: https://developer.apple.com/documentation/Foundation/FileWrapper/serializedRepresentation
func (f FileWrapper) SerializedRepresentation() INSData {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("serializedRepresentation"))
	return NSDataFromID(objc.ID(rv))
}
// The filename of the file wrapper object
//
// # Discussion
// 
// This property contains the file wrapper’s filename, or `nil` when the
// file wrapper has no corresponding file-system node.
// 
// The filename is used for record-keeping purposes only and is set
// automatically when the file wrapper is created from the file system using
// [InitWithURLOptionsError] and when it’s saved to the file system using
// [WriteToURLOptionsOriginalContentsURLError] (although this method allows
// you to request that the filename not be updated).
// 
// The filename is usually the same as the preferred filename, but might
// instead be a name derived from the preferred filename. You can use this
// method to get the name of a child that’s just been read. Don’t use this
// method to get the name of a child that’s about to be written, because the
// name might be about to change; send [KeyForFileWrapper] to the parent
// instead.
//
// See: https://developer.apple.com/documentation/Foundation/FileWrapper/filename
func (f FileWrapper) Filename() string {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("filename"))
	return NSStringFromID(rv).String()
}
func (f FileWrapper) SetFilename(value string) {
	objc.Send[struct{}](f.ID, objc.Sel("setFilename:"), objc.String(value))
}
// The preferred filename for the file wrapper object.
//
// # Discussion
// 
// This name is normally used as the dictionary key when a child file wrapper
// is added to a directory file wrapper. However, if another file wrapper with
// the same preferred name already exists in the directory file wrapper when
// this object is added, the filename assigned as the dictionary key may
// differ from the preferred filename.
// 
// When you change the preferred filename, the default implementation of this
// property causes the existing parent directory file wrappers to remove and
// re-add the child to accommodate the change. Preferred filenames of children
// are not preserved when you write a file wrapper to disk and then later
// instantiate another file wrapper by reading the file from disk. If you need
// to preserve the user-visible names of attachments, you have to store the
// names yourself.
//
// See: https://developer.apple.com/documentation/Foundation/FileWrapper/preferredFilename
func (f FileWrapper) PreferredFilename() string {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("preferredFilename"))
	return NSStringFromID(rv).String()
}
func (f FileWrapper) SetPreferredFilename(value string) {
	objc.Send[struct{}](f.ID, objc.Sel("setPreferredFilename:"), objc.String(value))
}
// A dictionary of file attributes.
//
// # Discussion
// 
// The file attributes’ dictionary is the same format as that returned by
// [AttributesOfItemAtPathError] ([NSFileManager]).
//
// See: https://developer.apple.com/documentation/Foundation/FileWrapper/fileAttributes
func (f FileWrapper) FileAttributes() INSDictionary {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("fileAttributes"))
	return NSDictionaryFromID(objc.ID(rv))
}
func (f FileWrapper) SetFileAttributes(value INSDictionary) {
	objc.Send[struct{}](f.ID, objc.Sel("setFileAttributes:"), value)
}
// The contents of the file-system node associated with a regular-file file
// wrapper.
//
// # Discussion
// 
// This property may contain `nil` if the user modifies the file after you
// call [ReadFromURLOptionsError] or [InitWithURLOptionsError] but before
// [NSFileWrapper] has read the contents of the file. Use the
// [FileWrapperReadingImmediate] reading option to reduce the likelihood of
// that problem.
// 
// # Special Considerations
// 
// This property raises [NSInternalInconsistencyException] if the file wrapper
// object is not a regular-file file wrapper.
//
// See: https://developer.apple.com/documentation/Foundation/FileWrapper/regularFileContents
func (f FileWrapper) RegularFileContents() INSData {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("regularFileContents"))
	return NSDataFromID(objc.ID(rv))
}
// The icon that represents the file wrapper.
//
// # Discussion
// 
// You don’t have to use this icon; for example, a file viewer typically
// looks up icons automatically based on file extensions, and so wouldn’t
// need this one. Similarly, if a file wrapper represents an image file, you
// can display the image directly rather than a file icon.
// 
// This method may return `nil` if the file wrapper is a child created when
// its parent was read from the file system, and the child was modified before
// it was read. Use the [NSFileWrapperReadingImmediate] reading option to
// reduce the likelihood of that problem.
// 
// Because the [NSImage] object that’s returned might be shared by many
// [NSFileWrapper] objects, you must not mutate it. If you need to mutate the
// returned object, make a copy first and mutate the copy instead.
//
// See: https://developer.apple.com/documentation/Foundation/FileWrapper/icon
func (f FileWrapper) Icon() objc.ID {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("icon"))
	return rv
}
func (f FileWrapper) SetIcon(value objc.ID) {
	objc.Send[struct{}](f.ID, objc.Sel("setIcon:"), value)
}

			// Protocol methods for NSSecureCoding
			

