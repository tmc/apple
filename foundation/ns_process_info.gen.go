// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"context"
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [ProcessInfo] class.
var (
	_ProcessInfoClass     ProcessInfoClass
	_ProcessInfoClassOnce sync.Once
)

func getProcessInfoClass() ProcessInfoClass {
	_ProcessInfoClassOnce.Do(func() {
		_ProcessInfoClass = ProcessInfoClass{class: objc.GetClass("NSProcessInfo")}
	})
	return _ProcessInfoClass
}

// GetProcessInfoClass returns the class object for NSProcessInfo.
func GetProcessInfoClass() ProcessInfoClass {
	return getProcessInfoClass()
}

type ProcessInfoClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (pc ProcessInfoClass) Class() objc.Class {
	return pc.class
}

// Alloc allocates memory for a new instance of the class.
func (pc ProcessInfoClass) Alloc() ProcessInfo {
	rv := objc.Send[ProcessInfo](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// A collection of information about the current process.
//
// # Overview
//
// Each process has a single, shared [NSProcessInfo] object known as a that
// can return information such as arguments, environment variables, host name,
// and process name. The [ProcessInfo] class method returns the shared agent
// for the current process. For example, the following line returns the
// [NSProcessInfo] object, which then provides the name of the current
// process:
//
// The [NSProcessInfo] class also includes the [OperatingSystemVersion]
// property, which returns an [OperatingSystemVersion] structure identifying
// the operating system version on which the process is executing.
//
// [NSProcessInfo] objects attempt to interpret environment variables and
// command-line arguments in the user’s default C string encoding if they
// can’t convert to Unicode as UTF-8 strings. If neither the Unicode nor C
// string conversion works, the [NSProcessInfo] object ignores these values.
//
// # Manage Activities
//
// The system has heuristics to improve battery life, performance, and
// responsiveness of applications for the benefit of the user. You can use the
// following methods to manage that give hints to the system that your
// application has special requirements:
//
// - [BeginActivityWithOptionsReason] - [EndActivity] -
// [PerformActivityWithOptionsReasonUsingBlock]
//
// In response to creating an activity, the system disables some or all of the
// heuristics so your application can finish quickly while still providing
// responsive behavior if the user needs it.
//
// You use activities when your application performs a long-running operation.
// If the activity can take different amounts of time (for example,
// calculating the next move in a chess game), it should use this API to
// ensure correct behavior when the amount of data or the capabilities of the
// user’s computer varies. Activities fall into two major categories:
//
// - activities are explicitly started by the user. Examples include exporting
// or downloading a user-specified file. - activities perform the normal
// operations of your application and aren’t explicitly started by the user.
// Examples include autosaving, indexing, and automatic downloading of files.
//
// In addition, if your application requires high priority input/output (I/O),
// you can include the [NSActivityLatencyCritical] flag (using a bitwise
// [OR]). You should only use this flag for activities like audio or video
// recording that require high priority I/O.
//
// If your activity takes place synchronously inside an event callback on the
// main thread, you don’t need to use this API.
//
// Be aware that failing to end these activities for an extended period of
// time can have significant negative impacts on the performance of your
// user’s computer, so be sure to use only the minimum amount of time
// required. User preferences may override your application’s request.
//
// You can also use this API to control automatic termination or sudden
// termination (see [NSProcessInfo]). For example, the following code brackets
// the work to protect it from sudden termination:
//
// The above example is equivalent to the following code, which uses the
// [DisableAutomaticTermination] method:
//
// Because this API returns an object, it may be easier to pair begins and
// ends than when using the automatic termination API. If your app deallocates
// the object before the [EndActivity] call, the activity ends automatically.
//
// This API also provides a mechanism to disable system-wide idle sleep and
// display idle sleep. These can have a large impact on the user experience,
// so be careful to end activities that disable sleep (including
// [NSActivityUserInitiated]).
//
// # Support Sudden Termination
//
// macOS 10.6 and later includes a mechanism that allows the system to log out
// or shut down more quickly by, whenever possible, killing applications
// instead of requesting that they quit themselves.
//
// Your application can enable this capability on a global basis and then
// manually override its availability during actions that could cause data
// corruption or a poor user experience by allowing sudden termination.
//
// Alternatively, your application can manually enable and disable this
// functionality. Creating a process assigns a counter that indicates if the
// process is safe to terminate. You decrement and increment the counter using
// the methods [EnableSuddenTermination] and [DisableSuddenTermination]. A
// value of `0` enables the system to terminate the process without first
// sending a notification or event.
//
// Your application can support sudden termination upon launch by adding a key
// to the application’s `Info.Plist()` file. If the
// [NSSupportsSuddenTermination] key exists in the `Info.Plist()` file and has
// a value of true, it’s the equivalent of calling [EnableSuddenTermination]
// during your application launch. This allows the system to terminate the
// process immediately. You can still override this behavior by invoking
// [DisableSuddenTermination].
//
// Typically, you disable sudden termination whenever your app defers work
// that the app must complete before it terminates. If, for example, your app
// defers writing data to disk and enables sudden termination, you should
// bracket the sensitive operations with a call to [DisableSuddenTermination],
// perform the necessary operations, and then send a balancing
// [EnableSuddenTermination] message.
//
// In agents or daemon executables that don’t depend on AppKit, you can
// manually invoke [EnableSuddenTermination] right away. You can then use the
// enable and disable methods whenever the process has work it must do before
// it terminates.
//
// Some AppKit functionality automatically disables sudden termination on a
// temporary basis to ensure data integrity.
//
// - [NSUserDefaults] temporarily disables sudden termination to prevent the
// process from terminating between the time at which it sets the default and
// the time at which it writes the preferences file — including that default
// — to disk. - [NSDocument] temporarily disables sudden termination to
// prevent the process from terminating between the time at which the user has
// made a change to a document and the time at which [NSDocument] writes the
// user’s change to disk.
//
// # Monitor Thermal State to Adjust App Performance
//
// indicates the level of heat generated by logic components as they run apps.
// As the thermal state increases, the system decreases heat by reducing the
// speed of the processors. Optimize your app’s performance by monitoring
// the thermal state and reducing system usage as the thermal state increases.
// Query the current state with [ThermalState] to determine if your app needs
// to reduce system usage. You can register the
// [thermalStateDidChangeNotification] for notifications of a change in
// thermal state. For recommended actions, see [ProcessInfo.ThermalState].
//
// # Accessing Process Information
//
//   - [ProcessInfo.Arguments]: Array of strings with the command-line arguments for the process.
//   - [ProcessInfo.Environment]: The variable names (keys) and their values in the environment from which the process was launched.
//   - [ProcessInfo.GloballyUniqueString]: Global unique identifier for the process.
//   - [ProcessInfo.MacCatalystApp]: A Boolean value that indicates whether the process originated as an iOS app and runs on macOS.
//   - [ProcessInfo.IOSAppOnMac]: A Boolean value that indicates whether the process is an iPhone or iPad app running on a Mac.
//   - [ProcessInfo.IOSAppOnVision]: A Boolean value that indicates whether the process is an iPhone or iPad app running on visionOS.
//   - [ProcessInfo.ProcessIdentifier]: The identifier of the process (often called process ID).
//   - [ProcessInfo.ProcessName]: The name of the process.
//   - [ProcessInfo.SetProcessName]
//
// # Accessing User Information
//
//   - [ProcessInfo.UserName]: Returns the account name of the current user.
//   - [ProcessInfo.FullUserName]: Returns the full name of the current user.
//
// # Sudden Application Termination
//
//   - [ProcessInfo.DisableSuddenTermination]: Disables the application for quickly killing using sudden termination.
//   - [ProcessInfo.EnableSuddenTermination]: Enables the application for quick killing using sudden termination.
//
// # Controlling Automatic Termination
//
//   - [ProcessInfo.DisableAutomaticTermination]: Disables automatic termination for the application.
//   - [ProcessInfo.EnableAutomaticTermination]: Enables automatic termination for the application.
//   - [ProcessInfo.AutomaticTerminationSupportEnabled]: A Boolean value indicating whether the app supports automatic termination.
//   - [ProcessInfo.SetAutomaticTerminationSupportEnabled]
//
// # Getting Host Information
//
//   - [ProcessInfo.HostName]: The name of the host computer on which the process is executing.
//   - [ProcessInfo.OperatingSystemVersionString]: A string containing the version of the operating system on which the process is executing.
//   - [ProcessInfo.OperatingSystemVersion]: The version of the operating system on which the process is executing.
//   - [ProcessInfo.IsOperatingSystemAtLeastVersion]: Returns a Boolean value indicating whether the version of the operating system on which the process is executing is the same or later than the given version.
//
// # Getting Computer Information
//
//   - [ProcessInfo.ProcessorCount]: The number of processing cores available on the computer.
//   - [ProcessInfo.ActiveProcessorCount]: The number of active processing cores available on the computer.
//   - [ProcessInfo.PhysicalMemory]: The amount of physical memory on the computer in bytes.
//   - [ProcessInfo.IsDeviceCertifiedFor]: Indicates whether the device supports the requested performance tier.
//   - [ProcessInfo.HasPerformanceProfile]: Indicates whether an app is running under a known performance profile.
//   - [ProcessInfo.SystemUptime]: The amount of time the system has been awake since the last time it was restarted.
//
// # Managing Activities
//
//   - [ProcessInfo.BeginActivityWithOptionsReason]: Begin an activity using the given options and reason.
//   - [ProcessInfo.EndActivity]: Ends the given activity.
//   - [ProcessInfo.PerformActivityWithOptionsReasonUsingBlock]: Synchronously perform an activity defined by a given block using the given options.
//
// # Getting the Thermal State
//
//   - [ProcessInfo.ThermalState]: The current thermal state of the system.
//
// # Determining Whether Low Power Mode is Enabled
//
//   - [ProcessInfo.LowPowerModeEnabled]: A Boolean value that indicates the current state of Low Power Mode.
//
// # Notifications
//
//   - [ProcessInfo.NSProcessInfoPowerStateDidChange]: Posts when the power state of a device changes.
//
// See: https://developer.apple.com/documentation/Foundation/ProcessInfo
//
// [NSDocument]: https://developer.apple.com/documentation/AppKit/NSDocument
// [NSSupportsSuddenTermination]: https://developer.apple.com/documentation/BundleResources/Information-Property-List/NSSupportsSuddenTermination
// [OperatingSystemVersion]: https://developer.apple.com/documentation/Foundation/OperatingSystemVersion
// [ProcessInfo.ThermalState]: https://developer.apple.com/documentation/Foundation/ProcessInfo/ThermalState-swift.enum
// [thermalStateDidChangeNotification]: https://developer.apple.com/documentation/Foundation/ProcessInfo/thermalStateDidChangeNotification
type ProcessInfo struct {
	objectivec.Object
}

// ProcessInfoFromID constructs a [ProcessInfo] from an objc.ID.
//
// A collection of information about the current process.
func ProcessInfoFromID(id objc.ID) ProcessInfo {
	return NSProcessInfo{objectivec.Object{ID: id}}
}

// NSProcessInfoFromID is an alias for [ProcessInfoFromID] for cross-framework compatibility.
func NSProcessInfoFromID(id objc.ID) ProcessInfo { return ProcessInfoFromID(id) }

// NOTE: ProcessInfo adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [ProcessInfo] class.
//
// # Accessing Process Information
//
//   - [IProcessInfo.Arguments]: Array of strings with the command-line arguments for the process.
//   - [IProcessInfo.Environment]: The variable names (keys) and their values in the environment from which the process was launched.
//   - [IProcessInfo.GloballyUniqueString]: Global unique identifier for the process.
//   - [IProcessInfo.MacCatalystApp]: A Boolean value that indicates whether the process originated as an iOS app and runs on macOS.
//   - [IProcessInfo.IOSAppOnMac]: A Boolean value that indicates whether the process is an iPhone or iPad app running on a Mac.
//   - [IProcessInfo.IOSAppOnVision]: A Boolean value that indicates whether the process is an iPhone or iPad app running on visionOS.
//   - [IProcessInfo.ProcessIdentifier]: The identifier of the process (often called process ID).
//   - [IProcessInfo.ProcessName]: The name of the process.
//   - [IProcessInfo.SetProcessName]
//
// # Accessing User Information
//
//   - [IProcessInfo.UserName]: Returns the account name of the current user.
//   - [IProcessInfo.FullUserName]: Returns the full name of the current user.
//
// # Sudden Application Termination
//
//   - [IProcessInfo.DisableSuddenTermination]: Disables the application for quickly killing using sudden termination.
//   - [IProcessInfo.EnableSuddenTermination]: Enables the application for quick killing using sudden termination.
//
// # Controlling Automatic Termination
//
//   - [IProcessInfo.DisableAutomaticTermination]: Disables automatic termination for the application.
//   - [IProcessInfo.EnableAutomaticTermination]: Enables automatic termination for the application.
//   - [IProcessInfo.AutomaticTerminationSupportEnabled]: A Boolean value indicating whether the app supports automatic termination.
//   - [IProcessInfo.SetAutomaticTerminationSupportEnabled]
//
// # Getting Host Information
//
//   - [IProcessInfo.HostName]: The name of the host computer on which the process is executing.
//   - [IProcessInfo.OperatingSystemVersionString]: A string containing the version of the operating system on which the process is executing.
//   - [IProcessInfo.OperatingSystemVersion]: The version of the operating system on which the process is executing.
//   - [IProcessInfo.IsOperatingSystemAtLeastVersion]: Returns a Boolean value indicating whether the version of the operating system on which the process is executing is the same or later than the given version.
//
// # Getting Computer Information
//
//   - [IProcessInfo.ProcessorCount]: The number of processing cores available on the computer.
//   - [IProcessInfo.ActiveProcessorCount]: The number of active processing cores available on the computer.
//   - [IProcessInfo.PhysicalMemory]: The amount of physical memory on the computer in bytes.
//   - [IProcessInfo.IsDeviceCertifiedFor]: Indicates whether the device supports the requested performance tier.
//   - [IProcessInfo.HasPerformanceProfile]: Indicates whether an app is running under a known performance profile.
//   - [IProcessInfo.SystemUptime]: The amount of time the system has been awake since the last time it was restarted.
//
// # Managing Activities
//
//   - [IProcessInfo.BeginActivityWithOptionsReason]: Begin an activity using the given options and reason.
//   - [IProcessInfo.EndActivity]: Ends the given activity.
//   - [IProcessInfo.PerformActivityWithOptionsReasonUsingBlock]: Synchronously perform an activity defined by a given block using the given options.
//
// # Getting the Thermal State
//
//   - [IProcessInfo.ThermalState]: The current thermal state of the system.
//
// # Determining Whether Low Power Mode is Enabled
//
//   - [IProcessInfo.LowPowerModeEnabled]: A Boolean value that indicates the current state of Low Power Mode.
//
// # Notifications
//
//   - [IProcessInfo.NSProcessInfoPowerStateDidChange]: Posts when the power state of a device changes.
//
// See: https://developer.apple.com/documentation/Foundation/ProcessInfo
type IProcessInfo interface {
	objectivec.IObject

	// Topic: Accessing Process Information

	// Array of strings with the command-line arguments for the process.
	Arguments() []string
	// The variable names (keys) and their values in the environment from which the process was launched.
	Environment() INSDictionary
	// Global unique identifier for the process.
	GloballyUniqueString() string
	// A Boolean value that indicates whether the process originated as an iOS app and runs on macOS.
	MacCatalystApp() bool
	// A Boolean value that indicates whether the process is an iPhone or iPad app running on a Mac.
	IOSAppOnMac() bool
	// A Boolean value that indicates whether the process is an iPhone or iPad app running on visionOS.
	IOSAppOnVision() bool
	// The identifier of the process (often called process ID).
	ProcessIdentifier() int
	// The name of the process.
	ProcessName() string
	SetProcessName(value string)

	// Topic: Accessing User Information

	// Returns the account name of the current user.
	UserName() string
	// Returns the full name of the current user.
	FullUserName() string

	// Topic: Sudden Application Termination

	// Disables the application for quickly killing using sudden termination.
	DisableSuddenTermination()
	// Enables the application for quick killing using sudden termination.
	EnableSuddenTermination()

	// Topic: Controlling Automatic Termination

	// Disables automatic termination for the application.
	DisableAutomaticTermination(reason string)
	// Enables automatic termination for the application.
	EnableAutomaticTermination(reason string)
	// A Boolean value indicating whether the app supports automatic termination.
	AutomaticTerminationSupportEnabled() bool
	SetAutomaticTerminationSupportEnabled(value bool)

	// Topic: Getting Host Information

	// The name of the host computer on which the process is executing.
	HostName() string
	// A string containing the version of the operating system on which the process is executing.
	OperatingSystemVersionString() string
	// The version of the operating system on which the process is executing.
	OperatingSystemVersion() OperatingSystemVersion
	// Returns a Boolean value indicating whether the version of the operating system on which the process is executing is the same or later than the given version.
	IsOperatingSystemAtLeastVersion(version NSOperatingSystemVersion) bool

	// Topic: Getting Computer Information

	// The number of processing cores available on the computer.
	ProcessorCount() uint
	// The number of active processing cores available on the computer.
	ActiveProcessorCount() uint
	// The amount of physical memory on the computer in bytes.
	PhysicalMemory() uint64
	// Indicates whether the device supports the requested performance tier.
	IsDeviceCertifiedFor(performanceTier int) bool
	// Indicates whether an app is running under a known performance profile.
	HasPerformanceProfile(performanceProfile int) bool
	// The amount of time the system has been awake since the last time it was restarted.
	SystemUptime() float64

	// Topic: Managing Activities

	// Begin an activity using the given options and reason.
	BeginActivityWithOptionsReason(options NSActivityOptions, reason string) objectivec.Object
	// Ends the given activity.
	EndActivity(activity objectivec.Object)
	// Synchronously perform an activity defined by a given block using the given options.
	PerformActivityWithOptionsReasonUsingBlock(options NSActivityOptions, reason string, block VoidHandler)

	// Topic: Getting the Thermal State

	// The current thermal state of the system.
	ThermalState() NSProcessInfoThermalState

	// Topic: Determining Whether Low Power Mode is Enabled

	// A Boolean value that indicates the current state of Low Power Mode.
	LowPowerModeEnabled() bool

	// Topic: Notifications

	// Posts when the power state of a device changes.
	NSProcessInfoPowerStateDidChange() NSNotificationName

	// A flag to indicate the activity requires the highest amount of timer and I/O precision available.
	LatencyCritical() NSActivityOptions
	SetLatencyCritical(value NSActivityOptions)
	// A flag to indicate the app is performing a user-requested action.
	UserInitiated() NSActivityOptions
	SetUserInitiated(value NSActivityOptions)
}

// Init initializes the instance.
func (p ProcessInfo) Init() ProcessInfo {
	rv := objc.Send[ProcessInfo](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p ProcessInfo) Autorelease() ProcessInfo {
	rv := objc.Send[ProcessInfo](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewProcessInfo creates a new ProcessInfo instance.
func NewProcessInfo() ProcessInfo {
	class := getProcessInfoClass()
	rv := objc.Send[ProcessInfo](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Disables the application for quickly killing using sudden termination.
//
// # Discussion
//
// This method increments the sudden termination counter. When the termination
// counter reaches `0` the application allows sudden termination.
//
// By default the sudden termination counter is set to 1. This can be
// overridden in your application Info.plist. See [NSProcessInfo] for more
// information and debugging suggestions.
//
// See: https://developer.apple.com/documentation/Foundation/ProcessInfo/disableSuddenTermination()
func (p ProcessInfo) DisableSuddenTermination() {
	objc.Send[objc.ID](p.ID, objc.Sel("disableSuddenTermination"))
}

// Enables the application for quick killing using sudden termination.
//
// # Discussion
//
// This method decrements the sudden termination counter. When the termination
// counter reaches `0` the application allows sudden termination.
//
// By default the sudden termination counter is set to 1. This can be
// overridden in your application Info.plist. See [NSProcessInfo] for more
// information and debugging suggestions.
//
// See: https://developer.apple.com/documentation/Foundation/ProcessInfo/enableSuddenTermination()
func (p ProcessInfo) EnableSuddenTermination() {
	objc.Send[objc.ID](p.ID, objc.Sel("enableSuddenTermination"))
}

// Disables automatic termination for the application.
//
// reason: The reason why automatic termination is being disabled.
//
// # Discussion
//
// This method increments the automatic termination counter. When the counter
// is greater than `0`, the application is considered active and ineligible
// for automatic termination. For example, you could disable automatic
// termination when the user of an instant messaging application signs on,
// because the application requires a background connection to be maintained
// even if the application is otherwise inactive.
//
// The reason parameter is used to track why an application is or is not
// automatically terminable and can be inspected by debugging tools. For
// example, you could pass the string `@"file transfer in progress"` if you
// disable automatic termination before transferring a file over the network.
// When you reenable automatic termination after the transfer is complete
// using [EnableAutomaticTermination], you should pass the matching string. A
// given reason can be used more than once at the same time; for example, if
// two files were being transferred at the same time, automatic termination
// could be disabled for each, passing the same reason string.
//
// See: https://developer.apple.com/documentation/Foundation/ProcessInfo/disableAutomaticTermination(_:)
func (p ProcessInfo) DisableAutomaticTermination(reason string) {
	objc.Send[objc.ID](p.ID, objc.Sel("disableAutomaticTermination:"), objc.String(reason))
}

// Enables automatic termination for the application.
//
// reason: The reason why automatic termination is being enabled.
//
// # Discussion
//
// This method decrements the automatic termination counter. When the counter
// is `0`, the application is eligible for automatic termination.
//
// The reason parameter is used to track why an application is or is not
// automatically terminable and can be inspected by debugging tools. For
// example, you could pass the string `@"file transfer in progress"` if you
// disable automatic termination before transferring a file over the network.
// When you reenable automatic termination after the transfer is complete
// using [EnableAutomaticTermination], you should pass the matching string. A
// given reason can be used more than once at the same time; for example, if
// two files were being transferred at the same time, automatic termination
// could be disabled for each, passing the same reason string.
//
// See: https://developer.apple.com/documentation/Foundation/ProcessInfo/enableAutomaticTermination(_:)
func (p ProcessInfo) EnableAutomaticTermination(reason string) {
	objc.Send[objc.ID](p.ID, objc.Sel("enableAutomaticTermination:"), objc.String(reason))
}

// Returns a Boolean value indicating whether the version of the operating
// system on which the process is executing is the same or later than the
// given version.
//
// version: The operating system version to test against.
//
// # Return Value
//
// true if the operating system on which the process is executing is the same
// or later than the given version; otherwise false.
//
// # Discussion
//
// This method accounts for major, minor, and update versions of the operating
// system.
//
// See: https://developer.apple.com/documentation/Foundation/ProcessInfo/isOperatingSystemAtLeast(_:)
func (p ProcessInfo) IsOperatingSystemAtLeastVersion(version NSOperatingSystemVersion) bool {
	rv := objc.Send[bool](p.ID, objc.Sel("isOperatingSystemAtLeastVersion:"), version)
	return rv
}

// Indicates whether the device supports the requested performance tier.
//
// performanceTier: The desired system performance tier. [iPhonePerformanceGaming] is the only
// performance tier.
//
// # Return Value
//
// [True] if the device meets the requirements for the given performance tier.
//
// See: https://developer.apple.com/documentation/Foundation/ProcessInfo/isDeviceCertified(for:)
//
// [iPhonePerformanceGaming]: https://developer.apple.com/documentation/Metal/NSDeviceCertification/iPhonePerformanceGaming
func (p ProcessInfo) IsDeviceCertifiedFor(performanceTier int) bool {
	rv := objc.Send[bool](p.ID, objc.Sel("isDeviceCertifiedFor:"), performanceTier)
	return rv
}

// Indicates whether an app is running under a known performance profile.
//
// performanceProfile: The desired performance profile. Choose between: [default] and [sustained].
//
// # Return Value
//
// True if the system is running under the given performance profile. If the
// profile isn’t [sustained], the app might cause the device to throttle
// under a heavy workload.
//
// See: https://developer.apple.com/documentation/Foundation/ProcessInfo/hasPerformanceProfile(_:)
//
// [default]: https://developer.apple.com/documentation/Metal/NSProcessPerformanceProfile/default
// [sustained]: https://developer.apple.com/documentation/Metal/NSProcessPerformanceProfile/sustained
//
// [sustained]: https://developer.apple.com/documentation/Metal/NSProcessPerformanceProfile/sustained
func (p ProcessInfo) HasPerformanceProfile(performanceProfile int) bool {
	rv := objc.Send[bool](p.ID, objc.Sel("hasPerformanceProfile:"), performanceProfile)
	return rv
}

// Begin an activity using the given options and reason.
//
// options: Options for the activity. See [ProcessInfo.ActivityOptions] for possible
// values.
//
// reason: A string used in debugging to indicate the reason the activity began.
//
// # Return Value
//
// An object token representing the activity.
//
// # Discussion
//
// Indicate completion of the activity by calling [EndActivity] passing the
// returned object as the argument.
//
// See: https://developer.apple.com/documentation/Foundation/ProcessInfo/beginActivity(options:reason:)
//
// [ProcessInfo.ActivityOptions]: https://developer.apple.com/documentation/Foundation/ProcessInfo/ActivityOptions
func (p ProcessInfo) BeginActivityWithOptionsReason(options NSActivityOptions, reason string) objectivec.Object {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("beginActivityWithOptions:reason:"), options, objc.String(reason))
	return objectivec.ObjectFromID(rv)
}

// Ends the given activity.
//
// activity: An activity object returned by [BeginActivityWithOptionsReason].
//
// See: https://developer.apple.com/documentation/Foundation/ProcessInfo/endActivity(_:)
func (p ProcessInfo) EndActivity(activity objectivec.Object) {
	objc.Send[objc.ID](p.ID, objc.Sel("endActivity:"), activity)
}

// Synchronously perform an activity defined by a given block using the given
// options.
//
// options: Options for the activity. See [ProcessInfo.ActivityOptions] for possible
// values.
//
// reason: A string used in debugging to indicate the reason the activity began.
//
// block: A block containing the work to be performed by the activity.
//
// # Discussion
//
// The activity will be automatically ended after `block` returns.
//
// See: https://developer.apple.com/documentation/Foundation/ProcessInfo/performActivity(options:reason:using:)
//
// [ProcessInfo.ActivityOptions]: https://developer.apple.com/documentation/Foundation/ProcessInfo/ActivityOptions
func (p ProcessInfo) PerformActivityWithOptionsReasonUsingBlock(options NSActivityOptions, reason string, block VoidHandler) {
	_block2, _ := NewVoidBlock(block)
	objc.Send[objc.ID](p.ID, objc.Sel("performActivityWithOptions:reason:usingBlock:"), options, objc.String(reason), _block2)
}

// Array of strings with the command-line arguments for the process.
//
// # Discussion
//
// This array contains all the information passed in the `argv` array,
// including the executable name in the first element.
//
// See: https://developer.apple.com/documentation/Foundation/ProcessInfo/arguments
func (p ProcessInfo) Arguments() []string {
	rv := objc.Send[[]objc.ID](p.ID, objc.Sel("arguments"))
	return objc.ConvertSliceToStrings(rv)
}

// The variable names (keys) and their values in the environment from which
// the process was launched.
//
// See: https://developer.apple.com/documentation/Foundation/ProcessInfo/environment
func (p ProcessInfo) Environment() INSDictionary {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("environment"))
	return NSDictionaryFromID(objc.ID(rv))
}

// Global unique identifier for the process.
//
// # Discussion
//
// The global ID for the process includes the host name, process ID, and a
// time stamp, which ensures that the ID is unique for the network. This
// property generates a new string each time its getter is invoked, and it
// uses a counter to guarantee that strings created from the same process are
// unique.
//
// See: https://developer.apple.com/documentation/Foundation/ProcessInfo/globallyUniqueString
func (p ProcessInfo) GloballyUniqueString() string {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("globallyUniqueString"))
	return NSStringFromID(rv).String()
}

// A Boolean value that indicates whether the process originated as an iOS app
// and runs on macOS.
//
// # Discussion
//
// The value of this property is true when the process is:
//
// - A Mac app built with Mac Catalyst, or an iOS app running on Apple
// silicon. - Running on a Mac.
//
// Frameworks that support iOS and macOS use this property to determine if the
// process is a Mac app built with Mac Catalyst. To conditionally compile
// source code intended to run only in macOS, use `#if
// targetEnvironment(macCatalyst)` (or `#if TARGET_OS_MACCATALYST` in
// Objective-C) instead.
//
// See: https://developer.apple.com/documentation/Foundation/ProcessInfo/isMacCatalystApp
func (p ProcessInfo) MacCatalystApp() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("isMacCatalystApp"))
	return rv
}

// A Boolean value that indicates whether the process is an iPhone or iPad app
// running on a Mac.
//
// # Discussion
//
// The value of this property is true only when the process is an iOS app
// running on a Mac. The value of the property is false for all other apps on
// the Mac, including Mac apps built using Mac Catalyst. The property is also
// false for processes running on platforms other than macOS.
//
// See: https://developer.apple.com/documentation/Foundation/ProcessInfo/isiOSAppOnMac
func (p ProcessInfo) IOSAppOnMac() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("isiOSAppOnMac"))
	return rv
}

