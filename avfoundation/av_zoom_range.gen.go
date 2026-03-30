// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVZoomRange] class.
var (
	_AVZoomRangeClass     AVZoomRangeClass
	_AVZoomRangeClassOnce sync.Once
)

func getAVZoomRangeClass() AVZoomRangeClass {
	_AVZoomRangeClassOnce.Do(func() {
		_AVZoomRangeClass = AVZoomRangeClass{class: objc.GetClass("AVZoomRange")}
	})
	return _AVZoomRangeClass
}

// GetAVZoomRangeClass returns the class object for AVZoomRange.
func GetAVZoomRangeClass() AVZoomRangeClass {
	return getAVZoomRangeClass()
}

type AVZoomRangeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVZoomRangeClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVZoomRangeClass) Alloc() AVZoomRange {
	rv := objc.Send[AVZoomRange](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that defines an inclusive range of zoom values.
//
// # Inspecting a range
//
//   - [AVZoomRange.MinZoomFactor]: The range’s minimum zoom factor.
//   - [AVZoomRange.MaxZoomFactor]: The range’s maximum zoom factor.
//   - [AVZoomRange.ContainsZoomFactor]: Returns a Boolean value that indicates whether the specified zoom factor exists in the range.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVZoomRange
type AVZoomRange struct {
	objectivec.Object
}

// AVZoomRangeFromID constructs a [AVZoomRange] from an objc.ID.
//
// An object that defines an inclusive range of zoom values.
func AVZoomRangeFromID(id objc.ID) AVZoomRange {
	return AVZoomRange{objectivec.Object{ID: id}}
}

// Ensure AVZoomRange implements IAVZoomRange.
var _ IAVZoomRange = AVZoomRange{}

// An interface definition for the [AVZoomRange] class.
//
// # Inspecting a range
//
//   - [IAVZoomRange.MinZoomFactor]: The range’s minimum zoom factor.
//   - [IAVZoomRange.MaxZoomFactor]: The range’s maximum zoom factor.
//   - [IAVZoomRange.ContainsZoomFactor]: Returns a Boolean value that indicates whether the specified zoom factor exists in the range.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVZoomRange
type IAVZoomRange interface {
	objectivec.IObject

	// Topic: Inspecting a range

	// The range’s minimum zoom factor.
	MinZoomFactor() float64
	// The range’s maximum zoom factor.
	MaxZoomFactor() float64
	// Returns a Boolean value that indicates whether the specified zoom factor exists in the range.
	ContainsZoomFactor(zoomFactor float64) bool

	// A maximum zoom factor the format allows.
	VideoMaxZoomFactor() float64
	SetVideoMaxZoomFactor(value float64)
	// A threshold at which the system upscales pixel data.
	VideoZoomFactorUpscaleThreshold() float64
	SetVideoZoomFactorUpscaleThreshold(value float64)
	// A Boolean value that indicates whether the format supports zoom factors outside the range supported for depth delivery.
	ZoomFactorsOutsideOfVideoZoomRangesForDepthDeliverySupported() bool
	SetZoomFactorsOutsideOfVideoZoomRangesForDepthDeliverySupported(value bool)
}

// Init initializes the instance.
func (z AVZoomRange) Init() AVZoomRange {
	rv := objc.Send[AVZoomRange](z.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (z AVZoomRange) Autorelease() AVZoomRange {
	rv := objc.Send[AVZoomRange](z.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVZoomRange creates a new AVZoomRange instance.
func NewAVZoomRange() AVZoomRange {
	class := getAVZoomRangeClass()
	rv := objc.Send[AVZoomRange](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Returns a Boolean value that indicates whether the specified zoom factor
// exists in the range.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVZoomRange/containsZoomFactor:
func (z AVZoomRange) ContainsZoomFactor(zoomFactor float64) bool {
	rv := objc.Send[bool](z.ID, objc.Sel("containsZoomFactor:"), zoomFactor)
	return rv
}

// The range’s minimum zoom factor.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVZoomRange/minZoomFactor
func (z AVZoomRange) MinZoomFactor() float64 {
	rv := objc.Send[float64](z.ID, objc.Sel("minZoomFactor"))
	return rv
}

// The range’s maximum zoom factor.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVZoomRange/maxZoomFactor
func (z AVZoomRange) MaxZoomFactor() float64 {
	rv := objc.Send[float64](z.ID, objc.Sel("maxZoomFactor"))
	return rv
}

// A maximum zoom factor the format allows.
//
// See: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/videomaxzoomfactor
func (z AVZoomRange) VideoMaxZoomFactor() float64 {
	rv := objc.Send[float64](z.ID, objc.Sel("videoMaxZoomFactor"))
	return rv
}
func (z AVZoomRange) SetVideoMaxZoomFactor(value float64) {
	objc.Send[struct{}](z.ID, objc.Sel("setVideoMaxZoomFactor:"), value)
}

// A threshold at which the system upscales pixel data.
//
// See: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/videozoomfactorupscalethreshold
func (z AVZoomRange) VideoZoomFactorUpscaleThreshold() float64 {
	rv := objc.Send[float64](z.ID, objc.Sel("videoZoomFactorUpscaleThreshold"))
	return rv
}
func (z AVZoomRange) SetVideoZoomFactorUpscaleThreshold(value float64) {
	objc.Send[struct{}](z.ID, objc.Sel("setVideoZoomFactorUpscaleThreshold:"), value)
}

// A Boolean value that indicates whether the format supports zoom factors
// outside the range supported for depth delivery.
//
// See: https://developer.apple.com/documentation/avfoundation/avcapturedevice/format/zoomfactorsoutsideofvideozoomrangesfordepthdeliverysupported
func (z AVZoomRange) ZoomFactorsOutsideOfVideoZoomRangesForDepthDeliverySupported() bool {
	rv := objc.Send[bool](z.ID, objc.Sel("zoomFactorsOutsideOfVideoZoomRangesForDepthDeliverySupported"))
	return rv
}
func (z AVZoomRange) SetZoomFactorsOutsideOfVideoZoomRangesForDepthDeliverySupported(value bool) {
	objc.Send[struct{}](z.ID, objc.Sel("setZoomFactorsOutsideOfVideoZoomRangesForDepthDeliverySupported:"), value)
}
