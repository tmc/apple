// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"context"
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSUserActivity] class.
var (
	_NSUserActivityClass     NSUserActivityClass
	_NSUserActivityClassOnce sync.Once
)

func getNSUserActivityClass() NSUserActivityClass {
	_NSUserActivityClassOnce.Do(func() {
		_NSUserActivityClass = NSUserActivityClass{class: objc.GetClass("NSUserActivity")}
	})
	return _NSUserActivityClass
}

// GetNSUserActivityClass returns the class object for NSUserActivity.
func GetNSUserActivityClass() NSUserActivityClass {
	return getNSUserActivityClass()
}

type NSUserActivityClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSUserActivityClass) Alloc() NSUserActivity {
	rv := objc.Send[NSUserActivity](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// A representation of the state of your app at a moment in time.
//
// # Overview
// 
// An [NSUserActivity] object provides a lightweight way to capture the state
// of your app and put it to use later. Create this object to capture
// information about what a person was doing, such as viewing app content,
// editing a document, viewing a web page, or watching a video. When the
// system launches your app and an activity object is available, your app can
// use the information in that object to restore itself to an appropriate
// state. Spotlight also uses these objects to improve search results for
// people. To allow people to continue an activity on another device, see
// [Implementing Handoff in Your App].
// 
// # Siri
// 
// If SiriKit needs to launch your app for any reason, it creates a user
// activity object and assigns an appropriate [INInteraction] object to its
// [NSUserActivity.Interaction] property. Your app can use the interaction information to
// configure itself and display information related to the interaction started
// by SiriKit. You can also provide SiriKit with a custom user activity object
// containing additional data that you want passed to your app.
// 
// In iOS 15 and later, a person can share content they’re viewing by asking
// Siri to “share this”. Apps built with Mac Catalyst provide the same
// capability with an [NSSharingServicePickerToolbarItem] in the toolbar. You
// can use [NSUserActivity.ActivityItemsConfiguration] or [NSUserActivity.ActivityItemsConfigurationSource]
// to provide shareable content. In iOS, if both of those properties are
// [nil], Siri uses the [NSUserActivity.WebpageURL] property of your app’s current user
// activity as a fallback value.
// 
// # Quick Note
// 
// Quick Note on macOS and iOS can link to any app content represented as an
// [NSUserActivity]. To appear as a link, the content must be the app’s
// current activity, and provide at least one of the following identifiers:
// 
// [NSUserActivity.WebpageURL]: An `` URL, ideally in a canonical form that’s consistent
// every time a person visits the same content. [NSUserActivity.PersistentIdentifier]: A
// string that uniquely identifies the content in this domain. The identifier
// should identify the same content across devices. [NSUserActivity.TargetContentIdentifier]:
// A string that uniquely identifies the content in this domain, but also
// allows disambiguating between multiple scenes of an app. The identifier
// should identify the same content across devices.
// 
// To work well with Quick Note, content must adhere to the following
// guidelines:
// 
// - The activity [NSUserActivity.Title] should be clear and concise. This text describes the
// content of the link, like “Photo taken on July 27, 2020” or
// “Conversation with Maria”. Use nouns for activity titles. - Keep the
// app’s current activity up to date, using [NSUserActivity.BecomeCurrent] and
// [NSUserActivity.ResignCurrent]. - Linkable identifiers (listed above) must be stable and
// consistent for the same content. When you link from a note to a document in
// an app, and later revisit that document, the system shows an indicator
// linking back to the note. The system compares identifiers to check that the
// document is the same as the original source of the link. - Maintain support
// for activities provided by your app, and support navigating to linked
// content indefinitely. Links added to notes are important to people, who may
// feel that a broken link indicates data loss. - Gracefully handle attempts
// to navigate to an activity that points to content that doesn’t exist. For
// example, you can redirect to the new location of moved content, or show an
// error message. This situation may happen with shared notes, when a person
// links to content that exists only on another person’s device.
// 
// # Search results
// 
// If your [NSUserActivity] objects contain information that a person might
// want to search for later, set the [NSUserActivity.EligibleForSearch] property to [true].
// When you enable search, Spotlight indexes your user activity objects and
// considers them during subsequent on-device searches. For example, if a
// person viewed information about a particular restaurant in your app,
// you’d enable search for the corresponding user activity object.
// Subsequent searches for restaurants using Spotlight could then include the
// results obtained from your user activity object.
// 
// In addition to on-device searches, you can contribute URLs accessed by your
// app with the global Spotlight search engine. Sharing a URL helps Spotlight
// improve its own search results for other people. To contribute a URL, put
// the URL in the [NSUserActivity.WebpageURL] property of your activity object and set the
// [NSUserActivity.EligibleForPublicIndexing] property to [true].
// 
// Employ user activity objects to record user-initiated activities, not as a
// general-purpose indexing mechanism of your app’s data. To index all of
// your app’s content, and not just the content touched by people, use the
// APIs of the [Core Spotlight] framework.
//
// [Core Spotlight]: https://developer.apple.com/documentation/CoreSpotlight
// [INInteraction]: https://developer.apple.com/documentation/Intents/INInteraction
// [Implementing Handoff in Your App]: https://developer.apple.com/documentation/Foundation/implementing-handoff-in-your-app
// [NSSharingServicePickerToolbarItem]: https://developer.apple.com/documentation/AppKit/NSSharingServicePickerToolbarItem
// [nil]: https://developer.apple.com/documentation/ObjectiveC/nil-227m0
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Creating a user activity object
//
//   - [NSUserActivity.InitWithActivityType]: Creates a user activity object with the specified type.
//
// # Accessing activity information
//
//   - [NSUserActivity.ActivityType]: The user activity object’s activity type.
//   - [NSUserActivity.Title]: An optional, user-visible title for this activity, such as a document name or web page title.
//   - [NSUserActivity.SetTitle]
//   - [NSUserActivity.RequiredUserInfoKeys]: A set of keys that represent the minimal information about the activity that should be stored for later restoration.
//   - [NSUserActivity.SetRequiredUserInfoKeys]
//   - [NSUserActivity.UserInfo]: A dictionary containing app-specific state information needed to continue an activity on another device.
//   - [NSUserActivity.SetUserInfo]
//   - [NSUserActivity.AddUserInfoEntriesFromDictionary]: Adds the contents of the specified dictionary to the user info dictionary.
//   - [NSUserActivity.TargetContentIdentifier]: A string that identifies the user activity’s content.
//   - [NSUserActivity.SetTargetContentIdentifier]
//   - [NSUserActivity.NeedsSave]: A Boolean value that indicates whether the state of the activity needs to be updated.
//   - [NSUserActivity.SetNeedsSave]
//   - [NSUserActivity.ContentAttributeSet]: A set of properties that describe the activity.
//   - [NSUserActivity.SetContentAttributeSet]
//   - [NSUserActivity.Keywords]: A set of localized keywords that can help users find the activity in search results.
//   - [NSUserActivity.SetKeywords]
//   - [NSUserActivity.PersistentIdentifier]: A value used to identify the user activity.
//   - [NSUserActivity.SetPersistentIdentifier]
//
// # Specifying the activity’s eligibility
//
//   - [NSUserActivity.EligibleForHandoff]: A Boolean value that indicates whether the activity can be continued on another device using Handoff.
//   - [NSUserActivity.SetEligibleForHandoff]
//   - [NSUserActivity.EligibleForSearch]: A Boolean value that indicates whether the activity should be added to the on-device index.
//   - [NSUserActivity.SetEligibleForSearch]
//   - [NSUserActivity.EligibleForPublicIndexing]: A Boolean value that indicates whether the activity can be publicly accessed by all iOS users.
//   - [NSUserActivity.SetEligibleForPublicIndexing]
//   - [NSUserActivity.ExpirationDate]: The date after which the activity is no longer eligible for Handoff or indexing.
//   - [NSUserActivity.SetExpirationDate]
//
// # Registering and invalidating user activities
//
//   - [NSUserActivity.BecomeCurrent]: Marks the activity as currently in use by the user.
//   - [NSUserActivity.ResignCurrent]: Marks this activity object as inactive without invalidating it.
//   - [NSUserActivity.Invalidate]: Invalidates an activity and marks it as no longer eligible for continuation.
//
// # Accessing the delegate
//
//   - [NSUserActivity.Delegate]: The user activity object’s delegate.
//   - [NSUserActivity.SetDelegate]
//
// # Working with continuation streams
//
//   - [NSUserActivity.SupportsContinuationStreams]: A Boolean value that determines whether the continuing app can request streams to be opened back to the originating app.
//   - [NSUserActivity.SetSupportsContinuationStreams]
//   - [NSUserActivity.GetContinuationStreamsWithCompletionHandler]: Requests streams back to the originating app.
//
// # Continuing web browsing
//
//   - [NSUserActivity.WebpageURL]: The URL of the webpage to load in a browser to continue the activity.
//   - [NSUserActivity.SetWebpageURL]
//   - [NSUserActivity.ReferrerURL]: The URL of the webpage that linked to the webpage URL.
//   - [NSUserActivity.SetReferrerURL]
//
// # Donating to Siri Shortcuts
//
//   - [NSUserActivity.SuggestedInvocationPhrase]: A phrase suggested to the user when they create a shortcut.
//   - [NSUserActivity.SetSuggestedInvocationPhrase]
//
// # Continuing Siri interactions
//
//   - [NSUserActivity.Interaction]: The SiriKit interaction object to use when configuring your app.
//
// # Processing barcodes
//
//   - [NSUserActivity.DetectedBarcodeDescriptor]: The barcode that the system scanner passes in.
//
// # Sharing map item information
//
//   - [NSUserActivity.MapItem]: Attaches the specified map item to a user activity object.
//   - [NSUserActivity.SetMapItem]
//
// # Working with ClassKit
//
//   - [NSUserActivity.IsClassKitDeepLink]: A Boolean value that indicates whether a user activity represents a ClassKit context.
//   - [NSUserActivity.ContextIdentifierPath]: The identifier path associated with a user activity generated by an app that adopts ClassKit.
//
// # Identifying activity types
//
//   - [NSUserActivity.NSUserActivityTypeBrowsingWeb]: An activity that continues from Handoff or a universal link.
//   - [NSUserActivity.TVUserActivityTypeBrowsingChannelGuide]: An activity for viewing your app’s channel guide.
//
// # Reporting errors
//
//   - [NSUserActivity.NSUserActivityConnectionUnavailableError]: The user activity couldn’t be continued because a required connection wasn’t available.
//   - [NSUserActivity.SetNSUserActivityConnectionUnavailableError]
//   - [NSUserActivity.NSUserActivityErrorMaximum]: The end of the range of error codes reserved for user activity errors.
//   - [NSUserActivity.SetNSUserActivityErrorMaximum]
//   - [NSUserActivity.NSUserActivityErrorMinimum]: The start of the range of error codes reserved for user activity errors.
//   - [NSUserActivity.SetNSUserActivityErrorMinimum]
//   - [NSUserActivity.NSUserActivityHandoffFailedError]: The data for the user activity wasn’t available.
//   - [NSUserActivity.SetNSUserActivityHandoffFailedError]
//   - [NSUserActivity.NSUserActivityHandoffUserInfoTooLargeError]: The user info dictionary was too large to receive.
//   - [NSUserActivity.SetNSUserActivityHandoffUserInfoTooLargeError]
//   - [NSUserActivity.NSUserActivityRemoteApplicationTimedOutError]: The remote application failed to send data within the specified time.
//   - [NSUserActivity.SetNSUserActivityRemoteApplicationTimedOutError]
//
// # Instance Properties
//
//   - [NSUserActivity.AppEntityIdentifier]: The identifier of an app entity that you associate with the user activity.
//   - [NSUserActivity.SetAppEntityIdentifier]
//
// See: https://developer.apple.com/documentation/Foundation/NSUserActivity
type NSUserActivity struct {
	objectivec.Object
}

// NSUserActivityFromID constructs a [NSUserActivity] from an objc.ID.
//
// A representation of the state of your app at a moment in time.
func NSUserActivityFromID(id objc.ID) NSUserActivity {
	return NSUserActivity{objectivec.Object{id}}
}
// NOTE: NSUserActivity adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [NSUserActivity] class.
//
// # Creating a user activity object
//
//   - [INSUserActivity.InitWithActivityType]: Creates a user activity object with the specified type.
//
// # Accessing activity information
//
//   - [INSUserActivity.ActivityType]: The user activity object’s activity type.
//   - [INSUserActivity.Title]: An optional, user-visible title for this activity, such as a document name or web page title.
//   - [INSUserActivity.SetTitle]
//   - [INSUserActivity.RequiredUserInfoKeys]: A set of keys that represent the minimal information about the activity that should be stored for later restoration.
//   - [INSUserActivity.SetRequiredUserInfoKeys]
//   - [INSUserActivity.UserInfo]: A dictionary containing app-specific state information needed to continue an activity on another device.
//   - [INSUserActivity.SetUserInfo]
//   - [INSUserActivity.AddUserInfoEntriesFromDictionary]: Adds the contents of the specified dictionary to the user info dictionary.
//   - [INSUserActivity.TargetContentIdentifier]: A string that identifies the user activity’s content.
//   - [INSUserActivity.SetTargetContentIdentifier]
//   - [INSUserActivity.NeedsSave]: A Boolean value that indicates whether the state of the activity needs to be updated.
//   - [INSUserActivity.SetNeedsSave]
//   - [INSUserActivity.ContentAttributeSet]: A set of properties that describe the activity.
//   - [INSUserActivity.SetContentAttributeSet]
//   - [INSUserActivity.Keywords]: A set of localized keywords that can help users find the activity in search results.
//   - [INSUserActivity.SetKeywords]
//   - [INSUserActivity.PersistentIdentifier]: A value used to identify the user activity.
//   - [INSUserActivity.SetPersistentIdentifier]
//
// # Specifying the activity’s eligibility
//
//   - [INSUserActivity.EligibleForHandoff]: A Boolean value that indicates whether the activity can be continued on another device using Handoff.
//   - [INSUserActivity.SetEligibleForHandoff]
//   - [INSUserActivity.EligibleForSearch]: A Boolean value that indicates whether the activity should be added to the on-device index.
//   - [INSUserActivity.SetEligibleForSearch]
//   - [INSUserActivity.EligibleForPublicIndexing]: A Boolean value that indicates whether the activity can be publicly accessed by all iOS users.
//   - [INSUserActivity.SetEligibleForPublicIndexing]
//   - [INSUserActivity.ExpirationDate]: The date after which the activity is no longer eligible for Handoff or indexing.
//   - [INSUserActivity.SetExpirationDate]
//
// # Registering and invalidating user activities
//
//   - [INSUserActivity.BecomeCurrent]: Marks the activity as currently in use by the user.
//   - [INSUserActivity.ResignCurrent]: Marks this activity object as inactive without invalidating it.
//   - [INSUserActivity.Invalidate]: Invalidates an activity and marks it as no longer eligible for continuation.
//
// # Accessing the delegate
//
//   - [INSUserActivity.Delegate]: The user activity object’s delegate.
//   - [INSUserActivity.SetDelegate]
//
// # Working with continuation streams
//
//   - [INSUserActivity.SupportsContinuationStreams]: A Boolean value that determines whether the continuing app can request streams to be opened back to the originating app.
//   - [INSUserActivity.SetSupportsContinuationStreams]
//   - [INSUserActivity.GetContinuationStreamsWithCompletionHandler]: Requests streams back to the originating app.
//
// # Continuing web browsing
//
//   - [INSUserActivity.WebpageURL]: The URL of the webpage to load in a browser to continue the activity.
//   - [INSUserActivity.SetWebpageURL]
//   - [INSUserActivity.ReferrerURL]: The URL of the webpage that linked to the webpage URL.
//   - [INSUserActivity.SetReferrerURL]
//
// # Donating to Siri Shortcuts
//
//   - [INSUserActivity.SuggestedInvocationPhrase]: A phrase suggested to the user when they create a shortcut.
//   - [INSUserActivity.SetSuggestedInvocationPhrase]
//
// # Continuing Siri interactions
//
//   - [INSUserActivity.Interaction]: The SiriKit interaction object to use when configuring your app.
//
// # Processing barcodes
//
//   - [INSUserActivity.DetectedBarcodeDescriptor]: The barcode that the system scanner passes in.
//
// # Sharing map item information
//
//   - [INSUserActivity.MapItem]: Attaches the specified map item to a user activity object.
//   - [INSUserActivity.SetMapItem]
//
// # Working with ClassKit
//
//   - [INSUserActivity.IsClassKitDeepLink]: A Boolean value that indicates whether a user activity represents a ClassKit context.
//   - [INSUserActivity.ContextIdentifierPath]: The identifier path associated with a user activity generated by an app that adopts ClassKit.
//
// # Identifying activity types
//
//   - [INSUserActivity.NSUserActivityTypeBrowsingWeb]: An activity that continues from Handoff or a universal link.
//   - [INSUserActivity.TVUserActivityTypeBrowsingChannelGuide]: An activity for viewing your app’s channel guide.
//
// # Reporting errors
//
//   - [INSUserActivity.NSUserActivityConnectionUnavailableError]: The user activity couldn’t be continued because a required connection wasn’t available.
//   - [INSUserActivity.SetNSUserActivityConnectionUnavailableError]
//   - [INSUserActivity.NSUserActivityErrorMaximum]: The end of the range of error codes reserved for user activity errors.
//   - [INSUserActivity.SetNSUserActivityErrorMaximum]
//   - [INSUserActivity.NSUserActivityErrorMinimum]: The start of the range of error codes reserved for user activity errors.
//   - [INSUserActivity.SetNSUserActivityErrorMinimum]
//   - [INSUserActivity.NSUserActivityHandoffFailedError]: The data for the user activity wasn’t available.
//   - [INSUserActivity.SetNSUserActivityHandoffFailedError]
//   - [INSUserActivity.NSUserActivityHandoffUserInfoTooLargeError]: The user info dictionary was too large to receive.
//   - [INSUserActivity.SetNSUserActivityHandoffUserInfoTooLargeError]
//   - [INSUserActivity.NSUserActivityRemoteApplicationTimedOutError]: The remote application failed to send data within the specified time.
//   - [INSUserActivity.SetNSUserActivityRemoteApplicationTimedOutError]
//
// # Instance Properties
//
//   - [INSUserActivity.AppEntityIdentifier]: The identifier of an app entity that you associate with the user activity.
//   - [INSUserActivity.SetAppEntityIdentifier]
//
// See: https://developer.apple.com/documentation/Foundation/NSUserActivity
type INSUserActivity interface {
	objectivec.IObject
	NSItemProviderReading
	NSItemProviderWriting

	// Topic: Creating a user activity object

	// Creates a user activity object with the specified type.
	InitWithActivityType(activityType string) NSUserActivity

	// Topic: Accessing activity information

	// The user activity object’s activity type.
	ActivityType() string
	// An optional, user-visible title for this activity, such as a document name or web page title.
	Title() string
	SetTitle(value string)
	// A set of keys that represent the minimal information about the activity that should be stored for later restoration.
	RequiredUserInfoKeys() INSSet
	SetRequiredUserInfoKeys(value INSSet)
	// A dictionary containing app-specific state information needed to continue an activity on another device.
	UserInfo() INSDictionary
	SetUserInfo(value INSDictionary)
	// Adds the contents of the specified dictionary to the user info dictionary.
	AddUserInfoEntriesFromDictionary(otherDictionary INSDictionary)
	// A string that identifies the user activity’s content.
	TargetContentIdentifier() string
	SetTargetContentIdentifier(value string)
	// A Boolean value that indicates whether the state of the activity needs to be updated.
	NeedsSave() bool
	SetNeedsSave(value bool)
	// A set of properties that describe the activity.
	ContentAttributeSet() objectivec.IObject
	SetContentAttributeSet(value objectivec.IObject)
	// A set of localized keywords that can help users find the activity in search results.
	Keywords() INSSet
	SetKeywords(value INSSet)
	// A value used to identify the user activity.
	PersistentIdentifier() NSUserActivityPersistentIdentifier
	SetPersistentIdentifier(value NSUserActivityPersistentIdentifier)

	// Topic: Specifying the activity’s eligibility

	// A Boolean value that indicates whether the activity can be continued on another device using Handoff.
	EligibleForHandoff() bool
	SetEligibleForHandoff(value bool)
	// A Boolean value that indicates whether the activity should be added to the on-device index.
	EligibleForSearch() bool
	SetEligibleForSearch(value bool)
	// A Boolean value that indicates whether the activity can be publicly accessed by all iOS users.
	EligibleForPublicIndexing() bool
	SetEligibleForPublicIndexing(value bool)
	// The date after which the activity is no longer eligible for Handoff or indexing.
	ExpirationDate() INSDate
	SetExpirationDate(value INSDate)

	// Topic: Registering and invalidating user activities

	// Marks the activity as currently in use by the user.
	BecomeCurrent()
	// Marks this activity object as inactive without invalidating it.
	ResignCurrent()
	// Invalidates an activity and marks it as no longer eligible for continuation.
	Invalidate()

	// Topic: Accessing the delegate

	// The user activity object’s delegate.
	Delegate() NSUserActivityDelegate
	SetDelegate(value NSUserActivityDelegate)

	// Topic: Working with continuation streams

	// A Boolean value that determines whether the continuing app can request streams to be opened back to the originating app.
	SupportsContinuationStreams() bool
	SetSupportsContinuationStreams(value bool)
	// Requests streams back to the originating app.
	GetContinuationStreamsWithCompletionHandler(completionHandler InputStreamOutputStreamErrorHandler)

	// Topic: Continuing web browsing

	// The URL of the webpage to load in a browser to continue the activity.
	WebpageURL() INSURL
	SetWebpageURL(value INSURL)
	// The URL of the webpage that linked to the webpage URL.
	ReferrerURL() INSURL
	SetReferrerURL(value INSURL)

	// Topic: Donating to Siri Shortcuts

	// A phrase suggested to the user when they create a shortcut.
	SuggestedInvocationPhrase() string
	SetSuggestedInvocationPhrase(value string)

	// Topic: Continuing Siri interactions

	// The SiriKit interaction object to use when configuring your app.
	Interaction() objectivec.IObject

	// Topic: Processing barcodes

	// The barcode that the system scanner passes in.
	DetectedBarcodeDescriptor() objectivec.IObject

	// Topic: Sharing map item information

	// Attaches the specified map item to a user activity object.
	MapItem() objectivec.IObject
	SetMapItem(value objectivec.IObject)

	// Topic: Working with ClassKit

	// A Boolean value that indicates whether a user activity represents a ClassKit context.
	IsClassKitDeepLink() bool
	// The identifier path associated with a user activity generated by an app that adopts ClassKit.
	ContextIdentifierPath() []string

	// Topic: Identifying activity types

	// An activity that continues from Handoff or a universal link.
	NSUserActivityTypeBrowsingWeb() string
	// An activity for viewing your app’s channel guide.
	TVUserActivityTypeBrowsingChannelGuide() string

	// Topic: Reporting errors

	// The user activity couldn’t be continued because a required connection wasn’t available.
	NSUserActivityConnectionUnavailableError() int
	SetNSUserActivityConnectionUnavailableError(value int)
	// The end of the range of error codes reserved for user activity errors.
	NSUserActivityErrorMaximum() int
	SetNSUserActivityErrorMaximum(value int)
	// The start of the range of error codes reserved for user activity errors.
	NSUserActivityErrorMinimum() int
	SetNSUserActivityErrorMinimum(value int)
	// The data for the user activity wasn’t available.
	NSUserActivityHandoffFailedError() int
	SetNSUserActivityHandoffFailedError(value int)
	// The user info dictionary was too large to receive.
	NSUserActivityHandoffUserInfoTooLargeError() int
	SetNSUserActivityHandoffUserInfoTooLargeError(value int)
	// The remote application failed to send data within the specified time.
	NSUserActivityRemoteApplicationTimedOutError() int
	SetNSUserActivityRemoteApplicationTimedOutError(value int)

	// Topic: Instance Properties

	// The identifier of an app entity that you associate with the user activity.
	AppEntityIdentifier() string
	SetAppEntityIdentifier(value string)

	// An object or value that specifies items to share.
	ActivityItemsConfiguration() objectivec.IObject
	SetActivityItemsConfiguration(value objectivec.IObject)
	// An object that can provide shareable items for a scene.
	ActivityItemsConfigurationSource() objectivec.IObject
	SetActivityItemsConfigurationSource(value objectivec.IObject)
}





// Init initializes the instance.
func (u NSUserActivity) Init() NSUserActivity {
	rv := objc.Send[NSUserActivity](u.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u NSUserActivity) Autorelease() NSUserActivity {
	rv := objc.Send[NSUserActivity](u.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSUserActivity creates a new NSUserActivity instance.
func NewNSUserActivity() NSUserActivity {
	class := getNSUserActivityClass()
	rv := objc.Send[NSUserActivity](objc.ID(class.class), objc.Sel("new"))
	return rv
}






// Creates a user activity object with the specified type.
//
// activityType: The type of the activity. The value is a developer-defined string in
// reverse-DNS format by convention, for example,
// `com.MyCompanyXCUIElementTypeMyEditorXCUIElementTypeEditing()`.
//
// # Return Value
// 
// An [NSUserActivity] object.
//
// See: https://developer.apple.com/documentation/Foundation/NSUserActivity/init(activityType:)
func NewUserActivityWithActivityType(activityType string) NSUserActivity {
	instance := getNSUserActivityClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithActivityType:"), objc.String(activityType))
	return NSUserActivityFromID(rv)
}







// Creates a user activity object with the specified type.
//
// activityType: The type of the activity. The value is a developer-defined string in
// reverse-DNS format by convention, for example,
// `com.MyCompanyXCUIElementTypeMyEditorXCUIElementTypeEditing()`.
//
// # Return Value
// 
// An [NSUserActivity] object.
//
// See: https://developer.apple.com/documentation/Foundation/NSUserActivity/init(activityType:)
func (u NSUserActivity) InitWithActivityType(activityType string) NSUserActivity {
	rv := objc.Send[NSUserActivity](u.ID, objc.Sel("initWithActivityType:"), objc.String(activityType))
	return rv
}

// Adds the contents of the specified dictionary to the user info dictionary.
//
// otherDictionary: The dictionary containing entries to be added.
//
// # Discussion
// 
// Use this method to add the keys from `otherDictionary` into the dictionary
// in the [UserInfo] property. If the same key is in both dictionaries, the
// value of the key is set to the value in the `otherDictionary` parameter.
// 
// It’s recommended that you keep the [UserInfo] dictionary as small as
// possible. The larger the dictionary, the longer it takes to deliver that
// payload and resume the activity.
//
// See: https://developer.apple.com/documentation/Foundation/NSUserActivity/addUserInfoEntries(from:)
func (u NSUserActivity) AddUserInfoEntriesFromDictionary(otherDictionary INSDictionary) {
	objc.Send[objc.ID](u.ID, objc.Sel("addUserInfoEntriesFromDictionary:"), otherDictionary)
}

// Marks the activity as currently in use by the user.
//
// # Discussion
// 
// Call this method to let the system know that the user is performing the
// associated activity. The system makes this object the current user activity
// object, which makes it available for Handoff and search indexing. If
// another user activity object was previously active, that object is made
// inactive.
// 
// Don’t call this method when providing a user activity object for a Siri
// request. Siri holds on to user activity objects and passes them along to
// your app automatically in response to specific events.
// 
// If you previously called the [Invalidate] method on the current object,
// calling this method has no effect.
//
// See: https://developer.apple.com/documentation/Foundation/NSUserActivity/becomeCurrent()
func (u NSUserActivity) BecomeCurrent() {
	objc.Send[objc.ID](u.ID, objc.Sel("becomeCurrent"))
}

// Marks this activity object as inactive without invalidating it.
//
// # Discussion
// 
// Calling this method marks the user activity as no longer current, but
// doesn’t invalidate it entirely. You can call this method when you want to
// stop advertising the activity for continuation and search indexing only
// temporarily. You may call [BecomeCurrent] later to restore this object as
// the current activity.
//
// See: https://developer.apple.com/documentation/Foundation/NSUserActivity/resignCurrent()
func (u NSUserActivity) ResignCurrent() {
	objc.Send[objc.ID](u.ID, objc.Sel("resignCurrent"))
}

// Invalidates an activity and marks it as no longer eligible for
// continuation.
//
// # Discussion
// 
// Call this method when the user stops engaging in the associated activity
// and that activity is no longer available. For example, you might call this
// method when the user closes the window associated with the activity. After
// calling this method on a user activity object, calling the [BecomeCurrent]
// method on that object has no effect.
//
// See: https://developer.apple.com/documentation/Foundation/NSUserActivity/invalidate()
func (u NSUserActivity) Invalidate() {
	objc.Send[objc.ID](u.ID, objc.Sel("invalidate"))
}

// Requests streams back to the originating app.
//
// completionHandler: The completion handler block that returns streams.
// 
// The block takes three arguments:
// 
// `inputStream`: The stream from which the continuing app can read data
// written by the originating app. `outputStream`: The stream to which the
// continuing app writes data to be read by the originating app. `error`: If
// successful, `nil`; if not successful, an [NSError] object that encapsulates
// the reason why the streams could not be created.
//
// # Discussion
// 
// When an app is launched for a continuation event, it can request streams
// back to the originating app. Streams can be successfully retrieved only
// from the [NSUserActivity] object in the [NSApplication] or [UIApplication]
// delegate that is called for a continuation event. The streams are provided
// by the completion handler in an unopened state, and the delegate should
// open them immediately to start communicating with the continuing side.
// 
// Continuation streams are an optional feature of Handoff, and most user
// activities do not need them for successful continuation. When streams are
// needed, a simple request from the continuing app accompanied by a response
// from the originating app is enough for most continuation events.
//
// [NSApplication]: https://developer.apple.com/documentation/AppKit/NSApplication
// [UIApplication]: https://developer.apple.com/documentation/UIKit/UIApplication
//
// See: https://developer.apple.com/documentation/Foundation/NSUserActivity/getContinuationStreams(completionHandler:)
func (u NSUserActivity) GetContinuationStreamsWithCompletionHandler(completionHandler InputStreamOutputStreamErrorHandler) {
		_block0, _cleanup0 := NewInputStreamOutputStreamErrorBlock(completionHandler)
	defer _cleanup0()
		objc.Send[objc.ID](u.ID, objc.Sel("getContinuationStreamsWithCompletionHandler:"), _block0)
}

// Asks the item provider for the representation visibility specification for
// the given UTI.
//
// typeIdentifier: A uniform type identifier (UTI).
//
// # Return Value
// 
// A representation visibility specification for the given UTI.
//
// See: https://developer.apple.com/documentation/Foundation/NSItemProviderWriting/itemProviderVisibilityForRepresentation(withTypeIdentifier:)-swift.method
func (u NSUserActivity) ItemProviderVisibilityForRepresentationWithTypeIdentifier(typeIdentifier string) NSItemProviderRepresentationVisibility {
	rv := objc.Send[NSItemProviderRepresentationVisibility](u.ID, objc.Sel("itemProviderVisibilityForRepresentationWithTypeIdentifier:"), objc.String(typeIdentifier))
	return NSItemProviderRepresentationVisibility(rv)
}

// Loads data of a particular type, identified by the given UTI.
//
// typeIdentifier: The uniform type identifier (UTI) identifying the type of data to load.
//
// completionHandler: The handler that’s called after the data is loaded.
//
// # Return Value
// 
// The progress of the data load process.
//
// # Discussion
// 
// When the system calls this method, the `typeIdentifier` parameter is set to
// one of the elements in the `writableTypeIdentifiersForItemProvider` array.
//
// See: https://developer.apple.com/documentation/Foundation/NSItemProviderWriting/loadData(withTypeIdentifier:forItemProviderCompletionHandler:)
func (u NSUserActivity) LoadDataWithTypeIdentifierForItemProviderCompletionHandler(typeIdentifier string, completionHandler DataErrorHandler) INSProgress {
		_block1, _cleanup1 := NewDataErrorBlock(completionHandler)
	defer _cleanup1()
		rv := objc.Send[objc.ID](u.ID, objc.Sel("loadDataWithTypeIdentifier:forItemProviderCompletionHandler:"), objc.String(typeIdentifier), _block1)
	return NSProgressFromID(rv)
}





// Deletes all user activities created by your app.
//
// handler: The block that the system invokes after deleting the user activities. Wait
// for the system to call this block to ensure that the system deletes the
// activities (or marks them for deletion).
//
// # Discussion
// 
// Deletes all user activities stored by Core Spotlight or donated as Siri
// shortcuts.
//
// See: https://developer.apple.com/documentation/Foundation/NSUserActivity/deleteAllSavedUserActivities(completionHandler:)
func (_NSUserActivityClass NSUserActivityClass) DeleteAllSavedUserActivitiesWithCompletionHandler(handler VoidHandler) {
		_block0, _cleanup0 := NewVoidBlock(handler)
	defer _cleanup0()
		objc.Send[objc.ID](objc.ID(_NSUserActivityClass.class), objc.Sel("deleteAllSavedUserActivitiesWithCompletionHandler:"), _block0)
}

// Deletes user activities created by your app that have the specified
// persistent identifiers.
//
// persistentIdentifiers: The list of persistent identifiers that the system uses to determine which
// user activities to delete.
//
// handler: The block that the system invokes after deleting the user activities. Wait
// for the system to call this block to ensure that the system deletes the
// activities (or marks them for deletion).
//
// # Discussion
// 
// Deletes user activities with a persistent identifier matching any
// identifier in the `persistentIdentifiers` array.
//
// See: https://developer.apple.com/documentation/Foundation/NSUserActivity/deleteSavedUserActivities(withPersistentIdentifiers:completionHandler:)
func (_NSUserActivityClass NSUserActivityClass) DeleteSavedUserActivitiesWithPersistentIdentifiersCompletionHandler(persistentIdentifiers []string, handler VoidHandler) {
		_block1, _cleanup1 := NewVoidBlock(handler)
	defer _cleanup1()
		objc.Send[objc.ID](objc.ID(_NSUserActivityClass.class), objc.Sel("deleteSavedUserActivitiesWithPersistentIdentifiers:completionHandler:"), persistentIdentifiers, _block1)
}

// Creates a new instance of a class using the given data and UTI string.
//
// data: The data used to create the object.
//
// typeIdentifier: The uniform type identifier (UTI) representing the data type of `data`.
//
// # Return Value
// 
// An object created from the given data.
//
// See: https://developer.apple.com/documentation/Foundation/NSItemProviderReading/object(withItemProviderData:typeIdentifier:)
func (_NSUserActivityClass NSUserActivityClass) ObjectWithItemProviderDataTypeIdentifierError(data INSData, typeIdentifier string) (NSUserActivity, error) {
			var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_NSUserActivityClass.class), objc.Sel("objectWithItemProviderData:typeIdentifier:error:"), data, objc.String(typeIdentifier), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return NSUserActivity{}, NSErrorFrom(errorPtr)
	}
	return NSUserActivityFromID(rv), nil

}








// The user activity object’s activity type.
//
// # Discussion
// 
// This property is set at initialization time and can’t be changed later.
// Typically, you specify activity type strings using a reverse-DNS format
// that uniquely identifies the activity.
//
// See: https://developer.apple.com/documentation/Foundation/NSUserActivity/activityType
func (u NSUserActivity) ActivityType() string {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("activityType"))
	return NSStringFromID(rv).String()
}



// An optional, user-visible title for this activity, such as a document name
// or web page title.
//
// # Discussion
// 
// Always specify a title string for activity objects that are eligible for
// searches, and it’s recommended that you include a title string for all
// user activity objects. For search-related user activity objects, this
// string is displayed in the search results.
//
// See: https://developer.apple.com/documentation/Foundation/NSUserActivity/title
func (u NSUserActivity) Title() string {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("title"))
	return NSStringFromID(rv).String()
}
func (u NSUserActivity) SetTitle(value string) {
	objc.Send[struct{}](u.ID, objc.Sel("setTitle:"), objc.String(value))
}



// A set of keys that represent the minimal information about the activity
// that should be stored for later restoration.
//
// # Discussion
// 
// The keys come from the [UserInfo] property.
//
// See: https://developer.apple.com/documentation/Foundation/NSUserActivity/requiredUserInfoKeys
func (u NSUserActivity) RequiredUserInfoKeys() INSSet {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("requiredUserInfoKeys"))
	return NSSetFromID(objc.ID(rv))
}
func (u NSUserActivity) SetRequiredUserInfoKeys(value INSSet) {
	objc.Send[struct{}](u.ID, objc.Sel("setRequiredUserInfoKeys:"), value)
}



