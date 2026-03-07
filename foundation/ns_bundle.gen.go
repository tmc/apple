// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"errors"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [Bundle] class.
var (
	_BundleClass     BundleClass
	_BundleClassOnce sync.Once
)

func getBundleClass() BundleClass {
	_BundleClassOnce.Do(func() {
		_BundleClass = BundleClass{class: objc.GetClass("NSBundle")}
	})
	return _BundleClass
}

// GetBundleClass returns the class object for NSBundle.
func GetBundleClass() BundleClass {
	return getBundleClass()
}

type BundleClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (bc BundleClass) Alloc() Bundle {
	rv := objc.Send[Bundle](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}







// A representation of the code and resources stored in a bundle directory on
// disk.
//
// # Overview
// 
// Apple uses bundles to represent apps, frameworks, plug-ins, and many other
// specific types of content. Bundles organize their contained resources into
// well-defined subdirectories, and bundle structures vary depending on the
// platform and the type of the bundle. By using a bundle object, you can
// access a bundle’s resources without knowing the structure of the bundle.
// The bundle object provides a single interface for locating items, taking
// into account the bundle structure, user preferences, available
// localizations, and other relevant factors.
// 
// Any executable can use a bundle object to locate resources, either inside
// an app’s bundle or in a known bundle located elsewhere. You don’t use a
// bundle object to locate files in a container directory or in other parts of
// the file system.
// 
// The general pattern for using a bundle object is as follows:
// 
// - Create a bundle object for the intended bundle directory. - Use the
// methods of the bundle object to locate or load the needed resource. - Use
// other system APIs to interact with the resource.
// 
// Some types of frequently used resources can be located and opened without a
// bundle. For example, when loading images, you store images in asset
// catalogs and load them using the [init(named:)] methods of [UIImage] or
// [NSImage]. Similarly, for string resources, you use [NSLocalizedString] to
// load individual strings instead of loading the entire
// `XCUIElementTypeStrings` file yourself.
// 
// # Finding and Opening a Bundle
// 
// Before you can locate a resource, you must first specify which bundle
// contains it. The [NSBundle] class has many constructors, but the one you
// use most often is [MainBundle]. The main bundle represents the bundle
// directory that contains the currently executing code. So for an app, the
// main bundle object gives you access to the resources that shipped with your
// app.
// 
// If your app interacts directly with plug-ins, frameworks, or other bundled
// content, you can use other methods of this class to create appropriate
// bundle objects. You can always create bundle objects from a known URL or
// path, but other methods make it easier to access bundles your app is
// already using. For example, if you link to a framework, you can use the
// [BundleForClass] method to locate the framework bundle based on a class
// defined in that framework.
// 
// In Swift, use the [bundle()] macro to insert a bundle instance appropriate
// to the current execution context, whether an app, app extension, framework,
// or Swift package.
// 
// # Locating Resources in a Bundle
// 
// You use [NSBundle] objects to obtain the location of specific resources
// inside the bundle. When looking for resources, you provide the name of the
// resource and its type at a minimum. For resources in a specific
// subdirectory, you can also specify that directory. After locating the
// resource, the bundle routines return a path string or URL that you can use
// to open the file.
// 
// Locating a single resource in a bundle
// 
// Bundle objects follow a specific search pattern when looking for resources
// on disk. Global resources—that is, resources not in a language-specific
// `XCUIElementTypeLproj` directory—are returned first, followed by region-
// and language-specific resources. This search pattern means that the bundle
// looks for resources in the following order:
// 
// - Global (nonlocalized) resources - Region-specific localized resources
// (based on the user’s region preferences) - Language-specific localized
// resources (based on the user’s language preferences) - Development
// language resources (as specified by the [CFBundleDevelopmentRegion] key in
// the bundle’s Info.plist file)
// 
// Because global resources take precedence over language-specific resources,
// you should never include both a global and localized version of a given
// resource in your app. When a global version of a resource exists,
// language-specific versions are never returned. The reason for this
// precedence is performance. If localized resources were searched first, the
// bundle object might waste time searching for a nonexistent localized
// resource before returning the global resource.
// 
// When locating resource files, the bundle object automatically considers
// many standard filename modifiers when determining which file to return.
// Resources may be tagged for a specific device (`~iphone`, `~ipad`) or for a
// specific screen resolution (`@2x`, `@3x`). Do not include these modifiers
// when specifying the name of the resource you want. The bundle object
// selects the file that is most appropriate for the underlying device. For
// more information, see [App Icons on iPhone, iPad and Apple Watch].
// 
// # Understanding Bundle Structures
// 
// Bundle structures vary depending on the target platform and the type of
// bundle you are building. The [NSBundle] class hides this underlying
// structure in most (but not all) cases. Many of the methods you use to load
// resources from a bundle automatically locate the appropriate starting
// directory and look for resources in known places. You can also use the
// methods and properties of this class to get the location of known bundle
// directories and to retrieve resources specifically from those directories.
// 
// For information about the bundle structure of iOS and macOS apps, see
// [Bundle Programming Guide]. For information about the structure of
// framework bundles, see [Framework Programming Guide]. For information about
// the structure of macOS plug-ins, see [Code Loading Programming Topics].
//
// [App Icons on iPhone, iPad and Apple Watch]: https://developer.apple.com/library/archive/qa/qa1686/_index.html#//apple_ref/doc/uid/DTS40009882
// [Bundle Programming Guide]: https://developer.apple.com/library/archive/documentation/CoreFoundation/Conceptual/CFBundles/Introduction/Introduction.html#//apple_ref/doc/uid/10000123i
// [CFBundleDevelopmentRegion]: https://developer.apple.com/library/archive/documentation/General/Reference/InfoPlistKeyReference/Articles/CoreFoundationKeys.html#//apple_ref/doc/uid/20001431-130430
// [Code Loading Programming Topics]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/LoadingCode/LoadingCode.html#//apple_ref/doc/uid/10000052i
// [Framework Programming Guide]: https://developer.apple.com/library/archive/documentation/MacOSX/Conceptual/BPFrameworks/Frameworks.html#//apple_ref/doc/uid/10000183i
// [NSImage]: https://developer.apple.com/documentation/AppKit/NSImage
// [NSLocalizedString]: https://developer.apple.com/documentation/Foundation/NSLocalizedString
// [UIImage]: https://developer.apple.com/documentation/UIKit/UIImage
// [bundle()]: https://developer.apple.com/documentation/Foundation/bundle()
// [init(named:)]: https://developer.apple.com/documentation/UIKit/UIImage/init(named:)
//
// # Creating and Initializing a Bundle
//
//   - [Bundle.InitWithURL]: Returns an [NSBundle] object initialized to correspond to the specified file URL.
//   - [Bundle.InitWithPath]: Returns an [NSBundle] object initialized to correspond to the specified directory.
//
// # Loading Nib Files
//
//   - [Bundle.LoadNibNamedOwnerTopLevelObjects]: Loads a nib from the bundle with the specified file name and owner.
//
// # Finding Resource Files
//
//   - [Bundle.URLForResourceWithExtensionSubdirectory]: Returns the file URL for the resource file identified by the specified name and extension and residing in a given bundle directory.
//   - [Bundle.URLForResourceWithExtension]: Returns the file URL for the resource identified by the specified name and file extension.
//   - [Bundle.URLsForResourcesWithExtensionSubdirectory]: Returns an array of file URLs for all resources identified by the specified file extension and located in the specified bundle subdirectory.
//   - [Bundle.URLForResourceWithExtensionSubdirectoryLocalization]: Returns the file URL for the resource identified by the specified name and file extension, located in the specified bundle subdirectory, and limited to global resources and those associated with the specified localization.
//   - [Bundle.URLsForResourcesWithExtensionSubdirectoryLocalization]: Returns an array containing the file URLs for all bundle resources having the specified filename extension, residing in the specified resource subdirectory, and limited to global resources and those associated with the specified localization.
//   - [Bundle.PathForResourceOfType]: Returns the full pathname for the resource identified by the specified name and file extension.
//   - [Bundle.PathForResourceOfTypeInDirectory]: Returns the full pathname for the resource identified by the specified name and file extension and located in the specified bundle subdirectory.
//   - [Bundle.PathForResourceOfTypeInDirectoryForLocalization]: Returns the full pathname for the resource identified by the specified name and file extension, located in the specified bundle subdirectory, and limited to global resources and those associated with the specified localization.
//   - [Bundle.PathsForResourcesOfTypeInDirectory]: Returns an array containing the pathnames for all bundle resources having the specified filename extension and residing in the resource subdirectory.
//   - [Bundle.PathsForResourcesOfTypeInDirectoryForLocalization]: Returns an array containing the file for all bundle resources having the specified filename extension, residing in the specified resource subdirectory, and limited to global resources and those associated with the specified localization.
//
// # Finding Image Resources
//
//   - [Bundle.URLForImageResource]: Returns the location of the specified image resource as an NSURL.
//   - [Bundle.PathForImageResource]: Returns the location of the specified image resource file.
//   - [Bundle.ImageForResource]: Returns an [NSImage] instance associated with the specified name, which can be backed by multiple files representing different resolution versions of the image.
//
// # Finding Sound Resources
//
//   - [Bundle.PathForSoundResource]: Returns the location of the specified sound resource file.
//
// # Fetching Localized Strings
//
//   - [Bundle.LocalizedStringForKeyValueTable]: Returns a localized version of the string designated by the specified key and residing in the specified table.
//
// # Fetching Context Help Resources
//
//   - [Bundle.ContextHelpForKey]: Returns the context-sensitive help for the specified key from the bundle’s help file.
//
// # Getting the Standard Bundle Directories
//
//   - [Bundle.ResourceURL]: The file URL of the bundle’s subdirectory containing resource files.
//   - [Bundle.ExecutableURL]: The file URL of the receiver’s executable file.
//   - [Bundle.PrivateFrameworksURL]: The file URL of the bundle’s subdirectory containing private frameworks.
//   - [Bundle.SharedFrameworksURL]: The file URL of the receiver’s subdirectory containing shared frameworks.
//   - [Bundle.BuiltInPlugInsURL]: The file URL of the receiver’s subdirectory containing plug-ins.
//   - [Bundle.URLForAuxiliaryExecutable]: Returns the file URL of the executable with the specified name in the receiver’s bundle.
//   - [Bundle.SharedSupportURL]: The file URL of the bundle’s subdirectory containing shared support files.
//   - [Bundle.AppStoreReceiptURL]: The file URL for the bundle’s App Store receipt.
//   - [Bundle.ResourcePath]: The full pathname of the bundle’s subdirectory containing resources.
//   - [Bundle.ExecutablePath]: The full pathname of the receiver’s executable file.
//   - [Bundle.PrivateFrameworksPath]: The full pathname of the bundle’s subdirectory containing private frameworks.
//   - [Bundle.SharedFrameworksPath]: The full pathname of the bundle’s subdirectory containing shared frameworks.
//   - [Bundle.BuiltInPlugInsPath]: The full pathname of the receiver’s subdirectory containing plug-ins.
//   - [Bundle.PathForAuxiliaryExecutable]: Returns the full pathname of the executable with the specified name in the receiver’s bundle.
//   - [Bundle.SharedSupportPath]: The full pathname of the bundle’s subdirectory containing shared support files.
//
// # Getting Bundle Information
//
//   - [Bundle.BundleURL]: The full URL of the receiver’s bundle directory.
//   - [Bundle.BundlePath]: The full pathname of the receiver’s bundle directory.
//   - [Bundle.BundleIdentifier]: The receiver’s bundle identifier.
//   - [Bundle.InfoDictionary]: A dictionary, constructed from the bundle’s `Info.Plist()` file, that contains information about the receiver.
//   - [Bundle.ObjectForInfoDictionaryKey]: Returns the value associated with the specified key in the receiver’s information property list.
//
// # Getting Localization Information
//
//   - [Bundle.Localizations]: A list of all the localizations contained in the bundle.
//   - [Bundle.PreferredLocalizations]: An ordered list of preferred localizations contained in the bundle.
//   - [Bundle.DevelopmentLocalization]: The localization for the development language.
//   - [Bundle.LocalizedInfoDictionary]: A dictionary with the keys from the bundle’s localized property list.
//
// # Getting Classes from a Bundle
//
//   - [Bundle.ClassNamed]: Returns the [Class] object for the specified name.
//   - [Bundle.PrincipalClass]: The bundle’s principal class.
//   - [Bundle.NSLoadedClasses]: A constant used as a key for the 
//
// # Loading Code from a Bundle
//
//   - [Bundle.ExecutableArchitectures]: An array of numbers indicating the architecture types supported by the bundle’s executable.
//   - [Bundle.PreflightAndReturnError]: Returns a Boolean value indicating whether the bundle’s executable code could be loaded successfully.
//   - [Bundle.LoadAndReturnError]: Loads the bundle’s executable code and returns any errors.
//   - [Bundle.Unload]: Unloads the code associated with the receiver.
//   - [Bundle.Loaded]: The load status of a bundle.
//
// # Errors
//
//   - [Bundle.NSExecutableErrorMinimum]: The beginning of the range of error codes reserved for errors related to executable files.
//   - [Bundle.SetNSExecutableErrorMinimum]
//   - [Bundle.NSExecutableNotLoadableError]: The executable type isn’t loadable in the current process.
//   - [Bundle.SetNSExecutableNotLoadableError]
//   - [Bundle.NSExecutableArchitectureMismatchError]: The executable doesn’t provide an architecture compatible with the current process.
//   - [Bundle.SetNSExecutableArchitectureMismatchError]
//   - [Bundle.NSExecutableRuntimeMismatchError]: The executable has Objective-C runtime information that’s incompatible with the current process.
//   - [Bundle.SetNSExecutableRuntimeMismatchError]
//   - [Bundle.NSExecutableLoadError]: Executable cannot be loaded for an otherwise-unspecified reason.
//   - [Bundle.SetNSExecutableLoadError]
//   - [Bundle.NSExecutableLinkError]: The executable failed due to linking issues.
//   - [Bundle.SetNSExecutableLinkError]
//   - [Bundle.NSExecutableErrorMaximum]: The end of the range of error codes reserved for errors related to executable files.
//   - [Bundle.SetNSExecutableErrorMaximum]
//
// # Instance Methods
//
//   - [Bundle.LoadAppleScriptObjectiveCScripts]
//
// See: https://developer.apple.com/documentation/Foundation/Bundle
type Bundle struct {
	objectivec.Object
}

// BundleFromID constructs a [Bundle] from an objc.ID.
//
// A representation of the code and resources stored in a bundle directory on
// disk.
func BundleFromID(id objc.ID) Bundle {
	return NSBundle{objectivec.Object{id}}
}

// NSBundleFromID is an alias for [BundleFromID] for cross-framework compatibility.
func NSBundleFromID(id objc.ID) Bundle { return BundleFromID(id) }
// NOTE: Bundle adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [Bundle] class.
//
// # Creating and Initializing a Bundle
//
//   - [IBundle.InitWithURL]: Returns an [NSBundle] object initialized to correspond to the specified file URL.
//   - [IBundle.InitWithPath]: Returns an [NSBundle] object initialized to correspond to the specified directory.
//
// # Loading Nib Files
//
//   - [IBundle.LoadNibNamedOwnerTopLevelObjects]: Loads a nib from the bundle with the specified file name and owner.
//
// # Finding Resource Files
//
//   - [IBundle.URLForResourceWithExtensionSubdirectory]: Returns the file URL for the resource file identified by the specified name and extension and residing in a given bundle directory.
//   - [IBundle.URLForResourceWithExtension]: Returns the file URL for the resource identified by the specified name and file extension.
//   - [IBundle.URLsForResourcesWithExtensionSubdirectory]: Returns an array of file URLs for all resources identified by the specified file extension and located in the specified bundle subdirectory.
//   - [IBundle.URLForResourceWithExtensionSubdirectoryLocalization]: Returns the file URL for the resource identified by the specified name and file extension, located in the specified bundle subdirectory, and limited to global resources and those associated with the specified localization.
//   - [IBundle.URLsForResourcesWithExtensionSubdirectoryLocalization]: Returns an array containing the file URLs for all bundle resources having the specified filename extension, residing in the specified resource subdirectory, and limited to global resources and those associated with the specified localization.
//   - [IBundle.PathForResourceOfType]: Returns the full pathname for the resource identified by the specified name and file extension.
//   - [IBundle.PathForResourceOfTypeInDirectory]: Returns the full pathname for the resource identified by the specified name and file extension and located in the specified bundle subdirectory.
//   - [IBundle.PathForResourceOfTypeInDirectoryForLocalization]: Returns the full pathname for the resource identified by the specified name and file extension, located in the specified bundle subdirectory, and limited to global resources and those associated with the specified localization.
//   - [IBundle.PathsForResourcesOfTypeInDirectory]: Returns an array containing the pathnames for all bundle resources having the specified filename extension and residing in the resource subdirectory.
//   - [IBundle.PathsForResourcesOfTypeInDirectoryForLocalization]: Returns an array containing the file for all bundle resources having the specified filename extension, residing in the specified resource subdirectory, and limited to global resources and those associated with the specified localization.
//
// # Finding Image Resources
//
//   - [IBundle.URLForImageResource]: Returns the location of the specified image resource as an NSURL.
//   - [IBundle.PathForImageResource]: Returns the location of the specified image resource file.
//   - [IBundle.ImageForResource]: Returns an [NSImage] instance associated with the specified name, which can be backed by multiple files representing different resolution versions of the image.
//
// # Finding Sound Resources
//
//   - [IBundle.PathForSoundResource]: Returns the location of the specified sound resource file.
//
// # Fetching Localized Strings
//
//   - [IBundle.LocalizedStringForKeyValueTable]: Returns a localized version of the string designated by the specified key and residing in the specified table.
//
// # Fetching Context Help Resources
//
//   - [IBundle.ContextHelpForKey]: Returns the context-sensitive help for the specified key from the bundle’s help file.
//
// # Getting the Standard Bundle Directories
//
//   - [IBundle.ResourceURL]: The file URL of the bundle’s subdirectory containing resource files.
//   - [IBundle.ExecutableURL]: The file URL of the receiver’s executable file.
//   - [IBundle.PrivateFrameworksURL]: The file URL of the bundle’s subdirectory containing private frameworks.
//   - [IBundle.SharedFrameworksURL]: The file URL of the receiver’s subdirectory containing shared frameworks.
//   - [IBundle.BuiltInPlugInsURL]: The file URL of the receiver’s subdirectory containing plug-ins.
//   - [IBundle.URLForAuxiliaryExecutable]: Returns the file URL of the executable with the specified name in the receiver’s bundle.
//   - [IBundle.SharedSupportURL]: The file URL of the bundle’s subdirectory containing shared support files.
//   - [IBundle.AppStoreReceiptURL]: The file URL for the bundle’s App Store receipt.
//   - [IBundle.ResourcePath]: The full pathname of the bundle’s subdirectory containing resources.
//   - [IBundle.ExecutablePath]: The full pathname of the receiver’s executable file.
//   - [IBundle.PrivateFrameworksPath]: The full pathname of the bundle’s subdirectory containing private frameworks.
//   - [IBundle.SharedFrameworksPath]: The full pathname of the bundle’s subdirectory containing shared frameworks.
//   - [IBundle.BuiltInPlugInsPath]: The full pathname of the receiver’s subdirectory containing plug-ins.
//   - [IBundle.PathForAuxiliaryExecutable]: Returns the full pathname of the executable with the specified name in the receiver’s bundle.
//   - [IBundle.SharedSupportPath]: The full pathname of the bundle’s subdirectory containing shared support files.
//
// # Getting Bundle Information
//
//   - [IBundle.BundleURL]: The full URL of the receiver’s bundle directory.
//   - [IBundle.BundlePath]: The full pathname of the receiver’s bundle directory.
//   - [IBundle.BundleIdentifier]: The receiver’s bundle identifier.
//   - [IBundle.InfoDictionary]: A dictionary, constructed from the bundle’s `Info.Plist()` file, that contains information about the receiver.
//   - [IBundle.ObjectForInfoDictionaryKey]: Returns the value associated with the specified key in the receiver’s information property list.
//
// # Getting Localization Information
//
//   - [IBundle.Localizations]: A list of all the localizations contained in the bundle.
//   - [IBundle.PreferredLocalizations]: An ordered list of preferred localizations contained in the bundle.
//   - [IBundle.DevelopmentLocalization]: The localization for the development language.
//   - [IBundle.LocalizedInfoDictionary]: A dictionary with the keys from the bundle’s localized property list.
//
// # Getting Classes from a Bundle
//
//   - [IBundle.ClassNamed]: Returns the [Class] object for the specified name.
//   - [IBundle.PrincipalClass]: The bundle’s principal class.
//   - [IBundle.NSLoadedClasses]: A constant used as a key for the 
//
// # Loading Code from a Bundle
//
//   - [IBundle.ExecutableArchitectures]: An array of numbers indicating the architecture types supported by the bundle’s executable.
//   - [IBundle.PreflightAndReturnError]: Returns a Boolean value indicating whether the bundle’s executable code could be loaded successfully.
//   - [IBundle.LoadAndReturnError]: Loads the bundle’s executable code and returns any errors.
//   - [IBundle.Unload]: Unloads the code associated with the receiver.
//   - [IBundle.Loaded]: The load status of a bundle.
//
// # Errors
//
//   - [IBundle.NSExecutableErrorMinimum]: The beginning of the range of error codes reserved for errors related to executable files.
//   - [IBundle.SetNSExecutableErrorMinimum]
//   - [IBundle.NSExecutableNotLoadableError]: The executable type isn’t loadable in the current process.
//   - [IBundle.SetNSExecutableNotLoadableError]
//   - [IBundle.NSExecutableArchitectureMismatchError]: The executable doesn’t provide an architecture compatible with the current process.
//   - [IBundle.SetNSExecutableArchitectureMismatchError]
//   - [IBundle.NSExecutableRuntimeMismatchError]: The executable has Objective-C runtime information that’s incompatible with the current process.
//   - [IBundle.SetNSExecutableRuntimeMismatchError]
//   - [IBundle.NSExecutableLoadError]: Executable cannot be loaded for an otherwise-unspecified reason.
//   - [IBundle.SetNSExecutableLoadError]
//   - [IBundle.NSExecutableLinkError]: The executable failed due to linking issues.
//   - [IBundle.SetNSExecutableLinkError]
//   - [IBundle.NSExecutableErrorMaximum]: The end of the range of error codes reserved for errors related to executable files.
//   - [IBundle.SetNSExecutableErrorMaximum]
//
// # Instance Methods
//
//   - [IBundle.LoadAppleScriptObjectiveCScripts]
//
// See: https://developer.apple.com/documentation/Foundation/Bundle
type IBundle interface {
	objectivec.IObject

	// Topic: Creating and Initializing a Bundle

	// Returns an [NSBundle] object initialized to correspond to the specified file URL.
	InitWithURL(url INSURL) Bundle
	// Returns an [NSBundle] object initialized to correspond to the specified directory.
	InitWithPath(path string) Bundle

	// Topic: Loading Nib Files

	// Loads a nib from the bundle with the specified file name and owner.
	LoadNibNamedOwnerTopLevelObjects(nibName INSString, owner objectivec.IObject, topLevelObjects INSArray) bool

	// Topic: Finding Resource Files

	// Returns the file URL for the resource file identified by the specified name and extension and residing in a given bundle directory.
	URLForResourceWithExtensionSubdirectory(name string, ext string, subpath string) INSURL
	// Returns the file URL for the resource identified by the specified name and file extension.
	URLForResourceWithExtension(name string, ext string) INSURL
	// Returns an array of file URLs for all resources identified by the specified file extension and located in the specified bundle subdirectory.
	URLsForResourcesWithExtensionSubdirectory(ext string, subpath string) []NSURL
	// Returns the file URL for the resource identified by the specified name and file extension, located in the specified bundle subdirectory, and limited to global resources and those associated with the specified localization.
	URLForResourceWithExtensionSubdirectoryLocalization(name string, ext string, subpath string, localizationName string) INSURL
	// Returns an array containing the file URLs for all bundle resources having the specified filename extension, residing in the specified resource subdirectory, and limited to global resources and those associated with the specified localization.
	URLsForResourcesWithExtensionSubdirectoryLocalization(ext string, subpath string, localizationName string) []NSURL
	// Returns the full pathname for the resource identified by the specified name and file extension.
	PathForResourceOfType(name string, ext string) string
	// Returns the full pathname for the resource identified by the specified name and file extension and located in the specified bundle subdirectory.
	PathForResourceOfTypeInDirectory(name string, ext string, subpath string) string
	// Returns the full pathname for the resource identified by the specified name and file extension, located in the specified bundle subdirectory, and limited to global resources and those associated with the specified localization.
	PathForResourceOfTypeInDirectoryForLocalization(name string, ext string, subpath string, localizationName string) string
	// Returns an array containing the pathnames for all bundle resources having the specified filename extension and residing in the resource subdirectory.
	PathsForResourcesOfTypeInDirectory(ext string, subpath string) []string
	// Returns an array containing the file for all bundle resources having the specified filename extension, residing in the specified resource subdirectory, and limited to global resources and those associated with the specified localization.
	PathsForResourcesOfTypeInDirectoryForLocalization(ext string, subpath string, localizationName string) []string

	// Topic: Finding Image Resources

	// Returns the location of the specified image resource as an NSURL.
	URLForImageResource(name INSString) INSURL
	// Returns the location of the specified image resource file.
	PathForImageResource(name INSString) string
	// Returns an [NSImage] instance associated with the specified name, which can be backed by multiple files representing different resolution versions of the image.
	ImageForResource(name INSString) objc.ID

	// Topic: Finding Sound Resources

	// Returns the location of the specified sound resource file.
	PathForSoundResource(name INSString) string

	// Topic: Fetching Localized Strings

	// Returns a localized version of the string designated by the specified key and residing in the specified table.
	LocalizedStringForKeyValueTable(key string, value string, tableName string) string

	// Topic: Fetching Context Help Resources

	// Returns the context-sensitive help for the specified key from the bundle’s help file.
	ContextHelpForKey(key INSString) INSAttributedString

	// Topic: Getting the Standard Bundle Directories

	// The file URL of the bundle’s subdirectory containing resource files.
	ResourceURL() INSURL
	// The file URL of the receiver’s executable file.
	ExecutableURL() INSURL
	// The file URL of the bundle’s subdirectory containing private frameworks.
	PrivateFrameworksURL() INSURL
	// The file URL of the receiver’s subdirectory containing shared frameworks.
	SharedFrameworksURL() INSURL
	// The file URL of the receiver’s subdirectory containing plug-ins.
	BuiltInPlugInsURL() INSURL
	// Returns the file URL of the executable with the specified name in the receiver’s bundle.
	URLForAuxiliaryExecutable(executableName string) INSURL
	// The file URL of the bundle’s subdirectory containing shared support files.
	SharedSupportURL() INSURL
	// The file URL for the bundle’s App Store receipt.
	AppStoreReceiptURL() INSURL
	// The full pathname of the bundle’s subdirectory containing resources.
	ResourcePath() string
	// The full pathname of the receiver’s executable file.
	ExecutablePath() string
	// The full pathname of the bundle’s subdirectory containing private frameworks.
	PrivateFrameworksPath() string
	// The full pathname of the bundle’s subdirectory containing shared frameworks.
	SharedFrameworksPath() string
	// The full pathname of the receiver’s subdirectory containing plug-ins.
	BuiltInPlugInsPath() string
	// Returns the full pathname of the executable with the specified name in the receiver’s bundle.
	PathForAuxiliaryExecutable(executableName string) string
	// The full pathname of the bundle’s subdirectory containing shared support files.
	SharedSupportPath() string

	// Topic: Getting Bundle Information

	// The full URL of the receiver’s bundle directory.
	BundleURL() INSURL
	// The full pathname of the receiver’s bundle directory.
	BundlePath() string
	// The receiver’s bundle identifier.
	BundleIdentifier() string
	// A dictionary, constructed from the bundle’s `Info.Plist()` file, that contains information about the receiver.
	InfoDictionary() INSDictionary
	// Returns the value associated with the specified key in the receiver’s information property list.
	ObjectForInfoDictionaryKey(key string) objectivec.IObject

	// Topic: Getting Localization Information

	// A list of all the localizations contained in the bundle.
	Localizations() []string
	// An ordered list of preferred localizations contained in the bundle.
	PreferredLocalizations() []string
	// The localization for the development language.
	DevelopmentLocalization() string
	// A dictionary with the keys from the bundle’s localized property list.
	LocalizedInfoDictionary() INSDictionary

	// Topic: Getting Classes from a Bundle

	// Returns the [Class] object for the specified name.
	ClassNamed(className string) objc.Class
	// The bundle’s principal class.
	PrincipalClass() objc.Class
	// A constant used as a key for the 
	NSLoadedClasses() string

	// Topic: Loading Code from a Bundle

	// An array of numbers indicating the architecture types supported by the bundle’s executable.
	ExecutableArchitectures() []NSNumber
	// Returns a Boolean value indicating whether the bundle’s executable code could be loaded successfully.
	PreflightAndReturnError() (bool, error)
	// Loads the bundle’s executable code and returns any errors.
	LoadAndReturnError() (bool, error)
	// Unloads the code associated with the receiver.
	Unload() bool
	// The load status of a bundle.
	Loaded() bool

	// Topic: Errors

	// The beginning of the range of error codes reserved for errors related to executable files.
	NSExecutableErrorMinimum() int
	SetNSExecutableErrorMinimum(value int)
	// The executable type isn’t loadable in the current process.
	NSExecutableNotLoadableError() int
	SetNSExecutableNotLoadableError(value int)
	// The executable doesn’t provide an architecture compatible with the current process.
	NSExecutableArchitectureMismatchError() int
	SetNSExecutableArchitectureMismatchError(value int)
	// The executable has Objective-C runtime information that’s incompatible with the current process.
	NSExecutableRuntimeMismatchError() int
	SetNSExecutableRuntimeMismatchError(value int)
	// Executable cannot be loaded for an otherwise-unspecified reason.
	NSExecutableLoadError() int
	SetNSExecutableLoadError(value int)
	// The executable failed due to linking issues.
	NSExecutableLinkError() int
	SetNSExecutableLinkError(value int)
	// The end of the range of error codes reserved for errors related to executable files.
	NSExecutableErrorMaximum() int
	SetNSExecutableErrorMaximum(value int)

	// Topic: Instance Methods

	LoadAppleScriptObjectiveCScripts()

	LocalizedAttributedStringForKeyValueTable(key string, value string, tableName string) INSAttributedString
	// Look up a localized string given a list of available localizations.
	LocalizedStringForKeyValueTableLocalizations(key string, value string, tableName string, localizations []string) string
}





// Init initializes the instance.
func (b Bundle) Init() Bundle {
	rv := objc.Send[Bundle](b.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b Bundle) Autorelease() Bundle {
	rv := objc.Send[Bundle](b.ID, objc.Sel("autorelease"))
	return rv
}

// NewBundle creates a new Bundle instance.
func NewBundle() Bundle {
	class := getBundleClass()
	rv := objc.Send[Bundle](objc.ID(class.class), objc.Sel("new"))
	return rv
}






// Returns the [NSBundle] object with which the specified class is associated.
//
// aClass: A class.
//
// # Return Value
// 
// The [NSBundle] object that dynamically loaded `aClass` (a loadable bundle),
// the [NSBundle] object for the framework in which `aClass` is defined, or
// the main bundle object if `aClass` was not dynamically loaded or is not
// defined in a framework. This method creates and returns a new [NSBundle]
// object if there is no existing bundle associated with `aClass`. Otherwise,
// the existing instance is returned.
//
// See: https://developer.apple.com/documentation/Foundation/Bundle/init(for:)
func NewBundleForClass(aClass objc.Class) Bundle {
	rv := objc.Send[objc.ID](objc.ID(getBundleClass().class), objc.Sel("bundleForClass:"), aClass)
	return BundleFromID(rv)
}


// Returns the [NSBundle] instance that has the specified bundle identifier.
//
// identifier: The identifier for an existing [NSBundle] instance.
//
// # Return Value
// 
// The [NSBundle] object with the bundle identifier `identifier`, or `nil` if
// the requested bundle is not found on the system. This method creates and
// returns a new [NSBundle] object if there is no existing bundle associated
// with `identifier`. Otherwise, the existing instance is returned.
//
// # Discussion
// 
// This method is typically used by frameworks and plug-ins to locate their
// own bundle at runtime. This method may be somewhat more efficient than
// trying to locate the bundle using the [BundleForClass] method. However, if
// the initial lookup of an already loaded and cached bundle with the
// specified identifier fails, this method uses potentially time-consuming
// heuristics to attempt to locate the bundle. As an optimization, you can use
// the [BundleWithPath] or [BundleWithURL] method instead to avoid file system
// traversal.
//
// See: https://developer.apple.com/documentation/Foundation/Bundle/init(identifier:)
func NewBundleWithIdentifier(identifier string) Bundle {
	rv := objc.Send[objc.ID](objc.ID(getBundleClass().class), objc.Sel("bundleWithIdentifier:"), objc.String(identifier))
	return BundleFromID(rv)
}


// Returns an [NSBundle] object initialized to correspond to the specified
// directory.
//
// path: The path to a directory. This must be a full pathname for a directory; if
// it contains any symbolic links, they must be resolvable.
//
// # Return Value
// 
// An [NSBundle] object initialized to correspond to `fullPath`. This method
// initializes and returns a new instance only if there is no existing bundle
// associated with `fullPath`, otherwise it deallocates `self` and returns the
// existing object. If `fullPath` doesn’t exist or the user doesn’t have
// access to it, returns `nil`.
//
// # Discussion
// 
// It’s not necessary to allocate and initialize an instance for the main
// bundle; use the [MainBundle] class method to get this instance. You can
// also use the [BundleWithPath] class method to obtain a bundle identified by
// its directory path.
//
// See: https://developer.apple.com/documentation/Foundation/Bundle/init(path:)
func NewBundleWithPath(path string) Bundle {
	instance := getBundleClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithPath:"), objc.String(path))
	return BundleFromID(rv)
}


// Returns an [NSBundle] object initialized to correspond to the specified
// file URL.
//
// url: The file URL to a directory. This must be a full URL for a directory; if it
// contains any symbolic links, they must be resolvable.
//
// # Return Value
// 
// An [NSBundle] object initialized to correspond to `url`. This method
// initializes and returns a new instance only if there is no existing bundle
// associated with `url`, otherwise it deallocates `self` and returns the
// existing object. If `url` doesn’t exist or the user doesn’t have access
// to it, returns `nil`.
//
// # Discussion
// 
// It’s not necessary to allocate and initialize an instance for the main
// bundle; use the [MainBundle] class method to get this instance. You can
// also use the [BundleWithURL] class method to obtain a bundle identified by
// its file URL.
//
// See: https://developer.apple.com/documentation/Foundation/Bundle/init(url:)
func NewBundleWithURL(url INSURL) Bundle {
	instance := getBundleClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithURL:"), url)
	return BundleFromID(rv)
}







// Returns an [NSBundle] object initialized to correspond to the specified
// file URL.
//
// url: The file URL to a directory. This must be a full URL for a directory; if it
// contains any symbolic links, they must be resolvable.
//
// # Return Value
// 
// An [NSBundle] object initialized to correspond to `url`. This method
// initializes and returns a new instance only if there is no existing bundle
// associated with `url`, otherwise it deallocates `self` and returns the
// existing object. If `url` doesn’t exist or the user doesn’t have access
// to it, returns `nil`.
//
// # Discussion
// 
// It’s not necessary to allocate and initialize an instance for the main
// bundle; use the [MainBundle] class method to get this instance. You can
// also use the [BundleWithURL] class method to obtain a bundle identified by
// its file URL.
//
// See: https://developer.apple.com/documentation/Foundation/Bundle/init(url:)
func (b Bundle) InitWithURL(url INSURL) Bundle {
	rv := objc.Send[Bundle](b.ID, objc.Sel("initWithURL:"), url)
	return rv
}

// Returns an [NSBundle] object initialized to correspond to the specified
// directory.
//
// path: The path to a directory. This must be a full pathname for a directory; if
// it contains any symbolic links, they must be resolvable.
//
// # Return Value
// 
// An [NSBundle] object initialized to correspond to `fullPath`. This method
// initializes and returns a new instance only if there is no existing bundle
// associated with `fullPath`, otherwise it deallocates `self` and returns the
// existing object. If `fullPath` doesn’t exist or the user doesn’t have
// access to it, returns `nil`.
//
// # Discussion
// 
// It’s not necessary to allocate and initialize an instance for the main
// bundle; use the [MainBundle] class method to get this instance. You can
// also use the [BundleWithPath] class method to obtain a bundle identified by
// its directory path.
//
// See: https://developer.apple.com/documentation/Foundation/Bundle/init(path:)
func (b Bundle) InitWithPath(path string) Bundle {
	rv := objc.Send[Bundle](b.ID, objc.Sel("initWithPath:"), objc.String(path))
	return rv
}

// Loads a nib from the bundle with the specified file name and owner.
//
// nibName: The name of the nib.
//
// owner: The object that will be the nib’s owner.
//
// topLevelObjects: This by-reference parameter is populated with the top level objects of the
// nib.
//
// # Return Value
// 
// [true] if the nib file was loaded successfully; otherwise, [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// Unlike legacy methods, the objects adhere to the standard cocoa memory
// management rules; it is necessary to keep a strong reference to them by
// using IBOutlets or holding a reference to the array to prevent the nib
// contents from being deallocated.
// 
// Outlets to top-level objects should be strong references to demonstrate
// ownership and prevent deallocation.
// 
// For more information on Nibs, see [NSNib].
//
// [NSNib]: https://developer.apple.com/documentation/AppKit/NSNib
//
// See: https://developer.apple.com/documentation/Foundation/Bundle/loadNibNamed(_:owner:topLevelObjects:)
func (b Bundle) LoadNibNamedOwnerTopLevelObjects(nibName INSString, owner objectivec.IObject, topLevelObjects INSArray) bool {
	rv := objc.Send[bool](b.ID, objc.Sel("loadNibNamed:owner:topLevelObjects:"), nibName, owner, topLevelObjects)
	return rv
}

// Returns the file URL for the resource file identified by the specified name
// and extension and residing in a given bundle directory.
//
// name: The name of a resource file contained in the directory specified by
// `subpath`.
// 
// If you specify `nil`, the method returns the first resource file it finds
// with the specified extension in that directory.
//
// ext: The filename extension of the file to locate.
// 
// If you specify an empty string or `nil`, the extension is assumed not to
// exist and the file URL is the first file encountered that exactly matches
// `name`.
//
// subpath: The path of a top-level bundle directory. This must be a valid path. For
// example, to specify the bundle directory for a Mac app, you might specify
// the path `/Applications/MyApp.App()`.
//
// # Return Value
// 
// The file URL for the resource file or `nil` if the file could not be
// located. This method also returns `nil` if the bundle specified by the
// `bundlePath` parameter does not exist or is not a readable directory.
//
// # Discussion
// 
// For details on how localized resources are found, read [The Bundle Search
// Pattern] in [Bundle Programming Guide].
//
// [Bundle Programming Guide]: https://developer.apple.com/library/archive/documentation/CoreFoundation/Conceptual/CFBundles/Introduction/Introduction.html#//apple_ref/doc/uid/10000123i
// [The Bundle Search Pattern]: https://developer.apple.com/library/archive/documentation/CoreFoundation/Conceptual/CFBundles/AccessingaBundlesContents/AccessingaBundlesContents.html#//apple_ref/doc/uid/10000123i-CH104-SW7
//
// See: https://developer.apple.com/documentation/Foundation/Bundle/url(forResource:withExtension:subdirectory:)
func (b Bundle) URLForResourceWithExtensionSubdirectory(name string, ext string, subpath string) INSURL {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("URLForResource:withExtension:subdirectory:"), objc.String(name), objc.String(ext), objc.String(subpath))
	return NSURLFromID(rv)
}

// Returns the file URL for the resource identified by the specified name and
// file extension.
//
// name: The name of the resource file.
// 
// If you specify `nil`, the method returns the first resource file it finds
// with the specified extension.
//
// ext: The extension of the resource file.
// 
// If `extension` is an empty string or `nil`, the extension is assumed not to
// exist and the file URL is the first file encountered that exactly matches
// `name`.
//
// # Return Value
// 
// The file URL for the resource file or `nil` if the file could not be
// located.
//
// # Discussion
// 
// If `extension` is an empty string or `nil`, the returned pathname is the
// first one encountered where the file name exactly matches `name`. For
// details on how localized resources are found, read [The Bundle Search
// Pattern] in [Bundle Programming Guide].
//
// [Bundle Programming Guide]: https://developer.apple.com/library/archive/documentation/CoreFoundation/Conceptual/CFBundles/Introduction/Introduction.html#//apple_ref/doc/uid/10000123i
// [The Bundle Search Pattern]: https://developer.apple.com/library/archive/documentation/CoreFoundation/Conceptual/CFBundles/AccessingaBundlesContents/AccessingaBundlesContents.html#//apple_ref/doc/uid/10000123i-CH104-SW7
//
// See: https://developer.apple.com/documentation/Foundation/Bundle/url(forResource:withExtension:)
func (b Bundle) URLForResourceWithExtension(name string, ext string) INSURL {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("URLForResource:withExtension:"), objc.String(name), objc.String(ext))
	return NSURLFromID(rv)
}

