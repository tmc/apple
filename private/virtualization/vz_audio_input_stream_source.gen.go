// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
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

// # Methods
//
//   - [VZAudioInputStreamSource._init]
//   - [VZAudioInputStreamSource._attachment]
//   - [VZAudioInputStreamSource.DebugDescription]
//   - [VZAudioInputStreamSource.Description]
//   - [VZAudioInputStreamSource.Hash]
//   - [VZAudioInputStreamSource.Superclass]
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
type IVZAudioInputStreamSource interface {
	objectivec.IObject

	// Topic: Methods

	_init() objectivec.IObject
	_attachment() unsafe.Pointer
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objectivec.Class
}

// Init initializes the instance.
func (v VZAudioInputStreamSource) Init() VZAudioInputStreamSource {
	rv := objc.Send[VZAudioInputStreamSource](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZAudioInputStreamSource) Autorelease() VZAudioInputStreamSource {
	rv := objc.Send[VZAudioInputStreamSource](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZAudioInputStreamSource creates a new VZAudioInputStreamSource instance.
func NewVZAudioInputStreamSource() VZAudioInputStreamSource {
	class := getVZAudioInputStreamSourceClass()
	rv := objc.Send[VZAudioInputStreamSource](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (v VZAudioInputStreamSource) _init() objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("_init"))
	return objectivec.Object{ID: rv}
}

func (v VZAudioInputStreamSource) _attachment() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v.ID, objc.Sel("_attachment"))
	return rv
}

// CanAttachment reports whether the receiver responds to the private selector _attachment.
func (v VZAudioInputStreamSource) CanAttachment() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_attachment"))
}

// Attachment is an exported wrapper for the private property _attachment.
func (v VZAudioInputStreamSource) Attachment() (unsafe.Pointer, error) {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_attachment")) {
		return nil, &objc.UnrecognizedSelectorError{Selector: "_attachment"}
	}
	return v._attachment(), nil
}
func (v VZAudioInputStreamSource) DebugDescription() string {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
func (v VZAudioInputStreamSource) Description() string {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
func (v VZAudioInputStreamSource) Hash() uint64 {
	rv := objc.Send[uint64](v.ID, objc.Sel("hash"))
	return rv
}
func (v VZAudioInputStreamSource) Superclass() objectivec.Class {
	rv := objc.Send[objectivec.Class](v.ID, objc.Sel("superclass"))
	return objectivec.Class(rv)
}
