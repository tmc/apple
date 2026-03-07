// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NeuralNetworksShuffleSampler] class.
var (
	_NeuralNetworksShuffleSamplerClass     NeuralNetworksShuffleSamplerClass
	_NeuralNetworksShuffleSamplerClassOnce sync.Once
)

func getNeuralNetworksShuffleSamplerClass() NeuralNetworksShuffleSamplerClass {
	_NeuralNetworksShuffleSamplerClassOnce.Do(func() {
		_NeuralNetworksShuffleSamplerClass = NeuralNetworksShuffleSamplerClass{class: objc.GetClass("NeuralNetworks.ShuffleSampler")}
	})
	return _NeuralNetworksShuffleSamplerClass
}

// GetNeuralNetworksShuffleSamplerClass returns the class object for NeuralNetworks.ShuffleSampler.
func GetNeuralNetworksShuffleSamplerClass() NeuralNetworksShuffleSamplerClass {
	return getNeuralNetworksShuffleSamplerClass()
}

type NeuralNetworksShuffleSamplerClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NeuralNetworksShuffleSamplerClass) Alloc() NeuralNetworksShuffleSampler {
	rv := objc.Send[NeuralNetworksShuffleSampler](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.ShuffleSampler
type NeuralNetworksShuffleSampler struct {
	objectivec.Object
}

// NeuralNetworksShuffleSamplerFromID constructs a [NeuralNetworksShuffleSampler] from an objc.ID.
func NeuralNetworksShuffleSamplerFromID(id objc.ID) NeuralNetworksShuffleSampler {
	return NeuralNetworksShuffleSampler{objectivec.Object{id}}
}
// Ensure NeuralNetworksShuffleSampler implements INeuralNetworksShuffleSampler.
var _ INeuralNetworksShuffleSampler = NeuralNetworksShuffleSampler{}





// An interface definition for the [NeuralNetworksShuffleSampler] class.
//
// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.ShuffleSampler
type INeuralNetworksShuffleSampler interface {
	objectivec.IObject
}





// Init initializes the instance.
func (n NeuralNetworksShuffleSampler) Init() NeuralNetworksShuffleSampler {
	rv := objc.Send[NeuralNetworksShuffleSampler](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n NeuralNetworksShuffleSampler) Autorelease() NeuralNetworksShuffleSampler {
	rv := objc.Send[NeuralNetworksShuffleSampler](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewNeuralNetworksShuffleSampler creates a new NeuralNetworksShuffleSampler instance.
func NewNeuralNetworksShuffleSampler() NeuralNetworksShuffleSampler {
	class := getNeuralNetworksShuffleSamplerClass()
	rv := objc.Send[NeuralNetworksShuffleSampler](objc.ID(class.class), objc.Sel("new"))
	return rv
}



































