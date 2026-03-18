// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSUbiquitousKeyValueStore] class.
var (
	_NSUbiquitousKeyValueStoreClass     NSUbiquitousKeyValueStoreClass
	_NSUbiquitousKeyValueStoreClassOnce sync.Once
)

func getNSUbiquitousKeyValueStoreClass() NSUbiquitousKeyValueStoreClass {
	_NSUbiquitousKeyValueStoreClassOnce.Do(func() {
		_NSUbiquitousKeyValueStoreClass = NSUbiquitousKeyValueStoreClass{class: objc.GetClass("NSUbiquitousKeyValueStore")}
	})
	return _NSUbiquitousKeyValueStoreClass
}

// GetNSUbiquitousKeyValueStoreClass returns the class object for NSUbiquitousKeyValueStore.
func GetNSUbiquitousKeyValueStoreClass() NSUbiquitousKeyValueStoreClass {
	return getNSUbiquitousKeyValueStoreClass()
}

type NSUbiquitousKeyValueStoreClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSUbiquitousKeyValueStoreClass) Alloc() NSUbiquitousKeyValueStore {
	rv := objc.Send[NSUbiquitousKeyValueStore](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// An iCloud-based container of key-value pairs you share among instances of
// your app running on a person’s devices.
//
// # Overview
// 
// Use the shared [NSUbiquitousKeyValueStore] object to store settings,
// configuration information, and app-specific data in a person’s iCloud
// account and share it among instances of your app running on all of the
// person’s devices. The object stores a dictionary of key-value pairs that
// you provide, and propagates that data to devices with the same Apple
// account. Sharing data among different devices gives you a way to coordinate
// your app’s behavior on those devices. For example, a textbook app might
// save the current page number on someone’s iPhone so that the person can
// continue reading from the same place on their other devices.
// 
// Each app has a single iCloud key-value store object, which you retrieve
// from the [NSUbiquitousKeyValueStore.DefaultStore] class property. Use this same object throughout
// your app to read and write values. Don’t subclass
// [NSUbiquitousKeyValueStore].
// 
// The keys in the iCloud key-value store identify the item and its purpose in
// your app, and the value is a data object you use to implement the
// corresponding behavior in your app. Values must be property list types such
// as [Int64], [Float], [Double], [Bool], [String], [NSNumber], [Date],
// [Array], or [Dictionary]. To include other types of objects in the
// key-value store, archive them to a [Data] object first and store that
// object instead. Prefer simple types over custom objects whenever possible.
// 
// When you write a new value, the iCloud key-value store saves it in memory
// initially and writes it to disk asynchronously later. If the device
// doesn’t have an active Apple account, the changes remain only on the
// current device. When the person signs into their account, the system
// forwards any changes to the iCloud server and reconciles the values there
// with the local ones. As you make more changes, the system keeps the local
// and server-based copies of the data synchronized, updating each one at
// appropriate times.
// 
// When a value changes on one device, iCloud forwards that change to the
// person’s other devices. If your app is running on one of those other
// devices, the system posts [didChangeExternallyNotification] to report the
// change. Register for that notification to keep all instances of your app in
// sync.
// 
// When designing the keys and values you intend to save for your app,
// consider the following size limitations:
// 
// - Your app can have no more than 1024 keys in the iCloud key-value store. -
// The total amount of available storage space for all values is 1 megabyte. -
// The maximum size for a single value is 1 megabyte. Therefore, if you
// associate 1 megabyte of data with a single key, you can’t write other
// keys to the store. - The maximum length for each key string is 128
// characters using the UTF-16 encoding. Key strings don’t count against the
// 1 megabyte quota for values.
// 
// If you exceed any of the prescribed limits during a write operation, the
// operation fails and the system doesn’t add the keys or values to the
// store. If a key string exceeds the maximum length, the system raises an
// exception. If a write operation would exceed your app’s quota, the system
// posts [didChangeExternallyNotification] notification with the change reason
// set to [UbiquitousKeyValueStoreQuotaViolationChange].
//
// [Array]: https://developer.apple.com/documentation/Swift/Array
// [Bool]: https://developer.apple.com/documentation/Swift/Bool
// [Data]: https://developer.apple.com/documentation/Foundation/Data
// [Date]: https://developer.apple.com/documentation/Foundation/Date
// [Dictionary]: https://developer.apple.com/documentation/Swift/Dictionary
// [Double]: https://developer.apple.com/documentation/Swift/Double
// [Float]: https://developer.apple.com/documentation/Swift/Float
// [Int64]: https://developer.apple.com/documentation/Swift/Int64
// [String]: https://developer.apple.com/documentation/Swift/String
// [didChangeExternallyNotification]: https://developer.apple.com/documentation/Foundation/NSUbiquitousKeyValueStore/didChangeExternallyNotification
//
// # Getting values
//
//   - [NSUbiquitousKeyValueStore.BoolForKey]: Returns the Boolean value associated with the specified key.
//   - [NSUbiquitousKeyValueStore.DoubleForKey]: Returns the double value associated with the specified key.
//   - [NSUbiquitousKeyValueStore.LongLongForKey]: Returns the 64-bit integer value associated with the specified key.
//   - [NSUbiquitousKeyValueStore.StringForKey]: Returns the string associated with the specified key.
//   - [NSUbiquitousKeyValueStore.DataForKey]: Returns the data object associated with the specified key.
//   - [NSUbiquitousKeyValueStore.ObjectForKey]: Returns the object associated with the specified key.
//   - [NSUbiquitousKeyValueStore.ArrayForKey]: Returns the array associated with the specified key.
//   - [NSUbiquitousKeyValueStore.DictionaryForKey]: Returns the dictionary object associated with the specified key.
//   - [NSUbiquitousKeyValueStore.DictionaryRepresentation]: A dictionary with all of the key-value pairs in the iCloud key-value store.
//
// # Setting values
//
//   - [NSUbiquitousKeyValueStore.SetBoolForKey]: Sets the value of the specified key to a Boolean value.
//   - [NSUbiquitousKeyValueStore.SetDoubleForKey]: Sets the value of the specified key to a double value.
//   - [NSUbiquitousKeyValueStore.SetLongLongForKey]: Sets the value of the specified key to a 64-bit integer value.
//   - [NSUbiquitousKeyValueStore.SetStringForKey]: Sets the value of the specified key to a string value.
//   - [NSUbiquitousKeyValueStore.SetDataForKey]: Sets the value of the specified key to a data object.
//   - [NSUbiquitousKeyValueStore.SetObjectForKey]: Sets the value of the specified key to a property list object.
//   - [NSUbiquitousKeyValueStore.SetArrayForKey]: Sets the value of the specified key to an array of property list objects.
//   - [NSUbiquitousKeyValueStore.SetDictionaryForKey]: Sets the value of the specified key to a dictionary of property list objects.
//
// # Synchronizing the in-memory cache with iCloud
//
//   - [NSUbiquitousKeyValueStore.Synchronize]: Synchronizes the in-memory keys and values with the ones stored in iCloud.
//
// # Removing keys
//
//   - [NSUbiquitousKeyValueStore.RemoveObjectForKey]: Removes the value for the specified key from the iCloud key-value store.
//
// # Detecting changes to values
//
//   - [NSUbiquitousKeyValueStore.NSUbiquitousKeyValueStoreChangeReasonKey]: A key that indicates the reason why the key-value store changed.
//   - [NSUbiquitousKeyValueStore.NSUbiquitousKeyValueStoreChangedKeysKey]: A key that indicates which keys changed in the iCloud key-value store.
//   - [NSUbiquitousKeyValueStore.NSUbiquitousKeyValueStoreServerChange]: A constant that indicates a value changed in iCloud.
//   - [NSUbiquitousKeyValueStore.SetNSUbiquitousKeyValueStoreServerChange]
//   - [NSUbiquitousKeyValueStore.NSUbiquitousKeyValueStoreInitialSyncChange]: A constant that indicates the initial attempt to load keys and values from iCloud is in progress.
//   - [NSUbiquitousKeyValueStore.SetNSUbiquitousKeyValueStoreInitialSyncChange]
//   - [NSUbiquitousKeyValueStore.NSUbiquitousKeyValueStoreQuotaViolationChange]: A constant that indicates an attempt to write data exceeded the quota limits.
//   - [NSUbiquitousKeyValueStore.SetNSUbiquitousKeyValueStoreQuotaViolationChange]
//   - [NSUbiquitousKeyValueStore.NSUbiquitousKeyValueStoreAccountChange]: A constant that indicates the current Apple account changed.
//   - [NSUbiquitousKeyValueStore.SetNSUbiquitousKeyValueStoreAccountChange]
//
// See: https://developer.apple.com/documentation/Foundation/NSUbiquitousKeyValueStore
type NSUbiquitousKeyValueStore struct {
	objectivec.Object
}

// NSUbiquitousKeyValueStoreFromID constructs a [NSUbiquitousKeyValueStore] from an objc.ID.
//
// An iCloud-based container of key-value pairs you share among instances of
// your app running on a person’s devices.
func NSUbiquitousKeyValueStoreFromID(id objc.ID) NSUbiquitousKeyValueStore {
	return NSUbiquitousKeyValueStore{objectivec.Object{ID: id}}
}
// NOTE: NSUbiquitousKeyValueStore adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSUbiquitousKeyValueStore] class.
//
// # Getting values
//
//   - [INSUbiquitousKeyValueStore.BoolForKey]: Returns the Boolean value associated with the specified key.
//   - [INSUbiquitousKeyValueStore.DoubleForKey]: Returns the double value associated with the specified key.
//   - [INSUbiquitousKeyValueStore.LongLongForKey]: Returns the 64-bit integer value associated with the specified key.
//   - [INSUbiquitousKeyValueStore.StringForKey]: Returns the string associated with the specified key.
//   - [INSUbiquitousKeyValueStore.DataForKey]: Returns the data object associated with the specified key.
//   - [INSUbiquitousKeyValueStore.ObjectForKey]: Returns the object associated with the specified key.
//   - [INSUbiquitousKeyValueStore.ArrayForKey]: Returns the array associated with the specified key.
//   - [INSUbiquitousKeyValueStore.DictionaryForKey]: Returns the dictionary object associated with the specified key.
//   - [INSUbiquitousKeyValueStore.DictionaryRepresentation]: A dictionary with all of the key-value pairs in the iCloud key-value store.
//
// # Setting values
//
//   - [INSUbiquitousKeyValueStore.SetBoolForKey]: Sets the value of the specified key to a Boolean value.
//   - [INSUbiquitousKeyValueStore.SetDoubleForKey]: Sets the value of the specified key to a double value.
//   - [INSUbiquitousKeyValueStore.SetLongLongForKey]: Sets the value of the specified key to a 64-bit integer value.
//   - [INSUbiquitousKeyValueStore.SetStringForKey]: Sets the value of the specified key to a string value.
//   - [INSUbiquitousKeyValueStore.SetDataForKey]: Sets the value of the specified key to a data object.
//   - [INSUbiquitousKeyValueStore.SetObjectForKey]: Sets the value of the specified key to a property list object.
//   - [INSUbiquitousKeyValueStore.SetArrayForKey]: Sets the value of the specified key to an array of property list objects.
//   - [INSUbiquitousKeyValueStore.SetDictionaryForKey]: Sets the value of the specified key to a dictionary of property list objects.
//
// # Synchronizing the in-memory cache with iCloud
//
//   - [INSUbiquitousKeyValueStore.Synchronize]: Synchronizes the in-memory keys and values with the ones stored in iCloud.
//
// # Removing keys
//
//   - [INSUbiquitousKeyValueStore.RemoveObjectForKey]: Removes the value for the specified key from the iCloud key-value store.
//
// # Detecting changes to values
//
//   - [INSUbiquitousKeyValueStore.NSUbiquitousKeyValueStoreChangeReasonKey]: A key that indicates the reason why the key-value store changed.
//   - [INSUbiquitousKeyValueStore.NSUbiquitousKeyValueStoreChangedKeysKey]: A key that indicates which keys changed in the iCloud key-value store.
//   - [INSUbiquitousKeyValueStore.NSUbiquitousKeyValueStoreServerChange]: A constant that indicates a value changed in iCloud.
//   - [INSUbiquitousKeyValueStore.SetNSUbiquitousKeyValueStoreServerChange]
//   - [INSUbiquitousKeyValueStore.NSUbiquitousKeyValueStoreInitialSyncChange]: A constant that indicates the initial attempt to load keys and values from iCloud is in progress.
//   - [INSUbiquitousKeyValueStore.SetNSUbiquitousKeyValueStoreInitialSyncChange]
//   - [INSUbiquitousKeyValueStore.NSUbiquitousKeyValueStoreQuotaViolationChange]: A constant that indicates an attempt to write data exceeded the quota limits.
//   - [INSUbiquitousKeyValueStore.SetNSUbiquitousKeyValueStoreQuotaViolationChange]
//   - [INSUbiquitousKeyValueStore.NSUbiquitousKeyValueStoreAccountChange]: A constant that indicates the current Apple account changed.
//   - [INSUbiquitousKeyValueStore.SetNSUbiquitousKeyValueStoreAccountChange]
//
// See: https://developer.apple.com/documentation/Foundation/NSUbiquitousKeyValueStore
type INSUbiquitousKeyValueStore interface {
	objectivec.IObject

	// Topic: Getting values

	// Returns the Boolean value associated with the specified key.
	BoolForKey(aKey string) bool
	// Returns the double value associated with the specified key.
	DoubleForKey(aKey string) float64
	// Returns the 64-bit integer value associated with the specified key.
	LongLongForKey(aKey string) int64
	// Returns the string associated with the specified key.
	StringForKey(aKey string) string
	// Returns the data object associated with the specified key.
	DataForKey(aKey string) INSData
	// Returns the object associated with the specified key.
	ObjectForKey(aKey string) objectivec.IObject
	// Returns the array associated with the specified key.
	ArrayForKey(aKey string) INSArray
	// Returns the dictionary object associated with the specified key.
	DictionaryForKey(aKey string) INSDictionary
	// A dictionary with all of the key-value pairs in the iCloud key-value store.
	DictionaryRepresentation() INSDictionary

	// Topic: Setting values

	// Sets the value of the specified key to a Boolean value.
	SetBoolForKey(value bool, aKey string)
	// Sets the value of the specified key to a double value.
	SetDoubleForKey(value float64, aKey string)
	// Sets the value of the specified key to a 64-bit integer value.
	SetLongLongForKey(value int64, aKey string)
	// Sets the value of the specified key to a string value.
	SetStringForKey(aString string, aKey string)
	// Sets the value of the specified key to a data object.
	SetDataForKey(aData INSData, aKey string)
	// Sets the value of the specified key to a property list object.
	SetObjectForKey(anObject objectivec.IObject, aKey string)
	// Sets the value of the specified key to an array of property list objects.
	SetArrayForKey(anArray INSArray, aKey string)
	// Sets the value of the specified key to a dictionary of property list objects.
	SetDictionaryForKey(aDictionary INSDictionary, aKey string)

	// Topic: Synchronizing the in-memory cache with iCloud

	// Synchronizes the in-memory keys and values with the ones stored in iCloud.
	Synchronize() bool

	// Topic: Removing keys

	// Removes the value for the specified key from the iCloud key-value store.
	RemoveObjectForKey(aKey string)

	// Topic: Detecting changes to values

	// A key that indicates the reason why the key-value store changed.
	NSUbiquitousKeyValueStoreChangeReasonKey() string
	// A key that indicates which keys changed in the iCloud key-value store.
	NSUbiquitousKeyValueStoreChangedKeysKey() string
	// A constant that indicates a value changed in iCloud.
	NSUbiquitousKeyValueStoreServerChange() int
	SetNSUbiquitousKeyValueStoreServerChange(value int)
	// A constant that indicates the initial attempt to load keys and values from iCloud is in progress.
	NSUbiquitousKeyValueStoreInitialSyncChange() int
	SetNSUbiquitousKeyValueStoreInitialSyncChange(value int)
	// A constant that indicates an attempt to write data exceeded the quota limits.
	NSUbiquitousKeyValueStoreQuotaViolationChange() int
	SetNSUbiquitousKeyValueStoreQuotaViolationChange(value int)
	// A constant that indicates the current Apple account changed.
	NSUbiquitousKeyValueStoreAccountChange() int
	SetNSUbiquitousKeyValueStoreAccountChange(value int)
}

// Init initializes the instance.
func (u NSUbiquitousKeyValueStore) Init() NSUbiquitousKeyValueStore {
	rv := objc.Send[NSUbiquitousKeyValueStore](u.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u NSUbiquitousKeyValueStore) Autorelease() NSUbiquitousKeyValueStore {
	rv := objc.Send[NSUbiquitousKeyValueStore](u.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSUbiquitousKeyValueStore creates a new NSUbiquitousKeyValueStore instance.
func NewNSUbiquitousKeyValueStore() NSUbiquitousKeyValueStore {
	class := getNSUbiquitousKeyValueStoreClass()
	rv := objc.Send[NSUbiquitousKeyValueStore](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Returns the Boolean value associated with the specified key.
//
// aKey: The key to retrieve from the iCloud key-value store.
//
// # Return Value
// 
// The Boolean value associated with `aKey`, or `false` if the key isn’t
// present.
//
// # Discussion
// 
// This method automatically coerces certain values to their equivalent
// Boolean meanings. For example, it coerces the numbers 1 and 1.0, and the
// strings “true”, “YES”, and “1” to the value `true`.
//
// See: https://developer.apple.com/documentation/Foundation/NSUbiquitousKeyValueStore/bool(forKey:)
func (u NSUbiquitousKeyValueStore) BoolForKey(aKey string) bool {
	rv := objc.Send[bool](u.ID, objc.Sel("boolForKey:"), objc.String(aKey))
	return rv
}

// Returns the double value associated with the specified key.
//
// aKey: The key to retrieve from the iCloud key-value store.
//
// # Return Value
// 
// The double value associated with `aKey`, or `0.0` if the key isn’t
// present.
//
// # Discussion
// 
// This method automatically coerces certain types to their equivalent double
// values. The Boolean value `true` becomes `1.0` and `false` becomes `0.0`.
// An integer becomes the equivalent double –– for example, `2` becomes
// `2.0`. A string that contains a numerical value contains the equivalent
// double — for example, “123.4” becomes `123.4`.
//
// See: https://developer.apple.com/documentation/Foundation/NSUbiquitousKeyValueStore/double(forKey:)
func (u NSUbiquitousKeyValueStore) DoubleForKey(aKey string) float64 {
	rv := objc.Send[float64](u.ID, objc.Sel("doubleForKey:"), objc.String(aKey))
	return rv
}

// Returns the 64-bit integer value associated with the specified key.
//
// aKey: The key to retrieve from the iCloud key-value store.
//
// # Return Value
// 
// The integer value associated with `aKey`, or `0` if the key isn’t
// present.
//
// # Discussion
// 
// This method automatically coerces certain types to their equivalent long
// integer values. The Boolean value `true` becomes `1` and `false` becomes
// `0`. A floating-point number becomes the greatest integer that’s less
// than the stored number –– for example, `2.67` becomes `2`. A string
// that contains a numerical value contains the equivalent integer value —
// for example, “123” becomes `123`.
//
// See: https://developer.apple.com/documentation/Foundation/NSUbiquitousKeyValueStore/longLong(forKey:)
func (u NSUbiquitousKeyValueStore) LongLongForKey(aKey string) int64 {
	rv := objc.Send[int64](u.ID, objc.Sel("longLongForKey:"), objc.String(aKey))
	return rv
}

// Returns the string associated with the specified key.
//
// aKey: The key to retrieve from the iCloud key-value store.
//
// # Return Value
// 
// The string associated with `aKey`, or `nil` if the key isn’t present.
//
// # Discussion
// 
// This method automatically coerces certain types to equivalent strings. For
// example, it converts the integer `123` to the string “123”. If the key
// is present, but the method can’t use it to create an appropriate string,
// this method returns `nil`.
//
// See: https://developer.apple.com/documentation/Foundation/NSUbiquitousKeyValueStore/string(forKey:)
func (u NSUbiquitousKeyValueStore) StringForKey(aKey string) string {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("stringForKey:"), objc.String(aKey))
	return NSStringFromID(rv).String()
}

// Returns the data object associated with the specified key.
//
// aKey: The key to retrieve from the iCloud key-value store.
//
// # Return Value
// 
// The data object associated with `aKey`, or `nil` if the key isn’t
// present. This method also returns `nil` if the retrieved value isn’t a
// data object.
//
// See: https://developer.apple.com/documentation/Foundation/NSUbiquitousKeyValueStore/data(forKey:)
func (u NSUbiquitousKeyValueStore) DataForKey(aKey string) INSData {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("dataForKey:"), objc.String(aKey))
	return NSDataFromID(rv)
}

// Returns the object associated with the specified key.
//
// aKey: The key to retrieve from the iCloud key-value store.
//
// # Return Value
// 
// The object associated with `aKey`, or `nil` if the key isn’t present.
//
// See: https://developer.apple.com/documentation/Foundation/NSUbiquitousKeyValueStore/object(forKey:)
func (u NSUbiquitousKeyValueStore) ObjectForKey(aKey string) objectivec.IObject {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("objectForKey:"), objc.String(aKey))
	return objectivec.Object{ID: rv}
}

// Returns the array associated with the specified key.
//
// aKey: The key to retrieve from the iCloud key-value store.
//
// # Return Value
// 
// The array associated with `aKey`, or `nil` if the key isn’t present. This
// method also returns `nil` if the retrieved value isn’t an array.
//
// See: https://developer.apple.com/documentation/Foundation/NSUbiquitousKeyValueStore/array(forKey:)
func (u NSUbiquitousKeyValueStore) ArrayForKey(aKey string) INSArray {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("arrayForKey:"), objc.String(aKey))
	return NSArrayFromID(rv)
}

// Returns the dictionary object associated with the specified key.
//
// aKey: The key to retrieve from the iCloud key-value store.
//
// # Return Value
// 
// The dictionary object associated with `aKey`, or `nil` if the key isn’t
// present. This method also returns `nil` if the retrieved value isn’t a
// dictionary object.
//
// See: https://developer.apple.com/documentation/Foundation/NSUbiquitousKeyValueStore/dictionary(forKey:)
func (u NSUbiquitousKeyValueStore) DictionaryForKey(aKey string) INSDictionary {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("dictionaryForKey:"), objc.String(aKey))
	return NSDictionaryFromID(rv)
}

// Sets the value of the specified key to a Boolean value.
//
// value: The Boolean value to save to the iCloud key-value store.
//
// aKey: The key to associate with the value.
//
// # Discussion
// 
// This method places the Boolean value in an [NSNumber] type before writing
// the key and value to the store.
//
// See: https://developer.apple.com/documentation/Foundation/NSUbiquitousKeyValueStore/set(_:forKey:)-8o8mq
func (u NSUbiquitousKeyValueStore) SetBoolForKey(value bool, aKey string) {
	objc.Send[objc.ID](u.ID, objc.Sel("setBool:forKey:"), value, objc.String(aKey))
}

// Sets the value of the specified key to a double value.
//
// value: The double value to save to the iCloud key-value store.
//
// aKey: The key to associate with the value.
//
// # Discussion
// 
// This method places the double value in an [NSNumber] type before writing
// the key and value to the store.
//
// See: https://developer.apple.com/documentation/Foundation/NSUbiquitousKeyValueStore/set(_:forKey:)-1xml0
func (u NSUbiquitousKeyValueStore) SetDoubleForKey(value float64, aKey string) {
	objc.Send[objc.ID](u.ID, objc.Sel("setDouble:forKey:"), value, objc.String(aKey))
}

// Sets the value of the specified key to a 64-bit integer value.
//
// value: The integer value to save to the iCloud key-value store.
//
// aKey: The key to associate with the value.
//
// # Discussion
// 
// This method places the integer value in an [NSNumber] type before writing
// the key and value to the store.
//
// See: https://developer.apple.com/documentation/Foundation/NSUbiquitousKeyValueStore/set(_:forKey:)-7tt20
func (u NSUbiquitousKeyValueStore) SetLongLongForKey(value int64, aKey string) {
	objc.Send[objc.ID](u.ID, objc.Sel("setLongLong:forKey:"), value, objc.String(aKey))
}

// Sets the value of the specified key to a string value.
//
// aString: The string value to save to the iCloud key-value store.
//
// aKey: The key to associate with the value.
//
// See: https://developer.apple.com/documentation/Foundation/NSUbiquitousKeyValueStore/set(_:forKey:)-2rlp
func (u NSUbiquitousKeyValueStore) SetStringForKey(aString string, aKey string) {
	objc.Send[objc.ID](u.ID, objc.Sel("setString:forKey:"), objc.String(aString), objc.String(aKey))
}

// Sets the value of the specified key to a data object.
//
// aData: The data object to save to the iCloud key-value store.
//
// aKey: The key to associate with the value.
//
// # Discussion
// 
// To store types that aren’t property list objects, archive them to an
// [NSData] object first and add that object to the store using this method.
// Exercise caution when saving custom objects to iCloud. Instances of your
// app on a person’s other devices must also be able to extract the objects
// and use them. Design your objects to be portable, and design new versions
// of your app to support previous versions of your custom types.
//
// See: https://developer.apple.com/documentation/Foundation/NSUbiquitousKeyValueStore/set(_:forKey:)-3ga7z
func (u NSUbiquitousKeyValueStore) SetDataForKey(aData INSData, aKey string) {
	objc.Send[objc.ID](u.ID, objc.Sel("setData:forKey:"), aData, objc.String(aKey))
}

// Sets the value of the specified key to a property list object.
//
// anObject: The property list type to save to the iCloud key-value store.
//
// aKey: The key to associate with the value.
//
// # Discussion
// 
// Use this method to write property list object types to the iCloud key-value
// store.
//
// See: https://developer.apple.com/documentation/Foundation/NSUbiquitousKeyValueStore/set(_:forKey:)-9e3de
func (u NSUbiquitousKeyValueStore) SetObjectForKey(anObject objectivec.IObject, aKey string) {
	objc.Send[objc.ID](u.ID, objc.Sel("setObject:forKey:"), anObject, objc.String(aKey))
}

// Sets the value of the specified key to an array of property list objects.
//
// anArray: The array object to save in the iCloud key-value store. The array must
// contain only property list types.
//
// aKey: The key to associate with the value.
//
// See: https://developer.apple.com/documentation/Foundation/NSUbiquitousKeyValueStore/set(_:forKey:)-40a8f
func (u NSUbiquitousKeyValueStore) SetArrayForKey(anArray INSArray, aKey string) {
	objc.Send[objc.ID](u.ID, objc.Sel("setArray:forKey:"), anArray, objc.String(aKey))
}

// Sets the value of the specified key to a dictionary of property list
// objects.
//
// aDictionary: The dictionary object to save in the iCloud key-value store. The dictionary
// must contain only property list types.
//
// aKey: The key to associate with the value.
//
// See: https://developer.apple.com/documentation/Foundation/NSUbiquitousKeyValueStore/set(_:forKey:)-9vmlm
func (u NSUbiquitousKeyValueStore) SetDictionaryForKey(aDictionary INSDictionary, aKey string) {
	objc.Send[objc.ID](u.ID, objc.Sel("setDictionary:forKey:"), aDictionary, objc.String(aKey))
}

// Synchronizes the in-memory keys and values with the ones stored in iCloud.
//
// # Return Value
// 
// `true` if the in-memory and iCloud keys are synchronized, or `false` if an
// error occurred. For example, this method returns `false` if the app
// doesn’t have the required entitlements to access the iCloud key-value
// store.
//
// # Discussion
// 
// Call this method sparingly to synchronize the in-memory copy of the keys
// and values with the version stored in iCloud. Typically, you call this
// method only at launch or when your app returns to the foreground.
// 
// Most of the time, you don’t need to call this method directly. The system
// automatically synchronizes your app’s in-memory copy of the keys and
// values with the on-disk version at appropriate times. For example, it
// synchronizes them when your app moves to the background and when iCloud
// reports a change to a key or value. When you change keys and values
// locally, the system also synchronizes your changes automatically after a
// short delay.
// 
// Don’t rely on keys and values being available on the person’s other
// devices immediately. This method doesn’t force the system to write new
// keys and values to iCloud. Instead, it notifies iCloud that new keys and
// values are available. iCloud determines the best time to retrieve those
// keys and synchronize them with the person’s other devices. Typically,
// iCloud limits updates to several times per minute.
//
// See: https://developer.apple.com/documentation/Foundation/NSUbiquitousKeyValueStore/synchronize()
func (u NSUbiquitousKeyValueStore) Synchronize() bool {
	rv := objc.Send[bool](u.ID, objc.Sel("synchronize"))
	return rv
}

// Removes the value for the specified key from the iCloud key-value store.
//
// aKey: The key with the value you want to remove.
//
// # Discussion
// 
// This method removes the specified key and value from the in-memory version
// of the store’s data. The next time the system synchronizes the data, it
// removes the key from the on-disk storage and iCloud server. If the key is
// not in the key-value store, this method does nothing.
//
// See: https://developer.apple.com/documentation/Foundation/NSUbiquitousKeyValueStore/removeObject(forKey:)
func (u NSUbiquitousKeyValueStore) RemoveObjectForKey(aKey string) {
	objc.Send[objc.ID](u.ID, objc.Sel("removeObjectForKey:"), objc.String(aKey))
}

// A dictionary with all of the key-value pairs in the iCloud key-value store.
//
// # Discussion
// 
// Getting this property retrieves the in-memory copy of the keys and values.
// If changes to the keys and values are pending, the system fetches those
// changes from iCloud and updates the dictionary before returning it. To
// ensure the dictionary contains all recent changes, call [Synchronize]
// shortly before accessing this property. All of the values in the dictionary
// are property list object types.
//
// See: https://developer.apple.com/documentation/Foundation/NSUbiquitousKeyValueStore/dictionaryRepresentation
func (u NSUbiquitousKeyValueStore) DictionaryRepresentation() INSDictionary {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("dictionaryRepresentation"))
	return NSDictionaryFromID(objc.ID(rv))
}

// A key that indicates the reason why the key-value store changed.
//
// See: https://developer.apple.com/documentation/foundation/nsubiquitouskeyvaluestorechangereasonkey
func (u NSUbiquitousKeyValueStore) NSUbiquitousKeyValueStoreChangeReasonKey() string {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("NSUbiquitousKeyValueStoreChangeReasonKey"))
	return NSStringFromID(rv).String()
}

