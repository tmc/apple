// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MTLFunctionConstant] class.
var (
	_MTLFunctionConstantClass     MTLFunctionConstantClass
	_MTLFunctionConstantClassOnce sync.Once
)

func getMTLFunctionConstantClass() MTLFunctionConstantClass {
	_MTLFunctionConstantClassOnce.Do(func() {
		_MTLFunctionConstantClass = MTLFunctionConstantClass{class: objc.GetClass("MTLFunctionConstant")}
	})
	return _MTLFunctionConstantClass
}

// GetMTLFunctionConstantClass returns the class object for MTLFunctionConstant.
func GetMTLFunctionConstantClass() MTLFunctionConstantClass {
	return getMTLFunctionConstantClass()
}

type MTLFunctionConstantClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MTLFunctionConstantClass) Alloc() MTLFunctionConstant {
	rv := objc.Send[MTLFunctionConstant](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A constant that specializes the behavior of a shader.
//
// # Overview
// 
// Don’t create an [MTLFunctionConstant] instance directly. Instead, the
// list of function constants for a function by querying the
// `functionConstants` property of an [MTLFunction] instance.
// 
// An [MTLFunctionConstant] instance should only be obtained from a
// nonspecialized function created with the [NewFunctionWithName] method. You
// only need an [MTLFunctionConstant] instance if you don’t have sufficient
// information to create an [MTLFunctionConstantValues] instance used to
// create a specialized function with the
// [NewFunctionWithNameConstantValuesError] or
// [NewFunctionWithNameConstantValuesCompletionHandler] method.
//
// # Reading the function constant’s properties
//
//   - [MTLFunctionConstant.Name]: The name of the function constant.
//   - [MTLFunctionConstant.Type]: The data type of the function constant.
//   - [MTLFunctionConstant.Index]: The index of the function constant.
//   - [MTLFunctionConstant.Required]: A Boolean value indicating whether the function constant needs to be provided to specialize the function.
//
// See: https://developer.apple.com/documentation/Metal/MTLFunctionConstant
type MTLFunctionConstant struct {
	objectivec.Object
}

// MTLFunctionConstantFromID constructs a [MTLFunctionConstant] from an objc.ID.
//
// A constant that specializes the behavior of a shader.
func MTLFunctionConstantFromID(id objc.ID) MTLFunctionConstant {
	return MTLFunctionConstant{objectivec.Object{ID: id}}
}
// NOTE: MTLFunctionConstant adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MTLFunctionConstant] class.
//
// # Reading the function constant’s properties
//
//   - [IMTLFunctionConstant.Name]: The name of the function constant.
//   - [IMTLFunctionConstant.Type]: The data type of the function constant.
//   - [IMTLFunctionConstant.Index]: The index of the function constant.
//   - [IMTLFunctionConstant.Required]: A Boolean value indicating whether the function constant needs to be provided to specialize the function.
//
// See: https://developer.apple.com/documentation/Metal/MTLFunctionConstant
type IMTLFunctionConstant interface {
	objectivec.IObject

	// Topic: Reading the function constant’s properties

	// The name of the function constant.
	Name() string
	// The data type of the function constant.
	Type() MTLDataType
	// The index of the function constant.
	Index() uint
	// A Boolean value indicating whether the function constant needs to be provided to specialize the function.
	Required() bool
}

// Init initializes the instance.
func (f MTLFunctionConstant) Init() MTLFunctionConstant {
	rv := objc.Send[MTLFunctionConstant](f.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f MTLFunctionConstant) Autorelease() MTLFunctionConstant {
	rv := objc.Send[MTLFunctionConstant](f.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTLFunctionConstant creates a new MTLFunctionConstant instance.
func NewMTLFunctionConstant() MTLFunctionConstant {
	class := getMTLFunctionConstantClass()
	rv := objc.Send[MTLFunctionConstant](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The name of the function constant.
//
// See: https://developer.apple.com/documentation/Metal/MTLFunctionConstant/name
func (f MTLFunctionConstant) Name() string {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("name"))
	return foundation.NSStringFromID(rv).String()
}
// The data type of the function constant.
//
// See: https://developer.apple.com/documentation/Metal/MTLFunctionConstant/type
func (f MTLFunctionConstant) Type() MTLDataType {
	rv := objc.Send[MTLDataType](f.ID, objc.Sel("type"))
	return MTLDataType(rv)
}
// The index of the function constant.
//
// See: https://developer.apple.com/documentation/Metal/MTLFunctionConstant/index
func (f MTLFunctionConstant) Index() uint {
	rv := objc.Send[uint](f.ID, objc.Sel("index"))
	return rv
}
// A Boolean value indicating whether the function constant needs to be
// provided to specialize the function.
//
// # Discussion
// 
// This value is [true] if a constant value needs to be provided for the
// function constant. A function constant is optional only if it is referenced
// in a call to the built-in `is_function_constant_defined(name)` function.
// 
// Refer to the [Metal Shading Language Guide] for more information.
//
// [Metal Shading Language Guide]: https://developer.apple.com/library/archive/documentation/Metal/Reference/MetalShadingLanguageGuide/Introduction/Introduction.html#//apple_ref/doc/uid/TP40014364
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Metal/MTLFunctionConstant/required
func (f MTLFunctionConstant) Required() bool {
	rv := objc.Send[bool](f.ID, objc.Sel("required"))
	return rv
}

