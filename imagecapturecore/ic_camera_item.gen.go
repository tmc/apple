// Code generated from Apple documentation for ImageCaptureCore. DO NOT EDIT.

package imagecapturecore

import (
	"sync"

	"github.com/tmc/apple/coregraphics"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [ICCameraItem] class.
var (
	_ICCameraItemClass     ICCameraItemClass
	_ICCameraItemClassOnce sync.Once
)

func getICCameraItemClass() ICCameraItemClass {
	_ICCameraItemClassOnce.Do(func() {
		_ICCameraItemClass = ICCameraItemClass{class: objc.GetClass("ICCameraItem")}
	})
	return _ICCameraItemClass
}

// GetICCameraItemClass returns the class object for ICCameraItem.
func GetICCameraItemClass() ICCameraItemClass {
	return getICCameraItemClass()
}

type ICCameraItemClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ic ICCameraItemClass) Class() objc.Class {
	return ic.class
}

// Alloc allocates memory for a new instance of the class.
func (ic ICCameraItemClass) Alloc() ICCameraItem {
	rv := objc.Send[ICCameraItem](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// An abstract class that represents a camera item.
//
// # Overview
//
// The ImageCaptureCore framework defines two concrete subclasses of camera
// items: [ICCameraFolder] and [ICCameraFile].
//
// # Inspecting an Item’s Name and Type
//
//   - [ICCameraItem.UTI]: The item’s uniform type identifier (UTI) string.
//   - [ICCameraItem.Name]: The item’s name.
//   - [ICCameraItem.PtpObjectHandle]: The item’s [PTP] object handle value, if the camera uses the [PTP] protocol.
//   - [ICCameraItem.IsRaw]: A Boolean value indicating whether the item is a raw image file.
//
// # Determining an Item’s Change Dates
//
//   - [ICCameraItem.CreationDate]: The item’s creation date, usually the same as its [EXIF] creation date.
//   - [ICCameraItem.ModificationDate]: The item’s modification date, usually the same as its [EXIF] modification date.
//   - [ICCameraItem.AddedAfterContentCatalogCompleted]: A Boolean value indicating whether the item was captured on the camera after the camera’s content had been fully enumerated.
//
// # Locating an Item
//
//   - [ICCameraItem.Device]: The item’s parent device.
//   - [ICCameraItem.FileSystemPath]: The item’s file system path on a camera using the mass storage transport type.
//   - [ICCameraItem.ParentFolder]: This item’s parent folder.
//   - [ICCameraItem.IsInTemporaryStore]: A Boolean value that indicates whether this item is in a temporary store.
//
// # Requesting Metadata
//
//   - [ICCameraItem.RequestMetadata]: Requests metadata for the item.
//   - [ICCameraItem.Metadata]: The item’s metadata.
//   - [ICCameraItem.FlushMetadataCache]: Deletes the item’s cached metadata.
//
// # Requesting Thumbnails
//
//   - [ICCameraItem.RequestThumbnail]: Requests a thumbnail for the item.
//   - [ICCameraItem.Thumbnail]: The item’s thumbnail.
//   - [ICCameraItem.FlushThumbnailCache]: Deletes the item’s cached thumbnail.
//
// # Accessing a Protected Item
//
//   - [ICCameraItem.IsLocked]: A Boolean value that indicates whether the storage card in the camera is locked.
//
// # Storing Information
//
//   - [ICCameraItem.UserData]: A mutable dictionary to store arbitrary key-value pairs associated with a camera item.
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICCameraItem
type ICCameraItem struct {
	objectivec.Object
}

// ICCameraItemFromID constructs a [ICCameraItem] from an objc.ID.
//
// An abstract class that represents a camera item.
func ICCameraItemFromID(id objc.ID) ICCameraItem {
	return ICCameraItem{objectivec.Object{ID: id}}
}

// NOTE: ICCameraItem adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [ICCameraItem] class.
//
// # Inspecting an Item’s Name and Type
//
//   - [IICCameraItem.UTI]: The item’s uniform type identifier (UTI) string.
//   - [IICCameraItem.Name]: The item’s name.
//   - [IICCameraItem.PtpObjectHandle]: The item’s [PTP] object handle value, if the camera uses the [PTP] protocol.
//   - [IICCameraItem.IsRaw]: A Boolean value indicating whether the item is a raw image file.
//
// # Determining an Item’s Change Dates
//
//   - [IICCameraItem.CreationDate]: The item’s creation date, usually the same as its [EXIF] creation date.
//   - [IICCameraItem.ModificationDate]: The item’s modification date, usually the same as its [EXIF] modification date.
//   - [IICCameraItem.AddedAfterContentCatalogCompleted]: A Boolean value indicating whether the item was captured on the camera after the camera’s content had been fully enumerated.
//
// # Locating an Item
//
//   - [IICCameraItem.Device]: The item’s parent device.
//   - [IICCameraItem.FileSystemPath]: The item’s file system path on a camera using the mass storage transport type.
//   - [IICCameraItem.ParentFolder]: This item’s parent folder.
//   - [IICCameraItem.IsInTemporaryStore]: A Boolean value that indicates whether this item is in a temporary store.
//
// # Requesting Metadata
//
//   - [IICCameraItem.RequestMetadata]: Requests metadata for the item.
//   - [IICCameraItem.Metadata]: The item’s metadata.
//   - [IICCameraItem.FlushMetadataCache]: Deletes the item’s cached metadata.
//
// # Requesting Thumbnails
//
//   - [IICCameraItem.RequestThumbnail]: Requests a thumbnail for the item.
//   - [IICCameraItem.Thumbnail]: The item’s thumbnail.
//   - [IICCameraItem.FlushThumbnailCache]: Deletes the item’s cached thumbnail.
//
// # Accessing a Protected Item
//
//   - [IICCameraItem.IsLocked]: A Boolean value that indicates whether the storage card in the camera is locked.
//
// # Storing Information
//
//   - [IICCameraItem.UserData]: A mutable dictionary to store arbitrary key-value pairs associated with a camera item.
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICCameraItem
type IICCameraItem interface {
	objectivec.IObject

	// Topic: Inspecting an Item’s Name and Type

	// The item’s uniform type identifier (UTI) string.
	UTI() string
	// The item’s name.
	Name() string
	// The item’s [PTP] object handle value, if the camera uses the [PTP] protocol.
	PtpObjectHandle() uint32
	// A Boolean value indicating whether the item is a raw image file.
	IsRaw() bool

	// Topic: Determining an Item’s Change Dates

	// The item’s creation date, usually the same as its [EXIF] creation date.
	CreationDate() foundation.NSDate
	// The item’s modification date, usually the same as its [EXIF] modification date.
	ModificationDate() foundation.NSDate
	// A Boolean value indicating whether the item was captured on the camera after the camera’s content had been fully enumerated.
	AddedAfterContentCatalogCompleted() bool

	// Topic: Locating an Item

	// The item’s parent device.
	Device() IICCameraDevice
	// The item’s file system path on a camera using the mass storage transport type.
	FileSystemPath() string
	// This item’s parent folder.
	ParentFolder() IICCameraFolder
	// A Boolean value that indicates whether this item is in a temporary store.
	IsInTemporaryStore() bool

	// Topic: Requesting Metadata

	// Requests metadata for the item.
	RequestMetadata()
	// The item’s metadata.
	Metadata() foundation.INSDictionary
	// Deletes the item’s cached metadata.
	FlushMetadataCache()

	// Topic: Requesting Thumbnails

	// Requests a thumbnail for the item.
	RequestThumbnail()
	// The item’s thumbnail.
	Thumbnail() coregraphics.CGImageRef
	// Deletes the item’s cached thumbnail.
	FlushThumbnailCache()

	// Topic: Accessing a Protected Item

	// A Boolean value that indicates whether the storage card in the camera is locked.
	IsLocked() bool

	// Topic: Storing Information

	// A mutable dictionary to store arbitrary key-value pairs associated with a camera item.
	UserData() foundation.INSDictionary
}

// Init initializes the instance.
func (c ICCameraItem) Init() ICCameraItem {
	rv := objc.Send[ICCameraItem](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c ICCameraItem) Autorelease() ICCameraItem {
	rv := objc.Send[ICCameraItem](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewICCameraItem creates a new ICCameraItem instance.
func NewICCameraItem() ICCameraItem {
	class := getICCameraItemClass()
	rv := objc.Send[ICCameraItem](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Requests metadata for the item.
//
// # Discussion
//
// If metadata for the item is not readily available, accessing this property
// requests metadata from the camera, then notifies the delegate by calling
// [CameraDeviceDidReceiveMetadataForItemError].
//
// Execution of the delegate callback occurs on the main thread.
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICCameraItem/requestMetadata()
func (c ICCameraItem) RequestMetadata() {
	objc.Send[objc.ID](c.ID, objc.Sel("requestMetadata"))
}

// Deletes the item’s cached metadata.
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICCameraItem/flushMetadataCache()
func (c ICCameraItem) FlushMetadataCache() {
	objc.Send[objc.ID](c.ID, objc.Sel("flushMetadataCache"))
}

// Requests a thumbnail for the item.
//
// # Discussion
//
// If a thumbnail is not readily available, accessing this property will send
// a message to the device requesting a thumbnail for the file. The delegate
// of the device will be notified via method
// [CameraDeviceDidReceiveThumbnailForItemError], if this method is
// implemented by the delegate. Execution of the delegate callback will occur
// on the main thread.
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICCameraItem/requestThumbnail()
func (c ICCameraItem) RequestThumbnail() {
	objc.Send[objc.ID](c.ID, objc.Sel("requestThumbnail"))
}

// Deletes the item’s cached thumbnail.
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICCameraItem/flushThumbnailCache()
func (c ICCameraItem) FlushThumbnailCache() {
	objc.Send[objc.ID](c.ID, objc.Sel("flushThumbnailCache"))
}

// The item’s uniform type identifier (UTI) string.
//
// # Discussion
//
// The [UTI] options are `kUTTypeFolder`, `kUTTypeImage`, `kUTTypeMovie`,
// `kUTTypeAudio`, or `kUTTypeData`.
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICCameraItem/uti
func (c ICCameraItem) UTI() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("UTI"))
	return foundation.NSStringFromID(rv).String()
}

// The item’s name.
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICCameraItem/name
func (c ICCameraItem) Name() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("name"))
	return foundation.NSStringFromID(rv).String()
}

// The item’s [PTP] object handle value, if the camera uses the [PTP]
// protocol.
//
// # Discussion
//
// The value of this property is set to `0` if the camera does not use PTP
// protocol.
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICCameraItem/ptpObjectHandle
func (c ICCameraItem) PtpObjectHandle() uint32 {
	rv := objc.Send[uint32](c.ID, objc.Sel("ptpObjectHandle"))
	return rv
}

// A Boolean value indicating whether the item is a raw image file.
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICCameraItem/isRaw
func (c ICCameraItem) IsRaw() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isRaw"))
	return rv
}

// The item’s creation date, usually the same as its [EXIF] creation date.
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICCameraItem/creationDate
func (c ICCameraItem) CreationDate() foundation.NSDate {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("creationDate"))
	return foundation.NSDateFromID(objc.ID(rv))
}

// The item’s modification date, usually the same as its [EXIF] modification
// date.
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICCameraItem/modificationDate
func (c ICCameraItem) ModificationDate() foundation.NSDate {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("modificationDate"))
	return foundation.NSDateFromID(objc.ID(rv))
}

