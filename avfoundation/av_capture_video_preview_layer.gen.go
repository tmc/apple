// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/quartzcore"
)

// The class instance for the [AVCaptureVideoPreviewLayer] class.
var (
	_AVCaptureVideoPreviewLayerClass     AVCaptureVideoPreviewLayerClass
	_AVCaptureVideoPreviewLayerClassOnce sync.Once
)

func getAVCaptureVideoPreviewLayerClass() AVCaptureVideoPreviewLayerClass {
	_AVCaptureVideoPreviewLayerClassOnce.Do(func() {
		_AVCaptureVideoPreviewLayerClass = AVCaptureVideoPreviewLayerClass{class: objc.GetClass("AVCaptureVideoPreviewLayer")}
	})
	return _AVCaptureVideoPreviewLayerClass
}

// GetAVCaptureVideoPreviewLayerClass returns the class object for AVCaptureVideoPreviewLayer.
func GetAVCaptureVideoPreviewLayerClass() AVCaptureVideoPreviewLayerClass {
	return getAVCaptureVideoPreviewLayerClass()
}

type AVCaptureVideoPreviewLayerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVCaptureVideoPreviewLayerClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVCaptureVideoPreviewLayerClass) Alloc() AVCaptureVideoPreviewLayer {
	rv := objc.Send[AVCaptureVideoPreviewLayer](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// A Core Animation layer that displays video from a camera device.
//
// # Overview
//
// Use this layer to provide a preview of the content the camera captures. A
// convenient way to use this class in iOS is to set it as the backing layer
// for a view as shown below.
//
// # Creating a preview layer
//
//   - [AVCaptureVideoPreviewLayer.InitWithSession]: Creates a layer to preview the visual output of a capture session.
//   - [AVCaptureVideoPreviewLayer.InitWithSessionWithNoConnection]: Creates a layer to preview the visual output of a capture session, without making connections to eligible video inputs.
//
// # Layer configuration
//
//   - [AVCaptureVideoPreviewLayer.VideoGravity]: A value that indicates how the layer displays video content within its bounds.
//   - [AVCaptureVideoPreviewLayer.SetVideoGravity]
//
// # Configuring deferred start
//
//   - [AVCaptureVideoPreviewLayer.DeferredStartSupported]: A [BOOL] value that indicates whether the preview layer supports deferred start.
//   - [AVCaptureVideoPreviewLayer.DeferredStartEnabled]: A [BOOL] value that indicates whether to defer starting this preview layer.
//   - [AVCaptureVideoPreviewLayer.SetDeferredStartEnabled]
//
// # Session configuration
//
//   - [AVCaptureVideoPreviewLayer.Session]: A capture session with visual output to preview.
//   - [AVCaptureVideoPreviewLayer.SetSession]
//   - [AVCaptureVideoPreviewLayer.Connection]: An object that describes the connection from the layer to a particular input port.
//   - [AVCaptureVideoPreviewLayer.SetSessionWithNoConnection]: Associates a session with the layer without automatically forming a connection to an eligible input port.
//
// # Converting between coordinate spaces
//
//   - [AVCaptureVideoPreviewLayer.PointForCaptureDevicePointOfInterest]: Converts a point from the coordinate space of the capture device to the coordinate space of the layer.
//   - [AVCaptureVideoPreviewLayer.CaptureDevicePointOfInterestForPoint]: Converts a point from layer coordinates to the coordinate space of the capture device.
//   - [AVCaptureVideoPreviewLayer.RectForMetadataOutputRectOfInterest]: Converts a rectangle from metadata output coordinates to the coordinate space of the layer.
//   - [AVCaptureVideoPreviewLayer.MetadataOutputRectOfInterestForRect]: Converts a rectangle from layer coordinates to the coordinate space of the metadata output.
//   - [AVCaptureVideoPreviewLayer.TransformedMetadataObjectForMetadataObject]: Converts a metadata object’s visual properties to layer coordinates.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureVideoPreviewLayer
type AVCaptureVideoPreviewLayer struct {
	quartzcore.CALayer
}

// AVCaptureVideoPreviewLayerFromID constructs a [AVCaptureVideoPreviewLayer] from an objc.ID.
//
// A Core Animation layer that displays video from a camera device.
func AVCaptureVideoPreviewLayerFromID(id objc.ID) AVCaptureVideoPreviewLayer {
	return AVCaptureVideoPreviewLayer{CALayer: quartzcore.CALayerFromID(id)}
}

// NOTE: AVCaptureVideoPreviewLayer adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVCaptureVideoPreviewLayer] class.
//
// # Creating a preview layer
//
//   - [IAVCaptureVideoPreviewLayer.InitWithSession]: Creates a layer to preview the visual output of a capture session.
//   - [IAVCaptureVideoPreviewLayer.InitWithSessionWithNoConnection]: Creates a layer to preview the visual output of a capture session, without making connections to eligible video inputs.
//
// # Layer configuration
//
//   - [IAVCaptureVideoPreviewLayer.VideoGravity]: A value that indicates how the layer displays video content within its bounds.
//   - [IAVCaptureVideoPreviewLayer.SetVideoGravity]
//
// # Configuring deferred start
//
//   - [IAVCaptureVideoPreviewLayer.DeferredStartSupported]: A [BOOL] value that indicates whether the preview layer supports deferred start.
//   - [IAVCaptureVideoPreviewLayer.DeferredStartEnabled]: A [BOOL] value that indicates whether to defer starting this preview layer.
//   - [IAVCaptureVideoPreviewLayer.SetDeferredStartEnabled]
//
// # Session configuration
//
//   - [IAVCaptureVideoPreviewLayer.Session]: A capture session with visual output to preview.
//   - [IAVCaptureVideoPreviewLayer.SetSession]
//   - [IAVCaptureVideoPreviewLayer.Connection]: An object that describes the connection from the layer to a particular input port.
//   - [IAVCaptureVideoPreviewLayer.SetSessionWithNoConnection]: Associates a session with the layer without automatically forming a connection to an eligible input port.
//
// # Converting between coordinate spaces
//
//   - [IAVCaptureVideoPreviewLayer.PointForCaptureDevicePointOfInterest]: Converts a point from the coordinate space of the capture device to the coordinate space of the layer.
//   - [IAVCaptureVideoPreviewLayer.CaptureDevicePointOfInterestForPoint]: Converts a point from layer coordinates to the coordinate space of the capture device.
//   - [IAVCaptureVideoPreviewLayer.RectForMetadataOutputRectOfInterest]: Converts a rectangle from metadata output coordinates to the coordinate space of the layer.
//   - [IAVCaptureVideoPreviewLayer.MetadataOutputRectOfInterestForRect]: Converts a rectangle from layer coordinates to the coordinate space of the metadata output.
//   - [IAVCaptureVideoPreviewLayer.TransformedMetadataObjectForMetadataObject]: Converts a metadata object’s visual properties to layer coordinates.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureVideoPreviewLayer
type IAVCaptureVideoPreviewLayer interface {
	quartzcore.ICALayer

	// Topic: Creating a preview layer

	// Creates a layer to preview the visual output of a capture session.
	InitWithSession(session IAVCaptureSession) AVCaptureVideoPreviewLayer
	// Creates a layer to preview the visual output of a capture session, without making connections to eligible video inputs.
	InitWithSessionWithNoConnection(session IAVCaptureSession) AVCaptureVideoPreviewLayer

	// Topic: Layer configuration

	// A value that indicates how the layer displays video content within its bounds.
	VideoGravity() AVLayerVideoGravity
	SetVideoGravity(value AVLayerVideoGravity)

	// Topic: Configuring deferred start

	// A [BOOL] value that indicates whether the preview layer supports deferred start.
	DeferredStartSupported() bool
	// A [BOOL] value that indicates whether to defer starting this preview layer.
	DeferredStartEnabled() bool
	SetDeferredStartEnabled(value bool)

	// Topic: Session configuration

	// A capture session with visual output to preview.
	Session() IAVCaptureSession
	SetSession(value IAVCaptureSession)
	// An object that describes the connection from the layer to a particular input port.
	Connection() IAVCaptureConnection
	// Associates a session with the layer without automatically forming a connection to an eligible input port.
	SetSessionWithNoConnection(session IAVCaptureSession)

	// Topic: Converting between coordinate spaces

	// Converts a point from the coordinate space of the capture device to the coordinate space of the layer.
	PointForCaptureDevicePointOfInterest(captureDevicePointOfInterest corefoundation.CGPoint) corefoundation.CGPoint
	// Converts a point from layer coordinates to the coordinate space of the capture device.
	CaptureDevicePointOfInterestForPoint(pointInLayer corefoundation.CGPoint) corefoundation.CGPoint
	// Converts a rectangle from metadata output coordinates to the coordinate space of the layer.
	RectForMetadataOutputRectOfInterest(rectInMetadataOutputCoordinates corefoundation.CGRect) corefoundation.CGRect
	// Converts a rectangle from layer coordinates to the coordinate space of the metadata output.
	MetadataOutputRectOfInterestForRect(rectInLayerCoordinates corefoundation.CGRect) corefoundation.CGRect
	// Converts a metadata object’s visual properties to layer coordinates.
	TransformedMetadataObjectForMetadataObject(metadataObject IAVMetadataObject) IAVMetadataObject
}

// Init initializes the instance.
func (c AVCaptureVideoPreviewLayer) Init() AVCaptureVideoPreviewLayer {
	rv := objc.Send[AVCaptureVideoPreviewLayer](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c AVCaptureVideoPreviewLayer) Autorelease() AVCaptureVideoPreviewLayer {
	rv := objc.Send[AVCaptureVideoPreviewLayer](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVCaptureVideoPreviewLayer creates a new AVCaptureVideoPreviewLayer instance.
func NewAVCaptureVideoPreviewLayer() AVCaptureVideoPreviewLayer {
	class := getAVCaptureVideoPreviewLayerClass()
	rv := objc.Send[AVCaptureVideoPreviewLayer](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a layer to preview the visual output of a capture session.
//
// session: A capture session to preview.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureVideoPreviewLayer/init(session:)
func NewCaptureVideoPreviewLayerWithSession(session IAVCaptureSession) AVCaptureVideoPreviewLayer {
	instance := getAVCaptureVideoPreviewLayerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSession:"), session)
	return AVCaptureVideoPreviewLayerFromID(rv)
}

// Creates a layer to preview the visual output of a capture session, without
// making connections to eligible video inputs.
//
// session: A capture session to preview.
//
// # Discussion
//
// Only use this initializer if you intend to manually connect the layer to a
// particular [AVCaptureInputPort] by calling the session’s [AddConnection]
// method.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureVideoPreviewLayer/init(sessionWithNoConnection:)
func NewCaptureVideoPreviewLayerWithSessionWithNoConnection(session IAVCaptureSession) AVCaptureVideoPreviewLayer {
	instance := getAVCaptureVideoPreviewLayerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSessionWithNoConnection:"), session)
	return AVCaptureVideoPreviewLayerFromID(rv)
}

// Creates a layer to preview the visual output of a capture session.
//
// session: A capture session to preview.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureVideoPreviewLayer/init(session:)
func (c AVCaptureVideoPreviewLayer) InitWithSession(session IAVCaptureSession) AVCaptureVideoPreviewLayer {
	rv := objc.Send[AVCaptureVideoPreviewLayer](c.ID, objc.Sel("initWithSession:"), session)
	return rv
}

// Creates a layer to preview the visual output of a capture session, without
// making connections to eligible video inputs.
//
// session: A capture session to preview.
//
// # Discussion
//
// Only use this initializer if you intend to manually connect the layer to a
// particular [AVCaptureInputPort] by calling the session’s [AddConnection]
// method.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureVideoPreviewLayer/init(sessionWithNoConnection:)
func (c AVCaptureVideoPreviewLayer) InitWithSessionWithNoConnection(session IAVCaptureSession) AVCaptureVideoPreviewLayer {
	rv := objc.Send[AVCaptureVideoPreviewLayer](c.ID, objc.Sel("initWithSessionWithNoConnection:"), session)
	return rv
}

// Associates a session with the layer without automatically forming a
// connection to an eligible input port.
//
// session: A capture session.
//
// # Discussion
//
// Only use this method if you intend to manually create a connection between
// the layer and a particular [AVCaptureInputPort], and add it to the session
// using its [AddConnection] method.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureVideoPreviewLayer/setSessionWithNoConnection(_:)
func (c AVCaptureVideoPreviewLayer) SetSessionWithNoConnection(session IAVCaptureSession) {
	objc.Send[objc.ID](c.ID, objc.Sel("setSessionWithNoConnection:"), session)
}

// Converts a point from the coordinate space of the capture device to the
// coordinate space of the layer.
//
// captureDevicePointOfInterest: A point in capture device coordinates to convert.
//
// # Return Value
//
// A point in layer coordinates.
//
// # Discussion
//
// A capture device’s [FocusPointOfInterest] and [ExposurePointOfInterest]
// properties provide a [CGPoint] value where `{0,0}` represents the top-left
// and `{1,1}` represents the bottom-right of the unrotated image.
//
// The system takes the layer’s frame size and its [VideoGravity] into
// consideration when making the conversion.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureVideoPreviewLayer/layerPointConverted(fromCaptureDevicePoint:)
//
// [CGPoint]: https://developer.apple.com/documentation/CoreFoundation/CGPoint
func (c AVCaptureVideoPreviewLayer) PointForCaptureDevicePointOfInterest(captureDevicePointOfInterest corefoundation.CGPoint) corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](c.ID, objc.Sel("pointForCaptureDevicePointOfInterest:"), captureDevicePointOfInterest)
	return corefoundation.CGPoint(rv)
}

// Converts a point from layer coordinates to the coordinate space of the
// capture device.
//
// pointInLayer: A point in layer coordinates to convert.
//
// # Return Value
//
// A point in capture device coordinates.
//
// # Discussion
//
// A capture device’s [FocusPointOfInterest] and [ExposurePointOfInterest]
// properties provide a [CGPoint] value where `{0,0}` represents the top-left
// and `{1,1}` represents the bottom-right of the unrotated image.
//
// The conversion performed by this method takes the layer’s frame size and
// its [VideoGravity] into consideration.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureVideoPreviewLayer/captureDevicePointConverted(fromLayerPoint:)
//
// [CGPoint]: https://developer.apple.com/documentation/CoreFoundation/CGPoint
func (c AVCaptureVideoPreviewLayer) CaptureDevicePointOfInterestForPoint(pointInLayer corefoundation.CGPoint) corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](c.ID, objc.Sel("captureDevicePointOfInterestForPoint:"), pointInLayer)
	return corefoundation.CGPoint(rv)
}

// Converts a rectangle from metadata output coordinates to the coordinate
// space of the layer.
//
// rectInMetadataOutputCoordinates: A rectangle in the [AVCaptureMetadataOutput] coordinate system.
//
// # Return Value
//
// A rectangle in the preview layer’s coordinate system.
//
// # Discussion
//
// A metadata capture output’s [RectOfInterest] a [CGRect] value where
// `{0,0}` represents the top-left of the picture area, and `{1,1}` represents
// the bottom-right on an unrotated image.
//
// The system takes the layer’s frame size and its [VideoGravity] into
// consideration when making the conversion.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureVideoPreviewLayer/layerRectConverted(fromMetadataOutputRect:)
//
// [CGRect]: https://developer.apple.com/documentation/CoreFoundation/CGRect
func (c AVCaptureVideoPreviewLayer) RectForMetadataOutputRectOfInterest(rectInMetadataOutputCoordinates corefoundation.CGRect) corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](c.ID, objc.Sel("rectForMetadataOutputRectOfInterest:"), rectInMetadataOutputCoordinates)
	return corefoundation.CGRect(rv)
}

// Converts a rectangle from layer coordinates to the coordinate space of the
// metadata output.
//
// rectInLayerCoordinates: A rectangle in the [AVCaptureVideoPreviewLayer] object’s coordinate
// system.
//
// # Return Value
//
// A rectangle in the metadata output’s coordinate system.
//
// # Discussion
//
// A metadata capture output’s [RectOfInterest] a [CGRect] value where
// `{0,0}` represents the top-left of the picture area, and `{1,1}` represents
// the bottom-right on an unrotated image.
//
// The system takes the layer’s frame size and its [VideoGravity] into
// consideration when making the conversion.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureVideoPreviewLayer/metadataOutputRectConverted(fromLayerRect:)
//
// [CGRect]: https://developer.apple.com/documentation/CoreFoundation/CGRect
func (c AVCaptureVideoPreviewLayer) MetadataOutputRectOfInterestForRect(rectInLayerCoordinates corefoundation.CGRect) corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](c.ID, objc.Sel("metadataOutputRectOfInterestForRect:"), rectInLayerCoordinates)
	return corefoundation.CGRect(rv)
}

