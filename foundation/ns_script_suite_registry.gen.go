// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSScriptSuiteRegistry] class.
var (
	_NSScriptSuiteRegistryClass     NSScriptSuiteRegistryClass
	_NSScriptSuiteRegistryClassOnce sync.Once
)

func getNSScriptSuiteRegistryClass() NSScriptSuiteRegistryClass {
	_NSScriptSuiteRegistryClassOnce.Do(func() {
		_NSScriptSuiteRegistryClass = NSScriptSuiteRegistryClass{class: objc.GetClass("NSScriptSuiteRegistry")}
	})
	return _NSScriptSuiteRegistryClass
}

// GetNSScriptSuiteRegistryClass returns the class object for NSScriptSuiteRegistry.
func GetNSScriptSuiteRegistryClass() NSScriptSuiteRegistryClass {
	return getNSScriptSuiteRegistryClass()
}

type NSScriptSuiteRegistryClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSScriptSuiteRegistryClass) Alloc() NSScriptSuiteRegistry {
	rv := objc.Send[NSScriptSuiteRegistry](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// The top-level repository of scriptability information for an app at
// runtime.
//
// # Overview
// 
// Scriptability information specifies the terminology available for use in
// scripts that target an application. It also provides information, used by
// AppleScript and by Cocoa, about how support for that terminology is
// implemented in the application. This information includes descriptions of
// the scriptable object classes in an application and of the commands the
// application supports.
// 
// There are two standard formats for supplying scriptability information: the
// older script suite format, consisting of a script suite file and one or
// more script terminology files, and the newer scripting definition (or sdef)
// format, consisting of a single sdef file.
// 
// There is one instance of [NSScriptSuiteRegistry] per scriptable
// application. This registry object collects scriptability information when
// the application first needs to respond to an Apple event for which Cocoa
// hasn’t installed a default event handler. It then creates one instance of
// [NSScriptClassDescription] for each object class and one instance of
// [NSScriptCommandDescription] for each command class, and installs a command
// handler for each command.
// 
// When a user executes an AppleScript script, Apple events are sent to the
// targeted application. Using the information stored in the registry object,
// Cocoa automatically converts incoming Apple events into script commands
// (based on [NSScriptCommand] or a subclass) that manipulate objects in the
// application.
// 
// The public methods of [NSScriptSuiteRegistry] are used primarily by
// Cocoa’s built-in scripting support. You should not need to create a
// subclass of [NSScriptSuiteRegistry].
// 
// For information on scriptability information formats, loading of
// scriptability information, and related topics, see “Scriptability
// Information” in [Overview of Cocoa Support for Scriptable Applications]
// in [Cocoa Scripting Guide].
//
// [Cocoa Scripting Guide]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/ScriptableCocoaApplications/SApps_intro/SAppsIntro.html#//apple_ref/doc/uid/TP40002164
// [Overview of Cocoa Support for Scriptable Applications]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/ScriptableCocoaApplications/SApps_about_apps/SAppsAboutApps.html#//apple_ref/doc/uid/TP40001976
//
// # Getting Suite Information
//
//   - [NSScriptSuiteRegistry.SuiteForAppleEventCode]: Returns the name of the suite definition associated with the given four-character Apple event code, `code`.
//   - [NSScriptSuiteRegistry.SuiteNames]: Returns the names of the suite definitions currently loaded by the application.
//
// # Getting and Registering Class Descriptions
//
//   - [NSScriptSuiteRegistry.ClassDescriptionsInSuite]: Returns the class descriptions contained in the suite identified by `suiteName`.
//   - [NSScriptSuiteRegistry.ClassDescriptionWithAppleEventCode]: Returns the class description associated with the given four-character Apple event code, `code`.
//   - [NSScriptSuiteRegistry.RegisterClassDescription]: Registers class description `classDescription` for use by Cocoa’s built-in scripting support by storing it in a per-suite internal dictionary under the class name.
//
// # Getting and Registering Command Descriptions
//
//   - [NSScriptSuiteRegistry.CommandDescriptionsInSuite]: Returns the command descriptions contained in the suite identified by `suiteName`.
//   - [NSScriptSuiteRegistry.CommandDescriptionWithAppleEventClassAndAppleEventCode]: Returns the command description identified by a suite’s four-character Apple event code of the class (`eventClass`) and the four-character Apple event code of the command (`commandCode`).
//   - [NSScriptSuiteRegistry.RegisterCommandDescription]: Registers command description `commandDesc` for use by Cocoa’s built-in scripting support by storing it in a per-suite internal dictionary under the command name.
//
// # Getting Other Suite Information
//
//   - [NSScriptSuiteRegistry.AeteResource]: Returns an [NSData] object that contains data in `'aete'` resource format describing the scriptability information currently known to the application.
//   - [NSScriptSuiteRegistry.AppleEventCodeForSuite]: Returns the Apple event code associated with the suite named `suiteName`, such as `‘core’` for the Core suite.
//   - [NSScriptSuiteRegistry.BundleForSuite]: Returns the bundle containing the suite-definition property list (extension `XCUIElementTypeScriptSuite`) identified by `suiteName`.
//
// # Loading Suites
//
//   - [NSScriptSuiteRegistry.LoadSuiteWithDictionaryFromBundle]: Loads the suite definition encapsulated in `dictionary`; previously, this suite definition was parsed from a `XCUIElementTypeScriptSuite` property list contained in a framework or in `bundle`.
//   - [NSScriptSuiteRegistry.LoadSuitesFromBundle]: Loads the suite definitions in bundle `aBundle`, invoking [loadSuite(with:from:)](<doc://com.apple.foundation/documentation/Foundation/NSScriptSuiteRegistry/loadSuite(with:from:)>) for each suite found.
//
// See: https://developer.apple.com/documentation/Foundation/NSScriptSuiteRegistry
type NSScriptSuiteRegistry struct {
	objectivec.Object
}

// NSScriptSuiteRegistryFromID constructs a [NSScriptSuiteRegistry] from an objc.ID.
//
// The top-level repository of scriptability information for an app at
// runtime.
func NSScriptSuiteRegistryFromID(id objc.ID) NSScriptSuiteRegistry {
	return NSScriptSuiteRegistry{objectivec.Object{ID: id}}
}
// NOTE: NSScriptSuiteRegistry adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSScriptSuiteRegistry] class.
//
// # Getting Suite Information
//
//   - [INSScriptSuiteRegistry.SuiteForAppleEventCode]: Returns the name of the suite definition associated with the given four-character Apple event code, `code`.
//   - [INSScriptSuiteRegistry.SuiteNames]: Returns the names of the suite definitions currently loaded by the application.
//
// # Getting and Registering Class Descriptions
//
//   - [INSScriptSuiteRegistry.ClassDescriptionsInSuite]: Returns the class descriptions contained in the suite identified by `suiteName`.
//   - [INSScriptSuiteRegistry.ClassDescriptionWithAppleEventCode]: Returns the class description associated with the given four-character Apple event code, `code`.
//   - [INSScriptSuiteRegistry.RegisterClassDescription]: Registers class description `classDescription` for use by Cocoa’s built-in scripting support by storing it in a per-suite internal dictionary under the class name.
//
// # Getting and Registering Command Descriptions
//
//   - [INSScriptSuiteRegistry.CommandDescriptionsInSuite]: Returns the command descriptions contained in the suite identified by `suiteName`.
//   - [INSScriptSuiteRegistry.CommandDescriptionWithAppleEventClassAndAppleEventCode]: Returns the command description identified by a suite’s four-character Apple event code of the class (`eventClass`) and the four-character Apple event code of the command (`commandCode`).
//   - [INSScriptSuiteRegistry.RegisterCommandDescription]: Registers command description `commandDesc` for use by Cocoa’s built-in scripting support by storing it in a per-suite internal dictionary under the command name.
//
// # Getting Other Suite Information
//
//   - [INSScriptSuiteRegistry.AeteResource]: Returns an [NSData] object that contains data in `'aete'` resource format describing the scriptability information currently known to the application.
//   - [INSScriptSuiteRegistry.AppleEventCodeForSuite]: Returns the Apple event code associated with the suite named `suiteName`, such as `‘core’` for the Core suite.
//   - [INSScriptSuiteRegistry.BundleForSuite]: Returns the bundle containing the suite-definition property list (extension `XCUIElementTypeScriptSuite`) identified by `suiteName`.
//
// # Loading Suites
//
//   - [INSScriptSuiteRegistry.LoadSuiteWithDictionaryFromBundle]: Loads the suite definition encapsulated in `dictionary`; previously, this suite definition was parsed from a `XCUIElementTypeScriptSuite` property list contained in a framework or in `bundle`.
//   - [INSScriptSuiteRegistry.LoadSuitesFromBundle]: Loads the suite definitions in bundle `aBundle`, invoking [loadSuite(with:from:)](<doc://com.apple.foundation/documentation/Foundation/NSScriptSuiteRegistry/loadSuite(with:from:)>) for each suite found.
//
// See: https://developer.apple.com/documentation/Foundation/NSScriptSuiteRegistry
type INSScriptSuiteRegistry interface {
	objectivec.IObject

	// Topic: Getting Suite Information

	// Returns the name of the suite definition associated with the given four-character Apple event code, `code`.
	SuiteForAppleEventCode(appleEventCode uint32) string
	// Returns the names of the suite definitions currently loaded by the application.
	SuiteNames() []string

	// Topic: Getting and Registering Class Descriptions

	// Returns the class descriptions contained in the suite identified by `suiteName`.
	ClassDescriptionsInSuite(suiteName string) INSDictionary
	// Returns the class description associated with the given four-character Apple event code, `code`.
	ClassDescriptionWithAppleEventCode(appleEventCode uint32) INSScriptClassDescription
	// Registers class description `classDescription` for use by Cocoa’s built-in scripting support by storing it in a per-suite internal dictionary under the class name.
	RegisterClassDescription(classDescription INSScriptClassDescription)

	// Topic: Getting and Registering Command Descriptions

	// Returns the command descriptions contained in the suite identified by `suiteName`.
	CommandDescriptionsInSuite(suiteName string) INSDictionary
	// Returns the command description identified by a suite’s four-character Apple event code of the class (`eventClass`) and the four-character Apple event code of the command (`commandCode`).
	CommandDescriptionWithAppleEventClassAndAppleEventCode(appleEventClassCode uint32, appleEventIDCode uint32) INSScriptCommandDescription
	// Registers command description `commandDesc` for use by Cocoa’s built-in scripting support by storing it in a per-suite internal dictionary under the command name.
	RegisterCommandDescription(commandDescription INSScriptCommandDescription)

	// Topic: Getting Other Suite Information

	// Returns an [NSData] object that contains data in `'aete'` resource format describing the scriptability information currently known to the application.
	AeteResource(languageName string) INSData
	// Returns the Apple event code associated with the suite named `suiteName`, such as `‘core’` for the Core suite.
	AppleEventCodeForSuite(suiteName string) uint32
	// Returns the bundle containing the suite-definition property list (extension `XCUIElementTypeScriptSuite`) identified by `suiteName`.
	BundleForSuite(suiteName string) INSBundle

	// Topic: Loading Suites

	// Loads the suite definition encapsulated in `dictionary`; previously, this suite definition was parsed from a `XCUIElementTypeScriptSuite` property list contained in a framework or in `bundle`.
	LoadSuiteWithDictionaryFromBundle(suiteDeclaration INSDictionary, bundle INSBundle)
	// Loads the suite definitions in bundle `aBundle`, invoking [loadSuite(with:from:)](<doc://com.apple.foundation/documentation/Foundation/NSScriptSuiteRegistry/loadSuite(with:from:)>) for each suite found.
	LoadSuitesFromBundle(bundle INSBundle)
}

// Init initializes the instance.
func (s NSScriptSuiteRegistry) Init() NSScriptSuiteRegistry {
	rv := objc.Send[NSScriptSuiteRegistry](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s NSScriptSuiteRegistry) Autorelease() NSScriptSuiteRegistry {
	rv := objc.Send[NSScriptSuiteRegistry](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSScriptSuiteRegistry creates a new NSScriptSuiteRegistry instance.
func NewNSScriptSuiteRegistry() NSScriptSuiteRegistry {
	class := getNSScriptSuiteRegistryClass()
	rv := objc.Send[NSScriptSuiteRegistry](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Returns the name of the suite definition associated with the given
// four-character Apple event code, `code`.
//
// See: https://developer.apple.com/documentation/Foundation/NSScriptSuiteRegistry/suite(forAppleEventCode:)
func (s NSScriptSuiteRegistry) SuiteForAppleEventCode(appleEventCode uint32) string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("suiteForAppleEventCode:"), appleEventCode)
	return NSStringFromID(rv).String()
}
// Returns the class descriptions contained in the suite identified by
// `suiteName`.
//
// # Discussion
// 
// Each class description (instance of [NSScriptClassDescription]) in the
// returned dictionary is identified by class name.
//
// See: https://developer.apple.com/documentation/Foundation/NSScriptSuiteRegistry/classDescriptions(inSuite:)
func (s NSScriptSuiteRegistry) ClassDescriptionsInSuite(suiteName string) INSDictionary {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("classDescriptionsInSuite:"), objc.String(suiteName))
	return NSDictionaryFromID(rv)
}
// Returns the class description associated with the given four-character
// Apple event code, `code`.
//
// # Discussion
// 
// Overriding behavior is important here. Multiple classes can have the same
// code if the classes have an uninterrupted linear inheritance from one
// another. For example, if class B is a subclass of A and class C is a
// subclass of B, and all three classes have the same four-character Apple
// event code, then this method returns the class description for class C.
//
// See: https://developer.apple.com/documentation/Foundation/NSScriptSuiteRegistry/classDescription(withAppleEventCode:)
func (s NSScriptSuiteRegistry) ClassDescriptionWithAppleEventCode(appleEventCode uint32) INSScriptClassDescription {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("classDescriptionWithAppleEventCode:"), appleEventCode)
	return NSScriptClassDescriptionFromID(rv)
}
// Registers class description `classDescription` for use by Cocoa’s
// built-in scripting support by storing it in a per-suite internal dictionary
// under the class name.
//
// See: https://developer.apple.com/documentation/Foundation/NSScriptSuiteRegistry/register(_:)-9aplw
func (s NSScriptSuiteRegistry) RegisterClassDescription(classDescription INSScriptClassDescription) {
	objc.Send[objc.ID](s.ID, objc.Sel("registerClassDescription:"), classDescription)
}
// Returns the command descriptions contained in the suite identified by
// `suiteName`.
//
// # Discussion
// 
// Each command description (instance of [NSScriptCommandDescription]) in the
// returned dictionary is identified by command name.
//
// See: https://developer.apple.com/documentation/Foundation/NSScriptSuiteRegistry/commandDescriptions(inSuite:)
func (s NSScriptSuiteRegistry) CommandDescriptionsInSuite(suiteName string) INSDictionary {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("commandDescriptionsInSuite:"), objc.String(suiteName))
	return NSDictionaryFromID(rv)
}
// Returns the command description identified by a suite’s four-character
// Apple event code of the class (`eventClass`) and the four-character Apple
// event code of the command (`commandCode`).
//
// See: https://developer.apple.com/documentation/Foundation/NSScriptSuiteRegistry/commandDescription(withAppleEventClass:andAppleEventCode:)
func (s NSScriptSuiteRegistry) CommandDescriptionWithAppleEventClassAndAppleEventCode(appleEventClassCode uint32, appleEventIDCode uint32) INSScriptCommandDescription {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("commandDescriptionWithAppleEventClass:andAppleEventCode:"), appleEventClassCode, appleEventIDCode)
	return NSScriptCommandDescriptionFromID(rv)
}
// Registers command description `commandDesc` for use by Cocoa’s built-in
// scripting support by storing it in a per-suite internal dictionary under
// the command name.
//
// # Discussion
// 
// Also registers with the single, shared instance of [NSAppleEventManager] to
// handle incoming Apple events that should be handled by the command.
//
// See: https://developer.apple.com/documentation/Foundation/NSScriptSuiteRegistry/register(_:)-5mq91
func (s NSScriptSuiteRegistry) RegisterCommandDescription(commandDescription INSScriptCommandDescription) {
	objc.Send[objc.ID](s.ID, objc.Sel("registerCommandDescription:"), commandDescription)
}
// Returns an [NSData] object that contains data in `'aete'` resource format
// describing the scriptability information currently known to the
// application.
//
// # Discussion
// 
// This method is typically invoked to implement the `get aete` Apple event
// for an application that provides scriptability information in the script
// suite format. The `languageName` argument is the name of a language for
// which a localized resource directory (such as `English.Lproj()`) exists.
// This language indication specifies the set of
// `XCUIElementTypeScriptTerminology` files to be used to generate the data.
// [NSScriptSuiteRegistry] does not create an `'aete'` resource unless this
// method is called.
//
// See: https://developer.apple.com/documentation/Foundation/NSScriptSuiteRegistry/aeteResource(_:)
func (s NSScriptSuiteRegistry) AeteResource(languageName string) INSData {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("aeteResource:"), objc.String(languageName))
	return NSDataFromID(rv)
}
// Returns the Apple event code associated with the suite named `suiteName`,
// such as `‘core’` for the Core suite.
//
// See: https://developer.apple.com/documentation/Foundation/NSScriptSuiteRegistry/appleEventCode(forSuite:)
func (s NSScriptSuiteRegistry) AppleEventCodeForSuite(suiteName string) uint32 {
	rv := objc.Send[uint32](s.ID, objc.Sel("appleEventCodeForSuite:"), objc.String(suiteName))
	return rv
}
// Returns the bundle containing the suite-definition property list (extension
// `XCUIElementTypeScriptSuite`) identified by `suiteName`.
//
// See: https://developer.apple.com/documentation/Foundation/NSScriptSuiteRegistry/bundle(forSuite:)
func (s NSScriptSuiteRegistry) BundleForSuite(suiteName string) INSBundle {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("bundleForSuite:"), objc.String(suiteName))
	return NSBundleFromID(rv)
}
// Loads the suite definition encapsulated in `dictionary`; previously, this
// suite definition was parsed from a `XCUIElementTypeScriptSuite` property
// list contained in a framework or in `bundle`.
//
// # Discussion
// 
// The method extracts information from the dictionary and caches it in
// various internal collection objects. If keys are missing or values are of
// the wrong type, it logs messages to the console. It also registers class
// descriptions and command descriptions. In registering a class description,
// it invokes the [NSClassDescription] class method
// [RegisterClassDescriptionForClass]. In registering a command description,
// it arranges for the Apple event translator to handle incoming Apple events
// that represent the defined commands.
// 
// This method is invoked when the shared instance is initialized and when
// bundles are loaded at runtime. Prior to invoking it,
// [NSScriptSuiteRegistry] creates the dictionary argument from the
// `XCUIElementTypeScriptSuite` property list. If you invoke this method in
// your code, you should try to do it before the application receives its
// first Apple event.
//
// See: https://developer.apple.com/documentation/Foundation/NSScriptSuiteRegistry/loadSuite(with:from:)
func (s NSScriptSuiteRegistry) LoadSuiteWithDictionaryFromBundle(suiteDeclaration INSDictionary, bundle INSBundle) {
	objc.Send[objc.ID](s.ID, objc.Sel("loadSuiteWithDictionary:fromBundle:"), suiteDeclaration, bundle)
}
// Loads the suite definitions in bundle `aBundle`, invoking
// [LoadSuiteWithDictionaryFromBundle] for each suite found.
//
// # Discussion
// 
// If errors occur while method is parsing a suite-definition file, the method
// logs error messages to the console.
//
// See: https://developer.apple.com/documentation/Foundation/NSScriptSuiteRegistry/loadSuites(from:)
func (s NSScriptSuiteRegistry) LoadSuitesFromBundle(bundle INSBundle) {
	objc.Send[objc.ID](s.ID, objc.Sel("loadSuitesFromBundle:"), bundle)
}