// A dictionary containing app-specific state information needed to continue
// an activity on another device.
//
// # Discussion
// 
// Each key and value must be of the following types: [NSArray], [NSData],
// [NSDate], [NSDictionary], [NSNull], [NSNumber], [NSSet], [NSString], or
// [NSURL]. The system may translate file scheme URLs that refer to iCloud
// documents to valid file URLs on a continuing device.
//
// See: https://developer.apple.com/documentation/Foundation/NSUserActivity/userInfo
func (u NSUserActivity) UserInfo() INSDictionary {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("userInfo"))
	return NSDictionaryFromID(objc.ID(rv))
}
func (u NSUserActivity) SetUserInfo(value INSDictionary) {
	objc.Send[struct{}](u.ID, objc.Sel("setUserInfo:"), value)
}



// A string that identifies the user activity’s content.
//
// # Discussion
// 
// A target content identifier is a string you define within your app. This
// string provides a unique identifier for specific content in your app, like
// a particular document or the location of a piece of data in a database.
// This string isn’t visible to the user.
// 
// If you set this property, when the system delivers an [NSUserActivity]
// object to an app with multiple scenes, it chooses the [UIScene] whose
// [UISceneActivationConditions] have the best match with the target content
// identifier. For more information, see [UISceneActivationConditions].
// 
// This property is optional but is highly recommended to create a great
// multitasking experience for apps that run on iPad. Setting this property
// doesn’t automatically set [NeedsSave] to [true].
//
// [UISceneActivationConditions]: https://developer.apple.com/documentation/UIKit/UISceneActivationConditions
// [UIScene]: https://developer.apple.com/documentation/UIKit/UIScene
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Foundation/NSUserActivity/targetContentIdentifier
func (u NSUserActivity) TargetContentIdentifier() string {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("targetContentIdentifier"))
	return NSStringFromID(rv).String()
}
func (u NSUserActivity) SetTargetContentIdentifier(value string) {
	objc.Send[struct{}](u.ID, objc.Sel("setTargetContentIdentifier:"), objc.String(value))
}



