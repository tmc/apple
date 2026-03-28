// Code generated from Apple documentation for gtshaderprofiler. DO NOT EDIT.

package gtshaderprofiler

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [GTShaderProfilerDiassemblyRegisterPressure] class.
var (
	_GTShaderProfilerDiassemblyRegisterPressureClass     GTShaderProfilerDiassemblyRegisterPressureClass
	_GTShaderProfilerDiassemblyRegisterPressureClassOnce sync.Once
)

func getGTShaderProfilerDiassemblyRegisterPressureClass() GTShaderProfilerDiassemblyRegisterPressureClass {
	_GTShaderProfilerDiassemblyRegisterPressureClassOnce.Do(func() {
		_GTShaderProfilerDiassemblyRegisterPressureClass = GTShaderProfilerDiassemblyRegisterPressureClass{class: objc.GetClass("GTShaderProfilerDiassemblyRegisterPressure")}
	})
	return _GTShaderProfilerDiassemblyRegisterPressureClass
}

// GetGTShaderProfilerDiassemblyRegisterPressureClass returns the class object for GTShaderProfilerDiassemblyRegisterPressure.
func GetGTShaderProfilerDiassemblyRegisterPressureClass() GTShaderProfilerDiassemblyRegisterPressureClass {
	return getGTShaderProfilerDiassemblyRegisterPressureClass()
}

type GTShaderProfilerDiassemblyRegisterPressureClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (gc GTShaderProfilerDiassemblyRegisterPressureClass) Class() objc.Class {
	return gc.class
}

