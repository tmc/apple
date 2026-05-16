// Code generated from Apple documentation for FSKit. DO NOT EDIT.

package fskit

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [FSVolume] class.
var (
	_FSVolumeClass     FSVolumeClass
	_FSVolumeClassOnce sync.Once
)

func getFSVolumeClass() FSVolumeClass {
	_FSVolumeClassOnce.Do(func() {
		_FSVolumeClass = FSVolumeClass{class: objc.GetClass("FSVolume")}
	})
	return _FSVolumeClass
}

// GetFSVolumeClass returns the class object for FSVolume.
func GetFSVolumeClass() FSVolumeClass {
	return getFSVolumeClass()
}

type FSVolumeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (fc FSVolumeClass) Class() objc.Class {
	return fc.class
}

// Alloc allocates memory for a new instance of the class.
func (fc FSVolumeClass) Alloc() FSVolume {
	rv := objc.Send[FSVolume](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// A directory structure for files and folders.
//
// # Overview
//
// A file system, depending on its type, provides one or more volumes to
// clients. The [FSUnaryFileSystem] by definition provides only one volume,
// while an [FSFileSystem] supports multiple volumes.
//
// You implement a volume for your file system type by subclassing this class,
// and also conforming to the [FSVolumeOperations] and
// [FSVolumePathConfOperations] protocols. This protocol defines the minimum
// set of operations supported by a volume, such as mounting, activating,
// creating and removing items, and more.
//
// Your volume can provide additional functionality by conforming to other
// volume operations protocols. These protocols add support for operations
// like open and close, read and write, extended attribute (Xattr)
// manipulation, and more.
//
// # Creating a volume
//
//   - [FSVolume.InitWithVolumeIDVolumeName]: Creates a volume with the given identifier and name.
//
// # Accessing volume properties
//
//   - [FSVolume.VolumeID]: An identifier that uniquely identifies the volume.
//   - [FSVolume.Name]: The name of the volume.
//   - [FSVolume.SetName]
//
// See: https://developer.apple.com/documentation/FSKit/FSVolume
type FSVolume struct {
	objectivec.Object
}

// FSVolumeFromID constructs a [FSVolume] from an objc.ID.
//
// A directory structure for files and folders.
func FSVolumeFromID(id objc.ID) FSVolume {
	return FSVolume{objectivec.Object{ID: id}}
}

// NOTE: FSVolume adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [FSVolume] class.
//
// # Creating a volume
//
//   - [IFSVolume.InitWithVolumeIDVolumeName]: Creates a volume with the given identifier and name.
//
// # Accessing volume properties
//
//   - [IFSVolume.VolumeID]: An identifier that uniquely identifies the volume.
//   - [IFSVolume.Name]: The name of the volume.
//   - [IFSVolume.SetName]
//
// See: https://developer.apple.com/documentation/FSKit/FSVolume
type IFSVolume interface {
	objectivec.IObject

	// Topic: Creating a volume

	// Creates a volume with the given identifier and name.
	InitWithVolumeIDVolumeName(volumeID IFSVolumeIdentifier, volumeName IFSFileName) FSVolume

	// Topic: Accessing volume properties

	// An identifier that uniquely identifies the volume.
	VolumeID() IFSVolumeIdentifier
	// The name of the volume.
	Name() IFSFileName
	SetName(value IFSFileName)
}

// Init initializes the instance.
func (v FSVolume) Init() FSVolume {
	rv := objc.Send[FSVolume](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v FSVolume) Autorelease() FSVolume {
	rv := objc.Send[FSVolume](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewFSVolume creates a new FSVolume instance.
func NewFSVolume() FSVolume {
	class := getFSVolumeClass()
	rv := objc.Send[FSVolume](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a volume with the given identifier and name.
//
// volumeID: An [FSVolumeIdentifier] to uniquely identify the volume. For a network file
// system that supports multiple authenticated users, disambiguate the users
// by using qualifying data in the identifier.
//
// volumeName: A name for the volume.
//
// See: https://developer.apple.com/documentation/FSKit/FSVolume/init(volumeID:volumeName:)
func NewVolumeWithVolumeIDVolumeName(volumeID IFSVolumeIdentifier, volumeName IFSFileName) FSVolume {
	instance := getFSVolumeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithVolumeID:volumeName:"), volumeID, volumeName)
	return FSVolumeFromID(rv)
}

// Creates a volume with the given identifier and name.
//
// volumeID: An [FSVolumeIdentifier] to uniquely identify the volume. For a network file
// system that supports multiple authenticated users, disambiguate the users
// by using qualifying data in the identifier.
//
// volumeName: A name for the volume.
//
// See: https://developer.apple.com/documentation/FSKit/FSVolume/init(volumeID:volumeName:)
func (v FSVolume) InitWithVolumeIDVolumeName(volumeID IFSVolumeIdentifier, volumeName IFSFileName) FSVolume {
	rv := objc.Send[FSVolume](v.ID, objc.Sel("initWithVolumeID:volumeName:"), volumeID, volumeName)
	return rv
}

// An identifier that uniquely identifies the volume.
//
// See: https://developer.apple.com/documentation/FSKit/FSVolume/volumeID
func (v FSVolume) VolumeID() IFSVolumeIdentifier {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("volumeID"))
	return FSVolumeIdentifierFromID(objc.ID(rv))
}

// The name of the volume.
//
// See: https://developer.apple.com/documentation/FSKit/FSVolume/name
func (v FSVolume) Name() IFSFileName {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("name"))
	return FSFileNameFromID(objc.ID(rv))
}
func (v FSVolume) SetName(value IFSFileName) {
	objc.Send[struct{}](v.ID, objc.Sel("setName:"), value)
}
