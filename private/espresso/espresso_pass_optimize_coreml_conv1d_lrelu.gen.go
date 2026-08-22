// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [EspressoPassOptimizeCoremlConv1dLrelu] class.
var (
	_EspressoPassOptimizeCoremlConv1dLreluClass     EspressoPassOptimizeCoremlConv1dLreluClass
	_EspressoPassOptimizeCoremlConv1dLreluClassOnce sync.Once
)

func getEspressoPassOptimizeCoremlConv1dLreluClass() EspressoPassOptimizeCoremlConv1dLreluClass {
	_EspressoPassOptimizeCoremlConv1dLreluClassOnce.Do(func() {
		_EspressoPassOptimizeCoremlConv1dLreluClass = EspressoPassOptimizeCoremlConv1dLreluClass{class: objc.GetClass("EspressoPass_optimize_coreml_conv1d_lrelu")}
	})
	return _EspressoPassOptimizeCoremlConv1dLreluClass
}

// GetEspressoPassOptimizeCoremlConv1dLreluClass returns the class object for EspressoPass_optimize_coreml_conv1d_lrelu.
func GetEspressoPassOptimizeCoremlConv1dLreluClass() EspressoPassOptimizeCoremlConv1dLreluClass {
	return getEspressoPassOptimizeCoremlConv1dLreluClass()
}

type EspressoPassOptimizeCoremlConv1dLreluClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ec EspressoPassOptimizeCoremlConv1dLreluClass) Class() objc.Class {
	return ec.class
}

// Alloc allocates memory for a new instance of the class.
func (ec EspressoPassOptimizeCoremlConv1dLreluClass) Alloc() EspressoPassOptimizeCoremlConv1dLrelu {
	rv := objc.SendIfResponds[EspressoPassOptimizeCoremlConv1dLrelu](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

type EspressoPassOptimizeCoremlConv1dLrelu struct {
	EspressoCustomPass
}

// EspressoPassOptimizeCoremlConv1dLreluFromID constructs a [EspressoPassOptimizeCoremlConv1dLrelu] from an objc.ID.
func EspressoPassOptimizeCoremlConv1dLreluFromID(id objc.ID) EspressoPassOptimizeCoremlConv1dLrelu {
	return EspressoPassOptimizeCoremlConv1dLrelu{EspressoCustomPass: EspressoCustomPassFromID(id)}
}

// EspressoPass_optimize_coreml_conv1d_lreluFromID is an alias for [EspressoPassOptimizeCoremlConv1dLreluFromID] for cross-framework compatibility.
func EspressoPass_optimize_coreml_conv1d_lreluFromID(id objc.ID) EspressoPassOptimizeCoremlConv1dLrelu {
	return EspressoPassOptimizeCoremlConv1dLreluFromID(id)
}

// Ensure EspressoPassOptimizeCoremlConv1dLrelu implements IEspressoPassOptimizeCoremlConv1dLrelu.
var _ IEspressoPassOptimizeCoremlConv1dLrelu = EspressoPassOptimizeCoremlConv1dLrelu{}

// An interface definition for the [EspressoPassOptimizeCoremlConv1dLrelu] class.
type IEspressoPassOptimizeCoremlConv1dLrelu interface {
	IEspressoCustomPass
}

// Init initializes the instance.
func (e EspressoPassOptimizeCoremlConv1dLrelu) Init() EspressoPassOptimizeCoremlConv1dLrelu {
	rv := objc.SendIfResponds[EspressoPassOptimizeCoremlConv1dLrelu](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e EspressoPassOptimizeCoremlConv1dLrelu) Autorelease() EspressoPassOptimizeCoremlConv1dLrelu {
	rv := objc.SendIfResponds[EspressoPassOptimizeCoremlConv1dLrelu](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewEspressoPassOptimizeCoremlConv1dLrelu creates a new EspressoPassOptimizeCoremlConv1dLrelu instance.
func NewEspressoPassOptimizeCoremlConv1dLrelu() EspressoPassOptimizeCoremlConv1dLrelu {
	class := getEspressoPassOptimizeCoremlConv1dLreluClass()
	rv := objc.SendIfResponds[EspressoPassOptimizeCoremlConv1dLrelu](objc.ID(class.class), objc.Sel("new"))
	return rv
}
