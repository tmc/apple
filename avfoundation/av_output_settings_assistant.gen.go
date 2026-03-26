// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVOutputSettingsAssistant] class.
var (
	_AVOutputSettingsAssistantClass     AVOutputSettingsAssistantClass
	_AVOutputSettingsAssistantClassOnce sync.Once
)

func getAVOutputSettingsAssistantClass() AVOutputSettingsAssistantClass {
	_AVOutputSettingsAssistantClassOnce.Do(func() {
		_AVOutputSettingsAssistantClass = AVOutputSettingsAssistantClass{class: objc.GetClass("AVOutputSettingsAssistant")}
	})
	return _AVOutputSettingsAssistantClass
}

// GetAVOutputSettingsAssistantClass returns the class object for AVOutputSettingsAssistant.
func GetAVOutputSettingsAssistantClass() AVOutputSettingsAssistantClass {
	return getAVOutputSettingsAssistantClass()
}

type AVOutputSettingsAssistantClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVOutputSettingsAssistantClass) Alloc() AVOutputSettingsAssistant {
	rv := objc.Send[AVOutputSettingsAssistant](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that builds audio and video output settings dictionaries.
//
// # Overview
// 
// Use an output settings assistant to create the audio and video settings
// that you use to configure instances of [AVAssetWriter] and
// [AVAssetWriterInput]. You create an assistant with a specific preset
// configuration, such as [hevc3840x2160WithAlpha] or [preset1920x1080]. You
// can accept the settings dictionaries as is to generate a file that conforms
// to the criteria that the preset implies. You may also use the dictionaries
// it generates as a base configuration that you can customize as you require.
// 
// Providing the assistant additional details about your source media helps it
// generate more complete results. For example, setting a value for its
// [AVOutputSettingsAssistant.SourceVideoFormat] property ensures that the assistant generates settings
// that don’t scale up video frames from a smaller size.
//
// [hevc3840x2160WithAlpha]: https://developer.apple.com/documentation/AVFoundation/AVOutputSettingsPreset/hevc3840x2160WithAlpha
// [preset1920x1080]: https://developer.apple.com/documentation/AVFoundation/AVOutputSettingsPreset/preset1920x1080
//
// # Configuring output settings
//
//   - [AVOutputSettingsAssistant.OutputFileType]: A uniform type identifier (UTI) that indicates the type of file to write.
//   - [AVOutputSettingsAssistant.AudioSettings]: An audio settings dictionary.
//   - [AVOutputSettingsAssistant.VideoSettings]: A video settings dictionary.
//   - [AVOutputSettingsAssistant.SourceVideoMinFrameDuration]: A time value that describes the minimum frame duration of the video data.
//   - [AVOutputSettingsAssistant.SetSourceVideoMinFrameDuration]
//   - [AVOutputSettingsAssistant.SourceVideoAverageFrameDuration]: A time value that describes the average frame duration of the video data.
//   - [AVOutputSettingsAssistant.SetSourceVideoAverageFrameDuration]
//
// See: https://developer.apple.com/documentation/AVFoundation/AVOutputSettingsAssistant
type AVOutputSettingsAssistant struct {
	objectivec.Object
}

// AVOutputSettingsAssistantFromID constructs a [AVOutputSettingsAssistant] from an objc.ID.
//
// An object that builds audio and video output settings dictionaries.
func AVOutputSettingsAssistantFromID(id objc.ID) AVOutputSettingsAssistant {
	return AVOutputSettingsAssistant{objectivec.Object{ID: id}}
}
// NOTE: AVOutputSettingsAssistant adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVOutputSettingsAssistant] class.
//
// # Configuring output settings
//
//   - [IAVOutputSettingsAssistant.OutputFileType]: A uniform type identifier (UTI) that indicates the type of file to write.
//   - [IAVOutputSettingsAssistant.AudioSettings]: An audio settings dictionary.
//   - [IAVOutputSettingsAssistant.VideoSettings]: A video settings dictionary.
//   - [IAVOutputSettingsAssistant.SourceVideoMinFrameDuration]: A time value that describes the minimum frame duration of the video data.
//   - [IAVOutputSettingsAssistant.SetSourceVideoMinFrameDuration]
//   - [IAVOutputSettingsAssistant.SourceVideoAverageFrameDuration]: A time value that describes the average frame duration of the video data.
//   - [IAVOutputSettingsAssistant.SetSourceVideoAverageFrameDuration]
//
// See: https://developer.apple.com/documentation/AVFoundation/AVOutputSettingsAssistant
type IAVOutputSettingsAssistant interface {
	objectivec.IObject

	// Topic: Configuring output settings

	// A uniform type identifier (UTI) that indicates the type of file to write.
	OutputFileType() AVFileType
	// An audio settings dictionary.
	AudioSettings() foundation.INSDictionary
	// A video settings dictionary.
	VideoSettings() foundation.INSDictionary
	// A time value that describes the minimum frame duration of the video data.
	SourceVideoMinFrameDuration() objectivec.IObject
	SetSourceVideoMinFrameDuration(value objectivec.IObject)
	// A time value that describes the average frame duration of the video data.
	SourceVideoAverageFrameDuration() objectivec.IObject
	SetSourceVideoAverageFrameDuration(value objectivec.IObject)

	// A preset for HEVC with Alpha video at 3840 by 2160 pixels.
	Hevc3840x2160WithAlpha() AVOutputSettingsPreset
	// A preset for H.264 video at 1920 by 1080 pixels.
	Preset1920x1080() AVOutputSettingsPreset
}

// Init initializes the instance.
func (o AVOutputSettingsAssistant) Init() AVOutputSettingsAssistant {
	rv := objc.Send[AVOutputSettingsAssistant](o.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (o AVOutputSettingsAssistant) Autorelease() AVOutputSettingsAssistant {
	rv := objc.Send[AVOutputSettingsAssistant](o.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVOutputSettingsAssistant creates a new AVOutputSettingsAssistant instance.
func NewAVOutputSettingsAssistant() AVOutputSettingsAssistant {
	class := getAVOutputSettingsAssistantClass()
	rv := objc.Send[AVOutputSettingsAssistant](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates an output setting assistant with a preset configuration.
//
// # Discussion
// 
// - presetIdentifier: A preset configuration for the object.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVOutputSettingsAssistant/init(preset:)
func NewOutputSettingsAssistantWithPreset(presetIdentifier AVOutputSettingsPreset) AVOutputSettingsAssistant {
	rv := objc.Send[objc.ID](objc.ID(getAVOutputSettingsAssistantClass().class), objc.Sel("outputSettingsAssistantWithPreset:"), objc.String(string(presetIdentifier)))
	return AVOutputSettingsAssistantFromID(rv)
}

// Returns an array of preset values to use to initialize an output settings
// assistant.
//
// # Return Value
// 
// An array of available output settings presets.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVOutputSettingsAssistant/availableOutputSettingsPresets()
func (_AVOutputSettingsAssistantClass AVOutputSettingsAssistantClass) AvailableOutputSettingsPresets() []string {
	rv := objc.Send[[]objc.ID](objc.ID(_AVOutputSettingsAssistantClass.class), objc.Sel("availableOutputSettingsPresets"))
	return objc.ConvertSliceToStrings(rv)
}

// A uniform type identifier (UTI) that indicates the type of file to write.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVOutputSettingsAssistant/outputFileType
func (o AVOutputSettingsAssistant) OutputFileType() AVFileType {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputFileType"))
	return AVFileType(foundation.NSStringFromID(rv).String())
}
// An audio settings dictionary.
//
// # Discussion
// 
// The value of this property may change as a result of setting a new value
// for the [SourceAudioFormat] property. See [Audio settings] for keys and
// values.
//
// [Audio settings]: https://developer.apple.com/documentation/AVFoundation/audio-settings
//
// See: https://developer.apple.com/documentation/AVFoundation/AVOutputSettingsAssistant/audioSettings
func (o AVOutputSettingsAssistant) AudioSettings() foundation.INSDictionary {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("audioSettings"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
// A video settings dictionary.
//
// # Discussion
// 
// The value of this property may change as a result of setting a new value
// for the [SourceVideoFormat] property. See [Video settings] for the
// supported keys and values.
//
// [Video settings]: https://developer.apple.com/documentation/AVFoundation/video-settings
//
// See: https://developer.apple.com/documentation/AVFoundation/AVOutputSettingsAssistant/videoSettings
func (o AVOutputSettingsAssistant) VideoSettings() foundation.INSDictionary {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("videoSettings"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
// A time value that describes the minimum frame duration of the video data.
//
// # Discussion
// 
// Setting this property enables the output settings assistant to generate
// more complete video settings. After setting a value, requery the
// [VideoSettings] property to get the latest values.
// 
// If the source of the video data is an instance of [AVAssetReaderOutput],
// you can discover the minimum frame duration of your source asset using the
// [AVAssetTrack] instance’s [minFrameDuration] property.
// 
// The default value is `1/30`, which means the output settings assistant
// assumes that the source video has a maximum frame rate of 30fps.
//
// [minFrameDuration]: https://developer.apple.com/documentation/AVFoundation/AVAssetTrack/minFrameDuration
//
// See: https://developer.apple.com/documentation/AVFoundation/AVOutputSettingsAssistant/sourceVideoMinFrameDuration
func (o AVOutputSettingsAssistant) SourceVideoMinFrameDuration() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("sourceVideoMinFrameDuration"))
	return objectivec.Object{ID: rv}
}
func (o AVOutputSettingsAssistant) SetSourceVideoMinFrameDuration(value objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setSourceVideoMinFrameDuration:"), value)
}
// A time value that describes the average frame duration of the video data.
//
// # Discussion
// 
// Setting this property enables the output settings assistant to generate
// more complete video settings. After setting a value, requery the
// [VideoSettings] property to get the latest values.
// 
// The default value is `1/30`, which means the output settings assistant
// assumes that your source video has a frame rate of 30fps.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVOutputSettingsAssistant/sourceVideoAverageFrameDuration
func (o AVOutputSettingsAssistant) SourceVideoAverageFrameDuration() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("sourceVideoAverageFrameDuration"))
	return objectivec.Object{ID: rv}
}
func (o AVOutputSettingsAssistant) SetSourceVideoAverageFrameDuration(value objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setSourceVideoAverageFrameDuration:"), value)
}
// A preset for HEVC with Alpha video at 3840 by 2160 pixels.
//
// See: https://developer.apple.com/documentation/avfoundation/avoutputsettingspreset/hevc3840x2160withalpha
func (o AVOutputSettingsAssistant) Hevc3840x2160WithAlpha() AVOutputSettingsPreset {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("AVOutputSettingsPresetHEVC3840x2160WithAlpha"))
	return AVOutputSettingsPreset(foundation.NSStringFromID(rv).String())
}
// A preset for H.264 video at 1920 by 1080 pixels.
//
// See: https://developer.apple.com/documentation/avfoundation/avoutputsettingspreset/preset1920x1080
func (o AVOutputSettingsAssistant) Preset1920x1080() AVOutputSettingsPreset {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("AVOutputSettingsPreset1920x1080"))
	return AVOutputSettingsPreset(foundation.NSStringFromID(rv).String())
}

