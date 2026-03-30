// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZAudioInputStreamSource] class.
var (
	_VZAudioInputStreamSourceClass     VZAudioInputStreamSourceClass
	_VZAudioInputStreamSourceClassOnce sync.Once
)

func getVZAudioInputStreamSourceClass() VZAudioInputStreamSourceClass {
	_VZAudioInputStreamSourceClassOnce.Do(func() {
		_VZAudioInputStreamSourceClass = VZAudioInputStreamSourceClass{class: objc.GetClass("VZAudioInputStreamSource")}
	})
	return _VZAudioInputStreamSourceClass
}

// GetVZAudioInputStreamSourceClass returns the class object for VZAudioInputStreamSource.
func GetVZAudioInputStreamSourceClass() VZAudioInputStreamSourceClass {
	return getVZAudioInputStreamSourceClass()
}

type VZAudioInputStreamSourceClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZAudioInputStreamSourceClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZAudioInputStreamSourceClass) Alloc() VZAudioInputStreamSource {
	rv := objc.Send[VZAudioInputStreamSource](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// The base class for an audio input stream source.
//
// # Overview
//
// An audio input stream source defines how th guest produces audio input data
// on the host system.
//
// Don’t instantiate [VZAudioInputStreamSource] directly, use one of its
// subclasses such as [VZHostAudioInputStreamSource] instead.
//
// See: https://developer.apple.com/documentation/Virtualization/VZAudioInputStreamSource
type VZAudioInputStreamSource struct {
	objectivec.Object
}

// VZAudioInputStreamSourceFromID constructs a [VZAudioInputStreamSource] from an objc.ID.
//
// The base class for an audio input stream source.
func VZAudioInputStreamSourceFromID(id objc.ID) VZAudioInputStreamSource {
	return VZAudioInputStreamSource{objectivec.Object{ID: id}}
}

// NOTE: VZAudioInputStreamSource adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [VZAudioInputStreamSource] class.
//
// See: https://developer.apple.com/documentation/Virtualization/VZAudioInputStreamSource
type IVZAudioInputStreamSource interface {
	objectivec.IObject
}

// Init initializes the instance.
func (a VZAudioInputStreamSource) Init() VZAudioInputStreamSource {
	rv := objc.Send[VZAudioInputStreamSource](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a VZAudioInputStreamSource) Autorelease() VZAudioInputStreamSource {
	rv := objc.Send[VZAudioInputStreamSource](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZAudioInputStreamSource creates a new VZAudioInputStreamSource instance.
func NewVZAudioInputStreamSource() VZAudioInputStreamSource {
	class := getVZAudioInputStreamSourceClass()
	rv := objc.Send[VZAudioInputStreamSource](objc.ID(class.class), objc.Sel("new"))
	return rv
}
