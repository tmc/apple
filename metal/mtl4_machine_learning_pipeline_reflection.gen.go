// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MTL4MachineLearningPipelineReflection] class.
var (
	_MTL4MachineLearningPipelineReflectionClass     MTL4MachineLearningPipelineReflectionClass
	_MTL4MachineLearningPipelineReflectionClassOnce sync.Once
)

func getMTL4MachineLearningPipelineReflectionClass() MTL4MachineLearningPipelineReflectionClass {
	_MTL4MachineLearningPipelineReflectionClassOnce.Do(func() {
		_MTL4MachineLearningPipelineReflectionClass = MTL4MachineLearningPipelineReflectionClass{class: objc.GetClass("MTL4MachineLearningPipelineReflection")}
	})
	return _MTL4MachineLearningPipelineReflectionClass
}

// GetMTL4MachineLearningPipelineReflectionClass returns the class object for MTL4MachineLearningPipelineReflection.
func GetMTL4MachineLearningPipelineReflectionClass() MTL4MachineLearningPipelineReflectionClass {
	return getMTL4MachineLearningPipelineReflectionClass()
}

type MTL4MachineLearningPipelineReflectionClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MTL4MachineLearningPipelineReflectionClass) Alloc() MTL4MachineLearningPipelineReflection {
	rv := objc.Send[MTL4MachineLearningPipelineReflection](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}







// Represents reflection information for a machine learning pipeline state.
//
// # Instance Properties
//
//   - [MTL4MachineLearningPipelineReflection.Bindings]: Describes every input and output of the pipeline.
//
// See: https://developer.apple.com/documentation/Metal/MTL4MachineLearningPipelineReflection
type MTL4MachineLearningPipelineReflection struct {
	objectivec.Object
}

// MTL4MachineLearningPipelineReflectionFromID constructs a [MTL4MachineLearningPipelineReflection] from an objc.ID.
//
// Represents reflection information for a machine learning pipeline state.
func MTL4MachineLearningPipelineReflectionFromID(id objc.ID) MTL4MachineLearningPipelineReflection {
	return MTL4MachineLearningPipelineReflection{objectivec.Object{id}}
}
// NOTE: MTL4MachineLearningPipelineReflection adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [MTL4MachineLearningPipelineReflection] class.
//
// # Instance Properties
//
//   - [IMTL4MachineLearningPipelineReflection.Bindings]: Describes every input and output of the pipeline.
//
// See: https://developer.apple.com/documentation/Metal/MTL4MachineLearningPipelineReflection
type IMTL4MachineLearningPipelineReflection interface {
	objectivec.IObject

	// Topic: Instance Properties

	// Describes every input and output of the pipeline.
	Bindings() []objectivec.IObject
}





// Init initializes the instance.
func (m MTL4MachineLearningPipelineReflection) Init() MTL4MachineLearningPipelineReflection {
	rv := objc.Send[MTL4MachineLearningPipelineReflection](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MTL4MachineLearningPipelineReflection) Autorelease() MTL4MachineLearningPipelineReflection {
	rv := objc.Send[MTL4MachineLearningPipelineReflection](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTL4MachineLearningPipelineReflection creates a new MTL4MachineLearningPipelineReflection instance.
func NewMTL4MachineLearningPipelineReflection() MTL4MachineLearningPipelineReflection {
	class := getMTL4MachineLearningPipelineReflectionClass()
	rv := objc.Send[MTL4MachineLearningPipelineReflection](objc.ID(class.class), objc.Sel("new"))
	return rv
}





















// Describes every input and output of the pipeline.
//
// See: https://developer.apple.com/documentation/Metal/MTL4MachineLearningPipelineReflection/bindings
func (m MTL4MachineLearningPipelineReflection) Bindings() []objectivec.IObject {
	rv := objc.Send[[]objc.ID](m.ID, objc.Sel("bindings"))
	return objc.ConvertSlice(rv, func(id objc.ID) objectivec.IObject {
		return objectivec.Object{ID: id}
	})
}

























