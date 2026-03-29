// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLLayerExecutionSchedule] class.
var (
	_MLLayerExecutionScheduleClass     MLLayerExecutionScheduleClass
	_MLLayerExecutionScheduleClassOnce sync.Once
)

func getMLLayerExecutionScheduleClass() MLLayerExecutionScheduleClass {
	_MLLayerExecutionScheduleClassOnce.Do(func() {
		_MLLayerExecutionScheduleClass = MLLayerExecutionScheduleClass{class: objc.GetClass("MLLayerExecutionSchedule")}
	})
	return _MLLayerExecutionScheduleClass
}

// GetMLLayerExecutionScheduleClass returns the class object for MLLayerExecutionSchedule.
func GetMLLayerExecutionScheduleClass() MLLayerExecutionScheduleClass {
	return getMLLayerExecutionScheduleClass()
}

type MLLayerExecutionScheduleClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLLayerExecutionScheduleClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLLayerExecutionScheduleClass) Alloc() MLLayerExecutionSchedule {
	rv := objc.Send[MLLayerExecutionSchedule](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [MLLayerExecutionSchedule.ComputeTimeRatio]
//   - [MLLayerExecutionSchedule.LayerError]
//   - [MLLayerExecutionSchedule.LayerIndex]
//   - [MLLayerExecutionSchedule.LayerName]
//   - [MLLayerExecutionSchedule.LayerTypeName]
//   - [MLLayerExecutionSchedule.PreferredComputeUnit]
//   - [MLLayerExecutionSchedule.SupportMessagesPerComputeUnit]
//   - [MLLayerExecutionSchedule.SupportedComputeUnits]
//   - [MLLayerExecutionSchedule.InitWithComputeUnitsLayerNameLayerErrorLayerTypeNameSupportedComputeUnitsLayerIndex]
//   - [MLLayerExecutionSchedule.InitWithComputeUnitsLayerNameLayerErrorLayerTypeNameSupportedComputeUnitsLayerIndexSupportMessagesComputeTimeRatio]
//   - [MLLayerExecutionSchedule.InitWithLayerError]
// See: https://developer.apple.com/documentation/CoreML/MLLayerExecutionSchedule
type MLLayerExecutionSchedule struct {
	objectivec.Object
}

// MLLayerExecutionScheduleFromID constructs a [MLLayerExecutionSchedule] from an objc.ID.
func MLLayerExecutionScheduleFromID(id objc.ID) MLLayerExecutionSchedule {
	return MLLayerExecutionSchedule{objectivec.Object{ID: id}}
}
// Ensure MLLayerExecutionSchedule implements IMLLayerExecutionSchedule.
var _ IMLLayerExecutionSchedule = MLLayerExecutionSchedule{}

// An interface definition for the [MLLayerExecutionSchedule] class.
//
// # Methods
//
//   - [IMLLayerExecutionSchedule.ComputeTimeRatio]
//   - [IMLLayerExecutionSchedule.LayerError]
//   - [IMLLayerExecutionSchedule.LayerIndex]
//   - [IMLLayerExecutionSchedule.LayerName]
//   - [IMLLayerExecutionSchedule.LayerTypeName]
//   - [IMLLayerExecutionSchedule.PreferredComputeUnit]
//   - [IMLLayerExecutionSchedule.SupportMessagesPerComputeUnit]
//   - [IMLLayerExecutionSchedule.SupportedComputeUnits]
//   - [IMLLayerExecutionSchedule.InitWithComputeUnitsLayerNameLayerErrorLayerTypeNameSupportedComputeUnitsLayerIndex]
//   - [IMLLayerExecutionSchedule.InitWithComputeUnitsLayerNameLayerErrorLayerTypeNameSupportedComputeUnitsLayerIndexSupportMessagesComputeTimeRatio]
//   - [IMLLayerExecutionSchedule.InitWithLayerError]
//
// See: https://developer.apple.com/documentation/CoreML/MLLayerExecutionSchedule
type IMLLayerExecutionSchedule interface {
	objectivec.IObject

	// Topic: Methods

	ComputeTimeRatio() float64
	LayerError() int64
	LayerIndex() int64
	LayerName() string
	LayerTypeName() string
	PreferredComputeUnit() uint64
	SupportMessagesPerComputeUnit() foundation.INSDictionary
	SupportedComputeUnits() uint64
	InitWithComputeUnitsLayerNameLayerErrorLayerTypeNameSupportedComputeUnitsLayerIndex(units uint64, name objectivec.IObject, error_ int64, name2 objectivec.IObject, units2 uint64, index int64) MLLayerExecutionSchedule
	InitWithComputeUnitsLayerNameLayerErrorLayerTypeNameSupportedComputeUnitsLayerIndexSupportMessagesComputeTimeRatio(units uint64, name objectivec.IObject, error_ int64, name2 objectivec.IObject, units2 uint64, index int64, messages objectivec.IObject, ratio float64) MLLayerExecutionSchedule
	InitWithLayerError(error_ int64) MLLayerExecutionSchedule
}

// Init initializes the instance.
func (l MLLayerExecutionSchedule) Init() MLLayerExecutionSchedule {
	rv := objc.Send[MLLayerExecutionSchedule](l.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (l MLLayerExecutionSchedule) Autorelease() MLLayerExecutionSchedule {
	rv := objc.Send[MLLayerExecutionSchedule](l.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLLayerExecutionSchedule creates a new MLLayerExecutionSchedule instance.
func NewMLLayerExecutionSchedule() MLLayerExecutionSchedule {
	class := getMLLayerExecutionScheduleClass()
	rv := objc.Send[MLLayerExecutionSchedule](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/CoreML/MLLayerExecutionSchedule/initWithComputeUnits:layerName:layerError:layerTypeName:supportedComputeUnits:layerIndex:
func NewLayerExecutionScheduleWithComputeUnitsLayerNameLayerErrorLayerTypeNameSupportedComputeUnitsLayerIndex(units uint64, name objectivec.IObject, error_ int64, name2 objectivec.IObject, units2 uint64, index int64) MLLayerExecutionSchedule {
	instance := getMLLayerExecutionScheduleClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithComputeUnits:layerName:layerError:layerTypeName:supportedComputeUnits:layerIndex:"), units, name, error_, name2, units2, index)
	return MLLayerExecutionScheduleFromID(rv)
}

//
// See: https://developer.apple.com/documentation/CoreML/MLLayerExecutionSchedule/initWithComputeUnits:layerName:layerError:layerTypeName:supportedComputeUnits:layerIndex:supportMessages:computeTimeRatio:
func NewLayerExecutionScheduleWithComputeUnitsLayerNameLayerErrorLayerTypeNameSupportedComputeUnitsLayerIndexSupportMessagesComputeTimeRatio(units uint64, name objectivec.IObject, error_ int64, name2 objectivec.IObject, units2 uint64, index int64, messages objectivec.IObject, ratio float64) MLLayerExecutionSchedule {
	instance := getMLLayerExecutionScheduleClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithComputeUnits:layerName:layerError:layerTypeName:supportedComputeUnits:layerIndex:supportMessages:computeTimeRatio:"), units, name, error_, name2, units2, index, messages, ratio)
	return MLLayerExecutionScheduleFromID(rv)
}

//
// See: https://developer.apple.com/documentation/CoreML/MLLayerExecutionSchedule/initWithLayerError:
func NewLayerExecutionScheduleWithLayerError(error_ int64) MLLayerExecutionSchedule {
	instance := getMLLayerExecutionScheduleClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithLayerError:"), error_)
	return MLLayerExecutionScheduleFromID(rv)
}

//
// See: https://developer.apple.com/documentation/CoreML/MLLayerExecutionSchedule/initWithComputeUnits:layerName:layerError:layerTypeName:supportedComputeUnits:layerIndex:
func (l MLLayerExecutionSchedule) InitWithComputeUnitsLayerNameLayerErrorLayerTypeNameSupportedComputeUnitsLayerIndex(units uint64, name objectivec.IObject, error_ int64, name2 objectivec.IObject, units2 uint64, index int64) MLLayerExecutionSchedule {
	rv := objc.Send[MLLayerExecutionSchedule](l.ID, objc.Sel("initWithComputeUnits:layerName:layerError:layerTypeName:supportedComputeUnits:layerIndex:"), units, name, error_, name2, units2, index)
	return rv
}
//
// See: https://developer.apple.com/documentation/CoreML/MLLayerExecutionSchedule/initWithComputeUnits:layerName:layerError:layerTypeName:supportedComputeUnits:layerIndex:supportMessages:computeTimeRatio:
func (l MLLayerExecutionSchedule) InitWithComputeUnitsLayerNameLayerErrorLayerTypeNameSupportedComputeUnitsLayerIndexSupportMessagesComputeTimeRatio(units uint64, name objectivec.IObject, error_ int64, name2 objectivec.IObject, units2 uint64, index int64, messages objectivec.IObject, ratio float64) MLLayerExecutionSchedule {
	rv := objc.Send[MLLayerExecutionSchedule](l.ID, objc.Sel("initWithComputeUnits:layerName:layerError:layerTypeName:supportedComputeUnits:layerIndex:supportMessages:computeTimeRatio:"), units, name, error_, name2, units2, index, messages, ratio)
	return rv
}
//
// See: https://developer.apple.com/documentation/CoreML/MLLayerExecutionSchedule/initWithLayerError:
func (l MLLayerExecutionSchedule) InitWithLayerError(error_ int64) MLLayerExecutionSchedule {
	rv := objc.Send[MLLayerExecutionSchedule](l.ID, objc.Sel("initWithLayerError:"), error_)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLLayerExecutionSchedule/computeTimeRatio
func (l MLLayerExecutionSchedule) ComputeTimeRatio() float64 {
	rv := objc.Send[float64](l.ID, objc.Sel("computeTimeRatio"))
	return rv
}
// See: https://developer.apple.com/documentation/CoreML/MLLayerExecutionSchedule/layerError
func (l MLLayerExecutionSchedule) LayerError() int64 {
	rv := objc.Send[int64](l.ID, objc.Sel("layerError"))
	return rv
}
// See: https://developer.apple.com/documentation/CoreML/MLLayerExecutionSchedule/layerIndex
func (l MLLayerExecutionSchedule) LayerIndex() int64 {
	rv := objc.Send[int64](l.ID, objc.Sel("layerIndex"))
	return rv
}
// See: https://developer.apple.com/documentation/CoreML/MLLayerExecutionSchedule/layerName
func (l MLLayerExecutionSchedule) LayerName() string {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("layerName"))
	return foundation.NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/CoreML/MLLayerExecutionSchedule/layerTypeName
func (l MLLayerExecutionSchedule) LayerTypeName() string {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("layerTypeName"))
	return foundation.NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/CoreML/MLLayerExecutionSchedule/preferredComputeUnit
func (l MLLayerExecutionSchedule) PreferredComputeUnit() uint64 {
	rv := objc.Send[uint64](l.ID, objc.Sel("preferredComputeUnit"))
	return rv
}
// See: https://developer.apple.com/documentation/CoreML/MLLayerExecutionSchedule/supportMessagesPerComputeUnit
func (l MLLayerExecutionSchedule) SupportMessagesPerComputeUnit() foundation.INSDictionary {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("supportMessagesPerComputeUnit"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/CoreML/MLLayerExecutionSchedule/supportedComputeUnits
func (l MLLayerExecutionSchedule) SupportedComputeUnits() uint64 {
	rv := objc.Send[uint64](l.ID, objc.Sel("supportedComputeUnits"))
	return rv
}