// A Boolean value that indicates whether the state of the activity needs to
// be updated.
//
// # Discussion
// 
// If [true], the delegate for this user activity receives a
// [UserActivityWillSave] callback before the activity is sent for
// continuation on another device.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Foundation/NSUserActivity/needsSave
func (u NSUserActivity) NeedsSave() bool {
	rv := objc.Send[bool](u.ID, objc.Sel("needsSave"))
	return rv
}
func (u NSUserActivity) SetNeedsSave(value bool) {
	objc.Send[struct{}](u.ID, objc.Sel("setNeedsSave:"), value)
}



// A set of properties that describe the activity.
//
// # Discussion
// 
// A [CSSearchableItemAttributeSet] object encapsulates the set of properties
// you want to display for a searchable activity.
//
// [CSSearchableItemAttributeSet]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet
//
// See: https://developer.apple.com/documentation/Foundation/NSUserActivity/contentAttributeSet
func (u NSUserActivity) ContentAttributeSet() objectivec.IObject {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("contentAttributeSet"))
	return objectivec.Object{ID: rv}
}
func (u NSUserActivity) SetContentAttributeSet(value objectivec.IObject) {
	objc.Send[struct{}](u.ID, objc.Sel("setContentAttributeSet:"), value)
}



// A set of localized keywords that can help users find the activity in search
// results.
//
// # Discussion
// 
// The default value of this property is `nil`. The system indexes the
// keywords you provide.
//
// See: https://developer.apple.com/documentation/Foundation/NSUserActivity/keywords
func (u NSUserActivity) Keywords() INSSet {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("keywords"))
	return NSSetFromID(objc.ID(rv))
}
func (u NSUserActivity) SetKeywords(value INSSet) {
	objc.Send[struct{}](u.ID, objc.Sel("setKeywords:"), value)
}



