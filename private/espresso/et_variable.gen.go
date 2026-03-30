// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [ETVariable] class.
var (
	_ETVariableClass     ETVariableClass
	_ETVariableClassOnce sync.Once
)

func getETVariableClass() ETVariableClass {
	_ETVariableClassOnce.Do(func() {
		_ETVariableClass = ETVariableClass{class: objc.GetClass("ETVariable")}
	})
	return _ETVariableClass
}

// GetETVariableClass returns the class object for ETVariable.
func GetETVariableClass() ETVariableClass {
	return getETVariableClass()
}

type ETVariableClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ec ETVariableClass) Class() objc.Class {
	return ec.class
}

// Alloc allocates memory for a new instance of the class.
func (ec ETVariableClass) Alloc() ETVariable {
	rv := objc.Send[ETVariable](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [ETVariable.CopyData]
//   - [ETVariable.GetOpaqueCopy]
//   - [ETVariable.InitializationAlpha]
//   - [ETVariable.SetInitializationAlpha]
//   - [ETVariable.InitializationBeta]
//   - [ETVariable.SetInitializationBeta]
//   - [ETVariable.InitializationMode]
//   - [ETVariable.SetInitializationMode]
//   - [ETVariable.Kind]
//   - [ETVariable.SetKind]
//   - [ETVariable.LayerName]
//   - [ETVariable.SetLayerName]
//   - [ETVariable.Model]
//   - [ETVariable.SetModel]
//   - [ETVariable.Name]
//   - [ETVariable.SetName]
//   - [ETVariable.SwapWithOpaqueCopy]
//   - [ETVariable.UpdateWithData]
//   - [ETVariable.InitWithModelDef]
//
// See: https://developer.apple.com/documentation/Espresso/ETVariable
type ETVariable struct {
	objectivec.Object
}

// ETVariableFromID constructs a [ETVariable] from an objc.ID.
func ETVariableFromID(id objc.ID) ETVariable {
	return ETVariable{objectivec.Object{ID: id}}
}

// Ensure ETVariable implements IETVariable.
var _ IETVariable = ETVariable{}

// An interface definition for the [ETVariable] class.
//
// # Methods
//
//   - [IETVariable.CopyData]
//   - [IETVariable.GetOpaqueCopy]
//   - [IETVariable.InitializationAlpha]
//   - [IETVariable.SetInitializationAlpha]
//   - [IETVariable.InitializationBeta]
//   - [IETVariable.SetInitializationBeta]
//   - [IETVariable.InitializationMode]
//   - [IETVariable.SetInitializationMode]
//   - [IETVariable.Kind]
//   - [IETVariable.SetKind]
//   - [IETVariable.LayerName]
//   - [IETVariable.SetLayerName]
//   - [IETVariable.Model]
//   - [IETVariable.SetModel]
//   - [IETVariable.Name]
//   - [IETVariable.SetName]
//   - [IETVariable.SwapWithOpaqueCopy]
//   - [IETVariable.UpdateWithData]
//   - [IETVariable.InitWithModelDef]
//
// See: https://developer.apple.com/documentation/Espresso/ETVariable
type IETVariable interface {
	objectivec.IObject

	// Topic: Methods

	CopyData() objectivec.IObject
	GetOpaqueCopy() objectivec.IObject
	InitializationAlpha() float32
	SetInitializationAlpha(value float32)
	InitializationBeta() float32
	SetInitializationBeta(value float32)
	InitializationMode() uint64
	SetInitializationMode(value uint64)
	Kind() uint64
	SetKind(value uint64)
	LayerName() string
	SetLayerName(value string)
	Model() IETModelDef
	SetModel(value IETModelDef)
	Name() string
	SetName(value string)
	SwapWithOpaqueCopy(copy_ objectivec.IObject) objectivec.IObject
	UpdateWithData(data objectivec.IObject)
	InitWithModelDef(def objectivec.IObject) ETVariable
}

// Init initializes the instance.
func (e ETVariable) Init() ETVariable {
	rv := objc.Send[ETVariable](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e ETVariable) Autorelease() ETVariable {
	rv := objc.Send[ETVariable](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewETVariable creates a new ETVariable instance.
func NewETVariable() ETVariable {
	class := getETVariableClass()
	rv := objc.Send[ETVariable](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Espresso/ETVariable/initWithModelDef:
func NewETVariableWithModelDef(def objectivec.IObject) ETVariable {
	instance := getETVariableClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithModelDef:"), def)
	return ETVariableFromID(rv)
}

// See: https://developer.apple.com/documentation/Espresso/ETVariable/copyData
func (e ETVariable) CopyData() objectivec.IObject {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("copyData"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/Espresso/ETVariable/getOpaqueCopy
func (e ETVariable) GetOpaqueCopy() objectivec.IObject {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("getOpaqueCopy"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/Espresso/ETVariable/swapWithOpaqueCopy:
func (e ETVariable) SwapWithOpaqueCopy(copy_ objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("swapWithOpaqueCopy:"), copy_)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/Espresso/ETVariable/updateWithData:
func (e ETVariable) UpdateWithData(data objectivec.IObject) {
	objc.Send[objc.ID](e.ID, objc.Sel("updateWithData:"), data)
}

// See: https://developer.apple.com/documentation/Espresso/ETVariable/initWithModelDef:
func (e ETVariable) InitWithModelDef(def objectivec.IObject) ETVariable {
	rv := objc.Send[ETVariable](e.ID, objc.Sel("initWithModelDef:"), def)
	return rv
}

// See: https://developer.apple.com/documentation/Espresso/ETVariable/initializationAlpha
func (e ETVariable) InitializationAlpha() float32 {
	rv := objc.Send[float32](e.ID, objc.Sel("initializationAlpha"))
	return rv
}
func (e ETVariable) SetInitializationAlpha(value float32) {
	objc.Send[struct{}](e.ID, objc.Sel("setInitializationAlpha:"), value)
}

// See: https://developer.apple.com/documentation/Espresso/ETVariable/initializationBeta
func (e ETVariable) InitializationBeta() float32 {
	rv := objc.Send[float32](e.ID, objc.Sel("initializationBeta"))
	return rv
}
func (e ETVariable) SetInitializationBeta(value float32) {
	objc.Send[struct{}](e.ID, objc.Sel("setInitializationBeta:"), value)
}

// See: https://developer.apple.com/documentation/Espresso/ETVariable/initializationMode
func (e ETVariable) InitializationMode() uint64 {
	rv := objc.Send[uint64](e.ID, objc.Sel("initializationMode"))
	return rv
}
func (e ETVariable) SetInitializationMode(value uint64) {
	objc.Send[struct{}](e.ID, objc.Sel("setInitializationMode:"), value)
}

// See: https://developer.apple.com/documentation/Espresso/ETVariable/kind
func (e ETVariable) Kind() uint64 {
	rv := objc.Send[uint64](e.ID, objc.Sel("kind"))
	return rv
}
func (e ETVariable) SetKind(value uint64) {
	objc.Send[struct{}](e.ID, objc.Sel("setKind:"), value)
}

// See: https://developer.apple.com/documentation/Espresso/ETVariable/layerName
func (e ETVariable) LayerName() string {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("layerName"))
	return foundation.NSStringFromID(rv).String()
}
func (e ETVariable) SetLayerName(value string) {
	objc.Send[struct{}](e.ID, objc.Sel("setLayerName:"), objc.String(value))
}

// See: https://developer.apple.com/documentation/Espresso/ETVariable/model
func (e ETVariable) Model() IETModelDef {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("model"))
	return ETModelDefFromID(objc.ID(rv))
}
func (e ETVariable) SetModel(value IETModelDef) {
	objc.Send[struct{}](e.ID, objc.Sel("setModel:"), value)
}

// See: https://developer.apple.com/documentation/Espresso/ETVariable/name
func (e ETVariable) Name() string {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("name"))
	return foundation.NSStringFromID(rv).String()
}
func (e ETVariable) SetName(value string) {
	objc.Send[struct{}](e.ID, objc.Sel("setName:"), objc.String(value))
}
