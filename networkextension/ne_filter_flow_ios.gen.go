// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.
//go:build ios
// +build ios

package networkextension

import (
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// A byte string that uniquely identifies the binary for each build of the app
// that is the source of the flow.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEFilterFlow/sourceAppUniqueIdentifier
func (f NEFilterFlow) SourceAppUniqueIdentifier() foundation.INSData {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("sourceAppUniqueIdentifier"))
	return foundation.NSDataFromID(rv)
}

// A string containing the identifier of the source app of the flow.
//
// # Discussion
//
// This identifier remains the same for all versions and builds of the app and
// is unique among all apps.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEFilterFlow/sourceAppIdentifier
func (f NEFilterFlow) SourceAppIdentifier() string {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("sourceAppIdentifier"))
	return foundation.NSStringFromID(rv).String()
}

// The short version string of the app that is the source of the flow.
//
// # Discussion
//
// This property is `nil` if the app info is unavailable.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEFilterFlow/sourceAppVersion
func (f NEFilterFlow) SourceAppVersion() string {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("sourceAppVersion"))
	return foundation.NSStringFromID(rv).String()
}
