// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"fmt"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)
var _ = fmt.Sprintf

// Methods for monitoring progress and receiving results from a photo capture output.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoCaptureDelegate
type AVCapturePhotoCaptureDelegate interface {
	objectivec.IObject
}

// AVCapturePhotoCaptureDelegateObject wraps an existing Objective-C object that conforms to the AVCapturePhotoCaptureDelegate protocol.
type AVCapturePhotoCaptureDelegateObject struct {
	objectivec.Object
}
func (o AVCapturePhotoCaptureDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}

// AVCapturePhotoCaptureDelegateObjectFromID constructs a [AVCapturePhotoCaptureDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func AVCapturePhotoCaptureDelegateObjectFromID(id objc.ID) AVCapturePhotoCaptureDelegateObject {
	return AVCapturePhotoCaptureDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Notifies the delegate that the capture output has resolved settings and
// will soon begin its capture process.
//
// output: The photo output performing the capture.
//
// resolvedSettings: An object describing the settings used for this capture. Match this
// object’s [UniqueID] value to the [UniqueID] property of the photo
// settings object you initiated capture with to determine which capture
// request this delegate call corresponds to. You can also use this object to
// find out which values the photo output has chosen for automatic settings.
//
// # Discussion
// 
// The photo output calls this method when it has committed to a choice of
// settings and will soon begin the capture process. This call occurs as early
// as possible after your call to the [CapturePhotoWithSettingsDelegate]
// method, letting you know what to expect for other delegate method calls
// related to the same capture.
// 
// Use this method and the [AVCaptureResolvedPhotoSettings] it provides to
// find out at the earliest possible opportunity which values the photo output
// has chosen for automatic settings, and what the output dimensions for
// captured images and movies will be. For example, if you requested capture
// with the [FlashMode] property set to [CaptureFlashModeAuto], the resolved
// photo settings’ [FlashEnabled] property indicates whether the flash will
// fire during capture.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoCaptureDelegate/photoOutput(_:willBeginCaptureFor:)
func (o AVCapturePhotoCaptureDelegateObject) CaptureOutputWillBeginCaptureForResolvedSettings(output IAVCapturePhotoOutput, resolvedSettings IAVCaptureResolvedPhotoSettings) {
	objc.Send[struct{}](o.ID, objc.Sel("captureOutput:willBeginCaptureForResolvedSettings:"), output, resolvedSettings)
	}
// Notifies the delegate that photo capture is about to occur.
//
// output: The photo output performing the capture.
//
// resolvedSettings: An object describing the settings used for this capture. Match this
// object’s [UniqueID] value to the [UniqueID] property of the photo
// settings object you initiated capture with to determine which capture
// request this delegate call corresponds to. You can also use this object to
// find out which values the photo output has chosen for automatic settings.
//
// # Discussion
// 
// The photo output calls this method as close as possible to the initial
// moment of capture. If the shutter sound is enabled, this call occurs
// immediately after the photo output begins playing the shutter sound.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoCaptureDelegate/photoOutput(_:willCapturePhotoFor:)
func (o AVCapturePhotoCaptureDelegateObject) CaptureOutputWillCapturePhotoForResolvedSettings(output IAVCapturePhotoOutput, resolvedSettings IAVCaptureResolvedPhotoSettings) {
	objc.Send[struct{}](o.ID, objc.Sel("captureOutput:willCapturePhotoForResolvedSettings:"), output, resolvedSettings)
	}