// A Boolean value that indicates whether the process is an iPhone or iPad app
// running on visionOS.
//
// # Discussion
//
// The value of this property is true only when the process is an iOS app
// running on a visionOS device. The value of the property is false for all
// other apps on visionOS. The property is also false for processes running on
// platforms other than visonOS.
//
// See: https://developer.apple.com/documentation/Foundation/ProcessInfo/isiOSAppOnVision
func (p ProcessInfo) IOSAppOnVision() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("isiOSAppOnVision"))
	return rv
}

// The identifier of the process (often called process ID).
//
// See: https://developer.apple.com/documentation/Foundation/ProcessInfo/processIdentifier
func (p ProcessInfo) ProcessIdentifier() int {
	rv := objc.Send[int](p.ID, objc.Sel("processIdentifier"))
	return rv
}

// The name of the process.
//
// # Discussion
//
// The process name is used to register application defaults and is used in
// error messages. It does not uniquely identify the process.
//
// See: https://developer.apple.com/documentation/Foundation/ProcessInfo/processName
func (p ProcessInfo) ProcessName() string {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("processName"))
	return NSStringFromID(rv).String()
}
func (p ProcessInfo) SetProcessName(value string) {
	objc.Send[struct{}](p.ID, objc.Sel("setProcessName:"), objc.String(value))
}

