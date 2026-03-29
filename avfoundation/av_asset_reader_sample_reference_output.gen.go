// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
)

// The class instance for the [AVAssetReaderSampleReferenceOutput] class.
var (
	_AVAssetReaderSampleReferenceOutputClass     AVAssetReaderSampleReferenceOutputClass
	_AVAssetReaderSampleReferenceOutputClassOnce sync.Once
)

func getAVAssetReaderSampleReferenceOutputClass() AVAssetReaderSampleReferenceOutputClass {
	_AVAssetReaderSampleReferenceOutputClassOnce.Do(func() {
		_AVAssetReaderSampleReferenceOutputClass = AVAssetReaderSampleReferenceOutputClass{class: objc.GetClass("AVAssetReaderSampleReferenceOutput")}
	})
	return _AVAssetReaderSampleReferenceOutputClass
}

// GetAVAssetReaderSampleReferenceOutputClass returns the class object for AVAssetReaderSampleReferenceOutput.
func GetAVAssetReaderSampleReferenceOutputClass() AVAssetReaderSampleReferenceOutputClass {
	return getAVAssetReaderSampleReferenceOutputClass()
}

type AVAssetReaderSampleReferenceOutputClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVAssetReaderSampleReferenceOutputClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVAssetReaderSampleReferenceOutputClass) Alloc() AVAssetReaderSampleReferenceOutput {
	rv := objc.Send[AVAssetReaderSampleReferenceOutput](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that reads sample references from an asset track.
//
// # Overview
// 
// Apps can extract information about the location of samples in a track —
// the file URL and offset — by adding an instance of this class to an asset
// reader. Read the [KCMSampleBufferAttachmentKey_SampleReferenceURL] and
// [KCMSampleBufferAttachmentKey_SampleReferenceByteOffset] attachments on the
// extracted sample buffers to get the location of the sample data.
// 
// You can also append sample buffers that you extract using this class to an
// [AVAssetWriterInput] instance to create movie tracks that aren’t
// self-contained and reference data in the original file instead. To write
// tracks that aren’t self-contained, use instances of [AVAssetWriter] that
// you configure to write files of type [mov].
// 
// Because this output doesn’t return sample data, it ignores the value of
// the [alwaysCopiesSampleData] property.
//
// [alwaysCopiesSampleData]: https://developer.apple.com/documentation/AVFoundation/AVAssetReaderOutput/alwaysCopiesSampleData
// [mov]: https://developer.apple.com/documentation/AVFoundation/AVFileType/mov
//
// # Creating a sample reference output
//
//   - [AVAssetReaderSampleReferenceOutput.InitWithTrack]: Creates an object that supplies sample references.
//
// # Inspecting the track
//
//   - [AVAssetReaderSampleReferenceOutput.Track]: The track from which the output reads sample references.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetReaderSampleReferenceOutput
type AVAssetReaderSampleReferenceOutput struct {
	AVAssetReaderOutput
}

// AVAssetReaderSampleReferenceOutputFromID constructs a [AVAssetReaderSampleReferenceOutput] from an objc.ID.
//
// An object that reads sample references from an asset track.
func AVAssetReaderSampleReferenceOutputFromID(id objc.ID) AVAssetReaderSampleReferenceOutput {
	return AVAssetReaderSampleReferenceOutput{AVAssetReaderOutput: AVAssetReaderOutputFromID(id)}
}
// NOTE: AVAssetReaderSampleReferenceOutput adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVAssetReaderSampleReferenceOutput] class.
//
// # Creating a sample reference output
//
//   - [IAVAssetReaderSampleReferenceOutput.InitWithTrack]: Creates an object that supplies sample references.
//
// # Inspecting the track
//
//   - [IAVAssetReaderSampleReferenceOutput.Track]: The track from which the output reads sample references.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetReaderSampleReferenceOutput
type IAVAssetReaderSampleReferenceOutput interface {
	IAVAssetReaderOutput

	// Topic: Creating a sample reference output

	// Creates an object that supplies sample references.
	InitWithTrack(track IAVAssetTrack) AVAssetReaderSampleReferenceOutput

	// Topic: Inspecting the track

	// The track from which the output reads sample references.
	Track() IAVAssetTrack

	// A Boolean value that indicates whether the output vends copied sample data.
	AlwaysCopiesSampleData() bool
	SetAlwaysCopiesSampleData(value bool)
	// Indicates the byte offset at which the sample data begins (type `CFNumber`).
	KCMSampleBufferAttachmentKey_SampleReferenceByteOffset() corefoundation.CFStringRef
	// Indicates the URL where the sample data is (type `CFURL`).
	KCMSampleBufferAttachmentKey_SampleReferenceURL() corefoundation.CFStringRef
	// The UTI for the QuickTime movie file format.
	Mov() AVFileType
}

// Init initializes the instance.
func (a AVAssetReaderSampleReferenceOutput) Init() AVAssetReaderSampleReferenceOutput {
	rv := objc.Send[AVAssetReaderSampleReferenceOutput](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AVAssetReaderSampleReferenceOutput) Autorelease() AVAssetReaderSampleReferenceOutput {
	rv := objc.Send[AVAssetReaderSampleReferenceOutput](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVAssetReaderSampleReferenceOutput creates a new AVAssetReaderSampleReferenceOutput instance.
func NewAVAssetReaderSampleReferenceOutput() AVAssetReaderSampleReferenceOutput {
	class := getAVAssetReaderSampleReferenceOutputClass()
	rv := objc.Send[AVAssetReaderSampleReferenceOutput](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates an object that supplies sample references.
//
// track: The track for which to provide sample references.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetReaderSampleReferenceOutput/init(track:)
func NewAssetReaderSampleReferenceOutputWithTrack(track IAVAssetTrack) AVAssetReaderSampleReferenceOutput {
	instance := getAVAssetReaderSampleReferenceOutputClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTrack:"), track)
	return AVAssetReaderSampleReferenceOutputFromID(rv)
}

// Creates an object that supplies sample references.
//
// track: The track for which to provide sample references.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetReaderSampleReferenceOutput/init(track:)
func (a AVAssetReaderSampleReferenceOutput) InitWithTrack(track IAVAssetTrack) AVAssetReaderSampleReferenceOutput {
	rv := objc.Send[AVAssetReaderSampleReferenceOutput](a.ID, objc.Sel("initWithTrack:"), track)
	return rv
}

// Returns a new object that supplies sample references.
//
// track: The track for which to provide sample references.
//
// # Return Value
// 
// A sample rererence output object.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetReaderSampleReferenceOutput/assetReaderSampleReferenceOutputWithTrack:
func (_AVAssetReaderSampleReferenceOutputClass AVAssetReaderSampleReferenceOutputClass) AssetReaderSampleReferenceOutputWithTrack(track IAVAssetTrack) AVAssetReaderSampleReferenceOutput {
	rv := objc.Send[objc.ID](objc.ID(_AVAssetReaderSampleReferenceOutputClass.class), objc.Sel("assetReaderSampleReferenceOutputWithTrack:"), track)
	return AVAssetReaderSampleReferenceOutputFromID(rv)
}

// The track from which the output reads sample references.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetReaderSampleReferenceOutput/track
func (a AVAssetReaderSampleReferenceOutput) Track() IAVAssetTrack {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("track"))
	return AVAssetTrackFromID(objc.ID(rv))
}
// A Boolean value that indicates whether the output vends copied sample data.
//
// See: https://developer.apple.com/documentation/avfoundation/avassetreaderoutput/alwayscopiessampledata
func (a AVAssetReaderSampleReferenceOutput) AlwaysCopiesSampleData() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("alwaysCopiesSampleData"))
	return rv
}
func (a AVAssetReaderSampleReferenceOutput) SetAlwaysCopiesSampleData(value bool) {
	objc.Send[struct{}](a.ID, objc.Sel("setAlwaysCopiesSampleData:"), value)
}
// Indicates the byte offset at which the sample data begins (type
// `CFNumber`).
//
// See: https://developer.apple.com/documentation/CoreMedia/kCMSampleBufferAttachmentKey_SampleReferenceByteOffset
func (a AVAssetReaderSampleReferenceOutput) KCMSampleBufferAttachmentKey_SampleReferenceByteOffset() corefoundation.CFStringRef {
	rv := objc.Send[corefoundation.CFStringRef](a.ID, objc.Sel("kCMSampleBufferAttachmentKey_SampleReferenceByteOffset"))
	return corefoundation.CFStringRef(rv)
}
// Indicates the URL where the sample data is (type `CFURL`).
//
// See: https://developer.apple.com/documentation/CoreMedia/kCMSampleBufferAttachmentKey_SampleReferenceURL
func (a AVAssetReaderSampleReferenceOutput) KCMSampleBufferAttachmentKey_SampleReferenceURL() corefoundation.CFStringRef {
	rv := objc.Send[corefoundation.CFStringRef](a.ID, objc.Sel("kCMSampleBufferAttachmentKey_SampleReferenceURL"))
	return corefoundation.CFStringRef(rv)
}
// The UTI for the QuickTime movie file format.
//
// See: https://developer.apple.com/documentation/avfoundation/avfiletype/mov
func (a AVAssetReaderSampleReferenceOutput) Mov() AVFileType {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("AVFileTypeQuickTimeMovie"))
	return AVFileType(foundation.NSStringFromID(rv).String())
}