// A value used to identify the user activity.
//
// # Discussion
// 
// Set this property to a value that identifies the user activity so you can
// later delete it with
// [DeleteSavedUserActivitiesWithPersistentIdentifiersCompletionHandler]. For
// example, if the user checks the weather for Cupertino each morning from
// home, the weather app sets the persistent identifier to the city name
// (Cupertino). When the user deletes Cupertino from the weather app, the app
// deletes the user activity associated with the identifier, “Cupertino”.
//
// See: https://developer.apple.com/documentation/Foundation/NSUserActivity/persistentIdentifier
func (u NSUserActivity) PersistentIdentifier() NSUserActivityPersistentIdentifier {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("persistentIdentifier"))
	return NSUserActivityPersistentIdentifier(NSStringFromID(rv).String())
}
func (u NSUserActivity) SetPersistentIdentifier(value NSUserActivityPersistentIdentifier) {
	objc.Send[struct{}](u.ID, objc.Sel("setPersistentIdentifier:"), objc.String(string(value)))
}



// A Boolean value that indicates whether the activity can be continued on
// another device using Handoff.
//
// # Discussion
// 
// The default value of this property is [true].
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Foundation/NSUserActivity/isEligibleForHandoff
func (u NSUserActivity) EligibleForHandoff() bool {
	rv := objc.Send[bool](u.ID, objc.Sel("isEligibleForHandoff"))
	return rv
}
func (u NSUserActivity) SetEligibleForHandoff(value bool) {
	objc.Send[struct{}](u.ID, objc.Sel("setEligibleForHandoff:"), value)
}



