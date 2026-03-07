// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [NSUserAutomatorTask] class.
var (
	_NSUserAutomatorTaskClass     NSUserAutomatorTaskClass
	_NSUserAutomatorTaskClassOnce sync.Once
)

func getNSUserAutomatorTaskClass() NSUserAutomatorTaskClass {
	_NSUserAutomatorTaskClassOnce.Do(func() {
		_NSUserAutomatorTaskClass = NSUserAutomatorTaskClass{class: objc.GetClass("NSUserAutomatorTask")}
	})
	return _NSUserAutomatorTaskClass
}

// GetNSUserAutomatorTaskClass returns the class object for NSUserAutomatorTask.
func GetNSUserAutomatorTaskClass() NSUserAutomatorTaskClass {
	return getNSUserAutomatorTaskClass()
}

type NSUserAutomatorTaskClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSUserAutomatorTaskClass) Alloc() NSUserAutomatorTask {
	rv := objc.Send[NSUserAutomatorTask](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// An object that executes Automator workflows.
//
// # Overview
// 
// The [NSUserAutomatorTask] class is intended to run Automator workflows from
// your application. It is intended to execute user-supplied workflows, and
// will execute them outside of the application’s sandbox, if any.
// 
// The class is not intended to execute scripts built into an application; for
// that, use one of the [NSTask] or [AMWorkflow] classes. If the application
// is sandboxed, then the script must be in the [ApplicationScriptsDirectory]
// folder. A sandboxed application may read from, but not write to, this
// folder.
// 
// If you simply need to execute scripts without regard to input or output,
// use [NSUserScriptTask], which can execute any of the specific types. If you
// need specific control over the input to or output from the workflow, use
// this class.
//
// [AMWorkflow]: https://developer.apple.com/documentation/Automator/AMWorkflow
//
// # Executing Automator Tasks
//
//   - [NSUserAutomatorTask.ExecuteWithInputCompletionHandler]: Execute the Automator workflow by providing it as securely coded input.
//   - [NSUserAutomatorTask.Variables]: The variables required by the Automator workflow.
//   - [NSUserAutomatorTask.SetVariables]
//
// See: https://developer.apple.com/documentation/Foundation/NSUserAutomatorTask
type NSUserAutomatorTask struct {
	NSUserScriptTask
}

// NSUserAutomatorTaskFromID constructs a [NSUserAutomatorTask] from an objc.ID.
//
// An object that executes Automator workflows.
func NSUserAutomatorTaskFromID(id objc.ID) NSUserAutomatorTask {
	return NSUserAutomatorTask{NSUserScriptTask: NSUserScriptTaskFromID(id)}
}
// NOTE: NSUserAutomatorTask adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [NSUserAutomatorTask] class.
//
// # Executing Automator Tasks
//
//   - [INSUserAutomatorTask.ExecuteWithInputCompletionHandler]: Execute the Automator workflow by providing it as securely coded input.
//   - [INSUserAutomatorTask.Variables]: The variables required by the Automator workflow.
//   - [INSUserAutomatorTask.SetVariables]
//
// See: https://developer.apple.com/documentation/Foundation/NSUserAutomatorTask
type INSUserAutomatorTask interface {
	INSUserScriptTask

	// Topic: Executing Automator Tasks

	// Execute the Automator workflow by providing it as securely coded input.
	ExecuteWithInputCompletionHandler(input NSSecureCoding, handler ErrorHandler)
	// The variables required by the Automator workflow.
	Variables() INSDictionary
	SetVariables(value INSDictionary)
}





// Init initializes the instance.
func (u NSUserAutomatorTask) Init() NSUserAutomatorTask {
	rv := objc.Send[NSUserAutomatorTask](u.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u NSUserAutomatorTask) Autorelease() NSUserAutomatorTask {
	rv := objc.Send[NSUserAutomatorTask](u.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSUserAutomatorTask creates a new NSUserAutomatorTask instance.
func NewNSUserAutomatorTask() NSUserAutomatorTask {
	class := getNSUserAutomatorTaskClass()
	rv := objc.Send[NSUserAutomatorTask](objc.ID(class.class), objc.Sel("new"))
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
func NewUserAutomatorTaskWithURLError(url INSURL) (NSUserAutomatorTask, error) {
	var errorPtr objc.ID
	instance := getNSUserAutomatorTaskClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithURL:error:"), url, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return NSUserAutomatorTaskFromID(rv), NSErrorFrom(errorPtr)
	}
	return NSUserAutomatorTaskFromID(rv), nil
}







// Execute the Automator workflow by providing it as securely coded input.
//
// input: The automator task.
//
// handler: The completion handler Block that returns the result or an error. See
// [NSUserAutomatorTaskCompletionHandler].
//
// # Discussion
// 
// The Automator workflow will execute using the [Variables] property values.
// 
// This method should be invoked no more than once for a given instance of the
// class.
// 
// If the script completed normally, the completion handler’s `error`
// parameter will be `nil`.
//
// See: https://developer.apple.com/documentation/Foundation/NSUserAutomatorTask/execute(withInput:completionHandler:)
func (u NSUserAutomatorTask) ExecuteWithInputCompletionHandler(input NSSecureCoding, handler ErrorHandler) {
		_block1, _cleanup1 := NewErrorBlock(handler)
	defer _cleanup1()
		objc.Send[objc.ID](u.ID, objc.Sel("executeWithInput:completionHandler:"), input, _block1)
}












// The variables required by the Automator workflow.
//
// See: https://developer.apple.com/documentation/Foundation/NSUserAutomatorTask/variables
func (u NSUserAutomatorTask) Variables() INSDictionary {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("variables"))
	return NSDictionaryFromID(objc.ID(rv))
}
func (u NSUserAutomatorTask) SetVariables(value INSDictionary) {
	objc.Send[struct{}](u.ID, objc.Sel("setVariables:"), value)
}

























