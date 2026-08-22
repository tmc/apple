// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CKRecordZoneID] class.
var (
	_CKRecordZoneIDClass     CKRecordZoneIDClass
	_CKRecordZoneIDClassOnce sync.Once
)

func getCKRecordZoneIDClass() CKRecordZoneIDClass {
	_CKRecordZoneIDClassOnce.Do(func() {
		_CKRecordZoneIDClass = CKRecordZoneIDClass{class: objc.GetClass("CKRecordZoneID")}
	})
	return _CKRecordZoneIDClass
}

// GetCKRecordZoneIDClass returns the class object for CKRecordZoneID.
func GetCKRecordZoneIDClass() CKRecordZoneIDClass {
	return getCKRecordZoneIDClass()
}

type CKRecordZoneIDClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CKRecordZoneIDClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CKRecordZoneIDClass) Alloc() CKRecordZoneID {
	rv := objc.Send[CKRecordZoneID](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// An object that uniquely identifies a record zone in a database.
//
// # Overview
//
// Zones are a mechanism for grouping related records together. You create
// zone ID objects when you want to fetch an existing zone object or create a
// new zone with a specific name.
//
// A record zone ID distinguishes one zone from another by a name string and
// the ID of the user who creates the zone. Both strings must be ASCII strings
// that don’t exceed 255 characters. When creating your own record zone ID
// objects, you can use names that have more meaning to your app or to the
// user, providing each zone name is unique within the specified database. The
// owner name must be either the current user name or the name of another
// user. Get the current user name from [CKCurrentUserDefaultName] or by
// calling [CKContainer.FetchUserRecordIDWithCompletionHandler].
//
// When creating new record zones, make the name string in the record zone ID
// unique in the target database. Public databases don’t support custom
// zones, and only the user who owns the database can create zones in private
// databases.
//
// Don’t create subclasses of this class.
//
// # Interacting with Record Zone IDs
//
// After you create a record zone ID, interactions with it typically include:
//
// - Creating a [CKRecordID] object so that you can fetch or create records in
// that zone. - Retrieving an existing [CKRecordZone] object from the
// database.
//
// You don’t need to create a record zone ID to create a record zone. The
// [CKRecordZone] class has initialization methods that create a record zone
// ID using the name string you provide.
//
// # Creating Record Zone IDs for Records
//
// To create a new record in a custom zone, create a record zone ID that
// specifies the zone name. Use the record zone ID to create a [CKRecordID],
// and then use the record ID to create the record.
//
// # Fetching a Record Zone Object from the Database
//
// To fetch a record zone from the database, use a
// [CKFetchRecordZonesOperation] object or the
// [CKDatabase.FetchRecordZoneWithIDCompletionHandler] method of [CKDatabase].
// Both techniques accept a record zone ID that you provide and retrieve the
// corresponding record zone object asynchronously. If you use the operation
// object, you can retrieve multiple record zones at the same time.
//
// # Getting the Record Zone ID Attributes
//
//   - [CKRecordZoneID.ZoneName]: The unique name of the record zone.
//   - [CKRecordZoneID.OwnerName]: The ID of the user who owns the record zone.
//
// See: https://developer.apple.com/documentation/CloudKit/CKRecordZone/ID
//
// [CKCurrentUserDefaultName]: https://developer.apple.com/documentation/CloudKit/CKCurrentUserDefaultName
type CKRecordZoneID struct {
	objectivec.Object
}

// CKRecordZoneIDFromID constructs a [CKRecordZoneID] from an objc.ID.
//
// An object that uniquely identifies a record zone in a database.
func CKRecordZoneIDFromID(id objc.ID) CKRecordZoneID {
	return CKRecordZoneID{objectivec.Object{ID: id}}
}

// NOTE: CKRecordZoneID adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CKRecordZoneID] class.
//
// # Getting the Record Zone ID Attributes
//
//   - [ICKRecordZoneID.ZoneName]: The unique name of the record zone.
//   - [ICKRecordZoneID.OwnerName]: The ID of the user who owns the record zone.
//
// See: https://developer.apple.com/documentation/CloudKit/CKRecordZone/ID
type ICKRecordZoneID interface {
	objectivec.IObject

	// Topic: Getting the Record Zone ID Attributes

	// The unique name of the record zone.
	ZoneName() string
	// The ID of the user who owns the record zone.
	OwnerName() string

	EncodeWithCoder(coder foundation.INSCoder)
	InitWithZoneNameOwnerName(zoneName string, ownerName string) CKRecordZoneID
}

// Init initializes the instance.
func (c CKRecordZoneID) Init() CKRecordZoneID {
	rv := objc.Send[CKRecordZoneID](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CKRecordZoneID) Autorelease() CKRecordZoneID {
	rv := objc.Send[CKRecordZoneID](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKRecordZoneID creates a new CKRecordZoneID instance.
func NewCKRecordZoneID() CKRecordZoneID {
	class := getCKRecordZoneIDClass()
	rv := objc.Send[CKRecordZoneID](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (c CKRecordZoneID) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](c.ID, objc.Sel("encodeWithCoder:"), coder)
}
func (c CKRecordZoneID) InitWithZoneNameOwnerName(zoneName string, ownerName string) CKRecordZoneID {
	rv := objc.Send[CKRecordZoneID](c.ID, objc.Sel("initWithZoneName:ownerName:"), objc.String(zoneName), objc.String(ownerName))
	return rv
}

// The unique name of the record zone.
//
// See: https://developer.apple.com/documentation/CloudKit/CKRecordZone/ID/zoneName
func (c CKRecordZoneID) ZoneName() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("zoneName"))
	return foundation.NSStringFromID(rv).String()
}

// The ID of the user who owns the record zone.
//
// See: https://developer.apple.com/documentation/CloudKit/CKRecordZone/ID/ownerName
func (c CKRecordZoneID) OwnerName() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("ownerName"))
	return foundation.NSStringFromID(rv).String()
}