// Sets the single, shared instance of [NSScriptSuiteRegistry] to `registry`.
//
// See: https://developer.apple.com/documentation/Foundation/NSScriptSuiteRegistry/setShared(_:)
func (_NSScriptSuiteRegistryClass NSScriptSuiteRegistryClass) SetSharedScriptSuiteRegistry(registry INSScriptSuiteRegistry) {
	objc.Send[objc.ID](objc.ID(_NSScriptSuiteRegistryClass.class), objc.Sel("setSharedScriptSuiteRegistry:"), registry)
}
// Returns the single, shared instance of [NSScriptSuiteRegistry], creating it
// first if it doesn’t exist.
//
// # Discussion
// 
// If it creates an instance, and if the application provides scriptability
// information in the script suite format, the method loads suite definitions
// in all frameworks and other bundles that the application currently imports
// or includes; if information is provided in the sdef format, the method
// loads information only from the specified sdef file. If in reading
// scriptability information an exception is `raised` because of parsing
// errors, it handles the exception by printing a line of information to the
// console.
//
// See: https://developer.apple.com/documentation/Foundation/NSScriptSuiteRegistry/shared()
func (_NSScriptSuiteRegistryClass NSScriptSuiteRegistryClass) SharedScriptSuiteRegistry() NSScriptSuiteRegistry {
	rv := objc.Send[objc.ID](objc.ID(_NSScriptSuiteRegistryClass.class), objc.Sel("sharedScriptSuiteRegistry"))
	return NSScriptSuiteRegistryFromID(rv)
}

// Returns the names of the suite definitions currently loaded by the
// application.
//
// See: https://developer.apple.com/documentation/Foundation/NSScriptSuiteRegistry/suiteNames
func (s NSScriptSuiteRegistry) SuiteNames() []string {
	rv := objc.Send[[]objc.ID](s.ID, objc.Sel("suiteNames"))
	return objc.ConvertSliceToStrings(rv)
}

