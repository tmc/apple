// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVCaptureConnection] class.
var (
	_AVCaptureConnectionClass     AVCaptureConnectionClass
	_AVCaptureConnectionClassOnce sync.Once
)

func getAVCaptureConnectionClass() AVCaptureConnectionClass {
	_AVCaptureConnectionClassOnce.Do(func() {
		_AVCaptureConnectionClass = AVCaptureConnectionClass{class: objc.GetClass("AVCaptureConnection")}
	})
	return _AVCaptureConnectionClass
}

// GetAVCaptureConnectionClass returns the class object for AVCaptureConnection.
func GetAVCaptureConnectionClass() AVCaptureConnectionClass {
	return getAVCaptureConnectionClass()
}

type AVCaptureConnectionClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVCaptureConnectionClass) Alloc() AVCaptureConnection {
	rv := objc.Send[AVCaptureConnection](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that represents a connection from a capture input to a capture
// output.
//
// # Overview
// 
// Capture inputs have one or more input ports (instances of
// [AVCaptureInputPort]). Capture outputs can accept data from one or more
// sources (for example, an [AVCaptureMovieFileOutput] object accepts both
// video and audio data).
// 
// You can add an [AVCaptureConnection] instance to a session using the
// [AddConnection] method only if the [CanAddConnection] method returns
// [true]. When using the [AddInput] or [AddOutput] method, the session forms
// connections automatically between all compatible inputs and outputs. You
// only need to add connections manually when adding an input or output with
// no connections. You can also use connections to enable or disable the flow
// of data from a given input or to a given output.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Creating a connection
//
//   - [AVCaptureConnection.InitWithInputPortsOutput]: Creates a capture connection that represents a connection between multiple input ports and an output.
//   - [AVCaptureConnection.InitWithInputPortVideoPreviewLayer]: Creates a capture connection that represents a connection between an input port and a video preview layer.
//
// # Enabling a connection
//
//   - [AVCaptureConnection.Enabled]: Turns the connection on and off.
//   - [AVCaptureConnection.SetEnabled]
//   - [AVCaptureConnection.Active]: Indicates whether the connection is active.
//
// # Inspecting a connection
//
//   - [AVCaptureConnection.InputPorts]: An array of the connection’s input ports.
//   - [AVCaptureConnection.Output]: The connection’s output port, if applicable.
//   - [AVCaptureConnection.VideoPreviewLayer]: The video preview layer associated with the connection.
//   - [AVCaptureConnection.AudioChannels]: An array of audio channels that the connection provides.
//
// # Rotating a video
//
//   - [AVCaptureConnection.IsVideoRotationAngleSupported]: Returns a Boolean value that indicates whether the connection supports a rotation angle.
//   - [AVCaptureConnection.VideoRotationAngle]: A rotation angle the connection applies to a video flowing through it.
//   - [AVCaptureConnection.SetVideoRotationAngle]
//
// # Mirroring a video
//
//   - [AVCaptureConnection.SupportsVideoMirroring]: A Boolean value that indicates whether the connection supports video mirroring.
//   - [AVCaptureConnection.VideoMirrored]: A Boolean value that indicates whether the connection horizontally flips the video flowing through it.
//   - [AVCaptureConnection.SetVideoMirrored]
//   - [AVCaptureConnection.AutomaticallyAdjustsVideoMirroring]: A Boolean value that indicates whether you can enable mirroring based on a session’s configuration.
//   - [AVCaptureConnection.SetAutomaticallyAdjustsVideoMirroring]
//
// # Configuring a video’s frame rate
//
//   - [AVCaptureConnection.SupportsVideoMinFrameDuration]: A Boolean value that indicates whether the connection supports a minimum frame duration.
//   - [AVCaptureConnection.VideoMinFrameDuration]: The smallest time interval the connection can apply between consecutive video frames.
//   - [AVCaptureConnection.SetVideoMinFrameDuration]
//   - [AVCaptureConnection.SupportsVideoMaxFrameDuration]: A Boolean value that indicates whether the connection supports a maximum frame duration.
//   - [AVCaptureConnection.VideoMaxFrameDuration]: The largest time interval the connection can apply between consecutive video frames.
//   - [AVCaptureConnection.SetVideoMaxFrameDuration]
//
// # Interlacing video
//
//   - [AVCaptureConnection.SupportsVideoFieldMode]: A Boolean value that indicates whether the connection supports setting a video field mode.
//   - [AVCaptureConnection.VideoFieldMode]: A setting that tells the connection how to interlace video flowing through it.
//   - [AVCaptureConnection.SetVideoFieldMode]
//
// # Deprecated
//
//   - [AVCaptureConnection.SupportsVideoOrientation]: A Boolean value that indicates whether the connection supports changing the orientation of the video.
//   - [AVCaptureConnection.VideoOrientation]: An orientation that tells the connection how to rotate a video flowing through it.
//   - [AVCaptureConnection.SetVideoOrientation]
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureConnection
type AVCaptureConnection struct {
	objectivec.Object
}

// AVCaptureConnectionFromID constructs a [AVCaptureConnection] from an objc.ID.
//
// An object that represents a connection from a capture input to a capture
// output.
func AVCaptureConnectionFromID(id objc.ID) AVCaptureConnection {
	return AVCaptureConnection{objectivec.Object{ID: id}}
}
// NOTE: AVCaptureConnection adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVCaptureConnection] class.
//
// # Creating a connection
//
//   - [IAVCaptureConnection.InitWithInputPortsOutput]: Creates a capture connection that represents a connection between multiple input ports and an output.
//   - [IAVCaptureConnection.InitWithInputPortVideoPreviewLayer]: Creates a capture connection that represents a connection between an input port and a video preview layer.
//
// # Enabling a connection
//
//   - [IAVCaptureConnection.Enabled]: Turns the connection on and off.
//   - [IAVCaptureConnection.SetEnabled]
//   - [IAVCaptureConnection.Active]: Indicates whether the connection is active.
//
// # Inspecting a connection
//
//   - [IAVCaptureConnection.InputPorts]: An array of the connection’s input ports.
//   - [IAVCaptureConnection.Output]: The connection’s output port, if applicable.
//   - [IAVCaptureConnection.VideoPreviewLayer]: The video preview layer associated with the connection.
//   - [IAVCaptureConnection.AudioChannels]: An array of audio channels that the connection provides.
//
// # Rotating a video
//
//   - [IAVCaptureConnection.IsVideoRotationAngleSupported]: Returns a Boolean value that indicates whether the connection supports a rotation angle.
//   - [IAVCaptureConnection.VideoRotationAngle]: A rotation angle the connection applies to a video flowing through it.
//   - [IAVCaptureConnection.SetVideoRotationAngle]
//
// # Mirroring a video
//
//   - [IAVCaptureConnection.SupportsVideoMirroring]: A Boolean value that indicates whether the connection supports video mirroring.
//   - [IAVCaptureConnection.VideoMirrored]: A Boolean value that indicates whether the connection horizontally flips the video flowing through it.
//   - [IAVCaptureConnection.SetVideoMirrored]
//   - [IAVCaptureConnection.AutomaticallyAdjustsVideoMirroring]: A Boolean value that indicates whether you can enable mirroring based on a session’s configuration.
//   - [IAVCaptureConnection.SetAutomaticallyAdjustsVideoMirroring]
//
// # Configuring a video’s frame rate
//
//   - [IAVCaptureConnection.SupportsVideoMinFrameDuration]: A Boolean value that indicates whether the connection supports a minimum frame duration.
//   - [IAVCaptureConnection.VideoMinFrameDuration]: The smallest time interval the connection can apply between consecutive video frames.
//   - [IAVCaptureConnection.SetVideoMinFrameDuration]
//   - [IAVCaptureConnection.SupportsVideoMaxFrameDuration]: A Boolean value that indicates whether the connection supports a maximum frame duration.
//   - [IAVCaptureConnection.VideoMaxFrameDuration]: The largest time interval the connection can apply between consecutive video frames.
//   - [IAVCaptureConnection.SetVideoMaxFrameDuration]
//
// # Interlacing video
//
//   - [IAVCaptureConnection.SupportsVideoFieldMode]: A Boolean value that indicates whether the connection supports setting a video field mode.
//   - [IAVCaptureConnection.VideoFieldMode]: A setting that tells the connection how to interlace video flowing through it.
//   - [IAVCaptureConnection.SetVideoFieldMode]
//
// # Deprecated
//
//   - [IAVCaptureConnection.SupportsVideoOrientation]: A Boolean value that indicates whether the connection supports changing the orientation of the video.
//   - [IAVCaptureConnection.VideoOrientation]: An orientation that tells the connection how to rotate a video flowing through it.
//   - [IAVCaptureConnection.SetVideoOrientation]
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureConnection
type IAVCaptureConnection interface {
	objectivec.IObject

	// Topic: Creating a connection

	// Creates a capture connection that represents a connection between multiple input ports and an output.
	InitWithInputPortsOutput(ports []AVCaptureInputPort, output IAVCaptureOutput) AVCaptureConnection
	// Creates a capture connection that represents a connection between an input port and a video preview layer.
	InitWithInputPortVideoPreviewLayer(port IAVCaptureInputPort, layer IAVCaptureVideoPreviewLayer) AVCaptureConnection

	// Topic: Enabling a connection

	// Turns the connection on and off.
	Enabled() bool
	SetEnabled(value bool)
	// Indicates whether the connection is active.
	Active() bool

	// Topic: Inspecting a connection

	// An array of the connection’s input ports.
	InputPorts() []AVCaptureInputPort
	// The connection’s output port, if applicable.
	Output() IAVCaptureOutput
	// The video preview layer associated with the connection.
	VideoPreviewLayer() IAVCaptureVideoPreviewLayer
	// An array of audio channels that the connection provides.
	AudioChannels() []AVCaptureAudioChannel

	// Topic: Rotating a video

	// Returns a Boolean value that indicates whether the connection supports a rotation angle.
	IsVideoRotationAngleSupported(videoRotationAngle float64) bool
	// A rotation angle the connection applies to a video flowing through it.
	VideoRotationAngle() float64
	SetVideoRotationAngle(value float64)

	// Topic: Mirroring a video

	// A Boolean value that indicates whether the connection supports video mirroring.
	SupportsVideoMirroring() bool
	// A Boolean value that indicates whether the connection horizontally flips the video flowing through it.
	VideoMirrored() bool
	SetVideoMirrored(value bool)
	// A Boolean value that indicates whether you can enable mirroring based on a session’s configuration.
	AutomaticallyAdjustsVideoMirroring() bool
	SetAutomaticallyAdjustsVideoMirroring(value bool)

	// Topic: Configuring a video’s frame rate

	// A Boolean value that indicates whether the connection supports a minimum frame duration.
	SupportsVideoMinFrameDuration() bool
	// The smallest time interval the connection can apply between consecutive video frames.
	VideoMinFrameDuration() objectivec.IObject
	SetVideoMinFrameDuration(value objectivec.IObject)
	// A Boolean value that indicates whether the connection supports a maximum frame duration.
	SupportsVideoMaxFrameDuration() bool
	// The largest time interval the connection can apply between consecutive video frames.
	VideoMaxFrameDuration() objectivec.IObject
	SetVideoMaxFrameDuration(value objectivec.IObject)

	// Topic: Interlacing video

	// A Boolean value that indicates whether the connection supports setting a video field mode.
	SupportsVideoFieldMode() bool
	// A setting that tells the connection how to interlace video flowing through it.
	VideoFieldMode() AVVideoFieldMode
	SetVideoFieldMode(value AVVideoFieldMode)

	// Topic: Deprecated

	// A Boolean value that indicates whether the connection supports changing the orientation of the video.
	SupportsVideoOrientation() bool
	// An orientation that tells the connection how to rotate a video flowing through it.
	VideoOrientation() AVCaptureVideoOrientation
	SetVideoOrientation(value AVCaptureVideoOrientation)
}

// Init initializes the instance.
func (c AVCaptureConnection) Init() AVCaptureConnection {
	rv := objc.Send[AVCaptureConnection](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c AVCaptureConnection) Autorelease() AVCaptureConnection {
	rv := objc.Send[AVCaptureConnection](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVCaptureConnection creates a new AVCaptureConnection instance.
func NewAVCaptureConnection() AVCaptureConnection {
	class := getAVCaptureConnectionClass()
	rv := objc.Send[AVCaptureConnection](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a capture connection that represents a connection between an input
// port and a video preview layer.
//
// port: An [AVCaptureInputPort] instance that relates to an [AVCaptureInput]
// instance.
//
// layer: An [AVCaptureVideoPreviewLayer] instance.
//
// # Return Value
// 
// A capture connection that represents a connection between `port` and
// `layer`.
//
// # Discussion
// 
// You can add the connection this method returns to an [AVCaptureSession]
// instance with the [AddConnection] method.
// 
// The [AddInput]: or [AddOutput] methods automatically form connections
// between all compatible inputs and outputs. You don’t need to manually
// create and add connections to the session unless you use the primitive
// [AddInputWithNoConnections] and [AddOutputWithNoConnections] methods.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureConnection/init(inputPort:videoPreviewLayer:)
func NewCaptureConnectionWithInputPortVideoPreviewLayer(port IAVCaptureInputPort, layer IAVCaptureVideoPreviewLayer) AVCaptureConnection {
	instance := getAVCaptureConnectionClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithInputPort:videoPreviewLayer:"), port, layer)
	return AVCaptureConnectionFromID(rv)
}

// Creates a capture connection that represents a connection between multiple
// input ports and an output.
//
// ports: An array of [AVCaptureInputPort] instances that relate to [AVCaptureInput]
// instances.
//
// output: An [AVCaptureOutput] instance.
//
// # Return Value
// 
// A capture connection that represents a connection between `ports` and
// `output`.
//
// # Discussion
// 
// You can add the connection this method returns to an [AVCaptureSession]
// instance with the [AddConnection] method.
// 
// The [AddInput]: or [AddOutput] methods automatically form connections
// between all compatible inputs and outputs. You don’t need to manually
// create and add connections to the session unless you use the primitive
// [AddInputWithNoConnections] and [AddOutputWithNoConnections] methods.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureConnection/init(inputPorts:output:)
func NewCaptureConnectionWithInputPortsOutput(ports []AVCaptureInputPort, output IAVCaptureOutput) AVCaptureConnection {
	instance := getAVCaptureConnectionClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithInputPorts:output:"), objectivec.IObjectSliceToNSArray(ports), output)
	return AVCaptureConnectionFromID(rv)
}

// Creates a capture connection that represents a connection between multiple
// input ports and an output.
//
// ports: An array of [AVCaptureInputPort] instances that relate to [AVCaptureInput]
// instances.
//
// output: An [AVCaptureOutput] instance.
//
// # Return Value
// 
// A capture connection that represents a connection between `ports` and
// `output`.
//
// # Discussion
// 
// You can add the connection this method returns to an [AVCaptureSession]
// instance with the [AddConnection] method.
// 
// The [AddInput]: or [AddOutput] methods automatically form connections
// between all compatible inputs and outputs. You don’t need to manually
// create and add connections to the session unless you use the primitive
// [AddInputWithNoConnections] and [AddOutputWithNoConnections] methods.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureConnection/init(inputPorts:output:)
func (c AVCaptureConnection) InitWithInputPortsOutput(ports []AVCaptureInputPort, output IAVCaptureOutput) AVCaptureConnection {
	rv := objc.Send[AVCaptureConnection](c.ID, objc.Sel("initWithInputPorts:output:"), objectivec.IObjectSliceToNSArray(ports), output)
	return rv
}
// Creates a capture connection that represents a connection between an input
// port and a video preview layer.
//
// port: An [AVCaptureInputPort] instance that relates to an [AVCaptureInput]
// instance.
//
// layer: An [AVCaptureVideoPreviewLayer] instance.
//
// # Return Value
// 
// A capture connection that represents a connection between `port` and
// `layer`.
//
// # Discussion
// 
// You can add the connection this method returns to an [AVCaptureSession]
// instance with the [AddConnection] method.
// 
// The [AddInput]: or [AddOutput] methods automatically form connections
// between all compatible inputs and outputs. You don’t need to manually
// create and add connections to the session unless you use the primitive
// [AddInputWithNoConnections] and [AddOutputWithNoConnections] methods.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureConnection/init(inputPort:videoPreviewLayer:)
func (c AVCaptureConnection) InitWithInputPortVideoPreviewLayer(port IAVCaptureInputPort, layer IAVCaptureVideoPreviewLayer) AVCaptureConnection {
	rv := objc.Send[AVCaptureConnection](c.ID, objc.Sel("initWithInputPort:videoPreviewLayer:"), port, layer)
	return rv
}
// Returns a Boolean value that indicates whether the connection supports a
// rotation angle.
//
// videoRotationAngle: A rotation angle in degrees.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureConnection/isVideoRotationAngleSupported(_:)
func (c AVCaptureConnection) IsVideoRotationAngleSupported(videoRotationAngle float64) bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isVideoRotationAngleSupported:"), videoRotationAngle)
	return rv
}

// Returns a capture connection that represents a connection between an input
// port and a video preview layer.
//
// port: An [AVCaptureInputPort] instance that relates to an [AVCaptureInput]
// instance.
//
// layer: An [AVCaptureVideoPreviewLayer] instance.
//
// # Return Value
// 
// A capture connection that represents a connection between `port` and
// `layer`.
//
// # Discussion
// 
// You can add the connection this method returns to an [AVCaptureSession]
// instance with the [AddConnection] method.
// 
// The [AddInput]: or [AddOutput] methods automatically form connections
// between all compatible inputs and outputs. You don’t need to manually
// create and add connections to the session unless you use the primitive
// [AddInputWithNoConnections] and [AddOutputWithNoConnections] methods.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureConnection/connectionWithInputPort:videoPreviewLayer:
func (_AVCaptureConnectionClass AVCaptureConnectionClass) ConnectionWithInputPortVideoPreviewLayer(port IAVCaptureInputPort, layer IAVCaptureVideoPreviewLayer) AVCaptureConnection {
	rv := objc.Send[objc.ID](objc.ID(_AVCaptureConnectionClass.class), objc.Sel("connectionWithInputPort:videoPreviewLayer:"), port, layer)
	return AVCaptureConnectionFromID(rv)
}
// Returns a capture connection that represents a connection between multiple
// input ports and an output.
//
// ports: An array of [AVCaptureInputPort] instances that relate to [AVCaptureInput]
// instances.
//
// output: An [AVCaptureOutput] instance.
//
// # Return Value
// 
// A capture connection that represents a connection between `ports` and
// `output`.
//
// # Discussion
// 
// You can add the connection this method returns to an [AVCaptureSession]
// instance with the [AddConnection] method.
// 
// The [AddInput]: or [AddOutput] methods automatically form connections
// between all compatible inputs and outputs. You don’t need to manually
// create and add connections to the session unless you use the primitive
// [AddInputWithNoConnections] and [AddOutputWithNoConnections] methods.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureConnection/connectionWithInputPorts:output:
func (_AVCaptureConnectionClass AVCaptureConnectionClass) ConnectionWithInputPortsOutput(ports []AVCaptureInputPort, output IAVCaptureOutput) AVCaptureConnection {
	rv := objc.Send[objc.ID](objc.ID(_AVCaptureConnectionClass.class), objc.Sel("connectionWithInputPorts:output:"), objectivec.IObjectSliceToNSArray(ports), output)
	return AVCaptureConnectionFromID(rv)
}

// Turns the connection on and off.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureConnection/isEnabled
func (c AVCaptureConnection) Enabled() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isEnabled"))
	return rv
}
func (c AVCaptureConnection) SetEnabled(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setEnabled:"), value)
}
// Indicates whether the connection is active.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureConnection/isActive
func (c AVCaptureConnection) Active() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isActive"))
	return rv
}
// An array of the connection’s input ports.
//
// # Discussion
// 
// Input ports are instances of [AVCaptureInputPort].
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureConnection/inputPorts
func (c AVCaptureConnection) InputPorts() []AVCaptureInputPort {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("inputPorts"))
	return objc.ConvertSlice(rv, func(id objc.ID) AVCaptureInputPort {
		return AVCaptureInputPortFromID(id)
	})
}
// The connection’s output port, if applicable.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureConnection/output
func (c AVCaptureConnection) Output() IAVCaptureOutput {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("output"))
	return AVCaptureOutputFromID(objc.ID(rv))
}
// The video preview layer associated with the connection.
//
// # Discussion
// 
// The connection sets the property in its
// [InitWithInputPortVideoPreviewLayer] initializer.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureConnection/videoPreviewLayer
func (c AVCaptureConnection) VideoPreviewLayer() IAVCaptureVideoPreviewLayer {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("videoPreviewLayer"))
	return AVCaptureVideoPreviewLayerFromID(objc.ID(rv))
}
// An array of audio channels that the connection provides.
//
// # Discussion
// 
// The property only applies to a video connection.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureConnection/audioChannels
func (c AVCaptureConnection) AudioChannels() []AVCaptureAudioChannel {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("audioChannels"))
	return objc.ConvertSlice(rv, func(id objc.ID) AVCaptureAudioChannel {
		return AVCaptureAudioChannelFromID(id)
	})
}
// A rotation angle the connection applies to a video flowing through it.
//
// # Discussion
// 
// Your app can set a video rotation angle that it gets from an
// [AVCaptureDeviceRotationCoordinator] instance’s
// [VideoRotationAngleForHorizonLevelCapture] or
// [VideoRotationAngleForHorizonLevelPreview] property. The rotation angle
// only applies to video or depth connections, similar to [VideoMirrored], and
// can be any angle that [IsVideoRotationAngleSupported] returns [true] for.
// 
// Not all capture connections rotate each frame. For example, a video
// connection to an [AVCaptureMovieFileOutput] or [AVCapturePhotoOutput]
// instance applies a rotation with a QuickTime track matrix or with EXIF
// tags, respectively.
// 
// Capture connections to [AVCaptureVideoDataOutput] and
// [AVCaptureDepthDataOutput] instances rotate video frames they provide to
// their [CaptureOutputDidOutputSampleBufferFromConnection] and
// [depthDataOutput(_:didOutput:timestamp:connection:)] delegate methods,
// respectively. Each [AVCaptureVideoDataOutput] instance uses hardware
// acceleration to rotate every frame.
// 
// You can rotate the video of a movie file you record with an [AVAssetWriter]
// instance by applying the rotation to an [AVAssetWriterInput] instance’s
// [Transform] property. This approach avoids the performance costs that come
// with rotating each video frame.
//
// [AVCaptureDepthDataOutput]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDepthDataOutput
// [depthDataOutput(_:didOutput:timestamp:connection:)]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDepthDataOutputDelegate/depthDataOutput(_:didOutput:timestamp:connection:)
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureConnection/videoRotationAngle
func (c AVCaptureConnection) VideoRotationAngle() float64 {
	rv := objc.Send[float64](c.ID, objc.Sel("videoRotationAngle"))
	return rv
}
func (c AVCaptureConnection) SetVideoRotationAngle(value float64) {
	objc.Send[struct{}](c.ID, objc.Sel("setVideoRotationAngle:"), value)
}
// A Boolean value that indicates whether the connection supports video
// mirroring.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureConnection/isVideoMirroringSupported
func (c AVCaptureConnection) SupportsVideoMirroring() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isVideoMirroringSupported"))
	return rv
}
// A Boolean value that indicates whether the connection horizontally flips
// the video flowing through it.
//
// # Discussion
// 
// You can apply a mirror-image effect to a video flowing through the
// connection by setting the value to [true]. The mirroring effect only
// applies to video or depth connections, similar to [VideoRotationAngle], and
// if [SupportsVideoMirroring] is [true].
// 
// Not all capture connections mirror each frame. For example, a video
// connection to an [AVCaptureMovieFileOutput] or [AVCapturePhotoOutput]
// instance applies the mirror effect with a QuickTime track matrix or with
// EXIF tags, respectively.
// 
// Capture connections to [AVCaptureVideoDataOutput] and
// [AVCaptureDepthDataOutput] instances mirror video frames they provide to
// their [CaptureOutputDidOutputSampleBufferFromConnection] and
// [depthDataOutput(_:didOutput:timestamp:connection:)] delegate methods,
// respectively. Each [AVCaptureVideoDataOutput] instance uses hardware
// acceleration to mirror every frame.
// 
// You can mirror the video of a movie file you record with an [AVAssetWriter]
// instance by applying a scale factor to the [Transform] property of its
// [AVAssetWriterInput]. For example, you can horizontally flip an image by
// scaling the x-axis by `-1`. This approach avoids the performance costs that
// come with rotating each video frame.
//
// [AVCaptureDepthDataOutput]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDepthDataOutput
// [depthDataOutput(_:didOutput:timestamp:connection:)]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDepthDataOutputDelegate/depthDataOutput(_:didOutput:timestamp:connection:)
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureConnection/isVideoMirrored
func (c AVCaptureConnection) VideoMirrored() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isVideoMirrored"))
	return rv
}
func (c AVCaptureConnection) SetVideoMirrored(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setVideoMirrored:"), value)
}
// A Boolean value that indicates whether you can enable mirroring based on a
// session’s configuration.
//
// # Discussion
// 
// For some session configurations, the connection mirrors the video data by
// default. When the value of this property is [true], the value of
// [VideoMirrored] may change, depending on the configuration of the session.
// For example, the value may change after switching to a different capture
// device input.
// 
// The default value is [true].
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureConnection/automaticallyAdjustsVideoMirroring
func (c AVCaptureConnection) AutomaticallyAdjustsVideoMirroring() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("automaticallyAdjustsVideoMirroring"))
	return rv
}
func (c AVCaptureConnection) SetAutomaticallyAdjustsVideoMirroring(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setAutomaticallyAdjustsVideoMirroring:"), value)
}
// A Boolean value that indicates whether the connection supports a minimum
// frame duration.
//
// # Discussion
// 
// The property indicates whether the connection honors the
// [VideoMinFrameDuration] property for a video connection.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureConnection/isVideoMinFrameDurationSupported
func (c AVCaptureConnection) SupportsVideoMinFrameDuration() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isVideoMinFrameDurationSupported"))
	return rv
}
// The smallest time interval the connection can apply between consecutive
// video frames.
//
// # Discussion
// 
// When [SupportsVideoMinFrameDuration] is [true], the value of the property
// configures the lower bound for the amount of time a video connection
// separates consecutive frames. The value is equivalent to the reciprocal of
// the minimum frame rate.
// 
// You can set an unlimited frame rate with [zero] or [invalid] (which is the
// default).
//
// [invalid]: https://developer.apple.com/documentation/CoreMedia/CMTime/invalid
// [true]: https://developer.apple.com/documentation/Swift/true
// [zero]: https://developer.apple.com/documentation/CoreMedia/CMTime/zero
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureConnection/videoMinFrameDuration
func (c AVCaptureConnection) VideoMinFrameDuration() objectivec.IObject {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("videoMinFrameDuration"))
	return objectivec.Object{ID: rv}
}
func (c AVCaptureConnection) SetVideoMinFrameDuration(value objectivec.IObject) {
	objc.Send[struct{}](c.ID, objc.Sel("setVideoMinFrameDuration:"), value)
}
// A Boolean value that indicates whether the connection supports a maximum
// frame duration.
//
// # Discussion
// 
// The property indicates whether the connection honors the
// [VideoMaxFrameDuration] property for a video connection.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureConnection/isVideoMaxFrameDurationSupported
func (c AVCaptureConnection) SupportsVideoMaxFrameDuration() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isVideoMaxFrameDurationSupported"))
	return rv
}
// The largest time interval the connection can apply between consecutive
// video frames.
//
// # Discussion
// 
// When [SupportsVideoMaxFrameDuration] is [true], the value of the property
// configures the upper bound for the amount of time a video connection
// separates consecutive frames. The value is equivalent to the reciprocal of
// the minimum frame rate.
// 
// You can set an unlimited frame rate with [zero] or [invalid] (which is the
// default).
//
// [invalid]: https://developer.apple.com/documentation/CoreMedia/CMTime/invalid
// [true]: https://developer.apple.com/documentation/Swift/true
// [zero]: https://developer.apple.com/documentation/CoreMedia/CMTime/zero
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureConnection/videoMaxFrameDuration
func (c AVCaptureConnection) VideoMaxFrameDuration() objectivec.IObject {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("videoMaxFrameDuration"))
	return objectivec.Object{ID: rv}
}
func (c AVCaptureConnection) SetVideoMaxFrameDuration(value objectivec.IObject) {
	objc.Send[struct{}](c.ID, objc.Sel("setVideoMaxFrameDuration:"), value)
}
// A Boolean value that indicates whether the connection supports setting a
// video field mode.
//
// # Discussion
// 
// The property only applies to a video connection’s [VideoFieldMode]
// property.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureConnection/isVideoFieldModeSupported
func (c AVCaptureConnection) SupportsVideoFieldMode() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isVideoFieldModeSupported"))
	return rv
}
// A setting that tells the connection how to interlace video flowing through
// it.
//
// # Discussion
// 
// The property only applies to a video connection and when
// [SupportsVideoFieldMode] is [true].
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureConnection/videoFieldMode
func (c AVCaptureConnection) VideoFieldMode() AVVideoFieldMode {
	rv := objc.Send[AVVideoFieldMode](c.ID, objc.Sel("videoFieldMode"))
	return AVVideoFieldMode(rv)
}
func (c AVCaptureConnection) SetVideoFieldMode(value AVVideoFieldMode) {
	objc.Send[struct{}](c.ID, objc.Sel("setVideoFieldMode:"), value)
}
// A Boolean value that indicates whether the connection supports changing the
// orientation of the video.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureConnection/isVideoOrientationSupported
func (c AVCaptureConnection) SupportsVideoOrientation() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isVideoOrientationSupported"))
	return rv
}
// An orientation that tells the connection how to rotate a video flowing
// through it.
//
// # Discussion
// 
// The property only applies to a video connection.
// 
// If the value of [SupportsVideoOrientation] is [true], you can set
// `videoOrientation` to rotate the video buffers consumed by the
// connection’s output. Setting `videoOrientation` doesn’t necessarily
// result in a physical rotation of video buffers. For example, a video
// connection to an [AVCaptureMovieFileOutput] object handles orientation
// using a QuickTime track matrix. A video connection to an
// [AVCaptureStillImageOutput] object handles orientation using EXIF tags.
// 
// [AVCaptureVideoDataOutput] clients may receive physically rotated pixel
// buffers in their [CaptureOutputDidOutputSampleBufferFromConnection]
// delegate callback. The [AVCaptureVideoDataOutput] hardware accelerates the
// rotation operation and supports all four [AVCaptureVideoOrientation] modes.
// A client sets `videoOrientation` or [VideoMirrored] on the video data
// output’s video [AVCaptureConnection] to request physical buffer rotation.
//
// [AVCaptureStillImageOutput]: https://developer.apple.com/documentation/AVFoundation/AVCaptureStillImageOutput
// [AVCaptureVideoOrientation]: https://developer.apple.com/documentation/AVFoundation/AVCaptureVideoOrientation
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureConnection/videoOrientation
func (c AVCaptureConnection) VideoOrientation() AVCaptureVideoOrientation {
	rv := objc.Send[AVCaptureVideoOrientation](c.ID, objc.Sel("videoOrientation"))
	return AVCaptureVideoOrientation(rv)
}
func (c AVCaptureConnection) SetVideoOrientation(value AVCaptureVideoOrientation) {
	objc.Send[struct{}](c.ID, objc.Sel("setVideoOrientation:"), value)
}

