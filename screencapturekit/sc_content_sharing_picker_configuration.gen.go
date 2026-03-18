// Code generated from Apple documentation for ScreenCaptureKit. DO NOT EDIT.

package screencapturekit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SCContentSharingPickerConfiguration] class.
var (
	_SCContentSharingPickerConfigurationClass     SCContentSharingPickerConfigurationClass
	_SCContentSharingPickerConfigurationClassOnce sync.Once
)

func getSCContentSharingPickerConfigurationClass() SCContentSharingPickerConfigurationClass {
	_SCContentSharingPickerConfigurationClassOnce.Do(func() {
		_SCContentSharingPickerConfigurationClass = SCContentSharingPickerConfigurationClass{class: objc.GetClass("SCContentSharingPickerConfiguration")}
	})
	return _SCContentSharingPickerConfigurationClass
}

// GetSCContentSharingPickerConfigurationClass returns the class object for SCContentSharingPickerConfiguration.
func GetSCContentSharingPickerConfigurationClass() SCContentSharingPickerConfigurationClass {
	return getSCContentSharingPickerConfigurationClass()
}

type SCContentSharingPickerConfigurationClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (sc SCContentSharingPickerConfigurationClass) Alloc() SCContentSharingPickerConfiguration {
	rv := objc.Send[SCContentSharingPickerConfiguration](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// An instance for configuring the system content-sharing picker.
//
// # Control streaming selections
//
//   - [SCContentSharingPickerConfiguration.AllowedPickerModes]: The content-selection modes supported by the picker.
//   - [SCContentSharingPickerConfiguration.SetAllowedPickerModes]
//   - [SCContentSharingPickerConfiguration.AllowsChangingSelectedContent]: A Boolean value that indicates if the present stream can change to a different source.
//   - [SCContentSharingPickerConfiguration.SetAllowsChangingSelectedContent]
//   - [SCContentSharingPickerConfiguration.ExcludedBundleIDs]: A list of bundle IDs to exclude from the sharing picker.
//   - [SCContentSharingPickerConfiguration.SetExcludedBundleIDs]
//   - [SCContentSharingPickerConfiguration.ExcludedWindowIDs]: A list of window IDs to exclude from the sharing picker.
//   - [SCContentSharingPickerConfiguration.SetExcludedWindowIDs]
//
// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCContentSharingPickerConfiguration-c.class
type SCContentSharingPickerConfiguration struct {
	objectivec.Object
}

// SCContentSharingPickerConfigurationFromID constructs a [SCContentSharingPickerConfiguration] from an objc.ID.
//
// An instance for configuring the system content-sharing picker.
func SCContentSharingPickerConfigurationFromID(id objc.ID) SCContentSharingPickerConfiguration {
	return SCContentSharingPickerConfiguration{objectivec.Object{ID: id}}
}
// Ensure SCContentSharingPickerConfiguration implements ISCContentSharingPickerConfiguration.
var _ ISCContentSharingPickerConfiguration = SCContentSharingPickerConfiguration{}

// An interface definition for the [SCContentSharingPickerConfiguration] class.
//
// # Control streaming selections
//
//   - [ISCContentSharingPickerConfiguration.AllowedPickerModes]: The content-selection modes supported by the picker.
//   - [ISCContentSharingPickerConfiguration.SetAllowedPickerModes]
//   - [ISCContentSharingPickerConfiguration.AllowsChangingSelectedContent]: A Boolean value that indicates if the present stream can change to a different source.
//   - [ISCContentSharingPickerConfiguration.SetAllowsChangingSelectedContent]
//   - [ISCContentSharingPickerConfiguration.ExcludedBundleIDs]: A list of bundle IDs to exclude from the sharing picker.
//   - [ISCContentSharingPickerConfiguration.SetExcludedBundleIDs]
//   - [ISCContentSharingPickerConfiguration.ExcludedWindowIDs]: A list of window IDs to exclude from the sharing picker.
//   - [ISCContentSharingPickerConfiguration.SetExcludedWindowIDs]
//
// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCContentSharingPickerConfiguration-c.class
type ISCContentSharingPickerConfiguration interface {
	objectivec.IObject

	// Topic: Control streaming selections

	// The content-selection modes supported by the picker.
	AllowedPickerModes() SCContentSharingPickerMode
	SetAllowedPickerModes(value SCContentSharingPickerMode)
	// A Boolean value that indicates if the present stream can change to a different source.
	AllowsChangingSelectedContent() bool
	SetAllowsChangingSelectedContent(value bool)
	// A list of bundle IDs to exclude from the sharing picker.
	ExcludedBundleIDs() []string
	SetExcludedBundleIDs(value []string)
	// A list of window IDs to exclude from the sharing picker.
	ExcludedWindowIDs() []foundation.NSNumber
	SetExcludedWindowIDs(value []foundation.NSNumber)
}

// Init initializes the instance.
func (c SCContentSharingPickerConfiguration) Init() SCContentSharingPickerConfiguration {
	rv := objc.Send[SCContentSharingPickerConfiguration](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c SCContentSharingPickerConfiguration) Autorelease() SCContentSharingPickerConfiguration {
	rv := objc.Send[SCContentSharingPickerConfiguration](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewSCContentSharingPickerConfiguration creates a new SCContentSharingPickerConfiguration instance.
func NewSCContentSharingPickerConfiguration() SCContentSharingPickerConfiguration {
	class := getSCContentSharingPickerConfigurationClass()
	rv := objc.Send[SCContentSharingPickerConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The content-selection modes supported by the picker.
//
// # Discussion
// 
// The default value doesn’t exclude selecting any content for streaming.
// There isn’t an equivalent [SCContentSharingPickerMode] available to reset
// this property once changed.
//
// [SCContentSharingPickerMode]: https://developer.apple.com/documentation/ScreenCaptureKit/SCContentSharingPickerMode
//
// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCContentSharingPickerConfiguration-c.class/allowedPickerModes
func (c SCContentSharingPickerConfiguration) AllowedPickerModes() SCContentSharingPickerMode {
	rv := objc.Send[SCContentSharingPickerMode](c.ID, objc.Sel("allowedPickerModes"))
	return SCContentSharingPickerMode(rv)
}
func (c SCContentSharingPickerConfiguration) SetAllowedPickerModes(value SCContentSharingPickerMode) {
	objc.Send[struct{}](c.ID, objc.Sel("setAllowedPickerModes:"), value)
}

// A Boolean value that indicates if the present stream can change to a
// different source.
//
// # Discussion
// 
// The default value is `true`.
//
// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCContentSharingPickerConfiguration-c.class/allowsChangingSelectedContent
func (c SCContentSharingPickerConfiguration) AllowsChangingSelectedContent() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("allowsChangingSelectedContent"))
	return rv
}
func (c SCContentSharingPickerConfiguration) SetAllowsChangingSelectedContent(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setAllowsChangingSelectedContent:"), value)
}

// A list of bundle IDs to exclude from the sharing picker.
//
// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCContentSharingPickerConfiguration-c.class/excludedBundleIDs
func (c SCContentSharingPickerConfiguration) ExcludedBundleIDs() []string {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("excludedBundleIDs"))
	return objc.ConvertSliceToStrings(rv)
}
func (c SCContentSharingPickerConfiguration) SetExcludedBundleIDs(value []string) {
	objc.Send[struct{}](c.ID, objc.Sel("setExcludedBundleIDs:"), objectivec.StringSliceToNSArray(value))
}

// A list of window IDs to exclude from the sharing picker.
//
// # Discussion
//
// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCContentSharingPickerConfiguration-c.class/excludedWindowIDs
func (c SCContentSharingPickerConfiguration) ExcludedWindowIDs() []foundation.NSNumber {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("excludedWindowIDs"))
	return objc.ConvertSlice(rv, func(id objc.ID) foundation.NSNumber {
		return foundation.NSNumberFromID(id)
	})
}
func (c SCContentSharingPickerConfiguration) SetExcludedWindowIDs(value []foundation.NSNumber) {
	objc.Send[struct{}](c.ID, objc.Sel("setExcludedWindowIDs:"), objectivec.IObjectSliceToNSArray(value))
}

