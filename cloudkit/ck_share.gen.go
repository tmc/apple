// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CKShare] class.
var (
	_CKShareClass     CKShareClass
	_CKShareClassOnce sync.Once
)

func getCKShareClass() CKShareClass {
	_CKShareClassOnce.Do(func() {
		_CKShareClass = CKShareClass{class: objc.GetClass("CKShare")}
	})
	return _CKShareClass
}

// GetCKShareClass returns the class object for CKShare.
func GetCKShareClass() CKShareClass {
	return getCKShareClass()
}

type CKShareClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CKShareClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CKShareClass) Alloc() CKShare {
	rv := objc.Send[CKShare](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// A specialized record type that manages a collection of shared records.
//
// # Overview
//
// A share is a specialized record type that facilitates the sharing of one or
// more records with many participants. You store shareable records in a
// custom record zone in the user’s private database. As you create records
// in that zone, they become eligible for record zone sharing. If you want to
// share a specific hierarchy of related records, rather than the entire
// record zone, set each record’s [CKShare.Parent] property to define the
// relationship with its parent. CloudKit infers the shared hierarchy using
// only the [CKShare.Parent] property, and ignores any custom reference fields.
//
// You create a share with either the ID of the record zone to share, or the
// root record, which defines the point in a record hierarchy where you want
// to start sharing. CloudKit shares all the records in the record zone, or
// every record in the hierarchy below the root. If you set the root
// record’s [CKShare.Parent] property, CloudKit ignores it. A record can take part
// in only a single share. This applies to every record in the shared record
// zone or hierarchy. If a record is participating in another share, any
// attempt to save the share fails, and CloudKit returns an [CKShare.AlreadyShared]
// error.
//
// Use [CKModifyRecordsOperation] to save the share to the server. The initial
// set of records the share includes must exist on the server or be part of
// the same save operation to succeed. CloudKit then updates the share’s
// [CKShare.URL] property. Use [UICloudSharingController] to present options to the
// user for sharing the URL. Otherwise, distribute the URL to any participants
// you add to the share. You can allow anyone with the URL to take part in the
// share by setting [CKShare.PublicPermission] to a value more permissive than
// [CKShare.ParticipantPermission.none].
//
// After CloudKit saves the share, a participant can fetch its corresponding
// metadata, which includes a reference to the share, information about the
// user’s participation, and, for shared hierarchies, the root record’s
// record ID. Create an instance of [CKFetchShareMetadataOperation] using the
// share’s URL and add it to the container’s queue to execute it. The
// operation returns an instance of [CKShareMetadata] for each URL you
// provide. This is only applicable if you manually process share acceptance.
// If a user receives the share URL and taps or clicks it, CloudKit
// automatically processes their participation.
//
// To determine the configuration of a fetched share, inspect the [CKShare.RecordName]
// property of its [CKShare.RecordID]. If the value is [CKShare.CKRecordNameZoneWideShare],
// the share is managing a shared record zone; otherwise, it’s managing a
// shared record hierarchy.
//
// CloudKit limits the number of participants in a share to 100, and each
// participant must have an active iCloud account. You don’t create
// participants. Instead, use [UICloudSharingController] to manage a share’s
// participants and their permissions. Alternatively, create an instance of
// [CKUserIdentityLookupInfo] for each user. Provide the user’s email
// address or phone number, and use [CKFetchShareParticipantsOperation] to
// fetch the corresponding participants. CloudKit queries iCloud for
// corresponding accounts as part of the operation. If it doesn’t find an
// account, the server updates the participant’s [CKShare.UserIdentity] to reflect
// that by setting the [CKShare.HasiCloudAccount] property to false. CloudKit
// associates the participant with their iCloud account when they accept the
// share if they launch the process by tapping or clicking the share URL.
//
// Participants with write permissions can modify or delete any record that
// you include in the share. However, only the owner can delete a shared
// hierarchy’s root record. If a participant attempts to delete the share,
// CloudKit removes the participant. The share remains active for all other
// participants. If the owner deletes a share that manages a record hierarchy,
// CloudKit sets the root record’s [CKShare.Share] property to `nil`. CloudKit
// deletes the share if the owner of the shared heirarchy deletes its root
// record.
//
// You can customize the title and image the system displays when initiating a
// share or accepting an invitation to participate. You can also provide a
// custom UTI to indicate the content of the shared records. Use the keys that
// [CKShare.SystemFieldKey] defines, as the following example shows:
//
// # Creating a Share
//
//   - [CKShare.InitWithRootRecord]: Creates a new share for the specified record.
//   - [CKShare.InitWithRootRecordShareID]: Creates a new share for the specified record and record ID.
//   - [CKShare.InitWithRecordZoneID]: Creates a new share for the specified record zone.
//
// # Accessing the Share’s Attributes
//
//   - [CKShare.Owner]: The participant that represents the share’s owner.
//   - [CKShare.CurrentUserParticipant]: The participant that represents the current user.
//   - [CKShare.Participants]: An array that contains the share’s participants.
//   - [CKShare.URL]: The URL for inviting participants to the share.
//
// # Configuring the Share
//
//   - [CKShare.PublicPermission]: The permission for anyone with access to the share’s URL.
//   - [CKShare.SetPublicPermission]
//   - [CKShare.AddParticipant]: Adds a participant to the share.
//   - [CKShare.RemoveParticipant]: Removes a participant from the share.
//
// # Instance Properties
//
//   - [CKShare.AllowsAccessRequests]: Indicates whether uninvited users can request access to this share.
//   - [CKShare.SetAllowsAccessRequests]
//   - [CKShare.BlockedIdentities]: A list of users blocked from requesting access to this share.
//   - [CKShare.Requesters]: A list of all uninvited users who have requested access to this share.
//
// # Instance Methods
//
//   - [CKShare.BlockRequesters]: Blocks specified users from requesting access to this share.
//   - [CKShare.DenyRequesters]: Denies access requests from specified users.
//   - [CKShare.UnblockIdentities]: Unblocks previously blocked users, allowing them to request access again.
//
// See: https://developer.apple.com/documentation/CloudKit/CKShare
//
// [CKShare.ParticipantPermission.none]: https://developer.apple.com/documentation/CloudKit/CKShare/ParticipantPermission/none
// [CKShare.SystemFieldKey]: https://developer.apple.com/documentation/CloudKit/CKShare/SystemFieldKey
// [UICloudSharingController]: https://developer.apple.com/documentation/UIKit/UICloudSharingController
type CKShare struct {
	CKRecord
}

// CKShareFromID constructs a [CKShare] from an objc.ID.
//
// A specialized record type that manages a collection of shared records.
func CKShareFromID(id objc.ID) CKShare {
	return CKShare{CKRecord: CKRecordFromID(id)}
}

// NOTE: CKShare adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CKShare] class.
//
// # Creating a Share
//
//   - [ICKShare.InitWithRootRecord]: Creates a new share for the specified record.
//   - [ICKShare.InitWithRootRecordShareID]: Creates a new share for the specified record and record ID.
//   - [ICKShare.InitWithRecordZoneID]: Creates a new share for the specified record zone.
//
// # Accessing the Share’s Attributes
//
//   - [ICKShare.Owner]: The participant that represents the share’s owner.
//   - [ICKShare.CurrentUserParticipant]: The participant that represents the current user.
//   - [ICKShare.Participants]: An array that contains the share’s participants.
//   - [ICKShare.URL]: The URL for inviting participants to the share.
//
// # Configuring the Share
//
//   - [ICKShare.PublicPermission]: The permission for anyone with access to the share’s URL.
//   - [ICKShare.SetPublicPermission]
//   - [ICKShare.AddParticipant]: Adds a participant to the share.
//   - [ICKShare.RemoveParticipant]: Removes a participant from the share.
//
// # Instance Properties
//
//   - [ICKShare.AllowsAccessRequests]: Indicates whether uninvited users can request access to this share.
//   - [ICKShare.SetAllowsAccessRequests]
//   - [ICKShare.BlockedIdentities]: A list of users blocked from requesting access to this share.
//   - [ICKShare.Requesters]: A list of all uninvited users who have requested access to this share.
//
// # Instance Methods
//
//   - [ICKShare.BlockRequesters]: Blocks specified users from requesting access to this share.
//   - [ICKShare.DenyRequesters]: Denies access requests from specified users.
//   - [ICKShare.UnblockIdentities]: Unblocks previously blocked users, allowing them to request access again.
//
// See: https://developer.apple.com/documentation/CloudKit/CKShare
type ICKShare interface {
	ICKRecord

	// Topic: Creating a Share

	// Creates a new share for the specified record.
	InitWithRootRecord(rootRecord ICKRecord) CKShare
	// Creates a new share for the specified record and record ID.
	InitWithRootRecordShareID(rootRecord ICKRecord, shareID ICKRecordID) CKShare
	// Creates a new share for the specified record zone.
	InitWithRecordZoneID(recordZoneID ICKRecordZoneID) CKShare

	// Topic: Accessing the Share’s Attributes

	// The participant that represents the share’s owner.
	Owner() ICKShareParticipant
	// The participant that represents the current user.
	CurrentUserParticipant() ICKShareParticipant
	// An array that contains the share’s participants.
	Participants() []CKShareParticipant
	// The URL for inviting participants to the share.
	URL() foundation.NSURL

	// Topic: Configuring the Share

	// The permission for anyone with access to the share’s URL.
	PublicPermission() CKShareParticipantPermission
	SetPublicPermission(value CKShareParticipantPermission)
	// Adds a participant to the share.
	AddParticipant(participant ICKShareParticipant)
	// Removes a participant from the share.
	RemoveParticipant(participant ICKShareParticipant)

	// Topic: Instance Properties

	// Indicates whether uninvited users can request access to this share.
	AllowsAccessRequests() bool
	SetAllowsAccessRequests(value bool)
	// A list of users blocked from requesting access to this share.
	BlockedIdentities() []CKShareBlockedIdentity
	// A list of all uninvited users who have requested access to this share.
	Requesters() []CKShareAccessRequester

	// Topic: Instance Methods

	// Blocks specified users from requesting access to this share.
	BlockRequesters(requesters []CKShareAccessRequester)
	// Denies access requests from specified users.
	DenyRequesters(requesters []CKShareAccessRequester)
	// Unblocks previously blocked users, allowing them to request access again.
	UnblockIdentities(blockedIdentities []CKShareBlockedIdentity)

	// The name of a share record that manages a shared record zone.
	CKRecordNameZoneWideShare() string
	// An error that occurs when CloudKit attempts to share a record with an existing share.
	AlreadyShared() CKErrorCode
	SetAlreadyShared(value CKErrorCode)
	// A Boolean value that indicates whether the user has an iCloud account.
	HasiCloudAccount() bool
	SetHasiCloudAccount(value bool)
	// The unique name of the record.
	RecordName() string
	SetRecordName(value string)
	// The identity of the participant.
	UserIdentity() ICKUserIdentity
	SetUserIdentity(value ICKUserIdentity)
}

// Init initializes the instance.
func (c CKShare) Init() CKShare {
	rv := objc.Send[CKShare](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CKShare) Autorelease() CKShare {
	rv := objc.Send[CKShare](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKShare creates a new CKShare instance.
func NewCKShare() CKShare {
	class := getCKShareClass()
	rv := objc.Send[CKShare](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a share from a serialized instance.
//
// aDecoder: The coder to use when deserializing the share.
//
// # Discussion
//
// When saving a newly created [CKShare], you must save the share and its
// [RootRecord] in the same [CKModifyRecordsOperation] batch.
//
// See: https://developer.apple.com/documentation/CloudKit/CKShare/init(coder:)
func NewCKShareWithCoder(aDecoder foundation.INSCoder) CKShare {
	instance := getCKShareClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return CKShareFromID(rv)
}

// Creates a new share for the specified record zone.
//
// recordZoneID: The ID of the record zone to share.
//
// # Discussion
//
// A shared record zone must have the [CKRecordZoneCapabilityZoneWideSharing]
// capability. Custom record zones that you create in the user’s private
// database have this capability by default. A record zone, and the records it
// contains, can take part in only a single share.
//
// After accepting a share invite, CloudKit adds the records of the shared
// record zone to a new zone in the participant’s shared database. Use
// [CKFetchDatabaseChangesOperation] to fetch the ID of the new record zone.
// Then configure [CKFetchRecordZoneChangesOperation] with that record zone ID
// and execute the operation to fetch the records.
//
// If you use [CKFetchShareMetadataOperation] to fetch the metadata for a
// shared record zone, the operation ignores the [ShouldFetchRootRecord] and
// [RootRecordDesiredKeys] properties because, unlike a shared record
// hierarchy, a record zone doesn’t have a nominated root record.
//
// See: https://developer.apple.com/documentation/CloudKit/CKShare/init(recordZoneID:)
func NewCKShareWithRecordZoneID(recordZoneID ICKRecordZoneID) CKShare {
	instance := getCKShareClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithRecordZoneID:"), recordZoneID)
	return CKShareFromID(rv)
}

// Creates a new share for the specified record.
//
// rootRecord: The record to share.
//
// # Discussion
//
// When saving a newly created [CKShare], you must save the share and its
// [RootRecord] in the same [CKModifyRecordsOperation] batch.
//
// See: https://developer.apple.com/documentation/CloudKit/CKShare/init(rootRecord:)
func NewCKShareWithRootRecord(rootRecord ICKRecord) CKShare {
	instance := getCKShareClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithRootRecord:"), rootRecord)
	return CKShareFromID(rv)
}

// Creates a new share for the specified record and record ID.
//
// rootRecord: The record to share.
//
// shareID: The [CKRecordID] for the share.
//
// # Discussion
//
// When saving a newly created [CKShare], save the share and its [RootRecord]
// in the same [CKModifyRecordsOperation] batch.
//
// See: https://developer.apple.com/documentation/CloudKit/CKShare/init(rootRecord:shareID:)
func NewCKShareWithRootRecordShareID(rootRecord ICKRecord, shareID ICKRecordID) CKShare {
	instance := getCKShareClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithRootRecord:shareID:"), rootRecord, shareID)
	return CKShareFromID(rv)
}

// Creates a new share for the specified record.
//
// rootRecord: The record to share.
//
// # Discussion
//
// When saving a newly created [CKShare], you must save the share and its
// [RootRecord] in the same [CKModifyRecordsOperation] batch.
//
// See: https://developer.apple.com/documentation/CloudKit/CKShare/init(rootRecord:)
func (c CKShare) InitWithRootRecord(rootRecord ICKRecord) CKShare {
	rv := objc.Send[CKShare](c.ID, objc.Sel("initWithRootRecord:"), rootRecord)
	return rv
}

// Creates a new share for the specified record and record ID.
//
// rootRecord: The record to share.
//
// shareID: The [CKRecordID] for the share.
//
// # Discussion
//
// When saving a newly created [CKShare], save the share and its [RootRecord]
// in the same [CKModifyRecordsOperation] batch.
//
// See: https://developer.apple.com/documentation/CloudKit/CKShare/init(rootRecord:shareID:)
func (c CKShare) InitWithRootRecordShareID(rootRecord ICKRecord, shareID ICKRecordID) CKShare {
	rv := objc.Send[CKShare](c.ID, objc.Sel("initWithRootRecord:shareID:"), rootRecord, shareID)
	return rv
}

// Creates a new share for the specified record zone.
//
// recordZoneID: The ID of the record zone to share.
//
// # Discussion
//
// A shared record zone must have the [CKRecordZoneCapabilityZoneWideSharing]
// capability. Custom record zones that you create in the user’s private
// database have this capability by default. A record zone, and the records it
// contains, can take part in only a single share.
//
// After accepting a share invite, CloudKit adds the records of the shared
// record zone to a new zone in the participant’s shared database. Use
// [CKFetchDatabaseChangesOperation] to fetch the ID of the new record zone.
// Then configure [CKFetchRecordZoneChangesOperation] with that record zone ID
// and execute the operation to fetch the records.
//
// If you use [CKFetchShareMetadataOperation] to fetch the metadata for a
// shared record zone, the operation ignores the [ShouldFetchRootRecord] and
// [RootRecordDesiredKeys] properties because, unlike a shared record
// hierarchy, a record zone doesn’t have a nominated root record.
//
// See: https://developer.apple.com/documentation/CloudKit/CKShare/init(recordZoneID:)
func (c CKShare) InitWithRecordZoneID(recordZoneID ICKRecordZoneID) CKShare {
	rv := objc.Send[CKShare](c.ID, objc.Sel("initWithRecordZoneID:"), recordZoneID)
	return rv
}

// Adds a participant to the share.
//
// participant: The participant to add to the share.
//
// # Discussion
//
// If a participant with a matching [UserIdentity] already exists in the
// share, the system updates the existing participant’s properties and
// doesn’t add a new participant.
//
// To modify the list of participants, a share’s [PublicPermission] must be
// [CKShare.ParticipantPermission.none]. You can’t mix and match public and
// private users in the same share. You can only add certain participant types
// with this API. See [CKShareParticipant] for more information.
//
// See: https://developer.apple.com/documentation/CloudKit/CKShare/addParticipant(_:)
//
// [CKShare.ParticipantPermission.none]: https://developer.apple.com/documentation/CloudKit/CKShare/ParticipantPermission/none
func (c CKShare) AddParticipant(participant ICKShareParticipant) {
	objc.Send[objc.ID](c.ID, objc.Sel("addParticipant:"), participant)
}

// Removes a participant from the share.
//
// participant: The participant to remove from the share.
//
// # Discussion
//
// To modify the list of participants, a share’s [PublicPermission] must be
// [CKShare.ParticipantPermission.none]. You can’t mix and match public and
// private users in the same share. You can only add certain participant types
// with this API. See [CKShareParticipant] for more information.
//
// See: https://developer.apple.com/documentation/CloudKit/CKShare/removeParticipant(_:)
//
// [CKShare.ParticipantPermission.none]: https://developer.apple.com/documentation/CloudKit/CKShare/ParticipantPermission/none
func (c CKShare) RemoveParticipant(participant ICKShareParticipant) {
	objc.Send[objc.ID](c.ID, objc.Sel("removeParticipant:"), participant)
}

// Blocks specified users from requesting access to this share.
//
// requesters: An array of [CKShareAccessRequester] objects to block.
//
// # Discussion
//
// Blocking prevents users from submitting future access requests and removes
// existing participants from the share. Blocked requesters appear in the
// [BlockedIdentities] array.
//
// To persist this change, save the share to the server after calling this
// method.
//
// Only the share owner or an administrator can invoke this method. Attempts
// by other participants result in an exception.
//
// See: https://developer.apple.com/documentation/CloudKit/CKShare/blockRequesters(_:)
func (c CKShare) BlockRequesters(requesters []CKShareAccessRequester) {
	objc.Send[objc.ID](c.ID, objc.Sel("blockRequesters:"), objectivec.IObjectSliceToNSArray(requesters))
}

// Denies access requests from specified users.
//
// requesters: An array of [CKShareAccessRequester] objects to deny.
//
// # Discussion
//
// Use this method to deny pending access requests from uninvited users.
// CloudKit removes denied requesters from the [Requesters] array. To persist
// the changes, save the share to the server after calling this method.
//
// After denial, requesters can still submit new access requests unless
// explicitly blocked using [BlockRequesters].
//
// Only the share owner or an administrator can invoke this method. Attempts
// by other participants result in an exception.
//
// See: https://developer.apple.com/documentation/CloudKit/CKShare/denyRequesters(_:)
func (c CKShare) DenyRequesters(requesters []CKShareAccessRequester) {
	objc.Send[objc.ID](c.ID, objc.Sel("denyRequesters:"), objectivec.IObjectSliceToNSArray(requesters))
}

// Unblocks previously blocked users, allowing them to request access again.
//
// blockedIdentities: An array of [CKShareBlockedIdentity] objects to unblock.
//
// # Discussion
//
// Use this method to remove specified identities from the [BlockedIdentities]
// array. Unblocked identities can request access again if access requests are
// enabled.
//
// To persist this change, save the share to the server after calling this
// method.
//
// Only the share owner or an administrator can invoke this method. Attempts
// by other participants result in an exception.
//
// See: https://developer.apple.com/documentation/CloudKit/CKShare/unblockIdentities(_:)
func (c CKShare) UnblockIdentities(blockedIdentities []CKShareBlockedIdentity) {
	objc.Send[objc.ID](c.ID, objc.Sel("unblockIdentities:"), objectivec.IObjectSliceToNSArray(blockedIdentities))
}

// The participant that represents the share’s owner.
//
// See: https://developer.apple.com/documentation/CloudKit/CKShare/owner
func (c CKShare) Owner() ICKShareParticipant {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("owner"))
	return CKShareParticipantFromID(objc.ID(rv))
}

// The participant that represents the current user.
//
// See: https://developer.apple.com/documentation/CloudKit/CKShare/currentUserParticipant
func (c CKShare) CurrentUserParticipant() ICKShareParticipant {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("currentUserParticipant"))
	return CKShareParticipantFromID(objc.ID(rv))
}

// An array that contains the share’s participants.
//
// # Discussion
//
// The property’s value contains all of the share’s participants that the
// current user has permissions to see. At a minimum, it includes the
// share’s owner and the current user.
//
// See: https://developer.apple.com/documentation/CloudKit/CKShare/participants
func (c CKShare) Participants() []CKShareParticipant {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("participants"))
	return objc.ConvertSlice(rv, func(id objc.ID) CKShareParticipant {
		return CKShareParticipantFromID(id)
	})
}

// The URL for inviting participants to the share.
//
// # Discussion
//
// This property is only available after saving a share record to the server.
// This URL is stable and persists across shares and reshares of the same root
// record.
//
// See: https://developer.apple.com/documentation/CloudKit/CKShare/url
func (c CKShare) URL() foundation.NSURL {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("URL"))
	return foundation.NSURLFromID(objc.ID(rv))
}

// The permission for anyone with access to the share’s URL.
//
// # Discussion
//
// Setting this property’s value to be more permissive than
// [CKShare.ParticipantPermission.none] allows any user with the share’s URL
// to join. CloudKit removes all public participants when you save the share
// if you set the property’s value to [CKShare.ParticipantPermission.none].
//
// The default value is [CKShare.ParticipantPermission.none]
//
// See: https://developer.apple.com/documentation/CloudKit/CKShare/publicPermission
//
// [CKShare.ParticipantPermission.none]: https://developer.apple.com/documentation/CloudKit/CKShare/ParticipantPermission/none
func (c CKShare) PublicPermission() CKShareParticipantPermission {
	rv := objc.Send[CKShareParticipantPermission](c.ID, objc.Sel("publicPermission"))
	return CKShareParticipantPermission(rv)
}
func (c CKShare) SetPublicPermission(value CKShareParticipantPermission) {
	objc.Send[struct{}](c.ID, objc.Sel("setPublicPermission:"), value)
}

// Indicates whether uninvited users can request access to this share.
//
// # Discussion
//
// By default, this property is [NO]. When this property is [YES], uninvited
// users can request access to the share if they discover the share URL. When
// this property is [NO], the server prevents uninvited users from requesting
// access and does not indicate whether the share exists.
//
// Only the share owner or an administrator can modify this property. If
// another participant attempts to modify this property, CloudKit throws an
// exception.
//
// See: https://developer.apple.com/documentation/CloudKit/CKShare/allowsAccessRequests
func (c CKShare) AllowsAccessRequests() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("allowsAccessRequests"))
	return rv
}
func (c CKShare) SetAllowsAccessRequests(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setAllowsAccessRequests:"), value)
}

// A list of users blocked from requesting access to this share.
//
// # Discussion
//
// Identities remain in this list until an owner or administrator calls
// [UnblockIdentities].
//
// See: https://developer.apple.com/documentation/CloudKit/CKShare/blockedIdentities
func (c CKShare) BlockedIdentities() []CKShareBlockedIdentity {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("blockedIdentities"))
	return objc.ConvertSlice(rv, func(id objc.ID) CKShareBlockedIdentity {
		return CKShareBlockedIdentityFromID(id)
	})
}

