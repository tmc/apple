// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSAppleScript] class.
var (
	_NSAppleScriptClass     NSAppleScriptClass
	_NSAppleScriptClassOnce sync.Once
)

func getNSAppleScriptClass() NSAppleScriptClass {
	_NSAppleScriptClassOnce.Do(func() {
		_NSAppleScriptClass = NSAppleScriptClass{class: objc.GetClass("NSAppleScript")}
	})
	return _NSAppleScriptClass
}

// GetNSAppleScriptClass returns the class object for NSAppleScript.
func GetNSAppleScriptClass() NSAppleScriptClass {
	return getNSAppleScriptClass()
}

type NSAppleScriptClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSAppleScriptClass) Alloc() NSAppleScript {
	rv := objc.Send[NSAppleScript](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// An object that provides the ability to load, compile, and execute scripts.
//
// # Overview
// 
// This class provides applications with the ability to
// 
// - load a script from a URL or from a text string - compile or execute a
// script or an individual Apple event - obtain an [NSAppleEventDescriptor]
// containing the reply from an executed script or event - obtain an
// attributed string for a compiled script, suitable for display in a script
// editor - obtain various kinds of information about any errors that may
// occur
// 
// When you create an instance of [NSAppleScript] object, you can use a URL to
// specify a script that can be in either text or compiled form, or you can
// supply the script as a string. Should an error occur when compiling or
// executing the script, several of the methods return a dictionary containing
// error information. The keys for obtaining error information, such as
// [errorMessage], are described in the Constants section.
// 
// See also NSAppleScript Additions Reference in the Application Kit
// framework, which defines a method that returns the syntax-highlighted
// source code for a script.
//
// [errorMessage]: https://developer.apple.com/documentation/Foundation/NSAppleScript/errorMessage
//
// # Initializing a Script
//
//   - [NSAppleScript.InitWithContentsOfURLError]: Initializes a newly allocated script instance from the source identified by the passed URL.
//   - [NSAppleScript.InitWithSource]: Initializes a newly allocated script instance from the passed source.
//
// # Getting Information About a Script
//
//   - [NSAppleScript.Compiled]: A Boolean value that indicates whether the receiver’s script has been compiled.
//   - [NSAppleScript.Source]: The script source for the receiver.
//
// # Compiling and Executing a Script
//
//   - [NSAppleScript.CompileAndReturnError]: Compiles the receiver, if it is not already compiled.
//   - [NSAppleScript.ExecuteAndReturnError]: Executes the receiver, compiling it first if it is not already compiled.
//   - [NSAppleScript.ExecuteAppleEventError]: Executes an Apple event in the context of the receiver, as a means of allowing the application to invoke a handler in the script.
//
// # Instance Properties
//
//   - [NSAppleScript.RichTextSource]: Returns the syntax-highlighted source code of the receiver if the receiver has been compiled and its source code is available.
//
// See: https://developer.apple.com/documentation/Foundation/NSAppleScript
type NSAppleScript struct {
	objectivec.Object
}

// NSAppleScriptFromID constructs a [NSAppleScript] from an objc.ID.
//
// An object that provides the ability to load, compile, and execute scripts.
func NSAppleScriptFromID(id objc.ID) NSAppleScript {
	return NSAppleScript{objectivec.Object{ID: id}}
}
// NOTE: NSAppleScript adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSAppleScript] class.
//
// # Initializing a Script
//
//   - [INSAppleScript.InitWithContentsOfURLError]: Initializes a newly allocated script instance from the source identified by the passed URL.
//   - [INSAppleScript.InitWithSource]: Initializes a newly allocated script instance from the passed source.
//
// # Getting Information About a Script
//
//   - [INSAppleScript.Compiled]: A Boolean value that indicates whether the receiver’s script has been compiled.
//   - [INSAppleScript.Source]: The script source for the receiver.
//
// # Compiling and Executing a Script
//
//   - [INSAppleScript.CompileAndReturnError]: Compiles the receiver, if it is not already compiled.
//   - [INSAppleScript.ExecuteAndReturnError]: Executes the receiver, compiling it first if it is not already compiled.
//   - [INSAppleScript.ExecuteAppleEventError]: Executes an Apple event in the context of the receiver, as a means of allowing the application to invoke a handler in the script.
//
// # Instance Properties
//
//   - [INSAppleScript.RichTextSource]: Returns the syntax-highlighted source code of the receiver if the receiver has been compiled and its source code is available.
//
// See: https://developer.apple.com/documentation/Foundation/NSAppleScript
type INSAppleScript interface {
	objectivec.IObject
	NSCopying

	// Topic: Initializing a Script

	// Initializes a newly allocated script instance from the source identified by the passed URL.
	InitWithContentsOfURLError(url INSURL, errorInfo INSDictionary) NSAppleScript
	// Initializes a newly allocated script instance from the passed source.
	InitWithSource(source string) NSAppleScript

	// Topic: Getting Information About a Script

	// A Boolean value that indicates whether the receiver’s script has been compiled.
	Compiled() bool
	// The script source for the receiver.
	Source() string

	// Topic: Compiling and Executing a Script

	// Compiles the receiver, if it is not already compiled.
	CompileAndReturnError(errorInfo INSDictionary) bool
	// Executes the receiver, compiling it first if it is not already compiled.
	ExecuteAndReturnError(errorInfo INSDictionary) INSAppleEventDescriptor
	// Executes an Apple event in the context of the receiver, as a means of allowing the application to invoke a handler in the script.
	ExecuteAppleEventError(event INSAppleEventDescriptor, errorInfo INSDictionary) INSAppleEventDescriptor

	// Topic: Instance Properties

	// Returns the syntax-highlighted source code of the receiver if the receiver has been compiled and its source code is available.
	RichTextSource() INSAttributedString
}

// Init initializes the instance.
func (a NSAppleScript) Init() NSAppleScript {
	rv := objc.Send[NSAppleScript](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a NSAppleScript) Autorelease() NSAppleScript {
	rv := objc.Send[NSAppleScript](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSAppleScript creates a new NSAppleScript instance.
func NewNSAppleScript() NSAppleScript {
	class := getNSAppleScriptClass()
	rv := objc.Send[NSAppleScript](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Initializes a newly allocated script instance from the source identified by
// the passed URL.
//
// url: A URL that locates a script, in either text or compiled form.
//
// errorInfo: On return, if an error occurs, a pointer to an error information
// dictionary.
//
// # Return Value
// 
// The initialized script object, `nil` if an error occurs.
//
// # Discussion
// 
// This method is a designated initializer for [NSAppleScript].
//
// See: https://developer.apple.com/documentation/Foundation/NSAppleScript/init(contentsOf:error:)
func NewAppleScriptWithContentsOfURLError(url INSURL, errorInfo INSDictionary) NSAppleScript {
	instance := getNSAppleScriptClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithContentsOfURL:error:"), url, errorInfo)
	return NSAppleScriptFromID(rv)
}

// Initializes a newly allocated script instance from the passed source.
//
// source: A string containing the source code of a script.
//
// # Return Value
// 
// The initialized script object, `nil` if an error occurs.
//
// # Discussion
// 
// This method is a designated initializer for [NSAppleScript].
//
// See: https://developer.apple.com/documentation/Foundation/NSAppleScript/init(source:)
func NewAppleScriptWithSource(source string) NSAppleScript {
	instance := getNSAppleScriptClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSource:"), objc.String(source))
	return NSAppleScriptFromID(rv)
}

// Initializes a newly allocated script instance from the source identified by
// the passed URL.
//
// url: A URL that locates a script, in either text or compiled form.
//
// errorInfo: On return, if an error occurs, a pointer to an error information
// dictionary.
//
// # Return Value
// 
// The initialized script object, `nil` if an error occurs.
//
// # Discussion
// 
// This method is a designated initializer for [NSAppleScript].
//
// See: https://developer.apple.com/documentation/Foundation/NSAppleScript/init(contentsOf:error:)
func (a NSAppleScript) InitWithContentsOfURLError(url INSURL, errorInfo INSDictionary) NSAppleScript {
	rv := objc.Send[NSAppleScript](a.ID, objc.Sel("initWithContentsOfURL:error:"), url, errorInfo)
	return rv
}
// Initializes a newly allocated script instance from the passed source.
//
// source: A string containing the source code of a script.
//
// # Return Value
// 
// The initialized script object, `nil` if an error occurs.
//
// # Discussion
// 
// This method is a designated initializer for [NSAppleScript].
//
// See: https://developer.apple.com/documentation/Foundation/NSAppleScript/init(source:)
func (a NSAppleScript) InitWithSource(source string) NSAppleScript {
	rv := objc.Send[NSAppleScript](a.ID, objc.Sel("initWithSource:"), objc.String(source))
	return rv
}
// Compiles the receiver, if it is not already compiled.
//
// errorInfo: On return, if an error occurs, a pointer to an error information
// dictionary.
//
// # Return Value
// 
// [true] for success or if the script was already compiled, [false]
// otherwise.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Foundation/NSAppleScript/compileAndReturnError(_:)
func (a NSAppleScript) CompileAndReturnError(errorInfo INSDictionary) bool {
	rv := objc.Send[bool](a.ID, objc.Sel("compileAndReturnError:"), errorInfo)
	return rv
}
// Executes the receiver, compiling it first if it is not already compiled.
//
// errorInfo: On return, if an error occurs, a pointer to an error information
// dictionary.
//
// # Return Value
// 
// The result of executing the event, or `nil` if an error occurs.
//
// # Discussion
// 
// Any changes to property values caused by executing the script do not
// persist.
//
// See: https://developer.apple.com/documentation/Foundation/NSAppleScript/executeAndReturnError(_:)
func (a NSAppleScript) ExecuteAndReturnError(errorInfo INSDictionary) INSAppleEventDescriptor {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("executeAndReturnError:"), errorInfo)
	return NSAppleEventDescriptorFromID(rv)
}
// Executes an Apple event in the context of the receiver, as a means of
// allowing the application to invoke a handler in the script.
//
// event: The Apple event to execute.
//
// errorInfo: On return, if an error occurs, a pointer to an error information
// dictionary.
//
// # Return Value
// 
// The result of executing the event, or `nil` if an error occurs.
//
// # Discussion
// 
// Compiles the receiver before executing it if it is not already compiled.
//
// See: https://developer.apple.com/documentation/Foundation/NSAppleScript/executeAppleEvent(_:error:)
func (a NSAppleScript) ExecuteAppleEventError(event INSAppleEventDescriptor, errorInfo INSDictionary) INSAppleEventDescriptor {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("executeAppleEvent:error:"), event, errorInfo)
	return NSAppleEventDescriptorFromID(rv)
}

// A Boolean value that indicates whether the receiver’s script has been
// compiled.
//
// See: https://developer.apple.com/documentation/Foundation/NSAppleScript/isCompiled
func (a NSAppleScript) Compiled() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("isCompiled"))
	return rv
}
// The script source for the receiver.
//
// # Discussion
// 
// It is possible for an [NSAppleScript] that has been instantiated with
// [InitWithContentsOfURLError] to be a script for which the source code is
// not available but is nonetheless executable.
//
// See: https://developer.apple.com/documentation/Foundation/NSAppleScript/source
func (a NSAppleScript) Source() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("source"))
	return NSStringFromID(rv).String()
}
// Returns the syntax-highlighted source code of the receiver if the receiver
// has been compiled and its source code is available.
//
// # Discussion
// 
// Returns `nil` otherwise. It is possible for an instance of [NSAppleScript]
// that has been instantiated with [InitWithContentsOfURLError] to be a script
// for which the source code is not available, but is nonetheless executable.
//
// See: https://developer.apple.com/documentation/Foundation/NSAppleScript/richTextSource
func (a NSAppleScript) RichTextSource() INSAttributedString {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("richTextSource"))
	return NSAttributedStringFromID(objc.ID(rv))
}

			// Protocol methods for NSCopying
			

