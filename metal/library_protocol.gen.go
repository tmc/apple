// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// A collection of Metal shader functions.
//
// See: https://developer.apple.com/documentation/Metal/MTLLibrary
type MTLLibrary interface {
	objectivec.IObject

	// The installation name for a dynamic library.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLLibrary/installName
	InstallName() string

	// The library’s basic type.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLLibrary/type
	Type() MTLLibraryType

	// The names of all public functions in the library.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLLibrary/functionNames
	FunctionNames() []string

	// Creates an instance that represents a shader function in the library.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLLibrary/makeFunction(name:)
	NewFunctionWithName(functionName string) MTLFunction

	// Asynchronously creates a specialized shader function.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLLibrary/makeFunction(name:constantValues:completionHandler:)
	NewFunctionWithNameConstantValuesCompletionHandler(name string, constantValues IMTLFunctionConstantValues, completionHandler MTLFunctionErrorHandler)

	// Synchronously creates a specialized shader function.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLLibrary/makeFunction(name:constantValues:)
	NewFunctionWithNameConstantValuesError(name string, constantValues IMTLFunctionConstantValues) (MTLFunction, error)

	// Asynchronously creates an object representing a shader function, using the specified descriptor.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLLibrary/makeFunction(descriptor:completionHandler:)
	NewFunctionWithDescriptorCompletionHandler(descriptor IMTLFunctionDescriptor, completionHandler MTLFunctionErrorHandler)

	// Synchronously creates an object representing a shader function, using the specified descriptor.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLLibrary/makeFunction(descriptor:)
	NewFunctionWithDescriptorError(descriptor IMTLFunctionDescriptor) (MTLFunction, error)

	// Asynchronously creates an object representing a ray-tracing intersection function, using the specified descriptor.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLLibrary/makeIntersectionFunction(descriptor:completionHandler:)
	NewIntersectionFunctionWithDescriptorCompletionHandler(descriptor IMTLIntersectionFunctionDescriptor, completionHandler MTLFunctionErrorHandler)

	// Synchronously creates an object representing a ray-tracing intersection function, using the specified descriptor.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLLibrary/makeIntersectionFunction(descriptor:)
	NewIntersectionFunctionWithDescriptorError(descriptor IMTLIntersectionFunctionDescriptor) (MTLFunction, error)

	// The Metal device object that created the library.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLLibrary/device
	Device() MTLDevice

	// A string that identifies the library.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLLibrary/label
	Label() string

	// Retrieves reflection information for a function in the library.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLLibrary/reflection(functionName:)
	ReflectionForFunctionWithName(functionName string) IMTLFunctionReflection

	// A string that identifies the library.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLLibrary/label
	SetLabel(value string)
}

// MTLLibraryObject wraps an existing Objective-C object that conforms to the MTLLibrary protocol.
type MTLLibraryObject struct {
	objectivec.Object
}
func (o MTLLibraryObject) BaseObject() objectivec.Object {
	return o.Object
}

// MTLLibraryObjectFromID constructs a [MTLLibraryObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func MTLLibraryObjectFromID(id objc.ID) MTLLibraryObject {
	return MTLLibraryObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The installation name for a dynamic library.
//
// See: https://developer.apple.com/documentation/Metal/MTLLibrary/installName

func (o MTLLibraryObject) InstallName() string {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("installName"))
	return foundation.NSStringFromID(rv).String()
	}

// The library’s basic type.
//
// See: https://developer.apple.com/documentation/Metal/MTLLibrary/type

func (o MTLLibraryObject) Type() MTLLibraryType {
	
	rv := objc.Send[MTLLibraryType](o.ID, objc.Sel("type"))
	return rv
	}

// The names of all public functions in the library.
//
// See: https://developer.apple.com/documentation/Metal/MTLLibrary/functionNames

func (o MTLLibraryObject) FunctionNames() []string {
	
	rv := objc.Send[[]objc.ID](o.ID, objc.Sel("functionNames"))
	return objc.ConvertSliceToStrings(rv)
	}

