// Code generated from Apple documentation for ObjectiveC. DO NOT EDIT.

package objectivec

import (
	"unsafe"

	"github.com/tmc/apple/objc"
)

// Object wraps an Objective-C object ID.
// This is the base type embedded by all generated class types.
type Object struct {
	ID objc.ID
}

// IObject is the interface implemented by all Objective-C object wrappers.
type IObject interface {
	GetID() objc.ID
}

// GetID returns the underlying Objective-C object ID.
func (o Object) GetID() objc.ID {
	return o.ID
}

// RespondsToSelector returns true if the object responds to the given selector.
func (o Object) RespondsToSelector(sel objc.SEL) bool {
	return objc.Send[bool](o.ID, objc.Sel("respondsToSelector:"), sel)
}

// Description returns the object's description as a Go string.
func (o Object) Description() string {
	descID := objc.Send[objc.ID](o.ID, objc.Sel("description"))
	if descID == 0 {
		return ""
	}
	cstr := objc.Send[*byte](descID, objc.Sel("UTF8String"))
	return objc.GoString(cstr)
}

// DebugDescription returns the object's debug description as a Go string.
func (o Object) DebugDescription() string {
	descID := objc.Send[objc.ID](o.ID, objc.Sel("debugDescription"))
	if descID == 0 {
		return ""
	}
	cstr := objc.Send[*byte](descID, objc.Sel("UTF8String"))
	return objc.GoString(cstr)
}

// String implements fmt.Stringer by calling Description.
func (o Object) String() string {
	return o.Description()
}

// ObjectFromID creates an Object from an objc.ID.
func ObjectFromID(id objc.ID) Object {
	return Object{ID: id}
}

// C struct types

// NXHashState
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NXHashState
type NXHashState struct {
	I int
	J int
}

// NXHashTable
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NXHashTable
type NXHashTable struct {
	Buckets   unsafe.Pointer
	Count     uint
	Info      unsafe.Pointer
	NbBuckets uint
	Prototype *NXHashTablePrototype
}

// NXHashTablePrototype
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NXHashTablePrototype
type NXHashTablePrototype struct {
	Free    func(unsafe.Pointer, unsafe.Pointer)
	Hash    func(unsafe.Pointer, unsafe.Pointer) uint
	IsEqual func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) int
	Style   int
}

// Objc_method_description - Defines an Objective-C method.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/objc_method_description
type Objc_method_description struct {
	Name  SEL   // The name of the method at runtime.
	Types *byte // The types of the method arguments.

}

// Objc_object - Represents an instance of a class.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/objc_object
type Objc_object struct {
	Isa objc.Class // A pointer to the class definition of which this object is an instance.

}

// Objc_property_attribute_t - Defines a property attribute.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/objc_property_attribute_t
type Objc_property_attribute_t struct {
	Name  *byte // The name of the attribute.
	Value *byte // The value of the attribute (usually empty).

}

// Objc_super - Specifies the superclass of an instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/objc_super-swift.struct
type Objc_super struct {
	Receiver    unsafe.Pointer // A pointer of type [objc_object](<doc://com.apple.objectivec/documentation/ObjectiveC/objc_object>). Specifies an instance of a class.
	Super_class objc.Class     // A pointer to a [Class](<doc://com.apple.objectivec/documentation/ObjectiveC/Class>) data structure. Specifies the particular superclass of the instance to message.

}

// Manual definitions for opaque types not fully parsed.

type Objc_class struct{}
type Objc_method struct{}
type Objc_ivar struct{}
type Objc_selector struct{}
type Objc_property struct{}
type Objc_category struct{}
