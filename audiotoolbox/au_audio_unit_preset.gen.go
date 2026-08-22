// Code generated from Apple documentation for AudioToolbox. DO NOT EDIT.

package audiotoolbox

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AUAudioUnitPreset] class.
var (
	_AUAudioUnitPresetClass     AUAudioUnitPresetClass
	_AUAudioUnitPresetClassOnce sync.Once
)

func getAUAudioUnitPresetClass() AUAudioUnitPresetClass {
	_AUAudioUnitPresetClassOnce.Do(func() {
		_AUAudioUnitPresetClass = AUAudioUnitPresetClass{class: objc.GetClass("AUAudioUnitPreset")}
	})
	return _AUAudioUnitPresetClass
}

// GetAUAudioUnitPresetClass returns the class object for AUAudioUnitPreset.
func GetAUAudioUnitPresetClass() AUAudioUnitPresetClass {
	return getAUAudioUnitPresetClass()
}

type AUAudioUnitPresetClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AUAudioUnitPresetClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AUAudioUnitPresetClass) Alloc() AUAudioUnitPreset {
	rv := objc.Send[AUAudioUnitPreset](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// A class that describes an interface for custom parameter settings provided
// by the audio unit developer.
//
// # Overview
//
// These presets often produce a useful sound or starting point.
//
// For more details on working with Audio Unit presets, see [Audio Units - How
// to correctly save and restore Audio Unit presets.] Note that the version 3
// [AUAudioUnit.FullState] property is bridged to the version 2
// `kAudioUnitProperty_ClassInfo` API. Similarly, the version 3
// [AUAudioUnit.FullStateForDocument] property is bridged to the version 2
// `kAudioUnitProperty_ClassInfoFromDocument` API.
//
// # Preset Properties
//
//   - [AUAudioUnitPreset.Name]: The preset’s name.
//   - [AUAudioUnitPreset.SetName]
//   - [AUAudioUnitPreset.Number]: The preset’s unique numeric identifier.
//   - [AUAudioUnitPreset.SetNumber]
//
// # Initializers
//
//   - [AUAudioUnitPreset.InitWithCoder]
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnitPreset
//
// [Audio Units - How to correctly save and restore Audio Unit presets.]: https://developer.apple.com/library/archive/technotes/tn2157/_index.html#//apple_ref/doc/uid/DTS40011953
type AUAudioUnitPreset struct {
	objectivec.Object
}

// AUAudioUnitPresetFromID constructs a [AUAudioUnitPreset] from an objc.ID.
//
// A class that describes an interface for custom parameter settings provided
// by the audio unit developer.
func AUAudioUnitPresetFromID(id objc.ID) AUAudioUnitPreset {
	return AUAudioUnitPreset{objectivec.Object{ID: id}}
}

// NOTE: AUAudioUnitPreset adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AUAudioUnitPreset] class.
//
// # Preset Properties
//
//   - [IAUAudioUnitPreset.Name]: The preset’s name.
//   - [IAUAudioUnitPreset.SetName]
//   - [IAUAudioUnitPreset.Number]: The preset’s unique numeric identifier.
//   - [IAUAudioUnitPreset.SetNumber]
//
// # Initializers
//
//   - [IAUAudioUnitPreset.InitWithCoder]
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnitPreset
type IAUAudioUnitPreset interface {
	objectivec.IObject

	// Topic: Preset Properties

	// The preset’s name.
	Name() string
	SetName(value string)
	// The preset’s unique numeric identifier.
	Number() int
	SetNumber(value int)

	// Topic: Initializers

	InitWithCoder(coder foundation.INSCoder) AUAudioUnitPreset

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (a AUAudioUnitPreset) Init() AUAudioUnitPreset {
	rv := objc.Send[AUAudioUnitPreset](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AUAudioUnitPreset) Autorelease() AUAudioUnitPreset {
	rv := objc.Send[AUAudioUnitPreset](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAUAudioUnitPreset creates a new AUAudioUnitPreset instance.
func NewAUAudioUnitPreset() AUAudioUnitPreset {
	class := getAUAudioUnitPresetClass()
	rv := objc.Send[AUAudioUnitPreset](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnitPreset/init(coder:)
func NewAudioUnitPresetWithCoder(coder foundation.INSCoder) AUAudioUnitPreset {
	instance := getAUAudioUnitPresetClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return AUAudioUnitPresetFromID(rv)
}

// See: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnitPreset/init(coder:)
func (a AUAudioUnitPreset) InitWithCoder(coder foundation.INSCoder) AUAudioUnitPreset {
	rv := objc.Send[AUAudioUnitPreset](a.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}
func (a AUAudioUnitPreset) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](a.ID, objc.Sel("encodeWithCoder:"), coder)
}

// The preset’s name.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnitPreset/name
func (a AUAudioUnitPreset) Name() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("name"))
	return foundation.NSStringFromID(rv).String()
}
func (a AUAudioUnitPreset) SetName(value string) {
	objc.Send[struct{}](a.ID, objc.Sel("setName:"), objc.String(value))
}

// The preset’s unique numeric identifier.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnitPreset/number
func (a AUAudioUnitPreset) Number() int {
	rv := objc.Send[int](a.ID, objc.Sel("number"))
	return rv
}
func (a AUAudioUnitPreset) SetNumber(value int) {
	objc.Send[struct{}](a.ID, objc.Sel("setNumber:"), value)
}
