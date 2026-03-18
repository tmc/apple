// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"context"
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"errors"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
	"github.com/tmc/apple/uniformtypeidentifiers"
)

// The class instance for the [NSWorkspace] class.
var (
	_NSWorkspaceClass     NSWorkspaceClass
	_NSWorkspaceClassOnce sync.Once
)

func getNSWorkspaceClass() NSWorkspaceClass {
	_NSWorkspaceClassOnce.Do(func() {
		_NSWorkspaceClass = NSWorkspaceClass{class: objc.GetClass("NSWorkspace")}
	})
	return _NSWorkspaceClass
}

// GetNSWorkspaceClass returns the class object for NSWorkspace.
func GetNSWorkspaceClass() NSWorkspaceClass {
	return getNSWorkspaceClass()
}

type NSWorkspaceClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSWorkspaceClass) Alloc() NSWorkspace {
	rv := objc.Send[NSWorkspace](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A workspace that can launch other apps and perform a variety of
// file-handling services.
//
// # Overview
// 
// There is one shared [NSWorkspace] object per app. You use the class method
// [NSWorkspace.SharedWorkspace] to access it. For example, the following statement uses
// an [NSWorkspace] object to request that a file be opened in the TextEdit
// app:
// 
// You can use the workspace object to:
// 
// - Open, manipulate, and get information about files and devices. - Track
// changes to the file system, devices, and the user database. - Get and set
// Finder information for files. - Launch apps.
//
// # Accessing the Workspace Notification Center
//
//   - [NSWorkspace.NotificationCenter]: The notification center for workspace notifications.
//
// # Opening URLs
//
//   - [NSWorkspace.OpenURLConfigurationCompletionHandler]: Opens a URL asynchronously using the provided options.
//   - [NSWorkspace.OpenURLsWithApplicationAtURLConfigurationCompletionHandler]: Opens one or more URLs asynchronously in the specified app using the provided options.
//   - [NSWorkspace.OpenURL]: Opens the location at the specified URL.
//
// # Launching and Hiding Apps
//
//   - [NSWorkspace.OpenApplicationAtURLConfigurationCompletionHandler]: Launches the app at the specified URL and asynchronously reports back on the app’s status.
//   - [NSWorkspace.HideOtherApplications]: Hides all applications other than the sender.
//
// # Manipulating Files
//
//   - [NSWorkspace.ActivateFileViewerSelectingURLs]: Activates the Finder, and opens one or more windows selecting the specified files.
//   - [NSWorkspace.SelectFileInFileViewerRootedAtPath]: Selects the file at the specified path.
//
// # Requesting Information
//
//   - [NSWorkspace.URLForApplicationToOpenURL]: Returns the URL to the default app to open the specified URL.
//   - [NSWorkspace.URLForApplicationToOpenContentType]: Returns the URL to the default app to open the specified content type.
//   - [NSWorkspace.URLForApplicationWithBundleIdentifier]: Returns the URL to the default app with the specified bundle identifier.
//   - [NSWorkspace.URLsForApplicationsToOpenURL]: Returns an array of URLs to all available applications that can open the URL.
//   - [NSWorkspace.URLsForApplicationsToOpenContentType]: Returns an array of URLs to all available applications that can open the specified content type.
//   - [NSWorkspace.URLsForApplicationsWithBundleIdentifier]: Returns an array of URLs to all available applications that can open the specified bundle identifier.
//   - [NSWorkspace.GetFileSystemInfoForPathIsRemovableIsWritableIsUnmountableDescriptionType]: Returns information about the file system at the specified path.
//   - [NSWorkspace.IsFilePackageAtPath]: Determines whether the specified path is a file package.
//   - [NSWorkspace.FrontmostApplication]: Returns the frontmost app, which is the app that receives key events.
//   - [NSWorkspace.RunningApplications]: Returns an array of running apps.
//   - [NSWorkspace.MenuBarOwningApplication]: Returns the app that owns the currently displayed menu bar.
//
// # Setting Default Application Information
//
//   - [NSWorkspace.SetDefaultApplicationAtURLToOpenFileAtURLCompletionHandler]: Sets the default app to use when opening a specific file.
//   - [NSWorkspace.SetDefaultApplicationAtURLToOpenContentTypeCompletionHandler]: Sets the default app to use when opening files of a specific content type.
//   - [NSWorkspace.SetDefaultApplicationAtURLToOpenContentTypeOfFileAtURLCompletionHandler]: Sets the default app to use when opening files of a specific content type defined by a file URL.
//   - [NSWorkspace.SetDefaultApplicationAtURLToOpenURLsWithSchemeCompletionHandler]: Sets the default app to use when opening files of a specific scheme.
//
// # Managing Icons
//
//   - [NSWorkspace.IconForFile]: Returns an image containing the icon for the specified file.
//   - [NSWorkspace.IconForFiles]: Returns an image containing the icon for the specified files.
//   - [NSWorkspace.IconForContentType]: Returns an image containing the icon for the specified content type.
//   - [NSWorkspace.SetIconForFileOptions]: Sets the icon for the file or directory at the specified path.
//
// # Unmounting a Device
//
//   - [NSWorkspace.UnmountAndEjectDeviceAtPath]: Unmounts and ejects the device at the specified path.
//   - [NSWorkspace.UnmountAndEjectDeviceAtURLError]: Attempts to eject the volume mounted at the given path.
//
// # Managing the Desktop Image
//
//   - [NSWorkspace.DesktopImageURLForScreen]: Returns the URL for the desktop image for the given screen.
//   - [NSWorkspace.SetDesktopImageURLForScreenOptionsError]: Sets the desktop image for the given screen to the image at the specified URL.
//   - [NSWorkspace.DesktopImageOptionsForScreen]: Returns the desktop image options for the given screen.
//
// # Performing Finder Spotlight Searches
//
//   - [NSWorkspace.ShowSearchResultsForQueryString]: Displays a Spotlight search results window in Finder for the specified query string.
//
// # Finder File Labels
//
//   - [NSWorkspace.FileLabels]: The array of file labels, returned as strings.
//   - [NSWorkspace.FileLabelColors]: The array of colors for the file labels.
//
// # Tracking Changes to the File System
//
//   - [NSWorkspace.NoteFileSystemChanged]: Informs the workspace object that the file system changed at the specified path.
//
// # Requesting Additional Time Before Logout
//
//   - [NSWorkspace.ExtendPowerOffBy]: Requests the system wait for the specified amount of time before turning off the power or logging out the user.
//
// # Supporting Accessibility
//
//   - [NSWorkspace.AccessibilityDisplayShouldDifferentiateWithoutColor]: A Boolean value that indicates whether the app avoids conveying information through color alone.
//   - [NSWorkspace.AccessibilityDisplayShouldIncreaseContrast]: A Boolean value that indicates whether the app presents a high-contrast user interface.
//   - [NSWorkspace.AccessibilityDisplayShouldReduceTransparency]: A Boolean value that indicates whether the app avoids using semitransparent backgrounds.
//   - [NSWorkspace.AccessibilityDisplayShouldInvertColors]: A Boolean value that indicates whether the accessibility option to invert colors is in an enabled state.
//   - [NSWorkspace.AccessibilityDisplayShouldReduceMotion]: A Boolean value that indicates whether the accessibility option to reduce motion is in an enabled state.
//   - [NSWorkspace.SwitchControlEnabled]: A Boolean value that indicates whether Switch Control is currently running.
//   - [NSWorkspace.VoiceOverEnabled]: A Boolean value that indicates whether VoiceOver is currently running.
//
// # Performing Privileged Operations
//
//   - [NSWorkspace.RequestAuthorizationOfTypeCompletionHandler]: Requests authorization to perform a privileged file operation.
//
// See: https://developer.apple.com/documentation/AppKit/NSWorkspace
type NSWorkspace struct {
	objectivec.Object
}

// NSWorkspaceFromID constructs a [NSWorkspace] from an objc.ID.
//
// A workspace that can launch other apps and perform a variety of
// file-handling services.
func NSWorkspaceFromID(id objc.ID) NSWorkspace {
	return NSWorkspace{objectivec.Object{ID: id}}
}
// NOTE: NSWorkspace adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSWorkspace] class.
//
// # Accessing the Workspace Notification Center
//
//   - [INSWorkspace.NotificationCenter]: The notification center for workspace notifications.
//
// # Opening URLs
//
//   - [INSWorkspace.OpenURLConfigurationCompletionHandler]: Opens a URL asynchronously using the provided options.
//   - [INSWorkspace.OpenURLsWithApplicationAtURLConfigurationCompletionHandler]: Opens one or more URLs asynchronously in the specified app using the provided options.
//   - [INSWorkspace.OpenURL]: Opens the location at the specified URL.
//
// # Launching and Hiding Apps
//
//   - [INSWorkspace.OpenApplicationAtURLConfigurationCompletionHandler]: Launches the app at the specified URL and asynchronously reports back on the app’s status.
//   - [INSWorkspace.HideOtherApplications]: Hides all applications other than the sender.
//
// # Manipulating Files
//
//   - [INSWorkspace.ActivateFileViewerSelectingURLs]: Activates the Finder, and opens one or more windows selecting the specified files.
//   - [INSWorkspace.SelectFileInFileViewerRootedAtPath]: Selects the file at the specified path.
//
// # Requesting Information
//
//   - [INSWorkspace.URLForApplicationToOpenURL]: Returns the URL to the default app to open the specified URL.
//   - [INSWorkspace.URLForApplicationToOpenContentType]: Returns the URL to the default app to open the specified content type.
//   - [INSWorkspace.URLForApplicationWithBundleIdentifier]: Returns the URL to the default app with the specified bundle identifier.
//   - [INSWorkspace.URLsForApplicationsToOpenURL]: Returns an array of URLs to all available applications that can open the URL.
//   - [INSWorkspace.URLsForApplicationsToOpenContentType]: Returns an array of URLs to all available applications that can open the specified content type.
//   - [INSWorkspace.URLsForApplicationsWithBundleIdentifier]: Returns an array of URLs to all available applications that can open the specified bundle identifier.
//   - [INSWorkspace.GetFileSystemInfoForPathIsRemovableIsWritableIsUnmountableDescriptionType]: Returns information about the file system at the specified path.
//   - [INSWorkspace.IsFilePackageAtPath]: Determines whether the specified path is a file package.
//   - [INSWorkspace.FrontmostApplication]: Returns the frontmost app, which is the app that receives key events.
//   - [INSWorkspace.RunningApplications]: Returns an array of running apps.
//   - [INSWorkspace.MenuBarOwningApplication]: Returns the app that owns the currently displayed menu bar.
//
// # Setting Default Application Information
//
//   - [INSWorkspace.SetDefaultApplicationAtURLToOpenFileAtURLCompletionHandler]: Sets the default app to use when opening a specific file.
//   - [INSWorkspace.SetDefaultApplicationAtURLToOpenContentTypeCompletionHandler]: Sets the default app to use when opening files of a specific content type.
//   - [INSWorkspace.SetDefaultApplicationAtURLToOpenContentTypeOfFileAtURLCompletionHandler]: Sets the default app to use when opening files of a specific content type defined by a file URL.
//   - [INSWorkspace.SetDefaultApplicationAtURLToOpenURLsWithSchemeCompletionHandler]: Sets the default app to use when opening files of a specific scheme.
//
// # Managing Icons
//
//   - [INSWorkspace.IconForFile]: Returns an image containing the icon for the specified file.
//   - [INSWorkspace.IconForFiles]: Returns an image containing the icon for the specified files.
//   - [INSWorkspace.IconForContentType]: Returns an image containing the icon for the specified content type.
//   - [INSWorkspace.SetIconForFileOptions]: Sets the icon for the file or directory at the specified path.
//
// # Unmounting a Device
//
//   - [INSWorkspace.UnmountAndEjectDeviceAtPath]: Unmounts and ejects the device at the specified path.
//   - [INSWorkspace.UnmountAndEjectDeviceAtURLError]: Attempts to eject the volume mounted at the given path.
//
// # Managing the Desktop Image
//
//   - [INSWorkspace.DesktopImageURLForScreen]: Returns the URL for the desktop image for the given screen.
//   - [INSWorkspace.SetDesktopImageURLForScreenOptionsError]: Sets the desktop image for the given screen to the image at the specified URL.
//   - [INSWorkspace.DesktopImageOptionsForScreen]: Returns the desktop image options for the given screen.
//
// # Performing Finder Spotlight Searches
//
//   - [INSWorkspace.ShowSearchResultsForQueryString]: Displays a Spotlight search results window in Finder for the specified query string.
//
// # Finder File Labels
//
//   - [INSWorkspace.FileLabels]: The array of file labels, returned as strings.
//   - [INSWorkspace.FileLabelColors]: The array of colors for the file labels.
//
// # Tracking Changes to the File System
//
//   - [INSWorkspace.NoteFileSystemChanged]: Informs the workspace object that the file system changed at the specified path.
//
// # Requesting Additional Time Before Logout
//
//   - [INSWorkspace.ExtendPowerOffBy]: Requests the system wait for the specified amount of time before turning off the power or logging out the user.
//
// # Supporting Accessibility
//
//   - [INSWorkspace.AccessibilityDisplayShouldDifferentiateWithoutColor]: A Boolean value that indicates whether the app avoids conveying information through color alone.
//   - [INSWorkspace.AccessibilityDisplayShouldIncreaseContrast]: A Boolean value that indicates whether the app presents a high-contrast user interface.
//   - [INSWorkspace.AccessibilityDisplayShouldReduceTransparency]: A Boolean value that indicates whether the app avoids using semitransparent backgrounds.
//   - [INSWorkspace.AccessibilityDisplayShouldInvertColors]: A Boolean value that indicates whether the accessibility option to invert colors is in an enabled state.
//   - [INSWorkspace.AccessibilityDisplayShouldReduceMotion]: A Boolean value that indicates whether the accessibility option to reduce motion is in an enabled state.
//   - [INSWorkspace.SwitchControlEnabled]: A Boolean value that indicates whether Switch Control is currently running.
//   - [INSWorkspace.VoiceOverEnabled]: A Boolean value that indicates whether VoiceOver is currently running.
//
// # Performing Privileged Operations
//
//   - [INSWorkspace.RequestAuthorizationOfTypeCompletionHandler]: Requests authorization to perform a privileged file operation.
//
// See: https://developer.apple.com/documentation/AppKit/NSWorkspace
type INSWorkspace interface {
	objectivec.IObject

	// Topic: Accessing the Workspace Notification Center

	// The notification center for workspace notifications.
	NotificationCenter() foundation.NSNotificationCenter

	// Topic: Opening URLs

	// Opens a URL asynchronously using the provided options.
	OpenURLConfigurationCompletionHandler(url foundation.INSURL, configuration INSWorkspaceOpenConfiguration, completionHandler RunningApplicationErrorHandler)
	// Opens one or more URLs asynchronously in the specified app using the provided options.
	OpenURLsWithApplicationAtURLConfigurationCompletionHandler(urls []foundation.NSURL, applicationURL foundation.INSURL, configuration INSWorkspaceOpenConfiguration, completionHandler RunningApplicationErrorHandler)
	// Opens the location at the specified URL.
	OpenURL(url foundation.INSURL) bool

	// Topic: Launching and Hiding Apps

	// Launches the app at the specified URL and asynchronously reports back on the app’s status.
	OpenApplicationAtURLConfigurationCompletionHandler(applicationURL foundation.INSURL, configuration INSWorkspaceOpenConfiguration, completionHandler RunningApplicationErrorHandler)
	// Hides all applications other than the sender.
	HideOtherApplications()

	// Topic: Manipulating Files

	// Activates the Finder, and opens one or more windows selecting the specified files.
	ActivateFileViewerSelectingURLs(fileURLs []foundation.NSURL)
	// Selects the file at the specified path.
	SelectFileInFileViewerRootedAtPath(fullPath string, rootFullPath string) bool

	// Topic: Requesting Information

	// Returns the URL to the default app to open the specified URL.
	URLForApplicationToOpenURL(url foundation.INSURL) foundation.INSURL
	// Returns the URL to the default app to open the specified content type.
	URLForApplicationToOpenContentType(contentType uniformtypeidentifiers.UTType) foundation.INSURL
	// Returns the URL to the default app with the specified bundle identifier.
	URLForApplicationWithBundleIdentifier(bundleIdentifier string) foundation.INSURL
	// Returns an array of URLs to all available applications that can open the URL.
	URLsForApplicationsToOpenURL(url foundation.INSURL) []foundation.NSURL
	// Returns an array of URLs to all available applications that can open the specified content type.
	URLsForApplicationsToOpenContentType(contentType uniformtypeidentifiers.UTType) []foundation.NSURL
	// Returns an array of URLs to all available applications that can open the specified bundle identifier.
	URLsForApplicationsWithBundleIdentifier(bundleIdentifier string) []foundation.NSURL
	// Returns information about the file system at the specified path.
	GetFileSystemInfoForPathIsRemovableIsWritableIsUnmountableDescriptionType(fullPath string, description string, fileSystemType string) (bool, bool, bool, bool)
	// Determines whether the specified path is a file package.
	IsFilePackageAtPath(fullPath string) bool
	// Returns the frontmost app, which is the app that receives key events.
	FrontmostApplication() INSRunningApplication
	// Returns an array of running apps.
	RunningApplications() []NSRunningApplication
	// Returns the app that owns the currently displayed menu bar.
	MenuBarOwningApplication() INSRunningApplication

	// Topic: Setting Default Application Information

	// Sets the default app to use when opening a specific file.
	SetDefaultApplicationAtURLToOpenFileAtURLCompletionHandler(applicationURL foundation.INSURL, url foundation.INSURL, completionHandler ErrorHandler)
	// Sets the default app to use when opening files of a specific content type.
	SetDefaultApplicationAtURLToOpenContentTypeCompletionHandler(applicationURL foundation.INSURL, contentType uniformtypeidentifiers.UTType, completionHandler ErrorHandler)
	// Sets the default app to use when opening files of a specific content type defined by a file URL.
	SetDefaultApplicationAtURLToOpenContentTypeOfFileAtURLCompletionHandler(applicationURL foundation.INSURL, url foundation.INSURL, completionHandler ErrorHandler)
	// Sets the default app to use when opening files of a specific scheme.
	SetDefaultApplicationAtURLToOpenURLsWithSchemeCompletionHandler(applicationURL foundation.INSURL, urlScheme string, completionHandler ErrorHandler)

	// Topic: Managing Icons

	// Returns an image containing the icon for the specified file.
	IconForFile(fullPath string) INSImage
	// Returns an image containing the icon for the specified files.
	IconForFiles(fullPaths []string) INSImage
	// Returns an image containing the icon for the specified content type.
	IconForContentType(contentType uniformtypeidentifiers.UTType) INSImage
	// Sets the icon for the file or directory at the specified path.
	SetIconForFileOptions(image INSImage, fullPath string, options NSWorkspaceIconCreationOptions) bool

	// Topic: Unmounting a Device

	// Unmounts and ejects the device at the specified path.
	UnmountAndEjectDeviceAtPath(path string) bool
	// Attempts to eject the volume mounted at the given path.
	UnmountAndEjectDeviceAtURLError(url foundation.INSURL) (bool, error)

	// Topic: Managing the Desktop Image

	// Returns the URL for the desktop image for the given screen.
	DesktopImageURLForScreen(screen INSScreen) foundation.INSURL
	// Sets the desktop image for the given screen to the image at the specified URL.
	SetDesktopImageURLForScreenOptionsError(url foundation.INSURL, screen INSScreen, options foundation.INSDictionary) (bool, error)
	// Returns the desktop image options for the given screen.
	DesktopImageOptionsForScreen(screen INSScreen) foundation.INSDictionary

	// Topic: Performing Finder Spotlight Searches

	// Displays a Spotlight search results window in Finder for the specified query string.
	ShowSearchResultsForQueryString(queryString string) bool

	// Topic: Finder File Labels

	// The array of file labels, returned as strings.
	FileLabels() []string
	// The array of colors for the file labels.
	FileLabelColors() []NSColor

	// Topic: Tracking Changes to the File System

	// Informs the workspace object that the file system changed at the specified path.
	NoteFileSystemChanged(path string)

	// Topic: Requesting Additional Time Before Logout

	// Requests the system wait for the specified amount of time before turning off the power or logging out the user.
	ExtendPowerOffBy(requested int) int

	// Topic: Supporting Accessibility

	// A Boolean value that indicates whether the app avoids conveying information through color alone.
	AccessibilityDisplayShouldDifferentiateWithoutColor() bool
	// A Boolean value that indicates whether the app presents a high-contrast user interface.
	AccessibilityDisplayShouldIncreaseContrast() bool
	// A Boolean value that indicates whether the app avoids using semitransparent backgrounds.
	AccessibilityDisplayShouldReduceTransparency() bool
	// A Boolean value that indicates whether the accessibility option to invert colors is in an enabled state.
	AccessibilityDisplayShouldInvertColors() bool
	// A Boolean value that indicates whether the accessibility option to reduce motion is in an enabled state.
	AccessibilityDisplayShouldReduceMotion() bool
	// A Boolean value that indicates whether Switch Control is currently running.
	SwitchControlEnabled() bool
	// A Boolean value that indicates whether VoiceOver is currently running.
	VoiceOverEnabled() bool

	// Topic: Performing Privileged Operations

	// Requests authorization to perform a privileged file operation.
	RequestAuthorizationOfTypeCompletionHandler(type_ NSWorkspaceAuthorizationType, completionHandler WorkspaceAuthorizationErrorHandler)
}

// Init initializes the instance.
func (w NSWorkspace) Init() NSWorkspace {
	rv := objc.Send[NSWorkspace](w.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (w NSWorkspace) Autorelease() NSWorkspace {
	rv := objc.Send[NSWorkspace](w.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSWorkspace creates a new NSWorkspace instance.
func NewNSWorkspace() NSWorkspace {
	class := getNSWorkspaceClass()
	rv := objc.Send[NSWorkspace](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Opens a URL asynchronously using the provided options.
//
// url: The URL to open.
//
// configuration: The options that indicate how you want to open the URL.
//
// completionHandler: The completion handler block to call asynchronously with the results.
// AppKit executes the completion handler on a concurrent queue. The handler
// block has no return value and takes the following parameters:
// 
// app: On success, this parameter contains a reference to the app that opened
// the URL. If the app didn’t open the URL successfully, this parameter is
// `nil`. error: On failure, this parameter contains an [NSError] object
// indicating the reason for the failure. If the method opened the URL
// successfully, this parameter is `nil`.
// //
// [NSError]: https://developer.apple.com/documentation/Foundation/NSError
//
// # Discussion
// 
// You may call this method safely from any thread of your app.
//
// See: https://developer.apple.com/documentation/AppKit/NSWorkspace/open(_:configuration:completionHandler:)
func (w NSWorkspace) OpenURLConfigurationCompletionHandler(url foundation.INSURL, configuration INSWorkspaceOpenConfiguration, completionHandler RunningApplicationErrorHandler) {
_block2, _cleanup2 := NewRunningApplicationErrorBlock(completionHandler)
	defer _cleanup2()
	objc.Send[objc.ID](w.ID, objc.Sel("openURL:configuration:completionHandler:"), url, configuration, _block2)
}

// Opens one or more URLs asynchronously in the specified app using the
// provided options.
//
// urls: The URL to open.
//
// applicationURL: A URL specifying the location of the app in the file system.
//
// configuration: The options that indicate how you want to open the URLs.
//
// completionHandler: The completion handler block to call asynchronously with the results.
// AppKit executes the completion handler on a concurrent queue. The handler
// block has no return value and takes the following parameters:
// 
// app: On success, this parameter contains a reference to the app that opened
// the URLs. If the app didn’t open the URLs successfully, this parameter is
// `nil`. error: On failure, this parameter contains an [NSError] object
// indicating the reason for the failure. If the method opened the URLs
// successfully, this parameter is `nil`.
// //
// [NSError]: https://developer.apple.com/documentation/Foundation/NSError
//
// # Discussion
// 
// You may call this method safely from any thread of your app.
//
// See: https://developer.apple.com/documentation/AppKit/NSWorkspace/open(_:withApplicationAt:configuration:completionHandler:)
func (w NSWorkspace) OpenURLsWithApplicationAtURLConfigurationCompletionHandler(urls []foundation.NSURL, applicationURL foundation.INSURL, configuration INSWorkspaceOpenConfiguration, completionHandler RunningApplicationErrorHandler) {
_block3, _cleanup3 := NewRunningApplicationErrorBlock(completionHandler)
	defer _cleanup3()
	objc.Send[objc.ID](w.ID, objc.Sel("openURLs:withApplicationAtURL:configuration:completionHandler:"), urls, applicationURL, configuration, _block3)
}

// Opens the location at the specified URL.
//
// url: A URL specifying the location to open.
//
// # Return Value
// 
// [true] if the location was successfully opened; otherwise, [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// You can call this method safely from any thread in macOS 10.6 and later.
//
// See: https://developer.apple.com/documentation/AppKit/NSWorkspace/open(_:)
func (w NSWorkspace) OpenURL(url foundation.INSURL) bool {
	rv := objc.Send[bool](w.ID, objc.Sel("openURL:"), url)
	return rv
}

// Launches the app at the specified URL and asynchronously reports back on
// the app’s status.
//
// applicationURL: A URL specifying the location of the app in the file system.
//
// configuration: The options that indicate how you want to launch the app.
//
// completionHandler: The completion handler block to call asynchronously with the results.
// AppKit executes the completion handler on a concurrent queue. The handler
// block has no return value and takes the following parameters:
// 
// app: On success, this parameter contains a reference to the launched app.
// If the app wasn’t launched, this parameter is `nil`. error: On failure,
// this parameter contains an [NSError] object indicating the reason for the
// failure. If the app launched successfully, this parameter is `nil`.
// //
// [NSError]: https://developer.apple.com/documentation/Foundation/NSError
//
// # Discussion
//
// See: https://developer.apple.com/documentation/AppKit/NSWorkspace/openApplication(at:configuration:completionHandler:)
func (w NSWorkspace) OpenApplicationAtURLConfigurationCompletionHandler(applicationURL foundation.INSURL, configuration INSWorkspaceOpenConfiguration, completionHandler RunningApplicationErrorHandler) {
_block2, _cleanup2 := NewRunningApplicationErrorBlock(completionHandler)
	defer _cleanup2()
	objc.Send[objc.ID](w.ID, objc.Sel("openApplicationAtURL:configuration:completionHandler:"), applicationURL, configuration, _block2)
}

// Hides all applications other than the sender.
//
// # Discussion
// 
// In order to hide all apps except the current one, the user can
// Command-Option-click an app’s Dock icon.
// 
// You must call this method from your app’s main thread.
//
// See: https://developer.apple.com/documentation/AppKit/NSWorkspace/hideOtherApplications()
func (w NSWorkspace) HideOtherApplications() {
	objc.Send[objc.ID](w.ID, objc.Sel("hideOtherApplications"))
}

// Activates the Finder, and opens one or more windows selecting the specified
// files.
//
// fileURLs: The files to select and display in the Finder.
//
// # Discussion
// 
// You can safely call this method from any thread of your app.
//
// See: https://developer.apple.com/documentation/AppKit/NSWorkspace/activateFileViewerSelecting(_:)
func (w NSWorkspace) ActivateFileViewerSelectingURLs(fileURLs []foundation.NSURL) {
	objc.Send[objc.ID](w.ID, objc.Sel("activateFileViewerSelectingURLs:"), objectivec.IObjectSliceToNSArray(fileURLs))
}

// Selects the file at the specified path.
//
// fullPath: The full path of the file to select.
//
// rootFullPath: The path to use for the file viewer. If you specify a nonempty path string,
// this method opens a new file viewer. If you specify an empty string
// (`@""`), this method selects the file in the main viewer.
//
// # Return Value
// 
// [true] if the file was successfully selected; otherwise, [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// In macOS 10.5 and later, this method does not follow symlinks when
// selecting the file. If the `fullPath` parameter contains any symlinks, this
// method selects the symlink instead of the file it targets. If you want to
// select the target file, use the [resolvingSymlinksInPath] method to resolve
// any symlinks before calling this method.
// 
// You can safely call this method from any thread of your app.
//
// [resolvingSymlinksInPath]: https://developer.apple.com/documentation/Foundation/NSString/resolvingSymlinksInPath
//
// See: https://developer.apple.com/documentation/AppKit/NSWorkspace/selectFile(_:inFileViewerRootedAtPath:)
func (w NSWorkspace) SelectFileInFileViewerRootedAtPath(fullPath string, rootFullPath string) bool {
	rv := objc.Send[bool](w.ID, objc.Sel("selectFile:inFileViewerRootedAtPath:"), objc.String(fullPath), objc.String(rootFullPath))
	return rv
}

// Returns the URL to the default app to open the specified URL.
//
// url: The URL of the file to open.
//
// # Return Value
// 
// The URL of the default app that would open the specified `url`. Returns
// `nil` if no app can open the URL, or if the file URL does not exist.
//
// # Discussion
// 
// This method is the programmatic equivalent of double-clicking a document in
// the Finder.
// 
// You can safely call this method from any thread of your app.
//
// See: https://developer.apple.com/documentation/AppKit/NSWorkspace/urlForApplication(toOpen:)-7qkzf
func (w NSWorkspace) URLForApplicationToOpenURL(url foundation.INSURL) foundation.INSURL {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("URLForApplicationToOpenURL:"), url)
	return foundation.NSURLFromID(rv)
}

// Returns the URL to the default app to open the specified content type.
//
// contentType: The content type to open.
//
// # Return Value
// 
// The URL of the default app that opens the specified `contentType`. Returns
// `nil` if no app can open the content type.
//
// See: https://developer.apple.com/documentation/AppKit/NSWorkspace/urlForApplication(toOpen:)-95cvp
func (w NSWorkspace) URLForApplicationToOpenContentType(contentType uniformtypeidentifiers.UTType) foundation.INSURL {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("URLForApplicationToOpenContentType:"), contentType)
	return foundation.NSURLFromID(rv)
}

// Returns the URL to the default app with the specified bundle identifier.
//
// bundleIdentifier: The bundle identifier for the app.
//
// # Return Value
// 
// The URL of the app, or `nil` if no app has the bundle identifier.
//
// # Discussion
// 
// This method uses various heuristics in case multiple apps have the same
// bundle ID.
// 
// You can safely call this method from any thread of your app.
//
// See: https://developer.apple.com/documentation/AppKit/NSWorkspace/urlForApplication(withBundleIdentifier:)
func (w NSWorkspace) URLForApplicationWithBundleIdentifier(bundleIdentifier string) foundation.INSURL {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("URLForApplicationWithBundleIdentifier:"), objc.String(bundleIdentifier))
	return foundation.NSURLFromID(rv)
}

// Returns an array of URLs to all available applications that can open the
// URL.
//
// url: The URL of the file to open.
//
// # Return Value
// 
// An array of URLs to available apps that can open the specified `url`.
// Returns an empty array if no app can open the URL, or if the file URL
// doesn’t exist.
//
// # Discussion
// 
// The system sorts the resulting array according to each app’s suitability
// to open the URL. The returned array lists the best match first.
//
// See: https://developer.apple.com/documentation/AppKit/NSWorkspace/urlsForApplications(toOpen:)-ualk
func (w NSWorkspace) URLsForApplicationsToOpenURL(url foundation.INSURL) []foundation.NSURL {
	rv := objc.Send[[]objc.ID](w.ID, objc.Sel("URLsForApplicationsToOpenURL:"), url)
	return objc.ConvertSlice(rv, func(id objc.ID) foundation.NSURL {
		return foundation.NSURLFromID(id)
	})
}

// Returns an array of URLs to all available applications that can open the
// specified content type.
//
// contentType: The content type to open.
//
// # Return Value
// 
// An array of URLs to available apps that can open the specified
// `contentType`. Returns an empty array if no app can open the content type.
//
// # Discussion
// 
// The system sorts the resulting array according to each app’s suitability
// to open the `contentType`. The returned array lists the best match first.
//
// See: https://developer.apple.com/documentation/AppKit/NSWorkspace/urlsForApplications(toOpen:)-60rkm
func (w NSWorkspace) URLsForApplicationsToOpenContentType(contentType uniformtypeidentifiers.UTType) []foundation.NSURL {
	rv := objc.Send[[]objc.ID](w.ID, objc.Sel("URLsForApplicationsToOpenContentType:"), contentType)
	return objc.ConvertSlice(rv, func(id objc.ID) foundation.NSURL {
		return foundation.NSURLFromID(id)
	})
}

// Returns an array of URLs to all available applications that can open the
// specified bundle identifier.
//
// bundleIdentifier: The bundle identifier for the apps to open.
//
// # Return Value
// 
// An array of URLs to available apps that can open the specified
// `bundleIdentifier`. Returns an empty array if no app associates with the
// bundle identifier.
//
// # Discussion
// 
// The system sorts the resulting array accounts according to each app’s
// suitability to launch. The returned array lists the best match first.
//
// See: https://developer.apple.com/documentation/AppKit/NSWorkspace/urlsForApplications(withBundleIdentifier:)
func (w NSWorkspace) URLsForApplicationsWithBundleIdentifier(bundleIdentifier string) []foundation.NSURL {
	rv := objc.Send[[]objc.ID](w.ID, objc.Sel("URLsForApplicationsWithBundleIdentifier:"), objc.String(bundleIdentifier))
	return objc.ConvertSlice(rv, func(id objc.ID) foundation.NSURL {
		return foundation.NSURLFromID(id)
	})
}

// Returns information about the file system at the specified path.
//
// fullPath: The path to the file system mount point.
//
// removableFlag: On input, a Boolean variable; on return, this variable contains [true] if
// the file system is on removable media.
// //
// [true]: https://developer.apple.com/documentation/Swift/true
//
// writableFlag: On input, a Boolean variable; on return, this variable contains [true] if
// the file system writable.
// //
// [true]: https://developer.apple.com/documentation/Swift/true
//
// unmountableFlag: On input, a Boolean variable; on return, this variable contains [true] if
// the file system is unmountable.
// //
// [true]: https://developer.apple.com/documentation/Swift/true
//
// description: On input, a pointer to a string object variable; on return, if the method
// was successful, this variable contains a string object that describes the
// file system. You should not rely on this description for program logic but
// can use it in message strings. Values can include “hard,” “nfs,”
// and “foreign.”
//
// fileSystemType: On input, a pointer to a string object variable; on return, if the method
// was successful, this variable contains the file system type. Values can
// include “HFS,” “UFS,” or other values.
//
// # Return Value
// 
// [true] if the information was returned; otherwise, [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// You can safely call this method from any thread of your app.
//
// See: https://developer.apple.com/documentation/AppKit/NSWorkspace/getFileSystemInfo(forPath:isRemovable:isWritable:isUnmountable:description:type:)
func (w NSWorkspace) GetFileSystemInfoForPathIsRemovableIsWritableIsUnmountableDescriptionType(fullPath string, description string, fileSystemType string) (bool, bool, bool, bool) {
	var removableFlag bool
	var writableFlag bool
	var unmountableFlag bool
	rv := objc.Send[bool](w.ID, objc.Sel("getFileSystemInfoForPath:isRemovable:isWritable:isUnmountable:description:type:"), objc.String(fullPath), unsafe.Pointer(&removableFlag), unsafe.Pointer(&writableFlag), unsafe.Pointer(&unmountableFlag), objc.String(description), objc.String(fileSystemType))
	return removableFlag, writableFlag, unmountableFlag, rv
}

// Determines whether the specified path is a file package.
//
// fullPath: The full path to examine.
//
// # Return Value
// 
// [true] if the path identifies a file package; otherwise, [false] if the
// path does not exist, is not a directory, or is not a file package.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// You can safely call this method from any thread of your app.
//
// See: https://developer.apple.com/documentation/AppKit/NSWorkspace/isFilePackage(atPath:)
func (w NSWorkspace) IsFilePackageAtPath(fullPath string) bool {
	rv := objc.Send[bool](w.ID, objc.Sel("isFilePackageAtPath:"), objc.String(fullPath))
	return rv
}

// Sets the default app to use when opening a specific file.
//
// applicationURL: The URL of the default app to use when opening the file.
//
// url: The URL of the file to open.
//
// completionHandler: The completion handler to call after the operation completes.
//
// # Discussion
// 
// This method sets the default app to use for a specific file (rather than
// all files of that content type). The system requires write access to the
// specified `url` before it can make the change.
// 
// If a change requires user consent, the system asks the user for consent
// asynchronously before invoking the completion handler.
// 
// This function doesn’t apply to non-file URLs.
//
// See: https://developer.apple.com/documentation/AppKit/NSWorkspace/setDefaultApplication(at:toOpenFileAt:completion:)
func (w NSWorkspace) SetDefaultApplicationAtURLToOpenFileAtURLCompletionHandler(applicationURL foundation.INSURL, url foundation.INSURL, completionHandler ErrorHandler) {
_block2, _ := NewErrorBlock(completionHandler)
	objc.Send[objc.ID](w.ID, objc.Sel("setDefaultApplicationAtURL:toOpenFileAtURL:completionHandler:"), applicationURL, url, _block2)
}

// Sets the default app to use when opening files of a specific content type.
//
// applicationURL: The URL of the default application.
//
// contentType: The content type to open.
//
// completionHandler: The completion handler to call after the operation completes.
//
// # Discussion
// 
// This method sets the default app to open for files of the specified
// `contentType`. If a change requires user consent, the system asks the user
// for consent asynchronously before invoking the completion handler.
//
// See: https://developer.apple.com/documentation/AppKit/NSWorkspace/setDefaultApplication(at:toOpen:completion:)
func (w NSWorkspace) SetDefaultApplicationAtURLToOpenContentTypeCompletionHandler(applicationURL foundation.INSURL, contentType uniformtypeidentifiers.UTType, completionHandler ErrorHandler) {
_block2, _ := NewErrorBlock(completionHandler)
	objc.Send[objc.ID](w.ID, objc.Sel("setDefaultApplicationAtURL:toOpenContentType:completionHandler:"), applicationURL, contentType, _block2)
}

// Sets the default app to use when opening files of a specific content type
// defined by a file URL.
//
// applicationURL: The URL of the default application.
//
// url: The URL of the file specifying the content type to open.
//
// completionHandler: The completion handler to call after the operation completes.
//
// # Discussion
// 
// This method sets the default app to open files of the type specified by the
// file `url`. If a change requires user consent, the system asks the user for
// consent asynchronously before invoking the completion handler.
//
// See: https://developer.apple.com/documentation/AppKit/NSWorkspace/setDefaultApplication(at:toOpenContentTypeOfFileAt:completion:)
func (w NSWorkspace) SetDefaultApplicationAtURLToOpenContentTypeOfFileAtURLCompletionHandler(applicationURL foundation.INSURL, url foundation.INSURL, completionHandler ErrorHandler) {
_block2, _ := NewErrorBlock(completionHandler)
	objc.Send[objc.ID](w.ID, objc.Sel("setDefaultApplicationAtURL:toOpenContentTypeOfFileAtURL:completionHandler:"), applicationURL, url, _block2)
}

// Sets the default app to use when opening files of a specific scheme.
//
// applicationURL: The URL of the default application.
//
// urlScheme: The URL of the scheme to open.
//
// completionHandler: The completion handler to call after the operation completes.
//
// # Discussion
// 
// This method sets the default app to open files of the type specified by the
// `urlScheme`. If a change requires user consent, the system asks the for
// consent asynchronously before invoking the completion handler.
//
// See: https://developer.apple.com/documentation/AppKit/NSWorkspace/setDefaultApplication(at:toOpenURLsWithScheme:completion:)
func (w NSWorkspace) SetDefaultApplicationAtURLToOpenURLsWithSchemeCompletionHandler(applicationURL foundation.INSURL, urlScheme string, completionHandler ErrorHandler) {
_block2, _ := NewErrorBlock(completionHandler)
	objc.Send[objc.ID](w.ID, objc.Sel("setDefaultApplicationAtURL:toOpenURLsWithScheme:completionHandler:"), applicationURL, objc.String(urlScheme), _block2)
}

// Returns an image containing the icon for the specified file.
//
// fullPath: The full path to the file.
//
// # Return Value
// 
// The icon associated with the file.
//
// # Discussion
// 
// The returned image has an initial size of 32 pixels by 32 pixels.
// 
// You can safely call this method from any thread of your app.
//
// See: https://developer.apple.com/documentation/AppKit/NSWorkspace/icon(forFile:)
func (w NSWorkspace) IconForFile(fullPath string) INSImage {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("iconForFile:"), objc.String(fullPath))
	return NSImageFromID(rv)
}

// Returns an image containing the icon for the specified files.
//
// fullPaths: An array of [NSString] objects, each of which contains the full path to a
// file.
//
// # Return Value
// 
// The icon associated with the group of files.
//
// # Discussion
// 
// If `fullPaths` specifies one file, that file’s icon is returned. If
// `fullPaths` specifies more than one file, an icon representing the multiple
// selection is returned.
// 
// You can safely call this method from any thread of your app.
//
// See: https://developer.apple.com/documentation/AppKit/NSWorkspace/icon(forFiles:)
func (w NSWorkspace) IconForFiles(fullPaths []string) INSImage {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("iconForFiles:"), objectivec.StringSliceToNSArray(fullPaths))
	return NSImageFromID(rv)
}

// Returns an image containing the icon for the specified content type.
//
// contentType: An object representing a uniform type of content.
//
// # Return Value
// 
// The icon associated with the content type.
//
// # Discussion
// 
// This method returns a default icon if the operation fails.
//
// See: https://developer.apple.com/documentation/AppKit/NSWorkspace/icon(for:)
func (w NSWorkspace) IconForContentType(contentType uniformtypeidentifiers.UTType) INSImage {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("iconForContentType:"), contentType)
	return NSImageFromID(rv)
}

// Sets the icon for the file or directory at the specified path.
//
// image: The image to use as the icon for the file or directory.
//
// fullPath: The full path of the file or directory.
//
// options: The icon representations to generate from the image. You specify this value
// by combining the appropriate [NSWorkspace.IconCreationOptions] constants,
// using the C bitwise [OR] operator. Specify `0` if you want to generate
// icons in all available icon representation formats.
// //
// [NSWorkspace.IconCreationOptions]: https://developer.apple.com/documentation/AppKit/NSWorkspace/IconCreationOptions
//
// # Return Value
// 
// [true] if the icon was set; otherwise, [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// The `image` can be an arbitrary image, with or without transparency. The
// method automatically scales this image (as needed) to generate the icon
// representations. The file or folder must exist and be writable by the user.
// 
// This method uses the image to set an icon with a size of 512 pixels by 512
// pixels. If you specify the [Exclude10_4ElementsIconCreationOption] option
// (not recommended), this method creates an icon that is compatible with the
// Finder from macOS 10.2 or earlier.
// 
// You can safely call this method from any of your app’s threads, but you
// must call it from only one thread at a time.
//
// See: https://developer.apple.com/documentation/AppKit/NSWorkspace/setIcon(_:forFile:options:)
func (w NSWorkspace) SetIconForFileOptions(image INSImage, fullPath string, options NSWorkspaceIconCreationOptions) bool {
	rv := objc.Send[bool](w.ID, objc.Sel("setIcon:forFile:options:"), image, objc.String(fullPath), options)
	return rv
}

// Unmounts and ejects the device at the specified path.
//
// path: The path to the device.
//
// # Return Value
// 
// [true] if the system unmounted the device; otherwise, [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// When this method begins, it posts an [willUnmountNotification] to the
// [NSWorkspace] object’s notification center. When it is finished, it posts
// an [didUnmountNotification].
// 
// Prefer the [UnmountAndEjectDeviceAtURLError] method because it provides
// more detailed error information.
// 
// You can safely call this method from any thread of your app.
//
// [didUnmountNotification]: https://developer.apple.com/documentation/AppKit/NSWorkspace/didUnmountNotification
// [willUnmountNotification]: https://developer.apple.com/documentation/AppKit/NSWorkspace/willUnmountNotification
//
// See: https://developer.apple.com/documentation/AppKit/NSWorkspace/unmountAndEjectDevice(atPath:)
func (w NSWorkspace) UnmountAndEjectDeviceAtPath(path string) bool {
	rv := objc.Send[bool](w.ID, objc.Sel("unmountAndEjectDeviceAtPath:"), objc.String(path))
	return rv
}

// Attempts to eject the volume mounted at the given path.
//
// url: The URL of the volume to eject.
//
// # Discussion
// 
// You can safely call this method from any thread of your app.
//
// See: https://developer.apple.com/documentation/AppKit/NSWorkspace/unmountAndEjectDevice(at:)
func (w NSWorkspace) UnmountAndEjectDeviceAtURLError(url foundation.INSURL) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](w.ID, objc.Sel("unmountAndEjectDeviceAtURL:error:"), url, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("unmountAndEjectDeviceAtURL:error: returned NO with nil NSError")
	}
	return rv, nil

}

