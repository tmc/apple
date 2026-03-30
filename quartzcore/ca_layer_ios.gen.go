// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.
//go:build ios
// +build ios

package quartzcore

import (
	"github.com/tmc/apple/objc"
)

// See: https://developer.apple.com/documentation/QuartzCore/CALayer/wantsDynamicContentScaling
func (l CALayer) WantsDynamicContentScaling() bool {
	rv := objc.Send[bool](l.ID, objc.Sel("wantsDynamicContentScaling"))
	return rv
}
func (l CALayer) SetWantsDynamicContentScaling(value bool) {
	objc.Send[struct{}](l.ID, objc.Sel("setWantsDynamicContentScaling:"), value)
}
