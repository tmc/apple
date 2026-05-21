// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"
	"unsafe"

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
type ICPXRemoteViewEventServerDefaultConfig interface {
	objectivec.IObject

	// Topic: Methods

	ListenerDomain() objectivec.IObject
	ListenerService() objectivec.IObject
	RemoteViewEventManagerForConnection(connection objectivec.IObject) objectivec.IObject
	ServiceInterface() objectivec.IObject
	SignEventSigningKey(event unsafe.Pointer, key SLSSigningKeyRef)
	SigningKeyForConnection(connection objectivec.IObject) objectivec.IObject
	VerifyEventOrderMostRecentEventTime(order unsafe.Pointer, time uint64) bool
	VerifyEventSignatureSigningKey(signature unsafe.Pointer, key SLSSigningKeyRef) bool
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objectivec.Class
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

func (c CPXRemoteViewEventServerDefaultConfig) ListenerDomain() objectivec.IObject {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("listenerDomain"))
	return objectivec.Object{ID: rv}
}
func (c CPXRemoteViewEventServerDefaultConfig) ListenerService() objectivec.IObject {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("listenerService"))
	return objectivec.Object{ID: rv}
}
func (c CPXRemoteViewEventServerDefaultConfig) RemoteViewEventManagerForConnection(connection objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("remoteViewEventManagerForConnection:"), connection)
	return objectivec.Object{ID: rv}
}
func (c CPXRemoteViewEventServerDefaultConfig) ServiceInterface() objectivec.IObject {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("serviceInterface"))
	return objectivec.Object{ID: rv}
}
func (c CPXRemoteViewEventServerDefaultConfig) SignEventSigningKey(event unsafe.Pointer, key SLSSigningKeyRef) {
	objc.Send[objc.ID](c.ID, objc.Sel("signEvent:signingKey:"), event, key)
}
func (c CPXRemoteViewEventServerDefaultConfig) SigningKeyForConnection(connection objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("signingKeyForConnection:"), connection)
	return objectivec.Object{ID: rv}
}
func (c CPXRemoteViewEventServerDefaultConfig) VerifyEventOrderMostRecentEventTime(order unsafe.Pointer, time uint64) bool {
	rv := objc.Send[bool](c.ID, objc.Sel("verifyEventOrder:mostRecentEventTime:"), order, time)
	return rv
}
func (c CPXRemoteViewEventServerDefaultConfig) VerifyEventSignatureSigningKey(signature unsafe.Pointer, key SLSSigningKeyRef) bool {
	rv := objc.Send[bool](c.ID, objc.Sel("verifyEventSignature:signingKey:"), signature, key)
	return rv
}

func (c CPXRemoteViewEventServerDefaultConfig) DebugDescription() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
func (c CPXRemoteViewEventServerDefaultConfig) Description() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
func (c CPXRemoteViewEventServerDefaultConfig) Hash() uint64 {
	rv := objc.Send[uint64](c.ID, objc.Sel("hash"))
	return rv
}
func (c CPXRemoteViewEventServerDefaultConfig) Superclass() objectivec.Class {
	rv := objc.Send[objectivec.Class](c.ID, objc.Sel("superclass"))
	return objectivec.Class(rv)
}
