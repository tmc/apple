// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [EspressoPass_merge_tf_normalization] class.
var (
	_EspressoPass_merge_tf_normalizationClass     EspressoPass_merge_tf_normalizationClass
	_EspressoPass_merge_tf_normalizationClassOnce sync.Once
)

func getEspressoPass_merge_tf_normalizationClass() EspressoPass_merge_tf_normalizationClass {
	_EspressoPass_merge_tf_normalizationClassOnce.Do(func() {
		_EspressoPass_merge_tf_normalizationClass = EspressoPass_merge_tf_normalizationClass{class: objc.GetClass("EspressoPass_merge_tf_normalization")}
	})
	return _EspressoPass_merge_tf_normalizationClass
}

// GetEspressoPass_merge_tf_normalizationClass returns the class object for EspressoPass_merge_tf_normalization.
func GetEspressoPass_merge_tf_normalizationClass() EspressoPass_merge_tf_normalizationClass {
	return getEspressoPass_merge_tf_normalizationClass()
}

type EspressoPass_merge_tf_normalizationClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (ec EspressoPass_merge_tf_normalizationClass) Alloc() EspressoPass_merge_tf_normalization {
	rv := objc.Send[EspressoPass_merge_tf_normalization](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/Espresso/EspressoPass_merge_tf_normalization
type EspressoPass_merge_tf_normalization struct {
	EspressoCustomPass
}

// EspressoPass_merge_tf_normalizationFromID constructs a [EspressoPass_merge_tf_normalization] from an objc.ID.
func EspressoPass_merge_tf_normalizationFromID(id objc.ID) EspressoPass_merge_tf_normalization {
	return EspressoPass_merge_tf_normalization{EspressoCustomPass: EspressoCustomPassFromID(id)}
}
// Ensure EspressoPass_merge_tf_normalization implements IEspressoPass_merge_tf_normalization.
var _ IEspressoPass_merge_tf_normalization = EspressoPass_merge_tf_normalization{}





// An interface definition for the [EspressoPass_merge_tf_normalization] class.
//
// See: https://developer.apple.com/documentation/Espresso/EspressoPass_merge_tf_normalization
type IEspressoPass_merge_tf_normalization interface {
	IEspressoCustomPass
}





// Init initializes the instance.
func (e EspressoPass_merge_tf_normalization) Init() EspressoPass_merge_tf_normalization {
	rv := objc.Send[EspressoPass_merge_tf_normalization](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e EspressoPass_merge_tf_normalization) Autorelease() EspressoPass_merge_tf_normalization {
	rv := objc.Send[EspressoPass_merge_tf_normalization](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewEspressoPass_merge_tf_normalization creates a new EspressoPass_merge_tf_normalization instance.
func NewEspressoPass_merge_tf_normalization() EspressoPass_merge_tf_normalization {
	class := getEspressoPass_merge_tf_normalizationClass()
	rv := objc.Send[EspressoPass_merge_tf_normalization](objc.ID(class.class), objc.Sel("new"))
	return rv
}



































