// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NEPacket] class.
var (
	_NEPacketClass     NEPacketClass
	_NEPacketClassOnce sync.Once
)

func getNEPacketClass() NEPacketClass {
	_NEPacketClassOnce.Do(func() {
		_NEPacketClass = NEPacketClass{class: objc.GetClass("NEPacket")}
	})
	return _NEPacketClass
}

// GetNEPacketClass returns the class object for NEPacket.
func GetNEPacketClass() NEPacketClass {
	return getNEPacketClass()
}

type NEPacketClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NEPacketClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NEPacketClass) Alloc() NEPacket {
	rv := objc.Send[NEPacket](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A network packet and its associated properties.
//
// # Initializing a packet
//
//   - [NEPacket.InitWithDataProtocolFamily]
//
// # Accessing packet properties
//
//   - [NEPacket.Data]
//   - [NEPacket.Metadata]
//   - [NEPacket.ProtocolFamily]
//   - [NEPacket.Direction]: The direction of the packet.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEPacket
type NEPacket struct {
	objectivec.Object
}

// NEPacketFromID constructs a [NEPacket] from an objc.ID.
//
// A network packet and its associated properties.
func NEPacketFromID(id objc.ID) NEPacket {
	return NEPacket{objectivec.Object{ID: id}}
}

// NOTE: NEPacket adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NEPacket] class.
//
// # Initializing a packet
//
//   - [INEPacket.InitWithDataProtocolFamily]
//
// # Accessing packet properties
//
//   - [INEPacket.Data]
//   - [INEPacket.Metadata]
//   - [INEPacket.ProtocolFamily]
//   - [INEPacket.Direction]: The direction of the packet.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEPacket
type INEPacket interface {
	objectivec.IObject

	// Topic: Initializing a packet

	InitWithDataProtocolFamily(data foundation.INSData, protocolFamily uint8) NEPacket

	// Topic: Accessing packet properties

	Data() foundation.INSData
	Metadata() INEFlowMetaData
	ProtocolFamily() uint8
	// The direction of the packet.
	Direction() NETrafficDirection

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (p NEPacket) Init() NEPacket {
	rv := objc.Send[NEPacket](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p NEPacket) Autorelease() NEPacket {
	rv := objc.Send[NEPacket](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEPacket creates a new NEPacket instance.
func NewNEPacket() NEPacket {
	class := getNEPacketClass()
	rv := objc.Send[NEPacket](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/NetworkExtension/NEPacket/init(data:protocolFamily:)
func NewPacketWithDataProtocolFamily(data foundation.INSData, protocolFamily uint8) NEPacket {
	instance := getNEPacketClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithData:protocolFamily:"), data, protocolFamily)
	return NEPacketFromID(rv)
}

// See: https://developer.apple.com/documentation/NetworkExtension/NEPacket/init(data:protocolFamily:)
func (p NEPacket) InitWithDataProtocolFamily(data foundation.INSData, protocolFamily uint8) NEPacket {
	rv := objc.Send[NEPacket](p.ID, objc.Sel("initWithData:protocolFamily:"), data, protocolFamily)
	return rv
}
func (p NEPacket) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](p.ID, objc.Sel("encodeWithCoder:"), coder)
}

// See: https://developer.apple.com/documentation/NetworkExtension/NEPacket/data
func (p NEPacket) Data() foundation.INSData {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("data"))
	return foundation.NSDataFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/NetworkExtension/NEPacket/metadata
func (p NEPacket) Metadata() INEFlowMetaData {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("metadata"))
	return NEFlowMetaDataFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/NetworkExtension/NEPacket/protocolFamily
func (p NEPacket) ProtocolFamily() uint8 {
	rv := objc.Send[uint8](p.ID, objc.Sel("protocolFamily"))
	return rv
}

// The direction of the packet.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEPacket/direction
func (p NEPacket) Direction() NETrafficDirection {
	rv := objc.Send[NETrafficDirection](p.ID, objc.Sel("direction"))
	return NETrafficDirection(rv)
}
