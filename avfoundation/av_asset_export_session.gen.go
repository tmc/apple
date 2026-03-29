// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"context"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVAssetExportSession] class.
var (
	_AVAssetExportSessionClass     AVAssetExportSessionClass
	_AVAssetExportSessionClassOnce sync.Once
)

func getAVAssetExportSessionClass() AVAssetExportSessionClass {
	_AVAssetExportSessionClassOnce.Do(func() {
		_AVAssetExportSessionClass = AVAssetExportSessionClass{class: objc.GetClass("AVAssetExportSession")}
	})
	return _AVAssetExportSessionClass
}

// GetAVAssetExportSessionClass returns the class object for AVAssetExportSession.
func GetAVAssetExportSessionClass() AVAssetExportSessionClass {
	return getAVAssetExportSessionClass()
}

type AVAssetExportSessionClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVAssetExportSessionClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVAssetExportSessionClass) Alloc() AVAssetExportSession {
	rv := objc.Send[AVAssetExportSession](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that exports assets in a format that you specify using an export
// preset.
//
// # Overview
// 
// You configure this object to export an instance of [AVAsset] by setting an
// export preset, an output file type, and an output URL.
//
// # Creating an export session
//
//   - [AVAssetExportSession.InitWithAssetPresetName]: Creates an export session with a preset configuration.
//
// # Accessing export presets
//
//   - [AVAssetExportSession.PresetName]: The name of the preset that the asset export session uses.
//   - [AVAssetExportSession.DetermineCompatibleFileTypesWithCompletionHandler]: Determines the output file types an asset export session supports writing in its current configuration.
//
// # Configuring output
//
//   - [AVAssetExportSession.SupportedFileTypes]: An array containing the types of files the session can write.
//   - [AVAssetExportSession.AllowsParallelizedExport]: A Boolean value that indicates whether the session can parallelize its export operation.
//   - [AVAssetExportSession.SetAllowsParallelizedExport]
//   - [AVAssetExportSession.ShouldOptimizeForNetworkUse]: A Boolean value that indicates whether to optimize the movie for network use.
//   - [AVAssetExportSession.SetShouldOptimizeForNetworkUse]
//   - [AVAssetExportSession.CanPerformMultiplePassesOverSourceMediaData]: A Boolean value that indicates whether the export session can perform multiple passes over the source media to achieve better results.
//   - [AVAssetExportSession.SetCanPerformMultiplePassesOverSourceMediaData]
//   - [AVAssetExportSession.TimeRange]: The time range of the source asset to export.
//   - [AVAssetExportSession.SetTimeRange]
//   - [AVAssetExportSession.FileLengthLimit]: The file length that the output of the session must not exceed.
//   - [AVAssetExportSession.SetFileLengthLimit]
//   - [AVAssetExportSession.DirectoryForTemporaryFiles]: A directory suitable to store temporary files that the export process generates.
//   - [AVAssetExportSession.SetDirectoryForTemporaryFiles]
//
// # Configuring metadata
//
//   - [AVAssetExportSession.Metadata]: The metadata an export session writes to the output container file.
//   - [AVAssetExportSession.SetMetadata]
//   - [AVAssetExportSession.MetadataItemFilter]: An object the export session uses to filter the metadata items it transfers to the output asset.
//   - [AVAssetExportSession.SetMetadataItemFilter]
//
// # Configuring video output
//
//   - [AVAssetExportSession.VideoComposition]: An optional object that provides instructions for how to composite frames of video.
//   - [AVAssetExportSession.SetVideoComposition]
//   - [AVAssetExportSession.CustomVideoCompositor]: An optional custom object to use when compositing video frames.
//
// # Configuring track groups
//
//   - [AVAssetExportSession.AudioTrackGroupHandling]: A policy that defines how the session exports alternate audio tracks.
//   - [AVAssetExportSession.SetAudioTrackGroupHandling]
//
// # Configuring audio output
//
//   - [AVAssetExportSession.AudioMix]: The parameters for audio mixing and an indication of whether to enable nondefault audio mixing for export.
//   - [AVAssetExportSession.SetAudioMix]
//   - [AVAssetExportSession.AudioTimePitchAlgorithm]: A processing algorithm for managing audio pitch for scaled audio edits.
//   - [AVAssetExportSession.SetAudioTimePitchAlgorithm]
//
// # Monitoring export progress
//
//   - [AVAssetExportSession.Status]: The status of the export session.
//   - [AVAssetExportSession.Error]: An optional error object.
//
// # Estimating file length and duration
//
//   - [AVAssetExportSession.EstimateOutputFileLengthWithCompletionHandler]: Starts estimating the output file length of the export while considering the asset, preset, and time range configuration of the export session.
//   - [AVAssetExportSession.EstimatedOutputFileLength]: The estimated length of the exported file, in bytes.
//
// # Estimating duration
//
//   - [AVAssetExportSession.EstimateMaximumDurationWithCompletionHandler]: Starts estimating the maximum duration of the export while considering the asset, preset, and time range configuration of the export session.
//   - [AVAssetExportSession.MaxDuration]: Provides an estimate of the maximum duration of the exported media.
//
// # Accessing the asset
//
//   - [AVAssetExportSession.Asset]: An asset that a session exports.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetExportSession
type AVAssetExportSession struct {
	objectivec.Object
}

// AVAssetExportSessionFromID constructs a [AVAssetExportSession] from an objc.ID.
//
// An object that exports assets in a format that you specify using an export
// preset.
func AVAssetExportSessionFromID(id objc.ID) AVAssetExportSession {
	return AVAssetExportSession{objectivec.Object{ID: id}}
}
// NOTE: AVAssetExportSession adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVAssetExportSession] class.
//
// # Creating an export session
//
//   - [IAVAssetExportSession.InitWithAssetPresetName]: Creates an export session with a preset configuration.
//
// # Accessing export presets
//
//   - [IAVAssetExportSession.PresetName]: The name of the preset that the asset export session uses.
//   - [IAVAssetExportSession.DetermineCompatibleFileTypesWithCompletionHandler]: Determines the output file types an asset export session supports writing in its current configuration.
//
// # Configuring output
//
//   - [IAVAssetExportSession.SupportedFileTypes]: An array containing the types of files the session can write.
//   - [IAVAssetExportSession.AllowsParallelizedExport]: A Boolean value that indicates whether the session can parallelize its export operation.
//   - [IAVAssetExportSession.SetAllowsParallelizedExport]
//   - [IAVAssetExportSession.ShouldOptimizeForNetworkUse]: A Boolean value that indicates whether to optimize the movie for network use.
//   - [IAVAssetExportSession.SetShouldOptimizeForNetworkUse]
//   - [IAVAssetExportSession.CanPerformMultiplePassesOverSourceMediaData]: A Boolean value that indicates whether the export session can perform multiple passes over the source media to achieve better results.
//   - [IAVAssetExportSession.SetCanPerformMultiplePassesOverSourceMediaData]
//   - [IAVAssetExportSession.TimeRange]: The time range of the source asset to export.
//   - [IAVAssetExportSession.SetTimeRange]
//   - [IAVAssetExportSession.FileLengthLimit]: The file length that the output of the session must not exceed.
//   - [IAVAssetExportSession.SetFileLengthLimit]
//   - [IAVAssetExportSession.DirectoryForTemporaryFiles]: A directory suitable to store temporary files that the export process generates.
//   - [IAVAssetExportSession.SetDirectoryForTemporaryFiles]
//
// # Configuring metadata
//
//   - [IAVAssetExportSession.Metadata]: The metadata an export session writes to the output container file.
//   - [IAVAssetExportSession.SetMetadata]
//   - [IAVAssetExportSession.MetadataItemFilter]: An object the export session uses to filter the metadata items it transfers to the output asset.
//   - [IAVAssetExportSession.SetMetadataItemFilter]
//
// # Configuring video output
//
//   - [IAVAssetExportSession.VideoComposition]: An optional object that provides instructions for how to composite frames of video.
//   - [IAVAssetExportSession.SetVideoComposition]
//   - [IAVAssetExportSession.CustomVideoCompositor]: An optional custom object to use when compositing video frames.
//
// # Configuring track groups
//
//   - [IAVAssetExportSession.AudioTrackGroupHandling]: A policy that defines how the session exports alternate audio tracks.
//   - [IAVAssetExportSession.SetAudioTrackGroupHandling]
//
// # Configuring audio output
//
//   - [IAVAssetExportSession.AudioMix]: The parameters for audio mixing and an indication of whether to enable nondefault audio mixing for export.
//   - [IAVAssetExportSession.SetAudioMix]
//   - [IAVAssetExportSession.AudioTimePitchAlgorithm]: A processing algorithm for managing audio pitch for scaled audio edits.
//   - [IAVAssetExportSession.SetAudioTimePitchAlgorithm]
//
// # Monitoring export progress
//
//   - [IAVAssetExportSession.Status]: The status of the export session.
//   - [IAVAssetExportSession.Error]: An optional error object.
//
// # Estimating file length and duration
//
//   - [IAVAssetExportSession.EstimateOutputFileLengthWithCompletionHandler]: Starts estimating the output file length of the export while considering the asset, preset, and time range configuration of the export session.
//   - [IAVAssetExportSession.EstimatedOutputFileLength]: The estimated length of the exported file, in bytes.
//
// # Estimating duration
//
//   - [IAVAssetExportSession.EstimateMaximumDurationWithCompletionHandler]: Starts estimating the maximum duration of the export while considering the asset, preset, and time range configuration of the export session.
//   - [IAVAssetExportSession.MaxDuration]: Provides an estimate of the maximum duration of the exported media.
//
// # Accessing the asset
//
//   - [IAVAssetExportSession.Asset]: An asset that a session exports.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetExportSession
type IAVAssetExportSession interface {
	objectivec.IObject

	// Topic: Creating an export session

	// Creates an export session with a preset configuration.
	InitWithAssetPresetName(asset IAVAsset, presetName string) AVAssetExportSession

	// Topic: Accessing export presets

	// The name of the preset that the asset export session uses.
	PresetName() string
	// Determines the output file types an asset export session supports writing in its current configuration.
	DetermineCompatibleFileTypesWithCompletionHandler(handler VoidHandler)

	// Topic: Configuring output

	// An array containing the types of files the session can write.
	SupportedFileTypes() []string
	// A Boolean value that indicates whether the session can parallelize its export operation.
	AllowsParallelizedExport() bool
	SetAllowsParallelizedExport(value bool)
	// A Boolean value that indicates whether to optimize the movie for network use.
	ShouldOptimizeForNetworkUse() bool
	SetShouldOptimizeForNetworkUse(value bool)
	// A Boolean value that indicates whether the export session can perform multiple passes over the source media to achieve better results.
	CanPerformMultiplePassesOverSourceMediaData() bool
	SetCanPerformMultiplePassesOverSourceMediaData(value bool)
	// The time range of the source asset to export.
	TimeRange() uintptr
	SetTimeRange(value uintptr)
	// The file length that the output of the session must not exceed.
	FileLengthLimit() int64
	SetFileLengthLimit(value int64)
	// A directory suitable to store temporary files that the export process generates.
	DirectoryForTemporaryFiles() foundation.INSURL
	SetDirectoryForTemporaryFiles(value foundation.INSURL)

	// Topic: Configuring metadata

	// The metadata an export session writes to the output container file.
	Metadata() []AVMetadataItem
	SetMetadata(value []AVMetadataItem)
	// An object the export session uses to filter the metadata items it transfers to the output asset.
	MetadataItemFilter() IAVMetadataItemFilter
	SetMetadataItemFilter(value IAVMetadataItemFilter)

	// Topic: Configuring video output

	// An optional object that provides instructions for how to composite frames of video.
	VideoComposition() IAVVideoComposition
	SetVideoComposition(value IAVVideoComposition)
	// An optional custom object to use when compositing video frames.
	CustomVideoCompositor() AVVideoCompositing

	// Topic: Configuring track groups

	// A policy that defines how the session exports alternate audio tracks.
	AudioTrackGroupHandling() AVAssetTrackGroupOutputHandling
	SetAudioTrackGroupHandling(value AVAssetTrackGroupOutputHandling)

	// Topic: Configuring audio output

	// The parameters for audio mixing and an indication of whether to enable nondefault audio mixing for export.
	AudioMix() IAVAudioMix
	SetAudioMix(value IAVAudioMix)
	// A processing algorithm for managing audio pitch for scaled audio edits.
	AudioTimePitchAlgorithm() AVAudioTimePitchAlgorithm
	SetAudioTimePitchAlgorithm(value AVAudioTimePitchAlgorithm)

	// Topic: Monitoring export progress

	// The status of the export session.
	Status() AVAssetExportSessionStatus
	// An optional error object.
	Error() foundation.INSError

	// Topic: Estimating file length and duration

	// Starts estimating the output file length of the export while considering the asset, preset, and time range configuration of the export session.
	EstimateOutputFileLengthWithCompletionHandler(handler int64_tErrorHandler)
	// The estimated length of the exported file, in bytes.
	EstimatedOutputFileLength() int64

	// Topic: Estimating duration

	// Starts estimating the maximum duration of the export while considering the asset, preset, and time range configuration of the export session.
	EstimateMaximumDurationWithCompletionHandler(handler CMTimeErrorHandler)
	// Provides an estimate of the maximum duration of the exported media.
	MaxDuration() uintptr

	// Topic: Accessing the asset

	// An asset that a session exports.
	Asset() IAVAsset
}

// Init initializes the instance.
func (a AVAssetExportSession) Init() AVAssetExportSession {
	rv := objc.Send[AVAssetExportSession](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AVAssetExportSession) Autorelease() AVAssetExportSession {
	rv := objc.Send[AVAssetExportSession](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVAssetExportSession creates a new AVAssetExportSession instance.
func NewAVAssetExportSession() AVAssetExportSession {
	class := getAVAssetExportSessionClass()
	rv := objc.Send[AVAssetExportSession](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates an export session with a preset configuration.
//
// asset: The asset to export.
//
// presetName: A string constant that specifies the preset template for the export. See
// [Export presets] for available values.
// //
// [Export presets]: https://developer.apple.com/documentation/AVFoundation/export-presets
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetExportSession/init(asset:presetName:)
func NewAssetExportSessionWithAssetPresetName(asset IAVAsset, presetName string) AVAssetExportSession {
	instance := getAVAssetExportSessionClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithAsset:presetName:"), asset, objc.String(presetName))
	return AVAssetExportSessionFromID(rv)
}

// Creates an export session with a preset configuration.
//
// asset: The asset to export.
//
// presetName: A string constant that specifies the preset template for the export. See
// [Export presets] for available values.
// //
// [Export presets]: https://developer.apple.com/documentation/AVFoundation/export-presets
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetExportSession/init(asset:presetName:)
func (a AVAssetExportSession) InitWithAssetPresetName(asset IAVAsset, presetName string) AVAssetExportSession {
	rv := objc.Send[AVAssetExportSession](a.ID, objc.Sel("initWithAsset:presetName:"), asset, objc.String(presetName))
	return rv
}
// Determines the output file types an asset export session supports writing
// in its current configuration.
//
// handler: A callback the system passes an array of [AVFileType] structures when it
// determines the compatible file types.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetExportSession/determineCompatibleFileTypes(completionHandler:)
func (a AVAssetExportSession) DetermineCompatibleFileTypesWithCompletionHandler(handler VoidHandler) {
_block0, _ := NewVoidBlock(handler)
	objc.Send[objc.ID](a.ID, objc.Sel("determineCompatibleFileTypesWithCompletionHandler:"), _block0)
}
// Starts estimating the output file length of the export while considering
// the asset, preset, and time range configuration of the export session.
//
// handler: A callback the system invokes when it finishes its estimation. It passes
// the callback the following parameters:
// 
// `estimatedOutputFileLength`: The system’s estimation of the output file
// length. `error`: An error object if the request fails; otherwise, `nil`.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetExportSession/estimateOutputFileLength(completionHandler:)
func (a AVAssetExportSession) EstimateOutputFileLengthWithCompletionHandler(handler int64_tErrorHandler) {
_block0, _ := Newint64_tErrorBlock(handler)
	objc.Send[objc.ID](a.ID, objc.Sel("estimateOutputFileLengthWithCompletionHandler:"), _block0)
}
// Starts estimating the maximum duration of the export while considering the
// asset, preset, and time range configuration of the export session.
//
// handler: A callback the system invokes when it finishes its estimation. It passes
// the callback the following parameters:
// 
// `estimatedMaximumDuration`: The system’s estimation of the maximum
// duration. `error`: An optional error object that indicates if an error
// occurred during processing.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetExportSession/estimateMaximumDuration(completionHandler:)
func (a AVAssetExportSession) EstimateMaximumDurationWithCompletionHandler(handler CMTimeErrorHandler) {
_block0, _ := NewCMTimeErrorBlock(handler)
	objc.Send[objc.ID](a.ID, objc.Sel("estimateMaximumDurationWithCompletionHandler:"), _block0)
}

// Returns all available export preset names.
//
// # Return Value
// 
// See [Export presets] for values an asset export session supports.
//
// [Export presets]: https://developer.apple.com/documentation/AVFoundation/export-presets
//
// # Discussion
// 
// Not all presets are compatible with all assets.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetExportSession/allExportPresets()
func (_AVAssetExportSessionClass AVAssetExportSessionClass) AllExportPresets() []string {
	rv := objc.Send[[]objc.ID](objc.ID(_AVAssetExportSessionClass.class), objc.Sel("allExportPresets"))
	return objc.ConvertSliceToStrings(rv)
}
// Determines an export preset’s compatibility to export the asset in a
// container of the output file type.
//
// presetName: The name of the preset whose compatibility you want to test. See [Export
// presets] for preset values an asset export session supports.
// //
// [Export presets]: https://developer.apple.com/documentation/AVFoundation/export-presets
//
// asset: The asset to export.
//
// outputFileType: The file type of the output container.
//
// handler: A callback the system passes a Boolean result when it determines the
// compatibility of the preset.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetExportSession/determineCompatibility(ofExportPreset:with:outputFileType:completionHandler:)
func (_AVAssetExportSessionClass AVAssetExportSessionClass) DetermineCompatibilityOfExportPresetWithAssetOutputFileTypeCompletionHandler(presetName string, asset IAVAsset, outputFileType AVFileType, handler BoolHandler) {
_block3, _ := NewBoolBlock(handler)
	objc.Send[objc.ID](objc.ID(_AVAssetExportSessionClass.class), objc.Sel("determineCompatibilityOfExportPreset:withAsset:outputFileType:completionHandler:"), objc.String(presetName), asset, outputFileType, _block3)
}
// Returns a new asset export session that uses the specified preset.
//
// asset: The asset to export.
//
// presetName: A string constant that specifies the preset template for the export. See
// [Export presets] for values an asset export session supports.
// //
// [Export presets]: https://developer.apple.com/documentation/AVFoundation/export-presets
//
// # Return Value
// 
// An asset export session.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetExportSession/exportSessionWithAsset:presetName:
func (_AVAssetExportSessionClass AVAssetExportSessionClass) ExportSessionWithAssetPresetName(asset IAVAsset, presetName string) AVAssetExportSession {
	rv := objc.Send[objc.ID](objc.ID(_AVAssetExportSessionClass.class), objc.Sel("exportSessionWithAsset:presetName:"), asset, objc.String(presetName))
	return AVAssetExportSessionFromID(rv)
}

// The name of the preset that the asset export session uses.
//
// # Discussion
// 
// See [Export presets] for values an asset export session supports.
// 
// This property is key-value observable.
//
// [Export presets]: https://developer.apple.com/documentation/AVFoundation/export-presets
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetExportSession/presetName
func (a AVAssetExportSession) PresetName() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("presetName"))
	return foundation.NSStringFromID(rv).String()
}
// An array containing the types of files the session can write.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetExportSession/supportedFileTypes
func (a AVAssetExportSession) SupportedFileTypes() []string {
	rv := objc.Send[[]objc.ID](a.ID, objc.Sel("supportedFileTypes"))
	return objc.ConvertSliceToStrings(rv)
}
// A Boolean value that indicates whether the session can parallelize its
// export operation.
//
// # Discussion
// 
// This value is [true] by default, which indicates that the export session is
// allowed to expedite its processing by using additional resources in
// parallel on select Mac systems. If parallelization isn’t achievable,
// export proceeds as normal.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetExportSession/allowsParallelizedExport
func (a AVAssetExportSession) AllowsParallelizedExport() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("allowsParallelizedExport"))
	return rv
}
func (a AVAssetExportSession) SetAllowsParallelizedExport(value bool) {
	objc.Send[struct{}](a.ID, objc.Sel("setAllowsParallelizedExport:"), value)
}
// A Boolean value that indicates whether to optimize the movie for network
// use.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetExportSession/shouldOptimizeForNetworkUse
func (a AVAssetExportSession) ShouldOptimizeForNetworkUse() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("shouldOptimizeForNetworkUse"))
	return rv
}
func (a AVAssetExportSession) SetShouldOptimizeForNetworkUse(value bool) {
	objc.Send[struct{}](a.ID, objc.Sel("setShouldOptimizeForNetworkUse:"), value)
}
// A Boolean value that indicates whether the export session can perform
// multiple passes over the source media to achieve better results.
//
// # Discussion
// 
// When the value for this property is [true], the export session can produce
// higher quality results at the expense of longer export times. Setting this
// property to [true] may also require the export session to write temporary
// data to disk during the export. To control the location of temporary data,
// use the property [DirectoryForTemporaryFiles].
// 
// The default value is [false]. Not all export session configurations can
// benefit from performing multiple passes over the source media. In these
// cases, setting this property to [true] has no effect.
// 
// You can’t set this property after the export starts.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetExportSession/canPerformMultiplePassesOverSourceMediaData
func (a AVAssetExportSession) CanPerformMultiplePassesOverSourceMediaData() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("canPerformMultiplePassesOverSourceMediaData"))
	return rv
}
func (a AVAssetExportSession) SetCanPerformMultiplePassesOverSourceMediaData(value bool) {
	objc.Send[struct{}](a.ID, objc.Sel("setCanPerformMultiplePassesOverSourceMediaData:"), value)
}
// The time range of the source asset to export.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetExportSession/timeRange
func (a AVAssetExportSession) TimeRange() uintptr {
	rv := objc.Send[uintptr](a.ID, objc.Sel("timeRange"))
	return rv
}
func (a AVAssetExportSession) SetTimeRange(value uintptr) {
	objc.Send[struct{}](a.ID, objc.Sel("setTimeRange:"), value)
}
// The file length that the output of the session must not exceed.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetExportSession/fileLengthLimit
func (a AVAssetExportSession) FileLengthLimit() int64 {
	rv := objc.Send[int64](a.ID, objc.Sel("fileLengthLimit"))
	return rv
}
func (a AVAssetExportSession) SetFileLengthLimit(value int64) {
	objc.Send[struct{}](a.ID, objc.Sel("setFileLengthLimit:"), value)
}
// A directory suitable to store temporary files that the export process
// generates.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetExportSession/directoryForTemporaryFiles
func (a AVAssetExportSession) DirectoryForTemporaryFiles() foundation.INSURL {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("directoryForTemporaryFiles"))
	return foundation.NSURLFromID(objc.ID(rv))
}
func (a AVAssetExportSession) SetDirectoryForTemporaryFiles(value foundation.INSURL) {
	objc.Send[struct{}](a.ID, objc.Sel("setDirectoryForTemporaryFiles:"), value)
}
// The metadata an export session writes to the output container file.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetExportSession/metadata
func (a AVAssetExportSession) Metadata() []AVMetadataItem {
	rv := objc.Send[[]objc.ID](a.ID, objc.Sel("metadata"))
	return objc.ConvertSlice(rv, func(id objc.ID) AVMetadataItem {
		return AVMetadataItemFromID(id)
	})
}
func (a AVAssetExportSession) SetMetadata(value []AVMetadataItem) {
	objc.Send[struct{}](a.ID, objc.Sel("setMetadata:"), objectivec.IObjectSliceToNSArray(value))
}
// An object the export session uses to filter the metadata items it transfers
// to the output asset.
//
// # Discussion
// 
// The default value is `nil`.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetExportSession/metadataItemFilter
func (a AVAssetExportSession) MetadataItemFilter() IAVMetadataItemFilter {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("metadataItemFilter"))
	return AVMetadataItemFilterFromID(objc.ID(rv))
}
func (a AVAssetExportSession) SetMetadataItemFilter(value IAVMetadataItemFilter) {
	objc.Send[struct{}](a.ID, objc.Sel("setMetadataItemFilter:"), value)
}
// An optional object that provides instructions for how to composite frames
// of video.
//
// # Discussion
// 
// The default value is `nil`.
// 
// This property is key-value observable.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetExportSession/videoComposition
func (a AVAssetExportSession) VideoComposition() IAVVideoComposition {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("videoComposition"))
	return AVVideoCompositionFromID(objc.ID(rv))
}
func (a AVAssetExportSession) SetVideoComposition(value IAVVideoComposition) {
	objc.Send[struct{}](a.ID, objc.Sel("setVideoComposition:"), value)
}
// An optional custom object to use when compositing video frames.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetExportSession/customVideoCompositor
func (a AVAssetExportSession) CustomVideoCompositor() AVVideoCompositing {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("customVideoCompositor"))
	return AVVideoCompositingObjectFromID(rv)
}
// A policy that defines how the session exports alternate audio tracks.
//
// # Discussion
// 
// By default, a session exports only the enabled audio tracks within an
// alternate track group from the source asset. You can specify that the
// session preserve all audio tracks within an alternate track group by
// setting this value to
// [AssetTrackGroupOutputHandlingPreserveAlternateTracks].
// 
// If no audio alternate track group is present, the value of this property
// has no effect. You can query the [trackGroups] property of [AVAsset] to
// determine whether it contains audio track groups.
//
// [trackGroups]: https://developer.apple.com/documentation/AVFoundation/AVPartialAsyncProperty/trackGroups
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetExportSession/audioTrackGroupHandling
func (a AVAssetExportSession) AudioTrackGroupHandling() AVAssetTrackGroupOutputHandling {
	rv := objc.Send[AVAssetTrackGroupOutputHandling](a.ID, objc.Sel("audioTrackGroupHandling"))
	return AVAssetTrackGroupOutputHandling(rv)
}
func (a AVAssetExportSession) SetAudioTrackGroupHandling(value AVAssetTrackGroupOutputHandling) {
	objc.Send[struct{}](a.ID, objc.Sel("setAudioTrackGroupHandling:"), value)
}
// The parameters for audio mixing and an indication of whether to enable
// nondefault audio mixing for export.
//
// # Discussion
// 
// This value is key-value observable.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetExportSession/audioMix
func (a AVAssetExportSession) AudioMix() IAVAudioMix {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("audioMix"))
	return AVAudioMixFromID(objc.ID(rv))
}
func (a AVAssetExportSession) SetAudioMix(value IAVAudioMix) {
	objc.Send[struct{}](a.ID, objc.Sel("setAudioMix:"), value)
}
// A processing algorithm for managing audio pitch for scaled audio edits.
//
// # Discussion
// 
// The default value is [spectral].
//
// [spectral]: https://developer.apple.com/documentation/AVFoundation/AVAudioTimePitchAlgorithm/spectral
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetExportSession/audioTimePitchAlgorithm
func (a AVAssetExportSession) AudioTimePitchAlgorithm() AVAudioTimePitchAlgorithm {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("audioTimePitchAlgorithm"))
	return AVAudioTimePitchAlgorithm(foundation.NSStringFromID(rv).String())
}
func (a AVAssetExportSession) SetAudioTimePitchAlgorithm(value AVAudioTimePitchAlgorithm) {
	objc.Send[struct{}](a.ID, objc.Sel("setAudioTimePitchAlgorithm:"), objc.String(string(value)))
}
// The status of the export session.
//
// # Discussion
// 
// For possible values, see [AVAssetExportSession.Status].
// 
// This value is key-value observable.
//
// [AVAssetExportSession.Status]: https://developer.apple.com/documentation/AVFoundation/AVAssetExportSession/Status-swift.enum
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetExportSession/status-swift.property
func (a AVAssetExportSession) Status() AVAssetExportSessionStatus {
	rv := objc.Send[AVAssetExportSessionStatus](a.ID, objc.Sel("status"))
	return AVAssetExportSessionStatus(rv)
}
// An optional error object.
//
// # Discussion
// 
// The default value of this property is `nil`. The export session sets it to
// an error object if its status changes to [AssetExportSessionStatusFailed]
// or [AssetExportSessionStatusCancelled].
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetExportSession/error
func (a AVAssetExportSession) Error() foundation.INSError {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("error"))
	return foundation.NSErrorFromID(objc.ID(rv))
}
// The estimated length of the exported file, in bytes.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetExportSession/estimatedOutputFileLength
func (a AVAssetExportSession) EstimatedOutputFileLength() int64 {
	rv := objc.Send[int64](a.ID, objc.Sel("estimatedOutputFileLength"))
	return rv
}
// Provides an estimate of the maximum duration of the exported media.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetExportSession/maxDuration
func (a AVAssetExportSession) MaxDuration() uintptr {
	rv := objc.Send[uintptr](a.ID, objc.Sel("maxDuration"))
	return rv
}
// An asset that a session exports.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetExportSession/asset
func (a AVAssetExportSession) Asset() IAVAsset {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("asset"))
	return AVAssetFromID(objc.ID(rv))
}

