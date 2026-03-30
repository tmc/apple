// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NERelay] class.
var (
	_NERelayClass     NERelayClass
	_NERelayClassOnce sync.Once
)

func getNERelayClass() NERelayClass {
	_NERelayClassOnce.Do(func() {
		_NERelayClass = NERelayClass{class: objc.GetClass("NERelay")}
	})
	return _NERelayClass
}

// GetNERelayClass returns the class object for NERelay.
func GetNERelayClass() NERelayClass {
	return getNERelayClass()
}

type NERelayClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NERelayClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NERelayClass) Alloc() NERelay {
	rv := objc.Send[NERelay](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A single relay server configuration that you can chain together with other
// relays.
//
// # Overview
//
// Relay servers are secure HTTP proxies that allow proxying TCP traffic using
// the [CONNECT] method and UDP traffic using the `connect-udp` protocol
// defined in [RFC 9298].
//
// # Configuring server properties
//
//   - [NERelay.HTTP3RelayURL]: A URL identifying the relay server accessible using HTTP/3.
//   - [NERelay.SetHTTP3RelayURL]
//   - [NERelay.HTTP2RelayURL]: A URL identifying the relay server accessible using HTTP/2.
//   - [NERelay.SetHTTP2RelayURL]
//   - [NERelay.DnsOverHTTPSURL]: The URL of a DNS-over-HTTPS (DoH) resolver accessible from the relay.
//   - [NERelay.SetDnsOverHTTPSURL]
//   - [NERelay.RawPublicKeys]: An array of TLS raw public keys that the relay server can present during the TLS handshake.
//   - [NERelay.SetRawPublicKeys]
//
// # Configuring client properties
//
//   - [NERelay.AdditionalHTTPHeaderFields]: A dictionary of additional HTTP headers to send as part of [CONNECT] requests to the relay.
//   - [NERelay.SetAdditionalHTTPHeaderFields]
//   - [NERelay.IdentityData]: The PKCS12 data for the relay client authentication.
//   - [NERelay.SetIdentityData]
//   - [NERelay.IdentityDataPassword]: The password the relay uses to decrypt the PKCS12 identity data.
//   - [NERelay.SetIdentityDataPassword]
//   - [NERelay.SyntheticDNSAnswerIPv4Prefix]: An IPv4 address prefix the relay uses to handle address info requests.
//   - [NERelay.SetSyntheticDNSAnswerIPv4Prefix]
//   - [NERelay.SyntheticDNSAnswerIPv6Prefix]: An IPv6 address prefix the relay uses to handle address info requests.
//   - [NERelay.SetSyntheticDNSAnswerIPv6Prefix]
//
// See: https://developer.apple.com/documentation/NetworkExtension/NERelay
//
// [RFC 9298]: https://www.rfc-editor.org/rfc/rfc9298.html
type NERelay struct {
	objectivec.Object
}

// NERelayFromID constructs a [NERelay] from an objc.ID.
//
// A single relay server configuration that you can chain together with other
// relays.
func NERelayFromID(id objc.ID) NERelay {
	return NERelay{objectivec.Object{ID: id}}
}

// NOTE: NERelay adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NERelay] class.
//
// # Configuring server properties
//
//   - [INERelay.HTTP3RelayURL]: A URL identifying the relay server accessible using HTTP/3.
//   - [INERelay.SetHTTP3RelayURL]
//   - [INERelay.HTTP2RelayURL]: A URL identifying the relay server accessible using HTTP/2.
//   - [INERelay.SetHTTP2RelayURL]
//   - [INERelay.DnsOverHTTPSURL]: The URL of a DNS-over-HTTPS (DoH) resolver accessible from the relay.
//   - [INERelay.SetDnsOverHTTPSURL]
//   - [INERelay.RawPublicKeys]: An array of TLS raw public keys that the relay server can present during the TLS handshake.
//   - [INERelay.SetRawPublicKeys]
//
// # Configuring client properties
//
//   - [INERelay.AdditionalHTTPHeaderFields]: A dictionary of additional HTTP headers to send as part of [CONNECT] requests to the relay.
//   - [INERelay.SetAdditionalHTTPHeaderFields]
//   - [INERelay.IdentityData]: The PKCS12 data for the relay client authentication.
//   - [INERelay.SetIdentityData]
//   - [INERelay.IdentityDataPassword]: The password the relay uses to decrypt the PKCS12 identity data.
//   - [INERelay.SetIdentityDataPassword]
//   - [INERelay.SyntheticDNSAnswerIPv4Prefix]: An IPv4 address prefix the relay uses to handle address info requests.
//   - [INERelay.SetSyntheticDNSAnswerIPv4Prefix]
//   - [INERelay.SyntheticDNSAnswerIPv6Prefix]: An IPv6 address prefix the relay uses to handle address info requests.
//   - [INERelay.SetSyntheticDNSAnswerIPv6Prefix]
//
// See: https://developer.apple.com/documentation/NetworkExtension/NERelay
type INERelay interface {
	objectivec.IObject

	// Topic: Configuring server properties

	// A URL identifying the relay server accessible using HTTP/3.
	HTTP3RelayURL() foundation.INSURL
	SetHTTP3RelayURL(value foundation.INSURL)
	// A URL identifying the relay server accessible using HTTP/2.
	HTTP2RelayURL() foundation.INSURL
	SetHTTP2RelayURL(value foundation.INSURL)
	// The URL of a DNS-over-HTTPS (DoH) resolver accessible from the relay.
	DnsOverHTTPSURL() foundation.INSURL
	SetDnsOverHTTPSURL(value foundation.INSURL)
	// An array of TLS raw public keys that the relay server can present during the TLS handshake.
	RawPublicKeys() []foundation.NSData
	SetRawPublicKeys(value []foundation.NSData)

	// Topic: Configuring client properties

	// A dictionary of additional HTTP headers to send as part of [CONNECT] requests to the relay.
	AdditionalHTTPHeaderFields() foundation.INSDictionary
	SetAdditionalHTTPHeaderFields(value foundation.INSDictionary)
	// The PKCS12 data for the relay client authentication.
	IdentityData() foundation.INSData
	SetIdentityData(value foundation.INSData)
	// The password the relay uses to decrypt the PKCS12 identity data.
	IdentityDataPassword() string
	SetIdentityDataPassword(value string)
	// An IPv4 address prefix the relay uses to handle address info requests.
	SyntheticDNSAnswerIPv4Prefix() string
	SetSyntheticDNSAnswerIPv4Prefix(value string)
	// An IPv6 address prefix the relay uses to handle address info requests.
	SyntheticDNSAnswerIPv6Prefix() string
	SetSyntheticDNSAnswerIPv6Prefix(value string)

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (r NERelay) Init() NERelay {
	rv := objc.Send[NERelay](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r NERelay) Autorelease() NERelay {
	rv := objc.Send[NERelay](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewNERelay creates a new NERelay instance.
func NewNERelay() NERelay {
	class := getNERelayClass()
	rv := objc.Send[NERelay](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (r NERelay) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](r.ID, objc.Sel("encodeWithCoder:"), coder)
}

// A URL identifying the relay server accessible using HTTP/3.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NERelay/http3RelayURL
func (r NERelay) HTTP3RelayURL() foundation.INSURL {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("HTTP3RelayURL"))
	return foundation.NSURLFromID(objc.ID(rv))
}
func (r NERelay) SetHTTP3RelayURL(value foundation.INSURL) {
	objc.Send[struct{}](r.ID, objc.Sel("setHTTP3RelayURL:"), value)
}

// A URL identifying the relay server accessible using HTTP/2.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NERelay/http2RelayURL
func (r NERelay) HTTP2RelayURL() foundation.INSURL {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("HTTP2RelayURL"))
	return foundation.NSURLFromID(objc.ID(rv))
}
func (r NERelay) SetHTTP2RelayURL(value foundation.INSURL) {
	objc.Send[struct{}](r.ID, objc.Sel("setHTTP2RelayURL:"), value)
}

// The URL of a DNS-over-HTTPS (DoH) resolver accessible from the relay.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NERelay/dnsOverHTTPSURL
func (r NERelay) DnsOverHTTPSURL() foundation.INSURL {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("dnsOverHTTPSURL"))
	return foundation.NSURLFromID(objc.ID(rv))
}
func (r NERelay) SetDnsOverHTTPSURL(value foundation.INSURL) {
	objc.Send[struct{}](r.ID, objc.Sel("setDnsOverHTTPSURL:"), value)
}

// An array of TLS raw public keys that the relay server can present during
// the TLS handshake.
//
// # Discussion
//
// If you set one or more keys, the raw public keys are used to authenticate
// the relay server. If no keys are set, or if the array is `nil`, default TLS
// server certificate evaluation is used.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NERelay/rawPublicKeys
func (r NERelay) RawPublicKeys() []foundation.NSData {
	rv := objc.Send[[]objc.ID](r.ID, objc.Sel("rawPublicKeys"))
	return objc.ConvertSlice(rv, func(id objc.ID) foundation.NSData {
		return foundation.NSDataFromID(id)
	})
}
func (r NERelay) SetRawPublicKeys(value []foundation.NSData) {
	objc.Send[struct{}](r.ID, objc.Sel("setRawPublicKeys:"), objectivec.IObjectSliceToNSArray(value))
}

// A dictionary of additional HTTP headers to send as part of [CONNECT]
// requests to the relay.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NERelay/additionalHTTPHeaderFields
func (r NERelay) AdditionalHTTPHeaderFields() foundation.INSDictionary {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("additionalHTTPHeaderFields"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (r NERelay) SetAdditionalHTTPHeaderFields(value foundation.INSDictionary) {
	objc.Send[struct{}](r.ID, objc.Sel("setAdditionalHTTPHeaderFields:"), value)
}

// The PKCS12 data for the relay client authentication.
//
// # Discussion
//
// The value is a [NSData] object in PKCS12 format.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NERelay/identityData
//
// [NSData]: https://developer.apple.com/documentation/Foundation/NSData
func (r NERelay) IdentityData() foundation.INSData {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("identityData"))
	return foundation.NSDataFromID(objc.ID(rv))
}
func (r NERelay) SetIdentityData(value foundation.INSData) {
	objc.Send[struct{}](r.ID, objc.Sel("setIdentityData:"), value)
}

// The password the relay uses to decrypt the PKCS12 identity data.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NERelay/identityDataPassword
func (r NERelay) IdentityDataPassword() string {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("identityDataPassword"))
	return foundation.NSStringFromID(rv).String()
}
func (r NERelay) SetIdentityDataPassword(value string) {
	objc.Send[struct{}](r.ID, objc.Sel("setIdentityDataPassword:"), objc.String(value))
}

// An IPv4 address prefix the relay uses to handle address info requests.
//
// # Discussion
//
// The value of this property is an address prefix, such as `192.0.2.0/24`.
// The relay manager uses this prefix to synthesize DNS answers for apps that
// use `getaddrinfo()` to resolve domains included in [MatchDomains].
//
// See: https://developer.apple.com/documentation/NetworkExtension/NERelay/syntheticDNSAnswerIPv4Prefix
func (r NERelay) SyntheticDNSAnswerIPv4Prefix() string {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("syntheticDNSAnswerIPv4Prefix"))
	return foundation.NSStringFromID(rv).String()
}
func (r NERelay) SetSyntheticDNSAnswerIPv4Prefix(value string) {
	objc.Send[struct{}](r.ID, objc.Sel("setSyntheticDNSAnswerIPv4Prefix:"), objc.String(value))
}

// An IPv6 address prefix the relay uses to handle address info requests.
//
// # Discussion
//
// The value of this property is an address prefix, such as `2001:DB8::/32`.
// The relay manager uses this prefix to synthesize DNS answers for apps that
// use `getaddrinfo()` to resolve domains included in [MatchDomains].
//
// See: https://developer.apple.com/documentation/NetworkExtension/NERelay/syntheticDNSAnswerIPv6Prefix
func (r NERelay) SyntheticDNSAnswerIPv6Prefix() string {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("syntheticDNSAnswerIPv6Prefix"))
	return foundation.NSStringFromID(rv).String()
}
func (r NERelay) SetSyntheticDNSAnswerIPv6Prefix(value string) {
	objc.Send[struct{}](r.ID, objc.Sel("setSyntheticDNSAnswerIPv6Prefix:"), objc.String(value))
}