// Notifies the delegate that the photo has been taken.
//
// output: The photo output performing the capture.
//
// resolvedSettings: An object describing the settings used for this capture. Match this
// object’s [UniqueID] value to the [UniqueID] property of the photo
// settings object you initiated capture with to determine which capture
// request this delegate call corresponds to. You can also use this object to
// find out which values the photo output has chosen for automatic settings.
//
// # Discussion
// 
// The photo output calls this method as soon as the first step of capture
// ends—that is, at the end of the photographic exposure time.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoCaptureDelegate/photoOutput(_:didCapturePhotoFor:)
func (o AVCapturePhotoCaptureDelegateObject) CaptureOutputDidCapturePhotoForResolvedSettings(output IAVCapturePhotoOutput, resolvedSettings IAVCaptureResolvedPhotoSettings) {
	objc.Send[struct{}](o.ID, objc.Sel("captureOutput:didCapturePhotoForResolvedSettings:"), output, resolvedSettings)
	}
// Notifies the delegate that the capture process is complete.
//
// output: The photo output performing the capture.
//
// resolvedSettings: An object describing the settings used for this capture. Match this
// object’s [UniqueID] value to the [UniqueID] property of the photo
// settings object you initiated capture with to determine which capture
// request this delegate call corresponds to. You can also use this object to
// find out which values the photo output has chosen for automatic settings.
//
// error: If the capture process did not complete successfully, an error object
// describing the failure; otherwise, `nil`.
//
// # Discussion
// 
// The photo output calls this method when the entire capture process has
// finished, and no more delegate messages will be sent for this capture
// request. Use this time to clean up any resources you’ve allocated that
// relate to this capture request.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoCaptureDelegate/photoOutput(_:didFinishCaptureFor:error:)
func (o AVCapturePhotoCaptureDelegateObject) CaptureOutputDidFinishCaptureForResolvedSettingsError(output IAVCapturePhotoOutput, resolvedSettings IAVCaptureResolvedPhotoSettings, error_ foundation.INSError) {
	objc.Send[struct{}](o.ID, objc.Sel("captureOutput:didFinishCaptureForResolvedSettings:error:"), output, resolvedSettings, error_)
	}
// Provides the delegate with the captured image and associated metadata
// resulting from a photo capture.
//
// output: The photo output performing the capture.
//
// photo: An object containing the captured image pixel buffer, along with any
// metadata and attachments captured along with the photo (such as a preview
// image or depth map).
// 
// This parameter is always non-`nil`: if an error prevented successful
// capture, this object still contains metadata for the intended capture.
//
// error: If the capture process could not proceed successfully, an error object
// describing the failure; otherwise, `nil`.
//
// # Discussion
// 
// Use this method to receive the results of photo capture regardless of
// format.
// 
// The photo output calls this method once for each primary image to be
// delivered in a capture request. If you request capture in both RAW and
// processed formats, this method fires once for each format. If you request a
// bracketed capture with multiple exposures, this method fires once for each
// exposure.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoCaptureDelegate/photoOutput(_:didFinishProcessingPhoto:error:)
func (o AVCapturePhotoCaptureDelegateObject) CaptureOutputDidFinishProcessingPhotoError(output IAVCapturePhotoOutput, photo IAVCapturePhoto, error_ foundation.INSError) {
	objc.Send[struct{}](o.ID, objc.Sel("captureOutput:didFinishProcessingPhoto:error:"), output, photo, error_)
	}