// Returns the account name of the current user.
//
// See: https://developer.apple.com/documentation/Foundation/ProcessInfo/userName
func (p ProcessInfo) UserName() string {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("userName"))
	return NSStringFromID(rv).String()
}

// Returns the full name of the current user.
//
// See: https://developer.apple.com/documentation/Foundation/ProcessInfo/fullUserName
func (p ProcessInfo) FullUserName() string {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("fullUserName"))
	return NSStringFromID(rv).String()
}

// A Boolean value indicating whether the app supports automatic termination.
//
// # Discussion
//
// Without setting this property or setting the equivalent `Info.Plist()` key
// ([NSSupportsAutomaticTermination]), the methods
// [DisableAutomaticTermination] and [EnableAutomaticTermination] have no
// effect, although the counter tracking automatic termination opt-outs is
// still kept up to date to ensure correctness if this is called later.
// Currently, setting this property to false has no effect. This property
// should be set in the app delegate method
// [applicationDidFinishLaunching(_:)] or earlier.
//
// See: https://developer.apple.com/documentation/Foundation/ProcessInfo/automaticTerminationSupportEnabled
//
// [applicationDidFinishLaunching(_:)]: https://developer.apple.com/documentation/AppKit/NSApplicationDelegate/applicationDidFinishLaunching(_:)
func (p ProcessInfo) AutomaticTerminationSupportEnabled() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("automaticTerminationSupportEnabled"))
	return rv
}
func (p ProcessInfo) SetAutomaticTerminationSupportEnabled(value bool) {
	objc.Send[struct{}](p.ID, objc.Sel("setAutomaticTerminationSupportEnabled:"), value)
}

