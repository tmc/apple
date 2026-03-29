// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [ETLossDefinition] class.
var (
	_ETLossDefinitionClass     ETLossDefinitionClass
	_ETLossDefinitionClassOnce sync.Once
)

func getETLossDefinitionClass() ETLossDefinitionClass {
	_ETLossDefinitionClassOnce.Do(func() {
		_ETLossDefinitionClass = ETLossDefinitionClass{class: objc.GetClass("ETLossDefinition")}
	})
	return _ETLossDefinitionClass
}

// GetETLossDefinitionClass returns the class object for ETLossDefinition.
func GetETLossDefinitionClass() ETLossDefinitionClass {
	return getETLossDefinitionClass()
}

type ETLossDefinitionClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ec ETLossDefinitionClass) Class() objc.Class {
	return ec.class
}

// Alloc allocates memory for a new instance of the class.
func (ec ETLossDefinitionClass) Alloc() ETLossDefinition {
	rv := objc.Send[ETLossDefinition](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [ETLossDefinition.InputName]
//   - [ETLossDefinition.LossOutputName]
//   - [ETLossDefinition.Mode]
//   - [ETLossDefinition.OutputName]
//   - [ETLossDefinition.SetOutputName]
//   - [ETLossDefinition.TargetInputName]
//   - [ETLossDefinition.InitWithModeInputNameTargetNameLossOutputName]
// See: https://developer.apple.com/documentation/Espresso/ETLossDefinition
type ETLossDefinition struct {
	objectivec.Object
}

// ETLossDefinitionFromID constructs a [ETLossDefinition] from an objc.ID.
func ETLossDefinitionFromID(id objc.ID) ETLossDefinition {
	return ETLossDefinition{objectivec.Object{ID: id}}
}
// Ensure ETLossDefinition implements IETLossDefinition.
var _ IETLossDefinition = ETLossDefinition{}

// An interface definition for the [ETLossDefinition] class.
//
// # Methods
//
//   - [IETLossDefinition.InputName]
//   - [IETLossDefinition.LossOutputName]
//   - [IETLossDefinition.Mode]
//   - [IETLossDefinition.OutputName]
//   - [IETLossDefinition.SetOutputName]
//   - [IETLossDefinition.TargetInputName]
//   - [IETLossDefinition.InitWithModeInputNameTargetNameLossOutputName]
//
// See: https://developer.apple.com/documentation/Espresso/ETLossDefinition
type IETLossDefinition interface {
	objectivec.IObject

	// Topic: Methods

	InputName() string
	LossOutputName() string
	Mode() uint64
	OutputName() string
	SetOutputName(value string)
	TargetInputName() string
	InitWithModeInputNameTargetNameLossOutputName(mode uint64, name objectivec.IObject, name2 objectivec.IObject, name3 objectivec.IObject) ETLossDefinition
}

// Init initializes the instance.
func (e ETLossDefinition) Init() ETLossDefinition {
	rv := objc.Send[ETLossDefinition](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e ETLossDefinition) Autorelease() ETLossDefinition {
	rv := objc.Send[ETLossDefinition](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewETLossDefinition creates a new ETLossDefinition instance.
func NewETLossDefinition() ETLossDefinition {
	class := getETLossDefinitionClass()
	rv := objc.Send[ETLossDefinition](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/Espresso/ETLossDefinition/initWithMode:inputName:targetName:lossOutputName:
func NewETLossDefinitionWithModeInputNameTargetNameLossOutputName(mode uint64, name objectivec.IObject, name2 objectivec.IObject, name3 objectivec.IObject) ETLossDefinition {
	instance := getETLossDefinitionClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithMode:inputName:targetName:lossOutputName:"), mode, name, name2, name3)
	return ETLossDefinitionFromID(rv)
}

//
// See: https://developer.apple.com/documentation/Espresso/ETLossDefinition/initWithMode:inputName:targetName:lossOutputName:
func (e ETLossDefinition) InitWithModeInputNameTargetNameLossOutputName(mode uint64, name objectivec.IObject, name2 objectivec.IObject, name3 objectivec.IObject) ETLossDefinition {
	rv := objc.Send[ETLossDefinition](e.ID, objc.Sel("initWithMode:inputName:targetName:lossOutputName:"), mode, name, name2, name3)
	return rv
}

//
// See: https://developer.apple.com/documentation/Espresso/ETLossDefinition/BuiltInLoss:
func (_ETLossDefinitionClass ETLossDefinitionClass) BuiltInLoss(loss objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ETLossDefinitionClass.class), objc.Sel("BuiltInLoss:"), loss)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/Espresso/ETLossDefinition/L2LossWithInputName:targetInputName:lossOutputName:
func (_ETLossDefinitionClass ETLossDefinitionClass) L2LossWithInputNameTargetInputNameLossOutputName(name objectivec.IObject, name2 objectivec.IObject, name3 objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ETLossDefinitionClass.class), objc.Sel("L2LossWithInputName:targetInputName:lossOutputName:"), name, name2, name3)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/Espresso/ETLossDefinition/crossEntropyLossWithInputName:targetInputName:lossOutputName:
func (_ETLossDefinitionClass ETLossDefinitionClass) CrossEntropyLossWithInputNameTargetInputNameLossOutputName(name objectivec.IObject, name2 objectivec.IObject, name3 objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ETLossDefinitionClass.class), objc.Sel("crossEntropyLossWithInputName:targetInputName:lossOutputName:"), name, name2, name3)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/Espresso/ETLossDefinition/inputName
func (e ETLossDefinition) InputName() string {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("inputName"))
	return foundation.NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/Espresso/ETLossDefinition/lossOutputName
func (e ETLossDefinition) LossOutputName() string {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("lossOutputName"))
	return foundation.NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/Espresso/ETLossDefinition/mode
func (e ETLossDefinition) Mode() uint64 {
	rv := objc.Send[uint64](e.ID, objc.Sel("mode"))
	return rv
}
// See: https://developer.apple.com/documentation/Espresso/ETLossDefinition/outputName
func (e ETLossDefinition) OutputName() string {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("outputName"))
	return foundation.NSStringFromID(rv).String()
}
func (e ETLossDefinition) SetOutputName(value string) {
	objc.Send[struct{}](e.ID, objc.Sel("setOutputName:"), objc.String(value))
}
// See: https://developer.apple.com/documentation/Espresso/ETLossDefinition/targetInputName
func (e ETLossDefinition) TargetInputName() string {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("targetInputName"))
	return foundation.NSStringFromID(rv).String()
}

