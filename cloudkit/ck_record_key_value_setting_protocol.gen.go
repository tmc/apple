// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// A protocol for managing the key-value pairs of a CloudKit record.
//
// See: https://developer.apple.com/documentation/CloudKit/CKRecordKeyValueSetting
type CKRecordKeyValueSetting interface {
	objectivec.IObject

	// Returns the object that the record stores for the specified key.
	//
	// See: https://developer.apple.com/documentation/CloudKit/CKRecordKeyValueSetting/object(forKey:)
	ObjectForKey(key CKRecordFieldKey) CKRecordValue

	// Returns the object that the record stores for the specified key.
	//
	// See: https://developer.apple.com/documentation/CloudKit/CKRecordKeyValueSetting/subscript(_:)
	ObjectForKeyedSubscript(key CKRecordFieldKey) CKRecordValue

	// Stores an object in the record using the specified key.
	//
	// See: https://developer.apple.com/documentation/CloudKit/CKRecordKeyValueSetting/setObject(_:forKey:)
	SetObjectForKey(object CKRecordValue, key CKRecordFieldKey)

	// Returns an array of the record’s keys.
	//
	// See: https://developer.apple.com/documentation/CloudKit/CKRecordKeyValueSetting/allKeys()
	AllKeys() []string

	// Returns an array of keys with recent changes to their values.
	//
	// See: https://developer.apple.com/documentation/CloudKit/CKRecordKeyValueSetting/changedKeys()
	ChangedKeys() []string
}

// CKRecordKeyValueSettingObject wraps an existing Objective-C object that conforms to the CKRecordKeyValueSetting protocol.
type CKRecordKeyValueSettingObject struct {
	objectivec.Object
}

func (o CKRecordKeyValueSettingObject) BaseObject() objectivec.Object {
	return o.Object
}

// CKRecordKeyValueSettingObjectFromID constructs a [CKRecordKeyValueSettingObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CKRecordKeyValueSettingObjectFromID(id objc.ID) CKRecordKeyValueSettingObject {
	return CKRecordKeyValueSettingObject{
		Object: objectivec.ObjectFromID(id),
	}
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
func (o CKRecordKeyValueSettingObject) ObjectForKey(key CKRecordFieldKey) CKRecordValue {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("objectForKey:"), objc.String(string(key)))
	return CKRecordValueObjectFromID(rv)
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
// See: https://developer.apple.com/documentation/CloudKit/CKRecordKeyValueSetting/subscript(_:)
func (o CKRecordKeyValueSettingObject) ObjectForKeyedSubscript(key CKRecordFieldKey) CKRecordValue {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("objectForKeyedSubscript:"), objc.String(string(key)))
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
func (o CKRecordKeyValueSettingObject) SetObjectForKey(object CKRecordValue, key CKRecordFieldKey) {
	objc.Send[struct{}](o.ID, objc.Sel("setObject:forKey:"), object, objc.String(string(key)))
}

// Returns an array of the record’s keys.
//
// # Return Value
//
// An array of keys, or an empty array if the record doesn’t contain any
// keys.
//
// See: https://developer.apple.com/documentation/CloudKit/CKRecordKeyValueSetting/allKeys()
func (o CKRecordKeyValueSettingObject) AllKeys() []string {
	rv := objc.Send[[]objc.ID](o.ID, objc.Sel("allKeys"))
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
func (o CKRecordKeyValueSettingObject) ChangedKeys() []string {
	rv := objc.Send[[]objc.ID](o.ID, objc.Sel("changedKeys"))
	return objc.ConvertSliceToStrings(rv)
}
