// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// A protocol that defines the methods custom video compositors must implement.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVVideoCompositing
type AVVideoCompositing interface {
	objectivec.IObject

	// Tells the compositor that the composition changed render contexts.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVVideoCompositing/renderContextChanged(_:)
	RenderContextChanged(newRenderContext IAVVideoCompositionRenderContext)

	// Directs a custom video compositor object to create a new pixel buffer composed asynchronously from a collection of sources.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVVideoCompositing/startRequest(_:)
	StartVideoCompositionRequest(asyncVideoCompositionRequest IAVAsynchronousVideoCompositionRequest)

	// The pixel buffer attributes that the compositor accepts for source frames.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVVideoCompositing/sourcePixelBufferAttributes
	SourcePixelBufferAttributes() foundation.INSDictionary

	// The pixel buffer attributes that the compositor requires for pixel buffers that it creates.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVVideoCompositing/requiredPixelBufferAttributesForRenderContext
	RequiredPixelBufferAttributesForRenderContext() foundation.INSDictionary

	// A Boolean value that indicates whether the compositor handles source frames that contain high dynamic range (HDR) properties.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVVideoCompositing/supportsHDRSourceFrames
	SupportsHDRSourceFrames() bool

	// A Boolean value that indicates whether the compositor handles source frames that contains wide color properties.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVVideoCompositing/supportsWideColorSourceFrames
	SupportsWideColorSourceFrames() bool

	// A Boolean value that indicates whether the compositor conforms the color space of source frames to the composition color space.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVVideoCompositing/canConformColorOfSourceFrames
	CanConformColorOfSourceFrames() bool

	// supportsSourceTaggedBuffers protocol.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVVideoCompositing/supportsSourceTaggedBuffers
	SupportsSourceTaggedBuffers() bool
}

// AVVideoCompositingObject wraps an existing Objective-C object that conforms to the AVVideoCompositing protocol.
type AVVideoCompositingObject struct {
	objectivec.Object
}

func (o AVVideoCompositingObject) BaseObject() objectivec.Object {
	return o.Object
}

// AVVideoCompositingObjectFromID constructs a [AVVideoCompositingObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func AVVideoCompositingObjectFromID(id objc.ID) AVVideoCompositingObject {
	return AVVideoCompositingObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Tells the compositor that the composition changed render contexts.
//
// newRenderContext: The new render context of the video composition.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVVideoCompositing/renderContextChanged(_:)
func (o AVVideoCompositingObject) RenderContextChanged(newRenderContext IAVVideoCompositionRenderContext) {
	objc.Send[struct{}](o.ID, objc.Sel("renderContextChanged:"), newRenderContext)
}

// Directs a custom video compositor object to create a new pixel buffer
// composed asynchronously from a collection of sources.
//
// asyncVideoCompositionRequest: An instance of [AVAsynchronousVideoCompositionRequest] that provides
// context for the requested composition.
//
// # Discussion
//
// The custom compositor is expected to invoke, either subsequently or
// immediately, the `asyncVideoCompositionRequest` object’s
// [finish(withComposedVideoFrame:)] or [FinishWithError] methods.
//
// If you intend to finish rendering the frame after handling of this message
// returns, you must retain `asyncVideoCompositionRequest` until after
// composition is finished.
//
// Note that if the custom compositor’s implementation of this method
// returns without finishing the composition immediately, it may be invoked
// again with another composition request before the prior request is
// finished; in such cases the custom compositor should be prepared to manage
// multiple composition requests.
//
// If the rendered frame is exactly the same as one of the source frames, with
// no letterboxing, pillboxing or cropping needed, then the appropriate source
// pixel buffer may be returned, after [CFRetain] has been called on it).
//
// See: https://developer.apple.com/documentation/AVFoundation/AVVideoCompositing/startRequest(_:)
//
// [CFRetain]: https://developer.apple.com/documentation/CoreFoundation/CFRetain
// [finish(withComposedVideoFrame:)]: https://developer.apple.com/documentation/AVFoundation/AVAsynchronousVideoCompositionRequest/finish(withComposedVideoFrame:)
func (o AVVideoCompositingObject) StartVideoCompositionRequest(asyncVideoCompositionRequest IAVAsynchronousVideoCompositionRequest) {
	objc.Send[struct{}](o.ID, objc.Sel("startVideoCompositionRequest:"), asyncVideoCompositionRequest)
}

