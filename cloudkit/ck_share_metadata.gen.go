// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CKShareMetadata] class.
var (
	_CKShareMetadataClass     CKShareMetadataClass
	_CKShareMetadataClassOnce sync.Once
)

func getCKShareMetadataClass() CKShareMetadataClass {
	_CKShareMetadataClassOnce.Do(func() {
		_CKShareMetadataClass = CKShareMetadataClass{class: objc.GetClass("CKShareMetadata")}
	})
	return _CKShareMetadataClass
}

// GetCKShareMetadataClass returns the class object for CKShareMetadata.
func GetCKShareMetadataClass() CKShareMetadataClass {
	return getCKShareMetadataClass()
}

type CKShareMetadataClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CKShareMetadataClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CKShareMetadataClass) Alloc() CKShareMetadata {
	rv := objc.Send[CKShareMetadata](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// An object that describes a shared record’s metadata.
//
// # Overview
//
// A share’s metadata is an intermediary object that provides access to the
// share, its owner, and, for a shared record hierarchy, its root record.
// Metadata also includes details about the current user’s participation in
// the share.
//
// You don’t create metadata. CloudKit provides it to your app when the user
// taps or clicks a share’s [CKShare.URL], such as in an email or a message.
// The method CloudKit calls varies by platform and app configuration, and
// includes the following:
//
// - For a scene-based iOS app in a running or suspended state, CloudKit calls
// the [windowScene(_:userDidAcceptCloudKitShareWith:)] method on your window
// scene delegate. - For a scene-based iOS app that’s not running, the
// system launches your app in response to the tap or click, and calls the
// [scene(_:willConnectTo:options:)] method on your scene delegate. The
// `connectionOptions` parameter contains the metadata. Use its
// [cloudKitShareMetadata] property to access it. - For an iOS app that
// doesn’t use scenes, CloudKit calls your app delegate’s
// [application(_:userDidAcceptCloudKitShareWith:)] method. - For a macOS app,
// CloudKit calls your app delegate’s
// [application(_:userDidAcceptCloudKitShareWith:)] method. - For a watchOS
// app, CloudKit calls the [userDidAcceptCloudKitShare(with:)] method on your
// watch extension delegate.
//
// Respond by checking the [CKShareMetadata.ParticipantStatus] of the provided
// metadata. If the status is `pending`, use [CKAcceptSharesOperation] to
// accept participation in the share. You can also fetch metadata independent
// of this flow using [CKFetchShareMetadataOperation].
//
// For a shared record hierarchy, the
// [CKShareMetadata.HierarchicalRootRecordID] property contains the ID of the
// share’s root record. When using [CKFetchShareMetadataOperation] to fetch
// metadata, you can include the entire root record by setting the
// operation’s [CKFetchShareMetadataOperation.ShouldFetchRootRecord]
// property to true. CloudKit then populates the [CKShareMetadata.RootRecord]
// property before it returns the metadata. You can further customize this
// behavior using the operation’s
// [CKFetchShareMetadataOperation.RootRecordDesiredKeys] property to specify
// which fields to return. This functionality isn’t applicable for a shared
// record zone because, unlike a shared record hierarchy, it doesn’t have a
// nominated root record.
//
// The participant properties provide the current user’s acceptance status,
// permissions, and role. Use these values to determine what functionality to
// provide to the user. For example, only display editing controls for
// accepted participants with `readWrite` permissions.
//
// # Accessing the Share
//
//   - [CKShareMetadata.Share]: The share that owns the metadata.
//   - [CKShareMetadata.ContainerIdentifier]: The ID of the share’s container.
//   - [CKShareMetadata.OwnerIdentity]: The identity of the share’s owner.
//
// # Accessing the Root Record
//
//   - [CKShareMetadata.HierarchicalRootRecordID]: The record ID of the shared hierarchy’s root record.
//   - [CKShareMetadata.RootRecord]: The share’s root record.
//
// # Accessing the Participant’s Capabilities
//
//   - [CKShareMetadata.ParticipantRole]: The share’s participant role for the user who retrieves the metadata.
//   - [CKShareMetadata.ParticipantPermission]: The share’s permissions for the user who retrieves the metadata.
//   - [CKShareMetadata.ParticipantStatus]: The share’s participation status for the user who retrieves the metadata.
//   - [CKShareMetadata.ParticipantType]: The share’s participation type for the user who retrieves the metadata.
//
// See: https://developer.apple.com/documentation/CloudKit/CKShare/Metadata
//
// [application(_:userDidAcceptCloudKitShareWith:)]: https://developer.apple.com/documentation/AppKit/NSApplicationDelegate/application(_:userDidAcceptCloudKitShareWith:)
// [cloudKitShareMetadata]: https://developer.apple.com/documentation/UIKit/UIScene/ConnectionOptions/cloudKitShareMetadata
// [scene(_:willConnectTo:options:)]: https://developer.apple.com/documentation/UIKit/UISceneDelegate/scene(_:willConnectTo:options:)
// [userDidAcceptCloudKitShare(with:)]: https://developer.apple.com/documentation/WatchKit/WKExtensionDelegate/userDidAcceptCloudKitShare(with:)
// [windowScene(_:userDidAcceptCloudKitShareWith:)]: https://developer.apple.com/documentation/UIKit/UIWindowSceneDelegate/windowScene(_:userDidAcceptCloudKitShareWith:)
type CKShareMetadata struct {
	objectivec.Object
}

// CKShareMetadataFromID constructs a [CKShareMetadata] from an objc.ID.
//
// An object that describes a shared record’s metadata.
func CKShareMetadataFromID(id objc.ID) CKShareMetadata {
	return CKShareMetadata{objectivec.Object{ID: id}}
}

// NOTE: CKShareMetadata adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CKShareMetadata] class.
//
// # Accessing the Share
//
//   - [ICKShareMetadata.Share]: The share that owns the metadata.
//   - [ICKShareMetadata.ContainerIdentifier]: The ID of the share’s container.
//   - [ICKShareMetadata.OwnerIdentity]: The identity of the share’s owner.
//
// # Accessing the Root Record
//
//   - [ICKShareMetadata.HierarchicalRootRecordID]: The record ID of the shared hierarchy’s root record.
//   - [ICKShareMetadata.RootRecord]: The share’s root record.
//
// # Accessing the Participant’s Capabilities
//
//   - [ICKShareMetadata.ParticipantRole]: The share’s participant role for the user who retrieves the metadata.
//   - [ICKShareMetadata.ParticipantPermission]: The share’s permissions for the user who retrieves the metadata.
//   - [ICKShareMetadata.ParticipantStatus]: The share’s participation status for the user who retrieves the metadata.
//   - [ICKShareMetadata.ParticipantType]: The share’s participation type for the user who retrieves the metadata.
//
// See: https://developer.apple.com/documentation/CloudKit/CKShare/Metadata
type ICKShareMetadata interface {
	objectivec.IObject

	// Topic: Accessing the Share

	// The share that owns the metadata.
	Share() ICKShare
	// The ID of the share’s container.
	ContainerIdentifier() string
	// The identity of the share’s owner.
	OwnerIdentity() ICKUserIdentity

	// Topic: Accessing the Root Record

	// The record ID of the shared hierarchy’s root record.
	HierarchicalRootRecordID() ICKRecordID
	// The share’s root record.
	RootRecord() ICKRecord

	// Topic: Accessing the Participant’s Capabilities

	// The share’s participant role for the user who retrieves the metadata.
	ParticipantRole() CKShareParticipantRole
	// The share’s permissions for the user who retrieves the metadata.
	ParticipantPermission() CKShareParticipantPermission
	// The share’s participation status for the user who retrieves the metadata.
	ParticipantStatus() CKShareParticipantAcceptanceStatus
	// The share’s participation type for the user who retrieves the metadata.
	ParticipantType() unsafe.Pointer

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (c CKShareMetadata) Init() CKShareMetadata {
	rv := objc.Send[CKShareMetadata](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CKShareMetadata) Autorelease() CKShareMetadata {
	rv := objc.Send[CKShareMetadata](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKShareMetadata creates a new CKShareMetadata instance.
func NewCKShareMetadata() CKShareMetadata {
	class := getCKShareMetadataClass()
	rv := objc.Send[CKShareMetadata](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (c CKShareMetadata) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](c.ID, objc.Sel("encodeWithCoder:"), coder)
}

// The share that owns the metadata.
//
// See: https://developer.apple.com/documentation/CloudKit/CKShare/Metadata/share
func (c CKShareMetadata) Share() ICKShare {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("share"))
	return CKShareFromID(objc.ID(rv))
}

// The ID of the share’s container.
//
// See: https://developer.apple.com/documentation/CloudKit/CKShare/Metadata/containerIdentifier
func (c CKShareMetadata) ContainerIdentifier() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("containerIdentifier"))
	return foundation.NSStringFromID(rv).String()
}

// The identity of the share’s owner.
//
// See: https://developer.apple.com/documentation/CloudKit/CKShare/Metadata/ownerIdentity
func (c CKShareMetadata) OwnerIdentity() ICKUserIdentity {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("ownerIdentity"))
	return CKUserIdentityFromID(objc.ID(rv))
}

// The record ID of the shared hierarchy’s root record.
//
// # Discussion
//
// CloudKit populates this property only for metadata that belongs to a shared
// record hierarchy. If the metadata is part of a shared record zone, the
// property is `nil`. This is because, unlike a shared record hierarchy, a
// shared record zone doesn’t have a nominated root record.
//
// See: https://developer.apple.com/documentation/CloudKit/CKShare/Metadata/hierarchicalRootRecordID
func (c CKShareMetadata) HierarchicalRootRecordID() ICKRecordID {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("hierarchicalRootRecordID"))
	return CKRecordIDFromID(objc.ID(rv))
}

// The share’s root record.
//
// # Discussion
//
// This property contains the root record of the shared record hierarchy if
// you set the [CKFetchShareMetadataOperation.ShouldFetchRootRecord] property
// of the operation that fetches the metadata to true. You can specify which
// fields CloudKit returns by setting the same operation’s
// [CKFetchShareMetadataOperation.RootRecordDesiredKeys] property.
//
// The operation ignores the
// [CKFetchShareMetadataOperation.ShouldFetchRootRecord] and
// [CKFetchShareMetadataOperation.RootRecordDesiredKeys] properties when
// fetching a shared record zone’s metadata because, unlike a shared record
// hierarchy, a record zone doesn’t have a nominated root record.
//
// See: https://developer.apple.com/documentation/CloudKit/CKShare/Metadata/rootRecord
func (c CKShareMetadata) RootRecord() ICKRecord {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("rootRecord"))
	return CKRecordFromID(objc.ID(rv))
}

