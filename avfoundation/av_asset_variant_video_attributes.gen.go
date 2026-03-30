// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVAssetVariantVideoAttributes] class.
var (
	_AVAssetVariantVideoAttributesClass     AVAssetVariantVideoAttributesClass
	_AVAssetVariantVideoAttributesClassOnce sync.Once
)

func getAVAssetVariantVideoAttributesClass() AVAssetVariantVideoAttributesClass {
	_AVAssetVariantVideoAttributesClassOnce.Do(func() {
		_AVAssetVariantVideoAttributesClass = AVAssetVariantVideoAttributesClass{class: objc.GetClass("AVAssetVariantVideoAttributes")}
	})
	return _AVAssetVariantVideoAttributesClass
}

// GetAVAssetVariantVideoAttributesClass returns the class object for AVAssetVariantVideoAttributes.
func GetAVAssetVariantVideoAttributesClass() AVAssetVariantVideoAttributesClass {
	return getAVAssetVariantVideoAttributesClass()
}

type AVAssetVariantVideoAttributesClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVAssetVariantVideoAttributesClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVAssetVariantVideoAttributesClass) Alloc() AVAssetVariantVideoAttributes {
	rv := objc.Send[AVAssetVariantVideoAttributes](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that defines the video attributes for an asset variant.
//
// # Inspecting the attributes
//
//   - [AVAssetVariantVideoAttributes.PresentationSize]: The presentation size of the variant’s renditions.
//   - [AVAssetVariantVideoAttributes.VideoRange]: The video range of the variant.
//   - [AVAssetVariantVideoAttributes.VideoLayoutAttributes]: Attributes that describe the layout of the video content.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetVariant/VideoAttributes-swift.class
type AVAssetVariantVideoAttributes struct {
	objectivec.Object
}

// AVAssetVariantVideoAttributesFromID constructs a [AVAssetVariantVideoAttributes] from an objc.ID.
//
// An object that defines the video attributes for an asset variant.
func AVAssetVariantVideoAttributesFromID(id objc.ID) AVAssetVariantVideoAttributes {
	return AVAssetVariantVideoAttributes{objectivec.Object{ID: id}}
}

// NOTE: AVAssetVariantVideoAttributes adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVAssetVariantVideoAttributes] class.
//
// # Inspecting the attributes
//
//   - [IAVAssetVariantVideoAttributes.PresentationSize]: The presentation size of the variant’s renditions.
//   - [IAVAssetVariantVideoAttributes.VideoRange]: The video range of the variant.
//   - [IAVAssetVariantVideoAttributes.VideoLayoutAttributes]: Attributes that describe the layout of the video content.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetVariant/VideoAttributes-swift.class
type IAVAssetVariantVideoAttributes interface {
	objectivec.IObject

	// Topic: Inspecting the attributes

	// The presentation size of the variant’s renditions.
	PresentationSize() corefoundation.CGSize
	// The video range of the variant.
	VideoRange() AVVideoRange
	// Attributes that describe the layout of the video content.
	VideoLayoutAttributes() []AVAssetVariantVideoLayoutAttributes

	// The audio rendition attributes for the variant.
	AudioAttributes() IAVAssetVariantAudioAttributes
	SetAudioAttributes(value IAVAssetVariantAudioAttributes)
	// The video sample codec types present in the variant’s renditions.
	CodecTypes() []foundation.NSNumber
	// The nominal frame rate of the variant’s renditions.
	NominalFrameRate() float64
	// The video rendition attributes for the variant.
	VideoAttributes() IAVAssetVariantVideoAttributes
	SetVideoAttributes(value IAVAssetVariantVideoAttributes)
}

// Init initializes the instance.
func (a AVAssetVariantVideoAttributes) Init() AVAssetVariantVideoAttributes {
	rv := objc.Send[AVAssetVariantVideoAttributes](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AVAssetVariantVideoAttributes) Autorelease() AVAssetVariantVideoAttributes {
	rv := objc.Send[AVAssetVariantVideoAttributes](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVAssetVariantVideoAttributes creates a new AVAssetVariantVideoAttributes instance.
func NewAVAssetVariantVideoAttributes() AVAssetVariantVideoAttributes {
	class := getAVAssetVariantVideoAttributesClass()
	rv := objc.Send[AVAssetVariantVideoAttributes](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The presentation size of the variant’s renditions.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetVariant/VideoAttributes-swift.class/presentationSize
func (a AVAssetVariantVideoAttributes) PresentationSize() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](a.ID, objc.Sel("presentationSize"))
	return corefoundation.CGSize(rv)
}

// The video range of the variant.
//
// # Discussion
//
// The property defaults to [sdr].
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetVariant/VideoAttributes-swift.class/videoRange
//
// [sdr]: https://developer.apple.com/documentation/AVFoundation/AVVideoRange/sdr
func (a AVAssetVariantVideoAttributes) VideoRange() AVVideoRange {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("videoRange"))
	return AVVideoRange(foundation.NSStringFromID(rv).String())
}

// Attributes that describe the layout of the video content.
//
// # Discussion
//
// This property may contain more that one element if the variant contains a
// collection of differing video layout media attributes over time.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetVariant/VideoAttributes-swift.class/videoLayoutAttributes
func (a AVAssetVariantVideoAttributes) VideoLayoutAttributes() []AVAssetVariantVideoLayoutAttributes {
	rv := objc.Send[[]objc.ID](a.ID, objc.Sel("videoLayoutAttributes"))
	return objc.ConvertSlice(rv, func(id objc.ID) AVAssetVariantVideoLayoutAttributes {
		return AVAssetVariantVideoLayoutAttributesFromID(id)
	})
}

// The audio rendition attributes for the variant.
//
// See: https://developer.apple.com/documentation/avfoundation/avassetvariant/audioattributes-swift.property
func (a AVAssetVariantVideoAttributes) AudioAttributes() IAVAssetVariantAudioAttributes {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("audioAttributes"))
	return AVAssetVariantAudioAttributesFromID(objc.ID(rv))
}
func (a AVAssetVariantVideoAttributes) SetAudioAttributes(value IAVAssetVariantAudioAttributes) {
	objc.Send[struct{}](a.ID, objc.Sel("setAudioAttributes:"), value)
}

// The video sample codec types present in the variant’s renditions.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetVariantVideoAttributes/codecTypes
func (a AVAssetVariantVideoAttributes) CodecTypes() []foundation.NSNumber {
	rv := objc.Send[[]objc.ID](a.ID, objc.Sel("codecTypes"))
	return objc.ConvertSlice(rv, func(id objc.ID) foundation.NSNumber {
		return foundation.NSNumberFromID(id)
	})
}

// The nominal frame rate of the variant’s renditions.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetVariantVideoAttributes/nominalFrameRate
func (a AVAssetVariantVideoAttributes) NominalFrameRate() float64 {
	rv := objc.Send[float64](a.ID, objc.Sel("nominalFrameRate"))
	return rv
}

// The video rendition attributes for the variant.
//
// See: https://developer.apple.com/documentation/avfoundation/avassetvariant/videoattributes-swift.property
func (a AVAssetVariantVideoAttributes) VideoAttributes() IAVAssetVariantVideoAttributes {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("videoAttributes"))
	return AVAssetVariantVideoAttributesFromID(objc.ID(rv))
}
func (a AVAssetVariantVideoAttributes) SetVideoAttributes(value IAVAssetVariantVideoAttributes) {
	objc.Send[struct{}](a.ID, objc.Sel("setVideoAttributes:"), value)
}
