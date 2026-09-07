// Code generated from Apple documentation for texttospeech. DO NOT EDIT.

package texttospeech

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [BFSyncSpeechElement] class.
var (
	_BFSyncSpeechElementClass     BFSyncSpeechElementClass
	_BFSyncSpeechElementClassOnce sync.Once
)

func getBFSyncSpeechElementClass() BFSyncSpeechElementClass {
	_BFSyncSpeechElementClassOnce.Do(func() {
		_BFSyncSpeechElementClass = BFSyncSpeechElementClass{class: objc.GetClass("BFSyncSpeechElement")}
	})
	return _BFSyncSpeechElementClass
}

// GetBFSyncSpeechElementClass returns the class object for BFSyncSpeechElement.
func GetBFSyncSpeechElementClass() BFSyncSpeechElementClass {
	return getBFSyncSpeechElementClass()
}

type BFSyncSpeechElementClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (bc BFSyncSpeechElementClass) Class() objc.Class {
	return bc.class
}

// Alloc allocates memory for a new instance of the class.
func (bc BFSyncSpeechElementClass) Alloc() BFSyncSpeechElement {
	rv := objc.SendIfResponds[BFSyncSpeechElement](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [BFSyncSpeechElement.Name]
//   - [BFSyncSpeechElement.SetName]
//   - [BFSyncSpeechElement.InitWithNameAndRange]
type BFSyncSpeechElement struct {
	BFSpeechElement
}

// BFSyncSpeechElementFromID constructs a [BFSyncSpeechElement] from an objc.ID.
func BFSyncSpeechElementFromID(id objc.ID) BFSyncSpeechElement {
	return BFSyncSpeechElement{BFSpeechElement: BFSpeechElementFromID(id)}
}

// Ensure BFSyncSpeechElement implements IBFSyncSpeechElement.
var _ IBFSyncSpeechElement = BFSyncSpeechElement{}

// An interface definition for the [BFSyncSpeechElement] class.
//
// # Methods
//
//   - [IBFSyncSpeechElement.Name]
//   - [IBFSyncSpeechElement.SetName]
//   - [IBFSyncSpeechElement.InitWithNameAndRange]
type IBFSyncSpeechElement interface {
	IBFSpeechElement

	// Topic: Methods

	Name() string
	SetName(value string)
	InitWithNameAndRange(name objectivec.IObject, range_ foundation.NSRange) BFSyncSpeechElement
}

// Init initializes the instance.
func (b BFSyncSpeechElement) Init() BFSyncSpeechElement {
	rv := objc.SendIfResponds[BFSyncSpeechElement](b.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b BFSyncSpeechElement) Autorelease() BFSyncSpeechElement {
	rv := objc.SendIfResponds[BFSyncSpeechElement](b.ID, objc.Sel("autorelease"))
	return rv
}

// NewBFSyncSpeechElement creates a new BFSyncSpeechElement instance.
func NewBFSyncSpeechElement() BFSyncSpeechElement {
	class := getBFSyncSpeechElementClass()
	rv := objc.SendIfResponds[BFSyncSpeechElement](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewBFSyncSpeechElementWithNameAndRange(name objectivec.IObject, range_ foundation.NSRange) BFSyncSpeechElement {
	instance := getBFSyncSpeechElementClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithName:andRange:"), name, range_)
	return BFSyncSpeechElementFromID(rv)
}

func NewBFSyncSpeechElementWithRange(range_ foundation.NSRange) BFSyncSpeechElement {
	instance := getBFSyncSpeechElementClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithRange:"), range_)
	return BFSyncSpeechElementFromID(rv)
}

func (b BFSyncSpeechElement) InitWithNameAndRange(name objectivec.IObject, range_ foundation.NSRange) BFSyncSpeechElement {
	rv := objc.SendIfResponds[BFSyncSpeechElement](b.ID, objc.Sel("initWithName:andRange:"), name, range_)
	return rv
}

func (b BFSyncSpeechElement) Name() string {
	rv := objc.SendIfResponds[objc.ID](b.ID, objc.Sel("name"))
	return foundation.NSStringFromID(rv).String()
}
func (b BFSyncSpeechElement) SetName(value string) {
	objc.SendIfResponds[struct{}](b.ID, objc.Sel("setName:"), objc.String(value))
}
