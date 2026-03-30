// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"errors"
	"sync"
	"unsafe"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSTask] class.
var (
	_NSTaskClass     NSTaskClass
	_NSTaskClassOnce sync.Once
)

func getNSTaskClass() NSTaskClass {
	_NSTaskClassOnce.Do(func() {
		_NSTaskClass = NSTaskClass{class: objc.GetClass("NSTask")}
	})
	return _NSTaskClass
}

// GetNSTaskClass returns the class object for NSTask.
func GetNSTaskClass() NSTaskClass {
	return getNSTaskClass()
}

type NSTaskClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSTaskClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSTaskClass) Alloc() NSTask {
	rv := objc.Send[NSTask](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// An object that represents a subprocess of the current process.
//
// # Overview
//
// Using this class, your program can run another program as a subprocess and
// monitor that program’s execution. Unlike [NSThread], it doesn’t share
// memory space with the process that creates it.
//
// A process operates within an environment defined by the current values for
// several items: the current directory, standard input, standard output,
// standard error, and the values of any environment variables, inheriting its
// environment from the process that launches it. If there are any environment
// variables that should be different for the subprocess (for example, if the
// current directory needs to change), change it in the instance after
// initialization, before your app launches it. Your app can’t change a
// process’s environment while it’s running.
//
// You can only run the subprocess once per instance. Subsequent attempts
// raise an error.
//
// # Returning Information
//
//   - [NSTask.ProcessIdentifier]: The receiver’s process identifier.
//
// # Running and Stopping
//
//   - [NSTask.LaunchAndReturnError]: Runs the process with the current environment.
//   - [NSTask.Interrupt]: Sends an interrupt signal to the receiver and all of its subtasks.
//   - [NSTask.Resume]: Resumes execution of a suspended task.
//   - [NSTask.Suspend]: Suspends execution of the receiver task.
//   - [NSTask.Terminate]: Sends a terminate signal to the receiver and all of its subtasks.
//   - [NSTask.WaitUntilExit]: Blocks the process until the receiver is finished.
//
// # Querying the State
//
//   - [NSTask.Running]: A status that indicates whether the receiver is still running.
//   - [NSTask.TerminationStatus]: The exit status the receiver’s executable returns.
//   - [NSTask.TerminationReason]: The reason the system terminated the task.
//
// # Configuring
//
//   - [NSTask.Arguments]: The command arguments that the system uses to launch the executable.
//   - [NSTask.SetArguments]
//   - [NSTask.CurrentDirectoryURL]: The current directory for the receiver.
//   - [NSTask.SetCurrentDirectoryURL]
//   - [NSTask.Environment]: The environment for the receiver.
//   - [NSTask.SetEnvironment]
//   - [NSTask.ExecutableURL]: The receiver’s executable.
//   - [NSTask.SetExecutableURL]
//   - [NSTask.QualityOfService]: The default quality of service level the system applies to operations the task executes.
//   - [NSTask.SetQualityOfService]
//   - [NSTask.StandardError]: The standard error for the receiver.
//   - [NSTask.SetStandardError]
//   - [NSTask.StandardInput]: The standard input for the receiver.
//   - [NSTask.SetStandardInput]
//   - [NSTask.StandardOutput]: The standard output for the receiver.
//   - [NSTask.SetStandardOutput]
//
// # Termination Handler
//
//   - [NSTask.TerminationHandler]: A completion block the system invokes when the task completes.
//   - [NSTask.SetTerminationHandler]
//
// # Deprecated
//
//   - [NSTask.CurrentDirectoryPath]: Sets the current directory for the receiver.
//   - [NSTask.SetCurrentDirectoryPath]
//   - [NSTask.LaunchPath]: Sets the receiver’s executable.
//   - [NSTask.SetLaunchPath]
//
// # Instance Properties
//
//   - [NSTask.LaunchRequirement]
//   - [NSTask.SetLaunchRequirement]
//   - [NSTask.LaunchRequirementData]
//   - [NSTask.SetLaunchRequirementData]
//
// See: https://developer.apple.com/documentation/Foundation/Process
type NSTask struct {
	objectivec.Object
}

// NSTaskFromID constructs a [NSTask] from an objc.ID.
//
// An object that represents a subprocess of the current process.
func NSTaskFromID(id objc.ID) NSTask {
	return NSTask{objectivec.Object{ID: id}}
}

// NOTE: NSTask adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSTask] class.
//
// # Returning Information
//
//   - [INSTask.ProcessIdentifier]: The receiver’s process identifier.
//
// # Running and Stopping
//
//   - [INSTask.LaunchAndReturnError]: Runs the process with the current environment.
//   - [INSTask.Interrupt]: Sends an interrupt signal to the receiver and all of its subtasks.
//   - [INSTask.Resume]: Resumes execution of a suspended task.
//   - [INSTask.Suspend]: Suspends execution of the receiver task.
//   - [INSTask.Terminate]: Sends a terminate signal to the receiver and all of its subtasks.
//   - [INSTask.WaitUntilExit]: Blocks the process until the receiver is finished.
//
// # Querying the State
//
//   - [INSTask.Running]: A status that indicates whether the receiver is still running.
//   - [INSTask.TerminationStatus]: The exit status the receiver’s executable returns.
//   - [INSTask.TerminationReason]: The reason the system terminated the task.
//
// # Configuring
//
//   - [INSTask.Arguments]: The command arguments that the system uses to launch the executable.
//   - [INSTask.SetArguments]
//   - [INSTask.CurrentDirectoryURL]: The current directory for the receiver.
//   - [INSTask.SetCurrentDirectoryURL]
//   - [INSTask.Environment]: The environment for the receiver.
//   - [INSTask.SetEnvironment]
//   - [INSTask.ExecutableURL]: The receiver’s executable.
//   - [INSTask.SetExecutableURL]
//   - [INSTask.QualityOfService]: The default quality of service level the system applies to operations the task executes.
//   - [INSTask.SetQualityOfService]
//   - [INSTask.StandardError]: The standard error for the receiver.
//   - [INSTask.SetStandardError]
//   - [INSTask.StandardInput]: The standard input for the receiver.
//   - [INSTask.SetStandardInput]
//   - [INSTask.StandardOutput]: The standard output for the receiver.
//   - [INSTask.SetStandardOutput]
//
// # Termination Handler
//
//   - [INSTask.TerminationHandler]: A completion block the system invokes when the task completes.
//   - [INSTask.SetTerminationHandler]
//
// # Deprecated
//
//   - [INSTask.CurrentDirectoryPath]: Sets the current directory for the receiver.
//   - [INSTask.SetCurrentDirectoryPath]
//   - [INSTask.LaunchPath]: Sets the receiver’s executable.
//   - [INSTask.SetLaunchPath]
//
// # Instance Properties
//
//   - [INSTask.LaunchRequirement]
//   - [INSTask.SetLaunchRequirement]
//   - [INSTask.LaunchRequirementData]
//   - [INSTask.SetLaunchRequirementData]
//
// See: https://developer.apple.com/documentation/Foundation/Process
type INSTask interface {
	objectivec.IObject

	// Topic: Returning Information

	// The receiver’s process identifier.
	ProcessIdentifier() int

	// Topic: Running and Stopping

	// Runs the process with the current environment.
	LaunchAndReturnError() (bool, error)
	// Sends an interrupt signal to the receiver and all of its subtasks.
	Interrupt()
	// Resumes execution of a suspended task.
	Resume() bool
	// Suspends execution of the receiver task.
	Suspend() bool
	// Sends a terminate signal to the receiver and all of its subtasks.
	Terminate()
	// Blocks the process until the receiver is finished.
	WaitUntilExit()

	// Topic: Querying the State

	// A status that indicates whether the receiver is still running.
	Running() bool
	// The exit status the receiver’s executable returns.
	TerminationStatus() int
	// The reason the system terminated the task.
	TerminationReason() NSTaskTerminationReason

	// Topic: Configuring

	// The command arguments that the system uses to launch the executable.
	Arguments() []string
	SetArguments(value []string)
	// The current directory for the receiver.
	CurrentDirectoryURL() INSURL
	SetCurrentDirectoryURL(value INSURL)
	// The environment for the receiver.
	Environment() INSDictionary
	SetEnvironment(value INSDictionary)
	// The receiver’s executable.
	ExecutableURL() INSURL
	SetExecutableURL(value INSURL)
	// The default quality of service level the system applies to operations the task executes.
	QualityOfService() QualityOfService
	SetQualityOfService(value QualityOfService)
	// The standard error for the receiver.
	StandardError() objectivec.IObject
	SetStandardError(value objectivec.IObject)
	// The standard input for the receiver.
	StandardInput() objectivec.IObject
	SetStandardInput(value objectivec.IObject)
	// The standard output for the receiver.
	StandardOutput() objectivec.IObject
	SetStandardOutput(value objectivec.IObject)

	// Topic: Termination Handler

	// A completion block the system invokes when the task completes.
	TerminationHandler() TaskHandler
	SetTerminationHandler(value TaskHandler)

	// Topic: Deprecated

	// Sets the current directory for the receiver.
	CurrentDirectoryPath() string
	SetCurrentDirectoryPath(value string)
	// Sets the receiver’s executable.
	LaunchPath() string
	SetLaunchPath(value string)

	// Topic: Instance Properties

	LaunchRequirement() objectivec.IObject
	SetLaunchRequirement(value objectivec.IObject)
	LaunchRequirementData() INSData
	SetLaunchRequirementData(value INSData)
}

// Init initializes the instance.
func (t NSTask) Init() NSTask {
	rv := objc.Send[NSTask](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t NSTask) Autorelease() NSTask {
	rv := objc.Send[NSTask](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSTask creates a new NSTask instance.
func NewNSTask() NSTask {
	class := getNSTaskClass()
	rv := objc.Send[NSTask](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Runs the process with the current environment.
//
// See: https://developer.apple.com/documentation/Foundation/Process/run()
func (t NSTask) LaunchAndReturnError() (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](t.ID, objc.Sel("launchAndReturnError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("launchAndReturnError: returned NO with nil NSError")
	}
	return rv, nil

}

// Sends an interrupt signal to the receiver and all of its subtasks.
//
// # Discussion
//
// If the task terminates as a result, which is the default behavior, an
// [didTerminateNotification] gets sent to the default notification center.
// This method has no effect if the receiver was already launched and has
// already finished executing. If the system hasn’t launched the receiver,
// this method raises an [NSInvalidArgumentException].
//
// It isn’t always possible to interrupt the receiver because it might be
// ignoring the interrupt signal. The [Interrupt] method sends [SIGINT].
//
// See: https://developer.apple.com/documentation/Foundation/Process/interrupt()
//
// [didTerminateNotification]: https://developer.apple.com/documentation/Foundation/Process/didTerminateNotification
func (t NSTask) Interrupt() {
	objc.Send[objc.ID](t.ID, objc.Sel("interrupt"))
}

// Resumes execution of a suspended task.
//
// # Return Value
//
// true if the receiver was able to resume execution, false otherwise.
//
// # Discussion
//
// If the system sent multiple [Suspend] messages to the receiver, an equal
// number of [Resume] messages must be sent before the task resumes execution.
//
// See: https://developer.apple.com/documentation/Foundation/Process/resume()
func (t NSTask) Resume() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("resume"))
	return rv
}

// Suspends execution of the receiver task.
//
// # Return Value
//
// true if the receiver was successfully suspended, false otherwise.
//
// # Discussion
//
// Multiple [Suspend] messages can be sent, but they must be balanced with an
// equal number of [Resume] messages before the task resumes execution.
//
// See: https://developer.apple.com/documentation/Foundation/Process/suspend()
func (t NSTask) Suspend() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("suspend"))
	return rv
}

// Sends a terminate signal to the receiver and all of its subtasks.
//
// # Discussion
//
// If the task terminates as a result, which is the default behavior, an
// [didTerminateNotification] gets sent to the default notification center.
// This method has no effect if the receiver was already launched and has
// already finished executing. If the receiver hasn’t been launched yet,
// this method raises an [NSInvalidArgumentException].
//
// It’s not always possible to terminate the receiver because it might be
// ignoring the terminate signal. The [Terminate] method sends [SIGTERM].
//
// See: https://developer.apple.com/documentation/Foundation/Process/terminate()
//
// [didTerminateNotification]: https://developer.apple.com/documentation/Foundation/Process/didTerminateNotification
func (t NSTask) Terminate() {
	objc.Send[objc.ID](t.ID, objc.Sel("terminate"))
}

// Blocks the process until the receiver is finished.
//
// # Discussion
//
// This method first checks to see if the receiver is still running using
// [Running]. Then it polls the current run loop using [NSDefaultRunLoopMode]
// until the task completes.
//
// [WaitUntilExit] does not guarantee that the [TerminationHandler] block has
// been fully executed before [WaitUntilExit] returns.
//
// See: https://developer.apple.com/documentation/Foundation/Process/waitUntilExit()
func (t NSTask) WaitUntilExit() {
	objc.Send[objc.ID](t.ID, objc.Sel("waitUntilExit"))
}

// Creates and runs a task with a specified executable and arguments.
//
// url: The URL for the executable.
//
// arguments: An array of [NSString] objects that supplies the arguments to the task. If
// `arguments` is `nil`, the system raises an [NSInvalidArgumentException].
//
// terminationHandler: The system invokes this completion block when the task has completed.
//
// # Return Value
//
// An initialized [NSTask] object with the environment of the current process.
//
// See: https://developer.apple.com/documentation/Foundation/Process/run(_:arguments:terminationHandler:)
func (_NSTaskClass NSTaskClass) LaunchedTaskWithExecutableURLArgumentsErrorTerminationHandler(url INSURL, arguments []string, error_ INSError, terminationHandler TaskHandler) NSTask {
	_block3, _ := NewTaskBlock(terminationHandler)
	rv := objc.Send[objc.ID](objc.ID(_NSTaskClass.class), objc.Sel("launchedTaskWithExecutableURL:arguments:error:terminationHandler:"), url, arguments, error_, _block3)
	return NSTaskFromID(rv)
}

// The receiver’s process identifier.
//
// # Return Value
//
// The receiver’s process identifier.
//
// See: https://developer.apple.com/documentation/Foundation/Process/processIdentifier
func (t NSTask) ProcessIdentifier() int {
	rv := objc.Send[int](t.ID, objc.Sel("processIdentifier"))
	return rv
}

// A status that indicates whether the receiver is still running.
//
// # Return Value
//
// true if the receiver is still running, otherwise false. false means either
// the receiver could not run or it has terminated.
//
// See: https://developer.apple.com/documentation/Foundation/Process/isRunning
func (t NSTask) Running() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("isRunning"))
	return rv
}

// The exit status the receiver’s executable returns.
//
// # Return Value
//
// The exit status returned by the receiver’s executable.
//
// # Discussion
//
// Each task defines and documents how your app should interpret the return
// value. For example, many commands return 0 if they complete successfully or
// an error code if they don’t. You’ll need to look at the documentation
// for that task to learn what values it returns under what circumstances.
//
// This method raises an [NSInvalidArgumentException] if the receiver is still
// running. Verify that the receiver isn’t running before you use it.
//
// See: https://developer.apple.com/documentation/Foundation/Process/terminationStatus
func (t NSTask) TerminationStatus() int {
	rv := objc.Send[int](t.ID, objc.Sel("terminationStatus"))
	return rv
}

// The reason the system terminated the task.
//
// # Return Value
//
// The termination status. The possible values are described in
// [Process.TerminationReason].
//
// See: https://developer.apple.com/documentation/Foundation/Process/terminationReason-swift.property
//
// [Process.TerminationReason]: https://developer.apple.com/documentation/Foundation/Process/TerminationReason-swift.enum
func (t NSTask) TerminationReason() NSTaskTerminationReason {
	rv := objc.Send[NSTaskTerminationReason](t.ID, objc.Sel("terminationReason"))
	return NSTaskTerminationReason(rv)
}

// The command arguments that the system uses to launch the executable.
//
// # Discussion
//
// The [NSTask] object converts both `path` and the strings in `arguments` to
// appropriate C-style strings (using [FileSystemRepresentation]) before
// passing them to the task through `argv[]`. The strings in `arguments`
// don’t undergo shell expansion, so you don’t need to do special quoting,
// and shell variables, such as `$PWD`, aren’t resolved.
//
// See: https://developer.apple.com/documentation/Foundation/Process/arguments
func (t NSTask) Arguments() []string {
	rv := objc.Send[[]objc.ID](t.ID, objc.Sel("arguments"))
	return objc.ConvertSliceToStrings(rv)
}
func (t NSTask) SetArguments(value []string) {
	objc.Send[struct{}](t.ID, objc.Sel("setArguments:"), objectivec.StringSliceToNSArray(value))
}

// The current directory for the receiver.
//
// See: https://developer.apple.com/documentation/Foundation/Process/currentDirectoryURL
func (t NSTask) CurrentDirectoryURL() INSURL {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("currentDirectoryURL"))
	return NSURLFromID(objc.ID(rv))
}
func (t NSTask) SetCurrentDirectoryURL(value INSURL) {
	objc.Send[struct{}](t.ID, objc.Sel("setCurrentDirectoryURL:"), value)
}

