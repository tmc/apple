// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [EspressoPassFuseBlizzardFinal1x1Convolutions] class.
var (
	_EspressoPassFuseBlizzardFinal1x1ConvolutionsClass     EspressoPassFuseBlizzardFinal1x1ConvolutionsClass
	_EspressoPassFuseBlizzardFinal1x1ConvolutionsClassOnce sync.Once
)

func getEspressoPassFuseBlizzardFinal1x1ConvolutionsClass() EspressoPassFuseBlizzardFinal1x1ConvolutionsClass {
	_EspressoPassFuseBlizzardFinal1x1ConvolutionsClassOnce.Do(func() {
		_EspressoPassFuseBlizzardFinal1x1ConvolutionsClass = EspressoPassFuseBlizzardFinal1x1ConvolutionsClass{class: objc.GetClass("EspressoPass_fuse_blizzard_final_1x1_convolutions")}
	})
	return _EspressoPassFuseBlizzardFinal1x1ConvolutionsClass
}

// GetEspressoPassFuseBlizzardFinal1x1ConvolutionsClass returns the class object for EspressoPass_fuse_blizzard_final_1x1_convolutions.
func GetEspressoPassFuseBlizzardFinal1x1ConvolutionsClass() EspressoPassFuseBlizzardFinal1x1ConvolutionsClass {
	return getEspressoPassFuseBlizzardFinal1x1ConvolutionsClass()
}

type EspressoPassFuseBlizzardFinal1x1ConvolutionsClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ec EspressoPassFuseBlizzardFinal1x1ConvolutionsClass) Class() objc.Class {
	return ec.class
}

// Alloc allocates memory for a new instance of the class.
func (ec EspressoPassFuseBlizzardFinal1x1ConvolutionsClass) Alloc() EspressoPassFuseBlizzardFinal1x1Convolutions {
	rv := objc.Send[EspressoPassFuseBlizzardFinal1x1Convolutions](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

type EspressoPassFuseBlizzardFinal1x1Convolutions struct {
	EspressoCustomPass
}

// EspressoPassFuseBlizzardFinal1x1ConvolutionsFromID constructs a [EspressoPassFuseBlizzardFinal1x1Convolutions] from an objc.ID.
func EspressoPassFuseBlizzardFinal1x1ConvolutionsFromID(id objc.ID) EspressoPassFuseBlizzardFinal1x1Convolutions {
	return EspressoPassFuseBlizzardFinal1x1Convolutions{EspressoCustomPass: EspressoCustomPassFromID(id)}
}

// EspressoPass_fuse_blizzard_final_1x1_convolutionsFromID is an alias for [EspressoPassFuseBlizzardFinal1x1ConvolutionsFromID] for cross-framework compatibility.
func EspressoPass_fuse_blizzard_final_1x1_convolutionsFromID(id objc.ID) EspressoPassFuseBlizzardFinal1x1Convolutions {
	return EspressoPassFuseBlizzardFinal1x1ConvolutionsFromID(id)
}

// Ensure EspressoPassFuseBlizzardFinal1x1Convolutions implements IEspressoPassFuseBlizzardFinal1x1Convolutions.
var _ IEspressoPassFuseBlizzardFinal1x1Convolutions = EspressoPassFuseBlizzardFinal1x1Convolutions{}

// An interface definition for the [EspressoPassFuseBlizzardFinal1x1Convolutions] class.
type IEspressoPassFuseBlizzardFinal1x1Convolutions interface {
	IEspressoCustomPass
}

// Init initializes the instance.
func (e EspressoPassFuseBlizzardFinal1x1Convolutions) Init() EspressoPassFuseBlizzardFinal1x1Convolutions {
	rv := objc.Send[EspressoPassFuseBlizzardFinal1x1Convolutions](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e EspressoPassFuseBlizzardFinal1x1Convolutions) Autorelease() EspressoPassFuseBlizzardFinal1x1Convolutions {
	rv := objc.Send[EspressoPassFuseBlizzardFinal1x1Convolutions](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewEspressoPassFuseBlizzardFinal1x1Convolutions creates a new EspressoPassFuseBlizzardFinal1x1Convolutions instance.
func NewEspressoPassFuseBlizzardFinal1x1Convolutions() EspressoPassFuseBlizzardFinal1x1Convolutions {
	class := getEspressoPassFuseBlizzardFinal1x1ConvolutionsClass()
	rv := objc.Send[EspressoPassFuseBlizzardFinal1x1Convolutions](objc.ID(class.class), objc.Sel("new"))
	return rv
}
