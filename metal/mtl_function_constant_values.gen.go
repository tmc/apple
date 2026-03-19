// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MTLFunctionConstantValues] class.
var (
	_MTLFunctionConstantValuesClass     MTLFunctionConstantValuesClass
	_MTLFunctionConstantValuesClassOnce sync.Once
)

func getMTLFunctionConstantValuesClass() MTLFunctionConstantValuesClass {
	_MTLFunctionConstantValuesClassOnce.Do(func() {
		_MTLFunctionConstantValuesClass = MTLFunctionConstantValuesClass{class: objc.GetClass("MTLFunctionConstantValues")}
	})
	return _MTLFunctionConstantValuesClass
}

// GetMTLFunctionConstantValuesClass returns the class object for MTLFunctionConstantValues.
func GetMTLFunctionConstantValuesClass() MTLFunctionConstantValuesClass {
	return getMTLFunctionConstantValuesClass()
}

type MTLFunctionConstantValuesClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MTLFunctionConstantValuesClass) Alloc() MTLFunctionConstantValues {
	rv := objc.Send[MTLFunctionConstantValues](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A set of constant values that specialize a graphics or compute GPU
// function.
//
// # Overview
// 
// An [MTLFunctionConstantValues] instance sets constant values for function
// constants. You declare function constants with the `[[
// function_constant(index) ]]` attribute in MSL (Metal Shading Language)
// source code. See the [Metal Shading Language specification] for more
// information.
// 
// With an [MTLFunctionConstantValues] instance, you can set each constant
// value individually with an index or a name, or set multiple constant values
// with an index range.
// 
// You can apply a single [MTLFunctionConstantValues] instance to multiple
// [MTLFunction] instances of any kind, such as a vertex function and a
// fragment function. When you create a specialized function, subsequent
// changes to its constant values have no effect. However, you can reset, add,
// or modify a constant value in your [MTLFunctionConstantValues] instance and
// reuse it to create another [MTLFunction] instance.
//
// [Metal Shading Language specification]: https://developer.apple.com/metal/Metal-Shading-Language-Specification.pdf
//
// # Setting constant values
//
//   - [MTLFunctionConstantValues.SetConstantValueTypeAtIndex]: Sets a value for a function constant at a specific index.
//   - [MTLFunctionConstantValues.SetConstantValueTypeWithName]: Sets a value for a function constant with a specific name.
//
// # Resetting constant values
//
//   - [MTLFunctionConstantValues.Reset]: Deletes all previously set constant values.
//
// See: https://developer.apple.com/documentation/Metal/MTLFunctionConstantValues
type MTLFunctionConstantValues struct {
	objectivec.Object
}

// MTLFunctionConstantValuesFromID constructs a [MTLFunctionConstantValues] from an objc.ID.
//
// A set of constant values that specialize a graphics or compute GPU
// function.
func MTLFunctionConstantValuesFromID(id objc.ID) MTLFunctionConstantValues {
	return MTLFunctionConstantValues{objectivec.Object{ID: id}}
}
// NOTE: MTLFunctionConstantValues adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MTLFunctionConstantValues] class.
//
// # Setting constant values
//
//   - [IMTLFunctionConstantValues.SetConstantValueTypeAtIndex]: Sets a value for a function constant at a specific index.
//   - [IMTLFunctionConstantValues.SetConstantValueTypeWithName]: Sets a value for a function constant with a specific name.
//
// # Resetting constant values
//
//   - [IMTLFunctionConstantValues.Reset]: Deletes all previously set constant values.
//
// See: https://developer.apple.com/documentation/Metal/MTLFunctionConstantValues
type IMTLFunctionConstantValues interface {
	objectivec.IObject

	// Topic: Setting constant values

	// Sets a value for a function constant at a specific index.
	SetConstantValueTypeAtIndex(value unsafe.Pointer, type_ MTLDataType, index uint)
	// Sets a value for a function constant with a specific name.
	SetConstantValueTypeWithName(value unsafe.Pointer, type_ MTLDataType, name string)

	// Topic: Resetting constant values

	// Deletes all previously set constant values.
	Reset()

	// Sets values for a group of function constants within a specific index range.
	SetConstantValuesTypeWithRange(values unsafe.Pointer, type_ MTLDataType, range_ foundation.NSRange)
}

// Init initializes the instance.
func (f MTLFunctionConstantValues) Init() MTLFunctionConstantValues {
	rv := objc.Send[MTLFunctionConstantValues](f.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f MTLFunctionConstantValues) Autorelease() MTLFunctionConstantValues {
	rv := objc.Send[MTLFunctionConstantValues](f.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTLFunctionConstantValues creates a new MTLFunctionConstantValues instance.
func NewMTLFunctionConstantValues() MTLFunctionConstantValues {
	class := getMTLFunctionConstantValuesClass()
	rv := objc.Send[MTLFunctionConstantValues](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Sets a value for a function constant at a specific index.
//
// value: A pointer to the constant value.
//
// type: The data type of the function constant.
//
// index: The index of the function constant.
//
// # Discussion
// 
// Declare a single function constant in Metal Shading Language (MSL).
// 
// Set its value by assigning with a specific index.
//
// See: https://developer.apple.com/documentation/Metal/MTLFunctionConstantValues/setConstantValue(_:type:index:)
func (f MTLFunctionConstantValues) SetConstantValueTypeAtIndex(value unsafe.Pointer, type_ MTLDataType, index uint) {
	objc.Send[objc.ID](f.ID, objc.Sel("setConstantValue:type:atIndex:"), value, type_, index)
}
// Sets a value for a function constant with a specific name.
//
// value: A pointer to the constant value.
//
// type: The data type of the function constant.
//
// name: The name of the function constant.
//
// # Discussion
// 
// The first example declares a single function constant in a Metal Shading
// Language file.
// 
// The next example sets that Boolean value by providing its specific name.
//
// See: https://developer.apple.com/documentation/Metal/MTLFunctionConstantValues/setConstantValue(_:type:withName:)
func (f MTLFunctionConstantValues) SetConstantValueTypeWithName(value unsafe.Pointer, type_ MTLDataType, name string) {
	objc.Send[objc.ID](f.ID, objc.Sel("setConstantValue:type:withName:"), value, type_, objc.String(name))
}
// Deletes all previously set constant values.
//
// # Discussion
// 
// You don’t need to call this method if you don’t set any constant
// values.
//
// See: https://developer.apple.com/documentation/Metal/MTLFunctionConstantValues/reset()
func (f MTLFunctionConstantValues) Reset() {
	objc.Send[objc.ID](f.ID, objc.Sel("reset"))
}
// Sets values for a group of function constants within a specific index
// range.
//
// values: A pointer to the constant values.
//
// type: The data type of the function constants.
//
// range: The range of the function constant indices.
//
// # Discussion
// 
// Declare multiple function constants in Metal Shading Language (MSL).
// 
// Set their values by assigning an index range of an array.
//
// See: https://developer.apple.com/documentation/Metal/MTLFunctionConstantValues/setConstantValues:type:withRange:
func (f MTLFunctionConstantValues) SetConstantValuesTypeWithRange(values unsafe.Pointer, type_ MTLDataType, range_ foundation.NSRange) {
	objc.Send[objc.ID](f.ID, objc.Sel("setConstantValues:type:withRange:"), values, type_, range_)
}

