// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLSScreenTelemetryResultsSnapshotZoneDataWrapper] class.
var (
	_SLSScreenTelemetryResultsSnapshotZoneDataWrapperClass     SLSScreenTelemetryResultsSnapshotZoneDataWrapperClass
	_SLSScreenTelemetryResultsSnapshotZoneDataWrapperClassOnce sync.Once
)

func getSLSScreenTelemetryResultsSnapshotZoneDataWrapperClass() SLSScreenTelemetryResultsSnapshotZoneDataWrapperClass {
	_SLSScreenTelemetryResultsSnapshotZoneDataWrapperClassOnce.Do(func() {
		_SLSScreenTelemetryResultsSnapshotZoneDataWrapperClass = SLSScreenTelemetryResultsSnapshotZoneDataWrapperClass{class: objc.GetClass("SLSScreenTelemetryResultsSnapshotZoneDataWrapper")}
	})
	return _SLSScreenTelemetryResultsSnapshotZoneDataWrapperClass
}

// GetSLSScreenTelemetryResultsSnapshotZoneDataWrapperClass returns the class object for SLSScreenTelemetryResultsSnapshotZoneDataWrapper.
func GetSLSScreenTelemetryResultsSnapshotZoneDataWrapperClass() SLSScreenTelemetryResultsSnapshotZoneDataWrapperClass {
	return getSLSScreenTelemetryResultsSnapshotZoneDataWrapperClass()
}

