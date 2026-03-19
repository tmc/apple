// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// Represents a binary function.
//
// See: https://developer.apple.com/documentation/Metal/MTL4BinaryFunction
type MTL4BinaryFunction interface {
	objectivec.IObject

	// Describes the type of this binary function.
	//
	// See: https://developer.apple.com/documentation/Metal/MTL4BinaryFunction/functionType
	FunctionType() MTLFunctionType

	// Obtains the optional name of this binary function.
	//
	// See: https://developer.apple.com/documentation/Metal/MTL4BinaryFunction/name
	Name() string
}

// MTL4BinaryFunctionObject wraps an existing Objective-C object that conforms to the MTL4BinaryFunction protocol.
type MTL4BinaryFunctionObject struct {
	objectivec.Object
}
func (o MTL4BinaryFunctionObject) BaseObject() objectivec.Object {
	return o.Object
}

// MTL4BinaryFunctionObjectFromID constructs a [MTL4BinaryFunctionObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func MTL4BinaryFunctionObjectFromID(id objc.ID) MTL4BinaryFunctionObject {
	return MTL4BinaryFunctionObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Describes the type of this binary function.
//
// See: https://developer.apple.com/documentation/Metal/MTL4BinaryFunction/functionType
func (o MTL4BinaryFunctionObject) FunctionType() MTLFunctionType {
	
	rv := objc.Send[MTLFunctionType](o.ID, objc.Sel("functionType"))
	return rv
	}
// Obtains the optional name of this binary function.
//
// See: https://developer.apple.com/documentation/Metal/MTL4BinaryFunction/name
func (o MTL4BinaryFunctionObject) Name() string {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("name"))
	return foundation.NSStringFromID(rv).String()
	}

