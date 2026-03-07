// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NotificationQueue] class.
var (
	_NotificationQueueClass     NotificationQueueClass
	_NotificationQueueClassOnce sync.Once
)

func getNotificationQueueClass() NotificationQueueClass {
	_NotificationQueueClassOnce.Do(func() {
		_NotificationQueueClass = NotificationQueueClass{class: objc.GetClass("NSNotificationQueue")}
	})
	return _NotificationQueueClass
}

// GetNotificationQueueClass returns the class object for NSNotificationQueue.
func GetNotificationQueueClass() NotificationQueueClass {
	return getNotificationQueueClass()
}

type NotificationQueueClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NotificationQueueClass) Alloc() NotificationQueue {
	rv := objc.Send[NotificationQueue](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// A notification center buffer.
//
// # Overview
// 
// Whereas a notification center distributes notifications when posted,
// notifications placed into the queue can be delayed until the end of the
// current pass through the run loop or until the run loop is idle. Duplicate
// notifications can be coalesced so that only one notification is sent
// although multiple notifications are posted.
// 
// A notification queue maintains notifications in first in, first out (FIFO)
// order. When a notification moves to the front of the queue, the queue posts
// it to the notification center, which in turn dispatches the notification to
// all objects registered as observers.
// 
// Every thread has a default notification queue, which is associated with the
// default notification center for the process. You can create your own
// notification queues and have multiple queues per center and thread.
//
// # Creating Notification Queues
//
//   - [NotificationQueue.InitWithNotificationCenter]: Initializes and returns a notification queue for the specified notification center.
//
// # Managing Notifications
//
//   - [NotificationQueue.EnqueueNotificationPostingStyleCoalesceMaskForModes]: Adds a notification to the notification queue with a specified posting style, criteria for coalescing, and run loop mode.
//   - [NotificationQueue.EnqueueNotificationPostingStyle]: Adds a notification to the notification queue with a specified posting style.
//   - [NotificationQueue.DequeueNotificationsMatchingCoalesceMask]: Removes all notifications from the queue that match a provided notification using provided matching criteria.
//
// See: https://developer.apple.com/documentation/Foundation/NotificationQueue
type NotificationQueue struct {
	objectivec.Object
}

// NotificationQueueFromID constructs a [NotificationQueue] from an objc.ID.
//
// A notification center buffer.
func NotificationQueueFromID(id objc.ID) NotificationQueue {
	return NSNotificationQueue{objectivec.Object{id}}
}

// NSNotificationQueueFromID is an alias for [NotificationQueueFromID] for cross-framework compatibility.
func NSNotificationQueueFromID(id objc.ID) NotificationQueue { return NotificationQueueFromID(id) }
// NOTE: NotificationQueue adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [NotificationQueue] class.
//
// # Creating Notification Queues
//
//   - [INotificationQueue.InitWithNotificationCenter]: Initializes and returns a notification queue for the specified notification center.
//
// # Managing Notifications
//
//   - [INotificationQueue.EnqueueNotificationPostingStyleCoalesceMaskForModes]: Adds a notification to the notification queue with a specified posting style, criteria for coalescing, and run loop mode.
//   - [INotificationQueue.EnqueueNotificationPostingStyle]: Adds a notification to the notification queue with a specified posting style.
//   - [INotificationQueue.DequeueNotificationsMatchingCoalesceMask]: Removes all notifications from the queue that match a provided notification using provided matching criteria.
//
// See: https://developer.apple.com/documentation/Foundation/NotificationQueue
type INotificationQueue interface {
	objectivec.IObject

	// Topic: Creating Notification Queues

	// Initializes and returns a notification queue for the specified notification center.
	InitWithNotificationCenter(notificationCenter INSNotificationCenter) NotificationQueue

	// Topic: Managing Notifications

	// Adds a notification to the notification queue with a specified posting style, criteria for coalescing, and run loop mode.
	EnqueueNotificationPostingStyleCoalesceMaskForModes(notification INSNotification, postingStyle NSPostingStyle, coalesceMask NSNotificationCoalescing, modes []string)
	// Adds a notification to the notification queue with a specified posting style.
	EnqueueNotificationPostingStyle(notification INSNotification, postingStyle NSPostingStyle)
	// Removes all notifications from the queue that match a provided notification using provided matching criteria.
	DequeueNotificationsMatchingCoalesceMask(notification INSNotification, coalesceMask uint)
}





// Init initializes the instance.
func (n NotificationQueue) Init() NotificationQueue {
	rv := objc.Send[NotificationQueue](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n NotificationQueue) Autorelease() NotificationQueue {
	rv := objc.Send[NotificationQueue](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewNotificationQueue creates a new NotificationQueue instance.
func NewNotificationQueue() NotificationQueue {
	class := getNotificationQueueClass()
	rv := objc.Send[NotificationQueue](objc.ID(class.class), objc.Sel("new"))
	return rv
}






// Initializes and returns a notification queue for the specified notification
// center.
//
// notificationCenter: The notification center to be used by the notification queue.
//
// # Return Value
// 
// The newly initialized notification queue.
//
// # Discussion
// 
// This is the designated initializer for the [NSNotificationQueue] class.
//
// See: https://developer.apple.com/documentation/Foundation/NotificationQueue/init(notificationCenter:)
func NewNotificationQueueWithNotificationCenter(notificationCenter INSNotificationCenter) NotificationQueue {
	instance := getNotificationQueueClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithNotificationCenter:"), notificationCenter)
	return NotificationQueueFromID(rv)
}







// Initializes and returns a notification queue for the specified notification
// center.
//
// notificationCenter: The notification center to be used by the notification queue.
//
// # Return Value
// 
// The newly initialized notification queue.
//
// # Discussion
// 
// This is the designated initializer for the [NSNotificationQueue] class.
//
// See: https://developer.apple.com/documentation/Foundation/NotificationQueue/init(notificationCenter:)
func (n NotificationQueue) InitWithNotificationCenter(notificationCenter INSNotificationCenter) NotificationQueue {
	rv := objc.Send[NotificationQueue](n.ID, objc.Sel("initWithNotificationCenter:"), notificationCenter)
	return rv
}

// Adds a notification to the notification queue with a specified posting
// style, criteria for coalescing, and run loop mode.
//
// notification: The notification to add to the queue.
//
// postingStyle: The posting style for the notification. The posting style indicates when
// the notification queue should post the notification to its notification
// center.
//
// coalesceMask: A mask indicating what criteria to use when matching attributes of
// `notification` to attributes of notifications in the queue. The mask is
// created by combining any of the constants [NotificationNoCoalescing],
// [NotificationCoalescingOnName], and [NotificationCoalescingOnSender].
//
// modes: The list of modes the notification may be posted in. The notification queue
// will only post the notification to its notification center if the run loop
// is in one of the modes provided in the array.
// 
// This parameter may be `nil`, in which case it defaults to [default].
// //
// [default]: https://developer.apple.com/documentation/Foundation/RunLoop/Mode/default
//
// See: https://developer.apple.com/documentation/Foundation/NotificationQueue/enqueue(_:postingStyle:coalesceMask:forModes:)
func (n NotificationQueue) EnqueueNotificationPostingStyleCoalesceMaskForModes(notification INSNotification, postingStyle NSPostingStyle, coalesceMask NSNotificationCoalescing, modes []string) {
	objc.Send[objc.ID](n.ID, objc.Sel("enqueueNotification:postingStyle:coalesceMask:forModes:"), notification, postingStyle, coalesceMask, objectivec.StringSliceToNSArray(modes))
}

// Adds a notification to the notification queue with a specified posting
// style.
//
// notification: The notification to add to the queue.
//
// postingStyle: The posting style for the notification. The posting style indicates when
// the notification queue should post the notification to its notification
// center.
//
// # Discussion
// 
// This is a convenience method for calling
// [EnqueueNotificationPostingStyleCoalesceMaskForModes] with coalescing
// criteria that will coalesce only notifications that match both the
// notification’s name and object and the runloop mode [default].
//
// [default]: https://developer.apple.com/documentation/Foundation/RunLoop/Mode/default
//
// See: https://developer.apple.com/documentation/Foundation/NotificationQueue/enqueue(_:postingStyle:)
func (n NotificationQueue) EnqueueNotificationPostingStyle(notification INSNotification, postingStyle NSPostingStyle) {
	objc.Send[objc.ID](n.ID, objc.Sel("enqueueNotification:postingStyle:"), notification, postingStyle)
}

// Removes all notifications from the queue that match a provided notification
// using provided matching criteria.
//
// notification: The notification used for matching notifications to remove from the
// notification queue.
//
// coalesceMask: A mask indicating what criteria to use when matching attributes of
// `notification` to attributes of notifications in the queue. The mask is
// created by combining any of the constants [NotificationNoCoalescing],
// [NotificationCoalescingOnName], and [NotificationCoalescingOnSender].
//
// See: https://developer.apple.com/documentation/Foundation/NotificationQueue/dequeueNotifications(matching:coalesceMask:)
func (n NotificationQueue) DequeueNotificationsMatchingCoalesceMask(notification INSNotification, coalesceMask uint) {
	objc.Send[objc.ID](n.ID, objc.Sel("dequeueNotificationsMatching:coalesceMask:"), notification, coalesceMask)
}
















// Returns the default notification queue for the current thread.
//
// # Return Value
// 
// Returns the default notification queue for the current thread.
// 
// # Discussion
// 
// This notification queue uses the default notification center.
//
// See: https://developer.apple.com/documentation/Foundation/NotificationQueue/default
func (_NotificationQueueClass NotificationQueueClass) DefaultQueue() NotificationQueue {
	rv := objc.Send[objc.ID](objc.ID(_NotificationQueueClass.class), objc.Sel("defaultQueue"))
	return NSNotificationQueueFromID(objc.ID(rv))
}



