// Returns the URL for the desktop image for the given screen.
//
// screen: The screen for which to get the desktop image.
//
// # Return Value
// 
// The desktop image.
//
// # Discussion
// 
// You must call this method from your app’s main thread.
//
// See: https://developer.apple.com/documentation/AppKit/NSWorkspace/desktopImageURL(for:)
func (w NSWorkspace) DesktopImageURLForScreen(screen INSScreen) foundation.INSURL {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("desktopImageURLForScreen:"), screen)
	return foundation.NSURLFromID(rv)
}

// Sets the desktop image for the given screen to the image at the specified
// URL.
//
// url: A file URL to the image. The URL must not be `nil`.
//
// screen: The screen on which to set the desktop image.
//
// options: The options dictionary may contain any of the keys in
// [NSWorkspaceDesktopImageOptionKey], which control how the image is scaled
// on the screen.
//
// # Discussion
// 
// Instead of presenting a user interface for picking the options, choose
// appropriate defaults and allow the user to adjust them in the System
// Preference Pane.
// 
// You must call this method from your app’s main thread.
//
// See: https://developer.apple.com/documentation/AppKit/NSWorkspace/setDesktopImageURL(_:for:options:)
func (w NSWorkspace) SetDesktopImageURLForScreenOptionsError(url foundation.INSURL, screen INSScreen, options foundation.INSDictionary) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](w.ID, objc.Sel("setDesktopImageURL:forScreen:options:error:"), url, screen, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("setDesktopImageURL:forScreen:options:error: returned NO with nil NSError")
	}
	return rv, nil

}

