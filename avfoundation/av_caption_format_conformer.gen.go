// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVCaptionFormatConformer] class.
var (
	_AVCaptionFormatConformerClass     AVCaptionFormatConformerClass
	_AVCaptionFormatConformerClassOnce sync.Once
)

func getAVCaptionFormatConformerClass() AVCaptionFormatConformerClass {
	_AVCaptionFormatConformerClassOnce.Do(func() {
		_AVCaptionFormatConformerClass = AVCaptionFormatConformerClass{class: objc.GetClass("AVCaptionFormatConformer")}
	})
	return _AVCaptionFormatConformerClass
}

// GetAVCaptionFormatConformerClass returns the class object for AVCaptionFormatConformer.
func GetAVCaptionFormatConformerClass() AVCaptionFormatConformerClass {
	return getAVCaptionFormatConformerClass()
}

type AVCaptionFormatConformerClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVCaptionFormatConformerClass) Alloc() AVCaptionFormatConformer {
	rv := objc.Send[AVCaptionFormatConformer](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that converts a canonical caption to a specific format.
//
// # Creating a format conformer
//
//   - [AVCaptionFormatConformer.InitWithConversionSettings]: Creates a new object with format conversion settings.
//
// # Conforming captions
//
//   - [AVCaptionFormatConformer.ConformsCaptionsToTimeRange]: A Boolean value that indicates whether to conform the time range of a canonical caption.
//   - [AVCaptionFormatConformer.SetConformsCaptionsToTimeRange]
//   - [AVCaptionFormatConformer.ConformedCaptionForCaptionError]: Creates a caption that conforms to a specific format.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptionFormatConformer
type AVCaptionFormatConformer struct {
	objectivec.Object
}

// AVCaptionFormatConformerFromID constructs a [AVCaptionFormatConformer] from an objc.ID.
//
// An object that converts a canonical caption to a specific format.
func AVCaptionFormatConformerFromID(id objc.ID) AVCaptionFormatConformer {
	return AVCaptionFormatConformer{objectivec.Object{ID: id}}
}
// NOTE: AVCaptionFormatConformer adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVCaptionFormatConformer] class.
//
// # Creating a format conformer
//
//   - [IAVCaptionFormatConformer.InitWithConversionSettings]: Creates a new object with format conversion settings.
//
// # Conforming captions
//
//   - [IAVCaptionFormatConformer.ConformsCaptionsToTimeRange]: A Boolean value that indicates whether to conform the time range of a canonical caption.
//   - [IAVCaptionFormatConformer.SetConformsCaptionsToTimeRange]
//   - [IAVCaptionFormatConformer.ConformedCaptionForCaptionError]: Creates a caption that conforms to a specific format.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptionFormatConformer
type IAVCaptionFormatConformer interface {
	objectivec.IObject

	// Topic: Creating a format conformer

	// Creates a new object with format conversion settings.
	InitWithConversionSettings(conversionSettings foundation.INSDictionary) AVCaptionFormatConformer

	// Topic: Conforming captions

	// A Boolean value that indicates whether to conform the time range of a canonical caption.
	ConformsCaptionsToTimeRange() bool
	SetConformsCaptionsToTimeRange(value bool)
	// Creates a caption that conforms to a specific format.
	ConformedCaptionForCaptionError(caption IAVCaption) (IAVCaption, error)
}

// Init initializes the instance.
func (c AVCaptionFormatConformer) Init() AVCaptionFormatConformer {
	rv := objc.Send[AVCaptionFormatConformer](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c AVCaptionFormatConformer) Autorelease() AVCaptionFormatConformer {
	rv := objc.Send[AVCaptionFormatConformer](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVCaptionFormatConformer creates a new AVCaptionFormatConformer instance.
func NewAVCaptionFormatConformer() AVCaptionFormatConformer {
	class := getAVCaptionFormatConformerClass()
	rv := objc.Send[AVCaptionFormatConformer](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a new object with format conversion settings.
//
// conversionSettings: A dictionary that specifies the conversion settings that this instance
// uses.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptionFormatConformer/init(conversionSettings:)
func NewCaptionFormatConformerWithConversionSettings(conversionSettings foundation.INSDictionary) AVCaptionFormatConformer {
	instance := getAVCaptionFormatConformerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithConversionSettings:"), conversionSettings)
	return AVCaptionFormatConformerFromID(rv)
}

// Creates a new object with format conversion settings.
//
// conversionSettings: A dictionary that specifies the conversion settings that this instance
// uses.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptionFormatConformer/init(conversionSettings:)
func (c AVCaptionFormatConformer) InitWithConversionSettings(conversionSettings foundation.INSDictionary) AVCaptionFormatConformer {
	rv := objc.Send[AVCaptionFormatConformer](c.ID, objc.Sel("initWithConversionSettings:"), conversionSettings)
	return rv
}
// Creates a caption that conforms to a specific format.
//
// caption: The caption to conform.
//
// # Return Value
// 
// A caption that conforms to the defined caption format.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptionFormatConformer/conformedCaption(for:)
func (c AVCaptionFormatConformer) ConformedCaptionForCaptionError(caption IAVCaption) (IAVCaption, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](c.ID, objc.Sel("conformedCaptionForCaption:error:"), caption, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return AVCaption{}, foundation.NSErrorFrom(errorPtr)
	}
	return AVCaptionFromID(rv), nil

}

// A class method that creates a new object with format conversion settings.
//
// conversionSettings: A dictionary that specifies the conversion settings that this instance
// uses.
//
// # Return Value
// 
// A new instance of [AVCaptionFormatConformer].
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptionFormatConformer/captionFormatConformerWithConversionSettings:
func (_AVCaptionFormatConformerClass AVCaptionFormatConformerClass) CaptionFormatConformerWithConversionSettings(conversionSettings foundation.INSDictionary) AVCaptionFormatConformer {
	rv := objc.Send[objc.ID](objc.ID(_AVCaptionFormatConformerClass.class), objc.Sel("captionFormatConformerWithConversionSettings:"), conversionSettings)
	return AVCaptionFormatConformerFromID(rv)
}

// A Boolean value that indicates whether to conform the time range of a
// canonical caption.
//
// # Discussion
// 
// By default, this property value is [false]. If you set the value to [true],
// this object conforms the time range of captions to fit its encoded data.
// 
// When this object conforms captions to CAE608 format, it encodes them so
// that each CAE608 2-byte control code fits into one frame duration
// (1001/30000).
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptionFormatConformer/conformsCaptionsToTimeRange
func (c AVCaptionFormatConformer) ConformsCaptionsToTimeRange() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("conformsCaptionsToTimeRange"))
	return rv
}
func (c AVCaptionFormatConformer) SetConformsCaptionsToTimeRange(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setConformsCaptionsToTimeRange:"), value)
}