// Converts a metadata object’s visual properties to layer coordinates.
//
// metadataObject: The metadata object whose visual properties you want to convert. The
// metadata object must originate from the same [AVCaptureInput] as the
// preview layer.
//
// # Return Value
//
// A metadata object with coordinates converted into layer coordinates, or
// `nil` if the metadata object originates from an input source other than
// that of the preview layer.
//
// # Discussion
//
// The system provides the metadata object’s bounds as a rectangle where
// `{0,0}` represents the top-left of the picture area, and `{1,1}` represents
// the bottom-right on an unrotated image. Face metadata objects also provide
// [YawAngle] and [RollAngle] values with respect to an unrotated picture.
//
// The conversion takes orientation, mirroring, layer bounds and video gravity
// into consideration.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureVideoPreviewLayer/transformedMetadataObject(for:)
func (c AVCaptureVideoPreviewLayer) TransformedMetadataObjectForMetadataObject(metadataObject IAVMetadataObject) IAVMetadataObject {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("transformedMetadataObjectForMetadataObject:"), metadataObject)
	return AVMetadataObjectFromID(rv)
}

// Returns a new layer to preview the visual output of a capture session.
//
// session: The capture session from which to source the preview.
//
// # Return Value
//
// A preview layer.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureVideoPreviewLayer/layerWithSession:
func (_AVCaptureVideoPreviewLayerClass AVCaptureVideoPreviewLayerClass) LayerWithSession(session IAVCaptureSession) AVCaptureVideoPreviewLayer {
	rv := objc.Send[objc.ID](objc.ID(_AVCaptureVideoPreviewLayerClass.class), objc.Sel("layerWithSession:"), session)
	return AVCaptureVideoPreviewLayerFromID(rv)
}

