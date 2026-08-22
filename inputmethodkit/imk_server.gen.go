// Code generated from Apple documentation for InputMethodKit. DO NOT EDIT.

package inputmethodkit

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [IMKServer] class.
var (
	_IMKServerClass     IMKServerClass
	_IMKServerClassOnce sync.Once
)

func getIMKServerClass() IMKServerClass {
	_IMKServerClassOnce.Do(func() {
		_IMKServerClass = IMKServerClass{class: objc.GetClass("IMKServer")}
	})
	return _IMKServerClass
}

// GetIMKServerClass returns the class object for IMKServer.
func GetIMKServerClass() IMKServerClass {
	return getIMKServerClass()
}

type IMKServerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ic IMKServerClass) Class() objc.Class {
	return ic.class
}

// Alloc allocates memory for a new instance of the class.
func (ic IMKServerClass) Alloc() IMKServer {
	rv := objc.Send[IMKServer](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// The [IMKServer] class manages client connections to your input method. When
// you write the main function for your input method, you create an
// [IMKServer] object. You should never need to override this class.
//
// # Initializing a Server Object
//
//   - [IMKServer.InitWithNameBundleIdentifier]: Creates and returns a server object from property list information contained in the provided bundle.
//   - [IMKServer.InitWithNameControllerClassDelegateClass]: Creates and returns a server object initialized with the provided parameters.
//
// # Getting a Bundle for the Input Method
//
//   - [IMKServer.Bundle]: Returns an [NSBundle] object for the input method.
//
// # Instance Methods
//
//   - [IMKServer.LastKeyEventWasDeadKey]
//   - [IMKServer.PaletteWillTerminate]
//
// See: https://developer.apple.com/documentation/InputMethodKit/IMKServer
type IMKServer struct {
	objectivec.Object
}

// IMKServerFromID constructs a [IMKServer] from an objc.ID.
//
// The [IMKServer] class manages client connections to your input method. When
// you write the main function for your input method, you create an
// [IMKServer] object. You should never need to override this class.
func IMKServerFromID(id objc.ID) IMKServer {
	return IMKServer{objectivec.Object{ID: id}}
}

// NOTE: IMKServer adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [IMKServer] class.
//
// # Initializing a Server Object
//
//   - [IIMKServer.InitWithNameBundleIdentifier]: Creates and returns a server object from property list information contained in the provided bundle.
//   - [IIMKServer.InitWithNameControllerClassDelegateClass]: Creates and returns a server object initialized with the provided parameters.
//
// # Getting a Bundle for the Input Method
//
//   - [IIMKServer.Bundle]: Returns an [NSBundle] object for the input method.
//
// # Instance Methods
//
//   - [IIMKServer.LastKeyEventWasDeadKey]
//   - [IIMKServer.PaletteWillTerminate]
//
// See: https://developer.apple.com/documentation/InputMethodKit/IMKServer
type IIMKServer interface {
	objectivec.IObject

	// Topic: Initializing a Server Object

	// Creates and returns a server object from property list information contained in the provided bundle.
	InitWithNameBundleIdentifier(name string, bundleIdentifier string) IMKServer
	// Creates and returns a server object initialized with the provided parameters.
	InitWithNameControllerClassDelegateClass(name string, controllerClassID objectivec.Class, delegateClassID objectivec.Class) IMKServer

	// Topic: Getting a Bundle for the Input Method

	// Returns an [NSBundle] object for the input method.
	Bundle() foundation.Bundle

	// Topic: Instance Methods

	LastKeyEventWasDeadKey() bool
	PaletteWillTerminate() bool
}

// Init initializes the instance.
func (i IMKServer) Init() IMKServer {
	rv := objc.Send[IMKServer](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i IMKServer) Autorelease() IMKServer {
	rv := objc.Send[IMKServer](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewIMKServer creates a new IMKServer instance.
func NewIMKServer() IMKServer {
	class := getIMKServerClass()
	rv := objc.Send[IMKServer](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates and returns a server object from property list information
// contained in the provided bundle.
//
// name: The name to initialize the server object with.
//
// bundleIdentifier: The bundle identifier.
//
// # Return Value
//
// An initialized server object.
//
// # Discussion
//
// This method examines the `Info.Plist()` file for the entries shown in the
// table below. The class names are loaded, but no classes are instantiated.
// Additionally, an [NSConnection] object is allocated and registered using
// the input method connection name supplied in the `Info.Plist()` file.
//
// [Table data omitted]
//
// See: https://developer.apple.com/documentation/InputMethodKit/IMKServer/init(name:bundleIdentifier:)
//
// [NSConnection]: https://developer.apple.com/documentation/Foundation/NSConnection
func NewIMKServerWithNameBundleIdentifier(name string, bundleIdentifier string) IMKServer {
	instance := getIMKServerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithName:bundleIdentifier:"), objc.String(name), objc.String(bundleIdentifier))
	return IMKServerFromID(rv)
}

// Creates and returns a server object initialized with the provided
// parameters.
//
// name: The name to initialize the server object with.
//
// controllerClassID: The id for the input controller class.
//
// delegateClassID: The id for the delegate class.
//
// # Return Value
//
// An initialized server object.
//
// See: https://developer.apple.com/documentation/InputMethodKit/IMKServer/init(name:controllerClass:delegateClass:)
func NewIMKServerWithNameControllerClassDelegateClass(name string, controllerClassID objectivec.Class, delegateClassID objectivec.Class) IMKServer {
	instance := getIMKServerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithName:controllerClass:delegateClass:"), objc.String(name), controllerClassID, delegateClassID)
	return IMKServerFromID(rv)
}

// Creates and returns a server object from property list information
// contained in the provided bundle.
//
// name: The name to initialize the server object with.
//
// bundleIdentifier: The bundle identifier.
//
// # Return Value
//
// An initialized server object.
//
// # Discussion
//
// This method examines the `Info.Plist()` file for the entries shown in the
// table below. The class names are loaded, but no classes are instantiated.
// Additionally, an [NSConnection] object is allocated and registered using
// the input method connection name supplied in the `Info.Plist()` file.
//
// [Table data omitted]
//
// See: https://developer.apple.com/documentation/InputMethodKit/IMKServer/init(name:bundleIdentifier:)
//
// [NSConnection]: https://developer.apple.com/documentation/Foundation/NSConnection
func (i IMKServer) InitWithNameBundleIdentifier(name string, bundleIdentifier string) IMKServer {
	rv := objc.Send[IMKServer](i.ID, objc.Sel("initWithName:bundleIdentifier:"), objc.String(name), objc.String(bundleIdentifier))
	return rv
}

// Creates and returns a server object initialized with the provided
// parameters.
//
// name: The name to initialize the server object with.
//
// controllerClassID: The id for the input controller class.
//
// delegateClassID: The id for the delegate class.
//
// # Return Value
//
// An initialized server object.
//
// See: https://developer.apple.com/documentation/InputMethodKit/IMKServer/init(name:controllerClass:delegateClass:)
func (i IMKServer) InitWithNameControllerClassDelegateClass(name string, controllerClassID objectivec.Class, delegateClassID objectivec.Class) IMKServer {
	rv := objc.Send[IMKServer](i.ID, objc.Sel("initWithName:controllerClass:delegateClass:"), objc.String(name), controllerClassID, delegateClassID)
	return rv
}

// Returns an [NSBundle] object for the input method.
//
// # Return Value
//
// An [NSBundle] object that is either created from the bundle identifier
// contained in the server object, or from the main bundle.
//
// See: https://developer.apple.com/documentation/InputMethodKit/IMKServer/bundle()
func (i IMKServer) Bundle() foundation.Bundle {
	rv := objc.Send[objc.ID](i.ID, objc.Sel("bundle"))
	return foundation.NSBundleFromID(rv)
}

// See: https://developer.apple.com/documentation/InputMethodKit/IMKServer/lastKeyEventWasDeadKey()
func (i IMKServer) LastKeyEventWasDeadKey() bool {
	rv := objc.Send[bool](i.ID, objc.Sel("lastKeyEventWasDeadKey"))
	return rv
}

// See: https://developer.apple.com/documentation/InputMethodKit/IMKServer/paletteWillTerminate()
func (i IMKServer) PaletteWillTerminate() bool {
	rv := objc.Send[bool](i.ID, objc.Sel("paletteWillTerminate"))
	return rv
}
