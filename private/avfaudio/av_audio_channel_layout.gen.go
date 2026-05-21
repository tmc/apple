// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVAudioChannelLayout] class.
var (
	_AVAudioChannelLayoutClass     AVAudioChannelLayoutClass
	_AVAudioChannelLayoutClassOnce sync.Once
)

func getAVAudioChannelLayoutClass() AVAudioChannelLayoutClass {
	_AVAudioChannelLayoutClassOnce.Do(func() {
		_AVAudioChannelLayoutClass = AVAudioChannelLayoutClass{class: objc.GetClass("AVAudioChannelLayout")}
	})
	return _AVAudioChannelLayoutClass
}

// GetAVAudioChannelLayoutClass returns the class object for AVAudioChannelLayout.
func GetAVAudioChannelLayoutClass() AVAudioChannelLayoutClass {
	return getAVAudioChannelLayoutClass()
}

type AVAudioChannelLayoutClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVAudioChannelLayoutClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVAudioChannelLayoutClass) Alloc() AVAudioChannelLayout {
	rv := objc.Send[AVAudioChannelLayout](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [AVAudioChannelLayout.LayoutSize]
type AVAudioChannelLayout struct {
	objectivec.Object
}

// AVAudioChannelLayoutFromID constructs a [AVAudioChannelLayout] from an objc.ID.
func AVAudioChannelLayoutFromID(id objc.ID) AVAudioChannelLayout {
	return AVAudioChannelLayout{objectivec.Object{ID: id}}
}

// Ensure AVAudioChannelLayout implements IAVAudioChannelLayout.
var _ IAVAudioChannelLayout = AVAudioChannelLayout{}

// An interface definition for the [AVAudioChannelLayout] class.
//
// # Methods
//
//   - [IAVAudioChannelLayout.LayoutSize]
type IAVAudioChannelLayout interface {
	objectivec.IObject

	// Topic: Methods

	LayoutSize() uint64
}

// Init initializes the instance.
func (a AVAudioChannelLayout) Init() AVAudioChannelLayout {
	rv := objc.Send[AVAudioChannelLayout](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AVAudioChannelLayout) Autorelease() AVAudioChannelLayout {
	rv := objc.Send[AVAudioChannelLayout](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVAudioChannelLayout creates a new AVAudioChannelLayout instance.
func NewAVAudioChannelLayout() AVAudioChannelLayout {
	class := getAVAudioChannelLayoutClass()
	rv := objc.Send[AVAudioChannelLayout](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (a AVAudioChannelLayout) LayoutSize() uint64 {
	rv := objc.Send[uint64](a.ID, objc.Sel("layoutSize"))
	return rv
}

func (_AVAudioChannelLayoutClass AVAudioChannelLayoutClass) SupportsSecureCoding() bool {
	rv := objc.Send[bool](objc.ID(_AVAudioChannelLayoutClass.class), objc.Sel("supportsSecureCoding"))
	return rv
}
