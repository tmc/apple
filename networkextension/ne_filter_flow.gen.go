// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NEFilterFlow] class.
var (
	_NEFilterFlowClass     NEFilterFlowClass
	_NEFilterFlowClassOnce sync.Once
)

func getNEFilterFlowClass() NEFilterFlowClass {
	_NEFilterFlowClassOnce.Do(func() {
		_NEFilterFlowClass = NEFilterFlowClass{class: objc.GetClass("NEFilterFlow")}
	})
	return _NEFilterFlowClass
}

// GetNEFilterFlowClass returns the class object for NEFilterFlow.
func GetNEFilterFlowClass() NEFilterFlowClass {
	return getNEFilterFlowClass()
}

type NEFilterFlowClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NEFilterFlowClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NEFilterFlowClass) Alloc() NEFilterFlow {
	rv := objc.Send[NEFilterFlow](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// The abstract base class for types that represent flows of network data.
//
// # Inspecting flow properties
//
//   - [NEFilterFlow.URL]: The flow’s HTTP URL.
//   - [NEFilterFlow.Identifier]: The unique identifier of the flow.
//   - [NEFilterFlow.Direction]: The initial direction of the flow: incoming or outgoing.
//   - [NEFilterFlow.NEFilterFlowBytesMax]: The maximum number of bytes to pass or peek for a flow.
//   - [NEFilterFlow.SetNEFilterFlowBytesMax]
//
// # Source app identification
//
//   - [NEFilterFlow.SourceAppAuditToken]: The audit token of the source application of the flow.
//   - [NEFilterFlow.SourceProcessAuditToken]: The audit token of the process that created the flow.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEFilterFlow
type NEFilterFlow struct {
	objectivec.Object
}

// NEFilterFlowFromID constructs a [NEFilterFlow] from an objc.ID.
//
// The abstract base class for types that represent flows of network data.
func NEFilterFlowFromID(id objc.ID) NEFilterFlow {
	return NEFilterFlow{objectivec.Object{ID: id}}
}
// NOTE: NEFilterFlow adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NEFilterFlow] class.
//
// # Inspecting flow properties
//
//   - [INEFilterFlow.URL]: The flow’s HTTP URL.
//   - [INEFilterFlow.Identifier]: The unique identifier of the flow.
//   - [INEFilterFlow.Direction]: The initial direction of the flow: incoming or outgoing.
//   - [INEFilterFlow.NEFilterFlowBytesMax]: The maximum number of bytes to pass or peek for a flow.
//   - [INEFilterFlow.SetNEFilterFlowBytesMax]
//
// # Source app identification
//
//   - [INEFilterFlow.SourceAppAuditToken]: The audit token of the source application of the flow.
//   - [INEFilterFlow.SourceProcessAuditToken]: The audit token of the process that created the flow.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEFilterFlow
type INEFilterFlow interface {
	objectivec.IObject

	// Topic: Inspecting flow properties

	// The flow’s HTTP URL.
	URL() foundation.INSURL
	// The unique identifier of the flow.
	Identifier() foundation.NSUUID
	// The initial direction of the flow: incoming or outgoing.
	Direction() NETrafficDirection
	// The maximum number of bytes to pass or peek for a flow.
	NEFilterFlowBytesMax() uint64
	SetNEFilterFlowBytesMax(value uint64)

	// Topic: Source app identification

	// The audit token of the source application of the flow.
	SourceAppAuditToken() foundation.INSData
	// The audit token of the process that created the flow.
	SourceProcessAuditToken() foundation.INSData

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (f NEFilterFlow) Init() NEFilterFlow {
	rv := objc.Send[NEFilterFlow](f.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f NEFilterFlow) Autorelease() NEFilterFlow {
	rv := objc.Send[NEFilterFlow](f.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEFilterFlow creates a new NEFilterFlow instance.
func NewNEFilterFlow() NEFilterFlow {
	class := getNEFilterFlowClass()
	rv := objc.Send[NEFilterFlow](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (f NEFilterFlow) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](f.ID, objc.Sel("encodeWithCoder:"), coder)
}

// The flow’s HTTP URL.
//
// # Discussion
// 
// This parameter is only non-`nil` for flows that originate from WebKit
// browser objects.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEFilterFlow/url
func (f NEFilterFlow) URL() foundation.INSURL {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("URL"))
	return foundation.NSURLFromID(objc.ID(rv))
}
// The unique identifier of the flow.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEFilterFlow/identifier
func (f NEFilterFlow) Identifier() foundation.NSUUID {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("identifier"))
	return foundation.NSUUIDFromID(objc.ID(rv))
}
// The initial direction of the flow: incoming or outgoing.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEFilterFlow/direction
func (f NEFilterFlow) Direction() NETrafficDirection {
	rv := objc.Send[NETrafficDirection](f.ID, objc.Sel("direction"))
	return NETrafficDirection(rv)
}
// The maximum number of bytes to pass or peek for a flow.
//
// See: https://developer.apple.com/documentation/networkextension/nefilterflowbytesmax
func (f NEFilterFlow) NEFilterFlowBytesMax() uint64 {
	rv := objc.Send[uint64](f.ID, objc.Sel("NEFilterFlowBytesMax"))
	return rv
}
func (f NEFilterFlow) SetNEFilterFlowBytesMax(value uint64) {
	objc.Send[struct{}](f.ID, objc.Sel("setNEFilterFlowBytesMax:"), value)
}
// The audit token of the source application of the flow.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEFilterFlow/sourceAppAuditToken
func (f NEFilterFlow) SourceAppAuditToken() foundation.INSData {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("sourceAppAuditToken"))
	return foundation.NSDataFromID(objc.ID(rv))
}
// The audit token of the process that created the flow.
//
// # Discussion
// 
// In cases where a system process creates the connection on behalf of a
// source app, this value is different from [SourceAppAuditToken]. In cases
// where the source app directly creates the connection, these values are
// identical.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEFilterFlow/sourceProcessAuditToken
func (f NEFilterFlow) SourceProcessAuditToken() foundation.INSData {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("sourceProcessAuditToken"))
	return foundation.NSDataFromID(objc.ID(rv))
}

