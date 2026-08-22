// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MTLTensorAuxiliaryPlaneDescriptorMap] class.
var (
	_MTLTensorAuxiliaryPlaneDescriptorMapClass     MTLTensorAuxiliaryPlaneDescriptorMapClass
	_MTLTensorAuxiliaryPlaneDescriptorMapClassOnce sync.Once
)

func getMTLTensorAuxiliaryPlaneDescriptorMapClass() MTLTensorAuxiliaryPlaneDescriptorMapClass {
	_MTLTensorAuxiliaryPlaneDescriptorMapClassOnce.Do(func() {
		_MTLTensorAuxiliaryPlaneDescriptorMapClass = MTLTensorAuxiliaryPlaneDescriptorMapClass{class: objc.GetClass("MTLTensorAuxiliaryPlaneDescriptorMap")}
	})
	return _MTLTensorAuxiliaryPlaneDescriptorMapClass
}

// GetMTLTensorAuxiliaryPlaneDescriptorMapClass returns the class object for MTLTensorAuxiliaryPlaneDescriptorMap.
func GetMTLTensorAuxiliaryPlaneDescriptorMapClass() MTLTensorAuxiliaryPlaneDescriptorMapClass {
	return getMTLTensorAuxiliaryPlaneDescriptorMapClass()
}

