// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVVCDuckLevel] class.
var (
	_AVVCDuckLevelClass     AVVCDuckLevelClass
	_AVVCDuckLevelClassOnce sync.Once
)

func getAVVCDuckLevelClass() AVVCDuckLevelClass {
	_AVVCDuckLevelClassOnce.Do(func() {
		_AVVCDuckLevelClass = AVVCDuckLevelClass{class: objc.GetClass("AVVCDuckLevel")}
	})
	return _AVVCDuckLevelClass
}

// GetAVVCDuckLevelClass returns the class object for AVVCDuckLevel.
func GetAVVCDuckLevelClass() AVVCDuckLevelClass {
	return getAVVCDuckLevelClass()
}

type AVVCDuckLevelClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVVCDuckLevelClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVVCDuckLevelClass) Alloc() AVVCDuckLevel {
	rv := objc.Send[AVVCDuckLevel](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [AVVCDuckLevel.IsBlur]
//   - [AVVCDuckLevel.SetIsBlur]
//   - [AVVCDuckLevel.Value]
//   - [AVVCDuckLevel.SetValue]
//
// See: https://developer.apple.com/documentation/AVFAudio/AVVCDuckLevel
type AVVCDuckLevel struct {
	objectivec.Object
}

// AVVCDuckLevelFromID constructs a [AVVCDuckLevel] from an objc.ID.
func AVVCDuckLevelFromID(id objc.ID) AVVCDuckLevel {
	return AVVCDuckLevel{objectivec.Object{ID: id}}
}

// Ensure AVVCDuckLevel implements IAVVCDuckLevel.
var _ IAVVCDuckLevel = AVVCDuckLevel{}

// An interface definition for the [AVVCDuckLevel] class.
//
// # Methods
//
//   - [IAVVCDuckLevel.IsBlur]
//   - [IAVVCDuckLevel.SetIsBlur]
//   - [IAVVCDuckLevel.Value]
//   - [IAVVCDuckLevel.SetValue]
//
// See: https://developer.apple.com/documentation/AVFAudio/AVVCDuckLevel
type IAVVCDuckLevel interface {
	objectivec.IObject

	// Topic: Methods

	IsBlur() bool
	SetIsBlur(value bool)
	Value() foundation.NSNumber
	SetValue(value foundation.NSNumber)
}

// Init initializes the instance.
func (v AVVCDuckLevel) Init() AVVCDuckLevel {
	rv := objc.Send[AVVCDuckLevel](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v AVVCDuckLevel) Autorelease() AVVCDuckLevel {
	rv := objc.Send[AVVCDuckLevel](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVVCDuckLevel creates a new AVVCDuckLevel instance.
func NewAVVCDuckLevel() AVVCDuckLevel {
	class := getAVVCDuckLevelClass()
	rv := objc.Send[AVVCDuckLevel](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/AVFAudio/AVVCDuckLevel/isBlur
func (v AVVCDuckLevel) IsBlur() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("isBlur"))
	return rv
}
func (v AVVCDuckLevel) SetIsBlur(value bool) {
	objc.Send[struct{}](v.ID, objc.Sel("setIsBlur:"), value)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVVCDuckLevel/value
func (v AVVCDuckLevel) Value() foundation.NSNumber {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("value"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
func (v AVVCDuckLevel) SetValue(value foundation.NSNumber) {
	objc.Send[struct{}](v.ID, objc.Sel("setValue:"), value)
}
