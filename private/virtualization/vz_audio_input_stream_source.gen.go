// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
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

//
// # Methods
//
//   - [VZAudioInputStreamSource._init]
//   - [VZAudioInputStreamSource._attachment]
//   - [VZAudioInputStreamSource.DebugDescription]
//   - [VZAudioInputStreamSource.Description]
//   - [VZAudioInputStreamSource.Hash]
//   - [VZAudioInputStreamSource.Superclass]
// See: https://developer.apple.com/documentation/Virtualization/VZAudioInputStreamSource
type VZAudioInputStreamSource struct {
	objectivec.Object
}

// VZAudioInputStreamSourceFromID constructs a [VZAudioInputStreamSource] from an objc.ID.
func VZAudioInputStreamSourceFromID(id objc.ID) VZAudioInputStreamSource {
	return VZAudioInputStreamSource{objectivec.Object{ID: id}}
}
// Ensure VZAudioInputStreamSource implements IVZAudioInputStreamSource.
var _ IVZAudioInputStreamSource = VZAudioInputStreamSource{}

// An interface definition for the [VZAudioInputStreamSource] class.
//
// # Methods
//
//   - [IVZAudioInputStreamSource._init]
//   - [IVZAudioInputStreamSource._attachment]
//   - [IVZAudioInputStreamSource.DebugDescription]
//   - [IVZAudioInputStreamSource.Description]
//   - [IVZAudioInputStreamSource.Hash]
//   - [IVZAudioInputStreamSource.Superclass]
//
// See: https://developer.apple.com/documentation/Virtualization/VZAudioInputStreamSource
type IVZAudioInputStreamSource interface {
	objectivec.IObject

	// Topic: Methods

	_init() objectivec.IObject
	_attachment() objectivec.IObject
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objc.Class
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

// See: https://developer.apple.com/documentation/Virtualization/VZAudioInputStreamSource/_init
func (a VZAudioInputStreamSource) _init() objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("_init"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/Virtualization/VZAudioInputStreamSource/_attachment
func (a VZAudioInputStreamSource) _attachment() objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("_attachment"))
	return objectivec.Object{ID: rv}
}
// See: https://developer.apple.com/documentation/Virtualization/VZAudioInputStreamSource/debugDescription
func (a VZAudioInputStreamSource) DebugDescription() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/Virtualization/VZAudioInputStreamSource/description
func (a VZAudioInputStreamSource) Description() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/Virtualization/VZAudioInputStreamSource/hash
func (a VZAudioInputStreamSource) Hash() uint64 {
	rv := objc.Send[uint64](a.ID, objc.Sel("hash"))
	return rv
}
// See: https://developer.apple.com/documentation/Virtualization/VZAudioInputStreamSource/superclass
func (a VZAudioInputStreamSource) Superclass() objc.Class {
	rv := objc.Send[objc.Class](a.ID, objc.Sel("superclass"))
	return rv
}