// Returns the desktop image options for the given screen.
//
// screen: The screen for which to get the desktop image options.
//
// # Return Value
// 
// A dictionary containing the keys found in
// [NSWorkspaceDesktopImageOptionKey].
//
// # Discussion
// 
// You must call this method from your app’s main thread.
//
// See: https://developer.apple.com/documentation/AppKit/NSWorkspace/desktopImageOptions(for:)
func (w NSWorkspace) DesktopImageOptionsForScreen(screen INSScreen) foundation.INSDictionary {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("desktopImageOptionsForScreen:"), screen)
	return foundation.NSDictionaryFromID(rv)
}

// Displays a Spotlight search results window in Finder for the specified
// query string.
//
// queryString: The string to search for.
//
// # Return Value
// 
// [true] if the method communicated successfully with Finder; otherwise,
// [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// Finder becomes the active app, if possible. The user can further refine the
// search via the Finder user interface.
// 
// You can safely call this method from any thread of your app.
//
// See: https://developer.apple.com/documentation/AppKit/NSWorkspace/showSearchResults(forQueryString:)
func (w NSWorkspace) ShowSearchResultsForQueryString(queryString string) bool {
	rv := objc.Send[bool](w.ID, objc.Sel("showSearchResultsForQueryString:"), objc.String(queryString))
	return rv
}

