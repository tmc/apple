// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVAssetVariantAudioRenditionSpecificAttributes] class.
var (
	_AVAssetVariantAudioRenditionSpecificAttributesClass     AVAssetVariantAudioRenditionSpecificAttributesClass
	_AVAssetVariantAudioRenditionSpecificAttributesClassOnce sync.Once
)

func getAVAssetVariantAudioRenditionSpecificAttributesClass() AVAssetVariantAudioRenditionSpecificAttributesClass {
	_AVAssetVariantAudioRenditionSpecificAttributesClassOnce.Do(func() {
		_AVAssetVariantAudioRenditionSpecificAttributesClass = AVAssetVariantAudioRenditionSpecificAttributesClass{class: objc.GetClass("AVAssetVariantAudioRenditionSpecificAttributes")}
	})
	return _AVAssetVariantAudioRenditionSpecificAttributesClass
}

// GetAVAssetVariantAudioRenditionSpecificAttributesClass returns the class object for AVAssetVariantAudioRenditionSpecificAttributes.
func GetAVAssetVariantAudioRenditionSpecificAttributesClass() AVAssetVariantAudioRenditionSpecificAttributesClass {
	return getAVAssetVariantAudioRenditionSpecificAttributesClass()
}

type AVAssetVariantAudioRenditionSpecificAttributesClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVAssetVariantAudioRenditionSpecificAttributesClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVAssetVariantAudioRenditionSpecificAttributesClass) Alloc() AVAssetVariantAudioRenditionSpecificAttributes {
	rv := objc.Send[AVAssetVariantAudioRenditionSpecificAttributes](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that represents attributes specific to a particular rendition.
//
// # Accessing attributes
//
//   - [AVAssetVariantAudioRenditionSpecificAttributes.Binaural]: A Boolean value that indicates the variant is best suited for delivery to headphones.
//   - [AVAssetVariantAudioRenditionSpecificAttributes.Immersive]: A Boolean value that indicates whether this variant contains virtualized or otherwise preprocessed audio content suitable for various purposes.
//   - [AVAssetVariantAudioRenditionSpecificAttributes.Downmix]: A Boolean value that indicates whether the variant is a downmix derivative of other media of greater channel count.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetVariant/AudioAttributes-swift.class/RenditionSpecificAttributes
type AVAssetVariantAudioRenditionSpecificAttributes struct {
	objectivec.Object
}

// AVAssetVariantAudioRenditionSpecificAttributesFromID constructs a [AVAssetVariantAudioRenditionSpecificAttributes] from an objc.ID.
//
// An object that represents attributes specific to a particular rendition.
func AVAssetVariantAudioRenditionSpecificAttributesFromID(id objc.ID) AVAssetVariantAudioRenditionSpecificAttributes {
	return AVAssetVariantAudioRenditionSpecificAttributes{objectivec.Object{ID: id}}
}

// NOTE: AVAssetVariantAudioRenditionSpecificAttributes adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVAssetVariantAudioRenditionSpecificAttributes] class.
//
// # Accessing attributes
//
//   - [IAVAssetVariantAudioRenditionSpecificAttributes.Binaural]: A Boolean value that indicates the variant is best suited for delivery to headphones.
//   - [IAVAssetVariantAudioRenditionSpecificAttributes.Immersive]: A Boolean value that indicates whether this variant contains virtualized or otherwise preprocessed audio content suitable for various purposes.
//   - [IAVAssetVariantAudioRenditionSpecificAttributes.Downmix]: A Boolean value that indicates whether the variant is a downmix derivative of other media of greater channel count.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetVariant/AudioAttributes-swift.class/RenditionSpecificAttributes
type IAVAssetVariantAudioRenditionSpecificAttributes interface {
	objectivec.IObject

	// Topic: Accessing attributes

	// A Boolean value that indicates the variant is best suited for delivery to headphones.
	Binaural() bool
	// A Boolean value that indicates whether this variant contains virtualized or otherwise preprocessed audio content suitable for various purposes.
	Immersive() bool
	// A Boolean value that indicates whether the variant is a downmix derivative of other media of greater channel count.
	Downmix() bool

	// The count of audio channels in the rendition.
	ChannelCount() int
	// The audio formats of the renditions present in the variant.
	FormatIDs() objectivec.IObject
	SetFormatIDs(value objectivec.IObject)
}

// Init initializes the instance.
func (a AVAssetVariantAudioRenditionSpecificAttributes) Init() AVAssetVariantAudioRenditionSpecificAttributes {
	rv := objc.Send[AVAssetVariantAudioRenditionSpecificAttributes](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AVAssetVariantAudioRenditionSpecificAttributes) Autorelease() AVAssetVariantAudioRenditionSpecificAttributes {
	rv := objc.Send[AVAssetVariantAudioRenditionSpecificAttributes](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVAssetVariantAudioRenditionSpecificAttributes creates a new AVAssetVariantAudioRenditionSpecificAttributes instance.
func NewAVAssetVariantAudioRenditionSpecificAttributes() AVAssetVariantAudioRenditionSpecificAttributes {
	class := getAVAssetVariantAudioRenditionSpecificAttributesClass()
	rv := objc.Send[AVAssetVariantAudioRenditionSpecificAttributes](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// A Boolean value that indicates the variant is best suited for delivery to
// headphones.
//
// # Discussion
//
// A binaural variant may originate from a direct binaural recording or from
// the processing of a multichannel audio source.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetVariant/AudioAttributes-swift.class/RenditionSpecificAttributes/isBinaural
func (a AVAssetVariantAudioRenditionSpecificAttributes) Binaural() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("isBinaural"))
	return rv
}

// A Boolean value that indicates whether this variant contains virtualized or
// otherwise preprocessed audio content suitable for various purposes.
//
// # Discussion
//
// If a variant audio redition is immersive it’s eligible for rendering to
// headphones or speakers.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetVariant/AudioAttributes-swift.class/RenditionSpecificAttributes/isImmersive
func (a AVAssetVariantAudioRenditionSpecificAttributes) Immersive() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("isImmersive"))
	return rv
}

// A Boolean value that indicates whether the variant is a downmix derivative
// of other media of greater channel count.
//
// # Discussion
//
// If the stream provides one or more multichannel variants, the system
// assumes the dowmix rendition to be compatible in its internal timing and
// other attributes with the other variants. Typically, this is because the
// variant derives its content from the same source. You can use a downmix as
// a suitable substitute for a multichannel variant under some conditions.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetVariant/AudioAttributes-swift.class/RenditionSpecificAttributes/isDownmix
func (a AVAssetVariantAudioRenditionSpecificAttributes) Downmix() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("isDownmix"))
	return rv
}

// The count of audio channels in the rendition.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetVariantAudioRenditionSpecificAttributes/channelCount
func (a AVAssetVariantAudioRenditionSpecificAttributes) ChannelCount() int {
	rv := objc.Send[int](a.ID, objc.Sel("channelCount"))
	return rv
}

// The audio formats of the renditions present in the variant.
//
// See: https://developer.apple.com/documentation/avfoundation/avassetvariant/audioattributes-swift.class/formatids
func (a AVAssetVariantAudioRenditionSpecificAttributes) FormatIDs() objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("formatIDs"))
	return objectivec.Object{ID: rv}
}
func (a AVAssetVariantAudioRenditionSpecificAttributes) SetFormatIDs(value objectivec.IObject) {
	objc.Send[struct{}](a.ID, objc.Sel("setFormatIDs:"), value)
}
