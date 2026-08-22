// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
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
	rv := objc.SendIfResponds[MLComputePlan](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLComputePlan.ComputeDevicesBySupportedComputeUnits]
//   - [MLComputePlan.Configuration]
//   - [MLComputePlan.ExecutionScheduleByModelStructurePath]
//   - [MLComputePlan.ModelAssetStorageType]
//   - [MLComputePlan.ModelDescription]
//   - [MLComputePlan.InitWithModelStructureModelDescriptionModelAssetStorageTypeExecutionScheduleConfiguration]
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
func (m MLComputePlan) Init() MLComputePlan {
	rv := objc.SendIfResponds[MLComputePlan](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLComputePlan) Autorelease() MLComputePlan {
	rv := objc.SendIfResponds[MLComputePlan](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLComputePlan creates a new MLComputePlan instance.
func NewMLComputePlan() MLComputePlan {
	class := getMLComputePlanClass()
	rv := objc.SendIfResponds[MLComputePlan](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewComputePlanWithModelStructureModelDescriptionModelAssetStorageTypeExecutionScheduleConfiguration(structure objectivec.IObject, description objectivec.IObject, type_ int64, schedule objectivec.IObject, configuration objectivec.IObject) MLComputePlan {
	instance := getMLComputePlanClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithModelStructure:modelDescription:modelAssetStorageType:executionSchedule:configuration:"), structure, description, type_, schedule, configuration)
	return MLComputePlanFromID(rv)
}

func (m MLComputePlan) InitWithModelStructureModelDescriptionModelAssetStorageTypeExecutionScheduleConfiguration(structure objectivec.IObject, description objectivec.IObject, type_ int64, schedule objectivec.IObject, configuration objectivec.IObject) MLComputePlan {
	rv := objc.SendIfResponds[MLComputePlan](m.ID, objc.Sel("initWithModelStructure:modelDescription:modelAssetStorageType:executionSchedule:configuration:"), structure, description, type_, schedule, configuration)
	return rv
}

func (_MLComputePlanClass MLComputePlanClass) ComputePlanOfModelStructureModelAssetConfigurationError(structure objectivec.IObject, asset objectivec.IObject, configuration objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLComputePlanClass.class), objc.Sel("computePlanOfModelStructure:modelAsset:configuration:error:"), structure, asset, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

func (m MLComputePlan) ComputeDevicesBySupportedComputeUnits() foundation.INSDictionary {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("computeDevicesBySupportedComputeUnits"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (m MLComputePlan) Configuration() IMLModelConfiguration {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("configuration"))
	return MLModelConfigurationFromID(objc.ID(rv))
}
func (m MLComputePlan) ExecutionScheduleByModelStructurePath() foundation.INSDictionary {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("executionScheduleByModelStructurePath"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (m MLComputePlan) ModelAssetStorageType() int64 {
	rv := objc.SendIfResponds[int64](m.ID, objc.Sel("modelAssetStorageType"))
	return rv
}
func (m MLComputePlan) ModelDescription() IMLModelDescription {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("modelDescription"))
	return MLModelDescriptionFromID(objc.ID(rv))
}