// The name of the host computer on which the process is executing.
//
// See: https://developer.apple.com/documentation/Foundation/ProcessInfo/hostName
func (p ProcessInfo) HostName() string {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("hostName"))
	return NSStringFromID(rv).String()
}

// A string containing the version of the operating system on which the
// process is executing.
//
// # Discussion
//
// The operating system version string is human readable, localized, and is
// appropriate for displaying to the user. This string is appropriate for
// parsing.
//
// See: https://developer.apple.com/documentation/Foundation/ProcessInfo/operatingSystemVersionString
func (p ProcessInfo) OperatingSystemVersionString() string {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("operatingSystemVersionString"))
	return NSStringFromID(rv).String()
}

// The version of the operating system on which the process is executing.
//
// See: https://developer.apple.com/documentation/Foundation/ProcessInfo/operatingSystemVersion
func (p ProcessInfo) OperatingSystemVersion() OperatingSystemVersion {
	rv := objc.Send[OperatingSystemVersion](p.ID, objc.Sel("operatingSystemVersion"))
	return OperatingSystemVersion(rv)
}

// The number of processing cores available on the computer.
//
// # Discussion
//
// This property value is equal to the result of entering the command `sysctl
// -n hw.Ncpu()` on the current system.
//
// See: https://developer.apple.com/documentation/Foundation/ProcessInfo/processorCount
func (p ProcessInfo) ProcessorCount() uint {
	rv := objc.Send[uint](p.ID, objc.Sel("processorCount"))
	return rv
}

