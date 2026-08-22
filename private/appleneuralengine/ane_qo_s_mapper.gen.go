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
	rv := objc.SendIfResponds[ANEQoSMapper](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

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
type IANEQoSMapper interface {
	objectivec.IObject
}

// Init initializes the instance.
func (a ANEQoSMapper) Init() ANEQoSMapper {
	rv := objc.SendIfResponds[ANEQoSMapper](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a ANEQoSMapper) Autorelease() ANEQoSMapper {
	rv := objc.SendIfResponds[ANEQoSMapper](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewANEQoSMapper creates a new ANEQoSMapper instance.
func NewANEQoSMapper() ANEQoSMapper {
	class := getANEQoSMapperClass()
	rv := objc.SendIfResponds[ANEQoSMapper](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (_ANEQoSMapperClass ANEQoSMapperClass) AneBackgroundTaskQoS() uint32 {
	rv := objc.SendIfResponds[uint32](objc.ID(_ANEQoSMapperClass.class), objc.Sel("aneBackgroundTaskQoS"))
	return rv
}
func (_ANEQoSMapperClass ANEQoSMapperClass) AneDefaultTaskQoS() uint32 {
	rv := objc.SendIfResponds[uint32](objc.ID(_ANEQoSMapperClass.class), objc.Sel("aneDefaultTaskQoS"))
	return rv
}
func (_ANEQoSMapperClass ANEQoSMapperClass) AneRealTimeTaskQoS() uint32 {
	rv := objc.SendIfResponds[uint32](objc.ID(_ANEQoSMapperClass.class), objc.Sel("aneRealTimeTaskQoS"))
	return rv
}
func (_ANEQoSMapperClass ANEQoSMapperClass) AneUserInitiatedTaskQoS() uint32 {
	rv := objc.SendIfResponds[uint32](objc.ID(_ANEQoSMapperClass.class), objc.Sel("aneUserInitiatedTaskQoS"))
	return rv
}
func (_ANEQoSMapperClass ANEQoSMapperClass) AneUserInteractiveTaskQoS() uint32 {
	rv := objc.SendIfResponds[uint32](objc.ID(_ANEQoSMapperClass.class), objc.Sel("aneUserInteractiveTaskQoS"))
	return rv
}
func (_ANEQoSMapperClass ANEQoSMapperClass) AneUtilityTaskQoS() uint32 {
	rv := objc.SendIfResponds[uint32](objc.ID(_ANEQoSMapperClass.class), objc.Sel("aneUtilityTaskQoS"))
	return rv
}
func (_ANEQoSMapperClass ANEQoSMapperClass) DispatchQueueArrayByMappingPrioritiesWithTag(tag objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_ANEQoSMapperClass.class), objc.Sel("dispatchQueueArrayByMappingPrioritiesWithTag:"), tag)
	return objectivec.Object{ID: rv}
}
func (_ANEQoSMapperClass ANEQoSMapperClass) ProgramPriorityForQoS(s uint32) int {
	rv := objc.SendIfResponds[int](objc.ID(_ANEQoSMapperClass.class), objc.Sel("programPriorityForQoS:"), s)
	return rv
}
func (_ANEQoSMapperClass ANEQoSMapperClass) QosForProgramPriority(priority int) uint32 {
	rv := objc.SendIfResponds[uint32](objc.ID(_ANEQoSMapperClass.class), objc.Sel("qosForProgramPriority:"), priority)
	return rv
}
func (_ANEQoSMapperClass ANEQoSMapperClass) QueueIndexForQoS(s uint32) uint64 {
	rv := objc.SendIfResponds[uint64](objc.ID(_ANEQoSMapperClass.class), objc.Sel("queueIndexForQoS:"), s)
	return rv
}
func (_ANEQoSMapperClass ANEQoSMapperClass) RealTimeProgramPriority() int {
	rv := objc.SendIfResponds[int](objc.ID(_ANEQoSMapperClass.class), objc.Sel("realTimeProgramPriority"))
	return rv
}
func (_ANEQoSMapperClass ANEQoSMapperClass) RealTimeQueueIndex() uint64 {
	rv := objc.SendIfResponds[uint64](objc.ID(_ANEQoSMapperClass.class), objc.Sel("realTimeQueueIndex"))
	return rv
}
