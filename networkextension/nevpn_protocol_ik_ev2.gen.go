// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
)

// The class instance for the [NEVPNProtocolIKEv2] class.
var (
	_NEVPNProtocolIKEv2Class     NEVPNProtocolIKEv2Class
	_NEVPNProtocolIKEv2ClassOnce sync.Once
)

func getNEVPNProtocolIKEv2Class() NEVPNProtocolIKEv2Class {
	_NEVPNProtocolIKEv2ClassOnce.Do(func() {
		_NEVPNProtocolIKEv2Class = NEVPNProtocolIKEv2Class{class: objc.GetClass("NEVPNProtocolIKEv2")}
	})
	return _NEVPNProtocolIKEv2Class
}

// GetNEVPNProtocolIKEv2Class returns the class object for NEVPNProtocolIKEv2.
func GetNEVPNProtocolIKEv2Class() NEVPNProtocolIKEv2Class {
	return getNEVPNProtocolIKEv2Class()
}

type NEVPNProtocolIKEv2Class struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NEVPNProtocolIKEv2Class) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NEVPNProtocolIKEv2Class) Alloc() NEVPNProtocolIKEv2 {
	rv := objc.Send[NEVPNProtocolIKEv2](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// Settings for an IKEv2 VPN configuration.
//
// # Overview
// 
// Instances of this class are thread safe.
//
// # Accessing IKEv2 Security Association parameters
//
//   - [NEVPNProtocolIKEv2.IKESecurityAssociationParameters]: An [NEVPNIKEv2SecurityAssociationParameters](<doc://com.apple.networkextension/documentation/NetworkExtension/NEVPNIKEv2SecurityAssociationParameters>) object containing the parameters for the initial IKE security association to be negotiated with the IKEv2 server.
//   - [NEVPNProtocolIKEv2.ChildSecurityAssociationParameters]: An [NEVPNIKEv2SecurityAssociationParameters](<doc://com.apple.networkextension/documentation/NetworkExtension/NEVPNIKEv2SecurityAssociationParameters>) object containing the parameters for the child IPSec security associations to be negotiated for each IKEv2 policy.
//
// # Accessing certificate properties
//
//   - [NEVPNProtocolIKEv2.ServerCertificateIssuerCommonName]: A string containing the value of the Subject Common Name field of the Certificate Authority certificate that issued the IKEv2 server’s certificate.
//   - [NEVPNProtocolIKEv2.SetServerCertificateIssuerCommonName]
//   - [NEVPNProtocolIKEv2.ServerCertificateCommonName]: A string containing the value of the Subject Common Name field of the IKEv2 server’s certificate.
//   - [NEVPNProtocolIKEv2.SetServerCertificateCommonName]
//   - [NEVPNProtocolIKEv2.CertificateType]: The type of the certificate in the identity configured in `identityReference` or `identityData`.
//   - [NEVPNProtocolIKEv2.SetCertificateType]
//
// # Accessing TLS version properties
//
//   - [NEVPNProtocolIKEv2.MinimumTLSVersion]: The minimum TLS version to allow for EAP-TLS authentication.
//   - [NEVPNProtocolIKEv2.SetMinimumTLSVersion]
//   - [NEVPNProtocolIKEv2.MaximumTLSVersion]: The minimum TLS version to allow for EAP-TLS authentication.
//   - [NEVPNProtocolIKEv2.SetMaximumTLSVersion]
//
// # Accessing other IKEv2 properties
//
//   - [NEVPNProtocolIKEv2.DeadPeerDetectionRate]: The frequency at which the IKEv2 client will run the dead peer detection algorithm.
//   - [NEVPNProtocolIKEv2.SetDeadPeerDetectionRate]
//   - [NEVPNProtocolIKEv2.UseConfigurationAttributeInternalIPSubnet]: A Boolean indicating whether or not the IKEv2 client should use the INTERNAL_IP4_SUBNET and/or INTERNAL_IP6_SUBNET attributes sent by the IKEv2 server.
//   - [NEVPNProtocolIKEv2.SetUseConfigurationAttributeInternalIPSubnet]
//   - [NEVPNProtocolIKEv2.DisableMOBIKE]: A Boolean indicating whether or not MOBIKE should be disabled for the IKEv2 sessions.
//   - [NEVPNProtocolIKEv2.SetDisableMOBIKE]
//   - [NEVPNProtocolIKEv2.DisableRedirect]: A Boolean indicating whether or not IKEv2 server redirects are disabled.
//   - [NEVPNProtocolIKEv2.SetDisableRedirect]
//   - [NEVPNProtocolIKEv2.EnablePFS]: A Boolean indicating whether or not Perfect Forward Secrecy is enabled.
//   - [NEVPNProtocolIKEv2.SetEnablePFS]
//   - [NEVPNProtocolIKEv2.EnableRevocationCheck]: Enable revocation checking of the IKEv2 server certificate.
//   - [NEVPNProtocolIKEv2.SetEnableRevocationCheck]
//   - [NEVPNProtocolIKEv2.StrictRevocationCheck]: Require a “not revoked” result when checking if the certificate identifying the server is revoked.
//   - [NEVPNProtocolIKEv2.SetStrictRevocationCheck]
//   - [NEVPNProtocolIKEv2.Mtu]: The Maximum Transmission Unit (MTU) size in bytes to assign to the tunnel interface.
//   - [NEVPNProtocolIKEv2.SetMtu]
//
// # Supporting quantum-secure cryptography
//
//   - [NEVPNProtocolIKEv2.AllowPostQuantumKeyExchangeFallback]: A Boolean value that indicates whether servers that don’t support post-quantum key exchanges can skip them.
//   - [NEVPNProtocolIKEv2.SetAllowPostQuantumKeyExchangeFallback]
//   - [NEVPNProtocolIKEv2.PpkConfiguration]: The configuration for a post-quantum pre-shared key (PPK).
//   - [NEVPNProtocolIKEv2.SetPpkConfiguration]
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocolIKEv2
type NEVPNProtocolIKEv2 struct {
	NEVPNProtocolIPSec
}

// NEVPNProtocolIKEv2FromID constructs a [NEVPNProtocolIKEv2] from an objc.ID.
//
// Settings for an IKEv2 VPN configuration.
func NEVPNProtocolIKEv2FromID(id objc.ID) NEVPNProtocolIKEv2 {
	return NEVPNProtocolIKEv2{NEVPNProtocolIPSec: NEVPNProtocolIPSecFromID(id)}
}
// NOTE: NEVPNProtocolIKEv2 adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NEVPNProtocolIKEv2] class.
//
// # Accessing IKEv2 Security Association parameters
//
//   - [INEVPNProtocolIKEv2.IKESecurityAssociationParameters]: An [NEVPNIKEv2SecurityAssociationParameters](<doc://com.apple.networkextension/documentation/NetworkExtension/NEVPNIKEv2SecurityAssociationParameters>) object containing the parameters for the initial IKE security association to be negotiated with the IKEv2 server.
//   - [INEVPNProtocolIKEv2.ChildSecurityAssociationParameters]: An [NEVPNIKEv2SecurityAssociationParameters](<doc://com.apple.networkextension/documentation/NetworkExtension/NEVPNIKEv2SecurityAssociationParameters>) object containing the parameters for the child IPSec security associations to be negotiated for each IKEv2 policy.
//
// # Accessing certificate properties
//
//   - [INEVPNProtocolIKEv2.ServerCertificateIssuerCommonName]: A string containing the value of the Subject Common Name field of the Certificate Authority certificate that issued the IKEv2 server’s certificate.
//   - [INEVPNProtocolIKEv2.SetServerCertificateIssuerCommonName]
//   - [INEVPNProtocolIKEv2.ServerCertificateCommonName]: A string containing the value of the Subject Common Name field of the IKEv2 server’s certificate.
//   - [INEVPNProtocolIKEv2.SetServerCertificateCommonName]
//   - [INEVPNProtocolIKEv2.CertificateType]: The type of the certificate in the identity configured in `identityReference` or `identityData`.
//   - [INEVPNProtocolIKEv2.SetCertificateType]
//
// # Accessing TLS version properties
//
//   - [INEVPNProtocolIKEv2.MinimumTLSVersion]: The minimum TLS version to allow for EAP-TLS authentication.
//   - [INEVPNProtocolIKEv2.SetMinimumTLSVersion]
//   - [INEVPNProtocolIKEv2.MaximumTLSVersion]: The minimum TLS version to allow for EAP-TLS authentication.
//   - [INEVPNProtocolIKEv2.SetMaximumTLSVersion]
//
// # Accessing other IKEv2 properties
//
//   - [INEVPNProtocolIKEv2.DeadPeerDetectionRate]: The frequency at which the IKEv2 client will run the dead peer detection algorithm.
//   - [INEVPNProtocolIKEv2.SetDeadPeerDetectionRate]
//   - [INEVPNProtocolIKEv2.UseConfigurationAttributeInternalIPSubnet]: A Boolean indicating whether or not the IKEv2 client should use the INTERNAL_IP4_SUBNET and/or INTERNAL_IP6_SUBNET attributes sent by the IKEv2 server.
//   - [INEVPNProtocolIKEv2.SetUseConfigurationAttributeInternalIPSubnet]
//   - [INEVPNProtocolIKEv2.DisableMOBIKE]: A Boolean indicating whether or not MOBIKE should be disabled for the IKEv2 sessions.
//   - [INEVPNProtocolIKEv2.SetDisableMOBIKE]
//   - [INEVPNProtocolIKEv2.DisableRedirect]: A Boolean indicating whether or not IKEv2 server redirects are disabled.
//   - [INEVPNProtocolIKEv2.SetDisableRedirect]
//   - [INEVPNProtocolIKEv2.EnablePFS]: A Boolean indicating whether or not Perfect Forward Secrecy is enabled.
//   - [INEVPNProtocolIKEv2.SetEnablePFS]
//   - [INEVPNProtocolIKEv2.EnableRevocationCheck]: Enable revocation checking of the IKEv2 server certificate.
//   - [INEVPNProtocolIKEv2.SetEnableRevocationCheck]
//   - [INEVPNProtocolIKEv2.StrictRevocationCheck]: Require a “not revoked” result when checking if the certificate identifying the server is revoked.
//   - [INEVPNProtocolIKEv2.SetStrictRevocationCheck]
//   - [INEVPNProtocolIKEv2.Mtu]: The Maximum Transmission Unit (MTU) size in bytes to assign to the tunnel interface.
//   - [INEVPNProtocolIKEv2.SetMtu]
//
// # Supporting quantum-secure cryptography
//
//   - [INEVPNProtocolIKEv2.AllowPostQuantumKeyExchangeFallback]: A Boolean value that indicates whether servers that don’t support post-quantum key exchanges can skip them.
//   - [INEVPNProtocolIKEv2.SetAllowPostQuantumKeyExchangeFallback]
//   - [INEVPNProtocolIKEv2.PpkConfiguration]: The configuration for a post-quantum pre-shared key (PPK).
//   - [INEVPNProtocolIKEv2.SetPpkConfiguration]
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocolIKEv2
type INEVPNProtocolIKEv2 interface {
	INEVPNProtocolIPSec

	// Topic: Accessing IKEv2 Security Association parameters

	// An [NEVPNIKEv2SecurityAssociationParameters](<doc://com.apple.networkextension/documentation/NetworkExtension/NEVPNIKEv2SecurityAssociationParameters>) object containing the parameters for the initial IKE security association to be negotiated with the IKEv2 server.
	IKESecurityAssociationParameters() INEVPNIKEv2SecurityAssociationParameters
	// An [NEVPNIKEv2SecurityAssociationParameters](<doc://com.apple.networkextension/documentation/NetworkExtension/NEVPNIKEv2SecurityAssociationParameters>) object containing the parameters for the child IPSec security associations to be negotiated for each IKEv2 policy.
	ChildSecurityAssociationParameters() INEVPNIKEv2SecurityAssociationParameters

	// Topic: Accessing certificate properties

	// A string containing the value of the Subject Common Name field of the Certificate Authority certificate that issued the IKEv2 server’s certificate.
	ServerCertificateIssuerCommonName() string
	SetServerCertificateIssuerCommonName(value string)
	// A string containing the value of the Subject Common Name field of the IKEv2 server’s certificate.
	ServerCertificateCommonName() string
	SetServerCertificateCommonName(value string)
	// The type of the certificate in the identity configured in `identityReference` or `identityData`.
	CertificateType() NEVPNIKEv2CertificateType
	SetCertificateType(value NEVPNIKEv2CertificateType)

	// Topic: Accessing TLS version properties

	// The minimum TLS version to allow for EAP-TLS authentication.
	MinimumTLSVersion() NEVPNIKEv2TLSVersion
	SetMinimumTLSVersion(value NEVPNIKEv2TLSVersion)
	// The minimum TLS version to allow for EAP-TLS authentication.
	MaximumTLSVersion() NEVPNIKEv2TLSVersion
	SetMaximumTLSVersion(value NEVPNIKEv2TLSVersion)

	// Topic: Accessing other IKEv2 properties

	// The frequency at which the IKEv2 client will run the dead peer detection algorithm.
	DeadPeerDetectionRate() NEVPNIKEv2DeadPeerDetectionRate
	SetDeadPeerDetectionRate(value NEVPNIKEv2DeadPeerDetectionRate)
	// A Boolean indicating whether or not the IKEv2 client should use the INTERNAL_IP4_SUBNET and/or INTERNAL_IP6_SUBNET attributes sent by the IKEv2 server.
	UseConfigurationAttributeInternalIPSubnet() bool
	SetUseConfigurationAttributeInternalIPSubnet(value bool)
	// A Boolean indicating whether or not MOBIKE should be disabled for the IKEv2 sessions.
	DisableMOBIKE() bool
	SetDisableMOBIKE(value bool)
	// A Boolean indicating whether or not IKEv2 server redirects are disabled.
	DisableRedirect() bool
	SetDisableRedirect(value bool)
	// A Boolean indicating whether or not Perfect Forward Secrecy is enabled.
	EnablePFS() bool
	SetEnablePFS(value bool)
	// Enable revocation checking of the IKEv2 server certificate.
	EnableRevocationCheck() bool
	SetEnableRevocationCheck(value bool)
	// Require a “not revoked” result when checking if the certificate identifying the server is revoked.
	StrictRevocationCheck() bool
	SetStrictRevocationCheck(value bool)
	// The Maximum Transmission Unit (MTU) size in bytes to assign to the tunnel interface.
	Mtu() uint
	SetMtu(value uint)

	// Topic: Supporting quantum-secure cryptography

	// A Boolean value that indicates whether servers that don’t support post-quantum key exchanges can skip them.
	AllowPostQuantumKeyExchangeFallback() bool
	SetAllowPostQuantumKeyExchangeFallback(value bool)
	// The configuration for a post-quantum pre-shared key (PPK).
	PpkConfiguration() INEVPNIKEv2PPKConfiguration
	SetPpkConfiguration(value INEVPNIKEv2PPKConfiguration)
}

// Init initializes the instance.
func (v NEVPNProtocolIKEv2) Init() NEVPNProtocolIKEv2 {
	rv := objc.Send[NEVPNProtocolIKEv2](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v NEVPNProtocolIKEv2) Autorelease() NEVPNProtocolIKEv2 {
	rv := objc.Send[NEVPNProtocolIKEv2](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEVPNProtocolIKEv2 creates a new NEVPNProtocolIKEv2 instance.
func NewNEVPNProtocolIKEv2() NEVPNProtocolIKEv2 {
	class := getNEVPNProtocolIKEv2Class()
	rv := objc.Send[NEVPNProtocolIKEv2](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// An [NEVPNIKEv2SecurityAssociationParameters] object containing the
// parameters for the initial IKE security association to be negotiated with
// the IKEv2 server.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocolIKEv2/ikeSecurityAssociationParameters
func (v NEVPNProtocolIKEv2) IKESecurityAssociationParameters() INEVPNIKEv2SecurityAssociationParameters {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("IKESecurityAssociationParameters"))
	return NEVPNIKEv2SecurityAssociationParametersFromID(objc.ID(rv))
}
// An [NEVPNIKEv2SecurityAssociationParameters] object containing the
// parameters for the child IPSec security associations to be negotiated for
// each IKEv2 policy.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocolIKEv2/childSecurityAssociationParameters
func (v NEVPNProtocolIKEv2) ChildSecurityAssociationParameters() INEVPNIKEv2SecurityAssociationParameters {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("childSecurityAssociationParameters"))
	return NEVPNIKEv2SecurityAssociationParametersFromID(objc.ID(rv))
}
// A string containing the value of the Subject Common Name field of the
// Certificate Authority certificate that issued the IKEv2 server’s
// certificate.
//
// # Discussion
// 
// This string helps verify the identity of the IKEv2 server.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocolIKEv2/serverCertificateIssuerCommonName
func (v NEVPNProtocolIKEv2) ServerCertificateIssuerCommonName() string {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("serverCertificateIssuerCommonName"))
	return foundation.NSStringFromID(rv).String()
}
func (v NEVPNProtocolIKEv2) SetServerCertificateIssuerCommonName(value string) {
	objc.Send[struct{}](v.ID, objc.Sel("setServerCertificateIssuerCommonName:"), objc.String(value))
}
// A string containing the value of the Subject Common Name field of the IKEv2
// server’s certificate.
//
// # Discussion
// 
// This string is used to help verify the identity of the IKEv2 server.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocolIKEv2/serverCertificateCommonName
func (v NEVPNProtocolIKEv2) ServerCertificateCommonName() string {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("serverCertificateCommonName"))
	return foundation.NSStringFromID(rv).String()
}
func (v NEVPNProtocolIKEv2) SetServerCertificateCommonName(value string) {
	objc.Send[struct{}](v.ID, objc.Sel("setServerCertificateCommonName:"), objc.String(value))
}
// The type of the certificate in the identity configured in
// `identityReference` or `identityData`.
//
// # Discussion
// 
// The default value is [NEVPNIKEv2CertificateType.RSA].
//
// [NEVPNIKEv2CertificateType.RSA]: https://developer.apple.com/documentation/NetworkExtension/NEVPNIKEv2CertificateType/RSA
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocolIKEv2/certificateType
func (v NEVPNProtocolIKEv2) CertificateType() NEVPNIKEv2CertificateType {
	rv := objc.Send[NEVPNIKEv2CertificateType](v.ID, objc.Sel("certificateType"))
	return NEVPNIKEv2CertificateType(rv)
}
func (v NEVPNProtocolIKEv2) SetCertificateType(value NEVPNIKEv2CertificateType) {
	objc.Send[struct{}](v.ID, objc.Sel("setCertificateType:"), value)
}
// The minimum TLS version to allow for EAP-TLS authentication.
//
// # Discussion
// 
// The default value of this property is
// [NEVPNIKEv2TLSVersion.versionDefault].
//
// [NEVPNIKEv2TLSVersion.versionDefault]: https://developer.apple.com/documentation/NetworkExtension/NEVPNIKEv2TLSVersion/versionDefault
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocolIKEv2/minimumTLSVersion
func (v NEVPNProtocolIKEv2) MinimumTLSVersion() NEVPNIKEv2TLSVersion {
	rv := objc.Send[NEVPNIKEv2TLSVersion](v.ID, objc.Sel("minimumTLSVersion"))
	return NEVPNIKEv2TLSVersion(rv)
}
func (v NEVPNProtocolIKEv2) SetMinimumTLSVersion(value NEVPNIKEv2TLSVersion) {
	objc.Send[struct{}](v.ID, objc.Sel("setMinimumTLSVersion:"), value)
}
// The minimum TLS version to allow for EAP-TLS authentication.
//
// # Discussion
// 
// The default value of this property is
// [NEVPNIKEv2TLSVersion.versionDefault].
//
// [NEVPNIKEv2TLSVersion.versionDefault]: https://developer.apple.com/documentation/NetworkExtension/NEVPNIKEv2TLSVersion/versionDefault
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocolIKEv2/maximumTLSVersion
func (v NEVPNProtocolIKEv2) MaximumTLSVersion() NEVPNIKEv2TLSVersion {
	rv := objc.Send[NEVPNIKEv2TLSVersion](v.ID, objc.Sel("maximumTLSVersion"))
	return NEVPNIKEv2TLSVersion(rv)
}
func (v NEVPNProtocolIKEv2) SetMaximumTLSVersion(value NEVPNIKEv2TLSVersion) {
	objc.Send[struct{}](v.ID, objc.Sel("setMaximumTLSVersion:"), value)
}
// The frequency at which the IKEv2 client will run the dead peer detection
// algorithm.
//
// # Discussion
// 
// The IKEv2 client periodically communicates with the IKEv2 server to detect
// when communication with the IKEv2 server has been interrupted. This
// property specifies how frequently this communication takes place. The
// default is [NEVPNIKEv2DeadPeerDetectionRate.medium].
//
// [NEVPNIKEv2DeadPeerDetectionRate.medium]: https://developer.apple.com/documentation/NetworkExtension/NEVPNIKEv2DeadPeerDetectionRate/medium
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocolIKEv2/deadPeerDetectionRate
func (v NEVPNProtocolIKEv2) DeadPeerDetectionRate() NEVPNIKEv2DeadPeerDetectionRate {
	rv := objc.Send[NEVPNIKEv2DeadPeerDetectionRate](v.ID, objc.Sel("deadPeerDetectionRate"))
	return NEVPNIKEv2DeadPeerDetectionRate(rv)
}
func (v NEVPNProtocolIKEv2) SetDeadPeerDetectionRate(value NEVPNIKEv2DeadPeerDetectionRate) {
	objc.Send[struct{}](v.ID, objc.Sel("setDeadPeerDetectionRate:"), value)
}
// A Boolean indicating whether or not the IKEv2 client should use the
// INTERNAL_IP4_SUBNET and/or INTERNAL_IP6_SUBNET attributes sent by the IKEv2
// server.
//
// # Discussion
// 
// If this property is [false], split tunnel configurations may still be
// communicated via Traffic Selectors during IKE negotiation.
// 
// Some IKEv2 servers use the INTERNAL_IP4_SUBNET and INTERNAL_IP6_SUBNET
// protocol message attributes to communicate split tunnel routes to IKEv2
// clients. The default value of this property is [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocolIKEv2/useConfigurationAttributeInternalIPSubnet
func (v NEVPNProtocolIKEv2) UseConfigurationAttributeInternalIPSubnet() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("useConfigurationAttributeInternalIPSubnet"))
	return rv
}
func (v NEVPNProtocolIKEv2) SetUseConfigurationAttributeInternalIPSubnet(value bool) {
	objc.Send[struct{}](v.ID, objc.Sel("setUseConfigurationAttributeInternalIPSubnet:"), value)
}
// A Boolean indicating whether or not MOBIKE should be disabled for the IKEv2
// sessions.
//
// # Discussion
// 
// The default value is [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocolIKEv2/disableMOBIKE
func (v NEVPNProtocolIKEv2) DisableMOBIKE() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("disableMOBIKE"))
	return rv
}
func (v NEVPNProtocolIKEv2) SetDisableMOBIKE(value bool) {
	objc.Send[struct{}](v.ID, objc.Sel("setDisableMOBIKE:"), value)
}
// A Boolean indicating whether or not IKEv2 server redirects are disabled.
//
// # Discussion
// 
// The default value is [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocolIKEv2/disableRedirect
func (v NEVPNProtocolIKEv2) DisableRedirect() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("disableRedirect"))
	return rv
}
func (v NEVPNProtocolIKEv2) SetDisableRedirect(value bool) {
	objc.Send[struct{}](v.ID, objc.Sel("setDisableRedirect:"), value)
}
// A Boolean indicating whether or not Perfect Forward Secrecy is enabled.
//
// # Discussion
// 
// The default value is [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocolIKEv2/enablePFS
func (v NEVPNProtocolIKEv2) EnablePFS() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("enablePFS"))
	return rv
}
func (v NEVPNProtocolIKEv2) SetEnablePFS(value bool) {
	objc.Send[struct{}](v.ID, objc.Sel("setEnablePFS:"), value)
}
// Enable revocation checking of the IKEv2 server certificate.
//
// # Discussion
// 
// The default value is NO. If this property is set to YES, then during IKEv2
// negotiation the certificate identifying the server is checked to see if it
// has been revoked.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocolIKEv2/enableRevocationCheck
func (v NEVPNProtocolIKEv2) EnableRevocationCheck() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("enableRevocationCheck"))
	return rv
}
func (v NEVPNProtocolIKEv2) SetEnableRevocationCheck(value bool) {
	objc.Send[struct{}](v.ID, objc.Sel("setEnableRevocationCheck:"), value)
}
// Require a “not revoked” result when checking if the certificate
// identifying the server is revoked.
//
// # Discussion
// 
// The default value is NO. If this property is set to NO, then either a
// “not revoked” result from the certificate revocation server or a
// failure to communicate with the certificate revocation server will result
// in a successful revocation check. If this property is set to YES, then only
// a “not revoked” result from the certificate revocation server will
// result in a successful revocation check.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocolIKEv2/strictRevocationCheck
func (v NEVPNProtocolIKEv2) StrictRevocationCheck() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("strictRevocationCheck"))
	return rv
}
func (v NEVPNProtocolIKEv2) SetStrictRevocationCheck(value bool) {
	objc.Send[struct{}](v.ID, objc.Sel("setStrictRevocationCheck:"), value)
}
// The Maximum Transmission Unit (MTU) size in bytes to assign to the tunnel
// interface.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocolIKEv2/mtu
func (v NEVPNProtocolIKEv2) Mtu() uint {
	rv := objc.Send[uint](v.ID, objc.Sel("mtu"))
	return rv
}
func (v NEVPNProtocolIKEv2) SetMtu(value uint) {
	objc.Send[struct{}](v.ID, objc.Sel("setMtu:"), value)
}
// A Boolean value that indicates whether servers that don’t support
// post-quantum key exchanges can skip them.
//
// # Discussion
// 
// This property has no effect if you don’t configure any post-quantum key
// exchange methods in the [NEVPNIKEv2SecurityAssociationParameters]. The
// property’s default value is `false`.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocolIKEv2/allowPostQuantumKeyExchangeFallback
func (v NEVPNProtocolIKEv2) AllowPostQuantumKeyExchangeFallback() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("allowPostQuantumKeyExchangeFallback"))
	return rv
}
func (v NEVPNProtocolIKEv2) SetAllowPostQuantumKeyExchangeFallback(value bool) {
	objc.Send[struct{}](v.ID, objc.Sel("setAllowPostQuantumKeyExchangeFallback:"), value)
}
// The configuration for a post-quantum pre-shared key (PPK).
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocolIKEv2/ppkConfiguration
func (v NEVPNProtocolIKEv2) PpkConfiguration() INEVPNIKEv2PPKConfiguration {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("ppkConfiguration"))
	return NEVPNIKEv2PPKConfigurationFromID(objc.ID(rv))
}
func (v NEVPNProtocolIKEv2) SetPpkConfiguration(value INEVPNIKEv2PPKConfiguration) {
	objc.Send[struct{}](v.ID, objc.Sel("setPpkConfiguration:"), value)
}

