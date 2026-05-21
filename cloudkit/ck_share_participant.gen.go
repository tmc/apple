// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CKShareParticipant] class.
var (
	_CKShareParticipantClass     CKShareParticipantClass
	_CKShareParticipantClassOnce sync.Once
)

func getCKShareParticipantClass() CKShareParticipantClass {
	_CKShareParticipantClassOnce.Do(func() {
		_CKShareParticipantClass = CKShareParticipantClass{class: objc.GetClass("CKShareParticipant")}
	})
	return _CKShareParticipantClass
}

// GetCKShareParticipantClass returns the class object for CKShareParticipant.
func GetCKShareParticipantClass() CKShareParticipantClass {
	return getCKShareParticipantClass()
}

type CKShareParticipantClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CKShareParticipantClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CKShareParticipantClass) Alloc() CKShareParticipant {
	rv := objc.Send[CKShareParticipant](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// An object that describes a user’s participation in a share.
//
// # Overview
//
// Participants are a key element of sharing in CloudKit. A participant
// provides information about an iCloud user and their participation in a
// share, including their identity, acceptance status, permissions, and role.
//
// The acceptance status determines the participant’s visibilty of the
// shared records. Statuses are: `pending`, `accepted`, `removed`, and
// `unknown`. If the status is `pending`, use [CKAcceptSharesOperation] to
// accept the share. Upon acceptance, CloudKit makes the shared records
// available in the participant’s shared database. The records remain
// accessible for as long as the participant’s status is `accepted`.
//
// You don’t create participants. Use the share’s [CKShare.Participants]
// property to access its existing participants. Use
// [UICloudSharingController] to manage the share’s participants and their
// permissions. Alternatively, you can generate participants using
// [CKFetchShareParticipantsOperation]. Participants must have an active
// iCloud account.
//
// Anyone with the URL of a public share can become a participant in that
// share. Participants of a public share assume the `publicUser` role. For
// private shares, the owner manages the participants. An owner is any
// participant with the `owner` role. A participant of a private share can’t
// accept the share unless the owner adds them first. Private share
// participants assume the `privateUser` role. CloudKit removes any pending
// participants if the owner changes the share’s [CKShare.PublicPermission].
// CloudKit removes all participants if the new permission is `none`.
//
// Participants with write permissions can modify or delete any record that
// you include in the share. However, only the owner can delete a shared
// hierarchy’s root record. If a participant attempts to delete the share,
// CloudKit removes the participant. The share remains active for all other
// participants.
//
// # Accessing the Participant’s Status
//
//   - [CKShareParticipant.AcceptanceStatus]: The current state of the user’s acceptance of the share.
//
// # Accessing the Participant’s Identity
//
//   - [CKShareParticipant.UserIdentity]: The identity of the participant.
//
// # Managing the Participant’s Capabilites
//
//   - [CKShareParticipant.Permission]: The participant’s permission level for the share.
//   - [CKShareParticipant.SetPermission]
//   - [CKShareParticipant.Role]: The participant’s role for the share.
//   - [CKShareParticipant.SetRole]
//
// # Instance Properties
//
//   - [CKShareParticipant.DateAddedToShare]: The date and time when the participant was added to the share.
//   - [CKShareParticipant.IsApprovedRequester]: Indicates whether the participant was originally a requester who was approved to join the share.
//
// See: https://developer.apple.com/documentation/CloudKit/CKShare/Participant
//
// [UICloudSharingController]: https://developer.apple.com/documentation/UIKit/UICloudSharingController
type CKShareParticipant struct {
	objectivec.Object
}

// CKShareParticipantFromID constructs a [CKShareParticipant] from an objc.ID.
//
// An object that describes a user’s participation in a share.
func CKShareParticipantFromID(id objc.ID) CKShareParticipant {
	return CKShareParticipant{objectivec.Object{ID: id}}
}

// NOTE: CKShareParticipant adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CKShareParticipant] class.
//
// # Accessing the Participant’s Status
//
//   - [ICKShareParticipant.AcceptanceStatus]: The current state of the user’s acceptance of the share.
//
// # Accessing the Participant’s Identity
//
//   - [ICKShareParticipant.UserIdentity]: The identity of the participant.
//
// # Managing the Participant’s Capabilites
//
//   - [ICKShareParticipant.Permission]: The participant’s permission level for the share.
//   - [ICKShareParticipant.SetPermission]
//   - [ICKShareParticipant.Role]: The participant’s role for the share.
//   - [ICKShareParticipant.SetRole]
//
// # Instance Properties
//
//   - [ICKShareParticipant.DateAddedToShare]: The date and time when the participant was added to the share.
//   - [ICKShareParticipant.IsApprovedRequester]: Indicates whether the participant was originally a requester who was approved to join the share.
//
// See: https://developer.apple.com/documentation/CloudKit/CKShare/Participant
type ICKShareParticipant interface {
	objectivec.IObject

	// Topic: Accessing the Participant’s Status

	// The current state of the user’s acceptance of the share.
	AcceptanceStatus() CKShareParticipantAcceptanceStatus

	// Topic: Accessing the Participant’s Identity

	// The identity of the participant.
	UserIdentity() ICKUserIdentity

	// Topic: Managing the Participant’s Capabilites

	// The participant’s permission level for the share.
	Permission() CKShareParticipantPermission
	SetPermission(value CKShareParticipantPermission)
	// The participant’s role for the share.
	Role() CKShareParticipantRole
	SetRole(value CKShareParticipantRole)

	// Topic: Instance Properties

	// The date and time when the participant was added to the share.
	DateAddedToShare() foundation.NSDate
	// Indicates whether the participant was originally a requester who was approved to join the share.
	IsApprovedRequester() bool

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (c CKShareParticipant) Init() CKShareParticipant {
	rv := objc.Send[CKShareParticipant](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CKShareParticipant) Autorelease() CKShareParticipant {
	rv := objc.Send[CKShareParticipant](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKShareParticipant creates a new CKShareParticipant instance.
func NewCKShareParticipant() CKShareParticipant {
	class := getCKShareParticipantClass()
	rv := objc.Send[CKShareParticipant](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (c CKShareParticipant) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](c.ID, objc.Sel("encodeWithCoder:"), coder)
}

// Generate a unique URL for inviting a participant without knowing their
// handle
//
// # Discussion
//
// When a participant’s email address / phone number / userRecordID isn’t
// known up-front, a [CKShareParticipantClass.OneTimeURLParticipant] can be
// added to the share. Once the share is saved, a custom invitation link or
// one-time URL is available for the added participant via
// [oneTimeURLForParticipantID:]. This custom link can be used by any
// recipient user to fetch share metadata and accept the share.
//
// Note that a one-time URL participant in the
// [CKShare.ParticipantAcceptanceStatus.pending] state has empty
// [CKUserIdentity.NameComponents] and a nil [CKUserIdentity.LookupInfo].
//
// See: https://developer.apple.com/documentation/CloudKit/CKShare/Participant/oneTimeURLParticipant()
//
// [CKShare.ParticipantAcceptanceStatus.pending]: https://developer.apple.com/documentation/CloudKit/CKShare/ParticipantAcceptanceStatus/pending
// [oneTimeURLForParticipantID:]: https://developer.apple.com/documentation/CloudKit/CKShare/oneTimeURLForParticipantID:
func (_CKShareParticipantClass CKShareParticipantClass) OneTimeURLParticipant() CKShareParticipant {
	rv := objc.Send[objc.ID](objc.ID(_CKShareParticipantClass.class), objc.Sel("oneTimeURLParticipant"))
	return CKShareParticipantFromID(rv)
}

// The current state of the user’s acceptance of the share.
//
// # Discussion
//
// This property contains the current state of the participant’s acceptance
// of the share. For a list of possible values, see
// [CKShare.ParticipantAcceptanceStatus].
//
// See: https://developer.apple.com/documentation/CloudKit/CKShare/Participant/acceptanceStatus-swift.property
//
// [CKShare.ParticipantAcceptanceStatus]: https://developer.apple.com/documentation/CloudKit/CKShare/ParticipantAcceptanceStatus
func (c CKShareParticipant) AcceptanceStatus() CKShareParticipantAcceptanceStatus {
	rv := objc.Send[CKShareParticipantAcceptanceStatus](c.ID, objc.Sel("acceptanceStatus"))
	return CKShareParticipantAcceptanceStatus(rv)
}

// The identity of the participant.
//
// # Discussion
//
// This property contains a reference to the user identity for the share
// participant.
//
// See: https://developer.apple.com/documentation/CloudKit/CKShare/Participant/userIdentity
func (c CKShareParticipant) UserIdentity() ICKUserIdentity {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("userIdentity"))
	return CKUserIdentityFromID(objc.ID(rv))
}

// The participant’s permission level for the share.
//
// # Discussion
//
// This property controls the permissions that the participant has for the
// share. For a list of possible values, see [CKShare.ParticipantPermission].
//
// See: https://developer.apple.com/documentation/CloudKit/CKShare/Participant/permission-swift.property
//
// [CKShare.ParticipantPermission]: https://developer.apple.com/documentation/CloudKit/CKShare/ParticipantPermission
func (c CKShareParticipant) Permission() CKShareParticipantPermission {
	rv := objc.Send[CKShareParticipantPermission](c.ID, objc.Sel("permission"))
	return CKShareParticipantPermission(rv)
}
func (c CKShareParticipant) SetPermission(value CKShareParticipantPermission) {
	objc.Send[struct{}](c.ID, objc.Sel("setPermission:"), value)
}

// The participant’s role for the share.
//
// See: https://developer.apple.com/documentation/CloudKit/CKShare/Participant/role-swift.property
func (c CKShareParticipant) Role() CKShareParticipantRole {
	rv := objc.Send[CKShareParticipantRole](c.ID, objc.Sel("role"))
	return CKShareParticipantRole(rv)
}
func (c CKShareParticipant) SetRole(value CKShareParticipantRole) {
	objc.Send[struct{}](c.ID, objc.Sel("setRole:"), value)
}

// The date and time when the participant was added to the share.
//
// # Discussion
//
// This timestamp is set when the share is successfully saved to the server.
//
// See: https://developer.apple.com/documentation/CloudKit/CKShare/Participant/dateAddedToShare
func (c CKShareParticipant) DateAddedToShare() foundation.NSDate {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("dateAddedToShare"))
	return foundation.NSDateFromID(objc.ID(rv))
}

// Indicates whether the participant was originally a requester who was
// approved to join the share.
//
// See: https://developer.apple.com/documentation/CloudKit/CKShare/Participant/isApprovedRequester
func (c CKShareParticipant) IsApprovedRequester() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isApprovedRequester"))
	return rv
}
