// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [VZHostAudioOutputStreamSink] class.
var (
	_VZHostAudioOutputStreamSinkClass     VZHostAudioOutputStreamSinkClass
	_VZHostAudioOutputStreamSinkClassOnce sync.Once
)

func getVZHostAudioOutputStreamSinkClass() VZHostAudioOutputStreamSinkClass {
	_VZHostAudioOutputStreamSinkClassOnce.Do(func() {
		_VZHostAudioOutputStreamSinkClass = VZHostAudioOutputStreamSinkClass{class: objc.GetClass("VZHostAudioOutputStreamSink")}
	})
	return _VZHostAudioOutputStreamSinkClass
}

// GetVZHostAudioOutputStreamSinkClass returns the class object for VZHostAudioOutputStreamSink.
func GetVZHostAudioOutputStreamSinkClass() VZHostAudioOutputStreamSinkClass {
	return getVZHostAudioOutputStreamSinkClass()
}

type VZHostAudioOutputStreamSinkClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZHostAudioOutputStreamSinkClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZHostAudioOutputStreamSinkClass) Alloc() VZHostAudioOutputStreamSink {
	rv := objc.Send[VZHostAudioOutputStreamSink](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// Host audio output stream sink plays audio to the host system’s default
// output device.
//
// # Overview
//
// Host output data goes to the same device that
// [AudioQueueNewOutput(_:_:_:_:_:_:_:)] uses.
//
// See: https://developer.apple.com/documentation/Virtualization/VZHostAudioOutputStreamSink
//
// [AudioQueueNewOutput(_:_:_:_:_:_:_:)]: https://developer.apple.com/documentation/AudioToolbox/AudioQueueNewOutput(_:_:_:_:_:_:_:)
type VZHostAudioOutputStreamSink struct {
	VZAudioOutputStreamSink
}

// VZHostAudioOutputStreamSinkFromID constructs a [VZHostAudioOutputStreamSink] from an objc.ID.
//
// Host audio output stream sink plays audio to the host system’s default
// output device.
func VZHostAudioOutputStreamSinkFromID(id objc.ID) VZHostAudioOutputStreamSink {
	return VZHostAudioOutputStreamSink{VZAudioOutputStreamSink: VZAudioOutputStreamSinkFromID(id)}
}

// NOTE: VZHostAudioOutputStreamSink adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [VZHostAudioOutputStreamSink] class.
//
// See: https://developer.apple.com/documentation/Virtualization/VZHostAudioOutputStreamSink
type IVZHostAudioOutputStreamSink interface {
	IVZAudioOutputStreamSink
}

// Init initializes the instance.
func (h VZHostAudioOutputStreamSink) Init() VZHostAudioOutputStreamSink {
	rv := objc.Send[VZHostAudioOutputStreamSink](h.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h VZHostAudioOutputStreamSink) Autorelease() VZHostAudioOutputStreamSink {
	rv := objc.Send[VZHostAudioOutputStreamSink](h.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZHostAudioOutputStreamSink creates a new VZHostAudioOutputStreamSink instance.
func NewVZHostAudioOutputStreamSink() VZHostAudioOutputStreamSink {
	class := getVZHostAudioOutputStreamSinkClass()
	rv := objc.Send[VZHostAudioOutputStreamSink](objc.ID(class.class), objc.Sel("new"))
	return rv
}
