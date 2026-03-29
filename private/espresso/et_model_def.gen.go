// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [ETModelDef] class.
var (
	_ETModelDefClass     ETModelDefClass
	_ETModelDefClassOnce sync.Once
)

func getETModelDefClass() ETModelDefClass {
	_ETModelDefClassOnce.Do(func() {
		_ETModelDefClass = ETModelDefClass{class: objc.GetClass("ETModelDef")}
	})
	return _ETModelDefClass
}

// GetETModelDefClass returns the class object for ETModelDef.
func GetETModelDefClass() ETModelDefClass {
	return getETModelDefClass()
}

type ETModelDefClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ec ETModelDefClass) Class() objc.Class {
	return ec.class
}

// Alloc allocates memory for a new instance of the class.
func (ec ETModelDefClass) Alloc() ETModelDef {
	rv := objc.Send[ETModelDef](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [ETModelDef.All_variables]
//   - [ETModelDef.SetAll_variables]
//   - [ETModelDef.BiasesForLayer]
//   - [ETModelDef.ConfigureLayersToTrainReinitializeVariables]
//   - [ETModelDef.Gb]
//   - [ETModelDef.SetGb]
//   - [ETModelDef.LayerForName]
//   - [ETModelDef.LayerNames]
//   - [ETModelDef.Network]
//   - [ETModelDef.SetNetwork]
//   - [ETModelDef.RandomizeWeightsForLayerWithSeed]
//   - [ETModelDef.SetupVariablesDef]
//   - [ETModelDef.TopNamesForLayerIndex]
//   - [ETModelDef.TransformForTraining]
//   - [ETModelDef.UpdateLayerWithBiasesLength]
//   - [ETModelDef.UpdateLayerWithWeightsLength]
//   - [ETModelDef.VariableForLayerKind]
//   - [ETModelDef.VariableNameForLayerKind]
//   - [ETModelDef.Variables]
//   - [ETModelDef.WeightsForLayer]
//   - [ETModelDef.InitWithNetwork]
// See: https://developer.apple.com/documentation/Espresso/ETModelDef
type ETModelDef struct {
	objectivec.Object
}

// ETModelDefFromID constructs a [ETModelDef] from an objc.ID.
func ETModelDefFromID(id objc.ID) ETModelDef {
	return ETModelDef{objectivec.Object{ID: id}}
}
// Ensure ETModelDef implements IETModelDef.
var _ IETModelDef = ETModelDef{}

// An interface definition for the [ETModelDef] class.
//
// # Methods
//
//   - [IETModelDef.All_variables]
//   - [IETModelDef.SetAll_variables]
//   - [IETModelDef.BiasesForLayer]
//   - [IETModelDef.ConfigureLayersToTrainReinitializeVariables]
//   - [IETModelDef.Gb]
//   - [IETModelDef.SetGb]
//   - [IETModelDef.LayerForName]
//   - [IETModelDef.LayerNames]
//   - [IETModelDef.Network]
//   - [IETModelDef.SetNetwork]
//   - [IETModelDef.RandomizeWeightsForLayerWithSeed]
//   - [IETModelDef.SetupVariablesDef]
//   - [IETModelDef.TopNamesForLayerIndex]
//   - [IETModelDef.TransformForTraining]
//   - [IETModelDef.UpdateLayerWithBiasesLength]
//   - [IETModelDef.UpdateLayerWithWeightsLength]
//   - [IETModelDef.VariableForLayerKind]
//   - [IETModelDef.VariableNameForLayerKind]
//   - [IETModelDef.Variables]
//   - [IETModelDef.WeightsForLayer]
//   - [IETModelDef.InitWithNetwork]
//
// See: https://developer.apple.com/documentation/Espresso/ETModelDef
type IETModelDef interface {
	objectivec.IObject

	// Topic: Methods

	All_variables() foundation.INSArray
	SetAll_variables(value foundation.INSArray)
	BiasesForLayer(layer objectivec.IObject) objectivec.IObject
	ConfigureLayersToTrainReinitializeVariables(train objectivec.IObject, variables bool) int
	Gb() objectivec.IObject
	SetGb(value objectivec.IObject)
	LayerForName(name objectivec.IObject) unsafe.Pointer
	LayerNames() objectivec.IObject
	Network() objectivec.IObject
	SetNetwork(value objectivec.IObject)
	RandomizeWeightsForLayerWithSeed(layer objectivec.IObject, seed float32)
	SetupVariablesDef()
	TopNamesForLayerIndex(index int) objectivec.IObject
	TransformForTraining(training objectivec.IObject)
	UpdateLayerWithBiasesLength(layer objectivec.IObject, biases objectivec.IObject, length uint64)
	UpdateLayerWithWeightsLength(layer objectivec.IObject, weights objectivec.IObject, length uint64)
	VariableForLayerKind(layer objectivec.IObject, kind uint64) objectivec.IObject
	VariableNameForLayerKind(layer objectivec.IObject, kind uint64) objectivec.IObject
	Variables() objectivec.IObject
	WeightsForLayer(layer objectivec.IObject) objectivec.IObject
	InitWithNetwork(network objectivec.IObject) ETModelDef
}

// Init initializes the instance.
func (e ETModelDef) Init() ETModelDef {
	rv := objc.Send[ETModelDef](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e ETModelDef) Autorelease() ETModelDef {
	rv := objc.Send[ETModelDef](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewETModelDef creates a new ETModelDef instance.
func NewETModelDef() ETModelDef {
	class := getETModelDefClass()
	rv := objc.Send[ETModelDef](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/Espresso/ETModelDef/initWithNetwork:
func NewETModelDefWithNetwork(network objectivec.IObject) ETModelDef {
	instance := getETModelDefClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithNetwork:"), network)
	return ETModelDefFromID(rv)
}

//
// See: https://developer.apple.com/documentation/Espresso/ETModelDef/biasesForLayer:
func (e ETModelDef) BiasesForLayer(layer objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("biasesForLayer:"), layer)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/Espresso/ETModelDef/configureLayersToTrain:reinitializeVariables:
func (e ETModelDef) ConfigureLayersToTrainReinitializeVariables(train objectivec.IObject, variables bool) int {
	rv := objc.Send[int](e.ID, objc.Sel("configureLayersToTrain:reinitializeVariables:"), train, variables)
	return rv
}
//
// See: https://developer.apple.com/documentation/Espresso/ETModelDef/layerForName:
func (e ETModelDef) LayerForName(name objectivec.IObject) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e.ID, objc.Sel("layerForName:"), name)
	return rv
}
// See: https://developer.apple.com/documentation/Espresso/ETModelDef/layerNames
func (e ETModelDef) LayerNames() objectivec.IObject {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("layerNames"))
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/Espresso/ETModelDef/randomizeWeightsForLayer:withSeed:
func (e ETModelDef) RandomizeWeightsForLayerWithSeed(layer objectivec.IObject, seed float32) {
	objc.Send[objc.ID](e.ID, objc.Sel("randomizeWeightsForLayer:withSeed:"), layer, seed)
}
// See: https://developer.apple.com/documentation/Espresso/ETModelDef/setupVariablesDef
func (e ETModelDef) SetupVariablesDef() {
	objc.Send[objc.ID](e.ID, objc.Sel("setupVariablesDef"))
}
//
// See: https://developer.apple.com/documentation/Espresso/ETModelDef/topNamesForLayerIndex:
func (e ETModelDef) TopNamesForLayerIndex(index int) objectivec.IObject {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("topNamesForLayerIndex:"), index)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/Espresso/ETModelDef/transformForTraining:
func (e ETModelDef) TransformForTraining(training objectivec.IObject) {
	objc.Send[objc.ID](e.ID, objc.Sel("transformForTraining:"), training)
}
//
// See: https://developer.apple.com/documentation/Espresso/ETModelDef/updateLayer:withBiases:length:
func (e ETModelDef) UpdateLayerWithBiasesLength(layer objectivec.IObject, biases objectivec.IObject, length uint64) {
	objc.Send[objc.ID](e.ID, objc.Sel("updateLayer:withBiases:length:"), layer, biases, length)
}
//
// See: https://developer.apple.com/documentation/Espresso/ETModelDef/updateLayer:withWeights:length:
func (e ETModelDef) UpdateLayerWithWeightsLength(layer objectivec.IObject, weights objectivec.IObject, length uint64) {
	objc.Send[objc.ID](e.ID, objc.Sel("updateLayer:withWeights:length:"), layer, weights, length)
}
//
// See: https://developer.apple.com/documentation/Espresso/ETModelDef/variableForLayer:kind:
func (e ETModelDef) VariableForLayerKind(layer objectivec.IObject, kind uint64) objectivec.IObject {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("variableForLayer:kind:"), layer, kind)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/Espresso/ETModelDef/variableNameForLayer:kind:
func (e ETModelDef) VariableNameForLayerKind(layer objectivec.IObject, kind uint64) objectivec.IObject {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("variableNameForLayer:kind:"), layer, kind)
	return objectivec.Object{ID: rv}
}
// See: https://developer.apple.com/documentation/Espresso/ETModelDef/variables
func (e ETModelDef) Variables() objectivec.IObject {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("variables"))
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/Espresso/ETModelDef/weightsForLayer:
func (e ETModelDef) WeightsForLayer(layer objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("weightsForLayer:"), layer)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/Espresso/ETModelDef/initWithNetwork:
func (e ETModelDef) InitWithNetwork(network objectivec.IObject) ETModelDef {
	rv := objc.Send[ETModelDef](e.ID, objc.Sel("initWithNetwork:"), network)
	return rv
}

// See: https://developer.apple.com/documentation/Espresso/ETModelDef/all_variables
func (e ETModelDef) All_variables() foundation.INSArray {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("all_variables"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (e ETModelDef) SetAll_variables(value foundation.INSArray) {
	objc.Send[struct{}](e.ID, objc.Sel("setAll_variables:"), value)
}
// See: https://developer.apple.com/documentation/Espresso/ETModelDef/gb
func (e ETModelDef) Gb() objectivec.IObject {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("gb"))
	return objectivec.Object{ID: rv}
}
func (e ETModelDef) SetGb(value objectivec.IObject) {
	objc.Send[struct{}](e.ID, objc.Sel("setGb:"), value)
}
// See: https://developer.apple.com/documentation/Espresso/ETModelDef/network
func (e ETModelDef) Network() objectivec.IObject {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("network"))
	return objectivec.Object{ID: rv}
}
func (e ETModelDef) SetNetwork(value objectivec.IObject) {
	objc.Send[struct{}](e.ID, objc.Sel("setNetwork:"), value)
}

