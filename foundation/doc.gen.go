// Code generated from Apple documentation for Foundation. DO NOT EDIT.

// Package foundation provides Go bindings for the Foundation framework.
//
// Access essential data types, collections, and operating-system services to
// define the base layer of functionality for your app.
//
// The Foundation framework provides a base layer of functionality for apps
// and frameworks, including data storage and persistence, text processing,
// date and time calculations, sorting and filtering, and networking. The
// classes, protocols, and data types defined by Foundation are used
// throughout the macOS, iOS, watchOS, and tvOS SDKs.
//
// # Fundamentals
//
//   - [Numbers, Data, and Basic Values]: Work with primitive values and other fundamental types used throughout Cocoa. ([Decimal], [NumberFormatter], [Data], [URL], [URLComponents])
//   - [Strings and Text]: Create and process strings of Unicode characters, use regular expressions to find patterns, and perform natural language analysis of text. ([String], [AttributedString], [NSAttributedString], [NSMutableAttributedString], [CharacterSet])
//   - Collections: Use arrays, dictionaries, sets, and specialized collections to store and iterate groups of objects or values. ([Array], [Dictionary], [Set], [IndexPath], [IndexSet])
//   - [Dates and Times]: Compare dates and times, and perform calendar and time zone calculations. ([Date], [DateInterval], [TimeInterval], [DateComponents], [Calendar])
//   - [Units and Measurement]: Label numeric quantities with physical dimensions to allow locale-aware formatting and conversion between related units. ([Measurement], [NSMeasurement], [Unit], [Dimension], [UnitConverter])
//   - [Data Formatting]: Convert numbers, dates, measurements, and other values to and from locale-aware string representations. ([NumberFormatter], [PersonNameComponentsFormatter], [PersonNameComponents], [DateFormatter], [DateComponentsFormatter])
//   - [Filters and Sorting]: Use predicates, expressions, and sort descriptors to examine elements in collections and other services. ([Predicate], [NSPredicate], [NSExpression], [NSComparisonPredicate], [NSCompoundPredicate])
//
// # App Support
//
//   - [Task Management]: Manage your app’s work and how it interacts with system services like Handoff and Shortcuts. ([UndoManager], [ProgressReporting], [Progress], [Operation], [OperationQueue])
//   - Resources: Access assets and other data bundled with your app. ([Bundle])
//   - Notifications: Design patterns for broadcasting information and for subscribing to broadcasts. ([Notification], [NotificationCenter], [NotificationQueue], [DistributedNotificationCenter])
//   - [App Extension Support]: Manage the interaction between an app extension and its hosting app. ([NSExtensionRequestHandling], [NSExtensionContext], [NSItemProvider], [NSExtensionItem], [NSUserActivity])
//   - [Errors and Exceptions]: Respond to problem situations in your interactions with APIs, and fine-tune your app for better debugging. ([Error], [NSError], [NSAssertionHandler], [NSException])
//   - [Scripting Support]: Allow users to control your app with AppleScript and other automation technologies, or run scripts from within your app. ([NSAppleScript], [NSAppleEventDescriptor], [NSAppleEventManager], [NSScriptCommand], [NSQuitCommand])
//
// # Files and Data Persistence
//
//   - [File System]: Create, read, write, and examine files and folders in the file system. ([FileManager], [FileManagerDelegate], [NSFilePresenter], [NSFileAccessIntent], [NSFileCoordinator])
//   - [Archives and Serialization]: Convert objects and values to and from property list, JSON, and other flat binary representations. ([NSCoding], [NSSecureCoding], [JSONSerialization], [PropertyListSerialization], [NSKeyedArchiver])
//   - Settings: Configure your app using data you store persistently on the local disk or in iCloud. ([UserDefaults], [NSUbiquitousKeyValueStore])
//   - Spotlight: Search for files and other items on the local device, and index your app’s content for searching. ([NSMetadataQuery], [NSMetadataQueryDelegate], [NSMetadataItem])
//   - iCloud: Manage files and key-value data that automatically synchronize among a user’s iCloud devices. ([FileManager], [FileManagerDelegate], [NSUbiquitousKeyValueStore], [NSMetadataQuery], [NSMetadataQueryDelegate])
//   - [Optimizing Your App’s Data for iCloud Backup]: Minimize the space and time that backups take to create by excluding purgeable and nonpurgeable data from backups.
//
// # Networking
//
//   - [URL Loading System]: Interact with URLs and communicate with servers using standard Internet protocols. ([URLSession], [URLSessionTask], [URLRequest], [NSURLRequest], [NSMutableURLRequest])
//   - Bonjour: Advertise services for easy discovery on local networks, or discover services advertised by others. ([NetService], [NetServiceDelegate], [NetServiceBrowser], [NetServiceBrowserDelegate])
//
// # Low-Level Utilities
//
//   - XPC: Manage secure interprocess communication. ([NSXPCProxyCreating], [NSXPCConnection], [NSXPCInterface], [NSXPCCoder], [NSXPCListener])
//   - [Object Runtime]: Get low-level support for basic Objective-C features, Cocoa design patterns, and Swift integration. ([NSCopying], [NSMutableCopying], [NSNumber], [NSValue], [ValueTransformer])
//   - [Processes and Threads]: Manage your app’s interaction with the host operating system and other processes, and implement low-level concurrency features. ([RunLoop], [Timer], [ProcessInfo], [Thread], [NSLocking])
//   - [Streams, Sockets, and Ports]: Use low-level Unix features to manage input and output among files, processes, and the network. ([Stream], [InputStream], [OutputStream], [StreamDelegate], [Pipe])
//
// # Variables
//
//   - [ListItemDelimiterAttributeName]
//   - [CalendarIdentifierBangla]
//   - [CalendarIdentifierDangi]
//   - [CalendarIdentifierGujarati]
//   - [CalendarIdentifierKannada]
//   - [CalendarIdentifierMalayalam]
//   - [CalendarIdentifierMarathi]
//   - [CalendarIdentifierOdia]
//   - [CalendarIdentifierTamil]
//   - [CalendarIdentifierTelugu]
//   - [CalendarIdentifierVietnamese]
//   - [CalendarIdentifierVikram]
//   - [URLUbiquitousItemIsSyncPausedKey]
//   - [URLUbiquitousItemSupportedSyncControlsKey]
//
// # Macros
//
//   - NS_SWIFT_NONISOLATED_UNSAFE//
//
// # Key Types
//
//   - [NSString] - A static, plain-text Unicode string object.
//   - [NSCoder] - An abstract class that serves as the basis for objects that enable archiving and distribution of other objects.
//   - [NSURL] - An object that represents the location of a resource, such as an item on a remote server or the path to a local file.
//   - [NSAttributedString] - A string of text that manages data, layout, and stylistic information for ranges of characters to support rendering.
//   - [NSArray] - A static ordered collection of objects.
//   - [NSNumber] - An object wrapper for primitive scalar numeric values.
//   - [NSDecimalNumber] - An object for representing and performing arithmetic on base-10 numbers.
//   - [NSNumberFormatter] - A formatter that converts between numeric values and their textual representations.
//   - [NSOrderedSet] - A static, ordered collection of unique objects.
//   - [NSFileManager] - A convenient interface to the contents of the file system, and the primary means of interacting with it.
//
// [App Extension Support]: https://developer.apple.com/documentation/foundation/app-extension-support
// [Archives and Serialization]: https://developer.apple.com/documentation/foundation/archives-and-serialization
// [Data Formatting]: https://developer.apple.com/documentation/foundation/data-formatting
// [Dates and Times]: https://developer.apple.com/documentation/foundation/dates-and-times
// [Errors and Exceptions]: https://developer.apple.com/documentation/foundation/errors-and-exceptions
// [File System]: https://developer.apple.com/documentation/foundation/file-system
// [Filters and Sorting]: https://developer.apple.com/documentation/foundation/filters-and-sorting
// [Numbers, Data, and Basic Values]: https://developer.apple.com/documentation/foundation/numbers-data-and-basic-values
// [Object Runtime]: https://developer.apple.com/documentation/foundation/object-runtime
// [Optimizing Your App’s Data for iCloud Backup]: https://developer.apple.com/documentation/foundation/optimizing-your-app-s-data-for-icloud-backup
// [Processes and Threads]: https://developer.apple.com/documentation/foundation/processes-and-threads
// [Scripting Support]: https://developer.apple.com/documentation/foundation/scripting-support
// [Streams, Sockets, and Ports]: https://developer.apple.com/documentation/foundation/streams-sockets-and-ports
// [Strings and Text]: https://developer.apple.com/documentation/foundation/strings-and-text
// [Task Management]: https://developer.apple.com/documentation/foundation/task-management
// [URL Loading System]: https://developer.apple.com/documentation/foundation/url-loading-system
// [Units and Measurement]: https://developer.apple.com/documentation/foundation/units-and-measurement
package foundation

import (
	"fmt"
	"os"

	"github.com/ebitengine/purego"
)

// frameworkPaths lists paths to try when loading the Foundation library.
// The framework bundle path is tried first; a /usr/lib dylib fallback covers
// C-API frameworks that are not in the dyld shared cache as bundles.
var frameworkPaths = []string{
	"/System/Library/Frameworks/Foundation.framework/Foundation",
	"/usr/lib/libFoundation.dylib",
}

// frameworkHandle is the handle to the loaded framework.
var frameworkHandle uintptr

func init() {
	for _, path := range frameworkPaths {
		h, err := purego.Dlopen(path, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
		if err == nil {
			frameworkHandle = h
			return
		}
	}
	fmt.Fprintf(os.Stderr, "warning: Foundation: failed to load framework from any known path\n")
}
