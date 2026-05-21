// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"errors"
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLParameterContainer] class.
var (
	_MLParameterContainerClass     MLParameterContainerClass
	_MLParameterContainerClassOnce sync.Once
)

func getMLParameterContainerClass() MLParameterContainerClass {
	_MLParameterContainerClassOnce.Do(func() {
		_MLParameterContainerClass = MLParameterContainerClass{class: objc.GetClass("MLParameterContainer")}
	})
	return _MLParameterContainerClass
}

// GetMLParameterContainerClass returns the class object for MLParameterContainer.
func GetMLParameterContainerClass() MLParameterContainerClass {
	return getMLParameterContainerClass()
}

type MLParameterContainerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLParameterContainerClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLParameterContainerClass) Alloc() MLParameterContainer {
	rv := objc.Send[MLParameterContainer](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLParameterContainer.CurrentParameterValues]
//   - [MLParameterContainer.SetCurrentParameterValues]
//   - [MLParameterContainer.EncodeWithCoder]
//   - [MLParameterContainer.ParameterDescriptions]
//   - [MLParameterContainer.SetParameterDescriptions]
//   - [MLParameterContainer.ParameterKeys]
//   - [MLParameterContainer.SetParameterKeys]
//   - [MLParameterContainer.SetCurrentValueForKeyError]
//   - [MLParameterContainer.ValidateParameterValueGivenConstraint]
//   - [MLParameterContainer.InitWithCoder]
type MLParameterContainer struct {
	objectivec.Object
}

// MLParameterContainerFromID constructs a [MLParameterContainer] from an objc.ID.
func MLParameterContainerFromID(id objc.ID) MLParameterContainer {
	return MLParameterContainer{objectivec.Object{ID: id}}
}

// Ensure MLParameterContainer implements IMLParameterContainer.
var _ IMLParameterContainer = MLParameterContainer{}

// An interface definition for the [MLParameterContainer] class.
//
// # Methods
//
//   - [IMLParameterContainer.CurrentParameterValues]
//   - [IMLParameterContainer.SetCurrentParameterValues]
//   - [IMLParameterContainer.EncodeWithCoder]
//   - [IMLParameterContainer.ParameterDescriptions]
//   - [IMLParameterContainer.SetParameterDescriptions]
//   - [IMLParameterContainer.ParameterKeys]
//   - [IMLParameterContainer.SetParameterKeys]
//   - [IMLParameterContainer.SetCurrentValueForKeyError]
//   - [IMLParameterContainer.ValidateParameterValueGivenConstraint]
//   - [IMLParameterContainer.InitWithCoder]
type IMLParameterContainer interface {
	objectivec.IObject

	// Topic: Methods

	CurrentParameterValues() foundation.INSDictionary
	SetCurrentParameterValues(value foundation.INSDictionary)
	EncodeWithCoder(coder foundation.INSCoder)
	ParameterDescriptions() foundation.INSDictionary
	SetParameterDescriptions(value foundation.INSDictionary)
	ParameterKeys() foundation.INSArray
	SetParameterKeys(value foundation.INSArray)
	SetCurrentValueForKeyError(value objectivec.IObject, key objectivec.IObject) (bool, error)
	ValidateParameterValueGivenConstraint(value objectivec.IObject, constraint objectivec.IObject) bool
	InitWithCoder(coder foundation.INSCoder) MLParameterContainer
}

// Init initializes the instance.
func (m MLParameterContainer) Init() MLParameterContainer {
	rv := objc.Send[MLParameterContainer](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLParameterContainer) Autorelease() MLParameterContainer {
	rv := objc.Send[MLParameterContainer](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLParameterContainer creates a new MLParameterContainer instance.
func NewMLParameterContainer() MLParameterContainer {
	class := getMLParameterContainerClass()
	rv := objc.Send[MLParameterContainer](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewParameterContainerWithCoder(coder objectivec.IObject) MLParameterContainer {
	instance := getMLParameterContainerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return MLParameterContainerFromID(rv)
}

func (m MLParameterContainer) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](m.ID, objc.Sel("encodeWithCoder:"), coder)
}
func (m MLParameterContainer) SetCurrentValueForKeyError(value objectivec.IObject, key objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](m.ID, objc.Sel("setCurrentValue:forKey:error:"), value, key, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("setCurrentValue:forKey:error: returned NO with nil NSError")
	}
	return rv, nil

}
func (m MLParameterContainer) ValidateParameterValueGivenConstraint(value objectivec.IObject, constraint objectivec.IObject) bool {
	rv := objc.Send[bool](m.ID, objc.Sel("validateParameterValue:givenConstraint:"), value, constraint)
	return rv
}
func (m MLParameterContainer) InitWithCoder(coder foundation.INSCoder) MLParameterContainer {
	rv := objc.Send[MLParameterContainer](m.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}

func (_MLParameterContainerClass MLParameterContainerClass) ParameterContainerForDescriptions(for_ objectivec.IObject, descriptions objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLParameterContainerClass.class), objc.Sel("parameterContainerFor:descriptions:"), for_, descriptions)
	return objectivec.Object{ID: rv}
}
func (_MLParameterContainerClass MLParameterContainerClass) SupportsSecureCoding() bool {
	rv := objc.Send[bool](objc.ID(_MLParameterContainerClass.class), objc.Sel("supportsSecureCoding"))
	return rv
}

func (m MLParameterContainer) CurrentParameterValues() foundation.INSDictionary {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("currentParameterValues"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (m MLParameterContainer) SetCurrentParameterValues(value foundation.INSDictionary) {
	objc.Send[struct{}](m.ID, objc.Sel("setCurrentParameterValues:"), value)
}
func (m MLParameterContainer) ParameterDescriptions() foundation.INSDictionary {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("parameterDescriptions"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (m MLParameterContainer) SetParameterDescriptions(value foundation.INSDictionary) {
	objc.Send[struct{}](m.ID, objc.Sel("setParameterDescriptions:"), value)
}
func (m MLParameterContainer) ParameterKeys() foundation.INSArray {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("parameterKeys"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (m MLParameterContainer) SetParameterKeys(value foundation.INSArray) {
	objc.Send[struct{}](m.ID, objc.Sel("setParameterKeys:"), value)
}
