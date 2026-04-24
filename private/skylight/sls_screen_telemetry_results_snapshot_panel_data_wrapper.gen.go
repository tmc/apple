// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLSScreenTelemetryResultsSnapshotPanelDataWrapper] class.
var (
	_SLSScreenTelemetryResultsSnapshotPanelDataWrapperClass     SLSScreenTelemetryResultsSnapshotPanelDataWrapperClass
	_SLSScreenTelemetryResultsSnapshotPanelDataWrapperClassOnce sync.Once
)

func getSLSScreenTelemetryResultsSnapshotPanelDataWrapperClass() SLSScreenTelemetryResultsSnapshotPanelDataWrapperClass {
	_SLSScreenTelemetryResultsSnapshotPanelDataWrapperClassOnce.Do(func() {
		_SLSScreenTelemetryResultsSnapshotPanelDataWrapperClass = SLSScreenTelemetryResultsSnapshotPanelDataWrapperClass{class: objc.GetClass("SLSScreenTelemetryResultsSnapshotPanelDataWrapper")}
	})
	return _SLSScreenTelemetryResultsSnapshotPanelDataWrapperClass
}

// GetSLSScreenTelemetryResultsSnapshotPanelDataWrapperClass returns the class object for SLSScreenTelemetryResultsSnapshotPanelDataWrapper.
func GetSLSScreenTelemetryResultsSnapshotPanelDataWrapperClass() SLSScreenTelemetryResultsSnapshotPanelDataWrapperClass {
	return getSLSScreenTelemetryResultsSnapshotPanelDataWrapperClass()
}

