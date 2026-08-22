// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSCNNLossDataDescriptor] class.
var (
	_MPSCNNLossDataDescriptorClass     MPSCNNLossDataDescriptorClass
	_MPSCNNLossDataDescriptorClassOnce sync.Once
)

func getMPSCNNLossDataDescriptorClass() MPSCNNLossDataDescriptorClass {
	_MPSCNNLossDataDescriptorClassOnce.Do(func() {
		_MPSCNNLossDataDescriptorClass = MPSCNNLossDataDescriptorClass{class: objc.GetClass("MPSCNNLossDataDescriptor")}
	})
	return _MPSCNNLossDataDescriptorClass
}

// GetMPSCNNLossDataDescriptorClass returns the class object for MPSCNNLossDataDescriptor.
func GetMPSCNNLossDataDescriptorClass() MPSCNNLossDataDescriptorClass {
	return getMPSCNNLossDataDescriptorClass()
}

type MPSCNNLossDataDescriptorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSCNNLossDataDescriptorClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSCNNLossDataDescriptorClass) Alloc() MPSCNNLossDataDescriptor {
	rv := objc.Send[MPSCNNLossDataDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// An object that specifies properties used by a loss data descriptor.
//
// # Instance Properties
//
//   - [MPSCNNLossDataDescriptor.BytesPerImage]
//   - [MPSCNNLossDataDescriptor.SetBytesPerImage]
//   - [MPSCNNLossDataDescriptor.BytesPerRow]
//   - [MPSCNNLossDataDescriptor.SetBytesPerRow]
//   - [MPSCNNLossDataDescriptor.Layout]
//   - [MPSCNNLossDataDescriptor.Size]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNLossDataDescriptor
type MPSCNNLossDataDescriptor struct {
	objectivec.Object
}

// MPSCNNLossDataDescriptorFromID constructs a [MPSCNNLossDataDescriptor] from an objc.ID.
//
// An object that specifies properties used by a loss data descriptor.
func MPSCNNLossDataDescriptorFromID(id objc.ID) MPSCNNLossDataDescriptor {
	return MPSCNNLossDataDescriptor{objectivec.Object{ID: id}}
}

// NOTE: MPSCNNLossDataDescriptor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSCNNLossDataDescriptor] class.
//
// # Instance Properties
//
//   - [IMPSCNNLossDataDescriptor.BytesPerImage]
//   - [IMPSCNNLossDataDescriptor.SetBytesPerImage]
//   - [IMPSCNNLossDataDescriptor.BytesPerRow]
//   - [IMPSCNNLossDataDescriptor.SetBytesPerRow]
//   - [IMPSCNNLossDataDescriptor.Layout]
//   - [IMPSCNNLossDataDescriptor.Size]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNLossDataDescriptor
type IMPSCNNLossDataDescriptor interface {
	objectivec.IObject

	// Topic: Instance Properties

	BytesPerImage() uint
	SetBytesPerImage(value uint)
	BytesPerRow() uint
	SetBytesPerRow(value uint)
	Layout() MPSDataLayout
	Size() metal.MTLSize
}

// Init initializes the instance.
func (c MPSCNNLossDataDescriptor) Init() MPSCNNLossDataDescriptor {
	rv := objc.Send[MPSCNNLossDataDescriptor](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSCNNLossDataDescriptor) Autorelease() MPSCNNLossDataDescriptor {
	rv := objc.Send[MPSCNNLossDataDescriptor](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSCNNLossDataDescriptor creates a new MPSCNNLossDataDescriptor instance.
func NewMPSCNNLossDataDescriptor() MPSCNNLossDataDescriptor {
	class := getMPSCNNLossDataDescriptorClass()
	rv := objc.Send[MPSCNNLossDataDescriptor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNLossDataDescriptor/init(data:layout:size:)
func NewCNNLossDataDescriptorCnnLossDataDescriptorWithDataLayoutSize(data foundation.NSData, layout MPSDataLayout, size metal.MTLSize) MPSCNNLossDataDescriptor {
	rv := objc.Send[objc.ID](objc.ID(getMPSCNNLossDataDescriptorClass().class), objc.Sel("cnnLossDataDescriptorWithData:layout:size:"), data, layout, size)
	return MPSCNNLossDataDescriptorFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNLossDataDescriptor/bytesPerImage
func (c MPSCNNLossDataDescriptor) BytesPerImage() uint {
	rv := objc.Send[uint](c.ID, objc.Sel("bytesPerImage"))
	return rv
}
func (c MPSCNNLossDataDescriptor) SetBytesPerImage(value uint) {
	objc.Send[struct{}](c.ID, objc.Sel("setBytesPerImage:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNLossDataDescriptor/bytesPerRow
func (c MPSCNNLossDataDescriptor) BytesPerRow() uint {
	rv := objc.Send[uint](c.ID, objc.Sel("bytesPerRow"))
	return rv
}
func (c MPSCNNLossDataDescriptor) SetBytesPerRow(value uint) {
	objc.Send[struct{}](c.ID, objc.Sel("setBytesPerRow:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNLossDataDescriptor/layout
func (c MPSCNNLossDataDescriptor) Layout() MPSDataLayout {
	rv := objc.Send[MPSDataLayout](c.ID, objc.Sel("layout"))
	return MPSDataLayout(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNLossDataDescriptor/size
func (c MPSCNNLossDataDescriptor) Size() metal.MTLSize {
	rv := objc.Send[metal.MTLSize](c.ID, objc.Sel("size"))
	return metal.MTLSize(rv)
}
