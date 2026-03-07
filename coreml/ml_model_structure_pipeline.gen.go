// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLModelStructurePipeline] class.
var (
	_MLModelStructurePipelineClass     MLModelStructurePipelineClass
	_MLModelStructurePipelineClassOnce sync.Once
)

func getMLModelStructurePipelineClass() MLModelStructurePipelineClass {
	_MLModelStructurePipelineClassOnce.Do(func() {
		_MLModelStructurePipelineClass = MLModelStructurePipelineClass{class: objc.GetClass("MLModelStructurePipeline")}
	})
	return _MLModelStructurePipelineClass
}

// GetMLModelStructurePipelineClass returns the class object for MLModelStructurePipeline.
func GetMLModelStructurePipelineClass() MLModelStructurePipelineClass {
	return getMLModelStructurePipelineClass()
}

type MLModelStructurePipelineClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLModelStructurePipelineClass) Alloc() MLModelStructurePipeline {
	rv := objc.Send[MLModelStructurePipeline](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}







// A class representing the structure of a Pipeline model.
//
// # Accessing the models and model names
//
//   - [MLModelStructurePipeline.SubModelNames]: The names of the sub models in the pipeline.
//   - [MLModelStructurePipeline.SubModels]: The structure of the sub models in the pipeline.
//
// See: https://developer.apple.com/documentation/CoreML/MLModelStructurePipeline
type MLModelStructurePipeline struct {
	objectivec.Object
}

// MLModelStructurePipelineFromID constructs a [MLModelStructurePipeline] from an objc.ID.
//
// A class representing the structure of a Pipeline model.
func MLModelStructurePipelineFromID(id objc.ID) MLModelStructurePipeline {
	return MLModelStructurePipeline{objectivec.Object{id}}
}
// Ensure MLModelStructurePipeline implements IMLModelStructurePipeline.
var _ IMLModelStructurePipeline = MLModelStructurePipeline{}





// An interface definition for the [MLModelStructurePipeline] class.
//
// # Accessing the models and model names
//
//   - [IMLModelStructurePipeline.SubModelNames]: The names of the sub models in the pipeline.
//   - [IMLModelStructurePipeline.SubModels]: The structure of the sub models in the pipeline.
//
// See: https://developer.apple.com/documentation/CoreML/MLModelStructurePipeline
type IMLModelStructurePipeline interface {
	objectivec.IObject

	// Topic: Accessing the models and model names

	// The names of the sub models in the pipeline.
	SubModelNames() []string
	// The structure of the sub models in the pipeline.
	SubModels() []MLModelStructure
}





// Init initializes the instance.
func (m MLModelStructurePipeline) Init() MLModelStructurePipeline {
	rv := objc.Send[MLModelStructurePipeline](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLModelStructurePipeline) Autorelease() MLModelStructurePipeline {
	rv := objc.Send[MLModelStructurePipeline](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLModelStructurePipeline creates a new MLModelStructurePipeline instance.
func NewMLModelStructurePipeline() MLModelStructurePipeline {
	class := getMLModelStructurePipelineClass()
	rv := objc.Send[MLModelStructurePipeline](objc.ID(class.class), objc.Sel("new"))
	return rv
}





















// The names of the sub models in the pipeline.
//
// See: https://developer.apple.com/documentation/CoreML/MLModelStructurePipeline/subModelNames
func (m MLModelStructurePipeline) SubModelNames() []string {
	rv := objc.Send[[]objc.ID](m.ID, objc.Sel("subModelNames"))
	return objc.ConvertSliceToStrings(rv)
}



// The structure of the sub models in the pipeline.
//
// See: https://developer.apple.com/documentation/CoreML/MLModelStructurePipeline/subModels
func (m MLModelStructurePipeline) SubModels() []MLModelStructure {
	rv := objc.Send[[]objc.ID](m.ID, objc.Sel("subModels"))
	return objc.ConvertSlice(rv, func(id objc.ID) MLModelStructure {
		return MLModelStructureFromID(id)
	})
}

















