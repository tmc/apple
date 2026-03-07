// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NeuralNetworksTensorOperationIdentityGenerator] class.
var (
	_NeuralNetworksTensorOperationIdentityGeneratorClass     NeuralNetworksTensorOperationIdentityGeneratorClass
	_NeuralNetworksTensorOperationIdentityGeneratorClassOnce sync.Once
)

func getNeuralNetworksTensorOperationIdentityGeneratorClass() NeuralNetworksTensorOperationIdentityGeneratorClass {
	_NeuralNetworksTensorOperationIdentityGeneratorClassOnce.Do(func() {
		_NeuralNetworksTensorOperationIdentityGeneratorClass = NeuralNetworksTensorOperationIdentityGeneratorClass{class: objc.GetClass("NeuralNetworks.TensorOperationIdentityGenerator")}
	})
	return _NeuralNetworksTensorOperationIdentityGeneratorClass
}

// GetNeuralNetworksTensorOperationIdentityGeneratorClass returns the class object for NeuralNetworks.TensorOperationIdentityGenerator.
func GetNeuralNetworksTensorOperationIdentityGeneratorClass() NeuralNetworksTensorOperationIdentityGeneratorClass {
	return getNeuralNetworksTensorOperationIdentityGeneratorClass()
}

type NeuralNetworksTensorOperationIdentityGeneratorClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NeuralNetworksTensorOperationIdentityGeneratorClass) Alloc() NeuralNetworksTensorOperationIdentityGenerator {
	rv := objc.Send[NeuralNetworksTensorOperationIdentityGenerator](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.TensorOperationIdentityGenerator
type NeuralNetworksTensorOperationIdentityGenerator struct {
	objectivec.Object
}

// NeuralNetworksTensorOperationIdentityGeneratorFromID constructs a [NeuralNetworksTensorOperationIdentityGenerator] from an objc.ID.
func NeuralNetworksTensorOperationIdentityGeneratorFromID(id objc.ID) NeuralNetworksTensorOperationIdentityGenerator {
	return NeuralNetworksTensorOperationIdentityGenerator{objectivec.Object{id}}
}
// Ensure NeuralNetworksTensorOperationIdentityGenerator implements INeuralNetworksTensorOperationIdentityGenerator.
var _ INeuralNetworksTensorOperationIdentityGenerator = NeuralNetworksTensorOperationIdentityGenerator{}





// An interface definition for the [NeuralNetworksTensorOperationIdentityGenerator] class.
//
// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.TensorOperationIdentityGenerator
type INeuralNetworksTensorOperationIdentityGenerator interface {
	objectivec.IObject
}





// Init initializes the instance.
func (n NeuralNetworksTensorOperationIdentityGenerator) Init() NeuralNetworksTensorOperationIdentityGenerator {
	rv := objc.Send[NeuralNetworksTensorOperationIdentityGenerator](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n NeuralNetworksTensorOperationIdentityGenerator) Autorelease() NeuralNetworksTensorOperationIdentityGenerator {
	rv := objc.Send[NeuralNetworksTensorOperationIdentityGenerator](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewNeuralNetworksTensorOperationIdentityGenerator creates a new NeuralNetworksTensorOperationIdentityGenerator instance.
func NewNeuralNetworksTensorOperationIdentityGenerator() NeuralNetworksTensorOperationIdentityGenerator {
	class := getNeuralNetworksTensorOperationIdentityGeneratorClass()
	rv := objc.Send[NeuralNetworksTensorOperationIdentityGenerator](objc.ID(class.class), objc.Sel("new"))
	return rv
}



































