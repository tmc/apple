// Code generated from Apple documentation for texttospeech. DO NOT EDIT.

package texttospeech

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [BFIsoProsodicTextSpeechElement] class.
var (
	_BFIsoProsodicTextSpeechElementClass     BFIsoProsodicTextSpeechElementClass
	_BFIsoProsodicTextSpeechElementClassOnce sync.Once
)

func getBFIsoProsodicTextSpeechElementClass() BFIsoProsodicTextSpeechElementClass {
	_BFIsoProsodicTextSpeechElementClassOnce.Do(func() {
		_BFIsoProsodicTextSpeechElementClass = BFIsoProsodicTextSpeechElementClass{class: objc.GetClass("BFIsoProsodicTextSpeechElement")}
	})
	return _BFIsoProsodicTextSpeechElementClass
}

// GetBFIsoProsodicTextSpeechElementClass returns the class object for BFIsoProsodicTextSpeechElement.
func GetBFIsoProsodicTextSpeechElementClass() BFIsoProsodicTextSpeechElementClass {
	return getBFIsoProsodicTextSpeechElementClass()
}

type BFIsoProsodicTextSpeechElementClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (bc BFIsoProsodicTextSpeechElementClass) Class() objc.Class {
	return bc.class
}

// Alloc allocates memory for a new instance of the class.
func (bc BFIsoProsodicTextSpeechElementClass) Alloc() BFIsoProsodicTextSpeechElement {
	rv := objc.SendIfResponds[BFIsoProsodicTextSpeechElement](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [BFIsoProsodicTextSpeechElement.Mode]
//   - [BFIsoProsodicTextSpeechElement.SetMode]
//   - [BFIsoProsodicTextSpeechElement.ProsodicState]
//   - [BFIsoProsodicTextSpeechElement.SetProsodicState]
//   - [BFIsoProsodicTextSpeechElement.Text]
//   - [BFIsoProsodicTextSpeechElement.SetText]
//   - [BFIsoProsodicTextSpeechElement.InitWithTextOriginalRangeProsodicStateMode]
type BFIsoProsodicTextSpeechElement struct {
	BFSpeechElement
}

// BFIsoProsodicTextSpeechElementFromID constructs a [BFIsoProsodicTextSpeechElement] from an objc.ID.
func BFIsoProsodicTextSpeechElementFromID(id objc.ID) BFIsoProsodicTextSpeechElement {
	return BFIsoProsodicTextSpeechElement{BFSpeechElement: BFSpeechElementFromID(id)}
}

// Ensure BFIsoProsodicTextSpeechElement implements IBFIsoProsodicTextSpeechElement.
var _ IBFIsoProsodicTextSpeechElement = BFIsoProsodicTextSpeechElement{}

// An interface definition for the [BFIsoProsodicTextSpeechElement] class.
//
// # Methods
//
//   - [IBFIsoProsodicTextSpeechElement.Mode]
//   - [IBFIsoProsodicTextSpeechElement.SetMode]
//   - [IBFIsoProsodicTextSpeechElement.ProsodicState]
//   - [IBFIsoProsodicTextSpeechElement.SetProsodicState]
//   - [IBFIsoProsodicTextSpeechElement.Text]
//   - [IBFIsoProsodicTextSpeechElement.SetText]
//   - [IBFIsoProsodicTextSpeechElement.InitWithTextOriginalRangeProsodicStateMode]
type IBFIsoProsodicTextSpeechElement interface {
	IBFSpeechElement

	// Topic: Methods

	Mode() uint64
	SetMode(value uint64)
	ProsodicState() IBFProsodicState
	SetProsodicState(value IBFProsodicState)
	Text() string
	SetText(value string)
	InitWithTextOriginalRangeProsodicStateMode(text objectivec.IObject, range_ foundation.NSRange, state objectivec.IObject, mode uint64) BFIsoProsodicTextSpeechElement
}

// Init initializes the instance.
func (b BFIsoProsodicTextSpeechElement) Init() BFIsoProsodicTextSpeechElement {
	rv := objc.SendIfResponds[BFIsoProsodicTextSpeechElement](b.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b BFIsoProsodicTextSpeechElement) Autorelease() BFIsoProsodicTextSpeechElement {
	rv := objc.SendIfResponds[BFIsoProsodicTextSpeechElement](b.ID, objc.Sel("autorelease"))
	return rv
}

// NewBFIsoProsodicTextSpeechElement creates a new BFIsoProsodicTextSpeechElement instance.
func NewBFIsoProsodicTextSpeechElement() BFIsoProsodicTextSpeechElement {
	class := getBFIsoProsodicTextSpeechElementClass()
	rv := objc.SendIfResponds[BFIsoProsodicTextSpeechElement](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewBFIsoProsodicTextSpeechElementWithRange(range_ foundation.NSRange) BFIsoProsodicTextSpeechElement {
	instance := getBFIsoProsodicTextSpeechElementClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithRange:"), range_)
	return BFIsoProsodicTextSpeechElementFromID(rv)
}

func NewBFIsoProsodicTextSpeechElementWithTextOriginalRangeProsodicStateMode(text objectivec.IObject, range_ foundation.NSRange, state objectivec.IObject, mode uint64) BFIsoProsodicTextSpeechElement {
	instance := getBFIsoProsodicTextSpeechElementClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithText:originalRange:prosodicState:mode:"), text, range_, state, mode)
	return BFIsoProsodicTextSpeechElementFromID(rv)
}

func (b BFIsoProsodicTextSpeechElement) InitWithTextOriginalRangeProsodicStateMode(text objectivec.IObject, range_ foundation.NSRange, state objectivec.IObject, mode uint64) BFIsoProsodicTextSpeechElement {
	rv := objc.SendIfResponds[BFIsoProsodicTextSpeechElement](b.ID, objc.Sel("initWithText:originalRange:prosodicState:mode:"), text, range_, state, mode)
	return rv
}

func (b BFIsoProsodicTextSpeechElement) Mode() uint64 {
	rv := objc.SendIfResponds[uint64](b.ID, objc.Sel("mode"))
	return rv
}
func (b BFIsoProsodicTextSpeechElement) SetMode(value uint64) {
	objc.SendIfResponds[struct{}](b.ID, objc.Sel("setMode:"), value)
}
func (b BFIsoProsodicTextSpeechElement) ProsodicState() IBFProsodicState {
	rv := objc.SendIfResponds[objc.ID](b.ID, objc.Sel("prosodicState"))
	return BFProsodicStateFromID(objc.ID(rv))
}
func (b BFIsoProsodicTextSpeechElement) SetProsodicState(value IBFProsodicState) {
	objc.SendIfResponds[struct{}](b.ID, objc.Sel("setProsodicState:"), value)
}
func (b BFIsoProsodicTextSpeechElement) Text() string {
	rv := objc.SendIfResponds[objc.ID](b.ID, objc.Sel("text"))
	return foundation.NSStringFromID(rv).String()
}
func (b BFIsoProsodicTextSpeechElement) SetText(value string) {
	objc.SendIfResponds[struct{}](b.ID, objc.Sel("setText:"), objc.String(value))
}