// The environment for the receiver.
//
// # Discussion
//
// If this method isn’t used, the environment is inherited from the process
// that created the receiver. This method raises an
// [NSInvalidArgumentException] if the system has launched the receiver.
//
// See: https://developer.apple.com/documentation/Foundation/Process/environment
func (t NSTask) Environment() INSDictionary {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("environment"))
	return NSDictionaryFromID(objc.ID(rv))
}
func (t NSTask) SetEnvironment(value INSDictionary) {
	objc.Send[struct{}](t.ID, objc.Sel("setEnvironment:"), value)
}

// The receiver’s executable.
//
// See: https://developer.apple.com/documentation/Foundation/Process/executableURL
func (t NSTask) ExecutableURL() INSURL {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("executableURL"))
	return NSURLFromID(objc.ID(rv))
}
func (t NSTask) SetExecutableURL(value INSURL) {
	objc.Send[struct{}](t.ID, objc.Sel("setExecutableURL:"), value)
}

// The default quality of service level the system applies to operations the
// task executes.
//
// See: https://developer.apple.com/documentation/Foundation/Process/qualityOfService
func (t NSTask) QualityOfService() QualityOfService {
	rv := objc.Send[QualityOfService](t.ID, objc.Sel("qualityOfService"))
	return QualityOfService(rv)
}
func (t NSTask) SetQualityOfService(value QualityOfService) {
	objc.Send[struct{}](t.ID, objc.Sel("setQualityOfService:"), value)
}

