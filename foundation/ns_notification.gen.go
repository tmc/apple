// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSNotification] class.
var (
	_NSNotificationClass     NSNotificationClass
	_NSNotificationClassOnce sync.Once
)

func getNSNotificationClass() NSNotificationClass {
	_NSNotificationClassOnce.Do(func() {
		_NSNotificationClass = NSNotificationClass{class: objc.GetClass("NSNotification")}
	})
	return _NSNotificationClass
}

// GetNSNotificationClass returns the class object for NSNotification.
func GetNSNotificationClass() NSNotificationClass {
	return getNSNotificationClass()
}

type NSNotificationClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSNotificationClass) Alloc() NSNotification {
	rv := objc.Send[NSNotification](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A container for information broadcast through a notification center to all
// registered observers.
//
// # Overview
// 
// In Swift, this object bridges to [Notification]; use [NSNotification] when
// you need reference semantics or other Foundation-specific behavior.
// 
// A notification contains a name, an object, and an optional dictionary, and
// is broadcast to by instances of [NSNotificationCenter] or
// [NSDistributedNotificationCenter]. The name is a tag identifying the
// notification. The object is any object that the poster of the notification
// wants to send to observers of that notification (typically, the object
// posting the notification). The dictionary stores other related objects, if
// any. [NSNotification] objects are immutable.
// 
// You don’t usually create your own notifications directly, but instead
// call the [NSNotificationCenter] methods [PostNotificationNameObject] and
// [PostNotificationNameObjectUserInfo].
// 
// # Object Comparison
// 
// The objects of a notification are compared using pointer equality for local
// notifications. Distributed notifications use strings as their objects, and
// those strings are compared using [isEqual(_:)], because pointer equality
// doesn’t make sense across process boundaries.
// 
// # Creating Subclasses
// 
// You can subclass [NSNotification] to contain information in addition to the
// notification name, object, and dictionary. This extra data must be agreed
// upon between notifiers and observers.
// 
// [NSNotificationCenter] is a class cluster with no instance variables. As
// such, you must subclass [NSNotification] and override the primitive methods
// [Name], [NSNotification.GetObject], and [UserInfo]. You can choose any designated
// initializer you like, but be sure that your initializer does not call
// [NSNotification.Init] on `super` ([NSNotification] is not meant to be instantiated
// directly, and its `init` method raises an exception).
//
// [Notification]: https://developer.apple.com/documentation/Foundation/Notification
// [isEqual(_:)]: https://developer.apple.com/documentation/ObjectiveC/NSObjectProtocol/isEqual(_:)
//
// # Creating Notifications
//
//   - [NSNotification.InitWithNameObjectUserInfo]: Initializes a notification with a specified name, object, and user information.
//
// # Getting Notification Information
//
//   - [NSNotification.Name]: The name of the notification.
//   - [NSNotification.GetObject]: The object associated with the notification.
//   - [NSNotification.UserInfo]: The user information dictionary associated with the notification.
//
// See: https://developer.apple.com/documentation/Foundation/NSNotification
type NSNotification struct {
	objectivec.Object
}

// NSNotificationFromID constructs a [NSNotification] from an objc.ID.
//
// A container for information broadcast through a notification center to all
// registered observers.
func NSNotificationFromID(id objc.ID) NSNotification {
	return NSNotification{objectivec.Object{ID: id}}
}
// NOTE: NSNotification adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSNotification] class.
//
// # Creating Notifications
//
//   - [INSNotification.InitWithNameObjectUserInfo]: Initializes a notification with a specified name, object, and user information.
//
// # Getting Notification Information
//
//   - [INSNotification.Name]: The name of the notification.
//   - [INSNotification.GetObject]: The object associated with the notification.
//   - [INSNotification.UserInfo]: The user information dictionary associated with the notification.
//
// See: https://developer.apple.com/documentation/Foundation/NSNotification
type INSNotification interface {
	objectivec.IObject
	NSCoding
	NSCopying

	// Topic: Creating Notifications

	// Initializes a notification with a specified name, object, and user information.
	InitWithNameObjectUserInfo(name NSNotificationName, object objectivec.IObject, userInfo INSDictionary) NSNotification

	// Topic: Getting Notification Information

	// The name of the notification.
	Name() NSNotificationName
	// The object associated with the notification.
	GetObject() objectivec.IObject
	// The user information dictionary associated with the notification.
	UserInfo() INSDictionary
}

// Init initializes the instance.
func (n NSNotification) Init() NSNotification {
	rv := objc.Send[NSNotification](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n NSNotification) Autorelease() NSNotification {
	rv := objc.Send[NSNotification](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSNotification creates a new NSNotification instance.
func NewNSNotification() NSNotification {
	class := getNSNotificationClass()
	rv := objc.Send[NSNotification](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Initializes a notification with the data from an unarchiver.
//
// See: https://developer.apple.com/documentation/Foundation/NSNotification/init(coder:)
func NewNotificationWithCoder(coder INSCoder) NSNotification {
	instance := getNSNotificationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSNotificationFromID(rv)
}

// Returns a new notification object with a specified name and object.
//
// aName: The name for the new notification. May not be `nil`.
//
// anObject: The object for the new notification.
//
// See: https://developer.apple.com/documentation/Foundation/NSNotification/init(name:object:)
func NewNotificationWithNameObject(aName NSNotificationName, anObject objectivec.IObject) NSNotification {
	rv := objc.Send[objc.ID](objc.ID(getNSNotificationClass().class), objc.Sel("notificationWithName:object:"), objc.String(string(aName)), anObject)
	return NSNotificationFromID(rv)
}

// Initializes a notification with a specified name, object, and user
// information.
//
// name: The name for the new notification. May not be `nil`.
//
// object: The object for the new notification.
//
// userInfo: The user information dictionary for the new notification. May be `nil`.
//
// See: https://developer.apple.com/documentation/Foundation/NSNotification/init(name:object:userInfo:)
func NewNotificationWithNameObjectUserInfo(name NSNotificationName, object objectivec.IObject, userInfo INSDictionary) NSNotification {
	instance := getNSNotificationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithName:object:userInfo:"), objc.String(string(name)), object, userInfo)
	return NSNotificationFromID(rv)
}

// Initializes a notification with the data from an unarchiver.
//
// See: https://developer.apple.com/documentation/Foundation/NSNotification/init(coder:)
func (n NSNotification) InitWithCoder(coder INSCoder) NSNotification {
	rv := objc.Send[NSNotification](n.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}

// Initializes a notification with a specified name, object, and user
// information.
//
// name: The name for the new notification. May not be `nil`.
//
// object: The object for the new notification.
//
// userInfo: The user information dictionary for the new notification. May be `nil`.
//
// See: https://developer.apple.com/documentation/Foundation/NSNotification/init(name:object:userInfo:)
func (n NSNotification) InitWithNameObjectUserInfo(name NSNotificationName, object objectivec.IObject, userInfo INSDictionary) NSNotification {
	rv := objc.Send[NSNotification](n.ID, objc.Sel("initWithName:object:userInfo:"), objc.String(string(name)), object, userInfo)
	return rv
}

// Encodes the receiver using a given archiver.
//
// coder: An archiver object.
//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/encode(with:)
func (n NSNotification) EncodeWithCoder(coder INSCoder) {
	objc.Send[objc.ID](n.ID, objc.Sel("encodeWithCoder:"), coder)
}

// Returns a notification object with a specified name, object, and user
// information.
//
// aName: The name for the new notification. May not be `nil`.
//
// anObject: The object for the new notification.
//
// aUserInfo: The user information dictionary for the new notification. May be `nil`.
//
// See: https://developer.apple.com/documentation/Foundation/NSNotification/notificationWithName:object:userInfo:
func (_NSNotificationClass NSNotificationClass) NotificationWithNameObjectUserInfo(aName NSNotificationName, anObject objectivec.IObject, aUserInfo INSDictionary) NSNotification {
	rv := objc.Send[objc.ID](objc.ID(_NSNotificationClass.class), objc.Sel("notificationWithName:object:userInfo:"), objc.String(string(aName)), anObject, aUserInfo)
	return NSNotificationFromID(rv)
}

// The name of the notification.
//
// # Discussion
// 
// Typically you use this property to find out what kind of notification you
// are dealing with when you receive a notification.
// 
// # Special Considerations
// 
// Notification names can be any string. To avoid name collisions, you might
// want to use a prefix that’s specific to your application.
//
// See: https://developer.apple.com/documentation/Foundation/NSNotification/name-swift.property
func (n NSNotification) Name() NSNotificationName {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("name"))
	return NSNotificationName(NSStringFromID(rv).String())
}

// The object associated with the notification.
//
// # Discussion
// 
// This is often the object that posted this notification. In Objective-C, it
// may be `nil`.
// 
// Typically you use this method to find out what object a notification
// applies to when you receive a notification.
// 
// For example, suppose you’ve registered an object to receive the message
// `` when the [PortInvalid] notification is posted to the notification center
// and that `` needs to access the object monitoring the port that is now
// invalid. `` can retrieve that object as shown here:
// 
// Example of accessing notification object in Objective-C:
//
// See: https://developer.apple.com/documentation/Foundation/NSNotification/object
func (n NSNotification) GetObject() objectivec.IObject {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("object"))
	return objectivec.Object{ID: rv}
}

// The user information dictionary associated with the notification.
//
// # Discussion
// 
// In Objective-C, this may be `nil`.
// 
// The user information dictionary stores any additional objects that objects
// receiving the notification might use.
// 
// For example, in AppKit, [NSControl] objects post the
// [textDidChangeNotification] whenever the field editor (an [NSText] object)
// changes text inside the [NSControl]. This notification provides the
// [NSControl] object as the notification’s associated object. In order to
// provide access to the field editor, the [NSControl] object posting the
// notification adds the field editor to the notification’s user information
// dictionary. Objects receiving the notification can access the field editor
// and the [NSControl] object posting the notification as follows:
//
// [NSControl]: https://developer.apple.com/documentation/AppKit/NSControl
// [NSText]: https://developer.apple.com/documentation/AppKit/NSText
// [textDidChangeNotification]: https://developer.apple.com/documentation/AppKit/NSControl/textDidChangeNotification
//
// See: https://developer.apple.com/documentation/Foundation/NSNotification/userInfo
func (n NSNotification) UserInfo() INSDictionary {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("userInfo"))
	return NSDictionaryFromID(objc.ID(rv))
}

			// Protocol methods for NSCoding
			

			// Protocol methods for NSCopying
			

