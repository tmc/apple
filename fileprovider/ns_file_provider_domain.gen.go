// Code generated from Apple documentation for FileProvider. DO NOT EDIT.

package fileprovider

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSFileProviderDomain] class.
var (
	_NSFileProviderDomainClass     NSFileProviderDomainClass
	_NSFileProviderDomainClassOnce sync.Once
)

func getNSFileProviderDomainClass() NSFileProviderDomainClass {
	_NSFileProviderDomainClassOnce.Do(func() {
		_NSFileProviderDomainClass = NSFileProviderDomainClass{class: objc.GetClass("NSFileProviderDomain")}
	})
	return _NSFileProviderDomainClass
}

// GetNSFileProviderDomainClass returns the class object for NSFileProviderDomain.
func GetNSFileProviderDomainClass() NSFileProviderDomainClass {
	return getNSFileProviderDomainClass()
}

type NSFileProviderDomainClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSFileProviderDomainClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSFileProviderDomainClass) Alloc() NSFileProviderDomain {
	rv := objc.Send[NSFileProviderDomain](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A File Provider extension’s domain.
//
// # Overview
//
// You can use domains to partition a file provider’s content. When you use
// domains, a single file provider can act as if multiple file providers were
// installed, and users can dynamically switch from one domain to another. You
// can use domains to represent different accounts or locations.
//
// By default, a File Provider extension has no domain. You can register
// domains by calling the [NSFileProviderManager] class’s
// [NSFileProviderManagerClass.AddDomainCompletionHandler] method. A new
// [NSFileProviderExtension] instance is created for each domain that you
// register. The [NSFileProviderExtension] object’s [domain] property
// indicates which domain the file provider belongs to. Any items returned by
// that file provider also belong to the domain.
//
// # Creating domains
//
//   - [NSFileProviderDomain.InitWithIdentifierDisplayName]: Creates a new file provider domain with the specified identifier and display name.
//   - [NSFileProviderDomain.InitWithDisplayNameUserInfoVolumeURL]: Creates a new file provider domain with the specified URL and display name.
//
// # Accessing data
//
//   - [NSFileProviderDomain.DisplayName]: The name of the domain displayed in the user interface.
//   - [NSFileProviderDomain.Identifier]: The domain’s unique identifier.
//   - [NSFileProviderDomain.IsReplicated]
//   - [NSFileProviderDomain.BackingStoreIdentity]: A unique identifier for the backing store used by the system.
//   - [NSFileProviderDomain.IsHidden]: A Boolean value that determines whether the domain is visible to users.
//   - [NSFileProviderDomain.SetHidden]
//   - [NSFileProviderDomain.UserEnabled]: A Boolean value that indicates whether the user has enabled or disabled the domain.
//   - [NSFileProviderDomain.IsDisconnected]: A Boolean value indicating that the domain is present, but disconnected from the file extension.
//   - [NSFileProviderDomain.SupportsSyncingTrash]
//   - [NSFileProviderDomain.SetSupportsSyncingTrash]
//   - [NSFileProviderDomain.UserInfo]
//   - [NSFileProviderDomain.SetUserInfo]
//   - [NSFileProviderDomain.VolumeUUID]
//
// # Syncing Desktop and Documents folders
//
//   - [NSFileProviderDomain.ReplicatedKnownFolders]: A list of known folders that the domain currently replicates.
//   - [NSFileProviderDomain.SupportedKnownFolders]: A list of known folders that the domain can replicate.
//   - [NSFileProviderDomain.SetSupportedKnownFolders]
//
// # Supporting search
//
//   - [NSFileProviderDomain.SupportsStringSearchRequest]: A Boolean value that indicates whether the provider supports search.
//   - [NSFileProviderDomain.SetSupportsStringSearchRequest]
//
// # Testing
//
//   - [NSFileProviderDomain.TestingModes]: A mode that gives the File Provider extension more control over the system’s behavior during testing.
//   - [NSFileProviderDomain.SetTestingModes]
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderDomain
//
// [NSFileProviderExtension]: https://developer.apple.com/documentation/FileProvider/NSFileProviderExtension
// [domain]: https://developer.apple.com/documentation/FileProvider/NSFileProviderExtension/domain
type NSFileProviderDomain struct {
	objectivec.Object
}

// NSFileProviderDomainFromID constructs a [NSFileProviderDomain] from an objc.ID.
//
// A File Provider extension’s domain.
func NSFileProviderDomainFromID(id objc.ID) NSFileProviderDomain {
	return NSFileProviderDomain{objectivec.Object{ID: id}}
}

// NOTE: NSFileProviderDomain adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSFileProviderDomain] class.
//
// # Creating domains
//
//   - [INSFileProviderDomain.InitWithIdentifierDisplayName]: Creates a new file provider domain with the specified identifier and display name.
//   - [INSFileProviderDomain.InitWithDisplayNameUserInfoVolumeURL]: Creates a new file provider domain with the specified URL and display name.
//
// # Accessing data
//
//   - [INSFileProviderDomain.DisplayName]: The name of the domain displayed in the user interface.
//   - [INSFileProviderDomain.Identifier]: The domain’s unique identifier.
//   - [INSFileProviderDomain.IsReplicated]
//   - [INSFileProviderDomain.BackingStoreIdentity]: A unique identifier for the backing store used by the system.
//   - [INSFileProviderDomain.IsHidden]: A Boolean value that determines whether the domain is visible to users.
//   - [INSFileProviderDomain.SetHidden]
//   - [INSFileProviderDomain.UserEnabled]: A Boolean value that indicates whether the user has enabled or disabled the domain.
//   - [INSFileProviderDomain.IsDisconnected]: A Boolean value indicating that the domain is present, but disconnected from the file extension.
//   - [INSFileProviderDomain.SupportsSyncingTrash]
//   - [INSFileProviderDomain.SetSupportsSyncingTrash]
//   - [INSFileProviderDomain.UserInfo]
//   - [INSFileProviderDomain.SetUserInfo]
//   - [INSFileProviderDomain.VolumeUUID]
//
// # Syncing Desktop and Documents folders
//
//   - [INSFileProviderDomain.ReplicatedKnownFolders]: A list of known folders that the domain currently replicates.
//   - [INSFileProviderDomain.SupportedKnownFolders]: A list of known folders that the domain can replicate.
//   - [INSFileProviderDomain.SetSupportedKnownFolders]
//
// # Supporting search
//
//   - [INSFileProviderDomain.SupportsStringSearchRequest]: A Boolean value that indicates whether the provider supports search.
//   - [INSFileProviderDomain.SetSupportsStringSearchRequest]
//
// # Testing
//
//   - [INSFileProviderDomain.TestingModes]: A mode that gives the File Provider extension more control over the system’s behavior during testing.
//   - [INSFileProviderDomain.SetTestingModes]
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderDomain
type INSFileProviderDomain interface {
	objectivec.IObject

	// Topic: Creating domains

	// Creates a new file provider domain with the specified identifier and display name.
	InitWithIdentifierDisplayName(identifier NSFileProviderDomainIdentifier, displayName string) NSFileProviderDomain
	// Creates a new file provider domain with the specified URL and display name.
	InitWithDisplayNameUserInfoVolumeURL(displayName string, userInfo foundation.INSDictionary, volumeURL foundation.NSURL) NSFileProviderDomain

	// Topic: Accessing data

	// The name of the domain displayed in the user interface.
	DisplayName() string
	// The domain’s unique identifier.
	Identifier() NSFileProviderDomainIdentifier
	IsReplicated() bool
	// A unique identifier for the backing store used by the system.
	BackingStoreIdentity() foundation.NSData
	// A Boolean value that determines whether the domain is visible to users.
	IsHidden() bool
	SetHidden(value bool)
	// A Boolean value that indicates whether the user has enabled or disabled the domain.
	UserEnabled() bool
	// A Boolean value indicating that the domain is present, but disconnected from the file extension.
	IsDisconnected() bool
	SupportsSyncingTrash() bool
	SetSupportsSyncingTrash(value bool)
	UserInfo() foundation.INSDictionary
	SetUserInfo(value foundation.INSDictionary)
	VolumeUUID() foundation.NSUUID

	// Topic: Syncing Desktop and Documents folders

	// A list of known folders that the domain currently replicates.
	ReplicatedKnownFolders() NSFileProviderKnownFolders
	// A list of known folders that the domain can replicate.
	SupportedKnownFolders() NSFileProviderKnownFolders
	SetSupportedKnownFolders(value NSFileProviderKnownFolders)

	// Topic: Supporting search

	// A Boolean value that indicates whether the provider supports search.
	SupportsStringSearchRequest() bool
	SetSupportsStringSearchRequest(value bool)

	// Topic: Testing

	// A mode that gives the File Provider extension more control over the system’s behavior during testing.
	TestingModes() NSFileProviderDomainTestingModes
	SetTestingModes(value NSFileProviderDomainTestingModes)
}

// Init initializes the instance.
func (f NSFileProviderDomain) Init() NSFileProviderDomain {
	rv := objc.Send[NSFileProviderDomain](f.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f NSFileProviderDomain) Autorelease() NSFileProviderDomain {
	rv := objc.Send[NSFileProviderDomain](f.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSFileProviderDomain creates a new NSFileProviderDomain instance.
func NewNSFileProviderDomain() NSFileProviderDomain {
	class := getNSFileProviderDomainClass()
	rv := objc.Send[NSFileProviderDomain](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a new file provider domain with the specified URL and display name.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderDomain/init(displayName:userInfo:volumeURL:)
func NewFileProviderDomainWithDisplayNameUserInfoVolumeURL(displayName string, userInfo foundation.INSDictionary, volumeURL foundation.NSURL) NSFileProviderDomain {
	instance := getNSFileProviderDomainClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDisplayName:userInfo:volumeURL:"), objc.String(displayName), userInfo, volumeURL)
	return NSFileProviderDomainFromID(rv)
}

// Creates a new file provider domain with the specified identifier and
// display name.
//
// identifier: A string that identifies the domain. The file provider extension can select
// any string to uniquely identify the domain, as long as it doesn’t contain
// the colon (:) or slash (/) symbols.
//
// displayName: The name for the domain that the system shows to the user.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderDomain/init(identifier:displayName:)
func NewFileProviderDomainWithIdentifierDisplayName(identifier NSFileProviderDomainIdentifier, displayName string) NSFileProviderDomain {
	instance := getNSFileProviderDomainClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithIdentifier:displayName:"), objc.String(string(identifier)), objc.String(displayName))
	return NSFileProviderDomainFromID(rv)
}

// Returns a newly instantiated domain.
//
// identifier: A string that identifies the domain. The file provider extension can select
// any string to uniquely identify the domain, as long as it doesn’t contain
// the colon (:) or slash (/) symbols.
//
// displayName: The name for the domain that the system shows to the user.
//
// pathRelativeToDocumentStorage: A path relative to the file provider extension’s [documentStorageURL]
// that the system uses to store the domain’s content.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderDomain/init(identifier:displayName:pathRelativeToDocumentStorage:)
//
// [documentStorageURL]: https://developer.apple.com/documentation/FileProvider/NSFileProviderExtension/documentStorageURL
func NewFileProviderDomainWithIdentifierDisplayNamePathRelativeToDocumentStorage(identifier NSFileProviderDomainIdentifier, displayName string, pathRelativeToDocumentStorage string) NSFileProviderDomain {
	instance := getNSFileProviderDomainClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithIdentifier:displayName:pathRelativeToDocumentStorage:"), objc.String(string(identifier)), objc.String(displayName), objc.String(pathRelativeToDocumentStorage))
	return NSFileProviderDomainFromID(rv)
}

// Creates a new file provider domain with the specified identifier and
// display name.
//
// identifier: A string that identifies the domain. The file provider extension can select
// any string to uniquely identify the domain, as long as it doesn’t contain
// the colon (:) or slash (/) symbols.
//
// displayName: The name for the domain that the system shows to the user.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderDomain/init(identifier:displayName:)
func (f NSFileProviderDomain) InitWithIdentifierDisplayName(identifier NSFileProviderDomainIdentifier, displayName string) NSFileProviderDomain {
	rv := objc.Send[NSFileProviderDomain](f.ID, objc.Sel("initWithIdentifier:displayName:"), objc.String(string(identifier)), objc.String(displayName))
	return rv
}

// Creates a new file provider domain with the specified URL and display name.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderDomain/init(displayName:userInfo:volumeURL:)
func (f NSFileProviderDomain) InitWithDisplayNameUserInfoVolumeURL(displayName string, userInfo foundation.INSDictionary, volumeURL foundation.NSURL) NSFileProviderDomain {
	rv := objc.Send[NSFileProviderDomain](f.ID, objc.Sel("initWithDisplayName:userInfo:volumeURL:"), objc.String(displayName), userInfo, volumeURL)
	return rv
}

// The name of the domain displayed in the user interface.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderDomain/displayName
func (f NSFileProviderDomain) DisplayName() string {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("displayName"))
	return foundation.NSStringFromID(rv).String()
}

// The domain’s unique identifier.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderDomain/identifier
func (f NSFileProviderDomain) Identifier() NSFileProviderDomainIdentifier {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("identifier"))
	return NSFileProviderDomainIdentifier(foundation.NSStringFromID(rv).String())
}

// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderDomain/isReplicated
func (f NSFileProviderDomain) IsReplicated() bool {
	rv := objc.Send[bool](f.ID, objc.Sel("isReplicated"))
	return rv
}

// A unique identifier for the backing store used by the system.
//
// # Discussion
//
// Changes to this identifier indicate that the system has dropped its backing
// store and is creating a new one. The system may create a new backing store
// if the old store becomes corrupted. The file provider extension can also
// request a new backing store by calling
// [NSFileProviderManager.ReimportItemsBelowItemWithIdentifierCompletionHandler].
//
// While rebuilding the backing store, the system invalidates any extension
// instances associated with the domain. As a result, the system guarantees
// that the [NSFileProviderDomain.BackingStoreIdentity] remains stable
// throughout the lifetime of an [NSFileProviderReplicatedExtension] instance.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderDomain/backingStoreIdentity
func (f NSFileProviderDomain) BackingStoreIdentity() foundation.NSData {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("backingStoreIdentity"))
	return foundation.NSDataFromID(objc.ID(rv))
}

// A Boolean value that determines whether the domain is visible to users.
//
// # Discussion
//
// The system stores the files on disk, but it doesn’t display them to the
// user. For example, you could set this value to false when performing a dry
// run of a migration.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderDomain/isHidden
func (f NSFileProviderDomain) IsHidden() bool {
	rv := objc.Send[bool](f.ID, objc.Sel("isHidden"))
	return rv
}
func (f NSFileProviderDomain) SetHidden(value bool) {
	objc.Send[struct{}](f.ID, objc.Sel("setHidden:"), value)
}

// A Boolean value that indicates whether the user has enabled or disabled the
// domain.
//
// # Discussion
//
// By default, the property is true; however, If the user disables the domain
// in the System Preferences, the property is false.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderDomain/userEnabled
func (f NSFileProviderDomain) UserEnabled() bool {
	rv := objc.Send[bool](f.ID, objc.Sel("userEnabled"))
	return rv
}

// A Boolean value indicating that the domain is present, but disconnected
// from the file extension.
//
// # Discussion
//
// Users can continue to browse the content from a disconnected domain;
// however, the File Provider extension no longer sends or receives updates
// about modifications to the files.
//
// To change the disconnected state, create a new [NSFileProviderDomain] using
// the same identifier, and pass it to
// [NSFileProviderManagerClass.AddDomainCompletionHandler].
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderDomain/isDisconnected
func (f NSFileProviderDomain) IsDisconnected() bool {
	rv := objc.Send[bool](f.ID, objc.Sel("isDisconnected"))
	return rv
}

// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderDomain/supportsSyncingTrash
func (f NSFileProviderDomain) SupportsSyncingTrash() bool {
	rv := objc.Send[bool](f.ID, objc.Sel("supportsSyncingTrash"))
	return rv
}
func (f NSFileProviderDomain) SetSupportsSyncingTrash(value bool) {
	objc.Send[struct{}](f.ID, objc.Sel("setSupportsSyncingTrash:"), value)
}

// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderDomain/userInfo
func (f NSFileProviderDomain) UserInfo() foundation.INSDictionary {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("userInfo"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (f NSFileProviderDomain) SetUserInfo(value foundation.INSDictionary) {
	objc.Send[struct{}](f.ID, objc.Sel("setUserInfo:"), value)
}

// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderDomain/volumeUUID
func (f NSFileProviderDomain) VolumeUUID() foundation.NSUUID {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("volumeUUID"))
	return foundation.NSUUIDFromID(objc.ID(rv))
}

// A list of known folders that the domain currently replicates.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderDomain/replicatedKnownFolders
func (f NSFileProviderDomain) ReplicatedKnownFolders() NSFileProviderKnownFolders {
	rv := objc.Send[NSFileProviderKnownFolders](f.ID, objc.Sel("replicatedKnownFolders"))
	return NSFileProviderKnownFolders(rv)
}

// A list of known folders that the domain can replicate.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderDomain/supportedKnownFolders
func (f NSFileProviderDomain) SupportedKnownFolders() NSFileProviderKnownFolders {
	rv := objc.Send[NSFileProviderKnownFolders](f.ID, objc.Sel("supportedKnownFolders"))
	return NSFileProviderKnownFolders(rv)
}
func (f NSFileProviderDomain) SetSupportedKnownFolders(value NSFileProviderKnownFolders) {
	objc.Send[struct{}](f.ID, objc.Sel("setSupportedKnownFolders:"), value)
}

// A Boolean value that indicates whether the provider supports search.
//
// # Discussion
//
// If this value is `true`, the framework uses the extension’s
// [NSFileProviderSearching] implementation to support search.
//
// The property defaults to `false` (Swift) or [NO] (Objective-C).
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderDomain/supportsStringSearchRequest
func (f NSFileProviderDomain) SupportsStringSearchRequest() bool {
	rv := objc.Send[bool](f.ID, objc.Sel("supportsStringSearchRequest"))
	return rv
}
func (f NSFileProviderDomain) SetSupportsStringSearchRequest(value bool) {
	objc.Send[struct{}](f.ID, objc.Sel("setSupportsStringSearchRequest:"), value)
}

// A mode that gives the File Provider extension more control over the
// system’s behavior during testing.
//
// # Discussion
//
// By default, the value is `[]` (Swift) or `0` (Objective-C) and all testing
// modes are disabled. To enable a testing mode, assign its value to this
// property. You can combine multiple modes:
//
// The system registers the domain’s testing mode when you add the domain by
// calling [NSFileProviderManagerClass.AddDomainCompletionHandler]. You
// can’t change the test mode after you add the domain.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderDomain/testingModes-swift.property
func (f NSFileProviderDomain) TestingModes() NSFileProviderDomainTestingModes {
	rv := objc.Send[NSFileProviderDomainTestingModes](f.ID, objc.Sel("testingModes"))
	return NSFileProviderDomainTestingModes(rv)
}
func (f NSFileProviderDomain) SetTestingModes(value NSFileProviderDomainTestingModes) {
	objc.Send[struct{}](f.ID, objc.Sel("setTestingModes:"), value)
}
