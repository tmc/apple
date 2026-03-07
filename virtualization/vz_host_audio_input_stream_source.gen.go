// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [VZHostAudioInputStreamSource] class.
var (
	_VZHostAudioInputStreamSourceClass     VZHostAudioInputStreamSourceClass
	_VZHostAudioInputStreamSourceClassOnce sync.Once
)

func getVZHostAudioInputStreamSourceClass() VZHostAudioInputStreamSourceClass {
	_VZHostAudioInputStreamSourceClassOnce.Do(func() {
		_VZHostAudioInputStreamSourceClass = VZHostAudioInputStreamSourceClass{class: objc.GetClass("VZHostAudioInputStreamSource")}
	})
	return _VZHostAudioInputStreamSourceClass
}

// GetVZHostAudioInputStreamSourceClass returns the class object for VZHostAudioInputStreamSource.
func GetVZHostAudioInputStreamSourceClass() VZHostAudioInputStreamSourceClass {
	return getVZHostAudioInputStreamSourceClass()
}

type VZHostAudioInputStreamSourceClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZHostAudioInputStreamSourceClass) Alloc() VZHostAudioInputStreamSource {
	rv := objc.Send[VZHostAudioInputStreamSource](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}







// The host audio input stream source that provides audio from the host
// system’s default input device.
//
// # Overview
// 
// The host input data comes from the same device that
// [AudioQueueNewInput(_:_:_:_:_:_:_:)] uses.
//
// [AudioQueueNewInput(_:_:_:_:_:_:_:)]: https://developer.apple.com/documentation/AudioToolbox/AudioQueueNewInput(_:_:_:_:_:_:_:)
//
// See: https://developer.apple.com/documentation/Virtualization/VZHostAudioInputStreamSource
type VZHostAudioInputStreamSource struct {
	VZAudioInputStreamSource
}

// VZHostAudioInputStreamSourceFromID constructs a [VZHostAudioInputStreamSource] from an objc.ID.
//
// The host audio input stream source that provides audio from the host
// system’s default input device.
func VZHostAudioInputStreamSourceFromID(id objc.ID) VZHostAudioInputStreamSource {
	return VZHostAudioInputStreamSource{VZAudioInputStreamSource: VZAudioInputStreamSourceFromID(id)}
}
// NOTE: VZHostAudioInputStreamSource adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [VZHostAudioInputStreamSource] class.
//
// See: https://developer.apple.com/documentation/Virtualization/VZHostAudioInputStreamSource
type IVZHostAudioInputStreamSource interface {
	IVZAudioInputStreamSource
}





// Init initializes the instance.
func (h VZHostAudioInputStreamSource) Init() VZHostAudioInputStreamSource {
	rv := objc.Send[VZHostAudioInputStreamSource](h.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h VZHostAudioInputStreamSource) Autorelease() VZHostAudioInputStreamSource {
	rv := objc.Send[VZHostAudioInputStreamSource](h.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZHostAudioInputStreamSource creates a new VZHostAudioInputStreamSource instance.
func NewVZHostAudioInputStreamSource() VZHostAudioInputStreamSource {
	class := getVZHostAudioInputStreamSourceClass()
	rv := objc.Send[VZHostAudioInputStreamSource](objc.ID(class.class), objc.Sel("new"))
	return rv
}










































