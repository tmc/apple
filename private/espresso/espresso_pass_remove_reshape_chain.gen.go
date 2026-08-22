// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [EspressoPassRemoveReshapeChain] class.
var (
	_EspressoPassRemoveReshapeChainClass     EspressoPassRemoveReshapeChainClass
	_EspressoPassRemoveReshapeChainClassOnce sync.Once
)

func getEspressoPassRemoveReshapeChainClass() EspressoPassRemoveReshapeChainClass {
	_EspressoPassRemoveReshapeChainClassOnce.Do(func() {
		_EspressoPassRemoveReshapeChainClass = EspressoPassRemoveReshapeChainClass{class: objc.GetClass("EspressoPass_remove_reshape_chain")}
	})
	return _EspressoPassRemoveReshapeChainClass
}

// GetEspressoPassRemoveReshapeChainClass returns the class object for EspressoPass_remove_reshape_chain.
func GetEspressoPassRemoveReshapeChainClass() EspressoPassRemoveReshapeChainClass {
	return getEspressoPassRemoveReshapeChainClass()
}

type EspressoPassRemoveReshapeChainClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ec EspressoPassRemoveReshapeChainClass) Class() objc.Class {
	return ec.class
}

// Alloc allocates memory for a new instance of the class.
func (ec EspressoPassRemoveReshapeChainClass) Alloc() EspressoPassRemoveReshapeChain {
	rv := objc.SendIfResponds[EspressoPassRemoveReshapeChain](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

type EspressoPassRemoveReshapeChain struct {
	EspressoCustomPass
}

// EspressoPassRemoveReshapeChainFromID constructs a [EspressoPassRemoveReshapeChain] from an objc.ID.
func EspressoPassRemoveReshapeChainFromID(id objc.ID) EspressoPassRemoveReshapeChain {
	return EspressoPassRemoveReshapeChain{EspressoCustomPass: EspressoCustomPassFromID(id)}
}

// EspressoPass_remove_reshape_chainFromID is an alias for [EspressoPassRemoveReshapeChainFromID] for cross-framework compatibility.
func EspressoPass_remove_reshape_chainFromID(id objc.ID) EspressoPassRemoveReshapeChain {
	return EspressoPassRemoveReshapeChainFromID(id)
}

// Ensure EspressoPassRemoveReshapeChain implements IEspressoPassRemoveReshapeChain.
var _ IEspressoPassRemoveReshapeChain = EspressoPassRemoveReshapeChain{}

// An interface definition for the [EspressoPassRemoveReshapeChain] class.
type IEspressoPassRemoveReshapeChain interface {
	IEspressoCustomPass
}

// Init initializes the instance.
func (e EspressoPassRemoveReshapeChain) Init() EspressoPassRemoveReshapeChain {
	rv := objc.SendIfResponds[EspressoPassRemoveReshapeChain](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e EspressoPassRemoveReshapeChain) Autorelease() EspressoPassRemoveReshapeChain {
	rv := objc.SendIfResponds[EspressoPassRemoveReshapeChain](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewEspressoPassRemoveReshapeChain creates a new EspressoPassRemoveReshapeChain instance.
func NewEspressoPassRemoveReshapeChain() EspressoPassRemoveReshapeChain {
	class := getEspressoPassRemoveReshapeChainClass()
	rv := objc.SendIfResponds[EspressoPassRemoveReshapeChain](objc.ID(class.class), objc.Sel("new"))
	return rv
}
