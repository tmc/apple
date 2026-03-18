// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [UserDefaults] class.
var (
	_UserDefaultsClass     UserDefaultsClass
	_UserDefaultsClassOnce sync.Once
)

func getUserDefaultsClass() UserDefaultsClass {
	_UserDefaultsClassOnce.Do(func() {
		_UserDefaultsClass = UserDefaultsClass{class: objc.GetClass("NSUserDefaults")}
	})
	return _UserDefaultsClass
}

// GetUserDefaultsClass returns the class object for NSUserDefaults.
func GetUserDefaultsClass() UserDefaultsClass {
	return getUserDefaultsClass()
}

type UserDefaultsClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (uc UserDefaultsClass) Alloc() UserDefaults {
	rv := objc.Send[UserDefaults](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// An interface to the user’s defaults database, which stores system-wide
// and app-specific settings.
//
// # Overview
// 
// A [UserDefaults] object provides access to the defaults system, which is a
// persistent store for app-specific and system-wide settings. You use this
// system to store nonsensitive information, such as app-specific
// configuration details. The system also stores configuration details that
// apply to all apps, such as the current language settings for the device. In
// your code, you check values from this system and use them to dynamically
// alter your app’s appearance or behavior. The term refers to the fact that
// the stored data determines the default startup state and behavior.
// 
// To access the defaults system, obtain a [UserDefaults] object and call its
// methods to read and write values. The [StandardUserDefaults] object is a
// shared object you use to read and write your app’s standard settings. You
// can also create unique [UserDefaults] objects to manage specific sets of
// settings. For example, you can create a [UserDefaults] object that reads
// and writes settings your app shares with an app extension. Don’t subclass
// [UserDefaults].
// 
// Each item you store in a defaults object consists of a key-value pair,
// where each key is a string that you use to locate the item and each value
// is a data object. The defaults database supports the same value types found
// in property list files, including types like [Int], [Float], [Double],
// [Bool], [String], [URL], [NSNumber], [Date], [Array], and [Dictionary]. To
// include other types of objects in the defaults database, archive them to a
// [Data] object first and store that object instead. Prefer simple types over
// custom objects whenever possible.
// 
// With the exception of managed devices in educational institutions, the
// system stores defaults locally on the current device. When you write values
// to a [UserDefaults] object, the object updates its in-memory version of
// that information right away, and writes the value to disk asynchronously.
// When someone backs up their device, the system includes any persistent
// defaults databases in the backup data. Because the data is device-specific,
// you don’t use the defaults system to share data between devices. To share
// data between someone’s devices, use the [NSUbiquitousKeyValueStore]
// instead.
// 
// While your app is running, the defaults system generates notifications to
// let you know when values change. To observe changes to individual settings,
// add a [Using Key-Value Observing in Swift] to your [UserDefaults] object,
// using key names to build the path to the setting you want. To observe
// changes for all settings, register for a [UserDefaults.DidChangeMessage] or
// [didChangeNotification] with your [UserDefaults] object.
// 
// The [UserDefaults] type is thread-safe, and you can use the same object in
// multiple threads or tasks simultaneously.
// 
// # Domains and settings search paths
// 
// To integrate settings from different sources, the defaults system organizes
// them into domains. An app defines its own custom settings, but the system
// defines settings that apply to all apps. Similarly, you might choose to
// override a specific setting temporarily to test one of your app’s
// features. The defaults system provides domains for each of these cases
// along with several others.
// 
// When you request the value of a setting, the [UserDefaults] object searches
// its domains in a specific order until it finds the value you want. The
// following table lists the key domains that the defaults system supports and
// their search order. Some domains might not be present for all apps. For
// example, the managed domain is present only on administrator-managed
// devices.
// 
// [Table data omitted]
// 
// The system stores data for most persistent domains on the current device,
// and doesn’t share that data with other devices. To share settings among
// all of a person’s devices, save them using an [NSUbiquitousKeyValueStore]
// object instead.
// 
// # Settings in managed environments
// 
// If your app supports managed environments, an administrator might configure
// any managed devices with a default set of settings. For example, in a
// computer lab or classrom environment, a teacher might set default settings
// that the lessons require. Apps can’t write to managed domains, so if your
// app encounters a managed setting, disable or hide any controls that someone
// might use to change that setting’s value. To determine if a setting is
// managed, call the [ObjectIsForcedForKey] or [ObjectIsForcedForKeyInDomain]
// method of your [UserDefaults] object.
// 
// An app running on a managed device can use [NSUbiquitousKeyValueStore] to
// share small amounts of data with the person’s other devices. Use this
// store for data that your app can safely share with other instances of
// itself. For example, a textbook app might save the current page number so
// that the person can continue reading from the same place on any of their
// devices.
// 
// For more details about managing devices, see [Device Management].
// 
// # Sandbox considerations
// 
// A sandboxed app cannot access or modify the settings of another app or
// process, with the following exceptions:
// 
// - An app can modify settings for one of its app extensions. - An app can
// modify settings for an app group to which it belongs.
// 
// If you use the [AddSuiteNamed] method to add the identifier for an
// unrelated app, the method doesn’t give you access to the other app’s
// settings. Instead, the system writes changes to your app’s settings, not
// to the third-party app’s settings.
//
// [Array]: https://developer.apple.com/documentation/Swift/Array
// [Bool]: https://developer.apple.com/documentation/Swift/Bool
// [Data]: https://developer.apple.com/documentation/Foundation/Data
// [Date]: https://developer.apple.com/documentation/Foundation/Date
// [Device Management]: https://developer.apple.com/documentation/DeviceManagement
// [Dictionary]: https://developer.apple.com/documentation/Swift/Dictionary
// [Double]: https://developer.apple.com/documentation/Swift/Double
// [Float]: https://developer.apple.com/documentation/Swift/Float
// [Int]: https://developer.apple.com/documentation/Swift/Int
// [String]: https://developer.apple.com/documentation/Swift/String
// [URL]: https://developer.apple.com/documentation/Foundation/URL
// [UserDefaults.DidChangeMessage]: https://developer.apple.com/documentation/Foundation/UserDefaults/DidChangeMessage
// [Using Key-Value Observing in Swift]: https://developer.apple.com/documentation/Swift/using-key-value-observing-in-swift
// [didChangeNotification]: https://developer.apple.com/documentation/Foundation/UserDefaults/didChangeNotification
//
// # Creating a user defaults object
//
//   - [UserDefaults.InitWithSuiteName]: Creates a new defaults object and initializes it with the settings from the specified database.
//
// # Registering default settings
//
//   - [UserDefaults.RegisterDefaults]: Specifies the set of default settings and values to use as a fallback in cases where the app domain doesn’t have them.
//
// # Getting the value of a key
//
//   - [UserDefaults.BoolForKey]: Returns the Boolean value associated with the specified key.
//   - [UserDefaults.IntegerForKey]: Returns the integer value associated with the specified key.
//   - [UserDefaults.FloatForKey]: Returns the floating-point value associated with the specified key.
//   - [UserDefaults.DoubleForKey]: Returns the double value associated with the specified key.
//   - [UserDefaults.URLForKey]: Returns the URL associated with the specified key.
//   - [UserDefaults.StringForKey]: Returns the string associated with the specified key.
//   - [UserDefaults.StringArrayForKey]: Returns the array of strings associated with the specified key.
//   - [UserDefaults.DataForKey]: Returns the data object associated with the specified key.
//   - [UserDefaults.ObjectForKey]: Returns the object associated with the specified key.
//   - [UserDefaults.ArrayForKey]: Returns the array associated with the specified key.
//   - [UserDefaults.DictionaryForKey]: Returns the dictionary object associated with the specified key.
//   - [UserDefaults.DictionaryRepresentation]: Returns a dictionary with the union of all key-value pairs found from all domains.
//
// # Setting the value for a key
//
//   - [UserDefaults.SetBoolForKey]: Sets the value of the specified key to a Boolean value.
//   - [UserDefaults.SetIntegerForKey]: Sets the value of the specified key to an integer.
//   - [UserDefaults.SetFloatForKey]: Sets the value of the specified key to a floating-point number.
//   - [UserDefaults.SetDoubleForKey]: Sets the value of the specified key to a double.
//   - [UserDefaults.SetURLForKey]: Sets the value of the specified key to a URL.
//   - [UserDefaults.SetObjectForKey]: Sets the value of the specified key to a property list object.
//
// # Removing settings values
//
//   - [UserDefaults.RemoveObjectForKey]: Removes the value for the specified key from the defaults database.
//
// # Adding and removing search domains
//
//   - [UserDefaults.AddSuiteNamed]: Inserts settings for the specified domain into the search list of the current object.
//   - [UserDefaults.RemoveSuiteNamed]: Removes the specified domain from the search list of the current object.
//
// # Getting the domain names
//
//   - [UserDefaults.VolatileDomainNames]: An array of identifiers for the volatile domains associated with the current object.
//
// # Managing domain-specific values
//
//   - [UserDefaults.PersistentDomainForName]: Retrieves the settings from the specified persistent domain.
//   - [UserDefaults.SetPersistentDomainForName]: Replaces the keys and values in the specified domain with the new keys and values you supply.
//   - [UserDefaults.VolatileDomainForName]: Retrieves the settings from the specified volatile domain.
//   - [UserDefaults.SetVolatileDomainForName]: Replaces the keys and values in the specified domain with the new keys and values you supply.
//   - [UserDefaults.RemovePersistentDomainForName]: Removes the keys and values from the specified persistent domain.
//   - [UserDefaults.RemoveVolatileDomainForName]: Removes the keys and values from the specified volatile domain.
//
// # Checking for managed keys
//
//   - [UserDefaults.ObjectIsForcedForKey]: Returns a Boolean value that indicates whether an administrator provided the value for the specified key.
//   - [UserDefaults.ObjectIsForcedForKeyInDomain]: Returns a Boolean value that indicates whether an administrator provided the value for the key in the specified domain.
//
// # Deprecated
//
//   - [UserDefaults.Synchronize]: Waits for any pending asynchronous updates to the defaults database and returns; this method is unnecessary and shouldn’t be used.
//
// See: https://developer.apple.com/documentation/Foundation/UserDefaults
type UserDefaults struct {
	objectivec.Object
}

// UserDefaultsFromID constructs a [UserDefaults] from an objc.ID.
//
// An interface to the user’s defaults database, which stores system-wide
// and app-specific settings.
func UserDefaultsFromID(id objc.ID) UserDefaults {
	return NSUserDefaults{objectivec.Object{ID: id}}
}

// NSUserDefaultsFromID is an alias for [UserDefaultsFromID] for cross-framework compatibility.
func NSUserDefaultsFromID(id objc.ID) UserDefaults { return UserDefaultsFromID(id) }
// NOTE: UserDefaults adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [UserDefaults] class.
//
// # Creating a user defaults object
//
//   - [IUserDefaults.InitWithSuiteName]: Creates a new defaults object and initializes it with the settings from the specified database.
//
// # Registering default settings
//
//   - [IUserDefaults.RegisterDefaults]: Specifies the set of default settings and values to use as a fallback in cases where the app domain doesn’t have them.
//
// # Getting the value of a key
//
//   - [IUserDefaults.BoolForKey]: Returns the Boolean value associated with the specified key.
//   - [IUserDefaults.IntegerForKey]: Returns the integer value associated with the specified key.
//   - [IUserDefaults.FloatForKey]: Returns the floating-point value associated with the specified key.
//   - [IUserDefaults.DoubleForKey]: Returns the double value associated with the specified key.
//   - [IUserDefaults.URLForKey]: Returns the URL associated with the specified key.
//   - [IUserDefaults.StringForKey]: Returns the string associated with the specified key.
//   - [IUserDefaults.StringArrayForKey]: Returns the array of strings associated with the specified key.
//   - [IUserDefaults.DataForKey]: Returns the data object associated with the specified key.
//   - [IUserDefaults.ObjectForKey]: Returns the object associated with the specified key.
//   - [IUserDefaults.ArrayForKey]: Returns the array associated with the specified key.
//   - [IUserDefaults.DictionaryForKey]: Returns the dictionary object associated with the specified key.
//   - [IUserDefaults.DictionaryRepresentation]: Returns a dictionary with the union of all key-value pairs found from all domains.
//
// # Setting the value for a key
//
//   - [IUserDefaults.SetBoolForKey]: Sets the value of the specified key to a Boolean value.
//   - [IUserDefaults.SetIntegerForKey]: Sets the value of the specified key to an integer.
//   - [IUserDefaults.SetFloatForKey]: Sets the value of the specified key to a floating-point number.
//   - [IUserDefaults.SetDoubleForKey]: Sets the value of the specified key to a double.
//   - [IUserDefaults.SetURLForKey]: Sets the value of the specified key to a URL.
//   - [IUserDefaults.SetObjectForKey]: Sets the value of the specified key to a property list object.
//
// # Removing settings values
//
//   - [IUserDefaults.RemoveObjectForKey]: Removes the value for the specified key from the defaults database.
//
// # Adding and removing search domains
//
//   - [IUserDefaults.AddSuiteNamed]: Inserts settings for the specified domain into the search list of the current object.
//   - [IUserDefaults.RemoveSuiteNamed]: Removes the specified domain from the search list of the current object.
//
// # Getting the domain names
//
//   - [IUserDefaults.VolatileDomainNames]: An array of identifiers for the volatile domains associated with the current object.
//
// # Managing domain-specific values
//
//   - [IUserDefaults.PersistentDomainForName]: Retrieves the settings from the specified persistent domain.
//   - [IUserDefaults.SetPersistentDomainForName]: Replaces the keys and values in the specified domain with the new keys and values you supply.
//   - [IUserDefaults.VolatileDomainForName]: Retrieves the settings from the specified volatile domain.
//   - [IUserDefaults.SetVolatileDomainForName]: Replaces the keys and values in the specified domain with the new keys and values you supply.
//   - [IUserDefaults.RemovePersistentDomainForName]: Removes the keys and values from the specified persistent domain.
//   - [IUserDefaults.RemoveVolatileDomainForName]: Removes the keys and values from the specified volatile domain.
//
// # Checking for managed keys
//
//   - [IUserDefaults.ObjectIsForcedForKey]: Returns a Boolean value that indicates whether an administrator provided the value for the specified key.
//   - [IUserDefaults.ObjectIsForcedForKeyInDomain]: Returns a Boolean value that indicates whether an administrator provided the value for the key in the specified domain.
//
// # Deprecated
//
//   - [IUserDefaults.Synchronize]: Waits for any pending asynchronous updates to the defaults database and returns; this method is unnecessary and shouldn’t be used.
//
// See: https://developer.apple.com/documentation/Foundation/UserDefaults
type IUserDefaults interface {
	objectivec.IObject

	// Topic: Creating a user defaults object

	// Creates a new defaults object and initializes it with the settings from the specified database.
	InitWithSuiteName(suitename string) UserDefaults

	// Topic: Registering default settings

	// Specifies the set of default settings and values to use as a fallback in cases where the app domain doesn’t have them.
	RegisterDefaults(registrationDictionary INSDictionary)

	// Topic: Getting the value of a key

	// Returns the Boolean value associated with the specified key.
	BoolForKey(defaultName string) bool
	// Returns the integer value associated with the specified key.
	IntegerForKey(defaultName string) int
	// Returns the floating-point value associated with the specified key.
	FloatForKey(defaultName string) float32
	// Returns the double value associated with the specified key.
	DoubleForKey(defaultName string) float64
	// Returns the URL associated with the specified key.
	URLForKey(defaultName string) INSURL
	// Returns the string associated with the specified key.
	StringForKey(defaultName string) string
	// Returns the array of strings associated with the specified key.
	StringArrayForKey(defaultName string) []string
	// Returns the data object associated with the specified key.
	DataForKey(defaultName string) INSData
	// Returns the object associated with the specified key.
	ObjectForKey(defaultName string) objectivec.IObject
	// Returns the array associated with the specified key.
	ArrayForKey(defaultName string) INSArray
	// Returns the dictionary object associated with the specified key.
	DictionaryForKey(defaultName string) INSDictionary
	// Returns a dictionary with the union of all key-value pairs found from all domains.
	DictionaryRepresentation() INSDictionary

	// Topic: Setting the value for a key

	// Sets the value of the specified key to a Boolean value.
	SetBoolForKey(value bool, defaultName string)
	// Sets the value of the specified key to an integer.
	SetIntegerForKey(value int, defaultName string)
	// Sets the value of the specified key to a floating-point number.
	SetFloatForKey(value float32, defaultName string)
	// Sets the value of the specified key to a double.
	SetDoubleForKey(value float64, defaultName string)
	// Sets the value of the specified key to a URL.
	SetURLForKey(url INSURL, defaultName string)
	// Sets the value of the specified key to a property list object.
	SetObjectForKey(value objectivec.IObject, defaultName string)

	// Topic: Removing settings values

	// Removes the value for the specified key from the defaults database.
	RemoveObjectForKey(defaultName string)

	// Topic: Adding and removing search domains

	// Inserts settings for the specified domain into the search list of the current object.
	AddSuiteNamed(suiteName string)
	// Removes the specified domain from the search list of the current object.
	RemoveSuiteNamed(suiteName string)

	// Topic: Getting the domain names

	// An array of identifiers for the volatile domains associated with the current object.
	VolatileDomainNames() []string

	// Topic: Managing domain-specific values

	// Retrieves the settings from the specified persistent domain.
	PersistentDomainForName(domainName string) INSDictionary
	// Replaces the keys and values in the specified domain with the new keys and values you supply.
	SetPersistentDomainForName(domain INSDictionary, domainName string)
	// Retrieves the settings from the specified volatile domain.
	VolatileDomainForName(domainName string) INSDictionary
	// Replaces the keys and values in the specified domain with the new keys and values you supply.
	SetVolatileDomainForName(domain INSDictionary, domainName string)
	// Removes the keys and values from the specified persistent domain.
	RemovePersistentDomainForName(domainName string)
	// Removes the keys and values from the specified volatile domain.
	RemoveVolatileDomainForName(domainName string)

	// Topic: Checking for managed keys

	// Returns a Boolean value that indicates whether an administrator provided the value for the specified key.
	ObjectIsForcedForKey(key string) bool
	// Returns a Boolean value that indicates whether an administrator provided the value for the key in the specified domain.
	ObjectIsForcedForKeyInDomain(key string, domain string) bool

	// Topic: Deprecated

	// Waits for any pending asynchronous updates to the defaults database and returns; this method is unnecessary and shouldn’t be used.
	Synchronize() bool
}

// Init initializes the instance.
func (u UserDefaults) Init() UserDefaults {
	rv := objc.Send[UserDefaults](u.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u UserDefaults) Autorelease() UserDefaults {
	rv := objc.Send[UserDefaults](u.ID, objc.Sel("autorelease"))
	return rv
}

// NewUserDefaults creates a new UserDefaults instance.
func NewUserDefaults() UserDefaults {
	class := getUserDefaultsClass()
	rv := objc.Send[UserDefaults](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a new defaults object and initializes it with the settings from the
// specified database.
//
// suitename: The name of the app group or suite to add to the search list. To read and
// write settings for a shared app group, specify the app group identifier.
// Don’t specify the [GlobalDomain] or your app’s bundle identifier. If
// you specify `nil`, this method returns a defaults object that reads and
// writes from the current app’s settings.
//
// # Discussion
// 
// Use this method to create a defaults object that reads settings from the
// custom domain you specify. For example, you might use this method to access
// settings you share among multiple apps or between your app and an app
// extension. The returned object writes settings to the domain you specified.
// Every instance of ``UserDefaults shares the contents of the argument and
// registration domains.
// 
// The `suiteName` parameter matches the domain parameter of the corresponding
// CFPreferences APIs, except when translating between Foundation and Core
// Foundation constants. The following example shows two equivalent
// statements. For more details, see [Preferences Utilities].
// 
// Equivalent statements using NSUserDefaults and CFPreferences APIs
// 
// In macOS, specify another app’s bundle identifier to search that app’s
// settings. You can’t search another app’s settings if either app runs in
// an [App Sandbox] and you don’t have the proper entitlements.
//
// [App Sandbox]: https://developer.apple.com/documentation/Security/app-sandbox
// [Preferences Utilities]: https://developer.apple.com/documentation/CoreFoundation/preferences-utilities
//
// See: https://developer.apple.com/documentation/Foundation/UserDefaults/init(suiteName:)
func NewUserDefaultsWithSuiteName(suitename string) UserDefaults {
	instance := getUserDefaultsClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSuiteName:"), objc.String(suitename))
	return UserDefaultsFromID(rv)
}

// Creates a user defaults object initialized with the defaults for the
// specified user account.
//
// username: The name of the user account.
//
// # Return Value
// 
// An initialized [NSUserDefaults] object whose argument and registration
// domains are already set up. If the current user does not have access to the
// specified user account, this method returns `nil`.
//
// # Discussion
// 
// This method doesn’t put anything in the search list. Invoke it only if
// you’ve allocated your own [NSUserDefaults] instance instead of using the
// shared one.
// 
// You do not normally use this method to initialize an instance of
// [NSUserDefaults]. Applications used by a superuser might use this method to
// update the defaults databases for a number of users. The user who started
// the application must have appropriate access (read, write, or both) to the
// defaults database of the new user, or this method returns `nil`.
// 
// # Special Considerations
// 
// This method was never implemented to do anything except return the defaults
// for the current user.
//
// See: https://developer.apple.com/documentation/Foundation/UserDefaults/init(user:)
func NewUserDefaultsWithUser(username string) UserDefaults {
	instance := getUserDefaultsClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithUser:"), objc.String(username))
	return UserDefaultsFromID(rv)
}

// Creates a new defaults object and initializes it with the settings from the
// specified database.
//
// suitename: The name of the app group or suite to add to the search list. To read and
// write settings for a shared app group, specify the app group identifier.
// Don’t specify the [GlobalDomain] or your app’s bundle identifier. If
// you specify `nil`, this method returns a defaults object that reads and
// writes from the current app’s settings.
//
// # Discussion
// 
// Use this method to create a defaults object that reads settings from the
// custom domain you specify. For example, you might use this method to access
// settings you share among multiple apps or between your app and an app
// extension. The returned object writes settings to the domain you specified.
// Every instance of ``UserDefaults shares the contents of the argument and
// registration domains.
// 
// The `suiteName` parameter matches the domain parameter of the corresponding
// CFPreferences APIs, except when translating between Foundation and Core
// Foundation constants. The following example shows two equivalent
// statements. For more details, see [Preferences Utilities].
// 
// Equivalent statements using NSUserDefaults and CFPreferences APIs
// 
// In macOS, specify another app’s bundle identifier to search that app’s
// settings. You can’t search another app’s settings if either app runs in
// an [App Sandbox] and you don’t have the proper entitlements.
//
// [App Sandbox]: https://developer.apple.com/documentation/Security/app-sandbox
// [Preferences Utilities]: https://developer.apple.com/documentation/CoreFoundation/preferences-utilities
//
// See: https://developer.apple.com/documentation/Foundation/UserDefaults/init(suiteName:)
func (u UserDefaults) InitWithSuiteName(suitename string) UserDefaults {
	rv := objc.Send[UserDefaults](u.ID, objc.Sel("initWithSuiteName:"), objc.String(suitename))
	return rv
}

// Specifies the set of default settings and values to use as a fallback in
// cases where the app domain doesn’t have them.
//
// registrationDictionary: The dictionary of key-value pairs that contain the default values for your
// app’s settings. Build this dictionary programmatically or load it from a
// property list resource file in your app’s bundle.
//
// # Discussion
// 
// Call this method shortly after launch to specify the default values for
// your app’s settings. This method assigns the key-value pairs you provide
// to the registration domain, which is typically the last domain in the
// search list. The registration domain is volatile, so you must register the
// set of default values each time your app launches.
//
// See: https://developer.apple.com/documentation/Foundation/UserDefaults/register(defaults:)
func (u UserDefaults) RegisterDefaults(registrationDictionary INSDictionary) {
	objc.Send[objc.ID](u.ID, objc.Sel("registerDefaults:"), registrationDictionary)
}

// Returns the Boolean value associated with the specified key.
//
// defaultName: The key to retrieve from the defaults database.
//
// # Return Value
// 
// The Boolean value associated with `defaultName`, or `false` if the key
// isn’t present in the defaults database.
//
// # Discussion
// 
// This method automatically coerces certain values to their equivalent
// Boolean meanings. For example, it coerces the numbers `1` and `1.0`, and
// the strings “true”, “YES”, and “1” to the value `true`.
//
// See: https://developer.apple.com/documentation/Foundation/UserDefaults/bool(forKey:)
func (u UserDefaults) BoolForKey(defaultName string) bool {
	rv := objc.Send[bool](u.ID, objc.Sel("boolForKey:"), objc.String(defaultName))
	return rv
}

// Returns the integer value associated with the specified key.
//
// defaultName: The key to retrieve from the defaults database.
//
// # Return Value
// 
// The integer value associated with `defaultName`, or `0` if the key isn’t
// present in the defaults database.
//
// # Discussion
// 
// This method automatically coerces certain types to their equivalent integer
// values. The Boolean value `true` becomes `1` and `false` becomes `0`. A
// floating-point number becomes the greatest integer that’s less than the
// stored number –– for example, `2.67` becomes `2`. A string that
// contains a numerical value contains the equivalent integer value — for
// example, “123” becomes `123`.
//
// See: https://developer.apple.com/documentation/Foundation/UserDefaults/integer(forKey:)
func (u UserDefaults) IntegerForKey(defaultName string) int {
	rv := objc.Send[int](u.ID, objc.Sel("integerForKey:"), objc.String(defaultName))
	return rv
}

// Returns the floating-point value associated with the specified key.
//
// defaultName: The key to retrieve from the defaults database.
//
// # Return Value
// 
// The floating-point value associated with `defaultName`, or `0.0` if the key
// isn’t present in the defaults database.
//
// # Discussion
// 
// This method automatically coerces certain types to their equivalent integer
// values. The Boolean value `true` becomes `1.0` and `false` becomes `0.0`.
// An integer becomes the equivalent floating-point number –– for example,
// `2` becomes `2.0`. A string that contains a numerical value contains the
// equivalent floating-point value — for example, “123.4” becomes
// `123.4`.
//
// See: https://developer.apple.com/documentation/Foundation/UserDefaults/float(forKey:)
func (u UserDefaults) FloatForKey(defaultName string) float32 {
	rv := objc.Send[float32](u.ID, objc.Sel("floatForKey:"), objc.String(defaultName))
	return rv
}

// Returns the double value associated with the specified key.
//
// defaultName: The key to retrieve from the defaults database.
//
// # Return Value
// 
// The double value associated with `defaultName`, or `0.0` if the key isn’t
// present in the defaults database.
//
// # Discussion
// 
// This method automatically coerces certain types to their equivalent double
// values. The Boolean value `true` becomes `1.0` and `false` becomes `0.0`.
// An integer becomes the equivalent double –– for example, `2` becomes
// `2.0`. A string that contains a numerical value contains the equivalent
// double — for example, “123.4” becomes `123.4`.
//
// See: https://developer.apple.com/documentation/Foundation/UserDefaults/double(forKey:)
func (u UserDefaults) DoubleForKey(defaultName string) float64 {
	rv := objc.Send[float64](u.ID, objc.Sel("doubleForKey:"), objc.String(defaultName))
	return rv
}

// Returns the URL associated with the specified key.
//
// defaultName: The key to retrieve from the defaults database.
//
// # Return Value
// 
// The URL associated with `defaultName`, or `nil` if the key isn’t present
// in the defaults database.
//
// # Discussion
// 
// This method uses the data for the specified key to create and return a URL
// type. If the key is present but the method can’t use it to create a URL,
// this method returns `nil`. If a file URL contains a tilde (~) character in
// its path, this method replaces the tilde with an expanded path. If you
// saved a bookmark URL for the key previously, use the
// [URLByResolvingBookmarkDataOptionsRelativeToURLBookmarkDataIsStaleError]
// method to resolve the bookmark data and retrieve an equivalent file URL.
//
// See: https://developer.apple.com/documentation/Foundation/UserDefaults/url(forKey:)
func (u UserDefaults) URLForKey(defaultName string) INSURL {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("URLForKey:"), objc.String(defaultName))
	return NSURLFromID(rv)
}

// Returns the string associated with the specified key.
//
// defaultName: The key to retrieve from the defaults database.
//
// # Return Value
// 
// The string associated with `defaultName`, or `nil` if the key isn’t
// present in the defaults database.
//
// # Discussion
// 
// This method automatically coerces certain types to equivalent strings. For
// example, it converts the integer `123` to the string “123”. If the key
// is present, but the method can’t use it to create an appropriate string,
// this method returns `nil`.
// 
// This returned string is immutable, even if you originally set the key to a
// mutable string.
//
// See: https://developer.apple.com/documentation/Foundation/UserDefaults/string(forKey:)
func (u UserDefaults) StringForKey(defaultName string) string {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("stringForKey:"), objc.String(defaultName))
	return NSStringFromID(rv).String()
}

// Returns the array of strings associated with the specified key.
//
// defaultName: The key to retrieve from the defaults database.
//
// # Return Value
// 
// The array of strings associated with `defaultName`, or `nil` if the key
// isn’t present in the defaults database. This method also returns `nil` if
// the retrieved value isn’t an array, or if any item in the array isn’t a
// string.
//
// # Discussion
// 
// The returned array and its contents are immutable, even if you originally
// set the key to mutable values.
//
// See: https://developer.apple.com/documentation/Foundation/UserDefaults/stringArray(forKey:)
func (u UserDefaults) StringArrayForKey(defaultName string) []string {
	rv := objc.Send[[]objc.ID](u.ID, objc.Sel("stringArrayForKey:"), objc.String(defaultName))
	return objc.ConvertSliceToStrings(rv)
}

// Returns the data object associated with the specified key.
//
// defaultName: The key to retrieve from the defaults database.
//
// # Return Value
// 
// The data object associated with `defaultName`, or `nil` if the key isn’t
// present in the defaults database. This method also returns `nil` if the
// retrieved value isn’t a data object.
//
// # Discussion
// 
// The returned object is immutable, even if you originally set the key to a
// mutable value.
//
// See: https://developer.apple.com/documentation/Foundation/UserDefaults/data(forKey:)
func (u UserDefaults) DataForKey(defaultName string) INSData {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("dataForKey:"), objc.String(defaultName))
	return NSDataFromID(rv)
}

// Returns the object associated with the specified key.
//
// defaultName: The key to retrieve from the defaults database.
//
// # Return Value
// 
// The object associated with `defaultName`, or `nil` if the key isn’t
// present in the defaults database.
//
// # Discussion
// 
// The returned object is immutable, even if you originally set the key to a
// mutable value.
//
// See: https://developer.apple.com/documentation/Foundation/UserDefaults/object(forKey:)
func (u UserDefaults) ObjectForKey(defaultName string) objectivec.IObject {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("objectForKey:"), objc.String(defaultName))
	return objectivec.Object{ID: rv}
}

// Returns the array associated with the specified key.
//
// defaultName: The key to retrieve from the defaults database.
//
// # Return Value
// 
// The array associated with `defaultName`, or `nil` if the key isn’t
// present in the defaults database. This method also returns `nil` if the
// retrieved value isn’t an array.
//
// # Discussion
// 
// The returned array and its contents are immutable, even if you originally
// set the key to mutable values.
//
// See: https://developer.apple.com/documentation/Foundation/UserDefaults/array(forKey:)
func (u UserDefaults) ArrayForKey(defaultName string) INSArray {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("arrayForKey:"), objc.String(defaultName))
	return NSArrayFromID(rv)
}

