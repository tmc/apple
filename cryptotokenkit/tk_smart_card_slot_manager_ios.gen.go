// Code generated from Apple documentation for CryptoTokenKit. DO NOT EDIT.
//go:build ios
// +build ios

package cryptotokenkit

import (
	"github.com/tmc/apple/objc"
)

// Determines whether NFC (Near Field Communication) is supported on this
// device.
//
// # Return Value
//
// [YES] if NFC is supported and available for use, NO otherwise.
//
// See: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardSlotManager/isNFCSupported()
func (t TKSmartCardSlotManager) IsNFCSupported() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("isNFCSupported"))
	return rv
}
