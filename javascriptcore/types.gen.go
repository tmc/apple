// Code generated from Apple documentation for JavaScriptCore. DO NOT EDIT.

package javascriptcore

// C struct types

// JSClassDefinition - A structure that contains properties and callbacks that define a type of object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSClassDefinition
type JSClassDefinition struct {
	Version           int                               // The version of the class definition structure.
	Attributes        JSClassAttributes                 // A set of class attributes to give to the class.
	ClassName         *byte                             // A null-terminated UTF-8 string that contains the class’s name.
	ParentClass       JSClassRef                        // A JavaScript class to set as the class’s parent class.
	StaticValues      *JSStaticValue                    // An array that contains the class’s statically declared value properties.
	StaticFunctions   *JSStaticFunction                 // An array that contains the class’s statically declared function properties.
	Initialize        JSObjectInitializeCallback        // The callback for creating the object.
	Finalize          JSObjectFinalizeCallback          // The callback for preparing the object for garbage collection.
	HasProperty       JSObjectHasPropertyCallback       // The callback for determining whether an object has a property.
	GetProperty       JSObjectGetPropertyCallback       // The callback for getting a property’s value.
	SetProperty       JSObjectSetPropertyCallback       // The callback for setting a property’s value.
	DeleteProperty    JSObjectDeletePropertyCallback    // The callback for deleting a property.
	GetPropertyNames  JSObjectGetPropertyNamesCallback  // The callback for collecting the names of an object’s properties.
	CallAsFunction    JSObjectCallAsFunctionCallback    // The callback for calling an object as a function.
	CallAsConstructor JSObjectCallAsConstructorCallback // The callback for using an object as a constructor.
	HasInstance       JSObjectHasInstanceCallback       // The callback for checking whether an object is an instance of a particular type.
	ConvertToType     JSObjectConvertToTypeCallback     // The callback for converting an object to a particular JavaScript type.

}

// JSStaticFunction - A statically declared function property.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSStaticFunction
type JSStaticFunction struct {
	Name           *byte                          // A null-terminated UTF-8 string that contains the property’s name.
	CallAsFunction JSObjectCallAsFunctionCallback // A callback to invoke when calling the property as a function.
	Attributes     JSPropertyAttributes           // A set of property attributes to give to the property.

}

// JSStaticValue - A statically declared value property.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/JavaScriptCore/JSStaticValue
type JSStaticValue struct {
	Name        *byte                       // A null-terminated UTF-8 string that contains the property’s name.
	GetProperty JSObjectGetPropertyCallback // A callback to invoke when getting the property’s value.
	SetProperty JSObjectSetPropertyCallback // A callback to invoke when setting the property’s value.
	Attributes  JSPropertyAttributes        // A set of property attributes to give to the property.

}