// Informs the workspace object that the file system changed at the specified
// path.
//
// path: The full path that changed.
//
// # Discussion
// 
// Avoid calling this method if possible. If you want to track changes to
// files and directories, use the FSEvents API described in [FSEvents].
// 
// The [NSWorkspace] object uses this method to track changes to all the files
// and directories in which it is interested.
//
// See: https://developer.apple.com/documentation/AppKit/NSWorkspace/noteFileSystemChanged(_:)
func (w NSWorkspace) NoteFileSystemChanged(path string) {
	objc.Send[objc.ID](w.ID, objc.Sel("noteFileSystemChanged:"), objc.String(path))
}

// Requests the system wait for the specified amount of time before turning
// off the power or logging out the user.
//
// requested: The number of milliseconds to wait before turning off the power or logging
// off the user.
//
// # Return Value
// 
// The number of milliseconds granted by the system. This method currently
// returns `0`.
//
// # Discussion
// 
// This method currently does nothing. Do not call it.
//
// See: https://developer.apple.com/documentation/AppKit/NSWorkspace/extendPowerOff(by:)
func (w NSWorkspace) ExtendPowerOffBy(requested int) int {
	rv := objc.Send[int](w.ID, objc.Sel("extendPowerOffBy:"), requested)
	return rv
}