// A Boolean value that indicates whether the activity should be added to the
// on-device index.
//
// # Discussion
// 
// The default value of this property is [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
//
// See: https://developer.apple.com/documentation/Foundation/NSUserActivity/isEligibleForSearch
func (u NSUserActivity) EligibleForSearch() bool {
	rv := objc.Send[bool](u.ID, objc.Sel("isEligibleForSearch"))
	return rv
}
func (u NSUserActivity) SetEligibleForSearch(value bool) {
	objc.Send[struct{}](u.ID, objc.Sel("setEligibleForSearch:"), value)
}



// A Boolean value that indicates whether the activity can be publicly
// accessed by all iOS users.
//
// # Discussion
// 
// The default value of this property is [false], which indicates that the
// activity object contains private or sensitive information or that the
// activity isn’t useful to other users. When the value of this property is
// [true], the system identifies this activity as one that can be shared
// publicly. When you make an activity public, the system indexes the values
// in the [RequiredUserInfoKeys] or [WebpageURL] properties, and you must
// provide a value for one of those properties.
// 
// Identifying an activity as public confers an advantage when you also add
// web markup to the content on your related website. Specifically, when users
// engage with your app’s public activities in search results, it indicates
// to Apple that public information on your website is popular, which can help
// increase your ranking and potentially lead to expanded indexing of your
// website’s content.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Foundation/NSUserActivity/isEligibleForPublicIndexing
func (u NSUserActivity) EligibleForPublicIndexing() bool {
	rv := objc.Send[bool](u.ID, objc.Sel("isEligibleForPublicIndexing"))
	return rv
}
func (u NSUserActivity) SetEligibleForPublicIndexing(value bool) {
	objc.Send[struct{}](u.ID, objc.Sel("setEligibleForPublicIndexing:"), value)
}



