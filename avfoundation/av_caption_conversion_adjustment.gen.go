// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVCaptionConversionAdjustment] class.
var (
	_AVCaptionConversionAdjustmentClass     AVCaptionConversionAdjustmentClass
	_AVCaptionConversionAdjustmentClassOnce sync.Once
)

func getAVCaptionConversionAdjustmentClass() AVCaptionConversionAdjustmentClass {
	_AVCaptionConversionAdjustmentClassOnce.Do(func() {
		_AVCaptionConversionAdjustmentClass = AVCaptionConversionAdjustmentClass{class: objc.GetClass("AVCaptionConversionAdjustment")}
	})
	return _AVCaptionConversionAdjustmentClass
}

// GetAVCaptionConversionAdjustmentClass returns the class object for AVCaptionConversionAdjustment.
func GetAVCaptionConversionAdjustmentClass() AVCaptionConversionAdjustmentClass {
	return getAVCaptionConversionAdjustmentClass()
}

type AVCaptionConversionAdjustmentClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVCaptionConversionAdjustmentClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVCaptionConversionAdjustmentClass) Alloc() AVCaptionConversionAdjustment {
	rv := objc.Send[AVCaptionConversionAdjustment](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that describes an adjustment to correct a problem found during
// validation of a caption conversion.
//
// # Accessing the adjustment type
//
//   - [AVCaptionConversionAdjustment.AdjustmentType]: The type of caption conversion adjustment.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptionConversionAdjustment
type AVCaptionConversionAdjustment struct {
	objectivec.Object
}

// AVCaptionConversionAdjustmentFromID constructs a [AVCaptionConversionAdjustment] from an objc.ID.
//
// An object that describes an adjustment to correct a problem found during
// validation of a caption conversion.
func AVCaptionConversionAdjustmentFromID(id objc.ID) AVCaptionConversionAdjustment {
	return AVCaptionConversionAdjustment{objectivec.Object{ID: id}}
}

// NOTE: AVCaptionConversionAdjustment adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVCaptionConversionAdjustment] class.
//
// # Accessing the adjustment type
//
//   - [IAVCaptionConversionAdjustment.AdjustmentType]: The type of caption conversion adjustment.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptionConversionAdjustment
type IAVCaptionConversionAdjustment interface {
	objectivec.IObject

	// Topic: Accessing the adjustment type

	// The type of caption conversion adjustment.
	AdjustmentType() AVCaptionConversionAdjustmentType
}

// Init initializes the instance.
func (c AVCaptionConversionAdjustment) Init() AVCaptionConversionAdjustment {
	rv := objc.Send[AVCaptionConversionAdjustment](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c AVCaptionConversionAdjustment) Autorelease() AVCaptionConversionAdjustment {
	rv := objc.Send[AVCaptionConversionAdjustment](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVCaptionConversionAdjustment creates a new AVCaptionConversionAdjustment instance.
func NewAVCaptionConversionAdjustment() AVCaptionConversionAdjustment {
	class := getAVCaptionConversionAdjustmentClass()
	rv := objc.Send[AVCaptionConversionAdjustment](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The type of caption conversion adjustment.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptionConversionAdjustment/adjustmentType-swift.property
func (c AVCaptionConversionAdjustment) AdjustmentType() AVCaptionConversionAdjustmentType {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("adjustmentType"))
	return AVCaptionConversionAdjustmentType(foundation.NSStringFromID(rv).String())
}
