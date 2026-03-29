// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NEVPNIKEv2SecurityAssociationParameters] class.
var (
	_NEVPNIKEv2SecurityAssociationParametersClass     NEVPNIKEv2SecurityAssociationParametersClass
	_NEVPNIKEv2SecurityAssociationParametersClassOnce sync.Once
)

func getNEVPNIKEv2SecurityAssociationParametersClass() NEVPNIKEv2SecurityAssociationParametersClass {
	_NEVPNIKEv2SecurityAssociationParametersClassOnce.Do(func() {
		_NEVPNIKEv2SecurityAssociationParametersClass = NEVPNIKEv2SecurityAssociationParametersClass{class: objc.GetClass("NEVPNIKEv2SecurityAssociationParameters")}
	})
	return _NEVPNIKEv2SecurityAssociationParametersClass
}

// GetNEVPNIKEv2SecurityAssociationParametersClass returns the class object for NEVPNIKEv2SecurityAssociationParameters.
func GetNEVPNIKEv2SecurityAssociationParametersClass() NEVPNIKEv2SecurityAssociationParametersClass {
	return getNEVPNIKEv2SecurityAssociationParametersClass()
}

type NEVPNIKEv2SecurityAssociationParametersClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NEVPNIKEv2SecurityAssociationParametersClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NEVPNIKEv2SecurityAssociationParametersClass) Alloc() NEVPNIKEv2SecurityAssociationParameters {
	rv := objc.Send[NEVPNIKEv2SecurityAssociationParameters](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// Parameters for an IKEv2 Security Association.
//
// # IKEv2 Security Association parameters
//
//   - [NEVPNIKEv2SecurityAssociationParameters.EncryptionAlgorithm]: The algorithm used by the Security Association to encrypt and decrypt data.
//   - [NEVPNIKEv2SecurityAssociationParameters.SetEncryptionAlgorithm]
//   - [NEVPNIKEv2SecurityAssociationParameters.IntegrityAlgorithm]: The algorithm used by the Security Association to verify the integrity of data.
//   - [NEVPNIKEv2SecurityAssociationParameters.SetIntegrityAlgorithm]
//   - [NEVPNIKEv2SecurityAssociationParameters.DiffieHellmanGroup]: The Diffie Hellman group used by the Security Association.
//   - [NEVPNIKEv2SecurityAssociationParameters.SetDiffieHellmanGroup]
//   - [NEVPNIKEv2SecurityAssociationParameters.LifetimeMinutes]: The duration of the lifetime of the Security Association, in minutes.
//   - [NEVPNIKEv2SecurityAssociationParameters.SetLifetimeMinutes]
//   - [NEVPNIKEv2SecurityAssociationParameters.PostQuantumKeyExchangeMethods]: A list of the quantum-secure key exchange methods the Security Association uses.
//   - [NEVPNIKEv2SecurityAssociationParameters.SetPostQuantumKeyExchangeMethods]
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEVPNIKEv2SecurityAssociationParameters
type NEVPNIKEv2SecurityAssociationParameters struct {
	objectivec.Object
}

// NEVPNIKEv2SecurityAssociationParametersFromID constructs a [NEVPNIKEv2SecurityAssociationParameters] from an objc.ID.
//
// Parameters for an IKEv2 Security Association.
func NEVPNIKEv2SecurityAssociationParametersFromID(id objc.ID) NEVPNIKEv2SecurityAssociationParameters {
	return NEVPNIKEv2SecurityAssociationParameters{objectivec.Object{ID: id}}
}
// NOTE: NEVPNIKEv2SecurityAssociationParameters adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NEVPNIKEv2SecurityAssociationParameters] class.
//
// # IKEv2 Security Association parameters
//
//   - [INEVPNIKEv2SecurityAssociationParameters.EncryptionAlgorithm]: The algorithm used by the Security Association to encrypt and decrypt data.
//   - [INEVPNIKEv2SecurityAssociationParameters.SetEncryptionAlgorithm]
//   - [INEVPNIKEv2SecurityAssociationParameters.IntegrityAlgorithm]: The algorithm used by the Security Association to verify the integrity of data.
//   - [INEVPNIKEv2SecurityAssociationParameters.SetIntegrityAlgorithm]
//   - [INEVPNIKEv2SecurityAssociationParameters.DiffieHellmanGroup]: The Diffie Hellman group used by the Security Association.
//   - [INEVPNIKEv2SecurityAssociationParameters.SetDiffieHellmanGroup]
//   - [INEVPNIKEv2SecurityAssociationParameters.LifetimeMinutes]: The duration of the lifetime of the Security Association, in minutes.
//   - [INEVPNIKEv2SecurityAssociationParameters.SetLifetimeMinutes]
//   - [INEVPNIKEv2SecurityAssociationParameters.PostQuantumKeyExchangeMethods]: A list of the quantum-secure key exchange methods the Security Association uses.
//   - [INEVPNIKEv2SecurityAssociationParameters.SetPostQuantumKeyExchangeMethods]
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEVPNIKEv2SecurityAssociationParameters
type INEVPNIKEv2SecurityAssociationParameters interface {
	objectivec.IObject

	// Topic: IKEv2 Security Association parameters

	// The algorithm used by the Security Association to encrypt and decrypt data.
	EncryptionAlgorithm() NEVPNIKEv2EncryptionAlgorithm
	SetEncryptionAlgorithm(value NEVPNIKEv2EncryptionAlgorithm)
	// The algorithm used by the Security Association to verify the integrity of data.
	IntegrityAlgorithm() NEVPNIKEv2IntegrityAlgorithm
	SetIntegrityAlgorithm(value NEVPNIKEv2IntegrityAlgorithm)
	// The Diffie Hellman group used by the Security Association.
	DiffieHellmanGroup() NEVPNIKEv2DiffieHellmanGroup
	SetDiffieHellmanGroup(value NEVPNIKEv2DiffieHellmanGroup)
	// The duration of the lifetime of the Security Association, in minutes.
	LifetimeMinutes() int32
	SetLifetimeMinutes(value int32)
	// A list of the quantum-secure key exchange methods the Security Association uses.
	PostQuantumKeyExchangeMethods() NEVPNIKEv2PostQuantumKeyExchangeMethod
	SetPostQuantumKeyExchangeMethods(value NEVPNIKEv2PostQuantumKeyExchangeMethod)

	// An 
	ChildSecurityAssociationParameters() INEVPNIKEv2SecurityAssociationParameters
	SetChildSecurityAssociationParameters(value INEVPNIKEv2SecurityAssociationParameters)
	// An 
	IkeSecurityAssociationParameters() INEVPNIKEv2SecurityAssociationParameters
	SetIkeSecurityAssociationParameters(value INEVPNIKEv2SecurityAssociationParameters)
	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (v NEVPNIKEv2SecurityAssociationParameters) Init() NEVPNIKEv2SecurityAssociationParameters {
	rv := objc.Send[NEVPNIKEv2SecurityAssociationParameters](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v NEVPNIKEv2SecurityAssociationParameters) Autorelease() NEVPNIKEv2SecurityAssociationParameters {
	rv := objc.Send[NEVPNIKEv2SecurityAssociationParameters](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEVPNIKEv2SecurityAssociationParameters creates a new NEVPNIKEv2SecurityAssociationParameters instance.
func NewNEVPNIKEv2SecurityAssociationParameters() NEVPNIKEv2SecurityAssociationParameters {
	class := getNEVPNIKEv2SecurityAssociationParametersClass()
	rv := objc.Send[NEVPNIKEv2SecurityAssociationParameters](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (v NEVPNIKEv2SecurityAssociationParameters) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](v.ID, objc.Sel("encodeWithCoder:"), coder)
}

// The algorithm used by the Security Association to encrypt and decrypt data.
//
// # Discussion
// 
// The default value of this property is
// [NEVPNIKEv2EncryptionAlgorithm.algorithmAES256], except on tvOS where the
// default is [NEVPNIKEv2EncryptionAlgorithm.algorithmAES256GCM].
//
// [NEVPNIKEv2EncryptionAlgorithm.algorithmAES256GCM]: https://developer.apple.com/documentation/NetworkExtension/NEVPNIKEv2EncryptionAlgorithm/algorithmAES256GCM
// [NEVPNIKEv2EncryptionAlgorithm.algorithmAES256]: https://developer.apple.com/documentation/NetworkExtension/NEVPNIKEv2EncryptionAlgorithm/algorithmAES256
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEVPNIKEv2SecurityAssociationParameters/encryptionAlgorithm
func (v NEVPNIKEv2SecurityAssociationParameters) EncryptionAlgorithm() NEVPNIKEv2EncryptionAlgorithm {
	rv := objc.Send[NEVPNIKEv2EncryptionAlgorithm](v.ID, objc.Sel("encryptionAlgorithm"))
	return NEVPNIKEv2EncryptionAlgorithm(rv)
}
func (v NEVPNIKEv2SecurityAssociationParameters) SetEncryptionAlgorithm(value NEVPNIKEv2EncryptionAlgorithm) {
	objc.Send[struct{}](v.ID, objc.Sel("setEncryptionAlgorithm:"), value)
}
// The algorithm used by the Security Association to verify the integrity of
// data.
//
// # Discussion
// 
// The default value of this property is
// [NEVPNIKEv2IntegrityAlgorithm.SHA256].
// 
// The system infers its IKE psedo-random number generation algorithm based on
// the integrity algorithm.
//
// [NEVPNIKEv2IntegrityAlgorithm.SHA256]: https://developer.apple.com/documentation/NetworkExtension/NEVPNIKEv2IntegrityAlgorithm/SHA256
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEVPNIKEv2SecurityAssociationParameters/integrityAlgorithm
func (v NEVPNIKEv2SecurityAssociationParameters) IntegrityAlgorithm() NEVPNIKEv2IntegrityAlgorithm {
	rv := objc.Send[NEVPNIKEv2IntegrityAlgorithm](v.ID, objc.Sel("integrityAlgorithm"))
	return NEVPNIKEv2IntegrityAlgorithm(rv)
}
func (v NEVPNIKEv2SecurityAssociationParameters) SetIntegrityAlgorithm(value NEVPNIKEv2IntegrityAlgorithm) {
	objc.Send[struct{}](v.ID, objc.Sel("setIntegrityAlgorithm:"), value)
}
// The Diffie Hellman group used by the Security Association.
//
// # Discussion
// 
// The default value of this property is
// [NEVPNIKEv2DiffieHellmanGroup.group14].
// 
// The value of this property on [ChildSecurityAssociationParameters] of
// [NEVPNProtocolIKEv2] only takes effect if the [EnablePFS] of
// [NEVPNProtocolIKEv2] is [true] (its default value is [false]).
//
// [NEVPNIKEv2DiffieHellmanGroup.group14]: https://developer.apple.com/documentation/NetworkExtension/NEVPNIKEv2DiffieHellmanGroup/group14
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEVPNIKEv2SecurityAssociationParameters/diffieHellmanGroup
func (v NEVPNIKEv2SecurityAssociationParameters) DiffieHellmanGroup() NEVPNIKEv2DiffieHellmanGroup {
	rv := objc.Send[NEVPNIKEv2DiffieHellmanGroup](v.ID, objc.Sel("diffieHellmanGroup"))
	return NEVPNIKEv2DiffieHellmanGroup(rv)
}
func (v NEVPNIKEv2SecurityAssociationParameters) SetDiffieHellmanGroup(value NEVPNIKEv2DiffieHellmanGroup) {
	objc.Send[struct{}](v.ID, objc.Sel("setDiffieHellmanGroup:"), value)
}
// The duration of the lifetime of the Security Association, in minutes.
//
// # Discussion
// 
// The default is 60 for IKE Security Associations, and 30 for Child Security
// Associations. Before the end of the lifetime is reached, IKEv2 will attempt
// to negotiate new keys for the Security Association in order to maintain the
// IKEv2 session.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEVPNIKEv2SecurityAssociationParameters/lifetimeMinutes
func (v NEVPNIKEv2SecurityAssociationParameters) LifetimeMinutes() int32 {
	rv := objc.Send[int32](v.ID, objc.Sel("lifetimeMinutes"))
	return rv
}
func (v NEVPNIKEv2SecurityAssociationParameters) SetLifetimeMinutes(value int32) {
	objc.Send[struct{}](v.ID, objc.Sel("setLifetimeMinutes:"), value)
}
// A list of the quantum-secure key exchange methods the Security Association
// uses.
//
// See: https://developer.apple.com/documentation/networkextension/nevpnikev2securityassociationparameters/postquantumkeyexchangemethods-3173s
func (v NEVPNIKEv2SecurityAssociationParameters) PostQuantumKeyExchangeMethods() NEVPNIKEv2PostQuantumKeyExchangeMethod {
	rv := objc.Send[NEVPNIKEv2PostQuantumKeyExchangeMethod](v.ID, objc.Sel("postQuantumKeyExchangeMethods"))
	return NEVPNIKEv2PostQuantumKeyExchangeMethod(rv)
}
func (v NEVPNIKEv2SecurityAssociationParameters) SetPostQuantumKeyExchangeMethods(value NEVPNIKEv2PostQuantumKeyExchangeMethod) {
	objc.Send[struct{}](v.ID, objc.Sel("setPostQuantumKeyExchangeMethods:"), value)
}
// An
//
// See: https://developer.apple.com/documentation/networkextension/nevpnprotocolikev2/childsecurityassociationparameters
func (v NEVPNIKEv2SecurityAssociationParameters) ChildSecurityAssociationParameters() INEVPNIKEv2SecurityAssociationParameters {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("childSecurityAssociationParameters"))
	return NEVPNIKEv2SecurityAssociationParametersFromID(objc.ID(rv))
}
func (v NEVPNIKEv2SecurityAssociationParameters) SetChildSecurityAssociationParameters(value INEVPNIKEv2SecurityAssociationParameters) {
	objc.Send[struct{}](v.ID, objc.Sel("setChildSecurityAssociationParameters:"), value)
}
// An
//
// See: https://developer.apple.com/documentation/networkextension/nevpnprotocolikev2/ikesecurityassociationparameters
func (v NEVPNIKEv2SecurityAssociationParameters) IkeSecurityAssociationParameters() INEVPNIKEv2SecurityAssociationParameters {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("IKESecurityAssociationParameters"))
	return NEVPNIKEv2SecurityAssociationParametersFromID(objc.ID(rv))
}
func (v NEVPNIKEv2SecurityAssociationParameters) SetIkeSecurityAssociationParameters(value INEVPNIKEv2SecurityAssociationParameters) {
	objc.Send[struct{}](v.ID, objc.Sel("setIKESecurityAssociationParameters:"), value)
}