// Returns a new layer to preview the visual output of a capture session,
// without making connections to eligible video inputs.
//
// session: A capture session to preview.
//
// # Return Value
//
// # A preview layer with no connections to a session’s eligible video inputs
//
// # Discussion
//
// Only use this initializer if you intend to manually connect the layer to a
// particular [AVCaptureInputPort] by calling the session’s [AddConnection]
// method.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureVideoPreviewLayer/layerWithSessionWithNoConnection:
func (_AVCaptureVideoPreviewLayerClass AVCaptureVideoPreviewLayerClass) LayerWithSessionWithNoConnection(session IAVCaptureSession) AVCaptureVideoPreviewLayer {
	rv := objc.Send[objc.ID](objc.ID(_AVCaptureVideoPreviewLayerClass.class), objc.Sel("layerWithSessionWithNoConnection:"), session)
	return AVCaptureVideoPreviewLayerFromID(rv)
}

// A value that indicates how the layer displays video content within its
// bounds.
//
// # Discussion
//
// Options are [resizeAspect], [resizeAspectFill], and [resize]. The default
// is [resizeAspect].
//
// This property is animatable.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureVideoPreviewLayer/videoGravity
//
// [resizeAspectFill]: https://developer.apple.com/documentation/AVFoundation/AVLayerVideoGravity/resizeAspectFill
// [resizeAspect]: https://developer.apple.com/documentation/AVFoundation/AVLayerVideoGravity/resizeAspect
// [resize]: https://developer.apple.com/documentation/AVFoundation/AVLayerVideoGravity/resize
func (c AVCaptureVideoPreviewLayer) VideoGravity() AVLayerVideoGravity {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("videoGravity"))
	return AVLayerVideoGravity(foundation.NSStringFromID(rv).String())
}
func (c AVCaptureVideoPreviewLayer) SetVideoGravity(value AVLayerVideoGravity) {
	objc.Send[struct{}](c.ID, objc.Sel("setVideoGravity:"), objc.String(string(value)))
}

