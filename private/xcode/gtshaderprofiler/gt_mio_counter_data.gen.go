// Code generated from Apple documentation for gtshaderprofiler. DO NOT EDIT.

package gtshaderprofiler

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [GTMioCounterData] class.
var (
	_GTMioCounterDataClass     GTMioCounterDataClass
	_GTMioCounterDataClassOnce sync.Once
)

func getGTMioCounterDataClass() GTMioCounterDataClass {
	_GTMioCounterDataClassOnce.Do(func() {
		_GTMioCounterDataClass = GTMioCounterDataClass{class: objc.GetClass("GTMioCounterData")}
	})
	return _GTMioCounterDataClass
}

// GetGTMioCounterDataClass returns the class object for GTMioCounterData.
func GetGTMioCounterDataClass() GTMioCounterDataClass {
	return getGTMioCounterDataClass()
}

type GTMioCounterDataClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (gc GTMioCounterDataClass) Class() objc.Class {
	return gc.class
}

// Alloc allocates memory for a new instance of the class.
func (gc GTMioCounterDataClass) Alloc() GTMioCounterData {
	rv := objc.SendIfResponds[GTMioCounterData](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [GTMioCounterData.CounterIndex]
//   - [GTMioCounterData.DataType]
//   - [GTMioCounterData.MaxValue]
//   - [GTMioCounterData.MinValue]
//   - [GTMioCounterData.Name]
//   - [GTMioCounterData.SampleCount]
//   - [GTMioCounterData.SampleInterval]
//   - [GTMioCounterData.Scope]
//   - [GTMioCounterData.ScopeIndex]
//   - [GTMioCounterData.Timestamps]
//   - [GTMioCounterData.ValueType]
//   - [GTMioCounterData.Values]
//   - [GTMioCounterData.InitWithContainerIndexScopeScopeIndex]
type GTMioCounterData struct {
	objectivec.Object
}

// GTMioCounterDataFromID constructs a [GTMioCounterData] from an objc.ID.
func GTMioCounterDataFromID(id objc.ID) GTMioCounterData {
	return GTMioCounterData{objectivec.Object{ID: id}}
}

// Ensure GTMioCounterData implements IGTMioCounterData.
var _ IGTMioCounterData = GTMioCounterData{}

// An interface definition for the [GTMioCounterData] class.
//
// # Methods
//
//   - [IGTMioCounterData.CounterIndex]
//   - [IGTMioCounterData.DataType]
//   - [IGTMioCounterData.MaxValue]
//   - [IGTMioCounterData.MinValue]
//   - [IGTMioCounterData.Name]
//   - [IGTMioCounterData.SampleCount]
//   - [IGTMioCounterData.SampleInterval]
//   - [IGTMioCounterData.Scope]
//   - [IGTMioCounterData.ScopeIndex]
//   - [IGTMioCounterData.Timestamps]
//   - [IGTMioCounterData.ValueType]
//   - [IGTMioCounterData.Values]
//   - [IGTMioCounterData.InitWithContainerIndexScopeScopeIndex]
type IGTMioCounterData interface {
	objectivec.IObject

	// Topic: Methods

	CounterIndex() uint64
	DataType() uint32
	MaxValue() float64
	MinValue() float64
	Name() string
	SampleCount() uint64
	SampleInterval() uint64
	Scope() uint16
	ScopeIndex() uint64
	Timestamps() unsafe.Pointer
	ValueType() uint16
	Values() unsafe.Pointer
	InitWithContainerIndexScopeScopeIndex(container unsafe.Pointer, index uint64, scope uint16, index2 uint64) GTMioCounterData
}

// Init initializes the instance.
func (g GTMioCounterData) Init() GTMioCounterData {
	rv := objc.SendIfResponds[GTMioCounterData](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g GTMioCounterData) Autorelease() GTMioCounterData {
	rv := objc.SendIfResponds[GTMioCounterData](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewGTMioCounterData creates a new GTMioCounterData instance.
func NewGTMioCounterData() GTMioCounterData {
	class := getGTMioCounterDataClass()
	rv := objc.SendIfResponds[GTMioCounterData](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewGTMioCounterDataWithContainerIndexScopeScopeIndex(container unsafe.Pointer, index uint64, scope uint16, index2 uint64) GTMioCounterData {
	instance := getGTMioCounterDataClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithContainer:index:scope:scopeIndex:"), container, index, scope, index2)
	return GTMioCounterDataFromID(rv)
}

func (g GTMioCounterData) InitWithContainerIndexScopeScopeIndex(container unsafe.Pointer, index uint64, scope uint16, index2 uint64) GTMioCounterData {
	rv := objc.SendIfResponds[GTMioCounterData](g.ID, objc.Sel("initWithContainer:index:scope:scopeIndex:"), container, index, scope, index2)
	return rv
}

func (g GTMioCounterData) CounterIndex() uint64 {
	rv := objc.SendIfResponds[uint64](g.ID, objc.Sel("counterIndex"))
	return rv
}
func (g GTMioCounterData) DataType() uint32 {
	rv := objc.SendIfResponds[uint32](g.ID, objc.Sel("dataType"))
	return rv
}
func (g GTMioCounterData) MaxValue() float64 {
	rv := objc.SendIfResponds[float64](g.ID, objc.Sel("maxValue"))
	return rv
}
func (g GTMioCounterData) MinValue() float64 {
	rv := objc.SendIfResponds[float64](g.ID, objc.Sel("minValue"))
	return rv
}
func (g GTMioCounterData) Name() string {
	rv := objc.SendIfResponds[objc.ID](g.ID, objc.Sel("name"))
	return foundation.NSStringFromID(rv).String()
}
func (g GTMioCounterData) SampleCount() uint64 {
	rv := objc.SendIfResponds[uint64](g.ID, objc.Sel("sampleCount"))
	return rv
}
func (g GTMioCounterData) SampleInterval() uint64 {
	rv := objc.SendIfResponds[uint64](g.ID, objc.Sel("sampleInterval"))
	return rv
}
func (g GTMioCounterData) Scope() uint16 {
	rv := objc.SendIfResponds[uint16](g.ID, objc.Sel("scope"))
	return rv
}
func (g GTMioCounterData) ScopeIndex() uint64 {
	rv := objc.SendIfResponds[uint64](g.ID, objc.Sel("scopeIndex"))
	return rv
}
func (g GTMioCounterData) Timestamps() unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](g.ID, objc.Sel("timestamps"))
	return rv
}
func (g GTMioCounterData) ValueType() uint16 {
	rv := objc.SendIfResponds[uint16](g.ID, objc.Sel("valueType"))
	return rv
}
func (g GTMioCounterData) Values() unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](g.ID, objc.Sel("values"))
	return rv
}
