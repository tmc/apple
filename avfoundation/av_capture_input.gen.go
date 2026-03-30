// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVCaptureInput] class.
var (
	_AVCaptureInputClass     AVCaptureInputClass
	_AVCaptureInputClassOnce sync.Once
)

func getAVCaptureInputClass() AVCaptureInputClass {
	_AVCaptureInputClassOnce.Do(func() {
		_AVCaptureInputClass = AVCaptureInputClass{class: objc.GetClass("AVCaptureInput")}
	})
	return _AVCaptureInputClass
}

// GetAVCaptureInputClass returns the class object for AVCaptureInput.
func GetAVCaptureInputClass() AVCaptureInputClass {
	return getAVCaptureInputClass()
}

type AVCaptureInputClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVCaptureInputClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVCaptureInputClass) Alloc() AVCaptureInput {
	rv := objc.Send[AVCaptureInput](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An abstract superclass for objects that provide input data to a capture
// session.
//
// # Overview
//
// You create concrete instances of this class, such as
// [AVCaptureDeviceInput], to add inputs to a capture session. An input
// provides one or more streams of media data. For example, input devices can
// provide both audio and video data. The framework represents each media
// stream that an input provides as an [AVCaptureInputPort] object.
//
// A capture makes connections between capture inputs and capture outputs
// using a [AVCaptureConnection] object. The connection defines the mapping
// between a set of port objects and an [AVCaptureOutput].
//
// # Accessing ports
//
//   - [AVCaptureInput.Ports]: The ports available on a capture input.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureInput
type AVCaptureInput struct {
	objectivec.Object
}

// AVCaptureInputFromID constructs a [AVCaptureInput] from an objc.ID.
//
// An abstract superclass for objects that provide input data to a capture
// session.
func AVCaptureInputFromID(id objc.ID) AVCaptureInput {
	return AVCaptureInput{objectivec.Object{ID: id}}
}

// NOTE: AVCaptureInput adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVCaptureInput] class.
//
// # Accessing ports
//
//   - [IAVCaptureInput.Ports]: The ports available on a capture input.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureInput
type IAVCaptureInput interface {
	objectivec.IObject

	// Topic: Accessing ports

	// The ports available on a capture input.
	Ports() []AVCaptureInputPort
}

// Init initializes the instance.
func (c AVCaptureInput) Init() AVCaptureInput {
	rv := objc.Send[AVCaptureInput](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c AVCaptureInput) Autorelease() AVCaptureInput {
	rv := objc.Send[AVCaptureInput](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVCaptureInput creates a new AVCaptureInput instance.
func NewAVCaptureInput() AVCaptureInput {
	class := getAVCaptureInputClass()
	rv := objc.Send[AVCaptureInput](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The ports available on a capture input.
//
// # Discussion
//
// Individual ports post an [formatDescriptionDidChangeNotification]
// notification when their [FormatDescription] changes.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureInput/ports
//
// [formatDescriptionDidChangeNotification]: https://developer.apple.com/documentation/AVFoundation/AVCaptureInput/Port/formatDescriptionDidChangeNotification
func (c AVCaptureInput) Ports() []AVCaptureInputPort {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("ports"))
	return objc.ConvertSlice(rv, func(id objc.ID) AVCaptureInputPort {
		return AVCaptureInputPortFromID(id)
	})
}