// Returns the dictionary object associated with the specified key.
//
// defaultName: The key to retrieve from the defaults database.
//
// # Return Value
// 
// The dictionary object associated with `defaultName`, or `nil` if the key
// isn’t present in the defaults database. This method also returns `nil` if
// the retrieved value isn’t a dictionary object.
//
// # Discussion
// 
// The returned dictionary and its contents are immutable, even if you
// originally set the key to mutable values.
//
// See: https://developer.apple.com/documentation/Foundation/UserDefaults/dictionary(forKey:)
func (u UserDefaults) DictionaryForKey(defaultName string) INSDictionary {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("dictionaryForKey:"), objc.String(defaultName))
	return NSDictionaryFromID(rv)
}

// Returns a dictionary with the union of all key-value pairs found from all
// domains.
//
// # Return Value
// 
// A dictionary with the combined set of keys and values from all domains.
//
// # Discussion
// 
// Use this method to retrieve a union of the keys and values available to
// your app. The dictionary contains the data from all of the available
// domains. If multiple domains contain a value for the same key, the
// dictionary includes the value from the earliest occurrence of that key and
// discards the values in subsequent domains.
// 
// The values in the dictionary are one of the property list object types,
// such as [NSNumber], [NSString], [Data], [Date], [NSArray], or
// [NSDictionary].
//
// [Data]: https://developer.apple.com/documentation/Foundation/Data
// [Date]: https://developer.apple.com/documentation/Foundation/Date
//
// See: https://developer.apple.com/documentation/Foundation/UserDefaults/dictionaryRepresentation()
func (u UserDefaults) DictionaryRepresentation() INSDictionary {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("dictionaryRepresentation"))
	return NSDictionaryFromID(rv)
}

