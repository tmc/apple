// Code generated from Apple documentation for iogpu. DO NOT EDIT.

package iogpu

import (
	"unsafe"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// MTLTensorSPI protocol.
type MTLTensorSPI interface {
	objectivec.IObject

	// GetBytesStridesFromSlice protocol.
	GetBytesStridesFromSlice(bytes unsafe.Pointer, strides objectivec.IObject, slice MTLTensorSlice)

	// InternalMTLBuffer protocol.
	InternalMTLBuffer() objectivec.IObject

	// IsTensorViewableWithReshapedDescriptor protocol.
	IsTensorViewableWithReshapedDescriptor(descriptor objectivec.IObject) bool

	// NewTensorViewWithReshapedDescriptorError protocol.
	NewTensorViewWithReshapedDescriptorError(descriptor objectivec.IObject) (objectivec.IObject, error)

	// NewTensorViewWithSliceError protocol.
	NewTensorViewWithSliceError(slice MTLTensorSlice) (objectivec.IObject, error)

	// Offset protocol.
	Offset() uint64

	// ParentTensor protocol.
	ParentTensor() objectivec.IObject

	// ReplaceSliceWithBytesStrides protocol.
	ReplaceSliceWithBytesStrides(slice MTLTensorSlice, bytes unsafe.Pointer, strides objectivec.IObject)

	// ResourceIndex protocol.
	ResourceIndex() uint64
}

// MTLTensorSPIObject wraps an existing Objective-C object that conforms to the MTLTensorSPI protocol.
type MTLTensorSPIObject struct {
	objectivec.Object
}

func (o MTLTensorSPIObject) BaseObject() objectivec.Object {
	return o.Object
}

// MTLTensorSPIObjectFromID constructs a [MTLTensorSPIObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func MTLTensorSPIObjectFromID(id objc.ID) MTLTensorSPIObject {
	return MTLTensorSPIObject{
		Object: objectivec.ObjectFromID(id),
	}
}

func (o MTLTensorSPIObject) GetBytesStridesFromSlice(bytes unsafe.Pointer, strides objectivec.IObject, slice MTLTensorSlice) {
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("getBytes:strides:fromSlice:"), bytes, strides, slice)
}
func (o MTLTensorSPIObject) InternalMTLBuffer() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](o.ID, objc.Sel("internalMTLBuffer"))
	return objectivec.Object{ID: rv}
}
func (o MTLTensorSPIObject) IsTensorViewableWithReshapedDescriptor(descriptor objectivec.IObject) bool {
	rv := objc.SendIfResponds[bool](o.ID, objc.Sel("isTensorViewableWithReshapedDescriptor:"), descriptor)
	return rv
}
func (o MTLTensorSPIObject) NewTensorViewWithReshapedDescriptorError(descriptor objectivec.IObject) (objectivec.IObject, error) {
	rv, err := objc.SendWithError[objc.ID](o.ID, objc.Sel("newTensorViewWithReshapedDescriptor:error:"), descriptor)
	if err != nil {
		return nil, err
	}
	return objectivec.Object{ID: rv}, nil
}
func (o MTLTensorSPIObject) NewTensorViewWithSliceError(slice MTLTensorSlice) (objectivec.IObject, error) {
	rv, err := objc.SendWithError[objc.ID](o.ID, objc.Sel("newTensorViewWithSlice:error:"), slice)
	if err != nil {
		return nil, err
	}
	return objectivec.Object{ID: rv}, nil
}
func (o MTLTensorSPIObject) Offset() uint64 {
	rv := objc.SendIfResponds[uint64](o.ID, objc.Sel("offset"))
	return rv
}
func (o MTLTensorSPIObject) ParentTensor() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](o.ID, objc.Sel("parentTensor"))
	return objectivec.Object{ID: rv}
}
func (o MTLTensorSPIObject) ReplaceSliceWithBytesStrides(slice MTLTensorSlice, bytes unsafe.Pointer, strides objectivec.IObject) {
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("replaceSlice:withBytes:strides:"), slice, bytes, strides)
}
func (o MTLTensorSPIObject) ResourceIndex() uint64 {
	rv := objc.SendIfResponds[uint64](o.ID, objc.Sel("resourceIndex"))
	return rv
}