// Informs a custom video compositor about upcoming rendering requests.
//
// renderHint: Information about the upcoming composition requests.
//
// # Discussion
//
// In this method, the compositor can load composition resources, such as
// overlay images, that will be needed in the anticipated rendering time
// range.
//
// Unlike the [StartVideoCompositionRequest] method, which is invoked only
// when the frame compositing is necessary, this method is typically called
// every frame duration. It allows the custom compositor to load and unload a
// composition resource such as overlay images at an appropriate time.
//
// In forward playback, the render hint’s [StartCompositionTime] is less
// than its [EndCompositionTime]. In reverse playback, its
// [EndCompositionTime] is less than its [StartCompositionTime]. For seeking,
// the two values are equivalent, which means the upcoming composition request
// time range is unknown.
//
// This method is guaranteed to be called before
// [StartVideoCompositionRequest] for a given composition time.
//
// This method is synchronous. Make sure that your implementation returns
// quickly to ensure that playback doesn’t stall and cause frame drops.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVVideoCompositing/anticipateRendering(using:)
func (o AVVideoCompositingObject) AnticipateRenderingUsingHint(renderHint IAVVideoCompositionRenderHint) {
	objc.Send[struct{}](o.ID, objc.Sel("anticipateRenderingUsingHint:"), renderHint)
}

// Tells a custom video compositor to perform any work in the prerolling
// phase.
//
// renderHint: Information about the upcoming composition requests.
//
// # Discussion
//
// The AVFoundation framework may perform prerolling to load media data to
// prime the render pipelines for smoother playback. This method is called in
// the prerolling phase so that the compositor can load composition resources,
// such as overlay images, that will be needed as soon as the playback starts.
//
// Not all rendering scenarios use prerolling. For example, this method
// won’t be called during seeking.
//
// If this method is called, it is guaranteed to be invoked before the first
// [StartVideoCompositionRequest] call.
//
// This method is synchronous. The prerolling won’t finish until the method
// returns.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVVideoCompositing/prerollForRendering(using:)
func (o AVVideoCompositingObject) PrerollForRenderingUsingHint(renderHint IAVVideoCompositionRenderHint) {
	objc.Send[struct{}](o.ID, objc.Sel("prerollForRenderingUsingHint:"), renderHint)
}

// Directs a custom video compositor object to cancel or finish all pending
// video composition requests.
//
// # Discussion
//
// Upon receiving this message, a custom video compositor must block until it
// has either cancelled all pending frame requests, and called the
// [FinishCancelledRequest] method for each of them. If cancellation is not
// possible, the method must block until it has finished processing of all the
// frames and called the [finish(withComposedVideoFrame:)] method for each of
// them.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVVideoCompositing/cancelAllPendingVideoCompositionRequests()
//
// [finish(withComposedVideoFrame:)]: https://developer.apple.com/documentation/AVFoundation/AVAsynchronousVideoCompositionRequest/finish(withComposedVideoFrame:)
func (o AVVideoCompositingObject) CancelAllPendingVideoCompositionRequests() {
	objc.Send[struct{}](o.ID, objc.Sel("cancelAllPendingVideoCompositionRequests"))
}

// The pixel buffer attributes that the compositor accepts for source frames.
//
// # Discussion
//
// The property is required to provide a [kCVPixelBufferPixelFormatTypeKey]
// key in the dictionary, along with attributes for which the compositor needs
// specific values to work properly. Omitted attributes will be supplied by
// the composition engine to allow for the best performance. If the attribute
// [kCVPixelBufferPixelFormatTypeKey] key is not in the dictionary an
// exception will be raised. The value of the
// [kCVPixelBufferPixelFormatTypeKey] is an array of `kCVPixelFormatType_*`
// constants as defined in Pixel_Format_Types.
//
// If the custom compositor is meant to be used with an
// [AVVideoCompositionCoreAnimationTool] created using the
// [VideoCompositionCoreAnimationToolWithAdditionalLayerAsTrackID] method,
// [KCVPixelFormatType_32BGRA] should be included as one of the supported
// pixel format types.
//
// Missing attributes will be set by the composition engine to values allowing
// the best performance.
//
// This property is queried once before any composition request is sent to the
// compositor. Changing source buffer attributes afterwards is not supported.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVVideoCompositing/sourcePixelBufferAttributes
//
// [kCVPixelBufferPixelFormatTypeKey]: https://developer.apple.com/documentation/CoreVideo/kCVPixelBufferPixelFormatTypeKey
func (o AVVideoCompositingObject) SourcePixelBufferAttributes() foundation.INSDictionary {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("sourcePixelBufferAttributes"))
	return foundation.NSDictionaryFromID(rv)
}

