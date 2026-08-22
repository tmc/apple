// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [EspressoPassRemoveNhwcNchwTransposes] class.
var (
	_EspressoPassRemoveNhwcNchwTransposesClass     EspressoPassRemoveNhwcNchwTransposesClass
	_EspressoPassRemoveNhwcNchwTransposesClassOnce sync.Once
)

func getEspressoPassRemoveNhwcNchwTransposesClass() EspressoPassRemoveNhwcNchwTransposesClass {
	_EspressoPassRemoveNhwcNchwTransposesClassOnce.Do(func() {
		_EspressoPassRemoveNhwcNchwTransposesClass = EspressoPassRemoveNhwcNchwTransposesClass{class: objc.GetClass("EspressoPass_remove_nhwc_nchw_transposes")}
	})
	return _EspressoPassRemoveNhwcNchwTransposesClass
}

// GetEspressoPassRemoveNhwcNchwTransposesClass returns the class object for EspressoPass_remove_nhwc_nchw_transposes.
func GetEspressoPassRemoveNhwcNchwTransposesClass() EspressoPassRemoveNhwcNchwTransposesClass {
	return getEspressoPassRemoveNhwcNchwTransposesClass()
}

type EspressoPassRemoveNhwcNchwTransposesClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ec EspressoPassRemoveNhwcNchwTransposesClass) Class() objc.Class {
	return ec.class
}

// Alloc allocates memory for a new instance of the class.
func (ec EspressoPassRemoveNhwcNchwTransposesClass) Alloc() EspressoPassRemoveNhwcNchwTransposes {
	rv := objc.SendIfResponds[EspressoPassRemoveNhwcNchwTransposes](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

type EspressoPassRemoveNhwcNchwTransposes struct {
	EspressoCustomPass
}

// EspressoPassRemoveNhwcNchwTransposesFromID constructs a [EspressoPassRemoveNhwcNchwTransposes] from an objc.ID.
func EspressoPassRemoveNhwcNchwTransposesFromID(id objc.ID) EspressoPassRemoveNhwcNchwTransposes {
	return EspressoPassRemoveNhwcNchwTransposes{EspressoCustomPass: EspressoCustomPassFromID(id)}
}

// EspressoPass_remove_nhwc_nchw_transposesFromID is an alias for [EspressoPassRemoveNhwcNchwTransposesFromID] for cross-framework compatibility.
func EspressoPass_remove_nhwc_nchw_transposesFromID(id objc.ID) EspressoPassRemoveNhwcNchwTransposes {
	return EspressoPassRemoveNhwcNchwTransposesFromID(id)
}

// Ensure EspressoPassRemoveNhwcNchwTransposes implements IEspressoPassRemoveNhwcNchwTransposes.
var _ IEspressoPassRemoveNhwcNchwTransposes = EspressoPassRemoveNhwcNchwTransposes{}

// An interface definition for the [EspressoPassRemoveNhwcNchwTransposes] class.
type IEspressoPassRemoveNhwcNchwTransposes interface {
	IEspressoCustomPass
}

// Init initializes the instance.
func (e EspressoPassRemoveNhwcNchwTransposes) Init() EspressoPassRemoveNhwcNchwTransposes {
	rv := objc.SendIfResponds[EspressoPassRemoveNhwcNchwTransposes](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e EspressoPassRemoveNhwcNchwTransposes) Autorelease() EspressoPassRemoveNhwcNchwTransposes {
	rv := objc.SendIfResponds[EspressoPassRemoveNhwcNchwTransposes](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewEspressoPassRemoveNhwcNchwTransposes creates a new EspressoPassRemoveNhwcNchwTransposes instance.
func NewEspressoPassRemoveNhwcNchwTransposes() EspressoPassRemoveNhwcNchwTransposes {
	class := getEspressoPassRemoveNhwcNchwTransposesClass()
	rv := objc.SendIfResponds[EspressoPassRemoveNhwcNchwTransposes](objc.ID(class.class), objc.Sel("new"))
	return rv
}
