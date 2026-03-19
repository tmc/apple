// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
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

// Alloc allocates memory for a new instance of the class.
func (mc MLImageSizeClass) Alloc() MLImageSize {
	rv := objc.Send[MLImageSize](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// The width and height of an image feature size.
//
// # Accessing an image size
//
//   - [MLImageSize.PixelsHigh]: The height of an image feature in pixels.
//   - [MLImageSize.PixelsWide]: The width of an image feature in pixels.
//
// See: https://developer.apple.com/documentation/CoreML/MLImageSize
type MLImageSize struct {
	objectivec.Object
}

// MLImageSizeFromID constructs a [MLImageSize] from an objc.ID.
//
// The width and height of an image feature size.
func MLImageSizeFromID(id objc.ID) MLImageSize {
	return MLImageSize{objectivec.Object{ID: id}}
}
// NOTE: MLImageSize adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MLImageSize] class.
//
// # Accessing an image size
//
//   - [IMLImageSize.PixelsHigh]: The height of an image feature in pixels.
//   - [IMLImageSize.PixelsWide]: The width of an image feature in pixels.
//
// See: https://developer.apple.com/documentation/CoreML/MLImageSize
type IMLImageSize interface {
	objectivec.IObject

	// Topic: Accessing an image size

	// The height of an image feature in pixels.
	PixelsHigh() int
	// The width of an image feature in pixels.
	PixelsWide() int

	// An array of image sizes a model’s image feature accepts as input or produces as output.
	EnumeratedImageSizes() IMLImageSize
	SetEnumeratedImageSizes(value IMLImageSize)
	EncodeWithCoder(coder foundation.INSCoder)
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

func (i MLImageSize) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](i.ID, objc.Sel("encodeWithCoder:"), coder)
}

// The height of an image feature in pixels.
//
// See: https://developer.apple.com/documentation/CoreML/MLImageSize/pixelsHigh
func (i MLImageSize) PixelsHigh() int {
	rv := objc.Send[int](i.ID, objc.Sel("pixelsHigh"))
	return rv
}
// The width of an image feature in pixels.
//
// See: https://developer.apple.com/documentation/CoreML/MLImageSize/pixelsWide
func (i MLImageSize) PixelsWide() int {
	rv := objc.Send[int](i.ID, objc.Sel("pixelsWide"))
	return rv
}
// An array of image sizes a model’s image feature accepts as input or
// produces as output.
//
// See: https://developer.apple.com/documentation/coreml/mlimagesizeconstraint/enumeratedimagesizes
func (i MLImageSize) EnumeratedImageSizes() IMLImageSize {
	rv := objc.Send[objc.ID](i.ID, objc.Sel("enumeratedImageSizes"))
	return MLImageSizeFromID(objc.ID(rv))
}
func (i MLImageSize) SetEnumeratedImageSizes(value IMLImageSize) {
	objc.Send[struct{}](i.ID, objc.Sel("setEnumeratedImageSizes:"), value)
}

