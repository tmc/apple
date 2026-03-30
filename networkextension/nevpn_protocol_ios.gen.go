// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.
//go:build ios
// +build ios

package networkextension

import (
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// See: https://developer.apple.com/documentation/NetworkExtension/NEVPNProtocol/sliceUUID
func (v NEVPNProtocol) SliceUUID() string {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("sliceUUID"))
	return foundation.NSStringFromID(rv).String()
}
func (v NEVPNProtocol) SetSliceUUID(value string) {
	objc.Send[struct{}](v.ID, objc.Sel("setSliceUUID:"), objc.String(value))
}
