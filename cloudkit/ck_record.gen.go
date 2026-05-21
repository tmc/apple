// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/kernel"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CKRecord] class.
var (
	_CKRecordClass     CKRecordClass
	_CKRecordClassOnce sync.Once
)

func getCKRecordClass() CKRecordClass {
	_CKRecordClassOnce.Do(func() {
		_CKRecordClass = CKRecordClass{class: objc.GetClass("CKRecord")}
	})
	return _CKRecordClass
}

// GetCKRecordClass returns the class object for CKRecord.
func GetCKRecordClass() CKRecordClass {
	return getCKRecordClass()
}

type CKRecordClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CKRecordClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CKRecordClass) Alloc() CKRecord {
	rv := objc.Send[CKRecord](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// A collection of key-value pairs that store your app’s data.
//
// # Overview
//
// Records are the fundamental objects that manage data in CloudKit. You can
// define any number of record types for your app, with each record type
// corresponding to a different type of information. Within a record type, you
// then define one or more fields, each with a name and a value. Records can
// contain simple data types, such as strings and numbers, or more complex
// types, such as geographic locations or pointers to other records.
//
// An important step in using CloudKit is defining the record types your app
// supports. A new record object doesn’t contain any keys or values. During
// development, you can add new keys and values at any time. The first time
// you set a value for a key and save the record, the server associates that
// type with the key for all records of the same type. The [CKRecord] class
// doesn’t impose these type constraints or do any local validation of a
// record’s contents. CloudKit enforces these constraints when you save the
// records.
//
// Although records behave like dictionaries, there are limitations to the
// types of values you can assign to keys. The following are the object types
// that the [CKRecord] class supports. Attempting to specify objects of any
// other type results in failure. Fields of all types are searchable unless
// otherwise noted.
//
// # Supported Data Types
//
// [CKRecord] fields support the following data types:
//
// [NSString]: Stores relatively small amounts of text. Although strings
// themselves can be any length, use a [CKAsset] to store large amounts of
// text. [NSNumber]: Stores any numerical information, including integers and
// floating-point numbers. [NSData]: Stores arbitrary bytes of data. A typical
// use for data objects is to map the bytes that they contain to a `struct`.
// Don’t use data objects for storing large binary data files; use a
// [CKAsset] instead. Data fields aren’t searchable. [NSDate]: Stores day
// and time information in an accessible form. [NSArray]: Stores one or more
// objects of any other type in this table. You can store arrays of strings,
// arrays of numbers, arrays of references, and so on. [CLLocation]: Stores
// geographic coordinate data. You use locations in conjunction with the Core
// Location framework and any other services that handle location information.
// [CKAsset]: Associates a disk-based file with the record. Although assets
// have a close association with records, you manage them separately. For more
// information about using assets, see [CKAsset]. [CKReference]: Creates a
// link to a related record. A reference stores the ID of the target record.
// The advantage of using a reference instead of storing the ID as a string is
// that references can initiate cascade deletions of dependent records. The
// disadvantage is that references can only link between records in the same
// record zone. For more information, see [CKReference].
//
// # Defining Records
//
// The process for defining your record types depends entirely on your app and
// the data you’re trying to represent. It’s best to design records that
// encapsulate data for one unit of information. For example, you might use
// one record type to store an employee’s name, job title, and date of hire,
// and use a separate record type to store the employee’s address
// information. Using different record types lets you manage, manipulate, and
// validate the two types of information separately.
//
// Use fields that contain [CKReference] objects to establish relationships
// between different types of records. After you define your record types, use
// the iCloud Dashboard to set them up. During development, you can also
// create new record types programmatically.
//
// # Indexing the Fields of a Record
//
// Indexes make it possible to search the contents of your records
// efficiently. During development, the server indexes all fields with data
// types it can use in the predicate of a query. This automatic indexing makes
// it easier to experiment with queries during development, but the indexes
// require space in a database, and require time to generate and maintain.
//
// To manage the indexing behavior of your records in the production
// environment, use CloudKit Dashboard. When migrating your schema from the
// development environment to the production environment, enable indexing only
// for the fields that your app uses in queries, and disable it for all other
// fields.
//
// # Customizing Records
//
// Use this class as-is to manage data coming from or going to the server, and
// don’t subclass it.
//
// # Storing Records Locally
//
// If you store records in a local database, use the
// [CKRecord.EncodeSystemFieldsWithCoder] method to encode and store the
// record’s metadata. The metadata contains the record ID and the change
// tag, which you need later to sync records in a local database with those in
// CloudKit.
//
// # Accessing the Record’s Metadata
//
//   - [CKRecord.RecordID]: The unique ID of the record.
//   - [CKRecord.RecordType]: The value that your app defines to identify the type of record.
//   - [CKRecord.SetRecordType]
//   - [CKRecord.CreationDate]: The time when CloudKit first saves the record to the server.
//   - [CKRecord.CreatorUserRecordID]: The ID of the user who creates the record.
//   - [CKRecord.ModificationDate]: The most recent time that CloudKit saved the record to the server.
//   - [CKRecord.LastModifiedUserRecordID]: The ID of the user who most recently modified the record.
//   - [CKRecord.RecordChangeTag]: The server change token for the record.
//
// # Encrypting the Record’s Values
//
//   - [CKRecord.EncryptedValues]: An object that manages the record’s encrypted key-value pairs.
//
// # Getting Data for Full-Text Searches
//
//   - [CKRecord.AllTokens]: Returns an array of strings to use for full-text searches of the field’s string-based values.
//
// # Encoding the Record’s Metadata
//
//   - [CKRecord.EncodeSystemFieldsWithCoder]: Encodes the record’s system fields using the specified archiver.
//
// # Sharing Records
//
//   - [CKRecord.Parent]: A reference to the record’s parent record.
//   - [CKRecord.SetParent]
//   - [CKRecord.Share]: A reference to the share object that determines the share status of the record.
//   - [CKRecord.SetParentReferenceFromRecord]: Creates and sets a reference object for a parent from its record.
//   - [CKRecord.SetParentReferenceFromRecordID]: Creates and sets a reference object for a parent from the parent’s record ID.
//
// # Initializers
//
//   - [CKRecord.InitWithCoder]
//
// See: https://developer.apple.com/documentation/CloudKit/CKRecord
//
// [CLLocation]: https://developer.apple.com/documentation/CoreLocation/CLLocation
// [NSArray]: https://developer.apple.com/documentation/Foundation/NSArray
// [NSData]: https://developer.apple.com/documentation/Foundation/NSData
// [NSDate]: https://developer.apple.com/documentation/Foundation/NSDate
// [NSNumber]: https://developer.apple.com/documentation/Foundation/NSNumber
// [NSString]: https://developer.apple.com/documentation/Foundation/NSString
type CKRecord struct {
	objectivec.Object
}

// CKRecordFromID constructs a [CKRecord] from an objc.ID.
//
// A collection of key-value pairs that store your app’s data.
func CKRecordFromID(id objc.ID) CKRecord {
	return CKRecord{objectivec.Object{ID: id}}
}

// NOTE: CKRecord adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CKRecord] class.
//
// # Accessing the Record’s Metadata
//
//   - [ICKRecord.RecordID]: The unique ID of the record.
//   - [ICKRecord.RecordType]: The value that your app defines to identify the type of record.
//   - [ICKRecord.SetRecordType]
//   - [ICKRecord.CreationDate]: The time when CloudKit first saves the record to the server.
//   - [ICKRecord.CreatorUserRecordID]: The ID of the user who creates the record.
//   - [ICKRecord.ModificationDate]: The most recent time that CloudKit saved the record to the server.
//   - [ICKRecord.LastModifiedUserRecordID]: The ID of the user who most recently modified the record.
//   - [ICKRecord.RecordChangeTag]: The server change token for the record.
//
// # Encrypting the Record’s Values
//
//   - [ICKRecord.EncryptedValues]: An object that manages the record’s encrypted key-value pairs.
//
// # Getting Data for Full-Text Searches
//
//   - [ICKRecord.AllTokens]: Returns an array of strings to use for full-text searches of the field’s string-based values.
//
// # Encoding the Record’s Metadata
//
//   - [ICKRecord.EncodeSystemFieldsWithCoder]: Encodes the record’s system fields using the specified archiver.
//
// # Sharing Records
//
//   - [ICKRecord.Parent]: A reference to the record’s parent record.
//   - [ICKRecord.SetParent]
//   - [ICKRecord.Share]: A reference to the share object that determines the share status of the record.
//   - [ICKRecord.SetParentReferenceFromRecord]: Creates and sets a reference object for a parent from its record.
//   - [ICKRecord.SetParentReferenceFromRecordID]: Creates and sets a reference object for a parent from the parent’s record ID.
//
// # Initializers
//
//   - [ICKRecord.InitWithCoder]
//
// See: https://developer.apple.com/documentation/CloudKit/CKRecord
type ICKRecord interface {
	objectivec.IObject

	// Topic: Accessing the Record’s Metadata

	// The unique ID of the record.
	RecordID() ICKRecordID
	// The value that your app defines to identify the type of record.
	RecordType() unsafe.Pointer
	SetRecordType(value kernel.Pointer)
	// The time when CloudKit first saves the record to the server.
	CreationDate() foundation.NSDate
	// The ID of the user who creates the record.
	CreatorUserRecordID() ICKRecordID
	// The most recent time that CloudKit saved the record to the server.
	ModificationDate() foundation.NSDate
	// The ID of the user who most recently modified the record.
	LastModifiedUserRecordID() ICKRecordID
	// The server change token for the record.
	RecordChangeTag() string

	// Topic: Encrypting the Record’s Values

	// An object that manages the record’s encrypted key-value pairs.
	EncryptedValues() CKRecordKeyValueSetting

	// Topic: Getting Data for Full-Text Searches

	// Returns an array of strings to use for full-text searches of the field’s string-based values.
	AllTokens() []string

	// Topic: Encoding the Record’s Metadata

	// Encodes the record’s system fields using the specified archiver.
	EncodeSystemFieldsWithCoder(coder foundation.INSCoder)

	// Topic: Sharing Records

	// A reference to the record’s parent record.
	Parent() ICKReference
	SetParent(value ICKReference)
	// A reference to the share object that determines the share status of the record.
	Share() ICKReference
	// Creates and sets a reference object for a parent from its record.
	SetParentReferenceFromRecord(parentRecord ICKRecord)
	// Creates and sets a reference object for a parent from the parent’s record ID.
	SetParentReferenceFromRecordID(parentRecordID ICKRecordID)

	// Topic: Initializers

	InitWithCoder(coder foundation.INSCoder) CKRecord

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (c CKRecord) Init() CKRecord {
	rv := objc.Send[CKRecord](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CKRecord) Autorelease() CKRecord {
	rv := objc.Send[CKRecord](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKRecord creates a new CKRecord instance.
func NewCKRecord() CKRecord {
	class := getCKRecordClass()
	rv := objc.Send[CKRecord](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CloudKit/CKRecord/init(coder:)
func NewCKRecordWithCoder(coder foundation.INSCoder) CKRecord {
	instance := getCKRecordClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return CKRecordFromID(rv)
}

// Returns the object that the record stores for the specified key.
//
// key: The string that identifies a field in the record. A key must consist of one
// or more alphanumeric characters and must start with a letter. CloudKit
// permits the use of underscores, but not spaces.
//
// # Return Value
//
// The object for the specified key, or `nil` if no such key exists in the
// record.
//
// # Discussion
//
// See: https://developer.apple.com/documentation/CloudKit/CKRecord/subscript(_:)-51whk
func (c CKRecord) ObjectForKeyedSubscript(key CKRecordFieldKey) CKRecordValue {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("objectForKeyedSubscript:"), objc.String(string(key)))
	return CKRecordValueObjectFromID(rv)
}

// Returns an array of strings to use for full-text searches of the field’s
// string-based values.
//
// # Return Value
//
// An array of strings that contains data from the record’s string-based
// fields.
//
// # Discussion
//
// When performing your own full-text searches, you can use this method to get
// a list of strings for your search. The method acts only on keys with string
// values. It breaks each value string apart at whitespace boundaries, creates
// new strings for each word, adds the new strings to an array, and returns
// the array. This tokenized version of the record’s string values makes it
// easier to do string-based comparisons of individual words.
//
// See: https://developer.apple.com/documentation/CloudKit/CKRecord/allTokens()
func (c CKRecord) AllTokens() []string {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("allTokens"))
	return objc.ConvertSliceToStrings(rv)
}

// Encodes the record’s system fields using the specified archiver.
//
// coder: An archiver object.
//
// # Discussion
//
// Use this method to encode the record’s metadata that CloudKit provides.
// Every record has keys that the system defines that correspond to record
// metadata, such as the record ID, record type, creation date, and so on.
// This method encodes those keys in the specified archiver. This method
// doesn’t include any keys you add to the record. It also doesn’t encode
// the keys that the [changedKeys] method returns.
//
// You might use this method when you want to store only the system metadata
// because you store the actual record data elsewhere.
//
// See: https://developer.apple.com/documentation/CloudKit/CKRecord/encodeSystemFields(with:)
//
// [changedKeys]: https://developer.apple.com/documentation/CloudKit/CKRecord/changedKeys
func (c CKRecord) EncodeSystemFieldsWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](c.ID, objc.Sel("encodeSystemFieldsWithCoder:"), coder)
}

// Creates and sets a reference object for a parent from its record.
//
// parentRecord: A record that you want to set as the parent to this record.
//
// # Discussion
//
// This method creates and sets a [CKReference] object that points to the
// record you provide. The resulting [CKReference] has an action of
// [CKRecord.ReferenceAction.none].
//
// See: https://developer.apple.com/documentation/CloudKit/CKRecord/setParent(_:)-23du1
//
// [CKRecord.ReferenceAction.none]: https://developer.apple.com/documentation/CloudKit/CKRecord/ReferenceAction/none
func (c CKRecord) SetParentReferenceFromRecord(parentRecord ICKRecord) {
	objc.Send[objc.ID](c.ID, objc.Sel("setParentReferenceFromRecord:"), parentRecord)
}

// Creates and sets a reference object for a parent from the parent’s record
// ID.
//
// parentRecordID: The [CKRecordID] object for the record that you want to set as this
// record’s parent.
//
// # Discussion
//
// This method creates and sets a [CKReference] object that points to the
// record you provide. The resulting [CKReference] has an action of
// [CKRecord.ReferenceAction.none].
//
// See: https://developer.apple.com/documentation/CloudKit/CKRecord/setParent(_:)-7egcx
//
// [CKRecord.ReferenceAction.none]: https://developer.apple.com/documentation/CloudKit/CKRecord/ReferenceAction/none
func (c CKRecord) SetParentReferenceFromRecordID(parentRecordID ICKRecordID) {
	objc.Send[objc.ID](c.ID, objc.Sel("setParentReferenceFromRecordID:"), parentRecordID)
}

// See: https://developer.apple.com/documentation/CloudKit/CKRecord/init(coder:)
func (c CKRecord) InitWithCoder(coder foundation.INSCoder) CKRecord {
	rv := objc.Send[CKRecord](c.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}

// Returns an array of the record’s keys.
//
// # Return Value
//
// An array of keys, or an empty array if the record doesn’t contain any
// keys.
//
// See: https://developer.apple.com/documentation/CloudKit/CKRecordKeyValueSetting/allKeys()
func (c CKRecord) AllKeys() []string {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("allKeys"))
	return objc.ConvertSliceToStrings(rv)
}

// Returns an array of keys with recent changes to their values.
//
// # Return Value
//
// An array of keys with changed values since downloading or saving the
// record. If there aren’t any changed keys, this method returns an empty
// array.
//
// See: https://developer.apple.com/documentation/CloudKit/CKRecordKeyValueSetting/changedKeys()
func (c CKRecord) ChangedKeys() []string {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("changedKeys"))
	return objc.ConvertSliceToStrings(rv)
}

// Returns the object that the record stores for the specified key.
//
// key: The string that identifies a field in the record. A key must consist of one
// or more alphanumeric characters and must start with a letter. CloudKit
// permits the use of underscores, but not spaces.
//
// # Return Value
//
// The object for the specified key, or `nil` if no such key exists in the
// record.
//
// See: https://developer.apple.com/documentation/CloudKit/CKRecordKeyValueSetting/object(forKey:)
func (c CKRecord) ObjectForKey(key CKRecordFieldKey) CKRecordValue {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("objectForKey:"), objc.String(string(key)))
	return CKRecordValueObjectFromID(rv)
}

// Stores an object in the record using the specified key.
//
// object: The object to store using the specified key. It must be one of the data
// types in [CKRecord]. You receive an error if you use a data type that
// CloudKit doesn’t support. If you specify `nil`, CloudKit removes any
// object that the record associates with the key.
//
// key: The key to associate with `object`. Use this key to retrieve the value
// later. A key must consist of one or more alphanumeric characters and must
// start with a letter. CloudKit permits the use of underscores, but not
// spaces. Avoid using a key that matches the name of any property of
// [CKRecord].
//
// See: https://developer.apple.com/documentation/CloudKit/CKRecordKeyValueSetting/setObject(_:forKey:)
func (c CKRecord) SetObjectForKey(object CKRecordValue, key CKRecordFieldKey) {
	objc.Send[objc.ID](c.ID, objc.Sel("setObject:forKey:"), object, objc.String(string(key)))
}
func (c CKRecord) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](c.ID, objc.Sel("encodeWithCoder:"), coder)
}

