// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLImageSize] class.
var (
	_MLImageSizeClass     MLImageSizeClass
	_MLImageSizeClassOnce sync.Once
)

func getMLImageSizeClass() MLImageSizeClass {
	_MLImageSizeClassOnce.Do(func() {
		_MLImageSizeClass = MLImageSizeClass{class: objc.GetClass("MLImageSize")}
	})
	return _MLImageSizeClass
}

// GetMLImageSizeClass returns the class object for MLImageSize.
func GetMLImageSizeClass() MLImageSizeClass {
	return getMLImageSizeClass()
}

type MLImageSizeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLImageSizeClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLImageSizeClass) Alloc() MLImageSize {
	rv := objc.Send[MLImageSize](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLImageSize.IsEqualToImageSize]
//   - [MLImageSize.InitWithCoder]
//   - [MLImageSize.InitWithPixelsWidePixelsHigh]
//
// See: https://developer.apple.com/documentation/CoreML/MLImageSize
type MLImageSize struct {
	objectivec.Object
}

// MLImageSizeFromID constructs a [MLImageSize] from an objc.ID.
func MLImageSizeFromID(id objc.ID) MLImageSize {
	return MLImageSize{objectivec.Object{ID: id}}
}

// Ensure MLImageSize implements IMLImageSize.
var _ IMLImageSize = MLImageSize{}

// An interface definition for the [MLImageSize] class.
//
// # Methods
//
//   - [IMLImageSize.IsEqualToImageSize]
//   - [IMLImageSize.InitWithCoder]
//   - [IMLImageSize.InitWithPixelsWidePixelsHigh]
//
// See: https://developer.apple.com/documentation/CoreML/MLImageSize
type IMLImageSize interface {
	objectivec.IObject

	// Topic: Methods

	IsEqualToImageSize(size objectivec.IObject) bool
	InitWithCoder(coder foundation.INSCoder) MLImageSize
	InitWithPixelsWidePixelsHigh(wide int64, high int64) MLImageSize
}

// Init initializes the instance.
func (i MLImageSize) Init() MLImageSize {
	rv := objc.Send[MLImageSize](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i MLImageSize) Autorelease() MLImageSize {
	rv := objc.Send[MLImageSize](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLImageSize creates a new MLImageSize instance.
func NewMLImageSize() MLImageSize {
	class := getMLImageSizeClass()
	rv := objc.Send[MLImageSize](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLImageSize/initWithCoder:
func NewImageSizeWithCoder(coder objectivec.IObject) MLImageSize {
	instance := getMLImageSizeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return MLImageSizeFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLImageSize/initWithPixelsWide:pixelsHigh:
func NewImageSizeWithPixelsWidePixelsHigh(wide int64, high int64) MLImageSize {
	instance := getMLImageSizeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithPixelsWide:pixelsHigh:"), wide, high)
	return MLImageSizeFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLImageSize/isEqualToImageSize:
func (i MLImageSize) IsEqualToImageSize(size objectivec.IObject) bool {
	rv := objc.Send[bool](i.ID, objc.Sel("isEqualToImageSize:"), size)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLImageSize/initWithCoder:
func (i MLImageSize) InitWithCoder(coder foundation.INSCoder) MLImageSize {
	rv := objc.Send[MLImageSize](i.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLImageSize/initWithPixelsWide:pixelsHigh:
func (i MLImageSize) InitWithPixelsWidePixelsHigh(wide int64, high int64) MLImageSize {
	rv := objc.Send[MLImageSize](i.ID, objc.Sel("initWithPixelsWide:pixelsHigh:"), wide, high)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLImageSize/supportsSecureCoding
func (_MLImageSizeClass MLImageSizeClass) SupportsSecureCoding() bool {
	rv := objc.Send[bool](objc.ID(_MLImageSizeClass.class), objc.Sel("supportsSecureCoding"))
	return rv
}
