// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MaskPattern] class.
var (
	_MaskPatternClass     MaskPatternClass
	_MaskPatternClassOnce sync.Once
)

func getMaskPatternClass() MaskPatternClass {
	_MaskPatternClassOnce.Do(func() {
		_MaskPatternClass = MaskPatternClass{class: objc.GetClass("maskPattern")}
	})
	return _MaskPatternClass
}

// GetMaskPatternClass returns the class object for maskPattern.
func GetMaskPatternClass() MaskPatternClass {
	return getMaskPatternClass()
}

type MaskPatternClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MaskPatternClass) Alloc() MaskPattern {
	rv := objc.Send[MaskPattern](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreImage/CIQRCodeDescriptor/maskPattern-c.ivar
type MaskPattern struct {
	objectivec.Object
}

// MaskPatternFromID constructs a [MaskPattern] from an objc.ID.
func MaskPatternFromID(id objc.ID) MaskPattern {
	return MaskPattern{objectivec.Object{ID: id}}
}
// Ensure MaskPattern implements IMaskPattern.
var _ IMaskPattern = MaskPattern{}

// An interface definition for the [MaskPattern] class.
//
// See: https://developer.apple.com/documentation/CoreImage/CIQRCodeDescriptor/maskPattern-c.ivar
type IMaskPattern interface {
	objectivec.IObject
}

// Init initializes the instance.
func (m MaskPattern) Init() MaskPattern {
	rv := objc.Send[MaskPattern](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MaskPattern) Autorelease() MaskPattern {
	rv := objc.Send[MaskPattern](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMaskPattern creates a new MaskPattern instance.
func NewMaskPattern() MaskPattern {
	class := getMaskPatternClass()
	rv := objc.Send[MaskPattern](objc.ID(class.class), objc.Sel("new"))
	return rv
}

