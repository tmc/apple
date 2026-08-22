// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
	"github.com/tmc/apple/private/appleneuralengine"
)

// The class instance for the [ANECompilerAnalytics] class.
var (
	_ANECompilerAnalyticsClass     ANECompilerAnalyticsClass
	_ANECompilerAnalyticsClassOnce sync.Once
)

func getANECompilerAnalyticsClass() ANECompilerAnalyticsClass {
	_ANECompilerAnalyticsClassOnce.Do(func() {
		_ANECompilerAnalyticsClass = ANECompilerAnalyticsClass{class: objc.GetClass("_ANECompilerAnalytics")}
	})
	return _ANECompilerAnalyticsClass
}

// GetANECompilerAnalyticsClass returns the class object for _ANECompilerAnalytics.
func GetANECompilerAnalyticsClass() ANECompilerAnalyticsClass {
	return getANECompilerAnalyticsClass()
}

type ANECompilerAnalyticsClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac ANECompilerAnalyticsClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac ANECompilerAnalyticsClass) Alloc() ANECompilerAnalytics {
	rv := objc.SendIfResponds[ANECompilerAnalytics](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [ANECompilerAnalytics.AnalyticsBuffer]
//   - [ANECompilerAnalytics.BufferSizeInBytes]
//   - [ANECompilerAnalytics.DataInfoAt]
//   - [ANECompilerAnalytics.GetBOOLDataValueAt]
//   - [ANECompilerAnalytics.GetDataValueAt]
//   - [ANECompilerAnalytics.GroupInfoAt]
//   - [ANECompilerAnalytics.LayerInfoAt]
//   - [ANECompilerAnalytics.OffsetTableAtCount]
//   - [ANECompilerAnalytics.PopulateAnalytics]
//   - [ANECompilerAnalytics.ProcedureAnalytics]
//   - [ANECompilerAnalytics.SetProcedureAnalytics]
//   - [ANECompilerAnalytics.ProcedureInfoAt]
//   - [ANECompilerAnalytics.Serialize]
//   - [ANECompilerAnalytics.StringForAnalyticsType]
//   - [ANECompilerAnalytics.TaskInfoAt]
//   - [ANECompilerAnalytics.InitWithBuffer]
type ANECompilerAnalytics struct {
	objectivec.Object
}

// ANECompilerAnalyticsFromID constructs a [ANECompilerAnalytics] from an objc.ID.
func ANECompilerAnalyticsFromID(id objc.ID) ANECompilerAnalytics {
	return ANECompilerAnalytics{objectivec.Object{ID: id}}
}

// Ensure ANECompilerAnalytics implements IANECompilerAnalytics.
var _ IANECompilerAnalytics = ANECompilerAnalytics{}

// An interface definition for the [ANECompilerAnalytics] class.
//
// # Methods
//
//   - [IANECompilerAnalytics.AnalyticsBuffer]
//   - [IANECompilerAnalytics.BufferSizeInBytes]
//   - [IANECompilerAnalytics.DataInfoAt]
//   - [IANECompilerAnalytics.GetBOOLDataValueAt]
//   - [IANECompilerAnalytics.GetDataValueAt]
//   - [IANECompilerAnalytics.GroupInfoAt]
//   - [IANECompilerAnalytics.LayerInfoAt]
//   - [IANECompilerAnalytics.OffsetTableAtCount]
//   - [IANECompilerAnalytics.PopulateAnalytics]
//   - [IANECompilerAnalytics.ProcedureAnalytics]
//   - [IANECompilerAnalytics.SetProcedureAnalytics]
//   - [IANECompilerAnalytics.ProcedureInfoAt]
//   - [IANECompilerAnalytics.Serialize]
//   - [IANECompilerAnalytics.StringForAnalyticsType]
//   - [IANECompilerAnalytics.TaskInfoAt]
//   - [IANECompilerAnalytics.InitWithBuffer]
type IANECompilerAnalytics interface {
	objectivec.IObject

	// Topic: Methods

	AnalyticsBuffer() foundation.NSData
	BufferSizeInBytes() foundation.NSNumber
	DataInfoAt(at uint64) appleneuralengine.AnalyticsData
	GetBOOLDataValueAt(at uint64) bool
	GetDataValueAt(at uint64) uint64
	GroupInfoAt(at uint64) appleneuralengine.AnalyticsGroupInfo
	LayerInfoAt(at uint64) appleneuralengine.AnalyticsLayerInfo
	OffsetTableAtCount(at uint64, count uint32) unsafe.Pointer
	PopulateAnalytics() bool
	ProcedureAnalytics() foundation.INSArray
	SetProcedureAnalytics(value foundation.INSArray)
	ProcedureInfoAt(at uint64) appleneuralengine.AnalyticsProcedureInfo
	Serialize() objectivec.IObject
	StringForAnalyticsType(type_ uint32) objectivec.IObject
	TaskInfoAt(at uint64) appleneuralengine.AnalyticsTaskInfo
	InitWithBuffer(buffer objectivec.IObject) ANECompilerAnalytics
}

// Init initializes the instance.
func (a ANECompilerAnalytics) Init() ANECompilerAnalytics {
	rv := objc.SendIfResponds[ANECompilerAnalytics](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a ANECompilerAnalytics) Autorelease() ANECompilerAnalytics {
	rv := objc.SendIfResponds[ANECompilerAnalytics](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewANECompilerAnalytics creates a new ANECompilerAnalytics instance.
func NewANECompilerAnalytics() ANECompilerAnalytics {
	class := getANECompilerAnalyticsClass()
	rv := objc.SendIfResponds[ANECompilerAnalytics](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewANECompilerAnalyticsWithBuffer(buffer objectivec.IObject) ANECompilerAnalytics {
	instance := getANECompilerAnalyticsClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithBuffer:"), buffer)
	return ANECompilerAnalyticsFromID(rv)
}

func (a ANECompilerAnalytics) DataInfoAt(at uint64) appleneuralengine.AnalyticsData {
	rv := objc.SendIfResponds[appleneuralengine.AnalyticsData](a.ID, objc.Sel("dataInfoAt:"), at)
	return appleneuralengine.AnalyticsData(rv)
}
func (a ANECompilerAnalytics) GetBOOLDataValueAt(at uint64) bool {
	rv := objc.SendIfResponds[bool](a.ID, objc.Sel("getBOOLDataValueAt:"), at)
	return rv
}
func (a ANECompilerAnalytics) GetDataValueAt(at uint64) uint64 {
	rv := objc.SendIfResponds[uint64](a.ID, objc.Sel("getDataValueAt:"), at)
	return rv
}
func (a ANECompilerAnalytics) GroupInfoAt(at uint64) appleneuralengine.AnalyticsGroupInfo {
	rv := objc.SendIfResponds[appleneuralengine.AnalyticsGroupInfo](a.ID, objc.Sel("groupInfoAt:"), at)
	return appleneuralengine.AnalyticsGroupInfo(rv)
}
func (a ANECompilerAnalytics) LayerInfoAt(at uint64) appleneuralengine.AnalyticsLayerInfo {
	rv := objc.SendIfResponds[appleneuralengine.AnalyticsLayerInfo](a.ID, objc.Sel("layerInfoAt:"), at)
	return appleneuralengine.AnalyticsLayerInfo(rv)
}
func (a ANECompilerAnalytics) OffsetTableAtCount(at uint64, count uint32) unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](a.ID, objc.Sel("offsetTableAt:count:"), at, count)
	return rv
}
func (a ANECompilerAnalytics) PopulateAnalytics() bool {
	rv := objc.SendIfResponds[bool](a.ID, objc.Sel("populateAnalytics"))
	return rv
}
func (a ANECompilerAnalytics) ProcedureInfoAt(at uint64) appleneuralengine.AnalyticsProcedureInfo {
	rv := objc.SendIfResponds[appleneuralengine.AnalyticsProcedureInfo](a.ID, objc.Sel("procedureInfoAt:"), at)
	return appleneuralengine.AnalyticsProcedureInfo(rv)
}
func (a ANECompilerAnalytics) Serialize() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](a.ID, objc.Sel("serialize"))
	return objectivec.Object{ID: rv}
}
func (a ANECompilerAnalytics) StringForAnalyticsType(type_ uint32) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](a.ID, objc.Sel("stringForAnalyticsType:"), type_)
	return objectivec.Object{ID: rv}
}
func (a ANECompilerAnalytics) TaskInfoAt(at uint64) appleneuralengine.AnalyticsTaskInfo {
	rv := objc.SendIfResponds[appleneuralengine.AnalyticsTaskInfo](a.ID, objc.Sel("taskInfoAt:"), at)
	return appleneuralengine.AnalyticsTaskInfo(rv)
}
func (a ANECompilerAnalytics) InitWithBuffer(buffer objectivec.IObject) ANECompilerAnalytics {
	rv := objc.SendIfResponds[ANECompilerAnalytics](a.ID, objc.Sel("initWithBuffer:"), buffer)
	return rv
}

func (_ANECompilerAnalyticsClass ANECompilerAnalyticsClass) ObjectWithBuffer(buffer objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_ANECompilerAnalyticsClass.class), objc.Sel("objectWithBuffer:"), buffer)
	return objectivec.Object{ID: rv}
}

func (a ANECompilerAnalytics) AnalyticsBuffer() foundation.NSData {
	rv := objc.SendIfResponds[objc.ID](a.ID, objc.Sel("analyticsBuffer"))
	return foundation.NSDataFromID(objc.ID(rv))
}
func (a ANECompilerAnalytics) BufferSizeInBytes() foundation.NSNumber {
	rv := objc.SendIfResponds[objc.ID](a.ID, objc.Sel("bufferSizeInBytes"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
func (a ANECompilerAnalytics) ProcedureAnalytics() foundation.INSArray {
	rv := objc.SendIfResponds[objc.ID](a.ID, objc.Sel("procedureAnalytics"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (a ANECompilerAnalytics) SetProcedureAnalytics(value foundation.INSArray) {
	objc.SendIfResponds[struct{}](a.ID, objc.Sel("setProcedureAnalytics:"), value)
}