// The date after which the activity is no longer eligible for Handoff or
// indexing.
//
// # Discussion
// 
// If you don’t set the value of this property, the system automatically
// expires the activity after a period of time.
//
// See: https://developer.apple.com/documentation/Foundation/NSUserActivity/expirationDate
func (u NSUserActivity) ExpirationDate() INSDate {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("expirationDate"))
	return NSDateFromID(objc.ID(rv))
}
func (u NSUserActivity) SetExpirationDate(value INSDate) {
	objc.Send[struct{}](u.ID, objc.Sel("setExpirationDate:"), value)
}



// The user activity object’s delegate.
//
// # Discussion
// 
// The user activity delegate is informed when the activity is being saved or
// continued. For more information on how to implement the delegate, see
// [NSUserActivityDelegate].
//
// See: https://developer.apple.com/documentation/Foundation/NSUserActivity/delegate
func (u NSUserActivity) Delegate() NSUserActivityDelegate {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("delegate"))
	return NSUserActivityDelegateObjectFromID(rv)
}
func (u NSUserActivity) SetDelegate(value NSUserActivityDelegate) {
	objc.Send[struct{}](u.ID, objc.Sel("setDelegate:"), value)
}



// A Boolean value that determines whether the continuing app can request
// streams to be opened back to the originating app.
//
// # Discussion
// 
// If the value of this property is [true], the continuing app can connect
// back to the originating app for more information using streams. The default
// value of this property is [false]. It can dynamically be set to [true] to
// selectively support continuation streams based on the state of the user
// activity.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Foundation/NSUserActivity/supportsContinuationStreams
func (u NSUserActivity) SupportsContinuationStreams() bool {
	rv := objc.Send[bool](u.ID, objc.Sel("supportsContinuationStreams"))
	return rv
}
func (u NSUserActivity) SetSupportsContinuationStreams(value bool) {
	objc.Send[struct{}](u.ID, objc.Sel("setSupportsContinuationStreams:"), value)
}