// Notifies the delegate that the movie content of a Live Photo has finished
// recording.
//
// output: The photo output performing the capture.
//
// outputFileURL: The file URL at which the Live Photo movie will be written.
//
// resolvedSettings: An object describing the settings used for this capture. Match this
// object’s [UniqueID] value to the [UniqueID] property of the photo
// settings object you initiated capture with to determine which capture
// request this delegate call corresponds to. You can also use this object to
// find out which values the photo output has chosen for automatic settings.
//
// # Discussion
// 
// The photo output calls this method as soon as it has captured all movie
// data for a Live Photo. However, at this moment, that media content has not
// yet been processed or written to storage. (To be notified when the complete
// movie file has finished writing and is ready for consumption, implement the
// [CaptureOutputDidFinishProcessingLivePhotoToMovieFileAtURLDurationPhotoDisplayTimeResolvedSettingsError]
// method.)
// 
// Use this method to determine when it is appropriate to change your
// displayed UI to indicate that Live Photo movie capture is no longer in
// progress. For example, the Camera app displays a “LIVE” icon when the
// user presses the shutter button, then hides that icon when movie capture
// ends.
// 
// The photo output calls this method only once for each Live Photo capture.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoCaptureDelegate/photoOutput(_:didFinishRecordingLivePhotoMovieForEventualFileAt:resolvedSettings:)
func (o AVCapturePhotoCaptureDelegateObject) CaptureOutputDidFinishRecordingLivePhotoMovieForEventualFileAtURLResolvedSettings(output IAVCapturePhotoOutput, outputFileURL foundation.INSURL, resolvedSettings IAVCaptureResolvedPhotoSettings) {
	objc.Send[struct{}](o.ID, objc.Sel("captureOutput:didFinishRecordingLivePhotoMovieForEventualFileAtURL:resolvedSettings:"), output, outputFileURL, resolvedSettings)
	}
// Provides the delegate the movie file URL resulting from a Live Photo
// capture.
//
// output: The photo output performing the capture.
//
// outputFileURL: The file URL at which the movie content of the Live Photo was written.
//
// duration: The duration of the Live Photo movie.
//
// photoDisplayTime: The timestamp within the movie to which the still image part of the Live
// Photo corresponds.
//
// resolvedSettings: An object describing the settings used for this capture. Match this
// object’s [UniqueID] value to the [UniqueID] property of the photo
// settings object you initiated capture with to determine which capture
// request this delegate call corresponds to. You can also use this object to
// find out which values the photo output has chosen for automatic settings.
//
// error: If the capture process could not proceed successfully, an error object
// describing the failure; otherwise, `nil`.
//
// # Discussion
// 
// Use this method to receive the results of a Live Photo capture. When the
// photo output calls this method, the movie component of the Live Photo has
// been written to the location specified by the `outputFileURL` parameter and
// the Live Photo is ready for consumption. (To receive the still image
// component of the Live Photo, implement the
// [CaptureOutputDidFinishProcessingPhotoSampleBufferPreviewPhotoSampleBufferResolvedSettingsBracketSettingsError]
// method.)
// 
// You don’t need to implement this method if you’re not requesting Live
// Photo capture.
// 
// The photo output calls this method only once for each Live Photo capture.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoCaptureDelegate/photoOutput(_:didFinishProcessingLivePhotoToMovieFileAt:duration:photoDisplayTime:resolvedSettings:error:)
func (o AVCapturePhotoCaptureDelegateObject) CaptureOutputDidFinishProcessingLivePhotoToMovieFileAtURLDurationPhotoDisplayTimeResolvedSettingsError(output IAVCapturePhotoOutput, outputFileURL foundation.INSURL, duration uintptr, photoDisplayTime uintptr, resolvedSettings IAVCaptureResolvedPhotoSettings, error_ foundation.INSError) {
	objc.Send[struct{}](o.ID, objc.Sel("captureOutput:didFinishProcessingLivePhotoToMovieFileAtURL:duration:photoDisplayTime:resolvedSettings:error:"), output, outputFileURL, duration, photoDisplayTime, resolvedSettings, error_)
	}