// Requests authorization to perform a privileged file operation.
//
// type: The type of file operation to perform.
//
// completionHandler: The completion handler to call when the authorization request is completed.
// 
// The completion handler takes two parameters:
// 
// authorization: The authorization granted for this app. Use it when creating
// a new [FileManager] with [init(authorization:)]. error: `nil` if the app is
// authorized; otherwise, a pointer to the authorization error.
// //
// [FileManager]: https://developer.apple.com/documentation/Foundation/FileManager
// [init(authorization:)]: https://developer.apple.com/documentation/Foundation/FileManager/init(authorization:)
//
// # Discussion
//
// See: https://developer.apple.com/documentation/AppKit/NSWorkspace/requestAuthorization(to:completionHandler:)
func (w NSWorkspace) RequestAuthorizationOfTypeCompletionHandler(type_ NSWorkspaceAuthorizationType, completionHandler WorkspaceAuthorizationErrorHandler) {
_block1, _cleanup1 := NewWorkspaceAuthorizationErrorBlock(completionHandler)
	defer _cleanup1()
	objc.Send[objc.ID](w.ID, objc.Sel("requestAuthorizationOfType:completionHandler:"), type_, _block1)
}

// The notification center for workspace notifications.
//
// # Return Value
// 
// The notification center object associated with the workspace
// 
// # Discussion
// 
// This notification center object delivers the workspace-related
// notifications described in Responding to Environment Notifications.
// 
// You can access this object safely from any thread in your app in macOS 10.6
// and later.
//
// See: https://developer.apple.com/documentation/AppKit/NSWorkspace/notificationCenter
func (w NSWorkspace) NotificationCenter() foundation.NSNotificationCenter {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("notificationCenter"))
	return foundation.NSNotificationCenterFromID(objc.ID(rv))
}

