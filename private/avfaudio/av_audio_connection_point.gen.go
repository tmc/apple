// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVAudioConnectionPoint] class.
var (
	_AVAudioConnectionPointClass     AVAudioConnectionPointClass
	_AVAudioConnectionPointClassOnce sync.Once
)

func getAVAudioConnectionPointClass() AVAudioConnectionPointClass {
	_AVAudioConnectionPointClassOnce.Do(func() {
		_AVAudioConnectionPointClass = AVAudioConnectionPointClass{class: objc.GetClass("AVAudioConnectionPoint")}
	})
	return _AVAudioConnectionPointClass
}

// GetAVAudioConnectionPointClass returns the class object for AVAudioConnectionPoint.
func GetAVAudioConnectionPointClass() AVAudioConnectionPointClass {
	return getAVAudioConnectionPointClass()
}

type AVAudioConnectionPointClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVAudioConnectionPointClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVAudioConnectionPointClass) Alloc() AVAudioConnectionPoint {
	rv := objc.Send[AVAudioConnectionPoint](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioConnectionPoint
type AVAudioConnectionPoint struct {
	objectivec.Object
}

// AVAudioConnectionPointFromID constructs a [AVAudioConnectionPoint] from an objc.ID.
func AVAudioConnectionPointFromID(id objc.ID) AVAudioConnectionPoint {
	return AVAudioConnectionPoint{objectivec.Object{ID: id}}
}
// Ensure AVAudioConnectionPoint implements IAVAudioConnectionPoint.
var _ IAVAudioConnectionPoint = AVAudioConnectionPoint{}

// An interface definition for the [AVAudioConnectionPoint] class.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioConnectionPoint
type IAVAudioConnectionPoint interface {
	objectivec.IObject
}

// Init initializes the instance.
func (a AVAudioConnectionPoint) Init() AVAudioConnectionPoint {
	rv := objc.Send[AVAudioConnectionPoint](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AVAudioConnectionPoint) Autorelease() AVAudioConnectionPoint {
	rv := objc.Send[AVAudioConnectionPoint](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVAudioConnectionPoint creates a new AVAudioConnectionPoint instance.
func NewAVAudioConnectionPoint() AVAudioConnectionPoint {
	class := getAVAudioConnectionPointClass()
	rv := objc.Send[AVAudioConnectionPoint](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioConnectionPoint/connectionPointWithNode:bus:
func (_AVAudioConnectionPointClass AVAudioConnectionPointClass) ConnectionPointWithNodeBus(node objectivec.IObject, bus uint64) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_AVAudioConnectionPointClass.class), objc.Sel("connectionPointWithNode:bus:"), node, bus)
	return objectivec.Object{ID: rv}
}

