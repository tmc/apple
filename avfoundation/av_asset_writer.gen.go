// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"context"
	"sync"
	"unsafe"

	"github.com/tmc/apple/coremedia"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
	"github.com/tmc/apple/uniformtypeidentifiers"
)

// The class instance for the [AVAssetWriter] class.
var (
	_AVAssetWriterClass     AVAssetWriterClass
	_AVAssetWriterClassOnce sync.Once
)

func getAVAssetWriterClass() AVAssetWriterClass {
	_AVAssetWriterClassOnce.Do(func() {
		_AVAssetWriterClass = AVAssetWriterClass{class: objc.GetClass("AVAssetWriter")}
	})
	return _AVAssetWriterClass
}

// GetAVAssetWriterClass returns the class object for AVAssetWriter.
func GetAVAssetWriterClass() AVAssetWriterClass {
	return getAVAssetWriterClass()
}

type AVAssetWriterClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVAssetWriterClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVAssetWriterClass) Alloc() AVAssetWriter {
	rv := objc.Send[AVAssetWriter](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that writes media data to a container file.
//
// # Overview
//
// You use an asset writer to write media to file formats such as the
// QuickTime movie file format and MPEG-4 file format. An asset writer
// automatically supports interleaving media data from concurrent tracks for
// efficient playback and storage. It can reencode media samples it writes to
// the output file, and may also write collections of metadata to the output
// file.
//
// # Creating an asset writer
//
//   - [AVAssetWriter.InitWithURLFileTypeError]: Creates an object that writes media data to a container file at the output URL.
//   - [AVAssetWriter.InitWithContentType]: Creates an object that outputs segment data in a specified container format.
//
// # Configuring inputs
//
//   - [AVAssetWriter.Inputs]: The inputs an asset writer contains.
//   - [AVAssetWriter.AvailableMediaTypes]: The media types the asset writer supports adding as inputs.
//   - [AVAssetWriter.CanApplyOutputSettingsForMediaType]: Determines whether the output file format supports the output settings for a specific media type.
//   - [AVAssetWriter.CanAddInput]: Determines whether the asset writer supports adding the input.
//
// # Configuring input groups
//
//   - [AVAssetWriter.InputGroups]: The input groups an asset writer contains.
//   - [AVAssetWriter.CanAddInputGroup]: Determines whether the asset writer supports adding the input group.
//   - [AVAssetWriter.AddInputGroup]: Adds an input group to an asset writer.
//
// # Configuring output
//
//   - [AVAssetWriter.Metadata]: An array of metadata items to write to the output file.
//   - [AVAssetWriter.SetMetadata]
//   - [AVAssetWriter.ShouldOptimizeForNetworkUse]: A Boolean value that indicates whether to write the output file to make it more suitable for playback over a network.
//   - [AVAssetWriter.SetShouldOptimizeForNetworkUse]
//   - [AVAssetWriter.DirectoryForTemporaryFiles]: A directory to contain temporary files that the export process generates.
//   - [AVAssetWriter.SetDirectoryForTemporaryFiles]
//
// # Configuring fragment output
//
//   - [AVAssetWriter.MovieFragmentInterval]: The interval at which to write movie fragments.
//   - [AVAssetWriter.SetMovieFragmentInterval]
//   - [AVAssetWriter.InitialMovieFragmentInterval]: The interval at which to write the initial movie fragment.
//   - [AVAssetWriter.SetInitialMovieFragmentInterval]
//   - [AVAssetWriter.InitialMovieFragmentSequenceNumber]: The sequence number of the initial movie fragment.
//   - [AVAssetWriter.SetInitialMovieFragmentSequenceNumber]
//   - [AVAssetWriter.ProducesCombinableFragments]: A Boolean value that indicates whether the asset writer outputs movie fragments suitable for combining with others.
//   - [AVAssetWriter.SetProducesCombinableFragments]
//   - [AVAssetWriter.OverallDurationHint]: A hint of the final duration of the output file.
//   - [AVAssetWriter.SetOverallDurationHint]
//   - [AVAssetWriter.MovieTimeScale]: The time scale of the movie.
//   - [AVAssetWriter.SetMovieTimeScale]
//
// # Managing writing sessions
//
//   - [AVAssetWriter.StartSessionAtSourceTime]: Starts an asset-writing session.
//   - [AVAssetWriter.EndSessionAtSourceTime]: Finishes an asset-writing session.
//   - [AVAssetWriter.FinishWritingWithCompletionHandler]: Marks all unfinished inputs as finished and completes the writing of the output file.
//   - [AVAssetWriter.CancelWriting]: Cancels the creation of the output file.
//
// # Inspecting writing status
//
//   - [AVAssetWriter.Status]: The status of writing samples to the output file.
//   - [AVAssetWriter.Error]: An error object that describes an asset-writing failure.
//
// # Configuring segment writing
//
//   - [AVAssetWriter.Delegate]: A delegate object that responds to asset-writing events.
//   - [AVAssetWriter.SetDelegate]
//   - [AVAssetWriter.PreferredOutputSegmentInterval]: The interval of output segments that you prefer.
//   - [AVAssetWriter.SetPreferredOutputSegmentInterval]
//   - [AVAssetWriter.InitialSegmentStartTime]: The start time of the initial segment.
//   - [AVAssetWriter.SetInitialSegmentStartTime]
//   - [AVAssetWriter.OutputFileTypeProfile]: A profile for the output file type.
//   - [AVAssetWriter.SetOutputFileTypeProfile]
//   - [AVAssetWriter.FlushSegment]: Closes the current segment and outputs it to a delegate method.
//
// # Accessing output settings
//
//   - [AVAssetWriter.OutputURL]: The location of the container file that the writer outputs.
//   - [AVAssetWriter.OutputFileType]: The type of container file that the writer outputs.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter
type AVAssetWriter struct {
	objectivec.Object
}

// AVAssetWriterFromID constructs a [AVAssetWriter] from an objc.ID.
//
// An object that writes media data to a container file.
func AVAssetWriterFromID(id objc.ID) AVAssetWriter {
	return AVAssetWriter{objectivec.Object{ID: id}}
}

// NOTE: AVAssetWriter adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVAssetWriter] class.
//
// # Creating an asset writer
//
//   - [IAVAssetWriter.InitWithURLFileTypeError]: Creates an object that writes media data to a container file at the output URL.
//   - [IAVAssetWriter.InitWithContentType]: Creates an object that outputs segment data in a specified container format.
//
// # Configuring inputs
//
//   - [IAVAssetWriter.Inputs]: The inputs an asset writer contains.
//   - [IAVAssetWriter.AvailableMediaTypes]: The media types the asset writer supports adding as inputs.
//   - [IAVAssetWriter.CanApplyOutputSettingsForMediaType]: Determines whether the output file format supports the output settings for a specific media type.
//   - [IAVAssetWriter.CanAddInput]: Determines whether the asset writer supports adding the input.
//
// # Configuring input groups
//
//   - [IAVAssetWriter.InputGroups]: The input groups an asset writer contains.
//   - [IAVAssetWriter.CanAddInputGroup]: Determines whether the asset writer supports adding the input group.
//   - [IAVAssetWriter.AddInputGroup]: Adds an input group to an asset writer.
//
// # Configuring output
//
//   - [IAVAssetWriter.Metadata]: An array of metadata items to write to the output file.
//   - [IAVAssetWriter.SetMetadata]
//   - [IAVAssetWriter.ShouldOptimizeForNetworkUse]: A Boolean value that indicates whether to write the output file to make it more suitable for playback over a network.
//   - [IAVAssetWriter.SetShouldOptimizeForNetworkUse]
//   - [IAVAssetWriter.DirectoryForTemporaryFiles]: A directory to contain temporary files that the export process generates.
//   - [IAVAssetWriter.SetDirectoryForTemporaryFiles]
//
// # Configuring fragment output
//
//   - [IAVAssetWriter.MovieFragmentInterval]: The interval at which to write movie fragments.
//   - [IAVAssetWriter.SetMovieFragmentInterval]
//   - [IAVAssetWriter.InitialMovieFragmentInterval]: The interval at which to write the initial movie fragment.
//   - [IAVAssetWriter.SetInitialMovieFragmentInterval]
//   - [IAVAssetWriter.InitialMovieFragmentSequenceNumber]: The sequence number of the initial movie fragment.
//   - [IAVAssetWriter.SetInitialMovieFragmentSequenceNumber]
//   - [IAVAssetWriter.ProducesCombinableFragments]: A Boolean value that indicates whether the asset writer outputs movie fragments suitable for combining with others.
//   - [IAVAssetWriter.SetProducesCombinableFragments]
//   - [IAVAssetWriter.OverallDurationHint]: A hint of the final duration of the output file.
//   - [IAVAssetWriter.SetOverallDurationHint]
//   - [IAVAssetWriter.MovieTimeScale]: The time scale of the movie.
//   - [IAVAssetWriter.SetMovieTimeScale]
//
// # Managing writing sessions
//
//   - [IAVAssetWriter.StartSessionAtSourceTime]: Starts an asset-writing session.
//   - [IAVAssetWriter.EndSessionAtSourceTime]: Finishes an asset-writing session.
//   - [IAVAssetWriter.FinishWritingWithCompletionHandler]: Marks all unfinished inputs as finished and completes the writing of the output file.
//   - [IAVAssetWriter.CancelWriting]: Cancels the creation of the output file.
//
// # Inspecting writing status
//
//   - [IAVAssetWriter.Status]: The status of writing samples to the output file.
//   - [IAVAssetWriter.Error]: An error object that describes an asset-writing failure.
//
// # Configuring segment writing
//
//   - [IAVAssetWriter.Delegate]: A delegate object that responds to asset-writing events.
//   - [IAVAssetWriter.SetDelegate]
//   - [IAVAssetWriter.PreferredOutputSegmentInterval]: The interval of output segments that you prefer.
//   - [IAVAssetWriter.SetPreferredOutputSegmentInterval]
//   - [IAVAssetWriter.InitialSegmentStartTime]: The start time of the initial segment.
//   - [IAVAssetWriter.SetInitialSegmentStartTime]
//   - [IAVAssetWriter.OutputFileTypeProfile]: A profile for the output file type.
//   - [IAVAssetWriter.SetOutputFileTypeProfile]
//   - [IAVAssetWriter.FlushSegment]: Closes the current segment and outputs it to a delegate method.
//
// # Accessing output settings
//
//   - [IAVAssetWriter.OutputURL]: The location of the container file that the writer outputs.
//   - [IAVAssetWriter.OutputFileType]: The type of container file that the writer outputs.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter
type IAVAssetWriter interface {
	objectivec.IObject

	// Topic: Creating an asset writer

	// Creates an object that writes media data to a container file at the output URL.
	InitWithURLFileTypeError(outputURL foundation.INSURL, outputFileType AVFileType) (AVAssetWriter, error)
	// Creates an object that outputs segment data in a specified container format.
	InitWithContentType(outputContentType uniformtypeidentifiers.UTType) AVAssetWriter

	// Topic: Configuring inputs

	// The inputs an asset writer contains.
	Inputs() []AVAssetWriterInput
	// The media types the asset writer supports adding as inputs.
	AvailableMediaTypes() []string
	// Determines whether the output file format supports the output settings for a specific media type.
	CanApplyOutputSettingsForMediaType(outputSettings foundation.INSDictionary, mediaType AVMediaType) bool
	// Determines whether the asset writer supports adding the input.
	CanAddInput(input IAVAssetWriterInput) bool

	// Topic: Configuring input groups

	// The input groups an asset writer contains.
	InputGroups() []AVAssetWriterInputGroup
	// Determines whether the asset writer supports adding the input group.
	CanAddInputGroup(inputGroup IAVAssetWriterInputGroup) bool
	// Adds an input group to an asset writer.
	AddInputGroup(inputGroup IAVAssetWriterInputGroup)

	// Topic: Configuring output

	// An array of metadata items to write to the output file.
	Metadata() []AVMetadataItem
	SetMetadata(value []AVMetadataItem)
	// A Boolean value that indicates whether to write the output file to make it more suitable for playback over a network.
	ShouldOptimizeForNetworkUse() bool
	SetShouldOptimizeForNetworkUse(value bool)
	// A directory to contain temporary files that the export process generates.
	DirectoryForTemporaryFiles() foundation.INSURL
	SetDirectoryForTemporaryFiles(value foundation.INSURL)

	// Topic: Configuring fragment output

	// The interval at which to write movie fragments.
	MovieFragmentInterval() coremedia.CMTime
	SetMovieFragmentInterval(value coremedia.CMTime)
	// The interval at which to write the initial movie fragment.
	InitialMovieFragmentInterval() coremedia.CMTime
	SetInitialMovieFragmentInterval(value coremedia.CMTime)
	// The sequence number of the initial movie fragment.
	InitialMovieFragmentSequenceNumber() int
	SetInitialMovieFragmentSequenceNumber(value int)
	// A Boolean value that indicates whether the asset writer outputs movie fragments suitable for combining with others.
	ProducesCombinableFragments() bool
	SetProducesCombinableFragments(value bool)
	// A hint of the final duration of the output file.
	OverallDurationHint() coremedia.CMTime
	SetOverallDurationHint(value coremedia.CMTime)
	// The time scale of the movie.
	MovieTimeScale() int32
	SetMovieTimeScale(value int32)

	// Topic: Managing writing sessions

	// Starts an asset-writing session.
	StartSessionAtSourceTime(startTime coremedia.CMTime)
	// Finishes an asset-writing session.
	EndSessionAtSourceTime(endTime coremedia.CMTime)
	// Marks all unfinished inputs as finished and completes the writing of the output file.
	FinishWritingWithCompletionHandler(handler VoidHandler)
	// Cancels the creation of the output file.
	CancelWriting()

	// Topic: Inspecting writing status

	// The status of writing samples to the output file.
	Status() AVAssetWriterStatus
	// An error object that describes an asset-writing failure.
	Error() foundation.INSError

	// Topic: Configuring segment writing

	// A delegate object that responds to asset-writing events.
	Delegate() AVAssetWriterDelegate
	SetDelegate(value AVAssetWriterDelegate)
	// The interval of output segments that you prefer.
	PreferredOutputSegmentInterval() coremedia.CMTime
	SetPreferredOutputSegmentInterval(value coremedia.CMTime)
	// The start time of the initial segment.
	InitialSegmentStartTime() coremedia.CMTime
	SetInitialSegmentStartTime(value coremedia.CMTime)
	// A profile for the output file type.
	OutputFileTypeProfile() AVFileTypeProfile
	SetOutputFileTypeProfile(value AVFileTypeProfile)
	// Closes the current segment and outputs it to a delegate method.
	FlushSegment()

	// Topic: Accessing output settings

	// The location of the container file that the writer outputs.
	OutputURL() foundation.INSURL
	// The type of container file that the writer outputs.
	OutputFileType() AVFileType
}

// Init initializes the instance.
func (a AVAssetWriter) Init() AVAssetWriter {
	rv := objc.Send[AVAssetWriter](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AVAssetWriter) Autorelease() AVAssetWriter {
	rv := objc.Send[AVAssetWriter](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVAssetWriter creates a new AVAssetWriter instance.
func NewAVAssetWriter() AVAssetWriter {
	class := getAVAssetWriterClass()
	rv := objc.Send[AVAssetWriter](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates an object that outputs segment data in a specified container
// format.
//
// outputContentType: A type that indicates the format of the segment data to output.
//
// # Discussion
//
// Use this initializer to create an asset writer that outputs segment data to
// an adopter of the [AVAssetWriterDelegate] protocol. For example, you can
// create an asset writer object to write an MPEG-4 file as shown below:
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/init(contentType:)
func NewAssetWriterWithContentType(outputContentType uniformtypeidentifiers.UTType) AVAssetWriter {
	instance := getAVAssetWriterClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithContentType:"), outputContentType)
	return AVAssetWriterFromID(rv)
}

// Creates an object that writes media data to a container file at the output
// URL.
//
// outputURL: The location of the file to write.
//
// outputFileType: The type of container file to write.
//
// # Discussion
//
// Writing fails if a file already exists at the output URL.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/init(outputURL:fileType:)
func NewAssetWriterWithURLFileTypeError(outputURL foundation.INSURL, outputFileType AVFileType) (AVAssetWriter, error) {
	var errorPtr objc.ID
	instance := getAVAssetWriterClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithURL:fileType:error:"), outputURL, objc.String(string(outputFileType)), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return AVAssetWriter{}, foundation.NSErrorFrom(errorPtr)
	}
	return AVAssetWriterFromID(rv), nil
}

// Creates an object that writes media data to a container file at the output
// URL.
//
// outputURL: The location of the file to write.
//
// outputFileType: The type of container file to write.
//
// # Discussion
//
// Writing fails if a file already exists at the output URL.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/init(outputURL:fileType:)
func (a AVAssetWriter) InitWithURLFileTypeError(outputURL foundation.INSURL, outputFileType AVFileType) (AVAssetWriter, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](a.ID, objc.Sel("initWithURL:fileType:error:"), outputURL, objc.String(string(outputFileType)), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return AVAssetWriter{}, foundation.NSErrorFrom(errorPtr)
	}
	return AVAssetWriterFromID(rv), nil

}

// Creates an object that outputs segment data in a specified container
// format.
//
// outputContentType: A type that indicates the format of the segment data to output.
//
// # Discussion
//
// Use this initializer to create an asset writer that outputs segment data to
// an adopter of the [AVAssetWriterDelegate] protocol. For example, you can
// create an asset writer object to write an MPEG-4 file as shown below:
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/init(contentType:)
func (a AVAssetWriter) InitWithContentType(outputContentType uniformtypeidentifiers.UTType) AVAssetWriter {
	rv := objc.Send[AVAssetWriter](a.ID, objc.Sel("initWithContentType:"), outputContentType)
	return rv
}

// Determines whether the output file format supports the output settings for
// a specific media type.
//
// outputSettings: The output settings to validate.
//
// mediaType: The media type to validate the output settings for.
//
// # Return Value
//
// true if the format supports the output settings; otherwise, false.
//
// # Discussion
//
// Use this method to determine the compatibility of output settings for a
// particular media type. For example, video compression settings that specify
// H.264 compression aren’t compatible with file formats that don’t
// contain H.264-compressed video.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/canApply(outputSettings:forMediaType:)
func (a AVAssetWriter) CanApplyOutputSettingsForMediaType(outputSettings foundation.INSDictionary, mediaType AVMediaType) bool {
	rv := objc.Send[bool](a.ID, objc.Sel("canApplyOutputSettings:forMediaType:"), outputSettings, objc.String(string(mediaType)))
	return rv
}

// Determines whether the asset writer supports adding the input.
//
// input: The asset writer input to add.
//
// # Return Value
//
// true if you can add the input to the asset writer; otherwise false.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/canAdd(_:)-6al7j
func (a AVAssetWriter) CanAddInput(input IAVAssetWriterInput) bool {
	rv := objc.Send[bool](a.ID, objc.Sel("canAddInput:"), input)
	return rv
}

// Determines whether the asset writer supports adding the input group.
//
// inputGroup: The asset writer input group to add.
//
// # Return Value
//
// true if you can add the input group to the asset writer; otherwise false.
//
// # Discussion
//
// This method returns false if the asset writer’s output file type
// doesn’t support mutually exclusive relationships among tracks, or if the
// input group contains inputs with media types that you can’t relate.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/canAdd(_:)-8s1oh
func (a AVAssetWriter) CanAddInputGroup(inputGroup IAVAssetWriterInputGroup) bool {
	rv := objc.Send[bool](a.ID, objc.Sel("canAddInputGroup:"), inputGroup)
	return rv
}

// Adds an input group to an asset writer.
//
// inputGroup: A compatible asset writer input group to add.
//
// # Discussion
//
// An asset writer marks tracks associated with grouped inputs as mutually
// exclusive to each other for playback or other processing, if the output
// container format supports mutually exclusive relationships among tracks.
//
// When you add an input group to an asset writer, the system sets the value
// of the default input’s [MarksOutputTrackAsEnabled] property to true and
// sets the values of the group’s other inputs to false.
//
// You can’t add input groups after writing starts.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/add(_:)-3san4
func (a AVAssetWriter) AddInputGroup(inputGroup IAVAssetWriterInputGroup) {
	objc.Send[objc.ID](a.ID, objc.Sel("addInputGroup:"), inputGroup)
}

// Starts an asset-writing session.
//
// startTime: The starting asset time for the sample-writing session, in the timeline of
// the source samples.
//
// # Discussion
//
// You must call this method after you call [startWriting()], but before you
// append sample data to asset writer inputs.
//
// Each writing session has a start time that, where allowed by the file
// format you’re writing, defines the mapping from the timeline of source
// samples to the timeline of the written file. In the case of the QuickTime
// movie file format, the first session begins at movie time `0`, so a sample
// you append with timestamp [T] plays at movie time (`T-startTime`). The
// writer adds samples with timestamps earlier than the start time to the
// output file, but they don’t display during playback. If the earliest
// sample for an input has a timestamp later than the start time, the system
// inserts an empty edit to preserve synchronization between tracks of the
// output asset.
//
// To end a session, call [EndSessionAtSourceTime]or
// [FinishWritingWithCompletionHandler]
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/startSession(atSourceTime:)
//
// [startWriting()]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/startWriting()
func (a AVAssetWriter) StartSessionAtSourceTime(startTime coremedia.CMTime) {
	objc.Send[objc.ID](a.ID, objc.Sel("startSessionAtSourceTime:"), startTime)
}

// Finishes an asset-writing session.
//
// endTime: The ending asset time for the session, in the timeline of the source
// samples.
//
// # Discussion
//
// Call this method to complete a session that you started with
// [StartSessionAtSourceTime].
//
// The end time defines the moment on the timeline of source samples at which
// the session ends. In the case of the QuickTime movie file format, each
// sample-writing session’s start and end time pair corresponds to a period
// of movie time into which a writer inserts samples. The writer adds samples
// with timestamps that are later than the session end time to the written
// file but they aren’t presented during playback. For example, if the first
// session has duration `D1 = endTime - startTime`, the writer inserts it into
// the written file at time `0` through [D1]; the second session would insert
// into the written file at time [D1] through `D1 + D2`, and so on. It’s
// legal to have a session with no samples; this causes the creation of an
// empty edit of the prescribed duration.
//
// If you don’t explicitly call this method, the system invokes it
// automatically when you call [FinishWritingWithCompletionHandler]. In that
// case, the session’s effective end time is the timestamp of the last
// sample you append.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/endSession(atSourceTime:)
func (a AVAssetWriter) EndSessionAtSourceTime(endTime coremedia.CMTime) {
	objc.Send[objc.ID](a.ID, objc.Sel("endSessionAtSourceTime:"), endTime)
}

// Marks all unfinished inputs as finished and completes the writing of the
// output file.
//
// handler: A completion handler the system invokes when it finishes writing. Determine
// the success or failure of the writing session by querying the asset
// writer’s [Status] property value.
//
// # Discussion
//
// To ensure the asset writer finishes writing all samples, call this method
// only after all calls to [append(_:)] or [append(_:withPresentationTime:)]
// return.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/finishWriting(completionHandler:)
//
// [append(_:)]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriterInput/append(_:)
// [append(_:withPresentationTime:)]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriterInputPixelBufferAdaptor/append(_:withPresentationTime:)
func (a AVAssetWriter) FinishWritingWithCompletionHandler(handler VoidHandler) {
	_block0, _ := NewVoidBlock(handler)
	objc.Send[objc.ID](a.ID, objc.Sel("finishWritingWithCompletionHandler:"), _block0)
}

// Cancels the creation of the output file.
//
// # Discussion
//
// If the asset writer is in [AVAssetWriterStatusFailed] or
// [AVAssetWriterStatusCompleted] state, calling this method has no effect.
// Otherwise, invoking it blocks the calling thread until the asset writer
// finishes canceling the writing session.
//
// If the asset writer created an output file during the writing process,
// calling this method deletes the file.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/cancelWriting()
func (a AVAssetWriter) CancelWriting() {
	objc.Send[objc.ID](a.ID, objc.Sel("cancelWriting"))
}

// Closes the current segment and outputs it to a delegate method.
//
// # Discussion
//
// Call this method only when the [PreferredOutputSegmentInterval] property
// value is [indefinite].
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/flushSegment()
//
// [indefinite]: https://developer.apple.com/documentation/CoreMedia/CMTime/indefinite
func (a AVAssetWriter) FlushSegment() {
	objc.Send[objc.ID](a.ID, objc.Sel("flushSegment"))
}

// Returns a new object that writes media data to a container file at the
// output URL.
//
// outputURL: The location of the file to write.
//
// outputFileType: The type of container file to write.
//
// # Return Value
//
// A new asset writer.
//
// # Discussion
//
// Writing fails if a file already exists at the output URL.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/init(url:fileType:)
func (_AVAssetWriterClass AVAssetWriterClass) AssetWriterWithURLFileTypeError(outputURL foundation.INSURL, outputFileType AVFileType) (AVAssetWriter, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_AVAssetWriterClass.class), objc.Sel("assetWriterWithURL:fileType:error:"), outputURL, objc.String(string(outputFileType)), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return AVAssetWriter{}, foundation.NSErrorFrom(errorPtr)
	}
	return AVAssetWriterFromID(rv), nil

}

// The inputs an asset writer contains.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/inputs
func (a AVAssetWriter) Inputs() []AVAssetWriterInput {
	rv := objc.Send[[]objc.ID](a.ID, objc.Sel("inputs"))
	return objc.ConvertSlice(rv, func(id objc.ID) AVAssetWriterInput {
		return AVAssetWriterInputFromID(id)
	})
}

// The media types the asset writer supports adding as inputs.
//
// # Discussion
//
// Use this property to determine the acceptable media types for the container
// file that the asset writer outputs.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/availableMediaTypes
func (a AVAssetWriter) AvailableMediaTypes() []string {
	rv := objc.Send[[]objc.ID](a.ID, objc.Sel("availableMediaTypes"))
	return objc.ConvertSliceToStrings(rv)
}

// The input groups an asset writer contains.
//
// # Discussion
//
// Add input groups to the asset writer using its [AddInputGroup] method.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/inputGroups
func (a AVAssetWriter) InputGroups() []AVAssetWriterInputGroup {
	rv := objc.Send[[]objc.ID](a.ID, objc.Sel("inputGroups"))
	return objc.ConvertSlice(rv, func(id objc.ID) AVAssetWriterInputGroup {
		return AVAssetWriterInputGroupFromID(id)
	})
}

// An array of metadata items to write to the output file.
//
// # Discussion
//
// You can’t modify this property value after writing starts.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/metadata
func (a AVAssetWriter) Metadata() []AVMetadataItem {
	rv := objc.Send[[]objc.ID](a.ID, objc.Sel("metadata"))
	return objc.ConvertSlice(rv, func(id objc.ID) AVMetadataItem {
		return AVMetadataItemFromID(id)
	})
}
func (a AVAssetWriter) SetMetadata(value []AVMetadataItem) {
	objc.Send[struct{}](a.ID, objc.Sel("setMetadata:"), objectivec.IObjectSliceToNSArray(value))
}

// A Boolean value that indicates whether to write the output file to make it
// more suitable for playback over a network.
//
// # Discussion
//
// Setting this value to true writes the output file in a form that enables a
// player to begin playing the media after downloading only a small portion of
// it.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/shouldOptimizeForNetworkUse
func (a AVAssetWriter) ShouldOptimizeForNetworkUse() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("shouldOptimizeForNetworkUse"))
	return rv
}
func (a AVAssetWriter) SetShouldOptimizeForNetworkUse(value bool) {
	objc.Send[struct{}](a.ID, objc.Sel("setShouldOptimizeForNetworkUse:"), value)
}

// A directory to contain temporary files that the export process generates.
//
// # Discussion
//
// In some configurations, such as performing multipass encoding, an asset
// writer may need to write temporary files. Use this property value to set
// the file-system location where it writes temporary files. The system
// deletes all temporary files after it finishes writing successfully, fails,
// or you cancel the writing session.
//
// You can set this value after writing starts.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/directoryForTemporaryFiles
func (a AVAssetWriter) DirectoryForTemporaryFiles() foundation.INSURL {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("directoryForTemporaryFiles"))
	return foundation.NSURLFromID(objc.ID(rv))
}
func (a AVAssetWriter) SetDirectoryForTemporaryFiles(value foundation.INSURL) {
	objc.Send[struct{}](a.ID, objc.Sel("setDirectoryForTemporaryFiles:"), value)
}

