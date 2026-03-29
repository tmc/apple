// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSStoryboard] class.
var (
	_NSStoryboardClass     NSStoryboardClass
	_NSStoryboardClassOnce sync.Once
)

func getNSStoryboardClass() NSStoryboardClass {
	_NSStoryboardClassOnce.Do(func() {
		_NSStoryboardClass = NSStoryboardClass{class: objc.GetClass("NSStoryboard")}
	})
	return _NSStoryboardClass
}

// GetNSStoryboardClass returns the class object for NSStoryboard.
func GetNSStoryboardClass() NSStoryboardClass {
	return getNSStoryboardClass()
}

type NSStoryboardClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSStoryboardClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSStoryboardClass) Alloc() NSStoryboard {
	rv := objc.Send[NSStoryboard](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// An encapsulation of the design-time view controller and window controller
// graph represented in an Interface Builder storyboard resource file.
//
// # Overview
// 
// You can use storyboard files to define the view and window controllers for
// all or part of an app’s user interface. Typically, AppKit creates these
// objects automatically in response to actions defined within a storyboard
// file itself, such as the clicking of a button or the choosing of a menu
// item. However, you can use a storyboard object to directly instantiate the
// initial view controller from a storyboard file or to instantiate other view
// or window controllers that you want to present programmatically. In the
// context of a storyboard file, each contained controller is called a .
// 
// A transition from one scene to another in a storyboard is called a . This
// same term, and the same Cocoa APIs, express a containment relationship
// between two scenes. In macOS, containment (rather than transition) is the
// more common notion for storyboards. For descriptions of the related APIs,
// refer to [NSStoryboardSegue] and [NSSeguePerforming].
//
// # Loading the Initial View Controller
//
//   - [NSStoryboard.InstantiateInitialController]: Creates the initial view controller or window controller from a storyboard.
//
// # Instantiating Storyboard Controllers
//
//   - [NSStoryboard.InstantiateControllerWithIdentifier]: Instantiates a specified view controller or window controller from a storyboard.
//
// See: https://developer.apple.com/documentation/AppKit/NSStoryboard
type NSStoryboard struct {
	objectivec.Object
}

// NSStoryboardFromID constructs a [NSStoryboard] from an objc.ID.
//
// An encapsulation of the design-time view controller and window controller
// graph represented in an Interface Builder storyboard resource file.
func NSStoryboardFromID(id objc.ID) NSStoryboard {
	return NSStoryboard{objectivec.Object{ID: id}}
}
// NOTE: NSStoryboard adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSStoryboard] class.
//
// # Loading the Initial View Controller
//
//   - [INSStoryboard.InstantiateInitialController]: Creates the initial view controller or window controller from a storyboard.
//
// # Instantiating Storyboard Controllers
//
//   - [INSStoryboard.InstantiateControllerWithIdentifier]: Instantiates a specified view controller or window controller from a storyboard.
//
// See: https://developer.apple.com/documentation/AppKit/NSStoryboard
type INSStoryboard interface {
	objectivec.IObject

	// Topic: Loading the Initial View Controller

	// Creates the initial view controller or window controller from a storyboard.
	InstantiateInitialController() objectivec.IObject

	// Topic: Instantiating Storyboard Controllers

	// Instantiates a specified view controller or window controller from a storyboard.
	InstantiateControllerWithIdentifier(identifier NSStoryboardSceneIdentifier) objectivec.IObject
}

// Init initializes the instance.
func (s NSStoryboard) Init() NSStoryboard {
	rv := objc.Send[NSStoryboard](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s NSStoryboard) Autorelease() NSStoryboard {
	rv := objc.Send[NSStoryboard](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSStoryboard creates a new NSStoryboard instance.
func NewNSStoryboard() NSStoryboard {
	class := getNSStoryboardClass()
	rv := objc.Send[NSStoryboard](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a storyboard based on the named storyboard file in the specified
// bundle.
//
// name: The name of the storyboard file, without the filename extension. This
// method raises an exception if this parameter’s value is `nil`.
//
// storyboardBundleOrNil: The bundle used to resolve references to resources, typically images, in
// the archived controllers represented in the storyboard file. If you specify
// `nil`, AppKit uses the app’s main bundle.
//
// # Return Value
// 
// A new storyboard object.
//
// See: https://developer.apple.com/documentation/AppKit/NSStoryboard/init(name:bundle:)
func NewStoryboardWithNameBundle(name NSStoryboardName, storyboardBundleOrNil foundation.NSBundle) NSStoryboard {
	rv := objc.Send[objc.ID](objc.ID(getNSStoryboardClass().class), objc.Sel("storyboardWithName:bundle:"), objc.String(string(name)), storyboardBundleOrNil)
	return NSStoryboardFromID(rv)
}

// Creates the initial view controller or window controller from a storyboard.
//
// # Return Value
// 
// The initial view controller or window controller for the storyboard.
//
// # Discussion
// 
// Every storyboard has an initial view controller or window controller that
// represents its starting point. For the main storyboard, this is usually the
// first controller presented to the user at launch time. Designate the
// initial view controller in Interface Builder when configuring the
// storyboard file.
// 
// Typically, you call this method only when transitioning to the initial view
// controller in a different storyboard file. For your app’s main storyboard
// file—that is, the storyboard file specified in the app’s `Info.Plist()`
// file using the [UIMainStoryboardFile] key—the initial view controller is
// loaded into memory and presented automatically.
// 
// Each time you call this method, it creates a new instance of the initial
// controller.
//
// See: https://developer.apple.com/documentation/AppKit/NSStoryboard/instantiateInitialController()
func (s NSStoryboard) InstantiateInitialController() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("instantiateInitialController"))
	return objectivec.Object{ID: rv}
}
// Instantiates a specified view controller or window controller from a
// storyboard.
//
// identifier: The unique identifier for the controller, which you have specified using
// the Identity inspector in Interface Builder.
//
// # Return Value
// 
// The instantiated view controller or window controller identified by the
// `identifier` parameter, from the storyboard file. If the specified
// identifier does not exist (or is `nil`) in the storyboard file, this method
// raises an exception.
//
// # Discussion
// 
// Use this method to create a view controller or window controller object
// that you want to manipulate and present programmatically. Your controller
// object must have an identifier string. In Xcode’s storyboard editor,
// select your controller, display the identity inspector, and place this
// string in the Storyboard ID field.
// 
// This method creates a new instance of the specified controller each time
// you call it.
//
// See: https://developer.apple.com/documentation/AppKit/NSStoryboard/instantiateController(withIdentifier:)
func (s NSStoryboard) InstantiateControllerWithIdentifier(identifier NSStoryboardSceneIdentifier) objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("instantiateControllerWithIdentifier:"), objc.String(string(identifier)))
	return objectivec.Object{ID: rv}
}

// The app’s main storyboard.
//
// # Discussion
// 
// The name of the main storyboard is stored in the [NSMainStoryboardFile] key
// of the app’s `Info.Plist()` file.
//
// [NSMainStoryboardFile]: https://developer.apple.com/documentation/BundleResources/Information-Property-List/NSMainStoryboardFile
//
// See: https://developer.apple.com/documentation/AppKit/NSStoryboard/main
func (_NSStoryboardClass NSStoryboardClass) MainStoryboard() NSStoryboard {
	rv := objc.Send[objc.ID](objc.ID(_NSStoryboardClass.class), objc.Sel("mainStoryboard"))
	return NSStoryboardFromID(objc.ID(rv))
}

