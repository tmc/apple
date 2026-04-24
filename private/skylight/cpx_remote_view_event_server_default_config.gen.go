// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CPXRemoteViewEventServerDefaultConfig] class.
var (
	_CPXRemoteViewEventServerDefaultConfigClass     CPXRemoteViewEventServerDefaultConfigClass
	_CPXRemoteViewEventServerDefaultConfigClassOnce sync.Once
)

func getCPXRemoteViewEventServerDefaultConfigClass() CPXRemoteViewEventServerDefaultConfigClass {
	_CPXRemoteViewEventServerDefaultConfigClassOnce.Do(func() {
		_CPXRemoteViewEventServerDefaultConfigClass = CPXRemoteViewEventServerDefaultConfigClass{class: objc.GetClass("CPXRemoteViewEventServerDefaultConfig")}
	})
	return _CPXRemoteViewEventServerDefaultConfigClass
}

// GetCPXRemoteViewEventServerDefaultConfigClass returns the class object for CPXRemoteViewEventServerDefaultConfig.
func GetCPXRemoteViewEventServerDefaultConfigClass() CPXRemoteViewEventServerDefaultConfigClass {
	return getCPXRemoteViewEventServerDefaultConfigClass()
}

type CPXRemoteViewEventServerDefaultConfigClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CPXRemoteViewEventServerDefaultConfigClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CPXRemoteViewEventServerDefaultConfigClass) Alloc() CPXRemoteViewEventServerDefaultConfig {
	rv := objc.Send[CPXRemoteViewEventServerDefaultConfig](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [CPXRemoteViewEventServerDefaultConfig.ListenerDomain]
//   - [CPXRemoteViewEventServerDefaultConfig.ListenerService]
//   - [CPXRemoteViewEventServerDefaultConfig.RemoteViewEventManagerForConnection]
//   - [CPXRemoteViewEventServerDefaultConfig.ServiceInterface]
//   - [CPXRemoteViewEventServerDefaultConfig.SignEventSigningKey]
//   - [CPXRemoteViewEventServerDefaultConfig.SigningKeyForConnection]
//   - [CPXRemoteViewEventServerDefaultConfig.VerifyEventOrderMostRecentEventTime]
//   - [CPXRemoteViewEventServerDefaultConfig.VerifyEventSignatureSigningKey]
//   - [CPXRemoteViewEventServerDefaultConfig.DebugDescription]
//   - [CPXRemoteViewEventServerDefaultConfig.Description]
//   - [CPXRemoteViewEventServerDefaultConfig.Hash]
//   - [CPXRemoteViewEventServerDefaultConfig.Superclass]
//
// See: https://developer.apple.com/documentation/SkyLight/CPXRemoteViewEventServerDefaultConfig
type CPXRemoteViewEventServerDefaultConfig struct {
	objectivec.Object
}

// CPXRemoteViewEventServerDefaultConfigFromID constructs a [CPXRemoteViewEventServerDefaultConfig] from an objc.ID.
func CPXRemoteViewEventServerDefaultConfigFromID(id objc.ID) CPXRemoteViewEventServerDefaultConfig {
	return CPXRemoteViewEventServerDefaultConfig{objectivec.Object{ID: id}}
}

// Ensure CPXRemoteViewEventServerDefaultConfig implements ICPXRemoteViewEventServerDefaultConfig.
var _ ICPXRemoteViewEventServerDefaultConfig = CPXRemoteViewEventServerDefaultConfig{}

// An interface definition for the [CPXRemoteViewEventServerDefaultConfig] class.
//
// # Methods
//
//   - [ICPXRemoteViewEventServerDefaultConfig.ListenerDomain]
//   - [ICPXRemoteViewEventServerDefaultConfig.ListenerService]
//   - [ICPXRemoteViewEventServerDefaultConfig.RemoteViewEventManagerForConnection]
//   - [ICPXRemoteViewEventServerDefaultConfig.ServiceInterface]
//   - [ICPXRemoteViewEventServerDefaultConfig.SignEventSigningKey]
//   - [ICPXRemoteViewEventServerDefaultConfig.SigningKeyForConnection]
//   - [ICPXRemoteViewEventServerDefaultConfig.VerifyEventOrderMostRecentEventTime]
//   - [ICPXRemoteViewEventServerDefaultConfig.VerifyEventSignatureSigningKey]
//   - [ICPXRemoteViewEventServerDefaultConfig.DebugDescription]
//   - [ICPXRemoteViewEventServerDefaultConfig.Description]
//   - [ICPXRemoteViewEventServerDefaultConfig.Hash]
//   - [ICPXRemoteViewEventServerDefaultConfig.Superclass]
//
// See: https://developer.apple.com/documentation/SkyLight/CPXRemoteViewEventServerDefaultConfig
type ICPXRemoteViewEventServerDefaultConfig interface {
	objectivec.IObject

	// Topic: Methods

	ListenerDomain() objectivec.IObject
	ListenerService() objectivec.IObject
	RemoteViewEventManagerForConnection(connection objectivec.IObject) objectivec.IObject
	ServiceInterface() objectivec.IObject
	SignEventSigningKey(event objectivec.IObject, key SLSSigningKeyRef)
	SigningKeyForConnection(connection objectivec.IObject) objectivec.IObject
	VerifyEventOrderMostRecentEventTime(order objectivec.IObject, time uint64) bool
	VerifyEventSignatureSigningKey(signature objectivec.IObject, key SLSSigningKeyRef) bool
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objc.Class
}

// Init initializes the instance.
func (c CPXRemoteViewEventServerDefaultConfig) Init() CPXRemoteViewEventServerDefaultConfig {
	rv := objc.Send[CPXRemoteViewEventServerDefaultConfig](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CPXRemoteViewEventServerDefaultConfig) Autorelease() CPXRemoteViewEventServerDefaultConfig {
	rv := objc.Send[CPXRemoteViewEventServerDefaultConfig](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCPXRemoteViewEventServerDefaultConfig creates a new CPXRemoteViewEventServerDefaultConfig instance.
func NewCPXRemoteViewEventServerDefaultConfig() CPXRemoteViewEventServerDefaultConfig {
	class := getCPXRemoteViewEventServerDefaultConfigClass()
	rv := objc.Send[CPXRemoteViewEventServerDefaultConfig](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/CPXRemoteViewEventServerDefaultConfig/listenerDomain
func (c CPXRemoteViewEventServerDefaultConfig) ListenerDomain() objectivec.IObject {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("listenerDomain"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/CPXRemoteViewEventServerDefaultConfig/listenerService
func (c CPXRemoteViewEventServerDefaultConfig) ListenerService() objectivec.IObject {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("listenerService"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/CPXRemoteViewEventServerDefaultConfig/remoteViewEventManagerForConnection:
func (c CPXRemoteViewEventServerDefaultConfig) RemoteViewEventManagerForConnection(connection objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("remoteViewEventManagerForConnection:"), connection)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/CPXRemoteViewEventServerDefaultConfig/serviceInterface
func (c CPXRemoteViewEventServerDefaultConfig) ServiceInterface() objectivec.IObject {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("serviceInterface"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/CPXRemoteViewEventServerDefaultConfig/signEvent:signingKey:
func (c CPXRemoteViewEventServerDefaultConfig) SignEventSigningKey(event objectivec.IObject, key SLSSigningKeyRef) {
	objc.Send[objc.ID](c.ID, objc.Sel("signEvent:signingKey:"), event, key)
}

// See: https://developer.apple.com/documentation/SkyLight/CPXRemoteViewEventServerDefaultConfig/signingKeyForConnection:
func (c CPXRemoteViewEventServerDefaultConfig) SigningKeyForConnection(connection objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("signingKeyForConnection:"), connection)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/CPXRemoteViewEventServerDefaultConfig/verifyEventOrder:mostRecentEventTime:
func (c CPXRemoteViewEventServerDefaultConfig) VerifyEventOrderMostRecentEventTime(order objectivec.IObject, time uint64) bool {
	rv := objc.Send[bool](c.ID, objc.Sel("verifyEventOrder:mostRecentEventTime:"), order, time)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/CPXRemoteViewEventServerDefaultConfig/verifyEventSignature:signingKey:
func (c CPXRemoteViewEventServerDefaultConfig) VerifyEventSignatureSigningKey(signature objectivec.IObject, key SLSSigningKeyRef) bool {
	rv := objc.Send[bool](c.ID, objc.Sel("verifyEventSignature:signingKey:"), signature, key)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/CPXRemoteViewEventServerDefaultConfig/debugDescription
func (c CPXRemoteViewEventServerDefaultConfig) DebugDescription() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/SkyLight/CPXRemoteViewEventServerDefaultConfig/description
func (c CPXRemoteViewEventServerDefaultConfig) Description() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/SkyLight/CPXRemoteViewEventServerDefaultConfig/hash
func (c CPXRemoteViewEventServerDefaultConfig) Hash() uint64 {
	rv := objc.Send[uint64](c.ID, objc.Sel("hash"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/CPXRemoteViewEventServerDefaultConfig/superclass
func (c CPXRemoteViewEventServerDefaultConfig) Superclass() objc.Class {
	rv := objc.Send[objc.Class](c.ID, objc.Sel("superclass"))
	return rv
}
