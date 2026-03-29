// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NENetworkRule] class.
var (
	_NENetworkRuleClass     NENetworkRuleClass
	_NENetworkRuleClassOnce sync.Once
)

func getNENetworkRuleClass() NENetworkRuleClass {
	_NENetworkRuleClassOnce.Do(func() {
		_NENetworkRuleClass = NENetworkRuleClass{class: objc.GetClass("NENetworkRule")}
	})
	return _NENetworkRuleClass
}

// GetNENetworkRuleClass returns the class object for NENetworkRule.
func GetNENetworkRuleClass() NENetworkRuleClass {
	return getNENetworkRuleClass()
}

type NENetworkRuleClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NENetworkRuleClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NENetworkRuleClass) Alloc() NENetworkRule {
	rv := objc.Send[NENetworkRule](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A rule to match attributes of network traffic.
//
// # Matching network traffic characteristics
//
//   - [NENetworkRule.MatchRemoteEndpoint]: The remote endpoint that the rule matches.
//   - [NENetworkRule.MatchRemotePrefix]: A number that specifies the remote sub-network that the rule matches.
//   - [NENetworkRule.MatchLocalNetwork]: The local network that the rule matches.
//   - [NENetworkRule.MatchLocalPrefix]: A number that specifies the local sub-network that the rule matches.
//   - [NENetworkRule.MatchProtocol]: The protocol that the rule matches.
//   - [NENetworkRule.MatchDirection]: The direction of network traffic that the rule matches.
//
// # Instance Properties
//
//   - [NENetworkRule.MatchLocalNetworkEndpoint]
//   - [NENetworkRule.SetMatchLocalNetworkEndpoint]
//   - [NENetworkRule.MatchRemoteHostOrNetworkEndpoint]
//   - [NENetworkRule.SetMatchRemoteHostOrNetworkEndpoint]
//
// See: https://developer.apple.com/documentation/NetworkExtension/NENetworkRule
type NENetworkRule struct {
	objectivec.Object
}

// NENetworkRuleFromID constructs a [NENetworkRule] from an objc.ID.
//
// A rule to match attributes of network traffic.
func NENetworkRuleFromID(id objc.ID) NENetworkRule {
	return NENetworkRule{objectivec.Object{ID: id}}
}
// NOTE: NENetworkRule adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NENetworkRule] class.
//
// # Matching network traffic characteristics
//
//   - [INENetworkRule.MatchRemoteEndpoint]: The remote endpoint that the rule matches.
//   - [INENetworkRule.MatchRemotePrefix]: A number that specifies the remote sub-network that the rule matches.
//   - [INENetworkRule.MatchLocalNetwork]: The local network that the rule matches.
//   - [INENetworkRule.MatchLocalPrefix]: A number that specifies the local sub-network that the rule matches.
//   - [INENetworkRule.MatchProtocol]: The protocol that the rule matches.
//   - [INENetworkRule.MatchDirection]: The direction of network traffic that the rule matches.
//
// # Instance Properties
//
//   - [INENetworkRule.MatchLocalNetworkEndpoint]
//   - [INENetworkRule.SetMatchLocalNetworkEndpoint]
//   - [INENetworkRule.MatchRemoteHostOrNetworkEndpoint]
//   - [INENetworkRule.SetMatchRemoteHostOrNetworkEndpoint]
//
// See: https://developer.apple.com/documentation/NetworkExtension/NENetworkRule
type INENetworkRule interface {
	objectivec.IObject

	// Topic: Matching network traffic characteristics

	// The remote endpoint that the rule matches.
	MatchRemoteEndpoint() INWHostEndpoint
	// A number that specifies the remote sub-network that the rule matches.
	MatchRemotePrefix() uint
	// The local network that the rule matches.
	MatchLocalNetwork() INWHostEndpoint
	// A number that specifies the local sub-network that the rule matches.
	MatchLocalPrefix() uint
	// The protocol that the rule matches.
	MatchProtocol() NENetworkRuleProtocol
	// The direction of network traffic that the rule matches.
	MatchDirection() NETrafficDirection

	// Topic: Instance Properties

	MatchLocalNetworkEndpoint() INWEndpoint
	SetMatchLocalNetworkEndpoint(value INWEndpoint)
	MatchRemoteHostOrNetworkEndpoint() INWEndpoint
	SetMatchRemoteHostOrNetworkEndpoint(value INWEndpoint)

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (n NENetworkRule) Init() NENetworkRule {
	rv := objc.Send[NENetworkRule](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n NENetworkRule) Autorelease() NENetworkRule {
	rv := objc.Send[NENetworkRule](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewNENetworkRule creates a new NENetworkRule instance.
func NewNENetworkRule() NENetworkRule {
	class := getNENetworkRuleClass()
	rv := objc.Send[NENetworkRule](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a rule that matches network traffic destined for a host within a
// specific DNS domain.
//
// hostEndpoint: An endpoint instance that contains the port and hostname or domain that the
// rule matches. This endpoint must contain a hostname, not an address.
//
// protocol: The protocol that the rule matches.
//
// # Discussion
// 
// If the port string of `destinationHost` is `0` or is the empty string, then
// the rule matches traffic on any port destined for the given hostname or
// domain.
// 
// If the hostname string of `destinationHost` consists of a single label,
// then the rule matches traffic destined to the specific host with that
// single label as its name.
// 
// If the hostname string of `destinationHost` consists of two or more labels,
// then the rule matches traffic destined to hosts within the domain specified
// by the hostname string.
// 
// # Examples
// 
// The following example makes a rule that matches all TCP and UDP traffic to
// a host named `com` in Swift.
// 
// Here’s the same example in ObjectiveC.
// 
// The next example matches all TCP and UDP traffic to hosts in the
// `example.Com()` DNS domain, including all DNS queries for names in the
// `example.Com()` DNS domain.
// 
// Here’s the same example in ObjectiveC.
// 
// The next example makes a rule that matches all DNS queries and responses
// for hosts in the `example.Com()` domain.
// 
// Here’s the same example in ObjectiveC.
// 
// The last example makes a rule that matches all TCP port 443 traffic to
// hosts in the `example.Com()` domain.
// 
// Here’s the same example in ObjectiveC.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NENetworkRule/init(destinationHost:protocol:)
func NewNetworkRuleWithDestinationHostProtocol(hostEndpoint INWHostEndpoint, protocol_ NENetworkRuleProtocol) NENetworkRule {
	instance := getNENetworkRuleClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDestinationHost:protocol:"), hostEndpoint, protocol_)
	return NENetworkRuleFromID(rv)
}

// Creates a rule that matches network traffic destined for a host within a
// specific network.
//
// networkEndpoint: An endpoint instance that matches the port and address or network that the
// rule matches. This endpoint must contain an address, not a hostname.
//
// destinationPrefix: An integer that in combination with the address in the endpoint specifies
// the destination network that the rule matches.
//
// protocol: The protocol that the rule matches.
//
// # Discussion
// 
// If the port string of `networkEndpoint` is `0` or the empty string, the
// rule matches traffic on any port destined for the given address or network.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NENetworkRule/init(destinationNetwork:prefix:protocol:)
func NewNetworkRuleWithDestinationNetworkPrefixProtocol(networkEndpoint INWHostEndpoint, destinationPrefix uint, protocol_ NENetworkRuleProtocol) NENetworkRule {
	instance := getNENetworkRuleClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDestinationNetwork:prefix:protocol:"), networkEndpoint, destinationPrefix, protocol_)
	return NENetworkRuleFromID(rv)
}

// Creates a rule that matches traffic by remote network, local network,
// protocol, and direction.
//
// remoteNetwork: An endpoint instance that contains the remote port and the remote address
// or network that the rule matches. This endpoint must contain an address,
// not a hostname.
//
// remotePrefix: An integer that in combination with the address in `remoteNetwork`
// specifies the remote network that the rule matches.
//
// localNetwork: An endpoint instance that contains the local port and the local address or
// network that the rule matches. This endpoint must contain an address, not a
// hostname.
//
// localPrefix: An integer that in combination with the address in localNetwork specifies
// the local network that the rule matches. The rule ignores this parameter if
// `localNetwork` is `nil`.
//
// protocol: The protocol that the rule matches.
//
// direction: The direction of network traffic that the rule matches.
//
// # Discussion
// 
// If the port string of `remoteNetwork` is `0` or the empty string, then the
// rule matches traffic on any port coming from the remote network. If
// `remoteNetwork` is `nil`, the rule matches any remote network.
// 
// If the port string of `localNetwork` is `0` or the empty string, then the
// rule matches traffic on any port coming from the local network. If
// `localNetwork` is `nil`, the rule matches any local network.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NENetworkRule/init(remoteNetwork:remotePrefix:localNetwork:localPrefix:protocol:direction:)
func NewNetworkRuleWithRemoteNetworkRemotePrefixLocalNetworkLocalPrefixProtocolDirection(remoteNetwork INWHostEndpoint, remotePrefix uint, localNetwork INWHostEndpoint, localPrefix uint, protocol_ NENetworkRuleProtocol, direction NETrafficDirection) NENetworkRule {
	instance := getNENetworkRuleClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithRemoteNetwork:remotePrefix:localNetwork:localPrefix:protocol:direction:"), remoteNetwork, remotePrefix, localNetwork, localPrefix, protocol_, direction)
	return NENetworkRuleFromID(rv)
}

func (n NENetworkRule) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](n.ID, objc.Sel("encodeWithCoder:"), coder)
}

// The remote endpoint that the rule matches.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NENetworkRule/matchRemoteEndpoint
func (n NENetworkRule) MatchRemoteEndpoint() INWHostEndpoint {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("matchRemoteEndpoint"))
	return NWHostEndpointFromID(objc.ID(rv))
}
// A number that specifies the remote sub-network that the rule matches.
//
// # Discussion
// 
// This property is [NSNotFound] for rules where [MatchRemoteEndpoint]
// doesn’t contain an IP address.
//
// [NSNotFound]: https://developer.apple.com/documentation/Foundation/NSNotFound-4qp9h
//
// See: https://developer.apple.com/documentation/NetworkExtension/NENetworkRule/matchRemotePrefix
func (n NENetworkRule) MatchRemotePrefix() uint {
	rv := objc.Send[uint](n.ID, objc.Sel("matchRemotePrefix"))
	return rv
}
// The local network that the rule matches.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NENetworkRule/matchLocalNetwork
func (n NENetworkRule) MatchLocalNetwork() INWHostEndpoint {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("matchLocalNetwork"))
	return NWHostEndpointFromID(objc.ID(rv))
}
// A number that specifies the local sub-network that the rule matches.
//
// # Discussion
// 
// This property is [NSNotFound] for rules whose [MatchLocalNetwork] property
// is `nil.`
//
// [NSNotFound]: https://developer.apple.com/documentation/Foundation/NSNotFound-4qp9h
//
// See: https://developer.apple.com/documentation/NetworkExtension/NENetworkRule/matchLocalPrefix
func (n NENetworkRule) MatchLocalPrefix() uint {
	rv := objc.Send[uint](n.ID, objc.Sel("matchLocalPrefix"))
	return rv
}
// The protocol that the rule matches.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NENetworkRule/matchProtocol
func (n NENetworkRule) MatchProtocol() NENetworkRuleProtocol {
	rv := objc.Send[NENetworkRuleProtocol](n.ID, objc.Sel("matchProtocol"))
	return NENetworkRuleProtocol(rv)
}
// The direction of network traffic that the rule matches.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NENetworkRule/matchDirection
func (n NENetworkRule) MatchDirection() NETrafficDirection {
	rv := objc.Send[NETrafficDirection](n.ID, objc.Sel("matchDirection"))
	return NETrafficDirection(rv)
}
// See: https://developer.apple.com/documentation/networkextension/nenetworkrule/matchlocalnetworkendpoint-62ttv
func (n NENetworkRule) MatchLocalNetworkEndpoint() INWEndpoint {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("matchLocalNetworkEndpoint"))
	return NWEndpointFromID(objc.ID(rv))
}
func (n NENetworkRule) SetMatchLocalNetworkEndpoint(value INWEndpoint) {
	objc.Send[struct{}](n.ID, objc.Sel("setMatchLocalNetworkEndpoint:"), value)
}
// See: https://developer.apple.com/documentation/networkextension/nenetworkrule/matchremotehostornetworkendpoint-4a5ht
func (n NENetworkRule) MatchRemoteHostOrNetworkEndpoint() INWEndpoint {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("matchRemoteHostOrNetworkEndpoint"))
	return NWEndpointFromID(objc.ID(rv))
}
func (n NENetworkRule) SetMatchRemoteHostOrNetworkEndpoint(value INWEndpoint) {
	objc.Send[struct{}](n.ID, objc.Sel("setMatchRemoteHostOrNetworkEndpoint:"), value)
}

