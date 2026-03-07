// Code generated from Apple documentation for ScreenCaptureKit. DO NOT EDIT.

package screencapturekit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SCContentSharingPicker] class.
var (
	_SCContentSharingPickerClass     SCContentSharingPickerClass
	_SCContentSharingPickerClassOnce sync.Once
)

func getSCContentSharingPickerClass() SCContentSharingPickerClass {
	_SCContentSharingPickerClassOnce.Do(func() {
		_SCContentSharingPickerClass = SCContentSharingPickerClass{class: objc.GetClass("SCContentSharingPicker")}
	})
	return _SCContentSharingPickerClass
}

// GetSCContentSharingPickerClass returns the class object for SCContentSharingPicker.
func GetSCContentSharingPickerClass() SCContentSharingPickerClass {
	return getSCContentSharingPickerClass()
}

type SCContentSharingPickerClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (sc SCContentSharingPickerClass) Alloc() SCContentSharingPicker {
	rv := objc.Send[SCContentSharingPicker](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}







// An instance of a picker presented by the operating system for managing
// frame-capture streams.
//
// # Overview
//
// # Picker availability
//
//   - [SCContentSharingPicker.Active]: A Boolean value that indicates if the picker is active.
//   - [SCContentSharingPicker.SetActive]
//
// # Stream configuration
//
//   - [SCContentSharingPicker.Configuration]: Sets the configuration for the content capture picker for all streams, providing allowed selection modes and content excluded from selection.
//   - [SCContentSharingPicker.SetConfiguration]
//   - [SCContentSharingPicker.DefaultConfiguration]: The default configuration to use for the content capture picker.
//   - [SCContentSharingPicker.SetDefaultConfiguration]
//   - [SCContentSharingPicker.MaximumStreamCount]: The maximum number of streams the content capture picker allows.
//   - [SCContentSharingPicker.SetMaximumStreamCount]
//
// # Manage observers
//
//   - [SCContentSharingPicker.AddObserver]: Adds an observer instance to notify of changes in the content-sharing picker.
//   - [SCContentSharingPicker.RemoveObserver]: Removes an observer instance from the content-sharing picker.
//
// # Picker display
//
//   - [SCContentSharingPicker.Present]: Displays the picker with no active selection for capture.
//   - [SCContentSharingPicker.PresentPickerForStream]: Displays the picker with an already running capture stream.
//   - [SCContentSharingPicker.PresentPickerUsingContentStyle]: Displays the picker for a single type of capture selection.
//   - [SCContentSharingPicker.PresentPickerForStreamUsingContentStyle]: Displays the picker with an existing capture stream, allowing for a single type of capture selection.
//
// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCContentSharingPicker
type SCContentSharingPicker struct {
	objectivec.Object
}

// SCContentSharingPickerFromID constructs a [SCContentSharingPicker] from an objc.ID.
//
// An instance of a picker presented by the operating system for managing
// frame-capture streams.
func SCContentSharingPickerFromID(id objc.ID) SCContentSharingPicker {
	return SCContentSharingPicker{objectivec.Object{id}}
}
// NOTE: SCContentSharingPicker adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [SCContentSharingPicker] class.
//
// # Picker availability
//
//   - [ISCContentSharingPicker.Active]: A Boolean value that indicates if the picker is active.
//   - [ISCContentSharingPicker.SetActive]
//
// # Stream configuration
//
//   - [ISCContentSharingPicker.Configuration]: Sets the configuration for the content capture picker for all streams, providing allowed selection modes and content excluded from selection.
//   - [ISCContentSharingPicker.SetConfiguration]
//   - [ISCContentSharingPicker.DefaultConfiguration]: The default configuration to use for the content capture picker.
//   - [ISCContentSharingPicker.SetDefaultConfiguration]
//   - [ISCContentSharingPicker.MaximumStreamCount]: The maximum number of streams the content capture picker allows.
//   - [ISCContentSharingPicker.SetMaximumStreamCount]
//
// # Manage observers
//
//   - [ISCContentSharingPicker.AddObserver]: Adds an observer instance to notify of changes in the content-sharing picker.
//   - [ISCContentSharingPicker.RemoveObserver]: Removes an observer instance from the content-sharing picker.
//
// # Picker display
//
//   - [ISCContentSharingPicker.Present]: Displays the picker with no active selection for capture.
//   - [ISCContentSharingPicker.PresentPickerForStream]: Displays the picker with an already running capture stream.
//   - [ISCContentSharingPicker.PresentPickerUsingContentStyle]: Displays the picker for a single type of capture selection.
//   - [ISCContentSharingPicker.PresentPickerForStreamUsingContentStyle]: Displays the picker with an existing capture stream, allowing for a single type of capture selection.
//
// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCContentSharingPicker
type ISCContentSharingPicker interface {
	objectivec.IObject

	// Topic: Picker availability

	// A Boolean value that indicates if the picker is active.
	Active() bool
	SetActive(value bool)

	// Topic: Stream configuration

	// Sets the configuration for the content capture picker for all streams, providing allowed selection modes and content excluded from selection.
	Configuration() ISCContentSharingPickerConfiguration
	SetConfiguration(value ISCContentSharingPickerConfiguration)
	// The default configuration to use for the content capture picker.
	DefaultConfiguration() ISCContentSharingPickerConfiguration
	SetDefaultConfiguration(value ISCContentSharingPickerConfiguration)
	// The maximum number of streams the content capture picker allows.
	MaximumStreamCount() int
	SetMaximumStreamCount(value int)

	// Topic: Manage observers

	// Adds an observer instance to notify of changes in the content-sharing picker.
	AddObserver(observer SCContentSharingPickerObserver)
	// Removes an observer instance from the content-sharing picker.
	RemoveObserver(observer SCContentSharingPickerObserver)

	// Topic: Picker display

	// Displays the picker with no active selection for capture.
	Present()
	// Displays the picker with an already running capture stream.
	PresentPickerForStream(stream ISCStream)
	// Displays the picker for a single type of capture selection.
	PresentPickerUsingContentStyle(contentStyle SCShareableContentStyle)
	// Displays the picker with an existing capture stream, allowing for a single type of capture selection.
	PresentPickerForStreamUsingContentStyle(stream ISCStream, contentStyle SCShareableContentStyle)
}





// Init initializes the instance.
func (c SCContentSharingPicker) Init() SCContentSharingPicker {
	rv := objc.Send[SCContentSharingPicker](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c SCContentSharingPicker) Autorelease() SCContentSharingPicker {
	rv := objc.Send[SCContentSharingPicker](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewSCContentSharingPicker creates a new SCContentSharingPicker instance.
func NewSCContentSharingPicker() SCContentSharingPicker {
	class := getSCContentSharingPickerClass()
	rv := objc.Send[SCContentSharingPicker](objc.ID(class.class), objc.Sel("new"))
	return rv
}










// Adds an observer instance to notify of changes in the content-sharing
// picker.
//
// observer: The observer instance to send notifications to.
//
// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCContentSharingPicker/add(_:)
func (c SCContentSharingPicker) AddObserver(observer SCContentSharingPickerObserver) {
	objc.Send[objc.ID](c.ID, objc.Sel("addObserver:"), observer)
}

// Removes an observer instance from the content-sharing picker.
//
// observer: The observer instance to remove.
//
// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCContentSharingPicker/remove(_:)
func (c SCContentSharingPicker) RemoveObserver(observer SCContentSharingPickerObserver) {
	objc.Send[objc.ID](c.ID, objc.Sel("removeObserver:"), observer)
}

// Displays the picker with no active selection for capture.
//
// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCContentSharingPicker/present()
func (c SCContentSharingPicker) Present() {
	objc.Send[objc.ID](c.ID, objc.Sel("present"))
}

// Displays the picker with an already running capture stream.
//
// stream: The capture stream to display in the picker.
//
// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCContentSharingPicker/present(for:)
func (c SCContentSharingPicker) PresentPickerForStream(stream ISCStream) {
	objc.Send[objc.ID](c.ID, objc.Sel("presentPickerForStream:"), stream)
}

// Displays the picker for a single type of capture selection.
//
// contentStyle: The type of streaming content selection allowed through the presented
// picker.
//
// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCContentSharingPicker/present(using:)
func (c SCContentSharingPicker) PresentPickerUsingContentStyle(contentStyle SCShareableContentStyle) {
	objc.Send[objc.ID](c.ID, objc.Sel("presentPickerUsingContentStyle:"), contentStyle)
}

// Displays the picker with an existing capture stream, allowing for a single
// type of capture selection.
//
// stream: The stream to display in the picker.
//
// contentStyle: The type of streaming content selection allowed through the presented
// picker.
//
// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCContentSharingPicker/present(for:using:)
func (c SCContentSharingPicker) PresentPickerForStreamUsingContentStyle(stream ISCStream, contentStyle SCShareableContentStyle) {
	objc.Send[objc.ID](c.ID, objc.Sel("presentPickerForStream:usingContentStyle:"), stream, contentStyle)
}












// A Boolean value that indicates if the picker is active.
//
// # Discussion
// 
// When this value is `true`, the capture stream picker is active, available
// for managing capture. The default value is `false`.
//
// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCContentSharingPicker/isActive
func (c SCContentSharingPicker) Active() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isActive"))
	return rv
}
func (c SCContentSharingPicker) SetActive(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setActive:"), value)
}



// Sets the configuration for the content capture picker for all streams,
// providing allowed selection modes and content excluded from selection.
//
// See: https://developer.apple.com/documentation/screencapturekit/sccontentsharingpicker/configuration
func (c SCContentSharingPicker) Configuration() ISCContentSharingPickerConfiguration {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("configuration"))
	return SCContentSharingPickerConfigurationFromID(objc.ID(rv))
}
func (c SCContentSharingPicker) SetConfiguration(value ISCContentSharingPickerConfiguration) {
	objc.Send[struct{}](c.ID, objc.Sel("setConfiguration:"), value)
}



// The default configuration to use for the content capture picker.
//
// See: https://developer.apple.com/documentation/screencapturekit/sccontentsharingpicker/defaultconfiguration-94q2b
func (c SCContentSharingPicker) DefaultConfiguration() ISCContentSharingPickerConfiguration {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("defaultConfiguration"))
	return SCContentSharingPickerConfigurationFromID(objc.ID(rv))
}
func (c SCContentSharingPicker) SetDefaultConfiguration(value ISCContentSharingPickerConfiguration) {
	objc.Send[struct{}](c.ID, objc.Sel("setDefaultConfiguration:"), value)
}



// The maximum number of streams the content capture picker allows.
//
// See: https://developer.apple.com/documentation/screencapturekit/sccontentsharingpicker/maximumstreamcount-2kuaa
func (c SCContentSharingPicker) MaximumStreamCount() int {
	rv := objc.Send[int](c.ID, objc.Sel("maximumStreamCount"))
	return rv
}
func (c SCContentSharingPicker) SetMaximumStreamCount(value int) {
	objc.Send[struct{}](c.ID, objc.Sel("setMaximumStreamCount:"), value)
}







// The system-provided picker UI instance for capturing display and audio
// content from someone’s Mac.
//
// # Discussion
// 
// The picker gives a person control over what information on their Mac they
// wish to let your app view or record such as specific applications,
// displays, and windows.
//
// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCContentSharingPicker/shared
func (_SCContentSharingPickerClass SCContentSharingPickerClass) SharedPicker() SCContentSharingPicker {
	rv := objc.Send[objc.ID](objc.ID(_SCContentSharingPickerClass.class), objc.Sel("sharedPicker"))
	return SCContentSharingPickerFromID(objc.ID(rv))
}



