type SLSScreenTelemetryResultsSnapshotZoneDataWrapperClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLSScreenTelemetryResultsSnapshotZoneDataWrapperClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLSScreenTelemetryResultsSnapshotZoneDataWrapperClass) Alloc() SLSScreenTelemetryResultsSnapshotZoneDataWrapper {
	rv := objc.Send[SLSScreenTelemetryResultsSnapshotZoneDataWrapper](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLSScreenTelemetryResultsSnapshotZoneDataWrapper.Column]
//   - [SLSScreenTelemetryResultsSnapshotZoneDataWrapper.Count]
//   - [SLSScreenTelemetryResultsSnapshotZoneDataWrapper.Data]
//   - [SLSScreenTelemetryResultsSnapshotZoneDataWrapper.RawData]
//   - [SLSScreenTelemetryResultsSnapshotZoneDataWrapper.Row]
//   - [SLSScreenTelemetryResultsSnapshotZoneDataWrapper.ZAverage]
//   - [SLSScreenTelemetryResultsSnapshotZoneDataWrapper.ZMax]
//   - [SLSScreenTelemetryResultsSnapshotZoneDataWrapper.ZMin]
//   - [SLSScreenTelemetryResultsSnapshotZoneDataWrapper.InitWithObject]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSScreenTelemetryResultsSnapshotZoneDataWrapper
type SLSScreenTelemetryResultsSnapshotZoneDataWrapper struct {
	objectivec.Object
}

// SLSScreenTelemetryResultsSnapshotZoneDataWrapperFromID constructs a [SLSScreenTelemetryResultsSnapshotZoneDataWrapper] from an objc.ID.
func SLSScreenTelemetryResultsSnapshotZoneDataWrapperFromID(id objc.ID) SLSScreenTelemetryResultsSnapshotZoneDataWrapper {
	return SLSScreenTelemetryResultsSnapshotZoneDataWrapper{objectivec.Object{ID: id}}
}

// Ensure SLSScreenTelemetryResultsSnapshotZoneDataWrapper implements ISLSScreenTelemetryResultsSnapshotZoneDataWrapper.
var _ ISLSScreenTelemetryResultsSnapshotZoneDataWrapper = SLSScreenTelemetryResultsSnapshotZoneDataWrapper{}

// An interface definition for the [SLSScreenTelemetryResultsSnapshotZoneDataWrapper] class.
//
// # Methods
//
//   - [ISLSScreenTelemetryResultsSnapshotZoneDataWrapper.Column]
//   - [ISLSScreenTelemetryResultsSnapshotZoneDataWrapper.Count]
//   - [ISLSScreenTelemetryResultsSnapshotZoneDataWrapper.Data]
//   - [ISLSScreenTelemetryResultsSnapshotZoneDataWrapper.RawData]
//   - [ISLSScreenTelemetryResultsSnapshotZoneDataWrapper.Row]
//   - [ISLSScreenTelemetryResultsSnapshotZoneDataWrapper.ZAverage]
//   - [ISLSScreenTelemetryResultsSnapshotZoneDataWrapper.ZMax]
//   - [ISLSScreenTelemetryResultsSnapshotZoneDataWrapper.ZMin]
//   - [ISLSScreenTelemetryResultsSnapshotZoneDataWrapper.InitWithObject]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSScreenTelemetryResultsSnapshotZoneDataWrapper
type ISLSScreenTelemetryResultsSnapshotZoneDataWrapper interface {
	objectivec.IObject

	// Topic: Methods

	Column() uint64
	Count() float32
	Data() objectivec.IObject
	RawData() unsafe.Pointer
	Row() uint64
	ZAverage() float32
	ZMax() float32
	ZMin() float32
	InitWithObject(object objectivec.IObject) SLSScreenTelemetryResultsSnapshotZoneDataWrapper
}

// Init initializes the instance.
func (s SLSScreenTelemetryResultsSnapshotZoneDataWrapper) Init() SLSScreenTelemetryResultsSnapshotZoneDataWrapper {
	rv := objc.Send[SLSScreenTelemetryResultsSnapshotZoneDataWrapper](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLSScreenTelemetryResultsSnapshotZoneDataWrapper) Autorelease() SLSScreenTelemetryResultsSnapshotZoneDataWrapper {
	rv := objc.Send[SLSScreenTelemetryResultsSnapshotZoneDataWrapper](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLSScreenTelemetryResultsSnapshotZoneDataWrapper creates a new SLSScreenTelemetryResultsSnapshotZoneDataWrapper instance.
func NewSLSScreenTelemetryResultsSnapshotZoneDataWrapper() SLSScreenTelemetryResultsSnapshotZoneDataWrapper {
	class := getSLSScreenTelemetryResultsSnapshotZoneDataWrapperClass()
	rv := objc.Send[SLSScreenTelemetryResultsSnapshotZoneDataWrapper](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSScreenTelemetryResultsSnapshotZoneDataWrapper/initWithObject:
func NewSLSScreenTelemetryResultsSnapshotZoneDataWrapperWithObject(object objectivec.IObject) SLSScreenTelemetryResultsSnapshotZoneDataWrapper {
	instance := getSLSScreenTelemetryResultsSnapshotZoneDataWrapperClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithObject:"), object)
	return SLSScreenTelemetryResultsSnapshotZoneDataWrapperFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSScreenTelemetryResultsSnapshotZoneDataWrapper/column
func (s SLSScreenTelemetryResultsSnapshotZoneDataWrapper) Column() uint64 {
	rv := objc.Send[uint64](s.ID, objc.Sel("column"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSScreenTelemetryResultsSnapshotZoneDataWrapper/count
func (s SLSScreenTelemetryResultsSnapshotZoneDataWrapper) Count() float32 {
	rv := objc.Send[float32](s.ID, objc.Sel("count"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSScreenTelemetryResultsSnapshotZoneDataWrapper/rawData
func (s SLSScreenTelemetryResultsSnapshotZoneDataWrapper) RawData() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s.ID, objc.Sel("rawData"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSScreenTelemetryResultsSnapshotZoneDataWrapper/row
func (s SLSScreenTelemetryResultsSnapshotZoneDataWrapper) Row() uint64 {
	rv := objc.Send[uint64](s.ID, objc.Sel("row"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSScreenTelemetryResultsSnapshotZoneDataWrapper/zAverage
func (s SLSScreenTelemetryResultsSnapshotZoneDataWrapper) ZAverage() float32 {
	rv := objc.Send[float32](s.ID, objc.Sel("zAverage"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSScreenTelemetryResultsSnapshotZoneDataWrapper/zMax
func (s SLSScreenTelemetryResultsSnapshotZoneDataWrapper) ZMax() float32 {
	rv := objc.Send[float32](s.ID, objc.Sel("zMax"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSScreenTelemetryResultsSnapshotZoneDataWrapper/zMin
func (s SLSScreenTelemetryResultsSnapshotZoneDataWrapper) ZMin() float32 {
	rv := objc.Send[float32](s.ID, objc.Sel("zMin"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSScreenTelemetryResultsSnapshotZoneDataWrapper/initWithObject:
func (s SLSScreenTelemetryResultsSnapshotZoneDataWrapper) InitWithObject(object objectivec.IObject) SLSScreenTelemetryResultsSnapshotZoneDataWrapper {
	rv := objc.Send[SLSScreenTelemetryResultsSnapshotZoneDataWrapper](s.ID, objc.Sel("initWithObject:"), object)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSScreenTelemetryResultsSnapshotZoneDataWrapper/wrapperWithObject:
func (_SLSScreenTelemetryResultsSnapshotZoneDataWrapperClass SLSScreenTelemetryResultsSnapshotZoneDataWrapperClass) WrapperWithObject(object objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_SLSScreenTelemetryResultsSnapshotZoneDataWrapperClass.class), objc.Sel("wrapperWithObject:"), object)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/SLSScreenTelemetryResultsSnapshotZoneDataWrapper/data
func (s SLSScreenTelemetryResultsSnapshotZoneDataWrapper) Data() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("data"))
	return objectivec.Object{ID: rv}
}
