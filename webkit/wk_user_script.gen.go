// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [WKUserScript] class.
var (
	_WKUserScriptClass     WKUserScriptClass
	_WKUserScriptClassOnce sync.Once
)

func getWKUserScriptClass() WKUserScriptClass {
	_WKUserScriptClassOnce.Do(func() {
		_WKUserScriptClass = WKUserScriptClass{class: objc.GetClass("WKUserScript")}
	})
	return _WKUserScriptClass
}

// GetWKUserScriptClass returns the class object for WKUserScript.
func GetWKUserScriptClass() WKUserScriptClass {
	return getWKUserScriptClass()
}

type WKUserScriptClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (wc WKUserScriptClass) Class() objc.Class {
	return wc.class
}

// Alloc allocates memory for a new instance of the class.
func (wc WKUserScriptClass) Alloc() WKUserScript {
	rv := objc.Send[WKUserScript](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

// A script that the web view injects into a webpage.
//
// # Overview
//
// Create a [WKUserScript] object when you want to inject custom script code
// into the pages of your web view. Use this object to specify the JavaScript
// code to inject, and parameters relating to when and how to inject that
// code. Before you create the web view, add this object to the
// [WKUserContentController] object associated with your web view’s
// configuration.
//
// # Creating a User Script Object
//
//   - [WKUserScript.InitWithSourceInjectionTimeForMainFrameOnly]: Creates a user script object that contains the specified source code and attributes.
//   - [WKUserScript.InitWithSourceInjectionTimeForMainFrameOnlyInContentWorld]: Creates a user script object that is scoped to a particular content world.
//
// # Inspecting Script Information
//
//   - [WKUserScript.Source]: The script’s source code.
//   - [WKUserScript.InjectionTime]: The time at which to inject the script into the webpage.
//   - [WKUserScript.ForMainFrameOnly]: A Boolean value that indicates whether to inject the script into the main frame or all frames.
//
// See: https://developer.apple.com/documentation/WebKit/WKUserScript
type WKUserScript struct {
	objectivec.Object
}

// WKUserScriptFromID constructs a [WKUserScript] from an objc.ID.
//
// A script that the web view injects into a webpage.
func WKUserScriptFromID(id objc.ID) WKUserScript {
	return WKUserScript{objectivec.Object{ID: id}}
}

// NOTE: WKUserScript adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [WKUserScript] class.
//
// # Creating a User Script Object
//
//   - [IWKUserScript.InitWithSourceInjectionTimeForMainFrameOnly]: Creates a user script object that contains the specified source code and attributes.
//   - [IWKUserScript.InitWithSourceInjectionTimeForMainFrameOnlyInContentWorld]: Creates a user script object that is scoped to a particular content world.
//
// # Inspecting Script Information
//
//   - [IWKUserScript.Source]: The script’s source code.
//   - [IWKUserScript.InjectionTime]: The time at which to inject the script into the webpage.
//   - [IWKUserScript.ForMainFrameOnly]: A Boolean value that indicates whether to inject the script into the main frame or all frames.
//
// See: https://developer.apple.com/documentation/WebKit/WKUserScript
type IWKUserScript interface {
	objectivec.IObject

	// Topic: Creating a User Script Object

	// Creates a user script object that contains the specified source code and attributes.
	InitWithSourceInjectionTimeForMainFrameOnly(source string, injectionTime WKUserScriptInjectionTime, forMainFrameOnly bool) WKUserScript
	// Creates a user script object that is scoped to a particular content world.
	InitWithSourceInjectionTimeForMainFrameOnlyInContentWorld(source string, injectionTime WKUserScriptInjectionTime, forMainFrameOnly bool, contentWorld IWKContentWorld) WKUserScript

	// Topic: Inspecting Script Information

	// The script’s source code.
	Source() string
	// The time at which to inject the script into the webpage.
	InjectionTime() WKUserScriptInjectionTime
	// A Boolean value that indicates whether to inject the script into the main frame or all frames.
	ForMainFrameOnly() bool
}

// Init initializes the instance.
func (u WKUserScript) Init() WKUserScript {
	rv := objc.Send[WKUserScript](u.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u WKUserScript) Autorelease() WKUserScript {
	rv := objc.Send[WKUserScript](u.ID, objc.Sel("autorelease"))
	return rv
}

// NewWKUserScript creates a new WKUserScript instance.
func NewWKUserScript() WKUserScript {
	class := getWKUserScriptClass()
	rv := objc.Send[WKUserScript](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a user script object that contains the specified source code and
// attributes.
//
// source: The script’s source code.
//
// injectionTime: The time at which to inject the script into the webpage. For a list of
// possible values, see [WKUserScriptInjectionTime].
//
// forMainFrameOnly: A Boolean value that indicates whether to inject the script into the main
// frame. Specify true to inject the script only into the main frame, or false
// to inject it into all frames.
//
// # Return Value
//
// An initialized user script, or `nil` if initialization failed.
//
// # Discussion
//
// This method sets the script’s content world to the object in the
// [PageWorld] property of [WKContentWorld].
//
// See: https://developer.apple.com/documentation/WebKit/WKUserScript/init(source:injectionTime:forMainFrameOnly:)
//
// [WKUserScriptInjectionTime]: https://developer.apple.com/documentation/WebKit/WKUserScriptInjectionTime
func NewUserScriptWithSourceInjectionTimeForMainFrameOnly(source string, injectionTime WKUserScriptInjectionTime, forMainFrameOnly bool) WKUserScript {
	instance := getWKUserScriptClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSource:injectionTime:forMainFrameOnly:"), objc.String(source), injectionTime, forMainFrameOnly)
	return WKUserScriptFromID(rv)
}

// Creates a user script object that is scoped to a particular content world.
//
// source: The script’s source code.
//
// injectionTime: The time at which to inject the script into the webpage. For a list of
// possible values, see [WKUserScriptInjectionTime].
//
// forMainFrameOnly: A Boolean value that indicates whether to inject the script into the main
// frame. Specify true to inject the script only into the main frame, or false
// to inject it into all frames.
//
// contentWorld: The namespace in which to evaluate the script. This parameter doesn’t
// apply to changes your script makes to the underlying web content, such as
// the document’s DOM structure. Those changes remain visible to all
// scripts, regardless of which content world you specify. For more
// information about content worlds, see [WKContentWorld].
//
// # Return Value
//
// An initialized user script, or `nil` if initialization failed.
//
// See: https://developer.apple.com/documentation/WebKit/WKUserScript/init(source:injectionTime:forMainFrameOnly:in:)
//
// [WKUserScriptInjectionTime]: https://developer.apple.com/documentation/WebKit/WKUserScriptInjectionTime
func NewUserScriptWithSourceInjectionTimeForMainFrameOnlyInContentWorld(source string, injectionTime WKUserScriptInjectionTime, forMainFrameOnly bool, contentWorld IWKContentWorld) WKUserScript {
	instance := getWKUserScriptClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSource:injectionTime:forMainFrameOnly:inContentWorld:"), objc.String(source), injectionTime, forMainFrameOnly, contentWorld)
	return WKUserScriptFromID(rv)
}

// Creates a user script object that contains the specified source code and
// attributes.
//
// source: The script’s source code.
//
// injectionTime: The time at which to inject the script into the webpage. For a list of
// possible values, see [WKUserScriptInjectionTime].
//
// forMainFrameOnly: A Boolean value that indicates whether to inject the script into the main
// frame. Specify true to inject the script only into the main frame, or false
// to inject it into all frames.
//
// # Return Value
//
// An initialized user script, or `nil` if initialization failed.
//
// # Discussion
//
// This method sets the script’s content world to the object in the
// [PageWorld] property of [WKContentWorld].
//
// See: https://developer.apple.com/documentation/WebKit/WKUserScript/init(source:injectionTime:forMainFrameOnly:)
//
// [WKUserScriptInjectionTime]: https://developer.apple.com/documentation/WebKit/WKUserScriptInjectionTime
func (u WKUserScript) InitWithSourceInjectionTimeForMainFrameOnly(source string, injectionTime WKUserScriptInjectionTime, forMainFrameOnly bool) WKUserScript {
	rv := objc.Send[WKUserScript](u.ID, objc.Sel("initWithSource:injectionTime:forMainFrameOnly:"), objc.String(source), injectionTime, forMainFrameOnly)
	return rv
}

// Creates a user script object that is scoped to a particular content world.
//
// source: The script’s source code.
//
// injectionTime: The time at which to inject the script into the webpage. For a list of
// possible values, see [WKUserScriptInjectionTime].
//
// forMainFrameOnly: A Boolean value that indicates whether to inject the script into the main
// frame. Specify true to inject the script only into the main frame, or false
// to inject it into all frames.
//
// contentWorld: The namespace in which to evaluate the script. This parameter doesn’t
// apply to changes your script makes to the underlying web content, such as
// the document’s DOM structure. Those changes remain visible to all
// scripts, regardless of which content world you specify. For more
// information about content worlds, see [WKContentWorld].
//
// # Return Value
//
// An initialized user script, or `nil` if initialization failed.
//
// See: https://developer.apple.com/documentation/WebKit/WKUserScript/init(source:injectionTime:forMainFrameOnly:in:)
//
// [WKUserScriptInjectionTime]: https://developer.apple.com/documentation/WebKit/WKUserScriptInjectionTime
func (u WKUserScript) InitWithSourceInjectionTimeForMainFrameOnlyInContentWorld(source string, injectionTime WKUserScriptInjectionTime, forMainFrameOnly bool, contentWorld IWKContentWorld) WKUserScript {
	rv := objc.Send[WKUserScript](u.ID, objc.Sel("initWithSource:injectionTime:forMainFrameOnly:inContentWorld:"), objc.String(source), injectionTime, forMainFrameOnly, contentWorld)
	return rv
}

// The script’s source code.
//
// See: https://developer.apple.com/documentation/WebKit/WKUserScript/source
func (u WKUserScript) Source() string {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("source"))
	return foundation.NSStringFromID(rv).String()
}

// The time at which to inject the script into the webpage.
//
// # Discussion
//
// The value is one of the constants of the enumerated type
// [WKUserScriptInjectionTime].
//
// See: https://developer.apple.com/documentation/WebKit/WKUserScript/injectionTime
//
// [WKUserScriptInjectionTime]: https://developer.apple.com/documentation/WebKit/WKUserScriptInjectionTime
func (u WKUserScript) InjectionTime() WKUserScriptInjectionTime {
	rv := objc.Send[WKUserScriptInjectionTime](u.ID, objc.Sel("injectionTime"))
	return WKUserScriptInjectionTime(rv)
}

// A Boolean value that indicates whether to inject the script into the main
// frame or all frames.
//
// # Discussion
//
// When the value of this property is true, the web view injects the script
// only into the main frame. When the value is false, the web view injects the
// script into all frames.
//
// See: https://developer.apple.com/documentation/WebKit/WKUserScript/isForMainFrameOnly
func (u WKUserScript) ForMainFrameOnly() bool {
	rv := objc.Send[bool](u.ID, objc.Sel("isForMainFrameOnly"))
	return rv
}
