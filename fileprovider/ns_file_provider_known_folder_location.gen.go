// Code generated from Apple documentation for FileProvider. DO NOT EDIT.

package fileprovider

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSFileProviderKnownFolderLocation] class.
var (
	_NSFileProviderKnownFolderLocationClass     NSFileProviderKnownFolderLocationClass
	_NSFileProviderKnownFolderLocationClassOnce sync.Once
)

func getNSFileProviderKnownFolderLocationClass() NSFileProviderKnownFolderLocationClass {
	_NSFileProviderKnownFolderLocationClassOnce.Do(func() {
		_NSFileProviderKnownFolderLocationClass = NSFileProviderKnownFolderLocationClass{class: objc.GetClass("NSFileProviderKnownFolderLocation")}
	})
	return _NSFileProviderKnownFolderLocationClass
}

// GetNSFileProviderKnownFolderLocationClass returns the class object for NSFileProviderKnownFolderLocation.
func GetNSFileProviderKnownFolderLocationClass() NSFileProviderKnownFolderLocationClass {
	return getNSFileProviderKnownFolderLocationClass()
}

type NSFileProviderKnownFolderLocationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSFileProviderKnownFolderLocationClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSFileProviderKnownFolderLocationClass) Alloc() NSFileProviderKnownFolderLocation {
	rv := objc.Send[NSFileProviderKnownFolderLocation](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// # Initializers
//
//   - [NSFileProviderKnownFolderLocation.InitWithExistingItemIdentifier]: Initialize a location with the item identifier of a folder that already exists on the server.
//   - [NSFileProviderKnownFolderLocation.InitWithParentItemIdentifierFilename]: Initialize a location with the filename of the folder in a specified parent.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderKnownFolderLocations/Location
type NSFileProviderKnownFolderLocation struct {
	objectivec.Object
}

// NSFileProviderKnownFolderLocationFromID constructs a [NSFileProviderKnownFolderLocation] from an objc.ID.
func NSFileProviderKnownFolderLocationFromID(id objc.ID) NSFileProviderKnownFolderLocation {
	return NSFileProviderKnownFolderLocation{objectivec.Object{ID: id}}
}

// NOTE: NSFileProviderKnownFolderLocation adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSFileProviderKnownFolderLocation] class.
//
// # Initializers
//
//   - [INSFileProviderKnownFolderLocation.InitWithExistingItemIdentifier]: Initialize a location with the item identifier of a folder that already exists on the server.
//   - [INSFileProviderKnownFolderLocation.InitWithParentItemIdentifierFilename]: Initialize a location with the filename of the folder in a specified parent.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderKnownFolderLocations/Location
type INSFileProviderKnownFolderLocation interface {
	objectivec.IObject

	// Topic: Initializers

	// Initialize a location with the item identifier of a folder that already exists on the server.
	InitWithExistingItemIdentifier(existingItemIdentifier NSFileProviderItemIdentifier) NSFileProviderKnownFolderLocation
	// Initialize a location with the filename of the folder in a specified parent.
	InitWithParentItemIdentifierFilename(parentItemIdentifier NSFileProviderItemIdentifier, filename string) NSFileProviderKnownFolderLocation
}

// Init initializes the instance.
func (f NSFileProviderKnownFolderLocation) Init() NSFileProviderKnownFolderLocation {
	rv := objc.Send[NSFileProviderKnownFolderLocation](f.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f NSFileProviderKnownFolderLocation) Autorelease() NSFileProviderKnownFolderLocation {
	rv := objc.Send[NSFileProviderKnownFolderLocation](f.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSFileProviderKnownFolderLocation creates a new NSFileProviderKnownFolderLocation instance.
func NewNSFileProviderKnownFolderLocation() NSFileProviderKnownFolderLocation {
	class := getNSFileProviderKnownFolderLocationClass()
	rv := objc.Send[NSFileProviderKnownFolderLocation](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Initialize a location with the item identifier of a folder that already
// exists on the server.
//
// # Discussion
//
// If the known folder already exists on the server, the provider can specify
// the exact identifier of the item that needs to be used to back the known
// folder.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderKnownFolderLocations/Location/init(existingItemIdentifier:)
func NewFileProviderKnownFolderLocationWithExistingItemIdentifier(existingItemIdentifier NSFileProviderItemIdentifier) NSFileProviderKnownFolderLocation {
	instance := getNSFileProviderKnownFolderLocationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithExistingItemIdentifier:"), objc.String(string(existingItemIdentifier)))
	return NSFileProviderKnownFolderLocationFromID(rv)
}

// Initialize a location with the filename of the folder in a specified
// parent.
//
// # Discussion
//
// When replicating a known folder the system will reuse a folder located at
// the specified filename within the parent if one exists, or create a new
// item at this location if none exists yet.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderKnownFolderLocations/Location/init(parentItemIdentifier:filename:)
func NewFileProviderKnownFolderLocationWithParentItemIdentifierFilename(parentItemIdentifier NSFileProviderItemIdentifier, filename string) NSFileProviderKnownFolderLocation {
	instance := getNSFileProviderKnownFolderLocationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithParentItemIdentifier:filename:"), objc.String(string(parentItemIdentifier)), objc.String(filename))
	return NSFileProviderKnownFolderLocationFromID(rv)
}

// Initialize a location with the item identifier of a folder that already
// exists on the server.
//
// # Discussion
//
// If the known folder already exists on the server, the provider can specify
// the exact identifier of the item that needs to be used to back the known
// folder.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderKnownFolderLocations/Location/init(existingItemIdentifier:)
func (f NSFileProviderKnownFolderLocation) InitWithExistingItemIdentifier(existingItemIdentifier NSFileProviderItemIdentifier) NSFileProviderKnownFolderLocation {
	rv := objc.Send[NSFileProviderKnownFolderLocation](f.ID, objc.Sel("initWithExistingItemIdentifier:"), objc.String(string(existingItemIdentifier)))
	return rv
}

// Initialize a location with the filename of the folder in a specified
// parent.
//
// # Discussion
//
// When replicating a known folder the system will reuse a folder located at
// the specified filename within the parent if one exists, or create a new
// item at this location if none exists yet.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderKnownFolderLocations/Location/init(parentItemIdentifier:filename:)
func (f NSFileProviderKnownFolderLocation) InitWithParentItemIdentifierFilename(parentItemIdentifier NSFileProviderItemIdentifier, filename string) NSFileProviderKnownFolderLocation {
	rv := objc.Send[NSFileProviderKnownFolderLocation](f.ID, objc.Sel("initWithParentItemIdentifier:filename:"), objc.String(string(parentItemIdentifier)), objc.String(filename))
	return rv
}
