// Code generated from Apple documentation for OpenDirectory. DO NOT EDIT.

package opendirectory

import (
	"errors"
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [ODConfiguration] class.
var (
	_ODConfigurationClass     ODConfigurationClass
	_ODConfigurationClassOnce sync.Once
)

func getODConfigurationClass() ODConfigurationClass {
	_ODConfigurationClassOnce.Do(func() {
		_ODConfigurationClass = ODConfigurationClass{class: objc.GetClass("ODConfiguration")}
	})
	return _ODConfigurationClass
}

// GetODConfigurationClass returns the class object for ODConfiguration.
func GetODConfigurationClass() ODConfigurationClass {
	return getODConfigurationClass()
}

type ODConfigurationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (oc ODConfigurationClass) Class() objc.Class {
	return oc.class
}

// Alloc allocates memory for a new instance of the class.
func (oc ODConfigurationClass) Alloc() ODConfiguration {
	rv := objc.Send[ODConfiguration](objc.ID(oc.class), objc.Sel("alloc"))
	return rv
}

// # Instance Properties
//
//   - [ODConfiguration.AuthenticationModuleEntries]
//   - [ODConfiguration.SetAuthenticationModuleEntries]
//   - [ODConfiguration.Comment]
//   - [ODConfiguration.SetComment]
//   - [ODConfiguration.ConnectionIdleTimeoutInSeconds]
//   - [ODConfiguration.SetConnectionIdleTimeoutInSeconds]
//   - [ODConfiguration.ConnectionSetupTimeoutInSeconds]
//   - [ODConfiguration.SetConnectionSetupTimeoutInSeconds]
//   - [ODConfiguration.DefaultMappings]
//   - [ODConfiguration.SetDefaultMappings]
//   - [ODConfiguration.DefaultModuleEntries]
//   - [ODConfiguration.SetDefaultModuleEntries]
//   - [ODConfiguration.DiscoveryModuleEntries]
//   - [ODConfiguration.SetDiscoveryModuleEntries]
//   - [ODConfiguration.GeneralModuleEntries]
//   - [ODConfiguration.SetGeneralModuleEntries]
//   - [ODConfiguration.HideRegistration]
//   - [ODConfiguration.SetHideRegistration]
//   - [ODConfiguration.ManInTheMiddleProtection]
//   - [ODConfiguration.SetManInTheMiddleProtection]
//   - [ODConfiguration.NodeName]
//   - [ODConfiguration.SetNodeName]
//   - [ODConfiguration.PacketEncryption]
//   - [ODConfiguration.SetPacketEncryption]
//   - [ODConfiguration.PacketSigning]
//   - [ODConfiguration.SetPacketSigning]
//   - [ODConfiguration.PreferredDestinationHostName]
//   - [ODConfiguration.SetPreferredDestinationHostName]
//   - [ODConfiguration.PreferredDestinationHostPort]
//   - [ODConfiguration.SetPreferredDestinationHostPort]
//   - [ODConfiguration.QueryTimeoutInSeconds]
//   - [ODConfiguration.SetQueryTimeoutInSeconds]
//   - [ODConfiguration.TemplateName]
//   - [ODConfiguration.SetTemplateName]
//   - [ODConfiguration.TrustAccount]
//   - [ODConfiguration.TrustKerberosPrincipal]
//   - [ODConfiguration.TrustMetaAccount]
//   - [ODConfiguration.TrustType]
//   - [ODConfiguration.TrustUsesKerberosKeytab]
//   - [ODConfiguration.TrustUsesMutualAuthentication]
//   - [ODConfiguration.TrustUsesSystemKeychain]
//   - [ODConfiguration.VirtualSubnodes]
//   - [ODConfiguration.SetVirtualSubnodes]
//
// # Instance Methods
//
//   - [ODConfiguration.AddTrustTypeTrustAccountTrustPasswordUsernamePasswordJoinExistingError]
//   - [ODConfiguration.RemoveTrustUsingUsernamePasswordDeleteTrustAccountError]
//   - [ODConfiguration.SaveUsingAuthorizationError]
//
// See: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration
type ODConfiguration struct {
	objectivec.Object
}

// ODConfigurationFromID constructs a [ODConfiguration] from an objc.ID.
func ODConfigurationFromID(id objc.ID) ODConfiguration {
	return ODConfiguration{objectivec.Object{ID: id}}
}

// NOTE: ODConfiguration adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [ODConfiguration] class.
//
// # Instance Properties
//
//   - [IODConfiguration.AuthenticationModuleEntries]
//   - [IODConfiguration.SetAuthenticationModuleEntries]
//   - [IODConfiguration.Comment]
//   - [IODConfiguration.SetComment]
//   - [IODConfiguration.ConnectionIdleTimeoutInSeconds]
//   - [IODConfiguration.SetConnectionIdleTimeoutInSeconds]
//   - [IODConfiguration.ConnectionSetupTimeoutInSeconds]
//   - [IODConfiguration.SetConnectionSetupTimeoutInSeconds]
//   - [IODConfiguration.DefaultMappings]
//   - [IODConfiguration.SetDefaultMappings]
//   - [IODConfiguration.DefaultModuleEntries]
//   - [IODConfiguration.SetDefaultModuleEntries]
//   - [IODConfiguration.DiscoveryModuleEntries]
//   - [IODConfiguration.SetDiscoveryModuleEntries]
//   - [IODConfiguration.GeneralModuleEntries]
//   - [IODConfiguration.SetGeneralModuleEntries]
//   - [IODConfiguration.HideRegistration]
//   - [IODConfiguration.SetHideRegistration]
//   - [IODConfiguration.ManInTheMiddleProtection]
//   - [IODConfiguration.SetManInTheMiddleProtection]
//   - [IODConfiguration.NodeName]
//   - [IODConfiguration.SetNodeName]
//   - [IODConfiguration.PacketEncryption]
//   - [IODConfiguration.SetPacketEncryption]
//   - [IODConfiguration.PacketSigning]
//   - [IODConfiguration.SetPacketSigning]
//   - [IODConfiguration.PreferredDestinationHostName]
//   - [IODConfiguration.SetPreferredDestinationHostName]
//   - [IODConfiguration.PreferredDestinationHostPort]
//   - [IODConfiguration.SetPreferredDestinationHostPort]
//   - [IODConfiguration.QueryTimeoutInSeconds]
//   - [IODConfiguration.SetQueryTimeoutInSeconds]
//   - [IODConfiguration.TemplateName]
//   - [IODConfiguration.SetTemplateName]
//   - [IODConfiguration.TrustAccount]
//   - [IODConfiguration.TrustKerberosPrincipal]
//   - [IODConfiguration.TrustMetaAccount]
//   - [IODConfiguration.TrustType]
//   - [IODConfiguration.TrustUsesKerberosKeytab]
//   - [IODConfiguration.TrustUsesMutualAuthentication]
//   - [IODConfiguration.TrustUsesSystemKeychain]
//   - [IODConfiguration.VirtualSubnodes]
//   - [IODConfiguration.SetVirtualSubnodes]
//
// # Instance Methods
//
//   - [IODConfiguration.AddTrustTypeTrustAccountTrustPasswordUsernamePasswordJoinExistingError]
//   - [IODConfiguration.RemoveTrustUsingUsernamePasswordDeleteTrustAccountError]
//   - [IODConfiguration.SaveUsingAuthorizationError]
//
// See: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration
type IODConfiguration interface {
	objectivec.IObject

	// Topic: Instance Properties

	AuthenticationModuleEntries() foundation.INSArray
	SetAuthenticationModuleEntries(value foundation.INSArray)
	Comment() string
	SetComment(value string)
	ConnectionIdleTimeoutInSeconds() int
	SetConnectionIdleTimeoutInSeconds(value int)
	ConnectionSetupTimeoutInSeconds() int
	SetConnectionSetupTimeoutInSeconds(value int)
	DefaultMappings() IODMappings
	SetDefaultMappings(value IODMappings)
	DefaultModuleEntries() foundation.INSArray
	SetDefaultModuleEntries(value foundation.INSArray)
	DiscoveryModuleEntries() foundation.INSArray
	SetDiscoveryModuleEntries(value foundation.INSArray)
	GeneralModuleEntries() foundation.INSArray
	SetGeneralModuleEntries(value foundation.INSArray)
	HideRegistration() bool
	SetHideRegistration(value bool)
	ManInTheMiddleProtection() bool
	SetManInTheMiddleProtection(value bool)
	NodeName() string
	SetNodeName(value string)
	PacketEncryption() int
	SetPacketEncryption(value int)
	PacketSigning() int
	SetPacketSigning(value int)
	PreferredDestinationHostName() string
	SetPreferredDestinationHostName(value string)
	PreferredDestinationHostPort() uint16
	SetPreferredDestinationHostPort(value uint16)
	QueryTimeoutInSeconds() int
	SetQueryTimeoutInSeconds(value int)
	TemplateName() string
	SetTemplateName(value string)
	TrustAccount() string
	TrustKerberosPrincipal() string
	TrustMetaAccount() string
	TrustType() string
	TrustUsesKerberosKeytab() bool
	TrustUsesMutualAuthentication() bool
	TrustUsesSystemKeychain() bool
	VirtualSubnodes() foundation.INSArray
	SetVirtualSubnodes(value foundation.INSArray)

	// Topic: Instance Methods

	AddTrustTypeTrustAccountTrustPasswordUsernamePasswordJoinExistingError(trustType string, account string, accountPassword string, username string, password string, join bool) (bool, error)
	RemoveTrustUsingUsernamePasswordDeleteTrustAccountError(username string, password string, deleteAccount bool) (bool, error)
	SaveUsingAuthorizationError(authorization objectivec.IObject) (bool, error)
}

// Init initializes the instance.
func (o ODConfiguration) Init() ODConfiguration {
	rv := objc.Send[ODConfiguration](o.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (o ODConfiguration) Autorelease() ODConfiguration {
	rv := objc.Send[ODConfiguration](o.ID, objc.Sel("autorelease"))
	return rv
}

// NewODConfiguration creates a new ODConfiguration instance.
func NewODConfiguration() ODConfiguration {
	class := getODConfigurationClass()
	rv := objc.Send[ODConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/addTrustType(_:trustAccount:trustPassword:username:password:joinExisting:)
func (o ODConfiguration) AddTrustTypeTrustAccountTrustPasswordUsernamePasswordJoinExistingError(trustType string, account string, accountPassword string, username string, password string, join bool) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](o.ID, objc.Sel("addTrustType:trustAccount:trustPassword:username:password:joinExisting:error:"), objc.String(trustType), objc.String(account), objc.String(accountPassword), objc.String(username), objc.String(password), join, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("addTrustType:trustAccount:trustPassword:username:password:joinExisting:error: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/removeTrust(usingUsername:password:deleteTrustAccount:)
func (o ODConfiguration) RemoveTrustUsingUsernamePasswordDeleteTrustAccountError(username string, password string, deleteAccount bool) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](o.ID, objc.Sel("removeTrustUsingUsername:password:deleteTrustAccount:error:"), objc.String(username), objc.String(password), deleteAccount, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("removeTrustUsingUsername:password:deleteTrustAccount:error: returned NO with nil NSError")
	}
	return rv, nil

}

// authorization is a [*securityfoundation.SFAuthorization].
//
// See: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/save(using:)
func (o ODConfiguration) SaveUsingAuthorizationError(authorization objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](o.ID, objc.Sel("saveUsingAuthorization:error:"), authorization, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("saveUsingAuthorization:error: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/suggestedTrustAccount(_:)
func (_ODConfigurationClass ODConfigurationClass) SuggestedTrustAccount(hostname string) string {
	rv := objc.Send[objc.ID](objc.ID(_ODConfigurationClass.class), objc.Sel("suggestedTrustAccount:"), objc.String(hostname))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/suggestedTrustPassword(_:)
func (_ODConfigurationClass ODConfigurationClass) SuggestedTrustPassword(length uintptr) string {
	rv := objc.Send[objc.ID](objc.ID(_ODConfigurationClass.class), objc.Sel("suggestedTrustPassword:"), length)
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/authenticationModuleEntries-swift.property
func (o ODConfiguration) AuthenticationModuleEntries() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("authenticationModuleEntries"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (o ODConfiguration) SetAuthenticationModuleEntries(value foundation.INSArray) {
	objc.Send[struct{}](o.ID, objc.Sel("setAuthenticationModuleEntries:"), value)
}

// See: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/comment-swift.property
func (o ODConfiguration) Comment() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("comment"))
	return foundation.NSStringFromID(rv).String()
}
func (o ODConfiguration) SetComment(value string) {
	objc.Send[struct{}](o.ID, objc.Sel("setComment:"), objc.String(value))
}

// See: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/connectionIdleTimeoutInSeconds-swift.property
func (o ODConfiguration) ConnectionIdleTimeoutInSeconds() int {
	rv := objc.Send[int](o.ID, objc.Sel("connectionIdleTimeoutInSeconds"))
	return rv
}
func (o ODConfiguration) SetConnectionIdleTimeoutInSeconds(value int) {
	objc.Send[struct{}](o.ID, objc.Sel("setConnectionIdleTimeoutInSeconds:"), value)
}

// See: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/connectionSetupTimeoutInSeconds-swift.property
func (o ODConfiguration) ConnectionSetupTimeoutInSeconds() int {
	rv := objc.Send[int](o.ID, objc.Sel("connectionSetupTimeoutInSeconds"))
	return rv
}
func (o ODConfiguration) SetConnectionSetupTimeoutInSeconds(value int) {
	objc.Send[struct{}](o.ID, objc.Sel("setConnectionSetupTimeoutInSeconds:"), value)
}

// See: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/defaultMappings-swift.property
func (o ODConfiguration) DefaultMappings() IODMappings {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("defaultMappings"))
	return ODMappingsFromID(objc.ID(rv))
}
func (o ODConfiguration) SetDefaultMappings(value IODMappings) {
	objc.Send[struct{}](o.ID, objc.Sel("setDefaultMappings:"), value)
}

// See: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/defaultModuleEntries-swift.property
func (o ODConfiguration) DefaultModuleEntries() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("defaultModuleEntries"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (o ODConfiguration) SetDefaultModuleEntries(value foundation.INSArray) {
	objc.Send[struct{}](o.ID, objc.Sel("setDefaultModuleEntries:"), value)
}

// See: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/discoveryModuleEntries-swift.property
func (o ODConfiguration) DiscoveryModuleEntries() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("discoveryModuleEntries"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (o ODConfiguration) SetDiscoveryModuleEntries(value foundation.INSArray) {
	objc.Send[struct{}](o.ID, objc.Sel("setDiscoveryModuleEntries:"), value)
}

// See: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/generalModuleEntries-swift.property
func (o ODConfiguration) GeneralModuleEntries() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("generalModuleEntries"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (o ODConfiguration) SetGeneralModuleEntries(value foundation.INSArray) {
	objc.Send[struct{}](o.ID, objc.Sel("setGeneralModuleEntries:"), value)
}

// See: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/hideRegistration-swift.property
func (o ODConfiguration) HideRegistration() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("hideRegistration"))
	return rv
}
func (o ODConfiguration) SetHideRegistration(value bool) {
	objc.Send[struct{}](o.ID, objc.Sel("setHideRegistration:"), value)
}

// See: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/manInTheMiddleProtection-swift.property
func (o ODConfiguration) ManInTheMiddleProtection() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("manInTheMiddleProtection"))
	return rv
}
func (o ODConfiguration) SetManInTheMiddleProtection(value bool) {
	objc.Send[struct{}](o.ID, objc.Sel("setManInTheMiddleProtection:"), value)
}

// See: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/nodeName-swift.property
func (o ODConfiguration) NodeName() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("nodeName"))
	return foundation.NSStringFromID(rv).String()
}
func (o ODConfiguration) SetNodeName(value string) {
	objc.Send[struct{}](o.ID, objc.Sel("setNodeName:"), objc.String(value))
}

// See: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/packetEncryption-swift.property
func (o ODConfiguration) PacketEncryption() int {
	rv := objc.Send[int](o.ID, objc.Sel("packetEncryption"))
	return rv
}
func (o ODConfiguration) SetPacketEncryption(value int) {
	objc.Send[struct{}](o.ID, objc.Sel("setPacketEncryption:"), value)
}

// See: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/packetSigning-swift.property
func (o ODConfiguration) PacketSigning() int {
	rv := objc.Send[int](o.ID, objc.Sel("packetSigning"))
	return rv
}
func (o ODConfiguration) SetPacketSigning(value int) {
	objc.Send[struct{}](o.ID, objc.Sel("setPacketSigning:"), value)
}

// See: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/preferredDestinationHostName-swift.property
func (o ODConfiguration) PreferredDestinationHostName() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("preferredDestinationHostName"))
	return foundation.NSStringFromID(rv).String()
}
func (o ODConfiguration) SetPreferredDestinationHostName(value string) {
	objc.Send[struct{}](o.ID, objc.Sel("setPreferredDestinationHostName:"), objc.String(value))
}

// See: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/preferredDestinationHostPort-swift.property
func (o ODConfiguration) PreferredDestinationHostPort() uint16 {
	rv := objc.Send[uint16](o.ID, objc.Sel("preferredDestinationHostPort"))
	return rv
}
func (o ODConfiguration) SetPreferredDestinationHostPort(value uint16) {
	objc.Send[struct{}](o.ID, objc.Sel("setPreferredDestinationHostPort:"), value)
}

// See: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/queryTimeoutInSeconds-swift.property
func (o ODConfiguration) QueryTimeoutInSeconds() int {
	rv := objc.Send[int](o.ID, objc.Sel("queryTimeoutInSeconds"))
	return rv
}
func (o ODConfiguration) SetQueryTimeoutInSeconds(value int) {
	objc.Send[struct{}](o.ID, objc.Sel("setQueryTimeoutInSeconds:"), value)
}

// See: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/templateName-swift.property
func (o ODConfiguration) TemplateName() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("templateName"))
	return foundation.NSStringFromID(rv).String()
}
func (o ODConfiguration) SetTemplateName(value string) {
	objc.Send[struct{}](o.ID, objc.Sel("setTemplateName:"), objc.String(value))
}

