// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [EspressoPassRemoveReshapeAroundRank2Softmax] class.
var (
	_EspressoPassRemoveReshapeAroundRank2SoftmaxClass     EspressoPassRemoveReshapeAroundRank2SoftmaxClass
	_EspressoPassRemoveReshapeAroundRank2SoftmaxClassOnce sync.Once
)

func getEspressoPassRemoveReshapeAroundRank2SoftmaxClass() EspressoPassRemoveReshapeAroundRank2SoftmaxClass {
	_EspressoPassRemoveReshapeAroundRank2SoftmaxClassOnce.Do(func() {
		_EspressoPassRemoveReshapeAroundRank2SoftmaxClass = EspressoPassRemoveReshapeAroundRank2SoftmaxClass{class: objc.GetClass("EspressoPass_remove_reshape_around_rank2_softmax")}
	})
	return _EspressoPassRemoveReshapeAroundRank2SoftmaxClass
}

// GetEspressoPassRemoveReshapeAroundRank2SoftmaxClass returns the class object for EspressoPass_remove_reshape_around_rank2_softmax.
func GetEspressoPassRemoveReshapeAroundRank2SoftmaxClass() EspressoPassRemoveReshapeAroundRank2SoftmaxClass {
	return getEspressoPassRemoveReshapeAroundRank2SoftmaxClass()
}

type EspressoPassRemoveReshapeAroundRank2SoftmaxClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ec EspressoPassRemoveReshapeAroundRank2SoftmaxClass) Class() objc.Class {
	return ec.class
}

// Alloc allocates memory for a new instance of the class.
func (ec EspressoPassRemoveReshapeAroundRank2SoftmaxClass) Alloc() EspressoPassRemoveReshapeAroundRank2Softmax {
	rv := objc.Send[EspressoPassRemoveReshapeAroundRank2Softmax](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/Espresso/EspressoPass_remove_reshape_around_rank2_softmax
type EspressoPassRemoveReshapeAroundRank2Softmax struct {
	EspressoCustomPass
}

// EspressoPassRemoveReshapeAroundRank2SoftmaxFromID constructs a [EspressoPassRemoveReshapeAroundRank2Softmax] from an objc.ID.
func EspressoPassRemoveReshapeAroundRank2SoftmaxFromID(id objc.ID) EspressoPassRemoveReshapeAroundRank2Softmax {
	return EspressoPassRemoveReshapeAroundRank2Softmax{EspressoCustomPass: EspressoCustomPassFromID(id)}
}

// EspressoPass_remove_reshape_around_rank2_softmaxFromID is an alias for [EspressoPassRemoveReshapeAroundRank2SoftmaxFromID] for cross-framework compatibility.
func EspressoPass_remove_reshape_around_rank2_softmaxFromID(id objc.ID) EspressoPassRemoveReshapeAroundRank2Softmax {
	return EspressoPassRemoveReshapeAroundRank2SoftmaxFromID(id)
}

// Ensure EspressoPassRemoveReshapeAroundRank2Softmax implements IEspressoPassRemoveReshapeAroundRank2Softmax.
var _ IEspressoPassRemoveReshapeAroundRank2Softmax = EspressoPassRemoveReshapeAroundRank2Softmax{}

// An interface definition for the [EspressoPassRemoveReshapeAroundRank2Softmax] class.
//
// See: https://developer.apple.com/documentation/Espresso/EspressoPass_remove_reshape_around_rank2_softmax
type IEspressoPassRemoveReshapeAroundRank2Softmax interface {
	IEspressoCustomPass
}

// Init initializes the instance.
func (e EspressoPassRemoveReshapeAroundRank2Softmax) Init() EspressoPassRemoveReshapeAroundRank2Softmax {
	rv := objc.Send[EspressoPassRemoveReshapeAroundRank2Softmax](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e EspressoPassRemoveReshapeAroundRank2Softmax) Autorelease() EspressoPassRemoveReshapeAroundRank2Softmax {
	rv := objc.Send[EspressoPassRemoveReshapeAroundRank2Softmax](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewEspressoPassRemoveReshapeAroundRank2Softmax creates a new EspressoPassRemoveReshapeAroundRank2Softmax instance.
func NewEspressoPassRemoveReshapeAroundRank2Softmax() EspressoPassRemoveReshapeAroundRank2Softmax {
	class := getEspressoPassRemoveReshapeAroundRank2SoftmaxClass()
	rv := objc.Send[EspressoPassRemoveReshapeAroundRank2Softmax](objc.ID(class.class), objc.Sel("new"))
	return rv
}