// The interval at which to write movie fragments.
//
// # Discussion
//
// Some container formats, such as QuickTime movies, support writing movies in
// fragments. Using this feature enables you to open and play a partially
// written movie in the event that an unexpected error or interruption occurs.
//
// An asset writer disables this feature by default and sets this value to
// [invalid]. To enable fragment writing, set a valid [CMTime] value. For best
// performance when writing to external storage devices, set the movie
// fragment interval to 10 seconds or greater.
//
// You can’t set this value after writing starts.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/movieFragmentInterval
//
// [CMTime]: https://developer.apple.com/documentation/CoreMedia/CMTime
// [invalid]: https://developer.apple.com/documentation/CoreMedia/CMTime/invalid
func (a AVAssetWriter) MovieFragmentInterval() coremedia.CMTime {
	rv := objc.Send[coremedia.CMTime](a.ID, objc.Sel("movieFragmentInterval"))
	return coremedia.CMTime(rv)
}
func (a AVAssetWriter) SetMovieFragmentInterval(value coremedia.CMTime) {
	objc.Send[struct{}](a.ID, objc.Sel("setMovieFragmentInterval:"), value)
}

// The interval at which to write the initial movie fragment.
//
// # Discussion
//
// When using fragment writing, you can set this property value to indicate
// the interval at which to write the initial fragment.
//
// The default value is [invalid], which indicates to use the interval set in
// the [MovieFragmentInterval] property. The [MovieFragmentInterval] property
// is typically set to 10 seconds, so if an error occurs before writing the
// first fragment, the movie file won’t be playable. To avoid this case,
// your app may want to set a shorter interval, such as 1 second, to write the
// initial fragment, and then use a 10 second interval for subsequent
// fragments.
//
// You can’t set this property after writing starts.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/initialMovieFragmentInterval
//
// [invalid]: https://developer.apple.com/documentation/CoreMedia/CMTime/invalid
func (a AVAssetWriter) InitialMovieFragmentInterval() coremedia.CMTime {
	rv := objc.Send[coremedia.CMTime](a.ID, objc.Sel("initialMovieFragmentInterval"))
	return coremedia.CMTime(rv)
}
func (a AVAssetWriter) SetInitialMovieFragmentInterval(value coremedia.CMTime) {
	objc.Send[struct{}](a.ID, objc.Sel("setInitialMovieFragmentInterval:"), value)
}