// Tells the delegate when the system finishes capturing the photo proxy.
//
// output: The output instance.
//
// deferredPhotoProxy: A [AVCaptureDeferredPhotoProxy] instance that contains a proxy
// [CVPixelBuffer] as a placeholder for the final image.
// //
// [AVCaptureDeferredPhotoProxy]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDeferredPhotoProxy
// [CVPixelBuffer]: https://developer.apple.com/documentation/CoreVideo/cvpixelbuffer-q2e
//
// error: If the system couldn’t create the photo proxy, or any of the underlying
// intermediate files, an error object that describes the failure.
//
// # Discussion
// 
// You can use the output’s [FileDataRepresentation] with
// [PHAssetCreationRequest] to eventually produce the final, processed photo
// into the user’s Photo Library. Add the in-memory proxy file data
// representation to the photo library as quickly as possible after this call
// to ensure that the photo library can begin background processing. It’s
// also important so that the intermediates aren’t removed by a periodic
// clean-up job looking for abandoned intermediates produced by using the
// deferred photo processing APIs.
// 
// Your delegate implementation must adopt this method to opt into deferred
// photo processing, otherwise calling [CapturePhotoWithSettingsDelegate]
// throws an exception.
//
// [PHAssetCreationRequest]: https://developer.apple.com/documentation/Photos/PHAssetCreationRequest
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoCaptureDelegate/photoOutput(_:didFinishCapturingDeferredPhotoProxy:error:)
func (o AVCapturePhotoCaptureDelegateObject) CaptureOutputDidFinishCapturingDeferredPhotoProxyError(output IAVCapturePhotoOutput, deferredPhotoProxy objectivec.IObject, error_ foundation.INSError) {
	objc.Send[struct{}](o.ID, objc.Sel("captureOutput:didFinishCapturingDeferredPhotoProxy:error:"), output, deferredPhotoProxy, error_)
	}

// AVCapturePhotoCaptureDelegateConfig holds optional typed callbacks for [AVCapturePhotoCaptureDelegate] methods.
// Set non-nil fields to register the corresponding Objective-C delegate method.
// Methods with nil callbacks are not registered, so [NSObject.RespondsToSelector]
// returns false for them — matching the Objective-C delegate pattern exactly.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/avfoundation/avcapturephotocapturedelegate
type AVCapturePhotoCaptureDelegateConfig struct {

	// Other Methods
	// CaptureOutputWillBeginCaptureForResolvedSettings — Notifies the delegate that the capture output has resolved settings and will soon begin its capture process.
	CaptureOutputWillBeginCaptureForResolvedSettings func(output AVCapturePhotoOutput, resolvedSettings AVCaptureResolvedPhotoSettings)
	// CaptureOutputWillCapturePhotoForResolvedSettings — Notifies the delegate that photo capture is about to occur.
	CaptureOutputWillCapturePhotoForResolvedSettings func(output AVCapturePhotoOutput, resolvedSettings AVCaptureResolvedPhotoSettings)
	// CaptureOutputDidCapturePhotoForResolvedSettings — Notifies the delegate that the photo has been taken.
	CaptureOutputDidCapturePhotoForResolvedSettings func(output AVCapturePhotoOutput, resolvedSettings AVCaptureResolvedPhotoSettings)
	// CaptureOutputDidFinishCaptureForResolvedSettingsError — Notifies the delegate that the capture process is complete.
	CaptureOutputDidFinishCaptureForResolvedSettingsError func(output AVCapturePhotoOutput, resolvedSettings AVCaptureResolvedPhotoSettings, error_ foundation.NSError)
	// CaptureOutputDidFinishProcessingPhotoError — Provides the delegate with the captured image and associated metadata resulting from a photo capture.
	CaptureOutputDidFinishProcessingPhotoError func(output AVCapturePhotoOutput, photo AVCapturePhoto, error_ foundation.NSError)
}

