// Code generated from Apple documentation for texttospeech. DO NOT EDIT.

package texttospeech

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [BFPauseSpeechElement] class.
var (
	_BFPauseSpeechElementClass     BFPauseSpeechElementClass
	_BFPauseSpeechElementClassOnce sync.Once
)

func getBFPauseSpeechElementClass() BFPauseSpeechElementClass {
	_BFPauseSpeechElementClassOnce.Do(func() {
		_BFPauseSpeechElementClass = BFPauseSpeechElementClass{class: objc.GetClass("BFPauseSpeechElement")}
	})
	return _BFPauseSpeechElementClass
}

// GetBFPauseSpeechElementClass returns the class object for BFPauseSpeechElement.
func GetBFPauseSpeechElementClass() BFPauseSpeechElementClass {
	return getBFPauseSpeechElementClass()
}

type BFPauseSpeechElementClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (bc BFPauseSpeechElementClass) Class() objc.Class {
	return bc.class
}

// Alloc allocates memory for a new instance of the class.
func (bc BFPauseSpeechElementClass) Alloc() BFPauseSpeechElement {
	rv := objc.SendIfResponds[BFPauseSpeechElement](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [BFPauseSpeechElement.Milliseconds]
//   - [BFPauseSpeechElement.SetMilliseconds]
//   - [BFPauseSpeechElement.InitWithMSDurationAndRange]
type BFPauseSpeechElement struct {
	BFSpeechElement
}

// BFPauseSpeechElementFromID constructs a [BFPauseSpeechElement] from an objc.ID.
func BFPauseSpeechElementFromID(id objc.ID) BFPauseSpeechElement {
	return BFPauseSpeechElement{BFSpeechElement: BFSpeechElementFromID(id)}
}

// Ensure BFPauseSpeechElement implements IBFPauseSpeechElement.
var _ IBFPauseSpeechElement = BFPauseSpeechElement{}

// An interface definition for the [BFPauseSpeechElement] class.
//
// # Methods
//
//   - [IBFPauseSpeechElement.Milliseconds]
//   - [IBFPauseSpeechElement.SetMilliseconds]
//   - [IBFPauseSpeechElement.InitWithMSDurationAndRange]
type IBFPauseSpeechElement interface {
	IBFSpeechElement

	// Topic: Methods

	Milliseconds() uint64
	SetMilliseconds(value uint64)
	InitWithMSDurationAndRange(mSDuration objectivec.IObject, range_ foundation.NSRange) BFPauseSpeechElement
}

// Init initializes the instance.
func (b BFPauseSpeechElement) Init() BFPauseSpeechElement {
	rv := objc.SendIfResponds[BFPauseSpeechElement](b.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b BFPauseSpeechElement) Autorelease() BFPauseSpeechElement {
	rv := objc.SendIfResponds[BFPauseSpeechElement](b.ID, objc.Sel("autorelease"))
	return rv
}

// NewBFPauseSpeechElement creates a new BFPauseSpeechElement instance.
func NewBFPauseSpeechElement() BFPauseSpeechElement {
	class := getBFPauseSpeechElementClass()
	rv := objc.SendIfResponds[BFPauseSpeechElement](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewBFPauseSpeechElementWithMSDurationAndRange(mSDuration objectivec.IObject, range_ foundation.NSRange) BFPauseSpeechElement {
	instance := getBFPauseSpeechElementClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithMSDuration:andRange:"), mSDuration, range_)
	return BFPauseSpeechElementFromID(rv)
}

func NewBFPauseSpeechElementWithRange(range_ foundation.NSRange) BFPauseSpeechElement {
	instance := getBFPauseSpeechElementClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithRange:"), range_)
	return BFPauseSpeechElementFromID(rv)
}

func (b BFPauseSpeechElement) InitWithMSDurationAndRange(mSDuration objectivec.IObject, range_ foundation.NSRange) BFPauseSpeechElement {
	rv := objc.SendIfResponds[BFPauseSpeechElement](b.ID, objc.Sel("initWithMSDuration:andRange:"), mSDuration, range_)
	return rv
}

func (b BFPauseSpeechElement) Milliseconds() uint64 {
	rv := objc.SendIfResponds[uint64](b.ID, objc.Sel("milliseconds"))
	return rv
}
func (b BFPauseSpeechElement) SetMilliseconds(value uint64) {
	objc.SendIfResponds[struct{}](b.ID, objc.Sel("setMilliseconds:"), value)
}