// Returns the frontmost app, which is the app that receives key events.
//
// # Return Value
// 
// The running app instance for the app that receives key events.
// 
// # Discussion
// 
// This value is key-value observing compliant.
//
// See: https://developer.apple.com/documentation/AppKit/NSWorkspace/frontmostApplication
func (w NSWorkspace) FrontmostApplication() INSRunningApplication {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("frontmostApplication"))
	return NSRunningApplicationFromID(objc.ID(rv))
}

// Returns an array of running apps.
//
// # Return Value
// 
// An array of [NSRunningApplication] instances. This value is key-value
// observing compliant.
// 
// # Discussion
// 
// The order of the array is unspecified, but it is stable, meaning that the
// relative order of particular apps will not change across multiple calls to
// `runningApplications`. See [NSRunningApplication] for more information on
// [NSRunningApplication].
// 
// Similar to the [NSRunningApplication] class’s properties, this property
// will only change when the main run loop runs in a common mode. Instead of
// polling, use key-value observing to be notified of changes to this array
// property.
// 
// You can safely call this method from any of your app’s threads. The
// method returns its value atomically.
//
// See: https://developer.apple.com/documentation/AppKit/NSWorkspace/runningApplications
func (w NSWorkspace) RunningApplications() []NSRunningApplication {
	rv := objc.Send[[]objc.ID](w.ID, objc.Sel("runningApplications"))
	return objc.ConvertSlice(rv, func(id objc.ID) NSRunningApplication {
		return NSRunningApplicationFromID(id)
	})
}

