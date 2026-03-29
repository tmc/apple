// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVVCDuckOverride] class.
var (
	_AVVCDuckOverrideClass     AVVCDuckOverrideClass
	_AVVCDuckOverrideClassOnce sync.Once
)

func getAVVCDuckOverrideClass() AVVCDuckOverrideClass {
	_AVVCDuckOverrideClassOnce.Do(func() {
		_AVVCDuckOverrideClass = AVVCDuckOverrideClass{class: objc.GetClass("AVVCDuckOverride")}
	})
	return _AVVCDuckOverrideClass
}

// GetAVVCDuckOverrideClass returns the class object for AVVCDuckOverride.
func GetAVVCDuckOverrideClass() AVVCDuckOverrideClass {
	return getAVVCDuckOverrideClass()
}

type AVVCDuckOverrideClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVVCDuckOverrideClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVVCDuckOverrideClass) Alloc() AVVCDuckOverride {
	rv := objc.Send[AVVCDuckOverride](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [AVVCDuckOverride.DuckOthers]
//   - [AVVCDuckOverride.SetDuckOthers]
//   - [AVVCDuckOverride.DuckToLevel]
//   - [AVVCDuckOverride.SetDuckToLevel]
//   - [AVVCDuckOverride.IsBlur]
//   - [AVVCDuckOverride.SetIsBlur]
//   - [AVVCDuckOverride.MixWithOthers]
//   - [AVVCDuckOverride.SetMixWithOthers]
//   - [AVVCDuckOverride.InitWithDuckOthersDuckToLevelMixWithOthers]
// See: https://developer.apple.com/documentation/AVFAudio/AVVCDuckOverride
type AVVCDuckOverride struct {
	objectivec.Object
}

// AVVCDuckOverrideFromID constructs a [AVVCDuckOverride] from an objc.ID.
func AVVCDuckOverrideFromID(id objc.ID) AVVCDuckOverride {
	return AVVCDuckOverride{objectivec.Object{ID: id}}
}
// Ensure AVVCDuckOverride implements IAVVCDuckOverride.
var _ IAVVCDuckOverride = AVVCDuckOverride{}

// An interface definition for the [AVVCDuckOverride] class.
//
// # Methods
//
//   - [IAVVCDuckOverride.DuckOthers]
//   - [IAVVCDuckOverride.SetDuckOthers]
//   - [IAVVCDuckOverride.DuckToLevel]
//   - [IAVVCDuckOverride.SetDuckToLevel]
//   - [IAVVCDuckOverride.IsBlur]
//   - [IAVVCDuckOverride.SetIsBlur]
//   - [IAVVCDuckOverride.MixWithOthers]
//   - [IAVVCDuckOverride.SetMixWithOthers]
//   - [IAVVCDuckOverride.InitWithDuckOthersDuckToLevelMixWithOthers]
//
// See: https://developer.apple.com/documentation/AVFAudio/AVVCDuckOverride
type IAVVCDuckOverride interface {
	objectivec.IObject

	// Topic: Methods

	DuckOthers() foundation.NSNumber
	SetDuckOthers(value foundation.NSNumber)
	DuckToLevel() foundation.NSNumber
	SetDuckToLevel(value foundation.NSNumber)
	IsBlur() bool
	SetIsBlur(value bool)
	MixWithOthers() foundation.NSNumber
	SetMixWithOthers(value foundation.NSNumber)
	InitWithDuckOthersDuckToLevelMixWithOthers(others objectivec.IObject, level objectivec.IObject, others2 objectivec.IObject) AVVCDuckOverride
}

// Init initializes the instance.
func (v AVVCDuckOverride) Init() AVVCDuckOverride {
	rv := objc.Send[AVVCDuckOverride](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v AVVCDuckOverride) Autorelease() AVVCDuckOverride {
	rv := objc.Send[AVVCDuckOverride](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVVCDuckOverride creates a new AVVCDuckOverride instance.
func NewAVVCDuckOverride() AVVCDuckOverride {
	class := getAVVCDuckOverrideClass()
	rv := objc.Send[AVVCDuckOverride](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/AVFAudio/AVVCDuckOverride/initWithDuckOthers:duckToLevel:mixWithOthers:
func NewVCDuckOverrideWithDuckOthersDuckToLevelMixWithOthers(others objectivec.IObject, level objectivec.IObject, others2 objectivec.IObject) AVVCDuckOverride {
	instance := getAVVCDuckOverrideClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDuckOthers:duckToLevel:mixWithOthers:"), others, level, others2)
	return AVVCDuckOverrideFromID(rv)
}

//
// See: https://developer.apple.com/documentation/AVFAudio/AVVCDuckOverride/initWithDuckOthers:duckToLevel:mixWithOthers:
func (v AVVCDuckOverride) InitWithDuckOthersDuckToLevelMixWithOthers(others objectivec.IObject, level objectivec.IObject, others2 objectivec.IObject) AVVCDuckOverride {
	rv := objc.Send[AVVCDuckOverride](v.ID, objc.Sel("initWithDuckOthers:duckToLevel:mixWithOthers:"), others, level, others2)
	return rv
}

// See: https://developer.apple.com/documentation/AVFAudio/AVVCDuckOverride/duckOthers
func (v AVVCDuckOverride) DuckOthers() foundation.NSNumber {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("duckOthers"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
func (v AVVCDuckOverride) SetDuckOthers(value foundation.NSNumber) {
	objc.Send[struct{}](v.ID, objc.Sel("setDuckOthers:"), value)
}
// See: https://developer.apple.com/documentation/AVFAudio/AVVCDuckOverride/duckToLevel
func (v AVVCDuckOverride) DuckToLevel() foundation.NSNumber {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("duckToLevel"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
func (v AVVCDuckOverride) SetDuckToLevel(value foundation.NSNumber) {
	objc.Send[struct{}](v.ID, objc.Sel("setDuckToLevel:"), value)
}
// See: https://developer.apple.com/documentation/AVFAudio/AVVCDuckOverride/isBlur
func (v AVVCDuckOverride) IsBlur() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("isBlur"))
	return rv
}
func (v AVVCDuckOverride) SetIsBlur(value bool) {
	objc.Send[struct{}](v.ID, objc.Sel("setIsBlur:"), value)
}
// See: https://developer.apple.com/documentation/AVFAudio/AVVCDuckOverride/mixWithOthers
func (v AVVCDuckOverride) MixWithOthers() foundation.NSNumber {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("mixWithOthers"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
func (v AVVCDuckOverride) SetMixWithOthers(value foundation.NSNumber) {
	objc.Send[struct{}](v.ID, objc.Sel("setMixWithOthers:"), value)
}