// Returns an array of file URLs for all resources identified by the specified
// file extension and located in the specified bundle subdirectory.
//
// ext: The filename extension of the files to locate.
// 
// If you specify an empty string or `nil`, the extension is assumed not to
// exist and all of the files in `subpath` are returned.
//
// subpath: The name of the bundle subdirectory.
//
// # Return Value
// 
// An array of file URLs for the resource files or `nil` if no files could be
// located at `subpath` with `extension`. Returns an empty array if no
// matching resource files are found.
//
// # Discussion
// 
// If `subpath` is `nil`, this method searches the top-level non-localized
// resource directory and the top-level of any language-specific directories.
// (In macOS, the top-level non-localized resource directory is typically
// called [Resources] but in iOS, it is the main bundle directory.)
// 
// For example, suppose you have a Mac app with a modern bundle and you
// specify `@"Documentation"` for the `subpath` parameter. This method would
// first look in the `Contents/Resources/Documentation` directory of the
// bundle, followed by the [Documentation] subdirectories of each
// language-specific `XCUIElementTypeLproj` directory. (The search order for
// the language-specific directories corresponds to the user’s preferences.)
// This method does not recurse through any other subdirectories at any of
// these locations. For more details see [The Bundle Search Pattern] in
// [Bundle Programming Guide].
//
// [Bundle Programming Guide]: https://developer.apple.com/library/archive/documentation/CoreFoundation/Conceptual/CFBundles/Introduction/Introduction.html#//apple_ref/doc/uid/10000123i
// [The Bundle Search Pattern]: https://developer.apple.com/library/archive/documentation/CoreFoundation/Conceptual/CFBundles/AccessingaBundlesContents/AccessingaBundlesContents.html#//apple_ref/doc/uid/10000123i-CH104-SW7
//
// See: https://developer.apple.com/documentation/Foundation/Bundle/urls(forResourcesWithExtension:subdirectory:)
func (b Bundle) URLsForResourcesWithExtensionSubdirectory(ext string, subpath string) []NSURL {
	rv := objc.Send[[]objc.ID](b.ID, objc.Sel("URLsForResourcesWithExtension:subdirectory:"), objc.String(ext), objc.String(subpath))
	return objc.ConvertSlice(rv, func(id objc.ID) NSURL {
		return NSURLFromID(id)
	})
}

