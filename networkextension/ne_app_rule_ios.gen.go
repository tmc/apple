// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.
//go:build ios
// +build ios

package networkextension

import (
	"github.com/tmc/apple/objc"
)

// Create an app rule that matches an app with a given signing identifier.
//
// signingIdentifier: The signing identifier of the app that matches the rule. For apps that are
// signed using Xcode, the app’s signing identifier is equivalent to the
// app’s bundle identifier.
//
// # Return Value
//
// A newly-initialized [NEAppRule] object.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEAppRule/init(signingIdentifier:)
func (a NEAppRule) InitWithSigningIdentifier(signingIdentifier string) NEAppRule {
	rv := objc.Send[NEAppRule](a.ID, objc.Sel("initWithSigningIdentifier:"), objc.String(signingIdentifier))
	return rv
}