// The standard error for the receiver.
//
// # Discussion
//
// If `file` is an [NSPipe] object, launching the receiver automatically
// closes the write end of the pipe in the current task. Don’t create a
// handle for the pipe and pass that as the argument, or the system won’t
// automatically close the write end of the pipe.
//
// If this method isn’t used, the standard error is inherited from the
// process that created the receiver. This method raises an
// [NSInvalidArgumentException] if the system has lauched the receiver.
//
// See: https://developer.apple.com/documentation/Foundation/Process/standardError
func (t NSTask) StandardError() objectivec.IObject {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("standardError"))
	return objectivec.Object{ID: rv}
}
func (t NSTask) SetStandardError(value objectivec.IObject) {
	objc.Send[struct{}](t.ID, objc.Sel("setStandardError:"), value)
}

// The standard input for the receiver.
//
// # Discussion
//
// If `file` is an [NSPipe] object, launching the receiver automatically
// closes the read end of the pipe in the current task. Don’t create a
// handle for the pipe and pass that as the argument, or the read end of the
// pipe won’t be closed automatically.
//
// If this method isn’t used, the standard input is inherited from the
// process that created the receiver. This method raises an
// [NSInvalidArgumentException] if the system has lauched the receiver.
//
// See: https://developer.apple.com/documentation/Foundation/Process/standardInput
func (t NSTask) StandardInput() objectivec.IObject {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("standardInput"))
	return objectivec.Object{ID: rv}
}
func (t NSTask) SetStandardInput(value objectivec.IObject) {
	objc.Send[struct{}](t.ID, objc.Sel("setStandardInput:"), value)
}