// Returns the file URL for the resource identified by the specified name and
// file extension, located in the specified bundle subdirectory, and limited
// to global resources and those associated with the specified localization.
//
// name: The name of the resource file.
// 
// If you specify `nil`, the method returns the first resource file it finds
// that matches the remaining criteria.
//
// ext: The filename extension of the file to locate.
// 
// If you specify an empty string or `nil`, the extension is assumed not to
// exist and the file URL is the first file encountered that exactly matches
// `name`.
//
// subpath: The name of the bundle subdirectory to search.
//
// localizationName: The language ID for the localization. This parameter should correspond to
// the name of one of the bundle’s language-specific resource directories
// without the `XCUIElementTypeLproj` extension.
//
// # Return Value
// 
// The file URL for the resource file or `nil` if the file could not be
// located.
//
// # Discussion
// 
// This method is equivalent to [URLsForResourcesWithExtensionSubdirectory],
// except that only nonlocalized resources and those in the language-specific
// `XCUIElementTypeLproj` directory specified by `localizationName` are
// searched.
// 
// There should typically be little reason to use this method—see Getting
// the Current Language and Locale. See also
// [PreferredLocalizationsFromArrayForPreferences] for how to determine what
// localizations are available.
//
// See: https://developer.apple.com/documentation/Foundation/Bundle/url(forResource:withExtension:subdirectory:localization:)
func (b Bundle) URLForResourceWithExtensionSubdirectoryLocalization(name string, ext string, subpath string, localizationName string) INSURL {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("URLForResource:withExtension:subdirectory:localization:"), objc.String(name), objc.String(ext), objc.String(subpath), objc.String(localizationName))
	return NSURLFromID(rv)
}

