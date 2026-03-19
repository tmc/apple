// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MTLLinkedFunctions] class.
var (
	_MTLLinkedFunctionsClass     MTLLinkedFunctionsClass
	_MTLLinkedFunctionsClassOnce sync.Once
)

func getMTLLinkedFunctionsClass() MTLLinkedFunctionsClass {
	_MTLLinkedFunctionsClassOnce.Do(func() {
		_MTLLinkedFunctionsClass = MTLLinkedFunctionsClass{class: objc.GetClass("MTLLinkedFunctions")}
	})
	return _MTLLinkedFunctionsClass
}

// GetMTLLinkedFunctionsClass returns the class object for MTLLinkedFunctions.
func GetMTLLinkedFunctionsClass() MTLLinkedFunctionsClass {
	return getMTLLinkedFunctionsClass()
}

type MTLLinkedFunctionsClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MTLLinkedFunctionsClass) Alloc() MTLLinkedFunctions {
	rv := objc.Send[MTLLinkedFunctions](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A set of related functions that Metal links to when necessary to create the
// function instance.
//
// # Overview
// 
// When you create a Metal function instance using an [MTLFunctionDescriptor],
// you specify additional functions that Metal needs to link to when it
// compiles and links the underlying shader code. Most often, you need to do
// this if your shader takes a visible function table as one or more of its
// arguments. For Metal to create the [MTLFunction] instance, it needs a
// complete list of functions that your shader can call so that it can resolve
// any dependencies and generate the correct code to run on the GPU.
//
// # Specifying related functions
//
//   - [MTLLinkedFunctions.Functions]: An array of function objects to link to the new function.
//   - [MTLLinkedFunctions.SetFunctions]
//   - [MTLLinkedFunctions.BinaryFunctions]: An array of function objects already compiled to a binary representation to link.
//   - [MTLLinkedFunctions.SetBinaryFunctions]
//   - [MTLLinkedFunctions.Groups]: An optional list of groups specifying which functions your shader can call at each call site.
//   - [MTLLinkedFunctions.SetGroups]
//   - [MTLLinkedFunctions.PrivateFunctions]: An array of function objects to link to the new function, without exporting the functions publicly.
//   - [MTLLinkedFunctions.SetPrivateFunctions]
//
// See: https://developer.apple.com/documentation/Metal/MTLLinkedFunctions
type MTLLinkedFunctions struct {
	objectivec.Object
}

// MTLLinkedFunctionsFromID constructs a [MTLLinkedFunctions] from an objc.ID.
//
// A set of related functions that Metal links to when necessary to create the
// function instance.
func MTLLinkedFunctionsFromID(id objc.ID) MTLLinkedFunctions {
	return MTLLinkedFunctions{objectivec.Object{ID: id}}
}
// NOTE: MTLLinkedFunctions adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MTLLinkedFunctions] class.
//
// # Specifying related functions
//
//   - [IMTLLinkedFunctions.Functions]: An array of function objects to link to the new function.
//   - [IMTLLinkedFunctions.SetFunctions]
//   - [IMTLLinkedFunctions.BinaryFunctions]: An array of function objects already compiled to a binary representation to link.
//   - [IMTLLinkedFunctions.SetBinaryFunctions]
//   - [IMTLLinkedFunctions.Groups]: An optional list of groups specifying which functions your shader can call at each call site.
//   - [IMTLLinkedFunctions.SetGroups]
//   - [IMTLLinkedFunctions.PrivateFunctions]: An array of function objects to link to the new function, without exporting the functions publicly.
//   - [IMTLLinkedFunctions.SetPrivateFunctions]
//
// See: https://developer.apple.com/documentation/Metal/MTLLinkedFunctions
type IMTLLinkedFunctions interface {
	objectivec.IObject

	// Topic: Specifying related functions

	// An array of function objects to link to the new function.
	Functions() []objectivec.IObject
	SetFunctions(value []objectivec.IObject)
	// An array of function objects already compiled to a binary representation to link.
	BinaryFunctions() []objectivec.IObject
	SetBinaryFunctions(value []objectivec.IObject)
	// An optional list of groups specifying which functions your shader can call at each call site.
	Groups() foundation.INSDictionary
	SetGroups(value foundation.INSDictionary)
	// An array of function objects to link to the new function, without exporting the functions publicly.
	PrivateFunctions() []objectivec.IObject
	SetPrivateFunctions(value []objectivec.IObject)

	// The binary archives to search for a previously-compiled version of this function.
	BinaryArchives() MTLBinaryArchive
	SetBinaryArchives(value MTLBinaryArchive)
	// The set of constant values assigned to the function constants.
	ConstantValues() IMTLFunctionConstantValues
	SetConstantValues(value IMTLFunctionConstantValues)
	// The name of the function to fetch from the library.
	Name() string
	SetName(value string)
	// Flags specifying how Metal should create the new function object.
	Options() MTLFunctionOptions
	SetOptions(value MTLFunctionOptions)
	// A new name for the created function object.
	SpecializedName() string
	SetSpecializedName(value string)
}

// Init initializes the instance.
func (l MTLLinkedFunctions) Init() MTLLinkedFunctions {
	rv := objc.Send[MTLLinkedFunctions](l.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (l MTLLinkedFunctions) Autorelease() MTLLinkedFunctions {
	rv := objc.Send[MTLLinkedFunctions](l.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTLLinkedFunctions creates a new MTLLinkedFunctions instance.
func NewMTLLinkedFunctions() MTLLinkedFunctions {
	class := getMTLLinkedFunctionsClass()
	rv := objc.Send[MTLLinkedFunctions](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates an empty linked functions object.
//
// See: https://developer.apple.com/documentation/Metal/MTLLinkedFunctions/linkedFunctions
func (_MTLLinkedFunctionsClass MTLLinkedFunctionsClass) LinkedFunctions() MTLLinkedFunctions {
	rv := objc.Send[objc.ID](objc.ID(_MTLLinkedFunctionsClass.class), objc.Sel("linkedFunctions"))
	return MTLLinkedFunctionsFromID(rv)
}

// An array of function objects to link to the new function.
//
// See: https://developer.apple.com/documentation/Metal/MTLLinkedFunctions/functions
func (l MTLLinkedFunctions) Functions() []objectivec.IObject {
	rv := objc.Send[[]objc.ID](l.ID, objc.Sel("functions"))
	return objc.ConvertSlice(rv, func(id objc.ID) objectivec.IObject {
		return objectivec.Object{ID: id}
	})
}
func (l MTLLinkedFunctions) SetFunctions(value []objectivec.IObject) {
	objc.Send[struct{}](l.ID, objc.Sel("setFunctions:"), objectivec.IObjectSliceToNSArray(value))
}
// An array of function objects already compiled to a binary representation to
// link.
//
// See: https://developer.apple.com/documentation/Metal/MTLLinkedFunctions/binaryFunctions
func (l MTLLinkedFunctions) BinaryFunctions() []objectivec.IObject {
	rv := objc.Send[[]objc.ID](l.ID, objc.Sel("binaryFunctions"))
	return objc.ConvertSlice(rv, func(id objc.ID) objectivec.IObject {
		return objectivec.Object{ID: id}
	})
}
func (l MTLLinkedFunctions) SetBinaryFunctions(value []objectivec.IObject) {
	objc.Send[struct{}](l.ID, objc.Sel("setBinaryFunctions:"), objectivec.IObjectSliceToNSArray(value))
}
// An optional list of groups specifying which functions your shader can call
// at each call site.
//
// # Discussion
// 
// The default value is `nil`.
// 
// The default behavior is conservative and assumes that your shader can call
// any linked function from every call site. If you know that the shader can
// only call a limited subset of functions at a call site, you can annotate
// those sites in the shader with a name of a group and then specify the list
// of functions for that call site using this property. Specifying call sites
// and callable functions more precisely can improve performance.
// 
// For more information on how to specify call site groups, see [Metal Shading
// Language Specification].
// 
// The value of this property is a dictionary whose keys are call site names
// and values are arrays specifying the list of functions that the shader can
// call from each site.
//
// [Metal Shading Language Specification]: https://developer.apple.com/metal/Metal-Shading-Language-Specification.pdf
//
// See: https://developer.apple.com/documentation/Metal/MTLLinkedFunctions/groups
func (l MTLLinkedFunctions) Groups() foundation.INSDictionary {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("groups"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (l MTLLinkedFunctions) SetGroups(value foundation.INSDictionary) {
	objc.Send[struct{}](l.ID, objc.Sel("setGroups:"), value)
}
// An array of function objects to link to the new function, without exporting
// the functions publicly.
//
// # Discussion
// 
// The pipeline doesn’t export these functions as [MTLFunctionHandle]
// instances because the Metal device doesn’t need to support function
// pointers to link private functions.
//
// See: https://developer.apple.com/documentation/Metal/MTLLinkedFunctions/privateFunctions
func (l MTLLinkedFunctions) PrivateFunctions() []objectivec.IObject {
	rv := objc.Send[[]objc.ID](l.ID, objc.Sel("privateFunctions"))
	return objc.ConvertSlice(rv, func(id objc.ID) objectivec.IObject {
		return objectivec.Object{ID: id}
	})
}
func (l MTLLinkedFunctions) SetPrivateFunctions(value []objectivec.IObject) {
	objc.Send[struct{}](l.ID, objc.Sel("setPrivateFunctions:"), objectivec.IObjectSliceToNSArray(value))
}
// The binary archives to search for a previously-compiled version of this
// function.
//
// See: https://developer.apple.com/documentation/metal/mtlfunctiondescriptor/binaryarchives
func (l MTLLinkedFunctions) BinaryArchives() MTLBinaryArchive {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("binaryArchives"))
	return MTLBinaryArchiveObjectFromID(rv)
}
func (l MTLLinkedFunctions) SetBinaryArchives(value MTLBinaryArchive) {
	objc.Send[struct{}](l.ID, objc.Sel("setBinaryArchives:"), value)
}
// The set of constant values assigned to the function constants.
//
// See: https://developer.apple.com/documentation/metal/mtlfunctiondescriptor/constantvalues
func (l MTLLinkedFunctions) ConstantValues() IMTLFunctionConstantValues {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("constantValues"))
	return MTLFunctionConstantValuesFromID(objc.ID(rv))
}
func (l MTLLinkedFunctions) SetConstantValues(value IMTLFunctionConstantValues) {
	objc.Send[struct{}](l.ID, objc.Sel("setConstantValues:"), value)
}
// The name of the function to fetch from the library.
//
// See: https://developer.apple.com/documentation/metal/mtlfunctiondescriptor/name
func (l MTLLinkedFunctions) Name() string {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("name"))
	return foundation.NSStringFromID(rv).String()
}
func (l MTLLinkedFunctions) SetName(value string) {
	objc.Send[struct{}](l.ID, objc.Sel("setName:"), objc.String(value))
}
// Flags specifying how Metal should create the new function object.
//
// See: https://developer.apple.com/documentation/metal/mtlfunctiondescriptor/options
func (l MTLLinkedFunctions) Options() MTLFunctionOptions {
	rv := objc.Send[MTLFunctionOptions](l.ID, objc.Sel("options"))
	return MTLFunctionOptions(rv)
}
func (l MTLLinkedFunctions) SetOptions(value MTLFunctionOptions) {
	objc.Send[struct{}](l.ID, objc.Sel("setOptions:"), value)
}
// A new name for the created function object.
//
// See: https://developer.apple.com/documentation/metal/mtlfunctiondescriptor/specializedname
func (l MTLLinkedFunctions) SpecializedName() string {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("specializedName"))
	return foundation.NSStringFromID(rv).String()
}
func (l MTLLinkedFunctions) SetSpecializedName(value string) {
	objc.Send[struct{}](l.ID, objc.Sel("setSpecializedName:"), objc.String(value))
}