// The pixel buffer attributes that the compositor requires for pixel buffers
// that it creates.
//
// # Discussion
//
// The property is required to provide a [kCVPixelBufferPixelFormatTypeKey]
// key in the dictionary, along with attributes for which the compositor needs
// specific values to work properly. Omitted attributes will be supplied by
// the composition engine to allow for the best performance. If the attribute
// [kCVPixelBufferPixelFormatTypeKey] key is not in the dictionary an
// exception will be raised. The value of the
// [kCVPixelBufferPixelFormatTypeKey] is an array of `kCVPixelFormatType_*`
// constants as defined in Pixel_Format_Types.
//
// The value of `requiredPixelBufferAttributesForRenderContext` is retrieved
// prior to the creation of a new render context; the combination of the
// attributes in the returned value and the additional attributes supplied by
// the composition engine will be used in the creation of subsequent render
// context’s pixelBuffers.
//
// This property is queried once before any composition request is sent to the
// compositor. Changing required buffer attributes afterwards is not
// supported.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVVideoCompositing/requiredPixelBufferAttributesForRenderContext
//
// [kCVPixelBufferPixelFormatTypeKey]: https://developer.apple.com/documentation/CoreVideo/kCVPixelBufferPixelFormatTypeKey
func (o AVVideoCompositingObject) RequiredPixelBufferAttributesForRenderContext() foundation.INSDictionary {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("requiredPixelBufferAttributesForRenderContext"))
	return foundation.NSDictionaryFromID(rv)
}

// A Boolean value that indicates whether the compositor handles source frames
// that contain high dynamic range (HDR) properties.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVVideoCompositing/supportsHDRSourceFrames
func (o AVVideoCompositingObject) SupportsHDRSourceFrames() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("supportsHDRSourceFrames"))
	return bool(rv)
}

// A Boolean value that indicates whether the compositor handles source frames
// that contains wide color properties.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVVideoCompositing/supportsWideColorSourceFrames
func (o AVVideoCompositingObject) SupportsWideColorSourceFrames() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("supportsWideColorSourceFrames"))
	return bool(rv)
}

// A Boolean value that indicates whether the compositor conforms the color
// space of source frames to the composition color space.
//
// # Discussion
//
// A custom compositor indicates its processing requirements through the
// [SourcePixelBufferAttributes] and [SupportsWideColorSourceFrames]
// properties. By default, the composition engine prepares source frames by
// converting them to meet the compositor’s configuration.
//
// When this property value is true, the engine doesn’t convert source pixel
// buffers that meet the compositor’s processing requirements. However, it
// does convert buffers that don’t meet the processing requirements, which
// includes the following cases:
//
// - The values of [SupportsWideColorSourceFrames] and
// [SupportsHDRSourceFrames] are false, but the source buffers contain wide
// color. In this case, the engine converts the color space of source pixel
// buffers to BT.709 color space. Note that when [SupportsHDRSourceFrames] is
// true, the engine also assumes [SupportsWideColorSourceFrames] is true. -
// The value of [SupportsHDRSourceFrames] is false and source buffers contain
// HDR color. In this case, the engine converts the color space of source
// pixel buffers to the composition color space. - The pixel format of the
// source buffers isn’t specified in [SourcePixelBufferAttributes]. In this
// case, the engine converts the pixel format to a supported format and
// converts the color space to the composition color space.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVVideoCompositing/canConformColorOfSourceFrames
func (o AVVideoCompositingObject) CanConformColorOfSourceFrames() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("canConformColorOfSourceFrames"))
	return bool(rv)
}

// See: https://developer.apple.com/documentation/AVFoundation/AVVideoCompositing/supportsSourceTaggedBuffers
func (o AVVideoCompositingObject) SupportsSourceTaggedBuffers() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("supportsSourceTaggedBuffers"))
	return bool(rv)
}