// Returns an array containing the file URLs for all bundle resources having
// the specified filename extension, residing in the specified resource
// subdirectory, and limited to global resources and those associated with the
// specified localization.
//
// ext: The filename extension of the files to locate.
// 
// If you specify an empty string or `nil`, the extension is assumed not to
// exist and all of the files in `subpath` are returned.
//
// subpath: The name of the bundle subdirectory to search.
//
// localizationName: The language ID for the localization. This parameter should correspond to
// the name of one of the bundle’s language-specific resource directories
// without the `XCUIElementTypeLproj` extension.
//
// # Return Value
// 
// An array containing the file URLs for all bundle resources matching the
// specified criteria. This method returns an empty array if no matching
// resource files are found.
//
// # Discussion
// 
// This method is equivalent to [URLsForResourcesWithExtensionSubdirectory],
// except that only nonlocalized resources and those in the language-specific
// `XCUIElementTypeLproj` directory specified by `localizationName` are
// searched.
//
// See: https://developer.apple.com/documentation/Foundation/Bundle/urls(forResourcesWithExtension:subdirectory:localization:)
func (b Bundle) URLsForResourcesWithExtensionSubdirectoryLocalization(ext string, subpath string, localizationName string) []NSURL {
	rv := objc.Send[[]objc.ID](b.ID, objc.Sel("URLsForResourcesWithExtension:subdirectory:localization:"), objc.String(ext), objc.String(subpath), objc.String(localizationName))
	return objc.ConvertSlice(rv, func(id objc.ID) NSURL {
		return NSURLFromID(id)
	})
}

// Returns the full pathname for the resource identified by the specified name
// and file extension.
//
// name: The name of the resource file.
// 
// If you specify `nil`, the method returns the first resource file it finds
// with the specified extension.
//
// ext: The filename extension of the file to locate.
// 
// If you specify an empty string or `nil`, the extension is assumed not to
// exist and the file is the first file encountered that exactly matches
// `name`.
//
// # Return Value
// 
// The full pathname for the resource file, or `nil` if the file could not be
// located.
//
// # Discussion
// 
// The method first looks for a matching resource file in the non-localized
// resource directory of the specified bundle. If a matching resource file is
// not found, it then looks in the top level of an available language-specific
// `XCUIElementTypeLproj` folder. (The search order for the language-specific
// folders corresponds to the user’s preferences.) It does not recurse
// through other subfolders at any of these locations. For more details on how
// localized resources are found, read [The Bundle Search Pattern] in [Bundle
// Programming Guide].
// 
// The following code fragment gets the path to a plist within the bundle, and
// loads it into an [NSDictionary]:
//
// [Bundle Programming Guide]: https://developer.apple.com/library/archive/documentation/CoreFoundation/Conceptual/CFBundles/Introduction/Introduction.html#//apple_ref/doc/uid/10000123i
// [The Bundle Search Pattern]: https://developer.apple.com/library/archive/documentation/CoreFoundation/Conceptual/CFBundles/AccessingaBundlesContents/AccessingaBundlesContents.html#//apple_ref/doc/uid/10000123i-CH104-SW7
//
// See: https://developer.apple.com/documentation/Foundation/Bundle/path(forResource:ofType:)
func (b Bundle) PathForResourceOfType(name string, ext string) string {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("pathForResource:ofType:"), objc.String(name), objc.String(ext))
	return NSStringFromID(rv).String()
}

