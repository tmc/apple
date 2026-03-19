// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MTL4CompilerDescriptor] class.
var (
	_MTL4CompilerDescriptorClass     MTL4CompilerDescriptorClass
	_MTL4CompilerDescriptorClassOnce sync.Once
)

func getMTL4CompilerDescriptorClass() MTL4CompilerDescriptorClass {
	_MTL4CompilerDescriptorClassOnce.Do(func() {
		_MTL4CompilerDescriptorClass = MTL4CompilerDescriptorClass{class: objc.GetClass("MTL4CompilerDescriptor")}
	})
	return _MTL4CompilerDescriptorClass
}

// GetMTL4CompilerDescriptorClass returns the class object for MTL4CompilerDescriptor.
func GetMTL4CompilerDescriptorClass() MTL4CompilerDescriptorClass {
	return getMTL4CompilerDescriptorClass()
}

type MTL4CompilerDescriptorClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MTL4CompilerDescriptorClass) Alloc() MTL4CompilerDescriptor {
	rv := objc.Send[MTL4CompilerDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// Groups together properties for creating a compiler context.
//
// # Instance Properties
//
//   - [MTL4CompilerDescriptor.Label]: Assigns an optional descriptor label to the compiler for debugging purposes.
//   - [MTL4CompilerDescriptor.SetLabel]
//   - [MTL4CompilerDescriptor.PipelineDataSetSerializer]: Assigns a pipeline data set serializer into which this compiler stores data for all pipelines it creates.
//   - [MTL4CompilerDescriptor.SetPipelineDataSetSerializer]
//
// See: https://developer.apple.com/documentation/Metal/MTL4CompilerDescriptor
type MTL4CompilerDescriptor struct {
	objectivec.Object
}

// MTL4CompilerDescriptorFromID constructs a [MTL4CompilerDescriptor] from an objc.ID.
//
// Groups together properties for creating a compiler context.
func MTL4CompilerDescriptorFromID(id objc.ID) MTL4CompilerDescriptor {
	return MTL4CompilerDescriptor{objectivec.Object{ID: id}}
}
// NOTE: MTL4CompilerDescriptor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MTL4CompilerDescriptor] class.
//
// # Instance Properties
//
//   - [IMTL4CompilerDescriptor.Label]: Assigns an optional descriptor label to the compiler for debugging purposes.
//   - [IMTL4CompilerDescriptor.SetLabel]
//   - [IMTL4CompilerDescriptor.PipelineDataSetSerializer]: Assigns a pipeline data set serializer into which this compiler stores data for all pipelines it creates.
//   - [IMTL4CompilerDescriptor.SetPipelineDataSetSerializer]
//
// See: https://developer.apple.com/documentation/Metal/MTL4CompilerDescriptor
type IMTL4CompilerDescriptor interface {
	objectivec.IObject

	// Topic: Instance Properties

	// Assigns an optional descriptor label to the compiler for debugging purposes.
	Label() string
	SetLabel(value string)
	// Assigns a pipeline data set serializer into which this compiler stores data for all pipelines it creates.
	PipelineDataSetSerializer() MTL4PipelineDataSetSerializer
	SetPipelineDataSetSerializer(value MTL4PipelineDataSetSerializer)
}

// Init initializes the instance.
func (m MTL4CompilerDescriptor) Init() MTL4CompilerDescriptor {
	rv := objc.Send[MTL4CompilerDescriptor](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MTL4CompilerDescriptor) Autorelease() MTL4CompilerDescriptor {
	rv := objc.Send[MTL4CompilerDescriptor](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTL4CompilerDescriptor creates a new MTL4CompilerDescriptor instance.
func NewMTL4CompilerDescriptor() MTL4CompilerDescriptor {
	class := getMTL4CompilerDescriptorClass()
	rv := objc.Send[MTL4CompilerDescriptor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Assigns an optional descriptor label to the compiler for debugging
// purposes.
//
// See: https://developer.apple.com/documentation/Metal/MTL4CompilerDescriptor/label
func (m MTL4CompilerDescriptor) Label() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("label"))
	return foundation.NSStringFromID(rv).String()
}
func (m MTL4CompilerDescriptor) SetLabel(value string) {
	objc.Send[struct{}](m.ID, objc.Sel("setLabel:"), objc.String(value))
}
// Assigns a pipeline data set serializer into which this compiler stores data
// for all pipelines it creates.
//
// See: https://developer.apple.com/documentation/Metal/MTL4CompilerDescriptor/pipelineDataSetSerializer
func (m MTL4CompilerDescriptor) PipelineDataSetSerializer() MTL4PipelineDataSetSerializer {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("pipelineDataSetSerializer"))
	return MTL4PipelineDataSetSerializerObjectFromID(rv)
}
func (m MTL4CompilerDescriptor) SetPipelineDataSetSerializer(value MTL4PipelineDataSetSerializer) {
	objc.Send[struct{}](m.ID, objc.Sel("setPipelineDataSetSerializer:"), value)
}

