// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLComputePlan] class.
var (
	_MLComputePlanClass     MLComputePlanClass
	_MLComputePlanClassOnce sync.Once
)

func getMLComputePlanClass() MLComputePlanClass {
	_MLComputePlanClassOnce.Do(func() {
		_MLComputePlanClass = MLComputePlanClass{class: objc.GetClass("MLComputePlan")}
	})
	return _MLComputePlanClass
}

// GetMLComputePlanClass returns the class object for MLComputePlan.
func GetMLComputePlanClass() MLComputePlanClass {
	return getMLComputePlanClass()
}

type MLComputePlanClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLComputePlanClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLComputePlanClass) Alloc() MLComputePlan {
	rv := objc.Send[MLComputePlan](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [MLComputePlan.ComputeDevicesBySupportedComputeUnits]
//   - [MLComputePlan.Configuration]
//   - [MLComputePlan.ExecutionScheduleByModelStructurePath]
//   - [MLComputePlan.ModelAssetStorageType]
//   - [MLComputePlan.ModelDescription]
//   - [MLComputePlan.InitWithModelStructureModelDescriptionModelAssetStorageTypeExecutionScheduleConfiguration]
// See: https://developer.apple.com/documentation/CoreML/MLComputePlan
type MLComputePlan struct {
	objectivec.Object
}

// MLComputePlanFromID constructs a [MLComputePlan] from an objc.ID.
func MLComputePlanFromID(id objc.ID) MLComputePlan {
	return MLComputePlan{objectivec.Object{ID: id}}
}
// Ensure MLComputePlan implements IMLComputePlan.
var _ IMLComputePlan = MLComputePlan{}

// An interface definition for the [MLComputePlan] class.
//
// # Methods
//
//   - [IMLComputePlan.ComputeDevicesBySupportedComputeUnits]
//   - [IMLComputePlan.Configuration]
//   - [IMLComputePlan.ExecutionScheduleByModelStructurePath]
//   - [IMLComputePlan.ModelAssetStorageType]
//   - [IMLComputePlan.ModelDescription]
//   - [IMLComputePlan.InitWithModelStructureModelDescriptionModelAssetStorageTypeExecutionScheduleConfiguration]
//
// See: https://developer.apple.com/documentation/CoreML/MLComputePlan
type IMLComputePlan interface {
	objectivec.IObject

	// Topic: Methods

	ComputeDevicesBySupportedComputeUnits() foundation.INSDictionary
	Configuration() IMLModelConfiguration
	ExecutionScheduleByModelStructurePath() foundation.INSDictionary
	ModelAssetStorageType() int64
	ModelDescription() IMLModelDescription
	InitWithModelStructureModelDescriptionModelAssetStorageTypeExecutionScheduleConfiguration(structure objectivec.IObject, description objectivec.IObject, type_ int64, schedule objectivec.IObject, configuration objectivec.IObject) MLComputePlan
}

// Init initializes the instance.
func (c MLComputePlan) Init() MLComputePlan {
	rv := objc.Send[MLComputePlan](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MLComputePlan) Autorelease() MLComputePlan {
	rv := objc.Send[MLComputePlan](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLComputePlan creates a new MLComputePlan instance.
func NewMLComputePlan() MLComputePlan {
	class := getMLComputePlanClass()
	rv := objc.Send[MLComputePlan](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/CoreML/MLComputePlan/initWithModelStructure:modelDescription:modelAssetStorageType:executionSchedule:configuration:
func NewComputePlanWithModelStructureModelDescriptionModelAssetStorageTypeExecutionScheduleConfiguration(structure objectivec.IObject, description objectivec.IObject, type_ int64, schedule objectivec.IObject, configuration objectivec.IObject) MLComputePlan {
	instance := getMLComputePlanClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithModelStructure:modelDescription:modelAssetStorageType:executionSchedule:configuration:"), structure, description, type_, schedule, configuration)
	return MLComputePlanFromID(rv)
}

//
// See: https://developer.apple.com/documentation/CoreML/MLComputePlan/initWithModelStructure:modelDescription:modelAssetStorageType:executionSchedule:configuration:
func (c MLComputePlan) InitWithModelStructureModelDescriptionModelAssetStorageTypeExecutionScheduleConfiguration(structure objectivec.IObject, description objectivec.IObject, type_ int64, schedule objectivec.IObject, configuration objectivec.IObject) MLComputePlan {
	rv := objc.Send[MLComputePlan](c.ID, objc.Sel("initWithModelStructure:modelDescription:modelAssetStorageType:executionSchedule:configuration:"), structure, description, type_, schedule, configuration)
	return rv
}

//
// See: https://developer.apple.com/documentation/CoreML/MLComputePlan/computePlanOfModelStructure:modelAsset:configuration:error:
func (_MLComputePlanClass MLComputePlanClass) ComputePlanOfModelStructureModelAssetConfigurationError(structure objectivec.IObject, asset objectivec.IObject, configuration objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLComputePlanClass.class), objc.Sel("computePlanOfModelStructure:modelAsset:configuration:error:"), structure, asset, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// See: https://developer.apple.com/documentation/CoreML/MLComputePlan/computeDevicesBySupportedComputeUnits
func (c MLComputePlan) ComputeDevicesBySupportedComputeUnits() foundation.INSDictionary {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("computeDevicesBySupportedComputeUnits"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/CoreML/MLComputePlan/configuration
func (c MLComputePlan) Configuration() IMLModelConfiguration {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("configuration"))
	return MLModelConfigurationFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/CoreML/MLComputePlan/executionScheduleByModelStructurePath
func (c MLComputePlan) ExecutionScheduleByModelStructurePath() foundation.INSDictionary {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("executionScheduleByModelStructurePath"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/CoreML/MLComputePlan/modelAssetStorageType
func (c MLComputePlan) ModelAssetStorageType() int64 {
	rv := objc.Send[int64](c.ID, objc.Sel("modelAssetStorageType"))
	return rv
}
// See: https://developer.apple.com/documentation/CoreML/MLComputePlan/modelDescription
func (c MLComputePlan) ModelDescription() IMLModelDescription {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("modelDescription"))
	return MLModelDescriptionFromID(objc.ID(rv))
}

