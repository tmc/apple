// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CKUserIdentity] class.
var (
	_CKUserIdentityClass     CKUserIdentityClass
	_CKUserIdentityClassOnce sync.Once
)

func getCKUserIdentityClass() CKUserIdentityClass {
	_CKUserIdentityClassOnce.Do(func() {
		_CKUserIdentityClass = CKUserIdentityClass{class: objc.GetClass("CKUserIdentity")}
	})
	return _CKUserIdentityClass
}

// GetCKUserIdentityClass returns the class object for CKUserIdentity.
func GetCKUserIdentityClass() CKUserIdentityClass {
	return getCKUserIdentityClass()
}

type CKUserIdentityClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CKUserIdentityClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CKUserIdentityClass) Alloc() CKUserIdentity {
	rv := objc.Send[CKUserIdentity](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// The identity of a user.
//
// # Overview
//
// A user identity provides identifiable data about an iCloud user, including
// their name, user record ID, and an email address or phone number. CloudKit
// retrieves this information from the user’s iCloud account. A user must
// give their consent to be discoverable before CloudKit can provide this data
// to your app. For more information, see
// [CKContainer.RequestApplicationPermissionCompletionHandler].
//
// You don’t create instances of this class. Instead, CloudKit provides them
// in certain contexts. A share’s owner has a user identity, as does each of
// its participants. When creating participants, CloudKit tries to find iCloud
// accounts it can use to populate their identities. If CloudKit doesn’t
// find an account, it sets the identity’s [CKUserIdentity.HasiCloudAccount]
// property to false.
//
// You can also discover the identities of your app’s users by executing one
// of the user discovery operations: [CKDiscoverAllUserIdentitiesOperation]
// and [CKDiscoverUserIdentitiesOperation]. Identities that CloudKit discovers
// using [CKDiscoverAllUserIdentitiesOperation] correspond to entries in the
// device’s Contacts database. These identities contain the identifiers of
// their Contact records, which you can use to fetch those records from the
// Contacts database. For more information, see
// [CKUserIdentity.ContactIdentifiers].
//
// # Accessing iCloud Information
//
//   - [CKUserIdentity.HasiCloudAccount]: A Boolean value that indicates whether the user has an iCloud account.
//   - [CKUserIdentity.LookupInfo]: The lookup info for retrieving the user identity.
//
// # Accessing User Information
//
//   - [CKUserIdentity.UserRecordID]: The user record ID for the corresponding user record.
//   - [CKUserIdentity.ContactIdentifiers]: Identifiers that match contacts in the local Contacts database.
//   - [CKUserIdentity.NameComponents]: The user’s name.
//
// # Initializers
//
//   - [CKUserIdentity.InitWithCoder]
//
// See: https://developer.apple.com/documentation/CloudKit/CKUserIdentity
type CKUserIdentity struct {
	objectivec.Object
}

// CKUserIdentityFromID constructs a [CKUserIdentity] from an objc.ID.
//
// The identity of a user.
func CKUserIdentityFromID(id objc.ID) CKUserIdentity {
	return CKUserIdentity{objectivec.Object{ID: id}}
}

// NOTE: CKUserIdentity adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CKUserIdentity] class.
//
// # Accessing iCloud Information
//
//   - [ICKUserIdentity.HasiCloudAccount]: A Boolean value that indicates whether the user has an iCloud account.
//   - [ICKUserIdentity.LookupInfo]: The lookup info for retrieving the user identity.
//
// # Accessing User Information
//
//   - [ICKUserIdentity.UserRecordID]: The user record ID for the corresponding user record.
//   - [ICKUserIdentity.ContactIdentifiers]: Identifiers that match contacts in the local Contacts database.
//   - [ICKUserIdentity.NameComponents]: The user’s name.
//
// # Initializers
//
//   - [ICKUserIdentity.InitWithCoder]
//
// See: https://developer.apple.com/documentation/CloudKit/CKUserIdentity
type ICKUserIdentity interface {
	objectivec.IObject

	// Topic: Accessing iCloud Information

	// A Boolean value that indicates whether the user has an iCloud account.
	HasiCloudAccount() bool
	// The lookup info for retrieving the user identity.
	LookupInfo() ICKUserIdentityLookupInfo

	// Topic: Accessing User Information

	// The user record ID for the corresponding user record.
	UserRecordID() ICKRecordID
	// Identifiers that match contacts in the local Contacts database.
	ContactIdentifiers() []string
	// The user’s name.
	NameComponents() foundation.NSPersonNameComponents

	// Topic: Initializers

	InitWithCoder(coder foundation.INSCoder) CKUserIdentity

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (c CKUserIdentity) Init() CKUserIdentity {
	rv := objc.Send[CKUserIdentity](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CKUserIdentity) Autorelease() CKUserIdentity {
	rv := objc.Send[CKUserIdentity](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKUserIdentity creates a new CKUserIdentity instance.
func NewCKUserIdentity() CKUserIdentity {
	class := getCKUserIdentityClass()
	rv := objc.Send[CKUserIdentity](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CloudKit/CKUserIdentity/init(coder:)
func NewCKUserIdentityWithCoder(coder foundation.INSCoder) CKUserIdentity {
	instance := getCKUserIdentityClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return CKUserIdentityFromID(rv)
}

// See: https://developer.apple.com/documentation/CloudKit/CKUserIdentity/init(coder:)
func (c CKUserIdentity) InitWithCoder(coder foundation.INSCoder) CKUserIdentity {
	rv := objc.Send[CKUserIdentity](c.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}
func (c CKUserIdentity) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](c.ID, objc.Sel("encodeWithCoder:"), coder)
}

// A Boolean value that indicates whether the user has an iCloud account.
//
// # Discussion
//
// true if the user identity has an iCloud account; otherwise, false.
//
// See: https://developer.apple.com/documentation/CloudKit/CKUserIdentity/hasiCloudAccount
func (c CKUserIdentity) HasiCloudAccount() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("hasiCloudAccount"))
	return rv
}

// The lookup info for retrieving the user identity.
//
// # Discussion
//
// Use this property’s value to retrieve the user identity when using the
// [CKDiscoverUserIdentitiesOperation] and [CKFetchShareParticipantsOperation]
// operations.
//
// See: https://developer.apple.com/documentation/CloudKit/CKUserIdentity/lookupInfo-swift.property
func (c CKUserIdentity) LookupInfo() ICKUserIdentityLookupInfo {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("lookupInfo"))
	return CKUserIdentityLookupInfoFromID(objc.ID(rv))
}

// The user record ID for the corresponding user record.
//
// See: https://developer.apple.com/documentation/CloudKit/CKUserIdentity/userRecordID
func (c CKUserIdentity) UserRecordID() ICKRecordID {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("userRecordID"))
	return CKRecordIDFromID(objc.ID(rv))
}

// Identifiers that match contacts in the local Contacts database.
//
// # Discussion
//
// Identities that CloudKit discovers using
// [CKDiscoverAllUserIdentitiesOperation] correspond to entries in the local
// Contacts database, matching the identifier on [CNContact]. Use these
// identifiers with the Contacts database to get additional information about
// the contacts. Multiple identifiers can exist for a single discovered user
// because multiple contacts can contain the same email addresses or phone
// numbers.
//
// To transform these identifiers into an array of unified contact
// identifiers, create a predicate by calling the
// [predicateForContacts(withIdentifiers:)] method, and then pass that
// predicate to the [unifiedContacts(matching:keysToFetch:)] method.
//
// See: https://developer.apple.com/documentation/CloudKit/CKUserIdentity/contactIdentifiers
//
// [CNContact]: https://developer.apple.com/documentation/Contacts/CNContact
// [predicateForContacts(withIdentifiers:)]: https://developer.apple.com/documentation/Contacts/CNContact/predicateForContacts(withIdentifiers:)
// [unifiedContacts(matching:keysToFetch:)]: https://developer.apple.com/documentation/Contacts/CNContactStore/unifiedContacts(matching:keysToFetch:)
func (c CKUserIdentity) ContactIdentifiers() []string {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("contactIdentifiers"))
	return objc.ConvertSliceToStrings(rv)
}

// The user’s name.
//
// # Discussion
//
// You can use this property to construct the user’s name for display. Use
// the components with an instance of [PersonNameComponentsFormatter] to
// create a string representation for the current locale.
//
// See: https://developer.apple.com/documentation/CloudKit/CKUserIdentity/nameComponents
//
// [PersonNameComponentsFormatter]: https://developer.apple.com/documentation/Foundation/PersonNameComponentsFormatter
func (c CKUserIdentity) NameComponents() foundation.NSPersonNameComponents {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("nameComponents"))
	return foundation.NSPersonNameComponentsFromID(objc.ID(rv))
}
