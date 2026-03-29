// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVCaptureMovieFileOutput] class.
var (
	_AVCaptureMovieFileOutputClass     AVCaptureMovieFileOutputClass
	_AVCaptureMovieFileOutputClassOnce sync.Once
)

func getAVCaptureMovieFileOutputClass() AVCaptureMovieFileOutputClass {
	_AVCaptureMovieFileOutputClassOnce.Do(func() {
		_AVCaptureMovieFileOutputClass = AVCaptureMovieFileOutputClass{class: objc.GetClass("AVCaptureMovieFileOutput")}
	})
	return _AVCaptureMovieFileOutputClass
}

// GetAVCaptureMovieFileOutputClass returns the class object for AVCaptureMovieFileOutput.
func GetAVCaptureMovieFileOutputClass() AVCaptureMovieFileOutputClass {
	return getAVCaptureMovieFileOutputClass()
}

type AVCaptureMovieFileOutputClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVCaptureMovieFileOutputClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVCaptureMovieFileOutputClass) Alloc() AVCaptureMovieFileOutput {
	rv := objc.Send[AVCaptureMovieFileOutput](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// A capture output that records video and audio to a QuickTime movie file.
//
// # Overview
// 
// A movie file output provides a complete file recording interface for
// writing media data to QuickTime movie files. It includes the ability to
// configure QuickTime-specific options, including writing metadata
// collections to each file, specify media encoding options for each track,
// and specify the interval at which it writes movie fragments.
//
// # Configuring movies
//
//   - [AVCaptureMovieFileOutput.MovieFragmentInterval]: The number of seconds of output that are written per fragment.
//   - [AVCaptureMovieFileOutput.SetMovieFragmentInterval]
//   - [AVCaptureMovieFileOutput.Metadata]: The metadata for the output file.
//   - [AVCaptureMovieFileOutput.SetMetadata]
//
// # Managing output settings
//
//   - [AVCaptureMovieFileOutput.OutputSettingsForConnection]: Returns the settings the output uses to encode media from the specified connection.
//   - [AVCaptureMovieFileOutput.SetOutputSettingsForConnection]: Sets the options the output uses to encode media from the given connection while recording.
//
// # Enabling spatial capture
//
//   - [AVCaptureMovieFileOutput.SpatialVideoCaptureSupported]: A Boolean value that indicates whether a movie file output supports capturing spatial videos.
//   - [AVCaptureMovieFileOutput.SpatialVideoCaptureEnabled]: A Boolean value that indicates whether a movie file output captures spatial videos.
//   - [AVCaptureMovieFileOutput.SetSpatialVideoCaptureEnabled]
//
// # Restricting camera switching
//
//   - [AVCaptureMovieFileOutput.PrimaryConstituentDeviceSwitchingBehaviorForRecordingEnabled]: A Boolean value that indicates whether to restrict constituent device switching behavior during recording.
//   - [AVCaptureMovieFileOutput.SetPrimaryConstituentDeviceSwitchingBehaviorForRecordingEnabled]
//   - [AVCaptureMovieFileOutput.SetPrimaryConstituentDeviceSwitchingBehaviorForRecordingRestrictedSwitchingBehaviorConditions]: Sets the camera switching behavior to use during recording.
//   - [AVCaptureMovieFileOutput.PrimaryConstituentDeviceSwitchingBehaviorForRecording]: The camera switching behavior to use for recording.
//   - [AVCaptureMovieFileOutput.PrimaryConstituentDeviceRestrictedSwitchingBehaviorConditionsForRecording]: The conditions during which camera switching may occur while recording.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureMovieFileOutput
type AVCaptureMovieFileOutput struct {
	AVCaptureFileOutput
}

// AVCaptureMovieFileOutputFromID constructs a [AVCaptureMovieFileOutput] from an objc.ID.
//
// A capture output that records video and audio to a QuickTime movie file.
func AVCaptureMovieFileOutputFromID(id objc.ID) AVCaptureMovieFileOutput {
	return AVCaptureMovieFileOutput{AVCaptureFileOutput: AVCaptureFileOutputFromID(id)}
}
// NOTE: AVCaptureMovieFileOutput adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVCaptureMovieFileOutput] class.
//
// # Configuring movies
//
//   - [IAVCaptureMovieFileOutput.MovieFragmentInterval]: The number of seconds of output that are written per fragment.
//   - [IAVCaptureMovieFileOutput.SetMovieFragmentInterval]
//   - [IAVCaptureMovieFileOutput.Metadata]: The metadata for the output file.
//   - [IAVCaptureMovieFileOutput.SetMetadata]
//
// # Managing output settings
//
//   - [IAVCaptureMovieFileOutput.OutputSettingsForConnection]: Returns the settings the output uses to encode media from the specified connection.
//   - [IAVCaptureMovieFileOutput.SetOutputSettingsForConnection]: Sets the options the output uses to encode media from the given connection while recording.
//
// # Enabling spatial capture
//
//   - [IAVCaptureMovieFileOutput.SpatialVideoCaptureSupported]: A Boolean value that indicates whether a movie file output supports capturing spatial videos.
//   - [IAVCaptureMovieFileOutput.SpatialVideoCaptureEnabled]: A Boolean value that indicates whether a movie file output captures spatial videos.
//   - [IAVCaptureMovieFileOutput.SetSpatialVideoCaptureEnabled]
//
// # Restricting camera switching
//
//   - [IAVCaptureMovieFileOutput.PrimaryConstituentDeviceSwitchingBehaviorForRecordingEnabled]: A Boolean value that indicates whether to restrict constituent device switching behavior during recording.
//   - [IAVCaptureMovieFileOutput.SetPrimaryConstituentDeviceSwitchingBehaviorForRecordingEnabled]
//   - [IAVCaptureMovieFileOutput.SetPrimaryConstituentDeviceSwitchingBehaviorForRecordingRestrictedSwitchingBehaviorConditions]: Sets the camera switching behavior to use during recording.
//   - [IAVCaptureMovieFileOutput.PrimaryConstituentDeviceSwitchingBehaviorForRecording]: The camera switching behavior to use for recording.
//   - [IAVCaptureMovieFileOutput.PrimaryConstituentDeviceRestrictedSwitchingBehaviorConditionsForRecording]: The conditions during which camera switching may occur while recording.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureMovieFileOutput
type IAVCaptureMovieFileOutput interface {
	IAVCaptureFileOutput

	// Topic: Configuring movies

	// The number of seconds of output that are written per fragment.
	MovieFragmentInterval() uintptr
	SetMovieFragmentInterval(value uintptr)
	// The metadata for the output file.
	Metadata() []AVMetadataItem
	SetMetadata(value []AVMetadataItem)

	// Topic: Managing output settings

	// Returns the settings the output uses to encode media from the specified connection.
	OutputSettingsForConnection(connection IAVCaptureConnection) foundation.INSDictionary
	// Sets the options the output uses to encode media from the given connection while recording.
	SetOutputSettingsForConnection(outputSettings foundation.INSDictionary, connection IAVCaptureConnection)

	// Topic: Enabling spatial capture

	// A Boolean value that indicates whether a movie file output supports capturing spatial videos.
	SpatialVideoCaptureSupported() bool
	// A Boolean value that indicates whether a movie file output captures spatial videos.
	SpatialVideoCaptureEnabled() bool
	SetSpatialVideoCaptureEnabled(value bool)

	// Topic: Restricting camera switching

	// A Boolean value that indicates whether to restrict constituent device switching behavior during recording.
	PrimaryConstituentDeviceSwitchingBehaviorForRecordingEnabled() bool
	SetPrimaryConstituentDeviceSwitchingBehaviorForRecordingEnabled(value bool)
	// Sets the camera switching behavior to use during recording.
	SetPrimaryConstituentDeviceSwitchingBehaviorForRecordingRestrictedSwitchingBehaviorConditions(switchingBehavior AVCapturePrimaryConstituentDeviceSwitchingBehavior, restrictedSwitchingBehaviorConditions AVCapturePrimaryConstituentDeviceRestrictedSwitchingBehaviorConditions)
	// The camera switching behavior to use for recording.
	PrimaryConstituentDeviceSwitchingBehaviorForRecording() AVCapturePrimaryConstituentDeviceSwitchingBehavior
	// The conditions during which camera switching may occur while recording.
	PrimaryConstituentDeviceRestrictedSwitchingBehaviorConditionsForRecording() AVCapturePrimaryConstituentDeviceRestrictedSwitchingBehaviorConditions
}

// Init initializes the instance.
func (c AVCaptureMovieFileOutput) Init() AVCaptureMovieFileOutput {
	rv := objc.Send[AVCaptureMovieFileOutput](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c AVCaptureMovieFileOutput) Autorelease() AVCaptureMovieFileOutput {
	rv := objc.Send[AVCaptureMovieFileOutput](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVCaptureMovieFileOutput creates a new AVCaptureMovieFileOutput instance.
func NewAVCaptureMovieFileOutput() AVCaptureMovieFileOutput {
	class := getAVCaptureMovieFileOutputClass()
	rv := objc.Send[AVCaptureMovieFileOutput](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Returns the settings the output uses to encode media from the specified
// connection.
//
// connection: The connection delivering the media to encode.
//
// # Return Value
// 
// A dictionary of output settings.
//
// # Discussion
// 
// If the returned value is an empty dictionary, the format of the media from
// the connection isn’t changed before writing to the file.
// 
// If you call [SetOutputSettingsForConnection] with a `nil` dictionary, this
// method returns a non-`nil` dictionary that reflects the settings used by
// the capture session’s [SessionPreset] value.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureMovieFileOutput/outputSettings(for:)
func (c AVCaptureMovieFileOutput) OutputSettingsForConnection(connection IAVCaptureConnection) foundation.INSDictionary {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("outputSettingsForConnection:"), connection)
	return foundation.NSDictionaryFromID(rv)
}
// Sets the options the output uses to encode media from the given connection
// while recording.
//
// outputSettings: A dictionary of output settings. Pass an empty dictionary to specify that
// the format of the media from the connection shouldn’t change before
// writing to the file. Pass `nil` to specify that the session preset
// determines output format.
//
// connection: The connection delivering the media to encode.
//
// # Discussion
// 
// For details on output settings, see [Video settings] for video connections
// and [Audio settings] for audio connections.
// 
// On iOS, your output settings dictionary may only contain keys listed
// returned from the [SupportedOutputSettingsKeysForConnection] method. If you
// specify any other key, the system throws an invalid argument exception.
// Additionally, the value you specify for [AVVideoCodecKey] should be present
// in the [AvailableVideoCodecTypes] array. If you specify
// [AVVideoCompressionPropertiesKey], you must also specify a valid value for
// [AVVideoCodecKey].
// 
// On iOS, the [OutputSettingsForConnection] method always provides a fully
// populated dictionary. If you call [OutputSettingsForConnection] with the
// intent of overriding a few of the values, you must exclude keys that
// aren’t supported before calling [SetOutputSettingsForConnection]. When
// providing an [AVVideoCompressionPropertiesKey] sub dictionary, you may
// specify a sparse dictionary. A movie file output object always fills in
// missing keys with default values for the current capture session
// configuration.
//
// [AVVideoCodecKey]: https://developer.apple.com/documentation/AVFoundation/AVVideoCodecKey
// [AVVideoCompressionPropertiesKey]: https://developer.apple.com/documentation/AVFoundation/AVVideoCompressionPropertiesKey
// [Audio settings]: https://developer.apple.com/documentation/AVFoundation/audio-settings
// [Video settings]: https://developer.apple.com/documentation/AVFoundation/video-settings
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureMovieFileOutput/setOutputSettings(_:for:)
func (c AVCaptureMovieFileOutput) SetOutputSettingsForConnection(outputSettings foundation.INSDictionary, connection IAVCaptureConnection) {
	objc.Send[objc.ID](c.ID, objc.Sel("setOutputSettings:forConnection:"), outputSettings, connection)
}
// Sets the camera switching behavior to use during recording.
//
// switchingBehavior: The switching behavior to set on the movie file output.
// 
// Attempting to restrict the switching behavior of a capture device that
// doesn’t support constituent device switching results in an error.
//
// restrictedSwitchingBehaviorConditions: The conditions during which camera switching occurs. Only set a condition
// when you set the switching behavior to
// [CapturePrimaryConstituentDeviceSwitchingBehaviorRestricted]. In all other
// cases, set the value to
// [CapturePrimaryConstituentDeviceRestrictedSwitchingBehaviorConditionNone].
//
// # Discussion
// 
// Use this method to control the camera switching behavior the system uses
// when recording a movie. The behavior you specify takes effect when you
// enable it by setting the value of
// [PrimaryConstituentDeviceSwitchingBehaviorForRecordingEnabled] to [true].
// 
// When a capture device doesn’t support constituent device selection,
// attempting to set a behavior other than
// [CapturePrimaryConstituentDeviceSwitchingBehaviorUnsupported] causes the
// system to throw an invalid argument exception.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureMovieFileOutput/setPrimaryConstituentDeviceSwitchingBehaviorForRecording(_:restrictedSwitchingBehaviorConditions:)
func (c AVCaptureMovieFileOutput) SetPrimaryConstituentDeviceSwitchingBehaviorForRecordingRestrictedSwitchingBehaviorConditions(switchingBehavior AVCapturePrimaryConstituentDeviceSwitchingBehavior, restrictedSwitchingBehaviorConditions AVCapturePrimaryConstituentDeviceRestrictedSwitchingBehaviorConditions) {
	objc.Send[objc.ID](c.ID, objc.Sel("setPrimaryConstituentDeviceSwitchingBehaviorForRecording:restrictedSwitchingBehaviorConditions:"), switchingBehavior, restrictedSwitchingBehaviorConditions)
}

// The number of seconds of output that are written per fragment.
//
// # Discussion
// 
// The default is 10 seconds. To disable movie fragment writing (not typically
// recommended), set to [invalid].
// 
// A QuickTime movie contains media samples and a table identifying their
// location in the file. A movie file without a sample table is unreadable.
// 
// In a processed file, the sample table typically appears at the beginning of
// the file. It may also appear at the end of the file, in which case the
// header contains a pointer to the sample table at the end. When a new movie
// file is being recorded, it is not possible to write the sample table since
// the size of the file is not yet known. Instead, the table is must be
// written when recording is complete. If no other action is taken, this means
// that if the recording does not complete successfully (for example, in the
// event of a crash), the file data is unusable (because there is no sample
// table). By periodically inserting “movie fragments” into the movie
// file, the sample table can be built up incrementally. This means that if
// the file is not written completely, the movie file is still usable (up to
// the point where the last fragment was written).
//
// [invalid]: https://developer.apple.com/documentation/CoreMedia/CMTime/invalid
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureMovieFileOutput/movieFragmentInterval
func (c AVCaptureMovieFileOutput) MovieFragmentInterval() uintptr {
	rv := objc.Send[uintptr](c.ID, objc.Sel("movieFragmentInterval"))
	return rv
}
func (c AVCaptureMovieFileOutput) SetMovieFragmentInterval(value uintptr) {
	objc.Send[struct{}](c.ID, objc.Sel("setMovieFragmentInterval:"), value)
}
// The metadata for the output file.
//
// # Discussion
// 
// This array contains [AVMetadataItem] objects. You use it to add metadata,
// such as copyright, creation date, and so on, to the recorded movie file.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureMovieFileOutput/metadata
func (c AVCaptureMovieFileOutput) Metadata() []AVMetadataItem {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("metadata"))
	return objc.ConvertSlice(rv, func(id objc.ID) AVMetadataItem {
		return AVMetadataItemFromID(id)
	})
}
func (c AVCaptureMovieFileOutput) SetMetadata(value []AVMetadataItem) {
	objc.Send[struct{}](c.ID, objc.Sel("setMetadata:"), objectivec.IObjectSliceToNSArray(value))
}
// A Boolean value that indicates whether a movie file output supports
// capturing spatial videos.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureMovieFileOutput/isSpatialVideoCaptureSupported
func (c AVCaptureMovieFileOutput) SpatialVideoCaptureSupported() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isSpatialVideoCaptureSupported"))
	return rv
}
// A Boolean value that indicates whether a movie file output captures spatial
// videos.
//
// # Discussion
// 
// Spatial capture lets you record your favorite moments in 3D for playback on
// Apple Vision Pro. This feature isn’t supported on all devices, so you can
// only enable this property when [SpatialVideoCaptureSupported] is [true].
// 
// The default value is [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureMovieFileOutput/isSpatialVideoCaptureEnabled
func (c AVCaptureMovieFileOutput) SpatialVideoCaptureEnabled() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isSpatialVideoCaptureEnabled"))
	return rv
}
func (c AVCaptureMovieFileOutput) SetSpatialVideoCaptureEnabled(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setSpatialVideoCaptureEnabled:"), value)
}
// A Boolean value that indicates whether to restrict constituent device
// switching behavior during recording.
//
// # Discussion
// 
// Use this property to enable camera switching restrictions when recording
// movies. You set restrictions by calling the output’s
// [SetPrimaryConstituentDeviceSwitchingBehaviorForRecordingRestrictedSwitchingBehaviorConditions]
// method. The restrictions take effect when you start recording, and revert
// to the behavior set by the capture device’s
// [PrimaryConstituentDeviceSwitchingBehavior] when you stop recording.
// 
// By default, this property is [true] when connected to a capture device that
// supports constituent device switching.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureMovieFileOutput/isPrimaryConstituentDeviceSwitchingBehaviorForRecordingEnabled
func (c AVCaptureMovieFileOutput) PrimaryConstituentDeviceSwitchingBehaviorForRecordingEnabled() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isPrimaryConstituentDeviceSwitchingBehaviorForRecordingEnabled"))
	return rv
}
func (c AVCaptureMovieFileOutput) SetPrimaryConstituentDeviceSwitchingBehaviorForRecordingEnabled(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setPrimaryConstituentDeviceSwitchingBehaviorForRecordingEnabled:"), value)
}
// The camera switching behavior to use for recording.
//
// # Discussion
// 
// The default value of this property is
// [CapturePrimaryConstituentDeviceSwitchingBehaviorRestricted].
// 
// This property is key-value observable.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureMovieFileOutput/primaryConstituentDeviceSwitchingBehaviorForRecording
func (c AVCaptureMovieFileOutput) PrimaryConstituentDeviceSwitchingBehaviorForRecording() AVCapturePrimaryConstituentDeviceSwitchingBehavior {
	rv := objc.Send[AVCapturePrimaryConstituentDeviceSwitchingBehavior](c.ID, objc.Sel("primaryConstituentDeviceSwitchingBehaviorForRecording"))
	return AVCapturePrimaryConstituentDeviceSwitchingBehavior(rv)
}
// The conditions during which camera switching may occur while recording.
//
// # Discussion
// 
// The default conditions include
// [CapturePrimaryConstituentDeviceRestrictedSwitchingBehaviorConditionVideoZoomChanged],
// [CapturePrimaryConstituentDeviceRestrictedSwitchingBehaviorConditionFocusModeChanged],
// and
// [CapturePrimaryConstituentDeviceRestrictedSwitchingBehaviorConditionExposureModeChanged].
// 
// This property is key-value observable.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureMovieFileOutput/primaryConstituentDeviceRestrictedSwitchingBehaviorConditionsForRecording
func (c AVCaptureMovieFileOutput) PrimaryConstituentDeviceRestrictedSwitchingBehaviorConditionsForRecording() AVCapturePrimaryConstituentDeviceRestrictedSwitchingBehaviorConditions {
	rv := objc.Send[AVCapturePrimaryConstituentDeviceRestrictedSwitchingBehaviorConditions](c.ID, objc.Sel("primaryConstituentDeviceRestrictedSwitchingBehaviorConditionsForRecording"))
	return AVCapturePrimaryConstituentDeviceRestrictedSwitchingBehaviorConditions(rv)
}

