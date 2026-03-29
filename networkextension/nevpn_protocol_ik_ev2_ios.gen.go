// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.
// +build ios

package networkextension

import (
"github.com/tmc/apple/objc"
)

// A property to enable the use of cellular data when Wi-Fi connectivity is
// poor.
//
// # Discussion
// 
// Wi-Fi Assist allows connections for foreground apps to switch over to
// cellular data when Wi-Fi connectivity is poor. When this value is `true`,
// the device brings up a tunnel over celluar to carry traffic that’s
// eligible for Wi-Fi Assist and also requires VPN. Enabling fallback requires
// that the server support multiple tunnels for a single user.
// 
// The default value of this property is `false`.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocolIKEv2/enableFallback
func (v NEVPNProtocolIKEv2) EnableFallback() bool {
rv := objc.Send[bool](v.ID, objc.Sel("enableFallback"))
		return rv
}
func (v NEVPNProtocolIKEv2) SetEnableFallback(value bool) {
objc.Send[struct{}](v.ID, objc.Sel("setEnableFallback:"), value)
}