// A [BOOL] value that indicates whether the preview layer supports deferred
// start.
//
// # Discussion
//
// You can only set the [DeferredStartEnabled] property to `true` if the
// preview layer supports deferred start.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureVideoPreviewLayer/isDeferredStartSupported
func (c AVCaptureVideoPreviewLayer) DeferredStartSupported() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isDeferredStartSupported"))
	return rv
}

// A [BOOL] value that indicates whether to defer starting this preview layer.
//
// # Discussion
//
// When this value is `true`, the session does not prepare the output’s
// resources until some time after [StartRunning] returns. You can start the
// visual parts of your user interface (e.g. preview) prior to other parts
// (e.g. photo/movie capture, metadata output, etc..) to improve startup
// performance. Set this value to `false` if your app needs video preview
// immediately for startup, and `true` if it does not.
//
// By default, this value is `false` for [AVCaptureVideoPreviewLayer] objects,
// since this object is used to display preview. For best session start
// performance, set [DeferredStartEnabled] to `false` for preview layers. If
// your app contains multiple preview layers, you may want to display the main
// preview layer as soon as possible and allow the remaining layers to display
// subsequently. In this case, set [DeferredStartEnabled] to `true` for the
// remaining layers.
//
// If [DeferredStartSupported] is `false`, setting this property value to
// `true` results in the session throwing an [NSInvalidArgumentException].
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureVideoPreviewLayer/isDeferredStartEnabled
func (c AVCaptureVideoPreviewLayer) DeferredStartEnabled() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isDeferredStartEnabled"))
	return rv
}
func (c AVCaptureVideoPreviewLayer) SetDeferredStartEnabled(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setDeferredStartEnabled:"), value)
}

// A capture session with visual output to preview.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureVideoPreviewLayer/session
func (c AVCaptureVideoPreviewLayer) Session() IAVCaptureSession {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("session"))
	return AVCaptureSessionFromID(objc.ID(rv))
}
func (c AVCaptureVideoPreviewLayer) SetSession(value IAVCaptureSession) {
	objc.Send[struct{}](c.ID, objc.Sel("setSession:"), value)
}

// An object that describes the connection from the layer to a particular
// input port.
//
// # Discussion
//
// When you associate a preview layer with a capture session, the session
// automatically creates a connection to the first eligible video
// [AVCaptureInputPort] object. If you detach a preview layer from a session,
// the connection property becomes `nil`.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureVideoPreviewLayer/connection
func (c AVCaptureVideoPreviewLayer) Connection() IAVCaptureConnection {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("connection"))
	return AVCaptureConnectionFromID(objc.ID(rv))
}
