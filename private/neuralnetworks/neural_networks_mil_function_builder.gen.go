// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NeuralNetworksMILFunctionBuilder] class.
var (
	_NeuralNetworksMILFunctionBuilderClass     NeuralNetworksMILFunctionBuilderClass
	_NeuralNetworksMILFunctionBuilderClassOnce sync.Once
)

func getNeuralNetworksMILFunctionBuilderClass() NeuralNetworksMILFunctionBuilderClass {
	_NeuralNetworksMILFunctionBuilderClassOnce.Do(func() {
		_NeuralNetworksMILFunctionBuilderClass = NeuralNetworksMILFunctionBuilderClass{class: objc.GetClass("NeuralNetworks.MILFunctionBuilder")}
	})
	return _NeuralNetworksMILFunctionBuilderClass
}

// GetNeuralNetworksMILFunctionBuilderClass returns the class object for NeuralNetworks.MILFunctionBuilder.
func GetNeuralNetworksMILFunctionBuilderClass() NeuralNetworksMILFunctionBuilderClass {
	return getNeuralNetworksMILFunctionBuilderClass()
}

type NeuralNetworksMILFunctionBuilderClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NeuralNetworksMILFunctionBuilderClass) Alloc() NeuralNetworksMILFunctionBuilder {
	rv := objc.Send[NeuralNetworksMILFunctionBuilder](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.MILFunctionBuilder
type NeuralNetworksMILFunctionBuilder struct {
	objectivec.Object
}

// NeuralNetworksMILFunctionBuilderFromID constructs a [NeuralNetworksMILFunctionBuilder] from an objc.ID.
func NeuralNetworksMILFunctionBuilderFromID(id objc.ID) NeuralNetworksMILFunctionBuilder {
	return NeuralNetworksMILFunctionBuilder{objectivec.Object{id}}
}
// Ensure NeuralNetworksMILFunctionBuilder implements INeuralNetworksMILFunctionBuilder.
var _ INeuralNetworksMILFunctionBuilder = NeuralNetworksMILFunctionBuilder{}





// An interface definition for the [NeuralNetworksMILFunctionBuilder] class.
//
// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.MILFunctionBuilder
type INeuralNetworksMILFunctionBuilder interface {
	objectivec.IObject
}





// Init initializes the instance.
func (n NeuralNetworksMILFunctionBuilder) Init() NeuralNetworksMILFunctionBuilder {
	rv := objc.Send[NeuralNetworksMILFunctionBuilder](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n NeuralNetworksMILFunctionBuilder) Autorelease() NeuralNetworksMILFunctionBuilder {
	rv := objc.Send[NeuralNetworksMILFunctionBuilder](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewNeuralNetworksMILFunctionBuilder creates a new NeuralNetworksMILFunctionBuilder instance.
func NewNeuralNetworksMILFunctionBuilder() NeuralNetworksMILFunctionBuilder {
	class := getNeuralNetworksMILFunctionBuilderClass()
	rv := objc.Send[NeuralNetworksMILFunctionBuilder](objc.ID(class.class), objc.Sel("new"))
	return rv
}



































