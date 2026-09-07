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
	rv := objc.SendIfResponds[VZHostAudioOutputStreamSink](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

type VZHostAudioOutputStreamSink struct {
	VZAudioOutputStreamSink
}

// VZHostAudioOutputStreamSinkFromID constructs a [VZHostAudioOutputStreamSink] from an objc.ID.
func VZHostAudioOutputStreamSinkFromID(id objc.ID) VZHostAudioOutputStreamSink {
	return VZHostAudioOutputStreamSink{VZAudioOutputStreamSink: VZAudioOutputStreamSinkFromID(id)}
}

// Ensure VZHostAudioOutputStreamSink implements IVZHostAudioOutputStreamSink.
var _ IVZHostAudioOutputStreamSink = VZHostAudioOutputStreamSink{}

// An interface definition for the [VZHostAudioOutputStreamSink] class.
type IVZHostAudioOutputStreamSink interface {
	IVZAudioOutputStreamSink
}

// Init initializes the instance.
func (v VZHostAudioOutputStreamSink) Init() VZHostAudioOutputStreamSink {
	rv := objc.SendIfResponds[VZHostAudioOutputStreamSink](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZHostAudioOutputStreamSink) Autorelease() VZHostAudioOutputStreamSink {
	rv := objc.SendIfResponds[VZHostAudioOutputStreamSink](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZHostAudioOutputStreamSink creates a new VZHostAudioOutputStreamSink instance.
func NewVZHostAudioOutputStreamSink() VZHostAudioOutputStreamSink {
	class := getVZHostAudioOutputStreamSinkClass()
	rv := objc.SendIfResponds[VZHostAudioOutputStreamSink](objc.ID(class.class), objc.Sel("new"))
	return rv
}
