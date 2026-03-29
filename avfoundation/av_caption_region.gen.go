// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVCaptionRegion] class.
var (
	_AVCaptionRegionClass     AVCaptionRegionClass
	_AVCaptionRegionClassOnce sync.Once
)

func getAVCaptionRegionClass() AVCaptionRegionClass {
	_AVCaptionRegionClassOnce.Do(func() {
		_AVCaptionRegionClass = AVCaptionRegionClass{class: objc.GetClass("AVCaptionRegion")}
	})
	return _AVCaptionRegionClass
}

// GetAVCaptionRegionClass returns the class object for AVCaptionRegion.
func GetAVCaptionRegionClass() AVCaptionRegionClass {
	return getAVCaptionRegionClass()
}

type AVCaptionRegionClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVCaptionRegionClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVCaptionRegionClass) Alloc() AVCaptionRegion {
	rv := objc.Send[AVCaptionRegion](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that represents the region in which the system presents a
// caption.
//
// # Overview
// 
// The framework defines four regions, and doesn’t support configuring
// region settings.
//
// # Identifying a region
//
//   - [AVCaptionRegion.Identifier]: A string that identifies the region.
//
// # Accessing the location
//
//   - [AVCaptionRegion.Origin]: The region’s top-left position.
//
// # Accessing the size
//
//   - [AVCaptionRegion.Size]: The height and width of the region.
//
// # Accessing the display alignment
//
//   - [AVCaptionRegion.DisplayAlignment]: The alignment of lines for the region.
//
// # Accessing the scroll mode
//
//   - [AVCaptionRegion.Scroll]: The scroll mode of the region.
//
// # Accessing the writing mode
//
//   - [AVCaptionRegion.WritingMode]: The block and inline progression direction of the region.
//
// # Processing regions
//
//   - [AVCaptionRegion.EncodeWithCoder]: Encodes the region using the specified encoder.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptionRegion
type AVCaptionRegion struct {
	objectivec.Object
}

// AVCaptionRegionFromID constructs a [AVCaptionRegion] from an objc.ID.
//
// An object that represents the region in which the system presents a
// caption.
func AVCaptionRegionFromID(id objc.ID) AVCaptionRegion {
	return AVCaptionRegion{objectivec.Object{ID: id}}
}
// NOTE: AVCaptionRegion adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVCaptionRegion] class.
//
// # Identifying a region
//
//   - [IAVCaptionRegion.Identifier]: A string that identifies the region.
//
// # Accessing the location
//
//   - [IAVCaptionRegion.Origin]: The region’s top-left position.
//
// # Accessing the size
//
//   - [IAVCaptionRegion.Size]: The height and width of the region.
//
// # Accessing the display alignment
//
//   - [IAVCaptionRegion.DisplayAlignment]: The alignment of lines for the region.
//
// # Accessing the scroll mode
//
//   - [IAVCaptionRegion.Scroll]: The scroll mode of the region.
//
// # Accessing the writing mode
//
//   - [IAVCaptionRegion.WritingMode]: The block and inline progression direction of the region.
//
// # Processing regions
//
//   - [IAVCaptionRegion.EncodeWithCoder]: Encodes the region using the specified encoder.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptionRegion
type IAVCaptionRegion interface {
	objectivec.IObject

	// Topic: Identifying a region

	// A string that identifies the region.
	Identifier() string

	// Topic: Accessing the location

	// The region’s top-left position.
	Origin() AVCaptionPoint

	// Topic: Accessing the size

	// The height and width of the region.
	Size() AVCaptionSize

	// Topic: Accessing the display alignment

	// The alignment of lines for the region.
	DisplayAlignment() AVCaptionRegionDisplayAlignment

	// Topic: Accessing the scroll mode

	// The scroll mode of the region.
	Scroll() AVCaptionRegionScroll

	// Topic: Accessing the writing mode

	// The block and inline progression direction of the region.
	WritingMode() AVCaptionRegionWritingMode

	// Topic: Processing regions

	// Encodes the region using the specified encoder.
	EncodeWithCoder(encoder foundation.INSCoder)
}

// Init initializes the instance.
func (c AVCaptionRegion) Init() AVCaptionRegion {
	rv := objc.Send[AVCaptionRegion](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c AVCaptionRegion) Autorelease() AVCaptionRegion {
	rv := objc.Send[AVCaptionRegion](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVCaptionRegion creates a new AVCaptionRegion instance.
func NewAVCaptionRegion() AVCaptionRegion {
	class := getAVCaptionRegionClass()
	rv := objc.Send[AVCaptionRegion](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Encodes the region using the specified encoder.
//
// encoder: An encoder instance to use.
//
// # Discussion
// 
// This method throws an exception if the caption region’s [Size] has
// different units for [width] and [height], or if the units are
// unrecognizeable.
//
// [height]: https://developer.apple.com/documentation/AVFoundation/AVCaptionSize/height
// [width]: https://developer.apple.com/documentation/AVFoundation/AVCaptionSize/width
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptionRegion/encode(with:)
func (c AVCaptionRegion) EncodeWithCoder(encoder foundation.INSCoder) {
	objc.Send[objc.ID](c.ID, objc.Sel("encodeWithCoder:"), encoder)
}

// A string that identifies the region.
//
// # Discussion
// 
// The system considers two regions the same if their region identifier is
// equal. Your app needs to ensure that equal caption regions have the same
// property values.
// 
// If this value is `nil`, the system instead treats two regions equal if
// their `position` and `endPosition` are the same. Captions referring to
// these regions belong to the same region when the system serializes them to
// a format like TTML.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptionRegion/identifier
func (c AVCaptionRegion) Identifier() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("identifier"))
	return foundation.NSStringFromID(rv).String()
}
// The region’s top-left position.
//
// # Discussion
// 
// The caption’s origin may not provide undefined [x] and [y] values, which
// indicates the region doesn’t have positioning information for that
// dimension.
//
// [x]: https://developer.apple.com/documentation/AVFoundation/AVCaptionPoint/x
// [y]: https://developer.apple.com/documentation/AVFoundation/AVCaptionPoint/y
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptionRegion/origin
func (c AVCaptionRegion) Origin() AVCaptionPoint {
	rv := objc.Send[AVCaptionPoint](c.ID, objc.Sel("origin"))
	return AVCaptionPoint(rv)
}
// The height and width of the region.
//
// # Discussion
// 
// CEA608 closed captions limit the [height] property’s value to 1 cell,
// except when the value of its [Scroll] property is
// [CaptionRegionScrollRollUp]. In this case, the [height] property’s value
// must be 2, 3 or 4 cells.
//
// [height]: https://developer.apple.com/documentation/AVFoundation/AVCaptionSize/height
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptionRegion/size
func (c AVCaptionRegion) Size() AVCaptionSize {
	rv := objc.Send[AVCaptionSize](c.ID, objc.Sel("size"))
	return AVCaptionSize(rv)
}
// The alignment of lines for the region.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptionRegion/displayAlignment-swift.property
func (c AVCaptionRegion) DisplayAlignment() AVCaptionRegionDisplayAlignment {
	rv := objc.Send[AVCaptionRegionDisplayAlignment](c.ID, objc.Sel("displayAlignment"))
	return AVCaptionRegionDisplayAlignment(rv)
}
// The scroll mode of the region.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptionRegion/scroll-swift.property
func (c AVCaptionRegion) Scroll() AVCaptionRegionScroll {
	rv := objc.Send[AVCaptionRegionScroll](c.ID, objc.Sel("scroll"))
	return AVCaptionRegionScroll(rv)
}
// The block and inline progression direction of the region.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptionRegion/writingMode-swift.property
func (c AVCaptionRegion) WritingMode() AVCaptionRegionWritingMode {
	rv := objc.Send[AVCaptionRegionWritingMode](c.ID, objc.Sel("writingMode"))
	return AVCaptionRegionWritingMode(rv)
}

// The top region for iTT format captions.
//
// # Discussion
// 
// This region is available when working with iTT captions. It occupies the
// top 15% of the display area, and it uses a LRTB layout where a line
// progresses from left to right and the block extends from top to bottom.
// Lines are top justified.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptionRegion/appleITTTop
func (_AVCaptionRegionClass AVCaptionRegionClass) AppleITTTopRegion() AVCaptionRegion {
	rv := objc.Send[objc.ID](objc.ID(_AVCaptionRegionClass.class), objc.Sel("appleITTTopRegion"))
	return AVCaptionRegionFromID(objc.ID(rv))
}
// The bottom region for iTT format captions.
//
// # Discussion
// 
// This region occupies the bottom 15% of the display area, and it uses a LRTB
// layout where a line progresses from left to right and the block extends
// from top to bottom. Lines are bottom justified.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptionRegion/appleITTBottom
func (_AVCaptionRegionClass AVCaptionRegionClass) AppleITTBottomRegion() AVCaptionRegion {
	rv := objc.Send[objc.ID](objc.ID(_AVCaptionRegionClass.class), objc.Sel("appleITTBottomRegion"))
	return AVCaptionRegionFromID(objc.ID(rv))
}
// The left region for iTT format captions.
//
// # Discussion
// 
// This region occupies the left 15% of the display area, and it uses a TBRL
// layout where a line progresses from top to bottom and the block extends
// from right to left. Lines are left justified.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptionRegion/appleITTLeft
func (_AVCaptionRegionClass AVCaptionRegionClass) AppleITTLeftRegion() AVCaptionRegion {
	rv := objc.Send[objc.ID](objc.ID(_AVCaptionRegionClass.class), objc.Sel("appleITTLeftRegion"))
	return AVCaptionRegionFromID(objc.ID(rv))
}
// The right region for iTT format captions.
//
// # Discussion
// 
// This region occupies the right 15% of the display area, and it uses a TBRL
// layout where a line progresses from top to bottom and the block extends
// from right to left. Lines are right justified.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptionRegion/appleITTRight
func (_AVCaptionRegionClass AVCaptionRegionClass) AppleITTRightRegion() AVCaptionRegion {
	rv := objc.Send[objc.ID](objc.ID(_AVCaptionRegionClass.class), objc.Sel("appleITTRightRegion"))
	return AVCaptionRegionFromID(objc.ID(rv))
}
// The bottom caption region for SubRip Text (SRT) format captions.
//
// # Discussion
// 
// This region is suitable for SRT format, and it occupies the entire video
// display area. The region uses a
// [CaptionRegionWritingModeLeftToRightAndTopToBottom] writing mode, where a
// line progresses left to right and the block extends from top to bottom. The
// system stacks each line of text with bottom justification.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptionRegion/subRipTextBottom
func (_AVCaptionRegionClass AVCaptionRegionClass) SubRipTextBottomRegion() AVCaptionRegion {
	rv := objc.Send[objc.ID](objc.ID(_AVCaptionRegionClass.class), objc.Sel("subRipTextBottomRegion"))
	return AVCaptionRegionFromID(objc.ID(rv))
}

