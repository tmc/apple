// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVPlayerMediaSelectionCriteria] class.
var (
	_AVPlayerMediaSelectionCriteriaClass     AVPlayerMediaSelectionCriteriaClass
	_AVPlayerMediaSelectionCriteriaClassOnce sync.Once
)

func getAVPlayerMediaSelectionCriteriaClass() AVPlayerMediaSelectionCriteriaClass {
	_AVPlayerMediaSelectionCriteriaClassOnce.Do(func() {
		_AVPlayerMediaSelectionCriteriaClass = AVPlayerMediaSelectionCriteriaClass{class: objc.GetClass("AVPlayerMediaSelectionCriteria")}
	})
	return _AVPlayerMediaSelectionCriteriaClass
}

// GetAVPlayerMediaSelectionCriteriaClass returns the class object for AVPlayerMediaSelectionCriteria.
func GetAVPlayerMediaSelectionCriteriaClass() AVPlayerMediaSelectionCriteriaClass {
	return getAVPlayerMediaSelectionCriteriaClass()
}

type AVPlayerMediaSelectionCriteriaClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVPlayerMediaSelectionCriteriaClass) Alloc() AVPlayerMediaSelectionCriteria {
	rv := objc.Send[AVPlayerMediaSelectionCriteria](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that specifies the preferred languages and media characteristics
// for a player.
//
// # Overview
// 
// An instance of this object represents the languages and media
// characteristics of assets that contain media selection options that a
// player attempts to select automatically when preparing and playing items.
// It lists the languages and media characteristics in their preferred order.
//
// # Creating media selection criteria
//
//   - [AVPlayerMediaSelectionCriteria.InitWithPreferredLanguagesPreferredMediaCharacteristics]: Creates media selection criteria with the preferred languages and media characteristics.
//   - [AVPlayerMediaSelectionCriteria.InitWithPrincipalMediaCharacteristicsPreferredLanguagesPreferredMediaCharacteristics]: Creates media selection criteria with the principal media characteristics, and preferred languages and media characteristics.
//
// # Retrieving selection criteria settings
//
//   - [AVPlayerMediaSelectionCriteria.PreferredLanguages]: An array of language identifiers in preferred order.
//   - [AVPlayerMediaSelectionCriteria.PreferredMediaCharacteristics]: An array of media characteristics in preferred order.
//   - [AVPlayerMediaSelectionCriteria.PrincipalMediaCharacteristics]: An array of media characteristics that are essential to select when choosing media with a particular characteristic.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerMediaSelectionCriteria
type AVPlayerMediaSelectionCriteria struct {
	objectivec.Object
}

// AVPlayerMediaSelectionCriteriaFromID constructs a [AVPlayerMediaSelectionCriteria] from an objc.ID.
//
// An object that specifies the preferred languages and media characteristics
// for a player.
func AVPlayerMediaSelectionCriteriaFromID(id objc.ID) AVPlayerMediaSelectionCriteria {
	return AVPlayerMediaSelectionCriteria{objectivec.Object{ID: id}}
}
// NOTE: AVPlayerMediaSelectionCriteria adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVPlayerMediaSelectionCriteria] class.
//
// # Creating media selection criteria
//
//   - [IAVPlayerMediaSelectionCriteria.InitWithPreferredLanguagesPreferredMediaCharacteristics]: Creates media selection criteria with the preferred languages and media characteristics.
//   - [IAVPlayerMediaSelectionCriteria.InitWithPrincipalMediaCharacteristicsPreferredLanguagesPreferredMediaCharacteristics]: Creates media selection criteria with the principal media characteristics, and preferred languages and media characteristics.
//
// # Retrieving selection criteria settings
//
//   - [IAVPlayerMediaSelectionCriteria.PreferredLanguages]: An array of language identifiers in preferred order.
//   - [IAVPlayerMediaSelectionCriteria.PreferredMediaCharacteristics]: An array of media characteristics in preferred order.
//   - [IAVPlayerMediaSelectionCriteria.PrincipalMediaCharacteristics]: An array of media characteristics that are essential to select when choosing media with a particular characteristic.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerMediaSelectionCriteria
type IAVPlayerMediaSelectionCriteria interface {
	objectivec.IObject

	// Topic: Creating media selection criteria

	// Creates media selection criteria with the preferred languages and media characteristics.
	InitWithPreferredLanguagesPreferredMediaCharacteristics(preferredLanguages []string, preferredMediaCharacteristics []string) AVPlayerMediaSelectionCriteria
	// Creates media selection criteria with the principal media characteristics, and preferred languages and media characteristics.
	InitWithPrincipalMediaCharacteristicsPreferredLanguagesPreferredMediaCharacteristics(principalMediaCharacteristics []string, preferredLanguages []string, preferredMediaCharacteristics []string) AVPlayerMediaSelectionCriteria

	// Topic: Retrieving selection criteria settings

	// An array of language identifiers in preferred order.
	PreferredLanguages() []string
	// An array of media characteristics in preferred order.
	PreferredMediaCharacteristics() []string
	// An array of media characteristics that are essential to select when choosing media with a particular characteristic.
	PrincipalMediaCharacteristics() []string
}

// Init initializes the instance.
func (p AVPlayerMediaSelectionCriteria) Init() AVPlayerMediaSelectionCriteria {
	rv := objc.Send[AVPlayerMediaSelectionCriteria](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p AVPlayerMediaSelectionCriteria) Autorelease() AVPlayerMediaSelectionCriteria {
	rv := objc.Send[AVPlayerMediaSelectionCriteria](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVPlayerMediaSelectionCriteria creates a new AVPlayerMediaSelectionCriteria instance.
func NewAVPlayerMediaSelectionCriteria() AVPlayerMediaSelectionCriteria {
	class := getAVPlayerMediaSelectionCriteriaClass()
	rv := objc.Send[AVPlayerMediaSelectionCriteria](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates media selection criteria with the preferred languages and media
// characteristics.
//
// preferredLanguages: An array of language identifier strings, in order of preference. This value
// may be `nil`.
//
// preferredMediaCharacteristics: An array of media characteristics, in order of preference. This value may
// be `nil`.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerMediaSelectionCriteria/init(preferredLanguages:preferredMediaCharacteristics:)
func NewPlayerMediaSelectionCriteriaWithPreferredLanguagesPreferredMediaCharacteristics(preferredLanguages []string, preferredMediaCharacteristics []string) AVPlayerMediaSelectionCriteria {
	instance := getAVPlayerMediaSelectionCriteriaClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithPreferredLanguages:preferredMediaCharacteristics:"), objectivec.StringSliceToNSArray(preferredLanguages), objectivec.StringSliceToNSArray(preferredMediaCharacteristics))
	return AVPlayerMediaSelectionCriteriaFromID(rv)
}

// Creates media selection criteria with the principal media characteristics,
// and preferred languages and media characteristics.
//
// principalMediaCharacteristics: An array of media characteristics that are essential to selecting media
// with the characteristic. This value may be `nil`.
//
// preferredLanguages: An array of language identifier strings, in order of preference. This value
// may be `nil`.
//
// preferredMediaCharacteristics: An array of media characteristics, in order of preference. This value may
// be `nil`.
//
// # Discussion
// 
// Principal media characteristics, when present, override language
// preferences when making selections within a specific media selection group.
// However, language preferences may still pertain to selections in other
// groups. For example, the system may consider language preferences when
// choosing whether to select nonforced subtitles for translation purposes.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerMediaSelectionCriteria/init(principalMediaCharacteristics:preferredLanguages:preferredMediaCharacteristics:)
func NewPlayerMediaSelectionCriteriaWithPrincipalMediaCharacteristicsPreferredLanguagesPreferredMediaCharacteristics(principalMediaCharacteristics []string, preferredLanguages []string, preferredMediaCharacteristics []string) AVPlayerMediaSelectionCriteria {
	instance := getAVPlayerMediaSelectionCriteriaClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithPrincipalMediaCharacteristics:preferredLanguages:preferredMediaCharacteristics:"), objectivec.StringSliceToNSArray(principalMediaCharacteristics), objectivec.StringSliceToNSArray(preferredLanguages), objectivec.StringSliceToNSArray(preferredMediaCharacteristics))
	return AVPlayerMediaSelectionCriteriaFromID(rv)
}

// Creates media selection criteria with the preferred languages and media
// characteristics.
//
// preferredLanguages: An array of language identifier strings, in order of preference. This value
// may be `nil`.
//
// preferredMediaCharacteristics: An array of media characteristics, in order of preference. This value may
// be `nil`.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerMediaSelectionCriteria/init(preferredLanguages:preferredMediaCharacteristics:)
func (p AVPlayerMediaSelectionCriteria) InitWithPreferredLanguagesPreferredMediaCharacteristics(preferredLanguages []string, preferredMediaCharacteristics []string) AVPlayerMediaSelectionCriteria {
	rv := objc.Send[AVPlayerMediaSelectionCriteria](p.ID, objc.Sel("initWithPreferredLanguages:preferredMediaCharacteristics:"), objectivec.StringSliceToNSArray(preferredLanguages), objectivec.StringSliceToNSArray(preferredMediaCharacteristics))
	return rv
}
// Creates media selection criteria with the principal media characteristics,
// and preferred languages and media characteristics.
//
// principalMediaCharacteristics: An array of media characteristics that are essential to selecting media
// with the characteristic. This value may be `nil`.
//
// preferredLanguages: An array of language identifier strings, in order of preference. This value
// may be `nil`.
//
// preferredMediaCharacteristics: An array of media characteristics, in order of preference. This value may
// be `nil`.
//
// # Discussion
// 
// Principal media characteristics, when present, override language
// preferences when making selections within a specific media selection group.
// However, language preferences may still pertain to selections in other
// groups. For example, the system may consider language preferences when
// choosing whether to select nonforced subtitles for translation purposes.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerMediaSelectionCriteria/init(principalMediaCharacteristics:preferredLanguages:preferredMediaCharacteristics:)
func (p AVPlayerMediaSelectionCriteria) InitWithPrincipalMediaCharacteristicsPreferredLanguagesPreferredMediaCharacteristics(principalMediaCharacteristics []string, preferredLanguages []string, preferredMediaCharacteristics []string) AVPlayerMediaSelectionCriteria {
	rv := objc.Send[AVPlayerMediaSelectionCriteria](p.ID, objc.Sel("initWithPrincipalMediaCharacteristics:preferredLanguages:preferredMediaCharacteristics:"), objectivec.StringSliceToNSArray(principalMediaCharacteristics), objectivec.StringSliceToNSArray(preferredLanguages), objectivec.StringSliceToNSArray(preferredMediaCharacteristics))
	return rv
}

// An array of language identifiers in preferred order.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerMediaSelectionCriteria/preferredLanguages
func (p AVPlayerMediaSelectionCriteria) PreferredLanguages() []string {
	rv := objc.Send[[]objc.ID](p.ID, objc.Sel("preferredLanguages"))
	return objc.ConvertSliceToStrings(rv)
}
// An array of media characteristics in preferred order.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerMediaSelectionCriteria/preferredMediaCharacteristics
func (p AVPlayerMediaSelectionCriteria) PreferredMediaCharacteristics() []string {
	rv := objc.Send[[]objc.ID](p.ID, objc.Sel("preferredMediaCharacteristics"))
	return objc.ConvertSliceToStrings(rv)
}
// An array of media characteristics that are essential to select when
// choosing media with a particular characteristic.
//
// # Discussion
// 
// If no option matches the principal media characteristics, the system
// chooses the default option in the group as the best match.
// 
// When making automatic selections, a player item treats principal media
// characteristics as criteria that supersede language preferences and
// preferred media characteristics.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerMediaSelectionCriteria/principalMediaCharacteristics
func (p AVPlayerMediaSelectionCriteria) PrincipalMediaCharacteristics() []string {
	rv := objc.Send[[]objc.ID](p.ID, objc.Sel("principalMediaCharacteristics"))
	return objc.ConvertSliceToStrings(rv)
}

