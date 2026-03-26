// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVCaptureInputPort] class.
var (
	_AVCaptureInputPortClass     AVCaptureInputPortClass
	_AVCaptureInputPortClassOnce sync.Once
)

func getAVCaptureInputPortClass() AVCaptureInputPortClass {
	_AVCaptureInputPortClassOnce.Do(func() {
		_AVCaptureInputPortClass = AVCaptureInputPortClass{class: objc.GetClass("AVCaptureInputPort")}
	})
	return _AVCaptureInputPortClass
}

// GetAVCaptureInputPortClass returns the class object for AVCaptureInputPort.
func GetAVCaptureInputPortClass() AVCaptureInputPortClass {
	return getAVCaptureInputPortClass()
}

type AVCaptureInputPortClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVCaptureInputPortClass) Alloc() AVCaptureInputPort {
	rv := objc.Send[AVCaptureInputPort](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that represents a stream of data that a capture input provides.
//
// # Overview
// 
// Instances of [AVCaptureInput] have one or more input ports, one for each
// data stream they can produce. For example, an [AVCaptureDeviceInput] object
// presenting one video data stream has one port.
//
// # Inspecting an input port
//
//   - [AVCaptureInputPort.Enabled]: A Boolean value that indicates whether the port is in an enabled state.
//   - [AVCaptureInputPort.SetEnabled]
//   - [AVCaptureInputPort.MediaType]: The media type of the port.
//   - [AVCaptureInputPort.FormatDescription]: A description of the port format.
//   - [AVCaptureInputPort.Clock]: An object that represents the capture device’s clock.
//
// # Accessing the input
//
//   - [AVCaptureInputPort.Input]: The input object that owns the port.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureInput/Port
type AVCaptureInputPort struct {
	objectivec.Object
}

// AVCaptureInputPortFromID constructs a [AVCaptureInputPort] from an objc.ID.
//
// An object that represents a stream of data that a capture input provides.
func AVCaptureInputPortFromID(id objc.ID) AVCaptureInputPort {
	return AVCaptureInputPort{objectivec.Object{ID: id}}
}
// NOTE: AVCaptureInputPort adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVCaptureInputPort] class.
//
// # Inspecting an input port
//
//   - [IAVCaptureInputPort.Enabled]: A Boolean value that indicates whether the port is in an enabled state.
//   - [IAVCaptureInputPort.SetEnabled]
//   - [IAVCaptureInputPort.MediaType]: The media type of the port.
//   - [IAVCaptureInputPort.FormatDescription]: A description of the port format.
//   - [IAVCaptureInputPort.Clock]: An object that represents the capture device’s clock.
//
// # Accessing the input
//
//   - [IAVCaptureInputPort.Input]: The input object that owns the port.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureInput/Port
type IAVCaptureInputPort interface {
	objectivec.IObject

	// Topic: Inspecting an input port

	// A Boolean value that indicates whether the port is in an enabled state.
	Enabled() bool
	SetEnabled(value bool)
	// The media type of the port.
	MediaType() AVMediaType
	// A description of the port format.
	FormatDescription() objectivec.IObject
	// An object that represents the capture device’s clock.
	Clock() objectivec.IObject

	// Topic: Accessing the input

	// The input object that owns the port.
	Input() IAVCaptureInput

	// The ports available on a capture input.
	Ports() IAVCaptureInputPort
	SetPorts(value IAVCaptureInputPort)
}

// Init initializes the instance.
func (c AVCaptureInputPort) Init() AVCaptureInputPort {
	rv := objc.Send[AVCaptureInputPort](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c AVCaptureInputPort) Autorelease() AVCaptureInputPort {
	rv := objc.Send[AVCaptureInputPort](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVCaptureInputPort creates a new AVCaptureInputPort instance.
func NewAVCaptureInputPort() AVCaptureInputPort {
	class := getAVCaptureInputPortClass()
	rv := objc.Send[AVCaptureInputPort](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// A Boolean value that indicates whether the port is in an enabled state.
//
// # Discussion
// 
// Ports are in an enabled state by default. If you want to capture only a
// subset of the media streams provided by a capture input, use this property
// to selectively disable streams.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureInput/Port/isEnabled
func (c AVCaptureInputPort) Enabled() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isEnabled"))
	return rv
}
func (c AVCaptureInputPort) SetEnabled(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setEnabled:"), value)
}
// The media type of the port.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureInput/Port/mediaType
func (c AVCaptureInputPort) MediaType() AVMediaType {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("mediaType"))
	return AVMediaType(foundation.NSStringFromID(rv).String())
}
// A description of the port format.
//
// # Discussion
// 
// A format description object describes the format of the media the port
// currently provides. To observe changes to a port’s format, observe
// notifications of type [formatDescriptionDidChangeNotification].
//
// [formatDescriptionDidChangeNotification]: https://developer.apple.com/documentation/AVFoundation/AVCaptureInput/Port/formatDescriptionDidChangeNotification
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureInput/Port/formatDescription
func (c AVCaptureInputPort) FormatDescription() objectivec.IObject {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("formatDescription"))
	return objectivec.Object{ID: rv}
}
// An object that represents the capture device’s clock.
//
// # Discussion
// 
// The value of this property is readonly and may not reflect the actual clock
// in the capture device.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureInput/Port/clock
func (c AVCaptureInputPort) Clock() objectivec.IObject {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("clock"))
	return objectivec.Object{ID: rv}
}
// The input object that owns the port.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureInput/Port/input
func (c AVCaptureInputPort) Input() IAVCaptureInput {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("input"))
	return AVCaptureInputFromID(objc.ID(rv))
}
// The ports available on a capture input.
//
// See: https://developer.apple.com/documentation/avfoundation/avcaptureinput/ports
func (c AVCaptureInputPort) Ports() IAVCaptureInputPort {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("ports"))
	return AVCaptureInputPortFromID(objc.ID(rv))
}
func (c AVCaptureInputPort) SetPorts(value IAVCaptureInputPort) {
	objc.Send[struct{}](c.ID, objc.Sel("setPorts:"), value)
}