// A list of all uninvited users who have requested access to this share.
//
// # Discussion
//
// When share access requests are allowed, uninvited users can request to join
// the share. All pending access requests appear in this array. Each requester
// is returned with name components and either an email or phone number.
//
// Either share owners or administrators can respond to these access requests.
//
// # Responding to Access Requests:
//
// - - Fetch the participant information by running
// [CKFetchShareParticipantsOperation] with the requester’s
// [ParticipantLookupInfo]. - Add the resulting participant to the share. - -
// Use [DenyRequesters] to remove the requester from the requesters list. - -
// Use [BlockRequesters] to block requesters. - Blocking a requester prevents
// them from sending future access requests to the share.
//
// See: https://developer.apple.com/documentation/CloudKit/CKShare/requesters
func (c CKShare) Requesters() []CKShareAccessRequester {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("requesters"))
	return objc.ConvertSlice(rv, func(id objc.ID) CKShareAccessRequester {
		return CKShareAccessRequesterFromID(id)
	})
}

// The name of a share record that manages a shared record zone.
//
// See: https://developer.apple.com/documentation/cloudkit/ckrecordnamezonewideshare
func (c CKShare) CKRecordNameZoneWideShare() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("CKRecordNameZoneWideShare"))
	return foundation.NSStringFromID(rv).String()
}

