// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"errors"
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLImageSizeConstraint] class.
var (
	_MLImageSizeConstraintClass     MLImageSizeConstraintClass
	_MLImageSizeConstraintClassOnce sync.Once
)

func getMLImageSizeConstraintClass() MLImageSizeConstraintClass {
	_MLImageSizeConstraintClassOnce.Do(func() {
		_MLImageSizeConstraintClass = MLImageSizeConstraintClass{class: objc.GetClass("MLImageSizeConstraint")}
	})
	return _MLImageSizeConstraintClass
}

// GetMLImageSizeConstraintClass returns the class object for MLImageSizeConstraint.
func GetMLImageSizeConstraintClass() MLImageSizeConstraintClass {
	return getMLImageSizeConstraintClass()
}

type MLImageSizeConstraintClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLImageSizeConstraintClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLImageSizeConstraintClass) Alloc() MLImageSizeConstraint {
	rv := objc.Send[MLImageSizeConstraint](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLImageSizeConstraint.AllowedImageSizeClosestToPixelsWidePixelsHighPreferDownScalingPreferInputAspectRatio]
//   - [MLImageSizeConstraint.ImageSizeSet]
//   - [MLImageSizeConstraint.IsAllowedImageSizeError]
//   - [MLImageSizeConstraint.InitUnspecified]
//   - [MLImageSizeConstraint.InitWithCoder]
//   - [MLImageSizeConstraint.InitWithEnumeratedImageSizes]
//
// See: https://developer.apple.com/documentation/CoreML/MLImageSizeConstraint
type MLImageSizeConstraint struct {
	objectivec.Object
}

// MLImageSizeConstraintFromID constructs a [MLImageSizeConstraint] from an objc.ID.
func MLImageSizeConstraintFromID(id objc.ID) MLImageSizeConstraint {
	return MLImageSizeConstraint{objectivec.Object{ID: id}}
}

// Ensure MLImageSizeConstraint implements IMLImageSizeConstraint.
var _ IMLImageSizeConstraint = MLImageSizeConstraint{}

// An interface definition for the [MLImageSizeConstraint] class.
//
// # Methods
//
//   - [IMLImageSizeConstraint.AllowedImageSizeClosestToPixelsWidePixelsHighPreferDownScalingPreferInputAspectRatio]
//   - [IMLImageSizeConstraint.ImageSizeSet]
//   - [IMLImageSizeConstraint.IsAllowedImageSizeError]
//   - [IMLImageSizeConstraint.InitUnspecified]
//   - [IMLImageSizeConstraint.InitWithCoder]
//   - [IMLImageSizeConstraint.InitWithEnumeratedImageSizes]
//
// See: https://developer.apple.com/documentation/CoreML/MLImageSizeConstraint
type IMLImageSizeConstraint interface {
	objectivec.IObject

	// Topic: Methods

	AllowedImageSizeClosestToPixelsWidePixelsHighPreferDownScalingPreferInputAspectRatio(wide int64, high int64, scaling bool, ratio bool) objectivec.IObject
	ImageSizeSet() foundation.INSOrderedSet
	IsAllowedImageSizeError(size objectivec.IObject) (bool, error)
	InitUnspecified() MLImageSizeConstraint
	InitWithCoder(coder foundation.INSCoder) MLImageSizeConstraint
	InitWithEnumeratedImageSizes(sizes objectivec.IObject) MLImageSizeConstraint
}

// Init initializes the instance.
func (i MLImageSizeConstraint) Init() MLImageSizeConstraint {
	rv := objc.Send[MLImageSizeConstraint](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i MLImageSizeConstraint) Autorelease() MLImageSizeConstraint {
	rv := objc.Send[MLImageSizeConstraint](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLImageSizeConstraint creates a new MLImageSizeConstraint instance.
func NewMLImageSizeConstraint() MLImageSizeConstraint {
	class := getMLImageSizeConstraintClass()
	rv := objc.Send[MLImageSizeConstraint](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLImageSizeConstraint/initUnspecified
func NewImageSizeConstraintUnspecified() MLImageSizeConstraint {
	instance := getMLImageSizeConstraintClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initUnspecified"))
	return MLImageSizeConstraintFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLImageSizeConstraint/initWithCoder:
func NewImageSizeConstraintWithCoder(coder objectivec.IObject) MLImageSizeConstraint {
	instance := getMLImageSizeConstraintClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return MLImageSizeConstraintFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLImageSizeConstraint/initWithEnumeratedImageSizes:
func NewImageSizeConstraintWithEnumeratedImageSizes(sizes objectivec.IObject) MLImageSizeConstraint {
	instance := getMLImageSizeConstraintClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithEnumeratedImageSizes:"), sizes)
	return MLImageSizeConstraintFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLImageSizeConstraint/allowedImageSizeClosestToPixelsWide:pixelsHigh:preferDownScaling:preferInputAspectRatio:
func (i MLImageSizeConstraint) AllowedImageSizeClosestToPixelsWidePixelsHighPreferDownScalingPreferInputAspectRatio(wide int64, high int64, scaling bool, ratio bool) objectivec.IObject {
	rv := objc.Send[objc.ID](i.ID, objc.Sel("allowedImageSizeClosestToPixelsWide:pixelsHigh:preferDownScaling:preferInputAspectRatio:"), wide, high, scaling, ratio)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/CoreML/MLImageSizeConstraint/isAllowedImageSize:error:
func (i MLImageSizeConstraint) IsAllowedImageSizeError(size objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](i.ID, objc.Sel("isAllowedImageSize:error:"), size, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("isAllowedImageSize:error: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/CoreML/MLImageSizeConstraint/initUnspecified
func (i MLImageSizeConstraint) InitUnspecified() MLImageSizeConstraint {
	rv := objc.Send[MLImageSizeConstraint](i.ID, objc.Sel("initUnspecified"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLImageSizeConstraint/initWithCoder:
func (i MLImageSizeConstraint) InitWithCoder(coder foundation.INSCoder) MLImageSizeConstraint {
	rv := objc.Send[MLImageSizeConstraint](i.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLImageSizeConstraint/initWithEnumeratedImageSizes:
func (i MLImageSizeConstraint) InitWithEnumeratedImageSizes(sizes objectivec.IObject) MLImageSizeConstraint {
	rv := objc.Send[MLImageSizeConstraint](i.ID, objc.Sel("initWithEnumeratedImageSizes:"), sizes)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLImageSizeConstraint/closestImageSizeInArray:toImageSize:preferDownScaling:
func (_MLImageSizeConstraintClass MLImageSizeConstraintClass) ClosestImageSizeInArrayToImageSizePreferDownScaling(array objectivec.IObject, size objectivec.IObject, scaling bool) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLImageSizeConstraintClass.class), objc.Sel("closestImageSizeInArray:toImageSize:preferDownScaling:"), array, size, scaling)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/CoreML/MLImageSizeConstraint/supportsSecureCoding
func (_MLImageSizeConstraintClass MLImageSizeConstraintClass) SupportsSecureCoding() bool {
	rv := objc.Send[bool](objc.ID(_MLImageSizeConstraintClass.class), objc.Sel("supportsSecureCoding"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLImageSizeConstraint/imageSizeSet
func (i MLImageSizeConstraint) ImageSizeSet() foundation.INSOrderedSet {
	rv := objc.Send[objc.ID](i.ID, objc.Sel("imageSizeSet"))
	return foundation.NSOrderedSetFromID(objc.ID(rv))
}
