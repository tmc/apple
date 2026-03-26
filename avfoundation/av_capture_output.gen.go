// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVCaptureOutput] class.
var (
	_AVCaptureOutputClass     AVCaptureOutputClass
	_AVCaptureOutputClassOnce sync.Once
)

func getAVCaptureOutputClass() AVCaptureOutputClass {
	_AVCaptureOutputClassOnce.Do(func() {
		_AVCaptureOutputClass = AVCaptureOutputClass{class: objc.GetClass("AVCaptureOutput")}
	})
	return _AVCaptureOutputClass
}

// GetAVCaptureOutputClass returns the class object for AVCaptureOutput.
func GetAVCaptureOutputClass() AVCaptureOutputClass {
	return getAVCaptureOutputClass()
}

type AVCaptureOutputClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVCaptureOutputClass) Alloc() AVCaptureOutput {
	rv := objc.Send[AVCaptureOutput](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An abstract superclass for objects that provide media output destinations
// for a capture session.
//
// # Overview
// 
// This class provides an abstract interface to connect capture output
// destinations, such as files and streams, to a capture session.
// 
// A capture output can have multiple connections, one for each stream of
// media that it receives from a capture input. A capture output doesn’t
// have any connections when you create it. When you add it to a capture
// session, the session automatically forms connections between compatible
// inputs and outputs.
//
// # Accessing connections
//
//   - [AVCaptureOutput.Connections]: The capture output object’s connections.
//   - [AVCaptureOutput.ConnectionWithMediaType]: Returns the first connection with an input port of a specified media type.
//
// # Managing deferred start
//
//   - [AVCaptureOutput.DeferredStartEnabled]: A Boolean value that indicates whether to defer starting this capture output.
//   - [AVCaptureOutput.SetDeferredStartEnabled]
//   - [AVCaptureOutput.DeferredStartSupported]: A [BOOL] value that indicates whether the output supports deferred start.
//
// # Converting between coordinate systems
//
//   - [AVCaptureOutput.TransformedMetadataObjectForMetadataObjectConnection]: Converts a metadata object’s visual properties to layer coordinates.
//   - [AVCaptureOutput.MetadataOutputRectOfInterestForRect]: Converts a rectangle in the capture output object’s coordinate system to one in the coordinate system used for metadata outputs.
//   - [AVCaptureOutput.RectForMetadataOutputRectOfInterest]: Converts a rectangle in the coordinate system used for metadata outputs to one in the capture output object’s coordinate system.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureOutput
type AVCaptureOutput struct {
	objectivec.Object
}

// AVCaptureOutputFromID constructs a [AVCaptureOutput] from an objc.ID.
//
// An abstract superclass for objects that provide media output destinations
// for a capture session.
func AVCaptureOutputFromID(id objc.ID) AVCaptureOutput {
	return AVCaptureOutput{objectivec.Object{ID: id}}
}
// NOTE: AVCaptureOutput adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVCaptureOutput] class.
//
// # Accessing connections
//
//   - [IAVCaptureOutput.Connections]: The capture output object’s connections.
//   - [IAVCaptureOutput.ConnectionWithMediaType]: Returns the first connection with an input port of a specified media type.
//
// # Managing deferred start
//
//   - [IAVCaptureOutput.DeferredStartEnabled]: A Boolean value that indicates whether to defer starting this capture output.
//   - [IAVCaptureOutput.SetDeferredStartEnabled]
//   - [IAVCaptureOutput.DeferredStartSupported]: A [BOOL] value that indicates whether the output supports deferred start.
//
// # Converting between coordinate systems
//
//   - [IAVCaptureOutput.TransformedMetadataObjectForMetadataObjectConnection]: Converts a metadata object’s visual properties to layer coordinates.
//   - [IAVCaptureOutput.MetadataOutputRectOfInterestForRect]: Converts a rectangle in the capture output object’s coordinate system to one in the coordinate system used for metadata outputs.
//   - [IAVCaptureOutput.RectForMetadataOutputRectOfInterest]: Converts a rectangle in the coordinate system used for metadata outputs to one in the capture output object’s coordinate system.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureOutput
type IAVCaptureOutput interface {
	objectivec.IObject

	// Topic: Accessing connections

	// The capture output object’s connections.
	Connections() []AVCaptureConnection
	// Returns the first connection with an input port of a specified media type.
	ConnectionWithMediaType(mediaType AVMediaType) IAVCaptureConnection

	// Topic: Managing deferred start

	// A Boolean value that indicates whether to defer starting this capture output.
	DeferredStartEnabled() bool
	SetDeferredStartEnabled(value bool)
	// A [BOOL] value that indicates whether the output supports deferred start.
	DeferredStartSupported() bool

	// Topic: Converting between coordinate systems

	// Converts a metadata object’s visual properties to layer coordinates.
	TransformedMetadataObjectForMetadataObjectConnection(metadataObject IAVMetadataObject, connection IAVCaptureConnection) IAVMetadataObject
	// Converts a rectangle in the capture output object’s coordinate system to one in the coordinate system used for metadata outputs.
	MetadataOutputRectOfInterestForRect(rectInOutputCoordinates corefoundation.CGRect) corefoundation.CGRect
	// Converts a rectangle in the coordinate system used for metadata outputs to one in the capture output object’s coordinate system.
	RectForMetadataOutputRectOfInterest(rectInMetadataOutputCoordinates corefoundation.CGRect) corefoundation.CGRect
}

// Init initializes the instance.
func (c AVCaptureOutput) Init() AVCaptureOutput {
	rv := objc.Send[AVCaptureOutput](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c AVCaptureOutput) Autorelease() AVCaptureOutput {
	rv := objc.Send[AVCaptureOutput](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVCaptureOutput creates a new AVCaptureOutput instance.
func NewAVCaptureOutput() AVCaptureOutput {
	class := getAVCaptureOutputClass()
	rv := objc.Send[AVCaptureOutput](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Returns the first connection with an input port of a specified media type.
//
// mediaType: A media type such as [video] or [audio].
// //
// [audio]: https://developer.apple.com/documentation/AVFoundation/AVMediaType/audio
// [video]: https://developer.apple.com/documentation/AVFoundation/AVMediaType/video
//
// # Return Value
// 
// The first capture connection that has the specified media type, or `nil` if
// no connection for the media type exists.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureOutput/connection(with:)
func (c AVCaptureOutput) ConnectionWithMediaType(mediaType AVMediaType) IAVCaptureConnection {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("connectionWithMediaType:"), objc.String(string(mediaType)))
	return AVCaptureConnectionFromID(rv)
}
// Converts a metadata object’s visual properties to layer coordinates.
//
// metadataObject: A metadata object that originates from the same capture input as the
// preview layer.
//
// connection: The output’s connection whose input matches that of the metadata object
// to convert.
//
// # Return Value
// 
// A metadata object whose properties are in output coordinates, or `nil` if
// the object originates from an input source other than the preview layer.
//
// # Discussion
// 
// A metadata object has a rectangular bounds where `0,0` is the top-left of
// the picture area, and `1,1` represents the bottom-right on an unrotated
// picture. Face metadata objects likewise express yaw and roll angles with
// respect to an unrotated picture.
// 
// This method converts the visual properties in the coordinate space of the
// supplied metadata object to the coordinate space of the output. The
// conversion takes orientation, mirroring, layer bounds and video gravity
// into consideration. If the provided metadata object originates from an
// input source other than the preview layer’s, the method returns `nil`.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureOutput/transformedMetadataObject(for:connection:)
func (c AVCaptureOutput) TransformedMetadataObjectForMetadataObjectConnection(metadataObject IAVMetadataObject, connection IAVCaptureConnection) IAVMetadataObject {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("transformedMetadataObjectForMetadataObject:connection:"), metadataObject, connection)
	return AVMetadataObjectFromID(rv)
}
// Converts a rectangle in the capture output object’s coordinate system to
// one in the coordinate system used for metadata outputs.
//
// rectInOutputCoordinates: A rectangle in the [AVCaptureOutput] object’s coordinate system.
//
// # Return Value
// 
// A rectangle in the [AVCaptureMetadataOutput] coordinate system.
//
// # Discussion
// 
// An [AVCaptureMetadataOutput] object expresses its [RectOfInterest] as a
// [CGRect] where 0,0 represents the top-left of the picture area, and 1,1
// represents the bottom-right on an unrotated picture. This convenience
// method converts a rectangle in the coordinate space of the output to a
// rectangle of interest in the coordinate space of a metadata output whose
// capture device provides input to the output. The conversion takes
// orientation, mirroring, and scaling into consideration.
// 
// See [TransformedMetadataObjectForMetadataObjectConnection] for a full
// discussion of how the system applies orientation and mirroring to sample
// buffers passing through the output.
//
// [CGRect]: https://developer.apple.com/documentation/CoreFoundation/CGRect
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureOutput/metadataOutputRectConverted(fromOutputRect:)
func (c AVCaptureOutput) MetadataOutputRectOfInterestForRect(rectInOutputCoordinates corefoundation.CGRect) corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](c.ID, objc.Sel("metadataOutputRectOfInterestForRect:"), rectInOutputCoordinates)
	return corefoundation.CGRect(rv)
}
// Converts a rectangle in the coordinate system used for metadata outputs to
// one in the capture output object’s coordinate system.
//
// rectInMetadataOutputCoordinates: A rectangle in the [AVCaptureMetadataOutput] coordinate system.
//
// # Return Value
// 
// A rectangle in the [AVCaptureOutput] object’s coordinate system.
//
// # Discussion
// 
// The rectangle of interest for an [AVCaptureMetadataOutput] object is in a
// coordinate system extending from `{0,0}` in the top-left to `{1,1}` in the
// bottom-right, relative to the device’s natural orientation. A capture
// output object uses a pixel coordinate space which you may zoom, rotate, or
// mirror.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureOutput/outputRectConverted(fromMetadataOutputRect:)
func (c AVCaptureOutput) RectForMetadataOutputRectOfInterest(rectInMetadataOutputCoordinates corefoundation.CGRect) corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](c.ID, objc.Sel("rectForMetadataOutputRectOfInterest:"), rectInMetadataOutputCoordinates)
	return corefoundation.CGRect(rv)
}

// The capture output object’s connections.
//
// # Discussion
// 
// Each connection object in the array describes the mapping between the
// output and the capture input ports.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureOutput/connections
func (c AVCaptureOutput) Connections() []AVCaptureConnection {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("connections"))
	return objc.ConvertSlice(rv, func(id objc.ID) AVCaptureConnection {
		return AVCaptureConnectionFromID(id)
	})
}
// A Boolean value that indicates whether to defer starting this capture
// output.
//
// # Discussion
// 
// When this value is `true`, the session doesn’t prepare the output’s
// resources until some time after [StartRunning] returns. You can start the
// visual parts of your user interface prior to other parts, such as photo or
// movie capture, metadata output, and so on, to improve startup performance.
// Set this value to `false` for outputs that your app needs for startup, and
// `true` for the ones that it doesn’t need to start immediately. For
// example, an [AVCaptureVideoDataOutput] that you intend to use for
// displaying preview should set this value to `false`, so that the frames are
// available as soon as possible.
// 
// For apps that are linked on or after iOS 26, this property value is `true`
// for [AVCapturePhotoOutput] and [AVCaptureFileOutput] subclasses if
// supported, and `false` otherwise. When set to `true` for
// [AVCapturePhotoOutput], if you want to support multiple capture requests
// before running deferred start, set [ResponsiveCaptureEnabled] to `true` on
// that output.
// 
// If [DeferredStartSupported] is `false`, setting this property value to
// `true` results in the system throwing an invalid argument exception.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureOutput/isDeferredStartEnabled
func (c AVCaptureOutput) DeferredStartEnabled() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isDeferredStartEnabled"))
	return rv
}
func (c AVCaptureOutput) SetDeferredStartEnabled(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setDeferredStartEnabled:"), value)
}
// A [BOOL] value that indicates whether the output supports deferred start.
//
// # Discussion
// 
// You can only set the [DeferredStartEnabled] property value to `true` if the
// output supports deferred start.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureOutput/isDeferredStartSupported
func (c AVCaptureOutput) DeferredStartSupported() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isDeferredStartSupported"))
	return rv
}

