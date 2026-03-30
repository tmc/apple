// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVMetricMediaRendition] class.
var (
	_AVMetricMediaRenditionClass     AVMetricMediaRenditionClass
	_AVMetricMediaRenditionClassOnce sync.Once
)

func getAVMetricMediaRenditionClass() AVMetricMediaRenditionClass {
	_AVMetricMediaRenditionClassOnce.Do(func() {
		_AVMetricMediaRenditionClass = AVMetricMediaRenditionClass{class: objc.GetClass("AVMetricMediaRendition")}
	})
	return _AVMetricMediaRenditionClass
}

// GetAVMetricMediaRenditionClass returns the class object for AVMetricMediaRendition.
func GetAVMetricMediaRenditionClass() AVMetricMediaRenditionClass {
	return getAVMetricMediaRenditionClass()
}

type AVMetricMediaRenditionClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVMetricMediaRenditionClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVMetricMediaRenditionClass) Alloc() AVMetricMediaRendition {
	rv := objc.Send[AVMetricMediaRendition](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// # Inspecting the rendition
//
//   - [AVMetricMediaRendition.StableID]
//   - [AVMetricMediaRendition.URL]
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMetricMediaRendition
type AVMetricMediaRendition struct {
	objectivec.Object
}

// AVMetricMediaRenditionFromID constructs a [AVMetricMediaRendition] from an objc.ID.
func AVMetricMediaRenditionFromID(id objc.ID) AVMetricMediaRendition {
	return AVMetricMediaRendition{objectivec.Object{ID: id}}
}

// NOTE: AVMetricMediaRendition adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVMetricMediaRendition] class.
//
// # Inspecting the rendition
//
//   - [IAVMetricMediaRendition.StableID]
//   - [IAVMetricMediaRendition.URL]
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMetricMediaRendition
type IAVMetricMediaRendition interface {
	objectivec.IObject

	// Topic: Inspecting the rendition

	StableID() string
	URL() foundation.INSURL

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (m AVMetricMediaRendition) Init() AVMetricMediaRendition {
	rv := objc.Send[AVMetricMediaRendition](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m AVMetricMediaRendition) Autorelease() AVMetricMediaRendition {
	rv := objc.Send[AVMetricMediaRendition](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVMetricMediaRendition creates a new AVMetricMediaRendition instance.
func NewAVMetricMediaRendition() AVMetricMediaRendition {
	class := getAVMetricMediaRenditionClass()
	rv := objc.Send[AVMetricMediaRendition](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (m AVMetricMediaRendition) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](m.ID, objc.Sel("encodeWithCoder:"), coder)
}

// # Discussion
//
// Provides ID corresponding to the rendition. This is equivalent to the
// STABLE-RENDITION-ID in the HLS playlist. If not available, value is nil.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMetricMediaRendition/stableID
func (m AVMetricMediaRendition) StableID() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("stableID"))
	return foundation.NSStringFromID(rv).String()
}

// # Discussion
//
// Provides URL corresponding to the rendition’s HLS playlist. If not
// available, value is nil.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMetricMediaRendition/url
func (m AVMetricMediaRendition) URL() foundation.INSURL {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("URL"))
	return foundation.NSURLFromID(objc.ID(rv))
}