// The share’s participant role for the user who retrieves the metadata.
//
// See: https://developer.apple.com/documentation/CloudKit/CKShare/Metadata/participantRole
func (c CKShareMetadata) ParticipantRole() CKShareParticipantRole {
	rv := objc.Send[CKShareParticipantRole](c.ID, objc.Sel("participantRole"))
	return CKShareParticipantRole(rv)
}

// The share’s permissions for the user who retrieves the metadata.
//
// See: https://developer.apple.com/documentation/CloudKit/CKShare/Metadata/participantPermission
func (c CKShareMetadata) ParticipantPermission() CKShareParticipantPermission {
	rv := objc.Send[CKShareParticipantPermission](c.ID, objc.Sel("participantPermission"))
	return CKShareParticipantPermission(rv)
}

// The share’s participation status for the user who retrieves the metadata.
//
// See: https://developer.apple.com/documentation/CloudKit/CKShare/Metadata/participantStatus
func (c CKShareMetadata) ParticipantStatus() CKShareParticipantAcceptanceStatus {
	rv := objc.Send[CKShareParticipantAcceptanceStatus](c.ID, objc.Sel("participantStatus"))
	return CKShareParticipantAcceptanceStatus(rv)
}

// The share’s participation type for the user who retrieves the metadata.
//
// See: https://developer.apple.com/documentation/CloudKit/CKShare/Metadata/participantType
func (c CKShareMetadata) ParticipantType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c.ID, objc.Sel("participantType"))
	return rv
}
