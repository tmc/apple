// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"

	"github.com/tmc/apple/coremedia"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVCompositionTrackFormatDescriptionReplacement] class.
var (
	_AVCompositionTrackFormatDescriptionReplacementClass     AVCompositionTrackFormatDescriptionReplacementClass
	_AVCompositionTrackFormatDescriptionReplacementClassOnce sync.Once
)

func getAVCompositionTrackFormatDescriptionReplacementClass() AVCompositionTrackFormatDescriptionReplacementClass {
	_AVCompositionTrackFormatDescriptionReplacementClassOnce.Do(func() {
		_AVCompositionTrackFormatDescriptionReplacementClass = AVCompositionTrackFormatDescriptionReplacementClass{class: objc.GetClass("AVCompositionTrackFormatDescriptionReplacement")}
	})
	return _AVCompositionTrackFormatDescriptionReplacementClass
}

// GetAVCompositionTrackFormatDescriptionReplacementClass returns the class object for AVCompositionTrackFormatDescriptionReplacement.
func GetAVCompositionTrackFormatDescriptionReplacementClass() AVCompositionTrackFormatDescriptionReplacementClass {
	return getAVCompositionTrackFormatDescriptionReplacementClass()
}

type AVCompositionTrackFormatDescriptionReplacementClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVCompositionTrackFormatDescriptionReplacementClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVCompositionTrackFormatDescriptionReplacementClass) Alloc() AVCompositionTrackFormatDescriptionReplacement {
	rv := objc.Send[AVCompositionTrackFormatDescriptionReplacement](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that represents a format description and its replacement.
//
// # Managing format descriptions
//
//   - [AVCompositionTrackFormatDescriptionReplacement.OriginalFormatDescription]: The format description to replace.
//   - [AVCompositionTrackFormatDescriptionReplacement.ReplacementFormatDescription]: The replacement format description.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrackFormatDescriptionReplacement
type AVCompositionTrackFormatDescriptionReplacement struct {
	objectivec.Object
}

// AVCompositionTrackFormatDescriptionReplacementFromID constructs a [AVCompositionTrackFormatDescriptionReplacement] from an objc.ID.
//
// An object that represents a format description and its replacement.
func AVCompositionTrackFormatDescriptionReplacementFromID(id objc.ID) AVCompositionTrackFormatDescriptionReplacement {
	return AVCompositionTrackFormatDescriptionReplacement{objectivec.Object{ID: id}}
}

// NOTE: AVCompositionTrackFormatDescriptionReplacement adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVCompositionTrackFormatDescriptionReplacement] class.
//
// # Managing format descriptions
//
//   - [IAVCompositionTrackFormatDescriptionReplacement.OriginalFormatDescription]: The format description to replace.
//   - [IAVCompositionTrackFormatDescriptionReplacement.ReplacementFormatDescription]: The replacement format description.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrackFormatDescriptionReplacement
type IAVCompositionTrackFormatDescriptionReplacement interface {
	objectivec.IObject

	// Topic: Managing format descriptions

	// The format description to replace.
	OriginalFormatDescription() coremedia.CMFormatDescriptionRef
	// The replacement format description.
	ReplacementFormatDescription() coremedia.CMFormatDescriptionRef

	InitWithCoder(coder foundation.INSCoder) AVCompositionTrackFormatDescriptionReplacement
	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (c AVCompositionTrackFormatDescriptionReplacement) Init() AVCompositionTrackFormatDescriptionReplacement {
	rv := objc.Send[AVCompositionTrackFormatDescriptionReplacement](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c AVCompositionTrackFormatDescriptionReplacement) Autorelease() AVCompositionTrackFormatDescriptionReplacement {
	rv := objc.Send[AVCompositionTrackFormatDescriptionReplacement](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVCompositionTrackFormatDescriptionReplacement creates a new AVCompositionTrackFormatDescriptionReplacement instance.
func NewAVCompositionTrackFormatDescriptionReplacement() AVCompositionTrackFormatDescriptionReplacement {
	class := getAVCompositionTrackFormatDescriptionReplacementClass()
	rv := objc.Send[AVCompositionTrackFormatDescriptionReplacement](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrackFormatDescriptionReplacement/init(coder:)
func NewCompositionTrackFormatDescriptionReplacementWithCoder(coder foundation.INSCoder) AVCompositionTrackFormatDescriptionReplacement {
	instance := getAVCompositionTrackFormatDescriptionReplacementClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return AVCompositionTrackFormatDescriptionReplacementFromID(rv)
}

// See: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrackFormatDescriptionReplacement/init(coder:)
func (c AVCompositionTrackFormatDescriptionReplacement) InitWithCoder(coder foundation.INSCoder) AVCompositionTrackFormatDescriptionReplacement {
	rv := objc.Send[AVCompositionTrackFormatDescriptionReplacement](c.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}
func (c AVCompositionTrackFormatDescriptionReplacement) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](c.ID, objc.Sel("encodeWithCoder:"), coder)
}

// The format description to replace.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrackFormatDescriptionReplacement/originalFormatDescription
func (c AVCompositionTrackFormatDescriptionReplacement) OriginalFormatDescription() coremedia.CMFormatDescriptionRef {
	rv := objc.Send[coremedia.CMFormatDescriptionRef](c.ID, objc.Sel("originalFormatDescription"))
	return coremedia.CMFormatDescriptionRef(rv)
}

// The replacement format description.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrackFormatDescriptionReplacement/replacementFormatDescription
func (c AVCompositionTrackFormatDescriptionReplacement) ReplacementFormatDescription() coremedia.CMFormatDescriptionRef {
	rv := objc.Send[coremedia.CMFormatDescriptionRef](c.ID, objc.Sel("replacementFormatDescription"))
	return coremedia.CMFormatDescriptionRef(rv)
}
