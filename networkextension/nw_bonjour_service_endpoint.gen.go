// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
)

// The class instance for the [NWBonjourServiceEndpoint] class.
var (
	_NWBonjourServiceEndpointClass     NWBonjourServiceEndpointClass
	_NWBonjourServiceEndpointClassOnce sync.Once
)

func getNWBonjourServiceEndpointClass() NWBonjourServiceEndpointClass {
	_NWBonjourServiceEndpointClassOnce.Do(func() {
		_NWBonjourServiceEndpointClass = NWBonjourServiceEndpointClass{class: objc.GetClass("NWBonjourServiceEndpoint")}
	})
	return _NWBonjourServiceEndpointClass
}

// GetNWBonjourServiceEndpointClass returns the class object for NWBonjourServiceEndpoint.
func GetNWBonjourServiceEndpointClass() NWBonjourServiceEndpointClass {
	return getNWBonjourServiceEndpointClass()
}

type NWBonjourServiceEndpointClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NWBonjourServiceEndpointClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NWBonjourServiceEndpointClass) Alloc() NWBonjourServiceEndpoint {
	rv := objc.Send[NWBonjourServiceEndpoint](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A network endpoint specified as a Bonjour service name, type, and domain.
//
// # Overview
// 
// For example, the Bonjour service `MyMusicStudio._music._tcp.Local().` has
// the name `"MyMusicStudio"`, the type `"_music._tcp"`, and the domain
// `"local"`.
//
// # Getting endpoint properties
//
//   - [NWBonjourServiceEndpoint.Name]: The endpoint’s Bonjour service name.
//   - [NWBonjourServiceEndpoint.Type]: The endpoint’s Bonjour service type.
//   - [NWBonjourServiceEndpoint.Domain]: The endpoint’s Bonjour service domain, such as `"local"`.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NWBonjourServiceEndpoint
type NWBonjourServiceEndpoint struct {
	NWEndpoint
}

// NWBonjourServiceEndpointFromID constructs a [NWBonjourServiceEndpoint] from an objc.ID.
//
// A network endpoint specified as a Bonjour service name, type, and domain.
func NWBonjourServiceEndpointFromID(id objc.ID) NWBonjourServiceEndpoint {
	return NWBonjourServiceEndpoint{NWEndpoint: NWEndpointFromID(id)}
}
// NOTE: NWBonjourServiceEndpoint adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NWBonjourServiceEndpoint] class.
//
// # Getting endpoint properties
//
//   - [INWBonjourServiceEndpoint.Name]: The endpoint’s Bonjour service name.
//   - [INWBonjourServiceEndpoint.Type]: The endpoint’s Bonjour service type.
//   - [INWBonjourServiceEndpoint.Domain]: The endpoint’s Bonjour service domain, such as `"local"`.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NWBonjourServiceEndpoint
type INWBonjourServiceEndpoint interface {
	INWEndpoint

	// Topic: Getting endpoint properties

	// The endpoint’s Bonjour service name.
	Name() string
	// The endpoint’s Bonjour service type.
	Type() string
	// The endpoint’s Bonjour service domain, such as `"local"`.
	Domain() string
}

// Init initializes the instance.
func (n NWBonjourServiceEndpoint) Init() NWBonjourServiceEndpoint {
	rv := objc.Send[NWBonjourServiceEndpoint](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n NWBonjourServiceEndpoint) Autorelease() NWBonjourServiceEndpoint {
	rv := objc.Send[NWBonjourServiceEndpoint](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewNWBonjourServiceEndpoint creates a new NWBonjourServiceEndpoint instance.
func NewNWBonjourServiceEndpoint() NWBonjourServiceEndpoint {
	class := getNWBonjourServiceEndpointClass()
	rv := objc.Send[NWBonjourServiceEndpoint](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Create an endpoint with a Bonjour service name, type, and domain. All
// fields must be specified.
//
// name: The Bonjour service name.
//
// type: The Bonjour service type.
//
// domain: The Bonjour service domain.
//
// # Return Value
// 
// The new [NWBonjourServiceEndpoint] object.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NWBonjourServiceEndpoint/init(name:type:domain:)
func NewNWBonjourServiceEndpointWithNameTypeDomain(name string, type_ string, domain string) NWBonjourServiceEndpoint {
	rv := objc.Send[objc.ID](objc.ID(getNWBonjourServiceEndpointClass().class), objc.Sel("endpointWithName:type:domain:"), objc.String(name), objc.String(type_), objc.String(domain))
	return NWBonjourServiceEndpointFromID(rv)
}

// The endpoint’s Bonjour service name.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NWBonjourServiceEndpoint/name
func (n NWBonjourServiceEndpoint) Name() string {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("name"))
	return foundation.NSStringFromID(rv).String()
}
// The endpoint’s Bonjour service type.
//
// # Discussion
// 
// For example, the service type could be `"_music._tcp"`.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NWBonjourServiceEndpoint/type
func (n NWBonjourServiceEndpoint) Type() string {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("type"))
	return foundation.NSStringFromID(rv).String()
}
// The endpoint’s Bonjour service domain, such as `"local"`.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NWBonjourServiceEndpoint/domain
func (n NWBonjourServiceEndpoint) Domain() string {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("domain"))
	return foundation.NSStringFromID(rv).String()
}