// The number of active processing cores available on the computer.
//
// # Discussion
//
// Whereas the [ProcessorCount] property reports the number of advertised
// processing cores, the [ActiveProcessorCount] property reflects the actual
// number of active processing cores on the system. There are a number of
// different factors that may cause a core to not be active, including boot
// arguments, thermal throttling, or a manufacturing defect.
//
// This property value is equal to the result of entering the command `sysctl
// -n hw.Logicalcpu()` on the current system.
//
// See: https://developer.apple.com/documentation/Foundation/ProcessInfo/activeProcessorCount
func (p ProcessInfo) ActiveProcessorCount() uint {
	rv := objc.Send[uint](p.ID, objc.Sel("activeProcessorCount"))
	return rv
}

// The amount of physical memory on the computer in bytes.
//
// See: https://developer.apple.com/documentation/Foundation/ProcessInfo/physicalMemory
func (p ProcessInfo) PhysicalMemory() uint64 {
	rv := objc.Send[uint64](p.ID, objc.Sel("physicalMemory"))
	return rv
}

// The amount of time the system has been awake since the last time it was
// restarted.
//
// # Discussion
//
// See: https://developer.apple.com/documentation/Foundation/ProcessInfo/systemUptime
func (p ProcessInfo) SystemUptime() float64 {
	rv := objc.Send[NSTimeInterval](p.ID, objc.Sel("systemUptime"))
	return float64(rv)
}

