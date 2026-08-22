// Code generated from Apple documentation for Accessibility. DO NOT EDIT.

package accessibility

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AXLiveAudioGraph] class.
var (
	_AXLiveAudioGraphClass     AXLiveAudioGraphClass
	_AXLiveAudioGraphClassOnce sync.Once
)

func getAXLiveAudioGraphClass() AXLiveAudioGraphClass {
	_AXLiveAudioGraphClassOnce.Do(func() {
		_AXLiveAudioGraphClass = AXLiveAudioGraphClass{class: objc.GetClass("AXLiveAudioGraph")}
	})
	return _AXLiveAudioGraphClass
}

// GetAXLiveAudioGraphClass returns the class object for AXLiveAudioGraph.
func GetAXLiveAudioGraphClass() AXLiveAudioGraphClass {
	return getAXLiveAudioGraphClass()
}

type AXLiveAudioGraphClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AXLiveAudioGraphClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AXLiveAudioGraphClass) Alloc() AXLiveAudioGraph {
	rv := objc.Send[AXLiveAudioGraph](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that represents an audio graph for a live-updating, continuous
// data series for VoiceOver.
//
// # Overview
//
// Use [AXLiveAudioGraph] to interact with an ongoing, continuous stream of
// data that updates with new data in real time.
//
// See: https://developer.apple.com/documentation/Accessibility/AXLiveAudioGraph
type AXLiveAudioGraph struct {
	objectivec.Object
}

// AXLiveAudioGraphFromID constructs a [AXLiveAudioGraph] from an objc.ID.
//
// An object that represents an audio graph for a live-updating, continuous
// data series for VoiceOver.
func AXLiveAudioGraphFromID(id objc.ID) AXLiveAudioGraph {
	return AXLiveAudioGraph{objectivec.Object{ID: id}}
}

// NOTE: AXLiveAudioGraph adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AXLiveAudioGraph] class.
//
// See: https://developer.apple.com/documentation/Accessibility/AXLiveAudioGraph
type IAXLiveAudioGraph interface {
	objectivec.IObject
}

// Init initializes the instance.
func (a AXLiveAudioGraph) Init() AXLiveAudioGraph {
	rv := objc.Send[AXLiveAudioGraph](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AXLiveAudioGraph) Autorelease() AXLiveAudioGraph {
	rv := objc.Send[AXLiveAudioGraph](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAXLiveAudioGraph creates a new AXLiveAudioGraph instance.
func NewAXLiveAudioGraph() AXLiveAudioGraph {
	class := getAXLiveAudioGraphClass()
	rv := objc.Send[AXLiveAudioGraph](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Begins the live audio graph session.
//
// See: https://developer.apple.com/documentation/Accessibility/AXLiveAudioGraph/start()
func (_AXLiveAudioGraphClass AXLiveAudioGraphClass) Start() {
	objc.Send[objc.ID](objc.ID(_AXLiveAudioGraphClass.class), objc.Sel("start"))
}

// Ends the live audio graph session.
//
// See: https://developer.apple.com/documentation/Accessibility/AXLiveAudioGraph/stop()
func (_AXLiveAudioGraphClass AXLiveAudioGraphClass) Stop() {
	objc.Send[objc.ID](objc.ID(_AXLiveAudioGraphClass.class), objc.Sel("stop"))
}

// Sets the pitch of the audio graph’s tone.
//
// value: A normalized value in the range [`0.0`, `1.0`], where `0.0` represents the
// minimum displayable y-axis value for your data series, and `1.0` represents
// the maximum displayable y-axis value for your data series.
//
// See: https://developer.apple.com/documentation/Accessibility/AXLiveAudioGraph/updateValue(_:)
func (_AXLiveAudioGraphClass AXLiveAudioGraphClass) UpdateValue(value float64) {
	objc.Send[objc.ID](objc.ID(_AXLiveAudioGraphClass.class), objc.Sel("updateValue:"), value)
}