// Alloc allocates memory for a new instance of the class.
func (gc GTShaderProfilerDiassemblyRegisterPressureClass) Alloc() GTShaderProfilerDiassemblyRegisterPressure {
	rv := objc.Send[GTShaderProfilerDiassemblyRegisterPressure](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [GTShaderProfilerDiassemblyRegisterPressure.Allocs]
//   - [GTShaderProfilerDiassemblyRegisterPressure.Combine]
//   - [GTShaderProfilerDiassemblyRegisterPressure.Defs]
//   - [GTShaderProfilerDiassemblyRegisterPressure.HighRegisterIndex]
//   - [GTShaderProfilerDiassemblyRegisterPressure.LastUses]
//   - [GTShaderProfilerDiassemblyRegisterPressure.Live]
//   - [GTShaderProfilerDiassemblyRegisterPressure.LiveRegisters]
//   - [GTShaderProfilerDiassemblyRegisterPressure.Uses]
//   - [GTShaderProfilerDiassemblyRegisterPressure.InitWithDict]
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerDiassemblyRegisterPressure
type GTShaderProfilerDiassemblyRegisterPressure struct {
	objectivec.Object
}

// GTShaderProfilerDiassemblyRegisterPressureFromID constructs a [GTShaderProfilerDiassemblyRegisterPressure] from an objc.ID.
func GTShaderProfilerDiassemblyRegisterPressureFromID(id objc.ID) GTShaderProfilerDiassemblyRegisterPressure {
	return GTShaderProfilerDiassemblyRegisterPressure{objectivec.Object{ID: id}}
}
// Ensure GTShaderProfilerDiassemblyRegisterPressure implements IGTShaderProfilerDiassemblyRegisterPressure.
var _ IGTShaderProfilerDiassemblyRegisterPressure = GTShaderProfilerDiassemblyRegisterPressure{}

// An interface definition for the [GTShaderProfilerDiassemblyRegisterPressure] class.
//
// # Methods
//
//   - [IGTShaderProfilerDiassemblyRegisterPressure.Allocs]
//   - [IGTShaderProfilerDiassemblyRegisterPressure.Combine]
//   - [IGTShaderProfilerDiassemblyRegisterPressure.Defs]
//   - [IGTShaderProfilerDiassemblyRegisterPressure.HighRegisterIndex]
//   - [IGTShaderProfilerDiassemblyRegisterPressure.LastUses]
//   - [IGTShaderProfilerDiassemblyRegisterPressure.Live]
//   - [IGTShaderProfilerDiassemblyRegisterPressure.LiveRegisters]
//   - [IGTShaderProfilerDiassemblyRegisterPressure.Uses]
//   - [IGTShaderProfilerDiassemblyRegisterPressure.InitWithDict]
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerDiassemblyRegisterPressure
type IGTShaderProfilerDiassemblyRegisterPressure interface {
	objectivec.IObject

	// Topic: Methods

	Allocs() IGTShaderProfilerRegisterUsage
	Combine(combine objectivec.IObject)
	Defs() IGTShaderProfilerRegisterUsage
	HighRegisterIndex() uint64
	LastUses() IGTShaderProfilerRegisterUsage
	Live() IGTShaderProfilerRegisterUsage
	LiveRegisters() uint64
	Uses() IGTShaderProfilerRegisterUsage
	InitWithDict(dict objectivec.IObject) GTShaderProfilerDiassemblyRegisterPressure
}

// Init initializes the instance.
func (g GTShaderProfilerDiassemblyRegisterPressure) Init() GTShaderProfilerDiassemblyRegisterPressure {
	rv := objc.Send[GTShaderProfilerDiassemblyRegisterPressure](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g GTShaderProfilerDiassemblyRegisterPressure) Autorelease() GTShaderProfilerDiassemblyRegisterPressure {
	rv := objc.Send[GTShaderProfilerDiassemblyRegisterPressure](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewGTShaderProfilerDiassemblyRegisterPressure creates a new GTShaderProfilerDiassemblyRegisterPressure instance.
func NewGTShaderProfilerDiassemblyRegisterPressure() GTShaderProfilerDiassemblyRegisterPressure {
	class := getGTShaderProfilerDiassemblyRegisterPressureClass()
	rv := objc.Send[GTShaderProfilerDiassemblyRegisterPressure](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerDiassemblyRegisterPressure/initWithDict:
func NewGTShaderProfilerDiassemblyRegisterPressureWithDict(dict objectivec.IObject) GTShaderProfilerDiassemblyRegisterPressure {
	instance := getGTShaderProfilerDiassemblyRegisterPressureClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDict:"), dict)
	return GTShaderProfilerDiassemblyRegisterPressureFromID(rv)
}

//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerDiassemblyRegisterPressure/combine:
func (g GTShaderProfilerDiassemblyRegisterPressure) Combine(combine objectivec.IObject) {
	objc.Send[objc.ID](g.ID, objc.Sel("combine:"), combine)
}
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerDiassemblyRegisterPressure/initWithDict:
func (g GTShaderProfilerDiassemblyRegisterPressure) InitWithDict(dict objectivec.IObject) GTShaderProfilerDiassemblyRegisterPressure {
	rv := objc.Send[GTShaderProfilerDiassemblyRegisterPressure](g.ID, objc.Sel("initWithDict:"), dict)
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerDiassemblyRegisterPressure/allocs
func (g GTShaderProfilerDiassemblyRegisterPressure) Allocs() IGTShaderProfilerRegisterUsage {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("allocs"))
	return GTShaderProfilerRegisterUsageFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerDiassemblyRegisterPressure/defs
func (g GTShaderProfilerDiassemblyRegisterPressure) Defs() IGTShaderProfilerRegisterUsage {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("defs"))
	return GTShaderProfilerRegisterUsageFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerDiassemblyRegisterPressure/highRegisterIndex
func (g GTShaderProfilerDiassemblyRegisterPressure) HighRegisterIndex() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("highRegisterIndex"))
	return rv
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerDiassemblyRegisterPressure/lastUses
func (g GTShaderProfilerDiassemblyRegisterPressure) LastUses() IGTShaderProfilerRegisterUsage {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("lastUses"))
	return GTShaderProfilerRegisterUsageFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerDiassemblyRegisterPressure/live
func (g GTShaderProfilerDiassemblyRegisterPressure) Live() IGTShaderProfilerRegisterUsage {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("live"))
	return GTShaderProfilerRegisterUsageFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerDiassemblyRegisterPressure/liveRegisters
func (g GTShaderProfilerDiassemblyRegisterPressure) LiveRegisters() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("liveRegisters"))
	return rv
}
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTShaderProfilerDiassemblyRegisterPressure/uses
func (g GTShaderProfilerDiassemblyRegisterPressure) Uses() IGTShaderProfilerRegisterUsage {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("uses"))
	return GTShaderProfilerRegisterUsageFromID(objc.ID(rv))
}

