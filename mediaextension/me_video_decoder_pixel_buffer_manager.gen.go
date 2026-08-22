// Code generated from Apple documentation for MediaExtension. DO NOT EDIT.

package mediaextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/corevideo"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MEVideoDecoderPixelBufferManager] class.
var (
	_MEVideoDecoderPixelBufferManagerClass     MEVideoDecoderPixelBufferManagerClass
	_MEVideoDecoderPixelBufferManagerClassOnce sync.Once
)

func getMEVideoDecoderPixelBufferManagerClass() MEVideoDecoderPixelBufferManagerClass {
	_MEVideoDecoderPixelBufferManagerClassOnce.Do(func() {
		_MEVideoDecoderPixelBufferManagerClass = MEVideoDecoderPixelBufferManagerClass{class: objc.GetClass("MEVideoDecoderPixelBufferManager")}
	})
	return _MEVideoDecoderPixelBufferManagerClass
}

// GetMEVideoDecoderPixelBufferManagerClass returns the class object for MEVideoDecoderPixelBufferManager.
func GetMEVideoDecoderPixelBufferManagerClass() MEVideoDecoderPixelBufferManagerClass {
	return getMEVideoDecoderPixelBufferManagerClass()
}

type MEVideoDecoderPixelBufferManagerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MEVideoDecoderPixelBufferManagerClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MEVideoDecoderPixelBufferManagerClass) Alloc() MEVideoDecoderPixelBufferManager {
	rv := objc.Send[MEVideoDecoderPixelBufferManager](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// Describes pixel buffer requirements and creates new pixel buffers.
//
// # Discussion
//
// Contains the interfaces that the [MEVideoDecoder] uses for two tasks.
// First, to declare its set of requirements for output [CVPixelBuffer]
// objects in the form of a
// [MEVideoDecoderPixelBufferManager.PixelBufferAttributes] dictionary.
// Second, to create pixel buffers that match decoder output requirements but
// also satisfy [Video Toolbox] and client requirements.
//
// # Creating a pixel buffer
//
//   - [MEVideoDecoderPixelBufferManager.PixelBufferAttributes]: A dictionary that contains the attributes Video Toolbox uses to create a pixel buffer for the decoder.
//   - [MEVideoDecoderPixelBufferManager.SetPixelBufferAttributes]
//   - [MEVideoDecoderPixelBufferManager.CreatePixelBufferAndReturnError]: Generates a pixel buffer using the session’s pixel buffer pool.
//
// # Registering Custom Pixel Formats
//
//   - [MEVideoDecoderPixelBufferManager.RegisterCustomPixelFormat]
//
// See: https://developer.apple.com/documentation/MediaExtension/MEVideoDecoderPixelBufferManager
//
// [Video Toolbox]: https://developer.apple.com/documentation/VideoToolbox
type MEVideoDecoderPixelBufferManager struct {
	objectivec.Object
}

// MEVideoDecoderPixelBufferManagerFromID constructs a [MEVideoDecoderPixelBufferManager] from an objc.ID.
//
// Describes pixel buffer requirements and creates new pixel buffers.
func MEVideoDecoderPixelBufferManagerFromID(id objc.ID) MEVideoDecoderPixelBufferManager {
	return MEVideoDecoderPixelBufferManager{objectivec.Object{ID: id}}
}

// NOTE: MEVideoDecoderPixelBufferManager adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MEVideoDecoderPixelBufferManager] class.
//
// # Creating a pixel buffer
//
//   - [IMEVideoDecoderPixelBufferManager.PixelBufferAttributes]: A dictionary that contains the attributes Video Toolbox uses to create a pixel buffer for the decoder.
//   - [IMEVideoDecoderPixelBufferManager.SetPixelBufferAttributes]
//   - [IMEVideoDecoderPixelBufferManager.CreatePixelBufferAndReturnError]: Generates a pixel buffer using the session’s pixel buffer pool.
//
// # Registering Custom Pixel Formats
//
//   - [IMEVideoDecoderPixelBufferManager.RegisterCustomPixelFormat]
//
// See: https://developer.apple.com/documentation/MediaExtension/MEVideoDecoderPixelBufferManager
type IMEVideoDecoderPixelBufferManager interface {
	objectivec.IObject

	// Topic: Creating a pixel buffer

	// A dictionary that contains the attributes Video Toolbox uses to create a pixel buffer for the decoder.
	PixelBufferAttributes() foundation.INSDictionary
	SetPixelBufferAttributes(value foundation.INSDictionary)
	// Generates a pixel buffer using the session’s pixel buffer pool.
	CreatePixelBufferAndReturnError() (corevideo.CVImageBufferRef, error)

	// Topic: Registering Custom Pixel Formats

	RegisterCustomPixelFormat(customPixelFormat foundation.INSDictionary)
}

// Init initializes the instance.
func (m MEVideoDecoderPixelBufferManager) Init() MEVideoDecoderPixelBufferManager {
	rv := objc.Send[MEVideoDecoderPixelBufferManager](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MEVideoDecoderPixelBufferManager) Autorelease() MEVideoDecoderPixelBufferManager {
	rv := objc.Send[MEVideoDecoderPixelBufferManager](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMEVideoDecoderPixelBufferManager creates a new MEVideoDecoderPixelBufferManager instance.
func NewMEVideoDecoderPixelBufferManager() MEVideoDecoderPixelBufferManager {
	class := getMEVideoDecoderPixelBufferManagerClass()
	rv := objc.Send[MEVideoDecoderPixelBufferManager](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Generates a pixel buffer using the session’s pixel buffer pool.
//
// # Return Value
//
// A pixel buffer that’s compatible with the extension’s most recently set
// pixel buffer attributes.
//
// See: https://developer.apple.com/documentation/MediaExtension/MEVideoDecoderPixelBufferManager/makePixelBuffer()
func (m MEVideoDecoderPixelBufferManager) CreatePixelBufferAndReturnError() (corevideo.CVImageBufferRef, error) {
	var errorPtr objc.ID
	rv := objc.Send[corevideo.CVImageBufferRef](m.ID, objc.Sel("createPixelBufferAndReturnError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return *new(corevideo.CVImageBufferRef), foundation.NSErrorFrom(errorPtr)
	}
	return rv, nil

}

// customPixelFormat: A dictionary containing a set of keys and values as described in
// [CVPixelFormatDescription] suitable for providing as the ‘description’
// parameter to [CVPixelFormatDescriptionCreateWithPixelFormatType(_:_:)].
// This must contain the custom pixel format fourCC as the value for the
// `kCVPixelFormatCodecType` key.
//
// # Discussion
//
// This property is appropriate for decoders which produce output in a custom
// pixel format. This will generally only be used by decoders which produce
// RAW output, where the decoder’s output buffers will only be consumed by
// an [MERAWProcessor] extension which registers the same pixel format. The
// [MERAWProcessor] needs to manually register the custom pixel format using
// [CVPixelFormatDescriptionCreateWithPixelFormatType(_:_:)].
//
// See: https://developer.apple.com/documentation/MediaExtension/MEVideoDecoderPixelBufferManager/registerCustomPixelFormat(_:)
//
// [CVPixelFormatDescriptionCreateWithPixelFormatType(_:_:)]: https://developer.apple.com/documentation/CoreVideo/CVPixelFormatDescriptionCreateWithPixelFormatType(_:_:)
// [CVPixelFormatDescription]: https://developer.apple.com/documentation/CoreVideo/cvpixelformatdescription
//
// [CVPixelFormatDescriptionCreateWithPixelFormatType(_:_:)]: https://developer.apple.com/documentation/CoreVideo/CVPixelFormatDescriptionCreateWithPixelFormatType(_:_:)
func (m MEVideoDecoderPixelBufferManager) RegisterCustomPixelFormat(customPixelFormat foundation.INSDictionary) {
	objc.Send[objc.ID](m.ID, objc.Sel("registerCustomPixelFormat:"), customPixelFormat)
}

// A dictionary that contains the attributes Video Toolbox uses to create a
// pixel buffer for the decoder.
//
// # Discussion
//
// The decoder can update this dictionary before it requests a new pixel
// buffer.
//
// See: https://developer.apple.com/documentation/MediaExtension/MEVideoDecoderPixelBufferManager/pixelBufferAttributes
func (m MEVideoDecoderPixelBufferManager) PixelBufferAttributes() foundation.INSDictionary {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("pixelBufferAttributes"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (m MEVideoDecoderPixelBufferManager) SetPixelBufferAttributes(value foundation.INSDictionary) {
	objc.Send[struct{}](m.ID, objc.Sel("setPixelBufferAttributes:"), value)
}
