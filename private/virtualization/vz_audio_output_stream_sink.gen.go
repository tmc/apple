// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"

	"github.com/tmc/apple/foundation"
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

// # Methods
//
//   - [VZAudioOutputStreamSink._init]
//   - [VZAudioOutputStreamSink._attachment]
//   - [VZAudioOutputStreamSink.DebugDescription]
//   - [VZAudioOutputStreamSink.Description]
//   - [VZAudioOutputStreamSink.Hash]
//   - [VZAudioOutputStreamSink.Superclass]
//
// See: https://developer.apple.com/documentation/Virtualization/VZAudioOutputStreamSink
type VZAudioOutputStreamSink struct {
	objectivec.Object
}

// VZAudioOutputStreamSinkFromID constructs a [VZAudioOutputStreamSink] from an objc.ID.
func VZAudioOutputStreamSinkFromID(id objc.ID) VZAudioOutputStreamSink {
	return VZAudioOutputStreamSink{objectivec.Object{ID: id}}
}

// Ensure VZAudioOutputStreamSink implements IVZAudioOutputStreamSink.
var _ IVZAudioOutputStreamSink = VZAudioOutputStreamSink{}

// An interface definition for the [VZAudioOutputStreamSink] class.
//
// # Methods
//
//   - [IVZAudioOutputStreamSink._init]
//   - [IVZAudioOutputStreamSink._attachment]
//   - [IVZAudioOutputStreamSink.DebugDescription]
//   - [IVZAudioOutputStreamSink.Description]
//   - [IVZAudioOutputStreamSink.Hash]
//   - [IVZAudioOutputStreamSink.Superclass]
//
// See: https://developer.apple.com/documentation/Virtualization/VZAudioOutputStreamSink
type IVZAudioOutputStreamSink interface {
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
func (v VZAudioOutputStreamSink) Init() VZAudioOutputStreamSink {
	rv := objc.Send[VZAudioOutputStreamSink](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZAudioOutputStreamSink) Autorelease() VZAudioOutputStreamSink {
	rv := objc.Send[VZAudioOutputStreamSink](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZAudioOutputStreamSink creates a new VZAudioOutputStreamSink instance.
func NewVZAudioOutputStreamSink() VZAudioOutputStreamSink {
	class := getVZAudioOutputStreamSinkClass()
	rv := objc.Send[VZAudioOutputStreamSink](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/VZAudioOutputStreamSink/_init
func (v VZAudioOutputStreamSink) _init() objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("_init"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/Virtualization/VZAudioOutputStreamSink/_attachment
func (v VZAudioOutputStreamSink) _attachment() objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("_attachment"))
	return objectivec.Object{ID: rv}
}

// CanAttachment reports whether the receiver responds to the private selector _attachment.
func (v VZAudioOutputStreamSink) CanAttachment() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_attachment"))
}

// Attachment is an exported wrapper for the private property _attachment.
func (v VZAudioOutputStreamSink) Attachment() (objectivec.IObject, error) {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_attachment")) {
		return nil, &objc.UnrecognizedSelectorError{Selector: "_attachment"}
	}
	return v._attachment(), nil
}

// See: https://developer.apple.com/documentation/Virtualization/VZAudioOutputStreamSink/debugDescription
func (v VZAudioOutputStreamSink) DebugDescription() string {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/Virtualization/VZAudioOutputStreamSink/description
func (v VZAudioOutputStreamSink) Description() string {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/Virtualization/VZAudioOutputStreamSink/hash
func (v VZAudioOutputStreamSink) Hash() uint64 {
	rv := objc.Send[uint64](v.ID, objc.Sel("hash"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/VZAudioOutputStreamSink/superclass
func (v VZAudioOutputStreamSink) Superclass() objc.Class {
	rv := objc.Send[objc.Class](v.ID, objc.Sel("superclass"))
	return rv
}
