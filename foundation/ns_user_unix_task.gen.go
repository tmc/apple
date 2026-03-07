// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [NSUserUnixTask] class.
var (
	_NSUserUnixTaskClass     NSUserUnixTaskClass
	_NSUserUnixTaskClassOnce sync.Once
)

func getNSUserUnixTaskClass() NSUserUnixTaskClass {
	_NSUserUnixTaskClassOnce.Do(func() {
		_NSUserUnixTaskClass = NSUserUnixTaskClass{class: objc.GetClass("NSUserUnixTask")}
	})
	return _NSUserUnixTaskClass
}

// GetNSUserUnixTaskClass returns the class object for NSUserUnixTask.
func GetNSUserUnixTaskClass() NSUserUnixTaskClass {
	return getNSUserUnixTaskClass()
}

type NSUserUnixTaskClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSUserUnixTaskClass) Alloc() NSUserUnixTask {
	rv := objc.Send[NSUserUnixTask](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// An object that executes unix applications.
//
// # Overview
// 
// The [NSUserUnixTask] class is intended to run unix applications, typically
// a shell script, from your application. It is intended to execute
// user-supplied scripts, and will execute them outside of the application’s
// sandbox, if any.
// 
// The class is not intended to execute scripts built into an application; for
// that, use one of the [NSTask], [NSAppleScript], or [AMWorkflow] classes. If
// the application is sandboxed, then the script must be in the
// [ApplicationScriptsDirectory] folder. A sandboxed application may read
// from, but not write to, this folder.
// 
// If you simply need to execute unix scripts without regard to input or
// output, use [NSUserScriptTask], which can execute any of the specific
// types. If you need specific control over the input to, or output from, or
// the error stream of the script, use this class.
//
// [AMWorkflow]: https://developer.apple.com/documentation/Automator/AMWorkflow
//
// # Executing the Unix Script
//
//   - [NSUserUnixTask.ExecuteWithArgumentsCompletionHandler]: Execute the unix script with the specified arguments.
//
// # Standard Unix Streams
//
//   - [NSUserUnixTask.StandardError]: The standard error stream.
//   - [NSUserUnixTask.SetStandardError]
//   - [NSUserUnixTask.StandardInput]: The standard input stream.
//   - [NSUserUnixTask.SetStandardInput]
//   - [NSUserUnixTask.StandardOutput]: The standard output stream.
//   - [NSUserUnixTask.SetStandardOutput]
//
// See: https://developer.apple.com/documentation/Foundation/NSUserUnixTask
type NSUserUnixTask struct {
	NSUserScriptTask
}

// NSUserUnixTaskFromID constructs a [NSUserUnixTask] from an objc.ID.
//
// An object that executes unix applications.
func NSUserUnixTaskFromID(id objc.ID) NSUserUnixTask {
	return NSUserUnixTask{NSUserScriptTask: NSUserScriptTaskFromID(id)}
}
// NOTE: NSUserUnixTask adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [NSUserUnixTask] class.
//
// # Executing the Unix Script
//
//   - [INSUserUnixTask.ExecuteWithArgumentsCompletionHandler]: Execute the unix script with the specified arguments.
//
// # Standard Unix Streams
//
//   - [INSUserUnixTask.StandardError]: The standard error stream.
//   - [INSUserUnixTask.SetStandardError]
//   - [INSUserUnixTask.StandardInput]: The standard input stream.
//   - [INSUserUnixTask.SetStandardInput]
//   - [INSUserUnixTask.StandardOutput]: The standard output stream.
//   - [INSUserUnixTask.SetStandardOutput]
//
// See: https://developer.apple.com/documentation/Foundation/NSUserUnixTask
type INSUserUnixTask interface {
	INSUserScriptTask

	// Topic: Executing the Unix Script

	// Execute the unix script with the specified arguments.
	ExecuteWithArgumentsCompletionHandler(arguments []string, handler ErrorHandler)

	// Topic: Standard Unix Streams

	// The standard error stream.
	StandardError() INSFileHandle
	SetStandardError(value INSFileHandle)
	// The standard input stream.
	StandardInput() INSFileHandle
	SetStandardInput(value INSFileHandle)
	// The standard output stream.
	StandardOutput() INSFileHandle
	SetStandardOutput(value INSFileHandle)
}





// Init initializes the instance.
func (u NSUserUnixTask) Init() NSUserUnixTask {
	rv := objc.Send[NSUserUnixTask](u.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u NSUserUnixTask) Autorelease() NSUserUnixTask {
	rv := objc.Send[NSUserUnixTask](u.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSUserUnixTask creates a new NSUserUnixTask instance.
func NewNSUserUnixTask() NSUserUnixTask {
	class := getNSUserUnixTaskClass()
	rv := objc.Send[NSUserUnixTask](objc.ID(class.class), objc.Sel("new"))
	return rv
}






// Return a user script task instance given a URL for a script file.
//
// url: The script URL.
//
// # Return Value
// 
// An instance of an [NSUserScriptTask] subclass or `nil` if the file does not
// appear to match any of the known types.
//
// # Discussion
// 
// The returned object will be of one of the specific sub-classes
// ([NSUserUnixTask], [NSUserAppleScriptTask], and [NSUserAutomatorTask]), or
// `nil` if the file does not appear to match any of the known types.
// 
// If invoked from a subclass, the result will be that class or `nil`.
//
// See: https://developer.apple.com/documentation/Foundation/NSUserScriptTask/init(url:)
func NewUserUnixTaskWithURLError(url INSURL) (NSUserUnixTask, error) {
	var errorPtr objc.ID
	instance := getNSUserUnixTaskClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithURL:error:"), url, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return NSUserUnixTaskFromID(rv), NSErrorFrom(errorPtr)
	}
	return NSUserUnixTaskFromID(rv), nil
}







// Execute the unix script with the specified arguments.
//
// arguments: An array of [NSString] objects containing the script arguments. The
// arguments do not undergo shell expansion, so you do not need to do special
// quoting, and shell variables are not resolved.
//
// handler: The completion handler Block that returns the result. See
// [NSUserUnixTaskCompletionHandler].
//
// # Discussion
// 
// This method should be invoked no more than once for a given instance of the
// class.
// 
// If the script completed normally, the completion handler’s `error`
// parameter will be `nil`.
//
// See: https://developer.apple.com/documentation/Foundation/NSUserUnixTask/execute(withArguments:completionHandler:)
func (u NSUserUnixTask) ExecuteWithArgumentsCompletionHandler(arguments []string, handler ErrorHandler) {
		_block1, _cleanup1 := NewErrorBlock(handler)
	defer _cleanup1()
		objc.Send[objc.ID](u.ID, objc.Sel("executeWithArguments:completionHandler:"), arguments, _block1)
}












// The standard error stream.
//
// # Discussion
// 
// Setting to `nil` will bind the stream to `/dev/null`.
// 
// The default is `nil`.
//
// See: https://developer.apple.com/documentation/Foundation/NSUserUnixTask/standardError
func (u NSUserUnixTask) StandardError() INSFileHandle {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("standardError"))
	return NSFileHandleFromID(objc.ID(rv))
}
func (u NSUserUnixTask) SetStandardError(value INSFileHandle) {
	objc.Send[struct{}](u.ID, objc.Sel("setStandardError:"), value)
}



// The standard input stream.
//
// # Discussion
// 
// Setting to `nil` will bind the stream to `/dev/null`.
// 
// The default is `nil`.
//
// See: https://developer.apple.com/documentation/Foundation/NSUserUnixTask/standardInput
func (u NSUserUnixTask) StandardInput() INSFileHandle {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("standardInput"))
	return NSFileHandleFromID(objc.ID(rv))
}
func (u NSUserUnixTask) SetStandardInput(value INSFileHandle) {
	objc.Send[struct{}](u.ID, objc.Sel("setStandardInput:"), value)
}



// The standard output stream.
//
// # Discussion
// 
// Setting to `nil` will bind the stream to `/dev/null`.
// 
// The default is `nil`.
//
// See: https://developer.apple.com/documentation/Foundation/NSUserUnixTask/standardOutput
func (u NSUserUnixTask) StandardOutput() INSFileHandle {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("standardOutput"))
	return NSFileHandleFromID(objc.ID(rv))
}
func (u NSUserUnixTask) SetStandardOutput(value INSFileHandle) {
	objc.Send[struct{}](u.ID, objc.Sel("setStandardOutput:"), value)
}

