// An error that occurs when CloudKit attempts to share a record with an
// existing share.
//
// See: https://developer.apple.com/documentation/cloudkit/ckerror/alreadyshared
func (c CKShare) AlreadyShared() CKErrorCode {
	rv := objc.Send[CKErrorCode](c.ID, objc.Sel("alreadyShared"))
	return CKErrorCode(rv)
}
func (c CKShare) SetAlreadyShared(value CKErrorCode) {
	objc.Send[struct{}](c.ID, objc.Sel("setAlreadyShared:"), value)
}

// A Boolean value that indicates whether the user has an iCloud account.
//
// See: https://developer.apple.com/documentation/cloudkit/ckuseridentity/hasicloudaccount
func (c CKShare) HasiCloudAccount() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("hasiCloudAccount"))
	return rv
}
func (c CKShare) SetHasiCloudAccount(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setHasiCloudAccount:"), value)
}

// The unique name of the record.
//
// See: https://developer.apple.com/documentation/cloudkit/ckrecord/id/recordname
func (c CKShare) RecordName() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("recordName"))
	return foundation.NSStringFromID(rv).String()
}
func (c CKShare) SetRecordName(value string) {
	objc.Send[struct{}](c.ID, objc.Sel("setRecordName:"), objc.String(value))
}

// The identity of the participant.
//
// See: https://developer.apple.com/documentation/cloudkit/ckshare/participant/useridentity
func (c CKShare) UserIdentity() ICKUserIdentity {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("userIdentity"))
	return CKUserIdentityFromID(objc.ID(rv))
}
func (c CKShare) SetUserIdentity(value ICKUserIdentity) {
	objc.Send[struct{}](c.ID, objc.Sel("setUserIdentity:"), value)
}