// The URL of the webpage to load in a browser to continue the activity.
//
// # Discussion
// 
// When no suitable app is installed on a resuming device and the [WebpageURL]
// property is set, the specified webpage is loaded and the user activity is
// continued in a web browser.
// 
// If your activity’s content can be restored on the web or you support
// Safari universal links, be sure to set this property so that the system can
// resume the activity in Safari or your app. After setting the [WebpageURL]
// property on an activity for which [EligibleForSearch] is [true], also set
// the [RequiredUserInfoKeys] property, using the keys of the [UserInfo]
// dictionary that must be stored. If you don’t also set the
// [RequiredUserInfoKeys] property, the [UserInfo] dictionary will be empty
// when the activity is restored.
// 
// If [EligibleForSearch] is [true] for this activity and you’re using both
// [NSUserActivity] and web markup to index the same item, set [WebpageURL] to
// the relevant URL on your website to avoid showing duplicate results in
// Spotlight. The [NSUserActivity] API does not perform any modifications to
// the URL that you specify. URL components, such as the query string and the
// fragment identifier, are used for matching the item against pages that are
// indexed by Applebot.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Foundation/NSUserActivity/webpageURL
func (u NSUserActivity) WebpageURL() INSURL {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("webpageURL"))
	return NSURLFromID(objc.ID(rv))
}
func (u NSUserActivity) SetWebpageURL(value INSURL) {
	objc.Send[struct{}](u.ID, objc.Sel("setWebpageURL:"), value)
}



// The URL of the webpage that linked to the webpage URL.
//
// See: https://developer.apple.com/documentation/Foundation/NSUserActivity/referrerURL
func (u NSUserActivity) ReferrerURL() INSURL {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("referrerURL"))
	return NSURLFromID(objc.ID(rv))
}
func (u NSUserActivity) SetReferrerURL(value INSURL) {
	objc.Send[struct{}](u.ID, objc.Sel("setReferrerURL:"), value)
}



// A phrase suggested to the user when they create a shortcut.
//
// # Discussion
// 
// The system displays the suggested invocation phrase to the user when they
// create the shortcut. Use a short, memorable phrase, such as “Soup
// time”.
// 
// [media-3020431]
//
// See: https://developer.apple.com/documentation/Foundation/NSUserActivity/suggestedInvocationPhrase
func (u NSUserActivity) SuggestedInvocationPhrase() string {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("suggestedInvocationPhrase"))
	return NSStringFromID(rv).String()
}
func (u NSUserActivity) SetSuggestedInvocationPhrase(value string) {
	objc.Send[struct{}](u.ID, objc.Sel("setSuggestedInvocationPhrase:"), objc.String(value))
}



// The SiriKit interaction object to use when configuring your app.
//
// # Discussion
// 
// When SiriKit launches your app, it fills this property with the intent and
// response information that are the reason for launching your app. Use the
// information in this property to configure your app’s interface and show
// any relevant interaction details. If your app wasn’t launched because of
// a Siri interaction, the value in this property is `nil`.
//
// See: https://developer.apple.com/documentation/Foundation/NSUserActivity/interaction
func (u NSUserActivity) Interaction() objectivec.IObject {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("interaction"))
	return objectivec.Object{ID: rv}
}



// The barcode that the system scanner passes in.
//
// # Discussion
// 
// This property is optional. This value is present if the user activity was
// created from a source that detected a barcode or a QR code.
//
// See: https://developer.apple.com/documentation/Foundation/NSUserActivity/detectedBarcodeDescriptor
func (u NSUserActivity) DetectedBarcodeDescriptor() objectivec.IObject {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("detectedBarcodeDescriptor"))
	return objectivec.Object{ID: rv}
}



// Attaches the specified map item to a user activity object.
//
// # Discussion
// 
// Use this property to share a map item’s location information with other
// apps so that users can benefit from a more contextually integrated
// experience. For example, if a user views a location in an app that provides
// restaurant reviews, this property can make the same location available to
// the user when they switch to an app that helps them make travel plans.
// 
// Setting the [MapItem] property also populates the user activity object’s
// [ContentAttributeSet] property. After attaching a map item to a user
// activity object, you can easily adopt app search by setting the object’s
// [EligibleForSearch] property to [true]. To learn more about participating
// in app search, see [App Search Programming Guide].
//
// [App Search Programming Guide]: https://developer.apple.com/library/archive/documentation/General/Conceptual/AppSearch/index.html#//apple_ref/doc/uid/TP40016308
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Foundation/NSUserActivity/mapItem
func (u NSUserActivity) MapItem() objectivec.IObject {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("mapItem"))
	return objectivec.Object{ID: rv}
}
func (u NSUserActivity) SetMapItem(value objectivec.IObject) {
	objc.Send[struct{}](u.ID, objc.Sel("setMapItem:"), value)
}



// A Boolean value that indicates whether a user activity represents a
// ClassKit context.
//
// # Discussion
// 
// When a student taps on an assignment associated with an app that adopts
// ClassKit, the framework redirects the student to the corresponding app
// using either a universal link or a user activity, depending on the app’s
// configuration. Use the [IsClassKitDeepLink] property of a [NSUserActivity]
// instance that you receive to test if that activity is from ClassKit. See
// [Linking directly to assignments] for more information.
//
// [Linking directly to assignments]: https://developer.apple.com/documentation/ClassKit/linking-directly-to-assignments
//
// See: https://developer.apple.com/documentation/Foundation/NSUserActivity/isClassKitDeepLink
func (u NSUserActivity) IsClassKitDeepLink() bool {
	rv := objc.Send[bool](u.ID, objc.Sel("isClassKitDeepLink"))
	return rv
}



// The identifier path associated with a user activity generated by an app
// that adopts ClassKit.
//
// # Discussion
// 
// If you receive an [NSUserActivity] instance that has its
// [IsClassKitDeepLink] property set to [true], then the activity’s
// [ContextIdentifierPath] contains the identifier path of the context
// associated with the assignment that the user tapped to generate the user
// activity. See [Linking directly to assignments] for more information.
//
// [Linking directly to assignments]: https://developer.apple.com/documentation/ClassKit/linking-directly-to-assignments
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Foundation/NSUserActivity/contextIdentifierPath
func (u NSUserActivity) ContextIdentifierPath() []string {
	rv := objc.Send[[]objc.ID](u.ID, objc.Sel("contextIdentifierPath"))
	return objc.ConvertSliceToStrings(rv)
}



// An activity that continues from Handoff or a universal link.
//
// See: https://developer.apple.com/documentation/foundation/nsuseractivitytypebrowsingweb
func (u NSUserActivity) NSUserActivityTypeBrowsingWeb() string {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("NSUserActivityTypeBrowsingWeb"))
	return NSStringFromID(rv).String()
}