// Returns the full pathname for the resource identified by the specified name
// and file extension and located in the specified bundle subdirectory.
//
// name: The name of the resource file.
// 
// If you specify `nil`, the method returns the first resource file it finds
// that matches the remaining criteria.
//
// ext: The filename extension of the files to locate.
// 
// If you specify an empty string or `nil`, all the files in `subpath` and its
// subdirectories are returned. If an extension is provided the subdirectories
// are not searched.
//
// subpath: The name of the bundle subdirectory.
//
// # Return Value
// 
// The full pathname for the resource file, or `nil` if the file could not be
// located.
//
// # Discussion
// 
// If `subpath` is `nil`, this method searches the top-level nonlocalized
// resource directory and the top-level of any language-specific directories.
// (In macOS, the top-level nonlocalized resource directory is typically
// called [Resources] but in iOS, it is the main bundle directory.) For
// example, suppose you have a Mac app with a modern bundle and you specify
// `@"Documentation"` for the `subpath` parameter. This method would first
// look in the `Contents/Resources/Documentation` directory of the bundle,
// followed by the [Documentation] subdirectories of each language-specific
// `XCUIElementTypeLproj` directory.
// 
// Whether this method recurses through subdirectories is dependent on the
// `extension` parameter. If `nil` or an empty string it will recurse,
// otherwise, it does not. (The search order for the language-specific
// directories corresponds to the user’s preferences.) For details on how
// localized resources are found, read [The Bundle Search Pattern] in [Bundle
// Programming Guide].
//
// [Bundle Programming Guide]: https://developer.apple.com/library/archive/documentation/CoreFoundation/Conceptual/CFBundles/Introduction/Introduction.html#//apple_ref/doc/uid/10000123i
// [The Bundle Search Pattern]: https://developer.apple.com/library/archive/documentation/CoreFoundation/Conceptual/CFBundles/AccessingaBundlesContents/AccessingaBundlesContents.html#//apple_ref/doc/uid/10000123i-CH104-SW7
//
// See: https://developer.apple.com/documentation/Foundation/Bundle/path(forResource:ofType:inDirectory:)-swift.method
func (b Bundle) PathForResourceOfTypeInDirectory(name string, ext string, subpath string) string {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("pathForResource:ofType:inDirectory:"), objc.String(name), objc.String(ext), objc.String(subpath))
	return NSStringFromID(rv).String()
}

// Returns the full pathname for the resource identified by the specified name
// and file extension, located in the specified bundle subdirectory, and
// limited to global resources and those associated with the specified
// localization.
//
// name: The name of the resource file.
// 
// If you specify `nil`, the method returns the first resource file it finds
// that matches the remaining criteria.
//
// ext: The filename extension of the files to locate.
// 
// If you specify an empty string or `nil`, the extension is assumed not to
// exist and the file is the first file encountered that exactly matches
// `name`.
//
// subpath: The name of the bundle subdirectory to search.
//
// localizationName: The language ID for of the localization. This parameter should correspond
// to the name of one of the bundle’s language-specific resource directories
// without the `XCUIElementTypeLproj` extension.
//
// # Return Value
// 
// The full pathname for the resource file or `nil` if the file could not be
// located.
//
// # Discussion
// 
// This method is equivalent to [PathForResourceOfTypeInDirectory], except
// that only nonlocalized resources and those in the language-specific
// `XCUIElementTypeLproj` directory specified by `localizationName` are
// searched.
// 
// There should typically be little reason to use this method—see Getting
// the Current Language and Locale. See also
// [PreferredLocalizationsFromArrayForPreferences] for how to determine what
// localizations are available.
//
// See: https://developer.apple.com/documentation/Foundation/Bundle/path(forResource:ofType:inDirectory:forLocalization:)
func (b Bundle) PathForResourceOfTypeInDirectoryForLocalization(name string, ext string, subpath string, localizationName string) string {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("pathForResource:ofType:inDirectory:forLocalization:"), objc.String(name), objc.String(ext), objc.String(subpath), objc.String(localizationName))
	return NSStringFromID(rv).String()
}

// Returns an array containing the pathnames for all bundle resources having
// the specified filename extension and residing in the resource subdirectory.
//
// ext: The filename extension of the files to locate.
// 
// If you specify an empty string or `nil`, the extension is assumed not to
// exist and all of the files in `subpath` are returned.
//
// subpath: The name of the bundle subdirectory to search.
//
// # Return Value
// 
// An array containing the full pathnames for all bundle resources matching
// the specified criteria. This method returns an empty array if no matching
// resource files are found.
//
// # Discussion
// 
// This method provides a means for dynamically discovering multiple bundle
// resources of the same type. If `extension` is an empty string or `nil`, all
// bundle resources in the specified resource directory are returned.
// 
// The argument `subpath` specifies the name of a specific subdirectory to
// search within the current bundle’s resource directory hierarchy. If
// `subpath` is `nil`, this method searches the top-level nonlocalized
// resource directory and the top-level of any language-specific directories.
// (In macOS, the top-level nonlocalized resource directory is typically
// called [Resources] but in iOS, it is the main bundle directory.) For
// example, suppose you have a Mac app with a modern bundle and you specify
// `@"Documentation"` for the `subpath` parameter. This method would first
// look in the `Contents/Resources/Documentation` directory of the bundle,
// followed by the [Documentation] subdirectories of each language-specific
// `XCUIElementTypeLproj` directory. (The search order for the
// language-specific directories corresponds to the user’s preferences.)
// This method does not recurse through any other subdirectories at any of
// these locations. For details on how localized resources are found, read
// [The Bundle Search Pattern] in [Bundle Programming Guide].
//
// [Bundle Programming Guide]: https://developer.apple.com/library/archive/documentation/CoreFoundation/Conceptual/CFBundles/Introduction/Introduction.html#//apple_ref/doc/uid/10000123i
// [The Bundle Search Pattern]: https://developer.apple.com/library/archive/documentation/CoreFoundation/Conceptual/CFBundles/AccessingaBundlesContents/AccessingaBundlesContents.html#//apple_ref/doc/uid/10000123i-CH104-SW7
//
// See: https://developer.apple.com/documentation/Foundation/Bundle/paths(forResourcesOfType:inDirectory:)-swift.method
func (b Bundle) PathsForResourcesOfTypeInDirectory(ext string, subpath string) []string {
	rv := objc.Send[[]objc.ID](b.ID, objc.Sel("pathsForResourcesOfType:inDirectory:"), objc.String(ext), objc.String(subpath))
	return objc.ConvertSliceToStrings(rv)
}

// Returns an array containing the file for all bundle resources having the
// specified filename extension, residing in the specified resource
// subdirectory, and limited to global resources and those associated with the
// specified localization.
//
// ext: The filename extension of the files to locate.
// 
// If you specify an empty string or `nil`, the extension is assumed not to
// exist and all of the files in `subpath` are returned.
//
// subpath: The name of the bundle subdirectory to search.
//
// localizationName: The language ID for the localization. This parameter should correspond to
// the name of one of the bundle’s language-specific resource directories
// without the `XCUIElementTypeLproj` extension.
//
// # Return Value
// 
// An array containing the full pathnames for all bundle resources matching
// the specified criteria. This method returns an empty array if no matching
// resource files are found.
//
// # Discussion
// 
// This method is equivalent to [PathsForResourcesOfTypeInDirectory], except
// that only nonlocalized resources and those in the language-specific
// `XCUIElementTypeLproj` directory specified by `localizationName` are
// searched.
//
// See: https://developer.apple.com/documentation/Foundation/Bundle/paths(forResourcesOfType:inDirectory:forLocalization:)
func (b Bundle) PathsForResourcesOfTypeInDirectoryForLocalization(ext string, subpath string, localizationName string) []string {
	rv := objc.Send[[]objc.ID](b.ID, objc.Sel("pathsForResourcesOfType:inDirectory:forLocalization:"), objc.String(ext), objc.String(subpath), objc.String(localizationName))
	return objc.ConvertSliceToStrings(rv)
}

// Returns the location of the specified image resource as an NSURL.
//
// name: The name of the image resource file. Including a filename extension is
// optional.
//
// # Return Value
// 
// An [NSURL] for the resource file or `nil` if the file was not found.
//
// See: https://developer.apple.com/documentation/Foundation/Bundle/urlForImageResource(_:)
func (b Bundle) URLForImageResource(name INSString) INSURL {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("URLForImageResource:"), name)
	return NSURLFromID(rv)
}

// Returns the location of the specified image resource file.
//
// name: The name of the image resource file, without any pathname information.
// Including a filename extension is optional.
//
// # Return Value
// 
// The absolute pathname of the resource file or `nil` if the file is not
// found.
//
// # Discussion
// 
// Image resources are those files in the bundle that are recognized by the
// [NSImage] class, including those that can be converted using the Image IO
// framework.
//
// See: https://developer.apple.com/documentation/Foundation/Bundle/pathForImageResource(_:)
func (b Bundle) PathForImageResource(name INSString) string {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("pathForImageResource:"), name)
	return NSStringFromID(rv).String()
}

// Returns an [NSImage] instance associated with the specified name, which can
// be backed by multiple files representing different resolution versions of
// the image.
//
// name: The filename of the image resource file. Including a filename extension is
// optional.
//
// # Return Value
// 
// The [NSImage] object associated with the specified name, or `nil` if no
// file is found.
//
// # Discussion
// 
// This method accommodates Apple’s naming conventions for high-resolution
// versions of the image. For example, if your bundle contains files named
// `button.Png()`, `button@2x.Png()`, and `button.Pdf()` then `@"button"`
// returns an [NSImage] object backed by all three files. Each time the
// [NSImage] object is drawn, it selects the representation best for the
// drawing context.
// 
// Images requested using this method whose name ends in the word [Template]
// are automatically marked as template images.
// 
// This method does not look up images based on [setName(_:)] or get named
// system images. Use [init(named:)] for that purpose.
// 
// This method does not cache its search results.
//
// [init(named:)]: https://developer.apple.com/documentation/AppKit/NSImage/init(named:)
// [setName(_:)]: https://developer.apple.com/documentation/AppKit/NSImage/setName(_:)
//
// See: https://developer.apple.com/documentation/Foundation/Bundle/image(forResource:)
func (b Bundle) ImageForResource(name INSString) objc.ID {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("imageForResource:"), name)
	return rv
}

// Returns the location of the specified sound resource file.
//
// name: The name of the sound resource file, without any pathname information.
// Including a filename extension is optional
//
// # Return Value
// 
// The absolute pathname of the resource file or `nil` if the file was not
// found.
//
// # Discussion
// 
// Sound resources are those files in the bundle that are recognized by the
// [NSSound] class. The types of sound files can be determined by calling the
// `soundUnfilteredFileTypes` method of [NSSound].
//
// See: https://developer.apple.com/documentation/Foundation/Bundle/path(forSoundResource:)
func (b Bundle) PathForSoundResource(name INSString) string {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("pathForSoundResource:"), name)
	return NSStringFromID(rv).String()
}

// Returns a localized version of the string designated by the specified key
// and residing in the specified table.
//
// key: The key for a string in the table identified by `tableName`.
//
// value: The value to return if `key` is `nil` or if a localized string for `key`
// can’t be found in the table.
//
// tableName: The receiver’s string table to search. If `tableName` is `nil` or is an
// empty string, the method attempts to use the table in
// `Localizable.Strings()`.
//
// # Return Value
// 
// A localized version of the string designated by `key` in table `tableName`.
// This method returns the following when key is `nil` or not found in table:
// 
// - If `key` is `nil` and `value` is `nil`, returns an empty string. - If
// `key` is `nil` and `value` is non-`nil`, returns value. - If `key` is not
// found and `value` is `nil` or an empty string, returns `key`. - If `key` is
// not found and `value` is non-`nil` and not empty, return `value`.
//
// # Discussion
// 
// For more details about string localization and the specification of a
// `XCUIElementTypeStrings` file, see “[String Resources].”
// 
// Using the user default [NSShowNonLocalizedStrings], you can alter the
// behavior of [LocalizedStringForKeyValueTable] to log a message when the
// method can’t find a localized string. If you set this default to [true]
// (in the global domain or in the application’s domain), then when the
// method can’t find a localized string in the table, it logs a message to
// the console and capitalizes `key` before returning it.
// 
// The following example cycles through a static array of keys when a button
// is clicked, gets the value for each key from a strings table named
// `Buttons.Strings()`, and sets the button title with the returned value:
//
// [String Resources]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/LoadingResources/Strings/Strings.html#//apple_ref/doc/uid/10000051i-CH6
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Foundation/Bundle/localizedString(forKey:value:table:)
func (b Bundle) LocalizedStringForKeyValueTable(key string, value string, tableName string) string {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("localizedStringForKey:value:table:"), objc.String(key), objc.String(value), objc.String(tableName))
	return NSStringFromID(rv).String()
}