// The sequence number of the initial movie fragment.
//
// # Discussion
//
// If you combine movie fragments that you create from multiple asset writers,
// movie fragment sequence numbers need to increase monotonically across the
// entire combined collection, in temporal order. The default value of this
// property is `1`.
//
// You can’t set this property after writing starts.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/initialMovieFragmentSequenceNumber
func (a AVAssetWriter) InitialMovieFragmentSequenceNumber() int {
	rv := objc.Send[int](a.ID, objc.Sel("initialMovieFragmentSequenceNumber"))
	return rv
}
func (a AVAssetWriter) SetInitialMovieFragmentSequenceNumber(value int) {
	objc.Send[struct{}](a.ID, objc.Sel("setInitialMovieFragmentSequenceNumber:"), value)
}

// A Boolean value that indicates whether the asset writer outputs movie
// fragments suitable for combining with others.
//
// # Discussion
//
// The default value is false. Set the value to true when you use multiple
// asset writers to produce distinct streams that complement each other, such
// as HLS encodings or bit rate variants.
//
// You can’t set this property after writing starts.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/producesCombinableFragments
func (a AVAssetWriter) ProducesCombinableFragments() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("producesCombinableFragments"))
	return rv
}
func (a AVAssetWriter) SetProducesCombinableFragments(value bool) {
	objc.Send[struct{}](a.ID, objc.Sel("setProducesCombinableFragments:"), value)
}