// Sets the value of the specified key to a Boolean value.
//
// value: The Boolean value to store in the defaults database.
//
// defaultName: The key that contains the setting’s name.
//
// # Discussion
// 
// This method places the Boolean value in an [NSNumber] type before writing
// the key and value to the defaults database. After you call this method, the
// system generates a [didChangeNotification] for registered observers.
//
// [didChangeNotification]: https://developer.apple.com/documentation/Foundation/UserDefaults/didChangeNotification
//
// See: https://developer.apple.com/documentation/Foundation/UserDefaults/set(_:forKey:)-3nn5m
func (u UserDefaults) SetBoolForKey(value bool, defaultName string) {
	objc.Send[objc.ID](u.ID, objc.Sel("setBool:forKey:"), value, objc.String(defaultName))
}

// Sets the value of the specified key to an integer.
//
// value: The integer value to store in the defaults database.
//
// defaultName: The key that contains the setting’s name.
//
// # Discussion
// 
// This method places the integer value in an [NSNumber] type before writing
// the key and value to the defaults database. After you call this method, the
// system generates a [didChangeNotification] for registered observers.
//
// [didChangeNotification]: https://developer.apple.com/documentation/Foundation/UserDefaults/didChangeNotification
//
// See: https://developer.apple.com/documentation/Foundation/UserDefaults/set(_:forKey:)-3v852
func (u UserDefaults) SetIntegerForKey(value int, defaultName string) {
	objc.Send[objc.ID](u.ID, objc.Sel("setInteger:forKey:"), value, objc.String(defaultName))
}

