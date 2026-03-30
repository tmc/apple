// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NWTLSParameters] class.
var (
	_NWTLSParametersClass     NWTLSParametersClass
	_NWTLSParametersClassOnce sync.Once
)

func getNWTLSParametersClass() NWTLSParametersClass {
	_NWTLSParametersClassOnce.Do(func() {
		_NWTLSParametersClass = NWTLSParametersClass{class: objc.GetClass("NWTLSParameters")}
	})
	return _NWTLSParametersClass
}

// GetNWTLSParametersClass returns the class object for NWTLSParameters.
func GetNWTLSParametersClass() NWTLSParametersClass {
	return getNWTLSParametersClass()
}

type NWTLSParametersClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NWTLSParametersClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NWTLSParametersClass) Alloc() NWTLSParameters {
	rv := objc.Send[NWTLSParameters](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// TLS properties for creating a connection.
//
// # Accessing TLS parameters
//
//   - [NWTLSParameters.TLSSessionID]: The Session ID to use for the associated TCP connection.
//   - [NWTLSParameters.SetTLSSessionID]
//   - [NWTLSParameters.SSLCipherSuites]: The set of allowed cipher suites when negotiating TLS.
//   - [NWTLSParameters.SetSSLCipherSuites]
//   - [NWTLSParameters.MinimumSSLProtocolVersion]: The minimum allowed [SSLProtocol] value to use when negotiating TLS.
//   - [NWTLSParameters.SetMinimumSSLProtocolVersion]
//   - [NWTLSParameters.MaximumSSLProtocolVersion]: The maximum allowed [SSLProtocol] value to use when negotiating TLS.
//   - [NWTLSParameters.SetMaximumSSLProtocolVersion]
//
// See: https://developer.apple.com/documentation/NetworkExtension/NWTLSParameters
type NWTLSParameters struct {
	objectivec.Object
}

// NWTLSParametersFromID constructs a [NWTLSParameters] from an objc.ID.
//
// TLS properties for creating a connection.
func NWTLSParametersFromID(id objc.ID) NWTLSParameters {
	return NWTLSParameters{objectivec.Object{ID: id}}
}

// NOTE: NWTLSParameters adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NWTLSParameters] class.
//
// # Accessing TLS parameters
//
//   - [INWTLSParameters.TLSSessionID]: The Session ID to use for the associated TCP connection.
//   - [INWTLSParameters.SetTLSSessionID]
//   - [INWTLSParameters.SSLCipherSuites]: The set of allowed cipher suites when negotiating TLS.
//   - [INWTLSParameters.SetSSLCipherSuites]
//   - [INWTLSParameters.MinimumSSLProtocolVersion]: The minimum allowed [SSLProtocol] value to use when negotiating TLS.
//   - [INWTLSParameters.SetMinimumSSLProtocolVersion]
//   - [INWTLSParameters.MaximumSSLProtocolVersion]: The maximum allowed [SSLProtocol] value to use when negotiating TLS.
//   - [INWTLSParameters.SetMaximumSSLProtocolVersion]
//
// See: https://developer.apple.com/documentation/NetworkExtension/NWTLSParameters
type INWTLSParameters interface {
	objectivec.IObject

	// Topic: Accessing TLS parameters

	// The Session ID to use for the associated TCP connection.
	TLSSessionID() foundation.INSData
	SetTLSSessionID(value foundation.INSData)
	// The set of allowed cipher suites when negotiating TLS.
	SSLCipherSuites() foundation.INSSet
	SetSSLCipherSuites(value foundation.INSSet)
	// The minimum allowed [SSLProtocol] value to use when negotiating TLS.
	MinimumSSLProtocolVersion() uint
	SetMinimumSSLProtocolVersion(value uint)
	// The maximum allowed [SSLProtocol] value to use when negotiating TLS.
	MaximumSSLProtocolVersion() uint
	SetMaximumSSLProtocolVersion(value uint)
}

// Init initializes the instance.
func (n NWTLSParameters) Init() NWTLSParameters {
	rv := objc.Send[NWTLSParameters](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n NWTLSParameters) Autorelease() NWTLSParameters {
	rv := objc.Send[NWTLSParameters](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewNWTLSParameters creates a new NWTLSParameters instance.
func NewNWTLSParameters() NWTLSParameters {
	class := getNWTLSParametersClass()
	rv := objc.Send[NWTLSParameters](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The Session ID to use for the associated TCP connection.
//
// # Discussion
//
// The Session ID is used for TLS session resumption.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NWTLSParameters/tlsSessionID
func (n NWTLSParameters) TLSSessionID() foundation.INSData {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("TLSSessionID"))
	return foundation.NSDataFromID(objc.ID(rv))
}
func (n NWTLSParameters) SetTLSSessionID(value foundation.INSData) {
	objc.Send[struct{}](n.ID, objc.Sel("setTLSSessionID:"), value)
}

// The set of allowed cipher suites when negotiating TLS.
//
// # Discussion
//
// Values for cipher suites are defined in “. These values should be wrapped
// as [NSNumber] objects in a set. If this property is set to `nil`, the
// default cipher suites will be used.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NWTLSParameters/sslCipherSuites
//
// [NSNumber]: https://developer.apple.com/documentation/Foundation/NSNumber
func (n NWTLSParameters) SSLCipherSuites() foundation.INSSet {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("SSLCipherSuites"))
	return foundation.NSSetFromID(objc.ID(rv))
}
func (n NWTLSParameters) SetSSLCipherSuites(value foundation.INSSet) {
	objc.Send[struct{}](n.ID, objc.Sel("setSSLCipherSuites:"), value)
}

// The minimum allowed [SSLProtocol] value to use when negotiating TLS.
//
// # Discussion
//
// Values for [SSLProtocol] are defined in “. If set to a non-zero value, the
// SSL handshake will not accept any protocol version less than the minimum.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NWTLSParameters/minimumSSLProtocolVersion
func (n NWTLSParameters) MinimumSSLProtocolVersion() uint {
	rv := objc.Send[uint](n.ID, objc.Sel("minimumSSLProtocolVersion"))
	return rv
}
func (n NWTLSParameters) SetMinimumSSLProtocolVersion(value uint) {
	objc.Send[struct{}](n.ID, objc.Sel("setMinimumSSLProtocolVersion:"), value)
}

// The maximum allowed [SSLProtocol] value to use when negotiating TLS.
//
// # Discussion
//
// Values for [SSLProtocol] are defined in “. If set to a non-zero value, the
// SSL handshake will not accept any protocol version greater than the
// maximum.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NWTLSParameters/maximumSSLProtocolVersion
func (n NWTLSParameters) MaximumSSLProtocolVersion() uint {
	rv := objc.Send[uint](n.ID, objc.Sel("maximumSSLProtocolVersion"))
	return rv
}
func (n NWTLSParameters) SetMaximumSSLProtocolVersion(value uint) {
	objc.Send[struct{}](n.ID, objc.Sel("setMaximumSSLProtocolVersion:"), value)
}
