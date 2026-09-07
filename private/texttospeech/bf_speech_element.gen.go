// Code generated from Apple documentation for texttospeech. DO NOT EDIT.

package texttospeech

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [BFSpeechElement] class.
var (
	_BFSpeechElementClass     BFSpeechElementClass
	_BFSpeechElementClassOnce sync.Once
)

func getBFSpeechElementClass() BFSpeechElementClass {
	_BFSpeechElementClassOnce.Do(func() {
		_BFSpeechElementClass = BFSpeechElementClass{class: objc.GetClass("BFSpeechElement")}
	})
	return _BFSpeechElementClass
}

// GetBFSpeechElementClass returns the class object for BFSpeechElement.
func GetBFSpeechElementClass() BFSpeechElementClass {
	return getBFSpeechElementClass()
}

type BFSpeechElementClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (bc BFSpeechElementClass) Class() objc.Class {
	return bc.class
}

// Alloc allocates memory for a new instance of the class.
func (bc BFSpeechElementClass) Alloc() BFSpeechElement {
	rv := objc.SendIfResponds[BFSpeechElement](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [BFSpeechElement.OriginalRange]
//   - [BFSpeechElement.SetOriginalRange]
//   - [BFSpeechElement.InitWithRange]
type BFSpeechElement struct {
	objectivec.Object
}

// BFSpeechElementFromID constructs a [BFSpeechElement] from an objc.ID.
func BFSpeechElementFromID(id objc.ID) BFSpeechElement {
	return BFSpeechElement{objectivec.Object{ID: id}}
}

// Ensure BFSpeechElement implements IBFSpeechElement.
var _ IBFSpeechElement = BFSpeechElement{}

// An interface definition for the [BFSpeechElement] class.
//
// # Methods
//
//   - [IBFSpeechElement.OriginalRange]
//   - [IBFSpeechElement.SetOriginalRange]
//   - [IBFSpeechElement.InitWithRange]
type IBFSpeechElement interface {
	objectivec.IObject

	// Topic: Methods

	OriginalRange() foundation.NSRange
	SetOriginalRange(value foundation.NSRange)
	InitWithRange(range_ foundation.NSRange) BFSpeechElement
}

// Init initializes the instance.
func (b BFSpeechElement) Init() BFSpeechElement {
	rv := objc.SendIfResponds[BFSpeechElement](b.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b BFSpeechElement) Autorelease() BFSpeechElement {
	rv := objc.SendIfResponds[BFSpeechElement](b.ID, objc.Sel("autorelease"))
	return rv
}

// NewBFSpeechElement creates a new BFSpeechElement instance.
func NewBFSpeechElement() BFSpeechElement {
	class := getBFSpeechElementClass()
	rv := objc.SendIfResponds[BFSpeechElement](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewBFSpeechElementWithRange(range_ foundation.NSRange) BFSpeechElement {
	instance := getBFSpeechElementClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithRange:"), range_)
	return BFSpeechElementFromID(rv)
}

func (b BFSpeechElement) InitWithRange(range_ foundation.NSRange) BFSpeechElement {
	rv := objc.SendIfResponds[BFSpeechElement](b.ID, objc.Sel("initWithRange:"), range_)
	return rv
}

func (b BFSpeechElement) OriginalRange() foundation.NSRange {
	rv := objc.SendIfResponds[foundation.NSRange](b.ID, objc.Sel("originalRange"))
	return foundation.NSRange(rv)
}
func (b BFSpeechElement) SetOriginalRange(value foundation.NSRange) {
	objc.SendIfResponds[struct{}](b.ID, objc.Sel("setOriginalRange:"), value)
}