// Sets the value of the specified key to a floating-point number.
//
// value: The floating-point value to store in the defaults database.
//
// defaultName: The key that contains the setting’s name.
//
// # Discussion
// 
// This method places the floating-point value in an [NSNumber] type before
// writing the key and value to the defaults database. After you call this
// method, the system generates a [didChangeNotification] for registered
// observers.
//
// [didChangeNotification]: https://developer.apple.com/documentation/Foundation/UserDefaults/didChangeNotification
//
// See: https://developer.apple.com/documentation/Foundation/UserDefaults/set(_:forKey:)-1t5ec
func (u UserDefaults) SetFloatForKey(value float32, defaultName string) {
	objc.Send[objc.ID](u.ID, objc.Sel("setFloat:forKey:"), value, objc.String(defaultName))
}

// Sets the value of the specified key to a double.
//
// value: The double value to store in the defaults database.
//
// defaultName: The key that contains the setting’s name.
//
// # Discussion
// 
// This method places the double value in an [NSNumber] type before writing
// the key and value to the defaults database. After you call this method, the
// system generates a [didChangeNotification] for registered observers.
//
// [didChangeNotification]: https://developer.apple.com/documentation/Foundation/UserDefaults/didChangeNotification
//
// See: https://developer.apple.com/documentation/Foundation/UserDefaults/set(_:forKey:)-2w22f
func (u UserDefaults) SetDoubleForKey(value float64, defaultName string) {
	objc.Send[objc.ID](u.ID, objc.Sel("setDouble:forKey:"), value, objc.String(defaultName))
}

