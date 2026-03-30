// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MTL4PipelineDescriptor] class.
var (
	_MTL4PipelineDescriptorClass     MTL4PipelineDescriptorClass
	_MTL4PipelineDescriptorClassOnce sync.Once
)

func getMTL4PipelineDescriptorClass() MTL4PipelineDescriptorClass {
	_MTL4PipelineDescriptorClassOnce.Do(func() {
		_MTL4PipelineDescriptorClass = MTL4PipelineDescriptorClass{class: objc.GetClass("MTL4PipelineDescriptor")}
	})
	return _MTL4PipelineDescriptorClass
}

// GetMTL4PipelineDescriptorClass returns the class object for MTL4PipelineDescriptor.
func GetMTL4PipelineDescriptorClass() MTL4PipelineDescriptorClass {
	return getMTL4PipelineDescriptorClass()
}

type MTL4PipelineDescriptorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MTL4PipelineDescriptorClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MTL4PipelineDescriptorClass) Alloc() MTL4PipelineDescriptor {
	rv := objc.Send[MTL4PipelineDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// Base type for descriptors you use for building pipeline state objects.
//
// # Instance Properties
//
//   - [MTL4PipelineDescriptor.Label]: Assigns an optional string that uniquely identifies a pipeline descriptor.
//   - [MTL4PipelineDescriptor.SetLabel]
//   - [MTL4PipelineDescriptor.Options]: Provides compile-time options when you build the pipeline.
//   - [MTL4PipelineDescriptor.SetOptions]
//
// See: https://developer.apple.com/documentation/Metal/MTL4PipelineDescriptor
type MTL4PipelineDescriptor struct {
	objectivec.Object
}

// MTL4PipelineDescriptorFromID constructs a [MTL4PipelineDescriptor] from an objc.ID.
//
// Base type for descriptors you use for building pipeline state objects.
func MTL4PipelineDescriptorFromID(id objc.ID) MTL4PipelineDescriptor {
	return MTL4PipelineDescriptor{objectivec.Object{ID: id}}
}

// NOTE: MTL4PipelineDescriptor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MTL4PipelineDescriptor] class.
//
// # Instance Properties
//
//   - [IMTL4PipelineDescriptor.Label]: Assigns an optional string that uniquely identifies a pipeline descriptor.
//   - [IMTL4PipelineDescriptor.SetLabel]
//   - [IMTL4PipelineDescriptor.Options]: Provides compile-time options when you build the pipeline.
//   - [IMTL4PipelineDescriptor.SetOptions]
//
// See: https://developer.apple.com/documentation/Metal/MTL4PipelineDescriptor
type IMTL4PipelineDescriptor interface {
	objectivec.IObject

	// Topic: Instance Properties

	// Assigns an optional string that uniquely identifies a pipeline descriptor.
	Label() string
	SetLabel(value string)
	// Provides compile-time options when you build the pipeline.
	Options() IMTL4PipelineOptions
	SetOptions(value IMTL4PipelineOptions)
}

// Init initializes the instance.
func (m MTL4PipelineDescriptor) Init() MTL4PipelineDescriptor {
	rv := objc.Send[MTL4PipelineDescriptor](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MTL4PipelineDescriptor) Autorelease() MTL4PipelineDescriptor {
	rv := objc.Send[MTL4PipelineDescriptor](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTL4PipelineDescriptor creates a new MTL4PipelineDescriptor instance.
func NewMTL4PipelineDescriptor() MTL4PipelineDescriptor {
	class := getMTL4PipelineDescriptorClass()
	rv := objc.Send[MTL4PipelineDescriptor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Assigns an optional string that uniquely identifies a pipeline descriptor.
//
// # Discussion
//
// After you provide this label, you can use it to look up a pipeline state
// object by name in a binary archive.
//
// See: https://developer.apple.com/documentation/Metal/MTL4PipelineDescriptor/label
func (m MTL4PipelineDescriptor) Label() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("label"))
	return foundation.NSStringFromID(rv).String()
}
func (m MTL4PipelineDescriptor) SetLabel(value string) {
	objc.Send[struct{}](m.ID, objc.Sel("setLabel:"), objc.String(value))
}

// Provides compile-time options when you build the pipeline.
//
// See: https://developer.apple.com/documentation/Metal/MTL4PipelineDescriptor/options
func (m MTL4PipelineDescriptor) Options() IMTL4PipelineOptions {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("options"))
	return MTL4PipelineOptionsFromID(objc.ID(rv))
}
func (m MTL4PipelineDescriptor) SetOptions(value IMTL4PipelineOptions) {
	objc.Send[struct{}](m.ID, objc.Sel("setOptions:"), value)
}
