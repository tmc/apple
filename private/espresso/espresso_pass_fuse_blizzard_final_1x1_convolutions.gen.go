// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [EspressoPass_fuse_blizzard_final_1x1_convolutions] class.
var (
	_EspressoPass_fuse_blizzard_final_1x1_convolutionsClass     EspressoPass_fuse_blizzard_final_1x1_convolutionsClass
	_EspressoPass_fuse_blizzard_final_1x1_convolutionsClassOnce sync.Once
)

func getEspressoPass_fuse_blizzard_final_1x1_convolutionsClass() EspressoPass_fuse_blizzard_final_1x1_convolutionsClass {
	_EspressoPass_fuse_blizzard_final_1x1_convolutionsClassOnce.Do(func() {
		_EspressoPass_fuse_blizzard_final_1x1_convolutionsClass = EspressoPass_fuse_blizzard_final_1x1_convolutionsClass{class: objc.GetClass("EspressoPass_fuse_blizzard_final_1x1_convolutions")}
	})
	return _EspressoPass_fuse_blizzard_final_1x1_convolutionsClass
}

// GetEspressoPass_fuse_blizzard_final_1x1_convolutionsClass returns the class object for EspressoPass_fuse_blizzard_final_1x1_convolutions.
func GetEspressoPass_fuse_blizzard_final_1x1_convolutionsClass() EspressoPass_fuse_blizzard_final_1x1_convolutionsClass {
	return getEspressoPass_fuse_blizzard_final_1x1_convolutionsClass()
}

type EspressoPass_fuse_blizzard_final_1x1_convolutionsClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (ec EspressoPass_fuse_blizzard_final_1x1_convolutionsClass) Alloc() EspressoPass_fuse_blizzard_final_1x1_convolutions {
	rv := objc.Send[EspressoPass_fuse_blizzard_final_1x1_convolutions](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/Espresso/EspressoPass_fuse_blizzard_final_1x1_convolutions
type EspressoPass_fuse_blizzard_final_1x1_convolutions struct {
	EspressoCustomPass
}

// EspressoPass_fuse_blizzard_final_1x1_convolutionsFromID constructs a [EspressoPass_fuse_blizzard_final_1x1_convolutions] from an objc.ID.
func EspressoPass_fuse_blizzard_final_1x1_convolutionsFromID(id objc.ID) EspressoPass_fuse_blizzard_final_1x1_convolutions {
	return EspressoPass_fuse_blizzard_final_1x1_convolutions{EspressoCustomPass: EspressoCustomPassFromID(id)}
}
// Ensure EspressoPass_fuse_blizzard_final_1x1_convolutions implements IEspressoPass_fuse_blizzard_final_1x1_convolutions.
var _ IEspressoPass_fuse_blizzard_final_1x1_convolutions = EspressoPass_fuse_blizzard_final_1x1_convolutions{}

// An interface definition for the [EspressoPass_fuse_blizzard_final_1x1_convolutions] class.
//
// See: https://developer.apple.com/documentation/Espresso/EspressoPass_fuse_blizzard_final_1x1_convolutions
type IEspressoPass_fuse_blizzard_final_1x1_convolutions interface {
	IEspressoCustomPass
}

// Init initializes the instance.
func (e EspressoPass_fuse_blizzard_final_1x1_convolutions) Init() EspressoPass_fuse_blizzard_final_1x1_convolutions {
	rv := objc.Send[EspressoPass_fuse_blizzard_final_1x1_convolutions](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e EspressoPass_fuse_blizzard_final_1x1_convolutions) Autorelease() EspressoPass_fuse_blizzard_final_1x1_convolutions {
	rv := objc.Send[EspressoPass_fuse_blizzard_final_1x1_convolutions](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewEspressoPass_fuse_blizzard_final_1x1_convolutions creates a new EspressoPass_fuse_blizzard_final_1x1_convolutions instance.
func NewEspressoPass_fuse_blizzard_final_1x1_convolutions() EspressoPass_fuse_blizzard_final_1x1_convolutions {
	class := getEspressoPass_fuse_blizzard_final_1x1_convolutionsClass()
	rv := objc.Send[EspressoPass_fuse_blizzard_final_1x1_convolutions](objc.ID(class.class), objc.Sel("new"))
	return rv
}

