// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [WKWebsiteDataRecord] class.
var (
	_WKWebsiteDataRecordClass     WKWebsiteDataRecordClass
	_WKWebsiteDataRecordClassOnce sync.Once
)

func getWKWebsiteDataRecordClass() WKWebsiteDataRecordClass {
	_WKWebsiteDataRecordClassOnce.Do(func() {
		_WKWebsiteDataRecordClass = WKWebsiteDataRecordClass{class: objc.GetClass("WKWebsiteDataRecord")}
	})
	return _WKWebsiteDataRecordClass
}

// GetWKWebsiteDataRecordClass returns the class object for WKWebsiteDataRecord.
func GetWKWebsiteDataRecordClass() WKWebsiteDataRecordClass {
	return getWKWebsiteDataRecordClass()
}

type WKWebsiteDataRecordClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (wc WKWebsiteDataRecordClass) Class() objc.Class {
	return wc.class
}

// Alloc allocates memory for a new instance of the class.
func (wc WKWebsiteDataRecordClass) Alloc() WKWebsiteDataRecord {
	rv := objc.Send[WKWebsiteDataRecord](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

// A record of the data that a particular website stores persistently.
//
// # Overview
//
// Use [WKWebsiteDataRecord] objects to discover the types of information that
// a website stores. Records identify the data types a website stores, but
// don’t identify the actual data. You might use this information to help
// the user manage website data. For example, Safari provides a way for users
// to view and remove website data. The domain name of each record contains
// the website’s domain name and suffix.
//
// You don’t create [WKWebsiteDataRecord] objects directly. WebKit creates
// these records and stores them in the web view’s data store. Use the
// [FetchDataRecordsOfTypesCompletionHandler] of that data store to retrieve
// the current record objects. You also use that object to remove unwanted
// records.
//
// # Getting the Record Information
//
//   - [WKWebsiteDataRecord.DisplayName]: The display name for the data record.
//
// # Getting the Data Type
//
//   - [WKWebsiteDataRecord.DataTypes]: The types of data associated with the record.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebsiteDataRecord
type WKWebsiteDataRecord struct {
	objectivec.Object
}

// WKWebsiteDataRecordFromID constructs a [WKWebsiteDataRecord] from an objc.ID.
//
// A record of the data that a particular website stores persistently.
func WKWebsiteDataRecordFromID(id objc.ID) WKWebsiteDataRecord {
	return WKWebsiteDataRecord{objectivec.Object{ID: id}}
}

// NOTE: WKWebsiteDataRecord adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [WKWebsiteDataRecord] class.
//
// # Getting the Record Information
//
//   - [IWKWebsiteDataRecord.DisplayName]: The display name for the data record.
//
// # Getting the Data Type
//
//   - [IWKWebsiteDataRecord.DataTypes]: The types of data associated with the record.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebsiteDataRecord
type IWKWebsiteDataRecord interface {
	objectivec.IObject

	// Topic: Getting the Record Information

	// The display name for the data record.
	DisplayName() string

	// Topic: Getting the Data Type

	// The types of data associated with the record.
	DataTypes() foundation.INSSet

	// The local files WebKit can access when loading content.
	ReadAccessURL() foundation.NSString
}

// Init initializes the instance.
func (w WKWebsiteDataRecord) Init() WKWebsiteDataRecord {
	rv := objc.Send[WKWebsiteDataRecord](w.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (w WKWebsiteDataRecord) Autorelease() WKWebsiteDataRecord {
	rv := objc.Send[WKWebsiteDataRecord](w.ID, objc.Sel("autorelease"))
	return rv
}

// NewWKWebsiteDataRecord creates a new WKWebsiteDataRecord instance.
func NewWKWebsiteDataRecord() WKWebsiteDataRecord {
	class := getWKWebsiteDataRecordClass()
	rv := objc.Send[WKWebsiteDataRecord](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The display name for the data record.
//
// # Discussion
//
// This property contains identifying information that you to display to
// users. Typically, the value is the website’s domain name and suffix taken
// from the host information in the resource’s security origin object. For
// more information, see [WKSecurityOrigin].
//
// See: https://developer.apple.com/documentation/WebKit/WKWebsiteDataRecord/displayName
func (w WKWebsiteDataRecord) DisplayName() string {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("displayName"))
	return foundation.NSStringFromID(rv).String()
}

// The types of data associated with the record.
//
// # Discussion
//
// Each record contains the set of types that the associated website stores.
// For a list of possible types, see [Data Store Record Types].
//
// See: https://developer.apple.com/documentation/WebKit/WKWebsiteDataRecord/dataTypes
//
// [Data Store Record Types]: https://developer.apple.com/documentation/WebKit/data-store-record-types
func (w WKWebsiteDataRecord) DataTypes() foundation.INSSet {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("dataTypes"))
	return foundation.NSSetFromID(objc.ID(rv))
}

// The local files WebKit can access when loading content.
//
// See: https://developer.apple.com/documentation/Foundation/NSAttributedString/DocumentReadingOptionKey/readAccessURL
func (w WKWebsiteDataRecord) ReadAccessURL() foundation.NSString {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("readAccessURL"))
	return foundation.NSStringFromID(objc.ID(rv))
}