// A Boolean value indicating whether the item was captured on the camera
// after the camera’s content had been fully enumerated.
//
// # Discussion
//
// This value does not apply to files added as a result of adding a new store
// to the camera.
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICCameraItem/wasAddedAfterContentCatalogCompleted
func (c ICCameraItem) AddedAfterContentCatalogCompleted() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("wasAddedAfterContentCatalogCompleted"))
	return rv
}

// The item’s parent device.
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICCameraItem/device
func (c ICCameraItem) Device() IICCameraDevice {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("device"))
	return ICCameraDeviceFromID(objc.ID(rv))
}

// The item’s file system path on a camera using the mass storage transport
// type.
//
// # Discussion
//
// This property is set for cameras whose [ICDevice.TransportType] is
// [transportTypeMassStorage].
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICCameraItem/fileSystemPath
//
// [transportTypeMassStorage]: https://developer.apple.com/documentation/ImageCaptureCore/ICDeviceTransport/transportTypeMassStorage
func (c ICCameraItem) FileSystemPath() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("fileSystemPath"))
	return foundation.NSStringFromID(rv).String()
}

// This item’s parent folder.
//
// # Discussion
//
// The value of this property on the root folder is nil.
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICCameraItem/parentFolder
func (c ICCameraItem) ParentFolder() IICCameraFolder {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("parentFolder"))
	return ICCameraFolderFromID(objc.ID(rv))
}