// A key that indicates which keys changed in the iCloud key-value store.
//
// See: https://developer.apple.com/documentation/foundation/nsubiquitouskeyvaluestorechangedkeyskey
func (u NSUbiquitousKeyValueStore) NSUbiquitousKeyValueStoreChangedKeysKey() string {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("NSUbiquitousKeyValueStoreChangedKeysKey"))
	return NSStringFromID(rv).String()
}

// A constant that indicates a value changed in iCloud.
//
// See: https://developer.apple.com/documentation/foundation/nsubiquitouskeyvaluestoreserverchange
func (u NSUbiquitousKeyValueStore) NSUbiquitousKeyValueStoreServerChange() int {
	rv := objc.Send[int](u.ID, objc.Sel("NSUbiquitousKeyValueStoreServerChange"))
	return rv
}
func (u NSUbiquitousKeyValueStore) SetNSUbiquitousKeyValueStoreServerChange(value int) {
	objc.Send[struct{}](u.ID, objc.Sel("setNSUbiquitousKeyValueStoreServerChange:"), value)
}

// A constant that indicates the initial attempt to load keys and values from
// iCloud is in progress.
//
// See: https://developer.apple.com/documentation/foundation/nsubiquitouskeyvaluestoreinitialsyncchange
func (u NSUbiquitousKeyValueStore) NSUbiquitousKeyValueStoreInitialSyncChange() int {
	rv := objc.Send[int](u.ID, objc.Sel("NSUbiquitousKeyValueStoreInitialSyncChange"))
	return rv
}
func (u NSUbiquitousKeyValueStore) SetNSUbiquitousKeyValueStoreInitialSyncChange(value int) {
	objc.Send[struct{}](u.ID, objc.Sel("setNSUbiquitousKeyValueStoreInitialSyncChange:"), value)
}

