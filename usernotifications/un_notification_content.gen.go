// Code generated from Apple documentation for UserNotifications. DO NOT EDIT.

package usernotifications

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [UNNotificationContent] class.
var (
	_UNNotificationContentClass     UNNotificationContentClass
	_UNNotificationContentClassOnce sync.Once
)

func getUNNotificationContentClass() UNNotificationContentClass {
	_UNNotificationContentClassOnce.Do(func() {
		_UNNotificationContentClass = UNNotificationContentClass{class: objc.GetClass("UNNotificationContent")}
	})
	return _UNNotificationContentClass
}

// GetUNNotificationContentClass returns the class object for UNNotificationContent.
func GetUNNotificationContentClass() UNNotificationContentClass {
	return getUNNotificationContentClass()
}

type UNNotificationContentClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (uc UNNotificationContentClass) Class() objc.Class {
	return uc.class
}

// Alloc allocates memory for a new instance of the class.
func (uc UNNotificationContentClass) Alloc() UNNotificationContent {
	rv := objc.Send[UNNotificationContent](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// The uneditable content of a notification.
//
// # Overview
// 
// A [UNNotificationContent] object contains the data associated with a
// notification. When your app receives a notification, the associated
// [UNNotificationRequest] object contains an object of this type with the
// content that your app received. Use the content object to get the details
// of the notification, including the type of notification that the system
// delivered, any custom data you stored in the [UNNotificationContent.UserInfo] dictionary before
// scheduling the notification, and any attachments.
// 
// Don’t create instances of this class directly. For remote notifications,
// the system derives the contents of this object from the JSON payload that
// your server sends to the APNS server. For local notifications, create a
// [UNMutableNotificationContent] object, and configure the contents of that
// object instead.
//
// # Accessing the primary content
//
//   - [UNNotificationContent.Title]: The localized text that provides the notification’s primary description.
//   - [UNNotificationContent.Subtitle]: The localized text that provides the notification’s secondary description.
//   - [UNNotificationContent.Body]: The localized text that provides the notification’s main content.
//
// # Accessing supplementary content
//
//   - [UNNotificationContent.Attachments]: The visual and audio attachments to display alongside the notification’s main content.
//   - [UNNotificationContent.UserInfo]: The custom data to associate with the notification.
//
// # Reading app configuration
//
//   - [UNNotificationContent.Badge]: The number that your app’s icon displays.
//   - [UNNotificationContent.TargetContentIdentifier]: The value your app uses to determine which scene to display to handle the notification.
//
// # Reading system configuration
//
//   - [UNNotificationContent.Sound]: The sound that plays when the system delivers the notification.
//   - [UNNotificationContent.InterruptionLevel]: The notification’s importance and required delivery timing.
//   - [UNNotificationContent.RelevanceScore]: The score the system uses to determine if the notification is the summary’s featured notification.
//   - [UNNotificationContent.FilterCriteria]: The criteria the system evaluates to determine if it displays the notification in the current Focus.
//
// # Retrieving group information
//
//   - [UNNotificationContent.ThreadIdentifier]: The identifier that groups related notifications.
//   - [UNNotificationContent.CategoryIdentifier]: The identifier of the notification’s category.
//   - [UNNotificationContent.SummaryArgument]: The text the system adds to the notification summary to provide additional context.
//   - [UNNotificationContent.SummaryArgumentCount]: The number the system adds to the notification summary when the notification represents multiple items.
//
// # Updating the notification’s content
//
//   - [UNNotificationContent.ContentByUpdatingWithProviderError]: Returns a copy of the notification that includes content from the specified provider.
//
// See: https://developer.apple.com/documentation/UserNotifications/UNNotificationContent
type UNNotificationContent struct {
	objectivec.Object
}

// UNNotificationContentFromID constructs a [UNNotificationContent] from an objc.ID.
//
// The uneditable content of a notification.
func UNNotificationContentFromID(id objc.ID) UNNotificationContent {
	return UNNotificationContent{objectivec.Object{ID: id}}
}
// NOTE: UNNotificationContent adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [UNNotificationContent] class.
//
// # Accessing the primary content
//
//   - [IUNNotificationContent.Title]: The localized text that provides the notification’s primary description.
//   - [IUNNotificationContent.Subtitle]: The localized text that provides the notification’s secondary description.
//   - [IUNNotificationContent.Body]: The localized text that provides the notification’s main content.
//
// # Accessing supplementary content
//
//   - [IUNNotificationContent.Attachments]: The visual and audio attachments to display alongside the notification’s main content.
//   - [IUNNotificationContent.UserInfo]: The custom data to associate with the notification.
//
// # Reading app configuration
//
//   - [IUNNotificationContent.Badge]: The number that your app’s icon displays.
//   - [IUNNotificationContent.TargetContentIdentifier]: The value your app uses to determine which scene to display to handle the notification.
//
// # Reading system configuration
//
//   - [IUNNotificationContent.Sound]: The sound that plays when the system delivers the notification.
//   - [IUNNotificationContent.InterruptionLevel]: The notification’s importance and required delivery timing.
//   - [IUNNotificationContent.RelevanceScore]: The score the system uses to determine if the notification is the summary’s featured notification.
//   - [IUNNotificationContent.FilterCriteria]: The criteria the system evaluates to determine if it displays the notification in the current Focus.
//
// # Retrieving group information
//
//   - [IUNNotificationContent.ThreadIdentifier]: The identifier that groups related notifications.
//   - [IUNNotificationContent.CategoryIdentifier]: The identifier of the notification’s category.
//   - [IUNNotificationContent.SummaryArgument]: The text the system adds to the notification summary to provide additional context.
//   - [IUNNotificationContent.SummaryArgumentCount]: The number the system adds to the notification summary when the notification represents multiple items.
//
// # Updating the notification’s content
//
//   - [IUNNotificationContent.ContentByUpdatingWithProviderError]: Returns a copy of the notification that includes content from the specified provider.
//
// See: https://developer.apple.com/documentation/UserNotifications/UNNotificationContent
type IUNNotificationContent interface {
	objectivec.IObject

	// Topic: Accessing the primary content

	// The localized text that provides the notification’s primary description.
	Title() string
	// The localized text that provides the notification’s secondary description.
	Subtitle() string
	// The localized text that provides the notification’s main content.
	Body() string

	// Topic: Accessing supplementary content

	// The visual and audio attachments to display alongside the notification’s main content.
	Attachments() []UNNotificationAttachment
	// The custom data to associate with the notification.
	UserInfo() foundation.INSDictionary

	// Topic: Reading app configuration

	// The number that your app’s icon displays.
	Badge() foundation.NSNumber
	// The value your app uses to determine which scene to display to handle the notification.
	TargetContentIdentifier() string

	// Topic: Reading system configuration

	// The sound that plays when the system delivers the notification.
	Sound() IUNNotificationSound
	// The notification’s importance and required delivery timing.
	InterruptionLevel() UNNotificationInterruptionLevel
	// The score the system uses to determine if the notification is the summary’s featured notification.
	RelevanceScore() float64
	// The criteria the system evaluates to determine if it displays the notification in the current Focus.
	FilterCriteria() string

	// Topic: Retrieving group information

	// The identifier that groups related notifications.
	ThreadIdentifier() string
	// The identifier of the notification’s category.
	CategoryIdentifier() string
	// The text the system adds to the notification summary to provide additional context.
	SummaryArgument() string
	// The number the system adds to the notification summary when the notification represents multiple items.
	SummaryArgumentCount() uint

	// Topic: Updating the notification’s content

	// Returns a copy of the notification that includes content from the specified provider.
	ContentByUpdatingWithProviderError(provider UNNotificationContentProviding) (IUNNotificationContent, error)

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (u UNNotificationContent) Init() UNNotificationContent {
	rv := objc.Send[UNNotificationContent](u.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u UNNotificationContent) Autorelease() UNNotificationContent {
	rv := objc.Send[UNNotificationContent](u.ID, objc.Sel("autorelease"))
	return rv
}

// NewUNNotificationContent creates a new UNNotificationContent instance.
func NewUNNotificationContent() UNNotificationContent {
	class := getUNNotificationContentClass()
	rv := objc.Send[UNNotificationContent](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Returns a copy of the notification that includes content from the specified
// provider.
//
// provider: The notification content providing object.
//
// # Return Value
// 
// The notification content object for the content handler.
//
// # Discussion
// 
// The system contextualizes your [UNNotificationContent] object with other
// Apple SDK objects conforming to [UNNotificationContentProviding]. The
// system specializes the notification and decorates its look and behavior
// accordingly.
// 
// For example, the system treats the notification as a message with an avatar
// and promotes it to the top of notification center if the object passed in
// is a valid [INSendMessageIntent] that conforms to
// [UNNotificationContentProviding]. The system throws an error with a
// [UNError.Code], if the [UNNotificationContentProviding] object is invalid.
// Pass the valid [UNNotificationContent] result directly to
// [UNUserNotificationCenter] without mutating.
// 
// Add this call to the [UNNotificationServiceExtension] in
// [DidReceiveNotificationRequestWithContentHandler]. Your app passes the
// returned [UNNotificationContent] to the `contentHandler` for incoming push
// notifications.
//
// [INSendMessageIntent]: https://developer.apple.com/documentation/Intents/INSendMessageIntent
// [UNError.Code]: https://developer.apple.com/documentation/UserNotifications/UNError/Code
//
// See: https://developer.apple.com/documentation/UserNotifications/UNNotificationContent/updating(from:)
func (u UNNotificationContent) ContentByUpdatingWithProviderError(provider UNNotificationContentProviding) (IUNNotificationContent, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](u.ID, objc.Sel("contentByUpdatingWithProvider:error:"), provider, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return UNNotificationContent{}, foundation.NSErrorFrom(errorPtr)
	}
	return UNNotificationContentFromID(rv), nil

}
func (u UNNotificationContent) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](u.ID, objc.Sel("encodeWithCoder:"), coder)
}

// The localized text that provides the notification’s primary description.
//
// # Discussion
// 
// When a title is present, the system attempts to display a notification
// alert. If your app isn’t authorized to display alert-based notifications,
// the system ignores this property.
// 
// Title strings should be short, usually only a couple of words describing
// the reason for the notification. In watchOS, the system displays the title
// string as part of the short look notification interface, which has limited
// space.
//
// See: https://developer.apple.com/documentation/UserNotifications/UNNotificationContent/title
func (u UNNotificationContent) Title() string {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("title"))
	return foundation.NSStringFromID(rv).String()
}
// The localized text that provides the notification’s secondary
// description.
//
// # Discussion
// 
// Subtitles offer additional context in cases where the title alone isn’t
// clear. Subtitles aren’t displayed in all cases. If your app isn’t
// authorized to display alert-based notifications, the system ignores this
// property.
//
// See: https://developer.apple.com/documentation/UserNotifications/UNNotificationContent/subtitle
func (u UNNotificationContent) Subtitle() string {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("subtitle"))
	return foundation.NSStringFromID(rv).String()
}
// The localized text that provides the notification’s main content.
//
// # Discussion
// 
// The body text contains the final text that you want to display. If your app
// isn’t authorized to display alert-based notifications, the system ignores
// this property.
// 
// If you specified two percent symbols (`%%`) in the message body, the system
// replaces it with a single percent symbol (`%`). The system strips all other
// printf style escape characters from your string prior to display.
//
// See: https://developer.apple.com/documentation/UserNotifications/UNNotificationContent/body
func (u UNNotificationContent) Body() string {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("body"))
	return foundation.NSStringFromID(rv).String()
}
// The visual and audio attachments to display alongside the notification’s
// main content.
//
// # Discussion
// 
// Use this property to retrieve the images, movies, and audio files
// associated with your notification’s content. A notification content app
// extension might use these values to add the associated content to its view
// controller.
//
// See: https://developer.apple.com/documentation/UserNotifications/UNNotificationContent/attachments
func (u UNNotificationContent) Attachments() []UNNotificationAttachment {
	rv := objc.Send[[]objc.ID](u.ID, objc.Sel("attachments"))
	return objc.ConvertSlice(rv, func(id objc.ID) UNNotificationAttachment {
		return UNNotificationAttachmentFromID(id)
	})
}
// The custom data to associate with the notification.
//
// # Discussion
// 
// For remote notifications, this property contains the entire notification
// payload. For local notifications, you configure the property directly
// before scheduling the notification.
// 
// The keys in this dictionary must be property-list types—that’s, they
// must be types that can be serialized into the property-list format. For
// information about property-list types, see [Property List Programming
// Guide].
//
// [Property List Programming Guide]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/PropertyLists/Introduction/Introduction.html#//apple_ref/doc/uid/10000048i
//
// See: https://developer.apple.com/documentation/UserNotifications/UNNotificationContent/userInfo
func (u UNNotificationContent) UserInfo() foundation.INSDictionary {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("userInfo"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
// The number that your app’s icon displays.
//
// # Discussion
// 
// When the number in this property is `0`, the system doesn’t display a
// badge. When the number is greater than `0`, the system displays the badge
// with the specified number. When the value in this property is `nil`, the
// system leaves the current badge unchanged.
//
// See: https://developer.apple.com/documentation/UserNotifications/UNNotificationContent/badge
func (u UNNotificationContent) Badge() foundation.NSNumber {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("badge"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
// The value your app uses to determine which scene to display to handle the
// notification.
//
// # Discussion
// 
// Use this value to determine the content to show in your app when the user
// taps the notification.
//
// See: https://developer.apple.com/documentation/UserNotifications/UNNotificationContent/targetContentIdentifier
func (u UNNotificationContent) TargetContentIdentifier() string {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("targetContentIdentifier"))
	return foundation.NSStringFromID(rv).String()
}
// The sound that plays when the system delivers the notification.
//
// # Discussion
// 
// Notifications can play a default sound or a custom sound. For information
// on how to specify custom sounds for your notifications, see
// [UNNotificationSound].
//
// See: https://developer.apple.com/documentation/UserNotifications/UNNotificationContent/sound
func (u UNNotificationContent) Sound() IUNNotificationSound {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("sound"))
	return UNNotificationSoundFromID(objc.ID(rv))
}
// The notification’s importance and required delivery timing.
//
// See: https://developer.apple.com/documentation/UserNotifications/UNNotificationContent/interruptionLevel
func (u UNNotificationContent) InterruptionLevel() UNNotificationInterruptionLevel {
	rv := objc.Send[UNNotificationInterruptionLevel](u.ID, objc.Sel("interruptionLevel"))
	return UNNotificationInterruptionLevel(rv)
}
// The score the system uses to determine if the notification is the
// summary’s featured notification.
//
// # Discussion
// 
// The system uses the `relevanceScore`, a value between `0` and `1`, to sort
// the notifications from your app. The highest score gets featured in the
// notification summary.
//
// See: https://developer.apple.com/documentation/UserNotifications/UNNotificationContent/relevanceScore
func (u UNNotificationContent) RelevanceScore() float64 {
	rv := objc.Send[float64](u.ID, objc.Sel("relevanceScore"))
	return rv
}
// The criteria the system evaluates to determine if it displays the
// notification in the current Focus.
//
// # Discussion
// 
// For more information, see [SetFocusFilterIntent].
//
// [SetFocusFilterIntent]: https://developer.apple.com/documentation/AppIntents/SetFocusFilterIntent
//
// See: https://developer.apple.com/documentation/UserNotifications/UNNotificationContent/filterCriteria
func (u UNNotificationContent) FilterCriteria() string {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("filterCriteria"))
	return foundation.NSStringFromID(rv).String()
}
// The identifier that groups related notifications.
//
// # Discussion
// 
// For remote notifications, the system sets this property to the value of the
// `thread-id` key in the `aps` dictionary.
//
// See: https://developer.apple.com/documentation/UserNotifications/UNNotificationContent/threadIdentifier
func (u UNNotificationContent) ThreadIdentifier() string {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("threadIdentifier"))
	return foundation.NSStringFromID(rv).String()
}
// The identifier of the notification’s category.
//
// # Discussion
// 
// Use notification types to distinguish between the different types of
// notifications your app supports. You use this support primarily to create
// actionable notifications with custom action buttons, and to redirect your
// notifications through either your notification service app extension or
// your notification content app extension.
// 
// For remote notifications, the system sets this property to the value of the
// `category` key in the `aps` dictionary.
//
// See: https://developer.apple.com/documentation/UserNotifications/UNNotificationContent/categoryIdentifier
func (u UNNotificationContent) CategoryIdentifier() string {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("categoryIdentifier"))
	return foundation.NSStringFromID(rv).String()
}
// The text the system adds to the notification summary to provide additional
// context.
//
// See: https://developer.apple.com/documentation/UserNotifications/UNNotificationContent/summaryArgument
func (u UNNotificationContent) SummaryArgument() string {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("summaryArgument"))
	return foundation.NSStringFromID(rv).String()
}
// The number the system adds to the notification summary when the
// notification represents multiple items.
//
// See: https://developer.apple.com/documentation/UserNotifications/UNNotificationContent/summaryArgumentCount
func (u UNNotificationContent) SummaryArgumentCount() uint {
	rv := objc.Send[uint](u.ID, objc.Sel("summaryArgumentCount"))
	return rv
}

