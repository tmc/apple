// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [NSSetCommand] class.
var (
	_NSSetCommandClass     NSSetCommandClass
	_NSSetCommandClassOnce sync.Once
)

func getNSSetCommandClass() NSSetCommandClass {
	_NSSetCommandClassOnce.Do(func() {
		_NSSetCommandClass = NSSetCommandClass{class: objc.GetClass("NSSetCommand")}
	})
	return _NSSetCommandClass
}

// GetNSSetCommandClass returns the class object for NSSetCommand.
func GetNSSetCommandClass() NSSetCommandClass {
	return getNSSetCommandClass()
}

type NSSetCommandClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSSetCommandClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSSetCommandClass) Alloc() NSSetCommand {
	rv := objc.Send[NSSetCommand](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A command that sets one or more attributes or relationships to one or more
// values.
//
// # Overview
//
// An instance of [NSSetCommand] sets one or more attributes or relationships
// to one or more values; for example, it may set the (x, y) coordinates for a
// window’s position or set the name of a document.
//
// [NSSetCommand] is part of Cocoa’s built-in scripting support. It works
// automatically to support the `set` command through key-value coding. Most
// applications don’t need to subclass [NSSetCommand] or call its methods.
//
// [NSSetCommand] uses available scripting class descriptions to determine
// whether it should set a value for an attribute (or property), or set a
// value for all elements (to-many objects). For the latter, it invokes
// [replaceValue(at:inPropertyWithKey:withValue:)]; for the former, it invokes
// [setValue(_:forKey:)] (or, if the receiver overrides
// [takeValue(_:forKey:)], it invokes that method, to support backward binary
// compatibility.)
//
// For information on working with `set` commands, see [Getting and Setting
// Properties and Elements] in [Cocoa Scripting Guide].
//
// # Working with specifiers
//
//   - [NSSetCommand.KeySpecifier]: Returns a specifier that identifies the attribute or relationship that is to be set for the receiver of the `set` AppleScript command.
//
// See: https://developer.apple.com/documentation/Foundation/NSSetCommand
//
// [Cocoa Scripting Guide]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/ScriptableCocoaApplications/SApps_intro/SAppsIntro.html#//apple_ref/doc/uid/TP40002164
// [Getting and Setting Properties and Elements]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/ScriptableCocoaApplications/SApps_get_set/SAppsGetSet.html#//apple_ref/doc/uid/TP40002164-CH18
// [replaceValue(at:inPropertyWithKey:withValue:)]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/replaceValue(at:inPropertyWithKey:withValue:)
// [setValue(_:forKey:)]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/setValue(_:forKey:)
// [takeValue(_:forKey:)]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/takeValue(_:forKey:)
type NSSetCommand struct {
	NSScriptCommand
}

// NSSetCommandFromID constructs a [NSSetCommand] from an objc.ID.
//
// A command that sets one or more attributes or relationships to one or more
// values.
func NSSetCommandFromID(id objc.ID) NSSetCommand {
	return NSSetCommand{NSScriptCommand: NSScriptCommandFromID(id)}
}

// NOTE: NSSetCommand adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSSetCommand] class.
//
// # Working with specifiers
//
//   - [INSSetCommand.KeySpecifier]: Returns a specifier that identifies the attribute or relationship that is to be set for the receiver of the `set` AppleScript command.
//
// See: https://developer.apple.com/documentation/Foundation/NSSetCommand
type INSSetCommand interface {
	INSScriptCommand

	// Topic: Working with specifiers

	// Returns a specifier that identifies the attribute or relationship that is to be set for the receiver of the `set` AppleScript command.
	KeySpecifier() INSScriptObjectSpecifier
}

// Init initializes the instance.
func (s NSSetCommand) Init() NSSetCommand {
	rv := objc.Send[NSSetCommand](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s NSSetCommand) Autorelease() NSSetCommand {
	rv := objc.Send[NSSetCommand](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSSetCommand creates a new NSSetCommand instance.
func NewNSSetCommand() NSSetCommand {
	class := getNSSetCommandClass()
	rv := objc.Send[NSSetCommand](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Foundation/NSScriptCommand/init(coder:)
func NewSetCommandWithCoder(inCoder INSCoder) NSSetCommand {
	instance := getNSSetCommandClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), inCoder)
	return NSSetCommandFromID(rv)
}

// Returns an a script command object initialized from the passed command
// description.
//
// commandDef: A command description for the command to be created.
//
// # Return Value
//
// A newly initialized instance of [NSScriptCommand] or a subclass.
//
// # Discussion
//
// To make this command object usable, you must set its receiving objects and
// arguments (if any) after invoking this method.
//
// See: https://developer.apple.com/documentation/Foundation/NSScriptCommand/init(commandDescription:)
func NewSetCommandWithCommandDescription(commandDef INSScriptCommandDescription) NSSetCommand {
	instance := getNSSetCommandClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCommandDescription:"), commandDef)
	return NSSetCommandFromID(rv)
}

// Returns a specifier that identifies the attribute or relationship that is
// to be set for the receiver of the `set` AppleScript command.
//
// # Return Value
//
// A specifier that identifies the attribute or relationship that is to be set
// for the receiver of the `set` AppleScript command.
//
// See: https://developer.apple.com/documentation/Foundation/NSSetCommand/keySpecifier
func (s NSSetCommand) KeySpecifier() INSScriptObjectSpecifier {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("keySpecifier"))
	return NSScriptObjectSpecifierFromID(objc.ID(rv))
}
