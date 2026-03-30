// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [WKWebExtensionDataRecord] class.
var (
	_WKWebExtensionDataRecordClass     WKWebExtensionDataRecordClass
	_WKWebExtensionDataRecordClassOnce sync.Once
)

func getWKWebExtensionDataRecordClass() WKWebExtensionDataRecordClass {
	_WKWebExtensionDataRecordClassOnce.Do(func() {
		_WKWebExtensionDataRecordClass = WKWebExtensionDataRecordClass{class: objc.GetClass("WKWebExtensionDataRecord")}
	})
	return _WKWebExtensionDataRecordClass
}

// GetWKWebExtensionDataRecordClass returns the class object for WKWebExtensionDataRecord.
func GetWKWebExtensionDataRecordClass() WKWebExtensionDataRecordClass {
	return getWKWebExtensionDataRecordClass()
}

type WKWebExtensionDataRecordClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (wc WKWebExtensionDataRecordClass) Class() objc.Class {
	return wc.class
}

// Alloc allocates memory for a new instance of the class.
func (wc WKWebExtensionDataRecordClass) Alloc() WKWebExtensionDataRecord {
	rv := objc.Send[WKWebExtensionDataRecord](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

// An object that represents a record of stored data for a specific web
// extension context.
//
// # Overview
//
// Contains properties and methods to query the data types and sizes.
//
// # Instance Properties
//
//   - [WKWebExtensionDataRecord.ContainedDataTypes]: The set of data types contained in this data record.
//   - [WKWebExtensionDataRecord.DisplayName]: The display name for the web extension to which this data record belongs.
//   - [WKWebExtensionDataRecord.Errors]: An array of errors that may have occurred when either calculating or deleting storage.
//   - [WKWebExtensionDataRecord.TotalSizeInBytes]: The total size in bytes of all data types contained in this data record.
//   - [WKWebExtensionDataRecord.UniqueIdentifier]: Unique identifier for the web extension context to which this data record belongs.
//
// # Instance Methods
//
//   - [WKWebExtensionDataRecord.SizeInBytesOfTypes]: Retrieves the size in bytes of the specific data types in this data record.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtension/DataRecord
type WKWebExtensionDataRecord struct {
	objectivec.Object
}

// WKWebExtensionDataRecordFromID constructs a [WKWebExtensionDataRecord] from an objc.ID.
//
// An object that represents a record of stored data for a specific web
// extension context.
func WKWebExtensionDataRecordFromID(id objc.ID) WKWebExtensionDataRecord {
	return WKWebExtensionDataRecord{objectivec.Object{ID: id}}
}

// NOTE: WKWebExtensionDataRecord adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [WKWebExtensionDataRecord] class.
//
// # Instance Properties
//
//   - [IWKWebExtensionDataRecord.ContainedDataTypes]: The set of data types contained in this data record.
//   - [IWKWebExtensionDataRecord.DisplayName]: The display name for the web extension to which this data record belongs.
//   - [IWKWebExtensionDataRecord.Errors]: An array of errors that may have occurred when either calculating or deleting storage.
//   - [IWKWebExtensionDataRecord.TotalSizeInBytes]: The total size in bytes of all data types contained in this data record.
//   - [IWKWebExtensionDataRecord.UniqueIdentifier]: Unique identifier for the web extension context to which this data record belongs.
//
// # Instance Methods
//
//   - [IWKWebExtensionDataRecord.SizeInBytesOfTypes]: Retrieves the size in bytes of the specific data types in this data record.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtension/DataRecord
type IWKWebExtensionDataRecord interface {
	objectivec.IObject

	// Topic: Instance Properties

	// The set of data types contained in this data record.
	ContainedDataTypes() foundation.INSSet
	// The display name for the web extension to which this data record belongs.
	DisplayName() string
	// An array of errors that may have occurred when either calculating or deleting storage.
	Errors() []foundation.NSError
	// The total size in bytes of all data types contained in this data record.
	TotalSizeInBytes() uint
	// Unique identifier for the web extension context to which this data record belongs.
	UniqueIdentifier() string

	// Topic: Instance Methods

	// Retrieves the size in bytes of the specific data types in this data record.
	SizeInBytesOfTypes(dataTypes foundation.INSSet) uint
}

// Init initializes the instance.
func (w WKWebExtensionDataRecord) Init() WKWebExtensionDataRecord {
	rv := objc.Send[WKWebExtensionDataRecord](w.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (w WKWebExtensionDataRecord) Autorelease() WKWebExtensionDataRecord {
	rv := objc.Send[WKWebExtensionDataRecord](w.ID, objc.Sel("autorelease"))
	return rv
}

// NewWKWebExtensionDataRecord creates a new WKWebExtensionDataRecord instance.
func NewWKWebExtensionDataRecord() WKWebExtensionDataRecord {
	class := getWKWebExtensionDataRecordClass()
	rv := objc.Send[WKWebExtensionDataRecord](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Retrieves the size in bytes of the specific data types in this data record.
//
// dataTypes: The set of data types to measure the size for.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtension/DataRecord/sizeInBytes(ofTypes:)
func (w WKWebExtensionDataRecord) SizeInBytesOfTypes(dataTypes foundation.INSSet) uint {
	rv := objc.Send[uint](w.ID, objc.Sel("sizeInBytesOfTypes:"), dataTypes)
	return rv
}

// The set of data types contained in this data record.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtension/DataRecord/containedDataTypes
func (w WKWebExtensionDataRecord) ContainedDataTypes() foundation.INSSet {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("containedDataTypes"))
	return foundation.NSSetFromID(objc.ID(rv))
}

// The display name for the web extension to which this data record belongs.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtension/DataRecord/displayName
func (w WKWebExtensionDataRecord) DisplayName() string {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("displayName"))
	return foundation.NSStringFromID(rv).String()
}

// An array of errors that may have occurred when either calculating or
// deleting storage.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtension/DataRecord/errors
func (w WKWebExtensionDataRecord) Errors() []foundation.NSError {
	rv := objc.Send[[]objc.ID](w.ID, objc.Sel("errors"))
	return objc.ConvertSlice(rv, func(id objc.ID) foundation.NSError {
		return foundation.NSErrorFromID(id)
	})
}

// The total size in bytes of all data types contained in this data record.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtension/DataRecord/totalSizeInBytes
func (w WKWebExtensionDataRecord) TotalSizeInBytes() uint {
	rv := objc.Send[uint](w.ID, objc.Sel("totalSizeInBytes"))
	return rv
}

// Unique identifier for the web extension context to which this data record
// belongs.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtension/DataRecord/uniqueIdentifier
func (w WKWebExtensionDataRecord) UniqueIdentifier() string {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("uniqueIdentifier"))
	return foundation.NSStringFromID(rv).String()
}
