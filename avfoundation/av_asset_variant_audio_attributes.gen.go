// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVAssetVariantAudioAttributes] class.
var (
	_AVAssetVariantAudioAttributesClass     AVAssetVariantAudioAttributesClass
	_AVAssetVariantAudioAttributesClassOnce sync.Once
)

func getAVAssetVariantAudioAttributesClass() AVAssetVariantAudioAttributesClass {
	_AVAssetVariantAudioAttributesClassOnce.Do(func() {
		_AVAssetVariantAudioAttributesClass = AVAssetVariantAudioAttributesClass{class: objc.GetClass("AVAssetVariantAudioAttributes")}
	})
	return _AVAssetVariantAudioAttributesClass
}

// GetAVAssetVariantAudioAttributesClass returns the class object for AVAssetVariantAudioAttributes.
func GetAVAssetVariantAudioAttributesClass() AVAssetVariantAudioAttributesClass {
	return getAVAssetVariantAudioAttributesClass()
}

type AVAssetVariantAudioAttributesClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVAssetVariantAudioAttributesClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVAssetVariantAudioAttributesClass) Alloc() AVAssetVariantAudioAttributes {
	rv := objc.Send[AVAssetVariantAudioAttributes](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that defines the audio attributes for an asset variant.
//
// # Inspecting audio attributes
//
//   - [AVAssetVariantAudioAttributes.RenditionSpecificAttributesForMediaOption]: Returns specific attributes for the media option.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetVariant/AudioAttributes-swift.class
type AVAssetVariantAudioAttributes struct {
	objectivec.Object
}

// AVAssetVariantAudioAttributesFromID constructs a [AVAssetVariantAudioAttributes] from an objc.ID.
//
// An object that defines the audio attributes for an asset variant.
func AVAssetVariantAudioAttributesFromID(id objc.ID) AVAssetVariantAudioAttributes {
	return AVAssetVariantAudioAttributes{objectivec.Object{ID: id}}
}
// NOTE: AVAssetVariantAudioAttributes adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVAssetVariantAudioAttributes] class.
//
// # Inspecting audio attributes
//
//   - [IAVAssetVariantAudioAttributes.RenditionSpecificAttributesForMediaOption]: Returns specific attributes for the media option.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetVariant/AudioAttributes-swift.class
type IAVAssetVariantAudioAttributes interface {
	objectivec.IObject

	// Topic: Inspecting audio attributes

	// Returns specific attributes for the media option.
	RenditionSpecificAttributesForMediaOption(mediaSelectionOption IAVMediaSelectionOption) IAVAssetVariantAudioRenditionSpecificAttributes

	// The audio rendition attributes for the variant.
	AudioAttributes() IAVAssetVariantAudioAttributes
	SetAudioAttributes(value IAVAssetVariantAudioAttributes)
	// The audio formats of the renditions present in the variant.
	FormatIDs() []foundation.NSNumber
	// The video rendition attributes for the variant.
	VideoAttributes() IAVAssetVariantVideoAttributes
	SetVideoAttributes(value IAVAssetVariantVideoAttributes)
}

// Init initializes the instance.
func (a AVAssetVariantAudioAttributes) Init() AVAssetVariantAudioAttributes {
	rv := objc.Send[AVAssetVariantAudioAttributes](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AVAssetVariantAudioAttributes) Autorelease() AVAssetVariantAudioAttributes {
	rv := objc.Send[AVAssetVariantAudioAttributes](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVAssetVariantAudioAttributes creates a new AVAssetVariantAudioAttributes instance.
func NewAVAssetVariantAudioAttributes() AVAssetVariantAudioAttributes {
	class := getAVAssetVariantAudioAttributesClass()
	rv := objc.Send[AVAssetVariantAudioAttributes](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Returns specific attributes for the media option.
//
// mediaSelectionOption: The media option for which to retrieve attributes.
//
// # Return Value
// 
// Attributes for the rendition, or `nil` of none exist.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetVariant/AudioAttributes-swift.class/renditionSpecificAttributes(for:)
func (a AVAssetVariantAudioAttributes) RenditionSpecificAttributesForMediaOption(mediaSelectionOption IAVMediaSelectionOption) IAVAssetVariantAudioRenditionSpecificAttributes {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("renditionSpecificAttributesForMediaOption:"), mediaSelectionOption)
	return AVAssetVariantAudioRenditionSpecificAttributesFromID(rv)
}

// The audio rendition attributes for the variant.
//
// See: https://developer.apple.com/documentation/avfoundation/avassetvariant/audioattributes-swift.property
func (a AVAssetVariantAudioAttributes) AudioAttributes() IAVAssetVariantAudioAttributes {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("audioAttributes"))
	return AVAssetVariantAudioAttributesFromID(objc.ID(rv))
}
func (a AVAssetVariantAudioAttributes) SetAudioAttributes(value IAVAssetVariantAudioAttributes) {
	objc.Send[struct{}](a.ID, objc.Sel("setAudioAttributes:"), value)
}
// The audio formats of the renditions present in the variant.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetVariantAudioAttributes/formatIDs
func (a AVAssetVariantAudioAttributes) FormatIDs() []foundation.NSNumber {
	rv := objc.Send[[]objc.ID](a.ID, objc.Sel("formatIDs"))
	return objc.ConvertSlice(rv, func(id objc.ID) foundation.NSNumber {
		return foundation.NSNumberFromID(id)
	})
}
// The video rendition attributes for the variant.
//
// See: https://developer.apple.com/documentation/avfoundation/avassetvariant/videoattributes-swift.property
func (a AVAssetVariantAudioAttributes) VideoAttributes() IAVAssetVariantVideoAttributes {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("videoAttributes"))
	return AVAssetVariantVideoAttributesFromID(objc.ID(rv))
}
func (a AVAssetVariantAudioAttributes) SetVideoAttributes(value IAVAssetVariantVideoAttributes) {
	objc.Send[struct{}](a.ID, objc.Sel("setVideoAttributes:"), value)
}

