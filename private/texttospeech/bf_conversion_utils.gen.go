// Code generated from Apple documentation for texttospeech. DO NOT EDIT.

package texttospeech

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [BFConversionUtils] class.
var (
	_BFConversionUtilsClass     BFConversionUtilsClass
	_BFConversionUtilsClassOnce sync.Once
)

func getBFConversionUtilsClass() BFConversionUtilsClass {
	_BFConversionUtilsClassOnce.Do(func() {
		_BFConversionUtilsClass = BFConversionUtilsClass{class: objc.GetClass("BFConversionUtils")}
	})
	return _BFConversionUtilsClass
}

// GetBFConversionUtilsClass returns the class object for BFConversionUtils.
func GetBFConversionUtilsClass() BFConversionUtilsClass {
	return getBFConversionUtilsClass()
}

type BFConversionUtilsClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (bc BFConversionUtilsClass) Class() objc.Class {
	return bc.class
}

// Alloc allocates memory for a new instance of the class.
func (bc BFConversionUtilsClass) Alloc() BFConversionUtils {
	rv := objc.SendIfResponds[BFConversionUtils](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [BFConversionUtils._assymetricLinearlyScaledValueForNumberSourceMinSourceMiddleSourceMaxDestinationMinDestinationMiddleDestinationMax]
//   - [BFConversionUtils.ConvertAbsolutePitchForVoice]
//   - [BFConversionUtils.ConvertPitchModulationForVoice]
//   - [BFConversionUtils.ConvertWPMRateForVoice]
//   - [BFConversionUtils.DefaultPitchForVoice]
//   - [BFConversionUtils.DefaultPitchModulationForVoice]
//   - [BFConversionUtils.DefaultRateForVoice]
//   - [BFConversionUtils.DefaultVolumeForVoice]
//   - [BFConversionUtils.VoiceDefaults]
//   - [BFConversionUtils.SetVoiceDefaults]
type BFConversionUtils struct {
	objectivec.Object
}

// BFConversionUtilsFromID constructs a [BFConversionUtils] from an objc.ID.
func BFConversionUtilsFromID(id objc.ID) BFConversionUtils {
	return BFConversionUtils{objectivec.Object{ID: id}}
}

// Ensure BFConversionUtils implements IBFConversionUtils.
var _ IBFConversionUtils = BFConversionUtils{}

// An interface definition for the [BFConversionUtils] class.
//
// # Methods
//
//   - [IBFConversionUtils._assymetricLinearlyScaledValueForNumberSourceMinSourceMiddleSourceMaxDestinationMinDestinationMiddleDestinationMax]
//   - [IBFConversionUtils.ConvertAbsolutePitchForVoice]
//   - [IBFConversionUtils.ConvertPitchModulationForVoice]
//   - [IBFConversionUtils.ConvertWPMRateForVoice]
//   - [IBFConversionUtils.DefaultPitchForVoice]
//   - [IBFConversionUtils.DefaultPitchModulationForVoice]
//   - [IBFConversionUtils.DefaultRateForVoice]
//   - [IBFConversionUtils.DefaultVolumeForVoice]
//   - [IBFConversionUtils.VoiceDefaults]
//   - [IBFConversionUtils.SetVoiceDefaults]
type IBFConversionUtils interface {
	objectivec.IObject

	// Topic: Methods

	_assymetricLinearlyScaledValueForNumberSourceMinSourceMiddleSourceMaxDestinationMinDestinationMiddleDestinationMax(number float32, min float32, middle float32, max float32, min2 float32, middle2 float32, max2 float32) float32
	ConvertAbsolutePitchForVoice(pitch objectivec.IObject, voice objectivec.IObject) float32
	ConvertPitchModulationForVoice(modulation objectivec.IObject, voice objectivec.IObject) float32
	ConvertWPMRateForVoice(wPMRate objectivec.IObject, voice objectivec.IObject) float32
	DefaultPitchForVoice(voice objectivec.IObject) objectivec.IObject
	DefaultPitchModulationForVoice(voice objectivec.IObject) objectivec.IObject
	DefaultRateForVoice(voice objectivec.IObject) objectivec.IObject
	DefaultVolumeForVoice(voice objectivec.IObject) objectivec.IObject
	VoiceDefaults() foundation.INSDictionary
	SetVoiceDefaults(value foundation.INSDictionary)
}

// Init initializes the instance.
func (b BFConversionUtils) Init() BFConversionUtils {
	rv := objc.SendIfResponds[BFConversionUtils](b.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b BFConversionUtils) Autorelease() BFConversionUtils {
	rv := objc.SendIfResponds[BFConversionUtils](b.ID, objc.Sel("autorelease"))
	return rv
}

// NewBFConversionUtils creates a new BFConversionUtils instance.
func NewBFConversionUtils() BFConversionUtils {
	class := getBFConversionUtilsClass()
	rv := objc.SendIfResponds[BFConversionUtils](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (b BFConversionUtils) _assymetricLinearlyScaledValueForNumberSourceMinSourceMiddleSourceMaxDestinationMinDestinationMiddleDestinationMax(number float32, min float32, middle float32, max float32, min2 float32, middle2 float32, max2 float32) float32 {
	rv := objc.SendIfResponds[float32](b.ID, objc.Sel("_assymetricLinearlyScaledValueForNumber:sourceMin:sourceMiddle:sourceMax:destinationMin:destinationMiddle:destinationMax:"), number, min, middle, max, min2, middle2, max2)
	return rv
}

// AssymetricLinearlyScaledValueForNumberSourceMinSourceMiddleSourceMaxDestinationMinDestinationMiddleDestinationMax is an exported wrapper for the private method _assymetricLinearlyScaledValueForNumberSourceMinSourceMiddleSourceMaxDestinationMinDestinationMiddleDestinationMax.
func (b BFConversionUtils) AssymetricLinearlyScaledValueForNumberSourceMinSourceMiddleSourceMaxDestinationMinDestinationMiddleDestinationMax(number float32, min float32, middle float32, max float32, min2 float32, middle2 float32, max2 float32) (float32, error) {
	if !objc.RespondsToSelector(b.ID, objc.Sel("_assymetricLinearlyScaledValueForNumber:sourceMin:sourceMiddle:sourceMax:destinationMin:destinationMiddle:destinationMax:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_assymetricLinearlyScaledValueForNumber:sourceMin:sourceMiddle:sourceMax:destinationMin:destinationMiddle:destinationMax:"}
		return 0.0, err
	}
	return b._assymetricLinearlyScaledValueForNumberSourceMinSourceMiddleSourceMaxDestinationMinDestinationMiddleDestinationMax(number, min, middle, max, min2, middle2, max2), nil
}

// CanAssymetricLinearlyScaledValueForNumberSourceMinSourceMiddleSourceMaxDestinationMinDestinationMiddleDestinationMax reports whether the receiver responds to the private selector _assymetricLinearlyScaledValueForNumber:sourceMin:sourceMiddle:sourceMax:destinationMin:destinationMiddle:destinationMax:.
func (b BFConversionUtils) CanAssymetricLinearlyScaledValueForNumberSourceMinSourceMiddleSourceMaxDestinationMinDestinationMiddleDestinationMax() bool {
	return objc.RespondsToSelector(b.ID, objc.Sel("_assymetricLinearlyScaledValueForNumber:sourceMin:sourceMiddle:sourceMax:destinationMin:destinationMiddle:destinationMax:"))
}
func (b BFConversionUtils) ConvertAbsolutePitchForVoice(pitch objectivec.IObject, voice objectivec.IObject) float32 {
	rv := objc.SendIfResponds[float32](b.ID, objc.Sel("convertAbsolutePitch:forVoice:"), pitch, voice)
	return rv
}
func (b BFConversionUtils) ConvertPitchModulationForVoice(modulation objectivec.IObject, voice objectivec.IObject) float32 {
	rv := objc.SendIfResponds[float32](b.ID, objc.Sel("convertPitchModulation:forVoice:"), modulation, voice)
	return rv
}
func (b BFConversionUtils) ConvertWPMRateForVoice(wPMRate objectivec.IObject, voice objectivec.IObject) float32 {
	rv := objc.SendIfResponds[float32](b.ID, objc.Sel("convertWPMRate:forVoice:"), wPMRate, voice)
	return rv
}
func (b BFConversionUtils) DefaultPitchForVoice(voice objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](b.ID, objc.Sel("defaultPitchForVoice:"), voice)
	return objectivec.Object{ID: rv}
}
func (b BFConversionUtils) DefaultPitchModulationForVoice(voice objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](b.ID, objc.Sel("defaultPitchModulationForVoice:"), voice)
	return objectivec.Object{ID: rv}
}
func (b BFConversionUtils) DefaultRateForVoice(voice objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](b.ID, objc.Sel("defaultRateForVoice:"), voice)
	return objectivec.Object{ID: rv}
}
func (b BFConversionUtils) DefaultVolumeForVoice(voice objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](b.ID, objc.Sel("defaultVolumeForVoice:"), voice)
	return objectivec.Object{ID: rv}
}

func (_BFConversionUtilsClass BFConversionUtilsClass) SharedInstance() BFConversionUtils {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_BFConversionUtilsClass.class), objc.Sel("sharedInstance"))
	return BFConversionUtilsFromID(rv)
}

func (b BFConversionUtils) VoiceDefaults() foundation.INSDictionary {
	rv := objc.SendIfResponds[objc.ID](b.ID, objc.Sel("voiceDefaults"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (b BFConversionUtils) SetVoiceDefaults(value foundation.INSDictionary) {
	objc.SendIfResponds[struct{}](b.ID, objc.Sel("setVoiceDefaults:"), value)
}