type MTLTensorAuxiliaryPlaneDescriptorMapClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MTLTensorAuxiliaryPlaneDescriptorMapClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MTLTensorAuxiliaryPlaneDescriptorMapClass) Alloc() MTLTensorAuxiliaryPlaneDescriptorMap {
	rv := objc.Send[MTLTensorAuxiliaryPlaneDescriptorMap](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A map of auxiliary plane descriptors keyed by plane type.
//
// # Overview
//
// Use this collection to associate [MTLTensorPlaneType] values with
// [MTLTensorAuxiliaryPlaneDescriptor] configurations, then attach it to a
// [MTLTensorDescriptor] to create a multi-plane tensor.
//
// # Instance Methods
//
//   - [MTLTensorAuxiliaryPlaneDescriptorMap.DescriptorForPlane]: Returns the auxiliary plane descriptor for the given plane type, or `nil` if none has been set.
//   - [MTLTensorAuxiliaryPlaneDescriptorMap.Reset]: Empties the map of all its elements.
//   - [MTLTensorAuxiliaryPlaneDescriptorMap.SetDescriptorForPlane]: Sets the auxiliary plane descriptor for the given plane type.
//
// See: https://developer.apple.com/documentation/Metal/MTLTensorAuxiliaryPlaneDescriptorMap
//
// [MTLTensorPlaneType]: https://developer.apple.com/documentation/Metal/MTLTensorPlaneType
type MTLTensorAuxiliaryPlaneDescriptorMap struct {
	objectivec.Object
}

// MTLTensorAuxiliaryPlaneDescriptorMapFromID constructs a [MTLTensorAuxiliaryPlaneDescriptorMap] from an objc.ID.
//
// A map of auxiliary plane descriptors keyed by plane type.
func MTLTensorAuxiliaryPlaneDescriptorMapFromID(id objc.ID) MTLTensorAuxiliaryPlaneDescriptorMap {
	return MTLTensorAuxiliaryPlaneDescriptorMap{objectivec.Object{ID: id}}
}

// NOTE: MTLTensorAuxiliaryPlaneDescriptorMap adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MTLTensorAuxiliaryPlaneDescriptorMap] class.
//
// # Instance Methods
//
//   - [IMTLTensorAuxiliaryPlaneDescriptorMap.DescriptorForPlane]: Returns the auxiliary plane descriptor for the given plane type, or `nil` if none has been set.
//   - [IMTLTensorAuxiliaryPlaneDescriptorMap.Reset]: Empties the map of all its elements.
//   - [IMTLTensorAuxiliaryPlaneDescriptorMap.SetDescriptorForPlane]: Sets the auxiliary plane descriptor for the given plane type.
//
// See: https://developer.apple.com/documentation/Metal/MTLTensorAuxiliaryPlaneDescriptorMap
type IMTLTensorAuxiliaryPlaneDescriptorMap interface {
	objectivec.IObject

	// Topic: Instance Methods

	// Returns the auxiliary plane descriptor for the given plane type, or `nil` if none has been set.
	DescriptorForPlane(plane MTLTensorPlaneType) IMTLTensorAuxiliaryPlaneDescriptor
	// Empties the map of all its elements.
	Reset()
	// Sets the auxiliary plane descriptor for the given plane type.
	SetDescriptorForPlane(descriptor IMTLTensorAuxiliaryPlaneDescriptor, plane MTLTensorPlaneType)
}

// Init initializes the instance.
func (t MTLTensorAuxiliaryPlaneDescriptorMap) Init() MTLTensorAuxiliaryPlaneDescriptorMap {
	rv := objc.Send[MTLTensorAuxiliaryPlaneDescriptorMap](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t MTLTensorAuxiliaryPlaneDescriptorMap) Autorelease() MTLTensorAuxiliaryPlaneDescriptorMap {
	rv := objc.Send[MTLTensorAuxiliaryPlaneDescriptorMap](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTLTensorAuxiliaryPlaneDescriptorMap creates a new MTLTensorAuxiliaryPlaneDescriptorMap instance.
func NewMTLTensorAuxiliaryPlaneDescriptorMap() MTLTensorAuxiliaryPlaneDescriptorMap {
	class := getMTLTensorAuxiliaryPlaneDescriptorMapClass()
	rv := objc.Send[MTLTensorAuxiliaryPlaneDescriptorMap](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Returns the auxiliary plane descriptor for the given plane type, or `nil`
// if none has been set.
//
// plane: The plane type to look up.
//
// # Return Value
//
// The descriptor for the given plane type, or `nil`.
//
// See: https://developer.apple.com/documentation/Metal/MTLTensorAuxiliaryPlaneDescriptorMap/descriptor(for:)
func (t MTLTensorAuxiliaryPlaneDescriptorMap) DescriptorForPlane(plane MTLTensorPlaneType) IMTLTensorAuxiliaryPlaneDescriptor {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("descriptorForPlane:"), plane)
	return MTLTensorAuxiliaryPlaneDescriptorFromID(rv)
}

// Empties the map of all its elements.
//
// See: https://developer.apple.com/documentation/Metal/MTLTensorAuxiliaryPlaneDescriptorMap/reset()
func (t MTLTensorAuxiliaryPlaneDescriptorMap) Reset() {
	objc.Send[objc.ID](t.ID, objc.Sel("reset"))
}

// Sets the auxiliary plane descriptor for the given plane type.
//
// descriptor: The descriptor configuring the auxiliary plane.
//
// plane: The plane type to associate the descriptor with.
//
// # Discussion
//
// [MTLTensorPlaneType.data] is not a valid plane type for this method. The
// data plane is always present, and you configure it directly on
// [MTLTensorDescriptor].
//
// [MTLTensorPlaneType.scales] auxiliary planes only support
// [MTLTensorDataType.metalFloat8ue8m0] as a data type.
//
// See: https://developer.apple.com/documentation/Metal/MTLTensorAuxiliaryPlaneDescriptorMap/setDescriptor(_:for:)
//
// [MTLTensorDataType.metalFloat8ue8m0]: https://developer.apple.com/documentation/Metal/MTLTensorDataType/metalFloat8ue8m0
// [MTLTensorPlaneType.data]: https://developer.apple.com/documentation/Metal/MTLTensorPlaneType/data
// [MTLTensorPlaneType.scales]: https://developer.apple.com/documentation/Metal/MTLTensorPlaneType/scales
func (t MTLTensorAuxiliaryPlaneDescriptorMap) SetDescriptorForPlane(descriptor IMTLTensorAuxiliaryPlaneDescriptor, plane MTLTensorPlaneType) {
	objc.Send[objc.ID](t.ID, objc.Sel("setDescriptor:forPlane:"), descriptor, plane)
}