// The unique ID of the record.
//
// # Discussion
//
// The system sets the ID of a new record at initialization time. If you use
// the [init(recordType:recordID:)] method to initialize the record, the ID
// derives from the [CKRecordID] object you provide. In all other cases, the
// record generates a UUID and bases its ID on that value. The ID of a record
// never changes during its lifetime.
//
// When you save a new record object to the server, the server validates the
// uniqueness of the record, but returns an error only if the save policy
// calls for it. Specifically, it returns an error when the save policy is
// [CKModifyRecordsOperation.RecordSavePolicy.ifServerRecordUnchanged], which
// is the default. For all other save policies, the server overwrites the
// contents of the existing record.
//
// See: https://developer.apple.com/documentation/CloudKit/CKRecord/recordID
//
// [CKModifyRecordsOperation.RecordSavePolicy.ifServerRecordUnchanged]: https://developer.apple.com/documentation/CloudKit/CKModifyRecordsOperation/RecordSavePolicy/ifServerRecordUnchanged
// [init(recordType:recordID:)]: https://developer.apple.com/documentation/CloudKit/CKRecord/init(recordType:recordID:)
func (c CKRecord) RecordID() ICKRecordID {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("recordID"))
	return CKRecordIDFromID(objc.ID(rv))
}

// The value that your app defines to identify the type of record.
//
// See: https://developer.apple.com/documentation/cloudkit/ckrecord/recordtype-6v7au
func (c CKRecord) RecordType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c.ID, objc.Sel("recordType"))
	return rv
}
func (c CKRecord) SetRecordType(value kernel.Pointer) {
	objc.Send[struct{}](c.ID, objc.Sel("setRecordType:"), value)
}

