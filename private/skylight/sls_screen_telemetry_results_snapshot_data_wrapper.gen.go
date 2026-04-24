// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLSScreenTelemetryResultsSnapshotDataWrapper] class.
var (
	_SLSScreenTelemetryResultsSnapshotDataWrapperClass     SLSScreenTelemetryResultsSnapshotDataWrapperClass
	_SLSScreenTelemetryResultsSnapshotDataWrapperClassOnce sync.Once
)

func getSLSScreenTelemetryResultsSnapshotDataWrapperClass() SLSScreenTelemetryResultsSnapshotDataWrapperClass {
	_SLSScreenTelemetryResultsSnapshotDataWrapperClassOnce.Do(func() {
		_SLSScreenTelemetryResultsSnapshotDataWrapperClass = SLSScreenTelemetryResultsSnapshotDataWrapperClass{class: objc.GetClass("SLSScreenTelemetryResultsSnapshotDataWrapper")}
	})
	return _SLSScreenTelemetryResultsSnapshotDataWrapperClass
}

// GetSLSScreenTelemetryResultsSnapshotDataWrapperClass returns the class object for SLSScreenTelemetryResultsSnapshotDataWrapper.
func GetSLSScreenTelemetryResultsSnapshotDataWrapperClass() SLSScreenTelemetryResultsSnapshotDataWrapperClass {
	return getSLSScreenTelemetryResultsSnapshotDataWrapperClass()
}