// Sets the value of the specified key to a URL.
//
// url: The value to store in the defaults database.
//
// defaultName: The key that contains the setting’s name.
//
// # Discussion
// 
// This method handles file URLs differently than other types of URLs. For
// most URLs, the method stores the URL object as the root object of a [Data]
// archive. If `url` contains a path to a file in the home directory, this
// method replaces the home directory portion of the path with a tilde (~)
// character before generating the data object. When you read the value back,
// the system expands the tilde character to the current home directory path.
// 
// If the location of a file might change, don’t use a file URL to specify
// its location. Instead, create a bookmark URL using the
// [BookmarkDataWithContentsOfURLError] method and save that URL instead.
// Bookmark URLs store additional information about the file so the system can
// locate the file later, even if the path to that file changes.
// 
// After you call this method, the system generates a [didChangeNotification]
// for registered observers.
//
// [Data]: https://developer.apple.com/documentation/Foundation/Data
// [didChangeNotification]: https://developer.apple.com/documentation/Foundation/UserDefaults/didChangeNotification
//
// See: https://developer.apple.com/documentation/Foundation/UserDefaults/set(_:forKey:)-2bqjt
func (u UserDefaults) SetURLForKey(url INSURL, defaultName string) {
	objc.Send[objc.ID](u.ID, objc.Sel("setURL:forKey:"), url, objc.String(defaultName))
}