// The standard output for the receiver.
//
// # Discussion
//
// If `file` is an [NSPipe] object, launching the receiver automatically
// closes the write end of the pipe in the current task. Don’t create a
// handle for the pipe and pass that as the argument, or the write end of the
// pipe won’t be closed automatically.
//
// If this method isn’t used, the standard output is inherited from the
// process that created the receiver. This method raises an
// [NSInvalidArgumentException] if the system has lauched the receiver.
//
// See: https://developer.apple.com/documentation/Foundation/Process/standardOutput
func (t NSTask) StandardOutput() objectivec.IObject {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("standardOutput"))
	return objectivec.Object{ID: rv}
}
func (t NSTask) SetStandardOutput(value objectivec.IObject) {
	objc.Send[struct{}](t.ID, objc.Sel("setStandardOutput:"), value)
}

// A completion block the system invokes when the task completes.
//
// # Discussion
//
// The system passes the task object to the block to allow access to the task
// parameters, for example to determine if the task completed successfully.
//
// This block isn’t guaranteed to be fully executed prior to [WaitUntilExit]
// returning.
//
// See: https://developer.apple.com/documentation/Foundation/Process/terminationHandler
func (t NSTask) TerminationHandler() TaskHandler {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("terminationHandler"))
	_ = rv
	return nil
}
func (t NSTask) SetTerminationHandler(value TaskHandler) {
	block, cleanup := NewTaskBlock(value)
	defer cleanup()
	objc.Send[struct{}](t.ID, objc.Sel("setTerminationHandler:"), block)
}

