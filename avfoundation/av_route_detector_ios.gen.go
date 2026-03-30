// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.
//go:build ios
// +build ios

package avfoundation

import (
	"github.com/tmc/apple/objc"
)

// A Boolean value that indicates whether route detection includes custom
// routes.
//
// # Discussion
//
// The default value is false. Only set it to true if your app uses an
// instance of [AVCustomRoutingController].
//
// See: https://developer.apple.com/documentation/AVFoundation/AVRouteDetector/detectsCustomRoutes
//
// [AVCustomRoutingController]: https://developer.apple.com/documentation/AVRouting/AVCustomRoutingController
func (r AVRouteDetector) DetectsCustomRoutes() bool {
	rv := objc.Send[bool](r.ID, objc.Sel("detectsCustomRoutes"))
	return rv
}
func (r AVRouteDetector) SetDetectsCustomRoutes(value bool) {
	objc.Send[struct{}](r.ID, objc.Sel("setDetectsCustomRoutes:"), value)
}
