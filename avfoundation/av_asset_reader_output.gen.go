// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVAssetReaderOutput] class.
var (
	_AVAssetReaderOutputClass     AVAssetReaderOutputClass
	_AVAssetReaderOutputClassOnce sync.Once
)

func getAVAssetReaderOutputClass() AVAssetReaderOutputClass {
	_AVAssetReaderOutputClassOnce.Do(func() {
		_AVAssetReaderOutputClass = AVAssetReaderOutputClass{class: objc.GetClass("AVAssetReaderOutput")}
	})
	return _AVAssetReaderOutputClass
}

// GetAVAssetReaderOutputClass returns the class object for AVAssetReaderOutput.
func GetAVAssetReaderOutputClass() AVAssetReaderOutputClass {
	return getAVAssetReaderOutputClass()
}

type AVAssetReaderOutputClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVAssetReaderOutputClass) Alloc() AVAssetReaderOutput {
	rv := objc.Send[AVAssetReaderOutput](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An abstract class that defines the interface to read media samples from an
// asset reader.
//
// # Overview
// 
// You add concrete output instances, such as [AVAssetReaderTrackOutput] or
// [AVAssetReaderVideoCompositionOutput], to an asset reader to perform
// specific tasks.
//
// # Inspecting the media type
//
//   - [AVAssetReaderOutput.MediaType]: The media type of samples that the output reads.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetReaderOutput
type AVAssetReaderOutput struct {
	objectivec.Object
}

// AVAssetReaderOutputFromID constructs a [AVAssetReaderOutput] from an objc.ID.
//
// An abstract class that defines the interface to read media samples from an
// asset reader.
func AVAssetReaderOutputFromID(id objc.ID) AVAssetReaderOutput {
	return AVAssetReaderOutput{objectivec.Object{ID: id}}
}
// NOTE: AVAssetReaderOutput adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVAssetReaderOutput] class.
//
// # Inspecting the media type
//
//   - [IAVAssetReaderOutput.MediaType]: The media type of samples that the output reads.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetReaderOutput
type IAVAssetReaderOutput interface {
	objectivec.IObject

	// Topic: Inspecting the media type

	// The media type of samples that the output reads.
	MediaType() AVMediaType
}

// Init initializes the instance.
func (a AVAssetReaderOutput) Init() AVAssetReaderOutput {
	rv := objc.Send[AVAssetReaderOutput](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AVAssetReaderOutput) Autorelease() AVAssetReaderOutput {
	rv := objc.Send[AVAssetReaderOutput](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVAssetReaderOutput creates a new AVAssetReaderOutput instance.
func NewAVAssetReaderOutput() AVAssetReaderOutput {
	class := getAVAssetReaderOutputClass()
	rv := objc.Send[AVAssetReaderOutput](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The media type of samples that the output reads.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetReaderOutput/mediaType
func (a AVAssetReaderOutput) MediaType() AVMediaType {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("mediaType"))
	return AVMediaType(foundation.NSStringFromID(rv).String())
}

