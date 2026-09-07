// Code generated from Apple documentation for texttospeech. DO NOT EDIT.

package texttospeech

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [BFProsodicState] class.
var (
	_BFProsodicStateClass     BFProsodicStateClass
	_BFProsodicStateClassOnce sync.Once
)

func getBFProsodicStateClass() BFProsodicStateClass {
	_BFProsodicStateClassOnce.Do(func() {
		_BFProsodicStateClass = BFProsodicStateClass{class: objc.GetClass("BFProsodicState")}
	})
	return _BFProsodicStateClass
}

// GetBFProsodicStateClass returns the class object for BFProsodicState.
func GetBFProsodicStateClass() BFProsodicStateClass {
	return getBFProsodicStateClass()
}

type BFProsodicStateClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (bc BFProsodicStateClass) Class() objc.Class {
	return bc.class
}

// Alloc allocates memory for a new instance of the class.
func (bc BFProsodicStateClass) Alloc() BFProsodicState {
	rv := objc.SendIfResponds[BFProsodicState](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [BFProsodicState.NumberLiteralMode]
//   - [BFProsodicState.SetNumberLiteralMode]
//   - [BFProsodicState.PitchBase]
//   - [BFProsodicState.SetPitchBase]
//   - [BFProsodicState.PitchModulation]
//   - [BFProsodicState.SetPitchModulation]
//   - [BFProsodicState.Rate]
//   - [BFProsodicState.SetRate]
//   - [BFProsodicState.TextLiteralMode]
//   - [BFProsodicState.SetTextLiteralMode]
//   - [BFProsodicState.Volume]
//   - [BFProsodicState.SetVolume]
type BFProsodicState struct {
	objectivec.Object
}

// BFProsodicStateFromID constructs a [BFProsodicState] from an objc.ID.
func BFProsodicStateFromID(id objc.ID) BFProsodicState {
	return BFProsodicState{objectivec.Object{ID: id}}
}

// Ensure BFProsodicState implements IBFProsodicState.
var _ IBFProsodicState = BFProsodicState{}

// An interface definition for the [BFProsodicState] class.
//
// # Methods
//
//   - [IBFProsodicState.NumberLiteralMode]
//   - [IBFProsodicState.SetNumberLiteralMode]
//   - [IBFProsodicState.PitchBase]
//   - [IBFProsodicState.SetPitchBase]
//   - [IBFProsodicState.PitchModulation]
//   - [IBFProsodicState.SetPitchModulation]
//   - [IBFProsodicState.Rate]
//   - [IBFProsodicState.SetRate]
//   - [IBFProsodicState.TextLiteralMode]
//   - [IBFProsodicState.SetTextLiteralMode]
//   - [IBFProsodicState.Volume]
//   - [IBFProsodicState.SetVolume]
type IBFProsodicState interface {
	objectivec.IObject

	// Topic: Methods

	NumberLiteralMode() bool
	SetNumberLiteralMode(value bool)
	PitchBase() foundation.NSNumber
	SetPitchBase(value foundation.NSNumber)
	PitchModulation() foundation.NSNumber
	SetPitchModulation(value foundation.NSNumber)
	Rate() foundation.NSNumber
	SetRate(value foundation.NSNumber)
	TextLiteralMode() bool
	SetTextLiteralMode(value bool)
	Volume() foundation.NSNumber
	SetVolume(value foundation.NSNumber)
}

// Init initializes the instance.
func (b BFProsodicState) Init() BFProsodicState {
	rv := objc.SendIfResponds[BFProsodicState](b.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b BFProsodicState) Autorelease() BFProsodicState {
	rv := objc.SendIfResponds[BFProsodicState](b.ID, objc.Sel("autorelease"))
	return rv
}

// NewBFProsodicState creates a new BFProsodicState instance.
func NewBFProsodicState() BFProsodicState {
	class := getBFProsodicStateClass()
	rv := objc.SendIfResponds[BFProsodicState](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (b BFProsodicState) NumberLiteralMode() bool {
	rv := objc.SendIfResponds[bool](b.ID, objc.Sel("numberLiteralMode"))
	return rv
}
func (b BFProsodicState) SetNumberLiteralMode(value bool) {
	objc.SendIfResponds[struct{}](b.ID, objc.Sel("setNumberLiteralMode:"), value)
}
func (b BFProsodicState) PitchBase() foundation.NSNumber {
	rv := objc.SendIfResponds[foundation.NSNumber](b.ID, objc.Sel("pitchBase"))
	return foundation.NSNumber(rv)
}
func (b BFProsodicState) SetPitchBase(value foundation.NSNumber) {
	objc.SendIfResponds[struct{}](b.ID, objc.Sel("setPitchBase:"), value)
}
func (b BFProsodicState) PitchModulation() foundation.NSNumber {
	rv := objc.SendIfResponds[foundation.NSNumber](b.ID, objc.Sel("pitchModulation"))
	return foundation.NSNumber(rv)
}
func (b BFProsodicState) SetPitchModulation(value foundation.NSNumber) {
	objc.SendIfResponds[struct{}](b.ID, objc.Sel("setPitchModulation:"), value)
}
func (b BFProsodicState) Rate() foundation.NSNumber {
	rv := objc.SendIfResponds[foundation.NSNumber](b.ID, objc.Sel("rate"))
	return foundation.NSNumber(rv)
}
func (b BFProsodicState) SetRate(value foundation.NSNumber) {
	objc.SendIfResponds[struct{}](b.ID, objc.Sel("setRate:"), value)
}
func (b BFProsodicState) TextLiteralMode() bool {
	rv := objc.SendIfResponds[bool](b.ID, objc.Sel("textLiteralMode"))
	return rv
}
func (b BFProsodicState) SetTextLiteralMode(value bool) {
	objc.SendIfResponds[struct{}](b.ID, objc.Sel("setTextLiteralMode:"), value)
}
func (b BFProsodicState) Volume() foundation.NSNumber {
	rv := objc.SendIfResponds[foundation.NSNumber](b.ID, objc.Sel("volume"))
	return foundation.NSNumber(rv)
}
func (b BFProsodicState) SetVolume(value foundation.NSNumber) {
	objc.SendIfResponds[struct{}](b.ID, objc.Sel("setVolume:"), value)
}
