// Code generated from Apple documentation for gtshaderprofiler. DO NOT EDIT.

package gtshaderprofiler

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [GTMioCounterDataPerDM] class.
var (
	_GTMioCounterDataPerDMClass     GTMioCounterDataPerDMClass
	_GTMioCounterDataPerDMClassOnce sync.Once
)

func getGTMioCounterDataPerDMClass() GTMioCounterDataPerDMClass {
	_GTMioCounterDataPerDMClassOnce.Do(func() {
		_GTMioCounterDataPerDMClass = GTMioCounterDataPerDMClass{class: objc.GetClass("GTMioCounterDataPerDM")}
	})
	return _GTMioCounterDataPerDMClass
}

// GetGTMioCounterDataPerDMClass returns the class object for GTMioCounterDataPerDM.
func GetGTMioCounterDataPerDMClass() GTMioCounterDataPerDMClass {
	return getGTMioCounterDataPerDMClass()
}

type GTMioCounterDataPerDMClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (gc GTMioCounterDataPerDMClass) Class() objc.Class {
	return gc.class
}

// Alloc allocates memory for a new instance of the class.
func (gc GTMioCounterDataPerDMClass) Alloc() GTMioCounterDataPerDM {
	rv := objc.Send[GTMioCounterDataPerDM](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [GTMioCounterDataPerDM._cacheValues]
//   - [GTMioCounterDataPerDM.CounterIndex]
//   - [GTMioCounterDataPerDM.DataMaster]
//   - [GTMioCounterDataPerDM.MaxValue]
//   - [GTMioCounterDataPerDM.MinValue]
//   - [GTMioCounterDataPerDM.Name]
//   - [GTMioCounterDataPerDM.OriginalName]
//   - [GTMioCounterDataPerDM.SampleCount]
//   - [GTMioCounterDataPerDM.Scope]
//   - [GTMioCounterDataPerDM.ScopeIndex]
//   - [GTMioCounterDataPerDM.Values]
//   - [GTMioCounterDataPerDM.InitWithContainerIndexDataMasterScopeScopeIndex]
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioCounterDataPerDM
type GTMioCounterDataPerDM struct {
	objectivec.Object
}

// GTMioCounterDataPerDMFromID constructs a [GTMioCounterDataPerDM] from an objc.ID.
func GTMioCounterDataPerDMFromID(id objc.ID) GTMioCounterDataPerDM {
	return GTMioCounterDataPerDM{objectivec.Object{ID: id}}
}

// Ensure GTMioCounterDataPerDM implements IGTMioCounterDataPerDM.
var _ IGTMioCounterDataPerDM = GTMioCounterDataPerDM{}

// An interface definition for the [GTMioCounterDataPerDM] class.
//
// # Methods
//
//   - [IGTMioCounterDataPerDM._cacheValues]
//   - [IGTMioCounterDataPerDM.CounterIndex]
//   - [IGTMioCounterDataPerDM.DataMaster]
//   - [IGTMioCounterDataPerDM.MaxValue]
//   - [IGTMioCounterDataPerDM.MinValue]
//   - [IGTMioCounterDataPerDM.Name]
//   - [IGTMioCounterDataPerDM.OriginalName]
//   - [IGTMioCounterDataPerDM.SampleCount]
//   - [IGTMioCounterDataPerDM.Scope]
//   - [IGTMioCounterDataPerDM.ScopeIndex]
//   - [IGTMioCounterDataPerDM.Values]
//   - [IGTMioCounterDataPerDM.InitWithContainerIndexDataMasterScopeScopeIndex]
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioCounterDataPerDM
type IGTMioCounterDataPerDM interface {
	objectivec.IObject

	// Topic: Methods

	_cacheValues()
	CounterIndex() uint64
	DataMaster() uint16
	MaxValue() float64
	MinValue() float64
	Name() string
	OriginalName() string
	SampleCount() uint64
	Scope() uint16
	ScopeIndex() uint64
	Values() []float64
	InitWithContainerIndexDataMasterScopeScopeIndex(container unsafe.Pointer, index uint64, master uint16, scope uint16, index2 uint64) GTMioCounterDataPerDM
}

// Init initializes the instance.
func (g GTMioCounterDataPerDM) Init() GTMioCounterDataPerDM {
	rv := objc.Send[GTMioCounterDataPerDM](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g GTMioCounterDataPerDM) Autorelease() GTMioCounterDataPerDM {
	rv := objc.Send[GTMioCounterDataPerDM](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewGTMioCounterDataPerDM creates a new GTMioCounterDataPerDM instance.
func NewGTMioCounterDataPerDM() GTMioCounterDataPerDM {
	class := getGTMioCounterDataPerDMClass()
	rv := objc.Send[GTMioCounterDataPerDM](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioCounterDataPerDM/initWithContainer:index:dataMaster:scope:scopeIndex:
func NewGTMioCounterDataPerDMWithContainerIndexDataMasterScopeScopeIndex(container unsafe.Pointer, index uint64, master uint16, scope uint16, index2 uint64) GTMioCounterDataPerDM {
	instance := getGTMioCounterDataPerDMClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithContainer:index:dataMaster:scope:scopeIndex:"), container, index, master, scope, index2)
	return GTMioCounterDataPerDMFromID(rv)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioCounterDataPerDM/_cacheValues
func (g GTMioCounterDataPerDM) _cacheValues() {
	objc.Send[objc.ID](g.ID, objc.Sel("_cacheValues"))
}

// CacheValues is an exported wrapper for the private method _cacheValues.
func (g GTMioCounterDataPerDM) CacheValues() {
	g._cacheValues()
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioCounterDataPerDM/initWithContainer:index:dataMaster:scope:scopeIndex:
func (g GTMioCounterDataPerDM) InitWithContainerIndexDataMasterScopeScopeIndex(container unsafe.Pointer, index uint64, master uint16, scope uint16, index2 uint64) GTMioCounterDataPerDM {
	rv := objc.Send[GTMioCounterDataPerDM](g.ID, objc.Sel("initWithContainer:index:dataMaster:scope:scopeIndex:"), container, index, master, scope, index2)
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioCounterDataPerDM/counterIndex
func (g GTMioCounterDataPerDM) CounterIndex() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("counterIndex"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioCounterDataPerDM/dataMaster
func (g GTMioCounterDataPerDM) DataMaster() uint16 {
	rv := objc.Send[uint16](g.ID, objc.Sel("dataMaster"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioCounterDataPerDM/maxValue
func (g GTMioCounterDataPerDM) MaxValue() float64 {
	rv := objc.Send[float64](g.ID, objc.Sel("maxValue"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioCounterDataPerDM/minValue
func (g GTMioCounterDataPerDM) MinValue() float64 {
	rv := objc.Send[float64](g.ID, objc.Sel("minValue"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioCounterDataPerDM/name
func (g GTMioCounterDataPerDM) Name() string {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("name"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioCounterDataPerDM/originalName
func (g GTMioCounterDataPerDM) OriginalName() string {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("originalName"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioCounterDataPerDM/sampleCount
func (g GTMioCounterDataPerDM) SampleCount() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("sampleCount"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioCounterDataPerDM/scope
func (g GTMioCounterDataPerDM) Scope() uint16 {
	rv := objc.Send[uint16](g.ID, objc.Sel("scope"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioCounterDataPerDM/scopeIndex
func (g GTMioCounterDataPerDM) ScopeIndex() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("scopeIndex"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTMioCounterDataPerDM/values
func (g GTMioCounterDataPerDM) Values() []float64 {
	rv := objc.Send[[]objc.ID](g.ID, objc.Sel("values"))
	return objc.ConvertSlice(rv, func(id objc.ID) float64 {
		return float64(id)
	})
}
