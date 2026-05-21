// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CKAsset] class.
var (
	_CKAssetClass     CKAssetClass
	_CKAssetClassOnce sync.Once
)

func getCKAssetClass() CKAssetClass {
	_CKAssetClassOnce.Do(func() {
		_CKAssetClass = CKAssetClass{class: objc.GetClass("CKAsset")}
	})
	return _CKAssetClass
}

// GetCKAssetClass returns the class object for CKAsset.
func GetCKAssetClass() CKAssetClass {
	return getCKAssetClass()
}

type CKAssetClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CKAssetClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CKAssetClass) Alloc() CKAsset {
	rv := objc.Send[CKAsset](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// An external file that belongs to a record.
//
// # Overview
//
// Use assets to incorporate external files into your app’s records, such as
// photos, videos, and binary files. Alternatively, use assets when a
// field’s value is more than a few kilobytes in size. To associate an
// instance of [CKAsset] with a record, assign it to one of its fields.
//
// CloudKit stores an asset’s data separately from a record that references
// it, but maintains an association with that record. When you save a record
// that has an asset, CloudKit saves both the record and the asset to the
// server. Similarly, when you fetch the record, the server returns the record
// and the asset.
//
// When you fetch a record that contains an asset, CloudKit stores the
// asset’s data in a staging area accessible to your app. Use the asset’s
// [CKAsset.FileURL] property to access its staged location. The system
// regularly deletes files in the staging area to reclaim disk space. To avoid
// this behavior, move the data into your app’s container as soon as you
// fetch it.
//
// If you don’t require the asset when retrieving records, use the
// operation’s `desiredKeys` property to exclude the field. For more
// information, see [CKFetchRecordsOperation], [CKQueryOperation], and
// [CKFetchRecordZoneChangesOperation].
//
// If you no longer require an asset that’s on the server, you don’t
// delete it. Instead, orphan the asset by setting any fields that contain the
// asset to `nil` and then saving the record. CloudKit periodically deletes
// orphaned assets from the server.
//
// # Creating an Asset
//
//   - [CKAsset.InitWithFileURL]: Creates an asset that references a file.
//
// # Getting the URL of the Asset
//
//   - [CKAsset.FileURL]: The URL for accessing the asset.
//
// See: https://developer.apple.com/documentation/CloudKit/CKAsset
type CKAsset struct {
	objectivec.Object
}

// CKAssetFromID constructs a [CKAsset] from an objc.ID.
//
// An external file that belongs to a record.
func CKAssetFromID(id objc.ID) CKAsset {
	return CKAsset{objectivec.Object{ID: id}}
}

// NOTE: CKAsset adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CKAsset] class.
//
// # Creating an Asset
//
//   - [ICKAsset.InitWithFileURL]: Creates an asset that references a file.
//
// # Getting the URL of the Asset
//
//   - [ICKAsset.FileURL]: The URL for accessing the asset.
//
// See: https://developer.apple.com/documentation/CloudKit/CKAsset
type ICKAsset interface {
	objectivec.IObject

	// Topic: Creating an Asset

	// Creates an asset that references a file.
	InitWithFileURL(fileURL foundation.NSURL) CKAsset

	// Topic: Getting the URL of the Asset

	// The URL for accessing the asset.
	FileURL() foundation.NSURL
}

// Init initializes the instance.
func (c CKAsset) Init() CKAsset {
	rv := objc.Send[CKAsset](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CKAsset) Autorelease() CKAsset {
	rv := objc.Send[CKAsset](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKAsset creates a new CKAsset instance.
func NewCKAsset() CKAsset {
	class := getCKAssetClass()
	rv := objc.Send[CKAsset](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates an asset that references a file.
//
// fileURL: The URL of the file that you want to store in CloudKit. The URL must be a
// file URL, and must not be `nil`.
//
// # Return Value
//
// An asset object that represents the specified file, or `nil` if the system
// can’t create the asset.
//
// # Discussion
//
// Use this method to initialize new file-based assets that you want to
// transfer to iCloud. After saving an asset to the server, CloudKit doesn’t
// delete the file at the specified URL. If you no longer need the file, you
// must delete it yourself. When you subsequently download a record that
// contains an asset, CloudKit downloads its own copy of the asset data to the
// local device and provides you with a URL to that file.
//
// You can assign only one record to the asset that this method returns. If
// you want multiple records to point to the same file, you must create
// separate assets for each one.
//
// See: https://developer.apple.com/documentation/CloudKit/CKAsset/init(fileURL:)
func NewCKAssetWithFileURL(fileURL foundation.NSURL) CKAsset {
	instance := getCKAssetClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFileURL:"), fileURL)
	return CKAssetFromID(rv)
}

// Creates an asset that references a file.
//
// fileURL: The URL of the file that you want to store in CloudKit. The URL must be a
// file URL, and must not be `nil`.
//
// # Return Value
//
// An asset object that represents the specified file, or `nil` if the system
// can’t create the asset.
//
// # Discussion
//
// Use this method to initialize new file-based assets that you want to
// transfer to iCloud. After saving an asset to the server, CloudKit doesn’t
// delete the file at the specified URL. If you no longer need the file, you
// must delete it yourself. When you subsequently download a record that
// contains an asset, CloudKit downloads its own copy of the asset data to the
// local device and provides you with a URL to that file.
//
// You can assign only one record to the asset that this method returns. If
// you want multiple records to point to the same file, you must create
// separate assets for each one.
//
// See: https://developer.apple.com/documentation/CloudKit/CKAsset/init(fileURL:)
func (c CKAsset) InitWithFileURL(fileURL foundation.NSURL) CKAsset {
	rv := objc.Send[CKAsset](c.ID, objc.Sel("initWithFileURL:"), fileURL)
	return rv
}

// The URL for accessing the asset.
//
// # Discussion
//
// After you create an asset, use the URL in this property to access the
// asset’s contents. The URL in this property is different from the one you
// specify when creating the asset.
//
// See: https://developer.apple.com/documentation/CloudKit/CKAsset/fileURL
func (c CKAsset) FileURL() foundation.NSURL {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("fileURL"))
	return foundation.NSURLFromID(objc.ID(rv))
}
