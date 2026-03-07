// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NeuralNetworksUnfairLock] class.
var (
	_NeuralNetworksUnfairLockClass     NeuralNetworksUnfairLockClass
	_NeuralNetworksUnfairLockClassOnce sync.Once
)

func getNeuralNetworksUnfairLockClass() NeuralNetworksUnfairLockClass {
	_NeuralNetworksUnfairLockClassOnce.Do(func() {
		_NeuralNetworksUnfairLockClass = NeuralNetworksUnfairLockClass{class: objc.GetClass("NeuralNetworks.UnfairLock")}
	})
	return _NeuralNetworksUnfairLockClass
}

// GetNeuralNetworksUnfairLockClass returns the class object for NeuralNetworks.UnfairLock.
func GetNeuralNetworksUnfairLockClass() NeuralNetworksUnfairLockClass {
	return getNeuralNetworksUnfairLockClass()
}

type NeuralNetworksUnfairLockClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NeuralNetworksUnfairLockClass) Alloc() NeuralNetworksUnfairLock {
	rv := objc.Send[NeuralNetworksUnfairLock](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.UnfairLock
type NeuralNetworksUnfairLock struct {
	objectivec.Object
}

// NeuralNetworksUnfairLockFromID constructs a [NeuralNetworksUnfairLock] from an objc.ID.
func NeuralNetworksUnfairLockFromID(id objc.ID) NeuralNetworksUnfairLock {
	return NeuralNetworksUnfairLock{objectivec.Object{id}}
}
// Ensure NeuralNetworksUnfairLock implements INeuralNetworksUnfairLock.
var _ INeuralNetworksUnfairLock = NeuralNetworksUnfairLock{}





// An interface definition for the [NeuralNetworksUnfairLock] class.
//
// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.UnfairLock
type INeuralNetworksUnfairLock interface {
	objectivec.IObject
}





// Init initializes the instance.
func (n NeuralNetworksUnfairLock) Init() NeuralNetworksUnfairLock {
	rv := objc.Send[NeuralNetworksUnfairLock](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n NeuralNetworksUnfairLock) Autorelease() NeuralNetworksUnfairLock {
	rv := objc.Send[NeuralNetworksUnfairLock](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewNeuralNetworksUnfairLock creates a new NeuralNetworksUnfairLock instance.
func NewNeuralNetworksUnfairLock() NeuralNetworksUnfairLock {
	class := getNeuralNetworksUnfairLockClass()
	rv := objc.Send[NeuralNetworksUnfairLock](objc.ID(class.class), objc.Sel("new"))
	return rv
}



































