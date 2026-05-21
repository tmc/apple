// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVSpeechSynthesisMarker] class.
var (
	_AVSpeechSynthesisMarkerClass     AVSpeechSynthesisMarkerClass
	_AVSpeechSynthesisMarkerClassOnce sync.Once
)

func getAVSpeechSynthesisMarkerClass() AVSpeechSynthesisMarkerClass {
	_AVSpeechSynthesisMarkerClassOnce.Do(func() {
		_AVSpeechSynthesisMarkerClass = AVSpeechSynthesisMarkerClass{class: objc.GetClass("AVSpeechSynthesisMarker")}
	})
	return _AVSpeechSynthesisMarkerClass
}

// GetAVSpeechSynthesisMarkerClass returns the class object for AVSpeechSynthesisMarker.
func GetAVSpeechSynthesisMarkerClass() AVSpeechSynthesisMarkerClass {
	return getAVSpeechSynthesisMarkerClass()
}

type AVSpeechSynthesisMarkerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVSpeechSynthesisMarkerClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVSpeechSynthesisMarkerClass) Alloc() AVSpeechSynthesisMarker {
	rv := objc.Send[AVSpeechSynthesisMarker](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

type AVSpeechSynthesisMarker struct {
	objectivec.Object
}

// AVSpeechSynthesisMarkerFromID constructs a [AVSpeechSynthesisMarker] from an objc.ID.
func AVSpeechSynthesisMarkerFromID(id objc.ID) AVSpeechSynthesisMarker {
	return AVSpeechSynthesisMarker{objectivec.Object{ID: id}}
}

// Ensure AVSpeechSynthesisMarker implements IAVSpeechSynthesisMarker.
var _ IAVSpeechSynthesisMarker = AVSpeechSynthesisMarker{}

// An interface definition for the [AVSpeechSynthesisMarker] class.
type IAVSpeechSynthesisMarker interface {
	objectivec.IObject
}

// Init initializes the instance.
func (a AVSpeechSynthesisMarker) Init() AVSpeechSynthesisMarker {
	rv := objc.Send[AVSpeechSynthesisMarker](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AVSpeechSynthesisMarker) Autorelease() AVSpeechSynthesisMarker {
	rv := objc.Send[AVSpeechSynthesisMarker](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVSpeechSynthesisMarker creates a new AVSpeechSynthesisMarker instance.
func NewAVSpeechSynthesisMarker() AVSpeechSynthesisMarker {
	class := getAVSpeechSynthesisMarkerClass()
	rv := objc.Send[AVSpeechSynthesisMarker](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (_AVSpeechSynthesisMarkerClass AVSpeechSynthesisMarkerClass) SupportsSecureCoding() bool {
	rv := objc.Send[bool](objc.ID(_AVSpeechSynthesisMarkerClass.class), objc.Sel("supportsSecureCoding"))
	return rv
}