// Creates an instance that represents a shader function in the library.
//
// functionName: The name of the function.
//
// # Return Value
// 
// An [MTLFunction], or `nil` if the named function isn’t found in the
// library.
//
// # Discussion
// 
// If you call this method to retrieve a function that doesn’t use function
// constants, it returns an [MTLFunction] instance that you can use to build a
// render or compute pipeline.
// 
// If you call this method to retrieve a function that uses function constants
// to specialize its behavior, you can only use the returned instance to query
// the `functionConstants` property for the list of function constants. You
// can’t use to create a render or compute pipeline. To get a specialized
// instance that you can use to create a pipeline instance, call the
// [NewFunctionWithNameConstantValuesCompletionHandler] method or
// [NewFunctionWithNameConstantValuesError] to generate a specialized
// function.
//
// See: https://developer.apple.com/documentation/Metal/MTLLibrary/makeFunction(name:)

func (o MTLLibraryObject) NewFunctionWithName(functionName string) MTLFunction {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("newFunctionWithName:"), objc.String(functionName))
	return MTLFunctionObjectFromID(rv)
	}

// Asynchronously creates a specialized shader function.
//
// name: The name of the specialized function.
//
// constantValues: The set of constant values assigned to the function constants. Compilation
// fails if you don’t provide valid constant values for all required
// function constants.
//
// completionHandler: A block of code that Metal calls after it creates the specialized function.
// 
// function: A specialized function, or `nil` if an error occurred. error: An
// error object that describes compilation problems, if any. This object
// contains compiler errors if the specialized function is `nil`, and compiler
// warnings if Metal created the specialized function with warnings. If Metal
// created the function without errors or warnings, this error object is
// `nil`.
//
// # Discussion
// 
// Function constant values are first looked up by their index, then by their
// name. Metal ignores any values that don’t correspond to a function
// constant in the named function without generating errors or warnings.
//
// See: https://developer.apple.com/documentation/Metal/MTLLibrary/makeFunction(name:constantValues:completionHandler:)

func (o MTLLibraryObject) NewFunctionWithNameConstantValuesCompletionHandler(name string, constantValues IMTLFunctionConstantValues, completionHandler MTLFunctionErrorHandler) {
	
	objc.Send[struct{}](o.ID, objc.Sel("newFunctionWithName:constantValues:completionHandler:"), objc.String(name), constantValues, completionHandler)
	}

// Synchronously creates a specialized shader function.
//
// name: The name of the specialized function.
//
// constantValues: The set of constant values for the function constants. The compiler can’t
// compile the function if any value is invalid for the function constants it
// requires.
//
// # Return Value
// 
// A new [MTLFunction] instance if the method completes successfully;
// otherwise Swift throws an error and Objective-C returns `nil`.
//
// # Discussion
// 
// Function constant values are first looked up by their index, then by their
// name. The compiler ignores any values that don’t correspond to a function
// constant in the named function, and doesn’t generate errors or warnings.
//
// See: https://developer.apple.com/documentation/Metal/MTLLibrary/makeFunction(name:constantValues:)

func (o MTLLibraryObject) NewFunctionWithNameConstantValuesError(name string, constantValues IMTLFunctionConstantValues) (MTLFunction, error) {
	
	rv, err := objc.SendWithError[objc.ID](o.ID, objc.Sel("newFunctionWithName:constantValues:error:"), objc.String(name), constantValues)
	if err != nil {
		return nil, err
	}
	return MTLFunctionObjectFromID(rv), nil
	}

// Asynchronously creates an object representing a shader function, using the
// specified descriptor.
//
// descriptor: The description of the function object to create.
//
// completionHandler: A Swift closure or an Objective-C block that Metal calls after it creates
// the function.
//
// See: https://developer.apple.com/documentation/Metal/MTLLibrary/makeFunction(descriptor:completionHandler:)

