// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NEFlowMetaData] class.
var (
	_NEFlowMetaDataClass     NEFlowMetaDataClass
	_NEFlowMetaDataClassOnce sync.Once
)

func getNEFlowMetaDataClass() NEFlowMetaDataClass {
	_NEFlowMetaDataClassOnce.Do(func() {
		_NEFlowMetaDataClass = NEFlowMetaDataClass{class: objc.GetClass("NEFlowMetaData")}
	})
	return _NEFlowMetaDataClass
}

// GetNEFlowMetaDataClass returns the class object for NEFlowMetaData.
func GetNEFlowMetaDataClass() NEFlowMetaDataClass {
	return getNEFlowMetaDataClass()
}

type NEFlowMetaDataClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NEFlowMetaDataClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NEFlowMetaDataClass) Alloc() NEFlowMetaData {
	rv := objc.Send[NEFlowMetaData](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// Additional information about data flowing through a per-app VPN provider.
//
// # Overview
//
// This metadata is only present for data flowing through per-app VPN
// providers, that is, app proxy providers and packet tunnel providers in
// per-app VPN mode, as indicated by the [NEFlowMetaData.RoutingMethod] property.
//
// # Getting source app information
//
//   - [NEFlowMetaData.SourceAppUniqueIdentifier]: A data instance that contains a unique hash value for the source application.
//   - [NEFlowMetaData.SourceAppSigningIdentifier]: A string that contains the signing identifier of the source application.
//   - [NEFlowMetaData.SourceAppAuditToken]: The audit token of the source application of the flow.
//
// # Getting flow information
//
//   - [NEFlowMetaData.FilterFlowIdentifier]: The identifier of the content filter flow corresponding to this flow.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEFlowMetaData
type NEFlowMetaData struct {
	objectivec.Object
}

// NEFlowMetaDataFromID constructs a [NEFlowMetaData] from an objc.ID.
//
// Additional information about data flowing through a per-app VPN provider.
func NEFlowMetaDataFromID(id objc.ID) NEFlowMetaData {
	return NEFlowMetaData{objectivec.Object{ID: id}}
}

// NOTE: NEFlowMetaData adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NEFlowMetaData] class.
//
// # Getting source app information
//
//   - [INEFlowMetaData.SourceAppUniqueIdentifier]: A data instance that contains a unique hash value for the source application.
//   - [INEFlowMetaData.SourceAppSigningIdentifier]: A string that contains the signing identifier of the source application.
//   - [INEFlowMetaData.SourceAppAuditToken]: The audit token of the source application of the flow.
//
// # Getting flow information
//
//   - [INEFlowMetaData.FilterFlowIdentifier]: The identifier of the content filter flow corresponding to this flow.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEFlowMetaData
type INEFlowMetaData interface {
	objectivec.IObject

	// Topic: Getting source app information

	// A data instance that contains a unique hash value for the source application.
	SourceAppUniqueIdentifier() foundation.INSData
	// A string that contains the signing identifier of the source application.
	SourceAppSigningIdentifier() string
	// The audit token of the source application of the flow.
	SourceAppAuditToken() foundation.INSData

	// Topic: Getting flow information

	// The identifier of the content filter flow corresponding to this flow.
	FilterFlowIdentifier() foundation.NSUUID

	// The method by which network traffic is routed to the tunnel.
	RoutingMethod() NETunnelProviderRoutingMethod
	SetRoutingMethod(value NETunnelProviderRoutingMethod)
	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (f NEFlowMetaData) Init() NEFlowMetaData {
	rv := objc.Send[NEFlowMetaData](f.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f NEFlowMetaData) Autorelease() NEFlowMetaData {
	rv := objc.Send[NEFlowMetaData](f.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEFlowMetaData creates a new NEFlowMetaData instance.
func NewNEFlowMetaData() NEFlowMetaData {
	class := getNEFlowMetaDataClass()
	rv := objc.Send[NEFlowMetaData](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (f NEFlowMetaData) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](f.ID, objc.Sel("encodeWithCoder:"), coder)
}

// A data instance that contains a unique hash value for the source
// application.
//
// # Discussion
//
// The property contains the Code Directory Hash for the application.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEFlowMetaData/sourceAppUniqueIdentifier
func (f NEFlowMetaData) SourceAppUniqueIdentifier() foundation.INSData {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("sourceAppUniqueIdentifier"))
	return foundation.NSDataFromID(objc.ID(rv))
}

// A string that contains the signing identifier of the source application.
//
// # Discussion
//
// For all apps that are signed in the standard way using Xcode, this value is
// identical to the app’s bundle identifier.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEFlowMetaData/sourceAppSigningIdentifier
func (f NEFlowMetaData) SourceAppSigningIdentifier() string {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("sourceAppSigningIdentifier"))
	return foundation.NSStringFromID(rv).String()
}

// The audit token of the source application of the flow.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEFlowMetaData/sourceAppAuditToken
func (f NEFlowMetaData) SourceAppAuditToken() foundation.INSData {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("sourceAppAuditToken"))
	return foundation.NSDataFromID(objc.ID(rv))
}

// The identifier of the content filter flow corresponding to this flow.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEFlowMetaData/filterFlowIdentifier
func (f NEFlowMetaData) FilterFlowIdentifier() foundation.NSUUID {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("filterFlowIdentifier"))
	return foundation.NSUUIDFromID(objc.ID(rv))
}

// The method by which network traffic is routed to the tunnel.
//
// See: https://developer.apple.com/documentation/networkextension/netunnelprovider/routingmethod
func (f NEFlowMetaData) RoutingMethod() NETunnelProviderRoutingMethod {
	rv := objc.Send[NETunnelProviderRoutingMethod](f.ID, objc.Sel("routingMethod"))
	return NETunnelProviderRoutingMethod(rv)
}
func (f NEFlowMetaData) SetRoutingMethod(value NETunnelProviderRoutingMethod) {
	objc.Send[struct{}](f.ID, objc.Sel("setRoutingMethod:"), value)
}
