// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZVirtioSoundDeviceStreamConfiguration] class.
var (
	_VZVirtioSoundDeviceStreamConfigurationClass     VZVirtioSoundDeviceStreamConfigurationClass
	_VZVirtioSoundDeviceStreamConfigurationClassOnce sync.Once
)

func getVZVirtioSoundDeviceStreamConfigurationClass() VZVirtioSoundDeviceStreamConfigurationClass {
	_VZVirtioSoundDeviceStreamConfigurationClassOnce.Do(func() {
		_VZVirtioSoundDeviceStreamConfigurationClass = VZVirtioSoundDeviceStreamConfigurationClass{class: objc.GetClass("VZVirtioSoundDeviceStreamConfiguration")}
	})
	return _VZVirtioSoundDeviceStreamConfigurationClass
}

// GetVZVirtioSoundDeviceStreamConfigurationClass returns the class object for VZVirtioSoundDeviceStreamConfiguration.
func GetVZVirtioSoundDeviceStreamConfigurationClass() VZVirtioSoundDeviceStreamConfigurationClass {
	return getVZVirtioSoundDeviceStreamConfigurationClass()
}

type VZVirtioSoundDeviceStreamConfigurationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZVirtioSoundDeviceStreamConfigurationClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZVirtioSoundDeviceStreamConfigurationClass) Alloc() VZVirtioSoundDeviceStreamConfiguration {
	rv := objc.Send[VZVirtioSoundDeviceStreamConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [VZVirtioSoundDeviceStreamConfiguration._init]
//   - [VZVirtioSoundDeviceStreamConfiguration._stream]
//   - [VZVirtioSoundDeviceStreamConfiguration.DebugDescription]
//   - [VZVirtioSoundDeviceStreamConfiguration.Description]
//   - [VZVirtioSoundDeviceStreamConfiguration.Hash]
//   - [VZVirtioSoundDeviceStreamConfiguration.Superclass]
type VZVirtioSoundDeviceStreamConfiguration struct {
	objectivec.Object
}

// VZVirtioSoundDeviceStreamConfigurationFromID constructs a [VZVirtioSoundDeviceStreamConfiguration] from an objc.ID.
func VZVirtioSoundDeviceStreamConfigurationFromID(id objc.ID) VZVirtioSoundDeviceStreamConfiguration {
	return VZVirtioSoundDeviceStreamConfiguration{objectivec.Object{ID: id}}
}

// Ensure VZVirtioSoundDeviceStreamConfiguration implements IVZVirtioSoundDeviceStreamConfiguration.
var _ IVZVirtioSoundDeviceStreamConfiguration = VZVirtioSoundDeviceStreamConfiguration{}

// An interface definition for the [VZVirtioSoundDeviceStreamConfiguration] class.
//
// # Methods
//
//   - [IVZVirtioSoundDeviceStreamConfiguration._init]
//   - [IVZVirtioSoundDeviceStreamConfiguration._stream]
//   - [IVZVirtioSoundDeviceStreamConfiguration.DebugDescription]
//   - [IVZVirtioSoundDeviceStreamConfiguration.Description]
//   - [IVZVirtioSoundDeviceStreamConfiguration.Hash]
//   - [IVZVirtioSoundDeviceStreamConfiguration.Superclass]
type IVZVirtioSoundDeviceStreamConfiguration interface {
	objectivec.IObject

	// Topic: Methods

	_init() objectivec.IObject
	_stream() unsafe.Pointer
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objectivec.Class
}

// Init initializes the instance.
func (v VZVirtioSoundDeviceStreamConfiguration) Init() VZVirtioSoundDeviceStreamConfiguration {
	rv := objc.Send[VZVirtioSoundDeviceStreamConfiguration](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZVirtioSoundDeviceStreamConfiguration) Autorelease() VZVirtioSoundDeviceStreamConfiguration {
	rv := objc.Send[VZVirtioSoundDeviceStreamConfiguration](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZVirtioSoundDeviceStreamConfiguration creates a new VZVirtioSoundDeviceStreamConfiguration instance.
func NewVZVirtioSoundDeviceStreamConfiguration() VZVirtioSoundDeviceStreamConfiguration {
	class := getVZVirtioSoundDeviceStreamConfigurationClass()
	rv := objc.Send[VZVirtioSoundDeviceStreamConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (v VZVirtioSoundDeviceStreamConfiguration) _init() objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("_init"))
	return objectivec.Object{ID: rv}
}

func (v VZVirtioSoundDeviceStreamConfiguration) _stream() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v.ID, objc.Sel("_stream"))
	return rv
}

// CanStream reports whether the receiver responds to the private selector _stream.
func (v VZVirtioSoundDeviceStreamConfiguration) CanStream() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_stream"))
}

// Stream is an exported wrapper for the private property _stream.
func (v VZVirtioSoundDeviceStreamConfiguration) Stream() (unsafe.Pointer, error) {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_stream")) {
		return nil, &objc.UnrecognizedSelectorError{Selector: "_stream"}
	}
	return v._stream(), nil
}
func (v VZVirtioSoundDeviceStreamConfiguration) DebugDescription() string {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
func (v VZVirtioSoundDeviceStreamConfiguration) Description() string {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
func (v VZVirtioSoundDeviceStreamConfiguration) Hash() uint64 {
	rv := objc.Send[uint64](v.ID, objc.Sel("hash"))
	return rv
}
func (v VZVirtioSoundDeviceStreamConfiguration) Superclass() objectivec.Class {
	rv := objc.Send[objectivec.Class](v.ID, objc.Sel("superclass"))
	return objectivec.Class(rv)
}
