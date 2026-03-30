// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVAudioUnitSplitter] class.
var (
	_AVAudioUnitSplitterClass     AVAudioUnitSplitterClass
	_AVAudioUnitSplitterClassOnce sync.Once
)

func getAVAudioUnitSplitterClass() AVAudioUnitSplitterClass {
	_AVAudioUnitSplitterClassOnce.Do(func() {
		_AVAudioUnitSplitterClass = AVAudioUnitSplitterClass{class: objc.GetClass("AVAudioUnitSplitter")}
	})
	return _AVAudioUnitSplitterClass
}

// GetAVAudioUnitSplitterClass returns the class object for AVAudioUnitSplitter.
func GetAVAudioUnitSplitterClass() AVAudioUnitSplitterClass {
	return getAVAudioUnitSplitterClass()
}

type AVAudioUnitSplitterClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVAudioUnitSplitterClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVAudioUnitSplitterClass) Alloc() AVAudioUnitSplitter {
	rv := objc.Send[AVAudioUnitSplitter](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [AVAudioUnitSplitter.InitWithAudioComponentDescription]
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitSplitter
type AVAudioUnitSplitter struct {
	AVAudioUnit
}

// AVAudioUnitSplitterFromID constructs a [AVAudioUnitSplitter] from an objc.ID.
func AVAudioUnitSplitterFromID(id objc.ID) AVAudioUnitSplitter {
	return AVAudioUnitSplitter{AVAudioUnit: AVAudioUnitFromID(id)}
}

// Ensure AVAudioUnitSplitter implements IAVAudioUnitSplitter.
var _ IAVAudioUnitSplitter = AVAudioUnitSplitter{}

// An interface definition for the [AVAudioUnitSplitter] class.
//
// # Methods
//
//   - [IAVAudioUnitSplitter.InitWithAudioComponentDescription]
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitSplitter
type IAVAudioUnitSplitter interface {
	IAVAudioUnit

	// Topic: Methods

	InitWithAudioComponentDescription(description objectivec.IObject) AVAudioUnitSplitter
}

// Init initializes the instance.
func (a AVAudioUnitSplitter) Init() AVAudioUnitSplitter {
	rv := objc.Send[AVAudioUnitSplitter](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AVAudioUnitSplitter) Autorelease() AVAudioUnitSplitter {
	rv := objc.Send[AVAudioUnitSplitter](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVAudioUnitSplitter creates a new AVAudioUnitSplitter instance.
func NewAVAudioUnitSplitter() AVAudioUnitSplitter {
	class := getAVAudioUnitSplitterClass()
	rv := objc.Send[AVAudioUnitSplitter](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitSplitter/initWithAudioComponentDescription:
func NewAudioUnitSplitterWithAudioComponentDescription(description objectivec.IObject) AVAudioUnitSplitter {
	instance := getAVAudioUnitSplitterClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithAudioComponentDescription:"), description)
	return AVAudioUnitSplitterFromID(rv)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioNode/initWithImpl:
func NewAudioUnitSplitterWithImpl(impl unsafe.Pointer) AVAudioUnitSplitter {
	instance := getAVAudioUnitSplitterClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithImpl:"), impl)
	return AVAudioUnitSplitterFromID(rv)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitSplitter/initWithAudioComponentDescription:
func (a AVAudioUnitSplitter) InitWithAudioComponentDescription(description objectivec.IObject) AVAudioUnitSplitter {
	rv := objc.Send[AVAudioUnitSplitter](a.ID, objc.Sel("initWithAudioComponentDescription:"), description)
	return rv
}