// The time when CloudKit first saves the record to the server.
//
// # Discussion
//
// The creation date reflects the time when CloudKit creates a record on the
// server with the current record’s ID. For new instances of this class, the
// value of this property is initially `nil`. When you save the record to the
// server, the value updates with the creation date for the record.
//
// See: https://developer.apple.com/documentation/CloudKit/CKRecord/creationDate
func (c CKRecord) CreationDate() foundation.NSDate {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("creationDate"))
	return foundation.NSDateFromID(objc.ID(rv))
}

// The ID of the user who creates the record.
//
// # Discussion
//
// Use this property’s value to retrieve the user record for the user who
// creates this record. Every user of the app has a unique user record that is
// empty by default. Apps can add data to the user record on behalf of the
// user, but don’t store sensitive data in it.
//
// See: https://developer.apple.com/documentation/CloudKit/CKRecord/creatorUserRecordID
func (c CKRecord) CreatorUserRecordID() ICKRecordID {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("creatorUserRecordID"))
	return CKRecordIDFromID(objc.ID(rv))
}

// The most recent time that CloudKit saved the record to the server.
//
// # Discussion
//
// The modification date reflects the most recent time that CloudKit saved a
// record with the current record’s ID to the server. For new instances of
// this class, the value of this property is initially `nil`. When you save
// the record to the server, the value updates with the modification date for
// the record.
//
// See: https://developer.apple.com/documentation/CloudKit/CKRecord/modificationDate
func (c CKRecord) ModificationDate() foundation.NSDate {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("modificationDate"))
	return foundation.NSDateFromID(objc.ID(rv))
}