// The current thermal state of the system.
//
// # Discussion
//
// At higher thermal states your app should reduce usage of system resources.
// For more information, see [ProcessInfo.ThermalState].
//
// See: https://developer.apple.com/documentation/Foundation/ProcessInfo/thermalState-swift.property
//
// [ProcessInfo.ThermalState]: https://developer.apple.com/documentation/Foundation/ProcessInfo/ThermalState-swift.enum
func (p ProcessInfo) ThermalState() NSProcessInfoThermalState {
	rv := objc.Send[NSProcessInfoThermalState](p.ID, objc.Sel("thermalState"))
	return NSProcessInfoThermalState(rv)
}

// A Boolean value that indicates the current state of Low Power Mode.
//
// # Discussion
//
// Users who wish to prolong their device’s battery life can enable Low
// Power Mode under Settings > Battery. In Low Power Mode, the system
// conserves battery life by enacting certain energy-saving measures, such as:
//
// - Reducing CPU and GPU performance. - Reducing screen brightness. - Pausing
// discretionary and background activities.
//
// Your app can query the [LowPowerModeEnabled] property at any time to
// determine whether Low Power Mode is active.
//
// Your app can also register to receive notifications when the Low Power Mode
// state of a device changes. To register for power state notifications, send
// the message [AddObserverSelectorNameObject] to the default notification
// center of your app (an instance of [NSNotificationCenter]). Pass it a
// selector to call and a notification name of
// [NSProcessInfoPowerStateDidChange]. When your app receives a notification
// of a power state change, query [LowPowerModeEnabled] to determine the
// current power state. If Low Power Mode is active, take appropriate steps to
// reduce activity in your app. Otherwise, your app can resume normal
// operations.
//
// For additional information, see [React to Low Power Mode on iPhones] in
// [Energy Efficiency Guide for iOS Apps].
//
// See: https://developer.apple.com/documentation/Foundation/ProcessInfo/isLowPowerModeEnabled
//
// [Energy Efficiency Guide for iOS Apps]: https://developer.apple.com/library/archive/documentation/Performance/Conceptual/EnergyGuide-iOS/index.html#//apple_ref/doc/uid/TP40015243
// [React to Low Power Mode on iPhones]: https://developer.apple.com/library/archive/documentation/Performance/Conceptual/EnergyGuide-iOS/LowPowerMode.html#//apple_ref/doc/uid/TP40015243-CH31
func (p ProcessInfo) LowPowerModeEnabled() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("isLowPowerModeEnabled"))
	return rv
}

