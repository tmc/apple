// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NeuralNetworksPersistentGraphScope] class.
var (
	_NeuralNetworksPersistentGraphScopeClass     NeuralNetworksPersistentGraphScopeClass
	_NeuralNetworksPersistentGraphScopeClassOnce sync.Once
)

func getNeuralNetworksPersistentGraphScopeClass() NeuralNetworksPersistentGraphScopeClass {
	_NeuralNetworksPersistentGraphScopeClassOnce.Do(func() {
		_NeuralNetworksPersistentGraphScopeClass = NeuralNetworksPersistentGraphScopeClass{class: objc.GetClass("NeuralNetworks.PersistentGraphScope")}
	})
	return _NeuralNetworksPersistentGraphScopeClass
}

// GetNeuralNetworksPersistentGraphScopeClass returns the class object for NeuralNetworks.PersistentGraphScope.
func GetNeuralNetworksPersistentGraphScopeClass() NeuralNetworksPersistentGraphScopeClass {
	return getNeuralNetworksPersistentGraphScopeClass()
}

type NeuralNetworksPersistentGraphScopeClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NeuralNetworksPersistentGraphScopeClass) Alloc() NeuralNetworksPersistentGraphScope {
	rv := objc.Send[NeuralNetworksPersistentGraphScope](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.PersistentGraphScope
type NeuralNetworksPersistentGraphScope struct {
	objectivec.Object
}

// NeuralNetworksPersistentGraphScopeFromID constructs a [NeuralNetworksPersistentGraphScope] from an objc.ID.
func NeuralNetworksPersistentGraphScopeFromID(id objc.ID) NeuralNetworksPersistentGraphScope {
	return NeuralNetworksPersistentGraphScope{objectivec.Object{id}}
}
// Ensure NeuralNetworksPersistentGraphScope implements INeuralNetworksPersistentGraphScope.
var _ INeuralNetworksPersistentGraphScope = NeuralNetworksPersistentGraphScope{}





// An interface definition for the [NeuralNetworksPersistentGraphScope] class.
//
// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.PersistentGraphScope
type INeuralNetworksPersistentGraphScope interface {
	objectivec.IObject
}





// Init initializes the instance.
func (n NeuralNetworksPersistentGraphScope) Init() NeuralNetworksPersistentGraphScope {
	rv := objc.Send[NeuralNetworksPersistentGraphScope](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n NeuralNetworksPersistentGraphScope) Autorelease() NeuralNetworksPersistentGraphScope {
	rv := objc.Send[NeuralNetworksPersistentGraphScope](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewNeuralNetworksPersistentGraphScope creates a new NeuralNetworksPersistentGraphScope instance.
func NewNeuralNetworksPersistentGraphScope() NeuralNetworksPersistentGraphScope {
	class := getNeuralNetworksPersistentGraphScopeClass()
	rv := objc.Send[NeuralNetworksPersistentGraphScope](objc.ID(class.class), objc.Sel("new"))
	return rv
}



































