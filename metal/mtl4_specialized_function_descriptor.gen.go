// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
)

// The class instance for the [MTL4SpecializedFunctionDescriptor] class.
var (
	_MTL4SpecializedFunctionDescriptorClass     MTL4SpecializedFunctionDescriptorClass
	_MTL4SpecializedFunctionDescriptorClassOnce sync.Once
)

func getMTL4SpecializedFunctionDescriptorClass() MTL4SpecializedFunctionDescriptorClass {
	_MTL4SpecializedFunctionDescriptorClassOnce.Do(func() {
		_MTL4SpecializedFunctionDescriptorClass = MTL4SpecializedFunctionDescriptorClass{class: objc.GetClass("MTL4SpecializedFunctionDescriptor")}
	})
	return _MTL4SpecializedFunctionDescriptorClass
}

// GetMTL4SpecializedFunctionDescriptorClass returns the class object for MTL4SpecializedFunctionDescriptor.
func GetMTL4SpecializedFunctionDescriptorClass() MTL4SpecializedFunctionDescriptorClass {
	return getMTL4SpecializedFunctionDescriptorClass()
}

type MTL4SpecializedFunctionDescriptorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MTL4SpecializedFunctionDescriptorClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MTL4SpecializedFunctionDescriptorClass) Alloc() MTL4SpecializedFunctionDescriptor {
	rv := objc.Send[MTL4SpecializedFunctionDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// Groups together properties to configure and create a specialized function
// by passing it to a factory method.
//
// # Overview
// 
// You can pass an instance of this class to any methods that accept a
// [MTL4FunctionDescriptor] parameter to provide extra configuration, such as
// function constants or a name.
//
// # Instance Properties
//
//   - [MTL4SpecializedFunctionDescriptor.ConstantValues]: Configures optional function constant values to associate with the function.
//   - [MTL4SpecializedFunctionDescriptor.SetConstantValues]
//   - [MTL4SpecializedFunctionDescriptor.FunctionDescriptor]: Provides a descriptor that corresponds to a base function that the specialization applies to.
//   - [MTL4SpecializedFunctionDescriptor.SetFunctionDescriptor]
//   - [MTL4SpecializedFunctionDescriptor.SpecializedName]: Assigns an optional name to the specialized function.
//   - [MTL4SpecializedFunctionDescriptor.SetSpecializedName]
//
// See: https://developer.apple.com/documentation/Metal/MTL4SpecializedFunctionDescriptor
type MTL4SpecializedFunctionDescriptor struct {
	MTL4FunctionDescriptor
}

// MTL4SpecializedFunctionDescriptorFromID constructs a [MTL4SpecializedFunctionDescriptor] from an objc.ID.
//
// Groups together properties to configure and create a specialized function
// by passing it to a factory method.
func MTL4SpecializedFunctionDescriptorFromID(id objc.ID) MTL4SpecializedFunctionDescriptor {
	return MTL4SpecializedFunctionDescriptor{MTL4FunctionDescriptor: MTL4FunctionDescriptorFromID(id)}
}
// NOTE: MTL4SpecializedFunctionDescriptor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MTL4SpecializedFunctionDescriptor] class.
//
// # Instance Properties
//
//   - [IMTL4SpecializedFunctionDescriptor.ConstantValues]: Configures optional function constant values to associate with the function.
//   - [IMTL4SpecializedFunctionDescriptor.SetConstantValues]
//   - [IMTL4SpecializedFunctionDescriptor.FunctionDescriptor]: Provides a descriptor that corresponds to a base function that the specialization applies to.
//   - [IMTL4SpecializedFunctionDescriptor.SetFunctionDescriptor]
//   - [IMTL4SpecializedFunctionDescriptor.SpecializedName]: Assigns an optional name to the specialized function.
//   - [IMTL4SpecializedFunctionDescriptor.SetSpecializedName]
//
// See: https://developer.apple.com/documentation/Metal/MTL4SpecializedFunctionDescriptor
type IMTL4SpecializedFunctionDescriptor interface {
	IMTL4FunctionDescriptor

	// Topic: Instance Properties

	// Configures optional function constant values to associate with the function.
	ConstantValues() IMTLFunctionConstantValues
	SetConstantValues(value IMTLFunctionConstantValues)
	// Provides a descriptor that corresponds to a base function that the specialization applies to.
	FunctionDescriptor() IMTL4FunctionDescriptor
	SetFunctionDescriptor(value IMTL4FunctionDescriptor)
	// Assigns an optional name to the specialized function.
	SpecializedName() string
	SetSpecializedName(value string)
}

// Init initializes the instance.
func (m MTL4SpecializedFunctionDescriptor) Init() MTL4SpecializedFunctionDescriptor {
	rv := objc.Send[MTL4SpecializedFunctionDescriptor](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MTL4SpecializedFunctionDescriptor) Autorelease() MTL4SpecializedFunctionDescriptor {
	rv := objc.Send[MTL4SpecializedFunctionDescriptor](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTL4SpecializedFunctionDescriptor creates a new MTL4SpecializedFunctionDescriptor instance.
func NewMTL4SpecializedFunctionDescriptor() MTL4SpecializedFunctionDescriptor {
	class := getMTL4SpecializedFunctionDescriptorClass()
	rv := objc.Send[MTL4SpecializedFunctionDescriptor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Configures optional function constant values to associate with the
// function.
//
// See: https://developer.apple.com/documentation/Metal/MTL4SpecializedFunctionDescriptor/constantValues
func (m MTL4SpecializedFunctionDescriptor) ConstantValues() IMTLFunctionConstantValues {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("constantValues"))
	return MTLFunctionConstantValuesFromID(objc.ID(rv))
}
func (m MTL4SpecializedFunctionDescriptor) SetConstantValues(value IMTLFunctionConstantValues) {
	objc.Send[struct{}](m.ID, objc.Sel("setConstantValues:"), value)
}
// Provides a descriptor that corresponds to a base function that the
// specialization applies to.
//
// See: https://developer.apple.com/documentation/Metal/MTL4SpecializedFunctionDescriptor/functionDescriptor
func (m MTL4SpecializedFunctionDescriptor) FunctionDescriptor() IMTL4FunctionDescriptor {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("functionDescriptor"))
	return MTL4FunctionDescriptorFromID(objc.ID(rv))
}
func (m MTL4SpecializedFunctionDescriptor) SetFunctionDescriptor(value IMTL4FunctionDescriptor) {
	objc.Send[struct{}](m.ID, objc.Sel("setFunctionDescriptor:"), value)
}
// Assigns an optional name to the specialized function.
//
// See: https://developer.apple.com/documentation/Metal/MTL4SpecializedFunctionDescriptor/specializedName
func (m MTL4SpecializedFunctionDescriptor) SpecializedName() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("specializedName"))
	return foundation.NSStringFromID(rv).String()
}
func (m MTL4SpecializedFunctionDescriptor) SetSpecializedName(value string) {
	objc.Send[struct{}](m.ID, objc.Sel("setSpecializedName:"), objc.String(value))
}

