// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [NSScriptClassDescription] class.
var (
	_NSScriptClassDescriptionClass     NSScriptClassDescriptionClass
	_NSScriptClassDescriptionClassOnce sync.Once
)

func getNSScriptClassDescriptionClass() NSScriptClassDescriptionClass {
	_NSScriptClassDescriptionClassOnce.Do(func() {
		_NSScriptClassDescriptionClass = NSScriptClassDescriptionClass{class: objc.GetClass("NSScriptClassDescription")}
	})
	return _NSScriptClassDescriptionClass
}

// GetNSScriptClassDescriptionClass returns the class object for NSScriptClassDescription.
func GetNSScriptClassDescriptionClass() NSScriptClassDescriptionClass {
	return getNSScriptClassDescriptionClass()
}

type NSScriptClassDescriptionClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSScriptClassDescriptionClass) Alloc() NSScriptClassDescription {
	rv := objc.Send[NSScriptClassDescription](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A scriptable class that a macOS app supports.
//
// # Overview
// 
// A scriptable application provides scriptability information that describes
// the commands and objects scripters can use in scripts that target the
// application. That includes information about the classes those scriptable
// objects are created from.
// 
// An application’s scriptability information is collected automatically by
// an instance of [NSScriptSuiteRegistry]. The registry object creates an
// [NSScriptClassDescription] for each class it finds and caches these objects
// in memory. Cocoa scripting uses registry information in handling scripting
// requests that target the application.
// 
// A class description instance stores the name, attributes, relationships,
// and supported commands for a class. For example, a scriptable `document`
// class for a drawing application might support attributes such as `file` and
// `file type`, relationships such as collections of `circles`, `rectangles`,
// and `lines`, and commands such as `align` and `rotate`.
// 
// As with many of the classes in Cocoa’s built-in scripting support, your
// application may never need to directly work with instances of
// [NSScriptClassDescription]. However, one case where you might need access
// to a class description is if you override `objectSpecifier` in a scriptable
// class. For information on how to do this, see [Object Specifiers] in [Cocoa
// Scripting Guide].
// 
// Another case where your application may need access to class description
// information is if you override `` in a specifier class.
// 
// Although you can subclass [NSScriptClassDescription], it is unlikely that
// you would need to do so, or even to create instances of it.
//
// [Cocoa Scripting Guide]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/ScriptableCocoaApplications/SApps_intro/SAppsIntro.html#//apple_ref/doc/uid/TP40002164
// [Object Specifiers]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/ScriptableCocoaApplications/SApps_object_specifiers/SAppsObjectSpecifiers.html#//apple_ref/doc/uid/TP40002164-CH3
//
// # Initializing a Script Class Description
//
//   - [NSScriptClassDescription.InitWithSuiteNameClassNameDictionary]: Initializes and returns a newly allocated instance of [NSScriptClassDescription].
//
// # Getting a Script Class Description
//
//   - [NSScriptClassDescription.ClassDescriptionForKey]: Returns the class description instance for the class type of the specified attribute or relationship.
//   - [NSScriptClassDescription.SuperclassDescription]: Returns the class description instance for the superclass of the receiver’s class.
//
// # Getting basic information about the script class
//
//   - [NSScriptClassDescription.ClassName]: Returns the name of the class the receiver describes, as provided at initialization time.
//   - [NSScriptClassDescription.DefaultSubcontainerAttributeKey]: Returns the value of the [DefaultSubcontainerAttribute] entry of the class dictionary from which the receiver was instantiated.
//   - [NSScriptClassDescription.ImplementationClassName]: Returns the name of the Objective-C class instantiated to implement the scripting class.
//   - [NSScriptClassDescription.IsLocationRequiredToCreateForKey]: Returns a Boolean value indicating whether an insertion location must be specified when creating a new object in the specified to-many relationship of the receiver.
//   - [NSScriptClassDescription.SuiteName]: Returns the name of the receiver’s suite.
//
// # Getting and comparing Apple event codes
//
//   - [NSScriptClassDescription.AppleEventCode]: Returns the Apple event code associated with the receiver’s class.
//   - [NSScriptClassDescription.AppleEventCodeForKey]: Returns the Apple event code for the specified attribute or relationship in the receiver.
//   - [NSScriptClassDescription.MatchesAppleEventCode]: Returns a Boolean value indicating whether a primary or secondary Apple event code in the receiver matches the passed code.
//
// # Getting attribute and relationship information
//
//   - [NSScriptClassDescription.HasOrderedToManyRelationshipForKey]: Returns a Boolean value indicating whether the described class has an ordered to-many relationship identified by the specified key.
//   - [NSScriptClassDescription.HasPropertyForKey]: Returns a Boolean value indicating whether the described class has a property identified by the specified key.
//   - [NSScriptClassDescription.HasReadablePropertyForKey]: Returns a Boolean value indicating whether the described class has a readable property identified by the specified key.
//   - [NSScriptClassDescription.HasWritablePropertyForKey]: Returns a Boolean value indicating whether the described class has a writable property identified by the specified key.
//   - [NSScriptClassDescription.KeyWithAppleEventCode]: Given an Apple event code that identifies a property or element class, returns the key for the corresponding attribute, one-to-one relationship, or one-to-many relationship.
//   - [NSScriptClassDescription.TypeForKey]: Returns the name of the declared type of the attribute or relationship identified by the passed key.
//
// # Getting command information
//
//   - [NSScriptClassDescription.SelectorForCommand]: Returns the selector associated with the receiver for the specified command description.
//   - [NSScriptClassDescription.SupportsCommand]: Returns a Boolean value indicating whether the receiver or any superclass supports the specified command.
//
// See: https://developer.apple.com/documentation/Foundation/NSScriptClassDescription
type NSScriptClassDescription struct {
	NSClassDescription
}

// NSScriptClassDescriptionFromID constructs a [NSScriptClassDescription] from an objc.ID.
//
// A scriptable class that a macOS app supports.
func NSScriptClassDescriptionFromID(id objc.ID) NSScriptClassDescription {
	return NSScriptClassDescription{NSClassDescription: NSClassDescriptionFromID(id)}
}
// NOTE: NSScriptClassDescription adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSScriptClassDescription] class.
//
// # Initializing a Script Class Description
//
//   - [INSScriptClassDescription.InitWithSuiteNameClassNameDictionary]: Initializes and returns a newly allocated instance of [NSScriptClassDescription].
//
// # Getting a Script Class Description
//
//   - [INSScriptClassDescription.ClassDescriptionForKey]: Returns the class description instance for the class type of the specified attribute or relationship.
//   - [INSScriptClassDescription.SuperclassDescription]: Returns the class description instance for the superclass of the receiver’s class.
//
// # Getting basic information about the script class
//
//   - [INSScriptClassDescription.ClassName]: Returns the name of the class the receiver describes, as provided at initialization time.
//   - [INSScriptClassDescription.DefaultSubcontainerAttributeKey]: Returns the value of the [DefaultSubcontainerAttribute] entry of the class dictionary from which the receiver was instantiated.
//   - [INSScriptClassDescription.ImplementationClassName]: Returns the name of the Objective-C class instantiated to implement the scripting class.
//   - [INSScriptClassDescription.IsLocationRequiredToCreateForKey]: Returns a Boolean value indicating whether an insertion location must be specified when creating a new object in the specified to-many relationship of the receiver.
//   - [INSScriptClassDescription.SuiteName]: Returns the name of the receiver’s suite.
//
// # Getting and comparing Apple event codes
//
//   - [INSScriptClassDescription.AppleEventCode]: Returns the Apple event code associated with the receiver’s class.
//   - [INSScriptClassDescription.AppleEventCodeForKey]: Returns the Apple event code for the specified attribute or relationship in the receiver.
//   - [INSScriptClassDescription.MatchesAppleEventCode]: Returns a Boolean value indicating whether a primary or secondary Apple event code in the receiver matches the passed code.
//
// # Getting attribute and relationship information
//
//   - [INSScriptClassDescription.HasOrderedToManyRelationshipForKey]: Returns a Boolean value indicating whether the described class has an ordered to-many relationship identified by the specified key.
//   - [INSScriptClassDescription.HasPropertyForKey]: Returns a Boolean value indicating whether the described class has a property identified by the specified key.
//   - [INSScriptClassDescription.HasReadablePropertyForKey]: Returns a Boolean value indicating whether the described class has a readable property identified by the specified key.
//   - [INSScriptClassDescription.HasWritablePropertyForKey]: Returns a Boolean value indicating whether the described class has a writable property identified by the specified key.
//   - [INSScriptClassDescription.KeyWithAppleEventCode]: Given an Apple event code that identifies a property or element class, returns the key for the corresponding attribute, one-to-one relationship, or one-to-many relationship.
//   - [INSScriptClassDescription.TypeForKey]: Returns the name of the declared type of the attribute or relationship identified by the passed key.
//
// # Getting command information
//
//   - [INSScriptClassDescription.SelectorForCommand]: Returns the selector associated with the receiver for the specified command description.
//   - [INSScriptClassDescription.SupportsCommand]: Returns a Boolean value indicating whether the receiver or any superclass supports the specified command.
//
// See: https://developer.apple.com/documentation/Foundation/NSScriptClassDescription
type INSScriptClassDescription interface {
	INSClassDescription

	// Topic: Initializing a Script Class Description

	// Initializes and returns a newly allocated instance of [NSScriptClassDescription].
	InitWithSuiteNameClassNameDictionary(suiteName string, className string, classDeclaration INSDictionary) NSScriptClassDescription

	// Topic: Getting a Script Class Description

	// Returns the class description instance for the class type of the specified attribute or relationship.
	ClassDescriptionForKey(key string) INSScriptClassDescription
	// Returns the class description instance for the superclass of the receiver’s class.
	SuperclassDescription() INSScriptClassDescription

	// Topic: Getting basic information about the script class

	// Returns the name of the class the receiver describes, as provided at initialization time.
	ClassName() string
	// Returns the value of the [DefaultSubcontainerAttribute] entry of the class dictionary from which the receiver was instantiated.
	DefaultSubcontainerAttributeKey() string
	// Returns the name of the Objective-C class instantiated to implement the scripting class.
	ImplementationClassName() string
	// Returns a Boolean value indicating whether an insertion location must be specified when creating a new object in the specified to-many relationship of the receiver.
	IsLocationRequiredToCreateForKey(toManyRelationshipKey string) bool
	// Returns the name of the receiver’s suite.
	SuiteName() string

	// Topic: Getting and comparing Apple event codes

	// Returns the Apple event code associated with the receiver’s class.
	AppleEventCode() uint32
	// Returns the Apple event code for the specified attribute or relationship in the receiver.
	AppleEventCodeForKey(key string) uint32
	// Returns a Boolean value indicating whether a primary or secondary Apple event code in the receiver matches the passed code.
	MatchesAppleEventCode(appleEventCode uint32) bool

	// Topic: Getting attribute and relationship information

	// Returns a Boolean value indicating whether the described class has an ordered to-many relationship identified by the specified key.
	HasOrderedToManyRelationshipForKey(key string) bool
	// Returns a Boolean value indicating whether the described class has a property identified by the specified key.
	HasPropertyForKey(key string) bool
	// Returns a Boolean value indicating whether the described class has a readable property identified by the specified key.
	HasReadablePropertyForKey(key string) bool
	// Returns a Boolean value indicating whether the described class has a writable property identified by the specified key.
	HasWritablePropertyForKey(key string) bool
	// Given an Apple event code that identifies a property or element class, returns the key for the corresponding attribute, one-to-one relationship, or one-to-many relationship.
	KeyWithAppleEventCode(appleEventCode uint32) string
	// Returns the name of the declared type of the attribute or relationship identified by the passed key.
	TypeForKey(key string) string

	// Topic: Getting command information

	// Returns the selector associated with the receiver for the specified command description.
	SelectorForCommand(commandDescription INSScriptCommandDescription) objc.SEL
	// Returns a Boolean value indicating whether the receiver or any superclass supports the specified command.
	SupportsCommand(commandDescription INSScriptCommandDescription) bool
}

// Init initializes the instance.
func (s NSScriptClassDescription) Init() NSScriptClassDescription {
	rv := objc.Send[NSScriptClassDescription](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s NSScriptClassDescription) Autorelease() NSScriptClassDescription {
	rv := objc.Send[NSScriptClassDescription](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSScriptClassDescription creates a new NSScriptClassDescription instance.
func NewNSScriptClassDescription() NSScriptClassDescription {
	class := getNSScriptClassDescriptionClass()
	rv := objc.Send[NSScriptClassDescription](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Returns the class description for the specified class or, if it is not
// scriptable, for the first superclass that is.
//
// aClass: The class whose description is needed.
//
// # Return Value
// 
// The class description for the class specified by `aClass` or, if that class
// isn’t scriptable, for the class description for the first superclass that
// is. Returns `nil` if it doesn’t find a scriptable class.
//
// See: https://developer.apple.com/documentation/Foundation/NSScriptClassDescription/init(for:)
func NewScriptClassDescriptionForClass(aClass objc.Class) NSScriptClassDescription {
	rv := objc.Send[objc.ID](objc.ID(getNSScriptClassDescriptionClass().class), objc.Sel("classDescriptionForClass:"), aClass)
	return NSScriptClassDescriptionFromID(rv)
}

// Initializes and returns a newly allocated instance of
// [NSScriptClassDescription].
//
// suiteName: The name of the suite (in the application’s scriptability information)
// that the class belongs to. For example, `"AppName Suite"`.
//
// className: The name of the class that this instance describes.
//
// classDeclaration: A class declaration dictionary of the sort that is valid in script suite
// property list files. This dictionary provides information about the class
// such as its attributes and relationships.
//
// # Return Value
// 
// The initialized instance. Returns `nil` if the event code value for the
// class description itself is missing or is not an [NSString]. Also returns
// `nil` if the superclass name or any of the subdictionaries of descriptions
// are not of the right type.
//
// # Discussion
// 
// This method registers `self` with the application’s global instance of
// [NSScriptSuiteRegistry].
//
// See: https://developer.apple.com/documentation/Foundation/NSScriptClassDescription/init(suiteName:className:dictionary:)
func NewScriptClassDescriptionWithSuiteNameClassNameDictionary(suiteName string, className string, classDeclaration INSDictionary) NSScriptClassDescription {
	instance := getNSScriptClassDescriptionClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSuiteName:className:dictionary:"), objc.String(suiteName), objc.String(className), classDeclaration)
	return NSScriptClassDescriptionFromID(rv)
}

// Initializes and returns a newly allocated instance of
// [NSScriptClassDescription].
//
// suiteName: The name of the suite (in the application’s scriptability information)
// that the class belongs to. For example, `"AppName Suite"`.
//
// className: The name of the class that this instance describes.
//
// classDeclaration: A class declaration dictionary of the sort that is valid in script suite
// property list files. This dictionary provides information about the class
// such as its attributes and relationships.
//
// # Return Value
// 
// The initialized instance. Returns `nil` if the event code value for the
// class description itself is missing or is not an [NSString]. Also returns
// `nil` if the superclass name or any of the subdictionaries of descriptions
// are not of the right type.
//
// # Discussion
// 
// This method registers `self` with the application’s global instance of
// [NSScriptSuiteRegistry].
//
// See: https://developer.apple.com/documentation/Foundation/NSScriptClassDescription/init(suiteName:className:dictionary:)
func (s NSScriptClassDescription) InitWithSuiteNameClassNameDictionary(suiteName string, className string, classDeclaration INSDictionary) NSScriptClassDescription {
	rv := objc.Send[NSScriptClassDescription](s.ID, objc.Sel("initWithSuiteName:className:dictionary:"), objc.String(suiteName), objc.String(className), classDeclaration)
	return rv
}
// Returns the class description instance for the class type of the specified
// attribute or relationship.
//
// key: The identifying key for an attribute or relationship of the receiver.
//
// # Return Value
// 
// The instance of [NSScriptClassDescription] for the type of the attribute or
// relationship specified by `key`. Returns `nil` if no scriptable property
// corresponds to `key`.
//
// See: https://developer.apple.com/documentation/Foundation/NSScriptClassDescription/forKey(_:)
func (s NSScriptClassDescription) ClassDescriptionForKey(key string) INSScriptClassDescription {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("classDescriptionForKey:"), objc.String(key))
	return NSScriptClassDescriptionFromID(rv)
}
// Returns a Boolean value indicating whether an insertion location must be
// specified when creating a new object in the specified to-many relationship
// of the receiver.
//
// toManyRelationshipKey: The key for the to-many relationship that may require an insertion
// location.
//
// # Return Value
// 
// [true] if an insertion location must be specified; otherwise, [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// A script command object that creates a new object in a to-many relationship
// needs to know whether an explicitly specified insertion location is
// required. It can get this information from an instance of
// [NSScriptClassDescription]. For example, [NSMakeCommand] uses this method
// to determine whether or not a specific `make` AppleScript command must have
// an `at` parameter.
//
// See: https://developer.apple.com/documentation/Foundation/NSScriptClassDescription/isLocationRequiredToCreate(forKey:)
func (s NSScriptClassDescription) IsLocationRequiredToCreateForKey(toManyRelationshipKey string) bool {
	rv := objc.Send[bool](s.ID, objc.Sel("isLocationRequiredToCreateForKey:"), objc.String(toManyRelationshipKey))
	return rv
}
// Returns the Apple event code for the specified attribute or relationship in
// the receiver.
//
// key: The identifying key for an attribute or relationship of the receiver.
//
// # Return Value
// 
// The four-character Apple event code associated with the attribute or
// relationship identified by `key` in the receiver or, if none exists, in the
// class description for the receiver’s superclass. Returns `0` if no such
// attribute or relationship is found.
//
// See: https://developer.apple.com/documentation/Foundation/NSScriptClassDescription/appleEventCode(forKey:)
func (s NSScriptClassDescription) AppleEventCodeForKey(key string) uint32 {
	rv := objc.Send[uint32](s.ID, objc.Sel("appleEventCodeForKey:"), objc.String(key))
	return rv
}
// Returns a Boolean value indicating whether a primary or secondary Apple
// event code in the receiver matches the passed code.
//
// appleEventCode: An Apple event code to compare against the receiver’s primary or
// secondary codes.
//
// # Return Value
// 
// [true] if the receiver’s primary four-character Apple event code or any
// of its secondary codes (its synonyms) matches `code`; otherwise, [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Foundation/NSScriptClassDescription/matchesAppleEventCode(_:)
func (s NSScriptClassDescription) MatchesAppleEventCode(appleEventCode uint32) bool {
	rv := objc.Send[bool](s.ID, objc.Sel("matchesAppleEventCode:"), appleEventCode)
	return rv
}
// Returns a Boolean value indicating whether the described class has an
// ordered to-many relationship identified by the specified key.
//
// key: The identifying key for a property of the receiver.
//
// # Return Value
// 
// [true] if the described class has an ordered to-many relationship
// identified by the specified key; otherwise, [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Foundation/NSScriptClassDescription/hasOrderedToManyRelationship(forKey:)
func (s NSScriptClassDescription) HasOrderedToManyRelationshipForKey(key string) bool {
	rv := objc.Send[bool](s.ID, objc.Sel("hasOrderedToManyRelationshipForKey:"), objc.String(key))
	return rv
}
// Returns a Boolean value indicating whether the described class has a
// property identified by the specified key.
//
// key: The identifying key for a property of the receiver.
//
// # Return Value
// 
// [true] if the described class has a property identified by the specified
// key; otherwise, [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Foundation/NSScriptClassDescription/hasProperty(forKey:)
func (s NSScriptClassDescription) HasPropertyForKey(key string) bool {
	rv := objc.Send[bool](s.ID, objc.Sel("hasPropertyForKey:"), objc.String(key))
	return rv
}
// Returns a Boolean value indicating whether the described class has a
// readable property identified by the specified key.
//
// key: The identifying key for a property of the receiver.
//
// # Return Value
// 
// [true] if the described class has a readable property identified by the
// specified key; otherwise, [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// To determine if a property is read-only, invoke
// [HasWritablePropertyForKey]/
//
// See: https://developer.apple.com/documentation/Foundation/NSScriptClassDescription/hasReadableProperty(forKey:)
func (s NSScriptClassDescription) HasReadablePropertyForKey(key string) bool {
	rv := objc.Send[bool](s.ID, objc.Sel("hasReadablePropertyForKey:"), objc.String(key))
	return rv
}
// Returns a Boolean value indicating whether the described class has a
// writable property identified by the specified key.
//
// key: The identifying key for a property of the receiver.
//
// # Return Value
// 
// [true] if the described class has a writable property identified by the
// specified key; otherwise, [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Foundation/NSScriptClassDescription/hasWritableProperty(forKey:)
func (s NSScriptClassDescription) HasWritablePropertyForKey(key string) bool {
	rv := objc.Send[bool](s.ID, objc.Sel("hasWritablePropertyForKey:"), objc.String(key))
	return rv
}
// Given an Apple event code that identifies a property or element class,
// returns the key for the corresponding attribute, one-to-one relationship,
// or one-to-many relationship.
//
// appleEventCode: An Apple event code that identifies a property or element class.
//
// # Return Value
// 
// The key that corresponds to the property or element class identified by
// `appleEventCode` in the receiver or, if none exists, in a class description
// in the receiver’s superclasses. The four-character Apple event code
// associated with the attribute or relationship identified by `key` Returns
// `0` if no such attribute or relationship is found. Returns `nil` if it
// cannot find any such attribute or relationship.
//
// See: https://developer.apple.com/documentation/Foundation/NSScriptClassDescription/key(withAppleEventCode:)
func (s NSScriptClassDescription) KeyWithAppleEventCode(appleEventCode uint32) string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("keyWithAppleEventCode:"), appleEventCode)
	return NSStringFromID(rv).String()
}
// Returns the name of the declared type of the attribute or relationship
// identified by the passed key.
//
// key: The identifying key for an attribute, one-to-one relationship, or
// one-to-many relationship of the receiver.
//
// # Return Value
// 
// The name of the declared type of the attribute or relationship identified
// by `key`; for example, “NSString”. Searches in the receiver first, then
// in any superclass. Returns `nil` if no match is found.
//
// See: https://developer.apple.com/documentation/Foundation/NSScriptClassDescription/type(forKey:)
func (s NSScriptClassDescription) TypeForKey(key string) string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("typeForKey:"), objc.String(key))
	return NSStringFromID(rv).String()
}
// Returns the selector associated with the receiver for the specified command
// description.
//
// commandDescription: A description for a script command, such as `duplicate`, `make`, or `move`.
// Encapsulates the scriptability information for that command, such as its
// Objective-C selector, its argument names and types, and its return type (if
// any).
//
// # Return Value
// 
// The selector from the receiver for the command specified by
// `commandDescription`. Searches in the receiver first, then in any
// superclass. Returns [NULL] if no matching selector is found.
//
// See: https://developer.apple.com/documentation/Foundation/NSScriptClassDescription/selector(forCommand:)
func (s NSScriptClassDescription) SelectorForCommand(commandDescription INSScriptCommandDescription) objc.SEL {
	rv := objc.Send[objc.SEL](s.ID, objc.Sel("selectorForCommand:"), commandDescription)
	return rv
}
// Returns a Boolean value indicating whether the receiver or any superclass
// supports the specified command.
//
// commandDescription: A description for a script command, such as `duplicate`, `make`, or `move`.
// Encapsulates the scriptability information for that command, such as its
// Objective-C selector, its argument names and types, and its return type (if
// any).
//
// # Return Value
// 
// [true] if an the receiver or the instance of [NSScriptClassDescription] of
// any superclass of the receiver’s class lists the command described by
// `commandDesc` among its supported commands; otherwise, [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Foundation/NSScriptClassDescription/supportsCommand(_:)
func (s NSScriptClassDescription) SupportsCommand(commandDescription INSScriptCommandDescription) bool {
	rv := objc.Send[bool](s.ID, objc.Sel("supportsCommand:"), commandDescription)
	return rv
}

// Returns the class description instance for the superclass of the
// receiver’s class.
//
// # Return Value
// 
// A class description instance that describes the superclass of the
// receiver’s class. Returns `nil` if the class has no superclass.
// 
// # Discussion
// 
// The instance of [NSScriptClassDescription] that describes the superclass
// can be in the same suite as the receiver or in a different suite.
//
// See: https://developer.apple.com/documentation/Foundation/NSScriptClassDescription/superclass
func (s NSScriptClassDescription) SuperclassDescription() INSScriptClassDescription {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("superclassDescription"))
	return NSScriptClassDescriptionFromID(objc.ID(rv))
}
// Returns the name of the class the receiver describes, as provided at
// initialization time.
//
// # Return Value
// 
// A class name. This may be either the human-readable name for the
// class—that is, the name that is used in a script—or the name of the
// Objective-C class that is instantiated to implement the class. To reliably
// obtain the implementation name, use [ImplementationClassName].
//
// See: https://developer.apple.com/documentation/Foundation/NSScriptClassDescription/className
func (s NSScriptClassDescription) ClassName() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("className"))
	return NSStringFromID(rv).String()
}
// Returns the value of the [DefaultSubcontainerAttribute] entry of the class
// dictionary from which the receiver was instantiated.
//
// # Return Value
// 
// The value of the default subcontainer attribute entry. Returns `nil` if the
// there was no such entry.
//
// See: https://developer.apple.com/documentation/Foundation/NSScriptClassDescription/defaultSubcontainerAttributeKey
func (s NSScriptClassDescription) DefaultSubcontainerAttributeKey() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("defaultSubcontainerAttributeKey"))
	return NSStringFromID(rv).String()
}
// Returns the name of the Objective-C class instantiated to implement the
// scripting class.
//
// # Return Value
// 
// An Objective-C class name.
// 
// # Discussion
// 
// The name returned by the [ClassName] method for an instance of
// [NSScriptClassDescription] resulting from an sdef class declaration is the
// human-readable name for the class—that is, the name that is used in a
// script. To obtain the name of the Objective-C class instantiated to
// implement the class, use `implementationClassName`.
//
// See: https://developer.apple.com/documentation/Foundation/NSScriptClassDescription/implementationClassName
func (s NSScriptClassDescription) ImplementationClassName() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("implementationClassName"))
	return NSStringFromID(rv).String()
}
// Returns the name of the receiver’s suite.
//
// # Return Value
// 
// The receiver’s suite name. Within an application’s scriptability
// information, named suites contain related sets of information.
//
// See: https://developer.apple.com/documentation/Foundation/NSScriptClassDescription/suiteName
func (s NSScriptClassDescription) SuiteName() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("suiteName"))
	return NSStringFromID(rv).String()
}
// Returns the Apple event code associated with the receiver’s class.
//
// # Return Value
// 
// The Apple event code associated with the receiver’s class. This is the
// primary code used to identify the class in Apple events.
//
// See: https://developer.apple.com/documentation/Foundation/NSScriptClassDescription/appleEventCode
func (s NSScriptClassDescription) AppleEventCode() uint32 {
	rv := objc.Send[uint32](s.ID, objc.Sel("appleEventCode"))
	return rv
}

