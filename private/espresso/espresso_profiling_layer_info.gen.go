// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [EspressoProfilingLayerInfo] class.
var (
	_EspressoProfilingLayerInfoClass     EspressoProfilingLayerInfoClass
	_EspressoProfilingLayerInfoClassOnce sync.Once
)

func getEspressoProfilingLayerInfoClass() EspressoProfilingLayerInfoClass {
	_EspressoProfilingLayerInfoClassOnce.Do(func() {
		_EspressoProfilingLayerInfoClass = EspressoProfilingLayerInfoClass{class: objc.GetClass("EspressoProfilingLayerInfo")}
	})
	return _EspressoProfilingLayerInfoClass
}

// GetEspressoProfilingLayerInfoClass returns the class object for EspressoProfilingLayerInfo.
func GetEspressoProfilingLayerInfoClass() EspressoProfilingLayerInfoClass {
	return getEspressoProfilingLayerInfoClass()
}

type EspressoProfilingLayerInfoClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (ec EspressoProfilingLayerInfoClass) Alloc() EspressoProfilingLayerInfo {
	rv := objc.Send[EspressoProfilingLayerInfo](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}







//
// # Methods
//
//   - [EspressoProfilingLayerInfo.Average_runtime]
//   - [EspressoProfilingLayerInfo.SetAverage_runtime]
//   - [EspressoProfilingLayerInfo.Debug_name]
//   - [EspressoProfilingLayerInfo.SetDebug_name]
//   - [EspressoProfilingLayerInfo.Main_engine_support]
//   - [EspressoProfilingLayerInfo.SetMain_engine_support]
//   - [EspressoProfilingLayerInfo.Name]
//   - [EspressoProfilingLayerInfo.SetName]
//   - [EspressoProfilingLayerInfo.Per_platform_support]
//   - [EspressoProfilingLayerInfo.SetPer_platform_support]
//   - [EspressoProfilingLayerInfo.Runtimes]
//   - [EspressoProfilingLayerInfo.SetRuntimes]
//   - [EspressoProfilingLayerInfo.Selected_runtime_engine]
//   - [EspressoProfilingLayerInfo.SetSelected_runtime_engine]
// See: https://developer.apple.com/documentation/Espresso/EspressoProfilingLayerInfo
type EspressoProfilingLayerInfo struct {
	objectivec.Object
}

// EspressoProfilingLayerInfoFromID constructs a [EspressoProfilingLayerInfo] from an objc.ID.
func EspressoProfilingLayerInfoFromID(id objc.ID) EspressoProfilingLayerInfo {
	return EspressoProfilingLayerInfo{objectivec.Object{id}}
}
// Ensure EspressoProfilingLayerInfo implements IEspressoProfilingLayerInfo.
var _ IEspressoProfilingLayerInfo = EspressoProfilingLayerInfo{}





// An interface definition for the [EspressoProfilingLayerInfo] class.
//
// # Methods
//
//   - [IEspressoProfilingLayerInfo.Average_runtime]
//   - [IEspressoProfilingLayerInfo.SetAverage_runtime]
//   - [IEspressoProfilingLayerInfo.Debug_name]
//   - [IEspressoProfilingLayerInfo.SetDebug_name]
//   - [IEspressoProfilingLayerInfo.Main_engine_support]
//   - [IEspressoProfilingLayerInfo.SetMain_engine_support]
//   - [IEspressoProfilingLayerInfo.Name]
//   - [IEspressoProfilingLayerInfo.SetName]
//   - [IEspressoProfilingLayerInfo.Per_platform_support]
//   - [IEspressoProfilingLayerInfo.SetPer_platform_support]
//   - [IEspressoProfilingLayerInfo.Runtimes]
//   - [IEspressoProfilingLayerInfo.SetRuntimes]
//   - [IEspressoProfilingLayerInfo.Selected_runtime_engine]
//   - [IEspressoProfilingLayerInfo.SetSelected_runtime_engine]
//
// See: https://developer.apple.com/documentation/Espresso/EspressoProfilingLayerInfo
type IEspressoProfilingLayerInfo interface {
	objectivec.IObject

	// Topic: Methods

	Average_runtime() float64
	SetAverage_runtime(value float64)
	Debug_name() string
	SetDebug_name(value string)
	Main_engine_support() IEspressoProfilingLayerSupportInfo
	SetMain_engine_support(value IEspressoProfilingLayerSupportInfo)
	Name() string
	SetName(value string)
	Per_platform_support() foundation.INSDictionary
	SetPer_platform_support(value foundation.INSDictionary)
	Runtimes() foundation.INSArray
	SetRuntimes(value foundation.INSArray)
	Selected_runtime_engine() int
	SetSelected_runtime_engine(value int)
}





// Init initializes the instance.
func (e EspressoProfilingLayerInfo) Init() EspressoProfilingLayerInfo {
	rv := objc.Send[EspressoProfilingLayerInfo](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e EspressoProfilingLayerInfo) Autorelease() EspressoProfilingLayerInfo {
	rv := objc.Send[EspressoProfilingLayerInfo](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewEspressoProfilingLayerInfo creates a new EspressoProfilingLayerInfo instance.
func NewEspressoProfilingLayerInfo() EspressoProfilingLayerInfo {
	class := getEspressoProfilingLayerInfoClass()
	rv := objc.Send[EspressoProfilingLayerInfo](objc.ID(class.class), objc.Sel("new"))
	return rv
}





















// See: https://developer.apple.com/documentation/Espresso/EspressoProfilingLayerInfo/average_runtime
func (e EspressoProfilingLayerInfo) Average_runtime() float64 {
	rv := objc.Send[float64](e.ID, objc.Sel("average_runtime"))
	return rv
}
func (e EspressoProfilingLayerInfo) SetAverage_runtime(value float64) {
	objc.Send[struct{}](e.ID, objc.Sel("setAverage_runtime:"), value)
}



// See: https://developer.apple.com/documentation/Espresso/EspressoProfilingLayerInfo/debug_name
func (e EspressoProfilingLayerInfo) Debug_name() string {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("debug_name"))
	return foundation.NSStringFromID(rv).String()
}
func (e EspressoProfilingLayerInfo) SetDebug_name(value string) {
	objc.Send[struct{}](e.ID, objc.Sel("setDebug_name:"), objc.String(value))
}



// See: https://developer.apple.com/documentation/Espresso/EspressoProfilingLayerInfo/main_engine_support
func (e EspressoProfilingLayerInfo) Main_engine_support() IEspressoProfilingLayerSupportInfo {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("main_engine_support"))
	return EspressoProfilingLayerSupportInfoFromID(objc.ID(rv))
}
func (e EspressoProfilingLayerInfo) SetMain_engine_support(value IEspressoProfilingLayerSupportInfo) {
	objc.Send[struct{}](e.ID, objc.Sel("setMain_engine_support:"), value)
}



// See: https://developer.apple.com/documentation/Espresso/EspressoProfilingLayerInfo/name
func (e EspressoProfilingLayerInfo) Name() string {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("name"))
	return foundation.NSStringFromID(rv).String()
}
func (e EspressoProfilingLayerInfo) SetName(value string) {
	objc.Send[struct{}](e.ID, objc.Sel("setName:"), objc.String(value))
}



// See: https://developer.apple.com/documentation/Espresso/EspressoProfilingLayerInfo/per_platform_support
func (e EspressoProfilingLayerInfo) Per_platform_support() foundation.INSDictionary {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("per_platform_support"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (e EspressoProfilingLayerInfo) SetPer_platform_support(value foundation.INSDictionary) {
	objc.Send[struct{}](e.ID, objc.Sel("setPer_platform_support:"), value)
}



// See: https://developer.apple.com/documentation/Espresso/EspressoProfilingLayerInfo/runtimes
func (e EspressoProfilingLayerInfo) Runtimes() foundation.INSArray {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("runtimes"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (e EspressoProfilingLayerInfo) SetRuntimes(value foundation.INSArray) {
	objc.Send[struct{}](e.ID, objc.Sel("setRuntimes:"), value)
}



// See: https://developer.apple.com/documentation/Espresso/EspressoProfilingLayerInfo/selected_runtime_engine
func (e EspressoProfilingLayerInfo) Selected_runtime_engine() int {
	rv := objc.Send[int](e.ID, objc.Sel("selected_runtime_engine"))
	return rv
}
func (e EspressoProfilingLayerInfo) SetSelected_runtime_engine(value int) {
	objc.Send[struct{}](e.ID, objc.Sel("setSelected_runtime_engine:"), value)
}

