// A hint of the final duration of the output file.
//
// # Discussion
//
// The default value of [invalid] indicates that the asset writer doesn’t
// write an overall duration hint to the file. The asset writer ignores this
// value if it doesn’t write movie fragments.
//
// You can’t set this property after writing starts.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/overallDurationHint
//
// [invalid]: https://developer.apple.com/documentation/CoreMedia/CMTime/invalid
func (a AVAssetWriter) OverallDurationHint() coremedia.CMTime {
	rv := objc.Send[coremedia.CMTime](a.ID, objc.Sel("overallDurationHint"))
	return coremedia.CMTime(rv)
}
func (a AVAssetWriter) SetOverallDurationHint(value coremedia.CMTime) {
	objc.Send[struct{}](a.ID, objc.Sel("setOverallDurationHint:"), value)
}

// The time scale of the movie.
//
// # Discussion
//
// The default value is `0`, which indicates that the asset writer chooses an
// appropriate value, if applicable.
//
// You can’t set this property value after writing starts.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/movieTimeScale
func (a AVAssetWriter) MovieTimeScale() int32 {
	rv := objc.Send[int32](a.ID, objc.Sel("movieTimeScale"))
	return rv
}
func (a AVAssetWriter) SetMovieTimeScale(value int32) {
	objc.Send[struct{}](a.ID, objc.Sel("setMovieTimeScale:"), value)
}