// Sets the current directory for the receiver.
//
// # Discussion
//
// If this method isn’t used, the current directory is inherited from the
// process that created the receiver. This method raises an
// [NSInvalidArgumentException] if the receiver has already been launched.
//
// See: https://developer.apple.com/documentation/Foundation/Process/currentDirectoryPath
func (t NSTask) CurrentDirectoryPath() string {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("currentDirectoryPath"))
	return NSStringFromID(rv).String()
}
func (t NSTask) SetCurrentDirectoryPath(value string) {
	objc.Send[struct{}](t.ID, objc.Sel("setCurrentDirectoryPath:"), objc.String(value))
}

// Sets the receiver’s executable.
//
// See: https://developer.apple.com/documentation/Foundation/Process/launchPath
func (t NSTask) LaunchPath() string {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("launchPath"))
	return NSStringFromID(rv).String()
}
func (t NSTask) SetLaunchPath(value string) {
	objc.Send[struct{}](t.ID, objc.Sel("setLaunchPath:"), objc.String(value))
}

// See: https://developer.apple.com/documentation/foundation/process/launchrequirement
func (t NSTask) LaunchRequirement() objectivec.IObject {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("launchRequirement"))
	return objectivec.Object{ID: rv}
}
func (t NSTask) SetLaunchRequirement(value objectivec.IObject) {
	objc.Send[struct{}](t.ID, objc.Sel("setLaunchRequirement:"), value)
}

// See: https://developer.apple.com/documentation/Foundation/Process/launchRequirementData
func (t NSTask) LaunchRequirementData() INSData {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("launchRequirementData"))
	return NSDataFromID(objc.ID(rv))
}
func (t NSTask) SetLaunchRequirementData(value INSData) {
	objc.Send[struct{}](t.ID, objc.Sel("setLaunchRequirementData:"), value)
}