// Returns the context-sensitive help for the specified key from the
// bundle’s help file.
//
// key: A key in your application’s `Help.Plist()` file that identifies the
// context-sensitive help to return.
//
// # Return Value
// 
// The help string or `nil` if the application does not have a `Help.Plist()`
// file or the file does not contain an entry for the specified `key`.
//
// # Discussion
// 
// When you build your application, you can merge multiple RTF-based help
// files together using the `/usr/bin/compileHelp` tool, which then packages
// your help file information into a property list named `Help.Plist()`. After
// placing this property-list file in your application bundle, you can use
// this method to extract context help information from it. To look up a
// particular entry, you specify the name of the original RTF help file in the
// `key` parameter of this method. For example, if your application project
// contains a help file named `Copy.Rtf()`, you would retrieve the text from
// this file by passing the value `@"Copy.Rtf()"` to the `key` parameter.
// 
// This method is declared in `NSHelpManager.H()`.
//
// See: https://developer.apple.com/documentation/Foundation/Bundle/contextHelp(forKey:)
func (b Bundle) ContextHelpForKey(key INSString) INSAttributedString {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("contextHelpForKey:"), key)
	return NSAttributedStringFromID(rv)
}

// Returns the file URL of the executable with the specified name in the
// receiver’s bundle.
//
// executableName: The name of an executable file.
//
// # Return Value
// 
// The file URL of the executable `executableName` in the receiver’s bundle.
//
// # Discussion
// 
// This method returns the appropriate path for modern application and
// framework bundles. This method may not return a URL for non-standard bundle
// formats or for some older bundle formats.
//
// See: https://developer.apple.com/documentation/Foundation/Bundle/url(forAuxiliaryExecutable:)
func (b Bundle) URLForAuxiliaryExecutable(executableName string) INSURL {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("URLForAuxiliaryExecutable:"), objc.String(executableName))
	return NSURLFromID(rv)
}

// Returns the full pathname of the executable with the specified name in the
// receiver’s bundle.
//
// executableName: The name of an executable file.
//
// # Return Value
// 
// The full pathname of the executable `executableName` in the receiver’s
// bundle.
//
// # Discussion
// 
// This method returns the appropriate path for modern application and
// framework bundles. This method may not return a path for non-standard
// bundle formats or for some older bundle formats.
//
// See: https://developer.apple.com/documentation/Foundation/Bundle/path(forAuxiliaryExecutable:)
func (b Bundle) PathForAuxiliaryExecutable(executableName string) string {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("pathForAuxiliaryExecutable:"), objc.String(executableName))
	return NSStringFromID(rv).String()
}

// Returns the value associated with the specified key in the receiver’s
// information property list.
//
// key: A key in the receiver’s property list.
//
// # Return Value
// 
// The value associated with `key` in the receiver’s property list
// (`Info.Plist()`). The localized value of a key is returned when one is
// available.
//
// # Discussion
// 
// Use of this method is preferred over other access methods because it
// returns the localized value of a key when one is available.
//
// See: https://developer.apple.com/documentation/Foundation/Bundle/object(forInfoDictionaryKey:)
func (b Bundle) ObjectForInfoDictionaryKey(key string) objectivec.IObject {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("objectForInfoDictionaryKey:"), objc.String(key))
	return objectivec.Object{ID: rv}
}

// Returns the [Class] object for the specified name.
//
// className: The name of a class.
//
// # Return Value
// 
// The [Class] object for `className`. Returns `nil` if `className` is not one
// of the classes associated with the receiver or if there is an error loading
// the executable code containing the class implementation.
//
// # Discussion
// 
// If the bundle’s executable code is not yet loaded, this method
// dynamically loads it into memory. Classes (and categories) are loaded from
// just one file within the bundle directory; this code file has the same name
// as the directory, but without the extension (”`XCUIElementTypeBundle`”,
// “`XCUIElementTypeApp`”, “`XCUIElementTypeFramework`”). As a side
// effect of code loading, the receiver posts [didLoadNotification] after all
// classes and categories have been loaded; see [Notifications] for details.
//
// [didLoadNotification]: https://developer.apple.com/documentation/Foundation/Bundle/didLoadNotification
//
// See: https://developer.apple.com/documentation/Foundation/Bundle/classNamed(_:)
func (b Bundle) ClassNamed(className string) objc.Class {
	rv := objc.Send[objc.Class](b.ID, objc.Sel("classNamed:"), objc.String(className))
	return rv
}

// Returns a Boolean value indicating whether the bundle’s executable code
// could be loaded successfully.
//
// # Discussion
// 
// This method does not actually load the bundle’s executable code. Instead,
// it performs several checks to see if the code could be loaded and with one
// exception returns the same errors that would occur during an actual load
// operation. The one exception is the [NSExecutableLinkError] error, which
// requires the actual loading of the code to verify link errors.
// 
// For a list of possible load errors, see the discussion for the
// [LoadAndReturnError] method.
//
// See: https://developer.apple.com/documentation/Foundation/Bundle/preflight()
func (b Bundle) PreflightAndReturnError() (bool, error) {
			var errorPtr objc.ID
	rv := objc.Send[bool](b.ID, objc.Sel("preflightAndReturnError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("preflightAndReturnError: returned NO with nil NSError")
	}
	return rv, nil

}

// Loads the bundle’s executable code and returns any errors.
//
// # Discussion
// 
// If this method returns [false] and you pass a value for the `error`
// parameter, a suitable error object is returned in that parameter. Potential
// errors returned are in the Cocoa error domain and include the types that
// follow. For a full list of error types, see `FoundationErrors.H()`.
// 
// - [NSFileNoSuchFileError] - returned if the bundle’s executable file was
// not located. - [NSExecutableNotLoadableError] - returned if the bundle’s
// executable file exists but could not be loaded. This error is returned if
// the executable is not recognized as a loadable executable. It can also be
// returned if the executable is a PEF/CFM executable but the current process
// does not support that type of executable. -
// [NSExecutableArchitectureMismatchError] - returned if the bundle executable
// does not include code that matches the processor architecture of the
// current processor. - [NSExecutableRuntimeMismatchError] - returned if the
// bundle’s required Objective-C runtime information is not compatible with
// the runtime of the current process. - [NSExecutableLoadError] - returned if
// the bundle’s executable failed to load for some detectable reason prior
// to linking. This error might occur if the bundle depends on a framework or
// library that is missing or if the required framework or library is not
// compatible with the current architecture or runtime version. -
// [NSExecutableLinkError] - returned if the executable failed to load due to
// link errors but is otherwise alright.
// 
// The error object may contain additional debugging information in its
// description that you can use to identify the cause of the error. (This
// debugging information should not be displayed to the user.) You can obtain
// the debugging information by invoking the error object’s `description`
// method in your code or by using the `print-object` command on the error
// object in gdb.
//
// [false]: https://developer.apple.com/documentation/Swift/false
//
// See: https://developer.apple.com/documentation/Foundation/Bundle/loadAndReturnError()
func (b Bundle) LoadAndReturnError() (bool, error) {
			var errorPtr objc.ID
	rv := objc.Send[bool](b.ID, objc.Sel("loadAndReturnError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("loadAndReturnError: returned NO with nil NSError")
	}
	return rv, nil

}

// Unloads the code associated with the receiver.
//
// # Return Value
// 
// [true] if the bundle was successfully unloaded or was not already loaded;
// otherwise, [false] if the bundle could not be unloaded.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// This method attempts to unload a bundle’s executable code using the
// underlying dynamic loader (typically `dyld`). You may use this method to
// unload plug-in and framework bundles when you no longer need the code they
// contain. You should use this method to unload bundles that were loaded
// using the methods of the [NSBundle] class only. Do not use this method to
// unload bundles that were originally loaded using the bundle-manipulation
// functions in Core Foundation.
// 
// It is the responsibility of the caller to ensure that no in-memory objects
// or data structures refer to the code being unloaded. For example, if you
// have an object whose class is defined in a bundle, you must release that
// object prior to unloading the bundle. Similarly, your code should not
// attempt to access any symbols defined in an unloaded bundle.
// 
// # Special Considerations
// 
// Prior to OS X version 10.5, code could not be unloaded once loaded, and
// this method would always return [false]. In macOS 10.5 and later, you can
// unload a bundle’s executable code using this method.
//
// [false]: https://developer.apple.com/documentation/Swift/false
//
// See: https://developer.apple.com/documentation/Foundation/Bundle/unload()
func (b Bundle) Unload() bool {
	rv := objc.Send[bool](b.ID, objc.Sel("unload"))
	return rv
}

// See: https://developer.apple.com/documentation/Foundation/Bundle/loadAppleScriptObjectiveCScripts()
func (b Bundle) LoadAppleScriptObjectiveCScripts() {
	objc.Send[objc.ID](b.ID, objc.Sel("loadAppleScriptObjectiveCScripts"))
}

//
// See: https://developer.apple.com/documentation/Foundation/NSBundle/localizedAttributedStringForKey:value:table:
func (b Bundle) LocalizedAttributedStringForKeyValueTable(key string, value string, tableName string) INSAttributedString {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("localizedAttributedStringForKey:value:table:"), objc.String(key), objc.String(value), objc.String(tableName))
	return NSAttributedStringFromID(rv)
}

// Look up a localized string given a list of available localizations.
//
// key: The key for the localized string to retrieve.
//
// value: A default value to return if a localized string for `key` cannot be found.
//
// tableName: The name of the strings file to search. If `nil`, the method uses tables in
// `Localizable.Strings()`.
//
// localizations: An array of BCP 47 language codes corresponding to available localizations.
// Bundle compares the array against its available localizations, and uses the
// best result to retrieve the localized string. If empty, we treat it as no
// localization is available, and may return a fallback.
//
// # Return Value
// 
// A localized version of the string designated by `key` in table `tableName`.
//
// See: https://developer.apple.com/documentation/Foundation/NSBundle/localizedStringForKey:value:table:localizations:
func (b Bundle) LocalizedStringForKeyValueTableLocalizations(key string, value string, tableName string, localizations []string) string {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("localizedStringForKey:value:table:localizations:"), objc.String(key), objc.String(value), objc.String(tableName), objectivec.StringSliceToNSArray(localizations))
	return NSStringFromID(rv).String()
}





// Creates and returns a file URL for the resource with the specified name and
// extension in the specified bundle.
//
// name: The name of the resource file.
// 
// If you specify `nil`, the method returns the first resource file it finds
// that matches the remaining criteria.
//
// ext: The filename extension of the file to locate.
// 
// If you specify an empty string or `nil`, the extension is assumed not to
// exist and the file URL is the first file encountered that exactly matches
// `name`.
//
// subpath: The name of the bundle subdirectory to search.
//
// bundleURL: The file URL of the bundle to search.
//
// # Return Value
// 
// The file URL for the resource file or `nil` if the file could not be
// located.
//
// See: https://developer.apple.com/documentation/Foundation/Bundle/url(forResource:withExtension:subdirectory:in:)
func (_BundleClass BundleClass) URLForResourceWithExtensionSubdirectoryInBundleWithURL(name string, ext string, subpath string, bundleURL INSURL) NSURL {
	rv := objc.Send[objc.ID](objc.ID(_BundleClass.class), objc.Sel("URLForResource:withExtension:subdirectory:inBundleWithURL:"), objc.String(name), objc.String(ext), objc.String(subpath), bundleURL)
	return NSURLFromID(rv)
}

