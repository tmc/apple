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

// Class returns the underlying Objective-C class pointer.
func (vc VZHostAudioInputStreamSourceClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZHostAudioInputStreamSourceClass) Alloc() VZHostAudioInputStreamSource {
	rv := objc.SendIfResponds[VZHostAudioInputStreamSource](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

type VZHostAudioInputStreamSource struct {
	VZAudioInputStreamSource
}

// VZHostAudioInputStreamSourceFromID constructs a [VZHostAudioInputStreamSource] from an objc.ID.
func VZHostAudioInputStreamSourceFromID(id objc.ID) VZHostAudioInputStreamSource {
	return VZHostAudioInputStreamSource{VZAudioInputStreamSource: VZAudioInputStreamSourceFromID(id)}
}

// Ensure VZHostAudioInputStreamSource implements IVZHostAudioInputStreamSource.
var _ IVZHostAudioInputStreamSource = VZHostAudioInputStreamSource{}

// An interface definition for the [VZHostAudioInputStreamSource] class.
type IVZHostAudioInputStreamSource interface {
	IVZAudioInputStreamSource
}

// Init initializes the instance.
func (v VZHostAudioInputStreamSource) Init() VZHostAudioInputStreamSource {
	rv := objc.SendIfResponds[VZHostAudioInputStreamSource](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZHostAudioInputStreamSource) Autorelease() VZHostAudioInputStreamSource {
	rv := objc.SendIfResponds[VZHostAudioInputStreamSource](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZHostAudioInputStreamSource creates a new VZHostAudioInputStreamSource instance.
func NewVZHostAudioInputStreamSource() VZHostAudioInputStreamSource {
	class := getVZHostAudioInputStreamSourceClass()
	rv := objc.SendIfResponds[VZHostAudioInputStreamSource](objc.ID(class.class), objc.Sel("new"))
	return rv
}
