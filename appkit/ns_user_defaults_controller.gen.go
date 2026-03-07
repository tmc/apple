// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSUserDefaultsController] class.
var (
	_NSUserDefaultsControllerClass     NSUserDefaultsControllerClass
	_NSUserDefaultsControllerClassOnce sync.Once
)

func getNSUserDefaultsControllerClass() NSUserDefaultsControllerClass {
	_NSUserDefaultsControllerClassOnce.Do(func() {
		_NSUserDefaultsControllerClass = NSUserDefaultsControllerClass{class: objc.GetClass("NSUserDefaultsController")}
	})
	return _NSUserDefaultsControllerClass
}

// GetNSUserDefaultsControllerClass returns the class object for NSUserDefaultsController.
func GetNSUserDefaultsControllerClass() NSUserDefaultsControllerClass {
	return getNSUserDefaultsControllerClass()
}

type NSUserDefaultsControllerClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSUserDefaultsControllerClass) Alloc() NSUserDefaultsController {
	rv := objc.Send[NSUserDefaultsController](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// A controller that accesses user preference information for your app from
// the user’s defaults database.
//
// # Overview
// 
// [NSUserDefaultsController] is a Cocoa bindings–compatible controller
// class. Properties of the shared instance of this class can be bound to user
// interface elements to access and modify values stored in [UserDefaults].
//
// [UserDefaults]: https://developer.apple.com/documentation/Foundation/UserDefaults
//
// # Initializing a user defaults controller
//
//   - [NSUserDefaultsController.InitWithDefaultsInitialValues]: Returns an initialized NSUserDefaultsController object using the NSUserDefaults instance specified in `defaults` and the initial default values contained in the `initialValues` dictionary.
//
// # Managing user defaults values
//
//   - [NSUserDefaultsController.Defaults]: Returns the instance of NSUserDefaults in use by the receiver.
//   - [NSUserDefaultsController.InitialValues]: Returns a dictionary containing the receiver’s initial default values.
//   - [NSUserDefaultsController.SetInitialValues]
//   - [NSUserDefaultsController.HasUnappliedChanges]: Returns whether the receiver has user default values that have not been saved to NSUserDefaults.
//   - [NSUserDefaultsController.AppliesImmediately]: Returns whether any changes made to bound user default properties are saved immediately.
//   - [NSUserDefaultsController.SetAppliesImmediately]
//   - [NSUserDefaultsController.Values]: Returns a key value coding compliant object that is used to access the user default properties.
//   - [NSUserDefaultsController.Revert]: Causes the receiver to discard any unsaved changes to bound user default properties, restoring their previous values.
//   - [NSUserDefaultsController.RevertToInitialValues]: Causes the receiver to discard all edits and replace the values of all the user default properties with any corresponding values in the [initialValues](<doc://com.apple.appkit/documentation/AppKit/NSUserDefaultsController/initialValues>) dictionary.
//   - [NSUserDefaultsController.Save]: Saves the values of the receiver’s user default properties.
//
// See: https://developer.apple.com/documentation/AppKit/NSUserDefaultsController
type NSUserDefaultsController struct {
	NSController
}

// NSUserDefaultsControllerFromID constructs a [NSUserDefaultsController] from an objc.ID.
//
// A controller that accesses user preference information for your app from
// the user’s defaults database.
func NSUserDefaultsControllerFromID(id objc.ID) NSUserDefaultsController {
	return NSUserDefaultsController{NSController: NSControllerFromID(id)}
}
// NOTE: NSUserDefaultsController adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [NSUserDefaultsController] class.
//
// # Initializing a user defaults controller
//
//   - [INSUserDefaultsController.InitWithDefaultsInitialValues]: Returns an initialized NSUserDefaultsController object using the NSUserDefaults instance specified in `defaults` and the initial default values contained in the `initialValues` dictionary.
//
// # Managing user defaults values
//
//   - [INSUserDefaultsController.Defaults]: Returns the instance of NSUserDefaults in use by the receiver.
//   - [INSUserDefaultsController.InitialValues]: Returns a dictionary containing the receiver’s initial default values.
//   - [INSUserDefaultsController.SetInitialValues]
//   - [INSUserDefaultsController.HasUnappliedChanges]: Returns whether the receiver has user default values that have not been saved to NSUserDefaults.
//   - [INSUserDefaultsController.AppliesImmediately]: Returns whether any changes made to bound user default properties are saved immediately.
//   - [INSUserDefaultsController.SetAppliesImmediately]
//   - [INSUserDefaultsController.Values]: Returns a key value coding compliant object that is used to access the user default properties.
//   - [INSUserDefaultsController.Revert]: Causes the receiver to discard any unsaved changes to bound user default properties, restoring their previous values.
//   - [INSUserDefaultsController.RevertToInitialValues]: Causes the receiver to discard all edits and replace the values of all the user default properties with any corresponding values in the [initialValues](<doc://com.apple.appkit/documentation/AppKit/NSUserDefaultsController/initialValues>) dictionary.
//   - [INSUserDefaultsController.Save]: Saves the values of the receiver’s user default properties.
//
// See: https://developer.apple.com/documentation/AppKit/NSUserDefaultsController
type INSUserDefaultsController interface {
	INSController

	// Topic: Initializing a user defaults controller

	// Returns an initialized NSUserDefaultsController object using the NSUserDefaults instance specified in `defaults` and the initial default values contained in the `initialValues` dictionary.
	InitWithDefaultsInitialValues(defaults *foundation.NSUserDefaults, initialValues foundation.INSDictionary) NSUserDefaultsController

	// Topic: Managing user defaults values

	// Returns the instance of NSUserDefaults in use by the receiver.
	Defaults() *foundation.NSUserDefaults
	// Returns a dictionary containing the receiver’s initial default values.
	InitialValues() foundation.INSDictionary
	SetInitialValues(value foundation.INSDictionary)
	// Returns whether the receiver has user default values that have not been saved to NSUserDefaults.
	HasUnappliedChanges() bool
	// Returns whether any changes made to bound user default properties are saved immediately.
	AppliesImmediately() bool
	SetAppliesImmediately(value bool)
	// Returns a key value coding compliant object that is used to access the user default properties.
	Values() objectivec.IObject
	// Causes the receiver to discard any unsaved changes to bound user default properties, restoring their previous values.
	Revert(sender objectivec.IObject)
	// Causes the receiver to discard all edits and replace the values of all the user default properties with any corresponding values in the [initialValues](<doc://com.apple.appkit/documentation/AppKit/NSUserDefaultsController/initialValues>) dictionary.
	RevertToInitialValues(sender objectivec.IObject)
	// Saves the values of the receiver’s user default properties.
	Save(sender objectivec.IObject)

	EncodeWithCoder(coder foundation.INSCoder)
}





// Init initializes the instance.
func (u NSUserDefaultsController) Init() NSUserDefaultsController {
	rv := objc.Send[NSUserDefaultsController](u.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u NSUserDefaultsController) Autorelease() NSUserDefaultsController {
	rv := objc.Send[NSUserDefaultsController](u.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSUserDefaultsController creates a new NSUserDefaultsController instance.
func NewNSUserDefaultsController() NSUserDefaultsController {
	class := getNSUserDefaultsControllerClass()
	rv := objc.Send[NSUserDefaultsController](objc.ID(class.class), objc.Sel("new"))
	return rv
}






//
// See: https://developer.apple.com/documentation/AppKit/NSUserDefaultsController/init(coder:)
func NewUserDefaultsControllerWithCoder(coder foundation.INSCoder) NSUserDefaultsController {
	instance := getNSUserDefaultsControllerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSUserDefaultsControllerFromID(rv)
}


// Returns an initialized NSUserDefaultsController object using the
// NSUserDefaults instance specified in `defaults` and the initial default
// values contained in the `initialValues` dictionary.
//
// # Discussion
// 
// If `defaults` is `nil`, the receiver uses `[NSUserDefaults
// standardUserDefaults]`.
// 
// This method is the designated initializer.
//
// See: https://developer.apple.com/documentation/AppKit/NSUserDefaultsController/init(defaults:initialValues:)
func NewUserDefaultsControllerWithDefaultsInitialValues(defaults *foundation.NSUserDefaults, initialValues foundation.INSDictionary) NSUserDefaultsController {
	instance := getNSUserDefaultsControllerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDefaults:initialValues:"), defaults, initialValues)
	return NSUserDefaultsControllerFromID(rv)
}







// Returns an initialized NSUserDefaultsController object using the
// NSUserDefaults instance specified in `defaults` and the initial default
// values contained in the `initialValues` dictionary.
//
// # Discussion
// 
// If `defaults` is `nil`, the receiver uses `[NSUserDefaults
// standardUserDefaults]`.
// 
// This method is the designated initializer.
//
// See: https://developer.apple.com/documentation/AppKit/NSUserDefaultsController/init(defaults:initialValues:)
func (u NSUserDefaultsController) InitWithDefaultsInitialValues(defaults *foundation.NSUserDefaults, initialValues foundation.INSDictionary) NSUserDefaultsController {
	rv := objc.Send[NSUserDefaultsController](u.ID, objc.Sel("initWithDefaults:initialValues:"), defaults, initialValues)
	return rv
}

// Causes the receiver to discard any unsaved changes to bound user default
// properties, restoring their previous values.
//
// # Discussion
// 
// The receiver invokes [discardEditing] on any currently registered editors.
// The `sender` is typically the object that invoked this method.
// 
// If [AppliesImmediately] is [true], this method only causes any bound
// editors with uncommitted changes to discard their edits.
//
// [discardEditing]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/discardEditing
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSUserDefaultsController/revert(_:)
func (u NSUserDefaultsController) Revert(sender objectivec.IObject) {
	objc.Send[objc.ID](u.ID, objc.Sel("revert:"), sender)
}

// Causes the receiver to discard all edits and replace the values of all the
// user default properties with any corresponding values in the
// [InitialValues] dictionary.
//
// # Discussion
// 
// This effectively sets the preferences that a user can change to their
// “out-of-the-box” values. This method has no effect if initial values
// were not specified. The `sender` is typically the object that invoked this
// method.
//
// See: https://developer.apple.com/documentation/AppKit/NSUserDefaultsController/revertToInitialValues(_:)
func (u NSUserDefaultsController) RevertToInitialValues(sender objectivec.IObject) {
	objc.Send[objc.ID](u.ID, objc.Sel("revertToInitialValues:"), sender)
}

// Saves the values of the receiver’s user default properties.
//
// # Discussion
// 
// This method has no effect if [AppliesImmediately] returns [true].
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSUserDefaultsController/save(_:)
func (u NSUserDefaultsController) Save(sender objectivec.IObject) {
	objc.Send[objc.ID](u.ID, objc.Sel("save:"), sender)
}
func (u NSUserDefaultsController) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](u.ID, objc.Sel("encodeWithCoder:"), coder)
}












// Returns the instance of NSUserDefaults in use by the receiver.
//
// # Discussion
// 
// This property is observable using key-value observing.
//
// See: https://developer.apple.com/documentation/AppKit/NSUserDefaultsController/defaults
func (u NSUserDefaultsController) Defaults() *foundation.NSUserDefaults {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("defaults"))
	if rv == 0 {
		return nil
	}
	val := foundation.NSUserDefaultsFromID(objc.ID(rv))
	return &val
}



// Returns a dictionary containing the receiver’s initial default values.
//
// # Discussion
// 
// These values are used when is no value found for the bound property in
// [Defaults].
// 
// This property is observable using key-value observing.
//
// See: https://developer.apple.com/documentation/AppKit/NSUserDefaultsController/initialValues
func (u NSUserDefaultsController) InitialValues() foundation.INSDictionary {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("initialValues"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (u NSUserDefaultsController) SetInitialValues(value foundation.INSDictionary) {
	objc.Send[struct{}](u.ID, objc.Sel("setInitialValues:"), value)
}



// Returns whether the receiver has user default values that have not been
// saved to NSUserDefaults.
//
// # Discussion
// 
// This property is observable using key-value observing.
//
// See: https://developer.apple.com/documentation/AppKit/NSUserDefaultsController/hasUnappliedChanges
func (u NSUserDefaultsController) HasUnappliedChanges() bool {
	rv := objc.Send[bool](u.ID, objc.Sel("hasUnappliedChanges"))
	return rv
}



// Returns whether any changes made to bound user default properties are saved
// immediately.
//
// # Discussion
// 
// Default is [true].
// 
// This property is observable using key-value observing.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSUserDefaultsController/appliesImmediately
func (u NSUserDefaultsController) AppliesImmediately() bool {
	rv := objc.Send[bool](u.ID, objc.Sel("appliesImmediately"))
	return rv
}
func (u NSUserDefaultsController) SetAppliesImmediately(value bool) {
	objc.Send[struct{}](u.ID, objc.Sel("setAppliesImmediately:"), value)
}



// Returns a key value coding compliant object that is used to access the user
// default properties.
//
// # Discussion
// 
// If present the value for the property in [Defaults] is returned, otherwise
// a corresponding value in [InitialValues] is returned.
// 
// This property is observable using key-value observing.
//
// See: https://developer.apple.com/documentation/AppKit/NSUserDefaultsController/values
func (u NSUserDefaultsController) Values() objectivec.IObject {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("values"))
	return objectivec.Object{ID: rv}
}







// Returns the shared instance of NSUserDefaultsController, creating it if
// necessary.
//
// # Discussion
// 
// This instance has no initial values, and uses `[NSUserDefaults
// standardUserDefaults]` to create the defaults. An application can get this
// object when an application launches and configure it as required.
//
// See: https://developer.apple.com/documentation/AppKit/NSUserDefaultsController/shared
func (_NSUserDefaultsControllerClass NSUserDefaultsControllerClass) SharedUserDefaultsController() NSUserDefaultsController {
	rv := objc.Send[objc.ID](objc.ID(_NSUserDefaultsControllerClass.class), objc.Sel("sharedUserDefaultsController"))
	return NSUserDefaultsControllerFromID(objc.ID(rv))
}
