// The status of writing samples to the output file.
//
// # Discussion
//
// This property is thread safe.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/status-swift.property
func (a AVAssetWriter) Status() AVAssetWriterStatus {
	rv := objc.Send[AVAssetWriterStatus](a.ID, objc.Sel("status"))
	return AVAssetWriterStatus(rv)
}

// An error object that describes an asset-writing failure.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/error
func (a AVAssetWriter) Error() foundation.INSError {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("error"))
	return foundation.NSErrorFromID(objc.ID(rv))
}

// A delegate object that responds to asset-writing events.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/delegate
func (a AVAssetWriter) Delegate() AVAssetWriterDelegate {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("delegate"))
	return AVAssetWriterDelegateObjectFromID(rv)
}
func (a AVAssetWriter) SetDelegate(value AVAssetWriterDelegate) {
	objc.Send[struct{}](a.ID, objc.Sel("setDelegate:"), value)
}

// The interval of output segments that you prefer.
//
// # Discussion
//
// The default value is [invalid], which indicates that the asset writer
// chooses an appropriate default value. You may also set a positive numeric
// or [indefinite] time. When the value is [indefinite], each call you make to
// [FlushSegment] outputs a segment data.
//
// You can’t change this value after writing starts.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/preferredOutputSegmentInterval
//
// [indefinite]: https://developer.apple.com/documentation/CoreMedia/CMTime/indefinite
// [invalid]: https://developer.apple.com/documentation/CoreMedia/CMTime/invalid
func (a AVAssetWriter) PreferredOutputSegmentInterval() coremedia.CMTime {
	rv := objc.Send[coremedia.CMTime](a.ID, objc.Sel("preferredOutputSegmentInterval"))
	return coremedia.CMTime(rv)
}
func (a AVAssetWriter) SetPreferredOutputSegmentInterval(value coremedia.CMTime) {
	objc.Send[struct{}](a.ID, objc.Sel("setPreferredOutputSegmentInterval:"), value)
}

