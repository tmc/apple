// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [DistributedNotificationCenter] class.
var (
	_DistributedNotificationCenterClass     DistributedNotificationCenterClass
	_DistributedNotificationCenterClassOnce sync.Once
)

func getDistributedNotificationCenterClass() DistributedNotificationCenterClass {
	_DistributedNotificationCenterClassOnce.Do(func() {
		_DistributedNotificationCenterClass = DistributedNotificationCenterClass{class: objc.GetClass("NSDistributedNotificationCenter")}
	})
	return _DistributedNotificationCenterClass
}

// GetDistributedNotificationCenterClass returns the class object for NSDistributedNotificationCenter.
func GetDistributedNotificationCenterClass() DistributedNotificationCenterClass {
	return getDistributedNotificationCenterClass()
}

type DistributedNotificationCenterClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (dc DistributedNotificationCenterClass) Class() objc.Class {
	return dc.class
}

// Alloc allocates memory for a new instance of the class.
func (dc DistributedNotificationCenterClass) Alloc() DistributedNotificationCenter {
	rv := objc.Send[DistributedNotificationCenter](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// A notification dispatch mechanism that enables the broadcast of
// notifications across task boundaries.
//
// # Overview
// 
// A [NSDistributedNotificationCenter] instance broadcasts [NSNotification]
// objects to objects in other tasks that have registered for the notification
// with their task’s default distributed notification center.
// 
// # Principal Attributes
// 
// - Notification dispatch table. See “Class at a Glance” > “Principal
// Attributes” in [NSNotificationCenter] for information about the dispatch
// table.
// 
// In addition to the notification name and sender, dispatch table entries for
// distributed notification centers specify when the notification center
// delivers notifications to its observers. See the
// [PostNotificationNameObjectUserInfoDeliverImmediately] method, Suspending
// and Resuming Notification Delivery, and
// [DistributedNotificationCenter.SuspensionBehavior] for details.
// 
// # Commonly Used Methods
// 
// [DefaultCenter]: Accesses the default distributed notification center.
// [AddObserverSelectorNameObjectSuspensionBehavior]: Registers an object to
// receive a notification with a specified behavior when notification delivery
// is suspended. [PostNotificationNameObjectUserInfoDeliverImmediately]:
// Creates and posts a notification. [RemoveObserverNameObject]: Specifies
// that an object no longer wants to receive certain notifications.
// 
// # Overview
// 
// Each task has a default distributed notification center that you access
// with the [DefaultCenter] class method. There may be different types of
// distributed notification centers. Currently there is a single
// type—[NSLocalNotificationCenterType]. This type of distributed
// notification center handles notifications that can be sent between tasks on
// a single computer. For communication between tasks on different computers,
// use [Distributed Objects Programming Topics].
// 
// Posting a is an expensive operation. The notification gets sent to a
// system-wide server that distributes it to all the tasks that have objects
// registered for distributed notifications. The latency between posting the
// notification and the notification’s arrival in another task is unbounded.
// In fact, when too many notifications are posted and the server’s queue
// fills up, notifications may be dropped.
// 
// Distributed notifications are delivered via a task’s run loop. A task
// must be running a run loop in one of the “common” modes, such as
// [NSDefaultRunLoopMode], to receive a distributed notification. For
// multithreaded applications running in macOS 10.3 and later, distributed
// notifications are always delivered to the main thread. For multithreaded
// applications running in OS X v10.2.8 and earlier, notifications are
// delivered to the thread that first used the distributed notifications API,
// which in most cases is the main thread.
//
// [Distributed Objects Programming Topics]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/DistrObjects/DistrObjects.html#//apple_ref/doc/uid/10000102i
// [DistributedNotificationCenter.SuspensionBehavior]: https://developer.apple.com/documentation/Foundation/DistributedNotificationCenter/SuspensionBehavior
//
// # Managing Observers
//
//   - [DistributedNotificationCenter.AddObserverSelectorNameObjectSuspensionBehavior]: Adds an entry to the receiver’s dispatch table with a specific observer and suspended-notifications behavior, and optional notification name and sender.
//
// # Posting Notifications
//
//   - [DistributedNotificationCenter.PostNotificationNameObjectUserInfoDeliverImmediately]: Creates a notification with information and an immediate-delivery specifier, and posts it to the receiver.
//   - [DistributedNotificationCenter.PostNotificationNameObjectUserInfoOptions]: Creates a notification with information, and posts it to the receiver.
//
// # Suspending and Resuming Notification Delivery
//
//   - [DistributedNotificationCenter.Suspended]: Suspends or resumes notification delivery.
//   - [DistributedNotificationCenter.SetSuspended]
//
// See: https://developer.apple.com/documentation/Foundation/DistributedNotificationCenter
type DistributedNotificationCenter struct {
	NSNotificationCenter
}

// DistributedNotificationCenterFromID constructs a [DistributedNotificationCenter] from an objc.ID.
//
// A notification dispatch mechanism that enables the broadcast of
// notifications across task boundaries.
func DistributedNotificationCenterFromID(id objc.ID) DistributedNotificationCenter {
	return NSDistributedNotificationCenter{NSNotificationCenter: NSNotificationCenterFromID(id)}
}

// NSDistributedNotificationCenterFromID is an alias for [DistributedNotificationCenterFromID] for cross-framework compatibility.
func NSDistributedNotificationCenterFromID(id objc.ID) DistributedNotificationCenter { return DistributedNotificationCenterFromID(id) }
// NOTE: DistributedNotificationCenter adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [DistributedNotificationCenter] class.
//
// # Managing Observers
//
//   - [IDistributedNotificationCenter.AddObserverSelectorNameObjectSuspensionBehavior]: Adds an entry to the receiver’s dispatch table with a specific observer and suspended-notifications behavior, and optional notification name and sender.
//
// # Posting Notifications
//
//   - [IDistributedNotificationCenter.PostNotificationNameObjectUserInfoDeliverImmediately]: Creates a notification with information and an immediate-delivery specifier, and posts it to the receiver.
//   - [IDistributedNotificationCenter.PostNotificationNameObjectUserInfoOptions]: Creates a notification with information, and posts it to the receiver.
//
// # Suspending and Resuming Notification Delivery
//
//   - [IDistributedNotificationCenter.Suspended]: Suspends or resumes notification delivery.
//   - [IDistributedNotificationCenter.SetSuspended]
//
// See: https://developer.apple.com/documentation/Foundation/DistributedNotificationCenter
type IDistributedNotificationCenter interface {
	INSNotificationCenter

	// Topic: Managing Observers

	// Adds an entry to the receiver’s dispatch table with a specific observer and suspended-notifications behavior, and optional notification name and sender.
	AddObserverSelectorNameObjectSuspensionBehavior(observer objectivec.IObject, selector objc.SEL, name NSNotificationName, object string, suspensionBehavior NSNotificationSuspensionBehavior)

	// Topic: Posting Notifications

	// Creates a notification with information and an immediate-delivery specifier, and posts it to the receiver.
	PostNotificationNameObjectUserInfoDeliverImmediately(name NSNotificationName, object string, userInfo INSDictionary, deliverImmediately bool)
	// Creates a notification with information, and posts it to the receiver.
	PostNotificationNameObjectUserInfoOptions(name NSNotificationName, object string, userInfo INSDictionary, options NSDistributedNotificationOptions)

	// Topic: Suspending and Resuming Notification Delivery

	// Suspends or resumes notification delivery.
	Suspended() bool
	SetSuspended(value bool)
}

// Init initializes the instance.
func (d DistributedNotificationCenter) Init() DistributedNotificationCenter {
	rv := objc.Send[DistributedNotificationCenter](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d DistributedNotificationCenter) Autorelease() DistributedNotificationCenter {
	rv := objc.Send[DistributedNotificationCenter](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewDistributedNotificationCenter creates a new DistributedNotificationCenter instance.
func NewDistributedNotificationCenter() DistributedNotificationCenter {
	class := getDistributedNotificationCenterClass()
	rv := objc.Send[DistributedNotificationCenter](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Adds an entry to the receiver’s dispatch table with a specific observer
// and suspended-notifications behavior, and optional notification name and
// sender.
//
// observer: Object registering as an observer. Must not be `nil`.
//
// selector: Selector that specifies the message the receiver sends
// `notificationObserver` to notify it of the notification posting. Must not
// be `0`.
//
// name: The name of the notification for which to register the observer; that is,
// only notifications with this name are delivered to the observer. When
// `nil`, the notification center doesn’t use a notification’s name to
// decide whether to deliver it to the observer.
//
// object: The object whose notifications the observer wants to receive; that is, only
// notifications sent by this sender are delivered to the observer. When
// `nil`, the notification center doesn’t use a notification’s sender to
// decide whether to deliver it to the observer.
//
// suspensionBehavior: Notification posting behavior when notification delivery is suspended.
//
// # Discussion
// 
// The receiver does not retain `notificationObserver`. Therefore, you should
// always send `NotificationCenter/removeObserver(_:)` or
// [RemoveObserverNameObject] to the receiver before releasing
// `notificationObserver`.
//
// See: https://developer.apple.com/documentation/Foundation/DistributedNotificationCenter/addObserver(_:selector:name:object:suspensionBehavior:)
func (d DistributedNotificationCenter) AddObserverSelectorNameObjectSuspensionBehavior(observer objectivec.IObject, selector objc.SEL, name NSNotificationName, object string, suspensionBehavior NSNotificationSuspensionBehavior) {
	objc.Send[objc.ID](d.ID, objc.Sel("addObserver:selector:name:object:suspensionBehavior:"), observer, selector, objc.String(string(name)), objc.String(object), suspensionBehavior)
}
// Creates a notification with information and an immediate-delivery
// specifier, and posts it to the receiver.
//
// name: Name of the notification to post. Must not be `nil`.
//
// object: Sender of the notification. May be `nil`.
//
// userInfo: Dictionary containing additional information. May be `nil`.
//
// deliverImmediately: Specifies when to deliver the notification. When [false], the receiver
// delivers notifications to their observers according to the
// suspended-notification behavior specified in the corresponding dispatch
// table entry. When [true], the receiver delivers the notification
// immediately to its observers.
// //
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// This is the preferred method for posting notifications.
// 
// The `notificationInfo` dictionary is serialized as a property list, so it
// can be passed to another task. In the receiving task, it is deserialized
// back into a dictionary. This serialization imposes some restrictions on the
// objects that can be placed in the `notificationInfo` dictionary. See XML
// Property Lists for details.
//
// See: https://developer.apple.com/documentation/Foundation/DistributedNotificationCenter/postNotificationName(_:object:userInfo:deliverImmediately:)
func (d DistributedNotificationCenter) PostNotificationNameObjectUserInfoDeliverImmediately(name NSNotificationName, object string, userInfo INSDictionary, deliverImmediately bool) {
	objc.Send[objc.ID](d.ID, objc.Sel("postNotificationName:object:userInfo:deliverImmediately:"), objc.String(string(name)), objc.String(object), userInfo, deliverImmediately)
}
// Creates a notification with information, and posts it to the receiver.
//
// name: Name of the notification to post. Must not be `nil`.
//
// object: Sender of the notification. May be `nil`.
//
// userInfo: Dictionary containing additional information. May be `nil`.
//
// options: Specifies how the notification is posted to the task and when to deliver it
// to its observers. See `Notification Posting Behavior` for details.
//
// # Discussion
// 
// The `userInfo` dictionary is serialized as a property list, so it can be
// passed to another task. In the receiving task, it is deserialized back into
// a dictionary. This serialization imposes some restrictions on the objects
// that can be placed in the `userInfo` dictionary. See XML Property Lists for
// details.
//
// See: https://developer.apple.com/documentation/Foundation/DistributedNotificationCenter/postNotificationName(_:object:userInfo:options:)
func (d DistributedNotificationCenter) PostNotificationNameObjectUserInfoOptions(name NSNotificationName, object string, userInfo INSDictionary, options NSDistributedNotificationOptions) {
	objc.Send[objc.ID](d.ID, objc.Sel("postNotificationName:object:userInfo:options:"), objc.String(string(name)), objc.String(object), userInfo, options)
}

// Returns the distributed notification center for a particular notification
// center type.
//
// notificationCenterType: Notification center type being inquired about.
//
// # Return Value
// 
// Distributed notification center for `notificationCenterType`.
//
// # Discussion
// 
// Currently only one type, [NSLocalNotificationCenterType], is supported.
//
// See: https://developer.apple.com/documentation/Foundation/DistributedNotificationCenter/forType(_:)
func (_DistributedNotificationCenterClass DistributedNotificationCenterClass) NotificationCenterForType(notificationCenterType NSDistributedNotificationCenterType) DistributedNotificationCenter {
	rv := objc.Send[objc.ID](objc.ID(_DistributedNotificationCenterClass.class), objc.Sel("notificationCenterForType:"), objc.String(string(notificationCenterType)))
	return NSDistributedNotificationCenterFromID(rv)
}

// Suspends or resumes notification delivery.
//
// # Discussion
// 
// See [DistributedNotificationCenter.SuspensionBehavior] for details on how
// the receiver delivers notifications to their observers when normal
// notification delivery is suspended.
// 
// The [NSApplication] class automatically suspends distributed notification
// delivery when the application is not active. Applications based on the
// Application Kit framework should let AppKit manage the suspension of
// notification delivery. Foundation-only programs may have occasional need to
// use this method.
//
// [DistributedNotificationCenter.SuspensionBehavior]: https://developer.apple.com/documentation/Foundation/DistributedNotificationCenter/SuspensionBehavior
// [NSApplication]: https://developer.apple.com/documentation/AppKit/NSApplication
//
// See: https://developer.apple.com/documentation/Foundation/DistributedNotificationCenter/suspended
func (d DistributedNotificationCenter) Suspended() bool {
	rv := objc.Send[bool](d.ID, objc.Sel("suspended"))
	return rv
}
func (d DistributedNotificationCenter) SetSuspended(value bool) {
	objc.Send[struct{}](d.ID, objc.Sel("setSuspended:"), value)
}