type SLSScreenTelemetryResultsSnapshotPanelDataWrapperClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLSScreenTelemetryResultsSnapshotPanelDataWrapperClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLSScreenTelemetryResultsSnapshotPanelDataWrapperClass) Alloc() SLSScreenTelemetryResultsSnapshotPanelDataWrapper {
	rv := objc.Send[SLSScreenTelemetryResultsSnapshotPanelDataWrapper](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLSScreenTelemetryResultsSnapshotPanelDataWrapper.AvgB]
//   - [SLSScreenTelemetryResultsSnapshotPanelDataWrapper.AvgG]
//   - [SLSScreenTelemetryResultsSnapshotPanelDataWrapper.AvgR]
//   - [SLSScreenTelemetryResultsSnapshotPanelDataWrapper.Data]
//   - [SLSScreenTelemetryResultsSnapshotPanelDataWrapper.RawData]
//   - [SLSScreenTelemetryResultsSnapshotPanelDataWrapper.TotalPixelCount]
//   - [SLSScreenTelemetryResultsSnapshotPanelDataWrapper.InitWithObject]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSScreenTelemetryResultsSnapshotPanelDataWrapper
type SLSScreenTelemetryResultsSnapshotPanelDataWrapper struct {
	objectivec.Object
}

// SLSScreenTelemetryResultsSnapshotPanelDataWrapperFromID constructs a [SLSScreenTelemetryResultsSnapshotPanelDataWrapper] from an objc.ID.
func SLSScreenTelemetryResultsSnapshotPanelDataWrapperFromID(id objc.ID) SLSScreenTelemetryResultsSnapshotPanelDataWrapper {
	return SLSScreenTelemetryResultsSnapshotPanelDataWrapper{objectivec.Object{ID: id}}
}

// Ensure SLSScreenTelemetryResultsSnapshotPanelDataWrapper implements ISLSScreenTelemetryResultsSnapshotPanelDataWrapper.
var _ ISLSScreenTelemetryResultsSnapshotPanelDataWrapper = SLSScreenTelemetryResultsSnapshotPanelDataWrapper{}

// An interface definition for the [SLSScreenTelemetryResultsSnapshotPanelDataWrapper] class.
//
// # Methods
//
//   - [ISLSScreenTelemetryResultsSnapshotPanelDataWrapper.AvgB]
//   - [ISLSScreenTelemetryResultsSnapshotPanelDataWrapper.AvgG]
//   - [ISLSScreenTelemetryResultsSnapshotPanelDataWrapper.AvgR]
//   - [ISLSScreenTelemetryResultsSnapshotPanelDataWrapper.Data]
//   - [ISLSScreenTelemetryResultsSnapshotPanelDataWrapper.RawData]
//   - [ISLSScreenTelemetryResultsSnapshotPanelDataWrapper.TotalPixelCount]
//   - [ISLSScreenTelemetryResultsSnapshotPanelDataWrapper.InitWithObject]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSScreenTelemetryResultsSnapshotPanelDataWrapper
type ISLSScreenTelemetryResultsSnapshotPanelDataWrapper interface {
	objectivec.IObject

	// Topic: Methods

	AvgB() float32
	AvgG() float32
	AvgR() float32
	Data() objectivec.IObject
	RawData() unsafe.Pointer
	TotalPixelCount() float32
	InitWithObject(object objectivec.IObject) SLSScreenTelemetryResultsSnapshotPanelDataWrapper
}

// Init initializes the instance.
func (s SLSScreenTelemetryResultsSnapshotPanelDataWrapper) Init() SLSScreenTelemetryResultsSnapshotPanelDataWrapper {
	rv := objc.Send[SLSScreenTelemetryResultsSnapshotPanelDataWrapper](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLSScreenTelemetryResultsSnapshotPanelDataWrapper) Autorelease() SLSScreenTelemetryResultsSnapshotPanelDataWrapper {
	rv := objc.Send[SLSScreenTelemetryResultsSnapshotPanelDataWrapper](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLSScreenTelemetryResultsSnapshotPanelDataWrapper creates a new SLSScreenTelemetryResultsSnapshotPanelDataWrapper instance.
func NewSLSScreenTelemetryResultsSnapshotPanelDataWrapper() SLSScreenTelemetryResultsSnapshotPanelDataWrapper {
	class := getSLSScreenTelemetryResultsSnapshotPanelDataWrapperClass()
	rv := objc.Send[SLSScreenTelemetryResultsSnapshotPanelDataWrapper](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSScreenTelemetryResultsSnapshotPanelDataWrapper/initWithObject:
func NewSLSScreenTelemetryResultsSnapshotPanelDataWrapperWithObject(object objectivec.IObject) SLSScreenTelemetryResultsSnapshotPanelDataWrapper {
	instance := getSLSScreenTelemetryResultsSnapshotPanelDataWrapperClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithObject:"), object)
	return SLSScreenTelemetryResultsSnapshotPanelDataWrapperFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSScreenTelemetryResultsSnapshotPanelDataWrapper/avgB
func (s SLSScreenTelemetryResultsSnapshotPanelDataWrapper) AvgB() float32 {
	rv := objc.Send[float32](s.ID, objc.Sel("avgB"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSScreenTelemetryResultsSnapshotPanelDataWrapper/avgG
func (s SLSScreenTelemetryResultsSnapshotPanelDataWrapper) AvgG() float32 {
	rv := objc.Send[float32](s.ID, objc.Sel("avgG"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSScreenTelemetryResultsSnapshotPanelDataWrapper/avgR
func (s SLSScreenTelemetryResultsSnapshotPanelDataWrapper) AvgR() float32 {
	rv := objc.Send[float32](s.ID, objc.Sel("avgR"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSScreenTelemetryResultsSnapshotPanelDataWrapper/rawData
func (s SLSScreenTelemetryResultsSnapshotPanelDataWrapper) RawData() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s.ID, objc.Sel("rawData"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSScreenTelemetryResultsSnapshotPanelDataWrapper/totalPixelCount
func (s SLSScreenTelemetryResultsSnapshotPanelDataWrapper) TotalPixelCount() float32 {
	rv := objc.Send[float32](s.ID, objc.Sel("totalPixelCount"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSScreenTelemetryResultsSnapshotPanelDataWrapper/initWithObject:
func (s SLSScreenTelemetryResultsSnapshotPanelDataWrapper) InitWithObject(object objectivec.IObject) SLSScreenTelemetryResultsSnapshotPanelDataWrapper {
	rv := objc.Send[SLSScreenTelemetryResultsSnapshotPanelDataWrapper](s.ID, objc.Sel("initWithObject:"), object)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSScreenTelemetryResultsSnapshotPanelDataWrapper/wrapperWithObject:
func (_SLSScreenTelemetryResultsSnapshotPanelDataWrapperClass SLSScreenTelemetryResultsSnapshotPanelDataWrapperClass) WrapperWithObject(object objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_SLSScreenTelemetryResultsSnapshotPanelDataWrapperClass.class), objc.Sel("wrapperWithObject:"), object)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/SLSScreenTelemetryResultsSnapshotPanelDataWrapper/data
func (s SLSScreenTelemetryResultsSnapshotPanelDataWrapper) Data() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("data"))
	return objectivec.Object{ID: rv}
}
