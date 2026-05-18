// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CKRecordZone] class.
var (
	_CKRecordZoneClass     CKRecordZoneClass
	_CKRecordZoneClassOnce sync.Once
)

func getCKRecordZoneClass() CKRecordZoneClass {
	_CKRecordZoneClassOnce.Do(func() {
		_CKRecordZoneClass = CKRecordZoneClass{class: objc.GetClass("CKRecordZone")}
	})
	return _CKRecordZoneClass
}

// GetCKRecordZoneClass returns the class object for CKRecordZone.
func GetCKRecordZoneClass() CKRecordZoneClass {
	return getCKRecordZoneClass()
}

type CKRecordZoneClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CKRecordZoneClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CKRecordZoneClass) Alloc() CKRecordZone {
	rv := objc.Send[CKRecordZone](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// A database partition that contains related records.
//
// # Overview
//
// Zones are an important part of how you organize your data. The public and
// private databases each have a single default zone. In the private database,
// you can use [CKRecordZone] objects to create additional custom zones as
// necessary. Use custom zones to arrange and encapsulate groups of related
// records in the private database. Custom zones support other capabilities
// too, such as the ability to write multiple records as a single atomic
// transaction.
//
// Treat each custom zone as a single unit of data that is separate from every
// other zone in the database. You can add records inside the zone. You can
// also create links between the records inside a zone by using the
// [CKReference] class. However, the [CKReference] class doesn’t support
// cross-zone linking, so each reference object must point to a record in the
// same zone as the current record.
//
// Use the [CKRecordZone] class as-is and don’t subclass it.
//
// # Creating a Custom Record Zone
//
// Generally, you use instances of this class to create and manage custom
// zones. Although you can use this class to retrieve a database’s default
// zone, most operations act on records in the default zone by default, so you
// rarely need to specify it explicitly.
//
// To create a custom zone, use [CKRecordZone] to create the zone object, and
// then save that zone to the user’s private database using a
// [CKModifyRecordZonesOperation] object. You can’t save any records in the
// zone until you save it to the database. When creating records, explicitly
// specify the zone ID if you want the records to reside in a specific zone;
// otherwise, they save to the default zone. You can’t create custom zones
// in a public database.
//
// After creating a [CKRecordZone] object and saving it to the database, you
// don’t interact with the object much. Instead, most interactions occur
// with its corresponding [CKRecordZoneID] object, which you use to refer to
// the zone when creating records.
//
// # Creating a Record Zone
//
//   - [CKRecordZone.InitWithZoneName]: Creates a record zone object with the specified zone name.
//   - [CKRecordZone.InitWithZoneID]: Creates a record zone object with the specified zone ID.
//
// # Getting the Zone Attributes
//
//   - [CKRecordZone.ZoneID]: The unique ID of the zone.
//   - [CKRecordZone.Capabilities]: The capabilities that the zone supports.
//
// # Sharing Records
//
//   - [CKRecordZone.Share]: A reference to the record zone’s share record.
//
// # Instance Properties
//
//   - [CKRecordZone.EncryptionScope]: The encryption scope determines the granularity at which encryption keys are stored within the zone.
//   - [CKRecordZone.SetEncryptionScope]
//
// See: https://developer.apple.com/documentation/CloudKit/CKRecordZone
type CKRecordZone struct {
	objectivec.Object
}

// CKRecordZoneFromID constructs a [CKRecordZone] from an objc.ID.
//
// A database partition that contains related records.
func CKRecordZoneFromID(id objc.ID) CKRecordZone {
	return CKRecordZone{objectivec.Object{ID: id}}
}

// NOTE: CKRecordZone adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CKRecordZone] class.
//
// # Creating a Record Zone
//
//   - [ICKRecordZone.InitWithZoneName]: Creates a record zone object with the specified zone name.
//   - [ICKRecordZone.InitWithZoneID]: Creates a record zone object with the specified zone ID.
//
// # Getting the Zone Attributes
//
//   - [ICKRecordZone.ZoneID]: The unique ID of the zone.
//   - [ICKRecordZone.Capabilities]: The capabilities that the zone supports.
//
// # Sharing Records
//
//   - [ICKRecordZone.Share]: A reference to the record zone’s share record.
//
// # Instance Properties
//
//   - [ICKRecordZone.EncryptionScope]: The encryption scope determines the granularity at which encryption keys are stored within the zone.
//   - [ICKRecordZone.SetEncryptionScope]
//
// See: https://developer.apple.com/documentation/CloudKit/CKRecordZone
type ICKRecordZone interface {
	objectivec.IObject

	// Topic: Creating a Record Zone

	// Creates a record zone object with the specified zone name.
	InitWithZoneName(zoneName string) CKRecordZone
	// Creates a record zone object with the specified zone ID.
	InitWithZoneID(zoneID ICKRecordZoneID) CKRecordZone

	// Topic: Getting the Zone Attributes

	// The unique ID of the zone.
	ZoneID() ICKRecordZoneID
	// The capabilities that the zone supports.
	Capabilities() CKRecordZoneCapabilities

	// Topic: Sharing Records

	// A reference to the record zone’s share record.
	Share() ICKReference

	// Topic: Instance Properties

	// The encryption scope determines the granularity at which encryption keys are stored within the zone.
	EncryptionScope() CKRecordZoneEncryptionScope
	SetEncryptionScope(value CKRecordZoneEncryptionScope)

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (c CKRecordZone) Init() CKRecordZone {
	rv := objc.Send[CKRecordZone](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CKRecordZone) Autorelease() CKRecordZone {
	rv := objc.Send[CKRecordZone](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKRecordZone creates a new CKRecordZone instance.
func NewCKRecordZone() CKRecordZone {
	class := getCKRecordZoneClass()
	rv := objc.Send[CKRecordZone](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a record zone object with the specified zone ID.
//
// zoneID: The ID for the new zone. This parameter must not be `nil`.
//
// # Return Value
//
// The custom record zone, or `nil` if CloudKit can’t create the zone.
//
// # Discussion
//
// Use this method when you want to create a new record zone from the
// information in a zone ID. After creating the zone, save it to the server
// using a [CKModifyRecordZonesOperation] object or the
// [SaveRecordZoneCompletionHandler] method of [CKDatabase].
//
// Don’t use this method to create a [CKRecordZone] object that corresponds
// to a zone that already exists in the database. If the zone exists, fetch it
// using a [CKFetchRecordZonesOperation] object or the
// [FetchRecordZoneWithIDCompletionHandler] method of [CKDatabase].
//
// See: https://developer.apple.com/documentation/CloudKit/CKRecordZone/init(zoneID:)
func NewCKRecordZoneWithZoneID(zoneID ICKRecordZoneID) CKRecordZone {
	instance := getCKRecordZoneClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithZoneID:"), zoneID)
	return CKRecordZoneFromID(rv)
}

// Creates a record zone object with the specified zone name.
//
// zoneName: The name of the new zone. Zone names inside a user’s private database are
// unique, consist of up to 255 ASCII characters, and don’t start with an
// underscore. One way to ensure the uniqueness of zone names is to create a
// string from a UUID, but you can also use other techniques.
//
// If this parameter is `nil` or is an empty string, the method throws an
// exception.
//
// # Return Value
//
// The new custom zone, or `nil` if CloudKit can’t create the zone.
//
// # Discussion
//
// Use this method to create a new record zone. The new zone has the name you
// provide and the zone’s owner is the current user. After creating the
// zone, save it to the server using a [CKModifyRecordZonesOperation] object
// or the [SaveRecordZoneCompletionHandler] method of [CKDatabase]. You must
// save the zone to the server before you attempt to save any records to that
// zone.
//
// Don’t use this method to create a [CKRecordZone] object that corresponds
// to a zone that already exists in the database. If the zone exists, fetch it
// using a [CKFetchRecordZonesOperation] object or the
// [FetchRecordZoneWithIDCompletionHandler] method of [CKDatabase].
//
// See: https://developer.apple.com/documentation/CloudKit/CKRecordZone/init(zoneName:)
func NewCKRecordZoneWithZoneName(zoneName string) CKRecordZone {
	instance := getCKRecordZoneClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithZoneName:"), objc.String(zoneName))
	return CKRecordZoneFromID(rv)
}

// Creates a record zone object with the specified zone name.
//
// zoneName: The name of the new zone. Zone names inside a user’s private database are
// unique, consist of up to 255 ASCII characters, and don’t start with an
// underscore. One way to ensure the uniqueness of zone names is to create a
// string from a UUID, but you can also use other techniques.
//
// If this parameter is `nil` or is an empty string, the method throws an
// exception.
//
// # Return Value
//
// The new custom zone, or `nil` if CloudKit can’t create the zone.
//
// # Discussion
//
// Use this method to create a new record zone. The new zone has the name you
// provide and the zone’s owner is the current user. After creating the
// zone, save it to the server using a [CKModifyRecordZonesOperation] object
// or the [SaveRecordZoneCompletionHandler] method of [CKDatabase]. You must
// save the zone to the server before you attempt to save any records to that
// zone.
//
// Don’t use this method to create a [CKRecordZone] object that corresponds
// to a zone that already exists in the database. If the zone exists, fetch it
// using a [CKFetchRecordZonesOperation] object or the
// [FetchRecordZoneWithIDCompletionHandler] method of [CKDatabase].
//
// See: https://developer.apple.com/documentation/CloudKit/CKRecordZone/init(zoneName:)
func (c CKRecordZone) InitWithZoneName(zoneName string) CKRecordZone {
	rv := objc.Send[CKRecordZone](c.ID, objc.Sel("initWithZoneName:"), objc.String(zoneName))
	return rv
}

// Creates a record zone object with the specified zone ID.
//
// zoneID: The ID for the new zone. This parameter must not be `nil`.
//
// # Return Value
//
// The custom record zone, or `nil` if CloudKit can’t create the zone.
//
// # Discussion
//
// Use this method when you want to create a new record zone from the
// information in a zone ID. After creating the zone, save it to the server
// using a [CKModifyRecordZonesOperation] object or the
// [SaveRecordZoneCompletionHandler] method of [CKDatabase].
//
// Don’t use this method to create a [CKRecordZone] object that corresponds
// to a zone that already exists in the database. If the zone exists, fetch it
// using a [CKFetchRecordZonesOperation] object or the
// [FetchRecordZoneWithIDCompletionHandler] method of [CKDatabase].
//
// See: https://developer.apple.com/documentation/CloudKit/CKRecordZone/init(zoneID:)
func (c CKRecordZone) InitWithZoneID(zoneID ICKRecordZoneID) CKRecordZone {
	rv := objc.Send[CKRecordZone](c.ID, objc.Sel("initWithZoneID:"), zoneID)
	return rv
}
func (c CKRecordZone) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](c.ID, objc.Sel("encodeWithCoder:"), coder)
}

// Returns the default record zone.
//
// # Discussion
//
// Always use this method to retrieve the default zone for a database. You can
// use the returned object to specify the default zone for either the public
// or private database of a container. You don’t need to save the returned
// zone object before using it. The owner of the zone is [CKOwnerDefaultName],
// which corresponds to the current user.
//
// The default zone of a database is a convenient place to store and access
// records. If you don’t explicitly assign a zone to a record, CloudKit puts
// the record in the default zone.
//
// The disadvantage of using the default zone for storing records is that it
// doesn’t have any special capabilities. You can’t save a group of
// records to iCloud atomically in the default zone. Similarly, you can’t
// use a [CKFetchRecordChangesOperation] object on records in the default
// zone.
//
// See: https://developer.apple.com/documentation/CloudKit/CKRecordZone/default()
//
// [CKFetchRecordChangesOperation]: https://developer.apple.com/documentation/CloudKit/CKFetchRecordChangesOperation
func (_CKRecordZoneClass CKRecordZoneClass) DefaultRecordZone() CKRecordZone {
	rv := objc.Send[objc.ID](objc.ID(_CKRecordZoneClass.class), objc.Sel("defaultRecordZone"))
	return CKRecordZoneFromID(rv)
}

// The unique ID of the zone.
//
// # Discussion
//
// The zone ID contains the name of the zone and the name of the user who owns
// the zone. Use this property to access both of those values.
//
// See: https://developer.apple.com/documentation/CloudKit/CKRecordZone/zoneID
func (c CKRecordZone) ZoneID() ICKRecordZoneID {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("zoneID"))
	return CKRecordZoneIDFromID(objc.ID(rv))
}

// The capabilities that the zone supports.
//
// # Discussion
//
// The server determines the capabilities of the zone and sets the value of
// this property when you save the record zone. Always check this property
// before performing tasks that require a specific capability.
//
// Default zones don’t support any special capabilities. Custom zones in a
// private database support the options that [CKRecordZone.Capabilities]
// provides.
//
// See: https://developer.apple.com/documentation/CloudKit/CKRecordZone/capabilities-swift.property
//
// [CKRecordZone.Capabilities]: https://developer.apple.com/documentation/CloudKit/CKRecordZone/Capabilities-swift.struct
func (c CKRecordZone) Capabilities() CKRecordZoneCapabilities {
	rv := objc.Send[CKRecordZoneCapabilities](c.ID, objc.Sel("capabilities"))
	return CKRecordZoneCapabilities(rv)
}

// A reference to the record zone’s share record.
//
// # Discussion
//
// CloudKit sets this property only for fetched record zones that contain a
// share record; otherwise, it’s `nil`.
//
// To share a record zone, create a share record using the
// [InitWithRecordZoneID] method and then save it to the server. Shared record
// zones must have the [CKRecordZoneCapabilityZoneWideSharing] capability,
// which CloudKit enables by default for new custom record zones in the
// user’s private database.
//
// A record zone, and the records it contains, can take part in only a single
// share. CloudKit returns an error if you attempt to share an already-shared
// record zone, or if that record zone contains previously shared records.
//
// Record zone sharing errors include the following:
//
// - [CKError.Code.serverRecordChanged], which CloudKit returns if you try to
// share an already-shared record zone. -
// [CKError.Code.serverRejectedRequest], which CloudKit returns if you try to
// share a record hierarchy from an already-shared record zone. -
// [CKError.Code.invalidArguments], which CloudKit returns if you try to share
// a record zone that contains one or more shared hierarchies.
//
// See: https://developer.apple.com/documentation/CloudKit/CKRecordZone/share
//
// [CKError.Code.invalidArguments]: https://developer.apple.com/documentation/CloudKit/CKError/Code/invalidArguments
// [CKError.Code.serverRecordChanged]: https://developer.apple.com/documentation/CloudKit/CKError/Code/serverRecordChanged
// [CKError.Code.serverRejectedRequest]: https://developer.apple.com/documentation/CloudKit/CKError/Code/serverRejectedRequest
func (c CKRecordZone) Share() ICKReference {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("share"))
	return CKReferenceFromID(objc.ID(rv))
}

// The encryption scope determines the granularity at which encryption keys
// are stored within the zone.
//
// # Discussion
//
// Zone encryption scope defaults to [CKRecordZoneEncryptionScopePerRecord]
// and can only be modified before zone creation. Attempting to change the
// encryption scope of an existing zone is invalid and will result in an
// error.
//
// Zones using [CKRecordZoneEncryptionScopePerZone] can only use zone-wide
// sharing and are not compatible with older device OS versions. Refer to
// [CKRecordZoneEncryptionScope] for more info.
//
// See: https://developer.apple.com/documentation/CloudKit/CKRecordZone/encryptionScope-swift.property
func (c CKRecordZone) EncryptionScope() CKRecordZoneEncryptionScope {
	rv := objc.Send[CKRecordZoneEncryptionScope](c.ID, objc.Sel("encryptionScope"))
	return CKRecordZoneEncryptionScope(rv)
}
func (c CKRecordZone) SetEncryptionScope(value CKRecordZoneEncryptionScope) {
	objc.Send[struct{}](c.ID, objc.Sel("setEncryptionScope:"), value)
}