// Posts when the power state of a device changes.
//
// See: https://developer.apple.com/documentation/foundation/nsnotification/name-swift.struct/nsprocessinfopowerstatedidchange
func (p ProcessInfo) NSProcessInfoPowerStateDidChange() NSNotificationName {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("NSProcessInfoPowerStateDidChangeNotification"))
	return NSNotificationName(NSStringFromID(rv).String())
}

// A flag to indicate the activity requires the highest amount of timer and
// I/O precision available.
//
// See: https://developer.apple.com/documentation/foundation/processinfo/activityoptions/latencycritical
func (p ProcessInfo) LatencyCritical() NSActivityOptions {
	rv := objc.Send[NSActivityOptions](p.ID, objc.Sel("NSActivityLatencyCritical"))
	return NSActivityOptions(rv)
}
func (p ProcessInfo) SetLatencyCritical(value NSActivityOptions) {
	objc.Send[struct{}](p.ID, objc.Sel("setNSActivityLatencyCritical:"), value)
}

// A flag to indicate the app is performing a user-requested action.
//
// See: https://developer.apple.com/documentation/foundation/processinfo/activityoptions/userinitiated
func (p ProcessInfo) UserInitiated() NSActivityOptions {
	rv := objc.Send[NSActivityOptions](p.ID, objc.Sel("NSActivityUserInitiated"))
	return NSActivityOptions(rv)
}
func (p ProcessInfo) SetUserInitiated(value NSActivityOptions) {
	objc.Send[struct{}](p.ID, objc.Sel("setNSActivityUserInitiated:"), value)
}

// Returns the process information agent for the process.
//
// # Return Value
//
// Shared process information agent for the process.
//
// # Discussion
//
// An [NSProcessInfo] object is created the first time this method is invoked,
// and that same object is returned on each subsequent invocation.
//
// See: https://developer.apple.com/documentation/Foundation/ProcessInfo/processInfo
func (_ProcessInfoClass ProcessInfoClass) ProcessInfo() ProcessInfo {
	rv := objc.Send[objc.ID](objc.ID(_ProcessInfoClass.class), objc.Sel("processInfo"))
	return NSProcessInfoFromID(objc.ID(rv))
}

// PerformActivityWithOptionsReasonUsingBlockSync is a synchronous wrapper around [ProcessInfo.PerformActivityWithOptionsReasonUsingBlock].
// It blocks until the completion handler fires or the context is cancelled.
func (p ProcessInfo) PerformActivityWithOptionsReasonUsingBlockSync(ctx context.Context, options NSActivityOptions, reason string) error {
	done := make(chan struct{}, 1)
	p.PerformActivityWithOptionsReasonUsingBlock(options, reason, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