// An activity for viewing your app’s channel guide.
//
// See: https://developer.apple.com/documentation/TVServices/TVUserActivityTypeBrowsingChannelGuide
func (u NSUserActivity) TVUserActivityTypeBrowsingChannelGuide() string {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("TVUserActivityTypeBrowsingChannelGuide"))
	return NSStringFromID(rv).String()
}



// The user activity couldn’t be continued because a required connection
// wasn’t available.
//
// See: https://developer.apple.com/documentation/foundation/nsuseractivityconnectionunavailableerror-swift.var
func (u NSUserActivity) NSUserActivityConnectionUnavailableError() int {
	rv := objc.Send[int](u.ID, objc.Sel("NSUserActivityConnectionUnavailableError"))
	return rv
}
func (u NSUserActivity) SetNSUserActivityConnectionUnavailableError(value int) {
	objc.Send[struct{}](u.ID, objc.Sel("setNSUserActivityConnectionUnavailableError:"), value)
}



// The end of the range of error codes reserved for user activity errors.
//
// See: https://developer.apple.com/documentation/foundation/nsuseractivityerrormaximum-swift.var
func (u NSUserActivity) NSUserActivityErrorMaximum() int {
	rv := objc.Send[int](u.ID, objc.Sel("NSUserActivityErrorMaximum"))
	return rv
}
func (u NSUserActivity) SetNSUserActivityErrorMaximum(value int) {
	objc.Send[struct{}](u.ID, objc.Sel("setNSUserActivityErrorMaximum:"), value)
}



// The start of the range of error codes reserved for user activity errors.
//
// See: https://developer.apple.com/documentation/foundation/nsuseractivityerrorminimum-swift.var
func (u NSUserActivity) NSUserActivityErrorMinimum() int {
	rv := objc.Send[int](u.ID, objc.Sel("NSUserActivityErrorMinimum"))
	return rv
}
func (u NSUserActivity) SetNSUserActivityErrorMinimum(value int) {
	objc.Send[struct{}](u.ID, objc.Sel("setNSUserActivityErrorMinimum:"), value)
}



// The data for the user activity wasn’t available.
//
// See: https://developer.apple.com/documentation/foundation/nsuseractivityhandofffailederror-swift.var
func (u NSUserActivity) NSUserActivityHandoffFailedError() int {
	rv := objc.Send[int](u.ID, objc.Sel("NSUserActivityHandoffFailedError"))
	return rv
}
func (u NSUserActivity) SetNSUserActivityHandoffFailedError(value int) {
	objc.Send[struct{}](u.ID, objc.Sel("setNSUserActivityHandoffFailedError:"), value)
}



// The user info dictionary was too large to receive.
//
// See: https://developer.apple.com/documentation/foundation/nsuseractivityhandoffuserinfotoolargeerror-swift.var
func (u NSUserActivity) NSUserActivityHandoffUserInfoTooLargeError() int {
	rv := objc.Send[int](u.ID, objc.Sel("NSUserActivityHandoffUserInfoTooLargeError"))
	return rv
}
func (u NSUserActivity) SetNSUserActivityHandoffUserInfoTooLargeError(value int) {
	objc.Send[struct{}](u.ID, objc.Sel("setNSUserActivityHandoffUserInfoTooLargeError:"), value)
}



// The remote application failed to send data within the specified time.
//
// See: https://developer.apple.com/documentation/foundation/nsuseractivityremoteapplicationtimedouterror-swift.var
func (u NSUserActivity) NSUserActivityRemoteApplicationTimedOutError() int {
	rv := objc.Send[int](u.ID, objc.Sel("NSUserActivityRemoteApplicationTimedOutError"))
	return rv
}
func (u NSUserActivity) SetNSUserActivityRemoteApplicationTimedOutError(value int) {
	objc.Send[struct{}](u.ID, objc.Sel("setNSUserActivityRemoteApplicationTimedOutError:"), value)
}



// The identifier of an app entity that you associate with the user activity.
//
// See: https://developer.apple.com/documentation/foundation/nsuseractivity/appentityidentifier
func (u NSUserActivity) AppEntityIdentifier() string {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("appEntityIdentifier"))
	return NSStringFromID(rv).String()
}
func (u NSUserActivity) SetAppEntityIdentifier(value string) {
	objc.Send[struct{}](u.ID, objc.Sel("setAppEntityIdentifier:"), objc.String(value))
}



// An object or value that specifies items to share.
//
// See: https://developer.apple.com/documentation/UIKit/UIActivityItemsConfigurationProviding/activityItemsConfiguration
func (u NSUserActivity) ActivityItemsConfiguration() objectivec.IObject {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("activityItemsConfiguration"))
	return objectivec.Object{ID: rv}
}
func (u NSUserActivity) SetActivityItemsConfiguration(value objectivec.IObject) {
	objc.Send[struct{}](u.ID, objc.Sel("setActivityItemsConfiguration:"), value)
}



// An object that can provide shareable items for a scene.
//
// See: https://developer.apple.com/documentation/UIKit/UIWindowScene/activityItemsConfigurationSource
func (u NSUserActivity) ActivityItemsConfigurationSource() objectivec.IObject {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("activityItemsConfigurationSource"))
	return objectivec.Object{ID: rv}
}
func (u NSUserActivity) SetActivityItemsConfigurationSource(value objectivec.IObject) {
	objc.Send[struct{}](u.ID, objc.Sel("setActivityItemsConfigurationSource:"), value)
}



// An array of UTI strings representing the types of data that can be loaded
// for an item provider.
//
// # Discussion
// 
// Provide uniform type identifiers (UTIs) in order from highest fidelity to
// lowest. If your app employs a native data representation, place that first
// in the array.
// 
// Use the instance version of this property when you initialize an item
// provider with an object. As possible, implement this property to provide an
// extended array of UTIs based on the object. For example, for an [NSURL]
// object, your implementation could offer the `public.File()-url` UTI, in
// addition to the `public.Url()` UTI, if your implementation detects that the
// stored URL uses the `//` scheme.
//
// See: https://developer.apple.com/documentation/Foundation/NSItemProviderWriting/writableTypeIdentifiersForItemProvider-swift.property
func (u NSUserActivity) WritableTypeIdentifiersForItemProvider() []string {
	rv := objc.Send[[]objc.ID](u.ID, objc.Sel("writableTypeIdentifiersForItemProvider"))
	return objc.ConvertSliceToStrings(rv)
}
















			// Protocol methods for NSItemProviderReading
			




			// Protocol methods for NSItemProviderWriting
			









// DeleteAllSavedUserActivities is a synchronous wrapper around [NSUserActivity.DeleteAllSavedUserActivitiesWithCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (uc NSUserActivityClass) DeleteAllSavedUserActivities(ctx context.Context) error {
	done := make(chan struct{}, 1)
	uc.DeleteAllSavedUserActivitiesWithCompletionHandler(func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// LoadDataWithTypeIdentifierForItemProvider is a synchronous wrapper around [NSUserActivity.LoadDataWithTypeIdentifierForItemProviderCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (u NSUserActivity) LoadDataWithTypeIdentifierForItemProvider(ctx context.Context, typeIdentifier string) (*NSData, error) {
	type result struct {
		val *NSData
		err error
	}
	done := make(chan result, 1)
	u.LoadDataWithTypeIdentifierForItemProviderCompletionHandler(typeIdentifier, func(val *NSData, err error) {
		done <- result{val, err}
	})
	select {
	case r := <-done:
		return r.val, r.err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}