// See: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/trustAccount-swift.property
func (o ODConfiguration) TrustAccount() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("trustAccount"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/trustKerberosPrincipal-swift.property
func (o ODConfiguration) TrustKerberosPrincipal() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("trustKerberosPrincipal"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/trustMetaAccount-swift.property
func (o ODConfiguration) TrustMetaAccount() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("trustMetaAccount"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/trustType-swift.property
func (o ODConfiguration) TrustType() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("trustType"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/trustUsesKerberosKeytab-swift.property
func (o ODConfiguration) TrustUsesKerberosKeytab() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("trustUsesKerberosKeytab"))
	return rv
}

// See: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/trustUsesMutualAuthentication-swift.property
func (o ODConfiguration) TrustUsesMutualAuthentication() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("trustUsesMutualAuthentication"))
	return rv
}

// See: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/trustUsesSystemKeychain-swift.property
func (o ODConfiguration) TrustUsesSystemKeychain() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("trustUsesSystemKeychain"))
	return rv
}

// See: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/virtualSubnodes-swift.property
func (o ODConfiguration) VirtualSubnodes() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("virtualSubnodes"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (o ODConfiguration) SetVirtualSubnodes(value foundation.INSArray) {
	objc.Send[struct{}](o.ID, objc.Sel("setVirtualSubnodes:"), value)
}
