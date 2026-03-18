// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [URLProtectionSpace] class.
var (
	_URLProtectionSpaceClass     URLProtectionSpaceClass
	_URLProtectionSpaceClassOnce sync.Once
)

func getURLProtectionSpaceClass() URLProtectionSpaceClass {
	_URLProtectionSpaceClassOnce.Do(func() {
		_URLProtectionSpaceClass = URLProtectionSpaceClass{class: objc.GetClass("NSURLProtectionSpace")}
	})
	return _URLProtectionSpaceClass
}

// GetURLProtectionSpaceClass returns the class object for NSURLProtectionSpace.
func GetURLProtectionSpaceClass() URLProtectionSpaceClass {
	return getURLProtectionSpaceClass()
}

type URLProtectionSpaceClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (uc URLProtectionSpaceClass) Alloc() URLProtectionSpace {
	rv := objc.Send[URLProtectionSpace](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// A server or an area on a server, commonly referred to as a realm, that
// requires authentication.
//
// # Overview
// 
// A protection space defines a series of matching constraints that determine
// which credential should be provided. For example, if a request provides
// your delegate with a [NSURLAuthenticationChallenge] object that requests a
// client username and password, your app should provide the correct username
// and password for the particular host, port, protocol, and realm, as
// specified in the challenge’s protection space.
//
// # Creating a protection space
//
//   - [URLProtectionSpace.InitWithHostPortProtocolRealmAuthenticationMethod]: Creates a protection space object from the given host, port, protocol, realm, and authentication method.
//   - [URLProtectionSpace.InitWithProxyHostPortTypeRealmAuthenticationMethod]: Creates a protection space object representing a proxy server.
//
// # Getting protection space properties
//
//   - [URLProtectionSpace.AuthenticationMethod]: The authentication method used by the receiver.
//   - [URLProtectionSpace.DistinguishedNames]: The acceptable certificate-issuing authorities for client certificate authentication.
//   - [URLProtectionSpace.Host]: The receiver’s host.
//   - [URLProtectionSpace.Port]: The receiver’s port.
//   - [URLProtectionSpace.Protocol]: The receiver’s protocol.
//   - [URLProtectionSpace.ProxyType]: The receiver’s proxy type.
//   - [URLProtectionSpace.Realm]: The receiver’s authentication realm
//   - [URLProtectionSpace.ReceivesCredentialSecurely]: A Boolean value that indicates whether the credentials for the protection space can be sent securely.
//   - [URLProtectionSpace.ServerTrust]: A representation of the server’s SSL transaction state.
//
// See: https://developer.apple.com/documentation/Foundation/URLProtectionSpace
type URLProtectionSpace struct {
	objectivec.Object
}

// URLProtectionSpaceFromID constructs a [URLProtectionSpace] from an objc.ID.
//
// A server or an area on a server, commonly referred to as a realm, that
// requires authentication.
func URLProtectionSpaceFromID(id objc.ID) URLProtectionSpace {
	return NSURLProtectionSpace{objectivec.Object{ID: id}}
}

// NSURLProtectionSpaceFromID is an alias for [URLProtectionSpaceFromID] for cross-framework compatibility.
func NSURLProtectionSpaceFromID(id objc.ID) URLProtectionSpace { return URLProtectionSpaceFromID(id) }
// NOTE: URLProtectionSpace adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [URLProtectionSpace] class.
//
// # Creating a protection space
//
//   - [IURLProtectionSpace.InitWithHostPortProtocolRealmAuthenticationMethod]: Creates a protection space object from the given host, port, protocol, realm, and authentication method.
//   - [IURLProtectionSpace.InitWithProxyHostPortTypeRealmAuthenticationMethod]: Creates a protection space object representing a proxy server.
//
// # Getting protection space properties
//
//   - [IURLProtectionSpace.AuthenticationMethod]: The authentication method used by the receiver.
//   - [IURLProtectionSpace.DistinguishedNames]: The acceptable certificate-issuing authorities for client certificate authentication.
//   - [IURLProtectionSpace.Host]: The receiver’s host.
//   - [IURLProtectionSpace.Port]: The receiver’s port.
//   - [IURLProtectionSpace.Protocol]: The receiver’s protocol.
//   - [IURLProtectionSpace.ProxyType]: The receiver’s proxy type.
//   - [IURLProtectionSpace.Realm]: The receiver’s authentication realm
//   - [IURLProtectionSpace.ReceivesCredentialSecurely]: A Boolean value that indicates whether the credentials for the protection space can be sent securely.
//   - [IURLProtectionSpace.ServerTrust]: A representation of the server’s SSL transaction state.
//
// See: https://developer.apple.com/documentation/Foundation/URLProtectionSpace
type IURLProtectionSpace interface {
	objectivec.IObject
	NSCoding
	NSCopying
	NSSecureCoding

	// Topic: Creating a protection space

	// Creates a protection space object from the given host, port, protocol, realm, and authentication method.
	InitWithHostPortProtocolRealmAuthenticationMethod(host string, port int, protocol_ string, realm string, authenticationMethod string) URLProtectionSpace
	// Creates a protection space object representing a proxy server.
	InitWithProxyHostPortTypeRealmAuthenticationMethod(host string, port int, type_ string, realm string, authenticationMethod string) URLProtectionSpace

	// Topic: Getting protection space properties

	// The authentication method used by the receiver.
	AuthenticationMethod() string
	// The acceptable certificate-issuing authorities for client certificate authentication.
	DistinguishedNames() []NSData
	// The receiver’s host.
	Host() string
	// The receiver’s port.
	Port() int
	// The receiver’s protocol.
	Protocol() string
	// The receiver’s proxy type.
	ProxyType() string
	// The receiver’s authentication realm
	Realm() string
	// A Boolean value that indicates whether the credentials for the protection space can be sent securely.
	ReceivesCredentialSecurely() bool
	// A representation of the server’s SSL transaction state.
	ServerTrust() objectivec.IObject

	// A Boolean value that indicates whether the receiver represents a proxy server.
	IsProxy() bool
}

// Init initializes the instance.
func (u URLProtectionSpace) Init() URLProtectionSpace {
	rv := objc.Send[URLProtectionSpace](u.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u URLProtectionSpace) Autorelease() URLProtectionSpace {
	rv := objc.Send[URLProtectionSpace](u.ID, objc.Sel("autorelease"))
	return rv
}

// NewURLProtectionSpace creates a new URLProtectionSpace instance.
func NewURLProtectionSpace() URLProtectionSpace {
	class := getURLProtectionSpaceClass()
	rv := objc.Send[URLProtectionSpace](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/init(coder:)
func NewURLProtectionSpaceWithCoder(coder INSCoder) URLProtectionSpace {
	instance := getURLProtectionSpaceClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return URLProtectionSpaceFromID(rv)
}

// Creates a protection space object from the given host, port, protocol,
// realm, and authentication method.
//
// host: The host name for the [NSURLProtectionSpace] object.
//
// port: The port for the protection space object. If `port` is 0, the default port
// for the specified protocol is used, for example, port 80 for HTTP. Note
// that servers can, and do, treat these values differently.
//
// protocol: The protocol for the protection space object. The value of `protocol` is
// equivalent to the scheme for a URL in the protection space, for example,
// “http”, “https”, “ftp”, etc.
//
// realm: A string indicating a protocol-specific subdivision of the host. `realm`
// may be `nil` if there is no specified realm or if the protocol doesn’t
// support realms.
//
// authenticationMethod: The type of authentication to use. `authenticationMethod` should be set to
// one of the values in [NSURLProtectionSpace authentication method constants]
// or `nil` to use the default, [NSURLAuthenticationMethodDefault].
// //
// [NSURLAuthenticationMethodDefault]: https://developer.apple.com/documentation/Foundation/NSURLAuthenticationMethodDefault
// [NSURLProtectionSpace authentication method constants]: https://developer.apple.com/documentation/Foundation/nsurlprotectionspace-authentication-method-constants
//
// # Return Value
// 
// A new protection space object, initialized with the given host, port,
// protocol, realm, and authentication method.
//
// See: https://developer.apple.com/documentation/Foundation/URLProtectionSpace/init(host:port:protocol:realm:authenticationMethod:)
func NewURLProtectionSpaceWithHostPortProtocolRealmAuthenticationMethod(host string, port int, protocol_ string, realm string, authenticationMethod string) URLProtectionSpace {
	instance := getURLProtectionSpaceClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithHost:port:protocol:realm:authenticationMethod:"), objc.String(host), port, objc.String(protocol_), objc.String(realm), objc.String(authenticationMethod))
	return URLProtectionSpaceFromID(rv)
}

// Creates a protection space object representing a proxy server.
//
// host: The host of the proxy server for the protection space object.
//
// port: The port for the protection space object. If `port` is 0 the default port
// for the specified proxy type is used, for example, port 80 for HTTP. Note
// that servers can, and do, treat these values differently.
//
// type: The type of proxy server. The value of `proxyType` should be set to one of
// the values specified in [NSURLProtectionSpace proxy types].
// //
// [NSURLProtectionSpace proxy types]: https://developer.apple.com/documentation/Foundation/nsurlprotectionspace-proxy-types
//
// realm: A string indicating a protocol specific subdivision of the host. `realm`
// may be `nil` if there is no specified realm or if the protocol doesn’t
// support realms.
//
// authenticationMethod: The type of authentication to use. `authenticationMethod` should be set to
// one of the values in [NSURLProtectionSpace authentication method constants]
// or `nil` to use the default, [NSURLAuthenticationMethodDefault].
// //
// [NSURLAuthenticationMethodDefault]: https://developer.apple.com/documentation/Foundation/NSURLAuthenticationMethodDefault
// [NSURLProtectionSpace authentication method constants]: https://developer.apple.com/documentation/Foundation/nsurlprotectionspace-authentication-method-constants
//
// # Return Value
// 
// A new protection space object, with the given host, port, proxyType, realm,
// and authenticationMethod.
//
// See: https://developer.apple.com/documentation/Foundation/URLProtectionSpace/init(proxyHost:port:type:realm:authenticationMethod:)
func NewURLProtectionSpaceWithProxyHostPortTypeRealmAuthenticationMethod(host string, port int, type_ string, realm string, authenticationMethod string) URLProtectionSpace {
	instance := getURLProtectionSpaceClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithProxyHost:port:type:realm:authenticationMethod:"), objc.String(host), port, objc.String(type_), objc.String(realm), objc.String(authenticationMethod))
	return URLProtectionSpaceFromID(rv)
}

// Creates a protection space object from the given host, port, protocol,
// realm, and authentication method.
//
// host: The host name for the [NSURLProtectionSpace] object.
//
// port: The port for the protection space object. If `port` is 0, the default port
// for the specified protocol is used, for example, port 80 for HTTP. Note
// that servers can, and do, treat these values differently.
//
// protocol: The protocol for the protection space object. The value of `protocol` is
// equivalent to the scheme for a URL in the protection space, for example,
// “http”, “https”, “ftp”, etc.
//
// realm: A string indicating a protocol-specific subdivision of the host. `realm`
// may be `nil` if there is no specified realm or if the protocol doesn’t
// support realms.
//
// authenticationMethod: The type of authentication to use. `authenticationMethod` should be set to
// one of the values in [NSURLProtectionSpace authentication method constants]
// or `nil` to use the default, [NSURLAuthenticationMethodDefault].
// //
// [NSURLAuthenticationMethodDefault]: https://developer.apple.com/documentation/Foundation/NSURLAuthenticationMethodDefault
// [NSURLProtectionSpace authentication method constants]: https://developer.apple.com/documentation/Foundation/nsurlprotectionspace-authentication-method-constants
//
// # Return Value
// 
// A new protection space object, initialized with the given host, port,
// protocol, realm, and authentication method.
//
// See: https://developer.apple.com/documentation/Foundation/URLProtectionSpace/init(host:port:protocol:realm:authenticationMethod:)
func (u URLProtectionSpace) InitWithHostPortProtocolRealmAuthenticationMethod(host string, port int, protocol_ string, realm string, authenticationMethod string) URLProtectionSpace {
	rv := objc.Send[URLProtectionSpace](u.ID, objc.Sel("initWithHost:port:protocol:realm:authenticationMethod:"), objc.String(host), port, objc.String(protocol_), objc.String(realm), objc.String(authenticationMethod))
	return rv
}

// Creates a protection space object representing a proxy server.
//
// host: The host of the proxy server for the protection space object.
//
// port: The port for the protection space object. If `port` is 0 the default port
// for the specified proxy type is used, for example, port 80 for HTTP. Note
// that servers can, and do, treat these values differently.
//
// type: The type of proxy server. The value of `proxyType` should be set to one of
// the values specified in [NSURLProtectionSpace proxy types].
// //
// [NSURLProtectionSpace proxy types]: https://developer.apple.com/documentation/Foundation/nsurlprotectionspace-proxy-types
//
// realm: A string indicating a protocol specific subdivision of the host. `realm`
// may be `nil` if there is no specified realm or if the protocol doesn’t
// support realms.
//
// authenticationMethod: The type of authentication to use. `authenticationMethod` should be set to
// one of the values in [NSURLProtectionSpace authentication method constants]
// or `nil` to use the default, [NSURLAuthenticationMethodDefault].
// //
// [NSURLAuthenticationMethodDefault]: https://developer.apple.com/documentation/Foundation/NSURLAuthenticationMethodDefault
// [NSURLProtectionSpace authentication method constants]: https://developer.apple.com/documentation/Foundation/nsurlprotectionspace-authentication-method-constants
//
// # Return Value
// 
// A new protection space object, with the given host, port, proxyType, realm,
// and authenticationMethod.
//
// See: https://developer.apple.com/documentation/Foundation/URLProtectionSpace/init(proxyHost:port:type:realm:authenticationMethod:)
func (u URLProtectionSpace) InitWithProxyHostPortTypeRealmAuthenticationMethod(host string, port int, type_ string, realm string, authenticationMethod string) URLProtectionSpace {
	rv := objc.Send[URLProtectionSpace](u.ID, objc.Sel("initWithProxyHost:port:type:realm:authenticationMethod:"), objc.String(host), port, objc.String(type_), objc.String(realm), objc.String(authenticationMethod))
	return rv
}

// Encodes the receiver using a given archiver.
//
// coder: An archiver object.
//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/encode(with:)
func (u URLProtectionSpace) EncodeWithCoder(coder INSCoder) {
	objc.Send[objc.ID](u.ID, objc.Sel("encodeWithCoder:"), coder)
}

//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/init(coder:)
func (u URLProtectionSpace) InitWithCoder(coder INSCoder) URLProtectionSpace {
	rv := objc.Send[URLProtectionSpace](u.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}

// The authentication method used by the receiver.
//
// # Discussion
// 
// The supported authentication methods are listed in [NSURLProtectionSpace
// authentication method constants].
//
// [NSURLProtectionSpace authentication method constants]: https://developer.apple.com/documentation/Foundation/nsurlprotectionspace-authentication-method-constants
//
// See: https://developer.apple.com/documentation/Foundation/URLProtectionSpace/authenticationMethod
func (u URLProtectionSpace) AuthenticationMethod() string {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("authenticationMethod"))
	return NSStringFromID(rv).String()
}

// The acceptable certificate-issuing authorities for client certificate
// authentication.
//
// # Discussion
// 
// This value is `nil` if the authentication method of the protection space is
// not client certificate. The returned issuing authorities are encoded with
// Distinguished Encoding Rules (DER).
//
// See: https://developer.apple.com/documentation/Foundation/URLProtectionSpace/distinguishedNames
func (u URLProtectionSpace) DistinguishedNames() []NSData {
	rv := objc.Send[[]objc.ID](u.ID, objc.Sel("distinguishedNames"))
	return objc.ConvertSlice(rv, func(id objc.ID) NSData {
		return NSDataFromID(id)
	})
}

// The receiver’s host.
//
// See: https://developer.apple.com/documentation/Foundation/URLProtectionSpace/host
func (u URLProtectionSpace) Host() string {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("host"))
	return NSStringFromID(rv).String()
}

// The receiver’s port.
//
// See: https://developer.apple.com/documentation/Foundation/URLProtectionSpace/port
func (u URLProtectionSpace) Port() int {
	rv := objc.Send[int](u.ID, objc.Sel("port"))
	return rv
}

// The receiver’s protocol.
//
// # Discussion
// 
// This value is `nil` if the receiver represents a proxy protection space.
//
// See: https://developer.apple.com/documentation/Foundation/URLProtectionSpace/protocol
func (u URLProtectionSpace) Protocol() string {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("protocol"))
	return NSStringFromID(rv).String()
}

// The receiver’s proxy type.
//
// # Discussion
// 
// This value is `nil` if the receiver does not represent a proxy protection
// space. The supported proxy types are listed in [NSURLProtectionSpace proxy
// types].
//
// [NSURLProtectionSpace proxy types]: https://developer.apple.com/documentation/Foundation/nsurlprotectionspace-proxy-types
//
// See: https://developer.apple.com/documentation/Foundation/URLProtectionSpace/proxyType
func (u URLProtectionSpace) ProxyType() string {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("proxyType"))
	return NSStringFromID(rv).String()
}

// The receiver’s authentication realm
//
// # Discussion
// 
// This value is `nil` if no realm has been set. A realm is generally only
// specified for HTTP and HTTPS authentication.
//
// See: https://developer.apple.com/documentation/Foundation/URLProtectionSpace/realm
func (u URLProtectionSpace) Realm() string {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("realm"))
	return NSStringFromID(rv).String()
}

// A Boolean value that indicates whether the credentials for the protection
// space can be sent securely.
//
// # Discussion
// 
// This value is [true] if the credentials for the protection space
// represented by the receiver can be sent securely, [false] otherwise.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Foundation/URLProtectionSpace/receivesCredentialSecurely
func (u URLProtectionSpace) ReceivesCredentialSecurely() bool {
	rv := objc.Send[bool](u.ID, objc.Sel("receivesCredentialSecurely"))
	return rv
}

// A representation of the server’s SSL transaction state.
//
// # Discussion
// 
// This value is `nil` if the authentication method of the protection space is
// not server trust.
//
// See: https://developer.apple.com/documentation/Foundation/URLProtectionSpace/serverTrust
func (u URLProtectionSpace) ServerTrust() objectivec.IObject {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("serverTrust"))
	return objectivec.Object{ID: rv}
}

// A Boolean value that indicates whether the receiver represents a proxy
// server.
//
// # Discussion
// 
// [true] if the receiver represents a proxy server, [false] otherwise.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Foundation/NSURLProtectionSpace/isProxy
func (u URLProtectionSpace) IsProxy() bool {
	rv := objc.Send[bool](u.ID, objc.Sel("isProxy"))
	return rv
}

			// Protocol methods for NSCopying
			

			// Protocol methods for NSSecureCoding
			

