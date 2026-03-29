// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [AVMutableCaptionRegion] class.
var (
	_AVMutableCaptionRegionClass     AVMutableCaptionRegionClass
	_AVMutableCaptionRegionClassOnce sync.Once
)

func getAVMutableCaptionRegionClass() AVMutableCaptionRegionClass {
	_AVMutableCaptionRegionClassOnce.Do(func() {
		_AVMutableCaptionRegionClass = AVMutableCaptionRegionClass{class: objc.GetClass("AVMutableCaptionRegion")}
	})
	return _AVMutableCaptionRegionClass
}

// GetAVMutableCaptionRegionClass returns the class object for AVMutableCaptionRegion.
func GetAVMutableCaptionRegionClass() AVMutableCaptionRegionClass {
	return getAVMutableCaptionRegionClass()
}

type AVMutableCaptionRegionClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVMutableCaptionRegionClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVMutableCaptionRegionClass) Alloc() AVMutableCaptionRegion {
	rv := objc.Send[AVMutableCaptionRegion](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// A mutable caption region subclass that you use to create new caption
// regions.
//
// # Creating a caption region
//
//   - [AVMutableCaptionRegion.InitWithIdentifier]: Creates a caption region that has an identifier.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableCaptionRegion
type AVMutableCaptionRegion struct {
	AVCaptionRegion
}

// AVMutableCaptionRegionFromID constructs a [AVMutableCaptionRegion] from an objc.ID.
//
// A mutable caption region subclass that you use to create new caption
// regions.
func AVMutableCaptionRegionFromID(id objc.ID) AVMutableCaptionRegion {
	return AVMutableCaptionRegion{AVCaptionRegion: AVCaptionRegionFromID(id)}
}
// NOTE: AVMutableCaptionRegion adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVMutableCaptionRegion] class.
//
// # Creating a caption region
//
//   - [IAVMutableCaptionRegion.InitWithIdentifier]: Creates a caption region that has an identifier.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableCaptionRegion
type IAVMutableCaptionRegion interface {
	IAVCaptionRegion

	// Topic: Creating a caption region

	// Creates a caption region that has an identifier.
	InitWithIdentifier(identifier string) AVMutableCaptionRegion
}

// Init initializes the instance.
func (m AVMutableCaptionRegion) Init() AVMutableCaptionRegion {
	rv := objc.Send[AVMutableCaptionRegion](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m AVMutableCaptionRegion) Autorelease() AVMutableCaptionRegion {
	rv := objc.Send[AVMutableCaptionRegion](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVMutableCaptionRegion creates a new AVMutableCaptionRegion instance.
func NewAVMutableCaptionRegion() AVMutableCaptionRegion {
	class := getAVMutableCaptionRegionClass()
	rv := objc.Send[AVMutableCaptionRegion](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a caption region that has an identifier.
//
// identifier: The identifier of the caption region.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableCaptionRegion/init(identifier:)
func NewMutableCaptionRegionWithIdentifier(identifier string) AVMutableCaptionRegion {
	instance := getAVMutableCaptionRegionClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithIdentifier:"), objc.String(identifier))
	return AVMutableCaptionRegionFromID(rv)
}

// Creates a caption region that has an identifier.
//
// identifier: The identifier of the caption region.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableCaptionRegion/init(identifier:)
func (m AVMutableCaptionRegion) InitWithIdentifier(identifier string) AVMutableCaptionRegion {
	rv := objc.Send[AVMutableCaptionRegion](m.ID, objc.Sel("initWithIdentifier:"), objc.String(identifier))
	return rv
}

