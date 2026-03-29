// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLPipeline] class.
var (
	_MLPipelineClass     MLPipelineClass
	_MLPipelineClassOnce sync.Once
)

func getMLPipelineClass() MLPipelineClass {
	_MLPipelineClassOnce.Do(func() {
		_MLPipelineClass = MLPipelineClass{class: objc.GetClass("MLPipeline")}
	})
	return _MLPipelineClass
}

// GetMLPipelineClass returns the class object for MLPipeline.
func GetMLPipelineClass() MLPipelineClass {
	return getMLPipelineClass()
}

type MLPipelineClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLPipelineClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLPipelineClass) Alloc() MLPipeline {
	rv := objc.Send[MLPipeline](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A parent class referenced by other CoreML classes. [Full Topic]
type MLPipeline struct {
	objectivec.Object
}

// MLPipelineFromID constructs a [MLPipeline] from an objc.ID.
//
// A parent class referenced by other CoreML classes.
func MLPipelineFromID(id objc.ID) MLPipeline {
	return MLPipeline{objectivec.Object{ID: id}}
}
// Ensure MLPipeline implements IMLPipeline.
var _ IMLPipeline = MLPipeline{}

// An interface definition for the [MLPipeline] class.
type IMLPipeline interface {
	objectivec.IObject
}

// Init initializes the instance.
func (p MLPipeline) Init() MLPipeline {
	rv := objc.Send[MLPipeline](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p MLPipeline) Autorelease() MLPipeline {
	rv := objc.Send[MLPipeline](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLPipeline creates a new MLPipeline instance.
func NewMLPipeline() MLPipeline {
	class := getMLPipelineClass()
	rv := objc.Send[MLPipeline](objc.ID(class.class), objc.Sel("new"))
	return rv
}

