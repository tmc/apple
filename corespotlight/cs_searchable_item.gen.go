// Code generated from Apple documentation for CoreSpotlight. DO NOT EDIT.

package corespotlight

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CSSearchableItem] class.
var (
	_CSSearchableItemClass     CSSearchableItemClass
	_CSSearchableItemClassOnce sync.Once
)

func getCSSearchableItemClass() CSSearchableItemClass {
	_CSSearchableItemClassOnce.Do(func() {
		_CSSearchableItemClass = CSSearchableItemClass{class: objc.GetClass("CSSearchableItem")}
	})
	return _CSSearchableItemClass
}

// GetCSSearchableItemClass returns the class object for CSSearchableItem.
func GetCSSearchableItemClass() CSSearchableItemClass {
	return getCSSearchableItemClass()
}

type CSSearchableItemClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CSSearchableItemClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CSSearchableItemClass) Alloc() CSSearchableItem {
	rv := objc.Send[CSSearchableItem](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// The details of your app-specific content that someone might search for on
// their devices.
//
// # Overview
//
// A [CSSearchableItem] uniquely identifies a part of your app’s content,
// and provides the metadata that Spotlight indexes and uses to find that
// content later. As part of indexing your app’s content, you create
// searchable items and fill them with details about your app’s content and
// where to find it. After indexing the content, you can then execute queries
// using the Core Spotlight APIs to find the items you indexed. People can
// also use the system’s Spotlight search interface to find your app’s
// content.
//
// When you create or update content in your app, create a [CSSearchableItem]
// for that content if you want it to be searchable. A searchable item
// contains identification strings you use to locate that item in your content
// and a [CSSearchableItemAttributeSet] object with details about the item.
// For the metadata, you typically want to provide values for the
// [CSSearchableItemAttributeSet.Title],
// [CSSearchableItemAttributeSet.DisplayName], and
// [CSSearchableItemAttributeSet.ContentType] attributes at a minimum. If
// you’re indexing a file on disk, provide a value for the
// [CSSearchableItemAttributeSet.ContentURL] attribute. Fill in as many other
// attributes as makes sense for the content you’re indexing.
//
// After creating a searchable item, index it using a [CSSearchableIndex]
// object. As you update your app’s content, update your [CSSearchableItem]
// objects for that content and index them right away. If you delete content,
// similarly delete the searchable items from the index. Keeping your app’s
// indexes current ensures that searches return valid information. For more
// information on indexing your content, see [Adding your app’s content to
// Spotlight indexes].
//
// # Getting a searchable item
//
//   - [CSSearchableItem.InitWithUniqueIdentifierDomainIdentifierAttributeSet]: Returns a searchable item associated with the specified identifier, domain identifier, and attribute set.
//   - [CSSearchableItem.InitWithCoder]
//
// # Setting attributes on a searchable item
//
//   - [CSSearchableItem.UniqueIdentifier]: The value that uniquely identifies the searchable item within your app.
//   - [CSSearchableItem.SetUniqueIdentifier]
//   - [CSSearchableItem.DomainIdentifier]: An optional identifier that represents the domain or owner of the item.
//   - [CSSearchableItem.SetDomainIdentifier]
//   - [CSSearchableItem.AttributeSet]: The set of attributes that contain metadata associated with the item in a [CSSearchableItemAttributeSet](<https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet>) object.
//   - [CSSearchableItem.SetAttributeSet]
//   - [CSSearchableItem.ExpirationDate]: The date after which the searchable item should no longer exist.
//   - [CSSearchableItem.SetExpirationDate]
//   - [CSSearchableItem.IsUpdate]: A Boolean value that indicates whether to treat the item as an update instead of a new item.
//   - [CSSearchableItem.SetIsUpdate]
//   - [CSSearchableItem.UpdateListenerOptions]
//   - [CSSearchableItem.SetUpdateListenerOptions]
//
// # Comparing items
//
//   - [CSSearchableItem.CompareByRank]: Compares two items by rank and returns the result.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItem
//
// [Adding your app’s content to Spotlight indexes]: https://developer.apple.com/documentation/CoreSpotlight/adding-your-app-s-content-to-spotlight-indexes
type CSSearchableItem struct {
	objectivec.Object
}

// CSSearchableItemFromID constructs a [CSSearchableItem] from an objc.ID.
//
// The details of your app-specific content that someone might search for on
// their devices.
func CSSearchableItemFromID(id objc.ID) CSSearchableItem {
	return CSSearchableItem{objectivec.Object{ID: id}}
}

// NOTE: CSSearchableItem adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CSSearchableItem] class.
//
// # Getting a searchable item
//
//   - [ICSSearchableItem.InitWithUniqueIdentifierDomainIdentifierAttributeSet]: Returns a searchable item associated with the specified identifier, domain identifier, and attribute set.
//   - [ICSSearchableItem.InitWithCoder]
//
// # Setting attributes on a searchable item
//
//   - [ICSSearchableItem.UniqueIdentifier]: The value that uniquely identifies the searchable item within your app.
//   - [ICSSearchableItem.SetUniqueIdentifier]
//   - [ICSSearchableItem.DomainIdentifier]: An optional identifier that represents the domain or owner of the item.
//   - [ICSSearchableItem.SetDomainIdentifier]
//   - [ICSSearchableItem.AttributeSet]: The set of attributes that contain metadata associated with the item in a [CSSearchableItemAttributeSet](<https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet>) object.
//   - [ICSSearchableItem.SetAttributeSet]
//   - [ICSSearchableItem.ExpirationDate]: The date after which the searchable item should no longer exist.
//   - [ICSSearchableItem.SetExpirationDate]
//   - [ICSSearchableItem.IsUpdate]: A Boolean value that indicates whether to treat the item as an update instead of a new item.
//   - [ICSSearchableItem.SetIsUpdate]
//   - [ICSSearchableItem.UpdateListenerOptions]
//   - [ICSSearchableItem.SetUpdateListenerOptions]
//
// # Comparing items
//
//   - [ICSSearchableItem.CompareByRank]: Compares two items by rank and returns the result.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItem
type ICSSearchableItem interface {
	objectivec.IObject

	// Topic: Getting a searchable item

	// Returns a searchable item associated with the specified identifier, domain identifier, and attribute set.
	InitWithUniqueIdentifierDomainIdentifierAttributeSet(uniqueIdentifier string, domainIdentifier string, attributeSet ICSSearchableItemAttributeSet) CSSearchableItem
	InitWithCoder(coder foundation.INSCoder) CSSearchableItem

	// Topic: Setting attributes on a searchable item

	// The value that uniquely identifies the searchable item within your app.
	UniqueIdentifier() string
	SetUniqueIdentifier(value string)
	// An optional identifier that represents the domain or owner of the item.
	DomainIdentifier() string
	SetDomainIdentifier(value string)
	// The set of attributes that contain metadata associated with the item in a [CSSearchableItemAttributeSet](<https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet>) object.
	AttributeSet() ICSSearchableItemAttributeSet
	SetAttributeSet(value ICSSearchableItemAttributeSet)
	// The date after which the searchable item should no longer exist.
	ExpirationDate() foundation.NSDate
	SetExpirationDate(value foundation.NSDate)
	// A Boolean value that indicates whether to treat the item as an update instead of a new item.
	IsUpdate() bool
	SetIsUpdate(value bool)
	UpdateListenerOptions() CSSearchableItemUpdateListenerOptions
	SetUpdateListenerOptions(value CSSearchableItemUpdateListenerOptions)

	// Topic: Comparing items

	// Compares two items by rank and returns the result.
	CompareByRank(other ICSSearchableItem) foundation.ComparisonResult

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (c CSSearchableItem) Init() CSSearchableItem {
	rv := objc.Send[CSSearchableItem](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CSSearchableItem) Autorelease() CSSearchableItem {
	rv := objc.Send[CSSearchableItem](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCSSearchableItem creates a new CSSearchableItem instance.
func NewCSSearchableItem() CSSearchableItem {
	class := getCSSearchableItemClass()
	rv := objc.Send[CSSearchableItem](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItem/init(coder:)
func NewCSSearchableItemWithCoder(coder foundation.INSCoder) CSSearchableItem {
	instance := getCSSearchableItemClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return CSSearchableItemFromID(rv)
}

// Returns a searchable item associated with the specified identifier, domain
// identifier, and attribute set.
//
// uniqueIdentifier: The unique identifier for the item. If you specify [NULL], an identifier is
// generated automatically.
//
// domainIdentifier: An identifier for a domain, such as an album, that helps you group items
// together in a way that makes sense.
//
// attributeSet: A set of properties that specify the metadata you want to display about an
// item in a search result. See [CSSearchableItemAttributeSet] for the types
// of properties you can use.
//
// # Return Value
//
// A searchable item that’s associated with the specified identifier, domain
// identifier, and attribute set.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItem/init(uniqueIdentifier:domainIdentifier:attributeSet:)
func NewCSSearchableItemWithUniqueIdentifierDomainIdentifierAttributeSet(uniqueIdentifier string, domainIdentifier string, attributeSet ICSSearchableItemAttributeSet) CSSearchableItem {
	instance := getCSSearchableItemClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithUniqueIdentifier:domainIdentifier:attributeSet:"), objc.String(uniqueIdentifier), objc.String(domainIdentifier), attributeSet)
	return CSSearchableItemFromID(rv)
}

// Returns a searchable item associated with the specified identifier, domain
// identifier, and attribute set.
//
// uniqueIdentifier: The unique identifier for the item. If you specify [NULL], an identifier is
// generated automatically.
//
// domainIdentifier: An identifier for a domain, such as an album, that helps you group items
// together in a way that makes sense.
//
// attributeSet: A set of properties that specify the metadata you want to display about an
// item in a search result. See [CSSearchableItemAttributeSet] for the types
// of properties you can use.
//
// # Return Value
//
// A searchable item that’s associated with the specified identifier, domain
// identifier, and attribute set.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItem/init(uniqueIdentifier:domainIdentifier:attributeSet:)
func (c CSSearchableItem) InitWithUniqueIdentifierDomainIdentifierAttributeSet(uniqueIdentifier string, domainIdentifier string, attributeSet ICSSearchableItemAttributeSet) CSSearchableItem {
	rv := objc.Send[CSSearchableItem](c.ID, objc.Sel("initWithUniqueIdentifier:domainIdentifier:attributeSet:"), objc.String(uniqueIdentifier), objc.String(domainIdentifier), attributeSet)
	return rv
}

// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItem/init(coder:)
func (c CSSearchableItem) InitWithCoder(coder foundation.INSCoder) CSSearchableItem {
	rv := objc.Send[CSSearchableItem](c.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}

// Compares two items by rank and returns the result.
//
// other: The other item to compare against the current one.
//
// # Return Value
//
// A comparison result that indicates the ranked order of the items.
//
// # Discussion
//
// Call this function when you want to compare the current item with the one
// you specify. The method compares the ranks of the items.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItem/compare(byRank:)
func (c CSSearchableItem) CompareByRank(other ICSSearchableItem) foundation.ComparisonResult {
	rv := objc.Send[foundation.NSComparisonResult](c.ID, objc.Sel("compareByRank:"), other)
	return foundation.ComparisonResult(rv)
}
func (c CSSearchableItem) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](c.ID, objc.Sel("encodeWithCoder:"), coder)
}

// The value that uniquely identifies the searchable item within your app.
//
// # Discussion
//
// This property is required because it’s the only way to identify
// searchable items in the index when you need to access or delete them. When
// you create a searchable item, the system generates a UUID by default, but
// you can replace the default value with a unique identifier that makes sense
// in the context of your app. If you want to use a custom value for
// [CSSearchableItem.UniqueIdentifier], be sure to set it before the item is
// indexed for the first time.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItem/uniqueIdentifier
func (c CSSearchableItem) UniqueIdentifier() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("uniqueIdentifier"))
	return foundation.NSStringFromID(rv).String()
}
func (c CSSearchableItem) SetUniqueIdentifier(value string) {
	objc.Send[struct{}](c.ID, objc.Sel("setUniqueIdentifier:"), objc.String(value))
}

// An optional identifier that represents the domain or owner of the item.
//
// # Discussion
//
// Specify a domain identifier to group items together and to make it easy to
// delete groups of items from the index. For example, you might specify an
// identifier for a mailbox in an account whose indexed data you want to
// remove when the account is deleted. In this example,
// [CSSearchableItem.DomainIdentifier] should be of the form `.`, where
// neither “ nor “ contain periods. To delete all items associated with the
// specified account and mailbox, you can call
// [CSSearchableIndex.DeleteSearchableItemsWithDomainIdentifiersCompletionHandler]
// with a [CSSearchableItem.DomainIdentifier] of `.`. Or to delete all items
// associated with all mailboxes in the specified account, you can call
// [CSSearchableIndex.DeleteSearchableItemsWithDomainIdentifiersCompletionHandler]
// with a [CSSearchableItem.DomainIdentifier] of “.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItem/domainIdentifier
func (c CSSearchableItem) DomainIdentifier() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("domainIdentifier"))
	return foundation.NSStringFromID(rv).String()
}
func (c CSSearchableItem) SetDomainIdentifier(value string) {
	objc.Send[struct{}](c.ID, objc.Sel("setDomainIdentifier:"), objc.String(value))
}

// The set of attributes that contain metadata associated with the item in a
// [CSSearchableItemAttributeSet] object.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItem/attributeSet
func (c CSSearchableItem) AttributeSet() ICSSearchableItemAttributeSet {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("attributeSet"))
	return CSSearchableItemAttributeSetFromID(objc.ID(rv))
}
func (c CSSearchableItem) SetAttributeSet(value ICSSearchableItemAttributeSet) {
	objc.Send[struct{}](c.ID, objc.Sel("setAttributeSet:"), value)
}

// The date after which the searchable item should no longer exist.
//
// # Discussion
//
// If you don’t set the [CSSearchableItem.ExpirationDate] property
// appropriately, the system automatically expires the item after a period of
// time.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItem/expirationDate
func (c CSSearchableItem) ExpirationDate() foundation.NSDate {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("expirationDate"))
	return foundation.NSDateFromID(objc.ID(rv))
}
func (c CSSearchableItem) SetExpirationDate(value foundation.NSDate) {
	objc.Send[struct{}](c.ID, objc.Sel("setExpirationDate:"), value)
}

// A Boolean value that indicates whether to treat the item as an update
// instead of a new item.
//
// # Discussion
//
// Set the value of this property to `true` if the item represents an update
// to information already in the index. Marking an item as an update makes the
// indexing process more efficient. If this property is `false` and the system
// encounters an item with the same identifier in the index, it deletes the
// old item and then inserts the new one. When the property is `true`, it
// updates the existing item, which saves time. If this property is `true` and
// the item doesn’t exist in the index, the system ignores the request to
// index the item and doesn’t create a new item.
//
// When configuring the attributes for the item, set an attribute to `nil` to
// remove its value.
//
// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItem/isUpdate
func (c CSSearchableItem) IsUpdate() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isUpdate"))
	return rv
}
func (c CSSearchableItem) SetIsUpdate(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setIsUpdate:"), value)
}

// See: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItem/updateListenerOptions-swift.property
func (c CSSearchableItem) UpdateListenerOptions() CSSearchableItemUpdateListenerOptions {
	rv := objc.Send[CSSearchableItemUpdateListenerOptions](c.ID, objc.Sel("updateListenerOptions"))
	return CSSearchableItemUpdateListenerOptions(rv)
}
func (c CSSearchableItem) SetUpdateListenerOptions(value CSSearchableItemUpdateListenerOptions) {
	objc.Send[struct{}](c.ID, objc.Sel("setUpdateListenerOptions:"), value)
}
