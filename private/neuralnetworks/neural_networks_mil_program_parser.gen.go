// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

package neuralnetworks

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NeuralNetworksMILProgramParser] class.
var (
	_NeuralNetworksMILProgramParserClass     NeuralNetworksMILProgramParserClass
	_NeuralNetworksMILProgramParserClassOnce sync.Once
)

func getNeuralNetworksMILProgramParserClass() NeuralNetworksMILProgramParserClass {
	_NeuralNetworksMILProgramParserClassOnce.Do(func() {
		_NeuralNetworksMILProgramParserClass = NeuralNetworksMILProgramParserClass{class: objc.GetClass("NeuralNetworks.MILProgramParser")}
	})
	return _NeuralNetworksMILProgramParserClass
}

// GetNeuralNetworksMILProgramParserClass returns the class object for NeuralNetworks.MILProgramParser.
func GetNeuralNetworksMILProgramParserClass() NeuralNetworksMILProgramParserClass {
	return getNeuralNetworksMILProgramParserClass()
}

type NeuralNetworksMILProgramParserClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NeuralNetworksMILProgramParserClass) Alloc() NeuralNetworksMILProgramParser {
	rv := objc.Send[NeuralNetworksMILProgramParser](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.MILProgramParser
type NeuralNetworksMILProgramParser struct {
	objectivec.Object
}

// NeuralNetworksMILProgramParserFromID constructs a [NeuralNetworksMILProgramParser] from an objc.ID.
func NeuralNetworksMILProgramParserFromID(id objc.ID) NeuralNetworksMILProgramParser {
	return NeuralNetworksMILProgramParser{objectivec.Object{id}}
}
// Ensure NeuralNetworksMILProgramParser implements INeuralNetworksMILProgramParser.
var _ INeuralNetworksMILProgramParser = NeuralNetworksMILProgramParser{}





// An interface definition for the [NeuralNetworksMILProgramParser] class.
//
// See: https://developer.apple.com/documentation/NeuralNetworks/NeuralNetworks.MILProgramParser
type INeuralNetworksMILProgramParser interface {
	objectivec.IObject
}





// Init initializes the instance.
func (n NeuralNetworksMILProgramParser) Init() NeuralNetworksMILProgramParser {
	rv := objc.Send[NeuralNetworksMILProgramParser](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n NeuralNetworksMILProgramParser) Autorelease() NeuralNetworksMILProgramParser {
	rv := objc.Send[NeuralNetworksMILProgramParser](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewNeuralNetworksMILProgramParser creates a new NeuralNetworksMILProgramParser instance.
func NewNeuralNetworksMILProgramParser() NeuralNetworksMILProgramParser {
	class := getNeuralNetworksMILProgramParserClass()
	rv := objc.Send[NeuralNetworksMILProgramParser](objc.ID(class.class), objc.Sel("new"))
	return rv
}



