// Returns an array containing the file URLs for all bundle resources having
// the specified filename extension, residing in the specified resource
// subdirectory, within the specified bundle.
//
// ext: The filename extension of the files to locate.
// 
// If you specify an empty string or `nil`, the extension is assumed not to
// exist and all of the files in `subpath` are returned.
//
// subpath: The name of the bundle subdirectory to search.
//
// bundleURL: The file URL of the bundle to search.
//
// # Return Value
// 
// An array of file URLs for the resource files matching the criteria or an
// empty array if no files could be located.
//
// See: https://developer.apple.com/documentation/Foundation/Bundle/urls(forResourcesWithExtension:subdirectory:in:)
func (_BundleClass BundleClass) URLsForResourcesWithExtensionSubdirectoryInBundleWithURL(ext string, subpath string, bundleURL INSURL) []NSURL {
	rv := objc.Send[[]objc.ID](objc.ID(_BundleClass.class), objc.Sel("URLsForResourcesWithExtension:subdirectory:inBundleWithURL:"), objc.String(ext), objc.String(subpath), bundleURL)
	return objc.ConvertSlice(rv, func(id objc.ID) NSURL {
		return NSURLFromID(id)
	})
}

// Returns one or more localizations from the specified list that a bundle
// object would use to locate resources for the current user.
//
// localizationsArray: An array of [NSString] objects, each of which specifies the language ID for
// a localization that the bundle supports.
//
// # Return Value
// 
// An array of [NSString] objects containing the preferred localizations.
// These strings are ordered in the array according to the user’s language
// preferences and are taken from the strings in the `localizationsArray`
// parameter.
//
// # Discussion
// 
// This method does not return all localizations in preference order but only
// those from which [NSBundle] would get localized content, typically either a
// single non-region-specific localization or a region-specific localization
// followed by a corresponding non-region-specific localization as a fallback.
// 
// However, clients who want all localizations in preference order can make
// repeated calls, each time taking the top localizations out of the list of
// localizations passed in.
//
// See: https://developer.apple.com/documentation/Foundation/Bundle/preferredLocalizations(from:)
func (_BundleClass BundleClass) PreferredLocalizationsFromArray(localizationsArray []string) []string {
	rv := objc.Send[[]objc.ID](objc.ID(_BundleClass.class), objc.Sel("preferredLocalizationsFromArray:"), objectivec.StringSliceToNSArray(localizationsArray))
	return objc.ConvertSliceToStrings(rv)
}

// Returns locale identifiers for which a bundle would provide localized
// content, given a specified list of candidates for a user’s language
// preferences.
//
// localizationsArray: An array of identifiers, each corresponding to a localization that a bundle
// can support.
//
// preferencesArray: An array of BCP 47 language codes corresponding to a user’s preferred
// languages.
// 
// If this parameter is `nil`, the method uses the current user’s language
// preferences.
//
// # Return Value
// 
// An array of locale identifiers, ordered according to user preference. If
// none of the user-preferred localizations are available, this method returns
// one of the values in `localizationsArray`.
//
// # Discussion
// 
// This method returns only the locale identifiers for which a bundle would
// provide localized content. Typically, this means one of the following:
// 
// - A single localization that isn’t region-specific - A region-specific
// localization, followed by a corresponding localization that isn’t
// region-specific, as a fallback
// 
// This method doesn’t return all localizations in order of user preference.
// To get this information, you can call this method repeatedly, each time
// removing the identifiers returned by the previous call.
//
// See: https://developer.apple.com/documentation/Foundation/Bundle/preferredLocalizations(from:forPreferences:)
func (_BundleClass BundleClass) PreferredLocalizationsFromArrayForPreferences(localizationsArray []string, preferencesArray []string) []string {
	rv := objc.Send[[]objc.ID](objc.ID(_BundleClass.class), objc.Sel("preferredLocalizationsFromArray:forPreferences:"), objectivec.StringSliceToNSArray(localizationsArray), objectivec.StringSliceToNSArray(preferencesArray))
	return objc.ConvertSliceToStrings(rv)
}

// Returns an [NSBundle] object that corresponds to the specified directory.
//
// path: The path to a directory. This must be a full pathname for a directory; if
// it contains any symbolic links, they must be resolvable.
//
// # Return Value
// 
// The [NSBundle] object that corresponds to `path`, or `nil` if `path` does
// not identify an accessible bundle directory.
//
// # Discussion
// 
// This method allocates and initializes the returned object if there is no
// existing [NSBundle] associated with `path`, in which case it returns the
// existing object.
//
// See: https://developer.apple.com/documentation/Foundation/NSBundle/bundleWithPath:
func (_BundleClass BundleClass) BundleWithPath(path string) Bundle {
	rv := objc.Send[objc.ID](objc.ID(_BundleClass.class), objc.Sel("bundleWithPath:"), objc.String(path))
	return NSBundleFromID(rv)
}

// Returns an [NSBundle] object that corresponds to the specified file URL.
//
// url: The URL to a directory. This must be a URL for a directory; if it contains
// any symbolic links, they must be resolvable.
//
// # Return Value
// 
// The [NSBundle] object that corresponds to `url`, or `nil` if `url` does not
// identify an accessible bundle directory.
//
// # Discussion
// 
// This method allocates and initializes the returned object if there is no
// existing [NSBundle] associated with `url`, in which case it returns the
// existing object.
//
// See: https://developer.apple.com/documentation/Foundation/NSBundle/bundleWithURL:
func (_BundleClass BundleClass) BundleWithURL(url INSURL) Bundle {
	rv := objc.Send[objc.ID](objc.ID(_BundleClass.class), objc.Sel("bundleWithURL:"), url)
	return NSBundleFromID(rv)
}








// The file URL of the bundle’s subdirectory containing resource files.
//
// # Discussion
// 
// This property contains the appropriate path for modern application and
// framework bundles. This property may not contain a path for non-standard
// bundle formats or for some older bundle formats.
//
// See: https://developer.apple.com/documentation/Foundation/Bundle/resourceURL
func (b Bundle) ResourceURL() INSURL {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("resourceURL"))
	return NSURLFromID(objc.ID(rv))
}



// The file URL of the receiver’s executable file.
//
// See: https://developer.apple.com/documentation/Foundation/Bundle/executableURL
func (b Bundle) ExecutableURL() INSURL {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("executableURL"))
	return NSURLFromID(objc.ID(rv))
}



// The file URL of the bundle’s subdirectory containing private frameworks.
//
// # Discussion
// 
// This property contains the appropriate path for modern application and
// framework bundles. This property may not be a URL for non-standard bundle
// formats or for some older bundle formats.
//
// See: https://developer.apple.com/documentation/Foundation/Bundle/privateFrameworksURL
func (b Bundle) PrivateFrameworksURL() INSURL {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("privateFrameworksURL"))
	return NSURLFromID(objc.ID(rv))
}



// The file URL of the receiver’s subdirectory containing shared frameworks.
//
// # Discussion
// 
// This property contains the appropriate path for modern application and
// framework bundles. This property may not contain a URL for non-standard
// bundle formats or for some older bundle formats.
//
// See: https://developer.apple.com/documentation/Foundation/Bundle/sharedFrameworksURL
func (b Bundle) SharedFrameworksURL() INSURL {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("sharedFrameworksURL"))
	return NSURLFromID(objc.ID(rv))
}



// The file URL of the receiver’s subdirectory containing plug-ins.
//
// # Discussion
// 
// This is the appropriate path for modern application and framework bundles.
// This may not be a URL for non-standard bundle formats or for some older
// bundle formats.
//
// See: https://developer.apple.com/documentation/Foundation/Bundle/builtInPlugInsURL
func (b Bundle) BuiltInPlugInsURL() INSURL {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("builtInPlugInsURL"))
	return NSURLFromID(objc.ID(rv))
}



// The file URL of the bundle’s subdirectory containing shared support
// files.
//
// # Discussion
// 
// This property contains the appropriate path for modern application and
// framework bundles. This property may not contain a path for non-standard
// bundle formats or for some older bundle formats.
//
// See: https://developer.apple.com/documentation/Foundation/Bundle/sharedSupportURL
func (b Bundle) SharedSupportURL() INSURL {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("sharedSupportURL"))
	return NSURLFromID(objc.ID(rv))
}



// The file URL for the bundle’s App Store receipt.
//
// # Discussion
// 
// Use this app bundle property to locate the app receipt if it’s present;
// this property is `nil` if the receipt isn’t present. In the rare case a
// receipt is invalid or missing in an app that a user downloads from the App
// Store, use [SKReceiptRefreshRequest] to request a new receipt. For
// information about validating receipts, see [Choosing a receipt validation
// technique].
// 
// You can’t use the general best practice of weak linking using the
// [responds(to:)] method here; the method’s implementation uses the
// [doesNotRecognizeSelector(_:)] method.
// 
// # Get the receipt in testing environments
// 
// Receipts aren’t initially present in iOS and iPadOS apps in the sandbox
// environment and in Xcode. Apps get a receipt after the tester completes the
// first in-app purchase. When your app checks [AppStoreReceiptURL] and finds
// that it’s `nil`, assume the tester is a new customer and has no access to
// premium content. For Mac apps running in TestFlight, the receipt is always
// present.
//
// [Choosing a receipt validation technique]: https://developer.apple.com/documentation/StoreKit/choosing-a-receipt-validation-technique
// [SKReceiptRefreshRequest]: https://developer.apple.com/documentation/StoreKit/SKReceiptRefreshRequest
// [doesNotRecognizeSelector(_:)]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/doesNotRecognizeSelector(_:)
// [responds(to:)]: https://developer.apple.com/documentation/ObjectiveC/NSObjectProtocol/responds(to:)
//
// See: https://developer.apple.com/documentation/Foundation/Bundle/appStoreReceiptURL
func (b Bundle) AppStoreReceiptURL() INSURL {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("appStoreReceiptURL"))
	return NSURLFromID(objc.ID(rv))
}



// The full pathname of the bundle’s subdirectory containing resources.
//
// See: https://developer.apple.com/documentation/Foundation/Bundle/resourcePath
func (b Bundle) ResourcePath() string {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("resourcePath"))
	return NSStringFromID(rv).String()
}



// The full pathname of the receiver’s executable file.
//
// See: https://developer.apple.com/documentation/Foundation/Bundle/executablePath
func (b Bundle) ExecutablePath() string {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("executablePath"))
	return NSStringFromID(rv).String()
}



// The full pathname of the bundle’s subdirectory containing private
// frameworks.
//
// # Discussion
// 
// This property contains the appropriate path for modern application and
// framework bundles. This property may not contain a path for non-standard
// bundle formats or for some older bundle formats.
//
// See: https://developer.apple.com/documentation/Foundation/Bundle/privateFrameworksPath
func (b Bundle) PrivateFrameworksPath() string {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("privateFrameworksPath"))
	return NSStringFromID(rv).String()
}



// The full pathname of the bundle’s subdirectory containing shared
// frameworks.
//
// # Discussion
// 
// This property contains the appropriate path for modern application and
// framework bundles. This property may not contain a path for non-standard
// bundle formats or for some older bundle formats.
//
// See: https://developer.apple.com/documentation/Foundation/Bundle/sharedFrameworksPath
func (b Bundle) SharedFrameworksPath() string {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("sharedFrameworksPath"))
	return NSStringFromID(rv).String()
}



// The full pathname of the receiver’s subdirectory containing plug-ins.
//
// # Discussion
// 
// This is the appropriate path for modern application and framework bundles.
// This may not be a path for non-standard bundle formats or for some older
// bundle formats.
//
// See: https://developer.apple.com/documentation/Foundation/Bundle/builtInPlugInsPath
func (b Bundle) BuiltInPlugInsPath() string {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("builtInPlugInsPath"))
	return NSStringFromID(rv).String()
}



// The full pathname of the bundle’s subdirectory containing shared support
// files.
//
// # Discussion
// 
// This property contains the appropriate path for modern application and
// framework bundles. This property may not contain a path for non-standard
// bundle formats or for some older bundle formats.
//
// See: https://developer.apple.com/documentation/Foundation/Bundle/sharedSupportPath
func (b Bundle) SharedSupportPath() string {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("sharedSupportPath"))
	return NSStringFromID(rv).String()
}



// The full URL of the receiver’s bundle directory.
//
// See: https://developer.apple.com/documentation/Foundation/Bundle/bundleURL
func (b Bundle) BundleURL() INSURL {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("bundleURL"))
	return NSURLFromID(objc.ID(rv))
}



// The full pathname of the receiver’s bundle directory.
//
// See: https://developer.apple.com/documentation/Foundation/Bundle/bundlePath
func (b Bundle) BundlePath() string {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("bundlePath"))
	return NSStringFromID(rv).String()
}



