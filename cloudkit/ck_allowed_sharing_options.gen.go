// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CKAllowedSharingOptions] class.
var (
	_CKAllowedSharingOptionsClass     CKAllowedSharingOptionsClass
	_CKAllowedSharingOptionsClassOnce sync.Once
)

func getCKAllowedSharingOptionsClass() CKAllowedSharingOptionsClass {
	_CKAllowedSharingOptionsClassOnce.Do(func() {
		_CKAllowedSharingOptionsClass = CKAllowedSharingOptionsClass{class: objc.GetClass("CKAllowedSharingOptions")}
	})
	return _CKAllowedSharingOptionsClass
}

// GetCKAllowedSharingOptionsClass returns the class object for CKAllowedSharingOptions.
func GetCKAllowedSharingOptionsClass() CKAllowedSharingOptionsClass {
	return getCKAllowedSharingOptionsClass()
}

type CKAllowedSharingOptionsClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CKAllowedSharingOptionsClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CKAllowedSharingOptionsClass) Alloc() CKAllowedSharingOptions {
	rv := objc.Send[CKAllowedSharingOptions](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// An object that controls participant access and permission options.
//
// # Overview
//
// Register an instance of this class with an [NSItemProvider] or when
// preparing a [CKShareTransferRepresentation.ExportedShare] before your app
// invokes the share sheet. The share sheet uses the registered
// [CKAllowedSharingOptions] object to let the user choose between the allowed
// options when sharing.
//
// # Creating sharing options
//
//   - [CKAllowedSharingOptions.InitWithAllowedParticipantPermissionOptionsAllowedParticipantAccessOptions]: Creates and initializes an allowed sharing options object.
//
// # Configuring the options
//
//   - [CKAllowedSharingOptions.AllowedParticipantAccessOptions]: The permission option the system uses to control whether a user can share publicly or privately.
//   - [CKAllowedSharingOptions.SetAllowedParticipantAccessOptions]
//   - [CKAllowedSharingOptions.AllowedParticipantPermissionOptions]: The permission option the system uses to control whether a user can grant read-only or write access.
//   - [CKAllowedSharingOptions.SetAllowedParticipantPermissionOptions]
//
// # Initializers
//
//   - [CKAllowedSharingOptions.InitWithCoder]
//
// # Instance Properties
//
//   - [CKAllowedSharingOptions.AllowsAccessRequests]: Default value is [NO]. If set, the system sharing UI will allow the user to configure whether access requests are enabled on the share.
//   - [CKAllowedSharingOptions.SetAllowsAccessRequests]
//   - [CKAllowedSharingOptions.AllowsParticipantsToInviteOthers]: Default value is [NO]. If set, the system sharing UI will allow the user to choose whether added participants can invite others to the share. Shares with [CKShare.ParticipantRole.administrator](<https://developer.apple.com/documentation/CloudKit/CKShare/ParticipantRole/administrator>) participants will be returned as read-only to devices running OS versions prior to this role being introduced. Administrator participants on these read-only shares will be returned as [CKShare.ParticipantRole.privateUser](<https://developer.apple.com/documentation/CloudKit/CKShare/ParticipantRole/privateUser>).
//   - [CKAllowedSharingOptions.SetAllowsParticipantsToInviteOthers]
//
// See: https://developer.apple.com/documentation/CloudKit/CKAllowedSharingOptions
//
// [CKShareTransferRepresentation.ExportedShare]: https://developer.apple.com/documentation/CloudKit/CKShareTransferRepresentation/ExportedShare
// [NSItemProvider]: https://developer.apple.com/documentation/Foundation/NSItemProvider
type CKAllowedSharingOptions struct {
	objectivec.Object
}

// CKAllowedSharingOptionsFromID constructs a [CKAllowedSharingOptions] from an objc.ID.
//
// An object that controls participant access and permission options.
func CKAllowedSharingOptionsFromID(id objc.ID) CKAllowedSharingOptions {
	return CKAllowedSharingOptions{objectivec.Object{ID: id}}
}

// NOTE: CKAllowedSharingOptions adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CKAllowedSharingOptions] class.
//
// # Creating sharing options
//
//   - [ICKAllowedSharingOptions.InitWithAllowedParticipantPermissionOptionsAllowedParticipantAccessOptions]: Creates and initializes an allowed sharing options object.
//
// # Configuring the options
//
//   - [ICKAllowedSharingOptions.AllowedParticipantAccessOptions]: The permission option the system uses to control whether a user can share publicly or privately.
//   - [ICKAllowedSharingOptions.SetAllowedParticipantAccessOptions]
//   - [ICKAllowedSharingOptions.AllowedParticipantPermissionOptions]: The permission option the system uses to control whether a user can grant read-only or write access.
//   - [ICKAllowedSharingOptions.SetAllowedParticipantPermissionOptions]
//
// # Initializers
//
//   - [ICKAllowedSharingOptions.InitWithCoder]
//
// # Instance Properties
//
//   - [ICKAllowedSharingOptions.AllowsAccessRequests]: Default value is [NO]. If set, the system sharing UI will allow the user to configure whether access requests are enabled on the share.
//   - [ICKAllowedSharingOptions.SetAllowsAccessRequests]
//   - [ICKAllowedSharingOptions.AllowsParticipantsToInviteOthers]: Default value is [NO]. If set, the system sharing UI will allow the user to choose whether added participants can invite others to the share. Shares with [CKShare.ParticipantRole.administrator](<https://developer.apple.com/documentation/CloudKit/CKShare/ParticipantRole/administrator>) participants will be returned as read-only to devices running OS versions prior to this role being introduced. Administrator participants on these read-only shares will be returned as [CKShare.ParticipantRole.privateUser](<https://developer.apple.com/documentation/CloudKit/CKShare/ParticipantRole/privateUser>).
//   - [ICKAllowedSharingOptions.SetAllowsParticipantsToInviteOthers]
//
// See: https://developer.apple.com/documentation/CloudKit/CKAllowedSharingOptions
type ICKAllowedSharingOptions interface {
	objectivec.IObject

	// Topic: Creating sharing options

	// Creates and initializes an allowed sharing options object.
	InitWithAllowedParticipantPermissionOptionsAllowedParticipantAccessOptions(allowedParticipantPermissionOptions CKSharingParticipantPermissionOption, allowedParticipantAccessOptions CKSharingParticipantAccessOption) CKAllowedSharingOptions

	// Topic: Configuring the options

	// The permission option the system uses to control whether a user can share publicly or privately.
	AllowedParticipantAccessOptions() CKSharingParticipantAccessOption
	SetAllowedParticipantAccessOptions(value CKSharingParticipantAccessOption)
	// The permission option the system uses to control whether a user can grant read-only or write access.
	AllowedParticipantPermissionOptions() CKSharingParticipantPermissionOption
	SetAllowedParticipantPermissionOptions(value CKSharingParticipantPermissionOption)

	// Topic: Initializers

	InitWithCoder(coder foundation.INSCoder) CKAllowedSharingOptions

	// Topic: Instance Properties

	// Default value is [NO]. If set, the system sharing UI will allow the user to configure whether access requests are enabled on the share.
	AllowsAccessRequests() bool
	SetAllowsAccessRequests(value bool)
	// Default value is [NO]. If set, the system sharing UI will allow the user to choose whether added participants can invite others to the share. Shares with [CKShare.ParticipantRole.administrator](<https://developer.apple.com/documentation/CloudKit/CKShare/ParticipantRole/administrator>) participants will be returned as read-only to devices running OS versions prior to this role being introduced. Administrator participants on these read-only shares will be returned as [CKShare.ParticipantRole.privateUser](<https://developer.apple.com/documentation/CloudKit/CKShare/ParticipantRole/privateUser>).
	AllowsParticipantsToInviteOthers() bool
	SetAllowsParticipantsToInviteOthers(value bool)

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (c CKAllowedSharingOptions) Init() CKAllowedSharingOptions {
	rv := objc.Send[CKAllowedSharingOptions](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CKAllowedSharingOptions) Autorelease() CKAllowedSharingOptions {
	rv := objc.Send[CKAllowedSharingOptions](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKAllowedSharingOptions creates a new CKAllowedSharingOptions instance.
func NewCKAllowedSharingOptions() CKAllowedSharingOptions {
	class := getCKAllowedSharingOptionsClass()
	rv := objc.Send[CKAllowedSharingOptions](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates and initializes an allowed sharing options object.
//
// allowedParticipantPermissionOptions: The [CKSharingParticipantPermissionOption] setting.
//
// allowedParticipantAccessOptions: The [CKSharingParticipantAccessOption] setting.
//
// See: https://developer.apple.com/documentation/CloudKit/CKAllowedSharingOptions/init(allowedParticipantPermissionOptions:allowedParticipantAccessOptions:)
//
// [CKSharingParticipantPermissionOption]: https://developer.apple.com/documentation/CloudKit/CKSharingParticipantPermissionOption
// [CKSharingParticipantAccessOption]: https://developer.apple.com/documentation/CloudKit/CKSharingParticipantAccessOption
func NewCKAllowedSharingOptionsWithAllowedParticipantPermissionOptionsAllowedParticipantAccessOptions(allowedParticipantPermissionOptions CKSharingParticipantPermissionOption, allowedParticipantAccessOptions CKSharingParticipantAccessOption) CKAllowedSharingOptions {
	instance := getCKAllowedSharingOptionsClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithAllowedParticipantPermissionOptions:allowedParticipantAccessOptions:"), allowedParticipantPermissionOptions, allowedParticipantAccessOptions)
	return CKAllowedSharingOptionsFromID(rv)
}

// See: https://developer.apple.com/documentation/CloudKit/CKAllowedSharingOptions/init(coder:)
func NewCKAllowedSharingOptionsWithCoder(coder foundation.INSCoder) CKAllowedSharingOptions {
	instance := getCKAllowedSharingOptionsClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return CKAllowedSharingOptionsFromID(rv)
}

// Creates and initializes an allowed sharing options object.
//
// allowedParticipantPermissionOptions: The [CKSharingParticipantPermissionOption] setting.
//
// allowedParticipantAccessOptions: The [CKSharingParticipantAccessOption] setting.
//
// See: https://developer.apple.com/documentation/CloudKit/CKAllowedSharingOptions/init(allowedParticipantPermissionOptions:allowedParticipantAccessOptions:)
//
// [CKSharingParticipantPermissionOption]: https://developer.apple.com/documentation/CloudKit/CKSharingParticipantPermissionOption
// [CKSharingParticipantAccessOption]: https://developer.apple.com/documentation/CloudKit/CKSharingParticipantAccessOption
func (c CKAllowedSharingOptions) InitWithAllowedParticipantPermissionOptionsAllowedParticipantAccessOptions(allowedParticipantPermissionOptions CKSharingParticipantPermissionOption, allowedParticipantAccessOptions CKSharingParticipantAccessOption) CKAllowedSharingOptions {
	rv := objc.Send[CKAllowedSharingOptions](c.ID, objc.Sel("initWithAllowedParticipantPermissionOptions:allowedParticipantAccessOptions:"), allowedParticipantPermissionOptions, allowedParticipantAccessOptions)
	return rv
}

// See: https://developer.apple.com/documentation/CloudKit/CKAllowedSharingOptions/init(coder:)
func (c CKAllowedSharingOptions) InitWithCoder(coder foundation.INSCoder) CKAllowedSharingOptions {
	rv := objc.Send[CKAllowedSharingOptions](c.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}
func (c CKAllowedSharingOptions) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](c.ID, objc.Sel("encodeWithCoder:"), coder)
}

// The permission option the system uses to control whether a user can share
// publicly or privately.
//
// See: https://developer.apple.com/documentation/CloudKit/CKAllowedSharingOptions/allowedParticipantAccessOptions
func (c CKAllowedSharingOptions) AllowedParticipantAccessOptions() CKSharingParticipantAccessOption {
	rv := objc.Send[CKSharingParticipantAccessOption](c.ID, objc.Sel("allowedParticipantAccessOptions"))
	return CKSharingParticipantAccessOption(rv)
}
func (c CKAllowedSharingOptions) SetAllowedParticipantAccessOptions(value CKSharingParticipantAccessOption) {
	objc.Send[struct{}](c.ID, objc.Sel("setAllowedParticipantAccessOptions:"), value)
}

// The permission option the system uses to control whether a user can grant
// read-only or write access.
//
// See: https://developer.apple.com/documentation/CloudKit/CKAllowedSharingOptions/allowedParticipantPermissionOptions
func (c CKAllowedSharingOptions) AllowedParticipantPermissionOptions() CKSharingParticipantPermissionOption {
	rv := objc.Send[CKSharingParticipantPermissionOption](c.ID, objc.Sel("allowedParticipantPermissionOptions"))
	return CKSharingParticipantPermissionOption(rv)
}
func (c CKAllowedSharingOptions) SetAllowedParticipantPermissionOptions(value CKSharingParticipantPermissionOption) {
	objc.Send[struct{}](c.ID, objc.Sel("setAllowedParticipantPermissionOptions:"), value)
}

// Default value is [NO]. If set, the system sharing UI will allow the user to
// configure whether access requests are enabled on the share.
//
// See: https://developer.apple.com/documentation/CloudKit/CKAllowedSharingOptions/allowsAccessRequests
func (c CKAllowedSharingOptions) AllowsAccessRequests() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("allowsAccessRequests"))
	return rv
}
func (c CKAllowedSharingOptions) SetAllowsAccessRequests(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setAllowsAccessRequests:"), value)
}

// Default value is [NO]. If set, the system sharing UI will allow the user to
// choose whether added participants can invite others to the share. Shares
// with [CKShare.ParticipantRole.administrator] participants will be returned
// as read-only to devices running OS versions prior to this role being
// introduced. Administrator participants on these read-only shares will be
// returned as [CKShare.ParticipantRole.privateUser].
//
// See: https://developer.apple.com/documentation/CloudKit/CKAllowedSharingOptions/allowsParticipantsToInviteOthers
//
// [CKShare.ParticipantRole.administrator]: https://developer.apple.com/documentation/CloudKit/CKShare/ParticipantRole/administrator
// [CKShare.ParticipantRole.privateUser]: https://developer.apple.com/documentation/CloudKit/CKShare/ParticipantRole/privateUser
func (c CKAllowedSharingOptions) AllowsParticipantsToInviteOthers() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("allowsParticipantsToInviteOthers"))
	return rv
}
func (c CKAllowedSharingOptions) SetAllowsParticipantsToInviteOthers(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setAllowsParticipantsToInviteOthers:"), value)
}

// An object set to the most permissive sharing options.
//
// # Discussion
//
// The `standardOptions` has
// [CKAllowedSharingOptions.AllowedParticipantPermissionOptions] set to
// [CKSharingParticipantPermissionOptionAny] and
// [CKAllowedSharingOptions.AllowedParticipantAccessOptions] set to
// [CKSharingParticipantAccessOptionAny].
//
// See: https://developer.apple.com/documentation/CloudKit/CKAllowedSharingOptions/standard
func (_CKAllowedSharingOptionsClass CKAllowedSharingOptionsClass) StandardOptions() CKAllowedSharingOptions {
	rv := objc.Send[objc.ID](objc.ID(_CKAllowedSharingOptionsClass.class), objc.Sel("standardOptions"))
	return CKAllowedSharingOptionsFromID(objc.ID(rv))
}
