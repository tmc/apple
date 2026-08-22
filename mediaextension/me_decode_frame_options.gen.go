// Code generated from Apple documentation for MediaExtension. DO NOT EDIT.

package mediaextension

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MEDecodeFrameOptions] class.
var (
	_MEDecodeFrameOptionsClass     MEDecodeFrameOptionsClass
	_MEDecodeFrameOptionsClassOnce sync.Once
)

func getMEDecodeFrameOptionsClass() MEDecodeFrameOptionsClass {
	_MEDecodeFrameOptionsClassOnce.Do(func() {
		_MEDecodeFrameOptionsClass = MEDecodeFrameOptionsClass{class: objc.GetClass("MEDecodeFrameOptions")}
	})
	return _MEDecodeFrameOptionsClass
}

// GetMEDecodeFrameOptionsClass returns the class object for MEDecodeFrameOptions.
func GetMEDecodeFrameOptionsClass() MEDecodeFrameOptionsClass {
	return getMEDecodeFrameOptionsClass()
}

type MEDecodeFrameOptionsClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MEDecodeFrameOptionsClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MEDecodeFrameOptionsClass) Alloc() MEDecodeFrameOptions {
	rv := objc.Send[MEDecodeFrameOptions](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// An object that guides the video decoder operation on a per-frame basis.
//
// # Inspecting frame decoding options
//
//   - [MEDecodeFrameOptions.DoNotOutputFrame]: A Boolean value that hints to the decoder whether or not it should emit an image buffer for the frame.
//   - [MEDecodeFrameOptions.SetDoNotOutputFrame]
//   - [MEDecodeFrameOptions.RealTimePlayback]: A Boolean value that hints to the decoder to use a low-power mode that can’t decode faster than 1x real-time.
//   - [MEDecodeFrameOptions.SetRealTimePlayback]
//
// See: https://developer.apple.com/documentation/MediaExtension/MEDecodeFrameOptions
type MEDecodeFrameOptions struct {
	objectivec.Object
}

// MEDecodeFrameOptionsFromID constructs a [MEDecodeFrameOptions] from an objc.ID.
//
// An object that guides the video decoder operation on a per-frame basis.
func MEDecodeFrameOptionsFromID(id objc.ID) MEDecodeFrameOptions {
	return MEDecodeFrameOptions{objectivec.Object{ID: id}}
}

// NOTE: MEDecodeFrameOptions adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MEDecodeFrameOptions] class.
//
// # Inspecting frame decoding options
//
//   - [IMEDecodeFrameOptions.DoNotOutputFrame]: A Boolean value that hints to the decoder whether or not it should emit an image buffer for the frame.
//   - [IMEDecodeFrameOptions.SetDoNotOutputFrame]
//   - [IMEDecodeFrameOptions.RealTimePlayback]: A Boolean value that hints to the decoder to use a low-power mode that can’t decode faster than 1x real-time.
//   - [IMEDecodeFrameOptions.SetRealTimePlayback]
//
// See: https://developer.apple.com/documentation/MediaExtension/MEDecodeFrameOptions
type IMEDecodeFrameOptions interface {
	objectivec.IObject

	// Topic: Inspecting frame decoding options

	// A Boolean value that hints to the decoder whether or not it should emit an image buffer for the frame.
	DoNotOutputFrame() bool
	SetDoNotOutputFrame(value bool)
	// A Boolean value that hints to the decoder to use a low-power mode that can’t decode faster than 1x real-time.
	RealTimePlayback() bool
	SetRealTimePlayback(value bool)
}

// Init initializes the instance.
func (m MEDecodeFrameOptions) Init() MEDecodeFrameOptions {
	rv := objc.Send[MEDecodeFrameOptions](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MEDecodeFrameOptions) Autorelease() MEDecodeFrameOptions {
	rv := objc.Send[MEDecodeFrameOptions](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMEDecodeFrameOptions creates a new MEDecodeFrameOptions instance.
func NewMEDecodeFrameOptions() MEDecodeFrameOptions {
	class := getMEDecodeFrameOptionsClass()
	rv := objc.Send[MEDecodeFrameOptions](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// A Boolean value that hints to the decoder whether or not it should emit an
// image buffer for the frame.
//
// # Discussion
//
// If true, the decoder emits nil instead of a [CVImageBuffer] instance.
//
// See: https://developer.apple.com/documentation/MediaExtension/MEDecodeFrameOptions/doNotOutputFrame
//
// [CVImageBuffer]: https://developer.apple.com/documentation/CoreVideo/cvimagebuffer-q40
func (m MEDecodeFrameOptions) DoNotOutputFrame() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("doNotOutputFrame"))
	return rv
}
func (m MEDecodeFrameOptions) SetDoNotOutputFrame(value bool) {
	objc.Send[struct{}](m.ID, objc.Sel("setDoNotOutputFrame:"), value)
}

// A Boolean value that hints to the decoder to use a low-power mode that
// can’t decode faster than 1x real-time.
//
// # Discussion
//
// The system sets this value to false during all uses other than 1x forward
// real-time playback, including seeking, playback at other rates, and export.
//
// This hint only applies to the current decode session. If multiple instances
// of a decoder operate at the same time, it may not be acceptable to use a
// low-power mode if the system can’t sustain real-time playback across all
// the streams.
//
// See: https://developer.apple.com/documentation/MediaExtension/MEDecodeFrameOptions/realTimePlayback
func (m MEDecodeFrameOptions) RealTimePlayback() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("realTimePlayback"))
	return rv
}
func (m MEDecodeFrameOptions) SetRealTimePlayback(value bool) {
	objc.Send[struct{}](m.ID, objc.Sel("setRealTimePlayback:"), value)
}