// Sets the value of the specified key to a property list object.
//
// value: The property-list type to store in the defaults database. If you specify an
// array or dictionary type, those collections must similarly contain only
// property list types.
//
// defaultName: The key that contains the setting’s name.
//
// # Discussion
// 
// Use this method to write property list object types to the defaults store.
// To store types that aren’t property list objects, archive them to a
// [Data] object and use this method to save that data object to the defaults
// store.
// 
// After you call this method, the system generates a [didChangeNotification]
// for registered observers.
//
// [Data]: https://developer.apple.com/documentation/Foundation/Data
// [didChangeNotification]: https://developer.apple.com/documentation/Foundation/UserDefaults/didChangeNotification
//
// See: https://developer.apple.com/documentation/Foundation/UserDefaults/set(_:forKey:)-8ab6d
func (u UserDefaults) SetObjectForKey(value objectivec.IObject, defaultName string) {
	objc.Send[objc.ID](u.ID, objc.Sel("setObject:forKey:"), value, objc.String(defaultName))
}

// Removes the value for the specified key from the defaults database.
//
// defaultName: The key with the value you want to remove.
//
// # Discussion
// 
// This method removes the specified key and value from the app-specific
// settings. If your [UserDefaults] object writes to settings for an app group
// or other shared settings file, the method removes the key from that file
// instead. This method removes the key and value only from the target domain,
// and doesn’t impact values for the same key in other domains. For example,
// it doesn’t remove keys and values from the global domain.
// 
// After you remove the key, the system generates a [didChangeNotification]
// for registered observers.
//
// [didChangeNotification]: https://developer.apple.com/documentation/Foundation/UserDefaults/didChangeNotification
//
// See: https://developer.apple.com/documentation/Foundation/UserDefaults/removeObject(forKey:)
func (u UserDefaults) RemoveObjectForKey(defaultName string) {
	objc.Send[objc.ID](u.ID, objc.Sel("removeObjectForKey:"), objc.String(defaultName))
}