// A constant that indicates an attempt to write data exceeded the quota
// limits.
//
// See: https://developer.apple.com/documentation/foundation/nsubiquitouskeyvaluestorequotaviolationchange
func (u NSUbiquitousKeyValueStore) NSUbiquitousKeyValueStoreQuotaViolationChange() int {
	rv := objc.Send[int](u.ID, objc.Sel("NSUbiquitousKeyValueStoreQuotaViolationChange"))
	return rv
}
func (u NSUbiquitousKeyValueStore) SetNSUbiquitousKeyValueStoreQuotaViolationChange(value int) {
	objc.Send[struct{}](u.ID, objc.Sel("setNSUbiquitousKeyValueStoreQuotaViolationChange:"), value)
}

// A constant that indicates the current Apple account changed.
//
// See: https://developer.apple.com/documentation/foundation/nsubiquitouskeyvaluestoreaccountchange
func (u NSUbiquitousKeyValueStore) NSUbiquitousKeyValueStoreAccountChange() int {
	rv := objc.Send[int](u.ID, objc.Sel("NSUbiquitousKeyValueStoreAccountChange"))
	return rv
}
func (u NSUbiquitousKeyValueStore) SetNSUbiquitousKeyValueStoreAccountChange(value int) {
	objc.Send[struct{}](u.ID, objc.Sel("setNSUbiquitousKeyValueStoreAccountChange:"), value)
}

// The shared iCloud key-value store object.
//
// # Discussion
// 
// Use this object to access the shared iCloud key-value store tied to your
// app and the current person. You must use this object to get and set stored
// values.
//
// See: https://developer.apple.com/documentation/Foundation/NSUbiquitousKeyValueStore/default
func (_NSUbiquitousKeyValueStoreClass NSUbiquitousKeyValueStoreClass) DefaultStore() NSUbiquitousKeyValueStore {
	rv := objc.Send[objc.ID](objc.ID(_NSUbiquitousKeyValueStoreClass.class), objc.Sel("defaultStore"))
	return NSUbiquitousKeyValueStoreFromID(objc.ID(rv))
}

