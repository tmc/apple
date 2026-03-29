// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NEProxyServer] class.
var (
	_NEProxyServerClass     NEProxyServerClass
	_NEProxyServerClassOnce sync.Once
)

func getNEProxyServerClass() NEProxyServerClass {
	_NEProxyServerClassOnce.Do(func() {
		_NEProxyServerClass = NEProxyServerClass{class: objc.GetClass("NEProxyServer")}
	})
	return _NEProxyServerClass
}

// GetNEProxyServerClass returns the class object for NEProxyServer.
func GetNEProxyServerClass() NEProxyServerClass {
	return getNEProxyServerClass()
}

type NEProxyServerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NEProxyServerClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NEProxyServerClass) Alloc() NEProxyServer {
	rv := objc.Send[NEProxyServer](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// [NEProxyServer] contains settings for a proxy server.
//
// # Overview
// 
// [NEProxyServer] instances are used inside of [NEProxySettings] instances to
// configure proxy settings for VPN connections.
//
// # Initializing a Proxy Server
//
//   - [NEProxyServer.InitWithAddressPort]: Initialize a newly-allocated [NEProxyServer] object
//
// # Accessing Proxy Server Properties
//
//   - [NEProxyServer.Address]: The address of the proxy server.
//   - [NEProxyServer.Port]: The TCP port on which the proxy server is listening for connections.
//   - [NEProxyServer.AuthenticationRequired]: A Boolean indicating if the server requires authentication credentials.
//   - [NEProxyServer.SetAuthenticationRequired]
//   - [NEProxyServer.Username]: The username portion of the authentication credential to be used to authenticate with the proxy server.
//   - [NEProxyServer.SetUsername]
//   - [NEProxyServer.Password]: The password portion of the authentication credential to be used to authenticate with the proxy server.
//   - [NEProxyServer.SetPassword]
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEProxyServer
type NEProxyServer struct {
	objectivec.Object
}

// NEProxyServerFromID constructs a [NEProxyServer] from an objc.ID.
//
// [NEProxyServer] contains settings for a proxy server.
func NEProxyServerFromID(id objc.ID) NEProxyServer {
	return NEProxyServer{objectivec.Object{ID: id}}
}
// NOTE: NEProxyServer adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NEProxyServer] class.
//
// # Initializing a Proxy Server
//
//   - [INEProxyServer.InitWithAddressPort]: Initialize a newly-allocated [NEProxyServer] object
//
// # Accessing Proxy Server Properties
//
//   - [INEProxyServer.Address]: The address of the proxy server.
//   - [INEProxyServer.Port]: The TCP port on which the proxy server is listening for connections.
//   - [INEProxyServer.AuthenticationRequired]: A Boolean indicating if the server requires authentication credentials.
//   - [INEProxyServer.SetAuthenticationRequired]
//   - [INEProxyServer.Username]: The username portion of the authentication credential to be used to authenticate with the proxy server.
//   - [INEProxyServer.SetUsername]
//   - [INEProxyServer.Password]: The password portion of the authentication credential to be used to authenticate with the proxy server.
//   - [INEProxyServer.SetPassword]
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEProxyServer
type INEProxyServer interface {
	objectivec.IObject

	// Topic: Initializing a Proxy Server

	// Initialize a newly-allocated [NEProxyServer] object
	InitWithAddressPort(address string, port int) NEProxyServer

	// Topic: Accessing Proxy Server Properties

	// The address of the proxy server.
	Address() string
	// The TCP port on which the proxy server is listening for connections.
	Port() int
	// A Boolean indicating if the server requires authentication credentials.
	AuthenticationRequired() bool
	SetAuthenticationRequired(value bool)
	// The username portion of the authentication credential to be used to authenticate with the proxy server.
	Username() string
	SetUsername(value string)
	// The password portion of the authentication credential to be used to authenticate with the proxy server.
	Password() string
	SetPassword(value string)

	// A Boolean indicating if a static HTTP proxy will be used.
	HttpEnabled() bool
	SetHttpEnabled(value bool)
	// An 
	HttpServer() INEProxyServer
	SetHttpServer(value INEProxyServer)
	// A Boolean indicating if a static HTTPS proxy will be used.
	HttpsEnabled() bool
	SetHttpsEnabled(value bool)
	// An 
	HttpsServer() INEProxyServer
	SetHttpsServer(value INEProxyServer)
	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (p NEProxyServer) Init() NEProxyServer {
	rv := objc.Send[NEProxyServer](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p NEProxyServer) Autorelease() NEProxyServer {
	rv := objc.Send[NEProxyServer](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEProxyServer creates a new NEProxyServer instance.
func NewNEProxyServer() NEProxyServer {
	class := getNEProxyServerClass()
	rv := objc.Send[NEProxyServer](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Initialize a newly-allocated [NEProxyServer] object
//
// address: The address of the proxy server.
//
// port: The TCP port on which the proxy server is listening for connections.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEProxyServer/init(address:port:)
func NewProxyServerWithAddressPort(address string, port int) NEProxyServer {
	instance := getNEProxyServerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithAddress:port:"), objc.String(address), port)
	return NEProxyServerFromID(rv)
}

// Initialize a newly-allocated [NEProxyServer] object
//
// address: The address of the proxy server.
//
// port: The TCP port on which the proxy server is listening for connections.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEProxyServer/init(address:port:)
func (p NEProxyServer) InitWithAddressPort(address string, port int) NEProxyServer {
	rv := objc.Send[NEProxyServer](p.ID, objc.Sel("initWithAddress:port:"), objc.String(address), port)
	return rv
}
func (p NEProxyServer) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](p.ID, objc.Sel("encodeWithCoder:"), coder)
}

// The address of the proxy server.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEProxyServer/address
func (p NEProxyServer) Address() string {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("address"))
	return foundation.NSStringFromID(rv).String()
}
// The TCP port on which the proxy server is listening for connections.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEProxyServer/port
func (p NEProxyServer) Port() int {
	rv := objc.Send[int](p.ID, objc.Sel("port"))
	return rv
}
// A Boolean indicating if the server requires authentication credentials.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEProxyServer/authenticationRequired
func (p NEProxyServer) AuthenticationRequired() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("authenticationRequired"))
	return rv
}
func (p NEProxyServer) SetAuthenticationRequired(value bool) {
	objc.Send[struct{}](p.ID, objc.Sel("setAuthenticationRequired:"), value)
}
// The username portion of the authentication credential to be used to
// authenticate with the proxy server.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEProxyServer/username
func (p NEProxyServer) Username() string {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("username"))
	return foundation.NSStringFromID(rv).String()
}
func (p NEProxyServer) SetUsername(value string) {
	objc.Send[struct{}](p.ID, objc.Sel("setUsername:"), objc.String(value))
}
// The password portion of the authentication credential to be used to
// authenticate with the proxy server.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEProxyServer/password
func (p NEProxyServer) Password() string {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("password"))
	return foundation.NSStringFromID(rv).String()
}
func (p NEProxyServer) SetPassword(value string) {
	objc.Send[struct{}](p.ID, objc.Sel("setPassword:"), objc.String(value))
}
// A Boolean indicating if a static HTTP proxy will be used.
//
// See: https://developer.apple.com/documentation/networkextension/neproxysettings/httpenabled
func (p NEProxyServer) HttpEnabled() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("HTTPEnabled"))
	return rv
}
func (p NEProxyServer) SetHttpEnabled(value bool) {
	objc.Send[struct{}](p.ID, objc.Sel("setHTTPEnabled:"), value)
}
// An
//
// See: https://developer.apple.com/documentation/networkextension/neproxysettings/httpserver
func (p NEProxyServer) HttpServer() INEProxyServer {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("HTTPServer"))
	return NEProxyServerFromID(objc.ID(rv))
}
func (p NEProxyServer) SetHttpServer(value INEProxyServer) {
	objc.Send[struct{}](p.ID, objc.Sel("setHTTPServer:"), value)
}
// A Boolean indicating if a static HTTPS proxy will be used.
//
// See: https://developer.apple.com/documentation/networkextension/neproxysettings/httpsenabled
func (p NEProxyServer) HttpsEnabled() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("HTTPSEnabled"))
	return rv
}
func (p NEProxyServer) SetHttpsEnabled(value bool) {
	objc.Send[struct{}](p.ID, objc.Sel("setHTTPSEnabled:"), value)
}
// An
//
// See: https://developer.apple.com/documentation/networkextension/neproxysettings/httpsserver
func (p NEProxyServer) HttpsServer() INEProxyServer {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("HTTPSServer"))
	return NEProxyServerFromID(objc.ID(rv))
}
func (p NEProxyServer) SetHttpsServer(value INEProxyServer) {
	objc.Send[struct{}](p.ID, objc.Sel("setHTTPSServer:"), value)
}

