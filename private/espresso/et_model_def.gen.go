// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
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

// # Methods
//
//   - [ETModelDef.All_variables]
//   - [ETModelDef.SetAll_variables]
//   - [ETModelDef.ConfigureLayersToTrainReinitializeVariables]
//   - [ETModelDef.Gb]
//   - [ETModelDef.SetGb]
//   - [ETModelDef.LayerForName]
//   - [ETModelDef.LayerNames]
//   - [ETModelDef.Network]
//   - [ETModelDef.SetNetwork]
//   - [ETModelDef.RandomizeWeightsForLayerWithSeed]
//   - [ETModelDef.SetupVariablesDef]
//   - [ETModelDef.UpdateLayerWithBiasesLength]
//   - [ETModelDef.UpdateLayerWithWeightsLength]
//   - [ETModelDef.VariableForLayerKind]
//   - [ETModelDef.VariableNameForLayerKind]
//   - [ETModelDef.Variables]
//   - [ETModelDef.InitWithNetwork]
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
//   - [IETModelDef.ConfigureLayersToTrainReinitializeVariables]
//   - [IETModelDef.Gb]
//   - [IETModelDef.SetGb]
//   - [IETModelDef.LayerForName]
//   - [IETModelDef.LayerNames]
//   - [IETModelDef.Network]
//   - [IETModelDef.SetNetwork]
//   - [IETModelDef.RandomizeWeightsForLayerWithSeed]
//   - [IETModelDef.SetupVariablesDef]
//   - [IETModelDef.UpdateLayerWithBiasesLength]
//   - [IETModelDef.UpdateLayerWithWeightsLength]
//   - [IETModelDef.VariableForLayerKind]
//   - [IETModelDef.VariableNameForLayerKind]
//   - [IETModelDef.Variables]
//   - [IETModelDef.InitWithNetwork]
type IETModelDef interface {
	objectivec.IObject

	// Topic: Methods

	All_variables() foundation.INSArray
	SetAll_variables(value foundation.INSArray)
	ConfigureLayersToTrainReinitializeVariables(train objectivec.IObject, variables bool) int
	Gb() unsafe.Pointer
	SetGb(value unsafe.Pointer)
	LayerForName(name objectivec.IObject) unsafe.Pointer
	LayerNames() objectivec.IObject
	Network() unsafe.Pointer
	SetNetwork(value unsafe.Pointer)
	RandomizeWeightsForLayerWithSeed(layer objectivec.IObject, seed float32)
	SetupVariablesDef()
	UpdateLayerWithBiasesLength(layer objectivec.IObject, biases unsafe.Pointer, length uint64)
	UpdateLayerWithWeightsLength(layer objectivec.IObject, weights unsafe.Pointer, length uint64)
	VariableForLayerKind(layer objectivec.IObject, kind uint64) objectivec.IObject
	VariableNameForLayerKind(layer objectivec.IObject, kind uint64) objectivec.IObject
	Variables() objectivec.IObject
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

func NewETModelDefWithNetwork(network objectivec.IObject) ETModelDef {
	instance := getETModelDefClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithNetwork:"), network)
	return ETModelDefFromID(rv)
}

func (e ETModelDef) ConfigureLayersToTrainReinitializeVariables(train objectivec.IObject, variables bool) int {
	rv := objc.Send[int](e.ID, objc.Sel("configureLayersToTrain:reinitializeVariables:"), train, variables)
	return rv
}
func (e ETModelDef) LayerForName(name objectivec.IObject) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e.ID, objc.Sel("layerForName:"), name)
	return rv
}
func (e ETModelDef) LayerNames() objectivec.IObject {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("layerNames"))
	return objectivec.Object{ID: rv}
}
func (e ETModelDef) RandomizeWeightsForLayerWithSeed(layer objectivec.IObject, seed float32) {
	objc.Send[objc.ID](e.ID, objc.Sel("randomizeWeightsForLayer:withSeed:"), layer, seed)
}
func (e ETModelDef) SetupVariablesDef() {
	objc.Send[objc.ID](e.ID, objc.Sel("setupVariablesDef"))
}
func (e ETModelDef) UpdateLayerWithBiasesLength(layer objectivec.IObject, biases unsafe.Pointer, length uint64) {
	objc.Send[objc.ID](e.ID, objc.Sel("updateLayer:withBiases:length:"), layer, biases, length)
}
func (e ETModelDef) UpdateLayerWithWeightsLength(layer objectivec.IObject, weights unsafe.Pointer, length uint64) {
	objc.Send[objc.ID](e.ID, objc.Sel("updateLayer:withWeights:length:"), layer, weights, length)
}
func (e ETModelDef) VariableForLayerKind(layer objectivec.IObject, kind uint64) objectivec.IObject {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("variableForLayer:kind:"), layer, kind)
	return objectivec.Object{ID: rv}
}
func (e ETModelDef) VariableNameForLayerKind(layer objectivec.IObject, kind uint64) objectivec.IObject {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("variableNameForLayer:kind:"), layer, kind)
	return objectivec.Object{ID: rv}
}
func (e ETModelDef) Variables() objectivec.IObject {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("variables"))
	return objectivec.Object{ID: rv}
}
func (e ETModelDef) InitWithNetwork(network objectivec.IObject) ETModelDef {
	rv := objc.Send[ETModelDef](e.ID, objc.Sel("initWithNetwork:"), network)
	return rv
}

func (e ETModelDef) All_variables() foundation.INSArray {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("all_variables"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (e ETModelDef) SetAll_variables(value foundation.INSArray) {
	objc.Send[struct{}](e.ID, objc.Sel("setAll_variables:"), value)
}
func (e ETModelDef) Gb() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e.ID, objc.Sel("gb"))
	return rv
}
func (e ETModelDef) SetGb(value unsafe.Pointer) {
	objc.Send[struct{}](e.ID, objc.Sel("setGb:"), value)
}
func (e ETModelDef) Network() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e.ID, objc.Sel("network"))
	return rv
}
func (e ETModelDef) SetNetwork(value unsafe.Pointer) {
	objc.Send[struct{}](e.ID, objc.Sel("setNetwork:"), value)
}
