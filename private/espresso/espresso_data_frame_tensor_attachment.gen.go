// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corevideo"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [EspressoDataFrameTensorAttachment] class.
var (
	_EspressoDataFrameTensorAttachmentClass     EspressoDataFrameTensorAttachmentClass
	_EspressoDataFrameTensorAttachmentClassOnce sync.Once
)

func getEspressoDataFrameTensorAttachmentClass() EspressoDataFrameTensorAttachmentClass {
	_EspressoDataFrameTensorAttachmentClassOnce.Do(func() {
		_EspressoDataFrameTensorAttachmentClass = EspressoDataFrameTensorAttachmentClass{class: objc.GetClass("EspressoDataFrameTensorAttachment")}
	})
	return _EspressoDataFrameTensorAttachmentClass
}

// GetEspressoDataFrameTensorAttachmentClass returns the class object for EspressoDataFrameTensorAttachment.
func GetEspressoDataFrameTensorAttachmentClass() EspressoDataFrameTensorAttachmentClass {
	return getEspressoDataFrameTensorAttachmentClass()
}

type EspressoDataFrameTensorAttachmentClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (ec EspressoDataFrameTensorAttachmentClass) Alloc() EspressoDataFrameTensorAttachment {
	rv := objc.Send[EspressoDataFrameTensorAttachment](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [EspressoDataFrameTensorAttachment.CopyAsEspressoBuffer]
// See: https://developer.apple.com/documentation/Espresso/EspressoDataFrameTensorAttachment
type EspressoDataFrameTensorAttachment struct {
	EspressoDataFrameAttachment
}

// EspressoDataFrameTensorAttachmentFromID constructs a [EspressoDataFrameTensorAttachment] from an objc.ID.
func EspressoDataFrameTensorAttachmentFromID(id objc.ID) EspressoDataFrameTensorAttachment {
	return EspressoDataFrameTensorAttachment{EspressoDataFrameAttachment: EspressoDataFrameAttachmentFromID(id)}
}
// Ensure EspressoDataFrameTensorAttachment implements IEspressoDataFrameTensorAttachment.
var _ IEspressoDataFrameTensorAttachment = EspressoDataFrameTensorAttachment{}

// An interface definition for the [EspressoDataFrameTensorAttachment] class.
//
// # Methods
//
//   - [IEspressoDataFrameTensorAttachment.CopyAsEspressoBuffer]
//
// See: https://developer.apple.com/documentation/Espresso/EspressoDataFrameTensorAttachment
type IEspressoDataFrameTensorAttachment interface {
	IEspressoDataFrameAttachment

	// Topic: Methods

	CopyAsEspressoBuffer() objectivec.IObject
}

// Init initializes the instance.
func (e EspressoDataFrameTensorAttachment) Init() EspressoDataFrameTensorAttachment {
	rv := objc.Send[EspressoDataFrameTensorAttachment](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e EspressoDataFrameTensorAttachment) Autorelease() EspressoDataFrameTensorAttachment {
	rv := objc.Send[EspressoDataFrameTensorAttachment](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewEspressoDataFrameTensorAttachment creates a new EspressoDataFrameTensorAttachment instance.
func NewEspressoDataFrameTensorAttachment() EspressoDataFrameTensorAttachment {
	class := getEspressoDataFrameTensorAttachmentClass()
	rv := objc.Send[EspressoDataFrameTensorAttachment](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Espresso/EspressoDataFrameTensorAttachment/copyAsEspressoBuffer
func (e EspressoDataFrameTensorAttachment) CopyAsEspressoBuffer() objectivec.IObject {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("copyAsEspressoBuffer"))
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/Espresso/EspressoDataFrameTensorAttachment/copyFromCVPixelBuffer:
func (_EspressoDataFrameTensorAttachmentClass EspressoDataFrameTensorAttachmentClass) CopyFromCVPixelBuffer(buffer corevideo.CVImageBufferRef) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_EspressoDataFrameTensorAttachmentClass.class), objc.Sel("copyFromCVPixelBuffer:"), buffer)
	return objectivec.Object{ID: rv}
}

