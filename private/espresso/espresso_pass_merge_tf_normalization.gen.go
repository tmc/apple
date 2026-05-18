// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [EspressoPassMergeTfNormalization] class.
var (
	_EspressoPassMergeTfNormalizationClass     EspressoPassMergeTfNormalizationClass
	_EspressoPassMergeTfNormalizationClassOnce sync.Once
)

func getEspressoPassMergeTfNormalizationClass() EspressoPassMergeTfNormalizationClass {
	_EspressoPassMergeTfNormalizationClassOnce.Do(func() {
		_EspressoPassMergeTfNormalizationClass = EspressoPassMergeTfNormalizationClass{class: objc.GetClass("EspressoPass_merge_tf_normalization")}
	})
	return _EspressoPassMergeTfNormalizationClass
}

// GetEspressoPassMergeTfNormalizationClass returns the class object for EspressoPass_merge_tf_normalization.
func GetEspressoPassMergeTfNormalizationClass() EspressoPassMergeTfNormalizationClass {
	return getEspressoPassMergeTfNormalizationClass()
}

type EspressoPassMergeTfNormalizationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ec EspressoPassMergeTfNormalizationClass) Class() objc.Class {
	return ec.class
}

// Alloc allocates memory for a new instance of the class.
func (ec EspressoPassMergeTfNormalizationClass) Alloc() EspressoPassMergeTfNormalization {
	rv := objc.Send[EspressoPassMergeTfNormalization](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/Espresso/EspressoPass_merge_tf_normalization
type EspressoPassMergeTfNormalization struct {
	EspressoCustomPass
}

// EspressoPassMergeTfNormalizationFromID constructs a [EspressoPassMergeTfNormalization] from an objc.ID.
func EspressoPassMergeTfNormalizationFromID(id objc.ID) EspressoPassMergeTfNormalization {
	return EspressoPassMergeTfNormalization{EspressoCustomPass: EspressoCustomPassFromID(id)}
}

// EspressoPass_merge_tf_normalizationFromID is an alias for [EspressoPassMergeTfNormalizationFromID] for cross-framework compatibility.
func EspressoPass_merge_tf_normalizationFromID(id objc.ID) EspressoPassMergeTfNormalization {
	return EspressoPassMergeTfNormalizationFromID(id)
}

// Ensure EspressoPassMergeTfNormalization implements IEspressoPassMergeTfNormalization.
var _ IEspressoPassMergeTfNormalization = EspressoPassMergeTfNormalization{}

// An interface definition for the [EspressoPassMergeTfNormalization] class.
//
// See: https://developer.apple.com/documentation/Espresso/EspressoPass_merge_tf_normalization
type IEspressoPassMergeTfNormalization interface {
	IEspressoCustomPass
}

// Init initializes the instance.
func (e EspressoPassMergeTfNormalization) Init() EspressoPassMergeTfNormalization {
	rv := objc.Send[EspressoPassMergeTfNormalization](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e EspressoPassMergeTfNormalization) Autorelease() EspressoPassMergeTfNormalization {
	rv := objc.Send[EspressoPassMergeTfNormalization](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewEspressoPassMergeTfNormalization creates a new EspressoPassMergeTfNormalization instance.
func NewEspressoPassMergeTfNormalization() EspressoPassMergeTfNormalization {
	class := getEspressoPassMergeTfNormalizationClass()
	rv := objc.Send[EspressoPassMergeTfNormalization](objc.ID(class.class), objc.Sel("new"))
	return rv
}
