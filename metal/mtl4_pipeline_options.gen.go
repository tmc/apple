// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MTL4PipelineOptions] class.
var (
	_MTL4PipelineOptionsClass     MTL4PipelineOptionsClass
	_MTL4PipelineOptionsClassOnce sync.Once
)

func getMTL4PipelineOptionsClass() MTL4PipelineOptionsClass {
	_MTL4PipelineOptionsClassOnce.Do(func() {
		_MTL4PipelineOptionsClass = MTL4PipelineOptionsClass{class: objc.GetClass("MTL4PipelineOptions")}
	})
	return _MTL4PipelineOptionsClass
}

// GetMTL4PipelineOptionsClass returns the class object for MTL4PipelineOptions.
func GetMTL4PipelineOptionsClass() MTL4PipelineOptionsClass {
	return getMTL4PipelineOptionsClass()
}

type MTL4PipelineOptionsClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MTL4PipelineOptionsClass) Alloc() MTL4PipelineOptions {
	rv := objc.Send[MTL4PipelineOptions](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}







// Provides options controlling how to compile a pipeline state.
//
// # Overview
// 
// You provide these options through the [MTL4PipelineDescriptor] class at
// compilation time.
//
// # Instance Properties
//
//   - [MTL4PipelineOptions.ShaderReflection]: Controls whether to include Metal shader reflection in this pipeline.
//   - [MTL4PipelineOptions.SetShaderReflection]
//   - [MTL4PipelineOptions.ShaderValidation]: Controls whether to enable or disable Metal Shader Validation for the pipeline.
//   - [MTL4PipelineOptions.SetShaderValidation]
//
// See: https://developer.apple.com/documentation/Metal/MTL4PipelineOptions
type MTL4PipelineOptions struct {
	objectivec.Object
}

// MTL4PipelineOptionsFromID constructs a [MTL4PipelineOptions] from an objc.ID.
//
// Provides options controlling how to compile a pipeline state.
func MTL4PipelineOptionsFromID(id objc.ID) MTL4PipelineOptions {
	return MTL4PipelineOptions{objectivec.Object{id}}
}
// NOTE: MTL4PipelineOptions adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [MTL4PipelineOptions] class.
//
// # Instance Properties
//
//   - [IMTL4PipelineOptions.ShaderReflection]: Controls whether to include Metal shader reflection in this pipeline.
//   - [IMTL4PipelineOptions.SetShaderReflection]
//   - [IMTL4PipelineOptions.ShaderValidation]: Controls whether to enable or disable Metal Shader Validation for the pipeline.
//   - [IMTL4PipelineOptions.SetShaderValidation]
//
// See: https://developer.apple.com/documentation/Metal/MTL4PipelineOptions
type IMTL4PipelineOptions interface {
	objectivec.IObject

	// Topic: Instance Properties

	// Controls whether to include Metal shader reflection in this pipeline.
	ShaderReflection() MTL4ShaderReflection
	SetShaderReflection(value MTL4ShaderReflection)
	// Controls whether to enable or disable Metal Shader Validation for the pipeline.
	ShaderValidation() MTLShaderValidation
	SetShaderValidation(value MTLShaderValidation)
}





// Init initializes the instance.
func (m MTL4PipelineOptions) Init() MTL4PipelineOptions {
	rv := objc.Send[MTL4PipelineOptions](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MTL4PipelineOptions) Autorelease() MTL4PipelineOptions {
	rv := objc.Send[MTL4PipelineOptions](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTL4PipelineOptions creates a new MTL4PipelineOptions instance.
func NewMTL4PipelineOptions() MTL4PipelineOptions {
	class := getMTL4PipelineOptionsClass()
	rv := objc.Send[MTL4PipelineOptions](objc.ID(class.class), objc.Sel("new"))
	return rv
}





















// Controls whether to include Metal shader reflection in this pipeline.
//
// See: https://developer.apple.com/documentation/Metal/MTL4PipelineOptions/shaderReflection
func (m MTL4PipelineOptions) ShaderReflection() MTL4ShaderReflection {
	rv := objc.Send[MTL4ShaderReflection](m.ID, objc.Sel("shaderReflection"))
	return MTL4ShaderReflection(rv)
}
func (m MTL4PipelineOptions) SetShaderReflection(value MTL4ShaderReflection) {
	objc.Send[struct{}](m.ID, objc.Sel("setShaderReflection:"), value)
}



// Controls whether to enable or disable Metal Shader Validation for the
// pipeline.
//
// See: https://developer.apple.com/documentation/Metal/MTL4PipelineOptions/shaderValidation
func (m MTL4PipelineOptions) ShaderValidation() MTLShaderValidation {
	rv := objc.Send[MTLShaderValidation](m.ID, objc.Sel("shaderValidation"))
	return MTLShaderValidation(rv)
}
func (m MTL4PipelineOptions) SetShaderValidation(value MTLShaderValidation) {
	objc.Send[struct{}](m.ID, objc.Sel("setShaderValidation:"), value)
}
























