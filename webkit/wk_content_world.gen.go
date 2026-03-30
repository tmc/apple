// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [WKContentWorld] class.
var (
	_WKContentWorldClass     WKContentWorldClass
	_WKContentWorldClassOnce sync.Once
)

func getWKContentWorldClass() WKContentWorldClass {
	_WKContentWorldClassOnce.Do(func() {
		_WKContentWorldClass = WKContentWorldClass{class: objc.GetClass("WKContentWorld")}
	})
	return _WKContentWorldClass
}

// GetWKContentWorldClass returns the class object for WKContentWorld.
func GetWKContentWorldClass() WKContentWorldClass {
	return getWKContentWorldClass()
}

type WKContentWorldClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (wc WKContentWorldClass) Class() objc.Class {
	return wc.class
}

// Alloc allocates memory for a new instance of the class.
func (wc WKContentWorldClass) Alloc() WKContentWorld {
	rv := objc.Send[WKContentWorld](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

// An object that defines a scope of execution for JavaScript code, and which
// you use to prevent conflicts between different scripts.
//
// # Overview
//
// Use a [WKContentWorld] object as a namespace to separate your app’s web
// environment from the environment of individual webpages or scripts you
// execute. Content worlds help prevent issues that occur when two scripts
// modify environment variables in conflicting ways. Executing a script in its
// own content world effectively gives it a separate copy of the environment
// variables to modify. You might use this support in the following scenarios:
//
// - You have complex script logic to bridge your web content to your app, but
// your web content has complex script libraries of its own. In that scenario,
// use one content world for your app-specific scripts and a separate content
// world for your content-specific scripts. - You implement a web browser that
// supports JavaScript extensions. In that scenario, create a unique content
// world for each extension to prevent conflicts between the extensions.
//
// A [WKContentWorld] object is a namespace and doesn’t persist data outside
// of the current web view or webpage. If you use the same content world in
// two [WKWebView] objects, variables in one web view’s content world
// don’t appear in the other web view. Similarly, when the user or your app
// navigates to a new webpage, variables from the previous page are gone, even
// if both pages share the same content world.
//
// Use the methods and properties of this class to fetch the content world you
// need. [WKContentWorld] provides a default content world for your app and a
// content world for the current web page. You can also create new content
// worlds. For example, you might create a custom content world for each
// JavaScript extension you manage. Specify the content world object when
// configuring or executing scripts associated with your content.
//
// # Retrieving a Custom Content World
//
//   - [WKContentWorld.Name]: The name of a custom content world.
//
// See: https://developer.apple.com/documentation/WebKit/WKContentWorld
type WKContentWorld struct {
	objectivec.Object
}

// WKContentWorldFromID constructs a [WKContentWorld] from an objc.ID.
//
// An object that defines a scope of execution for JavaScript code, and which
// you use to prevent conflicts between different scripts.
func WKContentWorldFromID(id objc.ID) WKContentWorld {
	return WKContentWorld{objectivec.Object{ID: id}}
}

// NOTE: WKContentWorld adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [WKContentWorld] class.
//
// # Retrieving a Custom Content World
//
//   - [IWKContentWorld.Name]: The name of a custom content world.
//
// See: https://developer.apple.com/documentation/WebKit/WKContentWorld
type IWKContentWorld interface {
	objectivec.IObject

	// Topic: Retrieving a Custom Content World

	// The name of a custom content world.
	Name() string
}

// Init initializes the instance.
func (c WKContentWorld) Init() WKContentWorld {
	rv := objc.Send[WKContentWorld](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c WKContentWorld) Autorelease() WKContentWorld {
	rv := objc.Send[WKContentWorld](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewWKContentWorld creates a new WKContentWorld instance.
func NewWKContentWorld() WKContentWorld {
	class := getWKContentWorldClass()
	rv := objc.Send[WKContentWorld](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Returns the custom content world with the specified name.
//
// name: The name of the content world you want. If no content world with that name
// exists, this method creates a new [WKContentWorld] object and returns it.
// The next time you request a content world with the same name, this method
// returns the object it previously created.
//
// # Return Value
//
// The content world with the specified name.
//
// # Discussion
//
// Use this method to create unique content worlds for your script code. For
// example, if you execute scripts from multiple JavaScript extensions, you
// might use this method to create a content world based on a unique string
// associated with that extension.
//
// See: https://developer.apple.com/documentation/WebKit/WKContentWorld/world(name:)
func (_WKContentWorldClass WKContentWorldClass) WorldWithName(name string) WKContentWorld {
	rv := objc.Send[objc.ID](objc.ID(_WKContentWorldClass.class), objc.Sel("worldWithName:"), objc.String(name))
	return WKContentWorldFromID(rv)
}

// The name of a custom content world.
//
// # Discussion
//
// This property contains a valid string only for content worlds you retrieve
// using the [WorldWithName] function. The value of this property is `nil` for
// the content worlds in the [DefaultClientWorld] and [PageWorld] properties.
//
// See: https://developer.apple.com/documentation/WebKit/WKContentWorld/name
func (c WKContentWorld) Name() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("name"))
	return foundation.NSStringFromID(rv).String()
}

// The default world for clients.
//
// See: https://developer.apple.com/documentation/WebKit/WKContentWorld/defaultClient
func (_WKContentWorldClass WKContentWorldClass) DefaultClientWorld() WKContentWorld {
	rv := objc.Send[objc.ID](objc.ID(_WKContentWorldClass.class), objc.Sel("defaultClientWorld"))
	return WKContentWorldFromID(objc.ID(rv))
}

// The content world for the current webpage’s content.
//
// # Discussion
//
// This property contains the content world for scripts that the current
// webpage executes. Be careful when manipulating variables in this content
// world. If you modify a variable with the same name as one the webpage uses,
// you may unintentionally disrupt the normal operation of that page.
//
// See: https://developer.apple.com/documentation/WebKit/WKContentWorld/page
func (_WKContentWorldClass WKContentWorldClass) PageWorld() WKContentWorld {
	rv := objc.Send[objc.ID](objc.ID(_WKContentWorldClass.class), objc.Sel("pageWorld"))
	return WKContentWorldFromID(objc.ID(rv))
}
