// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"context"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/coregraphics"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVAssetImageGenerator] class.
var (
	_AVAssetImageGeneratorClass     AVAssetImageGeneratorClass
	_AVAssetImageGeneratorClassOnce sync.Once
)

func getAVAssetImageGeneratorClass() AVAssetImageGeneratorClass {
	_AVAssetImageGeneratorClassOnce.Do(func() {
		_AVAssetImageGeneratorClass = AVAssetImageGeneratorClass{class: objc.GetClass("AVAssetImageGenerator")}
	})
	return _AVAssetImageGeneratorClass
}

// GetAVAssetImageGeneratorClass returns the class object for AVAssetImageGenerator.
func GetAVAssetImageGeneratorClass() AVAssetImageGeneratorClass {
	return getAVAssetImageGeneratorClass()
}

type AVAssetImageGeneratorClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVAssetImageGeneratorClass) Alloc() AVAssetImageGenerator {
	rv := objc.Send[AVAssetImageGenerator](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that generates images from a video asset.
//
// # Overview
// 
// Use an image generator to extract images from a video asset at particular
// times within its timeline.
//
// # Creating an image generator
//
//   - [AVAssetImageGenerator.InitWithAsset]: Creates an object that generates images for times within a video asset.
//
// # Configuring image generation
//
//   - [AVAssetImageGenerator.MaximumSize]: The maximum size of images to generate.
//   - [AVAssetImageGenerator.SetMaximumSize]
//   - [AVAssetImageGenerator.RequestedTimeToleranceBefore]: A maximum length of time before the requested time to allow image generation to occur.
//   - [AVAssetImageGenerator.SetRequestedTimeToleranceBefore]
//   - [AVAssetImageGenerator.RequestedTimeToleranceAfter]: A maximum length of time after the requested time to allow image generation to occur.
//   - [AVAssetImageGenerator.SetRequestedTimeToleranceAfter]
//   - [AVAssetImageGenerator.DynamicRangePolicy]: The dynamic range policy to use when generating images.
//   - [AVAssetImageGenerator.SetDynamicRangePolicy]
//   - [AVAssetImageGenerator.AppliesPreferredTrackTransform]: A Boolean value that specifies whether to apply the track matrix or matrices when generating an image from the asset.
//   - [AVAssetImageGenerator.SetAppliesPreferredTrackTransform]
//   - [AVAssetImageGenerator.ApertureMode]: Specifies the aperture mode for the generated image.
//   - [AVAssetImageGenerator.SetApertureMode]
//
// # Configuring compositing
//
//   - [AVAssetImageGenerator.VideoComposition]: A video composition to use when extracting images from assets with multiple video tracks.
//   - [AVAssetImageGenerator.SetVideoComposition]
//   - [AVAssetImageGenerator.CustomVideoCompositor]: A custom video compositor to use when extracting images from assets with multiple video tracks.
//
// # Generating images
//
//   - [AVAssetImageGenerator.GenerateCGImageAsynchronouslyForTimeCompletionHandler]: Generates an image asynchronously for a requested time, and returns the result in a callback.
//   - [AVAssetImageGenerator.GenerateCGImagesAsynchronouslyForTimesCompletionHandler]: Generates images asynchronously for an array of requested times, and returns the results in a callback.
//   - [AVAssetImageGenerator.CancelAllCGImageGeneration]: Cancels all pending image generation requests.
//
// # Accessing the asset
//
//   - [AVAssetImageGenerator.Asset]: The asset that initialized the image generator.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetImageGenerator
type AVAssetImageGenerator struct {
	objectivec.Object
}

// AVAssetImageGeneratorFromID constructs a [AVAssetImageGenerator] from an objc.ID.
//
// An object that generates images from a video asset.
func AVAssetImageGeneratorFromID(id objc.ID) AVAssetImageGenerator {
	return AVAssetImageGenerator{objectivec.Object{ID: id}}
}
// NOTE: AVAssetImageGenerator adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVAssetImageGenerator] class.
//
// # Creating an image generator
//
//   - [IAVAssetImageGenerator.InitWithAsset]: Creates an object that generates images for times within a video asset.
//
// # Configuring image generation
//
//   - [IAVAssetImageGenerator.MaximumSize]: The maximum size of images to generate.
//   - [IAVAssetImageGenerator.SetMaximumSize]
//   - [IAVAssetImageGenerator.RequestedTimeToleranceBefore]: A maximum length of time before the requested time to allow image generation to occur.
//   - [IAVAssetImageGenerator.SetRequestedTimeToleranceBefore]
//   - [IAVAssetImageGenerator.RequestedTimeToleranceAfter]: A maximum length of time after the requested time to allow image generation to occur.
//   - [IAVAssetImageGenerator.SetRequestedTimeToleranceAfter]
//   - [IAVAssetImageGenerator.DynamicRangePolicy]: The dynamic range policy to use when generating images.
//   - [IAVAssetImageGenerator.SetDynamicRangePolicy]
//   - [IAVAssetImageGenerator.AppliesPreferredTrackTransform]: A Boolean value that specifies whether to apply the track matrix or matrices when generating an image from the asset.
//   - [IAVAssetImageGenerator.SetAppliesPreferredTrackTransform]
//   - [IAVAssetImageGenerator.ApertureMode]: Specifies the aperture mode for the generated image.
//   - [IAVAssetImageGenerator.SetApertureMode]
//
// # Configuring compositing
//
//   - [IAVAssetImageGenerator.VideoComposition]: A video composition to use when extracting images from assets with multiple video tracks.
//   - [IAVAssetImageGenerator.SetVideoComposition]
//   - [IAVAssetImageGenerator.CustomVideoCompositor]: A custom video compositor to use when extracting images from assets with multiple video tracks.
//
// # Generating images
//
//   - [IAVAssetImageGenerator.GenerateCGImageAsynchronouslyForTimeCompletionHandler]: Generates an image asynchronously for a requested time, and returns the result in a callback.
//   - [IAVAssetImageGenerator.GenerateCGImagesAsynchronouslyForTimesCompletionHandler]: Generates images asynchronously for an array of requested times, and returns the results in a callback.
//   - [IAVAssetImageGenerator.CancelAllCGImageGeneration]: Cancels all pending image generation requests.
//
// # Accessing the asset
//
//   - [IAVAssetImageGenerator.Asset]: The asset that initialized the image generator.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetImageGenerator
type IAVAssetImageGenerator interface {
	objectivec.IObject

	// Topic: Creating an image generator

	// Creates an object that generates images for times within a video asset.
	InitWithAsset(asset IAVAsset) AVAssetImageGenerator

	// Topic: Configuring image generation

	// The maximum size of images to generate.
	MaximumSize() corefoundation.CGSize
	SetMaximumSize(value corefoundation.CGSize)
	// A maximum length of time before the requested time to allow image generation to occur.
	RequestedTimeToleranceBefore() objectivec.IObject
	SetRequestedTimeToleranceBefore(value objectivec.IObject)
	// A maximum length of time after the requested time to allow image generation to occur.
	RequestedTimeToleranceAfter() objectivec.IObject
	SetRequestedTimeToleranceAfter(value objectivec.IObject)
	// The dynamic range policy to use when generating images.
	DynamicRangePolicy() AVAssetImageGeneratorDynamicRangePolicy
	SetDynamicRangePolicy(value AVAssetImageGeneratorDynamicRangePolicy)
	// A Boolean value that specifies whether to apply the track matrix or matrices when generating an image from the asset.
	AppliesPreferredTrackTransform() bool
	SetAppliesPreferredTrackTransform(value bool)
	// Specifies the aperture mode for the generated image.
	ApertureMode() AVAssetImageGeneratorApertureMode
	SetApertureMode(value AVAssetImageGeneratorApertureMode)

	// Topic: Configuring compositing

	// A video composition to use when extracting images from assets with multiple video tracks.
	VideoComposition() IAVVideoComposition
	SetVideoComposition(value IAVVideoComposition)
	// A custom video compositor to use when extracting images from assets with multiple video tracks.
	CustomVideoCompositor() AVVideoCompositing

	// Topic: Generating images

	// Generates an image asynchronously for a requested time, and returns the result in a callback.
	GenerateCGImageAsynchronouslyForTimeCompletionHandler(requestedTime objectivec.IObject, handler CGImageRefErrorHandler)
	// Generates images asynchronously for an array of requested times, and returns the results in a callback.
	GenerateCGImagesAsynchronouslyForTimesCompletionHandler(requestedTimes []foundation.NSValue, handler ErrorHandler)
	// Cancels all pending image generation requests.
	CancelAllCGImageGeneration()

	// Topic: Accessing the asset

	// The asset that initialized the image generator.
	Asset() IAVAsset
}

// Init initializes the instance.
func (a AVAssetImageGenerator) Init() AVAssetImageGenerator {
	rv := objc.Send[AVAssetImageGenerator](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AVAssetImageGenerator) Autorelease() AVAssetImageGenerator {
	rv := objc.Send[AVAssetImageGenerator](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVAssetImageGenerator creates a new AVAssetImageGenerator instance.
func NewAVAssetImageGenerator() AVAssetImageGenerator {
	class := getAVAssetImageGeneratorClass()
	rv := objc.Send[AVAssetImageGenerator](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates an object that generates images for times within a video asset.
//
// asset: A video asset from which to generate images.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetImageGenerator/init(asset:)
func NewAssetImageGeneratorWithAsset(asset IAVAsset) AVAssetImageGenerator {
	instance := getAVAssetImageGeneratorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithAsset:"), asset)
	return AVAssetImageGeneratorFromID(rv)
}

// Creates an object that generates images for times within a video asset.
//
// asset: A video asset from which to generate images.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetImageGenerator/init(asset:)
func (a AVAssetImageGenerator) InitWithAsset(asset IAVAsset) AVAssetImageGenerator {
	rv := objc.Send[AVAssetImageGenerator](a.ID, objc.Sel("initWithAsset:"), asset)
	return rv
}
// Generates an image asynchronously for a requested time, and returns the
// result in a callback.
//
// requestedTime: A time in the video timeline for which to generate an image. The requested
// time and actual time at which it generates an image may differ depending on
// the generator’s time tolerance settings.
//
// handler: A callback that the image generator invokes with the result of the request.
//
// requestedTime is a [coremedia.CMTime].
//
// # Discussion
// 
// Swift clients should use the asynchronous [image(at:)] method instead.
//
// [image(at:)]: https://developer.apple.com/documentation/AVFoundation/AVAssetImageGenerator/image(at:)
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetImageGenerator/generateCGImageAsynchronously(for:completionHandler:)
func (a AVAssetImageGenerator) GenerateCGImageAsynchronouslyForTimeCompletionHandler(requestedTime objectivec.IObject, handler CGImageRefErrorHandler) {
_block1, _cleanup1 := NewCGImageRefErrorBlock(handler)
	defer _cleanup1()
	objc.Send[objc.ID](a.ID, objc.Sel("generateCGImageAsynchronouslyForTime:completionHandler:"), requestedTime, _block1)
}
// Generates images asynchronously for an array of requested times, and
// returns the results in a callback.
//
// requestedTimes: An array of times, contained in NSValue objects, in the video timeline for
// which to generate images.
//
// handler: A callback that the image generator invokes for each requested image time.
//
// # Discussion
// 
// Swift clients should prefer the asynchronous `images()` method instead.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetImageGenerator/generateCGImagesAsynchronously(forTimes:completionHandler:)
func (a AVAssetImageGenerator) GenerateCGImagesAsynchronouslyForTimesCompletionHandler(requestedTimes []foundation.NSValue, handler ErrorHandler) {
_block1, _cleanup1 := NewErrorBlock(handler)
	defer _cleanup1()
	objc.Send[objc.ID](a.ID, objc.Sel("generateCGImagesAsynchronouslyForTimes:completionHandler:"), requestedTimes, _block1)
}
// Cancels all pending image generation requests.
//
// # Discussion
// 
// Calling this method invokes the handler block with a result of
// [AssetImageGeneratorCancelled] for all requested times for which the
// generator hasn’t yet produced an image.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetImageGenerator/cancelAllCGImageGeneration()
func (a AVAssetImageGenerator) CancelAllCGImageGeneration() {
	objc.Send[objc.ID](a.ID, objc.Sel("cancelAllCGImageGeneration"))
}

// Returns a new object that generates images for times within a video asset.
//
// asset: A video asset from which to generate images.
//
// # Return Value
// 
// A new image generator.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetImageGenerator/assetImageGeneratorWithAsset:
func (_AVAssetImageGeneratorClass AVAssetImageGeneratorClass) AssetImageGeneratorWithAsset(asset IAVAsset) AVAssetImageGenerator {
	rv := objc.Send[objc.ID](objc.ID(_AVAssetImageGeneratorClass.class), objc.Sel("assetImageGeneratorWithAsset:"), asset)
	return AVAssetImageGeneratorFromID(rv)
}

// The maximum size of images to generate.
//
// # Discussion
// 
// The default value is [zero], which generates images at the asset’s
// unscaled dimensions.
// 
// Setting a size scales images to fit their defined bounding boxes. You
// define the aspect ratio of the scaled image by setting a value for the
// [ApertureMode] property.
//
// [zero]: https://developer.apple.com/documentation/CoreFoundation/CGSize/zero
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetImageGenerator/maximumSize
func (a AVAssetImageGenerator) MaximumSize() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](a.ID, objc.Sel("maximumSize"))
	return corefoundation.CGSize(rv)
}
func (a AVAssetImageGenerator) SetMaximumSize(value corefoundation.CGSize) {
	objc.Send[struct{}](a.ID, objc.Sel("setMaximumSize:"), value)
}
// A maximum length of time before the requested time to allow image
// generation to occur.
//
// # Discussion
// 
// The default value is [positiveInfinity]. Set the values of
// [RequestedTimeToleranceBefore] and [RequestedTimeToleranceAfter] to [zero]
// to request frame-accurate image generation; this may incur additional
// decoding delay.
//
// [positiveInfinity]: https://developer.apple.com/documentation/CoreMedia/CMTime/positiveInfinity
// [zero]: https://developer.apple.com/documentation/CoreMedia/CMTime/zero
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetImageGenerator/requestedTimeToleranceBefore
func (a AVAssetImageGenerator) RequestedTimeToleranceBefore() objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("requestedTimeToleranceBefore"))
	return objectivec.Object{ID: rv}
}
func (a AVAssetImageGenerator) SetRequestedTimeToleranceBefore(value objectivec.IObject) {
	objc.Send[struct{}](a.ID, objc.Sel("setRequestedTimeToleranceBefore:"), value)
}
// A maximum length of time after the requested time to allow image generation
// to occur.
//
// # Discussion
// 
// The default value is [positiveInfinity]. Set the values of
// [RequestedTimeToleranceBefore] and [RequestedTimeToleranceAfter] to [zero]
// to request frame-accurate image generation; this may incur additional
// decoding delay.
//
// [positiveInfinity]: https://developer.apple.com/documentation/CoreMedia/CMTime/positiveInfinity
// [zero]: https://developer.apple.com/documentation/CoreMedia/CMTime/zero
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetImageGenerator/requestedTimeToleranceAfter
func (a AVAssetImageGenerator) RequestedTimeToleranceAfter() objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("requestedTimeToleranceAfter"))
	return objectivec.Object{ID: rv}
}
func (a AVAssetImageGenerator) SetRequestedTimeToleranceAfter(value objectivec.IObject) {
	objc.Send[struct{}](a.ID, objc.Sel("setRequestedTimeToleranceAfter:"), value)
}
// The dynamic range policy to use when generating images.
//
// # Discussion
// 
// This property defaults to [forceSDR].
//
// [forceSDR]: https://developer.apple.com/documentation/AVFoundation/AVAssetImageGenerator/DynamicRangePolicy-swift.struct/forceSDR
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetImageGenerator/dynamicRangePolicy-swift.property
func (a AVAssetImageGenerator) DynamicRangePolicy() AVAssetImageGeneratorDynamicRangePolicy {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("dynamicRangePolicy"))
	return AVAssetImageGeneratorDynamicRangePolicy(foundation.NSStringFromID(rv).String())
}
func (a AVAssetImageGenerator) SetDynamicRangePolicy(value AVAssetImageGeneratorDynamicRangePolicy) {
	objc.Send[struct{}](a.ID, objc.Sel("setDynamicRangePolicy:"), objc.String(string(value)))
}
// A Boolean value that specifies whether to apply the track matrix or
// matrices when generating an image from the asset.
//
// # Discussion
// 
// The default value is [false]. This class only supports rotation by 90, 180,
// or 270 degrees.
// 
// The image generator ignores this property if you set a value for the
// [VideoComposition] property.
//
// [false]: https://developer.apple.com/documentation/Swift/false
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetImageGenerator/appliesPreferredTrackTransform
func (a AVAssetImageGenerator) AppliesPreferredTrackTransform() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("appliesPreferredTrackTransform"))
	return rv
}
func (a AVAssetImageGenerator) SetAppliesPreferredTrackTransform(value bool) {
	objc.Send[struct{}](a.ID, objc.Sel("setAppliesPreferredTrackTransform:"), value)
}
// Specifies the aperture mode for the generated image.
//
// # Discussion
// 
// The default value is [cleanAperture].
//
// [cleanAperture]: https://developer.apple.com/documentation/AVFoundation/AVAssetImageGenerator/ApertureMode-swift.struct/cleanAperture
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetImageGenerator/apertureMode-swift.property
func (a AVAssetImageGenerator) ApertureMode() AVAssetImageGeneratorApertureMode {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("apertureMode"))
	return AVAssetImageGeneratorApertureMode(foundation.NSStringFromID(rv).String())
}
func (a AVAssetImageGenerator) SetApertureMode(value AVAssetImageGeneratorApertureMode) {
	objc.Send[struct{}](a.ID, objc.Sel("setApertureMode:"), objc.String(string(value)))
}
// A video composition to use when extracting images from assets with multiple
// video tracks.
//
// # Discussion
// 
// If you don’t specify a video composition, the generator only uses the
// first enabled video track.
// 
// If specify a video composition, the image generator ignores the value of
// the [AppliesPreferredTrackTransform] property.
// 
// Setting a video composition with any of the following attributes generates
// an exception:
// 
// - A [RenderScale] not equal to `1.0`. - A [RenderSize] with a width or
// height less than `0`. - A [FrameDuration] that’s invalid, or less than or
// equal to [zero]. - A [SourceTrackIDForFrameTiming] less than [zero].
//
// [zero]: https://developer.apple.com/documentation/CoreMedia/CMTime/zero
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetImageGenerator/videoComposition
func (a AVAssetImageGenerator) VideoComposition() IAVVideoComposition {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("videoComposition"))
	return AVVideoCompositionFromID(objc.ID(rv))
}
func (a AVAssetImageGenerator) SetVideoComposition(value IAVVideoComposition) {
	objc.Send[struct{}](a.ID, objc.Sel("setVideoComposition:"), value)
}
// A custom video compositor to use when extracting images from assets with
// multiple video tracks.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetImageGenerator/customVideoCompositor
func (a AVAssetImageGenerator) CustomVideoCompositor() AVVideoCompositing {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("customVideoCompositor"))
	return AVVideoCompositingObjectFromID(rv)
}
// The asset that initialized the image generator.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetImageGenerator/asset
func (a AVAssetImageGenerator) Asset() IAVAsset {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("asset"))
	return AVAssetFromID(objc.ID(rv))
}

// GenerateCGImageAsynchronouslyForTime is a synchronous wrapper around [AVAssetImageGenerator.GenerateCGImageAsynchronouslyForTimeCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (a AVAssetImageGenerator) GenerateCGImageAsynchronouslyForTime(ctx context.Context, requestedTime objectivec.IObject) (coregraphics.CGImageRef, error) {
	type result struct {
		val coregraphics.CGImageRef
		err error
	}
	done := make(chan result, 1)
	a.GenerateCGImageAsynchronouslyForTimeCompletionHandler(requestedTime, func(val coregraphics.CGImageRef, err error) {
		done <- result{val, err}
	})
	select {
	case r := <-done:
		return r.val, r.err
	case <-ctx.Done():
		return 0, ctx.Err()
	}
}