// Inserts settings for the specified domain into the search list of the
// current object.
//
// suiteName: The bundle identifier for the domain you want to add. You don’t need to
// specify a bundle identifier for another app. Instead, you might specify the
// app group identifier you use to share data between multiple apps or between
// your app and an app extension. Don’t specify your app’s bundle
// identifier or the [GlobalDomain] identifier in this parameter.
//
// # Discussion
// 
// This method inserts the domain for your custom suite of settings after the
// app domain and before the global domain. This arrangement causes the
// [UserDefaults] object to return your app-specific settings first, followed
// by settings from the specified suite. If you call this method multiple
// times, the [UserDefaults] object searches your suites in the order you
// added them.
// 
// This method doesn’t affect the destination for write operations. If you
// want to write settings to a custom suite, use the [InitWithSuiteName]
// initializer to construct a [UserDefaults] object specifically for that
// suite.
//
// See: https://developer.apple.com/documentation/Foundation/UserDefaults/addSuite(named:)
func (u UserDefaults) AddSuiteNamed(suiteName string) {
	objc.Send[objc.ID](u.ID, objc.Sel("addSuiteNamed:"), objc.String(suiteName))
}

// Removes the specified domain from the search list of the current object.
//
// suiteName: The bundle identifier for the domain you want to remove. Specify the same
// string you used when you added the domain.
//
// See: https://developer.apple.com/documentation/Foundation/UserDefaults/removeSuite(named:)
func (u UserDefaults) RemoveSuiteNamed(suiteName string) {
	objc.Send[objc.ID](u.ID, objc.Sel("removeSuiteNamed:"), objc.String(suiteName))
}

// Retrieves the settings from the specified persistent domain.
//
// domainName: The name of the persistent domain. Specify your app’s bundle identifier
// to retrieve any app-specific keys. Specify the [GlobalDomain] identifier to
// retrieve keys in the global domain.
//
// # Return Value
// 
// A dictionary containing the keys and values from the specified domain. If
// the domain doesn’t contain any keys, or is a volatile domain, the method
// returns `nil`.
//
// # Discussion
// 
// This method retrieves only the keys and values from the specified domain.
// It doesn’t retrieve keys from other persistent or volatile domains.
//
// See: https://developer.apple.com/documentation/Foundation/UserDefaults/persistentDomain(forName:)
func (u UserDefaults) PersistentDomainForName(domainName string) INSDictionary {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("persistentDomainForName:"), objc.String(domainName))
	return NSDictionaryFromID(rv)
}

// Replaces the keys and values in the specified domain with the new keys and
// values you supply.
//
// domain: A dictionary of keys and values to assign to the domain.
//
// domainName: The name of the domain to update. If you specify the identifier for the
// argument or registration domain, this method throws an exception.
//
// # Discussion
// 
// This method removes the existing keys from the specified domain and then
// adds the new keys you provide. After updating the keys, this method
// generates a [didChangeNotification] for registered observers.
//
// [didChangeNotification]: https://developer.apple.com/documentation/Foundation/UserDefaults/didChangeNotification
//
// See: https://developer.apple.com/documentation/Foundation/UserDefaults/setPersistentDomain(_:forName:)
func (u UserDefaults) SetPersistentDomainForName(domain INSDictionary, domainName string) {
	objc.Send[objc.ID](u.ID, objc.Sel("setPersistentDomain:forName:"), domain, objc.String(domainName))
}

// Retrieves the settings from the specified volatile domain.
//
// domainName: The name of the volatile domain. For example, specify the [ArgumentDomain]
// identifier to retrieve the command-line settings.
//
// # Return Value
// 
// A dictionary containing the keys and values from the specified domain. If
// the domain doesn’t contain any keys, or is a persistent domain, this
// method returns `nil`.
//
// # Discussion
// 
// This method retrieves only the keys and values from the specified domain.
// It doesn’t retrieve keys from other persistent or volatile domains.
//
// See: https://developer.apple.com/documentation/Foundation/UserDefaults/volatileDomain(forName:)
func (u UserDefaults) VolatileDomainForName(domainName string) INSDictionary {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("volatileDomainForName:"), objc.String(domainName))
	return NSDictionaryFromID(rv)
}

// Replaces the keys and values in the specified domain with the new keys and
// values you supply.
//
// domain: A dictionary of keys and values to assign to the domain.
//
// domainName: The name of the domain to update.
//
// # Discussion
// 
// This method removes the existing keys from the specified domain and then
// adds the new keys you provide. After updating the keys, this method
// generates a [didChangeNotification] for registered observers.
//
// [didChangeNotification]: https://developer.apple.com/documentation/Foundation/UserDefaults/didChangeNotification
//
// See: https://developer.apple.com/documentation/Foundation/UserDefaults/setVolatileDomain(_:forName:)
func (u UserDefaults) SetVolatileDomainForName(domain INSDictionary, domainName string) {
	objc.Send[objc.ID](u.ID, objc.Sel("setVolatileDomain:forName:"), domain, objc.String(domainName))
}

// Removes the keys and values from the specified persistent domain.
//
// domainName: The name of the domain to clear. If you specify the identifier for the
// argument or registration domain, this method throws an exception.
//
// # Discussion
// 
// This method removes all of the keys and values from the specified domain.
// After clearing the domain’s contents, this method generates a
// [didChangeNotification] for registered observers.
//
// [didChangeNotification]: https://developer.apple.com/documentation/Foundation/UserDefaults/didChangeNotification
//
// See: https://developer.apple.com/documentation/Foundation/UserDefaults/removePersistentDomain(forName:)
func (u UserDefaults) RemovePersistentDomainForName(domainName string) {
	objc.Send[objc.ID](u.ID, objc.Sel("removePersistentDomainForName:"), objc.String(domainName))
}

// Removes the keys and values from the specified volatile domain.
//
// domainName: The name of the domain to clear.
//
// # Discussion
// 
// This method removes all of the keys and values from the specified domain.
// After clearing the domain’s contents, this method generates a
// [didChangeNotification] for registered observers.
//
// [didChangeNotification]: https://developer.apple.com/documentation/Foundation/UserDefaults/didChangeNotification
//
// See: https://developer.apple.com/documentation/Foundation/UserDefaults/removeVolatileDomain(forName:)
func (u UserDefaults) RemoveVolatileDomainForName(domainName string) {
	objc.Send[objc.ID](u.ID, objc.Sel("removeVolatileDomainForName:"), objc.String(domainName))
}

