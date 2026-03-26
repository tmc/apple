// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVVideoCompositionRenderContext] class.
var (
	_AVVideoCompositionRenderContextClass     AVVideoCompositionRenderContextClass
	_AVVideoCompositionRenderContextClassOnce sync.Once
)

func getAVVideoCompositionRenderContextClass() AVVideoCompositionRenderContextClass {
	_AVVideoCompositionRenderContextClassOnce.Do(func() {
		_AVVideoCompositionRenderContextClass = AVVideoCompositionRenderContextClass{class: objc.GetClass("AVVideoCompositionRenderContext")}
	})
	return _AVVideoCompositionRenderContextClass
}

// GetAVVideoCompositionRenderContextClass returns the class object for AVVideoCompositionRenderContext.
func GetAVVideoCompositionRenderContextClass() AVVideoCompositionRenderContextClass {
	return getAVVideoCompositionRenderContextClass()
}

type AVVideoCompositionRenderContextClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVVideoCompositionRenderContextClass) Alloc() AVVideoCompositionRenderContext {
	rv := objc.Send[AVVideoCompositionRenderContext](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that defines the context in which custom compositors render pixel
// buffers.
//
// # Overview
// 
// A render context provides size and scaling information and offers a service
// for efficiently providing pixel buffers from a managed pool of buffers.
//
// # Getting the render settings
//
//   - [AVVideoCompositionRenderContext.VideoComposition]: The video composition being rendered.
//   - [AVVideoCompositionRenderContext.HighQualityRendering]: The rendering quality to use.
//   - [AVVideoCompositionRenderContext.RenderScale]: A scaling ratio that is applied when rendering frames.
//   - [AVVideoCompositionRenderContext.RenderTransform]: A transform to apply to the source image.
//   - [AVVideoCompositionRenderContext.Size]: The width and height for the rendering frames.
//
// # Getting pixel and edge width information
//
//   - [AVVideoCompositionRenderContext.EdgeWidths]: The width of the edge processing region on the left, top, right, and bottom edges, in pixels.
//   - [AVVideoCompositionRenderContext.PixelAspectRatio]: The pixel aspect ratio for rendered frames.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVVideoCompositionRenderContext
type AVVideoCompositionRenderContext struct {
	objectivec.Object
}

// AVVideoCompositionRenderContextFromID constructs a [AVVideoCompositionRenderContext] from an objc.ID.
//
// An object that defines the context in which custom compositors render pixel
// buffers.
func AVVideoCompositionRenderContextFromID(id objc.ID) AVVideoCompositionRenderContext {
	return AVVideoCompositionRenderContext{objectivec.Object{ID: id}}
}
// NOTE: AVVideoCompositionRenderContext adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVVideoCompositionRenderContext] class.
//
// # Getting the render settings
//
//   - [IAVVideoCompositionRenderContext.VideoComposition]: The video composition being rendered.
//   - [IAVVideoCompositionRenderContext.HighQualityRendering]: The rendering quality to use.
//   - [IAVVideoCompositionRenderContext.RenderScale]: A scaling ratio that is applied when rendering frames.
//   - [IAVVideoCompositionRenderContext.RenderTransform]: A transform to apply to the source image.
//   - [IAVVideoCompositionRenderContext.Size]: The width and height for the rendering frames.
//
// # Getting pixel and edge width information
//
//   - [IAVVideoCompositionRenderContext.EdgeWidths]: The width of the edge processing region on the left, top, right, and bottom edges, in pixels.
//   - [IAVVideoCompositionRenderContext.PixelAspectRatio]: The pixel aspect ratio for rendered frames.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVVideoCompositionRenderContext
type IAVVideoCompositionRenderContext interface {
	objectivec.IObject

	// Topic: Getting the render settings

	// The video composition being rendered.
	VideoComposition() IAVVideoComposition
	// The rendering quality to use.
	HighQualityRendering() bool
	// A scaling ratio that is applied when rendering frames.
	RenderScale() float32
	// A transform to apply to the source image.
	RenderTransform() corefoundation.CGAffineTransform
	// The width and height for the rendering frames.
	Size() corefoundation.CGSize

	// Topic: Getting pixel and edge width information

	// The width of the edge processing region on the left, top, right, and bottom edges, in pixels.
	EdgeWidths() AVEdgeWidths
	// The pixel aspect ratio for rendered frames.
	PixelAspectRatio() AVPixelAspectRatio
}

// Init initializes the instance.
func (v AVVideoCompositionRenderContext) Init() AVVideoCompositionRenderContext {
	rv := objc.Send[AVVideoCompositionRenderContext](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v AVVideoCompositionRenderContext) Autorelease() AVVideoCompositionRenderContext {
	rv := objc.Send[AVVideoCompositionRenderContext](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVVideoCompositionRenderContext creates a new AVVideoCompositionRenderContext instance.
func NewAVVideoCompositionRenderContext() AVVideoCompositionRenderContext {
	class := getAVVideoCompositionRenderContextClass()
	rv := objc.Send[AVVideoCompositionRenderContext](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The video composition being rendered.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVVideoCompositionRenderContext/videoComposition
func (v AVVideoCompositionRenderContext) VideoComposition() IAVVideoComposition {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("videoComposition"))
	return AVVideoCompositionFromID(objc.ID(rv))
}
// The rendering quality to use.
//
// # Discussion
// 
// Specifies that the custom compositor should use higher quality, potentially
// slower algorithms.
// 
// Generally this property is [true] for non-real-time use cases.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AVFoundation/AVVideoCompositionRenderContext/highQualityRendering
func (v AVVideoCompositionRenderContext) HighQualityRendering() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("highQualityRendering"))
	return rv
}
// A scaling ratio that is applied when rendering frames.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVVideoCompositionRenderContext/renderScale
func (v AVVideoCompositionRenderContext) RenderScale() float32 {
	rv := objc.Send[float32](v.ID, objc.Sel("renderScale"))
	return rv
}
// A transform to apply to the source image.
//
// # Discussion
// 
// The transform to apply to the source image incorporating the [RenderScale],
// [PixelAspectRatio], and [EdgeWidths].
// 
// The coordinate system origin is the top left corner of the buffer.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVVideoCompositionRenderContext/renderTransform
func (v AVVideoCompositionRenderContext) RenderTransform() corefoundation.CGAffineTransform {
	rv := objc.Send[corefoundation.CGAffineTransform](v.ID, objc.Sel("renderTransform"))
	return corefoundation.CGAffineTransform(rv)
}
// The width and height for the rendering frames.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVVideoCompositionRenderContext/size
func (v AVVideoCompositionRenderContext) Size() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](v.ID, objc.Sel("size"))
	return corefoundation.CGSize(rv)
}
// The width of the edge processing region on the left, top, right, and bottom
// edges, in pixels.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVVideoCompositionRenderContext/edgeWidths
func (v AVVideoCompositionRenderContext) EdgeWidths() AVEdgeWidths {
	rv := objc.Send[AVEdgeWidths](v.ID, objc.Sel("edgeWidths"))
	return AVEdgeWidths(rv)
}
// The pixel aspect ratio for rendered frames.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVVideoCompositionRenderContext/pixelAspectRatio
func (v AVVideoCompositionRenderContext) PixelAspectRatio() AVPixelAspectRatio {
	rv := objc.Send[AVPixelAspectRatio](v.ID, objc.Sel("pixelAspectRatio"))
	return AVPixelAspectRatio(rv)
}

