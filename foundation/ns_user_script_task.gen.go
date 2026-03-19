// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSUserScriptTask] class.
var (
	_NSUserScriptTaskClass     NSUserScriptTaskClass
	_NSUserScriptTaskClassOnce sync.Once
)

func getNSUserScriptTaskClass() NSUserScriptTaskClass {
	_NSUserScriptTaskClassOnce.Do(func() {
		_NSUserScriptTaskClass = NSUserScriptTaskClass{class: objc.GetClass("NSUserScriptTask")}
	})
	return _NSUserScriptTaskClass
}

// GetNSUserScriptTaskClass returns the class object for NSUserScriptTask.
func GetNSUserScriptTaskClass() NSUserScriptTaskClass {
	return getNSUserScriptTaskClass()
}

type NSUserScriptTaskClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSUserScriptTaskClass) Alloc() NSUserScriptTask {
	rv := objc.Send[NSUserScriptTask](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// An object that executes scripts.
//
// # Overview
// 
// The [NSUserScriptTask] class is able to run all the scripts normally run by
// the one of its subclasses, however it ignores the results. It is intended
// to execute user-supplied scripts and will execute them outside of the
// application’s sandbox, if any.
// 
// If you need to execute scripts and get the input and output information use
// the [NSUserUnixTask], [NSUserAppleScriptTask], and [NSUserAutomatorTask]
// sub classes.
//
// # Specifying the Script
//
//   - [NSUserScriptTask.InitWithURLError]: Return a user script task instance given a URL for a script file.
//   - [NSUserScriptTask.ScriptURL]: The URL of the script file.
//
// # Executing the User Script
//
//   - [NSUserScriptTask.ExecuteWithCompletionHandler]: Executes the script with no input and ignoring any result.
//
// See: https://developer.apple.com/documentation/Foundation/NSUserScriptTask
type NSUserScriptTask struct {
	objectivec.Object
}

// NSUserScriptTaskFromID constructs a [NSUserScriptTask] from an objc.ID.
//
// An object that executes scripts.
func NSUserScriptTaskFromID(id objc.ID) NSUserScriptTask {
	return NSUserScriptTask{objectivec.Object{ID: id}}
}
// NOTE: NSUserScriptTask adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSUserScriptTask] class.
//
// # Specifying the Script
//
//   - [INSUserScriptTask.InitWithURLError]: Return a user script task instance given a URL for a script file.
//   - [INSUserScriptTask.ScriptURL]: The URL of the script file.
//
// # Executing the User Script
//
//   - [INSUserScriptTask.ExecuteWithCompletionHandler]: Executes the script with no input and ignoring any result.
//
// See: https://developer.apple.com/documentation/Foundation/NSUserScriptTask
type INSUserScriptTask interface {
	objectivec.IObject

	// Topic: Specifying the Script

	// Return a user script task instance given a URL for a script file.
	InitWithURLError(url INSURL) (NSUserScriptTask, error)
	// The URL of the script file.
	ScriptURL() INSURL

	// Topic: Executing the User Script

	// Executes the script with no input and ignoring any result.
	ExecuteWithCompletionHandler(handler ErrorHandler)
}

// Init initializes the instance.
func (u NSUserScriptTask) Init() NSUserScriptTask {
	rv := objc.Send[NSUserScriptTask](u.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u NSUserScriptTask) Autorelease() NSUserScriptTask {
	rv := objc.Send[NSUserScriptTask](u.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSUserScriptTask creates a new NSUserScriptTask instance.
func NewNSUserScriptTask() NSUserScriptTask {
	class := getNSUserScriptTaskClass()
	rv := objc.Send[NSUserScriptTask](objc.ID(class.class), objc.Sel("new"))
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
func NewUserScriptTaskWithURLError(url INSURL) (NSUserScriptTask, error) {
	var errorPtr objc.ID
	instance := getNSUserScriptTaskClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithURL:error:"), url, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return NSUserScriptTaskFromID(rv), NSErrorFrom(errorPtr)
	}
	return NSUserScriptTaskFromID(rv), nil
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
func (u NSUserScriptTask) InitWithURLError(url INSURL) (NSUserScriptTask, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](u.ID, objc.Sel("initWithURL:error:"), url, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return NSUserScriptTask{}, NSErrorFrom(errorPtr)
	}
	return NSUserScriptTaskFromID(rv), nil

}
// Executes the script with no input and ignoring any result.
//
// handler: The completion handler Block that returns the result or an error. See
// [NSUserScriptTaskCompletionHandler].
//
// # Discussion
// 
// This method should be invoked no more than once for a given instance of the
// class.
// 
// If the script completed normally, the completion handler’s `error`
// parameter will be `nil`.
//
// See: https://developer.apple.com/documentation/Foundation/NSUserScriptTask/execute(completionHandler:)
func (u NSUserScriptTask) ExecuteWithCompletionHandler(handler ErrorHandler) {
_block0, _cleanup0 := NewErrorBlock(handler)
	defer _cleanup0()
	objc.Send[objc.ID](u.ID, objc.Sel("executeWithCompletionHandler:"), _block0)
}

// The URL of the script file.
//
// # Return Value
// 
// The URL of the script file.
//
// See: https://developer.apple.com/documentation/Foundation/NSUserScriptTask/scriptURL
func (u NSUserScriptTask) ScriptURL() INSURL {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("scriptURL"))
	return NSURLFromID(objc.ID(rv))
}

