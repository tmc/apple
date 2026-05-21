// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [EspressoPassFuseFastGelu1] class.
var (
	_EspressoPassFuseFastGelu1Class     EspressoPassFuseFastGelu1Class
	_EspressoPassFuseFastGelu1ClassOnce sync.Once
)

func getEspressoPassFuseFastGelu1Class() EspressoPassFuseFastGelu1Class {
	_EspressoPassFuseFastGelu1ClassOnce.Do(func() {
		_EspressoPassFuseFastGelu1Class = EspressoPassFuseFastGelu1Class{class: objc.GetClass("EspressoPass_fuse_fast_gelu_1")}
	})
	return _EspressoPassFuseFastGelu1Class
}

// GetEspressoPassFuseFastGelu1Class returns the class object for EspressoPass_fuse_fast_gelu_1.
func GetEspressoPassFuseFastGelu1Class() EspressoPassFuseFastGelu1Class {
	return getEspressoPassFuseFastGelu1Class()
}

type EspressoPassFuseFastGelu1Class struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ec EspressoPassFuseFastGelu1Class) Class() objc.Class {
	return ec.class
}

// Alloc allocates memory for a new instance of the class.
func (ec EspressoPassFuseFastGelu1Class) Alloc() EspressoPassFuseFastGelu1 {
	rv := objc.Send[EspressoPassFuseFastGelu1](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

type EspressoPassFuseFastGelu1 struct {
	EspressoCustomPass
}

// EspressoPassFuseFastGelu1FromID constructs a [EspressoPassFuseFastGelu1] from an objc.ID.
func EspressoPassFuseFastGelu1FromID(id objc.ID) EspressoPassFuseFastGelu1 {
	return EspressoPassFuseFastGelu1{EspressoCustomPass: EspressoCustomPassFromID(id)}
}

// EspressoPass_fuse_fast_gelu_1FromID is an alias for [EspressoPassFuseFastGelu1FromID] for cross-framework compatibility.
func EspressoPass_fuse_fast_gelu_1FromID(id objc.ID) EspressoPassFuseFastGelu1 {
	return EspressoPassFuseFastGelu1FromID(id)
}

// Ensure EspressoPassFuseFastGelu1 implements IEspressoPassFuseFastGelu1.
var _ IEspressoPassFuseFastGelu1 = EspressoPassFuseFastGelu1{}

// An interface definition for the [EspressoPassFuseFastGelu1] class.
type IEspressoPassFuseFastGelu1 interface {
	IEspressoCustomPass
}

// Init initializes the instance.
func (e EspressoPassFuseFastGelu1) Init() EspressoPassFuseFastGelu1 {
	rv := objc.Send[EspressoPassFuseFastGelu1](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e EspressoPassFuseFastGelu1) Autorelease() EspressoPassFuseFastGelu1 {
	rv := objc.Send[EspressoPassFuseFastGelu1](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewEspressoPassFuseFastGelu1 creates a new EspressoPassFuseFastGelu1 instance.
func NewEspressoPassFuseFastGelu1() EspressoPassFuseFastGelu1 {
	class := getEspressoPassFuseFastGelu1Class()
	rv := objc.Send[EspressoPassFuseFastGelu1](objc.ID(class.class), objc.Sel("new"))
	return rv
}
