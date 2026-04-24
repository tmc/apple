// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLSScreenTelemetryResultsSnapshotZoneRowDataWrapper] class.
var (
	_SLSScreenTelemetryResultsSnapshotZoneRowDataWrapperClass     SLSScreenTelemetryResultsSnapshotZoneRowDataWrapperClass
	_SLSScreenTelemetryResultsSnapshotZoneRowDataWrapperClassOnce sync.Once
)

func getSLSScreenTelemetryResultsSnapshotZoneRowDataWrapperClass() SLSScreenTelemetryResultsSnapshotZoneRowDataWrapperClass {
	_SLSScreenTelemetryResultsSnapshotZoneRowDataWrapperClassOnce.Do(func() {
		_SLSScreenTelemetryResultsSnapshotZoneRowDataWrapperClass = SLSScreenTelemetryResultsSnapshotZoneRowDataWrapperClass{class: objc.GetClass("SLSScreenTelemetryResultsSnapshotZoneRowDataWrapper")}
	})
	return _SLSScreenTelemetryResultsSnapshotZoneRowDataWrapperClass
}

// GetSLSScreenTelemetryResultsSnapshotZoneRowDataWrapperClass returns the class object for SLSScreenTelemetryResultsSnapshotZoneRowDataWrapper.
func GetSLSScreenTelemetryResultsSnapshotZoneRowDataWrapperClass() SLSScreenTelemetryResultsSnapshotZoneRowDataWrapperClass {
	return getSLSScreenTelemetryResultsSnapshotZoneRowDataWrapperClass()
}