// NewAVCapturePhotoCaptureDelegate creates an Objective-C object implementing the [AVCapturePhotoCaptureDelegate] protocol.
//
// Each call registers a unique Objective-C class containing only the methods
// set in config. This means [NSObject.RespondsToSelector] works correctly
// for optional delegate methods — only non-nil callbacks are registered.
//
// The returned [AVCapturePhotoCaptureDelegateObject] satisfies the [AVCapturePhotoCaptureDelegate] interface
// and can be passed directly to SetDelegate and similar methods.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/avfoundation/avcapturephotocapturedelegate
func NewAVCapturePhotoCaptureDelegate(config AVCapturePhotoCaptureDelegateConfig) AVCapturePhotoCaptureDelegateObject {
	n := delegateClassCounter.Add(1)
	className := fmt.Sprintf("GoAVCapturePhotoCaptureDelegate_%d", n)

	var methods []objc.MethodDef

	if config.CaptureOutputWillBeginCaptureForResolvedSettings != nil {
		fn := config.CaptureOutputWillBeginCaptureForResolvedSettings
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("captureOutput:willBeginCaptureForResolvedSettings:"),
			Fn: func(self objc.ID, _cmd objc.SEL, outputID objc.ID, resolvedSettingsID objc.ID) {
				output := AVCapturePhotoOutputFromID(outputID)
				resolvedSettings := AVCaptureResolvedPhotoSettingsFromID(resolvedSettingsID)
				fn(output, resolvedSettings)
			},
		})
	}

	if config.CaptureOutputWillCapturePhotoForResolvedSettings != nil {
		fn := config.CaptureOutputWillCapturePhotoForResolvedSettings
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("captureOutput:willCapturePhotoForResolvedSettings:"),
			Fn: func(self objc.ID, _cmd objc.SEL, outputID objc.ID, resolvedSettingsID objc.ID) {
				output := AVCapturePhotoOutputFromID(outputID)
				resolvedSettings := AVCaptureResolvedPhotoSettingsFromID(resolvedSettingsID)
				fn(output, resolvedSettings)
			},
		})
	}

	if config.CaptureOutputDidCapturePhotoForResolvedSettings != nil {
		fn := config.CaptureOutputDidCapturePhotoForResolvedSettings
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("captureOutput:didCapturePhotoForResolvedSettings:"),
			Fn: func(self objc.ID, _cmd objc.SEL, outputID objc.ID, resolvedSettingsID objc.ID) {
				output := AVCapturePhotoOutputFromID(outputID)
				resolvedSettings := AVCaptureResolvedPhotoSettingsFromID(resolvedSettingsID)
				fn(output, resolvedSettings)
			},
		})
	}

	if config.CaptureOutputDidFinishCaptureForResolvedSettingsError != nil {
		fn := config.CaptureOutputDidFinishCaptureForResolvedSettingsError
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("captureOutput:didFinishCaptureForResolvedSettings:error:"),
			Fn: func(self objc.ID, _cmd objc.SEL, outputID objc.ID, resolvedSettingsID objc.ID, error_ID objc.ID) {
				output := AVCapturePhotoOutputFromID(outputID)
				resolvedSettings := AVCaptureResolvedPhotoSettingsFromID(resolvedSettingsID)
				error_ := foundation.NSErrorFromID(error_ID)
				fn(output, resolvedSettings, error_)
			},
		})
	}

	if config.CaptureOutputDidFinishProcessingPhotoError != nil {
		fn := config.CaptureOutputDidFinishProcessingPhotoError
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("captureOutput:didFinishProcessingPhoto:error:"),
			Fn: func(self objc.ID, _cmd objc.SEL, outputID objc.ID, photoID objc.ID, error_ID objc.ID) {
				output := AVCapturePhotoOutputFromID(outputID)
				photo := AVCapturePhotoFromID(photoID)
				error_ := foundation.NSErrorFromID(error_ID)
				fn(output, photo, error_)
			},
		})
	}

	nsObjectClass := objc.GetClass("NSObject")
	proto := objc.GetProtocol("AVCapturePhotoCaptureDelegate")

	var protocols []*objc.Protocol
	if proto != nil {
		protocols = append(protocols, proto)
	}

	cls, err := objc.RegisterClass(className, nsObjectClass, protocols, nil, methods)
	if err != nil {
		panic(fmt.Sprintf("NewAVCapturePhotoCaptureDelegate: RegisterClass %s: %v", className, err))
	}

	instance := objc.ID(cls).Send(objc.RegisterName("alloc")).Send(objc.RegisterName("init"))
	return AVCapturePhotoCaptureDelegateObjectFromID(instance)
}