// A Boolean value that indicates whether this item is in a temporary store.
//
// # Discussion
//
// A device may use a temporary store when it captures images while tethered
// to a computer.
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICCameraItem/isInTemporaryStore
func (c ICCameraItem) IsInTemporaryStore() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isInTemporaryStore"))
	return rv
}

// The item’s metadata.
//
// # Discussion
//
// The value of this property is `nil` unless a [ICCameraItem.RequestMetadata]
// message is sent to this object.
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICCameraItem/metadata
func (c ICCameraItem) Metadata() foundation.INSDictionary {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("metadata"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}

// The item’s thumbnail.
//
// # Discussion
//
// The value of this property is `nil` until you call
// [ICCameraItem.RequestThumbnail].
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICCameraItem/thumbnail
func (c ICCameraItem) Thumbnail() coregraphics.CGImageRef {
	rv := objc.Send[coregraphics.CGImageRef](c.ID, objc.Sel("thumbnail"))
	return coregraphics.CGImageRef(rv)
}

// A Boolean value that indicates whether the storage card in the camera is
// locked.
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICCameraItem/isLocked
func (c ICCameraItem) IsLocked() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isLocked"))
	return rv
}

// A mutable dictionary to store arbitrary key-value pairs associated with a
// camera item.
//
// # Discussion
//
// View objects can bind to this object to store “house-keeping”
// information.
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICCameraItem/userData
func (c ICCameraItem) UserData() foundation.INSDictionary {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("userData"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
