// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [EspressoProfilingNetworkANEInfo] class.
var (
	_EspressoProfilingNetworkANEInfoClass     EspressoProfilingNetworkANEInfoClass
	_EspressoProfilingNetworkANEInfoClassOnce sync.Once
)

func getEspressoProfilingNetworkANEInfoClass() EspressoProfilingNetworkANEInfoClass {
	_EspressoProfilingNetworkANEInfoClassOnce.Do(func() {
		_EspressoProfilingNetworkANEInfoClass = EspressoProfilingNetworkANEInfoClass{class: objc.GetClass("EspressoProfilingNetworkANEInfo")}
	})
	return _EspressoProfilingNetworkANEInfoClass
}

// GetEspressoProfilingNetworkANEInfoClass returns the class object for EspressoProfilingNetworkANEInfo.
func GetEspressoProfilingNetworkANEInfoClass() EspressoProfilingNetworkANEInfoClass {
	return getEspressoProfilingNetworkANEInfoClass()
}

type EspressoProfilingNetworkANEInfoClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (ec EspressoProfilingNetworkANEInfoClass) Alloc() EspressoProfilingNetworkANEInfo {
	rv := objc.Send[EspressoProfilingNetworkANEInfo](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}







//
// # Methods
//
//   - [EspressoProfilingNetworkANEInfo.Ane_time_per_eval_ns]
//   - [EspressoProfilingNetworkANEInfo.SetAne_time_per_eval_ns]
//   - [EspressoProfilingNetworkANEInfo.Total_ane_time_ns]
//   - [EspressoProfilingNetworkANEInfo.SetTotal_ane_time_ns]
// See: https://developer.apple.com/documentation/Espresso/EspressoProfilingNetworkANEInfo
type EspressoProfilingNetworkANEInfo struct {
	objectivec.Object
}

// EspressoProfilingNetworkANEInfoFromID constructs a [EspressoProfilingNetworkANEInfo] from an objc.ID.
func EspressoProfilingNetworkANEInfoFromID(id objc.ID) EspressoProfilingNetworkANEInfo {
	return EspressoProfilingNetworkANEInfo{objectivec.Object{id}}
}
// Ensure EspressoProfilingNetworkANEInfo implements IEspressoProfilingNetworkANEInfo.
var _ IEspressoProfilingNetworkANEInfo = EspressoProfilingNetworkANEInfo{}





// An interface definition for the [EspressoProfilingNetworkANEInfo] class.
//
// # Methods
//
//   - [IEspressoProfilingNetworkANEInfo.Ane_time_per_eval_ns]
//   - [IEspressoProfilingNetworkANEInfo.SetAne_time_per_eval_ns]
//   - [IEspressoProfilingNetworkANEInfo.Total_ane_time_ns]
//   - [IEspressoProfilingNetworkANEInfo.SetTotal_ane_time_ns]
//
// See: https://developer.apple.com/documentation/Espresso/EspressoProfilingNetworkANEInfo
type IEspressoProfilingNetworkANEInfo interface {
	objectivec.IObject

	// Topic: Methods

	Ane_time_per_eval_ns() uint64
	SetAne_time_per_eval_ns(value uint64)
	Total_ane_time_ns() uint64
	SetTotal_ane_time_ns(value uint64)
}





// Init initializes the instance.
func (e EspressoProfilingNetworkANEInfo) Init() EspressoProfilingNetworkANEInfo {
	rv := objc.Send[EspressoProfilingNetworkANEInfo](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e EspressoProfilingNetworkANEInfo) Autorelease() EspressoProfilingNetworkANEInfo {
	rv := objc.Send[EspressoProfilingNetworkANEInfo](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewEspressoProfilingNetworkANEInfo creates a new EspressoProfilingNetworkANEInfo instance.
func NewEspressoProfilingNetworkANEInfo() EspressoProfilingNetworkANEInfo {
	class := getEspressoProfilingNetworkANEInfoClass()
	rv := objc.Send[EspressoProfilingNetworkANEInfo](objc.ID(class.class), objc.Sel("new"))
	return rv
}





















// See: https://developer.apple.com/documentation/Espresso/EspressoProfilingNetworkANEInfo/ane_time_per_eval_ns
func (e EspressoProfilingNetworkANEInfo) Ane_time_per_eval_ns() uint64 {
	rv := objc.Send[uint64](e.ID, objc.Sel("ane_time_per_eval_ns"))
	return rv
}
func (e EspressoProfilingNetworkANEInfo) SetAne_time_per_eval_ns(value uint64) {
	objc.Send[struct{}](e.ID, objc.Sel("setAne_time_per_eval_ns:"), value)
}



// See: https://developer.apple.com/documentation/Espresso/EspressoProfilingNetworkANEInfo/total_ane_time_ns
func (e EspressoProfilingNetworkANEInfo) Total_ane_time_ns() uint64 {
	rv := objc.Send[uint64](e.ID, objc.Sel("total_ane_time_ns"))
	return rv
}
func (e EspressoProfilingNetworkANEInfo) SetTotal_ane_time_ns(value uint64) {
	objc.Send[struct{}](e.ID, objc.Sel("setTotal_ane_time_ns:"), value)
}

















