// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NeuralNetworksResolutionDependencies] class.
var (
	_NeuralNetworksResolutionDependenciesClass     NeuralNetworksResolutionDependenciesClass
	_NeuralNetworksResolutionDependenciesClassOnce sync.Once
)

func getNeuralNetworksResolutionDependenciesClass() NeuralNetworksResolutionDependenciesClass {
	_NeuralNetworksResolutionDependenciesClassOnce.Do(func() {
		_NeuralNetworksResolutionDependenciesClass = NeuralNetworksResolutionDependenciesClass{class: objc.GetClass("NeuralNetworks.ResolutionDependencies")}
	})
	return _NeuralNetworksResolutionDependenciesClass
}

// GetNeuralNetworksResolutionDependenciesClass returns the class object for NeuralNetworks.ResolutionDependencies.
func GetNeuralNetworksResolutionDependenciesClass() NeuralNetworksResolutionDependenciesClass {
	return getNeuralNetworksResolutionDependenciesClass()
}

type NeuralNetworksResolutionDependenciesClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NeuralNetworksResolutionDependenciesClass) Alloc() NeuralNetworksResolutionDependencies {
	rv := objc.Send[NeuralNetworksResolutionDependencies](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.ResolutionDependencies
type NeuralNetworksResolutionDependencies struct {
	objectivec.Object
}

// NeuralNetworksResolutionDependenciesFromID constructs a [NeuralNetworksResolutionDependencies] from an objc.ID.
func NeuralNetworksResolutionDependenciesFromID(id objc.ID) NeuralNetworksResolutionDependencies {
	return NeuralNetworksResolutionDependencies{objectivec.Object{id}}
}
// Ensure NeuralNetworksResolutionDependencies implements INeuralNetworksResolutionDependencies.
var _ INeuralNetworksResolutionDependencies = NeuralNetworksResolutionDependencies{}





// An interface definition for the [NeuralNetworksResolutionDependencies] class.
//
// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.ResolutionDependencies
type INeuralNetworksResolutionDependencies interface {
	objectivec.IObject
}





// Init initializes the instance.
func (n NeuralNetworksResolutionDependencies) Init() NeuralNetworksResolutionDependencies {
	rv := objc.Send[NeuralNetworksResolutionDependencies](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n NeuralNetworksResolutionDependencies) Autorelease() NeuralNetworksResolutionDependencies {
	rv := objc.Send[NeuralNetworksResolutionDependencies](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewNeuralNetworksResolutionDependencies creates a new NeuralNetworksResolutionDependencies instance.
func NewNeuralNetworksResolutionDependencies() NeuralNetworksResolutionDependencies {
	class := getNeuralNetworksResolutionDependenciesClass()
	rv := objc.Send[NeuralNetworksResolutionDependencies](objc.ID(class.class), objc.Sel("new"))
	return rv
}



































