// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSStateResourceList] class.
var (
	_MPSStateResourceListClass     MPSStateResourceListClass
	_MPSStateResourceListClassOnce sync.Once
)

func getMPSStateResourceListClass() MPSStateResourceListClass {
	_MPSStateResourceListClassOnce.Do(func() {
		_MPSStateResourceListClass = MPSStateResourceListClass{class: objc.GetClass("MPSStateResourceList")}
	})
	return _MPSStateResourceListClass
}

// GetMPSStateResourceListClass returns the class object for MPSStateResourceList.
func GetMPSStateResourceListClass() MPSStateResourceListClass {
	return getMPSStateResourceListClass()
}

type MPSStateResourceListClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSStateResourceListClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSStateResourceListClass) Alloc() MPSStateResourceList {
	rv := objc.Send[MPSStateResourceList](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// An interface for objects that define resources for Metal Performance
// Shaders state containers.
//
// # Instance Methods
//
//   - [MPSStateResourceList.AppendBuffer]
//   - [MPSStateResourceList.AppendTexture]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSStateResourceList
type MPSStateResourceList struct {
	objectivec.Object
}

// MPSStateResourceListFromID constructs a [MPSStateResourceList] from an objc.ID.
//
// An interface for objects that define resources for Metal Performance
// Shaders state containers.
func MPSStateResourceListFromID(id objc.ID) MPSStateResourceList {
	return MPSStateResourceList{objectivec.Object{ID: id}}
}

// NOTE: MPSStateResourceList adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSStateResourceList] class.
//
// # Instance Methods
//
//   - [IMPSStateResourceList.AppendBuffer]
//   - [IMPSStateResourceList.AppendTexture]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSStateResourceList
type IMPSStateResourceList interface {
	objectivec.IObject

	// Topic: Instance Methods

	AppendBuffer(size uint)
	AppendTexture(descriptor metal.MTLTextureDescriptor)
}

// Init initializes the instance.
func (s MPSStateResourceList) Init() MPSStateResourceList {
	rv := objc.Send[MPSStateResourceList](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s MPSStateResourceList) Autorelease() MPSStateResourceList {
	rv := objc.Send[MPSStateResourceList](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSStateResourceList creates a new MPSStateResourceList instance.
func NewMPSStateResourceList() MPSStateResourceList {
	class := getMPSStateResourceListClass()
	rv := objc.Send[MPSStateResourceList](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSStateResourceList/appendBuffer(_:)
func (s MPSStateResourceList) AppendBuffer(size uint) {
	objc.Send[objc.ID](s.ID, objc.Sel("appendBuffer:"), size)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSStateResourceList/appendTexture(_:)
func (s MPSStateResourceList) AppendTexture(descriptor metal.MTLTextureDescriptor) {
	objc.Send[objc.ID](s.ID, objc.Sel("appendTexture:"), descriptor)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSStateResourceList/resourceList
func (_MPSStateResourceListClass MPSStateResourceListClass) ResourceList() MPSStateResourceList {
	rv := objc.Send[objc.ID](objc.ID(_MPSStateResourceListClass.class), objc.Sel("resourceList"))
	return MPSStateResourceListFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSStateResourceList/resourceListWithBufferSizes:
func (_MPSStateResourceListClass MPSStateResourceListClass) ResourceListWithBufferSizes(firstSize uint) MPSStateResourceList {
	rv := objc.Send[objc.ID](objc.ID(_MPSStateResourceListClass.class), objc.Sel("resourceListWithBufferSizes:"), firstSize)
	return MPSStateResourceListFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSStateResourceList/resourceListWithTextureDescriptors:
func (_MPSStateResourceListClass MPSStateResourceListClass) ResourceListWithTextureDescriptors(d metal.MTLTextureDescriptor) MPSStateResourceList {
	rv := objc.Send[objc.ID](objc.ID(_MPSStateResourceListClass.class), objc.Sel("resourceListWithTextureDescriptors:"), d)
	return MPSStateResourceListFromID(rv)
}