// DetermineCompatibleFileTypes is a synchronous wrapper around [AVAssetExportSession.DetermineCompatibleFileTypesWithCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (a AVAssetExportSession) DetermineCompatibleFileTypes(ctx context.Context) error {
	done := make(chan struct{}, 1)
	a.DetermineCompatibleFileTypesWithCompletionHandler(func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// DetermineCompatibilityOfExportPresetWithAssetOutputFileType is a synchronous wrapper around [AVAssetExportSession.DetermineCompatibilityOfExportPresetWithAssetOutputFileTypeCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (ac AVAssetExportSessionClass) DetermineCompatibilityOfExportPresetWithAssetOutputFileType(ctx context.Context, presetName string, asset IAVAsset, outputFileType AVFileType) (bool, error) {
	done := make(chan bool, 1)
	ac.DetermineCompatibilityOfExportPresetWithAssetOutputFileTypeCompletionHandler(presetName, asset, outputFileType, func(val bool) {
		done <- val
	})
	select {
	case r := <-done:
		return r, nil
	case <-ctx.Done():
		return false, ctx.Err()
	}
}

// EstimateOutputFileLength is a synchronous wrapper around [AVAssetExportSession.EstimateOutputFileLengthWithCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (a AVAssetExportSession) EstimateOutputFileLength(ctx context.Context) (int64, error) {
	type result struct {
		val int64
		err error
	}
	done := make(chan result, 1)
	a.EstimateOutputFileLengthWithCompletionHandler(func(val int64, err error) {
		done <- result{val, err}
	})
	select {
	case r := <-done:
		return r.val, r.err
	case <-ctx.Done():
		return 0, ctx.Err()
	}
}

// EstimateMaximumDuration is a synchronous wrapper around [AVAssetExportSession.EstimateMaximumDurationWithCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (a AVAssetExportSession) EstimateMaximumDuration(ctx context.Context) (uintptr, error) {
	type result struct {
		val uintptr
		err error
	}
	done := make(chan result, 1)
	a.EstimateMaximumDurationWithCompletionHandler(func(val uintptr, err error) {
		done <- result{val, err}
	})
	select {
	case r := <-done:
		return r.val, r.err
	case <-ctx.Done():
		return 0, ctx.Err()
	}
}

