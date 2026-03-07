// Code generated from Apple documentation for appleneuralengine. DO NOT EDIT.

package appleneuralengine

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
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

// Alloc allocates memory for a new instance of the class.
func (ac ANECompilerAnalyticsClass) Alloc() ANECompilerAnalytics {
	rv := objc.Send[ANECompilerAnalytics](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}







//
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
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANECompilerAnalytics
type ANECompilerAnalytics struct {
	objectivec.Object
}

// ANECompilerAnalyticsFromID constructs a [ANECompilerAnalytics] from an objc.ID.
func ANECompilerAnalyticsFromID(id objc.ID) ANECompilerAnalytics {
	return ANECompilerAnalytics{objectivec.Object{id}}
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
//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANECompilerAnalytics
type IANECompilerAnalytics interface {
	objectivec.IObject

	// Topic: Methods

	AnalyticsBuffer() foundation.INSData
	BufferSizeInBytes() foundation.NSNumber
	DataInfoAt(at uint64) AnalyticsDataRef
	GetBOOLDataValueAt(at uint64) bool
	GetDataValueAt(at uint64) uint64
	GroupInfoAt(at uint64) AnalyticsGroupInfoRef
	LayerInfoAt(at uint64) AnalyticsLayerInfoRef
	OffsetTableAtCount(at uint64, count uint32) AnalyticsOffsetTableRef
	PopulateAnalytics() bool
	ProcedureAnalytics() foundation.INSArray
	SetProcedureAnalytics(value foundation.INSArray)
	ProcedureInfoAt(at uint64) AnalyticsProcedureInfoRef
	Serialize() objectivec.IObject
	StringForAnalyticsType(type_ uint32) objectivec.IObject
	TaskInfoAt(at uint64) AnalyticsTaskInfoRef
	InitWithBuffer(buffer objectivec.IObject) ANECompilerAnalytics
}





// Init initializes the instance.
func (a ANECompilerAnalytics) Init() ANECompilerAnalytics {
	rv := objc.Send[ANECompilerAnalytics](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a ANECompilerAnalytics) Autorelease() ANECompilerAnalytics {
	rv := objc.Send[ANECompilerAnalytics](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewANECompilerAnalytics creates a new ANECompilerAnalytics instance.
func NewANECompilerAnalytics() ANECompilerAnalytics {
	class := getANECompilerAnalyticsClass()
	rv := objc.Send[ANECompilerAnalytics](objc.ID(class.class), objc.Sel("new"))
	return rv
}






//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANECompilerAnalytics/initWithBuffer:
func NewANECompilerAnalyticsWithBuffer(buffer objectivec.IObject) ANECompilerAnalytics {
	instance := getANECompilerAnalyticsClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithBuffer:"), buffer)
	return ANECompilerAnalyticsFromID(rv)
}







//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANECompilerAnalytics/dataInfoAt:
func (a ANECompilerAnalytics) DataInfoAt(at uint64) AnalyticsDataRef {
	rv := objc.Send[AnalyticsDataRef](a.ID, objc.Sel("dataInfoAt:"), at)
	return AnalyticsDataRef(rv)
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANECompilerAnalytics/getBOOLDataValueAt:
func (a ANECompilerAnalytics) GetBOOLDataValueAt(at uint64) bool {
	rv := objc.Send[bool](a.ID, objc.Sel("getBOOLDataValueAt:"), at)
	return rv
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANECompilerAnalytics/getDataValueAt:
func (a ANECompilerAnalytics) GetDataValueAt(at uint64) uint64 {
	rv := objc.Send[uint64](a.ID, objc.Sel("getDataValueAt:"), at)
	return rv
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANECompilerAnalytics/groupInfoAt:
func (a ANECompilerAnalytics) GroupInfoAt(at uint64) AnalyticsGroupInfoRef {
	rv := objc.Send[AnalyticsGroupInfoRef](a.ID, objc.Sel("groupInfoAt:"), at)
	return AnalyticsGroupInfoRef(rv)
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANECompilerAnalytics/layerInfoAt:
func (a ANECompilerAnalytics) LayerInfoAt(at uint64) AnalyticsLayerInfoRef {
	rv := objc.Send[AnalyticsLayerInfoRef](a.ID, objc.Sel("layerInfoAt:"), at)
	return AnalyticsLayerInfoRef(rv)
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANECompilerAnalytics/offsetTableAt:count:
func (a ANECompilerAnalytics) OffsetTableAtCount(at uint64, count uint32) AnalyticsOffsetTableRef {
	rv := objc.Send[AnalyticsOffsetTableRef](a.ID, objc.Sel("offsetTableAt:count:"), at, count)
	return AnalyticsOffsetTableRef(rv)
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANECompilerAnalytics/populateAnalytics
func (a ANECompilerAnalytics) PopulateAnalytics() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("populateAnalytics"))
	return rv
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANECompilerAnalytics/procedureInfoAt:
func (a ANECompilerAnalytics) ProcedureInfoAt(at uint64) AnalyticsProcedureInfoRef {
	rv := objc.Send[AnalyticsProcedureInfoRef](a.ID, objc.Sel("procedureInfoAt:"), at)
	return AnalyticsProcedureInfoRef(rv)
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANECompilerAnalytics/serialize
func (a ANECompilerAnalytics) Serialize() objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("serialize"))
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANECompilerAnalytics/stringForAnalyticsType:
func (a ANECompilerAnalytics) StringForAnalyticsType(type_ uint32) objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("stringForAnalyticsType:"), type_)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANECompilerAnalytics/taskInfoAt:
func (a ANECompilerAnalytics) TaskInfoAt(at uint64) AnalyticsTaskInfoRef {
	rv := objc.Send[AnalyticsTaskInfoRef](a.ID, objc.Sel("taskInfoAt:"), at)
	return AnalyticsTaskInfoRef(rv)
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANECompilerAnalytics/initWithBuffer:
func (a ANECompilerAnalytics) InitWithBuffer(buffer objectivec.IObject) ANECompilerAnalytics {
	rv := objc.Send[ANECompilerAnalytics](a.ID, objc.Sel("initWithBuffer:"), buffer)
	return rv
}





//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANECompilerAnalytics/objectWithBuffer:
func (_ANECompilerAnalyticsClass ANECompilerAnalyticsClass) ObjectWithBuffer(buffer objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANECompilerAnalyticsClass.class), objc.Sel("objectWithBuffer:"), buffer)
	return objectivec.Object{ID: rv}
}








// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANECompilerAnalytics/analyticsBuffer
func (a ANECompilerAnalytics) AnalyticsBuffer() foundation.INSData {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("analyticsBuffer"))
	return foundation.NSDataFromID(objc.ID(rv))
}



// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANECompilerAnalytics/bufferSizeInBytes
func (a ANECompilerAnalytics) BufferSizeInBytes() foundation.NSNumber {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("bufferSizeInBytes"))
	return foundation.NSNumberFromID(objc.ID(rv))
}



// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANECompilerAnalytics/procedureAnalytics
func (a ANECompilerAnalytics) ProcedureAnalytics() foundation.INSArray {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("procedureAnalytics"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (a ANECompilerAnalytics) SetProcedureAnalytics(value foundation.INSArray) {
	objc.Send[struct{}](a.ID, objc.Sel("setProcedureAnalytics:"), value)
}

















