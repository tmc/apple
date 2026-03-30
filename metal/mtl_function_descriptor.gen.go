// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MTLFunctionDescriptor] class.
var (
	_MTLFunctionDescriptorClass     MTLFunctionDescriptorClass
	_MTLFunctionDescriptorClassOnce sync.Once
)

func getMTLFunctionDescriptorClass() MTLFunctionDescriptorClass {
	_MTLFunctionDescriptorClassOnce.Do(func() {
		_MTLFunctionDescriptorClass = MTLFunctionDescriptorClass{class: objc.GetClass("MTLFunctionDescriptor")}
	})
	return _MTLFunctionDescriptorClass
}

// GetMTLFunctionDescriptorClass returns the class object for MTLFunctionDescriptor.
func GetMTLFunctionDescriptorClass() MTLFunctionDescriptorClass {
	return getMTLFunctionDescriptorClass()
}

type MTLFunctionDescriptorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MTLFunctionDescriptorClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MTLFunctionDescriptorClass) Alloc() MTLFunctionDescriptor {
	rv := objc.Send[MTLFunctionDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A description of a function object to create.
//
// # Specifying the function configuration
//
//   - [MTLFunctionDescriptor.Name]: The name of the function to fetch from the library.
//   - [MTLFunctionDescriptor.SetName]
//   - [MTLFunctionDescriptor.SpecializedName]: A new name for the created function object.
//   - [MTLFunctionDescriptor.SetSpecializedName]
//   - [MTLFunctionDescriptor.ConstantValues]: The set of constant values assigned to the function constants.
//   - [MTLFunctionDescriptor.SetConstantValues]
//   - [MTLFunctionDescriptor.Options]: Flags specifying how Metal should create the new function object.
//   - [MTLFunctionDescriptor.SetOptions]
//   - [MTLFunctionDescriptor.BinaryArchives]: The binary archives to search for a previously-compiled version of this function.
//   - [MTLFunctionDescriptor.SetBinaryArchives]
//
// See: https://developer.apple.com/documentation/Metal/MTLFunctionDescriptor
type MTLFunctionDescriptor struct {
	objectivec.Object
}

// MTLFunctionDescriptorFromID constructs a [MTLFunctionDescriptor] from an objc.ID.
//
// A description of a function object to create.
func MTLFunctionDescriptorFromID(id objc.ID) MTLFunctionDescriptor {
	return MTLFunctionDescriptor{objectivec.Object{ID: id}}
}

// NOTE: MTLFunctionDescriptor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MTLFunctionDescriptor] class.
//
// # Specifying the function configuration
//
//   - [IMTLFunctionDescriptor.Name]: The name of the function to fetch from the library.
//   - [IMTLFunctionDescriptor.SetName]
//   - [IMTLFunctionDescriptor.SpecializedName]: A new name for the created function object.
//   - [IMTLFunctionDescriptor.SetSpecializedName]
//   - [IMTLFunctionDescriptor.ConstantValues]: The set of constant values assigned to the function constants.
//   - [IMTLFunctionDescriptor.SetConstantValues]
//   - [IMTLFunctionDescriptor.Options]: Flags specifying how Metal should create the new function object.
//   - [IMTLFunctionDescriptor.SetOptions]
//   - [IMTLFunctionDescriptor.BinaryArchives]: The binary archives to search for a previously-compiled version of this function.
//   - [IMTLFunctionDescriptor.SetBinaryArchives]
//
// See: https://developer.apple.com/documentation/Metal/MTLFunctionDescriptor
type IMTLFunctionDescriptor interface {
	objectivec.IObject

	// Topic: Specifying the function configuration

	// The name of the function to fetch from the library.
	Name() string
	SetName(value string)
	// A new name for the created function object.
	SpecializedName() string
	SetSpecializedName(value string)
	// The set of constant values assigned to the function constants.
	ConstantValues() IMTLFunctionConstantValues
	SetConstantValues(value IMTLFunctionConstantValues)
	// Flags specifying how Metal should create the new function object.
	Options() MTLFunctionOptions
	SetOptions(value MTLFunctionOptions)
	// The binary archives to search for a previously-compiled version of this function.
	BinaryArchives() []objectivec.IObject
	SetBinaryArchives(value []objectivec.IObject)
}

// Init initializes the instance.
func (f MTLFunctionDescriptor) Init() MTLFunctionDescriptor {
	rv := objc.Send[MTLFunctionDescriptor](f.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f MTLFunctionDescriptor) Autorelease() MTLFunctionDescriptor {
	rv := objc.Send[MTLFunctionDescriptor](f.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTLFunctionDescriptor creates a new MTLFunctionDescriptor instance.
func NewMTLFunctionDescriptor() MTLFunctionDescriptor {
	class := getMTLFunctionDescriptorClass()
	rv := objc.Send[MTLFunctionDescriptor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The name of the function to fetch from the library.
//
// See: https://developer.apple.com/documentation/Metal/MTLFunctionDescriptor/name
func (f MTLFunctionDescriptor) Name() string {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("name"))
	return foundation.NSStringFromID(rv).String()
}
func (f MTLFunctionDescriptor) SetName(value string) {
	objc.Send[struct{}](f.ID, objc.Sel("setName:"), objc.String(value))
}

// A new name for the created function object.
//
// # Discussion
//
// The default value is `nil`. If you specify a value for this property, Metal
// creates the new [MTLFunction] object with the new name. Use this property
// if you want to specialize a function with multiple variants and give each a
// distinct name.
//
// See: https://developer.apple.com/documentation/Metal/MTLFunctionDescriptor/specializedName
func (f MTLFunctionDescriptor) SpecializedName() string {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("specializedName"))
	return foundation.NSStringFromID(rv).String()
}
func (f MTLFunctionDescriptor) SetSpecializedName(value string) {
	objc.Send[struct{}](f.ID, objc.Sel("setSpecializedName:"), objc.String(value))
}

// The set of constant values assigned to the function constants.
//
// # Discussion
//
// The default value is `nil`. If you are creating a function object for a
// specialized function, you need to provide an array of valid constant values
// for all required function constants.
//
// See: https://developer.apple.com/documentation/Metal/MTLFunctionDescriptor/constantValues
func (f MTLFunctionDescriptor) ConstantValues() IMTLFunctionConstantValues {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("constantValues"))
	return MTLFunctionConstantValuesFromID(objc.ID(rv))
}
func (f MTLFunctionDescriptor) SetConstantValues(value IMTLFunctionConstantValues) {
	objc.Send[struct{}](f.ID, objc.Sel("setConstantValues:"), value)
}

// Flags specifying how Metal should create the new function object.
//
// See: https://developer.apple.com/documentation/Metal/MTLFunctionDescriptor/options
func (f MTLFunctionDescriptor) Options() MTLFunctionOptions {
	rv := objc.Send[MTLFunctionOptions](f.ID, objc.Sel("options"))
	return MTLFunctionOptions(rv)
}
func (f MTLFunctionDescriptor) SetOptions(value MTLFunctionOptions) {
	objc.Send[struct{}](f.ID, objc.Sel("setOptions:"), value)
}

// The binary archives to search for a previously-compiled version of this
// function.
//
// # Discussion
//
// If you specify an archive that includes a fully compiled version of this
// function, Metal uses the compiled version rather than creating a new one.
//
// See: https://developer.apple.com/documentation/Metal/MTLFunctionDescriptor/binaryArchives
func (f MTLFunctionDescriptor) BinaryArchives() []objectivec.IObject {
	rv := objc.Send[[]objc.ID](f.ID, objc.Sel("binaryArchives"))
	return objc.ConvertSlice(rv, func(id objc.ID) objectivec.IObject {
		return objectivec.Object{ID: id}
	})
}
func (f MTLFunctionDescriptor) SetBinaryArchives(value []objectivec.IObject) {
	objc.Send[struct{}](f.ID, objc.Sel("setBinaryArchives:"), objectivec.IObjectSliceToNSArray(value))
}
