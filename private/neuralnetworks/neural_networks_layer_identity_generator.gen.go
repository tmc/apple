// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NeuralNetworksLayerIdentityGenerator] class.
var (
	_NeuralNetworksLayerIdentityGeneratorClass     NeuralNetworksLayerIdentityGeneratorClass
	_NeuralNetworksLayerIdentityGeneratorClassOnce sync.Once
)

func getNeuralNetworksLayerIdentityGeneratorClass() NeuralNetworksLayerIdentityGeneratorClass {
	_NeuralNetworksLayerIdentityGeneratorClassOnce.Do(func() {
		_NeuralNetworksLayerIdentityGeneratorClass = NeuralNetworksLayerIdentityGeneratorClass{class: objc.GetClass("NeuralNetworks.LayerIdentityGenerator")}
	})
	return _NeuralNetworksLayerIdentityGeneratorClass
}

// GetNeuralNetworksLayerIdentityGeneratorClass returns the class object for NeuralNetworks.LayerIdentityGenerator.
func GetNeuralNetworksLayerIdentityGeneratorClass() NeuralNetworksLayerIdentityGeneratorClass {
	return getNeuralNetworksLayerIdentityGeneratorClass()
}

type NeuralNetworksLayerIdentityGeneratorClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NeuralNetworksLayerIdentityGeneratorClass) Alloc() NeuralNetworksLayerIdentityGenerator {
	rv := objc.Send[NeuralNetworksLayerIdentityGenerator](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.LayerIdentityGenerator
type NeuralNetworksLayerIdentityGenerator struct {
	objectivec.Object
}

// NeuralNetworksLayerIdentityGeneratorFromID constructs a [NeuralNetworksLayerIdentityGenerator] from an objc.ID.
func NeuralNetworksLayerIdentityGeneratorFromID(id objc.ID) NeuralNetworksLayerIdentityGenerator {
	return NeuralNetworksLayerIdentityGenerator{objectivec.Object{id}}
}
// Ensure NeuralNetworksLayerIdentityGenerator implements INeuralNetworksLayerIdentityGenerator.
var _ INeuralNetworksLayerIdentityGenerator = NeuralNetworksLayerIdentityGenerator{}





// An interface definition for the [NeuralNetworksLayerIdentityGenerator] class.
//
// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.LayerIdentityGenerator
type INeuralNetworksLayerIdentityGenerator interface {
	objectivec.IObject
}





// Init initializes the instance.
func (n NeuralNetworksLayerIdentityGenerator) Init() NeuralNetworksLayerIdentityGenerator {
	rv := objc.Send[NeuralNetworksLayerIdentityGenerator](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n NeuralNetworksLayerIdentityGenerator) Autorelease() NeuralNetworksLayerIdentityGenerator {
	rv := objc.Send[NeuralNetworksLayerIdentityGenerator](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewNeuralNetworksLayerIdentityGenerator creates a new NeuralNetworksLayerIdentityGenerator instance.
func NewNeuralNetworksLayerIdentityGenerator() NeuralNetworksLayerIdentityGenerator {
	class := getNeuralNetworksLayerIdentityGeneratorClass()
	rv := objc.Send[NeuralNetworksLayerIdentityGenerator](objc.ID(class.class), objc.Sel("new"))
	return rv
}



































