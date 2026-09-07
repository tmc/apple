// Code generated from Apple documentation for texttospeech. DO NOT EDIT.

package texttospeech

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// The class instance for the [BFSkipSpeechElement] class.
var (
	_BFSkipSpeechElementClass     BFSkipSpeechElementClass
	_BFSkipSpeechElementClassOnce sync.Once
)

func getBFSkipSpeechElementClass() BFSkipSpeechElementClass {
	_BFSkipSpeechElementClassOnce.Do(func() {
		_BFSkipSpeechElementClass = BFSkipSpeechElementClass{class: objc.GetClass("BFSkipSpeechElement")}
	})
	return _BFSkipSpeechElementClass
}

// GetBFSkipSpeechElementClass returns the class object for BFSkipSpeechElement.
func GetBFSkipSpeechElementClass() BFSkipSpeechElementClass {
	return getBFSkipSpeechElementClass()
}

type BFSkipSpeechElementClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (bc BFSkipSpeechElementClass) Class() objc.Class {
	return bc.class
}

// Alloc allocates memory for a new instance of the class.
func (bc BFSkipSpeechElementClass) Alloc() BFSkipSpeechElement {
	rv := objc.SendIfResponds[BFSkipSpeechElement](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

type BFSkipSpeechElement struct {
	BFSpeechElement
}

// BFSkipSpeechElementFromID constructs a [BFSkipSpeechElement] from an objc.ID.
func BFSkipSpeechElementFromID(id objc.ID) BFSkipSpeechElement {
	return BFSkipSpeechElement{BFSpeechElement: BFSpeechElementFromID(id)}
}

// Ensure BFSkipSpeechElement implements IBFSkipSpeechElement.
var _ IBFSkipSpeechElement = BFSkipSpeechElement{}

// An interface definition for the [BFSkipSpeechElement] class.
type IBFSkipSpeechElement interface {
	IBFSpeechElement
}

// Init initializes the instance.
func (b BFSkipSpeechElement) Init() BFSkipSpeechElement {
	rv := objc.SendIfResponds[BFSkipSpeechElement](b.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b BFSkipSpeechElement) Autorelease() BFSkipSpeechElement {
	rv := objc.SendIfResponds[BFSkipSpeechElement](b.ID, objc.Sel("autorelease"))
	return rv
}

// NewBFSkipSpeechElement creates a new BFSkipSpeechElement instance.
func NewBFSkipSpeechElement() BFSkipSpeechElement {
	class := getBFSkipSpeechElementClass()
	rv := objc.SendIfResponds[BFSkipSpeechElement](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewBFSkipSpeechElementWithRange(range_ foundation.NSRange) BFSkipSpeechElement {
	instance := getBFSkipSpeechElementClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithRange:"), range_)
	return BFSkipSpeechElementFromID(rv)
}
