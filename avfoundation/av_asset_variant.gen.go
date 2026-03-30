// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVAssetVariant] class.
var (
	_AVAssetVariantClass     AVAssetVariantClass
	_AVAssetVariantClassOnce sync.Once
)

func getAVAssetVariantClass() AVAssetVariantClass {
	_AVAssetVariantClassOnce.Do(func() {
		_AVAssetVariantClass = AVAssetVariantClass{class: objc.GetClass("AVAssetVariant")}
	})
	return _AVAssetVariantClass
}

// GetAVAssetVariantClass returns the class object for AVAssetVariant.
func GetAVAssetVariantClass() AVAssetVariantClass {
	return getAVAssetVariantClass()
}

type AVAssetVariantClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVAssetVariantClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVAssetVariantClass) Alloc() AVAssetVariant {
	rv := objc.Send[AVAssetVariant](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that represents a bit rate variant.
//
// # Configuring attributes
//
//   - [AVAssetVariant.AudioAttributes]: The audio rendition attributes for the variant.
//   - [AVAssetVariant.VideoAttributes]: The video rendition attributes for the variant.
//
// # Accessing the URL
//
//   - [AVAssetVariant.URL]: Provides URL to media playlist corresponding to variant
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetVariant
type AVAssetVariant struct {
	objectivec.Object
}

// AVAssetVariantFromID constructs a [AVAssetVariant] from an objc.ID.
//
// An object that represents a bit rate variant.
func AVAssetVariantFromID(id objc.ID) AVAssetVariant {
	return AVAssetVariant{objectivec.Object{ID: id}}
}

// NOTE: AVAssetVariant adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVAssetVariant] class.
//
// # Configuring attributes
//
//   - [IAVAssetVariant.AudioAttributes]: The audio rendition attributes for the variant.
//   - [IAVAssetVariant.VideoAttributes]: The video rendition attributes for the variant.
//
// # Accessing the URL
//
//   - [IAVAssetVariant.URL]: Provides URL to media playlist corresponding to variant
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetVariant
type IAVAssetVariant interface {
	objectivec.IObject

	// Topic: Configuring attributes

	// The audio rendition attributes for the variant.
	AudioAttributes() IAVAssetVariantAudioAttributes
	// The video rendition attributes for the variant.
	VideoAttributes() IAVAssetVariantVideoAttributes

	// Topic: Accessing the URL

	// Provides URL to media playlist corresponding to variant
	URL() foundation.INSURL

	// The average bit rate for the variant.
	AverageBitRate() float64
	// The peak bit rate for the variant.
	PeakBitRate() float64
}

// Init initializes the instance.
func (a AVAssetVariant) Init() AVAssetVariant {
	rv := objc.Send[AVAssetVariant](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AVAssetVariant) Autorelease() AVAssetVariant {
	rv := objc.Send[AVAssetVariant](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVAssetVariant creates a new AVAssetVariant instance.
func NewAVAssetVariant() AVAssetVariant {
	class := getAVAssetVariantClass()
	rv := objc.Send[AVAssetVariant](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The audio rendition attributes for the variant.
//
// # Discussion
//
// This property value is `nil` if the variant defines no audio attributes.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetVariant/audioAttributes-swift.property
func (a AVAssetVariant) AudioAttributes() IAVAssetVariantAudioAttributes {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("audioAttributes"))
	return AVAssetVariantAudioAttributesFromID(objc.ID(rv))
}

// The video rendition attributes for the variant.
//
// # Discussion
//
// This property value is `nil` if the variant defines no video attributes.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetVariant/videoAttributes-swift.property
func (a AVAssetVariant) VideoAttributes() IAVAssetVariantVideoAttributes {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("videoAttributes"))
	return AVAssetVariantVideoAttributesFromID(objc.ID(rv))
}

// Provides URL to media playlist corresponding to variant
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetVariant/url
func (a AVAssetVariant) URL() foundation.INSURL {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("URL"))
	return foundation.NSURLFromID(objc.ID(rv))
}

// The average bit rate for the variant.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetVariant/averageBitRate-7bnsq
func (a AVAssetVariant) AverageBitRate() float64 {
	rv := objc.Send[float64](a.ID, objc.Sel("averageBitRate"))
	return rv
}

// The peak bit rate for the variant.
//
// # Discussion
//
// If the variant doesn’t define a peak bit rate, the value is negative.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetVariant/peakBitRate-38p2b
func (a AVAssetVariant) PeakBitRate() float64 {
	rv := objc.Send[float64](a.ID, objc.Sel("peakBitRate"))
	return rv
}
