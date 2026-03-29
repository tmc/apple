// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [NEFilterPacketProvider] class.
var (
	_NEFilterPacketProviderClass     NEFilterPacketProviderClass
	_NEFilterPacketProviderClassOnce sync.Once
)

func getNEFilterPacketProviderClass() NEFilterPacketProviderClass {
	_NEFilterPacketProviderClassOnce.Do(func() {
		_NEFilterPacketProviderClass = NEFilterPacketProviderClass{class: objc.GetClass("NEFilterPacketProvider")}
	})
	return _NEFilterPacketProviderClass
}

// GetNEFilterPacketProviderClass returns the class object for NEFilterPacketProvider.
func GetNEFilterPacketProviderClass() NEFilterPacketProviderClass {
	return getNEFilterPacketProviderClass()
}

type NEFilterPacketProviderClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NEFilterPacketProviderClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NEFilterPacketProviderClass) Alloc() NEFilterPacketProvider {
	rv := objc.Send[NEFilterPacketProvider](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A filter provider that evaluates network packets and decides whether to
// block, allow, or delay the packets.
//
// # Filtering packets
//
//   - [NEFilterPacketProvider.PacketHandler]: A Swift closure or an ObjectiveC block that handles each packet received by the filter.
//   - [NEFilterPacketProvider.SetPacketHandler]
//
// # Delaying packets
//
//   - [NEFilterPacketProvider.DelayCurrentPacket]: Delay a packet currently processed by a packet handler.
//   - [NEFilterPacketProvider.AllowPacket]: Allow delivery of a previously-delayed packet.
//
// # Instance Properties
//
//   - [NEFilterPacketProvider.Handler]
//   - [NEFilterPacketProvider.SetHandler]
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEFilterPacketProvider
type NEFilterPacketProvider struct {
	NEFilterProvider
}

// NEFilterPacketProviderFromID constructs a [NEFilterPacketProvider] from an objc.ID.
//
// A filter provider that evaluates network packets and decides whether to
// block, allow, or delay the packets.
func NEFilterPacketProviderFromID(id objc.ID) NEFilterPacketProvider {
	return NEFilterPacketProvider{NEFilterProvider: NEFilterProviderFromID(id)}
}
// NOTE: NEFilterPacketProvider adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NEFilterPacketProvider] class.
//
// # Filtering packets
//
//   - [INEFilterPacketProvider.PacketHandler]: A Swift closure or an ObjectiveC block that handles each packet received by the filter.
//   - [INEFilterPacketProvider.SetPacketHandler]
//
// # Delaying packets
//
//   - [INEFilterPacketProvider.DelayCurrentPacket]: Delay a packet currently processed by a packet handler.
//   - [INEFilterPacketProvider.AllowPacket]: Allow delivery of a previously-delayed packet.
//
// # Instance Properties
//
//   - [INEFilterPacketProvider.Handler]
//   - [INEFilterPacketProvider.SetHandler]
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEFilterPacketProvider
type INEFilterPacketProvider interface {
	INEFilterProvider

	// Topic: Filtering packets

	// A Swift closure or an ObjectiveC block that handles each packet received by the filter.
	PacketHandler() NEFilterPacketHandler
	SetPacketHandler(value NEFilterPacketHandler)

	// Topic: Delaying packets

	// Delay a packet currently processed by a packet handler.
	DelayCurrentPacket(context INEFilterPacketContext) INEPacket
	// Allow delivery of a previously-delayed packet.
	AllowPacket(packet INEPacket)

	// Topic: Instance Properties

	Handler() NEFilterPacketProviderVerdict
	SetHandler(value NEFilterPacketProviderVerdict)
}

// Init initializes the instance.
func (f NEFilterPacketProvider) Init() NEFilterPacketProvider {
	rv := objc.Send[NEFilterPacketProvider](f.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f NEFilterPacketProvider) Autorelease() NEFilterPacketProvider {
	rv := objc.Send[NEFilterPacketProvider](f.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEFilterPacketProvider creates a new NEFilterPacketProvider instance.
func NewNEFilterPacketProvider() NEFilterPacketProvider {
	class := getNEFilterPacketProviderClass()
	rv := objc.Send[NEFilterPacketProvider](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Delay a packet currently processed by a packet handler.
//
// context: A context for the packet handler.
//
// # Discussion
// 
// This function is only valid within the [PacketHandler] Swift closure or
// ObjectiveC block, which must return [NEFilterPacketProvider.Verdict.delay]
// after delaying the packet. The framework prevents further delivery of the
// packet through the network stack until it’s allowed or dropped. Allow the
// packet by calling [AllowPacket], or drop it by releasing it.
//
// [NEFilterPacketProvider.Verdict.delay]: https://developer.apple.com/documentation/NetworkExtension/NEFilterPacketProvider/Verdict/delay
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEFilterPacketProvider/delayCurrentPacket(_:)
func (f NEFilterPacketProvider) DelayCurrentPacket(context INEFilterPacketContext) INEPacket {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("delayCurrentPacket:"), context)
	return NEPacketFromID(rv)
}
// Allow delivery of a previously-delayed packet.
//
// packet: The packet previously delayed by the packet handler.
//
// # Discussion
// 
// Use this method to allow a previously-delayed packet to continue its
// journey into or out of the networking stack.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEFilterPacketProvider/allow(_:)
func (f NEFilterPacketProvider) AllowPacket(packet INEPacket) {
	objc.Send[objc.ID](f.ID, objc.Sel("allowPacket:"), packet)
}

// A Swift closure or an ObjectiveC block that handles each packet received by
// the filter.
//
// # Discussion
// 
// Set this property to a handler that returns a
// [NEFilterPacketProvider.Verdict] for each packet it receives.
// 
// Since there may be multiple filtering sources presenting frames to the
// provider, multiple simultaneous threads may execute this packet handler.
// Therefore, the packet handler must be able to handle execution in a
// multi-threaded environment.
//
// [NEFilterPacketProvider.Verdict]: https://developer.apple.com/documentation/NetworkExtension/NEFilterPacketProvider/Verdict
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEFilterPacketProvider/packetHandler
func (f NEFilterPacketProvider) PacketHandler() NEFilterPacketHandler {
	rv := objc.Send[NEFilterPacketHandler](f.ID, objc.Sel("packetHandler"))
	return NEFilterPacketHandler(rv)
}
func (f NEFilterPacketProvider) SetPacketHandler(value NEFilterPacketHandler) {
	objc.Send[struct{}](f.ID, objc.Sel("setPacketHandler:"), value)
}
// See: https://developer.apple.com/documentation/networkextension/nefilterpacketprovider/handler
func (f NEFilterPacketProvider) Handler() NEFilterPacketProviderVerdict {
	rv := objc.Send[NEFilterPacketProviderVerdict](f.ID, objc.Sel("handler"))
	return NEFilterPacketProviderVerdict(rv)
}
func (f NEFilterPacketProvider) SetHandler(value NEFilterPacketProviderVerdict) {
	objc.Send[struct{}](f.ID, objc.Sel("setHandler:"), value)
}

