// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corevideo"
)

// The class instance for the [EspressoDataFrameImageAttachment] class.
var (
	_EspressoDataFrameImageAttachmentClass     EspressoDataFrameImageAttachmentClass
	_EspressoDataFrameImageAttachmentClassOnce sync.Once
)

func getEspressoDataFrameImageAttachmentClass() EspressoDataFrameImageAttachmentClass {
	_EspressoDataFrameImageAttachmentClassOnce.Do(func() {
		_EspressoDataFrameImageAttachmentClass = EspressoDataFrameImageAttachmentClass{class: objc.GetClass("EspressoDataFrameImageAttachment")}
	})
	return _EspressoDataFrameImageAttachmentClass
}

// GetEspressoDataFrameImageAttachmentClass returns the class object for EspressoDataFrameImageAttachment.
func GetEspressoDataFrameImageAttachmentClass() EspressoDataFrameImageAttachmentClass {
	return getEspressoDataFrameImageAttachmentClass()
}

type EspressoDataFrameImageAttachmentClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (ec EspressoDataFrameImageAttachmentClass) Alloc() EspressoDataFrameImageAttachment {
	rv := objc.Send[EspressoDataFrameImageAttachment](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [EspressoDataFrameImageAttachment.CopyAsImageGrayscaleOrBGRA]
//   - [EspressoDataFrameImageAttachment.NChannels]
//   - [EspressoDataFrameImageAttachment.SetNChannels]
// See: https://developer.apple.com/documentation/Espresso/EspressoDataFrameImageAttachment
type EspressoDataFrameImageAttachment struct {
	EspressoDataFrameAttachment
}

// EspressoDataFrameImageAttachmentFromID constructs a [EspressoDataFrameImageAttachment] from an objc.ID.
func EspressoDataFrameImageAttachmentFromID(id objc.ID) EspressoDataFrameImageAttachment {
	return EspressoDataFrameImageAttachment{EspressoDataFrameAttachment: EspressoDataFrameAttachmentFromID(id)}
}
// Ensure EspressoDataFrameImageAttachment implements IEspressoDataFrameImageAttachment.
var _ IEspressoDataFrameImageAttachment = EspressoDataFrameImageAttachment{}

// An interface definition for the [EspressoDataFrameImageAttachment] class.
//
// # Methods
//
//   - [IEspressoDataFrameImageAttachment.CopyAsImageGrayscaleOrBGRA]
//   - [IEspressoDataFrameImageAttachment.NChannels]
//   - [IEspressoDataFrameImageAttachment.SetNChannels]
//
// See: https://developer.apple.com/documentation/Espresso/EspressoDataFrameImageAttachment
type IEspressoDataFrameImageAttachment interface {
	IEspressoDataFrameAttachment

	// Topic: Methods

	CopyAsImageGrayscaleOrBGRA() unsafe.Pointer
	NChannels() int
	SetNChannels(value int)
}

// Init initializes the instance.
func (e EspressoDataFrameImageAttachment) Init() EspressoDataFrameImageAttachment {
	rv := objc.Send[EspressoDataFrameImageAttachment](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e EspressoDataFrameImageAttachment) Autorelease() EspressoDataFrameImageAttachment {
	rv := objc.Send[EspressoDataFrameImageAttachment](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewEspressoDataFrameImageAttachment creates a new EspressoDataFrameImageAttachment instance.
func NewEspressoDataFrameImageAttachment() EspressoDataFrameImageAttachment {
	class := getEspressoDataFrameImageAttachmentClass()
	rv := objc.Send[EspressoDataFrameImageAttachment](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Espresso/EspressoDataFrameImageAttachment/copyAsImageGrayscaleOrBGRA
func (e EspressoDataFrameImageAttachment) CopyAsImageGrayscaleOrBGRA() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e.ID, objc.Sel("copyAsImageGrayscaleOrBGRA"))
	return rv
}

//
// See: https://developer.apple.com/documentation/Espresso/EspressoDataFrameImageAttachment/createCVPixelBufferFromvImage:withPixelFormat:
func (_EspressoDataFrameImageAttachmentClass EspressoDataFrameImageAttachmentClass) CreateCVPixelBufferFromvImageWithPixelFormat(image unsafe.Pointer, format uint32) corevideo.CVImageBufferRef {
	rv := objc.Send[corevideo.CVImageBufferRef](objc.ID(_EspressoDataFrameImageAttachmentClass.class), objc.Sel("createCVPixelBufferFromvImage:withPixelFormat:"), image, format)
	return corevideo.CVImageBufferRef(rv)
}

// See: https://developer.apple.com/documentation/Espresso/EspressoDataFrameImageAttachment/nChannels
func (e EspressoDataFrameImageAttachment) NChannels() int {
	rv := objc.Send[int](e.ID, objc.Sel("nChannels"))
	return rv
}
func (e EspressoDataFrameImageAttachment) SetNChannels(value int) {
	objc.Send[struct{}](e.ID, objc.Sel("setNChannels:"), value)
}

