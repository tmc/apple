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

// The class instance for the [MLImageConstraint] class.
var (
	_MLImageConstraintClass     MLImageConstraintClass
	_MLImageConstraintClassOnce sync.Once
)

func getMLImageConstraintClass() MLImageConstraintClass {
	_MLImageConstraintClassOnce.Do(func() {
		_MLImageConstraintClass = MLImageConstraintClass{class: objc.GetClass("MLImageConstraint")}
	})
	return _MLImageConstraintClass
}

// GetMLImageConstraintClass returns the class object for MLImageConstraint.
func GetMLImageConstraintClass() MLImageConstraintClass {
	return getMLImageConstraintClass()
}

type MLImageConstraintClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLImageConstraintClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLImageConstraintClass) Alloc() MLImageConstraint {
	rv := objc.Send[MLImageConstraint](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLImageConstraint._stringForAllowedOSTypes]
//   - [MLImageConstraint.ImageHeight]
//   - [MLImageConstraint.ImageWidth]
//   - [MLImageConstraint.IsAllowedValueError]
//   - [MLImageConstraint.OsType]
//   - [MLImageConstraint.PixelType]
//   - [MLImageConstraint.InitWithCoder]
//   - [MLImageConstraint.InitWithPixelsWidePixelsHighPixelTypeSizeConstraint]
//
// See: https://developer.apple.com/documentation/CoreML/MLImageConstraint
type MLImageConstraint struct {
	objectivec.Object
}

// MLImageConstraintFromID constructs a [MLImageConstraint] from an objc.ID.
func MLImageConstraintFromID(id objc.ID) MLImageConstraint {
	return MLImageConstraint{objectivec.Object{ID: id}}
}

// Ensure MLImageConstraint implements IMLImageConstraint.
var _ IMLImageConstraint = MLImageConstraint{}

// An interface definition for the [MLImageConstraint] class.
//
// # Methods
//
//   - [IMLImageConstraint._stringForAllowedOSTypes]
//   - [IMLImageConstraint.ImageHeight]
//   - [IMLImageConstraint.ImageWidth]
//   - [IMLImageConstraint.IsAllowedValueError]
//   - [IMLImageConstraint.OsType]
//   - [IMLImageConstraint.PixelType]
//   - [IMLImageConstraint.InitWithCoder]
//   - [IMLImageConstraint.InitWithPixelsWidePixelsHighPixelTypeSizeConstraint]
//
// See: https://developer.apple.com/documentation/CoreML/MLImageConstraint
type IMLImageConstraint interface {
	objectivec.IObject

	// Topic: Methods

	_stringForAllowedOSTypes() objectivec.IObject
	ImageHeight() uint64
	ImageWidth() uint64
	IsAllowedValueError(value objectivec.IObject) (bool, error)
	OsType() uint32
	PixelType() uint64
	InitWithCoder(coder foundation.INSCoder) MLImageConstraint
	InitWithPixelsWidePixelsHighPixelTypeSizeConstraint(wide int64, high int64, type_ uint64, constraint objectivec.IObject) MLImageConstraint
}

// Init initializes the instance.
func (i MLImageConstraint) Init() MLImageConstraint {
	rv := objc.Send[MLImageConstraint](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i MLImageConstraint) Autorelease() MLImageConstraint {
	rv := objc.Send[MLImageConstraint](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLImageConstraint creates a new MLImageConstraint instance.
func NewMLImageConstraint() MLImageConstraint {
	class := getMLImageConstraintClass()
	rv := objc.Send[MLImageConstraint](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLImageConstraint/initWithCoder:
func NewImageConstraintWithCoder(coder objectivec.IObject) MLImageConstraint {
	instance := getMLImageConstraintClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return MLImageConstraintFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLImageConstraint/initWithPixelsWide:pixelsHigh:pixelType:sizeConstraint:
func NewImageConstraintWithPixelsWidePixelsHighPixelTypeSizeConstraint(wide int64, high int64, type_ uint64, constraint objectivec.IObject) MLImageConstraint {
	instance := getMLImageConstraintClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithPixelsWide:pixelsHigh:pixelType:sizeConstraint:"), wide, high, type_, constraint)
	return MLImageConstraintFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLImageConstraint/_stringForAllowedOSTypes
func (i MLImageConstraint) _stringForAllowedOSTypes() objectivec.IObject {
	rv := objc.Send[objc.ID](i.ID, objc.Sel("_stringForAllowedOSTypes"))
	return objectivec.Object{ID: rv}
}

// StringForAllowedOSTypes is an exported wrapper for the private method _stringForAllowedOSTypes.
func (i MLImageConstraint) StringForAllowedOSTypes() objectivec.IObject {
	return i._stringForAllowedOSTypes()
}

// See: https://developer.apple.com/documentation/CoreML/MLImageConstraint/isAllowedValue:error:
func (i MLImageConstraint) IsAllowedValueError(value objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](i.ID, objc.Sel("isAllowedValue:error:"), value, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("isAllowedValue:error: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/CoreML/MLImageConstraint/initWithCoder:
func (i MLImageConstraint) InitWithCoder(coder foundation.INSCoder) MLImageConstraint {
	rv := objc.Send[MLImageConstraint](i.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLImageConstraint/initWithPixelsWide:pixelsHigh:pixelType:sizeConstraint:
func (i MLImageConstraint) InitWithPixelsWidePixelsHighPixelTypeSizeConstraint(wide int64, high int64, type_ uint64, constraint objectivec.IObject) MLImageConstraint {
	rv := objc.Send[MLImageConstraint](i.ID, objc.Sel("initWithPixelsWide:pixelsHigh:pixelType:sizeConstraint:"), wide, high, type_, constraint)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLImageConstraint/constraintWithPixelsWide:pixelsHigh:pixelType:
func (_MLImageConstraintClass MLImageConstraintClass) ConstraintWithPixelsWidePixelsHighPixelType(wide int64, high int64, type_ uint64) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLImageConstraintClass.class), objc.Sel("constraintWithPixelsWide:pixelsHigh:pixelType:"), wide, high, type_)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/CoreML/MLImageConstraint/constraintWithPixelsWide:pixelsHigh:pixelType:sizeConstraint:
func (_MLImageConstraintClass MLImageConstraintClass) ConstraintWithPixelsWidePixelsHighPixelTypeSizeConstraint(wide int64, high int64, type_ uint64, constraint objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLImageConstraintClass.class), objc.Sel("constraintWithPixelsWide:pixelsHigh:pixelType:sizeConstraint:"), wide, high, type_, constraint)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/CoreML/MLImageConstraint/imagePixelTypeFromOSType:
func (_MLImageConstraintClass MLImageConstraintClass) ImagePixelTypeFromOSType(oSType uint32) uint64 {
	rv := objc.Send[uint64](objc.ID(_MLImageConstraintClass.class), objc.Sel("imagePixelTypeFromOSType:"), oSType)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLImageConstraint/stringForImagePixelType:
func (_MLImageConstraintClass MLImageConstraintClass) StringForImagePixelType(type_ uint64) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLImageConstraintClass.class), objc.Sel("stringForImagePixelType:"), type_)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/CoreML/MLImageConstraint/stringForPixelFormatType:
func (_MLImageConstraintClass MLImageConstraintClass) StringForPixelFormatType(type_ uint32) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLImageConstraintClass.class), objc.Sel("stringForPixelFormatType:"), type_)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/CoreML/MLImageConstraint/supportsSecureCoding
func (_MLImageConstraintClass MLImageConstraintClass) SupportsSecureCoding() bool {
	rv := objc.Send[bool](objc.ID(_MLImageConstraintClass.class), objc.Sel("supportsSecureCoding"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLImageConstraint/imageHeight
func (i MLImageConstraint) ImageHeight() uint64 {
	rv := objc.Send[uint64](i.ID, objc.Sel("imageHeight"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLImageConstraint/imageWidth
func (i MLImageConstraint) ImageWidth() uint64 {
	rv := objc.Send[uint64](i.ID, objc.Sel("imageWidth"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLImageConstraint/osType
func (i MLImageConstraint) OsType() uint32 {
	rv := objc.Send[uint32](i.ID, objc.Sel("osType"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLImageConstraint/pixelType
func (i MLImageConstraint) PixelType() uint64 {
	rv := objc.Send[uint64](i.ID, objc.Sel("pixelType"))
	return rv
}