// Returns the app that owns the currently displayed menu bar.
//
// # Discussion
// 
// This property contains the running app instance for the app that owns the
// displayed menu bar. This value is key-value observing compliant.
//
// See: https://developer.apple.com/documentation/AppKit/NSWorkspace/menuBarOwningApplication
func (w NSWorkspace) MenuBarOwningApplication() INSRunningApplication {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("menuBarOwningApplication"))
	return NSRunningApplicationFromID(objc.ID(rv))
}

// The array of file labels, returned as strings.
//
// # Return Value
// 
// An array of strings.
// 
// # Discussion
// 
// You can listen for notifications named [didChangeFileLabelsNotification] to
// be notified when file labels change.
// 
// You can safely call this method from any thread of your app.
//
// [didChangeFileLabelsNotification]: https://developer.apple.com/documentation/AppKit/NSWorkspace/didChangeFileLabelsNotification
//
// See: https://developer.apple.com/documentation/AppKit/NSWorkspace/fileLabels
func (w NSWorkspace) FileLabels() []string {
	rv := objc.Send[[]objc.ID](w.ID, objc.Sel("fileLabels"))
	return objc.ConvertSliceToStrings(rv)
}

// The array of colors for the file labels.
//
// # Return Value
// 
// An array of [NSColor] objects.
// 
// # Discussion
// 
// This array has the same number of elements as [FileLabels], and the color
// at a given index corresponds to the label at the same index.
// 
// You can listen for notifications named [didChangeFileLabelsNotification] to
// be notified when file labels change that may result in changes to the order
// of the `fileLabelColors`.
// 
// You can safely call this method from any thread of your app.
//
// [didChangeFileLabelsNotification]: https://developer.apple.com/documentation/AppKit/NSWorkspace/didChangeFileLabelsNotification
//
// See: https://developer.apple.com/documentation/AppKit/NSWorkspace/fileLabelColors
func (w NSWorkspace) FileLabelColors() []NSColor {
	rv := objc.Send[[]objc.ID](w.ID, objc.Sel("fileLabelColors"))
	return objc.ConvertSlice(rv, func(id objc.ID) NSColor {
		return NSColorFromID(id)
	})
}

// A Boolean value that indicates whether the app avoids conveying information
// through color alone.
//
// # Discussion
// 
// If this property is [true], the user interface avoids conveying information
// using color alone. Instead, use shapes or glyphs to convey important
// information.
// 
// Users can change this setting by choosing System Preferences >
// Accessibility > Display and selecting the “Differentiate without color”
// option. To receive updates when this setting changes, register for the
// [accessibilityDisplayOptionsDidChangeNotification] notification using
// [NotificationCenter].
//
// [accessibilityDisplayOptionsDidChangeNotification]: https://developer.apple.com/documentation/AppKit/NSWorkspace/accessibilityDisplayOptionsDidChangeNotification
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSWorkspace/accessibilityDisplayShouldDifferentiateWithoutColor
func (w NSWorkspace) AccessibilityDisplayShouldDifferentiateWithoutColor() bool {
	rv := objc.Send[bool](w.ID, objc.Sel("accessibilityDisplayShouldDifferentiateWithoutColor"))
	return rv
}

// A Boolean value that indicates whether the app presents a high-contrast
// user interface.
//
// # Discussion
// 
// When this method returns [true], present a high-contrast UI. For example,
// use a less subtle color palette or bolder lines.
// 
// Users can change this setting by choosing System Preferences >
// Accessibility > Display and selecting the “Increase contrast” option.
// To receive updates when this setting changes, register for the
// [accessibilityDisplayOptionsDidChangeNotification] notification using
// [NotificationCenter].
//
// [accessibilityDisplayOptionsDidChangeNotification]: https://developer.apple.com/documentation/AppKit/NSWorkspace/accessibilityDisplayOptionsDidChangeNotification
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSWorkspace/accessibilityDisplayShouldIncreaseContrast
func (w NSWorkspace) AccessibilityDisplayShouldIncreaseContrast() bool {
	rv := objc.Send[bool](w.ID, objc.Sel("accessibilityDisplayShouldIncreaseContrast"))
	return rv
}

