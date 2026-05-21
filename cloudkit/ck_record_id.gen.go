// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CKRecordID] class.
var (
	_CKRecordIDClass     CKRecordIDClass
	_CKRecordIDClassOnce sync.Once
)

func getCKRecordIDClass() CKRecordIDClass {
	_CKRecordIDClassOnce.Do(func() {
		_CKRecordIDClass = CKRecordIDClass{class: objc.GetClass("CKRecordID")}
	})
	return _CKRecordIDClass
}

// GetCKRecordIDClass returns the class object for CKRecordID.
func GetCKRecordIDClass() CKRecordIDClass {
	return getCKRecordIDClass()
}

type CKRecordIDClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CKRecordIDClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CKRecordIDClass) Alloc() CKRecordID {
	rv := objc.Send[CKRecordID](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// An object that uniquely identifies a record in a database.
//
// # Overview
//
// A record ID object consists of a name string and a zone ID. The name string
// is an ASCII string that doesn’t exceed 255 characters in length. When you
// create a record without specifying a record ID, the ID name string derives
// from a UUID and is, therefore, unique. When creating your own record ID
// objects, you can use names that have more meaning to your app or to the
// user, as long as each name is unique within the specified zone. For
// example, you might use a document name for the name string.
//
// CloudKit uniques records by recordID within a specified database, but you
// can reuse record IDs in different databases. Each container has a public
// and a private database, and the private database is different for each
// unique user. This configuration provides for the reusing of record IDs in
// each user’s private database, but ensures that only one record uses a
// specific record ID in the public database.
//
// CloudKit generally creates record IDs when it first saves a new record, but
// you might manually instantiate instances of [CKRecordID] in specific
// situations. For example, you must create an instance when saving a record
// in a zone other than the default zone. You also instantiate instances of
// [CKRecordID] when retrieving specific records from a database.
//
// Don’t subclass [CKRecordID].
//
// # Interacting with Record IDs
//
// After you create a [CKRecordID] object, interactions with that object
// typically involve creating a new record or retrieving an existing record
// from a database.
//
// You might also use record IDs when you can’t use a [CKReference] object
// to refer to a record. References are only valid within a single zone of a
// database. To refer to objects outside of the current zone or database, save
// the strings in the record’s [CKRecordID] and [CKRecordZoneID] objects.
// When you want to retrieve the record later, use those strings to recreate
// the record and zone ID objects so that you can fetch the record.
//
// # Creating Record IDs for New Records
//
// To assign a custom record ID to a new record, you must create the
// [CKRecordID] object first. You need to know the intended name and zone
// information for that record, which might also require creating a
// [CKRecordZoneID] object. After creating the record ID object, initialize
// your new record using its [init(recordType:recordID:)] method.
//
// # Using Record IDs to Fetch Records
//
// Use a record ID to fetch the corresponding [CKRecord] object from a
// database quickly. You perform the fetch operation using a
// [CKFetchRecordsOperation] object or the
// [CKDatabase.FetchRecordWithIDCompletionHandler] method of the [CKDatabase]
// class. In both cases, CloudKit returns the record asynchronously using the
// handler you provide.
//
// # Creating a Record ID
//
//   - [CKRecordID.InitWithRecordName]: Creates a new record ID with the specified name in the default zone.
//
// # Getting the Record ID’s Name
//
//   - [CKRecordID.RecordName]: The unique name of the record.
//
// # Getting the Record ID’s Zone
//
//   - [CKRecordID.ZoneID]: The ID of the zone that contains the record.
//
// See: https://developer.apple.com/documentation/CloudKit/CKRecord/ID
//
// [init(recordType:recordID:)]: https://developer.apple.com/documentation/CloudKit/CKRecord/init(recordType:recordID:)
type CKRecordID struct {
	objectivec.Object
}

// CKRecordIDFromID constructs a [CKRecordID] from an objc.ID.
//
// An object that uniquely identifies a record in a database.
func CKRecordIDFromID(id objc.ID) CKRecordID {
	return CKRecordID{objectivec.Object{ID: id}}
}

// NOTE: CKRecordID adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CKRecordID] class.
//
// # Creating a Record ID
//
//   - [ICKRecordID.InitWithRecordName]: Creates a new record ID with the specified name in the default zone.
//
// # Getting the Record ID’s Name
//
//   - [ICKRecordID.RecordName]: The unique name of the record.
//
// # Getting the Record ID’s Zone
//
//   - [ICKRecordID.ZoneID]: The ID of the zone that contains the record.
//
// See: https://developer.apple.com/documentation/CloudKit/CKRecord/ID
type ICKRecordID interface {
	objectivec.IObject

	// Topic: Creating a Record ID

	// Creates a new record ID with the specified name in the default zone.
	InitWithRecordName(recordName string) CKRecordID

	// Topic: Getting the Record ID’s Name

	// The unique name of the record.
	RecordName() string

	// Topic: Getting the Record ID’s Zone

	// The ID of the zone that contains the record.
	ZoneID() ICKRecordZoneID

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (c CKRecordID) Init() CKRecordID {
	rv := objc.Send[CKRecordID](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CKRecordID) Autorelease() CKRecordID {
	rv := objc.Send[CKRecordID](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKRecordID creates a new CKRecordID instance.
func NewCKRecordID() CKRecordID {
	class := getCKRecordIDClass()
	rv := objc.Send[CKRecordID](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a new record ID with the specified name in the default zone.
//
// recordName: The name that identifies the record. The string must contain only ASCII
// characters, must not exceed 255 characters, and must not start with an
// underscore. If you specify an empty string for this parameter, the method
// throws an exception.
//
// # Return Value
//
// An initialized record ID object.
//
// # Discussion
//
// Use this method when you’re creating or searching for records in the
// default zone.
//
// See: https://developer.apple.com/documentation/CloudKit/CKRecord/ID/init(recordName:)
func NewCKRecordIDWithRecordName(recordName string) CKRecordID {
	instance := getCKRecordIDClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithRecordName:"), objc.String(recordName))
	return CKRecordIDFromID(rv)
}

// Creates a new record ID with the specified name in the default zone.
//
// recordName: The name that identifies the record. The string must contain only ASCII
// characters, must not exceed 255 characters, and must not start with an
// underscore. If you specify an empty string for this parameter, the method
// throws an exception.
//
// # Return Value
//
// An initialized record ID object.
//
// # Discussion
//
// Use this method when you’re creating or searching for records in the
// default zone.
//
// See: https://developer.apple.com/documentation/CloudKit/CKRecord/ID/init(recordName:)
func (c CKRecordID) InitWithRecordName(recordName string) CKRecordID {
	rv := objc.Send[CKRecordID](c.ID, objc.Sel("initWithRecordName:"), objc.String(recordName))
	return rv
}
func (c CKRecordID) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](c.ID, objc.Sel("encodeWithCoder:"), coder)
}

// The unique name of the record.
//
// # Discussion
//
// For share records that manage a shared record zone, this property’s value
// is always [CKRecordNameZoneWideShare].
//
// See: https://developer.apple.com/documentation/CloudKit/CKRecord/ID/recordName
//
// [CKRecordNameZoneWideShare]: https://developer.apple.com/documentation/CloudKit/CKRecordNameZoneWideShare
func (c CKRecordID) RecordName() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("recordName"))
	return foundation.NSStringFromID(rv).String()
}

// The ID of the zone that contains the record.
//
// See: https://developer.apple.com/documentation/CloudKit/CKRecord/ID/zoneID
func (c CKRecordID) ZoneID() ICKRecordZoneID {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("zoneID"))
	return CKRecordZoneIDFromID(objc.ID(rv))
}
