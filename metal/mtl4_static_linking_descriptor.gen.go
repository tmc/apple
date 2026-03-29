// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MTL4StaticLinkingDescriptor] class.
var (
	_MTL4StaticLinkingDescriptorClass     MTL4StaticLinkingDescriptorClass
	_MTL4StaticLinkingDescriptorClassOnce sync.Once
)

func getMTL4StaticLinkingDescriptorClass() MTL4StaticLinkingDescriptorClass {
	_MTL4StaticLinkingDescriptorClassOnce.Do(func() {
		_MTL4StaticLinkingDescriptorClass = MTL4StaticLinkingDescriptorClass{class: objc.GetClass("MTL4StaticLinkingDescriptor")}
	})
	return _MTL4StaticLinkingDescriptorClass
}

// GetMTL4StaticLinkingDescriptorClass returns the class object for MTL4StaticLinkingDescriptor.
func GetMTL4StaticLinkingDescriptorClass() MTL4StaticLinkingDescriptorClass {
	return getMTL4StaticLinkingDescriptorClass()
}

type MTL4StaticLinkingDescriptorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MTL4StaticLinkingDescriptorClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MTL4StaticLinkingDescriptorClass) Alloc() MTL4StaticLinkingDescriptor {
	rv := objc.Send[MTL4StaticLinkingDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// Groups together properties to drive a static linking process.
//
// # Instance Properties
//
//   - [MTL4StaticLinkingDescriptor.FunctionDescriptors]: Provides an array of functions to link at the Metal IR level.
//   - [MTL4StaticLinkingDescriptor.SetFunctionDescriptors]
//   - [MTL4StaticLinkingDescriptor.Groups]: Assigns groups of functions to match call-site attributes in shader code.
//   - [MTL4StaticLinkingDescriptor.SetGroups]
//   - [MTL4StaticLinkingDescriptor.PrivateFunctionDescriptors]: Provides an array of private functions to link at the Metal IR level.
//   - [MTL4StaticLinkingDescriptor.SetPrivateFunctionDescriptors]
//
// See: https://developer.apple.com/documentation/Metal/MTL4StaticLinkingDescriptor
type MTL4StaticLinkingDescriptor struct {
	objectivec.Object
}

// MTL4StaticLinkingDescriptorFromID constructs a [MTL4StaticLinkingDescriptor] from an objc.ID.
//
// Groups together properties to drive a static linking process.
func MTL4StaticLinkingDescriptorFromID(id objc.ID) MTL4StaticLinkingDescriptor {
	return MTL4StaticLinkingDescriptor{objectivec.Object{ID: id}}
}
// NOTE: MTL4StaticLinkingDescriptor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MTL4StaticLinkingDescriptor] class.
//
// # Instance Properties
//
//   - [IMTL4StaticLinkingDescriptor.FunctionDescriptors]: Provides an array of functions to link at the Metal IR level.
//   - [IMTL4StaticLinkingDescriptor.SetFunctionDescriptors]
//   - [IMTL4StaticLinkingDescriptor.Groups]: Assigns groups of functions to match call-site attributes in shader code.
//   - [IMTL4StaticLinkingDescriptor.SetGroups]
//   - [IMTL4StaticLinkingDescriptor.PrivateFunctionDescriptors]: Provides an array of private functions to link at the Metal IR level.
//   - [IMTL4StaticLinkingDescriptor.SetPrivateFunctionDescriptors]
//
// See: https://developer.apple.com/documentation/Metal/MTL4StaticLinkingDescriptor
type IMTL4StaticLinkingDescriptor interface {
	objectivec.IObject

	// Topic: Instance Properties

	// Provides an array of functions to link at the Metal IR level.
	FunctionDescriptors() []MTL4FunctionDescriptor
	SetFunctionDescriptors(value []MTL4FunctionDescriptor)
	// Assigns groups of functions to match call-site attributes in shader code.
	Groups() foundation.INSDictionary
	SetGroups(value foundation.INSDictionary)
	// Provides an array of private functions to link at the Metal IR level.
	PrivateFunctionDescriptors() []MTL4FunctionDescriptor
	SetPrivateFunctionDescriptors(value []MTL4FunctionDescriptor)
}

// Init initializes the instance.
func (m MTL4StaticLinkingDescriptor) Init() MTL4StaticLinkingDescriptor {
	rv := objc.Send[MTL4StaticLinkingDescriptor](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MTL4StaticLinkingDescriptor) Autorelease() MTL4StaticLinkingDescriptor {
	rv := objc.Send[MTL4StaticLinkingDescriptor](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTL4StaticLinkingDescriptor creates a new MTL4StaticLinkingDescriptor instance.
func NewMTL4StaticLinkingDescriptor() MTL4StaticLinkingDescriptor {
	class := getMTL4StaticLinkingDescriptorClass()
	rv := objc.Send[MTL4StaticLinkingDescriptor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Provides an array of functions to link at the Metal IR level.
//
// See: https://developer.apple.com/documentation/Metal/MTL4StaticLinkingDescriptor/functionDescriptors
func (m MTL4StaticLinkingDescriptor) FunctionDescriptors() []MTL4FunctionDescriptor {
	rv := objc.Send[[]objc.ID](m.ID, objc.Sel("functionDescriptors"))
	return objc.ConvertSlice(rv, func(id objc.ID) MTL4FunctionDescriptor {
		return MTL4FunctionDescriptorFromID(id)
	})
}
func (m MTL4StaticLinkingDescriptor) SetFunctionDescriptors(value []MTL4FunctionDescriptor) {
	objc.Send[struct{}](m.ID, objc.Sel("setFunctionDescriptors:"), objectivec.IObjectSliceToNSArray(value))
}
// Assigns groups of functions to match call-site attributes in shader code.
//
// # Discussion
// 
// Function groups help the compiler reduce the number of candidate functions
// it needs to evaluate for shader function calls, potentially increasing
// runtime performance.
//
// See: https://developer.apple.com/documentation/Metal/MTL4StaticLinkingDescriptor/groups
func (m MTL4StaticLinkingDescriptor) Groups() foundation.INSDictionary {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("groups"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (m MTL4StaticLinkingDescriptor) SetGroups(value foundation.INSDictionary) {
	objc.Send[struct{}](m.ID, objc.Sel("setGroups:"), value)
}
// Provides an array of private functions to link at the Metal IR level.
//
// # Discussion
// 
// You specify private functions to link separately from [FunctionDescriptors]
// because pipelines don’t export private functions as [MTLFunctionHandle]
// instances.
//
// See: https://developer.apple.com/documentation/Metal/MTL4StaticLinkingDescriptor/privateFunctionDescriptors
func (m MTL4StaticLinkingDescriptor) PrivateFunctionDescriptors() []MTL4FunctionDescriptor {
	rv := objc.Send[[]objc.ID](m.ID, objc.Sel("privateFunctionDescriptors"))
	return objc.ConvertSlice(rv, func(id objc.ID) MTL4FunctionDescriptor {
		return MTL4FunctionDescriptorFromID(id)
	})
}
func (m MTL4StaticLinkingDescriptor) SetPrivateFunctionDescriptors(value []MTL4FunctionDescriptor) {
	objc.Send[struct{}](m.ID, objc.Sel("setPrivateFunctionDescriptors:"), objectivec.IObjectSliceToNSArray(value))
}

