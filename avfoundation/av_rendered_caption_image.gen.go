// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVRenderedCaptionImage] class.
var (
	_AVRenderedCaptionImageClass     AVRenderedCaptionImageClass
	_AVRenderedCaptionImageClassOnce sync.Once
)

func getAVRenderedCaptionImageClass() AVRenderedCaptionImageClass {
	_AVRenderedCaptionImageClassOnce.Do(func() {
		_AVRenderedCaptionImageClass = AVRenderedCaptionImageClass{class: objc.GetClass("AVRenderedCaptionImage")}
	})
	return _AVRenderedCaptionImageClass
}

// GetAVRenderedCaptionImageClass returns the class object for AVRenderedCaptionImage.
func GetAVRenderedCaptionImageClass() AVRenderedCaptionImageClass {
	return getAVRenderedCaptionImageClass()
}

type AVRenderedCaptionImageClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVRenderedCaptionImageClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVRenderedCaptionImageClass) Alloc() AVRenderedCaptionImage {
	rv := objc.Send[AVRenderedCaptionImage](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that provides a rendered pixel buffer and its position in pixels.
//
// # Inspecting the image
//
//   - [AVRenderedCaptionImage.ReadOnlyPixelBuffer]: A CVReadOnlyPixelBuffer that contains pixel data for the rendered caption
//   - [AVRenderedCaptionImage.SetReadOnlyPixelBuffer]
//   - [AVRenderedCaptionImage.Position]: A point that defines the position, in pixels, of the rendered caption image relative to the video frame.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVRenderedCaptionImage
type AVRenderedCaptionImage struct {
	objectivec.Object
}

// AVRenderedCaptionImageFromID constructs a [AVRenderedCaptionImage] from an objc.ID.
//
// An object that provides a rendered pixel buffer and its position in pixels.
func AVRenderedCaptionImageFromID(id objc.ID) AVRenderedCaptionImage {
	return AVRenderedCaptionImage{objectivec.Object{ID: id}}
}
// NOTE: AVRenderedCaptionImage adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVRenderedCaptionImage] class.
//
// # Inspecting the image
//
//   - [IAVRenderedCaptionImage.ReadOnlyPixelBuffer]: A CVReadOnlyPixelBuffer that contains pixel data for the rendered caption
//   - [IAVRenderedCaptionImage.SetReadOnlyPixelBuffer]
//   - [IAVRenderedCaptionImage.Position]: A point that defines the position, in pixels, of the rendered caption image relative to the video frame.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVRenderedCaptionImage
type IAVRenderedCaptionImage interface {
	objectivec.IObject

	// Topic: Inspecting the image

	// A CVReadOnlyPixelBuffer that contains pixel data for the rendered caption
	ReadOnlyPixelBuffer() objectivec.IObject
	SetReadOnlyPixelBuffer(value objectivec.IObject)
	// A point that defines the position, in pixels, of the rendered caption image relative to the video frame.
	Position() corefoundation.CGPoint
}

// Init initializes the instance.
func (r AVRenderedCaptionImage) Init() AVRenderedCaptionImage {
	rv := objc.Send[AVRenderedCaptionImage](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r AVRenderedCaptionImage) Autorelease() AVRenderedCaptionImage {
	rv := objc.Send[AVRenderedCaptionImage](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVRenderedCaptionImage creates a new AVRenderedCaptionImage instance.
func NewAVRenderedCaptionImage() AVRenderedCaptionImage {
	class := getAVRenderedCaptionImageClass()
	rv := objc.Send[AVRenderedCaptionImage](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// A CVReadOnlyPixelBuffer that contains pixel data for the rendered caption
//
// See: https://developer.apple.com/documentation/avfoundation/avrenderedcaptionimage/readonlypixelbuffer
func (r AVRenderedCaptionImage) ReadOnlyPixelBuffer() objectivec.IObject {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("readOnlyPixelBuffer"))
	return objectivec.Object{ID: rv}
}
func (r AVRenderedCaptionImage) SetReadOnlyPixelBuffer(value objectivec.IObject) {
	objc.Send[struct{}](r.ID, objc.Sel("setReadOnlyPixelBuffer:"), value)
}
// A point that defines the position, in pixels, of the rendered caption image
// relative to the video frame.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVRenderedCaptionImage/position
func (r AVRenderedCaptionImage) Position() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](r.ID, objc.Sel("position"))
	return corefoundation.CGPoint(rv)
}

