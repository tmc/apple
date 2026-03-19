// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [EspressoProfilingLayerRuntime] class.
var (
	_EspressoProfilingLayerRuntimeClass     EspressoProfilingLayerRuntimeClass
	_EspressoProfilingLayerRuntimeClassOnce sync.Once
)

func getEspressoProfilingLayerRuntimeClass() EspressoProfilingLayerRuntimeClass {
	_EspressoProfilingLayerRuntimeClassOnce.Do(func() {
		_EspressoProfilingLayerRuntimeClass = EspressoProfilingLayerRuntimeClass{class: objc.GetClass("EspressoProfilingLayerRuntime")}
	})
	return _EspressoProfilingLayerRuntimeClass
}

// GetEspressoProfilingLayerRuntimeClass returns the class object for EspressoProfilingLayerRuntime.
func GetEspressoProfilingLayerRuntimeClass() EspressoProfilingLayerRuntimeClass {
	return getEspressoProfilingLayerRuntimeClass()
}

type EspressoProfilingLayerRuntimeClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (ec EspressoProfilingLayerRuntimeClass) Alloc() EspressoProfilingLayerRuntime {
	rv := objc.Send[EspressoProfilingLayerRuntime](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [EspressoProfilingLayerRuntime.End_t]
//   - [EspressoProfilingLayerRuntime.SetEnd_t]
//   - [EspressoProfilingLayerRuntime.Start_t]
//   - [EspressoProfilingLayerRuntime.SetStart_t]
// See: https://developer.apple.com/documentation/Espresso/EspressoProfilingLayerRuntime
type EspressoProfilingLayerRuntime struct {
	objectivec.Object
}

// EspressoProfilingLayerRuntimeFromID constructs a [EspressoProfilingLayerRuntime] from an objc.ID.
func EspressoProfilingLayerRuntimeFromID(id objc.ID) EspressoProfilingLayerRuntime {
	return EspressoProfilingLayerRuntime{objectivec.Object{ID: id}}
}
// Ensure EspressoProfilingLayerRuntime implements IEspressoProfilingLayerRuntime.
var _ IEspressoProfilingLayerRuntime = EspressoProfilingLayerRuntime{}

// An interface definition for the [EspressoProfilingLayerRuntime] class.
//
// # Methods
//
//   - [IEspressoProfilingLayerRuntime.End_t]
//   - [IEspressoProfilingLayerRuntime.SetEnd_t]
//   - [IEspressoProfilingLayerRuntime.Start_t]
//   - [IEspressoProfilingLayerRuntime.SetStart_t]
//
// See: https://developer.apple.com/documentation/Espresso/EspressoProfilingLayerRuntime
type IEspressoProfilingLayerRuntime interface {
	objectivec.IObject

	// Topic: Methods

	End_t() float64
	SetEnd_t(value float64)
	Start_t() float64
	SetStart_t(value float64)
}

// Init initializes the instance.
func (e EspressoProfilingLayerRuntime) Init() EspressoProfilingLayerRuntime {
	rv := objc.Send[EspressoProfilingLayerRuntime](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e EspressoProfilingLayerRuntime) Autorelease() EspressoProfilingLayerRuntime {
	rv := objc.Send[EspressoProfilingLayerRuntime](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewEspressoProfilingLayerRuntime creates a new EspressoProfilingLayerRuntime instance.
func NewEspressoProfilingLayerRuntime() EspressoProfilingLayerRuntime {
	class := getEspressoProfilingLayerRuntimeClass()
	rv := objc.Send[EspressoProfilingLayerRuntime](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Espresso/EspressoProfilingLayerRuntime/end_t
func (e EspressoProfilingLayerRuntime) End_t() float64 {
	rv := objc.Send[float64](e.ID, objc.Sel("end_t"))
	return rv
}
func (e EspressoProfilingLayerRuntime) SetEnd_t(value float64) {
	objc.Send[struct{}](e.ID, objc.Sel("setEnd_t:"), value)
}
// See: https://developer.apple.com/documentation/Espresso/EspressoProfilingLayerRuntime/start_t
func (e EspressoProfilingLayerRuntime) Start_t() float64 {
	rv := objc.Send[float64](e.ID, objc.Sel("start_t"))
	return rv
}
func (e EspressoProfilingLayerRuntime) SetStart_t(value float64) {
	objc.Send[struct{}](e.ID, objc.Sel("setStart_t:"), value)
}