// The start time of the initial segment.
//
// # Discussion
//
// This value is relevant only when the [PreferredOutputSegmentInterval]
// property value is positive numeric, in which case you must set a numeric
// time.
//
// You can’t change this value after writing starts.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/initialSegmentStartTime
func (a AVAssetWriter) InitialSegmentStartTime() coremedia.CMTime {
	rv := objc.Send[coremedia.CMTime](a.ID, objc.Sel("initialSegmentStartTime"))
	return coremedia.CMTime(rv)
}
func (a AVAssetWriter) SetInitialSegmentStartTime(value coremedia.CMTime) {
	objc.Send[struct{}](a.ID, objc.Sel("setInitialSegmentStartTime:"), value)
}

// A profile for the output file type.
//
// # Discussion
//
// The default value is `nil`, which indicates that the writer chooses an
// appropriate default profile for the output file type. If your app requires
// segment data that’s suitable for streaming, set the value to
// [mpeg4AppleHLS] or [mpeg4CMAFCompliant] to output CMAF-compliant [mp4]
// data.
//
// You can’t change this value after writing starts.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/outputFileTypeProfile
//
// [mp4]: https://developer.apple.com/documentation/AVFoundation/AVFileType/mp4
// [mpeg4AppleHLS]: https://developer.apple.com/documentation/AVFoundation/AVFileTypeProfile/mpeg4AppleHLS
// [mpeg4CMAFCompliant]: https://developer.apple.com/documentation/AVFoundation/AVFileTypeProfile/mpeg4CMAFCompliant
func (a AVAssetWriter) OutputFileTypeProfile() AVFileTypeProfile {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("outputFileTypeProfile"))
	return AVFileTypeProfile(foundation.NSStringFromID(rv).String())
}
func (a AVAssetWriter) SetOutputFileTypeProfile(value AVFileTypeProfile) {
	objc.Send[struct{}](a.ID, objc.Sel("setOutputFileTypeProfile:"), objc.String(string(value)))
}

// The location of the container file that the writer outputs.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/outputURL
func (a AVAssetWriter) OutputURL() foundation.INSURL {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("outputURL"))
	return foundation.NSURLFromID(objc.ID(rv))
}

// The type of container file that the writer outputs.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/outputFileType
func (a AVAssetWriter) OutputFileType() AVFileType {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("outputFileType"))
	return AVFileType(foundation.NSStringFromID(rv).String())
}

// FinishWritingSync is a synchronous wrapper around [AVAssetWriter.FinishWritingWithCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (a AVAssetWriter) FinishWritingSync(ctx context.Context) error {
	done := make(chan struct{}, 1)
	a.FinishWritingWithCompletionHandler(func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
