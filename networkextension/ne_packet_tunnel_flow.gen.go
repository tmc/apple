// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NEPacketTunnelFlow] class.
var (
	_NEPacketTunnelFlowClass     NEPacketTunnelFlowClass
	_NEPacketTunnelFlowClassOnce sync.Once
)

func getNEPacketTunnelFlowClass() NEPacketTunnelFlowClass {
	_NEPacketTunnelFlowClassOnce.Do(func() {
		_NEPacketTunnelFlowClass = NEPacketTunnelFlowClass{class: objc.GetClass("NEPacketTunnelFlow")}
	})
	return _NEPacketTunnelFlowClass
}

// GetNEPacketTunnelFlowClass returns the class object for NEPacketTunnelFlow.
func GetNEPacketTunnelFlowClass() NEPacketTunnelFlowClass {
	return getNEPacketTunnelFlowClass()
}

type NEPacketTunnelFlowClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NEPacketTunnelFlowClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NEPacketTunnelFlowClass) Alloc() NEPacketTunnelFlow {
	rv := objc.Send[NEPacketTunnelFlow](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// An object you use to read and write packets to and from the tunnel’s
// virtual interface.
//
// # Overview
//
// Use the [NEPacketTunnelFlow] class to implement a custom-IP tunneling
// protocol for your packet tunnel. For example, use the APIs in this class to
// read packets from the virtual interface, so you can then encapsulate these
// packets and send them to a packet-tunnel server. Likewise, read packets
// from your packet-tunnel server and use these APIs to write the packets back
// to the tunnel’s virtual interface.
//
// # Handling IP packets
//
//   - [NEPacketTunnelFlow.WritePacketObjects]: Write multiple IP packets to the TUN interface.
//   - [NEPacketTunnelFlow.ReadPacketsWithCompletionHandler]: Reads IP packets from the TUN interface.
//   - [NEPacketTunnelFlow.WritePacketsWithProtocols]: Writes IP packets to the TUN interface.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEPacketTunnelFlow
type NEPacketTunnelFlow struct {
	objectivec.Object
}

// NEPacketTunnelFlowFromID constructs a [NEPacketTunnelFlow] from an objc.ID.
//
// An object you use to read and write packets to and from the tunnel’s
// virtual interface.
func NEPacketTunnelFlowFromID(id objc.ID) NEPacketTunnelFlow {
	return NEPacketTunnelFlow{objectivec.Object{ID: id}}
}

// NOTE: NEPacketTunnelFlow adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NEPacketTunnelFlow] class.
//
// # Handling IP packets
//
//   - [INEPacketTunnelFlow.WritePacketObjects]: Write multiple IP packets to the TUN interface.
//   - [INEPacketTunnelFlow.ReadPacketsWithCompletionHandler]: Reads IP packets from the TUN interface.
//   - [INEPacketTunnelFlow.WritePacketsWithProtocols]: Writes IP packets to the TUN interface.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEPacketTunnelFlow
type INEPacketTunnelFlow interface {
	objectivec.IObject

	// Topic: Handling IP packets

	// Write multiple IP packets to the TUN interface.
	WritePacketObjects(packets []NEPacket) bool
	// Reads IP packets from the TUN interface.
	ReadPacketsWithCompletionHandler(completionHandler VoidHandler)
	// Writes IP packets to the TUN interface.
	WritePacketsWithProtocols(packets []foundation.NSData, protocols []foundation.NSNumber) bool
}

// Init initializes the instance.
func (p NEPacketTunnelFlow) Init() NEPacketTunnelFlow {
	rv := objc.Send[NEPacketTunnelFlow](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p NEPacketTunnelFlow) Autorelease() NEPacketTunnelFlow {
	rv := objc.Send[NEPacketTunnelFlow](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEPacketTunnelFlow creates a new NEPacketTunnelFlow instance.
func NewNEPacketTunnelFlow() NEPacketTunnelFlow {
	class := getNEPacketTunnelFlowClass()
	rv := objc.Send[NEPacketTunnelFlow](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Write multiple IP packets to the TUN interface.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEPacketTunnelFlow/writePacketObjects(_:)
func (p NEPacketTunnelFlow) WritePacketObjects(packets []NEPacket) bool {
	rv := objc.Send[bool](p.ID, objc.Sel("writePacketObjects:"), objectivec.IObjectSliceToNSArray(packets))
	return rv
}

// Reads IP packets from the TUN interface.
//
// completionHandler: A Swift closure or an ObjectiveC block that runs when some packets are read
// from the TUN interface. The packets that were read are passed to this block
// in the `packets` array. The protocol numbers of the packets that were read
// are passed to this block in the `protocols` array. Each packet has a
// protocol number in the corresponding index in the `protocols` array. The
// protocol numbers are given in host byte order. Valid protocol numbers
// include `AF_INET` and `AF_INET6`. See `/usr/include/sys/socket.H()`.
//
// # Discussion
//
// Each call to this method results in a single execution of the completion
// handler. The caller should call this method after each `completionHandler`
// execution in order to continue to receive packets from the TUN interface.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEPacketTunnelFlow/readPackets(completionHandler:)
func (p NEPacketTunnelFlow) ReadPacketsWithCompletionHandler(completionHandler VoidHandler) {
	_block0, _ := NewVoidBlock(completionHandler)
	objc.Send[objc.ID](p.ID, objc.Sel("readPacketsWithCompletionHandler:"), _block0)
}

// Writes IP packets to the TUN interface.
//
// packets: An array of NSData objects containing the IP packets to the written.
//
// protocols: An array of NSNumber objects containing the protocol numbers (e.g. AF_INET
// or AF_INET6) of the IP packets in `packets` in host byte order.
//
// # Discussion
//
// The number of NSData objects in `packets` must be exactly equal to the
// number of NSNumber objects in `protocols`.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEPacketTunnelFlow/writePackets(_:withProtocols:)
func (p NEPacketTunnelFlow) WritePacketsWithProtocols(packets []foundation.NSData, protocols []foundation.NSNumber) bool {
	rv := objc.Send[bool](p.ID, objc.Sel("writePackets:withProtocols:"), objectivec.IObjectSliceToNSArray(packets), objectivec.IObjectSliceToNSArray(protocols))
	return rv
}
