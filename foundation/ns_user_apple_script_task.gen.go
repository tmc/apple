// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [NSUserAppleScriptTask] class.
var (
	_NSUserAppleScriptTaskClass     NSUserAppleScriptTaskClass
	_NSUserAppleScriptTaskClassOnce sync.Once
)

func getNSUserAppleScriptTaskClass() NSUserAppleScriptTaskClass {
	_NSUserAppleScriptTaskClassOnce.Do(func() {
		_NSUserAppleScriptTaskClass = NSUserAppleScriptTaskClass{class: objc.GetClass("NSUserAppleScriptTask")}
	})
	return _NSUserAppleScriptTaskClass
}

// GetNSUserAppleScriptTaskClass returns the class object for NSUserAppleScriptTask.
func GetNSUserAppleScriptTaskClass() NSUserAppleScriptTaskClass {
	return getNSUserAppleScriptTaskClass()
}

type NSUserAppleScriptTaskClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSUserAppleScriptTaskClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSUserAppleScriptTaskClass) Alloc() NSUserAppleScriptTask {
	rv := objc.Send[NSUserAppleScriptTask](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// An object that executes AppleScript scripts.
//
// # Overview
// 
// The [NSUserAppleScriptTask] class is intended to run AppleScript scripts
// from your application. It is intended to execute user-supplied scripts and
// will execute them outside of the application’s sandbox, if any.
// 
// The class is not intended to execute scripts built into an application; for
// that, use one of the [NSTask] classes. If the application is sandboxed,
// then the script must be in the [ApplicationScriptsDirectory] folder. A
// sandboxed application may read from, but not write to, this folder.
// 
// If you simply need to execute scripts without regard to input or output,
// use [NSUserScriptTask], which can execute any of the specific types. If you
// need specific control over the input to or output from the script, use this
// class.
//
// # Executing an AppleScript Script
//
//   - [NSUserAppleScriptTask.ExecuteWithAppleEventCompletionHandler]: Execute the AppleScript script by sending it the specified Apple event.
//
// See: https://developer.apple.com/documentation/Foundation/NSUserAppleScriptTask
type NSUserAppleScriptTask struct {
	NSUserScriptTask
}

// NSUserAppleScriptTaskFromID constructs a [NSUserAppleScriptTask] from an objc.ID.
//
// An object that executes AppleScript scripts.
func NSUserAppleScriptTaskFromID(id objc.ID) NSUserAppleScriptTask {
	return NSUserAppleScriptTask{NSUserScriptTask: NSUserScriptTaskFromID(id)}
}
// NOTE: NSUserAppleScriptTask adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSUserAppleScriptTask] class.
//
// # Executing an AppleScript Script
//
//   - [INSUserAppleScriptTask.ExecuteWithAppleEventCompletionHandler]: Execute the AppleScript script by sending it the specified Apple event.
//
// See: https://developer.apple.com/documentation/Foundation/NSUserAppleScriptTask
type INSUserAppleScriptTask interface {
	INSUserScriptTask

	// Topic: Executing an AppleScript Script

	// Execute the AppleScript script by sending it the specified Apple event.
	ExecuteWithAppleEventCompletionHandler(event INSAppleEventDescriptor, handler ErrorHandler)
}

// Init initializes the instance.
func (u NSUserAppleScriptTask) Init() NSUserAppleScriptTask {
	rv := objc.Send[NSUserAppleScriptTask](u.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u NSUserAppleScriptTask) Autorelease() NSUserAppleScriptTask {
	rv := objc.Send[NSUserAppleScriptTask](u.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSUserAppleScriptTask creates a new NSUserAppleScriptTask instance.
func NewNSUserAppleScriptTask() NSUserAppleScriptTask {
	class := getNSUserAppleScriptTaskClass()
	rv := objc.Send[NSUserAppleScriptTask](objc.ID(class.class), objc.Sel("new"))
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
func NewUserAppleScriptTaskWithURLError(url INSURL) (NSUserAppleScriptTask, error) {
	var errorPtr objc.ID
	instance := getNSUserAppleScriptTaskClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithURL:error:"), url, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return NSUserAppleScriptTask{}, NSErrorFrom(errorPtr)
	}
	return NSUserAppleScriptTaskFromID(rv), nil
}

// Execute the AppleScript script by sending it the specified Apple event.
//
// event: The Apple event.
//
// handler: The completion handler Block that returns the result or an error. See
// [NSUserAppleScriptTaskCompletionHandler].
//
// # Discussion
// 
// Pass `nil` as `event` to execute the script’s default “run” handler.
// 
// This method should be invoked no more than once for a given instance of the
// class.
// 
// If the script completed normally, the completion handler’s `error`
// parameter will be `nil`.
//
// See: https://developer.apple.com/documentation/Foundation/NSUserAppleScriptTask/execute(withAppleEvent:completionHandler:)
func (u NSUserAppleScriptTask) ExecuteWithAppleEventCompletionHandler(event INSAppleEventDescriptor, handler ErrorHandler) {
_block1, _ := NewErrorBlock(handler)
	objc.Send[objc.ID](u.ID, objc.Sel("executeWithAppleEvent:completionHandler:"), event, _block1)
}

