// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MTLAttributeDescriptorArray] class.
var (
	_MTLAttributeDescriptorArrayClass     MTLAttributeDescriptorArrayClass
	_MTLAttributeDescriptorArrayClassOnce sync.Once
)

func getMTLAttributeDescriptorArrayClass() MTLAttributeDescriptorArrayClass {
	_MTLAttributeDescriptorArrayClassOnce.Do(func() {
		_MTLAttributeDescriptorArrayClass = MTLAttributeDescriptorArrayClass{class: objc.GetClass("MTLAttributeDescriptorArray")}
	})
	return _MTLAttributeDescriptorArrayClass
}

// GetMTLAttributeDescriptorArrayClass returns the class object for MTLAttributeDescriptorArray.
func GetMTLAttributeDescriptorArrayClass() MTLAttributeDescriptorArrayClass {
	return getMTLAttributeDescriptorArrayClass()
}

type MTLAttributeDescriptorArrayClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MTLAttributeDescriptorArrayClass) Alloc() MTLAttributeDescriptorArray {
	rv := objc.Send[MTLAttributeDescriptorArray](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}







// An array of attribute descriptor objects.
//
// # Overview
// 
// An [MTLAttributeDescriptorArray] defines the data format and index binding
// for the attribute argument table, using [MTLAttributeDescriptor] instances.
//
// # Accessing attribute state objects
//
//   - [MTLAttributeDescriptorArray.ObjectAtIndexedSubscript]: Returns the state of the specified attribute.
//
// See: https://developer.apple.com/documentation/Metal/MTLAttributeDescriptorArray
type MTLAttributeDescriptorArray struct {
	objectivec.Object
}

// MTLAttributeDescriptorArrayFromID constructs a [MTLAttributeDescriptorArray] from an objc.ID.
//
// An array of attribute descriptor objects.
func MTLAttributeDescriptorArrayFromID(id objc.ID) MTLAttributeDescriptorArray {
	return MTLAttributeDescriptorArray{objectivec.Object{id}}
}
// NOTE: MTLAttributeDescriptorArray adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [MTLAttributeDescriptorArray] class.
//
// # Accessing attribute state objects
//
//   - [IMTLAttributeDescriptorArray.ObjectAtIndexedSubscript]: Returns the state of the specified attribute.
//
// See: https://developer.apple.com/documentation/Metal/MTLAttributeDescriptorArray
type IMTLAttributeDescriptorArray interface {
	objectivec.IObject

	// Topic: Accessing attribute state objects

	// Returns the state of the specified attribute.
	ObjectAtIndexedSubscript(index uint) IMTLAttributeDescriptor

	// The organization of input and output data for the next kernel call.
	StageInputDescriptor() IMTLStageInputOutputDescriptor
	SetStageInputDescriptor(value IMTLStageInputOutputDescriptor)
	// Sets state for the specified attribute.
	SetObjectAtIndexedSubscript(attributeDesc IMTLAttributeDescriptor, index uint)
}





// Init initializes the instance.
func (a MTLAttributeDescriptorArray) Init() MTLAttributeDescriptorArray {
	rv := objc.Send[MTLAttributeDescriptorArray](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a MTLAttributeDescriptorArray) Autorelease() MTLAttributeDescriptorArray {
	rv := objc.Send[MTLAttributeDescriptorArray](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTLAttributeDescriptorArray creates a new MTLAttributeDescriptorArray instance.
func NewMTLAttributeDescriptorArray() MTLAttributeDescriptorArray {
	class := getMTLAttributeDescriptorArrayClass()
	rv := objc.Send[MTLAttributeDescriptorArray](objc.ID(class.class), objc.Sel("new"))
	return rv
}










// Returns the state of the specified attribute.
//
// index: A specified index in the argument table bindings.
//
// # Return Value
// 
// The attribute descriptor for data bound at this index.
//
// See: https://developer.apple.com/documentation/Metal/MTLAttributeDescriptorArray/subscript(_:)
func (a MTLAttributeDescriptorArray) ObjectAtIndexedSubscript(index uint) IMTLAttributeDescriptor {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("objectAtIndexedSubscript:"), index)
	return MTLAttributeDescriptorFromID(rv)
}

// Sets state for the specified attribute.
//
// attributeDesc: A descriptor that contains attribute state.
//
// index: A specified index in the array of vertex attribute states.
//
// See: https://developer.apple.com/documentation/Metal/MTLAttributeDescriptorArray/setObject:atIndexedSubscript:
func (a MTLAttributeDescriptorArray) SetObjectAtIndexedSubscript(attributeDesc IMTLAttributeDescriptor, index uint) {
	objc.Send[objc.ID](a.ID, objc.Sel("setObject:atIndexedSubscript:"), attributeDesc, index)
}












// The organization of input and output data for the next kernel call.
//
// See: https://developer.apple.com/documentation/metal/mtlcomputepipelinedescriptor/stageinputdescriptor
func (a MTLAttributeDescriptorArray) StageInputDescriptor() IMTLStageInputOutputDescriptor {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("stageInputDescriptor"))
	return MTLStageInputOutputDescriptorFromID(objc.ID(rv))
}
func (a MTLAttributeDescriptorArray) SetStageInputDescriptor(value IMTLStageInputOutputDescriptor) {
	objc.Send[struct{}](a.ID, objc.Sel("setStageInputDescriptor:"), value)
}























