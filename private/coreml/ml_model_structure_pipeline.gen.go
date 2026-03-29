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

// Class returns the underlying Objective-C class pointer.
func (mc MLModelStructurePipelineClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLModelStructurePipelineClass) Alloc() MLModelStructurePipeline {
	rv := objc.Send[MLModelStructurePipeline](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [MLModelStructurePipeline.InitWithSubModelNamesSubModels]
// See: https://developer.apple.com/documentation/CoreML/MLModelStructurePipeline
type MLModelStructurePipeline struct {
	objectivec.Object
}

// MLModelStructurePipelineFromID constructs a [MLModelStructurePipeline] from an objc.ID.
func MLModelStructurePipelineFromID(id objc.ID) MLModelStructurePipeline {
	return MLModelStructurePipeline{objectivec.Object{ID: id}}
}
// Ensure MLModelStructurePipeline implements IMLModelStructurePipeline.
var _ IMLModelStructurePipeline = MLModelStructurePipeline{}

// An interface definition for the [MLModelStructurePipeline] class.
//
// # Methods
//
//   - [IMLModelStructurePipeline.InitWithSubModelNamesSubModels]
//
// See: https://developer.apple.com/documentation/CoreML/MLModelStructurePipeline
type IMLModelStructurePipeline interface {
	objectivec.IObject

	// Topic: Methods

	InitWithSubModelNamesSubModels(names objectivec.IObject, models objectivec.IObject) MLModelStructurePipeline
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

//
// See: https://developer.apple.com/documentation/CoreML/MLModelStructurePipeline/initWithSubModelNames:subModels:
func NewModelStructurePipelineWithSubModelNamesSubModels(names objectivec.IObject, models objectivec.IObject) MLModelStructurePipeline {
	instance := getMLModelStructurePipelineClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSubModelNames:subModels:"), names, models)
	return MLModelStructurePipelineFromID(rv)
}

//
// See: https://developer.apple.com/documentation/CoreML/MLModelStructurePipeline/initWithSubModelNames:subModels:
func (m MLModelStructurePipeline) InitWithSubModelNamesSubModels(names objectivec.IObject, models objectivec.IObject) MLModelStructurePipeline {
	rv := objc.Send[MLModelStructurePipeline](m.ID, objc.Sel("initWithSubModelNames:subModels:"), names, models)
	return rv
}

