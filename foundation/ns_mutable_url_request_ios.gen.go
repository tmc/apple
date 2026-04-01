// Code generated from Apple documentation for Foundation. DO NOT EDIT.
//go:build ios
// +build ios

package foundation

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// Binds a URL request to the network interface associated with the hotspot
// helper command instance.
//
// command: The hotspot helper command to bind the request to.
//
// command is a [*networkextension.NEHotspotHelperCommand].
//
// # Discussion
//
// Apps that participate in joining Wi-Fi hotspot networks use the APIs in the
// [Network Extension] framework to authenticate with hotspots. Ordinarily,
// [NSURLSession] will use the default interface, which may be WWAN. By
// binding to a hotspot helper command, you force a request to use Wi-Fi to
// communicate with the hotspot.
//
// See: https://developer.apple.com/documentation/Foundation/NSMutableURLRequest/bind(to:)
//
// [Network Extension]: https://developer.apple.com/documentation/NetworkExtension
func (m NSMutableURLRequest) BindToHotspotHelperCommand(command objectivec.IObject) {
	objc.Send[objc.ID](m.ID, objc.Sel("bindToHotspotHelperCommand:"), command)
}
