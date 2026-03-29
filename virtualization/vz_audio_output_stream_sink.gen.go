// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZAudioOutputStreamSink] class.
var (
	_VZAudioOutputStreamSinkClass     VZAudioOutputStreamSinkClass
	_VZAudioOutputStreamSinkClassOnce sync.Once
)

func getVZAudioOutputStreamSinkClass() VZAudioOutputStreamSinkClass {
	_VZAudioOutputStreamSinkClassOnce.Do(func() {
		_VZAudioOutputStreamSinkClass = VZAudioOutputStreamSinkClass{class: objc.GetClass("VZAudioOutputStreamSink")}
	})
	return _VZAudioOutputStreamSinkClass
}

// GetVZAudioOutputStreamSinkClass returns the class object for VZAudioOutputStreamSink.
func GetVZAudioOutputStreamSinkClass() VZAudioOutputStreamSinkClass {
	return getVZAudioOutputStreamSinkClass()
}

type VZAudioOutputStreamSinkClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZAudioOutputStreamSinkClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZAudioOutputStreamSinkClass) Alloc() VZAudioOutputStreamSink {
	rv := objc.Send[VZAudioOutputStreamSink](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// The base class for an audio output stream sink.
//
// # Overview
// 
// An audio output stream sink defines how the host system consumes audio data
// from a guest.
// 
// Don’t instantiate [VZAudioOutputStreamSink] directly, use one of its
// subclasses, such as [VZHostAudioOutputStreamSink] instead.
//
// See: https://developer.apple.com/documentation/Virtualization/VZAudioOutputStreamSink
type VZAudioOutputStreamSink struct {
	objectivec.Object
}

// VZAudioOutputStreamSinkFromID constructs a [VZAudioOutputStreamSink] from an objc.ID.
//
// The base class for an audio output stream sink.
func VZAudioOutputStreamSinkFromID(id objc.ID) VZAudioOutputStreamSink {
	return VZAudioOutputStreamSink{objectivec.Object{ID: id}}
}
// NOTE: VZAudioOutputStreamSink adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [VZAudioOutputStreamSink] class.
//
// See: https://developer.apple.com/documentation/Virtualization/VZAudioOutputStreamSink
type IVZAudioOutputStreamSink interface {
	objectivec.IObject
}

// Init initializes the instance.
func (a VZAudioOutputStreamSink) Init() VZAudioOutputStreamSink {
	rv := objc.Send[VZAudioOutputStreamSink](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a VZAudioOutputStreamSink) Autorelease() VZAudioOutputStreamSink {
	rv := objc.Send[VZAudioOutputStreamSink](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZAudioOutputStreamSink creates a new VZAudioOutputStreamSink instance.
func NewVZAudioOutputStreamSink() VZAudioOutputStreamSink {
	class := getVZAudioOutputStreamSinkClass()
	rv := objc.Send[VZAudioOutputStreamSink](objc.ID(class.class), objc.Sel("new"))
	return rv
}

