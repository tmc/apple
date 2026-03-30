// Code generated from Apple documentation for AppKit. DO NOT EDIT.
//go:build ios
// +build ios

package appkit

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The custom object from an app built with Mac Catalyst that provides the
// items to share.
//
// # Discussion
//
// Use this object to provide the set of items to share from your macOS
// window.
//
// If this property is `nil`, the item uses the activity items configuration
// the window scene provides. For more information, see
// [activityItemsConfigurationSource].
//
// See: https://developer.apple.com/documentation/AppKit/NSSharingServicePickerToolbarItem/activityItemsConfiguration
//
// [activityItemsConfigurationSource]: https://developer.apple.com/documentation/UIKit/UIWindowScene/activityItemsConfigurationSource
func (s NSSharingServicePickerToolbarItem) ActivityItemsConfiguration() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("activityItemsConfiguration"))
	return objectivec.Object{ID: rv}
}
func (s NSSharingServicePickerToolbarItem) SetActivityItemsConfiguration(value objectivec.IObject) {
	objc.Send[struct{}](s.ID, objc.Sel("setActivityItemsConfiguration:"), value)
}