type SLSScreenTelemetryResultsSnapshotZoneRowDataWrapperClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLSScreenTelemetryResultsSnapshotZoneRowDataWrapperClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLSScreenTelemetryResultsSnapshotZoneRowDataWrapperClass) Alloc() SLSScreenTelemetryResultsSnapshotZoneRowDataWrapper {
	rv := objc.Send[SLSScreenTelemetryResultsSnapshotZoneRowDataWrapper](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLSScreenTelemetryResultsSnapshotZoneRowDataWrapper.ColumnCount]
//   - [SLSScreenTelemetryResultsSnapshotZoneRowDataWrapper.Data]
//   - [SLSScreenTelemetryResultsSnapshotZoneRowDataWrapper.ObjectAtIndexedSubscript]
//   - [SLSScreenTelemetryResultsSnapshotZoneRowDataWrapper.RawData]
//   - [SLSScreenTelemetryResultsSnapshotZoneRowDataWrapper.Row]
//   - [SLSScreenTelemetryResultsSnapshotZoneRowDataWrapper.InitWithObject]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSScreenTelemetryResultsSnapshotZoneRowDataWrapper
type SLSScreenTelemetryResultsSnapshotZoneRowDataWrapper struct {
	objectivec.Object
}

// SLSScreenTelemetryResultsSnapshotZoneRowDataWrapperFromID constructs a [SLSScreenTelemetryResultsSnapshotZoneRowDataWrapper] from an objc.ID.
func SLSScreenTelemetryResultsSnapshotZoneRowDataWrapperFromID(id objc.ID) SLSScreenTelemetryResultsSnapshotZoneRowDataWrapper {
	return SLSScreenTelemetryResultsSnapshotZoneRowDataWrapper{objectivec.Object{ID: id}}
}

// Ensure SLSScreenTelemetryResultsSnapshotZoneRowDataWrapper implements ISLSScreenTelemetryResultsSnapshotZoneRowDataWrapper.
var _ ISLSScreenTelemetryResultsSnapshotZoneRowDataWrapper = SLSScreenTelemetryResultsSnapshotZoneRowDataWrapper{}

// An interface definition for the [SLSScreenTelemetryResultsSnapshotZoneRowDataWrapper] class.
//
// # Methods
//
//   - [ISLSScreenTelemetryResultsSnapshotZoneRowDataWrapper.ColumnCount]
//   - [ISLSScreenTelemetryResultsSnapshotZoneRowDataWrapper.Data]
//   - [ISLSScreenTelemetryResultsSnapshotZoneRowDataWrapper.ObjectAtIndexedSubscript]
//   - [ISLSScreenTelemetryResultsSnapshotZoneRowDataWrapper.RawData]
//   - [ISLSScreenTelemetryResultsSnapshotZoneRowDataWrapper.Row]
//   - [ISLSScreenTelemetryResultsSnapshotZoneRowDataWrapper.InitWithObject]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSScreenTelemetryResultsSnapshotZoneRowDataWrapper
type ISLSScreenTelemetryResultsSnapshotZoneRowDataWrapper interface {
	objectivec.IObject

	// Topic: Methods

	ColumnCount() uint64
	Data() objectivec.IObject
	ObjectAtIndexedSubscript(subscript uint64) objectivec.IObject
	RawData() unsafe.Pointer
	Row() uint64
	InitWithObject(object objectivec.IObject) SLSScreenTelemetryResultsSnapshotZoneRowDataWrapper
}

// Init initializes the instance.
func (s SLSScreenTelemetryResultsSnapshotZoneRowDataWrapper) Init() SLSScreenTelemetryResultsSnapshotZoneRowDataWrapper {
	rv := objc.Send[SLSScreenTelemetryResultsSnapshotZoneRowDataWrapper](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLSScreenTelemetryResultsSnapshotZoneRowDataWrapper) Autorelease() SLSScreenTelemetryResultsSnapshotZoneRowDataWrapper {
	rv := objc.Send[SLSScreenTelemetryResultsSnapshotZoneRowDataWrapper](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLSScreenTelemetryResultsSnapshotZoneRowDataWrapper creates a new SLSScreenTelemetryResultsSnapshotZoneRowDataWrapper instance.
func NewSLSScreenTelemetryResultsSnapshotZoneRowDataWrapper() SLSScreenTelemetryResultsSnapshotZoneRowDataWrapper {
	class := getSLSScreenTelemetryResultsSnapshotZoneRowDataWrapperClass()
	rv := objc.Send[SLSScreenTelemetryResultsSnapshotZoneRowDataWrapper](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSScreenTelemetryResultsSnapshotZoneRowDataWrapper/initWithObject:
func NewSLSScreenTelemetryResultsSnapshotZoneRowDataWrapperWithObject(object objectivec.IObject) SLSScreenTelemetryResultsSnapshotZoneRowDataWrapper {
	instance := getSLSScreenTelemetryResultsSnapshotZoneRowDataWrapperClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithObject:"), object)
	return SLSScreenTelemetryResultsSnapshotZoneRowDataWrapperFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSScreenTelemetryResultsSnapshotZoneRowDataWrapper/columnCount
func (s SLSScreenTelemetryResultsSnapshotZoneRowDataWrapper) ColumnCount() uint64 {
	rv := objc.Send[uint64](s.ID, objc.Sel("columnCount"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSScreenTelemetryResultsSnapshotZoneRowDataWrapper/objectAtIndexedSubscript:
func (s SLSScreenTelemetryResultsSnapshotZoneRowDataWrapper) ObjectAtIndexedSubscript(subscript uint64) objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("objectAtIndexedSubscript:"), subscript)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/SLSScreenTelemetryResultsSnapshotZoneRowDataWrapper/rawData
func (s SLSScreenTelemetryResultsSnapshotZoneRowDataWrapper) RawData() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s.ID, objc.Sel("rawData"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSScreenTelemetryResultsSnapshotZoneRowDataWrapper/row
func (s SLSScreenTelemetryResultsSnapshotZoneRowDataWrapper) Row() uint64 {
	rv := objc.Send[uint64](s.ID, objc.Sel("row"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSScreenTelemetryResultsSnapshotZoneRowDataWrapper/initWithObject:
func (s SLSScreenTelemetryResultsSnapshotZoneRowDataWrapper) InitWithObject(object objectivec.IObject) SLSScreenTelemetryResultsSnapshotZoneRowDataWrapper {
	rv := objc.Send[SLSScreenTelemetryResultsSnapshotZoneRowDataWrapper](s.ID, objc.Sel("initWithObject:"), object)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSScreenTelemetryResultsSnapshotZoneRowDataWrapper/wrapperWithObject:
func (_SLSScreenTelemetryResultsSnapshotZoneRowDataWrapperClass SLSScreenTelemetryResultsSnapshotZoneRowDataWrapperClass) WrapperWithObject(object objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_SLSScreenTelemetryResultsSnapshotZoneRowDataWrapperClass.class), objc.Sel("wrapperWithObject:"), object)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/SLSScreenTelemetryResultsSnapshotZoneRowDataWrapper/data
func (s SLSScreenTelemetryResultsSnapshotZoneRowDataWrapper) Data() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("data"))
	return objectivec.Object{ID: rv}
}
