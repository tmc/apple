// Code generated from Apple documentation for FileProvider. DO NOT EDIT.

package fileprovider

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSFileProviderKnownFolderLocations] class.
var (
	_NSFileProviderKnownFolderLocationsClass     NSFileProviderKnownFolderLocationsClass
	_NSFileProviderKnownFolderLocationsClassOnce sync.Once
)

func getNSFileProviderKnownFolderLocationsClass() NSFileProviderKnownFolderLocationsClass {
	_NSFileProviderKnownFolderLocationsClassOnce.Do(func() {
		_NSFileProviderKnownFolderLocationsClass = NSFileProviderKnownFolderLocationsClass{class: objc.GetClass("NSFileProviderKnownFolderLocations")}
	})
	return _NSFileProviderKnownFolderLocationsClass
}

// GetNSFileProviderKnownFolderLocationsClass returns the class object for NSFileProviderKnownFolderLocations.
func GetNSFileProviderKnownFolderLocationsClass() NSFileProviderKnownFolderLocationsClass {
	return getNSFileProviderKnownFolderLocationsClass()
}

type NSFileProviderKnownFolderLocationsClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSFileProviderKnownFolderLocationsClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSFileProviderKnownFolderLocationsClass) Alloc() NSFileProviderKnownFolderLocations {
	rv := objc.Send[NSFileProviderKnownFolderLocations](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A class for working with known-folder locations.
//
// # Identifying known-folder locations
//
//   - [NSFileProviderKnownFolderLocations.DesktopLocation]
//   - [NSFileProviderKnownFolderLocations.SetDesktopLocation]
//   - [NSFileProviderKnownFolderLocations.DocumentsLocation]
//   - [NSFileProviderKnownFolderLocations.SetDocumentsLocation]
//
// # Configuring folder options
//
//   - [NSFileProviderKnownFolderLocations.ShouldCreateBinaryCompatibilitySymlink]
//   - [NSFileProviderKnownFolderLocations.SetShouldCreateBinaryCompatibilitySymlink]
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderKnownFolderLocations
type NSFileProviderKnownFolderLocations struct {
	objectivec.Object
}

// NSFileProviderKnownFolderLocationsFromID constructs a [NSFileProviderKnownFolderLocations] from an objc.ID.
//
// A class for working with known-folder locations.
func NSFileProviderKnownFolderLocationsFromID(id objc.ID) NSFileProviderKnownFolderLocations {
	return NSFileProviderKnownFolderLocations{objectivec.Object{ID: id}}
}

// NOTE: NSFileProviderKnownFolderLocations adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSFileProviderKnownFolderLocations] class.
//
// # Identifying known-folder locations
//
//   - [INSFileProviderKnownFolderLocations.DesktopLocation]
//   - [INSFileProviderKnownFolderLocations.SetDesktopLocation]
//   - [INSFileProviderKnownFolderLocations.DocumentsLocation]
//   - [INSFileProviderKnownFolderLocations.SetDocumentsLocation]
//
// # Configuring folder options
//
//   - [INSFileProviderKnownFolderLocations.ShouldCreateBinaryCompatibilitySymlink]
//   - [INSFileProviderKnownFolderLocations.SetShouldCreateBinaryCompatibilitySymlink]
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderKnownFolderLocations
type INSFileProviderKnownFolderLocations interface {
	objectivec.IObject

	// Topic: Identifying known-folder locations

	DesktopLocation() INSFileProviderKnownFolderLocation
	SetDesktopLocation(value INSFileProviderKnownFolderLocation)
	DocumentsLocation() INSFileProviderKnownFolderLocation
	SetDocumentsLocation(value INSFileProviderKnownFolderLocation)

	// Topic: Configuring folder options

	ShouldCreateBinaryCompatibilitySymlink() bool
	SetShouldCreateBinaryCompatibilitySymlink(value bool)
}

// Init initializes the instance.
func (f NSFileProviderKnownFolderLocations) Init() NSFileProviderKnownFolderLocations {
	rv := objc.Send[NSFileProviderKnownFolderLocations](f.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f NSFileProviderKnownFolderLocations) Autorelease() NSFileProviderKnownFolderLocations {
	rv := objc.Send[NSFileProviderKnownFolderLocations](f.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSFileProviderKnownFolderLocations creates a new NSFileProviderKnownFolderLocations instance.
func NewNSFileProviderKnownFolderLocations() NSFileProviderKnownFolderLocations {
	class := getNSFileProviderKnownFolderLocationsClass()
	rv := objc.Send[NSFileProviderKnownFolderLocations](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderKnownFolderLocations/desktopLocation
func (f NSFileProviderKnownFolderLocations) DesktopLocation() INSFileProviderKnownFolderLocation {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("desktopLocation"))
	return NSFileProviderKnownFolderLocationFromID(objc.ID(rv))
}
func (f NSFileProviderKnownFolderLocations) SetDesktopLocation(value INSFileProviderKnownFolderLocation) {
	objc.Send[struct{}](f.ID, objc.Sel("setDesktopLocation:"), value)
}

// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderKnownFolderLocations/documentsLocation
func (f NSFileProviderKnownFolderLocations) DocumentsLocation() INSFileProviderKnownFolderLocation {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("documentsLocation"))
	return NSFileProviderKnownFolderLocationFromID(objc.ID(rv))
}
func (f NSFileProviderKnownFolderLocations) SetDocumentsLocation(value INSFileProviderKnownFolderLocation) {
	objc.Send[struct{}](f.ID, objc.Sel("setDocumentsLocation:"), value)
}

// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderKnownFolderLocations/shouldCreateBinaryCompatibilitySymlink
func (f NSFileProviderKnownFolderLocations) ShouldCreateBinaryCompatibilitySymlink() bool {
	rv := objc.Send[bool](f.ID, objc.Sel("shouldCreateBinaryCompatibilitySymlink"))
	return rv
}
func (f NSFileProviderKnownFolderLocations) SetShouldCreateBinaryCompatibilitySymlink(value bool) {
	objc.Send[struct{}](f.ID, objc.Sel("setShouldCreateBinaryCompatibilitySymlink:"), value)
}
