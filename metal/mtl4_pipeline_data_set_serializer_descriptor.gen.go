// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MTL4PipelineDataSetSerializerDescriptor] class.
var (
	_MTL4PipelineDataSetSerializerDescriptorClass     MTL4PipelineDataSetSerializerDescriptorClass
	_MTL4PipelineDataSetSerializerDescriptorClassOnce sync.Once
)

func getMTL4PipelineDataSetSerializerDescriptorClass() MTL4PipelineDataSetSerializerDescriptorClass {
	_MTL4PipelineDataSetSerializerDescriptorClassOnce.Do(func() {
		_MTL4PipelineDataSetSerializerDescriptorClass = MTL4PipelineDataSetSerializerDescriptorClass{class: objc.GetClass("MTL4PipelineDataSetSerializerDescriptor")}
	})
	return _MTL4PipelineDataSetSerializerDescriptorClass
}

// GetMTL4PipelineDataSetSerializerDescriptorClass returns the class object for MTL4PipelineDataSetSerializerDescriptor.
func GetMTL4PipelineDataSetSerializerDescriptorClass() MTL4PipelineDataSetSerializerDescriptorClass {
	return getMTL4PipelineDataSetSerializerDescriptorClass()
}

type MTL4PipelineDataSetSerializerDescriptorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MTL4PipelineDataSetSerializerDescriptorClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MTL4PipelineDataSetSerializerDescriptorClass) Alloc() MTL4PipelineDataSetSerializerDescriptor {
	rv := objc.Send[MTL4PipelineDataSetSerializerDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// Groups together properties to create a pipeline data set serializer.
//
// # Instance Properties
//
//   - [MTL4PipelineDataSetSerializerDescriptor.Configuration]: Specifies the configuration of the serialization process.
//   - [MTL4PipelineDataSetSerializerDescriptor.SetConfiguration]
//
// See: https://developer.apple.com/documentation/Metal/MTL4PipelineDataSetSerializerDescriptor
type MTL4PipelineDataSetSerializerDescriptor struct {
	objectivec.Object
}

// MTL4PipelineDataSetSerializerDescriptorFromID constructs a [MTL4PipelineDataSetSerializerDescriptor] from an objc.ID.
//
// Groups together properties to create a pipeline data set serializer.
func MTL4PipelineDataSetSerializerDescriptorFromID(id objc.ID) MTL4PipelineDataSetSerializerDescriptor {
	return MTL4PipelineDataSetSerializerDescriptor{objectivec.Object{ID: id}}
}

// NOTE: MTL4PipelineDataSetSerializerDescriptor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MTL4PipelineDataSetSerializerDescriptor] class.
//
// # Instance Properties
//
//   - [IMTL4PipelineDataSetSerializerDescriptor.Configuration]: Specifies the configuration of the serialization process.
//   - [IMTL4PipelineDataSetSerializerDescriptor.SetConfiguration]
//
// See: https://developer.apple.com/documentation/Metal/MTL4PipelineDataSetSerializerDescriptor
type IMTL4PipelineDataSetSerializerDescriptor interface {
	objectivec.IObject

	// Topic: Instance Properties

	// Specifies the configuration of the serialization process.
	Configuration() MTL4PipelineDataSetSerializerConfiguration
	SetConfiguration(value MTL4PipelineDataSetSerializerConfiguration)
}

// Init initializes the instance.
func (m MTL4PipelineDataSetSerializerDescriptor) Init() MTL4PipelineDataSetSerializerDescriptor {
	rv := objc.Send[MTL4PipelineDataSetSerializerDescriptor](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MTL4PipelineDataSetSerializerDescriptor) Autorelease() MTL4PipelineDataSetSerializerDescriptor {
	rv := objc.Send[MTL4PipelineDataSetSerializerDescriptor](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTL4PipelineDataSetSerializerDescriptor creates a new MTL4PipelineDataSetSerializerDescriptor instance.
func NewMTL4PipelineDataSetSerializerDescriptor() MTL4PipelineDataSetSerializerDescriptor {
	class := getMTL4PipelineDataSetSerializerDescriptorClass()
	rv := objc.Send[MTL4PipelineDataSetSerializerDescriptor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Specifies the configuration of the serialization process.
//
// # Discussion
//
// The configuration of the serialization process determines the mechanisms
// you use to serialize pipeline data sets.
//
// When this configuration contains
// [MTL4PipelineDataSetSerializerConfigurationCaptureDescriptors], use “ to
// serialize pipeline scripts.
//
// If this option contains
// [MTL4PipelineDataSetSerializerConfigurationCaptureBinaries], the serializer
// can additionally serialize to a binary archive by calling `:`.
//
// See: https://developer.apple.com/documentation/Metal/MTL4PipelineDataSetSerializerDescriptor/configuration
func (m MTL4PipelineDataSetSerializerDescriptor) Configuration() MTL4PipelineDataSetSerializerConfiguration {
	rv := objc.Send[MTL4PipelineDataSetSerializerConfiguration](m.ID, objc.Sel("configuration"))
	return MTL4PipelineDataSetSerializerConfiguration(rv)
}
func (m MTL4PipelineDataSetSerializerDescriptor) SetConfiguration(value MTL4PipelineDataSetSerializerConfiguration) {
	objc.Send[struct{}](m.ID, objc.Sel("setConfiguration:"), value)
}
