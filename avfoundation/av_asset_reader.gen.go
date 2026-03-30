// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/coremedia"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVAssetReader] class.
var (
	_AVAssetReaderClass     AVAssetReaderClass
	_AVAssetReaderClassOnce sync.Once
)

func getAVAssetReaderClass() AVAssetReaderClass {
	_AVAssetReaderClassOnce.Do(func() {
		_AVAssetReaderClass = AVAssetReaderClass{class: objc.GetClass("AVAssetReader")}
	})
	return _AVAssetReaderClass
}

// GetAVAssetReaderClass returns the class object for AVAssetReader.
func GetAVAssetReaderClass() AVAssetReaderClass {
	return getAVAssetReaderClass()
}

type AVAssetReaderClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVAssetReaderClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVAssetReaderClass) Alloc() AVAssetReader {
	rv := objc.Send[AVAssetReader](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that reads media data from an asset.
//
// # Overview
//
// Use an asset reader to read media data from instances of [AVAsset]. The
// assets you read may represent file-based media like QuickTime movies or
// MPEG-4 files, or media that you compose from multiple sources using
// [AVComposition].
//
// # Creating an asset reader
//
//   - [AVAssetReader.InitWithAssetError]: Creates an object to read media data from an asset.
//
// # Managing outputs
//
//   - [AVAssetReader.CanAddOutput]: Determines whether you can add the output to the asset reader.
//   - [AVAssetReader.Outputs]: The outputs from which you read media data.
//
// # Configuring reading
//
//   - [AVAssetReader.TimeRange]: The time range within the asset to read.
//   - [AVAssetReader.SetTimeRange]
//   - [AVAssetReader.Status]: The status of reading sample buffers from the asset.
//   - [AVAssetReader.Error]: An error that describes the reason for a failure.
//
// # Controlling reading
//
//   - [AVAssetReader.CancelReading]: Cancels any background work and stops the reader’s outputs from reading more samples.
//
// # Inspecting the asset
//
//   - [AVAssetReader.Asset]: The asset from which to read media data.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetReader
type AVAssetReader struct {
	objectivec.Object
}

// AVAssetReaderFromID constructs a [AVAssetReader] from an objc.ID.
//
// An object that reads media data from an asset.
func AVAssetReaderFromID(id objc.ID) AVAssetReader {
	return AVAssetReader{objectivec.Object{ID: id}}
}

// NOTE: AVAssetReader adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVAssetReader] class.
//
// # Creating an asset reader
//
//   - [IAVAssetReader.InitWithAssetError]: Creates an object to read media data from an asset.
//
// # Managing outputs
//
//   - [IAVAssetReader.CanAddOutput]: Determines whether you can add the output to the asset reader.
//   - [IAVAssetReader.Outputs]: The outputs from which you read media data.
//
// # Configuring reading
//
//   - [IAVAssetReader.TimeRange]: The time range within the asset to read.
//   - [IAVAssetReader.SetTimeRange]
//   - [IAVAssetReader.Status]: The status of reading sample buffers from the asset.
//   - [IAVAssetReader.Error]: An error that describes the reason for a failure.
//
// # Controlling reading
//
//   - [IAVAssetReader.CancelReading]: Cancels any background work and stops the reader’s outputs from reading more samples.
//
// # Inspecting the asset
//
//   - [IAVAssetReader.Asset]: The asset from which to read media data.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetReader
type IAVAssetReader interface {
	objectivec.IObject

	// Topic: Creating an asset reader

	// Creates an object to read media data from an asset.
	InitWithAssetError(asset IAVAsset) (AVAssetReader, error)

	// Topic: Managing outputs

	// Determines whether you can add the output to the asset reader.
	CanAddOutput(output IAVAssetReaderOutput) bool
	// The outputs from which you read media data.
	Outputs() []AVAssetReaderOutput

	// Topic: Configuring reading

	// The time range within the asset to read.
	TimeRange() coremedia.CMTimeRange
	SetTimeRange(value coremedia.CMTimeRange)
	// The status of reading sample buffers from the asset.
	Status() AVAssetReaderStatus
	// An error that describes the reason for a failure.
	Error() foundation.INSError

	// Topic: Controlling reading

	// Cancels any background work and stops the reader’s outputs from reading more samples.
	CancelReading()

	// Topic: Inspecting the asset

	// The asset from which to read media data.
	Asset() IAVAsset
}

// Init initializes the instance.
func (a AVAssetReader) Init() AVAssetReader {
	rv := objc.Send[AVAssetReader](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AVAssetReader) Autorelease() AVAssetReader {
	rv := objc.Send[AVAssetReader](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVAssetReader creates a new AVAssetReader instance.
func NewAVAssetReader() AVAssetReader {
	class := getAVAssetReaderClass()
	rv := objc.Send[AVAssetReader](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates an object to read media data from an asset.
//
// asset: The asset from which to read media data.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetReader/init(asset:)
func NewAssetReaderWithAssetError(asset IAVAsset) (AVAssetReader, error) {
	var errorPtr objc.ID
	instance := getAVAssetReaderClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithAsset:error:"), asset, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return AVAssetReader{}, foundation.NSErrorFrom(errorPtr)
	}
	return AVAssetReaderFromID(rv), nil
}

// Creates an object to read media data from an asset.
//
// asset: The asset from which to read media data.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetReader/init(asset:)
func (a AVAssetReader) InitWithAssetError(asset IAVAsset) (AVAssetReader, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](a.ID, objc.Sel("initWithAsset:error:"), asset, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return AVAssetReader{}, foundation.NSErrorFrom(errorPtr)
	}
	return AVAssetReaderFromID(rv), nil

}

// Determines whether you can add the output to the asset reader.
//
// output: The asset reader output to test.
//
// # Return Value
//
// true if you can add the output; otherwise, false.
//
// # Discussion
//
// You may only add outputs that retrieve media data from the asset that you
// associate with the asset reader.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetReader/canAdd(_:)
func (a AVAssetReader) CanAddOutput(output IAVAssetReaderOutput) bool {
	rv := objc.Send[bool](a.ID, objc.Sel("canAddOutput:"), output)
	return rv
}

// Cancels any background work and stops the reader’s outputs from reading
// more samples.
//
// # Discussion
//
// To stop reading samples before reaching the end of the time range, call
// this method to stop any in-progress background read operations.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetReader/cancelReading()
func (a AVAssetReader) CancelReading() {
	objc.Send[objc.ID](a.ID, objc.Sel("cancelReading"))
}

// Returns a new object to read media data from an asset.
//
// asset: The asset from which to read media data.
//
// outError: On return, if the initialization fails, a pointer to an error object
// provides the details of the failure.
//
// # Return Value
//
// A new asset reader object.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetReader/assetReaderWithAsset:error:
func (_AVAssetReaderClass AVAssetReaderClass) AssetReaderWithAssetError(asset IAVAsset) (AVAssetReader, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_AVAssetReaderClass.class), objc.Sel("assetReaderWithAsset:error:"), asset, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return AVAssetReader{}, foundation.NSErrorFrom(errorPtr)
	}
	return AVAssetReaderFromID(rv), nil

}

// The outputs from which you read media data.
//
// # Discussion
//
// The array contains the concrete instances of [AVAssetReaderOutput] that you
// associate with the reader.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetReader/outputs
func (a AVAssetReader) Outputs() []AVAssetReaderOutput {
	rv := objc.Send[[]objc.ID](a.ID, objc.Sel("outputs"))
	return objc.ConvertSlice(rv, func(id objc.ID) AVAssetReaderOutput {
		return AVAssetReaderOutputFromID(id)
	})
}

// The time range within the asset to read.
//
// # Discussion
//
// The default value is a time range with a start time of [zero] and a
// duration of [positiveInfinity].
//
// You can’t modify this value after reading starts.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetReader/timeRange
//
// [positiveInfinity]: https://developer.apple.com/documentation/CoreMedia/CMTime/positiveInfinity
// [zero]: https://developer.apple.com/documentation/CoreMedia/CMTime/zero
func (a AVAssetReader) TimeRange() coremedia.CMTimeRange {
	rv := objc.Send[coremedia.CMTimeRange](a.ID, objc.Sel("timeRange"))
	return coremedia.CMTimeRange(rv)
}
func (a AVAssetReader) SetTimeRange(value coremedia.CMTimeRange) {
	objc.Send[struct{}](a.ID, objc.Sel("setTimeRange:"), value)
}

// The status of reading sample buffers from the asset.
//
// # Discussion
//
// Check the value of this property when the [copyNextSampleBuffer()] method
// on [AVAssetReaderOutput] returns `nil` to determine why the output can’t
// read more data.
//
// This property is thread safe.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetReader/status-swift.property
//
// [copyNextSampleBuffer()]: https://developer.apple.com/documentation/AVFoundation/AVAssetReaderOutput/copyNextSampleBuffer()
func (a AVAssetReader) Status() AVAssetReaderStatus {
	rv := objc.Send[AVAssetReaderStatus](a.ID, objc.Sel("status"))
	return AVAssetReaderStatus(rv)
}

// An error that describes the reason for a failure.
//
// # Discussion
//
// The value is `nil` if the asset reader’s status isn’t
// [AVAssetReaderStatusFailed].
//
// This property is thread safe.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetReader/error
func (a AVAssetReader) Error() foundation.INSError {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("error"))
	return foundation.NSErrorFromID(objc.ID(rv))
}

// The asset from which to read media data.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetReader/asset
func (a AVAssetReader) Asset() IAVAsset {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("asset"))
	return AVAssetFromID(objc.ID(rv))
}
