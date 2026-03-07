// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NeuralNetworksMILProgramBuilder] class.
var (
	_NeuralNetworksMILProgramBuilderClass     NeuralNetworksMILProgramBuilderClass
	_NeuralNetworksMILProgramBuilderClassOnce sync.Once
)

func getNeuralNetworksMILProgramBuilderClass() NeuralNetworksMILProgramBuilderClass {
	_NeuralNetworksMILProgramBuilderClassOnce.Do(func() {
		_NeuralNetworksMILProgramBuilderClass = NeuralNetworksMILProgramBuilderClass{class: objc.GetClass("NeuralNetworks.MILProgramBuilder")}
	})
	return _NeuralNetworksMILProgramBuilderClass
}

// GetNeuralNetworksMILProgramBuilderClass returns the class object for NeuralNetworks.MILProgramBuilder.
func GetNeuralNetworksMILProgramBuilderClass() NeuralNetworksMILProgramBuilderClass {
	return getNeuralNetworksMILProgramBuilderClass()
}

type NeuralNetworksMILProgramBuilderClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NeuralNetworksMILProgramBuilderClass) Alloc() NeuralNetworksMILProgramBuilder {
	rv := objc.Send[NeuralNetworksMILProgramBuilder](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.MILProgramBuilder
type NeuralNetworksMILProgramBuilder struct {
	objectivec.Object
}

// NeuralNetworksMILProgramBuilderFromID constructs a [NeuralNetworksMILProgramBuilder] from an objc.ID.
func NeuralNetworksMILProgramBuilderFromID(id objc.ID) NeuralNetworksMILProgramBuilder {
	return NeuralNetworksMILProgramBuilder{objectivec.Object{id}}
}
// Ensure NeuralNetworksMILProgramBuilder implements INeuralNetworksMILProgramBuilder.
var _ INeuralNetworksMILProgramBuilder = NeuralNetworksMILProgramBuilder{}





// An interface definition for the [NeuralNetworksMILProgramBuilder] class.
//
// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.MILProgramBuilder
type INeuralNetworksMILProgramBuilder interface {
	objectivec.IObject
}





// Init initializes the instance.
func (n NeuralNetworksMILProgramBuilder) Init() NeuralNetworksMILProgramBuilder {
	rv := objc.Send[NeuralNetworksMILProgramBuilder](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n NeuralNetworksMILProgramBuilder) Autorelease() NeuralNetworksMILProgramBuilder {
	rv := objc.Send[NeuralNetworksMILProgramBuilder](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewNeuralNetworksMILProgramBuilder creates a new NeuralNetworksMILProgramBuilder instance.
func NewNeuralNetworksMILProgramBuilder() NeuralNetworksMILProgramBuilder {
	class := getNeuralNetworksMILProgramBuilderClass()
	rv := objc.Send[NeuralNetworksMILProgramBuilder](objc.ID(class.class), objc.Sel("new"))
	return rv
}



