// The ID of the user who most recently modified the record.
//
// # Discussion
//
// Use this property’s value to retrieve the user record of the user who
// most recently modified this record. Every user of the app has a unique user
// record that is empty by default. Apps can add data to the user record on
// behalf of the user, but don’t store sensitive data in it.
//
// See: https://developer.apple.com/documentation/CloudKit/CKRecord/lastModifiedUserRecordID
func (c CKRecord) LastModifiedUserRecordID() ICKRecordID {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("lastModifiedUserRecordID"))
	return CKRecordIDFromID(objc.ID(rv))
}

// The server change token for the record.
//
// # Discussion
//
// When you fetch a record from the server, you get the current version of
// that record as it exists on the server. However, at any time after you
// fetch a record, other users might save a newer version of it to the server.
// Every time CloudKit saves a record, the server updates the record’s
// change token to a new value. When you save your copy of the record, the
// server compares your record’s token with the token on the server. If the
// two tokens match, the server interprets that you modified the latest
// version of the record and that it can apply your changes immediately. If
// the two tokens don’t match, the server checks your app’s save policy to
// determine how to proceed.
//
// In your own code, you can use change tokens to distinguish between two
// different versions of the same record.
//
// See: https://developer.apple.com/documentation/CloudKit/CKRecord/recordChangeTag
func (c CKRecord) RecordChangeTag() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("recordChangeTag"))
	return foundation.NSStringFromID(rv).String()
}

