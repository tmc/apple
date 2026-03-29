// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVExposureBiasRange] class.
var (
	_AVExposureBiasRangeClass     AVExposureBiasRangeClass
	_AVExposureBiasRangeClassOnce sync.Once
)

func getAVExposureBiasRangeClass() AVExposureBiasRangeClass {
	_AVExposureBiasRangeClassOnce.Do(func() {
		_AVExposureBiasRangeClass = AVExposureBiasRangeClass{class: objc.GetClass("AVExposureBiasRange")}
	})
	return _AVExposureBiasRangeClass
}

// GetAVExposureBiasRangeClass returns the class object for AVExposureBiasRange.
func GetAVExposureBiasRangeClass() AVExposureBiasRangeClass {
	return getAVExposureBiasRangeClass()
}

type AVExposureBiasRangeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVExposureBiasRangeClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVExposureBiasRangeClass) Alloc() AVExposureBiasRange {
	rv := objc.Send[AVExposureBiasRange](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that expresses an inclusive range of supported exposure bias
// values, in EV units.
//
// # Overview
// 
// A [AVCaptureSystemExposureBiasSlider] defines its range using this type.
//
// # Inspecting the exposure bias range
//
//   - [AVExposureBiasRange.MinExposureBias]: The minimum exposure bias in EV units that this range supports.
//   - [AVExposureBiasRange.MaxExposureBias]: The maximum exposure bias in EV units that this range supports.
//   - [AVExposureBiasRange.ContainsExposureBias]: Determines whether the range contains the specified exposure bias.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVExposureBiasRange
type AVExposureBiasRange struct {
	objectivec.Object
}

// AVExposureBiasRangeFromID constructs a [AVExposureBiasRange] from an objc.ID.
//
// An object that expresses an inclusive range of supported exposure bias
// values, in EV units.
func AVExposureBiasRangeFromID(id objc.ID) AVExposureBiasRange {
	return AVExposureBiasRange{objectivec.Object{ID: id}}
}
// Ensure AVExposureBiasRange implements IAVExposureBiasRange.
var _ IAVExposureBiasRange = AVExposureBiasRange{}

// An interface definition for the [AVExposureBiasRange] class.
//
// # Inspecting the exposure bias range
//
//   - [IAVExposureBiasRange.MinExposureBias]: The minimum exposure bias in EV units that this range supports.
//   - [IAVExposureBiasRange.MaxExposureBias]: The maximum exposure bias in EV units that this range supports.
//   - [IAVExposureBiasRange.ContainsExposureBias]: Determines whether the range contains the specified exposure bias.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVExposureBiasRange
type IAVExposureBiasRange interface {
	objectivec.IObject

	// Topic: Inspecting the exposure bias range

	// The minimum exposure bias in EV units that this range supports.
	MinExposureBias() float32
	// The maximum exposure bias in EV units that this range supports.
	MaxExposureBias() float32
	// Determines whether the range contains the specified exposure bias.
	ContainsExposureBias(exposureBias float32) bool

	// A time value that indicates the maximum supported exposure duration.
	MaxExposureDuration() uintptr
	SetMaxExposureDuration(value uintptr)
	// A floating-point number that indicates the maximum supported exposure ISO value.
	MaxISO() float32
	SetMaxISO(value float32)
	// A time value that indicates the minimum supported exposure duration.
	MinExposureDuration() uintptr
	SetMinExposureDuration(value uintptr)
	// A floating-point number that indicates the minimum supported exposure ISO value.
	MinISO() float32
	SetMinISO(value float32)
}

// Init initializes the instance.
func (e AVExposureBiasRange) Init() AVExposureBiasRange {
	rv := objc.Send[AVExposureBiasRange](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e AVExposureBiasRange) Autorelease() AVExposureBiasRange {
	rv := objc.Send[AVExposureBiasRange](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVExposureBiasRange creates a new AVExposureBiasRange instance.
func NewAVExposureBiasRange() AVExposureBiasRange {
	class := getAVExposureBiasRangeClass()
	rv := objc.Send[AVExposureBiasRange](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Determines whether the range contains the specified exposure bias.
//
// exposureBias: The exposure bias to test, in EV units.
//
// # Return Value
// 
// `true` if the range contains the exposure bias; otherwise, `false`.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVExposureBiasRange/containsExposureBias:
func (e AVExposureBiasRange) ContainsExposureBias(exposureBias float32) bool {
	rv := objc.Send[bool](e.ID, objc.Sel("containsExposureBias:"), exposureBias)
	return rv
}

// The minimum exposure bias in EV units that this range supports.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVExposureBiasRange/minExposureBias
func (e AVExposureBiasRange) MinExposureBias() float32 {
	rv := objc.Send[float32](e.ID, objc.Sel("minExposureBias"))
	return rv
}
// The maximum exposure bias in EV units that this range supports.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVExposureBiasRange/maxExposureBias
func (e AVExposureBiasRange) MaxExposureBias() float32 {
	rv := objc.Send[float32](e.ID, objc.Sel("maxExposureBias"))
	return rv
}
// A time value that indicates the maximum supported exposure duration.
//
// See: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/maxexposureduration
func (e AVExposureBiasRange) MaxExposureDuration() uintptr {
	rv := objc.Send[uintptr](e.ID, objc.Sel("maxExposureDuration"))
	return rv
}
func (e AVExposureBiasRange) SetMaxExposureDuration(value uintptr) {
	objc.Send[struct{}](e.ID, objc.Sel("setMaxExposureDuration:"), value)
}
// A floating-point number that indicates the maximum supported exposure ISO
// value.
//
// See: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/maxiso
func (e AVExposureBiasRange) MaxISO() float32 {
	rv := objc.Send[float32](e.ID, objc.Sel("maxISO"))
	return rv
}
func (e AVExposureBiasRange) SetMaxISO(value float32) {
	objc.Send[struct{}](e.ID, objc.Sel("setMaxISO:"), value)
}
// A time value that indicates the minimum supported exposure duration.
//
// See: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/minexposureduration
func (e AVExposureBiasRange) MinExposureDuration() uintptr {
	rv := objc.Send[uintptr](e.ID, objc.Sel("minExposureDuration"))
	return rv
}
func (e AVExposureBiasRange) SetMinExposureDuration(value uintptr) {
	objc.Send[struct{}](e.ID, objc.Sel("setMinExposureDuration:"), value)
}
// A floating-point number that indicates the minimum supported exposure ISO
// value.
//
// See: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/miniso
func (e AVExposureBiasRange) MinISO() float32 {
	rv := objc.Send[float32](e.ID, objc.Sel("minISO"))
	return rv
}
func (e AVExposureBiasRange) SetMinISO(value float32) {
	objc.Send[struct{}](e.ID, objc.Sel("setMinISO:"), value)
}

