// Code generated from Apple documentation for AppKit. DO NOT EDIT.
// +build ios

package appkit

import (
"github.com/tmc/apple/objc"
"github.com/tmc/apple/objectivec"
)

// See: https://developer.apple.com/documentation/AppKit/NSSharingServicePickerTouchBarItem/activityItemsConfiguration
func (s NSSharingServicePickerTouchBarItem) ActivityItemsConfiguration() objectivec.IObject {
rv := objc.Send[objc.ID](s.ID, objc.Sel("activityItemsConfiguration"))
return objectivec.Object{ID: rv}
}
func (s NSSharingServicePickerTouchBarItem) SetActivityItemsConfiguration(value objectivec.IObject) {
objc.Send[struct{}](s.ID, objc.Sel("setActivityItemsConfiguration:"), value)
}