func (o MTLLibraryObject) NewFunctionWithDescriptorCompletionHandler(descriptor IMTLFunctionDescriptor, completionHandler MTLFunctionErrorHandler) {
	
	objc.Send[struct{}](o.ID, objc.Sel("newFunctionWithDescriptor:completionHandler:"), descriptor, completionHandler)
	}

// Synchronously creates an object representing a shader function, using the
// specified descriptor.
//
// descriptor: The description of the function object to create.
//
// # Return Value
// 
// A new [MTLFunction] instance if the method finds the function in the
// library; otherwise Swift throws an error and Objective-C returns `nil`.
//
// See: https://developer.apple.com/documentation/Metal/MTLLibrary/makeFunction(descriptor:)

func (o MTLLibraryObject) NewFunctionWithDescriptorError(descriptor IMTLFunctionDescriptor) (MTLFunction, error) {
	
	rv, err := objc.SendWithError[objc.ID](o.ID, objc.Sel("newFunctionWithDescriptor:error:"), descriptor)
	if err != nil {
		return nil, err
	}
	return MTLFunctionObjectFromID(rv), nil
	}

// Asynchronously creates an object representing a ray-tracing intersection
// function, using the specified descriptor.
//
// See: https://developer.apple.com/documentation/Metal/MTLLibrary/makeIntersectionFunction(descriptor:completionHandler:)

func (o MTLLibraryObject) NewIntersectionFunctionWithDescriptorCompletionHandler(descriptor IMTLIntersectionFunctionDescriptor, completionHandler MTLFunctionErrorHandler) {
	
	objc.Send[struct{}](o.ID, objc.Sel("newIntersectionFunctionWithDescriptor:completionHandler:"), descriptor, completionHandler)
	}

// Synchronously creates an object representing a ray-tracing intersection
// function, using the specified descriptor.
//
// See: https://developer.apple.com/documentation/Metal/MTLLibrary/makeIntersectionFunction(descriptor:)

func (o MTLLibraryObject) NewIntersectionFunctionWithDescriptorError(descriptor IMTLIntersectionFunctionDescriptor) (MTLFunction, error) {
	
	rv, err := objc.SendWithError[objc.ID](o.ID, objc.Sel("newIntersectionFunctionWithDescriptor:error:"), descriptor)
	if err != nil {
		return nil, err
	}
	return MTLFunctionObjectFromID(rv), nil
	}

// The Metal device object that created the library.
//
// See: https://developer.apple.com/documentation/Metal/MTLLibrary/device

func (o MTLLibraryObject) Device() MTLDevice {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("device"))
	return MTLDeviceObjectFromID(rv)
	}

// A string that identifies the library.
//
// See: https://developer.apple.com/documentation/Metal/MTLLibrary/label

func (o MTLLibraryObject) Label() string {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("label"))
	return foundation.NSStringFromID(rv).String()
	}

// Retrieves reflection information for a function in the library.
//
// functionName: The name of a GPU function in the library. The name needs to match one of
// the elements in the string array of library’s [FunctionNames] property.
//
// # Return Value
// 
// An [MTLFunctionReflection] instance when the method succeeds; otherwise
// `nil`.
//
// # Discussion
// 
// The reflection instance contains metadata information about a specific GPU
// function, which can include:
// 
// - Function parameters - Return types - Bindings - Annotations from a
// developer, if available
// 
// The method only returns reflection information if all of the following
// conditions apply:
// 
// - The library has a function with a name that matches `functionName`. - The
// deployment target is macOS 13.0 or later, or iOS 16.0 or later, or visionOS
// 2.0 or later.
//
// See: https://developer.apple.com/documentation/Metal/MTLLibrary/reflection(functionName:)

func (o MTLLibraryObject) ReflectionForFunctionWithName(functionName string) IMTLFunctionReflection {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("reflectionForFunctionWithName:"), objc.String(functionName))
	return MTLFunctionReflectionFromID(rv)
	}

func (o MTLLibraryObject) SetLabel(value string) {
	objc.Send[struct{}](o.ID, objc.Sel("setLabel:"), objc.String(value))
}

