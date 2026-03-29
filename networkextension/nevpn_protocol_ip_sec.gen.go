// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
)

// The class instance for the [NEVPNProtocolIPSec] class.
var (
	_NEVPNProtocolIPSecClass     NEVPNProtocolIPSecClass
	_NEVPNProtocolIPSecClassOnce sync.Once
)

func getNEVPNProtocolIPSecClass() NEVPNProtocolIPSecClass {
	_NEVPNProtocolIPSecClassOnce.Do(func() {
		_NEVPNProtocolIPSecClass = NEVPNProtocolIPSecClass{class: objc.GetClass("NEVPNProtocolIPSec")}
	})
	return _NEVPNProtocolIPSecClass
}

// GetNEVPNProtocolIPSecClass returns the class object for NEVPNProtocolIPSec.
func GetNEVPNProtocolIPSecClass() NEVPNProtocolIPSecClass {
	return getNEVPNProtocolIPSecClass()
}

type NEVPNProtocolIPSecClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NEVPNProtocolIPSecClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NEVPNProtocolIPSecClass) Alloc() NEVPNProtocolIPSec {
	rv := objc.Send[NEVPNProtocolIPSec](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// Settings for an IPsec VPN configuration.
//
// # Overview
// 
// To configure IKE version 2 (IKEv2), use the [NEVPNProtocolIKEv2] subclass.
// Instantiating [NEVPNProtocolIPSec] directly implies IKE version 1.
//
// # Accessing IPSec properties
//
//   - [NEVPNProtocolIPSec.AuthenticationMethod]: The method used to authenticate the device with the IPSec server. For IKE version 2, when using extended authentication, this authentication method only affects how the client validates the authentication payload presented by the server.
//   - [NEVPNProtocolIPSec.SetAuthenticationMethod]
//   - [NEVPNProtocolIPSec.UseExtendedAuthentication]: A flag indicating if extended authentication will be negotiated.
//   - [NEVPNProtocolIPSec.SetUseExtendedAuthentication]
//   - [NEVPNProtocolIPSec.SharedSecretReference]: A persistent keychain reference to a keychain item containing the IKE shared secret.
//   - [NEVPNProtocolIPSec.SetSharedSecretReference]
//   - [NEVPNProtocolIPSec.LocalIdentifier]: A string identifying the iOS or macOS device for authentication purposes
//   - [NEVPNProtocolIPSec.SetLocalIdentifier]
//   - [NEVPNProtocolIPSec.RemoteIdentifier]: A string identifying the IPSec server for authentication purposes
//   - [NEVPNProtocolIPSec.SetRemoteIdentifier]
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocolIPSec
type NEVPNProtocolIPSec struct {
	NEVPNProtocol
}

// NEVPNProtocolIPSecFromID constructs a [NEVPNProtocolIPSec] from an objc.ID.
//
// Settings for an IPsec VPN configuration.
func NEVPNProtocolIPSecFromID(id objc.ID) NEVPNProtocolIPSec {
	return NEVPNProtocolIPSec{NEVPNProtocol: NEVPNProtocolFromID(id)}
}
// NOTE: NEVPNProtocolIPSec adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NEVPNProtocolIPSec] class.
//
// # Accessing IPSec properties
//
//   - [INEVPNProtocolIPSec.AuthenticationMethod]: The method used to authenticate the device with the IPSec server. For IKE version 2, when using extended authentication, this authentication method only affects how the client validates the authentication payload presented by the server.
//   - [INEVPNProtocolIPSec.SetAuthenticationMethod]
//   - [INEVPNProtocolIPSec.UseExtendedAuthentication]: A flag indicating if extended authentication will be negotiated.
//   - [INEVPNProtocolIPSec.SetUseExtendedAuthentication]
//   - [INEVPNProtocolIPSec.SharedSecretReference]: A persistent keychain reference to a keychain item containing the IKE shared secret.
//   - [INEVPNProtocolIPSec.SetSharedSecretReference]
//   - [INEVPNProtocolIPSec.LocalIdentifier]: A string identifying the iOS or macOS device for authentication purposes
//   - [INEVPNProtocolIPSec.SetLocalIdentifier]
//   - [INEVPNProtocolIPSec.RemoteIdentifier]: A string identifying the IPSec server for authentication purposes
//   - [INEVPNProtocolIPSec.SetRemoteIdentifier]
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocolIPSec
type INEVPNProtocolIPSec interface {
	INEVPNProtocol

	// Topic: Accessing IPSec properties

	// The method used to authenticate the device with the IPSec server. For IKE version 2, when using extended authentication, this authentication method only affects how the client validates the authentication payload presented by the server.
	AuthenticationMethod() NEVPNIKEAuthenticationMethod
	SetAuthenticationMethod(value NEVPNIKEAuthenticationMethod)
	// A flag indicating if extended authentication will be negotiated.
	UseExtendedAuthentication() bool
	SetUseExtendedAuthentication(value bool)
	// A persistent keychain reference to a keychain item containing the IKE shared secret.
	SharedSecretReference() foundation.INSData
	SetSharedSecretReference(value foundation.INSData)
	// A string identifying the iOS or macOS device for authentication purposes
	LocalIdentifier() string
	SetLocalIdentifier(value string)
	// A string identifying the IPSec server for authentication purposes
	RemoteIdentifier() string
	SetRemoteIdentifier(value string)
}

// Init initializes the instance.
func (v NEVPNProtocolIPSec) Init() NEVPNProtocolIPSec {
	rv := objc.Send[NEVPNProtocolIPSec](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v NEVPNProtocolIPSec) Autorelease() NEVPNProtocolIPSec {
	rv := objc.Send[NEVPNProtocolIPSec](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEVPNProtocolIPSec creates a new NEVPNProtocolIPSec instance.
func NewNEVPNProtocolIPSec() NEVPNProtocolIPSec {
	class := getNEVPNProtocolIPSecClass()
	rv := objc.Send[NEVPNProtocolIPSec](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The method used to authenticate the device with the IPSec server. For IKE
// version 2, when using extended authentication, this authentication method
// only affects how the client validates the authentication payload presented
// by the server.
//
// # Discussion
// 
// If this property is set to [NEVPNIKEAuthenticationMethodNone], extended
// authentication will still be negotiated if [UseExtendedAuthentication] is
// set to [true].
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocolIPSec/authenticationMethod
func (v NEVPNProtocolIPSec) AuthenticationMethod() NEVPNIKEAuthenticationMethod {
	rv := objc.Send[NEVPNIKEAuthenticationMethod](v.ID, objc.Sel("authenticationMethod"))
	return NEVPNIKEAuthenticationMethod(rv)
}
func (v NEVPNProtocolIPSec) SetAuthenticationMethod(value NEVPNIKEAuthenticationMethod) {
	objc.Send[struct{}](v.ID, objc.Sel("setAuthenticationMethod:"), value)
}
// A flag indicating if extended authentication will be negotiated.
//
// # Discussion
// 
// This authentication is in addition to the IKE authentication used to
// authenticate the endpoints of the IKE session.
// 
// - For IKE version 1, when this flag is set X-Auth authentication will be
// negotiated as part of the IKE session, using the `username` and
// `passwordReference` properties as the credential. - For IKE version 2, when
// this flag is set EAP authentication will be negotiated as part of the IKE
// session, using the `username`, `passwordReference`, and/or
// `identityReference` properties as the credential depending on which EAP
// method the server requires.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocolIPSec/useExtendedAuthentication
func (v NEVPNProtocolIPSec) UseExtendedAuthentication() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("useExtendedAuthentication"))
	return rv
}
func (v NEVPNProtocolIPSec) SetUseExtendedAuthentication(value bool) {
	objc.Send[struct{}](v.ID, objc.Sel("setUseExtendedAuthentication:"), value)
}
// A persistent keychain reference to a keychain item containing the IKE
// shared secret.
//
// # Discussion
// 
// The persistent keychain reference must refer to a kerychain item of class
// [kSecClassGenericPassword]
//
// [kSecClassGenericPassword]: https://developer.apple.com/documentation/Security/kSecClassGenericPassword
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocolIPSec/sharedSecretReference
func (v NEVPNProtocolIPSec) SharedSecretReference() foundation.INSData {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("sharedSecretReference"))
	return foundation.NSDataFromID(objc.ID(rv))
}
func (v NEVPNProtocolIPSec) SetSharedSecretReference(value foundation.INSData) {
	objc.Send[struct{}](v.ID, objc.Sel("setSharedSecretReference:"), value)
}
// A string identifying the iOS or macOS device for authentication purposes
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocolIPSec/localIdentifier
func (v NEVPNProtocolIPSec) LocalIdentifier() string {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("localIdentifier"))
	return foundation.NSStringFromID(rv).String()
}
func (v NEVPNProtocolIPSec) SetLocalIdentifier(value string) {
	objc.Send[struct{}](v.ID, objc.Sel("setLocalIdentifier:"), objc.String(value))
}
// A string identifying the IPSec server for authentication purposes
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocolIPSec/remoteIdentifier
func (v NEVPNProtocolIPSec) RemoteIdentifier() string {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("remoteIdentifier"))
	return foundation.NSStringFromID(rv).String()
}
func (v NEVPNProtocolIPSec) SetRemoteIdentifier(value string) {
	objc.Send[struct{}](v.ID, objc.Sel("setRemoteIdentifier:"), objc.String(value))
}