// Returns a Boolean value that indicates whether an administrator provided
// the value for the specified key.
//
// key: The name of the key to check.
//
// # Return Value
// 
// `true` if an administrator provides a value for the key, otherwise `false`.
//
// # Discussion
// 
// Apps can’t change the value of managed keys, so use this method to
// determine if you can make changes to one of your app-specific keys. If a
// key is managed, disable any app-specific UI you use to change the value of
// that key.
//
// See: https://developer.apple.com/documentation/Foundation/UserDefaults/objectIsForced(forKey:)
func (u UserDefaults) ObjectIsForcedForKey(key string) bool {
	rv := objc.Send[bool](u.ID, objc.Sel("objectIsForcedForKey:"), objc.String(key))
	return rv
}

// Returns a Boolean value that indicates whether an administrator provided
// the value for the key in the specified domain.
//
// key: The name of the key to check.
//
// domain: The domain that contains the key.
//
// # Return Value
// 
// `true` if an administrator provides a value for the key, otherwise `false`.
//
// # Discussion
// 
// Apps can’t change the value of managed keys, so use this method to
// determine if you can make changes to a key in a specific domain. For
// example, you might use this method to check for overrides of settings
// belonging to a shared app group. If a key is managed, disable any
// app-specific UI you use to change the value of that key.
//
// See: https://developer.apple.com/documentation/Foundation/UserDefaults/objectIsForced(forKey:inDomain:)
func (u UserDefaults) ObjectIsForcedForKeyInDomain(key string, domain string) bool {
	rv := objc.Send[bool](u.ID, objc.Sel("objectIsForcedForKey:inDomain:"), objc.String(key), objc.String(domain))
	return rv
}

// Waits for any pending asynchronous updates to the defaults database and
// returns; this method is unnecessary and shouldn’t be used.
//
// # Return Value
// 
// [true] if the data was saved successfully to disk, otherwise [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Foundation/UserDefaults/synchronize()
func (u UserDefaults) Synchronize() bool {
	rv := objc.Send[bool](u.ID, objc.Sel("synchronize"))
	return rv
}

// This method has no effect and shouldn’t be used.
//
// See: https://developer.apple.com/documentation/Foundation/UserDefaults/resetStandardUserDefaults()
func (_UserDefaultsClass UserDefaultsClass) ResetStandardUserDefaults() {
	objc.Send[objc.ID](objc.ID(_UserDefaultsClass.class), objc.Sel("resetStandardUserDefaults"))
}

// An array of identifiers for the volatile domains associated with the
// current object.
//
// # Discussion
// 
// Each string in the array corresponds to one of the volatile domains this
// [UserDefaults] object searches. To get the contents of one of these
// domains, call the [VolatileDomainForName] method.
//
// See: https://developer.apple.com/documentation/Foundation/UserDefaults/volatileDomainNames
func (u UserDefaults) VolatileDomainNames() []string {
	rv := objc.Send[[]objc.ID](u.ID, objc.Sel("volatileDomainNames"))
	return objc.ConvertSliceToStrings(rv)
}

// The shared defaults object for the current app.
//
// # Return Value
// 
// The shared defaults object for the app.
// 
// # Discussion
// 
// Each app maintains a single, shared defaults object for you to use in your
// code. The first time your app retrieves the value of this property, it
// creates the shared object and caches the result. Subsequent retrieval
// attempts return the cached object.
// 
// The shared object retrieves settings from all of the standard domains. If
// you add a domain using the [AddSuiteNamed] method, the object retrieves
// values from that domain in addition to the standard ones. Custom domains
// remain in the search list until you remove them or the app exits. When you
// write settings using the shared object, it writes them to the current
// app’s settings.
//
// See: https://developer.apple.com/documentation/Foundation/UserDefaults/standard
func (_UserDefaultsClass UserDefaultsClass) StandardUserDefaults() UserDefaults {
	rv := objc.Send[objc.ID](objc.ID(_UserDefaultsClass.class), objc.Sel("standardUserDefaults"))
	return NSUserDefaultsFromID(objc.ID(rv))
}

// The identifier for the domain that contains command-line settings.
//
// See: https://developer.apple.com/documentation/Foundation/UserDefaults/argumentDomain
func (_UserDefaultsClass UserDefaultsClass) ArgumentDomain() string {
	rv := objc.Send[objc.ID](objc.ID(_UserDefaultsClass.class), objc.Sel("argumentDomain"))
	return NSStringFromID(rv).String()
}

// The identifier for the domain that contains system-specified settings for
// all apps.
//
// See: https://developer.apple.com/documentation/Foundation/UserDefaults/globalDomain
func (_UserDefaultsClass UserDefaultsClass) GlobalDomain() string {
	rv := objc.Send[objc.ID](objc.ID(_UserDefaultsClass.class), objc.Sel("globalDomain"))
	return NSStringFromID(rv).String()
}

// The identifier for the domain that contains your app’s registered default
// values.
//
// See: https://developer.apple.com/documentation/Foundation/UserDefaults/registrationDomain
func (_UserDefaultsClass UserDefaultsClass) RegistrationDomain() string {
	rv := objc.Send[objc.ID](objc.ID(_UserDefaultsClass.class), objc.Sel("registrationDomain"))
	return NSStringFromID(rv).String()
}

// Posted when ubiquitous defaults finish downloading data, either the first
// time a device is connected to an iCloud account or when a user switches
// their primary iCloud account.
//
// See: https://developer.apple.com/documentation/foundation/userdefaults/completedinitialcloudsyncnotification
func (_UserDefaultsClass UserDefaultsClass) CompletedInitialCloudSyncNotification() NSNotificationName {
	rv := objc.Send[objc.ID](objc.ID(_UserDefaultsClass.class), objc.Sel("NSUbiquitousUserDefaultsCompletedInitialSyncNotification"))
	return NSNotificationName(NSStringFromID(rv).String())
}

// Posted when the user changes the primary iCloud account.
//
// See: https://developer.apple.com/documentation/foundation/userdefaults/didchangecloudaccountsnotification
func (_UserDefaultsClass UserDefaultsClass) DidChangeCloudAccountsNotification() NSNotificationName {
	rv := objc.Send[objc.ID](objc.ID(_UserDefaultsClass.class), objc.Sel("NSUbiquitousUserDefaultsDidChangeAccountsNotification"))
	return NSNotificationName(NSStringFromID(rv).String())
}

// Posted when a cloud default is set, but no iCloud user is logged in.
//
// See: https://developer.apple.com/documentation/foundation/userdefaults/nocloudaccountnotification
func (_UserDefaultsClass UserDefaultsClass) NoCloudAccountNotification() NSNotificationName {
	rv := objc.Send[objc.ID](objc.ID(_UserDefaultsClass.class), objc.Sel("NSUbiquitousUserDefaultsNoCloudAccountNotification"))
	return NSNotificationName(NSStringFromID(rv).String())
}

