// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [WKSecurityOrigin] class.
var (
	_WKSecurityOriginClass     WKSecurityOriginClass
	_WKSecurityOriginClassOnce sync.Once
)

func getWKSecurityOriginClass() WKSecurityOriginClass {
	_WKSecurityOriginClassOnce.Do(func() {
		_WKSecurityOriginClass = WKSecurityOriginClass{class: objc.GetClass("WKSecurityOrigin")}
	})
	return _WKSecurityOriginClass
}

// GetWKSecurityOriginClass returns the class object for WKSecurityOrigin.
func GetWKSecurityOriginClass() WKSecurityOriginClass {
	return getWKSecurityOriginClass()
}

type WKSecurityOriginClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (wc WKSecurityOriginClass) Class() objc.Class {
	return wc.class
}

// Alloc allocates memory for a new instance of the class.
func (wc WKSecurityOriginClass) Alloc() WKSecurityOrigin {
	rv := objc.Send[WKSecurityOrigin](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

// An object that identifies the origin of a particular resource.
//
// # Overview
//
// A [WKSecurityOrigin] object is a transient, data-only object that
// identifies the host name, protocol, and port number associated with a
// particular resource. You don’t create [WKSecurityOrigin] objects
// directly. Instead, WebKit creates them for the resources it loads. A load
// is any load URL has the same security origin as the requesting web site.
// First-party webpages can access each other’s resources, such as scripts
// and databases.
//
// Because a [WKSecurityOrigin] object is transient, it doesn’t uniquely
// identify a security origin across multiple delegate method calls.
//
// # Getting the Host Information
//
//   - [WKSecurityOrigin.Host]: The security origin’s host.
//   - [WKSecurityOrigin.Port]: The security origin’s port.
//
// # Getting the Host Protocol
//
//   - [WKSecurityOrigin.Protocol]: The security origin’s protocol.
//
// See: https://developer.apple.com/documentation/WebKit/WKSecurityOrigin
type WKSecurityOrigin struct {
	objectivec.Object
}

// WKSecurityOriginFromID constructs a [WKSecurityOrigin] from an objc.ID.
//
// An object that identifies the origin of a particular resource.
func WKSecurityOriginFromID(id objc.ID) WKSecurityOrigin {
	return WKSecurityOrigin{objectivec.Object{ID: id}}
}

// NOTE: WKSecurityOrigin adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [WKSecurityOrigin] class.
//
// # Getting the Host Information
//
//   - [IWKSecurityOrigin.Host]: The security origin’s host.
//   - [IWKSecurityOrigin.Port]: The security origin’s port.
//
// # Getting the Host Protocol
//
//   - [IWKSecurityOrigin.Protocol]: The security origin’s protocol.
//
// See: https://developer.apple.com/documentation/WebKit/WKSecurityOrigin
type IWKSecurityOrigin interface {
	objectivec.IObject

	// Topic: Getting the Host Information

	// The security origin’s host.
	Host() string
	// The security origin’s port.
	Port() int

	// Topic: Getting the Host Protocol

	// The security origin’s protocol.
	Protocol() string
}

// Init initializes the instance.
func (s WKSecurityOrigin) Init() WKSecurityOrigin {
	rv := objc.Send[WKSecurityOrigin](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s WKSecurityOrigin) Autorelease() WKSecurityOrigin {
	rv := objc.Send[WKSecurityOrigin](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewWKSecurityOrigin creates a new WKSecurityOrigin instance.
func NewWKSecurityOrigin() WKSecurityOrigin {
	class := getWKSecurityOriginClass()
	rv := objc.Send[WKSecurityOrigin](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The security origin’s host.
//
// See: https://developer.apple.com/documentation/WebKit/WKSecurityOrigin/host
func (s WKSecurityOrigin) Host() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("host"))
	return foundation.NSStringFromID(rv).String()
}

// The security origin’s port.
//
// See: https://developer.apple.com/documentation/WebKit/WKSecurityOrigin/port
func (s WKSecurityOrigin) Port() int {
	rv := objc.Send[int](s.ID, objc.Sel("port"))
	return rv
}

// The security origin’s protocol.
//
// See: https://developer.apple.com/documentation/WebKit/WKSecurityOrigin/protocol
func (s WKSecurityOrigin) Protocol() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("protocol"))
	return foundation.NSStringFromID(rv).String()
}
