// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NeuralNetworksLazyTensorFunctionBuilder] class.
var (
	_NeuralNetworksLazyTensorFunctionBuilderClass     NeuralNetworksLazyTensorFunctionBuilderClass
	_NeuralNetworksLazyTensorFunctionBuilderClassOnce sync.Once
)

func getNeuralNetworksLazyTensorFunctionBuilderClass() NeuralNetworksLazyTensorFunctionBuilderClass {
	_NeuralNetworksLazyTensorFunctionBuilderClassOnce.Do(func() {
		_NeuralNetworksLazyTensorFunctionBuilderClass = NeuralNetworksLazyTensorFunctionBuilderClass{class: objc.GetClass("NeuralNetworks.LazyTensorFunctionBuilder")}
	})
	return _NeuralNetworksLazyTensorFunctionBuilderClass
}

// GetNeuralNetworksLazyTensorFunctionBuilderClass returns the class object for NeuralNetworks.LazyTensorFunctionBuilder.
func GetNeuralNetworksLazyTensorFunctionBuilderClass() NeuralNetworksLazyTensorFunctionBuilderClass {
	return getNeuralNetworksLazyTensorFunctionBuilderClass()
}

type NeuralNetworksLazyTensorFunctionBuilderClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NeuralNetworksLazyTensorFunctionBuilderClass) Alloc() NeuralNetworksLazyTensorFunctionBuilder {
	rv := objc.Send[NeuralNetworksLazyTensorFunctionBuilder](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.LazyTensorFunctionBuilder
type NeuralNetworksLazyTensorFunctionBuilder struct {
	objectivec.Object
}

// NeuralNetworksLazyTensorFunctionBuilderFromID constructs a [NeuralNetworksLazyTensorFunctionBuilder] from an objc.ID.
func NeuralNetworksLazyTensorFunctionBuilderFromID(id objc.ID) NeuralNetworksLazyTensorFunctionBuilder {
	return NeuralNetworksLazyTensorFunctionBuilder{objectivec.Object{id}}
}
// Ensure NeuralNetworksLazyTensorFunctionBuilder implements INeuralNetworksLazyTensorFunctionBuilder.
var _ INeuralNetworksLazyTensorFunctionBuilder = NeuralNetworksLazyTensorFunctionBuilder{}





// An interface definition for the [NeuralNetworksLazyTensorFunctionBuilder] class.
//
// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.LazyTensorFunctionBuilder
type INeuralNetworksLazyTensorFunctionBuilder interface {
	objectivec.IObject
}





// Init initializes the instance.
func (n NeuralNetworksLazyTensorFunctionBuilder) Init() NeuralNetworksLazyTensorFunctionBuilder {
	rv := objc.Send[NeuralNetworksLazyTensorFunctionBuilder](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n NeuralNetworksLazyTensorFunctionBuilder) Autorelease() NeuralNetworksLazyTensorFunctionBuilder {
	rv := objc.Send[NeuralNetworksLazyTensorFunctionBuilder](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewNeuralNetworksLazyTensorFunctionBuilder creates a new NeuralNetworksLazyTensorFunctionBuilder instance.
func NewNeuralNetworksLazyTensorFunctionBuilder() NeuralNetworksLazyTensorFunctionBuilder {
	class := getNeuralNetworksLazyTensorFunctionBuilderClass()
	rv := objc.Send[NeuralNetworksLazyTensorFunctionBuilder](objc.ID(class.class), objc.Sel("new"))
	return rv
}



