// The receiver’s bundle identifier.
//
// # Discussion
// 
// The bundle identifier is defined by the [CFBundleIdentifier] key in the
// bundle’s information property list.
//
// See: https://developer.apple.com/documentation/Foundation/Bundle/bundleIdentifier
func (b Bundle) BundleIdentifier() string {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("bundleIdentifier"))
	return NSStringFromID(rv).String()
}



// A dictionary, constructed from the bundle’s `Info.Plist()` file, that
// contains information about the receiver.
//
// # Discussion
// 
// If the bundle does not contain an `Info.Plist()` file, this dictionary
// contains only private keys that are used internally by the [NSBundle]
// class. The [NSBundle] class may add extra keys to this dictionary for its
// own use. Common keys for accessing the values of the dictionary are
// [CFBundleIdentifier], [NSMainNibFile], and [NSPrincipalClass].
//
// See: https://developer.apple.com/documentation/Foundation/Bundle/infoDictionary
func (b Bundle) InfoDictionary() INSDictionary {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("infoDictionary"))
	return NSDictionaryFromID(objc.ID(rv))
}



// A list of all the localizations contained in the bundle.
//
// # Discussion
// 
// An array of [NSString] objects containing language IDs for all the
// localizations contained in the bundle.
//
// See: https://developer.apple.com/documentation/Foundation/Bundle/localizations
func (b Bundle) Localizations() []string {
	rv := objc.Send[[]objc.ID](b.ID, objc.Sel("localizations"))
	return objc.ConvertSliceToStrings(rv)
}



// An ordered list of preferred localizations contained in the bundle.
//
// # Discussion
// 
// An array of [NSString] objects containing language IDs for localizations in
// the bundle. The strings are ordered according to the user’s language
// preferences and available localizations.
//
// See: https://developer.apple.com/documentation/Foundation/Bundle/preferredLocalizations
func (b Bundle) PreferredLocalizations() []string {
	rv := objc.Send[[]objc.ID](b.ID, objc.Sel("preferredLocalizations"))
	return objc.ConvertSliceToStrings(rv)
}



// The localization for the development language.
//
// # Discussion
// 
// This property corresponds to the value in the [CFBundleDevelopmentRegion]
// key of the bundle’s property list (`Info.Plist()`).
//
// See: https://developer.apple.com/documentation/Foundation/Bundle/developmentLocalization
func (b Bundle) DevelopmentLocalization() string {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("developmentLocalization"))
	return NSStringFromID(rv).String()
}



// A dictionary with the keys from the bundle’s localized property list.
//
// # Discussion
// 
// This property uses the preferred localization for the current user when
// determining which resources to include. If the preferred localization is
// not available, this property chooses the most appropriate localization
// found in the bundle.
//
// See: https://developer.apple.com/documentation/Foundation/Bundle/localizedInfoDictionary
func (b Bundle) LocalizedInfoDictionary() INSDictionary {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("localizedInfoDictionary"))
	return NSDictionaryFromID(objc.ID(rv))
}



// The bundle’s principal class.
//
// # Discussion
// 
// This property is set after ensuring that the code containing the definition
// of the class is dynamically loaded. If the bundle encounters errors in
// loading or if it can’t find the executable code file in the bundle
// directory, this property is `nil`.
// 
// The principal class typically controls all the other classes in the bundle;
// it should mediate between those classes and classes external to the bundle.
// Classes (and categories) are loaded from just one file within the bundle
// directory. The bundle obtains the name of the code file to load from the
// dictionary returned from [InfoDictionary], using “[NSExecutable]” as
// the key. The bundle determines its principal class in one of two ways:
// 
// - It first looks in its own information dictionary, which extracts the
// information encoded in the bundle’s property list (`Info.Plist()`). The
// bundle obtains the principal class from the dictionary using the key
// [NSPrincipalClass]. For non-loadable bundles (applications and frameworks),
// if the principal class is not specified in the property list, this property
// is `nil`. - If the principal class is not specified in the information
// dictionary, the bundle identifies the first class loaded as the principal
// class. When several classes are linked into a dynamically loadable file,
// the default principal class is the first one listed on the `ld` command
// line. In the following example, Reporter would be the principal class:
// 
// The order of classes in Xcode’s project browser is the order in which
// they will be linked. To designate the principal class, control-drag the
// file containing its implementation to the top of the list.
// 
// As a side effect of code loading, the receiver posts [didLoadNotification]
// after all classes and categories have been loaded; see [Notifications] for
// details.
// 
// The following method obtains a bundle by specifying its path
// ([BundleWithPath]), then loads the bundle with [PrincipalClass] and uses
// the principal class object to allocate and initialize an instance of that
// class:
//
// [didLoadNotification]: https://developer.apple.com/documentation/Foundation/Bundle/didLoadNotification
//
// See: https://developer.apple.com/documentation/Foundation/Bundle/principalClass
func (b Bundle) PrincipalClass() objc.Class {
	rv := objc.Send[objc.Class](b.ID, objc.Sel("principalClass"))
	return rv
}



// A constant used as a key for the
//
// See: https://developer.apple.com/documentation/foundation/nsloadedclasses
func (b Bundle) NSLoadedClasses() string {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("NSLoadedClasses"))
	return NSStringFromID(rv).String()
}



// An array of numbers indicating the architecture types supported by the
// bundle’s executable.
//
// # Discussion
// 
// An array of [NSNumber] objects, each of which contains an integer value
// corresponding to a supported processor architecture. For a list of common
// architecture types, see the constants in [Mach-O Architecture]. If the
// bundle does not contain a Mach-O executable, this is `nil`.
// 
// The bundle scans its Mach-O executable and returns all of the architecture
// types it finds. Because they are taken directly from the executable, the
// values may not always correspond to one of the well-known CPU types defined
// in [Mach-O Architecture].
//
// [Mach-O Architecture]: https://developer.apple.com/documentation/Foundation/1495005-mach-o-architecture
//
// See: https://developer.apple.com/documentation/Foundation/Bundle/executableArchitectures
func (b Bundle) ExecutableArchitectures() []NSNumber {
	rv := objc.Send[[]objc.ID](b.ID, objc.Sel("executableArchitectures"))
	return objc.ConvertSlice(rv, func(id objc.ID) NSNumber {
		return NSNumberFromID(id)
	})
}



// The load status of a bundle.
//
// # Discussion
// 
// [true] if the bundle’s code is currently loaded, otherwise [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Foundation/Bundle/isLoaded
func (b Bundle) Loaded() bool {
	rv := objc.Send[bool](b.ID, objc.Sel("isLoaded"))
	return rv
}



// The beginning of the range of error codes reserved for errors related to
// executable files.
//
// See: https://developer.apple.com/documentation/foundation/nsexecutableerrorminimum-swift.var
func (b Bundle) NSExecutableErrorMinimum() int {
	rv := objc.Send[int](b.ID, objc.Sel("NSExecutableErrorMinimum"))
	return rv
}
func (b Bundle) SetNSExecutableErrorMinimum(value int) {
	objc.Send[struct{}](b.ID, objc.Sel("setNSExecutableErrorMinimum:"), value)
}



// The executable type isn’t loadable in the current process.
//
// See: https://developer.apple.com/documentation/foundation/nsexecutablenotloadableerror-swift.var
func (b Bundle) NSExecutableNotLoadableError() int {
	rv := objc.Send[int](b.ID, objc.Sel("NSExecutableNotLoadableError"))
	return rv
}
func (b Bundle) SetNSExecutableNotLoadableError(value int) {
	objc.Send[struct{}](b.ID, objc.Sel("setNSExecutableNotLoadableError:"), value)
}



// The executable doesn’t provide an architecture compatible with the
// current process.
//
// See: https://developer.apple.com/documentation/foundation/nsexecutablearchitecturemismatcherror-swift.var
func (b Bundle) NSExecutableArchitectureMismatchError() int {
	rv := objc.Send[int](b.ID, objc.Sel("NSExecutableArchitectureMismatchError"))
	return rv
}
func (b Bundle) SetNSExecutableArchitectureMismatchError(value int) {
	objc.Send[struct{}](b.ID, objc.Sel("setNSExecutableArchitectureMismatchError:"), value)
}



// The executable has Objective-C runtime information that’s incompatible
// with the current process.
//
// See: https://developer.apple.com/documentation/foundation/nsexecutableruntimemismatcherror-swift.var
func (b Bundle) NSExecutableRuntimeMismatchError() int {
	rv := objc.Send[int](b.ID, objc.Sel("NSExecutableRuntimeMismatchError"))
	return rv
}
func (b Bundle) SetNSExecutableRuntimeMismatchError(value int) {
	objc.Send[struct{}](b.ID, objc.Sel("setNSExecutableRuntimeMismatchError:"), value)
}



// Executable cannot be loaded for an otherwise-unspecified reason.
//
// See: https://developer.apple.com/documentation/foundation/nsexecutableloaderror-swift.var
func (b Bundle) NSExecutableLoadError() int {
	rv := objc.Send[int](b.ID, objc.Sel("NSExecutableLoadError"))
	return rv
}
func (b Bundle) SetNSExecutableLoadError(value int) {
	objc.Send[struct{}](b.ID, objc.Sel("setNSExecutableLoadError:"), value)
}



// The executable failed due to linking issues.
//
// See: https://developer.apple.com/documentation/foundation/nsexecutablelinkerror-swift.var
func (b Bundle) NSExecutableLinkError() int {
	rv := objc.Send[int](b.ID, objc.Sel("NSExecutableLinkError"))
	return rv
}
func (b Bundle) SetNSExecutableLinkError(value int) {
	objc.Send[struct{}](b.ID, objc.Sel("setNSExecutableLinkError:"), value)
}



// The end of the range of error codes reserved for errors related to
// executable files.
//
// See: https://developer.apple.com/documentation/foundation/nsexecutableerrormaximum-swift.var
func (b Bundle) NSExecutableErrorMaximum() int {
	rv := objc.Send[int](b.ID, objc.Sel("NSExecutableErrorMaximum"))
	return rv
}
func (b Bundle) SetNSExecutableErrorMaximum(value int) {
	objc.Send[struct{}](b.ID, objc.Sel("setNSExecutableErrorMaximum:"), value)
}







// Returns the bundle object that contains the current executable.
//
// # Return Value
// 
// The [NSBundle] object corresponding to the bundle directory that contains
// the current executable. This method may return a valid bundle object even
// for unbundled apps. It may also return `nil` if the bundle object could not
// be created, so always check the return value.
// 
// # Discussion
// 
// The main bundle lets you access the resources in the same directory as the
// currently running executable. For a running app or code running in a
// framework, the main bundle offers access to the app’s bundle directory.
//
// See: https://developer.apple.com/documentation/Foundation/Bundle/main
func (_BundleClass BundleClass) MainBundle() Bundle {
	rv := objc.Send[objc.ID](objc.ID(_BundleClass.class), objc.Sel("mainBundle"))
	return NSBundleFromID(objc.ID(rv))
}



// Returns an array of all of the application’s bundles that represent
// frameworks.
//
// # Return Value
// 
// An array of all of the application’s bundles that represent frameworks.
// Only frameworks with one or more Objective-C classes in them are included.
// 
// # Discussion
// 
// The returned array includes frameworks that are linked into an application
// when the application is built and bundles for frameworks that have been
// dynamically created.
//
// See: https://developer.apple.com/documentation/Foundation/Bundle/allFrameworks
func (_BundleClass BundleClass) AllFrameworks() []Bundle {
	rv := objc.Send[[]objc.ID](objc.ID(_BundleClass.class), objc.Sel("allFrameworks"))
	return objc.ConvertSlice(rv, func(id objc.ID) NSBundle {
		return NSBundleFromID(id)
	})
}



// Returns an array of all the application’s non-framework bundles.
//
// # Return Value
// 
// An array of all the application’s non-framework bundles.
// 
// # Discussion
// 
// The returned array includes the main bundle and all bundles that have been
// dynamically created but doesn’t contain any bundles that represent
// frameworks.
//
// See: https://developer.apple.com/documentation/Foundation/Bundle/allBundles
func (_BundleClass BundleClass) AllBundles() []Bundle {
	rv := objc.Send[[]objc.ID](objc.ID(_BundleClass.class), objc.Sel("allBundles"))
	return objc.ConvertSlice(rv, func(id objc.ID) NSBundle {
		return NSBundleFromID(id)
	})
}




