type SLSScreenTelemetryResultsSnapshotDataWrapperClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLSScreenTelemetryResultsSnapshotDataWrapperClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLSScreenTelemetryResultsSnapshotDataWrapperClass) Alloc() SLSScreenTelemetryResultsSnapshotDataWrapper {
	rv := objc.Send[SLSScreenTelemetryResultsSnapshotDataWrapper](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLSScreenTelemetryResultsSnapshotDataWrapper.AZL_max]
//   - [SLSScreenTelemetryResultsSnapshotDataWrapper.AZL_mean]
//   - [SLSScreenTelemetryResultsSnapshotDataWrapper.AZL_min]
//   - [SLSScreenTelemetryResultsSnapshotDataWrapper.EDR_headroom]
//   - [SLSScreenTelemetryResultsSnapshotDataWrapper.ColumnCount]
//   - [SLSScreenTelemetryResultsSnapshotDataWrapper.Data]
//   - [SLSScreenTelemetryResultsSnapshotDataWrapper.ObjectAtIndexedSubscript]
//   - [SLSScreenTelemetryResultsSnapshotDataWrapper.Panel]
//   - [SLSScreenTelemetryResultsSnapshotDataWrapper.RawData]
//   - [SLSScreenTelemetryResultsSnapshotDataWrapper.RowCount]
//   - [SLSScreenTelemetryResultsSnapshotDataWrapper.Timestamp]
//   - [SLSScreenTelemetryResultsSnapshotDataWrapper.InitWithObject]
//   - [SLSScreenTelemetryResultsSnapshotDataWrapper.InitWithXPCObject]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSScreenTelemetryResultsSnapshotDataWrapper
type SLSScreenTelemetryResultsSnapshotDataWrapper struct {
	objectivec.Object
}

// SLSScreenTelemetryResultsSnapshotDataWrapperFromID constructs a [SLSScreenTelemetryResultsSnapshotDataWrapper] from an objc.ID.
func SLSScreenTelemetryResultsSnapshotDataWrapperFromID(id objc.ID) SLSScreenTelemetryResultsSnapshotDataWrapper {
	return SLSScreenTelemetryResultsSnapshotDataWrapper{objectivec.Object{ID: id}}
}

// Ensure SLSScreenTelemetryResultsSnapshotDataWrapper implements ISLSScreenTelemetryResultsSnapshotDataWrapper.
var _ ISLSScreenTelemetryResultsSnapshotDataWrapper = SLSScreenTelemetryResultsSnapshotDataWrapper{}

// An interface definition for the [SLSScreenTelemetryResultsSnapshotDataWrapper] class.
//
// # Methods
//
//   - [ISLSScreenTelemetryResultsSnapshotDataWrapper.AZL_max]
//   - [ISLSScreenTelemetryResultsSnapshotDataWrapper.AZL_mean]
//   - [ISLSScreenTelemetryResultsSnapshotDataWrapper.AZL_min]
//   - [ISLSScreenTelemetryResultsSnapshotDataWrapper.EDR_headroom]
//   - [ISLSScreenTelemetryResultsSnapshotDataWrapper.ColumnCount]
//   - [ISLSScreenTelemetryResultsSnapshotDataWrapper.Data]
//   - [ISLSScreenTelemetryResultsSnapshotDataWrapper.ObjectAtIndexedSubscript]
//   - [ISLSScreenTelemetryResultsSnapshotDataWrapper.Panel]
//   - [ISLSScreenTelemetryResultsSnapshotDataWrapper.RawData]
//   - [ISLSScreenTelemetryResultsSnapshotDataWrapper.RowCount]
//   - [ISLSScreenTelemetryResultsSnapshotDataWrapper.Timestamp]
//   - [ISLSScreenTelemetryResultsSnapshotDataWrapper.InitWithObject]
//   - [ISLSScreenTelemetryResultsSnapshotDataWrapper.InitWithXPCObject]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSScreenTelemetryResultsSnapshotDataWrapper
type ISLSScreenTelemetryResultsSnapshotDataWrapper interface {
	objectivec.IObject

	// Topic: Methods

	AZL_max() float32
	AZL_mean() float32
	AZL_min() float32
	EDR_headroom() float32
	ColumnCount() uint64
	Data() objectivec.IObject
	ObjectAtIndexedSubscript(subscript uint64) objectivec.IObject
	Panel() objectivec.IObject
	RawData() unsafe.Pointer
	RowCount() uint64
	Timestamp() float64
	InitWithObject(object objectivec.IObject) SLSScreenTelemetryResultsSnapshotDataWrapper
	InitWithXPCObject(xPCObject objectivec.IObject) SLSScreenTelemetryResultsSnapshotDataWrapper
}

// Init initializes the instance.
func (s SLSScreenTelemetryResultsSnapshotDataWrapper) Init() SLSScreenTelemetryResultsSnapshotDataWrapper {
	rv := objc.Send[SLSScreenTelemetryResultsSnapshotDataWrapper](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLSScreenTelemetryResultsSnapshotDataWrapper) Autorelease() SLSScreenTelemetryResultsSnapshotDataWrapper {
	rv := objc.Send[SLSScreenTelemetryResultsSnapshotDataWrapper](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLSScreenTelemetryResultsSnapshotDataWrapper creates a new SLSScreenTelemetryResultsSnapshotDataWrapper instance.
func NewSLSScreenTelemetryResultsSnapshotDataWrapper() SLSScreenTelemetryResultsSnapshotDataWrapper {
	class := getSLSScreenTelemetryResultsSnapshotDataWrapperClass()
	rv := objc.Send[SLSScreenTelemetryResultsSnapshotDataWrapper](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSScreenTelemetryResultsSnapshotDataWrapper/initWithObject:
func NewSLSScreenTelemetryResultsSnapshotDataWrapperWithObject(object objectivec.IObject) SLSScreenTelemetryResultsSnapshotDataWrapper {
	instance := getSLSScreenTelemetryResultsSnapshotDataWrapperClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithObject:"), object)
	return SLSScreenTelemetryResultsSnapshotDataWrapperFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSScreenTelemetryResultsSnapshotDataWrapper/initWithXPCObject:
func NewSLSScreenTelemetryResultsSnapshotDataWrapperWithXPCObject(xPCObject objectivec.IObject) SLSScreenTelemetryResultsSnapshotDataWrapper {
	instance := getSLSScreenTelemetryResultsSnapshotDataWrapperClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithXPCObject:"), xPCObject)
	return SLSScreenTelemetryResultsSnapshotDataWrapperFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSScreenTelemetryResultsSnapshotDataWrapper/AZL_max
func (s SLSScreenTelemetryResultsSnapshotDataWrapper) AZL_max() float32 {
	rv := objc.Send[float32](s.ID, objc.Sel("AZL_max"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSScreenTelemetryResultsSnapshotDataWrapper/AZL_mean
func (s SLSScreenTelemetryResultsSnapshotDataWrapper) AZL_mean() float32 {
	rv := objc.Send[float32](s.ID, objc.Sel("AZL_mean"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSScreenTelemetryResultsSnapshotDataWrapper/AZL_min
func (s SLSScreenTelemetryResultsSnapshotDataWrapper) AZL_min() float32 {
	rv := objc.Send[float32](s.ID, objc.Sel("AZL_min"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSScreenTelemetryResultsSnapshotDataWrapper/EDR_headroom
func (s SLSScreenTelemetryResultsSnapshotDataWrapper) EDR_headroom() float32 {
	rv := objc.Send[float32](s.ID, objc.Sel("EDR_headroom"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSScreenTelemetryResultsSnapshotDataWrapper/columnCount
func (s SLSScreenTelemetryResultsSnapshotDataWrapper) ColumnCount() uint64 {
	rv := objc.Send[uint64](s.ID, objc.Sel("columnCount"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSScreenTelemetryResultsSnapshotDataWrapper/objectAtIndexedSubscript:
func (s SLSScreenTelemetryResultsSnapshotDataWrapper) ObjectAtIndexedSubscript(subscript uint64) objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("objectAtIndexedSubscript:"), subscript)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/SLSScreenTelemetryResultsSnapshotDataWrapper/panel
func (s SLSScreenTelemetryResultsSnapshotDataWrapper) Panel() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("panel"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/SLSScreenTelemetryResultsSnapshotDataWrapper/rawData
func (s SLSScreenTelemetryResultsSnapshotDataWrapper) RawData() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s.ID, objc.Sel("rawData"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSScreenTelemetryResultsSnapshotDataWrapper/rowCount
func (s SLSScreenTelemetryResultsSnapshotDataWrapper) RowCount() uint64 {
	rv := objc.Send[uint64](s.ID, objc.Sel("rowCount"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSScreenTelemetryResultsSnapshotDataWrapper/timestamp
func (s SLSScreenTelemetryResultsSnapshotDataWrapper) Timestamp() float64 {
	rv := objc.Send[float64](s.ID, objc.Sel("timestamp"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSScreenTelemetryResultsSnapshotDataWrapper/initWithObject:
func (s SLSScreenTelemetryResultsSnapshotDataWrapper) InitWithObject(object objectivec.IObject) SLSScreenTelemetryResultsSnapshotDataWrapper {
	rv := objc.Send[SLSScreenTelemetryResultsSnapshotDataWrapper](s.ID, objc.Sel("initWithObject:"), object)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSScreenTelemetryResultsSnapshotDataWrapper/initWithXPCObject:
func (s SLSScreenTelemetryResultsSnapshotDataWrapper) InitWithXPCObject(xPCObject objectivec.IObject) SLSScreenTelemetryResultsSnapshotDataWrapper {
	rv := objc.Send[SLSScreenTelemetryResultsSnapshotDataWrapper](s.ID, objc.Sel("initWithXPCObject:"), xPCObject)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSScreenTelemetryResultsSnapshotDataWrapper/wrapperWithObject:
func (_SLSScreenTelemetryResultsSnapshotDataWrapperClass SLSScreenTelemetryResultsSnapshotDataWrapperClass) WrapperWithObject(object objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_SLSScreenTelemetryResultsSnapshotDataWrapperClass.class), objc.Sel("wrapperWithObject:"), object)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/SLSScreenTelemetryResultsSnapshotDataWrapper/wrapperWithXPCObject:
func (_SLSScreenTelemetryResultsSnapshotDataWrapperClass SLSScreenTelemetryResultsSnapshotDataWrapperClass) WrapperWithXPCObject(xPCObject objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_SLSScreenTelemetryResultsSnapshotDataWrapperClass.class), objc.Sel("wrapperWithXPCObject:"), xPCObject)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/SLSScreenTelemetryResultsSnapshotDataWrapper/data
func (s SLSScreenTelemetryResultsSnapshotDataWrapper) Data() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("data"))
	return objectivec.Object{ID: rv}
}