// A Boolean value that indicates whether the app avoids using semitransparent
// backgrounds.
//
// # Discussion
// 
// If this property is [true], don’t use semitransparent backgrounds in the
// user interface. For example, use only opaque windows.
// 
// Users can change this setting by choosing System Preferences >
// Accessibility > Display and selecting the “Reduce transparency” option.
// To receive updates when this setting changes, register to the
// [accessibilityDisplayOptionsDidChangeNotification] notification using
// [NotificationCenter].
//
// [accessibilityDisplayOptionsDidChangeNotification]: https://developer.apple.com/documentation/AppKit/NSWorkspace/accessibilityDisplayOptionsDidChangeNotification
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSWorkspace/accessibilityDisplayShouldReduceTransparency
func (w NSWorkspace) AccessibilityDisplayShouldReduceTransparency() bool {
	rv := objc.Send[bool](w.ID, objc.Sel("accessibilityDisplayShouldReduceTransparency"))
	return rv
}

// A Boolean value that indicates whether the accessibility option to invert
// colors is in an enabled state.
//
// # Discussion
// 
// If this property’s value is [true], the system inverts the display. In
// this case, you may need to adjust your app’s drawing for optimal display.
// To receive updates when this setting changes, register for the
// [accessibilityDisplayOptionsDidChangeNotification] notification using
// [NotificationCenter].
//
// [accessibilityDisplayOptionsDidChangeNotification]: https://developer.apple.com/documentation/AppKit/NSWorkspace/accessibilityDisplayOptionsDidChangeNotification
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSWorkspace/accessibilityDisplayShouldInvertColors
func (w NSWorkspace) AccessibilityDisplayShouldInvertColors() bool {
	rv := objc.Send[bool](w.ID, objc.Sel("accessibilityDisplayShouldInvertColors"))
	return rv
}

// A Boolean value that indicates whether the accessibility option to reduce
// motion is in an enabled state.
//
// # Discussion
// 
// If this property’s value is [true], avoid large animations, especially
// those that simulate the third dimension. To receive updates when this
// setting changes, register for the
// [accessibilityDisplayOptionsDidChangeNotification] notification using
// [NotificationCenter].
//
// [accessibilityDisplayOptionsDidChangeNotification]: https://developer.apple.com/documentation/AppKit/NSWorkspace/accessibilityDisplayOptionsDidChangeNotification
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSWorkspace/accessibilityDisplayShouldReduceMotion
func (w NSWorkspace) AccessibilityDisplayShouldReduceMotion() bool {
	rv := objc.Send[bool](w.ID, objc.Sel("accessibilityDisplayShouldReduceMotion"))
	return rv
}

// A Boolean value that indicates whether Switch Control is currently running.
//
// # Discussion
// 
// You can observe this property with key-value observing.
//
// See: https://developer.apple.com/documentation/AppKit/NSWorkspace/isSwitchControlEnabled
func (w NSWorkspace) SwitchControlEnabled() bool {
	rv := objc.Send[bool](w.ID, objc.Sel("isSwitchControlEnabled"))
	return rv
}

// A Boolean value that indicates whether VoiceOver is currently running.
//
// # Discussion
// 
// You can observe this property with key-value observing.
//
// See: https://developer.apple.com/documentation/AppKit/NSWorkspace/isVoiceOverEnabled
func (w NSWorkspace) VoiceOverEnabled() bool {
	rv := objc.Send[bool](w.ID, objc.Sel("isVoiceOverEnabled"))
	return rv
}

// The shared workspace object.
//
// # Return Value
// 
// The [NSWorkspace] object associated with the process.
// 
// # Discussion
// 
// You can access this object safely from any thread in your app.
//
// See: https://developer.apple.com/documentation/AppKit/NSWorkspace/shared
func (_NSWorkspaceClass NSWorkspaceClass) SharedWorkspace() NSWorkspace {
	rv := objc.Send[objc.ID](objc.ID(_NSWorkspaceClass.class), objc.Sel("sharedWorkspace"))
	return NSWorkspaceFromID(objc.ID(rv))
}

// OpenURLConfiguration is a synchronous wrapper around [NSWorkspace.OpenURLConfigurationCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (w NSWorkspace) OpenURLConfiguration(ctx context.Context, url foundation.INSURL, configuration INSWorkspaceOpenConfiguration) (*NSRunningApplication, error) {
	type result struct {
		val *NSRunningApplication
		err error
	}
	done := make(chan result, 1)
	w.OpenURLConfigurationCompletionHandler(url, configuration, func(val *NSRunningApplication, err error) {
		done <- result{val, err}
	})
	select {
	case r := <-done:
		return r.val, r.err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

// OpenApplicationAtURLConfiguration is a synchronous wrapper around [NSWorkspace.OpenApplicationAtURLConfigurationCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (w NSWorkspace) OpenApplicationAtURLConfiguration(ctx context.Context, applicationURL foundation.INSURL, configuration INSWorkspaceOpenConfiguration) (*NSRunningApplication, error) {
	type result struct {
		val *NSRunningApplication
		err error
	}
	done := make(chan result, 1)
	w.OpenApplicationAtURLConfigurationCompletionHandler(applicationURL, configuration, func(val *NSRunningApplication, err error) {
		done <- result{val, err}
	})
	select {
	case r := <-done:
		return r.val, r.err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

// SetDefaultApplicationAtURLToOpenFileAtURL is a synchronous wrapper around [NSWorkspace.SetDefaultApplicationAtURLToOpenFileAtURLCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (w NSWorkspace) SetDefaultApplicationAtURLToOpenFileAtURL(ctx context.Context, applicationURL foundation.INSURL, url foundation.INSURL) error {
	done := make(chan error, 1)
	w.SetDefaultApplicationAtURLToOpenFileAtURLCompletionHandler(applicationURL, url, func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

// SetDefaultApplicationAtURLToOpenContentType is a synchronous wrapper around [NSWorkspace.SetDefaultApplicationAtURLToOpenContentTypeCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (w NSWorkspace) SetDefaultApplicationAtURLToOpenContentType(ctx context.Context, applicationURL foundation.INSURL, contentType uniformtypeidentifiers.UTType) error {
	done := make(chan error, 1)
	w.SetDefaultApplicationAtURLToOpenContentTypeCompletionHandler(applicationURL, contentType, func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

// SetDefaultApplicationAtURLToOpenContentTypeOfFileAtURL is a synchronous wrapper around [NSWorkspace.SetDefaultApplicationAtURLToOpenContentTypeOfFileAtURLCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (w NSWorkspace) SetDefaultApplicationAtURLToOpenContentTypeOfFileAtURL(ctx context.Context, applicationURL foundation.INSURL, url foundation.INSURL) error {
	done := make(chan error, 1)
	w.SetDefaultApplicationAtURLToOpenContentTypeOfFileAtURLCompletionHandler(applicationURL, url, func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

// SetDefaultApplicationAtURLToOpenURLsWithScheme is a synchronous wrapper around [NSWorkspace.SetDefaultApplicationAtURLToOpenURLsWithSchemeCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (w NSWorkspace) SetDefaultApplicationAtURLToOpenURLsWithScheme(ctx context.Context, applicationURL foundation.INSURL, urlScheme string) error {
	done := make(chan error, 1)
	w.SetDefaultApplicationAtURLToOpenURLsWithSchemeCompletionHandler(applicationURL, urlScheme, func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

// RequestAuthorizationOfType is a synchronous wrapper around [NSWorkspace.RequestAuthorizationOfTypeCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (w NSWorkspace) RequestAuthorizationOfType(ctx context.Context, type_ NSWorkspaceAuthorizationType) (*NSWorkspaceAuthorization, error) {
	type result struct {
		val *NSWorkspaceAuthorization
		err error
	}
	done := make(chan result, 1)
	w.RequestAuthorizationOfTypeCompletionHandler(type_, func(val *NSWorkspaceAuthorization, err error) {
		done <- result{val, err}
	})
	select {
	case r := <-done:
		return r.val, r.err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

