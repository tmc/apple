// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NeuralNetworksVariableIdentityGenerator] class.
var (
	_NeuralNetworksVariableIdentityGeneratorClass     NeuralNetworksVariableIdentityGeneratorClass
	_NeuralNetworksVariableIdentityGeneratorClassOnce sync.Once
)

func getNeuralNetworksVariableIdentityGeneratorClass() NeuralNetworksVariableIdentityGeneratorClass {
	_NeuralNetworksVariableIdentityGeneratorClassOnce.Do(func() {
		_NeuralNetworksVariableIdentityGeneratorClass = NeuralNetworksVariableIdentityGeneratorClass{class: objc.GetClass("NeuralNetworks.VariableIdentityGenerator")}
	})
	return _NeuralNetworksVariableIdentityGeneratorClass
}

// GetNeuralNetworksVariableIdentityGeneratorClass returns the class object for NeuralNetworks.VariableIdentityGenerator.
func GetNeuralNetworksVariableIdentityGeneratorClass() NeuralNetworksVariableIdentityGeneratorClass {
	return getNeuralNetworksVariableIdentityGeneratorClass()
}

type NeuralNetworksVariableIdentityGeneratorClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NeuralNetworksVariableIdentityGeneratorClass) Alloc() NeuralNetworksVariableIdentityGenerator {
	rv := objc.Send[NeuralNetworksVariableIdentityGenerator](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.VariableIdentityGenerator
type NeuralNetworksVariableIdentityGenerator struct {
	objectivec.Object
}

// NeuralNetworksVariableIdentityGeneratorFromID constructs a [NeuralNetworksVariableIdentityGenerator] from an objc.ID.
func NeuralNetworksVariableIdentityGeneratorFromID(id objc.ID) NeuralNetworksVariableIdentityGenerator {
	return NeuralNetworksVariableIdentityGenerator{objectivec.Object{id}}
}
// Ensure NeuralNetworksVariableIdentityGenerator implements INeuralNetworksVariableIdentityGenerator.
var _ INeuralNetworksVariableIdentityGenerator = NeuralNetworksVariableIdentityGenerator{}





// An interface definition for the [NeuralNetworksVariableIdentityGenerator] class.
//
// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.VariableIdentityGenerator
type INeuralNetworksVariableIdentityGenerator interface {
	objectivec.IObject
}





// Init initializes the instance.
func (n NeuralNetworksVariableIdentityGenerator) Init() NeuralNetworksVariableIdentityGenerator {
	rv := objc.Send[NeuralNetworksVariableIdentityGenerator](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n NeuralNetworksVariableIdentityGenerator) Autorelease() NeuralNetworksVariableIdentityGenerator {
	rv := objc.Send[NeuralNetworksVariableIdentityGenerator](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewNeuralNetworksVariableIdentityGenerator creates a new NeuralNetworksVariableIdentityGenerator instance.
func NewNeuralNetworksVariableIdentityGenerator() NeuralNetworksVariableIdentityGenerator {
	class := getNeuralNetworksVariableIdentityGeneratorClass()
	rv := objc.Send[NeuralNetworksVariableIdentityGenerator](objc.ID(class.class), objc.Sel("new"))
	return rv
}



































