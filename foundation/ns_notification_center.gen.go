// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"context"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NotificationCenter] class.
var (
	_NotificationCenterClass     NotificationCenterClass
	_NotificationCenterClassOnce sync.Once
)

func getNotificationCenterClass() NotificationCenterClass {
	_NotificationCenterClassOnce.Do(func() {
		_NotificationCenterClass = NotificationCenterClass{class: objc.GetClass("NSNotificationCenter")}
	})
	return _NotificationCenterClass
}

// GetNotificationCenterClass returns the class object for NSNotificationCenter.
func GetNotificationCenterClass() NotificationCenterClass {
	return getNotificationCenterClass()
}

type NotificationCenterClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NotificationCenterClass) Alloc() NotificationCenter {
	rv := objc.Send[NotificationCenter](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A notification dispatch mechanism that enables the broadcast of information
// to registered observers.
//
// # Overview
// 
// Callers register with a notification center to receive one or both of the
// following:
// 
// - [NSNotification] objects, when working in Objective-C or with frameworks
// that only support [NSNotification]. Objects register with a notification
// center to receive notifications ([NSNotification] objects) using the
// [AddObserverSelectorNameObject] or
// [AddObserverForNameObjectQueueUsingBlock] methods, specifying a
// notification name and optionally a source object. When a caller adds itself
// as an observer, it specifies which notifications it should receive. -
// [NotificationCenter.MainActorMessage] and [NotificationCenter.AsyncMessage]
// instances for use with Swift code, providing strong typing, appropriate
// actor isolation, and a more idiomatic Swift experience. Callers register
// with the notification center using the various flavors of the
// `addObserver()` method, specifying either a message type or a convenience
// [NotificationCenter.MessageIdentifier] to identify the notification
// messages to receive. See [Notification center messages] for more
// information about this API.
// 
// Callers may add observers for many different notifications, or even the
// same notification name or message type as produced by different source
// objects.
// 
// Each running app has a [DefaultCenter] notification center, and you can
// create new notification centers to organize communications in particular
// contexts.
// 
// A notification center can deliver notifications only within a single
// program. On macOS, if you want to post a notification to other processes or
// receive notifications from other processes, use
// [NSDistributedNotificationCenter] instead.
//
// [Notification center messages]: https://developer.apple.com/documentation/Foundation/notification-center-messages
// [NotificationCenter.AsyncMessage]: https://developer.apple.com/documentation/Foundation/NotificationCenter/AsyncMessage
// [NotificationCenter.MainActorMessage]: https://developer.apple.com/documentation/Foundation/NotificationCenter/MainActorMessage
// [NotificationCenter.MessageIdentifier]: https://developer.apple.com/documentation/Foundation/NotificationCenter/MessageIdentifier
//
// # Adding and removing notification observers
//
//   - [NotificationCenter.AddObserverForNameObjectQueueUsingBlock]: Adds an entry to the notification center to receive notifications that passed to the provided block.
//   - [NotificationCenter.AddObserverSelectorNameObject]: Adds an entry to the notification center to call the provided selector with the notification.
//   - [NotificationCenter.RemoveObserverNameObject]: Removes matching entries from the notification center’s dispatch table.
//   - [NotificationCenter.RemoveObserver]: Removes all entries specifying an observer from the notification center’s dispatch table.
//
// # Posting notifications
//
//   - [NotificationCenter.PostNotification]: Posts a given notification to the notification center.
//   - [NotificationCenter.PostNotificationNameObjectUserInfo]: Creates a notification with a given name, sender, and information and posts it to the notification center.
//   - [NotificationCenter.PostNotificationNameObject]: Creates a notification with a given name and sender and posts it to the notification center.
//
// See: https://developer.apple.com/documentation/Foundation/NotificationCenter
type NotificationCenter struct {
	objectivec.Object
}

// NotificationCenterFromID constructs a [NotificationCenter] from an objc.ID.
//
// A notification dispatch mechanism that enables the broadcast of information
// to registered observers.
func NotificationCenterFromID(id objc.ID) NotificationCenter {
	return NSNotificationCenter{objectivec.Object{ID: id}}
}

// NSNotificationCenterFromID is an alias for [NotificationCenterFromID] for cross-framework compatibility.
func NSNotificationCenterFromID(id objc.ID) NotificationCenter { return NotificationCenterFromID(id) }
// NOTE: NotificationCenter adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NotificationCenter] class.
//
// # Adding and removing notification observers
//
//   - [INotificationCenter.AddObserverForNameObjectQueueUsingBlock]: Adds an entry to the notification center to receive notifications that passed to the provided block.
//   - [INotificationCenter.AddObserverSelectorNameObject]: Adds an entry to the notification center to call the provided selector with the notification.
//   - [INotificationCenter.RemoveObserverNameObject]: Removes matching entries from the notification center’s dispatch table.
//   - [INotificationCenter.RemoveObserver]: Removes all entries specifying an observer from the notification center’s dispatch table.
//
// # Posting notifications
//
//   - [INotificationCenter.PostNotification]: Posts a given notification to the notification center.
//   - [INotificationCenter.PostNotificationNameObjectUserInfo]: Creates a notification with a given name, sender, and information and posts it to the notification center.
//   - [INotificationCenter.PostNotificationNameObject]: Creates a notification with a given name and sender and posts it to the notification center.
//
// See: https://developer.apple.com/documentation/Foundation/NotificationCenter
type INotificationCenter interface {
	objectivec.IObject

	// Topic: Adding and removing notification observers

	// Adds an entry to the notification center to receive notifications that passed to the provided block.
	AddObserverForNameObjectQueueUsingBlock(name NSNotificationName, obj objectivec.IObject, queue INSOperationQueue, block NotificationHandler) objectivec.Object
	// Adds an entry to the notification center to call the provided selector with the notification.
	AddObserverSelectorNameObject(observer objectivec.IObject, aSelector objc.SEL, aName NSNotificationName, anObject objectivec.IObject)
	// Removes matching entries from the notification center’s dispatch table.
	RemoveObserverNameObject(observer objectivec.IObject, aName NSNotificationName, anObject objectivec.IObject)
	// Removes all entries specifying an observer from the notification center’s dispatch table.
	RemoveObserver(observer objectivec.IObject)

	// Topic: Posting notifications

	// Posts a given notification to the notification center.
	PostNotification(notification INSNotification)
	// Creates a notification with a given name, sender, and information and posts it to the notification center.
	PostNotificationNameObjectUserInfo(aName NSNotificationName, anObject objectivec.IObject, aUserInfo INSDictionary)
	// Creates a notification with a given name and sender and posts it to the notification center.
	PostNotificationNameObject(aName NSNotificationName, anObject objectivec.IObject)
}

// Init initializes the instance.
func (n NotificationCenter) Init() NotificationCenter {
	rv := objc.Send[NotificationCenter](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n NotificationCenter) Autorelease() NotificationCenter {
	rv := objc.Send[NotificationCenter](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewNotificationCenter creates a new NotificationCenter instance.
func NewNotificationCenter() NotificationCenter {
	class := getNotificationCenterClass()
	rv := objc.Send[NotificationCenter](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Adds an entry to the notification center to receive notifications that
// passed to the provided block.
//
// name: The name of the notification to register for delivery to the observer
// block. Specify a notification name to deliver only entries with this
// notification name.
// 
// When `nil`, the sender doesn’t use notification names as criteria for
// delivery.
//
// obj: The object that sends notifications to the observer block. Specify a sender
// to deliver only notifications from this sender.
// 
// When `nil`, the notification center doesn’t use the sender as criteria
// for the delivery.
//
// queue: The operation queue where the `block` runs.
// 
// When `nil`, the block runs synchronously on the posting thread.
//
// block: The block that executes when receiving a notification.
// 
// The notification center copies the block. The notification center strongly
// holds the copied block until you remove the observer registration.
// 
// The block takes one argument: the notification.
//
// # Return Value
// 
// An opaque object to act as the observer. Notification center strongly holds
// this return value until you remove the observer registration.
//
// # Discussion
// 
// If a notification triggers more than one observer block, the blocks can all
// execute concurrently (but on their queue or on the current thread).
// 
// The following example shows how you can register to receive locale change
// notifications. It stores the return value from
// [AddObserverForNameObjectQueueUsingBlock] in an instance property called
// `localeChangeObserver`.
// 
// Unregister an observer to stop receiving notifications. To unregister an
// observer, use `NotificationCenter/removeObserver(_:)` or
// [RemoveObserverNameObject] with the most specific detail possible. For
// example, if you used a name and object to register the observer, use the
// name and object to remove it.
// 
// You must invoke `NotificationCenter/removeObserver(_:)` or
// [RemoveObserverNameObject] before the system deallocates any object that
// [AddObserverForNameObjectQueueUsingBlock] specifies.
// 
// Another common practice is to create a one-time notification by removing
// the observer from within the observation block, as in the following
// example.
// 
// This example stores the opaque observer object in an instance property
// called `token`, which you can use to remove the observer prior to receiving
// the notification.
//
// See: https://developer.apple.com/documentation/Foundation/NotificationCenter/addObserver(forName:object:queue:using:)
func (n NotificationCenter) AddObserverForNameObjectQueueUsingBlock(name NSNotificationName, obj objectivec.IObject, queue INSOperationQueue, block NotificationHandler) objectivec.Object {
_block3, _cleanup3 := NewNotificationBlock(block)
	defer _cleanup3()
	rv := objc.Send[objc.ID](n.ID, objc.Sel("addObserverForName:object:queue:usingBlock:"), objc.String(string(name)), obj, queue, _block3)
	return objectivec.ObjectFromID(rv)
}

// Adds an entry to the notification center to call the provided selector with
// the notification.
//
// observer: An object to register as an observer.
//
// aSelector: A selector that specifies the message the receiver sends `observer` to
// alert it to the notification posting. The method that `aSelector` specifies
// must have one and only one argument (an instance of [NSNotification]).
//
// aName: The name of the notification to register for delivery to the observer.
// Specify a notification name to deliver only entries with this notification
// name.
// 
// When `nil`, the sender doesn’t use notification names as criteria for the
// delivery.
//
// anObject: The object that sends notifications to the observer. Specify a notification
// sender to deliver only notifications from this sender.
// 
// When `nil`, the notification center doesn’t use sender names as criteria
// for delivery.
//
// # Discussion
// 
// Unregister an observer to stop receiving notifications.
// 
// To unregister an observer, use `NotificationCenter/removeObserver(_:)` or
// [RemoveObserverNameObject] with the most specific detail possible. For
// example, if you used a name and object to register the observer, use the
// name and object to remove it.
// 
// If your app targets iOS 9.0 and later or macOS 10.11 and later, you do not
// need to unregister an observer that you created with this function. If you
// forget or are unable to remove an observer, the system cleans up the next
// time it would have posted to it.
//
// See: https://developer.apple.com/documentation/Foundation/NotificationCenter/addObserver(_:selector:name:object:)
func (n NotificationCenter) AddObserverSelectorNameObject(observer objectivec.IObject, aSelector objc.SEL, aName NSNotificationName, anObject objectivec.IObject) {
	objc.Send[objc.ID](n.ID, objc.Sel("addObserver:selector:name:object:"), observer, aSelector, objc.String(string(aName)), anObject)
}

// Removes matching entries from the notification center’s dispatch table.
//
// observer: The observer to remove from the dispatch table. Specify an observer to
// remove only entries for this observer.
//
// aName: The name of the notification to remove from the dispatch table. Specify a
// notification name to remove only entries with this notification name. When
// `nil`, the receiver does not use notification names as criteria for
// removal.
//
// anObject: The sender to remove from the dispatch table. Specify a notification sender
// to remove only entries with this sender. When `nil`, the receiver does not
// use a sender as criteria for removal.
//
// # Discussion
// 
// Removing the observer stops it from receiving notifications.
// 
// If you used [AddObserverForNameObjectQueueUsingBlock] to create your
// observer, you should call this method or
// `NotificationCenter/removeObserver(_:)` before the system deallocates any
// object that [AddObserverForNameObjectQueueUsingBlock] specifies.
// 
// If your app targets iOS 9.0 and later or macOS 10.11 and later, and you
// used [AddObserverSelectorNameObject] to create your observer, you do not
// need to unregister the observer. If you forget or are unable to remove the
// observer, the system cleans up the next time it would have posted to it.
// 
// When unregistering an observer, use the most specific detail possible. For
// example, if you used a name and object to register the observer, use
// [RemoveObserverNameObject] with the name and object.
//
// See: https://developer.apple.com/documentation/Foundation/NotificationCenter/removeObserver(_:name:object:)
func (n NotificationCenter) RemoveObserverNameObject(observer objectivec.IObject, aName NSNotificationName, anObject objectivec.IObject) {
	objc.Send[objc.ID](n.ID, objc.Sel("removeObserver:name:object:"), observer, objc.String(string(aName)), anObject)
}

// Removes all entries specifying an observer from the notification center’s
// dispatch table.
//
// observer: The observer to remove from the dispatch table. Specify an observer to
// remove only entries for this observer.
//
// # Discussion
// 
// Removing the observer stops it from receiving notifications.
// 
// If you used [AddObserverForNameObjectQueueUsingBlock] to create your
// observer, you should call this method or [RemoveObserverNameObject] before
// the system deallocates any object that
// [AddObserverForNameObjectQueueUsingBlock] specifies.
// 
// If your app targets iOS 9.0 and later or macOS 10.11 and later, and you
// used [AddObserverSelectorNameObject], you do not need to unregister the
// observer. If you forget or are unable to remove the observer, the system
// cleans up the next time it would have posted to it.
// 
// When removing an observer, remove it with the most specific detail
// possible. For example, if you used a name and object to register the
// observer, use [RemoveObserverNameObject] with the name and object.
// 
// The following example illustrates how to unregister `someObserver` for all
// previously registered notifications. This is safe to do in the [dealloc]
// method, but you shouldn’t use it otherwise (use
// [RemoveObserverNameObject] instead).
//
// [dealloc]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/dealloc
//
// See: https://developer.apple.com/documentation/Foundation/NotificationCenter/removeObserver(_:)-2yciv
func (n NotificationCenter) RemoveObserver(observer objectivec.IObject) {
	objc.Send[objc.ID](n.ID, objc.Sel("removeObserver:"), observer)
}

// Posts a given notification to the notification center.
//
// notification: The notification to post.
//
// See: https://developer.apple.com/documentation/Foundation/NotificationCenter/post(_:)-3x2st
func (n NotificationCenter) PostNotification(notification INSNotification) {
	objc.Send[objc.ID](n.ID, objc.Sel("postNotification:"), notification)
}

// Creates a notification with a given name, sender, and information and posts
// it to the notification center.
//
// aName: The name of the notification.
//
// anObject: The object posting the notification.
//
// aUserInfo: A user info dictionary with optional information about the notification.
//
// See: https://developer.apple.com/documentation/Foundation/NotificationCenter/post(name:object:userInfo:)
func (n NotificationCenter) PostNotificationNameObjectUserInfo(aName NSNotificationName, anObject objectivec.IObject, aUserInfo INSDictionary) {
	objc.Send[objc.ID](n.ID, objc.Sel("postNotificationName:object:userInfo:"), objc.String(string(aName)), anObject, aUserInfo)
}

// Creates a notification with a given name and sender and posts it to the
// notification center.
//
// aName: The name of the notification.
//
// anObject: The object posting the notification.
//
// # Discussion
// 
// This is a convenience method for calling
// [PostNotificationNameObjectUserInfo] and passing `nil` to `aUserInfo`.
//
// See: https://developer.apple.com/documentation/Foundation/NotificationCenter/post(name:object:)
func (n NotificationCenter) PostNotificationNameObject(aName NSNotificationName, anObject objectivec.IObject) {
	objc.Send[objc.ID](n.ID, objc.Sel("postNotificationName:object:"), objc.String(string(aName)), anObject)
}

// The app’s default notification center.
//
// # Discussion
// 
// All system notifications sent to an app are posted to the [DefaultCenter]
// notification center. You can also post your own notifications there.
// 
// If your app uses notifications extensively, you may want to create and post
// to your own notification centers rather than posting only to the
// [DefaultCenter] notification center. When a notification is posted to a
// notification center, the notification center scans through the list of
// registered observers, which may slow down your app. By organizing
// notifications functionally around one or more notification centers, less
// work is done each time a notification is posted, which can improve
// performance throughout your app.
//
// See: https://developer.apple.com/documentation/Foundation/NotificationCenter/default
func (_NotificationCenterClass NotificationCenterClass) DefaultCenter() NotificationCenter {
	rv := objc.Send[objc.ID](objc.ID(_NotificationCenterClass.class), objc.Sel("defaultCenter"))
	return NSNotificationCenterFromID(objc.ID(rv))
}

// AddObserverForNameObjectQueueUsingBlockSync is a synchronous wrapper around [NotificationCenter.AddObserverForNameObjectQueueUsingBlock].
// It blocks until the completion handler fires or the context is cancelled.
func (n NotificationCenter) AddObserverForNameObjectQueueUsingBlockSync(ctx context.Context, name NSNotificationName, obj objectivec.IObject, queue INSOperationQueue) (*NSNotification, error) {
	done := make(chan *NSNotification, 1)
	n.AddObserverForNameObjectQueueUsingBlock(name, obj, queue, func(val *NSNotification) {
		done <- val
	})
	select {
	case r := <-done:
		return r, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

