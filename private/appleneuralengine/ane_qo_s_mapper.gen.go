// Code generated from Apple documentation for appleneuralengine. DO NOT EDIT.

package appleneuralengine

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [ANEQoSMapper] class.
var (
	_ANEQoSMapperClass     ANEQoSMapperClass
	_ANEQoSMapperClassOnce sync.Once
)

func getANEQoSMapperClass() ANEQoSMapperClass {
	_ANEQoSMapperClassOnce.Do(func() {
		_ANEQoSMapperClass = ANEQoSMapperClass{class: objc.GetClass("_ANEQoSMapper")}
	})
	return _ANEQoSMapperClass
}

// GetANEQoSMapperClass returns the class object for _ANEQoSMapper.
func GetANEQoSMapperClass() ANEQoSMapperClass {
	return getANEQoSMapperClass()
}

type ANEQoSMapperClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac ANEQoSMapperClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac ANEQoSMapperClass) Alloc() ANEQoSMapper {
	rv := objc.Send[ANEQoSMapper](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEQoSMapper
type ANEQoSMapper struct {
	objectivec.Object
}

// ANEQoSMapperFromID constructs a [ANEQoSMapper] from an objc.ID.
func ANEQoSMapperFromID(id objc.ID) ANEQoSMapper {
	return ANEQoSMapper{objectivec.Object{ID: id}}
}

// Ensure ANEQoSMapper implements IANEQoSMapper.
var _ IANEQoSMapper = ANEQoSMapper{}

// An interface definition for the [ANEQoSMapper] class.
//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEQoSMapper
type IANEQoSMapper interface {
	objectivec.IObject
}

// Init initializes the instance.
func (a ANEQoSMapper) Init() ANEQoSMapper {
	rv := objc.Send[ANEQoSMapper](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a ANEQoSMapper) Autorelease() ANEQoSMapper {
	rv := objc.Send[ANEQoSMapper](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewANEQoSMapper creates a new ANEQoSMapper instance.
func NewANEQoSMapper() ANEQoSMapper {
	class := getANEQoSMapperClass()
	rv := objc.Send[ANEQoSMapper](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEQoSMapper/aneBackgroundTaskQoS
func (_ANEQoSMapperClass ANEQoSMapperClass) AneBackgroundTaskQoS() uint32 {
	rv := objc.Send[uint32](objc.ID(_ANEQoSMapperClass.class), objc.Sel("aneBackgroundTaskQoS"))
	return rv
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEQoSMapper/aneDefaultTaskQoS
func (_ANEQoSMapperClass ANEQoSMapperClass) AneDefaultTaskQoS() uint32 {
	rv := objc.Send[uint32](objc.ID(_ANEQoSMapperClass.class), objc.Sel("aneDefaultTaskQoS"))
	return rv
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEQoSMapper/aneRealTimeTaskQoS
func (_ANEQoSMapperClass ANEQoSMapperClass) AneRealTimeTaskQoS() uint32 {
	rv := objc.Send[uint32](objc.ID(_ANEQoSMapperClass.class), objc.Sel("aneRealTimeTaskQoS"))
	return rv
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEQoSMapper/aneUserInitiatedTaskQoS
func (_ANEQoSMapperClass ANEQoSMapperClass) AneUserInitiatedTaskQoS() uint32 {
	rv := objc.Send[uint32](objc.ID(_ANEQoSMapperClass.class), objc.Sel("aneUserInitiatedTaskQoS"))
	return rv
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEQoSMapper/aneUserInteractiveTaskQoS
func (_ANEQoSMapperClass ANEQoSMapperClass) AneUserInteractiveTaskQoS() uint32 {
	rv := objc.Send[uint32](objc.ID(_ANEQoSMapperClass.class), objc.Sel("aneUserInteractiveTaskQoS"))
	return rv
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEQoSMapper/aneUtilityTaskQoS
func (_ANEQoSMapperClass ANEQoSMapperClass) AneUtilityTaskQoS() uint32 {
	rv := objc.Send[uint32](objc.ID(_ANEQoSMapperClass.class), objc.Sel("aneUtilityTaskQoS"))
	return rv
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEQoSMapper/dispatchQueueArrayByMappingPrioritiesWithTag:
func (_ANEQoSMapperClass ANEQoSMapperClass) DispatchQueueArrayByMappingPrioritiesWithTag(tag objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEQoSMapperClass.class), objc.Sel("dispatchQueueArrayByMappingPrioritiesWithTag:"), tag)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEQoSMapper/programPriorityForQoS:
func (_ANEQoSMapperClass ANEQoSMapperClass) ProgramPriorityForQoS(s uint32) int {
	rv := objc.Send[int](objc.ID(_ANEQoSMapperClass.class), objc.Sel("programPriorityForQoS:"), s)
	return rv
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEQoSMapper/qosForProgramPriority:
func (_ANEQoSMapperClass ANEQoSMapperClass) QosForProgramPriority(priority int) uint32 {
	rv := objc.Send[uint32](objc.ID(_ANEQoSMapperClass.class), objc.Sel("qosForProgramPriority:"), priority)
	return rv
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEQoSMapper/queueIndexForQoS:
func (_ANEQoSMapperClass ANEQoSMapperClass) QueueIndexForQoS(s uint32) uint64 {
	rv := objc.Send[uint64](objc.ID(_ANEQoSMapperClass.class), objc.Sel("queueIndexForQoS:"), s)
	return rv
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEQoSMapper/realTimeProgramPriority
func (_ANEQoSMapperClass ANEQoSMapperClass) RealTimeProgramPriority() int {
	rv := objc.Send[int](objc.ID(_ANEQoSMapperClass.class), objc.Sel("realTimeProgramPriority"))
	return rv
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEQoSMapper/realTimeQueueIndex
func (_ANEQoSMapperClass ANEQoSMapperClass) RealTimeQueueIndex() uint64 {
	rv := objc.Send[uint64](objc.ID(_ANEQoSMapperClass.class), objc.Sel("realTimeQueueIndex"))
	return rv
}
