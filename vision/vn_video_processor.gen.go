// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"errors"
	"github.com/tmc/apple/coremedia"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VNVideoProcessor] class.
var (
	_VNVideoProcessorClass     VNVideoProcessorClass
	_VNVideoProcessorClassOnce sync.Once
)

func getVNVideoProcessorClass() VNVideoProcessorClass {
	_VNVideoProcessorClassOnce.Do(func() {
		_VNVideoProcessorClass = VNVideoProcessorClass{class: objc.GetClass("VNVideoProcessor")}
	})
	return _VNVideoProcessorClass
}

// GetVNVideoProcessorClass returns the class object for VNVideoProcessor.
func GetVNVideoProcessorClass() VNVideoProcessorClass {
	return getVNVideoProcessorClass()
}

type VNVideoProcessorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VNVideoProcessorClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VNVideoProcessorClass) Alloc() VNVideoProcessor {
	rv := objc.Send[VNVideoProcessor](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// An object that performs offline analysis of video content.
//
// # Creating a Video Processor
//
//   - [VNVideoProcessor.InitWithURL]: Creates a video processor to perform Vision requests against the specified video asset.
//
// # Performing Requests
//
//   - [VNVideoProcessor.AddRequestProcessingOptionsError]: Adds a request with processing options to the video processor.
//   - [VNVideoProcessor.RemoveRequestError]: Removes a Vision request from the video processor’s request queue.
//   - [VNVideoProcessor.AnalyzeTimeRangeError]: Analyzes a time range of video content.
//   - [VNVideoProcessor.Cancel]: Cancels the video processing.
//
// See: https://developer.apple.com/documentation/Vision/VNVideoProcessor
type VNVideoProcessor struct {
	objectivec.Object
}

// VNVideoProcessorFromID constructs a [VNVideoProcessor] from an objc.ID.
//
// An object that performs offline analysis of video content.
func VNVideoProcessorFromID(id objc.ID) VNVideoProcessor {
	return VNVideoProcessor{objectivec.Object{ID: id}}
}
// NOTE: VNVideoProcessor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [VNVideoProcessor] class.
//
// # Creating a Video Processor
//
//   - [IVNVideoProcessor.InitWithURL]: Creates a video processor to perform Vision requests against the specified video asset.
//
// # Performing Requests
//
//   - [IVNVideoProcessor.AddRequestProcessingOptionsError]: Adds a request with processing options to the video processor.
//   - [IVNVideoProcessor.RemoveRequestError]: Removes a Vision request from the video processor’s request queue.
//   - [IVNVideoProcessor.AnalyzeTimeRangeError]: Analyzes a time range of video content.
//   - [IVNVideoProcessor.Cancel]: Cancels the video processing.
//
// See: https://developer.apple.com/documentation/Vision/VNVideoProcessor
type IVNVideoProcessor interface {
	objectivec.IObject

	// Topic: Creating a Video Processor

	// Creates a video processor to perform Vision requests against the specified video asset.
	InitWithURL(videoURL foundation.INSURL) VNVideoProcessor

	// Topic: Performing Requests

	// Adds a request with processing options to the video processor.
	AddRequestProcessingOptionsError(request IVNRequest, processingOptions IVNVideoProcessorRequestProcessingOptions) (bool, error)
	// Removes a Vision request from the video processor’s request queue.
	RemoveRequestError(request IVNRequest) (bool, error)
	// Analyzes a time range of video content.
	AnalyzeTimeRangeError(timeRange coremedia.CMTimeRange) (bool, error)
	// Cancels the video processing.
	Cancel()
}

// Init initializes the instance.
func (v VNVideoProcessor) Init() VNVideoProcessor {
	rv := objc.Send[VNVideoProcessor](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VNVideoProcessor) Autorelease() VNVideoProcessor {
	rv := objc.Send[VNVideoProcessor](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVNVideoProcessor creates a new VNVideoProcessor instance.
func NewVNVideoProcessor() VNVideoProcessor {
	class := getVNVideoProcessorClass()
	rv := objc.Send[VNVideoProcessor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a video processor to perform Vision requests against the specified
// video asset.
//
// videoURL: The video asset URL. The specified asset must be a video format supported
// by AVFoundation.
//
// See: https://developer.apple.com/documentation/Vision/VNVideoProcessor/init(url:)
func NewVideoProcessorWithURL(videoURL foundation.INSURL) VNVideoProcessor {
	instance := getVNVideoProcessorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithURL:"), videoURL)
	return VNVideoProcessorFromID(rv)
}

// Creates a video processor to perform Vision requests against the specified
// video asset.
//
// videoURL: The video asset URL. The specified asset must be a video format supported
// by AVFoundation.
//
// See: https://developer.apple.com/documentation/Vision/VNVideoProcessor/init(url:)
func (v VNVideoProcessor) InitWithURL(videoURL foundation.INSURL) VNVideoProcessor {
	rv := objc.Send[VNVideoProcessor](v.ID, objc.Sel("initWithURL:"), videoURL)
	return rv
}
// Adds a request with processing options to the video processor.
//
// request: The Vision request to add.
//
// processingOptions: The processing options to apply.
//
// # Discussion
// 
// Call this method either before calling [AnalyzeTimeRangeError] or from
// within the completion handler of an already associated request.
//
// See: https://developer.apple.com/documentation/Vision/VNVideoProcessor/addRequest(_:processingOptions:)
func (v VNVideoProcessor) AddRequestProcessingOptionsError(request IVNRequest, processingOptions IVNVideoProcessorRequestProcessingOptions) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](v.ID, objc.Sel("addRequest:processingOptions:error:"), request, processingOptions, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("addRequest:processingOptions:error: returned NO with nil NSError")
	}
	return rv, nil

}
// Removes a Vision request from the video processor’s request queue.
//
// request: The request to remove.
//
// See: https://developer.apple.com/documentation/Vision/VNVideoProcessor/removeRequest(_:)
func (v VNVideoProcessor) RemoveRequestError(request IVNRequest) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](v.ID, objc.Sel("removeRequest:error:"), request, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("removeRequest:error: returned NO with nil NSError")
	}
	return rv, nil

}
// Analyzes a time range of video content.
//
// timeRange: The time range to analyze. The value must be within the time range of the
// video asset.
//
// # Discussion
// 
// The system executes this method synchronously, so you typically call it
// from a separate dispatch queue. It returns when the video processor
// finishes analyzing the time range or if an error prevents processing.
//
// See: https://developer.apple.com/documentation/Vision/VNVideoProcessor/analyze(_:)
func (v VNVideoProcessor) AnalyzeTimeRangeError(timeRange coremedia.CMTimeRange) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](v.ID, objc.Sel("analyzeTimeRange:error:"), timeRange, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("analyzeTimeRange:error: returned NO with nil NSError")
	}
	return rv, nil

}
// Cancels the video processing.
//
// See: https://developer.apple.com/documentation/Vision/VNVideoProcessor/cancel()
func (v VNVideoProcessor) Cancel() {
	objc.Send[objc.ID](v.ID, objc.Sel("cancel"))
}

