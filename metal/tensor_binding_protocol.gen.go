// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// An object that represents a tensor bound to a graphics or compute function or a machine learning function.
//
// See: https://developer.apple.com/documentation/Metal/MTLTensorBinding
type MTLTensorBinding interface {
	objectivec.IObject
	MTLBinding

	// The array of sizes, in elements, one for each dimension of this tensor.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLTensorBinding/dimensions
	Dimensions() IMTLTensorExtents

	// The data format you use for indexing into the tensor.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLTensorBinding/indexType
	IndexType() MTLDataType

	// The underlying data format of this tensor.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLTensorBinding/tensorDataType
	TensorDataType() MTLTensorDataType
}

// MTLTensorBindingObject wraps an existing Objective-C object that conforms to the MTLTensorBinding protocol.
type MTLTensorBindingObject struct {
	objectivec.Object
}
func (o MTLTensorBindingObject) BaseObject() objectivec.Object {
	return o.Object
}

// MTLTensorBindingObjectFromID constructs a [MTLTensorBindingObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func MTLTensorBindingObjectFromID(id objc.ID) MTLTensorBindingObject {
	return MTLTensorBindingObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The array of sizes, in elements, one for each dimension of this tensor.
//
// See: https://developer.apple.com/documentation/Metal/MTLTensorBinding/dimensions
func (o MTLTensorBindingObject) Dimensions() IMTLTensorExtents {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("dimensions"))
	return MTLTensorExtentsFromID(rv)
	}
// The data format you use for indexing into the tensor.
//
// See: https://developer.apple.com/documentation/Metal/MTLTensorBinding/indexType
func (o MTLTensorBindingObject) IndexType() MTLDataType {
	
	rv := objc.Send[MTLDataType](o.ID, objc.Sel("indexType"))
	return rv
	}
// The underlying data format of this tensor.
//
// See: https://developer.apple.com/documentation/Metal/MTLTensorBinding/tensorDataType
func (o MTLTensorBindingObject) TensorDataType() MTLTensorDataType {
	
	rv := objc.Send[MTLTensorDataType](o.ID, objc.Sel("tensorDataType"))
	return rv
	}
// See: https://developer.apple.com/documentation/Metal/MTLBinding/access
func (o MTLTensorBindingObject) Access() MTLBindingAccess {
	
	rv := objc.Send[MTLBindingAccess](o.ID, objc.Sel("access"))
	return rv
	}
// See: https://developer.apple.com/documentation/Metal/MTLBinding/index
func (o MTLTensorBindingObject) Index() uint {
	
	rv := objc.Send[uint](o.ID, objc.Sel("index"))
	return rv
	}
// See: https://developer.apple.com/documentation/Metal/MTLBinding/isArgument
func (o MTLTensorBindingObject) IsArgument() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("isArgument"))
	return rv
	}
// See: https://developer.apple.com/documentation/Metal/MTLBinding/isUsed
func (o MTLTensorBindingObject) IsUsed() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("isUsed"))
	return rv
	}
// See: https://developer.apple.com/documentation/Metal/MTLBinding/name
func (o MTLTensorBindingObject) Name() string {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("name"))
	return foundation.NSStringFromID(rv).String()
	}
// See: https://developer.apple.com/documentation/Metal/MTLBinding/type
func (o MTLTensorBindingObject) Type() MTLBindingType {
	
	rv := objc.Send[MTLBindingType](o.ID, objc.Sel("type"))
	return rv
	}

