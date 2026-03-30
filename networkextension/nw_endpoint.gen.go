// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NWEndpoint] class.
var (
	_NWEndpointClass     NWEndpointClass
	_NWEndpointClassOnce sync.Once
)

func getNWEndpointClass() NWEndpointClass {
	_NWEndpointClassOnce.Do(func() {
		_NWEndpointClass = NWEndpointClass{class: objc.GetClass("NWEndpoint")}
	})
	return _NWEndpointClass
}

// GetNWEndpointClass returns the class object for NWEndpoint.
func GetNWEndpointClass() NWEndpointClass {
	return getNWEndpointClass()
}

type NWEndpointClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NWEndpointClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NWEndpointClass) Alloc() NWEndpoint {
	rv := objc.Send[NWEndpoint](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// An abstract base class, shared by [NWHostEndpoint] or
// [NWBonjourServiceEndpoint], that represents the source or destination of a
// network connection.
//
// # Overview
//
// All endpoint objects are static collections of parameters that describe a
// network resource. They do not directly provide any resolution services, but
// instead must be used with other classes to be resolved and create
// connections.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NWEndpoint
type NWEndpoint struct {
	objectivec.Object
}

// NWEndpointFromID constructs a [NWEndpoint] from an objc.ID.
//
// An abstract base class, shared by [NWHostEndpoint] or
// [NWBonjourServiceEndpoint], that represents the source or destination of a
// network connection.
func NWEndpointFromID(id objc.ID) NWEndpoint {
	return NWEndpoint{objectivec.Object{ID: id}}
}

// NOTE: NWEndpoint adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NWEndpoint] class.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NWEndpoint
type INWEndpoint interface {
	objectivec.IObject

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (n NWEndpoint) Init() NWEndpoint {
	rv := objc.Send[NWEndpoint](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n NWEndpoint) Autorelease() NWEndpoint {
	rv := objc.Send[NWEndpoint](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewNWEndpoint creates a new NWEndpoint instance.
func NewNWEndpoint() NWEndpoint {
	class := getNWEndpointClass()
	rv := objc.Send[NWEndpoint](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (n NWEndpoint) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](n.ID, objc.Sel("encodeWithCoder:"), coder)
}
