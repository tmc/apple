// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [ValueTransformer] class.
var (
	_ValueTransformerClass     ValueTransformerClass
	_ValueTransformerClassOnce sync.Once
)

func getValueTransformerClass() ValueTransformerClass {
	_ValueTransformerClassOnce.Do(func() {
		_ValueTransformerClass = ValueTransformerClass{class: objc.GetClass("NSValueTransformer")}
	})
	return _ValueTransformerClass
}

// GetValueTransformerClass returns the class object for NSValueTransformer.
func GetValueTransformerClass() ValueTransformerClass {
	return getValueTransformerClass()
}

type ValueTransformerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc ValueTransformerClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc ValueTransformerClass) Alloc() ValueTransformer {
	rv := objc.Send[ValueTransformer](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// An abstract class used to transform values from one representation to
// another.
//
// # Overview
//
// You create a value transformer by subclassing [NSValueTransformer] and
// overriding the necessary methods to provide the required custom
// transformation. You then register the value transformer using the
// [SetValueTransformerForName] method, so that other parts of your app can
// access it by name with [ValueTransformerForName].
//
// Use the [TransformedValue] method to transform a value from one
// representation into another. If a value transformer designates that its
// transformation is reversible by returning true for
// [AllowsReverseTransformation], you can also use the
// [ReverseTransformedValue] to perform the transformation in reverse. For
// example, reversing the characters in a string is a reversible operation,
// whereas changing the characters in a string to be uppercase is a
// nonreversible operation.
//
// A value transformer can take inputs of one type and return a value of a
// different type. For example, a value transformer could take an [NSImage] or
// [UIImage] object and return an [NSData] object containing the PNG
// representation of that image.
//
// # Example Usage
//
// The following example defines a new value transformer that takes an object
// and returns a string based on the object’s class type. This transformer
// isn’t reversible because it doesn’t make sense to transform a class
// name into an object.
//
// # Transforming Values
//
//   - [ValueTransformer.TransformedValue]: Returns the result of transforming a given value.
//   - [ValueTransformer.ReverseTransformedValue]: Returns the result of the reverse transformation of a given value.
//
// See: https://developer.apple.com/documentation/Foundation/ValueTransformer
//
// [NSImage]: https://developer.apple.com/documentation/AppKit/NSImage
// [UIImage]: https://developer.apple.com/documentation/UIKit/UIImage
type ValueTransformer struct {
	objectivec.Object
}

// ValueTransformerFromID constructs a [ValueTransformer] from an objc.ID.
//
// An abstract class used to transform values from one representation to
// another.
func ValueTransformerFromID(id objc.ID) ValueTransformer {
	return NSValueTransformer{objectivec.Object{ID: id}}
}

// NSValueTransformerFromID is an alias for [ValueTransformerFromID] for cross-framework compatibility.
func NSValueTransformerFromID(id objc.ID) ValueTransformer { return ValueTransformerFromID(id) }

// NOTE: ValueTransformer adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [ValueTransformer] class.
//
// # Transforming Values
//
//   - [IValueTransformer.TransformedValue]: Returns the result of transforming a given value.
//   - [IValueTransformer.ReverseTransformedValue]: Returns the result of the reverse transformation of a given value.
//
// See: https://developer.apple.com/documentation/Foundation/ValueTransformer
type IValueTransformer interface {
	objectivec.IObject

	// Topic: Transforming Values

	// Returns the result of transforming a given value.
	TransformedValue(value objectivec.IObject) objectivec.IObject
	// Returns the result of the reverse transformation of a given value.
	ReverseTransformedValue(value objectivec.IObject) objectivec.IObject
}

// Init initializes the instance.
func (v ValueTransformer) Init() ValueTransformer {
	rv := objc.Send[ValueTransformer](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v ValueTransformer) Autorelease() ValueTransformer {
	rv := objc.Send[ValueTransformer](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewValueTransformer creates a new ValueTransformer instance.
func NewValueTransformer() ValueTransformer {
	class := getValueTransformerClass()
	rv := objc.Send[ValueTransformer](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Returns the value transformer identified by a given identifier.
//
// name: The transformer identifier.
//
// # Return Value
//
// The value transformer identified by `name` in the shared registry, or `nil`
// if not found.
//
// # Discussion
//
// If “ does not find a registered transformer instance for `name`, it will
// attempt to find a class with the specified name. If a corresponding class
// is found an instance will be created and initialized using its “ method
// and then automatically registered with `name`.
//
// See: https://developer.apple.com/documentation/Foundation/ValueTransformer/init(forName:)
func NewValueTransformerForName(name NSValueTransformerName) ValueTransformer {
	rv := objc.Send[objc.ID](objc.ID(getValueTransformerClass().class), objc.Sel("valueTransformerForName:"), objc.String(string(name)))
	return ValueTransformerFromID(rv)
}

// Returns the result of transforming a given value.
//
// value: The value to transform.
//
// # Return Value
//
// The result of transforming `value`.
//
// The default implementation simply returns `value`.
//
// # Discussion
//
// A subclass should override this method to transform and return an object
// based on `value`.
//
// See: https://developer.apple.com/documentation/Foundation/ValueTransformer/transformedValue(_:)
func (v ValueTransformer) TransformedValue(value objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("transformedValue:"), value)
	return objectivec.Object{ID: rv}
}

// Returns the result of the reverse transformation of a given value.
//
// value: The value to reverse transform.
//
// # Return Value
//
// The reverse transformation of `value`.
//
// # Discussion
//
// The default implementation raises an exception if
// [AllowsReverseTransformation] returns false; otherwise it will invoke
// [TransformedValue] with `value`.
//
// A subclass should override this method if they require a reverse
// transformation that is not the same as simply reapplying the original
// transform (as would be the case with negation, for example). For example,
// if a value transformer converts a value in Fahrenheit to Celsius, this
// method would converts a value from Celsius to Fahrenheit.
//
// See: https://developer.apple.com/documentation/Foundation/ValueTransformer/reverseTransformedValue(_:)
func (v ValueTransformer) ReverseTransformedValue(value objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("reverseTransformedValue:"), value)
	return objectivec.Object{ID: rv}
}

// Registers the provided value transformer with a given identifier.
//
// transformer: The transformer to register.
//
// name: The name for `transformer`.
//
// See: https://developer.apple.com/documentation/Foundation/ValueTransformer/setValueTransformer(_:forName:)
func (_ValueTransformerClass ValueTransformerClass) SetValueTransformerForName(transformer INSValueTransformer, name NSValueTransformerName) {
	objc.Send[objc.ID](objc.ID(_ValueTransformerClass.class), objc.Sel("setValueTransformer:forName:"), transformer, objc.String(string(name)))
}

// Returns an array of all the registered value transformers.
//
// # Return Value
//
// An array of all the registered value transformers.
//
// See: https://developer.apple.com/documentation/Foundation/ValueTransformer/valueTransformerNames()
func (_ValueTransformerClass ValueTransformerClass) ValueTransformerNames() []string {
	rv := objc.Send[[]objc.ID](objc.ID(_ValueTransformerClass.class), objc.Sel("valueTransformerNames"))
	return objc.ConvertSliceToStrings(rv)
}

// Returns a Boolean value that indicates whether the receiver can reverse a
// transformation.
//
// # Return Value
//
// true if the receiver supports reverse value transformations, otherwise
// false.
//
// The default is true.
//
// # Discussion
//
// Subclasses should override this method to return false if they do not
// support reverse value transformations.
//
// See: https://developer.apple.com/documentation/Foundation/ValueTransformer/allowsReverseTransformation()
func (_ValueTransformerClass ValueTransformerClass) AllowsReverseTransformation() bool {
	rv := objc.Send[bool](objc.ID(_ValueTransformerClass.class), objc.Sel("allowsReverseTransformation"))
	return rv
}

// Returns the class of the value returned by the receiver for a forward
// transformation.
//
// # Return Value
//
// The class of the value returned by the receiver for a forward
// transformation.
//
// # Discussion
//
// A subclass should override this method to return the appropriate class.
//
// See: https://developer.apple.com/documentation/Foundation/ValueTransformer/transformedValueClass()
func (_ValueTransformerClass ValueTransformerClass) TransformedValueClass() objc.Class {
	rv := objc.Send[objc.Class](objc.ID(_ValueTransformerClass.class), objc.Sel("transformedValueClass"))
	return rv
}