// An object that manages the record’s encrypted key-value pairs.
//
// # Discussion
//
// Use the object this property returns to read and write encrypted key-value
// pairs that you store on the record. You can encrypt values of any data type
// that CloudKit supports, except [CKAsset], which is encrypted by default,
// and [CKReference], which isn’t encrypted so it remains available for
// server-side use. Only encrypt new fields. CloudKit doesn’t allow
// encryption on fields that already exist in your app’s schema, or on
// records that you store in the public database.
//
// CloudKit encrypts the fields’ values on-device before saving them to
// iCloud, and decrypts the values only after fetching them from the server.
// When you enable Advanced Data Protection, the encryption keys are available
// exclusively to the record’s owner and, if the user shares the record,
// that share’s participants.
//
// The following example shows how to use `encryptedValues` to encrypt and
// decrypt a string value:
//
// See: https://developer.apple.com/documentation/CloudKit/CKRecord/encryptedValues
func (c CKRecord) EncryptedValues() CKRecordKeyValueSetting {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("encryptedValues"))
	return CKRecordKeyValueSettingObjectFromID(rv)
}

// A reference to the record’s parent record.
//
// # Discussion
//
// Use parent references to inform CloudKit about the hierarchy of your
// records. CloudKit shares the hierarchy when a [CKShare] includes a
// referenced record. Add relationships between records as you create them,
// even if you don’t plan to share them. This allows you to manage the
// sharing of a hierarchy by only modifying the root record’s
// [CKRecord.Share] reference.
//
// To indicate that a record belongs to its parent, set this property to a
// reference that points to the parent record. The reference must use the
// [CKRecord.ReferenceAction.none] action or CloudKit throws an exception. The
// parent record must exist on the server when you save the child, or must be
// part of the same save operation. Otherwise, the operation fails.
//
// See: https://developer.apple.com/documentation/CloudKit/CKRecord/parent
//
// [CKRecord.ReferenceAction.none]: https://developer.apple.com/documentation/CloudKit/CKRecord/ReferenceAction/none
func (c CKRecord) Parent() ICKReference {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("parent"))
	return CKReferenceFromID(objc.ID(rv))
}
func (c CKRecord) SetParent(value ICKReference) {
	objc.Send[struct{}](c.ID, objc.Sel("setParent:"), value)
}

// A reference to the share object that determines the share status of the
// record.
//
// # Discussion
//
// CloudKit clears this property’s value when it deletes the corresponding
// [CKShare] object on the server. Send this record in the same batch
// operation as the share object you’re deleting, and this property updates
// accordingly.
//
// CloudKit only supports sharing in zones with the
// [CKRecordZoneCapabilitySharing] capability. The default zone doesn’t
// support sharing.
//
// If any records have a parent reference to this record, CloudKit implicitly
// shares them along with this record.
//
// See: https://developer.apple.com/documentation/CloudKit/CKRecord/share
func (c CKRecord) Share() ICKReference {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("share"))
	return CKReferenceFromID(objc.ID(rv))
}

// Protocol methods for CKRecordKeyValueSetting
