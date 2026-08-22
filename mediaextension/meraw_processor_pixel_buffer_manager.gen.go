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

// The class instance for the [MERAWProcessorPixelBufferManager] class.
var (
	_MERAWProcessorPixelBufferManagerClass     MERAWProcessorPixelBufferManagerClass
	_MERAWProcessorPixelBufferManagerClassOnce sync.Once
)

func getMERAWProcessorPixelBufferManagerClass() MERAWProcessorPixelBufferManagerClass {
	_MERAWProcessorPixelBufferManagerClassOnce.Do(func() {
		_MERAWProcessorPixelBufferManagerClass = MERAWProcessorPixelBufferManagerClass{class: objc.GetClass("MERAWProcessorPixelBufferManager")}
	})
	return _MERAWProcessorPixelBufferManagerClass
}

// GetMERAWProcessorPixelBufferManagerClass returns the class object for MERAWProcessorPixelBufferManager.
func GetMERAWProcessorPixelBufferManagerClass() MERAWProcessorPixelBufferManagerClass {
	return getMERAWProcessorPixelBufferManagerClass()
}

type MERAWProcessorPixelBufferManagerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MERAWProcessorPixelBufferManagerClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MERAWProcessorPixelBufferManagerClass) Alloc() MERAWProcessorPixelBufferManager {
	rv := objc.Send[MERAWProcessorPixelBufferManager](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// Describes pixel buffer requirements and creates new pixel buffers.
//
// # Discussion
//
// It contains the interfaces that the [MERAWProcessor] uses for two tasks.
// First, to declare its set of requirements for output [CVPixelBuffer] in the
// form of a [pixelBufferAttributes] dictionary. Second, create pixel buffers
// that match processor output requirements and satisfy Video Toolbox and
// client requirements.
//
// # Creating a pixel buffer
//
//   - [MERAWProcessorPixelBufferManager.PixelBufferAttributes]: A dictionary that contains the attributes Video Toolbox uses to create a pixel buffer for the video RAW processor.
//   - [MERAWProcessorPixelBufferManager.SetPixelBufferAttributes]
//   - [MERAWProcessorPixelBufferManager.CreatePixelBufferAndReturnError]: Generates a pixel buffer using the session’s pixel buffer pool.
//
// See: https://developer.apple.com/documentation/MediaExtension/MERAWProcessorPixelBufferManager
//
// [CVPixelBuffer]: https://developer.apple.com/documentation/CoreVideo/cvpixelbuffer-q2e
// [pixelBufferAttributes]: https://developer.apple.com/documentation/MediaExtension/MERAWProcessorPixelBufferManager/pixelBufferAttributes-2cki6
type MERAWProcessorPixelBufferManager struct {
	objectivec.Object
}

// MERAWProcessorPixelBufferManagerFromID constructs a [MERAWProcessorPixelBufferManager] from an objc.ID.
//
// Describes pixel buffer requirements and creates new pixel buffers.
func MERAWProcessorPixelBufferManagerFromID(id objc.ID) MERAWProcessorPixelBufferManager {
	return MERAWProcessorPixelBufferManager{objectivec.Object{ID: id}}
}

// NOTE: MERAWProcessorPixelBufferManager adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MERAWProcessorPixelBufferManager] class.
//
// # Creating a pixel buffer
//
//   - [IMERAWProcessorPixelBufferManager.PixelBufferAttributes]: A dictionary that contains the attributes Video Toolbox uses to create a pixel buffer for the video RAW processor.
//   - [IMERAWProcessorPixelBufferManager.SetPixelBufferAttributes]
//   - [IMERAWProcessorPixelBufferManager.CreatePixelBufferAndReturnError]: Generates a pixel buffer using the session’s pixel buffer pool.
//
// See: https://developer.apple.com/documentation/MediaExtension/MERAWProcessorPixelBufferManager
type IMERAWProcessorPixelBufferManager interface {
	objectivec.IObject

	// Topic: Creating a pixel buffer

	// A dictionary that contains the attributes Video Toolbox uses to create a pixel buffer for the video RAW processor.
	PixelBufferAttributes() foundation.INSDictionary
	SetPixelBufferAttributes(value foundation.INSDictionary)
	// Generates a pixel buffer using the session’s pixel buffer pool.
	CreatePixelBufferAndReturnError() (corevideo.CVImageBufferRef, error)
}

// Init initializes the instance.
func (m MERAWProcessorPixelBufferManager) Init() MERAWProcessorPixelBufferManager {
	rv := objc.Send[MERAWProcessorPixelBufferManager](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MERAWProcessorPixelBufferManager) Autorelease() MERAWProcessorPixelBufferManager {
	rv := objc.Send[MERAWProcessorPixelBufferManager](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMERAWProcessorPixelBufferManager creates a new MERAWProcessorPixelBufferManager instance.
func NewMERAWProcessorPixelBufferManager() MERAWProcessorPixelBufferManager {
	class := getMERAWProcessorPixelBufferManagerClass()
	rv := objc.Send[MERAWProcessorPixelBufferManager](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Generates a pixel buffer using the session’s pixel buffer pool.
//
// # Return Value
//
// A pixel buffer that’s compatible with the extension’s most recently set
// pixel buffer attributes.
//
// See: https://developer.apple.com/documentation/MediaExtension/MERAWProcessorPixelBufferManager/makePixelBuffer()
func (m MERAWProcessorPixelBufferManager) CreatePixelBufferAndReturnError() (corevideo.CVImageBufferRef, error) {
	var errorPtr objc.ID
	rv := objc.Send[corevideo.CVImageBufferRef](m.ID, objc.Sel("createPixelBufferAndReturnError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return *new(corevideo.CVImageBufferRef), foundation.NSErrorFrom(errorPtr)
	}
	return rv, nil

}

// A dictionary that contains the attributes Video Toolbox uses to create a
// pixel buffer for the video RAW processor.
//
// See: https://developer.apple.com/documentation/mediaextension/merawprocessorpixelbuffermanager/pixelbufferattributes-4fe69
func (m MERAWProcessorPixelBufferManager) PixelBufferAttributes() foundation.INSDictionary {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("pixelBufferAttributes"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (m MERAWProcessorPixelBufferManager) SetPixelBufferAttributes(value foundation.INSDictionary) {
	objc.Send[struct{}](m.ID, objc.Sel("setPixelBufferAttributes:"), value)
}
