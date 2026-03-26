// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVAssetVariantQualifier] class.
var (
	_AVAssetVariantQualifierClass     AVAssetVariantQualifierClass
	_AVAssetVariantQualifierClassOnce sync.Once
)

func getAVAssetVariantQualifierClass() AVAssetVariantQualifierClass {
	_AVAssetVariantQualifierClassOnce.Do(func() {
		_AVAssetVariantQualifierClass = AVAssetVariantQualifierClass{class: objc.GetClass("AVAssetVariantQualifier")}
	})
	return _AVAssetVariantQualifierClass
}

// GetAVAssetVariantQualifierClass returns the class object for AVAssetVariantQualifier.
func GetAVAssetVariantQualifierClass() AVAssetVariantQualifierClass {
	return getAVAssetVariantQualifierClass()
}

type AVAssetVariantQualifierClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVAssetVariantQualifierClass) Alloc() AVAssetVariantQualifier {
	rv := objc.Send[AVAssetVariantQualifier](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that represents an HTTP Live Streaming asset variant.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetVariantQualifier
type AVAssetVariantQualifier struct {
	objectivec.Object
}

// AVAssetVariantQualifierFromID constructs a [AVAssetVariantQualifier] from an objc.ID.
//
// An object that represents an HTTP Live Streaming asset variant.
func AVAssetVariantQualifierFromID(id objc.ID) AVAssetVariantQualifier {
	return AVAssetVariantQualifier{objectivec.Object{ID: id}}
}
// NOTE: AVAssetVariantQualifier adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVAssetVariantQualifier] class.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetVariantQualifier
type IAVAssetVariantQualifier interface {
	objectivec.IObject

	// The media selections of an asset that a task downloads.
	MediaSelections() IAVMediaSelection
	SetMediaSelections(value IAVMediaSelection)
	// The variant qualifiers for this configuration.
	VariantQualifiers() IAVAssetVariantQualifier
	SetVariantQualifiers(value IAVAssetVariantQualifier)
}

// Init initializes the instance.
func (a AVAssetVariantQualifier) Init() AVAssetVariantQualifier {
	rv := objc.Send[AVAssetVariantQualifier](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AVAssetVariantQualifier) Autorelease() AVAssetVariantQualifier {
	rv := objc.Send[AVAssetVariantQualifier](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVAssetVariantQualifier creates a new AVAssetVariantQualifier instance.
func NewAVAssetVariantQualifier() AVAssetVariantQualifier {
	class := getAVAssetVariantQualifierClass()
	rv := objc.Send[AVAssetVariantQualifier](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a variant qualifier with a predicate.
//
// predicate: A predicate to find a particular asset variant.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetVariantQualifier/init(predicate:)
func NewAssetVariantQualifierWithPredicate(predicate foundation.INSPredicate) AVAssetVariantQualifier {
	rv := objc.Send[objc.ID](objc.ID(getAVAssetVariantQualifierClass().class), objc.Sel("assetVariantQualifierWithPredicate:"), predicate)
	return AVAssetVariantQualifierFromID(rv)
}

// Creates a variant qualifier with an asset variant.
//
// variant: The asset variant for the qualifier.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetVariantQualifier/init(variant:)
func NewAssetVariantQualifierWithVariant(variant IAVAssetVariant) AVAssetVariantQualifier {
	rv := objc.Send[objc.ID](objc.ID(getAVAssetVariantQualifierClass().class), objc.Sel("assetVariantQualifierWithVariant:"), variant)
	return AVAssetVariantQualifierFromID(rv)
}

// Creates a predicate for audio sample rate.
//
// sampleRate: The specified sample rate to match.
//
// mediaSelectionOption: The audio media selection option under consideration.
//
// operatorType: The operator type.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetVariantQualifier/predicate(forAudioSampleRate:mediaSelectionOption:operatorType:)
func (_AVAssetVariantQualifierClass AVAssetVariantQualifierClass) PredicateForAudioSampleRateMediaSelectionOptionOperatorType(sampleRate float64, mediaSelectionOption IAVMediaSelectionOption, operatorType foundation.NSPredicateOperatorType) foundation.NSPredicate {
	rv := objc.Send[objc.ID](objc.ID(_AVAssetVariantQualifierClass.class), objc.Sel("predicateForAudioSampleRate:mediaSelectionOption:operatorType:"), sampleRate, mediaSelectionOption, operatorType)
	return foundation.NSPredicateFromID(rv)
}
// Creates a NSPredicate for audio sample rate which can be used with other
// NSPredicates to express variant preferences.
//
// sampleRate: The RHS value for the sample rate in the predicate equation.
//
// operatorType: The valid values are NSLessThanPredicateOperatorType,
// NSLessThanOrEqualToPredicateOperatorType,
// NSGreaterThanPredicateOperatorType,
// NSGreaterThanOrEqualToPredicateOperatorType, NSEqualToPredicateOperatorType
// and NSNotEqualToPredicateOperatorType.
//
// # Discussion
// 
// Predicate will be evaluated on the media selection option selected for the
// asset. Media selection options for primary assets may be specified in the
// AVAssetDownloadConfiguration mediaSelections property. Media selection
// options for interstitial assets may be circumscribed by
// -[AVAssetDownloadConfiguration setInterstitialMediaSelectionCriteria:
// forMediaCharacteristic:].
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetVariantQualifier/predicate(forAudioSampleRate:operatorType:)
func (_AVAssetVariantQualifierClass AVAssetVariantQualifierClass) PredicateForAudioSampleRateOperatorType(sampleRate float64, operatorType foundation.NSPredicateOperatorType) foundation.NSPredicate {
	rv := objc.Send[objc.ID](objc.ID(_AVAssetVariantQualifierClass.class), objc.Sel("predicateForAudioSampleRate:operatorType:"), sampleRate, operatorType)
	return foundation.NSPredicateFromID(rv)
}
// Creates a NSPredicate for binaural which can be used with other
// NSPredicates to express variant preferences.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetVariantQualifier/predicate(forBinauralAudio:)
func (_AVAssetVariantQualifierClass AVAssetVariantQualifierClass) PredicateForBinauralAudio(isBinauralAudio bool) foundation.NSPredicate {
	rv := objc.Send[objc.ID](objc.ID(_AVAssetVariantQualifierClass.class), objc.Sel("predicateForBinauralAudio:"), isBinauralAudio)
	return foundation.NSPredicateFromID(rv)
}
// Creates a predicate for binaural audio.
//
// isBinauralAudio: A Boolean value that indicates the binaural state to use in the predicate
// equation.
//
// mediaSelectionOption: The media selection option for the variant.
//
// # Return Value
// 
// A predicate object that you use to to create an [AVAssetVariantQualifier].
//
// # Discussion
// 
// Use the returned value, along with other predicates, to express variant
// preferences.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetVariantQualifier/predicate(forBinauralAudio:mediaSelectionOption:)
func (_AVAssetVariantQualifierClass AVAssetVariantQualifierClass) PredicateForBinauralAudioMediaSelectionOption(isBinauralAudio bool, mediaSelectionOption IAVMediaSelectionOption) foundation.NSPredicate {
	rv := objc.Send[objc.ID](objc.ID(_AVAssetVariantQualifierClass.class), objc.Sel("predicateForBinauralAudio:mediaSelectionOption:"), isBinauralAudio, mediaSelectionOption)
	return foundation.NSPredicateFromID(rv)
}
// Creates a predicate with a channel count, media selection option, and
// operator type.
//
// channelCount: The number of channels in the variant.
//
// mediaSelectionOption: The media selection option for the variant.
//
// operatorType: The predicate operator.
//
// # Return Value
// 
// A predicate object that you use to to create an [AVAssetVariantQualifier].
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetVariantQualifier/predicate(forChannelCount:mediaSelectionOption:operatorType:)
func (_AVAssetVariantQualifierClass AVAssetVariantQualifierClass) PredicateForChannelCountMediaSelectionOptionOperatorType(channelCount int, mediaSelectionOption IAVMediaSelectionOption, operatorType foundation.NSPredicateOperatorType) foundation.NSPredicate {
	rv := objc.Send[objc.ID](objc.ID(_AVAssetVariantQualifierClass.class), objc.Sel("predicateForChannelCount:mediaSelectionOption:operatorType:"), channelCount, mediaSelectionOption, operatorType)
	return foundation.NSPredicateFromID(rv)
}
// Creates a NSPredicate for audio channel count which can be used with other
// NSPredicates to express variant preferences.
//
// channelCount: The RHS value for the channel count in the predicate equation.
//
// operatorType: The valid values are NSLessThanPredicateOperatorType,
// NSLessThanOrEqualToPredicateOperatorType,
// NSGreaterThanPredicateOperatorType,
// NSGreaterThanOrEqualToPredicateOperatorType, NSEqualToPredicateOperatorType
// and NSNotEqualToPredicateOperatorType.
//
// # Discussion
// 
// Predicate will be evaluated on the media selection option selected for the
// asset. Media selection options for primary assets may be specified in the
// AVAssetDownloadConfiguration mediaSelections property. Media selection
// options for interstitial assets may be circumscribed by
// -[AVAssetDownloadConfiguration setInterstitialMediaSelectionCriteria:
// forMediaCharacteristic:].
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetVariantQualifier/predicate(forChannelCount:operatorType:)
func (_AVAssetVariantQualifierClass AVAssetVariantQualifierClass) PredicateForChannelCountOperatorType(channelCount int, operatorType foundation.NSPredicateOperatorType) foundation.NSPredicate {
	rv := objc.Send[objc.ID](objc.ID(_AVAssetVariantQualifierClass.class), objc.Sel("predicateForChannelCount:operatorType:"), channelCount, operatorType)
	return foundation.NSPredicateFromID(rv)
}
// Creates a NSPredicate for immersive audio which can be used with other
// NSPredicates to express variant preferences.
//
// isDownmixAudio: The RHS value for the value of isDownmixAudio in the predicate equation.
//
// # Discussion
// 
// Predicate will be evaluated on the media selection option selected for the
// asset. Media selection options for primary assets may be specified in the
// AVAssetDownloadConfiguration mediaSelections property. Media selection
// options for interstitial assets may be circumscribed by
// -[AVAssetDownloadConfiguration setInterstitialMediaSelectionCriteria:
// forMediaCharacteristic:].
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetVariantQualifier/predicate(forDownmixAudio:)
func (_AVAssetVariantQualifierClass AVAssetVariantQualifierClass) PredicateForDownmixAudio(isDownmixAudio bool) foundation.NSPredicate {
	rv := objc.Send[objc.ID](objc.ID(_AVAssetVariantQualifierClass.class), objc.Sel("predicateForDownmixAudio:"), isDownmixAudio)
	return foundation.NSPredicateFromID(rv)
}
// Creates a predicate for downmix audio.
//
// mediaSelectionOption: The media selection option for the variant.
//
// # Return Value
// 
// A predicate object that you use to to create an [AVAssetVariantQualifier].
//
// # Discussion
// 
// Use the returned value, along with other predicates, to express variant
// preferences.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetVariantQualifier/predicate(forDownmixAudio:mediaSelectionOption:)
func (_AVAssetVariantQualifierClass AVAssetVariantQualifierClass) PredicateForDownmixAudioMediaSelectionOption(isDownmixAudio bool, mediaSelectionOption IAVMediaSelectionOption) foundation.NSPredicate {
	rv := objc.Send[objc.ID](objc.ID(_AVAssetVariantQualifierClass.class), objc.Sel("predicateForDownmixAudio:mediaSelectionOption:"), isDownmixAudio, mediaSelectionOption)
	return foundation.NSPredicateFromID(rv)
}
// Creates a NSPredicate for immersive audio which can be used with other
// NSPredicates to express variant preferences.
//
// isImmersiveAudio: The RHS value for the value of isImmersiveAudio in the predicate equation.
//
// # Discussion
// 
// Predicate will be evaluated on the media selection option selected for the
// asset. Media selection options for primary assets may be specified in the
// AVAssetDownloadConfiguration mediaSelections property. Media selection
// options for interstitial assets may be circumscribed by
// -[AVAssetDownloadConfiguration setInterstitialMediaSelectionCriteria:
// forMediaCharacteristic:].
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetVariantQualifier/predicate(forImmersiveAudio:)
func (_AVAssetVariantQualifierClass AVAssetVariantQualifierClass) PredicateForImmersiveAudio(isImmersiveAudio bool) foundation.NSPredicate {
	rv := objc.Send[objc.ID](objc.ID(_AVAssetVariantQualifierClass.class), objc.Sel("predicateForImmersiveAudio:"), isImmersiveAudio)
	return foundation.NSPredicateFromID(rv)
}
// Creates a predicate for immersive audio.
//
// mediaSelectionOption: The media selection option for the variant.
//
// # Return Value
// 
// A predicate object that you use to to create an [AVAssetVariantQualifier].
//
// # Discussion
// 
// Use the returned value, along with other predicates, to express variant
// preferences.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetVariantQualifier/predicate(forImmersiveAudio:mediaSelectionOption:)
func (_AVAssetVariantQualifierClass AVAssetVariantQualifierClass) PredicateForImmersiveAudioMediaSelectionOption(isImmersiveAudio bool, mediaSelectionOption IAVMediaSelectionOption) foundation.NSPredicate {
	rv := objc.Send[objc.ID](objc.ID(_AVAssetVariantQualifierClass.class), objc.Sel("predicateForImmersiveAudio:mediaSelectionOption:"), isImmersiveAudio, mediaSelectionOption)
	return foundation.NSPredicateFromID(rv)
}
// Creates a predicate with a height and operator type.
//
// height: The presentation height.
//
// operatorType: The predicate’s operator type.
//
// # Return Value
// 
// A predicate object that you use to to create an [AVAssetVariantQualifier].
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetVariantQualifier/predicate(forPresentationHeight:operatorType:)
func (_AVAssetVariantQualifierClass AVAssetVariantQualifierClass) PredicateForPresentationHeightOperatorType(height float64, operatorType foundation.NSPredicateOperatorType) foundation.NSPredicate {
	rv := objc.Send[objc.ID](objc.ID(_AVAssetVariantQualifierClass.class), objc.Sel("predicateForPresentationHeight:operatorType:"), height, operatorType)
	return foundation.NSPredicateFromID(rv)
}
// Creates a predicate with a width and operator type.
//
// width: The presentation width.
//
// operatorType: The predicate’s operator type.
//
// # Return Value
// 
// A predicate object that you use to to create an [AVAssetVariantQualifier].
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetVariantQualifier/predicate(forPresentationWidth:operatorType:)
func (_AVAssetVariantQualifierClass AVAssetVariantQualifierClass) PredicateForPresentationWidthOperatorType(width float64, operatorType foundation.NSPredicateOperatorType) foundation.NSPredicate {
	rv := objc.Send[objc.ID](objc.ID(_AVAssetVariantQualifierClass.class), objc.Sel("predicateForPresentationWidth:operatorType:"), width, operatorType)
	return foundation.NSPredicateFromID(rv)
}
// Returns a qualifer for finding variant with maximum value in the input key
// path
//
// keyPath: AVAssetVariant keyPath. Allowed keyPath values are peakBitRate,
// averageBitRate, videoAttributes.presentationSize. Must be a valid, non-nil
// NSString.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetVariantQualifier/assetVariantQualifierForMaximumValueInKeyPath:
func (_AVAssetVariantQualifierClass AVAssetVariantQualifierClass) AssetVariantQualifierForMaximumValueInKeyPath(keyPath string) AVAssetVariantQualifier {
	rv := objc.Send[objc.ID](objc.ID(_AVAssetVariantQualifierClass.class), objc.Sel("assetVariantQualifierForMaximumValueInKeyPath:"), objc.String(keyPath))
	return AVAssetVariantQualifierFromID(rv)
}
// Returns a qualifer for finding variant with minimum value in the input key
// path.
//
// keyPath: AVAssetVariant keyPath. Allowed keyPath values are peakBitRate,
// averageBitRate, videoAttributes.presentationSize. Must be a valid, non-nil
// NSString.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetVariantQualifier/assetVariantQualifierForMinimumValueInKeyPath:
func (_AVAssetVariantQualifierClass AVAssetVariantQualifierClass) AssetVariantQualifierForMinimumValueInKeyPath(keyPath string) AVAssetVariantQualifier {
	rv := objc.Send[objc.ID](objc.ID(_AVAssetVariantQualifierClass.class), objc.Sel("assetVariantQualifierForMinimumValueInKeyPath:"), objc.String(keyPath))
	return AVAssetVariantQualifierFromID(rv)
}

// The media selections of an asset that a task downloads.
//
// See: https://developer.apple.com/documentation/avfoundation/avassetdownloadcontentconfiguration/mediaselections
func (a AVAssetVariantQualifier) MediaSelections() IAVMediaSelection {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("mediaSelections"))
	return AVMediaSelectionFromID(objc.ID(rv))
}
func (a AVAssetVariantQualifier) SetMediaSelections(value IAVMediaSelection) {
	objc.Send[struct{}](a.ID, objc.Sel("setMediaSelections:"), value)
}
// The variant qualifiers for this configuration.
//
// See: https://developer.apple.com/documentation/avfoundation/avassetdownloadcontentconfiguration/variantqualifiers
func (a AVAssetVariantQualifier) VariantQualifiers() IAVAssetVariantQualifier {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("variantQualifiers"))
	return AVAssetVariantQualifierFromID(objc.ID(rv))
}
func (a AVAssetVariantQualifier) SetVariantQualifiers(value IAVAssetVariantQualifier) {
	objc.Send[struct{}](a.ID, objc.Sel("setVariantQualifiers:"), value)
}

