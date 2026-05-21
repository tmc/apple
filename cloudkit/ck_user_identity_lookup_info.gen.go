// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CKUserIdentityLookupInfo] class.
var (
	_CKUserIdentityLookupInfoClass     CKUserIdentityLookupInfoClass
	_CKUserIdentityLookupInfoClassOnce sync.Once
)

func getCKUserIdentityLookupInfoClass() CKUserIdentityLookupInfoClass {
	_CKUserIdentityLookupInfoClassOnce.Do(func() {
		_CKUserIdentityLookupInfoClass = CKUserIdentityLookupInfoClass{class: objc.GetClass("CKUserIdentityLookupInfo")}
	})
	return _CKUserIdentityLookupInfoClass
}

// GetCKUserIdentityLookupInfoClass returns the class object for CKUserIdentityLookupInfo.
func GetCKUserIdentityLookupInfoClass() CKUserIdentityLookupInfoClass {
	return getCKUserIdentityLookupInfoClass()
}

type CKUserIdentityLookupInfoClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CKUserIdentityLookupInfoClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CKUserIdentityLookupInfoClass) Alloc() CKUserIdentityLookupInfo {
	rv := objc.Send[CKUserIdentityLookupInfo](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// The criteria to use when searching for discoverable iCloud users.
//
// # Overview
//
// Use this object when you want to discover the identities of your app’s
// users with [CKDiscoverUserIdentitiesOperation], or to create a share’s
// participants with [CKFetchShareParticipantsOperation].
//
// You create individual instances by providing an email address, phone
// number, or user record ID. Alternatively, create an array of objects all at
// once by using one of the convenience methods, such as
// [CKUserIdentityLookupInfoClass.LookupInfosWithEmails].
//
// # Creating a Lookup Info
//
//   - [CKUserIdentityLookupInfo.InitWithEmailAddress]: Creates a lookup info for the specified email address.
//   - [CKUserIdentityLookupInfo.InitWithPhoneNumber]: Creates a lookup info for the specified phone number.
//   - [CKUserIdentityLookupInfo.InitWithUserRecordID]: Creates a lookup info for the specified user record ID.
//
// # Accessing the Criteria
//
//   - [CKUserIdentityLookupInfo.EmailAddress]: The user’s email address.
//   - [CKUserIdentityLookupInfo.PhoneNumber]: The user’s phone number.
//   - [CKUserIdentityLookupInfo.UserRecordID]: The ID of the user record.
//
// See: https://developer.apple.com/documentation/CloudKit/CKUserIdentity/LookupInfo-swift.class
type CKUserIdentityLookupInfo struct {
	objectivec.Object
}

// CKUserIdentityLookupInfoFromID constructs a [CKUserIdentityLookupInfo] from an objc.ID.
//
// The criteria to use when searching for discoverable iCloud users.
func CKUserIdentityLookupInfoFromID(id objc.ID) CKUserIdentityLookupInfo {
	return CKUserIdentityLookupInfo{objectivec.Object{ID: id}}
}

// NOTE: CKUserIdentityLookupInfo adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CKUserIdentityLookupInfo] class.
//
// # Creating a Lookup Info
//
//   - [ICKUserIdentityLookupInfo.InitWithEmailAddress]: Creates a lookup info for the specified email address.
//   - [ICKUserIdentityLookupInfo.InitWithPhoneNumber]: Creates a lookup info for the specified phone number.
//   - [ICKUserIdentityLookupInfo.InitWithUserRecordID]: Creates a lookup info for the specified user record ID.
//
// # Accessing the Criteria
//
//   - [ICKUserIdentityLookupInfo.EmailAddress]: The user’s email address.
//   - [ICKUserIdentityLookupInfo.PhoneNumber]: The user’s phone number.
//   - [ICKUserIdentityLookupInfo.UserRecordID]: The ID of the user record.
//
// See: https://developer.apple.com/documentation/CloudKit/CKUserIdentity/LookupInfo-swift.class
type ICKUserIdentityLookupInfo interface {
	objectivec.IObject

	// Topic: Creating a Lookup Info

	// Creates a lookup info for the specified email address.
	InitWithEmailAddress(emailAddress string) CKUserIdentityLookupInfo
	// Creates a lookup info for the specified phone number.
	InitWithPhoneNumber(phoneNumber string) CKUserIdentityLookupInfo
	// Creates a lookup info for the specified user record ID.
	InitWithUserRecordID(userRecordID ICKRecordID) CKUserIdentityLookupInfo

	// Topic: Accessing the Criteria

	// The user’s email address.
	EmailAddress() string
	// The user’s phone number.
	PhoneNumber() string
	// The ID of the user record.
	UserRecordID() ICKRecordID

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (c CKUserIdentityLookupInfo) Init() CKUserIdentityLookupInfo {
	rv := objc.Send[CKUserIdentityLookupInfo](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CKUserIdentityLookupInfo) Autorelease() CKUserIdentityLookupInfo {
	rv := objc.Send[CKUserIdentityLookupInfo](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKUserIdentityLookupInfo creates a new CKUserIdentityLookupInfo instance.
func NewCKUserIdentityLookupInfo() CKUserIdentityLookupInfo {
	class := getCKUserIdentityLookupInfoClass()
	rv := objc.Send[CKUserIdentityLookupInfo](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a lookup info for the specified email address.
//
// emailAddress: The email address for looking up the user identity.
//
// # Discussion
//
// After you create a lookup info, use the [CKDiscoverUserIdentitiesOperation]
// operation or the [CKFetchShareParticipantsOperation] operation to retrieve
// the corresponding user identity.
//
// See: https://developer.apple.com/documentation/CloudKit/CKUserIdentity/LookupInfo-swift.class/init(emailAddress:)
func NewCKUserIdentityLookupInfoWithEmailAddress(emailAddress string) CKUserIdentityLookupInfo {
	instance := getCKUserIdentityLookupInfoClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithEmailAddress:"), objc.String(emailAddress))
	return CKUserIdentityLookupInfoFromID(rv)
}

// Creates a lookup info for the specified phone number.
//
// phoneNumber: The phone number for looking up the user identity.
//
// # Discussion
//
// After you create a lookup info, use the [CKDiscoverUserIdentitiesOperation]
// operation or the [CKFetchShareParticipantsOperation] operation to retrieve
// the corresponding user identity.
//
// See: https://developer.apple.com/documentation/CloudKit/CKUserIdentity/LookupInfo-swift.class/init(phoneNumber:)
func NewCKUserIdentityLookupInfoWithPhoneNumber(phoneNumber string) CKUserIdentityLookupInfo {
	instance := getCKUserIdentityLookupInfoClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithPhoneNumber:"), objc.String(phoneNumber))
	return CKUserIdentityLookupInfoFromID(rv)
}

// Creates a lookup info for the specified user record ID.
//
// userRecordID: The user record ID for looking up the user identity.
//
// # Discussion
//
// After you create a lookup info, use the [CKDiscoverUserIdentitiesOperation]
// operation or the [CKFetchShareParticipantsOperation] operation to retrieve
// the corresponding user identity.
//
// See: https://developer.apple.com/documentation/CloudKit/CKUserIdentity/LookupInfo-swift.class/init(userRecordID:)
func NewCKUserIdentityLookupInfoWithUserRecordID(userRecordID ICKRecordID) CKUserIdentityLookupInfo {
	instance := getCKUserIdentityLookupInfoClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithUserRecordID:"), userRecordID)
	return CKUserIdentityLookupInfoFromID(rv)
}

// Creates a lookup info for the specified email address.
//
// emailAddress: The email address for looking up the user identity.
//
// # Discussion
//
// After you create a lookup info, use the [CKDiscoverUserIdentitiesOperation]
// operation or the [CKFetchShareParticipantsOperation] operation to retrieve
// the corresponding user identity.
//
// See: https://developer.apple.com/documentation/CloudKit/CKUserIdentity/LookupInfo-swift.class/init(emailAddress:)
func (c CKUserIdentityLookupInfo) InitWithEmailAddress(emailAddress string) CKUserIdentityLookupInfo {
	rv := objc.Send[CKUserIdentityLookupInfo](c.ID, objc.Sel("initWithEmailAddress:"), objc.String(emailAddress))
	return rv
}

// Creates a lookup info for the specified phone number.
//
// phoneNumber: The phone number for looking up the user identity.
//
// # Discussion
//
// After you create a lookup info, use the [CKDiscoverUserIdentitiesOperation]
// operation or the [CKFetchShareParticipantsOperation] operation to retrieve
// the corresponding user identity.
//
// See: https://developer.apple.com/documentation/CloudKit/CKUserIdentity/LookupInfo-swift.class/init(phoneNumber:)
func (c CKUserIdentityLookupInfo) InitWithPhoneNumber(phoneNumber string) CKUserIdentityLookupInfo {
	rv := objc.Send[CKUserIdentityLookupInfo](c.ID, objc.Sel("initWithPhoneNumber:"), objc.String(phoneNumber))
	return rv
}

// Creates a lookup info for the specified user record ID.
//
// userRecordID: The user record ID for looking up the user identity.
//
// # Discussion
//
// After you create a lookup info, use the [CKDiscoverUserIdentitiesOperation]
// operation or the [CKFetchShareParticipantsOperation] operation to retrieve
// the corresponding user identity.
//
// See: https://developer.apple.com/documentation/CloudKit/CKUserIdentity/LookupInfo-swift.class/init(userRecordID:)
func (c CKUserIdentityLookupInfo) InitWithUserRecordID(userRecordID ICKRecordID) CKUserIdentityLookupInfo {
	rv := objc.Send[CKUserIdentityLookupInfo](c.ID, objc.Sel("initWithUserRecordID:"), userRecordID)
	return rv
}
func (c CKUserIdentityLookupInfo) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](c.ID, objc.Sel("encodeWithCoder:"), coder)
}

// Returns an array of lookup infos for the specifed email addresses.
//
// emails: The email addresses for looking up the user identities.
//
// # Discussion
//
// Use the values that this method returns in an
// [CKDiscoverUserIdentitiesOperation] operation or an
// [CKFetchShareParticipantsOperation] operation to retrieve the corresponding
// user identities.
//
// See: https://developer.apple.com/documentation/CloudKit/CKUserIdentity/LookupInfo-swift.class/lookupInfos(withEmails:)
func (_CKUserIdentityLookupInfoClass CKUserIdentityLookupInfoClass) LookupInfosWithEmails(emails []string) []CKUserIdentityLookupInfo {
	rv := objc.Send[[]objc.ID](objc.ID(_CKUserIdentityLookupInfoClass.class), objc.Sel("lookupInfosWithEmails:"), objectivec.StringSliceToNSArray(emails))
	return objc.ConvertSlice(rv, func(id objc.ID) CKUserIdentityLookupInfo {
		return CKUserIdentityLookupInfoFromID(id)
	})
}

// Returns an array of lookup infos for the specifed phone numbers.
//
// phoneNumbers: The phone numbers for looking up the user identities.
//
// # Discussion
//
// Use the values that this method returns in an
// [CKDiscoverUserIdentitiesOperation] operation or an
// [CKFetchShareParticipantsOperation] operation to retrieve the corresponding
// user identities.
//
// See: https://developer.apple.com/documentation/CloudKit/CKUserIdentity/LookupInfo-swift.class/lookupInfos(withPhoneNumbers:)
func (_CKUserIdentityLookupInfoClass CKUserIdentityLookupInfoClass) LookupInfosWithPhoneNumbers(phoneNumbers []string) []CKUserIdentityLookupInfo {
	rv := objc.Send[[]objc.ID](objc.ID(_CKUserIdentityLookupInfoClass.class), objc.Sel("lookupInfosWithPhoneNumbers:"), objectivec.StringSliceToNSArray(phoneNumbers))
	return objc.ConvertSlice(rv, func(id objc.ID) CKUserIdentityLookupInfo {
		return CKUserIdentityLookupInfoFromID(id)
	})
}

// Returns an array of lookup infos for the specifed user record IDs.
//
// recordIDs: The user record IDs for looking up the user identities.
//
// # Discussion
//
// Use the values that this method returns in an
// [CKDiscoverUserIdentitiesOperation] operation or an
// [CKFetchShareParticipantsOperation] operation to retrieve the corresponding
// user identities.
//
// See: https://developer.apple.com/documentation/CloudKit/CKUserIdentity/LookupInfo-swift.class/lookupInfos(with:)
func (_CKUserIdentityLookupInfoClass CKUserIdentityLookupInfoClass) LookupInfosWithRecordIDs(recordIDs []CKRecordID) []CKUserIdentityLookupInfo {
	rv := objc.Send[[]objc.ID](objc.ID(_CKUserIdentityLookupInfoClass.class), objc.Sel("lookupInfosWithRecordIDs:"), objectivec.IObjectSliceToNSArray(recordIDs))
	return objc.ConvertSlice(rv, func(id objc.ID) CKUserIdentityLookupInfo {
		return CKUserIdentityLookupInfoFromID(id)
	})
}

// The user’s email address.
//
// See: https://developer.apple.com/documentation/CloudKit/CKUserIdentity/LookupInfo-swift.class/emailAddress
func (c CKUserIdentityLookupInfo) EmailAddress() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("emailAddress"))
	return foundation.NSStringFromID(rv).String()
}

// The user’s phone number.
//
// See: https://developer.apple.com/documentation/CloudKit/CKUserIdentity/LookupInfo-swift.class/phoneNumber
func (c CKUserIdentityLookupInfo) PhoneNumber() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("phoneNumber"))
	return foundation.NSStringFromID(rv).String()
}

// The ID of the user record.
//
// # Discussion
//
// Use this value to retrieve the user record for the user identity. The user
// record doesn’t contain any personal information about the user, by
// default. You can add data to the user record, but you shouldn’t add
// anything sensitive.
//
// See: https://developer.apple.com/documentation/CloudKit/CKUserIdentity/LookupInfo-swift.class/userRecordID
func (c CKUserIdentityLookupInfo) UserRecordID() ICKRecordID {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("userRecordID"))
	return CKRecordIDFromID(objc.ID(rv))
}
